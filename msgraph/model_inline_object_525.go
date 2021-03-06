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

// InlineObject525 struct for InlineObject525
type InlineObject525 struct {
	Shift *string `json:"shift,omitempty"`
}

// NewInlineObject525 instantiates a new InlineObject525 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject525() *InlineObject525 {
	this := InlineObject525{}
	return &this
}

// NewInlineObject525WithDefaults instantiates a new InlineObject525 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject525WithDefaults() *InlineObject525 {
	this := InlineObject525{}
	return &this
}

// GetShift returns the Shift field value if set, zero value otherwise.
func (o *InlineObject525) GetShift() string {
	if o == nil || o.Shift == nil {
		var ret string
		return ret
	}
	return *o.Shift
}

// GetShiftOk returns a tuple with the Shift field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject525) GetShiftOk() (*string, bool) {
	if o == nil || o.Shift == nil {
		return nil, false
	}
	return o.Shift, true
}

// HasShift returns a boolean if a field has been set.
func (o *InlineObject525) HasShift() bool {
	if o != nil && o.Shift != nil {
		return true
	}

	return false
}

// SetShift gets a reference to the given string and assigns it to the Shift field.
func (o *InlineObject525) SetShift(v string) {
	o.Shift = &v
}

func (o InlineObject525) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Shift != nil {
		toSerialize["shift"] = o.Shift
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject525 struct {
	value *InlineObject525
	isSet bool
}

func (v NullableInlineObject525) Get() *InlineObject525 {
	return v.value
}

func (v *NullableInlineObject525) Set(val *InlineObject525) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject525) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject525) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject525(val *InlineObject525) *NullableInlineObject525 {
	return &NullableInlineObject525{value: val, isSet: true}
}

func (v NullableInlineObject525) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject525) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


