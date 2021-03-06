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

// InlineObject1506 struct for InlineObject1506
type InlineObject1506 struct {
	Cost AnyOfobject `json:"cost,omitempty"`
	Salvage AnyOfobject `json:"salvage,omitempty"`
	Life AnyOfobject `json:"life,omitempty"`
}

// NewInlineObject1506 instantiates a new InlineObject1506 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject1506() *InlineObject1506 {
	this := InlineObject1506{}
	return &this
}

// NewInlineObject1506WithDefaults instantiates a new InlineObject1506 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject1506WithDefaults() *InlineObject1506 {
	this := InlineObject1506{}
	return &this
}

// GetCost returns the Cost field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1506) GetCost() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Cost
}

// GetCostOk returns a tuple with the Cost field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1506) GetCostOk() (*AnyOfobject, bool) {
	if o == nil || o.Cost == nil {
		return nil, false
	}
	return &o.Cost, true
}

// HasCost returns a boolean if a field has been set.
func (o *InlineObject1506) HasCost() bool {
	if o != nil && o.Cost != nil {
		return true
	}

	return false
}

// SetCost gets a reference to the given AnyOfobject and assigns it to the Cost field.
func (o *InlineObject1506) SetCost(v AnyOfobject) {
	o.Cost = v
}

// GetSalvage returns the Salvage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1506) GetSalvage() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Salvage
}

// GetSalvageOk returns a tuple with the Salvage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1506) GetSalvageOk() (*AnyOfobject, bool) {
	if o == nil || o.Salvage == nil {
		return nil, false
	}
	return &o.Salvage, true
}

// HasSalvage returns a boolean if a field has been set.
func (o *InlineObject1506) HasSalvage() bool {
	if o != nil && o.Salvage != nil {
		return true
	}

	return false
}

// SetSalvage gets a reference to the given AnyOfobject and assigns it to the Salvage field.
func (o *InlineObject1506) SetSalvage(v AnyOfobject) {
	o.Salvage = v
}

// GetLife returns the Life field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1506) GetLife() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Life
}

// GetLifeOk returns a tuple with the Life field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1506) GetLifeOk() (*AnyOfobject, bool) {
	if o == nil || o.Life == nil {
		return nil, false
	}
	return &o.Life, true
}

// HasLife returns a boolean if a field has been set.
func (o *InlineObject1506) HasLife() bool {
	if o != nil && o.Life != nil {
		return true
	}

	return false
}

// SetLife gets a reference to the given AnyOfobject and assigns it to the Life field.
func (o *InlineObject1506) SetLife(v AnyOfobject) {
	o.Life = v
}

func (o InlineObject1506) MarshalJSON() ([]byte, error) {
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
	return json.Marshal(toSerialize)
}

type NullableInlineObject1506 struct {
	value *InlineObject1506
	isSet bool
}

func (v NullableInlineObject1506) Get() *InlineObject1506 {
	return v.value
}

func (v *NullableInlineObject1506) Set(val *InlineObject1506) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject1506) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject1506) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject1506(val *InlineObject1506) *NullableInlineObject1506 {
	return &NullableInlineObject1506{value: val, isSet: true}
}

func (v NullableInlineObject1506) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject1506) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


