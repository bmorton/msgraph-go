# MicrosoftGraphReminder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChangeKey** | Pointer to **NullableString** | Identifies the version of the reminder. Every time the reminder is changed, changeKey changes as well. This allows Exchange to apply changes to the correct version of the object. | [optional] 
**EventEndTime** | Pointer to [**AnyOfmicrosoftGraphDateTimeTimeZone**](anyOf&lt;microsoft.graph.dateTimeTimeZone&gt;.md) | The date, time and time zone that the event ends. | [optional] 
**EventId** | Pointer to **NullableString** | The unique ID of the event. Read only. | [optional] 
**EventLocation** | Pointer to [**AnyOfmicrosoftGraphLocation**](anyOf&lt;microsoft.graph.location&gt;.md) | The location of the event. | [optional] 
**EventStartTime** | Pointer to [**AnyOfmicrosoftGraphDateTimeTimeZone**](anyOf&lt;microsoft.graph.dateTimeTimeZone&gt;.md) | The date, time, and time zone that the event starts. | [optional] 
**EventSubject** | Pointer to **NullableString** | The text of the event&#39;s subject line. | [optional] 
**EventWebLink** | Pointer to **NullableString** | The URL to open the event in Outlook on the web.The event will open in the browser if you are logged in to your mailbox via Outlook on the web. You will be prompted to login if you are not already logged in with the browser.This URL cannot be accessed from within an iFrame. | [optional] 
**ReminderFireTime** | Pointer to [**AnyOfmicrosoftGraphDateTimeTimeZone**](anyOf&lt;microsoft.graph.dateTimeTimeZone&gt;.md) | The date, time, and time zone that the reminder is set to occur. | [optional] 

## Methods

### NewMicrosoftGraphReminder

`func NewMicrosoftGraphReminder() *MicrosoftGraphReminder`

NewMicrosoftGraphReminder instantiates a new MicrosoftGraphReminder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphReminderWithDefaults

`func NewMicrosoftGraphReminderWithDefaults() *MicrosoftGraphReminder`

NewMicrosoftGraphReminderWithDefaults instantiates a new MicrosoftGraphReminder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChangeKey

`func (o *MicrosoftGraphReminder) GetChangeKey() string`

GetChangeKey returns the ChangeKey field if non-nil, zero value otherwise.

### GetChangeKeyOk

`func (o *MicrosoftGraphReminder) GetChangeKeyOk() (*string, bool)`

GetChangeKeyOk returns a tuple with the ChangeKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeKey

`func (o *MicrosoftGraphReminder) SetChangeKey(v string)`

SetChangeKey sets ChangeKey field to given value.

### HasChangeKey

`func (o *MicrosoftGraphReminder) HasChangeKey() bool`

HasChangeKey returns a boolean if a field has been set.

### SetChangeKeyNil

`func (o *MicrosoftGraphReminder) SetChangeKeyNil(b bool)`

 SetChangeKeyNil sets the value for ChangeKey to be an explicit nil

### UnsetChangeKey
`func (o *MicrosoftGraphReminder) UnsetChangeKey()`

UnsetChangeKey ensures that no value is present for ChangeKey, not even an explicit nil
### GetEventEndTime

`func (o *MicrosoftGraphReminder) GetEventEndTime() AnyOfmicrosoftGraphDateTimeTimeZone`

GetEventEndTime returns the EventEndTime field if non-nil, zero value otherwise.

### GetEventEndTimeOk

`func (o *MicrosoftGraphReminder) GetEventEndTimeOk() (*AnyOfmicrosoftGraphDateTimeTimeZone, bool)`

GetEventEndTimeOk returns a tuple with the EventEndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventEndTime

`func (o *MicrosoftGraphReminder) SetEventEndTime(v AnyOfmicrosoftGraphDateTimeTimeZone)`

SetEventEndTime sets EventEndTime field to given value.

### HasEventEndTime

`func (o *MicrosoftGraphReminder) HasEventEndTime() bool`

HasEventEndTime returns a boolean if a field has been set.

### SetEventEndTimeNil

`func (o *MicrosoftGraphReminder) SetEventEndTimeNil(b bool)`

 SetEventEndTimeNil sets the value for EventEndTime to be an explicit nil

### UnsetEventEndTime
`func (o *MicrosoftGraphReminder) UnsetEventEndTime()`

UnsetEventEndTime ensures that no value is present for EventEndTime, not even an explicit nil
### GetEventId

