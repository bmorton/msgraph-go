# WorkbookRangeFormat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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

### NewWorkbookRangeFormat

`func NewWorkbookRangeFormat() *WorkbookRangeFormat`

NewWorkbookRangeFormat instantiates a new WorkbookRangeFormat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkbookRangeFormatWithDefaults

`func NewWorkbookRangeFormatWithDefaults() *WorkbookRangeFormat`

NewWorkbookRangeFormatWithDefaults instantiates a new WorkbookRangeFormat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetColumnWidth

`func (o *WorkbookRangeFormat) GetColumnWidth() AnyOfnumberstringstring`

GetColumnWidth returns the ColumnWidth field if non-nil, zero value otherwise.

### GetColumnWidthOk

`func (o *WorkbookRangeFormat) GetColumnWidthOk() (*AnyOfnumberstringstring, bool)`

GetColumnWidthOk returns a tuple with the ColumnWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumnWidth

`func (o *WorkbookRangeFormat) SetColumnWidth(v AnyOfnumberstringstring)`

SetColumnWidth sets ColumnWidth field to given value.

### HasColumnWidth

`func (o *WorkbookRangeFormat) HasColumnWidth() bool`

HasColumnWidth returns a boolean if a field has been set.

### SetColumnWidthNil

`func (o *WorkbookRangeFormat) SetColumnWidthNil(b bool)`

 SetColumnWidthNil sets the value for ColumnWidth to be an explicit nil

### UnsetColumnWidth
`func (o *WorkbookRangeFormat) UnsetColumnWidth()`

UnsetColumnWidth ensures that no value is present for ColumnWidth, not even an explicit nil
### GetHorizontalAlignment

`func (o *WorkbookRangeFormat) GetHorizontalAlignment() string`

GetHorizontalAlignment returns the HorizontalAlignment field if non-nil, zero value otherwise.

### GetHorizontalAlignmentOk

`func (o *WorkbookRangeFormat) GetHorizontalAlignmentOk() (*string, bool)`

GetHorizontalAlignmentOk returns a tuple with the HorizontalAlignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHorizontalAlignment

`func (o *WorkbookRangeFormat) SetHorizontalAlignment(v string)`

SetHorizontalAlignment sets HorizontalAlignment field to given value.

### HasHorizontalAlignment

`func (o *WorkbookRangeFormat) HasHorizontalAlignment() bool`

HasHorizontalAlignment returns a boolean if a field has been set.

### SetHorizontalAlignmentNil

`func (o *WorkbookRangeFormat) SetHorizontalAlignmentNil(b bool)`

 SetHorizontalAlignmentNil sets the value for HorizontalAlignment to be an explicit nil

### UnsetHorizontalAlignment
`func (o *WorkbookRangeFormat) UnsetHorizontalAlignment()`

UnsetHorizontalAlignment ensures that no value is present for HorizontalAlignment, not even an explicit nil
### GetRowHeight

`func (o *WorkbookRangeFormat) GetRowHeight() AnyOfnumberstringstring`

GetRowHeight returns the RowHeight field if non-nil, zero value otherwise.

### GetRowHeightOk

`func (o *WorkbookRangeFormat) GetRowHeightOk() (*AnyOfnumberstringstring, bool)`

GetRowHeightOk returns a tuple with the RowHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRowHeight

`func (o *WorkbookRangeFormat) SetRowHeight(v AnyOfnumberstringstring)`

SetRowHeight sets RowHeight field to given value.

### HasRowHeight

`func (o *WorkbookRangeFormat) HasRowHeight() bool`

HasRowHeight returns a boolean if a field has been set.

### SetRowHeightNil

`func (o *WorkbookRangeFormat) SetRowHeightNil(b bool)`

 SetRowHeightNil sets the value for RowHeight to be an explicit nil

### UnsetRowHeight
`func (o *WorkbookRangeFormat) UnsetRowHeight()`

UnsetRowHeight ensures that no value is present for RowHeight, not even an explicit nil
### GetVerticalAlignment

`func (o *WorkbookRangeFormat) GetVerticalAlignment() string`

GetVerticalAlignment returns the VerticalAlignment field if non-nil, zero value otherwise.

### GetVerticalAlignmentOk

`func (o *WorkbookRangeFormat) GetVerticalAlignmentOk() (*string, bool)`

GetVerticalAlignmentOk returns a tuple with the VerticalAlignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerticalAlignment

`func (o *WorkbookRangeFormat) SetVerticalAlignment(v string)`

SetVerticalAlignment sets VerticalAlignment field to given value.

### HasVerticalAlignment

`func (o *WorkbookRangeFormat) HasVerticalAlignment() bool`

HasVerticalAlignment returns a boolean if a field has been set.

### SetVerticalAlignmentNil

`func (o *WorkbookRangeFormat) SetVerticalAlignmentNil(b bool)`

 SetVerticalAlignmentNil sets the value for VerticalAlignment to be an explicit nil

