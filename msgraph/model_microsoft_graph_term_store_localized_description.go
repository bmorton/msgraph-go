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

// MicrosoftGraphTermStoreLocalizedDescription struct for MicrosoftGraphTermStoreLocalizedDescription
type MicrosoftGraphTermStoreLocalizedDescription struct {
	// The description in the localized language.
	Description NullableString `json:"description,omitempty"`
	// The language tag for the label.
	LanguageTag NullableString `json:"languageTag,omitempty"`
}

// NewMicrosoftGraphTermStoreLocalizedDescription instantiates a new MicrosoftGraphTermStoreLocalizedDescription object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphTermStoreLocalizedDescription() *MicrosoftGraphTermStoreLocalizedDescription {
	this := MicrosoftGraphTermStoreLocalizedDescription{}
	return &this
}

// NewMicrosoftGraphTermStoreLocalizedDescriptionWithDefaults instantiates a new MicrosoftGraphTermStoreLocalizedDescription object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphTermStoreLocalizedDescriptionWithDefaults() *MicrosoftGraphTermStoreLocalizedDescription {
	this := MicrosoftGraphTermStoreLocalizedDescription{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphTermStoreLocalizedDescription) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphTermStoreLocalizedDescription) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *MicrosoftGraphTermStoreLocalizedDescription) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *MicrosoftGraphTermStoreLocalizedDescription) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *MicrosoftGraphTermStoreLocalizedDescription) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *MicrosoftGraphTermStoreLocalizedDescription) UnsetDescription() {
	o.Description.Unset()
}

// GetLanguageTag returns the LanguageTag field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphTermStoreLocalizedDescription) GetLanguageTag() string {
	if o == nil || o.LanguageTag.Get() == nil {
		var ret string
		return ret
	}
	return *o.LanguageTag.Get()
}

// GetLanguageTagOk returns a tuple with the LanguageTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphTermStoreLocalizedDescription) GetLanguageTagOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LanguageTag.Get(), o.LanguageTag.IsSet()
}

// HasLanguageTag returns a boolean if a field has been set.
func (o *MicrosoftGraphTermStoreLocalizedDescription) HasLanguageTag() bool {
	if o != nil && o.LanguageTag.IsSet() {
		return true
	}

	return false
}

// SetLanguageTag gets a reference to the given NullableString and assigns it to the LanguageTag field.
func (o *MicrosoftGraphTermStoreLocalizedDescription) SetLanguageTag(v string) {
	o.LanguageTag.Set(&v)
}
// SetLanguageTagNil sets the value for LanguageTag to be an explicit nil
func (o *MicrosoftGraphTermStoreLocalizedDescription) SetLanguageTagNil() {
	o.LanguageTag.Set(nil)
}

// UnsetLanguageTag ensures that no value is present for LanguageTag, not even an explicit nil
func (o *MicrosoftGraphTermStoreLocalizedDescription) UnsetLanguageTag() {
	o.LanguageTag.Unset()
}

func (o MicrosoftGraphTermStoreLocalizedDescription) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.LanguageTag.IsSet() {
		toSerialize["languageTag"] = o.LanguageTag.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphTermStoreLocalizedDescription struct {
	value *MicrosoftGraphTermStoreLocalizedDescription
	isSet bool
}

func (v NullableMicrosoftGraphTermStoreLocalizedDescription) Get() *MicrosoftGraphTermStoreLocalizedDescription {
	return v.value
}

func (v *NullableMicrosoftGraphTermStoreLocalizedDescription) Set(val *MicrosoftGraphTermStoreLocalizedDescription) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphTermStoreLocalizedDescription) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphTermStoreLocalizedDescription) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphTermStoreLocalizedDescription(val *MicrosoftGraphTermStoreLocalizedDescription) *NullableMicrosoftGraphTermStoreLocalizedDescription {
	return &NullableMicrosoftGraphTermStoreLocalizedDescription{value: val, isSet: true}
}

func (v NullableMicrosoftGraphTermStoreLocalizedDescription) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphTermStoreLocalizedDescription) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


