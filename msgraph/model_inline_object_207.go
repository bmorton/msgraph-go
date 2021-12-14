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

// InlineObject207 struct for InlineObject207
type InlineObject207 struct {
	Post *MicrosoftGraphPost `json:"Post,omitempty"`
}

// NewInlineObject207 instantiates a new InlineObject207 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject207() *InlineObject207 {
	this := InlineObject207{}
	return &this
}

// NewInlineObject207WithDefaults instantiates a new InlineObject207 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject207WithDefaults() *InlineObject207 {
	this := InlineObject207{}
	return &this
}

// GetPost returns the Post field value if set, zero value otherwise.
func (o *InlineObject207) GetPost() MicrosoftGraphPost {
	if o == nil || o.Post == nil {
		var ret MicrosoftGraphPost
		return ret
	}
	return *o.Post
}

// GetPostOk returns a tuple with the Post field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject207) GetPostOk() (*MicrosoftGraphPost, bool) {
	if o == nil || o.Post == nil {
		return nil, false
	}
	return o.Post, true
}

// HasPost returns a boolean if a field has been set.
func (o *InlineObject207) HasPost() bool {
	if o != nil && o.Post != nil {
		return true
	}

	return false
}

// SetPost gets a reference to the given MicrosoftGraphPost and assigns it to the Post field.
func (o *InlineObject207) SetPost(v MicrosoftGraphPost) {
	o.Post = &v
}

func (o InlineObject207) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Post != nil {
		toSerialize["Post"] = o.Post
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject207 struct {
	value *InlineObject207
	isSet bool
}

func (v NullableInlineObject207) Get() *InlineObject207 {
	return v.value
}

func (v *NullableInlineObject207) Set(val *InlineObject207) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject207) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject207) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject207(val *InlineObject207) *NullableInlineObject207 {
	return &NullableInlineObject207{value: val, isSet: true}
}

func (v NullableInlineObject207) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject207) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


