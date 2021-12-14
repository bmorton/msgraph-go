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

// InlineObject722 struct for InlineObject722
type InlineObject722 struct {
	KeyId *string `json:"keyId,omitempty"`
}

// NewInlineObject722 instantiates a new InlineObject722 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject722() *InlineObject722 {
	this := InlineObject722{}
	return &this
}

// NewInlineObject722WithDefaults instantiates a new InlineObject722 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject722WithDefaults() *InlineObject722 {
	this := InlineObject722{}
	return &this
}

// GetKeyId returns the KeyId field value if set, zero value otherwise.
func (o *InlineObject722) GetKeyId() string {
	if o == nil || o.KeyId == nil {
		var ret string
		return ret
	}
	return *o.KeyId
}

// GetKeyIdOk returns a tuple with the KeyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject722) GetKeyIdOk() (*string, bool) {
	if o == nil || o.KeyId == nil {
		return nil, false
	}
	return o.KeyId, true
}

// HasKeyId returns a boolean if a field has been set.
func (o *InlineObject722) HasKeyId() bool {
	if o != nil && o.KeyId != nil {
		return true
	}

	return false
}

// SetKeyId gets a reference to the given string and assigns it to the KeyId field.
func (o *InlineObject722) SetKeyId(v string) {
	o.KeyId = &v
}

func (o InlineObject722) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.KeyId != nil {
		toSerialize["keyId"] = o.KeyId
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject722 struct {
	value *InlineObject722
	isSet bool
}

func (v NullableInlineObject722) Get() *InlineObject722 {
	return v.value
}

func (v *NullableInlineObject722) Set(val *InlineObject722) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject722) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject722) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject722(val *InlineObject722) *NullableInlineObject722 {
	return &NullableInlineObject722{value: val, isSet: true}
}

func (v NullableInlineObject722) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject722) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


