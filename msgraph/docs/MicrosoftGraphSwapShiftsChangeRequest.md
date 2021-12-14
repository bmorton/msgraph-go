# MicrosoftGraphSwapShiftsChangeRequest

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
**RecipientActionDateTime** | Pointer to **NullableTime** | The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z | [optional] 
**RecipientActionMessage** | Pointer to **NullableString** | Custom message sent by recipient of the offer shift request. | [optional] 
**RecipientUserId** | Pointer to **NullableString** | User ID of the recipient of the offer shift request. | [optional] 
**SenderShiftId** | Pointer to **NullableString** | User ID of the sender of the offer shift request. | [optional] 
**RecipientShiftId** | Pointer to **NullableString** | ShiftId for the recipient user with whom the request is to swap. | [optional] 

## Methods

### NewMicrosoftGraphSwapShiftsChangeRequest

`func NewMicrosoftGraphSwapShiftsChangeRequest() *MicrosoftGraphSwapShiftsChangeRequest`

NewMicrosoftGraphSwapShiftsChangeRequest instantiates a new MicrosoftGraphSwapShiftsChangeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphSwapShiftsChangeRequestWithDefaults

`func NewMicrosoftGraphSwapShiftsChangeRequestWithDefaults() *MicrosoftGraphSwapShiftsChangeRequest`

NewMicrosoftGraphSwapShiftsChangeRequestWithDefaults instantiates a new MicrosoftGraphSwapShiftsChangeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphSwapShiftsChangeRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphSwapShiftsChangeRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphSwapShiftsChangeRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphSwapShiftsChangeRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedDateTime

`func (o *MicrosoftGraphSwapShiftsChangeRequest) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *MicrosoftGraphSwapShiftsChangeRequest) GetCreatedDateTimeOk() (*time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateTime

`func (o *MicrosoftGraphSwapShiftsChangeRequest) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime sets CreatedDateTime field to given value.

### HasCreatedDateTime

