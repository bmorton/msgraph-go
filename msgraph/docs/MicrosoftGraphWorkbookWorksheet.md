# MicrosoftGraphWorkbookWorksheet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**Name** | Pointer to **NullableString** | The display name of the worksheet. | [optional] 
**Position** | Pointer to **int32** | The zero-based position of the worksheet within the workbook. | [optional] 
**Visibility** | Pointer to **string** | The Visibility of the worksheet. The possible values are: Visible, Hidden, VeryHidden. | [optional] 
**Charts** | Pointer to [**[]MicrosoftGraphWorkbookChart**](MicrosoftGraphWorkbookChart.md) | Returns collection of charts that are part of the worksheet. Read-only. | [optional] 
**Names** | Pointer to [**[]MicrosoftGraphWorkbookNamedItem**](MicrosoftGraphWorkbookNamedItem.md) | Returns collection of names that are associated with the worksheet. Read-only. | [optional] 
**PivotTables** | Pointer to [**[]MicrosoftGraphWorkbookPivotTable**](MicrosoftGraphWorkbookPivotTable.md) | Collection of PivotTables that are part of the worksheet. | [optional] 
**Protection** | Pointer to [**AnyOfmicrosoftGraphWorkbookWorksheetProtection**](anyOf&lt;microsoft.graph.workbookWorksheetProtection&gt;.md) | Returns sheet protection object for a worksheet. Read-only. | [optional] 
**Tables** | Pointer to [**[]MicrosoftGraphWorkbookTable**](MicrosoftGraphWorkbookTable.md) | Collection of tables that are part of the worksheet. Read-only. | [optional] 

## Methods

### NewMicrosoftGraphWorkbookWorksheet

`func NewMicrosoftGraphWorkbookWorksheet() *MicrosoftGraphWorkbookWorksheet`

NewMicrosoftGraphWorkbookWorksheet instantiates a new MicrosoftGraphWorkbookWorksheet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphWorkbookWorksheetWithDefaults

`func NewMicrosoftGraphWorkbookWorksheetWithDefaults() *MicrosoftGraphWorkbookWorksheet`

NewMicrosoftGraphWorkbookWorksheetWithDefaults instantiates a new MicrosoftGraphWorkbookWorksheet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphWorkbookWorksheet) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphWorkbookWorksheet) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphWorkbookWorksheet) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphWorkbookWorksheet) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *MicrosoftGraphWorkbookWorksheet) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MicrosoftGraphWorkbookWorksheet) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MicrosoftGraphWorkbookWorksheet) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MicrosoftGraphWorkbookWorksheet) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *MicrosoftGraphWorkbookWorksheet) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *MicrosoftGraphWorkbookWorksheet) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetPosition

`func (o *MicrosoftGraphWorkbookWorksheet) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *MicrosoftGraphWorkbookWorksheet) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *MicrosoftGraphWorkbookWorksheet) SetPosition(v int32)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *MicrosoftGraphWorkbookWorksheet) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetVisibility

`func (o *MicrosoftGraphWorkbookWorksheet) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *MicrosoftGraphWorkbookWorksheet) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *MicrosoftGraphWorkbookWorksheet) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *MicrosoftGraphWorkbookWorksheet) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetCharts

`func (o *MicrosoftGraphWorkbookWorksheet) GetCharts() []MicrosoftGraphWorkbookChart`

GetCharts returns the Charts field if non-nil, zero value otherwise.

### GetChartsOk

`func (o *MicrosoftGraphWorkbookWorksheet) GetChartsOk() (*[]MicrosoftGraphWorkbookChart, bool)`

GetChartsOk returns a tuple with the Charts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharts

`func (o *MicrosoftGraphWorkbookWorksheet) SetCharts(v []MicrosoftGraphWorkbookChart)`

SetCharts sets Charts field to given value.

### HasCharts

`func (o *MicrosoftGraphWorkbookWorksheet) HasCharts() bool`

HasCharts returns a boolean if a field has been set.

### GetNames

`func (o *MicrosoftGraphWorkbookWorksheet) GetNames() []MicrosoftGraphWorkbookNamedItem`

GetNames returns the Names field if non-nil, zero value otherwise.

### GetNamesOk

`func (o *MicrosoftGraphWorkbookWorksheet) GetNamesOk() (*[]MicrosoftGraphWorkbookNamedItem, bool)`

GetNamesOk returns a tuple with the Names field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNames

`func (o *MicrosoftGraphWorkbookWorksheet) SetNames(v []MicrosoftGraphWorkbookNamedItem)`

SetNames sets Names field to given value.

### HasNames

`func (o *MicrosoftGraphWorkbookWorksheet) HasNames() bool`

HasNames returns a boolean if a field has been set.

### GetPivotTables

`func (o *MicrosoftGraphWorkbookWorksheet) GetPivotTables() []MicrosoftGraphWorkbookPivotTable`

GetPivotTables returns the PivotTables field if non-nil, zero value otherwise.

### GetPivotTablesOk

`func (o *MicrosoftGraphWorkbookWorksheet) GetPivotTablesOk() (*[]MicrosoftGraphWorkbookPivotTable, bool)`

GetPivotTablesOk returns a tuple with the PivotTables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPivotTables

`func (o *MicrosoftGraphWorkbookWorksheet) SetPivotTables(v []MicrosoftGraphWorkbookPivotTable)`

SetPivotTables sets PivotTables field to given value.

### HasPivotTables

`func (o *MicrosoftGraphWorkbookWorksheet) HasPivotTables() bool`

HasPivotTables returns a boolean if a field has been set.

### GetProtection

`func (o *MicrosoftGraphWorkbookWorksheet) GetProtection() AnyOfmicrosoftGraphWorkbookWorksheetProtection`

GetProtection returns the Protection field if non-nil, zero value otherwise.

### GetProtectionOk

`func (o *MicrosoftGraphWorkbookWorksheet) GetProtectionOk() (*AnyOfmicrosoftGraphWorkbookWorksheetProtection, bool)`

GetProtectionOk returns a tuple with the Protection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtection

`func (o *MicrosoftGraphWorkbookWorksheet) SetProtection(v AnyOfmicrosoftGraphWorkbookWorksheetProtection)`

SetProtection sets Protection field to given value.

### HasProtection

`func (o *MicrosoftGraphWorkbookWorksheet) HasProtection() bool`

HasProtection returns a boolean if a field has been set.

### SetProtectionNil

`func (o *MicrosoftGraphWorkbookWorksheet) SetProtectionNil(b bool)`

 SetProtectionNil sets the value for Protection to be an explicit nil

### UnsetProtection
`func (o *MicrosoftGraphWorkbookWorksheet) UnsetProtection()`

UnsetProtection ensures that no value is present for Protection, not even an explicit nil
### GetTables

`func (o *MicrosoftGraphWorkbookWorksheet) GetTables() []MicrosoftGraphWorkbookTable`

GetTables returns the Tables field if non-nil, zero value otherwise.

### GetTablesOk

`func (o *MicrosoftGraphWorkbookWorksheet) GetTablesOk() (*[]MicrosoftGraphWorkbookTable, bool)`

GetTablesOk returns a tuple with the Tables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTables

`func (o *MicrosoftGraphWorkbookWorksheet) SetTables(v []MicrosoftGraphWorkbookTable)`

SetTables sets Tables field to given value.

### HasTables

`func (o *MicrosoftGraphWorkbookWorksheet) HasTables() bool`

HasTables returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


