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

// MicrosoftGraphWorkbookWorksheetProtectionOptions struct for MicrosoftGraphWorkbookWorksheetProtectionOptions
type MicrosoftGraphWorkbookWorksheetProtectionOptions struct {
	// Represents the worksheet protection option of allowing using auto filter feature.
	AllowAutoFilter *bool `json:"allowAutoFilter,omitempty"`
	// Represents the worksheet protection option of allowing deleting columns.
	AllowDeleteColumns *bool `json:"allowDeleteColumns,omitempty"`
	// Represents the worksheet protection option of allowing deleting rows.
	AllowDeleteRows *bool `json:"allowDeleteRows,omitempty"`
	// Represents the worksheet protection option of allowing formatting cells.
	AllowFormatCells *bool `json:"allowFormatCells,omitempty"`
	// Represents the worksheet protection option of allowing formatting columns.
	AllowFormatColumns *bool `json:"allowFormatColumns,omitempty"`
	// Represents the worksheet protection option of allowing formatting rows.
	AllowFormatRows *bool `json:"allowFormatRows,omitempty"`
	// Represents the worksheet protection option of allowing inserting columns.
	AllowInsertColumns *bool `json:"allowInsertColumns,omitempty"`
	// Represents the worksheet protection option of allowing inserting hyperlinks.
	AllowInsertHyperlinks *bool `json:"allowInsertHyperlinks,omitempty"`
	// Represents the worksheet protection option of allowing inserting rows.
	AllowInsertRows *bool `json:"allowInsertRows,omitempty"`
	// Represents the worksheet protection option of allowing using pivot table feature.
	AllowPivotTables *bool `json:"allowPivotTables,omitempty"`
	// Represents the worksheet protection option of allowing using sort feature.
	AllowSort *bool `json:"allowSort,omitempty"`
}

// NewMicrosoftGraphWorkbookWorksheetProtectionOptions instantiates a new MicrosoftGraphWorkbookWorksheetProtectionOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphWorkbookWorksheetProtectionOptions() *MicrosoftGraphWorkbookWorksheetProtectionOptions {
	this := MicrosoftGraphWorkbookWorksheetProtectionOptions{}
	return &this
}

// NewMicrosoftGraphWorkbookWorksheetProtectionOptionsWithDefaults instantiates a new MicrosoftGraphWorkbookWorksheetProtectionOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphWorkbookWorksheetProtectionOptionsWithDefaults() *MicrosoftGraphWorkbookWorksheetProtectionOptions {
	this := MicrosoftGraphWorkbookWorksheetProtectionOptions{}
	return &this
}

// GetAllowAutoFilter returns the AllowAutoFilter field value if set, zero value otherwise.
func (o *MicrosoftGraphWorkbookWorksheetProtectionOptions) GetAllowAutoFilter() bool {
	if o == nil || o.AllowAutoFilter == nil {
		var ret bool
		return ret
	}
	return *o.AllowAutoFilter
}

// GetAllowAutoFilterOk returns a tuple with the AllowAutoFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphWorkbookWorksheetProtectionOptions) GetAllowAutoFilterOk() (*bool, bool) {
	if o == nil || o.AllowAutoFilter == nil {
		return nil, false
	}
	return o.AllowAutoFilter, true
}

// HasAllowAutoFilter returns a boolean if a field has been set.
func (o *MicrosoftGraphWorkbookWorksheetProtectionOptions) HasAllowAutoFilter() bool {
	if o != nil && o.AllowAutoFilter != nil {
		return true
	}

	return false
}

// SetAllowAutoFilter gets a reference to the given bool and assigns it to the AllowAutoFilter field.
func (o *MicrosoftGraphWorkbookWorksheetProtectionOptions) SetAllowAutoFilter(v bool) {
	o.AllowAutoFilter = &v
}

// GetAllowDeleteColumns returns the AllowDeleteColumns field value if set, zero value otherwise.
func (o *MicrosoftGraphWorkbookWorksheetProtectionOptions) GetAllowDeleteColumns() bool {
	if o == nil || o.AllowDeleteColumns == nil {
		var ret bool
		return ret
	}
	return *o.AllowDeleteColumns
}

// GetAllowDeleteColumnsOk returns a tuple with the AllowDeleteColumns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphWorkbookWorksheetProtectionOptions) GetAllowDeleteColumnsOk() (*bool, bool) {
	if o == nil || o.AllowDeleteColumns == nil {
		return nil, false
	}
	return o.AllowDeleteColumns, true
}

