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

// MicrosoftGraphVolumeType the model 'MicrosoftGraphVolumeType'
type MicrosoftGraphVolumeType string

// List of microsoft.graph.volumeType
const (
	OPERATING_SYSTEM_VOLUME MicrosoftGraphVolumeType = "operatingSystemVolume"
	FIXED_DATA_VOLUME MicrosoftGraphVolumeType = "fixedDataVolume"
	REMOVABLE_DATA_VOLUME MicrosoftGraphVolumeType = "removableDataVolume"
	UNKNOWN_FUTURE_VALUE MicrosoftGraphVolumeType = "unknownFutureValue"
)

// All allowed values of MicrosoftGraphVolumeType enum
var AllowedMicrosoftGraphVolumeTypeEnumValues = []MicrosoftGraphVolumeType{
	"operatingSystemVolume",
	"fixedDataVolume",
	"removableDataVolume",
	"unknownFutureValue",
}

func (v *MicrosoftGraphVolumeType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MicrosoftGraphVolumeType(value)
	for _, existing := range AllowedMicrosoftGraphVolumeTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MicrosoftGraphVolumeType", value)
}

// NewMicrosoftGraphVolumeTypeFromValue returns a pointer to a valid MicrosoftGraphVolumeType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMicrosoftGraphVolumeTypeFromValue(v string) (*MicrosoftGraphVolumeType, error) {
	ev := MicrosoftGraphVolumeType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MicrosoftGraphVolumeType: valid values are %v", v, AllowedMicrosoftGraphVolumeTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MicrosoftGraphVolumeType) IsValid() bool {
	for _, existing := range AllowedMicrosoftGraphVolumeTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to microsoft.graph.volumeType value
func (v MicrosoftGraphVolumeType) Ptr() *MicrosoftGraphVolumeType {
	return &v
}

type NullableMicrosoftGraphVolumeType struct {
	value *MicrosoftGraphVolumeType
	isSet bool
}

func (v NullableMicrosoftGraphVolumeType) Get() *MicrosoftGraphVolumeType {
	return v.value
}

func (v *NullableMicrosoftGraphVolumeType) Set(val *MicrosoftGraphVolumeType) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphVolumeType) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphVolumeType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphVolumeType(val *MicrosoftGraphVolumeType) *NullableMicrosoftGraphVolumeType {
	return &NullableMicrosoftGraphVolumeType{value: val, isSet: true}
}

func (v NullableMicrosoftGraphVolumeType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphVolumeType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

