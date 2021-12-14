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

// MicrosoftGraphTeamSpecialization the model 'MicrosoftGraphTeamSpecialization'
type MicrosoftGraphTeamSpecialization string

// List of microsoft.graph.teamSpecialization
const (
	NONE MicrosoftGraphTeamSpecialization = "none"
	EDUCATION_STANDARD MicrosoftGraphTeamSpecialization = "educationStandard"
	EDUCATION_CLASS MicrosoftGraphTeamSpecialization = "educationClass"
	EDUCATION_PROFESSIONAL_LEARNING_COMMUNITY MicrosoftGraphTeamSpecialization = "educationProfessionalLearningCommunity"
	EDUCATION_STAFF MicrosoftGraphTeamSpecialization = "educationStaff"
	HEALTHCARE_STANDARD MicrosoftGraphTeamSpecialization = "healthcareStandard"
	HEALTHCARE_CARE_COORDINATION MicrosoftGraphTeamSpecialization = "healthcareCareCoordination"
	UNKNOWN_FUTURE_VALUE MicrosoftGraphTeamSpecialization = "unknownFutureValue"
)

// All allowed values of MicrosoftGraphTeamSpecialization enum
var AllowedMicrosoftGraphTeamSpecializationEnumValues = []MicrosoftGraphTeamSpecialization{
	"none",
	"educationStandard",
	"educationClass",
	"educationProfessionalLearningCommunity",
	"educationStaff",
	"healthcareStandard",
	"healthcareCareCoordination",
	"unknownFutureValue",
}

func (v *MicrosoftGraphTeamSpecialization) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MicrosoftGraphTeamSpecialization(value)
	for _, existing := range AllowedMicrosoftGraphTeamSpecializationEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MicrosoftGraphTeamSpecialization", value)
}

// NewMicrosoftGraphTeamSpecializationFromValue returns a pointer to a valid MicrosoftGraphTeamSpecialization
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMicrosoftGraphTeamSpecializationFromValue(v string) (*MicrosoftGraphTeamSpecialization, error) {
	ev := MicrosoftGraphTeamSpecialization(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MicrosoftGraphTeamSpecialization: valid values are %v", v, AllowedMicrosoftGraphTeamSpecializationEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MicrosoftGraphTeamSpecialization) IsValid() bool {
	for _, existing := range AllowedMicrosoftGraphTeamSpecializationEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to microsoft.graph.teamSpecialization value
func (v MicrosoftGraphTeamSpecialization) Ptr() *MicrosoftGraphTeamSpecialization {
	return &v
}

type NullableMicrosoftGraphTeamSpecialization struct {
	value *MicrosoftGraphTeamSpecialization
	isSet bool
}

func (v NullableMicrosoftGraphTeamSpecialization) Get() *MicrosoftGraphTeamSpecialization {
	return v.value
}

func (v *NullableMicrosoftGraphTeamSpecialization) Set(val *MicrosoftGraphTeamSpecialization) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphTeamSpecialization) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphTeamSpecialization) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphTeamSpecialization(val *MicrosoftGraphTeamSpecialization) *NullableMicrosoftGraphTeamSpecialization {
	return &NullableMicrosoftGraphTeamSpecialization{value: val, isSet: true}
}

func (v NullableMicrosoftGraphTeamSpecialization) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphTeamSpecialization) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

