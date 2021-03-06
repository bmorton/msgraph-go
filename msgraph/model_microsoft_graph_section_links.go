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

// MicrosoftGraphSectionLinks struct for MicrosoftGraphSectionLinks
type MicrosoftGraphSectionLinks struct {
	// Opens the section in the OneNote native client if it's installed.
	OneNoteClientUrl AnyOfmicrosoftGraphExternalLink `json:"oneNoteClientUrl,omitempty"`
	// Opens the section in OneNote on the web.
	OneNoteWebUrl AnyOfmicrosoftGraphExternalLink `json:"oneNoteWebUrl,omitempty"`
}

// NewMicrosoftGraphSectionLinks instantiates a new MicrosoftGraphSectionLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphSectionLinks() *MicrosoftGraphSectionLinks {
	this := MicrosoftGraphSectionLinks{}
	return &this
}

// NewMicrosoftGraphSectionLinksWithDefaults instantiates a new MicrosoftGraphSectionLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphSectionLinksWithDefaults() *MicrosoftGraphSectionLinks {
	this := MicrosoftGraphSectionLinks{}
	return &this
}

// GetOneNoteClientUrl returns the OneNoteClientUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphSectionLinks) GetOneNoteClientUrl() AnyOfmicrosoftGraphExternalLink {
	if o == nil  {
		var ret AnyOfmicrosoftGraphExternalLink
		return ret
	}
	return o.OneNoteClientUrl
}

// GetOneNoteClientUrlOk returns a tuple with the OneNoteClientUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphSectionLinks) GetOneNoteClientUrlOk() (*AnyOfmicrosoftGraphExternalLink, bool) {
	if o == nil || o.OneNoteClientUrl == nil {
		return nil, false
	}
	return &o.OneNoteClientUrl, true
}

// HasOneNoteClientUrl returns a boolean if a field has been set.
func (o *MicrosoftGraphSectionLinks) HasOneNoteClientUrl() bool {
	if o != nil && o.OneNoteClientUrl != nil {
		return true
	}

	return false
}

// SetOneNoteClientUrl gets a reference to the given AnyOfmicrosoftGraphExternalLink and assigns it to the OneNoteClientUrl field.
func (o *MicrosoftGraphSectionLinks) SetOneNoteClientUrl(v AnyOfmicrosoftGraphExternalLink) {
	o.OneNoteClientUrl = v
}

// GetOneNoteWebUrl returns the OneNoteWebUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphSectionLinks) GetOneNoteWebUrl() AnyOfmicrosoftGraphExternalLink {
	if o == nil  {
		var ret AnyOfmicrosoftGraphExternalLink
		return ret
	}
	return o.OneNoteWebUrl
}

// GetOneNoteWebUrlOk returns a tuple with the OneNoteWebUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphSectionLinks) GetOneNoteWebUrlOk() (*AnyOfmicrosoftGraphExternalLink, bool) {
	if o == nil || o.OneNoteWebUrl == nil {
		return nil, false
	}
	return &o.OneNoteWebUrl, true
}

// HasOneNoteWebUrl returns a boolean if a field has been set.
func (o *MicrosoftGraphSectionLinks) HasOneNoteWebUrl() bool {
	if o != nil && o.OneNoteWebUrl != nil {
		return true
	}

	return false
}

// SetOneNoteWebUrl gets a reference to the given AnyOfmicrosoftGraphExternalLink and assigns it to the OneNoteWebUrl field.
func (o *MicrosoftGraphSectionLinks) SetOneNoteWebUrl(v AnyOfmicrosoftGraphExternalLink) {
	o.OneNoteWebUrl = v
}

func (o MicrosoftGraphSectionLinks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.OneNoteClientUrl != nil {
		toSerialize["oneNoteClientUrl"] = o.OneNoteClientUrl
	}
	if o.OneNoteWebUrl != nil {
		toSerialize["oneNoteWebUrl"] = o.OneNoteWebUrl
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphSectionLinks struct {
	value *MicrosoftGraphSectionLinks
	isSet bool
}

func (v NullableMicrosoftGraphSectionLinks) Get() *MicrosoftGraphSectionLinks {
	return v.value
}

func (v *NullableMicrosoftGraphSectionLinks) Set(val *MicrosoftGraphSectionLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphSectionLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphSectionLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphSectionLinks(val *MicrosoftGraphSectionLinks) *NullableMicrosoftGraphSectionLinks {
	return &NullableMicrosoftGraphSectionLinks{value: val, isSet: true}
}

func (v NullableMicrosoftGraphSectionLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphSectionLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


