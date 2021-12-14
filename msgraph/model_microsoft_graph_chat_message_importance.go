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

// MicrosoftGraphChatMessageImportance the model 'MicrosoftGraphChatMessageImportance'
type MicrosoftGraphChatMessageImportance string

// List of microsoft.graph.chatMessageImportance
const (
	NORMAL MicrosoftGraphChatMessageImportance = "normal"
	HIGH MicrosoftGraphChatMessageImportance = "high"
	URGENT MicrosoftGraphChatMessageImportance = "urgent"
	UNKNOWN_FUTURE_VALUE MicrosoftGraphChatMessageImportance = "unknownFutureValue"
)

// All allowed values of MicrosoftGraphChatMessageImportance enum
var AllowedMicrosoftGraphChatMessageImportanceEnumValues = []MicrosoftGraphChatMessageImportance{
	"normal",
	"high",
	"urgent",
	"unknownFutureValue",
}

func (v *MicrosoftGraphChatMessageImportance) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MicrosoftGraphChatMessageImportance(value)
	for _, existing := range AllowedMicrosoftGraphChatMessageImportanceEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MicrosoftGraphChatMessageImportance", value)
}

// NewMicrosoftGraphChatMessageImportanceFromValue returns a pointer to a valid MicrosoftGraphChatMessageImportance
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMicrosoftGraphChatMessageImportanceFromValue(v string) (*MicrosoftGraphChatMessageImportance, error) {
	ev := MicrosoftGraphChatMessageImportance(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MicrosoftGraphChatMessageImportance: valid values are %v", v, AllowedMicrosoftGraphChatMessageImportanceEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MicrosoftGraphChatMessageImportance) IsValid() bool {
	for _, existing := range AllowedMicrosoftGraphChatMessageImportanceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to microsoft.graph.chatMessageImportance value
func (v MicrosoftGraphChatMessageImportance) Ptr() *MicrosoftGraphChatMessageImportance {
	return &v
}

type NullableMicrosoftGraphChatMessageImportance struct {
	value *MicrosoftGraphChatMessageImportance
	isSet bool
}

func (v NullableMicrosoftGraphChatMessageImportance) Get() *MicrosoftGraphChatMessageImportance {
	return v.value
}

func (v *NullableMicrosoftGraphChatMessageImportance) Set(val *MicrosoftGraphChatMessageImportance) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphChatMessageImportance) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphChatMessageImportance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphChatMessageImportance(val *MicrosoftGraphChatMessageImportance) *NullableMicrosoftGraphChatMessageImportance {
	return &NullableMicrosoftGraphChatMessageImportance{value: val, isSet: true}
}

func (v NullableMicrosoftGraphChatMessageImportance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphChatMessageImportance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

