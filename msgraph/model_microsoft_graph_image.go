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

// MicrosoftGraphImage struct for MicrosoftGraphImage
type MicrosoftGraphImage struct {
	// Optional. Height of the image, in pixels. Read-only.
	Height NullableInt32 `json:"height,omitempty"`
	// Optional. Width of the image, in pixels. Read-only.
	Width NullableInt32 `json:"width,omitempty"`
}

// NewMicrosoftGraphImage instantiates a new MicrosoftGraphImage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphImage() *MicrosoftGraphImage {
	this := MicrosoftGraphImage{}
	return &this
}

// NewMicrosoftGraphImageWithDefaults instantiates a new MicrosoftGraphImage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphImageWithDefaults() *MicrosoftGraphImage {
	this := MicrosoftGraphImage{}
	return &this
}

// GetHeight returns the Height field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphImage) GetHeight() int32 {
	if o == nil || o.Height.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Height.Get()
}

// GetHeightOk returns a tuple with the Height field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphImage) GetHeightOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Height.Get(), o.Height.IsSet()
}

// HasHeight returns a boolean if a field has been set.
func (o *MicrosoftGraphImage) HasHeight() bool {
	if o != nil && o.Height.IsSet() {
		return true
	}

	return false
}

// SetHeight gets a reference to the given NullableInt32 and assigns it to the Height field.
func (o *MicrosoftGraphImage) SetHeight(v int32) {
	o.Height.Set(&v)
}
// SetHeightNil sets the value for Height to be an explicit nil
func (o *MicrosoftGraphImage) SetHeightNil() {
	o.Height.Set(nil)
}

// UnsetHeight ensures that no value is present for Height, not even an explicit nil
func (o *MicrosoftGraphImage) UnsetHeight() {
	o.Height.Unset()
}

// GetWidth returns the Width field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphImage) GetWidth() int32 {
	if o == nil || o.Width.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Width.Get()
}

// GetWidthOk returns a tuple with the Width field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphImage) GetWidthOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Width.Get(), o.Width.IsSet()
}

// HasWidth returns a boolean if a field has been set.
func (o *MicrosoftGraphImage) HasWidth() bool {
	if o != nil && o.Width.IsSet() {
		return true
	}

	return false
}

// SetWidth gets a reference to the given NullableInt32 and assigns it to the Width field.
func (o *MicrosoftGraphImage) SetWidth(v int32) {
	o.Width.Set(&v)
}
// SetWidthNil sets the value for Width to be an explicit nil
func (o *MicrosoftGraphImage) SetWidthNil() {
	o.Width.Set(nil)
}

// UnsetWidth ensures that no value is present for Width, not even an explicit nil
func (o *MicrosoftGraphImage) UnsetWidth() {
	o.Width.Unset()
}

func (o MicrosoftGraphImage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Height.IsSet() {
		toSerialize["height"] = o.Height.Get()
	}
	if o.Width.IsSet() {
		toSerialize["width"] = o.Width.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphImage struct {
	value *MicrosoftGraphImage
	isSet bool
}

func (v NullableMicrosoftGraphImage) Get() *MicrosoftGraphImage {
	return v.value
}

func (v *NullableMicrosoftGraphImage) Set(val *MicrosoftGraphImage) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphImage) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphImage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphImage(val *MicrosoftGraphImage) *NullableMicrosoftGraphImage {
	return &NullableMicrosoftGraphImage{value: val, isSet: true}
}

func (v NullableMicrosoftGraphImage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphImage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


