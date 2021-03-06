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

// MicrosoftGraphAttestationLevel the model 'MicrosoftGraphAttestationLevel'
type MicrosoftGraphAttestationLevel string

// List of microsoft.graph.attestationLevel
const (
	ATTESTED MicrosoftGraphAttestationLevel = "attested"
	NOT_ATTESTED MicrosoftGraphAttestationLevel = "notAttested"
	UNKNOWN_FUTURE_VALUE MicrosoftGraphAttestationLevel = "unknownFutureValue"
)

// All allowed values of MicrosoftGraphAttestationLevel enum
var AllowedMicrosoftGraphAttestationLevelEnumValues = []MicrosoftGraphAttestationLevel{
	"attested",
	"notAttested",
	"unknownFutureValue",
}

func (v *MicrosoftGraphAttestationLevel) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MicrosoftGraphAttestationLevel(value)
	for _, existing := range AllowedMicrosoftGraphAttestationLevelEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MicrosoftGraphAttestationLevel", value)
}

// NewMicrosoftGraphAttestationLevelFromValue returns a pointer to a valid MicrosoftGraphAttestationLevel
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMicrosoftGraphAttestationLevelFromValue(v string) (*MicrosoftGraphAttestationLevel, error) {
	ev := MicrosoftGraphAttestationLevel(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MicrosoftGraphAttestationLevel: valid values are %v", v, AllowedMicrosoftGraphAttestationLevelEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MicrosoftGraphAttestationLevel) IsValid() bool {
	for _, existing := range AllowedMicrosoftGraphAttestationLevelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to microsoft.graph.attestationLevel value
func (v MicrosoftGraphAttestationLevel) Ptr() *MicrosoftGraphAttestationLevel {
	return &v
}

type NullableMicrosoftGraphAttestationLevel struct {
	value *MicrosoftGraphAttestationLevel
	isSet bool
}

func (v NullableMicrosoftGraphAttestationLevel) Get() *MicrosoftGraphAttestationLevel {
	return v.value
}

func (v *NullableMicrosoftGraphAttestationLevel) Set(val *MicrosoftGraphAttestationLevel) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphAttestationLevel) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphAttestationLevel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphAttestationLevel(val *MicrosoftGraphAttestationLevel) *NullableMicrosoftGraphAttestationLevel {
	return &NullableMicrosoftGraphAttestationLevel{value: val, isSet: true}
}

func (v NullableMicrosoftGraphAttestationLevel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphAttestationLevel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

