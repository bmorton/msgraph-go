# MicrosoftGraphWorkbook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**Application** | Pointer to [**AnyOfmicrosoftGraphWorkbookApplication**](anyOf&lt;microsoft.graph.workbookApplication&gt;.md) |  | [optional] 
**Comments** | Pointer to [**[]MicrosoftGraphWorkbookComment**](MicrosoftGraphWorkbookComment.md) |  | [optional] 
**Functions** | Pointer to [**AnyOfmicrosoftGraphWorkbookFunctions**](anyOf&lt;microsoft.graph.workbookFunctions&gt;.md) |  | [optional] 
**Names** | Pointer to [**[]MicrosoftGraphWorkbookNamedItem**](MicrosoftGraphWorkbookNamedItem.md) | Represents a collection of workbooks scoped named items (named ranges and constants). Read-only. | [optional] 
**Operations** | Pointer to [**[]MicrosoftGraphWorkbookOperation**](MicrosoftGraphWorkbookOperation.md) | The status of workbook operations. Getting an operation collection is not supported, but you can get the status of a long-running operation if the Location header is returned in the response. Read-only. | [optional] 
**Tables** | Pointer to [**[]MicrosoftGraphWorkbookTable**](MicrosoftGraphWorkbookTable.md) | Represents a collection of tables associated with the workbook. Read-only. | [optional] 
**Worksheets** | Pointer to [**[]MicrosoftGraphWorkbookWorksheet**](MicrosoftGraphWorkbookWorksheet.md) | Represents a collection of worksheets associated with the workbook. Read-only. | [optional] 

## Methods

### NewMicrosoftGraphWorkbook

`func NewMicrosoftGraphWorkbook() *MicrosoftGraphWorkbook`

NewMicrosoftGraphWorkbook instantiates a new MicrosoftGraphWorkbook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphWorkbookWithDefaults

`func NewMicrosoftGraphWorkbookWithDefaults() *MicrosoftGraphWorkbook`

NewMicrosoftGraphWorkbookWithDefaults instantiates a new MicrosoftGraphWorkbook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphWorkbook) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphWorkbook) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphWorkbook) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphWorkbook) HasId() bool`

HasId returns a boolean if a field has been set.

### GetApplication

`func (o *MicrosoftGraphWorkbook) GetApplication() AnyOfmicrosoftGraphWorkbookApplication`

GetApplication returns the Application field if non-nil, zero value otherwise.

### GetApplicationOk

`func (o *MicrosoftGraphWorkbook) GetApplicationOk() (*AnyOfmicrosoftGraphWorkbookApplication, bool)`

GetApplicationOk returns a tuple with the Application field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplication

`func (o *MicrosoftGraphWorkbook) SetApplication(v AnyOfmicrosoftGraphWorkbookApplication)`

SetApplication sets Application field to given value.

### HasApplication

`func (o *MicrosoftGraphWorkbook) HasApplication() bool`

HasApplication returns a boolean if a field has been set.

### SetApplicationNil

`func (o *MicrosoftGraphWorkbook) SetApplicationNil(b bool)`

 SetApplicationNil sets the value for Application to be an explicit nil

### UnsetApplication
`func (o *MicrosoftGraphWorkbook) UnsetApplication()`

UnsetApplication ensures that no value is present for Application, not even an explicit nil
### GetComments

`func (o *MicrosoftGraphWorkbook) GetComments() []MicrosoftGraphWorkbookComment`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *MicrosoftGraphWorkbook) GetCommentsOk() (*[]MicrosoftGraphWorkbookComment, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *MicrosoftGraphWorkbook) SetComments(v []MicrosoftGraphWorkbookComment)`

SetComments sets Comments field to given value.

### HasComments

