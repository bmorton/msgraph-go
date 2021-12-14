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

// InlineObject186 struct for InlineObject186
type InlineObject186 struct {
	Comment NullableString `json:"Comment,omitempty"`
}

// NewInlineObject186 instantiates a new InlineObject186 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject186() *InlineObject186 {
	this := InlineObject186{}
	return &this
}

// NewInlineObject186WithDefaults instantiates a new InlineObject186 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject186WithDefaults() *InlineObject186 {
	this := InlineObject186{}
	return &this
}

// GetComment returns the Comment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject186) GetComment() string {
	if o == nil || o.Comment.Get() == nil {
		var ret string
		return ret
	}
	return *o.Comment.Get()
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject186) GetCommentOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Comment.Get(), o.Comment.IsSet()
}

// HasComment returns a boolean if a field has been set.
func (o *InlineObject186) HasComment() bool {
	if o != nil && o.Comment.IsSet() {
		return true
	}

	return false
}

// SetComment gets a reference to the given NullableString and assigns it to the Comment field.
func (o *InlineObject186) SetComment(v string) {
	o.Comment.Set(&v)
}
// SetCommentNil sets the value for Comment to be an explicit nil
func (o *InlineObject186) SetCommentNil() {
	o.Comment.Set(nil)
}

// UnsetComment ensures that no value is present for Comment, not even an explicit nil
func (o *InlineObject186) UnsetComment() {
	o.Comment.Unset()
}

func (o InlineObject186) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Comment.IsSet() {
		toSerialize["Comment"] = o.Comment.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject186 struct {
	value *InlineObject186
	isSet bool
}

func (v NullableInlineObject186) Get() *InlineObject186 {
	return v.value
}

func (v *NullableInlineObject186) Set(val *InlineObject186) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject186) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject186) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject186(val *InlineObject186) *NullableInlineObject186 {
	return &NullableInlineObject186{value: val, isSet: true}
}

func (v NullableInlineObject186) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject186) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


