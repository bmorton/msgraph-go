# MicrosoftGraphRbacApplication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**RoleAssignments** | Pointer to [**[]MicrosoftGraphUnifiedRoleAssignment**](MicrosoftGraphUnifiedRoleAssignment.md) | Resource to grant access to users or groups. | [optional] 
**RoleDefinitions** | Pointer to [**[]MicrosoftGraphUnifiedRoleDefinition**](MicrosoftGraphUnifiedRoleDefinition.md) | Resource representing the roles allowed by RBAC providers and the permissions assigned to the roles. | [optional] 

## Methods

### NewMicrosoftGraphRbacApplication

`func NewMicrosoftGraphRbacApplication() *MicrosoftGraphRbacApplication`

NewMicrosoftGraphRbacApplication instantiates a new MicrosoftGraphRbacApplication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphRbacApplicationWithDefaults

`func NewMicrosoftGraphRbacApplicationWithDefaults() *MicrosoftGraphRbacApplication`

NewMicrosoftGraphRbacApplicationWithDefaults instantiates a new MicrosoftGraphRbacApplication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphRbacApplication) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphRbacApplication) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphRbacApplication) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphRbacApplication) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRoleAssignments

`func (o *MicrosoftGraphRbacApplication) GetRoleAssignments() []MicrosoftGraphUnifiedRoleAssignment`

GetRoleAssignments returns the RoleAssignments field if non-nil, zero value otherwise.

### GetRoleAssignmentsOk

`func (o *MicrosoftGraphRbacApplication) GetRoleAssignmentsOk() (*[]MicrosoftGraphUnifiedRoleAssignment, bool)`

GetRoleAssignmentsOk returns a tuple with the RoleAssignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleAssignments

`func (o *MicrosoftGraphRbacApplication) SetRoleAssignments(v []MicrosoftGraphUnifiedRoleAssignment)`

SetRoleAssignments sets RoleAssignments field to given value.

### HasRoleAssignments

`func (o *MicrosoftGraphRbacApplication) HasRoleAssignments() bool`

HasRoleAssignments returns a boolean if a field has been set.

### GetRoleDefinitions

`func (o *MicrosoftGraphRbacApplication) GetRoleDefinitions() []MicrosoftGraphUnifiedRoleDefinition`

GetRoleDefinitions returns the RoleDefinitions field if non-nil, zero value otherwise.

### GetRoleDefinitionsOk

`func (o *MicrosoftGraphRbacApplication) GetRoleDefinitionsOk() (*[]MicrosoftGraphUnifiedRoleDefinition, bool)`

GetRoleDefinitionsOk returns a tuple with the RoleDefinitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleDefinitions

`func (o *MicrosoftGraphRbacApplication) SetRoleDefinitions(v []MicrosoftGraphUnifiedRoleDefinition)`

SetRoleDefinitions sets RoleDefinitions field to given value.

### HasRoleDefinitions

`func (o *MicrosoftGraphRbacApplication) HasRoleDefinitions() bool`

HasRoleDefinitions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


