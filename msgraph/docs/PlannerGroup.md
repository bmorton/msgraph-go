# PlannerGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Plans** | Pointer to [**[]MicrosoftGraphPlannerPlan**](MicrosoftGraphPlannerPlan.md) | Read-only. Nullable. Returns the plannerPlans owned by the group. | [optional] 

## Methods

### NewPlannerGroup

`func NewPlannerGroup() *PlannerGroup`

NewPlannerGroup instantiates a new PlannerGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlannerGroupWithDefaults

`func NewPlannerGroupWithDefaults() *PlannerGroup`

NewPlannerGroupWithDefaults instantiates a new PlannerGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlans

`func (o *PlannerGroup) GetPlans() []MicrosoftGraphPlannerPlan`

GetPlans returns the Plans field if non-nil, zero value otherwise.

### GetPlansOk

`func (o *PlannerGroup) GetPlansOk() (*[]MicrosoftGraphPlannerPlan, bool)`

GetPlansOk returns a tuple with the Plans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlans

`func (o *PlannerGroup) SetPlans(v []MicrosoftGraphPlannerPlan)`

SetPlans sets Plans field to given value.

### HasPlans

`func (o *PlannerGroup) HasPlans() bool`

HasPlans returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


