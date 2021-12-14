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

// MicrosoftGraphNotificationTemplateBrandingOptions Branding Options for the Message Template. Branding is defined in the Intune Admin Console.
type MicrosoftGraphNotificationTemplateBrandingOptions string

// List of microsoft.graph.notificationTemplateBrandingOptions
const (
	NONE MicrosoftGraphNotificationTemplateBrandingOptions = "none"
	INCLUDE_COMPANY_LOGO MicrosoftGraphNotificationTemplateBrandingOptions = "includeCompanyLogo"
	INCLUDE_COMPANY_NAME MicrosoftGraphNotificationTemplateBrandingOptions = "includeCompanyName"
	INCLUDE_CONTACT_INFORMATION MicrosoftGraphNotificationTemplateBrandingOptions = "includeContactInformation"
)

// All allowed values of MicrosoftGraphNotificationTemplateBrandingOptions enum
var AllowedMicrosoftGraphNotificationTemplateBrandingOptionsEnumValues = []MicrosoftGraphNotificationTemplateBrandingOptions{
	"none",
	"includeCompanyLogo",
	"includeCompanyName",
	"includeContactInformation",
}

func (v *MicrosoftGraphNotificationTemplateBrandingOptions) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MicrosoftGraphNotificationTemplateBrandingOptions(value)
	for _, existing := range AllowedMicrosoftGraphNotificationTemplateBrandingOptionsEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MicrosoftGraphNotificationTemplateBrandingOptions", value)
}

// NewMicrosoftGraphNotificationTemplateBrandingOptionsFromValue returns a pointer to a valid MicrosoftGraphNotificationTemplateBrandingOptions
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMicrosoftGraphNotificationTemplateBrandingOptionsFromValue(v string) (*MicrosoftGraphNotificationTemplateBrandingOptions, error) {
	ev := MicrosoftGraphNotificationTemplateBrandingOptions(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MicrosoftGraphNotificationTemplateBrandingOptions: valid values are %v", v, AllowedMicrosoftGraphNotificationTemplateBrandingOptionsEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MicrosoftGraphNotificationTemplateBrandingOptions) IsValid() bool {
	for _, existing := range AllowedMicrosoftGraphNotificationTemplateBrandingOptionsEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to microsoft.graph.notificationTemplateBrandingOptions value
func (v MicrosoftGraphNotificationTemplateBrandingOptions) Ptr() *MicrosoftGraphNotificationTemplateBrandingOptions {
	return &v
}

type NullableMicrosoftGraphNotificationTemplateBrandingOptions struct {
	value *MicrosoftGraphNotificationTemplateBrandingOptions
	isSet bool
}

func (v NullableMicrosoftGraphNotificationTemplateBrandingOptions) Get() *MicrosoftGraphNotificationTemplateBrandingOptions {
	return v.value
}

func (v *NullableMicrosoftGraphNotificationTemplateBrandingOptions) Set(val *MicrosoftGraphNotificationTemplateBrandingOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphNotificationTemplateBrandingOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphNotificationTemplateBrandingOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphNotificationTemplateBrandingOptions(val *MicrosoftGraphNotificationTemplateBrandingOptions) *NullableMicrosoftGraphNotificationTemplateBrandingOptions {
	return &NullableMicrosoftGraphNotificationTemplateBrandingOptions{value: val, isSet: true}
}

func (v NullableMicrosoftGraphNotificationTemplateBrandingOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphNotificationTemplateBrandingOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

