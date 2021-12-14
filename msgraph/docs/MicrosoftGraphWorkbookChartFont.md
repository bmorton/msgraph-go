# MicrosoftGraphWorkbookChartFont

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**Bold** | Pointer to **NullableBool** | Represents the bold status of font. | [optional] 
**Color** | Pointer to **NullableString** | HTML color code representation of the text color. E.g. #FF0000 represents Red. | [optional] 
**Italic** | Pointer to **NullableBool** | Represents the italic status of the font. | [optional] 
**Name** | Pointer to **NullableString** | Font name (e.g. &#39;Calibri&#39;) | [optional] 
**Size** | Pointer to [**AnyOfnumberstringstring**](anyOf&lt;number,string,string&gt;.md) | Size of the font (e.g. 11) | [optional] 
**Underline** | Pointer to **NullableString** | Type of underline applied to the font. The possible values are: None, Single. | [optional] 

## Methods

### NewMicrosoftGraphWorkbookChartFont

`func NewMicrosoftGraphWorkbookChartFont() *MicrosoftGraphWorkbookChartFont`

NewMicrosoftGraphWorkbookChartFont instantiates a new MicrosoftGraphWorkbookChartFont object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphWorkbookChartFontWithDefaults

`func NewMicrosoftGraphWorkbookChartFontWithDefaults() *MicrosoftGraphWorkbookChartFont`

NewMicrosoftGraphWorkbookChartFontWithDefaults instantiates a new MicrosoftGraphWorkbookChartFont object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphWorkbookChartFont) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphWorkbookChartFont) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphWorkbookChartFont) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphWorkbookChartFont) HasId() bool`

HasId returns a boolean if a field has been set.

### GetBold

`func (o *MicrosoftGraphWorkbookChartFont) GetBold() bool`

GetBold returns the Bold field if non-nil, zero value otherwise.

### GetBoldOk

`func (o *MicrosoftGraphWorkbookChartFont) GetBoldOk() (*bool, bool)`

GetBoldOk returns a tuple with the Bold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBold

`func (o *MicrosoftGraphWorkbookChartFont) SetBold(v bool)`

SetBold sets Bold field to given value.

### HasBold

`func (o *MicrosoftGraphWorkbookChartFont) HasBold() bool`

HasBold returns a boolean if a field has been set.

### SetBoldNil

`func (o *MicrosoftGraphWorkbookChartFont) SetBoldNil(b bool)`

 SetBoldNil sets the value for Bold to be an explicit nil

### UnsetBold
`func (o *MicrosoftGraphWorkbookChartFont) UnsetBold()`

UnsetBold ensures that no value is present for Bold, not even an explicit nil
### GetColor

`func (o *MicrosoftGraphWorkbookChartFont) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *MicrosoftGraphWorkbookChartFont) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *MicrosoftGraphWorkbookChartFont) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *MicrosoftGraphWorkbookChartFont) HasColor() bool`

HasColor returns a boolean if a field has been set.

### SetColorNil

`func (o *MicrosoftGraphWorkbookChartFont) SetColorNil(b bool)`

 SetColorNil sets the value for Color to be an explicit nil

### UnsetColor
`func (o *MicrosoftGraphWorkbookChartFont) UnsetColor()`

UnsetColor ensures that no value is present for Color, not even an explicit nil
### GetItalic

`func (o *MicrosoftGraphWorkbookChartFont) GetItalic() bool`

GetItalic returns the Italic field if non-nil, zero value otherwise.

### GetItalicOk

`func (o *MicrosoftGraphWorkbookChartFont) GetItalicOk() (*bool, bool)`

GetItalicOk returns a tuple with the Italic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItalic

`func (o *MicrosoftGraphWorkbookChartFont) SetItalic(v bool)`

SetItalic sets Italic field to given value.

### HasItalic

`func (o *MicrosoftGraphWorkbookChartFont) HasItalic() bool`

HasItalic returns a boolean if a field has been set.

### SetItalicNil

`func (o *MicrosoftGraphWorkbookChartFont) SetItalicNil(b bool)`

 SetItalicNil sets the value for Italic to be an explicit nil

### UnsetItalic
`func (o *MicrosoftGraphWorkbookChartFont) UnsetItalic()`

UnsetItalic ensures that no value is present for Italic, not even an explicit nil
### GetName

`func (o *MicrosoftGraphWorkbookChartFont) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MicrosoftGraphWorkbookChartFont) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MicrosoftGraphWorkbookChartFont) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MicrosoftGraphWorkbookChartFont) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *MicrosoftGraphWorkbookChartFont) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *MicrosoftGraphWorkbookChartFont) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetSize

`func (o *MicrosoftGraphWorkbookChartFont) GetSize() AnyOfnumberstringstring`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *MicrosoftGraphWorkbookChartFont) GetSizeOk() (*AnyOfnumberstringstring, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *MicrosoftGraphWorkbookChartFont) SetSize(v AnyOfnumberstringstring)`

SetSize sets Size field to given value.

### HasSize

`func (o *MicrosoftGraphWorkbookChartFont) HasSize() bool`

HasSize returns a boolean if a field has been set.

### SetSizeNil

`func (o *MicrosoftGraphWorkbookChartFont) SetSizeNil(b bool)`

 SetSizeNil sets the value for Size to be an explicit nil

### UnsetSize
`func (o *MicrosoftGraphWorkbookChartFont) UnsetSize()`

UnsetSize ensures that no value is present for Size, not even an explicit nil
### GetUnderline

`func (o *MicrosoftGraphWorkbookChartFont) GetUnderline() string`

GetUnderline returns the Underline field if non-nil, zero value otherwise.

### GetUnderlineOk

`func (o *MicrosoftGraphWorkbookChartFont) GetUnderlineOk() (*string, bool)`

GetUnderlineOk returns a tuple with the Underline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderline

`func (o *MicrosoftGraphWorkbookChartFont) SetUnderline(v string)`

SetUnderline sets Underline field to given value.

### HasUnderline

`func (o *MicrosoftGraphWorkbookChartFont) HasUnderline() bool`

HasUnderline returns a boolean if a field has been set.

### SetUnderlineNil

`func (o *MicrosoftGraphWorkbookChartFont) SetUnderlineNil(b bool)`

 SetUnderlineNil sets the value for Underline to be an explicit nil

### UnsetUnderline
`func (o *MicrosoftGraphWorkbookChartFont) UnsetUnderline()`

UnsetUnderline ensures that no value is present for Underline, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


