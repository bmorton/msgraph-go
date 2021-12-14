# WorkbookRange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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

### NewWorkbookRange

`func NewWorkbookRange() *WorkbookRange`

NewWorkbookRange instantiates a new WorkbookRange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkbookRangeWithDefaults

`func NewWorkbookRangeWithDefaults() *WorkbookRange`

NewWorkbookRangeWithDefaults instantiates a new WorkbookRange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *WorkbookRange) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *WorkbookRange) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *WorkbookRange) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *WorkbookRange) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### SetAddressNil

`func (o *WorkbookRange) SetAddressNil(b bool)`

 SetAddressNil sets the value for Address to be an explicit nil

### UnsetAddress
`func (o *WorkbookRange) UnsetAddress()`

UnsetAddress ensures that no value is present for Address, not even an explicit nil
### GetAddressLocal

`func (o *WorkbookRange) GetAddressLocal() string`

GetAddressLocal returns the AddressLocal field if non-nil, zero value otherwise.

### GetAddressLocalOk

`func (o *WorkbookRange) GetAddressLocalOk() (*string, bool)`

GetAddressLocalOk returns a tuple with the AddressLocal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLocal

`func (o *WorkbookRange) SetAddressLocal(v string)`

SetAddressLocal sets AddressLocal field to given value.

### HasAddressLocal

`func (o *WorkbookRange) HasAddressLocal() bool`

HasAddressLocal returns a boolean if a field has been set.

### SetAddressLocalNil

`func (o *WorkbookRange) SetAddressLocalNil(b bool)`

 SetAddressLocalNil sets the value for AddressLocal to be an explicit nil

### UnsetAddressLocal
`func (o *WorkbookRange) UnsetAddressLocal()`

UnsetAddressLocal ensures that no value is present for AddressLocal, not even an explicit nil
### GetCellCount

`func (o *WorkbookRange) GetCellCount() int32`

GetCellCount returns the CellCount field if non-nil, zero value otherwise.

### GetCellCountOk

`func (o *WorkbookRange) GetCellCountOk() (*int32, bool)`

GetCellCountOk returns a tuple with the CellCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCellCount

`func (o *WorkbookRange) SetCellCount(v int32)`

SetCellCount sets CellCount field to given value.

### HasCellCount

`func (o *WorkbookRange) HasCellCount() bool`

HasCellCount returns a boolean if a field has been set.

### GetColumnCount

`func (o *WorkbookRange) GetColumnCount() int32`

GetColumnCount returns the ColumnCount field if non-nil, zero value otherwise.

### GetColumnCountOk

`func (o *WorkbookRange) GetColumnCountOk() (*int32, bool)`

GetColumnCountOk returns a tuple with the ColumnCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumnCount

`func (o *WorkbookRange) SetColumnCount(v int32)`

SetColumnCount sets ColumnCount field to given value.

### HasColumnCount

`func (o *WorkbookRange) HasColumnCount() bool`

HasColumnCount returns a boolean if a field has been set.

### GetColumnHidden

`func (o *WorkbookRange) GetColumnHidden() bool`

GetColumnHidden returns the ColumnHidden field if non-nil, zero value otherwise.

### GetColumnHiddenOk

`func (o *WorkbookRange) GetColumnHiddenOk() (*bool, bool)`

GetColumnHiddenOk returns a tuple with the ColumnHidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumnHidden

`func (o *WorkbookRange) SetColumnHidden(v bool)`

SetColumnHidden sets ColumnHidden field to given value.

### HasColumnHidden

`func (o *WorkbookRange) HasColumnHidden() bool`

HasColumnHidden returns a boolean if a field has been set.

### SetColumnHiddenNil

`func (o *WorkbookRange) SetColumnHiddenNil(b bool)`

 SetColumnHiddenNil sets the value for ColumnHidden to be an explicit nil

### UnsetColumnHidden
`func (o *WorkbookRange) UnsetColumnHidden()`

UnsetColumnHidden ensures that no value is present for ColumnHidden, not even an explicit nil
### GetColumnIndex

`func (o *WorkbookRange) GetColumnIndex() int32`

GetColumnIndex returns the ColumnIndex field if non-nil, zero value otherwise.

