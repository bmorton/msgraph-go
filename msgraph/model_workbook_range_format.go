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

// WorkbookRangeFormat struct for WorkbookRangeFormat
type WorkbookRangeFormat struct {
	// Gets or sets the width of all colums within the range. If the column widths are not uniform, null will be returned.
	ColumnWidth AnyOfnumberstringstring `json:"columnWidth,omitempty"`
	// Represents the horizontal alignment for the specified object. The possible values are: General, Left, Center, Right, Fill, Justify, CenterAcrossSelection, Distributed.
	HorizontalAlignment NullableString `json:"horizontalAlignment,omitempty"`
	// Gets or sets the height of all rows in the range. If the row heights are not uniform null will be returned.
	RowHeight AnyOfnumberstringstring `json:"rowHeight,omitempty"`
	// Represents the vertical alignment for the specified object. The possible values are: Top, Center, Bottom, Justify, Distributed.
	VerticalAlignment NullableString `json:"verticalAlignment,omitempty"`
	// Indicates if Excel wraps the text in the object. A null value indicates that the entire range doesn't have uniform wrap setting
	WrapText NullableBool `json:"wrapText,omitempty"`
	// Collection of border objects that apply to the overall range selected Read-only.
	Borders *[]MicrosoftGraphWorkbookRangeBorder `json:"borders,omitempty"`
	// Returns the fill object defined on the overall range. Read-only.
	Fill AnyOfmicrosoftGraphWorkbookRangeFill `json:"fill,omitempty"`
	// Returns the font object defined on the overall range selected Read-only.
	Font AnyOfmicrosoftGraphWorkbookRangeFont `json:"font,omitempty"`
	// Returns the format protection object for a range. Read-only.
	Protection AnyOfmicrosoftGraphWorkbookFormatProtection `json:"protection,omitempty"`
}

// NewWorkbookRangeFormat instantiates a new WorkbookRangeFormat object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkbookRangeFormat() *WorkbookRangeFormat {
	this := WorkbookRangeFormat{}
	return &this
}

// NewWorkbookRangeFormatWithDefaults instantiates a new WorkbookRangeFormat object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkbookRangeFormatWithDefaults() *WorkbookRangeFormat {
	this := WorkbookRangeFormat{}
	return &this
}

// GetColumnWidth returns the ColumnWidth field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkbookRangeFormat) GetColumnWidth() AnyOfnumberstringstring {
	if o == nil  {
		var ret AnyOfnumberstringstring
		return ret
	}
	return o.ColumnWidth
}

// GetColumnWidthOk returns a tuple with the ColumnWidth field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkbookRangeFormat) GetColumnWidthOk() (*AnyOfnumberstringstring, bool) {
	if o == nil || o.ColumnWidth == nil {
		return nil, false
	}
	return &o.ColumnWidth, true
}

// HasColumnWidth returns a boolean if a field has been set.
func (o *WorkbookRangeFormat) HasColumnWidth() bool {
	if o != nil && o.ColumnWidth != nil {
		return true
	}

	return false
}

// SetColumnWidth gets a reference to the given AnyOfnumberstringstring and assigns it to the ColumnWidth field.
func (o *WorkbookRangeFormat) SetColumnWidth(v AnyOfnumberstringstring) {
	o.ColumnWidth = v
}

// GetHorizontalAlignment returns the HorizontalAlignment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkbookRangeFormat) GetHorizontalAlignment() string {
	if o == nil || o.HorizontalAlignment.Get() == nil {
		var ret string
		return ret
	}
	return *o.HorizontalAlignment.Get()
}

// GetHorizontalAlignmentOk returns a tuple with the HorizontalAlignment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkbookRangeFormat) GetHorizontalAlignmentOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.HorizontalAlignment.Get(), o.HorizontalAlignment.IsSet()
}

// HasHorizontalAlignment returns a boolean if a field has been set.
func (o *WorkbookRangeFormat) HasHorizontalAlignment() bool {
	if o != nil && o.HorizontalAlignment.IsSet() {
		return true
	}

	return false
}

// SetHorizontalAlignment gets a reference to the given NullableString and assigns it to the HorizontalAlignment field.
func (o *WorkbookRangeFormat) SetHorizontalAlignment(v string) {
	o.HorizontalAlignment.Set(&v)
}
// SetHorizontalAlignmentNil sets the value for HorizontalAlignment to be an explicit nil
func (o *WorkbookRangeFormat) SetHorizontalAlignmentNil() {
	o.HorizontalAlignment.Set(nil)
}

// UnsetHorizontalAlignment ensures that no value is present for HorizontalAlignment, not even an explicit nil
func (o *WorkbookRangeFormat) UnsetHorizontalAlignment() {
	o.HorizontalAlignment.Unset()
}

