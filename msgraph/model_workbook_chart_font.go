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

// WorkbookChartFont struct for WorkbookChartFont
type WorkbookChartFont struct {
	// Represents the bold status of font.
	Bold NullableBool `json:"bold,omitempty"`
	// HTML color code representation of the text color. E.g. #FF0000 represents Red.
	Color NullableString `json:"color,omitempty"`
	// Represents the italic status of the font.
	Italic NullableBool `json:"italic,omitempty"`
	// Font name (e.g. 'Calibri')
	Name NullableString `json:"name,omitempty"`
	// Size of the font (e.g. 11)
	Size AnyOfnumberstringstring `json:"size,omitempty"`
	// Type of underline applied to the font. The possible values are: None, Single.
	Underline NullableString `json:"underline,omitempty"`
}

// NewWorkbookChartFont instantiates a new WorkbookChartFont object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkbookChartFont() *WorkbookChartFont {
	this := WorkbookChartFont{}
	return &this
}

// NewWorkbookChartFontWithDefaults instantiates a new WorkbookChartFont object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkbookChartFontWithDefaults() *WorkbookChartFont {
	this := WorkbookChartFont{}
	return &this
}

// GetBold returns the Bold field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkbookChartFont) GetBold() bool {
	if o == nil || o.Bold.Get() == nil {
		var ret bool
		return ret
	}
	return *o.Bold.Get()
}

// GetBoldOk returns a tuple with the Bold field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkbookChartFont) GetBoldOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Bold.Get(), o.Bold.IsSet()
}

// HasBold returns a boolean if a field has been set.
func (o *WorkbookChartFont) HasBold() bool {
	if o != nil && o.Bold.IsSet() {
		return true
	}

	return false
}

// SetBold gets a reference to the given NullableBool and assigns it to the Bold field.
func (o *WorkbookChartFont) SetBold(v bool) {
	o.Bold.Set(&v)
}
// SetBoldNil sets the value for Bold to be an explicit nil
func (o *WorkbookChartFont) SetBoldNil() {
	o.Bold.Set(nil)
}

// UnsetBold ensures that no value is present for Bold, not even an explicit nil
func (o *WorkbookChartFont) UnsetBold() {
	o.Bold.Unset()
}

// GetColor returns the Color field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkbookChartFont) GetColor() string {
	if o == nil || o.Color.Get() == nil {
		var ret string
		return ret
	}
	return *o.Color.Get()
}

// GetColorOk returns a tuple with the Color field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkbookChartFont) GetColorOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Color.Get(), o.Color.IsSet()
}

// HasColor returns a boolean if a field has been set.
func (o *WorkbookChartFont) HasColor() bool {
	if o != nil && o.Color.IsSet() {
		return true
	}

	return false
}

// SetColor gets a reference to the given NullableString and assigns it to the Color field.
func (o *WorkbookChartFont) SetColor(v string) {
	o.Color.Set(&v)
}
// SetColorNil sets the value for Color to be an explicit nil
func (o *WorkbookChartFont) SetColorNil() {
	o.Color.Set(nil)
}

// UnsetColor ensures that no value is present for Color, not even an explicit nil
func (o *WorkbookChartFont) UnsetColor() {
	o.Color.Unset()
}

// GetItalic returns the Italic field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkbookChartFont) GetItalic() bool {
	if o == nil || o.Italic.Get() == nil {
		var ret bool
		return ret
	}
	return *o.Italic.Get()
}

// GetItalicOk returns a tuple with the Italic field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkbookChartFont) GetItalicOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Italic.Get(), o.Italic.IsSet()
}

// HasItalic returns a boolean if a field has been set.
func (o *WorkbookChartFont) HasItalic() bool {
	if o != nil && o.Italic.IsSet() {
		return true
	}

	return false
}

// SetItalic gets a reference to the given NullableBool and assigns it to the Italic field.
func (o *WorkbookChartFont) SetItalic(v bool) {
	o.Italic.Set(&v)
}
// SetItalicNil sets the value for Italic to be an explicit nil
func (o *WorkbookChartFont) SetItalicNil() {
	o.Italic.Set(nil)
}

// UnsetItalic ensures that no value is present for Italic, not even an explicit nil
func (o *WorkbookChartFont) UnsetItalic() {
	o.Italic.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkbookChartFont) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkbookChartFont) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *WorkbookChartFont) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *WorkbookChartFont) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *WorkbookChartFont) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *WorkbookChartFont) UnsetName() {
	o.Name.Unset()
}

// GetSize returns the Size field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkbookChartFont) GetSize() AnyOfnumberstringstring {
	if o == nil  {
		var ret AnyOfnumberstringstring
		return ret
	}
	return o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkbookChartFont) GetSizeOk() (*AnyOfnumberstringstring, bool) {
	if o == nil || o.Size == nil {
		return nil, false
	}
	return &o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *WorkbookChartFont) HasSize() bool {
	if o != nil && o.Size != nil {
		return true
	}

	return false
}

// SetSize gets a reference to the given AnyOfnumberstringstring and assigns it to the Size field.
func (o *WorkbookChartFont) SetSize(v AnyOfnumberstringstring) {
	o.Size = v
}

// GetUnderline returns the Underline field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkbookChartFont) GetUnderline() string {
	if o == nil || o.Underline.Get() == nil {
		var ret string
		return ret
	}
	return *o.Underline.Get()
}

// GetUnderlineOk returns a tuple with the Underline field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkbookChartFont) GetUnderlineOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Underline.Get(), o.Underline.IsSet()
}

// HasUnderline returns a boolean if a field has been set.
func (o *WorkbookChartFont) HasUnderline() bool {
	if o != nil && o.Underline.IsSet() {
		return true
	}

	return false
}

// SetUnderline gets a reference to the given NullableString and assigns it to the Underline field.
func (o *WorkbookChartFont) SetUnderline(v string) {
	o.Underline.Set(&v)
}
// SetUnderlineNil sets the value for Underline to be an explicit nil
func (o *WorkbookChartFont) SetUnderlineNil() {
	o.Underline.Set(nil)
}

// UnsetUnderline ensures that no value is present for Underline, not even an explicit nil
func (o *WorkbookChartFont) UnsetUnderline() {
	o.Underline.Unset()
}

func (o WorkbookChartFont) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Bold.IsSet() {
		toSerialize["bold"] = o.Bold.Get()
	}
	if o.Color.IsSet() {
		toSerialize["color"] = o.Color.Get()
	}
	if o.Italic.IsSet() {
		toSerialize["italic"] = o.Italic.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Size != nil {
		toSerialize["size"] = o.Size
	}
	if o.Underline.IsSet() {
		toSerialize["underline"] = o.Underline.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableWorkbookChartFont struct {
	value *WorkbookChartFont
	isSet bool
}

func (v NullableWorkbookChartFont) Get() *WorkbookChartFont {
	return v.value
}

func (v *NullableWorkbookChartFont) Set(val *WorkbookChartFont) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkbookChartFont) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkbookChartFont) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkbookChartFont(val *WorkbookChartFont) *NullableWorkbookChartFont {
	return &NullableWorkbookChartFont{value: val, isSet: true}
}

func (v NullableWorkbookChartFont) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkbookChartFont) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


