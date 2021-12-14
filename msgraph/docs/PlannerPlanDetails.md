# PlannerPlanDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CategoryDescriptions** | Pointer to [**AnyOfmicrosoftGraphPlannerCategoryDescriptions**](anyOf&lt;microsoft.graph.plannerCategoryDescriptions&gt;.md) | An object that specifies the descriptions of the six categories that can be associated with tasks in the plan | [optional] 
**SharedWith** | Pointer to [**AnyOfobject**](anyOf&lt;object&gt;.md) | Set of user ids that this plan is shared with. If you are leveraging Microsoft 365 groups, use the Groups API to manage group membership to share the group&#39;s plan. You can also add existing members of the group to this collection though it is not required for them to access the plan owned by the group. | [optional] 

## Methods

### NewPlannerPlanDetails

`func NewPlannerPlanDetails() *PlannerPlanDetails`

NewPlannerPlanDetails instantiates a new PlannerPlanDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlannerPlanDetailsWithDefaults

`func NewPlannerPlanDetailsWithDefaults() *PlannerPlanDetails`

NewPlannerPlanDetailsWithDefaults instantiates a new PlannerPlanDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategoryDescriptions

`func (o *PlannerPlanDetails) GetCategoryDescriptions() AnyOfmicrosoftGraphPlannerCategoryDescriptions`

GetCategoryDescriptions returns the CategoryDescriptions field if non-nil, zero value otherwise.

### GetCategoryDescriptionsOk

`func (o *PlannerPlanDetails) GetCategoryDescriptionsOk() (*AnyOfmicrosoftGraphPlannerCategoryDescriptions, bool)`

GetCategoryDescriptionsOk returns a tuple with the CategoryDescriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryDescriptions

`func (o *PlannerPlanDetails) SetCategoryDescriptions(v AnyOfmicrosoftGraphPlannerCategoryDescriptions)`

SetCategoryDescriptions sets CategoryDescriptions field to given value.

### HasCategoryDescriptions

`func (o *PlannerPlanDetails) HasCategoryDescriptions() bool`

HasCategoryDescriptions returns a boolean if a field has been set.

### SetCategoryDescriptionsNil

`func (o *PlannerPlanDetails) SetCategoryDescriptionsNil(b bool)`

 SetCategoryDescriptionsNil sets the value for CategoryDescriptions to be an explicit nil

### UnsetCategoryDescriptions
`func (o *PlannerPlanDetails) UnsetCategoryDescriptions()`

UnsetCategoryDescriptions ensures that no value is present for CategoryDescriptions, not even an explicit nil
### GetSharedWith

`func (o *PlannerPlanDetails) GetSharedWith() AnyOfobject`

GetSharedWith returns the SharedWith field if non-nil, zero value otherwise.

### GetSharedWithOk

`func (o *PlannerPlanDetails) GetSharedWithOk() (*AnyOfobject, bool)`

GetSharedWithOk returns a tuple with the SharedWith field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedWith

`func (o *PlannerPlanDetails) SetSharedWith(v AnyOfobject)`

SetSharedWith sets SharedWith field to given value.

### HasSharedWith

`func (o *PlannerPlanDetails) HasSharedWith() bool`

HasSharedWith returns a boolean if a field has been set.

### SetSharedWithNil

`func (o *PlannerPlanDetails) SetSharedWithNil(b bool)`

 SetSharedWithNil sets the value for SharedWith to be an explicit nil

### UnsetSharedWith
`func (o *PlannerPlanDetails) UnsetSharedWith()`

UnsetSharedWith ensures that no value is present for SharedWith, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


