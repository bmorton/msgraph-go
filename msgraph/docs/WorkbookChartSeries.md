# WorkbookChartSeries

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | Represents the name of a series in a chart. | [optional] 
**Format** | Pointer to [**AnyOfmicrosoftGraphWorkbookChartSeriesFormat**](anyOf&lt;microsoft.graph.workbookChartSeriesFormat&gt;.md) | Represents the formatting of a chart series, which includes fill and line formatting. Read-only. | [optional] 
**Points** | Pointer to [**[]MicrosoftGraphWorkbookChartPoint**](MicrosoftGraphWorkbookChartPoint.md) | Represents a collection of all points in the series. Read-only. | [optional] 

## Methods

### NewWorkbookChartSeries

`func NewWorkbookChartSeries() *WorkbookChartSeries`

NewWorkbookChartSeries instantiates a new WorkbookChartSeries object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkbookChartSeriesWithDefaults

`func NewWorkbookChartSeriesWithDefaults() *WorkbookChartSeries`

NewWorkbookChartSeriesWithDefaults instantiates a new WorkbookChartSeries object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *WorkbookChartSeries) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkbookChartSeries) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkbookChartSeries) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WorkbookChartSeries) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *WorkbookChartSeries) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *WorkbookChartSeries) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetFormat

`func (o *WorkbookChartSeries) GetFormat() AnyOfmicrosoftGraphWorkbookChartSeriesFormat`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *WorkbookChartSeries) GetFormatOk() (*AnyOfmicrosoftGraphWorkbookChartSeriesFormat, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *WorkbookChartSeries) SetFormat(v AnyOfmicrosoftGraphWorkbookChartSeriesFormat)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *WorkbookChartSeries) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### SetFormatNil

`func (o *WorkbookChartSeries) SetFormatNil(b bool)`

 SetFormatNil sets the value for Format to be an explicit nil

### UnsetFormat
`func (o *WorkbookChartSeries) UnsetFormat()`

UnsetFormat ensures that no value is present for Format, not even an explicit nil
### GetPoints

`func (o *WorkbookChartSeries) GetPoints() []MicrosoftGraphWorkbookChartPoint`

GetPoints returns the Points field if non-nil, zero value otherwise.

### GetPointsOk

`func (o *WorkbookChartSeries) GetPointsOk() (*[]MicrosoftGraphWorkbookChartPoint, bool)`

GetPointsOk returns a tuple with the Points field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoints

`func (o *WorkbookChartSeries) SetPoints(v []MicrosoftGraphWorkbookChartPoint)`

SetPoints sets Points field to given value.

### HasPoints

`func (o *WorkbookChartSeries) HasPoints() bool`

HasPoints returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


