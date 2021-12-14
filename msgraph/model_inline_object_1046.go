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

// InlineObject1046 struct for InlineObject1046
type InlineObject1046 struct {
	Assignments *[]*AnyOfmicrosoftGraphTargetedManagedAppPolicyAssignment `json:"assignments,omitempty"`
}

// NewInlineObject1046 instantiates a new InlineObject1046 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject1046() *InlineObject1046 {
	this := InlineObject1046{}
	return &this
}

// NewInlineObject1046WithDefaults instantiates a new InlineObject1046 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject1046WithDefaults() *InlineObject1046 {
	this := InlineObject1046{}
	return &this
}

// GetAssignments returns the Assignments field value if set, zero value otherwise.
func (o *InlineObject1046) GetAssignments() []*AnyOfmicrosoftGraphTargetedManagedAppPolicyAssignment {
	if o == nil || o.Assignments == nil {
		var ret []*AnyOfmicrosoftGraphTargetedManagedAppPolicyAssignment
		return ret
	}
	return *o.Assignments
}

// GetAssignmentsOk returns a tuple with the Assignments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject1046) GetAssignmentsOk() (*[]*AnyOfmicrosoftGraphTargetedManagedAppPolicyAssignment, bool) {
	if o == nil || o.Assignments == nil {
		return nil, false
	}
	return o.Assignments, true
}

// HasAssignments returns a boolean if a field has been set.
func (o *InlineObject1046) HasAssignments() bool {
	if o != nil && o.Assignments != nil {
		return true
	}

	return false
}

// SetAssignments gets a reference to the given []*AnyOfmicrosoftGraphTargetedManagedAppPolicyAssignment and assigns it to the Assignments field.
func (o *InlineObject1046) SetAssignments(v []*AnyOfmicrosoftGraphTargetedManagedAppPolicyAssignment) {
	o.Assignments = &v
}

func (o InlineObject1046) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Assignments != nil {
		toSerialize["assignments"] = o.Assignments
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject1046 struct {
	value *InlineObject1046
	isSet bool
}

func (v NullableInlineObject1046) Get() *InlineObject1046 {
	return v.value
}

func (v *NullableInlineObject1046) Set(val *InlineObject1046) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject1046) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject1046) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject1046(val *InlineObject1046) *NullableInlineObject1046 {
	return &NullableInlineObject1046{value: val, isSet: true}
}

func (v NullableInlineObject1046) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject1046) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


