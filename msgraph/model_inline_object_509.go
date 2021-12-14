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

// InlineObject509 struct for InlineObject509
type InlineObject509 struct {
	Shift *string `json:"shift,omitempty"`
}

// NewInlineObject509 instantiates a new InlineObject509 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject509() *InlineObject509 {
	this := InlineObject509{}
	return &this
}

// NewInlineObject509WithDefaults instantiates a new InlineObject509 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject509WithDefaults() *InlineObject509 {
	this := InlineObject509{}
	return &this
}

// GetShift returns the Shift field value if set, zero value otherwise.
func (o *InlineObject509) GetShift() string {
	if o == nil || o.Shift == nil {
		var ret string
		return ret
	}
	return *o.Shift
}

// GetShiftOk returns a tuple with the Shift field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject509) GetShiftOk() (*string, bool) {
	if o == nil || o.Shift == nil {
		return nil, false
	}
	return o.Shift, true
}

// HasShift returns a boolean if a field has been set.
func (o *InlineObject509) HasShift() bool {
	if o != nil && o.Shift != nil {
		return true
	}

	return false
}

// SetShift gets a reference to the given string and assigns it to the Shift field.
func (o *InlineObject509) SetShift(v string) {
	o.Shift = &v
}

func (o InlineObject509) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Shift != nil {
		toSerialize["shift"] = o.Shift
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject509 struct {
	value *InlineObject509
	isSet bool
}

func (v NullableInlineObject509) Get() *InlineObject509 {
	return v.value
}

func (v *NullableInlineObject509) Set(val *InlineObject509) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject509) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject509) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject509(val *InlineObject509) *NullableInlineObject509 {
	return &NullableInlineObject509{value: val, isSet: true}
}

func (v NullableInlineObject509) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject509) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


