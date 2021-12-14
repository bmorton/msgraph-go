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

// InlineObject579 struct for InlineObject579
type InlineObject579 struct {
	Message AnyOfmicrosoftGraphMessage `json:"Message,omitempty"`
	Comment NullableString `json:"Comment,omitempty"`
}

// NewInlineObject579 instantiates a new InlineObject579 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject579() *InlineObject579 {
	this := InlineObject579{}
	return &this
}

// NewInlineObject579WithDefaults instantiates a new InlineObject579 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject579WithDefaults() *InlineObject579 {
	this := InlineObject579{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject579) GetMessage() AnyOfmicrosoftGraphMessage {
	if o == nil  {
		var ret AnyOfmicrosoftGraphMessage
		return ret
	}
	return o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject579) GetMessageOk() (*AnyOfmicrosoftGraphMessage, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return &o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *InlineObject579) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given AnyOfmicrosoftGraphMessage and assigns it to the Message field.
func (o *InlineObject579) SetMessage(v AnyOfmicrosoftGraphMessage) {
	o.Message = v
}

// GetComment returns the Comment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject579) GetComment() string {
	if o == nil || o.Comment.Get() == nil {
		var ret string
		return ret
	}
	return *o.Comment.Get()
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject579) GetCommentOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Comment.Get(), o.Comment.IsSet()
}

// HasComment returns a boolean if a field has been set.
func (o *InlineObject579) HasComment() bool {
	if o != nil && o.Comment.IsSet() {
		return true
	}

	return false
}

// SetComment gets a reference to the given NullableString and assigns it to the Comment field.
func (o *InlineObject579) SetComment(v string) {
	o.Comment.Set(&v)
}
// SetCommentNil sets the value for Comment to be an explicit nil
func (o *InlineObject579) SetCommentNil() {
	o.Comment.Set(nil)
}

// UnsetComment ensures that no value is present for Comment, not even an explicit nil
func (o *InlineObject579) UnsetComment() {
	o.Comment.Unset()
}

func (o InlineObject579) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Message != nil {
		toSerialize["Message"] = o.Message
	}
	if o.Comment.IsSet() {
		toSerialize["Comment"] = o.Comment.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject579 struct {
	value *InlineObject579
	isSet bool
}

func (v NullableInlineObject579) Get() *InlineObject579 {
	return v.value
}

func (v *NullableInlineObject579) Set(val *InlineObject579) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject579) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject579) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject579(val *InlineObject579) *NullableInlineObject579 {
	return &NullableInlineObject579{value: val, isSet: true}
}

func (v NullableInlineObject579) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject579) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


