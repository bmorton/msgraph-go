# MicrosoftGraphEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**Categories** | Pointer to **[]string** | The categories associated with the item | [optional] 
**ChangeKey** | Pointer to **NullableString** | Identifies the version of the item. Every time the item is changed, changeKey changes as well. This allows Exchange to apply changes to the correct version of the object. Read-only. | [optional] 
**CreatedDateTime** | Pointer to **NullableTime** | The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z | [optional] 
**LastModifiedDateTime** | Pointer to **NullableTime** | The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z | [optional] 
**AllowNewTimeProposals** | Pointer to **NullableBool** | true if the meeting organizer allows invitees to propose a new time when responding; otherwise, false. Optional. Default is true. | [optional] 
**Attendees** | Pointer to [**[]AnyOfmicrosoftGraphAttendee**](AnyOfmicrosoftGraphAttendee.md) | The collection of attendees for the event. | [optional] 
**Body** | Pointer to [**AnyOfmicrosoftGraphItemBody**](anyOf&lt;microsoft.graph.itemBody&gt;.md) | The body of the message associated with the event. It can be in HTML or text format. | [optional] 
**BodyPreview** | Pointer to **NullableString** | The preview of the message associated with the event. It is in text format. | [optional] 
**End** | Pointer to [**AnyOfmicrosoftGraphDateTimeTimeZone**](anyOf&lt;microsoft.graph.dateTimeTimeZone&gt;.md) | The date, time, and time zone that the event ends. By default, the end time is in UTC. | [optional] 
**HasAttachments** | Pointer to **NullableBool** | Set to true if the event has attachments. | [optional] 
**HideAttendees** | Pointer to **NullableBool** | When set to true, each attendee only sees themselves in the meeting request and meeting Tracking list. Default is false. | [optional] 
**ICalUId** | Pointer to **NullableString** | A unique identifier for an event across calendars. This ID is different for each occurrence in a recurring series. Read-only. | [optional] 
**Importance** | Pointer to [**AnyOfmicrosoftGraphImportance**](anyOf&lt;microsoft.graph.importance&gt;.md) |  | [optional] 
**IsAllDay** | Pointer to **NullableBool** |  | [optional] 
**IsCancelled** | Pointer to **NullableBool** |  | [optional] 
**IsDraft** | Pointer to **NullableBool** |  | [optional] 
**IsOnlineMeeting** | Pointer to **NullableBool** |  | [optional] 
**IsOrganizer** | Pointer to **NullableBool** |  | [optional] 
**IsReminderOn** | Pointer to **NullableBool** |  | [optional] 
**Location** | Pointer to [**AnyOfmicrosoftGraphLocation**](anyOf&lt;microsoft.graph.location&gt;.md) |  | [optional] 
**Locations** | Pointer to [**[]AnyOfmicrosoftGraphLocation**](AnyOfmicrosoftGraphLocation.md) |  | [optional] 
**OnlineMeeting** | Pointer to [**AnyOfmicrosoftGraphOnlineMeetingInfo**](anyOf&lt;microsoft.graph.onlineMeetingInfo&gt;.md) |  | [optional] 
**OnlineMeetingProvider** | Pointer to [**AnyOfmicrosoftGraphOnlineMeetingProviderType**](anyOf&lt;microsoft.graph.onlineMeetingProviderType&gt;.md) |  | [optional] 
**OnlineMeetingUrl** | Pointer to **NullableString** |  | [optional] 
**Organizer** | Pointer to [**AnyOfmicrosoftGraphRecipient**](anyOf&lt;microsoft.graph.recipient&gt;.md) |  | [optional] 
**OriginalEndTimeZone** | Pointer to **NullableString** |  | [optional] 
**OriginalStart** | Pointer to **NullableTime** |  | [optional] 
**OriginalStartTimeZone** | Pointer to **NullableString** |  | [optional] 
**Recurrence** | Pointer to [**AnyOfmicrosoftGraphPatternedRecurrence**](anyOf&lt;microsoft.graph.patternedRecurrence&gt;.md) |  | [optional] 
**ReminderMinutesBeforeStart** | Pointer to **NullableInt32** |  | [optional] 
**ResponseRequested** | Pointer to **NullableBool** |  | [optional] 
**ResponseStatus** | Pointer to [**AnyOfmicrosoftGraphResponseStatus**](anyOf&lt;microsoft.graph.responseStatus&gt;.md) |  | [optional] 
**Sensitivity** | Pointer to [**AnyOfmicrosoftGraphSensitivity**](anyOf&lt;microsoft.graph.sensitivity&gt;.md) |  | [optional] 
**SeriesMasterId** | Pointer to **NullableString** |  | [optional] 
**ShowAs** | Pointer to [**AnyOfmicrosoftGraphFreeBusyStatus**](anyOf&lt;microsoft.graph.freeBusyStatus&gt;.md) |  | [optional] 
**Start** | Pointer to [**AnyOfmicrosoftGraphDateTimeTimeZone**](anyOf&lt;microsoft.graph.dateTimeTimeZone&gt;.md) |  | [optional] 
**Subject** | Pointer to **NullableString** |  | [optional] 
**TransactionId** | Pointer to **NullableString** |  | [optional] 
**Type** | Pointer to [**AnyOfmicrosoftGraphEventType**](anyOf&lt;microsoft.graph.eventType&gt;.md) |  | [optional] 
**WebLink** | Pointer to **NullableString** |  | [optional] 
**Attachments** | Pointer to [**[]MicrosoftGraphAttachment**](MicrosoftGraphAttachment.md) | The collection of FileAttachment, ItemAttachment, and referenceAttachment attachments for the event. Navigation property. Read-only. Nullable. | [optional] 
**Calendar** | Pointer to [**AnyOfmicrosoftGraphCalendar**](anyOf&lt;microsoft.graph.calendar&gt;.md) | The calendar that contains the event. Navigation property. Read-only. | [optional] 
**Extensions** | Pointer to [**[]MicrosoftGraphExtension**](MicrosoftGraphExtension.md) | The collection of open extensions defined for the event. Nullable. | [optional] 
**Instances** | Pointer to [**[]MicrosoftGraphEvent**](MicrosoftGraphEvent.md) | The occurrences of a recurring series, if the event is a series master. This property includes occurrences that are part of the recurrence pattern, and exceptions that have been modified, but does not include occurrences that have been cancelled from the series. Navigation property. Read-only. Nullable. | [optional] 
**MultiValueExtendedProperties** | Pointer to [**[]MicrosoftGraphMultiValueLegacyExtendedProperty**](MicrosoftGraphMultiValueLegacyExtendedProperty.md) | The collection of multi-value extended properties defined for the event. Read-only. Nullable. | [optional] 
**SingleValueExtendedProperties** | Pointer to [**[]MicrosoftGraphSingleValueLegacyExtendedProperty**](MicrosoftGraphSingleValueLegacyExtendedProperty.md) | The collection of single-value extended properties defined for the event. Read-only. Nullable. | [optional] 

