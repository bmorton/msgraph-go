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

// MicrosoftGraphWorkbook struct for MicrosoftGraphWorkbook
type MicrosoftGraphWorkbook struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
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

// NewMicrosoftGraphWorkbook instantiates a new MicrosoftGraphWorkbook object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphWorkbook() *MicrosoftGraphWorkbook {
	this := MicrosoftGraphWorkbook{}
	return &this
}

// NewMicrosoftGraphWorkbookWithDefaults instantiates a new MicrosoftGraphWorkbook object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphWorkbookWithDefaults() *MicrosoftGraphWorkbook {
	this := MicrosoftGraphWorkbook{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MicrosoftGraphWorkbook) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphWorkbook) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphWorkbook) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MicrosoftGraphWorkbook) SetId(v string) {
	o.Id = &v
}

// GetApplication returns the Application field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphWorkbook) GetApplication() AnyOfmicrosoftGraphWorkbookApplication {
	if o == nil  {
		var ret AnyOfmicrosoftGraphWorkbookApplication
		return ret
	}
	return o.Application
}

// GetApplicationOk returns a tuple with the Application field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphWorkbook) GetApplicationOk() (*AnyOfmicrosoftGraphWorkbookApplication, bool) {
	if o == nil || o.Application == nil {
		return nil, false
	}
	return &o.Application, true
}

// HasApplication returns a boolean if a field has been set.
func (o *MicrosoftGraphWorkbook) HasApplication() bool {
	if o != nil && o.Application != nil {
		return true
	}

	return false
}

// SetApplication gets a reference to the given AnyOfmicrosoftGraphWorkbookApplication and assigns it to the Application field.
func (o *MicrosoftGraphWorkbook) SetApplication(v AnyOfmicrosoftGraphWorkbookApplication) {
	o.Application = v
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *MicrosoftGraphWorkbook) GetComments() []MicrosoftGraphWorkbookComment {
	if o == nil || o.Comments == nil {
		var ret []MicrosoftGraphWorkbookComment
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphWorkbook) GetCommentsOk() (*[]MicrosoftGraphWorkbookComment, bool) {
	if o == nil || o.Comments == nil {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *MicrosoftGraphWorkbook) HasComments() bool {
	if o != nil && o.Comments != nil {
		return true
	}

	return false
}

// SetComments gets a reference to the given []MicrosoftGraphWorkbookComment and assigns it to the Comments field.
func (o *MicrosoftGraphWorkbook) SetComments(v []MicrosoftGraphWorkbookComment) {
	o.Comments = &v
}

// GetFunctions returns the Functions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphWorkbook) GetFunctions() AnyOfmicrosoftGraphWorkbookFunctions {
	if o == nil  {
		var ret AnyOfmicrosoftGraphWorkbookFunctions
		return ret
	}
	return o.Functions
}

// GetFunctionsOk returns a tuple with the Functions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphWorkbook) GetFunctionsOk() (*AnyOfmicrosoftGraphWorkbookFunctions, bool) {
	if o == nil || o.Functions == nil {
		return nil, false
	}
	return &o.Functions, true
}

// HasFunctions returns a boolean if a field has been set.
func (o *MicrosoftGraphWorkbook) HasFunctions() bool {
	if o != nil && o.Functions != nil {
		return true
	}

	return false
}

// SetFunctions gets a reference to the given AnyOfmicrosoftGraphWorkbookFunctions and assigns it to the Functions field.
func (o *MicrosoftGraphWorkbook) SetFunctions(v AnyOfmicrosoftGraphWorkbookFunctions) {
	o.Functions = v
}

// GetNames returns the Names field value if set, zero value otherwise.
func (o *MicrosoftGraphWorkbook) GetNames() []MicrosoftGraphWorkbookNamedItem {
	if o == nil || o.Names == nil {
		var ret []MicrosoftGraphWorkbookNamedItem
		return ret
	}
	return *o.Names
}

// GetNamesOk returns a tuple with the Names field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphWorkbook) GetNamesOk() (*[]MicrosoftGraphWorkbookNamedItem, bool) {
	if o == nil || o.Names == nil {
		return nil, false
	}
	return o.Names, true
}

// HasNames returns a boolean if a field has been set.
func (o *MicrosoftGraphWorkbook) HasNames() bool {
	if o != nil && o.Names != nil {
		return true
	}

	return false
}

