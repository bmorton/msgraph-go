# MicrosoftGraphTodoTaskList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**DisplayName** | Pointer to **NullableString** | The name of the task list. | [optional] 
**IsOwner** | Pointer to **bool** | True if the user is owner of the given task list. | [optional] 
**IsShared** | Pointer to **bool** | True if the task list is shared with other users | [optional] 
**WellknownListName** | Pointer to [**AnyOfmicrosoftGraphWellknownListName**](anyOf&lt;microsoft.graph.wellknownListName&gt;.md) | Property indicating the list name if the given list is a well-known list. Possible values are: none, defaultList, flaggedEmails, unknownFutureValue. | [optional] 
**Extensions** | Pointer to [**[]MicrosoftGraphExtension**](MicrosoftGraphExtension.md) | The collection of open extensions defined for the task list. Nullable. | [optional] 
**Tasks** | Pointer to [**[]MicrosoftGraphTodoTask**](MicrosoftGraphTodoTask.md) | The tasks in this task list. Read-only. Nullable. | [optional] 

## Methods

### NewMicrosoftGraphTodoTaskList

`func NewMicrosoftGraphTodoTaskList() *MicrosoftGraphTodoTaskList`

NewMicrosoftGraphTodoTaskList instantiates a new MicrosoftGraphTodoTaskList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphTodoTaskListWithDefaults

`func NewMicrosoftGraphTodoTaskListWithDefaults() *MicrosoftGraphTodoTaskList`

NewMicrosoftGraphTodoTaskListWithDefaults instantiates a new MicrosoftGraphTodoTaskList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphTodoTaskList) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphTodoTaskList) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphTodoTaskList) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphTodoTaskList) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDisplayName

`func (o *MicrosoftGraphTodoTaskList) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphTodoTaskList) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *MicrosoftGraphTodoTaskList) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *MicrosoftGraphTodoTaskList) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *MicrosoftGraphTodoTaskList) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *MicrosoftGraphTodoTaskList) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetIsOwner

`func (o *MicrosoftGraphTodoTaskList) GetIsOwner() bool`

GetIsOwner returns the IsOwner field if non-nil, zero value otherwise.

### GetIsOwnerOk

`func (o *MicrosoftGraphTodoTaskList) GetIsOwnerOk() (*bool, bool)`

GetIsOwnerOk returns a tuple with the IsOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOwner

`func (o *MicrosoftGraphTodoTaskList) SetIsOwner(v bool)`

SetIsOwner sets IsOwner field to given value.

### HasIsOwner

`func (o *MicrosoftGraphTodoTaskList) HasIsOwner() bool`

HasIsOwner returns a boolean if a field has been set.

### GetIsShared

`func (o *MicrosoftGraphTodoTaskList) GetIsShared() bool`

GetIsShared returns the IsShared field if non-nil, zero value otherwise.

### GetIsSharedOk

`func (o *MicrosoftGraphTodoTaskList) GetIsSharedOk() (*bool, bool)`

GetIsSharedOk returns a tuple with the IsShared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsShared

`func (o *MicrosoftGraphTodoTaskList) SetIsShared(v bool)`

SetIsShared sets IsShared field to given value.

### HasIsShared

`func (o *MicrosoftGraphTodoTaskList) HasIsShared() bool`

HasIsShared returns a boolean if a field has been set.

### GetWellknownListName

`func (o *MicrosoftGraphTodoTaskList) GetWellknownListName() AnyOfmicrosoftGraphWellknownListName`

GetWellknownListName returns the WellknownListName field if non-nil, zero value otherwise.

### GetWellknownListNameOk

`func (o *MicrosoftGraphTodoTaskList) GetWellknownListNameOk() (*AnyOfmicrosoftGraphWellknownListName, bool)`

GetWellknownListNameOk returns a tuple with the WellknownListName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWellknownListName

`func (o *MicrosoftGraphTodoTaskList) SetWellknownListName(v AnyOfmicrosoftGraphWellknownListName)`

SetWellknownListName sets WellknownListName field to given value.

### HasWellknownListName

`func (o *MicrosoftGraphTodoTaskList) HasWellknownListName() bool`

HasWellknownListName returns a boolean if a field has been set.

### SetWellknownListNameNil

`func (o *MicrosoftGraphTodoTaskList) SetWellknownListNameNil(b bool)`

 SetWellknownListNameNil sets the value for WellknownListName to be an explicit nil

### UnsetWellknownListName
`func (o *MicrosoftGraphTodoTaskList) UnsetWellknownListName()`

UnsetWellknownListName ensures that no value is present for WellknownListName, not even an explicit nil
### GetExtensions

`func (o *MicrosoftGraphTodoTaskList) GetExtensions() []MicrosoftGraphExtension`

GetExtensions returns the Extensions field if non-nil, zero value otherwise.

### GetExtensionsOk

`func (o *MicrosoftGraphTodoTaskList) GetExtensionsOk() (*[]MicrosoftGraphExtension, bool)`

GetExtensionsOk returns a tuple with the Extensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensions

`func (o *MicrosoftGraphTodoTaskList) SetExtensions(v []MicrosoftGraphExtension)`

SetExtensions sets Extensions field to given value.

### HasExtensions

`func (o *MicrosoftGraphTodoTaskList) HasExtensions() bool`

HasExtensions returns a boolean if a field has been set.

### GetTasks

`func (o *MicrosoftGraphTodoTaskList) GetTasks() []MicrosoftGraphTodoTask`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *MicrosoftGraphTodoTaskList) GetTasksOk() (*[]MicrosoftGraphTodoTask, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *MicrosoftGraphTodoTaskList) SetTasks(v []MicrosoftGraphTodoTask)`

SetTasks sets Tasks field to given value.

### HasTasks

`func (o *MicrosoftGraphTodoTaskList) HasTasks() bool`

HasTasks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