## Methods

### NewMicrosoftGraphEvent

`func NewMicrosoftGraphEvent() *MicrosoftGraphEvent`

NewMicrosoftGraphEvent instantiates a new MicrosoftGraphEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphEventWithDefaults

`func NewMicrosoftGraphEventWithDefaults() *MicrosoftGraphEvent`

NewMicrosoftGraphEventWithDefaults instantiates a new MicrosoftGraphEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphEvent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphEvent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphEvent) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCategories

`func (o *MicrosoftGraphEvent) GetCategories() []*string`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *MicrosoftGraphEvent) GetCategoriesOk() (*[]*string, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *MicrosoftGraphEvent) SetCategories(v []*string)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *MicrosoftGraphEvent) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### GetChangeKey

`func (o *MicrosoftGraphEvent) GetChangeKey() string`

GetChangeKey returns the ChangeKey field if non-nil, zero value otherwise.

### GetChangeKeyOk

`func (o *MicrosoftGraphEvent) GetChangeKeyOk() (*string, bool)`

GetChangeKeyOk returns a tuple with the ChangeKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeKey

`func (o *MicrosoftGraphEvent) SetChangeKey(v string)`

SetChangeKey sets ChangeKey field to given value.

### HasChangeKey

`func (o *MicrosoftGraphEvent) HasChangeKey() bool`

HasChangeKey returns a boolean if a field has been set.

### SetChangeKeyNil

`func (o *MicrosoftGraphEvent) SetChangeKeyNil(b bool)`

 SetChangeKeyNil sets the value for ChangeKey to be an explicit nil

### UnsetChangeKey
`func (o *MicrosoftGraphEvent) UnsetChangeKey()`

UnsetChangeKey ensures that no value is present for ChangeKey, not even an explicit nil
### GetCreatedDateTime

`func (o *MicrosoftGraphEvent) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *MicrosoftGraphEvent) GetCreatedDateTimeOk() (*time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateTime

`func (o *MicrosoftGraphEvent) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime sets CreatedDateTime field to given value.

### HasCreatedDateTime

`func (o *MicrosoftGraphEvent) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTimeNil

`func (o *MicrosoftGraphEvent) SetCreatedDateTimeNil(b bool)`

 SetCreatedDateTimeNil sets the value for CreatedDateTime to be an explicit nil

### UnsetCreatedDateTime
`func (o *MicrosoftGraphEvent) UnsetCreatedDateTime()`

UnsetCreatedDateTime ensures that no value is present for CreatedDateTime, not even an explicit nil
### GetLastModifiedDateTime

`func (o *MicrosoftGraphEvent) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *MicrosoftGraphEvent) GetLastModifiedDateTimeOk() (*time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedDateTime

`func (o *MicrosoftGraphEvent) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime sets LastModifiedDateTime field to given value.

### HasLastModifiedDateTime

`func (o *MicrosoftGraphEvent) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### SetLastModifiedDateTimeNil

`func (o *MicrosoftGraphEvent) SetLastModifiedDateTimeNil(b bool)`

 SetLastModifiedDateTimeNil sets the value for LastModifiedDateTime to be an explicit nil

### UnsetLastModifiedDateTime
`func (o *MicrosoftGraphEvent) UnsetLastModifiedDateTime()`

UnsetLastModifiedDateTime ensures that no value is present for LastModifiedDateTime, not even an explicit nil
### GetAllowNewTimeProposals

`func (o *MicrosoftGraphEvent) GetAllowNewTimeProposals() bool`

GetAllowNewTimeProposals returns the AllowNewTimeProposals field if non-nil, zero value otherwise.

### GetAllowNewTimeProposalsOk

`func (o *MicrosoftGraphEvent) GetAllowNewTimeProposalsOk() (*bool, bool)`

GetAllowNewTimeProposalsOk returns a tuple with the AllowNewTimeProposals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowNewTimeProposals

`func (o *MicrosoftGraphEvent) SetAllowNewTimeProposals(v bool)`

SetAllowNewTimeProposals sets AllowNewTimeProposals field to given value.

### HasAllowNewTimeProposals

`func (o *MicrosoftGraphEvent) HasAllowNewTimeProposals() bool`

HasAllowNewTimeProposals returns a boolean if a field has been set.

### SetAllowNewTimeProposalsNil

`func (o *MicrosoftGraphEvent) SetAllowNewTimeProposalsNil(b bool)`

 SetAllowNewTimeProposalsNil sets the value for AllowNewTimeProposals to be an explicit nil

### UnsetAllowNewTimeProposals
`func (o *MicrosoftGraphEvent) UnsetAllowNewTimeProposals()`

UnsetAllowNewTimeProposals ensures that no value is present for AllowNewTimeProposals, not even an explicit nil
### GetAttendees

`func (o *MicrosoftGraphEvent) GetAttendees() []*AnyOfmicrosoftGraphAttendee`

GetAttendees returns the Attendees field if non-nil, zero value otherwise.

### GetAttendeesOk

`func (o *MicrosoftGraphEvent) GetAttendeesOk() (*[]*AnyOfmicrosoftGraphAttendee, bool)`

GetAttendeesOk returns a tuple with the Attendees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttendees

`func (o *MicrosoftGraphEvent) SetAttendees(v []*AnyOfmicrosoftGraphAttendee)`

SetAttendees sets Attendees field to given value.

### HasAttendees

`func (o *MicrosoftGraphEvent) HasAttendees() bool`

HasAttendees returns a boolean if a field has been set.

### GetBody

`func (o *MicrosoftGraphEvent) GetBody() AnyOfmicrosoftGraphItemBody`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *MicrosoftGraphEvent) GetBodyOk() (*AnyOfmicrosoftGraphItemBody, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *MicrosoftGraphEvent) SetBody(v AnyOfmicrosoftGraphItemBody)`

SetBody sets Body field to given value.

### HasBody

`func (o *MicrosoftGraphEvent) HasBody() bool`

HasBody returns a boolean if a field has been set.

### SetBodyNil

`func (o *MicrosoftGraphEvent) SetBodyNil(b bool)`

 SetBodyNil sets the value for Body to be an explicit nil

### UnsetBody
`func (o *MicrosoftGraphEvent) UnsetBody()`

UnsetBody ensures that no value is present for Body, not even an explicit nil
### GetBodyPreview

`func (o *MicrosoftGraphEvent) GetBodyPreview() string`

GetBodyPreview returns the BodyPreview field if non-nil, zero value otherwise.

### GetBodyPreviewOk

`func (o *MicrosoftGraphEvent) GetBodyPreviewOk() (*string, bool)`

GetBodyPreviewOk returns a tuple with the BodyPreview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBodyPreview

`func (o *MicrosoftGraphEvent) SetBodyPreview(v string)`

SetBodyPreview sets BodyPreview field to given value.

### HasBodyPreview

`func (o *MicrosoftGraphEvent) HasBodyPreview() bool`

HasBodyPreview returns a boolean if a field has been set.

### SetBodyPreviewNil

`func (o *MicrosoftGraphEvent) SetBodyPreviewNil(b bool)`

 SetBodyPreviewNil sets the value for BodyPreview to be an explicit nil

### UnsetBodyPreview
`func (o *MicrosoftGraphEvent) UnsetBodyPreview()`

UnsetBodyPreview ensures that no value is present for BodyPreview, not even an explicit nil
### GetEnd

`func (o *MicrosoftGraphEvent) GetEnd() AnyOfmicrosoftGraphDateTimeTimeZone`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *MicrosoftGraphEvent) GetEndOk() (*AnyOfmicrosoftGraphDateTimeTimeZone, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *MicrosoftGraphEvent) SetEnd(v AnyOfmicrosoftGraphDateTimeTimeZone)`

