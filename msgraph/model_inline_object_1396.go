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

// InlineObject1396 struct for InlineObject1396
type InlineObject1396 struct {
	Value AnyOfobject `json:"value,omitempty"`
}

// NewInlineObject1396 instantiates a new InlineObject1396 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject1396() *InlineObject1396 {
	this := InlineObject1396{}
	return &this
}

// NewInlineObject1396WithDefaults instantiates a new InlineObject1396 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject1396WithDefaults() *InlineObject1396 {
	this := InlineObject1396{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1396) GetValue() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1396) GetValueOk() (*AnyOfobject, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return &o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *InlineObject1396) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given AnyOfobject and assigns it to the Value field.
func (o *InlineObject1396) SetValue(v AnyOfobject) {
	o.Value = v
}

func (o InlineObject1396) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject1396 struct {
	value *InlineObject1396
	isSet bool
}

func (v NullableInlineObject1396) Get() *InlineObject1396 {
	return v.value
}

func (v *NullableInlineObject1396) Set(val *InlineObject1396) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject1396) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject1396) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject1396(val *InlineObject1396) *NullableInlineObject1396 {
	return &NullableInlineObject1396{value: val, isSet: true}
}

func (v NullableInlineObject1396) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject1396) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


