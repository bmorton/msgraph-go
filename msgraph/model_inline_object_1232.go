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

// InlineObject1232 struct for InlineObject1232
type InlineObject1232 struct {
	X AnyOfobject `json:"x,omitempty"`
	N AnyOfobject `json:"n,omitempty"`
}

// NewInlineObject1232 instantiates a new InlineObject1232 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject1232() *InlineObject1232 {
	this := InlineObject1232{}
	return &this
}

// NewInlineObject1232WithDefaults instantiates a new InlineObject1232 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject1232WithDefaults() *InlineObject1232 {
	this := InlineObject1232{}
	return &this
}

// GetX returns the X field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1232) GetX() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.X
}

// GetXOk returns a tuple with the X field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1232) GetXOk() (*AnyOfobject, bool) {
	if o == nil || o.X == nil {
		return nil, false
	}
	return &o.X, true
}

// HasX returns a boolean if a field has been set.
func (o *InlineObject1232) HasX() bool {
	if o != nil && o.X != nil {
		return true
	}

	return false
}

// SetX gets a reference to the given AnyOfobject and assigns it to the X field.
func (o *InlineObject1232) SetX(v AnyOfobject) {
	o.X = v
}

// GetN returns the N field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1232) GetN() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.N
}

// GetNOk returns a tuple with the N field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1232) GetNOk() (*AnyOfobject, bool) {
	if o == nil || o.N == nil {
		return nil, false
	}
	return &o.N, true
}

// HasN returns a boolean if a field has been set.
func (o *InlineObject1232) HasN() bool {
	if o != nil && o.N != nil {
		return true
	}

	return false
}

// SetN gets a reference to the given AnyOfobject and assigns it to the N field.
func (o *InlineObject1232) SetN(v AnyOfobject) {
	o.N = v
}

func (o InlineObject1232) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.X != nil {
		toSerialize["x"] = o.X
	}
	if o.N != nil {
		toSerialize["n"] = o.N
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject1232 struct {
	value *InlineObject1232
	isSet bool
}

func (v NullableInlineObject1232) Get() *InlineObject1232 {
	return v.value
}

func (v *NullableInlineObject1232) Set(val *InlineObject1232) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject1232) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject1232) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject1232(val *InlineObject1232) *NullableInlineObject1232 {
	return &NullableInlineObject1232{value: val, isSet: true}
}

func (v NullableInlineObject1232) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject1232) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

