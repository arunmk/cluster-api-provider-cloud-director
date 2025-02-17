/*
   Copyright 2021 VMware, Inc.
   SPDX-License-Identifier: Apache-2.0
*/

package controllers

import (
	"bytes"
	"context"
	_ "embed" // this needs go 1.16+
	b64 "encoding/base64"
	"fmt"
	"github.com/pkg/errors"
	"github.com/replicatedhq/troubleshoot/pkg/redact"
	infrav1 "github.com/vmware/cluster-api-provider-cloud-director/api/v1alpha4"
	"github.com/vmware/cluster-api-provider-cloud-director/pkg/vcdclient"
	"github.com/vmware/go-vcloud-director/v2/govcd"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/klog"
	"net/http"
	clusterv1 "sigs.k8s.io/cluster-api/api/v1alpha4"
	"sigs.k8s.io/cluster-api/util"
	"sigs.k8s.io/cluster-api/util/annotations"
	"sigs.k8s.io/cluster-api/util/conditions"
	"sigs.k8s.io/cluster-api/util/patch"
	"sigs.k8s.io/cluster-api/util/predicates"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/source"
	"strconv"
	"strings"
	"time"
)

// The following `embed` directives read the file in the mentioned path and copy the content into the declared variable.
// These variables need to be global within the package.
//go:embed cluster_scripts/cloud_init_control_plane.yaml
var controlPlaneCloudInitScriptTemplate string

//go:embed cluster_scripts/cloud_init_node.yaml
var nodeCloudInitScriptTemplate string

// VCDMachineReconciler reconciles a VCDMachine object
type VCDMachineReconciler struct {
	client.Client
	//Scheme    *runtime.Scheme
	VcdClient *vcdclient.Client
}

//+kubebuilder:rbac:groups=infrastructure.cluster.x-k8s.io,resources=vcdmachines,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=infrastructure.cluster.x-k8s.io,resources=vcdmachines/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=infrastructure.cluster.x-k8s.io,resources=vcdmachines/finalizers,verbs=update
func (r *VCDMachineReconciler) Reconcile(ctx context.Context, req ctrl.Request) (_ ctrl.Result, rerr error) {
	log := ctrl.LoggerFrom(ctx)

	// Fetch the VCDMachine instance.
	vcdMachine := &infrav1.VCDMachine{}
	if err := r.Client.Get(ctx, req.NamespacedName, vcdMachine); err != nil {
		if apierrors.IsNotFound(err) {
			return ctrl.Result{}, nil
		}
		return ctrl.Result{}, err
	}

	machine, err := util.GetOwnerMachine(ctx, r.Client, vcdMachine.ObjectMeta)
	if err != nil {
		return ctrl.Result{}, err
	}
	if machine == nil {
		log.Info("Waiting for Machine Controller to set OwnerRef on VCDMachine")
		return ctrl.Result{}, nil
	}

	log = log.WithValues("machine", machine.Name)

	// Fetch the Cluster from k8s etcd.
	cluster, err := util.GetClusterFromMetadata(ctx, r.Client, machine.ObjectMeta)
	if err != nil {
		log.Info("VCDMachine owner Machine is missing cluster label or cluster does not exist")
		return ctrl.Result{}, err
	}
	if cluster == nil {
		log.Info("Please associate this machine with a cluster using the label", "label", clusterv1.ClusterLabelName)
		return ctrl.Result{}, nil
	}

	log = log.WithValues("cluster", cluster.Name)

	// Return early if the object or Cluster is paused.
	if annotations.IsPaused(cluster, vcdMachine) {
		log.Info("Reconciliation is paused for this object")
		return ctrl.Result{}, nil
	}

	machineBeingDeleted := !vcdMachine.ObjectMeta.DeletionTimestamp.IsZero()

	// Fetch the VCD Cluster.
	vcdCluster := &infrav1.VCDCluster{}
	vcdClusterName := client.ObjectKey{
		Namespace: vcdMachine.Namespace,
		Name:      cluster.Spec.InfrastructureRef.Name,
	}
	if err := r.Client.Get(ctx, vcdClusterName, vcdCluster); err != nil {
		log.Info("VCDCluster is not available yet")
		if !machineBeingDeleted {
			return ctrl.Result{}, nil
		} else {
			log.Info("Continuing to delete the VCDMachine, since deletion timestamp is set")
		}
	}

	// Initialize the patch helper
	patchHelper, err := patch.NewHelper(vcdMachine, r)
	if err != nil {
		return ctrl.Result{}, err
	}
	// Always attempt to Patch the VCDMachine object and status after each reconciliation.
	defer func() {
		if err := patchVCDMachine(ctx, patchHelper, vcdMachine); err != nil {
			log.Error(err, "Failed to patch VCDMachine")
			if rerr == nil {
				rerr = err
			}
		}
	}()

	// Add finalizer first if not exist to avoid the race condition between init and delete
	if !controllerutil.ContainsFinalizer(vcdMachine, infrav1.MachineFinalizer) {
		controllerutil.AddFinalizer(vcdMachine, infrav1.MachineFinalizer)
		return ctrl.Result{}, nil
	}

	// If the machine is not being deleted, check if the infrastructure is ready. If not ready, return and wait for
	// the cluster object to be updated
	if !machineBeingDeleted && !cluster.Status.InfrastructureReady {
		log.Info("Waiting for VCDCluster Controller to create cluster infrastructure")
		conditions.MarkFalse(vcdMachine, infrav1.ContainerProvisionedCondition,
			infrav1.WaitingForClusterInfrastructureReason, clusterv1.ConditionSeverityInfo, "")
		return ctrl.Result{}, nil
	}

	// Handle deleted machines
	if machineBeingDeleted {
		return r.reconcileDelete(ctx, cluster, machine, vcdMachine, vcdCluster)
	}

	// Handle non-deleted machines
	return r.reconcileNormal(ctx, cluster, machine, vcdMachine, vcdCluster)
}

