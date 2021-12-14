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

// MicrosoftGraphActionState State of the action on the device
type MicrosoftGraphActionState string

// List of microsoft.graph.actionState
const (
	NONE MicrosoftGraphActionState = "none"
	PENDING MicrosoftGraphActionState = "pending"
	CANCELED MicrosoftGraphActionState = "canceled"
	ACTIVE MicrosoftGraphActionState = "active"
	DONE MicrosoftGraphActionState = "done"
	FAILED MicrosoftGraphActionState = "failed"
	NOT_SUPPORTED MicrosoftGraphActionState = "notSupported"
)

// All allowed values of MicrosoftGraphActionState enum
var AllowedMicrosoftGraphActionStateEnumValues = []MicrosoftGraphActionState{
	"none",
	"pending",
	"canceled",
	"active",
	"done",
	"failed",
	"notSupported",
}

func (v *MicrosoftGraphActionState) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MicrosoftGraphActionState(value)
	for _, existing := range AllowedMicrosoftGraphActionStateEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MicrosoftGraphActionState", value)
}

// NewMicrosoftGraphActionStateFromValue returns a pointer to a valid MicrosoftGraphActionState
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMicrosoftGraphActionStateFromValue(v string) (*MicrosoftGraphActionState, error) {
	ev := MicrosoftGraphActionState(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MicrosoftGraphActionState: valid values are %v", v, AllowedMicrosoftGraphActionStateEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MicrosoftGraphActionState) IsValid() bool {
	for _, existing := range AllowedMicrosoftGraphActionStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to microsoft.graph.actionState value
func (v MicrosoftGraphActionState) Ptr() *MicrosoftGraphActionState {
	return &v
}

type NullableMicrosoftGraphActionState struct {
	value *MicrosoftGraphActionState
	isSet bool
}

func (v NullableMicrosoftGraphActionState) Get() *MicrosoftGraphActionState {
	return v.value
}

func (v *NullableMicrosoftGraphActionState) Set(val *MicrosoftGraphActionState) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphActionState) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphActionState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphActionState(val *MicrosoftGraphActionState) *NullableMicrosoftGraphActionState {
	return &NullableMicrosoftGraphActionState{value: val, isSet: true}
}

func (v NullableMicrosoftGraphActionState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphActionState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

