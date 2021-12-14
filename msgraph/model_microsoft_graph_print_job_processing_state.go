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

// MicrosoftGraphPrintJobProcessingState the model 'MicrosoftGraphPrintJobProcessingState'
type MicrosoftGraphPrintJobProcessingState string

// List of microsoft.graph.printJobProcessingState
const (
	UNKNOWN MicrosoftGraphPrintJobProcessingState = "unknown"
	PENDING MicrosoftGraphPrintJobProcessingState = "pending"
	PROCESSING MicrosoftGraphPrintJobProcessingState = "processing"
	PAUSED MicrosoftGraphPrintJobProcessingState = "paused"
	STOPPED MicrosoftGraphPrintJobProcessingState = "stopped"
	COMPLETED MicrosoftGraphPrintJobProcessingState = "completed"
	CANCELED MicrosoftGraphPrintJobProcessingState = "canceled"
	ABORTED MicrosoftGraphPrintJobProcessingState = "aborted"
	UNKNOWN_FUTURE_VALUE MicrosoftGraphPrintJobProcessingState = "unknownFutureValue"
)

// All allowed values of MicrosoftGraphPrintJobProcessingState enum
var AllowedMicrosoftGraphPrintJobProcessingStateEnumValues = []MicrosoftGraphPrintJobProcessingState{
	"unknown",
	"pending",
	"processing",
	"paused",
	"stopped",
	"completed",
	"canceled",
	"aborted",
	"unknownFutureValue",
}

func (v *MicrosoftGraphPrintJobProcessingState) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MicrosoftGraphPrintJobProcessingState(value)
	for _, existing := range AllowedMicrosoftGraphPrintJobProcessingStateEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MicrosoftGraphPrintJobProcessingState", value)
}

// NewMicrosoftGraphPrintJobProcessingStateFromValue returns a pointer to a valid MicrosoftGraphPrintJobProcessingState
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMicrosoftGraphPrintJobProcessingStateFromValue(v string) (*MicrosoftGraphPrintJobProcessingState, error) {
	ev := MicrosoftGraphPrintJobProcessingState(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MicrosoftGraphPrintJobProcessingState: valid values are %v", v, AllowedMicrosoftGraphPrintJobProcessingStateEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MicrosoftGraphPrintJobProcessingState) IsValid() bool {
	for _, existing := range AllowedMicrosoftGraphPrintJobProcessingStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to microsoft.graph.printJobProcessingState value
func (v MicrosoftGraphPrintJobProcessingState) Ptr() *MicrosoftGraphPrintJobProcessingState {
	return &v
}

type NullableMicrosoftGraphPrintJobProcessingState struct {
	value *MicrosoftGraphPrintJobProcessingState
	isSet bool
}

func (v NullableMicrosoftGraphPrintJobProcessingState) Get() *MicrosoftGraphPrintJobProcessingState {
	return v.value
}

func (v *NullableMicrosoftGraphPrintJobProcessingState) Set(val *MicrosoftGraphPrintJobProcessingState) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphPrintJobProcessingState) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphPrintJobProcessingState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphPrintJobProcessingState(val *MicrosoftGraphPrintJobProcessingState) *NullableMicrosoftGraphPrintJobProcessingState {
	return &NullableMicrosoftGraphPrintJobProcessingState{value: val, isSet: true}
}

func (v NullableMicrosoftGraphPrintJobProcessingState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphPrintJobProcessingState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

