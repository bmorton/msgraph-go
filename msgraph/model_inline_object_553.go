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

// InlineObject553 struct for InlineObject553
type InlineObject553 struct {
	ApplyTo *string `json:"applyTo,omitempty"`
}

// NewInlineObject553 instantiates a new InlineObject553 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject553() *InlineObject553 {
	this := InlineObject553{}
	return &this
}

// NewInlineObject553WithDefaults instantiates a new InlineObject553 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject553WithDefaults() *InlineObject553 {
	this := InlineObject553{}
	return &this
}

// GetApplyTo returns the ApplyTo field value if set, zero value otherwise.
func (o *InlineObject553) GetApplyTo() string {
	if o == nil || o.ApplyTo == nil {
		var ret string
		return ret
	}
	return *o.ApplyTo
}

// GetApplyToOk returns a tuple with the ApplyTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject553) GetApplyToOk() (*string, bool) {
	if o == nil || o.ApplyTo == nil {
		return nil, false
	}
	return o.ApplyTo, true
}

// HasApplyTo returns a boolean if a field has been set.
func (o *InlineObject553) HasApplyTo() bool {
	if o != nil && o.ApplyTo != nil {
		return true
	}

	return false
}

// SetApplyTo gets a reference to the given string and assigns it to the ApplyTo field.
func (o *InlineObject553) SetApplyTo(v string) {
	o.ApplyTo = &v
}

func (o InlineObject553) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ApplyTo != nil {
		toSerialize["applyTo"] = o.ApplyTo
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject553 struct {
	value *InlineObject553
	isSet bool
}

func (v NullableInlineObject553) Get() *InlineObject553 {
	return v.value
}

func (v *NullableInlineObject553) Set(val *InlineObject553) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject553) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject553) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject553(val *InlineObject553) *NullableInlineObject553 {
	return &NullableInlineObject553{value: val, isSet: true}
}

func (v NullableInlineObject553) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject553) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

