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

// MicrosoftGraphTeamsAppDistributionMethod the model 'MicrosoftGraphTeamsAppDistributionMethod'
type MicrosoftGraphTeamsAppDistributionMethod string

// List of microsoft.graph.teamsAppDistributionMethod
const (
	STORE MicrosoftGraphTeamsAppDistributionMethod = "store"
	ORGANIZATION MicrosoftGraphTeamsAppDistributionMethod = "organization"
	SIDELOADED MicrosoftGraphTeamsAppDistributionMethod = "sideloaded"
	UNKNOWN_FUTURE_VALUE MicrosoftGraphTeamsAppDistributionMethod = "unknownFutureValue"
)

// All allowed values of MicrosoftGraphTeamsAppDistributionMethod enum
var AllowedMicrosoftGraphTeamsAppDistributionMethodEnumValues = []MicrosoftGraphTeamsAppDistributionMethod{
	"store",
	"organization",
	"sideloaded",
	"unknownFutureValue",
}

func (v *MicrosoftGraphTeamsAppDistributionMethod) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MicrosoftGraphTeamsAppDistributionMethod(value)
	for _, existing := range AllowedMicrosoftGraphTeamsAppDistributionMethodEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MicrosoftGraphTeamsAppDistributionMethod", value)
}

// NewMicrosoftGraphTeamsAppDistributionMethodFromValue returns a pointer to a valid MicrosoftGraphTeamsAppDistributionMethod
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMicrosoftGraphTeamsAppDistributionMethodFromValue(v string) (*MicrosoftGraphTeamsAppDistributionMethod, error) {
	ev := MicrosoftGraphTeamsAppDistributionMethod(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MicrosoftGraphTeamsAppDistributionMethod: valid values are %v", v, AllowedMicrosoftGraphTeamsAppDistributionMethodEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MicrosoftGraphTeamsAppDistributionMethod) IsValid() bool {
	for _, existing := range AllowedMicrosoftGraphTeamsAppDistributionMethodEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to microsoft.graph.teamsAppDistributionMethod value
func (v MicrosoftGraphTeamsAppDistributionMethod) Ptr() *MicrosoftGraphTeamsAppDistributionMethod {
	return &v
}

type NullableMicrosoftGraphTeamsAppDistributionMethod struct {
	value *MicrosoftGraphTeamsAppDistributionMethod
	isSet bool
}

func (v NullableMicrosoftGraphTeamsAppDistributionMethod) Get() *MicrosoftGraphTeamsAppDistributionMethod {
	return v.value
}

func (v *NullableMicrosoftGraphTeamsAppDistributionMethod) Set(val *MicrosoftGraphTeamsAppDistributionMethod) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphTeamsAppDistributionMethod) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphTeamsAppDistributionMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphTeamsAppDistributionMethod(val *MicrosoftGraphTeamsAppDistributionMethod) *NullableMicrosoftGraphTeamsAppDistributionMethod {
	return &NullableMicrosoftGraphTeamsAppDistributionMethod{value: val, isSet: true}
}

func (v NullableMicrosoftGraphTeamsAppDistributionMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphTeamsAppDistributionMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
