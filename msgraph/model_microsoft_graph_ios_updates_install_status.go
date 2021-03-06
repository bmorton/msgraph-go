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

// MicrosoftGraphIosUpdatesInstallStatus the model 'MicrosoftGraphIosUpdatesInstallStatus'
type MicrosoftGraphIosUpdatesInstallStatus string

// List of microsoft.graph.iosUpdatesInstallStatus
const (
	DEVICE_OS_HIGHER_THAN_DESIRED_OS_VERSION MicrosoftGraphIosUpdatesInstallStatus = "deviceOsHigherThanDesiredOsVersion"
	SHARED_DEVICE_USER_LOGGED_IN_ERROR MicrosoftGraphIosUpdatesInstallStatus = "sharedDeviceUserLoggedInError"
	NOT_SUPPORTED_OPERATION MicrosoftGraphIosUpdatesInstallStatus = "notSupportedOperation"
	INSTALL_FAILED MicrosoftGraphIosUpdatesInstallStatus = "installFailed"
	INSTALL_PHONE_CALL_IN_PROGRESS MicrosoftGraphIosUpdatesInstallStatus = "installPhoneCallInProgress"
	INSTALL_INSUFFICIENT_POWER MicrosoftGraphIosUpdatesInstallStatus = "installInsufficientPower"
	INSTALL_INSUFFICIENT_SPACE MicrosoftGraphIosUpdatesInstallStatus = "installInsufficientSpace"
	INSTALLING MicrosoftGraphIosUpdatesInstallStatus = "installing"
	DOWNLOAD_INSUFFICIENT_NETWORK MicrosoftGraphIosUpdatesInstallStatus = "downloadInsufficientNetwork"
	DOWNLOAD_INSUFFICIENT_POWER MicrosoftGraphIosUpdatesInstallStatus = "downloadInsufficientPower"
	DOWNLOAD_INSUFFICIENT_SPACE MicrosoftGraphIosUpdatesInstallStatus = "downloadInsufficientSpace"
	DOWNLOAD_REQUIRES_COMPUTER MicrosoftGraphIosUpdatesInstallStatus = "downloadRequiresComputer"
	DOWNLOAD_FAILED MicrosoftGraphIosUpdatesInstallStatus = "downloadFailed"
	DOWNLOADING MicrosoftGraphIosUpdatesInstallStatus = "downloading"
	SUCCESS MicrosoftGraphIosUpdatesInstallStatus = "success"
	AVAILABLE MicrosoftGraphIosUpdatesInstallStatus = "available"
	IDLE MicrosoftGraphIosUpdatesInstallStatus = "idle"
	UNKNOWN MicrosoftGraphIosUpdatesInstallStatus = "unknown"
)

// All allowed values of MicrosoftGraphIosUpdatesInstallStatus enum
var AllowedMicrosoftGraphIosUpdatesInstallStatusEnumValues = []MicrosoftGraphIosUpdatesInstallStatus{
	"deviceOsHigherThanDesiredOsVersion",
	"sharedDeviceUserLoggedInError",
	"notSupportedOperation",
	"installFailed",
	"installPhoneCallInProgress",
	"installInsufficientPower",
	"installInsufficientSpace",
	"installing",
	"downloadInsufficientNetwork",
	"downloadInsufficientPower",
	"downloadInsufficientSpace",
	"downloadRequiresComputer",
	"downloadFailed",
	"downloading",
	"success",
	"available",
	"idle",
	"unknown",
}

func (v *MicrosoftGraphIosUpdatesInstallStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MicrosoftGraphIosUpdatesInstallStatus(value)
	for _, existing := range AllowedMicrosoftGraphIosUpdatesInstallStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MicrosoftGraphIosUpdatesInstallStatus", value)
}

// NewMicrosoftGraphIosUpdatesInstallStatusFromValue returns a pointer to a valid MicrosoftGraphIosUpdatesInstallStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMicrosoftGraphIosUpdatesInstallStatusFromValue(v string) (*MicrosoftGraphIosUpdatesInstallStatus, error) {
	ev := MicrosoftGraphIosUpdatesInstallStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MicrosoftGraphIosUpdatesInstallStatus: valid values are %v", v, AllowedMicrosoftGraphIosUpdatesInstallStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MicrosoftGraphIosUpdatesInstallStatus) IsValid() bool {
	for _, existing := range AllowedMicrosoftGraphIosUpdatesInstallStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to microsoft.graph.iosUpdatesInstallStatus value
func (v MicrosoftGraphIosUpdatesInstallStatus) Ptr() *MicrosoftGraphIosUpdatesInstallStatus {
	return &v
}

type NullableMicrosoftGraphIosUpdatesInstallStatus struct {
	value *MicrosoftGraphIosUpdatesInstallStatus
	isSet bool
}

func (v NullableMicrosoftGraphIosUpdatesInstallStatus) Get() *MicrosoftGraphIosUpdatesInstallStatus {
	return v.value
}

func (v *NullableMicrosoftGraphIosUpdatesInstallStatus) Set(val *MicrosoftGraphIosUpdatesInstallStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphIosUpdatesInstallStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphIosUpdatesInstallStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphIosUpdatesInstallStatus(val *MicrosoftGraphIosUpdatesInstallStatus) *NullableMicrosoftGraphIosUpdatesInstallStatus {
	return &NullableMicrosoftGraphIosUpdatesInstallStatus{value: val, isSet: true}
}

func (v NullableMicrosoftGraphIosUpdatesInstallStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphIosUpdatesInstallStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

