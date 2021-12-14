# MicrosoftGraphWorkbookChartSeriesFormat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**Fill** | Pointer to [**AnyOfmicrosoftGraphWorkbookChartFill**](anyOf&lt;microsoft.graph.workbookChartFill&gt;.md) | Represents the fill format of a chart series, which includes background formating information. Read-only. | [optional] 
**Line** | Pointer to [**AnyOfmicrosoftGraphWorkbookChartLineFormat**](anyOf&lt;microsoft.graph.workbookChartLineFormat&gt;.md) | Represents line formatting. Read-only. | [optional] 

## Methods

### NewMicrosoftGraphWorkbookChartSeriesFormat

`func NewMicrosoftGraphWorkbookChartSeriesFormat() *MicrosoftGraphWorkbookChartSeriesFormat`

NewMicrosoftGraphWorkbookChartSeriesFormat instantiates a new MicrosoftGraphWorkbookChartSeriesFormat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphWorkbookChartSeriesFormatWithDefaults

`func NewMicrosoftGraphWorkbookChartSeriesFormatWithDefaults() *MicrosoftGraphWorkbookChartSeriesFormat`

NewMicrosoftGraphWorkbookChartSeriesFormatWithDefaults instantiates a new MicrosoftGraphWorkbookChartSeriesFormat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphWorkbookChartSeriesFormat) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphWorkbookChartSeriesFormat) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphWorkbookChartSeriesFormat) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphWorkbookChartSeriesFormat) HasId() bool`

HasId returns a boolean if a field has been set.

### GetFill

`func (o *MicrosoftGraphWorkbookChartSeriesFormat) GetFill() AnyOfmicrosoftGraphWorkbookChartFill`

GetFill returns the Fill field if non-nil, zero value otherwise.

### GetFillOk

`func (o *MicrosoftGraphWorkbookChartSeriesFormat) GetFillOk() (*AnyOfmicrosoftGraphWorkbookChartFill, bool)`

GetFillOk returns a tuple with the Fill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFill

`func (o *MicrosoftGraphWorkbookChartSeriesFormat) SetFill(v AnyOfmicrosoftGraphWorkbookChartFill)`

SetFill sets Fill field to given value.

### HasFill

`func (o *MicrosoftGraphWorkbookChartSeriesFormat) HasFill() bool`

HasFill returns a boolean if a field has been set.

### SetFillNil

`func (o *MicrosoftGraphWorkbookChartSeriesFormat) SetFillNil(b bool)`

 SetFillNil sets the value for Fill to be an explicit nil

### UnsetFill
`func (o *MicrosoftGraphWorkbookChartSeriesFormat) UnsetFill()`

UnsetFill ensures that no value is present for Fill, not even an explicit nil
### GetLine

`func (o *MicrosoftGraphWorkbookChartSeriesFormat) GetLine() AnyOfmicrosoftGraphWorkbookChartLineFormat`

GetLine returns the Line field if non-nil, zero value otherwise.

### GetLineOk

`func (o *MicrosoftGraphWorkbookChartSeriesFormat) GetLineOk() (*AnyOfmicrosoftGraphWorkbookChartLineFormat, bool)`

GetLineOk returns a tuple with the Line field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine

`func (o *MicrosoftGraphWorkbookChartSeriesFormat) SetLine(v AnyOfmicrosoftGraphWorkbookChartLineFormat)`

SetLine sets Line field to given value.

### HasLine

`func (o *MicrosoftGraphWorkbookChartSeriesFormat) HasLine() bool`

HasLine returns a boolean if a field has been set.

### SetLineNil

`func (o *MicrosoftGraphWorkbookChartSeriesFormat) SetLineNil(b bool)`

 SetLineNil sets the value for Line to be an explicit nil

### UnsetLine
`func (o *MicrosoftGraphWorkbookChartSeriesFormat) UnsetLine()`

UnsetLine ensures that no value is present for Line, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


