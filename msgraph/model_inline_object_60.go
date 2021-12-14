/*
Partial Graph API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// InlineObject60 struct for InlineObject60
type InlineObject60 struct {
	Assignments *[]*AnyOfmicrosoftGraphTargetedManagedAppPolicyAssignment `json:"assignments,omitempty"`
}

// NewInlineObject60 instantiates a new InlineObject60 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject60() *InlineObject60 {
	this := InlineObject60{}
	return &this
}

// NewInlineObject60WithDefaults instantiates a new InlineObject60 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject60WithDefaults() *InlineObject60 {
	this := InlineObject60{}
	return &this
}

// GetAssignments returns the Assignments field value if set, zero value otherwise.
func (o *InlineObject60) GetAssignments() []*AnyOfmicrosoftGraphTargetedManagedAppPolicyAssignment {
	if o == nil || o.Assignments == nil {
		var ret []*AnyOfmicrosoftGraphTargetedManagedAppPolicyAssignment
		return ret
	}
	return *o.Assignments
}

// GetAssignmentsOk returns a tuple with the Assignments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject60) GetAssignmentsOk() (*[]*AnyOfmicrosoftGraphTargetedManagedAppPolicyAssignment, bool) {
	if o == nil || o.Assignments == nil {
		return nil, false
	}
	return o.Assignments, true
}

// HasAssignments returns a boolean if a field has been set.
func (o *InlineObject60) HasAssignments() bool {
	if o != nil && o.Assignments != nil {
		return true
	}

	return false
}

// SetAssignments gets a reference to the given []*AnyOfmicrosoftGraphTargetedManagedAppPolicyAssignment and assigns it to the Assignments field.
func (o *InlineObject60) SetAssignments(v []*AnyOfmicrosoftGraphTargetedManagedAppPolicyAssignment) {
	o.Assignments = &v
}

func (o InlineObject60) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Assignments != nil {
		toSerialize["assignments"] = o.Assignments
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject60 struct {
	value *InlineObject60
	isSet bool
}

func (v NullableInlineObject60) Get() *InlineObject60 {
	return v.value
}

func (v *NullableInlineObject60) Set(val *InlineObject60) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject60) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject60) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject60(val *InlineObject60) *NullableInlineObject60 {
	return &NullableInlineObject60{value: val, isSet: true}
}

func (v NullableInlineObject60) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject60) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

