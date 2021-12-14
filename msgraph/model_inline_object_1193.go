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

// InlineObject1193 struct for InlineObject1193
type InlineObject1193 struct {
	IsSyncedFromOnPremises NullableBool `json:"isSyncedFromOnPremises,omitempty"`
}

// NewInlineObject1193 instantiates a new InlineObject1193 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject1193() *InlineObject1193 {
	this := InlineObject1193{}
	var isSyncedFromOnPremises bool = false
	this.IsSyncedFromOnPremises = *NewNullableBool(&isSyncedFromOnPremises)
	return &this
}

// NewInlineObject1193WithDefaults instantiates a new InlineObject1193 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject1193WithDefaults() *InlineObject1193 {
	this := InlineObject1193{}
	var isSyncedFromOnPremises bool = false
	this.IsSyncedFromOnPremises = *NewNullableBool(&isSyncedFromOnPremises)
	return &this
}

// GetIsSyncedFromOnPremises returns the IsSyncedFromOnPremises field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1193) GetIsSyncedFromOnPremises() bool {
	if o == nil || o.IsSyncedFromOnPremises.Get() == nil {
		var ret bool
		return ret
	}
	return *o.IsSyncedFromOnPremises.Get()
}

// GetIsSyncedFromOnPremisesOk returns a tuple with the IsSyncedFromOnPremises field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1193) GetIsSyncedFromOnPremisesOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IsSyncedFromOnPremises.Get(), o.IsSyncedFromOnPremises.IsSet()
}

// HasIsSyncedFromOnPremises returns a boolean if a field has been set.
func (o *InlineObject1193) HasIsSyncedFromOnPremises() bool {
	if o != nil && o.IsSyncedFromOnPremises.IsSet() {
		return true
	}

	return false
}

// SetIsSyncedFromOnPremises gets a reference to the given NullableBool and assigns it to the IsSyncedFromOnPremises field.
func (o *InlineObject1193) SetIsSyncedFromOnPremises(v bool) {
	o.IsSyncedFromOnPremises.Set(&v)
}
// SetIsSyncedFromOnPremisesNil sets the value for IsSyncedFromOnPremises to be an explicit nil
func (o *InlineObject1193) SetIsSyncedFromOnPremisesNil() {
	o.IsSyncedFromOnPremises.Set(nil)
}

// UnsetIsSyncedFromOnPremises ensures that no value is present for IsSyncedFromOnPremises, not even an explicit nil
func (o *InlineObject1193) UnsetIsSyncedFromOnPremises() {
	o.IsSyncedFromOnPremises.Unset()
}

func (o InlineObject1193) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.IsSyncedFromOnPremises.IsSet() {
		toSerialize["isSyncedFromOnPremises"] = o.IsSyncedFromOnPremises.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject1193 struct {
	value *InlineObject1193
	isSet bool
}

func (v NullableInlineObject1193) Get() *InlineObject1193 {
	return v.value
}

func (v *NullableInlineObject1193) Set(val *InlineObject1193) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject1193) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject1193) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject1193(val *InlineObject1193) *NullableInlineObject1193 {
	return &NullableInlineObject1193{value: val, isSet: true}
}

func (v NullableInlineObject1193) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject1193) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


