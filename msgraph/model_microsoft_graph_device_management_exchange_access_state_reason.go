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

// MicrosoftGraphDeviceManagementExchangeAccessStateReason Device Exchange Access State Reason.
type MicrosoftGraphDeviceManagementExchangeAccessStateReason string

// List of microsoft.graph.deviceManagementExchangeAccessStateReason
const (
	NONE MicrosoftGraphDeviceManagementExchangeAccessStateReason = "none"
	UNKNOWN MicrosoftGraphDeviceManagementExchangeAccessStateReason = "unknown"
	EXCHANGE_GLOBAL_RULE MicrosoftGraphDeviceManagementExchangeAccessStateReason = "exchangeGlobalRule"
	EXCHANGE_INDIVIDUAL_RULE MicrosoftGraphDeviceManagementExchangeAccessStateReason = "exchangeIndividualRule"
	EXCHANGE_DEVICE_RULE MicrosoftGraphDeviceManagementExchangeAccessStateReason = "exchangeDeviceRule"
	EXCHANGE_UPGRADE MicrosoftGraphDeviceManagementExchangeAccessStateReason = "exchangeUpgrade"
	EXCHANGE_MAILBOX_POLICY MicrosoftGraphDeviceManagementExchangeAccessStateReason = "exchangeMailboxPolicy"
	OTHER MicrosoftGraphDeviceManagementExchangeAccessStateReason = "other"
	COMPLIANT MicrosoftGraphDeviceManagementExchangeAccessStateReason = "compliant"
	NOT_COMPLIANT MicrosoftGraphDeviceManagementExchangeAccessStateReason = "notCompliant"
	NOT_ENROLLED MicrosoftGraphDeviceManagementExchangeAccessStateReason = "notEnrolled"
	UNKNOWN_LOCATION MicrosoftGraphDeviceManagementExchangeAccessStateReason = "unknownLocation"
	MFA_REQUIRED MicrosoftGraphDeviceManagementExchangeAccessStateReason = "mfaRequired"
	AZURE_AD_BLOCK_DUE_TO_ACCESS_POLICY MicrosoftGraphDeviceManagementExchangeAccessStateReason = "azureADBlockDueToAccessPolicy"
	COMPROMISED_PASSWORD MicrosoftGraphDeviceManagementExchangeAccessStateReason = "compromisedPassword"
	DEVICE_NOT_KNOWN_WITH_MANAGED_APP MicrosoftGraphDeviceManagementExchangeAccessStateReason = "deviceNotKnownWithManagedApp"
)

// All allowed values of MicrosoftGraphDeviceManagementExchangeAccessStateReason enum
var AllowedMicrosoftGraphDeviceManagementExchangeAccessStateReasonEnumValues = []MicrosoftGraphDeviceManagementExchangeAccessStateReason{
	"none",
	"unknown",
	"exchangeGlobalRule",
	"exchangeIndividualRule",
	"exchangeDeviceRule",
	"exchangeUpgrade",
	"exchangeMailboxPolicy",
	"other",
	"compliant",
	"notCompliant",
	"notEnrolled",
	"unknownLocation",
	"mfaRequired",
	"azureADBlockDueToAccessPolicy",
	"compromisedPassword",
	"deviceNotKnownWithManagedApp",
}

func (v *MicrosoftGraphDeviceManagementExchangeAccessStateReason) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MicrosoftGraphDeviceManagementExchangeAccessStateReason(value)
	for _, existing := range AllowedMicrosoftGraphDeviceManagementExchangeAccessStateReasonEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MicrosoftGraphDeviceManagementExchangeAccessStateReason", value)
}

// NewMicrosoftGraphDeviceManagementExchangeAccessStateReasonFromValue returns a pointer to a valid MicrosoftGraphDeviceManagementExchangeAccessStateReason
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMicrosoftGraphDeviceManagementExchangeAccessStateReasonFromValue(v string) (*MicrosoftGraphDeviceManagementExchangeAccessStateReason, error) {
	ev := MicrosoftGraphDeviceManagementExchangeAccessStateReason(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MicrosoftGraphDeviceManagementExchangeAccessStateReason: valid values are %v", v, AllowedMicrosoftGraphDeviceManagementExchangeAccessStateReasonEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MicrosoftGraphDeviceManagementExchangeAccessStateReason) IsValid() bool {
	for _, existing := range AllowedMicrosoftGraphDeviceManagementExchangeAccessStateReasonEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to microsoft.graph.deviceManagementExchangeAccessStateReason value
func (v MicrosoftGraphDeviceManagementExchangeAccessStateReason) Ptr() *MicrosoftGraphDeviceManagementExchangeAccessStateReason {
	return &v
}

type NullableMicrosoftGraphDeviceManagementExchangeAccessStateReason struct {
	value *MicrosoftGraphDeviceManagementExchangeAccessStateReason
	isSet bool
}

func (v NullableMicrosoftGraphDeviceManagementExchangeAccessStateReason) Get() *MicrosoftGraphDeviceManagementExchangeAccessStateReason {
	return v.value
}

func (v *NullableMicrosoftGraphDeviceManagementExchangeAccessStateReason) Set(val *MicrosoftGraphDeviceManagementExchangeAccessStateReason) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphDeviceManagementExchangeAccessStateReason) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphDeviceManagementExchangeAccessStateReason) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphDeviceManagementExchangeAccessStateReason(val *MicrosoftGraphDeviceManagementExchangeAccessStateReason) *NullableMicrosoftGraphDeviceManagementExchangeAccessStateReason {
	return &NullableMicrosoftGraphDeviceManagementExchangeAccessStateReason{value: val, isSet: true}
}

func (v NullableMicrosoftGraphDeviceManagementExchangeAccessStateReason) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphDeviceManagementExchangeAccessStateReason) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
