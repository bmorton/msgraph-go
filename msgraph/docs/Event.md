# Event

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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

### NewEvent

`func NewEvent() *Event`

NewEvent instantiates a new Event object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventWithDefaults

`func NewEventWithDefaults() *Event`

NewEventWithDefaults instantiates a new Event object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowNewTimeProposals

`func (o *Event) GetAllowNewTimeProposals() bool`

GetAllowNewTimeProposals returns the AllowNewTimeProposals field if non-nil, zero value otherwise.

### GetAllowNewTimeProposalsOk

`func (o *Event) GetAllowNewTimeProposalsOk() (*bool, bool)`

GetAllowNewTimeProposalsOk returns a tuple with the AllowNewTimeProposals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowNewTimeProposals

`func (o *Event) SetAllowNewTimeProposals(v bool)`

SetAllowNewTimeProposals sets AllowNewTimeProposals field to given value.

### HasAllowNewTimeProposals

`func (o *Event) HasAllowNewTimeProposals() bool`

HasAllowNewTimeProposals returns a boolean if a field has been set.

### SetAllowNewTimeProposalsNil

`func (o *Event) SetAllowNewTimeProposalsNil(b bool)`

 SetAllowNewTimeProposalsNil sets the value for AllowNewTimeProposals to be an explicit nil

### UnsetAllowNewTimeProposals
`func (o *Event) UnsetAllowNewTimeProposals()`

UnsetAllowNewTimeProposals ensures that no value is present for AllowNewTimeProposals, not even an explicit nil
### GetAttendees

`func (o *Event) GetAttendees() []*AnyOfmicrosoftGraphAttendee`

GetAttendees returns the Attendees field if non-nil, zero value otherwise.

### GetAttendeesOk

`func (o *Event) GetAttendeesOk() (*[]*AnyOfmicrosoftGraphAttendee, bool)`

GetAttendeesOk returns a tuple with the Attendees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttendees

`func (o *Event) SetAttendees(v []*AnyOfmicrosoftGraphAttendee)`

SetAttendees sets Attendees field to given value.

### HasAttendees

`func (o *Event) HasAttendees() bool`

HasAttendees returns a boolean if a field has been set.

### GetBody

`func (o *Event) GetBody() AnyOfmicrosoftGraphItemBody`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *Event) GetBodyOk() (*AnyOfmicrosoftGraphItemBody, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *Event) SetBody(v AnyOfmicrosoftGraphItemBody)`

SetBody sets Body field to given value.

### HasBody

`func (o *Event) HasBody() bool`

HasBody returns a boolean if a field has been set.

### SetBodyNil

`func (o *Event) SetBodyNil(b bool)`

 SetBodyNil sets the value for Body to be an explicit nil

### UnsetBody
`func (o *Event) UnsetBody()`

UnsetBody ensures that no value is present for Body, not even an explicit nil
### GetBodyPreview

`func (o *Event) GetBodyPreview() string`

GetBodyPreview returns the BodyPreview field if non-nil, zero value otherwise.

### GetBodyPreviewOk

`func (o *Event) GetBodyPreviewOk() (*string, bool)`

GetBodyPreviewOk returns a tuple with the BodyPreview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBodyPreview

`func (o *Event) SetBodyPreview(v string)`

SetBodyPreview sets BodyPreview field to given value.

### HasBodyPreview

`func (o *Event) HasBodyPreview() bool`

HasBodyPreview returns a boolean if a field has been set.

### SetBodyPreviewNil

`func (o *Event) SetBodyPreviewNil(b bool)`

 SetBodyPreviewNil sets the value for BodyPreview to be an explicit nil

### UnsetBodyPreview
`func (o *Event) UnsetBodyPreview()`

UnsetBodyPreview ensures that no value is present for BodyPreview, not even an explicit nil
### GetEnd

`func (o *Event) GetEnd() AnyOfmicrosoftGraphDateTimeTimeZone`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *Event) GetEndOk() (*AnyOfmicrosoftGraphDateTimeTimeZone, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *Event) SetEnd(v AnyOfmicrosoftGraphDateTimeTimeZone)`

SetEnd sets End field to given value.

### HasEnd

