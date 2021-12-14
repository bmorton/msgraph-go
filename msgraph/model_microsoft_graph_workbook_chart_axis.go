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

// MicrosoftGraphWorkbookChartAxis struct for MicrosoftGraphWorkbookChartAxis
type MicrosoftGraphWorkbookChartAxis struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
	// Represents the interval between two major tick marks. Can be set to a numeric value or an empty string.  The returned value is always a number.
	MajorUnit AnyOfobject `json:"majorUnit,omitempty"`
	// Represents the maximum value on the value axis.  Can be set to a numeric value or an empty string (for automatic axis values).  The returned value is always a number.
	Maximum AnyOfobject `json:"maximum,omitempty"`
	// Represents the minimum value on the value axis. Can be set to a numeric value or an empty string (for automatic axis values).  The returned value is always a number.
	Minimum AnyOfobject `json:"minimum,omitempty"`
	// Represents the interval between two minor tick marks. 'Can be set to a numeric value or an empty string (for automatic axis values). The returned value is always a number.
	MinorUnit AnyOfobject `json:"minorUnit,omitempty"`
	// Represents the formatting of a chart object, which includes line and font formatting. Read-only.
	Format AnyOfmicrosoftGraphWorkbookChartAxisFormat `json:"format,omitempty"`
	// Returns a gridlines object that represents the major gridlines for the specified axis. Read-only.
	MajorGridlines AnyOfmicrosoftGraphWorkbookChartGridlines `json:"majorGridlines,omitempty"`
	// Returns a Gridlines object that represents the minor gridlines for the specified axis. Read-only.
	MinorGridlines AnyOfmicrosoftGraphWorkbookChartGridlines `json:"minorGridlines,omitempty"`
	// Represents the axis title. Read-only.
	Title AnyOfmicrosoftGraphWorkbookChartAxisTitle `json:"title,omitempty"`
}

// NewMicrosoftGraphWorkbookChartAxis instantiates a new MicrosoftGraphWorkbookChartAxis object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphWorkbookChartAxis() *MicrosoftGraphWorkbookChartAxis {
	this := MicrosoftGraphWorkbookChartAxis{}
	return &this
}

// NewMicrosoftGraphWorkbookChartAxisWithDefaults instantiates a new MicrosoftGraphWorkbookChartAxis object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphWorkbookChartAxisWithDefaults() *MicrosoftGraphWorkbookChartAxis {
	this := MicrosoftGraphWorkbookChartAxis{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MicrosoftGraphWorkbookChartAxis) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphWorkbookChartAxis) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphWorkbookChartAxis) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MicrosoftGraphWorkbookChartAxis) SetId(v string) {
	o.Id = &v
}

// GetMajorUnit returns the MajorUnit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphWorkbookChartAxis) GetMajorUnit() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.MajorUnit
}

// GetMajorUnitOk returns a tuple with the MajorUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphWorkbookChartAxis) GetMajorUnitOk() (*AnyOfobject, bool) {
	if o == nil || o.MajorUnit == nil {
		return nil, false
	}
	return &o.MajorUnit, true
}

// HasMajorUnit returns a boolean if a field has been set.
func (o *MicrosoftGraphWorkbookChartAxis) HasMajorUnit() bool {
	if o != nil && o.MajorUnit != nil {
		return true
	}

	return false
}

// SetMajorUnit gets a reference to the given AnyOfobject and assigns it to the MajorUnit field.
func (o *MicrosoftGraphWorkbookChartAxis) SetMajorUnit(v AnyOfobject) {
	o.MajorUnit = v
}

// GetMaximum returns the Maximum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphWorkbookChartAxis) GetMaximum() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Maximum
}

// GetMaximumOk returns a tuple with the Maximum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphWorkbookChartAxis) GetMaximumOk() (*AnyOfobject, bool) {
	if o == nil || o.Maximum == nil {
		return nil, false
	}
	return &o.Maximum, true
}

// HasMaximum returns a boolean if a field has been set.
func (o *MicrosoftGraphWorkbookChartAxis) HasMaximum() bool {
	if o != nil && o.Maximum != nil {
		return true
	}

	return false
}

