# MicrosoftGraphAppRoleAssignment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**DeletedDateTime** | Pointer to **NullableTime** |  | [optional] 
**AppRoleId** | Pointer to **string** | The identifier (id) for the app role which is assigned to the principal. This app role must be exposed in the appRoles property on the resource application&#39;s service principal (resourceId). If the resource application has not declared any app roles, a default app role ID of 00000000-0000-0000-0000-000000000000 can be specified to signal that the principal is assigned to the resource app without any specific app roles. Required on create. | [optional] 
**CreatedDateTime** | Pointer to **NullableTime** | The time when the app role assignment was created.The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only. | [optional] 
**PrincipalDisplayName** | Pointer to **NullableString** | The display name of the user, group, or service principal that was granted the app role assignment. Read-only. Supports $filter (eq and startswith). | [optional] 
**PrincipalId** | Pointer to **NullableString** | The unique identifier (id) for the user, group or service principal being granted the app role. Required on create. | [optional] 
**PrincipalType** | Pointer to **NullableString** | The type of the assigned principal. This can either be User, Group or ServicePrincipal. Read-only. | [optional] 
**ResourceDisplayName** | Pointer to **NullableString** | The display name of the resource app&#39;s service principal to which the assignment is made. | [optional] 
**ResourceId** | Pointer to **NullableString** | The unique identifier (id) for the resource service principal for which the assignment is made. Required on create. Supports $filter (eq only). | [optional] 

## Methods

### NewMicrosoftGraphAppRoleAssignment

`func NewMicrosoftGraphAppRoleAssignment() *MicrosoftGraphAppRoleAssignment`

NewMicrosoftGraphAppRoleAssignment instantiates a new MicrosoftGraphAppRoleAssignment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphAppRoleAssignmentWithDefaults

`func NewMicrosoftGraphAppRoleAssignmentWithDefaults() *MicrosoftGraphAppRoleAssignment`

NewMicrosoftGraphAppRoleAssignmentWithDefaults instantiates a new MicrosoftGraphAppRoleAssignment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphAppRoleAssignment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphAppRoleAssignment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphAppRoleAssignment) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphAppRoleAssignment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDeletedDateTime

`func (o *MicrosoftGraphAppRoleAssignment) GetDeletedDateTime() time.Time`

GetDeletedDateTime returns the DeletedDateTime field if non-nil, zero value otherwise.

### GetDeletedDateTimeOk

`func (o *MicrosoftGraphAppRoleAssignment) GetDeletedDateTimeOk() (*time.Time, bool)`

GetDeletedDateTimeOk returns a tuple with the DeletedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedDateTime

`func (o *MicrosoftGraphAppRoleAssignment) SetDeletedDateTime(v time.Time)`

SetDeletedDateTime sets DeletedDateTime field to given value.

### HasDeletedDateTime

`func (o *MicrosoftGraphAppRoleAssignment) HasDeletedDateTime() bool`

HasDeletedDateTime returns a boolean if a field has been set.

### SetDeletedDateTimeNil

`func (o *MicrosoftGraphAppRoleAssignment) SetDeletedDateTimeNil(b bool)`

 SetDeletedDateTimeNil sets the value for DeletedDateTime to be an explicit nil

### UnsetDeletedDateTime
`func (o *MicrosoftGraphAppRoleAssignment) UnsetDeletedDateTime()`

UnsetDeletedDateTime ensures that no value is present for DeletedDateTime, not even an explicit nil
### GetAppRoleId

`func (o *MicrosoftGraphAppRoleAssignment) GetAppRoleId() string`

GetAppRoleId returns the AppRoleId field if non-nil, zero value otherwise.

### GetAppRoleIdOk

`func (o *MicrosoftGraphAppRoleAssignment) GetAppRoleIdOk() (*string, bool)`

GetAppRoleIdOk returns a tuple with the AppRoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppRoleId

`func (o *MicrosoftGraphAppRoleAssignment) SetAppRoleId(v string)`

SetAppRoleId sets AppRoleId field to given value.

### HasAppRoleId

`func (o *MicrosoftGraphAppRoleAssignment) HasAppRoleId() bool`

HasAppRoleId returns a boolean if a field has been set.

### GetCreatedDateTime

`func (o *MicrosoftGraphAppRoleAssignment) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *MicrosoftGraphAppRoleAssignment) GetCreatedDateTimeOk() (*time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateTime

`func (o *MicrosoftGraphAppRoleAssignment) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime sets CreatedDateTime field to given value.

### HasCreatedDateTime

