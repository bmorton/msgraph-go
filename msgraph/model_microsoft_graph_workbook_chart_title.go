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

// MicrosoftGraphWorkbookChartTitle struct for MicrosoftGraphWorkbookChartTitle
type MicrosoftGraphWorkbookChartTitle struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
	// Boolean value representing if the chart title will overlay the chart or not.
	Overlay NullableBool `json:"overlay,omitempty"`
	// Represents the title text of a chart.
	Text NullableString `json:"text,omitempty"`
	// A boolean value the represents the visibility of a chart title object.
	Visible *bool `json:"visible,omitempty"`
	// Represents the formatting of a chart title, which includes fill and font formatting. Read-only.
	Format AnyOfmicrosoftGraphWorkbookChartTitleFormat `json:"format,omitempty"`
}

// NewMicrosoftGraphWorkbookChartTitle instantiates a new MicrosoftGraphWorkbookChartTitle object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphWorkbookChartTitle() *MicrosoftGraphWorkbookChartTitle {
	this := MicrosoftGraphWorkbookChartTitle{}
	return &this
}

// NewMicrosoftGraphWorkbookChartTitleWithDefaults instantiates a new MicrosoftGraphWorkbookChartTitle object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphWorkbookChartTitleWithDefaults() *MicrosoftGraphWorkbookChartTitle {
	this := MicrosoftGraphWorkbookChartTitle{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MicrosoftGraphWorkbookChartTitle) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphWorkbookChartTitle) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphWorkbookChartTitle) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MicrosoftGraphWorkbookChartTitle) SetId(v string) {
	o.Id = &v
}

// GetOverlay returns the Overlay field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphWorkbookChartTitle) GetOverlay() bool {
	if o == nil || o.Overlay.Get() == nil {
		var ret bool
		return ret
	}
	return *o.Overlay.Get()
}

// GetOverlayOk returns a tuple with the Overlay field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphWorkbookChartTitle) GetOverlayOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Overlay.Get(), o.Overlay.IsSet()
}

// HasOverlay returns a boolean if a field has been set.
func (o *MicrosoftGraphWorkbookChartTitle) HasOverlay() bool {
	if o != nil && o.Overlay.IsSet() {
		return true
	}

	return false
}

// SetOverlay gets a reference to the given NullableBool and assigns it to the Overlay field.
func (o *MicrosoftGraphWorkbookChartTitle) SetOverlay(v bool) {
	o.Overlay.Set(&v)
}
// SetOverlayNil sets the value for Overlay to be an explicit nil
func (o *MicrosoftGraphWorkbookChartTitle) SetOverlayNil() {
	o.Overlay.Set(nil)
}

// UnsetOverlay ensures that no value is present for Overlay, not even an explicit nil
func (o *MicrosoftGraphWorkbookChartTitle) UnsetOverlay() {
	o.Overlay.Unset()
}

// GetText returns the Text field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphWorkbookChartTitle) GetText() string {
	if o == nil || o.Text.Get() == nil {
		var ret string
		return ret
	}
	return *o.Text.Get()
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphWorkbookChartTitle) GetTextOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Text.Get(), o.Text.IsSet()
}

// HasText returns a boolean if a field has been set.
func (o *MicrosoftGraphWorkbookChartTitle) HasText() bool {
	if o != nil && o.Text.IsSet() {
		return true
	}

	return false
}

// SetText gets a reference to the given NullableString and assigns it to the Text field.
func (o *MicrosoftGraphWorkbookChartTitle) SetText(v string) {
	o.Text.Set(&v)
}
// SetTextNil sets the value for Text to be an explicit nil
func (o *MicrosoftGraphWorkbookChartTitle) SetTextNil() {
	o.Text.Set(nil)
}

// UnsetText ensures that no value is present for Text, not even an explicit nil
func (o *MicrosoftGraphWorkbookChartTitle) UnsetText() {
	o.Text.Unset()
}

// GetVisible returns the Visible field value if set, zero value otherwise.
func (o *MicrosoftGraphWorkbookChartTitle) GetVisible() bool {
	if o == nil || o.Visible == nil {
		var ret bool
		return ret
	}
	return *o.Visible
}

// GetVisibleOk returns a tuple with the Visible field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphWorkbookChartTitle) GetVisibleOk() (*bool, bool) {
	if o == nil || o.Visible == nil {
		return nil, false
	}
	return o.Visible, true
}

// HasVisible returns a boolean if a field has been set.
func (o *MicrosoftGraphWorkbookChartTitle) HasVisible() bool {
	if o != nil && o.Visible != nil {
		return true
	}

	return false
}

// SetVisible gets a reference to the given bool and assigns it to the Visible field.
func (o *MicrosoftGraphWorkbookChartTitle) SetVisible(v bool) {
	o.Visible = &v
}

// GetFormat returns the Format field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphWorkbookChartTitle) GetFormat() AnyOfmicrosoftGraphWorkbookChartTitleFormat {
	if o == nil  {
		var ret AnyOfmicrosoftGraphWorkbookChartTitleFormat
		return ret
	}
	return o.Format
}

// GetFormatOk returns a tuple with the Format field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphWorkbookChartTitle) GetFormatOk() (*AnyOfmicrosoftGraphWorkbookChartTitleFormat, bool) {
	if o == nil || o.Format == nil {
		return nil, false
	}
	return &o.Format, true
}

// HasFormat returns a boolean if a field has been set.
func (o *MicrosoftGraphWorkbookChartTitle) HasFormat() bool {
	if o != nil && o.Format != nil {
		return true
	}

	return false
}

// SetFormat gets a reference to the given AnyOfmicrosoftGraphWorkbookChartTitleFormat and assigns it to the Format field.
func (o *MicrosoftGraphWorkbookChartTitle) SetFormat(v AnyOfmicrosoftGraphWorkbookChartTitleFormat) {
	o.Format = v
}

func (o MicrosoftGraphWorkbookChartTitle) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Overlay.IsSet() {
		toSerialize["overlay"] = o.Overlay.Get()
	}
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

type NullableMicrosoftGraphWorkbookChartTitle struct {
	value *MicrosoftGraphWorkbookChartTitle
	isSet bool
}

func (v NullableMicrosoftGraphWorkbookChartTitle) Get() *MicrosoftGraphWorkbookChartTitle {
	return v.value
}

func (v *NullableMicrosoftGraphWorkbookChartTitle) Set(val *MicrosoftGraphWorkbookChartTitle) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphWorkbookChartTitle) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphWorkbookChartTitle) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphWorkbookChartTitle(val *MicrosoftGraphWorkbookChartTitle) *NullableMicrosoftGraphWorkbookChartTitle {
	return &NullableMicrosoftGraphWorkbookChartTitle{value: val, isSet: true}
}

func (v NullableMicrosoftGraphWorkbookChartTitle) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphWorkbookChartTitle) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


