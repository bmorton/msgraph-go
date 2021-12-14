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

// InlineObject461 struct for InlineObject461
type InlineObject461 struct {
	Comment NullableString `json:"Comment,omitempty"`
}

// NewInlineObject461 instantiates a new InlineObject461 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject461() *InlineObject461 {
	this := InlineObject461{}
	return &this
}

// NewInlineObject461WithDefaults instantiates a new InlineObject461 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject461WithDefaults() *InlineObject461 {
	this := InlineObject461{}
	return &this
}

// GetComment returns the Comment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject461) GetComment() string {
	if o == nil || o.Comment.Get() == nil {
		var ret string
		return ret
	}
	return *o.Comment.Get()
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject461) GetCommentOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Comment.Get(), o.Comment.IsSet()
}

// HasComment returns a boolean if a field has been set.
func (o *InlineObject461) HasComment() bool {
	if o != nil && o.Comment.IsSet() {
		return true
	}

	return false
}

// SetComment gets a reference to the given NullableString and assigns it to the Comment field.
func (o *InlineObject461) SetComment(v string) {
	o.Comment.Set(&v)
}
// SetCommentNil sets the value for Comment to be an explicit nil
func (o *InlineObject461) SetCommentNil() {
	o.Comment.Set(nil)
}

// UnsetComment ensures that no value is present for Comment, not even an explicit nil
func (o *InlineObject461) UnsetComment() {
	o.Comment.Unset()
}

func (o InlineObject461) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Comment.IsSet() {
		toSerialize["Comment"] = o.Comment.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject461 struct {
	value *InlineObject461
	isSet bool
}

func (v NullableInlineObject461) Get() *InlineObject461 {
	return v.value
}

func (v *NullableInlineObject461) Set(val *InlineObject461) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject461) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject461) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject461(val *InlineObject461) *NullableInlineObject461 {
	return &NullableInlineObject461{value: val, isSet: true}
}

func (v NullableInlineObject461) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject461) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


