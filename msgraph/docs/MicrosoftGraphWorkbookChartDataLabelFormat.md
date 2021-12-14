# MicrosoftGraphWorkbookChartDataLabelFormat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**Fill** | Pointer to [**AnyOfmicrosoftGraphWorkbookChartFill**](anyOf&lt;microsoft.graph.workbookChartFill&gt;.md) | Represents the fill format of the current chart data label. Read-only. | [optional] 
**Font** | Pointer to [**AnyOfmicrosoftGraphWorkbookChartFont**](anyOf&lt;microsoft.graph.workbookChartFont&gt;.md) | Represents the font attributes (font name, font size, color, etc.) for a chart data label. Read-only. | [optional] 

## Methods

### NewMicrosoftGraphWorkbookChartDataLabelFormat

`func NewMicrosoftGraphWorkbookChartDataLabelFormat() *MicrosoftGraphWorkbookChartDataLabelFormat`

NewMicrosoftGraphWorkbookChartDataLabelFormat instantiates a new MicrosoftGraphWorkbookChartDataLabelFormat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphWorkbookChartDataLabelFormatWithDefaults

`func NewMicrosoftGraphWorkbookChartDataLabelFormatWithDefaults() *MicrosoftGraphWorkbookChartDataLabelFormat`

NewMicrosoftGraphWorkbookChartDataLabelFormatWithDefaults instantiates a new MicrosoftGraphWorkbookChartDataLabelFormat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphWorkbookChartDataLabelFormat) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphWorkbookChartDataLabelFormat) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphWorkbookChartDataLabelFormat) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphWorkbookChartDataLabelFormat) HasId() bool`

HasId returns a boolean if a field has been set.

### GetFill

`func (o *MicrosoftGraphWorkbookChartDataLabelFormat) GetFill() AnyOfmicrosoftGraphWorkbookChartFill`

GetFill returns the Fill field if non-nil, zero value otherwise.

### GetFillOk

`func (o *MicrosoftGraphWorkbookChartDataLabelFormat) GetFillOk() (*AnyOfmicrosoftGraphWorkbookChartFill, bool)`

GetFillOk returns a tuple with the Fill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFill

`func (o *MicrosoftGraphWorkbookChartDataLabelFormat) SetFill(v AnyOfmicrosoftGraphWorkbookChartFill)`

SetFill sets Fill field to given value.

### HasFill

`func (o *MicrosoftGraphWorkbookChartDataLabelFormat) HasFill() bool`

HasFill returns a boolean if a field has been set.

### SetFillNil

`func (o *MicrosoftGraphWorkbookChartDataLabelFormat) SetFillNil(b bool)`

 SetFillNil sets the value for Fill to be an explicit nil

### UnsetFill
`func (o *MicrosoftGraphWorkbookChartDataLabelFormat) UnsetFill()`

UnsetFill ensures that no value is present for Fill, not even an explicit nil
### GetFont

`func (o *MicrosoftGraphWorkbookChartDataLabelFormat) GetFont() AnyOfmicrosoftGraphWorkbookChartFont`

GetFont returns the Font field if non-nil, zero value otherwise.

### GetFontOk

`func (o *MicrosoftGraphWorkbookChartDataLabelFormat) GetFontOk() (*AnyOfmicrosoftGraphWorkbookChartFont, bool)`

GetFontOk returns a tuple with the Font field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFont

`func (o *MicrosoftGraphWorkbookChartDataLabelFormat) SetFont(v AnyOfmicrosoftGraphWorkbookChartFont)`

SetFont sets Font field to given value.

### HasFont

`func (o *MicrosoftGraphWorkbookChartDataLabelFormat) HasFont() bool`

HasFont returns a boolean if a field has been set.

### SetFontNil

`func (o *MicrosoftGraphWorkbookChartDataLabelFormat) SetFontNil(b bool)`

 SetFontNil sets the value for Font to be an explicit nil

### UnsetFont
`func (o *MicrosoftGraphWorkbookChartDataLabelFormat) UnsetFont()`

UnsetFont ensures that no value is present for Font, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


