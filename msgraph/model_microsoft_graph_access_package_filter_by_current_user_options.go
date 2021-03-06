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

// MicrosoftGraphAccessPackageFilterByCurrentUserOptions the model 'MicrosoftGraphAccessPackageFilterByCurrentUserOptions'
type MicrosoftGraphAccessPackageFilterByCurrentUserOptions string

// List of microsoft.graph.accessPackageFilterByCurrentUserOptions
const (
	ALLOWED_REQUESTOR MicrosoftGraphAccessPackageFilterByCurrentUserOptions = "allowedRequestor"
	UNKNOWN_FUTURE_VALUE MicrosoftGraphAccessPackageFilterByCurrentUserOptions = "unknownFutureValue"
)

// All allowed values of MicrosoftGraphAccessPackageFilterByCurrentUserOptions enum
var AllowedMicrosoftGraphAccessPackageFilterByCurrentUserOptionsEnumValues = []MicrosoftGraphAccessPackageFilterByCurrentUserOptions{
	"allowedRequestor",
	"unknownFutureValue",
}

func (v *MicrosoftGraphAccessPackageFilterByCurrentUserOptions) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MicrosoftGraphAccessPackageFilterByCurrentUserOptions(value)
	for _, existing := range AllowedMicrosoftGraphAccessPackageFilterByCurrentUserOptionsEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MicrosoftGraphAccessPackageFilterByCurrentUserOptions", value)
}

// NewMicrosoftGraphAccessPackageFilterByCurrentUserOptionsFromValue returns a pointer to a valid MicrosoftGraphAccessPackageFilterByCurrentUserOptions
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMicrosoftGraphAccessPackageFilterByCurrentUserOptionsFromValue(v string) (*MicrosoftGraphAccessPackageFilterByCurrentUserOptions, error) {
	ev := MicrosoftGraphAccessPackageFilterByCurrentUserOptions(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MicrosoftGraphAccessPackageFilterByCurrentUserOptions: valid values are %v", v, AllowedMicrosoftGraphAccessPackageFilterByCurrentUserOptionsEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MicrosoftGraphAccessPackageFilterByCurrentUserOptions) IsValid() bool {
	for _, existing := range AllowedMicrosoftGraphAccessPackageFilterByCurrentUserOptionsEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to microsoft.graph.accessPackageFilterByCurrentUserOptions value
func (v MicrosoftGraphAccessPackageFilterByCurrentUserOptions) Ptr() *MicrosoftGraphAccessPackageFilterByCurrentUserOptions {
	return &v
}

type NullableMicrosoftGraphAccessPackageFilterByCurrentUserOptions struct {
	value *MicrosoftGraphAccessPackageFilterByCurrentUserOptions
	isSet bool
}

func (v NullableMicrosoftGraphAccessPackageFilterByCurrentUserOptions) Get() *MicrosoftGraphAccessPackageFilterByCurrentUserOptions {
	return v.value
}

func (v *NullableMicrosoftGraphAccessPackageFilterByCurrentUserOptions) Set(val *MicrosoftGraphAccessPackageFilterByCurrentUserOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphAccessPackageFilterByCurrentUserOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphAccessPackageFilterByCurrentUserOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphAccessPackageFilterByCurrentUserOptions(val *MicrosoftGraphAccessPackageFilterByCurrentUserOptions) *NullableMicrosoftGraphAccessPackageFilterByCurrentUserOptions {
	return &NullableMicrosoftGraphAccessPackageFilterByCurrentUserOptions{value: val, isSet: true}
}

func (v NullableMicrosoftGraphAccessPackageFilterByCurrentUserOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphAccessPackageFilterByCurrentUserOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

