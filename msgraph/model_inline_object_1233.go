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

// InlineObject1233 struct for InlineObject1233
type InlineObject1233 struct {
	X AnyOfobject `json:"x,omitempty"`
	N AnyOfobject `json:"n,omitempty"`
}

// NewInlineObject1233 instantiates a new InlineObject1233 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject1233() *InlineObject1233 {
	this := InlineObject1233{}
	return &this
}

// NewInlineObject1233WithDefaults instantiates a new InlineObject1233 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject1233WithDefaults() *InlineObject1233 {
	this := InlineObject1233{}
	return &this
}

// GetX returns the X field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1233) GetX() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.X
}

// GetXOk returns a tuple with the X field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1233) GetXOk() (*AnyOfobject, bool) {
	if o == nil || o.X == nil {
		return nil, false
	}
	return &o.X, true
}

// HasX returns a boolean if a field has been set.
func (o *InlineObject1233) HasX() bool {
	if o != nil && o.X != nil {
		return true
	}

	return false
}

// SetX gets a reference to the given AnyOfobject and assigns it to the X field.
func (o *InlineObject1233) SetX(v AnyOfobject) {
	o.X = v
}

// GetN returns the N field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1233) GetN() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.N
}

// GetNOk returns a tuple with the N field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1233) GetNOk() (*AnyOfobject, bool) {
	if o == nil || o.N == nil {
		return nil, false
	}
	return &o.N, true
}

// HasN returns a boolean if a field has been set.
func (o *InlineObject1233) HasN() bool {
	if o != nil && o.N != nil {
		return true
	}

	return false
}

// SetN gets a reference to the given AnyOfobject and assigns it to the N field.
func (o *InlineObject1233) SetN(v AnyOfobject) {
	o.N = v
}

func (o InlineObject1233) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.X != nil {
		toSerialize["x"] = o.X
	}
	if o.N != nil {
		toSerialize["n"] = o.N
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject1233 struct {
	value *InlineObject1233
	isSet bool
}

func (v NullableInlineObject1233) Get() *InlineObject1233 {
	return v.value
}

func (v *NullableInlineObject1233) Set(val *InlineObject1233) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject1233) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject1233) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject1233(val *InlineObject1233) *NullableInlineObject1233 {
	return &NullableInlineObject1233{value: val, isSet: true}
}

func (v NullableInlineObject1233) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject1233) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