`func (o *Event) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### SetEndNil

`func (o *Event) SetEndNil(b bool)`

 SetEndNil sets the value for End to be an explicit nil

### UnsetEnd
`func (o *Event) UnsetEnd()`

UnsetEnd ensures that no value is present for End, not even an explicit nil
### GetHasAttachments

`func (o *Event) GetHasAttachments() bool`

GetHasAttachments returns the HasAttachments field if non-nil, zero value otherwise.

### GetHasAttachmentsOk

`func (o *Event) GetHasAttachmentsOk() (*bool, bool)`

GetHasAttachmentsOk returns a tuple with the HasAttachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAttachments

`func (o *Event) SetHasAttachments(v bool)`

SetHasAttachments sets HasAttachments field to given value.

### HasHasAttachments

`func (o *Event) HasHasAttachments() bool`

HasHasAttachments returns a boolean if a field has been set.

### SetHasAttachmentsNil

`func (o *Event) SetHasAttachmentsNil(b bool)`

 SetHasAttachmentsNil sets the value for HasAttachments to be an explicit nil

### UnsetHasAttachments
`func (o *Event) UnsetHasAttachments()`

UnsetHasAttachments ensures that no value is present for HasAttachments, not even an explicit nil
### GetHideAttendees

`func (o *Event) GetHideAttendees() bool`

GetHideAttendees returns the HideAttendees field if non-nil, zero value otherwise.

### GetHideAttendeesOk

`func (o *Event) GetHideAttendeesOk() (*bool, bool)`

GetHideAttendeesOk returns a tuple with the HideAttendees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideAttendees

`func (o *Event) SetHideAttendees(v bool)`

SetHideAttendees sets HideAttendees field to given value.

### HasHideAttendees

`func (o *Event) HasHideAttendees() bool`

HasHideAttendees returns a boolean if a field has been set.

### SetHideAttendeesNil

`func (o *Event) SetHideAttendeesNil(b bool)`

 SetHideAttendeesNil sets the value for HideAttendees to be an explicit nil

### UnsetHideAttendees
`func (o *Event) UnsetHideAttendees()`

UnsetHideAttendees ensures that no value is present for HideAttendees, not even an explicit nil
### GetICalUId

`func (o *Event) GetICalUId() string`

GetICalUId returns the ICalUId field if non-nil, zero value otherwise.

### GetICalUIdOk

`func (o *Event) GetICalUIdOk() (*string, bool)`

GetICalUIdOk returns a tuple with the ICalUId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetICalUId

`func (o *Event) SetICalUId(v string)`

SetICalUId sets ICalUId field to given value.

### HasICalUId

`func (o *Event) HasICalUId() bool`

HasICalUId returns a boolean if a field has been set.

### SetICalUIdNil

`func (o *Event) SetICalUIdNil(b bool)`

 SetICalUIdNil sets the value for ICalUId to be an explicit nil

### UnsetICalUId
`func (o *Event) UnsetICalUId()`

UnsetICalUId ensures that no value is present for ICalUId, not even an explicit nil
### GetImportance

`func (o *Event) GetImportance() AnyOfmicrosoftGraphImportance`

GetImportance returns the Importance field if non-nil, zero value otherwise.

### GetImportanceOk

`func (o *Event) GetImportanceOk() (*AnyOfmicrosoftGraphImportance, bool)`

GetImportanceOk returns a tuple with the Importance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportance

`func (o *Event) SetImportance(v AnyOfmicrosoftGraphImportance)`

SetImportance sets Importance field to given value.

### HasImportance

`func (o *Event) HasImportance() bool`

HasImportance returns a boolean if a field has been set.

### SetImportanceNil

`func (o *Event) SetImportanceNil(b bool)`

 SetImportanceNil sets the value for Importance to be an explicit nil

### UnsetImportance
`func (o *Event) UnsetImportance()`

UnsetImportance ensures that no value is present for Importance, not even an explicit nil
### GetIsAllDay

`func (o *Event) GetIsAllDay() bool`

GetIsAllDay returns the IsAllDay field if non-nil, zero value otherwise.

### GetIsAllDayOk

`func (o *Event) GetIsAllDayOk() (*bool, bool)`

GetIsAllDayOk returns a tuple with the IsAllDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAllDay

`func (o *Event) SetIsAllDay(v bool)`

SetIsAllDay sets IsAllDay field to given value.

### HasIsAllDay

`func (o *Event) HasIsAllDay() bool`

HasIsAllDay returns a boolean if a field has been set.

### SetIsAllDayNil

`func (o *Event) SetIsAllDayNil(b bool)`

 SetIsAllDayNil sets the value for IsAllDay to be an explicit nil

### UnsetIsAllDay
`func (o *Event) UnsetIsAllDay()`

UnsetIsAllDay ensures that no value is present for IsAllDay, not even an explicit nil
### GetIsCancelled

`func (o *Event) GetIsCancelled() bool`

GetIsCancelled returns the IsCancelled field if non-nil, zero value otherwise.

### GetIsCancelledOk

`func (o *Event) GetIsCancelledOk() (*bool, bool)`

GetIsCancelledOk returns a tuple with the IsCancelled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCancelled

`func (o *Event) SetIsCancelled(v bool)`

SetIsCancelled sets IsCancelled field to given value.

### HasIsCancelled

`func (o *Event) HasIsCancelled() bool`

HasIsCancelled returns a boolean if a field has been set.

### SetIsCancelledNil

`func (o *Event) SetIsCancelledNil(b bool)`

 SetIsCancelledNil sets the value for IsCancelled to be an explicit nil

### UnsetIsCancelled
`func (o *Event) UnsetIsCancelled()`

UnsetIsCancelled ensures that no value is present for IsCancelled, not even an explicit nil
### GetIsDraft

`func (o *Event) GetIsDraft() bool`

GetIsDraft returns the IsDraft field if non-nil, zero value otherwise.

### GetIsDraftOk

`func (o *Event) GetIsDraftOk() (*bool, bool)`

GetIsDraftOk returns a tuple with the IsDraft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDraft

`func (o *Event) SetIsDraft(v bool)`

SetIsDraft sets IsDraft field to given value.

### HasIsDraft

`func (o *Event) HasIsDraft() bool`

HasIsDraft returns a boolean if a field has been set.

### SetIsDraftNil

`func (o *Event) SetIsDraftNil(b bool)`

 SetIsDraftNil sets the value for IsDraft to be an explicit nil

### UnsetIsDraft
`func (o *Event) UnsetIsDraft()`

UnsetIsDraft ensures that no value is present for IsDraft, not even an explicit nil
### GetIsOnlineMeeting

`func (o *Event) GetIsOnlineMeeting() bool`

GetIsOnlineMeeting returns the IsOnlineMeeting field if non-nil, zero value otherwise.

### GetIsOnlineMeetingOk

`func (o *Event) GetIsOnlineMeetingOk() (*bool, bool)`

GetIsOnlineMeetingOk returns a tuple with the IsOnlineMeeting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOnlineMeeting

`func (o *Event) SetIsOnlineMeeting(v bool)`

SetIsOnlineMeeting sets IsOnlineMeeting field to given value.

### HasIsOnlineMeeting

`func (o *Event) HasIsOnlineMeeting() bool`

HasIsOnlineMeeting returns a boolean if a field has been set.

### SetIsOnlineMeetingNil

`func (o *Event) SetIsOnlineMeetingNil(b bool)`

 SetIsOnlineMeetingNil sets the value for IsOnlineMeeting to be an explicit nil

### UnsetIsOnlineMeeting
`func (o *Event) UnsetIsOnlineMeeting()`

UnsetIsOnlineMeeting ensures that no value is present for IsOnlineMeeting, not even an explicit nil
### GetIsOrganizer

`func (o *Event) GetIsOrganizer() bool`

GetIsOrganizer returns the IsOrganizer field if non-nil, zero value otherwise.

### GetIsOrganizerOk

`func (o *Event) GetIsOrganizerOk() (*bool, bool)`

GetIsOrganizerOk returns a tuple with the IsOrganizer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOrganizer

`func (o *Event) SetIsOrganizer(v bool)`

SetIsOrganizer sets IsOrganizer field to given value.

### HasIsOrganizer

`func (o *Event) HasIsOrganizer() bool`

HasIsOrganizer returns a boolean if a field has been set.

### SetIsOrganizerNil

`func (o *Event) SetIsOrganizerNil(b bool)`

 SetIsOrganizerNil sets the value for IsOrganizer to be an explicit nil

### UnsetIsOrganizer
`func (o *Event) UnsetIsOrganizer()`

UnsetIsOrganizer ensures that no value is present for IsOrganizer, not even an explicit nil
### GetIsReminderOn

`func (o *Event) GetIsReminderOn() bool`

GetIsReminderOn returns the IsReminderOn field if non-nil, zero value otherwise.

### GetIsReminderOnOk

`func (o *Event) GetIsReminderOnOk() (*bool, bool)`

GetIsReminderOnOk returns a tuple with the IsReminderOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReminderOn

`func (o *Event) SetIsReminderOn(v bool)`

SetIsReminderOn sets IsReminderOn field to given value.

### HasIsReminderOn

`func (o *Event) HasIsReminderOn() bool`

HasIsReminderOn returns a boolean if a field has been set.

### SetIsReminderOnNil

`func (o *Event) SetIsReminderOnNil(b bool)`

 SetIsReminderOnNil sets the value for IsReminderOn to be an explicit nil

### UnsetIsReminderOn
`func (o *Event) UnsetIsReminderOn()`

UnsetIsReminderOn ensures that no value is present for IsReminderOn, not even an explicit nil
### GetLocation

`func (o *Event) GetLocation() AnyOfmicrosoftGraphLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *Event) GetLocationOk() (*AnyOfmicrosoftGraphLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *Event) SetLocation(v AnyOfmicrosoftGraphLocation)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *Event) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocationNil

`func (o *Event) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *Event) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil
### GetLocations

`func (o *Event) GetLocations() []*AnyOfmicrosoftGraphLocation`

GetLocations returns the Locations field if non-nil, zero value otherwise.

### GetLocationsOk

`func (o *Event) GetLocationsOk() (*[]*AnyOfmicrosoftGraphLocation, bool)`

GetLocationsOk returns a tuple with the Locations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocations

`func (o *Event) SetLocations(v []*AnyOfmicrosoftGraphLocation)`

SetLocations sets Locations field to given value.

### HasLocations

`func (o *Event) HasLocations() bool`

HasLocations returns a boolean if a field has been set.

### GetOnlineMeeting

`func (o *Event) GetOnlineMeeting() AnyOfmicrosoftGraphOnlineMeetingInfo`

GetOnlineMeeting returns the OnlineMeeting field if non-nil, zero value otherwise.

### GetOnlineMeetingOk

`func (o *Event) GetOnlineMeetingOk() (*AnyOfmicrosoftGraphOnlineMeetingInfo, bool)`

GetOnlineMeetingOk returns a tuple with the OnlineMeeting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnlineMeeting

`func (o *Event) SetOnlineMeeting(v AnyOfmicrosoftGraphOnlineMeetingInfo)`

SetOnlineMeeting sets OnlineMeeting field to given value.

### HasOnlineMeeting

`func (o *Event) HasOnlineMeeting() bool`

HasOnlineMeeting returns a boolean if a field has been set.

### SetOnlineMeetingNil

`func (o *Event) SetOnlineMeetingNil(b bool)`

 SetOnlineMeetingNil sets the value for OnlineMeeting to be an explicit nil

### UnsetOnlineMeeting
`func (o *Event) UnsetOnlineMeeting()`

UnsetOnlineMeeting ensures that no value is present for OnlineMeeting, not even an explicit nil
### GetOnlineMeetingProvider

`func (o *Event) GetOnlineMeetingProvider() AnyOfmicrosoftGraphOnlineMeetingProviderType`

GetOnlineMeetingProvider returns the OnlineMeetingProvider field if non-nil, zero value otherwise.

### GetOnlineMeetingProviderOk

`func (o *Event) GetOnlineMeetingProviderOk() (*AnyOfmicrosoftGraphOnlineMeetingProviderType, bool)`

GetOnlineMeetingProviderOk returns a tuple with the OnlineMeetingProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnlineMeetingProvider

`func (o *Event) SetOnlineMeetingProvider(v AnyOfmicrosoftGraphOnlineMeetingProviderType)`

SetOnlineMeetingProvider sets OnlineMeetingProvider field to given value.

### HasOnlineMeetingProvider

`func (o *Event) HasOnlineMeetingProvider() bool`

HasOnlineMeetingProvider returns a boolean if a field has been set.

### SetOnlineMeetingProviderNil

`func (o *Event) SetOnlineMeetingProviderNil(b bool)`

 SetOnlineMeetingProviderNil sets the value for OnlineMeetingProvider to be an explicit nil

### UnsetOnlineMeetingProvider
`func (o *Event) UnsetOnlineMeetingProvider()`

UnsetOnlineMeetingProvider ensures that no value is present for OnlineMeetingProvider, not even an explicit nil
### GetOnlineMeetingUrl

`func (o *Event) GetOnlineMeetingUrl() string`

GetOnlineMeetingUrl returns the OnlineMeetingUrl field if non-nil, zero value otherwise.

### GetOnlineMeetingUrlOk

`func (o *Event) GetOnlineMeetingUrlOk() (*string, bool)`

GetOnlineMeetingUrlOk returns a tuple with the OnlineMeetingUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnlineMeetingUrl

`func (o *Event) SetOnlineMeetingUrl(v string)`

SetOnlineMeetingUrl sets OnlineMeetingUrl field to given value.

### HasOnlineMeetingUrl

`func (o *Event) HasOnlineMeetingUrl() bool`

HasOnlineMeetingUrl returns a boolean if a field has been set.

### SetOnlineMeetingUrlNil

`func (o *Event) SetOnlineMeetingUrlNil(b bool)`

 SetOnlineMeetingUrlNil sets the value for OnlineMeetingUrl to be an explicit nil

### UnsetOnlineMeetingUrl
`func (o *Event) UnsetOnlineMeetingUrl()`

UnsetOnlineMeetingUrl ensures that no value is present for OnlineMeetingUrl, not even an explicit nil
### GetOrganizer

`func (o *Event) GetOrganizer() AnyOfmicrosoftGraphRecipient`

GetOrganizer returns the Organizer field if non-nil, zero value otherwise.

### GetOrganizerOk

`func (o *Event) GetOrganizerOk() (*AnyOfmicrosoftGraphRecipient, bool)`

GetOrganizerOk returns a tuple with the Organizer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizer

`func (o *Event) SetOrganizer(v AnyOfmicrosoftGraphRecipient)`

SetOrganizer sets Organizer field to given value.

### HasOrganizer

`func (o *Event) HasOrganizer() bool`

HasOrganizer returns a boolean if a field has been set.

### SetOrganizerNil

`func (o *Event) SetOrganizerNil(b bool)`

 SetOrganizerNil sets the value for Organizer to be an explicit nil

### UnsetOrganizer
`func (o *Event) UnsetOrganizer()`

UnsetOrganizer ensures that no value is present for Organizer, not even an explicit nil
### GetOriginalEndTimeZone

`func (o *Event) GetOriginalEndTimeZone() string`

GetOriginalEndTimeZone returns the OriginalEndTimeZone field if non-nil, zero value otherwise.

### GetOriginalEndTimeZoneOk

`func (o *Event) GetOriginalEndTimeZoneOk() (*string, bool)`

GetOriginalEndTimeZoneOk returns a tuple with the OriginalEndTimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalEndTimeZone

`func (o *Event) SetOriginalEndTimeZone(v string)`

SetOriginalEndTimeZone sets OriginalEndTimeZone field to given value.

### HasOriginalEndTimeZone

`func (o *Event) HasOriginalEndTimeZone() bool`

HasOriginalEndTimeZone returns a boolean if a field has been set.

### SetOriginalEndTimeZoneNil

`func (o *Event) SetOriginalEndTimeZoneNil(b bool)`

 SetOriginalEndTimeZoneNil sets the value for OriginalEndTimeZone to be an explicit nil

### UnsetOriginalEndTimeZone
`func (o *Event) UnsetOriginalEndTimeZone()`

UnsetOriginalEndTimeZone ensures that no value is present for OriginalEndTimeZone, not even an explicit nil
### GetOriginalStart

`func (o *Event) GetOriginalStart() time.Time`

GetOriginalStart returns the OriginalStart field if non-nil, zero value otherwise.

### GetOriginalStartOk

`func (o *Event) GetOriginalStartOk() (*time.Time, bool)`

GetOriginalStartOk returns a tuple with the OriginalStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalStart

`func (o *Event) SetOriginalStart(v time.Time)`

SetOriginalStart sets OriginalStart field to given value.

### HasOriginalStart

`func (o *Event) HasOriginalStart() bool`

HasOriginalStart returns a boolean if a field has been set.

### SetOriginalStartNil

`func (o *Event) SetOriginalStartNil(b bool)`

 SetOriginalStartNil sets the value for OriginalStart to be an explicit nil

### UnsetOriginalStart
`func (o *Event) UnsetOriginalStart()`

UnsetOriginalStart ensures that no value is present for OriginalStart, not even an explicit nil
### GetOriginalStartTimeZone

`func (o *Event) GetOriginalStartTimeZone() string`

GetOriginalStartTimeZone returns the OriginalStartTimeZone field if non-nil, zero value otherwise.

### GetOriginalStartTimeZoneOk

`func (o *Event) GetOriginalStartTimeZoneOk() (*string, bool)`

GetOriginalStartTimeZoneOk returns a tuple with the OriginalStartTimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalStartTimeZone

`func (o *Event) SetOriginalStartTimeZone(v string)`

SetOriginalStartTimeZone sets OriginalStartTimeZone field to given value.

### HasOriginalStartTimeZone

`func (o *Event) HasOriginalStartTimeZone() bool`

HasOriginalStartTimeZone returns a boolean if a field has been set.

### SetOriginalStartTimeZoneNil

`func (o *Event) SetOriginalStartTimeZoneNil(b bool)`

 SetOriginalStartTimeZoneNil sets the value for OriginalStartTimeZone to be an explicit nil

### UnsetOriginalStartTimeZone
`func (o *Event) UnsetOriginalStartTimeZone()`

UnsetOriginalStartTimeZone ensures that no value is present for OriginalStartTimeZone, not even an explicit nil
### GetRecurrence

`func (o *Event) GetRecurrence() AnyOfmicrosoftGraphPatternedRecurrence`

GetRecurrence returns the Recurrence field if non-nil, zero value otherwise.

### GetRecurrenceOk

`func (o *Event) GetRecurrenceOk() (*AnyOfmicrosoftGraphPatternedRecurrence, bool)`

GetRecurrenceOk returns a tuple with the Recurrence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurrence

`func (o *Event) SetRecurrence(v AnyOfmicrosoftGraphPatternedRecurrence)`

SetRecurrence sets Recurrence field to given value.

### HasRecurrence

`func (o *Event) HasRecurrence() bool`

HasRecurrence returns a boolean if a field has been set.

### SetRecurrenceNil

`func (o *Event) SetRecurrenceNil(b bool)`

 SetRecurrenceNil sets the value for Recurrence to be an explicit nil

### UnsetRecurrence
`func (o *Event) UnsetRecurrence()`

UnsetRecurrence ensures that no value is present for Recurrence, not even an explicit nil
### GetReminderMinutesBeforeStart

`func (o *Event) GetReminderMinutesBeforeStart() int32`

GetReminderMinutesBeforeStart returns the ReminderMinutesBeforeStart field if non-nil, zero value otherwise.

### GetReminderMinutesBeforeStartOk

`func (o *Event) GetReminderMinutesBeforeStartOk() (*int32, bool)`

GetReminderMinutesBeforeStartOk returns a tuple with the ReminderMinutesBeforeStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReminderMinutesBeforeStart

`func (o *Event) SetReminderMinutesBeforeStart(v int32)`

SetReminderMinutesBeforeStart sets ReminderMinutesBeforeStart field to given value.

### HasReminderMinutesBeforeStart

`func (o *Event) HasReminderMinutesBeforeStart() bool`

HasReminderMinutesBeforeStart returns a boolean if a field has been set.

### SetReminderMinutesBeforeStartNil

`func (o *Event) SetReminderMinutesBeforeStartNil(b bool)`

 SetReminderMinutesBeforeStartNil sets the value for ReminderMinutesBeforeStart to be an explicit nil

### UnsetReminderMinutesBeforeStart
`func (o *Event) UnsetReminderMinutesBeforeStart()`

UnsetReminderMinutesBeforeStart ensures that no value is present for ReminderMinutesBeforeStart, not even an explicit nil
### GetResponseRequested

`func (o *Event) GetResponseRequested() bool`

GetResponseRequested returns the ResponseRequested field if non-nil, zero value otherwise.

### GetResponseRequestedOk

`func (o *Event) GetResponseRequestedOk() (*bool, bool)`

GetResponseRequestedOk returns a tuple with the ResponseRequested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseRequested

`func (o *Event) SetResponseRequested(v bool)`

SetResponseRequested sets ResponseRequested field to given value.

### HasResponseRequested

`func (o *Event) HasResponseRequested() bool`

HasResponseRequested returns a boolean if a field has been set.

### SetResponseRequestedNil

`func (o *Event) SetResponseRequestedNil(b bool)`

 SetResponseRequestedNil sets the value for ResponseRequested to be an explicit nil

### UnsetResponseRequested
`func (o *Event) UnsetResponseRequested()`

UnsetResponseRequested ensures that no value is present for ResponseRequested, not even an explicit nil
### GetResponseStatus

`func (o *Event) GetResponseStatus() AnyOfmicrosoftGraphResponseStatus`

GetResponseStatus returns the ResponseStatus field if non-nil, zero value otherwise.

### GetResponseStatusOk

`func (o *Event) GetResponseStatusOk() (*AnyOfmicrosoftGraphResponseStatus, bool)`

GetResponseStatusOk returns a tuple with the ResponseStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseStatus

`func (o *Event) SetResponseStatus(v AnyOfmicrosoftGraphResponseStatus)`

SetResponseStatus sets ResponseStatus field to given value.

### HasResponseStatus

`func (o *Event) HasResponseStatus() bool`

HasResponseStatus returns a boolean if a field has been set.

### SetResponseStatusNil

`func (o *Event) SetResponseStatusNil(b bool)`

 SetResponseStatusNil sets the value for ResponseStatus to be an explicit nil

### UnsetResponseStatus
`func (o *Event) UnsetResponseStatus()`

UnsetResponseStatus ensures that no value is present for ResponseStatus, not even an explicit nil
### GetSensitivity

`func (o *Event) GetSensitivity() AnyOfmicrosoftGraphSensitivity`

GetSensitivity returns the Sensitivity field if non-nil, zero value otherwise.

### GetSensitivityOk

`func (o *Event) GetSensitivityOk() (*AnyOfmicrosoftGraphSensitivity, bool)`

GetSensitivityOk returns a tuple with the Sensitivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSensitivity

`func (o *Event) SetSensitivity(v AnyOfmicrosoftGraphSensitivity)`

SetSensitivity sets Sensitivity field to given value.

### HasSensitivity

`func (o *Event) HasSensitivity() bool`

HasSensitivity returns a boolean if a field has been set.

### SetSensitivityNil

`func (o *Event) SetSensitivityNil(b bool)`

 SetSensitivityNil sets the value for Sensitivity to be an explicit nil

### UnsetSensitivity
`func (o *Event) UnsetSensitivity()`

UnsetSensitivity ensures that no value is present for Sensitivity, not even an explicit nil
### GetSeriesMasterId

`func (o *Event) GetSeriesMasterId() string`

GetSeriesMasterId returns the SeriesMasterId field if non-nil, zero value otherwise.

### GetSeriesMasterIdOk

`func (o *Event) GetSeriesMasterIdOk() (*string, bool)`

GetSeriesMasterIdOk returns a tuple with the SeriesMasterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeriesMasterId

`func (o *Event) SetSeriesMasterId(v string)`

SetSeriesMasterId sets SeriesMasterId field to given value.

### HasSeriesMasterId

`func (o *Event) HasSeriesMasterId() bool`

HasSeriesMasterId returns a boolean if a field has been set.

### SetSeriesMasterIdNil

`func (o *Event) SetSeriesMasterIdNil(b bool)`

 SetSeriesMasterIdNil sets the value for SeriesMasterId to be an explicit nil

### UnsetSeriesMasterId
`func (o *Event) UnsetSeriesMasterId()`

UnsetSeriesMasterId ensures that no value is present for SeriesMasterId, not even an explicit nil
### GetShowAs

`func (o *Event) GetShowAs() AnyOfmicrosoftGraphFreeBusyStatus`

GetShowAs returns the ShowAs field if non-nil, zero value otherwise.

### GetShowAsOk

`func (o *Event) GetShowAsOk() (*AnyOfmicrosoftGraphFreeBusyStatus, bool)`

GetShowAsOk returns a tuple with the ShowAs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowAs

`func (o *Event) SetShowAs(v AnyOfmicrosoftGraphFreeBusyStatus)`

SetShowAs sets ShowAs field to given value.

### HasShowAs

`func (o *Event) HasShowAs() bool`

HasShowAs returns a boolean if a field has been set.

### SetShowAsNil

`func (o *Event) SetShowAsNil(b bool)`

 SetShowAsNil sets the value for ShowAs to be an explicit nil

### UnsetShowAs
`func (o *Event) UnsetShowAs()`

UnsetShowAs ensures that no value is present for ShowAs, not even an explicit nil
### GetStart

`func (o *Event) GetStart() AnyOfmicrosoftGraphDateTimeTimeZone`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *Event) GetStartOk() (*AnyOfmicrosoftGraphDateTimeTimeZone, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *Event) SetStart(v AnyOfmicrosoftGraphDateTimeTimeZone)`

