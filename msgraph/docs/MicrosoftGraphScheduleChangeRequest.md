# MicrosoftGraphScheduleChangeRequest

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

## Methods

### NewMicrosoftGraphScheduleChangeRequest

`func NewMicrosoftGraphScheduleChangeRequest() *MicrosoftGraphScheduleChangeRequest`

NewMicrosoftGraphScheduleChangeRequest instantiates a new MicrosoftGraphScheduleChangeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphScheduleChangeRequestWithDefaults

`func NewMicrosoftGraphScheduleChangeRequestWithDefaults() *MicrosoftGraphScheduleChangeRequest`

NewMicrosoftGraphScheduleChangeRequestWithDefaults instantiates a new MicrosoftGraphScheduleChangeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphScheduleChangeRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphScheduleChangeRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphScheduleChangeRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphScheduleChangeRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedDateTime

`func (o *MicrosoftGraphScheduleChangeRequest) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *MicrosoftGraphScheduleChangeRequest) GetCreatedDateTimeOk() (*time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateTime

`func (o *MicrosoftGraphScheduleChangeRequest) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime sets CreatedDateTime field to given value.

### HasCreatedDateTime

`func (o *MicrosoftGraphScheduleChangeRequest) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTimeNil

`func (o *MicrosoftGraphScheduleChangeRequest) SetCreatedDateTimeNil(b bool)`

 SetCreatedDateTimeNil sets the value for CreatedDateTime to be an explicit nil

### UnsetCreatedDateTime
`func (o *MicrosoftGraphScheduleChangeRequest) UnsetCreatedDateTime()`

UnsetCreatedDateTime ensures that no value is present for CreatedDateTime, not even an explicit nil
### GetLastModifiedBy

`func (o *MicrosoftGraphScheduleChangeRequest) GetLastModifiedBy() AnyOfmicrosoftGraphIdentitySet`

GetLastModifiedBy returns the LastModifiedBy field if non-nil, zero value otherwise.

### GetLastModifiedByOk

`func (o *MicrosoftGraphScheduleChangeRequest) GetLastModifiedByOk() (*AnyOfmicrosoftGraphIdentitySet, bool)`

GetLastModifiedByOk returns a tuple with the LastModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedBy

`func (o *MicrosoftGraphScheduleChangeRequest) SetLastModifiedBy(v AnyOfmicrosoftGraphIdentitySet)`

SetLastModifiedBy sets LastModifiedBy field to given value.

### HasLastModifiedBy

`func (o *MicrosoftGraphScheduleChangeRequest) HasLastModifiedBy() bool`

HasLastModifiedBy returns a boolean if a field has been set.

### SetLastModifiedByNil

`func (o *MicrosoftGraphScheduleChangeRequest) SetLastModifiedByNil(b bool)`

 SetLastModifiedByNil sets the value for LastModifiedBy to be an explicit nil

### UnsetLastModifiedBy
`func (o *MicrosoftGraphScheduleChangeRequest) UnsetLastModifiedBy()`

UnsetLastModifiedBy ensures that no value is present for LastModifiedBy, not even an explicit nil
### GetLastModifiedDateTime

`func (o *MicrosoftGraphScheduleChangeRequest) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *MicrosoftGraphScheduleChangeRequest) GetLastModifiedDateTimeOk() (*time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedDateTime

`func (o *MicrosoftGraphScheduleChangeRequest) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime sets LastModifiedDateTime field to given value.

### HasLastModifiedDateTime

`func (o *MicrosoftGraphScheduleChangeRequest) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### SetLastModifiedDateTimeNil

`func (o *MicrosoftGraphScheduleChangeRequest) SetLastModifiedDateTimeNil(b bool)`

 SetLastModifiedDateTimeNil sets the value for LastModifiedDateTime to be an explicit nil

### UnsetLastModifiedDateTime
`func (o *MicrosoftGraphScheduleChangeRequest) UnsetLastModifiedDateTime()`

UnsetLastModifiedDateTime ensures that no value is present for LastModifiedDateTime, not even an explicit nil
### GetAssignedTo

`func (o *MicrosoftGraphScheduleChangeRequest) GetAssignedTo() AnyOfmicrosoftGraphScheduleChangeRequestActor`

GetAssignedTo returns the AssignedTo field if non-nil, zero value otherwise.

### GetAssignedToOk

`func (o *MicrosoftGraphScheduleChangeRequest) GetAssignedToOk() (*AnyOfmicrosoftGraphScheduleChangeRequestActor, bool)`

GetAssignedToOk returns a tuple with the AssignedTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedTo

`func (o *MicrosoftGraphScheduleChangeRequest) SetAssignedTo(v AnyOfmicrosoftGraphScheduleChangeRequestActor)`

SetAssignedTo sets AssignedTo field to given value.

### HasAssignedTo

`func (o *MicrosoftGraphScheduleChangeRequest) HasAssignedTo() bool`

HasAssignedTo returns a boolean if a field has been set.

### SetAssignedToNil

`func (o *MicrosoftGraphScheduleChangeRequest) SetAssignedToNil(b bool)`

 SetAssignedToNil sets the value for AssignedTo to be an explicit nil

### UnsetAssignedTo
`func (o *MicrosoftGraphScheduleChangeRequest) UnsetAssignedTo()`

UnsetAssignedTo ensures that no value is present for AssignedTo, not even an explicit nil
### GetManagerActionDateTime

`func (o *MicrosoftGraphScheduleChangeRequest) GetManagerActionDateTime() time.Time`

GetManagerActionDateTime returns the ManagerActionDateTime field if non-nil, zero value otherwise.

### GetManagerActionDateTimeOk

`func (o *MicrosoftGraphScheduleChangeRequest) GetManagerActionDateTimeOk() (*time.Time, bool)`

GetManagerActionDateTimeOk returns a tuple with the ManagerActionDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagerActionDateTime

`func (o *MicrosoftGraphScheduleChangeRequest) SetManagerActionDateTime(v time.Time)`

SetManagerActionDateTime sets ManagerActionDateTime field to given value.

### HasManagerActionDateTime

`func (o *MicrosoftGraphScheduleChangeRequest) HasManagerActionDateTime() bool`

HasManagerActionDateTime returns a boolean if a field has been set.

### SetManagerActionDateTimeNil

`func (o *MicrosoftGraphScheduleChangeRequest) SetManagerActionDateTimeNil(b bool)`

 SetManagerActionDateTimeNil sets the value for ManagerActionDateTime to be an explicit nil

### UnsetManagerActionDateTime
`func (o *MicrosoftGraphScheduleChangeRequest) UnsetManagerActionDateTime()`

UnsetManagerActionDateTime ensures that no value is present for ManagerActionDateTime, not even an explicit nil
### GetManagerActionMessage

`func (o *MicrosoftGraphScheduleChangeRequest) GetManagerActionMessage() string`

GetManagerActionMessage returns the ManagerActionMessage field if non-nil, zero value otherwise.

### GetManagerActionMessageOk

`func (o *MicrosoftGraphScheduleChangeRequest) GetManagerActionMessageOk() (*string, bool)`

GetManagerActionMessageOk returns a tuple with the ManagerActionMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagerActionMessage

`func (o *MicrosoftGraphScheduleChangeRequest) SetManagerActionMessage(v string)`

SetManagerActionMessage sets ManagerActionMessage field to given value.

### HasManagerActionMessage

`func (o *MicrosoftGraphScheduleChangeRequest) HasManagerActionMessage() bool`

HasManagerActionMessage returns a boolean if a field has been set.

### SetManagerActionMessageNil

`func (o *MicrosoftGraphScheduleChangeRequest) SetManagerActionMessageNil(b bool)`

 SetManagerActionMessageNil sets the value for ManagerActionMessage to be an explicit nil

### UnsetManagerActionMessage
`func (o *MicrosoftGraphScheduleChangeRequest) UnsetManagerActionMessage()`

UnsetManagerActionMessage ensures that no value is present for ManagerActionMessage, not even an explicit nil
### GetManagerUserId

`func (o *MicrosoftGraphScheduleChangeRequest) GetManagerUserId() string`

GetManagerUserId returns the ManagerUserId field if non-nil, zero value otherwise.

### GetManagerUserIdOk

`func (o *MicrosoftGraphScheduleChangeRequest) GetManagerUserIdOk() (*string, bool)`

GetManagerUserIdOk returns a tuple with the ManagerUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagerUserId

`func (o *MicrosoftGraphScheduleChangeRequest) SetManagerUserId(v string)`

SetManagerUserId sets ManagerUserId field to given value.

### HasManagerUserId

`func (o *MicrosoftGraphScheduleChangeRequest) HasManagerUserId() bool`

HasManagerUserId returns a boolean if a field has been set.

### SetManagerUserIdNil

`func (o *MicrosoftGraphScheduleChangeRequest) SetManagerUserIdNil(b bool)`

 SetManagerUserIdNil sets the value for ManagerUserId to be an explicit nil

### UnsetManagerUserId
`func (o *MicrosoftGraphScheduleChangeRequest) UnsetManagerUserId()`

UnsetManagerUserId ensures that no value is present for ManagerUserId, not even an explicit nil
### GetSenderDateTime

`func (o *MicrosoftGraphScheduleChangeRequest) GetSenderDateTime() time.Time`

GetSenderDateTime returns the SenderDateTime field if non-nil, zero value otherwise.

### GetSenderDateTimeOk

`func (o *MicrosoftGraphScheduleChangeRequest) GetSenderDateTimeOk() (*time.Time, bool)`

GetSenderDateTimeOk returns a tuple with the SenderDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderDateTime

`func (o *MicrosoftGraphScheduleChangeRequest) SetSenderDateTime(v time.Time)`

SetSenderDateTime sets SenderDateTime field to given value.

### HasSenderDateTime

`func (o *MicrosoftGraphScheduleChangeRequest) HasSenderDateTime() bool`

HasSenderDateTime returns a boolean if a field has been set.

### SetSenderDateTimeNil

`func (o *MicrosoftGraphScheduleChangeRequest) SetSenderDateTimeNil(b bool)`

 SetSenderDateTimeNil sets the value for SenderDateTime to be an explicit nil

### UnsetSenderDateTime
`func (o *MicrosoftGraphScheduleChangeRequest) UnsetSenderDateTime()`

UnsetSenderDateTime ensures that no value is present for SenderDateTime, not even an explicit nil
### GetSenderMessage

`func (o *MicrosoftGraphScheduleChangeRequest) GetSenderMessage() string`

GetSenderMessage returns the SenderMessage field if non-nil, zero value otherwise.

### GetSenderMessageOk

`func (o *MicrosoftGraphScheduleChangeRequest) GetSenderMessageOk() (*string, bool)`

GetSenderMessageOk returns a tuple with the SenderMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderMessage

`func (o *MicrosoftGraphScheduleChangeRequest) SetSenderMessage(v string)`

SetSenderMessage sets SenderMessage field to given value.

### HasSenderMessage

`func (o *MicrosoftGraphScheduleChangeRequest) HasSenderMessage() bool`

HasSenderMessage returns a boolean if a field has been set.

### SetSenderMessageNil

`func (o *MicrosoftGraphScheduleChangeRequest) SetSenderMessageNil(b bool)`

 SetSenderMessageNil sets the value for SenderMessage to be an explicit nil

### UnsetSenderMessage
`func (o *MicrosoftGraphScheduleChangeRequest) UnsetSenderMessage()`

UnsetSenderMessage ensures that no value is present for SenderMessage, not even an explicit nil
### GetSenderUserId

`func (o *MicrosoftGraphScheduleChangeRequest) GetSenderUserId() string`

GetSenderUserId returns the SenderUserId field if non-nil, zero value otherwise.

### GetSenderUserIdOk

`func (o *MicrosoftGraphScheduleChangeRequest) GetSenderUserIdOk() (*string, bool)`

GetSenderUserIdOk returns a tuple with the SenderUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderUserId

`func (o *MicrosoftGraphScheduleChangeRequest) SetSenderUserId(v string)`

SetSenderUserId sets SenderUserId field to given value.

### HasSenderUserId

`func (o *MicrosoftGraphScheduleChangeRequest) HasSenderUserId() bool`

HasSenderUserId returns a boolean if a field has been set.

### SetSenderUserIdNil

`func (o *MicrosoftGraphScheduleChangeRequest) SetSenderUserIdNil(b bool)`

 SetSenderUserIdNil sets the value for SenderUserId to be an explicit nil

### UnsetSenderUserId
`func (o *MicrosoftGraphScheduleChangeRequest) UnsetSenderUserId()`

UnsetSenderUserId ensures that no value is present for SenderUserId, not even an explicit nil
### GetState

`func (o *MicrosoftGraphScheduleChangeRequest) GetState() AnyOfmicrosoftGraphScheduleChangeState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *MicrosoftGraphScheduleChangeRequest) GetStateOk() (*AnyOfmicrosoftGraphScheduleChangeState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *MicrosoftGraphScheduleChangeRequest) SetState(v AnyOfmicrosoftGraphScheduleChangeState)`

SetState sets State field to given value.

### HasState

`func (o *MicrosoftGraphScheduleChangeRequest) HasState() bool`

HasState returns a boolean if a field has been set.

### SetStateNil

`func (o *MicrosoftGraphScheduleChangeRequest) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *MicrosoftGraphScheduleChangeRequest) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


