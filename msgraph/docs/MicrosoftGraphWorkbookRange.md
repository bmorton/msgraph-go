# MicrosoftGraphWorkbookRange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**Address** | Pointer to **NullableString** | Represents the range reference in A1-style. Address value will contain the Sheet reference (e.g. Sheet1!A1:B4). Read-only. | [optional] 
**AddressLocal** | Pointer to **NullableString** | Represents range reference for the specified range in the language of the user. Read-only. | [optional] 
**CellCount** | Pointer to **int32** | Number of cells in the range. Read-only. | [optional] 
**ColumnCount** | Pointer to **int32** | Represents the total number of columns in the range. Read-only. | [optional] 
**ColumnHidden** | Pointer to **NullableBool** | Represents if all columns of the current range are hidden. | [optional] 
**ColumnIndex** | Pointer to **int32** | Represents the column number of the first cell in the range. Zero-indexed. Read-only. | [optional] 
**Formulas** | Pointer to [**AnyOfobject**](anyOf&lt;object&gt;.md) | Represents the formula in A1-style notation. | [optional] 
**FormulasLocal** | Pointer to [**AnyOfobject**](anyOf&lt;object&gt;.md) | Represents the formula in A1-style notation, in the user&#39;s language and number-formatting locale.  For example, the English &#39;&#x3D;SUM(A1, 1.5)&#39; formula would become &#39;&#x3D;SUMME(A1; 1,5)&#39; in German. | [optional] 
**FormulasR1C1** | Pointer to [**AnyOfobject**](anyOf&lt;object&gt;.md) | Represents the formula in R1C1-style notation. | [optional] 
**Hidden** | Pointer to **NullableBool** | Represents if all cells of the current range are hidden. Read-only. | [optional] 
**NumberFormat** | Pointer to [**AnyOfobject**](anyOf&lt;object&gt;.md) | Represents Excel&#39;s number format code for the given cell. | [optional] 
**RowCount** | Pointer to **int32** | Returns the total number of rows in the range. Read-only. | [optional] 
**RowHidden** | Pointer to **NullableBool** | Represents if all rows of the current range are hidden. | [optional] 
**RowIndex** | Pointer to **int32** | Returns the row number of the first cell in the range. Zero-indexed. Read-only. | [optional] 
**Text** | Pointer to [**AnyOfobject**](anyOf&lt;object&gt;.md) | Text values of the specified range. The Text value will not depend on the cell width. The # sign substitution that happens in Excel UI will not affect the text value returned by the API. Read-only. | [optional] 
**Values** | Pointer to [**AnyOfobject**](anyOf&lt;object&gt;.md) | Represents the raw values of the specified range. The data returned could be of type string, number, or a boolean. Cell that contain an error will return the error string. | [optional] 
**ValueTypes** | Pointer to [**AnyOfobject**](anyOf&lt;object&gt;.md) | Represents the type of data of each cell. The possible values are: Unknown, Empty, String, Integer, Double, Boolean, Error. Read-only. | [optional] 
**Format** | Pointer to [**AnyOfmicrosoftGraphWorkbookRangeFormat**](anyOf&lt;microsoft.graph.workbookRangeFormat&gt;.md) | Returns a format object, encapsulating the range&#39;s font, fill, borders, alignment, and other properties. Read-only. | [optional] 
**Sort** | Pointer to [**AnyOfmicrosoftGraphWorkbookRangeSort**](anyOf&lt;microsoft.graph.workbookRangeSort&gt;.md) | The worksheet containing the current range. Read-only. | [optional] 
**Worksheet** | Pointer to [**AnyOfmicrosoftGraphWorkbookWorksheet**](anyOf&lt;microsoft.graph.workbookWorksheet&gt;.md) | The worksheet containing the current range. Read-only. | [optional] 

## Methods

### NewMicrosoftGraphWorkbookRange

