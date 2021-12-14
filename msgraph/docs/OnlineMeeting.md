# OnlineMeeting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowAttendeeToEnableCamera** | Pointer to **NullableBool** | Indicates whether attendees can turn on their camera. | [optional] 
**AllowAttendeeToEnableMic** | Pointer to **NullableBool** | Indicates whether attendees can turn on their microphone. | [optional] 
**AllowedPresenters** | Pointer to [**AnyOfmicrosoftGraphOnlineMeetingPresenters**](anyOf&lt;microsoft.graph.onlineMeetingPresenters&gt;.md) | Specifies who can be a presenter in a meeting. Possible values are listed in the following table. | [optional] 
**AllowMeetingChat** | Pointer to [**AnyOfmicrosoftGraphMeetingChatMode**](anyOf&lt;microsoft.graph.meetingChatMode&gt;.md) | Specifies the mode of meeting chat. | [optional] 
**AllowTeamworkReactions** | Pointer to **NullableBool** | Indicates whether Teams reactions are enabled for the meeting. | [optional] 
**AttendeeReport** | Pointer to **NullableString** | The content stream of the attendee report of a Microsoft Teams live event. Read-only. | [optional] 
**AudioConferencing** | Pointer to [**AnyOfmicrosoftGraphAudioConferencing**](anyOf&lt;microsoft.graph.audioConferencing&gt;.md) | The phone access (dial-in) information for an online meeting. Read-only. | [optional] 
**BroadcastSettings** | Pointer to [**AnyOfmicrosoftGraphBroadcastMeetingSettings**](anyOf&lt;microsoft.graph.broadcastMeetingSettings&gt;.md) | Settings related to a live event. | [optional] 
**ChatInfo** | Pointer to [**AnyOfmicrosoftGraphChatInfo**](anyOf&lt;microsoft.graph.chatInfo&gt;.md) | The chat information associated with this online meeting. | [optional] 
**CreationDateTime** | Pointer to **NullableTime** | The meeting creation time in UTC. Read-only. | [optional] 
**EndDateTime** | Pointer to **NullableTime** | The meeting end time in UTC. | [optional] 
**ExternalId** | Pointer to **NullableString** | The external ID. A custom ID. Optional. | [optional] 
**IsBroadcast** | Pointer to **NullableBool** | Indicates if this is a Teams live event. | [optional] 
**IsEntryExitAnnounced** | Pointer to **NullableBool** | Indicates whether to announce when callers join or leave. | [optional] 
**JoinInformation** | Pointer to [**AnyOfmicrosoftGraphItemBody**](anyOf&lt;microsoft.graph.itemBody&gt;.md) | The join information in the language and locale variant specified in the Accept-Language request HTTP header. Read-only. | [optional] 
**JoinWebUrl** | Pointer to **NullableString** | The join URL of the online meeting. Read-only. | [optional] 
**LobbyBypassSettings** | Pointer to [**AnyOfmicrosoftGraphLobbyBypassSettings**](anyOf&lt;microsoft.graph.lobbyBypassSettings&gt;.md) | Specifies which participants can bypass the meeting   lobby. | [optional] 
**Participants** | Pointer to [**AnyOfmicrosoftGraphMeetingParticipants**](anyOf&lt;microsoft.graph.meetingParticipants&gt;.md) | The participants associated with the online meeting.  This includes the organizer and the attendees. | [optional] 
**RecordAutomatically** | Pointer to **NullableBool** | Indicates whether to record the meeting automatically. | [optional] 
**StartDateTime** | Pointer to **NullableTime** | The meeting start time in UTC. | [optional] 
**Subject** | Pointer to **NullableString** | The subject of the online meeting. | [optional] 
**VideoTeleconferenceId** | Pointer to **NullableString** | The video teleconferencing ID. Read-only. | [optional] 

## Methods

### NewOnlineMeeting

`func NewOnlineMeeting() *OnlineMeeting`

NewOnlineMeeting instantiates a new OnlineMeeting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOnlineMeetingWithDefaults

`func NewOnlineMeetingWithDefaults() *OnlineMeeting`

NewOnlineMeetingWithDefaults instantiates a new OnlineMeeting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowAttendeeToEnableCamera

`func (o *OnlineMeeting) GetAllowAttendeeToEnableCamera() bool`

GetAllowAttendeeToEnableCamera returns the AllowAttendeeToEnableCamera field if non-nil, zero value otherwise.

### GetAllowAttendeeToEnableCameraOk

`func (o *OnlineMeeting) GetAllowAttendeeToEnableCameraOk() (*bool, bool)`

GetAllowAttendeeToEnableCameraOk returns a tuple with the AllowAttendeeToEnableCamera field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowAttendeeToEnableCamera

`func (o *OnlineMeeting) SetAllowAttendeeToEnableCamera(v bool)`

SetAllowAttendeeToEnableCamera sets AllowAttendeeToEnableCamera field to given value.

### HasAllowAttendeeToEnableCamera

`func (o *OnlineMeeting) HasAllowAttendeeToEnableCamera() bool`

HasAllowAttendeeToEnableCamera returns a boolean if a field has been set.

### SetAllowAttendeeToEnableCameraNil

`func (o *OnlineMeeting) SetAllowAttendeeToEnableCameraNil(b bool)`

 SetAllowAttendeeToEnableCameraNil sets the value for AllowAttendeeToEnableCamera to be an explicit nil

### UnsetAllowAttendeeToEnableCamera
`func (o *OnlineMeeting) UnsetAllowAttendeeToEnableCamera()`

UnsetAllowAttendeeToEnableCamera ensures that no value is present for AllowAttendeeToEnableCamera, not even an explicit nil
### GetAllowAttendeeToEnableMic

`func (o *OnlineMeeting) GetAllowAttendeeToEnableMic() bool`

GetAllowAttendeeToEnableMic returns the AllowAttendeeToEnableMic field if non-nil, zero value otherwise.

### GetAllowAttendeeToEnableMicOk

`func (o *OnlineMeeting) GetAllowAttendeeToEnableMicOk() (*bool, bool)`

GetAllowAttendeeToEnableMicOk returns a tuple with the AllowAttendeeToEnableMic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowAttendeeToEnableMic

`func (o *OnlineMeeting) SetAllowAttendeeToEnableMic(v bool)`

SetAllowAttendeeToEnableMic sets AllowAttendeeToEnableMic field to given value.

### HasAllowAttendeeToEnableMic

`func (o *OnlineMeeting) HasAllowAttendeeToEnableMic() bool`

HasAllowAttendeeToEnableMic returns a boolean if a field has been set.

### SetAllowAttendeeToEnableMicNil

`func (o *OnlineMeeting) SetAllowAttendeeToEnableMicNil(b bool)`

 SetAllowAttendeeToEnableMicNil sets the value for AllowAttendeeToEnableMic to be an explicit nil

### UnsetAllowAttendeeToEnableMic
`func (o *OnlineMeeting) UnsetAllowAttendeeToEnableMic()`

UnsetAllowAttendeeToEnableMic ensures that no value is present for AllowAttendeeToEnableMic, not even an explicit nil
### GetAllowedPresenters

`func (o *OnlineMeeting) GetAllowedPresenters() AnyOfmicrosoftGraphOnlineMeetingPresenters`

GetAllowedPresenters returns the AllowedPresenters field if non-nil, zero value otherwise.

### GetAllowedPresentersOk

`func (o *OnlineMeeting) GetAllowedPresentersOk() (*AnyOfmicrosoftGraphOnlineMeetingPresenters, bool)`

GetAllowedPresentersOk returns a tuple with the AllowedPresenters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedPresenters

`func (o *OnlineMeeting) SetAllowedPresenters(v AnyOfmicrosoftGraphOnlineMeetingPresenters)`

SetAllowedPresenters sets AllowedPresenters field to given value.

### HasAllowedPresenters

`func (o *OnlineMeeting) HasAllowedPresenters() bool`

HasAllowedPresenters returns a boolean if a field has been set.

### SetAllowedPresentersNil

`func (o *OnlineMeeting) SetAllowedPresentersNil(b bool)`

 SetAllowedPresentersNil sets the value for AllowedPresenters to be an explicit nil

### UnsetAllowedPresenters
`func (o *OnlineMeeting) UnsetAllowedPresenters()`