SetEnd sets End field to given value.

### HasEnd

`func (o *MicrosoftGraphEvent) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### SetEndNil

`func (o *MicrosoftGraphEvent) SetEndNil(b bool)`

 SetEndNil sets the value for End to be an explicit nil

### UnsetEnd
`func (o *MicrosoftGraphEvent) UnsetEnd()`

UnsetEnd ensures that no value is present for End, not even an explicit nil
### GetHasAttachments

`func (o *MicrosoftGraphEvent) GetHasAttachments() bool`

GetHasAttachments returns the HasAttachments field if non-nil, zero value otherwise.

### GetHasAttachmentsOk

`func (o *MicrosoftGraphEvent) GetHasAttachmentsOk() (*bool, bool)`

GetHasAttachmentsOk returns a tuple with the HasAttachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAttachments

`func (o *MicrosoftGraphEvent) SetHasAttachments(v bool)`

SetHasAttachments sets HasAttachments field to given value.

### HasHasAttachments

`func (o *MicrosoftGraphEvent) HasHasAttachments() bool`

HasHasAttachments returns a boolean if a field has been set.

### SetHasAttachmentsNil

`func (o *MicrosoftGraphEvent) SetHasAttachmentsNil(b bool)`

 SetHasAttachmentsNil sets the value for HasAttachments to be an explicit nil

### UnsetHasAttachments
`func (o *MicrosoftGraphEvent) UnsetHasAttachments()`

UnsetHasAttachments ensures that no value is present for HasAttachments, not even an explicit nil
### GetHideAttendees

`func (o *MicrosoftGraphEvent) GetHideAttendees() bool`

GetHideAttendees returns the HideAttendees field if non-nil, zero value otherwise.

### GetHideAttendeesOk

`func (o *MicrosoftGraphEvent) GetHideAttendeesOk() (*bool, bool)`

GetHideAttendeesOk returns a tuple with the HideAttendees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideAttendees

`func (o *MicrosoftGraphEvent) SetHideAttendees(v bool)`

SetHideAttendees sets HideAttendees field to given value.

### HasHideAttendees

`func (o *MicrosoftGraphEvent) HasHideAttendees() bool`

HasHideAttendees returns a boolean if a field has been set.

### SetHideAttendeesNil

`func (o *MicrosoftGraphEvent) SetHideAttendeesNil(b bool)`

 SetHideAttendeesNil sets the value for HideAttendees to be an explicit nil

### UnsetHideAttendees
`func (o *MicrosoftGraphEvent) UnsetHideAttendees()`

UnsetHideAttendees ensures that no value is present for HideAttendees, not even an explicit nil
### GetICalUId

`func (o *MicrosoftGraphEvent) GetICalUId() string`

GetICalUId returns the ICalUId field if non-nil, zero value otherwise.

### GetICalUIdOk

`func (o *MicrosoftGraphEvent) GetICalUIdOk() (*string, bool)`

GetICalUIdOk returns a tuple with the ICalUId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetICalUId

`func (o *MicrosoftGraphEvent) SetICalUId(v string)`

SetICalUId sets ICalUId field to given value.

### HasICalUId

`func (o *MicrosoftGraphEvent) HasICalUId() bool`

HasICalUId returns a boolean if a field has been set.

### SetICalUIdNil

`func (o *MicrosoftGraphEvent) SetICalUIdNil(b bool)`

 SetICalUIdNil sets the value for ICalUId to be an explicit nil

### UnsetICalUId
`func (o *MicrosoftGraphEvent) UnsetICalUId()`

UnsetICalUId ensures that no value is present for ICalUId, not even an explicit nil
### GetImportance

`func (o *MicrosoftGraphEvent) GetImportance() AnyOfmicrosoftGraphImportance`

GetImportance returns the Importance field if non-nil, zero value otherwise.

### GetImportanceOk

`func (o *MicrosoftGraphEvent) GetImportanceOk() (*AnyOfmicrosoftGraphImportance, bool)`

GetImportanceOk returns a tuple with the Importance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportance

`func (o *MicrosoftGraphEvent) SetImportance(v AnyOfmicrosoftGraphImportance)`

SetImportance sets Importance field to given value.

### HasImportance

`func (o *MicrosoftGraphEvent) HasImportance() bool`

HasImportance returns a boolean if a field has been set.

### SetImportanceNil

`func (o *MicrosoftGraphEvent) SetImportanceNil(b bool)`

 SetImportanceNil sets the value for Importance to be an explicit nil

### UnsetImportance
`func (o *MicrosoftGraphEvent) UnsetImportance()`

UnsetImportance ensures that no value is present for Importance, not even an explicit nil
### GetIsAllDay

`func (o *MicrosoftGraphEvent) GetIsAllDay() bool`

GetIsAllDay returns the IsAllDay field if non-nil, zero value otherwise.

### GetIsAllDayOk

`func (o *MicrosoftGraphEvent) GetIsAllDayOk() (*bool, bool)`

GetIsAllDayOk returns a tuple with the IsAllDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAllDay

`func (o *MicrosoftGraphEvent) SetIsAllDay(v bool)`

SetIsAllDay sets IsAllDay field to given value.

### HasIsAllDay

`func (o *MicrosoftGraphEvent) HasIsAllDay() bool`

HasIsAllDay returns a boolean if a field has been set.

### SetIsAllDayNil

`func (o *MicrosoftGraphEvent) SetIsAllDayNil(b bool)`

 SetIsAllDayNil sets the value for IsAllDay to be an explicit nil

### UnsetIsAllDay
`func (o *MicrosoftGraphEvent) UnsetIsAllDay()`

UnsetIsAllDay ensures that no value is present for IsAllDay, not even an explicit nil
### GetIsCancelled

`func (o *MicrosoftGraphEvent) GetIsCancelled() bool`

GetIsCancelled returns the IsCancelled field if non-nil, zero value otherwise.

### GetIsCancelledOk

`func (o *MicrosoftGraphEvent) GetIsCancelledOk() (*bool, bool)`

GetIsCancelledOk returns a tuple with the IsCancelled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCancelled

`func (o *MicrosoftGraphEvent) SetIsCancelled(v bool)`

SetIsCancelled sets IsCancelled field to given value.

### HasIsCancelled

`func (o *MicrosoftGraphEvent) HasIsCancelled() bool`

HasIsCancelled returns a boolean if a field has been set.

### SetIsCancelledNil

`func (o *MicrosoftGraphEvent) SetIsCancelledNil(b bool)`

 SetIsCancelledNil sets the value for IsCancelled to be an explicit nil

### UnsetIsCancelled
`func (o *MicrosoftGraphEvent) UnsetIsCancelled()`

UnsetIsCancelled ensures that no value is present for IsCancelled, not even an explicit nil
### GetIsDraft

`func (o *MicrosoftGraphEvent) GetIsDraft() bool`

GetIsDraft returns the IsDraft field if non-nil, zero value otherwise.

### GetIsDraftOk

`func (o *MicrosoftGraphEvent) GetIsDraftOk() (*bool, bool)`

GetIsDraftOk returns a tuple with the IsDraft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDraft

`func (o *MicrosoftGraphEvent) SetIsDraft(v bool)`

SetIsDraft sets IsDraft field to given value.

### HasIsDraft

`func (o *MicrosoftGraphEvent) HasIsDraft() bool`

HasIsDraft returns a boolean if a field has been set.

### SetIsDraftNil

`func (o *MicrosoftGraphEvent) SetIsDraftNil(b bool)`

 SetIsDraftNil sets the value for IsDraft to be an explicit nil

### UnsetIsDraft
`func (o *MicrosoftGraphEvent) UnsetIsDraft()`

UnsetIsDraft ensures that no value is present for IsDraft, not even an explicit nil
### GetIsOnlineMeeting

`func (o *MicrosoftGraphEvent) GetIsOnlineMeeting() bool`

GetIsOnlineMeeting returns the IsOnlineMeeting field if non-nil, zero value otherwise.

### GetIsOnlineMeetingOk

`func (o *MicrosoftGraphEvent) GetIsOnlineMeetingOk() (*bool, bool)`

GetIsOnlineMeetingOk returns a tuple with the IsOnlineMeeting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOnlineMeeting

`func (o *MicrosoftGraphEvent) SetIsOnlineMeeting(v bool)`

SetIsOnlineMeeting sets IsOnlineMeeting field to given value.

### HasIsOnlineMeeting

`func (o *MicrosoftGraphEvent) HasIsOnlineMeeting() bool`

HasIsOnlineMeeting returns a boolean if a field has been set.

### SetIsOnlineMeetingNil

`func (o *MicrosoftGraphEvent) SetIsOnlineMeetingNil(b bool)`

 SetIsOnlineMeetingNil sets the value for IsOnlineMeeting to be an explicit nil

### UnsetIsOnlineMeeting
`func (o *MicrosoftGraphEvent) UnsetIsOnlineMeeting()`

UnsetIsOnlineMeeting ensures that no value is present for IsOnlineMeeting, not even an explicit nil
### GetIsOrganizer

`func (o *MicrosoftGraphEvent) GetIsOrganizer() bool`

GetIsOrganizer returns the IsOrganizer field if non-nil, zero value otherwise.

### GetIsOrganizerOk

`func (o *MicrosoftGraphEvent) GetIsOrganizerOk() (*bool, bool)`

GetIsOrganizerOk returns a tuple with the IsOrganizer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOrganizer

`func (o *MicrosoftGraphEvent) SetIsOrganizer(v bool)`

SetIsOrganizer sets IsOrganizer field to given value.

### HasIsOrganizer

`func (o *MicrosoftGraphEvent) HasIsOrganizer() bool`

HasIsOrganizer returns a boolean if a field has been set.

### SetIsOrganizerNil

`func (o *MicrosoftGraphEvent) SetIsOrganizerNil(b bool)`

 SetIsOrganizerNil sets the value for IsOrganizer to be an explicit nil

### UnsetIsOrganizer
`func (o *MicrosoftGraphEvent) UnsetIsOrganizer()`

UnsetIsOrganizer ensures that no value is present for IsOrganizer, not even an explicit nil
### GetIsReminderOn

`func (o *MicrosoftGraphEvent) GetIsReminderOn() bool`

GetIsReminderOn returns the IsReminderOn field if non-nil, zero value otherwise.

### GetIsReminderOnOk

`func (o *MicrosoftGraphEvent) GetIsReminderOnOk() (*bool, bool)`

GetIsReminderOnOk returns a tuple with the IsReminderOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReminderOn

`func (o *MicrosoftGraphEvent) SetIsReminderOn(v bool)`

SetIsReminderOn sets IsReminderOn field to given value.

### HasIsReminderOn

`func (o *MicrosoftGraphEvent) HasIsReminderOn() bool`

HasIsReminderOn returns a boolean if a field has been set.

### SetIsReminderOnNil

`func (o *MicrosoftGraphEvent) SetIsReminderOnNil(b bool)`

 SetIsReminderOnNil sets the value for IsReminderOn to be an explicit nil

### UnsetIsReminderOn
`func (o *MicrosoftGraphEvent) UnsetIsReminderOn()`

UnsetIsReminderOn ensures that no value is present for IsReminderOn, not even an explicit nil
### GetLocation

`func (o *MicrosoftGraphEvent) GetLocation() AnyOfmicrosoftGraphLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *MicrosoftGraphEvent) GetLocationOk() (*AnyOfmicrosoftGraphLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *MicrosoftGraphEvent) SetLocation(v AnyOfmicrosoftGraphLocation)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *MicrosoftGraphEvent) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocationNil

