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

// MicrosoftGraphChatMessageAttachment struct for MicrosoftGraphChatMessageAttachment
type MicrosoftGraphChatMessageAttachment struct {
	// The content of the attachment. If the attachment is a rich card, set the property to the rich card object. This property and contentUrl are mutually exclusive.
	Content NullableString `json:"content,omitempty"`
	// The media type of the content attachment. It can have the following values: reference: Attachment is a link to another file. Populate the contentURL with the link to the object.Any contentTypes supported by the Bot Framework's Attachment objectapplication/vnd.microsoft.card.codesnippet: A code snippet. application/vnd.microsoft.card.announcement: An announcement header.
	ContentType NullableString `json:"contentType,omitempty"`
	// URL for the content of the attachment. Supported protocols: http, https, file and data.
	ContentUrl NullableString `json:"contentUrl,omitempty"`
	// Read-only. Unique id of the attachment.
	Id NullableString `json:"id,omitempty"`
	// Name of the attachment.
	Name NullableString `json:"name,omitempty"`
	// URL to a thumbnail image that the channel can use if it supports using an alternative, smaller form of content or contentUrl. For example, if you set contentType to application/word and set contentUrl to the location of the Word document, you might include a thumbnail image that represents the document. The channel could display the thumbnail image instead of the document. When the user clicks the image, the channel would open the document.
	ThumbnailUrl NullableString `json:"thumbnailUrl,omitempty"`
}

// NewMicrosoftGraphChatMessageAttachment instantiates a new MicrosoftGraphChatMessageAttachment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphChatMessageAttachment() *MicrosoftGraphChatMessageAttachment {
	this := MicrosoftGraphChatMessageAttachment{}
	return &this
}

// NewMicrosoftGraphChatMessageAttachmentWithDefaults instantiates a new MicrosoftGraphChatMessageAttachment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphChatMessageAttachmentWithDefaults() *MicrosoftGraphChatMessageAttachment {
	this := MicrosoftGraphChatMessageAttachment{}
	return &this
}

// GetContent returns the Content field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphChatMessageAttachment) GetContent() string {
	if o == nil || o.Content.Get() == nil {
		var ret string
		return ret
	}
	return *o.Content.Get()
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphChatMessageAttachment) GetContentOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Content.Get(), o.Content.IsSet()
}

// HasContent returns a boolean if a field has been set.
func (o *MicrosoftGraphChatMessageAttachment) HasContent() bool {
	if o != nil && o.Content.IsSet() {
		return true
	}

	return false
}

// SetContent gets a reference to the given NullableString and assigns it to the Content field.
func (o *MicrosoftGraphChatMessageAttachment) SetContent(v string) {
	o.Content.Set(&v)
}
// SetContentNil sets the value for Content to be an explicit nil
func (o *MicrosoftGraphChatMessageAttachment) SetContentNil() {
	o.Content.Set(nil)
}

// UnsetContent ensures that no value is present for Content, not even an explicit nil
func (o *MicrosoftGraphChatMessageAttachment) UnsetContent() {
	o.Content.Unset()
}

// GetContentType returns the ContentType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphChatMessageAttachment) GetContentType() string {
	if o == nil || o.ContentType.Get() == nil {
		var ret string
		return ret
	}
	return *o.ContentType.Get()
}

// GetContentTypeOk returns a tuple with the ContentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphChatMessageAttachment) GetContentTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ContentType.Get(), o.ContentType.IsSet()
}

// HasContentType returns a boolean if a field has been set.
func (o *MicrosoftGraphChatMessageAttachment) HasContentType() bool {
	if o != nil && o.ContentType.IsSet() {
		return true
	}

	return false
}

// SetContentType gets a reference to the given NullableString and assigns it to the ContentType field.
func (o *MicrosoftGraphChatMessageAttachment) SetContentType(v string) {
	o.ContentType.Set(&v)
}
// SetContentTypeNil sets the value for ContentType to be an explicit nil
func (o *MicrosoftGraphChatMessageAttachment) SetContentTypeNil() {
	o.ContentType.Set(nil)
}

// UnsetContentType ensures that no value is present for ContentType, not even an explicit nil
func (o *MicrosoftGraphChatMessageAttachment) UnsetContentType() {
	o.ContentType.Unset()
}

// GetContentUrl returns the ContentUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphChatMessageAttachment) GetContentUrl() string {
	if o == nil || o.ContentUrl.Get() == nil {
		var ret string
		return ret
	}
	return *o.ContentUrl.Get()
}

