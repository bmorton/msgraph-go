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

// MicrosoftGraphTermStoreLocalizedLabel struct for MicrosoftGraphTermStoreLocalizedLabel
type MicrosoftGraphTermStoreLocalizedLabel struct {
	// Indicates whether the label is the default label.
	IsDefault NullableBool `json:"isDefault,omitempty"`
	// The language tag for the label.
	LanguageTag NullableString `json:"languageTag,omitempty"`
	// The name of the label.
	Name NullableString `json:"name,omitempty"`
}

// NewMicrosoftGraphTermStoreLocalizedLabel instantiates a new MicrosoftGraphTermStoreLocalizedLabel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphTermStoreLocalizedLabel() *MicrosoftGraphTermStoreLocalizedLabel {
	this := MicrosoftGraphTermStoreLocalizedLabel{}
	return &this
}

// NewMicrosoftGraphTermStoreLocalizedLabelWithDefaults instantiates a new MicrosoftGraphTermStoreLocalizedLabel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphTermStoreLocalizedLabelWithDefaults() *MicrosoftGraphTermStoreLocalizedLabel {
	this := MicrosoftGraphTermStoreLocalizedLabel{}
	return &this
}

// GetIsDefault returns the IsDefault field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphTermStoreLocalizedLabel) GetIsDefault() bool {
	if o == nil || o.IsDefault.Get() == nil {
		var ret bool
		return ret
	}
	return *o.IsDefault.Get()
}

// GetIsDefaultOk returns a tuple with the IsDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphTermStoreLocalizedLabel) GetIsDefaultOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IsDefault.Get(), o.IsDefault.IsSet()
}

// HasIsDefault returns a boolean if a field has been set.
func (o *MicrosoftGraphTermStoreLocalizedLabel) HasIsDefault() bool {
	if o != nil && o.IsDefault.IsSet() {
		return true
	}

	return false
}

// SetIsDefault gets a reference to the given NullableBool and assigns it to the IsDefault field.
func (o *MicrosoftGraphTermStoreLocalizedLabel) SetIsDefault(v bool) {
	o.IsDefault.Set(&v)
}
// SetIsDefaultNil sets the value for IsDefault to be an explicit nil
func (o *MicrosoftGraphTermStoreLocalizedLabel) SetIsDefaultNil() {
	o.IsDefault.Set(nil)
}

// UnsetIsDefault ensures that no value is present for IsDefault, not even an explicit nil
func (o *MicrosoftGraphTermStoreLocalizedLabel) UnsetIsDefault() {
	o.IsDefault.Unset()
}

// GetLanguageTag returns the LanguageTag field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphTermStoreLocalizedLabel) GetLanguageTag() string {
	if o == nil || o.LanguageTag.Get() == nil {
		var ret string
		return ret
	}
	return *o.LanguageTag.Get()
}

// GetLanguageTagOk returns a tuple with the LanguageTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphTermStoreLocalizedLabel) GetLanguageTagOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LanguageTag.Get(), o.LanguageTag.IsSet()
}

// HasLanguageTag returns a boolean if a field has been set.
func (o *MicrosoftGraphTermStoreLocalizedLabel) HasLanguageTag() bool {
	if o != nil && o.LanguageTag.IsSet() {
		return true
	}

	return false
}

// SetLanguageTag gets a reference to the given NullableString and assigns it to the LanguageTag field.
func (o *MicrosoftGraphTermStoreLocalizedLabel) SetLanguageTag(v string) {
	o.LanguageTag.Set(&v)
}
// SetLanguageTagNil sets the value for LanguageTag to be an explicit nil
func (o *MicrosoftGraphTermStoreLocalizedLabel) SetLanguageTagNil() {
	o.LanguageTag.Set(nil)
}

// UnsetLanguageTag ensures that no value is present for LanguageTag, not even an explicit nil
func (o *MicrosoftGraphTermStoreLocalizedLabel) UnsetLanguageTag() {
	o.LanguageTag.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphTermStoreLocalizedLabel) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphTermStoreLocalizedLabel) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *MicrosoftGraphTermStoreLocalizedLabel) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *MicrosoftGraphTermStoreLocalizedLabel) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *MicrosoftGraphTermStoreLocalizedLabel) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *MicrosoftGraphTermStoreLocalizedLabel) UnsetName() {
	o.Name.Unset()
}

func (o MicrosoftGraphTermStoreLocalizedLabel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.IsDefault.IsSet() {
		toSerialize["isDefault"] = o.IsDefault.Get()
	}
	if o.LanguageTag.IsSet() {
		toSerialize["languageTag"] = o.LanguageTag.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphTermStoreLocalizedLabel struct {
	value *MicrosoftGraphTermStoreLocalizedLabel
	isSet bool
}

func (v NullableMicrosoftGraphTermStoreLocalizedLabel) Get() *MicrosoftGraphTermStoreLocalizedLabel {
	return v.value
}

func (v *NullableMicrosoftGraphTermStoreLocalizedLabel) Set(val *MicrosoftGraphTermStoreLocalizedLabel) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphTermStoreLocalizedLabel) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphTermStoreLocalizedLabel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphTermStoreLocalizedLabel(val *MicrosoftGraphTermStoreLocalizedLabel) *NullableMicrosoftGraphTermStoreLocalizedLabel {
	return &NullableMicrosoftGraphTermStoreLocalizedLabel{value: val, isSet: true}
}

func (v NullableMicrosoftGraphTermStoreLocalizedLabel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphTermStoreLocalizedLabel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


