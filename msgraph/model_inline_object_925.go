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

// InlineObject925 struct for InlineObject925
type InlineObject925 struct {
	AttachmentItem *MicrosoftGraphAttachmentItem `json:"AttachmentItem,omitempty"`
}

// NewInlineObject925 instantiates a new InlineObject925 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject925() *InlineObject925 {
	this := InlineObject925{}
	return &this
}

// NewInlineObject925WithDefaults instantiates a new InlineObject925 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject925WithDefaults() *InlineObject925 {
	this := InlineObject925{}
	return &this
}

// GetAttachmentItem returns the AttachmentItem field value if set, zero value otherwise.
func (o *InlineObject925) GetAttachmentItem() MicrosoftGraphAttachmentItem {
	if o == nil || o.AttachmentItem == nil {
		var ret MicrosoftGraphAttachmentItem
		return ret
	}
	return *o.AttachmentItem
}

// GetAttachmentItemOk returns a tuple with the AttachmentItem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject925) GetAttachmentItemOk() (*MicrosoftGraphAttachmentItem, bool) {
	if o == nil || o.AttachmentItem == nil {
		return nil, false
	}
	return o.AttachmentItem, true
}

// HasAttachmentItem returns a boolean if a field has been set.
func (o *InlineObject925) HasAttachmentItem() bool {
	if o != nil && o.AttachmentItem != nil {
		return true
	}

	return false
}

// SetAttachmentItem gets a reference to the given MicrosoftGraphAttachmentItem and assigns it to the AttachmentItem field.
func (o *InlineObject925) SetAttachmentItem(v MicrosoftGraphAttachmentItem) {
	o.AttachmentItem = &v
}

func (o InlineObject925) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AttachmentItem != nil {
		toSerialize["AttachmentItem"] = o.AttachmentItem
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject925 struct {
	value *InlineObject925
	isSet bool
}

func (v NullableInlineObject925) Get() *InlineObject925 {
	return v.value
}

func (v *NullableInlineObject925) Set(val *InlineObject925) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject925) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject925) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject925(val *InlineObject925) *NullableInlineObject925 {
	return &NullableInlineObject925{value: val, isSet: true}
}

func (v NullableInlineObject925) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject925) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