// SetMaximum gets a reference to the given AnyOfobject and assigns it to the Maximum field.
func (o *MicrosoftGraphWorkbookChartAxis) SetMaximum(v AnyOfobject) {
	o.Maximum = v
}

// GetMinimum returns the Minimum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphWorkbookChartAxis) GetMinimum() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Minimum
}

// GetMinimumOk returns a tuple with the Minimum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphWorkbookChartAxis) GetMinimumOk() (*AnyOfobject, bool) {
	if o == nil || o.Minimum == nil {
		return nil, false
	}
	return &o.Minimum, true
}

// HasMinimum returns a boolean if a field has been set.
func (o *MicrosoftGraphWorkbookChartAxis) HasMinimum() bool {
	if o != nil && o.Minimum != nil {
		return true
	}

	return false
}

// SetMinimum gets a reference to the given AnyOfobject and assigns it to the Minimum field.
func (o *MicrosoftGraphWorkbookChartAxis) SetMinimum(v AnyOfobject) {
	o.Minimum = v
}

// GetMinorUnit returns the MinorUnit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphWorkbookChartAxis) GetMinorUnit() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.MinorUnit
}

// GetMinorUnitOk returns a tuple with the MinorUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphWorkbookChartAxis) GetMinorUnitOk() (*AnyOfobject, bool) {
	if o == nil || o.MinorUnit == nil {
		return nil, false
	}
	return &o.MinorUnit, true
}

// HasMinorUnit returns a boolean if a field has been set.
func (o *MicrosoftGraphWorkbookChartAxis) HasMinorUnit() bool {
	if o != nil && o.MinorUnit != nil {
		return true
	}

	return false
}

// SetMinorUnit gets a reference to the given AnyOfobject and assigns it to the MinorUnit field.
func (o *MicrosoftGraphWorkbookChartAxis) SetMinorUnit(v AnyOfobject) {
	o.MinorUnit = v
}

// GetFormat returns the Format field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphWorkbookChartAxis) GetFormat() AnyOfmicrosoftGraphWorkbookChartAxisFormat {
	if o == nil  {
		var ret AnyOfmicrosoftGraphWorkbookChartAxisFormat
		return ret
	}
	return o.Format
}

// GetFormatOk returns a tuple with the Format field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphWorkbookChartAxis) GetFormatOk() (*AnyOfmicrosoftGraphWorkbookChartAxisFormat, bool) {
	if o == nil || o.Format == nil {
		return nil, false
	}
	return &o.Format, true
}

// HasFormat returns a boolean if a field has been set.
func (o *MicrosoftGraphWorkbookChartAxis) HasFormat() bool {
	if o != nil && o.Format != nil {
		return true
	}

	return false
}

// SetFormat gets a reference to the given AnyOfmicrosoftGraphWorkbookChartAxisFormat and assigns it to the Format field.
func (o *MicrosoftGraphWorkbookChartAxis) SetFormat(v AnyOfmicrosoftGraphWorkbookChartAxisFormat) {
	o.Format = v
}

// GetMajorGridlines returns the MajorGridlines field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphWorkbookChartAxis) GetMajorGridlines() AnyOfmicrosoftGraphWorkbookChartGridlines {
	if o == nil  {
		var ret AnyOfmicrosoftGraphWorkbookChartGridlines
		return ret
	}
	return o.MajorGridlines
}

// GetMajorGridlinesOk returns a tuple with the MajorGridlines field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphWorkbookChartAxis) GetMajorGridlinesOk() (*AnyOfmicrosoftGraphWorkbookChartGridlines, bool) {
	if o == nil || o.MajorGridlines == nil {
		return nil, false
	}
	return &o.MajorGridlines, true
}

// HasMajorGridlines returns a boolean if a field has been set.
func (o *MicrosoftGraphWorkbookChartAxis) HasMajorGridlines() bool {
	if o != nil && o.MajorGridlines != nil {
		return true
	}

	return false
}

