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

// InlineObject549 struct for InlineObject549
type InlineObject549 struct {
	Message NullableString `json:"message,omitempty"`
}

// NewInlineObject549 instantiates a new InlineObject549 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject549() *InlineObject549 {
	this := InlineObject549{}
	return &this
}

// NewInlineObject549WithDefaults instantiates a new InlineObject549 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject549WithDefaults() *InlineObject549 {
	this := InlineObject549{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject549) GetMessage() string {
	if o == nil || o.Message.Get() == nil {
		var ret string
		return ret
	}
	return *o.Message.Get()
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject549) GetMessageOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Message.Get(), o.Message.IsSet()
}

// HasMessage returns a boolean if a field has been set.
func (o *InlineObject549) HasMessage() bool {
	if o != nil && o.Message.IsSet() {
		return true
	}

	return false
}

// SetMessage gets a reference to the given NullableString and assigns it to the Message field.
func (o *InlineObject549) SetMessage(v string) {
	o.Message.Set(&v)
}
// SetMessageNil sets the value for Message to be an explicit nil
func (o *InlineObject549) SetMessageNil() {
	o.Message.Set(nil)
}

// UnsetMessage ensures that no value is present for Message, not even an explicit nil
func (o *InlineObject549) UnsetMessage() {
	o.Message.Unset()
}

func (o InlineObject549) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Message.IsSet() {
		toSerialize["message"] = o.Message.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject549 struct {
	value *InlineObject549
	isSet bool
}

func (v NullableInlineObject549) Get() *InlineObject549 {
	return v.value
}

func (v *NullableInlineObject549) Set(val *InlineObject549) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject549) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject549) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject549(val *InlineObject549) *NullableInlineObject549 {
	return &NullableInlineObject549{value: val, isSet: true}
}

func (v NullableInlineObject549) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject549) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


