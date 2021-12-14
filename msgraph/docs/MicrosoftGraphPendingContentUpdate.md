# MicrosoftGraphPendingContentUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QueuedDateTime** | Pointer to **NullableTime** | Date and time the pending binary operation was queued in UTC time. Read-only. | [optional] 

## Methods

### NewMicrosoftGraphPendingContentUpdate

`func NewMicrosoftGraphPendingContentUpdate() *MicrosoftGraphPendingContentUpdate`

NewMicrosoftGraphPendingContentUpdate instantiates a new MicrosoftGraphPendingContentUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphPendingContentUpdateWithDefaults

`func NewMicrosoftGraphPendingContentUpdateWithDefaults() *MicrosoftGraphPendingContentUpdate`

NewMicrosoftGraphPendingContentUpdateWithDefaults instantiates a new MicrosoftGraphPendingContentUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQueuedDateTime

`func (o *MicrosoftGraphPendingContentUpdate) GetQueuedDateTime() time.Time`

GetQueuedDateTime returns the QueuedDateTime field if non-nil, zero value otherwise.

### GetQueuedDateTimeOk

`func (o *MicrosoftGraphPendingContentUpdate) GetQueuedDateTimeOk() (*time.Time, bool)`

GetQueuedDateTimeOk returns a tuple with the QueuedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueuedDateTime

`func (o *MicrosoftGraphPendingContentUpdate) SetQueuedDateTime(v time.Time)`

SetQueuedDateTime sets QueuedDateTime field to given value.

### HasQueuedDateTime

`func (o *MicrosoftGraphPendingContentUpdate) HasQueuedDateTime() bool`

HasQueuedDateTime returns a boolean if a field has been set.

### SetQueuedDateTimeNil

`func (o *MicrosoftGraphPendingContentUpdate) SetQueuedDateTimeNil(b bool)`

 SetQueuedDateTimeNil sets the value for QueuedDateTime to be an explicit nil

### UnsetQueuedDateTime
`func (o *MicrosoftGraphPendingContentUpdate) UnsetQueuedDateTime()`

UnsetQueuedDateTime ensures that no value is present for QueuedDateTime, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


