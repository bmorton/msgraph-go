# PlannerAssignedToTaskBoardTaskFormat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderHintsByAssignee** | Pointer to [**AnyOfobject**](anyOf&lt;object&gt;.md) | Dictionary of hints used to order tasks on the AssignedTo view of the Task Board. The key of each entry is one of the users the task is assigned to and the value is the order hint. The format of each value is defined as outlined here. | [optional] 
**UnassignedOrderHint** | Pointer to **NullableString** | Hint value used to order the task on the AssignedTo view of the Task Board when the task is not assigned to anyone, or if the orderHintsByAssignee dictionary does not provide an order hint for the user the task is assigned to. The format is defined as outlined here. | [optional] 

## Methods

### NewPlannerAssignedToTaskBoardTaskFormat

`func NewPlannerAssignedToTaskBoardTaskFormat() *PlannerAssignedToTaskBoardTaskFormat`

NewPlannerAssignedToTaskBoardTaskFormat instantiates a new PlannerAssignedToTaskBoardTaskFormat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlannerAssignedToTaskBoardTaskFormatWithDefaults

`func NewPlannerAssignedToTaskBoardTaskFormatWithDefaults() *PlannerAssignedToTaskBoardTaskFormat`

NewPlannerAssignedToTaskBoardTaskFormatWithDefaults instantiates a new PlannerAssignedToTaskBoardTaskFormat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderHintsByAssignee

`func (o *PlannerAssignedToTaskBoardTaskFormat) GetOrderHintsByAssignee() AnyOfobject`

GetOrderHintsByAssignee returns the OrderHintsByAssignee field if non-nil, zero value otherwise.

### GetOrderHintsByAssigneeOk

`func (o *PlannerAssignedToTaskBoardTaskFormat) GetOrderHintsByAssigneeOk() (*AnyOfobject, bool)`

GetOrderHintsByAssigneeOk returns a tuple with the OrderHintsByAssignee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderHintsByAssignee

`func (o *PlannerAssignedToTaskBoardTaskFormat) SetOrderHintsByAssignee(v AnyOfobject)`

SetOrderHintsByAssignee sets OrderHintsByAssignee field to given value.

### HasOrderHintsByAssignee

`func (o *PlannerAssignedToTaskBoardTaskFormat) HasOrderHintsByAssignee() bool`

HasOrderHintsByAssignee returns a boolean if a field has been set.

### SetOrderHintsByAssigneeNil

`func (o *PlannerAssignedToTaskBoardTaskFormat) SetOrderHintsByAssigneeNil(b bool)`

 SetOrderHintsByAssigneeNil sets the value for OrderHintsByAssignee to be an explicit nil

### UnsetOrderHintsByAssignee
`func (o *PlannerAssignedToTaskBoardTaskFormat) UnsetOrderHintsByAssignee()`

UnsetOrderHintsByAssignee ensures that no value is present for OrderHintsByAssignee, not even an explicit nil
### GetUnassignedOrderHint

`func (o *PlannerAssignedToTaskBoardTaskFormat) GetUnassignedOrderHint() string`

GetUnassignedOrderHint returns the UnassignedOrderHint field if non-nil, zero value otherwise.

### GetUnassignedOrderHintOk

`func (o *PlannerAssignedToTaskBoardTaskFormat) GetUnassignedOrderHintOk() (*string, bool)`

GetUnassignedOrderHintOk returns a tuple with the UnassignedOrderHint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnassignedOrderHint

`func (o *PlannerAssignedToTaskBoardTaskFormat) SetUnassignedOrderHint(v string)`

SetUnassignedOrderHint sets UnassignedOrderHint field to given value.

### HasUnassignedOrderHint

`func (o *PlannerAssignedToTaskBoardTaskFormat) HasUnassignedOrderHint() bool`

HasUnassignedOrderHint returns a boolean if a field has been set.

### SetUnassignedOrderHintNil

`func (o *PlannerAssignedToTaskBoardTaskFormat) SetUnassignedOrderHintNil(b bool)`

 SetUnassignedOrderHintNil sets the value for UnassignedOrderHint to be an explicit nil

### UnsetUnassignedOrderHint
`func (o *PlannerAssignedToTaskBoardTaskFormat) UnsetUnassignedOrderHint()`

UnsetUnassignedOrderHint ensures that no value is present for UnassignedOrderHint, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


