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

// InlineObject899 struct for InlineObject899
type InlineObject899 struct {
	Comment NullableString `json:"Comment,omitempty"`
}

// NewInlineObject899 instantiates a new InlineObject899 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject899() *InlineObject899 {
	this := InlineObject899{}
	return &this
}

// NewInlineObject899WithDefaults instantiates a new InlineObject899 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject899WithDefaults() *InlineObject899 {
	this := InlineObject899{}
	return &this
}

// GetComment returns the Comment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject899) GetComment() string {
	if o == nil || o.Comment.Get() == nil {
		var ret string
		return ret
	}
	return *o.Comment.Get()
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject899) GetCommentOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Comment.Get(), o.Comment.IsSet()
}

// HasComment returns a boolean if a field has been set.
func (o *InlineObject899) HasComment() bool {
	if o != nil && o.Comment.IsSet() {
		return true
	}

	return false
}

// SetComment gets a reference to the given NullableString and assigns it to the Comment field.
func (o *InlineObject899) SetComment(v string) {
	o.Comment.Set(&v)
}
// SetCommentNil sets the value for Comment to be an explicit nil
func (o *InlineObject899) SetCommentNil() {
	o.Comment.Set(nil)
}

// UnsetComment ensures that no value is present for Comment, not even an explicit nil
func (o *InlineObject899) UnsetComment() {
	o.Comment.Unset()
}

func (o InlineObject899) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Comment.IsSet() {
		toSerialize["Comment"] = o.Comment.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject899 struct {
	value *InlineObject899
	isSet bool
}

func (v NullableInlineObject899) Get() *InlineObject899 {
	return v.value
}

func (v *NullableInlineObject899) Set(val *InlineObject899) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject899) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject899) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject899(val *InlineObject899) *NullableInlineObject899 {
	return &NullableInlineObject899{value: val, isSet: true}
}

func (v NullableInlineObject899) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject899) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


