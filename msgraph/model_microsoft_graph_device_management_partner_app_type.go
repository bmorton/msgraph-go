/*
Partial Graph API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package msgraph

import (
	"encoding/json"
	"fmt"
)

// MicrosoftGraphDeviceManagementPartnerAppType Partner App Type.
type MicrosoftGraphDeviceManagementPartnerAppType string

// List of microsoft.graph.deviceManagementPartnerAppType
const (
	UNKNOWN MicrosoftGraphDeviceManagementPartnerAppType = "unknown"
	SINGLE_TENANT_APP MicrosoftGraphDeviceManagementPartnerAppType = "singleTenantApp"
	MULTI_TENANT_APP MicrosoftGraphDeviceManagementPartnerAppType = "multiTenantApp"
)

// All allowed values of MicrosoftGraphDeviceManagementPartnerAppType enum
var AllowedMicrosoftGraphDeviceManagementPartnerAppTypeEnumValues = []MicrosoftGraphDeviceManagementPartnerAppType{
	"unknown",
	"singleTenantApp",
	"multiTenantApp",
}

func (v *MicrosoftGraphDeviceManagementPartnerAppType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MicrosoftGraphDeviceManagementPartnerAppType(value)
	for _, existing := range AllowedMicrosoftGraphDeviceManagementPartnerAppTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MicrosoftGraphDeviceManagementPartnerAppType", value)
}

// NewMicrosoftGraphDeviceManagementPartnerAppTypeFromValue returns a pointer to a valid MicrosoftGraphDeviceManagementPartnerAppType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMicrosoftGraphDeviceManagementPartnerAppTypeFromValue(v string) (*MicrosoftGraphDeviceManagementPartnerAppType, error) {
	ev := MicrosoftGraphDeviceManagementPartnerAppType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MicrosoftGraphDeviceManagementPartnerAppType: valid values are %v", v, AllowedMicrosoftGraphDeviceManagementPartnerAppTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MicrosoftGraphDeviceManagementPartnerAppType) IsValid() bool {
	for _, existing := range AllowedMicrosoftGraphDeviceManagementPartnerAppTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to microsoft.graph.deviceManagementPartnerAppType value
func (v MicrosoftGraphDeviceManagementPartnerAppType) Ptr() *MicrosoftGraphDeviceManagementPartnerAppType {
	return &v
}

type NullableMicrosoftGraphDeviceManagementPartnerAppType struct {
	value *MicrosoftGraphDeviceManagementPartnerAppType
	isSet bool
}

func (v NullableMicrosoftGraphDeviceManagementPartnerAppType) Get() *MicrosoftGraphDeviceManagementPartnerAppType {
	return v.value
}

func (v *NullableMicrosoftGraphDeviceManagementPartnerAppType) Set(val *MicrosoftGraphDeviceManagementPartnerAppType) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphDeviceManagementPartnerAppType) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphDeviceManagementPartnerAppType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphDeviceManagementPartnerAppType(val *MicrosoftGraphDeviceManagementPartnerAppType) *NullableMicrosoftGraphDeviceManagementPartnerAppType {
	return &NullableMicrosoftGraphDeviceManagementPartnerAppType{value: val, isSet: true}
}

func (v NullableMicrosoftGraphDeviceManagementPartnerAppType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphDeviceManagementPartnerAppType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

