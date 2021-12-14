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

// InlineObject731 struct for InlineObject731
type InlineObject731 struct {
	ContentType *string `json:"contentType,omitempty"`
}

// NewInlineObject731 instantiates a new InlineObject731 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject731() *InlineObject731 {
	this := InlineObject731{}
	return &this
}

// NewInlineObject731WithDefaults instantiates a new InlineObject731 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject731WithDefaults() *InlineObject731 {
	this := InlineObject731{}
	return &this
}

// GetContentType returns the ContentType field value if set, zero value otherwise.
func (o *InlineObject731) GetContentType() string {
	if o == nil || o.ContentType == nil {
		var ret string
		return ret
	}
	return *o.ContentType
}

// GetContentTypeOk returns a tuple with the ContentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject731) GetContentTypeOk() (*string, bool) {
	if o == nil || o.ContentType == nil {
		return nil, false
	}
	return o.ContentType, true
}

// HasContentType returns a boolean if a field has been set.
func (o *InlineObject731) HasContentType() bool {
	if o != nil && o.ContentType != nil {
		return true
	}

	return false
}

// SetContentType gets a reference to the given string and assigns it to the ContentType field.
func (o *InlineObject731) SetContentType(v string) {
	o.ContentType = &v
}

func (o InlineObject731) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ContentType != nil {
		toSerialize["contentType"] = o.ContentType
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject731 struct {
	value *InlineObject731
	isSet bool
}

func (v NullableInlineObject731) Get() *InlineObject731 {
	return v.value
}

func (v *NullableInlineObject731) Set(val *InlineObject731) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject731) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject731) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject731(val *InlineObject731) *NullableInlineObject731 {
	return &NullableInlineObject731{value: val, isSet: true}
}

func (v NullableInlineObject731) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject731) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


