# MicrosoftGraphWorkbookChartAreaFormat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**Fill** | Pointer to [**AnyOfmicrosoftGraphWorkbookChartFill**](anyOf&lt;microsoft.graph.workbookChartFill&gt;.md) | Represents the fill format of an object, which includes background formatting information. Read-only. | [optional] 
**Font** | Pointer to [**AnyOfmicrosoftGraphWorkbookChartFont**](anyOf&lt;microsoft.graph.workbookChartFont&gt;.md) | Represents the font attributes (font name, font size, color, etc.) for the current object. Read-only. | [optional] 

## Methods

### NewMicrosoftGraphWorkbookChartAreaFormat

`func NewMicrosoftGraphWorkbookChartAreaFormat() *MicrosoftGraphWorkbookChartAreaFormat`

NewMicrosoftGraphWorkbookChartAreaFormat instantiates a new MicrosoftGraphWorkbookChartAreaFormat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphWorkbookChartAreaFormatWithDefaults

`func NewMicrosoftGraphWorkbookChartAreaFormatWithDefaults() *MicrosoftGraphWorkbookChartAreaFormat`

NewMicrosoftGraphWorkbookChartAreaFormatWithDefaults instantiates a new MicrosoftGraphWorkbookChartAreaFormat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphWorkbookChartAreaFormat) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphWorkbookChartAreaFormat) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphWorkbookChartAreaFormat) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphWorkbookChartAreaFormat) HasId() bool`

HasId returns a boolean if a field has been set.

### GetFill

`func (o *MicrosoftGraphWorkbookChartAreaFormat) GetFill() AnyOfmicrosoftGraphWorkbookChartFill`

GetFill returns the Fill field if non-nil, zero value otherwise.

### GetFillOk

`func (o *MicrosoftGraphWorkbookChartAreaFormat) GetFillOk() (*AnyOfmicrosoftGraphWorkbookChartFill, bool)`

GetFillOk returns a tuple with the Fill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFill

`func (o *MicrosoftGraphWorkbookChartAreaFormat) SetFill(v AnyOfmicrosoftGraphWorkbookChartFill)`

SetFill sets Fill field to given value.

### HasFill

`func (o *MicrosoftGraphWorkbookChartAreaFormat) HasFill() bool`

HasFill returns a boolean if a field has been set.

### SetFillNil

`func (o *MicrosoftGraphWorkbookChartAreaFormat) SetFillNil(b bool)`

 SetFillNil sets the value for Fill to be an explicit nil

### UnsetFill
`func (o *MicrosoftGraphWorkbookChartAreaFormat) UnsetFill()`

UnsetFill ensures that no value is present for Fill, not even an explicit nil
### GetFont

`func (o *MicrosoftGraphWorkbookChartAreaFormat) GetFont() AnyOfmicrosoftGraphWorkbookChartFont`

GetFont returns the Font field if non-nil, zero value otherwise.

### GetFontOk

`func (o *MicrosoftGraphWorkbookChartAreaFormat) GetFontOk() (*AnyOfmicrosoftGraphWorkbookChartFont, bool)`

GetFontOk returns a tuple with the Font field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFont

`func (o *MicrosoftGraphWorkbookChartAreaFormat) SetFont(v AnyOfmicrosoftGraphWorkbookChartFont)`

SetFont sets Font field to given value.

### HasFont

`func (o *MicrosoftGraphWorkbookChartAreaFormat) HasFont() bool`

HasFont returns a boolean if a field has been set.

### SetFontNil

`func (o *MicrosoftGraphWorkbookChartAreaFormat) SetFontNil(b bool)`

 SetFontNil sets the value for Font to be an explicit nil

### UnsetFont
`func (o *MicrosoftGraphWorkbookChartAreaFormat) UnsetFont()`

UnsetFont ensures that no value is present for Font, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