// HasAllowDeleteColumns returns a boolean if a field has been set.
func (o *MicrosoftGraphWorkbookWorksheetProtectionOptions) HasAllowDeleteColumns() bool {
	if o != nil && o.AllowDeleteColumns != nil {
		return true
	}

	return false
}

// SetAllowDeleteColumns gets a reference to the given bool and assigns it to the AllowDeleteColumns field.
func (o *MicrosoftGraphWorkbookWorksheetProtectionOptions) SetAllowDeleteColumns(v bool) {
	o.AllowDeleteColumns = &v
}

// GetAllowDeleteRows returns the AllowDeleteRows field value if set, zero value otherwise.
func (o *MicrosoftGraphWorkbookWorksheetProtectionOptions) GetAllowDeleteRows() bool {
	if o == nil || o.AllowDeleteRows == nil {
		var ret bool
		return ret
	}
	return *o.AllowDeleteRows
}

// GetAllowDeleteRowsOk returns a tuple with the AllowDeleteRows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphWorkbookWorksheetProtectionOptions) GetAllowDeleteRowsOk() (*bool, bool) {
	if o == nil || o.AllowDeleteRows == nil {
		return nil, false
	}
	return o.AllowDeleteRows, true
}

// HasAllowDeleteRows returns a boolean if a field has been set.
func (o *MicrosoftGraphWorkbookWorksheetProtectionOptions) HasAllowDeleteRows() bool {
	if o != nil && o.AllowDeleteRows != nil {
		return true
	}

	return false
}

// SetAllowDeleteRows gets a reference to the given bool and assigns it to the AllowDeleteRows field.
func (o *MicrosoftGraphWorkbookWorksheetProtectionOptions) SetAllowDeleteRows(v bool) {
	o.AllowDeleteRows = &v
}

// GetAllowFormatCells returns the AllowFormatCells field value if set, zero value otherwise.
func (o *MicrosoftGraphWorkbookWorksheetProtectionOptions) GetAllowFormatCells() bool {
	if o == nil || o.AllowFormatCells == nil {
		var ret bool
		return ret
	}
	return *o.AllowFormatCells
}

// GetAllowFormatCellsOk returns a tuple with the AllowFormatCells field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphWorkbookWorksheetProtectionOptions) GetAllowFormatCellsOk() (*bool, bool) {
	if o == nil || o.AllowFormatCells == nil {
		return nil, false
	}
	return o.AllowFormatCells, true
}

// HasAllowFormatCells returns a boolean if a field has been set.
func (o *MicrosoftGraphWorkbookWorksheetProtectionOptions) HasAllowFormatCells() bool {
	if o != nil && o.AllowFormatCells != nil {
		return true
	}

	return false
}

// SetAllowFormatCells gets a reference to the given bool and assigns it to the AllowFormatCells field.
func (o *MicrosoftGraphWorkbookWorksheetProtectionOptions) SetAllowFormatCells(v bool) {
	o.AllowFormatCells = &v
}

// GetAllowFormatColumns returns the AllowFormatColumns field value if set, zero value otherwise.
func (o *MicrosoftGraphWorkbookWorksheetProtectionOptions) GetAllowFormatColumns() bool {
	if o == nil || o.AllowFormatColumns == nil {
		var ret bool
		return ret
	}
	return *o.AllowFormatColumns
}

// GetAllowFormatColumnsOk returns a tuple with the AllowFormatColumns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphWorkbookWorksheetProtectionOptions) GetAllowFormatColumnsOk() (*bool, bool) {
	if o == nil || o.AllowFormatColumns == nil {
		return nil, false
	}
	return o.AllowFormatColumns, true
}

// HasAllowFormatColumns returns a boolean if a field has been set.
func (o *MicrosoftGraphWorkbookWorksheetProtectionOptions) HasAllowFormatColumns() bool {
	if o != nil && o.AllowFormatColumns != nil {
		return true
	}

	return false
}

// SetAllowFormatColumns gets a reference to the given bool and assigns it to the AllowFormatColumns field.
func (o *MicrosoftGraphWorkbookWorksheetProtectionOptions) SetAllowFormatColumns(v bool) {
	o.AllowFormatColumns = &v
}

// GetAllowFormatRows returns the AllowFormatRows field value if set, zero value otherwise.
func (o *MicrosoftGraphWorkbookWorksheetProtectionOptions) GetAllowFormatRows() bool {
	if o == nil || o.AllowFormatRows == nil {
		var ret bool
		return ret
	}
	return *o.AllowFormatRows
}

