# WorkbookWorksheet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | The display name of the worksheet. | [optional] 
**Position** | Pointer to **int32** | The zero-based position of the worksheet within the workbook. | [optional] 
**Visibility** | Pointer to **string** | The Visibility of the worksheet. The possible values are: Visible, Hidden, VeryHidden. | [optional] 
**Charts** | Pointer to [**[]MicrosoftGraphWorkbookChart**](MicrosoftGraphWorkbookChart.md) | Returns collection of charts that are part of the worksheet. Read-only. | [optional] 
**Names** | Pointer to [**[]MicrosoftGraphWorkbookNamedItem**](MicrosoftGraphWorkbookNamedItem.md) | Returns collection of names that are associated with the worksheet. Read-only. | [optional] 
**PivotTables** | Pointer to [**[]MicrosoftGraphWorkbookPivotTable**](MicrosoftGraphWorkbookPivotTable.md) | Collection of PivotTables that are part of the worksheet. | [optional] 
**Protection** | Pointer to [**AnyOfmicrosoftGraphWorkbookWorksheetProtection**](anyOf&lt;microsoft.graph.workbookWorksheetProtection&gt;.md) | Returns sheet protection object for a worksheet. Read-only. | [optional] 
**Tables** | Pointer to [**[]MicrosoftGraphWorkbookTable**](MicrosoftGraphWorkbookTable.md) | Collection of tables that are part of the worksheet. Read-only. | [optional] 

## Methods

### NewWorkbookWorksheet

`func NewWorkbookWorksheet() *WorkbookWorksheet`

NewWorkbookWorksheet instantiates a new WorkbookWorksheet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkbookWorksheetWithDefaults

`func NewWorkbookWorksheetWithDefaults() *WorkbookWorksheet`

NewWorkbookWorksheetWithDefaults instantiates a new WorkbookWorksheet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *WorkbookWorksheet) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkbookWorksheet) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkbookWorksheet) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WorkbookWorksheet) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *WorkbookWorksheet) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *WorkbookWorksheet) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetPosition

`func (o *WorkbookWorksheet) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *WorkbookWorksheet) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *WorkbookWorksheet) SetPosition(v int32)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *WorkbookWorksheet) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetVisibility

`func (o *WorkbookWorksheet) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *WorkbookWorksheet) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *WorkbookWorksheet) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *WorkbookWorksheet) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetCharts

`func (o *WorkbookWorksheet) GetCharts() []MicrosoftGraphWorkbookChart`

GetCharts returns the Charts field if non-nil, zero value otherwise.

### GetChartsOk

`func (o *WorkbookWorksheet) GetChartsOk() (*[]MicrosoftGraphWorkbookChart, bool)`

GetChartsOk returns a tuple with the Charts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharts

`func (o *WorkbookWorksheet) SetCharts(v []MicrosoftGraphWorkbookChart)`

SetCharts sets Charts field to given value.

### HasCharts

`func (o *WorkbookWorksheet) HasCharts() bool`

HasCharts returns a boolean if a field has been set.

### GetNames

`func (o *WorkbookWorksheet) GetNames() []MicrosoftGraphWorkbookNamedItem`

GetNames returns the Names field if non-nil, zero value otherwise.

### GetNamesOk

`func (o *WorkbookWorksheet) GetNamesOk() (*[]MicrosoftGraphWorkbookNamedItem, bool)`

GetNamesOk returns a tuple with the Names field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNames

`func (o *WorkbookWorksheet) SetNames(v []MicrosoftGraphWorkbookNamedItem)`

SetNames sets Names field to given value.

### HasNames

`func (o *WorkbookWorksheet) HasNames() bool`

HasNames returns a boolean if a field has been set.

### GetPivotTables

`func (o *WorkbookWorksheet) GetPivotTables() []MicrosoftGraphWorkbookPivotTable`

GetPivotTables returns the PivotTables field if non-nil, zero value otherwise.

### GetPivotTablesOk

`func (o *WorkbookWorksheet) GetPivotTablesOk() (*[]MicrosoftGraphWorkbookPivotTable, bool)`

GetPivotTablesOk returns a tuple with the PivotTables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPivotTables

`func (o *WorkbookWorksheet) SetPivotTables(v []MicrosoftGraphWorkbookPivotTable)`

SetPivotTables sets PivotTables field to given value.

### HasPivotTables

`func (o *WorkbookWorksheet) HasPivotTables() bool`

HasPivotTables returns a boolean if a field has been set.

### GetProtection

`func (o *WorkbookWorksheet) GetProtection() AnyOfmicrosoftGraphWorkbookWorksheetProtection`

GetProtection returns the Protection field if non-nil, zero value otherwise.

### GetProtectionOk

`func (o *WorkbookWorksheet) GetProtectionOk() (*AnyOfmicrosoftGraphWorkbookWorksheetProtection, bool)`

GetProtectionOk returns a tuple with the Protection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtection

`func (o *WorkbookWorksheet) SetProtection(v AnyOfmicrosoftGraphWorkbookWorksheetProtection)`

SetProtection sets Protection field to given value.

### HasProtection

`func (o *WorkbookWorksheet) HasProtection() bool`

HasProtection returns a boolean if a field has been set.

### SetProtectionNil

`func (o *WorkbookWorksheet) SetProtectionNil(b bool)`

 SetProtectionNil sets the value for Protection to be an explicit nil

### UnsetProtection
`func (o *WorkbookWorksheet) UnsetProtection()`

UnsetProtection ensures that no value is present for Protection, not even an explicit nil
### GetTables

`func (o *WorkbookWorksheet) GetTables() []MicrosoftGraphWorkbookTable`

GetTables returns the Tables field if non-nil, zero value otherwise.

### GetTablesOk

`func (o *WorkbookWorksheet) GetTablesOk() (*[]MicrosoftGraphWorkbookTable, bool)`

GetTablesOk returns a tuple with the Tables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTables

`func (o *WorkbookWorksheet) SetTables(v []MicrosoftGraphWorkbookTable)`

SetTables sets Tables field to given value.

### HasTables

`func (o *WorkbookWorksheet) HasTables() bool`

HasTables returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