func patchVCDMachine(ctx context.Context, patchHelper *patch.Helper, vcdMachine *infrav1.VCDMachine) error {
	conditions.SetSummary(vcdMachine,
		conditions.WithConditions(
			infrav1.ContainerProvisionedCondition,
			infrav1.BootstrapExecSucceededCondition,
		),
		conditions.WithStepCounterIf(vcdMachine.ObjectMeta.DeletionTimestamp.IsZero()),
	)

	return patchHelper.Patch(
		ctx,
		vcdMachine,
		patch.WithOwnedConditions{Conditions: []clusterv1.ConditionType{
			clusterv1.ReadyCondition,
			infrav1.ContainerProvisionedCondition,
			infrav1.BootstrapExecSucceededCondition,
		}},
	)
}

const (
	NetworkConfiguration                   = "guestinfo.postcustomization.networkconfiguration.status"
	KubeadmInit                            = "guestinfo.postcustomization.kubeinit.status"
	KubectlApplyCni                        = "guestinfo.postcustomization.kubectl.cni.install.status"
	KubectlApplyCpi                        = "guestinfo.postcustomization.kubectl.cpi.install.status"
	KubectlApplyCsi                        = "guestinfo.postcustomization.kubectl.csi.install.status"
	KubeadmTokenGenerate                   = "guestinfo.postcustomization.kubeadm.token.generate.status"
	KubeadmNodeJoin                        = "guestinfo.postcustomization.kubeadm.node.join.status"
	PostCustomizationScriptExecutionStatus = "guestinfo.post_customization_script_execution_status"
	PostCustomizationScriptFailureReason   = "guestinfo.post_customization_script_execution_failure_reason"
)

var controlPlanePostCustPhases = []string{
	NetworkConfiguration,
	KubeadmInit,
	KubectlApplyCni,
	KubectlApplyCpi,
	KubectlApplyCsi,
	KubeadmTokenGenerate,
}

var joinPostCustPhases = []string{
	NetworkConfiguration,
	KubeadmNodeJoin,
}

func removeFromSlice(remove string, arr []string) []string {
	for ind, str := range arr {
		if str == remove {
			return append(arr[:ind], arr[ind+1:]...)
		}
	}
	return arr
}

func strInSlice(findStr string, arr []string) bool {
	for _, str := range arr {
		if str == findStr {
			return true
		}
	}
	return false
}

const phaseSecondTimeout = 600

func redactCloudInit(cloudInitYaml string, path []string) (string, error) {
	yamlRunner := redact.NewYamlRedactor(strings.Join(path, "."), "", "cloudInitRedactor")
	outReader := yamlRunner.Redact(bytes.NewReader([]byte(cloudInitYaml)), "")
	gotBytes, err := ioutil.ReadAll(outReader)
	if err != nil {
		return cloudInitYaml, fmt.Errorf("failed to read redacted yaml output : %v", err)
	}
	return string(gotBytes), nil

}

func (r *VCDMachineReconciler) waitForPostCustomizationPhase(ctx context.Context, workloadVCDClient *vcdclient.Client, vm *govcd.VM, phase string) error {
	startTime := time.Now()
	possibleStatuses := []string{"", "in_progress", "successful"}
	currentStatus := possibleStatuses[0]
	for {
		if err := vm.Refresh(); err != nil {
			return errors.Wrapf(err, "unable to refresh vm [%s]: [%v]", vm.VM.Name, err)
		}
		newStatus, err := workloadVCDClient.GetExtraConfigValue(vm, phase)
		if err != nil {
			return errors.Wrapf(err, "unable to get extra config value for key [%s] for vm: [%s]: [%v]",
				phase, vm.VM.Name, err)
		}
		if !strInSlice(newStatus, possibleStatuses) {
			return errors.Wrapf(err, "invalid postcustomiation phase: [%s] for key [%s] for vm [%s]",
				newStatus, phase, vm.VM.Name)
		}
		if newStatus != currentStatus {
			possibleStatuses = removeFromSlice(currentStatus, possibleStatuses)
			currentStatus = newStatus
		}
		if newStatus == possibleStatuses[len(possibleStatuses)-1] { // successful status
			return nil
		}

		// catch intermediate script execution failure
		scriptExecutionStatus, err := workloadVCDClient.GetExtraConfigValue(vm, PostCustomizationScriptExecutionStatus)
		if err != nil {
			return errors.Wrapf(err, "unable to get extra config value for key [%s] for vm: [%s]: [%v]",
				PostCustomizationScriptExecutionStatus, vm.VM.Name, err)
		}
		if scriptExecutionStatus != "" {
			execStatus, err := strconv.Atoi(scriptExecutionStatus)
			if err != nil {
				return errors.Wrapf(err, "unable to convert script execution status [%s] to int: [%v]",
					scriptExecutionStatus, err)
			}
			if execStatus != 0 {
				scriptExecutionFailureReason, err := workloadVCDClient.GetExtraConfigValue(vm, PostCustomizationScriptFailureReason)
				if err != nil {
					return errors.Wrapf(err, "unable to get extra config value for key [%s] for vm, "+
						"(script execution status [%d]): [%s]: [%v]",
						PostCustomizationScriptFailureReason, execStatus, vm.VM.Name, err)
				}
				return fmt.Errorf("script failed with status [%d] and reason [%s]", execStatus, scriptExecutionFailureReason)
			}
		}

		if seconds := int(time.Since(startTime) / time.Second); seconds > phaseSecondTimeout {
			return fmt.Errorf("time for postcustomization status [%s] exceeded timeout [%d]",
				phase, phaseSecondTimeout)
		}
		time.Sleep(10 * time.Second)
	}

}

