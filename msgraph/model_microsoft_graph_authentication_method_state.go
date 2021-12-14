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

// MicrosoftGraphAuthenticationMethodState the model 'MicrosoftGraphAuthenticationMethodState'
type MicrosoftGraphAuthenticationMethodState string

// List of microsoft.graph.authenticationMethodState
const (
	ENABLED MicrosoftGraphAuthenticationMethodState = "enabled"
	DISABLED MicrosoftGraphAuthenticationMethodState = "disabled"
)

// All allowed values of MicrosoftGraphAuthenticationMethodState enum
var AllowedMicrosoftGraphAuthenticationMethodStateEnumValues = []MicrosoftGraphAuthenticationMethodState{
	"enabled",
	"disabled",
}

func (v *MicrosoftGraphAuthenticationMethodState) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MicrosoftGraphAuthenticationMethodState(value)
	for _, existing := range AllowedMicrosoftGraphAuthenticationMethodStateEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MicrosoftGraphAuthenticationMethodState", value)
}

// NewMicrosoftGraphAuthenticationMethodStateFromValue returns a pointer to a valid MicrosoftGraphAuthenticationMethodState
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMicrosoftGraphAuthenticationMethodStateFromValue(v string) (*MicrosoftGraphAuthenticationMethodState, error) {
	ev := MicrosoftGraphAuthenticationMethodState(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MicrosoftGraphAuthenticationMethodState: valid values are %v", v, AllowedMicrosoftGraphAuthenticationMethodStateEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MicrosoftGraphAuthenticationMethodState) IsValid() bool {
	for _, existing := range AllowedMicrosoftGraphAuthenticationMethodStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to microsoft.graph.authenticationMethodState value
func (v MicrosoftGraphAuthenticationMethodState) Ptr() *MicrosoftGraphAuthenticationMethodState {
	return &v
}

type NullableMicrosoftGraphAuthenticationMethodState struct {
	value *MicrosoftGraphAuthenticationMethodState
	isSet bool
}

func (v NullableMicrosoftGraphAuthenticationMethodState) Get() *MicrosoftGraphAuthenticationMethodState {
	return v.value
}

func (v *NullableMicrosoftGraphAuthenticationMethodState) Set(val *MicrosoftGraphAuthenticationMethodState) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphAuthenticationMethodState) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphAuthenticationMethodState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphAuthenticationMethodState(val *MicrosoftGraphAuthenticationMethodState) *NullableMicrosoftGraphAuthenticationMethodState {
	return &NullableMicrosoftGraphAuthenticationMethodState{value: val, isSet: true}
}

func (v NullableMicrosoftGraphAuthenticationMethodState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphAuthenticationMethodState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

