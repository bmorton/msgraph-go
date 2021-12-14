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

// MicrosoftGraphMeetingChatMode the model 'MicrosoftGraphMeetingChatMode'
type MicrosoftGraphMeetingChatMode string

// List of microsoft.graph.meetingChatMode
const (
	ENABLED MicrosoftGraphMeetingChatMode = "enabled"
	DISABLED MicrosoftGraphMeetingChatMode = "disabled"
	LIMITED MicrosoftGraphMeetingChatMode = "limited"
	UNKNOWN_FUTURE_VALUE MicrosoftGraphMeetingChatMode = "unknownFutureValue"
)

// All allowed values of MicrosoftGraphMeetingChatMode enum
var AllowedMicrosoftGraphMeetingChatModeEnumValues = []MicrosoftGraphMeetingChatMode{
	"enabled",
	"disabled",
	"limited",
	"unknownFutureValue",
}

func (v *MicrosoftGraphMeetingChatMode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MicrosoftGraphMeetingChatMode(value)
	for _, existing := range AllowedMicrosoftGraphMeetingChatModeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MicrosoftGraphMeetingChatMode", value)
}

// NewMicrosoftGraphMeetingChatModeFromValue returns a pointer to a valid MicrosoftGraphMeetingChatMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMicrosoftGraphMeetingChatModeFromValue(v string) (*MicrosoftGraphMeetingChatMode, error) {
	ev := MicrosoftGraphMeetingChatMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MicrosoftGraphMeetingChatMode: valid values are %v", v, AllowedMicrosoftGraphMeetingChatModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MicrosoftGraphMeetingChatMode) IsValid() bool {
	for _, existing := range AllowedMicrosoftGraphMeetingChatModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to microsoft.graph.meetingChatMode value
func (v MicrosoftGraphMeetingChatMode) Ptr() *MicrosoftGraphMeetingChatMode {
	return &v
}

type NullableMicrosoftGraphMeetingChatMode struct {
	value *MicrosoftGraphMeetingChatMode
	isSet bool
}

func (v NullableMicrosoftGraphMeetingChatMode) Get() *MicrosoftGraphMeetingChatMode {
	return v.value
}

func (v *NullableMicrosoftGraphMeetingChatMode) Set(val *MicrosoftGraphMeetingChatMode) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphMeetingChatMode) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphMeetingChatMode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphMeetingChatMode(val *MicrosoftGraphMeetingChatMode) *NullableMicrosoftGraphMeetingChatMode {
	return &NullableMicrosoftGraphMeetingChatMode{value: val, isSet: true}
}

func (v NullableMicrosoftGraphMeetingChatMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphMeetingChatMode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
