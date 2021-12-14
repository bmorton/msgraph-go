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

// InlineObject1348 struct for InlineObject1348
type InlineObject1348 struct {
	X AnyOfobject `json:"x,omitempty"`
}

// NewInlineObject1348 instantiates a new InlineObject1348 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject1348() *InlineObject1348 {
	this := InlineObject1348{}
	return &this
}

// NewInlineObject1348WithDefaults instantiates a new InlineObject1348 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject1348WithDefaults() *InlineObject1348 {
	this := InlineObject1348{}
	return &this
}

// GetX returns the X field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1348) GetX() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.X
}

// GetXOk returns a tuple with the X field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1348) GetXOk() (*AnyOfobject, bool) {
	if o == nil || o.X == nil {
		return nil, false
	}
	return &o.X, true
}

// HasX returns a boolean if a field has been set.
func (o *InlineObject1348) HasX() bool {
	if o != nil && o.X != nil {
		return true
	}

	return false
}

// SetX gets a reference to the given AnyOfobject and assigns it to the X field.
func (o *InlineObject1348) SetX(v AnyOfobject) {
	o.X = v
}

func (o InlineObject1348) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.X != nil {
		toSerialize["x"] = o.X
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject1348 struct {
	value *InlineObject1348
	isSet bool
}

func (v NullableInlineObject1348) Get() *InlineObject1348 {
	return v.value
}

func (v *NullableInlineObject1348) Set(val *InlineObject1348) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject1348) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject1348) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject1348(val *InlineObject1348) *NullableInlineObject1348 {
	return &NullableInlineObject1348{value: val, isSet: true}
}

func (v NullableInlineObject1348) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject1348) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


