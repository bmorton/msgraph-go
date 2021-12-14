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

// MicrosoftGraphExternalConnectorsAccessType the model 'MicrosoftGraphExternalConnectorsAccessType'
type MicrosoftGraphExternalConnectorsAccessType string

// List of microsoft.graph.externalConnectors.accessType
const (
	GRANT MicrosoftGraphExternalConnectorsAccessType = "grant"
	DENY MicrosoftGraphExternalConnectorsAccessType = "deny"
	UNKNOWN_FUTURE_VALUE MicrosoftGraphExternalConnectorsAccessType = "unknownFutureValue"
)

// All allowed values of MicrosoftGraphExternalConnectorsAccessType enum
var AllowedMicrosoftGraphExternalConnectorsAccessTypeEnumValues = []MicrosoftGraphExternalConnectorsAccessType{
	"grant",
	"deny",
	"unknownFutureValue",
}

func (v *MicrosoftGraphExternalConnectorsAccessType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MicrosoftGraphExternalConnectorsAccessType(value)
	for _, existing := range AllowedMicrosoftGraphExternalConnectorsAccessTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MicrosoftGraphExternalConnectorsAccessType", value)
}

// NewMicrosoftGraphExternalConnectorsAccessTypeFromValue returns a pointer to a valid MicrosoftGraphExternalConnectorsAccessType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMicrosoftGraphExternalConnectorsAccessTypeFromValue(v string) (*MicrosoftGraphExternalConnectorsAccessType, error) {
	ev := MicrosoftGraphExternalConnectorsAccessType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MicrosoftGraphExternalConnectorsAccessType: valid values are %v", v, AllowedMicrosoftGraphExternalConnectorsAccessTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MicrosoftGraphExternalConnectorsAccessType) IsValid() bool {
	for _, existing := range AllowedMicrosoftGraphExternalConnectorsAccessTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to microsoft.graph.externalConnectors.accessType value
func (v MicrosoftGraphExternalConnectorsAccessType) Ptr() *MicrosoftGraphExternalConnectorsAccessType {
	return &v
}

type NullableMicrosoftGraphExternalConnectorsAccessType struct {
	value *MicrosoftGraphExternalConnectorsAccessType
	isSet bool
}

func (v NullableMicrosoftGraphExternalConnectorsAccessType) Get() *MicrosoftGraphExternalConnectorsAccessType {
	return v.value
}

func (v *NullableMicrosoftGraphExternalConnectorsAccessType) Set(val *MicrosoftGraphExternalConnectorsAccessType) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphExternalConnectorsAccessType) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphExternalConnectorsAccessType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphExternalConnectorsAccessType(val *MicrosoftGraphExternalConnectorsAccessType) *NullableMicrosoftGraphExternalConnectorsAccessType {
	return &NullableMicrosoftGraphExternalConnectorsAccessType{value: val, isSet: true}
}

func (v NullableMicrosoftGraphExternalConnectorsAccessType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphExternalConnectorsAccessType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

