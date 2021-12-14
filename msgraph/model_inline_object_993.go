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

// InlineObject993 struct for InlineObject993
type InlineObject993 struct {
	FileEncryptionInfo AnyOfmicrosoftGraphFileEncryptionInfo `json:"fileEncryptionInfo,omitempty"`
}

// NewInlineObject993 instantiates a new InlineObject993 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject993() *InlineObject993 {
	this := InlineObject993{}
	return &this
}

// NewInlineObject993WithDefaults instantiates a new InlineObject993 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject993WithDefaults() *InlineObject993 {
	this := InlineObject993{}
	return &this
}

// GetFileEncryptionInfo returns the FileEncryptionInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject993) GetFileEncryptionInfo() AnyOfmicrosoftGraphFileEncryptionInfo {
	if o == nil  {
		var ret AnyOfmicrosoftGraphFileEncryptionInfo
		return ret
	}
	return o.FileEncryptionInfo
}

// GetFileEncryptionInfoOk returns a tuple with the FileEncryptionInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject993) GetFileEncryptionInfoOk() (*AnyOfmicrosoftGraphFileEncryptionInfo, bool) {
	if o == nil || o.FileEncryptionInfo == nil {
		return nil, false
	}
	return &o.FileEncryptionInfo, true
}

// HasFileEncryptionInfo returns a boolean if a field has been set.
func (o *InlineObject993) HasFileEncryptionInfo() bool {
	if o != nil && o.FileEncryptionInfo != nil {
		return true
	}

	return false
}

// SetFileEncryptionInfo gets a reference to the given AnyOfmicrosoftGraphFileEncryptionInfo and assigns it to the FileEncryptionInfo field.
func (o *InlineObject993) SetFileEncryptionInfo(v AnyOfmicrosoftGraphFileEncryptionInfo) {
	o.FileEncryptionInfo = v
}

func (o InlineObject993) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.FileEncryptionInfo != nil {
		toSerialize["fileEncryptionInfo"] = o.FileEncryptionInfo
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject993 struct {
	value *InlineObject993
	isSet bool
}

func (v NullableInlineObject993) Get() *InlineObject993 {
	return v.value
}

func (v *NullableInlineObject993) Set(val *InlineObject993) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject993) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject993) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject993(val *InlineObject993) *NullableInlineObject993 {
	return &NullableInlineObject993{value: val, isSet: true}
}

func (v NullableInlineObject993) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject993) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


