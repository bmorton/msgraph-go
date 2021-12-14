# MicrosoftGraphActivityHistoryItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**ActiveDurationSeconds** | Pointer to **NullableInt32** | Optional. The duration of active user engagement. if not supplied, this is calculated from the startedDateTime and lastActiveDateTime. | [optional] 
**CreatedDateTime** | Pointer to **NullableTime** | Set by the server. DateTime in UTC when the object was created on the server. | [optional] 
**ExpirationDateTime** | Pointer to **NullableTime** | Optional. UTC DateTime when the historyItem will undergo hard-delete. Can be set by the client. | [optional] 
**LastActiveDateTime** | Pointer to **NullableTime** | Optional. UTC DateTime when the historyItem (activity session) was last understood as active or finished - if null, historyItem status should be Ongoing. | [optional] 
**LastModifiedDateTime** | Pointer to **NullableTime** | Set by the server. DateTime in UTC when the object was modified on the server. | [optional] 
**StartedDateTime** | Pointer to **time.Time** | Required. UTC DateTime when the historyItem (activity session) was started. Required for timeline history. | [optional] 
**Status** | Pointer to [**AnyOfmicrosoftGraphStatus**](anyOf&lt;microsoft.graph.status&gt;.md) | Set by the server. A status code used to identify valid objects. Values: active, updated, deleted, ignored. | [optional] 
**UserTimezone** | Pointer to **NullableString** | Optional. The timezone in which the user&#39;s device used to generate the activity was located at activity creation time. Values supplied as Olson IDs in order to support cross-platform representation. | [optional] 
**Activity** | Pointer to [**MicrosoftGraphUserActivity**](MicrosoftGraphUserActivity.md) |  | [optional] 

## Methods

### NewMicrosoftGraphActivityHistoryItem

`func NewMicrosoftGraphActivityHistoryItem() *MicrosoftGraphActivityHistoryItem`

NewMicrosoftGraphActivityHistoryItem instantiates a new MicrosoftGraphActivityHistoryItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphActivityHistoryItemWithDefaults

`func NewMicrosoftGraphActivityHistoryItemWithDefaults() *MicrosoftGraphActivityHistoryItem`

NewMicrosoftGraphActivityHistoryItemWithDefaults instantiates a new MicrosoftGraphActivityHistoryItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphActivityHistoryItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphActivityHistoryItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphActivityHistoryItem) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphActivityHistoryItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetActiveDurationSeconds

`func (o *MicrosoftGraphActivityHistoryItem) GetActiveDurationSeconds() int32`

GetActiveDurationSeconds returns the ActiveDurationSeconds field if non-nil, zero value otherwise.

### GetActiveDurationSecondsOk

`func (o *MicrosoftGraphActivityHistoryItem) GetActiveDurationSecondsOk() (*int32, bool)`

GetActiveDurationSecondsOk returns a tuple with the ActiveDurationSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveDurationSeconds

`func (o *MicrosoftGraphActivityHistoryItem) SetActiveDurationSeconds(v int32)`

SetActiveDurationSeconds sets ActiveDurationSeconds field to given value.

### HasActiveDurationSeconds

`func (o *MicrosoftGraphActivityHistoryItem) HasActiveDurationSeconds() bool`

HasActiveDurationSeconds returns a boolean if a field has been set.

### SetActiveDurationSecondsNil

`func (o *MicrosoftGraphActivityHistoryItem) SetActiveDurationSecondsNil(b bool)`

 SetActiveDurationSecondsNil sets the value for ActiveDurationSeconds to be an explicit nil

### UnsetActiveDurationSeconds
`func (o *MicrosoftGraphActivityHistoryItem) UnsetActiveDurationSeconds()`

UnsetActiveDurationSeconds ensures that no value is present for ActiveDurationSeconds, not even an explicit nil
### GetCreatedDateTime

`func (o *MicrosoftGraphActivityHistoryItem) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *MicrosoftGraphActivityHistoryItem) GetCreatedDateTimeOk() (*time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateTime

`func (o *MicrosoftGraphActivityHistoryItem) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime sets CreatedDateTime field to given value.

### HasCreatedDateTime

