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

// MicrosoftGraphManagedAppFlaggedReason The reason for which a user has been flagged
type MicrosoftGraphManagedAppFlaggedReason string

// List of microsoft.graph.managedAppFlaggedReason
const (
	NONE MicrosoftGraphManagedAppFlaggedReason = "none"
	ROOTED_DEVICE MicrosoftGraphManagedAppFlaggedReason = "rootedDevice"
)

// All allowed values of MicrosoftGraphManagedAppFlaggedReason enum
var AllowedMicrosoftGraphManagedAppFlaggedReasonEnumValues = []MicrosoftGraphManagedAppFlaggedReason{
	"none",
	"rootedDevice",
}

func (v *MicrosoftGraphManagedAppFlaggedReason) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MicrosoftGraphManagedAppFlaggedReason(value)
	for _, existing := range AllowedMicrosoftGraphManagedAppFlaggedReasonEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MicrosoftGraphManagedAppFlaggedReason", value)
}

// NewMicrosoftGraphManagedAppFlaggedReasonFromValue returns a pointer to a valid MicrosoftGraphManagedAppFlaggedReason
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMicrosoftGraphManagedAppFlaggedReasonFromValue(v string) (*MicrosoftGraphManagedAppFlaggedReason, error) {
	ev := MicrosoftGraphManagedAppFlaggedReason(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MicrosoftGraphManagedAppFlaggedReason: valid values are %v", v, AllowedMicrosoftGraphManagedAppFlaggedReasonEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MicrosoftGraphManagedAppFlaggedReason) IsValid() bool {
	for _, existing := range AllowedMicrosoftGraphManagedAppFlaggedReasonEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to microsoft.graph.managedAppFlaggedReason value
func (v MicrosoftGraphManagedAppFlaggedReason) Ptr() *MicrosoftGraphManagedAppFlaggedReason {
	return &v
}

type NullableMicrosoftGraphManagedAppFlaggedReason struct {
	value *MicrosoftGraphManagedAppFlaggedReason
	isSet bool
}

func (v NullableMicrosoftGraphManagedAppFlaggedReason) Get() *MicrosoftGraphManagedAppFlaggedReason {
	return v.value
}

func (v *NullableMicrosoftGraphManagedAppFlaggedReason) Set(val *MicrosoftGraphManagedAppFlaggedReason) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphManagedAppFlaggedReason) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphManagedAppFlaggedReason) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphManagedAppFlaggedReason(val *MicrosoftGraphManagedAppFlaggedReason) *NullableMicrosoftGraphManagedAppFlaggedReason {
	return &NullableMicrosoftGraphManagedAppFlaggedReason{value: val, isSet: true}
}

func (v NullableMicrosoftGraphManagedAppFlaggedReason) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphManagedAppFlaggedReason) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
