# MicrosoftGraphAuthorizationPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**DeletedDateTime** | Pointer to **NullableTime** |  | [optional] 
**Description** | Pointer to **NullableString** | Description for this policy. Required. | [optional] 
**DisplayName** | Pointer to **NullableString** | Display name for this policy. Required. | [optional] 
**AllowedToSignUpEmailBasedSubscriptions** | Pointer to **bool** | Indicates whether users can sign up for email based subscriptions. | [optional] 
**AllowedToUseSSPR** | Pointer to **bool** | Indicates whether the Self-Serve Password Reset feature can be used by users on the tenant. | [optional] 
**AllowEmailVerifiedUsersToJoinOrganization** | Pointer to **bool** | Indicates whether a user can join the tenant by email validation. | [optional] 
**AllowInvitesFrom** | Pointer to [**AnyOfmicrosoftGraphAllowInvitesFrom**](anyOf&lt;microsoft.graph.allowInvitesFrom&gt;.md) | Indicates who can invite external users to the organization. Possible values are: none, adminsAndGuestInviters, adminsGuestInvitersAndAllMembers, everyone.  everyone is the default setting for all cloud environments except US Government. See more in the table below. | [optional] 
**BlockMsolPowerShell** | Pointer to **NullableBool** | To disable the use of MSOL PowerShell set this property to true. This will also disable user-based access to the legacy service endpoint used by MSOL PowerShell. This does not affect Azure AD Connect or Microsoft Graph. | [optional] 
**DefaultUserRolePermissions** | Pointer to [**MicrosoftGraphDefaultUserRolePermissions**](MicrosoftGraphDefaultUserRolePermissions.md) |  | [optional] 
**GuestUserRoleId** | Pointer to **NullableString** | Represents role templateId for the role that should be granted to guest user. Currently following roles are supported:  User (a0b1b346-4d3e-4e8b-98f8-753987be4970), Guest User (10dae51f-b6af-4016-8d66-8c2a99b929b3), and Restricted Guest User (2af84b1e-32c8-42b7-82bc-daa82404023b). | [optional] 

## Methods

### NewMicrosoftGraphAuthorizationPolicy

`func NewMicrosoftGraphAuthorizationPolicy() *MicrosoftGraphAuthorizationPolicy`

NewMicrosoftGraphAuthorizationPolicy instantiates a new MicrosoftGraphAuthorizationPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphAuthorizationPolicyWithDefaults

`func NewMicrosoftGraphAuthorizationPolicyWithDefaults() *MicrosoftGraphAuthorizationPolicy`

NewMicrosoftGraphAuthorizationPolicyWithDefaults instantiates a new MicrosoftGraphAuthorizationPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphAuthorizationPolicy) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphAuthorizationPolicy) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphAuthorizationPolicy) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphAuthorizationPolicy) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDeletedDateTime

`func (o *MicrosoftGraphAuthorizationPolicy) GetDeletedDateTime() time.Time`

GetDeletedDateTime returns the DeletedDateTime field if non-nil, zero value otherwise.

### GetDeletedDateTimeOk

`func (o *MicrosoftGraphAuthorizationPolicy) GetDeletedDateTimeOk() (*time.Time, bool)`

GetDeletedDateTimeOk returns a tuple with the DeletedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedDateTime

`func (o *MicrosoftGraphAuthorizationPolicy) SetDeletedDateTime(v time.Time)`

SetDeletedDateTime sets DeletedDateTime field to given value.

### HasDeletedDateTime

`func (o *MicrosoftGraphAuthorizationPolicy) HasDeletedDateTime() bool`

HasDeletedDateTime returns a boolean if a field has been set.

### SetDeletedDateTimeNil

`func (o *MicrosoftGraphAuthorizationPolicy) SetDeletedDateTimeNil(b bool)`

 SetDeletedDateTimeNil sets the value for DeletedDateTime to be an explicit nil

### UnsetDeletedDateTime
`func (o *MicrosoftGraphAuthorizationPolicy) UnsetDeletedDateTime()`

UnsetDeletedDateTime ensures that no value is present for DeletedDateTime, not even an explicit nil
### GetDescription

`func (o *MicrosoftGraphAuthorizationPolicy) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MicrosoftGraphAuthorizationPolicy) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MicrosoftGraphAuthorizationPolicy) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MicrosoftGraphAuthorizationPolicy) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *MicrosoftGraphAuthorizationPolicy) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *MicrosoftGraphAuthorizationPolicy) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDisplayName

`func (o *MicrosoftGraphAuthorizationPolicy) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphAuthorizationPolicy) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *MicrosoftGraphAuthorizationPolicy) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *MicrosoftGraphAuthorizationPolicy) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *MicrosoftGraphAuthorizationPolicy) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *MicrosoftGraphAuthorizationPolicy) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetAllowedToSignUpEmailBasedSubscriptions

