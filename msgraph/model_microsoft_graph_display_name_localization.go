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

// MicrosoftGraphDisplayNameLocalization struct for MicrosoftGraphDisplayNameLocalization
type MicrosoftGraphDisplayNameLocalization struct {
	// If present, the value of this field contains the displayName string that has been set for the language present in the languageTag field.
	DisplayName NullableString `json:"displayName,omitempty"`
	// Provides the language culture-code and friendly name of the language that the displayName field has been provided in.
	LanguageTag NullableString `json:"languageTag,omitempty"`
}

// NewMicrosoftGraphDisplayNameLocalization instantiates a new MicrosoftGraphDisplayNameLocalization object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphDisplayNameLocalization() *MicrosoftGraphDisplayNameLocalization {
	this := MicrosoftGraphDisplayNameLocalization{}
	return &this
}

// NewMicrosoftGraphDisplayNameLocalizationWithDefaults instantiates a new MicrosoftGraphDisplayNameLocalization object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphDisplayNameLocalizationWithDefaults() *MicrosoftGraphDisplayNameLocalization {
	this := MicrosoftGraphDisplayNameLocalization{}
	return &this
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphDisplayNameLocalization) GetDisplayName() string {
	if o == nil || o.DisplayName.Get() == nil {
		var ret string
		return ret
	}
	return *o.DisplayName.Get()
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphDisplayNameLocalization) GetDisplayNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DisplayName.Get(), o.DisplayName.IsSet()
}

// HasDisplayName returns a boolean if a field has been set.
func (o *MicrosoftGraphDisplayNameLocalization) HasDisplayName() bool {
	if o != nil && o.DisplayName.IsSet() {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given NullableString and assigns it to the DisplayName field.
func (o *MicrosoftGraphDisplayNameLocalization) SetDisplayName(v string) {
	o.DisplayName.Set(&v)
}
// SetDisplayNameNil sets the value for DisplayName to be an explicit nil
func (o *MicrosoftGraphDisplayNameLocalization) SetDisplayNameNil() {
	o.DisplayName.Set(nil)
}

// UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
func (o *MicrosoftGraphDisplayNameLocalization) UnsetDisplayName() {
	o.DisplayName.Unset()
}

// GetLanguageTag returns the LanguageTag field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphDisplayNameLocalization) GetLanguageTag() string {
	if o == nil || o.LanguageTag.Get() == nil {
		var ret string
		return ret
	}
	return *o.LanguageTag.Get()
}

// GetLanguageTagOk returns a tuple with the LanguageTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphDisplayNameLocalization) GetLanguageTagOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LanguageTag.Get(), o.LanguageTag.IsSet()
}

// HasLanguageTag returns a boolean if a field has been set.
func (o *MicrosoftGraphDisplayNameLocalization) HasLanguageTag() bool {
	if o != nil && o.LanguageTag.IsSet() {
		return true
	}

	return false
}

// SetLanguageTag gets a reference to the given NullableString and assigns it to the LanguageTag field.
func (o *MicrosoftGraphDisplayNameLocalization) SetLanguageTag(v string) {
	o.LanguageTag.Set(&v)
}
// SetLanguageTagNil sets the value for LanguageTag to be an explicit nil
func (o *MicrosoftGraphDisplayNameLocalization) SetLanguageTagNil() {
	o.LanguageTag.Set(nil)
}

// UnsetLanguageTag ensures that no value is present for LanguageTag, not even an explicit nil
func (o *MicrosoftGraphDisplayNameLocalization) UnsetLanguageTag() {
	o.LanguageTag.Unset()
}

func (o MicrosoftGraphDisplayNameLocalization) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DisplayName.IsSet() {
		toSerialize["displayName"] = o.DisplayName.Get()
	}
	if o.LanguageTag.IsSet() {
		toSerialize["languageTag"] = o.LanguageTag.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphDisplayNameLocalization struct {
	value *MicrosoftGraphDisplayNameLocalization
	isSet bool
}

func (v NullableMicrosoftGraphDisplayNameLocalization) Get() *MicrosoftGraphDisplayNameLocalization {
	return v.value
}

func (v *NullableMicrosoftGraphDisplayNameLocalization) Set(val *MicrosoftGraphDisplayNameLocalization) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphDisplayNameLocalization) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphDisplayNameLocalization) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphDisplayNameLocalization(val *MicrosoftGraphDisplayNameLocalization) *NullableMicrosoftGraphDisplayNameLocalization {
	return &NullableMicrosoftGraphDisplayNameLocalization{value: val, isSet: true}
}

func (v NullableMicrosoftGraphDisplayNameLocalization) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphDisplayNameLocalization) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


