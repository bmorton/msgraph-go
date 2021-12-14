# MicrosoftGraphScopedRoleMembership

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**AdministrativeUnitId** | Pointer to **string** | Unique identifier for the administrative unit that the directory role is scoped to | [optional] 
**RoleId** | Pointer to **string** | Unique identifier for the directory role that the member is in. | [optional] 
**RoleMemberInfo** | Pointer to [**MicrosoftGraphIdentity**](MicrosoftGraphIdentity.md) |  | [optional] 

## Methods

### NewMicrosoftGraphScopedRoleMembership

`func NewMicrosoftGraphScopedRoleMembership() *MicrosoftGraphScopedRoleMembership`

NewMicrosoftGraphScopedRoleMembership instantiates a new MicrosoftGraphScopedRoleMembership object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphScopedRoleMembershipWithDefaults

`func NewMicrosoftGraphScopedRoleMembershipWithDefaults() *MicrosoftGraphScopedRoleMembership`

NewMicrosoftGraphScopedRoleMembershipWithDefaults instantiates a new MicrosoftGraphScopedRoleMembership object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphScopedRoleMembership) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphScopedRoleMembership) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphScopedRoleMembership) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphScopedRoleMembership) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAdministrativeUnitId

`func (o *MicrosoftGraphScopedRoleMembership) GetAdministrativeUnitId() string`

GetAdministrativeUnitId returns the AdministrativeUnitId field if non-nil, zero value otherwise.

### GetAdministrativeUnitIdOk

`func (o *MicrosoftGraphScopedRoleMembership) GetAdministrativeUnitIdOk() (*string, bool)`

GetAdministrativeUnitIdOk returns a tuple with the AdministrativeUnitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdministrativeUnitId

`func (o *MicrosoftGraphScopedRoleMembership) SetAdministrativeUnitId(v string)`

SetAdministrativeUnitId sets AdministrativeUnitId field to given value.

### HasAdministrativeUnitId

`func (o *MicrosoftGraphScopedRoleMembership) HasAdministrativeUnitId() bool`

HasAdministrativeUnitId returns a boolean if a field has been set.

### GetRoleId

`func (o *MicrosoftGraphScopedRoleMembership) GetRoleId() string`

GetRoleId returns the RoleId field if non-nil, zero value otherwise.

### GetRoleIdOk

`func (o *MicrosoftGraphScopedRoleMembership) GetRoleIdOk() (*string, bool)`

GetRoleIdOk returns a tuple with the RoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleId

`func (o *MicrosoftGraphScopedRoleMembership) SetRoleId(v string)`

SetRoleId sets RoleId field to given value.

### HasRoleId

`func (o *MicrosoftGraphScopedRoleMembership) HasRoleId() bool`

HasRoleId returns a boolean if a field has been set.

### GetRoleMemberInfo

`func (o *MicrosoftGraphScopedRoleMembership) GetRoleMemberInfo() MicrosoftGraphIdentity`

GetRoleMemberInfo returns the RoleMemberInfo field if non-nil, zero value otherwise.

### GetRoleMemberInfoOk

`func (o *MicrosoftGraphScopedRoleMembership) GetRoleMemberInfoOk() (*MicrosoftGraphIdentity, bool)`

GetRoleMemberInfoOk returns a tuple with the RoleMemberInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleMemberInfo

`func (o *MicrosoftGraphScopedRoleMembership) SetRoleMemberInfo(v MicrosoftGraphIdentity)`

SetRoleMemberInfo sets RoleMemberInfo field to given value.

### HasRoleMemberInfo

`func (o *MicrosoftGraphScopedRoleMembership) HasRoleMemberInfo() bool`

HasRoleMemberInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


