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

// InlineObject1496 struct for InlineObject1496
type InlineObject1496 struct {
	Number AnyOfobject `json:"number,omitempty"`
}

// NewInlineObject1496 instantiates a new InlineObject1496 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject1496() *InlineObject1496 {
	this := InlineObject1496{}
	return &this
}

// NewInlineObject1496WithDefaults instantiates a new InlineObject1496 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject1496WithDefaults() *InlineObject1496 {
	this := InlineObject1496{}
	return &this
}

// GetNumber returns the Number field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1496) GetNumber() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Number
}

// GetNumberOk returns a tuple with the Number field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1496) GetNumberOk() (*AnyOfobject, bool) {
	if o == nil || o.Number == nil {
		return nil, false
	}
	return &o.Number, true
}

// HasNumber returns a boolean if a field has been set.
func (o *InlineObject1496) HasNumber() bool {
	if o != nil && o.Number != nil {
		return true
	}

	return false
}

// SetNumber gets a reference to the given AnyOfobject and assigns it to the Number field.
func (o *InlineObject1496) SetNumber(v AnyOfobject) {
	o.Number = v
}

func (o InlineObject1496) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Number != nil {
		toSerialize["number"] = o.Number
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject1496 struct {
	value *InlineObject1496
	isSet bool
}

func (v NullableInlineObject1496) Get() *InlineObject1496 {
	return v.value
}

func (v *NullableInlineObject1496) Set(val *InlineObject1496) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject1496) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject1496) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject1496(val *InlineObject1496) *NullableInlineObject1496 {
	return &NullableInlineObject1496{value: val, isSet: true}
}

func (v NullableInlineObject1496) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject1496) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


