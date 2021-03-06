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

// MicrosoftGraphWindowsInformationProtectionEnforcementLevel Possible values for WIP Protection enforcement levels
type MicrosoftGraphWindowsInformationProtectionEnforcementLevel string

// List of microsoft.graph.windowsInformationProtectionEnforcementLevel
const (
	NO_PROTECTION MicrosoftGraphWindowsInformationProtectionEnforcementLevel = "noProtection"
	ENCRYPT_AND_AUDIT_ONLY MicrosoftGraphWindowsInformationProtectionEnforcementLevel = "encryptAndAuditOnly"
	ENCRYPT_AUDIT_AND_PROMPT MicrosoftGraphWindowsInformationProtectionEnforcementLevel = "encryptAuditAndPrompt"
	ENCRYPT_AUDIT_AND_BLOCK MicrosoftGraphWindowsInformationProtectionEnforcementLevel = "encryptAuditAndBlock"
)

// All allowed values of MicrosoftGraphWindowsInformationProtectionEnforcementLevel enum
var AllowedMicrosoftGraphWindowsInformationProtectionEnforcementLevelEnumValues = []MicrosoftGraphWindowsInformationProtectionEnforcementLevel{
	"noProtection",
	"encryptAndAuditOnly",
	"encryptAuditAndPrompt",
	"encryptAuditAndBlock",
}

func (v *MicrosoftGraphWindowsInformationProtectionEnforcementLevel) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MicrosoftGraphWindowsInformationProtectionEnforcementLevel(value)
	for _, existing := range AllowedMicrosoftGraphWindowsInformationProtectionEnforcementLevelEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MicrosoftGraphWindowsInformationProtectionEnforcementLevel", value)
}

// NewMicrosoftGraphWindowsInformationProtectionEnforcementLevelFromValue returns a pointer to a valid MicrosoftGraphWindowsInformationProtectionEnforcementLevel
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMicrosoftGraphWindowsInformationProtectionEnforcementLevelFromValue(v string) (*MicrosoftGraphWindowsInformationProtectionEnforcementLevel, error) {
	ev := MicrosoftGraphWindowsInformationProtectionEnforcementLevel(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MicrosoftGraphWindowsInformationProtectionEnforcementLevel: valid values are %v", v, AllowedMicrosoftGraphWindowsInformationProtectionEnforcementLevelEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MicrosoftGraphWindowsInformationProtectionEnforcementLevel) IsValid() bool {
	for _, existing := range AllowedMicrosoftGraphWindowsInformationProtectionEnforcementLevelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to microsoft.graph.windowsInformationProtectionEnforcementLevel value
func (v MicrosoftGraphWindowsInformationProtectionEnforcementLevel) Ptr() *MicrosoftGraphWindowsInformationProtectionEnforcementLevel {
	return &v
}

type NullableMicrosoftGraphWindowsInformationProtectionEnforcementLevel struct {
	value *MicrosoftGraphWindowsInformationProtectionEnforcementLevel
	isSet bool
}

func (v NullableMicrosoftGraphWindowsInformationProtectionEnforcementLevel) Get() *MicrosoftGraphWindowsInformationProtectionEnforcementLevel {
	return v.value
}

func (v *NullableMicrosoftGraphWindowsInformationProtectionEnforcementLevel) Set(val *MicrosoftGraphWindowsInformationProtectionEnforcementLevel) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphWindowsInformationProtectionEnforcementLevel) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphWindowsInformationProtectionEnforcementLevel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphWindowsInformationProtectionEnforcementLevel(val *MicrosoftGraphWindowsInformationProtectionEnforcementLevel) *NullableMicrosoftGraphWindowsInformationProtectionEnforcementLevel {
	return &NullableMicrosoftGraphWindowsInformationProtectionEnforcementLevel{value: val, isSet: true}
}

func (v NullableMicrosoftGraphWindowsInformationProtectionEnforcementLevel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphWindowsInformationProtectionEnforcementLevel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