`func (o *MicrosoftGraphAppRoleAssignment) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTimeNil

`func (o *MicrosoftGraphAppRoleAssignment) SetCreatedDateTimeNil(b bool)`

 SetCreatedDateTimeNil sets the value for CreatedDateTime to be an explicit nil

### UnsetCreatedDateTime
`func (o *MicrosoftGraphAppRoleAssignment) UnsetCreatedDateTime()`

UnsetCreatedDateTime ensures that no value is present for CreatedDateTime, not even an explicit nil
### GetPrincipalDisplayName

`func (o *MicrosoftGraphAppRoleAssignment) GetPrincipalDisplayName() string`

GetPrincipalDisplayName returns the PrincipalDisplayName field if non-nil, zero value otherwise.

### GetPrincipalDisplayNameOk

`func (o *MicrosoftGraphAppRoleAssignment) GetPrincipalDisplayNameOk() (*string, bool)`

GetPrincipalDisplayNameOk returns a tuple with the PrincipalDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipalDisplayName

`func (o *MicrosoftGraphAppRoleAssignment) SetPrincipalDisplayName(v string)`

SetPrincipalDisplayName sets PrincipalDisplayName field to given value.

### HasPrincipalDisplayName

`func (o *MicrosoftGraphAppRoleAssignment) HasPrincipalDisplayName() bool`

HasPrincipalDisplayName returns a boolean if a field has been set.

### SetPrincipalDisplayNameNil

`func (o *MicrosoftGraphAppRoleAssignment) SetPrincipalDisplayNameNil(b bool)`

 SetPrincipalDisplayNameNil sets the value for PrincipalDisplayName to be an explicit nil

### UnsetPrincipalDisplayName
`func (o *MicrosoftGraphAppRoleAssignment) UnsetPrincipalDisplayName()`

UnsetPrincipalDisplayName ensures that no value is present for PrincipalDisplayName, not even an explicit nil
### GetPrincipalId

`func (o *MicrosoftGraphAppRoleAssignment) GetPrincipalId() string`

GetPrincipalId returns the PrincipalId field if non-nil, zero value otherwise.

### GetPrincipalIdOk

`func (o *MicrosoftGraphAppRoleAssignment) GetPrincipalIdOk() (*string, bool)`

GetPrincipalIdOk returns a tuple with the PrincipalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipalId

`func (o *MicrosoftGraphAppRoleAssignment) SetPrincipalId(v string)`

SetPrincipalId sets PrincipalId field to given value.

### HasPrincipalId

`func (o *MicrosoftGraphAppRoleAssignment) HasPrincipalId() bool`

HasPrincipalId returns a boolean if a field has been set.

### SetPrincipalIdNil

`func (o *MicrosoftGraphAppRoleAssignment) SetPrincipalIdNil(b bool)`

 SetPrincipalIdNil sets the value for PrincipalId to be an explicit nil

### UnsetPrincipalId
`func (o *MicrosoftGraphAppRoleAssignment) UnsetPrincipalId()`

UnsetPrincipalId ensures that no value is present for PrincipalId, not even an explicit nil
### GetPrincipalType

`func (o *MicrosoftGraphAppRoleAssignment) GetPrincipalType() string`

GetPrincipalType returns the PrincipalType field if non-nil, zero value otherwise.

### GetPrincipalTypeOk

`func (o *MicrosoftGraphAppRoleAssignment) GetPrincipalTypeOk() (*string, bool)`

GetPrincipalTypeOk returns a tuple with the PrincipalType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipalType

`func (o *MicrosoftGraphAppRoleAssignment) SetPrincipalType(v string)`

SetPrincipalType sets PrincipalType field to given value.

### HasPrincipalType

`func (o *MicrosoftGraphAppRoleAssignment) HasPrincipalType() bool`

HasPrincipalType returns a boolean if a field has been set.

### SetPrincipalTypeNil

`func (o *MicrosoftGraphAppRoleAssignment) SetPrincipalTypeNil(b bool)`

 SetPrincipalTypeNil sets the value for PrincipalType to be an explicit nil

### UnsetPrincipalType
`func (o *MicrosoftGraphAppRoleAssignment) UnsetPrincipalType()`

UnsetPrincipalType ensures that no value is present for PrincipalType, not even an explicit nil
### GetResourceDisplayName

`func (o *MicrosoftGraphAppRoleAssignment) GetResourceDisplayName() string`

GetResourceDisplayName returns the ResourceDisplayName field if non-nil, zero value otherwise.

### GetResourceDisplayNameOk

`func (o *MicrosoftGraphAppRoleAssignment) GetResourceDisplayNameOk() (*string, bool)`

GetResourceDisplayNameOk returns a tuple with the ResourceDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceDisplayName

`func (o *MicrosoftGraphAppRoleAssignment) SetResourceDisplayName(v string)`

SetResourceDisplayName sets ResourceDisplayName field to given value.

### HasResourceDisplayName

`func (o *MicrosoftGraphAppRoleAssignment) HasResourceDisplayName() bool`

HasResourceDisplayName returns a boolean if a field has been set.

### SetResourceDisplayNameNil

`func (o *MicrosoftGraphAppRoleAssignment) SetResourceDisplayNameNil(b bool)`

 SetResourceDisplayNameNil sets the value for ResourceDisplayName to be an explicit nil

### UnsetResourceDisplayName
`func (o *MicrosoftGraphAppRoleAssignment) UnsetResourceDisplayName()`

UnsetResourceDisplayName ensures that no value is present for ResourceDisplayName, not even an explicit nil
### GetResourceId

`func (o *MicrosoftGraphAppRoleAssignment) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *MicrosoftGraphAppRoleAssignment) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *MicrosoftGraphAppRoleAssignment) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.

### HasResourceId

`func (o *MicrosoftGraphAppRoleAssignment) HasResourceId() bool`

HasResourceId returns a boolean if a field has been set.

### SetResourceIdNil

`func (o *MicrosoftGraphAppRoleAssignment) SetResourceIdNil(b bool)`

 SetResourceIdNil sets the value for ResourceId to be an explicit nil

### UnsetResourceId
`func (o *MicrosoftGraphAppRoleAssignment) UnsetResourceId()`

UnsetResourceId ensures that no value is present for ResourceId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


