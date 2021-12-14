# MicrosoftGraphOpenShift

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**CreatedDateTime** | Pointer to **NullableTime** | The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z | [optional] 
**LastModifiedBy** | Pointer to [**AnyOfmicrosoftGraphIdentitySet**](anyOf&lt;microsoft.graph.identitySet&gt;.md) | Identity of the person who last modified the entity. | [optional] 
**LastModifiedDateTime** | Pointer to **NullableTime** | The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z | [optional] 
**DraftOpenShift** | Pointer to [**AnyOfmicrosoftGraphOpenShiftItem**](anyOf&lt;microsoft.graph.openShiftItem&gt;.md) | An unpublished open shift. | [optional] 
**SchedulingGroupId** | Pointer to **NullableString** | ID for the scheduling group that the open shift belongs to. | [optional] 
**SharedOpenShift** | Pointer to [**AnyOfmicrosoftGraphOpenShiftItem**](anyOf&lt;microsoft.graph.openShiftItem&gt;.md) | A published open shift. | [optional] 

## Methods

### NewMicrosoftGraphOpenShift

`func NewMicrosoftGraphOpenShift() *MicrosoftGraphOpenShift`

NewMicrosoftGraphOpenShift instantiates a new MicrosoftGraphOpenShift object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphOpenShiftWithDefaults

`func NewMicrosoftGraphOpenShiftWithDefaults() *MicrosoftGraphOpenShift`

NewMicrosoftGraphOpenShiftWithDefaults instantiates a new MicrosoftGraphOpenShift object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphOpenShift) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphOpenShift) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphOpenShift) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphOpenShift) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedDateTime

`func (o *MicrosoftGraphOpenShift) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *MicrosoftGraphOpenShift) GetCreatedDateTimeOk() (*time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateTime

`func (o *MicrosoftGraphOpenShift) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime sets CreatedDateTime field to given value.

### HasCreatedDateTime

