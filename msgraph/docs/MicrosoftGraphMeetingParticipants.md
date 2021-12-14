# MicrosoftGraphMeetingParticipants

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attendees** | Pointer to [**[]AnyOfmicrosoftGraphMeetingParticipantInfo**](AnyOfmicrosoftGraphMeetingParticipantInfo.md) | Information of the meeting attendees. | [optional] 
**Organizer** | Pointer to [**AnyOfmicrosoftGraphMeetingParticipantInfo**](anyOf&lt;microsoft.graph.meetingParticipantInfo&gt;.md) | Information of the meeting organizer. | [optional] 

## Methods

### NewMicrosoftGraphMeetingParticipants

`func NewMicrosoftGraphMeetingParticipants() *MicrosoftGraphMeetingParticipants`

NewMicrosoftGraphMeetingParticipants instantiates a new MicrosoftGraphMeetingParticipants object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphMeetingParticipantsWithDefaults

`func NewMicrosoftGraphMeetingParticipantsWithDefaults() *MicrosoftGraphMeetingParticipants`

NewMicrosoftGraphMeetingParticipantsWithDefaults instantiates a new MicrosoftGraphMeetingParticipants object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttendees

`func (o *MicrosoftGraphMeetingParticipants) GetAttendees() []*AnyOfmicrosoftGraphMeetingParticipantInfo`

GetAttendees returns the Attendees field if non-nil, zero value otherwise.

### GetAttendeesOk

`func (o *MicrosoftGraphMeetingParticipants) GetAttendeesOk() (*[]*AnyOfmicrosoftGraphMeetingParticipantInfo, bool)`

GetAttendeesOk returns a tuple with the Attendees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttendees

`func (o *MicrosoftGraphMeetingParticipants) SetAttendees(v []*AnyOfmicrosoftGraphMeetingParticipantInfo)`

SetAttendees sets Attendees field to given value.

### HasAttendees

`func (o *MicrosoftGraphMeetingParticipants) HasAttendees() bool`

HasAttendees returns a boolean if a field has been set.

### GetOrganizer

`func (o *MicrosoftGraphMeetingParticipants) GetOrganizer() AnyOfmicrosoftGraphMeetingParticipantInfo`

GetOrganizer returns the Organizer field if non-nil, zero value otherwise.

### GetOrganizerOk

`func (o *MicrosoftGraphMeetingParticipants) GetOrganizerOk() (*AnyOfmicrosoftGraphMeetingParticipantInfo, bool)`

GetOrganizerOk returns a tuple with the Organizer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizer

`func (o *MicrosoftGraphMeetingParticipants) SetOrganizer(v AnyOfmicrosoftGraphMeetingParticipantInfo)`

SetOrganizer sets Organizer field to given value.

### HasOrganizer

`func (o *MicrosoftGraphMeetingParticipants) HasOrganizer() bool`

HasOrganizer returns a boolean if a field has been set.

### SetOrganizerNil

`func (o *MicrosoftGraphMeetingParticipants) SetOrganizerNil(b bool)`

 SetOrganizerNil sets the value for Organizer to be an explicit nil

### UnsetOrganizer
`func (o *MicrosoftGraphMeetingParticipants) UnsetOrganizer()`

UnsetOrganizer ensures that no value is present for Organizer, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


