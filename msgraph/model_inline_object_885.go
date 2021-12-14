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

// InlineObject885 struct for InlineObject885
type InlineObject885 struct {
	Comment NullableString `json:"Comment,omitempty"`
}

// NewInlineObject885 instantiates a new InlineObject885 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject885() *InlineObject885 {
	this := InlineObject885{}
	return &this
}

// NewInlineObject885WithDefaults instantiates a new InlineObject885 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject885WithDefaults() *InlineObject885 {
	this := InlineObject885{}
	return &this
}

// GetComment returns the Comment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject885) GetComment() string {
	if o == nil || o.Comment.Get() == nil {
		var ret string
		return ret
	}
	return *o.Comment.Get()
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject885) GetCommentOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Comment.Get(), o.Comment.IsSet()
}

// HasComment returns a boolean if a field has been set.
func (o *InlineObject885) HasComment() bool {
	if o != nil && o.Comment.IsSet() {
		return true
	}

	return false
}

// SetComment gets a reference to the given NullableString and assigns it to the Comment field.
func (o *InlineObject885) SetComment(v string) {
	o.Comment.Set(&v)
}
// SetCommentNil sets the value for Comment to be an explicit nil
func (o *InlineObject885) SetCommentNil() {
	o.Comment.Set(nil)
}

// UnsetComment ensures that no value is present for Comment, not even an explicit nil
func (o *InlineObject885) UnsetComment() {
	o.Comment.Unset()
}

func (o InlineObject885) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Comment.IsSet() {
		toSerialize["Comment"] = o.Comment.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject885 struct {
	value *InlineObject885
	isSet bool
}

func (v NullableInlineObject885) Get() *InlineObject885 {
	return v.value
}

func (v *NullableInlineObject885) Set(val *InlineObject885) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject885) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject885) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject885(val *InlineObject885) *NullableInlineObject885 {
	return &NullableInlineObject885{value: val, isSet: true}
}

func (v NullableInlineObject885) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject885) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


