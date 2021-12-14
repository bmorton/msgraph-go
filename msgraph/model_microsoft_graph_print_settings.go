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

// MicrosoftGraphPrintSettings struct for MicrosoftGraphPrintSettings
type MicrosoftGraphPrintSettings struct {
	// Specifies whether document conversion is enabled for the tenant. If document conversion is enabled, Universal Print service will automatically convert documents into a format compatible with the printer (xps to pdf) when needed.
	DocumentConversionEnabled *bool `json:"documentConversionEnabled,omitempty"`
}

// NewMicrosoftGraphPrintSettings instantiates a new MicrosoftGraphPrintSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphPrintSettings() *MicrosoftGraphPrintSettings {
	this := MicrosoftGraphPrintSettings{}
	return &this
}

// NewMicrosoftGraphPrintSettingsWithDefaults instantiates a new MicrosoftGraphPrintSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphPrintSettingsWithDefaults() *MicrosoftGraphPrintSettings {
	this := MicrosoftGraphPrintSettings{}
	return &this
}

// GetDocumentConversionEnabled returns the DocumentConversionEnabled field value if set, zero value otherwise.
func (o *MicrosoftGraphPrintSettings) GetDocumentConversionEnabled() bool {
	if o == nil || o.DocumentConversionEnabled == nil {
		var ret bool
		return ret
	}
	return *o.DocumentConversionEnabled
}

// GetDocumentConversionEnabledOk returns a tuple with the DocumentConversionEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphPrintSettings) GetDocumentConversionEnabledOk() (*bool, bool) {
	if o == nil || o.DocumentConversionEnabled == nil {
		return nil, false
	}
	return o.DocumentConversionEnabled, true
}

// HasDocumentConversionEnabled returns a boolean if a field has been set.
func (o *MicrosoftGraphPrintSettings) HasDocumentConversionEnabled() bool {
	if o != nil && o.DocumentConversionEnabled != nil {
		return true
	}

	return false
}

// SetDocumentConversionEnabled gets a reference to the given bool and assigns it to the DocumentConversionEnabled field.
func (o *MicrosoftGraphPrintSettings) SetDocumentConversionEnabled(v bool) {
	o.DocumentConversionEnabled = &v
}

func (o MicrosoftGraphPrintSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DocumentConversionEnabled != nil {
		toSerialize["documentConversionEnabled"] = o.DocumentConversionEnabled
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphPrintSettings struct {
	value *MicrosoftGraphPrintSettings
	isSet bool
}

func (v NullableMicrosoftGraphPrintSettings) Get() *MicrosoftGraphPrintSettings {
	return v.value
}

func (v *NullableMicrosoftGraphPrintSettings) Set(val *MicrosoftGraphPrintSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphPrintSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphPrintSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphPrintSettings(val *MicrosoftGraphPrintSettings) *NullableMicrosoftGraphPrintSettings {
	return &NullableMicrosoftGraphPrintSettings{value: val, isSet: true}
}

func (v NullableMicrosoftGraphPrintSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphPrintSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


