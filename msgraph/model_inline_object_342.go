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

// InlineObject342 struct for InlineObject342
type InlineObject342 struct {
	Comment NullableString `json:"Comment,omitempty"`
	ToRecipients *[]MicrosoftGraphRecipient `json:"ToRecipients,omitempty"`
}

// NewInlineObject342 instantiates a new InlineObject342 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject342() *InlineObject342 {
	this := InlineObject342{}
	return &this
}

// NewInlineObject342WithDefaults instantiates a new InlineObject342 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject342WithDefaults() *InlineObject342 {
	this := InlineObject342{}
	return &this
}

// GetComment returns the Comment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject342) GetComment() string {
	if o == nil || o.Comment.Get() == nil {
		var ret string
		return ret
	}
	return *o.Comment.Get()
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject342) GetCommentOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Comment.Get(), o.Comment.IsSet()
}

// HasComment returns a boolean if a field has been set.
func (o *InlineObject342) HasComment() bool {
	if o != nil && o.Comment.IsSet() {
		return true
	}

	return false
}

// SetComment gets a reference to the given NullableString and assigns it to the Comment field.
func (o *InlineObject342) SetComment(v string) {
	o.Comment.Set(&v)
}
// SetCommentNil sets the value for Comment to be an explicit nil
func (o *InlineObject342) SetCommentNil() {
	o.Comment.Set(nil)
}

// UnsetComment ensures that no value is present for Comment, not even an explicit nil
func (o *InlineObject342) UnsetComment() {
	o.Comment.Unset()
}

// GetToRecipients returns the ToRecipients field value if set, zero value otherwise.
func (o *InlineObject342) GetToRecipients() []MicrosoftGraphRecipient {
	if o == nil || o.ToRecipients == nil {
		var ret []MicrosoftGraphRecipient
		return ret
	}
	return *o.ToRecipients
}

// GetToRecipientsOk returns a tuple with the ToRecipients field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject342) GetToRecipientsOk() (*[]MicrosoftGraphRecipient, bool) {
	if o == nil || o.ToRecipients == nil {
		return nil, false
	}
	return o.ToRecipients, true
}

// HasToRecipients returns a boolean if a field has been set.
func (o *InlineObject342) HasToRecipients() bool {
	if o != nil && o.ToRecipients != nil {
		return true
	}

	return false
}

// SetToRecipients gets a reference to the given []MicrosoftGraphRecipient and assigns it to the ToRecipients field.
func (o *InlineObject342) SetToRecipients(v []MicrosoftGraphRecipient) {
	o.ToRecipients = &v
}

func (o InlineObject342) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Comment.IsSet() {
		toSerialize["Comment"] = o.Comment.Get()
	}
	if o.ToRecipients != nil {
		toSerialize["ToRecipients"] = o.ToRecipients
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject342 struct {
	value *InlineObject342
	isSet bool
}

func (v NullableInlineObject342) Get() *InlineObject342 {
	return v.value
}

func (v *NullableInlineObject342) Set(val *InlineObject342) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject342) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject342) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject342(val *InlineObject342) *NullableInlineObject342 {
	return &NullableInlineObject342{value: val, isSet: true}
}

func (v NullableInlineObject342) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject342) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