`func NewMicrosoftGraphWorkbookRange() *MicrosoftGraphWorkbookRange`

NewMicrosoftGraphWorkbookRange instantiates a new MicrosoftGraphWorkbookRange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphWorkbookRangeWithDefaults

`func NewMicrosoftGraphWorkbookRangeWithDefaults() *MicrosoftGraphWorkbookRange`

NewMicrosoftGraphWorkbookRangeWithDefaults instantiates a new MicrosoftGraphWorkbookRange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphWorkbookRange) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphWorkbookRange) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphWorkbookRange) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphWorkbookRange) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAddress

`func (o *MicrosoftGraphWorkbookRange) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *MicrosoftGraphWorkbookRange) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *MicrosoftGraphWorkbookRange) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *MicrosoftGraphWorkbookRange) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### SetAddressNil

`func (o *MicrosoftGraphWorkbookRange) SetAddressNil(b bool)`

 SetAddressNil sets the value for Address to be an explicit nil

### UnsetAddress
`func (o *MicrosoftGraphWorkbookRange) UnsetAddress()`

UnsetAddress ensures that no value is present for Address, not even an explicit nil
### GetAddressLocal

`func (o *MicrosoftGraphWorkbookRange) GetAddressLocal() string`

GetAddressLocal returns the AddressLocal field if non-nil, zero value otherwise.

### GetAddressLocalOk

`func (o *MicrosoftGraphWorkbookRange) GetAddressLocalOk() (*string, bool)`

GetAddressLocalOk returns a tuple with the AddressLocal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLocal

`func (o *MicrosoftGraphWorkbookRange) SetAddressLocal(v string)`

SetAddressLocal sets AddressLocal field to given value.

### HasAddressLocal

`func (o *MicrosoftGraphWorkbookRange) HasAddressLocal() bool`

HasAddressLocal returns a boolean if a field has been set.

### SetAddressLocalNil

`func (o *MicrosoftGraphWorkbookRange) SetAddressLocalNil(b bool)`

 SetAddressLocalNil sets the value for AddressLocal to be an explicit nil

### UnsetAddressLocal
`func (o *MicrosoftGraphWorkbookRange) UnsetAddressLocal()`

UnsetAddressLocal ensures that no value is present for AddressLocal, not even an explicit nil
### GetCellCount

`func (o *MicrosoftGraphWorkbookRange) GetCellCount() int32`

GetCellCount returns the CellCount field if non-nil, zero value otherwise.

### GetCellCountOk

`func (o *MicrosoftGraphWorkbookRange) GetCellCountOk() (*int32, bool)`

GetCellCountOk returns a tuple with the CellCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCellCount

`func (o *MicrosoftGraphWorkbookRange) SetCellCount(v int32)`

SetCellCount sets CellCount field to given value.

### HasCellCount

`func (o *MicrosoftGraphWorkbookRange) HasCellCount() bool`

HasCellCount returns a boolean if a field has been set.

### GetColumnCount

`func (o *MicrosoftGraphWorkbookRange) GetColumnCount() int32`

GetColumnCount returns the ColumnCount field if non-nil, zero value otherwise.

### GetColumnCountOk

`func (o *MicrosoftGraphWorkbookRange) GetColumnCountOk() (*int32, bool)`

GetColumnCountOk returns a tuple with the ColumnCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumnCount

`func (o *MicrosoftGraphWorkbookRange) SetColumnCount(v int32)`

SetColumnCount sets ColumnCount field to given value.

### HasColumnCount

`func (o *MicrosoftGraphWorkbookRange) HasColumnCount() bool`

HasColumnCount returns a boolean if a field has been set.

### GetColumnHidden

`func (o *MicrosoftGraphWorkbookRange) GetColumnHidden() bool`

GetColumnHidden returns the ColumnHidden field if non-nil, zero value otherwise.

### GetColumnHiddenOk

`func (o *MicrosoftGraphWorkbookRange) GetColumnHiddenOk() (*bool, bool)`

GetColumnHiddenOk returns a tuple with the ColumnHidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumnHidden

`func (o *MicrosoftGraphWorkbookRange) SetColumnHidden(v bool)`

SetColumnHidden sets ColumnHidden field to given value.

### HasColumnHidden

`func (o *MicrosoftGraphWorkbookRange) HasColumnHidden() bool`

HasColumnHidden returns a boolean if a field has been set.

### SetColumnHiddenNil

`func (o *MicrosoftGraphWorkbookRange) SetColumnHiddenNil(b bool)`

 SetColumnHiddenNil sets the value for ColumnHidden to be an explicit nil

### UnsetColumnHidden
`func (o *MicrosoftGraphWorkbookRange) UnsetColumnHidden()`

UnsetColumnHidden ensures that no value is present for ColumnHidden, not even an explicit nil
### GetColumnIndex

`func (o *MicrosoftGraphWorkbookRange) GetColumnIndex() int32`

GetColumnIndex returns the ColumnIndex field if non-nil, zero value otherwise.

### GetColumnIndexOk

`func (o *MicrosoftGraphWorkbookRange) GetColumnIndexOk() (*int32, bool)`

GetColumnIndexOk returns a tuple with the ColumnIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumnIndex

`func (o *MicrosoftGraphWorkbookRange) SetColumnIndex(v int32)`

SetColumnIndex sets ColumnIndex field to given value.

### HasColumnIndex

`func (o *MicrosoftGraphWorkbookRange) HasColumnIndex() bool`

HasColumnIndex returns a boolean if a field has been set.

### GetFormulas

`func (o *MicrosoftGraphWorkbookRange) GetFormulas() AnyOfobject`

GetFormulas returns the Formulas field if non-nil, zero value otherwise.

### GetFormulasOk

`func (o *MicrosoftGraphWorkbookRange) GetFormulasOk() (*AnyOfobject, bool)`

GetFormulasOk returns a tuple with the Formulas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormulas

`func (o *MicrosoftGraphWorkbookRange) SetFormulas(v AnyOfobject)`

SetFormulas sets Formulas field to given value.

### HasFormulas

`func (o *MicrosoftGraphWorkbookRange) HasFormulas() bool`

HasFormulas returns a boolean if a field has been set.

### SetFormulasNil

`func (o *MicrosoftGraphWorkbookRange) SetFormulasNil(b bool)`

 SetFormulasNil sets the value for Formulas to be an explicit nil

### UnsetFormulas
`func (o *MicrosoftGraphWorkbookRange) UnsetFormulas()`

UnsetFormulas ensures that no value is present for Formulas, not even an explicit nil
### GetFormulasLocal

`func (o *MicrosoftGraphWorkbookRange) GetFormulasLocal() AnyOfobject`

GetFormulasLocal returns the FormulasLocal field if non-nil, zero value otherwise.

### GetFormulasLocalOk

`func (o *MicrosoftGraphWorkbookRange) GetFormulasLocalOk() (*AnyOfobject, bool)`

GetFormulasLocalOk returns a tuple with the FormulasLocal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormulasLocal

`func (o *MicrosoftGraphWorkbookRange) SetFormulasLocal(v AnyOfobject)`

SetFormulasLocal sets FormulasLocal field to given value.

### HasFormulasLocal

`func (o *MicrosoftGraphWorkbookRange) HasFormulasLocal() bool`

HasFormulasLocal returns a boolean if a field has been set.

### SetFormulasLocalNil

`func (o *MicrosoftGraphWorkbookRange) SetFormulasLocalNil(b bool)`

 SetFormulasLocalNil sets the value for FormulasLocal to be an explicit nil

### UnsetFormulasLocal
`func (o *MicrosoftGraphWorkbookRange) UnsetFormulasLocal()`

UnsetFormulasLocal ensures that no value is present for FormulasLocal, not even an explicit nil
### GetFormulasR1C1

`func (o *MicrosoftGraphWorkbookRange) GetFormulasR1C1() AnyOfobject`

GetFormulasR1C1 returns the FormulasR1C1 field if non-nil, zero value otherwise.

### GetFormulasR1C1Ok

`func (o *MicrosoftGraphWorkbookRange) GetFormulasR1C1Ok() (*AnyOfobject, bool)`

GetFormulasR1C1Ok returns a tuple with the FormulasR1C1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormulasR1C1

`func (o *MicrosoftGraphWorkbookRange) SetFormulasR1C1(v AnyOfobject)`

SetFormulasR1C1 sets FormulasR1C1 field to given value.

### HasFormulasR1C1

`func (o *MicrosoftGraphWorkbookRange) HasFormulasR1C1() bool`

HasFormulasR1C1 returns a boolean if a field has been set.

### SetFormulasR1C1Nil

`func (o *MicrosoftGraphWorkbookRange) SetFormulasR1C1Nil(b bool)`

 SetFormulasR1C1Nil sets the value for FormulasR1C1 to be an explicit nil

### UnsetFormulasR1C1
`func (o *MicrosoftGraphWorkbookRange) UnsetFormulasR1C1()`

UnsetFormulasR1C1 ensures that no value is present for FormulasR1C1, not even an explicit nil
### GetHidden

`func (o *MicrosoftGraphWorkbookRange) GetHidden() bool`

GetHidden returns the Hidden field if non-nil, zero value otherwise.

### GetHiddenOk

`func (o *MicrosoftGraphWorkbookRange) GetHiddenOk() (*bool, bool)`

GetHiddenOk returns a tuple with the Hidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidden

`func (o *MicrosoftGraphWorkbookRange) SetHidden(v bool)`

SetHidden sets Hidden field to given value.

### HasHidden

`func (o *MicrosoftGraphWorkbookRange) HasHidden() bool`

HasHidden returns a boolean if a field has been set.

### SetHiddenNil

`func (o *MicrosoftGraphWorkbookRange) SetHiddenNil(b bool)`

 SetHiddenNil sets the value for Hidden to be an explicit nil

### UnsetHidden
`func (o *MicrosoftGraphWorkbookRange) UnsetHidden()`

UnsetHidden ensures that no value is present for Hidden, not even an explicit nil
### GetNumberFormat

`func (o *MicrosoftGraphWorkbookRange) GetNumberFormat() AnyOfobject`

GetNumberFormat returns the NumberFormat field if non-nil, zero value otherwise.

### GetNumberFormatOk

`func (o *MicrosoftGraphWorkbookRange) GetNumberFormatOk() (*AnyOfobject, bool)`

GetNumberFormatOk returns a tuple with the NumberFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberFormat

`func (o *MicrosoftGraphWorkbookRange) SetNumberFormat(v AnyOfobject)`

SetNumberFormat sets NumberFormat field to given value.

### HasNumberFormat

`func (o *MicrosoftGraphWorkbookRange) HasNumberFormat() bool`

HasNumberFormat returns a boolean if a field has been set.

### SetNumberFormatNil

`func (o *MicrosoftGraphWorkbookRange) SetNumberFormatNil(b bool)`

 SetNumberFormatNil sets the value for NumberFormat to be an explicit nil

### UnsetNumberFormat
`func (o *MicrosoftGraphWorkbookRange) UnsetNumberFormat()`

UnsetNumberFormat ensures that no value is present for NumberFormat, not even an explicit nil
### GetRowCount

`func (o *MicrosoftGraphWorkbookRange) GetRowCount() int32`

GetRowCount returns the RowCount field if non-nil, zero value otherwise.

### GetRowCountOk

`func (o *MicrosoftGraphWorkbookRange) GetRowCountOk() (*int32, bool)`

GetRowCountOk returns a tuple with the RowCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRowCount

`func (o *MicrosoftGraphWorkbookRange) SetRowCount(v int32)`

SetRowCount sets RowCount field to given value.

### HasRowCount

`func (o *MicrosoftGraphWorkbookRange) HasRowCount() bool`

HasRowCount returns a boolean if a field has been set.

### GetRowHidden

`func (o *MicrosoftGraphWorkbookRange) GetRowHidden() bool`

GetRowHidden returns the RowHidden field if non-nil, zero value otherwise.

### GetRowHiddenOk

`func (o *MicrosoftGraphWorkbookRange) GetRowHiddenOk() (*bool, bool)`

GetRowHiddenOk returns a tuple with the RowHidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRowHidden

`func (o *MicrosoftGraphWorkbookRange) SetRowHidden(v bool)`

SetRowHidden sets RowHidden field to given value.

### HasRowHidden

`func (o *MicrosoftGraphWorkbookRange) HasRowHidden() bool`

HasRowHidden returns a boolean if a field has been set.

### SetRowHiddenNil

`func (o *MicrosoftGraphWorkbookRange) SetRowHiddenNil(b bool)`

 SetRowHiddenNil sets the value for RowHidden to be an explicit nil

### UnsetRowHidden
`func (o *MicrosoftGraphWorkbookRange) UnsetRowHidden()`

UnsetRowHidden ensures that no value is present for RowHidden, not even an explicit nil
### GetRowIndex

`func (o *MicrosoftGraphWorkbookRange) GetRowIndex() int32`

GetRowIndex returns the RowIndex field if non-nil, zero value otherwise.

### GetRowIndexOk

`func (o *MicrosoftGraphWorkbookRange) GetRowIndexOk() (*int32, bool)`

GetRowIndexOk returns a tuple with the RowIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRowIndex

`func (o *MicrosoftGraphWorkbookRange) SetRowIndex(v int32)`

SetRowIndex sets RowIndex field to given value.

### HasRowIndex

`func (o *MicrosoftGraphWorkbookRange) HasRowIndex() bool`

HasRowIndex returns a boolean if a field has been set.

### GetText

`func (o *MicrosoftGraphWorkbookRange) GetText() AnyOfobject`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *MicrosoftGraphWorkbookRange) GetTextOk() (*AnyOfobject, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *MicrosoftGraphWorkbookRange) SetText(v AnyOfobject)`

