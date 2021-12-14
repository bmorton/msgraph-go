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

// MicrosoftGraphFileHashType the model 'MicrosoftGraphFileHashType'
type MicrosoftGraphFileHashType string

// List of microsoft.graph.fileHashType
const (
	UNKNOWN MicrosoftGraphFileHashType = "unknown"
	SHA1 MicrosoftGraphFileHashType = "sha1"
	SHA256 MicrosoftGraphFileHashType = "sha256"
	MD5 MicrosoftGraphFileHashType = "md5"
	AUTHENTICODE_HASH256 MicrosoftGraphFileHashType = "authenticodeHash256"
	LS_HASH MicrosoftGraphFileHashType = "lsHash"
	CTPH MicrosoftGraphFileHashType = "ctph"
	UNKNOWN_FUTURE_VALUE MicrosoftGraphFileHashType = "unknownFutureValue"
)

// All allowed values of MicrosoftGraphFileHashType enum
var AllowedMicrosoftGraphFileHashTypeEnumValues = []MicrosoftGraphFileHashType{
	"unknown",
	"sha1",
	"sha256",
	"md5",
	"authenticodeHash256",
	"lsHash",
	"ctph",
	"unknownFutureValue",
}

func (v *MicrosoftGraphFileHashType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MicrosoftGraphFileHashType(value)
	for _, existing := range AllowedMicrosoftGraphFileHashTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MicrosoftGraphFileHashType", value)
}

// NewMicrosoftGraphFileHashTypeFromValue returns a pointer to a valid MicrosoftGraphFileHashType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMicrosoftGraphFileHashTypeFromValue(v string) (*MicrosoftGraphFileHashType, error) {
	ev := MicrosoftGraphFileHashType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MicrosoftGraphFileHashType: valid values are %v", v, AllowedMicrosoftGraphFileHashTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MicrosoftGraphFileHashType) IsValid() bool {
	for _, existing := range AllowedMicrosoftGraphFileHashTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to microsoft.graph.fileHashType value
func (v MicrosoftGraphFileHashType) Ptr() *MicrosoftGraphFileHashType {
	return &v
}

type NullableMicrosoftGraphFileHashType struct {
	value *MicrosoftGraphFileHashType
	isSet bool
}

func (v NullableMicrosoftGraphFileHashType) Get() *MicrosoftGraphFileHashType {
	return v.value
}

func (v *NullableMicrosoftGraphFileHashType) Set(val *MicrosoftGraphFileHashType) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphFileHashType) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphFileHashType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphFileHashType(val *MicrosoftGraphFileHashType) *NullableMicrosoftGraphFileHashType {
	return &NullableMicrosoftGraphFileHashType{value: val, isSet: true}
}

func (v NullableMicrosoftGraphFileHashType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphFileHashType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

