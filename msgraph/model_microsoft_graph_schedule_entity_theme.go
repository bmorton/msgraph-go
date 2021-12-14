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

// MicrosoftGraphScheduleEntityTheme the model 'MicrosoftGraphScheduleEntityTheme'
type MicrosoftGraphScheduleEntityTheme string

// List of microsoft.graph.scheduleEntityTheme
const (
	WHITE MicrosoftGraphScheduleEntityTheme = "white"
	BLUE MicrosoftGraphScheduleEntityTheme = "blue"
	GREEN MicrosoftGraphScheduleEntityTheme = "green"
	PURPLE MicrosoftGraphScheduleEntityTheme = "purple"
	PINK MicrosoftGraphScheduleEntityTheme = "pink"
	YELLOW MicrosoftGraphScheduleEntityTheme = "yellow"
	GRAY MicrosoftGraphScheduleEntityTheme = "gray"
	DARK_BLUE MicrosoftGraphScheduleEntityTheme = "darkBlue"
	DARK_GREEN MicrosoftGraphScheduleEntityTheme = "darkGreen"
	DARK_PURPLE MicrosoftGraphScheduleEntityTheme = "darkPurple"
	DARK_PINK MicrosoftGraphScheduleEntityTheme = "darkPink"
	DARK_YELLOW MicrosoftGraphScheduleEntityTheme = "darkYellow"
	UNKNOWN_FUTURE_VALUE MicrosoftGraphScheduleEntityTheme = "unknownFutureValue"
)

// All allowed values of MicrosoftGraphScheduleEntityTheme enum
var AllowedMicrosoftGraphScheduleEntityThemeEnumValues = []MicrosoftGraphScheduleEntityTheme{
	"white",
	"blue",
	"green",
	"purple",
	"pink",
	"yellow",
	"gray",
	"darkBlue",
	"darkGreen",
	"darkPurple",
	"darkPink",
	"darkYellow",
	"unknownFutureValue",
}

func (v *MicrosoftGraphScheduleEntityTheme) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MicrosoftGraphScheduleEntityTheme(value)
	for _, existing := range AllowedMicrosoftGraphScheduleEntityThemeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MicrosoftGraphScheduleEntityTheme", value)
}

// NewMicrosoftGraphScheduleEntityThemeFromValue returns a pointer to a valid MicrosoftGraphScheduleEntityTheme
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMicrosoftGraphScheduleEntityThemeFromValue(v string) (*MicrosoftGraphScheduleEntityTheme, error) {
	ev := MicrosoftGraphScheduleEntityTheme(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MicrosoftGraphScheduleEntityTheme: valid values are %v", v, AllowedMicrosoftGraphScheduleEntityThemeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MicrosoftGraphScheduleEntityTheme) IsValid() bool {
	for _, existing := range AllowedMicrosoftGraphScheduleEntityThemeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to microsoft.graph.scheduleEntityTheme value
func (v MicrosoftGraphScheduleEntityTheme) Ptr() *MicrosoftGraphScheduleEntityTheme {
	return &v
}

type NullableMicrosoftGraphScheduleEntityTheme struct {
	value *MicrosoftGraphScheduleEntityTheme
	isSet bool
}

func (v NullableMicrosoftGraphScheduleEntityTheme) Get() *MicrosoftGraphScheduleEntityTheme {
	return v.value
}

func (v *NullableMicrosoftGraphScheduleEntityTheme) Set(val *MicrosoftGraphScheduleEntityTheme) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphScheduleEntityTheme) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphScheduleEntityTheme) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphScheduleEntityTheme(val *MicrosoftGraphScheduleEntityTheme) *NullableMicrosoftGraphScheduleEntityTheme {
	return &NullableMicrosoftGraphScheduleEntityTheme{value: val, isSet: true}
}

func (v NullableMicrosoftGraphScheduleEntityTheme) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphScheduleEntityTheme) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