// SetNames gets a reference to the given []MicrosoftGraphWorkbookNamedItem and assigns it to the Names field.
func (o *MicrosoftGraphWorkbook) SetNames(v []MicrosoftGraphWorkbookNamedItem) {
	o.Names = &v
}

// GetOperations returns the Operations field value if set, zero value otherwise.
func (o *MicrosoftGraphWorkbook) GetOperations() []MicrosoftGraphWorkbookOperation {
	if o == nil || o.Operations == nil {
		var ret []MicrosoftGraphWorkbookOperation
		return ret
	}
	return *o.Operations
}

// GetOperationsOk returns a tuple with the Operations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphWorkbook) GetOperationsOk() (*[]MicrosoftGraphWorkbookOperation, bool) {
	if o == nil || o.Operations == nil {
		return nil, false
	}
	return o.Operations, true
}

// HasOperations returns a boolean if a field has been set.
func (o *MicrosoftGraphWorkbook) HasOperations() bool {
	if o != nil && o.Operations != nil {
		return true
	}

	return false
}

// SetOperations gets a reference to the given []MicrosoftGraphWorkbookOperation and assigns it to the Operations field.
func (o *MicrosoftGraphWorkbook) SetOperations(v []MicrosoftGraphWorkbookOperation) {
	o.Operations = &v
}

// GetTables returns the Tables field value if set, zero value otherwise.
func (o *MicrosoftGraphWorkbook) GetTables() []MicrosoftGraphWorkbookTable {
	if o == nil || o.Tables == nil {
		var ret []MicrosoftGraphWorkbookTable
		return ret
	}
	return *o.Tables
}

// GetTablesOk returns a tuple with the Tables field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphWorkbook) GetTablesOk() (*[]MicrosoftGraphWorkbookTable, bool) {
	if o == nil || o.Tables == nil {
		return nil, false
	}
	return o.Tables, true
}

// HasTables returns a boolean if a field has been set.
func (o *MicrosoftGraphWorkbook) HasTables() bool {
	if o != nil && o.Tables != nil {
		return true
	}

	return false
}

// SetTables gets a reference to the given []MicrosoftGraphWorkbookTable and assigns it to the Tables field.
func (o *MicrosoftGraphWorkbook) SetTables(v []MicrosoftGraphWorkbookTable) {
	o.Tables = &v
}

// GetWorksheets returns the Worksheets field value if set, zero value otherwise.
func (o *MicrosoftGraphWorkbook) GetWorksheets() []MicrosoftGraphWorkbookWorksheet {
	if o == nil || o.Worksheets == nil {
		var ret []MicrosoftGraphWorkbookWorksheet
		return ret
	}
	return *o.Worksheets
}

// GetWorksheetsOk returns a tuple with the Worksheets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphWorkbook) GetWorksheetsOk() (*[]MicrosoftGraphWorkbookWorksheet, bool) {
	if o == nil || o.Worksheets == nil {
		return nil, false
	}
	return o.Worksheets, true
}

// HasWorksheets returns a boolean if a field has been set.
func (o *MicrosoftGraphWorkbook) HasWorksheets() bool {
	if o != nil && o.Worksheets != nil {
		return true
	}

	return false
}

// SetWorksheets gets a reference to the given []MicrosoftGraphWorkbookWorksheet and assigns it to the Worksheets field.
func (o *MicrosoftGraphWorkbook) SetWorksheets(v []MicrosoftGraphWorkbookWorksheet) {
	o.Worksheets = &v
}

func (o MicrosoftGraphWorkbook) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
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

type NullableMicrosoftGraphWorkbook struct {
	value *MicrosoftGraphWorkbook
	isSet bool
}

func (v NullableMicrosoftGraphWorkbook) Get() *MicrosoftGraphWorkbook {
	return v.value
}

func (v *NullableMicrosoftGraphWorkbook) Set(val *MicrosoftGraphWorkbook) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphWorkbook) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphWorkbook) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphWorkbook(val *MicrosoftGraphWorkbook) *NullableMicrosoftGraphWorkbook {
	return &NullableMicrosoftGraphWorkbook{value: val, isSet: true}
}

func (v NullableMicrosoftGraphWorkbook) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphWorkbook) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


