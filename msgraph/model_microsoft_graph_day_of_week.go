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

// MicrosoftGraphDayOfWeek the model 'MicrosoftGraphDayOfWeek'
type MicrosoftGraphDayOfWeek string

// List of microsoft.graph.dayOfWeek
const (
	SUNDAY MicrosoftGraphDayOfWeek = "sunday"
	MONDAY MicrosoftGraphDayOfWeek = "monday"
	TUESDAY MicrosoftGraphDayOfWeek = "tuesday"
	WEDNESDAY MicrosoftGraphDayOfWeek = "wednesday"
	THURSDAY MicrosoftGraphDayOfWeek = "thursday"
	FRIDAY MicrosoftGraphDayOfWeek = "friday"
	SATURDAY MicrosoftGraphDayOfWeek = "saturday"
)

// All allowed values of MicrosoftGraphDayOfWeek enum
var AllowedMicrosoftGraphDayOfWeekEnumValues = []MicrosoftGraphDayOfWeek{
	"sunday",
	"monday",
	"tuesday",
	"wednesday",
	"thursday",
	"friday",
	"saturday",
}

func (v *MicrosoftGraphDayOfWeek) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MicrosoftGraphDayOfWeek(value)
	for _, existing := range AllowedMicrosoftGraphDayOfWeekEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MicrosoftGraphDayOfWeek", value)
}

// NewMicrosoftGraphDayOfWeekFromValue returns a pointer to a valid MicrosoftGraphDayOfWeek
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMicrosoftGraphDayOfWeekFromValue(v string) (*MicrosoftGraphDayOfWeek, error) {
	ev := MicrosoftGraphDayOfWeek(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MicrosoftGraphDayOfWeek: valid values are %v", v, AllowedMicrosoftGraphDayOfWeekEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MicrosoftGraphDayOfWeek) IsValid() bool {
	for _, existing := range AllowedMicrosoftGraphDayOfWeekEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to microsoft.graph.dayOfWeek value
func (v MicrosoftGraphDayOfWeek) Ptr() *MicrosoftGraphDayOfWeek {
	return &v
}

type NullableMicrosoftGraphDayOfWeek struct {
	value *MicrosoftGraphDayOfWeek
	isSet bool
}

func (v NullableMicrosoftGraphDayOfWeek) Get() *MicrosoftGraphDayOfWeek {
	return v.value
}

func (v *NullableMicrosoftGraphDayOfWeek) Set(val *MicrosoftGraphDayOfWeek) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphDayOfWeek) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphDayOfWeek) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphDayOfWeek(val *MicrosoftGraphDayOfWeek) *NullableMicrosoftGraphDayOfWeek {
	return &NullableMicrosoftGraphDayOfWeek{value: val, isSet: true}
}

func (v NullableMicrosoftGraphDayOfWeek) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphDayOfWeek) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

