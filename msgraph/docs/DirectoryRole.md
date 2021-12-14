# DirectoryRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **NullableString** | The description for the directory role. Read-only. | [optional] 
**DisplayName** | Pointer to **NullableString** | The display name for the directory role. Read-only. | [optional] 
**RoleTemplateId** | Pointer to **NullableString** | The id of the directoryRoleTemplate that this role is based on. The property must be specified when activating a directory role in a tenant with a POST operation. After the directory role has been activated, the property is read only. | [optional] 
**Members** | Pointer to [**[]MicrosoftGraphDirectoryObject**](MicrosoftGraphDirectoryObject.md) | Users that are members of this directory role. HTTP Methods: GET, POST, DELETE. Read-only. Nullable. | [optional] 
**ScopedMembers** | Pointer to [**[]MicrosoftGraphScopedRoleMembership**](MicrosoftGraphScopedRoleMembership.md) | Members of this directory role that are scoped to administrative units. Read-only. Nullable. | [optional] 

## Methods

### NewDirectoryRole

`func NewDirectoryRole() *DirectoryRole`

NewDirectoryRole instantiates a new DirectoryRole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDirectoryRoleWithDefaults

`func NewDirectoryRoleWithDefaults() *DirectoryRole`

NewDirectoryRoleWithDefaults instantiates a new DirectoryRole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *DirectoryRole) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DirectoryRole) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DirectoryRole) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DirectoryRole) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *DirectoryRole) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *DirectoryRole) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDisplayName

`func (o *DirectoryRole) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *DirectoryRole) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *DirectoryRole) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *DirectoryRole) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *DirectoryRole) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *DirectoryRole) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetRoleTemplateId

`func (o *DirectoryRole) GetRoleTemplateId() string`

GetRoleTemplateId returns the RoleTemplateId field if non-nil, zero value otherwise.

### GetRoleTemplateIdOk

`func (o *DirectoryRole) GetRoleTemplateIdOk() (*string, bool)`

GetRoleTemplateIdOk returns a tuple with the RoleTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleTemplateId

`func (o *DirectoryRole) SetRoleTemplateId(v string)`

SetRoleTemplateId sets RoleTemplateId field to given value.

### HasRoleTemplateId

`func (o *DirectoryRole) HasRoleTemplateId() bool`

HasRoleTemplateId returns a boolean if a field has been set.

### SetRoleTemplateIdNil

`func (o *DirectoryRole) SetRoleTemplateIdNil(b bool)`

 SetRoleTemplateIdNil sets the value for RoleTemplateId to be an explicit nil

### UnsetRoleTemplateId
`func (o *DirectoryRole) UnsetRoleTemplateId()`

UnsetRoleTemplateId ensures that no value is present for RoleTemplateId, not even an explicit nil
### GetMembers

`func (o *DirectoryRole) GetMembers() []MicrosoftGraphDirectoryObject`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *DirectoryRole) GetMembersOk() (*[]MicrosoftGraphDirectoryObject, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *DirectoryRole) SetMembers(v []MicrosoftGraphDirectoryObject)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *DirectoryRole) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### GetScopedMembers

`func (o *DirectoryRole) GetScopedMembers() []MicrosoftGraphScopedRoleMembership`

GetScopedMembers returns the ScopedMembers field if non-nil, zero value otherwise.

### GetScopedMembersOk

`func (o *DirectoryRole) GetScopedMembersOk() (*[]MicrosoftGraphScopedRoleMembership, bool)`

GetScopedMembersOk returns a tuple with the ScopedMembers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopedMembers

`func (o *DirectoryRole) SetScopedMembers(v []MicrosoftGraphScopedRoleMembership)`

SetScopedMembers sets ScopedMembers field to given value.

### HasScopedMembers

`func (o *DirectoryRole) HasScopedMembers() bool`

HasScopedMembers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


