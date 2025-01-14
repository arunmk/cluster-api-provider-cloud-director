package util

import (
	"encoding/json"
	"fmt"
	"github.com/vmware/cluster-api-provider-cloud-director/pkg/vcdtypes"
)

func ConvertMapToCAPVCDEntity(entityMap map[string]interface{}) (*vcdtypes.CAPVCDEntity, error) {
	var capvcdEntity vcdtypes.CAPVCDEntity
	entityByteArr, err := json.Marshal(&entityMap)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal entity map: [%v]", err)
	}
	err = json.Unmarshal(entityByteArr, &capvcdEntity)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal entity byte array to capvcd entity: [%v]", err)
	}
	return &capvcdEntity, nil
}

func ConvertCAPVCDEntityToMap(capvcdEntity *vcdtypes.CAPVCDEntity) (map[string]interface{}, error) {
	var entityMap map[string]interface{}
	entityByteArr, err := json.Marshal(&capvcdEntity)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal CAPVCD entity to byte array: [%v]", err)
	}
	err = json.Unmarshal(entityByteArr, &entityMap)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal CAPVCD entity data to a map: [%v]", err)
	}
	return entityMap, nil
}
