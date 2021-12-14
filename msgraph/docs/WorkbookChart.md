# WorkbookChart

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Height** | Pointer to [**AnyOfnumberstringstring**](anyOf&lt;number,string,string&gt;.md) | Represents the height, in points, of the chart object. | [optional] 
**Left** | Pointer to [**AnyOfnumberstringstring**](anyOf&lt;number,string,string&gt;.md) | The distance, in points, from the left side of the chart to the worksheet origin. | [optional] 
**Name** | Pointer to **NullableString** | Represents the name of a chart object. | [optional] 
**Top** | Pointer to [**AnyOfnumberstringstring**](anyOf&lt;number,string,string&gt;.md) | Represents the distance, in points, from the top edge of the object to the top of row 1 (on a worksheet) or the top of the chart area (on a chart). | [optional] 
**Width** | Pointer to [**AnyOfnumberstringstring**](anyOf&lt;number,string,string&gt;.md) | Represents the width, in points, of the chart object. | [optional] 
**Axes** | Pointer to [**AnyOfmicrosoftGraphWorkbookChartAxes**](anyOf&lt;microsoft.graph.workbookChartAxes&gt;.md) | Represents chart axes. Read-only. | [optional] 
**DataLabels** | Pointer to [**AnyOfmicrosoftGraphWorkbookChartDataLabels**](anyOf&lt;microsoft.graph.workbookChartDataLabels&gt;.md) | Represents the datalabels on the chart. Read-only. | [optional] 
**Format** | Pointer to [**AnyOfmicrosoftGraphWorkbookChartAreaFormat**](anyOf&lt;microsoft.graph.workbookChartAreaFormat&gt;.md) | Encapsulates the format properties for the chart area. Read-only. | [optional] 
**Legend** | Pointer to [**AnyOfmicrosoftGraphWorkbookChartLegend**](anyOf&lt;microsoft.graph.workbookChartLegend&gt;.md) | Represents the legend for the chart. Read-only. | [optional] 
**Series** | Pointer to [**[]MicrosoftGraphWorkbookChartSeries**](MicrosoftGraphWorkbookChartSeries.md) | Represents either a single series or collection of series in the chart. Read-only. | [optional] 
**Title** | Pointer to [**AnyOfmicrosoftGraphWorkbookChartTitle**](anyOf&lt;microsoft.graph.workbookChartTitle&gt;.md) | Represents the title of the specified chart, including the text, visibility, position and formating of the title. Read-only. | [optional] 
**Worksheet** | Pointer to [**AnyOfmicrosoftGraphWorkbookWorksheet**](anyOf&lt;microsoft.graph.workbookWorksheet&gt;.md) | The worksheet containing the current chart. Read-only. | [optional] 

## Methods

### NewWorkbookChart

`func NewWorkbookChart() *WorkbookChart`

NewWorkbookChart instantiates a new WorkbookChart object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkbookChartWithDefaults

`func NewWorkbookChartWithDefaults() *WorkbookChart`

NewWorkbookChartWithDefaults instantiates a new WorkbookChart object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHeight

`func (o *WorkbookChart) GetHeight() AnyOfnumberstringstring`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *WorkbookChart) GetHeightOk() (*AnyOfnumberstringstring, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *WorkbookChart) SetHeight(v AnyOfnumberstringstring)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *WorkbookChart) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### SetHeightNil

`func (o *WorkbookChart) SetHeightNil(b bool)`

 SetHeightNil sets the value for Height to be an explicit nil

### UnsetHeight
`func (o *WorkbookChart) UnsetHeight()`

UnsetHeight ensures that no value is present for Height, not even an explicit nil
### GetLeft

`func (o *WorkbookChart) GetLeft() AnyOfnumberstringstring`

GetLeft returns the Left field if non-nil, zero value otherwise.

### GetLeftOk

`func (o *WorkbookChart) GetLeftOk() (*AnyOfnumberstringstring, bool)`

GetLeftOk returns a tuple with the Left field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeft

`func (o *WorkbookChart) SetLeft(v AnyOfnumberstringstring)`

