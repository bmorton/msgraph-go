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

// InlineObject1281 struct for InlineObject1281
type InlineObject1281 struct {
	Number AnyOfobject `json:"number,omitempty"`
}

// NewInlineObject1281 instantiates a new InlineObject1281 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject1281() *InlineObject1281 {
	this := InlineObject1281{}
	return &this
}

// NewInlineObject1281WithDefaults instantiates a new InlineObject1281 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject1281WithDefaults() *InlineObject1281 {
	this := InlineObject1281{}
	return &this
}

// GetNumber returns the Number field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1281) GetNumber() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Number
}

// GetNumberOk returns a tuple with the Number field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1281) GetNumberOk() (*AnyOfobject, bool) {
	if o == nil || o.Number == nil {
		return nil, false
	}
	return &o.Number, true
}

// HasNumber returns a boolean if a field has been set.
func (o *InlineObject1281) HasNumber() bool {
	if o != nil && o.Number != nil {
		return true
	}

	return false
}

// SetNumber gets a reference to the given AnyOfobject and assigns it to the Number field.
func (o *InlineObject1281) SetNumber(v AnyOfobject) {
	o.Number = v
}

func (o InlineObject1281) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Number != nil {
		toSerialize["number"] = o.Number
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject1281 struct {
	value *InlineObject1281
	isSet bool
}

func (v NullableInlineObject1281) Get() *InlineObject1281 {
	return v.value
}

func (v *NullableInlineObject1281) Set(val *InlineObject1281) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject1281) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject1281) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject1281(val *InlineObject1281) *NullableInlineObject1281 {
	return &NullableInlineObject1281{value: val, isSet: true}
}

func (v NullableInlineObject1281) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject1281) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