### GetColumnIndexOk

`func (o *WorkbookRange) GetColumnIndexOk() (*int32, bool)`

GetColumnIndexOk returns a tuple with the ColumnIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumnIndex

`func (o *WorkbookRange) SetColumnIndex(v int32)`

SetColumnIndex sets ColumnIndex field to given value.

### HasColumnIndex

`func (o *WorkbookRange) HasColumnIndex() bool`

HasColumnIndex returns a boolean if a field has been set.

### GetFormulas

`func (o *WorkbookRange) GetFormulas() AnyOfobject`

GetFormulas returns the Formulas field if non-nil, zero value otherwise.

### GetFormulasOk

`func (o *WorkbookRange) GetFormulasOk() (*AnyOfobject, bool)`

GetFormulasOk returns a tuple with the Formulas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormulas

`func (o *WorkbookRange) SetFormulas(v AnyOfobject)`

SetFormulas sets Formulas field to given value.

### HasFormulas

`func (o *WorkbookRange) HasFormulas() bool`

HasFormulas returns a boolean if a field has been set.

### SetFormulasNil

`func (o *WorkbookRange) SetFormulasNil(b bool)`

 SetFormulasNil sets the value for Formulas to be an explicit nil

### UnsetFormulas
`func (o *WorkbookRange) UnsetFormulas()`

UnsetFormulas ensures that no value is present for Formulas, not even an explicit nil
### GetFormulasLocal

`func (o *WorkbookRange) GetFormulasLocal() AnyOfobject`

GetFormulasLocal returns the FormulasLocal field if non-nil, zero value otherwise.

### GetFormulasLocalOk

`func (o *WorkbookRange) GetFormulasLocalOk() (*AnyOfobject, bool)`

GetFormulasLocalOk returns a tuple with the FormulasLocal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormulasLocal

`func (o *WorkbookRange) SetFormulasLocal(v AnyOfobject)`

SetFormulasLocal sets FormulasLocal field to given value.

### HasFormulasLocal

`func (o *WorkbookRange) HasFormulasLocal() bool`

HasFormulasLocal returns a boolean if a field has been set.

### SetFormulasLocalNil

`func (o *WorkbookRange) SetFormulasLocalNil(b bool)`

 SetFormulasLocalNil sets the value for FormulasLocal to be an explicit nil

### UnsetFormulasLocal
`func (o *WorkbookRange) UnsetFormulasLocal()`

UnsetFormulasLocal ensures that no value is present for FormulasLocal, not even an explicit nil
### GetFormulasR1C1

`func (o *WorkbookRange) GetFormulasR1C1() AnyOfobject`

GetFormulasR1C1 returns the FormulasR1C1 field if non-nil, zero value otherwise.

### GetFormulasR1C1Ok

`func (o *WorkbookRange) GetFormulasR1C1Ok() (*AnyOfobject, bool)`

GetFormulasR1C1Ok returns a tuple with the FormulasR1C1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormulasR1C1

`func (o *WorkbookRange) SetFormulasR1C1(v AnyOfobject)`

SetFormulasR1C1 sets FormulasR1C1 field to given value.

### HasFormulasR1C1

`func (o *WorkbookRange) HasFormulasR1C1() bool`

HasFormulasR1C1 returns a boolean if a field has been set.

### SetFormulasR1C1Nil

`func (o *WorkbookRange) SetFormulasR1C1Nil(b bool)`

 SetFormulasR1C1Nil sets the value for FormulasR1C1 to be an explicit nil

### UnsetFormulasR1C1
`func (o *WorkbookRange) UnsetFormulasR1C1()`

UnsetFormulasR1C1 ensures that no value is present for FormulasR1C1, not even an explicit nil
### GetHidden

`func (o *WorkbookRange) GetHidden() bool`

GetHidden returns the Hidden field if non-nil, zero value otherwise.

### GetHiddenOk

`func (o *WorkbookRange) GetHiddenOk() (*bool, bool)`

GetHiddenOk returns a tuple with the Hidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidden

`func (o *WorkbookRange) SetHidden(v bool)`

SetHidden sets Hidden field to given value.

### HasHidden

`func (o *WorkbookRange) HasHidden() bool`