SetLeft sets Left field to given value.

### HasLeft

`func (o *WorkbookChart) HasLeft() bool`

HasLeft returns a boolean if a field has been set.

### SetLeftNil

`func (o *WorkbookChart) SetLeftNil(b bool)`

 SetLeftNil sets the value for Left to be an explicit nil

### UnsetLeft
`func (o *WorkbookChart) UnsetLeft()`

UnsetLeft ensures that no value is present for Left, not even an explicit nil
### GetName

`func (o *WorkbookChart) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkbookChart) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkbookChart) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WorkbookChart) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *WorkbookChart) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *WorkbookChart) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetTop

`func (o *WorkbookChart) GetTop() AnyOfnumberstringstring`

GetTop returns the Top field if non-nil, zero value otherwise.

### GetTopOk

`func (o *WorkbookChart) GetTopOk() (*AnyOfnumberstringstring, bool)`

GetTopOk returns a tuple with the Top field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTop

`func (o *WorkbookChart) SetTop(v AnyOfnumberstringstring)`

SetTop sets Top field to given value.

### HasTop

`func (o *WorkbookChart) HasTop() bool`

HasTop returns a boolean if a field has been set.

### SetTopNil

`func (o *WorkbookChart) SetTopNil(b bool)`

 SetTopNil sets the value for Top to be an explicit nil

### UnsetTop
`func (o *WorkbookChart) UnsetTop()`

UnsetTop ensures that no value is present for Top, not even an explicit nil
### GetWidth

`func (o *WorkbookChart) GetWidth() AnyOfnumberstringstring`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *WorkbookChart) GetWidthOk() (*AnyOfnumberstringstring, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *WorkbookChart) SetWidth(v AnyOfnumberstringstring)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *WorkbookChart) HasWidth() bool`

HasWidth returns a boolean if a field has been set.

### SetWidthNil

`func (o *WorkbookChart) SetWidthNil(b bool)`

 SetWidthNil sets the value for Width to be an explicit nil

### UnsetWidth
`func (o *WorkbookChart) UnsetWidth()`

UnsetWidth ensures that no value is present for Width, not even an explicit nil
### GetAxes

`func (o *WorkbookChart) GetAxes() AnyOfmicrosoftGraphWorkbookChartAxes`

GetAxes returns the Axes field if non-nil, zero value otherwise.

### GetAxesOk

`func (o *WorkbookChart) GetAxesOk() (*AnyOfmicrosoftGraphWorkbookChartAxes, bool)`

GetAxesOk returns a tuple with the Axes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAxes

`func (o *WorkbookChart) SetAxes(v AnyOfmicrosoftGraphWorkbookChartAxes)`

SetAxes sets Axes field to given value.

### HasAxes

`func (o *WorkbookChart) HasAxes() bool`

HasAxes returns a boolean if a field has been set.

### SetAxesNil

`func (o *WorkbookChart) SetAxesNil(b bool)`

 SetAxesNil sets the value for Axes to be an explicit nil

### UnsetAxes
`func (o *WorkbookChart) UnsetAxes()`

UnsetAxes ensures that no value is present for Axes, not even an explicit nil
### GetDataLabels

`func (o *WorkbookChart) GetDataLabels() AnyOfmicrosoftGraphWorkbookChartDataLabels`

GetDataLabels returns the DataLabels field if non-nil, zero value otherwise.

### GetDataLabelsOk

`func (o *WorkbookChart) GetDataLabelsOk() (*AnyOfmicrosoftGraphWorkbookChartDataLabels, bool)`

GetDataLabelsOk returns a tuple with the DataLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataLabels

`func (o *WorkbookChart) SetDataLabels(v AnyOfmicrosoftGraphWorkbookChartDataLabels)`

SetDataLabels sets DataLabels field to given value.

### HasDataLabels

`func (o *WorkbookChart) HasDataLabels() bool`

HasDataLabels returns a boolean if a field has been set.

### SetDataLabelsNil

`func (o *WorkbookChart) SetDataLabelsNil(b bool)`

 SetDataLabelsNil sets the value for DataLabels to be an explicit nil

