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

// InlineObject914 struct for InlineObject914
type InlineObject914 struct {
	Comment NullableString `json:"Comment,omitempty"`
}

// NewInlineObject914 instantiates a new InlineObject914 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject914() *InlineObject914 {
	this := InlineObject914{}
	return &this
}

// NewInlineObject914WithDefaults instantiates a new InlineObject914 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject914WithDefaults() *InlineObject914 {
	this := InlineObject914{}
	return &this
}

// GetComment returns the Comment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject914) GetComment() string {
	if o == nil || o.Comment.Get() == nil {
		var ret string
		return ret
	}
	return *o.Comment.Get()
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject914) GetCommentOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Comment.Get(), o.Comment.IsSet()
}

// HasComment returns a boolean if a field has been set.
func (o *InlineObject914) HasComment() bool {
	if o != nil && o.Comment.IsSet() {
		return true
	}

	return false
}

// SetComment gets a reference to the given NullableString and assigns it to the Comment field.
func (o *InlineObject914) SetComment(v string) {
	o.Comment.Set(&v)
}
// SetCommentNil sets the value for Comment to be an explicit nil
func (o *InlineObject914) SetCommentNil() {
	o.Comment.Set(nil)
}

// UnsetComment ensures that no value is present for Comment, not even an explicit nil
func (o *InlineObject914) UnsetComment() {
	o.Comment.Unset()
}

func (o InlineObject914) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Comment.IsSet() {
		toSerialize["Comment"] = o.Comment.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject914 struct {
	value *InlineObject914
	isSet bool
}

func (v NullableInlineObject914) Get() *InlineObject914 {
	return v.value
}

func (v *NullableInlineObject914) Set(val *InlineObject914) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject914) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject914) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject914(val *InlineObject914) *NullableInlineObject914 {
	return &NullableInlineObject914{value: val, isSet: true}
}

func (v NullableInlineObject914) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject914) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


