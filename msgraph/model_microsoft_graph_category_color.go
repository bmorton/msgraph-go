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

// MicrosoftGraphCategoryColor the model 'MicrosoftGraphCategoryColor'
type MicrosoftGraphCategoryColor string

// List of microsoft.graph.categoryColor
const (
	NONE MicrosoftGraphCategoryColor = "none"
	PRESET0 MicrosoftGraphCategoryColor = "preset0"
	PRESET1 MicrosoftGraphCategoryColor = "preset1"
	PRESET2 MicrosoftGraphCategoryColor = "preset2"
	PRESET3 MicrosoftGraphCategoryColor = "preset3"
	PRESET4 MicrosoftGraphCategoryColor = "preset4"
	PRESET5 MicrosoftGraphCategoryColor = "preset5"
	PRESET6 MicrosoftGraphCategoryColor = "preset6"
	PRESET7 MicrosoftGraphCategoryColor = "preset7"
	PRESET8 MicrosoftGraphCategoryColor = "preset8"
	PRESET9 MicrosoftGraphCategoryColor = "preset9"
	PRESET10 MicrosoftGraphCategoryColor = "preset10"
	PRESET11 MicrosoftGraphCategoryColor = "preset11"
	PRESET12 MicrosoftGraphCategoryColor = "preset12"
	PRESET13 MicrosoftGraphCategoryColor = "preset13"
	PRESET14 MicrosoftGraphCategoryColor = "preset14"
	PRESET15 MicrosoftGraphCategoryColor = "preset15"
	PRESET16 MicrosoftGraphCategoryColor = "preset16"
	PRESET17 MicrosoftGraphCategoryColor = "preset17"
	PRESET18 MicrosoftGraphCategoryColor = "preset18"
	PRESET19 MicrosoftGraphCategoryColor = "preset19"
	PRESET20 MicrosoftGraphCategoryColor = "preset20"
	PRESET21 MicrosoftGraphCategoryColor = "preset21"
	PRESET22 MicrosoftGraphCategoryColor = "preset22"
	PRESET23 MicrosoftGraphCategoryColor = "preset23"
	PRESET24 MicrosoftGraphCategoryColor = "preset24"
)

// All allowed values of MicrosoftGraphCategoryColor enum
var AllowedMicrosoftGraphCategoryColorEnumValues = []MicrosoftGraphCategoryColor{
	"none",
	"preset0",
	"preset1",
	"preset2",
	"preset3",
	"preset4",
	"preset5",
	"preset6",
	"preset7",
	"preset8",
	"preset9",
	"preset10",
	"preset11",
	"preset12",
	"preset13",
	"preset14",
	"preset15",
	"preset16",
	"preset17",
	"preset18",
	"preset19",
	"preset20",
	"preset21",
	"preset22",
	"preset23",
	"preset24",
}

func (v *MicrosoftGraphCategoryColor) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MicrosoftGraphCategoryColor(value)
	for _, existing := range AllowedMicrosoftGraphCategoryColorEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MicrosoftGraphCategoryColor", value)
}

// NewMicrosoftGraphCategoryColorFromValue returns a pointer to a valid MicrosoftGraphCategoryColor
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMicrosoftGraphCategoryColorFromValue(v string) (*MicrosoftGraphCategoryColor, error) {
	ev := MicrosoftGraphCategoryColor(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MicrosoftGraphCategoryColor: valid values are %v", v, AllowedMicrosoftGraphCategoryColorEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MicrosoftGraphCategoryColor) IsValid() bool {
	for _, existing := range AllowedMicrosoftGraphCategoryColorEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to microsoft.graph.categoryColor value
func (v MicrosoftGraphCategoryColor) Ptr() *MicrosoftGraphCategoryColor {
	return &v
}

type NullableMicrosoftGraphCategoryColor struct {
	value *MicrosoftGraphCategoryColor
	isSet bool
}

func (v NullableMicrosoftGraphCategoryColor) Get() *MicrosoftGraphCategoryColor {
	return v.value
}

func (v *NullableMicrosoftGraphCategoryColor) Set(val *MicrosoftGraphCategoryColor) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphCategoryColor) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphCategoryColor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphCategoryColor(val *MicrosoftGraphCategoryColor) *NullableMicrosoftGraphCategoryColor {
	return &NullableMicrosoftGraphCategoryColor{value: val, isSet: true}
}

func (v NullableMicrosoftGraphCategoryColor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphCategoryColor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
