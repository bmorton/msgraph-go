# MicrosoftGraphUnifiedRoleAssignment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**AppScopeId** | Pointer to **NullableString** | Identifier of the app-specific scope when the assignment scope is app-specific.  Either this property or directoryScopeId is required. App scopes are scopes that are defined and understood by this application only. Use / for tenant-wide app scopes. Use directoryScopeId to limit the scope to particular directory objects, for example, administrative units. Supports $filter (eq, in). | [optional] 
**Condition** | Pointer to **NullableString** |  | [optional] 
**DirectoryScopeId** | Pointer to **NullableString** | Identifier of the directory object representing the scope of the assignment.  Either this property or appScopeId is required. The scope of an assignment determines the set of resources for which the principal has been granted access. Directory scopes are shared scopes stored in the directory that are understood by multiple applications. Use / for tenant-wide scope. Use appScopeId to limit the scope to an application only. Supports $filter (eq, in). | [optional] 
**PrincipalId** | Pointer to **NullableString** | Identifier of the principal to which the assignment is granted. Supports $filter (eq, in). | [optional] 
**RoleDefinitionId** | Pointer to **NullableString** | Identifier of the role definition the assignment is for. Read only. Supports $filter (eq, in). | [optional] 
**AppScope** | Pointer to [**AnyOfmicrosoftGraphAppScope**](anyOf&lt;microsoft.graph.appScope&gt;.md) | Read-only property with details of the app specific scope when the assignment scope is app specific. Containment entity. Supports $expand. | [optional] 
**DirectoryScope** | Pointer to [**AnyOfmicrosoftGraphDirectoryObject**](anyOf&lt;microsoft.graph.directoryObject&gt;.md) | The directory object that is the scope of the assignment. Read-only. Supports $expand. | [optional] 
**Principal** | Pointer to [**AnyOfmicrosoftGraphDirectoryObject**](anyOf&lt;microsoft.graph.directoryObject&gt;.md) | Referencing the assigned principal. Read-only. Supports $expand. | [optional] 
**RoleDefinition** | Pointer to [**AnyOfmicrosoftGraphUnifiedRoleDefinition**](anyOf&lt;microsoft.graph.unifiedRoleDefinition&gt;.md) | The roleDefinition the assignment is for.  Supports $expand. roleDefinition.Id will be auto expanded. | [optional] 

## Methods

### NewMicrosoftGraphUnifiedRoleAssignment

`func NewMicrosoftGraphUnifiedRoleAssignment() *MicrosoftGraphUnifiedRoleAssignment`

NewMicrosoftGraphUnifiedRoleAssignment instantiates a new MicrosoftGraphUnifiedRoleAssignment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphUnifiedRoleAssignmentWithDefaults

`func NewMicrosoftGraphUnifiedRoleAssignmentWithDefaults() *MicrosoftGraphUnifiedRoleAssignment`

NewMicrosoftGraphUnifiedRoleAssignmentWithDefaults instantiates a new MicrosoftGraphUnifiedRoleAssignment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphUnifiedRoleAssignment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphUnifiedRoleAssignment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphUnifiedRoleAssignment) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphUnifiedRoleAssignment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAppScopeId

`func (o *MicrosoftGraphUnifiedRoleAssignment) GetAppScopeId() string`

GetAppScopeId returns the AppScopeId field if non-nil, zero value otherwise.

### GetAppScopeIdOk

`func (o *MicrosoftGraphUnifiedRoleAssignment) GetAppScopeIdOk() (*string, bool)`

GetAppScopeIdOk returns a tuple with the AppScopeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppScopeId

`func (o *MicrosoftGraphUnifiedRoleAssignment) SetAppScopeId(v string)`

SetAppScopeId sets AppScopeId field to given value.

### HasAppScopeId

`func (o *MicrosoftGraphUnifiedRoleAssignment) HasAppScopeId() bool`

HasAppScopeId returns a boolean if a field has been set.

### SetAppScopeIdNil

`func (o *MicrosoftGraphUnifiedRoleAssignment) SetAppScopeIdNil(b bool)`

 SetAppScopeIdNil sets the value for AppScopeId to be an explicit nil

### UnsetAppScopeId
`func (o *MicrosoftGraphUnifiedRoleAssignment) UnsetAppScopeId()`

UnsetAppScopeId ensures that no value is present for AppScopeId, not even an explicit nil
### GetCondition

`func (o *MicrosoftGraphUnifiedRoleAssignment) GetCondition() string`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *MicrosoftGraphUnifiedRoleAssignment) GetConditionOk() (*string, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *MicrosoftGraphUnifiedRoleAssignment) SetCondition(v string)`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *MicrosoftGraphUnifiedRoleAssignment) HasCondition() bool`

HasCondition returns a boolean if a field has been set.

### SetConditionNil

