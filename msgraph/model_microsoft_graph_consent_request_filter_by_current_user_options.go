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

// MicrosoftGraphConsentRequestFilterByCurrentUserOptions the model 'MicrosoftGraphConsentRequestFilterByCurrentUserOptions'
type MicrosoftGraphConsentRequestFilterByCurrentUserOptions string

// List of microsoft.graph.consentRequestFilterByCurrentUserOptions
const (
	REVIEWER MicrosoftGraphConsentRequestFilterByCurrentUserOptions = "reviewer"
	UNKNOWN_FUTURE_VALUE MicrosoftGraphConsentRequestFilterByCurrentUserOptions = "unknownFutureValue"
)

// All allowed values of MicrosoftGraphConsentRequestFilterByCurrentUserOptions enum
var AllowedMicrosoftGraphConsentRequestFilterByCurrentUserOptionsEnumValues = []MicrosoftGraphConsentRequestFilterByCurrentUserOptions{
	"reviewer",
	"unknownFutureValue",
}

func (v *MicrosoftGraphConsentRequestFilterByCurrentUserOptions) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MicrosoftGraphConsentRequestFilterByCurrentUserOptions(value)
	for _, existing := range AllowedMicrosoftGraphConsentRequestFilterByCurrentUserOptionsEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MicrosoftGraphConsentRequestFilterByCurrentUserOptions", value)
}

// NewMicrosoftGraphConsentRequestFilterByCurrentUserOptionsFromValue returns a pointer to a valid MicrosoftGraphConsentRequestFilterByCurrentUserOptions
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMicrosoftGraphConsentRequestFilterByCurrentUserOptionsFromValue(v string) (*MicrosoftGraphConsentRequestFilterByCurrentUserOptions, error) {
	ev := MicrosoftGraphConsentRequestFilterByCurrentUserOptions(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MicrosoftGraphConsentRequestFilterByCurrentUserOptions: valid values are %v", v, AllowedMicrosoftGraphConsentRequestFilterByCurrentUserOptionsEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MicrosoftGraphConsentRequestFilterByCurrentUserOptions) IsValid() bool {
	for _, existing := range AllowedMicrosoftGraphConsentRequestFilterByCurrentUserOptionsEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to microsoft.graph.consentRequestFilterByCurrentUserOptions value
func (v MicrosoftGraphConsentRequestFilterByCurrentUserOptions) Ptr() *MicrosoftGraphConsentRequestFilterByCurrentUserOptions {
	return &v
}

type NullableMicrosoftGraphConsentRequestFilterByCurrentUserOptions struct {
	value *MicrosoftGraphConsentRequestFilterByCurrentUserOptions
	isSet bool
}

func (v NullableMicrosoftGraphConsentRequestFilterByCurrentUserOptions) Get() *MicrosoftGraphConsentRequestFilterByCurrentUserOptions {
	return v.value
}

func (v *NullableMicrosoftGraphConsentRequestFilterByCurrentUserOptions) Set(val *MicrosoftGraphConsentRequestFilterByCurrentUserOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphConsentRequestFilterByCurrentUserOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphConsentRequestFilterByCurrentUserOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphConsentRequestFilterByCurrentUserOptions(val *MicrosoftGraphConsentRequestFilterByCurrentUserOptions) *NullableMicrosoftGraphConsentRequestFilterByCurrentUserOptions {
	return &NullableMicrosoftGraphConsentRequestFilterByCurrentUserOptions{value: val, isSet: true}
}

func (v NullableMicrosoftGraphConsentRequestFilterByCurrentUserOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphConsentRequestFilterByCurrentUserOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
