# MicrosoftGraphCalendarPermission

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**AllowedRoles** | Pointer to [**[]AnyOfmicrosoftGraphCalendarRoleType**](AnyOfmicrosoftGraphCalendarRoleType.md) | List of allowed sharing or delegating permission levels for the calendar. Possible values are: none, freeBusyRead, limitedRead, read, write, delegateWithoutPrivateEventAccess, delegateWithPrivateEventAccess, custom. | [optional] 
**EmailAddress** | Pointer to [**AnyOfmicrosoftGraphEmailAddress**](anyOf&lt;microsoft.graph.emailAddress&gt;.md) | Represents a sharee or delegate who has access to the calendar. For the &#39;My Organization&#39; sharee, the address property is null. Read-only. | [optional] 
**IsInsideOrganization** | Pointer to **NullableBool** | True if the user in context (sharee or delegate) is inside the same organization as the calendar owner. | [optional] 
**IsRemovable** | Pointer to **NullableBool** | True if the user can be removed from the list of sharees or delegates for the specified calendar, false otherwise. The &#39;My organization&#39; user determines the permissions other people within your organization have to the given calendar. You cannot remove &#39;My organization&#39; as a sharee to a calendar. | [optional] 
**Role** | Pointer to [**AnyOfmicrosoftGraphCalendarRoleType**](anyOf&lt;microsoft.graph.calendarRoleType&gt;.md) | Current permission level of the calendar sharee or delegate. | [optional] 

## Methods

### NewMicrosoftGraphCalendarPermission

`func NewMicrosoftGraphCalendarPermission() *MicrosoftGraphCalendarPermission`

NewMicrosoftGraphCalendarPermission instantiates a new MicrosoftGraphCalendarPermission object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphCalendarPermissionWithDefaults

`func NewMicrosoftGraphCalendarPermissionWithDefaults() *MicrosoftGraphCalendarPermission`

NewMicrosoftGraphCalendarPermissionWithDefaults instantiates a new MicrosoftGraphCalendarPermission object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphCalendarPermission) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphCalendarPermission) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphCalendarPermission) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphCalendarPermission) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAllowedRoles

`func (o *MicrosoftGraphCalendarPermission) GetAllowedRoles() []*AnyOfmicrosoftGraphCalendarRoleType`

GetAllowedRoles returns the AllowedRoles field if non-nil, zero value otherwise.

### GetAllowedRolesOk

`func (o *MicrosoftGraphCalendarPermission) GetAllowedRolesOk() (*[]*AnyOfmicrosoftGraphCalendarRoleType, bool)`

GetAllowedRolesOk returns a tuple with the AllowedRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedRoles

`func (o *MicrosoftGraphCalendarPermission) SetAllowedRoles(v []*AnyOfmicrosoftGraphCalendarRoleType)`

SetAllowedRoles sets AllowedRoles field to given value.

### HasAllowedRoles

`func (o *MicrosoftGraphCalendarPermission) HasAllowedRoles() bool`

HasAllowedRoles returns a boolean if a field has been set.

### GetEmailAddress

`func (o *MicrosoftGraphCalendarPermission) GetEmailAddress() AnyOfmicrosoftGraphEmailAddress`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *MicrosoftGraphCalendarPermission) GetEmailAddressOk() (*AnyOfmicrosoftGraphEmailAddress, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *MicrosoftGraphCalendarPermission) SetEmailAddress(v AnyOfmicrosoftGraphEmailAddress)`

SetEmailAddress sets EmailAddress field to given value.

### HasEmailAddress

`func (o *MicrosoftGraphCalendarPermission) HasEmailAddress() bool`

HasEmailAddress returns a boolean if a field has been set.

### SetEmailAddressNil

`func (o *MicrosoftGraphCalendarPermission) SetEmailAddressNil(b bool)`

 SetEmailAddressNil sets the value for EmailAddress to be an explicit nil

### UnsetEmailAddress
`func (o *MicrosoftGraphCalendarPermission) UnsetEmailAddress()`

UnsetEmailAddress ensures that no value is present for EmailAddress, not even an explicit nil
### GetIsInsideOrganization

`func (o *MicrosoftGraphCalendarPermission) GetIsInsideOrganization() bool`

GetIsInsideOrganization returns the IsInsideOrganization field if non-nil, zero value otherwise.

### GetIsInsideOrganizationOk

`func (o *MicrosoftGraphCalendarPermission) GetIsInsideOrganizationOk() (*bool, bool)`

GetIsInsideOrganizationOk returns a tuple with the IsInsideOrganization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInsideOrganization

`func (o *MicrosoftGraphCalendarPermission) SetIsInsideOrganization(v bool)`

SetIsInsideOrganization sets IsInsideOrganization field to given value.

### HasIsInsideOrganization

`func (o *MicrosoftGraphCalendarPermission) HasIsInsideOrganization() bool`

HasIsInsideOrganization returns a boolean if a field has been set.

### SetIsInsideOrganizationNil

`func (o *MicrosoftGraphCalendarPermission) SetIsInsideOrganizationNil(b bool)`

 SetIsInsideOrganizationNil sets the value for IsInsideOrganization to be an explicit nil

### UnsetIsInsideOrganization
`func (o *MicrosoftGraphCalendarPermission) UnsetIsInsideOrganization()`

UnsetIsInsideOrganization ensures that no value is present for IsInsideOrganization, not even an explicit nil
### GetIsRemovable

`func (o *MicrosoftGraphCalendarPermission) GetIsRemovable() bool`

GetIsRemovable returns the IsRemovable field if non-nil, zero value otherwise.

### GetIsRemovableOk

`func (o *MicrosoftGraphCalendarPermission) GetIsRemovableOk() (*bool, bool)`

GetIsRemovableOk returns a tuple with the IsRemovable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRemovable

`func (o *MicrosoftGraphCalendarPermission) SetIsRemovable(v bool)`

SetIsRemovable sets IsRemovable field to given value.

### HasIsRemovable

`func (o *MicrosoftGraphCalendarPermission) HasIsRemovable() bool`

HasIsRemovable returns a boolean if a field has been set.

### SetIsRemovableNil

`func (o *MicrosoftGraphCalendarPermission) SetIsRemovableNil(b bool)`

 SetIsRemovableNil sets the value for IsRemovable to be an explicit nil

### UnsetIsRemovable
`func (o *MicrosoftGraphCalendarPermission) UnsetIsRemovable()`

UnsetIsRemovable ensures that no value is present for IsRemovable, not even an explicit nil
### GetRole

`func (o *MicrosoftGraphCalendarPermission) GetRole() AnyOfmicrosoftGraphCalendarRoleType`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *MicrosoftGraphCalendarPermission) GetRoleOk() (*AnyOfmicrosoftGraphCalendarRoleType, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *MicrosoftGraphCalendarPermission) SetRole(v AnyOfmicrosoftGraphCalendarRoleType)`

SetRole sets Role field to given value.

### HasRole

`func (o *MicrosoftGraphCalendarPermission) HasRole() bool`

HasRole returns a boolean if a field has been set.

### SetRoleNil

`func (o *MicrosoftGraphCalendarPermission) SetRoleNil(b bool)`

 SetRoleNil sets the value for Role to be an explicit nil

### UnsetRole
`func (o *MicrosoftGraphCalendarPermission) UnsetRole()`

UnsetRole ensures that no value is present for Role, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


