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

// MicrosoftGraphDataPolicyOperationStatus the model 'MicrosoftGraphDataPolicyOperationStatus'
type MicrosoftGraphDataPolicyOperationStatus string

// List of microsoft.graph.dataPolicyOperationStatus
const (
	NOT_STARTED MicrosoftGraphDataPolicyOperationStatus = "notStarted"
	RUNNING MicrosoftGraphDataPolicyOperationStatus = "running"
	COMPLETE MicrosoftGraphDataPolicyOperationStatus = "complete"
	FAILED MicrosoftGraphDataPolicyOperationStatus = "failed"
	UNKNOWN_FUTURE_VALUE MicrosoftGraphDataPolicyOperationStatus = "unknownFutureValue"
)

// All allowed values of MicrosoftGraphDataPolicyOperationStatus enum
var AllowedMicrosoftGraphDataPolicyOperationStatusEnumValues = []MicrosoftGraphDataPolicyOperationStatus{
	"notStarted",
	"running",
	"complete",
	"failed",
	"unknownFutureValue",
}

func (v *MicrosoftGraphDataPolicyOperationStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MicrosoftGraphDataPolicyOperationStatus(value)
	for _, existing := range AllowedMicrosoftGraphDataPolicyOperationStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MicrosoftGraphDataPolicyOperationStatus", value)
}

// NewMicrosoftGraphDataPolicyOperationStatusFromValue returns a pointer to a valid MicrosoftGraphDataPolicyOperationStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMicrosoftGraphDataPolicyOperationStatusFromValue(v string) (*MicrosoftGraphDataPolicyOperationStatus, error) {
	ev := MicrosoftGraphDataPolicyOperationStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MicrosoftGraphDataPolicyOperationStatus: valid values are %v", v, AllowedMicrosoftGraphDataPolicyOperationStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MicrosoftGraphDataPolicyOperationStatus) IsValid() bool {
	for _, existing := range AllowedMicrosoftGraphDataPolicyOperationStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to microsoft.graph.dataPolicyOperationStatus value
func (v MicrosoftGraphDataPolicyOperationStatus) Ptr() *MicrosoftGraphDataPolicyOperationStatus {
	return &v
}

type NullableMicrosoftGraphDataPolicyOperationStatus struct {
	value *MicrosoftGraphDataPolicyOperationStatus
	isSet bool
}

func (v NullableMicrosoftGraphDataPolicyOperationStatus) Get() *MicrosoftGraphDataPolicyOperationStatus {
	return v.value
}

func (v *NullableMicrosoftGraphDataPolicyOperationStatus) Set(val *MicrosoftGraphDataPolicyOperationStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphDataPolicyOperationStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphDataPolicyOperationStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphDataPolicyOperationStatus(val *MicrosoftGraphDataPolicyOperationStatus) *NullableMicrosoftGraphDataPolicyOperationStatus {
	return &NullableMicrosoftGraphDataPolicyOperationStatus{value: val, isSet: true}
}

func (v NullableMicrosoftGraphDataPolicyOperationStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphDataPolicyOperationStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
