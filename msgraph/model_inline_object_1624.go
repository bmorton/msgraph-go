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

// InlineObject1624 struct for InlineObject1624
type InlineObject1624 struct {
	Color NullableString `json:"color,omitempty"`
}

// NewInlineObject1624 instantiates a new InlineObject1624 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject1624() *InlineObject1624 {
	this := InlineObject1624{}
	return &this
}

// NewInlineObject1624WithDefaults instantiates a new InlineObject1624 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject1624WithDefaults() *InlineObject1624 {
	this := InlineObject1624{}
	return &this
}

// GetColor returns the Color field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1624) GetColor() string {
	if o == nil || o.Color.Get() == nil {
		var ret string
		return ret
	}
	return *o.Color.Get()
}

// GetColorOk returns a tuple with the Color field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1624) GetColorOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Color.Get(), o.Color.IsSet()
}

// HasColor returns a boolean if a field has been set.
func (o *InlineObject1624) HasColor() bool {
	if o != nil && o.Color.IsSet() {
		return true
	}

	return false
}

// SetColor gets a reference to the given NullableString and assigns it to the Color field.
func (o *InlineObject1624) SetColor(v string) {
	o.Color.Set(&v)
}
// SetColorNil sets the value for Color to be an explicit nil
func (o *InlineObject1624) SetColorNil() {
	o.Color.Set(nil)
}

// UnsetColor ensures that no value is present for Color, not even an explicit nil
func (o *InlineObject1624) UnsetColor() {
	o.Color.Unset()
}

func (o InlineObject1624) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Color.IsSet() {
		toSerialize["color"] = o.Color.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject1624 struct {
	value *InlineObject1624
	isSet bool
}

func (v NullableInlineObject1624) Get() *InlineObject1624 {
	return v.value
}

func (v *NullableInlineObject1624) Set(val *InlineObject1624) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject1624) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject1624) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject1624(val *InlineObject1624) *NullableInlineObject1624 {
	return &NullableInlineObject1624{value: val, isSet: true}
}

func (v NullableInlineObject1624) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject1624) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


