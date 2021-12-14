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

// MicrosoftGraphAccessPackageCatalogState the model 'MicrosoftGraphAccessPackageCatalogState'
type MicrosoftGraphAccessPackageCatalogState string

// List of microsoft.graph.accessPackageCatalogState
const (
	UNPUBLISHED MicrosoftGraphAccessPackageCatalogState = "unpublished"
	PUBLISHED MicrosoftGraphAccessPackageCatalogState = "published"
	UNKNOWN_FUTURE_VALUE MicrosoftGraphAccessPackageCatalogState = "unknownFutureValue"
)

// All allowed values of MicrosoftGraphAccessPackageCatalogState enum
var AllowedMicrosoftGraphAccessPackageCatalogStateEnumValues = []MicrosoftGraphAccessPackageCatalogState{
	"unpublished",
	"published",
	"unknownFutureValue",
}

func (v *MicrosoftGraphAccessPackageCatalogState) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MicrosoftGraphAccessPackageCatalogState(value)
	for _, existing := range AllowedMicrosoftGraphAccessPackageCatalogStateEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MicrosoftGraphAccessPackageCatalogState", value)
}

// NewMicrosoftGraphAccessPackageCatalogStateFromValue returns a pointer to a valid MicrosoftGraphAccessPackageCatalogState
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMicrosoftGraphAccessPackageCatalogStateFromValue(v string) (*MicrosoftGraphAccessPackageCatalogState, error) {
	ev := MicrosoftGraphAccessPackageCatalogState(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MicrosoftGraphAccessPackageCatalogState: valid values are %v", v, AllowedMicrosoftGraphAccessPackageCatalogStateEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MicrosoftGraphAccessPackageCatalogState) IsValid() bool {
	for _, existing := range AllowedMicrosoftGraphAccessPackageCatalogStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to microsoft.graph.accessPackageCatalogState value
func (v MicrosoftGraphAccessPackageCatalogState) Ptr() *MicrosoftGraphAccessPackageCatalogState {
	return &v
}

type NullableMicrosoftGraphAccessPackageCatalogState struct {
	value *MicrosoftGraphAccessPackageCatalogState
	isSet bool
}

func (v NullableMicrosoftGraphAccessPackageCatalogState) Get() *MicrosoftGraphAccessPackageCatalogState {
	return v.value
}

func (v *NullableMicrosoftGraphAccessPackageCatalogState) Set(val *MicrosoftGraphAccessPackageCatalogState) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphAccessPackageCatalogState) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphAccessPackageCatalogState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphAccessPackageCatalogState(val *MicrosoftGraphAccessPackageCatalogState) *NullableMicrosoftGraphAccessPackageCatalogState {
	return &NullableMicrosoftGraphAccessPackageCatalogState{value: val, isSet: true}
}

func (v NullableMicrosoftGraphAccessPackageCatalogState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphAccessPackageCatalogState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

