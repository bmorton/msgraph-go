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

// MicrosoftGraphSettingSourceType the model 'MicrosoftGraphSettingSourceType'
type MicrosoftGraphSettingSourceType string

// List of microsoft.graph.settingSourceType
const (
	DEVICE_CONFIGURATION MicrosoftGraphSettingSourceType = "deviceConfiguration"
	DEVICE_INTENT MicrosoftGraphSettingSourceType = "deviceIntent"
)

// All allowed values of MicrosoftGraphSettingSourceType enum
var AllowedMicrosoftGraphSettingSourceTypeEnumValues = []MicrosoftGraphSettingSourceType{
	"deviceConfiguration",
	"deviceIntent",
}

func (v *MicrosoftGraphSettingSourceType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MicrosoftGraphSettingSourceType(value)
	for _, existing := range AllowedMicrosoftGraphSettingSourceTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MicrosoftGraphSettingSourceType", value)
}

// NewMicrosoftGraphSettingSourceTypeFromValue returns a pointer to a valid MicrosoftGraphSettingSourceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMicrosoftGraphSettingSourceTypeFromValue(v string) (*MicrosoftGraphSettingSourceType, error) {
	ev := MicrosoftGraphSettingSourceType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MicrosoftGraphSettingSourceType: valid values are %v", v, AllowedMicrosoftGraphSettingSourceTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MicrosoftGraphSettingSourceType) IsValid() bool {
	for _, existing := range AllowedMicrosoftGraphSettingSourceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to microsoft.graph.settingSourceType value
func (v MicrosoftGraphSettingSourceType) Ptr() *MicrosoftGraphSettingSourceType {
	return &v
}

type NullableMicrosoftGraphSettingSourceType struct {
	value *MicrosoftGraphSettingSourceType
	isSet bool
}

func (v NullableMicrosoftGraphSettingSourceType) Get() *MicrosoftGraphSettingSourceType {
	return v.value
}

func (v *NullableMicrosoftGraphSettingSourceType) Set(val *MicrosoftGraphSettingSourceType) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphSettingSourceType) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphSettingSourceType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphSettingSourceType(val *MicrosoftGraphSettingSourceType) *NullableMicrosoftGraphSettingSourceType {
	return &NullableMicrosoftGraphSettingSourceType{value: val, isSet: true}
}

func (v NullableMicrosoftGraphSettingSourceType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphSettingSourceType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

