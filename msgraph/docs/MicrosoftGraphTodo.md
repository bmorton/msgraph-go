# MicrosoftGraphTodo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**Lists** | Pointer to [**[]MicrosoftGraphTodoTaskList**](MicrosoftGraphTodoTaskList.md) | The task lists in the users mailbox. | [optional] 

## Methods

### NewMicrosoftGraphTodo

`func NewMicrosoftGraphTodo() *MicrosoftGraphTodo`

NewMicrosoftGraphTodo instantiates a new MicrosoftGraphTodo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphTodoWithDefaults

`func NewMicrosoftGraphTodoWithDefaults() *MicrosoftGraphTodo`

NewMicrosoftGraphTodoWithDefaults instantiates a new MicrosoftGraphTodo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphTodo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphTodo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphTodo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphTodo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLists

`func (o *MicrosoftGraphTodo) GetLists() []MicrosoftGraphTodoTaskList`

GetLists returns the Lists field if non-nil, zero value otherwise.

### GetListsOk

`func (o *MicrosoftGraphTodo) GetListsOk() (*[]MicrosoftGraphTodoTaskList, bool)`

GetListsOk returns a tuple with the Lists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLists

`func (o *MicrosoftGraphTodo) SetLists(v []MicrosoftGraphTodoTaskList)`

SetLists sets Lists field to given value.

### HasLists

`func (o *MicrosoftGraphTodo) HasLists() bool`

HasLists returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


