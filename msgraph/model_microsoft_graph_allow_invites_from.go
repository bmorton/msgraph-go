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

// MicrosoftGraphAllowInvitesFrom the model 'MicrosoftGraphAllowInvitesFrom'
type MicrosoftGraphAllowInvitesFrom string

// List of microsoft.graph.allowInvitesFrom
const (
	NONE MicrosoftGraphAllowInvitesFrom = "none"
	ADMINS_AND_GUEST_INVITERS MicrosoftGraphAllowInvitesFrom = "adminsAndGuestInviters"
	ADMINS_GUEST_INVITERS_AND_ALL_MEMBERS MicrosoftGraphAllowInvitesFrom = "adminsGuestInvitersAndAllMembers"
	EVERYONE MicrosoftGraphAllowInvitesFrom = "everyone"
	UNKNOWN_FUTURE_VALUE MicrosoftGraphAllowInvitesFrom = "unknownFutureValue"
)

// All allowed values of MicrosoftGraphAllowInvitesFrom enum
var AllowedMicrosoftGraphAllowInvitesFromEnumValues = []MicrosoftGraphAllowInvitesFrom{
	"none",
	"adminsAndGuestInviters",
	"adminsGuestInvitersAndAllMembers",
	"everyone",
	"unknownFutureValue",
}

func (v *MicrosoftGraphAllowInvitesFrom) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MicrosoftGraphAllowInvitesFrom(value)
	for _, existing := range AllowedMicrosoftGraphAllowInvitesFromEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MicrosoftGraphAllowInvitesFrom", value)
}

// NewMicrosoftGraphAllowInvitesFromFromValue returns a pointer to a valid MicrosoftGraphAllowInvitesFrom
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMicrosoftGraphAllowInvitesFromFromValue(v string) (*MicrosoftGraphAllowInvitesFrom, error) {
	ev := MicrosoftGraphAllowInvitesFrom(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MicrosoftGraphAllowInvitesFrom: valid values are %v", v, AllowedMicrosoftGraphAllowInvitesFromEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MicrosoftGraphAllowInvitesFrom) IsValid() bool {
	for _, existing := range AllowedMicrosoftGraphAllowInvitesFromEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to microsoft.graph.allowInvitesFrom value
func (v MicrosoftGraphAllowInvitesFrom) Ptr() *MicrosoftGraphAllowInvitesFrom {
	return &v
}

type NullableMicrosoftGraphAllowInvitesFrom struct {
	value *MicrosoftGraphAllowInvitesFrom
	isSet bool
}

func (v NullableMicrosoftGraphAllowInvitesFrom) Get() *MicrosoftGraphAllowInvitesFrom {
	return v.value
}

func (v *NullableMicrosoftGraphAllowInvitesFrom) Set(val *MicrosoftGraphAllowInvitesFrom) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphAllowInvitesFrom) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphAllowInvitesFrom) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphAllowInvitesFrom(val *MicrosoftGraphAllowInvitesFrom) *NullableMicrosoftGraphAllowInvitesFrom {
	return &NullableMicrosoftGraphAllowInvitesFrom{value: val, isSet: true}
}

func (v NullableMicrosoftGraphAllowInvitesFrom) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphAllowInvitesFrom) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

