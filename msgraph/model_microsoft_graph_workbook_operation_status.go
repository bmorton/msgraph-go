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

// MicrosoftGraphWorkbookOperationStatus the model 'MicrosoftGraphWorkbookOperationStatus'
type MicrosoftGraphWorkbookOperationStatus string

// List of microsoft.graph.workbookOperationStatus
const (
	NOT_STARTED MicrosoftGraphWorkbookOperationStatus = "notStarted"
	RUNNING MicrosoftGraphWorkbookOperationStatus = "running"
	SUCCEEDED MicrosoftGraphWorkbookOperationStatus = "succeeded"
	FAILED MicrosoftGraphWorkbookOperationStatus = "failed"
)

// All allowed values of MicrosoftGraphWorkbookOperationStatus enum
var AllowedMicrosoftGraphWorkbookOperationStatusEnumValues = []MicrosoftGraphWorkbookOperationStatus{
	"notStarted",
	"running",
	"succeeded",
	"failed",
}

func (v *MicrosoftGraphWorkbookOperationStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MicrosoftGraphWorkbookOperationStatus(value)
	for _, existing := range AllowedMicrosoftGraphWorkbookOperationStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MicrosoftGraphWorkbookOperationStatus", value)
}

// NewMicrosoftGraphWorkbookOperationStatusFromValue returns a pointer to a valid MicrosoftGraphWorkbookOperationStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMicrosoftGraphWorkbookOperationStatusFromValue(v string) (*MicrosoftGraphWorkbookOperationStatus, error) {
	ev := MicrosoftGraphWorkbookOperationStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MicrosoftGraphWorkbookOperationStatus: valid values are %v", v, AllowedMicrosoftGraphWorkbookOperationStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MicrosoftGraphWorkbookOperationStatus) IsValid() bool {
	for _, existing := range AllowedMicrosoftGraphWorkbookOperationStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to microsoft.graph.workbookOperationStatus value
func (v MicrosoftGraphWorkbookOperationStatus) Ptr() *MicrosoftGraphWorkbookOperationStatus {
	return &v
}

type NullableMicrosoftGraphWorkbookOperationStatus struct {
	value *MicrosoftGraphWorkbookOperationStatus
	isSet bool
}

func (v NullableMicrosoftGraphWorkbookOperationStatus) Get() *MicrosoftGraphWorkbookOperationStatus {
	return v.value
}

func (v *NullableMicrosoftGraphWorkbookOperationStatus) Set(val *MicrosoftGraphWorkbookOperationStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphWorkbookOperationStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphWorkbookOperationStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphWorkbookOperationStatus(val *MicrosoftGraphWorkbookOperationStatus) *NullableMicrosoftGraphWorkbookOperationStatus {
	return &NullableMicrosoftGraphWorkbookOperationStatus{value: val, isSet: true}
}

func (v NullableMicrosoftGraphWorkbookOperationStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphWorkbookOperationStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

