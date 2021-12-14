# Todo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Lists** | Pointer to [**[]MicrosoftGraphTodoTaskList**](MicrosoftGraphTodoTaskList.md) | The task lists in the users mailbox. | [optional] 

## Methods

### NewTodo

`func NewTodo() *Todo`

NewTodo instantiates a new Todo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTodoWithDefaults

`func NewTodoWithDefaults() *Todo`

NewTodoWithDefaults instantiates a new Todo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLists

`func (o *Todo) GetLists() []MicrosoftGraphTodoTaskList`

GetLists returns the Lists field if non-nil, zero value otherwise.

### GetListsOk

`func (o *Todo) GetListsOk() (*[]MicrosoftGraphTodoTaskList, bool)`

GetListsOk returns a tuple with the Lists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLists

`func (o *Todo) SetLists(v []MicrosoftGraphTodoTaskList)`

SetLists sets Lists field to given value.

### HasLists

`func (o *Todo) HasLists() bool`

HasLists returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


