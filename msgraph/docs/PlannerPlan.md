# PlannerPlan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedBy** | Pointer to [**AnyOfmicrosoftGraphIdentitySet**](anyOf&lt;microsoft.graph.identitySet&gt;.md) | Read-only. The user who created the plan. | [optional] 
**CreatedDateTime** | Pointer to **NullableTime** | Read-only. Date and time at which the plan is created. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z | [optional] 
**Owner** | Pointer to **NullableString** | ID of the Group that owns the plan. A valid group must exist before this field can be set. After it is set, this property can’t be updated. | [optional] 
**Title** | Pointer to **string** | Required. Title of the plan. | [optional] 
**Buckets** | Pointer to [**[]MicrosoftGraphPlannerBucket**](MicrosoftGraphPlannerBucket.md) | Read-only. Nullable. Collection of buckets in the plan. | [optional] 
**Details** | Pointer to [**AnyOfmicrosoftGraphPlannerPlanDetails**](anyOf&lt;microsoft.graph.plannerPlanDetails&gt;.md) | Read-only. Nullable. Additional details about the plan. | [optional] 
**Tasks** | Pointer to [**[]MicrosoftGraphPlannerTask**](MicrosoftGraphPlannerTask.md) | Read-only. Nullable. Collection of tasks in the plan. | [optional] 

## Methods

### NewPlannerPlan

`func NewPlannerPlan() *PlannerPlan`

NewPlannerPlan instantiates a new PlannerPlan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlannerPlanWithDefaults

`func NewPlannerPlanWithDefaults() *PlannerPlan`

NewPlannerPlanWithDefaults instantiates a new PlannerPlan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedBy

`func (o *PlannerPlan) GetCreatedBy() AnyOfmicrosoftGraphIdentitySet`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *PlannerPlan) GetCreatedByOk() (*AnyOfmicrosoftGraphIdentitySet, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *PlannerPlan) SetCreatedBy(v AnyOfmicrosoftGraphIdentitySet)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *PlannerPlan) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedByNil

`func (o *PlannerPlan) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *PlannerPlan) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
### GetCreatedDateTime

`func (o *PlannerPlan) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *PlannerPlan) GetCreatedDateTimeOk() (*time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateTime

`func (o *PlannerPlan) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime sets CreatedDateTime field to given value.

### HasCreatedDateTime

`func (o *PlannerPlan) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTimeNil

`func (o *PlannerPlan) SetCreatedDateTimeNil(b bool)`

 SetCreatedDateTimeNil sets the value for CreatedDateTime to be an explicit nil

### UnsetCreatedDateTime
`func (o *PlannerPlan) UnsetCreatedDateTime()`

UnsetCreatedDateTime ensures that no value is present for CreatedDateTime, not even an explicit nil
### GetOwner

`func (o *PlannerPlan) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *PlannerPlan) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *PlannerPlan) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *PlannerPlan) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *PlannerPlan) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *PlannerPlan) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetTitle

`func (o *PlannerPlan) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *PlannerPlan) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *PlannerPlan) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *PlannerPlan) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetBuckets

`func (o *PlannerPlan) GetBuckets() []MicrosoftGraphPlannerBucket`

GetBuckets returns the Buckets field if non-nil, zero value otherwise.

### GetBucketsOk

`func (o *PlannerPlan) GetBucketsOk() (*[]MicrosoftGraphPlannerBucket, bool)`

GetBucketsOk returns a tuple with the Buckets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuckets

`func (o *PlannerPlan) SetBuckets(v []MicrosoftGraphPlannerBucket)`

SetBuckets sets Buckets field to given value.

### HasBuckets

`func (o *PlannerPlan) HasBuckets() bool`

HasBuckets returns a boolean if a field has been set.

### GetDetails

`func (o *PlannerPlan) GetDetails() AnyOfmicrosoftGraphPlannerPlanDetails`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *PlannerPlan) GetDetailsOk() (*AnyOfmicrosoftGraphPlannerPlanDetails, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *PlannerPlan) SetDetails(v AnyOfmicrosoftGraphPlannerPlanDetails)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *PlannerPlan) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### SetDetailsNil

`func (o *PlannerPlan) SetDetailsNil(b bool)`

 SetDetailsNil sets the value for Details to be an explicit nil

### UnsetDetails
`func (o *PlannerPlan) UnsetDetails()`

UnsetDetails ensures that no value is present for Details, not even an explicit nil
### GetTasks

`func (o *PlannerPlan) GetTasks() []MicrosoftGraphPlannerTask`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *PlannerPlan) GetTasksOk() (*[]MicrosoftGraphPlannerTask, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *PlannerPlan) SetTasks(v []MicrosoftGraphPlannerTask)`

SetTasks sets Tasks field to given value.

### HasTasks

`func (o *PlannerPlan) HasTasks() bool`

HasTasks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


