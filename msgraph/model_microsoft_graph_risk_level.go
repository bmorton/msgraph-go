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

// MicrosoftGraphRiskLevel the model 'MicrosoftGraphRiskLevel'
type MicrosoftGraphRiskLevel string

// List of microsoft.graph.riskLevel
const (
	LOW MicrosoftGraphRiskLevel = "low"
	MEDIUM MicrosoftGraphRiskLevel = "medium"
	HIGH MicrosoftGraphRiskLevel = "high"
	HIDDEN MicrosoftGraphRiskLevel = "hidden"
	NONE MicrosoftGraphRiskLevel = "none"
	UNKNOWN_FUTURE_VALUE MicrosoftGraphRiskLevel = "unknownFutureValue"
)

// All allowed values of MicrosoftGraphRiskLevel enum
var AllowedMicrosoftGraphRiskLevelEnumValues = []MicrosoftGraphRiskLevel{
	"low",
	"medium",
	"high",
	"hidden",
	"none",
	"unknownFutureValue",
}

func (v *MicrosoftGraphRiskLevel) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MicrosoftGraphRiskLevel(value)
	for _, existing := range AllowedMicrosoftGraphRiskLevelEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MicrosoftGraphRiskLevel", value)
}

// NewMicrosoftGraphRiskLevelFromValue returns a pointer to a valid MicrosoftGraphRiskLevel
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMicrosoftGraphRiskLevelFromValue(v string) (*MicrosoftGraphRiskLevel, error) {
	ev := MicrosoftGraphRiskLevel(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MicrosoftGraphRiskLevel: valid values are %v", v, AllowedMicrosoftGraphRiskLevelEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MicrosoftGraphRiskLevel) IsValid() bool {
	for _, existing := range AllowedMicrosoftGraphRiskLevelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to microsoft.graph.riskLevel value
func (v MicrosoftGraphRiskLevel) Ptr() *MicrosoftGraphRiskLevel {
	return &v
}

type NullableMicrosoftGraphRiskLevel struct {
	value *MicrosoftGraphRiskLevel
	isSet bool
}

func (v NullableMicrosoftGraphRiskLevel) Get() *MicrosoftGraphRiskLevel {
	return v.value
}

func (v *NullableMicrosoftGraphRiskLevel) Set(val *MicrosoftGraphRiskLevel) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphRiskLevel) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphRiskLevel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphRiskLevel(val *MicrosoftGraphRiskLevel) *NullableMicrosoftGraphRiskLevel {
	return &NullableMicrosoftGraphRiskLevel{value: val, isSet: true}
}

func (v NullableMicrosoftGraphRiskLevel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphRiskLevel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

