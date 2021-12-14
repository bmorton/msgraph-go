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

// InlineObject1465 struct for InlineObject1465
type InlineObject1465 struct {
	Rate AnyOfobject `json:"rate,omitempty"`
	Nper AnyOfobject `json:"nper,omitempty"`
	Pv AnyOfobject `json:"pv,omitempty"`
	Fv AnyOfobject `json:"fv,omitempty"`
	Type AnyOfobject `json:"type,omitempty"`
}

// NewInlineObject1465 instantiates a new InlineObject1465 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject1465() *InlineObject1465 {
	this := InlineObject1465{}
	return &this
}

// NewInlineObject1465WithDefaults instantiates a new InlineObject1465 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject1465WithDefaults() *InlineObject1465 {
	this := InlineObject1465{}
	return &this
}

// GetRate returns the Rate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1465) GetRate() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Rate
}

// GetRateOk returns a tuple with the Rate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1465) GetRateOk() (*AnyOfobject, bool) {
	if o == nil || o.Rate == nil {
		return nil, false
	}
	return &o.Rate, true
}

// HasRate returns a boolean if a field has been set.
func (o *InlineObject1465) HasRate() bool {
	if o != nil && o.Rate != nil {
		return true
	}

	return false
}

// SetRate gets a reference to the given AnyOfobject and assigns it to the Rate field.
func (o *InlineObject1465) SetRate(v AnyOfobject) {
	o.Rate = v
}

// GetNper returns the Nper field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1465) GetNper() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Nper
}

// GetNperOk returns a tuple with the Nper field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1465) GetNperOk() (*AnyOfobject, bool) {
	if o == nil || o.Nper == nil {
		return nil, false
	}
	return &o.Nper, true
}

// HasNper returns a boolean if a field has been set.
func (o *InlineObject1465) HasNper() bool {
	if o != nil && o.Nper != nil {
		return true
	}

	return false
}

// SetNper gets a reference to the given AnyOfobject and assigns it to the Nper field.
func (o *InlineObject1465) SetNper(v AnyOfobject) {
	o.Nper = v
}

// GetPv returns the Pv field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1465) GetPv() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Pv
}

// GetPvOk returns a tuple with the Pv field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1465) GetPvOk() (*AnyOfobject, bool) {
	if o == nil || o.Pv == nil {
		return nil, false
	}
	return &o.Pv, true
}

// HasPv returns a boolean if a field has been set.
func (o *InlineObject1465) HasPv() bool {
	if o != nil && o.Pv != nil {
		return true
	}

	return false
}

// SetPv gets a reference to the given AnyOfobject and assigns it to the Pv field.
func (o *InlineObject1465) SetPv(v AnyOfobject) {
	o.Pv = v
}

// GetFv returns the Fv field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1465) GetFv() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Fv
}

// GetFvOk returns a tuple with the Fv field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1465) GetFvOk() (*AnyOfobject, bool) {
	if o == nil || o.Fv == nil {
		return nil, false
	}
	return &o.Fv, true
}

// HasFv returns a boolean if a field has been set.
func (o *InlineObject1465) HasFv() bool {
	if o != nil && o.Fv != nil {
		return true
	}

	return false
}

// SetFv gets a reference to the given AnyOfobject and assigns it to the Fv field.
func (o *InlineObject1465) SetFv(v AnyOfobject) {
	o.Fv = v
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1465) GetType() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1465) GetTypeOk() (*AnyOfobject, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return &o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *InlineObject1465) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given AnyOfobject and assigns it to the Type field.
func (o *InlineObject1465) SetType(v AnyOfobject) {
	o.Type = v
}

func (o InlineObject1465) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Rate != nil {
		toSerialize["rate"] = o.Rate
	}
	if o.Nper != nil {
		toSerialize["nper"] = o.Nper
	}
	if o.Pv != nil {
		toSerialize["pv"] = o.Pv
	}
	if o.Fv != nil {
		toSerialize["fv"] = o.Fv
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject1465 struct {
	value *InlineObject1465
	isSet bool
}

func (v NullableInlineObject1465) Get() *InlineObject1465 {
	return v.value
}

func (v *NullableInlineObject1465) Set(val *InlineObject1465) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject1465) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject1465) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject1465(val *InlineObject1465) *NullableInlineObject1465 {
	return &NullableInlineObject1465{value: val, isSet: true}
}

func (v NullableInlineObject1465) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject1465) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

