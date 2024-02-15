//go:build !ignore_autogenerated_conversions
// +build !ignore_autogenerated_conversions

/*
Copyright 2019 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by conversion-gen. DO NOT EDIT.

package v1beta1

import (
	unsafe "unsafe"

	v1beta2 "github.com/vmware/cluster-api-provider-cloud-director/api/v1beta2"
	v1 "k8s.io/api/core/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	apiv1beta1 "sigs.k8s.io/cluster-api/api/v1beta1"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*APIEndpoint)(nil), (*v1beta2.APIEndpoint)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_APIEndpoint_To_v1beta2_APIEndpoint(a.(*APIEndpoint), b.(*v1beta2.APIEndpoint), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta2.APIEndpoint)(nil), (*APIEndpoint)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta2_APIEndpoint_To_v1beta1_APIEndpoint(a.(*v1beta2.APIEndpoint), b.(*APIEndpoint), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*LoadBalancerConfig)(nil), (*v1beta2.LoadBalancerConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_LoadBalancerConfig_To_v1beta2_LoadBalancerConfig(a.(*LoadBalancerConfig), b.(*v1beta2.LoadBalancerConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta2.LoadBalancerConfig)(nil), (*LoadBalancerConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta2_LoadBalancerConfig_To_v1beta1_LoadBalancerConfig(a.(*v1beta2.LoadBalancerConfig), b.(*LoadBalancerConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*Ports)(nil), (*v1beta2.Ports)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_Ports_To_v1beta2_Ports(a.(*Ports), b.(*v1beta2.Ports), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta2.Ports)(nil), (*Ports)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta2_Ports_To_v1beta1_Ports(a.(*v1beta2.Ports), b.(*Ports), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ProxyConfig)(nil), (*v1beta2.ProxyConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_ProxyConfig_To_v1beta2_ProxyConfig(a.(*ProxyConfig), b.(*v1beta2.ProxyConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta2.ProxyConfig)(nil), (*ProxyConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta2_ProxyConfig_To_v1beta1_ProxyConfig(a.(*v1beta2.ProxyConfig), b.(*ProxyConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*UserCredentialsContext)(nil), (*v1beta2.UserCredentialsContext)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_UserCredentialsContext_To_v1beta2_UserCredentialsContext(a.(*UserCredentialsContext), b.(*v1beta2.UserCredentialsContext), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta2.UserCredentialsContext)(nil), (*UserCredentialsContext)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta2_UserCredentialsContext_To_v1beta1_UserCredentialsContext(a.(*v1beta2.UserCredentialsContext), b.(*UserCredentialsContext), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*VCDCluster)(nil), (*v1beta2.VCDCluster)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_VCDCluster_To_v1beta2_VCDCluster(a.(*VCDCluster), b.(*v1beta2.VCDCluster), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta2.VCDCluster)(nil), (*VCDCluster)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta2_VCDCluster_To_v1beta1_VCDCluster(a.(*v1beta2.VCDCluster), b.(*VCDCluster), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*VCDClusterList)(nil), (*v1beta2.VCDClusterList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_VCDClusterList_To_v1beta2_VCDClusterList(a.(*VCDClusterList), b.(*v1beta2.VCDClusterList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta2.VCDClusterList)(nil), (*VCDClusterList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta2_VCDClusterList_To_v1beta1_VCDClusterList(a.(*v1beta2.VCDClusterList), b.(*VCDClusterList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*VCDClusterSpec)(nil), (*v1beta2.VCDClusterSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_VCDClusterSpec_To_v1beta2_VCDClusterSpec(a.(*VCDClusterSpec), b.(*v1beta2.VCDClusterSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta2.VCDClusterSpec)(nil), (*VCDClusterSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta2_VCDClusterSpec_To_v1beta1_VCDClusterSpec(a.(*v1beta2.VCDClusterSpec), b.(*VCDClusterSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*VCDClusterStatus)(nil), (*v1beta2.VCDClusterStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_VCDClusterStatus_To_v1beta2_VCDClusterStatus(a.(*VCDClusterStatus), b.(*v1beta2.VCDClusterStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta2.VCDClusterStatus)(nil), (*VCDClusterStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta2_VCDClusterStatus_To_v1beta1_VCDClusterStatus(a.(*v1beta2.VCDClusterStatus), b.(*VCDClusterStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*VCDMachine)(nil), (*v1beta2.VCDMachine)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_VCDMachine_To_v1beta2_VCDMachine(a.(*VCDMachine), b.(*v1beta2.VCDMachine), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta2.VCDMachine)(nil), (*VCDMachine)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta2_VCDMachine_To_v1beta1_VCDMachine(a.(*v1beta2.VCDMachine), b.(*VCDMachine), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*VCDMachineList)(nil), (*v1beta2.VCDMachineList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_VCDMachineList_To_v1beta2_VCDMachineList(a.(*VCDMachineList), b.(*v1beta2.VCDMachineList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta2.VCDMachineList)(nil), (*VCDMachineList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta2_VCDMachineList_To_v1beta1_VCDMachineList(a.(*v1beta2.VCDMachineList), b.(*VCDMachineList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*VCDMachineSpec)(nil), (*v1beta2.VCDMachineSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_VCDMachineSpec_To_v1beta2_VCDMachineSpec(a.(*VCDMachineSpec), b.(*v1beta2.VCDMachineSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*VCDMachineStatus)(nil), (*v1beta2.VCDMachineStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_VCDMachineStatus_To_v1beta2_VCDMachineStatus(a.(*VCDMachineStatus), b.(*v1beta2.VCDMachineStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta2.VCDMachineStatus)(nil), (*VCDMachineStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta2_VCDMachineStatus_To_v1beta1_VCDMachineStatus(a.(*v1beta2.VCDMachineStatus), b.(*VCDMachineStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*VCDMachineTemplate)(nil), (*v1beta2.VCDMachineTemplate)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_VCDMachineTemplate_To_v1beta2_VCDMachineTemplate(a.(*VCDMachineTemplate), b.(*v1beta2.VCDMachineTemplate), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta2.VCDMachineTemplate)(nil), (*VCDMachineTemplate)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta2_VCDMachineTemplate_To_v1beta1_VCDMachineTemplate(a.(*v1beta2.VCDMachineTemplate), b.(*VCDMachineTemplate), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*VCDMachineTemplateList)(nil), (*v1beta2.VCDMachineTemplateList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_VCDMachineTemplateList_To_v1beta2_VCDMachineTemplateList(a.(*VCDMachineTemplateList), b.(*v1beta2.VCDMachineTemplateList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta2.VCDMachineTemplateList)(nil), (*VCDMachineTemplateList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta2_VCDMachineTemplateList_To_v1beta1_VCDMachineTemplateList(a.(*v1beta2.VCDMachineTemplateList), b.(*VCDMachineTemplateList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*VCDMachineTemplateResource)(nil), (*v1beta2.VCDMachineTemplateResource)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_VCDMachineTemplateResource_To_v1beta2_VCDMachineTemplateResource(a.(*VCDMachineTemplateResource), b.(*v1beta2.VCDMachineTemplateResource), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*VCDMachineTemplateSpec)(nil), (*v1beta2.VCDMachineTemplateSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_VCDMachineTemplateSpec_To_v1beta2_VCDMachineTemplateSpec(a.(*VCDMachineTemplateSpec), b.(*v1beta2.VCDMachineTemplateSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta2.VCDMachineTemplateSpec)(nil), (*VCDMachineTemplateSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta2_VCDMachineTemplateSpec_To_v1beta1_VCDMachineTemplateSpec(a.(*v1beta2.VCDMachineTemplateSpec), b.(*VCDMachineTemplateSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*VCDMachineTemplateStatus)(nil), (*v1beta2.VCDMachineTemplateStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_VCDMachineTemplateStatus_To_v1beta2_VCDMachineTemplateStatus(a.(*VCDMachineTemplateStatus), b.(*v1beta2.VCDMachineTemplateStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta2.VCDMachineTemplateStatus)(nil), (*VCDMachineTemplateStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta2_VCDMachineTemplateStatus_To_v1beta1_VCDMachineTemplateStatus(a.(*v1beta2.VCDMachineTemplateStatus), b.(*VCDMachineTemplateStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*v1beta2.VCDMachineSpec)(nil), (*VCDMachineSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta2_VCDMachineSpec_To_v1beta1_VCDMachineSpec(a.(*v1beta2.VCDMachineSpec), b.(*VCDMachineSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*v1beta2.VCDMachineTemplateResource)(nil), (*VCDMachineTemplateResource)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta2_VCDMachineTemplateResource_To_v1beta1_VCDMachineTemplateResource(a.(*v1beta2.VCDMachineTemplateResource), b.(*VCDMachineTemplateResource), scope)
	}); err != nil {
		return err
	}
	return nil
}

func Convert_v1beta2_VCDClusterStatus_To_v1beta1_VCDClusterStatus(in *v1beta2.VCDClusterStatus, out *VCDClusterStatus, s conversion.Scope) error {
    return autoConvert_v1beta2_VCDClusterStatus_To_v1beta1_VCDClusterStatus(in, out, s)
}

func autoConvert_v1beta1_APIEndpoint_To_v1beta2_APIEndpoint(in *APIEndpoint, out *v1beta2.APIEndpoint, s conversion.Scope) error {
	out.Host = in.Host
	out.Port = in.Port
	return nil
}

// Convert_v1beta1_APIEndpoint_To_v1beta2_APIEndpoint is an autogenerated conversion function.
func Convert_v1beta1_APIEndpoint_To_v1beta2_APIEndpoint(in *APIEndpoint, out *v1beta2.APIEndpoint, s conversion.Scope) error {
	return autoConvert_v1beta1_APIEndpoint_To_v1beta2_APIEndpoint(in, out, s)
}

func autoConvert_v1beta2_APIEndpoint_To_v1beta1_APIEndpoint(in *v1beta2.APIEndpoint, out *APIEndpoint, s conversion.Scope) error {
	out.Host = in.Host
	out.Port = in.Port
	return nil
}

// Convert_v1beta2_APIEndpoint_To_v1beta1_APIEndpoint is an autogenerated conversion function.
func Convert_v1beta2_APIEndpoint_To_v1beta1_APIEndpoint(in *v1beta2.APIEndpoint, out *APIEndpoint, s conversion.Scope) error {
	return autoConvert_v1beta2_APIEndpoint_To_v1beta1_APIEndpoint(in, out, s)
}

func autoConvert_v1beta1_LoadBalancerConfig_To_v1beta2_LoadBalancerConfig(in *LoadBalancerConfig, out *v1beta2.LoadBalancerConfig, s conversion.Scope) error {
	out.UseOneArm = in.UseOneArm
	out.VipSubnet = in.VipSubnet
	return nil
}

// Convert_v1beta1_LoadBalancerConfig_To_v1beta2_LoadBalancerConfig is an autogenerated conversion function.
func Convert_v1beta1_LoadBalancerConfig_To_v1beta2_LoadBalancerConfig(in *LoadBalancerConfig, out *v1beta2.LoadBalancerConfig, s conversion.Scope) error {
	return autoConvert_v1beta1_LoadBalancerConfig_To_v1beta2_LoadBalancerConfig(in, out, s)
}

func autoConvert_v1beta2_LoadBalancerConfig_To_v1beta1_LoadBalancerConfig(in *v1beta2.LoadBalancerConfig, out *LoadBalancerConfig, s conversion.Scope) error {
	out.UseOneArm = in.UseOneArm
	out.VipSubnet = in.VipSubnet
	return nil
}

// Convert_v1beta2_LoadBalancerConfig_To_v1beta1_LoadBalancerConfig is an autogenerated conversion function.
func Convert_v1beta2_LoadBalancerConfig_To_v1beta1_LoadBalancerConfig(in *v1beta2.LoadBalancerConfig, out *LoadBalancerConfig, s conversion.Scope) error {
	return autoConvert_v1beta2_LoadBalancerConfig_To_v1beta1_LoadBalancerConfig(in, out, s)
}

func autoConvert_v1beta1_Ports_To_v1beta2_Ports(in *Ports, out *v1beta2.Ports, s conversion.Scope) error {
	out.HTTP = in.HTTP
	out.HTTPS = in.HTTPS
	out.TCP = in.TCP
	return nil
}

// Convert_v1beta1_Ports_To_v1beta2_Ports is an autogenerated conversion function.
func Convert_v1beta1_Ports_To_v1beta2_Ports(in *Ports, out *v1beta2.Ports, s conversion.Scope) error {
	return autoConvert_v1beta1_Ports_To_v1beta2_Ports(in, out, s)
}

func autoConvert_v1beta2_Ports_To_v1beta1_Ports(in *v1beta2.Ports, out *Ports, s conversion.Scope) error {
	out.HTTP = in.HTTP
	out.HTTPS = in.HTTPS
	out.TCP = in.TCP
	return nil
}

// Convert_v1beta2_Ports_To_v1beta1_Ports is an autogenerated conversion function.
func Convert_v1beta2_Ports_To_v1beta1_Ports(in *v1beta2.Ports, out *Ports, s conversion.Scope) error {
	return autoConvert_v1beta2_Ports_To_v1beta1_Ports(in, out, s)
}

func autoConvert_v1beta1_ProxyConfig_To_v1beta2_ProxyConfig(in *ProxyConfig, out *v1beta2.ProxyConfig, s conversion.Scope) error {
	out.HTTPProxy = in.HTTPProxy
	out.HTTPSProxy = in.HTTPSProxy
	out.NoProxy = in.NoProxy
	return nil
}

// Convert_v1beta1_ProxyConfig_To_v1beta2_ProxyConfig is an autogenerated conversion function.
func Convert_v1beta1_ProxyConfig_To_v1beta2_ProxyConfig(in *ProxyConfig, out *v1beta2.ProxyConfig, s conversion.Scope) error {
	return autoConvert_v1beta1_ProxyConfig_To_v1beta2_ProxyConfig(in, out, s)
}

func autoConvert_v1beta2_ProxyConfig_To_v1beta1_ProxyConfig(in *v1beta2.ProxyConfig, out *ProxyConfig, s conversion.Scope) error {
	out.HTTPProxy = in.HTTPProxy
	out.HTTPSProxy = in.HTTPSProxy
	out.NoProxy = in.NoProxy
	return nil
}

// Convert_v1beta2_ProxyConfig_To_v1beta1_ProxyConfig is an autogenerated conversion function.
func Convert_v1beta2_ProxyConfig_To_v1beta1_ProxyConfig(in *v1beta2.ProxyConfig, out *ProxyConfig, s conversion.Scope) error {
	return autoConvert_v1beta2_ProxyConfig_To_v1beta1_ProxyConfig(in, out, s)
}

func autoConvert_v1beta1_UserCredentialsContext_To_v1beta2_UserCredentialsContext(in *UserCredentialsContext, out *v1beta2.UserCredentialsContext, s conversion.Scope) error {
	out.Username = in.Username
	out.Password = in.Password
	out.RefreshToken = in.RefreshToken
	out.SecretRef = (*v1.SecretReference)(unsafe.Pointer(in.SecretRef))
	return nil
}

// Convert_v1beta1_UserCredentialsContext_To_v1beta2_UserCredentialsContext is an autogenerated conversion function.
func Convert_v1beta1_UserCredentialsContext_To_v1beta2_UserCredentialsContext(in *UserCredentialsContext, out *v1beta2.UserCredentialsContext, s conversion.Scope) error {
	return autoConvert_v1beta1_UserCredentialsContext_To_v1beta2_UserCredentialsContext(in, out, s)
}

func autoConvert_v1beta2_UserCredentialsContext_To_v1beta1_UserCredentialsContext(in *v1beta2.UserCredentialsContext, out *UserCredentialsContext, s conversion.Scope) error {
	out.Username = in.Username
	out.Password = in.Password
	out.RefreshToken = in.RefreshToken
	out.SecretRef = (*v1.SecretReference)(unsafe.Pointer(in.SecretRef))
	return nil
}

// Convert_v1beta2_UserCredentialsContext_To_v1beta1_UserCredentialsContext is an autogenerated conversion function.
func Convert_v1beta2_UserCredentialsContext_To_v1beta1_UserCredentialsContext(in *v1beta2.UserCredentialsContext, out *UserCredentialsContext, s conversion.Scope) error {
	return autoConvert_v1beta2_UserCredentialsContext_To_v1beta1_UserCredentialsContext(in, out, s)
}

func autoConvert_v1beta1_VCDCluster_To_v1beta2_VCDCluster(in *VCDCluster, out *v1beta2.VCDCluster, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1beta1_VCDClusterSpec_To_v1beta2_VCDClusterSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1beta1_VCDClusterStatus_To_v1beta2_VCDClusterStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1beta1_VCDCluster_To_v1beta2_VCDCluster is an autogenerated conversion function.
func Convert_v1beta1_VCDCluster_To_v1beta2_VCDCluster(in *VCDCluster, out *v1beta2.VCDCluster, s conversion.Scope) error {
	return autoConvert_v1beta1_VCDCluster_To_v1beta2_VCDCluster(in, out, s)
}

func autoConvert_v1beta2_VCDCluster_To_v1beta1_VCDCluster(in *v1beta2.VCDCluster, out *VCDCluster, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1beta2_VCDClusterSpec_To_v1beta1_VCDClusterSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1beta2_VCDClusterStatus_To_v1beta1_VCDClusterStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1beta2_VCDCluster_To_v1beta1_VCDCluster is an autogenerated conversion function.
func Convert_v1beta2_VCDCluster_To_v1beta1_VCDCluster(in *v1beta2.VCDCluster, out *VCDCluster, s conversion.Scope) error {
	return autoConvert_v1beta2_VCDCluster_To_v1beta1_VCDCluster(in, out, s)
}

func autoConvert_v1beta1_VCDClusterList_To_v1beta2_VCDClusterList(in *VCDClusterList, out *v1beta2.VCDClusterList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]v1beta2.VCDCluster, len(*in))
		for i := range *in {
			if err := Convert_v1beta1_VCDCluster_To_v1beta2_VCDCluster(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_v1beta1_VCDClusterList_To_v1beta2_VCDClusterList is an autogenerated conversion function.
func Convert_v1beta1_VCDClusterList_To_v1beta2_VCDClusterList(in *VCDClusterList, out *v1beta2.VCDClusterList, s conversion.Scope) error {
	return autoConvert_v1beta1_VCDClusterList_To_v1beta2_VCDClusterList(in, out, s)
}

func autoConvert_v1beta2_VCDClusterList_To_v1beta1_VCDClusterList(in *v1beta2.VCDClusterList, out *VCDClusterList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VCDCluster, len(*in))
		for i := range *in {
			if err := Convert_v1beta2_VCDCluster_To_v1beta1_VCDCluster(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_v1beta2_VCDClusterList_To_v1beta1_VCDClusterList is an autogenerated conversion function.
func Convert_v1beta2_VCDClusterList_To_v1beta1_VCDClusterList(in *v1beta2.VCDClusterList, out *VCDClusterList, s conversion.Scope) error {
	return autoConvert_v1beta2_VCDClusterList_To_v1beta1_VCDClusterList(in, out, s)
}

func autoConvert_v1beta1_VCDClusterSpec_To_v1beta2_VCDClusterSpec(in *VCDClusterSpec, out *v1beta2.VCDClusterSpec, s conversion.Scope) error {
	if err := Convert_v1beta1_APIEndpoint_To_v1beta2_APIEndpoint(&in.ControlPlaneEndpoint, &out.ControlPlaneEndpoint, s); err != nil {
		return err
	}
	out.Site = in.Site
	out.Org = in.Org
	out.Ovdc = in.Ovdc
	out.OvdcNetwork = in.OvdcNetwork
	if err := Convert_v1beta1_UserCredentialsContext_To_v1beta2_UserCredentialsContext(&in.UserCredentialsContext, &out.UserCredentialsContext, s); err != nil {
		return err
	}
	out.RDEId = in.RDEId
	out.ParentUID = in.ParentUID
	out.UseAsManagementCluster = in.UseAsManagementCluster
	if err := Convert_v1beta1_ProxyConfig_To_v1beta2_ProxyConfig(&in.ProxyConfigSpec, &out.ProxyConfigSpec, s); err != nil {
		return err
	}
	if err := Convert_v1beta1_LoadBalancerConfig_To_v1beta2_LoadBalancerConfig(&in.LoadBalancerConfigSpec, &out.LoadBalancerConfigSpec, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1beta1_VCDClusterSpec_To_v1beta2_VCDClusterSpec is an autogenerated conversion function.
func Convert_v1beta1_VCDClusterSpec_To_v1beta2_VCDClusterSpec(in *VCDClusterSpec, out *v1beta2.VCDClusterSpec, s conversion.Scope) error {
	return autoConvert_v1beta1_VCDClusterSpec_To_v1beta2_VCDClusterSpec(in, out, s)
}

func autoConvert_v1beta2_VCDClusterSpec_To_v1beta1_VCDClusterSpec(in *v1beta2.VCDClusterSpec, out *VCDClusterSpec, s conversion.Scope) error {
	if err := Convert_v1beta2_APIEndpoint_To_v1beta1_APIEndpoint(&in.ControlPlaneEndpoint, &out.ControlPlaneEndpoint, s); err != nil {
		return err
	}
	out.Site = in.Site
	out.Org = in.Org
	out.Ovdc = in.Ovdc
	out.OvdcNetwork = in.OvdcNetwork
	if err := Convert_v1beta2_UserCredentialsContext_To_v1beta1_UserCredentialsContext(&in.UserCredentialsContext, &out.UserCredentialsContext, s); err != nil {
		return err
	}
	out.RDEId = in.RDEId
	out.ParentUID = in.ParentUID
	out.UseAsManagementCluster = in.UseAsManagementCluster
	if err := Convert_v1beta2_ProxyConfig_To_v1beta1_ProxyConfig(&in.ProxyConfigSpec, &out.ProxyConfigSpec, s); err != nil {
		return err
	}
	if err := Convert_v1beta2_LoadBalancerConfig_To_v1beta1_LoadBalancerConfig(&in.LoadBalancerConfigSpec, &out.LoadBalancerConfigSpec, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1beta2_VCDClusterSpec_To_v1beta1_VCDClusterSpec is an autogenerated conversion function.
func Convert_v1beta2_VCDClusterSpec_To_v1beta1_VCDClusterSpec(in *v1beta2.VCDClusterSpec, out *VCDClusterSpec, s conversion.Scope) error {
	return autoConvert_v1beta2_VCDClusterSpec_To_v1beta1_VCDClusterSpec(in, out, s)
}

func autoConvert_v1beta1_VCDClusterStatus_To_v1beta2_VCDClusterStatus(in *VCDClusterStatus, out *v1beta2.VCDClusterStatus, s conversion.Scope) error {
	out.Ready = in.Ready
	out.RdeVersionInUse = in.RdeVersionInUse
	out.VAppMetadataUpdated = in.VAppMetadataUpdated
	out.Conditions = *(*apiv1beta1.Conditions)(unsafe.Pointer(&in.Conditions))
	out.Site = in.Site
	out.Org = in.Org
	out.Ovdc = in.Ovdc
	out.OvdcNetwork = in.OvdcNetwork
	out.InfraId = in.InfraId
	out.ParentUID = in.ParentUID
	out.UseAsManagementCluster = in.UseAsManagementCluster
	if err := Convert_v1beta1_ProxyConfig_To_v1beta2_ProxyConfig(&in.ProxyConfig, &out.ProxyConfig, s); err != nil {
		return err
	}
	if err := Convert_v1beta1_LoadBalancerConfig_To_v1beta2_LoadBalancerConfig(&in.LoadBalancerConfig, &out.LoadBalancerConfig, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1beta1_VCDClusterStatus_To_v1beta2_VCDClusterStatus is an autogenerated conversion function.
func Convert_v1beta1_VCDClusterStatus_To_v1beta2_VCDClusterStatus(in *VCDClusterStatus, out *v1beta2.VCDClusterStatus, s conversion.Scope) error {
	return autoConvert_v1beta1_VCDClusterStatus_To_v1beta2_VCDClusterStatus(in, out, s)
}

func autoConvert_v1beta2_VCDClusterStatus_To_v1beta1_VCDClusterStatus(in *v1beta2.VCDClusterStatus, out *VCDClusterStatus, s conversion.Scope) error {
	out.Ready = in.Ready
	out.RdeVersionInUse = in.RdeVersionInUse
	out.VAppMetadataUpdated = in.VAppMetadataUpdated
	out.Conditions = *(*apiv1beta1.Conditions)(unsafe.Pointer(&in.Conditions))
	out.Site = in.Site
	// WARNING: in.VcdResourceMap requires manual conversion: does not exist in peer-type
	out.Org = in.Org
	out.Ovdc = in.Ovdc
	out.OvdcNetwork = in.OvdcNetwork
	out.InfraId = in.InfraId
	out.ParentUID = in.ParentUID
	out.UseAsManagementCluster = in.UseAsManagementCluster
	if err := Convert_v1beta2_ProxyConfig_To_v1beta1_ProxyConfig(&in.ProxyConfig, &out.ProxyConfig, s); err != nil {
		return err
	}
	if err := Convert_v1beta2_LoadBalancerConfig_To_v1beta1_LoadBalancerConfig(&in.LoadBalancerConfig, &out.LoadBalancerConfig, s); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1beta1_VCDMachine_To_v1beta2_VCDMachine(in *VCDMachine, out *v1beta2.VCDMachine, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1beta1_VCDMachineSpec_To_v1beta2_VCDMachineSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1beta1_VCDMachineStatus_To_v1beta2_VCDMachineStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1beta1_VCDMachine_To_v1beta2_VCDMachine is an autogenerated conversion function.
func Convert_v1beta1_VCDMachine_To_v1beta2_VCDMachine(in *VCDMachine, out *v1beta2.VCDMachine, s conversion.Scope) error {
	return autoConvert_v1beta1_VCDMachine_To_v1beta2_VCDMachine(in, out, s)
}

func autoConvert_v1beta2_VCDMachine_To_v1beta1_VCDMachine(in *v1beta2.VCDMachine, out *VCDMachine, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1beta2_VCDMachineSpec_To_v1beta1_VCDMachineSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1beta2_VCDMachineStatus_To_v1beta1_VCDMachineStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1beta2_VCDMachine_To_v1beta1_VCDMachine is an autogenerated conversion function.
func Convert_v1beta2_VCDMachine_To_v1beta1_VCDMachine(in *v1beta2.VCDMachine, out *VCDMachine, s conversion.Scope) error {
	return autoConvert_v1beta2_VCDMachine_To_v1beta1_VCDMachine(in, out, s)
}

func autoConvert_v1beta1_VCDMachineList_To_v1beta2_VCDMachineList(in *VCDMachineList, out *v1beta2.VCDMachineList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]v1beta2.VCDMachine, len(*in))
		for i := range *in {
			if err := Convert_v1beta1_VCDMachine_To_v1beta2_VCDMachine(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_v1beta1_VCDMachineList_To_v1beta2_VCDMachineList is an autogenerated conversion function.
func Convert_v1beta1_VCDMachineList_To_v1beta2_VCDMachineList(in *VCDMachineList, out *v1beta2.VCDMachineList, s conversion.Scope) error {
	return autoConvert_v1beta1_VCDMachineList_To_v1beta2_VCDMachineList(in, out, s)
}

func autoConvert_v1beta2_VCDMachineList_To_v1beta1_VCDMachineList(in *v1beta2.VCDMachineList, out *VCDMachineList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VCDMachine, len(*in))
		for i := range *in {
			if err := Convert_v1beta2_VCDMachine_To_v1beta1_VCDMachine(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_v1beta2_VCDMachineList_To_v1beta1_VCDMachineList is an autogenerated conversion function.
func Convert_v1beta2_VCDMachineList_To_v1beta1_VCDMachineList(in *v1beta2.VCDMachineList, out *VCDMachineList, s conversion.Scope) error {
	return autoConvert_v1beta2_VCDMachineList_To_v1beta1_VCDMachineList(in, out, s)
}

func autoConvert_v1beta1_VCDMachineSpec_To_v1beta2_VCDMachineSpec(in *VCDMachineSpec, out *v1beta2.VCDMachineSpec, s conversion.Scope) error {
	out.ProviderID = (*string)(unsafe.Pointer(in.ProviderID))
	out.Catalog = in.Catalog
	out.Template = in.Template
	out.SizingPolicy = in.SizingPolicy
	out.PlacementPolicy = in.PlacementPolicy
	out.StorageProfile = in.StorageProfile
	out.DiskSize = in.DiskSize
	out.Bootstrapped = in.Bootstrapped
	out.EnableNvidiaGPU = in.EnableNvidiaGPU
	return nil
}

// Convert_v1beta1_VCDMachineSpec_To_v1beta2_VCDMachineSpec is an autogenerated conversion function.
func Convert_v1beta1_VCDMachineSpec_To_v1beta2_VCDMachineSpec(in *VCDMachineSpec, out *v1beta2.VCDMachineSpec, s conversion.Scope) error {
	return autoConvert_v1beta1_VCDMachineSpec_To_v1beta2_VCDMachineSpec(in, out, s)
}

func autoConvert_v1beta2_VCDMachineSpec_To_v1beta1_VCDMachineSpec(in *v1beta2.VCDMachineSpec, out *VCDMachineSpec, s conversion.Scope) error {
	out.ProviderID = (*string)(unsafe.Pointer(in.ProviderID))
	out.Catalog = in.Catalog
	out.Template = in.Template
	out.SizingPolicy = in.SizingPolicy
	out.PlacementPolicy = in.PlacementPolicy
	out.StorageProfile = in.StorageProfile
	out.DiskSize = in.DiskSize
	out.Bootstrapped = in.Bootstrapped
	out.EnableNvidiaGPU = in.EnableNvidiaGPU
	// WARNING: in.ExtraOvdcNetworks requires manual conversion: does not exist in peer-type
	// WARNING: in.VmNamingTemplate requires manual conversion: does not exist in peer-type
	return nil
}

func autoConvert_v1beta1_VCDMachineStatus_To_v1beta2_VCDMachineStatus(in *VCDMachineStatus, out *v1beta2.VCDMachineStatus, s conversion.Scope) error {
	out.ProviderID = (*string)(unsafe.Pointer(in.ProviderID))
	out.Ready = in.Ready
	out.Addresses = *(*[]apiv1beta1.MachineAddress)(unsafe.Pointer(&in.Addresses))
	out.Template = in.Template
	out.SizingPolicy = in.SizingPolicy
	out.PlacementPolicy = in.PlacementPolicy
	out.NvidiaGPUEnabled = in.NvidiaGPUEnabled
	out.DiskSize = in.DiskSize
	out.Conditions = *(*apiv1beta1.Conditions)(unsafe.Pointer(&in.Conditions))
	return nil
}

// Convert_v1beta1_VCDMachineStatus_To_v1beta2_VCDMachineStatus is an autogenerated conversion function.
func Convert_v1beta1_VCDMachineStatus_To_v1beta2_VCDMachineStatus(in *VCDMachineStatus, out *v1beta2.VCDMachineStatus, s conversion.Scope) error {
	return autoConvert_v1beta1_VCDMachineStatus_To_v1beta2_VCDMachineStatus(in, out, s)
}

func autoConvert_v1beta2_VCDMachineStatus_To_v1beta1_VCDMachineStatus(in *v1beta2.VCDMachineStatus, out *VCDMachineStatus, s conversion.Scope) error {
	out.ProviderID = (*string)(unsafe.Pointer(in.ProviderID))
	out.Ready = in.Ready
	out.Addresses = *(*[]apiv1beta1.MachineAddress)(unsafe.Pointer(&in.Addresses))
	out.Template = in.Template
	out.SizingPolicy = in.SizingPolicy
	out.PlacementPolicy = in.PlacementPolicy
	out.NvidiaGPUEnabled = in.NvidiaGPUEnabled
	out.DiskSize = in.DiskSize
	out.Conditions = *(*apiv1beta1.Conditions)(unsafe.Pointer(&in.Conditions))
	return nil
}

// Convert_v1beta2_VCDMachineStatus_To_v1beta1_VCDMachineStatus is an autogenerated conversion function.
func Convert_v1beta2_VCDMachineStatus_To_v1beta1_VCDMachineStatus(in *v1beta2.VCDMachineStatus, out *VCDMachineStatus, s conversion.Scope) error {
	return autoConvert_v1beta2_VCDMachineStatus_To_v1beta1_VCDMachineStatus(in, out, s)
}

func autoConvert_v1beta1_VCDMachineTemplate_To_v1beta2_VCDMachineTemplate(in *VCDMachineTemplate, out *v1beta2.VCDMachineTemplate, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1beta1_VCDMachineTemplateSpec_To_v1beta2_VCDMachineTemplateSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1beta1_VCDMachineTemplateStatus_To_v1beta2_VCDMachineTemplateStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1beta1_VCDMachineTemplate_To_v1beta2_VCDMachineTemplate is an autogenerated conversion function.
func Convert_v1beta1_VCDMachineTemplate_To_v1beta2_VCDMachineTemplate(in *VCDMachineTemplate, out *v1beta2.VCDMachineTemplate, s conversion.Scope) error {
	return autoConvert_v1beta1_VCDMachineTemplate_To_v1beta2_VCDMachineTemplate(in, out, s)
}

func autoConvert_v1beta2_VCDMachineTemplate_To_v1beta1_VCDMachineTemplate(in *v1beta2.VCDMachineTemplate, out *VCDMachineTemplate, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1beta2_VCDMachineTemplateSpec_To_v1beta1_VCDMachineTemplateSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1beta2_VCDMachineTemplateStatus_To_v1beta1_VCDMachineTemplateStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1beta2_VCDMachineTemplate_To_v1beta1_VCDMachineTemplate is an autogenerated conversion function.
func Convert_v1beta2_VCDMachineTemplate_To_v1beta1_VCDMachineTemplate(in *v1beta2.VCDMachineTemplate, out *VCDMachineTemplate, s conversion.Scope) error {
	return autoConvert_v1beta2_VCDMachineTemplate_To_v1beta1_VCDMachineTemplate(in, out, s)
}

func autoConvert_v1beta1_VCDMachineTemplateList_To_v1beta2_VCDMachineTemplateList(in *VCDMachineTemplateList, out *v1beta2.VCDMachineTemplateList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]v1beta2.VCDMachineTemplate, len(*in))
		for i := range *in {
			if err := Convert_v1beta1_VCDMachineTemplate_To_v1beta2_VCDMachineTemplate(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_v1beta1_VCDMachineTemplateList_To_v1beta2_VCDMachineTemplateList is an autogenerated conversion function.
func Convert_v1beta1_VCDMachineTemplateList_To_v1beta2_VCDMachineTemplateList(in *VCDMachineTemplateList, out *v1beta2.VCDMachineTemplateList, s conversion.Scope) error {
	return autoConvert_v1beta1_VCDMachineTemplateList_To_v1beta2_VCDMachineTemplateList(in, out, s)
}

func autoConvert_v1beta2_VCDMachineTemplateList_To_v1beta1_VCDMachineTemplateList(in *v1beta2.VCDMachineTemplateList, out *VCDMachineTemplateList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VCDMachineTemplate, len(*in))
		for i := range *in {
			if err := Convert_v1beta2_VCDMachineTemplate_To_v1beta1_VCDMachineTemplate(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_v1beta2_VCDMachineTemplateList_To_v1beta1_VCDMachineTemplateList is an autogenerated conversion function.
func Convert_v1beta2_VCDMachineTemplateList_To_v1beta1_VCDMachineTemplateList(in *v1beta2.VCDMachineTemplateList, out *VCDMachineTemplateList, s conversion.Scope) error {
	return autoConvert_v1beta2_VCDMachineTemplateList_To_v1beta1_VCDMachineTemplateList(in, out, s)
}

func autoConvert_v1beta1_VCDMachineTemplateResource_To_v1beta2_VCDMachineTemplateResource(in *VCDMachineTemplateResource, out *v1beta2.VCDMachineTemplateResource, s conversion.Scope) error {
	if err := Convert_v1beta1_VCDMachineSpec_To_v1beta2_VCDMachineSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1beta1_VCDMachineTemplateResource_To_v1beta2_VCDMachineTemplateResource is an autogenerated conversion function.
func Convert_v1beta1_VCDMachineTemplateResource_To_v1beta2_VCDMachineTemplateResource(in *VCDMachineTemplateResource, out *v1beta2.VCDMachineTemplateResource, s conversion.Scope) error {
	return autoConvert_v1beta1_VCDMachineTemplateResource_To_v1beta2_VCDMachineTemplateResource(in, out, s)
}

func autoConvert_v1beta2_VCDMachineTemplateResource_To_v1beta1_VCDMachineTemplateResource(in *v1beta2.VCDMachineTemplateResource, out *VCDMachineTemplateResource, s conversion.Scope) error {
	// WARNING: in.ObjectMeta requires manual conversion: does not exist in peer-type
	if err := Convert_v1beta2_VCDMachineSpec_To_v1beta1_VCDMachineSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1beta1_VCDMachineTemplateSpec_To_v1beta2_VCDMachineTemplateSpec(in *VCDMachineTemplateSpec, out *v1beta2.VCDMachineTemplateSpec, s conversion.Scope) error {
	if err := Convert_v1beta1_VCDMachineTemplateResource_To_v1beta2_VCDMachineTemplateResource(&in.Template, &out.Template, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1beta1_VCDMachineTemplateSpec_To_v1beta2_VCDMachineTemplateSpec is an autogenerated conversion function.
func Convert_v1beta1_VCDMachineTemplateSpec_To_v1beta2_VCDMachineTemplateSpec(in *VCDMachineTemplateSpec, out *v1beta2.VCDMachineTemplateSpec, s conversion.Scope) error {
	return autoConvert_v1beta1_VCDMachineTemplateSpec_To_v1beta2_VCDMachineTemplateSpec(in, out, s)
}

func autoConvert_v1beta2_VCDMachineTemplateSpec_To_v1beta1_VCDMachineTemplateSpec(in *v1beta2.VCDMachineTemplateSpec, out *VCDMachineTemplateSpec, s conversion.Scope) error {
	if err := Convert_v1beta2_VCDMachineTemplateResource_To_v1beta1_VCDMachineTemplateResource(&in.Template, &out.Template, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1beta2_VCDMachineTemplateSpec_To_v1beta1_VCDMachineTemplateSpec is an autogenerated conversion function.
func Convert_v1beta2_VCDMachineTemplateSpec_To_v1beta1_VCDMachineTemplateSpec(in *v1beta2.VCDMachineTemplateSpec, out *VCDMachineTemplateSpec, s conversion.Scope) error {
	return autoConvert_v1beta2_VCDMachineTemplateSpec_To_v1beta1_VCDMachineTemplateSpec(in, out, s)
}

func autoConvert_v1beta1_VCDMachineTemplateStatus_To_v1beta2_VCDMachineTemplateStatus(in *VCDMachineTemplateStatus, out *v1beta2.VCDMachineTemplateStatus, s conversion.Scope) error {
	return nil
}

// Convert_v1beta1_VCDMachineTemplateStatus_To_v1beta2_VCDMachineTemplateStatus is an autogenerated conversion function.
func Convert_v1beta1_VCDMachineTemplateStatus_To_v1beta2_VCDMachineTemplateStatus(in *VCDMachineTemplateStatus, out *v1beta2.VCDMachineTemplateStatus, s conversion.Scope) error {
	return autoConvert_v1beta1_VCDMachineTemplateStatus_To_v1beta2_VCDMachineTemplateStatus(in, out, s)
}

func autoConvert_v1beta2_VCDMachineTemplateStatus_To_v1beta1_VCDMachineTemplateStatus(in *v1beta2.VCDMachineTemplateStatus, out *VCDMachineTemplateStatus, s conversion.Scope) error {
	return nil
}

// Convert_v1beta2_VCDMachineTemplateStatus_To_v1beta1_VCDMachineTemplateStatus is an autogenerated conversion function.
func Convert_v1beta2_VCDMachineTemplateStatus_To_v1beta1_VCDMachineTemplateStatus(in *v1beta2.VCDMachineTemplateStatus, out *VCDMachineTemplateStatus, s conversion.Scope) error {
	return autoConvert_v1beta2_VCDMachineTemplateStatus_To_v1beta1_VCDMachineTemplateStatus(in, out, s)
}
