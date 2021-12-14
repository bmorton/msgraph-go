/*
Partial Graph API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// WorkbookTable struct for WorkbookTable
type WorkbookTable struct {
	// Indicates whether the first column contains special formatting.
	HighlightFirstColumn *bool `json:"highlightFirstColumn,omitempty"`
	// Indicates whether the last column contains special formatting.
	HighlightLastColumn *bool `json:"highlightLastColumn,omitempty"`
	// Legacy Id used in older Excle clients. The value of the identifier remains the same even when the table is renamed. This property should be interpreted as an opaque string value and should not be parsed to any other type. Read-only.
	LegacyId NullableString `json:"legacyId,omitempty"`
	// Name of the table.
	Name NullableString `json:"name,omitempty"`
	// Indicates whether the columns show banded formatting in which odd columns are highlighted differently from even ones to make reading the table easier.
	ShowBandedColumns *bool `json:"showBandedColumns,omitempty"`
	// Indicates whether the rows show banded formatting in which odd rows are highlighted differently from even ones to make reading the table easier.
	ShowBandedRows *bool `json:"showBandedRows,omitempty"`
	// Indicates whether the filter buttons are visible at the top of each column header. Setting this is only allowed if the table contains a header row.
	ShowFilterButton *bool `json:"showFilterButton,omitempty"`
	// Indicates whether the header row is visible or not. This value can be set to show or remove the header row.
	ShowHeaders *bool `json:"showHeaders,omitempty"`
	// Indicates whether the total row is visible or not. This value can be set to show or remove the total row.
	ShowTotals *bool `json:"showTotals,omitempty"`
	// Constant value that represents the Table style. The possible values are: TableStyleLight1 thru TableStyleLight21, TableStyleMedium1 thru TableStyleMedium28, TableStyleStyleDark1 thru TableStyleStyleDark11. A custom user-defined style present in the workbook can also be specified.
	Style NullableString `json:"style,omitempty"`
	// Represents a collection of all the columns in the table. Read-only.
	Columns *[]MicrosoftGraphWorkbookTableColumn `json:"columns,omitempty"`
	// Represents a collection of all the rows in the table. Read-only.
	Rows *[]MicrosoftGraphWorkbookTableRow `json:"rows,omitempty"`
	// Represents the sorting for the table. Read-only.
	Sort AnyOfmicrosoftGraphWorkbookTableSort `json:"sort,omitempty"`
	// The worksheet containing the current table. Read-only.
	Worksheet AnyOfmicrosoftGraphWorkbookWorksheet `json:"worksheet,omitempty"`
}

// NewWorkbookTable instantiates a new WorkbookTable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkbookTable() *WorkbookTable {
	this := WorkbookTable{}
	return &this
}

// NewWorkbookTableWithDefaults instantiates a new WorkbookTable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkbookTableWithDefaults() *WorkbookTable {
	this := WorkbookTable{}
	return &this
}

// GetHighlightFirstColumn returns the HighlightFirstColumn field value if set, zero value otherwise.
func (o *WorkbookTable) GetHighlightFirstColumn() bool {
	if o == nil || o.HighlightFirstColumn == nil {
		var ret bool
		return ret
	}
	return *o.HighlightFirstColumn
}

// GetHighlightFirstColumnOk returns a tuple with the HighlightFirstColumn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkbookTable) GetHighlightFirstColumnOk() (*bool, bool) {
	if o == nil || o.HighlightFirstColumn == nil {
		return nil, false
	}
	return o.HighlightFirstColumn, true
}

// HasHighlightFirstColumn returns a boolean if a field has been set.
func (o *WorkbookTable) HasHighlightFirstColumn() bool {
	if o != nil && o.HighlightFirstColumn != nil {
		return true
	}

	return false
}

// SetHighlightFirstColumn gets a reference to the given bool and assigns it to the HighlightFirstColumn field.
func (o *WorkbookTable) SetHighlightFirstColumn(v bool) {
	o.HighlightFirstColumn = &v
}

// GetHighlightLastColumn returns the HighlightLastColumn field value if set, zero value otherwise.
func (o *WorkbookTable) GetHighlightLastColumn() bool {
	if o == nil || o.HighlightLastColumn == nil {
		var ret bool
		return ret
	}
	return *o.HighlightLastColumn
}

// GetHighlightLastColumnOk returns a tuple with the HighlightLastColumn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkbookTable) GetHighlightLastColumnOk() (*bool, bool) {
	if o == nil || o.HighlightLastColumn == nil {
		return nil, false
	}
	return o.HighlightLastColumn, true
}

// HasHighlightLastColumn returns a boolean if a field has been set.
func (o *WorkbookTable) HasHighlightLastColumn() bool {
	if o != nil && o.HighlightLastColumn != nil {
		return true
	}

	return false
}

// SetHighlightLastColumn gets a reference to the given bool and assigns it to the HighlightLastColumn field.
func (o *WorkbookTable) SetHighlightLastColumn(v bool) {
	o.HighlightLastColumn = &v
}

// GetLegacyId returns the LegacyId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkbookTable) GetLegacyId() string {
	if o == nil || o.LegacyId.Get() == nil {
		var ret string
		return ret
	}
	return *o.LegacyId.Get()
}

// GetLegacyIdOk returns a tuple with the LegacyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkbookTable) GetLegacyIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LegacyId.Get(), o.LegacyId.IsSet()
}

// HasLegacyId returns a boolean if a field has been set.
func (o *WorkbookTable) HasLegacyId() bool {
	if o != nil && o.LegacyId.IsSet() {
		return true
	}

	return false
}

// SetLegacyId gets a reference to the given NullableString and assigns it to the LegacyId field.
func (o *WorkbookTable) SetLegacyId(v string) {
	o.LegacyId.Set(&v)
}
// SetLegacyIdNil sets the value for LegacyId to be an explicit nil
func (o *WorkbookTable) SetLegacyIdNil() {
	o.LegacyId.Set(nil)
}

// UnsetLegacyId ensures that no value is present for LegacyId, not even an explicit nil
func (o *WorkbookTable) UnsetLegacyId() {
	o.LegacyId.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkbookTable) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkbookTable) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *WorkbookTable) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *WorkbookTable) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *WorkbookTable) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *WorkbookTable) UnsetName() {
	o.Name.Unset()
}

// GetShowBandedColumns returns the ShowBandedColumns field value if set, zero value otherwise.
func (o *WorkbookTable) GetShowBandedColumns() bool {
	if o == nil || o.ShowBandedColumns == nil {
		var ret bool
		return ret
	}
	return *o.ShowBandedColumns
}

// GetShowBandedColumnsOk returns a tuple with the ShowBandedColumns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkbookTable) GetShowBandedColumnsOk() (*bool, bool) {
	if o == nil || o.ShowBandedColumns == nil {
		return nil, false
	}
	return o.ShowBandedColumns, true
}

// HasShowBandedColumns returns a boolean if a field has been set.
func (o *WorkbookTable) HasShowBandedColumns() bool {
	if o != nil && o.ShowBandedColumns != nil {
		return true
	}

	return false
}

// SetShowBandedColumns gets a reference to the given bool and assigns it to the ShowBandedColumns field.
func (o *WorkbookTable) SetShowBandedColumns(v bool) {
	o.ShowBandedColumns = &v
}

// GetShowBandedRows returns the ShowBandedRows field value if set, zero value otherwise.
func (o *WorkbookTable) GetShowBandedRows() bool {
	if o == nil || o.ShowBandedRows == nil {
		var ret bool
		return ret
	}
	return *o.ShowBandedRows
}

// GetShowBandedRowsOk returns a tuple with the ShowBandedRows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkbookTable) GetShowBandedRowsOk() (*bool, bool) {
	if o == nil || o.ShowBandedRows == nil {
		return nil, false
	}
	return o.ShowBandedRows, true
}

// HasShowBandedRows returns a boolean if a field has been set.
func (o *WorkbookTable) HasShowBandedRows() bool {
	if o != nil && o.ShowBandedRows != nil {
		return true
	}

	return false
}

// SetShowBandedRows gets a reference to the given bool and assigns it to the ShowBandedRows field.
func (o *WorkbookTable) SetShowBandedRows(v bool) {
	o.ShowBandedRows = &v
}

// GetShowFilterButton returns the ShowFilterButton field value if set, zero value otherwise.
func (o *WorkbookTable) GetShowFilterButton() bool {
	if o == nil || o.ShowFilterButton == nil {
		var ret bool
		return ret
	}
	return *o.ShowFilterButton
}

// GetShowFilterButtonOk returns a tuple with the ShowFilterButton field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkbookTable) GetShowFilterButtonOk() (*bool, bool) {
	if o == nil || o.ShowFilterButton == nil {
		return nil, false
	}
	return o.ShowFilterButton, true
}

// HasShowFilterButton returns a boolean if a field has been set.
func (o *WorkbookTable) HasShowFilterButton() bool {
	if o != nil && o.ShowFilterButton != nil {
		return true
	}

	return false
}

// SetShowFilterButton gets a reference to the given bool and assigns it to the ShowFilterButton field.
func (o *WorkbookTable) SetShowFilterButton(v bool) {
	o.ShowFilterButton = &v
}

// GetShowHeaders returns the ShowHeaders field value if set, zero value otherwise.
func (o *WorkbookTable) GetShowHeaders() bool {
	if o == nil || o.ShowHeaders == nil {
		var ret bool
		return ret
	}
	return *o.ShowHeaders
}

// GetShowHeadersOk returns a tuple with the ShowHeaders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkbookTable) GetShowHeadersOk() (*bool, bool) {
	if o == nil || o.ShowHeaders == nil {
		return nil, false
	}
	return o.ShowHeaders, true
}

// HasShowHeaders returns a boolean if a field has been set.
func (o *WorkbookTable) HasShowHeaders() bool {
	if o != nil && o.ShowHeaders != nil {
		return true
	}

	return false
}

// SetShowHeaders gets a reference to the given bool and assigns it to the ShowHeaders field.
func (o *WorkbookTable) SetShowHeaders(v bool) {
	o.ShowHeaders = &v
}

// GetShowTotals returns the ShowTotals field value if set, zero value otherwise.
func (o *WorkbookTable) GetShowTotals() bool {
	if o == nil || o.ShowTotals == nil {
		var ret bool
		return ret
	}
	return *o.ShowTotals
}

// GetShowTotalsOk returns a tuple with the ShowTotals field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkbookTable) GetShowTotalsOk() (*bool, bool) {
	if o == nil || o.ShowTotals == nil {
		return nil, false
	}
	return o.ShowTotals, true
}

// HasShowTotals returns a boolean if a field has been set.
func (o *WorkbookTable) HasShowTotals() bool {
	if o != nil && o.ShowTotals != nil {
		return true
	}

	return false
}

// SetShowTotals gets a reference to the given bool and assigns it to the ShowTotals field.
func (o *WorkbookTable) SetShowTotals(v bool) {
	o.ShowTotals = &v
}

// GetStyle returns the Style field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkbookTable) GetStyle() string {
	if o == nil || o.Style.Get() == nil {
		var ret string
		return ret
	}
	return *o.Style.Get()
}

// GetStyleOk returns a tuple with the Style field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkbookTable) GetStyleOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Style.Get(), o.Style.IsSet()
}

// HasStyle returns a boolean if a field has been set.
func (o *WorkbookTable) HasStyle() bool {
	if o != nil && o.Style.IsSet() {
		return true
	}

	return false
}

// SetStyle gets a reference to the given NullableString and assigns it to the Style field.
func (o *WorkbookTable) SetStyle(v string) {
	o.Style.Set(&v)
}
// SetStyleNil sets the value for Style to be an explicit nil
func (o *WorkbookTable) SetStyleNil() {
	o.Style.Set(nil)
}

// UnsetStyle ensures that no value is present for Style, not even an explicit nil
func (o *WorkbookTable) UnsetStyle() {
	o.Style.Unset()
}

// GetColumns returns the Columns field value if set, zero value otherwise.
func (o *WorkbookTable) GetColumns() []MicrosoftGraphWorkbookTableColumn {
	if o == nil || o.Columns == nil {
		var ret []MicrosoftGraphWorkbookTableColumn
		return ret
	}
	return *o.Columns
}

// GetColumnsOk returns a tuple with the Columns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkbookTable) GetColumnsOk() (*[]MicrosoftGraphWorkbookTableColumn, bool) {
	if o == nil || o.Columns == nil {
		return nil, false
	}
	return o.Columns, true
}

// HasColumns returns a boolean if a field has been set.
func (o *WorkbookTable) HasColumns() bool {
	if o != nil && o.Columns != nil {
		return true
	}

	return false
}

// SetColumns gets a reference to the given []MicrosoftGraphWorkbookTableColumn and assigns it to the Columns field.
func (o *WorkbookTable) SetColumns(v []MicrosoftGraphWorkbookTableColumn) {
	o.Columns = &v
}

// GetRows returns the Rows field value if set, zero value otherwise.
func (o *WorkbookTable) GetRows() []MicrosoftGraphWorkbookTableRow {
	if o == nil || o.Rows == nil {
		var ret []MicrosoftGraphWorkbookTableRow
		return ret
	}
	return *o.Rows
}

// GetRowsOk returns a tuple with the Rows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkbookTable) GetRowsOk() (*[]MicrosoftGraphWorkbookTableRow, bool) {
	if o == nil || o.Rows == nil {
		return nil, false
	}
	return o.Rows, true
}

// HasRows returns a boolean if a field has been set.
func (o *WorkbookTable) HasRows() bool {
	if o != nil && o.Rows != nil {
		return true
	}

	return false
}

// SetRows gets a reference to the given []MicrosoftGraphWorkbookTableRow and assigns it to the Rows field.
func (o *WorkbookTable) SetRows(v []MicrosoftGraphWorkbookTableRow) {
	o.Rows = &v
}

// GetSort returns the Sort field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkbookTable) GetSort() AnyOfmicrosoftGraphWorkbookTableSort {
	if o == nil  {
		var ret AnyOfmicrosoftGraphWorkbookTableSort
		return ret
	}
	return o.Sort
}

// GetSortOk returns a tuple with the Sort field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkbookTable) GetSortOk() (*AnyOfmicrosoftGraphWorkbookTableSort, bool) {
	if o == nil || o.Sort == nil {
		return nil, false
	}
	return &o.Sort, true
}

// HasSort returns a boolean if a field has been set.
func (o *WorkbookTable) HasSort() bool {
	if o != nil && o.Sort != nil {
		return true
	}

	return false
}

// SetSort gets a reference to the given AnyOfmicrosoftGraphWorkbookTableSort and assigns it to the Sort field.
func (o *WorkbookTable) SetSort(v AnyOfmicrosoftGraphWorkbookTableSort) {
	o.Sort = v
}

// GetWorksheet returns the Worksheet field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkbookTable) GetWorksheet() AnyOfmicrosoftGraphWorkbookWorksheet {
	if o == nil  {
		var ret AnyOfmicrosoftGraphWorkbookWorksheet
		return ret
	}
	return o.Worksheet
}

// GetWorksheetOk returns a tuple with the Worksheet field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkbookTable) GetWorksheetOk() (*AnyOfmicrosoftGraphWorkbookWorksheet, bool) {
	if o == nil || o.Worksheet == nil {
		return nil, false
	}
	return &o.Worksheet, true
}

// HasWorksheet returns a boolean if a field has been set.
func (o *WorkbookTable) HasWorksheet() bool {
	if o != nil && o.Worksheet != nil {
		return true
	}

	return false
}

// SetWorksheet gets a reference to the given AnyOfmicrosoftGraphWorkbookWorksheet and assigns it to the Worksheet field.
func (o *WorkbookTable) SetWorksheet(v AnyOfmicrosoftGraphWorkbookWorksheet) {
	o.Worksheet = v
}

func (o WorkbookTable) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.HighlightFirstColumn != nil {
		toSerialize["highlightFirstColumn"] = o.HighlightFirstColumn
	}
	if o.HighlightLastColumn != nil {
		toSerialize["highlightLastColumn"] = o.HighlightLastColumn
	}
	if o.LegacyId.IsSet() {
		toSerialize["legacyId"] = o.LegacyId.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.ShowBandedColumns != nil {
		toSerialize["showBandedColumns"] = o.ShowBandedColumns
	}
	if o.ShowBandedRows != nil {
		toSerialize["showBandedRows"] = o.ShowBandedRows
	}
	if o.ShowFilterButton != nil {
		toSerialize["showFilterButton"] = o.ShowFilterButton
	}
	if o.ShowHeaders != nil {
		toSerialize["showHeaders"] = o.ShowHeaders
	}
	if o.ShowTotals != nil {
		toSerialize["showTotals"] = o.ShowTotals
	}
	if o.Style.IsSet() {
		toSerialize["style"] = o.Style.Get()
	}
	if o.Columns != nil {
		toSerialize["columns"] = o.Columns
	}
	if o.Rows != nil {
		toSerialize["rows"] = o.Rows
	}
	if o.Sort != nil {
		toSerialize["sort"] = o.Sort
	}
	if o.Worksheet != nil {
		toSerialize["worksheet"] = o.Worksheet
	}
	return json.Marshal(toSerialize)
}

type NullableWorkbookTable struct {
	value *WorkbookTable
	isSet bool
}

func (v NullableWorkbookTable) Get() *WorkbookTable {
	return v.value
}

func (v *NullableWorkbookTable) Set(val *WorkbookTable) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkbookTable) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkbookTable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkbookTable(val *WorkbookTable) *NullableWorkbookTable {
	return &NullableWorkbookTable{value: val, isSet: true}
}

func (v NullableWorkbookTable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkbookTable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


