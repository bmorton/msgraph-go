/*
Partial Graph API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package msgraph

import (
	"encoding/json"
)

// EducationAssignmentSettings struct for EducationAssignmentSettings
type EducationAssignmentSettings struct {
	// Indicates whether turn-in celebration animation will be shown. A value of true indicates that the animation will not be shown. Default value is false.
	SubmissionAnimationDisabled NullableBool `json:"submissionAnimationDisabled,omitempty"`
}

// NewEducationAssignmentSettings instantiates a new EducationAssignmentSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEducationAssignmentSettings() *EducationAssignmentSettings {
	this := EducationAssignmentSettings{}
	return &this
}

// NewEducationAssignmentSettingsWithDefaults instantiates a new EducationAssignmentSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEducationAssignmentSettingsWithDefaults() *EducationAssignmentSettings {
	this := EducationAssignmentSettings{}
	return &this
}

// GetSubmissionAnimationDisabled returns the SubmissionAnimationDisabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EducationAssignmentSettings) GetSubmissionAnimationDisabled() bool {
	if o == nil || o.SubmissionAnimationDisabled.Get() == nil {
		var ret bool
		return ret
	}
	return *o.SubmissionAnimationDisabled.Get()
}

// GetSubmissionAnimationDisabledOk returns a tuple with the SubmissionAnimationDisabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EducationAssignmentSettings) GetSubmissionAnimationDisabledOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SubmissionAnimationDisabled.Get(), o.SubmissionAnimationDisabled.IsSet()
}

// HasSubmissionAnimationDisabled returns a boolean if a field has been set.
func (o *EducationAssignmentSettings) HasSubmissionAnimationDisabled() bool {
	if o != nil && o.SubmissionAnimationDisabled.IsSet() {
		return true
	}

	return false
}

// SetSubmissionAnimationDisabled gets a reference to the given NullableBool and assigns it to the SubmissionAnimationDisabled field.
func (o *EducationAssignmentSettings) SetSubmissionAnimationDisabled(v bool) {
	o.SubmissionAnimationDisabled.Set(&v)
}
// SetSubmissionAnimationDisabledNil sets the value for SubmissionAnimationDisabled to be an explicit nil
func (o *EducationAssignmentSettings) SetSubmissionAnimationDisabledNil() {
	o.SubmissionAnimationDisabled.Set(nil)
}

// UnsetSubmissionAnimationDisabled ensures that no value is present for SubmissionAnimationDisabled, not even an explicit nil
func (o *EducationAssignmentSettings) UnsetSubmissionAnimationDisabled() {
	o.SubmissionAnimationDisabled.Unset()
}

func (o EducationAssignmentSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SubmissionAnimationDisabled.IsSet() {
		toSerialize["submissionAnimationDisabled"] = o.SubmissionAnimationDisabled.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableEducationAssignmentSettings struct {
	value *EducationAssignmentSettings
	isSet bool
}

func (v NullableEducationAssignmentSettings) Get() *EducationAssignmentSettings {
	return v.value
}

func (v *NullableEducationAssignmentSettings) Set(val *EducationAssignmentSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableEducationAssignmentSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableEducationAssignmentSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEducationAssignmentSettings(val *EducationAssignmentSettings) *NullableEducationAssignmentSettings {
	return &NullableEducationAssignmentSettings{value: val, isSet: true}
}

func (v NullableEducationAssignmentSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEducationAssignmentSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