UnsetAllowedPresenters ensures that no value is present for AllowedPresenters, not even an explicit nil
### GetAllowMeetingChat

`func (o *OnlineMeeting) GetAllowMeetingChat() AnyOfmicrosoftGraphMeetingChatMode`

GetAllowMeetingChat returns the AllowMeetingChat field if non-nil, zero value otherwise.

### GetAllowMeetingChatOk

`func (o *OnlineMeeting) GetAllowMeetingChatOk() (*AnyOfmicrosoftGraphMeetingChatMode, bool)`

GetAllowMeetingChatOk returns a tuple with the AllowMeetingChat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowMeetingChat

`func (o *OnlineMeeting) SetAllowMeetingChat(v AnyOfmicrosoftGraphMeetingChatMode)`

SetAllowMeetingChat sets AllowMeetingChat field to given value.

### HasAllowMeetingChat

`func (o *OnlineMeeting) HasAllowMeetingChat() bool`

HasAllowMeetingChat returns a boolean if a field has been set.

### SetAllowMeetingChatNil

`func (o *OnlineMeeting) SetAllowMeetingChatNil(b bool)`

 SetAllowMeetingChatNil sets the value for AllowMeetingChat to be an explicit nil

### UnsetAllowMeetingChat
`func (o *OnlineMeeting) UnsetAllowMeetingChat()`

UnsetAllowMeetingChat ensures that no value is present for AllowMeetingChat, not even an explicit nil
### GetAllowTeamworkReactions

`func (o *OnlineMeeting) GetAllowTeamworkReactions() bool`

GetAllowTeamworkReactions returns the AllowTeamworkReactions field if non-nil, zero value otherwise.

### GetAllowTeamworkReactionsOk

`func (o *OnlineMeeting) GetAllowTeamworkReactionsOk() (*bool, bool)`

GetAllowTeamworkReactionsOk returns a tuple with the AllowTeamworkReactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowTeamworkReactions

`func (o *OnlineMeeting) SetAllowTeamworkReactions(v bool)`

SetAllowTeamworkReactions sets AllowTeamworkReactions field to given value.

### HasAllowTeamworkReactions

`func (o *OnlineMeeting) HasAllowTeamworkReactions() bool`

HasAllowTeamworkReactions returns a boolean if a field has been set.

### SetAllowTeamworkReactionsNil

`func (o *OnlineMeeting) SetAllowTeamworkReactionsNil(b bool)`

 SetAllowTeamworkReactionsNil sets the value for AllowTeamworkReactions to be an explicit nil

### UnsetAllowTeamworkReactions
`func (o *OnlineMeeting) UnsetAllowTeamworkReactions()`

UnsetAllowTeamworkReactions ensures that no value is present for AllowTeamworkReactions, not even an explicit nil
### GetAttendeeReport

`func (o *OnlineMeeting) GetAttendeeReport() string`

GetAttendeeReport returns the AttendeeReport field if non-nil, zero value otherwise.

### GetAttendeeReportOk

`func (o *OnlineMeeting) GetAttendeeReportOk() (*string, bool)`

GetAttendeeReportOk returns a tuple with the AttendeeReport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttendeeReport

`func (o *OnlineMeeting) SetAttendeeReport(v string)`

SetAttendeeReport sets AttendeeReport field to given value.

### HasAttendeeReport

`func (o *OnlineMeeting) HasAttendeeReport() bool`

HasAttendeeReport returns a boolean if a field has been set.

### SetAttendeeReportNil

`func (o *OnlineMeeting) SetAttendeeReportNil(b bool)`

 SetAttendeeReportNil sets the value for AttendeeReport to be an explicit nil

### UnsetAttendeeReport
`func (o *OnlineMeeting) UnsetAttendeeReport()`

UnsetAttendeeReport ensures that no value is present for AttendeeReport, not even an explicit nil
### GetAudioConferencing

`func (o *OnlineMeeting) GetAudioConferencing() AnyOfmicrosoftGraphAudioConferencing`

GetAudioConferencing returns the AudioConferencing field if non-nil, zero value otherwise.

### GetAudioConferencingOk

`func (o *OnlineMeeting) GetAudioConferencingOk() (*AnyOfmicrosoftGraphAudioConferencing, bool)`

