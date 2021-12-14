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

// MicrosoftGraphWorkbookRange struct for MicrosoftGraphWorkbookRange
type MicrosoftGraphWorkbookRange struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
	// Represents the range reference in A1-style. Address value will contain the Sheet reference (e.g. Sheet1!A1:B4). Read-only.
	Address NullableString `json:"address,omitempty"`
	// Represents range reference for the specified range in the language of the user. Read-only.
	AddressLocal NullableString `json:"addressLocal,omitempty"`
	// Number of cells in the range. Read-only.
	CellCount *int32 `json:"cellCount,omitempty"`
	// Represents the total number of columns in the range. Read-only.
	ColumnCount *int32 `json:"columnCount,omitempty"`
	// Represents if all columns of the current range are hidden.
	ColumnHidden NullableBool `json:"columnHidden,omitempty"`
	// Represents the column number of the first cell in the range. Zero-indexed. Read-only.
	ColumnIndex *int32 `json:"columnIndex,omitempty"`
	// Represents the formula in A1-style notation.
	Formulas AnyOfobject `json:"formulas,omitempty"`
	// Represents the formula in A1-style notation, in the user's language and number-formatting locale.  For example, the English '=SUM(A1, 1.5)' formula would become '=SUMME(A1; 1,5)' in German.
	FormulasLocal AnyOfobject `json:"formulasLocal,omitempty"`
	// Represents the formula in R1C1-style notation.
	FormulasR1C1 AnyOfobject `json:"formulasR1C1,omitempty"`
	// Represents if all cells of the current range are hidden. Read-only.
	Hidden NullableBool `json:"hidden,omitempty"`
	// Represents Excel's number format code for the given cell.
	NumberFormat AnyOfobject `json:"numberFormat,omitempty"`
	// Returns the total number of rows in the range. Read-only.
	RowCount *int32 `json:"rowCount,omitempty"`
	// Represents if all rows of the current range are hidden.
	RowHidden NullableBool `json:"rowHidden,omitempty"`
	// Returns the row number of the first cell in the range. Zero-indexed. Read-only.
	RowIndex *int32 `json:"rowIndex,omitempty"`
	// Text values of the specified range. The Text value will not depend on the cell width. The # sign substitution that happens in Excel UI will not affect the text value returned by the API. Read-only.
	Text AnyOfobject `json:"text,omitempty"`
	// Represents the raw values of the specified range. The data returned could be of type string, number, or a boolean. Cell that contain an error will return the error string.
	Values AnyOfobject `json:"values,omitempty"`
	// Represents the type of data of each cell. The possible values are: Unknown, Empty, String, Integer, Double, Boolean, Error. Read-only.
	ValueTypes AnyOfobject `json:"valueTypes,omitempty"`
	// Returns a format object, encapsulating the range's font, fill, borders, alignment, and other properties. Read-only.
	Format AnyOfmicrosoftGraphWorkbookRangeFormat `json:"format,omitempty"`
	// The worksheet containing the current range. Read-only.
	Sort AnyOfmicrosoftGraphWorkbookRangeSort `json:"sort,omitempty"`
	// The worksheet containing the current range. Read-only.
	Worksheet AnyOfmicrosoftGraphWorkbookWorksheet `json:"worksheet,omitempty"`
}

// NewMicrosoftGraphWorkbookRange instantiates a new MicrosoftGraphWorkbookRange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphWorkbookRange() *MicrosoftGraphWorkbookRange {
	this := MicrosoftGraphWorkbookRange{}
	return &this
}

// NewMicrosoftGraphWorkbookRangeWithDefaults instantiates a new MicrosoftGraphWorkbookRange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphWorkbookRangeWithDefaults() *MicrosoftGraphWorkbookRange {
	this := MicrosoftGraphWorkbookRange{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MicrosoftGraphWorkbookRange) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphWorkbookRange) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphWorkbookRange) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MicrosoftGraphWorkbookRange) SetId(v string) {
	o.Id = &v
}

// GetAddress returns the Address field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphWorkbookRange) GetAddress() string {
	if o == nil || o.Address.Get() == nil {
		var ret string
		return ret
	}
	return *o.Address.Get()
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphWorkbookRange) GetAddressOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Address.Get(), o.Address.IsSet()
}

// HasAddress returns a boolean if a field has been set.
func (o *MicrosoftGraphWorkbookRange) HasAddress() bool {
	if o != nil && o.Address.IsSet() {
		return true
	}

	return false
}

