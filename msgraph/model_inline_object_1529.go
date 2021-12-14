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

// InlineObject1529 struct for InlineObject1529
type InlineObject1529 struct {
	Number AnyOfobject `json:"number,omitempty"`
}

// NewInlineObject1529 instantiates a new InlineObject1529 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject1529() *InlineObject1529 {
	this := InlineObject1529{}
	return &this
}

// NewInlineObject1529WithDefaults instantiates a new InlineObject1529 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject1529WithDefaults() *InlineObject1529 {
	this := InlineObject1529{}
	return &this
}

// GetNumber returns the Number field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1529) GetNumber() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Number
}

// GetNumberOk returns a tuple with the Number field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1529) GetNumberOk() (*AnyOfobject, bool) {
	if o == nil || o.Number == nil {
		return nil, false
	}
	return &o.Number, true
}

// HasNumber returns a boolean if a field has been set.
func (o *InlineObject1529) HasNumber() bool {
	if o != nil && o.Number != nil {
		return true
	}

	return false
}

// SetNumber gets a reference to the given AnyOfobject and assigns it to the Number field.
func (o *InlineObject1529) SetNumber(v AnyOfobject) {
	o.Number = v
}

func (o InlineObject1529) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Number != nil {
		toSerialize["number"] = o.Number
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject1529 struct {
	value *InlineObject1529
	isSet bool
}

func (v NullableInlineObject1529) Get() *InlineObject1529 {
	return v.value
}

func (v *NullableInlineObject1529) Set(val *InlineObject1529) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject1529) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject1529) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject1529(val *InlineObject1529) *NullableInlineObject1529 {
	return &NullableInlineObject1529{value: val, isSet: true}
}

func (v NullableInlineObject1529) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject1529) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


