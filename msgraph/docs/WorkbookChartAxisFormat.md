# WorkbookChartAxisFormat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Font** | Pointer to [**AnyOfmicrosoftGraphWorkbookChartFont**](anyOf&lt;microsoft.graph.workbookChartFont&gt;.md) | Represents the font attributes (font name, font size, color, etc.) for a chart axis element. Read-only. | [optional] 
**Line** | Pointer to [**AnyOfmicrosoftGraphWorkbookChartLineFormat**](anyOf&lt;microsoft.graph.workbookChartLineFormat&gt;.md) | Represents chart line formatting. Read-only. | [optional] 

## Methods

### NewWorkbookChartAxisFormat

`func NewWorkbookChartAxisFormat() *WorkbookChartAxisFormat`

NewWorkbookChartAxisFormat instantiates a new WorkbookChartAxisFormat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkbookChartAxisFormatWithDefaults

`func NewWorkbookChartAxisFormatWithDefaults() *WorkbookChartAxisFormat`

NewWorkbookChartAxisFormatWithDefaults instantiates a new WorkbookChartAxisFormat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFont

`func (o *WorkbookChartAxisFormat) GetFont() AnyOfmicrosoftGraphWorkbookChartFont`

GetFont returns the Font field if non-nil, zero value otherwise.

### GetFontOk

`func (o *WorkbookChartAxisFormat) GetFontOk() (*AnyOfmicrosoftGraphWorkbookChartFont, bool)`

GetFontOk returns a tuple with the Font field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFont

`func (o *WorkbookChartAxisFormat) SetFont(v AnyOfmicrosoftGraphWorkbookChartFont)`

SetFont sets Font field to given value.

### HasFont

`func (o *WorkbookChartAxisFormat) HasFont() bool`

HasFont returns a boolean if a field has been set.

### SetFontNil

`func (o *WorkbookChartAxisFormat) SetFontNil(b bool)`

 SetFontNil sets the value for Font to be an explicit nil

### UnsetFont
`func (o *WorkbookChartAxisFormat) UnsetFont()`

UnsetFont ensures that no value is present for Font, not even an explicit nil
### GetLine

`func (o *WorkbookChartAxisFormat) GetLine() AnyOfmicrosoftGraphWorkbookChartLineFormat`

GetLine returns the Line field if non-nil, zero value otherwise.

### GetLineOk

`func (o *WorkbookChartAxisFormat) GetLineOk() (*AnyOfmicrosoftGraphWorkbookChartLineFormat, bool)`

GetLineOk returns a tuple with the Line field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine

`func (o *WorkbookChartAxisFormat) SetLine(v AnyOfmicrosoftGraphWorkbookChartLineFormat)`

SetLine sets Line field to given value.

### HasLine

`func (o *WorkbookChartAxisFormat) HasLine() bool`

HasLine returns a boolean if a field has been set.

### SetLineNil

`func (o *WorkbookChartAxisFormat) SetLineNil(b bool)`

 SetLineNil sets the value for Line to be an explicit nil

### UnsetLine
`func (o *WorkbookChartAxisFormat) UnsetLine()`

UnsetLine ensures that no value is present for Line, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


