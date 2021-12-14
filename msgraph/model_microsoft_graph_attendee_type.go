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

// MicrosoftGraphAttendeeType the model 'MicrosoftGraphAttendeeType'
type MicrosoftGraphAttendeeType string

// List of microsoft.graph.attendeeType
const (
	REQUIRED MicrosoftGraphAttendeeType = "required"
	OPTIONAL MicrosoftGraphAttendeeType = "optional"
	RESOURCE MicrosoftGraphAttendeeType = "resource"
)

// All allowed values of MicrosoftGraphAttendeeType enum
var AllowedMicrosoftGraphAttendeeTypeEnumValues = []MicrosoftGraphAttendeeType{
	"required",
	"optional",
	"resource",
}

func (v *MicrosoftGraphAttendeeType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MicrosoftGraphAttendeeType(value)
	for _, existing := range AllowedMicrosoftGraphAttendeeTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MicrosoftGraphAttendeeType", value)
}

// NewMicrosoftGraphAttendeeTypeFromValue returns a pointer to a valid MicrosoftGraphAttendeeType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMicrosoftGraphAttendeeTypeFromValue(v string) (*MicrosoftGraphAttendeeType, error) {
	ev := MicrosoftGraphAttendeeType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MicrosoftGraphAttendeeType: valid values are %v", v, AllowedMicrosoftGraphAttendeeTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MicrosoftGraphAttendeeType) IsValid() bool {
	for _, existing := range AllowedMicrosoftGraphAttendeeTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to microsoft.graph.attendeeType value
func (v MicrosoftGraphAttendeeType) Ptr() *MicrosoftGraphAttendeeType {
	return &v
}

type NullableMicrosoftGraphAttendeeType struct {
	value *MicrosoftGraphAttendeeType
	isSet bool
}

func (v NullableMicrosoftGraphAttendeeType) Get() *MicrosoftGraphAttendeeType {
	return v.value
}

func (v *NullableMicrosoftGraphAttendeeType) Set(val *MicrosoftGraphAttendeeType) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphAttendeeType) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphAttendeeType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphAttendeeType(val *MicrosoftGraphAttendeeType) *NullableMicrosoftGraphAttendeeType {
	return &NullableMicrosoftGraphAttendeeType{value: val, isSet: true}
}

func (v NullableMicrosoftGraphAttendeeType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphAttendeeType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