// SetAddress gets a reference to the given NullableString and assigns it to the Address field.
func (o *MicrosoftGraphWorkbookRange) SetAddress(v string) {
	o.Address.Set(&v)
}
// SetAddressNil sets the value for Address to be an explicit nil
func (o *MicrosoftGraphWorkbookRange) SetAddressNil() {
	o.Address.Set(nil)
}

// UnsetAddress ensures that no value is present for Address, not even an explicit nil
func (o *MicrosoftGraphWorkbookRange) UnsetAddress() {
	o.Address.Unset()
}

// GetAddressLocal returns the AddressLocal field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphWorkbookRange) GetAddressLocal() string {
	if o == nil || o.AddressLocal.Get() == nil {
		var ret string
		return ret
	}
	return *o.AddressLocal.Get()
}

// GetAddressLocalOk returns a tuple with the AddressLocal field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphWorkbookRange) GetAddressLocalOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AddressLocal.Get(), o.AddressLocal.IsSet()
}

// HasAddressLocal returns a boolean if a field has been set.
func (o *MicrosoftGraphWorkbookRange) HasAddressLocal() bool {
	if o != nil && o.AddressLocal.IsSet() {
		return true
	}

	return false
}

// SetAddressLocal gets a reference to the given NullableString and assigns it to the AddressLocal field.
func (o *MicrosoftGraphWorkbookRange) SetAddressLocal(v string) {
	o.AddressLocal.Set(&v)
}
// SetAddressLocalNil sets the value for AddressLocal to be an explicit nil
func (o *MicrosoftGraphWorkbookRange) SetAddressLocalNil() {
	o.AddressLocal.Set(nil)
}

// UnsetAddressLocal ensures that no value is present for AddressLocal, not even an explicit nil
func (o *MicrosoftGraphWorkbookRange) UnsetAddressLocal() {
	o.AddressLocal.Unset()
}

// GetCellCount returns the CellCount field value if set, zero value otherwise.
func (o *MicrosoftGraphWorkbookRange) GetCellCount() int32 {
	if o == nil || o.CellCount == nil {
		var ret int32
		return ret
	}
	return *o.CellCount
}

// GetCellCountOk returns a tuple with the CellCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphWorkbookRange) GetCellCountOk() (*int32, bool) {
	if o == nil || o.CellCount == nil {
		return nil, false
	}
	return o.CellCount, true
}

// HasCellCount returns a boolean if a field has been set.
func (o *MicrosoftGraphWorkbookRange) HasCellCount() bool {
	if o != nil && o.CellCount != nil {
		return true
	}

	return false
}

// SetCellCount gets a reference to the given int32 and assigns it to the CellCount field.
func (o *MicrosoftGraphWorkbookRange) SetCellCount(v int32) {
	o.CellCount = &v
}

// GetColumnCount returns the ColumnCount field value if set, zero value otherwise.
func (o *MicrosoftGraphWorkbookRange) GetColumnCount() int32 {
	if o == nil || o.ColumnCount == nil {
		var ret int32
		return ret
	}
	return *o.ColumnCount
}

// GetColumnCountOk returns a tuple with the ColumnCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphWorkbookRange) GetColumnCountOk() (*int32, bool) {
	if o == nil || o.ColumnCount == nil {
		return nil, false
	}
	return o.ColumnCount, true
}

// HasColumnCount returns a boolean if a field has been set.
func (o *MicrosoftGraphWorkbookRange) HasColumnCount() bool {
	if o != nil && o.ColumnCount != nil {
		return true
	}

	return false
}

// SetColumnCount gets a reference to the given int32 and assigns it to the ColumnCount field.
func (o *MicrosoftGraphWorkbookRange) SetColumnCount(v int32) {
	o.ColumnCount = &v
}

// GetColumnHidden returns the ColumnHidden field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphWorkbookRange) GetColumnHidden() bool {
	if o == nil || o.ColumnHidden.Get() == nil {
		var ret bool
		return ret
	}
	return *o.ColumnHidden.Get()
}

// GetColumnHiddenOk returns a tuple with the ColumnHidden field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphWorkbookRange) GetColumnHiddenOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ColumnHidden.Get(), o.ColumnHidden.IsSet()
}