`func (o *MicrosoftGraphUnifiedRoleAssignment) SetConditionNil(b bool)`

 SetConditionNil sets the value for Condition to be an explicit nil

### UnsetCondition
`func (o *MicrosoftGraphUnifiedRoleAssignment) UnsetCondition()`

UnsetCondition ensures that no value is present for Condition, not even an explicit nil
### GetDirectoryScopeId

`func (o *MicrosoftGraphUnifiedRoleAssignment) GetDirectoryScopeId() string`

GetDirectoryScopeId returns the DirectoryScopeId field if non-nil, zero value otherwise.

### GetDirectoryScopeIdOk

`func (o *MicrosoftGraphUnifiedRoleAssignment) GetDirectoryScopeIdOk() (*string, bool)`

GetDirectoryScopeIdOk returns a tuple with the DirectoryScopeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectoryScopeId

`func (o *MicrosoftGraphUnifiedRoleAssignment) SetDirectoryScopeId(v string)`

SetDirectoryScopeId sets DirectoryScopeId field to given value.

### HasDirectoryScopeId

`func (o *MicrosoftGraphUnifiedRoleAssignment) HasDirectoryScopeId() bool`

HasDirectoryScopeId returns a boolean if a field has been set.

### SetDirectoryScopeIdNil

`func (o *MicrosoftGraphUnifiedRoleAssignment) SetDirectoryScopeIdNil(b bool)`

 SetDirectoryScopeIdNil sets the value for DirectoryScopeId to be an explicit nil

### UnsetDirectoryScopeId
`func (o *MicrosoftGraphUnifiedRoleAssignment) UnsetDirectoryScopeId()`

UnsetDirectoryScopeId ensures that no value is present for DirectoryScopeId, not even an explicit nil
### GetPrincipalId

`func (o *MicrosoftGraphUnifiedRoleAssignment) GetPrincipalId() string`

GetPrincipalId returns the PrincipalId field if non-nil, zero value otherwise.

### GetPrincipalIdOk

`func (o *MicrosoftGraphUnifiedRoleAssignment) GetPrincipalIdOk() (*string, bool)`

GetPrincipalIdOk returns a tuple with the PrincipalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipalId

`func (o *MicrosoftGraphUnifiedRoleAssignment) SetPrincipalId(v string)`

SetPrincipalId sets PrincipalId field to given value.

### HasPrincipalId

`func (o *MicrosoftGraphUnifiedRoleAssignment) HasPrincipalId() bool`

HasPrincipalId returns a boolean if a field has been set.

### SetPrincipalIdNil

`func (o *MicrosoftGraphUnifiedRoleAssignment) SetPrincipalIdNil(b bool)`

 SetPrincipalIdNil sets the value for PrincipalId to be an explicit nil

### UnsetPrincipalId
`func (o *MicrosoftGraphUnifiedRoleAssignment) UnsetPrincipalId()`

UnsetPrincipalId ensures that no value is present for PrincipalId, not even an explicit nil
### GetRoleDefinitionId

`func (o *MicrosoftGraphUnifiedRoleAssignment) GetRoleDefinitionId() string`

GetRoleDefinitionId returns the RoleDefinitionId field if non-nil, zero value otherwise.

### GetRoleDefinitionIdOk

`func (o *MicrosoftGraphUnifiedRoleAssignment) GetRoleDefinitionIdOk() (*string, bool)`

GetRoleDefinitionIdOk returns a tuple with the RoleDefinitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleDefinitionId

`func (o *MicrosoftGraphUnifiedRoleAssignment) SetRoleDefinitionId(v string)`

SetRoleDefinitionId sets RoleDefinitionId field to given value.

### HasRoleDefinitionId

`func (o *MicrosoftGraphUnifiedRoleAssignment) HasRoleDefinitionId() bool`

HasRoleDefinitionId returns a boolean if a field has been set.

### SetRoleDefinitionIdNil

`func (o *MicrosoftGraphUnifiedRoleAssignment) SetRoleDefinitionIdNil(b bool)`

 SetRoleDefinitionIdNil sets the value for RoleDefinitionId to be an explicit nil

### UnsetRoleDefinitionId
`func (o *MicrosoftGraphUnifiedRoleAssignment) UnsetRoleDefinitionId()`

UnsetRoleDefinitionId ensures that no value is present for RoleDefinitionId, not even an explicit nil
### GetAppScope

`func (o *MicrosoftGraphUnifiedRoleAssignment) GetAppScope() AnyOfmicrosoftGraphAppScope`

GetAppScope returns the AppScope field if non-nil, zero value otherwise.

### GetAppScopeOk

`func (o *MicrosoftGraphUnifiedRoleAssignment) GetAppScopeOk() (*AnyOfmicrosoftGraphAppScope, bool)`

GetAppScopeOk returns a tuple with the AppScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppScope

`func (o *MicrosoftGraphUnifiedRoleAssignment) SetAppScope(v AnyOfmicrosoftGraphAppScope)`

SetAppScope sets AppScope field to given value.

### HasAppScope

`func (o *MicrosoftGraphUnifiedRoleAssignment) HasAppScope() bool`

HasAppScope returns a boolean if a field has been set.

### SetAppScopeNil

`func (o *MicrosoftGraphUnifiedRoleAssignment) SetAppScopeNil(b bool)`

 SetAppScopeNil sets the value for AppScope to be an explicit nil

### UnsetAppScope
`func (o *MicrosoftGraphUnifiedRoleAssignment) UnsetAppScope()`

UnsetAppScope ensures that no value is present for AppScope, not even an explicit nil
### GetDirectoryScope

`func (o *MicrosoftGraphUnifiedRoleAssignment) GetDirectoryScope() AnyOfmicrosoftGraphDirectoryObject`

GetDirectoryScope returns the DirectoryScope field if non-nil, zero value otherwise.

### GetDirectoryScopeOk

`func (o *MicrosoftGraphUnifiedRoleAssignment) GetDirectoryScopeOk() (*AnyOfmicrosoftGraphDirectoryObject, bool)`

GetDirectoryScopeOk returns a tuple with the DirectoryScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectoryScope

`func (o *MicrosoftGraphUnifiedRoleAssignment) SetDirectoryScope(v AnyOfmicrosoftGraphDirectoryObject)`

SetDirectoryScope sets DirectoryScope field to given value.

### HasDirectoryScope

`func (o *MicrosoftGraphUnifiedRoleAssignment) HasDirectoryScope() bool`

HasDirectoryScope returns a boolean if a field has been set.

### SetDirectoryScopeNil

`func (o *MicrosoftGraphUnifiedRoleAssignment) SetDirectoryScopeNil(b bool)`

 SetDirectoryScopeNil sets the value for DirectoryScope to be an explicit nil

### UnsetDirectoryScope
`func (o *MicrosoftGraphUnifiedRoleAssignment) UnsetDirectoryScope()`

UnsetDirectoryScope ensures that no value is present for DirectoryScope, not even an explicit nil
### GetPrincipal

`func (o *MicrosoftGraphUnifiedRoleAssignment) GetPrincipal() AnyOfmicrosoftGraphDirectoryObject`

GetPrincipal returns the Principal field if non-nil, zero value otherwise.

### GetPrincipalOk

`func (o *MicrosoftGraphUnifiedRoleAssignment) GetPrincipalOk() (*AnyOfmicrosoftGraphDirectoryObject, bool)`

GetPrincipalOk returns a tuple with the Principal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipal

`func (o *MicrosoftGraphUnifiedRoleAssignment) SetPrincipal(v AnyOfmicrosoftGraphDirectoryObject)`

SetPrincipal sets Principal field to given value.

### HasPrincipal

`func (o *MicrosoftGraphUnifiedRoleAssignment) HasPrincipal() bool`

HasPrincipal returns a boolean if a field has been set.

### SetPrincipalNil

`func (o *MicrosoftGraphUnifiedRoleAssignment) SetPrincipalNil(b bool)`

 SetPrincipalNil sets the value for Principal to be an explicit nil

### UnsetPrincipal
`func (o *MicrosoftGraphUnifiedRoleAssignment) UnsetPrincipal()`

UnsetPrincipal ensures that no value is present for Principal, not even an explicit nil
### GetRoleDefinition

`func (o *MicrosoftGraphUnifiedRoleAssignment) GetRoleDefinition() AnyOfmicrosoftGraphUnifiedRoleDefinition`

GetRoleDefinition returns the RoleDefinition field if non-nil, zero value otherwise.

### GetRoleDefinitionOk

`func (o *MicrosoftGraphUnifiedRoleAssignment) GetRoleDefinitionOk() (*AnyOfmicrosoftGraphUnifiedRoleDefinition, bool)`

GetRoleDefinitionOk returns a tuple with the RoleDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleDefinition

`func (o *MicrosoftGraphUnifiedRoleAssignment) SetRoleDefinition(v AnyOfmicrosoftGraphUnifiedRoleDefinition)`

SetRoleDefinition sets RoleDefinition field to given value.

### HasRoleDefinition

`func (o *MicrosoftGraphUnifiedRoleAssignment) HasRoleDefinition() bool`

HasRoleDefinition returns a boolean if a field has been set.

### SetRoleDefinitionNil

`func (o *MicrosoftGraphUnifiedRoleAssignment) SetRoleDefinitionNil(b bool)`

 SetRoleDefinitionNil sets the value for RoleDefinition to be an explicit nil

### UnsetRoleDefinition
`func (o *MicrosoftGraphUnifiedRoleAssignment) UnsetRoleDefinition()`

UnsetRoleDefinition ensures that no value is present for RoleDefinition, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


