/*
Partial Graph API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package msgraph

import (
	"encoding/json"
)

// MicrosoftGraphCall struct for MicrosoftGraphCall
type MicrosoftGraphCall struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
	// The callback URL on which callbacks will be delivered. Must be https.
	CallbackUri *string `json:"callbackUri,omitempty"`
	// A unique identifier for all the participant calls in a conference or a unique identifier for two participant calls in a P2P call.  This needs to be copied over from Microsoft.Graph.Call.CallChainId.
	CallChainId NullableString `json:"callChainId,omitempty"`
	CallOptions AnyOfobject `json:"callOptions,omitempty"`
	// The routing information on how the call was retargeted. Read-only.
	CallRoutes *[]*AnyOfmicrosoftGraphCallRoute `json:"callRoutes,omitempty"`
	// The chat information. Required information for joining a meeting.
	ChatInfo AnyOfmicrosoftGraphChatInfo `json:"chatInfo,omitempty"`
	// The direction of the call. The possible value are incoming or outgoing. Read-only.
	Direction AnyOfmicrosoftGraphCallDirection `json:"direction,omitempty"`
	// The context associated with an incoming call. Read-only. Server generated.
	IncomingContext AnyOfmicrosoftGraphIncomingContext `json:"incomingContext,omitempty"`
	// The media configuration. Required.
	MediaConfig AnyOfobject `json:"mediaConfig,omitempty"`
	// Read-only. The call media state.
	MediaState AnyOfmicrosoftGraphCallMediaState `json:"mediaState,omitempty"`
	// The meeting information that's required for joining a meeting.
	MeetingInfo AnyOfobject `json:"meetingInfo,omitempty"`
	MyParticipantId NullableString `json:"myParticipantId,omitempty"`
	RequestedModalities *[]*AnyOfmicrosoftGraphModality `json:"requestedModalities,omitempty"`
	ResultInfo AnyOfmicrosoftGraphResultInfo `json:"resultInfo,omitempty"`
	Source AnyOfmicrosoftGraphParticipantInfo `json:"source,omitempty"`
	State AnyOfmicrosoftGraphCallState `json:"state,omitempty"`
	Subject NullableString `json:"subject,omitempty"`
	Targets *[]*AnyOfmicrosoftGraphInvitationParticipantInfo `json:"targets,omitempty"`
	TenantId NullableString `json:"tenantId,omitempty"`
	ToneInfo AnyOfmicrosoftGraphToneInfo `json:"toneInfo,omitempty"`
	// The transcription information for the call. Read-only.
	Transcription AnyOfmicrosoftGraphCallTranscriptionInfo `json:"transcription,omitempty"`
	// Read-only. Nullable.
	AudioRoutingGroups *[]MicrosoftGraphAudioRoutingGroup `json:"audioRoutingGroups,omitempty"`
	// Read-only. Nullable.
	Operations *[]MicrosoftGraphCommsOperation `json:"operations,omitempty"`
	// Read-only. Nullable.
	Participants *[]MicrosoftGraphParticipant `json:"participants,omitempty"`
}

// NewMicrosoftGraphCall instantiates a new MicrosoftGraphCall object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphCall() *MicrosoftGraphCall {
	this := MicrosoftGraphCall{}
	return &this
}

// NewMicrosoftGraphCallWithDefaults instantiates a new MicrosoftGraphCall object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphCallWithDefaults() *MicrosoftGraphCall {
	this := MicrosoftGraphCall{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MicrosoftGraphCall) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphCall) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphCall) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MicrosoftGraphCall) SetId(v string) {
	o.Id = &v
}

// GetCallbackUri returns the CallbackUri field value if set, zero value otherwise.
func (o *MicrosoftGraphCall) GetCallbackUri() string {
	if o == nil || o.CallbackUri == nil {
		var ret string
		return ret
	}
	return *o.CallbackUri
}

// GetCallbackUriOk returns a tuple with the CallbackUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphCall) GetCallbackUriOk() (*string, bool) {
	if o == nil || o.CallbackUri == nil {
		return nil, false
	}
	return o.CallbackUri, true
}

// HasCallbackUri returns a boolean if a field has been set.
func (o *MicrosoftGraphCall) HasCallbackUri() bool {
	if o != nil && o.CallbackUri != nil {
		return true
	}

	return false
}

// SetCallbackUri gets a reference to the given string and assigns it to the CallbackUri field.
func (o *MicrosoftGraphCall) SetCallbackUri(v string) {
	o.CallbackUri = &v
}

// GetCallChainId returns the CallChainId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphCall) GetCallChainId() string {
	if o == nil || o.CallChainId.Get() == nil {
		var ret string
		return ret
	}
	return *o.CallChainId.Get()
}

// GetCallChainIdOk returns a tuple with the CallChainId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphCall) GetCallChainIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CallChainId.Get(), o.CallChainId.IsSet()
}

// HasCallChainId returns a boolean if a field has been set.
func (o *MicrosoftGraphCall) HasCallChainId() bool {
	if o != nil && o.CallChainId.IsSet() {
		return true
	}

	return false
}

// SetCallChainId gets a reference to the given NullableString and assigns it to the CallChainId field.
func (o *MicrosoftGraphCall) SetCallChainId(v string) {
	o.CallChainId.Set(&v)
}
// SetCallChainIdNil sets the value for CallChainId to be an explicit nil
func (o *MicrosoftGraphCall) SetCallChainIdNil() {
	o.CallChainId.Set(nil)
}

// UnsetCallChainId ensures that no value is present for CallChainId, not even an explicit nil
func (o *MicrosoftGraphCall) UnsetCallChainId() {
	o.CallChainId.Unset()
}

// GetCallOptions returns the CallOptions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphCall) GetCallOptions() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.CallOptions
}

// GetCallOptionsOk returns a tuple with the CallOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphCall) GetCallOptionsOk() (*AnyOfobject, bool) {
	if o == nil || o.CallOptions == nil {
		return nil, false
	}
	return &o.CallOptions, true
}

// HasCallOptions returns a boolean if a field has been set.
func (o *MicrosoftGraphCall) HasCallOptions() bool {
	if o != nil && o.CallOptions != nil {
		return true
	}

	return false
}

// SetCallOptions gets a reference to the given AnyOfobject and assigns it to the CallOptions field.
func (o *MicrosoftGraphCall) SetCallOptions(v AnyOfobject) {
	o.CallOptions = v
}

// GetCallRoutes returns the CallRoutes field value if set, zero value otherwise.
func (o *MicrosoftGraphCall) GetCallRoutes() []*AnyOfmicrosoftGraphCallRoute {
	if o == nil || o.CallRoutes == nil {
		var ret []*AnyOfmicrosoftGraphCallRoute
		return ret
	}
	return *o.CallRoutes
}

// GetCallRoutesOk returns a tuple with the CallRoutes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphCall) GetCallRoutesOk() (*[]*AnyOfmicrosoftGraphCallRoute, bool) {
	if o == nil || o.CallRoutes == nil {
		return nil, false
	}
	return o.CallRoutes, true
}

// HasCallRoutes returns a boolean if a field has been set.
func (o *MicrosoftGraphCall) HasCallRoutes() bool {
	if o != nil && o.CallRoutes != nil {
		return true
	}

	return false
}

// SetCallRoutes gets a reference to the given []*AnyOfmicrosoftGraphCallRoute and assigns it to the CallRoutes field.
func (o *MicrosoftGraphCall) SetCallRoutes(v []*AnyOfmicrosoftGraphCallRoute) {
	o.CallRoutes = &v
}

// GetChatInfo returns the ChatInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphCall) GetChatInfo() AnyOfmicrosoftGraphChatInfo {
	if o == nil  {
		var ret AnyOfmicrosoftGraphChatInfo
		return ret
	}
	return o.ChatInfo
}

// GetChatInfoOk returns a tuple with the ChatInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphCall) GetChatInfoOk() (*AnyOfmicrosoftGraphChatInfo, bool) {
	if o == nil || o.ChatInfo == nil {
		return nil, false
	}
	return &o.ChatInfo, true
}

// HasChatInfo returns a boolean if a field has been set.
func (o *MicrosoftGraphCall) HasChatInfo() bool {
	if o != nil && o.ChatInfo != nil {
		return true
	}

	return false
}

// SetChatInfo gets a reference to the given AnyOfmicrosoftGraphChatInfo and assigns it to the ChatInfo field.
func (o *MicrosoftGraphCall) SetChatInfo(v AnyOfmicrosoftGraphChatInfo) {
	o.ChatInfo = v
}

// GetDirection returns the Direction field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphCall) GetDirection() AnyOfmicrosoftGraphCallDirection {
	if o == nil  {
		var ret AnyOfmicrosoftGraphCallDirection
		return ret
	}
	return o.Direction
}

// GetDirectionOk returns a tuple with the Direction field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphCall) GetDirectionOk() (*AnyOfmicrosoftGraphCallDirection, bool) {
	if o == nil || o.Direction == nil {
		return nil, false
	}
	return &o.Direction, true
}

// HasDirection returns a boolean if a field has been set.
func (o *MicrosoftGraphCall) HasDirection() bool {
	if o != nil && o.Direction != nil {
		return true
	}

	return false
}

// SetDirection gets a reference to the given AnyOfmicrosoftGraphCallDirection and assigns it to the Direction field.
func (o *MicrosoftGraphCall) SetDirection(v AnyOfmicrosoftGraphCallDirection) {
	o.Direction = v
}

// GetIncomingContext returns the IncomingContext field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphCall) GetIncomingContext() AnyOfmicrosoftGraphIncomingContext {
	if o == nil  {
		var ret AnyOfmicrosoftGraphIncomingContext
		return ret
	}
	return o.IncomingContext
}

// GetIncomingContextOk returns a tuple with the IncomingContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphCall) GetIncomingContextOk() (*AnyOfmicrosoftGraphIncomingContext, bool) {
	if o == nil || o.IncomingContext == nil {
		return nil, false
	}
	return &o.IncomingContext, true
}

// HasIncomingContext returns a boolean if a field has been set.
func (o *MicrosoftGraphCall) HasIncomingContext() bool {
	if o != nil && o.IncomingContext != nil {
		return true
	}

	return false
}

// SetIncomingContext gets a reference to the given AnyOfmicrosoftGraphIncomingContext and assigns it to the IncomingContext field.
func (o *MicrosoftGraphCall) SetIncomingContext(v AnyOfmicrosoftGraphIncomingContext) {
	o.IncomingContext = v
}

// GetMediaConfig returns the MediaConfig field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphCall) GetMediaConfig() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.MediaConfig
}

// GetMediaConfigOk returns a tuple with the MediaConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphCall) GetMediaConfigOk() (*AnyOfobject, bool) {
	if o == nil || o.MediaConfig == nil {
		return nil, false
	}
	return &o.MediaConfig, true
}

// HasMediaConfig returns a boolean if a field has been set.
func (o *MicrosoftGraphCall) HasMediaConfig() bool {
	if o != nil && o.MediaConfig != nil {
		return true
	}

	return false
}

// SetMediaConfig gets a reference to the given AnyOfobject and assigns it to the MediaConfig field.
func (o *MicrosoftGraphCall) SetMediaConfig(v AnyOfobject) {
	o.MediaConfig = v
}

// GetMediaState returns the MediaState field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphCall) GetMediaState() AnyOfmicrosoftGraphCallMediaState {
	if o == nil  {
		var ret AnyOfmicrosoftGraphCallMediaState
		return ret
	}
	return o.MediaState
}

// GetMediaStateOk returns a tuple with the MediaState field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphCall) GetMediaStateOk() (*AnyOfmicrosoftGraphCallMediaState, bool) {
	if o == nil || o.MediaState == nil {
		return nil, false
	}
	return &o.MediaState, true
}

// HasMediaState returns a boolean if a field has been set.
func (o *MicrosoftGraphCall) HasMediaState() bool {
	if o != nil && o.MediaState != nil {
		return true
	}

	return false
}

// SetMediaState gets a reference to the given AnyOfmicrosoftGraphCallMediaState and assigns it to the MediaState field.
func (o *MicrosoftGraphCall) SetMediaState(v AnyOfmicrosoftGraphCallMediaState) {
	o.MediaState = v
}

// GetMeetingInfo returns the MeetingInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphCall) GetMeetingInfo() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.MeetingInfo
}

// GetMeetingInfoOk returns a tuple with the MeetingInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphCall) GetMeetingInfoOk() (*AnyOfobject, bool) {
	if o == nil || o.MeetingInfo == nil {
		return nil, false
	}
	return &o.MeetingInfo, true
}

// HasMeetingInfo returns a boolean if a field has been set.
func (o *MicrosoftGraphCall) HasMeetingInfo() bool {
	if o != nil && o.MeetingInfo != nil {
		return true
	}

	return false
}

// SetMeetingInfo gets a reference to the given AnyOfobject and assigns it to the MeetingInfo field.
func (o *MicrosoftGraphCall) SetMeetingInfo(v AnyOfobject) {
	o.MeetingInfo = v
}

// GetMyParticipantId returns the MyParticipantId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphCall) GetMyParticipantId() string {
	if o == nil || o.MyParticipantId.Get() == nil {
		var ret string
		return ret
	}
	return *o.MyParticipantId.Get()
}

// GetMyParticipantIdOk returns a tuple with the MyParticipantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphCall) GetMyParticipantIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.MyParticipantId.Get(), o.MyParticipantId.IsSet()
}

// HasMyParticipantId returns a boolean if a field has been set.
func (o *MicrosoftGraphCall) HasMyParticipantId() bool {
	if o != nil && o.MyParticipantId.IsSet() {
		return true
	}

	return false
}

// SetMyParticipantId gets a reference to the given NullableString and assigns it to the MyParticipantId field.
func (o *MicrosoftGraphCall) SetMyParticipantId(v string) {
	o.MyParticipantId.Set(&v)
}
// SetMyParticipantIdNil sets the value for MyParticipantId to be an explicit nil
func (o *MicrosoftGraphCall) SetMyParticipantIdNil() {
	o.MyParticipantId.Set(nil)
}

// UnsetMyParticipantId ensures that no value is present for MyParticipantId, not even an explicit nil
func (o *MicrosoftGraphCall) UnsetMyParticipantId() {
	o.MyParticipantId.Unset()
}

// GetRequestedModalities returns the RequestedModalities field value if set, zero value otherwise.
func (o *MicrosoftGraphCall) GetRequestedModalities() []*AnyOfmicrosoftGraphModality {
	if o == nil || o.RequestedModalities == nil {
		var ret []*AnyOfmicrosoftGraphModality
		return ret
	}
	return *o.RequestedModalities
}

// GetRequestedModalitiesOk returns a tuple with the RequestedModalities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphCall) GetRequestedModalitiesOk() (*[]*AnyOfmicrosoftGraphModality, bool) {
	if o == nil || o.RequestedModalities == nil {
		return nil, false
	}
	return o.RequestedModalities, true
}

// HasRequestedModalities returns a boolean if a field has been set.
func (o *MicrosoftGraphCall) HasRequestedModalities() bool {
	if o != nil && o.RequestedModalities != nil {
		return true
	}

	return false
}

// SetRequestedModalities gets a reference to the given []*AnyOfmicrosoftGraphModality and assigns it to the RequestedModalities field.
func (o *MicrosoftGraphCall) SetRequestedModalities(v []*AnyOfmicrosoftGraphModality) {
	o.RequestedModalities = &v
}

// GetResultInfo returns the ResultInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphCall) GetResultInfo() AnyOfmicrosoftGraphResultInfo {
	if o == nil  {
		var ret AnyOfmicrosoftGraphResultInfo
		return ret
	}
	return o.ResultInfo
}

// GetResultInfoOk returns a tuple with the ResultInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphCall) GetResultInfoOk() (*AnyOfmicrosoftGraphResultInfo, bool) {
	if o == nil || o.ResultInfo == nil {
		return nil, false
	}
	return &o.ResultInfo, true
}

// HasResultInfo returns a boolean if a field has been set.
func (o *MicrosoftGraphCall) HasResultInfo() bool {
	if o != nil && o.ResultInfo != nil {
		return true
	}

	return false
}

// SetResultInfo gets a reference to the given AnyOfmicrosoftGraphResultInfo and assigns it to the ResultInfo field.
func (o *MicrosoftGraphCall) SetResultInfo(v AnyOfmicrosoftGraphResultInfo) {
	o.ResultInfo = v
}

// GetSource returns the Source field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphCall) GetSource() AnyOfmicrosoftGraphParticipantInfo {
	if o == nil  {
		var ret AnyOfmicrosoftGraphParticipantInfo
		return ret
	}
	return o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphCall) GetSourceOk() (*AnyOfmicrosoftGraphParticipantInfo, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return &o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *MicrosoftGraphCall) HasSource() bool {
	if o != nil && o.Source != nil {
		return true
	}

	return false
}

// SetSource gets a reference to the given AnyOfmicrosoftGraphParticipantInfo and assigns it to the Source field.
func (o *MicrosoftGraphCall) SetSource(v AnyOfmicrosoftGraphParticipantInfo) {
	o.Source = v
}

// GetState returns the State field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphCall) GetState() AnyOfmicrosoftGraphCallState {
	if o == nil  {
		var ret AnyOfmicrosoftGraphCallState
		return ret
	}
	return o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphCall) GetStateOk() (*AnyOfmicrosoftGraphCallState, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return &o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *MicrosoftGraphCall) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given AnyOfmicrosoftGraphCallState and assigns it to the State field.
func (o *MicrosoftGraphCall) SetState(v AnyOfmicrosoftGraphCallState) {
	o.State = v
}

// GetSubject returns the Subject field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphCall) GetSubject() string {
	if o == nil || o.Subject.Get() == nil {
		var ret string
		return ret
	}
	return *o.Subject.Get()
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphCall) GetSubjectOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Subject.Get(), o.Subject.IsSet()
}

// HasSubject returns a boolean if a field has been set.
func (o *MicrosoftGraphCall) HasSubject() bool {
	if o != nil && o.Subject.IsSet() {
		return true
	}

	return false
}

// SetSubject gets a reference to the given NullableString and assigns it to the Subject field.
func (o *MicrosoftGraphCall) SetSubject(v string) {
	o.Subject.Set(&v)
}
// SetSubjectNil sets the value for Subject to be an explicit nil
func (o *MicrosoftGraphCall) SetSubjectNil() {
	o.Subject.Set(nil)
}

// UnsetSubject ensures that no value is present for Subject, not even an explicit nil
func (o *MicrosoftGraphCall) UnsetSubject() {
	o.Subject.Unset()
}

// GetTargets returns the Targets field value if set, zero value otherwise.
func (o *MicrosoftGraphCall) GetTargets() []*AnyOfmicrosoftGraphInvitationParticipantInfo {
	if o == nil || o.Targets == nil {
		var ret []*AnyOfmicrosoftGraphInvitationParticipantInfo
		return ret
	}
	return *o.Targets
}

// GetTargetsOk returns a tuple with the Targets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphCall) GetTargetsOk() (*[]*AnyOfmicrosoftGraphInvitationParticipantInfo, bool) {
	if o == nil || o.Targets == nil {
		return nil, false
	}
	return o.Targets, true
}

// HasTargets returns a boolean if a field has been set.
func (o *MicrosoftGraphCall) HasTargets() bool {
	if o != nil && o.Targets != nil {
		return true
	}

	return false
}

// SetTargets gets a reference to the given []*AnyOfmicrosoftGraphInvitationParticipantInfo and assigns it to the Targets field.
func (o *MicrosoftGraphCall) SetTargets(v []*AnyOfmicrosoftGraphInvitationParticipantInfo) {
	o.Targets = &v
}

// GetTenantId returns the TenantId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphCall) GetTenantId() string {
	if o == nil || o.TenantId.Get() == nil {
		var ret string
		return ret
	}
	return *o.TenantId.Get()
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphCall) GetTenantIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TenantId.Get(), o.TenantId.IsSet()
}

// HasTenantId returns a boolean if a field has been set.
func (o *MicrosoftGraphCall) HasTenantId() bool {
	if o != nil && o.TenantId.IsSet() {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given NullableString and assigns it to the TenantId field.
func (o *MicrosoftGraphCall) SetTenantId(v string) {
	o.TenantId.Set(&v)
}
// SetTenantIdNil sets the value for TenantId to be an explicit nil
func (o *MicrosoftGraphCall) SetTenantIdNil() {
	o.TenantId.Set(nil)
}

// UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
func (o *MicrosoftGraphCall) UnsetTenantId() {
	o.TenantId.Unset()
}

// GetToneInfo returns the ToneInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphCall) GetToneInfo() AnyOfmicrosoftGraphToneInfo {
	if o == nil  {
		var ret AnyOfmicrosoftGraphToneInfo
		return ret
	}
	return o.ToneInfo
}

// GetToneInfoOk returns a tuple with the ToneInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphCall) GetToneInfoOk() (*AnyOfmicrosoftGraphToneInfo, bool) {
	if o == nil || o.ToneInfo == nil {
		return nil, false
	}
	return &o.ToneInfo, true
}

// HasToneInfo returns a boolean if a field has been set.
func (o *MicrosoftGraphCall) HasToneInfo() bool {
	if o != nil && o.ToneInfo != nil {
		return true
	}

	return false
}

// SetToneInfo gets a reference to the given AnyOfmicrosoftGraphToneInfo and assigns it to the ToneInfo field.
func (o *MicrosoftGraphCall) SetToneInfo(v AnyOfmicrosoftGraphToneInfo) {
	o.ToneInfo = v
}

// GetTranscription returns the Transcription field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphCall) GetTranscription() AnyOfmicrosoftGraphCallTranscriptionInfo {
	if o == nil  {
		var ret AnyOfmicrosoftGraphCallTranscriptionInfo
		return ret
	}
	return o.Transcription
}

// GetTranscriptionOk returns a tuple with the Transcription field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphCall) GetTranscriptionOk() (*AnyOfmicrosoftGraphCallTranscriptionInfo, bool) {
	if o == nil || o.Transcription == nil {
		return nil, false
	}
	return &o.Transcription, true
}

// HasTranscription returns a boolean if a field has been set.
func (o *MicrosoftGraphCall) HasTranscription() bool {
	if o != nil && o.Transcription != nil {
		return true
	}

	return false
}

// SetTranscription gets a reference to the given AnyOfmicrosoftGraphCallTranscriptionInfo and assigns it to the Transcription field.
func (o *MicrosoftGraphCall) SetTranscription(v AnyOfmicrosoftGraphCallTranscriptionInfo) {
	o.Transcription = v
}

// GetAudioRoutingGroups returns the AudioRoutingGroups field value if set, zero value otherwise.
func (o *MicrosoftGraphCall) GetAudioRoutingGroups() []MicrosoftGraphAudioRoutingGroup {
	if o == nil || o.AudioRoutingGroups == nil {
		var ret []MicrosoftGraphAudioRoutingGroup
		return ret
	}
	return *o.AudioRoutingGroups
}

// GetAudioRoutingGroupsOk returns a tuple with the AudioRoutingGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphCall) GetAudioRoutingGroupsOk() (*[]MicrosoftGraphAudioRoutingGroup, bool) {
	if o == nil || o.AudioRoutingGroups == nil {
		return nil, false
	}
	return o.AudioRoutingGroups, true
}

// HasAudioRoutingGroups returns a boolean if a field has been set.
func (o *MicrosoftGraphCall) HasAudioRoutingGroups() bool {
	if o != nil && o.AudioRoutingGroups != nil {
		return true
	}

	return false
}

// SetAudioRoutingGroups gets a reference to the given []MicrosoftGraphAudioRoutingGroup and assigns it to the AudioRoutingGroups field.
func (o *MicrosoftGraphCall) SetAudioRoutingGroups(v []MicrosoftGraphAudioRoutingGroup) {
	o.AudioRoutingGroups = &v
}

// GetOperations returns the Operations field value if set, zero value otherwise.
func (o *MicrosoftGraphCall) GetOperations() []MicrosoftGraphCommsOperation {
	if o == nil || o.Operations == nil {
		var ret []MicrosoftGraphCommsOperation
		return ret
	}
	return *o.Operations
}

// GetOperationsOk returns a tuple with the Operations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphCall) GetOperationsOk() (*[]MicrosoftGraphCommsOperation, bool) {
	if o == nil || o.Operations == nil {
		return nil, false
	}
	return o.Operations, true
}

// HasOperations returns a boolean if a field has been set.
func (o *MicrosoftGraphCall) HasOperations() bool {
	if o != nil && o.Operations != nil {
		return true
	}

	return false
}

// SetOperations gets a reference to the given []MicrosoftGraphCommsOperation and assigns it to the Operations field.
func (o *MicrosoftGraphCall) SetOperations(v []MicrosoftGraphCommsOperation) {
	o.Operations = &v
}

// GetParticipants returns the Participants field value if set, zero value otherwise.
func (o *MicrosoftGraphCall) GetParticipants() []MicrosoftGraphParticipant {
	if o == nil || o.Participants == nil {
		var ret []MicrosoftGraphParticipant
		return ret
	}
	return *o.Participants
}

// GetParticipantsOk returns a tuple with the Participants field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphCall) GetParticipantsOk() (*[]MicrosoftGraphParticipant, bool) {
	if o == nil || o.Participants == nil {
		return nil, false
	}
	return o.Participants, true
}

// HasParticipants returns a boolean if a field has been set.
func (o *MicrosoftGraphCall) HasParticipants() bool {
	if o != nil && o.Participants != nil {
		return true
	}

	return false
}

// SetParticipants gets a reference to the given []MicrosoftGraphParticipant and assigns it to the Participants field.
func (o *MicrosoftGraphCall) SetParticipants(v []MicrosoftGraphParticipant) {
	o.Participants = &v
}

func (o MicrosoftGraphCall) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.CallbackUri != nil {
		toSerialize["callbackUri"] = o.CallbackUri
	}
	if o.CallChainId.IsSet() {
		toSerialize["callChainId"] = o.CallChainId.Get()
	}
	if o.CallOptions != nil {
		toSerialize["callOptions"] = o.CallOptions
	}
	if o.CallRoutes != nil {
		toSerialize["callRoutes"] = o.CallRoutes
	}
	if o.ChatInfo != nil {
		toSerialize["chatInfo"] = o.ChatInfo
	}
	if o.Direction != nil {
		toSerialize["direction"] = o.Direction
	}
	if o.IncomingContext != nil {
		toSerialize["incomingContext"] = o.IncomingContext
	}
	if o.MediaConfig != nil {
		toSerialize["mediaConfig"] = o.MediaConfig
	}
	if o.MediaState != nil {
		toSerialize["mediaState"] = o.MediaState
	}
	if o.MeetingInfo != nil {
		toSerialize["meetingInfo"] = o.MeetingInfo
	}
	if o.MyParticipantId.IsSet() {
		toSerialize["myParticipantId"] = o.MyParticipantId.Get()
	}
	if o.RequestedModalities != nil {
		toSerialize["requestedModalities"] = o.RequestedModalities
	}
	if o.ResultInfo != nil {
		toSerialize["resultInfo"] = o.ResultInfo
	}
	if o.Source != nil {
		toSerialize["source"] = o.Source
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.Subject.IsSet() {
		toSerialize["subject"] = o.Subject.Get()
	}
	if o.Targets != nil {
		toSerialize["targets"] = o.Targets
	}
	if o.TenantId.IsSet() {
		toSerialize["tenantId"] = o.TenantId.Get()
	}
	if o.ToneInfo != nil {
		toSerialize["toneInfo"] = o.ToneInfo
	}
	if o.Transcription != nil {
		toSerialize["transcription"] = o.Transcription
	}
	if o.AudioRoutingGroups != nil {
		toSerialize["audioRoutingGroups"] = o.AudioRoutingGroups
	}
	if o.Operations != nil {
		toSerialize["operations"] = o.Operations
	}
	if o.Participants != nil {
		toSerialize["participants"] = o.Participants
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphCall struct {
	value *MicrosoftGraphCall
	isSet bool
}

func (v NullableMicrosoftGraphCall) Get() *MicrosoftGraphCall {
	return v.value
}

func (v *NullableMicrosoftGraphCall) Set(val *MicrosoftGraphCall) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphCall) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphCall) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphCall(val *MicrosoftGraphCall) *NullableMicrosoftGraphCall {
	return &NullableMicrosoftGraphCall{value: val, isSet: true}
}

func (v NullableMicrosoftGraphCall) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphCall) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