SetStart sets Start field to given value.

### HasStart

`func (o *Event) HasStart() bool`

HasStart returns a boolean if a field has been set.

### SetStartNil

`func (o *Event) SetStartNil(b bool)`

 SetStartNil sets the value for Start to be an explicit nil

### UnsetStart
`func (o *Event) UnsetStart()`

UnsetStart ensures that no value is present for Start, not even an explicit nil
### GetSubject

`func (o *Event) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *Event) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *Event) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *Event) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### SetSubjectNil

`func (o *Event) SetSubjectNil(b bool)`

 SetSubjectNil sets the value for Subject to be an explicit nil

### UnsetSubject
`func (o *Event) UnsetSubject()`

UnsetSubject ensures that no value is present for Subject, not even an explicit nil
### GetTransactionId

`func (o *Event) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *Event) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *Event) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.

### HasTransactionId

`func (o *Event) HasTransactionId() bool`

HasTransactionId returns a boolean if a field has been set.

### SetTransactionIdNil

`func (o *Event) SetTransactionIdNil(b bool)`

 SetTransactionIdNil sets the value for TransactionId to be an explicit nil

### UnsetTransactionId
`func (o *Event) UnsetTransactionId()`

UnsetTransactionId ensures that no value is present for TransactionId, not even an explicit nil
### GetType

