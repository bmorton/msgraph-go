# MicrosoftGraphTimeOffRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**CreatedDateTime** | Pointer to **NullableTime** | The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z | [optional] 
**LastModifiedBy** | Pointer to [**AnyOfmicrosoftGraphIdentitySet**](anyOf&lt;microsoft.graph.identitySet&gt;.md) | Identity of the person who last modified the entity. | [optional] 
**LastModifiedDateTime** | Pointer to **NullableTime** | The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z | [optional] 
**AssignedTo** | Pointer to [**AnyOfmicrosoftGraphScheduleChangeRequestActor**](anyOf&lt;microsoft.graph.scheduleChangeRequestActor&gt;.md) |  | [optional] 
**ManagerActionDateTime** | Pointer to **NullableTime** |  | [optional] 
**ManagerActionMessage** | Pointer to **NullableString** |  | [optional] 
**ManagerUserId** | Pointer to **NullableString** |  | [optional] 
**SenderDateTime** | Pointer to **NullableTime** |  | [optional] 
**SenderMessage** | Pointer to **NullableString** |  | [optional] 
**SenderUserId** | Pointer to **NullableString** |  | [optional] 
**State** | Pointer to [**AnyOfmicrosoftGraphScheduleChangeState**](anyOf&lt;microsoft.graph.scheduleChangeState&gt;.md) |  | [optional] 
**EndDateTime** | Pointer to **NullableTime** | The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z | [optional] 
**StartDateTime** | Pointer to **NullableTime** | The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z | [optional] 
**TimeOffReasonId** | Pointer to **NullableString** | The reason for the time off. | [optional] 

## Methods

### NewMicrosoftGraphTimeOffRequest

`func NewMicrosoftGraphTimeOffRequest() *MicrosoftGraphTimeOffRequest`

NewMicrosoftGraphTimeOffRequest instantiates a new MicrosoftGraphTimeOffRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphTimeOffRequestWithDefaults

`func NewMicrosoftGraphTimeOffRequestWithDefaults() *MicrosoftGraphTimeOffRequest`

NewMicrosoftGraphTimeOffRequestWithDefaults instantiates a new MicrosoftGraphTimeOffRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphTimeOffRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphTimeOffRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphTimeOffRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphTimeOffRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedDateTime

`func (o *MicrosoftGraphTimeOffRequest) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *MicrosoftGraphTimeOffRequest) GetCreatedDateTimeOk() (*time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateTime

`func (o *MicrosoftGraphTimeOffRequest) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime sets CreatedDateTime field to given value.

### HasCreatedDateTime

