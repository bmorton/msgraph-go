# RbacApplication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RoleAssignments** | Pointer to [**[]MicrosoftGraphUnifiedRoleAssignment**](MicrosoftGraphUnifiedRoleAssignment.md) | Resource to grant access to users or groups. | [optional] 
**RoleDefinitions** | Pointer to [**[]MicrosoftGraphUnifiedRoleDefinition**](MicrosoftGraphUnifiedRoleDefinition.md) | Resource representing the roles allowed by RBAC providers and the permissions assigned to the roles. | [optional] 

## Methods

### NewRbacApplication

`func NewRbacApplication() *RbacApplication`

NewRbacApplication instantiates a new RbacApplication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRbacApplicationWithDefaults

`func NewRbacApplicationWithDefaults() *RbacApplication`

NewRbacApplicationWithDefaults instantiates a new RbacApplication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoleAssignments

`func (o *RbacApplication) GetRoleAssignments() []MicrosoftGraphUnifiedRoleAssignment`

GetRoleAssignments returns the RoleAssignments field if non-nil, zero value otherwise.

### GetRoleAssignmentsOk

`func (o *RbacApplication) GetRoleAssignmentsOk() (*[]MicrosoftGraphUnifiedRoleAssignment, bool)`

GetRoleAssignmentsOk returns a tuple with the RoleAssignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleAssignments

`func (o *RbacApplication) SetRoleAssignments(v []MicrosoftGraphUnifiedRoleAssignment)`

SetRoleAssignments sets RoleAssignments field to given value.

### HasRoleAssignments

`func (o *RbacApplication) HasRoleAssignments() bool`

HasRoleAssignments returns a boolean if a field has been set.

### GetRoleDefinitions

`func (o *RbacApplication) GetRoleDefinitions() []MicrosoftGraphUnifiedRoleDefinition`

GetRoleDefinitions returns the RoleDefinitions field if non-nil, zero value otherwise.

### GetRoleDefinitionsOk

`func (o *RbacApplication) GetRoleDefinitionsOk() (*[]MicrosoftGraphUnifiedRoleDefinition, bool)`

GetRoleDefinitionsOk returns a tuple with the RoleDefinitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleDefinitions

`func (o *RbacApplication) SetRoleDefinitions(v []MicrosoftGraphUnifiedRoleDefinition)`

SetRoleDefinitions sets RoleDefinitions field to given value.

### HasRoleDefinitions

`func (o *RbacApplication) HasRoleDefinitions() bool`

HasRoleDefinitions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


