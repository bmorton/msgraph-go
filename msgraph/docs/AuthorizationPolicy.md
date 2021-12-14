# AuthorizationPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedToSignUpEmailBasedSubscriptions** | Pointer to **bool** | Indicates whether users can sign up for email based subscriptions. | [optional] 
**AllowedToUseSSPR** | Pointer to **bool** | Indicates whether the Self-Serve Password Reset feature can be used by users on the tenant. | [optional] 
**AllowEmailVerifiedUsersToJoinOrganization** | Pointer to **bool** | Indicates whether a user can join the tenant by email validation. | [optional] 
**AllowInvitesFrom** | Pointer to [**AnyOfmicrosoftGraphAllowInvitesFrom**](anyOf&lt;microsoft.graph.allowInvitesFrom&gt;.md) | Indicates who can invite external users to the organization. Possible values are: none, adminsAndGuestInviters, adminsGuestInvitersAndAllMembers, everyone.  everyone is the default setting for all cloud environments except US Government. See more in the table below. | [optional] 
**BlockMsolPowerShell** | Pointer to **NullableBool** | To disable the use of MSOL PowerShell set this property to true. This will also disable user-based access to the legacy service endpoint used by MSOL PowerShell. This does not affect Azure AD Connect or Microsoft Graph. | [optional] 
**DefaultUserRolePermissions** | Pointer to [**MicrosoftGraphDefaultUserRolePermissions**](MicrosoftGraphDefaultUserRolePermissions.md) |  | [optional] 
**GuestUserRoleId** | Pointer to **NullableString** | Represents role templateId for the role that should be granted to guest user. Currently following roles are supported:  User (a0b1b346-4d3e-4e8b-98f8-753987be4970), Guest User (10dae51f-b6af-4016-8d66-8c2a99b929b3), and Restricted Guest User (2af84b1e-32c8-42b7-82bc-daa82404023b). | [optional] 

## Methods

### NewAuthorizationPolicy

`func NewAuthorizationPolicy() *AuthorizationPolicy`

NewAuthorizationPolicy instantiates a new AuthorizationPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorizationPolicyWithDefaults

`func NewAuthorizationPolicyWithDefaults() *AuthorizationPolicy`

NewAuthorizationPolicyWithDefaults instantiates a new AuthorizationPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedToSignUpEmailBasedSubscriptions

`func (o *AuthorizationPolicy) GetAllowedToSignUpEmailBasedSubscriptions() bool`

GetAllowedToSignUpEmailBasedSubscriptions returns the AllowedToSignUpEmailBasedSubscriptions field if non-nil, zero value otherwise.

### GetAllowedToSignUpEmailBasedSubscriptionsOk

`func (o *AuthorizationPolicy) GetAllowedToSignUpEmailBasedSubscriptionsOk() (*bool, bool)`

GetAllowedToSignUpEmailBasedSubscriptionsOk returns a tuple with the AllowedToSignUpEmailBasedSubscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedToSignUpEmailBasedSubscriptions

`func (o *AuthorizationPolicy) SetAllowedToSignUpEmailBasedSubscriptions(v bool)`

SetAllowedToSignUpEmailBasedSubscriptions sets AllowedToSignUpEmailBasedSubscriptions field to given value.

### HasAllowedToSignUpEmailBasedSubscriptions

`func (o *AuthorizationPolicy) HasAllowedToSignUpEmailBasedSubscriptions() bool`

HasAllowedToSignUpEmailBasedSubscriptions returns a boolean if a field has been set.

### GetAllowedToUseSSPR

`func (o *AuthorizationPolicy) GetAllowedToUseSSPR() bool`

GetAllowedToUseSSPR returns the AllowedToUseSSPR field if non-nil, zero value otherwise.

### GetAllowedToUseSSPROk

`func (o *AuthorizationPolicy) GetAllowedToUseSSPROk() (*bool, bool)`

GetAllowedToUseSSPROk returns a tuple with the AllowedToUseSSPR field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedToUseSSPR

`func (o *AuthorizationPolicy) SetAllowedToUseSSPR(v bool)`

SetAllowedToUseSSPR sets AllowedToUseSSPR field to given value.

### HasAllowedToUseSSPR

`func (o *AuthorizationPolicy) HasAllowedToUseSSPR() bool`

HasAllowedToUseSSPR returns a boolean if a field has been set.

### GetAllowEmailVerifiedUsersToJoinOrganization

`func (o *AuthorizationPolicy) GetAllowEmailVerifiedUsersToJoinOrganization() bool`

GetAllowEmailVerifiedUsersToJoinOrganization returns the AllowEmailVerifiedUsersToJoinOrganization field if non-nil, zero value otherwise.

### GetAllowEmailVerifiedUsersToJoinOrganizationOk

`func (o *AuthorizationPolicy) GetAllowEmailVerifiedUsersToJoinOrganizationOk() (*bool, bool)`

GetAllowEmailVerifiedUsersToJoinOrganizationOk returns a tuple with the AllowEmailVerifiedUsersToJoinOrganization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowEmailVerifiedUsersToJoinOrganization

`func (o *AuthorizationPolicy) SetAllowEmailVerifiedUsersToJoinOrganization(v bool)`

SetAllowEmailVerifiedUsersToJoinOrganization sets AllowEmailVerifiedUsersToJoinOrganization field to given value.

