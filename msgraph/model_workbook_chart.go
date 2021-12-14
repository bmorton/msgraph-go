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

// WorkbookChart struct for WorkbookChart
type WorkbookChart struct {
	// Represents the height, in points, of the chart object.
	Height AnyOfnumberstringstring `json:"height,omitempty"`
	// The distance, in points, from the left side of the chart to the worksheet origin.
	Left AnyOfnumberstringstring `json:"left,omitempty"`
	// Represents the name of a chart object.
	Name NullableString `json:"name,omitempty"`
	// Represents the distance, in points, from the top edge of the object to the top of row 1 (on a worksheet) or the top of the chart area (on a chart).
	Top AnyOfnumberstringstring `json:"top,omitempty"`
	// Represents the width, in points, of the chart object.
	Width AnyOfnumberstringstring `json:"width,omitempty"`
	// Represents chart axes. Read-only.
	Axes AnyOfmicrosoftGraphWorkbookChartAxes `json:"axes,omitempty"`
	// Represents the datalabels on the chart. Read-only.
	DataLabels AnyOfmicrosoftGraphWorkbookChartDataLabels `json:"dataLabels,omitempty"`
	// Encapsulates the format properties for the chart area. Read-only.
	Format AnyOfmicrosoftGraphWorkbookChartAreaFormat `json:"format,omitempty"`
	// Represents the legend for the chart. Read-only.
	Legend AnyOfmicrosoftGraphWorkbookChartLegend `json:"legend,omitempty"`
	// Represents either a single series or collection of series in the chart. Read-only.
	Series *[]MicrosoftGraphWorkbookChartSeries `json:"series,omitempty"`
	// Represents the title of the specified chart, including the text, visibility, position and formating of the title. Read-only.
	Title AnyOfmicrosoftGraphWorkbookChartTitle `json:"title,omitempty"`
	// The worksheet containing the current chart. Read-only.
	Worksheet AnyOfmicrosoftGraphWorkbookWorksheet `json:"worksheet,omitempty"`
}

// NewWorkbookChart instantiates a new WorkbookChart object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkbookChart() *WorkbookChart {
	this := WorkbookChart{}
	return &this
}

// NewWorkbookChartWithDefaults instantiates a new WorkbookChart object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkbookChartWithDefaults() *WorkbookChart {
	this := WorkbookChart{}
	return &this
}

// GetHeight returns the Height field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkbookChart) GetHeight() AnyOfnumberstringstring {
	if o == nil  {
		var ret AnyOfnumberstringstring
		return ret
	}
	return o.Height
}

// GetHeightOk returns a tuple with the Height field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkbookChart) GetHeightOk() (*AnyOfnumberstringstring, bool) {
	if o == nil || o.Height == nil {
		return nil, false
	}
	return &o.Height, true
}

// HasHeight returns a boolean if a field has been set.
func (o *WorkbookChart) HasHeight() bool {
	if o != nil && o.Height != nil {
		return true
	}

	return false
}

// SetHeight gets a reference to the given AnyOfnumberstringstring and assigns it to the Height field.
func (o *WorkbookChart) SetHeight(v AnyOfnumberstringstring) {
	o.Height = v
}

// GetLeft returns the Left field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkbookChart) GetLeft() AnyOfnumberstringstring {
	if o == nil  {
		var ret AnyOfnumberstringstring
		return ret
	}
	return o.Left
}

// GetLeftOk returns a tuple with the Left field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkbookChart) GetLeftOk() (*AnyOfnumberstringstring, bool) {
	if o == nil || o.Left == nil {
		return nil, false
	}
	return &o.Left, true
}

// HasLeft returns a boolean if a field has been set.
func (o *WorkbookChart) HasLeft() bool {
	if o != nil && o.Left != nil {
		return true
	}

	return false
}

// SetLeft gets a reference to the given AnyOfnumberstringstring and assigns it to the Left field.
func (o *WorkbookChart) SetLeft(v AnyOfnumberstringstring) {
	o.Left = v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkbookChart) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkbookChart) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *WorkbookChart) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *WorkbookChart) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *WorkbookChart) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *WorkbookChart) UnsetName() {
	o.Name.Unset()
}

// GetTop returns the Top field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkbookChart) GetTop() AnyOfnumberstringstring {
	if o == nil  {
		var ret AnyOfnumberstringstring
		return ret
	}
	return o.Top
}

// GetTopOk returns a tuple with the Top field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkbookChart) GetTopOk() (*AnyOfnumberstringstring, bool) {
	if o == nil || o.Top == nil {
		return nil, false
	}
	return &o.Top, true
}

