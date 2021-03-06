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

// InlineObject896 struct for InlineObject896
type InlineObject896 struct {
	AttachmentItem *MicrosoftGraphAttachmentItem `json:"AttachmentItem,omitempty"`
}

// NewInlineObject896 instantiates a new InlineObject896 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject896() *InlineObject896 {
	this := InlineObject896{}
	return &this
}

// NewInlineObject896WithDefaults instantiates a new InlineObject896 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject896WithDefaults() *InlineObject896 {
	this := InlineObject896{}
	return &this
}

// GetAttachmentItem returns the AttachmentItem field value if set, zero value otherwise.
func (o *InlineObject896) GetAttachmentItem() MicrosoftGraphAttachmentItem {
	if o == nil || o.AttachmentItem == nil {
		var ret MicrosoftGraphAttachmentItem
		return ret
	}
	return *o.AttachmentItem
}

// GetAttachmentItemOk returns a tuple with the AttachmentItem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject896) GetAttachmentItemOk() (*MicrosoftGraphAttachmentItem, bool) {
	if o == nil || o.AttachmentItem == nil {
		return nil, false
	}
	return o.AttachmentItem, true
}

// HasAttachmentItem returns a boolean if a field has been set.
func (o *InlineObject896) HasAttachmentItem() bool {
	if o != nil && o.AttachmentItem != nil {
		return true
	}

	return false
}

// SetAttachmentItem gets a reference to the given MicrosoftGraphAttachmentItem and assigns it to the AttachmentItem field.
func (o *InlineObject896) SetAttachmentItem(v MicrosoftGraphAttachmentItem) {
	o.AttachmentItem = &v
}

func (o InlineObject896) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AttachmentItem != nil {
		toSerialize["AttachmentItem"] = o.AttachmentItem
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject896 struct {
	value *InlineObject896
	isSet bool
}

func (v NullableInlineObject896) Get() *InlineObject896 {
	return v.value
}

func (v *NullableInlineObject896) Set(val *InlineObject896) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject896) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject896) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject896(val *InlineObject896) *NullableInlineObject896 {
	return &NullableInlineObject896{value: val, isSet: true}
}

func (v NullableInlineObject896) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject896) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


