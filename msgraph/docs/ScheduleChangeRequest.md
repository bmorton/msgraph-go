# ScheduleChangeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssignedTo** | Pointer to [**AnyOfmicrosoftGraphScheduleChangeRequestActor**](anyOf&lt;microsoft.graph.scheduleChangeRequestActor&gt;.md) |  | [optional] 
**ManagerActionDateTime** | Pointer to **NullableTime** |  | [optional] 
**ManagerActionMessage** | Pointer to **NullableString** |  | [optional] 
**ManagerUserId** | Pointer to **NullableString** |  | [optional] 
**SenderDateTime** | Pointer to **NullableTime** |  | [optional] 
**SenderMessage** | Pointer to **NullableString** |  | [optional] 
**SenderUserId** | Pointer to **NullableString** |  | [optional] 
**State** | Pointer to [**AnyOfmicrosoftGraphScheduleChangeState**](anyOf&lt;microsoft.graph.scheduleChangeState&gt;.md) |  | [optional] 

## Methods

### NewScheduleChangeRequest

`func NewScheduleChangeRequest() *ScheduleChangeRequest`

NewScheduleChangeRequest instantiates a new ScheduleChangeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduleChangeRequestWithDefaults

`func NewScheduleChangeRequestWithDefaults() *ScheduleChangeRequest`

NewScheduleChangeRequestWithDefaults instantiates a new ScheduleChangeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssignedTo

`func (o *ScheduleChangeRequest) GetAssignedTo() AnyOfmicrosoftGraphScheduleChangeRequestActor`

GetAssignedTo returns the AssignedTo field if non-nil, zero value otherwise.

### GetAssignedToOk

`func (o *ScheduleChangeRequest) GetAssignedToOk() (*AnyOfmicrosoftGraphScheduleChangeRequestActor, bool)`

GetAssignedToOk returns a tuple with the AssignedTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedTo

`func (o *ScheduleChangeRequest) SetAssignedTo(v AnyOfmicrosoftGraphScheduleChangeRequestActor)`

SetAssignedTo sets AssignedTo field to given value.

### HasAssignedTo

`func (o *ScheduleChangeRequest) HasAssignedTo() bool`

HasAssignedTo returns a boolean if a field has been set.

### SetAssignedToNil

`func (o *ScheduleChangeRequest) SetAssignedToNil(b bool)`

 SetAssignedToNil sets the value for AssignedTo to be an explicit nil

### UnsetAssignedTo
`func (o *ScheduleChangeRequest) UnsetAssignedTo()`

UnsetAssignedTo ensures that no value is present for AssignedTo, not even an explicit nil
### GetManagerActionDateTime

`func (o *ScheduleChangeRequest) GetManagerActionDateTime() time.Time`

GetManagerActionDateTime returns the ManagerActionDateTime field if non-nil, zero value otherwise.

### GetManagerActionDateTimeOk

`func (o *ScheduleChangeRequest) GetManagerActionDateTimeOk() (*time.Time, bool)`

GetManagerActionDateTimeOk returns a tuple with the ManagerActionDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagerActionDateTime

`func (o *ScheduleChangeRequest) SetManagerActionDateTime(v time.Time)`

SetManagerActionDateTime sets ManagerActionDateTime field to given value.

### HasManagerActionDateTime

`func (o *ScheduleChangeRequest) HasManagerActionDateTime() bool`

HasManagerActionDateTime returns a boolean if a field has been set.

### SetManagerActionDateTimeNil

`func (o *ScheduleChangeRequest) SetManagerActionDateTimeNil(b bool)`

 SetManagerActionDateTimeNil sets the value for ManagerActionDateTime to be an explicit nil

### UnsetManagerActionDateTime
`func (o *ScheduleChangeRequest) UnsetManagerActionDateTime()`

UnsetManagerActionDateTime ensures that no value is present for ManagerActionDateTime, not even an explicit nil
### GetManagerActionMessage

`func (o *ScheduleChangeRequest) GetManagerActionMessage() string`

GetManagerActionMessage returns the ManagerActionMessage field if non-nil, zero value otherwise.

### GetManagerActionMessageOk

`func (o *ScheduleChangeRequest) GetManagerActionMessageOk() (*string, bool)`

GetManagerActionMessageOk returns a tuple with the ManagerActionMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagerActionMessage

`func (o *ScheduleChangeRequest) SetManagerActionMessage(v string)`

SetManagerActionMessage sets ManagerActionMessage field to given value.

### HasManagerActionMessage

`func (o *ScheduleChangeRequest) HasManagerActionMessage() bool`

HasManagerActionMessage returns a boolean if a field has been set.

### SetManagerActionMessageNil

`func (o *ScheduleChangeRequest) SetManagerActionMessageNil(b bool)`

 SetManagerActionMessageNil sets the value for ManagerActionMessage to be an explicit nil

### UnsetManagerActionMessage
`func (o *ScheduleChangeRequest) UnsetManagerActionMessage()`

