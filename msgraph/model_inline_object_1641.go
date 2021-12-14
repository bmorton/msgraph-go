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

// InlineObject1641 struct for InlineObject1641
type InlineObject1641 struct {
	Color NullableString `json:"color,omitempty"`
}

// NewInlineObject1641 instantiates a new InlineObject1641 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject1641() *InlineObject1641 {
	this := InlineObject1641{}
	return &this
}

// NewInlineObject1641WithDefaults instantiates a new InlineObject1641 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject1641WithDefaults() *InlineObject1641 {
	this := InlineObject1641{}
	return &this
}

// GetColor returns the Color field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1641) GetColor() string {
	if o == nil || o.Color.Get() == nil {
		var ret string
		return ret
	}
	return *o.Color.Get()
}

// GetColorOk returns a tuple with the Color field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1641) GetColorOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Color.Get(), o.Color.IsSet()
}

// HasColor returns a boolean if a field has been set.
func (o *InlineObject1641) HasColor() bool {
	if o != nil && o.Color.IsSet() {
		return true
	}

	return false
}

// SetColor gets a reference to the given NullableString and assigns it to the Color field.
func (o *InlineObject1641) SetColor(v string) {
	o.Color.Set(&v)
}
// SetColorNil sets the value for Color to be an explicit nil
func (o *InlineObject1641) SetColorNil() {
	o.Color.Set(nil)
}

// UnsetColor ensures that no value is present for Color, not even an explicit nil
func (o *InlineObject1641) UnsetColor() {
	o.Color.Unset()
}

func (o InlineObject1641) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Color.IsSet() {
		toSerialize["color"] = o.Color.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject1641 struct {
	value *InlineObject1641
	isSet bool
}

func (v NullableInlineObject1641) Get() *InlineObject1641 {
	return v.value
}

func (v *NullableInlineObject1641) Set(val *InlineObject1641) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject1641) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject1641) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject1641(val *InlineObject1641) *NullableInlineObject1641 {
	return &NullableInlineObject1641{value: val, isSet: true}
}

func (v NullableInlineObject1641) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject1641) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

