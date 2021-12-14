# MicrosoftGraphOpenShiftChangeRequest

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
**OpenShiftId** | Pointer to **NullableString** | ID for the open shift. | [optional] 

## Methods

### NewMicrosoftGraphOpenShiftChangeRequest

`func NewMicrosoftGraphOpenShiftChangeRequest() *MicrosoftGraphOpenShiftChangeRequest`

NewMicrosoftGraphOpenShiftChangeRequest instantiates a new MicrosoftGraphOpenShiftChangeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphOpenShiftChangeRequestWithDefaults

`func NewMicrosoftGraphOpenShiftChangeRequestWithDefaults() *MicrosoftGraphOpenShiftChangeRequest`

NewMicrosoftGraphOpenShiftChangeRequestWithDefaults instantiates a new MicrosoftGraphOpenShiftChangeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphOpenShiftChangeRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphOpenShiftChangeRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphOpenShiftChangeRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphOpenShiftChangeRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedDateTime

`func (o *MicrosoftGraphOpenShiftChangeRequest) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *MicrosoftGraphOpenShiftChangeRequest) GetCreatedDateTimeOk() (*time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateTime

`func (o *MicrosoftGraphOpenShiftChangeRequest) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime sets CreatedDateTime field to given value.

### HasCreatedDateTime

