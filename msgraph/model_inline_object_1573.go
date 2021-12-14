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

// InlineObject1573 struct for InlineObject1573
type InlineObject1573 struct {
	Color NullableString `json:"color,omitempty"`
}

// NewInlineObject1573 instantiates a new InlineObject1573 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject1573() *InlineObject1573 {
	this := InlineObject1573{}
	return &this
}

// NewInlineObject1573WithDefaults instantiates a new InlineObject1573 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject1573WithDefaults() *InlineObject1573 {
	this := InlineObject1573{}
	return &this
}

// GetColor returns the Color field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1573) GetColor() string {
	if o == nil || o.Color.Get() == nil {
		var ret string
		return ret
	}
	return *o.Color.Get()
}

// GetColorOk returns a tuple with the Color field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1573) GetColorOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Color.Get(), o.Color.IsSet()
}

// HasColor returns a boolean if a field has been set.
func (o *InlineObject1573) HasColor() bool {
	if o != nil && o.Color.IsSet() {
		return true
	}

	return false
}

// SetColor gets a reference to the given NullableString and assigns it to the Color field.
func (o *InlineObject1573) SetColor(v string) {
	o.Color.Set(&v)
}
// SetColorNil sets the value for Color to be an explicit nil
func (o *InlineObject1573) SetColorNil() {
	o.Color.Set(nil)
}

// UnsetColor ensures that no value is present for Color, not even an explicit nil
func (o *InlineObject1573) UnsetColor() {
	o.Color.Unset()
}

func (o InlineObject1573) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Color.IsSet() {
		toSerialize["color"] = o.Color.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject1573 struct {
	value *InlineObject1573
	isSet bool
}

func (v NullableInlineObject1573) Get() *InlineObject1573 {
	return v.value
}

func (v *NullableInlineObject1573) Set(val *InlineObject1573) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject1573) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject1573) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject1573(val *InlineObject1573) *NullableInlineObject1573 {
	return &NullableInlineObject1573{value: val, isSet: true}
}

func (v NullableInlineObject1573) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject1573) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


