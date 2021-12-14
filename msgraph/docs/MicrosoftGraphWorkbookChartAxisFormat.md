# MicrosoftGraphWorkbookChartAxisFormat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**Font** | Pointer to [**AnyOfmicrosoftGraphWorkbookChartFont**](anyOf&lt;microsoft.graph.workbookChartFont&gt;.md) | Represents the font attributes (font name, font size, color, etc.) for a chart axis element. Read-only. | [optional] 
**Line** | Pointer to [**AnyOfmicrosoftGraphWorkbookChartLineFormat**](anyOf&lt;microsoft.graph.workbookChartLineFormat&gt;.md) | Represents chart line formatting. Read-only. | [optional] 

## Methods

### NewMicrosoftGraphWorkbookChartAxisFormat

`func NewMicrosoftGraphWorkbookChartAxisFormat() *MicrosoftGraphWorkbookChartAxisFormat`

NewMicrosoftGraphWorkbookChartAxisFormat instantiates a new MicrosoftGraphWorkbookChartAxisFormat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphWorkbookChartAxisFormatWithDefaults

`func NewMicrosoftGraphWorkbookChartAxisFormatWithDefaults() *MicrosoftGraphWorkbookChartAxisFormat`

NewMicrosoftGraphWorkbookChartAxisFormatWithDefaults instantiates a new MicrosoftGraphWorkbookChartAxisFormat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphWorkbookChartAxisFormat) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphWorkbookChartAxisFormat) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphWorkbookChartAxisFormat) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphWorkbookChartAxisFormat) HasId() bool`

HasId returns a boolean if a field has been set.

### GetFont

`func (o *MicrosoftGraphWorkbookChartAxisFormat) GetFont() AnyOfmicrosoftGraphWorkbookChartFont`

GetFont returns the Font field if non-nil, zero value otherwise.

### GetFontOk

`func (o *MicrosoftGraphWorkbookChartAxisFormat) GetFontOk() (*AnyOfmicrosoftGraphWorkbookChartFont, bool)`

GetFontOk returns a tuple with the Font field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFont

`func (o *MicrosoftGraphWorkbookChartAxisFormat) SetFont(v AnyOfmicrosoftGraphWorkbookChartFont)`

SetFont sets Font field to given value.

### HasFont

`func (o *MicrosoftGraphWorkbookChartAxisFormat) HasFont() bool`

HasFont returns a boolean if a field has been set.

### SetFontNil

`func (o *MicrosoftGraphWorkbookChartAxisFormat) SetFontNil(b bool)`

 SetFontNil sets the value for Font to be an explicit nil

### UnsetFont
`func (o *MicrosoftGraphWorkbookChartAxisFormat) UnsetFont()`

UnsetFont ensures that no value is present for Font, not even an explicit nil
### GetLine

`func (o *MicrosoftGraphWorkbookChartAxisFormat) GetLine() AnyOfmicrosoftGraphWorkbookChartLineFormat`

GetLine returns the Line field if non-nil, zero value otherwise.

### GetLineOk

`func (o *MicrosoftGraphWorkbookChartAxisFormat) GetLineOk() (*AnyOfmicrosoftGraphWorkbookChartLineFormat, bool)`

GetLineOk returns a tuple with the Line field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine

`func (o *MicrosoftGraphWorkbookChartAxisFormat) SetLine(v AnyOfmicrosoftGraphWorkbookChartLineFormat)`

SetLine sets Line field to given value.

### HasLine

`func (o *MicrosoftGraphWorkbookChartAxisFormat) HasLine() bool`

HasLine returns a boolean if a field has been set.

### SetLineNil

`func (o *MicrosoftGraphWorkbookChartAxisFormat) SetLineNil(b bool)`

 SetLineNil sets the value for Line to be an explicit nil

### UnsetLine
`func (o *MicrosoftGraphWorkbookChartAxisFormat) UnsetLine()`

UnsetLine ensures that no value is present for Line, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