`func (o *MicrosoftGraphSwapShiftsChangeRequest) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTimeNil

`func (o *MicrosoftGraphSwapShiftsChangeRequest) SetCreatedDateTimeNil(b bool)`

 SetCreatedDateTimeNil sets the value for CreatedDateTime to be an explicit nil

### UnsetCreatedDateTime
`func (o *MicrosoftGraphSwapShiftsChangeRequest) UnsetCreatedDateTime()`

UnsetCreatedDateTime ensures that no value is present for CreatedDateTime, not even an explicit nil
### GetLastModifiedBy

`func (o *MicrosoftGraphSwapShiftsChangeRequest) GetLastModifiedBy() AnyOfmicrosoftGraphIdentitySet`

GetLastModifiedBy returns the LastModifiedBy field if non-nil, zero value otherwise.

### GetLastModifiedByOk

`func (o *MicrosoftGraphSwapShiftsChangeRequest) GetLastModifiedByOk() (*AnyOfmicrosoftGraphIdentitySet, bool)`

GetLastModifiedByOk returns a tuple with the LastModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedBy

`func (o *MicrosoftGraphSwapShiftsChangeRequest) SetLastModifiedBy(v AnyOfmicrosoftGraphIdentitySet)`

SetLastModifiedBy sets LastModifiedBy field to given value.

### HasLastModifiedBy

`func (o *MicrosoftGraphSwapShiftsChangeRequest) HasLastModifiedBy() bool`

HasLastModifiedBy returns a boolean if a field has been set.

### SetLastModifiedByNil

`func (o *MicrosoftGraphSwapShiftsChangeRequest) SetLastModifiedByNil(b bool)`

 SetLastModifiedByNil sets the value for LastModifiedBy to be an explicit nil

### UnsetLastModifiedBy
`func (o *MicrosoftGraphSwapShiftsChangeRequest) UnsetLastModifiedBy()`

UnsetLastModifiedBy ensures that no value is present for LastModifiedBy, not even an explicit nil
### GetLastModifiedDateTime

`func (o *MicrosoftGraphSwapShiftsChangeRequest) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *MicrosoftGraphSwapShiftsChangeRequest) GetLastModifiedDateTimeOk() (*time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedDateTime

`func (o *MicrosoftGraphSwapShiftsChangeRequest) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime sets LastModifiedDateTime field to given value.

### HasLastModifiedDateTime

`func (o *MicrosoftGraphSwapShiftsChangeRequest) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### SetLastModifiedDateTimeNil

`func (o *MicrosoftGraphSwapShiftsChangeRequest) SetLastModifiedDateTimeNil(b bool)`

 SetLastModifiedDateTimeNil sets the value for LastModifiedDateTime to be an explicit nil

### UnsetLastModifiedDateTime
`func (o *MicrosoftGraphSwapShiftsChangeRequest) UnsetLastModifiedDateTime()`

UnsetLastModifiedDateTime ensures that no value is present for LastModifiedDateTime, not even an explicit nil
### GetAssignedTo

`func (o *MicrosoftGraphSwapShiftsChangeRequest) GetAssignedTo() AnyOfmicrosoftGraphScheduleChangeRequestActor`

GetAssignedTo returns the AssignedTo field if non-nil, zero value otherwise.

### GetAssignedToOk

`func (o *MicrosoftGraphSwapShiftsChangeRequest) GetAssignedToOk() (*AnyOfmicrosoftGraphScheduleChangeRequestActor, bool)`

GetAssignedToOk returns a tuple with the AssignedTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedTo

`func (o *MicrosoftGraphSwapShiftsChangeRequest) SetAssignedTo(v AnyOfmicrosoftGraphScheduleChangeRequestActor)`

SetAssignedTo sets AssignedTo field to given value.

### HasAssignedTo

`func (o *MicrosoftGraphSwapShiftsChangeRequest) HasAssignedTo() bool`

HasAssignedTo returns a boolean if a field has been set.

### SetAssignedToNil

`func (o *MicrosoftGraphSwapShiftsChangeRequest) SetAssignedToNil(b bool)`

 SetAssignedToNil sets the value for AssignedTo to be an explicit nil

### UnsetAssignedTo
`func (o *MicrosoftGraphSwapShiftsChangeRequest) UnsetAssignedTo()`

UnsetAssignedTo ensures that no value is present for AssignedTo, not even an explicit nil
### GetManagerActionDateTime

`func (o *MicrosoftGraphSwapShiftsChangeRequest) GetManagerActionDateTime() time.Time`

GetManagerActionDateTime returns the ManagerActionDateTime field if non-nil, zero value otherwise.

### GetManagerActionDateTimeOk

`func (o *MicrosoftGraphSwapShiftsChangeRequest) GetManagerActionDateTimeOk() (*time.Time, bool)`

GetManagerActionDateTimeOk returns a tuple with the ManagerActionDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagerActionDateTime

`func (o *MicrosoftGraphSwapShiftsChangeRequest) SetManagerActionDateTime(v time.Time)`

SetManagerActionDateTime sets ManagerActionDateTime field to given value.

### HasManagerActionDateTime

`func (o *MicrosoftGraphSwapShiftsChangeRequest) HasManagerActionDateTime() bool`

HasManagerActionDateTime returns a boolean if a field has been set.

### SetManagerActionDateTimeNil

`func (o *MicrosoftGraphSwapShiftsChangeRequest) SetManagerActionDateTimeNil(b bool)`

 SetManagerActionDateTimeNil sets the value for ManagerActionDateTime to be an explicit nil

### UnsetManagerActionDateTime
`func (o *MicrosoftGraphSwapShiftsChangeRequest) UnsetManagerActionDateTime()`

UnsetManagerActionDateTime ensures that no value is present for ManagerActionDateTime, not even an explicit nil
### GetManagerActionMessage

`func (o *MicrosoftGraphSwapShiftsChangeRequest) GetManagerActionMessage() string`

GetManagerActionMessage returns the ManagerActionMessage field if non-nil, zero value otherwise.

### GetManagerActionMessageOk

`func (o *MicrosoftGraphSwapShiftsChangeRequest) GetManagerActionMessageOk() (*string, bool)`

GetManagerActionMessageOk returns a tuple with the ManagerActionMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagerActionMessage

`func (o *MicrosoftGraphSwapShiftsChangeRequest) SetManagerActionMessage(v string)`

SetManagerActionMessage sets ManagerActionMessage field to given value.

### HasManagerActionMessage

`func (o *MicrosoftGraphSwapShiftsChangeRequest) HasManagerActionMessage() bool`

HasManagerActionMessage returns a boolean if a field has been set.

### SetManagerActionMessageNil

`func (o *MicrosoftGraphSwapShiftsChangeRequest) SetManagerActionMessageNil(b bool)`

 SetManagerActionMessageNil sets the value for ManagerActionMessage to be an explicit nil

### UnsetManagerActionMessage
`func (o *MicrosoftGraphSwapShiftsChangeRequest) UnsetManagerActionMessage()`

UnsetManagerActionMessage ensures that no value is present for ManagerActionMessage, not even an explicit nil
### GetManagerUserId

`func (o *MicrosoftGraphSwapShiftsChangeRequest) GetManagerUserId() string`

GetManagerUserId returns the ManagerUserId field if non-nil, zero value otherwise.

### GetManagerUserIdOk

`func (o *MicrosoftGraphSwapShiftsChangeRequest) GetManagerUserIdOk() (*string, bool)`

GetManagerUserIdOk returns a tuple with the ManagerUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagerUserId

`func (o *MicrosoftGraphSwapShiftsChangeRequest) SetManagerUserId(v string)`

SetManagerUserId sets ManagerUserId field to given value.

### HasManagerUserId

`func (o *MicrosoftGraphSwapShiftsChangeRequest) HasManagerUserId() bool`

HasManagerUserId returns a boolean if a field has been set.

### SetManagerUserIdNil

`func (o *MicrosoftGraphSwapShiftsChangeRequest) SetManagerUserIdNil(b bool)`

 SetManagerUserIdNil sets the value for ManagerUserId to be an explicit nil

### UnsetManagerUserId
`func (o *MicrosoftGraphSwapShiftsChangeRequest) UnsetManagerUserId()`

UnsetManagerUserId ensures that no value is present for ManagerUserId, not even an explicit nil
### GetSenderDateTime

`func (o *MicrosoftGraphSwapShiftsChangeRequest) GetSenderDateTime() time.Time`

GetSenderDateTime returns the SenderDateTime field if non-nil, zero value otherwise.

### GetSenderDateTimeOk

`func (o *MicrosoftGraphSwapShiftsChangeRequest) GetSenderDateTimeOk() (*time.Time, bool)`

GetSenderDateTimeOk returns a tuple with the SenderDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderDateTime

`func (o *MicrosoftGraphSwapShiftsChangeRequest) SetSenderDateTime(v time.Time)`

SetSenderDateTime sets SenderDateTime field to given value.

### HasSenderDateTime

`func (o *MicrosoftGraphSwapShiftsChangeRequest) HasSenderDateTime() bool`

HasSenderDateTime returns a boolean if a field has been set.

### SetSenderDateTimeNil

`func (o *MicrosoftGraphSwapShiftsChangeRequest) SetSenderDateTimeNil(b bool)`

 SetSenderDateTimeNil sets the value for SenderDateTime to be an explicit nil

### UnsetSenderDateTime
`func (o *MicrosoftGraphSwapShiftsChangeRequest) UnsetSenderDateTime()`

UnsetSenderDateTime ensures that no value is present for SenderDateTime, not even an explicit nil
### GetSenderMessage

`func (o *MicrosoftGraphSwapShiftsChangeRequest) GetSenderMessage() string`

GetSenderMessage returns the SenderMessage field if non-nil, zero value otherwise.

### GetSenderMessageOk

`func (o *MicrosoftGraphSwapShiftsChangeRequest) GetSenderMessageOk() (*string, bool)`

GetSenderMessageOk returns a tuple with the SenderMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderMessage

`func (o *MicrosoftGraphSwapShiftsChangeRequest) SetSenderMessage(v string)`

SetSenderMessage sets SenderMessage field to given value.

### HasSenderMessage

`func (o *MicrosoftGraphSwapShiftsChangeRequest) HasSenderMessage() bool`

HasSenderMessage returns a boolean if a field has been set.

### SetSenderMessageNil

`func (o *MicrosoftGraphSwapShiftsChangeRequest) SetSenderMessageNil(b bool)`

 SetSenderMessageNil sets the value for SenderMessage to be an explicit nil

### UnsetSenderMessage
`func (o *MicrosoftGraphSwapShiftsChangeRequest) UnsetSenderMessage()`

UnsetSenderMessage ensures that no value is present for SenderMessage, not even an explicit nil
### GetSenderUserId

`func (o *MicrosoftGraphSwapShiftsChangeRequest) GetSenderUserId() string`

GetSenderUserId returns the SenderUserId field if non-nil, zero value otherwise.

### GetSenderUserIdOk

`func (o *MicrosoftGraphSwapShiftsChangeRequest) GetSenderUserIdOk() (*string, bool)`

GetSenderUserIdOk returns a tuple with the SenderUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderUserId

`func (o *MicrosoftGraphSwapShiftsChangeRequest) SetSenderUserId(v string)`

SetSenderUserId sets SenderUserId field to given value.

### HasSenderUserId

`func (o *MicrosoftGraphSwapShiftsChangeRequest) HasSenderUserId() bool`

HasSenderUserId returns a boolean if a field has been set.

### SetSenderUserIdNil

`func (o *MicrosoftGraphSwapShiftsChangeRequest) SetSenderUserIdNil(b bool)`

 SetSenderUserIdNil sets the value for SenderUserId to be an explicit nil

### UnsetSenderUserId
`func (o *MicrosoftGraphSwapShiftsChangeRequest) UnsetSenderUserId()`

UnsetSenderUserId ensures that no value is present for SenderUserId, not even an explicit nil
### GetState

`func (o *MicrosoftGraphSwapShiftsChangeRequest) GetState() AnyOfmicrosoftGraphScheduleChangeState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *MicrosoftGraphSwapShiftsChangeRequest) GetStateOk() (*AnyOfmicrosoftGraphScheduleChangeState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *MicrosoftGraphSwapShiftsChangeRequest) SetState(v AnyOfmicrosoftGraphScheduleChangeState)`

SetState sets State field to given value.

### HasState

`func (o *MicrosoftGraphSwapShiftsChangeRequest) HasState() bool`

HasState returns a boolean if a field has been set.

### SetStateNil

`func (o *MicrosoftGraphSwapShiftsChangeRequest) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *MicrosoftGraphSwapShiftsChangeRequest) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil
### GetRecipientActionDateTime

`func (o *MicrosoftGraphSwapShiftsChangeRequest) GetRecipientActionDateTime() time.Time`

GetRecipientActionDateTime returns the RecipientActionDateTime field if non-nil, zero value otherwise.

### GetRecipientActionDateTimeOk

`func (o *MicrosoftGraphSwapShiftsChangeRequest) GetRecipientActionDateTimeOk() (*time.Time, bool)`

GetRecipientActionDateTimeOk returns a tuple with the RecipientActionDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientActionDateTime

`func (o *MicrosoftGraphSwapShiftsChangeRequest) SetRecipientActionDateTime(v time.Time)`

SetRecipientActionDateTime sets RecipientActionDateTime field to given value.

### HasRecipientActionDateTime

`func (o *MicrosoftGraphSwapShiftsChangeRequest) HasRecipientActionDateTime() bool`

HasRecipientActionDateTime returns a boolean if a field has been set.

### SetRecipientActionDateTimeNil

`func (o *MicrosoftGraphSwapShiftsChangeRequest) SetRecipientActionDateTimeNil(b bool)`

 SetRecipientActionDateTimeNil sets the value for RecipientActionDateTime to be an explicit nil

### UnsetRecipientActionDateTime
`func (o *MicrosoftGraphSwapShiftsChangeRequest) UnsetRecipientActionDateTime()`

UnsetRecipientActionDateTime ensures that no value is present for RecipientActionDateTime, not even an explicit nil
### GetRecipientActionMessage

`func (o *MicrosoftGraphSwapShiftsChangeRequest) GetRecipientActionMessage() string`

GetRecipientActionMessage returns the RecipientActionMessage field if non-nil, zero value otherwise.

### GetRecipientActionMessageOk

`func (o *MicrosoftGraphSwapShiftsChangeRequest) GetRecipientActionMessageOk() (*string, bool)`

GetRecipientActionMessageOk returns a tuple with the RecipientActionMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientActionMessage

`func (o *MicrosoftGraphSwapShiftsChangeRequest) SetRecipientActionMessage(v string)`

SetRecipientActionMessage sets RecipientActionMessage field to given value.

### HasRecipientActionMessage

`func (o *MicrosoftGraphSwapShiftsChangeRequest) HasRecipientActionMessage() bool`

HasRecipientActionMessage returns a boolean if a field has been set.

### SetRecipientActionMessageNil

`func (o *MicrosoftGraphSwapShiftsChangeRequest) SetRecipientActionMessageNil(b bool)`

 SetRecipientActionMessageNil sets the value for RecipientActionMessage to be an explicit nil

### UnsetRecipientActionMessage
`func (o *MicrosoftGraphSwapShiftsChangeRequest) UnsetRecipientActionMessage()`

UnsetRecipientActionMessage ensures that no value is present for RecipientActionMessage, not even an explicit nil
### GetRecipientUserId

`func (o *MicrosoftGraphSwapShiftsChangeRequest) GetRecipientUserId() string`

GetRecipientUserId returns the RecipientUserId field if non-nil, zero value otherwise.

### GetRecipientUserIdOk

`func (o *MicrosoftGraphSwapShiftsChangeRequest) GetRecipientUserIdOk() (*string, bool)`

GetRecipientUserIdOk returns a tuple with the RecipientUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientUserId

`func (o *MicrosoftGraphSwapShiftsChangeRequest) SetRecipientUserId(v string)`

SetRecipientUserId sets RecipientUserId field to given value.

### HasRecipientUserId

`func (o *MicrosoftGraphSwapShiftsChangeRequest) HasRecipientUserId() bool`

HasRecipientUserId returns a boolean if a field has been set.

### SetRecipientUserIdNil

`func (o *MicrosoftGraphSwapShiftsChangeRequest) SetRecipientUserIdNil(b bool)`

 SetRecipientUserIdNil sets the value for RecipientUserId to be an explicit nil

### UnsetRecipientUserId
`func (o *MicrosoftGraphSwapShiftsChangeRequest) UnsetRecipientUserId()`

UnsetRecipientUserId ensures that no value is present for RecipientUserId, not even an explicit nil
### GetSenderShiftId

`func (o *MicrosoftGraphSwapShiftsChangeRequest) GetSenderShiftId() string`

GetSenderShiftId returns the SenderShiftId field if non-nil, zero value otherwise.

### GetSenderShiftIdOk

`func (o *MicrosoftGraphSwapShiftsChangeRequest) GetSenderShiftIdOk() (*string, bool)`

GetSenderShiftIdOk returns a tuple with the SenderShiftId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderShiftId

`func (o *MicrosoftGraphSwapShiftsChangeRequest) SetSenderShiftId(v string)`

SetSenderShiftId sets SenderShiftId field to given value.

### HasSenderShiftId

`func (o *MicrosoftGraphSwapShiftsChangeRequest) HasSenderShiftId() bool`

HasSenderShiftId returns a boolean if a field has been set.

### SetSenderShiftIdNil

`func (o *MicrosoftGraphSwapShiftsChangeRequest) SetSenderShiftIdNil(b bool)`

 SetSenderShiftIdNil sets the value for SenderShiftId to be an explicit nil

### UnsetSenderShiftId
`func (o *MicrosoftGraphSwapShiftsChangeRequest) UnsetSenderShiftId()`

UnsetSenderShiftId ensures that no value is present for SenderShiftId, not even an explicit nil
### GetRecipientShiftId

`func (o *MicrosoftGraphSwapShiftsChangeRequest) GetRecipientShiftId() string`

GetRecipientShiftId returns the RecipientShiftId field if non-nil, zero value otherwise.

### GetRecipientShiftIdOk

`func (o *MicrosoftGraphSwapShiftsChangeRequest) GetRecipientShiftIdOk() (*string, bool)`

GetRecipientShiftIdOk returns a tuple with the RecipientShiftId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientShiftId

`func (o *MicrosoftGraphSwapShiftsChangeRequest) SetRecipientShiftId(v string)`

SetRecipientShiftId sets RecipientShiftId field to given value.

### HasRecipientShiftId

`func (o *MicrosoftGraphSwapShiftsChangeRequest) HasRecipientShiftId() bool`

HasRecipientShiftId returns a boolean if a field has been set.

### SetRecipientShiftIdNil

`func (o *MicrosoftGraphSwapShiftsChangeRequest) SetRecipientShiftIdNil(b bool)`

 SetRecipientShiftIdNil sets the value for RecipientShiftId to be an explicit nil

### UnsetRecipientShiftId
`func (o *MicrosoftGraphSwapShiftsChangeRequest) UnsetRecipientShiftId()`

UnsetRecipientShiftId ensures that no value is present for RecipientShiftId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


