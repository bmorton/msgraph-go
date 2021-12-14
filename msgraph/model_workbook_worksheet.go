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

// WorkbookWorksheet struct for WorkbookWorksheet
type WorkbookWorksheet struct {
	// The display name of the worksheet.
	Name NullableString `json:"name,omitempty"`
	// The zero-based position of the worksheet within the workbook.
	Position *int32 `json:"position,omitempty"`
	// The Visibility of the worksheet. The possible values are: Visible, Hidden, VeryHidden.
	Visibility *string `json:"visibility,omitempty"`
	// Returns collection of charts that are part of the worksheet. Read-only.
	Charts *[]MicrosoftGraphWorkbookChart `json:"charts,omitempty"`
	// Returns collection of names that are associated with the worksheet. Read-only.
	Names *[]MicrosoftGraphWorkbookNamedItem `json:"names,omitempty"`
	// Collection of PivotTables that are part of the worksheet.
	PivotTables *[]MicrosoftGraphWorkbookPivotTable `json:"pivotTables,omitempty"`
	// Returns sheet protection object for a worksheet. Read-only.
	Protection AnyOfmicrosoftGraphWorkbookWorksheetProtection `json:"protection,omitempty"`
	// Collection of tables that are part of the worksheet. Read-only.
	Tables *[]MicrosoftGraphWorkbookTable `json:"tables,omitempty"`
}

// NewWorkbookWorksheet instantiates a new WorkbookWorksheet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkbookWorksheet() *WorkbookWorksheet {
	this := WorkbookWorksheet{}
	return &this
}

// NewWorkbookWorksheetWithDefaults instantiates a new WorkbookWorksheet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkbookWorksheetWithDefaults() *WorkbookWorksheet {
	this := WorkbookWorksheet{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkbookWorksheet) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkbookWorksheet) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *WorkbookWorksheet) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *WorkbookWorksheet) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *WorkbookWorksheet) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *WorkbookWorksheet) UnsetName() {
	o.Name.Unset()
}

// GetPosition returns the Position field value if set, zero value otherwise.
func (o *WorkbookWorksheet) GetPosition() int32 {
	if o == nil || o.Position == nil {
		var ret int32
		return ret
	}
	return *o.Position
}

// GetPositionOk returns a tuple with the Position field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkbookWorksheet) GetPositionOk() (*int32, bool) {
	if o == nil || o.Position == nil {
		return nil, false
	}
	return o.Position, true
}

// HasPosition returns a boolean if a field has been set.
func (o *WorkbookWorksheet) HasPosition() bool {
	if o != nil && o.Position != nil {
		return true
	}

	return false
}

// SetPosition gets a reference to the given int32 and assigns it to the Position field.
func (o *WorkbookWorksheet) SetPosition(v int32) {
	o.Position = &v
}

// GetVisibility returns the Visibility field value if set, zero value otherwise.
func (o *WorkbookWorksheet) GetVisibility() string {
	if o == nil || o.Visibility == nil {
		var ret string
		return ret
	}
	return *o.Visibility
}

// GetVisibilityOk returns a tuple with the Visibility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkbookWorksheet) GetVisibilityOk() (*string, bool) {
	if o == nil || o.Visibility == nil {
		return nil, false
	}
	return o.Visibility, true
}

// HasVisibility returns a boolean if a field has been set.
func (o *WorkbookWorksheet) HasVisibility() bool {
	if o != nil && o.Visibility != nil {
		return true
	}

	return false
}

// SetVisibility gets a reference to the given string and assigns it to the Visibility field.
func (o *WorkbookWorksheet) SetVisibility(v string) {
	o.Visibility = &v
}

// GetCharts returns the Charts field value if set, zero value otherwise.
func (o *WorkbookWorksheet) GetCharts() []MicrosoftGraphWorkbookChart {
	if o == nil || o.Charts == nil {
		var ret []MicrosoftGraphWorkbookChart
		return ret
	}
	return *o.Charts
}

// GetChartsOk returns a tuple with the Charts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkbookWorksheet) GetChartsOk() (*[]MicrosoftGraphWorkbookChart, bool) {
	if o == nil || o.Charts == nil {
		return nil, false
	}
	return o.Charts, true
}

// HasCharts returns a boolean if a field has been set.
func (o *WorkbookWorksheet) HasCharts() bool {
	if o != nil && o.Charts != nil {
		return true
	}

	return false
}

// SetCharts gets a reference to the given []MicrosoftGraphWorkbookChart and assigns it to the Charts field.
func (o *WorkbookWorksheet) SetCharts(v []MicrosoftGraphWorkbookChart) {
	o.Charts = &v
}

// GetNames returns the Names field value if set, zero value otherwise.
func (o *WorkbookWorksheet) GetNames() []MicrosoftGraphWorkbookNamedItem {
	if o == nil || o.Names == nil {
		var ret []MicrosoftGraphWorkbookNamedItem
		return ret
	}
	return *o.Names
}

