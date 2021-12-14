# WorkbookChartAxisTitle

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Text** | Pointer to **NullableString** | Represents the axis title. | [optional] 
**Visible** | Pointer to **bool** | A boolean that specifies the visibility of an axis title. | [optional] 
**Format** | Pointer to [**AnyOfmicrosoftGraphWorkbookChartAxisTitleFormat**](anyOf&lt;microsoft.graph.workbookChartAxisTitleFormat&gt;.md) | Represents the formatting of chart axis title. Read-only. | [optional] 

## Methods

### NewWorkbookChartAxisTitle

`func NewWorkbookChartAxisTitle() *WorkbookChartAxisTitle`

NewWorkbookChartAxisTitle instantiates a new WorkbookChartAxisTitle object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkbookChartAxisTitleWithDefaults

`func NewWorkbookChartAxisTitleWithDefaults() *WorkbookChartAxisTitle`

NewWorkbookChartAxisTitleWithDefaults instantiates a new WorkbookChartAxisTitle object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetText

`func (o *WorkbookChartAxisTitle) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *WorkbookChartAxisTitle) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *WorkbookChartAxisTitle) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *WorkbookChartAxisTitle) HasText() bool`

HasText returns a boolean if a field has been set.

### SetTextNil

`func (o *WorkbookChartAxisTitle) SetTextNil(b bool)`

 SetTextNil sets the value for Text to be an explicit nil

### UnsetText
`func (o *WorkbookChartAxisTitle) UnsetText()`

UnsetText ensures that no value is present for Text, not even an explicit nil
### GetVisible

`func (o *WorkbookChartAxisTitle) GetVisible() bool`

GetVisible returns the Visible field if non-nil, zero value otherwise.

### GetVisibleOk

`func (o *WorkbookChartAxisTitle) GetVisibleOk() (*bool, bool)`

GetVisibleOk returns a tuple with the Visible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisible

`func (o *WorkbookChartAxisTitle) SetVisible(v bool)`

SetVisible sets Visible field to given value.

### HasVisible

`func (o *WorkbookChartAxisTitle) HasVisible() bool`

HasVisible returns a boolean if a field has been set.

### GetFormat

`func (o *WorkbookChartAxisTitle) GetFormat() AnyOfmicrosoftGraphWorkbookChartAxisTitleFormat`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *WorkbookChartAxisTitle) GetFormatOk() (*AnyOfmicrosoftGraphWorkbookChartAxisTitleFormat, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *WorkbookChartAxisTitle) SetFormat(v AnyOfmicrosoftGraphWorkbookChartAxisTitleFormat)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *WorkbookChartAxisTitle) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### SetFormatNil

`func (o *WorkbookChartAxisTitle) SetFormatNil(b bool)`

 SetFormatNil sets the value for Format to be an explicit nil

### UnsetFormat
`func (o *WorkbookChartAxisTitle) UnsetFormat()`

UnsetFormat ensures that no value is present for Format, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