`func (o *MicrosoftGraphEvent) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *MicrosoftGraphEvent) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil
### GetLocations

`func (o *MicrosoftGraphEvent) GetLocations() []*AnyOfmicrosoftGraphLocation`

GetLocations returns the Locations field if non-nil, zero value otherwise.

### GetLocationsOk

`func (o *MicrosoftGraphEvent) GetLocationsOk() (*[]*AnyOfmicrosoftGraphLocation, bool)`

GetLocationsOk returns a tuple with the Locations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocations

`func (o *MicrosoftGraphEvent) SetLocations(v []*AnyOfmicrosoftGraphLocation)`

SetLocations sets Locations field to given value.

### HasLocations

`func (o *MicrosoftGraphEvent) HasLocations() bool`

HasLocations returns a boolean if a field has been set.

### GetOnlineMeeting

`func (o *MicrosoftGraphEvent) GetOnlineMeeting() AnyOfmicrosoftGraphOnlineMeetingInfo`

GetOnlineMeeting returns the OnlineMeeting field if non-nil, zero value otherwise.

### GetOnlineMeetingOk

`func (o *MicrosoftGraphEvent) GetOnlineMeetingOk() (*AnyOfmicrosoftGraphOnlineMeetingInfo, bool)`

GetOnlineMeetingOk returns a tuple with the OnlineMeeting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnlineMeeting

`func (o *MicrosoftGraphEvent) SetOnlineMeeting(v AnyOfmicrosoftGraphOnlineMeetingInfo)`

SetOnlineMeeting sets OnlineMeeting field to given value.

### HasOnlineMeeting

`func (o *MicrosoftGraphEvent) HasOnlineMeeting() bool`

HasOnlineMeeting returns a boolean if a field has been set.

### SetOnlineMeetingNil

`func (o *MicrosoftGraphEvent) SetOnlineMeetingNil(b bool)`

 SetOnlineMeetingNil sets the value for OnlineMeeting to be an explicit nil

### UnsetOnlineMeeting
`func (o *MicrosoftGraphEvent) UnsetOnlineMeeting()`

UnsetOnlineMeeting ensures that no value is present for OnlineMeeting, not even an explicit nil
### GetOnlineMeetingProvider

`func (o *MicrosoftGraphEvent) GetOnlineMeetingProvider() AnyOfmicrosoftGraphOnlineMeetingProviderType`

GetOnlineMeetingProvider returns the OnlineMeetingProvider field if non-nil, zero value otherwise.

### GetOnlineMeetingProviderOk

`func (o *MicrosoftGraphEvent) GetOnlineMeetingProviderOk() (*AnyOfmicrosoftGraphOnlineMeetingProviderType, bool)`

GetOnlineMeetingProviderOk returns a tuple with the OnlineMeetingProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnlineMeetingProvider

`func (o *MicrosoftGraphEvent) SetOnlineMeetingProvider(v AnyOfmicrosoftGraphOnlineMeetingProviderType)`

SetOnlineMeetingProvider sets OnlineMeetingProvider field to given value.

### HasOnlineMeetingProvider

`func (o *MicrosoftGraphEvent) HasOnlineMeetingProvider() bool`

HasOnlineMeetingProvider returns a boolean if a field has been set.

### SetOnlineMeetingProviderNil

`func (o *MicrosoftGraphEvent) SetOnlineMeetingProviderNil(b bool)`

 SetOnlineMeetingProviderNil sets the value for OnlineMeetingProvider to be an explicit nil

### UnsetOnlineMeetingProvider
`func (o *MicrosoftGraphEvent) UnsetOnlineMeetingProvider()`

UnsetOnlineMeetingProvider ensures that no value is present for OnlineMeetingProvider, not even an explicit nil
### GetOnlineMeetingUrl

`func (o *MicrosoftGraphEvent) GetOnlineMeetingUrl() string`

GetOnlineMeetingUrl returns the OnlineMeetingUrl field if non-nil, zero value otherwise.

### GetOnlineMeetingUrlOk

`func (o *MicrosoftGraphEvent) GetOnlineMeetingUrlOk() (*string, bool)`

GetOnlineMeetingUrlOk returns a tuple with the OnlineMeetingUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnlineMeetingUrl

`func (o *MicrosoftGraphEvent) SetOnlineMeetingUrl(v string)`

SetOnlineMeetingUrl sets OnlineMeetingUrl field to given value.

### HasOnlineMeetingUrl

`func (o *MicrosoftGraphEvent) HasOnlineMeetingUrl() bool`

HasOnlineMeetingUrl returns a boolean if a field has been set.

### SetOnlineMeetingUrlNil

`func (o *MicrosoftGraphEvent) SetOnlineMeetingUrlNil(b bool)`

 SetOnlineMeetingUrlNil sets the value for OnlineMeetingUrl to be an explicit nil

### UnsetOnlineMeetingUrl
`func (o *MicrosoftGraphEvent) UnsetOnlineMeetingUrl()`

UnsetOnlineMeetingUrl ensures that no value is present for OnlineMeetingUrl, not even an explicit nil
### GetOrganizer

`func (o *MicrosoftGraphEvent) GetOrganizer() AnyOfmicrosoftGraphRecipient`

GetOrganizer returns the Organizer field if non-nil, zero value otherwise.

### GetOrganizerOk

`func (o *MicrosoftGraphEvent) GetOrganizerOk() (*AnyOfmicrosoftGraphRecipient, bool)`

GetOrganizerOk returns a tuple with the Organizer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizer

`func (o *MicrosoftGraphEvent) SetOrganizer(v AnyOfmicrosoftGraphRecipient)`

SetOrganizer sets Organizer field to given value.

### HasOrganizer

`func (o *MicrosoftGraphEvent) HasOrganizer() bool`

HasOrganizer returns a boolean if a field has been set.

### SetOrganizerNil

`func (o *MicrosoftGraphEvent) SetOrganizerNil(b bool)`

 SetOrganizerNil sets the value for Organizer to be an explicit nil

### UnsetOrganizer
`func (o *MicrosoftGraphEvent) UnsetOrganizer()`

UnsetOrganizer ensures that no value is present for Organizer, not even an explicit nil
### GetOriginalEndTimeZone

`func (o *MicrosoftGraphEvent) GetOriginalEndTimeZone() string`

GetOriginalEndTimeZone returns the OriginalEndTimeZone field if non-nil, zero value otherwise.

### GetOriginalEndTimeZoneOk

`func (o *MicrosoftGraphEvent) GetOriginalEndTimeZoneOk() (*string, bool)`

GetOriginalEndTimeZoneOk returns a tuple with the OriginalEndTimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalEndTimeZone

`func (o *MicrosoftGraphEvent) SetOriginalEndTimeZone(v string)`

SetOriginalEndTimeZone sets OriginalEndTimeZone field to given value.

### HasOriginalEndTimeZone

`func (o *MicrosoftGraphEvent) HasOriginalEndTimeZone() bool`

HasOriginalEndTimeZone returns a boolean if a field has been set.

### SetOriginalEndTimeZoneNil

`func (o *MicrosoftGraphEvent) SetOriginalEndTimeZoneNil(b bool)`

 SetOriginalEndTimeZoneNil sets the value for OriginalEndTimeZone to be an explicit nil

### UnsetOriginalEndTimeZone
`func (o *MicrosoftGraphEvent) UnsetOriginalEndTimeZone()`

UnsetOriginalEndTimeZone ensures that no value is present for OriginalEndTimeZone, not even an explicit nil
### GetOriginalStart

`func (o *MicrosoftGraphEvent) GetOriginalStart() time.Time`

GetOriginalStart returns the OriginalStart field if non-nil, zero value otherwise.

### GetOriginalStartOk

`func (o *MicrosoftGraphEvent) GetOriginalStartOk() (*time.Time, bool)`

GetOriginalStartOk returns a tuple with the OriginalStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalStart

`func (o *MicrosoftGraphEvent) SetOriginalStart(v time.Time)`

SetOriginalStart sets OriginalStart field to given value.

### HasOriginalStart

`func (o *MicrosoftGraphEvent) HasOriginalStart() bool`

HasOriginalStart returns a boolean if a field has been set.

### SetOriginalStartNil

`func (o *MicrosoftGraphEvent) SetOriginalStartNil(b bool)`

 SetOriginalStartNil sets the value for OriginalStart to be an explicit nil

### UnsetOriginalStart
`func (o *MicrosoftGraphEvent) UnsetOriginalStart()`

UnsetOriginalStart ensures that no value is present for OriginalStart, not even an explicit nil
### GetOriginalStartTimeZone

`func (o *MicrosoftGraphEvent) GetOriginalStartTimeZone() string`

GetOriginalStartTimeZone returns the OriginalStartTimeZone field if non-nil, zero value otherwise.

### GetOriginalStartTimeZoneOk

`func (o *MicrosoftGraphEvent) GetOriginalStartTimeZoneOk() (*string, bool)`

GetOriginalStartTimeZoneOk returns a tuple with the OriginalStartTimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalStartTimeZone

`func (o *MicrosoftGraphEvent) SetOriginalStartTimeZone(v string)`

SetOriginalStartTimeZone sets OriginalStartTimeZone field to given value.

### HasOriginalStartTimeZone

`func (o *MicrosoftGraphEvent) HasOriginalStartTimeZone() bool`

HasOriginalStartTimeZone returns a boolean if a field has been set.

### SetOriginalStartTimeZoneNil

`func (o *MicrosoftGraphEvent) SetOriginalStartTimeZoneNil(b bool)`

 SetOriginalStartTimeZoneNil sets the value for OriginalStartTimeZone to be an explicit nil

### UnsetOriginalStartTimeZone
`func (o *MicrosoftGraphEvent) UnsetOriginalStartTimeZone()`

UnsetOriginalStartTimeZone ensures that no value is present for OriginalStartTimeZone, not even an explicit nil
### GetRecurrence

`func (o *MicrosoftGraphEvent) GetRecurrence() AnyOfmicrosoftGraphPatternedRecurrence`

GetRecurrence returns the Recurrence field if non-nil, zero value otherwise.

### GetRecurrenceOk

`func (o *MicrosoftGraphEvent) GetRecurrenceOk() (*AnyOfmicrosoftGraphPatternedRecurrence, bool)`

GetRecurrenceOk returns a tuple with the Recurrence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurrence

`func (o *MicrosoftGraphEvent) SetRecurrence(v AnyOfmicrosoftGraphPatternedRecurrence)`

SetRecurrence sets Recurrence field to given value.

### HasRecurrence

`func (o *MicrosoftGraphEvent) HasRecurrence() bool`

HasRecurrence returns a boolean if a field has been set.

### SetRecurrenceNil

`func (o *MicrosoftGraphEvent) SetRecurrenceNil(b bool)`

 SetRecurrenceNil sets the value for Recurrence to be an explicit nil

### UnsetRecurrence
`func (o *MicrosoftGraphEvent) UnsetRecurrence()`

UnsetRecurrence ensures that no value is present for Recurrence, not even an explicit nil
### GetReminderMinutesBeforeStart

`func (o *MicrosoftGraphEvent) GetReminderMinutesBeforeStart() int32`

GetReminderMinutesBeforeStart returns the ReminderMinutesBeforeStart field if non-nil, zero value otherwise.

### GetReminderMinutesBeforeStartOk

`func (o *MicrosoftGraphEvent) GetReminderMinutesBeforeStartOk() (*int32, bool)`

GetReminderMinutesBeforeStartOk returns a tuple with the ReminderMinutesBeforeStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReminderMinutesBeforeStart

`func (o *MicrosoftGraphEvent) SetReminderMinutesBeforeStart(v int32)`

SetReminderMinutesBeforeStart sets ReminderMinutesBeforeStart field to given value.

### HasReminderMinutesBeforeStart

`func (o *MicrosoftGraphEvent) HasReminderMinutesBeforeStart() bool`

HasReminderMinutesBeforeStart returns a boolean if a field has been set.

### SetReminderMinutesBeforeStartNil

`func (o *MicrosoftGraphEvent) SetReminderMinutesBeforeStartNil(b bool)`

 SetReminderMinutesBeforeStartNil sets the value for ReminderMinutesBeforeStart to be an explicit nil

### UnsetReminderMinutesBeforeStart
`func (o *MicrosoftGraphEvent) UnsetReminderMinutesBeforeStart()`

UnsetReminderMinutesBeforeStart ensures that no value is present for ReminderMinutesBeforeStart, not even an explicit nil
### GetResponseRequested

`func (o *MicrosoftGraphEvent) GetResponseRequested() bool`

GetResponseRequested returns the ResponseRequested field if non-nil, zero value otherwise.

### GetResponseRequestedOk

`func (o *MicrosoftGraphEvent) GetResponseRequestedOk() (*bool, bool)`

GetResponseRequestedOk returns a tuple with the ResponseRequested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseRequested

`func (o *MicrosoftGraphEvent) SetResponseRequested(v bool)`

SetResponseRequested sets ResponseRequested field to given value.

### HasResponseRequested

`func (o *MicrosoftGraphEvent) HasResponseRequested() bool`

HasResponseRequested returns a boolean if a field has been set.

### SetResponseRequestedNil

`func (o *MicrosoftGraphEvent) SetResponseRequestedNil(b bool)`

 SetResponseRequestedNil sets the value for ResponseRequested to be an explicit nil

### UnsetResponseRequested
`func (o *MicrosoftGraphEvent) UnsetResponseRequested()`

UnsetResponseRequested ensures that no value is present for ResponseRequested, not even an explicit nil
### GetResponseStatus

`func (o *MicrosoftGraphEvent) GetResponseStatus() AnyOfmicrosoftGraphResponseStatus`

GetResponseStatus returns the ResponseStatus field if non-nil, zero value otherwise.

### GetResponseStatusOk

`func (o *MicrosoftGraphEvent) GetResponseStatusOk() (*AnyOfmicrosoftGraphResponseStatus, bool)`

GetResponseStatusOk returns a tuple with the ResponseStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseStatus

`func (o *MicrosoftGraphEvent) SetResponseStatus(v AnyOfmicrosoftGraphResponseStatus)`

SetResponseStatus sets ResponseStatus field to given value.

### HasResponseStatus

`func (o *MicrosoftGraphEvent) HasResponseStatus() bool`

HasResponseStatus returns a boolean if a field has been set.

### SetResponseStatusNil

`func (o *MicrosoftGraphEvent) SetResponseStatusNil(b bool)`

 SetResponseStatusNil sets the value for ResponseStatus to be an explicit nil

### UnsetResponseStatus
`func (o *MicrosoftGraphEvent) UnsetResponseStatus()`

UnsetResponseStatus ensures that no value is present for ResponseStatus, not even an explicit nil
### GetSensitivity

`func (o *MicrosoftGraphEvent) GetSensitivity() AnyOfmicrosoftGraphSensitivity`

GetSensitivity returns the Sensitivity field if non-nil, zero value otherwise.

### GetSensitivityOk

`func (o *MicrosoftGraphEvent) GetSensitivityOk() (*AnyOfmicrosoftGraphSensitivity, bool)`

GetSensitivityOk returns a tuple with the Sensitivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSensitivity

`func (o *MicrosoftGraphEvent) SetSensitivity(v AnyOfmicrosoftGraphSensitivity)`

SetSensitivity sets Sensitivity field to given value.

### HasSensitivity

`func (o *MicrosoftGraphEvent) HasSensitivity() bool`

HasSensitivity returns a boolean if a field has been set.

### SetSensitivityNil

`func (o *MicrosoftGraphEvent) SetSensitivityNil(b bool)`

 SetSensitivityNil sets the value for Sensitivity to be an explicit nil

### UnsetSensitivity
`func (o *MicrosoftGraphEvent) UnsetSensitivity()`

UnsetSensitivity ensures that no value is present for Sensitivity, not even an explicit nil
### GetSeriesMasterId

`func (o *MicrosoftGraphEvent) GetSeriesMasterId() string`

GetSeriesMasterId returns the SeriesMasterId field if non-nil, zero value otherwise.

### GetSeriesMasterIdOk

`func (o *MicrosoftGraphEvent) GetSeriesMasterIdOk() (*string, bool)`

GetSeriesMasterIdOk returns a tuple with the SeriesMasterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeriesMasterId

`func (o *MicrosoftGraphEvent) SetSeriesMasterId(v string)`

SetSeriesMasterId sets SeriesMasterId field to given value.

### HasSeriesMasterId

`func (o *MicrosoftGraphEvent) HasSeriesMasterId() bool`

HasSeriesMasterId returns a boolean if a field has been set.

### SetSeriesMasterIdNil

`func (o *MicrosoftGraphEvent) SetSeriesMasterIdNil(b bool)`

 SetSeriesMasterIdNil sets the value for SeriesMasterId to be an explicit nil

### UnsetSeriesMasterId
`func (o *MicrosoftGraphEvent) UnsetSeriesMasterId()`

UnsetSeriesMasterId ensures that no value is present for SeriesMasterId, not even an explicit nil
### GetShowAs

`func (o *MicrosoftGraphEvent) GetShowAs() AnyOfmicrosoftGraphFreeBusyStatus`

GetShowAs returns the ShowAs field if non-nil, zero value otherwise.

### GetShowAsOk

`func (o *MicrosoftGraphEvent) GetShowAsOk() (*AnyOfmicrosoftGraphFreeBusyStatus, bool)`

GetShowAsOk returns a tuple with the ShowAs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowAs

`func (o *MicrosoftGraphEvent) SetShowAs(v AnyOfmicrosoftGraphFreeBusyStatus)`

SetShowAs sets ShowAs field to given value.

### HasShowAs

`func (o *MicrosoftGraphEvent) HasShowAs() bool`

HasShowAs returns a boolean if a field has been set.

### SetShowAsNil

`func (o *MicrosoftGraphEvent) SetShowAsNil(b bool)`

 SetShowAsNil sets the value for ShowAs to be an explicit nil

### UnsetShowAs
`func (o *MicrosoftGraphEvent) UnsetShowAs()`

UnsetShowAs ensures that no value is present for ShowAs, not even an explicit nil
### GetStart

`func (o *MicrosoftGraphEvent) GetStart() AnyOfmicrosoftGraphDateTimeTimeZone`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *MicrosoftGraphEvent) GetStartOk() (*AnyOfmicrosoftGraphDateTimeTimeZone, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *MicrosoftGraphEvent) SetStart(v AnyOfmicrosoftGraphDateTimeTimeZone)`

