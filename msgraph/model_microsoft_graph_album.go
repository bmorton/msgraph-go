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

// MicrosoftGraphAlbum struct for MicrosoftGraphAlbum
type MicrosoftGraphAlbum struct {
	// Unique identifier of the [driveItem][] that is the cover of the album.
	CoverImageItemId NullableString `json:"coverImageItemId,omitempty"`
}

// NewMicrosoftGraphAlbum instantiates a new MicrosoftGraphAlbum object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphAlbum() *MicrosoftGraphAlbum {
	this := MicrosoftGraphAlbum{}
	return &this
}

// NewMicrosoftGraphAlbumWithDefaults instantiates a new MicrosoftGraphAlbum object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphAlbumWithDefaults() *MicrosoftGraphAlbum {
	this := MicrosoftGraphAlbum{}
	return &this
}

// GetCoverImageItemId returns the CoverImageItemId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphAlbum) GetCoverImageItemId() string {
	if o == nil || o.CoverImageItemId.Get() == nil {
		var ret string
		return ret
	}
	return *o.CoverImageItemId.Get()
}

// GetCoverImageItemIdOk returns a tuple with the CoverImageItemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphAlbum) GetCoverImageItemIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CoverImageItemId.Get(), o.CoverImageItemId.IsSet()
}

// HasCoverImageItemId returns a boolean if a field has been set.
func (o *MicrosoftGraphAlbum) HasCoverImageItemId() bool {
	if o != nil && o.CoverImageItemId.IsSet() {
		return true
	}

	return false
}

// SetCoverImageItemId gets a reference to the given NullableString and assigns it to the CoverImageItemId field.
func (o *MicrosoftGraphAlbum) SetCoverImageItemId(v string) {
	o.CoverImageItemId.Set(&v)
}
// SetCoverImageItemIdNil sets the value for CoverImageItemId to be an explicit nil
func (o *MicrosoftGraphAlbum) SetCoverImageItemIdNil() {
	o.CoverImageItemId.Set(nil)
}

// UnsetCoverImageItemId ensures that no value is present for CoverImageItemId, not even an explicit nil
func (o *MicrosoftGraphAlbum) UnsetCoverImageItemId() {
	o.CoverImageItemId.Unset()
}

func (o MicrosoftGraphAlbum) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CoverImageItemId.IsSet() {
		toSerialize["coverImageItemId"] = o.CoverImageItemId.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphAlbum struct {
	value *MicrosoftGraphAlbum
	isSet bool
}

func (v NullableMicrosoftGraphAlbum) Get() *MicrosoftGraphAlbum {
	return v.value
}

func (v *NullableMicrosoftGraphAlbum) Set(val *MicrosoftGraphAlbum) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphAlbum) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphAlbum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphAlbum(val *MicrosoftGraphAlbum) *NullableMicrosoftGraphAlbum {
	return &NullableMicrosoftGraphAlbum{value: val, isSet: true}
}

func (v NullableMicrosoftGraphAlbum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphAlbum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


