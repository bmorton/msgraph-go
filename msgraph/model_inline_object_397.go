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

// InlineObject397 struct for InlineObject397
type InlineObject397 struct {
	Comment NullableString `json:"Comment,omitempty"`
}

// NewInlineObject397 instantiates a new InlineObject397 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject397() *InlineObject397 {
	this := InlineObject397{}
	return &this
}

// NewInlineObject397WithDefaults instantiates a new InlineObject397 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject397WithDefaults() *InlineObject397 {
	this := InlineObject397{}
	return &this
}

// GetComment returns the Comment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject397) GetComment() string {
	if o == nil || o.Comment.Get() == nil {
		var ret string
		return ret
	}
	return *o.Comment.Get()
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject397) GetCommentOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Comment.Get(), o.Comment.IsSet()
}

// HasComment returns a boolean if a field has been set.
func (o *InlineObject397) HasComment() bool {
	if o != nil && o.Comment.IsSet() {
		return true
	}

	return false
}

// SetComment gets a reference to the given NullableString and assigns it to the Comment field.
func (o *InlineObject397) SetComment(v string) {
	o.Comment.Set(&v)
}
// SetCommentNil sets the value for Comment to be an explicit nil
func (o *InlineObject397) SetCommentNil() {
	o.Comment.Set(nil)
}

// UnsetComment ensures that no value is present for Comment, not even an explicit nil
func (o *InlineObject397) UnsetComment() {
	o.Comment.Unset()
}

func (o InlineObject397) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Comment.IsSet() {
		toSerialize["Comment"] = o.Comment.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject397 struct {
	value *InlineObject397
	isSet bool
}

func (v NullableInlineObject397) Get() *InlineObject397 {
	return v.value
}

func (v *NullableInlineObject397) Set(val *InlineObject397) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject397) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject397) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject397(val *InlineObject397) *NullableInlineObject397 {
	return &NullableInlineObject397{value: val, isSet: true}
}

func (v NullableInlineObject397) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject397) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

