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

// InlineObject110 struct for InlineObject110
type InlineObject110 struct {
	IsSyncedFromOnPremises NullableBool `json:"isSyncedFromOnPremises,omitempty"`
}

// NewInlineObject110 instantiates a new InlineObject110 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject110() *InlineObject110 {
	this := InlineObject110{}
	var isSyncedFromOnPremises bool = false
	this.IsSyncedFromOnPremises = *NewNullableBool(&isSyncedFromOnPremises)
	return &this
}

// NewInlineObject110WithDefaults instantiates a new InlineObject110 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject110WithDefaults() *InlineObject110 {
	this := InlineObject110{}
	var isSyncedFromOnPremises bool = false
	this.IsSyncedFromOnPremises = *NewNullableBool(&isSyncedFromOnPremises)
	return &this
}

// GetIsSyncedFromOnPremises returns the IsSyncedFromOnPremises field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject110) GetIsSyncedFromOnPremises() bool {
	if o == nil || o.IsSyncedFromOnPremises.Get() == nil {
		var ret bool
		return ret
	}
	return *o.IsSyncedFromOnPremises.Get()
}

// GetIsSyncedFromOnPremisesOk returns a tuple with the IsSyncedFromOnPremises field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject110) GetIsSyncedFromOnPremisesOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IsSyncedFromOnPremises.Get(), o.IsSyncedFromOnPremises.IsSet()
}

// HasIsSyncedFromOnPremises returns a boolean if a field has been set.
func (o *InlineObject110) HasIsSyncedFromOnPremises() bool {
	if o != nil && o.IsSyncedFromOnPremises.IsSet() {
		return true
	}

	return false
}

// SetIsSyncedFromOnPremises gets a reference to the given NullableBool and assigns it to the IsSyncedFromOnPremises field.
func (o *InlineObject110) SetIsSyncedFromOnPremises(v bool) {
	o.IsSyncedFromOnPremises.Set(&v)
}
// SetIsSyncedFromOnPremisesNil sets the value for IsSyncedFromOnPremises to be an explicit nil
func (o *InlineObject110) SetIsSyncedFromOnPremisesNil() {
	o.IsSyncedFromOnPremises.Set(nil)
}

// UnsetIsSyncedFromOnPremises ensures that no value is present for IsSyncedFromOnPremises, not even an explicit nil
func (o *InlineObject110) UnsetIsSyncedFromOnPremises() {
	o.IsSyncedFromOnPremises.Unset()
}

func (o InlineObject110) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.IsSyncedFromOnPremises.IsSet() {
		toSerialize["isSyncedFromOnPremises"] = o.IsSyncedFromOnPremises.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject110 struct {
	value *InlineObject110
	isSet bool
}

func (v NullableInlineObject110) Get() *InlineObject110 {
	return v.value
}

func (v *NullableInlineObject110) Set(val *InlineObject110) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject110) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject110) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject110(val *InlineObject110) *NullableInlineObject110 {
	return &NullableInlineObject110{value: val, isSet: true}
}

func (v NullableInlineObject110) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject110) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