SetStart sets Start field to given value.

### HasStart

`func (o *MicrosoftGraphEvent) HasStart() bool`

HasStart returns a boolean if a field has been set.

### SetStartNil

`func (o *MicrosoftGraphEvent) SetStartNil(b bool)`

 SetStartNil sets the value for Start to be an explicit nil

### UnsetStart
`func (o *MicrosoftGraphEvent) UnsetStart()`

UnsetStart ensures that no value is present for Start, not even an explicit nil
### GetSubject

`func (o *MicrosoftGraphEvent) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *MicrosoftGraphEvent) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *MicrosoftGraphEvent) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *MicrosoftGraphEvent) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### SetSubjectNil

`func (o *MicrosoftGraphEvent) SetSubjectNil(b bool)`

 SetSubjectNil sets the value for Subject to be an explicit nil

### UnsetSubject
`func (o *MicrosoftGraphEvent) UnsetSubject()`

UnsetSubject ensures that no value is present for Subject, not even an explicit nil
### GetTransactionId

`func (o *MicrosoftGraphEvent) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *MicrosoftGraphEvent) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *MicrosoftGraphEvent) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.

### HasTransactionId

`func (o *MicrosoftGraphEvent) HasTransactionId() bool`

HasTransactionId returns a boolean if a field has been set.

