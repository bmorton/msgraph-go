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

// InlineObject76 struct for InlineObject76
type InlineObject76 struct {
	Assignments *[]*AnyOfmicrosoftGraphDeviceCompliancePolicyAssignment `json:"assignments,omitempty"`
}

// NewInlineObject76 instantiates a new InlineObject76 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject76() *InlineObject76 {
	this := InlineObject76{}
	return &this
}

// NewInlineObject76WithDefaults instantiates a new InlineObject76 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject76WithDefaults() *InlineObject76 {
	this := InlineObject76{}
	return &this
}

// GetAssignments returns the Assignments field value if set, zero value otherwise.
func (o *InlineObject76) GetAssignments() []*AnyOfmicrosoftGraphDeviceCompliancePolicyAssignment {
	if o == nil || o.Assignments == nil {
		var ret []*AnyOfmicrosoftGraphDeviceCompliancePolicyAssignment
		return ret
	}
	return *o.Assignments
}

// GetAssignmentsOk returns a tuple with the Assignments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject76) GetAssignmentsOk() (*[]*AnyOfmicrosoftGraphDeviceCompliancePolicyAssignment, bool) {
	if o == nil || o.Assignments == nil {
		return nil, false
	}
	return o.Assignments, true
}

// HasAssignments returns a boolean if a field has been set.
func (o *InlineObject76) HasAssignments() bool {
	if o != nil && o.Assignments != nil {
		return true
	}

	return false
}

// SetAssignments gets a reference to the given []*AnyOfmicrosoftGraphDeviceCompliancePolicyAssignment and assigns it to the Assignments field.
func (o *InlineObject76) SetAssignments(v []*AnyOfmicrosoftGraphDeviceCompliancePolicyAssignment) {
	o.Assignments = &v
}

func (o InlineObject76) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Assignments != nil {
		toSerialize["assignments"] = o.Assignments
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject76 struct {
	value *InlineObject76
	isSet bool
}

func (v NullableInlineObject76) Get() *InlineObject76 {
	return v.value
}

func (v *NullableInlineObject76) Set(val *InlineObject76) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject76) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject76) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject76(val *InlineObject76) *NullableInlineObject76 {
	return &NullableInlineObject76{value: val, isSet: true}
}

func (v NullableInlineObject76) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject76) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


