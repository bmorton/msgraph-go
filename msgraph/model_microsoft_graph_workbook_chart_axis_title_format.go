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

// MicrosoftGraphWorkbookChartAxisTitleFormat struct for MicrosoftGraphWorkbookChartAxisTitleFormat
type MicrosoftGraphWorkbookChartAxisTitleFormat struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
	// Represents the font attributes, such as font name, font size, color, etc. of chart axis title object. Read-only.
	Font AnyOfmicrosoftGraphWorkbookChartFont `json:"font,omitempty"`
}

// NewMicrosoftGraphWorkbookChartAxisTitleFormat instantiates a new MicrosoftGraphWorkbookChartAxisTitleFormat object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphWorkbookChartAxisTitleFormat() *MicrosoftGraphWorkbookChartAxisTitleFormat {
	this := MicrosoftGraphWorkbookChartAxisTitleFormat{}
	return &this
}

// NewMicrosoftGraphWorkbookChartAxisTitleFormatWithDefaults instantiates a new MicrosoftGraphWorkbookChartAxisTitleFormat object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphWorkbookChartAxisTitleFormatWithDefaults() *MicrosoftGraphWorkbookChartAxisTitleFormat {
	this := MicrosoftGraphWorkbookChartAxisTitleFormat{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MicrosoftGraphWorkbookChartAxisTitleFormat) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphWorkbookChartAxisTitleFormat) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphWorkbookChartAxisTitleFormat) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MicrosoftGraphWorkbookChartAxisTitleFormat) SetId(v string) {
	o.Id = &v
}

// GetFont returns the Font field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphWorkbookChartAxisTitleFormat) GetFont() AnyOfmicrosoftGraphWorkbookChartFont {
	if o == nil  {
		var ret AnyOfmicrosoftGraphWorkbookChartFont
		return ret
	}
	return o.Font
}

// GetFontOk returns a tuple with the Font field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphWorkbookChartAxisTitleFormat) GetFontOk() (*AnyOfmicrosoftGraphWorkbookChartFont, bool) {
	if o == nil || o.Font == nil {
		return nil, false
	}
	return &o.Font, true
}

// HasFont returns a boolean if a field has been set.
func (o *MicrosoftGraphWorkbookChartAxisTitleFormat) HasFont() bool {
	if o != nil && o.Font != nil {
		return true
	}

	return false
}

// SetFont gets a reference to the given AnyOfmicrosoftGraphWorkbookChartFont and assigns it to the Font field.
func (o *MicrosoftGraphWorkbookChartAxisTitleFormat) SetFont(v AnyOfmicrosoftGraphWorkbookChartFont) {
	o.Font = v
}

func (o MicrosoftGraphWorkbookChartAxisTitleFormat) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Font != nil {
		toSerialize["font"] = o.Font
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphWorkbookChartAxisTitleFormat struct {
	value *MicrosoftGraphWorkbookChartAxisTitleFormat
	isSet bool
}

func (v NullableMicrosoftGraphWorkbookChartAxisTitleFormat) Get() *MicrosoftGraphWorkbookChartAxisTitleFormat {
	return v.value
}

func (v *NullableMicrosoftGraphWorkbookChartAxisTitleFormat) Set(val *MicrosoftGraphWorkbookChartAxisTitleFormat) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphWorkbookChartAxisTitleFormat) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphWorkbookChartAxisTitleFormat) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphWorkbookChartAxisTitleFormat(val *MicrosoftGraphWorkbookChartAxisTitleFormat) *NullableMicrosoftGraphWorkbookChartAxisTitleFormat {
	return &NullableMicrosoftGraphWorkbookChartAxisTitleFormat{value: val, isSet: true}
}

func (v NullableMicrosoftGraphWorkbookChartAxisTitleFormat) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphWorkbookChartAxisTitleFormat) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


