# MicrosoftGraphWorkbookRangeView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**CellAddresses** | Pointer to [**AnyOfobject**](anyOf&lt;object&gt;.md) | Represents the cell addresses | [optional] 
**ColumnCount** | Pointer to **int32** | Returns the number of visible columns. Read-only. | [optional] 
**Formulas** | Pointer to [**AnyOfobject**](anyOf&lt;object&gt;.md) | Represents the formula in A1-style notation. | [optional] 
**FormulasLocal** | Pointer to [**AnyOfobject**](anyOf&lt;object&gt;.md) | Represents the formula in A1-style notation, in the user&#39;s language and number-formatting locale. For example, the English &#39;&#x3D;SUM(A1, 1.5)&#39; formula would become &#39;&#x3D;SUMME(A1; 1,5)&#39; in German. | [optional] 
**FormulasR1C1** | Pointer to [**AnyOfobject**](anyOf&lt;object&gt;.md) | Represents the formula in R1C1-style notation. | [optional] 
**Index** | Pointer to **int32** | Index of the range. | [optional] 
**NumberFormat** | Pointer to [**AnyOfobject**](anyOf&lt;object&gt;.md) | Represents Excel&#39;s number format code for the given cell. Read-only. | [optional] 
**RowCount** | Pointer to **int32** | Returns the number of visible rows. Read-only. | [optional] 
**Text** | Pointer to [**AnyOfobject**](anyOf&lt;object&gt;.md) | Text values of the specified range. The Text value will not depend on the cell width. The # sign substitution that happens in Excel UI will not affect the text value returned by the API. Read-only. | [optional] 
**Values** | Pointer to [**AnyOfobject**](anyOf&lt;object&gt;.md) | Represents the raw values of the specified range view. The data returned could be of type string, number, or a boolean. Cell that contain an error will return the error string. | [optional] 
**ValueTypes** | Pointer to [**AnyOfobject**](anyOf&lt;object&gt;.md) | Represents the type of data of each cell. Read-only. The possible values are: Unknown, Empty, String, Integer, Double, Boolean, Error. | [optional] 
**Rows** | Pointer to [**[]MicrosoftGraphWorkbookRangeView**](MicrosoftGraphWorkbookRangeView.md) | Represents a collection of range views associated with the range. Read-only. Read-only. | [optional] 

## Methods

### NewMicrosoftGraphWorkbookRangeView

`func NewMicrosoftGraphWorkbookRangeView() *MicrosoftGraphWorkbookRangeView`

NewMicrosoftGraphWorkbookRangeView instantiates a new MicrosoftGraphWorkbookRangeView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphWorkbookRangeViewWithDefaults

`func NewMicrosoftGraphWorkbookRangeViewWithDefaults() *MicrosoftGraphWorkbookRangeView`

NewMicrosoftGraphWorkbookRangeViewWithDefaults instantiates a new MicrosoftGraphWorkbookRangeView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphWorkbookRangeView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphWorkbookRangeView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphWorkbookRangeView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphWorkbookRangeView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCellAddresses

`func (o *MicrosoftGraphWorkbookRangeView) GetCellAddresses() AnyOfobject`

GetCellAddresses returns the CellAddresses field if non-nil, zero value otherwise.

### GetCellAddressesOk

`func (o *MicrosoftGraphWorkbookRangeView) GetCellAddressesOk() (*AnyOfobject, bool)`

GetCellAddressesOk returns a tuple with the CellAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCellAddresses

`func (o *MicrosoftGraphWorkbookRangeView) SetCellAddresses(v AnyOfobject)`

SetCellAddresses sets CellAddresses field to given value.

### HasCellAddresses

`func (o *MicrosoftGraphWorkbookRangeView) HasCellAddresses() bool`

HasCellAddresses returns a boolean if a field has been set.

### SetCellAddressesNil

`func (o *MicrosoftGraphWorkbookRangeView) SetCellAddressesNil(b bool)`

 SetCellAddressesNil sets the value for CellAddresses to be an explicit nil

### UnsetCellAddresses
`func (o *MicrosoftGraphWorkbookRangeView) UnsetCellAddresses()`

UnsetCellAddresses ensures that no value is present for CellAddresses, not even an explicit nil
### GetColumnCount

`func (o *MicrosoftGraphWorkbookRangeView) GetColumnCount() int32`

GetColumnCount returns the ColumnCount field if non-nil, zero value otherwise.

### GetColumnCountOk

`func (o *MicrosoftGraphWorkbookRangeView) GetColumnCountOk() (*int32, bool)`

GetColumnCountOk returns a tuple with the ColumnCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumnCount

`func (o *MicrosoftGraphWorkbookRangeView) SetColumnCount(v int32)`

