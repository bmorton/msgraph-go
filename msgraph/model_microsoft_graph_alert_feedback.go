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

// MicrosoftGraphAlertFeedback the model 'MicrosoftGraphAlertFeedback'
type MicrosoftGraphAlertFeedback string

// List of microsoft.graph.alertFeedback
const (
	UNKNOWN MicrosoftGraphAlertFeedback = "unknown"
	TRUE_POSITIVE MicrosoftGraphAlertFeedback = "truePositive"
	FALSE_POSITIVE MicrosoftGraphAlertFeedback = "falsePositive"
	BENIGN_POSITIVE MicrosoftGraphAlertFeedback = "benignPositive"
	UNKNOWN_FUTURE_VALUE MicrosoftGraphAlertFeedback = "unknownFutureValue"
)

// All allowed values of MicrosoftGraphAlertFeedback enum
var AllowedMicrosoftGraphAlertFeedbackEnumValues = []MicrosoftGraphAlertFeedback{
	"unknown",
	"truePositive",
	"falsePositive",
	"benignPositive",
	"unknownFutureValue",
}

func (v *MicrosoftGraphAlertFeedback) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MicrosoftGraphAlertFeedback(value)
	for _, existing := range AllowedMicrosoftGraphAlertFeedbackEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MicrosoftGraphAlertFeedback", value)
}

// NewMicrosoftGraphAlertFeedbackFromValue returns a pointer to a valid MicrosoftGraphAlertFeedback
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMicrosoftGraphAlertFeedbackFromValue(v string) (*MicrosoftGraphAlertFeedback, error) {
	ev := MicrosoftGraphAlertFeedback(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MicrosoftGraphAlertFeedback: valid values are %v", v, AllowedMicrosoftGraphAlertFeedbackEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MicrosoftGraphAlertFeedback) IsValid() bool {
	for _, existing := range AllowedMicrosoftGraphAlertFeedbackEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to microsoft.graph.alertFeedback value
func (v MicrosoftGraphAlertFeedback) Ptr() *MicrosoftGraphAlertFeedback {
	return &v
}

type NullableMicrosoftGraphAlertFeedback struct {
	value *MicrosoftGraphAlertFeedback
	isSet bool
}

func (v NullableMicrosoftGraphAlertFeedback) Get() *MicrosoftGraphAlertFeedback {
	return v.value
}

func (v *NullableMicrosoftGraphAlertFeedback) Set(val *MicrosoftGraphAlertFeedback) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphAlertFeedback) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphAlertFeedback) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphAlertFeedback(val *MicrosoftGraphAlertFeedback) *NullableMicrosoftGraphAlertFeedback {
	return &NullableMicrosoftGraphAlertFeedback{value: val, isSet: true}
}

func (v NullableMicrosoftGraphAlertFeedback) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphAlertFeedback) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

