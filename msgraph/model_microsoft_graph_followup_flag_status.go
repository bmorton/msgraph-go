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

// MicrosoftGraphFollowupFlagStatus the model 'MicrosoftGraphFollowupFlagStatus'
type MicrosoftGraphFollowupFlagStatus string

// List of microsoft.graph.followupFlagStatus
const (
	NOT_FLAGGED MicrosoftGraphFollowupFlagStatus = "notFlagged"
	COMPLETE MicrosoftGraphFollowupFlagStatus = "complete"
	FLAGGED MicrosoftGraphFollowupFlagStatus = "flagged"
)

// All allowed values of MicrosoftGraphFollowupFlagStatus enum
var AllowedMicrosoftGraphFollowupFlagStatusEnumValues = []MicrosoftGraphFollowupFlagStatus{
	"notFlagged",
	"complete",
	"flagged",
}

func (v *MicrosoftGraphFollowupFlagStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MicrosoftGraphFollowupFlagStatus(value)
	for _, existing := range AllowedMicrosoftGraphFollowupFlagStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MicrosoftGraphFollowupFlagStatus", value)
}

// NewMicrosoftGraphFollowupFlagStatusFromValue returns a pointer to a valid MicrosoftGraphFollowupFlagStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMicrosoftGraphFollowupFlagStatusFromValue(v string) (*MicrosoftGraphFollowupFlagStatus, error) {
	ev := MicrosoftGraphFollowupFlagStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MicrosoftGraphFollowupFlagStatus: valid values are %v", v, AllowedMicrosoftGraphFollowupFlagStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MicrosoftGraphFollowupFlagStatus) IsValid() bool {
	for _, existing := range AllowedMicrosoftGraphFollowupFlagStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to microsoft.graph.followupFlagStatus value
func (v MicrosoftGraphFollowupFlagStatus) Ptr() *MicrosoftGraphFollowupFlagStatus {
	return &v
}

type NullableMicrosoftGraphFollowupFlagStatus struct {
	value *MicrosoftGraphFollowupFlagStatus
	isSet bool
}

func (v NullableMicrosoftGraphFollowupFlagStatus) Get() *MicrosoftGraphFollowupFlagStatus {
	return v.value
}

func (v *NullableMicrosoftGraphFollowupFlagStatus) Set(val *MicrosoftGraphFollowupFlagStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphFollowupFlagStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphFollowupFlagStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphFollowupFlagStatus(val *MicrosoftGraphFollowupFlagStatus) *NullableMicrosoftGraphFollowupFlagStatus {
	return &NullableMicrosoftGraphFollowupFlagStatus{value: val, isSet: true}
}

func (v NullableMicrosoftGraphFollowupFlagStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphFollowupFlagStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

