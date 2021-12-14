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

// MicrosoftGraphIdentityUserFlowAttributeDataType the model 'MicrosoftGraphIdentityUserFlowAttributeDataType'
type MicrosoftGraphIdentityUserFlowAttributeDataType string

// List of microsoft.graph.identityUserFlowAttributeDataType
const (
	STRING MicrosoftGraphIdentityUserFlowAttributeDataType = "string"
	BOOLEAN MicrosoftGraphIdentityUserFlowAttributeDataType = "boolean"
	INT64 MicrosoftGraphIdentityUserFlowAttributeDataType = "int64"
	STRING_COLLECTION MicrosoftGraphIdentityUserFlowAttributeDataType = "stringCollection"
	DATE_TIME MicrosoftGraphIdentityUserFlowAttributeDataType = "dateTime"
	UNKNOWN_FUTURE_VALUE MicrosoftGraphIdentityUserFlowAttributeDataType = "unknownFutureValue"
)

// All allowed values of MicrosoftGraphIdentityUserFlowAttributeDataType enum
var AllowedMicrosoftGraphIdentityUserFlowAttributeDataTypeEnumValues = []MicrosoftGraphIdentityUserFlowAttributeDataType{
	"string",
	"boolean",
	"int64",
	"stringCollection",
	"dateTime",
	"unknownFutureValue",
}

func (v *MicrosoftGraphIdentityUserFlowAttributeDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MicrosoftGraphIdentityUserFlowAttributeDataType(value)
	for _, existing := range AllowedMicrosoftGraphIdentityUserFlowAttributeDataTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MicrosoftGraphIdentityUserFlowAttributeDataType", value)
}

// NewMicrosoftGraphIdentityUserFlowAttributeDataTypeFromValue returns a pointer to a valid MicrosoftGraphIdentityUserFlowAttributeDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMicrosoftGraphIdentityUserFlowAttributeDataTypeFromValue(v string) (*MicrosoftGraphIdentityUserFlowAttributeDataType, error) {
	ev := MicrosoftGraphIdentityUserFlowAttributeDataType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MicrosoftGraphIdentityUserFlowAttributeDataType: valid values are %v", v, AllowedMicrosoftGraphIdentityUserFlowAttributeDataTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MicrosoftGraphIdentityUserFlowAttributeDataType) IsValid() bool {
	for _, existing := range AllowedMicrosoftGraphIdentityUserFlowAttributeDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to microsoft.graph.identityUserFlowAttributeDataType value
func (v MicrosoftGraphIdentityUserFlowAttributeDataType) Ptr() *MicrosoftGraphIdentityUserFlowAttributeDataType {
	return &v
}

type NullableMicrosoftGraphIdentityUserFlowAttributeDataType struct {
	value *MicrosoftGraphIdentityUserFlowAttributeDataType
	isSet bool
}

func (v NullableMicrosoftGraphIdentityUserFlowAttributeDataType) Get() *MicrosoftGraphIdentityUserFlowAttributeDataType {
	return v.value
}

func (v *NullableMicrosoftGraphIdentityUserFlowAttributeDataType) Set(val *MicrosoftGraphIdentityUserFlowAttributeDataType) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphIdentityUserFlowAttributeDataType) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphIdentityUserFlowAttributeDataType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphIdentityUserFlowAttributeDataType(val *MicrosoftGraphIdentityUserFlowAttributeDataType) *NullableMicrosoftGraphIdentityUserFlowAttributeDataType {
	return &NullableMicrosoftGraphIdentityUserFlowAttributeDataType{value: val, isSet: true}
}

func (v NullableMicrosoftGraphIdentityUserFlowAttributeDataType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphIdentityUserFlowAttributeDataType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
