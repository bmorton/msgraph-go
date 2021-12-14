# MicrosoftGraphWorkbookTable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**HighlightFirstColumn** | Pointer to **bool** | Indicates whether the first column contains special formatting. | [optional] 
**HighlightLastColumn** | Pointer to **bool** | Indicates whether the last column contains special formatting. | [optional] 
**LegacyId** | Pointer to **NullableString** | Legacy Id used in older Excle clients. The value of the identifier remains the same even when the table is renamed. This property should be interpreted as an opaque string value and should not be parsed to any other type. Read-only. | [optional] 
**Name** | Pointer to **NullableString** | Name of the table. | [optional] 
**ShowBandedColumns** | Pointer to **bool** | Indicates whether the columns show banded formatting in which odd columns are highlighted differently from even ones to make reading the table easier. | [optional] 
**ShowBandedRows** | Pointer to **bool** | Indicates whether the rows show banded formatting in which odd rows are highlighted differently from even ones to make reading the table easier. | [optional] 
**ShowFilterButton** | Pointer to **bool** | Indicates whether the filter buttons are visible at the top of each column header. Setting this is only allowed if the table contains a header row. | [optional] 
**ShowHeaders** | Pointer to **bool** | Indicates whether the header row is visible or not. This value can be set to show or remove the header row. | [optional] 
**ShowTotals** | Pointer to **bool** | Indicates whether the total row is visible or not. This value can be set to show or remove the total row. | [optional] 
**Style** | Pointer to **NullableString** | Constant value that represents the Table style. The possible values are: TableStyleLight1 thru TableStyleLight21, TableStyleMedium1 thru TableStyleMedium28, TableStyleStyleDark1 thru TableStyleStyleDark11. A custom user-defined style present in the workbook can also be specified. | [optional] 
**Columns** | Pointer to [**[]MicrosoftGraphWorkbookTableColumn**](MicrosoftGraphWorkbookTableColumn.md) | Represents a collection of all the columns in the table. Read-only. | [optional] 
**Rows** | Pointer to [**[]MicrosoftGraphWorkbookTableRow**](MicrosoftGraphWorkbookTableRow.md) | Represents a collection of all the rows in the table. Read-only. | [optional] 
**Sort** | Pointer to [**AnyOfmicrosoftGraphWorkbookTableSort**](anyOf&lt;microsoft.graph.workbookTableSort&gt;.md) | Represents the sorting for the table. Read-only. | [optional] 
**Worksheet** | Pointer to [**AnyOfmicrosoftGraphWorkbookWorksheet**](anyOf&lt;microsoft.graph.workbookWorksheet&gt;.md) | The worksheet containing the current table. Read-only. | [optional] 

## Methods

### NewMicrosoftGraphWorkbookTable

`func NewMicrosoftGraphWorkbookTable() *MicrosoftGraphWorkbookTable`

NewMicrosoftGraphWorkbookTable instantiates a new MicrosoftGraphWorkbookTable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphWorkbookTableWithDefaults

`func NewMicrosoftGraphWorkbookTableWithDefaults() *MicrosoftGraphWorkbookTable`

NewMicrosoftGraphWorkbookTableWithDefaults instantiates a new MicrosoftGraphWorkbookTable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphWorkbookTable) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphWorkbookTable) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphWorkbookTable) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphWorkbookTable) HasId() bool`

HasId returns a boolean if a field has been set.

### GetHighlightFirstColumn

`func (o *MicrosoftGraphWorkbookTable) GetHighlightFirstColumn() bool`

GetHighlightFirstColumn returns the HighlightFirstColumn field if non-nil, zero value otherwise.

### GetHighlightFirstColumnOk

`func (o *MicrosoftGraphWorkbookTable) GetHighlightFirstColumnOk() (*bool, bool)`

GetHighlightFirstColumnOk returns a tuple with the HighlightFirstColumn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighlightFirstColumn

`func (o *MicrosoftGraphWorkbookTable) SetHighlightFirstColumn(v bool)`

SetHighlightFirstColumn sets HighlightFirstColumn field to given value.

### HasHighlightFirstColumn

`func (o *MicrosoftGraphWorkbookTable) HasHighlightFirstColumn() bool`

HasHighlightFirstColumn returns a boolean if a field has been set.

### GetHighlightLastColumn

`func (o *MicrosoftGraphWorkbookTable) GetHighlightLastColumn() bool`

GetHighlightLastColumn returns the HighlightLastColumn field if non-nil, zero value otherwise.

### GetHighlightLastColumnOk

`func (o *MicrosoftGraphWorkbookTable) GetHighlightLastColumnOk() (*bool, bool)`

GetHighlightLastColumnOk returns a tuple with the HighlightLastColumn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighlightLastColumn

`func (o *MicrosoftGraphWorkbookTable) SetHighlightLastColumn(v bool)`

SetHighlightLastColumn sets HighlightLastColumn field to given value.

### HasHighlightLastColumn

`func (o *MicrosoftGraphWorkbookTable) HasHighlightLastColumn() bool`

HasHighlightLastColumn returns a boolean if a field has been set.

### GetLegacyId

`func (o *MicrosoftGraphWorkbookTable) GetLegacyId() string`

GetLegacyId returns the LegacyId field if non-nil, zero value otherwise.

### GetLegacyIdOk

`func (o *MicrosoftGraphWorkbookTable) GetLegacyIdOk() (*string, bool)`

GetLegacyIdOk returns a tuple with the LegacyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegacyId

`func (o *MicrosoftGraphWorkbookTable) SetLegacyId(v string)`

SetLegacyId sets LegacyId field to given value.

### HasLegacyId

`func (o *MicrosoftGraphWorkbookTable) HasLegacyId() bool`

HasLegacyId returns a boolean if a field has been set.

### SetLegacyIdNil

`func (o *MicrosoftGraphWorkbookTable) SetLegacyIdNil(b bool)`

 SetLegacyIdNil sets the value for LegacyId to be an explicit nil

### UnsetLegacyId
`func (o *MicrosoftGraphWorkbookTable) UnsetLegacyId()`

UnsetLegacyId ensures that no value is present for LegacyId, not even an explicit nil
### GetName

`func (o *MicrosoftGraphWorkbookTable) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MicrosoftGraphWorkbookTable) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MicrosoftGraphWorkbookTable) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MicrosoftGraphWorkbookTable) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *MicrosoftGraphWorkbookTable) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *MicrosoftGraphWorkbookTable) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetShowBandedColumns

