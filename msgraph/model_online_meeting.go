/*
Partial Graph API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// OnlineMeeting struct for OnlineMeeting
type OnlineMeeting struct {
	// Indicates whether attendees can turn on their camera.
	AllowAttendeeToEnableCamera NullableBool `json:"allowAttendeeToEnableCamera,omitempty"`
	// Indicates whether attendees can turn on their microphone.
	AllowAttendeeToEnableMic NullableBool `json:"allowAttendeeToEnableMic,omitempty"`
	// Specifies who can be a presenter in a meeting. Possible values are listed in the following table.
	AllowedPresenters AnyOfmicrosoftGraphOnlineMeetingPresenters `json:"allowedPresenters,omitempty"`
	// Specifies the mode of meeting chat.
	AllowMeetingChat AnyOfmicrosoftGraphMeetingChatMode `json:"allowMeetingChat,omitempty"`
	// Indicates whether Teams reactions are enabled for the meeting.
	AllowTeamworkReactions NullableBool `json:"allowTeamworkReactions,omitempty"`
	// The content stream of the attendee report of a Microsoft Teams live event. Read-only.
	AttendeeReport NullableString `json:"attendeeReport,omitempty"`
	// The phone access (dial-in) information for an online meeting. Read-only.
	AudioConferencing AnyOfmicrosoftGraphAudioConferencing `json:"audioConferencing,omitempty"`
	// Settings related to a live event.
	BroadcastSettings AnyOfmicrosoftGraphBroadcastMeetingSettings `json:"broadcastSettings,omitempty"`
	// The chat information associated with this online meeting.
	ChatInfo AnyOfmicrosoftGraphChatInfo `json:"chatInfo,omitempty"`
	// The meeting creation time in UTC. Read-only.
	CreationDateTime NullableTime `json:"creationDateTime,omitempty"`
	// The meeting end time in UTC.
	EndDateTime NullableTime `json:"endDateTime,omitempty"`
	// The external ID. A custom ID. Optional.
	ExternalId NullableString `json:"externalId,omitempty"`
	// Indicates if this is a Teams live event.
	IsBroadcast NullableBool `json:"isBroadcast,omitempty"`
	// Indicates whether to announce when callers join or leave.
	IsEntryExitAnnounced NullableBool `json:"isEntryExitAnnounced,omitempty"`
	// The join information in the language and locale variant specified in the Accept-Language request HTTP header. Read-only.
	JoinInformation AnyOfmicrosoftGraphItemBody `json:"joinInformation,omitempty"`
	// The join URL of the online meeting. Read-only.
	JoinWebUrl NullableString `json:"joinWebUrl,omitempty"`
	// Specifies which participants can bypass the meeting   lobby.
	LobbyBypassSettings AnyOfmicrosoftGraphLobbyBypassSettings `json:"lobbyBypassSettings,omitempty"`
	// The participants associated with the online meeting.  This includes the organizer and the attendees.
	Participants AnyOfmicrosoftGraphMeetingParticipants `json:"participants,omitempty"`
	// Indicates whether to record the meeting automatically.
	RecordAutomatically NullableBool `json:"recordAutomatically,omitempty"`
	// The meeting start time in UTC.
	StartDateTime NullableTime `json:"startDateTime,omitempty"`
	// The subject of the online meeting.
	Subject NullableString `json:"subject,omitempty"`
	// The video teleconferencing ID. Read-only.
	VideoTeleconferenceId NullableString `json:"videoTeleconferenceId,omitempty"`
}

// NewOnlineMeeting instantiates a new OnlineMeeting object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOnlineMeeting() *OnlineMeeting {
	this := OnlineMeeting{}
	return &this
}

// NewOnlineMeetingWithDefaults instantiates a new OnlineMeeting object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOnlineMeetingWithDefaults() *OnlineMeeting {
	this := OnlineMeeting{}
	return &this
}

// GetAllowAttendeeToEnableCamera returns the AllowAttendeeToEnableCamera field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OnlineMeeting) GetAllowAttendeeToEnableCamera() bool {
	if o == nil || o.AllowAttendeeToEnableCamera.Get() == nil {
		var ret bool
		return ret
	}
	return *o.AllowAttendeeToEnableCamera.Get()
}

// GetAllowAttendeeToEnableCameraOk returns a tuple with the AllowAttendeeToEnableCamera field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OnlineMeeting) GetAllowAttendeeToEnableCameraOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AllowAttendeeToEnableCamera.Get(), o.AllowAttendeeToEnableCamera.IsSet()
}

// HasAllowAttendeeToEnableCamera returns a boolean if a field has been set.
func (o *OnlineMeeting) HasAllowAttendeeToEnableCamera() bool {
	if o != nil && o.AllowAttendeeToEnableCamera.IsSet() {
		return true
	}

	return false
}

// SetAllowAttendeeToEnableCamera gets a reference to the given NullableBool and assigns it to the AllowAttendeeToEnableCamera field.
func (o *OnlineMeeting) SetAllowAttendeeToEnableCamera(v bool) {
	o.AllowAttendeeToEnableCamera.Set(&v)
}
// SetAllowAttendeeToEnableCameraNil sets the value for AllowAttendeeToEnableCamera to be an explicit nil
func (o *OnlineMeeting) SetAllowAttendeeToEnableCameraNil() {
	o.AllowAttendeeToEnableCamera.Set(nil)
}

// UnsetAllowAttendeeToEnableCamera ensures that no value is present for AllowAttendeeToEnableCamera, not even an explicit nil
func (o *OnlineMeeting) UnsetAllowAttendeeToEnableCamera() {
	o.AllowAttendeeToEnableCamera.Unset()
}

// GetAllowAttendeeToEnableMic returns the AllowAttendeeToEnableMic field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OnlineMeeting) GetAllowAttendeeToEnableMic() bool {
	if o == nil || o.AllowAttendeeToEnableMic.Get() == nil {
		var ret bool
		return ret
	}
	return *o.AllowAttendeeToEnableMic.Get()
}

// GetAllowAttendeeToEnableMicOk returns a tuple with the AllowAttendeeToEnableMic field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OnlineMeeting) GetAllowAttendeeToEnableMicOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AllowAttendeeToEnableMic.Get(), o.AllowAttendeeToEnableMic.IsSet()
}

// HasAllowAttendeeToEnableMic returns a boolean if a field has been set.
func (o *OnlineMeeting) HasAllowAttendeeToEnableMic() bool {
	if o != nil && o.AllowAttendeeToEnableMic.IsSet() {
		return true
	}

	return false
}

// SetAllowAttendeeToEnableMic gets a reference to the given NullableBool and assigns it to the AllowAttendeeToEnableMic field.
func (o *OnlineMeeting) SetAllowAttendeeToEnableMic(v bool) {
	o.AllowAttendeeToEnableMic.Set(&v)
}
// SetAllowAttendeeToEnableMicNil sets the value for AllowAttendeeToEnableMic to be an explicit nil
func (o *OnlineMeeting) SetAllowAttendeeToEnableMicNil() {
	o.AllowAttendeeToEnableMic.Set(nil)
}

// UnsetAllowAttendeeToEnableMic ensures that no value is present for AllowAttendeeToEnableMic, not even an explicit nil
func (o *OnlineMeeting) UnsetAllowAttendeeToEnableMic() {
	o.AllowAttendeeToEnableMic.Unset()
}

// GetAllowedPresenters returns the AllowedPresenters field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OnlineMeeting) GetAllowedPresenters() AnyOfmicrosoftGraphOnlineMeetingPresenters {
	if o == nil  {
		var ret AnyOfmicrosoftGraphOnlineMeetingPresenters
		return ret
	}
	return o.AllowedPresenters
}

// GetAllowedPresentersOk returns a tuple with the AllowedPresenters field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OnlineMeeting) GetAllowedPresentersOk() (*AnyOfmicrosoftGraphOnlineMeetingPresenters, bool) {
	if o == nil || o.AllowedPresenters == nil {
		return nil, false
	}
	return &o.AllowedPresenters, true
}

// HasAllowedPresenters returns a boolean if a field has been set.
func (o *OnlineMeeting) HasAllowedPresenters() bool {
	if o != nil && o.AllowedPresenters != nil {
		return true
	}

	return false
}

// SetAllowedPresenters gets a reference to the given AnyOfmicrosoftGraphOnlineMeetingPresenters and assigns it to the AllowedPresenters field.
func (o *OnlineMeeting) SetAllowedPresenters(v AnyOfmicrosoftGraphOnlineMeetingPresenters) {
	o.AllowedPresenters = v
}

// GetAllowMeetingChat returns the AllowMeetingChat field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OnlineMeeting) GetAllowMeetingChat() AnyOfmicrosoftGraphMeetingChatMode {
	if o == nil  {
		var ret AnyOfmicrosoftGraphMeetingChatMode
		return ret
	}
	return o.AllowMeetingChat
}

// GetAllowMeetingChatOk returns a tuple with the AllowMeetingChat field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OnlineMeeting) GetAllowMeetingChatOk() (*AnyOfmicrosoftGraphMeetingChatMode, bool) {
	if o == nil || o.AllowMeetingChat == nil {
		return nil, false
	}
	return &o.AllowMeetingChat, true
}

// HasAllowMeetingChat returns a boolean if a field has been set.
func (o *OnlineMeeting) HasAllowMeetingChat() bool {
	if o != nil && o.AllowMeetingChat != nil {
		return true
	}

	return false
}

// SetAllowMeetingChat gets a reference to the given AnyOfmicrosoftGraphMeetingChatMode and assigns it to the AllowMeetingChat field.
func (o *OnlineMeeting) SetAllowMeetingChat(v AnyOfmicrosoftGraphMeetingChatMode) {
	o.AllowMeetingChat = v
}

// GetAllowTeamworkReactions returns the AllowTeamworkReactions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OnlineMeeting) GetAllowTeamworkReactions() bool {
	if o == nil || o.AllowTeamworkReactions.Get() == nil {
		var ret bool
		return ret
	}
	return *o.AllowTeamworkReactions.Get()
}

// GetAllowTeamworkReactionsOk returns a tuple with the AllowTeamworkReactions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OnlineMeeting) GetAllowTeamworkReactionsOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AllowTeamworkReactions.Get(), o.AllowTeamworkReactions.IsSet()
}

// HasAllowTeamworkReactions returns a boolean if a field has been set.
func (o *OnlineMeeting) HasAllowTeamworkReactions() bool {
	if o != nil && o.AllowTeamworkReactions.IsSet() {
		return true
	}

	return false
}

// SetAllowTeamworkReactions gets a reference to the given NullableBool and assigns it to the AllowTeamworkReactions field.
func (o *OnlineMeeting) SetAllowTeamworkReactions(v bool) {
	o.AllowTeamworkReactions.Set(&v)
}
// SetAllowTeamworkReactionsNil sets the value for AllowTeamworkReactions to be an explicit nil
func (o *OnlineMeeting) SetAllowTeamworkReactionsNil() {
	o.AllowTeamworkReactions.Set(nil)
}

// UnsetAllowTeamworkReactions ensures that no value is present for AllowTeamworkReactions, not even an explicit nil
func (o *OnlineMeeting) UnsetAllowTeamworkReactions() {
	o.AllowTeamworkReactions.Unset()
}

// GetAttendeeReport returns the AttendeeReport field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OnlineMeeting) GetAttendeeReport() string {
	if o == nil || o.AttendeeReport.Get() == nil {
		var ret string
		return ret
	}
	return *o.AttendeeReport.Get()
}

// GetAttendeeReportOk returns a tuple with the AttendeeReport field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OnlineMeeting) GetAttendeeReportOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AttendeeReport.Get(), o.AttendeeReport.IsSet()
}

// HasAttendeeReport returns a boolean if a field has been set.
func (o *OnlineMeeting) HasAttendeeReport() bool {
	if o != nil && o.AttendeeReport.IsSet() {
		return true
	}

	return false
}

// SetAttendeeReport gets a reference to the given NullableString and assigns it to the AttendeeReport field.
func (o *OnlineMeeting) SetAttendeeReport(v string) {
	o.AttendeeReport.Set(&v)
}
// SetAttendeeReportNil sets the value for AttendeeReport to be an explicit nil
func (o *OnlineMeeting) SetAttendeeReportNil() {
	o.AttendeeReport.Set(nil)
}

// UnsetAttendeeReport ensures that no value is present for AttendeeReport, not even an explicit nil
func (o *OnlineMeeting) UnsetAttendeeReport() {
	o.AttendeeReport.Unset()
}

// GetAudioConferencing returns the AudioConferencing field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OnlineMeeting) GetAudioConferencing() AnyOfmicrosoftGraphAudioConferencing {
	if o == nil  {
		var ret AnyOfmicrosoftGraphAudioConferencing
		return ret
	}
	return o.AudioConferencing
}

// GetAudioConferencingOk returns a tuple with the AudioConferencing field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OnlineMeeting) GetAudioConferencingOk() (*AnyOfmicrosoftGraphAudioConferencing, bool) {
	if o == nil || o.AudioConferencing == nil {
		return nil, false
	}
	return &o.AudioConferencing, true
}

// HasAudioConferencing returns a boolean if a field has been set.
func (o *OnlineMeeting) HasAudioConferencing() bool {
	if o != nil && o.AudioConferencing != nil {
		return true
	}

	return false
}

// SetAudioConferencing gets a reference to the given AnyOfmicrosoftGraphAudioConferencing and assigns it to the AudioConferencing field.
func (o *OnlineMeeting) SetAudioConferencing(v AnyOfmicrosoftGraphAudioConferencing) {
	o.AudioConferencing = v
}

// GetBroadcastSettings returns the BroadcastSettings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OnlineMeeting) GetBroadcastSettings() AnyOfmicrosoftGraphBroadcastMeetingSettings {
	if o == nil  {
		var ret AnyOfmicrosoftGraphBroadcastMeetingSettings
		return ret
	}
	return o.BroadcastSettings
}

// GetBroadcastSettingsOk returns a tuple with the BroadcastSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OnlineMeeting) GetBroadcastSettingsOk() (*AnyOfmicrosoftGraphBroadcastMeetingSettings, bool) {
	if o == nil || o.BroadcastSettings == nil {
		return nil, false
	}
	return &o.BroadcastSettings, true
}

// HasBroadcastSettings returns a boolean if a field has been set.
func (o *OnlineMeeting) HasBroadcastSettings() bool {
	if o != nil && o.BroadcastSettings != nil {
		return true
	}

	return false
}

// SetBroadcastSettings gets a reference to the given AnyOfmicrosoftGraphBroadcastMeetingSettings and assigns it to the BroadcastSettings field.
func (o *OnlineMeeting) SetBroadcastSettings(v AnyOfmicrosoftGraphBroadcastMeetingSettings) {
	o.BroadcastSettings = v
}

// GetChatInfo returns the ChatInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OnlineMeeting) GetChatInfo() AnyOfmicrosoftGraphChatInfo {
	if o == nil  {
		var ret AnyOfmicrosoftGraphChatInfo
		return ret
	}
	return o.ChatInfo
}

// GetChatInfoOk returns a tuple with the ChatInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OnlineMeeting) GetChatInfoOk() (*AnyOfmicrosoftGraphChatInfo, bool) {
	if o == nil || o.ChatInfo == nil {
		return nil, false
	}
	return &o.ChatInfo, true
}

// HasChatInfo returns a boolean if a field has been set.
func (o *OnlineMeeting) HasChatInfo() bool {
	if o != nil && o.ChatInfo != nil {
		return true
	}

	return false
}

// SetChatInfo gets a reference to the given AnyOfmicrosoftGraphChatInfo and assigns it to the ChatInfo field.
func (o *OnlineMeeting) SetChatInfo(v AnyOfmicrosoftGraphChatInfo) {
	o.ChatInfo = v
}

// GetCreationDateTime returns the CreationDateTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OnlineMeeting) GetCreationDateTime() time.Time {
	if o == nil || o.CreationDateTime.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.CreationDateTime.Get()
}

// GetCreationDateTimeOk returns a tuple with the CreationDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OnlineMeeting) GetCreationDateTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CreationDateTime.Get(), o.CreationDateTime.IsSet()
}

// HasCreationDateTime returns a boolean if a field has been set.
func (o *OnlineMeeting) HasCreationDateTime() bool {
	if o != nil && o.CreationDateTime.IsSet() {
		return true
	}

	return false
}

// SetCreationDateTime gets a reference to the given NullableTime and assigns it to the CreationDateTime field.
func (o *OnlineMeeting) SetCreationDateTime(v time.Time) {
	o.CreationDateTime.Set(&v)
}
// SetCreationDateTimeNil sets the value for CreationDateTime to be an explicit nil
func (o *OnlineMeeting) SetCreationDateTimeNil() {
	o.CreationDateTime.Set(nil)
}

// UnsetCreationDateTime ensures that no value is present for CreationDateTime, not even an explicit nil
func (o *OnlineMeeting) UnsetCreationDateTime() {
	o.CreationDateTime.Unset()
}

// GetEndDateTime returns the EndDateTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OnlineMeeting) GetEndDateTime() time.Time {
	if o == nil || o.EndDateTime.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.EndDateTime.Get()
}

// GetEndDateTimeOk returns a tuple with the EndDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OnlineMeeting) GetEndDateTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.EndDateTime.Get(), o.EndDateTime.IsSet()
}

// HasEndDateTime returns a boolean if a field has been set.
func (o *OnlineMeeting) HasEndDateTime() bool {
	if o != nil && o.EndDateTime.IsSet() {
		return true
	}

	return false
}

// SetEndDateTime gets a reference to the given NullableTime and assigns it to the EndDateTime field.
func (o *OnlineMeeting) SetEndDateTime(v time.Time) {
	o.EndDateTime.Set(&v)
}
// SetEndDateTimeNil sets the value for EndDateTime to be an explicit nil
func (o *OnlineMeeting) SetEndDateTimeNil() {
	o.EndDateTime.Set(nil)
}

// UnsetEndDateTime ensures that no value is present for EndDateTime, not even an explicit nil
func (o *OnlineMeeting) UnsetEndDateTime() {
	o.EndDateTime.Unset()
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OnlineMeeting) GetExternalId() string {
	if o == nil || o.ExternalId.Get() == nil {
		var ret string
		return ret
	}
	return *o.ExternalId.Get()
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OnlineMeeting) GetExternalIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ExternalId.Get(), o.ExternalId.IsSet()
}

// HasExternalId returns a boolean if a field has been set.
func (o *OnlineMeeting) HasExternalId() bool {
	if o != nil && o.ExternalId.IsSet() {
		return true
	}

	return false
}

// SetExternalId gets a reference to the given NullableString and assigns it to the ExternalId field.
func (o *OnlineMeeting) SetExternalId(v string) {
	o.ExternalId.Set(&v)
}
// SetExternalIdNil sets the value for ExternalId to be an explicit nil
func (o *OnlineMeeting) SetExternalIdNil() {
	o.ExternalId.Set(nil)
}

// UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
func (o *OnlineMeeting) UnsetExternalId() {
	o.ExternalId.Unset()
}

// GetIsBroadcast returns the IsBroadcast field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OnlineMeeting) GetIsBroadcast() bool {
	if o == nil || o.IsBroadcast.Get() == nil {
		var ret bool
		return ret
	}
	return *o.IsBroadcast.Get()
}

// GetIsBroadcastOk returns a tuple with the IsBroadcast field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OnlineMeeting) GetIsBroadcastOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IsBroadcast.Get(), o.IsBroadcast.IsSet()
}

// HasIsBroadcast returns a boolean if a field has been set.
func (o *OnlineMeeting) HasIsBroadcast() bool {
	if o != nil && o.IsBroadcast.IsSet() {
		return true
	}

	return false
}

// SetIsBroadcast gets a reference to the given NullableBool and assigns it to the IsBroadcast field.
func (o *OnlineMeeting) SetIsBroadcast(v bool) {
	o.IsBroadcast.Set(&v)
}
// SetIsBroadcastNil sets the value for IsBroadcast to be an explicit nil
func (o *OnlineMeeting) SetIsBroadcastNil() {
	o.IsBroadcast.Set(nil)
}

// UnsetIsBroadcast ensures that no value is present for IsBroadcast, not even an explicit nil
func (o *OnlineMeeting) UnsetIsBroadcast() {
	o.IsBroadcast.Unset()
}

// GetIsEntryExitAnnounced returns the IsEntryExitAnnounced field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OnlineMeeting) GetIsEntryExitAnnounced() bool {
	if o == nil || o.IsEntryExitAnnounced.Get() == nil {
		var ret bool
		return ret
	}
	return *o.IsEntryExitAnnounced.Get()
}

// GetIsEntryExitAnnouncedOk returns a tuple with the IsEntryExitAnnounced field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OnlineMeeting) GetIsEntryExitAnnouncedOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IsEntryExitAnnounced.Get(), o.IsEntryExitAnnounced.IsSet()
}

// HasIsEntryExitAnnounced returns a boolean if a field has been set.
func (o *OnlineMeeting) HasIsEntryExitAnnounced() bool {
	if o != nil && o.IsEntryExitAnnounced.IsSet() {
		return true
	}

	return false
}

// SetIsEntryExitAnnounced gets a reference to the given NullableBool and assigns it to the IsEntryExitAnnounced field.
func (o *OnlineMeeting) SetIsEntryExitAnnounced(v bool) {
	o.IsEntryExitAnnounced.Set(&v)
}
// SetIsEntryExitAnnouncedNil sets the value for IsEntryExitAnnounced to be an explicit nil
func (o *OnlineMeeting) SetIsEntryExitAnnouncedNil() {
	o.IsEntryExitAnnounced.Set(nil)
}

// UnsetIsEntryExitAnnounced ensures that no value is present for IsEntryExitAnnounced, not even an explicit nil
func (o *OnlineMeeting) UnsetIsEntryExitAnnounced() {
	o.IsEntryExitAnnounced.Unset()
}

// GetJoinInformation returns the JoinInformation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OnlineMeeting) GetJoinInformation() AnyOfmicrosoftGraphItemBody {
	if o == nil  {
		var ret AnyOfmicrosoftGraphItemBody
		return ret
	}
	return o.JoinInformation
}

// GetJoinInformationOk returns a tuple with the JoinInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OnlineMeeting) GetJoinInformationOk() (*AnyOfmicrosoftGraphItemBody, bool) {
	if o == nil || o.JoinInformation == nil {
		return nil, false
	}
	return &o.JoinInformation, true
}

// HasJoinInformation returns a boolean if a field has been set.
func (o *OnlineMeeting) HasJoinInformation() bool {
	if o != nil && o.JoinInformation != nil {
		return true
	}

	return false
}

// SetJoinInformation gets a reference to the given AnyOfmicrosoftGraphItemBody and assigns it to the JoinInformation field.
func (o *OnlineMeeting) SetJoinInformation(v AnyOfmicrosoftGraphItemBody) {
	o.JoinInformation = v
}

// GetJoinWebUrl returns the JoinWebUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OnlineMeeting) GetJoinWebUrl() string {
	if o == nil || o.JoinWebUrl.Get() == nil {
		var ret string
		return ret
	}
	return *o.JoinWebUrl.Get()
}

// GetJoinWebUrlOk returns a tuple with the JoinWebUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OnlineMeeting) GetJoinWebUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.JoinWebUrl.Get(), o.JoinWebUrl.IsSet()
}

// HasJoinWebUrl returns a boolean if a field has been set.
func (o *OnlineMeeting) HasJoinWebUrl() bool {
	if o != nil && o.JoinWebUrl.IsSet() {
		return true
	}

	return false
}

// SetJoinWebUrl gets a reference to the given NullableString and assigns it to the JoinWebUrl field.
func (o *OnlineMeeting) SetJoinWebUrl(v string) {
	o.JoinWebUrl.Set(&v)
}
// SetJoinWebUrlNil sets the value for JoinWebUrl to be an explicit nil
func (o *OnlineMeeting) SetJoinWebUrlNil() {
	o.JoinWebUrl.Set(nil)
}

// UnsetJoinWebUrl ensures that no value is present for JoinWebUrl, not even an explicit nil
func (o *OnlineMeeting) UnsetJoinWebUrl() {
	o.JoinWebUrl.Unset()
}

// GetLobbyBypassSettings returns the LobbyBypassSettings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OnlineMeeting) GetLobbyBypassSettings() AnyOfmicrosoftGraphLobbyBypassSettings {
	if o == nil  {
		var ret AnyOfmicrosoftGraphLobbyBypassSettings
		return ret
	}
	return o.LobbyBypassSettings
}

// GetLobbyBypassSettingsOk returns a tuple with the LobbyBypassSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OnlineMeeting) GetLobbyBypassSettingsOk() (*AnyOfmicrosoftGraphLobbyBypassSettings, bool) {
	if o == nil || o.LobbyBypassSettings == nil {
		return nil, false
	}
	return &o.LobbyBypassSettings, true
}

// HasLobbyBypassSettings returns a boolean if a field has been set.
func (o *OnlineMeeting) HasLobbyBypassSettings() bool {
	if o != nil && o.LobbyBypassSettings != nil {
		return true
	}

	return false
}

// SetLobbyBypassSettings gets a reference to the given AnyOfmicrosoftGraphLobbyBypassSettings and assigns it to the LobbyBypassSettings field.
func (o *OnlineMeeting) SetLobbyBypassSettings(v AnyOfmicrosoftGraphLobbyBypassSettings) {
	o.LobbyBypassSettings = v
}

// GetParticipants returns the Participants field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OnlineMeeting) GetParticipants() AnyOfmicrosoftGraphMeetingParticipants {
	if o == nil  {
		var ret AnyOfmicrosoftGraphMeetingParticipants
		return ret
	}
	return o.Participants
}

// GetParticipantsOk returns a tuple with the Participants field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OnlineMeeting) GetParticipantsOk() (*AnyOfmicrosoftGraphMeetingParticipants, bool) {
	if o == nil || o.Participants == nil {
		return nil, false
	}
	return &o.Participants, true
}

// HasParticipants returns a boolean if a field has been set.
func (o *OnlineMeeting) HasParticipants() bool {
	if o != nil && o.Participants != nil {
		return true
	}

	return false
}

// SetParticipants gets a reference to the given AnyOfmicrosoftGraphMeetingParticipants and assigns it to the Participants field.
func (o *OnlineMeeting) SetParticipants(v AnyOfmicrosoftGraphMeetingParticipants) {
	o.Participants = v
}

// GetRecordAutomatically returns the RecordAutomatically field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OnlineMeeting) GetRecordAutomatically() bool {
	if o == nil || o.RecordAutomatically.Get() == nil {
		var ret bool
		return ret
	}
	return *o.RecordAutomatically.Get()
}

// GetRecordAutomaticallyOk returns a tuple with the RecordAutomatically field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OnlineMeeting) GetRecordAutomaticallyOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RecordAutomatically.Get(), o.RecordAutomatically.IsSet()
}

// HasRecordAutomatically returns a boolean if a field has been set.
func (o *OnlineMeeting) HasRecordAutomatically() bool {
	if o != nil && o.RecordAutomatically.IsSet() {
		return true
	}

	return false
}

// SetRecordAutomatically gets a reference to the given NullableBool and assigns it to the RecordAutomatically field.
func (o *OnlineMeeting) SetRecordAutomatically(v bool) {
	o.RecordAutomatically.Set(&v)
}
// SetRecordAutomaticallyNil sets the value for RecordAutomatically to be an explicit nil
func (o *OnlineMeeting) SetRecordAutomaticallyNil() {
	o.RecordAutomatically.Set(nil)
}

// UnsetRecordAutomatically ensures that no value is present for RecordAutomatically, not even an explicit nil
func (o *OnlineMeeting) UnsetRecordAutomatically() {
	o.RecordAutomatically.Unset()
}

// GetStartDateTime returns the StartDateTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OnlineMeeting) GetStartDateTime() time.Time {
	if o == nil || o.StartDateTime.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.StartDateTime.Get()
}

// GetStartDateTimeOk returns a tuple with the StartDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OnlineMeeting) GetStartDateTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.StartDateTime.Get(), o.StartDateTime.IsSet()
}

// HasStartDateTime returns a boolean if a field has been set.
func (o *OnlineMeeting) HasStartDateTime() bool {
	if o != nil && o.StartDateTime.IsSet() {
		return true
	}

	return false
}

// SetStartDateTime gets a reference to the given NullableTime and assigns it to the StartDateTime field.
func (o *OnlineMeeting) SetStartDateTime(v time.Time) {
	o.StartDateTime.Set(&v)
}
// SetStartDateTimeNil sets the value for StartDateTime to be an explicit nil
func (o *OnlineMeeting) SetStartDateTimeNil() {
	o.StartDateTime.Set(nil)
}

// UnsetStartDateTime ensures that no value is present for StartDateTime, not even an explicit nil
func (o *OnlineMeeting) UnsetStartDateTime() {
	o.StartDateTime.Unset()
}

// GetSubject returns the Subject field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OnlineMeeting) GetSubject() string {
	if o == nil || o.Subject.Get() == nil {
		var ret string
		return ret
	}
	return *o.Subject.Get()
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OnlineMeeting) GetSubjectOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Subject.Get(), o.Subject.IsSet()
}

// HasSubject returns a boolean if a field has been set.
func (o *OnlineMeeting) HasSubject() bool {
	if o != nil && o.Subject.IsSet() {
		return true
	}

	return false
}

// SetSubject gets a reference to the given NullableString and assigns it to the Subject field.
func (o *OnlineMeeting) SetSubject(v string) {
	o.Subject.Set(&v)
}
// SetSubjectNil sets the value for Subject to be an explicit nil
func (o *OnlineMeeting) SetSubjectNil() {
	o.Subject.Set(nil)
}

// UnsetSubject ensures that no value is present for Subject, not even an explicit nil
func (o *OnlineMeeting) UnsetSubject() {
	o.Subject.Unset()
}

// GetVideoTeleconferenceId returns the VideoTeleconferenceId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OnlineMeeting) GetVideoTeleconferenceId() string {
	if o == nil || o.VideoTeleconferenceId.Get() == nil {
		var ret string
		return ret
	}
	return *o.VideoTeleconferenceId.Get()
}

// GetVideoTeleconferenceIdOk returns a tuple with the VideoTeleconferenceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OnlineMeeting) GetVideoTeleconferenceIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.VideoTeleconferenceId.Get(), o.VideoTeleconferenceId.IsSet()
}

// HasVideoTeleconferenceId returns a boolean if a field has been set.
func (o *OnlineMeeting) HasVideoTeleconferenceId() bool {
	if o != nil && o.VideoTeleconferenceId.IsSet() {
		return true
	}

	return false
}

// SetVideoTeleconferenceId gets a reference to the given NullableString and assigns it to the VideoTeleconferenceId field.
func (o *OnlineMeeting) SetVideoTeleconferenceId(v string) {
	o.VideoTeleconferenceId.Set(&v)
}
// SetVideoTeleconferenceIdNil sets the value for VideoTeleconferenceId to be an explicit nil
func (o *OnlineMeeting) SetVideoTeleconferenceIdNil() {
	o.VideoTeleconferenceId.Set(nil)
}

// UnsetVideoTeleconferenceId ensures that no value is present for VideoTeleconferenceId, not even an explicit nil
func (o *OnlineMeeting) UnsetVideoTeleconferenceId() {
	o.VideoTeleconferenceId.Unset()
}

func (o OnlineMeeting) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AllowAttendeeToEnableCamera.IsSet() {
		toSerialize["allowAttendeeToEnableCamera"] = o.AllowAttendeeToEnableCamera.Get()
	}
	if o.AllowAttendeeToEnableMic.IsSet() {
		toSerialize["allowAttendeeToEnableMic"] = o.AllowAttendeeToEnableMic.Get()
	}
	if o.AllowedPresenters != nil {
		toSerialize["allowedPresenters"] = o.AllowedPresenters
	}
	if o.AllowMeetingChat != nil {
		toSerialize["allowMeetingChat"] = o.AllowMeetingChat
	}
	if o.AllowTeamworkReactions.IsSet() {
		toSerialize["allowTeamworkReactions"] = o.AllowTeamworkReactions.Get()
	}
	if o.AttendeeReport.IsSet() {
		toSerialize["attendeeReport"] = o.AttendeeReport.Get()
	}
	if o.AudioConferencing != nil {
		toSerialize["audioConferencing"] = o.AudioConferencing
	}
	if o.BroadcastSettings != nil {
		toSerialize["broadcastSettings"] = o.BroadcastSettings
	}
	if o.ChatInfo != nil {
		toSerialize["chatInfo"] = o.ChatInfo
	}
	if o.CreationDateTime.IsSet() {
		toSerialize["creationDateTime"] = o.CreationDateTime.Get()
	}
	if o.EndDateTime.IsSet() {
		toSerialize["endDateTime"] = o.EndDateTime.Get()
	}
	if o.ExternalId.IsSet() {
		toSerialize["externalId"] = o.ExternalId.Get()
	}
	if o.IsBroadcast.IsSet() {
		toSerialize["isBroadcast"] = o.IsBroadcast.Get()
	}
	if o.IsEntryExitAnnounced.IsSet() {
		toSerialize["isEntryExitAnnounced"] = o.IsEntryExitAnnounced.Get()
	}
	if o.JoinInformation != nil {
		toSerialize["joinInformation"] = o.JoinInformation
	}
	if o.JoinWebUrl.IsSet() {
		toSerialize["joinWebUrl"] = o.JoinWebUrl.Get()
	}
	if o.LobbyBypassSettings != nil {
		toSerialize["lobbyBypassSettings"] = o.LobbyBypassSettings
	}
	if o.Participants != nil {
		toSerialize["participants"] = o.Participants
	}
	if o.RecordAutomatically.IsSet() {
		toSerialize["recordAutomatically"] = o.RecordAutomatically.Get()
	}
	if o.StartDateTime.IsSet() {
		toSerialize["startDateTime"] = o.StartDateTime.Get()
	}
	if o.Subject.IsSet() {
		toSerialize["subject"] = o.Subject.Get()
	}
	if o.VideoTeleconferenceId.IsSet() {
		toSerialize["videoTeleconferenceId"] = o.VideoTeleconferenceId.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableOnlineMeeting struct {
	value *OnlineMeeting
	isSet bool
}

func (v NullableOnlineMeeting) Get() *OnlineMeeting {
	return v.value
}

func (v *NullableOnlineMeeting) Set(val *OnlineMeeting) {
	v.value = val
	v.isSet = true
}

func (v NullableOnlineMeeting) IsSet() bool {
	return v.isSet
}

func (v *NullableOnlineMeeting) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOnlineMeeting(val *OnlineMeeting) *NullableOnlineMeeting {
	return &NullableOnlineMeeting{value: val, isSet: true}
}

func (v NullableOnlineMeeting) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOnlineMeeting) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


