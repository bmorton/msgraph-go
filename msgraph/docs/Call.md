# Call

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallbackUri** | Pointer to **string** | The callback URL on which callbacks will be delivered. Must be https. | [optional] 
**CallChainId** | Pointer to **NullableString** | A unique identifier for all the participant calls in a conference or a unique identifier for two participant calls in a P2P call.  This needs to be copied over from Microsoft.Graph.Call.CallChainId. | [optional] 
**CallOptions** | Pointer to [**AnyOfobject**](anyOf&lt;object&gt;.md) |  | [optional] 
**CallRoutes** | Pointer to [**[]AnyOfmicrosoftGraphCallRoute**](AnyOfmicrosoftGraphCallRoute.md) | The routing information on how the call was retargeted. Read-only. | [optional] 
**ChatInfo** | Pointer to [**AnyOfmicrosoftGraphChatInfo**](anyOf&lt;microsoft.graph.chatInfo&gt;.md) | The chat information. Required information for joining a meeting. | [optional] 
**Direction** | Pointer to [**AnyOfmicrosoftGraphCallDirection**](anyOf&lt;microsoft.graph.callDirection&gt;.md) | The direction of the call. The possible value are incoming or outgoing. Read-only. | [optional] 
**IncomingContext** | Pointer to [**AnyOfmicrosoftGraphIncomingContext**](anyOf&lt;microsoft.graph.incomingContext&gt;.md) | The context associated with an incoming call. Read-only. Server generated. | [optional] 
**MediaConfig** | Pointer to [**AnyOfobject**](anyOf&lt;object&gt;.md) | The media configuration. Required. | [optional] 
**MediaState** | Pointer to [**AnyOfmicrosoftGraphCallMediaState**](anyOf&lt;microsoft.graph.callMediaState&gt;.md) | Read-only. The call media state. | [optional] 
**MeetingInfo** | Pointer to [**AnyOfobject**](anyOf&lt;object&gt;.md) | The meeting information that&#39;s required for joining a meeting. | [optional] 
**MyParticipantId** | Pointer to **NullableString** |  | [optional] 
**RequestedModalities** | Pointer to [**[]AnyOfmicrosoftGraphModality**](AnyOfmicrosoftGraphModality.md) |  | [optional] 
**ResultInfo** | Pointer to [**AnyOfmicrosoftGraphResultInfo**](anyOf&lt;microsoft.graph.resultInfo&gt;.md) |  | [optional] 
**Source** | Pointer to [**AnyOfmicrosoftGraphParticipantInfo**](anyOf&lt;microsoft.graph.participantInfo&gt;.md) |  | [optional] 
**State** | Pointer to [**AnyOfmicrosoftGraphCallState**](anyOf&lt;microsoft.graph.callState&gt;.md) |  | [optional] 
**Subject** | Pointer to **NullableString** |  | [optional] 
**Targets** | Pointer to [**[]AnyOfmicrosoftGraphInvitationParticipantInfo**](AnyOfmicrosoftGraphInvitationParticipantInfo.md) |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**ToneInfo** | Pointer to [**AnyOfmicrosoftGraphToneInfo**](anyOf&lt;microsoft.graph.toneInfo&gt;.md) |  | [optional] 
**Transcription** | Pointer to [**AnyOfmicrosoftGraphCallTranscriptionInfo**](anyOf&lt;microsoft.graph.callTranscriptionInfo&gt;.md) | The transcription information for the call. Read-only. | [optional] 
**AudioRoutingGroups** | Pointer to [**[]MicrosoftGraphAudioRoutingGroup**](MicrosoftGraphAudioRoutingGroup.md) | Read-only. Nullable. | [optional] 
**Operations** | Pointer to [**[]MicrosoftGraphCommsOperation**](MicrosoftGraphCommsOperation.md) | Read-only. Nullable. | [optional] 
**Participants** | Pointer to [**[]MicrosoftGraphParticipant**](MicrosoftGraphParticipant.md) | Read-only. Nullable. | [optional] 

## Methods

### NewCall

`func NewCall() *Call`

NewCall instantiates a new Call object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCallWithDefaults

`func NewCallWithDefaults() *Call`

NewCallWithDefaults instantiates a new Call object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallbackUri

`func (o *Call) GetCallbackUri() string`

GetCallbackUri returns the CallbackUri field if non-nil, zero value otherwise.

### GetCallbackUriOk

