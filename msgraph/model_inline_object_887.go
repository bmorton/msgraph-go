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

// InlineObject887 struct for InlineObject887
type InlineObject887 struct {
	ToRecipients *[]*AnyOfmicrosoftGraphRecipient `json:"ToRecipients,omitempty"`
	Comment NullableString `json:"Comment,omitempty"`
}

// NewInlineObject887 instantiates a new InlineObject887 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject887() *InlineObject887 {
	this := InlineObject887{}
	return &this
}

// NewInlineObject887WithDefaults instantiates a new InlineObject887 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject887WithDefaults() *InlineObject887 {
	this := InlineObject887{}
	return &this
}

// GetToRecipients returns the ToRecipients field value if set, zero value otherwise.
func (o *InlineObject887) GetToRecipients() []*AnyOfmicrosoftGraphRecipient {
	if o == nil || o.ToRecipients == nil {
		var ret []*AnyOfmicrosoftGraphRecipient
		return ret
	}
	return *o.ToRecipients
}

// GetToRecipientsOk returns a tuple with the ToRecipients field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject887) GetToRecipientsOk() (*[]*AnyOfmicrosoftGraphRecipient, bool) {
	if o == nil || o.ToRecipients == nil {
		return nil, false
	}
	return o.ToRecipients, true
}

// HasToRecipients returns a boolean if a field has been set.
func (o *InlineObject887) HasToRecipients() bool {
	if o != nil && o.ToRecipients != nil {
		return true
	}

	return false
}

// SetToRecipients gets a reference to the given []*AnyOfmicrosoftGraphRecipient and assigns it to the ToRecipients field.
func (o *InlineObject887) SetToRecipients(v []*AnyOfmicrosoftGraphRecipient) {
	o.ToRecipients = &v
}

// GetComment returns the Comment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject887) GetComment() string {
	if o == nil || o.Comment.Get() == nil {
		var ret string
		return ret
	}
	return *o.Comment.Get()
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject887) GetCommentOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Comment.Get(), o.Comment.IsSet()
}

// HasComment returns a boolean if a field has been set.
func (o *InlineObject887) HasComment() bool {
	if o != nil && o.Comment.IsSet() {
		return true
	}

	return false
}

// SetComment gets a reference to the given NullableString and assigns it to the Comment field.
func (o *InlineObject887) SetComment(v string) {
	o.Comment.Set(&v)
}
// SetCommentNil sets the value for Comment to be an explicit nil
func (o *InlineObject887) SetCommentNil() {
	o.Comment.Set(nil)
}

// UnsetComment ensures that no value is present for Comment, not even an explicit nil
func (o *InlineObject887) UnsetComment() {
	o.Comment.Unset()
}

func (o InlineObject887) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ToRecipients != nil {
		toSerialize["ToRecipients"] = o.ToRecipients
	}
	if o.Comment.IsSet() {
		toSerialize["Comment"] = o.Comment.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject887 struct {
	value *InlineObject887
	isSet bool
}

func (v NullableInlineObject887) Get() *InlineObject887 {
	return v.value
}

func (v *NullableInlineObject887) Set(val *InlineObject887) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject887) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject887) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject887(val *InlineObject887) *NullableInlineObject887 {
	return &NullableInlineObject887{value: val, isSet: true}
}

func (v NullableInlineObject887) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject887) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


