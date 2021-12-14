# MicrosoftGraphFollowupFlag

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompletedDateTime** | Pointer to [**AnyOfmicrosoftGraphDateTimeTimeZone**](anyOf&lt;microsoft.graph.dateTimeTimeZone&gt;.md) | The date and time that the follow-up was finished. | [optional] 
**DueDateTime** | Pointer to [**AnyOfmicrosoftGraphDateTimeTimeZone**](anyOf&lt;microsoft.graph.dateTimeTimeZone&gt;.md) | The date and time that the follow up is to be finished. Note: To set the due date, you must also specify the startDateTime; otherwise, you will get a 400 Bad Request response. | [optional] 
**FlagStatus** | Pointer to [**AnyOfmicrosoftGraphFollowupFlagStatus**](anyOf&lt;microsoft.graph.followupFlagStatus&gt;.md) | The status for follow-up for an item. Possible values are notFlagged, complete, and flagged. | [optional] 
**StartDateTime** | Pointer to [**AnyOfmicrosoftGraphDateTimeTimeZone**](anyOf&lt;microsoft.graph.dateTimeTimeZone&gt;.md) | The date and time that the follow-up is to begin. | [optional] 

## Methods

### NewMicrosoftGraphFollowupFlag

`func NewMicrosoftGraphFollowupFlag() *MicrosoftGraphFollowupFlag`

NewMicrosoftGraphFollowupFlag instantiates a new MicrosoftGraphFollowupFlag object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphFollowupFlagWithDefaults

`func NewMicrosoftGraphFollowupFlagWithDefaults() *MicrosoftGraphFollowupFlag`

NewMicrosoftGraphFollowupFlagWithDefaults instantiates a new MicrosoftGraphFollowupFlag object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompletedDateTime

`func (o *MicrosoftGraphFollowupFlag) GetCompletedDateTime() AnyOfmicrosoftGraphDateTimeTimeZone`

GetCompletedDateTime returns the CompletedDateTime field if non-nil, zero value otherwise.

### GetCompletedDateTimeOk

`func (o *MicrosoftGraphFollowupFlag) GetCompletedDateTimeOk() (*AnyOfmicrosoftGraphDateTimeTimeZone, bool)`

GetCompletedDateTimeOk returns a tuple with the CompletedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedDateTime

`func (o *MicrosoftGraphFollowupFlag) SetCompletedDateTime(v AnyOfmicrosoftGraphDateTimeTimeZone)`

SetCompletedDateTime sets CompletedDateTime field to given value.

### HasCompletedDateTime

`func (o *MicrosoftGraphFollowupFlag) HasCompletedDateTime() bool`

HasCompletedDateTime returns a boolean if a field has been set.

### SetCompletedDateTimeNil

`func (o *MicrosoftGraphFollowupFlag) SetCompletedDateTimeNil(b bool)`

 SetCompletedDateTimeNil sets the value for CompletedDateTime to be an explicit nil

### UnsetCompletedDateTime
`func (o *MicrosoftGraphFollowupFlag) UnsetCompletedDateTime()`

UnsetCompletedDateTime ensures that no value is present for CompletedDateTime, not even an explicit nil
### GetDueDateTime

`func (o *MicrosoftGraphFollowupFlag) GetDueDateTime() AnyOfmicrosoftGraphDateTimeTimeZone`

GetDueDateTime returns the DueDateTime field if non-nil, zero value otherwise.

### GetDueDateTimeOk

`func (o *MicrosoftGraphFollowupFlag) GetDueDateTimeOk() (*AnyOfmicrosoftGraphDateTimeTimeZone, bool)`

GetDueDateTimeOk returns a tuple with the DueDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDateTime

`func (o *MicrosoftGraphFollowupFlag) SetDueDateTime(v AnyOfmicrosoftGraphDateTimeTimeZone)`

SetDueDateTime sets DueDateTime field to given value.

### HasDueDateTime

`func (o *MicrosoftGraphFollowupFlag) HasDueDateTime() bool`

HasDueDateTime returns a boolean if a field has been set.

### SetDueDateTimeNil

`func (o *MicrosoftGraphFollowupFlag) SetDueDateTimeNil(b bool)`

 SetDueDateTimeNil sets the value for DueDateTime to be an explicit nil

### UnsetDueDateTime
`func (o *MicrosoftGraphFollowupFlag) UnsetDueDateTime()`

UnsetDueDateTime ensures that no value is present for DueDateTime, not even an explicit nil
### GetFlagStatus

`func (o *MicrosoftGraphFollowupFlag) GetFlagStatus() AnyOfmicrosoftGraphFollowupFlagStatus`

GetFlagStatus returns the FlagStatus field if non-nil, zero value otherwise.

### GetFlagStatusOk

`func (o *MicrosoftGraphFollowupFlag) GetFlagStatusOk() (*AnyOfmicrosoftGraphFollowupFlagStatus, bool)`

GetFlagStatusOk returns a tuple with the FlagStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlagStatus

`func (o *MicrosoftGraphFollowupFlag) SetFlagStatus(v AnyOfmicrosoftGraphFollowupFlagStatus)`

SetFlagStatus sets FlagStatus field to given value.

### HasFlagStatus

`func (o *MicrosoftGraphFollowupFlag) HasFlagStatus() bool`

HasFlagStatus returns a boolean if a field has been set.

### SetFlagStatusNil

`func (o *MicrosoftGraphFollowupFlag) SetFlagStatusNil(b bool)`

 SetFlagStatusNil sets the value for FlagStatus to be an explicit nil

### UnsetFlagStatus
`func (o *MicrosoftGraphFollowupFlag) UnsetFlagStatus()`

UnsetFlagStatus ensures that no value is present for FlagStatus, not even an explicit nil
### GetStartDateTime

`func (o *MicrosoftGraphFollowupFlag) GetStartDateTime() AnyOfmicrosoftGraphDateTimeTimeZone`

GetStartDateTime returns the StartDateTime field if non-nil, zero value otherwise.

### GetStartDateTimeOk

`func (o *MicrosoftGraphFollowupFlag) GetStartDateTimeOk() (*AnyOfmicrosoftGraphDateTimeTimeZone, bool)`

GetStartDateTimeOk returns a tuple with the StartDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDateTime

`func (o *MicrosoftGraphFollowupFlag) SetStartDateTime(v AnyOfmicrosoftGraphDateTimeTimeZone)`

SetStartDateTime sets StartDateTime field to given value.

### HasStartDateTime

`func (o *MicrosoftGraphFollowupFlag) HasStartDateTime() bool`

HasStartDateTime returns a boolean if a field has been set.

### SetStartDateTimeNil

`func (o *MicrosoftGraphFollowupFlag) SetStartDateTimeNil(b bool)`

 SetStartDateTimeNil sets the value for StartDateTime to be an explicit nil

### UnsetStartDateTime
`func (o *MicrosoftGraphFollowupFlag) UnsetStartDateTime()`

UnsetStartDateTime ensures that no value is present for StartDateTime, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