`func (o *Call) GetCallbackUriOk() (*string, bool)`

GetCallbackUriOk returns a tuple with the CallbackUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackUri

`func (o *Call) SetCallbackUri(v string)`

SetCallbackUri sets CallbackUri field to given value.

### HasCallbackUri

`func (o *Call) HasCallbackUri() bool`

HasCallbackUri returns a boolean if a field has been set.

### GetCallChainId

`func (o *Call) GetCallChainId() string`

GetCallChainId returns the CallChainId field if non-nil, zero value otherwise.

### GetCallChainIdOk

`func (o *Call) GetCallChainIdOk() (*string, bool)`

GetCallChainIdOk returns a tuple with the CallChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallChainId

`func (o *Call) SetCallChainId(v string)`

SetCallChainId sets CallChainId field to given value.

### HasCallChainId

`func (o *Call) HasCallChainId() bool`

HasCallChainId returns a boolean if a field has been set.

### SetCallChainIdNil

`func (o *Call) SetCallChainIdNil(b bool)`

 SetCallChainIdNil sets the value for CallChainId to be an explicit nil

### UnsetCallChainId
`func (o *Call) UnsetCallChainId()`

UnsetCallChainId ensures that no value is present for CallChainId, not even an explicit nil
### GetCallOptions

`func (o *Call) GetCallOptions() AnyOfobject`

GetCallOptions returns the CallOptions field if non-nil, zero value otherwise.

### GetCallOptionsOk

`func (o *Call) GetCallOptionsOk() (*AnyOfobject, bool)`

GetCallOptionsOk returns a tuple with the CallOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallOptions

`func (o *Call) SetCallOptions(v AnyOfobject)`

SetCallOptions sets CallOptions field to given value.

### HasCallOptions

`func (o *Call) HasCallOptions() bool`

HasCallOptions returns a boolean if a field has been set.

### SetCallOptionsNil

`func (o *Call) SetCallOptionsNil(b bool)`

 SetCallOptionsNil sets the value for CallOptions to be an explicit nil

### UnsetCallOptions
`func (o *Call) UnsetCallOptions()`

UnsetCallOptions ensures that no value is present for CallOptions, not even an explicit nil
### GetCallRoutes

`func (o *Call) GetCallRoutes() []*AnyOfmicrosoftGraphCallRoute`

GetCallRoutes returns the CallRoutes field if non-nil, zero value otherwise.

### GetCallRoutesOk

`func (o *Call) GetCallRoutesOk() (*[]*AnyOfmicrosoftGraphCallRoute, bool)`

GetCallRoutesOk returns a tuple with the CallRoutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallRoutes

`func (o *Call) SetCallRoutes(v []*AnyOfmicrosoftGraphCallRoute)`

SetCallRoutes sets CallRoutes field to given value.

### HasCallRoutes

`func (o *Call) HasCallRoutes() bool`

HasCallRoutes returns a boolean if a field has been set.

### GetChatInfo

`func (o *Call) GetChatInfo() AnyOfmicrosoftGraphChatInfo`

GetChatInfo returns the ChatInfo field if non-nil, zero value otherwise.

### GetChatInfoOk

`func (o *Call) GetChatInfoOk() (*AnyOfmicrosoftGraphChatInfo, bool)`

GetChatInfoOk returns a tuple with the ChatInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatInfo

`func (o *Call) SetChatInfo(v AnyOfmicrosoftGraphChatInfo)`

SetChatInfo sets ChatInfo field to given value.

### HasChatInfo

`func (o *Call) HasChatInfo() bool`

HasChatInfo returns a boolean if a field has been set.

### SetChatInfoNil

`func (o *Call) SetChatInfoNil(b bool)`

 SetChatInfoNil sets the value for ChatInfo to be an explicit nil

### UnsetChatInfo
`func (o *Call) UnsetChatInfo()`

UnsetChatInfo ensures that no value is present for ChatInfo, not even an explicit nil
### GetDirection

`func (o *Call) GetDirection() AnyOfmicrosoftGraphCallDirection`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *Call) GetDirectionOk() (*AnyOfmicrosoftGraphCallDirection, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *Call) SetDirection(v AnyOfmicrosoftGraphCallDirection)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *Call) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### SetDirectionNil

`func (o *Call) SetDirectionNil(b bool)`

 SetDirectionNil sets the value for Direction to be an explicit nil

### UnsetDirection
`func (o *Call) UnsetDirection()`

