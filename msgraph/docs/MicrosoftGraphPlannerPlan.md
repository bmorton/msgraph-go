# MicrosoftGraphPlannerPlan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**CreatedBy** | Pointer to [**AnyOfmicrosoftGraphIdentitySet**](anyOf&lt;microsoft.graph.identitySet&gt;.md) | Read-only. The user who created the plan. | [optional] 
**CreatedDateTime** | Pointer to **NullableTime** | Read-only. Date and time at which the plan is created. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z | [optional] 
**Owner** | Pointer to **NullableString** | ID of the Group that owns the plan. A valid group must exist before this field can be set. After it is set, this property can’t be updated. | [optional] 
**Title** | Pointer to **string** | Required. Title of the plan. | [optional] 
**Buckets** | Pointer to [**[]MicrosoftGraphPlannerBucket**](MicrosoftGraphPlannerBucket.md) | Read-only. Nullable. Collection of buckets in the plan. | [optional] 
**Details** | Pointer to [**AnyOfmicrosoftGraphPlannerPlanDetails**](anyOf&lt;microsoft.graph.plannerPlanDetails&gt;.md) | Read-only. Nullable. Additional details about the plan. | [optional] 
**Tasks** | Pointer to [**[]MicrosoftGraphPlannerTask**](MicrosoftGraphPlannerTask.md) | Read-only. Nullable. Collection of tasks in the plan. | [optional] 

## Methods

### NewMicrosoftGraphPlannerPlan

`func NewMicrosoftGraphPlannerPlan() *MicrosoftGraphPlannerPlan`

NewMicrosoftGraphPlannerPlan instantiates a new MicrosoftGraphPlannerPlan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphPlannerPlanWithDefaults

`func NewMicrosoftGraphPlannerPlanWithDefaults() *MicrosoftGraphPlannerPlan`

NewMicrosoftGraphPlannerPlanWithDefaults instantiates a new MicrosoftGraphPlannerPlan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphPlannerPlan) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphPlannerPlan) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphPlannerPlan) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphPlannerPlan) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedBy

`func (o *MicrosoftGraphPlannerPlan) GetCreatedBy() AnyOfmicrosoftGraphIdentitySet`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *MicrosoftGraphPlannerPlan) GetCreatedByOk() (*AnyOfmicrosoftGraphIdentitySet, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *MicrosoftGraphPlannerPlan) SetCreatedBy(v AnyOfmicrosoftGraphIdentitySet)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *MicrosoftGraphPlannerPlan) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedByNil

`func (o *MicrosoftGraphPlannerPlan) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *MicrosoftGraphPlannerPlan) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
### GetCreatedDateTime

`func (o *MicrosoftGraphPlannerPlan) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *MicrosoftGraphPlannerPlan) GetCreatedDateTimeOk() (*time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateTime

`func (o *MicrosoftGraphPlannerPlan) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime sets CreatedDateTime field to given value.

### HasCreatedDateTime

`func (o *MicrosoftGraphPlannerPlan) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTimeNil

`func (o *MicrosoftGraphPlannerPlan) SetCreatedDateTimeNil(b bool)`

 SetCreatedDateTimeNil sets the value for CreatedDateTime to be an explicit nil

### UnsetCreatedDateTime
`func (o *MicrosoftGraphPlannerPlan) UnsetCreatedDateTime()`

UnsetCreatedDateTime ensures that no value is present for CreatedDateTime, not even an explicit nil
### GetOwner

`func (o *MicrosoftGraphPlannerPlan) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *MicrosoftGraphPlannerPlan) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *MicrosoftGraphPlannerPlan) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *MicrosoftGraphPlannerPlan) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *MicrosoftGraphPlannerPlan) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *MicrosoftGraphPlannerPlan) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetTitle

`func (o *MicrosoftGraphPlannerPlan) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *MicrosoftGraphPlannerPlan) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *MicrosoftGraphPlannerPlan) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *MicrosoftGraphPlannerPlan) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetBuckets

`func (o *MicrosoftGraphPlannerPlan) GetBuckets() []MicrosoftGraphPlannerBucket`

GetBuckets returns the Buckets field if non-nil, zero value otherwise.

### GetBucketsOk

`func (o *MicrosoftGraphPlannerPlan) GetBucketsOk() (*[]MicrosoftGraphPlannerBucket, bool)`

GetBucketsOk returns a tuple with the Buckets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuckets

`func (o *MicrosoftGraphPlannerPlan) SetBuckets(v []MicrosoftGraphPlannerBucket)`

SetBuckets sets Buckets field to given value.

### HasBuckets

`func (o *MicrosoftGraphPlannerPlan) HasBuckets() bool`

HasBuckets returns a boolean if a field has been set.

### GetDetails

`func (o *MicrosoftGraphPlannerPlan) GetDetails() AnyOfmicrosoftGraphPlannerPlanDetails`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *MicrosoftGraphPlannerPlan) GetDetailsOk() (*AnyOfmicrosoftGraphPlannerPlanDetails, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *MicrosoftGraphPlannerPlan) SetDetails(v AnyOfmicrosoftGraphPlannerPlanDetails)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *MicrosoftGraphPlannerPlan) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### SetDetailsNil

`func (o *MicrosoftGraphPlannerPlan) SetDetailsNil(b bool)`

 SetDetailsNil sets the value for Details to be an explicit nil

### UnsetDetails
`func (o *MicrosoftGraphPlannerPlan) UnsetDetails()`

UnsetDetails ensures that no value is present for Details, not even an explicit nil
### GetTasks

`func (o *MicrosoftGraphPlannerPlan) GetTasks() []MicrosoftGraphPlannerTask`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *MicrosoftGraphPlannerPlan) GetTasksOk() (*[]MicrosoftGraphPlannerTask, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *MicrosoftGraphPlannerPlan) SetTasks(v []MicrosoftGraphPlannerTask)`

SetTasks sets Tasks field to given value.

### HasTasks

`func (o *MicrosoftGraphPlannerPlan) HasTasks() bool`

HasTasks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


