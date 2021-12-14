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

// MicrosoftGraphWorkbookCommentReply struct for MicrosoftGraphWorkbookCommentReply
type MicrosoftGraphWorkbookCommentReply struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
	// The content of a comment reply.
	Content NullableString `json:"content,omitempty"`
	// Indicates the type for the comment reply.
	ContentType *string `json:"contentType,omitempty"`
}

// NewMicrosoftGraphWorkbookCommentReply instantiates a new MicrosoftGraphWorkbookCommentReply object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphWorkbookCommentReply() *MicrosoftGraphWorkbookCommentReply {
	this := MicrosoftGraphWorkbookCommentReply{}
	return &this
}

// NewMicrosoftGraphWorkbookCommentReplyWithDefaults instantiates a new MicrosoftGraphWorkbookCommentReply object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphWorkbookCommentReplyWithDefaults() *MicrosoftGraphWorkbookCommentReply {
	this := MicrosoftGraphWorkbookCommentReply{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MicrosoftGraphWorkbookCommentReply) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphWorkbookCommentReply) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphWorkbookCommentReply) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MicrosoftGraphWorkbookCommentReply) SetId(v string) {
	o.Id = &v
}

// GetContent returns the Content field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphWorkbookCommentReply) GetContent() string {
	if o == nil || o.Content.Get() == nil {
		var ret string
		return ret
	}
	return *o.Content.Get()
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphWorkbookCommentReply) GetContentOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Content.Get(), o.Content.IsSet()
}

// HasContent returns a boolean if a field has been set.
func (o *MicrosoftGraphWorkbookCommentReply) HasContent() bool {
	if o != nil && o.Content.IsSet() {
		return true
	}

	return false
}

// SetContent gets a reference to the given NullableString and assigns it to the Content field.
func (o *MicrosoftGraphWorkbookCommentReply) SetContent(v string) {
	o.Content.Set(&v)
}
// SetContentNil sets the value for Content to be an explicit nil
func (o *MicrosoftGraphWorkbookCommentReply) SetContentNil() {
	o.Content.Set(nil)
}

// UnsetContent ensures that no value is present for Content, not even an explicit nil
func (o *MicrosoftGraphWorkbookCommentReply) UnsetContent() {
	o.Content.Unset()
}

// GetContentType returns the ContentType field value if set, zero value otherwise.
func (o *MicrosoftGraphWorkbookCommentReply) GetContentType() string {
	if o == nil || o.ContentType == nil {
		var ret string
		return ret
	}
	return *o.ContentType
}

// GetContentTypeOk returns a tuple with the ContentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphWorkbookCommentReply) GetContentTypeOk() (*string, bool) {
	if o == nil || o.ContentType == nil {
		return nil, false
	}
	return o.ContentType, true
}

// HasContentType returns a boolean if a field has been set.
func (o *MicrosoftGraphWorkbookCommentReply) HasContentType() bool {
	if o != nil && o.ContentType != nil {
		return true
	}

	return false
}

// SetContentType gets a reference to the given string and assigns it to the ContentType field.
func (o *MicrosoftGraphWorkbookCommentReply) SetContentType(v string) {
	o.ContentType = &v
}

func (o MicrosoftGraphWorkbookCommentReply) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Content.IsSet() {
		toSerialize["content"] = o.Content.Get()
	}
	if o.ContentType != nil {
		toSerialize["contentType"] = o.ContentType
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphWorkbookCommentReply struct {
	value *MicrosoftGraphWorkbookCommentReply
	isSet bool
}

func (v NullableMicrosoftGraphWorkbookCommentReply) Get() *MicrosoftGraphWorkbookCommentReply {
	return v.value
}

func (v *NullableMicrosoftGraphWorkbookCommentReply) Set(val *MicrosoftGraphWorkbookCommentReply) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphWorkbookCommentReply) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphWorkbookCommentReply) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphWorkbookCommentReply(val *MicrosoftGraphWorkbookCommentReply) *NullableMicrosoftGraphWorkbookCommentReply {
	return &NullableMicrosoftGraphWorkbookCommentReply{value: val, isSet: true}
}

func (v NullableMicrosoftGraphWorkbookCommentReply) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphWorkbookCommentReply) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


