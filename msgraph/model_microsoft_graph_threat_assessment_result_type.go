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

// MicrosoftGraphThreatAssessmentResultType the model 'MicrosoftGraphThreatAssessmentResultType'
type MicrosoftGraphThreatAssessmentResultType string

// List of microsoft.graph.threatAssessmentResultType
const (
	CHECK_POLICY MicrosoftGraphThreatAssessmentResultType = "checkPolicy"
	RESCAN MicrosoftGraphThreatAssessmentResultType = "rescan"
	UNKNOWN_FUTURE_VALUE MicrosoftGraphThreatAssessmentResultType = "unknownFutureValue"
)

// All allowed values of MicrosoftGraphThreatAssessmentResultType enum
var AllowedMicrosoftGraphThreatAssessmentResultTypeEnumValues = []MicrosoftGraphThreatAssessmentResultType{
	"checkPolicy",
	"rescan",
	"unknownFutureValue",
}

func (v *MicrosoftGraphThreatAssessmentResultType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MicrosoftGraphThreatAssessmentResultType(value)
	for _, existing := range AllowedMicrosoftGraphThreatAssessmentResultTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MicrosoftGraphThreatAssessmentResultType", value)
}

// NewMicrosoftGraphThreatAssessmentResultTypeFromValue returns a pointer to a valid MicrosoftGraphThreatAssessmentResultType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMicrosoftGraphThreatAssessmentResultTypeFromValue(v string) (*MicrosoftGraphThreatAssessmentResultType, error) {
	ev := MicrosoftGraphThreatAssessmentResultType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MicrosoftGraphThreatAssessmentResultType: valid values are %v", v, AllowedMicrosoftGraphThreatAssessmentResultTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MicrosoftGraphThreatAssessmentResultType) IsValid() bool {
	for _, existing := range AllowedMicrosoftGraphThreatAssessmentResultTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to microsoft.graph.threatAssessmentResultType value
func (v MicrosoftGraphThreatAssessmentResultType) Ptr() *MicrosoftGraphThreatAssessmentResultType {
	return &v
}

type NullableMicrosoftGraphThreatAssessmentResultType struct {
	value *MicrosoftGraphThreatAssessmentResultType
	isSet bool
}

func (v NullableMicrosoftGraphThreatAssessmentResultType) Get() *MicrosoftGraphThreatAssessmentResultType {
	return v.value
}

func (v *NullableMicrosoftGraphThreatAssessmentResultType) Set(val *MicrosoftGraphThreatAssessmentResultType) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphThreatAssessmentResultType) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphThreatAssessmentResultType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphThreatAssessmentResultType(val *MicrosoftGraphThreatAssessmentResultType) *NullableMicrosoftGraphThreatAssessmentResultType {
	return &NullableMicrosoftGraphThreatAssessmentResultType{value: val, isSet: true}
}

func (v NullableMicrosoftGraphThreatAssessmentResultType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphThreatAssessmentResultType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
