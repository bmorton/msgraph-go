# MicrosoftGraphDefaultUserRolePermissions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedToCreateApps** | Pointer to **bool** | Indicates whether the default user role can create applications. | [optional] 
**AllowedToCreateSecurityGroups** | Pointer to **bool** | Indicates whether the default user role can create security groups. | [optional] 
**AllowedToReadOtherUsers** | Pointer to **bool** | Indicates whether the default user role can read other users. | [optional] 
**PermissionGrantPoliciesAssigned** | Pointer to **[]string** | Indicates if user consent to apps is allowed, and if it is, which permission to grant consent and which app consent policy (permissionGrantPolicy) govern the permission for users to grant consent. Value should be in the format managePermissionGrantsForSelf.{id}, where {id} is the id of a built-in or custom app consent policy. An empty list indicates user consent to apps is disabled. | [optional] 

## Methods

### NewMicrosoftGraphDefaultUserRolePermissions

`func NewMicrosoftGraphDefaultUserRolePermissions() *MicrosoftGraphDefaultUserRolePermissions`

NewMicrosoftGraphDefaultUserRolePermissions instantiates a new MicrosoftGraphDefaultUserRolePermissions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphDefaultUserRolePermissionsWithDefaults

`func NewMicrosoftGraphDefaultUserRolePermissionsWithDefaults() *MicrosoftGraphDefaultUserRolePermissions`

NewMicrosoftGraphDefaultUserRolePermissionsWithDefaults instantiates a new MicrosoftGraphDefaultUserRolePermissions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedToCreateApps

`func (o *MicrosoftGraphDefaultUserRolePermissions) GetAllowedToCreateApps() bool`

GetAllowedToCreateApps returns the AllowedToCreateApps field if non-nil, zero value otherwise.

### GetAllowedToCreateAppsOk

`func (o *MicrosoftGraphDefaultUserRolePermissions) GetAllowedToCreateAppsOk() (*bool, bool)`

GetAllowedToCreateAppsOk returns a tuple with the AllowedToCreateApps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedToCreateApps

`func (o *MicrosoftGraphDefaultUserRolePermissions) SetAllowedToCreateApps(v bool)`

SetAllowedToCreateApps sets AllowedToCreateApps field to given value.

### HasAllowedToCreateApps

`func (o *MicrosoftGraphDefaultUserRolePermissions) HasAllowedToCreateApps() bool`

HasAllowedToCreateApps returns a boolean if a field has been set.

### GetAllowedToCreateSecurityGroups

`func (o *MicrosoftGraphDefaultUserRolePermissions) GetAllowedToCreateSecurityGroups() bool`

GetAllowedToCreateSecurityGroups returns the AllowedToCreateSecurityGroups field if non-nil, zero value otherwise.

### GetAllowedToCreateSecurityGroupsOk

`func (o *MicrosoftGraphDefaultUserRolePermissions) GetAllowedToCreateSecurityGroupsOk() (*bool, bool)`

GetAllowedToCreateSecurityGroupsOk returns a tuple with the AllowedToCreateSecurityGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedToCreateSecurityGroups

`func (o *MicrosoftGraphDefaultUserRolePermissions) SetAllowedToCreateSecurityGroups(v bool)`

SetAllowedToCreateSecurityGroups sets AllowedToCreateSecurityGroups field to given value.

### HasAllowedToCreateSecurityGroups

`func (o *MicrosoftGraphDefaultUserRolePermissions) HasAllowedToCreateSecurityGroups() bool`

HasAllowedToCreateSecurityGroups returns a boolean if a field has been set.

### GetAllowedToReadOtherUsers

`func (o *MicrosoftGraphDefaultUserRolePermissions) GetAllowedToReadOtherUsers() bool`

GetAllowedToReadOtherUsers returns the AllowedToReadOtherUsers field if non-nil, zero value otherwise.

### GetAllowedToReadOtherUsersOk

`func (o *MicrosoftGraphDefaultUserRolePermissions) GetAllowedToReadOtherUsersOk() (*bool, bool)`

GetAllowedToReadOtherUsersOk returns a tuple with the AllowedToReadOtherUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedToReadOtherUsers

`func (o *MicrosoftGraphDefaultUserRolePermissions) SetAllowedToReadOtherUsers(v bool)`

SetAllowedToReadOtherUsers sets AllowedToReadOtherUsers field to given value.

### HasAllowedToReadOtherUsers

`func (o *MicrosoftGraphDefaultUserRolePermissions) HasAllowedToReadOtherUsers() bool`

HasAllowedToReadOtherUsers returns a boolean if a field has been set.

### GetPermissionGrantPoliciesAssigned

`func (o *MicrosoftGraphDefaultUserRolePermissions) GetPermissionGrantPoliciesAssigned() []*string`

GetPermissionGrantPoliciesAssigned returns the PermissionGrantPoliciesAssigned field if non-nil, zero value otherwise.

### GetPermissionGrantPoliciesAssignedOk

`func (o *MicrosoftGraphDefaultUserRolePermissions) GetPermissionGrantPoliciesAssignedOk() (*[]*string, bool)`

GetPermissionGrantPoliciesAssignedOk returns a tuple with the PermissionGrantPoliciesAssigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionGrantPoliciesAssigned

`func (o *MicrosoftGraphDefaultUserRolePermissions) SetPermissionGrantPoliciesAssigned(v []*string)`

SetPermissionGrantPoliciesAssigned sets PermissionGrantPoliciesAssigned field to given value.

### HasPermissionGrantPoliciesAssigned

`func (o *MicrosoftGraphDefaultUserRolePermissions) HasPermissionGrantPoliciesAssigned() bool`

HasPermissionGrantPoliciesAssigned returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


