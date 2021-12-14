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

// InlineObject1239 struct for InlineObject1239
type InlineObject1239 struct {
	Number AnyOfobject `json:"number,omitempty"`
	Places AnyOfobject `json:"places,omitempty"`
}

// NewInlineObject1239 instantiates a new InlineObject1239 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject1239() *InlineObject1239 {
	this := InlineObject1239{}
	return &this
}

// NewInlineObject1239WithDefaults instantiates a new InlineObject1239 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject1239WithDefaults() *InlineObject1239 {
	this := InlineObject1239{}
	return &this
}

// GetNumber returns the Number field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1239) GetNumber() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Number
}

// GetNumberOk returns a tuple with the Number field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1239) GetNumberOk() (*AnyOfobject, bool) {
	if o == nil || o.Number == nil {
		return nil, false
	}
	return &o.Number, true
}

// HasNumber returns a boolean if a field has been set.
func (o *InlineObject1239) HasNumber() bool {
	if o != nil && o.Number != nil {
		return true
	}

	return false
}

// SetNumber gets a reference to the given AnyOfobject and assigns it to the Number field.
func (o *InlineObject1239) SetNumber(v AnyOfobject) {
	o.Number = v
}

// GetPlaces returns the Places field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1239) GetPlaces() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Places
}

// GetPlacesOk returns a tuple with the Places field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1239) GetPlacesOk() (*AnyOfobject, bool) {
	if o == nil || o.Places == nil {
		return nil, false
	}
	return &o.Places, true
}

// HasPlaces returns a boolean if a field has been set.
func (o *InlineObject1239) HasPlaces() bool {
	if o != nil && o.Places != nil {
		return true
	}

	return false
}

// SetPlaces gets a reference to the given AnyOfobject and assigns it to the Places field.
func (o *InlineObject1239) SetPlaces(v AnyOfobject) {
	o.Places = v
}

func (o InlineObject1239) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Number != nil {
		toSerialize["number"] = o.Number
	}
	if o.Places != nil {
		toSerialize["places"] = o.Places
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject1239 struct {
	value *InlineObject1239
	isSet bool
}

func (v NullableInlineObject1239) Get() *InlineObject1239 {
	return v.value
}

func (v *NullableInlineObject1239) Set(val *InlineObject1239) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject1239) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject1239) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject1239(val *InlineObject1239) *NullableInlineObject1239 {
	return &NullableInlineObject1239{value: val, isSet: true}
}

func (v NullableInlineObject1239) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject1239) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


