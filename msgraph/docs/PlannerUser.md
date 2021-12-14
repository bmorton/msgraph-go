# PlannerUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Plans** | Pointer to [**[]MicrosoftGraphPlannerPlan**](MicrosoftGraphPlannerPlan.md) | Read-only. Nullable. Returns the plannerTasks assigned to the user. | [optional] 
**Tasks** | Pointer to [**[]MicrosoftGraphPlannerTask**](MicrosoftGraphPlannerTask.md) | Read-only. Nullable. Returns the plannerPlans shared with the user. | [optional] 

## Methods

### NewPlannerUser

`func NewPlannerUser() *PlannerUser`

NewPlannerUser instantiates a new PlannerUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlannerUserWithDefaults

`func NewPlannerUserWithDefaults() *PlannerUser`

NewPlannerUserWithDefaults instantiates a new PlannerUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlans

`func (o *PlannerUser) GetPlans() []MicrosoftGraphPlannerPlan`

GetPlans returns the Plans field if non-nil, zero value otherwise.

### GetPlansOk

`func (o *PlannerUser) GetPlansOk() (*[]MicrosoftGraphPlannerPlan, bool)`

GetPlansOk returns a tuple with the Plans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlans

`func (o *PlannerUser) SetPlans(v []MicrosoftGraphPlannerPlan)`

SetPlans sets Plans field to given value.

### HasPlans

`func (o *PlannerUser) HasPlans() bool`

HasPlans returns a boolean if a field has been set.

### GetTasks

`func (o *PlannerUser) GetTasks() []MicrosoftGraphPlannerTask`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *PlannerUser) GetTasksOk() (*[]MicrosoftGraphPlannerTask, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *PlannerUser) SetTasks(v []MicrosoftGraphPlannerTask)`

SetTasks sets Tasks field to given value.

### HasTasks

`func (o *PlannerUser) HasTasks() bool`

HasTasks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


