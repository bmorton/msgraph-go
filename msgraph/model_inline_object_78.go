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

// InlineObject78 struct for InlineObject78
type InlineObject78 struct {
	Assignments *[]*AnyOfmicrosoftGraphDeviceConfigurationAssignment `json:"assignments,omitempty"`
}

// NewInlineObject78 instantiates a new InlineObject78 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject78() *InlineObject78 {
	this := InlineObject78{}
	return &this
}

// NewInlineObject78WithDefaults instantiates a new InlineObject78 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject78WithDefaults() *InlineObject78 {
	this := InlineObject78{}
	return &this
}

// GetAssignments returns the Assignments field value if set, zero value otherwise.
func (o *InlineObject78) GetAssignments() []*AnyOfmicrosoftGraphDeviceConfigurationAssignment {
	if o == nil || o.Assignments == nil {
		var ret []*AnyOfmicrosoftGraphDeviceConfigurationAssignment
		return ret
	}
	return *o.Assignments
}

// GetAssignmentsOk returns a tuple with the Assignments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject78) GetAssignmentsOk() (*[]*AnyOfmicrosoftGraphDeviceConfigurationAssignment, bool) {
	if o == nil || o.Assignments == nil {
		return nil, false
	}
	return o.Assignments, true
}

// HasAssignments returns a boolean if a field has been set.
func (o *InlineObject78) HasAssignments() bool {
	if o != nil && o.Assignments != nil {
		return true
	}

	return false
}

// SetAssignments gets a reference to the given []*AnyOfmicrosoftGraphDeviceConfigurationAssignment and assigns it to the Assignments field.
func (o *InlineObject78) SetAssignments(v []*AnyOfmicrosoftGraphDeviceConfigurationAssignment) {
	o.Assignments = &v
}

func (o InlineObject78) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Assignments != nil {
		toSerialize["assignments"] = o.Assignments
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject78 struct {
	value *InlineObject78
	isSet bool
}

func (v NullableInlineObject78) Get() *InlineObject78 {
	return v.value
}

func (v *NullableInlineObject78) Set(val *InlineObject78) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject78) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject78) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject78(val *InlineObject78) *NullableInlineObject78 {
	return &NullableInlineObject78{value: val, isSet: true}
}

func (v NullableInlineObject78) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject78) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


