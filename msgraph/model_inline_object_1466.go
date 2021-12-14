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

// InlineObject1466 struct for InlineObject1466
type InlineObject1466 struct {
	X AnyOfobject `json:"x,omitempty"`
	Mean AnyOfobject `json:"mean,omitempty"`
	Cumulative AnyOfobject `json:"cumulative,omitempty"`
}

// NewInlineObject1466 instantiates a new InlineObject1466 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject1466() *InlineObject1466 {
	this := InlineObject1466{}
	return &this
}

// NewInlineObject1466WithDefaults instantiates a new InlineObject1466 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject1466WithDefaults() *InlineObject1466 {
	this := InlineObject1466{}
	return &this
}

// GetX returns the X field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1466) GetX() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.X
}

// GetXOk returns a tuple with the X field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1466) GetXOk() (*AnyOfobject, bool) {
	if o == nil || o.X == nil {
		return nil, false
	}
	return &o.X, true
}

// HasX returns a boolean if a field has been set.
func (o *InlineObject1466) HasX() bool {
	if o != nil && o.X != nil {
		return true
	}

	return false
}

// SetX gets a reference to the given AnyOfobject and assigns it to the X field.
func (o *InlineObject1466) SetX(v AnyOfobject) {
	o.X = v
}

// GetMean returns the Mean field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1466) GetMean() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Mean
}

// GetMeanOk returns a tuple with the Mean field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1466) GetMeanOk() (*AnyOfobject, bool) {
	if o == nil || o.Mean == nil {
		return nil, false
	}
	return &o.Mean, true
}

// HasMean returns a boolean if a field has been set.
func (o *InlineObject1466) HasMean() bool {
	if o != nil && o.Mean != nil {
		return true
	}

	return false
}

// SetMean gets a reference to the given AnyOfobject and assigns it to the Mean field.
func (o *InlineObject1466) SetMean(v AnyOfobject) {
	o.Mean = v
}

// GetCumulative returns the Cumulative field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1466) GetCumulative() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Cumulative
}

// GetCumulativeOk returns a tuple with the Cumulative field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1466) GetCumulativeOk() (*AnyOfobject, bool) {
	if o == nil || o.Cumulative == nil {
		return nil, false
	}
	return &o.Cumulative, true
}

// HasCumulative returns a boolean if a field has been set.
func (o *InlineObject1466) HasCumulative() bool {
	if o != nil && o.Cumulative != nil {
		return true
	}

	return false
}

// SetCumulative gets a reference to the given AnyOfobject and assigns it to the Cumulative field.
func (o *InlineObject1466) SetCumulative(v AnyOfobject) {
	o.Cumulative = v
}

func (o InlineObject1466) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.X != nil {
		toSerialize["x"] = o.X
	}
	if o.Mean != nil {
		toSerialize["mean"] = o.Mean
	}
	if o.Cumulative != nil {
		toSerialize["cumulative"] = o.Cumulative
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject1466 struct {
	value *InlineObject1466
	isSet bool
}

func (v NullableInlineObject1466) Get() *InlineObject1466 {
	return v.value
}

func (v *NullableInlineObject1466) Set(val *InlineObject1466) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject1466) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject1466) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject1466(val *InlineObject1466) *NullableInlineObject1466 {
	return &NullableInlineObject1466{value: val, isSet: true}
}

func (v NullableInlineObject1466) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject1466) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