// HasColumnHidden returns a boolean if a field has been set.
func (o *MicrosoftGraphWorkbookRange) HasColumnHidden() bool {
	if o != nil && o.ColumnHidden.IsSet() {
		return true
	}

	return false
}

// SetColumnHidden gets a reference to the given NullableBool and assigns it to the ColumnHidden field.
func (o *MicrosoftGraphWorkbookRange) SetColumnHidden(v bool) {
	o.ColumnHidden.Set(&v)
}
// SetColumnHiddenNil sets the value for ColumnHidden to be an explicit nil
func (o *MicrosoftGraphWorkbookRange) SetColumnHiddenNil() {
	o.ColumnHidden.Set(nil)
}

// UnsetColumnHidden ensures that no value is present for ColumnHidden, not even an explicit nil
func (o *MicrosoftGraphWorkbookRange) UnsetColumnHidden() {
	o.ColumnHidden.Unset()
}

// GetColumnIndex returns the ColumnIndex field value if set, zero value otherwise.
func (o *MicrosoftGraphWorkbookRange) GetColumnIndex() int32 {
	if o == nil || o.ColumnIndex == nil {
		var ret int32
		return ret
	}
	return *o.ColumnIndex
}

// GetColumnIndexOk returns a tuple with the ColumnIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphWorkbookRange) GetColumnIndexOk() (*int32, bool) {
	if o == nil || o.ColumnIndex == nil {
		return nil, false
	}
	return o.ColumnIndex, true
}

// HasColumnIndex returns a boolean if a field has been set.
func (o *MicrosoftGraphWorkbookRange) HasColumnIndex() bool {
	if o != nil && o.ColumnIndex != nil {
		return true
	}

	return false
}

// SetColumnIndex gets a reference to the given int32 and assigns it to the ColumnIndex field.
func (o *MicrosoftGraphWorkbookRange) SetColumnIndex(v int32) {
	o.ColumnIndex = &v
}

// GetFormulas returns the Formulas field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphWorkbookRange) GetFormulas() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Formulas
}

// GetFormulasOk returns a tuple with the Formulas field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphWorkbookRange) GetFormulasOk() (*AnyOfobject, bool) {
	if o == nil || o.Formulas == nil {
		return nil, false
	}
	return &o.Formulas, true
}

// HasFormulas returns a boolean if a field has been set.
func (o *MicrosoftGraphWorkbookRange) HasFormulas() bool {
	if o != nil && o.Formulas != nil {
		return true
	}

	return false
}

// SetFormulas gets a reference to the given AnyOfobject and assigns it to the Formulas field.
func (o *MicrosoftGraphWorkbookRange) SetFormulas(v AnyOfobject) {
	o.Formulas = v
}

// GetFormulasLocal returns the FormulasLocal field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphWorkbookRange) GetFormulasLocal() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.FormulasLocal
}

// GetFormulasLocalOk returns a tuple with the FormulasLocal field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphWorkbookRange) GetFormulasLocalOk() (*AnyOfobject, bool) {
	if o == nil || o.FormulasLocal == nil {
		return nil, false
	}
	return &o.FormulasLocal, true
}

// HasFormulasLocal returns a boolean if a field has been set.
func (o *MicrosoftGraphWorkbookRange) HasFormulasLocal() bool {
	if o != nil && o.FormulasLocal != nil {
		return true
	}

	return false
}

// SetFormulasLocal gets a reference to the given AnyOfobject and assigns it to the FormulasLocal field.
func (o *MicrosoftGraphWorkbookRange) SetFormulasLocal(v AnyOfobject) {
	o.FormulasLocal = v
}

// GetFormulasR1C1 returns the FormulasR1C1 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphWorkbookRange) GetFormulasR1C1() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.FormulasR1C1
}

// GetFormulasR1C1Ok returns a tuple with the FormulasR1C1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphWorkbookRange) GetFormulasR1C1Ok() (*AnyOfobject, bool) {
	if o == nil || o.FormulasR1C1 == nil {
		return nil, false
	}
	return &o.FormulasR1C1, true
}

// HasFormulasR1C1 returns a boolean if a field has been set.
func (o *MicrosoftGraphWorkbookRange) HasFormulasR1C1() bool {
	if o != nil && o.FormulasR1C1 != nil {
		return true
	}

	return false
}

// SetFormulasR1C1 gets a reference to the given AnyOfobject and assigns it to the FormulasR1C1 field.
func (o *MicrosoftGraphWorkbookRange) SetFormulasR1C1(v AnyOfobject) {
	o.FormulasR1C1 = v
}

