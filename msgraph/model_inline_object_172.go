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

// InlineObject172 struct for InlineObject172
type InlineObject172 struct {
	Comment NullableString `json:"Comment,omitempty"`
}

// NewInlineObject172 instantiates a new InlineObject172 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject172() *InlineObject172 {
	this := InlineObject172{}
	return &this
}

// NewInlineObject172WithDefaults instantiates a new InlineObject172 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject172WithDefaults() *InlineObject172 {
	this := InlineObject172{}
	return &this
}

// GetComment returns the Comment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject172) GetComment() string {
	if o == nil || o.Comment.Get() == nil {
		var ret string
		return ret
	}
	return *o.Comment.Get()
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject172) GetCommentOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Comment.Get(), o.Comment.IsSet()
}

// HasComment returns a boolean if a field has been set.
func (o *InlineObject172) HasComment() bool {
	if o != nil && o.Comment.IsSet() {
		return true
	}

	return false
}

// SetComment gets a reference to the given NullableString and assigns it to the Comment field.
func (o *InlineObject172) SetComment(v string) {
	o.Comment.Set(&v)
}
// SetCommentNil sets the value for Comment to be an explicit nil
func (o *InlineObject172) SetCommentNil() {
	o.Comment.Set(nil)
}

// UnsetComment ensures that no value is present for Comment, not even an explicit nil
func (o *InlineObject172) UnsetComment() {
	o.Comment.Unset()
}

func (o InlineObject172) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Comment.IsSet() {
		toSerialize["Comment"] = o.Comment.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject172 struct {
	value *InlineObject172
	isSet bool
}

func (v NullableInlineObject172) Get() *InlineObject172 {
	return v.value
}

func (v *NullableInlineObject172) Set(val *InlineObject172) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject172) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject172) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject172(val *InlineObject172) *NullableInlineObject172 {
	return &NullableInlineObject172{value: val, isSet: true}
}

func (v NullableInlineObject172) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject172) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


