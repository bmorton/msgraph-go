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

// InlineObject1012 struct for InlineObject1012
type InlineObject1012 struct {
	Message NullableString `json:"message,omitempty"`
}

// NewInlineObject1012 instantiates a new InlineObject1012 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject1012() *InlineObject1012 {
	this := InlineObject1012{}
	return &this
}

// NewInlineObject1012WithDefaults instantiates a new InlineObject1012 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject1012WithDefaults() *InlineObject1012 {
	this := InlineObject1012{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1012) GetMessage() string {
	if o == nil || o.Message.Get() == nil {
		var ret string
		return ret
	}
	return *o.Message.Get()
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1012) GetMessageOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Message.Get(), o.Message.IsSet()
}

// HasMessage returns a boolean if a field has been set.
func (o *InlineObject1012) HasMessage() bool {
	if o != nil && o.Message.IsSet() {
		return true
	}

	return false
}

// SetMessage gets a reference to the given NullableString and assigns it to the Message field.
func (o *InlineObject1012) SetMessage(v string) {
	o.Message.Set(&v)
}
// SetMessageNil sets the value for Message to be an explicit nil
func (o *InlineObject1012) SetMessageNil() {
	o.Message.Set(nil)
}

// UnsetMessage ensures that no value is present for Message, not even an explicit nil
func (o *InlineObject1012) UnsetMessage() {
	o.Message.Unset()
}

func (o InlineObject1012) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Message.IsSet() {
		toSerialize["message"] = o.Message.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject1012 struct {
	value *InlineObject1012
	isSet bool
}

func (v NullableInlineObject1012) Get() *InlineObject1012 {
	return v.value
}

func (v *NullableInlineObject1012) Set(val *InlineObject1012) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject1012) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject1012) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject1012(val *InlineObject1012) *NullableInlineObject1012 {
	return &NullableInlineObject1012{value: val, isSet: true}
}

func (v NullableInlineObject1012) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject1012) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


