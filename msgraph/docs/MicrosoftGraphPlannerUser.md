# MicrosoftGraphPlannerUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**Plans** | Pointer to [**[]MicrosoftGraphPlannerPlan**](MicrosoftGraphPlannerPlan.md) | Read-only. Nullable. Returns the plannerTasks assigned to the user. | [optional] 
**Tasks** | Pointer to [**[]MicrosoftGraphPlannerTask**](MicrosoftGraphPlannerTask.md) | Read-only. Nullable. Returns the plannerPlans shared with the user. | [optional] 

## Methods

### NewMicrosoftGraphPlannerUser

`func NewMicrosoftGraphPlannerUser() *MicrosoftGraphPlannerUser`

NewMicrosoftGraphPlannerUser instantiates a new MicrosoftGraphPlannerUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphPlannerUserWithDefaults

`func NewMicrosoftGraphPlannerUserWithDefaults() *MicrosoftGraphPlannerUser`

NewMicrosoftGraphPlannerUserWithDefaults instantiates a new MicrosoftGraphPlannerUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphPlannerUser) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphPlannerUser) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphPlannerUser) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphPlannerUser) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPlans

`func (o *MicrosoftGraphPlannerUser) GetPlans() []MicrosoftGraphPlannerPlan`

GetPlans returns the Plans field if non-nil, zero value otherwise.

### GetPlansOk

`func (o *MicrosoftGraphPlannerUser) GetPlansOk() (*[]MicrosoftGraphPlannerPlan, bool)`

GetPlansOk returns a tuple with the Plans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlans

`func (o *MicrosoftGraphPlannerUser) SetPlans(v []MicrosoftGraphPlannerPlan)`

SetPlans sets Plans field to given value.

### HasPlans

`func (o *MicrosoftGraphPlannerUser) HasPlans() bool`

HasPlans returns a boolean if a field has been set.

### GetTasks

`func (o *MicrosoftGraphPlannerUser) GetTasks() []MicrosoftGraphPlannerTask`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *MicrosoftGraphPlannerUser) GetTasksOk() (*[]MicrosoftGraphPlannerTask, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *MicrosoftGraphPlannerUser) SetTasks(v []MicrosoftGraphPlannerTask)`

SetTasks sets Tasks field to given value.

### HasTasks

`func (o *MicrosoftGraphPlannerUser) HasTasks() bool`

HasTasks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


