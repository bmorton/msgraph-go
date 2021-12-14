# ScopedRoleMembership

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdministrativeUnitId** | Pointer to **string** | Unique identifier for the administrative unit that the directory role is scoped to | [optional] 
**RoleId** | Pointer to **string** | Unique identifier for the directory role that the member is in. | [optional] 
**RoleMemberInfo** | Pointer to [**MicrosoftGraphIdentity**](MicrosoftGraphIdentity.md) |  | [optional] 

## Methods

### NewScopedRoleMembership

`func NewScopedRoleMembership() *ScopedRoleMembership`

NewScopedRoleMembership instantiates a new ScopedRoleMembership object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScopedRoleMembershipWithDefaults

`func NewScopedRoleMembershipWithDefaults() *ScopedRoleMembership`

NewScopedRoleMembershipWithDefaults instantiates a new ScopedRoleMembership object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdministrativeUnitId

`func (o *ScopedRoleMembership) GetAdministrativeUnitId() string`

GetAdministrativeUnitId returns the AdministrativeUnitId field if non-nil, zero value otherwise.

### GetAdministrativeUnitIdOk

`func (o *ScopedRoleMembership) GetAdministrativeUnitIdOk() (*string, bool)`

GetAdministrativeUnitIdOk returns a tuple with the AdministrativeUnitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdministrativeUnitId

`func (o *ScopedRoleMembership) SetAdministrativeUnitId(v string)`

SetAdministrativeUnitId sets AdministrativeUnitId field to given value.

### HasAdministrativeUnitId

`func (o *ScopedRoleMembership) HasAdministrativeUnitId() bool`

HasAdministrativeUnitId returns a boolean if a field has been set.

### GetRoleId

`func (o *ScopedRoleMembership) GetRoleId() string`

GetRoleId returns the RoleId field if non-nil, zero value otherwise.

### GetRoleIdOk

`func (o *ScopedRoleMembership) GetRoleIdOk() (*string, bool)`

GetRoleIdOk returns a tuple with the RoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleId

`func (o *ScopedRoleMembership) SetRoleId(v string)`

SetRoleId sets RoleId field to given value.

### HasRoleId

`func (o *ScopedRoleMembership) HasRoleId() bool`

HasRoleId returns a boolean if a field has been set.

### GetRoleMemberInfo

`func (o *ScopedRoleMembership) GetRoleMemberInfo() MicrosoftGraphIdentity`

GetRoleMemberInfo returns the RoleMemberInfo field if non-nil, zero value otherwise.

### GetRoleMemberInfoOk

`func (o *ScopedRoleMembership) GetRoleMemberInfoOk() (*MicrosoftGraphIdentity, bool)`

GetRoleMemberInfoOk returns a tuple with the RoleMemberInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleMemberInfo

`func (o *ScopedRoleMembership) SetRoleMemberInfo(v MicrosoftGraphIdentity)`

SetRoleMemberInfo sets RoleMemberInfo field to given value.

### HasRoleMemberInfo

`func (o *ScopedRoleMembership) HasRoleMemberInfo() bool`

HasRoleMemberInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


