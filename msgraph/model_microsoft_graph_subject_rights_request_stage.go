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

// MicrosoftGraphSubjectRightsRequestStage the model 'MicrosoftGraphSubjectRightsRequestStage'
type MicrosoftGraphSubjectRightsRequestStage string

// List of microsoft.graph.subjectRightsRequestStage
const (
	CONTENT_RETRIEVAL MicrosoftGraphSubjectRightsRequestStage = "contentRetrieval"
	CONTENT_REVIEW MicrosoftGraphSubjectRightsRequestStage = "contentReview"
	GENERATE_REPORT MicrosoftGraphSubjectRightsRequestStage = "generateReport"
	CONTENT_DELETION MicrosoftGraphSubjectRightsRequestStage = "contentDeletion"
	CASE_RESOLVED MicrosoftGraphSubjectRightsRequestStage = "caseResolved"
	CONTENT_ESTIMATE MicrosoftGraphSubjectRightsRequestStage = "contentEstimate"
	UNKNOWN_FUTURE_VALUE MicrosoftGraphSubjectRightsRequestStage = "unknownFutureValue"
)

// All allowed values of MicrosoftGraphSubjectRightsRequestStage enum
var AllowedMicrosoftGraphSubjectRightsRequestStageEnumValues = []MicrosoftGraphSubjectRightsRequestStage{
	"contentRetrieval",
	"contentReview",
	"generateReport",
	"contentDeletion",
	"caseResolved",
	"contentEstimate",
	"unknownFutureValue",
}

func (v *MicrosoftGraphSubjectRightsRequestStage) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MicrosoftGraphSubjectRightsRequestStage(value)
	for _, existing := range AllowedMicrosoftGraphSubjectRightsRequestStageEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MicrosoftGraphSubjectRightsRequestStage", value)
}

// NewMicrosoftGraphSubjectRightsRequestStageFromValue returns a pointer to a valid MicrosoftGraphSubjectRightsRequestStage
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMicrosoftGraphSubjectRightsRequestStageFromValue(v string) (*MicrosoftGraphSubjectRightsRequestStage, error) {
	ev := MicrosoftGraphSubjectRightsRequestStage(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MicrosoftGraphSubjectRightsRequestStage: valid values are %v", v, AllowedMicrosoftGraphSubjectRightsRequestStageEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MicrosoftGraphSubjectRightsRequestStage) IsValid() bool {
	for _, existing := range AllowedMicrosoftGraphSubjectRightsRequestStageEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to microsoft.graph.subjectRightsRequestStage value
func (v MicrosoftGraphSubjectRightsRequestStage) Ptr() *MicrosoftGraphSubjectRightsRequestStage {
	return &v
}

type NullableMicrosoftGraphSubjectRightsRequestStage struct {
	value *MicrosoftGraphSubjectRightsRequestStage
	isSet bool
}

func (v NullableMicrosoftGraphSubjectRightsRequestStage) Get() *MicrosoftGraphSubjectRightsRequestStage {
	return v.value
}

func (v *NullableMicrosoftGraphSubjectRightsRequestStage) Set(val *MicrosoftGraphSubjectRightsRequestStage) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphSubjectRightsRequestStage) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphSubjectRightsRequestStage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphSubjectRightsRequestStage(val *MicrosoftGraphSubjectRightsRequestStage) *NullableMicrosoftGraphSubjectRightsRequestStage {
	return &NullableMicrosoftGraphSubjectRightsRequestStage{value: val, isSet: true}
}

func (v NullableMicrosoftGraphSubjectRightsRequestStage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphSubjectRightsRequestStage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

