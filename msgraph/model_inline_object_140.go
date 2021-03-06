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

// InlineObject140 struct for InlineObject140
type InlineObject140 struct {
	ContentType *string `json:"contentType,omitempty"`
}

// NewInlineObject140 instantiates a new InlineObject140 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject140() *InlineObject140 {
	this := InlineObject140{}
	return &this
}

// NewInlineObject140WithDefaults instantiates a new InlineObject140 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject140WithDefaults() *InlineObject140 {
	this := InlineObject140{}
	return &this
}

// GetContentType returns the ContentType field value if set, zero value otherwise.
func (o *InlineObject140) GetContentType() string {
	if o == nil || o.ContentType == nil {
		var ret string
		return ret
	}
	return *o.ContentType
}

// GetContentTypeOk returns a tuple with the ContentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject140) GetContentTypeOk() (*string, bool) {
	if o == nil || o.ContentType == nil {
		return nil, false
	}
	return o.ContentType, true
}

// HasContentType returns a boolean if a field has been set.
func (o *InlineObject140) HasContentType() bool {
	if o != nil && o.ContentType != nil {
		return true
	}

	return false
}

// SetContentType gets a reference to the given string and assigns it to the ContentType field.
func (o *InlineObject140) SetContentType(v string) {
	o.ContentType = &v
}

func (o InlineObject140) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ContentType != nil {
		toSerialize["contentType"] = o.ContentType
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject140 struct {
	value *InlineObject140
	isSet bool
}

func (v NullableInlineObject140) Get() *InlineObject140 {
	return v.value
}

func (v *NullableInlineObject140) Set(val *InlineObject140) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject140) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject140) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject140(val *InlineObject140) *NullableInlineObject140 {
	return &NullableInlineObject140{value: val, isSet: true}
}

func (v NullableInlineObject140) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject140) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


