# WorkbookRangeFont

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bold** | Pointer to **NullableBool** | Represents the bold status of font. | [optional] 
**Color** | Pointer to **NullableString** | HTML color code representation of the text color. E.g. #FF0000 represents Red. | [optional] 
**Italic** | Pointer to **NullableBool** | Represents the italic status of the font. | [optional] 
**Name** | Pointer to **NullableString** | Font name (e.g. &#39;Calibri&#39;) | [optional] 
**Size** | Pointer to [**AnyOfnumberstringstring**](anyOf&lt;number,string,string&gt;.md) | Font size. | [optional] 
**Underline** | Pointer to **NullableString** | Type of underline applied to the font. The possible values are: None, Single, Double, SingleAccountant, DoubleAccountant. | [optional] 

## Methods

### NewWorkbookRangeFont

`func NewWorkbookRangeFont() *WorkbookRangeFont`

NewWorkbookRangeFont instantiates a new WorkbookRangeFont object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkbookRangeFontWithDefaults

`func NewWorkbookRangeFontWithDefaults() *WorkbookRangeFont`

NewWorkbookRangeFontWithDefaults instantiates a new WorkbookRangeFont object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBold

`func (o *WorkbookRangeFont) GetBold() bool`

GetBold returns the Bold field if non-nil, zero value otherwise.

### GetBoldOk

`func (o *WorkbookRangeFont) GetBoldOk() (*bool, bool)`

GetBoldOk returns a tuple with the Bold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBold

`func (o *WorkbookRangeFont) SetBold(v bool)`

SetBold sets Bold field to given value.

### HasBold

`func (o *WorkbookRangeFont) HasBold() bool`

HasBold returns a boolean if a field has been set.

### SetBoldNil

`func (o *WorkbookRangeFont) SetBoldNil(b bool)`

 SetBoldNil sets the value for Bold to be an explicit nil

### UnsetBold
`func (o *WorkbookRangeFont) UnsetBold()`

UnsetBold ensures that no value is present for Bold, not even an explicit nil
### GetColor

`func (o *WorkbookRangeFont) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *WorkbookRangeFont) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *WorkbookRangeFont) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *WorkbookRangeFont) HasColor() bool`

HasColor returns a boolean if a field has been set.

### SetColorNil

`func (o *WorkbookRangeFont) SetColorNil(b bool)`

 SetColorNil sets the value for Color to be an explicit nil

### UnsetColor
`func (o *WorkbookRangeFont) UnsetColor()`

UnsetColor ensures that no value is present for Color, not even an explicit nil
### GetItalic

`func (o *WorkbookRangeFont) GetItalic() bool`

GetItalic returns the Italic field if non-nil, zero value otherwise.

### GetItalicOk

`func (o *WorkbookRangeFont) GetItalicOk() (*bool, bool)`

GetItalicOk returns a tuple with the Italic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItalic

`func (o *WorkbookRangeFont) SetItalic(v bool)`

SetItalic sets Italic field to given value.

### HasItalic

`func (o *WorkbookRangeFont) HasItalic() bool`

HasItalic returns a boolean if a field has been set.

### SetItalicNil

`func (o *WorkbookRangeFont) SetItalicNil(b bool)`

 SetItalicNil sets the value for Italic to be an explicit nil

### UnsetItalic
`func (o *WorkbookRangeFont) UnsetItalic()`

UnsetItalic ensures that no value is present for Italic, not even an explicit nil
### GetName

`func (o *WorkbookRangeFont) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkbookRangeFont) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkbookRangeFont) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WorkbookRangeFont) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *WorkbookRangeFont) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *WorkbookRangeFont) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetSize

`func (o *WorkbookRangeFont) GetSize() AnyOfnumberstringstring`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *WorkbookRangeFont) GetSizeOk() (*AnyOfnumberstringstring, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *WorkbookRangeFont) SetSize(v AnyOfnumberstringstring)`

SetSize sets Size field to given value.

### HasSize

`func (o *WorkbookRangeFont) HasSize() bool`

HasSize returns a boolean if a field has been set.

### SetSizeNil

`func (o *WorkbookRangeFont) SetSizeNil(b bool)`

 SetSizeNil sets the value for Size to be an explicit nil

### UnsetSize
`func (o *WorkbookRangeFont) UnsetSize()`

UnsetSize ensures that no value is present for Size, not even an explicit nil
### GetUnderline

`func (o *WorkbookRangeFont) GetUnderline() string`

GetUnderline returns the Underline field if non-nil, zero value otherwise.

### GetUnderlineOk

`func (o *WorkbookRangeFont) GetUnderlineOk() (*string, bool)`

GetUnderlineOk returns a tuple with the Underline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderline

`func (o *WorkbookRangeFont) SetUnderline(v string)`

SetUnderline sets Underline field to given value.

### HasUnderline

`func (o *WorkbookRangeFont) HasUnderline() bool`

HasUnderline returns a boolean if a field has been set.

### SetUnderlineNil

`func (o *WorkbookRangeFont) SetUnderlineNil(b bool)`

 SetUnderlineNil sets the value for Underline to be an explicit nil

### UnsetUnderline
`func (o *WorkbookRangeFont) UnsetUnderline()`

UnsetUnderline ensures that no value is present for Underline, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


