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

// MicrosoftGraphAccessReviewScheduleDefinitionFilterByCurrentUserOptions the model 'MicrosoftGraphAccessReviewScheduleDefinitionFilterByCurrentUserOptions'
type MicrosoftGraphAccessReviewScheduleDefinitionFilterByCurrentUserOptions string

// List of microsoft.graph.accessReviewScheduleDefinitionFilterByCurrentUserOptions
const (
	REVIEWER MicrosoftGraphAccessReviewScheduleDefinitionFilterByCurrentUserOptions = "reviewer"
	UNKNOWN_FUTURE_VALUE MicrosoftGraphAccessReviewScheduleDefinitionFilterByCurrentUserOptions = "unknownFutureValue"
)

// All allowed values of MicrosoftGraphAccessReviewScheduleDefinitionFilterByCurrentUserOptions enum
var AllowedMicrosoftGraphAccessReviewScheduleDefinitionFilterByCurrentUserOptionsEnumValues = []MicrosoftGraphAccessReviewScheduleDefinitionFilterByCurrentUserOptions{
	"reviewer",
	"unknownFutureValue",
}

func (v *MicrosoftGraphAccessReviewScheduleDefinitionFilterByCurrentUserOptions) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MicrosoftGraphAccessReviewScheduleDefinitionFilterByCurrentUserOptions(value)
	for _, existing := range AllowedMicrosoftGraphAccessReviewScheduleDefinitionFilterByCurrentUserOptionsEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MicrosoftGraphAccessReviewScheduleDefinitionFilterByCurrentUserOptions", value)
}

// NewMicrosoftGraphAccessReviewScheduleDefinitionFilterByCurrentUserOptionsFromValue returns a pointer to a valid MicrosoftGraphAccessReviewScheduleDefinitionFilterByCurrentUserOptions
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMicrosoftGraphAccessReviewScheduleDefinitionFilterByCurrentUserOptionsFromValue(v string) (*MicrosoftGraphAccessReviewScheduleDefinitionFilterByCurrentUserOptions, error) {
	ev := MicrosoftGraphAccessReviewScheduleDefinitionFilterByCurrentUserOptions(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MicrosoftGraphAccessReviewScheduleDefinitionFilterByCurrentUserOptions: valid values are %v", v, AllowedMicrosoftGraphAccessReviewScheduleDefinitionFilterByCurrentUserOptionsEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MicrosoftGraphAccessReviewScheduleDefinitionFilterByCurrentUserOptions) IsValid() bool {
	for _, existing := range AllowedMicrosoftGraphAccessReviewScheduleDefinitionFilterByCurrentUserOptionsEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to microsoft.graph.accessReviewScheduleDefinitionFilterByCurrentUserOptions value
func (v MicrosoftGraphAccessReviewScheduleDefinitionFilterByCurrentUserOptions) Ptr() *MicrosoftGraphAccessReviewScheduleDefinitionFilterByCurrentUserOptions {
	return &v
}

type NullableMicrosoftGraphAccessReviewScheduleDefinitionFilterByCurrentUserOptions struct {
	value *MicrosoftGraphAccessReviewScheduleDefinitionFilterByCurrentUserOptions
	isSet bool
}

func (v NullableMicrosoftGraphAccessReviewScheduleDefinitionFilterByCurrentUserOptions) Get() *MicrosoftGraphAccessReviewScheduleDefinitionFilterByCurrentUserOptions {
	return v.value
}

func (v *NullableMicrosoftGraphAccessReviewScheduleDefinitionFilterByCurrentUserOptions) Set(val *MicrosoftGraphAccessReviewScheduleDefinitionFilterByCurrentUserOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphAccessReviewScheduleDefinitionFilterByCurrentUserOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphAccessReviewScheduleDefinitionFilterByCurrentUserOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphAccessReviewScheduleDefinitionFilterByCurrentUserOptions(val *MicrosoftGraphAccessReviewScheduleDefinitionFilterByCurrentUserOptions) *NullableMicrosoftGraphAccessReviewScheduleDefinitionFilterByCurrentUserOptions {
	return &NullableMicrosoftGraphAccessReviewScheduleDefinitionFilterByCurrentUserOptions{value: val, isSet: true}
}

func (v NullableMicrosoftGraphAccessReviewScheduleDefinitionFilterByCurrentUserOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphAccessReviewScheduleDefinitionFilterByCurrentUserOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
