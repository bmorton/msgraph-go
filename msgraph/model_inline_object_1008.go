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

// InlineObject1008 struct for InlineObject1008
type InlineObject1008 struct {
	FileEncryptionInfo AnyOfmicrosoftGraphFileEncryptionInfo `json:"fileEncryptionInfo,omitempty"`
}

// NewInlineObject1008 instantiates a new InlineObject1008 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject1008() *InlineObject1008 {
	this := InlineObject1008{}
	return &this
}

// NewInlineObject1008WithDefaults instantiates a new InlineObject1008 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject1008WithDefaults() *InlineObject1008 {
	this := InlineObject1008{}
	return &this
}

// GetFileEncryptionInfo returns the FileEncryptionInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1008) GetFileEncryptionInfo() AnyOfmicrosoftGraphFileEncryptionInfo {
	if o == nil  {
		var ret AnyOfmicrosoftGraphFileEncryptionInfo
		return ret
	}
	return o.FileEncryptionInfo
}

// GetFileEncryptionInfoOk returns a tuple with the FileEncryptionInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1008) GetFileEncryptionInfoOk() (*AnyOfmicrosoftGraphFileEncryptionInfo, bool) {
	if o == nil || o.FileEncryptionInfo == nil {
		return nil, false
	}
	return &o.FileEncryptionInfo, true
}

// HasFileEncryptionInfo returns a boolean if a field has been set.
func (o *InlineObject1008) HasFileEncryptionInfo() bool {
	if o != nil && o.FileEncryptionInfo != nil {
		return true
	}

	return false
}

// SetFileEncryptionInfo gets a reference to the given AnyOfmicrosoftGraphFileEncryptionInfo and assigns it to the FileEncryptionInfo field.
func (o *InlineObject1008) SetFileEncryptionInfo(v AnyOfmicrosoftGraphFileEncryptionInfo) {
	o.FileEncryptionInfo = v
}

func (o InlineObject1008) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.FileEncryptionInfo != nil {
		toSerialize["fileEncryptionInfo"] = o.FileEncryptionInfo
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject1008 struct {
	value *InlineObject1008
	isSet bool
}

func (v NullableInlineObject1008) Get() *InlineObject1008 {
	return v.value
}

func (v *NullableInlineObject1008) Set(val *InlineObject1008) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject1008) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject1008) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject1008(val *InlineObject1008) *NullableInlineObject1008 {
	return &NullableInlineObject1008{value: val, isSet: true}
}

func (v NullableInlineObject1008) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject1008) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