HasHidden returns a boolean if a field has been set.

### SetHiddenNil

`func (o *WorkbookRange) SetHiddenNil(b bool)`

 SetHiddenNil sets the value for Hidden to be an explicit nil

### UnsetHidden
`func (o *WorkbookRange) UnsetHidden()`

UnsetHidden ensures that no value is present for Hidden, not even an explicit nil
### GetNumberFormat

`func (o *WorkbookRange) GetNumberFormat() AnyOfobject`

GetNumberFormat returns the NumberFormat field if non-nil, zero value otherwise.

### GetNumberFormatOk

`func (o *WorkbookRange) GetNumberFormatOk() (*AnyOfobject, bool)`

GetNumberFormatOk returns a tuple with the NumberFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberFormat

`func (o *WorkbookRange) SetNumberFormat(v AnyOfobject)`

SetNumberFormat sets NumberFormat field to given value.

### HasNumberFormat

`func (o *WorkbookRange) HasNumberFormat() bool`

HasNumberFormat returns a boolean if a field has been set.

### SetNumberFormatNil

`func (o *WorkbookRange) SetNumberFormatNil(b bool)`

 SetNumberFormatNil sets the value for NumberFormat to be an explicit nil

### UnsetNumberFormat
`func (o *WorkbookRange) UnsetNumberFormat()`

UnsetNumberFormat ensures that no value is present for NumberFormat, not even an explicit nil
### GetRowCount

`func (o *WorkbookRange) GetRowCount() int32`

GetRowCount returns the RowCount field if non-nil, zero value otherwise.

### GetRowCountOk

`func (o *WorkbookRange) GetRowCountOk() (*int32, bool)`

GetRowCountOk returns a tuple with the RowCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRowCount

`func (o *WorkbookRange) SetRowCount(v int32)`

SetRowCount sets RowCount field to given value.

### HasRowCount

`func (o *WorkbookRange) HasRowCount() bool`

HasRowCount returns a boolean if a field has been set.

### GetRowHidden

`func (o *WorkbookRange) GetRowHidden() bool`

GetRowHidden returns the RowHidden field if non-nil, zero value otherwise.

### GetRowHiddenOk

`func (o *WorkbookRange) GetRowHiddenOk() (*bool, bool)`

GetRowHiddenOk returns a tuple with the RowHidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRowHidden

`func (o *WorkbookRange) SetRowHidden(v bool)`

SetRowHidden sets RowHidden field to given value.

### HasRowHidden

`func (o *WorkbookRange) HasRowHidden() bool`

HasRowHidden returns a boolean if a field has been set.

### SetRowHiddenNil

`func (o *WorkbookRange) SetRowHiddenNil(b bool)`

 SetRowHiddenNil sets the value for RowHidden to be an explicit nil

### UnsetRowHidden
`func (o *WorkbookRange) UnsetRowHidden()`

UnsetRowHidden ensures that no value is present for RowHidden, not even an explicit nil
### GetRowIndex

`func (o *WorkbookRange) GetRowIndex() int32`

GetRowIndex returns the RowIndex field if non-nil, zero value otherwise.

### GetRowIndexOk

`func (o *WorkbookRange) GetRowIndexOk() (*int32, bool)`

GetRowIndexOk returns a tuple with the RowIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRowIndex

`func (o *WorkbookRange) SetRowIndex(v int32)`

SetRowIndex sets RowIndex field to given value.

### HasRowIndex

`func (o *WorkbookRange) HasRowIndex() bool`

HasRowIndex returns a boolean if a field has been set.

### GetText

`func (o *WorkbookRange) GetText() AnyOfobject`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *WorkbookRange) GetTextOk() (*AnyOfobject, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *WorkbookRange) SetText(v AnyOfobject)`

SetText sets Text field to given value.

### HasText

`func (o *WorkbookRange) HasText() bool`

HasText returns a boolean if a field has been set.

### SetTextNil

`func (o *WorkbookRange) SetTextNil(b bool)`

 SetTextNil sets the value for Text to be an explicit nil

### UnsetText
`func (o *WorkbookRange) UnsetText()`

UnsetText ensures that no value is present for Text, not even an explicit nil
### GetValues

`func (o *WorkbookRange) GetValues() AnyOfobject`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *WorkbookRange) GetValuesOk() (*AnyOfobject, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *WorkbookRange) SetValues(v AnyOfobject)`

