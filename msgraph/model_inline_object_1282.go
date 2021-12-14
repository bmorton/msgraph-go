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

// InlineObject1282 struct for InlineObject1282
type InlineObject1282 struct {
	Number AnyOfobject `json:"number,omitempty"`
}

// NewInlineObject1282 instantiates a new InlineObject1282 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject1282() *InlineObject1282 {
	this := InlineObject1282{}
	return &this
}

// NewInlineObject1282WithDefaults instantiates a new InlineObject1282 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject1282WithDefaults() *InlineObject1282 {
	this := InlineObject1282{}
	return &this
}

// GetNumber returns the Number field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1282) GetNumber() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Number
}

// GetNumberOk returns a tuple with the Number field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1282) GetNumberOk() (*AnyOfobject, bool) {
	if o == nil || o.Number == nil {
		return nil, false
	}
	return &o.Number, true
}

// HasNumber returns a boolean if a field has been set.
func (o *InlineObject1282) HasNumber() bool {
	if o != nil && o.Number != nil {
		return true
	}

	return false
}

// SetNumber gets a reference to the given AnyOfobject and assigns it to the Number field.
func (o *InlineObject1282) SetNumber(v AnyOfobject) {
	o.Number = v
}

func (o InlineObject1282) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Number != nil {
		toSerialize["number"] = o.Number
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject1282 struct {
	value *InlineObject1282
	isSet bool
}

func (v NullableInlineObject1282) Get() *InlineObject1282 {
	return v.value
}

func (v *NullableInlineObject1282) Set(val *InlineObject1282) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject1282) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject1282) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject1282(val *InlineObject1282) *NullableInlineObject1282 {
	return &NullableInlineObject1282{value: val, isSet: true}
}

func (v NullableInlineObject1282) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject1282) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


