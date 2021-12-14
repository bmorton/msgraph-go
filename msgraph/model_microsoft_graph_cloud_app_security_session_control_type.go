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

// MicrosoftGraphCloudAppSecuritySessionControlType the model 'MicrosoftGraphCloudAppSecuritySessionControlType'
type MicrosoftGraphCloudAppSecuritySessionControlType string

// List of microsoft.graph.cloudAppSecuritySessionControlType
const (
	MCAS_CONFIGURED MicrosoftGraphCloudAppSecuritySessionControlType = "mcasConfigured"
	MONITOR_ONLY MicrosoftGraphCloudAppSecuritySessionControlType = "monitorOnly"
	BLOCK_DOWNLOADS MicrosoftGraphCloudAppSecuritySessionControlType = "blockDownloads"
	UNKNOWN_FUTURE_VALUE MicrosoftGraphCloudAppSecuritySessionControlType = "unknownFutureValue"
)

// All allowed values of MicrosoftGraphCloudAppSecuritySessionControlType enum
var AllowedMicrosoftGraphCloudAppSecuritySessionControlTypeEnumValues = []MicrosoftGraphCloudAppSecuritySessionControlType{
	"mcasConfigured",
	"monitorOnly",
	"blockDownloads",
	"unknownFutureValue",
}

func (v *MicrosoftGraphCloudAppSecuritySessionControlType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MicrosoftGraphCloudAppSecuritySessionControlType(value)
	for _, existing := range AllowedMicrosoftGraphCloudAppSecuritySessionControlTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MicrosoftGraphCloudAppSecuritySessionControlType", value)
}

// NewMicrosoftGraphCloudAppSecuritySessionControlTypeFromValue returns a pointer to a valid MicrosoftGraphCloudAppSecuritySessionControlType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMicrosoftGraphCloudAppSecuritySessionControlTypeFromValue(v string) (*MicrosoftGraphCloudAppSecuritySessionControlType, error) {
	ev := MicrosoftGraphCloudAppSecuritySessionControlType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MicrosoftGraphCloudAppSecuritySessionControlType: valid values are %v", v, AllowedMicrosoftGraphCloudAppSecuritySessionControlTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MicrosoftGraphCloudAppSecuritySessionControlType) IsValid() bool {
	for _, existing := range AllowedMicrosoftGraphCloudAppSecuritySessionControlTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to microsoft.graph.cloudAppSecuritySessionControlType value
func (v MicrosoftGraphCloudAppSecuritySessionControlType) Ptr() *MicrosoftGraphCloudAppSecuritySessionControlType {
	return &v
}

type NullableMicrosoftGraphCloudAppSecuritySessionControlType struct {
	value *MicrosoftGraphCloudAppSecuritySessionControlType
	isSet bool
}

func (v NullableMicrosoftGraphCloudAppSecuritySessionControlType) Get() *MicrosoftGraphCloudAppSecuritySessionControlType {
	return v.value
}

func (v *NullableMicrosoftGraphCloudAppSecuritySessionControlType) Set(val *MicrosoftGraphCloudAppSecuritySessionControlType) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphCloudAppSecuritySessionControlType) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphCloudAppSecuritySessionControlType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphCloudAppSecuritySessionControlType(val *MicrosoftGraphCloudAppSecuritySessionControlType) *NullableMicrosoftGraphCloudAppSecuritySessionControlType {
	return &NullableMicrosoftGraphCloudAppSecuritySessionControlType{value: val, isSet: true}
}

func (v NullableMicrosoftGraphCloudAppSecuritySessionControlType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphCloudAppSecuritySessionControlType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