`func (o *MicrosoftGraphWorkbookTable) GetShowBandedColumns() bool`

GetShowBandedColumns returns the ShowBandedColumns field if non-nil, zero value otherwise.

### GetShowBandedColumnsOk

`func (o *MicrosoftGraphWorkbookTable) GetShowBandedColumnsOk() (*bool, bool)`

GetShowBandedColumnsOk returns a tuple with the ShowBandedColumns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowBandedColumns

`func (o *MicrosoftGraphWorkbookTable) SetShowBandedColumns(v bool)`

SetShowBandedColumns sets ShowBandedColumns field to given value.

### HasShowBandedColumns

`func (o *MicrosoftGraphWorkbookTable) HasShowBandedColumns() bool`

HasShowBandedColumns returns a boolean if a field has been set.

### GetShowBandedRows

`func (o *MicrosoftGraphWorkbookTable) GetShowBandedRows() bool`

GetShowBandedRows returns the ShowBandedRows field if non-nil, zero value otherwise.

### GetShowBandedRowsOk

`func (o *MicrosoftGraphWorkbookTable) GetShowBandedRowsOk() (*bool, bool)`

GetShowBandedRowsOk returns a tuple with the ShowBandedRows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowBandedRows

`func (o *MicrosoftGraphWorkbookTable) SetShowBandedRows(v bool)`

SetShowBandedRows sets ShowBandedRows field to given value.

### HasShowBandedRows

`func (o *MicrosoftGraphWorkbookTable) HasShowBandedRows() bool`

HasShowBandedRows returns a boolean if a field has been set.

### GetShowFilterButton

`func (o *MicrosoftGraphWorkbookTable) GetShowFilterButton() bool`

GetShowFilterButton returns the ShowFilterButton field if non-nil, zero value otherwise.

### GetShowFilterButtonOk

`func (o *MicrosoftGraphWorkbookTable) GetShowFilterButtonOk() (*bool, bool)`

GetShowFilterButtonOk returns a tuple with the ShowFilterButton field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowFilterButton

`func (o *MicrosoftGraphWorkbookTable) SetShowFilterButton(v bool)`

SetShowFilterButton sets ShowFilterButton field to given value.

### HasShowFilterButton

`func (o *MicrosoftGraphWorkbookTable) HasShowFilterButton() bool`

HasShowFilterButton returns a boolean if a field has been set.

### GetShowHeaders

`func (o *MicrosoftGraphWorkbookTable) GetShowHeaders() bool`

GetShowHeaders returns the ShowHeaders field if non-nil, zero value otherwise.

### GetShowHeadersOk

`func (o *MicrosoftGraphWorkbookTable) GetShowHeadersOk() (*bool, bool)`

GetShowHeadersOk returns a tuple with the ShowHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowHeaders

`func (o *MicrosoftGraphWorkbookTable) SetShowHeaders(v bool)`

SetShowHeaders sets ShowHeaders field to given value.

### HasShowHeaders

`func (o *MicrosoftGraphWorkbookTable) HasShowHeaders() bool`

HasShowHeaders returns a boolean if a field has been set.

### GetShowTotals

`func (o *MicrosoftGraphWorkbookTable) GetShowTotals() bool`

