# TodoTaskList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **NullableString** | The name of the task list. | [optional] 
**IsOwner** | Pointer to **bool** | True if the user is owner of the given task list. | [optional] 
**IsShared** | Pointer to **bool** | True if the task list is shared with other users | [optional] 
**WellknownListName** | Pointer to [**AnyOfmicrosoftGraphWellknownListName**](anyOf&lt;microsoft.graph.wellknownListName&gt;.md) | Property indicating the list name if the given list is a well-known list. Possible values are: none, defaultList, flaggedEmails, unknownFutureValue. | [optional] 
**Extensions** | Pointer to [**[]MicrosoftGraphExtension**](MicrosoftGraphExtension.md) | The collection of open extensions defined for the task list. Nullable. | [optional] 
**Tasks** | Pointer to [**[]MicrosoftGraphTodoTask**](MicrosoftGraphTodoTask.md) | The tasks in this task list. Read-only. Nullable. | [optional] 

## Methods

### NewTodoTaskList

`func NewTodoTaskList() *TodoTaskList`

NewTodoTaskList instantiates a new TodoTaskList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTodoTaskListWithDefaults

`func NewTodoTaskListWithDefaults() *TodoTaskList`

NewTodoTaskListWithDefaults instantiates a new TodoTaskList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *TodoTaskList) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *TodoTaskList) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *TodoTaskList) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *TodoTaskList) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *TodoTaskList) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *TodoTaskList) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetIsOwner

`func (o *TodoTaskList) GetIsOwner() bool`

GetIsOwner returns the IsOwner field if non-nil, zero value otherwise.

### GetIsOwnerOk

`func (o *TodoTaskList) GetIsOwnerOk() (*bool, bool)`

GetIsOwnerOk returns a tuple with the IsOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOwner

`func (o *TodoTaskList) SetIsOwner(v bool)`

SetIsOwner sets IsOwner field to given value.

### HasIsOwner

`func (o *TodoTaskList) HasIsOwner() bool`

HasIsOwner returns a boolean if a field has been set.

### GetIsShared

`func (o *TodoTaskList) GetIsShared() bool`

GetIsShared returns the IsShared field if non-nil, zero value otherwise.

### GetIsSharedOk

`func (o *TodoTaskList) GetIsSharedOk() (*bool, bool)`

GetIsSharedOk returns a tuple with the IsShared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsShared

`func (o *TodoTaskList) SetIsShared(v bool)`

SetIsShared sets IsShared field to given value.

### HasIsShared

`func (o *TodoTaskList) HasIsShared() bool`

HasIsShared returns a boolean if a field has been set.

### GetWellknownListName

`func (o *TodoTaskList) GetWellknownListName() AnyOfmicrosoftGraphWellknownListName`

GetWellknownListName returns the WellknownListName field if non-nil, zero value otherwise.

### GetWellknownListNameOk

`func (o *TodoTaskList) GetWellknownListNameOk() (*AnyOfmicrosoftGraphWellknownListName, bool)`

GetWellknownListNameOk returns a tuple with the WellknownListName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWellknownListName

`func (o *TodoTaskList) SetWellknownListName(v AnyOfmicrosoftGraphWellknownListName)`

SetWellknownListName sets WellknownListName field to given value.

### HasWellknownListName

`func (o *TodoTaskList) HasWellknownListName() bool`

HasWellknownListName returns a boolean if a field has been set.

### SetWellknownListNameNil

`func (o *TodoTaskList) SetWellknownListNameNil(b bool)`

 SetWellknownListNameNil sets the value for WellknownListName to be an explicit nil

### UnsetWellknownListName
`func (o *TodoTaskList) UnsetWellknownListName()`

UnsetWellknownListName ensures that no value is present for WellknownListName, not even an explicit nil
### GetExtensions

`func (o *TodoTaskList) GetExtensions() []MicrosoftGraphExtension`

GetExtensions returns the Extensions field if non-nil, zero value otherwise.

### GetExtensionsOk

`func (o *TodoTaskList) GetExtensionsOk() (*[]MicrosoftGraphExtension, bool)`

GetExtensionsOk returns a tuple with the Extensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensions

`func (o *TodoTaskList) SetExtensions(v []MicrosoftGraphExtension)`

SetExtensions sets Extensions field to given value.

### HasExtensions

`func (o *TodoTaskList) HasExtensions() bool`

HasExtensions returns a boolean if a field has been set.

### GetTasks

`func (o *TodoTaskList) GetTasks() []MicrosoftGraphTodoTask`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *TodoTaskList) GetTasksOk() (*[]MicrosoftGraphTodoTask, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *TodoTaskList) SetTasks(v []MicrosoftGraphTodoTask)`

SetTasks sets Tasks field to given value.

### HasTasks

`func (o *TodoTaskList) HasTasks() bool`

HasTasks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


