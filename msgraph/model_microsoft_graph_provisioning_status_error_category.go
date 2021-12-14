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

// MicrosoftGraphProvisioningStatusErrorCategory the model 'MicrosoftGraphProvisioningStatusErrorCategory'
type MicrosoftGraphProvisioningStatusErrorCategory string

// List of microsoft.graph.provisioningStatusErrorCategory
const (
	FAILURE MicrosoftGraphProvisioningStatusErrorCategory = "failure"
	NON_SERVICE_FAILURE MicrosoftGraphProvisioningStatusErrorCategory = "nonServiceFailure"
	SUCCESS MicrosoftGraphProvisioningStatusErrorCategory = "success"
	UNKNOWN_FUTURE_VALUE MicrosoftGraphProvisioningStatusErrorCategory = "unknownFutureValue"
)

// All allowed values of MicrosoftGraphProvisioningStatusErrorCategory enum
var AllowedMicrosoftGraphProvisioningStatusErrorCategoryEnumValues = []MicrosoftGraphProvisioningStatusErrorCategory{
	"failure",
	"nonServiceFailure",
	"success",
	"unknownFutureValue",
}

func (v *MicrosoftGraphProvisioningStatusErrorCategory) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MicrosoftGraphProvisioningStatusErrorCategory(value)
	for _, existing := range AllowedMicrosoftGraphProvisioningStatusErrorCategoryEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MicrosoftGraphProvisioningStatusErrorCategory", value)
}

// NewMicrosoftGraphProvisioningStatusErrorCategoryFromValue returns a pointer to a valid MicrosoftGraphProvisioningStatusErrorCategory
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMicrosoftGraphProvisioningStatusErrorCategoryFromValue(v string) (*MicrosoftGraphProvisioningStatusErrorCategory, error) {
	ev := MicrosoftGraphProvisioningStatusErrorCategory(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MicrosoftGraphProvisioningStatusErrorCategory: valid values are %v", v, AllowedMicrosoftGraphProvisioningStatusErrorCategoryEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MicrosoftGraphProvisioningStatusErrorCategory) IsValid() bool {
	for _, existing := range AllowedMicrosoftGraphProvisioningStatusErrorCategoryEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to microsoft.graph.provisioningStatusErrorCategory value
func (v MicrosoftGraphProvisioningStatusErrorCategory) Ptr() *MicrosoftGraphProvisioningStatusErrorCategory {
	return &v
}

type NullableMicrosoftGraphProvisioningStatusErrorCategory struct {
	value *MicrosoftGraphProvisioningStatusErrorCategory
	isSet bool
}

func (v NullableMicrosoftGraphProvisioningStatusErrorCategory) Get() *MicrosoftGraphProvisioningStatusErrorCategory {
	return v.value
}

func (v *NullableMicrosoftGraphProvisioningStatusErrorCategory) Set(val *MicrosoftGraphProvisioningStatusErrorCategory) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphProvisioningStatusErrorCategory) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphProvisioningStatusErrorCategory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphProvisioningStatusErrorCategory(val *MicrosoftGraphProvisioningStatusErrorCategory) *NullableMicrosoftGraphProvisioningStatusErrorCategory {
	return &NullableMicrosoftGraphProvisioningStatusErrorCategory{value: val, isSet: true}
}

func (v NullableMicrosoftGraphProvisioningStatusErrorCategory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphProvisioningStatusErrorCategory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
