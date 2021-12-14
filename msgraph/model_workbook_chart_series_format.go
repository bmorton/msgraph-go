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

// WorkbookChartSeriesFormat struct for WorkbookChartSeriesFormat
type WorkbookChartSeriesFormat struct {
	// Represents the fill format of a chart series, which includes background formating information. Read-only.
	Fill AnyOfmicrosoftGraphWorkbookChartFill `json:"fill,omitempty"`
	// Represents line formatting. Read-only.
	Line AnyOfmicrosoftGraphWorkbookChartLineFormat `json:"line,omitempty"`
}

// NewWorkbookChartSeriesFormat instantiates a new WorkbookChartSeriesFormat object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkbookChartSeriesFormat() *WorkbookChartSeriesFormat {
	this := WorkbookChartSeriesFormat{}
	return &this
}

// NewWorkbookChartSeriesFormatWithDefaults instantiates a new WorkbookChartSeriesFormat object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkbookChartSeriesFormatWithDefaults() *WorkbookChartSeriesFormat {
	this := WorkbookChartSeriesFormat{}
	return &this
}

// GetFill returns the Fill field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkbookChartSeriesFormat) GetFill() AnyOfmicrosoftGraphWorkbookChartFill {
	if o == nil  {
		var ret AnyOfmicrosoftGraphWorkbookChartFill
		return ret
	}
	return o.Fill
}

// GetFillOk returns a tuple with the Fill field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkbookChartSeriesFormat) GetFillOk() (*AnyOfmicrosoftGraphWorkbookChartFill, bool) {
	if o == nil || o.Fill == nil {
		return nil, false
	}
	return &o.Fill, true
}

// HasFill returns a boolean if a field has been set.
func (o *WorkbookChartSeriesFormat) HasFill() bool {
	if o != nil && o.Fill != nil {
		return true
	}

	return false
}

// SetFill gets a reference to the given AnyOfmicrosoftGraphWorkbookChartFill and assigns it to the Fill field.
func (o *WorkbookChartSeriesFormat) SetFill(v AnyOfmicrosoftGraphWorkbookChartFill) {
	o.Fill = v
}

// GetLine returns the Line field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkbookChartSeriesFormat) GetLine() AnyOfmicrosoftGraphWorkbookChartLineFormat {
	if o == nil  {
		var ret AnyOfmicrosoftGraphWorkbookChartLineFormat
		return ret
	}
	return o.Line
}

// GetLineOk returns a tuple with the Line field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkbookChartSeriesFormat) GetLineOk() (*AnyOfmicrosoftGraphWorkbookChartLineFormat, bool) {
	if o == nil || o.Line == nil {
		return nil, false
	}
	return &o.Line, true
}

// HasLine returns a boolean if a field has been set.
func (o *WorkbookChartSeriesFormat) HasLine() bool {
	if o != nil && o.Line != nil {
		return true
	}

	return false
}

// SetLine gets a reference to the given AnyOfmicrosoftGraphWorkbookChartLineFormat and assigns it to the Line field.
func (o *WorkbookChartSeriesFormat) SetLine(v AnyOfmicrosoftGraphWorkbookChartLineFormat) {
	o.Line = v
}

func (o WorkbookChartSeriesFormat) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Fill != nil {
		toSerialize["fill"] = o.Fill
	}
	if o.Line != nil {
		toSerialize["line"] = o.Line
	}
	return json.Marshal(toSerialize)
}

type NullableWorkbookChartSeriesFormat struct {
	value *WorkbookChartSeriesFormat
	isSet bool
}

func (v NullableWorkbookChartSeriesFormat) Get() *WorkbookChartSeriesFormat {
	return v.value
}

func (v *NullableWorkbookChartSeriesFormat) Set(val *WorkbookChartSeriesFormat) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkbookChartSeriesFormat) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkbookChartSeriesFormat) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkbookChartSeriesFormat(val *WorkbookChartSeriesFormat) *NullableWorkbookChartSeriesFormat {
	return &NullableWorkbookChartSeriesFormat{value: val, isSet: true}
}

func (v NullableWorkbookChartSeriesFormat) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkbookChartSeriesFormat) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


