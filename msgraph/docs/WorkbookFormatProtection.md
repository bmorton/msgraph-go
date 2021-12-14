# WorkbookFormatProtection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FormulaHidden** | Pointer to **NullableBool** | Indicates if Excel hides the formula for the cells in the range. A null value indicates that the entire range doesn&#39;t have uniform formula hidden setting. | [optional] 
**Locked** | Pointer to **NullableBool** | Indicates if Excel locks the cells in the object. A null value indicates that the entire range doesn&#39;t have uniform lock setting. | [optional] 

## Methods

### NewWorkbookFormatProtection

`func NewWorkbookFormatProtection() *WorkbookFormatProtection`

NewWorkbookFormatProtection instantiates a new WorkbookFormatProtection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkbookFormatProtectionWithDefaults

`func NewWorkbookFormatProtectionWithDefaults() *WorkbookFormatProtection`

NewWorkbookFormatProtectionWithDefaults instantiates a new WorkbookFormatProtection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFormulaHidden

`func (o *WorkbookFormatProtection) GetFormulaHidden() bool`

GetFormulaHidden returns the FormulaHidden field if non-nil, zero value otherwise.

### GetFormulaHiddenOk

`func (o *WorkbookFormatProtection) GetFormulaHiddenOk() (*bool, bool)`

GetFormulaHiddenOk returns a tuple with the FormulaHidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormulaHidden

`func (o *WorkbookFormatProtection) SetFormulaHidden(v bool)`

SetFormulaHidden sets FormulaHidden field to given value.

### HasFormulaHidden

`func (o *WorkbookFormatProtection) HasFormulaHidden() bool`

HasFormulaHidden returns a boolean if a field has been set.

### SetFormulaHiddenNil

`func (o *WorkbookFormatProtection) SetFormulaHiddenNil(b bool)`

 SetFormulaHiddenNil sets the value for FormulaHidden to be an explicit nil

### UnsetFormulaHidden
`func (o *WorkbookFormatProtection) UnsetFormulaHidden()`

UnsetFormulaHidden ensures that no value is present for FormulaHidden, not even an explicit nil
### GetLocked

`func (o *WorkbookFormatProtection) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *WorkbookFormatProtection) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *WorkbookFormatProtection) SetLocked(v bool)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *WorkbookFormatProtection) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### SetLockedNil

`func (o *WorkbookFormatProtection) SetLockedNil(b bool)`

 SetLockedNil sets the value for Locked to be an explicit nil

### UnsetLocked
`func (o *WorkbookFormatProtection) UnsetLocked()`

UnsetLocked ensures that no value is present for Locked, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


