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

// WorkbookChartAxisTitleFormat struct for WorkbookChartAxisTitleFormat
type WorkbookChartAxisTitleFormat struct {
	// Represents the font attributes, such as font name, font size, color, etc. of chart axis title object. Read-only.
	Font AnyOfmicrosoftGraphWorkbookChartFont `json:"font,omitempty"`
}

// NewWorkbookChartAxisTitleFormat instantiates a new WorkbookChartAxisTitleFormat object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkbookChartAxisTitleFormat() *WorkbookChartAxisTitleFormat {
	this := WorkbookChartAxisTitleFormat{}
	return &this
}

// NewWorkbookChartAxisTitleFormatWithDefaults instantiates a new WorkbookChartAxisTitleFormat object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkbookChartAxisTitleFormatWithDefaults() *WorkbookChartAxisTitleFormat {
	this := WorkbookChartAxisTitleFormat{}
	return &this
}

// GetFont returns the Font field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkbookChartAxisTitleFormat) GetFont() AnyOfmicrosoftGraphWorkbookChartFont {
	if o == nil  {
		var ret AnyOfmicrosoftGraphWorkbookChartFont
		return ret
	}
	return o.Font
}

// GetFontOk returns a tuple with the Font field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkbookChartAxisTitleFormat) GetFontOk() (*AnyOfmicrosoftGraphWorkbookChartFont, bool) {
	if o == nil || o.Font == nil {
		return nil, false
	}
	return &o.Font, true
}

// HasFont returns a boolean if a field has been set.
func (o *WorkbookChartAxisTitleFormat) HasFont() bool {
	if o != nil && o.Font != nil {
		return true
	}

	return false
}

// SetFont gets a reference to the given AnyOfmicrosoftGraphWorkbookChartFont and assigns it to the Font field.
func (o *WorkbookChartAxisTitleFormat) SetFont(v AnyOfmicrosoftGraphWorkbookChartFont) {
	o.Font = v
}

func (o WorkbookChartAxisTitleFormat) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Font != nil {
		toSerialize["font"] = o.Font
	}
	return json.Marshal(toSerialize)
}

type NullableWorkbookChartAxisTitleFormat struct {
	value *WorkbookChartAxisTitleFormat
	isSet bool
}

func (v NullableWorkbookChartAxisTitleFormat) Get() *WorkbookChartAxisTitleFormat {
	return v.value
}

func (v *NullableWorkbookChartAxisTitleFormat) Set(val *WorkbookChartAxisTitleFormat) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkbookChartAxisTitleFormat) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkbookChartAxisTitleFormat) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkbookChartAxisTitleFormat(val *WorkbookChartAxisTitleFormat) *NullableWorkbookChartAxisTitleFormat {
	return &NullableWorkbookChartAxisTitleFormat{value: val, isSet: true}
}

func (v NullableWorkbookChartAxisTitleFormat) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkbookChartAxisTitleFormat) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

