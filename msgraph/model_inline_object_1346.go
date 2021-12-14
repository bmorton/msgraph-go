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

// InlineObject1346 struct for InlineObject1346
type InlineObject1346 struct {
	X AnyOfobject `json:"x,omitempty"`
	Alpha AnyOfobject `json:"alpha,omitempty"`
	Beta AnyOfobject `json:"beta,omitempty"`
	Cumulative AnyOfobject `json:"cumulative,omitempty"`
}

// NewInlineObject1346 instantiates a new InlineObject1346 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject1346() *InlineObject1346 {
	this := InlineObject1346{}
	return &this
}

// NewInlineObject1346WithDefaults instantiates a new InlineObject1346 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject1346WithDefaults() *InlineObject1346 {
	this := InlineObject1346{}
	return &this
}

// GetX returns the X field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1346) GetX() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.X
}

// GetXOk returns a tuple with the X field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1346) GetXOk() (*AnyOfobject, bool) {
	if o == nil || o.X == nil {
		return nil, false
	}
	return &o.X, true
}

// HasX returns a boolean if a field has been set.
func (o *InlineObject1346) HasX() bool {
	if o != nil && o.X != nil {
		return true
	}

	return false
}

// SetX gets a reference to the given AnyOfobject and assigns it to the X field.
func (o *InlineObject1346) SetX(v AnyOfobject) {
	o.X = v
}

// GetAlpha returns the Alpha field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1346) GetAlpha() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Alpha
}

// GetAlphaOk returns a tuple with the Alpha field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1346) GetAlphaOk() (*AnyOfobject, bool) {
	if o == nil || o.Alpha == nil {
		return nil, false
	}
	return &o.Alpha, true
}

// HasAlpha returns a boolean if a field has been set.
func (o *InlineObject1346) HasAlpha() bool {
	if o != nil && o.Alpha != nil {
		return true
	}

	return false
}

// SetAlpha gets a reference to the given AnyOfobject and assigns it to the Alpha field.
func (o *InlineObject1346) SetAlpha(v AnyOfobject) {
	o.Alpha = v
}

// GetBeta returns the Beta field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1346) GetBeta() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Beta
}

// GetBetaOk returns a tuple with the Beta field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1346) GetBetaOk() (*AnyOfobject, bool) {
	if o == nil || o.Beta == nil {
		return nil, false
	}
	return &o.Beta, true
}

// HasBeta returns a boolean if a field has been set.
func (o *InlineObject1346) HasBeta() bool {
	if o != nil && o.Beta != nil {
		return true
	}

	return false
}

// SetBeta gets a reference to the given AnyOfobject and assigns it to the Beta field.
func (o *InlineObject1346) SetBeta(v AnyOfobject) {
	o.Beta = v
}

// GetCumulative returns the Cumulative field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1346) GetCumulative() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Cumulative
}

// GetCumulativeOk returns a tuple with the Cumulative field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1346) GetCumulativeOk() (*AnyOfobject, bool) {
	if o == nil || o.Cumulative == nil {
		return nil, false
	}
	return &o.Cumulative, true
}

// HasCumulative returns a boolean if a field has been set.
func (o *InlineObject1346) HasCumulative() bool {
	if o != nil && o.Cumulative != nil {
		return true
	}

	return false
}

// SetCumulative gets a reference to the given AnyOfobject and assigns it to the Cumulative field.
func (o *InlineObject1346) SetCumulative(v AnyOfobject) {
	o.Cumulative = v
}

func (o InlineObject1346) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.X != nil {
		toSerialize["x"] = o.X
	}
	if o.Alpha != nil {
		toSerialize["alpha"] = o.Alpha
	}
	if o.Beta != nil {
		toSerialize["beta"] = o.Beta
	}
	if o.Cumulative != nil {
		toSerialize["cumulative"] = o.Cumulative
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject1346 struct {
	value *InlineObject1346
	isSet bool
}

func (v NullableInlineObject1346) Get() *InlineObject1346 {
	return v.value
}

func (v *NullableInlineObject1346) Set(val *InlineObject1346) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject1346) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject1346) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject1346(val *InlineObject1346) *NullableInlineObject1346 {
	return &NullableInlineObject1346{value: val, isSet: true}
}

func (v NullableInlineObject1346) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject1346) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


