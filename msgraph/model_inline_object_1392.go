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

// InlineObject1392 struct for InlineObject1392
type InlineObject1392 struct {
	Value AnyOfobject `json:"value,omitempty"`
}

// NewInlineObject1392 instantiates a new InlineObject1392 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject1392() *InlineObject1392 {
	this := InlineObject1392{}
	return &this
}

// NewInlineObject1392WithDefaults instantiates a new InlineObject1392 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject1392WithDefaults() *InlineObject1392 {
	this := InlineObject1392{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1392) GetValue() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1392) GetValueOk() (*AnyOfobject, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return &o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *InlineObject1392) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given AnyOfobject and assigns it to the Value field.
func (o *InlineObject1392) SetValue(v AnyOfobject) {
	o.Value = v
}

func (o InlineObject1392) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject1392 struct {
	value *InlineObject1392
	isSet bool
}

func (v NullableInlineObject1392) Get() *InlineObject1392 {
	return v.value
}

func (v *NullableInlineObject1392) Set(val *InlineObject1392) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject1392) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject1392) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject1392(val *InlineObject1392) *NullableInlineObject1392 {
	return &NullableInlineObject1392{value: val, isSet: true}
}

func (v NullableInlineObject1392) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject1392) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