GetShowTotals returns the ShowTotals field if non-nil, zero value otherwise.

### GetShowTotalsOk

`func (o *MicrosoftGraphWorkbookTable) GetShowTotalsOk() (*bool, bool)`

GetShowTotalsOk returns a tuple with the ShowTotals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowTotals

`func (o *MicrosoftGraphWorkbookTable) SetShowTotals(v bool)`

SetShowTotals sets ShowTotals field to given value.

### HasShowTotals

`func (o *MicrosoftGraphWorkbookTable) HasShowTotals() bool`

HasShowTotals returns a boolean if a field has been set.

### GetStyle

`func (o *MicrosoftGraphWorkbookTable) GetStyle() string`

GetStyle returns the Style field if non-nil, zero value otherwise.

### GetStyleOk

`func (o *MicrosoftGraphWorkbookTable) GetStyleOk() (*string, bool)`

GetStyleOk returns a tuple with the Style field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStyle

`func (o *MicrosoftGraphWorkbookTable) SetStyle(v string)`

SetStyle sets Style field to given value.

### HasStyle

`func (o *MicrosoftGraphWorkbookTable) HasStyle() bool`

HasStyle returns a boolean if a field has been set.

### SetStyleNil

`func (o *MicrosoftGraphWorkbookTable) SetStyleNil(b bool)`

 SetStyleNil sets the value for Style to be an explicit nil

### UnsetStyle
`func (o *MicrosoftGraphWorkbookTable) UnsetStyle()`

UnsetStyle ensures that no value is present for Style, not even an explicit nil
### GetColumns

`func (o *MicrosoftGraphWorkbookTable) GetColumns() []MicrosoftGraphWorkbookTableColumn`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *MicrosoftGraphWorkbookTable) GetColumnsOk() (*[]MicrosoftGraphWorkbookTableColumn, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumns

`func (o *MicrosoftGraphWorkbookTable) SetColumns(v []MicrosoftGraphWorkbookTableColumn)`

SetColumns sets Columns field to given value.

### HasColumns

`func (o *MicrosoftGraphWorkbookTable) HasColumns() bool`

HasColumns returns a boolean if a field has been set.

### GetRows

`func (o *MicrosoftGraphWorkbookTable) GetRows() []MicrosoftGraphWorkbookTableRow`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *MicrosoftGraphWorkbookTable) GetRowsOk() (*[]MicrosoftGraphWorkbookTableRow, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *MicrosoftGraphWorkbookTable) SetRows(v []MicrosoftGraphWorkbookTableRow)`

SetRows sets Rows field to given value.

### HasRows

`func (o *MicrosoftGraphWorkbookTable) HasRows() bool`

HasRows returns a boolean if a field has been set.

### GetSort

`func (o *MicrosoftGraphWorkbookTable) GetSort() AnyOfmicrosoftGraphWorkbookTableSort`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *MicrosoftGraphWorkbookTable) GetSortOk() (*AnyOfmicrosoftGraphWorkbookTableSort, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *MicrosoftGraphWorkbookTable) SetSort(v AnyOfmicrosoftGraphWorkbookTableSort)`

SetSort sets Sort field to given value.

### HasSort

`func (o *MicrosoftGraphWorkbookTable) HasSort() bool`

HasSort returns a boolean if a field has been set.

### SetSortNil

`func (o *MicrosoftGraphWorkbookTable) SetSortNil(b bool)`

 SetSortNil sets the value for Sort to be an explicit nil

### UnsetSort
`func (o *MicrosoftGraphWorkbookTable) UnsetSort()`

UnsetSort ensures that no value is present for Sort, not even an explicit nil
### GetWorksheet

`func (o *MicrosoftGraphWorkbookTable) GetWorksheet() AnyOfmicrosoftGraphWorkbookWorksheet`

GetWorksheet returns the Worksheet field if non-nil, zero value otherwise.

### GetWorksheetOk

`func (o *MicrosoftGraphWorkbookTable) GetWorksheetOk() (*AnyOfmicrosoftGraphWorkbookWorksheet, bool)`

GetWorksheetOk returns a tuple with the Worksheet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorksheet

`func (o *MicrosoftGraphWorkbookTable) SetWorksheet(v AnyOfmicrosoftGraphWorkbookWorksheet)`

SetWorksheet sets Worksheet field to given value.

### HasWorksheet

`func (o *MicrosoftGraphWorkbookTable) HasWorksheet() bool`

HasWorksheet returns a boolean if a field has been set.

### SetWorksheetNil

`func (o *MicrosoftGraphWorkbookTable) SetWorksheetNil(b bool)`

 SetWorksheetNil sets the value for Worksheet to be an explicit nil

### UnsetWorksheet
`func (o *MicrosoftGraphWorkbookTable) UnsetWorksheet()`

UnsetWorksheet ensures that no value is present for Worksheet, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


