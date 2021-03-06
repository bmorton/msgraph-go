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

// MicrosoftGraphEducationExternalSource the model 'MicrosoftGraphEducationExternalSource'
type MicrosoftGraphEducationExternalSource string

// List of microsoft.graph.educationExternalSource
const (
	SIS MicrosoftGraphEducationExternalSource = "sis"
	MANUAL MicrosoftGraphEducationExternalSource = "manual"
	UNKNOWN_FUTURE_VALUE MicrosoftGraphEducationExternalSource = "unknownFutureValue"
)

// All allowed values of MicrosoftGraphEducationExternalSource enum
var AllowedMicrosoftGraphEducationExternalSourceEnumValues = []MicrosoftGraphEducationExternalSource{
	"sis",
	"manual",
	"unknownFutureValue",
}

func (v *MicrosoftGraphEducationExternalSource) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MicrosoftGraphEducationExternalSource(value)
	for _, existing := range AllowedMicrosoftGraphEducationExternalSourceEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MicrosoftGraphEducationExternalSource", value)
}

// NewMicrosoftGraphEducationExternalSourceFromValue returns a pointer to a valid MicrosoftGraphEducationExternalSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMicrosoftGraphEducationExternalSourceFromValue(v string) (*MicrosoftGraphEducationExternalSource, error) {
	ev := MicrosoftGraphEducationExternalSource(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MicrosoftGraphEducationExternalSource: valid values are %v", v, AllowedMicrosoftGraphEducationExternalSourceEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MicrosoftGraphEducationExternalSource) IsValid() bool {
	for _, existing := range AllowedMicrosoftGraphEducationExternalSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to microsoft.graph.educationExternalSource value
func (v MicrosoftGraphEducationExternalSource) Ptr() *MicrosoftGraphEducationExternalSource {
	return &v
}

type NullableMicrosoftGraphEducationExternalSource struct {
	value *MicrosoftGraphEducationExternalSource
	isSet bool
}

func (v NullableMicrosoftGraphEducationExternalSource) Get() *MicrosoftGraphEducationExternalSource {
	return v.value
}

func (v *NullableMicrosoftGraphEducationExternalSource) Set(val *MicrosoftGraphEducationExternalSource) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphEducationExternalSource) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphEducationExternalSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphEducationExternalSource(val *MicrosoftGraphEducationExternalSource) *NullableMicrosoftGraphEducationExternalSource {
	return &NullableMicrosoftGraphEducationExternalSource{value: val, isSet: true}
}

func (v NullableMicrosoftGraphEducationExternalSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphEducationExternalSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

