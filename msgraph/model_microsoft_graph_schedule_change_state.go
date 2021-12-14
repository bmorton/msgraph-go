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

// MicrosoftGraphScheduleChangeState the model 'MicrosoftGraphScheduleChangeState'
type MicrosoftGraphScheduleChangeState string

// List of microsoft.graph.scheduleChangeState
const (
	PENDING MicrosoftGraphScheduleChangeState = "pending"
	APPROVED MicrosoftGraphScheduleChangeState = "approved"
	DECLINED MicrosoftGraphScheduleChangeState = "declined"
	UNKNOWN_FUTURE_VALUE MicrosoftGraphScheduleChangeState = "unknownFutureValue"
)

// All allowed values of MicrosoftGraphScheduleChangeState enum
var AllowedMicrosoftGraphScheduleChangeStateEnumValues = []MicrosoftGraphScheduleChangeState{
	"pending",
	"approved",
	"declined",
	"unknownFutureValue",
}

func (v *MicrosoftGraphScheduleChangeState) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MicrosoftGraphScheduleChangeState(value)
	for _, existing := range AllowedMicrosoftGraphScheduleChangeStateEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MicrosoftGraphScheduleChangeState", value)
}

// NewMicrosoftGraphScheduleChangeStateFromValue returns a pointer to a valid MicrosoftGraphScheduleChangeState
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMicrosoftGraphScheduleChangeStateFromValue(v string) (*MicrosoftGraphScheduleChangeState, error) {
	ev := MicrosoftGraphScheduleChangeState(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MicrosoftGraphScheduleChangeState: valid values are %v", v, AllowedMicrosoftGraphScheduleChangeStateEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MicrosoftGraphScheduleChangeState) IsValid() bool {
	for _, existing := range AllowedMicrosoftGraphScheduleChangeStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to microsoft.graph.scheduleChangeState value
func (v MicrosoftGraphScheduleChangeState) Ptr() *MicrosoftGraphScheduleChangeState {
	return &v
}

type NullableMicrosoftGraphScheduleChangeState struct {
	value *MicrosoftGraphScheduleChangeState
	isSet bool
}

func (v NullableMicrosoftGraphScheduleChangeState) Get() *MicrosoftGraphScheduleChangeState {
	return v.value
}

func (v *NullableMicrosoftGraphScheduleChangeState) Set(val *MicrosoftGraphScheduleChangeState) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphScheduleChangeState) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphScheduleChangeState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphScheduleChangeState(val *MicrosoftGraphScheduleChangeState) *NullableMicrosoftGraphScheduleChangeState {
	return &NullableMicrosoftGraphScheduleChangeState{value: val, isSet: true}
}

func (v NullableMicrosoftGraphScheduleChangeState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphScheduleChangeState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

