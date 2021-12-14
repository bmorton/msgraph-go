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

// MicrosoftGraphFilterMode the model 'MicrosoftGraphFilterMode'
type MicrosoftGraphFilterMode string

// List of microsoft.graph.filterMode
const (
	INCLUDE MicrosoftGraphFilterMode = "include"
	EXCLUDE MicrosoftGraphFilterMode = "exclude"
)

// All allowed values of MicrosoftGraphFilterMode enum
var AllowedMicrosoftGraphFilterModeEnumValues = []MicrosoftGraphFilterMode{
	"include",
	"exclude",
}

func (v *MicrosoftGraphFilterMode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MicrosoftGraphFilterMode(value)
	for _, existing := range AllowedMicrosoftGraphFilterModeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MicrosoftGraphFilterMode", value)
}

// NewMicrosoftGraphFilterModeFromValue returns a pointer to a valid MicrosoftGraphFilterMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMicrosoftGraphFilterModeFromValue(v string) (*MicrosoftGraphFilterMode, error) {
	ev := MicrosoftGraphFilterMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MicrosoftGraphFilterMode: valid values are %v", v, AllowedMicrosoftGraphFilterModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MicrosoftGraphFilterMode) IsValid() bool {
	for _, existing := range AllowedMicrosoftGraphFilterModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to microsoft.graph.filterMode value
func (v MicrosoftGraphFilterMode) Ptr() *MicrosoftGraphFilterMode {
	return &v
}

type NullableMicrosoftGraphFilterMode struct {
	value *MicrosoftGraphFilterMode
	isSet bool
}

func (v NullableMicrosoftGraphFilterMode) Get() *MicrosoftGraphFilterMode {
	return v.value
}

func (v *NullableMicrosoftGraphFilterMode) Set(val *MicrosoftGraphFilterMode) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphFilterMode) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphFilterMode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphFilterMode(val *MicrosoftGraphFilterMode) *NullableMicrosoftGraphFilterMode {
	return &NullableMicrosoftGraphFilterMode{value: val, isSet: true}
}

func (v NullableMicrosoftGraphFilterMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphFilterMode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