func (r *VCDMachineReconciler) reconcileNodeStatusInRDE(ctx context.Context, rdeID string, nodeName string, status string,
	workloadVCDClient *vcdclient.Client) error {

	if rdeID == "" || strings.HasPrefix(rdeID, NoRdePrefix) {
		return NewNoRDEError("RDE ID is empty or generated; hence will not be updated")
	}

	updatePatch := make(map[string]interface{})
	_, capvcdEntity, err := workloadVCDClient.GetCAPVCDEntity(ctx, rdeID)
	if err != nil {
		return fmt.Errorf("failed to get CAPVCD entity with ID [%s] to sync node details for machine [%s]: [%v]", rdeID, nodeName, err)
	}
	nodeStatusMap := capvcdEntity.Status.NodeStatus
	if nodeStatus, ok := nodeStatusMap[nodeName]; ok && nodeStatus == status {
		// no update needed
		return nil
	}
	if nodeStatusMap == nil {
		nodeStatusMap = make(map[string]string)
	}
	nodeStatusMap[nodeName] = status
	updatePatch["Status.NodeStatus"] = nodeStatusMap

	// update defined entity
	updatedRDE, err := workloadVCDClient.PatchRDE(ctx, updatePatch, rdeID)
	if err != nil {
		return fmt.Errorf("failed to update defined entity with ID [%s] with node status for VCDMachine [%s]: [%v]", rdeID, nodeName, err)
	}
	if updatedRDE.State != RDEStatusResolved {
		// try to resolve the defined entity
		entityState, resp, err := workloadVCDClient.ApiClient.DefinedEntityApi.ResolveDefinedEntity(ctx, updatedRDE.Id)
		if err != nil {
			return fmt.Errorf("failed to resolve defined entity with ID [%s] for cluster [%s]", updatedRDE.Id, updatedRDE.Name)
		}
		if resp.StatusCode != http.StatusOK {
			return fmt.Errorf("error while resolving defined entity with ID [%s] for cluster [%s] with message: [%s]", updatedRDE.Id, updatedRDE.Name, entityState.Message)
		}

		if entityState.State != RDEStatusResolved {
			return fmt.Errorf("defined entity resolution failed for RDE with ID [%s] for cluster [%s] with message: [%s]", updatedRDE.Id, updatedRDE.Name, entityState.Message)

		}
	}
	return nil
}

