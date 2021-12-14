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

// MicrosoftGraphHyperlinkOrPictureColumn struct for MicrosoftGraphHyperlinkOrPictureColumn
type MicrosoftGraphHyperlinkOrPictureColumn struct {
	// Specifies whether the display format used for URL columns is an image or a hyperlink.
	IsPicture NullableBool `json:"isPicture,omitempty"`
}

// NewMicrosoftGraphHyperlinkOrPictureColumn instantiates a new MicrosoftGraphHyperlinkOrPictureColumn object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphHyperlinkOrPictureColumn() *MicrosoftGraphHyperlinkOrPictureColumn {
	this := MicrosoftGraphHyperlinkOrPictureColumn{}
	return &this
}

// NewMicrosoftGraphHyperlinkOrPictureColumnWithDefaults instantiates a new MicrosoftGraphHyperlinkOrPictureColumn object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphHyperlinkOrPictureColumnWithDefaults() *MicrosoftGraphHyperlinkOrPictureColumn {
	this := MicrosoftGraphHyperlinkOrPictureColumn{}
	return &this
}

// GetIsPicture returns the IsPicture field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphHyperlinkOrPictureColumn) GetIsPicture() bool {
	if o == nil || o.IsPicture.Get() == nil {
		var ret bool
		return ret
	}
	return *o.IsPicture.Get()
}

// GetIsPictureOk returns a tuple with the IsPicture field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphHyperlinkOrPictureColumn) GetIsPictureOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IsPicture.Get(), o.IsPicture.IsSet()
}

// HasIsPicture returns a boolean if a field has been set.
func (o *MicrosoftGraphHyperlinkOrPictureColumn) HasIsPicture() bool {
	if o != nil && o.IsPicture.IsSet() {
		return true
	}

	return false
}

// SetIsPicture gets a reference to the given NullableBool and assigns it to the IsPicture field.
func (o *MicrosoftGraphHyperlinkOrPictureColumn) SetIsPicture(v bool) {
	o.IsPicture.Set(&v)
}
// SetIsPictureNil sets the value for IsPicture to be an explicit nil
func (o *MicrosoftGraphHyperlinkOrPictureColumn) SetIsPictureNil() {
	o.IsPicture.Set(nil)
}

// UnsetIsPicture ensures that no value is present for IsPicture, not even an explicit nil
func (o *MicrosoftGraphHyperlinkOrPictureColumn) UnsetIsPicture() {
	o.IsPicture.Unset()
}

func (o MicrosoftGraphHyperlinkOrPictureColumn) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.IsPicture.IsSet() {
		toSerialize["isPicture"] = o.IsPicture.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphHyperlinkOrPictureColumn struct {
	value *MicrosoftGraphHyperlinkOrPictureColumn
	isSet bool
}

func (v NullableMicrosoftGraphHyperlinkOrPictureColumn) Get() *MicrosoftGraphHyperlinkOrPictureColumn {
	return v.value
}

func (v *NullableMicrosoftGraphHyperlinkOrPictureColumn) Set(val *MicrosoftGraphHyperlinkOrPictureColumn) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphHyperlinkOrPictureColumn) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphHyperlinkOrPictureColumn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphHyperlinkOrPictureColumn(val *MicrosoftGraphHyperlinkOrPictureColumn) *NullableMicrosoftGraphHyperlinkOrPictureColumn {
	return &NullableMicrosoftGraphHyperlinkOrPictureColumn{value: val, isSet: true}
}

func (v NullableMicrosoftGraphHyperlinkOrPictureColumn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphHyperlinkOrPictureColumn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


