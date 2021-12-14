# MicrosoftGraphRoleAssignment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**Description** | Pointer to **NullableString** | Description of the Role Assignment. | [optional] 
**DisplayName** | Pointer to **NullableString** | The display or friendly name of the role Assignment. | [optional] 
**ResourceScopes** | Pointer to **[]string** | List of ids of role scope member security groups.  These are IDs from Azure Active Directory. | [optional] 
**RoleDefinition** | Pointer to [**AnyOfmicrosoftGraphRoleDefinition**](anyOf&lt;microsoft.graph.roleDefinition&gt;.md) | Role definition this assignment is part of. | [optional] 

## Methods

### NewMicrosoftGraphRoleAssignment

`func NewMicrosoftGraphRoleAssignment() *MicrosoftGraphRoleAssignment`

NewMicrosoftGraphRoleAssignment instantiates a new MicrosoftGraphRoleAssignment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphRoleAssignmentWithDefaults

`func NewMicrosoftGraphRoleAssignmentWithDefaults() *MicrosoftGraphRoleAssignment`

NewMicrosoftGraphRoleAssignmentWithDefaults instantiates a new MicrosoftGraphRoleAssignment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphRoleAssignment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphRoleAssignment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphRoleAssignment) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphRoleAssignment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDescription

`func (o *MicrosoftGraphRoleAssignment) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MicrosoftGraphRoleAssignment) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MicrosoftGraphRoleAssignment) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MicrosoftGraphRoleAssignment) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *MicrosoftGraphRoleAssignment) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *MicrosoftGraphRoleAssignment) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDisplayName

`func (o *MicrosoftGraphRoleAssignment) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphRoleAssignment) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *MicrosoftGraphRoleAssignment) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *MicrosoftGraphRoleAssignment) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *MicrosoftGraphRoleAssignment) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *MicrosoftGraphRoleAssignment) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetResourceScopes

`func (o *MicrosoftGraphRoleAssignment) GetResourceScopes() []*string`

GetResourceScopes returns the ResourceScopes field if non-nil, zero value otherwise.

### GetResourceScopesOk

`func (o *MicrosoftGraphRoleAssignment) GetResourceScopesOk() (*[]*string, bool)`

GetResourceScopesOk returns a tuple with the ResourceScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceScopes

`func (o *MicrosoftGraphRoleAssignment) SetResourceScopes(v []*string)`

SetResourceScopes sets ResourceScopes field to given value.

### HasResourceScopes

`func (o *MicrosoftGraphRoleAssignment) HasResourceScopes() bool`

HasResourceScopes returns a boolean if a field has been set.

### GetRoleDefinition

`func (o *MicrosoftGraphRoleAssignment) GetRoleDefinition() AnyOfmicrosoftGraphRoleDefinition`

GetRoleDefinition returns the RoleDefinition field if non-nil, zero value otherwise.

### GetRoleDefinitionOk

`func (o *MicrosoftGraphRoleAssignment) GetRoleDefinitionOk() (*AnyOfmicrosoftGraphRoleDefinition, bool)`

GetRoleDefinitionOk returns a tuple with the RoleDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleDefinition

`func (o *MicrosoftGraphRoleAssignment) SetRoleDefinition(v AnyOfmicrosoftGraphRoleDefinition)`

SetRoleDefinition sets RoleDefinition field to given value.

### HasRoleDefinition

`func (o *MicrosoftGraphRoleAssignment) HasRoleDefinition() bool`

HasRoleDefinition returns a boolean if a field has been set.

### SetRoleDefinitionNil

`func (o *MicrosoftGraphRoleAssignment) SetRoleDefinitionNil(b bool)`

 SetRoleDefinitionNil sets the value for RoleDefinition to be an explicit nil

### UnsetRoleDefinition
`func (o *MicrosoftGraphRoleAssignment) UnsetRoleDefinition()`

UnsetRoleDefinition ensures that no value is present for RoleDefinition, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


