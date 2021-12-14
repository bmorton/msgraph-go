# UnifiedRoleAssignment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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

### NewUnifiedRoleAssignment

`func NewUnifiedRoleAssignment() *UnifiedRoleAssignment`

NewUnifiedRoleAssignment instantiates a new UnifiedRoleAssignment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnifiedRoleAssignmentWithDefaults

`func NewUnifiedRoleAssignmentWithDefaults() *UnifiedRoleAssignment`

NewUnifiedRoleAssignmentWithDefaults instantiates a new UnifiedRoleAssignment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppScopeId

`func (o *UnifiedRoleAssignment) GetAppScopeId() string`

GetAppScopeId returns the AppScopeId field if non-nil, zero value otherwise.

### GetAppScopeIdOk

`func (o *UnifiedRoleAssignment) GetAppScopeIdOk() (*string, bool)`

GetAppScopeIdOk returns a tuple with the AppScopeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppScopeId

`func (o *UnifiedRoleAssignment) SetAppScopeId(v string)`

SetAppScopeId sets AppScopeId field to given value.

### HasAppScopeId

`func (o *UnifiedRoleAssignment) HasAppScopeId() bool`

HasAppScopeId returns a boolean if a field has been set.

### SetAppScopeIdNil

`func (o *UnifiedRoleAssignment) SetAppScopeIdNil(b bool)`

 SetAppScopeIdNil sets the value for AppScopeId to be an explicit nil

### UnsetAppScopeId
`func (o *UnifiedRoleAssignment) UnsetAppScopeId()`

UnsetAppScopeId ensures that no value is present for AppScopeId, not even an explicit nil
### GetCondition

`func (o *UnifiedRoleAssignment) GetCondition() string`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *UnifiedRoleAssignment) GetConditionOk() (*string, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *UnifiedRoleAssignment) SetCondition(v string)`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *UnifiedRoleAssignment) HasCondition() bool`

HasCondition returns a boolean if a field has been set.

### SetConditionNil

`func (o *UnifiedRoleAssignment) SetConditionNil(b bool)`

 SetConditionNil sets the value for Condition to be an explicit nil

### UnsetCondition
`func (o *UnifiedRoleAssignment) UnsetCondition()`

UnsetCondition ensures that no value is present for Condition, not even an explicit nil
### GetDirectoryScopeId

`func (o *UnifiedRoleAssignment) GetDirectoryScopeId() string`

GetDirectoryScopeId returns the DirectoryScopeId field if non-nil, zero value otherwise.

### GetDirectoryScopeIdOk

`func (o *UnifiedRoleAssignment) GetDirectoryScopeIdOk() (*string, bool)`

GetDirectoryScopeIdOk returns a tuple with the DirectoryScopeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectoryScopeId

`func (o *UnifiedRoleAssignment) SetDirectoryScopeId(v string)`

SetDirectoryScopeId sets DirectoryScopeId field to given value.

### HasDirectoryScopeId

`func (o *UnifiedRoleAssignment) HasDirectoryScopeId() bool`

HasDirectoryScopeId returns a boolean if a field has been set.

### SetDirectoryScopeIdNil

`func (o *UnifiedRoleAssignment) SetDirectoryScopeIdNil(b bool)`

 SetDirectoryScopeIdNil sets the value for DirectoryScopeId to be an explicit nil

### UnsetDirectoryScopeId
`func (o *UnifiedRoleAssignment) UnsetDirectoryScopeId()`

UnsetDirectoryScopeId ensures that no value is present for DirectoryScopeId, not even an explicit nil
### GetPrincipalId

`func (o *UnifiedRoleAssignment) GetPrincipalId() string`

GetPrincipalId returns the PrincipalId field if non-nil, zero value otherwise.

### GetPrincipalIdOk

`func (o *UnifiedRoleAssignment) GetPrincipalIdOk() (*string, bool)`

GetPrincipalIdOk returns a tuple with the PrincipalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipalId

`func (o *UnifiedRoleAssignment) SetPrincipalId(v string)`

SetPrincipalId sets PrincipalId field to given value.

### HasPrincipalId

`func (o *UnifiedRoleAssignment) HasPrincipalId() bool`

HasPrincipalId returns a boolean if a field has been set.

### SetPrincipalIdNil

`func (o *UnifiedRoleAssignment) SetPrincipalIdNil(b bool)`

 SetPrincipalIdNil sets the value for PrincipalId to be an explicit nil

### UnsetPrincipalId
`func (o *UnifiedRoleAssignment) UnsetPrincipalId()`

UnsetPrincipalId ensures that no value is present for PrincipalId, not even an explicit nil
### GetRoleDefinitionId

`func (o *UnifiedRoleAssignment) GetRoleDefinitionId() string`

GetRoleDefinitionId returns the RoleDefinitionId field if non-nil, zero value otherwise.

### GetRoleDefinitionIdOk

`func (o *UnifiedRoleAssignment) GetRoleDefinitionIdOk() (*string, bool)`

GetRoleDefinitionIdOk returns a tuple with the RoleDefinitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleDefinitionId

`func (o *UnifiedRoleAssignment) SetRoleDefinitionId(v string)`

SetRoleDefinitionId sets RoleDefinitionId field to given value.

### HasRoleDefinitionId

`func (o *UnifiedRoleAssignment) HasRoleDefinitionId() bool`