GetAudioConferencingOk returns a tuple with the AudioConferencing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioConferencing

`func (o *OnlineMeeting) SetAudioConferencing(v AnyOfmicrosoftGraphAudioConferencing)`

SetAudioConferencing sets AudioConferencing field to given value.

### HasAudioConferencing

`func (o *OnlineMeeting) HasAudioConferencing() bool`

HasAudioConferencing returns a boolean if a field has been set.

### SetAudioConferencingNil

`func (o *OnlineMeeting) SetAudioConferencingNil(b bool)`

 SetAudioConferencingNil sets the value for AudioConferencing to be an explicit nil

### UnsetAudioConferencing
`func (o *OnlineMeeting) UnsetAudioConferencing()`

UnsetAudioConferencing ensures that no value is present for AudioConferencing, not even an explicit nil
### GetBroadcastSettings

`func (o *OnlineMeeting) GetBroadcastSettings() AnyOfmicrosoftGraphBroadcastMeetingSettings`

GetBroadcastSettings returns the BroadcastSettings field if non-nil, zero value otherwise.

### GetBroadcastSettingsOk

`func (o *OnlineMeeting) GetBroadcastSettingsOk() (*AnyOfmicrosoftGraphBroadcastMeetingSettings, bool)`

GetBroadcastSettingsOk returns a tuple with the BroadcastSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBroadcastSettings

`func (o *OnlineMeeting) SetBroadcastSettings(v AnyOfmicrosoftGraphBroadcastMeetingSettings)`

SetBroadcastSettings sets BroadcastSettings field to given value.

### HasBroadcastSettings

`func (o *OnlineMeeting) HasBroadcastSettings() bool`

HasBroadcastSettings returns a boolean if a field has been set.

### SetBroadcastSettingsNil

`func (o *OnlineMeeting) SetBroadcastSettingsNil(b bool)`

 SetBroadcastSettingsNil sets the value for BroadcastSettings to be an explicit nil

### UnsetBroadcastSettings
`func (o *OnlineMeeting) UnsetBroadcastSettings()`

UnsetBroadcastSettings ensures that no value is present for BroadcastSettings, not even an explicit nil
### GetChatInfo

`func (o *OnlineMeeting) GetChatInfo() AnyOfmicrosoftGraphChatInfo`

GetChatInfo returns the ChatInfo field if non-nil, zero value otherwise.

### GetChatInfoOk

`func (o *OnlineMeeting) GetChatInfoOk() (*AnyOfmicrosoftGraphChatInfo, bool)`

GetChatInfoOk returns a tuple with the ChatInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatInfo

`func (o *OnlineMeeting) SetChatInfo(v AnyOfmicrosoftGraphChatInfo)`

SetChatInfo sets ChatInfo field to given value.

### HasChatInfo

`func (o *OnlineMeeting) HasChatInfo() bool`

HasChatInfo returns a boolean if a field has been set.

### SetChatInfoNil

`func (o *OnlineMeeting) SetChatInfoNil(b bool)`

 SetChatInfoNil sets the value for ChatInfo to be an explicit nil

### UnsetChatInfo
`func (o *OnlineMeeting) UnsetChatInfo()`

UnsetChatInfo ensures that no value is present for ChatInfo, not even an explicit nil
### GetCreationDateTime

`func (o *OnlineMeeting) GetCreationDateTime() time.Time`

GetCreationDateTime returns the CreationDateTime field if non-nil, zero value otherwise.

### GetCreationDateTimeOk

`func (o *OnlineMeeting) GetCreationDateTimeOk() (*time.Time, bool)`

GetCreationDateTimeOk returns a tuple with the CreationDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDateTime

`func (o *OnlineMeeting) SetCreationDateTime(v time.Time)`

SetCreationDateTime sets CreationDateTime field to given value.

### HasCreationDateTime

`func (o *OnlineMeeting) HasCreationDateTime() bool`

HasCreationDateTime returns a boolean if a field has been set.

### SetCreationDateTimeNil

`func (o *OnlineMeeting) SetCreationDateTimeNil(b bool)`

 SetCreationDateTimeNil sets the value for CreationDateTime to be an explicit nil