`func (o *MicrosoftGraphOpenShift) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTimeNil

`func (o *MicrosoftGraphOpenShift) SetCreatedDateTimeNil(b bool)`

 SetCreatedDateTimeNil sets the value for CreatedDateTime to be an explicit nil

### UnsetCreatedDateTime
`func (o *MicrosoftGraphOpenShift) UnsetCreatedDateTime()`

UnsetCreatedDateTime ensures that no value is present for CreatedDateTime, not even an explicit nil
### GetLastModifiedBy

`func (o *MicrosoftGraphOpenShift) GetLastModifiedBy() AnyOfmicrosoftGraphIdentitySet`

GetLastModifiedBy returns the LastModifiedBy field if non-nil, zero value otherwise.

### GetLastModifiedByOk

`func (o *MicrosoftGraphOpenShift) GetLastModifiedByOk() (*AnyOfmicrosoftGraphIdentitySet, bool)`

GetLastModifiedByOk returns a tuple with the LastModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedBy

`func (o *MicrosoftGraphOpenShift) SetLastModifiedBy(v AnyOfmicrosoftGraphIdentitySet)`

SetLastModifiedBy sets LastModifiedBy field to given value.

### HasLastModifiedBy

`func (o *MicrosoftGraphOpenShift) HasLastModifiedBy() bool`

HasLastModifiedBy returns a boolean if a field has been set.

### SetLastModifiedByNil

`func (o *MicrosoftGraphOpenShift) SetLastModifiedByNil(b bool)`

 SetLastModifiedByNil sets the value for LastModifiedBy to be an explicit nil

### UnsetLastModifiedBy
`func (o *MicrosoftGraphOpenShift) UnsetLastModifiedBy()`

UnsetLastModifiedBy ensures that no value is present for LastModifiedBy, not even an explicit nil
### GetLastModifiedDateTime

`func (o *MicrosoftGraphOpenShift) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *MicrosoftGraphOpenShift) GetLastModifiedDateTimeOk() (*time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedDateTime

`func (o *MicrosoftGraphOpenShift) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime sets LastModifiedDateTime field to given value.

### HasLastModifiedDateTime

`func (o *MicrosoftGraphOpenShift) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### SetLastModifiedDateTimeNil

`func (o *MicrosoftGraphOpenShift) SetLastModifiedDateTimeNil(b bool)`

 SetLastModifiedDateTimeNil sets the value for LastModifiedDateTime to be an explicit nil

### UnsetLastModifiedDateTime
`func (o *MicrosoftGraphOpenShift) UnsetLastModifiedDateTime()`

UnsetLastModifiedDateTime ensures that no value is present for LastModifiedDateTime, not even an explicit nil
### GetDraftOpenShift

`func (o *MicrosoftGraphOpenShift) GetDraftOpenShift() AnyOfmicrosoftGraphOpenShiftItem`

GetDraftOpenShift returns the DraftOpenShift field if non-nil, zero value otherwise.

### GetDraftOpenShiftOk

`func (o *MicrosoftGraphOpenShift) GetDraftOpenShiftOk() (*AnyOfmicrosoftGraphOpenShiftItem, bool)`

GetDraftOpenShiftOk returns a tuple with the DraftOpenShift field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDraftOpenShift

`func (o *MicrosoftGraphOpenShift) SetDraftOpenShift(v AnyOfmicrosoftGraphOpenShiftItem)`

SetDraftOpenShift sets DraftOpenShift field to given value.

### HasDraftOpenShift

`func (o *MicrosoftGraphOpenShift) HasDraftOpenShift() bool`

HasDraftOpenShift returns a boolean if a field has been set.

### SetDraftOpenShiftNil

`func (o *MicrosoftGraphOpenShift) SetDraftOpenShiftNil(b bool)`

 SetDraftOpenShiftNil sets the value for DraftOpenShift to be an explicit nil

### UnsetDraftOpenShift
`func (o *MicrosoftGraphOpenShift) UnsetDraftOpenShift()`

UnsetDraftOpenShift ensures that no value is present for DraftOpenShift, not even an explicit nil
### GetSchedulingGroupId

`func (o *MicrosoftGraphOpenShift) GetSchedulingGroupId() string`

GetSchedulingGroupId returns the SchedulingGroupId field if non-nil, zero value otherwise.

### GetSchedulingGroupIdOk

`func (o *MicrosoftGraphOpenShift) GetSchedulingGroupIdOk() (*string, bool)`

GetSchedulingGroupIdOk returns a tuple with the SchedulingGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedulingGroupId

`func (o *MicrosoftGraphOpenShift) SetSchedulingGroupId(v string)`

SetSchedulingGroupId sets SchedulingGroupId field to given value.

### HasSchedulingGroupId

`func (o *MicrosoftGraphOpenShift) HasSchedulingGroupId() bool`

HasSchedulingGroupId returns a boolean if a field has been set.

### SetSchedulingGroupIdNil

`func (o *MicrosoftGraphOpenShift) SetSchedulingGroupIdNil(b bool)`

 SetSchedulingGroupIdNil sets the value for SchedulingGroupId to be an explicit nil

### UnsetSchedulingGroupId
`func (o *MicrosoftGraphOpenShift) UnsetSchedulingGroupId()`

UnsetSchedulingGroupId ensures that no value is present for SchedulingGroupId, not even an explicit nil
### GetSharedOpenShift

`func (o *MicrosoftGraphOpenShift) GetSharedOpenShift() AnyOfmicrosoftGraphOpenShiftItem`

GetSharedOpenShift returns the SharedOpenShift field if non-nil, zero value otherwise.

### GetSharedOpenShiftOk

`func (o *MicrosoftGraphOpenShift) GetSharedOpenShiftOk() (*AnyOfmicrosoftGraphOpenShiftItem, bool)`

GetSharedOpenShiftOk returns a tuple with the SharedOpenShift field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedOpenShift

`func (o *MicrosoftGraphOpenShift) SetSharedOpenShift(v AnyOfmicrosoftGraphOpenShiftItem)`

SetSharedOpenShift sets SharedOpenShift field to given value.

### HasSharedOpenShift

`func (o *MicrosoftGraphOpenShift) HasSharedOpenShift() bool`

HasSharedOpenShift returns a boolean if a field has been set.

### SetSharedOpenShiftNil

`func (o *MicrosoftGraphOpenShift) SetSharedOpenShiftNil(b bool)`

 SetSharedOpenShiftNil sets the value for SharedOpenShift to be an explicit nil

### UnsetSharedOpenShift
`func (o *MicrosoftGraphOpenShift) UnsetSharedOpenShift()`

UnsetSharedOpenShift ensures that no value is present for SharedOpenShift, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


