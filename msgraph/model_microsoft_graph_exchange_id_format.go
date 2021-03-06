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

// MicrosoftGraphExchangeIdFormat the model 'MicrosoftGraphExchangeIdFormat'
type MicrosoftGraphExchangeIdFormat string

// List of microsoft.graph.exchangeIdFormat
const (
	ENTRY_ID MicrosoftGraphExchangeIdFormat = "entryId"
	EWS_ID MicrosoftGraphExchangeIdFormat = "ewsId"
	IMMUTABLE_ENTRY_ID MicrosoftGraphExchangeIdFormat = "immutableEntryId"
	REST_ID MicrosoftGraphExchangeIdFormat = "restId"
	REST_IMMUTABLE_ENTRY_ID MicrosoftGraphExchangeIdFormat = "restImmutableEntryId"
)

// All allowed values of MicrosoftGraphExchangeIdFormat enum
var AllowedMicrosoftGraphExchangeIdFormatEnumValues = []MicrosoftGraphExchangeIdFormat{
	"entryId",
	"ewsId",
	"immutableEntryId",
	"restId",
	"restImmutableEntryId",
}

func (v *MicrosoftGraphExchangeIdFormat) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MicrosoftGraphExchangeIdFormat(value)
	for _, existing := range AllowedMicrosoftGraphExchangeIdFormatEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MicrosoftGraphExchangeIdFormat", value)
}

// NewMicrosoftGraphExchangeIdFormatFromValue returns a pointer to a valid MicrosoftGraphExchangeIdFormat
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMicrosoftGraphExchangeIdFormatFromValue(v string) (*MicrosoftGraphExchangeIdFormat, error) {
	ev := MicrosoftGraphExchangeIdFormat(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MicrosoftGraphExchangeIdFormat: valid values are %v", v, AllowedMicrosoftGraphExchangeIdFormatEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MicrosoftGraphExchangeIdFormat) IsValid() bool {
	for _, existing := range AllowedMicrosoftGraphExchangeIdFormatEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to microsoft.graph.exchangeIdFormat value
func (v MicrosoftGraphExchangeIdFormat) Ptr() *MicrosoftGraphExchangeIdFormat {
	return &v
}

type NullableMicrosoftGraphExchangeIdFormat struct {
	value *MicrosoftGraphExchangeIdFormat
	isSet bool
}

func (v NullableMicrosoftGraphExchangeIdFormat) Get() *MicrosoftGraphExchangeIdFormat {
	return v.value
}

func (v *NullableMicrosoftGraphExchangeIdFormat) Set(val *MicrosoftGraphExchangeIdFormat) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphExchangeIdFormat) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphExchangeIdFormat) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphExchangeIdFormat(val *MicrosoftGraphExchangeIdFormat) *NullableMicrosoftGraphExchangeIdFormat {
	return &NullableMicrosoftGraphExchangeIdFormat{value: val, isSet: true}
}

func (v NullableMicrosoftGraphExchangeIdFormat) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphExchangeIdFormat) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

