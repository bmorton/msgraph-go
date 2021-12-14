# MicrosoftGraphWorkbookChartSeries

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**Name** | Pointer to **NullableString** | Represents the name of a series in a chart. | [optional] 
**Format** | Pointer to [**AnyOfmicrosoftGraphWorkbookChartSeriesFormat**](anyOf&lt;microsoft.graph.workbookChartSeriesFormat&gt;.md) | Represents the formatting of a chart series, which includes fill and line formatting. Read-only. | [optional] 
**Points** | Pointer to [**[]MicrosoftGraphWorkbookChartPoint**](MicrosoftGraphWorkbookChartPoint.md) | Represents a collection of all points in the series. Read-only. | [optional] 

## Methods

### NewMicrosoftGraphWorkbookChartSeries

`func NewMicrosoftGraphWorkbookChartSeries() *MicrosoftGraphWorkbookChartSeries`

NewMicrosoftGraphWorkbookChartSeries instantiates a new MicrosoftGraphWorkbookChartSeries object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphWorkbookChartSeriesWithDefaults

`func NewMicrosoftGraphWorkbookChartSeriesWithDefaults() *MicrosoftGraphWorkbookChartSeries`

NewMicrosoftGraphWorkbookChartSeriesWithDefaults instantiates a new MicrosoftGraphWorkbookChartSeries object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphWorkbookChartSeries) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphWorkbookChartSeries) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphWorkbookChartSeries) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphWorkbookChartSeries) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *MicrosoftGraphWorkbookChartSeries) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MicrosoftGraphWorkbookChartSeries) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MicrosoftGraphWorkbookChartSeries) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MicrosoftGraphWorkbookChartSeries) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *MicrosoftGraphWorkbookChartSeries) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *MicrosoftGraphWorkbookChartSeries) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetFormat

`func (o *MicrosoftGraphWorkbookChartSeries) GetFormat() AnyOfmicrosoftGraphWorkbookChartSeriesFormat`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *MicrosoftGraphWorkbookChartSeries) GetFormatOk() (*AnyOfmicrosoftGraphWorkbookChartSeriesFormat, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *MicrosoftGraphWorkbookChartSeries) SetFormat(v AnyOfmicrosoftGraphWorkbookChartSeriesFormat)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *MicrosoftGraphWorkbookChartSeries) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### SetFormatNil

`func (o *MicrosoftGraphWorkbookChartSeries) SetFormatNil(b bool)`

 SetFormatNil sets the value for Format to be an explicit nil

### UnsetFormat
`func (o *MicrosoftGraphWorkbookChartSeries) UnsetFormat()`

UnsetFormat ensures that no value is present for Format, not even an explicit nil
### GetPoints

`func (o *MicrosoftGraphWorkbookChartSeries) GetPoints() []MicrosoftGraphWorkbookChartPoint`

GetPoints returns the Points field if non-nil, zero value otherwise.

### GetPointsOk

`func (o *MicrosoftGraphWorkbookChartSeries) GetPointsOk() (*[]MicrosoftGraphWorkbookChartPoint, bool)`

GetPointsOk returns a tuple with the Points field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoints

`func (o *MicrosoftGraphWorkbookChartSeries) SetPoints(v []MicrosoftGraphWorkbookChartPoint)`

SetPoints sets Points field to given value.

### HasPoints

`func (o *MicrosoftGraphWorkbookChartSeries) HasPoints() bool`

HasPoints returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


