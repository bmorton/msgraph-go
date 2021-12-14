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

// InlineObject1307 struct for InlineObject1307
type InlineObject1307 struct {
	Number AnyOfobject `json:"number,omitempty"`
	Decimals AnyOfobject `json:"decimals,omitempty"`
}

// NewInlineObject1307 instantiates a new InlineObject1307 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject1307() *InlineObject1307 {
	this := InlineObject1307{}
	return &this
}

// NewInlineObject1307WithDefaults instantiates a new InlineObject1307 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject1307WithDefaults() *InlineObject1307 {
	this := InlineObject1307{}
	return &this
}

// GetNumber returns the Number field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1307) GetNumber() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Number
}

// GetNumberOk returns a tuple with the Number field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1307) GetNumberOk() (*AnyOfobject, bool) {
	if o == nil || o.Number == nil {
		return nil, false
	}
	return &o.Number, true
}

// HasNumber returns a boolean if a field has been set.
func (o *InlineObject1307) HasNumber() bool {
	if o != nil && o.Number != nil {
		return true
	}

	return false
}

// SetNumber gets a reference to the given AnyOfobject and assigns it to the Number field.
func (o *InlineObject1307) SetNumber(v AnyOfobject) {
	o.Number = v
}

// GetDecimals returns the Decimals field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1307) GetDecimals() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Decimals
}

// GetDecimalsOk returns a tuple with the Decimals field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1307) GetDecimalsOk() (*AnyOfobject, bool) {
	if o == nil || o.Decimals == nil {
		return nil, false
	}
	return &o.Decimals, true
}

// HasDecimals returns a boolean if a field has been set.
func (o *InlineObject1307) HasDecimals() bool {
	if o != nil && o.Decimals != nil {
		return true
	}

	return false
}

// SetDecimals gets a reference to the given AnyOfobject and assigns it to the Decimals field.
func (o *InlineObject1307) SetDecimals(v AnyOfobject) {
	o.Decimals = v
}

func (o InlineObject1307) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Number != nil {
		toSerialize["number"] = o.Number
	}
	if o.Decimals != nil {
		toSerialize["decimals"] = o.Decimals
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject1307 struct {
	value *InlineObject1307
	isSet bool
}

func (v NullableInlineObject1307) Get() *InlineObject1307 {
	return v.value
}

func (v *NullableInlineObject1307) Set(val *InlineObject1307) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject1307) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject1307) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject1307(val *InlineObject1307) *NullableInlineObject1307 {
	return &NullableInlineObject1307{value: val, isSet: true}
}

func (v NullableInlineObject1307) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject1307) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


