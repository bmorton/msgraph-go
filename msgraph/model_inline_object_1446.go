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

// InlineObject1446 struct for InlineObject1446
type InlineObject1446 struct {
	Rate AnyOfobject `json:"rate,omitempty"`
	Values AnyOfobject `json:"values,omitempty"`
}

// NewInlineObject1446 instantiates a new InlineObject1446 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject1446() *InlineObject1446 {
	this := InlineObject1446{}
	return &this
}

// NewInlineObject1446WithDefaults instantiates a new InlineObject1446 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject1446WithDefaults() *InlineObject1446 {
	this := InlineObject1446{}
	return &this
}

// GetRate returns the Rate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1446) GetRate() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Rate
}

// GetRateOk returns a tuple with the Rate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1446) GetRateOk() (*AnyOfobject, bool) {
	if o == nil || o.Rate == nil {
		return nil, false
	}
	return &o.Rate, true
}

// HasRate returns a boolean if a field has been set.
func (o *InlineObject1446) HasRate() bool {
	if o != nil && o.Rate != nil {
		return true
	}

	return false
}

// SetRate gets a reference to the given AnyOfobject and assigns it to the Rate field.
func (o *InlineObject1446) SetRate(v AnyOfobject) {
	o.Rate = v
}

// GetValues returns the Values field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1446) GetValues() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1446) GetValuesOk() (*AnyOfobject, bool) {
	if o == nil || o.Values == nil {
		return nil, false
	}
	return &o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *InlineObject1446) HasValues() bool {
	if o != nil && o.Values != nil {
		return true
	}

	return false
}

// SetValues gets a reference to the given AnyOfobject and assigns it to the Values field.
func (o *InlineObject1446) SetValues(v AnyOfobject) {
	o.Values = v
}

func (o InlineObject1446) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Rate != nil {
		toSerialize["rate"] = o.Rate
	}
	if o.Values != nil {
		toSerialize["values"] = o.Values
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject1446 struct {
	value *InlineObject1446
	isSet bool
}

func (v NullableInlineObject1446) Get() *InlineObject1446 {
	return v.value
}

func (v *NullableInlineObject1446) Set(val *InlineObject1446) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject1446) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject1446) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject1446(val *InlineObject1446) *NullableInlineObject1446 {
	return &NullableInlineObject1446{value: val, isSet: true}
}

func (v NullableInlineObject1446) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject1446) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


