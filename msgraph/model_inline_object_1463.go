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

// InlineObject1463 struct for InlineObject1463
type InlineObject1463 struct {
	Number AnyOfobject `json:"number,omitempty"`
	NumberChosen AnyOfobject `json:"numberChosen,omitempty"`
}

// NewInlineObject1463 instantiates a new InlineObject1463 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject1463() *InlineObject1463 {
	this := InlineObject1463{}
	return &this
}

// NewInlineObject1463WithDefaults instantiates a new InlineObject1463 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject1463WithDefaults() *InlineObject1463 {
	this := InlineObject1463{}
	return &this
}

// GetNumber returns the Number field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1463) GetNumber() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Number
}

// GetNumberOk returns a tuple with the Number field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1463) GetNumberOk() (*AnyOfobject, bool) {
	if o == nil || o.Number == nil {
		return nil, false
	}
	return &o.Number, true
}

// HasNumber returns a boolean if a field has been set.
func (o *InlineObject1463) HasNumber() bool {
	if o != nil && o.Number != nil {
		return true
	}

	return false
}

// SetNumber gets a reference to the given AnyOfobject and assigns it to the Number field.
func (o *InlineObject1463) SetNumber(v AnyOfobject) {
	o.Number = v
}

// GetNumberChosen returns the NumberChosen field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1463) GetNumberChosen() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.NumberChosen
}

// GetNumberChosenOk returns a tuple with the NumberChosen field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1463) GetNumberChosenOk() (*AnyOfobject, bool) {
	if o == nil || o.NumberChosen == nil {
		return nil, false
	}
	return &o.NumberChosen, true
}

// HasNumberChosen returns a boolean if a field has been set.
func (o *InlineObject1463) HasNumberChosen() bool {
	if o != nil && o.NumberChosen != nil {
		return true
	}

	return false
}

// SetNumberChosen gets a reference to the given AnyOfobject and assigns it to the NumberChosen field.
func (o *InlineObject1463) SetNumberChosen(v AnyOfobject) {
	o.NumberChosen = v
}

func (o InlineObject1463) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Number != nil {
		toSerialize["number"] = o.Number
	}
	if o.NumberChosen != nil {
		toSerialize["numberChosen"] = o.NumberChosen
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject1463 struct {
	value *InlineObject1463
	isSet bool
}

func (v NullableInlineObject1463) Get() *InlineObject1463 {
	return v.value
}

func (v *NullableInlineObject1463) Set(val *InlineObject1463) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject1463) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject1463) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject1463(val *InlineObject1463) *NullableInlineObject1463 {
	return &NullableInlineObject1463{value: val, isSet: true}
}

func (v NullableInlineObject1463) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject1463) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