`func (o *Event) GetType() AnyOfmicrosoftGraphEventType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Event) GetTypeOk() (*AnyOfmicrosoftGraphEventType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Event) SetType(v AnyOfmicrosoftGraphEventType)`

SetType sets Type field to given value.

### HasType

`func (o *Event) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *Event) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *Event) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetWebLink

`func (o *Event) GetWebLink() string`

GetWebLink returns the WebLink field if non-nil, zero value otherwise.

### GetWebLinkOk

`func (o *Event) GetWebLinkOk() (*string, bool)`

GetWebLinkOk returns a tuple with the WebLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebLink

`func (o *Event) SetWebLink(v string)`

SetWebLink sets WebLink field to given value.

### HasWebLink

`func (o *Event) HasWebLink() bool`

HasWebLink returns a boolean if a field has been set.

### SetWebLinkNil

`func (o *Event) SetWebLinkNil(b bool)`

 SetWebLinkNil sets the value for WebLink to be an explicit nil

### UnsetWebLink
`func (o *Event) UnsetWebLink()`

UnsetWebLink ensures that no value is present for WebLink, not even an explicit nil
### GetAttachments

`func (o *Event) GetAttachments() []MicrosoftGraphAttachment`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *Event) GetAttachmentsOk() (*[]MicrosoftGraphAttachment, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *Event) SetAttachments(v []MicrosoftGraphAttachment)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *Event) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetCalendar

`func (o *Event) GetCalendar() AnyOfmicrosoftGraphCalendar`

GetCalendar returns the Calendar field if non-nil, zero value otherwise.

### GetCalendarOk

`func (o *Event) GetCalendarOk() (*AnyOfmicrosoftGraphCalendar, bool)`

GetCalendarOk returns a tuple with the Calendar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalendar

`func (o *Event) SetCalendar(v AnyOfmicrosoftGraphCalendar)`

SetCalendar sets Calendar field to given value.

### HasCalendar

`func (o *Event) HasCalendar() bool`

HasCalendar returns a boolean if a field has been set.

### SetCalendarNil

`func (o *Event) SetCalendarNil(b bool)`

 SetCalendarNil sets the value for Calendar to be an explicit nil

### UnsetCalendar
`func (o *Event) UnsetCalendar()`

UnsetCalendar ensures that no value is present for Calendar, not even an explicit nil
### GetExtensions

`func (o *Event) GetExtensions() []MicrosoftGraphExtension`

GetExtensions returns the Extensions field if non-nil, zero value otherwise.

### GetExtensionsOk

`func (o *Event) GetExtensionsOk() (*[]MicrosoftGraphExtension, bool)`

GetExtensionsOk returns a tuple with the Extensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensions

`func (o *Event) SetExtensions(v []MicrosoftGraphExtension)`

SetExtensions sets Extensions field to given value.

### HasExtensions

`func (o *Event) HasExtensions() bool`

HasExtensions returns a boolean if a field has been set.

### GetInstances

`func (o *Event) GetInstances() []MicrosoftGraphEvent`

GetInstances returns the Instances field if non-nil, zero value otherwise.

### GetInstancesOk

`func (o *Event) GetInstancesOk() (*[]MicrosoftGraphEvent, bool)`

GetInstancesOk returns a tuple with the Instances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstances

`func (o *Event) SetInstances(v []MicrosoftGraphEvent)`

SetInstances sets Instances field to given value.

### HasInstances

`func (o *Event) HasInstances() bool`

HasInstances returns a boolean if a field has been set.

### GetMultiValueExtendedProperties

`func (o *Event) GetMultiValueExtendedProperties() []MicrosoftGraphMultiValueLegacyExtendedProperty`

GetMultiValueExtendedProperties returns the MultiValueExtendedProperties field if non-nil, zero value otherwise.

### GetMultiValueExtendedPropertiesOk

`func (o *Event) GetMultiValueExtendedPropertiesOk() (*[]MicrosoftGraphMultiValueLegacyExtendedProperty, bool)`

GetMultiValueExtendedPropertiesOk returns a tuple with the MultiValueExtendedProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiValueExtendedProperties

`func (o *Event) SetMultiValueExtendedProperties(v []MicrosoftGraphMultiValueLegacyExtendedProperty)`

SetMultiValueExtendedProperties sets MultiValueExtendedProperties field to given value.

### HasMultiValueExtendedProperties

`func (o *Event) HasMultiValueExtendedProperties() bool`

HasMultiValueExtendedProperties returns a boolean if a field has been set.

### GetSingleValueExtendedProperties

`func (o *Event) GetSingleValueExtendedProperties() []MicrosoftGraphSingleValueLegacyExtendedProperty`

GetSingleValueExtendedProperties returns the SingleValueExtendedProperties field if non-nil, zero value otherwise.

### GetSingleValueExtendedPropertiesOk

`func (o *Event) GetSingleValueExtendedPropertiesOk() (*[]MicrosoftGraphSingleValueLegacyExtendedProperty, bool)`

GetSingleValueExtendedPropertiesOk returns a tuple with the SingleValueExtendedProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSingleValueExtendedProperties

`func (o *Event) SetSingleValueExtendedProperties(v []MicrosoftGraphSingleValueLegacyExtendedProperty)`

SetSingleValueExtendedProperties sets SingleValueExtendedProperties field to given value.

### HasSingleValueExtendedProperties

`func (o *Event) HasSingleValueExtendedProperties() bool`

HasSingleValueExtendedProperties returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