// GetHidden returns the Hidden field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphWorkbookRange) GetHidden() bool {
	if o == nil || o.Hidden.Get() == nil {
		var ret bool
		return ret
	}
	return *o.Hidden.Get()
}

// GetHiddenOk returns a tuple with the Hidden field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphWorkbookRange) GetHiddenOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Hidden.Get(), o.Hidden.IsSet()
}

// HasHidden returns a boolean if a field has been set.
func (o *MicrosoftGraphWorkbookRange) HasHidden() bool {
	if o != nil && o.Hidden.IsSet() {
		return true
	}

	return false
}

// SetHidden gets a reference to the given NullableBool and assigns it to the Hidden field.
func (o *MicrosoftGraphWorkbookRange) SetHidden(v bool) {
	o.Hidden.Set(&v)
}
// SetHiddenNil sets the value for Hidden to be an explicit nil
func (o *MicrosoftGraphWorkbookRange) SetHiddenNil() {
	o.Hidden.Set(nil)
}

// UnsetHidden ensures that no value is present for Hidden, not even an explicit nil
func (o *MicrosoftGraphWorkbookRange) UnsetHidden() {
	o.Hidden.Unset()
}

// GetNumberFormat returns the NumberFormat field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphWorkbookRange) GetNumberFormat() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.NumberFormat
}

// GetNumberFormatOk returns a tuple with the NumberFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphWorkbookRange) GetNumberFormatOk() (*AnyOfobject, bool) {
	if o == nil || o.NumberFormat == nil {
		return nil, false
	}
	return &o.NumberFormat, true
}

// HasNumberFormat returns a boolean if a field has been set.
func (o *MicrosoftGraphWorkbookRange) HasNumberFormat() bool {
	if o != nil && o.NumberFormat != nil {
		return true
	}

	return false
}

// SetNumberFormat gets a reference to the given AnyOfobject and assigns it to the NumberFormat field.
func (o *MicrosoftGraphWorkbookRange) SetNumberFormat(v AnyOfobject) {
	o.NumberFormat = v
}

// GetRowCount returns the RowCount field value if set, zero value otherwise.
func (o *MicrosoftGraphWorkbookRange) GetRowCount() int32 {
	if o == nil || o.RowCount == nil {
		var ret int32
		return ret
	}
	return *o.RowCount
}

// GetRowCountOk returns a tuple with the RowCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphWorkbookRange) GetRowCountOk() (*int32, bool) {
	if o == nil || o.RowCount == nil {
		return nil, false
	}
	return o.RowCount, true
}

// HasRowCount returns a boolean if a field has been set.
func (o *MicrosoftGraphWorkbookRange) HasRowCount() bool {
	if o != nil && o.RowCount != nil {
		return true
	}

	return false
}

// SetRowCount gets a reference to the given int32 and assigns it to the RowCount field.
func (o *MicrosoftGraphWorkbookRange) SetRowCount(v int32) {
	o.RowCount = &v
}

// GetRowHidden returns the RowHidden field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphWorkbookRange) GetRowHidden() bool {
	if o == nil || o.RowHidden.Get() == nil {
		var ret bool
		return ret
	}
	return *o.RowHidden.Get()
}

// GetRowHiddenOk returns a tuple with the RowHidden field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphWorkbookRange) GetRowHiddenOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RowHidden.Get(), o.RowHidden.IsSet()
}

// HasRowHidden returns a boolean if a field has been set.
func (o *MicrosoftGraphWorkbookRange) HasRowHidden() bool {
	if o != nil && o.RowHidden.IsSet() {
		return true
	}

	return false
}

// SetRowHidden gets a reference to the given NullableBool and assigns it to the RowHidden field.
func (o *MicrosoftGraphWorkbookRange) SetRowHidden(v bool) {
	o.RowHidden.Set(&v)
}
// SetRowHiddenNil sets the value for RowHidden to be an explicit nil
func (o *MicrosoftGraphWorkbookRange) SetRowHiddenNil() {
	o.RowHidden.Set(nil)
}

// UnsetRowHidden ensures that no value is present for RowHidden, not even an explicit nil
func (o *MicrosoftGraphWorkbookRange) UnsetRowHidden() {
	o.RowHidden.Unset()
}

