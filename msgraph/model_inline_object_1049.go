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

// InlineObject1049 struct for InlineObject1049
type InlineObject1049 struct {
	Shift *string `json:"shift,omitempty"`
}

// NewInlineObject1049 instantiates a new InlineObject1049 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject1049() *InlineObject1049 {
	this := InlineObject1049{}
	return &this
}

// NewInlineObject1049WithDefaults instantiates a new InlineObject1049 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject1049WithDefaults() *InlineObject1049 {
	this := InlineObject1049{}
	return &this
}

// GetShift returns the Shift field value if set, zero value otherwise.
func (o *InlineObject1049) GetShift() string {
	if o == nil || o.Shift == nil {
		var ret string
		return ret
	}
	return *o.Shift
}

// GetShiftOk returns a tuple with the Shift field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject1049) GetShiftOk() (*string, bool) {
	if o == nil || o.Shift == nil {
		return nil, false
	}
	return o.Shift, true
}

// HasShift returns a boolean if a field has been set.
func (o *InlineObject1049) HasShift() bool {
	if o != nil && o.Shift != nil {
		return true
	}

	return false
}

// SetShift gets a reference to the given string and assigns it to the Shift field.
func (o *InlineObject1049) SetShift(v string) {
	o.Shift = &v
}

func (o InlineObject1049) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Shift != nil {
		toSerialize["shift"] = o.Shift
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject1049 struct {
	value *InlineObject1049
	isSet bool
}

func (v NullableInlineObject1049) Get() *InlineObject1049 {
	return v.value
}

func (v *NullableInlineObject1049) Set(val *InlineObject1049) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject1049) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject1049) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject1049(val *InlineObject1049) *NullableInlineObject1049 {
	return &NullableInlineObject1049{value: val, isSet: true}
}

func (v NullableInlineObject1049) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject1049) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


