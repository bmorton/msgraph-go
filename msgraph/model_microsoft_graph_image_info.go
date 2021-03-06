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

// MicrosoftGraphImageInfo struct for MicrosoftGraphImageInfo
type MicrosoftGraphImageInfo struct {
	// Optional; parameter used to indicate the server is able to render image dynamically in response to parameterization. For example – a high contrast image
	AddImageQuery NullableBool `json:"addImageQuery,omitempty"`
	// Optional; alt-text accessible content for the image
	AlternateText NullableString `json:"alternateText,omitempty"`
	AlternativeText NullableString `json:"alternativeText,omitempty"`
	// Optional; URI that points to an icon which represents the application used to generate the activity
	IconUrl NullableString `json:"iconUrl,omitempty"`
}

// NewMicrosoftGraphImageInfo instantiates a new MicrosoftGraphImageInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphImageInfo() *MicrosoftGraphImageInfo {
	this := MicrosoftGraphImageInfo{}
	return &this
}

// NewMicrosoftGraphImageInfoWithDefaults instantiates a new MicrosoftGraphImageInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphImageInfoWithDefaults() *MicrosoftGraphImageInfo {
	this := MicrosoftGraphImageInfo{}
	return &this
}

// GetAddImageQuery returns the AddImageQuery field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphImageInfo) GetAddImageQuery() bool {
	if o == nil || o.AddImageQuery.Get() == nil {
		var ret bool
		return ret
	}
	return *o.AddImageQuery.Get()
}

// GetAddImageQueryOk returns a tuple with the AddImageQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphImageInfo) GetAddImageQueryOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AddImageQuery.Get(), o.AddImageQuery.IsSet()
}

// HasAddImageQuery returns a boolean if a field has been set.
func (o *MicrosoftGraphImageInfo) HasAddImageQuery() bool {
	if o != nil && o.AddImageQuery.IsSet() {
		return true
	}

	return false
}

// SetAddImageQuery gets a reference to the given NullableBool and assigns it to the AddImageQuery field.
func (o *MicrosoftGraphImageInfo) SetAddImageQuery(v bool) {
	o.AddImageQuery.Set(&v)
}
// SetAddImageQueryNil sets the value for AddImageQuery to be an explicit nil
func (o *MicrosoftGraphImageInfo) SetAddImageQueryNil() {
	o.AddImageQuery.Set(nil)
}

// UnsetAddImageQuery ensures that no value is present for AddImageQuery, not even an explicit nil
func (o *MicrosoftGraphImageInfo) UnsetAddImageQuery() {
	o.AddImageQuery.Unset()
}

// GetAlternateText returns the AlternateText field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphImageInfo) GetAlternateText() string {
	if o == nil || o.AlternateText.Get() == nil {
		var ret string
		return ret
	}
	return *o.AlternateText.Get()
}

// GetAlternateTextOk returns a tuple with the AlternateText field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphImageInfo) GetAlternateTextOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AlternateText.Get(), o.AlternateText.IsSet()
}

// HasAlternateText returns a boolean if a field has been set.
func (o *MicrosoftGraphImageInfo) HasAlternateText() bool {
	if o != nil && o.AlternateText.IsSet() {
		return true
	}

	return false
}

// SetAlternateText gets a reference to the given NullableString and assigns it to the AlternateText field.
func (o *MicrosoftGraphImageInfo) SetAlternateText(v string) {
	o.AlternateText.Set(&v)
}
// SetAlternateTextNil sets the value for AlternateText to be an explicit nil
func (o *MicrosoftGraphImageInfo) SetAlternateTextNil() {
	o.AlternateText.Set(nil)
}

// UnsetAlternateText ensures that no value is present for AlternateText, not even an explicit nil
func (o *MicrosoftGraphImageInfo) UnsetAlternateText() {
	o.AlternateText.Unset()
}

// GetAlternativeText returns the AlternativeText field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphImageInfo) GetAlternativeText() string {
	if o == nil || o.AlternativeText.Get() == nil {
		var ret string
		return ret
	}
	return *o.AlternativeText.Get()
}

// GetAlternativeTextOk returns a tuple with the AlternativeText field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphImageInfo) GetAlternativeTextOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AlternativeText.Get(), o.AlternativeText.IsSet()
}

// HasAlternativeText returns a boolean if a field has been set.
func (o *MicrosoftGraphImageInfo) HasAlternativeText() bool {
	if o != nil && o.AlternativeText.IsSet() {
		return true
	}

	return false
}

// SetAlternativeText gets a reference to the given NullableString and assigns it to the AlternativeText field.
func (o *MicrosoftGraphImageInfo) SetAlternativeText(v string) {
	o.AlternativeText.Set(&v)
}
// SetAlternativeTextNil sets the value for AlternativeText to be an explicit nil
func (o *MicrosoftGraphImageInfo) SetAlternativeTextNil() {
	o.AlternativeText.Set(nil)
}

// UnsetAlternativeText ensures that no value is present for AlternativeText, not even an explicit nil
func (o *MicrosoftGraphImageInfo) UnsetAlternativeText() {
	o.AlternativeText.Unset()
}

// GetIconUrl returns the IconUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphImageInfo) GetIconUrl() string {
	if o == nil || o.IconUrl.Get() == nil {
		var ret string
		return ret
	}
	return *o.IconUrl.Get()
}

// GetIconUrlOk returns a tuple with the IconUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphImageInfo) GetIconUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IconUrl.Get(), o.IconUrl.IsSet()
}

// HasIconUrl returns a boolean if a field has been set.
func (o *MicrosoftGraphImageInfo) HasIconUrl() bool {
	if o != nil && o.IconUrl.IsSet() {
		return true
	}

	return false
}

// SetIconUrl gets a reference to the given NullableString and assigns it to the IconUrl field.
func (o *MicrosoftGraphImageInfo) SetIconUrl(v string) {
	o.IconUrl.Set(&v)
}
// SetIconUrlNil sets the value for IconUrl to be an explicit nil
func (o *MicrosoftGraphImageInfo) SetIconUrlNil() {
	o.IconUrl.Set(nil)
}

// UnsetIconUrl ensures that no value is present for IconUrl, not even an explicit nil
func (o *MicrosoftGraphImageInfo) UnsetIconUrl() {
	o.IconUrl.Unset()
}

func (o MicrosoftGraphImageInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AddImageQuery.IsSet() {
		toSerialize["addImageQuery"] = o.AddImageQuery.Get()
	}
	if o.AlternateText.IsSet() {
		toSerialize["alternateText"] = o.AlternateText.Get()
	}
	if o.AlternativeText.IsSet() {
		toSerialize["alternativeText"] = o.AlternativeText.Get()
	}
	if o.IconUrl.IsSet() {
		toSerialize["iconUrl"] = o.IconUrl.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphImageInfo struct {
	value *MicrosoftGraphImageInfo
	isSet bool
}

func (v NullableMicrosoftGraphImageInfo) Get() *MicrosoftGraphImageInfo {
	return v.value
}

func (v *NullableMicrosoftGraphImageInfo) Set(val *MicrosoftGraphImageInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphImageInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphImageInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphImageInfo(val *MicrosoftGraphImageInfo) *NullableMicrosoftGraphImageInfo {
	return &NullableMicrosoftGraphImageInfo{value: val, isSet: true}
}

func (v NullableMicrosoftGraphImageInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphImageInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


