# OfferShiftRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecipientActionDateTime** | Pointer to **NullableTime** | The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z | [optional] 
**RecipientActionMessage** | Pointer to **NullableString** | Custom message sent by recipient of the offer shift request. | [optional] 
**RecipientUserId** | Pointer to **NullableString** | User ID of the recipient of the offer shift request. | [optional] 
**SenderShiftId** | Pointer to **NullableString** | User ID of the sender of the offer shift request. | [optional] 

## Methods

### NewOfferShiftRequest

`func NewOfferShiftRequest() *OfferShiftRequest`

NewOfferShiftRequest instantiates a new OfferShiftRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOfferShiftRequestWithDefaults

`func NewOfferShiftRequestWithDefaults() *OfferShiftRequest`

NewOfferShiftRequestWithDefaults instantiates a new OfferShiftRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecipientActionDateTime

`func (o *OfferShiftRequest) GetRecipientActionDateTime() time.Time`

GetRecipientActionDateTime returns the RecipientActionDateTime field if non-nil, zero value otherwise.

### GetRecipientActionDateTimeOk

`func (o *OfferShiftRequest) GetRecipientActionDateTimeOk() (*time.Time, bool)`

GetRecipientActionDateTimeOk returns a tuple with the RecipientActionDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientActionDateTime

`func (o *OfferShiftRequest) SetRecipientActionDateTime(v time.Time)`

SetRecipientActionDateTime sets RecipientActionDateTime field to given value.

### HasRecipientActionDateTime

`func (o *OfferShiftRequest) HasRecipientActionDateTime() bool`

HasRecipientActionDateTime returns a boolean if a field has been set.

### SetRecipientActionDateTimeNil

`func (o *OfferShiftRequest) SetRecipientActionDateTimeNil(b bool)`

 SetRecipientActionDateTimeNil sets the value for RecipientActionDateTime to be an explicit nil

### UnsetRecipientActionDateTime
`func (o *OfferShiftRequest) UnsetRecipientActionDateTime()`

UnsetRecipientActionDateTime ensures that no value is present for RecipientActionDateTime, not even an explicit nil
### GetRecipientActionMessage

`func (o *OfferShiftRequest) GetRecipientActionMessage() string`

GetRecipientActionMessage returns the RecipientActionMessage field if non-nil, zero value otherwise.

### GetRecipientActionMessageOk

`func (o *OfferShiftRequest) GetRecipientActionMessageOk() (*string, bool)`

GetRecipientActionMessageOk returns a tuple with the RecipientActionMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientActionMessage

`func (o *OfferShiftRequest) SetRecipientActionMessage(v string)`

SetRecipientActionMessage sets RecipientActionMessage field to given value.

### HasRecipientActionMessage

`func (o *OfferShiftRequest) HasRecipientActionMessage() bool`

HasRecipientActionMessage returns a boolean if a field has been set.

### SetRecipientActionMessageNil

`func (o *OfferShiftRequest) SetRecipientActionMessageNil(b bool)`

 SetRecipientActionMessageNil sets the value for RecipientActionMessage to be an explicit nil

### UnsetRecipientActionMessage
`func (o *OfferShiftRequest) UnsetRecipientActionMessage()`

UnsetRecipientActionMessage ensures that no value is present for RecipientActionMessage, not even an explicit nil
### GetRecipientUserId

`func (o *OfferShiftRequest) GetRecipientUserId() string`

GetRecipientUserId returns the RecipientUserId field if non-nil, zero value otherwise.

### GetRecipientUserIdOk

`func (o *OfferShiftRequest) GetRecipientUserIdOk() (*string, bool)`

GetRecipientUserIdOk returns a tuple with the RecipientUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientUserId

`func (o *OfferShiftRequest) SetRecipientUserId(v string)`

SetRecipientUserId sets RecipientUserId field to given value.

### HasRecipientUserId

`func (o *OfferShiftRequest) HasRecipientUserId() bool`

HasRecipientUserId returns a boolean if a field has been set.

### SetRecipientUserIdNil

`func (o *OfferShiftRequest) SetRecipientUserIdNil(b bool)`

 SetRecipientUserIdNil sets the value for RecipientUserId to be an explicit nil

### UnsetRecipientUserId
`func (o *OfferShiftRequest) UnsetRecipientUserId()`

UnsetRecipientUserId ensures that no value is present for RecipientUserId, not even an explicit nil
### GetSenderShiftId

`func (o *OfferShiftRequest) GetSenderShiftId() string`

GetSenderShiftId returns the SenderShiftId field if non-nil, zero value otherwise.

### GetSenderShiftIdOk

`func (o *OfferShiftRequest) GetSenderShiftIdOk() (*string, bool)`

GetSenderShiftIdOk returns a tuple with the SenderShiftId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderShiftId

`func (o *OfferShiftRequest) SetSenderShiftId(v string)`

SetSenderShiftId sets SenderShiftId field to given value.

### HasSenderShiftId

`func (o *OfferShiftRequest) HasSenderShiftId() bool`

HasSenderShiftId returns a boolean if a field has been set.

### SetSenderShiftIdNil

`func (o *OfferShiftRequest) SetSenderShiftIdNil(b bool)`

 SetSenderShiftIdNil sets the value for SenderShiftId to be an explicit nil

### UnsetSenderShiftId
`func (o *OfferShiftRequest) UnsetSenderShiftId()`

UnsetSenderShiftId ensures that no value is present for SenderShiftId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


