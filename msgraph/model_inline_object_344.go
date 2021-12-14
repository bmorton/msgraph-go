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

// InlineObject344 struct for InlineObject344
type InlineObject344 struct {
	Comment NullableString `json:"Comment,omitempty"`
	ToRecipients *[]MicrosoftGraphRecipient `json:"ToRecipients,omitempty"`
}

// NewInlineObject344 instantiates a new InlineObject344 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject344() *InlineObject344 {
	this := InlineObject344{}
	return &this
}

// NewInlineObject344WithDefaults instantiates a new InlineObject344 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject344WithDefaults() *InlineObject344 {
	this := InlineObject344{}
	return &this
}

// GetComment returns the Comment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject344) GetComment() string {
	if o == nil || o.Comment.Get() == nil {
		var ret string
		return ret
	}
	return *o.Comment.Get()
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject344) GetCommentOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Comment.Get(), o.Comment.IsSet()
}

// HasComment returns a boolean if a field has been set.
func (o *InlineObject344) HasComment() bool {
	if o != nil && o.Comment.IsSet() {
		return true
	}

	return false
}

// SetComment gets a reference to the given NullableString and assigns it to the Comment field.
func (o *InlineObject344) SetComment(v string) {
	o.Comment.Set(&v)
}
// SetCommentNil sets the value for Comment to be an explicit nil
func (o *InlineObject344) SetCommentNil() {
	o.Comment.Set(nil)
}

// UnsetComment ensures that no value is present for Comment, not even an explicit nil
func (o *InlineObject344) UnsetComment() {
	o.Comment.Unset()
}

// GetToRecipients returns the ToRecipients field value if set, zero value otherwise.
func (o *InlineObject344) GetToRecipients() []MicrosoftGraphRecipient {
	if o == nil || o.ToRecipients == nil {
		var ret []MicrosoftGraphRecipient
		return ret
	}
	return *o.ToRecipients
}

// GetToRecipientsOk returns a tuple with the ToRecipients field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject344) GetToRecipientsOk() (*[]MicrosoftGraphRecipient, bool) {
	if o == nil || o.ToRecipients == nil {
		return nil, false
	}
	return o.ToRecipients, true
}

// HasToRecipients returns a boolean if a field has been set.
func (o *InlineObject344) HasToRecipients() bool {
	if o != nil && o.ToRecipients != nil {
		return true
	}

	return false
}

// SetToRecipients gets a reference to the given []MicrosoftGraphRecipient and assigns it to the ToRecipients field.
func (o *InlineObject344) SetToRecipients(v []MicrosoftGraphRecipient) {
	o.ToRecipients = &v
}

func (o InlineObject344) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Comment.IsSet() {
		toSerialize["Comment"] = o.Comment.Get()
	}
	if o.ToRecipients != nil {
		toSerialize["ToRecipients"] = o.ToRecipients
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject344 struct {
	value *InlineObject344
	isSet bool
}

func (v NullableInlineObject344) Get() *InlineObject344 {
	return v.value
}

func (v *NullableInlineObject344) Set(val *InlineObject344) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject344) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject344) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject344(val *InlineObject344) *NullableInlineObject344 {
	return &NullableInlineObject344{value: val, isSet: true}
}

func (v NullableInlineObject344) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject344) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

