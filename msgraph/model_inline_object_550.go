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

// InlineObject550 struct for InlineObject550
type InlineObject550 struct {
	Assignments *[]*AnyOfmicrosoftGraphTargetedManagedAppPolicyAssignment `json:"assignments,omitempty"`
}

// NewInlineObject550 instantiates a new InlineObject550 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject550() *InlineObject550 {
	this := InlineObject550{}
	return &this
}

// NewInlineObject550WithDefaults instantiates a new InlineObject550 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject550WithDefaults() *InlineObject550 {
	this := InlineObject550{}
	return &this
}

// GetAssignments returns the Assignments field value if set, zero value otherwise.
func (o *InlineObject550) GetAssignments() []*AnyOfmicrosoftGraphTargetedManagedAppPolicyAssignment {
	if o == nil || o.Assignments == nil {
		var ret []*AnyOfmicrosoftGraphTargetedManagedAppPolicyAssignment
		return ret
	}
	return *o.Assignments
}

// GetAssignmentsOk returns a tuple with the Assignments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject550) GetAssignmentsOk() (*[]*AnyOfmicrosoftGraphTargetedManagedAppPolicyAssignment, bool) {
	if o == nil || o.Assignments == nil {
		return nil, false
	}
	return o.Assignments, true
}

// HasAssignments returns a boolean if a field has been set.
func (o *InlineObject550) HasAssignments() bool {
	if o != nil && o.Assignments != nil {
		return true
	}

	return false
}

// SetAssignments gets a reference to the given []*AnyOfmicrosoftGraphTargetedManagedAppPolicyAssignment and assigns it to the Assignments field.
func (o *InlineObject550) SetAssignments(v []*AnyOfmicrosoftGraphTargetedManagedAppPolicyAssignment) {
	o.Assignments = &v
}

func (o InlineObject550) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Assignments != nil {
		toSerialize["assignments"] = o.Assignments
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject550 struct {
	value *InlineObject550
	isSet bool
}

func (v NullableInlineObject550) Get() *InlineObject550 {
	return v.value
}

func (v *NullableInlineObject550) Set(val *InlineObject550) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject550) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject550) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject550(val *InlineObject550) *NullableInlineObject550 {
	return &NullableInlineObject550{value: val, isSet: true}
}

func (v NullableInlineObject550) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject550) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


