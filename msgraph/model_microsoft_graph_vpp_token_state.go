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

// MicrosoftGraphVppTokenState Possible states associated with an Apple Volume Purchase Program token.
type MicrosoftGraphVppTokenState string

// List of microsoft.graph.vppTokenState
const (
	UNKNOWN MicrosoftGraphVppTokenState = "unknown"
	VALID MicrosoftGraphVppTokenState = "valid"
	EXPIRED MicrosoftGraphVppTokenState = "expired"
	INVALID MicrosoftGraphVppTokenState = "invalid"
	ASSIGNED_TO_EXTERNAL_MDM MicrosoftGraphVppTokenState = "assignedToExternalMDM"
)

// All allowed values of MicrosoftGraphVppTokenState enum
var AllowedMicrosoftGraphVppTokenStateEnumValues = []MicrosoftGraphVppTokenState{
	"unknown",
	"valid",
	"expired",
	"invalid",
	"assignedToExternalMDM",
}

func (v *MicrosoftGraphVppTokenState) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MicrosoftGraphVppTokenState(value)
	for _, existing := range AllowedMicrosoftGraphVppTokenStateEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MicrosoftGraphVppTokenState", value)
}

// NewMicrosoftGraphVppTokenStateFromValue returns a pointer to a valid MicrosoftGraphVppTokenState
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMicrosoftGraphVppTokenStateFromValue(v string) (*MicrosoftGraphVppTokenState, error) {
	ev := MicrosoftGraphVppTokenState(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MicrosoftGraphVppTokenState: valid values are %v", v, AllowedMicrosoftGraphVppTokenStateEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MicrosoftGraphVppTokenState) IsValid() bool {
	for _, existing := range AllowedMicrosoftGraphVppTokenStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to microsoft.graph.vppTokenState value
func (v MicrosoftGraphVppTokenState) Ptr() *MicrosoftGraphVppTokenState {
	return &v
}

type NullableMicrosoftGraphVppTokenState struct {
	value *MicrosoftGraphVppTokenState
	isSet bool
}

func (v NullableMicrosoftGraphVppTokenState) Get() *MicrosoftGraphVppTokenState {
	return v.value
}

func (v *NullableMicrosoftGraphVppTokenState) Set(val *MicrosoftGraphVppTokenState) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphVppTokenState) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphVppTokenState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphVppTokenState(val *MicrosoftGraphVppTokenState) *NullableMicrosoftGraphVppTokenState {
	return &NullableMicrosoftGraphVppTokenState{value: val, isSet: true}
}

func (v NullableMicrosoftGraphVppTokenState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphVppTokenState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

