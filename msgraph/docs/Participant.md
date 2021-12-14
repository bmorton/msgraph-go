# Participant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Info** | Pointer to [**MicrosoftGraphParticipantInfo**](MicrosoftGraphParticipantInfo.md) |  | [optional] 
**IsInLobby** | Pointer to **bool** | true if the participant is in lobby. | [optional] 
**IsMuted** | Pointer to **bool** | true if the participant is muted (client or server muted). | [optional] 
**MediaStreams** | Pointer to [**[]AnyOfmicrosoftGraphMediaStream**](AnyOfmicrosoftGraphMediaStream.md) | The list of media streams. | [optional] 
**Metadata** | Pointer to **NullableString** | A blob of data provided by the participant in the roster. | [optional] 
**RecordingInfo** | Pointer to [**AnyOfmicrosoftGraphRecordingInfo**](anyOf&lt;microsoft.graph.recordingInfo&gt;.md) | Information about whether the participant has recording capability. | [optional] 

## Methods

### NewParticipant

`func NewParticipant() *Participant`

NewParticipant instantiates a new Participant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParticipantWithDefaults

`func NewParticipantWithDefaults() *Participant`

NewParticipantWithDefaults instantiates a new Participant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInfo

`func (o *Participant) GetInfo() MicrosoftGraphParticipantInfo`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *Participant) GetInfoOk() (*MicrosoftGraphParticipantInfo, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *Participant) SetInfo(v MicrosoftGraphParticipantInfo)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *Participant) HasInfo() bool`

HasInfo returns a boolean if a field has been set.

### GetIsInLobby

`func (o *Participant) GetIsInLobby() bool`

GetIsInLobby returns the IsInLobby field if non-nil, zero value otherwise.

### GetIsInLobbyOk

`func (o *Participant) GetIsInLobbyOk() (*bool, bool)`

GetIsInLobbyOk returns a tuple with the IsInLobby field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInLobby

`func (o *Participant) SetIsInLobby(v bool)`

SetIsInLobby sets IsInLobby field to given value.

### HasIsInLobby

`func (o *Participant) HasIsInLobby() bool`

HasIsInLobby returns a boolean if a field has been set.

### GetIsMuted

`func (o *Participant) GetIsMuted() bool`

GetIsMuted returns the IsMuted field if non-nil, zero value otherwise.

### GetIsMutedOk

`func (o *Participant) GetIsMutedOk() (*bool, bool)`

GetIsMutedOk returns a tuple with the IsMuted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMuted

`func (o *Participant) SetIsMuted(v bool)`

SetIsMuted sets IsMuted field to given value.

### HasIsMuted

`func (o *Participant) HasIsMuted() bool`

HasIsMuted returns a boolean if a field has been set.

### GetMediaStreams

`func (o *Participant) GetMediaStreams() []*AnyOfmicrosoftGraphMediaStream`

GetMediaStreams returns the MediaStreams field if non-nil, zero value otherwise.

### GetMediaStreamsOk

`func (o *Participant) GetMediaStreamsOk() (*[]*AnyOfmicrosoftGraphMediaStream, bool)`

GetMediaStreamsOk returns a tuple with the MediaStreams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaStreams

`func (o *Participant) SetMediaStreams(v []*AnyOfmicrosoftGraphMediaStream)`

SetMediaStreams sets MediaStreams field to given value.

### HasMediaStreams

`func (o *Participant) HasMediaStreams() bool`

HasMediaStreams returns a boolean if a field has been set.

### GetMetadata

`func (o *Participant) GetMetadata() string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Participant) GetMetadataOk() (*string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Participant) SetMetadata(v string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Participant) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *Participant) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *Participant) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetRecordingInfo

`func (o *Participant) GetRecordingInfo() AnyOfmicrosoftGraphRecordingInfo`

GetRecordingInfo returns the RecordingInfo field if non-nil, zero value otherwise.

### GetRecordingInfoOk

`func (o *Participant) GetRecordingInfoOk() (*AnyOfmicrosoftGraphRecordingInfo, bool)`

GetRecordingInfoOk returns a tuple with the RecordingInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordingInfo

`func (o *Participant) SetRecordingInfo(v AnyOfmicrosoftGraphRecordingInfo)`

SetRecordingInfo sets RecordingInfo field to given value.

### HasRecordingInfo

`func (o *Participant) HasRecordingInfo() bool`

HasRecordingInfo returns a boolean if a field has been set.

### SetRecordingInfoNil

`func (o *Participant) SetRecordingInfoNil(b bool)`

 SetRecordingInfoNil sets the value for RecordingInfo to be an explicit nil

### UnsetRecordingInfo
`func (o *Participant) UnsetRecordingInfo()`

UnsetRecordingInfo ensures that no value is present for RecordingInfo, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


