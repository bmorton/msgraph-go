# PlannerBucket

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the bucket. | [optional] 
**OrderHint** | Pointer to **NullableString** | Hint used to order items of this type in a list view. The format is defined as outlined here. | [optional] 
**PlanId** | Pointer to **NullableString** | Plan ID to which the bucket belongs. | [optional] 
**Tasks** | Pointer to [**[]MicrosoftGraphPlannerTask**](MicrosoftGraphPlannerTask.md) | Read-only. Nullable. The collection of tasks in the bucket. | [optional] 

## Methods

### NewPlannerBucket

`func NewPlannerBucket() *PlannerBucket`

NewPlannerBucket instantiates a new PlannerBucket object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlannerBucketWithDefaults

`func NewPlannerBucketWithDefaults() *PlannerBucket`

NewPlannerBucketWithDefaults instantiates a new PlannerBucket object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PlannerBucket) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PlannerBucket) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PlannerBucket) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PlannerBucket) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrderHint

`func (o *PlannerBucket) GetOrderHint() string`

GetOrderHint returns the OrderHint field if non-nil, zero value otherwise.

### GetOrderHintOk

`func (o *PlannerBucket) GetOrderHintOk() (*string, bool)`

GetOrderHintOk returns a tuple with the OrderHint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderHint

`func (o *PlannerBucket) SetOrderHint(v string)`

SetOrderHint sets OrderHint field to given value.

### HasOrderHint

`func (o *PlannerBucket) HasOrderHint() bool`

HasOrderHint returns a boolean if a field has been set.

### SetOrderHintNil

`func (o *PlannerBucket) SetOrderHintNil(b bool)`

 SetOrderHintNil sets the value for OrderHint to be an explicit nil

### UnsetOrderHint
`func (o *PlannerBucket) UnsetOrderHint()`

UnsetOrderHint ensures that no value is present for OrderHint, not even an explicit nil
### GetPlanId

`func (o *PlannerBucket) GetPlanId() string`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *PlannerBucket) GetPlanIdOk() (*string, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *PlannerBucket) SetPlanId(v string)`

SetPlanId sets PlanId field to given value.

### HasPlanId

`func (o *PlannerBucket) HasPlanId() bool`

HasPlanId returns a boolean if a field has been set.

### SetPlanIdNil

`func (o *PlannerBucket) SetPlanIdNil(b bool)`

 SetPlanIdNil sets the value for PlanId to be an explicit nil

### UnsetPlanId
`func (o *PlannerBucket) UnsetPlanId()`

UnsetPlanId ensures that no value is present for PlanId, not even an explicit nil
### GetTasks

`func (o *PlannerBucket) GetTasks() []MicrosoftGraphPlannerTask`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *PlannerBucket) GetTasksOk() (*[]MicrosoftGraphPlannerTask, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *PlannerBucket) SetTasks(v []MicrosoftGraphPlannerTask)`

SetTasks sets Tasks field to given value.

### HasTasks

`func (o *PlannerBucket) HasTasks() bool`

HasTasks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


