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

// MicrosoftGraphWindowsInformationProtectionPinCharacterRequirements Pin Character Requirements
type MicrosoftGraphWindowsInformationProtectionPinCharacterRequirements string

// List of microsoft.graph.windowsInformationProtectionPinCharacterRequirements
const (
	NOT_ALLOW MicrosoftGraphWindowsInformationProtectionPinCharacterRequirements = "notAllow"
	REQUIRE_AT_LEAST_ONE MicrosoftGraphWindowsInformationProtectionPinCharacterRequirements = "requireAtLeastOne"
	ALLOW MicrosoftGraphWindowsInformationProtectionPinCharacterRequirements = "allow"
)

// All allowed values of MicrosoftGraphWindowsInformationProtectionPinCharacterRequirements enum
var AllowedMicrosoftGraphWindowsInformationProtectionPinCharacterRequirementsEnumValues = []MicrosoftGraphWindowsInformationProtectionPinCharacterRequirements{
	"notAllow",
	"requireAtLeastOne",
	"allow",
}

func (v *MicrosoftGraphWindowsInformationProtectionPinCharacterRequirements) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MicrosoftGraphWindowsInformationProtectionPinCharacterRequirements(value)
	for _, existing := range AllowedMicrosoftGraphWindowsInformationProtectionPinCharacterRequirementsEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MicrosoftGraphWindowsInformationProtectionPinCharacterRequirements", value)
}

// NewMicrosoftGraphWindowsInformationProtectionPinCharacterRequirementsFromValue returns a pointer to a valid MicrosoftGraphWindowsInformationProtectionPinCharacterRequirements
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMicrosoftGraphWindowsInformationProtectionPinCharacterRequirementsFromValue(v string) (*MicrosoftGraphWindowsInformationProtectionPinCharacterRequirements, error) {
	ev := MicrosoftGraphWindowsInformationProtectionPinCharacterRequirements(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MicrosoftGraphWindowsInformationProtectionPinCharacterRequirements: valid values are %v", v, AllowedMicrosoftGraphWindowsInformationProtectionPinCharacterRequirementsEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MicrosoftGraphWindowsInformationProtectionPinCharacterRequirements) IsValid() bool {
	for _, existing := range AllowedMicrosoftGraphWindowsInformationProtectionPinCharacterRequirementsEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to microsoft.graph.windowsInformationProtectionPinCharacterRequirements value
func (v MicrosoftGraphWindowsInformationProtectionPinCharacterRequirements) Ptr() *MicrosoftGraphWindowsInformationProtectionPinCharacterRequirements {
	return &v
}

type NullableMicrosoftGraphWindowsInformationProtectionPinCharacterRequirements struct {
	value *MicrosoftGraphWindowsInformationProtectionPinCharacterRequirements
	isSet bool
}

func (v NullableMicrosoftGraphWindowsInformationProtectionPinCharacterRequirements) Get() *MicrosoftGraphWindowsInformationProtectionPinCharacterRequirements {
	return v.value
}

func (v *NullableMicrosoftGraphWindowsInformationProtectionPinCharacterRequirements) Set(val *MicrosoftGraphWindowsInformationProtectionPinCharacterRequirements) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphWindowsInformationProtectionPinCharacterRequirements) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphWindowsInformationProtectionPinCharacterRequirements) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphWindowsInformationProtectionPinCharacterRequirements(val *MicrosoftGraphWindowsInformationProtectionPinCharacterRequirements) *NullableMicrosoftGraphWindowsInformationProtectionPinCharacterRequirements {
	return &NullableMicrosoftGraphWindowsInformationProtectionPinCharacterRequirements{value: val, isSet: true}
}

func (v NullableMicrosoftGraphWindowsInformationProtectionPinCharacterRequirements) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphWindowsInformationProtectionPinCharacterRequirements) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