func (r *VCDMachineReconciler) reconcileNormal(ctx context.Context, cluster *clusterv1.Cluster,
	machine *clusterv1.Machine, vcdMachine *infrav1.VCDMachine, vcdCluster *infrav1.VCDCluster) (res ctrl.Result, retErr error) {

	log := ctrl.LoggerFrom(ctx, "machine", machine.Name, "cluster", vcdCluster.Name)

	workloadVCDClient, err := vcdclient.NewVCDClientFromSecrets(vcdCluster.Spec.Site, vcdCluster.Spec.Org,
		vcdCluster.Spec.Ovdc, vcdCluster.Name, vcdCluster.Spec.OvdcNetwork, r.VcdClient.IPAMSubnet,
		r.VcdClient.VcdAuthConfig.UserOrg, vcdCluster.Spec.UserCredentialsContext.Username,
		vcdCluster.Spec.UserCredentialsContext.Password, vcdCluster.Spec.UserCredentialsContext.RefreshToken,
		true, vcdCluster.Status.InfraId, r.VcdClient.OneArm, 0, 0, r.VcdClient.TCPPort,
		true, "", r.VcdClient.CsiVersion, r.VcdClient.CpiVersion, r.VcdClient.CniVersion,
		r.VcdClient.CAPVCDVersion)
	if err != nil {
		return ctrl.Result{}, errors.Wrapf(err, "Unable to create VCD client to reconcile infrastructure for the Machine [%s]", machine.Name)
	}

	if vcdMachine.Spec.ProviderID != nil {
		err := r.reconcileNodeStatusInRDE(ctx, vcdCluster.Status.InfraId, machine.Name, machine.Status.Phase,
			workloadVCDClient)
		if err != nil {
			if _, ok := err.(*NoRDEError); ok {
				log.V(3).Info("RDE NOT set up to track this cluster.",
					"infraID", vcdCluster.Status.InfraId)
			} else {
				log.Error(err, "Error during RDE reconciliation of the Node status")
			}
		}
		vcdMachine.Status.Ready = true
		conditions.MarkTrue(vcdMachine, infrav1.ContainerProvisionedCondition)
		return ctrl.Result{}, nil
	}

	err = r.reconcileNodeStatusInRDE(ctx, vcdCluster.Status.InfraId, machine.Name, machine.Status.Phase,
		workloadVCDClient)
	if err != nil {
		if _, ok := err.(*NoRDEError); ok {
			log.V(3).Info("RDE NOT set up to track this cluster.", "infraID", vcdCluster.Status.InfraId)
		} else {
			log.Error(err, "Error during RDE reconciliation of the Node status")
		}
	}

	if machine.Spec.Bootstrap.DataSecretName == nil {
		if !util.IsControlPlaneMachine(machine) && !conditions.IsTrue(cluster,
			clusterv1.ControlPlaneInitializedCondition) {

			log.Info("Waiting for the control plane to be initialized")
			conditions.MarkFalse(vcdMachine, infrav1.ContainerProvisionedCondition,
				clusterv1.WaitingForControlPlaneAvailableReason, clusterv1.ConditionSeverityInfo, "")
			return ctrl.Result{}, nil
		}

		log.Info("Waiting for the Bootstrap provider controller to set bootstrap data")
		conditions.MarkFalse(vcdMachine, infrav1.ContainerProvisionedCondition,
			infrav1.WaitingForBootstrapDataReason, clusterv1.ConditionSeverityInfo, "")
		return ctrl.Result{}, nil
	}

	patchHelper, err := patch.NewHelper(vcdMachine, r.Client)
	if err != nil {
		return ctrl.Result{}, errors.Wrapf(err, "Error patching VCDMachine [%s] of cluster [%s]", vcdMachine.Name, vcdCluster.Name)
	}
	conditions.MarkTrue(vcdMachine, infrav1.ContainerProvisionedCondition)

	if !conditions.Has(vcdMachine, infrav1.BootstrapExecSucceededCondition) {
		conditions.MarkFalse(vcdMachine, infrav1.BootstrapExecSucceededCondition,
			infrav1.BootstrappingReason, clusterv1.ConditionSeverityInfo, "")
		if err := patchVCDMachine(ctx, patchHelper, vcdMachine); err != nil {
			return ctrl.Result{}, errors.Wrapf(err, "Error patching VCDMachine [%s] of cluster [%s]", vcdMachine.Name, vcdCluster.Name)
		}
	}

	vdcManager := vcdclient.VdcManager{
		VdcName: workloadVCDClient.ClusterOVDCName,
		OrgName: workloadVCDClient.ClusterOrgName,
		Client:  workloadVCDClient,
		Vdc:     workloadVCDClient.Vdc,
	}

	// The vApp should have already been created, so this is more of a Get of the vApp
	vAppName := cluster.Name
	vApp, err := vdcManager.Vdc.GetVAppByName(vAppName, true)
	if err != nil {
		return ctrl.Result{}, errors.Wrapf(err, "Error provisioning infrastructure for the machine [%s] of the cluster [%s]", machine.Name, vcdCluster.Name)
	}

	bootstrapJinjaScript, err := r.getBootstrapData(ctx, machine)
	if err != nil {
		return ctrl.Result{}, errors.Wrapf(err, "Error retrieving bootstrap data for machine [%s] of the cluster [%s]",
			machine.Name, vcdCluster.Name)
	}
	// In a multimaster cluster, the initial control plane node runs `kubeadm init`; additional control plane nodes
	// run `kubeadm join`. The joining control planes run `kubeadm join`, so these nodes use the join script.
	// Although it is sufficient to just check if `kubeadm join` is in the bootstrap script, using the
	// isControlPlaneMachine function is a simpler operation, so this function is called first.
	useControlPlaneScript := true
	if !util.IsControlPlaneMachine(machine) || strings.Contains(bootstrapJinjaScript, "kubeadm join") {
		useControlPlaneScript = false
	}

	// We have control over the content in the guest Cloud Init Script. However, we can't control the content
	// in the Jinja script. Hence, do any fmt.Sprintf calls first and then merge the two scripts. This is cleaner
	// than calling custom sanitization libraries since there doesn't seem to be a clearly good one. Also cleaner
	// than handcrafting one.
	guestCloudInit := ""
	if !vcdMachine.Spec.Bootstrapped {
		guestCloudInitTemplate := controlPlaneCloudInitScriptTemplate
		if !useControlPlaneScript {
			guestCloudInitTemplate = nodeCloudInitScriptTemplate
		}

		switch {
		case useControlPlaneScript:
			orgUserStr := fmt.Sprintf("%s/%s", workloadVCDClient.VcdAuthConfig.UserOrg,
				workloadVCDClient.VcdAuthConfig.User)
			b64OrgUser := b64.StdEncoding.EncodeToString([]byte(orgUserStr))
			b64Password := b64.StdEncoding.EncodeToString([]byte(vcdCluster.Spec.UserCredentialsContext.Password))
			b64RefreshToken := b64.StdEncoding.EncodeToString([]byte(
				vcdCluster.Spec.UserCredentialsContext.RefreshToken))
			vcdHostFormatted := strings.Replace(vcdCluster.Spec.Site, "/", "\\/", -1)
			guestCloudInit = fmt.Sprintf(
				guestCloudInitTemplate,            // template script
				b64OrgUser,                        // base 64 org/username
				b64Password,                       // base64 password
				b64RefreshToken,                   // refresh token
				workloadVCDClient.CniVersion,      // cni version
				workloadVCDClient.CpiVersion,      // cpi version
				vcdHostFormatted,                  // vcd host
				workloadVCDClient.ClusterOrgName,  // org
				workloadVCDClient.ClusterOVDCName, // ovdc
				workloadVCDClient.NetworkName,     // network
				"",                                // vip subnet cidr - empty for now for CPI to select subnet
				vAppName,                          // vApp name
				workloadVCDClient.ClusterID,       // cluster id
				workloadVCDClient.CsiVersion,      // csi version
				vcdHostFormatted,                  // vcd host,
				workloadVCDClient.ClusterOrgName,  // org
				workloadVCDClient.ClusterOVDCName, // ovdc
				vAppName,                          // vApp
				workloadVCDClient.ClusterID,       // cluster id
				machine.Name,                      // vm host name
			)

		default:
			guestCloudInit = fmt.Sprintf(
				guestCloudInitTemplate, // template script
				machine.Name,           // vm host name
			)
		}
	}

	mergedCloudInitBytes, err := MergeJinjaToCloudInitScript(guestCloudInit, bootstrapJinjaScript)
	if err != nil {
		return ctrl.Result{}, errors.Wrapf(err,
			"Error merging bootstrap jinja script with the cloudInit script for [%s/%s] [%s]",
			vAppName, machine.Name, bootstrapJinjaScript)
	}

	redactedCloudInit := string(mergedCloudInitBytes)
	if util.IsControlPlaneMachine(machine) {
		// redact secrets
		// NOTE: the position of the key in cluster_scripts/cloud_init_control_plane.yaml is important as the following
		// code expects the secret to be the first element in write_files.
		redactedCloudInit, err = redactCloudInit(string(mergedCloudInitBytes), []string{"write_files", "0", "content"})
		if err != nil {
			log.Error(err, "failed to redact cloud init script")
		}
	}

	log.Info(fmt.Sprintf("Cloud init Script: [%s]", redactedCloudInit))

	vmExists := true
	vm, err := vApp.GetVMByName(machine.Name, true)
	if err != nil && err != govcd.ErrorEntityNotFound {
		return ctrl.Result{}, errors.Wrapf(err, "Error provisioning infrastructure for the machine; unable to query for VM [%s] in vApp [%s]",
			machine.Name, vAppName)
	} else if err == govcd.ErrorEntityNotFound {
		vmExists = false
	}
	if !vmExists {
		log.Info("Adding infra VM for the machine")
		err = vdcManager.AddNewVM(machine.Name, vApp.VApp.Name, 1,
			vcdMachine.Spec.Catalog, vcdMachine.Spec.Template, "",
			vcdMachine.Spec.ComputePolicy, "", false)
		if err != nil {
			return ctrl.Result{}, errors.Wrapf(err, "Error provisioning infrastructure for the machine; unable to create VM [%s] in vApp [%s]",
				machine.Name, vApp.VApp.Name)
		}
		vm, err = vApp.GetVMByName(machine.Name, true)
		if err != nil {
			return ctrl.Result{}, errors.Wrapf(err, "Error provisioning infrastructure for the machine; unable to find newly created VM [%s] in vApp [%s]",
				vm.VM.Name, vAppName)
		}
	}

	// set address in machine status
	if vm.VM == nil ||
		vm.VM.NetworkConnectionSection == nil ||
		len(vm.VM.NetworkConnectionSection.NetworkConnection) == 0 ||
		vm.VM.NetworkConnectionSection.NetworkConnection[0] == nil ||
		vm.VM.NetworkConnectionSection.NetworkConnection[0].IPAddress == "" {

		log.Error(nil, fmt.Sprintf("Requeuing...; failed to get the machine address of vm [%#v]", vm))
		return ctrl.Result{RequeueAfter: 5 * time.Second}, nil
	}

	machineAddress := vm.VM.NetworkConnectionSection.NetworkConnection[0].IPAddress
	vcdMachine.Status.Addresses = []clusterv1.MachineAddress{
		{
			Type:    clusterv1.MachineHostName,
			Address: vm.VM.Name,
		},
		{
			Type:    clusterv1.MachineInternalIP,
			Address: machineAddress,
		},
		{
			Type:    clusterv1.MachineExternalIP,
			Address: machineAddress,
		},
	}

	gateway := &vcdclient.GatewayManager{
		NetworkName:        workloadVCDClient.NetworkName,
		Client:             workloadVCDClient,
		GatewayRef:         workloadVCDClient.GatewayRef,
		NetworkBackingType: workloadVCDClient.NetworkBackingType,
	}

	// Update loadbalancer pool with the IP of the control plane node as a new member.
	// Note that this must be done before booting on the VM!
	if util.IsControlPlaneMachine(machine) {
		lbPoolName := vcdCluster.Name + "-" + vcdCluster.Status.InfraId + "-tcp"
		lbPoolRef, err := gateway.GetLoadBalancerPool(ctx, lbPoolName)
		if err != nil {
			return ctrl.Result{}, errors.Wrapf(err, "Error retrieving/updating load balancer pool [%s] for the "+
				"control plane machine [%s] of the cluster [%s]", lbPoolName, machine.Name, vcdCluster.Name)
		}
		controlPlaneIPs, err := gateway.GetLoadBalancerPoolMemberIPs(ctx, lbPoolRef)
		if err != nil {
			return ctrl.Result{}, errors.Wrapf(err,
				"Error retrieving/updating load balancer pool members [%s] for the "+
					"control plane machine [%s] of the cluster [%s]", lbPoolName, machine.Name, vcdCluster.Name)
		}

		updatedIPs := append(controlPlaneIPs, machineAddress)
		err = gateway.UpdateLoadBalancer(ctx, lbPoolName, updatedIPs, int32(6443))
		if err != nil {
			return ctrl.Result{}, errors.Wrapf(err,
				"Error updating the load balancer pool [%s] for the "+
					"control plane machine [%s] of the cluster [%s]", lbPoolName, machine.Name, vcdCluster.Name)
		}
		log.Info("Updated the load balancer pool with the control plane machine IP", "lbpool", lbPoolName)
	}

	vmStatus, err := vm.GetStatus()
	if err != nil {
		return ctrl.Result{},
			errors.Wrapf(err, "Error while provisioning the infrastructure VM for the machine [%s] of the cluster [%s]; failed to get status of vm", vm.VM.Name, vApp.VApp.Name)
	}

	if vmStatus != "POWERED_ON" {
		// try to power on the VM
		b64CloudInitScript := b64.StdEncoding.EncodeToString(mergedCloudInitBytes)
		keyVals := map[string]string{
			"guestinfo.userdata":          b64CloudInitScript,
			"guestinfo.userdata.encoding": "base64",
			"disk.enableUUID":             "1",
		}

		for key, val := range keyVals {
			err = workloadVCDClient.SetVmExtraConfigKeyValue(vm, key, val, true)
			if err != nil {
				return ctrl.Result{}, errors.Wrapf(err, "Error while enabling cloudinit on the machine [%s/%s]; unable to set vm extra config key [%s] for vm ",
					vcdCluster.Name, vm.VM.Name, key)
			}

			if err = vm.Refresh(); err != nil {
				return ctrl.Result{}, errors.Wrapf(err, "Error while enabling cloudinit on the machine [%s/%s]; unable to refresh vm", vcdCluster.Name, vm.VM.Name)
			}

			if err = vApp.Refresh(); err != nil {
				return ctrl.Result{}, errors.Wrapf(err, "Error while enabling cloudinit on the machine [%s/%s]; unable to refresh vapp", vAppName, vm.VM.Name)
			}

			log.Info(fmt.Sprintf("Configured the infra machine with variable [%s] to enable cloud-init", key))
		}

		task, err := vm.PowerOn()
		if err != nil {
			return ctrl.Result{}, errors.Wrapf(err, "Error while deploying infra for the machine [%s/%s]; unable to power on VM", vcdCluster.Name, vm.VM.Name)
		}
		if err = task.WaitTaskCompletion(); err != nil {
			return ctrl.Result{}, errors.Wrapf(err, "Error while deploying infra for the machine [%s/%s]; error waiting for VM power-on task completion", vcdCluster.Name, vm.VM.Name)
		}

		if err = vApp.Refresh(); err != nil {
			return ctrl.Result{}, errors.Wrapf(err, "Error while deploying infra for the machine [%s/%s]; unable to refresh vapp after VM power-on", vAppName, vm.VM.Name)
		}
	}
	if hasCloudInitFailedBefore, err := r.hasCloudInitExecutionFailedBefore(ctx, workloadVCDClient, vm); hasCloudInitFailedBefore {
		return ctrl.Result{}, errors.Wrapf(err, "Error bootstrapping the machine [%s/%s]; machine is probably in unreconciliable state", vAppName, vm.VM.Name)
	}

	// wait for each vm phase
	phases := controlPlanePostCustPhases
	if !useControlPlaneScript {
		phases = joinPostCustPhases
	}
	for _, phase := range phases {
		if err = vApp.Refresh(); err != nil {
			return ctrl.Result{}, errors.Wrapf(err, "Error while bootstrapping the machine [%s/%s]; unable to refresh vapp", vAppName, vm.VM.Name)
		}
		log.Info(fmt.Sprintf("Start: waiting for the bootstrapping phase [%s] to complete", phase))
		if err = r.waitForPostCustomizationPhase(ctx, workloadVCDClient, vm, phase); err != nil {
			log.Error(err, fmt.Sprintf("Error waiting for the bootstrapping phase [%s] to complete", phase))
			return ctrl.Result{}, errors.Wrapf(err, "Error while bootstrapping the machine [%s/%s]; unable to wait for post customization phase [%s]",
				vAppName, vm.VM.Name, phase)
		}
		log.Info(fmt.Sprintf("End: waiting for the bootstrapping phase [%s] to complete", phase))
	}

	log.Info("Successfully bootstrapped the machine")

	if err = vm.Refresh(); err != nil {
		return ctrl.Result{}, errors.Wrapf(err, "Unexpected error after the machine [%s/%s] is bootstrapped; unable to refresh vm", vAppName, vm.VM.Name)
	}
	if err = vApp.Refresh(); err != nil {
		return ctrl.Result{}, errors.Wrapf(err, "Unexpected error after the machine [%s/%s] is bootstrapped; unable to refresh vapp", vAppName, vm.VM.Name)
	}

	vcdMachine.Spec.Bootstrapped = true
	conditions.MarkTrue(vcdMachine, infrav1.BootstrapExecSucceededCondition)

	// Set ProviderID so the Cluster API Machine Controller can pull it
	providerID := fmt.Sprintf("%s://%s", infrav1.VCDProviderID, vm.VM.ID)
	vcdMachine.Spec.ProviderID = &providerID
	vcdMachine.Status.Ready = true
	conditions.MarkTrue(vcdMachine, infrav1.ContainerProvisionedCondition)
	err = r.reconcileNodeStatusInRDE(ctx, vcdCluster.Status.InfraId, machine.Name, machine.Status.Phase, workloadVCDClient)
	if err != nil {
		if _, ok := err.(*NoRDEError); ok {
			log.V(3).Info("RDE NOT set up to track this cluster.", "infraID", vcdCluster.Status.InfraId)
		} else {
			log.Error(err, "Error reconciling node status of the RDE",
				"RDEId", vcdCluster.Status.InfraId, "nodeStatus", machine.Status.Phase)
		}
	}

	return ctrl.Result{}, nil
}

