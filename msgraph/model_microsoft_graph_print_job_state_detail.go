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

// MicrosoftGraphPrintJobStateDetail the model 'MicrosoftGraphPrintJobStateDetail'
type MicrosoftGraphPrintJobStateDetail string

// List of microsoft.graph.printJobStateDetail
const (
	UPLOAD_PENDING MicrosoftGraphPrintJobStateDetail = "uploadPending"
	TRANSFORMING MicrosoftGraphPrintJobStateDetail = "transforming"
	COMPLETED_SUCCESSFULLY MicrosoftGraphPrintJobStateDetail = "completedSuccessfully"
	COMPLETED_WITH_WARNINGS MicrosoftGraphPrintJobStateDetail = "completedWithWarnings"
	COMPLETED_WITH_ERRORS MicrosoftGraphPrintJobStateDetail = "completedWithErrors"
	RELEASE_WAIT MicrosoftGraphPrintJobStateDetail = "releaseWait"
	INTERPRETING MicrosoftGraphPrintJobStateDetail = "interpreting"
	UNKNOWN_FUTURE_VALUE MicrosoftGraphPrintJobStateDetail = "unknownFutureValue"
)

// All allowed values of MicrosoftGraphPrintJobStateDetail enum
var AllowedMicrosoftGraphPrintJobStateDetailEnumValues = []MicrosoftGraphPrintJobStateDetail{
	"uploadPending",
	"transforming",
	"completedSuccessfully",
	"completedWithWarnings",
	"completedWithErrors",
	"releaseWait",
	"interpreting",
	"unknownFutureValue",
}

func (v *MicrosoftGraphPrintJobStateDetail) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MicrosoftGraphPrintJobStateDetail(value)
	for _, existing := range AllowedMicrosoftGraphPrintJobStateDetailEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MicrosoftGraphPrintJobStateDetail", value)
}

// NewMicrosoftGraphPrintJobStateDetailFromValue returns a pointer to a valid MicrosoftGraphPrintJobStateDetail
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMicrosoftGraphPrintJobStateDetailFromValue(v string) (*MicrosoftGraphPrintJobStateDetail, error) {
	ev := MicrosoftGraphPrintJobStateDetail(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MicrosoftGraphPrintJobStateDetail: valid values are %v", v, AllowedMicrosoftGraphPrintJobStateDetailEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MicrosoftGraphPrintJobStateDetail) IsValid() bool {
	for _, existing := range AllowedMicrosoftGraphPrintJobStateDetailEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to microsoft.graph.printJobStateDetail value
func (v MicrosoftGraphPrintJobStateDetail) Ptr() *MicrosoftGraphPrintJobStateDetail {
	return &v
}

type NullableMicrosoftGraphPrintJobStateDetail struct {
	value *MicrosoftGraphPrintJobStateDetail
	isSet bool
}

func (v NullableMicrosoftGraphPrintJobStateDetail) Get() *MicrosoftGraphPrintJobStateDetail {
	return v.value
}

func (v *NullableMicrosoftGraphPrintJobStateDetail) Set(val *MicrosoftGraphPrintJobStateDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphPrintJobStateDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphPrintJobStateDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphPrintJobStateDetail(val *MicrosoftGraphPrintJobStateDetail) *NullableMicrosoftGraphPrintJobStateDetail {
	return &NullableMicrosoftGraphPrintJobStateDetail{value: val, isSet: true}
}

func (v NullableMicrosoftGraphPrintJobStateDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphPrintJobStateDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

