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

// MicrosoftGraphProvisioningAction the model 'MicrosoftGraphProvisioningAction'
type MicrosoftGraphProvisioningAction string

// List of microsoft.graph.provisioningAction
const (
	OTHER MicrosoftGraphProvisioningAction = "other"
	CREATE MicrosoftGraphProvisioningAction = "create"
	DELETE MicrosoftGraphProvisioningAction = "delete"
	DISABLE MicrosoftGraphProvisioningAction = "disable"
	UPDATE MicrosoftGraphProvisioningAction = "update"
	STAGED_DELETE MicrosoftGraphProvisioningAction = "stagedDelete"
	UNKNOWN_FUTURE_VALUE MicrosoftGraphProvisioningAction = "unknownFutureValue"
)

// All allowed values of MicrosoftGraphProvisioningAction enum
var AllowedMicrosoftGraphProvisioningActionEnumValues = []MicrosoftGraphProvisioningAction{
	"other",
	"create",
	"delete",
	"disable",
	"update",
	"stagedDelete",
	"unknownFutureValue",
}

func (v *MicrosoftGraphProvisioningAction) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MicrosoftGraphProvisioningAction(value)
	for _, existing := range AllowedMicrosoftGraphProvisioningActionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MicrosoftGraphProvisioningAction", value)
}

// NewMicrosoftGraphProvisioningActionFromValue returns a pointer to a valid MicrosoftGraphProvisioningAction
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMicrosoftGraphProvisioningActionFromValue(v string) (*MicrosoftGraphProvisioningAction, error) {
	ev := MicrosoftGraphProvisioningAction(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MicrosoftGraphProvisioningAction: valid values are %v", v, AllowedMicrosoftGraphProvisioningActionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MicrosoftGraphProvisioningAction) IsValid() bool {
	for _, existing := range AllowedMicrosoftGraphProvisioningActionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to microsoft.graph.provisioningAction value
func (v MicrosoftGraphProvisioningAction) Ptr() *MicrosoftGraphProvisioningAction {
	return &v
}

type NullableMicrosoftGraphProvisioningAction struct {
	value *MicrosoftGraphProvisioningAction
	isSet bool
}

func (v NullableMicrosoftGraphProvisioningAction) Get() *MicrosoftGraphProvisioningAction {
	return v.value
}

func (v *NullableMicrosoftGraphProvisioningAction) Set(val *MicrosoftGraphProvisioningAction) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphProvisioningAction) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphProvisioningAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphProvisioningAction(val *MicrosoftGraphProvisioningAction) *NullableMicrosoftGraphProvisioningAction {
	return &NullableMicrosoftGraphProvisioningAction{value: val, isSet: true}
}

func (v NullableMicrosoftGraphProvisioningAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphProvisioningAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