UnsetManagerActionMessage ensures that no value is present for ManagerActionMessage, not even an explicit nil
### GetManagerUserId

`func (o *ScheduleChangeRequest) GetManagerUserId() string`

GetManagerUserId returns the ManagerUserId field if non-nil, zero value otherwise.

### GetManagerUserIdOk

`func (o *ScheduleChangeRequest) GetManagerUserIdOk() (*string, bool)`

GetManagerUserIdOk returns a tuple with the ManagerUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagerUserId

`func (o *ScheduleChangeRequest) SetManagerUserId(v string)`

SetManagerUserId sets ManagerUserId field to given value.

### HasManagerUserId

`func (o *ScheduleChangeRequest) HasManagerUserId() bool`

HasManagerUserId returns a boolean if a field has been set.

### SetManagerUserIdNil

`func (o *ScheduleChangeRequest) SetManagerUserIdNil(b bool)`

 SetManagerUserIdNil sets the value for ManagerUserId to be an explicit nil

### UnsetManagerUserId
`func (o *ScheduleChangeRequest) UnsetManagerUserId()`

UnsetManagerUserId ensures that no value is present for ManagerUserId, not even an explicit nil
### GetSenderDateTime

`func (o *ScheduleChangeRequest) GetSenderDateTime() time.Time`

GetSenderDateTime returns the SenderDateTime field if non-nil, zero value otherwise.

### GetSenderDateTimeOk

`func (o *ScheduleChangeRequest) GetSenderDateTimeOk() (*time.Time, bool)`

GetSenderDateTimeOk returns a tuple with the SenderDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderDateTime

`func (o *ScheduleChangeRequest) SetSenderDateTime(v time.Time)`

SetSenderDateTime sets SenderDateTime field to given value.

### HasSenderDateTime

`func (o *ScheduleChangeRequest) HasSenderDateTime() bool`

HasSenderDateTime returns a boolean if a field has been set.

### SetSenderDateTimeNil

`func (o *ScheduleChangeRequest) SetSenderDateTimeNil(b bool)`

 SetSenderDateTimeNil sets the value for SenderDateTime to be an explicit nil

### UnsetSenderDateTime
`func (o *ScheduleChangeRequest) UnsetSenderDateTime()`

UnsetSenderDateTime ensures that no value is present for SenderDateTime, not even an explicit nil
### GetSenderMessage

`func (o *ScheduleChangeRequest) GetSenderMessage() string`

GetSenderMessage returns the SenderMessage field if non-nil, zero value otherwise.

### GetSenderMessageOk

`func (o *ScheduleChangeRequest) GetSenderMessageOk() (*string, bool)`

GetSenderMessageOk returns a tuple with the SenderMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderMessage

`func (o *ScheduleChangeRequest) SetSenderMessage(v string)`

SetSenderMessage sets SenderMessage field to given value.

### HasSenderMessage

`func (o *ScheduleChangeRequest) HasSenderMessage() bool`

HasSenderMessage returns a boolean if a field has been set.

### SetSenderMessageNil

`func (o *ScheduleChangeRequest) SetSenderMessageNil(b bool)`

 SetSenderMessageNil sets the value for SenderMessage to be an explicit nil

### UnsetSenderMessage
`func (o *ScheduleChangeRequest) UnsetSenderMessage()`

UnsetSenderMessage ensures that no value is present for SenderMessage, not even an explicit nil
### GetSenderUserId

`func (o *ScheduleChangeRequest) GetSenderUserId() string`

GetSenderUserId returns the SenderUserId field if non-nil, zero value otherwise.

### GetSenderUserIdOk

`func (o *ScheduleChangeRequest) GetSenderUserIdOk() (*string, bool)`

GetSenderUserIdOk returns a tuple with the SenderUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderUserId

`func (o *ScheduleChangeRequest) SetSenderUserId(v string)`

SetSenderUserId sets SenderUserId field to given value.

### HasSenderUserId

`func (o *ScheduleChangeRequest) HasSenderUserId() bool`

HasSenderUserId returns a boolean if a field has been set.

### SetSenderUserIdNil

`func (o *ScheduleChangeRequest) SetSenderUserIdNil(b bool)`

 SetSenderUserIdNil sets the value for SenderUserId to be an explicit nil

### UnsetSenderUserId
`func (o *ScheduleChangeRequest) UnsetSenderUserId()`

UnsetSenderUserId ensures that no value is present for SenderUserId, not even an explicit nil
### GetState

`func (o *ScheduleChangeRequest) GetState() AnyOfmicrosoftGraphScheduleChangeState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ScheduleChangeRequest) GetStateOk() (*AnyOfmicrosoftGraphScheduleChangeState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ScheduleChangeRequest) SetState(v AnyOfmicrosoftGraphScheduleChangeState)`

SetState sets State field to given value.

### HasState

`func (o *ScheduleChangeRequest) HasState() bool`

HasState returns a boolean if a field has been set.

### SetStateNil

`func (o *ScheduleChangeRequest) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *ScheduleChangeRequest) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


