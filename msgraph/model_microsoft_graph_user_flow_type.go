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

// MicrosoftGraphUserFlowType the model 'MicrosoftGraphUserFlowType'
type MicrosoftGraphUserFlowType string

// List of microsoft.graph.userFlowType
const (
	SIGN_UP MicrosoftGraphUserFlowType = "signUp"
	SIGN_IN MicrosoftGraphUserFlowType = "signIn"
	SIGN_UP_OR_SIGN_IN MicrosoftGraphUserFlowType = "signUpOrSignIn"
	PASSWORD_RESET MicrosoftGraphUserFlowType = "passwordReset"
	PROFILE_UPDATE MicrosoftGraphUserFlowType = "profileUpdate"
	RESOURCE_OWNER MicrosoftGraphUserFlowType = "resourceOwner"
	UNKNOWN_FUTURE_VALUE MicrosoftGraphUserFlowType = "unknownFutureValue"
)

// All allowed values of MicrosoftGraphUserFlowType enum
var AllowedMicrosoftGraphUserFlowTypeEnumValues = []MicrosoftGraphUserFlowType{
	"signUp",
	"signIn",
	"signUpOrSignIn",
	"passwordReset",
	"profileUpdate",
	"resourceOwner",
	"unknownFutureValue",
}

func (v *MicrosoftGraphUserFlowType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MicrosoftGraphUserFlowType(value)
	for _, existing := range AllowedMicrosoftGraphUserFlowTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MicrosoftGraphUserFlowType", value)
}

// NewMicrosoftGraphUserFlowTypeFromValue returns a pointer to a valid MicrosoftGraphUserFlowType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMicrosoftGraphUserFlowTypeFromValue(v string) (*MicrosoftGraphUserFlowType, error) {
	ev := MicrosoftGraphUserFlowType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MicrosoftGraphUserFlowType: valid values are %v", v, AllowedMicrosoftGraphUserFlowTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MicrosoftGraphUserFlowType) IsValid() bool {
	for _, existing := range AllowedMicrosoftGraphUserFlowTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to microsoft.graph.userFlowType value
func (v MicrosoftGraphUserFlowType) Ptr() *MicrosoftGraphUserFlowType {
	return &v
}

type NullableMicrosoftGraphUserFlowType struct {
	value *MicrosoftGraphUserFlowType
	isSet bool
}

func (v NullableMicrosoftGraphUserFlowType) Get() *MicrosoftGraphUserFlowType {
	return v.value
}

func (v *NullableMicrosoftGraphUserFlowType) Set(val *MicrosoftGraphUserFlowType) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphUserFlowType) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphUserFlowType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphUserFlowType(val *MicrosoftGraphUserFlowType) *NullableMicrosoftGraphUserFlowType {
	return &NullableMicrosoftGraphUserFlowType{value: val, isSet: true}
}

func (v NullableMicrosoftGraphUserFlowType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphUserFlowType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