`func (o *MicrosoftGraphOpenShiftChangeRequest) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTimeNil

`func (o *MicrosoftGraphOpenShiftChangeRequest) SetCreatedDateTimeNil(b bool)`

 SetCreatedDateTimeNil sets the value for CreatedDateTime to be an explicit nil

### UnsetCreatedDateTime
`func (o *MicrosoftGraphOpenShiftChangeRequest) UnsetCreatedDateTime()`

UnsetCreatedDateTime ensures that no value is present for CreatedDateTime, not even an explicit nil
### GetLastModifiedBy

`func (o *MicrosoftGraphOpenShiftChangeRequest) GetLastModifiedBy() AnyOfmicrosoftGraphIdentitySet`

GetLastModifiedBy returns the LastModifiedBy field if non-nil, zero value otherwise.

### GetLastModifiedByOk

`func (o *MicrosoftGraphOpenShiftChangeRequest) GetLastModifiedByOk() (*AnyOfmicrosoftGraphIdentitySet, bool)`

GetLastModifiedByOk returns a tuple with the LastModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedBy

`func (o *MicrosoftGraphOpenShiftChangeRequest) SetLastModifiedBy(v AnyOfmicrosoftGraphIdentitySet)`

SetLastModifiedBy sets LastModifiedBy field to given value.

### HasLastModifiedBy

`func (o *MicrosoftGraphOpenShiftChangeRequest) HasLastModifiedBy() bool`

HasLastModifiedBy returns a boolean if a field has been set.

### SetLastModifiedByNil

`func (o *MicrosoftGraphOpenShiftChangeRequest) SetLastModifiedByNil(b bool)`

 SetLastModifiedByNil sets the value for LastModifiedBy to be an explicit nil

### UnsetLastModifiedBy
`func (o *MicrosoftGraphOpenShiftChangeRequest) UnsetLastModifiedBy()`

UnsetLastModifiedBy ensures that no value is present for LastModifiedBy, not even an explicit nil
### GetLastModifiedDateTime

`func (o *MicrosoftGraphOpenShiftChangeRequest) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *MicrosoftGraphOpenShiftChangeRequest) GetLastModifiedDateTimeOk() (*time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedDateTime

`func (o *MicrosoftGraphOpenShiftChangeRequest) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime sets LastModifiedDateTime field to given value.

### HasLastModifiedDateTime

`func (o *MicrosoftGraphOpenShiftChangeRequest) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### SetLastModifiedDateTimeNil

`func (o *MicrosoftGraphOpenShiftChangeRequest) SetLastModifiedDateTimeNil(b bool)`

 SetLastModifiedDateTimeNil sets the value for LastModifiedDateTime to be an explicit nil

### UnsetLastModifiedDateTime
`func (o *MicrosoftGraphOpenShiftChangeRequest) UnsetLastModifiedDateTime()`

UnsetLastModifiedDateTime ensures that no value is present for LastModifiedDateTime, not even an explicit nil
### GetAssignedTo

`func (o *MicrosoftGraphOpenShiftChangeRequest) GetAssignedTo() AnyOfmicrosoftGraphScheduleChangeRequestActor`

GetAssignedTo returns the AssignedTo field if non-nil, zero value otherwise.

### GetAssignedToOk

`func (o *MicrosoftGraphOpenShiftChangeRequest) GetAssignedToOk() (*AnyOfmicrosoftGraphScheduleChangeRequestActor, bool)`

GetAssignedToOk returns a tuple with the AssignedTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedTo

`func (o *MicrosoftGraphOpenShiftChangeRequest) SetAssignedTo(v AnyOfmicrosoftGraphScheduleChangeRequestActor)`

SetAssignedTo sets AssignedTo field to given value.

### HasAssignedTo

`func (o *MicrosoftGraphOpenShiftChangeRequest) HasAssignedTo() bool`

HasAssignedTo returns a boolean if a field has been set.

### SetAssignedToNil

`func (o *MicrosoftGraphOpenShiftChangeRequest) SetAssignedToNil(b bool)`

 SetAssignedToNil sets the value for AssignedTo to be an explicit nil

### UnsetAssignedTo
`func (o *MicrosoftGraphOpenShiftChangeRequest) UnsetAssignedTo()`

UnsetAssignedTo ensures that no value is present for AssignedTo, not even an explicit nil
### GetManagerActionDateTime

`func (o *MicrosoftGraphOpenShiftChangeRequest) GetManagerActionDateTime() time.Time`

GetManagerActionDateTime returns the ManagerActionDateTime field if non-nil, zero value otherwise.

### GetManagerActionDateTimeOk

`func (o *MicrosoftGraphOpenShiftChangeRequest) GetManagerActionDateTimeOk() (*time.Time, bool)`

GetManagerActionDateTimeOk returns a tuple with the ManagerActionDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagerActionDateTime

`func (o *MicrosoftGraphOpenShiftChangeRequest) SetManagerActionDateTime(v time.Time)`

SetManagerActionDateTime sets ManagerActionDateTime field to given value.

### HasManagerActionDateTime

`func (o *MicrosoftGraphOpenShiftChangeRequest) HasManagerActionDateTime() bool`

HasManagerActionDateTime returns a boolean if a field has been set.

### SetManagerActionDateTimeNil

`func (o *MicrosoftGraphOpenShiftChangeRequest) SetManagerActionDateTimeNil(b bool)`

 SetManagerActionDateTimeNil sets the value for ManagerActionDateTime to be an explicit nil

### UnsetManagerActionDateTime
`func (o *MicrosoftGraphOpenShiftChangeRequest) UnsetManagerActionDateTime()`

UnsetManagerActionDateTime ensures that no value is present for ManagerActionDateTime, not even an explicit nil
### GetManagerActionMessage

`func (o *MicrosoftGraphOpenShiftChangeRequest) GetManagerActionMessage() string`

GetManagerActionMessage returns the ManagerActionMessage field if non-nil, zero value otherwise.

### GetManagerActionMessageOk

`func (o *MicrosoftGraphOpenShiftChangeRequest) GetManagerActionMessageOk() (*string, bool)`

GetManagerActionMessageOk returns a tuple with the ManagerActionMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagerActionMessage

`func (o *MicrosoftGraphOpenShiftChangeRequest) SetManagerActionMessage(v string)`

SetManagerActionMessage sets ManagerActionMessage field to given value.

### HasManagerActionMessage

`func (o *MicrosoftGraphOpenShiftChangeRequest) HasManagerActionMessage() bool`

HasManagerActionMessage returns a boolean if a field has been set.

### SetManagerActionMessageNil

`func (o *MicrosoftGraphOpenShiftChangeRequest) SetManagerActionMessageNil(b bool)`

 SetManagerActionMessageNil sets the value for ManagerActionMessage to be an explicit nil

### UnsetManagerActionMessage
`func (o *MicrosoftGraphOpenShiftChangeRequest) UnsetManagerActionMessage()`

UnsetManagerActionMessage ensures that no value is present for ManagerActionMessage, not even an explicit nil
### GetManagerUserId

`func (o *MicrosoftGraphOpenShiftChangeRequest) GetManagerUserId() string`

GetManagerUserId returns the ManagerUserId field if non-nil, zero value otherwise.

### GetManagerUserIdOk

`func (o *MicrosoftGraphOpenShiftChangeRequest) GetManagerUserIdOk() (*string, bool)`

GetManagerUserIdOk returns a tuple with the ManagerUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagerUserId

`func (o *MicrosoftGraphOpenShiftChangeRequest) SetManagerUserId(v string)`

SetManagerUserId sets ManagerUserId field to given value.

### HasManagerUserId

`func (o *MicrosoftGraphOpenShiftChangeRequest) HasManagerUserId() bool`

HasManagerUserId returns a boolean if a field has been set.

### SetManagerUserIdNil

`func (o *MicrosoftGraphOpenShiftChangeRequest) SetManagerUserIdNil(b bool)`

 SetManagerUserIdNil sets the value for ManagerUserId to be an explicit nil

### UnsetManagerUserId
`func (o *MicrosoftGraphOpenShiftChangeRequest) UnsetManagerUserId()`

UnsetManagerUserId ensures that no value is present for ManagerUserId, not even an explicit nil
### GetSenderDateTime

`func (o *MicrosoftGraphOpenShiftChangeRequest) GetSenderDateTime() time.Time`

GetSenderDateTime returns the SenderDateTime field if non-nil, zero value otherwise.

### GetSenderDateTimeOk

`func (o *MicrosoftGraphOpenShiftChangeRequest) GetSenderDateTimeOk() (*time.Time, bool)`

GetSenderDateTimeOk returns a tuple with the SenderDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderDateTime

`func (o *MicrosoftGraphOpenShiftChangeRequest) SetSenderDateTime(v time.Time)`

SetSenderDateTime sets SenderDateTime field to given value.

### HasSenderDateTime

`func (o *MicrosoftGraphOpenShiftChangeRequest) HasSenderDateTime() bool`

HasSenderDateTime returns a boolean if a field has been set.

### SetSenderDateTimeNil

`func (o *MicrosoftGraphOpenShiftChangeRequest) SetSenderDateTimeNil(b bool)`

 SetSenderDateTimeNil sets the value for SenderDateTime to be an explicit nil

### UnsetSenderDateTime
`func (o *MicrosoftGraphOpenShiftChangeRequest) UnsetSenderDateTime()`

UnsetSenderDateTime ensures that no value is present for SenderDateTime, not even an explicit nil
### GetSenderMessage

`func (o *MicrosoftGraphOpenShiftChangeRequest) GetSenderMessage() string`

GetSenderMessage returns the SenderMessage field if non-nil, zero value otherwise.

### GetSenderMessageOk

`func (o *MicrosoftGraphOpenShiftChangeRequest) GetSenderMessageOk() (*string, bool)`

GetSenderMessageOk returns a tuple with the SenderMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderMessage

`func (o *MicrosoftGraphOpenShiftChangeRequest) SetSenderMessage(v string)`

SetSenderMessage sets SenderMessage field to given value.

### HasSenderMessage

`func (o *MicrosoftGraphOpenShiftChangeRequest) HasSenderMessage() bool`

HasSenderMessage returns a boolean if a field has been set.

### SetSenderMessageNil

`func (o *MicrosoftGraphOpenShiftChangeRequest) SetSenderMessageNil(b bool)`

 SetSenderMessageNil sets the value for SenderMessage to be an explicit nil

### UnsetSenderMessage
`func (o *MicrosoftGraphOpenShiftChangeRequest) UnsetSenderMessage()`

UnsetSenderMessage ensures that no value is present for SenderMessage, not even an explicit nil
### GetSenderUserId

`func (o *MicrosoftGraphOpenShiftChangeRequest) GetSenderUserId() string`

GetSenderUserId returns the SenderUserId field if non-nil, zero value otherwise.

### GetSenderUserIdOk

`func (o *MicrosoftGraphOpenShiftChangeRequest) GetSenderUserIdOk() (*string, bool)`

GetSenderUserIdOk returns a tuple with the SenderUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderUserId

`func (o *MicrosoftGraphOpenShiftChangeRequest) SetSenderUserId(v string)`

SetSenderUserId sets SenderUserId field to given value.

### HasSenderUserId

`func (o *MicrosoftGraphOpenShiftChangeRequest) HasSenderUserId() bool`

HasSenderUserId returns a boolean if a field has been set.

### SetSenderUserIdNil

`func (o *MicrosoftGraphOpenShiftChangeRequest) SetSenderUserIdNil(b bool)`

 SetSenderUserIdNil sets the value for SenderUserId to be an explicit nil

### UnsetSenderUserId
`func (o *MicrosoftGraphOpenShiftChangeRequest) UnsetSenderUserId()`

UnsetSenderUserId ensures that no value is present for SenderUserId, not even an explicit nil
### GetState

`func (o *MicrosoftGraphOpenShiftChangeRequest) GetState() AnyOfmicrosoftGraphScheduleChangeState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *MicrosoftGraphOpenShiftChangeRequest) GetStateOk() (*AnyOfmicrosoftGraphScheduleChangeState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *MicrosoftGraphOpenShiftChangeRequest) SetState(v AnyOfmicrosoftGraphScheduleChangeState)`

SetState sets State field to given value.

### HasState

`func (o *MicrosoftGraphOpenShiftChangeRequest) HasState() bool`

HasState returns a boolean if a field has been set.

### SetStateNil

`func (o *MicrosoftGraphOpenShiftChangeRequest) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *MicrosoftGraphOpenShiftChangeRequest) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil
### GetOpenShiftId

`func (o *MicrosoftGraphOpenShiftChangeRequest) GetOpenShiftId() string`

GetOpenShiftId returns the OpenShiftId field if non-nil, zero value otherwise.

### GetOpenShiftIdOk

`func (o *MicrosoftGraphOpenShiftChangeRequest) GetOpenShiftIdOk() (*string, bool)`

GetOpenShiftIdOk returns a tuple with the OpenShiftId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenShiftId

`func (o *MicrosoftGraphOpenShiftChangeRequest) SetOpenShiftId(v string)`

SetOpenShiftId sets OpenShiftId field to given value.

### HasOpenShiftId

`func (o *MicrosoftGraphOpenShiftChangeRequest) HasOpenShiftId() bool`

HasOpenShiftId returns a boolean if a field has been set.

### SetOpenShiftIdNil

`func (o *MicrosoftGraphOpenShiftChangeRequest) SetOpenShiftIdNil(b bool)`

 SetOpenShiftIdNil sets the value for OpenShiftId to be an explicit nil

### UnsetOpenShiftId
`func (o *MicrosoftGraphOpenShiftChangeRequest) UnsetOpenShiftId()`

UnsetOpenShiftId ensures that no value is present for OpenShiftId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


