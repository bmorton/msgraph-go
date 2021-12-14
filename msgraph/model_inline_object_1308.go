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

// InlineObject1308 struct for InlineObject1308
type InlineObject1308 struct {
	FractionalDollar AnyOfobject `json:"fractionalDollar,omitempty"`
	Fraction AnyOfobject `json:"fraction,omitempty"`
}

// NewInlineObject1308 instantiates a new InlineObject1308 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject1308() *InlineObject1308 {
	this := InlineObject1308{}
	return &this
}

// NewInlineObject1308WithDefaults instantiates a new InlineObject1308 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject1308WithDefaults() *InlineObject1308 {
	this := InlineObject1308{}
	return &this
}

// GetFractionalDollar returns the FractionalDollar field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1308) GetFractionalDollar() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.FractionalDollar
}

// GetFractionalDollarOk returns a tuple with the FractionalDollar field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1308) GetFractionalDollarOk() (*AnyOfobject, bool) {
	if o == nil || o.FractionalDollar == nil {
		return nil, false
	}
	return &o.FractionalDollar, true
}

// HasFractionalDollar returns a boolean if a field has been set.
func (o *InlineObject1308) HasFractionalDollar() bool {
	if o != nil && o.FractionalDollar != nil {
		return true
	}

	return false
}

// SetFractionalDollar gets a reference to the given AnyOfobject and assigns it to the FractionalDollar field.
func (o *InlineObject1308) SetFractionalDollar(v AnyOfobject) {
	o.FractionalDollar = v
}

// GetFraction returns the Fraction field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1308) GetFraction() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Fraction
}

// GetFractionOk returns a tuple with the Fraction field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1308) GetFractionOk() (*AnyOfobject, bool) {
	if o == nil || o.Fraction == nil {
		return nil, false
	}
	return &o.Fraction, true
}

// HasFraction returns a boolean if a field has been set.
func (o *InlineObject1308) HasFraction() bool {
	if o != nil && o.Fraction != nil {
		return true
	}

	return false
}

// SetFraction gets a reference to the given AnyOfobject and assigns it to the Fraction field.
func (o *InlineObject1308) SetFraction(v AnyOfobject) {
	o.Fraction = v
}

func (o InlineObject1308) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.FractionalDollar != nil {
		toSerialize["fractionalDollar"] = o.FractionalDollar
	}
	if o.Fraction != nil {
		toSerialize["fraction"] = o.Fraction
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject1308 struct {
	value *InlineObject1308
	isSet bool
}

func (v NullableInlineObject1308) Get() *InlineObject1308 {
	return v.value
}

func (v *NullableInlineObject1308) Set(val *InlineObject1308) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject1308) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject1308) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject1308(val *InlineObject1308) *NullableInlineObject1308 {
	return &NullableInlineObject1308{value: val, isSet: true}
}

func (v NullableInlineObject1308) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject1308) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

