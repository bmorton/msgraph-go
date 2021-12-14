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

// MicrosoftGraphUserFlowLanguageConfiguration struct for MicrosoftGraphUserFlowLanguageConfiguration
type MicrosoftGraphUserFlowLanguageConfiguration struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
	// The language name to display. This property is read-only.
	DisplayName NullableString `json:"displayName,omitempty"`
	// Indicates whether the language is enabled within the user flow.
	IsEnabled *bool `json:"isEnabled,omitempty"`
	// Collection of pages with the default content to display in a user flow for a specified language. This collection does not allow any kind of modification.
	DefaultPages *[]MicrosoftGraphUserFlowLanguagePage `json:"defaultPages,omitempty"`
	// Collection of pages with the overrides messages to display in a user flow for a specified language. This collection only allows to modify the content of the page, any other modification is not allowed (creation or deletion of pages).
	OverridesPages *[]MicrosoftGraphUserFlowLanguagePage `json:"overridesPages,omitempty"`
}

// NewMicrosoftGraphUserFlowLanguageConfiguration instantiates a new MicrosoftGraphUserFlowLanguageConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphUserFlowLanguageConfiguration() *MicrosoftGraphUserFlowLanguageConfiguration {
	this := MicrosoftGraphUserFlowLanguageConfiguration{}
	return &this
}

// NewMicrosoftGraphUserFlowLanguageConfigurationWithDefaults instantiates a new MicrosoftGraphUserFlowLanguageConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphUserFlowLanguageConfigurationWithDefaults() *MicrosoftGraphUserFlowLanguageConfiguration {
	this := MicrosoftGraphUserFlowLanguageConfiguration{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MicrosoftGraphUserFlowLanguageConfiguration) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphUserFlowLanguageConfiguration) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphUserFlowLanguageConfiguration) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MicrosoftGraphUserFlowLanguageConfiguration) SetId(v string) {
	o.Id = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphUserFlowLanguageConfiguration) GetDisplayName() string {
	if o == nil || o.DisplayName.Get() == nil {
		var ret string
		return ret
	}
	return *o.DisplayName.Get()
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphUserFlowLanguageConfiguration) GetDisplayNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DisplayName.Get(), o.DisplayName.IsSet()
}

// HasDisplayName returns a boolean if a field has been set.
func (o *MicrosoftGraphUserFlowLanguageConfiguration) HasDisplayName() bool {
	if o != nil && o.DisplayName.IsSet() {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given NullableString and assigns it to the DisplayName field.
func (o *MicrosoftGraphUserFlowLanguageConfiguration) SetDisplayName(v string) {
	o.DisplayName.Set(&v)
}
// SetDisplayNameNil sets the value for DisplayName to be an explicit nil
func (o *MicrosoftGraphUserFlowLanguageConfiguration) SetDisplayNameNil() {
	o.DisplayName.Set(nil)
}

// UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
func (o *MicrosoftGraphUserFlowLanguageConfiguration) UnsetDisplayName() {
	o.DisplayName.Unset()
}

// GetIsEnabled returns the IsEnabled field value if set, zero value otherwise.
func (o *MicrosoftGraphUserFlowLanguageConfiguration) GetIsEnabled() bool {
	if o == nil || o.IsEnabled == nil {
		var ret bool
		return ret
	}
	return *o.IsEnabled
}

// GetIsEnabledOk returns a tuple with the IsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphUserFlowLanguageConfiguration) GetIsEnabledOk() (*bool, bool) {
	if o == nil || o.IsEnabled == nil {
		return nil, false
	}
	return o.IsEnabled, true
}

// HasIsEnabled returns a boolean if a field has been set.
func (o *MicrosoftGraphUserFlowLanguageConfiguration) HasIsEnabled() bool {
	if o != nil && o.IsEnabled != nil {
		return true
	}

	return false
}

// SetIsEnabled gets a reference to the given bool and assigns it to the IsEnabled field.
func (o *MicrosoftGraphUserFlowLanguageConfiguration) SetIsEnabled(v bool) {
	o.IsEnabled = &v
}

// GetDefaultPages returns the DefaultPages field value if set, zero value otherwise.
func (o *MicrosoftGraphUserFlowLanguageConfiguration) GetDefaultPages() []MicrosoftGraphUserFlowLanguagePage {
	if o == nil || o.DefaultPages == nil {
		var ret []MicrosoftGraphUserFlowLanguagePage
		return ret
	}
	return *o.DefaultPages
}

// GetDefaultPagesOk returns a tuple with the DefaultPages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphUserFlowLanguageConfiguration) GetDefaultPagesOk() (*[]MicrosoftGraphUserFlowLanguagePage, bool) {
	if o == nil || o.DefaultPages == nil {
		return nil, false
	}
	return o.DefaultPages, true
}