// GetRowHeight returns the RowHeight field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkbookRangeFormat) GetRowHeight() AnyOfnumberstringstring {
	if o == nil  {
		var ret AnyOfnumberstringstring
		return ret
	}
	return o.RowHeight
}

// GetRowHeightOk returns a tuple with the RowHeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkbookRangeFormat) GetRowHeightOk() (*AnyOfnumberstringstring, bool) {
	if o == nil || o.RowHeight == nil {
		return nil, false
	}
	return &o.RowHeight, true
}

// HasRowHeight returns a boolean if a field has been set.
func (o *WorkbookRangeFormat) HasRowHeight() bool {
	if o != nil && o.RowHeight != nil {
		return true
	}

	return false
}

// SetRowHeight gets a reference to the given AnyOfnumberstringstring and assigns it to the RowHeight field.
func (o *WorkbookRangeFormat) SetRowHeight(v AnyOfnumberstringstring) {
	o.RowHeight = v
}

// GetVerticalAlignment returns the VerticalAlignment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkbookRangeFormat) GetVerticalAlignment() string {
	if o == nil || o.VerticalAlignment.Get() == nil {
		var ret string
		return ret
	}
	return *o.VerticalAlignment.Get()
}

// GetVerticalAlignmentOk returns a tuple with the VerticalAlignment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkbookRangeFormat) GetVerticalAlignmentOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.VerticalAlignment.Get(), o.VerticalAlignment.IsSet()
}

// HasVerticalAlignment returns a boolean if a field has been set.
func (o *WorkbookRangeFormat) HasVerticalAlignment() bool {
	if o != nil && o.VerticalAlignment.IsSet() {
		return true
	}

	return false
}

// SetVerticalAlignment gets a reference to the given NullableString and assigns it to the VerticalAlignment field.
func (o *WorkbookRangeFormat) SetVerticalAlignment(v string) {
	o.VerticalAlignment.Set(&v)
}
// SetVerticalAlignmentNil sets the value for VerticalAlignment to be an explicit nil
func (o *WorkbookRangeFormat) SetVerticalAlignmentNil() {
	o.VerticalAlignment.Set(nil)
}

// UnsetVerticalAlignment ensures that no value is present for VerticalAlignment, not even an explicit nil
func (o *WorkbookRangeFormat) UnsetVerticalAlignment() {
	o.VerticalAlignment.Unset()
}

// GetWrapText returns the WrapText field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkbookRangeFormat) GetWrapText() bool {
	if o == nil || o.WrapText.Get() == nil {
		var ret bool
		return ret
	}
	return *o.WrapText.Get()
}

// GetWrapTextOk returns a tuple with the WrapText field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkbookRangeFormat) GetWrapTextOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.WrapText.Get(), o.WrapText.IsSet()
}

// HasWrapText returns a boolean if a field has been set.
func (o *WorkbookRangeFormat) HasWrapText() bool {
	if o != nil && o.WrapText.IsSet() {
		return true
	}

	return false
}

// SetWrapText gets a reference to the given NullableBool and assigns it to the WrapText field.
func (o *WorkbookRangeFormat) SetWrapText(v bool) {
	o.WrapText.Set(&v)
}
// SetWrapTextNil sets the value for WrapText to be an explicit nil
func (o *WorkbookRangeFormat) SetWrapTextNil() {
	o.WrapText.Set(nil)
}

// UnsetWrapText ensures that no value is present for WrapText, not even an explicit nil
func (o *WorkbookRangeFormat) UnsetWrapText() {
	o.WrapText.Unset()
}

// GetBorders returns the Borders field value if set, zero value otherwise.
func (o *WorkbookRangeFormat) GetBorders() []MicrosoftGraphWorkbookRangeBorder {
	if o == nil || o.Borders == nil {
		var ret []MicrosoftGraphWorkbookRangeBorder
		return ret
	}
	return *o.Borders
}

// GetBordersOk returns a tuple with the Borders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkbookRangeFormat) GetBordersOk() (*[]MicrosoftGraphWorkbookRangeBorder, bool) {
	if o == nil || o.Borders == nil {
		return nil, false
	}
	return o.Borders, true
}

// HasBorders returns a boolean if a field has been set.
func (o *WorkbookRangeFormat) HasBorders() bool {
	if o != nil && o.Borders != nil {
		return true
	}

	return false
}

// SetBorders gets a reference to the given []MicrosoftGraphWorkbookRangeBorder and assigns it to the Borders field.
func (o *WorkbookRangeFormat) SetBorders(v []MicrosoftGraphWorkbookRangeBorder) {
	o.Borders = &v
}

