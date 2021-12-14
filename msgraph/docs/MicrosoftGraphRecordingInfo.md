# MicrosoftGraphRecordingInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Initiator** | Pointer to [**AnyOfmicrosoftGraphIdentitySet**](anyOf&lt;microsoft.graph.identitySet&gt;.md) | The identities of the recording initiator. | [optional] 
**RecordingStatus** | Pointer to [**AnyOfmicrosoftGraphRecordingStatus**](anyOf&lt;microsoft.graph.recordingStatus&gt;.md) | Possible values are: unknown, notRecording, recording, or failed. | [optional] 

## Methods

### NewMicrosoftGraphRecordingInfo

`func NewMicrosoftGraphRecordingInfo() *MicrosoftGraphRecordingInfo`

NewMicrosoftGraphRecordingInfo instantiates a new MicrosoftGraphRecordingInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphRecordingInfoWithDefaults

`func NewMicrosoftGraphRecordingInfoWithDefaults() *MicrosoftGraphRecordingInfo`

NewMicrosoftGraphRecordingInfoWithDefaults instantiates a new MicrosoftGraphRecordingInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInitiator

`func (o *MicrosoftGraphRecordingInfo) GetInitiator() AnyOfmicrosoftGraphIdentitySet`

GetInitiator returns the Initiator field if non-nil, zero value otherwise.

### GetInitiatorOk

`func (o *MicrosoftGraphRecordingInfo) GetInitiatorOk() (*AnyOfmicrosoftGraphIdentitySet, bool)`

GetInitiatorOk returns a tuple with the Initiator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiator

`func (o *MicrosoftGraphRecordingInfo) SetInitiator(v AnyOfmicrosoftGraphIdentitySet)`

SetInitiator sets Initiator field to given value.

### HasInitiator

`func (o *MicrosoftGraphRecordingInfo) HasInitiator() bool`

HasInitiator returns a boolean if a field has been set.

### SetInitiatorNil

`func (o *MicrosoftGraphRecordingInfo) SetInitiatorNil(b bool)`

 SetInitiatorNil sets the value for Initiator to be an explicit nil

### UnsetInitiator
`func (o *MicrosoftGraphRecordingInfo) UnsetInitiator()`

UnsetInitiator ensures that no value is present for Initiator, not even an explicit nil
### GetRecordingStatus

`func (o *MicrosoftGraphRecordingInfo) GetRecordingStatus() AnyOfmicrosoftGraphRecordingStatus`

GetRecordingStatus returns the RecordingStatus field if non-nil, zero value otherwise.

### GetRecordingStatusOk

`func (o *MicrosoftGraphRecordingInfo) GetRecordingStatusOk() (*AnyOfmicrosoftGraphRecordingStatus, bool)`

GetRecordingStatusOk returns a tuple with the RecordingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordingStatus

`func (o *MicrosoftGraphRecordingInfo) SetRecordingStatus(v AnyOfmicrosoftGraphRecordingStatus)`

SetRecordingStatus sets RecordingStatus field to given value.

### HasRecordingStatus

`func (o *MicrosoftGraphRecordingInfo) HasRecordingStatus() bool`

HasRecordingStatus returns a boolean if a field has been set.

### SetRecordingStatusNil

`func (o *MicrosoftGraphRecordingInfo) SetRecordingStatusNil(b bool)`

 SetRecordingStatusNil sets the value for RecordingStatus to be an explicit nil

### UnsetRecordingStatus
`func (o *MicrosoftGraphRecordingInfo) UnsetRecordingStatus()`

UnsetRecordingStatus ensures that no value is present for RecordingStatus, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


