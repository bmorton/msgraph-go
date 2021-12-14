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

// InlineObject1410 struct for InlineObject1410
type InlineObject1410 struct {
	Text AnyOfobject `json:"text,omitempty"`
	NumBytes AnyOfobject `json:"numBytes,omitempty"`
}

// NewInlineObject1410 instantiates a new InlineObject1410 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject1410() *InlineObject1410 {
	this := InlineObject1410{}
	return &this
}

// NewInlineObject1410WithDefaults instantiates a new InlineObject1410 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject1410WithDefaults() *InlineObject1410 {
	this := InlineObject1410{}
	return &this
}

// GetText returns the Text field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1410) GetText() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1410) GetTextOk() (*AnyOfobject, bool) {
	if o == nil || o.Text == nil {
		return nil, false
	}
	return &o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *InlineObject1410) HasText() bool {
	if o != nil && o.Text != nil {
		return true
	}

	return false
}

// SetText gets a reference to the given AnyOfobject and assigns it to the Text field.
func (o *InlineObject1410) SetText(v AnyOfobject) {
	o.Text = v
}

// GetNumBytes returns the NumBytes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1410) GetNumBytes() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.NumBytes
}

// GetNumBytesOk returns a tuple with the NumBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1410) GetNumBytesOk() (*AnyOfobject, bool) {
	if o == nil || o.NumBytes == nil {
		return nil, false
	}
	return &o.NumBytes, true
}

// HasNumBytes returns a boolean if a field has been set.
func (o *InlineObject1410) HasNumBytes() bool {
	if o != nil && o.NumBytes != nil {
		return true
	}

	return false
}

// SetNumBytes gets a reference to the given AnyOfobject and assigns it to the NumBytes field.
func (o *InlineObject1410) SetNumBytes(v AnyOfobject) {
	o.NumBytes = v
}

func (o InlineObject1410) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Text != nil {
		toSerialize["text"] = o.Text
	}
	if o.NumBytes != nil {
		toSerialize["numBytes"] = o.NumBytes
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject1410 struct {
	value *InlineObject1410
	isSet bool
}

func (v NullableInlineObject1410) Get() *InlineObject1410 {
	return v.value
}

func (v *NullableInlineObject1410) Set(val *InlineObject1410) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject1410) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject1410) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject1410(val *InlineObject1410) *NullableInlineObject1410 {
	return &NullableInlineObject1410{value: val, isSet: true}
}

func (v NullableInlineObject1410) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject1410) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


