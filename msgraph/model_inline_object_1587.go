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

// InlineObject1587 struct for InlineObject1587
type InlineObject1587 struct {
	Percent *int32 `json:"percent,omitempty"`
}

// NewInlineObject1587 instantiates a new InlineObject1587 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject1587() *InlineObject1587 {
	this := InlineObject1587{}
	return &this
}

// NewInlineObject1587WithDefaults instantiates a new InlineObject1587 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject1587WithDefaults() *InlineObject1587 {
	this := InlineObject1587{}
	return &this
}

// GetPercent returns the Percent field value if set, zero value otherwise.
func (o *InlineObject1587) GetPercent() int32 {
	if o == nil || o.Percent == nil {
		var ret int32
		return ret
	}
	return *o.Percent
}

// GetPercentOk returns a tuple with the Percent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject1587) GetPercentOk() (*int32, bool) {
	if o == nil || o.Percent == nil {
		return nil, false
	}
	return o.Percent, true
}

// HasPercent returns a boolean if a field has been set.
func (o *InlineObject1587) HasPercent() bool {
	if o != nil && o.Percent != nil {
		return true
	}

	return false
}

// SetPercent gets a reference to the given int32 and assigns it to the Percent field.
func (o *InlineObject1587) SetPercent(v int32) {
	o.Percent = &v
}

func (o InlineObject1587) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Percent != nil {
		toSerialize["percent"] = o.Percent
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject1587 struct {
	value *InlineObject1587
	isSet bool
}

func (v NullableInlineObject1587) Get() *InlineObject1587 {
	return v.value
}

func (v *NullableInlineObject1587) Set(val *InlineObject1587) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject1587) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject1587) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject1587(val *InlineObject1587) *NullableInlineObject1587 {
	return &NullableInlineObject1587{value: val, isSet: true}
}

func (v NullableInlineObject1587) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject1587) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


