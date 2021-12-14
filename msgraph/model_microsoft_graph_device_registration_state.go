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

// MicrosoftGraphDeviceRegistrationState Device registration status.
type MicrosoftGraphDeviceRegistrationState string

// List of microsoft.graph.deviceRegistrationState
const (
	NOT_REGISTERED MicrosoftGraphDeviceRegistrationState = "notRegistered"
	REGISTERED MicrosoftGraphDeviceRegistrationState = "registered"
	REVOKED MicrosoftGraphDeviceRegistrationState = "revoked"
	KEY_CONFLICT MicrosoftGraphDeviceRegistrationState = "keyConflict"
	APPROVAL_PENDING MicrosoftGraphDeviceRegistrationState = "approvalPending"
	CERTIFICATE_RESET MicrosoftGraphDeviceRegistrationState = "certificateReset"
	NOT_REGISTERED_PENDING_ENROLLMENT MicrosoftGraphDeviceRegistrationState = "notRegisteredPendingEnrollment"
	UNKNOWN MicrosoftGraphDeviceRegistrationState = "unknown"
)

// All allowed values of MicrosoftGraphDeviceRegistrationState enum
var AllowedMicrosoftGraphDeviceRegistrationStateEnumValues = []MicrosoftGraphDeviceRegistrationState{
	"notRegistered",
	"registered",
	"revoked",
	"keyConflict",
	"approvalPending",
	"certificateReset",
	"notRegisteredPendingEnrollment",
	"unknown",
}

func (v *MicrosoftGraphDeviceRegistrationState) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MicrosoftGraphDeviceRegistrationState(value)
	for _, existing := range AllowedMicrosoftGraphDeviceRegistrationStateEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MicrosoftGraphDeviceRegistrationState", value)
}

// NewMicrosoftGraphDeviceRegistrationStateFromValue returns a pointer to a valid MicrosoftGraphDeviceRegistrationState
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMicrosoftGraphDeviceRegistrationStateFromValue(v string) (*MicrosoftGraphDeviceRegistrationState, error) {
	ev := MicrosoftGraphDeviceRegistrationState(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MicrosoftGraphDeviceRegistrationState: valid values are %v", v, AllowedMicrosoftGraphDeviceRegistrationStateEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MicrosoftGraphDeviceRegistrationState) IsValid() bool {
	for _, existing := range AllowedMicrosoftGraphDeviceRegistrationStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to microsoft.graph.deviceRegistrationState value
func (v MicrosoftGraphDeviceRegistrationState) Ptr() *MicrosoftGraphDeviceRegistrationState {
	return &v
}

type NullableMicrosoftGraphDeviceRegistrationState struct {
	value *MicrosoftGraphDeviceRegistrationState
	isSet bool
}

func (v NullableMicrosoftGraphDeviceRegistrationState) Get() *MicrosoftGraphDeviceRegistrationState {
	return v.value
}

func (v *NullableMicrosoftGraphDeviceRegistrationState) Set(val *MicrosoftGraphDeviceRegistrationState) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphDeviceRegistrationState) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphDeviceRegistrationState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphDeviceRegistrationState(val *MicrosoftGraphDeviceRegistrationState) *NullableMicrosoftGraphDeviceRegistrationState {
	return &NullableMicrosoftGraphDeviceRegistrationState{value: val, isSet: true}
}

func (v NullableMicrosoftGraphDeviceRegistrationState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphDeviceRegistrationState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

