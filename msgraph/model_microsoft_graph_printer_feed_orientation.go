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

// MicrosoftGraphPrinterFeedOrientation the model 'MicrosoftGraphPrinterFeedOrientation'
type MicrosoftGraphPrinterFeedOrientation string

// List of microsoft.graph.printerFeedOrientation
const (
	LONG_EDGE_FIRST MicrosoftGraphPrinterFeedOrientation = "longEdgeFirst"
	SHORT_EDGE_FIRST MicrosoftGraphPrinterFeedOrientation = "shortEdgeFirst"
	UNKNOWN_FUTURE_VALUE MicrosoftGraphPrinterFeedOrientation = "unknownFutureValue"
)

// All allowed values of MicrosoftGraphPrinterFeedOrientation enum
var AllowedMicrosoftGraphPrinterFeedOrientationEnumValues = []MicrosoftGraphPrinterFeedOrientation{
	"longEdgeFirst",
	"shortEdgeFirst",
	"unknownFutureValue",
}

func (v *MicrosoftGraphPrinterFeedOrientation) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MicrosoftGraphPrinterFeedOrientation(value)
	for _, existing := range AllowedMicrosoftGraphPrinterFeedOrientationEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MicrosoftGraphPrinterFeedOrientation", value)
}

// NewMicrosoftGraphPrinterFeedOrientationFromValue returns a pointer to a valid MicrosoftGraphPrinterFeedOrientation
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMicrosoftGraphPrinterFeedOrientationFromValue(v string) (*MicrosoftGraphPrinterFeedOrientation, error) {
	ev := MicrosoftGraphPrinterFeedOrientation(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MicrosoftGraphPrinterFeedOrientation: valid values are %v", v, AllowedMicrosoftGraphPrinterFeedOrientationEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MicrosoftGraphPrinterFeedOrientation) IsValid() bool {
	for _, existing := range AllowedMicrosoftGraphPrinterFeedOrientationEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to microsoft.graph.printerFeedOrientation value
func (v MicrosoftGraphPrinterFeedOrientation) Ptr() *MicrosoftGraphPrinterFeedOrientation {
	return &v
}

type NullableMicrosoftGraphPrinterFeedOrientation struct {
	value *MicrosoftGraphPrinterFeedOrientation
	isSet bool
}

func (v NullableMicrosoftGraphPrinterFeedOrientation) Get() *MicrosoftGraphPrinterFeedOrientation {
	return v.value
}

func (v *NullableMicrosoftGraphPrinterFeedOrientation) Set(val *MicrosoftGraphPrinterFeedOrientation) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphPrinterFeedOrientation) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphPrinterFeedOrientation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphPrinterFeedOrientation(val *MicrosoftGraphPrinterFeedOrientation) *NullableMicrosoftGraphPrinterFeedOrientation {
	return &NullableMicrosoftGraphPrinterFeedOrientation{value: val, isSet: true}
}

func (v NullableMicrosoftGraphPrinterFeedOrientation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphPrinterFeedOrientation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