// HasDefaultPages returns a boolean if a field has been set.
func (o *MicrosoftGraphUserFlowLanguageConfiguration) HasDefaultPages() bool {
	if o != nil && o.DefaultPages != nil {
		return true
	}

	return false
}

// SetDefaultPages gets a reference to the given []MicrosoftGraphUserFlowLanguagePage and assigns it to the DefaultPages field.
func (o *MicrosoftGraphUserFlowLanguageConfiguration) SetDefaultPages(v []MicrosoftGraphUserFlowLanguagePage) {
	o.DefaultPages = &v
}

// GetOverridesPages returns the OverridesPages field value if set, zero value otherwise.
func (o *MicrosoftGraphUserFlowLanguageConfiguration) GetOverridesPages() []MicrosoftGraphUserFlowLanguagePage {
	if o == nil || o.OverridesPages == nil {
		var ret []MicrosoftGraphUserFlowLanguagePage
		return ret
	}
	return *o.OverridesPages
}

// GetOverridesPagesOk returns a tuple with the OverridesPages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphUserFlowLanguageConfiguration) GetOverridesPagesOk() (*[]MicrosoftGraphUserFlowLanguagePage, bool) {
	if o == nil || o.OverridesPages == nil {
		return nil, false
	}
	return o.OverridesPages, true
}

// HasOverridesPages returns a boolean if a field has been set.
func (o *MicrosoftGraphUserFlowLanguageConfiguration) HasOverridesPages() bool {
	if o != nil && o.OverridesPages != nil {
		return true
	}

	return false
}

// SetOverridesPages gets a reference to the given []MicrosoftGraphUserFlowLanguagePage and assigns it to the OverridesPages field.
func (o *MicrosoftGraphUserFlowLanguageConfiguration) SetOverridesPages(v []MicrosoftGraphUserFlowLanguagePage) {
	o.OverridesPages = &v
}

func (o MicrosoftGraphUserFlowLanguageConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.DisplayName.IsSet() {
		toSerialize["displayName"] = o.DisplayName.Get()
	}
	if o.IsEnabled != nil {
		toSerialize["isEnabled"] = o.IsEnabled
	}
	if o.DefaultPages != nil {
		toSerialize["defaultPages"] = o.DefaultPages
	}
	if o.OverridesPages != nil {
		toSerialize["overridesPages"] = o.OverridesPages
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphUserFlowLanguageConfiguration struct {
	value *MicrosoftGraphUserFlowLanguageConfiguration
	isSet bool
}

func (v NullableMicrosoftGraphUserFlowLanguageConfiguration) Get() *MicrosoftGraphUserFlowLanguageConfiguration {
	return v.value
}

func (v *NullableMicrosoftGraphUserFlowLanguageConfiguration) Set(val *MicrosoftGraphUserFlowLanguageConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphUserFlowLanguageConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphUserFlowLanguageConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphUserFlowLanguageConfiguration(val *MicrosoftGraphUserFlowLanguageConfiguration) *NullableMicrosoftGraphUserFlowLanguageConfiguration {
	return &NullableMicrosoftGraphUserFlowLanguageConfiguration{value: val, isSet: true}
}

func (v NullableMicrosoftGraphUserFlowLanguageConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphUserFlowLanguageConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

