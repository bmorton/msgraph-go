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

// MicrosoftGraphThumbnail struct for MicrosoftGraphThumbnail
type MicrosoftGraphThumbnail struct {
	// The content stream for the thumbnail.
	Content NullableString `json:"content,omitempty"`
	// The height of the thumbnail, in pixels.
	Height NullableInt32 `json:"height,omitempty"`
	// The unique identifier of the item that provided the thumbnail. This is only available when a folder thumbnail is requested.
	SourceItemId NullableString `json:"sourceItemId,omitempty"`
	// The URL used to fetch the thumbnail content.
	Url NullableString `json:"url,omitempty"`
	// The width of the thumbnail, in pixels.
	Width NullableInt32 `json:"width,omitempty"`
}

// NewMicrosoftGraphThumbnail instantiates a new MicrosoftGraphThumbnail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphThumbnail() *MicrosoftGraphThumbnail {
	this := MicrosoftGraphThumbnail{}
	return &this
}

// NewMicrosoftGraphThumbnailWithDefaults instantiates a new MicrosoftGraphThumbnail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphThumbnailWithDefaults() *MicrosoftGraphThumbnail {
	this := MicrosoftGraphThumbnail{}
	return &this
}

// GetContent returns the Content field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphThumbnail) GetContent() string {
	if o == nil || o.Content.Get() == nil {
		var ret string
		return ret
	}
	return *o.Content.Get()
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphThumbnail) GetContentOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Content.Get(), o.Content.IsSet()
}

// HasContent returns a boolean if a field has been set.
func (o *MicrosoftGraphThumbnail) HasContent() bool {
	if o != nil && o.Content.IsSet() {
		return true
	}

	return false
}

// SetContent gets a reference to the given NullableString and assigns it to the Content field.
func (o *MicrosoftGraphThumbnail) SetContent(v string) {
	o.Content.Set(&v)
}
// SetContentNil sets the value for Content to be an explicit nil
func (o *MicrosoftGraphThumbnail) SetContentNil() {
	o.Content.Set(nil)
}

// UnsetContent ensures that no value is present for Content, not even an explicit nil
func (o *MicrosoftGraphThumbnail) UnsetContent() {
	o.Content.Unset()
}

// GetHeight returns the Height field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphThumbnail) GetHeight() int32 {
	if o == nil || o.Height.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Height.Get()
}

// GetHeightOk returns a tuple with the Height field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphThumbnail) GetHeightOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Height.Get(), o.Height.IsSet()
}

// HasHeight returns a boolean if a field has been set.
func (o *MicrosoftGraphThumbnail) HasHeight() bool {
	if o != nil && o.Height.IsSet() {
		return true
	}

	return false
}

// SetHeight gets a reference to the given NullableInt32 and assigns it to the Height field.
func (o *MicrosoftGraphThumbnail) SetHeight(v int32) {
	o.Height.Set(&v)
}
// SetHeightNil sets the value for Height to be an explicit nil
func (o *MicrosoftGraphThumbnail) SetHeightNil() {
	o.Height.Set(nil)
}

// UnsetHeight ensures that no value is present for Height, not even an explicit nil
func (o *MicrosoftGraphThumbnail) UnsetHeight() {
	o.Height.Unset()
}

// GetSourceItemId returns the SourceItemId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphThumbnail) GetSourceItemId() string {
	if o == nil || o.SourceItemId.Get() == nil {
		var ret string
		return ret
	}
	return *o.SourceItemId.Get()
}

// GetSourceItemIdOk returns a tuple with the SourceItemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphThumbnail) GetSourceItemIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SourceItemId.Get(), o.SourceItemId.IsSet()
}

// HasSourceItemId returns a boolean if a field has been set.
func (o *MicrosoftGraphThumbnail) HasSourceItemId() bool {
	if o != nil && o.SourceItemId.IsSet() {
		return true
	}

	return false
}

// SetSourceItemId gets a reference to the given NullableString and assigns it to the SourceItemId field.
func (o *MicrosoftGraphThumbnail) SetSourceItemId(v string) {
	o.SourceItemId.Set(&v)
}
// SetSourceItemIdNil sets the value for SourceItemId to be an explicit nil
func (o *MicrosoftGraphThumbnail) SetSourceItemIdNil() {
	o.SourceItemId.Set(nil)
}

// UnsetSourceItemId ensures that no value is present for SourceItemId, not even an explicit nil
func (o *MicrosoftGraphThumbnail) UnsetSourceItemId() {
	o.SourceItemId.Unset()
}

// GetUrl returns the Url field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphThumbnail) GetUrl() string {
	if o == nil || o.Url.Get() == nil {
		var ret string
		return ret
	}
	return *o.Url.Get()
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphThumbnail) GetUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Url.Get(), o.Url.IsSet()
}

// HasUrl returns a boolean if a field has been set.
func (o *MicrosoftGraphThumbnail) HasUrl() bool {
	if o != nil && o.Url.IsSet() {
		return true
	}

	return false
}

// SetUrl gets a reference to the given NullableString and assigns it to the Url field.
func (o *MicrosoftGraphThumbnail) SetUrl(v string) {
	o.Url.Set(&v)
}
// SetUrlNil sets the value for Url to be an explicit nil
func (o *MicrosoftGraphThumbnail) SetUrlNil() {
	o.Url.Set(nil)
}

// UnsetUrl ensures that no value is present for Url, not even an explicit nil
func (o *MicrosoftGraphThumbnail) UnsetUrl() {
	o.Url.Unset()
}

// GetWidth returns the Width field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphThumbnail) GetWidth() int32 {
	if o == nil || o.Width.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Width.Get()
}

// GetWidthOk returns a tuple with the Width field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphThumbnail) GetWidthOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Width.Get(), o.Width.IsSet()
}

// HasWidth returns a boolean if a field has been set.
func (o *MicrosoftGraphThumbnail) HasWidth() bool {
	if o != nil && o.Width.IsSet() {
		return true
	}

	return false
}

// SetWidth gets a reference to the given NullableInt32 and assigns it to the Width field.
func (o *MicrosoftGraphThumbnail) SetWidth(v int32) {
	o.Width.Set(&v)
}
// SetWidthNil sets the value for Width to be an explicit nil
func (o *MicrosoftGraphThumbnail) SetWidthNil() {
	o.Width.Set(nil)
}

// UnsetWidth ensures that no value is present for Width, not even an explicit nil
func (o *MicrosoftGraphThumbnail) UnsetWidth() {
	o.Width.Unset()
}

func (o MicrosoftGraphThumbnail) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Content.IsSet() {
		toSerialize["content"] = o.Content.Get()
	}
	if o.Height.IsSet() {
		toSerialize["height"] = o.Height.Get()
	}
	if o.SourceItemId.IsSet() {
		toSerialize["sourceItemId"] = o.SourceItemId.Get()
	}
	if o.Url.IsSet() {
		toSerialize["url"] = o.Url.Get()
	}
	if o.Width.IsSet() {
		toSerialize["width"] = o.Width.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphThumbnail struct {
	value *MicrosoftGraphThumbnail
	isSet bool
}

func (v NullableMicrosoftGraphThumbnail) Get() *MicrosoftGraphThumbnail {
	return v.value
}

func (v *NullableMicrosoftGraphThumbnail) Set(val *MicrosoftGraphThumbnail) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphThumbnail) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphThumbnail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphThumbnail(val *MicrosoftGraphThumbnail) *NullableMicrosoftGraphThumbnail {
	return &NullableMicrosoftGraphThumbnail{value: val, isSet: true}
}

func (v NullableMicrosoftGraphThumbnail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphThumbnail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

