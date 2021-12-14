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

// MicrosoftGraphConnectionDirection the model 'MicrosoftGraphConnectionDirection'
type MicrosoftGraphConnectionDirection string

// List of microsoft.graph.connectionDirection
const (
	UNKNOWN MicrosoftGraphConnectionDirection = "unknown"
	INBOUND MicrosoftGraphConnectionDirection = "inbound"
	OUTBOUND MicrosoftGraphConnectionDirection = "outbound"
	UNKNOWN_FUTURE_VALUE MicrosoftGraphConnectionDirection = "unknownFutureValue"
)

// All allowed values of MicrosoftGraphConnectionDirection enum
var AllowedMicrosoftGraphConnectionDirectionEnumValues = []MicrosoftGraphConnectionDirection{
	"unknown",
	"inbound",
	"outbound",
	"unknownFutureValue",
}

func (v *MicrosoftGraphConnectionDirection) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MicrosoftGraphConnectionDirection(value)
	for _, existing := range AllowedMicrosoftGraphConnectionDirectionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MicrosoftGraphConnectionDirection", value)
}

// NewMicrosoftGraphConnectionDirectionFromValue returns a pointer to a valid MicrosoftGraphConnectionDirection
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMicrosoftGraphConnectionDirectionFromValue(v string) (*MicrosoftGraphConnectionDirection, error) {
	ev := MicrosoftGraphConnectionDirection(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MicrosoftGraphConnectionDirection: valid values are %v", v, AllowedMicrosoftGraphConnectionDirectionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MicrosoftGraphConnectionDirection) IsValid() bool {
	for _, existing := range AllowedMicrosoftGraphConnectionDirectionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to microsoft.graph.connectionDirection value
func (v MicrosoftGraphConnectionDirection) Ptr() *MicrosoftGraphConnectionDirection {
	return &v
}

type NullableMicrosoftGraphConnectionDirection struct {
	value *MicrosoftGraphConnectionDirection
	isSet bool
}

func (v NullableMicrosoftGraphConnectionDirection) Get() *MicrosoftGraphConnectionDirection {
	return v.value
}

func (v *NullableMicrosoftGraphConnectionDirection) Set(val *MicrosoftGraphConnectionDirection) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphConnectionDirection) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphConnectionDirection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphConnectionDirection(val *MicrosoftGraphConnectionDirection) *NullableMicrosoftGraphConnectionDirection {
	return &NullableMicrosoftGraphConnectionDirection{value: val, isSet: true}
}

func (v NullableMicrosoftGraphConnectionDirection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphConnectionDirection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

