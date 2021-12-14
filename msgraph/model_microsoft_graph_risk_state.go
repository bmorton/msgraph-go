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

// MicrosoftGraphRiskState the model 'MicrosoftGraphRiskState'
type MicrosoftGraphRiskState string

// List of microsoft.graph.riskState
const (
	NONE MicrosoftGraphRiskState = "none"
	CONFIRMED_SAFE MicrosoftGraphRiskState = "confirmedSafe"
	REMEDIATED MicrosoftGraphRiskState = "remediated"
	DISMISSED MicrosoftGraphRiskState = "dismissed"
	AT_RISK MicrosoftGraphRiskState = "atRisk"
	CONFIRMED_COMPROMISED MicrosoftGraphRiskState = "confirmedCompromised"
	UNKNOWN_FUTURE_VALUE MicrosoftGraphRiskState = "unknownFutureValue"
)

// All allowed values of MicrosoftGraphRiskState enum
var AllowedMicrosoftGraphRiskStateEnumValues = []MicrosoftGraphRiskState{
	"none",
	"confirmedSafe",
	"remediated",
	"dismissed",
	"atRisk",
	"confirmedCompromised",
	"unknownFutureValue",
}

func (v *MicrosoftGraphRiskState) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MicrosoftGraphRiskState(value)
	for _, existing := range AllowedMicrosoftGraphRiskStateEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MicrosoftGraphRiskState", value)
}

// NewMicrosoftGraphRiskStateFromValue returns a pointer to a valid MicrosoftGraphRiskState
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMicrosoftGraphRiskStateFromValue(v string) (*MicrosoftGraphRiskState, error) {
	ev := MicrosoftGraphRiskState(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MicrosoftGraphRiskState: valid values are %v", v, AllowedMicrosoftGraphRiskStateEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MicrosoftGraphRiskState) IsValid() bool {
	for _, existing := range AllowedMicrosoftGraphRiskStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to microsoft.graph.riskState value
func (v MicrosoftGraphRiskState) Ptr() *MicrosoftGraphRiskState {
	return &v
}

type NullableMicrosoftGraphRiskState struct {
	value *MicrosoftGraphRiskState
	isSet bool
}

func (v NullableMicrosoftGraphRiskState) Get() *MicrosoftGraphRiskState {
	return v.value
}

func (v *NullableMicrosoftGraphRiskState) Set(val *MicrosoftGraphRiskState) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphRiskState) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphRiskState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphRiskState(val *MicrosoftGraphRiskState) *NullableMicrosoftGraphRiskState {
	return &NullableMicrosoftGraphRiskState{value: val, isSet: true}
}

func (v NullableMicrosoftGraphRiskState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphRiskState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