`func (o *MicrosoftGraphTimeOffRequest) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTimeNil

`func (o *MicrosoftGraphTimeOffRequest) SetCreatedDateTimeNil(b bool)`

 SetCreatedDateTimeNil sets the value for CreatedDateTime to be an explicit nil

### UnsetCreatedDateTime
`func (o *MicrosoftGraphTimeOffRequest) UnsetCreatedDateTime()`

UnsetCreatedDateTime ensures that no value is present for CreatedDateTime, not even an explicit nil
### GetLastModifiedBy

`func (o *MicrosoftGraphTimeOffRequest) GetLastModifiedBy() AnyOfmicrosoftGraphIdentitySet`

GetLastModifiedBy returns the LastModifiedBy field if non-nil, zero value otherwise.

### GetLastModifiedByOk

`func (o *MicrosoftGraphTimeOffRequest) GetLastModifiedByOk() (*AnyOfmicrosoftGraphIdentitySet, bool)`

GetLastModifiedByOk returns a tuple with the LastModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedBy

`func (o *MicrosoftGraphTimeOffRequest) SetLastModifiedBy(v AnyOfmicrosoftGraphIdentitySet)`

SetLastModifiedBy sets LastModifiedBy field to given value.

### HasLastModifiedBy

`func (o *MicrosoftGraphTimeOffRequest) HasLastModifiedBy() bool`

HasLastModifiedBy returns a boolean if a field has been set.

### SetLastModifiedByNil

`func (o *MicrosoftGraphTimeOffRequest) SetLastModifiedByNil(b bool)`

 SetLastModifiedByNil sets the value for LastModifiedBy to be an explicit nil

### UnsetLastModifiedBy
`func (o *MicrosoftGraphTimeOffRequest) UnsetLastModifiedBy()`

UnsetLastModifiedBy ensures that no value is present for LastModifiedBy, not even an explicit nil
### GetLastModifiedDateTime

`func (o *MicrosoftGraphTimeOffRequest) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *MicrosoftGraphTimeOffRequest) GetLastModifiedDateTimeOk() (*time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedDateTime

`func (o *MicrosoftGraphTimeOffRequest) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime sets LastModifiedDateTime field to given value.

### HasLastModifiedDateTime

`func (o *MicrosoftGraphTimeOffRequest) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### SetLastModifiedDateTimeNil

`func (o *MicrosoftGraphTimeOffRequest) SetLastModifiedDateTimeNil(b bool)`

 SetLastModifiedDateTimeNil sets the value for LastModifiedDateTime to be an explicit nil

### UnsetLastModifiedDateTime
`func (o *MicrosoftGraphTimeOffRequest) UnsetLastModifiedDateTime()`

UnsetLastModifiedDateTime ensures that no value is present for LastModifiedDateTime, not even an explicit nil
### GetAssignedTo

`func (o *MicrosoftGraphTimeOffRequest) GetAssignedTo() AnyOfmicrosoftGraphScheduleChangeRequestActor`

GetAssignedTo returns the AssignedTo field if non-nil, zero value otherwise.

### GetAssignedToOk

`func (o *MicrosoftGraphTimeOffRequest) GetAssignedToOk() (*AnyOfmicrosoftGraphScheduleChangeRequestActor, bool)`

GetAssignedToOk returns a tuple with the AssignedTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedTo

`func (o *MicrosoftGraphTimeOffRequest) SetAssignedTo(v AnyOfmicrosoftGraphScheduleChangeRequestActor)`

SetAssignedTo sets AssignedTo field to given value.

### HasAssignedTo

`func (o *MicrosoftGraphTimeOffRequest) HasAssignedTo() bool`

HasAssignedTo returns a boolean if a field has been set.

### SetAssignedToNil

`func (o *MicrosoftGraphTimeOffRequest) SetAssignedToNil(b bool)`

 SetAssignedToNil sets the value for AssignedTo to be an explicit nil

### UnsetAssignedTo
`func (o *MicrosoftGraphTimeOffRequest) UnsetAssignedTo()`

UnsetAssignedTo ensures that no value is present for AssignedTo, not even an explicit nil
### GetManagerActionDateTime

`func (o *MicrosoftGraphTimeOffRequest) GetManagerActionDateTime() time.Time`

GetManagerActionDateTime returns the ManagerActionDateTime field if non-nil, zero value otherwise.

### GetManagerActionDateTimeOk

`func (o *MicrosoftGraphTimeOffRequest) GetManagerActionDateTimeOk() (*time.Time, bool)`

GetManagerActionDateTimeOk returns a tuple with the ManagerActionDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagerActionDateTime

`func (o *MicrosoftGraphTimeOffRequest) SetManagerActionDateTime(v time.Time)`

SetManagerActionDateTime sets ManagerActionDateTime field to given value.

### HasManagerActionDateTime

`func (o *MicrosoftGraphTimeOffRequest) HasManagerActionDateTime() bool`

HasManagerActionDateTime returns a boolean if a field has been set.

### SetManagerActionDateTimeNil

`func (o *MicrosoftGraphTimeOffRequest) SetManagerActionDateTimeNil(b bool)`

 SetManagerActionDateTimeNil sets the value for ManagerActionDateTime to be an explicit nil

### UnsetManagerActionDateTime
`func (o *MicrosoftGraphTimeOffRequest) UnsetManagerActionDateTime()`

UnsetManagerActionDateTime ensures that no value is present for ManagerActionDateTime, not even an explicit nil
### GetManagerActionMessage

`func (o *MicrosoftGraphTimeOffRequest) GetManagerActionMessage() string`

GetManagerActionMessage returns the ManagerActionMessage field if non-nil, zero value otherwise.

### GetManagerActionMessageOk

`func (o *MicrosoftGraphTimeOffRequest) GetManagerActionMessageOk() (*string, bool)`

GetManagerActionMessageOk returns a tuple with the ManagerActionMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagerActionMessage

`func (o *MicrosoftGraphTimeOffRequest) SetManagerActionMessage(v string)`

SetManagerActionMessage sets ManagerActionMessage field to given value.

### HasManagerActionMessage

`func (o *MicrosoftGraphTimeOffRequest) HasManagerActionMessage() bool`

HasManagerActionMessage returns a boolean if a field has been set.

### SetManagerActionMessageNil

`func (o *MicrosoftGraphTimeOffRequest) SetManagerActionMessageNil(b bool)`

 SetManagerActionMessageNil sets the value for ManagerActionMessage to be an explicit nil

### UnsetManagerActionMessage
`func (o *MicrosoftGraphTimeOffRequest) UnsetManagerActionMessage()`

UnsetManagerActionMessage ensures that no value is present for ManagerActionMessage, not even an explicit nil
### GetManagerUserId

`func (o *MicrosoftGraphTimeOffRequest) GetManagerUserId() string`

GetManagerUserId returns the ManagerUserId field if non-nil, zero value otherwise.

### GetManagerUserIdOk

`func (o *MicrosoftGraphTimeOffRequest) GetManagerUserIdOk() (*string, bool)`

GetManagerUserIdOk returns a tuple with the ManagerUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagerUserId

`func (o *MicrosoftGraphTimeOffRequest) SetManagerUserId(v string)`

SetManagerUserId sets ManagerUserId field to given value.

### HasManagerUserId

`func (o *MicrosoftGraphTimeOffRequest) HasManagerUserId() bool`

HasManagerUserId returns a boolean if a field has been set.

### SetManagerUserIdNil

`func (o *MicrosoftGraphTimeOffRequest) SetManagerUserIdNil(b bool)`

 SetManagerUserIdNil sets the value for ManagerUserId to be an explicit nil

### UnsetManagerUserId
`func (o *MicrosoftGraphTimeOffRequest) UnsetManagerUserId()`

UnsetManagerUserId ensures that no value is present for ManagerUserId, not even an explicit nil
### GetSenderDateTime

`func (o *MicrosoftGraphTimeOffRequest) GetSenderDateTime() time.Time`

GetSenderDateTime returns the SenderDateTime field if non-nil, zero value otherwise.

### GetSenderDateTimeOk

`func (o *MicrosoftGraphTimeOffRequest) GetSenderDateTimeOk() (*time.Time, bool)`

GetSenderDateTimeOk returns a tuple with the SenderDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderDateTime

`func (o *MicrosoftGraphTimeOffRequest) SetSenderDateTime(v time.Time)`

SetSenderDateTime sets SenderDateTime field to given value.

### HasSenderDateTime

`func (o *MicrosoftGraphTimeOffRequest) HasSenderDateTime() bool`

HasSenderDateTime returns a boolean if a field has been set.

### SetSenderDateTimeNil

`func (o *MicrosoftGraphTimeOffRequest) SetSenderDateTimeNil(b bool)`

 SetSenderDateTimeNil sets the value for SenderDateTime to be an explicit nil

### UnsetSenderDateTime
`func (o *MicrosoftGraphTimeOffRequest) UnsetSenderDateTime()`

UnsetSenderDateTime ensures that no value is present for SenderDateTime, not even an explicit nil
### GetSenderMessage

`func (o *MicrosoftGraphTimeOffRequest) GetSenderMessage() string`

GetSenderMessage returns the SenderMessage field if non-nil, zero value otherwise.

### GetSenderMessageOk

`func (o *MicrosoftGraphTimeOffRequest) GetSenderMessageOk() (*string, bool)`

GetSenderMessageOk returns a tuple with the SenderMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderMessage

`func (o *MicrosoftGraphTimeOffRequest) SetSenderMessage(v string)`

SetSenderMessage sets SenderMessage field to given value.

### HasSenderMessage

`func (o *MicrosoftGraphTimeOffRequest) HasSenderMessage() bool`

HasSenderMessage returns a boolean if a field has been set.

### SetSenderMessageNil

`func (o *MicrosoftGraphTimeOffRequest) SetSenderMessageNil(b bool)`

 SetSenderMessageNil sets the value for SenderMessage to be an explicit nil

### UnsetSenderMessage
`func (o *MicrosoftGraphTimeOffRequest) UnsetSenderMessage()`

UnsetSenderMessage ensures that no value is present for SenderMessage, not even an explicit nil
### GetSenderUserId

`func (o *MicrosoftGraphTimeOffRequest) GetSenderUserId() string`

GetSenderUserId returns the SenderUserId field if non-nil, zero value otherwise.

### GetSenderUserIdOk

`func (o *MicrosoftGraphTimeOffRequest) GetSenderUserIdOk() (*string, bool)`

GetSenderUserIdOk returns a tuple with the SenderUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderUserId

`func (o *MicrosoftGraphTimeOffRequest) SetSenderUserId(v string)`

SetSenderUserId sets SenderUserId field to given value.

### HasSenderUserId

`func (o *MicrosoftGraphTimeOffRequest) HasSenderUserId() bool`

HasSenderUserId returns a boolean if a field has been set.

### SetSenderUserIdNil

`func (o *MicrosoftGraphTimeOffRequest) SetSenderUserIdNil(b bool)`

 SetSenderUserIdNil sets the value for SenderUserId to be an explicit nil

### UnsetSenderUserId
`func (o *MicrosoftGraphTimeOffRequest) UnsetSenderUserId()`

UnsetSenderUserId ensures that no value is present for SenderUserId, not even an explicit nil
### GetState

`func (o *MicrosoftGraphTimeOffRequest) GetState() AnyOfmicrosoftGraphScheduleChangeState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *MicrosoftGraphTimeOffRequest) GetStateOk() (*AnyOfmicrosoftGraphScheduleChangeState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *MicrosoftGraphTimeOffRequest) SetState(v AnyOfmicrosoftGraphScheduleChangeState)`

SetState sets State field to given value.

### HasState

`func (o *MicrosoftGraphTimeOffRequest) HasState() bool`

HasState returns a boolean if a field has been set.

### SetStateNil

`func (o *MicrosoftGraphTimeOffRequest) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *MicrosoftGraphTimeOffRequest) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil
### GetEndDateTime

`func (o *MicrosoftGraphTimeOffRequest) GetEndDateTime() time.Time`

GetEndDateTime returns the EndDateTime field if non-nil, zero value otherwise.

### GetEndDateTimeOk

`func (o *MicrosoftGraphTimeOffRequest) GetEndDateTimeOk() (*time.Time, bool)`

GetEndDateTimeOk returns a tuple with the EndDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDateTime

`func (o *MicrosoftGraphTimeOffRequest) SetEndDateTime(v time.Time)`

SetEndDateTime sets EndDateTime field to given value.

### HasEndDateTime

`func (o *MicrosoftGraphTimeOffRequest) HasEndDateTime() bool`

HasEndDateTime returns a boolean if a field has been set.

### SetEndDateTimeNil

`func (o *MicrosoftGraphTimeOffRequest) SetEndDateTimeNil(b bool)`

 SetEndDateTimeNil sets the value for EndDateTime to be an explicit nil

### UnsetEndDateTime
`func (o *MicrosoftGraphTimeOffRequest) UnsetEndDateTime()`

UnsetEndDateTime ensures that no value is present for EndDateTime, not even an explicit nil
### GetStartDateTime

`func (o *MicrosoftGraphTimeOffRequest) GetStartDateTime() time.Time`

GetStartDateTime returns the StartDateTime field if non-nil, zero value otherwise.

### GetStartDateTimeOk

`func (o *MicrosoftGraphTimeOffRequest) GetStartDateTimeOk() (*time.Time, bool)`

GetStartDateTimeOk returns a tuple with the StartDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDateTime

`func (o *MicrosoftGraphTimeOffRequest) SetStartDateTime(v time.Time)`

SetStartDateTime sets StartDateTime field to given value.

### HasStartDateTime

`func (o *MicrosoftGraphTimeOffRequest) HasStartDateTime() bool`

HasStartDateTime returns a boolean if a field has been set.

### SetStartDateTimeNil

`func (o *MicrosoftGraphTimeOffRequest) SetStartDateTimeNil(b bool)`

 SetStartDateTimeNil sets the value for StartDateTime to be an explicit nil

### UnsetStartDateTime
`func (o *MicrosoftGraphTimeOffRequest) UnsetStartDateTime()`

UnsetStartDateTime ensures that no value is present for StartDateTime, not even an explicit nil
### GetTimeOffReasonId

`func (o *MicrosoftGraphTimeOffRequest) GetTimeOffReasonId() string`

GetTimeOffReasonId returns the TimeOffReasonId field if non-nil, zero value otherwise.

### GetTimeOffReasonIdOk

`func (o *MicrosoftGraphTimeOffRequest) GetTimeOffReasonIdOk() (*string, bool)`

GetTimeOffReasonIdOk returns a tuple with the TimeOffReasonId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeOffReasonId

`func (o *MicrosoftGraphTimeOffRequest) SetTimeOffReasonId(v string)`

SetTimeOffReasonId sets TimeOffReasonId field to given value.

### HasTimeOffReasonId

`func (o *MicrosoftGraphTimeOffRequest) HasTimeOffReasonId() bool`

HasTimeOffReasonId returns a boolean if a field has been set.

### SetTimeOffReasonIdNil

`func (o *MicrosoftGraphTimeOffRequest) SetTimeOffReasonIdNil(b bool)`

 SetTimeOffReasonIdNil sets the value for TimeOffReasonId to be an explicit nil

### UnsetTimeOffReasonId
`func (o *MicrosoftGraphTimeOffRequest) UnsetTimeOffReasonId()`

UnsetTimeOffReasonId ensures that no value is present for TimeOffReasonId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


