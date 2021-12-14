# RoleDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **NullableString** | Description of the Role definition. | [optional] 
**DisplayName** | Pointer to **NullableString** | Display Name of the Role definition. | [optional] 
**IsBuiltIn** | Pointer to **bool** | Type of Role. Set to True if it is built-in, or set to False if it is a custom role definition. | [optional] 
**RolePermissions** | Pointer to [**[]AnyOfmicrosoftGraphRolePermission**](AnyOfmicrosoftGraphRolePermission.md) | List of Role Permissions this role is allowed to perform. These must match the actionName that is defined as part of the rolePermission. | [optional] 
**RoleAssignments** | Pointer to [**[]MicrosoftGraphRoleAssignment**](MicrosoftGraphRoleAssignment.md) | List of Role assignments for this role definition. | [optional] 

## Methods

### NewRoleDefinition

`func NewRoleDefinition() *RoleDefinition`

NewRoleDefinition instantiates a new RoleDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoleDefinitionWithDefaults

`func NewRoleDefinitionWithDefaults() *RoleDefinition`

NewRoleDefinitionWithDefaults instantiates a new RoleDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *RoleDefinition) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RoleDefinition) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RoleDefinition) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RoleDefinition) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *RoleDefinition) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *RoleDefinition) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDisplayName

`func (o *RoleDefinition) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *RoleDefinition) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *RoleDefinition) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *RoleDefinition) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *RoleDefinition) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *RoleDefinition) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetIsBuiltIn

`func (o *RoleDefinition) GetIsBuiltIn() bool`

GetIsBuiltIn returns the IsBuiltIn field if non-nil, zero value otherwise.

### GetIsBuiltInOk

`func (o *RoleDefinition) GetIsBuiltInOk() (*bool, bool)`

GetIsBuiltInOk returns a tuple with the IsBuiltIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBuiltIn

`func (o *RoleDefinition) SetIsBuiltIn(v bool)`

SetIsBuiltIn sets IsBuiltIn field to given value.

### HasIsBuiltIn

`func (o *RoleDefinition) HasIsBuiltIn() bool`

HasIsBuiltIn returns a boolean if a field has been set.

### GetRolePermissions

`func (o *RoleDefinition) GetRolePermissions() []*AnyOfmicrosoftGraphRolePermission`

GetRolePermissions returns the RolePermissions field if non-nil, zero value otherwise.

### GetRolePermissionsOk

`func (o *RoleDefinition) GetRolePermissionsOk() (*[]*AnyOfmicrosoftGraphRolePermission, bool)`

GetRolePermissionsOk returns a tuple with the RolePermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRolePermissions

`func (o *RoleDefinition) SetRolePermissions(v []*AnyOfmicrosoftGraphRolePermission)`

SetRolePermissions sets RolePermissions field to given value.

### HasRolePermissions

`func (o *RoleDefinition) HasRolePermissions() bool`

HasRolePermissions returns a boolean if a field has been set.

### GetRoleAssignments

`func (o *RoleDefinition) GetRoleAssignments() []MicrosoftGraphRoleAssignment`

GetRoleAssignments returns the RoleAssignments field if non-nil, zero value otherwise.

### GetRoleAssignmentsOk

`func (o *RoleDefinition) GetRoleAssignmentsOk() (*[]MicrosoftGraphRoleAssignment, bool)`

GetRoleAssignmentsOk returns a tuple with the RoleAssignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleAssignments

`func (o *RoleDefinition) SetRoleAssignments(v []MicrosoftGraphRoleAssignment)`

SetRoleAssignments sets RoleAssignments field to given value.

### HasRoleAssignments

`func (o *RoleDefinition) HasRoleAssignments() bool`

HasRoleAssignments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


