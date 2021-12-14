# MicrosoftGraphWorkbookFormatProtection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**FormulaHidden** | Pointer to **NullableBool** | Indicates if Excel hides the formula for the cells in the range. A null value indicates that the entire range doesn&#39;t have uniform formula hidden setting. | [optional] 
**Locked** | Pointer to **NullableBool** | Indicates if Excel locks the cells in the object. A null value indicates that the entire range doesn&#39;t have uniform lock setting. | [optional] 

## Methods

### NewMicrosoftGraphWorkbookFormatProtection

`func NewMicrosoftGraphWorkbookFormatProtection() *MicrosoftGraphWorkbookFormatProtection`

NewMicrosoftGraphWorkbookFormatProtection instantiates a new MicrosoftGraphWorkbookFormatProtection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphWorkbookFormatProtectionWithDefaults

`func NewMicrosoftGraphWorkbookFormatProtectionWithDefaults() *MicrosoftGraphWorkbookFormatProtection`

NewMicrosoftGraphWorkbookFormatProtectionWithDefaults instantiates a new MicrosoftGraphWorkbookFormatProtection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphWorkbookFormatProtection) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphWorkbookFormatProtection) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphWorkbookFormatProtection) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphWorkbookFormatProtection) HasId() bool`

HasId returns a boolean if a field has been set.

### GetFormulaHidden

`func (o *MicrosoftGraphWorkbookFormatProtection) GetFormulaHidden() bool`

GetFormulaHidden returns the FormulaHidden field if non-nil, zero value otherwise.

### GetFormulaHiddenOk

`func (o *MicrosoftGraphWorkbookFormatProtection) GetFormulaHiddenOk() (*bool, bool)`

GetFormulaHiddenOk returns a tuple with the FormulaHidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormulaHidden

`func (o *MicrosoftGraphWorkbookFormatProtection) SetFormulaHidden(v bool)`

SetFormulaHidden sets FormulaHidden field to given value.

### HasFormulaHidden

`func (o *MicrosoftGraphWorkbookFormatProtection) HasFormulaHidden() bool`

HasFormulaHidden returns a boolean if a field has been set.

### SetFormulaHiddenNil

`func (o *MicrosoftGraphWorkbookFormatProtection) SetFormulaHiddenNil(b bool)`

 SetFormulaHiddenNil sets the value for FormulaHidden to be an explicit nil

### UnsetFormulaHidden
`func (o *MicrosoftGraphWorkbookFormatProtection) UnsetFormulaHidden()`

UnsetFormulaHidden ensures that no value is present for FormulaHidden, not even an explicit nil
### GetLocked

`func (o *MicrosoftGraphWorkbookFormatProtection) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *MicrosoftGraphWorkbookFormatProtection) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *MicrosoftGraphWorkbookFormatProtection) SetLocked(v bool)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *MicrosoftGraphWorkbookFormatProtection) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### SetLockedNil

`func (o *MicrosoftGraphWorkbookFormatProtection) SetLockedNil(b bool)`

 SetLockedNil sets the value for Locked to be an explicit nil

### UnsetLocked
`func (o *MicrosoftGraphWorkbookFormatProtection) UnsetLocked()`

UnsetLocked ensures that no value is present for Locked, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


