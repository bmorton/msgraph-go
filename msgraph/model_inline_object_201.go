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

// InlineObject201 struct for InlineObject201
type InlineObject201 struct {
	ToRecipients *[]*AnyOfmicrosoftGraphRecipient `json:"ToRecipients,omitempty"`
	Comment NullableString `json:"Comment,omitempty"`
}

// NewInlineObject201 instantiates a new InlineObject201 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject201() *InlineObject201 {
	this := InlineObject201{}
	return &this
}

// NewInlineObject201WithDefaults instantiates a new InlineObject201 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject201WithDefaults() *InlineObject201 {
	this := InlineObject201{}
	return &this
}

// GetToRecipients returns the ToRecipients field value if set, zero value otherwise.
func (o *InlineObject201) GetToRecipients() []*AnyOfmicrosoftGraphRecipient {
	if o == nil || o.ToRecipients == nil {
		var ret []*AnyOfmicrosoftGraphRecipient
		return ret
	}
	return *o.ToRecipients
}

// GetToRecipientsOk returns a tuple with the ToRecipients field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject201) GetToRecipientsOk() (*[]*AnyOfmicrosoftGraphRecipient, bool) {
	if o == nil || o.ToRecipients == nil {
		return nil, false
	}
	return o.ToRecipients, true
}

// HasToRecipients returns a boolean if a field has been set.
func (o *InlineObject201) HasToRecipients() bool {
	if o != nil && o.ToRecipients != nil {
		return true
	}

	return false
}

// SetToRecipients gets a reference to the given []*AnyOfmicrosoftGraphRecipient and assigns it to the ToRecipients field.
func (o *InlineObject201) SetToRecipients(v []*AnyOfmicrosoftGraphRecipient) {
	o.ToRecipients = &v
}

// GetComment returns the Comment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject201) GetComment() string {
	if o == nil || o.Comment.Get() == nil {
		var ret string
		return ret
	}
	return *o.Comment.Get()
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject201) GetCommentOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Comment.Get(), o.Comment.IsSet()
}

// HasComment returns a boolean if a field has been set.
func (o *InlineObject201) HasComment() bool {
	if o != nil && o.Comment.IsSet() {
		return true
	}

	return false
}

// SetComment gets a reference to the given NullableString and assigns it to the Comment field.
func (o *InlineObject201) SetComment(v string) {
	o.Comment.Set(&v)
}
// SetCommentNil sets the value for Comment to be an explicit nil
func (o *InlineObject201) SetCommentNil() {
	o.Comment.Set(nil)
}

// UnsetComment ensures that no value is present for Comment, not even an explicit nil
func (o *InlineObject201) UnsetComment() {
	o.Comment.Unset()
}

func (o InlineObject201) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ToRecipients != nil {
		toSerialize["ToRecipients"] = o.ToRecipients
	}
	if o.Comment.IsSet() {
		toSerialize["Comment"] = o.Comment.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject201 struct {
	value *InlineObject201
	isSet bool
}

func (v NullableInlineObject201) Get() *InlineObject201 {
	return v.value
}

func (v *NullableInlineObject201) Set(val *InlineObject201) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject201) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject201) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject201(val *InlineObject201) *NullableInlineObject201 {
	return &NullableInlineObject201{value: val, isSet: true}
}

func (v NullableInlineObject201) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject201) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