### UnsetCreationDateTime
`func (o *OnlineMeeting) UnsetCreationDateTime()`

UnsetCreationDateTime ensures that no value is present for CreationDateTime, not even an explicit nil
### GetEndDateTime

`func (o *OnlineMeeting) GetEndDateTime() time.Time`

GetEndDateTime returns the EndDateTime field if non-nil, zero value otherwise.

### GetEndDateTimeOk

`func (o *OnlineMeeting) GetEndDateTimeOk() (*time.Time, bool)`

GetEndDateTimeOk returns a tuple with the EndDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDateTime

`func (o *OnlineMeeting) SetEndDateTime(v time.Time)`

SetEndDateTime sets EndDateTime field to given value.

### HasEndDateTime

`func (o *OnlineMeeting) HasEndDateTime() bool`

HasEndDateTime returns a boolean if a field has been set.

### SetEndDateTimeNil

`func (o *OnlineMeeting) SetEndDateTimeNil(b bool)`

 SetEndDateTimeNil sets the value for EndDateTime to be an explicit nil

### UnsetEndDateTime
`func (o *OnlineMeeting) UnsetEndDateTime()`

UnsetEndDateTime ensures that no value is present for EndDateTime, not even an explicit nil
### GetExternalId

`func (o *OnlineMeeting) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *OnlineMeeting) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *OnlineMeeting) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *OnlineMeeting) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *OnlineMeeting) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *OnlineMeeting) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetIsBroadcast

`func (o *OnlineMeeting) GetIsBroadcast() bool`

GetIsBroadcast returns the IsBroadcast field if non-nil, zero value otherwise.

### GetIsBroadcastOk

`func (o *OnlineMeeting) GetIsBroadcastOk() (*bool, bool)`

GetIsBroadcastOk returns a tuple with the IsBroadcast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBroadcast

`func (o *OnlineMeeting) SetIsBroadcast(v bool)`

SetIsBroadcast sets IsBroadcast field to given value.

### HasIsBroadcast

`func (o *OnlineMeeting) HasIsBroadcast() bool`

HasIsBroadcast returns a boolean if a field has been set.

### SetIsBroadcastNil

`func (o *OnlineMeeting) SetIsBroadcastNil(b bool)`

 SetIsBroadcastNil sets the value for IsBroadcast to be an explicit nil

### UnsetIsBroadcast
`func (o *OnlineMeeting) UnsetIsBroadcast()`

UnsetIsBroadcast ensures that no value is present for IsBroadcast, not even an explicit nil
### GetIsEntryExitAnnounced

`func (o *OnlineMeeting) GetIsEntryExitAnnounced() bool`

GetIsEntryExitAnnounced returns the IsEntryExitAnnounced field if non-nil, zero value otherwise.

### GetIsEntryExitAnnouncedOk

`func (o *OnlineMeeting) GetIsEntryExitAnnouncedOk() (*bool, bool)`

GetIsEntryExitAnnouncedOk returns a tuple with the IsEntryExitAnnounced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEntryExitAnnounced

`func (o *OnlineMeeting) SetIsEntryExitAnnounced(v bool)`

SetIsEntryExitAnnounced sets IsEntryExitAnnounced field to given value.

### HasIsEntryExitAnnounced

`func (o *OnlineMeeting) HasIsEntryExitAnnounced() bool`

HasIsEntryExitAnnounced returns a boolean if a field has been set.

### SetIsEntryExitAnnouncedNil

`func (o *OnlineMeeting) SetIsEntryExitAnnouncedNil(b bool)`

 SetIsEntryExitAnnouncedNil sets the value for IsEntryExitAnnounced to be an explicit nil

### UnsetIsEntryExitAnnounced
`func (o *OnlineMeeting) UnsetIsEntryExitAnnounced()`

UnsetIsEntryExitAnnounced ensures that no value is present for IsEntryExitAnnounced, not even an explicit nil
### GetJoinInformation

`func (o *OnlineMeeting) GetJoinInformation() AnyOfmicrosoftGraphItemBody`

GetJoinInformation returns the JoinInformation field if non-nil, zero value otherwise.

### GetJoinInformationOk

`func (o *OnlineMeeting) GetJoinInformationOk() (*AnyOfmicrosoftGraphItemBody, bool)`

GetJoinInformationOk returns a tuple with the JoinInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinInformation

`func (o *OnlineMeeting) SetJoinInformation(v AnyOfmicrosoftGraphItemBody)`

SetJoinInformation sets JoinInformation field to given value.

### HasJoinInformation

`func (o *OnlineMeeting) HasJoinInformation() bool`

HasJoinInformation returns a boolean if a field has been set.

### SetJoinInformationNil

`func (o *OnlineMeeting) SetJoinInformationNil(b bool)`

 SetJoinInformationNil sets the value for JoinInformation to be an explicit nil

### UnsetJoinInformation
`func (o *OnlineMeeting) UnsetJoinInformation()`

UnsetJoinInformation ensures that no value is present for JoinInformation, not even an explicit nil
### GetJoinWebUrl

`func (o *OnlineMeeting) GetJoinWebUrl() string`

GetJoinWebUrl returns the JoinWebUrl field if non-nil, zero value otherwise.

### GetJoinWebUrlOk

`func (o *OnlineMeeting) GetJoinWebUrlOk() (*string, bool)`

GetJoinWebUrlOk returns a tuple with the JoinWebUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinWebUrl

`func (o *OnlineMeeting) SetJoinWebUrl(v string)`

SetJoinWebUrl sets JoinWebUrl field to given value.

### HasJoinWebUrl

`func (o *OnlineMeeting) HasJoinWebUrl() bool`

HasJoinWebUrl returns a boolean if a field has been set.

### SetJoinWebUrlNil

`func (o *OnlineMeeting) SetJoinWebUrlNil(b bool)`

 SetJoinWebUrlNil sets the value for JoinWebUrl to be an explicit nil

### UnsetJoinWebUrl
`func (o *OnlineMeeting) UnsetJoinWebUrl()`

UnsetJoinWebUrl ensures that no value is present for JoinWebUrl, not even an explicit nil
### GetLobbyBypassSettings

`func (o *OnlineMeeting) GetLobbyBypassSettings() AnyOfmicrosoftGraphLobbyBypassSettings`

GetLobbyBypassSettings returns the LobbyBypassSettings field if non-nil, zero value otherwise.

### GetLobbyBypassSettingsOk

`func (o *OnlineMeeting) GetLobbyBypassSettingsOk() (*AnyOfmicrosoftGraphLobbyBypassSettings, bool)`

GetLobbyBypassSettingsOk returns a tuple with the LobbyBypassSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLobbyBypassSettings

`func (o *OnlineMeeting) SetLobbyBypassSettings(v AnyOfmicrosoftGraphLobbyBypassSettings)`

SetLobbyBypassSettings sets LobbyBypassSettings field to given value.

### HasLobbyBypassSettings

`func (o *OnlineMeeting) HasLobbyBypassSettings() bool`

HasLobbyBypassSettings returns a boolean if a field has been set.

### SetLobbyBypassSettingsNil

`func (o *OnlineMeeting) SetLobbyBypassSettingsNil(b bool)`

 SetLobbyBypassSettingsNil sets the value for LobbyBypassSettings to be an explicit nil

### UnsetLobbyBypassSettings
`func (o *OnlineMeeting) UnsetLobbyBypassSettings()`

UnsetLobbyBypassSettings ensures that no value is present for LobbyBypassSettings, not even an explicit nil
### GetParticipants

`func (o *OnlineMeeting) GetParticipants() AnyOfmicrosoftGraphMeetingParticipants`

GetParticipants returns the Participants field if non-nil, zero value otherwise.

### GetParticipantsOk

`func (o *OnlineMeeting) GetParticipantsOk() (*AnyOfmicrosoftGraphMeetingParticipants, bool)`

GetParticipantsOk returns a tuple with the Participants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipants

`func (o *OnlineMeeting) SetParticipants(v AnyOfmicrosoftGraphMeetingParticipants)`

SetParticipants sets Participants field to given value.

### HasParticipants

`func (o *OnlineMeeting) HasParticipants() bool`

HasParticipants returns a boolean if a field has been set.

### SetParticipantsNil

`func (o *OnlineMeeting) SetParticipantsNil(b bool)`

 SetParticipantsNil sets the value for Participants to be an explicit nil

### UnsetParticipants
`func (o *OnlineMeeting) UnsetParticipants()`

UnsetParticipants ensures that no value is present for Participants, not even an explicit nil
### GetRecordAutomatically

`func (o *OnlineMeeting) GetRecordAutomatically() bool`

GetRecordAutomatically returns the RecordAutomatically field if non-nil, zero value otherwise.

### GetRecordAutomaticallyOk

`func (o *OnlineMeeting) GetRecordAutomaticallyOk() (*bool, bool)`

GetRecordAutomaticallyOk returns a tuple with the RecordAutomatically field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordAutomatically

`func (o *OnlineMeeting) SetRecordAutomatically(v bool)`

SetRecordAutomatically sets RecordAutomatically field to given value.

### HasRecordAutomatically

`func (o *OnlineMeeting) HasRecordAutomatically() bool`

HasRecordAutomatically returns a boolean if a field has been set.

### SetRecordAutomaticallyNil

`func (o *OnlineMeeting) SetRecordAutomaticallyNil(b bool)`

 SetRecordAutomaticallyNil sets the value for RecordAutomatically to be an explicit nil

### UnsetRecordAutomatically
`func (o *OnlineMeeting) UnsetRecordAutomatically()`

UnsetRecordAutomatically ensures that no value is present for RecordAutomatically, not even an explicit nil
### GetStartDateTime

`func (o *OnlineMeeting) GetStartDateTime() time.Time`

GetStartDateTime returns the StartDateTime field if non-nil, zero value otherwise.

### GetStartDateTimeOk

`func (o *OnlineMeeting) GetStartDateTimeOk() (*time.Time, bool)`

GetStartDateTimeOk returns a tuple with the StartDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDateTime

`func (o *OnlineMeeting) SetStartDateTime(v time.Time)`

SetStartDateTime sets StartDateTime field to given value.

### HasStartDateTime

`func (o *OnlineMeeting) HasStartDateTime() bool`

HasStartDateTime returns a boolean if a field has been set.

### SetStartDateTimeNil

`func (o *OnlineMeeting) SetStartDateTimeNil(b bool)`

 SetStartDateTimeNil sets the value for StartDateTime to be an explicit nil

### UnsetStartDateTime
`func (o *OnlineMeeting) UnsetStartDateTime()`

UnsetStartDateTime ensures that no value is present for StartDateTime, not even an explicit nil
### GetSubject

`func (o *OnlineMeeting) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *OnlineMeeting) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *OnlineMeeting) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *OnlineMeeting) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### SetSubjectNil

