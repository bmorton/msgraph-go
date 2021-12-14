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

// MicrosoftGraphCallRecordsWifiRadioType the model 'MicrosoftGraphCallRecordsWifiRadioType'
type MicrosoftGraphCallRecordsWifiRadioType string

// List of microsoft.graph.callRecords.wifiRadioType
const (
	UNKNOWN MicrosoftGraphCallRecordsWifiRadioType = "unknown"
	WIFI80211A MicrosoftGraphCallRecordsWifiRadioType = "wifi80211a"
	WIFI80211B MicrosoftGraphCallRecordsWifiRadioType = "wifi80211b"
	WIFI80211G MicrosoftGraphCallRecordsWifiRadioType = "wifi80211g"
	WIFI80211N MicrosoftGraphCallRecordsWifiRadioType = "wifi80211n"
	WIFI80211AC MicrosoftGraphCallRecordsWifiRadioType = "wifi80211ac"
	WIFI80211AX MicrosoftGraphCallRecordsWifiRadioType = "wifi80211ax"
	UNKNOWN_FUTURE_VALUE MicrosoftGraphCallRecordsWifiRadioType = "unknownFutureValue"
)

// All allowed values of MicrosoftGraphCallRecordsWifiRadioType enum
var AllowedMicrosoftGraphCallRecordsWifiRadioTypeEnumValues = []MicrosoftGraphCallRecordsWifiRadioType{
	"unknown",
	"wifi80211a",
	"wifi80211b",
	"wifi80211g",
	"wifi80211n",
	"wifi80211ac",
	"wifi80211ax",
	"unknownFutureValue",
}

func (v *MicrosoftGraphCallRecordsWifiRadioType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MicrosoftGraphCallRecordsWifiRadioType(value)
	for _, existing := range AllowedMicrosoftGraphCallRecordsWifiRadioTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MicrosoftGraphCallRecordsWifiRadioType", value)
}

// NewMicrosoftGraphCallRecordsWifiRadioTypeFromValue returns a pointer to a valid MicrosoftGraphCallRecordsWifiRadioType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMicrosoftGraphCallRecordsWifiRadioTypeFromValue(v string) (*MicrosoftGraphCallRecordsWifiRadioType, error) {
	ev := MicrosoftGraphCallRecordsWifiRadioType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MicrosoftGraphCallRecordsWifiRadioType: valid values are %v", v, AllowedMicrosoftGraphCallRecordsWifiRadioTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MicrosoftGraphCallRecordsWifiRadioType) IsValid() bool {
	for _, existing := range AllowedMicrosoftGraphCallRecordsWifiRadioTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to microsoft.graph.callRecords.wifiRadioType value
func (v MicrosoftGraphCallRecordsWifiRadioType) Ptr() *MicrosoftGraphCallRecordsWifiRadioType {
	return &v
}

type NullableMicrosoftGraphCallRecordsWifiRadioType struct {
	value *MicrosoftGraphCallRecordsWifiRadioType
	isSet bool
}

func (v NullableMicrosoftGraphCallRecordsWifiRadioType) Get() *MicrosoftGraphCallRecordsWifiRadioType {
	return v.value
}

func (v *NullableMicrosoftGraphCallRecordsWifiRadioType) Set(val *MicrosoftGraphCallRecordsWifiRadioType) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphCallRecordsWifiRadioType) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphCallRecordsWifiRadioType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphCallRecordsWifiRadioType(val *MicrosoftGraphCallRecordsWifiRadioType) *NullableMicrosoftGraphCallRecordsWifiRadioType {
	return &NullableMicrosoftGraphCallRecordsWifiRadioType{value: val, isSet: true}
}

func (v NullableMicrosoftGraphCallRecordsWifiRadioType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphCallRecordsWifiRadioType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
