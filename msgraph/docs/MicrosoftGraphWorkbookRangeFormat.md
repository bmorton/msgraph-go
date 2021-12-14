# MicrosoftGraphWorkbookRangeFormat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**ColumnWidth** | Pointer to [**AnyOfnumberstringstring**](anyOf&lt;number,string,string&gt;.md) | Gets or sets the width of all colums within the range. If the column widths are not uniform, null will be returned. | [optional] 
**HorizontalAlignment** | Pointer to **NullableString** | Represents the horizontal alignment for the specified object. The possible values are: General, Left, Center, Right, Fill, Justify, CenterAcrossSelection, Distributed. | [optional] 
**RowHeight** | Pointer to [**AnyOfnumberstringstring**](anyOf&lt;number,string,string&gt;.md) | Gets or sets the height of all rows in the range. If the row heights are not uniform null will be returned. | [optional] 
**VerticalAlignment** | Pointer to **NullableString** | Represents the vertical alignment for the specified object. The possible values are: Top, Center, Bottom, Justify, Distributed. | [optional] 
**WrapText** | Pointer to **NullableBool** | Indicates if Excel wraps the text in the object. A null value indicates that the entire range doesn&#39;t have uniform wrap setting | [optional] 
**Borders** | Pointer to [**[]MicrosoftGraphWorkbookRangeBorder**](MicrosoftGraphWorkbookRangeBorder.md) | Collection of border objects that apply to the overall range selected Read-only. | [optional] 
**Fill** | Pointer to [**AnyOfmicrosoftGraphWorkbookRangeFill**](anyOf&lt;microsoft.graph.workbookRangeFill&gt;.md) | Returns the fill object defined on the overall range. Read-only. | [optional] 
**Font** | Pointer to [**AnyOfmicrosoftGraphWorkbookRangeFont**](anyOf&lt;microsoft.graph.workbookRangeFont&gt;.md) | Returns the font object defined on the overall range selected Read-only. | [optional] 
**Protection** | Pointer to [**AnyOfmicrosoftGraphWorkbookFormatProtection**](anyOf&lt;microsoft.graph.workbookFormatProtection&gt;.md) | Returns the format protection object for a range. Read-only. | [optional] 

## Methods

### NewMicrosoftGraphWorkbookRangeFormat

`func NewMicrosoftGraphWorkbookRangeFormat() *MicrosoftGraphWorkbookRangeFormat`

NewMicrosoftGraphWorkbookRangeFormat instantiates a new MicrosoftGraphWorkbookRangeFormat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphWorkbookRangeFormatWithDefaults

`func NewMicrosoftGraphWorkbookRangeFormatWithDefaults() *MicrosoftGraphWorkbookRangeFormat`

