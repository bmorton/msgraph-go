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

// InlineObject1199 struct for InlineObject1199
type InlineObject1199 struct {
	Item AnyOfmicrosoftGraphDriveItemUploadableProperties `json:"item,omitempty"`
}

// NewInlineObject1199 instantiates a new InlineObject1199 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject1199() *InlineObject1199 {
	this := InlineObject1199{}
	return &this
}

// NewInlineObject1199WithDefaults instantiates a new InlineObject1199 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject1199WithDefaults() *InlineObject1199 {
	this := InlineObject1199{}
	return &this
}

// GetItem returns the Item field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1199) GetItem() AnyOfmicrosoftGraphDriveItemUploadableProperties {
	if o == nil  {
		var ret AnyOfmicrosoftGraphDriveItemUploadableProperties
		return ret
	}
	return o.Item
}

// GetItemOk returns a tuple with the Item field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1199) GetItemOk() (*AnyOfmicrosoftGraphDriveItemUploadableProperties, bool) {
	if o == nil || o.Item == nil {
		return nil, false
	}
	return &o.Item, true
}

// HasItem returns a boolean if a field has been set.
func (o *InlineObject1199) HasItem() bool {
	if o != nil && o.Item != nil {
		return true
	}

	return false
}

// SetItem gets a reference to the given AnyOfmicrosoftGraphDriveItemUploadableProperties and assigns it to the Item field.
func (o *InlineObject1199) SetItem(v AnyOfmicrosoftGraphDriveItemUploadableProperties) {
	o.Item = v
}

func (o InlineObject1199) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Item != nil {
		toSerialize["item"] = o.Item
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject1199 struct {
	value *InlineObject1199
	isSet bool
}

func (v NullableInlineObject1199) Get() *InlineObject1199 {
	return v.value
}

func (v *NullableInlineObject1199) Set(val *InlineObject1199) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject1199) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject1199) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject1199(val *InlineObject1199) *NullableInlineObject1199 {
	return &NullableInlineObject1199{value: val, isSet: true}
}

func (v NullableInlineObject1199) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject1199) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

