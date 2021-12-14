/*
Partial Graph API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// MicrosoftGraphManagementAgentType the model 'MicrosoftGraphManagementAgentType'
type MicrosoftGraphManagementAgentType string

// List of microsoft.graph.managementAgentType
const (
	EAS MicrosoftGraphManagementAgentType = "eas"
	MDM MicrosoftGraphManagementAgentType = "mdm"
	EAS_MDM MicrosoftGraphManagementAgentType = "easMdm"
	INTUNE_CLIENT MicrosoftGraphManagementAgentType = "intuneClient"
	EAS_INTUNE_CLIENT MicrosoftGraphManagementAgentType = "easIntuneClient"
	CONFIGURATION_MANAGER_CLIENT MicrosoftGraphManagementAgentType = "configurationManagerClient"
	CONFIGURATION_MANAGER_CLIENT_MDM MicrosoftGraphManagementAgentType = "configurationManagerClientMdm"
	CONFIGURATION_MANAGER_CLIENT_MDM_EAS MicrosoftGraphManagementAgentType = "configurationManagerClientMdmEas"
	UNKNOWN MicrosoftGraphManagementAgentType = "unknown"
	JAMF MicrosoftGraphManagementAgentType = "jamf"
	GOOGLE_CLOUD_DEVICE_POLICY_CONTROLLER MicrosoftGraphManagementAgentType = "googleCloudDevicePolicyController"
	MICROSOFT365_MANAGED_MDM MicrosoftGraphManagementAgentType = "microsoft365ManagedMdm"
	MS_SENSE MicrosoftGraphManagementAgentType = "msSense"
)

// All allowed values of MicrosoftGraphManagementAgentType enum
var AllowedMicrosoftGraphManagementAgentTypeEnumValues = []MicrosoftGraphManagementAgentType{
	"eas",
	"mdm",
	"easMdm",
	"intuneClient",
	"easIntuneClient",
	"configurationManagerClient",
	"configurationManagerClientMdm",
	"configurationManagerClientMdmEas",
	"unknown",
	"jamf",
	"googleCloudDevicePolicyController",
	"microsoft365ManagedMdm",
	"msSense",
}

func (v *MicrosoftGraphManagementAgentType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MicrosoftGraphManagementAgentType(value)
	for _, existing := range AllowedMicrosoftGraphManagementAgentTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MicrosoftGraphManagementAgentType", value)
}

// NewMicrosoftGraphManagementAgentTypeFromValue returns a pointer to a valid MicrosoftGraphManagementAgentType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMicrosoftGraphManagementAgentTypeFromValue(v string) (*MicrosoftGraphManagementAgentType, error) {
	ev := MicrosoftGraphManagementAgentType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MicrosoftGraphManagementAgentType: valid values are %v", v, AllowedMicrosoftGraphManagementAgentTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MicrosoftGraphManagementAgentType) IsValid() bool {
	for _, existing := range AllowedMicrosoftGraphManagementAgentTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to microsoft.graph.managementAgentType value
func (v MicrosoftGraphManagementAgentType) Ptr() *MicrosoftGraphManagementAgentType {
	return &v
}

type NullableMicrosoftGraphManagementAgentType struct {
	value *MicrosoftGraphManagementAgentType
	isSet bool
}

func (v NullableMicrosoftGraphManagementAgentType) Get() *MicrosoftGraphManagementAgentType {
	return v.value
}

func (v *NullableMicrosoftGraphManagementAgentType) Set(val *MicrosoftGraphManagementAgentType) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphManagementAgentType) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphManagementAgentType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphManagementAgentType(val *MicrosoftGraphManagementAgentType) *NullableMicrosoftGraphManagementAgentType {
	return &NullableMicrosoftGraphManagementAgentType{value: val, isSet: true}
}

func (v NullableMicrosoftGraphManagementAgentType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphManagementAgentType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
