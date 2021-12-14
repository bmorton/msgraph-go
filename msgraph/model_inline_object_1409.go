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

// InlineObject1409 struct for InlineObject1409
type InlineObject1409 struct {
	Text AnyOfobject `json:"text,omitempty"`
	NumChars AnyOfobject `json:"numChars,omitempty"`
}

// NewInlineObject1409 instantiates a new InlineObject1409 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject1409() *InlineObject1409 {
	this := InlineObject1409{}
	return &this
}

// NewInlineObject1409WithDefaults instantiates a new InlineObject1409 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject1409WithDefaults() *InlineObject1409 {
	this := InlineObject1409{}
	return &this
}

// GetText returns the Text field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1409) GetText() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1409) GetTextOk() (*AnyOfobject, bool) {
	if o == nil || o.Text == nil {
		return nil, false
	}
	return &o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *InlineObject1409) HasText() bool {
	if o != nil && o.Text != nil {
		return true
	}

	return false
}

// SetText gets a reference to the given AnyOfobject and assigns it to the Text field.
func (o *InlineObject1409) SetText(v AnyOfobject) {
	o.Text = v
}

// GetNumChars returns the NumChars field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1409) GetNumChars() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.NumChars
}

// GetNumCharsOk returns a tuple with the NumChars field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1409) GetNumCharsOk() (*AnyOfobject, bool) {
	if o == nil || o.NumChars == nil {
		return nil, false
	}
	return &o.NumChars, true
}

// HasNumChars returns a boolean if a field has been set.
func (o *InlineObject1409) HasNumChars() bool {
	if o != nil && o.NumChars != nil {
		return true
	}

	return false
}

// SetNumChars gets a reference to the given AnyOfobject and assigns it to the NumChars field.
func (o *InlineObject1409) SetNumChars(v AnyOfobject) {
	o.NumChars = v
}

func (o InlineObject1409) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Text != nil {
		toSerialize["text"] = o.Text
	}
	if o.NumChars != nil {
		toSerialize["numChars"] = o.NumChars
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject1409 struct {
	value *InlineObject1409
	isSet bool
}

func (v NullableInlineObject1409) Get() *InlineObject1409 {
	return v.value
}

func (v *NullableInlineObject1409) Set(val *InlineObject1409) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject1409) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject1409) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject1409(val *InlineObject1409) *NullableInlineObject1409 {
	return &NullableInlineObject1409{value: val, isSet: true}
}

func (v NullableInlineObject1409) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject1409) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