NewMicrosoftGraphWorkbookRangeFormatWithDefaults instantiates a new MicrosoftGraphWorkbookRangeFormat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphWorkbookRangeFormat) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphWorkbookRangeFormat) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphWorkbookRangeFormat) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphWorkbookRangeFormat) HasId() bool`

HasId returns a boolean if a field has been set.

### GetColumnWidth

`func (o *MicrosoftGraphWorkbookRangeFormat) GetColumnWidth() AnyOfnumberstringstring`

GetColumnWidth returns the ColumnWidth field if non-nil, zero value otherwise.

### GetColumnWidthOk

`func (o *MicrosoftGraphWorkbookRangeFormat) GetColumnWidthOk() (*AnyOfnumberstringstring, bool)`

GetColumnWidthOk returns a tuple with the ColumnWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumnWidth

`func (o *MicrosoftGraphWorkbookRangeFormat) SetColumnWidth(v AnyOfnumberstringstring)`

SetColumnWidth sets ColumnWidth field to given value.

### HasColumnWidth

`func (o *MicrosoftGraphWorkbookRangeFormat) HasColumnWidth() bool`

HasColumnWidth returns a boolean if a field has been set.

### SetColumnWidthNil

`func (o *MicrosoftGraphWorkbookRangeFormat) SetColumnWidthNil(b bool)`

 SetColumnWidthNil sets the value for ColumnWidth to be an explicit nil

### UnsetColumnWidth
`func (o *MicrosoftGraphWorkbookRangeFormat) UnsetColumnWidth()`

UnsetColumnWidth ensures that no value is present for ColumnWidth, not even an explicit nil
### GetHorizontalAlignment

`func (o *MicrosoftGraphWorkbookRangeFormat) GetHorizontalAlignment() string`

GetHorizontalAlignment returns the HorizontalAlignment field if non-nil, zero value otherwise.

### GetHorizontalAlignmentOk

`func (o *MicrosoftGraphWorkbookRangeFormat) GetHorizontalAlignmentOk() (*string, bool)`

GetHorizontalAlignmentOk returns a tuple with the HorizontalAlignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHorizontalAlignment

`func (o *MicrosoftGraphWorkbookRangeFormat) SetHorizontalAlignment(v string)`

SetHorizontalAlignment sets HorizontalAlignment field to given value.

### HasHorizontalAlignment

`func (o *MicrosoftGraphWorkbookRangeFormat) HasHorizontalAlignment() bool`

HasHorizontalAlignment returns a boolean if a field has been set.

### SetHorizontalAlignmentNil

`func (o *MicrosoftGraphWorkbookRangeFormat) SetHorizontalAlignmentNil(b bool)`

 SetHorizontalAlignmentNil sets the value for HorizontalAlignment to be an explicit nil

### UnsetHorizontalAlignment
`func (o *MicrosoftGraphWorkbookRangeFormat) UnsetHorizontalAlignment()`

UnsetHorizontalAlignment ensures that no value is present for HorizontalAlignment, not even an explicit nil
### GetRowHeight

`func (o *MicrosoftGraphWorkbookRangeFormat) GetRowHeight() AnyOfnumberstringstring`

GetRowHeight returns the RowHeight field if non-nil, zero value otherwise.

### GetRowHeightOk

`func (o *MicrosoftGraphWorkbookRangeFormat) GetRowHeightOk() (*AnyOfnumberstringstring, bool)`

GetRowHeightOk returns a tuple with the RowHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRowHeight

`func (o *MicrosoftGraphWorkbookRangeFormat) SetRowHeight(v AnyOfnumberstringstring)`

SetRowHeight sets RowHeight field to given value.

### HasRowHeight

`func (o *MicrosoftGraphWorkbookRangeFormat) HasRowHeight() bool`

HasRowHeight returns a boolean if a field has been set.

### SetRowHeightNil

`func (o *MicrosoftGraphWorkbookRangeFormat) SetRowHeightNil(b bool)`

 SetRowHeightNil sets the value for RowHeight to be an explicit nil

### UnsetRowHeight
`func (o *MicrosoftGraphWorkbookRangeFormat) UnsetRowHeight()`

UnsetRowHeight ensures that no value is present for RowHeight, not even an explicit nil
### GetVerticalAlignment

`func (o *MicrosoftGraphWorkbookRangeFormat) GetVerticalAlignment() string`

GetVerticalAlignment returns the VerticalAlignment field if non-nil, zero value otherwise.

### GetVerticalAlignmentOk

`func (o *MicrosoftGraphWorkbookRangeFormat) GetVerticalAlignmentOk() (*string, bool)`

GetVerticalAlignmentOk returns a tuple with the VerticalAlignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerticalAlignment

`func (o *MicrosoftGraphWorkbookRangeFormat) SetVerticalAlignment(v string)`

SetVerticalAlignment sets VerticalAlignment field to given value.

### HasVerticalAlignment

`func (o *MicrosoftGraphWorkbookRangeFormat) HasVerticalAlignment() bool`

HasVerticalAlignment returns a boolean if a field has been set.

### SetVerticalAlignmentNil

`func (o *MicrosoftGraphWorkbookRangeFormat) SetVerticalAlignmentNil(b bool)`

 SetVerticalAlignmentNil sets the value for VerticalAlignment to be an explicit nil

### UnsetVerticalAlignment
`func (o *MicrosoftGraphWorkbookRangeFormat) UnsetVerticalAlignment()`

UnsetVerticalAlignment ensures that no value is present for VerticalAlignment, not even an explicit nil
### GetWrapText

`func (o *MicrosoftGraphWorkbookRangeFormat) GetWrapText() bool`

GetWrapText returns the WrapText field if non-nil, zero value otherwise.

### GetWrapTextOk

`func (o *MicrosoftGraphWorkbookRangeFormat) GetWrapTextOk() (*bool, bool)`

GetWrapTextOk returns a tuple with the WrapText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWrapText

`func (o *MicrosoftGraphWorkbookRangeFormat) SetWrapText(v bool)`

SetWrapText sets WrapText field to given value.

### HasWrapText

`func (o *MicrosoftGraphWorkbookRangeFormat) HasWrapText() bool`

HasWrapText returns a boolean if a field has been set.

### SetWrapTextNil

`func (o *MicrosoftGraphWorkbookRangeFormat) SetWrapTextNil(b bool)`

 SetWrapTextNil sets the value for WrapText to be an explicit nil

### UnsetWrapText
`func (o *MicrosoftGraphWorkbookRangeFormat) UnsetWrapText()`

UnsetWrapText ensures that no value is present for WrapText, not even an explicit nil
### GetBorders

`func (o *MicrosoftGraphWorkbookRangeFormat) GetBorders() []MicrosoftGraphWorkbookRangeBorder`

GetBorders returns the Borders field if non-nil, zero value otherwise.

### GetBordersOk

`func (o *MicrosoftGraphWorkbookRangeFormat) GetBordersOk() (*[]MicrosoftGraphWorkbookRangeBorder, bool)`

GetBordersOk returns a tuple with the Borders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBorders

`func (o *MicrosoftGraphWorkbookRangeFormat) SetBorders(v []MicrosoftGraphWorkbookRangeBorder)`

SetBorders sets Borders field to given value.

### HasBorders

`func (o *MicrosoftGraphWorkbookRangeFormat) HasBorders() bool`

HasBorders returns a boolean if a field has been set.

### GetFill

`func (o *MicrosoftGraphWorkbookRangeFormat) GetFill() AnyOfmicrosoftGraphWorkbookRangeFill`

GetFill returns the Fill field if non-nil, zero value otherwise.

### GetFillOk

`func (o *MicrosoftGraphWorkbookRangeFormat) GetFillOk() (*AnyOfmicrosoftGraphWorkbookRangeFill, bool)`

GetFillOk returns a tuple with the Fill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFill

`func (o *MicrosoftGraphWorkbookRangeFormat) SetFill(v AnyOfmicrosoftGraphWorkbookRangeFill)`

SetFill sets Fill field to given value.

### HasFill

`func (o *MicrosoftGraphWorkbookRangeFormat) HasFill() bool`

HasFill returns a boolean if a field has been set.

### SetFillNil

`func (o *MicrosoftGraphWorkbookRangeFormat) SetFillNil(b bool)`

 SetFillNil sets the value for Fill to be an explicit nil

### UnsetFill
`func (o *MicrosoftGraphWorkbookRangeFormat) UnsetFill()`

UnsetFill ensures that no value is present for Fill, not even an explicit nil
### GetFont

`func (o *MicrosoftGraphWorkbookRangeFormat) GetFont() AnyOfmicrosoftGraphWorkbookRangeFont`

GetFont returns the Font field if non-nil, zero value otherwise.

### GetFontOk

`func (o *MicrosoftGraphWorkbookRangeFormat) GetFontOk() (*AnyOfmicrosoftGraphWorkbookRangeFont, bool)`

GetFontOk returns a tuple with the Font field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFont

`func (o *MicrosoftGraphWorkbookRangeFormat) SetFont(v AnyOfmicrosoftGraphWorkbookRangeFont)`

SetFont sets Font field to given value.

### HasFont

`func (o *MicrosoftGraphWorkbookRangeFormat) HasFont() bool`

HasFont returns a boolean if a field has been set.

### SetFontNil

`func (o *MicrosoftGraphWorkbookRangeFormat) SetFontNil(b bool)`

 SetFontNil sets the value for Font to be an explicit nil

### UnsetFont
`func (o *MicrosoftGraphWorkbookRangeFormat) UnsetFont()`

UnsetFont ensures that no value is present for Font, not even an explicit nil
### GetProtection

`func (o *MicrosoftGraphWorkbookRangeFormat) GetProtection() AnyOfmicrosoftGraphWorkbookFormatProtection`

GetProtection returns the Protection field if non-nil, zero value otherwise.

### GetProtectionOk

`func (o *MicrosoftGraphWorkbookRangeFormat) GetProtectionOk() (*AnyOfmicrosoftGraphWorkbookFormatProtection, bool)`

GetProtectionOk returns a tuple with the Protection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtection

`func (o *MicrosoftGraphWorkbookRangeFormat) SetProtection(v AnyOfmicrosoftGraphWorkbookFormatProtection)`

SetProtection sets Protection field to given value.

### HasProtection

`func (o *MicrosoftGraphWorkbookRangeFormat) HasProtection() bool`

HasProtection returns a boolean if a field has been set.

### SetProtectionNil

`func (o *MicrosoftGraphWorkbookRangeFormat) SetProtectionNil(b bool)`

 SetProtectionNil sets the value for Protection to be an explicit nil

### UnsetProtection
`func (o *MicrosoftGraphWorkbookRangeFormat) UnsetProtection()`

UnsetProtection ensures that no value is present for Protection, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


