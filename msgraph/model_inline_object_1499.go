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

// InlineObject1499 struct for InlineObject1499
type InlineObject1499 struct {
	Value AnyOfobject `json:"value,omitempty"`
}

// NewInlineObject1499 instantiates a new InlineObject1499 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject1499() *InlineObject1499 {
	this := InlineObject1499{}
	return &this
}

// NewInlineObject1499WithDefaults instantiates a new InlineObject1499 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject1499WithDefaults() *InlineObject1499 {
	this := InlineObject1499{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1499) GetValue() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1499) GetValueOk() (*AnyOfobject, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return &o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *InlineObject1499) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given AnyOfobject and assigns it to the Value field.
func (o *InlineObject1499) SetValue(v AnyOfobject) {
	o.Value = v
}

func (o InlineObject1499) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject1499 struct {
	value *InlineObject1499
	isSet bool
}

func (v NullableInlineObject1499) Get() *InlineObject1499 {
	return v.value
}

func (v *NullableInlineObject1499) Set(val *InlineObject1499) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject1499) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject1499) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject1499(val *InlineObject1499) *NullableInlineObject1499 {
	return &NullableInlineObject1499{value: val, isSet: true}
}

func (v NullableInlineObject1499) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject1499) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


