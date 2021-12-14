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

// InlineObject1490 struct for InlineObject1490
type InlineObject1490 struct {
	Number AnyOfobject `json:"number,omitempty"`
	NumDigits AnyOfobject `json:"numDigits,omitempty"`
}

// NewInlineObject1490 instantiates a new InlineObject1490 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject1490() *InlineObject1490 {
	this := InlineObject1490{}
	return &this
}

// NewInlineObject1490WithDefaults instantiates a new InlineObject1490 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject1490WithDefaults() *InlineObject1490 {
	this := InlineObject1490{}
	return &this
}

// GetNumber returns the Number field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1490) GetNumber() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Number
}

// GetNumberOk returns a tuple with the Number field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1490) GetNumberOk() (*AnyOfobject, bool) {
	if o == nil || o.Number == nil {
		return nil, false
	}
	return &o.Number, true
}

// HasNumber returns a boolean if a field has been set.
func (o *InlineObject1490) HasNumber() bool {
	if o != nil && o.Number != nil {
		return true
	}

	return false
}

// SetNumber gets a reference to the given AnyOfobject and assigns it to the Number field.
func (o *InlineObject1490) SetNumber(v AnyOfobject) {
	o.Number = v
}

// GetNumDigits returns the NumDigits field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1490) GetNumDigits() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.NumDigits
}

// GetNumDigitsOk returns a tuple with the NumDigits field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1490) GetNumDigitsOk() (*AnyOfobject, bool) {
	if o == nil || o.NumDigits == nil {
		return nil, false
	}
	return &o.NumDigits, true
}

// HasNumDigits returns a boolean if a field has been set.
func (o *InlineObject1490) HasNumDigits() bool {
	if o != nil && o.NumDigits != nil {
		return true
	}

	return false
}

// SetNumDigits gets a reference to the given AnyOfobject and assigns it to the NumDigits field.
func (o *InlineObject1490) SetNumDigits(v AnyOfobject) {
	o.NumDigits = v
}

func (o InlineObject1490) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Number != nil {
		toSerialize["number"] = o.Number
	}
	if o.NumDigits != nil {
		toSerialize["numDigits"] = o.NumDigits
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject1490 struct {
	value *InlineObject1490
	isSet bool
}

func (v NullableInlineObject1490) Get() *InlineObject1490 {
	return v.value
}

func (v *NullableInlineObject1490) Set(val *InlineObject1490) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject1490) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject1490) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject1490(val *InlineObject1490) *NullableInlineObject1490 {
	return &NullableInlineObject1490{value: val, isSet: true}
}

func (v NullableInlineObject1490) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject1490) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