func (r *VCDMachineReconciler) getBootstrapData(ctx context.Context, machine *clusterv1.Machine) (string, error) {
	log := ctrl.LoggerFrom(ctx)
	if machine.Spec.Bootstrap.DataSecretName == nil {
		return "", errors.New("error retrieving bootstrap data: linked Machine's bootstrap.dataSecretName is nil")
	}

	s := &corev1.Secret{}
	key := client.ObjectKey{Namespace: machine.GetNamespace(), Name: *machine.Spec.Bootstrap.DataSecretName}
	if err := r.Client.Get(ctx, key, s); err != nil {
		return "", errors.Wrapf(err, "failed to retrieve bootstrap data secret for VCDMachine %s/%s", machine.GetNamespace(), machine.GetName())
	}

	value, ok := s.Data["value"]
	if !ok {
		return "", errors.New("error retrieving bootstrap data: secret value key is missing")
	}

	log.Info(fmt.Sprintf("Auto-generated bootstrap script: [%s]", string(value)))

	return string(value), nil
}

func (r *VCDMachineReconciler) reconcileDelete(ctx context.Context, cluster *clusterv1.Cluster, machine *clusterv1.Machine,
	vcdMachine *infrav1.VCDMachine, vcdCluster *infrav1.VCDCluster) (ctrl.Result, error) {
	log := ctrl.LoggerFrom(ctx, "machine", machine.Name, "cluster", vcdCluster.Name)

	patchHelper, err := patch.NewHelper(vcdMachine, r.Client)
	if err != nil {
		return ctrl.Result{}, err
	}

	conditions.MarkFalse(vcdMachine, infrav1.ContainerProvisionedCondition,
		clusterv1.DeletingReason, clusterv1.ConditionSeverityInfo, "")
	if err := patchVCDMachine(ctx, patchHelper, vcdMachine); err != nil {
		return ctrl.Result{}, errors.Wrapf(err, "Failed to patch VCDMachine [%s/%s]", vcdCluster.Name, vcdMachine.Name)
	}

	if vcdCluster.Spec.Site == "" {
		controllerutil.RemoveFinalizer(vcdMachine, infrav1.MachineFinalizer)
		return ctrl.Result{}, nil
	}

	workloadVCDClient, err := vcdclient.NewVCDClientFromSecrets(vcdCluster.Spec.Site, vcdCluster.Spec.Org,
		vcdCluster.Spec.Ovdc, vcdCluster.Name, vcdCluster.Spec.OvdcNetwork, r.VcdClient.IPAMSubnet,
		r.VcdClient.VcdAuthConfig.UserOrg, vcdCluster.Spec.UserCredentialsContext.Username,
		vcdCluster.Spec.UserCredentialsContext.Password, vcdCluster.Spec.UserCredentialsContext.RefreshToken,
		true, vcdCluster.Status.InfraId, r.VcdClient.OneArm, 0, 0, r.VcdClient.TCPPort,
		true, "", r.VcdClient.CsiVersion, r.VcdClient.CpiVersion,
		r.VcdClient.CniVersion, r.VcdClient.CAPVCDVersion)
	if err != nil {
		return ctrl.Result{}, errors.Wrapf(err,
			"Error creating VCD client to reconcile the machine [%s/%s] deletion",
			vcdCluster.Name, vcdMachine.Name)
	}

	gateway := &vcdclient.GatewayManager{
		NetworkName:        workloadVCDClient.NetworkName,
		Client:             workloadVCDClient,
		GatewayRef:         workloadVCDClient.GatewayRef,
		NetworkBackingType: workloadVCDClient.NetworkBackingType,
	}
	if util.IsControlPlaneMachine(machine) {
		// remove the address from the lbpool
		log.Info("Deleting the control plane IP from the load balancer pool")
		lbPoolName := vcdCluster.Name + "-" + vcdCluster.Status.InfraId + "-tcp"
		lbPoolRef, err := gateway.GetLoadBalancerPool(ctx, lbPoolName)
		if err != nil && err != govcd.ErrorEntityNotFound {
			return ctrl.Result{}, errors.Wrapf(err, "Error while deleting the infra resources of the machine [%s/%s]; failed to get load balancer pool [%s]", vcdCluster.Name, vcdMachine.Name, lbPoolName)
		}
		// Do not try to update the load balancer if lbPool is not found
		if err != govcd.ErrorEntityNotFound {
			controlPlaneIPs, err := gateway.GetLoadBalancerPoolMemberIPs(ctx, lbPoolRef)
			if err != nil {
				return ctrl.Result{}, errors.Wrapf(err,
					"Error while deleting the infra resources of the machine [%s/%s]; failed to retrieve members from the load balancer pool [%s]",
					vcdCluster.Name, vcdMachine.Name, lbPoolName)
			}
			addresses := vcdMachine.Status.Addresses
			addressToBeDeleted := ""
			for _, address := range addresses {
				if address.Type == clusterv1.MachineInternalIP {
					addressToBeDeleted = address.Address
				}
			}
			updatedIPs := controlPlaneIPs
			for i, IP := range controlPlaneIPs {
				if IP == addressToBeDeleted {
					updatedIPs = append(controlPlaneIPs[:i], controlPlaneIPs[i+1:]...)
				}
			}
			err = gateway.UpdateLoadBalancer(ctx, lbPoolName, updatedIPs, int32(6443))
			if err != nil {
				return ctrl.Result{}, errors.Wrapf(err,
					"Error while deleting the infra resources of the machine [%s/%s]; error deleting the control plane from the load balancer pool [%s]",
					vcdCluster.Name, vcdMachine.Name, lbPoolName)
			}
		}
	}

	vdcManager := vcdclient.VdcManager{
		VdcName: workloadVCDClient.ClusterOVDCName,
		OrgName: workloadVCDClient.ClusterOrgName,
		Client:  workloadVCDClient,
		Vdc:     workloadVCDClient.Vdc,
	}

	// get the vApp
	vAppName := cluster.Name
	vApp, err := vdcManager.Vdc.GetVAppByName(vAppName, true)
	if err != nil {
		if err == govcd.ErrorEntityNotFound {
			log.Error(err, "Error while deleting the machine; vApp not found")
		} else {
			return ctrl.Result{}, errors.Wrapf(err, "Error while deleting the machine [%s/%s]; failed to find vapp by name", vAppName, machine.Name)
		}
	}
	if vApp != nil {
		// delete the vm
		vm, err := vApp.GetVMByName(machine.Name, true)
		if err != nil {
			if err == govcd.ErrorEntityNotFound {
				log.Error(err, "Error while deleting the machine; VM  not found")
			} else {
				return ctrl.Result{}, errors.Wrapf(err, "Error while deleting the machine [%s/%s]; unable to check if vm exists in vapp", vAppName, machine.Name)
			}
		}
		if vm != nil {
			// power-off the VM if it is powered on
			vmStatus, err := vm.GetStatus()
			if err != nil {
				klog.Warningf("Unable to get VM status for VM [%s]: [%v]", vm.VM.Name, err)
			} else {
				// continue and try to power-off in any case
				klog.Infof("VM [%s] has status [%s]", vm.VM.Name, vmStatus)
				task, err := vm.PowerOff()
				if err != nil {
					klog.Warningf("Error while powering off VM [%s]: [%v]", vm.VM.Name, err)
				} else {
					if err = task.WaitTaskCompletion(); err != nil {
						return ctrl.Result{}, fmt.Errorf("error waiting for task completion after reconfiguring vm: [%v]", err)
					}
				}
			}

			// in any case try to delete the machine
			log.Info("Deleting the infra VM of the machine")
			if err := vm.Delete(); err != nil {
				return ctrl.Result{}, errors.Wrapf(err, "error deleting the machine [%s/%s]", vAppName, vm.VM.Name)
			}
		}
		log.Info("Successfully deleted infra resources of the machine")
	}

	err = r.reconcileNodeStatusInRDE(ctx, vcdCluster.Status.InfraId, machine.Name, machine.Status.Phase, workloadVCDClient)
	if err != nil {
		if _, ok := err.(*NoRDEError); ok {
			log.V(3).Info("RDE NOT set up to track this cluster.", "infraID", vcdCluster.Status.InfraId)
		} else {
			log.Error(err, "Error reconciling the node status in the RDE",
				"InfraId", vcdCluster.Status.InfraId)
		}
	}

	controllerutil.RemoveFinalizer(vcdMachine, infrav1.MachineFinalizer)
	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *VCDMachineReconciler) SetupWithManager(ctx context.Context, mgr ctrl.Manager,
	options controller.Options) error {
	clusterToVCDMachines, err := util.ClusterToObjectsMapper(mgr.GetClient(),
		&infrav1.VCDMachineList{}, mgr.GetScheme())
	if err != nil {
		return err
	}

	c, err := ctrl.NewControllerManagedBy(mgr).
		For(&infrav1.VCDMachine{}).
		WithOptions(options).
		WithEventFilter(predicates.ResourceNotPaused(ctrl.LoggerFrom(ctx))).
		Watches(
			&source.Kind{Type: &clusterv1.Machine{}},
			handler.EnqueueRequestsFromMapFunc(
				util.MachineToInfrastructureMapFunc(infrav1.GroupVersion.WithKind("VCDMachine"))),
		).
		Watches(
			&source.Kind{Type: &infrav1.VCDCluster{}},
			handler.EnqueueRequestsFromMapFunc(r.VCDClusterToVCDMachines),
		).
		Build(r)
	if err != nil {
		return err
	}
	return c.Watch(
		&source.Kind{Type: &clusterv1.Cluster{}},
		handler.EnqueueRequestsFromMapFunc(clusterToVCDMachines),
		predicates.ClusterUnpausedAndInfrastructureReady(ctrl.LoggerFrom(ctx)),
	)
}

