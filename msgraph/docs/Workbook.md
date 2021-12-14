# Workbook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Application** | Pointer to [**AnyOfmicrosoftGraphWorkbookApplication**](anyOf&lt;microsoft.graph.workbookApplication&gt;.md) |  | [optional] 
**Comments** | Pointer to [**[]MicrosoftGraphWorkbookComment**](MicrosoftGraphWorkbookComment.md) |  | [optional] 
**Functions** | Pointer to [**AnyOfmicrosoftGraphWorkbookFunctions**](anyOf&lt;microsoft.graph.workbookFunctions&gt;.md) |  | [optional] 
**Names** | Pointer to [**[]MicrosoftGraphWorkbookNamedItem**](MicrosoftGraphWorkbookNamedItem.md) | Represents a collection of workbooks scoped named items (named ranges and constants). Read-only. | [optional] 
**Operations** | Pointer to [**[]MicrosoftGraphWorkbookOperation**](MicrosoftGraphWorkbookOperation.md) | The status of workbook operations. Getting an operation collection is not supported, but you can get the status of a long-running operation if the Location header is returned in the response. Read-only. | [optional] 
**Tables** | Pointer to [**[]MicrosoftGraphWorkbookTable**](MicrosoftGraphWorkbookTable.md) | Represents a collection of tables associated with the workbook. Read-only. | [optional] 
**Worksheets** | Pointer to [**[]MicrosoftGraphWorkbookWorksheet**](MicrosoftGraphWorkbookWorksheet.md) | Represents a collection of worksheets associated with the workbook. Read-only. | [optional] 

## Methods

### NewWorkbook

`func NewWorkbook() *Workbook`

NewWorkbook instantiates a new Workbook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkbookWithDefaults

`func NewWorkbookWithDefaults() *Workbook`

NewWorkbookWithDefaults instantiates a new Workbook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplication

`func (o *Workbook) GetApplication() AnyOfmicrosoftGraphWorkbookApplication`

GetApplication returns the Application field if non-nil, zero value otherwise.

### GetApplicationOk

`func (o *Workbook) GetApplicationOk() (*AnyOfmicrosoftGraphWorkbookApplication, bool)`

GetApplicationOk returns a tuple with the Application field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplication

`func (o *Workbook) SetApplication(v AnyOfmicrosoftGraphWorkbookApplication)`

SetApplication sets Application field to given value.

### HasApplication

`func (o *Workbook) HasApplication() bool`

HasApplication returns a boolean if a field has been set.

### SetApplicationNil

`func (o *Workbook) SetApplicationNil(b bool)`

 SetApplicationNil sets the value for Application to be an explicit nil

### UnsetApplication
`func (o *Workbook) UnsetApplication()`

UnsetApplication ensures that no value is present for Application, not even an explicit nil
### GetComments

`func (o *Workbook) GetComments() []MicrosoftGraphWorkbookComment`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *Workbook) GetCommentsOk() (*[]MicrosoftGraphWorkbookComment, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *Workbook) SetComments(v []MicrosoftGraphWorkbookComment)`

SetComments sets Comments field to given value.

### HasComments

`func (o *Workbook) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetFunctions

`func (o *Workbook) GetFunctions() AnyOfmicrosoftGraphWorkbookFunctions`

GetFunctions returns the Functions field if non-nil, zero value otherwise.

### GetFunctionsOk

`func (o *Workbook) GetFunctionsOk() (*AnyOfmicrosoftGraphWorkbookFunctions, bool)`

GetFunctionsOk returns a tuple with the Functions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctions

`func (o *Workbook) SetFunctions(v AnyOfmicrosoftGraphWorkbookFunctions)`

SetFunctions sets Functions field to given value.

### HasFunctions

`func (o *Workbook) HasFunctions() bool`

HasFunctions returns a boolean if a field has been set.

### SetFunctionsNil

`func (o *Workbook) SetFunctionsNil(b bool)`

 SetFunctionsNil sets the value for Functions to be an explicit nil

### UnsetFunctions
`func (o *Workbook) UnsetFunctions()`

UnsetFunctions ensures that no value is present for Functions, not even an explicit nil
### GetNames

`func (o *Workbook) GetNames() []MicrosoftGraphWorkbookNamedItem`

GetNames returns the Names field if non-nil, zero value otherwise.

### GetNamesOk

`func (o *Workbook) GetNamesOk() (*[]MicrosoftGraphWorkbookNamedItem, bool)`

GetNamesOk returns a tuple with the Names field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNames

`func (o *Workbook) SetNames(v []MicrosoftGraphWorkbookNamedItem)`

SetNames sets Names field to given value.

### HasNames

`func (o *Workbook) HasNames() bool`

HasNames returns a boolean if a field has been set.

### GetOperations

`func (o *Workbook) GetOperations() []MicrosoftGraphWorkbookOperation`

GetOperations returns the Operations field if non-nil, zero value otherwise.

### GetOperationsOk

`func (o *Workbook) GetOperationsOk() (*[]MicrosoftGraphWorkbookOperation, bool)`

GetOperationsOk returns a tuple with the Operations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperations

`func (o *Workbook) SetOperations(v []MicrosoftGraphWorkbookOperation)`

SetOperations sets Operations field to given value.

### HasOperations

`func (o *Workbook) HasOperations() bool`

HasOperations returns a boolean if a field has been set.

### GetTables

`func (o *Workbook) GetTables() []MicrosoftGraphWorkbookTable`

GetTables returns the Tables field if non-nil, zero value otherwise.

### GetTablesOk

`func (o *Workbook) GetTablesOk() (*[]MicrosoftGraphWorkbookTable, bool)`

GetTablesOk returns a tuple with the Tables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTables

`func (o *Workbook) SetTables(v []MicrosoftGraphWorkbookTable)`

SetTables sets Tables field to given value.

### HasTables

`func (o *Workbook) HasTables() bool`

HasTables returns a boolean if a field has been set.

### GetWorksheets

`func (o *Workbook) GetWorksheets() []MicrosoftGraphWorkbookWorksheet`

GetWorksheets returns the Worksheets field if non-nil, zero value otherwise.

### GetWorksheetsOk

`func (o *Workbook) GetWorksheetsOk() (*[]MicrosoftGraphWorkbookWorksheet, bool)`

GetWorksheetsOk returns a tuple with the Worksheets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorksheets

`func (o *Workbook) SetWorksheets(v []MicrosoftGraphWorkbookWorksheet)`

SetWorksheets sets Worksheets field to given value.

### HasWorksheets

`func (o *Workbook) HasWorksheets() bool`

HasWorksheets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


