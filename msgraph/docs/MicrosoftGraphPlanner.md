# MicrosoftGraphPlanner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**Buckets** | Pointer to [**[]MicrosoftGraphPlannerBucket**](MicrosoftGraphPlannerBucket.md) | Read-only. Nullable. Returns a collection of the specified buckets | [optional] 
**Plans** | Pointer to [**[]MicrosoftGraphPlannerPlan**](MicrosoftGraphPlannerPlan.md) | Read-only. Nullable. Returns a collection of the specified plans | [optional] 
**Tasks** | Pointer to [**[]MicrosoftGraphPlannerTask**](MicrosoftGraphPlannerTask.md) | Read-only. Nullable. Returns a collection of the specified tasks | [optional] 

## Methods

### NewMicrosoftGraphPlanner

`func NewMicrosoftGraphPlanner() *MicrosoftGraphPlanner`

NewMicrosoftGraphPlanner instantiates a new MicrosoftGraphPlanner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphPlannerWithDefaults

`func NewMicrosoftGraphPlannerWithDefaults() *MicrosoftGraphPlanner`

NewMicrosoftGraphPlannerWithDefaults instantiates a new MicrosoftGraphPlanner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphPlanner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphPlanner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphPlanner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphPlanner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetBuckets

`func (o *MicrosoftGraphPlanner) GetBuckets() []MicrosoftGraphPlannerBucket`

GetBuckets returns the Buckets field if non-nil, zero value otherwise.

### GetBucketsOk

`func (o *MicrosoftGraphPlanner) GetBucketsOk() (*[]MicrosoftGraphPlannerBucket, bool)`

GetBucketsOk returns a tuple with the Buckets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuckets

`func (o *MicrosoftGraphPlanner) SetBuckets(v []MicrosoftGraphPlannerBucket)`

SetBuckets sets Buckets field to given value.

### HasBuckets

`func (o *MicrosoftGraphPlanner) HasBuckets() bool`

HasBuckets returns a boolean if a field has been set.

### GetPlans

`func (o *MicrosoftGraphPlanner) GetPlans() []MicrosoftGraphPlannerPlan`

GetPlans returns the Plans field if non-nil, zero value otherwise.

### GetPlansOk

`func (o *MicrosoftGraphPlanner) GetPlansOk() (*[]MicrosoftGraphPlannerPlan, bool)`

GetPlansOk returns a tuple with the Plans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlans

`func (o *MicrosoftGraphPlanner) SetPlans(v []MicrosoftGraphPlannerPlan)`

SetPlans sets Plans field to given value.

### HasPlans

`func (o *MicrosoftGraphPlanner) HasPlans() bool`

HasPlans returns a boolean if a field has been set.

### GetTasks

`func (o *MicrosoftGraphPlanner) GetTasks() []MicrosoftGraphPlannerTask`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *MicrosoftGraphPlanner) GetTasksOk() (*[]MicrosoftGraphPlannerTask, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *MicrosoftGraphPlanner) SetTasks(v []MicrosoftGraphPlannerTask)`

SetTasks sets Tasks field to given value.

### HasTasks

`func (o *MicrosoftGraphPlanner) HasTasks() bool`

HasTasks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


