# WorkbookTable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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

### NewWorkbookTable

`func NewWorkbookTable() *WorkbookTable`

NewWorkbookTable instantiates a new WorkbookTable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkbookTableWithDefaults

`func NewWorkbookTableWithDefaults() *WorkbookTable`

NewWorkbookTableWithDefaults instantiates a new WorkbookTable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHighlightFirstColumn

`func (o *WorkbookTable) GetHighlightFirstColumn() bool`

GetHighlightFirstColumn returns the HighlightFirstColumn field if non-nil, zero value otherwise.

### GetHighlightFirstColumnOk

`func (o *WorkbookTable) GetHighlightFirstColumnOk() (*bool, bool)`

GetHighlightFirstColumnOk returns a tuple with the HighlightFirstColumn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighlightFirstColumn

`func (o *WorkbookTable) SetHighlightFirstColumn(v bool)`

SetHighlightFirstColumn sets HighlightFirstColumn field to given value.

### HasHighlightFirstColumn

`func (o *WorkbookTable) HasHighlightFirstColumn() bool`

HasHighlightFirstColumn returns a boolean if a field has been set.

### GetHighlightLastColumn

`func (o *WorkbookTable) GetHighlightLastColumn() bool`

GetHighlightLastColumn returns the HighlightLastColumn field if non-nil, zero value otherwise.

### GetHighlightLastColumnOk

`func (o *WorkbookTable) GetHighlightLastColumnOk() (*bool, bool)`

GetHighlightLastColumnOk returns a tuple with the HighlightLastColumn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighlightLastColumn

`func (o *WorkbookTable) SetHighlightLastColumn(v bool)`

SetHighlightLastColumn sets HighlightLastColumn field to given value.

### HasHighlightLastColumn

`func (o *WorkbookTable) HasHighlightLastColumn() bool`

HasHighlightLastColumn returns a boolean if a field has been set.

### GetLegacyId

`func (o *WorkbookTable) GetLegacyId() string`

GetLegacyId returns the LegacyId field if non-nil, zero value otherwise.

### GetLegacyIdOk

`func (o *WorkbookTable) GetLegacyIdOk() (*string, bool)`

GetLegacyIdOk returns a tuple with the LegacyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegacyId

`func (o *WorkbookTable) SetLegacyId(v string)`

SetLegacyId sets LegacyId field to given value.

### HasLegacyId

`func (o *WorkbookTable) HasLegacyId() bool`

HasLegacyId returns a boolean if a field has been set.

### SetLegacyIdNil

`func (o *WorkbookTable) SetLegacyIdNil(b bool)`

 SetLegacyIdNil sets the value for LegacyId to be an explicit nil

### UnsetLegacyId
`func (o *WorkbookTable) UnsetLegacyId()`

UnsetLegacyId ensures that no value is present for LegacyId, not even an explicit nil
### GetName

