# MicrosoftGraphPlannerBucket

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**Name** | Pointer to **string** | Name of the bucket. | [optional] 
**OrderHint** | Pointer to **NullableString** | Hint used to order items of this type in a list view. The format is defined as outlined here. | [optional] 
**PlanId** | Pointer to **NullableString** | Plan ID to which the bucket belongs. | [optional] 
**Tasks** | Pointer to [**[]MicrosoftGraphPlannerTask**](MicrosoftGraphPlannerTask.md) | Read-only. Nullable. The collection of tasks in the bucket. | [optional] 

## Methods

### NewMicrosoftGraphPlannerBucket

`func NewMicrosoftGraphPlannerBucket() *MicrosoftGraphPlannerBucket`

NewMicrosoftGraphPlannerBucket instantiates a new MicrosoftGraphPlannerBucket object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphPlannerBucketWithDefaults

`func NewMicrosoftGraphPlannerBucketWithDefaults() *MicrosoftGraphPlannerBucket`

NewMicrosoftGraphPlannerBucketWithDefaults instantiates a new MicrosoftGraphPlannerBucket object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphPlannerBucket) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphPlannerBucket) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphPlannerBucket) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphPlannerBucket) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *MicrosoftGraphPlannerBucket) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MicrosoftGraphPlannerBucket) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MicrosoftGraphPlannerBucket) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MicrosoftGraphPlannerBucket) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrderHint

`func (o *MicrosoftGraphPlannerBucket) GetOrderHint() string`

GetOrderHint returns the OrderHint field if non-nil, zero value otherwise.

### GetOrderHintOk

`func (o *MicrosoftGraphPlannerBucket) GetOrderHintOk() (*string, bool)`

GetOrderHintOk returns a tuple with the OrderHint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderHint

`func (o *MicrosoftGraphPlannerBucket) SetOrderHint(v string)`

SetOrderHint sets OrderHint field to given value.

### HasOrderHint

`func (o *MicrosoftGraphPlannerBucket) HasOrderHint() bool`

HasOrderHint returns a boolean if a field has been set.

### SetOrderHintNil

`func (o *MicrosoftGraphPlannerBucket) SetOrderHintNil(b bool)`

 SetOrderHintNil sets the value for OrderHint to be an explicit nil

### UnsetOrderHint
`func (o *MicrosoftGraphPlannerBucket) UnsetOrderHint()`

UnsetOrderHint ensures that no value is present for OrderHint, not even an explicit nil
### GetPlanId

`func (o *MicrosoftGraphPlannerBucket) GetPlanId() string`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *MicrosoftGraphPlannerBucket) GetPlanIdOk() (*string, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *MicrosoftGraphPlannerBucket) SetPlanId(v string)`

SetPlanId sets PlanId field to given value.

### HasPlanId

`func (o *MicrosoftGraphPlannerBucket) HasPlanId() bool`

HasPlanId returns a boolean if a field has been set.

### SetPlanIdNil

`func (o *MicrosoftGraphPlannerBucket) SetPlanIdNil(b bool)`

 SetPlanIdNil sets the value for PlanId to be an explicit nil

### UnsetPlanId
`func (o *MicrosoftGraphPlannerBucket) UnsetPlanId()`

UnsetPlanId ensures that no value is present for PlanId, not even an explicit nil
### GetTasks

`func (o *MicrosoftGraphPlannerBucket) GetTasks() []MicrosoftGraphPlannerTask`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *MicrosoftGraphPlannerBucket) GetTasksOk() (*[]MicrosoftGraphPlannerTask, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *MicrosoftGraphPlannerBucket) SetTasks(v []MicrosoftGraphPlannerTask)`

SetTasks sets Tasks field to given value.

### HasTasks

`func (o *MicrosoftGraphPlannerBucket) HasTasks() bool`

HasTasks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


