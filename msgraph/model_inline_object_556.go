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

// InlineObject556 struct for InlineObject556
type InlineObject556 struct {
	Across *bool `json:"across,omitempty"`
}

// NewInlineObject556 instantiates a new InlineObject556 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject556() *InlineObject556 {
	this := InlineObject556{}
	var across bool = false
	this.Across = &across
	return &this
}

// NewInlineObject556WithDefaults instantiates a new InlineObject556 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject556WithDefaults() *InlineObject556 {
	this := InlineObject556{}
	var across bool = false
	this.Across = &across
	return &this
}

// GetAcross returns the Across field value if set, zero value otherwise.
func (o *InlineObject556) GetAcross() bool {
	if o == nil || o.Across == nil {
		var ret bool
		return ret
	}
	return *o.Across
}

// GetAcrossOk returns a tuple with the Across field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject556) GetAcrossOk() (*bool, bool) {
	if o == nil || o.Across == nil {
		return nil, false
	}
	return o.Across, true
}

// HasAcross returns a boolean if a field has been set.
func (o *InlineObject556) HasAcross() bool {
	if o != nil && o.Across != nil {
		return true
	}

	return false
}

// SetAcross gets a reference to the given bool and assigns it to the Across field.
func (o *InlineObject556) SetAcross(v bool) {
	o.Across = &v
}

func (o InlineObject556) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Across != nil {
		toSerialize["across"] = o.Across
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject556 struct {
	value *InlineObject556
	isSet bool
}

func (v NullableInlineObject556) Get() *InlineObject556 {
	return v.value
}

func (v *NullableInlineObject556) Set(val *InlineObject556) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject556) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject556) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject556(val *InlineObject556) *NullableInlineObject556 {
	return &NullableInlineObject556{value: val, isSet: true}
}

func (v NullableInlineObject556) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject556) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


