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

// MicrosoftGraphTimeOffReasonIconType the model 'MicrosoftGraphTimeOffReasonIconType'
type MicrosoftGraphTimeOffReasonIconType string

// List of microsoft.graph.timeOffReasonIconType
const (
	NONE MicrosoftGraphTimeOffReasonIconType = "none"
	CAR MicrosoftGraphTimeOffReasonIconType = "car"
	CALENDAR MicrosoftGraphTimeOffReasonIconType = "calendar"
	RUNNING MicrosoftGraphTimeOffReasonIconType = "running"
	PLANE MicrosoftGraphTimeOffReasonIconType = "plane"
	FIRST_AID MicrosoftGraphTimeOffReasonIconType = "firstAid"
	DOCTOR MicrosoftGraphTimeOffReasonIconType = "doctor"
	NOT_WORKING MicrosoftGraphTimeOffReasonIconType = "notWorking"
	CLOCK MicrosoftGraphTimeOffReasonIconType = "clock"
	JURY_DUTY MicrosoftGraphTimeOffReasonIconType = "juryDuty"
	GLOBE MicrosoftGraphTimeOffReasonIconType = "globe"
	CUP MicrosoftGraphTimeOffReasonIconType = "cup"
	PHONE MicrosoftGraphTimeOffReasonIconType = "phone"
	WEATHER MicrosoftGraphTimeOffReasonIconType = "weather"
	UMBRELLA MicrosoftGraphTimeOffReasonIconType = "umbrella"
	PIGGY_BANK MicrosoftGraphTimeOffReasonIconType = "piggyBank"
	DOG MicrosoftGraphTimeOffReasonIconType = "dog"
	CAKE MicrosoftGraphTimeOffReasonIconType = "cake"
	TRAFFIC_CONE MicrosoftGraphTimeOffReasonIconType = "trafficCone"
	PIN MicrosoftGraphTimeOffReasonIconType = "pin"
	SUNNY MicrosoftGraphTimeOffReasonIconType = "sunny"
	UNKNOWN_FUTURE_VALUE MicrosoftGraphTimeOffReasonIconType = "unknownFutureValue"
)

// All allowed values of MicrosoftGraphTimeOffReasonIconType enum
var AllowedMicrosoftGraphTimeOffReasonIconTypeEnumValues = []MicrosoftGraphTimeOffReasonIconType{
	"none",
	"car",
	"calendar",
	"running",
	"plane",
	"firstAid",
	"doctor",
	"notWorking",
	"clock",
	"juryDuty",
	"globe",
	"cup",
	"phone",
	"weather",
	"umbrella",
	"piggyBank",
	"dog",
	"cake",
	"trafficCone",
	"pin",
	"sunny",
	"unknownFutureValue",
}

func (v *MicrosoftGraphTimeOffReasonIconType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MicrosoftGraphTimeOffReasonIconType(value)
	for _, existing := range AllowedMicrosoftGraphTimeOffReasonIconTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MicrosoftGraphTimeOffReasonIconType", value)
}

// NewMicrosoftGraphTimeOffReasonIconTypeFromValue returns a pointer to a valid MicrosoftGraphTimeOffReasonIconType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMicrosoftGraphTimeOffReasonIconTypeFromValue(v string) (*MicrosoftGraphTimeOffReasonIconType, error) {
	ev := MicrosoftGraphTimeOffReasonIconType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MicrosoftGraphTimeOffReasonIconType: valid values are %v", v, AllowedMicrosoftGraphTimeOffReasonIconTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MicrosoftGraphTimeOffReasonIconType) IsValid() bool {
	for _, existing := range AllowedMicrosoftGraphTimeOffReasonIconTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to microsoft.graph.timeOffReasonIconType value
func (v MicrosoftGraphTimeOffReasonIconType) Ptr() *MicrosoftGraphTimeOffReasonIconType {
	return &v
}

type NullableMicrosoftGraphTimeOffReasonIconType struct {
	value *MicrosoftGraphTimeOffReasonIconType
	isSet bool
}

func (v NullableMicrosoftGraphTimeOffReasonIconType) Get() *MicrosoftGraphTimeOffReasonIconType {
	return v.value
}

func (v *NullableMicrosoftGraphTimeOffReasonIconType) Set(val *MicrosoftGraphTimeOffReasonIconType) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphTimeOffReasonIconType) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphTimeOffReasonIconType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphTimeOffReasonIconType(val *MicrosoftGraphTimeOffReasonIconType) *NullableMicrosoftGraphTimeOffReasonIconType {
	return &NullableMicrosoftGraphTimeOffReasonIconType{value: val, isSet: true}
}

func (v NullableMicrosoftGraphTimeOffReasonIconType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphTimeOffReasonIconType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

