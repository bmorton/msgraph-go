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

// InlineObject182 struct for InlineObject182
type InlineObject182 struct {
	ToRecipients *[]*AnyOfmicrosoftGraphRecipient `json:"ToRecipients,omitempty"`
	Comment NullableString `json:"Comment,omitempty"`
}

// NewInlineObject182 instantiates a new InlineObject182 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject182() *InlineObject182 {
	this := InlineObject182{}
	return &this
}

// NewInlineObject182WithDefaults instantiates a new InlineObject182 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject182WithDefaults() *InlineObject182 {
	this := InlineObject182{}
	return &this
}

// GetToRecipients returns the ToRecipients field value if set, zero value otherwise.
func (o *InlineObject182) GetToRecipients() []*AnyOfmicrosoftGraphRecipient {
	if o == nil || o.ToRecipients == nil {
		var ret []*AnyOfmicrosoftGraphRecipient
		return ret
	}
	return *o.ToRecipients
}

// GetToRecipientsOk returns a tuple with the ToRecipients field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject182) GetToRecipientsOk() (*[]*AnyOfmicrosoftGraphRecipient, bool) {
	if o == nil || o.ToRecipients == nil {
		return nil, false
	}
	return o.ToRecipients, true
}

// HasToRecipients returns a boolean if a field has been set.
func (o *InlineObject182) HasToRecipients() bool {
	if o != nil && o.ToRecipients != nil {
		return true
	}

	return false
}

// SetToRecipients gets a reference to the given []*AnyOfmicrosoftGraphRecipient and assigns it to the ToRecipients field.
func (o *InlineObject182) SetToRecipients(v []*AnyOfmicrosoftGraphRecipient) {
	o.ToRecipients = &v
}

// GetComment returns the Comment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject182) GetComment() string {
	if o == nil || o.Comment.Get() == nil {
		var ret string
		return ret
	}
	return *o.Comment.Get()
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject182) GetCommentOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Comment.Get(), o.Comment.IsSet()
}

// HasComment returns a boolean if a field has been set.
func (o *InlineObject182) HasComment() bool {
	if o != nil && o.Comment.IsSet() {
		return true
	}

	return false
}

// SetComment gets a reference to the given NullableString and assigns it to the Comment field.
func (o *InlineObject182) SetComment(v string) {
	o.Comment.Set(&v)
}
// SetCommentNil sets the value for Comment to be an explicit nil
func (o *InlineObject182) SetCommentNil() {
	o.Comment.Set(nil)
}

// UnsetComment ensures that no value is present for Comment, not even an explicit nil
func (o *InlineObject182) UnsetComment() {
	o.Comment.Unset()
}

func (o InlineObject182) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ToRecipients != nil {
		toSerialize["ToRecipients"] = o.ToRecipients
	}
	if o.Comment.IsSet() {
		toSerialize["Comment"] = o.Comment.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject182 struct {
	value *InlineObject182
	isSet bool
}

func (v NullableInlineObject182) Get() *InlineObject182 {
	return v.value
}

func (v *NullableInlineObject182) Set(val *InlineObject182) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject182) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject182) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject182(val *InlineObject182) *NullableInlineObject182 {
	return &NullableInlineObject182{value: val, isSet: true}
}

func (v NullableInlineObject182) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject182) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


