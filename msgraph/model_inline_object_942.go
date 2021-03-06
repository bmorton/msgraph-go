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

// InlineObject942 struct for InlineObject942
type InlineObject942 struct {
	Comment NullableString `json:"Comment,omitempty"`
}

// NewInlineObject942 instantiates a new InlineObject942 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject942() *InlineObject942 {
	this := InlineObject942{}
	return &this
}

// NewInlineObject942WithDefaults instantiates a new InlineObject942 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject942WithDefaults() *InlineObject942 {
	this := InlineObject942{}
	return &this
}

// GetComment returns the Comment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject942) GetComment() string {
	if o == nil || o.Comment.Get() == nil {
		var ret string
		return ret
	}
	return *o.Comment.Get()
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject942) GetCommentOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Comment.Get(), o.Comment.IsSet()
}

// HasComment returns a boolean if a field has been set.
func (o *InlineObject942) HasComment() bool {
	if o != nil && o.Comment.IsSet() {
		return true
	}

	return false
}

// SetComment gets a reference to the given NullableString and assigns it to the Comment field.
func (o *InlineObject942) SetComment(v string) {
	o.Comment.Set(&v)
}
// SetCommentNil sets the value for Comment to be an explicit nil
func (o *InlineObject942) SetCommentNil() {
	o.Comment.Set(nil)
}

// UnsetComment ensures that no value is present for Comment, not even an explicit nil
func (o *InlineObject942) UnsetComment() {
	o.Comment.Unset()
}

func (o InlineObject942) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Comment.IsSet() {
		toSerialize["Comment"] = o.Comment.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject942 struct {
	value *InlineObject942
	isSet bool
}

func (v NullableInlineObject942) Get() *InlineObject942 {
	return v.value
}

func (v *NullableInlineObject942) Set(val *InlineObject942) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject942) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject942) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject942(val *InlineObject942) *NullableInlineObject942 {
	return &NullableInlineObject942{value: val, isSet: true}
}

func (v NullableInlineObject942) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject942) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


