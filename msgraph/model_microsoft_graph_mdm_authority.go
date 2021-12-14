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

// MicrosoftGraphMdmAuthority Mobile device management authority.
type MicrosoftGraphMdmAuthority string

// List of microsoft.graph.mdmAuthority
const (
	UNKNOWN MicrosoftGraphMdmAuthority = "unknown"
	INTUNE MicrosoftGraphMdmAuthority = "intune"
	SCCM MicrosoftGraphMdmAuthority = "sccm"
	OFFICE365 MicrosoftGraphMdmAuthority = "office365"
)

// All allowed values of MicrosoftGraphMdmAuthority enum
var AllowedMicrosoftGraphMdmAuthorityEnumValues = []MicrosoftGraphMdmAuthority{
	"unknown",
	"intune",
	"sccm",
	"office365",
}

func (v *MicrosoftGraphMdmAuthority) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MicrosoftGraphMdmAuthority(value)
	for _, existing := range AllowedMicrosoftGraphMdmAuthorityEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MicrosoftGraphMdmAuthority", value)
}

// NewMicrosoftGraphMdmAuthorityFromValue returns a pointer to a valid MicrosoftGraphMdmAuthority
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMicrosoftGraphMdmAuthorityFromValue(v string) (*MicrosoftGraphMdmAuthority, error) {
	ev := MicrosoftGraphMdmAuthority(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MicrosoftGraphMdmAuthority: valid values are %v", v, AllowedMicrosoftGraphMdmAuthorityEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MicrosoftGraphMdmAuthority) IsValid() bool {
	for _, existing := range AllowedMicrosoftGraphMdmAuthorityEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to microsoft.graph.mdmAuthority value
func (v MicrosoftGraphMdmAuthority) Ptr() *MicrosoftGraphMdmAuthority {
	return &v
}

type NullableMicrosoftGraphMdmAuthority struct {
	value *MicrosoftGraphMdmAuthority
	isSet bool
}

func (v NullableMicrosoftGraphMdmAuthority) Get() *MicrosoftGraphMdmAuthority {
	return v.value
}

func (v *NullableMicrosoftGraphMdmAuthority) Set(val *MicrosoftGraphMdmAuthority) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphMdmAuthority) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphMdmAuthority) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphMdmAuthority(val *MicrosoftGraphMdmAuthority) *NullableMicrosoftGraphMdmAuthority {
	return &NullableMicrosoftGraphMdmAuthority{value: val, isSet: true}
}

func (v NullableMicrosoftGraphMdmAuthority) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphMdmAuthority) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