// GetRowIndex returns the RowIndex field value if set, zero value otherwise.
func (o *MicrosoftGraphWorkbookRange) GetRowIndex() int32 {
	if o == nil || o.RowIndex == nil {
		var ret int32
		return ret
	}
	return *o.RowIndex
}

// GetRowIndexOk returns a tuple with the RowIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphWorkbookRange) GetRowIndexOk() (*int32, bool) {
	if o == nil || o.RowIndex == nil {
		return nil, false
	}
	return o.RowIndex, true
}

// HasRowIndex returns a boolean if a field has been set.
func (o *MicrosoftGraphWorkbookRange) HasRowIndex() bool {
	if o != nil && o.RowIndex != nil {
		return true
	}

	return false
}

// SetRowIndex gets a reference to the given int32 and assigns it to the RowIndex field.
func (o *MicrosoftGraphWorkbookRange) SetRowIndex(v int32) {
	o.RowIndex = &v
}

// GetText returns the Text field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphWorkbookRange) GetText() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphWorkbookRange) GetTextOk() (*AnyOfobject, bool) {
	if o == nil || o.Text == nil {
		return nil, false
	}
	return &o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *MicrosoftGraphWorkbookRange) HasText() bool {
	if o != nil && o.Text != nil {
		return true
	}

	return false
}

// SetText gets a reference to the given AnyOfobject and assigns it to the Text field.
func (o *MicrosoftGraphWorkbookRange) SetText(v AnyOfobject) {
	o.Text = v
}

// GetValues returns the Values field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphWorkbookRange) GetValues() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphWorkbookRange) GetValuesOk() (*AnyOfobject, bool) {
	if o == nil || o.Values == nil {
		return nil, false
	}
	return &o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *MicrosoftGraphWorkbookRange) HasValues() bool {
	if o != nil && o.Values != nil {
		return true
	}

	return false
}

// SetValues gets a reference to the given AnyOfobject and assigns it to the Values field.
func (o *MicrosoftGraphWorkbookRange) SetValues(v AnyOfobject) {
	o.Values = v
}

// GetValueTypes returns the ValueTypes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphWorkbookRange) GetValueTypes() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.ValueTypes
}

// GetValueTypesOk returns a tuple with the ValueTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphWorkbookRange) GetValueTypesOk() (*AnyOfobject, bool) {
	if o == nil || o.ValueTypes == nil {
		return nil, false
	}
	return &o.ValueTypes, true
}

// HasValueTypes returns a boolean if a field has been set.
func (o *MicrosoftGraphWorkbookRange) HasValueTypes() bool {
	if o != nil && o.ValueTypes != nil {
		return true
	}

	return false
}

// SetValueTypes gets a reference to the given AnyOfobject and assigns it to the ValueTypes field.
func (o *MicrosoftGraphWorkbookRange) SetValueTypes(v AnyOfobject) {
	o.ValueTypes = v
}

// GetFormat returns the Format field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphWorkbookRange) GetFormat() AnyOfmicrosoftGraphWorkbookRangeFormat {
	if o == nil  {
		var ret AnyOfmicrosoftGraphWorkbookRangeFormat
		return ret
	}
	return o.Format
}

// GetFormatOk returns a tuple with the Format field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphWorkbookRange) GetFormatOk() (*AnyOfmicrosoftGraphWorkbookRangeFormat, bool) {
	if o == nil || o.Format == nil {
		return nil, false
	}
	return &o.Format, true
}

// HasFormat returns a boolean if a field has been set.
func (o *MicrosoftGraphWorkbookRange) HasFormat() bool {
	if o != nil && o.Format != nil {
		return true
	}

	return false
}

// SetFormat gets a reference to the given AnyOfmicrosoftGraphWorkbookRangeFormat and assigns it to the Format field.
func (o *MicrosoftGraphWorkbookRange) SetFormat(v AnyOfmicrosoftGraphWorkbookRangeFormat) {
	o.Format = v
}

// GetSort returns the Sort field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphWorkbookRange) GetSort() AnyOfmicrosoftGraphWorkbookRangeSort {
	if o == nil  {
		var ret AnyOfmicrosoftGraphWorkbookRangeSort
		return ret
	}
	return o.Sort
}

// GetSortOk returns a tuple with the Sort field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphWorkbookRange) GetSortOk() (*AnyOfmicrosoftGraphWorkbookRangeSort, bool) {
	if o == nil || o.Sort == nil {
		return nil, false
	}
	return &o.Sort, true
}

