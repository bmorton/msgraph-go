/*
Partial Graph API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package msgraph

import (
	"encoding/json"
	"fmt"
)

// MicrosoftGraphChatMessagePolicyViolationUserActionTypes the model 'MicrosoftGraphChatMessagePolicyViolationUserActionTypes'
type MicrosoftGraphChatMessagePolicyViolationUserActionTypes string

// List of microsoft.graph.chatMessagePolicyViolationUserActionTypes
const (
	NONE MicrosoftGraphChatMessagePolicyViolationUserActionTypes = "none"
	OVERRIDE MicrosoftGraphChatMessagePolicyViolationUserActionTypes = "override"
	REPORT_FALSE_POSITIVE MicrosoftGraphChatMessagePolicyViolationUserActionTypes = "reportFalsePositive"
)

// All allowed values of MicrosoftGraphChatMessagePolicyViolationUserActionTypes enum
var AllowedMicrosoftGraphChatMessagePolicyViolationUserActionTypesEnumValues = []MicrosoftGraphChatMessagePolicyViolationUserActionTypes{
	"none",
	"override",
	"reportFalsePositive",
}

func (v *MicrosoftGraphChatMessagePolicyViolationUserActionTypes) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MicrosoftGraphChatMessagePolicyViolationUserActionTypes(value)
	for _, existing := range AllowedMicrosoftGraphChatMessagePolicyViolationUserActionTypesEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MicrosoftGraphChatMessagePolicyViolationUserActionTypes", value)
}

// NewMicrosoftGraphChatMessagePolicyViolationUserActionTypesFromValue returns a pointer to a valid MicrosoftGraphChatMessagePolicyViolationUserActionTypes
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMicrosoftGraphChatMessagePolicyViolationUserActionTypesFromValue(v string) (*MicrosoftGraphChatMessagePolicyViolationUserActionTypes, error) {
	ev := MicrosoftGraphChatMessagePolicyViolationUserActionTypes(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MicrosoftGraphChatMessagePolicyViolationUserActionTypes: valid values are %v", v, AllowedMicrosoftGraphChatMessagePolicyViolationUserActionTypesEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MicrosoftGraphChatMessagePolicyViolationUserActionTypes) IsValid() bool {
	for _, existing := range AllowedMicrosoftGraphChatMessagePolicyViolationUserActionTypesEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to microsoft.graph.chatMessagePolicyViolationUserActionTypes value
func (v MicrosoftGraphChatMessagePolicyViolationUserActionTypes) Ptr() *MicrosoftGraphChatMessagePolicyViolationUserActionTypes {
	return &v
}

type NullableMicrosoftGraphChatMessagePolicyViolationUserActionTypes struct {
	value *MicrosoftGraphChatMessagePolicyViolationUserActionTypes
	isSet bool
}

func (v NullableMicrosoftGraphChatMessagePolicyViolationUserActionTypes) Get() *MicrosoftGraphChatMessagePolicyViolationUserActionTypes {
	return v.value
}

func (v *NullableMicrosoftGraphChatMessagePolicyViolationUserActionTypes) Set(val *MicrosoftGraphChatMessagePolicyViolationUserActionTypes) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphChatMessagePolicyViolationUserActionTypes) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphChatMessagePolicyViolationUserActionTypes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphChatMessagePolicyViolationUserActionTypes(val *MicrosoftGraphChatMessagePolicyViolationUserActionTypes) *NullableMicrosoftGraphChatMessagePolicyViolationUserActionTypes {
	return &NullableMicrosoftGraphChatMessagePolicyViolationUserActionTypes{value: val, isSet: true}
}

func (v NullableMicrosoftGraphChatMessagePolicyViolationUserActionTypes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphChatMessagePolicyViolationUserActionTypes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