// GetFill returns the Fill field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkbookRangeFormat) GetFill() AnyOfmicrosoftGraphWorkbookRangeFill {
	if o == nil  {
		var ret AnyOfmicrosoftGraphWorkbookRangeFill
		return ret
	}
	return o.Fill
}

// GetFillOk returns a tuple with the Fill field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkbookRangeFormat) GetFillOk() (*AnyOfmicrosoftGraphWorkbookRangeFill, bool) {
	if o == nil || o.Fill == nil {
		return nil, false
	}
	return &o.Fill, true
}

// HasFill returns a boolean if a field has been set.
func (o *WorkbookRangeFormat) HasFill() bool {
	if o != nil && o.Fill != nil {
		return true
	}

	return false
}

// SetFill gets a reference to the given AnyOfmicrosoftGraphWorkbookRangeFill and assigns it to the Fill field.
func (o *WorkbookRangeFormat) SetFill(v AnyOfmicrosoftGraphWorkbookRangeFill) {
	o.Fill = v
}

// GetFont returns the Font field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkbookRangeFormat) GetFont() AnyOfmicrosoftGraphWorkbookRangeFont {
	if o == nil  {
		var ret AnyOfmicrosoftGraphWorkbookRangeFont
		return ret
	}
	return o.Font
}

// GetFontOk returns a tuple with the Font field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkbookRangeFormat) GetFontOk() (*AnyOfmicrosoftGraphWorkbookRangeFont, bool) {
	if o == nil || o.Font == nil {
		return nil, false
	}
	return &o.Font, true
}

// HasFont returns a boolean if a field has been set.
func (o *WorkbookRangeFormat) HasFont() bool {
	if o != nil && o.Font != nil {
		return true
	}

	return false
}

// SetFont gets a reference to the given AnyOfmicrosoftGraphWorkbookRangeFont and assigns it to the Font field.
func (o *WorkbookRangeFormat) SetFont(v AnyOfmicrosoftGraphWorkbookRangeFont) {
	o.Font = v
}

// GetProtection returns the Protection field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkbookRangeFormat) GetProtection() AnyOfmicrosoftGraphWorkbookFormatProtection {
	if o == nil  {
		var ret AnyOfmicrosoftGraphWorkbookFormatProtection
		return ret
	}
	return o.Protection
}

// GetProtectionOk returns a tuple with the Protection field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkbookRangeFormat) GetProtectionOk() (*AnyOfmicrosoftGraphWorkbookFormatProtection, bool) {
	if o == nil || o.Protection == nil {
		return nil, false
	}
	return &o.Protection, true
}

// HasProtection returns a boolean if a field has been set.
func (o *WorkbookRangeFormat) HasProtection() bool {
	if o != nil && o.Protection != nil {
		return true
	}

	return false
}

// SetProtection gets a reference to the given AnyOfmicrosoftGraphWorkbookFormatProtection and assigns it to the Protection field.
func (o *WorkbookRangeFormat) SetProtection(v AnyOfmicrosoftGraphWorkbookFormatProtection) {
	o.Protection = v
}

func (o WorkbookRangeFormat) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ColumnWidth != nil {
		toSerialize["columnWidth"] = o.ColumnWidth
	}
	if o.HorizontalAlignment.IsSet() {
		toSerialize["horizontalAlignment"] = o.HorizontalAlignment.Get()
	}
	if o.RowHeight != nil {
		toSerialize["rowHeight"] = o.RowHeight
	}
	if o.VerticalAlignment.IsSet() {
		toSerialize["verticalAlignment"] = o.VerticalAlignment.Get()
	}
	if o.WrapText.IsSet() {
		toSerialize["wrapText"] = o.WrapText.Get()
	}
	if o.Borders != nil {
		toSerialize["borders"] = o.Borders
	}
	if o.Fill != nil {
		toSerialize["fill"] = o.Fill
	}
	if o.Font != nil {
		toSerialize["font"] = o.Font
	}
	if o.Protection != nil {
		toSerialize["protection"] = o.Protection
	}
	return json.Marshal(toSerialize)
}

type NullableWorkbookRangeFormat struct {
	value *WorkbookRangeFormat
	isSet bool
}

func (v NullableWorkbookRangeFormat) Get() *WorkbookRangeFormat {
	return v.value
}

func (v *NullableWorkbookRangeFormat) Set(val *WorkbookRangeFormat) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkbookRangeFormat) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkbookRangeFormat) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkbookRangeFormat(val *WorkbookRangeFormat) *NullableWorkbookRangeFormat {
	return &NullableWorkbookRangeFormat{value: val, isSet: true}
}

func (v NullableWorkbookRangeFormat) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkbookRangeFormat) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


