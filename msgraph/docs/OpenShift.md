# OpenShift

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DraftOpenShift** | Pointer to [**AnyOfmicrosoftGraphOpenShiftItem**](anyOf&lt;microsoft.graph.openShiftItem&gt;.md) | An unpublished open shift. | [optional] 
**SchedulingGroupId** | Pointer to **NullableString** | ID for the scheduling group that the open shift belongs to. | [optional] 
**SharedOpenShift** | Pointer to [**AnyOfmicrosoftGraphOpenShiftItem**](anyOf&lt;microsoft.graph.openShiftItem&gt;.md) | A published open shift. | [optional] 

## Methods

### NewOpenShift

`func NewOpenShift() *OpenShift`

NewOpenShift instantiates a new OpenShift object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenShiftWithDefaults

`func NewOpenShiftWithDefaults() *OpenShift`

NewOpenShiftWithDefaults instantiates a new OpenShift object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDraftOpenShift

`func (o *OpenShift) GetDraftOpenShift() AnyOfmicrosoftGraphOpenShiftItem`

GetDraftOpenShift returns the DraftOpenShift field if non-nil, zero value otherwise.

### GetDraftOpenShiftOk

`func (o *OpenShift) GetDraftOpenShiftOk() (*AnyOfmicrosoftGraphOpenShiftItem, bool)`

GetDraftOpenShiftOk returns a tuple with the DraftOpenShift field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDraftOpenShift

`func (o *OpenShift) SetDraftOpenShift(v AnyOfmicrosoftGraphOpenShiftItem)`

SetDraftOpenShift sets DraftOpenShift field to given value.

### HasDraftOpenShift

`func (o *OpenShift) HasDraftOpenShift() bool`

HasDraftOpenShift returns a boolean if a field has been set.

### SetDraftOpenShiftNil

`func (o *OpenShift) SetDraftOpenShiftNil(b bool)`

 SetDraftOpenShiftNil sets the value for DraftOpenShift to be an explicit nil

### UnsetDraftOpenShift
`func (o *OpenShift) UnsetDraftOpenShift()`

UnsetDraftOpenShift ensures that no value is present for DraftOpenShift, not even an explicit nil
### GetSchedulingGroupId

`func (o *OpenShift) GetSchedulingGroupId() string`

GetSchedulingGroupId returns the SchedulingGroupId field if non-nil, zero value otherwise.

### GetSchedulingGroupIdOk

`func (o *OpenShift) GetSchedulingGroupIdOk() (*string, bool)`

GetSchedulingGroupIdOk returns a tuple with the SchedulingGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedulingGroupId

`func (o *OpenShift) SetSchedulingGroupId(v string)`

SetSchedulingGroupId sets SchedulingGroupId field to given value.

### HasSchedulingGroupId

`func (o *OpenShift) HasSchedulingGroupId() bool`

HasSchedulingGroupId returns a boolean if a field has been set.

### SetSchedulingGroupIdNil

`func (o *OpenShift) SetSchedulingGroupIdNil(b bool)`

 SetSchedulingGroupIdNil sets the value for SchedulingGroupId to be an explicit nil

### UnsetSchedulingGroupId
`func (o *OpenShift) UnsetSchedulingGroupId()`

UnsetSchedulingGroupId ensures that no value is present for SchedulingGroupId, not even an explicit nil
### GetSharedOpenShift

`func (o *OpenShift) GetSharedOpenShift() AnyOfmicrosoftGraphOpenShiftItem`

GetSharedOpenShift returns the SharedOpenShift field if non-nil, zero value otherwise.

### GetSharedOpenShiftOk

`func (o *OpenShift) GetSharedOpenShiftOk() (*AnyOfmicrosoftGraphOpenShiftItem, bool)`

GetSharedOpenShiftOk returns a tuple with the SharedOpenShift field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedOpenShift

`func (o *OpenShift) SetSharedOpenShift(v AnyOfmicrosoftGraphOpenShiftItem)`

SetSharedOpenShift sets SharedOpenShift field to given value.

### HasSharedOpenShift

`func (o *OpenShift) HasSharedOpenShift() bool`

HasSharedOpenShift returns a boolean if a field has been set.

### SetSharedOpenShiftNil

`func (o *OpenShift) SetSharedOpenShiftNil(b bool)`

 SetSharedOpenShiftNil sets the value for SharedOpenShift to be an explicit nil

### UnsetSharedOpenShift
`func (o *OpenShift) UnsetSharedOpenShift()`

UnsetSharedOpenShift ensures that no value is present for SharedOpenShift, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


