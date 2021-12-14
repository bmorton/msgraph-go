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

// MicrosoftGraphAutomaticRepliesStatus the model 'MicrosoftGraphAutomaticRepliesStatus'
type MicrosoftGraphAutomaticRepliesStatus string

// List of microsoft.graph.automaticRepliesStatus
const (
	DISABLED MicrosoftGraphAutomaticRepliesStatus = "disabled"
	ALWAYS_ENABLED MicrosoftGraphAutomaticRepliesStatus = "alwaysEnabled"
	SCHEDULED MicrosoftGraphAutomaticRepliesStatus = "scheduled"
)

// All allowed values of MicrosoftGraphAutomaticRepliesStatus enum
var AllowedMicrosoftGraphAutomaticRepliesStatusEnumValues = []MicrosoftGraphAutomaticRepliesStatus{
	"disabled",
	"alwaysEnabled",
	"scheduled",
}

func (v *MicrosoftGraphAutomaticRepliesStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MicrosoftGraphAutomaticRepliesStatus(value)
	for _, existing := range AllowedMicrosoftGraphAutomaticRepliesStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MicrosoftGraphAutomaticRepliesStatus", value)
}

// NewMicrosoftGraphAutomaticRepliesStatusFromValue returns a pointer to a valid MicrosoftGraphAutomaticRepliesStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMicrosoftGraphAutomaticRepliesStatusFromValue(v string) (*MicrosoftGraphAutomaticRepliesStatus, error) {
	ev := MicrosoftGraphAutomaticRepliesStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MicrosoftGraphAutomaticRepliesStatus: valid values are %v", v, AllowedMicrosoftGraphAutomaticRepliesStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MicrosoftGraphAutomaticRepliesStatus) IsValid() bool {
	for _, existing := range AllowedMicrosoftGraphAutomaticRepliesStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to microsoft.graph.automaticRepliesStatus value
func (v MicrosoftGraphAutomaticRepliesStatus) Ptr() *MicrosoftGraphAutomaticRepliesStatus {
	return &v
}

type NullableMicrosoftGraphAutomaticRepliesStatus struct {
	value *MicrosoftGraphAutomaticRepliesStatus
	isSet bool
}

func (v NullableMicrosoftGraphAutomaticRepliesStatus) Get() *MicrosoftGraphAutomaticRepliesStatus {
	return v.value
}

func (v *NullableMicrosoftGraphAutomaticRepliesStatus) Set(val *MicrosoftGraphAutomaticRepliesStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphAutomaticRepliesStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphAutomaticRepliesStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphAutomaticRepliesStatus(val *MicrosoftGraphAutomaticRepliesStatus) *NullableMicrosoftGraphAutomaticRepliesStatus {
	return &NullableMicrosoftGraphAutomaticRepliesStatus{value: val, isSet: true}
}

func (v NullableMicrosoftGraphAutomaticRepliesStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphAutomaticRepliesStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

