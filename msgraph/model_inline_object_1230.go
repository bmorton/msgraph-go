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

// InlineObject1230 struct for InlineObject1230
type InlineObject1230 struct {
	Number AnyOfobject `json:"number,omitempty"`
	Radix AnyOfobject `json:"radix,omitempty"`
	MinLength AnyOfobject `json:"minLength,omitempty"`
}

// NewInlineObject1230 instantiates a new InlineObject1230 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject1230() *InlineObject1230 {
	this := InlineObject1230{}
	return &this
}

// NewInlineObject1230WithDefaults instantiates a new InlineObject1230 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject1230WithDefaults() *InlineObject1230 {
	this := InlineObject1230{}
	return &this
}

// GetNumber returns the Number field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1230) GetNumber() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Number
}

// GetNumberOk returns a tuple with the Number field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1230) GetNumberOk() (*AnyOfobject, bool) {
	if o == nil || o.Number == nil {
		return nil, false
	}
	return &o.Number, true
}

// HasNumber returns a boolean if a field has been set.
func (o *InlineObject1230) HasNumber() bool {
	if o != nil && o.Number != nil {
		return true
	}

	return false
}

// SetNumber gets a reference to the given AnyOfobject and assigns it to the Number field.
func (o *InlineObject1230) SetNumber(v AnyOfobject) {
	o.Number = v
}

// GetRadix returns the Radix field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1230) GetRadix() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Radix
}

// GetRadixOk returns a tuple with the Radix field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1230) GetRadixOk() (*AnyOfobject, bool) {
	if o == nil || o.Radix == nil {
		return nil, false
	}
	return &o.Radix, true
}

// HasRadix returns a boolean if a field has been set.
func (o *InlineObject1230) HasRadix() bool {
	if o != nil && o.Radix != nil {
		return true
	}

	return false
}

// SetRadix gets a reference to the given AnyOfobject and assigns it to the Radix field.
func (o *InlineObject1230) SetRadix(v AnyOfobject) {
	o.Radix = v
}

// GetMinLength returns the MinLength field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1230) GetMinLength() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.MinLength
}

// GetMinLengthOk returns a tuple with the MinLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1230) GetMinLengthOk() (*AnyOfobject, bool) {
	if o == nil || o.MinLength == nil {
		return nil, false
	}
	return &o.MinLength, true
}

// HasMinLength returns a boolean if a field has been set.
func (o *InlineObject1230) HasMinLength() bool {
	if o != nil && o.MinLength != nil {
		return true
	}

	return false
}

// SetMinLength gets a reference to the given AnyOfobject and assigns it to the MinLength field.
func (o *InlineObject1230) SetMinLength(v AnyOfobject) {
	o.MinLength = v
}

func (o InlineObject1230) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Number != nil {
		toSerialize["number"] = o.Number
	}
	if o.Radix != nil {
		toSerialize["radix"] = o.Radix
	}
	if o.MinLength != nil {
		toSerialize["minLength"] = o.MinLength
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject1230 struct {
	value *InlineObject1230
	isSet bool
}

func (v NullableInlineObject1230) Get() *InlineObject1230 {
	return v.value
}

func (v *NullableInlineObject1230) Set(val *InlineObject1230) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject1230) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject1230) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject1230(val *InlineObject1230) *NullableInlineObject1230 {
	return &NullableInlineObject1230{value: val, isSet: true}
}

func (v NullableInlineObject1230) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject1230) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


