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

// InlineObject723 struct for InlineObject723
type InlineObject723 struct {
	IsSyncedFromOnPremises NullableBool `json:"isSyncedFromOnPremises,omitempty"`
}

// NewInlineObject723 instantiates a new InlineObject723 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject723() *InlineObject723 {
	this := InlineObject723{}
	var isSyncedFromOnPremises bool = false
	this.IsSyncedFromOnPremises = *NewNullableBool(&isSyncedFromOnPremises)
	return &this
}

// NewInlineObject723WithDefaults instantiates a new InlineObject723 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject723WithDefaults() *InlineObject723 {
	this := InlineObject723{}
	var isSyncedFromOnPremises bool = false
	this.IsSyncedFromOnPremises = *NewNullableBool(&isSyncedFromOnPremises)
	return &this
}

// GetIsSyncedFromOnPremises returns the IsSyncedFromOnPremises field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject723) GetIsSyncedFromOnPremises() bool {
	if o == nil || o.IsSyncedFromOnPremises.Get() == nil {
		var ret bool
		return ret
	}
	return *o.IsSyncedFromOnPremises.Get()
}

// GetIsSyncedFromOnPremisesOk returns a tuple with the IsSyncedFromOnPremises field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject723) GetIsSyncedFromOnPremisesOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IsSyncedFromOnPremises.Get(), o.IsSyncedFromOnPremises.IsSet()
}

// HasIsSyncedFromOnPremises returns a boolean if a field has been set.
func (o *InlineObject723) HasIsSyncedFromOnPremises() bool {
	if o != nil && o.IsSyncedFromOnPremises.IsSet() {
		return true
	}

	return false
}

// SetIsSyncedFromOnPremises gets a reference to the given NullableBool and assigns it to the IsSyncedFromOnPremises field.
func (o *InlineObject723) SetIsSyncedFromOnPremises(v bool) {
	o.IsSyncedFromOnPremises.Set(&v)
}
// SetIsSyncedFromOnPremisesNil sets the value for IsSyncedFromOnPremises to be an explicit nil
func (o *InlineObject723) SetIsSyncedFromOnPremisesNil() {
	o.IsSyncedFromOnPremises.Set(nil)
}

// UnsetIsSyncedFromOnPremises ensures that no value is present for IsSyncedFromOnPremises, not even an explicit nil
func (o *InlineObject723) UnsetIsSyncedFromOnPremises() {
	o.IsSyncedFromOnPremises.Unset()
}

func (o InlineObject723) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.IsSyncedFromOnPremises.IsSet() {
		toSerialize["isSyncedFromOnPremises"] = o.IsSyncedFromOnPremises.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject723 struct {
	value *InlineObject723
	isSet bool
}

func (v NullableInlineObject723) Get() *InlineObject723 {
	return v.value
}

func (v *NullableInlineObject723) Set(val *InlineObject723) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject723) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject723) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject723(val *InlineObject723) *NullableInlineObject723 {
	return &NullableInlineObject723{value: val, isSet: true}
}

func (v NullableInlineObject723) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject723) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


