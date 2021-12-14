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

// MicrosoftGraphChatMessageHostedContent struct for MicrosoftGraphChatMessageHostedContent
type MicrosoftGraphChatMessageHostedContent struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
	// Write only. Bytes for the hosted content (such as images).
	ContentBytes NullableString `json:"contentBytes,omitempty"`
	// Write only. Content type. sicj as image/png, image/jpg.
	ContentType NullableString `json:"contentType,omitempty"`
}

// NewMicrosoftGraphChatMessageHostedContent instantiates a new MicrosoftGraphChatMessageHostedContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphChatMessageHostedContent() *MicrosoftGraphChatMessageHostedContent {
	this := MicrosoftGraphChatMessageHostedContent{}
	return &this
}

// NewMicrosoftGraphChatMessageHostedContentWithDefaults instantiates a new MicrosoftGraphChatMessageHostedContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphChatMessageHostedContentWithDefaults() *MicrosoftGraphChatMessageHostedContent {
	this := MicrosoftGraphChatMessageHostedContent{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MicrosoftGraphChatMessageHostedContent) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphChatMessageHostedContent) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphChatMessageHostedContent) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MicrosoftGraphChatMessageHostedContent) SetId(v string) {
	o.Id = &v
}

// GetContentBytes returns the ContentBytes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphChatMessageHostedContent) GetContentBytes() string {
	if o == nil || o.ContentBytes.Get() == nil {
		var ret string
		return ret
	}
	return *o.ContentBytes.Get()
}

// GetContentBytesOk returns a tuple with the ContentBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphChatMessageHostedContent) GetContentBytesOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ContentBytes.Get(), o.ContentBytes.IsSet()
}

// HasContentBytes returns a boolean if a field has been set.
func (o *MicrosoftGraphChatMessageHostedContent) HasContentBytes() bool {
	if o != nil && o.ContentBytes.IsSet() {
		return true
	}

	return false
}

// SetContentBytes gets a reference to the given NullableString and assigns it to the ContentBytes field.
func (o *MicrosoftGraphChatMessageHostedContent) SetContentBytes(v string) {
	o.ContentBytes.Set(&v)
}
// SetContentBytesNil sets the value for ContentBytes to be an explicit nil
func (o *MicrosoftGraphChatMessageHostedContent) SetContentBytesNil() {
	o.ContentBytes.Set(nil)
}

// UnsetContentBytes ensures that no value is present for ContentBytes, not even an explicit nil
func (o *MicrosoftGraphChatMessageHostedContent) UnsetContentBytes() {
	o.ContentBytes.Unset()
}

// GetContentType returns the ContentType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphChatMessageHostedContent) GetContentType() string {
	if o == nil || o.ContentType.Get() == nil {
		var ret string
		return ret
	}
	return *o.ContentType.Get()
}

// GetContentTypeOk returns a tuple with the ContentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphChatMessageHostedContent) GetContentTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ContentType.Get(), o.ContentType.IsSet()
}

// HasContentType returns a boolean if a field has been set.
func (o *MicrosoftGraphChatMessageHostedContent) HasContentType() bool {
	if o != nil && o.ContentType.IsSet() {
		return true
	}

	return false
}

// SetContentType gets a reference to the given NullableString and assigns it to the ContentType field.
func (o *MicrosoftGraphChatMessageHostedContent) SetContentType(v string) {
	o.ContentType.Set(&v)
}
// SetContentTypeNil sets the value for ContentType to be an explicit nil
func (o *MicrosoftGraphChatMessageHostedContent) SetContentTypeNil() {
	o.ContentType.Set(nil)
}

// UnsetContentType ensures that no value is present for ContentType, not even an explicit nil
func (o *MicrosoftGraphChatMessageHostedContent) UnsetContentType() {
	o.ContentType.Unset()
}

func (o MicrosoftGraphChatMessageHostedContent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.ContentBytes.IsSet() {
		toSerialize["contentBytes"] = o.ContentBytes.Get()
	}
	if o.ContentType.IsSet() {
		toSerialize["contentType"] = o.ContentType.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphChatMessageHostedContent struct {
	value *MicrosoftGraphChatMessageHostedContent
	isSet bool
}

func (v NullableMicrosoftGraphChatMessageHostedContent) Get() *MicrosoftGraphChatMessageHostedContent {
	return v.value
}

func (v *NullableMicrosoftGraphChatMessageHostedContent) Set(val *MicrosoftGraphChatMessageHostedContent) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphChatMessageHostedContent) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphChatMessageHostedContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphChatMessageHostedContent(val *MicrosoftGraphChatMessageHostedContent) *NullableMicrosoftGraphChatMessageHostedContent {
	return &NullableMicrosoftGraphChatMessageHostedContent{value: val, isSet: true}
}

func (v NullableMicrosoftGraphChatMessageHostedContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphChatMessageHostedContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


