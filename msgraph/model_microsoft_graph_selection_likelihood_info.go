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

// MicrosoftGraphSelectionLikelihoodInfo the model 'MicrosoftGraphSelectionLikelihoodInfo'
type MicrosoftGraphSelectionLikelihoodInfo string

// List of microsoft.graph.selectionLikelihoodInfo
const (
	NOT_SPECIFIED MicrosoftGraphSelectionLikelihoodInfo = "notSpecified"
	HIGH MicrosoftGraphSelectionLikelihoodInfo = "high"
)

// All allowed values of MicrosoftGraphSelectionLikelihoodInfo enum
var AllowedMicrosoftGraphSelectionLikelihoodInfoEnumValues = []MicrosoftGraphSelectionLikelihoodInfo{
	"notSpecified",
	"high",
}

func (v *MicrosoftGraphSelectionLikelihoodInfo) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MicrosoftGraphSelectionLikelihoodInfo(value)
	for _, existing := range AllowedMicrosoftGraphSelectionLikelihoodInfoEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MicrosoftGraphSelectionLikelihoodInfo", value)
}

// NewMicrosoftGraphSelectionLikelihoodInfoFromValue returns a pointer to a valid MicrosoftGraphSelectionLikelihoodInfo
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMicrosoftGraphSelectionLikelihoodInfoFromValue(v string) (*MicrosoftGraphSelectionLikelihoodInfo, error) {
	ev := MicrosoftGraphSelectionLikelihoodInfo(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MicrosoftGraphSelectionLikelihoodInfo: valid values are %v", v, AllowedMicrosoftGraphSelectionLikelihoodInfoEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MicrosoftGraphSelectionLikelihoodInfo) IsValid() bool {
	for _, existing := range AllowedMicrosoftGraphSelectionLikelihoodInfoEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to microsoft.graph.selectionLikelihoodInfo value
func (v MicrosoftGraphSelectionLikelihoodInfo) Ptr() *MicrosoftGraphSelectionLikelihoodInfo {
	return &v
}

type NullableMicrosoftGraphSelectionLikelihoodInfo struct {
	value *MicrosoftGraphSelectionLikelihoodInfo
	isSet bool
}

func (v NullableMicrosoftGraphSelectionLikelihoodInfo) Get() *MicrosoftGraphSelectionLikelihoodInfo {
	return v.value
}

func (v *NullableMicrosoftGraphSelectionLikelihoodInfo) Set(val *MicrosoftGraphSelectionLikelihoodInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphSelectionLikelihoodInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphSelectionLikelihoodInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphSelectionLikelihoodInfo(val *MicrosoftGraphSelectionLikelihoodInfo) *NullableMicrosoftGraphSelectionLikelihoodInfo {
	return &NullableMicrosoftGraphSelectionLikelihoodInfo{value: val, isSet: true}
}

func (v NullableMicrosoftGraphSelectionLikelihoodInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphSelectionLikelihoodInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

