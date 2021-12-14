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

// MicrosoftGraphDelegateMeetingMessageDeliveryOptions the model 'MicrosoftGraphDelegateMeetingMessageDeliveryOptions'
type MicrosoftGraphDelegateMeetingMessageDeliveryOptions string

// List of microsoft.graph.delegateMeetingMessageDeliveryOptions
const (
	SEND_TO_DELEGATE_AND_INFORMATION_TO_PRINCIPAL MicrosoftGraphDelegateMeetingMessageDeliveryOptions = "sendToDelegateAndInformationToPrincipal"
	SEND_TO_DELEGATE_AND_PRINCIPAL MicrosoftGraphDelegateMeetingMessageDeliveryOptions = "sendToDelegateAndPrincipal"
	SEND_TO_DELEGATE_ONLY MicrosoftGraphDelegateMeetingMessageDeliveryOptions = "sendToDelegateOnly"
)

// All allowed values of MicrosoftGraphDelegateMeetingMessageDeliveryOptions enum
var AllowedMicrosoftGraphDelegateMeetingMessageDeliveryOptionsEnumValues = []MicrosoftGraphDelegateMeetingMessageDeliveryOptions{
	"sendToDelegateAndInformationToPrincipal",
	"sendToDelegateAndPrincipal",
	"sendToDelegateOnly",
}

func (v *MicrosoftGraphDelegateMeetingMessageDeliveryOptions) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MicrosoftGraphDelegateMeetingMessageDeliveryOptions(value)
	for _, existing := range AllowedMicrosoftGraphDelegateMeetingMessageDeliveryOptionsEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MicrosoftGraphDelegateMeetingMessageDeliveryOptions", value)
}

// NewMicrosoftGraphDelegateMeetingMessageDeliveryOptionsFromValue returns a pointer to a valid MicrosoftGraphDelegateMeetingMessageDeliveryOptions
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMicrosoftGraphDelegateMeetingMessageDeliveryOptionsFromValue(v string) (*MicrosoftGraphDelegateMeetingMessageDeliveryOptions, error) {
	ev := MicrosoftGraphDelegateMeetingMessageDeliveryOptions(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MicrosoftGraphDelegateMeetingMessageDeliveryOptions: valid values are %v", v, AllowedMicrosoftGraphDelegateMeetingMessageDeliveryOptionsEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MicrosoftGraphDelegateMeetingMessageDeliveryOptions) IsValid() bool {
	for _, existing := range AllowedMicrosoftGraphDelegateMeetingMessageDeliveryOptionsEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to microsoft.graph.delegateMeetingMessageDeliveryOptions value
func (v MicrosoftGraphDelegateMeetingMessageDeliveryOptions) Ptr() *MicrosoftGraphDelegateMeetingMessageDeliveryOptions {
	return &v
}

type NullableMicrosoftGraphDelegateMeetingMessageDeliveryOptions struct {
	value *MicrosoftGraphDelegateMeetingMessageDeliveryOptions
	isSet bool
}

func (v NullableMicrosoftGraphDelegateMeetingMessageDeliveryOptions) Get() *MicrosoftGraphDelegateMeetingMessageDeliveryOptions {
	return v.value
}

func (v *NullableMicrosoftGraphDelegateMeetingMessageDeliveryOptions) Set(val *MicrosoftGraphDelegateMeetingMessageDeliveryOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphDelegateMeetingMessageDeliveryOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphDelegateMeetingMessageDeliveryOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphDelegateMeetingMessageDeliveryOptions(val *MicrosoftGraphDelegateMeetingMessageDeliveryOptions) *NullableMicrosoftGraphDelegateMeetingMessageDeliveryOptions {
	return &NullableMicrosoftGraphDelegateMeetingMessageDeliveryOptions{value: val, isSet: true}
}

func (v NullableMicrosoftGraphDelegateMeetingMessageDeliveryOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphDelegateMeetingMessageDeliveryOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

