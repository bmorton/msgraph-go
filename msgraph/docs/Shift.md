# Shift

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DraftShift** | Pointer to [**AnyOfmicrosoftGraphShiftItem**](anyOf&lt;microsoft.graph.shiftItem&gt;.md) | The draft version of this shift that is viewable by managers. Required. | [optional] 
**SchedulingGroupId** | Pointer to **NullableString** | ID of the scheduling group the shift is part of. Required. | [optional] 
**SharedShift** | Pointer to [**AnyOfmicrosoftGraphShiftItem**](anyOf&lt;microsoft.graph.shiftItem&gt;.md) | The shared version of this shift that is viewable by both employees and managers. Required. | [optional] 
**UserId** | Pointer to **NullableString** | ID of the user assigned to the shift. Required. | [optional] 

## Methods

### NewShift

`func NewShift() *Shift`

NewShift instantiates a new Shift object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShiftWithDefaults

`func NewShiftWithDefaults() *Shift`

NewShiftWithDefaults instantiates a new Shift object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDraftShift

`func (o *Shift) GetDraftShift() AnyOfmicrosoftGraphShiftItem`

GetDraftShift returns the DraftShift field if non-nil, zero value otherwise.

### GetDraftShiftOk

`func (o *Shift) GetDraftShiftOk() (*AnyOfmicrosoftGraphShiftItem, bool)`

GetDraftShiftOk returns a tuple with the DraftShift field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDraftShift

`func (o *Shift) SetDraftShift(v AnyOfmicrosoftGraphShiftItem)`

SetDraftShift sets DraftShift field to given value.

### HasDraftShift

`func (o *Shift) HasDraftShift() bool`

HasDraftShift returns a boolean if a field has been set.

### SetDraftShiftNil

`func (o *Shift) SetDraftShiftNil(b bool)`

 SetDraftShiftNil sets the value for DraftShift to be an explicit nil

### UnsetDraftShift
`func (o *Shift) UnsetDraftShift()`

UnsetDraftShift ensures that no value is present for DraftShift, not even an explicit nil
### GetSchedulingGroupId

`func (o *Shift) GetSchedulingGroupId() string`

GetSchedulingGroupId returns the SchedulingGroupId field if non-nil, zero value otherwise.

### GetSchedulingGroupIdOk

`func (o *Shift) GetSchedulingGroupIdOk() (*string, bool)`

GetSchedulingGroupIdOk returns a tuple with the SchedulingGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedulingGroupId

`func (o *Shift) SetSchedulingGroupId(v string)`

SetSchedulingGroupId sets SchedulingGroupId field to given value.

### HasSchedulingGroupId

`func (o *Shift) HasSchedulingGroupId() bool`

HasSchedulingGroupId returns a boolean if a field has been set.

### SetSchedulingGroupIdNil

`func (o *Shift) SetSchedulingGroupIdNil(b bool)`

 SetSchedulingGroupIdNil sets the value for SchedulingGroupId to be an explicit nil

### UnsetSchedulingGroupId
`func (o *Shift) UnsetSchedulingGroupId()`

UnsetSchedulingGroupId ensures that no value is present for SchedulingGroupId, not even an explicit nil
### GetSharedShift

`func (o *Shift) GetSharedShift() AnyOfmicrosoftGraphShiftItem`

GetSharedShift returns the SharedShift field if non-nil, zero value otherwise.

### GetSharedShiftOk

`func (o *Shift) GetSharedShiftOk() (*AnyOfmicrosoftGraphShiftItem, bool)`

GetSharedShiftOk returns a tuple with the SharedShift field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedShift

`func (o *Shift) SetSharedShift(v AnyOfmicrosoftGraphShiftItem)`

SetSharedShift sets SharedShift field to given value.

### HasSharedShift

`func (o *Shift) HasSharedShift() bool`

HasSharedShift returns a boolean if a field has been set.

### SetSharedShiftNil

`func (o *Shift) SetSharedShiftNil(b bool)`

 SetSharedShiftNil sets the value for SharedShift to be an explicit nil

### UnsetSharedShift
`func (o *Shift) UnsetSharedShift()`

UnsetSharedShift ensures that no value is present for SharedShift, not even an explicit nil
### GetUserId

`func (o *Shift) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *Shift) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *Shift) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *Shift) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *Shift) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *Shift) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


