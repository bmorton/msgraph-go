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

// InlineObject1065 struct for InlineObject1065
type InlineObject1065 struct {
	KeepUserData *bool `json:"keepUserData,omitempty"`
}

// NewInlineObject1065 instantiates a new InlineObject1065 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject1065() *InlineObject1065 {
	this := InlineObject1065{}
	var keepUserData bool = false
	this.KeepUserData = &keepUserData
	return &this
}

// NewInlineObject1065WithDefaults instantiates a new InlineObject1065 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject1065WithDefaults() *InlineObject1065 {
	this := InlineObject1065{}
	var keepUserData bool = false
	this.KeepUserData = &keepUserData
	return &this
}

// GetKeepUserData returns the KeepUserData field value if set, zero value otherwise.
func (o *InlineObject1065) GetKeepUserData() bool {
	if o == nil || o.KeepUserData == nil {
		var ret bool
		return ret
	}
	return *o.KeepUserData
}

// GetKeepUserDataOk returns a tuple with the KeepUserData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject1065) GetKeepUserDataOk() (*bool, bool) {
	if o == nil || o.KeepUserData == nil {
		return nil, false
	}
	return o.KeepUserData, true
}

// HasKeepUserData returns a boolean if a field has been set.
func (o *InlineObject1065) HasKeepUserData() bool {
	if o != nil && o.KeepUserData != nil {
		return true
	}

	return false
}

// SetKeepUserData gets a reference to the given bool and assigns it to the KeepUserData field.
func (o *InlineObject1065) SetKeepUserData(v bool) {
	o.KeepUserData = &v
}

func (o InlineObject1065) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.KeepUserData != nil {
		toSerialize["keepUserData"] = o.KeepUserData
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject1065 struct {
	value *InlineObject1065
	isSet bool
}

func (v NullableInlineObject1065) Get() *InlineObject1065 {
	return v.value
}

func (v *NullableInlineObject1065) Set(val *InlineObject1065) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject1065) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject1065) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject1065(val *InlineObject1065) *NullableInlineObject1065 {
	return &NullableInlineObject1065{value: val, isSet: true}
}

func (v NullableInlineObject1065) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject1065) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