// HasSort returns a boolean if a field has been set.
func (o *MicrosoftGraphWorkbookRange) HasSort() bool {
	if o != nil && o.Sort != nil {
		return true
	}

	return false
}

// SetSort gets a reference to the given AnyOfmicrosoftGraphWorkbookRangeSort and assigns it to the Sort field.
func (o *MicrosoftGraphWorkbookRange) SetSort(v AnyOfmicrosoftGraphWorkbookRangeSort) {
	o.Sort = v
}

// GetWorksheet returns the Worksheet field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphWorkbookRange) GetWorksheet() AnyOfmicrosoftGraphWorkbookWorksheet {
	if o == nil  {
		var ret AnyOfmicrosoftGraphWorkbookWorksheet
		return ret
	}
	return o.Worksheet
}

// GetWorksheetOk returns a tuple with the Worksheet field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphWorkbookRange) GetWorksheetOk() (*AnyOfmicrosoftGraphWorkbookWorksheet, bool) {
	if o == nil || o.Worksheet == nil {
		return nil, false
	}
	return &o.Worksheet, true
}

// HasWorksheet returns a boolean if a field has been set.
func (o *MicrosoftGraphWorkbookRange) HasWorksheet() bool {
	if o != nil && o.Worksheet != nil {
		return true
	}

	return false
}

// SetWorksheet gets a reference to the given AnyOfmicrosoftGraphWorkbookWorksheet and assigns it to the Worksheet field.
func (o *MicrosoftGraphWorkbookRange) SetWorksheet(v AnyOfmicrosoftGraphWorkbookWorksheet) {
	o.Worksheet = v
}

func (o MicrosoftGraphWorkbookRange) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Address.IsSet() {
		toSerialize["address"] = o.Address.Get()
	}
	if o.AddressLocal.IsSet() {
		toSerialize["addressLocal"] = o.AddressLocal.Get()
	}
	if o.CellCount != nil {
		toSerialize["cellCount"] = o.CellCount
	}
	if o.ColumnCount != nil {
		toSerialize["columnCount"] = o.ColumnCount
	}
	if o.ColumnHidden.IsSet() {
		toSerialize["columnHidden"] = o.ColumnHidden.Get()
	}
	if o.ColumnIndex != nil {
		toSerialize["columnIndex"] = o.ColumnIndex
	}
	if o.Formulas != nil {
		toSerialize["formulas"] = o.Formulas
	}
	if o.FormulasLocal != nil {
		toSerialize["formulasLocal"] = o.FormulasLocal
	}
	if o.FormulasR1C1 != nil {
		toSerialize["formulasR1C1"] = o.FormulasR1C1
	}
	if o.Hidden.IsSet() {
		toSerialize["hidden"] = o.Hidden.Get()
	}
	if o.NumberFormat != nil {
		toSerialize["numberFormat"] = o.NumberFormat
	}
	if o.RowCount != nil {
		toSerialize["rowCount"] = o.RowCount
	}
	if o.RowHidden.IsSet() {
		toSerialize["rowHidden"] = o.RowHidden.Get()
	}
	if o.RowIndex != nil {
		toSerialize["rowIndex"] = o.RowIndex
	}
	if o.Text != nil {
		toSerialize["text"] = o.Text
	}
	if o.Values != nil {
		toSerialize["values"] = o.Values
	}
	if o.ValueTypes != nil {
		toSerialize["valueTypes"] = o.ValueTypes
	}
	if o.Format != nil {
		toSerialize["format"] = o.Format
	}
	if o.Sort != nil {
		toSerialize["sort"] = o.Sort
	}
	if o.Worksheet != nil {
		toSerialize["worksheet"] = o.Worksheet
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphWorkbookRange struct {
	value *MicrosoftGraphWorkbookRange
	isSet bool
}

func (v NullableMicrosoftGraphWorkbookRange) Get() *MicrosoftGraphWorkbookRange {
	return v.value
}

func (v *NullableMicrosoftGraphWorkbookRange) Set(val *MicrosoftGraphWorkbookRange) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphWorkbookRange) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphWorkbookRange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphWorkbookRange(val *MicrosoftGraphWorkbookRange) *NullableMicrosoftGraphWorkbookRange {
	return &NullableMicrosoftGraphWorkbookRange{value: val, isSet: true}
}

func (v NullableMicrosoftGraphWorkbookRange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphWorkbookRange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


