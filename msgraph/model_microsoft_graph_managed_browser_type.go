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

// MicrosoftGraphManagedBrowserType Type of managed browser
type MicrosoftGraphManagedBrowserType string

// List of microsoft.graph.managedBrowserType
const (
	NOT_CONFIGURED MicrosoftGraphManagedBrowserType = "notConfigured"
	MICROSOFT_EDGE MicrosoftGraphManagedBrowserType = "microsoftEdge"
)

// All allowed values of MicrosoftGraphManagedBrowserType enum
var AllowedMicrosoftGraphManagedBrowserTypeEnumValues = []MicrosoftGraphManagedBrowserType{
	"notConfigured",
	"microsoftEdge",
}

func (v *MicrosoftGraphManagedBrowserType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MicrosoftGraphManagedBrowserType(value)
	for _, existing := range AllowedMicrosoftGraphManagedBrowserTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MicrosoftGraphManagedBrowserType", value)
}

// NewMicrosoftGraphManagedBrowserTypeFromValue returns a pointer to a valid MicrosoftGraphManagedBrowserType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMicrosoftGraphManagedBrowserTypeFromValue(v string) (*MicrosoftGraphManagedBrowserType, error) {
	ev := MicrosoftGraphManagedBrowserType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MicrosoftGraphManagedBrowserType: valid values are %v", v, AllowedMicrosoftGraphManagedBrowserTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MicrosoftGraphManagedBrowserType) IsValid() bool {
	for _, existing := range AllowedMicrosoftGraphManagedBrowserTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to microsoft.graph.managedBrowserType value
func (v MicrosoftGraphManagedBrowserType) Ptr() *MicrosoftGraphManagedBrowserType {
	return &v
}

type NullableMicrosoftGraphManagedBrowserType struct {
	value *MicrosoftGraphManagedBrowserType
	isSet bool
}

func (v NullableMicrosoftGraphManagedBrowserType) Get() *MicrosoftGraphManagedBrowserType {
	return v.value
}

func (v *NullableMicrosoftGraphManagedBrowserType) Set(val *MicrosoftGraphManagedBrowserType) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphManagedBrowserType) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphManagedBrowserType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphManagedBrowserType(val *MicrosoftGraphManagedBrowserType) *NullableMicrosoftGraphManagedBrowserType {
	return &NullableMicrosoftGraphManagedBrowserType{value: val, isSet: true}
}

func (v NullableMicrosoftGraphManagedBrowserType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphManagedBrowserType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