UnsetDirection ensures that no value is present for Direction, not even an explicit nil
### GetIncomingContext

`func (o *Call) GetIncomingContext() AnyOfmicrosoftGraphIncomingContext`

GetIncomingContext returns the IncomingContext field if non-nil, zero value otherwise.

### GetIncomingContextOk

`func (o *Call) GetIncomingContextOk() (*AnyOfmicrosoftGraphIncomingContext, bool)`

GetIncomingContextOk returns a tuple with the IncomingContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncomingContext

`func (o *Call) SetIncomingContext(v AnyOfmicrosoftGraphIncomingContext)`

SetIncomingContext sets IncomingContext field to given value.

### HasIncomingContext

`func (o *Call) HasIncomingContext() bool`

HasIncomingContext returns a boolean if a field has been set.

### SetIncomingContextNil

`func (o *Call) SetIncomingContextNil(b bool)`

 SetIncomingContextNil sets the value for IncomingContext to be an explicit nil

### UnsetIncomingContext
`func (o *Call) UnsetIncomingContext()`

UnsetIncomingContext ensures that no value is present for IncomingContext, not even an explicit nil
### GetMediaConfig

`func (o *Call) GetMediaConfig() AnyOfobject`

GetMediaConfig returns the MediaConfig field if non-nil, zero value otherwise.

### GetMediaConfigOk

`func (o *Call) GetMediaConfigOk() (*AnyOfobject, bool)`

GetMediaConfigOk returns a tuple with the MediaConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaConfig

`func (o *Call) SetMediaConfig(v AnyOfobject)`

SetMediaConfig sets MediaConfig field to given value.

### HasMediaConfig

`func (o *Call) HasMediaConfig() bool`

HasMediaConfig returns a boolean if a field has been set.

### SetMediaConfigNil

`func (o *Call) SetMediaConfigNil(b bool)`

 SetMediaConfigNil sets the value for MediaConfig to be an explicit nil

### UnsetMediaConfig
`func (o *Call) UnsetMediaConfig()`

UnsetMediaConfig ensures that no value is present for MediaConfig, not even an explicit nil
### GetMediaState

`func (o *Call) GetMediaState() AnyOfmicrosoftGraphCallMediaState`

GetMediaState returns the MediaState field if non-nil, zero value otherwise.

### GetMediaStateOk

`func (o *Call) GetMediaStateOk() (*AnyOfmicrosoftGraphCallMediaState, bool)`

GetMediaStateOk returns a tuple with the MediaState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaState

`func (o *Call) SetMediaState(v AnyOfmicrosoftGraphCallMediaState)`

SetMediaState sets MediaState field to given value.

### HasMediaState

`func (o *Call) HasMediaState() bool`

HasMediaState returns a boolean if a field has been set.

### SetMediaStateNil

`func (o *Call) SetMediaStateNil(b bool)`

 SetMediaStateNil sets the value for MediaState to be an explicit nil

### UnsetMediaState
`func (o *Call) UnsetMediaState()`

UnsetMediaState ensures that no value is present for MediaState, not even an explicit nil
### GetMeetingInfo

`func (o *Call) GetMeetingInfo() AnyOfobject`

GetMeetingInfo returns the MeetingInfo field if non-nil, zero value otherwise.

### GetMeetingInfoOk

`func (o *Call) GetMeetingInfoOk() (*AnyOfobject, bool)`

GetMeetingInfoOk returns a tuple with the MeetingInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeetingInfo

`func (o *Call) SetMeetingInfo(v AnyOfobject)`

SetMeetingInfo sets MeetingInfo field to given value.

### HasMeetingInfo

`func (o *Call) HasMeetingInfo() bool`

HasMeetingInfo returns a boolean if a field has been set.

### SetMeetingInfoNil

`func (o *Call) SetMeetingInfoNil(b bool)`

 SetMeetingInfoNil sets the value for MeetingInfo to be an explicit nil

### UnsetMeetingInfo
`func (o *Call) UnsetMeetingInfo()`

UnsetMeetingInfo ensures that no value is present for MeetingInfo, not even an explicit nil
### GetMyParticipantId

`func (o *Call) GetMyParticipantId() string`

GetMyParticipantId returns the MyParticipantId field if non-nil, zero value otherwise.

### GetMyParticipantIdOk

`func (o *Call) GetMyParticipantIdOk() (*string, bool)`