SetValues sets Values field to given value.

### HasValues

`func (o *WorkbookRange) HasValues() bool`

HasValues returns a boolean if a field has been set.

### SetValuesNil

`func (o *WorkbookRange) SetValuesNil(b bool)`

 SetValuesNil sets the value for Values to be an explicit nil

### UnsetValues
`func (o *WorkbookRange) UnsetValues()`

UnsetValues ensures that no value is present for Values, not even an explicit nil
### GetValueTypes

`func (o *WorkbookRange) GetValueTypes() AnyOfobject`

GetValueTypes returns the ValueTypes field if non-nil, zero value otherwise.

### GetValueTypesOk

`func (o *WorkbookRange) GetValueTypesOk() (*AnyOfobject, bool)`

GetValueTypesOk returns a tuple with the ValueTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueTypes

`func (o *WorkbookRange) SetValueTypes(v AnyOfobject)`

SetValueTypes sets ValueTypes field to given value.

### HasValueTypes

`func (o *WorkbookRange) HasValueTypes() bool`

HasValueTypes returns a boolean if a field has been set.

### SetValueTypesNil

`func (o *WorkbookRange) SetValueTypesNil(b bool)`

 SetValueTypesNil sets the value for ValueTypes to be an explicit nil

### UnsetValueTypes
`func (o *WorkbookRange) UnsetValueTypes()`

UnsetValueTypes ensures that no value is present for ValueTypes, not even an explicit nil
### GetFormat

`func (o *WorkbookRange) GetFormat() AnyOfmicrosoftGraphWorkbookRangeFormat`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *WorkbookRange) GetFormatOk() (*AnyOfmicrosoftGraphWorkbookRangeFormat, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *WorkbookRange) SetFormat(v AnyOfmicrosoftGraphWorkbookRangeFormat)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *WorkbookRange) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### SetFormatNil

`func (o *WorkbookRange) SetFormatNil(b bool)`

 SetFormatNil sets the value for Format to be an explicit nil

### UnsetFormat
`func (o *WorkbookRange) UnsetFormat()`

UnsetFormat ensures that no value is present for Format, not even an explicit nil
### GetSort

`func (o *WorkbookRange) GetSort() AnyOfmicrosoftGraphWorkbookRangeSort`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *WorkbookRange) GetSortOk() (*AnyOfmicrosoftGraphWorkbookRangeSort, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *WorkbookRange) SetSort(v AnyOfmicrosoftGraphWorkbookRangeSort)`

SetSort sets Sort field to given value.

### HasSort

`func (o *WorkbookRange) HasSort() bool`

HasSort returns a boolean if a field has been set.

### SetSortNil

`func (o *WorkbookRange) SetSortNil(b bool)`

 SetSortNil sets the value for Sort to be an explicit nil

### UnsetSort
`func (o *WorkbookRange) UnsetSort()`

UnsetSort ensures that no value is present for Sort, not even an explicit nil
### GetWorksheet

`func (o *WorkbookRange) GetWorksheet() AnyOfmicrosoftGraphWorkbookWorksheet`

GetWorksheet returns the Worksheet field if non-nil, zero value otherwise.

### GetWorksheetOk

`func (o *WorkbookRange) GetWorksheetOk() (*AnyOfmicrosoftGraphWorkbookWorksheet, bool)`

GetWorksheetOk returns a tuple with the Worksheet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorksheet

`func (o *WorkbookRange) SetWorksheet(v AnyOfmicrosoftGraphWorkbookWorksheet)`

SetWorksheet sets Worksheet field to given value.

### HasWorksheet

`func (o *WorkbookRange) HasWorksheet() bool`

HasWorksheet returns a boolean if a field has been set.

### SetWorksheetNil

`func (o *WorkbookRange) SetWorksheetNil(b bool)`

 SetWorksheetNil sets the value for Worksheet to be an explicit nil

### UnsetWorksheet
`func (o *WorkbookRange) UnsetWorksheet()`

UnsetWorksheet ensures that no value is present for Worksheet, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


