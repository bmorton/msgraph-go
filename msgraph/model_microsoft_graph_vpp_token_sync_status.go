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

// MicrosoftGraphVppTokenSyncStatus Possible sync statuses associated with an Apple Volume Purchase Program token.
type MicrosoftGraphVppTokenSyncStatus string

// List of microsoft.graph.vppTokenSyncStatus
const (
	NONE MicrosoftGraphVppTokenSyncStatus = "none"
	IN_PROGRESS MicrosoftGraphVppTokenSyncStatus = "inProgress"
	COMPLETED MicrosoftGraphVppTokenSyncStatus = "completed"
	FAILED MicrosoftGraphVppTokenSyncStatus = "failed"
)

// All allowed values of MicrosoftGraphVppTokenSyncStatus enum
var AllowedMicrosoftGraphVppTokenSyncStatusEnumValues = []MicrosoftGraphVppTokenSyncStatus{
	"none",
	"inProgress",
	"completed",
	"failed",
}

func (v *MicrosoftGraphVppTokenSyncStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MicrosoftGraphVppTokenSyncStatus(value)
	for _, existing := range AllowedMicrosoftGraphVppTokenSyncStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MicrosoftGraphVppTokenSyncStatus", value)
}

// NewMicrosoftGraphVppTokenSyncStatusFromValue returns a pointer to a valid MicrosoftGraphVppTokenSyncStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMicrosoftGraphVppTokenSyncStatusFromValue(v string) (*MicrosoftGraphVppTokenSyncStatus, error) {
	ev := MicrosoftGraphVppTokenSyncStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MicrosoftGraphVppTokenSyncStatus: valid values are %v", v, AllowedMicrosoftGraphVppTokenSyncStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MicrosoftGraphVppTokenSyncStatus) IsValid() bool {
	for _, existing := range AllowedMicrosoftGraphVppTokenSyncStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to microsoft.graph.vppTokenSyncStatus value
func (v MicrosoftGraphVppTokenSyncStatus) Ptr() *MicrosoftGraphVppTokenSyncStatus {
	return &v
}

type NullableMicrosoftGraphVppTokenSyncStatus struct {
	value *MicrosoftGraphVppTokenSyncStatus
	isSet bool
}

func (v NullableMicrosoftGraphVppTokenSyncStatus) Get() *MicrosoftGraphVppTokenSyncStatus {
	return v.value
}

func (v *NullableMicrosoftGraphVppTokenSyncStatus) Set(val *MicrosoftGraphVppTokenSyncStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphVppTokenSyncStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphVppTokenSyncStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphVppTokenSyncStatus(val *MicrosoftGraphVppTokenSyncStatus) *NullableMicrosoftGraphVppTokenSyncStatus {
	return &NullableMicrosoftGraphVppTokenSyncStatus{value: val, isSet: true}
}

func (v NullableMicrosoftGraphVppTokenSyncStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphVppTokenSyncStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

