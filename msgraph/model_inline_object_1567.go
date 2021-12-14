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

// InlineObject1567 struct for InlineObject1567
type InlineObject1567 struct {
	Color NullableString `json:"color,omitempty"`
}

// NewInlineObject1567 instantiates a new InlineObject1567 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject1567() *InlineObject1567 {
	this := InlineObject1567{}
	return &this
}

// NewInlineObject1567WithDefaults instantiates a new InlineObject1567 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject1567WithDefaults() *InlineObject1567 {
	this := InlineObject1567{}
	return &this
}

// GetColor returns the Color field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1567) GetColor() string {
	if o == nil || o.Color.Get() == nil {
		var ret string
		return ret
	}
	return *o.Color.Get()
}

// GetColorOk returns a tuple with the Color field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1567) GetColorOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Color.Get(), o.Color.IsSet()
}

// HasColor returns a boolean if a field has been set.
func (o *InlineObject1567) HasColor() bool {
	if o != nil && o.Color.IsSet() {
		return true
	}

	return false
}

// SetColor gets a reference to the given NullableString and assigns it to the Color field.
func (o *InlineObject1567) SetColor(v string) {
	o.Color.Set(&v)
}
// SetColorNil sets the value for Color to be an explicit nil
func (o *InlineObject1567) SetColorNil() {
	o.Color.Set(nil)
}

// UnsetColor ensures that no value is present for Color, not even an explicit nil
func (o *InlineObject1567) UnsetColor() {
	o.Color.Unset()
}

func (o InlineObject1567) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Color.IsSet() {
		toSerialize["color"] = o.Color.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject1567 struct {
	value *InlineObject1567
	isSet bool
}

func (v NullableInlineObject1567) Get() *InlineObject1567 {
	return v.value
}

func (v *NullableInlineObject1567) Set(val *InlineObject1567) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject1567) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject1567) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject1567(val *InlineObject1567) *NullableInlineObject1567 {
	return &NullableInlineObject1567{value: val, isSet: true}
}

func (v NullableInlineObject1567) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject1567) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


