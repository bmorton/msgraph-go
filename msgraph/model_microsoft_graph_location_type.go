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

// MicrosoftGraphLocationType the model 'MicrosoftGraphLocationType'
type MicrosoftGraphLocationType string

// List of microsoft.graph.locationType
const (
	DEFAULT MicrosoftGraphLocationType = "default"
	CONFERENCE_ROOM MicrosoftGraphLocationType = "conferenceRoom"
	HOME_ADDRESS MicrosoftGraphLocationType = "homeAddress"
	BUSINESS_ADDRESS MicrosoftGraphLocationType = "businessAddress"
	GEO_COORDINATES MicrosoftGraphLocationType = "geoCoordinates"
	STREET_ADDRESS MicrosoftGraphLocationType = "streetAddress"
	HOTEL MicrosoftGraphLocationType = "hotel"
	RESTAURANT MicrosoftGraphLocationType = "restaurant"
	LOCAL_BUSINESS MicrosoftGraphLocationType = "localBusiness"
	POSTAL_ADDRESS MicrosoftGraphLocationType = "postalAddress"
)

// All allowed values of MicrosoftGraphLocationType enum
var AllowedMicrosoftGraphLocationTypeEnumValues = []MicrosoftGraphLocationType{
	"default",
	"conferenceRoom",
	"homeAddress",
	"businessAddress",
	"geoCoordinates",
	"streetAddress",
	"hotel",
	"restaurant",
	"localBusiness",
	"postalAddress",
}

func (v *MicrosoftGraphLocationType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MicrosoftGraphLocationType(value)
	for _, existing := range AllowedMicrosoftGraphLocationTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MicrosoftGraphLocationType", value)
}

// NewMicrosoftGraphLocationTypeFromValue returns a pointer to a valid MicrosoftGraphLocationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMicrosoftGraphLocationTypeFromValue(v string) (*MicrosoftGraphLocationType, error) {
	ev := MicrosoftGraphLocationType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MicrosoftGraphLocationType: valid values are %v", v, AllowedMicrosoftGraphLocationTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MicrosoftGraphLocationType) IsValid() bool {
	for _, existing := range AllowedMicrosoftGraphLocationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to microsoft.graph.locationType value
func (v MicrosoftGraphLocationType) Ptr() *MicrosoftGraphLocationType {
	return &v
}

type NullableMicrosoftGraphLocationType struct {
	value *MicrosoftGraphLocationType
	isSet bool
}

func (v NullableMicrosoftGraphLocationType) Get() *MicrosoftGraphLocationType {
	return v.value
}

func (v *NullableMicrosoftGraphLocationType) Set(val *MicrosoftGraphLocationType) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphLocationType) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphLocationType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphLocationType(val *MicrosoftGraphLocationType) *NullableMicrosoftGraphLocationType {
	return &NullableMicrosoftGraphLocationType{value: val, isSet: true}
}

func (v NullableMicrosoftGraphLocationType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphLocationType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