`func (o *MicrosoftGraphWorkbook) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetFunctions

`func (o *MicrosoftGraphWorkbook) GetFunctions() AnyOfmicrosoftGraphWorkbookFunctions`

GetFunctions returns the Functions field if non-nil, zero value otherwise.

### GetFunctionsOk

`func (o *MicrosoftGraphWorkbook) GetFunctionsOk() (*AnyOfmicrosoftGraphWorkbookFunctions, bool)`

GetFunctionsOk returns a tuple with the Functions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctions

`func (o *MicrosoftGraphWorkbook) SetFunctions(v AnyOfmicrosoftGraphWorkbookFunctions)`

SetFunctions sets Functions field to given value.

### HasFunctions

`func (o *MicrosoftGraphWorkbook) HasFunctions() bool`

HasFunctions returns a boolean if a field has been set.

### SetFunctionsNil

`func (o *MicrosoftGraphWorkbook) SetFunctionsNil(b bool)`

 SetFunctionsNil sets the value for Functions to be an explicit nil

### UnsetFunctions
`func (o *MicrosoftGraphWorkbook) UnsetFunctions()`

UnsetFunctions ensures that no value is present for Functions, not even an explicit nil
### GetNames

`func (o *MicrosoftGraphWorkbook) GetNames() []MicrosoftGraphWorkbookNamedItem`

GetNames returns the Names field if non-nil, zero value otherwise.

### GetNamesOk

`func (o *MicrosoftGraphWorkbook) GetNamesOk() (*[]MicrosoftGraphWorkbookNamedItem, bool)`

GetNamesOk returns a tuple with the Names field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNames

`func (o *MicrosoftGraphWorkbook) SetNames(v []MicrosoftGraphWorkbookNamedItem)`

SetNames sets Names field to given value.

### HasNames

`func (o *MicrosoftGraphWorkbook) HasNames() bool`

HasNames returns a boolean if a field has been set.

### GetOperations

`func (o *MicrosoftGraphWorkbook) GetOperations() []MicrosoftGraphWorkbookOperation`

GetOperations returns the Operations field if non-nil, zero value otherwise.

### GetOperationsOk

`func (o *MicrosoftGraphWorkbook) GetOperationsOk() (*[]MicrosoftGraphWorkbookOperation, bool)`

GetOperationsOk returns a tuple with the Operations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperations

`func (o *MicrosoftGraphWorkbook) SetOperations(v []MicrosoftGraphWorkbookOperation)`

SetOperations sets Operations field to given value.

### HasOperations

`func (o *MicrosoftGraphWorkbook) HasOperations() bool`

HasOperations returns a boolean if a field has been set.

### GetTables

`func (o *MicrosoftGraphWorkbook) GetTables() []MicrosoftGraphWorkbookTable`

GetTables returns the Tables field if non-nil, zero value otherwise.

### GetTablesOk

`func (o *MicrosoftGraphWorkbook) GetTablesOk() (*[]MicrosoftGraphWorkbookTable, bool)`

GetTablesOk returns a tuple with the Tables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTables

`func (o *MicrosoftGraphWorkbook) SetTables(v []MicrosoftGraphWorkbookTable)`

SetTables sets Tables field to given value.

### HasTables

`func (o *MicrosoftGraphWorkbook) HasTables() bool`

HasTables returns a boolean if a field has been set.

### GetWorksheets

`func (o *MicrosoftGraphWorkbook) GetWorksheets() []MicrosoftGraphWorkbookWorksheet`

GetWorksheets returns the Worksheets field if non-nil, zero value otherwise.

### GetWorksheetsOk

`func (o *MicrosoftGraphWorkbook) GetWorksheetsOk() (*[]MicrosoftGraphWorkbookWorksheet, bool)`

GetWorksheetsOk returns a tuple with the Worksheets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorksheets

`func (o *MicrosoftGraphWorkbook) SetWorksheets(v []MicrosoftGraphWorkbookWorksheet)`

SetWorksheets sets Worksheets field to given value.

### HasWorksheets

`func (o *MicrosoftGraphWorkbook) HasWorksheets() bool`

HasWorksheets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


