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

// InlineObject1299 struct for InlineObject1299
type InlineObject1299 struct {
	Number AnyOfobject `json:"number,omitempty"`
	Radix AnyOfobject `json:"radix,omitempty"`
}

// NewInlineObject1299 instantiates a new InlineObject1299 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject1299() *InlineObject1299 {
	this := InlineObject1299{}
	return &this
}

// NewInlineObject1299WithDefaults instantiates a new InlineObject1299 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject1299WithDefaults() *InlineObject1299 {
	this := InlineObject1299{}
	return &this
}

// GetNumber returns the Number field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1299) GetNumber() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Number
}

// GetNumberOk returns a tuple with the Number field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1299) GetNumberOk() (*AnyOfobject, bool) {
	if o == nil || o.Number == nil {
		return nil, false
	}
	return &o.Number, true
}

// HasNumber returns a boolean if a field has been set.
func (o *InlineObject1299) HasNumber() bool {
	if o != nil && o.Number != nil {
		return true
	}

	return false
}

// SetNumber gets a reference to the given AnyOfobject and assigns it to the Number field.
func (o *InlineObject1299) SetNumber(v AnyOfobject) {
	o.Number = v
}

// GetRadix returns the Radix field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1299) GetRadix() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Radix
}

// GetRadixOk returns a tuple with the Radix field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1299) GetRadixOk() (*AnyOfobject, bool) {
	if o == nil || o.Radix == nil {
		return nil, false
	}
	return &o.Radix, true
}

// HasRadix returns a boolean if a field has been set.
func (o *InlineObject1299) HasRadix() bool {
	if o != nil && o.Radix != nil {
		return true
	}

	return false
}

// SetRadix gets a reference to the given AnyOfobject and assigns it to the Radix field.
func (o *InlineObject1299) SetRadix(v AnyOfobject) {
	o.Radix = v
}

func (o InlineObject1299) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Number != nil {
		toSerialize["number"] = o.Number
	}
	if o.Radix != nil {
		toSerialize["radix"] = o.Radix
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject1299 struct {
	value *InlineObject1299
	isSet bool
}

func (v NullableInlineObject1299) Get() *InlineObject1299 {
	return v.value
}

func (v *NullableInlineObject1299) Set(val *InlineObject1299) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject1299) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject1299) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject1299(val *InlineObject1299) *NullableInlineObject1299 {
	return &NullableInlineObject1299{value: val, isSet: true}
}

func (v NullableInlineObject1299) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject1299) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


