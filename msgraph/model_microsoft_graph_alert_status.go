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

// MicrosoftGraphAlertStatus the model 'MicrosoftGraphAlertStatus'
type MicrosoftGraphAlertStatus string

// List of microsoft.graph.alertStatus
const (
	UNKNOWN MicrosoftGraphAlertStatus = "unknown"
	NEW_ALERT MicrosoftGraphAlertStatus = "newAlert"
	IN_PROGRESS MicrosoftGraphAlertStatus = "inProgress"
	RESOLVED MicrosoftGraphAlertStatus = "resolved"
	DISMISSED MicrosoftGraphAlertStatus = "dismissed"
	UNKNOWN_FUTURE_VALUE MicrosoftGraphAlertStatus = "unknownFutureValue"
)

// All allowed values of MicrosoftGraphAlertStatus enum
var AllowedMicrosoftGraphAlertStatusEnumValues = []MicrosoftGraphAlertStatus{
	"unknown",
	"newAlert",
	"inProgress",
	"resolved",
	"dismissed",
	"unknownFutureValue",
}

func (v *MicrosoftGraphAlertStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MicrosoftGraphAlertStatus(value)
	for _, existing := range AllowedMicrosoftGraphAlertStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MicrosoftGraphAlertStatus", value)
}

// NewMicrosoftGraphAlertStatusFromValue returns a pointer to a valid MicrosoftGraphAlertStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMicrosoftGraphAlertStatusFromValue(v string) (*MicrosoftGraphAlertStatus, error) {
	ev := MicrosoftGraphAlertStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MicrosoftGraphAlertStatus: valid values are %v", v, AllowedMicrosoftGraphAlertStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MicrosoftGraphAlertStatus) IsValid() bool {
	for _, existing := range AllowedMicrosoftGraphAlertStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to microsoft.graph.alertStatus value
func (v MicrosoftGraphAlertStatus) Ptr() *MicrosoftGraphAlertStatus {
	return &v
}

type NullableMicrosoftGraphAlertStatus struct {
	value *MicrosoftGraphAlertStatus
	isSet bool
}

func (v NullableMicrosoftGraphAlertStatus) Get() *MicrosoftGraphAlertStatus {
	return v.value
}

func (v *NullableMicrosoftGraphAlertStatus) Set(val *MicrosoftGraphAlertStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphAlertStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphAlertStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphAlertStatus(val *MicrosoftGraphAlertStatus) *NullableMicrosoftGraphAlertStatus {
	return &NullableMicrosoftGraphAlertStatus{value: val, isSet: true}
}

func (v NullableMicrosoftGraphAlertStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphAlertStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

