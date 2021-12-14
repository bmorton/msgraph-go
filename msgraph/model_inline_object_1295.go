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

// InlineObject1295 struct for InlineObject1295
type InlineObject1295 struct {
	Cost AnyOfobject `json:"cost,omitempty"`
	Salvage AnyOfobject `json:"salvage,omitempty"`
	Life AnyOfobject `json:"life,omitempty"`
	Period AnyOfobject `json:"period,omitempty"`
	Factor AnyOfobject `json:"factor,omitempty"`
}

// NewInlineObject1295 instantiates a new InlineObject1295 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject1295() *InlineObject1295 {
	this := InlineObject1295{}
	return &this
}

// NewInlineObject1295WithDefaults instantiates a new InlineObject1295 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject1295WithDefaults() *InlineObject1295 {
	this := InlineObject1295{}
	return &this
}

// GetCost returns the Cost field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1295) GetCost() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Cost
}

// GetCostOk returns a tuple with the Cost field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1295) GetCostOk() (*AnyOfobject, bool) {
	if o == nil || o.Cost == nil {
		return nil, false
	}
	return &o.Cost, true
}

// HasCost returns a boolean if a field has been set.
func (o *InlineObject1295) HasCost() bool {
	if o != nil && o.Cost != nil {
		return true
	}

	return false
}

// SetCost gets a reference to the given AnyOfobject and assigns it to the Cost field.
func (o *InlineObject1295) SetCost(v AnyOfobject) {
	o.Cost = v
}

// GetSalvage returns the Salvage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1295) GetSalvage() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Salvage
}

// GetSalvageOk returns a tuple with the Salvage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1295) GetSalvageOk() (*AnyOfobject, bool) {
	if o == nil || o.Salvage == nil {
		return nil, false
	}
	return &o.Salvage, true
}

// HasSalvage returns a boolean if a field has been set.
func (o *InlineObject1295) HasSalvage() bool {
	if o != nil && o.Salvage != nil {
		return true
	}

	return false
}

// SetSalvage gets a reference to the given AnyOfobject and assigns it to the Salvage field.
func (o *InlineObject1295) SetSalvage(v AnyOfobject) {
	o.Salvage = v
}

// GetLife returns the Life field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1295) GetLife() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Life
}

// GetLifeOk returns a tuple with the Life field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1295) GetLifeOk() (*AnyOfobject, bool) {
	if o == nil || o.Life == nil {
		return nil, false
	}
	return &o.Life, true
}

// HasLife returns a boolean if a field has been set.
func (o *InlineObject1295) HasLife() bool {
	if o != nil && o.Life != nil {
		return true
	}

	return false
}

// SetLife gets a reference to the given AnyOfobject and assigns it to the Life field.
func (o *InlineObject1295) SetLife(v AnyOfobject) {
	o.Life = v
}

// GetPeriod returns the Period field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1295) GetPeriod() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Period
}

// GetPeriodOk returns a tuple with the Period field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1295) GetPeriodOk() (*AnyOfobject, bool) {
	if o == nil || o.Period == nil {
		return nil, false
	}
	return &o.Period, true
}

// HasPeriod returns a boolean if a field has been set.
func (o *InlineObject1295) HasPeriod() bool {
	if o != nil && o.Period != nil {
		return true
	}

	return false
}

// SetPeriod gets a reference to the given AnyOfobject and assigns it to the Period field.
func (o *InlineObject1295) SetPeriod(v AnyOfobject) {
	o.Period = v
}

// GetFactor returns the Factor field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1295) GetFactor() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Factor
}

// GetFactorOk returns a tuple with the Factor field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1295) GetFactorOk() (*AnyOfobject, bool) {
	if o == nil || o.Factor == nil {
		return nil, false
	}
	return &o.Factor, true
}

// HasFactor returns a boolean if a field has been set.
func (o *InlineObject1295) HasFactor() bool {
	if o != nil && o.Factor != nil {
		return true
	}

	return false
}

// SetFactor gets a reference to the given AnyOfobject and assigns it to the Factor field.
func (o *InlineObject1295) SetFactor(v AnyOfobject) {
	o.Factor = v
}

func (o InlineObject1295) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Cost != nil {
		toSerialize["cost"] = o.Cost
	}
	if o.Salvage != nil {
		toSerialize["salvage"] = o.Salvage
	}
	if o.Life != nil {
		toSerialize["life"] = o.Life
	}
	if o.Period != nil {
		toSerialize["period"] = o.Period
	}
	if o.Factor != nil {
		toSerialize["factor"] = o.Factor
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject1295 struct {
	value *InlineObject1295
	isSet bool
}

func (v NullableInlineObject1295) Get() *InlineObject1295 {
	return v.value
}

func (v *NullableInlineObject1295) Set(val *InlineObject1295) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject1295) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject1295) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject1295(val *InlineObject1295) *NullableInlineObject1295 {
	return &NullableInlineObject1295{value: val, isSet: true}
}

func (v NullableInlineObject1295) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject1295) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