`func (o *WorkbookTable) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkbookTable) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkbookTable) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WorkbookTable) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *WorkbookTable) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *WorkbookTable) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetShowBandedColumns

`func (o *WorkbookTable) GetShowBandedColumns() bool`

GetShowBandedColumns returns the ShowBandedColumns field if non-nil, zero value otherwise.

### GetShowBandedColumnsOk

`func (o *WorkbookTable) GetShowBandedColumnsOk() (*bool, bool)`

GetShowBandedColumnsOk returns a tuple with the ShowBandedColumns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowBandedColumns

`func (o *WorkbookTable) SetShowBandedColumns(v bool)`

SetShowBandedColumns sets ShowBandedColumns field to given value.

### HasShowBandedColumns

`func (o *WorkbookTable) HasShowBandedColumns() bool`

HasShowBandedColumns returns a boolean if a field has been set.

### GetShowBandedRows

`func (o *WorkbookTable) GetShowBandedRows() bool`

GetShowBandedRows returns the ShowBandedRows field if non-nil, zero value otherwise.

### GetShowBandedRowsOk

`func (o *WorkbookTable) GetShowBandedRowsOk() (*bool, bool)`

GetShowBandedRowsOk returns a tuple with the ShowBandedRows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowBandedRows

`func (o *WorkbookTable) SetShowBandedRows(v bool)`

SetShowBandedRows sets ShowBandedRows field to given value.

### HasShowBandedRows

`func (o *WorkbookTable) HasShowBandedRows() bool`

HasShowBandedRows returns a boolean if a field has been set.

### GetShowFilterButton

`func (o *WorkbookTable) GetShowFilterButton() bool`

GetShowFilterButton returns the ShowFilterButton field if non-nil, zero value otherwise.

### GetShowFilterButtonOk

`func (o *WorkbookTable) GetShowFilterButtonOk() (*bool, bool)`

GetShowFilterButtonOk returns a tuple with the ShowFilterButton field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowFilterButton

`func (o *WorkbookTable) SetShowFilterButton(v bool)`

SetShowFilterButton sets ShowFilterButton field to given value.

### HasShowFilterButton

`func (o *WorkbookTable) HasShowFilterButton() bool`

HasShowFilterButton returns a boolean if a field has been set.

### GetShowHeaders

`func (o *WorkbookTable) GetShowHeaders() bool`

GetShowHeaders returns the ShowHeaders field if non-nil, zero value otherwise.

### GetShowHeadersOk

`func (o *WorkbookTable) GetShowHeadersOk() (*bool, bool)`

GetShowHeadersOk returns a tuple with the ShowHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowHeaders

`func (o *WorkbookTable) SetShowHeaders(v bool)`

SetShowHeaders sets ShowHeaders field to given value.

### HasShowHeaders

`func (o *WorkbookTable) HasShowHeaders() bool`

HasShowHeaders returns a boolean if a field has been set.

### GetShowTotals

`func (o *WorkbookTable) GetShowTotals() bool`

GetShowTotals returns the ShowTotals field if non-nil, zero value otherwise.

### GetShowTotalsOk

`func (o *WorkbookTable) GetShowTotalsOk() (*bool, bool)`

GetShowTotalsOk returns a tuple with the ShowTotals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowTotals

`func (o *WorkbookTable) SetShowTotals(v bool)`

SetShowTotals sets ShowTotals field to given value.

### HasShowTotals

`func (o *WorkbookTable) HasShowTotals() bool`

HasShowTotals returns a boolean if a field has been set.

### GetStyle

`func (o *WorkbookTable) GetStyle() string`

GetStyle returns the Style field if non-nil, zero value otherwise.

### GetStyleOk

`func (o *WorkbookTable) GetStyleOk() (*string, bool)`

GetStyleOk returns a tuple with the Style field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStyle

`func (o *WorkbookTable) SetStyle(v string)`

SetStyle sets Style field to given value.

### HasStyle

`func (o *WorkbookTable) HasStyle() bool`

HasStyle returns a boolean if a field has been set.

### SetStyleNil

`func (o *WorkbookTable) SetStyleNil(b bool)`

 SetStyleNil sets the value for Style to be an explicit nil

### UnsetStyle
`func (o *WorkbookTable) UnsetStyle()`

UnsetStyle ensures that no value is present for Style, not even an explicit nil
### GetColumns

`func (o *WorkbookTable) GetColumns() []MicrosoftGraphWorkbookTableColumn`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *WorkbookTable) GetColumnsOk() (*[]MicrosoftGraphWorkbookTableColumn, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumns

`func (o *WorkbookTable) SetColumns(v []MicrosoftGraphWorkbookTableColumn)`

SetColumns sets Columns field to given value.

### HasColumns

`func (o *WorkbookTable) HasColumns() bool`

HasColumns returns a boolean if a field has been set.

### GetRows

`func (o *WorkbookTable) GetRows() []MicrosoftGraphWorkbookTableRow`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *WorkbookTable) GetRowsOk() (*[]MicrosoftGraphWorkbookTableRow, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *WorkbookTable) SetRows(v []MicrosoftGraphWorkbookTableRow)`

SetRows sets Rows field to given value.

### HasRows

`func (o *WorkbookTable) HasRows() bool`

HasRows returns a boolean if a field has been set.

### GetSort

`func (o *WorkbookTable) GetSort() AnyOfmicrosoftGraphWorkbookTableSort`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *WorkbookTable) GetSortOk() (*AnyOfmicrosoftGraphWorkbookTableSort, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *WorkbookTable) SetSort(v AnyOfmicrosoftGraphWorkbookTableSort)`

SetSort sets Sort field to given value.

### HasSort

`func (o *WorkbookTable) HasSort() bool`

HasSort returns a boolean if a field has been set.

### SetSortNil

`func (o *WorkbookTable) SetSortNil(b bool)`

 SetSortNil sets the value for Sort to be an explicit nil

### UnsetSort
`func (o *WorkbookTable) UnsetSort()`

UnsetSort ensures that no value is present for Sort, not even an explicit nil
### GetWorksheet

`func (o *WorkbookTable) GetWorksheet() AnyOfmicrosoftGraphWorkbookWorksheet`

GetWorksheet returns the Worksheet field if non-nil, zero value otherwise.

### GetWorksheetOk

`func (o *WorkbookTable) GetWorksheetOk() (*AnyOfmicrosoftGraphWorkbookWorksheet, bool)`

GetWorksheetOk returns a tuple with the Worksheet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorksheet

`func (o *WorkbookTable) SetWorksheet(v AnyOfmicrosoftGraphWorkbookWorksheet)`

SetWorksheet sets Worksheet field to given value.

### HasWorksheet

`func (o *WorkbookTable) HasWorksheet() bool`

HasWorksheet returns a boolean if a field has been set.

### SetWorksheetNil

`func (o *WorkbookTable) SetWorksheetNil(b bool)`

 SetWorksheetNil sets the value for Worksheet to be an explicit nil

### UnsetWorksheet
`func (o *WorkbookTable) UnsetWorksheet()`

UnsetWorksheet ensures that no value is present for Worksheet, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


