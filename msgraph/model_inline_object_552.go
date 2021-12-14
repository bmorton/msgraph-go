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

// InlineObject552 struct for InlineObject552
type InlineObject552 struct {
	Assignments *[]*AnyOfmicrosoftGraphTargetedManagedAppPolicyAssignment `json:"assignments,omitempty"`
}

// NewInlineObject552 instantiates a new InlineObject552 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject552() *InlineObject552 {
	this := InlineObject552{}
	return &this
}

// NewInlineObject552WithDefaults instantiates a new InlineObject552 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject552WithDefaults() *InlineObject552 {
	this := InlineObject552{}
	return &this
}

// GetAssignments returns the Assignments field value if set, zero value otherwise.
func (o *InlineObject552) GetAssignments() []*AnyOfmicrosoftGraphTargetedManagedAppPolicyAssignment {
	if o == nil || o.Assignments == nil {
		var ret []*AnyOfmicrosoftGraphTargetedManagedAppPolicyAssignment
		return ret
	}
	return *o.Assignments
}

// GetAssignmentsOk returns a tuple with the Assignments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject552) GetAssignmentsOk() (*[]*AnyOfmicrosoftGraphTargetedManagedAppPolicyAssignment, bool) {
	if o == nil || o.Assignments == nil {
		return nil, false
	}
	return o.Assignments, true
}

// HasAssignments returns a boolean if a field has been set.
func (o *InlineObject552) HasAssignments() bool {
	if o != nil && o.Assignments != nil {
		return true
	}

	return false
}

// SetAssignments gets a reference to the given []*AnyOfmicrosoftGraphTargetedManagedAppPolicyAssignment and assigns it to the Assignments field.
func (o *InlineObject552) SetAssignments(v []*AnyOfmicrosoftGraphTargetedManagedAppPolicyAssignment) {
	o.Assignments = &v
}

func (o InlineObject552) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Assignments != nil {
		toSerialize["assignments"] = o.Assignments
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject552 struct {
	value *InlineObject552
	isSet bool
}

func (v NullableInlineObject552) Get() *InlineObject552 {
	return v.value
}

func (v *NullableInlineObject552) Set(val *InlineObject552) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject552) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject552) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject552(val *InlineObject552) *NullableInlineObject552 {
	return &NullableInlineObject552{value: val, isSet: true}
}

func (v NullableInlineObject552) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject552) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