// GetAllowFormatRowsOk returns a tuple with the AllowFormatRows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphWorkbookWorksheetProtectionOptions) GetAllowFormatRowsOk() (*bool, bool) {
	if o == nil || o.AllowFormatRows == nil {
		return nil, false
	}
	return o.AllowFormatRows, true
}

// HasAllowFormatRows returns a boolean if a field has been set.
func (o *MicrosoftGraphWorkbookWorksheetProtectionOptions) HasAllowFormatRows() bool {
	if o != nil && o.AllowFormatRows != nil {
		return true
	}

	return false
}

// SetAllowFormatRows gets a reference to the given bool and assigns it to the AllowFormatRows field.
func (o *MicrosoftGraphWorkbookWorksheetProtectionOptions) SetAllowFormatRows(v bool) {
	o.AllowFormatRows = &v
}

// GetAllowInsertColumns returns the AllowInsertColumns field value if set, zero value otherwise.
func (o *MicrosoftGraphWorkbookWorksheetProtectionOptions) GetAllowInsertColumns() bool {
	if o == nil || o.AllowInsertColumns == nil {
		var ret bool
		return ret
	}
	return *o.AllowInsertColumns
}

// GetAllowInsertColumnsOk returns a tuple with the AllowInsertColumns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphWorkbookWorksheetProtectionOptions) GetAllowInsertColumnsOk() (*bool, bool) {
	if o == nil || o.AllowInsertColumns == nil {
		return nil, false
	}
	return o.AllowInsertColumns, true
}

// HasAllowInsertColumns returns a boolean if a field has been set.
func (o *MicrosoftGraphWorkbookWorksheetProtectionOptions) HasAllowInsertColumns() bool {
	if o != nil && o.AllowInsertColumns != nil {
		return true
	}

	return false
}

// SetAllowInsertColumns gets a reference to the given bool and assigns it to the AllowInsertColumns field.
func (o *MicrosoftGraphWorkbookWorksheetProtectionOptions) SetAllowInsertColumns(v bool) {
	o.AllowInsertColumns = &v
}

// GetAllowInsertHyperlinks returns the AllowInsertHyperlinks field value if set, zero value otherwise.
func (o *MicrosoftGraphWorkbookWorksheetProtectionOptions) GetAllowInsertHyperlinks() bool {
	if o == nil || o.AllowInsertHyperlinks == nil {
		var ret bool
		return ret
	}
	return *o.AllowInsertHyperlinks
}

// GetAllowInsertHyperlinksOk returns a tuple with the AllowInsertHyperlinks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphWorkbookWorksheetProtectionOptions) GetAllowInsertHyperlinksOk() (*bool, bool) {
	if o == nil || o.AllowInsertHyperlinks == nil {
		return nil, false
	}
	return o.AllowInsertHyperlinks, true
}

// HasAllowInsertHyperlinks returns a boolean if a field has been set.
func (o *MicrosoftGraphWorkbookWorksheetProtectionOptions) HasAllowInsertHyperlinks() bool {
	if o != nil && o.AllowInsertHyperlinks != nil {
		return true
	}

	return false
}

// SetAllowInsertHyperlinks gets a reference to the given bool and assigns it to the AllowInsertHyperlinks field.
func (o *MicrosoftGraphWorkbookWorksheetProtectionOptions) SetAllowInsertHyperlinks(v bool) {
	o.AllowInsertHyperlinks = &v
}

// GetAllowInsertRows returns the AllowInsertRows field value if set, zero value otherwise.
func (o *MicrosoftGraphWorkbookWorksheetProtectionOptions) GetAllowInsertRows() bool {
	if o == nil || o.AllowInsertRows == nil {
		var ret bool
		return ret
	}
	return *o.AllowInsertRows
}

// GetAllowInsertRowsOk returns a tuple with the AllowInsertRows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphWorkbookWorksheetProtectionOptions) GetAllowInsertRowsOk() (*bool, bool) {
	if o == nil || o.AllowInsertRows == nil {
		return nil, false
	}
	return o.AllowInsertRows, true
}

// HasAllowInsertRows returns a boolean if a field has been set.
func (o *MicrosoftGraphWorkbookWorksheetProtectionOptions) HasAllowInsertRows() bool {
	if o != nil && o.AllowInsertRows != nil {
		return true
	}

	return false
}

// SetAllowInsertRows gets a reference to the given bool and assigns it to the AllowInsertRows field.
func (o *MicrosoftGraphWorkbookWorksheetProtectionOptions) SetAllowInsertRows(v bool) {
	o.AllowInsertRows = &v
}

