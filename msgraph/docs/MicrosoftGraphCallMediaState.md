# MicrosoftGraphCallMediaState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Audio** | Pointer to [**AnyOfmicrosoftGraphMediaState**](anyOf&lt;microsoft.graph.mediaState&gt;.md) | The audio media state. Possible values are: active, inactive, unknownFutureValue. | [optional] 

## Methods

### NewMicrosoftGraphCallMediaState

`func NewMicrosoftGraphCallMediaState() *MicrosoftGraphCallMediaState`

NewMicrosoftGraphCallMediaState instantiates a new MicrosoftGraphCallMediaState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphCallMediaStateWithDefaults

`func NewMicrosoftGraphCallMediaStateWithDefaults() *MicrosoftGraphCallMediaState`

NewMicrosoftGraphCallMediaStateWithDefaults instantiates a new MicrosoftGraphCallMediaState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAudio

`func (o *MicrosoftGraphCallMediaState) GetAudio() AnyOfmicrosoftGraphMediaState`

GetAudio returns the Audio field if non-nil, zero value otherwise.

### GetAudioOk

`func (o *MicrosoftGraphCallMediaState) GetAudioOk() (*AnyOfmicrosoftGraphMediaState, bool)`

GetAudioOk returns a tuple with the Audio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudio

`func (o *MicrosoftGraphCallMediaState) SetAudio(v AnyOfmicrosoftGraphMediaState)`

SetAudio sets Audio field to given value.

### HasAudio

`func (o *MicrosoftGraphCallMediaState) HasAudio() bool`

HasAudio returns a boolean if a field has been set.

### SetAudioNil

`func (o *MicrosoftGraphCallMediaState) SetAudioNil(b bool)`

 SetAudioNil sets the value for Audio to be an explicit nil

### UnsetAudio
`func (o *MicrosoftGraphCallMediaState) UnsetAudio()`

UnsetAudio ensures that no value is present for Audio, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


