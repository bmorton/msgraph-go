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

// MicrosoftGraphWebsiteType the model 'MicrosoftGraphWebsiteType'
type MicrosoftGraphWebsiteType string

// List of microsoft.graph.websiteType
const (
	OTHER MicrosoftGraphWebsiteType = "other"
	HOME MicrosoftGraphWebsiteType = "home"
	WORK MicrosoftGraphWebsiteType = "work"
	BLOG MicrosoftGraphWebsiteType = "blog"
	PROFILE MicrosoftGraphWebsiteType = "profile"
)

// All allowed values of MicrosoftGraphWebsiteType enum
var AllowedMicrosoftGraphWebsiteTypeEnumValues = []MicrosoftGraphWebsiteType{
	"other",
	"home",
	"work",
	"blog",
	"profile",
}

func (v *MicrosoftGraphWebsiteType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MicrosoftGraphWebsiteType(value)
	for _, existing := range AllowedMicrosoftGraphWebsiteTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MicrosoftGraphWebsiteType", value)
}

// NewMicrosoftGraphWebsiteTypeFromValue returns a pointer to a valid MicrosoftGraphWebsiteType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMicrosoftGraphWebsiteTypeFromValue(v string) (*MicrosoftGraphWebsiteType, error) {
	ev := MicrosoftGraphWebsiteType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MicrosoftGraphWebsiteType: valid values are %v", v, AllowedMicrosoftGraphWebsiteTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MicrosoftGraphWebsiteType) IsValid() bool {
	for _, existing := range AllowedMicrosoftGraphWebsiteTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to microsoft.graph.websiteType value
func (v MicrosoftGraphWebsiteType) Ptr() *MicrosoftGraphWebsiteType {
	return &v
}

type NullableMicrosoftGraphWebsiteType struct {
	value *MicrosoftGraphWebsiteType
	isSet bool
}

func (v NullableMicrosoftGraphWebsiteType) Get() *MicrosoftGraphWebsiteType {
	return v.value
}

func (v *NullableMicrosoftGraphWebsiteType) Set(val *MicrosoftGraphWebsiteType) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphWebsiteType) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphWebsiteType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphWebsiteType(val *MicrosoftGraphWebsiteType) *NullableMicrosoftGraphWebsiteType {
	return &NullableMicrosoftGraphWebsiteType{value: val, isSet: true}
}

func (v NullableMicrosoftGraphWebsiteType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphWebsiteType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