// GetContentUrlOk returns a tuple with the ContentUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphChatMessageAttachment) GetContentUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ContentUrl.Get(), o.ContentUrl.IsSet()
}

// HasContentUrl returns a boolean if a field has been set.
func (o *MicrosoftGraphChatMessageAttachment) HasContentUrl() bool {
	if o != nil && o.ContentUrl.IsSet() {
		return true
	}

	return false
}

// SetContentUrl gets a reference to the given NullableString and assigns it to the ContentUrl field.
func (o *MicrosoftGraphChatMessageAttachment) SetContentUrl(v string) {
	o.ContentUrl.Set(&v)
}
// SetContentUrlNil sets the value for ContentUrl to be an explicit nil
func (o *MicrosoftGraphChatMessageAttachment) SetContentUrlNil() {
	o.ContentUrl.Set(nil)
}

// UnsetContentUrl ensures that no value is present for ContentUrl, not even an explicit nil
func (o *MicrosoftGraphChatMessageAttachment) UnsetContentUrl() {
	o.ContentUrl.Unset()
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphChatMessageAttachment) GetId() string {
	if o == nil || o.Id.Get() == nil {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphChatMessageAttachment) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphChatMessageAttachment) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *MicrosoftGraphChatMessageAttachment) SetId(v string) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *MicrosoftGraphChatMessageAttachment) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *MicrosoftGraphChatMessageAttachment) UnsetId() {
	o.Id.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphChatMessageAttachment) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphChatMessageAttachment) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *MicrosoftGraphChatMessageAttachment) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *MicrosoftGraphChatMessageAttachment) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *MicrosoftGraphChatMessageAttachment) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *MicrosoftGraphChatMessageAttachment) UnsetName() {
	o.Name.Unset()
}

// GetThumbnailUrl returns the ThumbnailUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphChatMessageAttachment) GetThumbnailUrl() string {
	if o == nil || o.ThumbnailUrl.Get() == nil {
		var ret string
		return ret
	}
	return *o.ThumbnailUrl.Get()
}

// GetThumbnailUrlOk returns a tuple with the ThumbnailUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphChatMessageAttachment) GetThumbnailUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ThumbnailUrl.Get(), o.ThumbnailUrl.IsSet()
}

// HasThumbnailUrl returns a boolean if a field has been set.
func (o *MicrosoftGraphChatMessageAttachment) HasThumbnailUrl() bool {
	if o != nil && o.ThumbnailUrl.IsSet() {
		return true
	}

	return false
}

// SetThumbnailUrl gets a reference to the given NullableString and assigns it to the ThumbnailUrl field.
func (o *MicrosoftGraphChatMessageAttachment) SetThumbnailUrl(v string) {
	o.ThumbnailUrl.Set(&v)
}
// SetThumbnailUrlNil sets the value for ThumbnailUrl to be an explicit nil
func (o *MicrosoftGraphChatMessageAttachment) SetThumbnailUrlNil() {
	o.ThumbnailUrl.Set(nil)
}

// UnsetThumbnailUrl ensures that no value is present for ThumbnailUrl, not even an explicit nil
func (o *MicrosoftGraphChatMessageAttachment) UnsetThumbnailUrl() {
	o.ThumbnailUrl.Unset()
}

func (o MicrosoftGraphChatMessageAttachment) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Content.IsSet() {
		toSerialize["content"] = o.Content.Get()
	}
	if o.ContentType.IsSet() {
		toSerialize["contentType"] = o.ContentType.Get()
	}
	if o.ContentUrl.IsSet() {
		toSerialize["contentUrl"] = o.ContentUrl.Get()
	}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.ThumbnailUrl.IsSet() {
		toSerialize["thumbnailUrl"] = o.ThumbnailUrl.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphChatMessageAttachment struct {
	value *MicrosoftGraphChatMessageAttachment
	isSet bool
}

func (v NullableMicrosoftGraphChatMessageAttachment) Get() *MicrosoftGraphChatMessageAttachment {
	return v.value
}

func (v *NullableMicrosoftGraphChatMessageAttachment) Set(val *MicrosoftGraphChatMessageAttachment) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphChatMessageAttachment) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphChatMessageAttachment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphChatMessageAttachment(val *MicrosoftGraphChatMessageAttachment) *NullableMicrosoftGraphChatMessageAttachment {
	return &NullableMicrosoftGraphChatMessageAttachment{value: val, isSet: true}
}

func (v NullableMicrosoftGraphChatMessageAttachment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphChatMessageAttachment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


