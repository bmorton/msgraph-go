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

// WorkbookChartSeries struct for WorkbookChartSeries
type WorkbookChartSeries struct {
	// Represents the name of a series in a chart.
	Name NullableString `json:"name,omitempty"`
	// Represents the formatting of a chart series, which includes fill and line formatting. Read-only.
	Format AnyOfmicrosoftGraphWorkbookChartSeriesFormat `json:"format,omitempty"`
	// Represents a collection of all points in the series. Read-only.
	Points *[]MicrosoftGraphWorkbookChartPoint `json:"points,omitempty"`
}

// NewWorkbookChartSeries instantiates a new WorkbookChartSeries object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkbookChartSeries() *WorkbookChartSeries {
	this := WorkbookChartSeries{}
	return &this
}

// NewWorkbookChartSeriesWithDefaults instantiates a new WorkbookChartSeries object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkbookChartSeriesWithDefaults() *WorkbookChartSeries {
	this := WorkbookChartSeries{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkbookChartSeries) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkbookChartSeries) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *WorkbookChartSeries) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *WorkbookChartSeries) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *WorkbookChartSeries) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *WorkbookChartSeries) UnsetName() {
	o.Name.Unset()
}

// GetFormat returns the Format field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkbookChartSeries) GetFormat() AnyOfmicrosoftGraphWorkbookChartSeriesFormat {
	if o == nil  {
		var ret AnyOfmicrosoftGraphWorkbookChartSeriesFormat
		return ret
	}
	return o.Format
}

// GetFormatOk returns a tuple with the Format field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkbookChartSeries) GetFormatOk() (*AnyOfmicrosoftGraphWorkbookChartSeriesFormat, bool) {
	if o == nil || o.Format == nil {
		return nil, false
	}
	return &o.Format, true
}

// HasFormat returns a boolean if a field has been set.
func (o *WorkbookChartSeries) HasFormat() bool {
	if o != nil && o.Format != nil {
		return true
	}

	return false
}

// SetFormat gets a reference to the given AnyOfmicrosoftGraphWorkbookChartSeriesFormat and assigns it to the Format field.
func (o *WorkbookChartSeries) SetFormat(v AnyOfmicrosoftGraphWorkbookChartSeriesFormat) {
	o.Format = v
}

// GetPoints returns the Points field value if set, zero value otherwise.
func (o *WorkbookChartSeries) GetPoints() []MicrosoftGraphWorkbookChartPoint {
	if o == nil || o.Points == nil {
		var ret []MicrosoftGraphWorkbookChartPoint
		return ret
	}
	return *o.Points
}

// GetPointsOk returns a tuple with the Points field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkbookChartSeries) GetPointsOk() (*[]MicrosoftGraphWorkbookChartPoint, bool) {
	if o == nil || o.Points == nil {
		return nil, false
	}
	return o.Points, true
}

// HasPoints returns a boolean if a field has been set.
func (o *WorkbookChartSeries) HasPoints() bool {
	if o != nil && o.Points != nil {
		return true
	}

	return false
}

// SetPoints gets a reference to the given []MicrosoftGraphWorkbookChartPoint and assigns it to the Points field.
func (o *WorkbookChartSeries) SetPoints(v []MicrosoftGraphWorkbookChartPoint) {
	o.Points = &v
}

func (o WorkbookChartSeries) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Format != nil {
		toSerialize["format"] = o.Format
	}
	if o.Points != nil {
		toSerialize["points"] = o.Points
	}
	return json.Marshal(toSerialize)
}

type NullableWorkbookChartSeries struct {
	value *WorkbookChartSeries
	isSet bool
}

func (v NullableWorkbookChartSeries) Get() *WorkbookChartSeries {
	return v.value
}

func (v *NullableWorkbookChartSeries) Set(val *WorkbookChartSeries) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkbookChartSeries) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkbookChartSeries) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkbookChartSeries(val *WorkbookChartSeries) *NullableWorkbookChartSeries {
	return &NullableWorkbookChartSeries{value: val, isSet: true}
}

func (v NullableWorkbookChartSeries) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkbookChartSeries) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


