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

// InlineObject1511 struct for InlineObject1511
type InlineObject1511 struct {
	Values AnyOfobject `json:"values,omitempty"`
}

// NewInlineObject1511 instantiates a new InlineObject1511 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject1511() *InlineObject1511 {
	this := InlineObject1511{}
	return &this
}

// NewInlineObject1511WithDefaults instantiates a new InlineObject1511 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject1511WithDefaults() *InlineObject1511 {
	this := InlineObject1511{}
	return &this
}

// GetValues returns the Values field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1511) GetValues() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1511) GetValuesOk() (*AnyOfobject, bool) {
	if o == nil || o.Values == nil {
		return nil, false
	}
	return &o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *InlineObject1511) HasValues() bool {
	if o != nil && o.Values != nil {
		return true
	}

	return false
}

// SetValues gets a reference to the given AnyOfobject and assigns it to the Values field.
func (o *InlineObject1511) SetValues(v AnyOfobject) {
	o.Values = v
}

func (o InlineObject1511) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Values != nil {
		toSerialize["values"] = o.Values
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject1511 struct {
	value *InlineObject1511
	isSet bool
}

func (v NullableInlineObject1511) Get() *InlineObject1511 {
	return v.value
}

func (v *NullableInlineObject1511) Set(val *InlineObject1511) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject1511) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject1511) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject1511(val *InlineObject1511) *NullableInlineObject1511 {
	return &NullableInlineObject1511{value: val, isSet: true}
}

func (v NullableInlineObject1511) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject1511) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


