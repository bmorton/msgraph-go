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

// MicrosoftGraphGroupType the model 'MicrosoftGraphGroupType'
type MicrosoftGraphGroupType string

// List of microsoft.graph.groupType
const (
	UNIFIED_GROUPS MicrosoftGraphGroupType = "unifiedGroups"
	AZURE_AD MicrosoftGraphGroupType = "azureAD"
	UNKNOWN_FUTURE_VALUE MicrosoftGraphGroupType = "unknownFutureValue"
)

// All allowed values of MicrosoftGraphGroupType enum
var AllowedMicrosoftGraphGroupTypeEnumValues = []MicrosoftGraphGroupType{
	"unifiedGroups",
	"azureAD",
	"unknownFutureValue",
}

func (v *MicrosoftGraphGroupType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MicrosoftGraphGroupType(value)
	for _, existing := range AllowedMicrosoftGraphGroupTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MicrosoftGraphGroupType", value)
}

// NewMicrosoftGraphGroupTypeFromValue returns a pointer to a valid MicrosoftGraphGroupType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMicrosoftGraphGroupTypeFromValue(v string) (*MicrosoftGraphGroupType, error) {
	ev := MicrosoftGraphGroupType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MicrosoftGraphGroupType: valid values are %v", v, AllowedMicrosoftGraphGroupTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MicrosoftGraphGroupType) IsValid() bool {
	for _, existing := range AllowedMicrosoftGraphGroupTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to microsoft.graph.groupType value
func (v MicrosoftGraphGroupType) Ptr() *MicrosoftGraphGroupType {
	return &v
}

type NullableMicrosoftGraphGroupType struct {
	value *MicrosoftGraphGroupType
	isSet bool
}

func (v NullableMicrosoftGraphGroupType) Get() *MicrosoftGraphGroupType {
	return v.value
}

func (v *NullableMicrosoftGraphGroupType) Set(val *MicrosoftGraphGroupType) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphGroupType) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphGroupType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphGroupType(val *MicrosoftGraphGroupType) *NullableMicrosoftGraphGroupType {
	return &NullableMicrosoftGraphGroupType{value: val, isSet: true}
}

func (v NullableMicrosoftGraphGroupType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphGroupType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

