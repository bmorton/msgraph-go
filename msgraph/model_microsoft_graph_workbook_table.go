/*
Partial Graph API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package msgraph

import (
	"encoding/json"
)

// MicrosoftGraphWorkbookTable struct for MicrosoftGraphWorkbookTable
type MicrosoftGraphWorkbookTable struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
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

// NewMicrosoftGraphWorkbookTable instantiates a new MicrosoftGraphWorkbookTable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphWorkbookTable() *MicrosoftGraphWorkbookTable {
	this := MicrosoftGraphWorkbookTable{}
	return &this
}

// NewMicrosoftGraphWorkbookTableWithDefaults instantiates a new MicrosoftGraphWorkbookTable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphWorkbookTableWithDefaults() *MicrosoftGraphWorkbookTable {
	this := MicrosoftGraphWorkbookTable{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MicrosoftGraphWorkbookTable) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphWorkbookTable) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphWorkbookTable) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MicrosoftGraphWorkbookTable) SetId(v string) {
	o.Id = &v
}

// GetHighlightFirstColumn returns the HighlightFirstColumn field value if set, zero value otherwise.
func (o *MicrosoftGraphWorkbookTable) GetHighlightFirstColumn() bool {
	if o == nil || o.HighlightFirstColumn == nil {
		var ret bool
		return ret
	}
	return *o.HighlightFirstColumn
}

// GetHighlightFirstColumnOk returns a tuple with the HighlightFirstColumn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphWorkbookTable) GetHighlightFirstColumnOk() (*bool, bool) {
	if o == nil || o.HighlightFirstColumn == nil {
		return nil, false
	}
	return o.HighlightFirstColumn, true
}

// HasHighlightFirstColumn returns a boolean if a field has been set.
func (o *MicrosoftGraphWorkbookTable) HasHighlightFirstColumn() bool {
	if o != nil && o.HighlightFirstColumn != nil {
		return true
	}

	return false
}

// SetHighlightFirstColumn gets a reference to the given bool and assigns it to the HighlightFirstColumn field.
func (o *MicrosoftGraphWorkbookTable) SetHighlightFirstColumn(v bool) {
	o.HighlightFirstColumn = &v
}

// GetHighlightLastColumn returns the HighlightLastColumn field value if set, zero value otherwise.
func (o *MicrosoftGraphWorkbookTable) GetHighlightLastColumn() bool {
	if o == nil || o.HighlightLastColumn == nil {
		var ret bool
		return ret
	}
	return *o.HighlightLastColumn
}

// GetHighlightLastColumnOk returns a tuple with the HighlightLastColumn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphWorkbookTable) GetHighlightLastColumnOk() (*bool, bool) {
	if o == nil || o.HighlightLastColumn == nil {
		return nil, false
	}
	return o.HighlightLastColumn, true
}

// HasHighlightLastColumn returns a boolean if a field has been set.
func (o *MicrosoftGraphWorkbookTable) HasHighlightLastColumn() bool {
	if o != nil && o.HighlightLastColumn != nil {
		return true
	}

	return false
}

// SetHighlightLastColumn gets a reference to the given bool and assigns it to the HighlightLastColumn field.
func (o *MicrosoftGraphWorkbookTable) SetHighlightLastColumn(v bool) {
	o.HighlightLastColumn = &v
}

// GetLegacyId returns the LegacyId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphWorkbookTable) GetLegacyId() string {
	if o == nil || o.LegacyId.Get() == nil {
		var ret string
		return ret
	}
	return *o.LegacyId.Get()
}

// GetLegacyIdOk returns a tuple with the LegacyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphWorkbookTable) GetLegacyIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LegacyId.Get(), o.LegacyId.IsSet()
}

// HasLegacyId returns a boolean if a field has been set.
func (o *MicrosoftGraphWorkbookTable) HasLegacyId() bool {
	if o != nil && o.LegacyId.IsSet() {
		return true
	}

	return false
}

// SetLegacyId gets a reference to the given NullableString and assigns it to the LegacyId field.
func (o *MicrosoftGraphWorkbookTable) SetLegacyId(v string) {
	o.LegacyId.Set(&v)
}
// SetLegacyIdNil sets the value for LegacyId to be an explicit nil
func (o *MicrosoftGraphWorkbookTable) SetLegacyIdNil() {
	o.LegacyId.Set(nil)
}

// UnsetLegacyId ensures that no value is present for LegacyId, not even an explicit nil
func (o *MicrosoftGraphWorkbookTable) UnsetLegacyId() {
	o.LegacyId.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphWorkbookTable) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphWorkbookTable) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *MicrosoftGraphWorkbookTable) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *MicrosoftGraphWorkbookTable) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *MicrosoftGraphWorkbookTable) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *MicrosoftGraphWorkbookTable) UnsetName() {
	o.Name.Unset()
}

// GetShowBandedColumns returns the ShowBandedColumns field value if set, zero value otherwise.
func (o *MicrosoftGraphWorkbookTable) GetShowBandedColumns() bool {
	if o == nil || o.ShowBandedColumns == nil {
		var ret bool
		return ret
	}
	return *o.ShowBandedColumns
}

// GetShowBandedColumnsOk returns a tuple with the ShowBandedColumns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphWorkbookTable) GetShowBandedColumnsOk() (*bool, bool) {
	if o == nil || o.ShowBandedColumns == nil {
		return nil, false
	}
	return o.ShowBandedColumns, true
}

// HasShowBandedColumns returns a boolean if a field has been set.
func (o *MicrosoftGraphWorkbookTable) HasShowBandedColumns() bool {
	if o != nil && o.ShowBandedColumns != nil {
		return true
	}

	return false
}

// SetShowBandedColumns gets a reference to the given bool and assigns it to the ShowBandedColumns field.
func (o *MicrosoftGraphWorkbookTable) SetShowBandedColumns(v bool) {
	o.ShowBandedColumns = &v
}

// GetShowBandedRows returns the ShowBandedRows field value if set, zero value otherwise.
func (o *MicrosoftGraphWorkbookTable) GetShowBandedRows() bool {
	if o == nil || o.ShowBandedRows == nil {
		var ret bool
		return ret
	}
	return *o.ShowBandedRows
}

// GetShowBandedRowsOk returns a tuple with the ShowBandedRows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphWorkbookTable) GetShowBandedRowsOk() (*bool, bool) {
	if o == nil || o.ShowBandedRows == nil {
		return nil, false
	}
	return o.ShowBandedRows, true
}

// HasShowBandedRows returns a boolean if a field has been set.
func (o *MicrosoftGraphWorkbookTable) HasShowBandedRows() bool {
	if o != nil && o.ShowBandedRows != nil {
		return true
	}

	return false
}

// SetShowBandedRows gets a reference to the given bool and assigns it to the ShowBandedRows field.
func (o *MicrosoftGraphWorkbookTable) SetShowBandedRows(v bool) {
	o.ShowBandedRows = &v
}

// GetShowFilterButton returns the ShowFilterButton field value if set, zero value otherwise.
func (o *MicrosoftGraphWorkbookTable) GetShowFilterButton() bool {
	if o == nil || o.ShowFilterButton == nil {
		var ret bool
		return ret
	}
	return *o.ShowFilterButton
}

// GetShowFilterButtonOk returns a tuple with the ShowFilterButton field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphWorkbookTable) GetShowFilterButtonOk() (*bool, bool) {
	if o == nil || o.ShowFilterButton == nil {
		return nil, false
	}
	return o.ShowFilterButton, true
}

// HasShowFilterButton returns a boolean if a field has been set.
func (o *MicrosoftGraphWorkbookTable) HasShowFilterButton() bool {
	if o != nil && o.ShowFilterButton != nil {
		return true
	}

	return false
}

// SetShowFilterButton gets a reference to the given bool and assigns it to the ShowFilterButton field.
func (o *MicrosoftGraphWorkbookTable) SetShowFilterButton(v bool) {
	o.ShowFilterButton = &v
}

// GetShowHeaders returns the ShowHeaders field value if set, zero value otherwise.
func (o *MicrosoftGraphWorkbookTable) GetShowHeaders() bool {
	if o == nil || o.ShowHeaders == nil {
		var ret bool
		return ret
	}
	return *o.ShowHeaders
}

// GetShowHeadersOk returns a tuple with the ShowHeaders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphWorkbookTable) GetShowHeadersOk() (*bool, bool) {
	if o == nil || o.ShowHeaders == nil {
		return nil, false
	}
	return o.ShowHeaders, true
}

// HasShowHeaders returns a boolean if a field has been set.
func (o *MicrosoftGraphWorkbookTable) HasShowHeaders() bool {
	if o != nil && o.ShowHeaders != nil {
		return true
	}

	return false
}

// SetShowHeaders gets a reference to the given bool and assigns it to the ShowHeaders field.
func (o *MicrosoftGraphWorkbookTable) SetShowHeaders(v bool) {
	o.ShowHeaders = &v
}

// GetShowTotals returns the ShowTotals field value if set, zero value otherwise.
func (o *MicrosoftGraphWorkbookTable) GetShowTotals() bool {
	if o == nil || o.ShowTotals == nil {
		var ret bool
		return ret
	}
	return *o.ShowTotals
}

// GetShowTotalsOk returns a tuple with the ShowTotals field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphWorkbookTable) GetShowTotalsOk() (*bool, bool) {
	if o == nil || o.ShowTotals == nil {
		return nil, false
	}
	return o.ShowTotals, true
}

// HasShowTotals returns a boolean if a field has been set.
func (o *MicrosoftGraphWorkbookTable) HasShowTotals() bool {
	if o != nil && o.ShowTotals != nil {
		return true
	}

	return false
}

// SetShowTotals gets a reference to the given bool and assigns it to the ShowTotals field.
func (o *MicrosoftGraphWorkbookTable) SetShowTotals(v bool) {
	o.ShowTotals = &v
}

// GetStyle returns the Style field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphWorkbookTable) GetStyle() string {
	if o == nil || o.Style.Get() == nil {
		var ret string
		return ret
	}
	return *o.Style.Get()
}

// GetStyleOk returns a tuple with the Style field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphWorkbookTable) GetStyleOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Style.Get(), o.Style.IsSet()
}

// HasStyle returns a boolean if a field has been set.
func (o *MicrosoftGraphWorkbookTable) HasStyle() bool {
	if o != nil && o.Style.IsSet() {
		return true
	}

	return false
}

// SetStyle gets a reference to the given NullableString and assigns it to the Style field.
func (o *MicrosoftGraphWorkbookTable) SetStyle(v string) {
	o.Style.Set(&v)
}
// SetStyleNil sets the value for Style to be an explicit nil
func (o *MicrosoftGraphWorkbookTable) SetStyleNil() {
	o.Style.Set(nil)
}

// UnsetStyle ensures that no value is present for Style, not even an explicit nil
func (o *MicrosoftGraphWorkbookTable) UnsetStyle() {
	o.Style.Unset()
}

// GetColumns returns the Columns field value if set, zero value otherwise.
func (o *MicrosoftGraphWorkbookTable) GetColumns() []MicrosoftGraphWorkbookTableColumn {
	if o == nil || o.Columns == nil {
		var ret []MicrosoftGraphWorkbookTableColumn
		return ret
	}
	return *o.Columns
}

// GetColumnsOk returns a tuple with the Columns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphWorkbookTable) GetColumnsOk() (*[]MicrosoftGraphWorkbookTableColumn, bool) {
	if o == nil || o.Columns == nil {
		return nil, false
	}
	return o.Columns, true
}

// HasColumns returns a boolean if a field has been set.
func (o *MicrosoftGraphWorkbookTable) HasColumns() bool {
	if o != nil && o.Columns != nil {
		return true
	}

	return false
}

// SetColumns gets a reference to the given []MicrosoftGraphWorkbookTableColumn and assigns it to the Columns field.
func (o *MicrosoftGraphWorkbookTable) SetColumns(v []MicrosoftGraphWorkbookTableColumn) {
	o.Columns = &v
}

// GetRows returns the Rows field value if set, zero value otherwise.
func (o *MicrosoftGraphWorkbookTable) GetRows() []MicrosoftGraphWorkbookTableRow {
	if o == nil || o.Rows == nil {
		var ret []MicrosoftGraphWorkbookTableRow
		return ret
	}
	return *o.Rows
}

// GetRowsOk returns a tuple with the Rows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphWorkbookTable) GetRowsOk() (*[]MicrosoftGraphWorkbookTableRow, bool) {
	if o == nil || o.Rows == nil {
		return nil, false
	}
	return o.Rows, true
}

// HasRows returns a boolean if a field has been set.
func (o *MicrosoftGraphWorkbookTable) HasRows() bool {
	if o != nil && o.Rows != nil {
		return true
	}

	return false
}

// SetRows gets a reference to the given []MicrosoftGraphWorkbookTableRow and assigns it to the Rows field.
func (o *MicrosoftGraphWorkbookTable) SetRows(v []MicrosoftGraphWorkbookTableRow) {
	o.Rows = &v
}

// GetSort returns the Sort field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphWorkbookTable) GetSort() AnyOfmicrosoftGraphWorkbookTableSort {
	if o == nil  {
		var ret AnyOfmicrosoftGraphWorkbookTableSort
		return ret
	}
	return o.Sort
}

// GetSortOk returns a tuple with the Sort field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphWorkbookTable) GetSortOk() (*AnyOfmicrosoftGraphWorkbookTableSort, bool) {
	if o == nil || o.Sort == nil {
		return nil, false
	}
	return &o.Sort, true
}

// HasSort returns a boolean if a field has been set.
func (o *MicrosoftGraphWorkbookTable) HasSort() bool {
	if o != nil && o.Sort != nil {
		return true
	}

	return false
}

// SetSort gets a reference to the given AnyOfmicrosoftGraphWorkbookTableSort and assigns it to the Sort field.
func (o *MicrosoftGraphWorkbookTable) SetSort(v AnyOfmicrosoftGraphWorkbookTableSort) {
	o.Sort = v
}

// GetWorksheet returns the Worksheet field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphWorkbookTable) GetWorksheet() AnyOfmicrosoftGraphWorkbookWorksheet {
	if o == nil  {
		var ret AnyOfmicrosoftGraphWorkbookWorksheet
		return ret
	}
	return o.Worksheet
}

// GetWorksheetOk returns a tuple with the Worksheet field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphWorkbookTable) GetWorksheetOk() (*AnyOfmicrosoftGraphWorkbookWorksheet, bool) {
	if o == nil || o.Worksheet == nil {
		return nil, false
	}
	return &o.Worksheet, true
}

// HasWorksheet returns a boolean if a field has been set.
func (o *MicrosoftGraphWorkbookTable) HasWorksheet() bool {
	if o != nil && o.Worksheet != nil {
		return true
	}

	return false
}

// SetWorksheet gets a reference to the given AnyOfmicrosoftGraphWorkbookWorksheet and assigns it to the Worksheet field.
func (o *MicrosoftGraphWorkbookTable) SetWorksheet(v AnyOfmicrosoftGraphWorkbookWorksheet) {
	o.Worksheet = v
}

func (o MicrosoftGraphWorkbookTable) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
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

type NullableMicrosoftGraphWorkbookTable struct {
	value *MicrosoftGraphWorkbookTable
	isSet bool
}

func (v NullableMicrosoftGraphWorkbookTable) Get() *MicrosoftGraphWorkbookTable {
	return v.value
}

func (v *NullableMicrosoftGraphWorkbookTable) Set(val *MicrosoftGraphWorkbookTable) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphWorkbookTable) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphWorkbookTable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphWorkbookTable(val *MicrosoftGraphWorkbookTable) *NullableMicrosoftGraphWorkbookTable {
	return &NullableMicrosoftGraphWorkbookTable{value: val, isSet: true}
}

func (v NullableMicrosoftGraphWorkbookTable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphWorkbookTable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