`func (o *MicrosoftGraphReminder) GetEventId() string`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *MicrosoftGraphReminder) GetEventIdOk() (*string, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *MicrosoftGraphReminder) SetEventId(v string)`

SetEventId sets EventId field to given value.

### HasEventId

`func (o *MicrosoftGraphReminder) HasEventId() bool`

HasEventId returns a boolean if a field has been set.

### SetEventIdNil

`func (o *MicrosoftGraphReminder) SetEventIdNil(b bool)`

 SetEventIdNil sets the value for EventId to be an explicit nil

### UnsetEventId
`func (o *MicrosoftGraphReminder) UnsetEventId()`

UnsetEventId ensures that no value is present for EventId, not even an explicit nil
### GetEventLocation

`func (o *MicrosoftGraphReminder) GetEventLocation() AnyOfmicrosoftGraphLocation`

GetEventLocation returns the EventLocation field if non-nil, zero value otherwise.

### GetEventLocationOk

`func (o *MicrosoftGraphReminder) GetEventLocationOk() (*AnyOfmicrosoftGraphLocation, bool)`

GetEventLocationOk returns a tuple with the EventLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventLocation

`func (o *MicrosoftGraphReminder) SetEventLocation(v AnyOfmicrosoftGraphLocation)`

SetEventLocation sets EventLocation field to given value.

### HasEventLocation

`func (o *MicrosoftGraphReminder) HasEventLocation() bool`

HasEventLocation returns a boolean if a field has been set.

### SetEventLocationNil

`func (o *MicrosoftGraphReminder) SetEventLocationNil(b bool)`

 SetEventLocationNil sets the value for EventLocation to be an explicit nil

### UnsetEventLocation
`func (o *MicrosoftGraphReminder) UnsetEventLocation()`

UnsetEventLocation ensures that no value is present for EventLocation, not even an explicit nil
### GetEventStartTime

`func (o *MicrosoftGraphReminder) GetEventStartTime() AnyOfmicrosoftGraphDateTimeTimeZone`

GetEventStartTime returns the EventStartTime field if non-nil, zero value otherwise.

### GetEventStartTimeOk

`func (o *MicrosoftGraphReminder) GetEventStartTimeOk() (*AnyOfmicrosoftGraphDateTimeTimeZone, bool)`

GetEventStartTimeOk returns a tuple with the EventStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventStartTime

`func (o *MicrosoftGraphReminder) SetEventStartTime(v AnyOfmicrosoftGraphDateTimeTimeZone)`

SetEventStartTime sets EventStartTime field to given value.

### HasEventStartTime

`func (o *MicrosoftGraphReminder) HasEventStartTime() bool`

HasEventStartTime returns a boolean if a field has been set.

### SetEventStartTimeNil

`func (o *MicrosoftGraphReminder) SetEventStartTimeNil(b bool)`

 SetEventStartTimeNil sets the value for EventStartTime to be an explicit nil

### UnsetEventStartTime
`func (o *MicrosoftGraphReminder) UnsetEventStartTime()`

UnsetEventStartTime ensures that no value is present for EventStartTime, not even an explicit nil
### GetEventSubject

`func (o *MicrosoftGraphReminder) GetEventSubject() string`

GetEventSubject returns the EventSubject field if non-nil, zero value otherwise.

### GetEventSubjectOk

`func (o *MicrosoftGraphReminder) GetEventSubjectOk() (*string, bool)`

GetEventSubjectOk returns a tuple with the EventSubject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventSubject

`func (o *MicrosoftGraphReminder) SetEventSubject(v string)`

SetEventSubject sets EventSubject field to given value.

### HasEventSubject

`func (o *MicrosoftGraphReminder) HasEventSubject() bool`

HasEventSubject returns a boolean if a field has been set.

### SetEventSubjectNil

`func (o *MicrosoftGraphReminder) SetEventSubjectNil(b bool)`

 SetEventSubjectNil sets the value for EventSubject to be an explicit nil

### UnsetEventSubject
`func (o *MicrosoftGraphReminder) UnsetEventSubject()`

UnsetEventSubject ensures that no value is present for EventSubject, not even an explicit nil
### GetEventWebLink

`func (o *MicrosoftGraphReminder) GetEventWebLink() string`

GetEventWebLink returns the EventWebLink field if non-nil, zero value otherwise.

### GetEventWebLinkOk

`func (o *MicrosoftGraphReminder) GetEventWebLinkOk() (*string, bool)`

GetEventWebLinkOk returns a tuple with the EventWebLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventWebLink

`func (o *MicrosoftGraphReminder) SetEventWebLink(v string)`

SetEventWebLink sets EventWebLink field to given value.

### HasEventWebLink

`func (o *MicrosoftGraphReminder) HasEventWebLink() bool`

HasEventWebLink returns a boolean if a field has been set.

### SetEventWebLinkNil

`func (o *MicrosoftGraphReminder) SetEventWebLinkNil(b bool)`

 SetEventWebLinkNil sets the value for EventWebLink to be an explicit nil

### UnsetEventWebLink
`func (o *MicrosoftGraphReminder) UnsetEventWebLink()`

UnsetEventWebLink ensures that no value is present for EventWebLink, not even an explicit nil
### GetReminderFireTime

`func (o *MicrosoftGraphReminder) GetReminderFireTime() AnyOfmicrosoftGraphDateTimeTimeZone`

GetReminderFireTime returns the ReminderFireTime field if non-nil, zero value otherwise.

### GetReminderFireTimeOk

`func (o *MicrosoftGraphReminder) GetReminderFireTimeOk() (*AnyOfmicrosoftGraphDateTimeTimeZone, bool)`

GetReminderFireTimeOk returns a tuple with the ReminderFireTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReminderFireTime

`func (o *MicrosoftGraphReminder) SetReminderFireTime(v AnyOfmicrosoftGraphDateTimeTimeZone)`

SetReminderFireTime sets ReminderFireTime field to given value.

### HasReminderFireTime

`func (o *MicrosoftGraphReminder) HasReminderFireTime() bool`

HasReminderFireTime returns a boolean if a field has been set.

### SetReminderFireTimeNil

`func (o *MicrosoftGraphReminder) SetReminderFireTimeNil(b bool)`

 SetReminderFireTimeNil sets the value for ReminderFireTime to be an explicit nil

### UnsetReminderFireTime
`func (o *MicrosoftGraphReminder) UnsetReminderFireTime()`

UnsetReminderFireTime ensures that no value is present for ReminderFireTime, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


