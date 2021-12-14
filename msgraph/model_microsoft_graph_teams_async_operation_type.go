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

// MicrosoftGraphTeamsAsyncOperationType the model 'MicrosoftGraphTeamsAsyncOperationType'
type MicrosoftGraphTeamsAsyncOperationType string

// List of microsoft.graph.teamsAsyncOperationType
const (
	INVALID MicrosoftGraphTeamsAsyncOperationType = "invalid"
	CLONE_TEAM MicrosoftGraphTeamsAsyncOperationType = "cloneTeam"
	ARCHIVE_TEAM MicrosoftGraphTeamsAsyncOperationType = "archiveTeam"
	UNARCHIVE_TEAM MicrosoftGraphTeamsAsyncOperationType = "unarchiveTeam"
	CREATE_TEAM MicrosoftGraphTeamsAsyncOperationType = "createTeam"
	UNKNOWN_FUTURE_VALUE MicrosoftGraphTeamsAsyncOperationType = "unknownFutureValue"
)

// All allowed values of MicrosoftGraphTeamsAsyncOperationType enum
var AllowedMicrosoftGraphTeamsAsyncOperationTypeEnumValues = []MicrosoftGraphTeamsAsyncOperationType{
	"invalid",
	"cloneTeam",
	"archiveTeam",
	"unarchiveTeam",
	"createTeam",
	"unknownFutureValue",
}

func (v *MicrosoftGraphTeamsAsyncOperationType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MicrosoftGraphTeamsAsyncOperationType(value)
	for _, existing := range AllowedMicrosoftGraphTeamsAsyncOperationTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MicrosoftGraphTeamsAsyncOperationType", value)
}

// NewMicrosoftGraphTeamsAsyncOperationTypeFromValue returns a pointer to a valid MicrosoftGraphTeamsAsyncOperationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMicrosoftGraphTeamsAsyncOperationTypeFromValue(v string) (*MicrosoftGraphTeamsAsyncOperationType, error) {
	ev := MicrosoftGraphTeamsAsyncOperationType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MicrosoftGraphTeamsAsyncOperationType: valid values are %v", v, AllowedMicrosoftGraphTeamsAsyncOperationTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MicrosoftGraphTeamsAsyncOperationType) IsValid() bool {
	for _, existing := range AllowedMicrosoftGraphTeamsAsyncOperationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to microsoft.graph.teamsAsyncOperationType value
func (v MicrosoftGraphTeamsAsyncOperationType) Ptr() *MicrosoftGraphTeamsAsyncOperationType {
	return &v
}

type NullableMicrosoftGraphTeamsAsyncOperationType struct {
	value *MicrosoftGraphTeamsAsyncOperationType
	isSet bool
}

func (v NullableMicrosoftGraphTeamsAsyncOperationType) Get() *MicrosoftGraphTeamsAsyncOperationType {
	return v.value
}

func (v *NullableMicrosoftGraphTeamsAsyncOperationType) Set(val *MicrosoftGraphTeamsAsyncOperationType) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphTeamsAsyncOperationType) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphTeamsAsyncOperationType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphTeamsAsyncOperationType(val *MicrosoftGraphTeamsAsyncOperationType) *NullableMicrosoftGraphTeamsAsyncOperationType {
	return &NullableMicrosoftGraphTeamsAsyncOperationType{value: val, isSet: true}
}

func (v NullableMicrosoftGraphTeamsAsyncOperationType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphTeamsAsyncOperationType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

