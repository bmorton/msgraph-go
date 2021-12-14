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

// WorkbookChartAxisTitle struct for WorkbookChartAxisTitle
type WorkbookChartAxisTitle struct {
	// Represents the axis title.
	Text NullableString `json:"text,omitempty"`
	// A boolean that specifies the visibility of an axis title.
	Visible *bool `json:"visible,omitempty"`
	// Represents the formatting of chart axis title. Read-only.
	Format AnyOfmicrosoftGraphWorkbookChartAxisTitleFormat `json:"format,omitempty"`
}

// NewWorkbookChartAxisTitle instantiates a new WorkbookChartAxisTitle object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkbookChartAxisTitle() *WorkbookChartAxisTitle {
	this := WorkbookChartAxisTitle{}
	return &this
}

// NewWorkbookChartAxisTitleWithDefaults instantiates a new WorkbookChartAxisTitle object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkbookChartAxisTitleWithDefaults() *WorkbookChartAxisTitle {
	this := WorkbookChartAxisTitle{}
	return &this
}

// GetText returns the Text field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkbookChartAxisTitle) GetText() string {
	if o == nil || o.Text.Get() == nil {
		var ret string
		return ret
	}
	return *o.Text.Get()
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkbookChartAxisTitle) GetTextOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Text.Get(), o.Text.IsSet()
}

// HasText returns a boolean if a field has been set.
func (o *WorkbookChartAxisTitle) HasText() bool {
	if o != nil && o.Text.IsSet() {
		return true
	}

	return false
}

// SetText gets a reference to the given NullableString and assigns it to the Text field.
func (o *WorkbookChartAxisTitle) SetText(v string) {
	o.Text.Set(&v)
}
// SetTextNil sets the value for Text to be an explicit nil
func (o *WorkbookChartAxisTitle) SetTextNil() {
	o.Text.Set(nil)
}

// UnsetText ensures that no value is present for Text, not even an explicit nil
func (o *WorkbookChartAxisTitle) UnsetText() {
	o.Text.Unset()
}

// GetVisible returns the Visible field value if set, zero value otherwise.
func (o *WorkbookChartAxisTitle) GetVisible() bool {
	if o == nil || o.Visible == nil {
		var ret bool
		return ret
	}
	return *o.Visible
}

// GetVisibleOk returns a tuple with the Visible field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkbookChartAxisTitle) GetVisibleOk() (*bool, bool) {
	if o == nil || o.Visible == nil {
		return nil, false
	}
	return o.Visible, true
}

// HasVisible returns a boolean if a field has been set.
func (o *WorkbookChartAxisTitle) HasVisible() bool {
	if o != nil && o.Visible != nil {
		return true
	}

	return false
}

// SetVisible gets a reference to the given bool and assigns it to the Visible field.
func (o *WorkbookChartAxisTitle) SetVisible(v bool) {
	o.Visible = &v
}

// GetFormat returns the Format field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkbookChartAxisTitle) GetFormat() AnyOfmicrosoftGraphWorkbookChartAxisTitleFormat {
	if o == nil  {
		var ret AnyOfmicrosoftGraphWorkbookChartAxisTitleFormat
		return ret
	}
	return o.Format
}

// GetFormatOk returns a tuple with the Format field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkbookChartAxisTitle) GetFormatOk() (*AnyOfmicrosoftGraphWorkbookChartAxisTitleFormat, bool) {
	if o == nil || o.Format == nil {
		return nil, false
	}
	return &o.Format, true
}

// HasFormat returns a boolean if a field has been set.
func (o *WorkbookChartAxisTitle) HasFormat() bool {
	if o != nil && o.Format != nil {
		return true
	}

	return false
}

// SetFormat gets a reference to the given AnyOfmicrosoftGraphWorkbookChartAxisTitleFormat and assigns it to the Format field.
func (o *WorkbookChartAxisTitle) SetFormat(v AnyOfmicrosoftGraphWorkbookChartAxisTitleFormat) {
	o.Format = v
}

func (o WorkbookChartAxisTitle) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Text.IsSet() {
		toSerialize["text"] = o.Text.Get()
	}
	if o.Visible != nil {
		toSerialize["visible"] = o.Visible
	}
	if o.Format != nil {
		toSerialize["format"] = o.Format
	}
	return json.Marshal(toSerialize)
}

type NullableWorkbookChartAxisTitle struct {
	value *WorkbookChartAxisTitle
	isSet bool
}

func (v NullableWorkbookChartAxisTitle) Get() *WorkbookChartAxisTitle {
	return v.value
}

func (v *NullableWorkbookChartAxisTitle) Set(val *WorkbookChartAxisTitle) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkbookChartAxisTitle) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkbookChartAxisTitle) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkbookChartAxisTitle(val *WorkbookChartAxisTitle) *NullableWorkbookChartAxisTitle {
	return &NullableWorkbookChartAxisTitle{value: val, isSet: true}
}

func (v NullableWorkbookChartAxisTitle) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkbookChartAxisTitle) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


