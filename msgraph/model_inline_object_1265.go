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

// InlineObject1265 struct for InlineObject1265
type InlineObject1265 struct {
	Number AnyOfobject `json:"number,omitempty"`
	FromUnit AnyOfobject `json:"fromUnit,omitempty"`
	ToUnit AnyOfobject `json:"toUnit,omitempty"`
}

// NewInlineObject1265 instantiates a new InlineObject1265 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject1265() *InlineObject1265 {
	this := InlineObject1265{}
	return &this
}

// NewInlineObject1265WithDefaults instantiates a new InlineObject1265 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject1265WithDefaults() *InlineObject1265 {
	this := InlineObject1265{}
	return &this
}

// GetNumber returns the Number field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1265) GetNumber() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Number
}

// GetNumberOk returns a tuple with the Number field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1265) GetNumberOk() (*AnyOfobject, bool) {
	if o == nil || o.Number == nil {
		return nil, false
	}
	return &o.Number, true
}

// HasNumber returns a boolean if a field has been set.
func (o *InlineObject1265) HasNumber() bool {
	if o != nil && o.Number != nil {
		return true
	}

	return false
}

// SetNumber gets a reference to the given AnyOfobject and assigns it to the Number field.
func (o *InlineObject1265) SetNumber(v AnyOfobject) {
	o.Number = v
}

// GetFromUnit returns the FromUnit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1265) GetFromUnit() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.FromUnit
}

// GetFromUnitOk returns a tuple with the FromUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1265) GetFromUnitOk() (*AnyOfobject, bool) {
	if o == nil || o.FromUnit == nil {
		return nil, false
	}
	return &o.FromUnit, true
}

// HasFromUnit returns a boolean if a field has been set.
func (o *InlineObject1265) HasFromUnit() bool {
	if o != nil && o.FromUnit != nil {
		return true
	}

	return false
}

// SetFromUnit gets a reference to the given AnyOfobject and assigns it to the FromUnit field.
func (o *InlineObject1265) SetFromUnit(v AnyOfobject) {
	o.FromUnit = v
}

// GetToUnit returns the ToUnit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1265) GetToUnit() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.ToUnit
}

// GetToUnitOk returns a tuple with the ToUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1265) GetToUnitOk() (*AnyOfobject, bool) {
	if o == nil || o.ToUnit == nil {
		return nil, false
	}
	return &o.ToUnit, true
}

// HasToUnit returns a boolean if a field has been set.
func (o *InlineObject1265) HasToUnit() bool {
	if o != nil && o.ToUnit != nil {
		return true
	}

	return false
}

// SetToUnit gets a reference to the given AnyOfobject and assigns it to the ToUnit field.
func (o *InlineObject1265) SetToUnit(v AnyOfobject) {
	o.ToUnit = v
}

func (o InlineObject1265) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Number != nil {
		toSerialize["number"] = o.Number
	}
	if o.FromUnit != nil {
		toSerialize["fromUnit"] = o.FromUnit
	}
	if o.ToUnit != nil {
		toSerialize["toUnit"] = o.ToUnit
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject1265 struct {
	value *InlineObject1265
	isSet bool
}

func (v NullableInlineObject1265) Get() *InlineObject1265 {
	return v.value
}

func (v *NullableInlineObject1265) Set(val *InlineObject1265) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject1265) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject1265) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject1265(val *InlineObject1265) *NullableInlineObject1265 {
	return &NullableInlineObject1265{value: val, isSet: true}
}

func (v NullableInlineObject1265) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject1265) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


