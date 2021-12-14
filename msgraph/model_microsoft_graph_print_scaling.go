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

// MicrosoftGraphPrintScaling the model 'MicrosoftGraphPrintScaling'
type MicrosoftGraphPrintScaling string

// List of microsoft.graph.printScaling
const (
	AUTO MicrosoftGraphPrintScaling = "auto"
	SHRINK_TO_FIT MicrosoftGraphPrintScaling = "shrinkToFit"
	FILL MicrosoftGraphPrintScaling = "fill"
	FIT MicrosoftGraphPrintScaling = "fit"
	NONE MicrosoftGraphPrintScaling = "none"
	UNKNOWN_FUTURE_VALUE MicrosoftGraphPrintScaling = "unknownFutureValue"
)

// All allowed values of MicrosoftGraphPrintScaling enum
var AllowedMicrosoftGraphPrintScalingEnumValues = []MicrosoftGraphPrintScaling{
	"auto",
	"shrinkToFit",
	"fill",
	"fit",
	"none",
	"unknownFutureValue",
}

func (v *MicrosoftGraphPrintScaling) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MicrosoftGraphPrintScaling(value)
	for _, existing := range AllowedMicrosoftGraphPrintScalingEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MicrosoftGraphPrintScaling", value)
}

// NewMicrosoftGraphPrintScalingFromValue returns a pointer to a valid MicrosoftGraphPrintScaling
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMicrosoftGraphPrintScalingFromValue(v string) (*MicrosoftGraphPrintScaling, error) {
	ev := MicrosoftGraphPrintScaling(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MicrosoftGraphPrintScaling: valid values are %v", v, AllowedMicrosoftGraphPrintScalingEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MicrosoftGraphPrintScaling) IsValid() bool {
	for _, existing := range AllowedMicrosoftGraphPrintScalingEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to microsoft.graph.printScaling value
func (v MicrosoftGraphPrintScaling) Ptr() *MicrosoftGraphPrintScaling {
	return &v
}

type NullableMicrosoftGraphPrintScaling struct {
	value *MicrosoftGraphPrintScaling
	isSet bool
}

func (v NullableMicrosoftGraphPrintScaling) Get() *MicrosoftGraphPrintScaling {
	return v.value
}

func (v *NullableMicrosoftGraphPrintScaling) Set(val *MicrosoftGraphPrintScaling) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphPrintScaling) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphPrintScaling) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphPrintScaling(val *MicrosoftGraphPrintScaling) *NullableMicrosoftGraphPrintScaling {
	return &NullableMicrosoftGraphPrintScaling{value: val, isSet: true}
}

func (v NullableMicrosoftGraphPrintScaling) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphPrintScaling) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

