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

// InlineObject407 struct for InlineObject407
type InlineObject407 struct {
	ToRecipients *[]*AnyOfmicrosoftGraphRecipient `json:"ToRecipients,omitempty"`
	Comment NullableString `json:"Comment,omitempty"`
}

// NewInlineObject407 instantiates a new InlineObject407 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject407() *InlineObject407 {
	this := InlineObject407{}
	return &this
}

// NewInlineObject407WithDefaults instantiates a new InlineObject407 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject407WithDefaults() *InlineObject407 {
	this := InlineObject407{}
	return &this
}

// GetToRecipients returns the ToRecipients field value if set, zero value otherwise.
func (o *InlineObject407) GetToRecipients() []*AnyOfmicrosoftGraphRecipient {
	if o == nil || o.ToRecipients == nil {
		var ret []*AnyOfmicrosoftGraphRecipient
		return ret
	}
	return *o.ToRecipients
}

// GetToRecipientsOk returns a tuple with the ToRecipients field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject407) GetToRecipientsOk() (*[]*AnyOfmicrosoftGraphRecipient, bool) {
	if o == nil || o.ToRecipients == nil {
		return nil, false
	}
	return o.ToRecipients, true
}

// HasToRecipients returns a boolean if a field has been set.
func (o *InlineObject407) HasToRecipients() bool {
	if o != nil && o.ToRecipients != nil {
		return true
	}

	return false
}

// SetToRecipients gets a reference to the given []*AnyOfmicrosoftGraphRecipient and assigns it to the ToRecipients field.
func (o *InlineObject407) SetToRecipients(v []*AnyOfmicrosoftGraphRecipient) {
	o.ToRecipients = &v
}

// GetComment returns the Comment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject407) GetComment() string {
	if o == nil || o.Comment.Get() == nil {
		var ret string
		return ret
	}
	return *o.Comment.Get()
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject407) GetCommentOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Comment.Get(), o.Comment.IsSet()
}

// HasComment returns a boolean if a field has been set.
func (o *InlineObject407) HasComment() bool {
	if o != nil && o.Comment.IsSet() {
		return true
	}

	return false
}

// SetComment gets a reference to the given NullableString and assigns it to the Comment field.
func (o *InlineObject407) SetComment(v string) {
	o.Comment.Set(&v)
}
// SetCommentNil sets the value for Comment to be an explicit nil
func (o *InlineObject407) SetCommentNil() {
	o.Comment.Set(nil)
}

// UnsetComment ensures that no value is present for Comment, not even an explicit nil
func (o *InlineObject407) UnsetComment() {
	o.Comment.Unset()
}

func (o InlineObject407) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ToRecipients != nil {
		toSerialize["ToRecipients"] = o.ToRecipients
	}
	if o.Comment.IsSet() {
		toSerialize["Comment"] = o.Comment.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject407 struct {
	value *InlineObject407
	isSet bool
}

func (v NullableInlineObject407) Get() *InlineObject407 {
	return v.value
}

func (v *NullableInlineObject407) Set(val *InlineObject407) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject407) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject407) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject407(val *InlineObject407) *NullableInlineObject407 {
	return &NullableInlineObject407{value: val, isSet: true}
}

func (v NullableInlineObject407) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject407) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