### HasAllowEmailVerifiedUsersToJoinOrganization

`func (o *AuthorizationPolicy) HasAllowEmailVerifiedUsersToJoinOrganization() bool`

HasAllowEmailVerifiedUsersToJoinOrganization returns a boolean if a field has been set.

### GetAllowInvitesFrom

`func (o *AuthorizationPolicy) GetAllowInvitesFrom() AnyOfmicrosoftGraphAllowInvitesFrom`

GetAllowInvitesFrom returns the AllowInvitesFrom field if non-nil, zero value otherwise.

### GetAllowInvitesFromOk

`func (o *AuthorizationPolicy) GetAllowInvitesFromOk() (*AnyOfmicrosoftGraphAllowInvitesFrom, bool)`

GetAllowInvitesFromOk returns a tuple with the AllowInvitesFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowInvitesFrom

`func (o *AuthorizationPolicy) SetAllowInvitesFrom(v AnyOfmicrosoftGraphAllowInvitesFrom)`

SetAllowInvitesFrom sets AllowInvitesFrom field to given value.

### HasAllowInvitesFrom

`func (o *AuthorizationPolicy) HasAllowInvitesFrom() bool`

HasAllowInvitesFrom returns a boolean if a field has been set.

### SetAllowInvitesFromNil

`func (o *AuthorizationPolicy) SetAllowInvitesFromNil(b bool)`

 SetAllowInvitesFromNil sets the value for AllowInvitesFrom to be an explicit nil

### UnsetAllowInvitesFrom
`func (o *AuthorizationPolicy) UnsetAllowInvitesFrom()`

UnsetAllowInvitesFrom ensures that no value is present for AllowInvitesFrom, not even an explicit nil
### GetBlockMsolPowerShell

`func (o *AuthorizationPolicy) GetBlockMsolPowerShell() bool`

GetBlockMsolPowerShell returns the BlockMsolPowerShell field if non-nil, zero value otherwise.

### GetBlockMsolPowerShellOk

`func (o *AuthorizationPolicy) GetBlockMsolPowerShellOk() (*bool, bool)`

GetBlockMsolPowerShellOk returns a tuple with the BlockMsolPowerShell field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockMsolPowerShell

`func (o *AuthorizationPolicy) SetBlockMsolPowerShell(v bool)`

SetBlockMsolPowerShell sets BlockMsolPowerShell field to given value.

### HasBlockMsolPowerShell

`func (o *AuthorizationPolicy) HasBlockMsolPowerShell() bool`

HasBlockMsolPowerShell returns a boolean if a field has been set.

### SetBlockMsolPowerShellNil

`func (o *AuthorizationPolicy) SetBlockMsolPowerShellNil(b bool)`

 SetBlockMsolPowerShellNil sets the value for BlockMsolPowerShell to be an explicit nil

### UnsetBlockMsolPowerShell
`func (o *AuthorizationPolicy) UnsetBlockMsolPowerShell()`

UnsetBlockMsolPowerShell ensures that no value is present for BlockMsolPowerShell, not even an explicit nil
### GetDefaultUserRolePermissions

`func (o *AuthorizationPolicy) GetDefaultUserRolePermissions() MicrosoftGraphDefaultUserRolePermissions`

GetDefaultUserRolePermissions returns the DefaultUserRolePermissions field if non-nil, zero value otherwise.

### GetDefaultUserRolePermissionsOk

`func (o *AuthorizationPolicy) GetDefaultUserRolePermissionsOk() (*MicrosoftGraphDefaultUserRolePermissions, bool)`

GetDefaultUserRolePermissionsOk returns a tuple with the DefaultUserRolePermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultUserRolePermissions

`func (o *AuthorizationPolicy) SetDefaultUserRolePermissions(v MicrosoftGraphDefaultUserRolePermissions)`

SetDefaultUserRolePermissions sets DefaultUserRolePermissions field to given value.

### HasDefaultUserRolePermissions

`func (o *AuthorizationPolicy) HasDefaultUserRolePermissions() bool`

HasDefaultUserRolePermissions returns a boolean if a field has been set.

### GetGuestUserRoleId

`func (o *AuthorizationPolicy) GetGuestUserRoleId() string`

GetGuestUserRoleId returns the GuestUserRoleId field if non-nil, zero value otherwise.

### GetGuestUserRoleIdOk

`func (o *AuthorizationPolicy) GetGuestUserRoleIdOk() (*string, bool)`

GetGuestUserRoleIdOk returns a tuple with the GuestUserRoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestUserRoleId

`func (o *AuthorizationPolicy) SetGuestUserRoleId(v string)`

SetGuestUserRoleId sets GuestUserRoleId field to given value.

### HasGuestUserRoleId

`func (o *AuthorizationPolicy) HasGuestUserRoleId() bool`

HasGuestUserRoleId returns a boolean if a field has been set.

### SetGuestUserRoleIdNil

`func (o *AuthorizationPolicy) SetGuestUserRoleIdNil(b bool)`

 SetGuestUserRoleIdNil sets the value for GuestUserRoleId to be an explicit nil

### UnsetGuestUserRoleId
`func (o *AuthorizationPolicy) UnsetGuestUserRoleId()`

UnsetGuestUserRoleId ensures that no value is present for GuestUserRoleId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