// GetAllowPivotTables returns the AllowPivotTables field value if set, zero value otherwise.
func (o *MicrosoftGraphWorkbookWorksheetProtectionOptions) GetAllowPivotTables() bool {
	if o == nil || o.AllowPivotTables == nil {
		var ret bool
		return ret
	}
	return *o.AllowPivotTables
}

// GetAllowPivotTablesOk returns a tuple with the AllowPivotTables field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphWorkbookWorksheetProtectionOptions) GetAllowPivotTablesOk() (*bool, bool) {
	if o == nil || o.AllowPivotTables == nil {
		return nil, false
	}
	return o.AllowPivotTables, true
}

// HasAllowPivotTables returns a boolean if a field has been set.
func (o *MicrosoftGraphWorkbookWorksheetProtectionOptions) HasAllowPivotTables() bool {
	if o != nil && o.AllowPivotTables != nil {
		return true
	}

	return false
}

// SetAllowPivotTables gets a reference to the given bool and assigns it to the AllowPivotTables field.
func (o *MicrosoftGraphWorkbookWorksheetProtectionOptions) SetAllowPivotTables(v bool) {
	o.AllowPivotTables = &v
}

// GetAllowSort returns the AllowSort field value if set, zero value otherwise.
func (o *MicrosoftGraphWorkbookWorksheetProtectionOptions) GetAllowSort() bool {
	if o == nil || o.AllowSort == nil {
		var ret bool
		return ret
	}
	return *o.AllowSort
}

// GetAllowSortOk returns a tuple with the AllowSort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphWorkbookWorksheetProtectionOptions) GetAllowSortOk() (*bool, bool) {
	if o == nil || o.AllowSort == nil {
		return nil, false
	}
	return o.AllowSort, true
}

// HasAllowSort returns a boolean if a field has been set.
func (o *MicrosoftGraphWorkbookWorksheetProtectionOptions) HasAllowSort() bool {
	if o != nil && o.AllowSort != nil {
		return true
	}

	return false
}

// SetAllowSort gets a reference to the given bool and assigns it to the AllowSort field.
func (o *MicrosoftGraphWorkbookWorksheetProtectionOptions) SetAllowSort(v bool) {
	o.AllowSort = &v
}

func (o MicrosoftGraphWorkbookWorksheetProtectionOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AllowAutoFilter != nil {
		toSerialize["allowAutoFilter"] = o.AllowAutoFilter
	}
	if o.AllowDeleteColumns != nil {
		toSerialize["allowDeleteColumns"] = o.AllowDeleteColumns
	}
	if o.AllowDeleteRows != nil {
		toSerialize["allowDeleteRows"] = o.AllowDeleteRows
	}
	if o.AllowFormatCells != nil {
		toSerialize["allowFormatCells"] = o.AllowFormatCells
	}
	if o.AllowFormatColumns != nil {
		toSerialize["allowFormatColumns"] = o.AllowFormatColumns
	}
	if o.AllowFormatRows != nil {
		toSerialize["allowFormatRows"] = o.AllowFormatRows
	}
	if o.AllowInsertColumns != nil {
		toSerialize["allowInsertColumns"] = o.AllowInsertColumns
	}
	if o.AllowInsertHyperlinks != nil {
		toSerialize["allowInsertHyperlinks"] = o.AllowInsertHyperlinks
	}
	if o.AllowInsertRows != nil {
		toSerialize["allowInsertRows"] = o.AllowInsertRows
	}
	if o.AllowPivotTables != nil {
		toSerialize["allowPivotTables"] = o.AllowPivotTables
	}
	if o.AllowSort != nil {
		toSerialize["allowSort"] = o.AllowSort
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphWorkbookWorksheetProtectionOptions struct {
	value *MicrosoftGraphWorkbookWorksheetProtectionOptions
	isSet bool
}

func (v NullableMicrosoftGraphWorkbookWorksheetProtectionOptions) Get() *MicrosoftGraphWorkbookWorksheetProtectionOptions {
	return v.value
}

func (v *NullableMicrosoftGraphWorkbookWorksheetProtectionOptions) Set(val *MicrosoftGraphWorkbookWorksheetProtectionOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphWorkbookWorksheetProtectionOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphWorkbookWorksheetProtectionOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphWorkbookWorksheetProtectionOptions(val *MicrosoftGraphWorkbookWorksheetProtectionOptions) *NullableMicrosoftGraphWorkbookWorksheetProtectionOptions {
	return &NullableMicrosoftGraphWorkbookWorksheetProtectionOptions{value: val, isSet: true}
}

func (v NullableMicrosoftGraphWorkbookWorksheetProtectionOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphWorkbookWorksheetProtectionOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

