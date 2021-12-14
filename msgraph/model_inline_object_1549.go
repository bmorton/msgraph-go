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

// InlineObject1549 struct for InlineObject1549
type InlineObject1549 struct {
	Cost AnyOfobject `json:"cost,omitempty"`
	Salvage AnyOfobject `json:"salvage,omitempty"`
	Life AnyOfobject `json:"life,omitempty"`
	StartPeriod AnyOfobject `json:"startPeriod,omitempty"`
	EndPeriod AnyOfobject `json:"endPeriod,omitempty"`
	Factor AnyOfobject `json:"factor,omitempty"`
	NoSwitch AnyOfobject `json:"noSwitch,omitempty"`
}

// NewInlineObject1549 instantiates a new InlineObject1549 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject1549() *InlineObject1549 {
	this := InlineObject1549{}
	return &this
}

// NewInlineObject1549WithDefaults instantiates a new InlineObject1549 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject1549WithDefaults() *InlineObject1549 {
	this := InlineObject1549{}
	return &this
}

// GetCost returns the Cost field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1549) GetCost() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Cost
}

// GetCostOk returns a tuple with the Cost field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1549) GetCostOk() (*AnyOfobject, bool) {
	if o == nil || o.Cost == nil {
		return nil, false
	}
	return &o.Cost, true
}

// HasCost returns a boolean if a field has been set.
func (o *InlineObject1549) HasCost() bool {
	if o != nil && o.Cost != nil {
		return true
	}

	return false
}

// SetCost gets a reference to the given AnyOfobject and assigns it to the Cost field.
func (o *InlineObject1549) SetCost(v AnyOfobject) {
	o.Cost = v
}

// GetSalvage returns the Salvage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1549) GetSalvage() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Salvage
}

// GetSalvageOk returns a tuple with the Salvage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1549) GetSalvageOk() (*AnyOfobject, bool) {
	if o == nil || o.Salvage == nil {
		return nil, false
	}
	return &o.Salvage, true
}

// HasSalvage returns a boolean if a field has been set.
func (o *InlineObject1549) HasSalvage() bool {
	if o != nil && o.Salvage != nil {
		return true
	}

	return false
}

// SetSalvage gets a reference to the given AnyOfobject and assigns it to the Salvage field.
func (o *InlineObject1549) SetSalvage(v AnyOfobject) {
	o.Salvage = v
}

// GetLife returns the Life field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1549) GetLife() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Life
}

// GetLifeOk returns a tuple with the Life field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1549) GetLifeOk() (*AnyOfobject, bool) {
	if o == nil || o.Life == nil {
		return nil, false
	}
	return &o.Life, true
}

// HasLife returns a boolean if a field has been set.
func (o *InlineObject1549) HasLife() bool {
	if o != nil && o.Life != nil {
		return true
	}

	return false
}

// SetLife gets a reference to the given AnyOfobject and assigns it to the Life field.
func (o *InlineObject1549) SetLife(v AnyOfobject) {
	o.Life = v
}

// GetStartPeriod returns the StartPeriod field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1549) GetStartPeriod() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.StartPeriod
}

// GetStartPeriodOk returns a tuple with the StartPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1549) GetStartPeriodOk() (*AnyOfobject, bool) {
	if o == nil || o.StartPeriod == nil {
		return nil, false
	}
	return &o.StartPeriod, true
}

// HasStartPeriod returns a boolean if a field has been set.
func (o *InlineObject1549) HasStartPeriod() bool {
	if o != nil && o.StartPeriod != nil {
		return true
	}

	return false
}

// SetStartPeriod gets a reference to the given AnyOfobject and assigns it to the StartPeriod field.
func (o *InlineObject1549) SetStartPeriod(v AnyOfobject) {
	o.StartPeriod = v
}

// GetEndPeriod returns the EndPeriod field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1549) GetEndPeriod() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.EndPeriod
}

// GetEndPeriodOk returns a tuple with the EndPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1549) GetEndPeriodOk() (*AnyOfobject, bool) {
	if o == nil || o.EndPeriod == nil {
		return nil, false
	}
	return &o.EndPeriod, true
}

// HasEndPeriod returns a boolean if a field has been set.
func (o *InlineObject1549) HasEndPeriod() bool {
	if o != nil && o.EndPeriod != nil {
		return true
	}

	return false
}

// SetEndPeriod gets a reference to the given AnyOfobject and assigns it to the EndPeriod field.
func (o *InlineObject1549) SetEndPeriod(v AnyOfobject) {
	o.EndPeriod = v
}

// GetFactor returns the Factor field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1549) GetFactor() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Factor
}

// GetFactorOk returns a tuple with the Factor field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1549) GetFactorOk() (*AnyOfobject, bool) {
	if o == nil || o.Factor == nil {
		return nil, false
	}
	return &o.Factor, true
}

// HasFactor returns a boolean if a field has been set.
func (o *InlineObject1549) HasFactor() bool {
	if o != nil && o.Factor != nil {
		return true
	}

	return false
}

// SetFactor gets a reference to the given AnyOfobject and assigns it to the Factor field.
func (o *InlineObject1549) SetFactor(v AnyOfobject) {
	o.Factor = v
}

// GetNoSwitch returns the NoSwitch field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1549) GetNoSwitch() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.NoSwitch
}

// GetNoSwitchOk returns a tuple with the NoSwitch field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1549) GetNoSwitchOk() (*AnyOfobject, bool) {
	if o == nil || o.NoSwitch == nil {
		return nil, false
	}
	return &o.NoSwitch, true
}

// HasNoSwitch returns a boolean if a field has been set.
func (o *InlineObject1549) HasNoSwitch() bool {
	if o != nil && o.NoSwitch != nil {
		return true
	}

	return false
}

// SetNoSwitch gets a reference to the given AnyOfobject and assigns it to the NoSwitch field.
func (o *InlineObject1549) SetNoSwitch(v AnyOfobject) {
	o.NoSwitch = v
}

func (o InlineObject1549) MarshalJSON() ([]byte, error) {
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
	if o.StartPeriod != nil {
		toSerialize["startPeriod"] = o.StartPeriod
	}
	if o.EndPeriod != nil {
		toSerialize["endPeriod"] = o.EndPeriod
	}
	if o.Factor != nil {
		toSerialize["factor"] = o.Factor
	}
	if o.NoSwitch != nil {
		toSerialize["noSwitch"] = o.NoSwitch
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject1549 struct {
	value *InlineObject1549
	isSet bool
}

func (v NullableInlineObject1549) Get() *InlineObject1549 {
	return v.value
}

func (v *NullableInlineObject1549) Set(val *InlineObject1549) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject1549) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject1549) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject1549(val *InlineObject1549) *NullableInlineObject1549 {
	return &NullableInlineObject1549{value: val, isSet: true}
}

func (v NullableInlineObject1549) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject1549) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


