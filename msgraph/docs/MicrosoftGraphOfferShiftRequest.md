# MicrosoftGraphOfferShiftRequest

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

## Methods

### NewMicrosoftGraphOfferShiftRequest

`func NewMicrosoftGraphOfferShiftRequest() *MicrosoftGraphOfferShiftRequest`

NewMicrosoftGraphOfferShiftRequest instantiates a new MicrosoftGraphOfferShiftRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphOfferShiftRequestWithDefaults

`func NewMicrosoftGraphOfferShiftRequestWithDefaults() *MicrosoftGraphOfferShiftRequest`

NewMicrosoftGraphOfferShiftRequestWithDefaults instantiates a new MicrosoftGraphOfferShiftRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphOfferShiftRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphOfferShiftRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphOfferShiftRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphOfferShiftRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedDateTime

`func (o *MicrosoftGraphOfferShiftRequest) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *MicrosoftGraphOfferShiftRequest) GetCreatedDateTimeOk() (*time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateTime

`func (o *MicrosoftGraphOfferShiftRequest) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime sets CreatedDateTime field to given value.

### HasCreatedDateTime

`func (o *MicrosoftGraphOfferShiftRequest) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTimeNil

`func (o *MicrosoftGraphOfferShiftRequest) SetCreatedDateTimeNil(b bool)`

 SetCreatedDateTimeNil sets the value for CreatedDateTime to be an explicit nil

### UnsetCreatedDateTime
`func (o *MicrosoftGraphOfferShiftRequest) UnsetCreatedDateTime()`

UnsetCreatedDateTime ensures that no value is present for CreatedDateTime, not even an explicit nil
### GetLastModifiedBy

`func (o *MicrosoftGraphOfferShiftRequest) GetLastModifiedBy() AnyOfmicrosoftGraphIdentitySet`

GetLastModifiedBy returns the LastModifiedBy field if non-nil, zero value otherwise.

### GetLastModifiedByOk

`func (o *MicrosoftGraphOfferShiftRequest) GetLastModifiedByOk() (*AnyOfmicrosoftGraphIdentitySet, bool)`

GetLastModifiedByOk returns a tuple with the LastModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedBy

`func (o *MicrosoftGraphOfferShiftRequest) SetLastModifiedBy(v AnyOfmicrosoftGraphIdentitySet)`

SetLastModifiedBy sets LastModifiedBy field to given value.

### HasLastModifiedBy

`func (o *MicrosoftGraphOfferShiftRequest) HasLastModifiedBy() bool`

HasLastModifiedBy returns a boolean if a field has been set.

### SetLastModifiedByNil

`func (o *MicrosoftGraphOfferShiftRequest) SetLastModifiedByNil(b bool)`

 SetLastModifiedByNil sets the value for LastModifiedBy to be an explicit nil

### UnsetLastModifiedBy
`func (o *MicrosoftGraphOfferShiftRequest) UnsetLastModifiedBy()`

UnsetLastModifiedBy ensures that no value is present for LastModifiedBy, not even an explicit nil
### GetLastModifiedDateTime

`func (o *MicrosoftGraphOfferShiftRequest) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *MicrosoftGraphOfferShiftRequest) GetLastModifiedDateTimeOk() (*time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedDateTime

`func (o *MicrosoftGraphOfferShiftRequest) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime sets LastModifiedDateTime field to given value.

### HasLastModifiedDateTime

`func (o *MicrosoftGraphOfferShiftRequest) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### SetLastModifiedDateTimeNil

`func (o *MicrosoftGraphOfferShiftRequest) SetLastModifiedDateTimeNil(b bool)`

 SetLastModifiedDateTimeNil sets the value for LastModifiedDateTime to be an explicit nil

### UnsetLastModifiedDateTime
`func (o *MicrosoftGraphOfferShiftRequest) UnsetLastModifiedDateTime()`

UnsetLastModifiedDateTime ensures that no value is present for LastModifiedDateTime, not even an explicit nil
### GetAssignedTo

`func (o *MicrosoftGraphOfferShiftRequest) GetAssignedTo() AnyOfmicrosoftGraphScheduleChangeRequestActor`

GetAssignedTo returns the AssignedTo field if non-nil, zero value otherwise.

### GetAssignedToOk

`func (o *MicrosoftGraphOfferShiftRequest) GetAssignedToOk() (*AnyOfmicrosoftGraphScheduleChangeRequestActor, bool)`

GetAssignedToOk returns a tuple with the AssignedTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedTo

`func (o *MicrosoftGraphOfferShiftRequest) SetAssignedTo(v AnyOfmicrosoftGraphScheduleChangeRequestActor)`

SetAssignedTo sets AssignedTo field to given value.

### HasAssignedTo

`func (o *MicrosoftGraphOfferShiftRequest) HasAssignedTo() bool`

HasAssignedTo returns a boolean if a field has been set.

### SetAssignedToNil

`func (o *MicrosoftGraphOfferShiftRequest) SetAssignedToNil(b bool)`

 SetAssignedToNil sets the value for AssignedTo to be an explicit nil

### UnsetAssignedTo
`func (o *MicrosoftGraphOfferShiftRequest) UnsetAssignedTo()`

UnsetAssignedTo ensures that no value is present for AssignedTo, not even an explicit nil
### GetManagerActionDateTime

`func (o *MicrosoftGraphOfferShiftRequest) GetManagerActionDateTime() time.Time`

GetManagerActionDateTime returns the ManagerActionDateTime field if non-nil, zero value otherwise.

### GetManagerActionDateTimeOk

`func (o *MicrosoftGraphOfferShiftRequest) GetManagerActionDateTimeOk() (*time.Time, bool)`

GetManagerActionDateTimeOk returns a tuple with the ManagerActionDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagerActionDateTime

`func (o *MicrosoftGraphOfferShiftRequest) SetManagerActionDateTime(v time.Time)`

SetManagerActionDateTime sets ManagerActionDateTime field to given value.

### HasManagerActionDateTime

`func (o *MicrosoftGraphOfferShiftRequest) HasManagerActionDateTime() bool`

HasManagerActionDateTime returns a boolean if a field has been set.

### SetManagerActionDateTimeNil

`func (o *MicrosoftGraphOfferShiftRequest) SetManagerActionDateTimeNil(b bool)`

 SetManagerActionDateTimeNil sets the value for ManagerActionDateTime to be an explicit nil

### UnsetManagerActionDateTime
`func (o *MicrosoftGraphOfferShiftRequest) UnsetManagerActionDateTime()`

UnsetManagerActionDateTime ensures that no value is present for ManagerActionDateTime, not even an explicit nil
### GetManagerActionMessage

`func (o *MicrosoftGraphOfferShiftRequest) GetManagerActionMessage() string`

GetManagerActionMessage returns the ManagerActionMessage field if non-nil, zero value otherwise.

### GetManagerActionMessageOk

`func (o *MicrosoftGraphOfferShiftRequest) GetManagerActionMessageOk() (*string, bool)`

GetManagerActionMessageOk returns a tuple with the ManagerActionMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagerActionMessage

`func (o *MicrosoftGraphOfferShiftRequest) SetManagerActionMessage(v string)`

SetManagerActionMessage sets ManagerActionMessage field to given value.

### HasManagerActionMessage

`func (o *MicrosoftGraphOfferShiftRequest) HasManagerActionMessage() bool`

HasManagerActionMessage returns a boolean if a field has been set.

### SetManagerActionMessageNil

`func (o *MicrosoftGraphOfferShiftRequest) SetManagerActionMessageNil(b bool)`

 SetManagerActionMessageNil sets the value for ManagerActionMessage to be an explicit nil

### UnsetManagerActionMessage
`func (o *MicrosoftGraphOfferShiftRequest) UnsetManagerActionMessage()`

UnsetManagerActionMessage ensures that no value is present for ManagerActionMessage, not even an explicit nil
### GetManagerUserId

`func (o *MicrosoftGraphOfferShiftRequest) GetManagerUserId() string`

GetManagerUserId returns the ManagerUserId field if non-nil, zero value otherwise.

### GetManagerUserIdOk

`func (o *MicrosoftGraphOfferShiftRequest) GetManagerUserIdOk() (*string, bool)`

GetManagerUserIdOk returns a tuple with the ManagerUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagerUserId

`func (o *MicrosoftGraphOfferShiftRequest) SetManagerUserId(v string)`

SetManagerUserId sets ManagerUserId field to given value.

### HasManagerUserId

`func (o *MicrosoftGraphOfferShiftRequest) HasManagerUserId() bool`

HasManagerUserId returns a boolean if a field has been set.

### SetManagerUserIdNil

`func (o *MicrosoftGraphOfferShiftRequest) SetManagerUserIdNil(b bool)`

 SetManagerUserIdNil sets the value for ManagerUserId to be an explicit nil

### UnsetManagerUserId
`func (o *MicrosoftGraphOfferShiftRequest) UnsetManagerUserId()`

UnsetManagerUserId ensures that no value is present for ManagerUserId, not even an explicit nil
### GetSenderDateTime

`func (o *MicrosoftGraphOfferShiftRequest) GetSenderDateTime() time.Time`

GetSenderDateTime returns the SenderDateTime field if non-nil, zero value otherwise.

### GetSenderDateTimeOk

`func (o *MicrosoftGraphOfferShiftRequest) GetSenderDateTimeOk() (*time.Time, bool)`

GetSenderDateTimeOk returns a tuple with the SenderDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderDateTime

`func (o *MicrosoftGraphOfferShiftRequest) SetSenderDateTime(v time.Time)`

SetSenderDateTime sets SenderDateTime field to given value.

### HasSenderDateTime

`func (o *MicrosoftGraphOfferShiftRequest) HasSenderDateTime() bool`

HasSenderDateTime returns a boolean if a field has been set.

### SetSenderDateTimeNil

`func (o *MicrosoftGraphOfferShiftRequest) SetSenderDateTimeNil(b bool)`

 SetSenderDateTimeNil sets the value for SenderDateTime to be an explicit nil

### UnsetSenderDateTime
`func (o *MicrosoftGraphOfferShiftRequest) UnsetSenderDateTime()`

UnsetSenderDateTime ensures that no value is present for SenderDateTime, not even an explicit nil
### GetSenderMessage

`func (o *MicrosoftGraphOfferShiftRequest) GetSenderMessage() string`

GetSenderMessage returns the SenderMessage field if non-nil, zero value otherwise.

### GetSenderMessageOk

`func (o *MicrosoftGraphOfferShiftRequest) GetSenderMessageOk() (*string, bool)`

GetSenderMessageOk returns a tuple with the SenderMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderMessage

`func (o *MicrosoftGraphOfferShiftRequest) SetSenderMessage(v string)`

SetSenderMessage sets SenderMessage field to given value.

### HasSenderMessage

`func (o *MicrosoftGraphOfferShiftRequest) HasSenderMessage() bool`

HasSenderMessage returns a boolean if a field has been set.

### SetSenderMessageNil

`func (o *MicrosoftGraphOfferShiftRequest) SetSenderMessageNil(b bool)`

 SetSenderMessageNil sets the value for SenderMessage to be an explicit nil

### UnsetSenderMessage
`func (o *MicrosoftGraphOfferShiftRequest) UnsetSenderMessage()`

UnsetSenderMessage ensures that no value is present for SenderMessage, not even an explicit nil
### GetSenderUserId

`func (o *MicrosoftGraphOfferShiftRequest) GetSenderUserId() string`

GetSenderUserId returns the SenderUserId field if non-nil, zero value otherwise.

### GetSenderUserIdOk

`func (o *MicrosoftGraphOfferShiftRequest) GetSenderUserIdOk() (*string, bool)`

GetSenderUserIdOk returns a tuple with the SenderUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderUserId

`func (o *MicrosoftGraphOfferShiftRequest) SetSenderUserId(v string)`

SetSenderUserId sets SenderUserId field to given value.

### HasSenderUserId

`func (o *MicrosoftGraphOfferShiftRequest) HasSenderUserId() bool`

HasSenderUserId returns a boolean if a field has been set.

### SetSenderUserIdNil

`func (o *MicrosoftGraphOfferShiftRequest) SetSenderUserIdNil(b bool)`

 SetSenderUserIdNil sets the value for SenderUserId to be an explicit nil

### UnsetSenderUserId
`func (o *MicrosoftGraphOfferShiftRequest) UnsetSenderUserId()`

UnsetSenderUserId ensures that no value is present for SenderUserId, not even an explicit nil
### GetState

`func (o *MicrosoftGraphOfferShiftRequest) GetState() AnyOfmicrosoftGraphScheduleChangeState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *MicrosoftGraphOfferShiftRequest) GetStateOk() (*AnyOfmicrosoftGraphScheduleChangeState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *MicrosoftGraphOfferShiftRequest) SetState(v AnyOfmicrosoftGraphScheduleChangeState)`

SetState sets State field to given value.

### HasState

`func (o *MicrosoftGraphOfferShiftRequest) HasState() bool`

HasState returns a boolean if a field has been set.

### SetStateNil

`func (o *MicrosoftGraphOfferShiftRequest) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *MicrosoftGraphOfferShiftRequest) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil
### GetRecipientActionDateTime

`func (o *MicrosoftGraphOfferShiftRequest) GetRecipientActionDateTime() time.Time`

GetRecipientActionDateTime returns the RecipientActionDateTime field if non-nil, zero value otherwise.

### GetRecipientActionDateTimeOk

`func (o *MicrosoftGraphOfferShiftRequest) GetRecipientActionDateTimeOk() (*time.Time, bool)`

GetRecipientActionDateTimeOk returns a tuple with the RecipientActionDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientActionDateTime

`func (o *MicrosoftGraphOfferShiftRequest) SetRecipientActionDateTime(v time.Time)`

SetRecipientActionDateTime sets RecipientActionDateTime field to given value.

### HasRecipientActionDateTime

`func (o *MicrosoftGraphOfferShiftRequest) HasRecipientActionDateTime() bool`

HasRecipientActionDateTime returns a boolean if a field has been set.

### SetRecipientActionDateTimeNil

`func (o *MicrosoftGraphOfferShiftRequest) SetRecipientActionDateTimeNil(b bool)`

 SetRecipientActionDateTimeNil sets the value for RecipientActionDateTime to be an explicit nil

### UnsetRecipientActionDateTime
`func (o *MicrosoftGraphOfferShiftRequest) UnsetRecipientActionDateTime()`

UnsetRecipientActionDateTime ensures that no value is present for RecipientActionDateTime, not even an explicit nil
### GetRecipientActionMessage

`func (o *MicrosoftGraphOfferShiftRequest) GetRecipientActionMessage() string`

GetRecipientActionMessage returns the RecipientActionMessage field if non-nil, zero value otherwise.

### GetRecipientActionMessageOk

`func (o *MicrosoftGraphOfferShiftRequest) GetRecipientActionMessageOk() (*string, bool)`

GetRecipientActionMessageOk returns a tuple with the RecipientActionMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientActionMessage

`func (o *MicrosoftGraphOfferShiftRequest) SetRecipientActionMessage(v string)`

SetRecipientActionMessage sets RecipientActionMessage field to given value.

### HasRecipientActionMessage

`func (o *MicrosoftGraphOfferShiftRequest) HasRecipientActionMessage() bool`

HasRecipientActionMessage returns a boolean if a field has been set.

### SetRecipientActionMessageNil

`func (o *MicrosoftGraphOfferShiftRequest) SetRecipientActionMessageNil(b bool)`

 SetRecipientActionMessageNil sets the value for RecipientActionMessage to be an explicit nil

### UnsetRecipientActionMessage
`func (o *MicrosoftGraphOfferShiftRequest) UnsetRecipientActionMessage()`

UnsetRecipientActionMessage ensures that no value is present for RecipientActionMessage, not even an explicit nil
### GetRecipientUserId

`func (o *MicrosoftGraphOfferShiftRequest) GetRecipientUserId() string`

GetRecipientUserId returns the RecipientUserId field if non-nil, zero value otherwise.

### GetRecipientUserIdOk

`func (o *MicrosoftGraphOfferShiftRequest) GetRecipientUserIdOk() (*string, bool)`

GetRecipientUserIdOk returns a tuple with the RecipientUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientUserId

`func (o *MicrosoftGraphOfferShiftRequest) SetRecipientUserId(v string)`

SetRecipientUserId sets RecipientUserId field to given value.

### HasRecipientUserId

`func (o *MicrosoftGraphOfferShiftRequest) HasRecipientUserId() bool`

HasRecipientUserId returns a boolean if a field has been set.

### SetRecipientUserIdNil

`func (o *MicrosoftGraphOfferShiftRequest) SetRecipientUserIdNil(b bool)`

 SetRecipientUserIdNil sets the value for RecipientUserId to be an explicit nil

### UnsetRecipientUserId
`func (o *MicrosoftGraphOfferShiftRequest) UnsetRecipientUserId()`

UnsetRecipientUserId ensures that no value is present for RecipientUserId, not even an explicit nil
### GetSenderShiftId

`func (o *MicrosoftGraphOfferShiftRequest) GetSenderShiftId() string`

GetSenderShiftId returns the SenderShiftId field if non-nil, zero value otherwise.

### GetSenderShiftIdOk

`func (o *MicrosoftGraphOfferShiftRequest) GetSenderShiftIdOk() (*string, bool)`

GetSenderShiftIdOk returns a tuple with the SenderShiftId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderShiftId

`func (o *MicrosoftGraphOfferShiftRequest) SetSenderShiftId(v string)`

SetSenderShiftId sets SenderShiftId field to given value.

### HasSenderShiftId

`func (o *MicrosoftGraphOfferShiftRequest) HasSenderShiftId() bool`

HasSenderShiftId returns a boolean if a field has been set.

### SetSenderShiftIdNil

`func (o *MicrosoftGraphOfferShiftRequest) SetSenderShiftIdNil(b bool)`

 SetSenderShiftIdNil sets the value for SenderShiftId to be an explicit nil

### UnsetSenderShiftId
`func (o *MicrosoftGraphOfferShiftRequest) UnsetSenderShiftId()`

UnsetSenderShiftId ensures that no value is present for SenderShiftId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


