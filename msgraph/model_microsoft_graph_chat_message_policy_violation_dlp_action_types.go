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

// MicrosoftGraphChatMessagePolicyViolationDlpActionTypes the model 'MicrosoftGraphChatMessagePolicyViolationDlpActionTypes'
type MicrosoftGraphChatMessagePolicyViolationDlpActionTypes string

// List of microsoft.graph.chatMessagePolicyViolationDlpActionTypes
const (
	NONE MicrosoftGraphChatMessagePolicyViolationDlpActionTypes = "none"
	NOTIFY_SENDER MicrosoftGraphChatMessagePolicyViolationDlpActionTypes = "notifySender"
	BLOCK_ACCESS MicrosoftGraphChatMessagePolicyViolationDlpActionTypes = "blockAccess"
	BLOCK_ACCESS_EXTERNAL MicrosoftGraphChatMessagePolicyViolationDlpActionTypes = "blockAccessExternal"
)

// All allowed values of MicrosoftGraphChatMessagePolicyViolationDlpActionTypes enum
var AllowedMicrosoftGraphChatMessagePolicyViolationDlpActionTypesEnumValues = []MicrosoftGraphChatMessagePolicyViolationDlpActionTypes{
	"none",
	"notifySender",
	"blockAccess",
	"blockAccessExternal",
}

func (v *MicrosoftGraphChatMessagePolicyViolationDlpActionTypes) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MicrosoftGraphChatMessagePolicyViolationDlpActionTypes(value)
	for _, existing := range AllowedMicrosoftGraphChatMessagePolicyViolationDlpActionTypesEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MicrosoftGraphChatMessagePolicyViolationDlpActionTypes", value)
}

// NewMicrosoftGraphChatMessagePolicyViolationDlpActionTypesFromValue returns a pointer to a valid MicrosoftGraphChatMessagePolicyViolationDlpActionTypes
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMicrosoftGraphChatMessagePolicyViolationDlpActionTypesFromValue(v string) (*MicrosoftGraphChatMessagePolicyViolationDlpActionTypes, error) {
	ev := MicrosoftGraphChatMessagePolicyViolationDlpActionTypes(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MicrosoftGraphChatMessagePolicyViolationDlpActionTypes: valid values are %v", v, AllowedMicrosoftGraphChatMessagePolicyViolationDlpActionTypesEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MicrosoftGraphChatMessagePolicyViolationDlpActionTypes) IsValid() bool {
	for _, existing := range AllowedMicrosoftGraphChatMessagePolicyViolationDlpActionTypesEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to microsoft.graph.chatMessagePolicyViolationDlpActionTypes value
func (v MicrosoftGraphChatMessagePolicyViolationDlpActionTypes) Ptr() *MicrosoftGraphChatMessagePolicyViolationDlpActionTypes {
	return &v
}

type NullableMicrosoftGraphChatMessagePolicyViolationDlpActionTypes struct {
	value *MicrosoftGraphChatMessagePolicyViolationDlpActionTypes
	isSet bool
}

func (v NullableMicrosoftGraphChatMessagePolicyViolationDlpActionTypes) Get() *MicrosoftGraphChatMessagePolicyViolationDlpActionTypes {
	return v.value
}

func (v *NullableMicrosoftGraphChatMessagePolicyViolationDlpActionTypes) Set(val *MicrosoftGraphChatMessagePolicyViolationDlpActionTypes) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphChatMessagePolicyViolationDlpActionTypes) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphChatMessagePolicyViolationDlpActionTypes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphChatMessagePolicyViolationDlpActionTypes(val *MicrosoftGraphChatMessagePolicyViolationDlpActionTypes) *NullableMicrosoftGraphChatMessagePolicyViolationDlpActionTypes {
	return &NullableMicrosoftGraphChatMessagePolicyViolationDlpActionTypes{value: val, isSet: true}
}

func (v NullableMicrosoftGraphChatMessagePolicyViolationDlpActionTypes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphChatMessagePolicyViolationDlpActionTypes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

