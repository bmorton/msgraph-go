# MicrosoftGraphShift

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**CreatedDateTime** | Pointer to **NullableTime** | The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z | [optional] 
**LastModifiedBy** | Pointer to [**AnyOfmicrosoftGraphIdentitySet**](anyOf&lt;microsoft.graph.identitySet&gt;.md) | Identity of the person who last modified the entity. | [optional] 
**LastModifiedDateTime** | Pointer to **NullableTime** | The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z | [optional] 
**DraftShift** | Pointer to [**AnyOfmicrosoftGraphShiftItem**](anyOf&lt;microsoft.graph.shiftItem&gt;.md) | The draft version of this shift that is viewable by managers. Required. | [optional] 
**SchedulingGroupId** | Pointer to **NullableString** | ID of the scheduling group the shift is part of. Required. | [optional] 
**SharedShift** | Pointer to [**AnyOfmicrosoftGraphShiftItem**](anyOf&lt;microsoft.graph.shiftItem&gt;.md) | The shared version of this shift that is viewable by both employees and managers. Required. | [optional] 
**UserId** | Pointer to **NullableString** | ID of the user assigned to the shift. Required. | [optional] 

## Methods

### NewMicrosoftGraphShift

`func NewMicrosoftGraphShift() *MicrosoftGraphShift`

NewMicrosoftGraphShift instantiates a new MicrosoftGraphShift object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphShiftWithDefaults

`func NewMicrosoftGraphShiftWithDefaults() *MicrosoftGraphShift`

NewMicrosoftGraphShiftWithDefaults instantiates a new MicrosoftGraphShift object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphShift) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphShift) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphShift) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphShift) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedDateTime

`func (o *MicrosoftGraphShift) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *MicrosoftGraphShift) GetCreatedDateTimeOk() (*time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateTime

`func (o *MicrosoftGraphShift) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime sets CreatedDateTime field to given value.

### HasCreatedDateTime

