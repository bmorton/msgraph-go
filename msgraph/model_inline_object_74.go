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

// InlineObject74 struct for InlineObject74
type InlineObject74 struct {
	Assignments *[]*AnyOfmicrosoftGraphTargetedManagedAppPolicyAssignment `json:"assignments,omitempty"`
}

// NewInlineObject74 instantiates a new InlineObject74 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject74() *InlineObject74 {
	this := InlineObject74{}
	return &this
}

// NewInlineObject74WithDefaults instantiates a new InlineObject74 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject74WithDefaults() *InlineObject74 {
	this := InlineObject74{}
	return &this
}

// GetAssignments returns the Assignments field value if set, zero value otherwise.
func (o *InlineObject74) GetAssignments() []*AnyOfmicrosoftGraphTargetedManagedAppPolicyAssignment {
	if o == nil || o.Assignments == nil {
		var ret []*AnyOfmicrosoftGraphTargetedManagedAppPolicyAssignment
		return ret
	}
	return *o.Assignments
}

// GetAssignmentsOk returns a tuple with the Assignments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject74) GetAssignmentsOk() (*[]*AnyOfmicrosoftGraphTargetedManagedAppPolicyAssignment, bool) {
	if o == nil || o.Assignments == nil {
		return nil, false
	}
	return o.Assignments, true
}

// HasAssignments returns a boolean if a field has been set.
func (o *InlineObject74) HasAssignments() bool {
	if o != nil && o.Assignments != nil {
		return true
	}

	return false
}

// SetAssignments gets a reference to the given []*AnyOfmicrosoftGraphTargetedManagedAppPolicyAssignment and assigns it to the Assignments field.
func (o *InlineObject74) SetAssignments(v []*AnyOfmicrosoftGraphTargetedManagedAppPolicyAssignment) {
	o.Assignments = &v
}

func (o InlineObject74) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Assignments != nil {
		toSerialize["assignments"] = o.Assignments
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject74 struct {
	value *InlineObject74
	isSet bool
}

func (v NullableInlineObject74) Get() *InlineObject74 {
	return v.value
}

func (v *NullableInlineObject74) Set(val *InlineObject74) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject74) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject74) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject74(val *InlineObject74) *NullableInlineObject74 {
	return &NullableInlineObject74{value: val, isSet: true}
}

func (v NullableInlineObject74) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject74) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