### SetTransactionIdNil

`func (o *MicrosoftGraphEvent) SetTransactionIdNil(b bool)`

 SetTransactionIdNil sets the value for TransactionId to be an explicit nil

### UnsetTransactionId
`func (o *MicrosoftGraphEvent) UnsetTransactionId()`

UnsetTransactionId ensures that no value is present for TransactionId, not even an explicit nil
### GetType

`func (o *MicrosoftGraphEvent) GetType() AnyOfmicrosoftGraphEventType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MicrosoftGraphEvent) GetTypeOk() (*AnyOfmicrosoftGraphEventType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MicrosoftGraphEvent) SetType(v AnyOfmicrosoftGraphEventType)`

SetType sets Type field to given value.

### HasType

`func (o *MicrosoftGraphEvent) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *MicrosoftGraphEvent) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *MicrosoftGraphEvent) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetWebLink

`func (o *MicrosoftGraphEvent) GetWebLink() string`

GetWebLink returns the WebLink field if non-nil, zero value otherwise.

### GetWebLinkOk

`func (o *MicrosoftGraphEvent) GetWebLinkOk() (*string, bool)`

GetWebLinkOk returns a tuple with the WebLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebLink

`func (o *MicrosoftGraphEvent) SetWebLink(v string)`

SetWebLink sets WebLink field to given value.

### HasWebLink

`func (o *MicrosoftGraphEvent) HasWebLink() bool`

HasWebLink returns a boolean if a field has been set.

### SetWebLinkNil

`func (o *MicrosoftGraphEvent) SetWebLinkNil(b bool)`

 SetWebLinkNil sets the value for WebLink to be an explicit nil

### UnsetWebLink
`func (o *MicrosoftGraphEvent) UnsetWebLink()`

UnsetWebLink ensures that no value is present for WebLink, not even an explicit nil
### GetAttachments

`func (o *MicrosoftGraphEvent) GetAttachments() []MicrosoftGraphAttachment`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *MicrosoftGraphEvent) GetAttachmentsOk() (*[]MicrosoftGraphAttachment, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *MicrosoftGraphEvent) SetAttachments(v []MicrosoftGraphAttachment)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *MicrosoftGraphEvent) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetCalendar

`func (o *MicrosoftGraphEvent) GetCalendar() AnyOfmicrosoftGraphCalendar`

GetCalendar returns the Calendar field if non-nil, zero value otherwise.

### GetCalendarOk

`func (o *MicrosoftGraphEvent) GetCalendarOk() (*AnyOfmicrosoftGraphCalendar, bool)`

GetCalendarOk returns a tuple with the Calendar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalendar

`func (o *MicrosoftGraphEvent) SetCalendar(v AnyOfmicrosoftGraphCalendar)`

SetCalendar sets Calendar field to given value.

### HasCalendar

`func (o *MicrosoftGraphEvent) HasCalendar() bool`

HasCalendar returns a boolean if a field has been set.

### SetCalendarNil

`func (o *MicrosoftGraphEvent) SetCalendarNil(b bool)`

 SetCalendarNil sets the value for Calendar to be an explicit nil

### UnsetCalendar
`func (o *MicrosoftGraphEvent) UnsetCalendar()`

UnsetCalendar ensures that no value is present for Calendar, not even an explicit nil
### GetExtensions

`func (o *MicrosoftGraphEvent) GetExtensions() []MicrosoftGraphExtension`

GetExtensions returns the Extensions field if non-nil, zero value otherwise.

### GetExtensionsOk

`func (o *MicrosoftGraphEvent) GetExtensionsOk() (*[]MicrosoftGraphExtension, bool)`

GetExtensionsOk returns a tuple with the Extensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensions

`func (o *MicrosoftGraphEvent) SetExtensions(v []MicrosoftGraphExtension)`

SetExtensions sets Extensions field to given value.

### HasExtensions

`func (o *MicrosoftGraphEvent) HasExtensions() bool`

HasExtensions returns a boolean if a field has been set.

### GetInstances

`func (o *MicrosoftGraphEvent) GetInstances() []MicrosoftGraphEvent`

GetInstances returns the Instances field if non-nil, zero value otherwise.

### GetInstancesOk

`func (o *MicrosoftGraphEvent) GetInstancesOk() (*[]MicrosoftGraphEvent, bool)`

GetInstancesOk returns a tuple with the Instances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstances

`func (o *MicrosoftGraphEvent) SetInstances(v []MicrosoftGraphEvent)`

SetInstances sets Instances field to given value.

### HasInstances

`func (o *MicrosoftGraphEvent) HasInstances() bool`

HasInstances returns a boolean if a field has been set.

### GetMultiValueExtendedProperties

`func (o *MicrosoftGraphEvent) GetMultiValueExtendedProperties() []MicrosoftGraphMultiValueLegacyExtendedProperty`

GetMultiValueExtendedProperties returns the MultiValueExtendedProperties field if non-nil, zero value otherwise.

### GetMultiValueExtendedPropertiesOk

`func (o *MicrosoftGraphEvent) GetMultiValueExtendedPropertiesOk() (*[]MicrosoftGraphMultiValueLegacyExtendedProperty, bool)`

GetMultiValueExtendedPropertiesOk returns a tuple with the MultiValueExtendedProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiValueExtendedProperties

`func (o *MicrosoftGraphEvent) SetMultiValueExtendedProperties(v []MicrosoftGraphMultiValueLegacyExtendedProperty)`

SetMultiValueExtendedProperties sets MultiValueExtendedProperties field to given value.

### HasMultiValueExtendedProperties

`func (o *MicrosoftGraphEvent) HasMultiValueExtendedProperties() bool`

HasMultiValueExtendedProperties returns a boolean if a field has been set.

### GetSingleValueExtendedProperties

`func (o *MicrosoftGraphEvent) GetSingleValueExtendedProperties() []MicrosoftGraphSingleValueLegacyExtendedProperty`

GetSingleValueExtendedProperties returns the SingleValueExtendedProperties field if non-nil, zero value otherwise.

### GetSingleValueExtendedPropertiesOk

`func (o *MicrosoftGraphEvent) GetSingleValueExtendedPropertiesOk() (*[]MicrosoftGraphSingleValueLegacyExtendedProperty, bool)`

GetSingleValueExtendedPropertiesOk returns a tuple with the SingleValueExtendedProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSingleValueExtendedProperties

`func (o *MicrosoftGraphEvent) SetSingleValueExtendedProperties(v []MicrosoftGraphSingleValueLegacyExtendedProperty)`

SetSingleValueExtendedProperties sets SingleValueExtendedProperties field to given value.

### HasSingleValueExtendedProperties

`func (o *MicrosoftGraphEvent) HasSingleValueExtendedProperties() bool`

HasSingleValueExtendedProperties returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