`func (o *OnlineMeeting) SetSubjectNil(b bool)`

 SetSubjectNil sets the value for Subject to be an explicit nil

### UnsetSubject
`func (o *OnlineMeeting) UnsetSubject()`

UnsetSubject ensures that no value is present for Subject, not even an explicit nil
### GetVideoTeleconferenceId

`func (o *OnlineMeeting) GetVideoTeleconferenceId() string`

GetVideoTeleconferenceId returns the VideoTeleconferenceId field if non-nil, zero value otherwise.

### GetVideoTeleconferenceIdOk

`func (o *OnlineMeeting) GetVideoTeleconferenceIdOk() (*string, bool)`

GetVideoTeleconferenceIdOk returns a tuple with the VideoTeleconferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoTeleconferenceId

`func (o *OnlineMeeting) SetVideoTeleconferenceId(v string)`

SetVideoTeleconferenceId sets VideoTeleconferenceId field to given value.

### HasVideoTeleconferenceId

`func (o *OnlineMeeting) HasVideoTeleconferenceId() bool`

HasVideoTeleconferenceId returns a boolean if a field has been set.

### SetVideoTeleconferenceIdNil

`func (o *OnlineMeeting) SetVideoTeleconferenceIdNil(b bool)`

 SetVideoTeleconferenceIdNil sets the value for VideoTeleconferenceId to be an explicit nil

### UnsetVideoTeleconferenceId
`func (o *OnlineMeeting) UnsetVideoTeleconferenceId()`

UnsetVideoTeleconferenceId ensures that no value is present for VideoTeleconferenceId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


