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

// WorkbookChartLegendFormat struct for WorkbookChartLegendFormat
type WorkbookChartLegendFormat struct {
	// Represents the fill format of an object, which includes background formating information. Read-only.
	Fill AnyOfmicrosoftGraphWorkbookChartFill `json:"fill,omitempty"`
	// Represents the font attributes such as font name, font size, color, etc. of a chart legend. Read-only.
	Font AnyOfmicrosoftGraphWorkbookChartFont `json:"font,omitempty"`
}

// NewWorkbookChartLegendFormat instantiates a new WorkbookChartLegendFormat object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkbookChartLegendFormat() *WorkbookChartLegendFormat {
	this := WorkbookChartLegendFormat{}
	return &this
}

// NewWorkbookChartLegendFormatWithDefaults instantiates a new WorkbookChartLegendFormat object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkbookChartLegendFormatWithDefaults() *WorkbookChartLegendFormat {
	this := WorkbookChartLegendFormat{}
	return &this
}

// GetFill returns the Fill field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkbookChartLegendFormat) GetFill() AnyOfmicrosoftGraphWorkbookChartFill {
	if o == nil  {
		var ret AnyOfmicrosoftGraphWorkbookChartFill
		return ret
	}
	return o.Fill
}

// GetFillOk returns a tuple with the Fill field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkbookChartLegendFormat) GetFillOk() (*AnyOfmicrosoftGraphWorkbookChartFill, bool) {
	if o == nil || o.Fill == nil {
		return nil, false
	}
	return &o.Fill, true
}

// HasFill returns a boolean if a field has been set.
func (o *WorkbookChartLegendFormat) HasFill() bool {
	if o != nil && o.Fill != nil {
		return true
	}

	return false
}

// SetFill gets a reference to the given AnyOfmicrosoftGraphWorkbookChartFill and assigns it to the Fill field.
func (o *WorkbookChartLegendFormat) SetFill(v AnyOfmicrosoftGraphWorkbookChartFill) {
	o.Fill = v
}

// GetFont returns the Font field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkbookChartLegendFormat) GetFont() AnyOfmicrosoftGraphWorkbookChartFont {
	if o == nil  {
		var ret AnyOfmicrosoftGraphWorkbookChartFont
		return ret
	}
	return o.Font
}

// GetFontOk returns a tuple with the Font field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkbookChartLegendFormat) GetFontOk() (*AnyOfmicrosoftGraphWorkbookChartFont, bool) {
	if o == nil || o.Font == nil {
		return nil, false
	}
	return &o.Font, true
}

// HasFont returns a boolean if a field has been set.
func (o *WorkbookChartLegendFormat) HasFont() bool {
	if o != nil && o.Font != nil {
		return true
	}

	return false
}

// SetFont gets a reference to the given AnyOfmicrosoftGraphWorkbookChartFont and assigns it to the Font field.
func (o *WorkbookChartLegendFormat) SetFont(v AnyOfmicrosoftGraphWorkbookChartFont) {
	o.Font = v
}

func (o WorkbookChartLegendFormat) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Fill != nil {
		toSerialize["fill"] = o.Fill
	}
	if o.Font != nil {
		toSerialize["font"] = o.Font
	}
	return json.Marshal(toSerialize)
}

type NullableWorkbookChartLegendFormat struct {
	value *WorkbookChartLegendFormat
	isSet bool
}

func (v NullableWorkbookChartLegendFormat) Get() *WorkbookChartLegendFormat {
	return v.value
}

func (v *NullableWorkbookChartLegendFormat) Set(val *WorkbookChartLegendFormat) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkbookChartLegendFormat) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkbookChartLegendFormat) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkbookChartLegendFormat(val *WorkbookChartLegendFormat) *NullableWorkbookChartLegendFormat {
	return &NullableWorkbookChartLegendFormat{value: val, isSet: true}
}

func (v NullableWorkbookChartLegendFormat) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkbookChartLegendFormat) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


