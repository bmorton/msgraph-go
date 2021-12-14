# MicrosoftGraphDeviceAndAppManagementRoleAssignment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**Description** | Pointer to **NullableString** | Description of the Role Assignment. | [optional] 
**DisplayName** | Pointer to **NullableString** | The display or friendly name of the role Assignment. | [optional] 
**ResourceScopes** | Pointer to **[]string** | List of ids of role scope member security groups.  These are IDs from Azure Active Directory. | [optional] 
**RoleDefinition** | Pointer to [**AnyOfmicrosoftGraphRoleDefinition**](anyOf&lt;microsoft.graph.roleDefinition&gt;.md) | Role definition this assignment is part of. | [optional] 
**Members** | Pointer to **[]string** | The list of ids of role member security groups. These are IDs from Azure Active Directory. | [optional] 

## Methods

### NewMicrosoftGraphDeviceAndAppManagementRoleAssignment

`func NewMicrosoftGraphDeviceAndAppManagementRoleAssignment() *MicrosoftGraphDeviceAndAppManagementRoleAssignment`

NewMicrosoftGraphDeviceAndAppManagementRoleAssignment instantiates a new MicrosoftGraphDeviceAndAppManagementRoleAssignment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphDeviceAndAppManagementRoleAssignmentWithDefaults

`func NewMicrosoftGraphDeviceAndAppManagementRoleAssignmentWithDefaults() *MicrosoftGraphDeviceAndAppManagementRoleAssignment`

NewMicrosoftGraphDeviceAndAppManagementRoleAssignmentWithDefaults instantiates a new MicrosoftGraphDeviceAndAppManagementRoleAssignment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphDeviceAndAppManagementRoleAssignment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphDeviceAndAppManagementRoleAssignment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphDeviceAndAppManagementRoleAssignment) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphDeviceAndAppManagementRoleAssignment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDescription

`func (o *MicrosoftGraphDeviceAndAppManagementRoleAssignment) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MicrosoftGraphDeviceAndAppManagementRoleAssignment) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MicrosoftGraphDeviceAndAppManagementRoleAssignment) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MicrosoftGraphDeviceAndAppManagementRoleAssignment) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *MicrosoftGraphDeviceAndAppManagementRoleAssignment) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *MicrosoftGraphDeviceAndAppManagementRoleAssignment) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDisplayName

`func (o *MicrosoftGraphDeviceAndAppManagementRoleAssignment) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphDeviceAndAppManagementRoleAssignment) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *MicrosoftGraphDeviceAndAppManagementRoleAssignment) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *MicrosoftGraphDeviceAndAppManagementRoleAssignment) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *MicrosoftGraphDeviceAndAppManagementRoleAssignment) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *MicrosoftGraphDeviceAndAppManagementRoleAssignment) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetResourceScopes

`func (o *MicrosoftGraphDeviceAndAppManagementRoleAssignment) GetResourceScopes() []*string`

GetResourceScopes returns the ResourceScopes field if non-nil, zero value otherwise.

### GetResourceScopesOk

`func (o *MicrosoftGraphDeviceAndAppManagementRoleAssignment) GetResourceScopesOk() (*[]*string, bool)`

GetResourceScopesOk returns a tuple with the ResourceScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceScopes

`func (o *MicrosoftGraphDeviceAndAppManagementRoleAssignment) SetResourceScopes(v []*string)`

SetResourceScopes sets ResourceScopes field to given value.

### HasResourceScopes

`func (o *MicrosoftGraphDeviceAndAppManagementRoleAssignment) HasResourceScopes() bool`

HasResourceScopes returns a boolean if a field has been set.

### GetRoleDefinition

`func (o *MicrosoftGraphDeviceAndAppManagementRoleAssignment) GetRoleDefinition() AnyOfmicrosoftGraphRoleDefinition`

GetRoleDefinition returns the RoleDefinition field if non-nil, zero value otherwise.

### GetRoleDefinitionOk

`func (o *MicrosoftGraphDeviceAndAppManagementRoleAssignment) GetRoleDefinitionOk() (*AnyOfmicrosoftGraphRoleDefinition, bool)`

GetRoleDefinitionOk returns a tuple with the RoleDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleDefinition

`func (o *MicrosoftGraphDeviceAndAppManagementRoleAssignment) SetRoleDefinition(v AnyOfmicrosoftGraphRoleDefinition)`

SetRoleDefinition sets RoleDefinition field to given value.

### HasRoleDefinition

`func (o *MicrosoftGraphDeviceAndAppManagementRoleAssignment) HasRoleDefinition() bool`

HasRoleDefinition returns a boolean if a field has been set.

### SetRoleDefinitionNil

`func (o *MicrosoftGraphDeviceAndAppManagementRoleAssignment) SetRoleDefinitionNil(b bool)`

 SetRoleDefinitionNil sets the value for RoleDefinition to be an explicit nil

### UnsetRoleDefinition
`func (o *MicrosoftGraphDeviceAndAppManagementRoleAssignment) UnsetRoleDefinition()`

UnsetRoleDefinition ensures that no value is present for RoleDefinition, not even an explicit nil
### GetMembers

`func (o *MicrosoftGraphDeviceAndAppManagementRoleAssignment) GetMembers() []*string`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *MicrosoftGraphDeviceAndAppManagementRoleAssignment) GetMembersOk() (*[]*string, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *MicrosoftGraphDeviceAndAppManagementRoleAssignment) SetMembers(v []*string)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *MicrosoftGraphDeviceAndAppManagementRoleAssignment) HasMembers() bool`

HasMembers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


