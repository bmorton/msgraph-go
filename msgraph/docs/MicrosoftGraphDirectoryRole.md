# MicrosoftGraphDirectoryRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**DeletedDateTime** | Pointer to **NullableTime** |  | [optional] 
**Description** | Pointer to **NullableString** | The description for the directory role. Read-only. | [optional] 
**DisplayName** | Pointer to **NullableString** | The display name for the directory role. Read-only. | [optional] 
**RoleTemplateId** | Pointer to **NullableString** | The id of the directoryRoleTemplate that this role is based on. The property must be specified when activating a directory role in a tenant with a POST operation. After the directory role has been activated, the property is read only. | [optional] 
**Members** | Pointer to [**[]MicrosoftGraphDirectoryObject**](MicrosoftGraphDirectoryObject.md) | Users that are members of this directory role. HTTP Methods: GET, POST, DELETE. Read-only. Nullable. | [optional] 
**ScopedMembers** | Pointer to [**[]MicrosoftGraphScopedRoleMembership**](MicrosoftGraphScopedRoleMembership.md) | Members of this directory role that are scoped to administrative units. Read-only. Nullable. | [optional] 

## Methods

### NewMicrosoftGraphDirectoryRole

`func NewMicrosoftGraphDirectoryRole() *MicrosoftGraphDirectoryRole`

NewMicrosoftGraphDirectoryRole instantiates a new MicrosoftGraphDirectoryRole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphDirectoryRoleWithDefaults

`func NewMicrosoftGraphDirectoryRoleWithDefaults() *MicrosoftGraphDirectoryRole`

NewMicrosoftGraphDirectoryRoleWithDefaults instantiates a new MicrosoftGraphDirectoryRole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphDirectoryRole) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphDirectoryRole) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphDirectoryRole) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphDirectoryRole) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDeletedDateTime

`func (o *MicrosoftGraphDirectoryRole) GetDeletedDateTime() time.Time`

GetDeletedDateTime returns the DeletedDateTime field if non-nil, zero value otherwise.

### GetDeletedDateTimeOk

`func (o *MicrosoftGraphDirectoryRole) GetDeletedDateTimeOk() (*time.Time, bool)`

GetDeletedDateTimeOk returns a tuple with the DeletedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedDateTime

`func (o *MicrosoftGraphDirectoryRole) SetDeletedDateTime(v time.Time)`

SetDeletedDateTime sets DeletedDateTime field to given value.

### HasDeletedDateTime

`func (o *MicrosoftGraphDirectoryRole) HasDeletedDateTime() bool`

HasDeletedDateTime returns a boolean if a field has been set.

### SetDeletedDateTimeNil

`func (o *MicrosoftGraphDirectoryRole) SetDeletedDateTimeNil(b bool)`

 SetDeletedDateTimeNil sets the value for DeletedDateTime to be an explicit nil

### UnsetDeletedDateTime
`func (o *MicrosoftGraphDirectoryRole) UnsetDeletedDateTime()`

UnsetDeletedDateTime ensures that no value is present for DeletedDateTime, not even an explicit nil
### GetDescription

`func (o *MicrosoftGraphDirectoryRole) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MicrosoftGraphDirectoryRole) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MicrosoftGraphDirectoryRole) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MicrosoftGraphDirectoryRole) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *MicrosoftGraphDirectoryRole) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *MicrosoftGraphDirectoryRole) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDisplayName

`func (o *MicrosoftGraphDirectoryRole) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphDirectoryRole) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *MicrosoftGraphDirectoryRole) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *MicrosoftGraphDirectoryRole) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *MicrosoftGraphDirectoryRole) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *MicrosoftGraphDirectoryRole) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetRoleTemplateId

`func (o *MicrosoftGraphDirectoryRole) GetRoleTemplateId() string`

GetRoleTemplateId returns the RoleTemplateId field if non-nil, zero value otherwise.

### GetRoleTemplateIdOk

`func (o *MicrosoftGraphDirectoryRole) GetRoleTemplateIdOk() (*string, bool)`

GetRoleTemplateIdOk returns a tuple with the RoleTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleTemplateId

`func (o *MicrosoftGraphDirectoryRole) SetRoleTemplateId(v string)`

SetRoleTemplateId sets RoleTemplateId field to given value.

### HasRoleTemplateId

`func (o *MicrosoftGraphDirectoryRole) HasRoleTemplateId() bool`

HasRoleTemplateId returns a boolean if a field has been set.

### SetRoleTemplateIdNil

`func (o *MicrosoftGraphDirectoryRole) SetRoleTemplateIdNil(b bool)`

 SetRoleTemplateIdNil sets the value for RoleTemplateId to be an explicit nil

### UnsetRoleTemplateId
`func (o *MicrosoftGraphDirectoryRole) UnsetRoleTemplateId()`

UnsetRoleTemplateId ensures that no value is present for RoleTemplateId, not even an explicit nil
### GetMembers

`func (o *MicrosoftGraphDirectoryRole) GetMembers() []MicrosoftGraphDirectoryObject`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *MicrosoftGraphDirectoryRole) GetMembersOk() (*[]MicrosoftGraphDirectoryObject, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *MicrosoftGraphDirectoryRole) SetMembers(v []MicrosoftGraphDirectoryObject)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *MicrosoftGraphDirectoryRole) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### GetScopedMembers

`func (o *MicrosoftGraphDirectoryRole) GetScopedMembers() []MicrosoftGraphScopedRoleMembership`

GetScopedMembers returns the ScopedMembers field if non-nil, zero value otherwise.

### GetScopedMembersOk

`func (o *MicrosoftGraphDirectoryRole) GetScopedMembersOk() (*[]MicrosoftGraphScopedRoleMembership, bool)`

GetScopedMembersOk returns a tuple with the ScopedMembers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopedMembers

`func (o *MicrosoftGraphDirectoryRole) SetScopedMembers(v []MicrosoftGraphScopedRoleMembership)`

SetScopedMembers sets ScopedMembers field to given value.

### HasScopedMembers

`func (o *MicrosoftGraphDirectoryRole) HasScopedMembers() bool`

HasScopedMembers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


