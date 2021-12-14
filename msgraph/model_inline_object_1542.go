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

// InlineObject1542 struct for InlineObject1542
type InlineObject1542 struct {
	Text AnyOfobject `json:"text,omitempty"`
}

// NewInlineObject1542 instantiates a new InlineObject1542 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject1542() *InlineObject1542 {
	this := InlineObject1542{}
	return &this
}

// NewInlineObject1542WithDefaults instantiates a new InlineObject1542 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject1542WithDefaults() *InlineObject1542 {
	this := InlineObject1542{}
	return &this
}

// GetText returns the Text field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1542) GetText() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1542) GetTextOk() (*AnyOfobject, bool) {
	if o == nil || o.Text == nil {
		return nil, false
	}
	return &o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *InlineObject1542) HasText() bool {
	if o != nil && o.Text != nil {
		return true
	}

	return false
}

// SetText gets a reference to the given AnyOfobject and assigns it to the Text field.
func (o *InlineObject1542) SetText(v AnyOfobject) {
	o.Text = v
}

func (o InlineObject1542) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Text != nil {
		toSerialize["text"] = o.Text
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject1542 struct {
	value *InlineObject1542
	isSet bool
}

func (v NullableInlineObject1542) Get() *InlineObject1542 {
	return v.value
}

func (v *NullableInlineObject1542) Set(val *InlineObject1542) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject1542) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject1542) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject1542(val *InlineObject1542) *NullableInlineObject1542 {
	return &NullableInlineObject1542{value: val, isSet: true}
}

func (v NullableInlineObject1542) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject1542) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