`func (o *MicrosoftGraphAuthorizationPolicy) GetAllowedToSignUpEmailBasedSubscriptions() bool`

GetAllowedToSignUpEmailBasedSubscriptions returns the AllowedToSignUpEmailBasedSubscriptions field if non-nil, zero value otherwise.

### GetAllowedToSignUpEmailBasedSubscriptionsOk

`func (o *MicrosoftGraphAuthorizationPolicy) GetAllowedToSignUpEmailBasedSubscriptionsOk() (*bool, bool)`

GetAllowedToSignUpEmailBasedSubscriptionsOk returns a tuple with the AllowedToSignUpEmailBasedSubscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedToSignUpEmailBasedSubscriptions

`func (o *MicrosoftGraphAuthorizationPolicy) SetAllowedToSignUpEmailBasedSubscriptions(v bool)`

SetAllowedToSignUpEmailBasedSubscriptions sets AllowedToSignUpEmailBasedSubscriptions field to given value.

### HasAllowedToSignUpEmailBasedSubscriptions

`func (o *MicrosoftGraphAuthorizationPolicy) HasAllowedToSignUpEmailBasedSubscriptions() bool`

HasAllowedToSignUpEmailBasedSubscriptions returns a boolean if a field has been set.

### GetAllowedToUseSSPR

`func (o *MicrosoftGraphAuthorizationPolicy) GetAllowedToUseSSPR() bool`

GetAllowedToUseSSPR returns the AllowedToUseSSPR field if non-nil, zero value otherwise.

### GetAllowedToUseSSPROk

`func (o *MicrosoftGraphAuthorizationPolicy) GetAllowedToUseSSPROk() (*bool, bool)`

GetAllowedToUseSSPROk returns a tuple with the AllowedToUseSSPR field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedToUseSSPR

`func (o *MicrosoftGraphAuthorizationPolicy) SetAllowedToUseSSPR(v bool)`

SetAllowedToUseSSPR sets AllowedToUseSSPR field to given value.

### HasAllowedToUseSSPR

`func (o *MicrosoftGraphAuthorizationPolicy) HasAllowedToUseSSPR() bool`

HasAllowedToUseSSPR returns a boolean if a field has been set.

### GetAllowEmailVerifiedUsersToJoinOrganization

`func (o *MicrosoftGraphAuthorizationPolicy) GetAllowEmailVerifiedUsersToJoinOrganization() bool`

GetAllowEmailVerifiedUsersToJoinOrganization returns the AllowEmailVerifiedUsersToJoinOrganization field if non-nil, zero value otherwise.

### GetAllowEmailVerifiedUsersToJoinOrganizationOk

`func (o *MicrosoftGraphAuthorizationPolicy) GetAllowEmailVerifiedUsersToJoinOrganizationOk() (*bool, bool)`

GetAllowEmailVerifiedUsersToJoinOrganizationOk returns a tuple with the AllowEmailVerifiedUsersToJoinOrganization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowEmailVerifiedUsersToJoinOrganization

`func (o *MicrosoftGraphAuthorizationPolicy) SetAllowEmailVerifiedUsersToJoinOrganization(v bool)`

SetAllowEmailVerifiedUsersToJoinOrganization sets AllowEmailVerifiedUsersToJoinOrganization field to given value.

### HasAllowEmailVerifiedUsersToJoinOrganization

`func (o *MicrosoftGraphAuthorizationPolicy) HasAllowEmailVerifiedUsersToJoinOrganization() bool`

HasAllowEmailVerifiedUsersToJoinOrganization returns a boolean if a field has been set.

### GetAllowInvitesFrom

`func (o *MicrosoftGraphAuthorizationPolicy) GetAllowInvitesFrom() AnyOfmicrosoftGraphAllowInvitesFrom`

GetAllowInvitesFrom returns the AllowInvitesFrom field if non-nil, zero value otherwise.

### GetAllowInvitesFromOk

`func (o *MicrosoftGraphAuthorizationPolicy) GetAllowInvitesFromOk() (*AnyOfmicrosoftGraphAllowInvitesFrom, bool)`

GetAllowInvitesFromOk returns a tuple with the AllowInvitesFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowInvitesFrom

`func (o *MicrosoftGraphAuthorizationPolicy) SetAllowInvitesFrom(v AnyOfmicrosoftGraphAllowInvitesFrom)`

SetAllowInvitesFrom sets AllowInvitesFrom field to given value.