SetText sets Text field to given value.

### HasText

`func (o *MicrosoftGraphWorkbookRange) HasText() bool`

HasText returns a boolean if a field has been set.

### SetTextNil

`func (o *MicrosoftGraphWorkbookRange) SetTextNil(b bool)`

 SetTextNil sets the value for Text to be an explicit nil

### UnsetText
`func (o *MicrosoftGraphWorkbookRange) UnsetText()`

UnsetText ensures that no value is present for Text, not even an explicit nil
### GetValues

`func (o *MicrosoftGraphWorkbookRange) GetValues() AnyOfobject`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *MicrosoftGraphWorkbookRange) GetValuesOk() (*AnyOfobject, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *MicrosoftGraphWorkbookRange) SetValues(v AnyOfobject)`

SetValues sets Values field to given value.

### HasValues

`func (o *MicrosoftGraphWorkbookRange) HasValues() bool`

HasValues returns a boolean if a field has been set.

### SetValuesNil

`func (o *MicrosoftGraphWorkbookRange) SetValuesNil(b bool)`

 SetValuesNil sets the value for Values to be an explicit nil

### UnsetValues
`func (o *MicrosoftGraphWorkbookRange) UnsetValues()`

UnsetValues ensures that no value is present for Values, not even an explicit nil
### GetValueTypes

`func (o *MicrosoftGraphWorkbookRange) GetValueTypes() AnyOfobject`

GetValueTypes returns the ValueTypes field if non-nil, zero value otherwise.

### GetValueTypesOk

`func (o *MicrosoftGraphWorkbookRange) GetValueTypesOk() (*AnyOfobject, bool)`

GetValueTypesOk returns a tuple with the ValueTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueTypes

`func (o *MicrosoftGraphWorkbookRange) SetValueTypes(v AnyOfobject)`

SetValueTypes sets ValueTypes field to given value.

### HasValueTypes

`func (o *MicrosoftGraphWorkbookRange) HasValueTypes() bool`

HasValueTypes returns a boolean if a field has been set.

### SetValueTypesNil

`func (o *MicrosoftGraphWorkbookRange) SetValueTypesNil(b bool)`

 SetValueTypesNil sets the value for ValueTypes to be an explicit nil

### UnsetValueTypes
`func (o *MicrosoftGraphWorkbookRange) UnsetValueTypes()`

UnsetValueTypes ensures that no value is present for ValueTypes, not even an explicit nil
### GetFormat

`func (o *MicrosoftGraphWorkbookRange) GetFormat() AnyOfmicrosoftGraphWorkbookRangeFormat`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *MicrosoftGraphWorkbookRange) GetFormatOk() (*AnyOfmicrosoftGraphWorkbookRangeFormat, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *MicrosoftGraphWorkbookRange) SetFormat(v AnyOfmicrosoftGraphWorkbookRangeFormat)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *MicrosoftGraphWorkbookRange) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### SetFormatNil

`func (o *MicrosoftGraphWorkbookRange) SetFormatNil(b bool)`

 SetFormatNil sets the value for Format to be an explicit nil

### UnsetFormat
`func (o *MicrosoftGraphWorkbookRange) UnsetFormat()`

UnsetFormat ensures that no value is present for Format, not even an explicit nil
### GetSort

`func (o *MicrosoftGraphWorkbookRange) GetSort() AnyOfmicrosoftGraphWorkbookRangeSort`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *MicrosoftGraphWorkbookRange) GetSortOk() (*AnyOfmicrosoftGraphWorkbookRangeSort, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *MicrosoftGraphWorkbookRange) SetSort(v AnyOfmicrosoftGraphWorkbookRangeSort)`