`func (o *MicrosoftGraphShift) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTimeNil

`func (o *MicrosoftGraphShift) SetCreatedDateTimeNil(b bool)`

 SetCreatedDateTimeNil sets the value for CreatedDateTime to be an explicit nil

### UnsetCreatedDateTime
`func (o *MicrosoftGraphShift) UnsetCreatedDateTime()`

UnsetCreatedDateTime ensures that no value is present for CreatedDateTime, not even an explicit nil
### GetLastModifiedBy

`func (o *MicrosoftGraphShift) GetLastModifiedBy() AnyOfmicrosoftGraphIdentitySet`

GetLastModifiedBy returns the LastModifiedBy field if non-nil, zero value otherwise.

### GetLastModifiedByOk

`func (o *MicrosoftGraphShift) GetLastModifiedByOk() (*AnyOfmicrosoftGraphIdentitySet, bool)`

GetLastModifiedByOk returns a tuple with the LastModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedBy

`func (o *MicrosoftGraphShift) SetLastModifiedBy(v AnyOfmicrosoftGraphIdentitySet)`

SetLastModifiedBy sets LastModifiedBy field to given value.

### HasLastModifiedBy

`func (o *MicrosoftGraphShift) HasLastModifiedBy() bool`

HasLastModifiedBy returns a boolean if a field has been set.

### SetLastModifiedByNil

`func (o *MicrosoftGraphShift) SetLastModifiedByNil(b bool)`

 SetLastModifiedByNil sets the value for LastModifiedBy to be an explicit nil

### UnsetLastModifiedBy
`func (o *MicrosoftGraphShift) UnsetLastModifiedBy()`

UnsetLastModifiedBy ensures that no value is present for LastModifiedBy, not even an explicit nil
### GetLastModifiedDateTime

`func (o *MicrosoftGraphShift) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *MicrosoftGraphShift) GetLastModifiedDateTimeOk() (*time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedDateTime

`func (o *MicrosoftGraphShift) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime sets LastModifiedDateTime field to given value.

### HasLastModifiedDateTime

`func (o *MicrosoftGraphShift) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### SetLastModifiedDateTimeNil

`func (o *MicrosoftGraphShift) SetLastModifiedDateTimeNil(b bool)`

 SetLastModifiedDateTimeNil sets the value for LastModifiedDateTime to be an explicit nil

### UnsetLastModifiedDateTime
`func (o *MicrosoftGraphShift) UnsetLastModifiedDateTime()`

UnsetLastModifiedDateTime ensures that no value is present for LastModifiedDateTime, not even an explicit nil
### GetDraftShift

`func (o *MicrosoftGraphShift) GetDraftShift() AnyOfmicrosoftGraphShiftItem`

GetDraftShift returns the DraftShift field if non-nil, zero value otherwise.

### GetDraftShiftOk

`func (o *MicrosoftGraphShift) GetDraftShiftOk() (*AnyOfmicrosoftGraphShiftItem, bool)`

GetDraftShiftOk returns a tuple with the DraftShift field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDraftShift

`func (o *MicrosoftGraphShift) SetDraftShift(v AnyOfmicrosoftGraphShiftItem)`

SetDraftShift sets DraftShift field to given value.

### HasDraftShift

`func (o *MicrosoftGraphShift) HasDraftShift() bool`

HasDraftShift returns a boolean if a field has been set.

### SetDraftShiftNil

`func (o *MicrosoftGraphShift) SetDraftShiftNil(b bool)`

 SetDraftShiftNil sets the value for DraftShift to be an explicit nil

### UnsetDraftShift
`func (o *MicrosoftGraphShift) UnsetDraftShift()`

UnsetDraftShift ensures that no value is present for DraftShift, not even an explicit nil
### GetSchedulingGroupId

`func (o *MicrosoftGraphShift) GetSchedulingGroupId() string`

GetSchedulingGroupId returns the SchedulingGroupId field if non-nil, zero value otherwise.

### GetSchedulingGroupIdOk

`func (o *MicrosoftGraphShift) GetSchedulingGroupIdOk() (*string, bool)`

GetSchedulingGroupIdOk returns a tuple with the SchedulingGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedulingGroupId

`func (o *MicrosoftGraphShift) SetSchedulingGroupId(v string)`

SetSchedulingGroupId sets SchedulingGroupId field to given value.

### HasSchedulingGroupId

`func (o *MicrosoftGraphShift) HasSchedulingGroupId() bool`

HasSchedulingGroupId returns a boolean if a field has been set.

### SetSchedulingGroupIdNil

`func (o *MicrosoftGraphShift) SetSchedulingGroupIdNil(b bool)`

 SetSchedulingGroupIdNil sets the value for SchedulingGroupId to be an explicit nil

### UnsetSchedulingGroupId
`func (o *MicrosoftGraphShift) UnsetSchedulingGroupId()`

UnsetSchedulingGroupId ensures that no value is present for SchedulingGroupId, not even an explicit nil
### GetSharedShift

`func (o *MicrosoftGraphShift) GetSharedShift() AnyOfmicrosoftGraphShiftItem`

GetSharedShift returns the SharedShift field if non-nil, zero value otherwise.

### GetSharedShiftOk

`func (o *MicrosoftGraphShift) GetSharedShiftOk() (*AnyOfmicrosoftGraphShiftItem, bool)`

GetSharedShiftOk returns a tuple with the SharedShift field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedShift

`func (o *MicrosoftGraphShift) SetSharedShift(v AnyOfmicrosoftGraphShiftItem)`

SetSharedShift sets SharedShift field to given value.

### HasSharedShift

`func (o *MicrosoftGraphShift) HasSharedShift() bool`

HasSharedShift returns a boolean if a field has been set.

### SetSharedShiftNil

`func (o *MicrosoftGraphShift) SetSharedShiftNil(b bool)`

 SetSharedShiftNil sets the value for SharedShift to be an explicit nil

### UnsetSharedShift
`func (o *MicrosoftGraphShift) UnsetSharedShift()`

UnsetSharedShift ensures that no value is present for SharedShift, not even an explicit nil
### GetUserId

`func (o *MicrosoftGraphShift) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *MicrosoftGraphShift) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *MicrosoftGraphShift) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *MicrosoftGraphShift) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *MicrosoftGraphShift) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *MicrosoftGraphShift) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