### UnsetVerticalAlignment
`func (o *WorkbookRangeFormat) UnsetVerticalAlignment()`

UnsetVerticalAlignment ensures that no value is present for VerticalAlignment, not even an explicit nil
### GetWrapText

`func (o *WorkbookRangeFormat) GetWrapText() bool`

GetWrapText returns the WrapText field if non-nil, zero value otherwise.

### GetWrapTextOk

`func (o *WorkbookRangeFormat) GetWrapTextOk() (*bool, bool)`

GetWrapTextOk returns a tuple with the WrapText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWrapText

`func (o *WorkbookRangeFormat) SetWrapText(v bool)`

SetWrapText sets WrapText field to given value.

### HasWrapText

`func (o *WorkbookRangeFormat) HasWrapText() bool`

HasWrapText returns a boolean if a field has been set.

### SetWrapTextNil

`func (o *WorkbookRangeFormat) SetWrapTextNil(b bool)`

 SetWrapTextNil sets the value for WrapText to be an explicit nil

### UnsetWrapText
`func (o *WorkbookRangeFormat) UnsetWrapText()`

UnsetWrapText ensures that no value is present for WrapText, not even an explicit nil
### GetBorders

`func (o *WorkbookRangeFormat) GetBorders() []MicrosoftGraphWorkbookRangeBorder`

GetBorders returns the Borders field if non-nil, zero value otherwise.

### GetBordersOk

`func (o *WorkbookRangeFormat) GetBordersOk() (*[]MicrosoftGraphWorkbookRangeBorder, bool)`

GetBordersOk returns a tuple with the Borders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBorders

`func (o *WorkbookRangeFormat) SetBorders(v []MicrosoftGraphWorkbookRangeBorder)`

SetBorders sets Borders field to given value.

### HasBorders

`func (o *WorkbookRangeFormat) HasBorders() bool`

HasBorders returns a boolean if a field has been set.

### GetFill

`func (o *WorkbookRangeFormat) GetFill() AnyOfmicrosoftGraphWorkbookRangeFill`

GetFill returns the Fill field if non-nil, zero value otherwise.

### GetFillOk

`func (o *WorkbookRangeFormat) GetFillOk() (*AnyOfmicrosoftGraphWorkbookRangeFill, bool)`

GetFillOk returns a tuple with the Fill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFill

`func (o *WorkbookRangeFormat) SetFill(v AnyOfmicrosoftGraphWorkbookRangeFill)`

SetFill sets Fill field to given value.

### HasFill

`func (o *WorkbookRangeFormat) HasFill() bool`

HasFill returns a boolean if a field has been set.

### SetFillNil

`func (o *WorkbookRangeFormat) SetFillNil(b bool)`

 SetFillNil sets the value for Fill to be an explicit nil

### UnsetFill
`func (o *WorkbookRangeFormat) UnsetFill()`

UnsetFill ensures that no value is present for Fill, not even an explicit nil
### GetFont

`func (o *WorkbookRangeFormat) GetFont() AnyOfmicrosoftGraphWorkbookRangeFont`

GetFont returns the Font field if non-nil, zero value otherwise.

### GetFontOk

`func (o *WorkbookRangeFormat) GetFontOk() (*AnyOfmicrosoftGraphWorkbookRangeFont, bool)`

GetFontOk returns a tuple with the Font field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFont

`func (o *WorkbookRangeFormat) SetFont(v AnyOfmicrosoftGraphWorkbookRangeFont)`

SetFont sets Font field to given value.

### HasFont

`func (o *WorkbookRangeFormat) HasFont() bool`

HasFont returns a boolean if a field has been set.

### SetFontNil

`func (o *WorkbookRangeFormat) SetFontNil(b bool)`

 SetFontNil sets the value for Font to be an explicit nil

### UnsetFont
`func (o *WorkbookRangeFormat) UnsetFont()`

UnsetFont ensures that no value is present for Font, not even an explicit nil
### GetProtection

`func (o *WorkbookRangeFormat) GetProtection() AnyOfmicrosoftGraphWorkbookFormatProtection`

GetProtection returns the Protection field if non-nil, zero value otherwise.

### GetProtectionOk

`func (o *WorkbookRangeFormat) GetProtectionOk() (*AnyOfmicrosoftGraphWorkbookFormatProtection, bool)`

GetProtectionOk returns a tuple with the Protection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtection

`func (o *WorkbookRangeFormat) SetProtection(v AnyOfmicrosoftGraphWorkbookFormatProtection)`

SetProtection sets Protection field to given value.

### HasProtection

`func (o *WorkbookRangeFormat) HasProtection() bool`

HasProtection returns a boolean if a field has been set.

### SetProtectionNil

`func (o *WorkbookRangeFormat) SetProtectionNil(b bool)`

 SetProtectionNil sets the value for Protection to be an explicit nil

### UnsetProtection
`func (o *WorkbookRangeFormat) UnsetProtection()`

UnsetProtection ensures that no value is present for Protection, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


