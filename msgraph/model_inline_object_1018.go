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

// InlineObject1018 struct for InlineObject1018
type InlineObject1018 struct {
	Shift *string `json:"shift,omitempty"`
}

// NewInlineObject1018 instantiates a new InlineObject1018 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject1018() *InlineObject1018 {
	this := InlineObject1018{}
	return &this
}

// NewInlineObject1018WithDefaults instantiates a new InlineObject1018 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject1018WithDefaults() *InlineObject1018 {
	this := InlineObject1018{}
	return &this
}

// GetShift returns the Shift field value if set, zero value otherwise.
func (o *InlineObject1018) GetShift() string {
	if o == nil || o.Shift == nil {
		var ret string
		return ret
	}
	return *o.Shift
}

// GetShiftOk returns a tuple with the Shift field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject1018) GetShiftOk() (*string, bool) {
	if o == nil || o.Shift == nil {
		return nil, false
	}
	return o.Shift, true
}

// HasShift returns a boolean if a field has been set.
func (o *InlineObject1018) HasShift() bool {
	if o != nil && o.Shift != nil {
		return true
	}

	return false
}

// SetShift gets a reference to the given string and assigns it to the Shift field.
func (o *InlineObject1018) SetShift(v string) {
	o.Shift = &v
}

func (o InlineObject1018) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Shift != nil {
		toSerialize["shift"] = o.Shift
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject1018 struct {
	value *InlineObject1018
	isSet bool
}

func (v NullableInlineObject1018) Get() *InlineObject1018 {
	return v.value
}

func (v *NullableInlineObject1018) Set(val *InlineObject1018) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject1018) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject1018) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject1018(val *InlineObject1018) *NullableInlineObject1018 {
	return &NullableInlineObject1018{value: val, isSet: true}
}

func (v NullableInlineObject1018) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject1018) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