`func (o *MicrosoftGraphActivityHistoryItem) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTimeNil

`func (o *MicrosoftGraphActivityHistoryItem) SetCreatedDateTimeNil(b bool)`

 SetCreatedDateTimeNil sets the value for CreatedDateTime to be an explicit nil

### UnsetCreatedDateTime
`func (o *MicrosoftGraphActivityHistoryItem) UnsetCreatedDateTime()`

UnsetCreatedDateTime ensures that no value is present for CreatedDateTime, not even an explicit nil
### GetExpirationDateTime

`func (o *MicrosoftGraphActivityHistoryItem) GetExpirationDateTime() time.Time`

GetExpirationDateTime returns the ExpirationDateTime field if non-nil, zero value otherwise.

### GetExpirationDateTimeOk

`func (o *MicrosoftGraphActivityHistoryItem) GetExpirationDateTimeOk() (*time.Time, bool)`

GetExpirationDateTimeOk returns a tuple with the ExpirationDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDateTime

`func (o *MicrosoftGraphActivityHistoryItem) SetExpirationDateTime(v time.Time)`

SetExpirationDateTime sets ExpirationDateTime field to given value.

### HasExpirationDateTime

`func (o *MicrosoftGraphActivityHistoryItem) HasExpirationDateTime() bool`

HasExpirationDateTime returns a boolean if a field has been set.

### SetExpirationDateTimeNil

`func (o *MicrosoftGraphActivityHistoryItem) SetExpirationDateTimeNil(b bool)`

 SetExpirationDateTimeNil sets the value for ExpirationDateTime to be an explicit nil

### UnsetExpirationDateTime
`func (o *MicrosoftGraphActivityHistoryItem) UnsetExpirationDateTime()`

UnsetExpirationDateTime ensures that no value is present for ExpirationDateTime, not even an explicit nil
### GetLastActiveDateTime

`func (o *MicrosoftGraphActivityHistoryItem) GetLastActiveDateTime() time.Time`

GetLastActiveDateTime returns the LastActiveDateTime field if non-nil, zero value otherwise.

### GetLastActiveDateTimeOk

`func (o *MicrosoftGraphActivityHistoryItem) GetLastActiveDateTimeOk() (*time.Time, bool)`

GetLastActiveDateTimeOk returns a tuple with the LastActiveDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastActiveDateTime

`func (o *MicrosoftGraphActivityHistoryItem) SetLastActiveDateTime(v time.Time)`

SetLastActiveDateTime sets LastActiveDateTime field to given value.

### HasLastActiveDateTime

`func (o *MicrosoftGraphActivityHistoryItem) HasLastActiveDateTime() bool`

HasLastActiveDateTime returns a boolean if a field has been set.

### SetLastActiveDateTimeNil

`func (o *MicrosoftGraphActivityHistoryItem) SetLastActiveDateTimeNil(b bool)`

 SetLastActiveDateTimeNil sets the value for LastActiveDateTime to be an explicit nil

### UnsetLastActiveDateTime
`func (o *MicrosoftGraphActivityHistoryItem) UnsetLastActiveDateTime()`

UnsetLastActiveDateTime ensures that no value is present for LastActiveDateTime, not even an explicit nil
### GetLastModifiedDateTime

`func (o *MicrosoftGraphActivityHistoryItem) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *MicrosoftGraphActivityHistoryItem) GetLastModifiedDateTimeOk() (*time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedDateTime

`func (o *MicrosoftGraphActivityHistoryItem) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime sets LastModifiedDateTime field to given value.

### HasLastModifiedDateTime

`func (o *MicrosoftGraphActivityHistoryItem) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### SetLastModifiedDateTimeNil

`func (o *MicrosoftGraphActivityHistoryItem) SetLastModifiedDateTimeNil(b bool)`

 SetLastModifiedDateTimeNil sets the value for LastModifiedDateTime to be an explicit nil

### UnsetLastModifiedDateTime
`func (o *MicrosoftGraphActivityHistoryItem) UnsetLastModifiedDateTime()`

UnsetLastModifiedDateTime ensures that no value is present for LastModifiedDateTime, not even an explicit nil
### GetStartedDateTime

`func (o *MicrosoftGraphActivityHistoryItem) GetStartedDateTime() time.Time`

GetStartedDateTime returns the StartedDateTime field if non-nil, zero value otherwise.

### GetStartedDateTimeOk

`func (o *MicrosoftGraphActivityHistoryItem) GetStartedDateTimeOk() (*time.Time, bool)`

GetStartedDateTimeOk returns a tuple with the StartedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedDateTime

`func (o *MicrosoftGraphActivityHistoryItem) SetStartedDateTime(v time.Time)`

SetStartedDateTime sets StartedDateTime field to given value.

### HasStartedDateTime

`func (o *MicrosoftGraphActivityHistoryItem) HasStartedDateTime() bool`

HasStartedDateTime returns a boolean if a field has been set.

### GetStatus

`func (o *MicrosoftGraphActivityHistoryItem) GetStatus() AnyOfmicrosoftGraphStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MicrosoftGraphActivityHistoryItem) GetStatusOk() (*AnyOfmicrosoftGraphStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MicrosoftGraphActivityHistoryItem) SetStatus(v AnyOfmicrosoftGraphStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MicrosoftGraphActivityHistoryItem) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *MicrosoftGraphActivityHistoryItem) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *MicrosoftGraphActivityHistoryItem) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetUserTimezone

`func (o *MicrosoftGraphActivityHistoryItem) GetUserTimezone() string`

GetUserTimezone returns the UserTimezone field if non-nil, zero value otherwise.

### GetUserTimezoneOk

`func (o *MicrosoftGraphActivityHistoryItem) GetUserTimezoneOk() (*string, bool)`

GetUserTimezoneOk returns a tuple with the UserTimezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserTimezone

`func (o *MicrosoftGraphActivityHistoryItem) SetUserTimezone(v string)`

SetUserTimezone sets UserTimezone field to given value.

### HasUserTimezone

`func (o *MicrosoftGraphActivityHistoryItem) HasUserTimezone() bool`

HasUserTimezone returns a boolean if a field has been set.

### SetUserTimezoneNil

`func (o *MicrosoftGraphActivityHistoryItem) SetUserTimezoneNil(b bool)`

 SetUserTimezoneNil sets the value for UserTimezone to be an explicit nil

### UnsetUserTimezone
`func (o *MicrosoftGraphActivityHistoryItem) UnsetUserTimezone()`

UnsetUserTimezone ensures that no value is present for UserTimezone, not even an explicit nil
### GetActivity

`func (o *MicrosoftGraphActivityHistoryItem) GetActivity() MicrosoftGraphUserActivity`

GetActivity returns the Activity field if non-nil, zero value otherwise.

### GetActivityOk

`func (o *MicrosoftGraphActivityHistoryItem) GetActivityOk() (*MicrosoftGraphUserActivity, bool)`

GetActivityOk returns a tuple with the Activity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivity

`func (o *MicrosoftGraphActivityHistoryItem) SetActivity(v MicrosoftGraphUserActivity)`

SetActivity sets Activity field to given value.

### HasActivity

`func (o *MicrosoftGraphActivityHistoryItem) HasActivity() bool`

HasActivity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


