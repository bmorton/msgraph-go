# MicrosoftGraphMediaStream

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Direction** | Pointer to [**AnyOfmicrosoftGraphMediaDirection**](anyOf&lt;microsoft.graph.mediaDirection&gt;.md) | The direction. The possible values are inactive, sendOnly, receiveOnly, sendReceive. | [optional] 
**Label** | Pointer to **NullableString** | The media stream label. | [optional] 
**MediaType** | Pointer to [**AnyOfmicrosoftGraphModality**](anyOf&lt;microsoft.graph.modality&gt;.md) | The media type. The possible value are unknown, audio, video, videoBasedScreenSharing, data. | [optional] 
**ServerMuted** | Pointer to **bool** | If the media is muted by the server. | [optional] 
**SourceId** | Pointer to **string** | The source ID. | [optional] 

## Methods

### NewMicrosoftGraphMediaStream

`func NewMicrosoftGraphMediaStream() *MicrosoftGraphMediaStream`

NewMicrosoftGraphMediaStream instantiates a new MicrosoftGraphMediaStream object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphMediaStreamWithDefaults

`func NewMicrosoftGraphMediaStreamWithDefaults() *MicrosoftGraphMediaStream`

NewMicrosoftGraphMediaStreamWithDefaults instantiates a new MicrosoftGraphMediaStream object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDirection

`func (o *MicrosoftGraphMediaStream) GetDirection() AnyOfmicrosoftGraphMediaDirection`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *MicrosoftGraphMediaStream) GetDirectionOk() (*AnyOfmicrosoftGraphMediaDirection, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *MicrosoftGraphMediaStream) SetDirection(v AnyOfmicrosoftGraphMediaDirection)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *MicrosoftGraphMediaStream) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### SetDirectionNil

`func (o *MicrosoftGraphMediaStream) SetDirectionNil(b bool)`

 SetDirectionNil sets the value for Direction to be an explicit nil

### UnsetDirection
`func (o *MicrosoftGraphMediaStream) UnsetDirection()`

UnsetDirection ensures that no value is present for Direction, not even an explicit nil
### GetLabel

`func (o *MicrosoftGraphMediaStream) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *MicrosoftGraphMediaStream) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *MicrosoftGraphMediaStream) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *MicrosoftGraphMediaStream) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### SetLabelNil

`func (o *MicrosoftGraphMediaStream) SetLabelNil(b bool)`

 SetLabelNil sets the value for Label to be an explicit nil

### UnsetLabel
`func (o *MicrosoftGraphMediaStream) UnsetLabel()`

UnsetLabel ensures that no value is present for Label, not even an explicit nil
### GetMediaType

`func (o *MicrosoftGraphMediaStream) GetMediaType() AnyOfmicrosoftGraphModality`

GetMediaType returns the MediaType field if non-nil, zero value otherwise.

### GetMediaTypeOk

`func (o *MicrosoftGraphMediaStream) GetMediaTypeOk() (*AnyOfmicrosoftGraphModality, bool)`

GetMediaTypeOk returns a tuple with the MediaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaType

`func (o *MicrosoftGraphMediaStream) SetMediaType(v AnyOfmicrosoftGraphModality)`

SetMediaType sets MediaType field to given value.

### HasMediaType

`func (o *MicrosoftGraphMediaStream) HasMediaType() bool`

HasMediaType returns a boolean if a field has been set.

### SetMediaTypeNil

`func (o *MicrosoftGraphMediaStream) SetMediaTypeNil(b bool)`

 SetMediaTypeNil sets the value for MediaType to be an explicit nil

### UnsetMediaType
`func (o *MicrosoftGraphMediaStream) UnsetMediaType()`

UnsetMediaType ensures that no value is present for MediaType, not even an explicit nil
### GetServerMuted

`func (o *MicrosoftGraphMediaStream) GetServerMuted() bool`

GetServerMuted returns the ServerMuted field if non-nil, zero value otherwise.

### GetServerMutedOk

`func (o *MicrosoftGraphMediaStream) GetServerMutedOk() (*bool, bool)`

GetServerMutedOk returns a tuple with the ServerMuted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerMuted

`func (o *MicrosoftGraphMediaStream) SetServerMuted(v bool)`

SetServerMuted sets ServerMuted field to given value.

### HasServerMuted

`func (o *MicrosoftGraphMediaStream) HasServerMuted() bool`

HasServerMuted returns a boolean if a field has been set.

### GetSourceId

`func (o *MicrosoftGraphMediaStream) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *MicrosoftGraphMediaStream) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *MicrosoftGraphMediaStream) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *MicrosoftGraphMediaStream) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