HasRoleDefinitionId returns a boolean if a field has been set.

### SetRoleDefinitionIdNil

`func (o *UnifiedRoleAssignment) SetRoleDefinitionIdNil(b bool)`

 SetRoleDefinitionIdNil sets the value for RoleDefinitionId to be an explicit nil

### UnsetRoleDefinitionId
`func (o *UnifiedRoleAssignment) UnsetRoleDefinitionId()`

UnsetRoleDefinitionId ensures that no value is present for RoleDefinitionId, not even an explicit nil
### GetAppScope

`func (o *UnifiedRoleAssignment) GetAppScope() AnyOfmicrosoftGraphAppScope`

GetAppScope returns the AppScope field if non-nil, zero value otherwise.

### GetAppScopeOk

`func (o *UnifiedRoleAssignment) GetAppScopeOk() (*AnyOfmicrosoftGraphAppScope, bool)`

GetAppScopeOk returns a tuple with the AppScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppScope

`func (o *UnifiedRoleAssignment) SetAppScope(v AnyOfmicrosoftGraphAppScope)`

SetAppScope sets AppScope field to given value.

### HasAppScope

`func (o *UnifiedRoleAssignment) HasAppScope() bool`

HasAppScope returns a boolean if a field has been set.

### SetAppScopeNil

`func (o *UnifiedRoleAssignment) SetAppScopeNil(b bool)`

 SetAppScopeNil sets the value for AppScope to be an explicit nil

### UnsetAppScope
`func (o *UnifiedRoleAssignment) UnsetAppScope()`

UnsetAppScope ensures that no value is present for AppScope, not even an explicit nil
### GetDirectoryScope

`func (o *UnifiedRoleAssignment) GetDirectoryScope() AnyOfmicrosoftGraphDirectoryObject`

GetDirectoryScope returns the DirectoryScope field if non-nil, zero value otherwise.

### GetDirectoryScopeOk

`func (o *UnifiedRoleAssignment) GetDirectoryScopeOk() (*AnyOfmicrosoftGraphDirectoryObject, bool)`

GetDirectoryScopeOk returns a tuple with the DirectoryScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectoryScope

`func (o *UnifiedRoleAssignment) SetDirectoryScope(v AnyOfmicrosoftGraphDirectoryObject)`

SetDirectoryScope sets DirectoryScope field to given value.

### HasDirectoryScope

`func (o *UnifiedRoleAssignment) HasDirectoryScope() bool`

HasDirectoryScope returns a boolean if a field has been set.

### SetDirectoryScopeNil

`func (o *UnifiedRoleAssignment) SetDirectoryScopeNil(b bool)`

 SetDirectoryScopeNil sets the value for DirectoryScope to be an explicit nil

### UnsetDirectoryScope
`func (o *UnifiedRoleAssignment) UnsetDirectoryScope()`

UnsetDirectoryScope ensures that no value is present for DirectoryScope, not even an explicit nil
### GetPrincipal

`func (o *UnifiedRoleAssignment) GetPrincipal() AnyOfmicrosoftGraphDirectoryObject`

GetPrincipal returns the Principal field if non-nil, zero value otherwise.

### GetPrincipalOk

`func (o *UnifiedRoleAssignment) GetPrincipalOk() (*AnyOfmicrosoftGraphDirectoryObject, bool)`

GetPrincipalOk returns a tuple with the Principal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipal

`func (o *UnifiedRoleAssignment) SetPrincipal(v AnyOfmicrosoftGraphDirectoryObject)`

SetPrincipal sets Principal field to given value.

### HasPrincipal

`func (o *UnifiedRoleAssignment) HasPrincipal() bool`

HasPrincipal returns a boolean if a field has been set.

### SetPrincipalNil

`func (o *UnifiedRoleAssignment) SetPrincipalNil(b bool)`

 SetPrincipalNil sets the value for Principal to be an explicit nil

### UnsetPrincipal
`func (o *UnifiedRoleAssignment) UnsetPrincipal()`

UnsetPrincipal ensures that no value is present for Principal, not even an explicit nil
### GetRoleDefinition

`func (o *UnifiedRoleAssignment) GetRoleDefinition() AnyOfmicrosoftGraphUnifiedRoleDefinition`

GetRoleDefinition returns the RoleDefinition field if non-nil, zero value otherwise.

### GetRoleDefinitionOk

`func (o *UnifiedRoleAssignment) GetRoleDefinitionOk() (*AnyOfmicrosoftGraphUnifiedRoleDefinition, bool)`

GetRoleDefinitionOk returns a tuple with the RoleDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleDefinition

`func (o *UnifiedRoleAssignment) SetRoleDefinition(v AnyOfmicrosoftGraphUnifiedRoleDefinition)`

SetRoleDefinition sets RoleDefinition field to given value.

### HasRoleDefinition

`func (o *UnifiedRoleAssignment) HasRoleDefinition() bool`

HasRoleDefinition returns a boolean if a field has been set.

### SetRoleDefinitionNil

`func (o *UnifiedRoleAssignment) SetRoleDefinitionNil(b bool)`

 SetRoleDefinitionNil sets the value for RoleDefinition to be an explicit nil

### UnsetRoleDefinition
`func (o *UnifiedRoleAssignment) UnsetRoleDefinition()`

UnsetRoleDefinition ensures that no value is present for RoleDefinition, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