// HasTop returns a boolean if a field has been set.
func (o *WorkbookChart) HasTop() bool {
	if o != nil && o.Top != nil {
		return true
	}

	return false
}

// SetTop gets a reference to the given AnyOfnumberstringstring and assigns it to the Top field.
func (o *WorkbookChart) SetTop(v AnyOfnumberstringstring) {
	o.Top = v
}

// GetWidth returns the Width field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkbookChart) GetWidth() AnyOfnumberstringstring {
	if o == nil  {
		var ret AnyOfnumberstringstring
		return ret
	}
	return o.Width
}

// GetWidthOk returns a tuple with the Width field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkbookChart) GetWidthOk() (*AnyOfnumberstringstring, bool) {
	if o == nil || o.Width == nil {
		return nil, false
	}
	return &o.Width, true
}

// HasWidth returns a boolean if a field has been set.
func (o *WorkbookChart) HasWidth() bool {
	if o != nil && o.Width != nil {
		return true
	}

	return false
}

// SetWidth gets a reference to the given AnyOfnumberstringstring and assigns it to the Width field.
func (o *WorkbookChart) SetWidth(v AnyOfnumberstringstring) {
	o.Width = v
}

// GetAxes returns the Axes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkbookChart) GetAxes() AnyOfmicrosoftGraphWorkbookChartAxes {
	if o == nil  {
		var ret AnyOfmicrosoftGraphWorkbookChartAxes
		return ret
	}
	return o.Axes
}

// GetAxesOk returns a tuple with the Axes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkbookChart) GetAxesOk() (*AnyOfmicrosoftGraphWorkbookChartAxes, bool) {
	if o == nil || o.Axes == nil {
		return nil, false
	}
	return &o.Axes, true
}

// HasAxes returns a boolean if a field has been set.
func (o *WorkbookChart) HasAxes() bool {
	if o != nil && o.Axes != nil {
		return true
	}

	return false
}

// SetAxes gets a reference to the given AnyOfmicrosoftGraphWorkbookChartAxes and assigns it to the Axes field.
func (o *WorkbookChart) SetAxes(v AnyOfmicrosoftGraphWorkbookChartAxes) {
	o.Axes = v
}

// GetDataLabels returns the DataLabels field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkbookChart) GetDataLabels() AnyOfmicrosoftGraphWorkbookChartDataLabels {
	if o == nil  {
		var ret AnyOfmicrosoftGraphWorkbookChartDataLabels
		return ret
	}
	return o.DataLabels
}

// GetDataLabelsOk returns a tuple with the DataLabels field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkbookChart) GetDataLabelsOk() (*AnyOfmicrosoftGraphWorkbookChartDataLabels, bool) {
	if o == nil || o.DataLabels == nil {
		return nil, false
	}
	return &o.DataLabels, true
}

// HasDataLabels returns a boolean if a field has been set.
func (o *WorkbookChart) HasDataLabels() bool {
	if o != nil && o.DataLabels != nil {
		return true
	}

	return false
}

// SetDataLabels gets a reference to the given AnyOfmicrosoftGraphWorkbookChartDataLabels and assigns it to the DataLabels field.
func (o *WorkbookChart) SetDataLabels(v AnyOfmicrosoftGraphWorkbookChartDataLabels) {
	o.DataLabels = v
}

// GetFormat returns the Format field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkbookChart) GetFormat() AnyOfmicrosoftGraphWorkbookChartAreaFormat {
	if o == nil  {
		var ret AnyOfmicrosoftGraphWorkbookChartAreaFormat
		return ret
	}
	return o.Format
}

// GetFormatOk returns a tuple with the Format field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkbookChart) GetFormatOk() (*AnyOfmicrosoftGraphWorkbookChartAreaFormat, bool) {
	if o == nil || o.Format == nil {
		return nil, false
	}
	return &o.Format, true
}

// HasFormat returns a boolean if a field has been set.
func (o *WorkbookChart) HasFormat() bool {
	if o != nil && o.Format != nil {
		return true
	}

	return false
}

// SetFormat gets a reference to the given AnyOfmicrosoftGraphWorkbookChartAreaFormat and assigns it to the Format field.
func (o *WorkbookChart) SetFormat(v AnyOfmicrosoftGraphWorkbookChartAreaFormat) {
	o.Format = v
}

// GetLegend returns the Legend field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkbookChart) GetLegend() AnyOfmicrosoftGraphWorkbookChartLegend {
	if o == nil  {
		var ret AnyOfmicrosoftGraphWorkbookChartLegend
		return ret
	}
	return o.Legend
}

// GetLegendOk returns a tuple with the Legend field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkbookChart) GetLegendOk() (*AnyOfmicrosoftGraphWorkbookChartLegend, bool) {
	if o == nil || o.Legend == nil {
		return nil, false
	}
	return &o.Legend, true
}

