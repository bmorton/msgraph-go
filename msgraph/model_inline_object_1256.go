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

// InlineObject1256 struct for InlineObject1256
type InlineObject1256 struct {
	Text AnyOfobject `json:"text,omitempty"`
}

// NewInlineObject1256 instantiates a new InlineObject1256 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject1256() *InlineObject1256 {
	this := InlineObject1256{}
	return &this
}

// NewInlineObject1256WithDefaults instantiates a new InlineObject1256 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject1256WithDefaults() *InlineObject1256 {
	this := InlineObject1256{}
	return &this
}

// GetText returns the Text field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1256) GetText() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1256) GetTextOk() (*AnyOfobject, bool) {
	if o == nil || o.Text == nil {
		return nil, false
	}
	return &o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *InlineObject1256) HasText() bool {
	if o != nil && o.Text != nil {
		return true
	}

	return false
}

// SetText gets a reference to the given AnyOfobject and assigns it to the Text field.
func (o *InlineObject1256) SetText(v AnyOfobject) {
	o.Text = v
}

func (o InlineObject1256) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Text != nil {
		toSerialize["text"] = o.Text
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject1256 struct {
	value *InlineObject1256
	isSet bool
}

func (v NullableInlineObject1256) Get() *InlineObject1256 {
	return v.value
}

func (v *NullableInlineObject1256) Set(val *InlineObject1256) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject1256) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject1256) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject1256(val *InlineObject1256) *NullableInlineObject1256 {
	return &NullableInlineObject1256{value: val, isSet: true}
}

func (v NullableInlineObject1256) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject1256) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