// VCDClusterToVCDMachines is a handler.ToRequestsFunc to be used to enqueue
// requests for reconciliation of VCDMachines.
func (r *VCDMachineReconciler) VCDClusterToVCDMachines(o client.Object) []ctrl.Request {
	var result []ctrl.Request
	c, ok := o.(*infrav1.VCDCluster)
	if !ok {
		klog.Errorf("Expected a VCDCluster found [%T]", o)
		return nil
	}

	cluster, err := util.GetOwnerCluster(context.TODO(), r.Client, c.ObjectMeta)
	switch {
	case apierrors.IsNotFound(err) || cluster == nil:
		return result
	case err != nil:
		return result
	}

	labels := map[string]string{clusterv1.ClusterLabelName: cluster.Name}
	machineList := &clusterv1.MachineList{}
	if err := r.Client.List(context.TODO(), machineList, client.InNamespace(c.Namespace),
		client.MatchingLabels(labels)); err != nil {
		return nil
	}
	for _, m := range machineList.Items {
		if m.Spec.InfrastructureRef.Name == "" {
			continue
		}
		name := client.ObjectKey{Namespace: m.Namespace, Name: m.Name}
		result = append(result, ctrl.Request{NamespacedName: name})
	}

	return result
}

func (r *VCDMachineReconciler) hasCloudInitExecutionFailedBefore(ctx context.Context, workloadVCDClient *vcdclient.Client, vm *govcd.VM) (bool, error) {
	scriptExecutionStatus, err := workloadVCDClient.GetExtraConfigValue(vm, PostCustomizationScriptExecutionStatus)
	if err != nil {
		return false, errors.Wrapf(err, "unable to get extra config value for key [%s] for vm: [%s]: [%v]",
			PostCustomizationScriptExecutionStatus, vm.VM.Name, err)
	}
	if scriptExecutionStatus != "" {
		execStatus, err := strconv.Atoi(scriptExecutionStatus)
		if err != nil {
			return false, errors.Wrapf(err, "unable to convert script execution status [%s] to int: [%v]",
				scriptExecutionStatus, err)
		}
		if execStatus != 0 {
			scriptExecutionFailureReason, err := workloadVCDClient.GetExtraConfigValue(vm, PostCustomizationScriptFailureReason)
			if err != nil {
				return false, errors.Wrapf(err, "unable to get extra config value for key [%s] for vm, "+
					"(script execution status [%d]): [%s]: [%v]",
					PostCustomizationScriptFailureReason, execStatus, vm.VM.Name, err)
			}
			return true, fmt.Errorf("script failed with status [%d] and reason [%s]", execStatus, scriptExecutionFailureReason)
		}
	}
	return false, nil
}

