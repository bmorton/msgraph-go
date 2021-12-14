# MicrosoftGraphWorkbookRangeFont

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**Bold** | Pointer to **NullableBool** | Represents the bold status of font. | [optional] 
**Color** | Pointer to **NullableString** | HTML color code representation of the text color. E.g. #FF0000 represents Red. | [optional] 
**Italic** | Pointer to **NullableBool** | Represents the italic status of the font. | [optional] 
**Name** | Pointer to **NullableString** | Font name (e.g. &#39;Calibri&#39;) | [optional] 
**Size** | Pointer to [**AnyOfnumberstringstring**](anyOf&lt;number,string,string&gt;.md) | Font size. | [optional] 
**Underline** | Pointer to **NullableString** | Type of underline applied to the font. The possible values are: None, Single, Double, SingleAccountant, DoubleAccountant. | [optional] 

## Methods

### NewMicrosoftGraphWorkbookRangeFont

`func NewMicrosoftGraphWorkbookRangeFont() *MicrosoftGraphWorkbookRangeFont`

NewMicrosoftGraphWorkbookRangeFont instantiates a new MicrosoftGraphWorkbookRangeFont object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphWorkbookRangeFontWithDefaults

`func NewMicrosoftGraphWorkbookRangeFontWithDefaults() *MicrosoftGraphWorkbookRangeFont`

NewMicrosoftGraphWorkbookRangeFontWithDefaults instantiates a new MicrosoftGraphWorkbookRangeFont object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphWorkbookRangeFont) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphWorkbookRangeFont) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphWorkbookRangeFont) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphWorkbookRangeFont) HasId() bool`

HasId returns a boolean if a field has been set.

### GetBold

`func (o *MicrosoftGraphWorkbookRangeFont) GetBold() bool`

GetBold returns the Bold field if non-nil, zero value otherwise.

### GetBoldOk

`func (o *MicrosoftGraphWorkbookRangeFont) GetBoldOk() (*bool, bool)`

GetBoldOk returns a tuple with the Bold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBold

`func (o *MicrosoftGraphWorkbookRangeFont) SetBold(v bool)`

SetBold sets Bold field to given value.

### HasBold

`func (o *MicrosoftGraphWorkbookRangeFont) HasBold() bool`

HasBold returns a boolean if a field has been set.

### SetBoldNil

`func (o *MicrosoftGraphWorkbookRangeFont) SetBoldNil(b bool)`

 SetBoldNil sets the value for Bold to be an explicit nil

### UnsetBold
`func (o *MicrosoftGraphWorkbookRangeFont) UnsetBold()`

UnsetBold ensures that no value is present for Bold, not even an explicit nil
### GetColor

`func (o *MicrosoftGraphWorkbookRangeFont) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *MicrosoftGraphWorkbookRangeFont) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *MicrosoftGraphWorkbookRangeFont) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *MicrosoftGraphWorkbookRangeFont) HasColor() bool`

HasColor returns a boolean if a field has been set.

### SetColorNil

`func (o *MicrosoftGraphWorkbookRangeFont) SetColorNil(b bool)`

 SetColorNil sets the value for Color to be an explicit nil

### UnsetColor
`func (o *MicrosoftGraphWorkbookRangeFont) UnsetColor()`

UnsetColor ensures that no value is present for Color, not even an explicit nil
### GetItalic

`func (o *MicrosoftGraphWorkbookRangeFont) GetItalic() bool`

GetItalic returns the Italic field if non-nil, zero value otherwise.

### GetItalicOk

`func (o *MicrosoftGraphWorkbookRangeFont) GetItalicOk() (*bool, bool)`

GetItalicOk returns a tuple with the Italic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItalic

`func (o *MicrosoftGraphWorkbookRangeFont) SetItalic(v bool)`

SetItalic sets Italic field to given value.

### HasItalic

`func (o *MicrosoftGraphWorkbookRangeFont) HasItalic() bool`

HasItalic returns a boolean if a field has been set.

### SetItalicNil

`func (o *MicrosoftGraphWorkbookRangeFont) SetItalicNil(b bool)`

 SetItalicNil sets the value for Italic to be an explicit nil

### UnsetItalic
`func (o *MicrosoftGraphWorkbookRangeFont) UnsetItalic()`

UnsetItalic ensures that no value is present for Italic, not even an explicit nil
### GetName

`func (o *MicrosoftGraphWorkbookRangeFont) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MicrosoftGraphWorkbookRangeFont) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MicrosoftGraphWorkbookRangeFont) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MicrosoftGraphWorkbookRangeFont) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *MicrosoftGraphWorkbookRangeFont) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *MicrosoftGraphWorkbookRangeFont) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetSize

`func (o *MicrosoftGraphWorkbookRangeFont) GetSize() AnyOfnumberstringstring`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *MicrosoftGraphWorkbookRangeFont) GetSizeOk() (*AnyOfnumberstringstring, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *MicrosoftGraphWorkbookRangeFont) SetSize(v AnyOfnumberstringstring)`

SetSize sets Size field to given value.

### HasSize

`func (o *MicrosoftGraphWorkbookRangeFont) HasSize() bool`

HasSize returns a boolean if a field has been set.

### SetSizeNil

`func (o *MicrosoftGraphWorkbookRangeFont) SetSizeNil(b bool)`

 SetSizeNil sets the value for Size to be an explicit nil

### UnsetSize
`func (o *MicrosoftGraphWorkbookRangeFont) UnsetSize()`

UnsetSize ensures that no value is present for Size, not even an explicit nil
### GetUnderline

`func (o *MicrosoftGraphWorkbookRangeFont) GetUnderline() string`

GetUnderline returns the Underline field if non-nil, zero value otherwise.

### GetUnderlineOk

`func (o *MicrosoftGraphWorkbookRangeFont) GetUnderlineOk() (*string, bool)`

GetUnderlineOk returns a tuple with the Underline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderline

`func (o *MicrosoftGraphWorkbookRangeFont) SetUnderline(v string)`

SetUnderline sets Underline field to given value.

### HasUnderline

`func (o *MicrosoftGraphWorkbookRangeFont) HasUnderline() bool`

HasUnderline returns a boolean if a field has been set.

### SetUnderlineNil

`func (o *MicrosoftGraphWorkbookRangeFont) SetUnderlineNil(b bool)`

 SetUnderlineNil sets the value for Underline to be an explicit nil

### UnsetUnderline
`func (o *MicrosoftGraphWorkbookRangeFont) UnsetUnderline()`

UnsetUnderline ensures that no value is present for Underline, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