// SetMajorGridlines gets a reference to the given AnyOfmicrosoftGraphWorkbookChartGridlines and assigns it to the MajorGridlines field.
func (o *MicrosoftGraphWorkbookChartAxis) SetMajorGridlines(v AnyOfmicrosoftGraphWorkbookChartGridlines) {
	o.MajorGridlines = v
}

// GetMinorGridlines returns the MinorGridlines field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphWorkbookChartAxis) GetMinorGridlines() AnyOfmicrosoftGraphWorkbookChartGridlines {
	if o == nil  {
		var ret AnyOfmicrosoftGraphWorkbookChartGridlines
		return ret
	}
	return o.MinorGridlines
}

// GetMinorGridlinesOk returns a tuple with the MinorGridlines field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphWorkbookChartAxis) GetMinorGridlinesOk() (*AnyOfmicrosoftGraphWorkbookChartGridlines, bool) {
	if o == nil || o.MinorGridlines == nil {
		return nil, false
	}
	return &o.MinorGridlines, true
}

// HasMinorGridlines returns a boolean if a field has been set.
func (o *MicrosoftGraphWorkbookChartAxis) HasMinorGridlines() bool {
	if o != nil && o.MinorGridlines != nil {
		return true
	}

	return false
}

// SetMinorGridlines gets a reference to the given AnyOfmicrosoftGraphWorkbookChartGridlines and assigns it to the MinorGridlines field.
func (o *MicrosoftGraphWorkbookChartAxis) SetMinorGridlines(v AnyOfmicrosoftGraphWorkbookChartGridlines) {
	o.MinorGridlines = v
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphWorkbookChartAxis) GetTitle() AnyOfmicrosoftGraphWorkbookChartAxisTitle {
	if o == nil  {
		var ret AnyOfmicrosoftGraphWorkbookChartAxisTitle
		return ret
	}
	return o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphWorkbookChartAxis) GetTitleOk() (*AnyOfmicrosoftGraphWorkbookChartAxisTitle, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return &o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *MicrosoftGraphWorkbookChartAxis) HasTitle() bool {
	if o != nil && o.Title != nil {
		return true
	}

	return false
}

// SetTitle gets a reference to the given AnyOfmicrosoftGraphWorkbookChartAxisTitle and assigns it to the Title field.
func (o *MicrosoftGraphWorkbookChartAxis) SetTitle(v AnyOfmicrosoftGraphWorkbookChartAxisTitle) {
	o.Title = v
}

func (o MicrosoftGraphWorkbookChartAxis) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.MajorUnit != nil {
		toSerialize["majorUnit"] = o.MajorUnit
	}
	if o.Maximum != nil {
		toSerialize["maximum"] = o.Maximum
	}
	if o.Minimum != nil {
		toSerialize["minimum"] = o.Minimum
	}
	if o.MinorUnit != nil {
		toSerialize["minorUnit"] = o.MinorUnit
	}
	if o.Format != nil {
		toSerialize["format"] = o.Format
	}
	if o.MajorGridlines != nil {
		toSerialize["majorGridlines"] = o.MajorGridlines
	}
	if o.MinorGridlines != nil {
		toSerialize["minorGridlines"] = o.MinorGridlines
	}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphWorkbookChartAxis struct {
	value *MicrosoftGraphWorkbookChartAxis
	isSet bool
}

func (v NullableMicrosoftGraphWorkbookChartAxis) Get() *MicrosoftGraphWorkbookChartAxis {
	return v.value
}

func (v *NullableMicrosoftGraphWorkbookChartAxis) Set(val *MicrosoftGraphWorkbookChartAxis) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphWorkbookChartAxis) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphWorkbookChartAxis) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphWorkbookChartAxis(val *MicrosoftGraphWorkbookChartAxis) *NullableMicrosoftGraphWorkbookChartAxis {
	return &NullableMicrosoftGraphWorkbookChartAxis{value: val, isSet: true}
}

func (v NullableMicrosoftGraphWorkbookChartAxis) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphWorkbookChartAxis) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