### UnsetDataLabels
`func (o *WorkbookChart) UnsetDataLabels()`

UnsetDataLabels ensures that no value is present for DataLabels, not even an explicit nil
### GetFormat

`func (o *WorkbookChart) GetFormat() AnyOfmicrosoftGraphWorkbookChartAreaFormat`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *WorkbookChart) GetFormatOk() (*AnyOfmicrosoftGraphWorkbookChartAreaFormat, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *WorkbookChart) SetFormat(v AnyOfmicrosoftGraphWorkbookChartAreaFormat)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *WorkbookChart) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### SetFormatNil

`func (o *WorkbookChart) SetFormatNil(b bool)`

 SetFormatNil sets the value for Format to be an explicit nil

### UnsetFormat
`func (o *WorkbookChart) UnsetFormat()`

UnsetFormat ensures that no value is present for Format, not even an explicit nil
### GetLegend

`func (o *WorkbookChart) GetLegend() AnyOfmicrosoftGraphWorkbookChartLegend`

GetLegend returns the Legend field if non-nil, zero value otherwise.

### GetLegendOk

`func (o *WorkbookChart) GetLegendOk() (*AnyOfmicrosoftGraphWorkbookChartLegend, bool)`

GetLegendOk returns a tuple with the Legend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegend

`func (o *WorkbookChart) SetLegend(v AnyOfmicrosoftGraphWorkbookChartLegend)`

SetLegend sets Legend field to given value.

### HasLegend

`func (o *WorkbookChart) HasLegend() bool`

HasLegend returns a boolean if a field has been set.

### SetLegendNil

`func (o *WorkbookChart) SetLegendNil(b bool)`

 SetLegendNil sets the value for Legend to be an explicit nil

### UnsetLegend
`func (o *WorkbookChart) UnsetLegend()`

UnsetLegend ensures that no value is present for Legend, not even an explicit nil
### GetSeries

`func (o *WorkbookChart) GetSeries() []MicrosoftGraphWorkbookChartSeries`

GetSeries returns the Series field if non-nil, zero value otherwise.

### GetSeriesOk

`func (o *WorkbookChart) GetSeriesOk() (*[]MicrosoftGraphWorkbookChartSeries, bool)`

GetSeriesOk returns a tuple with the Series field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeries

`func (o *WorkbookChart) SetSeries(v []MicrosoftGraphWorkbookChartSeries)`

SetSeries sets Series field to given value.

### HasSeries

`func (o *WorkbookChart) HasSeries() bool`

HasSeries returns a boolean if a field has been set.

### GetTitle

`func (o *WorkbookChart) GetTitle() AnyOfmicrosoftGraphWorkbookChartTitle`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *WorkbookChart) GetTitleOk() (*AnyOfmicrosoftGraphWorkbookChartTitle, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *WorkbookChart) SetTitle(v AnyOfmicrosoftGraphWorkbookChartTitle)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *WorkbookChart) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *WorkbookChart) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *WorkbookChart) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetWorksheet

`func (o *WorkbookChart) GetWorksheet() AnyOfmicrosoftGraphWorkbookWorksheet`

GetWorksheet returns the Worksheet field if non-nil, zero value otherwise.

### GetWorksheetOk

`func (o *WorkbookChart) GetWorksheetOk() (*AnyOfmicrosoftGraphWorkbookWorksheet, bool)`

GetWorksheetOk returns a tuple with the Worksheet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorksheet

`func (o *WorkbookChart) SetWorksheet(v AnyOfmicrosoftGraphWorkbookWorksheet)`

SetWorksheet sets Worksheet field to given value.

### HasWorksheet

`func (o *WorkbookChart) HasWorksheet() bool`

HasWorksheet returns a boolean if a field has been set.

### SetWorksheetNil

`func (o *WorkbookChart) SetWorksheetNil(b bool)`

 SetWorksheetNil sets the value for Worksheet to be an explicit nil

### UnsetWorksheet
`func (o *WorkbookChart) UnsetWorksheet()`

UnsetWorksheet ensures that no value is present for Worksheet, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


