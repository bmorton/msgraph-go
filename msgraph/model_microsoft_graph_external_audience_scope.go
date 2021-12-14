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

// MicrosoftGraphExternalAudienceScope the model 'MicrosoftGraphExternalAudienceScope'
type MicrosoftGraphExternalAudienceScope string

// List of microsoft.graph.externalAudienceScope
const (
	NONE MicrosoftGraphExternalAudienceScope = "none"
	CONTACTS_ONLY MicrosoftGraphExternalAudienceScope = "contactsOnly"
	ALL MicrosoftGraphExternalAudienceScope = "all"
)

// All allowed values of MicrosoftGraphExternalAudienceScope enum
var AllowedMicrosoftGraphExternalAudienceScopeEnumValues = []MicrosoftGraphExternalAudienceScope{
	"none",
	"contactsOnly",
	"all",
}

func (v *MicrosoftGraphExternalAudienceScope) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MicrosoftGraphExternalAudienceScope(value)
	for _, existing := range AllowedMicrosoftGraphExternalAudienceScopeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MicrosoftGraphExternalAudienceScope", value)
}

// NewMicrosoftGraphExternalAudienceScopeFromValue returns a pointer to a valid MicrosoftGraphExternalAudienceScope
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMicrosoftGraphExternalAudienceScopeFromValue(v string) (*MicrosoftGraphExternalAudienceScope, error) {
	ev := MicrosoftGraphExternalAudienceScope(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MicrosoftGraphExternalAudienceScope: valid values are %v", v, AllowedMicrosoftGraphExternalAudienceScopeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MicrosoftGraphExternalAudienceScope) IsValid() bool {
	for _, existing := range AllowedMicrosoftGraphExternalAudienceScopeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to microsoft.graph.externalAudienceScope value
func (v MicrosoftGraphExternalAudienceScope) Ptr() *MicrosoftGraphExternalAudienceScope {
	return &v
}

type NullableMicrosoftGraphExternalAudienceScope struct {
	value *MicrosoftGraphExternalAudienceScope
	isSet bool
}

func (v NullableMicrosoftGraphExternalAudienceScope) Get() *MicrosoftGraphExternalAudienceScope {
	return v.value
}

func (v *NullableMicrosoftGraphExternalAudienceScope) Set(val *MicrosoftGraphExternalAudienceScope) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphExternalAudienceScope) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphExternalAudienceScope) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphExternalAudienceScope(val *MicrosoftGraphExternalAudienceScope) *NullableMicrosoftGraphExternalAudienceScope {
	return &NullableMicrosoftGraphExternalAudienceScope{value: val, isSet: true}
}

func (v NullableMicrosoftGraphExternalAudienceScope) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphExternalAudienceScope) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