GetMyParticipantIdOk returns a tuple with the MyParticipantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMyParticipantId

`func (o *Call) SetMyParticipantId(v string)`

SetMyParticipantId sets MyParticipantId field to given value.

### HasMyParticipantId

`func (o *Call) HasMyParticipantId() bool`

HasMyParticipantId returns a boolean if a field has been set.

### SetMyParticipantIdNil

`func (o *Call) SetMyParticipantIdNil(b bool)`

 SetMyParticipantIdNil sets the value for MyParticipantId to be an explicit nil

### UnsetMyParticipantId
`func (o *Call) UnsetMyParticipantId()`

UnsetMyParticipantId ensures that no value is present for MyParticipantId, not even an explicit nil
### GetRequestedModalities

`func (o *Call) GetRequestedModalities() []*AnyOfmicrosoftGraphModality`

GetRequestedModalities returns the RequestedModalities field if non-nil, zero value otherwise.

### GetRequestedModalitiesOk

`func (o *Call) GetRequestedModalitiesOk() (*[]*AnyOfmicrosoftGraphModality, bool)`

GetRequestedModalitiesOk returns a tuple with the RequestedModalities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedModalities

`func (o *Call) SetRequestedModalities(v []*AnyOfmicrosoftGraphModality)`

SetRequestedModalities sets RequestedModalities field to given value.

### HasRequestedModalities

`func (o *Call) HasRequestedModalities() bool`

HasRequestedModalities returns a boolean if a field has been set.

### GetResultInfo

`func (o *Call) GetResultInfo() AnyOfmicrosoftGraphResultInfo`

GetResultInfo returns the ResultInfo field if non-nil, zero value otherwise.

### GetResultInfoOk

`func (o *Call) GetResultInfoOk() (*AnyOfmicrosoftGraphResultInfo, bool)`

GetResultInfoOk returns a tuple with the ResultInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultInfo

`func (o *Call) SetResultInfo(v AnyOfmicrosoftGraphResultInfo)`

SetResultInfo sets ResultInfo field to given value.

### HasResultInfo

`func (o *Call) HasResultInfo() bool`

HasResultInfo returns a boolean if a field has been set.

### SetResultInfoNil

`func (o *Call) SetResultInfoNil(b bool)`

 SetResultInfoNil sets the value for ResultInfo to be an explicit nil

### UnsetResultInfo
`func (o *Call) UnsetResultInfo()`

UnsetResultInfo ensures that no value is present for ResultInfo, not even an explicit nil
### GetSource

`func (o *Call) GetSource() AnyOfmicrosoftGraphParticipantInfo`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *Call) GetSourceOk() (*AnyOfmicrosoftGraphParticipantInfo, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *Call) SetSource(v AnyOfmicrosoftGraphParticipantInfo)`

SetSource sets Source field to given value.

### HasSource

`func (o *Call) HasSource() bool`

HasSource returns a boolean if a field has been set.

### SetSourceNil

`func (o *Call) SetSourceNil(b bool)`

 SetSourceNil sets the value for Source to be an explicit nil

### UnsetSource
`func (o *Call) UnsetSource()`

UnsetSource ensures that no value is present for Source, not even an explicit nil
### GetState

`func (o *Call) GetState() AnyOfmicrosoftGraphCallState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Call) GetStateOk() (*AnyOfmicrosoftGraphCallState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Call) SetState(v AnyOfmicrosoftGraphCallState)`

SetState sets State field to given value.

### HasState

`func (o *Call) HasState() bool`

HasState returns a boolean if a field has been set.

### SetStateNil

`func (o *Call) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *Call) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil
### GetSubject

`func (o *Call) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *Call) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *Call) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *Call) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### SetSubjectNil

`func (o *Call) SetSubjectNil(b bool)`

 SetSubjectNil sets the value for Subject to be an explicit nil

### UnsetSubject
`func (o *Call) UnsetSubject()`

UnsetSubject ensures that no value is present for Subject, not even an explicit nil
### GetTargets

`func (o *Call) GetTargets() []*AnyOfmicrosoftGraphInvitationParticipantInfo`

GetTargets returns the Targets field if non-nil, zero value otherwise.

### GetTargetsOk

`func (o *Call) GetTargetsOk() (*[]*AnyOfmicrosoftGraphInvitationParticipantInfo, bool)`

GetTargetsOk returns a tuple with the Targets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargets

`func (o *Call) SetTargets(v []*AnyOfmicrosoftGraphInvitationParticipantInfo)`

SetTargets sets Targets field to given value.

### HasTargets

`func (o *Call) HasTargets() bool`

HasTargets returns a boolean if a field has been set.

### GetTenantId

`func (o *Call) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *Call) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *Call) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *Call) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *Call) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *Call) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetToneInfo

