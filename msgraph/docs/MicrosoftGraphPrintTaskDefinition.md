# MicrosoftGraphPrintTaskDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**CreatedBy** | Pointer to [**MicrosoftGraphAppIdentity**](MicrosoftGraphAppIdentity.md) |  | [optional] 
**DisplayName** | Pointer to **string** | The name of the printTaskDefinition. | [optional] 
**Tasks** | Pointer to [**[]MicrosoftGraphPrintTask**](MicrosoftGraphPrintTask.md) | A list of tasks that have been created based on this definition. The list includes currently running tasks and recently completed tasks. Read-only. | [optional] 

## Methods

### NewMicrosoftGraphPrintTaskDefinition

`func NewMicrosoftGraphPrintTaskDefinition() *MicrosoftGraphPrintTaskDefinition`

NewMicrosoftGraphPrintTaskDefinition instantiates a new MicrosoftGraphPrintTaskDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphPrintTaskDefinitionWithDefaults

`func NewMicrosoftGraphPrintTaskDefinitionWithDefaults() *MicrosoftGraphPrintTaskDefinition`

NewMicrosoftGraphPrintTaskDefinitionWithDefaults instantiates a new MicrosoftGraphPrintTaskDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphPrintTaskDefinition) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphPrintTaskDefinition) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphPrintTaskDefinition) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphPrintTaskDefinition) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedBy

`func (o *MicrosoftGraphPrintTaskDefinition) GetCreatedBy() MicrosoftGraphAppIdentity`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *MicrosoftGraphPrintTaskDefinition) GetCreatedByOk() (*MicrosoftGraphAppIdentity, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *MicrosoftGraphPrintTaskDefinition) SetCreatedBy(v MicrosoftGraphAppIdentity)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *MicrosoftGraphPrintTaskDefinition) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetDisplayName

`func (o *MicrosoftGraphPrintTaskDefinition) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphPrintTaskDefinition) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *MicrosoftGraphPrintTaskDefinition) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *MicrosoftGraphPrintTaskDefinition) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetTasks

`func (o *MicrosoftGraphPrintTaskDefinition) GetTasks() []MicrosoftGraphPrintTask`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *MicrosoftGraphPrintTaskDefinition) GetTasksOk() (*[]MicrosoftGraphPrintTask, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *MicrosoftGraphPrintTaskDefinition) SetTasks(v []MicrosoftGraphPrintTask)`

SetTasks sets Tasks field to given value.

### HasTasks

`func (o *MicrosoftGraphPrintTaskDefinition) HasTasks() bool`

HasTasks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


