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

// InlineObject1062 struct for InlineObject1062
type InlineObject1062 struct {
	Message AnyOfmicrosoftGraphMessage `json:"Message,omitempty"`
	Comment NullableString `json:"Comment,omitempty"`
}

// NewInlineObject1062 instantiates a new InlineObject1062 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject1062() *InlineObject1062 {
	this := InlineObject1062{}
	return &this
}

// NewInlineObject1062WithDefaults instantiates a new InlineObject1062 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject1062WithDefaults() *InlineObject1062 {
	this := InlineObject1062{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1062) GetMessage() AnyOfmicrosoftGraphMessage {
	if o == nil  {
		var ret AnyOfmicrosoftGraphMessage
		return ret
	}
	return o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1062) GetMessageOk() (*AnyOfmicrosoftGraphMessage, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return &o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *InlineObject1062) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given AnyOfmicrosoftGraphMessage and assigns it to the Message field.
func (o *InlineObject1062) SetMessage(v AnyOfmicrosoftGraphMessage) {
	o.Message = v
}

// GetComment returns the Comment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1062) GetComment() string {
	if o == nil || o.Comment.Get() == nil {
		var ret string
		return ret
	}
	return *o.Comment.Get()
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1062) GetCommentOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Comment.Get(), o.Comment.IsSet()
}

// HasComment returns a boolean if a field has been set.
func (o *InlineObject1062) HasComment() bool {
	if o != nil && o.Comment.IsSet() {
		return true
	}

	return false
}

// SetComment gets a reference to the given NullableString and assigns it to the Comment field.
func (o *InlineObject1062) SetComment(v string) {
	o.Comment.Set(&v)
}
// SetCommentNil sets the value for Comment to be an explicit nil
func (o *InlineObject1062) SetCommentNil() {
	o.Comment.Set(nil)
}

// UnsetComment ensures that no value is present for Comment, not even an explicit nil
func (o *InlineObject1062) UnsetComment() {
	o.Comment.Unset()
}

func (o InlineObject1062) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Message != nil {
		toSerialize["Message"] = o.Message
	}
	if o.Comment.IsSet() {
		toSerialize["Comment"] = o.Comment.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject1062 struct {
	value *InlineObject1062
	isSet bool
}

func (v NullableInlineObject1062) Get() *InlineObject1062 {
	return v.value
}

func (v *NullableInlineObject1062) Set(val *InlineObject1062) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject1062) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject1062) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject1062(val *InlineObject1062) *NullableInlineObject1062 {
	return &NullableInlineObject1062{value: val, isSet: true}
}

func (v NullableInlineObject1062) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject1062) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


