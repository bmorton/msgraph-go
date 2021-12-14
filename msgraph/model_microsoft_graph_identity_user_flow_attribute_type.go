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

// MicrosoftGraphIdentityUserFlowAttributeType the model 'MicrosoftGraphIdentityUserFlowAttributeType'
type MicrosoftGraphIdentityUserFlowAttributeType string

// List of microsoft.graph.identityUserFlowAttributeType
const (
	BUILT_IN MicrosoftGraphIdentityUserFlowAttributeType = "builtIn"
	CUSTOM MicrosoftGraphIdentityUserFlowAttributeType = "custom"
	REQUIRED MicrosoftGraphIdentityUserFlowAttributeType = "required"
	UNKNOWN_FUTURE_VALUE MicrosoftGraphIdentityUserFlowAttributeType = "unknownFutureValue"
)

// All allowed values of MicrosoftGraphIdentityUserFlowAttributeType enum
var AllowedMicrosoftGraphIdentityUserFlowAttributeTypeEnumValues = []MicrosoftGraphIdentityUserFlowAttributeType{
	"builtIn",
	"custom",
	"required",
	"unknownFutureValue",
}

func (v *MicrosoftGraphIdentityUserFlowAttributeType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MicrosoftGraphIdentityUserFlowAttributeType(value)
	for _, existing := range AllowedMicrosoftGraphIdentityUserFlowAttributeTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MicrosoftGraphIdentityUserFlowAttributeType", value)
}

// NewMicrosoftGraphIdentityUserFlowAttributeTypeFromValue returns a pointer to a valid MicrosoftGraphIdentityUserFlowAttributeType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMicrosoftGraphIdentityUserFlowAttributeTypeFromValue(v string) (*MicrosoftGraphIdentityUserFlowAttributeType, error) {
	ev := MicrosoftGraphIdentityUserFlowAttributeType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MicrosoftGraphIdentityUserFlowAttributeType: valid values are %v", v, AllowedMicrosoftGraphIdentityUserFlowAttributeTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MicrosoftGraphIdentityUserFlowAttributeType) IsValid() bool {
	for _, existing := range AllowedMicrosoftGraphIdentityUserFlowAttributeTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to microsoft.graph.identityUserFlowAttributeType value
func (v MicrosoftGraphIdentityUserFlowAttributeType) Ptr() *MicrosoftGraphIdentityUserFlowAttributeType {
	return &v
}

type NullableMicrosoftGraphIdentityUserFlowAttributeType struct {
	value *MicrosoftGraphIdentityUserFlowAttributeType
	isSet bool
}

func (v NullableMicrosoftGraphIdentityUserFlowAttributeType) Get() *MicrosoftGraphIdentityUserFlowAttributeType {
	return v.value
}

func (v *NullableMicrosoftGraphIdentityUserFlowAttributeType) Set(val *MicrosoftGraphIdentityUserFlowAttributeType) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphIdentityUserFlowAttributeType) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphIdentityUserFlowAttributeType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphIdentityUserFlowAttributeType(val *MicrosoftGraphIdentityUserFlowAttributeType) *NullableMicrosoftGraphIdentityUserFlowAttributeType {
	return &NullableMicrosoftGraphIdentityUserFlowAttributeType{value: val, isSet: true}
}

func (v NullableMicrosoftGraphIdentityUserFlowAttributeType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphIdentityUserFlowAttributeType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
