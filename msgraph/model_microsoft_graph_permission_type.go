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

// MicrosoftGraphPermissionType the model 'MicrosoftGraphPermissionType'
type MicrosoftGraphPermissionType string

// List of microsoft.graph.permissionType
const (
	APPLICATION MicrosoftGraphPermissionType = "application"
	DELEGATED MicrosoftGraphPermissionType = "delegated"
	DELEGATED_USER_CONSENTABLE MicrosoftGraphPermissionType = "delegatedUserConsentable"
)

// All allowed values of MicrosoftGraphPermissionType enum
var AllowedMicrosoftGraphPermissionTypeEnumValues = []MicrosoftGraphPermissionType{
	"application",
	"delegated",
	"delegatedUserConsentable",
}

func (v *MicrosoftGraphPermissionType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MicrosoftGraphPermissionType(value)
	for _, existing := range AllowedMicrosoftGraphPermissionTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MicrosoftGraphPermissionType", value)
}

// NewMicrosoftGraphPermissionTypeFromValue returns a pointer to a valid MicrosoftGraphPermissionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMicrosoftGraphPermissionTypeFromValue(v string) (*MicrosoftGraphPermissionType, error) {
	ev := MicrosoftGraphPermissionType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MicrosoftGraphPermissionType: valid values are %v", v, AllowedMicrosoftGraphPermissionTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MicrosoftGraphPermissionType) IsValid() bool {
	for _, existing := range AllowedMicrosoftGraphPermissionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to microsoft.graph.permissionType value
func (v MicrosoftGraphPermissionType) Ptr() *MicrosoftGraphPermissionType {
	return &v
}

type NullableMicrosoftGraphPermissionType struct {
	value *MicrosoftGraphPermissionType
	isSet bool
}

func (v NullableMicrosoftGraphPermissionType) Get() *MicrosoftGraphPermissionType {
	return v.value
}

func (v *NullableMicrosoftGraphPermissionType) Set(val *MicrosoftGraphPermissionType) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphPermissionType) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphPermissionType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphPermissionType(val *MicrosoftGraphPermissionType) *NullableMicrosoftGraphPermissionType {
	return &NullableMicrosoftGraphPermissionType{value: val, isSet: true}
}

func (v NullableMicrosoftGraphPermissionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphPermissionType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
