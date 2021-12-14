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

// MicrosoftGraphProcessIntegrityLevel the model 'MicrosoftGraphProcessIntegrityLevel'
type MicrosoftGraphProcessIntegrityLevel string

// List of microsoft.graph.processIntegrityLevel
const (
	UNKNOWN MicrosoftGraphProcessIntegrityLevel = "unknown"
	UNTRUSTED MicrosoftGraphProcessIntegrityLevel = "untrusted"
	LOW MicrosoftGraphProcessIntegrityLevel = "low"
	MEDIUM MicrosoftGraphProcessIntegrityLevel = "medium"
	HIGH MicrosoftGraphProcessIntegrityLevel = "high"
	SYSTEM MicrosoftGraphProcessIntegrityLevel = "system"
	UNKNOWN_FUTURE_VALUE MicrosoftGraphProcessIntegrityLevel = "unknownFutureValue"
)

// All allowed values of MicrosoftGraphProcessIntegrityLevel enum
var AllowedMicrosoftGraphProcessIntegrityLevelEnumValues = []MicrosoftGraphProcessIntegrityLevel{
	"unknown",
	"untrusted",
	"low",
	"medium",
	"high",
	"system",
	"unknownFutureValue",
}

func (v *MicrosoftGraphProcessIntegrityLevel) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MicrosoftGraphProcessIntegrityLevel(value)
	for _, existing := range AllowedMicrosoftGraphProcessIntegrityLevelEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MicrosoftGraphProcessIntegrityLevel", value)
}

// NewMicrosoftGraphProcessIntegrityLevelFromValue returns a pointer to a valid MicrosoftGraphProcessIntegrityLevel
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMicrosoftGraphProcessIntegrityLevelFromValue(v string) (*MicrosoftGraphProcessIntegrityLevel, error) {
	ev := MicrosoftGraphProcessIntegrityLevel(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MicrosoftGraphProcessIntegrityLevel: valid values are %v", v, AllowedMicrosoftGraphProcessIntegrityLevelEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MicrosoftGraphProcessIntegrityLevel) IsValid() bool {
	for _, existing := range AllowedMicrosoftGraphProcessIntegrityLevelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to microsoft.graph.processIntegrityLevel value
func (v MicrosoftGraphProcessIntegrityLevel) Ptr() *MicrosoftGraphProcessIntegrityLevel {
	return &v
}

type NullableMicrosoftGraphProcessIntegrityLevel struct {
	value *MicrosoftGraphProcessIntegrityLevel
	isSet bool
}

func (v NullableMicrosoftGraphProcessIntegrityLevel) Get() *MicrosoftGraphProcessIntegrityLevel {
	return v.value
}

func (v *NullableMicrosoftGraphProcessIntegrityLevel) Set(val *MicrosoftGraphProcessIntegrityLevel) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphProcessIntegrityLevel) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphProcessIntegrityLevel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphProcessIntegrityLevel(val *MicrosoftGraphProcessIntegrityLevel) *NullableMicrosoftGraphProcessIntegrityLevel {
	return &NullableMicrosoftGraphProcessIntegrityLevel{value: val, isSet: true}
}

func (v NullableMicrosoftGraphProcessIntegrityLevel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphProcessIntegrityLevel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

