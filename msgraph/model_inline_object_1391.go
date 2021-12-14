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

// InlineObject1391 struct for InlineObject1391
type InlineObject1391 struct {
	Values AnyOfobject `json:"values,omitempty"`
	Guess AnyOfobject `json:"guess,omitempty"`
}

// NewInlineObject1391 instantiates a new InlineObject1391 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject1391() *InlineObject1391 {
	this := InlineObject1391{}
	return &this
}

// NewInlineObject1391WithDefaults instantiates a new InlineObject1391 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject1391WithDefaults() *InlineObject1391 {
	this := InlineObject1391{}
	return &this
}

// GetValues returns the Values field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1391) GetValues() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1391) GetValuesOk() (*AnyOfobject, bool) {
	if o == nil || o.Values == nil {
		return nil, false
	}
	return &o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *InlineObject1391) HasValues() bool {
	if o != nil && o.Values != nil {
		return true
	}

	return false
}

// SetValues gets a reference to the given AnyOfobject and assigns it to the Values field.
func (o *InlineObject1391) SetValues(v AnyOfobject) {
	o.Values = v
}

// GetGuess returns the Guess field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1391) GetGuess() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Guess
}

// GetGuessOk returns a tuple with the Guess field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1391) GetGuessOk() (*AnyOfobject, bool) {
	if o == nil || o.Guess == nil {
		return nil, false
	}
	return &o.Guess, true
}

// HasGuess returns a boolean if a field has been set.
func (o *InlineObject1391) HasGuess() bool {
	if o != nil && o.Guess != nil {
		return true
	}

	return false
}

// SetGuess gets a reference to the given AnyOfobject and assigns it to the Guess field.
func (o *InlineObject1391) SetGuess(v AnyOfobject) {
	o.Guess = v
}

func (o InlineObject1391) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Values != nil {
		toSerialize["values"] = o.Values
	}
	if o.Guess != nil {
		toSerialize["guess"] = o.Guess
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject1391 struct {
	value *InlineObject1391
	isSet bool
}

func (v NullableInlineObject1391) Get() *InlineObject1391 {
	return v.value
}

func (v *NullableInlineObject1391) Set(val *InlineObject1391) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject1391) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject1391) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject1391(val *InlineObject1391) *NullableInlineObject1391 {
	return &NullableInlineObject1391{value: val, isSet: true}
}

func (v NullableInlineObject1391) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject1391) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


