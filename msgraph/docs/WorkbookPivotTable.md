# WorkbookPivotTable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | Name of the PivotTable. | [optional] 
**Worksheet** | Pointer to [**AnyOfmicrosoftGraphWorkbookWorksheet**](anyOf&lt;microsoft.graph.workbookWorksheet&gt;.md) | The worksheet containing the current PivotTable. Read-only. | [optional] 

## Methods

### NewWorkbookPivotTable

`func NewWorkbookPivotTable() *WorkbookPivotTable`

NewWorkbookPivotTable instantiates a new WorkbookPivotTable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkbookPivotTableWithDefaults

`func NewWorkbookPivotTableWithDefaults() *WorkbookPivotTable`

NewWorkbookPivotTableWithDefaults instantiates a new WorkbookPivotTable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *WorkbookPivotTable) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkbookPivotTable) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkbookPivotTable) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WorkbookPivotTable) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *WorkbookPivotTable) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *WorkbookPivotTable) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetWorksheet

`func (o *WorkbookPivotTable) GetWorksheet() AnyOfmicrosoftGraphWorkbookWorksheet`

GetWorksheet returns the Worksheet field if non-nil, zero value otherwise.

### GetWorksheetOk

`func (o *WorkbookPivotTable) GetWorksheetOk() (*AnyOfmicrosoftGraphWorkbookWorksheet, bool)`

GetWorksheetOk returns a tuple with the Worksheet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorksheet

`func (o *WorkbookPivotTable) SetWorksheet(v AnyOfmicrosoftGraphWorkbookWorksheet)`

SetWorksheet sets Worksheet field to given value.

### HasWorksheet

`func (o *WorkbookPivotTable) HasWorksheet() bool`

HasWorksheet returns a boolean if a field has been set.

### SetWorksheetNil

`func (o *WorkbookPivotTable) SetWorksheetNil(b bool)`

 SetWorksheetNil sets the value for Worksheet to be an explicit nil

### UnsetWorksheet
`func (o *WorkbookPivotTable) UnsetWorksheet()`

UnsetWorksheet ensures that no value is present for Worksheet, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


