# RoleAssignment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **NullableString** | Description of the Role Assignment. | [optional] 
**DisplayName** | Pointer to **NullableString** | The display or friendly name of the role Assignment. | [optional] 
**ResourceScopes** | Pointer to **[]string** | List of ids of role scope member security groups.  These are IDs from Azure Active Directory. | [optional] 
**RoleDefinition** | Pointer to [**AnyOfmicrosoftGraphRoleDefinition**](anyOf&lt;microsoft.graph.roleDefinition&gt;.md) | Role definition this assignment is part of. | [optional] 

## Methods

### NewRoleAssignment

`func NewRoleAssignment() *RoleAssignment`

NewRoleAssignment instantiates a new RoleAssignment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoleAssignmentWithDefaults

`func NewRoleAssignmentWithDefaults() *RoleAssignment`

NewRoleAssignmentWithDefaults instantiates a new RoleAssignment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *RoleAssignment) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RoleAssignment) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RoleAssignment) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RoleAssignment) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *RoleAssignment) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *RoleAssignment) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDisplayName

`func (o *RoleAssignment) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *RoleAssignment) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *RoleAssignment) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *RoleAssignment) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *RoleAssignment) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *RoleAssignment) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetResourceScopes

`func (o *RoleAssignment) GetResourceScopes() []*string`

GetResourceScopes returns the ResourceScopes field if non-nil, zero value otherwise.

### GetResourceScopesOk

`func (o *RoleAssignment) GetResourceScopesOk() (*[]*string, bool)`

GetResourceScopesOk returns a tuple with the ResourceScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceScopes

`func (o *RoleAssignment) SetResourceScopes(v []*string)`

SetResourceScopes sets ResourceScopes field to given value.

### HasResourceScopes

`func (o *RoleAssignment) HasResourceScopes() bool`

HasResourceScopes returns a boolean if a field has been set.

### GetRoleDefinition

`func (o *RoleAssignment) GetRoleDefinition() AnyOfmicrosoftGraphRoleDefinition`

GetRoleDefinition returns the RoleDefinition field if non-nil, zero value otherwise.

### GetRoleDefinitionOk

`func (o *RoleAssignment) GetRoleDefinitionOk() (*AnyOfmicrosoftGraphRoleDefinition, bool)`

GetRoleDefinitionOk returns a tuple with the RoleDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleDefinition

`func (o *RoleAssignment) SetRoleDefinition(v AnyOfmicrosoftGraphRoleDefinition)`

SetRoleDefinition sets RoleDefinition field to given value.

### HasRoleDefinition

`func (o *RoleAssignment) HasRoleDefinition() bool`

HasRoleDefinition returns a boolean if a field has been set.

### SetRoleDefinitionNil

`func (o *RoleAssignment) SetRoleDefinitionNil(b bool)`

 SetRoleDefinitionNil sets the value for RoleDefinition to be an explicit nil

### UnsetRoleDefinition
`func (o *RoleAssignment) UnsetRoleDefinition()`

UnsetRoleDefinition ensures that no value is present for RoleDefinition, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


