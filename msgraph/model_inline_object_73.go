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

// InlineObject73 struct for InlineObject73
type InlineObject73 struct {
	MobileAppAssignments *[]*AnyOfmicrosoftGraphMobileAppAssignment `json:"mobileAppAssignments,omitempty"`
}

// NewInlineObject73 instantiates a new InlineObject73 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject73() *InlineObject73 {
	this := InlineObject73{}
	return &this
}

// NewInlineObject73WithDefaults instantiates a new InlineObject73 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject73WithDefaults() *InlineObject73 {
	this := InlineObject73{}
	return &this
}

// GetMobileAppAssignments returns the MobileAppAssignments field value if set, zero value otherwise.
func (o *InlineObject73) GetMobileAppAssignments() []*AnyOfmicrosoftGraphMobileAppAssignment {
	if o == nil || o.MobileAppAssignments == nil {
		var ret []*AnyOfmicrosoftGraphMobileAppAssignment
		return ret
	}
	return *o.MobileAppAssignments
}

// GetMobileAppAssignmentsOk returns a tuple with the MobileAppAssignments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject73) GetMobileAppAssignmentsOk() (*[]*AnyOfmicrosoftGraphMobileAppAssignment, bool) {
	if o == nil || o.MobileAppAssignments == nil {
		return nil, false
	}
	return o.MobileAppAssignments, true
}

// HasMobileAppAssignments returns a boolean if a field has been set.
func (o *InlineObject73) HasMobileAppAssignments() bool {
	if o != nil && o.MobileAppAssignments != nil {
		return true
	}

	return false
}

// SetMobileAppAssignments gets a reference to the given []*AnyOfmicrosoftGraphMobileAppAssignment and assigns it to the MobileAppAssignments field.
func (o *InlineObject73) SetMobileAppAssignments(v []*AnyOfmicrosoftGraphMobileAppAssignment) {
	o.MobileAppAssignments = &v
}

func (o InlineObject73) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.MobileAppAssignments != nil {
		toSerialize["mobileAppAssignments"] = o.MobileAppAssignments
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject73 struct {
	value *InlineObject73
	isSet bool
}

func (v NullableInlineObject73) Get() *InlineObject73 {
	return v.value
}

func (v *NullableInlineObject73) Set(val *InlineObject73) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject73) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject73) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject73(val *InlineObject73) *NullableInlineObject73 {
	return &NullableInlineObject73{value: val, isSet: true}
}

func (v NullableInlineObject73) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject73) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


