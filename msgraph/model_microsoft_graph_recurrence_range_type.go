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

// MicrosoftGraphRecurrenceRangeType the model 'MicrosoftGraphRecurrenceRangeType'
type MicrosoftGraphRecurrenceRangeType string

// List of microsoft.graph.recurrenceRangeType
const (
	END_DATE MicrosoftGraphRecurrenceRangeType = "endDate"
	NO_END MicrosoftGraphRecurrenceRangeType = "noEnd"
	NUMBERED MicrosoftGraphRecurrenceRangeType = "numbered"
)

// All allowed values of MicrosoftGraphRecurrenceRangeType enum
var AllowedMicrosoftGraphRecurrenceRangeTypeEnumValues = []MicrosoftGraphRecurrenceRangeType{
	"endDate",
	"noEnd",
	"numbered",
}

func (v *MicrosoftGraphRecurrenceRangeType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MicrosoftGraphRecurrenceRangeType(value)
	for _, existing := range AllowedMicrosoftGraphRecurrenceRangeTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MicrosoftGraphRecurrenceRangeType", value)
}

// NewMicrosoftGraphRecurrenceRangeTypeFromValue returns a pointer to a valid MicrosoftGraphRecurrenceRangeType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMicrosoftGraphRecurrenceRangeTypeFromValue(v string) (*MicrosoftGraphRecurrenceRangeType, error) {
	ev := MicrosoftGraphRecurrenceRangeType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MicrosoftGraphRecurrenceRangeType: valid values are %v", v, AllowedMicrosoftGraphRecurrenceRangeTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MicrosoftGraphRecurrenceRangeType) IsValid() bool {
	for _, existing := range AllowedMicrosoftGraphRecurrenceRangeTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to microsoft.graph.recurrenceRangeType value
func (v MicrosoftGraphRecurrenceRangeType) Ptr() *MicrosoftGraphRecurrenceRangeType {
	return &v
}

type NullableMicrosoftGraphRecurrenceRangeType struct {
	value *MicrosoftGraphRecurrenceRangeType
	isSet bool
}

func (v NullableMicrosoftGraphRecurrenceRangeType) Get() *MicrosoftGraphRecurrenceRangeType {
	return v.value
}

func (v *NullableMicrosoftGraphRecurrenceRangeType) Set(val *MicrosoftGraphRecurrenceRangeType) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphRecurrenceRangeType) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphRecurrenceRangeType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphRecurrenceRangeType(val *MicrosoftGraphRecurrenceRangeType) *NullableMicrosoftGraphRecurrenceRangeType {
	return &NullableMicrosoftGraphRecurrenceRangeType{value: val, isSet: true}
}

func (v NullableMicrosoftGraphRecurrenceRangeType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphRecurrenceRangeType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
