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

// MicrosoftGraphConditionalAccessPolicyState the model 'MicrosoftGraphConditionalAccessPolicyState'
type MicrosoftGraphConditionalAccessPolicyState string

// List of microsoft.graph.conditionalAccessPolicyState
const (
	ENABLED MicrosoftGraphConditionalAccessPolicyState = "enabled"
	DISABLED MicrosoftGraphConditionalAccessPolicyState = "disabled"
	ENABLED_FOR_REPORTING_BUT_NOT_ENFORCED MicrosoftGraphConditionalAccessPolicyState = "enabledForReportingButNotEnforced"
)

// All allowed values of MicrosoftGraphConditionalAccessPolicyState enum
var AllowedMicrosoftGraphConditionalAccessPolicyStateEnumValues = []MicrosoftGraphConditionalAccessPolicyState{
	"enabled",
	"disabled",
	"enabledForReportingButNotEnforced",
}

func (v *MicrosoftGraphConditionalAccessPolicyState) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MicrosoftGraphConditionalAccessPolicyState(value)
	for _, existing := range AllowedMicrosoftGraphConditionalAccessPolicyStateEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MicrosoftGraphConditionalAccessPolicyState", value)
}

// NewMicrosoftGraphConditionalAccessPolicyStateFromValue returns a pointer to a valid MicrosoftGraphConditionalAccessPolicyState
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMicrosoftGraphConditionalAccessPolicyStateFromValue(v string) (*MicrosoftGraphConditionalAccessPolicyState, error) {
	ev := MicrosoftGraphConditionalAccessPolicyState(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MicrosoftGraphConditionalAccessPolicyState: valid values are %v", v, AllowedMicrosoftGraphConditionalAccessPolicyStateEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MicrosoftGraphConditionalAccessPolicyState) IsValid() bool {
	for _, existing := range AllowedMicrosoftGraphConditionalAccessPolicyStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to microsoft.graph.conditionalAccessPolicyState value
func (v MicrosoftGraphConditionalAccessPolicyState) Ptr() *MicrosoftGraphConditionalAccessPolicyState {
	return &v
}

type NullableMicrosoftGraphConditionalAccessPolicyState struct {
	value *MicrosoftGraphConditionalAccessPolicyState
	isSet bool
}

func (v NullableMicrosoftGraphConditionalAccessPolicyState) Get() *MicrosoftGraphConditionalAccessPolicyState {
	return v.value
}

func (v *NullableMicrosoftGraphConditionalAccessPolicyState) Set(val *MicrosoftGraphConditionalAccessPolicyState) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphConditionalAccessPolicyState) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphConditionalAccessPolicyState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphConditionalAccessPolicyState(val *MicrosoftGraphConditionalAccessPolicyState) *NullableMicrosoftGraphConditionalAccessPolicyState {
	return &NullableMicrosoftGraphConditionalAccessPolicyState{value: val, isSet: true}
}

func (v NullableMicrosoftGraphConditionalAccessPolicyState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphConditionalAccessPolicyState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

