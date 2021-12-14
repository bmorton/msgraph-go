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

// MicrosoftGraphApplicationType Possible types of Application
type MicrosoftGraphApplicationType string

// List of microsoft.graph.applicationType
const (
	UNIVERSAL MicrosoftGraphApplicationType = "universal"
	DESKTOP MicrosoftGraphApplicationType = "desktop"
)

// All allowed values of MicrosoftGraphApplicationType enum
var AllowedMicrosoftGraphApplicationTypeEnumValues = []MicrosoftGraphApplicationType{
	"universal",
	"desktop",
}

func (v *MicrosoftGraphApplicationType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MicrosoftGraphApplicationType(value)
	for _, existing := range AllowedMicrosoftGraphApplicationTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MicrosoftGraphApplicationType", value)
}

// NewMicrosoftGraphApplicationTypeFromValue returns a pointer to a valid MicrosoftGraphApplicationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMicrosoftGraphApplicationTypeFromValue(v string) (*MicrosoftGraphApplicationType, error) {
	ev := MicrosoftGraphApplicationType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MicrosoftGraphApplicationType: valid values are %v", v, AllowedMicrosoftGraphApplicationTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MicrosoftGraphApplicationType) IsValid() bool {
	for _, existing := range AllowedMicrosoftGraphApplicationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to microsoft.graph.applicationType value
func (v MicrosoftGraphApplicationType) Ptr() *MicrosoftGraphApplicationType {
	return &v
}

type NullableMicrosoftGraphApplicationType struct {
	value *MicrosoftGraphApplicationType
	isSet bool
}

func (v NullableMicrosoftGraphApplicationType) Get() *MicrosoftGraphApplicationType {
	return v.value
}

func (v *NullableMicrosoftGraphApplicationType) Set(val *MicrosoftGraphApplicationType) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphApplicationType) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphApplicationType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphApplicationType(val *MicrosoftGraphApplicationType) *NullableMicrosoftGraphApplicationType {
	return &NullableMicrosoftGraphApplicationType{value: val, isSet: true}
}

func (v NullableMicrosoftGraphApplicationType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphApplicationType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
