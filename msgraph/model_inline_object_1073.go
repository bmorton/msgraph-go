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

// InlineObject1073 struct for InlineObject1073
type InlineObject1073 struct {
	Message AnyOfmicrosoftGraphMessage `json:"Message,omitempty"`
	Comment NullableString `json:"Comment,omitempty"`
}

// NewInlineObject1073 instantiates a new InlineObject1073 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject1073() *InlineObject1073 {
	this := InlineObject1073{}
	return &this
}

// NewInlineObject1073WithDefaults instantiates a new InlineObject1073 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject1073WithDefaults() *InlineObject1073 {
	this := InlineObject1073{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1073) GetMessage() AnyOfmicrosoftGraphMessage {
	if o == nil  {
		var ret AnyOfmicrosoftGraphMessage
		return ret
	}
	return o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1073) GetMessageOk() (*AnyOfmicrosoftGraphMessage, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return &o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *InlineObject1073) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given AnyOfmicrosoftGraphMessage and assigns it to the Message field.
func (o *InlineObject1073) SetMessage(v AnyOfmicrosoftGraphMessage) {
	o.Message = v
}

// GetComment returns the Comment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1073) GetComment() string {
	if o == nil || o.Comment.Get() == nil {
		var ret string
		return ret
	}
	return *o.Comment.Get()
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1073) GetCommentOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Comment.Get(), o.Comment.IsSet()
}

// HasComment returns a boolean if a field has been set.
func (o *InlineObject1073) HasComment() bool {
	if o != nil && o.Comment.IsSet() {
		return true
	}

	return false
}

// SetComment gets a reference to the given NullableString and assigns it to the Comment field.
func (o *InlineObject1073) SetComment(v string) {
	o.Comment.Set(&v)
}
// SetCommentNil sets the value for Comment to be an explicit nil
func (o *InlineObject1073) SetCommentNil() {
	o.Comment.Set(nil)
}

// UnsetComment ensures that no value is present for Comment, not even an explicit nil
func (o *InlineObject1073) UnsetComment() {
	o.Comment.Unset()
}

func (o InlineObject1073) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Message != nil {
		toSerialize["Message"] = o.Message
	}
	if o.Comment.IsSet() {
		toSerialize["Comment"] = o.Comment.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject1073 struct {
	value *InlineObject1073
	isSet bool
}

func (v NullableInlineObject1073) Get() *InlineObject1073 {
	return v.value
}

func (v *NullableInlineObject1073) Set(val *InlineObject1073) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject1073) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject1073) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject1073(val *InlineObject1073) *NullableInlineObject1073 {
	return &NullableInlineObject1073{value: val, isSet: true}
}

func (v NullableInlineObject1073) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject1073) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


