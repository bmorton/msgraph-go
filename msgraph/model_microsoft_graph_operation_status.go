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

// MicrosoftGraphOperationStatus the model 'MicrosoftGraphOperationStatus'
type MicrosoftGraphOperationStatus string

// List of microsoft.graph.operationStatus
const (
	NOT_STARTED MicrosoftGraphOperationStatus = "NotStarted"
	RUNNING MicrosoftGraphOperationStatus = "Running"
	COMPLETED MicrosoftGraphOperationStatus = "Completed"
	FAILED MicrosoftGraphOperationStatus = "Failed"
)

// All allowed values of MicrosoftGraphOperationStatus enum
var AllowedMicrosoftGraphOperationStatusEnumValues = []MicrosoftGraphOperationStatus{
	"NotStarted",
	"Running",
	"Completed",
	"Failed",
}

func (v *MicrosoftGraphOperationStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MicrosoftGraphOperationStatus(value)
	for _, existing := range AllowedMicrosoftGraphOperationStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MicrosoftGraphOperationStatus", value)
}

// NewMicrosoftGraphOperationStatusFromValue returns a pointer to a valid MicrosoftGraphOperationStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMicrosoftGraphOperationStatusFromValue(v string) (*MicrosoftGraphOperationStatus, error) {
	ev := MicrosoftGraphOperationStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MicrosoftGraphOperationStatus: valid values are %v", v, AllowedMicrosoftGraphOperationStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MicrosoftGraphOperationStatus) IsValid() bool {
	for _, existing := range AllowedMicrosoftGraphOperationStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to microsoft.graph.operationStatus value
func (v MicrosoftGraphOperationStatus) Ptr() *MicrosoftGraphOperationStatus {
	return &v
}

type NullableMicrosoftGraphOperationStatus struct {
	value *MicrosoftGraphOperationStatus
	isSet bool
}

func (v NullableMicrosoftGraphOperationStatus) Get() *MicrosoftGraphOperationStatus {
	return v.value
}

func (v *NullableMicrosoftGraphOperationStatus) Set(val *MicrosoftGraphOperationStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphOperationStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphOperationStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphOperationStatus(val *MicrosoftGraphOperationStatus) *NullableMicrosoftGraphOperationStatus {
	return &NullableMicrosoftGraphOperationStatus{value: val, isSet: true}
}

func (v NullableMicrosoftGraphOperationStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphOperationStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

