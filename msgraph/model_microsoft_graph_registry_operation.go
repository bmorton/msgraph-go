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

// MicrosoftGraphRegistryOperation the model 'MicrosoftGraphRegistryOperation'
type MicrosoftGraphRegistryOperation string

// List of microsoft.graph.registryOperation
const (
	UNKNOWN MicrosoftGraphRegistryOperation = "unknown"
	CREATE MicrosoftGraphRegistryOperation = "create"
	MODIFY MicrosoftGraphRegistryOperation = "modify"
	DELETE MicrosoftGraphRegistryOperation = "delete"
	UNKNOWN_FUTURE_VALUE MicrosoftGraphRegistryOperation = "unknownFutureValue"
)

// All allowed values of MicrosoftGraphRegistryOperation enum
var AllowedMicrosoftGraphRegistryOperationEnumValues = []MicrosoftGraphRegistryOperation{
	"unknown",
	"create",
	"modify",
	"delete",
	"unknownFutureValue",
}

func (v *MicrosoftGraphRegistryOperation) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MicrosoftGraphRegistryOperation(value)
	for _, existing := range AllowedMicrosoftGraphRegistryOperationEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MicrosoftGraphRegistryOperation", value)
}

// NewMicrosoftGraphRegistryOperationFromValue returns a pointer to a valid MicrosoftGraphRegistryOperation
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMicrosoftGraphRegistryOperationFromValue(v string) (*MicrosoftGraphRegistryOperation, error) {
	ev := MicrosoftGraphRegistryOperation(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MicrosoftGraphRegistryOperation: valid values are %v", v, AllowedMicrosoftGraphRegistryOperationEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MicrosoftGraphRegistryOperation) IsValid() bool {
	for _, existing := range AllowedMicrosoftGraphRegistryOperationEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to microsoft.graph.registryOperation value
func (v MicrosoftGraphRegistryOperation) Ptr() *MicrosoftGraphRegistryOperation {
	return &v
}

type NullableMicrosoftGraphRegistryOperation struct {
	value *MicrosoftGraphRegistryOperation
	isSet bool
}

func (v NullableMicrosoftGraphRegistryOperation) Get() *MicrosoftGraphRegistryOperation {
	return v.value
}

func (v *NullableMicrosoftGraphRegistryOperation) Set(val *MicrosoftGraphRegistryOperation) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphRegistryOperation) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphRegistryOperation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphRegistryOperation(val *MicrosoftGraphRegistryOperation) *NullableMicrosoftGraphRegistryOperation {
	return &NullableMicrosoftGraphRegistryOperation{value: val, isSet: true}
}

func (v NullableMicrosoftGraphRegistryOperation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphRegistryOperation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