### HasAllowInvitesFrom

`func (o *MicrosoftGraphAuthorizationPolicy) HasAllowInvitesFrom() bool`

HasAllowInvitesFrom returns a boolean if a field has been set.

### SetAllowInvitesFromNil

`func (o *MicrosoftGraphAuthorizationPolicy) SetAllowInvitesFromNil(b bool)`

 SetAllowInvitesFromNil sets the value for AllowInvitesFrom to be an explicit nil

### UnsetAllowInvitesFrom
`func (o *MicrosoftGraphAuthorizationPolicy) UnsetAllowInvitesFrom()`

UnsetAllowInvitesFrom ensures that no value is present for AllowInvitesFrom, not even an explicit nil
### GetBlockMsolPowerShell

`func (o *MicrosoftGraphAuthorizationPolicy) GetBlockMsolPowerShell() bool`

GetBlockMsolPowerShell returns the BlockMsolPowerShell field if non-nil, zero value otherwise.

### GetBlockMsolPowerShellOk

`func (o *MicrosoftGraphAuthorizationPolicy) GetBlockMsolPowerShellOk() (*bool, bool)`

GetBlockMsolPowerShellOk returns a tuple with the BlockMsolPowerShell field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockMsolPowerShell

`func (o *MicrosoftGraphAuthorizationPolicy) SetBlockMsolPowerShell(v bool)`

SetBlockMsolPowerShell sets BlockMsolPowerShell field to given value.

### HasBlockMsolPowerShell

`func (o *MicrosoftGraphAuthorizationPolicy) HasBlockMsolPowerShell() bool`

HasBlockMsolPowerShell returns a boolean if a field has been set.

### SetBlockMsolPowerShellNil

`func (o *MicrosoftGraphAuthorizationPolicy) SetBlockMsolPowerShellNil(b bool)`

 SetBlockMsolPowerShellNil sets the value for BlockMsolPowerShell to be an explicit nil

### UnsetBlockMsolPowerShell
`func (o *MicrosoftGraphAuthorizationPolicy) UnsetBlockMsolPowerShell()`

UnsetBlockMsolPowerShell ensures that no value is present for BlockMsolPowerShell, not even an explicit nil
### GetDefaultUserRolePermissions

`func (o *MicrosoftGraphAuthorizationPolicy) GetDefaultUserRolePermissions() MicrosoftGraphDefaultUserRolePermissions`

GetDefaultUserRolePermissions returns the DefaultUserRolePermissions field if non-nil, zero value otherwise.

### GetDefaultUserRolePermissionsOk

`func (o *MicrosoftGraphAuthorizationPolicy) GetDefaultUserRolePermissionsOk() (*MicrosoftGraphDefaultUserRolePermissions, bool)`

GetDefaultUserRolePermissionsOk returns a tuple with the DefaultUserRolePermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultUserRolePermissions

`func (o *MicrosoftGraphAuthorizationPolicy) SetDefaultUserRolePermissions(v MicrosoftGraphDefaultUserRolePermissions)`

SetDefaultUserRolePermissions sets DefaultUserRolePermissions field to given value.

### HasDefaultUserRolePermissions

`func (o *MicrosoftGraphAuthorizationPolicy) HasDefaultUserRolePermissions() bool`

HasDefaultUserRolePermissions returns a boolean if a field has been set.

### GetGuestUserRoleId

`func (o *MicrosoftGraphAuthorizationPolicy) GetGuestUserRoleId() string`

GetGuestUserRoleId returns the GuestUserRoleId field if non-nil, zero value otherwise.

### GetGuestUserRoleIdOk

`func (o *MicrosoftGraphAuthorizationPolicy) GetGuestUserRoleIdOk() (*string, bool)`

GetGuestUserRoleIdOk returns a tuple with the GuestUserRoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestUserRoleId

`func (o *MicrosoftGraphAuthorizationPolicy) SetGuestUserRoleId(v string)`

SetGuestUserRoleId sets GuestUserRoleId field to given value.

### HasGuestUserRoleId

`func (o *MicrosoftGraphAuthorizationPolicy) HasGuestUserRoleId() bool`

HasGuestUserRoleId returns a boolean if a field has been set.

### SetGuestUserRoleIdNil

`func (o *MicrosoftGraphAuthorizationPolicy) SetGuestUserRoleIdNil(b bool)`

 SetGuestUserRoleIdNil sets the value for GuestUserRoleId to be an explicit nil

### UnsetGuestUserRoleId
`func (o *MicrosoftGraphAuthorizationPolicy) UnsetGuestUserRoleId()`

UnsetGuestUserRoleId ensures that no value is present for GuestUserRoleId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