// HasLegend returns a boolean if a field has been set.
func (o *WorkbookChart) HasLegend() bool {
	if o != nil && o.Legend != nil {
		return true
	}

	return false
}

// SetLegend gets a reference to the given AnyOfmicrosoftGraphWorkbookChartLegend and assigns it to the Legend field.
func (o *WorkbookChart) SetLegend(v AnyOfmicrosoftGraphWorkbookChartLegend) {
	o.Legend = v
}

// GetSeries returns the Series field value if set, zero value otherwise.
func (o *WorkbookChart) GetSeries() []MicrosoftGraphWorkbookChartSeries {
	if o == nil || o.Series == nil {
		var ret []MicrosoftGraphWorkbookChartSeries
		return ret
	}
	return *o.Series
}

// GetSeriesOk returns a tuple with the Series field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkbookChart) GetSeriesOk() (*[]MicrosoftGraphWorkbookChartSeries, bool) {
	if o == nil || o.Series == nil {
		return nil, false
	}
	return o.Series, true
}

// HasSeries returns a boolean if a field has been set.
func (o *WorkbookChart) HasSeries() bool {
	if o != nil && o.Series != nil {
		return true
	}

	return false
}

// SetSeries gets a reference to the given []MicrosoftGraphWorkbookChartSeries and assigns it to the Series field.
func (o *WorkbookChart) SetSeries(v []MicrosoftGraphWorkbookChartSeries) {
	o.Series = &v
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkbookChart) GetTitle() AnyOfmicrosoftGraphWorkbookChartTitle {
	if o == nil  {
		var ret AnyOfmicrosoftGraphWorkbookChartTitle
		return ret
	}
	return o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkbookChart) GetTitleOk() (*AnyOfmicrosoftGraphWorkbookChartTitle, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return &o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *WorkbookChart) HasTitle() bool {
	if o != nil && o.Title != nil {
		return true
	}

	return false
}

// SetTitle gets a reference to the given AnyOfmicrosoftGraphWorkbookChartTitle and assigns it to the Title field.
func (o *WorkbookChart) SetTitle(v AnyOfmicrosoftGraphWorkbookChartTitle) {
	o.Title = v
}

// GetWorksheet returns the Worksheet field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkbookChart) GetWorksheet() AnyOfmicrosoftGraphWorkbookWorksheet {
	if o == nil  {
		var ret AnyOfmicrosoftGraphWorkbookWorksheet
		return ret
	}
	return o.Worksheet
}

// GetWorksheetOk returns a tuple with the Worksheet field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkbookChart) GetWorksheetOk() (*AnyOfmicrosoftGraphWorkbookWorksheet, bool) {
	if o == nil || o.Worksheet == nil {
		return nil, false
	}
	return &o.Worksheet, true
}

// HasWorksheet returns a boolean if a field has been set.
func (o *WorkbookChart) HasWorksheet() bool {
	if o != nil && o.Worksheet != nil {
		return true
	}

	return false
}

// SetWorksheet gets a reference to the given AnyOfmicrosoftGraphWorkbookWorksheet and assigns it to the Worksheet field.
func (o *WorkbookChart) SetWorksheet(v AnyOfmicrosoftGraphWorkbookWorksheet) {
	o.Worksheet = v
}

func (o WorkbookChart) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Height != nil {
		toSerialize["height"] = o.Height
	}
	if o.Left != nil {
		toSerialize["left"] = o.Left
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Top != nil {
		toSerialize["top"] = o.Top
	}
	if o.Width != nil {
		toSerialize["width"] = o.Width
	}
	if o.Axes != nil {
		toSerialize["axes"] = o.Axes
	}
	if o.DataLabels != nil {
		toSerialize["dataLabels"] = o.DataLabels
	}
	if o.Format != nil {
		toSerialize["format"] = o.Format
	}
	if o.Legend != nil {
		toSerialize["legend"] = o.Legend
	}
	if o.Series != nil {
		toSerialize["series"] = o.Series
	}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	if o.Worksheet != nil {
		toSerialize["worksheet"] = o.Worksheet
	}
	return json.Marshal(toSerialize)
}

type NullableWorkbookChart struct {
	value *WorkbookChart
	isSet bool
}

func (v NullableWorkbookChart) Get() *WorkbookChart {
	return v.value
}

func (v *NullableWorkbookChart) Set(val *WorkbookChart) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkbookChart) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkbookChart) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkbookChart(val *WorkbookChart) *NullableWorkbookChart {
	return &NullableWorkbookChart{value: val, isSet: true}
}

func (v NullableWorkbookChart) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkbookChart) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

