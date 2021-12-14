# WorkbookChartSeriesFormat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fill** | Pointer to [**AnyOfmicrosoftGraphWorkbookChartFill**](anyOf&lt;microsoft.graph.workbookChartFill&gt;.md) | Represents the fill format of a chart series, which includes background formating information. Read-only. | [optional] 
**Line** | Pointer to [**AnyOfmicrosoftGraphWorkbookChartLineFormat**](anyOf&lt;microsoft.graph.workbookChartLineFormat&gt;.md) | Represents line formatting. Read-only. | [optional] 

## Methods

### NewWorkbookChartSeriesFormat

`func NewWorkbookChartSeriesFormat() *WorkbookChartSeriesFormat`

NewWorkbookChartSeriesFormat instantiates a new WorkbookChartSeriesFormat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkbookChartSeriesFormatWithDefaults

`func NewWorkbookChartSeriesFormatWithDefaults() *WorkbookChartSeriesFormat`

NewWorkbookChartSeriesFormatWithDefaults instantiates a new WorkbookChartSeriesFormat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFill

`func (o *WorkbookChartSeriesFormat) GetFill() AnyOfmicrosoftGraphWorkbookChartFill`

GetFill returns the Fill field if non-nil, zero value otherwise.

### GetFillOk

`func (o *WorkbookChartSeriesFormat) GetFillOk() (*AnyOfmicrosoftGraphWorkbookChartFill, bool)`

GetFillOk returns a tuple with the Fill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFill

`func (o *WorkbookChartSeriesFormat) SetFill(v AnyOfmicrosoftGraphWorkbookChartFill)`

SetFill sets Fill field to given value.

### HasFill

`func (o *WorkbookChartSeriesFormat) HasFill() bool`

HasFill returns a boolean if a field has been set.

### SetFillNil

`func (o *WorkbookChartSeriesFormat) SetFillNil(b bool)`

 SetFillNil sets the value for Fill to be an explicit nil

### UnsetFill
`func (o *WorkbookChartSeriesFormat) UnsetFill()`

UnsetFill ensures that no value is present for Fill, not even an explicit nil
### GetLine

`func (o *WorkbookChartSeriesFormat) GetLine() AnyOfmicrosoftGraphWorkbookChartLineFormat`

GetLine returns the Line field if non-nil, zero value otherwise.

### GetLineOk

`func (o *WorkbookChartSeriesFormat) GetLineOk() (*AnyOfmicrosoftGraphWorkbookChartLineFormat, bool)`

GetLineOk returns a tuple with the Line field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine

`func (o *WorkbookChartSeriesFormat) SetLine(v AnyOfmicrosoftGraphWorkbookChartLineFormat)`

SetLine sets Line field to given value.

### HasLine

`func (o *WorkbookChartSeriesFormat) HasLine() bool`

HasLine returns a boolean if a field has been set.

### SetLineNil

`func (o *WorkbookChartSeriesFormat) SetLineNil(b bool)`

 SetLineNil sets the value for Line to be an explicit nil

### UnsetLine
`func (o *WorkbookChartSeriesFormat) UnsetLine()`

UnsetLine ensures that no value is present for Line, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


