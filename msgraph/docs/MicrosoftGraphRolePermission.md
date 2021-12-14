# MicrosoftGraphRolePermission

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceActions** | Pointer to [**[]AnyOfmicrosoftGraphResourceAction**](AnyOfmicrosoftGraphResourceAction.md) | Resource Actions each containing a set of allowed and not allowed permissions. | [optional] 

## Methods

### NewMicrosoftGraphRolePermission

`func NewMicrosoftGraphRolePermission() *MicrosoftGraphRolePermission`

NewMicrosoftGraphRolePermission instantiates a new MicrosoftGraphRolePermission object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphRolePermissionWithDefaults

`func NewMicrosoftGraphRolePermissionWithDefaults() *MicrosoftGraphRolePermission`

NewMicrosoftGraphRolePermissionWithDefaults instantiates a new MicrosoftGraphRolePermission object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceActions

`func (o *MicrosoftGraphRolePermission) GetResourceActions() []*AnyOfmicrosoftGraphResourceAction`

GetResourceActions returns the ResourceActions field if non-nil, zero value otherwise.

### GetResourceActionsOk

`func (o *MicrosoftGraphRolePermission) GetResourceActionsOk() (*[]*AnyOfmicrosoftGraphResourceAction, bool)`

GetResourceActionsOk returns a tuple with the ResourceActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceActions

`func (o *MicrosoftGraphRolePermission) SetResourceActions(v []*AnyOfmicrosoftGraphResourceAction)`

SetResourceActions sets ResourceActions field to given value.

### HasResourceActions

`func (o *MicrosoftGraphRolePermission) HasResourceActions() bool`

HasResourceActions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