SetColumnCount sets ColumnCount field to given value.

### HasColumnCount

`func (o *MicrosoftGraphWorkbookRangeView) HasColumnCount() bool`

HasColumnCount returns a boolean if a field has been set.

### GetFormulas

`func (o *MicrosoftGraphWorkbookRangeView) GetFormulas() AnyOfobject`

GetFormulas returns the Formulas field if non-nil, zero value otherwise.

### GetFormulasOk

`func (o *MicrosoftGraphWorkbookRangeView) GetFormulasOk() (*AnyOfobject, bool)`

GetFormulasOk returns a tuple with the Formulas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormulas

`func (o *MicrosoftGraphWorkbookRangeView) SetFormulas(v AnyOfobject)`

SetFormulas sets Formulas field to given value.

### HasFormulas

`func (o *MicrosoftGraphWorkbookRangeView) HasFormulas() bool`

HasFormulas returns a boolean if a field has been set.

### SetFormulasNil

`func (o *MicrosoftGraphWorkbookRangeView) SetFormulasNil(b bool)`

 SetFormulasNil sets the value for Formulas to be an explicit nil

### UnsetFormulas
`func (o *MicrosoftGraphWorkbookRangeView) UnsetFormulas()`

UnsetFormulas ensures that no value is present for Formulas, not even an explicit nil
### GetFormulasLocal

`func (o *MicrosoftGraphWorkbookRangeView) GetFormulasLocal() AnyOfobject`

GetFormulasLocal returns the FormulasLocal field if non-nil, zero value otherwise.

### GetFormulasLocalOk

`func (o *MicrosoftGraphWorkbookRangeView) GetFormulasLocalOk() (*AnyOfobject, bool)`

GetFormulasLocalOk returns a tuple with the FormulasLocal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormulasLocal

`func (o *MicrosoftGraphWorkbookRangeView) SetFormulasLocal(v AnyOfobject)`

SetFormulasLocal sets FormulasLocal field to given value.

### HasFormulasLocal

`func (o *MicrosoftGraphWorkbookRangeView) HasFormulasLocal() bool`

HasFormulasLocal returns a boolean if a field has been set.

### SetFormulasLocalNil

`func (o *MicrosoftGraphWorkbookRangeView) SetFormulasLocalNil(b bool)`

 SetFormulasLocalNil sets the value for FormulasLocal to be an explicit nil

### UnsetFormulasLocal
`func (o *MicrosoftGraphWorkbookRangeView) UnsetFormulasLocal()`

UnsetFormulasLocal ensures that no value is present for FormulasLocal, not even an explicit nil
### GetFormulasR1C1

`func (o *MicrosoftGraphWorkbookRangeView) GetFormulasR1C1() AnyOfobject`

GetFormulasR1C1 returns the FormulasR1C1 field if non-nil, zero value otherwise.

### GetFormulasR1C1Ok

`func (o *MicrosoftGraphWorkbookRangeView) GetFormulasR1C1Ok() (*AnyOfobject, bool)`

GetFormulasR1C1Ok returns a tuple with the FormulasR1C1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormulasR1C1

`func (o *MicrosoftGraphWorkbookRangeView) SetFormulasR1C1(v AnyOfobject)`

SetFormulasR1C1 sets FormulasR1C1 field to given value.

### HasFormulasR1C1

`func (o *MicrosoftGraphWorkbookRangeView) HasFormulasR1C1() bool`

HasFormulasR1C1 returns a boolean if a field has been set.

### SetFormulasR1C1Nil

`func (o *MicrosoftGraphWorkbookRangeView) SetFormulasR1C1Nil(b bool)`

 SetFormulasR1C1Nil sets the value for FormulasR1C1 to be an explicit nil

### UnsetFormulasR1C1
`func (o *MicrosoftGraphWorkbookRangeView) UnsetFormulasR1C1()`

UnsetFormulasR1C1 ensures that no value is present for FormulasR1C1, not even an explicit nil
### GetIndex

