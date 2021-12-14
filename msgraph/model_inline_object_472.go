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

// InlineObject472 struct for InlineObject472
type InlineObject472 struct {
	AttachmentItem *MicrosoftGraphAttachmentItem `json:"AttachmentItem,omitempty"`
}

// NewInlineObject472 instantiates a new InlineObject472 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject472() *InlineObject472 {
	this := InlineObject472{}
	return &this
}

// NewInlineObject472WithDefaults instantiates a new InlineObject472 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject472WithDefaults() *InlineObject472 {
	this := InlineObject472{}
	return &this
}

// GetAttachmentItem returns the AttachmentItem field value if set, zero value otherwise.
func (o *InlineObject472) GetAttachmentItem() MicrosoftGraphAttachmentItem {
	if o == nil || o.AttachmentItem == nil {
		var ret MicrosoftGraphAttachmentItem
		return ret
	}
	return *o.AttachmentItem
}

// GetAttachmentItemOk returns a tuple with the AttachmentItem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject472) GetAttachmentItemOk() (*MicrosoftGraphAttachmentItem, bool) {
	if o == nil || o.AttachmentItem == nil {
		return nil, false
	}
	return o.AttachmentItem, true
}

// HasAttachmentItem returns a boolean if a field has been set.
func (o *InlineObject472) HasAttachmentItem() bool {
	if o != nil && o.AttachmentItem != nil {
		return true
	}

	return false
}

// SetAttachmentItem gets a reference to the given MicrosoftGraphAttachmentItem and assigns it to the AttachmentItem field.
func (o *InlineObject472) SetAttachmentItem(v MicrosoftGraphAttachmentItem) {
	o.AttachmentItem = &v
}

func (o InlineObject472) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AttachmentItem != nil {
		toSerialize["AttachmentItem"] = o.AttachmentItem
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject472 struct {
	value *InlineObject472
	isSet bool
}

func (v NullableInlineObject472) Get() *InlineObject472 {
	return v.value
}

func (v *NullableInlineObject472) Set(val *InlineObject472) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject472) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject472) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject472(val *InlineObject472) *NullableInlineObject472 {
	return &NullableInlineObject472{value: val, isSet: true}
}

func (v NullableInlineObject472) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject472) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


