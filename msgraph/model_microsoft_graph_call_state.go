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

// MicrosoftGraphCallState the model 'MicrosoftGraphCallState'
type MicrosoftGraphCallState string

// List of microsoft.graph.callState
const (
	INCOMING MicrosoftGraphCallState = "incoming"
	ESTABLISHING MicrosoftGraphCallState = "establishing"
	ESTABLISHED MicrosoftGraphCallState = "established"
	HOLD MicrosoftGraphCallState = "hold"
	TRANSFERRING MicrosoftGraphCallState = "transferring"
	TRANSFER_ACCEPTED MicrosoftGraphCallState = "transferAccepted"
	REDIRECTING MicrosoftGraphCallState = "redirecting"
	TERMINATING MicrosoftGraphCallState = "terminating"
	TERMINATED MicrosoftGraphCallState = "terminated"
	UNKNOWN_FUTURE_VALUE MicrosoftGraphCallState = "unknownFutureValue"
)

// All allowed values of MicrosoftGraphCallState enum
var AllowedMicrosoftGraphCallStateEnumValues = []MicrosoftGraphCallState{
	"incoming",
	"establishing",
	"established",
	"hold",
	"transferring",
	"transferAccepted",
	"redirecting",
	"terminating",
	"terminated",
	"unknownFutureValue",
}

func (v *MicrosoftGraphCallState) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MicrosoftGraphCallState(value)
	for _, existing := range AllowedMicrosoftGraphCallStateEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MicrosoftGraphCallState", value)
}

// NewMicrosoftGraphCallStateFromValue returns a pointer to a valid MicrosoftGraphCallState
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMicrosoftGraphCallStateFromValue(v string) (*MicrosoftGraphCallState, error) {
	ev := MicrosoftGraphCallState(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MicrosoftGraphCallState: valid values are %v", v, AllowedMicrosoftGraphCallStateEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MicrosoftGraphCallState) IsValid() bool {
	for _, existing := range AllowedMicrosoftGraphCallStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to microsoft.graph.callState value
func (v MicrosoftGraphCallState) Ptr() *MicrosoftGraphCallState {
	return &v
}

type NullableMicrosoftGraphCallState struct {
	value *MicrosoftGraphCallState
	isSet bool
}

func (v NullableMicrosoftGraphCallState) Get() *MicrosoftGraphCallState {
	return v.value
}

func (v *NullableMicrosoftGraphCallState) Set(val *MicrosoftGraphCallState) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphCallState) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphCallState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphCallState(val *MicrosoftGraphCallState) *NullableMicrosoftGraphCallState {
	return &NullableMicrosoftGraphCallState{value: val, isSet: true}
}

func (v NullableMicrosoftGraphCallState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphCallState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

