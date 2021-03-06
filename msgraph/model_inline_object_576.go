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

// InlineObject576 struct for InlineObject576
type InlineObject576 struct {
	AttachmentItem *MicrosoftGraphAttachmentItem `json:"AttachmentItem,omitempty"`
}

// NewInlineObject576 instantiates a new InlineObject576 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject576() *InlineObject576 {
	this := InlineObject576{}
	return &this
}

// NewInlineObject576WithDefaults instantiates a new InlineObject576 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject576WithDefaults() *InlineObject576 {
	this := InlineObject576{}
	return &this
}

// GetAttachmentItem returns the AttachmentItem field value if set, zero value otherwise.
func (o *InlineObject576) GetAttachmentItem() MicrosoftGraphAttachmentItem {
	if o == nil || o.AttachmentItem == nil {
		var ret MicrosoftGraphAttachmentItem
		return ret
	}
	return *o.AttachmentItem
}

// GetAttachmentItemOk returns a tuple with the AttachmentItem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject576) GetAttachmentItemOk() (*MicrosoftGraphAttachmentItem, bool) {
	if o == nil || o.AttachmentItem == nil {
		return nil, false
	}
	return o.AttachmentItem, true
}

// HasAttachmentItem returns a boolean if a field has been set.
func (o *InlineObject576) HasAttachmentItem() bool {
	if o != nil && o.AttachmentItem != nil {
		return true
	}

	return false
}

// SetAttachmentItem gets a reference to the given MicrosoftGraphAttachmentItem and assigns it to the AttachmentItem field.
func (o *InlineObject576) SetAttachmentItem(v MicrosoftGraphAttachmentItem) {
	o.AttachmentItem = &v
}

func (o InlineObject576) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AttachmentItem != nil {
		toSerialize["AttachmentItem"] = o.AttachmentItem
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject576 struct {
	value *InlineObject576
	isSet bool
}

func (v NullableInlineObject576) Get() *InlineObject576 {
	return v.value
}

func (v *NullableInlineObject576) Set(val *InlineObject576) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject576) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject576) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject576(val *InlineObject576) *NullableInlineObject576 {
	return &NullableInlineObject576{value: val, isSet: true}
}

func (v NullableInlineObject576) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject576) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


