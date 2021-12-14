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

// MicrosoftGraphConditionalAccessStatus the model 'MicrosoftGraphConditionalAccessStatus'
type MicrosoftGraphConditionalAccessStatus string

// List of microsoft.graph.conditionalAccessStatus
const (
	SUCCESS MicrosoftGraphConditionalAccessStatus = "success"
	FAILURE MicrosoftGraphConditionalAccessStatus = "failure"
	NOT_APPLIED MicrosoftGraphConditionalAccessStatus = "notApplied"
	UNKNOWN_FUTURE_VALUE MicrosoftGraphConditionalAccessStatus = "unknownFutureValue"
)

// All allowed values of MicrosoftGraphConditionalAccessStatus enum
var AllowedMicrosoftGraphConditionalAccessStatusEnumValues = []MicrosoftGraphConditionalAccessStatus{
	"success",
	"failure",
	"notApplied",
	"unknownFutureValue",
}

func (v *MicrosoftGraphConditionalAccessStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MicrosoftGraphConditionalAccessStatus(value)
	for _, existing := range AllowedMicrosoftGraphConditionalAccessStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MicrosoftGraphConditionalAccessStatus", value)
}

// NewMicrosoftGraphConditionalAccessStatusFromValue returns a pointer to a valid MicrosoftGraphConditionalAccessStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMicrosoftGraphConditionalAccessStatusFromValue(v string) (*MicrosoftGraphConditionalAccessStatus, error) {
	ev := MicrosoftGraphConditionalAccessStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MicrosoftGraphConditionalAccessStatus: valid values are %v", v, AllowedMicrosoftGraphConditionalAccessStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MicrosoftGraphConditionalAccessStatus) IsValid() bool {
	for _, existing := range AllowedMicrosoftGraphConditionalAccessStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to microsoft.graph.conditionalAccessStatus value
func (v MicrosoftGraphConditionalAccessStatus) Ptr() *MicrosoftGraphConditionalAccessStatus {
	return &v
}

type NullableMicrosoftGraphConditionalAccessStatus struct {
	value *MicrosoftGraphConditionalAccessStatus
	isSet bool
}

func (v NullableMicrosoftGraphConditionalAccessStatus) Get() *MicrosoftGraphConditionalAccessStatus {
	return v.value
}

func (v *NullableMicrosoftGraphConditionalAccessStatus) Set(val *MicrosoftGraphConditionalAccessStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphConditionalAccessStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphConditionalAccessStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphConditionalAccessStatus(val *MicrosoftGraphConditionalAccessStatus) *NullableMicrosoftGraphConditionalAccessStatus {
	return &NullableMicrosoftGraphConditionalAccessStatus{value: val, isSet: true}
}

func (v NullableMicrosoftGraphConditionalAccessStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphConditionalAccessStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

