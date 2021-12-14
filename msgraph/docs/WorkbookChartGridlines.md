# WorkbookChartGridlines

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Visible** | Pointer to **bool** | Boolean value representing if the axis gridlines are visible or not. | [optional] 
**Format** | Pointer to [**AnyOfmicrosoftGraphWorkbookChartGridlinesFormat**](anyOf&lt;microsoft.graph.workbookChartGridlinesFormat&gt;.md) | Represents the formatting of chart gridlines. Read-only. | [optional] 

## Methods

### NewWorkbookChartGridlines

`func NewWorkbookChartGridlines() *WorkbookChartGridlines`

NewWorkbookChartGridlines instantiates a new WorkbookChartGridlines object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkbookChartGridlinesWithDefaults

`func NewWorkbookChartGridlinesWithDefaults() *WorkbookChartGridlines`

NewWorkbookChartGridlinesWithDefaults instantiates a new WorkbookChartGridlines object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVisible

`func (o *WorkbookChartGridlines) GetVisible() bool`

GetVisible returns the Visible field if non-nil, zero value otherwise.

### GetVisibleOk

`func (o *WorkbookChartGridlines) GetVisibleOk() (*bool, bool)`

GetVisibleOk returns a tuple with the Visible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisible

`func (o *WorkbookChartGridlines) SetVisible(v bool)`

SetVisible sets Visible field to given value.

### HasVisible

`func (o *WorkbookChartGridlines) HasVisible() bool`

HasVisible returns a boolean if a field has been set.

### GetFormat

`func (o *WorkbookChartGridlines) GetFormat() AnyOfmicrosoftGraphWorkbookChartGridlinesFormat`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *WorkbookChartGridlines) GetFormatOk() (*AnyOfmicrosoftGraphWorkbookChartGridlinesFormat, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *WorkbookChartGridlines) SetFormat(v AnyOfmicrosoftGraphWorkbookChartGridlinesFormat)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *WorkbookChartGridlines) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### SetFormatNil

`func (o *WorkbookChartGridlines) SetFormatNil(b bool)`

 SetFormatNil sets the value for Format to be an explicit nil

### UnsetFormat
`func (o *WorkbookChartGridlines) UnsetFormat()`

UnsetFormat ensures that no value is present for Format, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


