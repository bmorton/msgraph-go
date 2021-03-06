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

// InlineObject1238 struct for InlineObject1238
type InlineObject1238 struct {
	Number AnyOfobject `json:"number,omitempty"`
	Places AnyOfobject `json:"places,omitempty"`
}

// NewInlineObject1238 instantiates a new InlineObject1238 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject1238() *InlineObject1238 {
	this := InlineObject1238{}
	return &this
}

// NewInlineObject1238WithDefaults instantiates a new InlineObject1238 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject1238WithDefaults() *InlineObject1238 {
	this := InlineObject1238{}
	return &this
}

// GetNumber returns the Number field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1238) GetNumber() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Number
}

// GetNumberOk returns a tuple with the Number field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1238) GetNumberOk() (*AnyOfobject, bool) {
	if o == nil || o.Number == nil {
		return nil, false
	}
	return &o.Number, true
}

// HasNumber returns a boolean if a field has been set.
func (o *InlineObject1238) HasNumber() bool {
	if o != nil && o.Number != nil {
		return true
	}

	return false
}

// SetNumber gets a reference to the given AnyOfobject and assigns it to the Number field.
func (o *InlineObject1238) SetNumber(v AnyOfobject) {
	o.Number = v
}

// GetPlaces returns the Places field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1238) GetPlaces() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Places
}

// GetPlacesOk returns a tuple with the Places field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1238) GetPlacesOk() (*AnyOfobject, bool) {
	if o == nil || o.Places == nil {
		return nil, false
	}
	return &o.Places, true
}

// HasPlaces returns a boolean if a field has been set.
func (o *InlineObject1238) HasPlaces() bool {
	if o != nil && o.Places != nil {
		return true
	}

	return false
}

// SetPlaces gets a reference to the given AnyOfobject and assigns it to the Places field.
func (o *InlineObject1238) SetPlaces(v AnyOfobject) {
	o.Places = v
}

func (o InlineObject1238) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Number != nil {
		toSerialize["number"] = o.Number
	}
	if o.Places != nil {
		toSerialize["places"] = o.Places
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject1238 struct {
	value *InlineObject1238
	isSet bool
}

func (v NullableInlineObject1238) Get() *InlineObject1238 {
	return v.value
}

func (v *NullableInlineObject1238) Set(val *InlineObject1238) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject1238) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject1238) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject1238(val *InlineObject1238) *NullableInlineObject1238 {
	return &NullableInlineObject1238{value: val, isSet: true}
}

func (v NullableInlineObject1238) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject1238) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


