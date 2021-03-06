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

// InlineObject487 struct for InlineObject487
type InlineObject487 struct {
	Comment NullableString `json:"Comment,omitempty"`
}

// NewInlineObject487 instantiates a new InlineObject487 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject487() *InlineObject487 {
	this := InlineObject487{}
	return &this
}

// NewInlineObject487WithDefaults instantiates a new InlineObject487 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject487WithDefaults() *InlineObject487 {
	this := InlineObject487{}
	return &this
}

// GetComment returns the Comment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject487) GetComment() string {
	if o == nil || o.Comment.Get() == nil {
		var ret string
		return ret
	}
	return *o.Comment.Get()
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject487) GetCommentOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Comment.Get(), o.Comment.IsSet()
}

// HasComment returns a boolean if a field has been set.
func (o *InlineObject487) HasComment() bool {
	if o != nil && o.Comment.IsSet() {
		return true
	}

	return false
}

// SetComment gets a reference to the given NullableString and assigns it to the Comment field.
func (o *InlineObject487) SetComment(v string) {
	o.Comment.Set(&v)
}
// SetCommentNil sets the value for Comment to be an explicit nil
func (o *InlineObject487) SetCommentNil() {
	o.Comment.Set(nil)
}

// UnsetComment ensures that no value is present for Comment, not even an explicit nil
func (o *InlineObject487) UnsetComment() {
	o.Comment.Unset()
}

func (o InlineObject487) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Comment.IsSet() {
		toSerialize["Comment"] = o.Comment.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject487 struct {
	value *InlineObject487
	isSet bool
}

func (v NullableInlineObject487) Get() *InlineObject487 {
	return v.value
}

func (v *NullableInlineObject487) Set(val *InlineObject487) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject487) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject487) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject487(val *InlineObject487) *NullableInlineObject487 {
	return &NullableInlineObject487{value: val, isSet: true}
}

func (v NullableInlineObject487) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject487) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


