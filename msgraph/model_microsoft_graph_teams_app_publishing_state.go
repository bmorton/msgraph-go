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

// MicrosoftGraphTeamsAppPublishingState the model 'MicrosoftGraphTeamsAppPublishingState'
type MicrosoftGraphTeamsAppPublishingState string

// List of microsoft.graph.teamsAppPublishingState
const (
	SUBMITTED MicrosoftGraphTeamsAppPublishingState = "submitted"
	REJECTED MicrosoftGraphTeamsAppPublishingState = "rejected"
	PUBLISHED MicrosoftGraphTeamsAppPublishingState = "published"
	UNKNOWN_FUTURE_VALUE MicrosoftGraphTeamsAppPublishingState = "unknownFutureValue"
)

// All allowed values of MicrosoftGraphTeamsAppPublishingState enum
var AllowedMicrosoftGraphTeamsAppPublishingStateEnumValues = []MicrosoftGraphTeamsAppPublishingState{
	"submitted",
	"rejected",
	"published",
	"unknownFutureValue",
}

func (v *MicrosoftGraphTeamsAppPublishingState) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MicrosoftGraphTeamsAppPublishingState(value)
	for _, existing := range AllowedMicrosoftGraphTeamsAppPublishingStateEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MicrosoftGraphTeamsAppPublishingState", value)
}

// NewMicrosoftGraphTeamsAppPublishingStateFromValue returns a pointer to a valid MicrosoftGraphTeamsAppPublishingState
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMicrosoftGraphTeamsAppPublishingStateFromValue(v string) (*MicrosoftGraphTeamsAppPublishingState, error) {
	ev := MicrosoftGraphTeamsAppPublishingState(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MicrosoftGraphTeamsAppPublishingState: valid values are %v", v, AllowedMicrosoftGraphTeamsAppPublishingStateEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MicrosoftGraphTeamsAppPublishingState) IsValid() bool {
	for _, existing := range AllowedMicrosoftGraphTeamsAppPublishingStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to microsoft.graph.teamsAppPublishingState value
func (v MicrosoftGraphTeamsAppPublishingState) Ptr() *MicrosoftGraphTeamsAppPublishingState {
	return &v
}

type NullableMicrosoftGraphTeamsAppPublishingState struct {
	value *MicrosoftGraphTeamsAppPublishingState
	isSet bool
}

func (v NullableMicrosoftGraphTeamsAppPublishingState) Get() *MicrosoftGraphTeamsAppPublishingState {
	return v.value
}

func (v *NullableMicrosoftGraphTeamsAppPublishingState) Set(val *MicrosoftGraphTeamsAppPublishingState) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphTeamsAppPublishingState) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphTeamsAppPublishingState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphTeamsAppPublishingState(val *MicrosoftGraphTeamsAppPublishingState) *NullableMicrosoftGraphTeamsAppPublishingState {
	return &NullableMicrosoftGraphTeamsAppPublishingState{value: val, isSet: true}
}

func (v NullableMicrosoftGraphTeamsAppPublishingState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphTeamsAppPublishingState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