`func (o *MicrosoftGraphWorkbookRangeView) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *MicrosoftGraphWorkbookRangeView) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *MicrosoftGraphWorkbookRangeView) SetIndex(v int32)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *MicrosoftGraphWorkbookRangeView) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetNumberFormat

`func (o *MicrosoftGraphWorkbookRangeView) GetNumberFormat() AnyOfobject`

GetNumberFormat returns the NumberFormat field if non-nil, zero value otherwise.

### GetNumberFormatOk

`func (o *MicrosoftGraphWorkbookRangeView) GetNumberFormatOk() (*AnyOfobject, bool)`

GetNumberFormatOk returns a tuple with the NumberFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberFormat

`func (o *MicrosoftGraphWorkbookRangeView) SetNumberFormat(v AnyOfobject)`

SetNumberFormat sets NumberFormat field to given value.

### HasNumberFormat

`func (o *MicrosoftGraphWorkbookRangeView) HasNumberFormat() bool`

HasNumberFormat returns a boolean if a field has been set.

### SetNumberFormatNil

`func (o *MicrosoftGraphWorkbookRangeView) SetNumberFormatNil(b bool)`

 SetNumberFormatNil sets the value for NumberFormat to be an explicit nil

### UnsetNumberFormat
`func (o *MicrosoftGraphWorkbookRangeView) UnsetNumberFormat()`

UnsetNumberFormat ensures that no value is present for NumberFormat, not even an explicit nil
### GetRowCount

`func (o *MicrosoftGraphWorkbookRangeView) GetRowCount() int32`

GetRowCount returns the RowCount field if non-nil, zero value otherwise.

### GetRowCountOk

`func (o *MicrosoftGraphWorkbookRangeView) GetRowCountOk() (*int32, bool)`

GetRowCountOk returns a tuple with the RowCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRowCount

`func (o *MicrosoftGraphWorkbookRangeView) SetRowCount(v int32)`

SetRowCount sets RowCount field to given value.

### HasRowCount

`func (o *MicrosoftGraphWorkbookRangeView) HasRowCount() bool`

HasRowCount returns a boolean if a field has been set.

### GetText

`func (o *MicrosoftGraphWorkbookRangeView) GetText() AnyOfobject`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *MicrosoftGraphWorkbookRangeView) GetTextOk() (*AnyOfobject, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *MicrosoftGraphWorkbookRangeView) SetText(v AnyOfobject)`

SetText sets Text field to given value.

### HasText

`func (o *MicrosoftGraphWorkbookRangeView) HasText() bool`

HasText returns a boolean if a field has been set.

### SetTextNil

`func (o *MicrosoftGraphWorkbookRangeView) SetTextNil(b bool)`

 SetTextNil sets the value for Text to be an explicit nil

### UnsetText
`func (o *MicrosoftGraphWorkbookRangeView) UnsetText()`

UnsetText ensures that no value is present for Text, not even an explicit nil
### GetValues

`func (o *MicrosoftGraphWorkbookRangeView) GetValues() AnyOfobject`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *MicrosoftGraphWorkbookRangeView) GetValuesOk() (*AnyOfobject, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *MicrosoftGraphWorkbookRangeView) SetValues(v AnyOfobject)`

SetValues sets Values field to given value.

### HasValues

`func (o *MicrosoftGraphWorkbookRangeView) HasValues() bool`

HasValues returns a boolean if a field has been set.

### SetValuesNil

`func (o *MicrosoftGraphWorkbookRangeView) SetValuesNil(b bool)`

 SetValuesNil sets the value for Values to be an explicit nil

### UnsetValues
`func (o *MicrosoftGraphWorkbookRangeView) UnsetValues()`

UnsetValues ensures that no value is present for Values, not even an explicit nil
### GetValueTypes

`func (o *MicrosoftGraphWorkbookRangeView) GetValueTypes() AnyOfobject`

GetValueTypes returns the ValueTypes field if non-nil, zero value otherwise.

### GetValueTypesOk

`func (o *MicrosoftGraphWorkbookRangeView) GetValueTypesOk() (*AnyOfobject, bool)`

GetValueTypesOk returns a tuple with the ValueTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueTypes

`func (o *MicrosoftGraphWorkbookRangeView) SetValueTypes(v AnyOfobject)`

SetValueTypes sets ValueTypes field to given value.

### HasValueTypes

`func (o *MicrosoftGraphWorkbookRangeView) HasValueTypes() bool`

HasValueTypes returns a boolean if a field has been set.

### SetValueTypesNil

`func (o *MicrosoftGraphWorkbookRangeView) SetValueTypesNil(b bool)`

 SetValueTypesNil sets the value for ValueTypes to be an explicit nil

### UnsetValueTypes
`func (o *MicrosoftGraphWorkbookRangeView) UnsetValueTypes()`

UnsetValueTypes ensures that no value is present for ValueTypes, not even an explicit nil
### GetRows

`func (o *MicrosoftGraphWorkbookRangeView) GetRows() []MicrosoftGraphWorkbookRangeView`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *MicrosoftGraphWorkbookRangeView) GetRowsOk() (*[]MicrosoftGraphWorkbookRangeView, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *MicrosoftGraphWorkbookRangeView) SetRows(v []MicrosoftGraphWorkbookRangeView)`

SetRows sets Rows field to given value.

### HasRows

`func (o *MicrosoftGraphWorkbookRangeView) HasRows() bool`

HasRows returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