// GetNamesOk returns a tuple with the Names field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkbookWorksheet) GetNamesOk() (*[]MicrosoftGraphWorkbookNamedItem, bool) {
	if o == nil || o.Names == nil {
		return nil, false
	}
	return o.Names, true
}

// HasNames returns a boolean if a field has been set.
func (o *WorkbookWorksheet) HasNames() bool {
	if o != nil && o.Names != nil {
		return true
	}

	return false
}

// SetNames gets a reference to the given []MicrosoftGraphWorkbookNamedItem and assigns it to the Names field.
func (o *WorkbookWorksheet) SetNames(v []MicrosoftGraphWorkbookNamedItem) {
	o.Names = &v
}

// GetPivotTables returns the PivotTables field value if set, zero value otherwise.
func (o *WorkbookWorksheet) GetPivotTables() []MicrosoftGraphWorkbookPivotTable {
	if o == nil || o.PivotTables == nil {
		var ret []MicrosoftGraphWorkbookPivotTable
		return ret
	}
	return *o.PivotTables
}

// GetPivotTablesOk returns a tuple with the PivotTables field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkbookWorksheet) GetPivotTablesOk() (*[]MicrosoftGraphWorkbookPivotTable, bool) {
	if o == nil || o.PivotTables == nil {
		return nil, false
	}
	return o.PivotTables, true
}

// HasPivotTables returns a boolean if a field has been set.
func (o *WorkbookWorksheet) HasPivotTables() bool {
	if o != nil && o.PivotTables != nil {
		return true
	}

	return false
}

// SetPivotTables gets a reference to the given []MicrosoftGraphWorkbookPivotTable and assigns it to the PivotTables field.
func (o *WorkbookWorksheet) SetPivotTables(v []MicrosoftGraphWorkbookPivotTable) {
	o.PivotTables = &v
}

// GetProtection returns the Protection field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkbookWorksheet) GetProtection() AnyOfmicrosoftGraphWorkbookWorksheetProtection {
	if o == nil  {
		var ret AnyOfmicrosoftGraphWorkbookWorksheetProtection
		return ret
	}
	return o.Protection
}

// GetProtectionOk returns a tuple with the Protection field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkbookWorksheet) GetProtectionOk() (*AnyOfmicrosoftGraphWorkbookWorksheetProtection, bool) {
	if o == nil || o.Protection == nil {
		return nil, false
	}
	return &o.Protection, true
}

// HasProtection returns a boolean if a field has been set.
func (o *WorkbookWorksheet) HasProtection() bool {
	if o != nil && o.Protection != nil {
		return true
	}

	return false
}

// SetProtection gets a reference to the given AnyOfmicrosoftGraphWorkbookWorksheetProtection and assigns it to the Protection field.
func (o *WorkbookWorksheet) SetProtection(v AnyOfmicrosoftGraphWorkbookWorksheetProtection) {
	o.Protection = v
}

// GetTables returns the Tables field value if set, zero value otherwise.
func (o *WorkbookWorksheet) GetTables() []MicrosoftGraphWorkbookTable {
	if o == nil || o.Tables == nil {
		var ret []MicrosoftGraphWorkbookTable
		return ret
	}
	return *o.Tables
}

// GetTablesOk returns a tuple with the Tables field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkbookWorksheet) GetTablesOk() (*[]MicrosoftGraphWorkbookTable, bool) {
	if o == nil || o.Tables == nil {
		return nil, false
	}
	return o.Tables, true
}

// HasTables returns a boolean if a field has been set.
func (o *WorkbookWorksheet) HasTables() bool {
	if o != nil && o.Tables != nil {
		return true
	}

	return false
}

// SetTables gets a reference to the given []MicrosoftGraphWorkbookTable and assigns it to the Tables field.
func (o *WorkbookWorksheet) SetTables(v []MicrosoftGraphWorkbookTable) {
	o.Tables = &v
}

func (o WorkbookWorksheet) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Position != nil {
		toSerialize["position"] = o.Position
	}
	if o.Visibility != nil {
		toSerialize["visibility"] = o.Visibility
	}
	if o.Charts != nil {
		toSerialize["charts"] = o.Charts
	}
	if o.Names != nil {
		toSerialize["names"] = o.Names
	}
	if o.PivotTables != nil {
		toSerialize["pivotTables"] = o.PivotTables
	}
	if o.Protection != nil {
		toSerialize["protection"] = o.Protection
	}
	if o.Tables != nil {
		toSerialize["tables"] = o.Tables
	}
	return json.Marshal(toSerialize)
}

type NullableWorkbookWorksheet struct {
	value *WorkbookWorksheet
	isSet bool
}

func (v NullableWorkbookWorksheet) Get() *WorkbookWorksheet {
	return v.value
}

func (v *NullableWorkbookWorksheet) Set(val *WorkbookWorksheet) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkbookWorksheet) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkbookWorksheet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkbookWorksheet(val *WorkbookWorksheet) *NullableWorkbookWorksheet {
	return &NullableWorkbookWorksheet{value: val, isSet: true}
}

func (v NullableWorkbookWorksheet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkbookWorksheet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

