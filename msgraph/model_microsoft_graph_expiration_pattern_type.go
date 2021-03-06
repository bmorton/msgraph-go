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

// MicrosoftGraphExpirationPatternType the model 'MicrosoftGraphExpirationPatternType'
type MicrosoftGraphExpirationPatternType string

// List of microsoft.graph.expirationPatternType
const (
	NOT_SPECIFIED MicrosoftGraphExpirationPatternType = "notSpecified"
	NO_EXPIRATION MicrosoftGraphExpirationPatternType = "noExpiration"
	AFTER_DATE_TIME MicrosoftGraphExpirationPatternType = "afterDateTime"
	AFTER_DURATION MicrosoftGraphExpirationPatternType = "afterDuration"
)

// All allowed values of MicrosoftGraphExpirationPatternType enum
var AllowedMicrosoftGraphExpirationPatternTypeEnumValues = []MicrosoftGraphExpirationPatternType{
	"notSpecified",
	"noExpiration",
	"afterDateTime",
	"afterDuration",
}

func (v *MicrosoftGraphExpirationPatternType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MicrosoftGraphExpirationPatternType(value)
	for _, existing := range AllowedMicrosoftGraphExpirationPatternTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MicrosoftGraphExpirationPatternType", value)
}

// NewMicrosoftGraphExpirationPatternTypeFromValue returns a pointer to a valid MicrosoftGraphExpirationPatternType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMicrosoftGraphExpirationPatternTypeFromValue(v string) (*MicrosoftGraphExpirationPatternType, error) {
	ev := MicrosoftGraphExpirationPatternType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MicrosoftGraphExpirationPatternType: valid values are %v", v, AllowedMicrosoftGraphExpirationPatternTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MicrosoftGraphExpirationPatternType) IsValid() bool {
	for _, existing := range AllowedMicrosoftGraphExpirationPatternTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to microsoft.graph.expirationPatternType value
func (v MicrosoftGraphExpirationPatternType) Ptr() *MicrosoftGraphExpirationPatternType {
	return &v
}

type NullableMicrosoftGraphExpirationPatternType struct {
	value *MicrosoftGraphExpirationPatternType
	isSet bool
}

func (v NullableMicrosoftGraphExpirationPatternType) Get() *MicrosoftGraphExpirationPatternType {
	return v.value
}

func (v *NullableMicrosoftGraphExpirationPatternType) Set(val *MicrosoftGraphExpirationPatternType) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphExpirationPatternType) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphExpirationPatternType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphExpirationPatternType(val *MicrosoftGraphExpirationPatternType) *NullableMicrosoftGraphExpirationPatternType {
	return &NullableMicrosoftGraphExpirationPatternType{value: val, isSet: true}
}

func (v NullableMicrosoftGraphExpirationPatternType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphExpirationPatternType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

