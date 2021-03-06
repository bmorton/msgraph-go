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

// WorkbookRangeFont struct for WorkbookRangeFont
type WorkbookRangeFont struct {
	// Represents the bold status of font.
	Bold NullableBool `json:"bold,omitempty"`
	// HTML color code representation of the text color. E.g. #FF0000 represents Red.
	Color NullableString `json:"color,omitempty"`
	// Represents the italic status of the font.
	Italic NullableBool `json:"italic,omitempty"`
	// Font name (e.g. 'Calibri')
	Name NullableString `json:"name,omitempty"`
	// Font size.
	Size AnyOfnumberstringstring `json:"size,omitempty"`
	// Type of underline applied to the font. The possible values are: None, Single, Double, SingleAccountant, DoubleAccountant.
	Underline NullableString `json:"underline,omitempty"`
}

// NewWorkbookRangeFont instantiates a new WorkbookRangeFont object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkbookRangeFont() *WorkbookRangeFont {
	this := WorkbookRangeFont{}
	return &this
}

// NewWorkbookRangeFontWithDefaults instantiates a new WorkbookRangeFont object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkbookRangeFontWithDefaults() *WorkbookRangeFont {
	this := WorkbookRangeFont{}
	return &this
}

// GetBold returns the Bold field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkbookRangeFont) GetBold() bool {
	if o == nil || o.Bold.Get() == nil {
		var ret bool
		return ret
	}
	return *o.Bold.Get()
}

// GetBoldOk returns a tuple with the Bold field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkbookRangeFont) GetBoldOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Bold.Get(), o.Bold.IsSet()
}

// HasBold returns a boolean if a field has been set.
func (o *WorkbookRangeFont) HasBold() bool {
	if o != nil && o.Bold.IsSet() {
		return true
	}

	return false
}

// SetBold gets a reference to the given NullableBool and assigns it to the Bold field.
func (o *WorkbookRangeFont) SetBold(v bool) {
	o.Bold.Set(&v)
}
// SetBoldNil sets the value for Bold to be an explicit nil
func (o *WorkbookRangeFont) SetBoldNil() {
	o.Bold.Set(nil)
}

// UnsetBold ensures that no value is present for Bold, not even an explicit nil
func (o *WorkbookRangeFont) UnsetBold() {
	o.Bold.Unset()
}

// GetColor returns the Color field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkbookRangeFont) GetColor() string {
	if o == nil || o.Color.Get() == nil {
		var ret string
		return ret
	}
	return *o.Color.Get()
}

// GetColorOk returns a tuple with the Color field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkbookRangeFont) GetColorOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Color.Get(), o.Color.IsSet()
}

// HasColor returns a boolean if a field has been set.
func (o *WorkbookRangeFont) HasColor() bool {
	if o != nil && o.Color.IsSet() {
		return true
	}

	return false
}

// SetColor gets a reference to the given NullableString and assigns it to the Color field.
func (o *WorkbookRangeFont) SetColor(v string) {
	o.Color.Set(&v)
}
// SetColorNil sets the value for Color to be an explicit nil
func (o *WorkbookRangeFont) SetColorNil() {
	o.Color.Set(nil)
}

// UnsetColor ensures that no value is present for Color, not even an explicit nil
func (o *WorkbookRangeFont) UnsetColor() {
	o.Color.Unset()
}

// GetItalic returns the Italic field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkbookRangeFont) GetItalic() bool {
	if o == nil || o.Italic.Get() == nil {
		var ret bool
		return ret
	}
	return *o.Italic.Get()
}

// GetItalicOk returns a tuple with the Italic field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkbookRangeFont) GetItalicOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Italic.Get(), o.Italic.IsSet()
}

// HasItalic returns a boolean if a field has been set.
func (o *WorkbookRangeFont) HasItalic() bool {
	if o != nil && o.Italic.IsSet() {
		return true
	}

	return false
}

// SetItalic gets a reference to the given NullableBool and assigns it to the Italic field.
func (o *WorkbookRangeFont) SetItalic(v bool) {
	o.Italic.Set(&v)
}
// SetItalicNil sets the value for Italic to be an explicit nil
func (o *WorkbookRangeFont) SetItalicNil() {
	o.Italic.Set(nil)
}

// UnsetItalic ensures that no value is present for Italic, not even an explicit nil
func (o *WorkbookRangeFont) UnsetItalic() {
	o.Italic.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkbookRangeFont) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkbookRangeFont) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *WorkbookRangeFont) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *WorkbookRangeFont) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *WorkbookRangeFont) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *WorkbookRangeFont) UnsetName() {
	o.Name.Unset()
}

// GetSize returns the Size field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkbookRangeFont) GetSize() AnyOfnumberstringstring {
	if o == nil  {
		var ret AnyOfnumberstringstring
		return ret
	}
	return o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkbookRangeFont) GetSizeOk() (*AnyOfnumberstringstring, bool) {
	if o == nil || o.Size == nil {
		return nil, false
	}
	return &o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *WorkbookRangeFont) HasSize() bool {
	if o != nil && o.Size != nil {
		return true
	}

	return false
}

// SetSize gets a reference to the given AnyOfnumberstringstring and assigns it to the Size field.
func (o *WorkbookRangeFont) SetSize(v AnyOfnumberstringstring) {
	o.Size = v
}

// GetUnderline returns the Underline field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkbookRangeFont) GetUnderline() string {
	if o == nil || o.Underline.Get() == nil {
		var ret string
		return ret
	}
	return *o.Underline.Get()
}

// GetUnderlineOk returns a tuple with the Underline field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkbookRangeFont) GetUnderlineOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Underline.Get(), o.Underline.IsSet()
}

// HasUnderline returns a boolean if a field has been set.
func (o *WorkbookRangeFont) HasUnderline() bool {
	if o != nil && o.Underline.IsSet() {
		return true
	}

	return false
}

// SetUnderline gets a reference to the given NullableString and assigns it to the Underline field.
func (o *WorkbookRangeFont) SetUnderline(v string) {
	o.Underline.Set(&v)
}
// SetUnderlineNil sets the value for Underline to be an explicit nil
func (o *WorkbookRangeFont) SetUnderlineNil() {
	o.Underline.Set(nil)
}

// UnsetUnderline ensures that no value is present for Underline, not even an explicit nil
func (o *WorkbookRangeFont) UnsetUnderline() {
	o.Underline.Unset()
}

func (o WorkbookRangeFont) MarshalJSON() ([]byte, error) {
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

type NullableWorkbookRangeFont struct {
	value *WorkbookRangeFont
	isSet bool
}

func (v NullableWorkbookRangeFont) Get() *WorkbookRangeFont {
	return v.value
}

func (v *NullableWorkbookRangeFont) Set(val *WorkbookRangeFont) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkbookRangeFont) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkbookRangeFont) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkbookRangeFont(val *WorkbookRangeFont) *NullableWorkbookRangeFont {
	return &NullableWorkbookRangeFont{value: val, isSet: true}
}

func (v NullableWorkbookRangeFont) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkbookRangeFont) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


