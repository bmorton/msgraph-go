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

// MicrosoftGraphOnenotePagePreview struct for MicrosoftGraphOnenotePagePreview
type MicrosoftGraphOnenotePagePreview struct {
	Links AnyOfmicrosoftGraphOnenotePagePreviewLinks `json:"links,omitempty"`
	PreviewText NullableString `json:"previewText,omitempty"`
}

// NewMicrosoftGraphOnenotePagePreview instantiates a new MicrosoftGraphOnenotePagePreview object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphOnenotePagePreview() *MicrosoftGraphOnenotePagePreview {
	this := MicrosoftGraphOnenotePagePreview{}
	return &this
}

// NewMicrosoftGraphOnenotePagePreviewWithDefaults instantiates a new MicrosoftGraphOnenotePagePreview object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphOnenotePagePreviewWithDefaults() *MicrosoftGraphOnenotePagePreview {
	this := MicrosoftGraphOnenotePagePreview{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphOnenotePagePreview) GetLinks() AnyOfmicrosoftGraphOnenotePagePreviewLinks {
	if o == nil  {
		var ret AnyOfmicrosoftGraphOnenotePagePreviewLinks
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphOnenotePagePreview) GetLinksOk() (*AnyOfmicrosoftGraphOnenotePagePreviewLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return &o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *MicrosoftGraphOnenotePagePreview) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given AnyOfmicrosoftGraphOnenotePagePreviewLinks and assigns it to the Links field.
func (o *MicrosoftGraphOnenotePagePreview) SetLinks(v AnyOfmicrosoftGraphOnenotePagePreviewLinks) {
	o.Links = v
}

// GetPreviewText returns the PreviewText field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphOnenotePagePreview) GetPreviewText() string {
	if o == nil || o.PreviewText.Get() == nil {
		var ret string
		return ret
	}
	return *o.PreviewText.Get()
}

// GetPreviewTextOk returns a tuple with the PreviewText field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphOnenotePagePreview) GetPreviewTextOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PreviewText.Get(), o.PreviewText.IsSet()
}

// HasPreviewText returns a boolean if a field has been set.
func (o *MicrosoftGraphOnenotePagePreview) HasPreviewText() bool {
	if o != nil && o.PreviewText.IsSet() {
		return true
	}

	return false
}

// SetPreviewText gets a reference to the given NullableString and assigns it to the PreviewText field.
func (o *MicrosoftGraphOnenotePagePreview) SetPreviewText(v string) {
	o.PreviewText.Set(&v)
}
// SetPreviewTextNil sets the value for PreviewText to be an explicit nil
func (o *MicrosoftGraphOnenotePagePreview) SetPreviewTextNil() {
	o.PreviewText.Set(nil)
}

// UnsetPreviewText ensures that no value is present for PreviewText, not even an explicit nil
func (o *MicrosoftGraphOnenotePagePreview) UnsetPreviewText() {
	o.PreviewText.Unset()
}

func (o MicrosoftGraphOnenotePagePreview) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.PreviewText.IsSet() {
		toSerialize["previewText"] = o.PreviewText.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphOnenotePagePreview struct {
	value *MicrosoftGraphOnenotePagePreview
	isSet bool
}

func (v NullableMicrosoftGraphOnenotePagePreview) Get() *MicrosoftGraphOnenotePagePreview {
	return v.value
}

func (v *NullableMicrosoftGraphOnenotePagePreview) Set(val *MicrosoftGraphOnenotePagePreview) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphOnenotePagePreview) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphOnenotePagePreview) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphOnenotePagePreview(val *MicrosoftGraphOnenotePagePreview) *NullableMicrosoftGraphOnenotePagePreview {
	return &NullableMicrosoftGraphOnenotePagePreview{value: val, isSet: true}
}

func (v NullableMicrosoftGraphOnenotePagePreview) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphOnenotePagePreview) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