// MergeJinjaToCloudInitScript : merges the cloud init config with a jinja config and adds a
// `#cloudconfig` header. Does a couple of special handling: takes jinja's runcmd and embeds
// it into a fixed location in the cloudInitConfig. Returns the merged bytes or nil and error.
func MergeJinjaToCloudInitScript(cloudInitConfig string, jinjaConfig string) ([]byte, error) {
	jinja := make(map[string]interface{})
	if err := yaml.Unmarshal([]byte(jinjaConfig), &jinja); err != nil {
		return nil, fmt.Errorf("unable to unmarshal yaml [%s]: [%v]", jinjaConfig, err)
	}

	// handle runcmd before parsing vcd cloud init yaml all to simplify things
	cloudInitModified := ""
	jinjaRunCmd, ok := jinja["runcmd"]
	if ok {
		jinjaLines, ok := jinjaRunCmd.([]interface{})
		if !ok {
			return nil, fmt.Errorf("expected []interface{}, found [%T] for jinja runcmd [%v]",
				jinjaRunCmd, jinjaRunCmd)
		}

		formattedJinjaCmd := "\n"
		indent := strings.Repeat(" ", 4)
		for _, jinjaLine := range jinjaLines {
			jinjaLineStr, ok := jinjaLine.(string)
			if !ok {
				return nil, fmt.Errorf("unable to convert [%#v] to string", jinjaLineStr)
			}
			formattedJinjaCmd += indent + jinjaLineStr + "\n"
		}

		cloudInitModified = strings.Trim(
			strings.Replace(cloudInitConfig,
				"__JINJA_RUNCMD_REPLACE_ME__", strings.Trim(formattedJinjaCmd, "\n"), 1), "\r\n")
	}

	vcdCloudInit := make(map[string]interface{})
	if err := yaml.Unmarshal([]byte(cloudInitModified), &vcdCloudInit); err != nil {
		return nil, fmt.Errorf("unable to unmarshal cloud init with embedded jinja script: [%v]: [%v]",
			cloudInitModified, err)
	}

	mergedCloudInit := make(map[string]interface{})
	for key, vcdVal := range vcdCloudInit {
		jinjaVal, ok := jinja[key]
		if !ok || key == "runcmd" {
			mergedCloudInit[key] = vcdVal
			continue
		}

		switch vcdVal.(type) {
		case []interface{}:
			mergedCloudInit[key] = append(vcdVal.([]interface{}), jinjaVal.([]interface{})...)
		default:
			return nil, fmt.Errorf("unable to handle type [%T] for key [%v]", vcdVal, key)
		}
	}

	// consume the remaining keys not used in VCD
	for key, jinjaVal := range jinja {
		if _, ok := vcdCloudInit[key]; !ok {
			mergedCloudInit[key] = jinjaVal
			continue
		}
	}

	out := []byte("#cloud-config\n")
	for _, key := range []string{
		"write_files",
		"runcmd",
		"users",
		"timezone",
		"disable_root",
		"preserve_hostname",
		"hostname",
		"final_message",
	} {
		val, ok := mergedCloudInit[key]
		if !ok {
			continue
		}

		deltaMap := map[string]interface{}{
			key: val,
		}
		delta, err := yaml.Marshal(deltaMap)
		if err != nil {
			return nil, fmt.Errorf("unable to marshal [%#v]", deltaMap)
		}

		out = append(out, delta...)
	}

	return out, nil
}