SetSort sets Sort field to given value.

### HasSort

`func (o *MicrosoftGraphWorkbookRange) HasSort() bool`

HasSort returns a boolean if a field has been set.

### SetSortNil

`func (o *MicrosoftGraphWorkbookRange) SetSortNil(b bool)`

 SetSortNil sets the value for Sort to be an explicit nil

### UnsetSort
`func (o *MicrosoftGraphWorkbookRange) UnsetSort()`

UnsetSort ensures that no value is present for Sort, not even an explicit nil
### GetWorksheet

`func (o *MicrosoftGraphWorkbookRange) GetWorksheet() AnyOfmicrosoftGraphWorkbookWorksheet`

GetWorksheet returns the Worksheet field if non-nil, zero value otherwise.

### GetWorksheetOk

`func (o *MicrosoftGraphWorkbookRange) GetWorksheetOk() (*AnyOfmicrosoftGraphWorkbookWorksheet, bool)`

GetWorksheetOk returns a tuple with the Worksheet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorksheet

`func (o *MicrosoftGraphWorkbookRange) SetWorksheet(v AnyOfmicrosoftGraphWorkbookWorksheet)`

SetWorksheet sets Worksheet field to given value.

### HasWorksheet

`func (o *MicrosoftGraphWorkbookRange) HasWorksheet() bool`

HasWorksheet returns a boolean if a field has been set.

### SetWorksheetNil

`func (o *MicrosoftGraphWorkbookRange) SetWorksheetNil(b bool)`

 SetWorksheetNil sets the value for Worksheet to be an explicit nil

### UnsetWorksheet
`func (o *MicrosoftGraphWorkbookRange) UnsetWorksheet()`

UnsetWorksheet ensures that no value is present for Worksheet, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