`func (o *Call) GetToneInfo() AnyOfmicrosoftGraphToneInfo`

GetToneInfo returns the ToneInfo field if non-nil, zero value otherwise.

### GetToneInfoOk

`func (o *Call) GetToneInfoOk() (*AnyOfmicrosoftGraphToneInfo, bool)`

GetToneInfoOk returns a tuple with the ToneInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToneInfo

`func (o *Call) SetToneInfo(v AnyOfmicrosoftGraphToneInfo)`

SetToneInfo sets ToneInfo field to given value.

### HasToneInfo

`func (o *Call) HasToneInfo() bool`

HasToneInfo returns a boolean if a field has been set.

### SetToneInfoNil

`func (o *Call) SetToneInfoNil(b bool)`

 SetToneInfoNil sets the value for ToneInfo to be an explicit nil

### UnsetToneInfo
`func (o *Call) UnsetToneInfo()`

UnsetToneInfo ensures that no value is present for ToneInfo, not even an explicit nil
### GetTranscription

`func (o *Call) GetTranscription() AnyOfmicrosoftGraphCallTranscriptionInfo`

GetTranscription returns the Transcription field if non-nil, zero value otherwise.

### GetTranscriptionOk

`func (o *Call) GetTranscriptionOk() (*AnyOfmicrosoftGraphCallTranscriptionInfo, bool)`

GetTranscriptionOk returns a tuple with the Transcription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscription

`func (o *Call) SetTranscription(v AnyOfmicrosoftGraphCallTranscriptionInfo)`

SetTranscription sets Transcription field to given value.

### HasTranscription

`func (o *Call) HasTranscription() bool`

HasTranscription returns a boolean if a field has been set.

### SetTranscriptionNil

`func (o *Call) SetTranscriptionNil(b bool)`

 SetTranscriptionNil sets the value for Transcription to be an explicit nil

### UnsetTranscription
`func (o *Call) UnsetTranscription()`

UnsetTranscription ensures that no value is present for Transcription, not even an explicit nil
### GetAudioRoutingGroups

`func (o *Call) GetAudioRoutingGroups() []MicrosoftGraphAudioRoutingGroup`

GetAudioRoutingGroups returns the AudioRoutingGroups field if non-nil, zero value otherwise.

### GetAudioRoutingGroupsOk

`func (o *Call) GetAudioRoutingGroupsOk() (*[]MicrosoftGraphAudioRoutingGroup, bool)`

GetAudioRoutingGroupsOk returns a tuple with the AudioRoutingGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioRoutingGroups

`func (o *Call) SetAudioRoutingGroups(v []MicrosoftGraphAudioRoutingGroup)`

SetAudioRoutingGroups sets AudioRoutingGroups field to given value.

### HasAudioRoutingGroups

`func (o *Call) HasAudioRoutingGroups() bool`

HasAudioRoutingGroups returns a boolean if a field has been set.

### GetOperations

`func (o *Call) GetOperations() []MicrosoftGraphCommsOperation`

GetOperations returns the Operations field if non-nil, zero value otherwise.

### GetOperationsOk

`func (o *Call) GetOperationsOk() (*[]MicrosoftGraphCommsOperation, bool)`

GetOperationsOk returns a tuple with the Operations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperations

`func (o *Call) SetOperations(v []MicrosoftGraphCommsOperation)`

SetOperations sets Operations field to given value.

### HasOperations

`func (o *Call) HasOperations() bool`

HasOperations returns a boolean if a field has been set.

### GetParticipants

`func (o *Call) GetParticipants() []MicrosoftGraphParticipant`

GetParticipants returns the Participants field if non-nil, zero value otherwise.

### GetParticipantsOk

`func (o *Call) GetParticipantsOk() (*[]MicrosoftGraphParticipant, bool)`

GetParticipantsOk returns a tuple with the Participants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipants

`func (o *Call) SetParticipants(v []MicrosoftGraphParticipant)`

SetParticipants sets Participants field to given value.

### HasParticipants

`func (o *Call) HasParticipants() bool`

HasParticipants returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


