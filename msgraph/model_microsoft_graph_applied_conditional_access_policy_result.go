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

// MicrosoftGraphAppliedConditionalAccessPolicyResult the model 'MicrosoftGraphAppliedConditionalAccessPolicyResult'
type MicrosoftGraphAppliedConditionalAccessPolicyResult string

// List of microsoft.graph.appliedConditionalAccessPolicyResult
const (
	SUCCESS MicrosoftGraphAppliedConditionalAccessPolicyResult = "success"
	FAILURE MicrosoftGraphAppliedConditionalAccessPolicyResult = "failure"
	NOT_APPLIED MicrosoftGraphAppliedConditionalAccessPolicyResult = "notApplied"
	NOT_ENABLED MicrosoftGraphAppliedConditionalAccessPolicyResult = "notEnabled"
	UNKNOWN MicrosoftGraphAppliedConditionalAccessPolicyResult = "unknown"
	UNKNOWN_FUTURE_VALUE MicrosoftGraphAppliedConditionalAccessPolicyResult = "unknownFutureValue"
)

// All allowed values of MicrosoftGraphAppliedConditionalAccessPolicyResult enum
var AllowedMicrosoftGraphAppliedConditionalAccessPolicyResultEnumValues = []MicrosoftGraphAppliedConditionalAccessPolicyResult{
	"success",
	"failure",
	"notApplied",
	"notEnabled",
	"unknown",
	"unknownFutureValue",
}

func (v *MicrosoftGraphAppliedConditionalAccessPolicyResult) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MicrosoftGraphAppliedConditionalAccessPolicyResult(value)
	for _, existing := range AllowedMicrosoftGraphAppliedConditionalAccessPolicyResultEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MicrosoftGraphAppliedConditionalAccessPolicyResult", value)
}

// NewMicrosoftGraphAppliedConditionalAccessPolicyResultFromValue returns a pointer to a valid MicrosoftGraphAppliedConditionalAccessPolicyResult
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMicrosoftGraphAppliedConditionalAccessPolicyResultFromValue(v string) (*MicrosoftGraphAppliedConditionalAccessPolicyResult, error) {
	ev := MicrosoftGraphAppliedConditionalAccessPolicyResult(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MicrosoftGraphAppliedConditionalAccessPolicyResult: valid values are %v", v, AllowedMicrosoftGraphAppliedConditionalAccessPolicyResultEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MicrosoftGraphAppliedConditionalAccessPolicyResult) IsValid() bool {
	for _, existing := range AllowedMicrosoftGraphAppliedConditionalAccessPolicyResultEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to microsoft.graph.appliedConditionalAccessPolicyResult value
func (v MicrosoftGraphAppliedConditionalAccessPolicyResult) Ptr() *MicrosoftGraphAppliedConditionalAccessPolicyResult {
	return &v
}

type NullableMicrosoftGraphAppliedConditionalAccessPolicyResult struct {
	value *MicrosoftGraphAppliedConditionalAccessPolicyResult
	isSet bool
}

func (v NullableMicrosoftGraphAppliedConditionalAccessPolicyResult) Get() *MicrosoftGraphAppliedConditionalAccessPolicyResult {
	return v.value
}

func (v *NullableMicrosoftGraphAppliedConditionalAccessPolicyResult) Set(val *MicrosoftGraphAppliedConditionalAccessPolicyResult) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphAppliedConditionalAccessPolicyResult) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphAppliedConditionalAccessPolicyResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphAppliedConditionalAccessPolicyResult(val *MicrosoftGraphAppliedConditionalAccessPolicyResult) *NullableMicrosoftGraphAppliedConditionalAccessPolicyResult {
	return &NullableMicrosoftGraphAppliedConditionalAccessPolicyResult{value: val, isSet: true}
}

func (v NullableMicrosoftGraphAppliedConditionalAccessPolicyResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphAppliedConditionalAccessPolicyResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

