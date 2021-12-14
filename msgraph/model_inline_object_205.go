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

// InlineObject205 struct for InlineObject205
type InlineObject205 struct {
	AttachmentItem *MicrosoftGraphAttachmentItem `json:"AttachmentItem,omitempty"`
}

// NewInlineObject205 instantiates a new InlineObject205 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject205() *InlineObject205 {
	this := InlineObject205{}
	return &this
}

// NewInlineObject205WithDefaults instantiates a new InlineObject205 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject205WithDefaults() *InlineObject205 {
	this := InlineObject205{}
	return &this
}

// GetAttachmentItem returns the AttachmentItem field value if set, zero value otherwise.
func (o *InlineObject205) GetAttachmentItem() MicrosoftGraphAttachmentItem {
	if o == nil || o.AttachmentItem == nil {
		var ret MicrosoftGraphAttachmentItem
		return ret
	}
	return *o.AttachmentItem
}

// GetAttachmentItemOk returns a tuple with the AttachmentItem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject205) GetAttachmentItemOk() (*MicrosoftGraphAttachmentItem, bool) {
	if o == nil || o.AttachmentItem == nil {
		return nil, false
	}
	return o.AttachmentItem, true
}

// HasAttachmentItem returns a boolean if a field has been set.
func (o *InlineObject205) HasAttachmentItem() bool {
	if o != nil && o.AttachmentItem != nil {
		return true
	}

	return false
}

// SetAttachmentItem gets a reference to the given MicrosoftGraphAttachmentItem and assigns it to the AttachmentItem field.
func (o *InlineObject205) SetAttachmentItem(v MicrosoftGraphAttachmentItem) {
	o.AttachmentItem = &v
}

func (o InlineObject205) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AttachmentItem != nil {
		toSerialize["AttachmentItem"] = o.AttachmentItem
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject205 struct {
	value *InlineObject205
	isSet bool
}

func (v NullableInlineObject205) Get() *InlineObject205 {
	return v.value
}

func (v *NullableInlineObject205) Set(val *InlineObject205) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject205) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject205) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject205(val *InlineObject205) *NullableInlineObject205 {
	return &NullableInlineObject205{value: val, isSet: true}
}

func (v NullableInlineObject205) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject205) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

