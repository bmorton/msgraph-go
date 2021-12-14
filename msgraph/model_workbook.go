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

// Workbook struct for Workbook
type Workbook struct {
	Application AnyOfmicrosoftGraphWorkbookApplication `json:"application,omitempty"`
	Comments *[]MicrosoftGraphWorkbookComment `json:"comments,omitempty"`
	Functions AnyOfmicrosoftGraphWorkbookFunctions `json:"functions,omitempty"`
	// Represents a collection of workbooks scoped named items (named ranges and constants). Read-only.
	Names *[]MicrosoftGraphWorkbookNamedItem `json:"names,omitempty"`
	// The status of workbook operations. Getting an operation collection is not supported, but you can get the status of a long-running operation if the Location header is returned in the response. Read-only.
	Operations *[]MicrosoftGraphWorkbookOperation `json:"operations,omitempty"`
	// Represents a collection of tables associated with the workbook. Read-only.
	Tables *[]MicrosoftGraphWorkbookTable `json:"tables,omitempty"`
	// Represents a collection of worksheets associated with the workbook. Read-only.
	Worksheets *[]MicrosoftGraphWorkbookWorksheet `json:"worksheets,omitempty"`
}

// NewWorkbook instantiates a new Workbook object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkbook() *Workbook {
	this := Workbook{}
	return &this
}

// NewWorkbookWithDefaults instantiates a new Workbook object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkbookWithDefaults() *Workbook {
	this := Workbook{}
	return &this
}

// GetApplication returns the Application field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Workbook) GetApplication() AnyOfmicrosoftGraphWorkbookApplication {
	if o == nil  {
		var ret AnyOfmicrosoftGraphWorkbookApplication
		return ret
	}
	return o.Application
}

// GetApplicationOk returns a tuple with the Application field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Workbook) GetApplicationOk() (*AnyOfmicrosoftGraphWorkbookApplication, bool) {
	if o == nil || o.Application == nil {
		return nil, false
	}
	return &o.Application, true
}

// HasApplication returns a boolean if a field has been set.
func (o *Workbook) HasApplication() bool {
	if o != nil && o.Application != nil {
		return true
	}

	return false
}

// SetApplication gets a reference to the given AnyOfmicrosoftGraphWorkbookApplication and assigns it to the Application field.
func (o *Workbook) SetApplication(v AnyOfmicrosoftGraphWorkbookApplication) {
	o.Application = v
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *Workbook) GetComments() []MicrosoftGraphWorkbookComment {
	if o == nil || o.Comments == nil {
		var ret []MicrosoftGraphWorkbookComment
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Workbook) GetCommentsOk() (*[]MicrosoftGraphWorkbookComment, bool) {
	if o == nil || o.Comments == nil {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *Workbook) HasComments() bool {
	if o != nil && o.Comments != nil {
		return true
	}

	return false
}

// SetComments gets a reference to the given []MicrosoftGraphWorkbookComment and assigns it to the Comments field.
func (o *Workbook) SetComments(v []MicrosoftGraphWorkbookComment) {
	o.Comments = &v
}

// GetFunctions returns the Functions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Workbook) GetFunctions() AnyOfmicrosoftGraphWorkbookFunctions {
	if o == nil  {
		var ret AnyOfmicrosoftGraphWorkbookFunctions
		return ret
	}
	return o.Functions
}

// GetFunctionsOk returns a tuple with the Functions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Workbook) GetFunctionsOk() (*AnyOfmicrosoftGraphWorkbookFunctions, bool) {
	if o == nil || o.Functions == nil {
		return nil, false
	}
	return &o.Functions, true
}

// HasFunctions returns a boolean if a field has been set.
func (o *Workbook) HasFunctions() bool {
	if o != nil && o.Functions != nil {
		return true
	}

	return false
}

// SetFunctions gets a reference to the given AnyOfmicrosoftGraphWorkbookFunctions and assigns it to the Functions field.
func (o *Workbook) SetFunctions(v AnyOfmicrosoftGraphWorkbookFunctions) {
	o.Functions = v
}

// GetNames returns the Names field value if set, zero value otherwise.
func (o *Workbook) GetNames() []MicrosoftGraphWorkbookNamedItem {
	if o == nil || o.Names == nil {
		var ret []MicrosoftGraphWorkbookNamedItem
		return ret
	}
	return *o.Names
}

// GetNamesOk returns a tuple with the Names field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Workbook) GetNamesOk() (*[]MicrosoftGraphWorkbookNamedItem, bool) {
	if o == nil || o.Names == nil {
		return nil, false
	}
	return o.Names, true
}

