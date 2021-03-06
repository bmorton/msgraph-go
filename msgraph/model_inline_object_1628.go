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

// InlineObject1628 struct for InlineObject1628
type InlineObject1628 struct {
	Color NullableString `json:"color,omitempty"`
}

// NewInlineObject1628 instantiates a new InlineObject1628 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject1628() *InlineObject1628 {
	this := InlineObject1628{}
	return &this
}

// NewInlineObject1628WithDefaults instantiates a new InlineObject1628 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject1628WithDefaults() *InlineObject1628 {
	this := InlineObject1628{}
	return &this
}

// GetColor returns the Color field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1628) GetColor() string {
	if o == nil || o.Color.Get() == nil {
		var ret string
		return ret
	}
	return *o.Color.Get()
}

// GetColorOk returns a tuple with the Color field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1628) GetColorOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Color.Get(), o.Color.IsSet()
}

// HasColor returns a boolean if a field has been set.
func (o *InlineObject1628) HasColor() bool {
	if o != nil && o.Color.IsSet() {
		return true
	}

	return false
}

// SetColor gets a reference to the given NullableString and assigns it to the Color field.
func (o *InlineObject1628) SetColor(v string) {
	o.Color.Set(&v)
}
// SetColorNil sets the value for Color to be an explicit nil
func (o *InlineObject1628) SetColorNil() {
	o.Color.Set(nil)
}

// UnsetColor ensures that no value is present for Color, not even an explicit nil
func (o *InlineObject1628) UnsetColor() {
	o.Color.Unset()
}

func (o InlineObject1628) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Color.IsSet() {
		toSerialize["color"] = o.Color.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject1628 struct {
	value *InlineObject1628
	isSet bool
}

func (v NullableInlineObject1628) Get() *InlineObject1628 {
	return v.value
}

func (v *NullableInlineObject1628) Set(val *InlineObject1628) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject1628) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject1628) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject1628(val *InlineObject1628) *NullableInlineObject1628 {
	return &NullableInlineObject1628{value: val, isSet: true}
}

func (v NullableInlineObject1628) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject1628) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


