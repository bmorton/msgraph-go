# Planner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Buckets** | Pointer to [**[]MicrosoftGraphPlannerBucket**](MicrosoftGraphPlannerBucket.md) | Read-only. Nullable. Returns a collection of the specified buckets | [optional] 
**Plans** | Pointer to [**[]MicrosoftGraphPlannerPlan**](MicrosoftGraphPlannerPlan.md) | Read-only. Nullable. Returns a collection of the specified plans | [optional] 
**Tasks** | Pointer to [**[]MicrosoftGraphPlannerTask**](MicrosoftGraphPlannerTask.md) | Read-only. Nullable. Returns a collection of the specified tasks | [optional] 

## Methods

### NewPlanner

`func NewPlanner() *Planner`

NewPlanner instantiates a new Planner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlannerWithDefaults

`func NewPlannerWithDefaults() *Planner`

NewPlannerWithDefaults instantiates a new Planner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuckets

`func (o *Planner) GetBuckets() []MicrosoftGraphPlannerBucket`

GetBuckets returns the Buckets field if non-nil, zero value otherwise.

### GetBucketsOk

`func (o *Planner) GetBucketsOk() (*[]MicrosoftGraphPlannerBucket, bool)`

GetBucketsOk returns a tuple with the Buckets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuckets

`func (o *Planner) SetBuckets(v []MicrosoftGraphPlannerBucket)`

SetBuckets sets Buckets field to given value.

### HasBuckets

`func (o *Planner) HasBuckets() bool`

HasBuckets returns a boolean if a field has been set.

### GetPlans

`func (o *Planner) GetPlans() []MicrosoftGraphPlannerPlan`

GetPlans returns the Plans field if non-nil, zero value otherwise.

### GetPlansOk

`func (o *Planner) GetPlansOk() (*[]MicrosoftGraphPlannerPlan, bool)`

GetPlansOk returns a tuple with the Plans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlans

`func (o *Planner) SetPlans(v []MicrosoftGraphPlannerPlan)`

SetPlans sets Plans field to given value.

### HasPlans

`func (o *Planner) HasPlans() bool`

HasPlans returns a boolean if a field has been set.

### GetTasks

`func (o *Planner) GetTasks() []MicrosoftGraphPlannerTask`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *Planner) GetTasksOk() (*[]MicrosoftGraphPlannerTask, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *Planner) SetTasks(v []MicrosoftGraphPlannerTask)`

SetTasks sets Tasks field to given value.

### HasTasks

`func (o *Planner) HasTasks() bool`

HasTasks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