// HasNames returns a boolean if a field has been set.
func (o *Workbook) HasNames() bool {
	if o != nil && o.Names != nil {
		return true
	}

	return false
}

// SetNames gets a reference to the given []MicrosoftGraphWorkbookNamedItem and assigns it to the Names field.
func (o *Workbook) SetNames(v []MicrosoftGraphWorkbookNamedItem) {
	o.Names = &v
}

// GetOperations returns the Operations field value if set, zero value otherwise.
func (o *Workbook) GetOperations() []MicrosoftGraphWorkbookOperation {
	if o == nil || o.Operations == nil {
		var ret []MicrosoftGraphWorkbookOperation
		return ret
	}
	return *o.Operations
}

// GetOperationsOk returns a tuple with the Operations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Workbook) GetOperationsOk() (*[]MicrosoftGraphWorkbookOperation, bool) {
	if o == nil || o.Operations == nil {
		return nil, false
	}
	return o.Operations, true
}

// HasOperations returns a boolean if a field has been set.
func (o *Workbook) HasOperations() bool {
	if o != nil && o.Operations != nil {
		return true
	}

	return false
}

// SetOperations gets a reference to the given []MicrosoftGraphWorkbookOperation and assigns it to the Operations field.
func (o *Workbook) SetOperations(v []MicrosoftGraphWorkbookOperation) {
	o.Operations = &v
}

// GetTables returns the Tables field value if set, zero value otherwise.
func (o *Workbook) GetTables() []MicrosoftGraphWorkbookTable {
	if o == nil || o.Tables == nil {
		var ret []MicrosoftGraphWorkbookTable
		return ret
	}
	return *o.Tables
}

// GetTablesOk returns a tuple with the Tables field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Workbook) GetTablesOk() (*[]MicrosoftGraphWorkbookTable, bool) {
	if o == nil || o.Tables == nil {
		return nil, false
	}
	return o.Tables, true
}

// HasTables returns a boolean if a field has been set.
func (o *Workbook) HasTables() bool {
	if o != nil && o.Tables != nil {
		return true
	}

	return false
}

// SetTables gets a reference to the given []MicrosoftGraphWorkbookTable and assigns it to the Tables field.
func (o *Workbook) SetTables(v []MicrosoftGraphWorkbookTable) {
	o.Tables = &v
}

// GetWorksheets returns the Worksheets field value if set, zero value otherwise.
func (o *Workbook) GetWorksheets() []MicrosoftGraphWorkbookWorksheet {
	if o == nil || o.Worksheets == nil {
		var ret []MicrosoftGraphWorkbookWorksheet
		return ret
	}
	return *o.Worksheets
}

// GetWorksheetsOk returns a tuple with the Worksheets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Workbook) GetWorksheetsOk() (*[]MicrosoftGraphWorkbookWorksheet, bool) {
	if o == nil || o.Worksheets == nil {
		return nil, false
	}
	return o.Worksheets, true
}

// HasWorksheets returns a boolean if a field has been set.
func (o *Workbook) HasWorksheets() bool {
	if o != nil && o.Worksheets != nil {
		return true
	}

	return false
}

// SetWorksheets gets a reference to the given []MicrosoftGraphWorkbookWorksheet and assigns it to the Worksheets field.
func (o *Workbook) SetWorksheets(v []MicrosoftGraphWorkbookWorksheet) {
	o.Worksheets = &v
}

func (o Workbook) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Application != nil {
		toSerialize["application"] = o.Application
	}
	if o.Comments != nil {
		toSerialize["comments"] = o.Comments
	}
	if o.Functions != nil {
		toSerialize["functions"] = o.Functions
	}
	if o.Names != nil {
		toSerialize["names"] = o.Names
	}
	if o.Operations != nil {
		toSerialize["operations"] = o.Operations
	}
	if o.Tables != nil {
		toSerialize["tables"] = o.Tables
	}
	if o.Worksheets != nil {
		toSerialize["worksheets"] = o.Worksheets
	}
	return json.Marshal(toSerialize)
}

type NullableWorkbook struct {
	value *Workbook
	isSet bool
}

func (v NullableWorkbook) Get() *Workbook {
	return v.value
}

func (v *NullableWorkbook) Set(val *Workbook) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkbook) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkbook) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkbook(val *Workbook) *NullableWorkbook {
	return &NullableWorkbook{value: val, isSet: true}
}

func (v NullableWorkbook) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkbook) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


