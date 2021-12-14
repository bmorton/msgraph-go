# MicrosoftGraphUnifiedRolePermission

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedResourceActions** | Pointer to **[]string** | Set of tasks that can be performed on a resource. Required. | [optional] 
**Condition** | Pointer to **NullableString** | Optional constraints that must be met for the permission to be effective. | [optional] 
**ExcludedResourceActions** | Pointer to **[]string** | Set of tasks that may not be performed on a resource. Not yet supported. | [optional] 

## Methods

### NewMicrosoftGraphUnifiedRolePermission

`func NewMicrosoftGraphUnifiedRolePermission() *MicrosoftGraphUnifiedRolePermission`

NewMicrosoftGraphUnifiedRolePermission instantiates a new MicrosoftGraphUnifiedRolePermission object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphUnifiedRolePermissionWithDefaults

`func NewMicrosoftGraphUnifiedRolePermissionWithDefaults() *MicrosoftGraphUnifiedRolePermission`

NewMicrosoftGraphUnifiedRolePermissionWithDefaults instantiates a new MicrosoftGraphUnifiedRolePermission object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedResourceActions

`func (o *MicrosoftGraphUnifiedRolePermission) GetAllowedResourceActions() []string`

GetAllowedResourceActions returns the AllowedResourceActions field if non-nil, zero value otherwise.

### GetAllowedResourceActionsOk

`func (o *MicrosoftGraphUnifiedRolePermission) GetAllowedResourceActionsOk() (*[]string, bool)`

GetAllowedResourceActionsOk returns a tuple with the AllowedResourceActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedResourceActions

`func (o *MicrosoftGraphUnifiedRolePermission) SetAllowedResourceActions(v []string)`

SetAllowedResourceActions sets AllowedResourceActions field to given value.

### HasAllowedResourceActions

`func (o *MicrosoftGraphUnifiedRolePermission) HasAllowedResourceActions() bool`

HasAllowedResourceActions returns a boolean if a field has been set.

### GetCondition

`func (o *MicrosoftGraphUnifiedRolePermission) GetCondition() string`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *MicrosoftGraphUnifiedRolePermission) GetConditionOk() (*string, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *MicrosoftGraphUnifiedRolePermission) SetCondition(v string)`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *MicrosoftGraphUnifiedRolePermission) HasCondition() bool`

HasCondition returns a boolean if a field has been set.

### SetConditionNil

`func (o *MicrosoftGraphUnifiedRolePermission) SetConditionNil(b bool)`

 SetConditionNil sets the value for Condition to be an explicit nil

### UnsetCondition
`func (o *MicrosoftGraphUnifiedRolePermission) UnsetCondition()`

UnsetCondition ensures that no value is present for Condition, not even an explicit nil
### GetExcludedResourceActions

`func (o *MicrosoftGraphUnifiedRolePermission) GetExcludedResourceActions() []*string`

GetExcludedResourceActions returns the ExcludedResourceActions field if non-nil, zero value otherwise.

### GetExcludedResourceActionsOk

`func (o *MicrosoftGraphUnifiedRolePermission) GetExcludedResourceActionsOk() (*[]*string, bool)`

GetExcludedResourceActionsOk returns a tuple with the ExcludedResourceActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedResourceActions

`func (o *MicrosoftGraphUnifiedRolePermission) SetExcludedResourceActions(v []*string)`

SetExcludedResourceActions sets ExcludedResourceActions field to given value.

### HasExcludedResourceActions

`func (o *MicrosoftGraphUnifiedRolePermission) HasExcludedResourceActions() bool`

HasExcludedResourceActions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


