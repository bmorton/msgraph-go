# MicrosoftGraphVideo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AudioBitsPerSample** | Pointer to **NullableInt32** | Number of audio bits per sample. | [optional] 
**AudioChannels** | Pointer to **NullableInt32** | Number of audio channels. | [optional] 
**AudioFormat** | Pointer to **NullableString** | Name of the audio format (AAC, MP3, etc.). | [optional] 
**AudioSamplesPerSecond** | Pointer to **NullableInt32** | Number of audio samples per second. | [optional] 
**Bitrate** | Pointer to **NullableInt32** | Bit rate of the video in bits per second. | [optional] 
**Duration** | Pointer to **NullableInt64** | Duration of the file in milliseconds. | [optional] 
**FourCC** | Pointer to **NullableString** | &#39;Four character code&#39; name of the video format. | [optional] 
**FrameRate** | Pointer to [**AnyOfnumberstringstring**](anyOf&lt;number,string,string&gt;.md) | Frame rate of the video. | [optional] 
**Height** | Pointer to **NullableInt32** | Height of the video, in pixels. | [optional] 
**Width** | Pointer to **NullableInt32** | Width of the video, in pixels. | [optional] 

## Methods

### NewMicrosoftGraphVideo

`func NewMicrosoftGraphVideo() *MicrosoftGraphVideo`

NewMicrosoftGraphVideo instantiates a new MicrosoftGraphVideo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphVideoWithDefaults

`func NewMicrosoftGraphVideoWithDefaults() *MicrosoftGraphVideo`

NewMicrosoftGraphVideoWithDefaults instantiates a new MicrosoftGraphVideo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAudioBitsPerSample

`func (o *MicrosoftGraphVideo) GetAudioBitsPerSample() int32`

GetAudioBitsPerSample returns the AudioBitsPerSample field if non-nil, zero value otherwise.

### GetAudioBitsPerSampleOk

`func (o *MicrosoftGraphVideo) GetAudioBitsPerSampleOk() (*int32, bool)`

GetAudioBitsPerSampleOk returns a tuple with the AudioBitsPerSample field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioBitsPerSample

`func (o *MicrosoftGraphVideo) SetAudioBitsPerSample(v int32)`

SetAudioBitsPerSample sets AudioBitsPerSample field to given value.

### HasAudioBitsPerSample

`func (o *MicrosoftGraphVideo) HasAudioBitsPerSample() bool`

HasAudioBitsPerSample returns a boolean if a field has been set.

### SetAudioBitsPerSampleNil

`func (o *MicrosoftGraphVideo) SetAudioBitsPerSampleNil(b bool)`

 SetAudioBitsPerSampleNil sets the value for AudioBitsPerSample to be an explicit nil

### UnsetAudioBitsPerSample
`func (o *MicrosoftGraphVideo) UnsetAudioBitsPerSample()`

UnsetAudioBitsPerSample ensures that no value is present for AudioBitsPerSample, not even an explicit nil
### GetAudioChannels

`func (o *MicrosoftGraphVideo) GetAudioChannels() int32`

GetAudioChannels returns the AudioChannels field if non-nil, zero value otherwise.

### GetAudioChannelsOk

`func (o *MicrosoftGraphVideo) GetAudioChannelsOk() (*int32, bool)`

GetAudioChannelsOk returns a tuple with the AudioChannels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioChannels

`func (o *MicrosoftGraphVideo) SetAudioChannels(v int32)`

SetAudioChannels sets AudioChannels field to given value.

### HasAudioChannels

`func (o *MicrosoftGraphVideo) HasAudioChannels() bool`

HasAudioChannels returns a boolean if a field has been set.

### SetAudioChannelsNil

`func (o *MicrosoftGraphVideo) SetAudioChannelsNil(b bool)`

 SetAudioChannelsNil sets the value for AudioChannels to be an explicit nil

### UnsetAudioChannels
`func (o *MicrosoftGraphVideo) UnsetAudioChannels()`

UnsetAudioChannels ensures that no value is present for AudioChannels, not even an explicit nil
### GetAudioFormat

`func (o *MicrosoftGraphVideo) GetAudioFormat() string`

GetAudioFormat returns the AudioFormat field if non-nil, zero value otherwise.

### GetAudioFormatOk

`func (o *MicrosoftGraphVideo) GetAudioFormatOk() (*string, bool)`

GetAudioFormatOk returns a tuple with the AudioFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioFormat

`func (o *MicrosoftGraphVideo) SetAudioFormat(v string)`

SetAudioFormat sets AudioFormat field to given value.

### HasAudioFormat

`func (o *MicrosoftGraphVideo) HasAudioFormat() bool`

HasAudioFormat returns a boolean if a field has been set.

### SetAudioFormatNil

`func (o *MicrosoftGraphVideo) SetAudioFormatNil(b bool)`

 SetAudioFormatNil sets the value for AudioFormat to be an explicit nil

### UnsetAudioFormat
`func (o *MicrosoftGraphVideo) UnsetAudioFormat()`

UnsetAudioFormat ensures that no value is present for AudioFormat, not even an explicit nil
### GetAudioSamplesPerSecond

`func (o *MicrosoftGraphVideo) GetAudioSamplesPerSecond() int32`

GetAudioSamplesPerSecond returns the AudioSamplesPerSecond field if non-nil, zero value otherwise.

### GetAudioSamplesPerSecondOk

`func (o *MicrosoftGraphVideo) GetAudioSamplesPerSecondOk() (*int32, bool)`

GetAudioSamplesPerSecondOk returns a tuple with the AudioSamplesPerSecond field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioSamplesPerSecond

`func (o *MicrosoftGraphVideo) SetAudioSamplesPerSecond(v int32)`

SetAudioSamplesPerSecond sets AudioSamplesPerSecond field to given value.

### HasAudioSamplesPerSecond

`func (o *MicrosoftGraphVideo) HasAudioSamplesPerSecond() bool`

HasAudioSamplesPerSecond returns a boolean if a field has been set.

### SetAudioSamplesPerSecondNil

`func (o *MicrosoftGraphVideo) SetAudioSamplesPerSecondNil(b bool)`

 SetAudioSamplesPerSecondNil sets the value for AudioSamplesPerSecond to be an explicit nil

### UnsetAudioSamplesPerSecond
`func (o *MicrosoftGraphVideo) UnsetAudioSamplesPerSecond()`

UnsetAudioSamplesPerSecond ensures that no value is present for AudioSamplesPerSecond, not even an explicit nil
### GetBitrate

`func (o *MicrosoftGraphVideo) GetBitrate() int32`

GetBitrate returns the Bitrate field if non-nil, zero value otherwise.

### GetBitrateOk

`func (o *MicrosoftGraphVideo) GetBitrateOk() (*int32, bool)`

GetBitrateOk returns a tuple with the Bitrate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBitrate

`func (o *MicrosoftGraphVideo) SetBitrate(v int32)`

SetBitrate sets Bitrate field to given value.

### HasBitrate

`func (o *MicrosoftGraphVideo) HasBitrate() bool`

HasBitrate returns a boolean if a field has been set.

### SetBitrateNil

`func (o *MicrosoftGraphVideo) SetBitrateNil(b bool)`

 SetBitrateNil sets the value for Bitrate to be an explicit nil

### UnsetBitrate
`func (o *MicrosoftGraphVideo) UnsetBitrate()`

UnsetBitrate ensures that no value is present for Bitrate, not even an explicit nil
### GetDuration

`func (o *MicrosoftGraphVideo) GetDuration() int64`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *MicrosoftGraphVideo) GetDurationOk() (*int64, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *MicrosoftGraphVideo) SetDuration(v int64)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *MicrosoftGraphVideo) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### SetDurationNil

`func (o *MicrosoftGraphVideo) SetDurationNil(b bool)`

 SetDurationNil sets the value for Duration to be an explicit nil

### UnsetDuration
`func (o *MicrosoftGraphVideo) UnsetDuration()`

UnsetDuration ensures that no value is present for Duration, not even an explicit nil
### GetFourCC

`func (o *MicrosoftGraphVideo) GetFourCC() string`

GetFourCC returns the FourCC field if non-nil, zero value otherwise.

### GetFourCCOk

`func (o *MicrosoftGraphVideo) GetFourCCOk() (*string, bool)`

GetFourCCOk returns a tuple with the FourCC field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFourCC

`func (o *MicrosoftGraphVideo) SetFourCC(v string)`

SetFourCC sets FourCC field to given value.

### HasFourCC

`func (o *MicrosoftGraphVideo) HasFourCC() bool`

HasFourCC returns a boolean if a field has been set.

### SetFourCCNil

`func (o *MicrosoftGraphVideo) SetFourCCNil(b bool)`

 SetFourCCNil sets the value for FourCC to be an explicit nil

### UnsetFourCC
`func (o *MicrosoftGraphVideo) UnsetFourCC()`

UnsetFourCC ensures that no value is present for FourCC, not even an explicit nil
### GetFrameRate

`func (o *MicrosoftGraphVideo) GetFrameRate() AnyOfnumberstringstring`

GetFrameRate returns the FrameRate field if non-nil, zero value otherwise.

### GetFrameRateOk

`func (o *MicrosoftGraphVideo) GetFrameRateOk() (*AnyOfnumberstringstring, bool)`

GetFrameRateOk returns a tuple with the FrameRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrameRate

`func (o *MicrosoftGraphVideo) SetFrameRate(v AnyOfnumberstringstring)`

SetFrameRate sets FrameRate field to given value.

### HasFrameRate

`func (o *MicrosoftGraphVideo) HasFrameRate() bool`

HasFrameRate returns a boolean if a field has been set.

### SetFrameRateNil

`func (o *MicrosoftGraphVideo) SetFrameRateNil(b bool)`

 SetFrameRateNil sets the value for FrameRate to be an explicit nil

### UnsetFrameRate
`func (o *MicrosoftGraphVideo) UnsetFrameRate()`

UnsetFrameRate ensures that no value is present for FrameRate, not even an explicit nil
### GetHeight

`func (o *MicrosoftGraphVideo) GetHeight() int32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *MicrosoftGraphVideo) GetHeightOk() (*int32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *MicrosoftGraphVideo) SetHeight(v int32)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *MicrosoftGraphVideo) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### SetHeightNil

`func (o *MicrosoftGraphVideo) SetHeightNil(b bool)`

 SetHeightNil sets the value for Height to be an explicit nil

### UnsetHeight
`func (o *MicrosoftGraphVideo) UnsetHeight()`

UnsetHeight ensures that no value is present for Height, not even an explicit nil
### GetWidth

`func (o *MicrosoftGraphVideo) GetWidth() int32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *MicrosoftGraphVideo) GetWidthOk() (*int32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *MicrosoftGraphVideo) SetWidth(v int32)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *MicrosoftGraphVideo) HasWidth() bool`

HasWidth returns a boolean if a field has been set.

### SetWidthNil

`func (o *MicrosoftGraphVideo) SetWidthNil(b bool)`

 SetWidthNil sets the value for Width to be an explicit nil

### UnsetWidth
`func (o *MicrosoftGraphVideo) UnsetWidth()`

UnsetWidth ensures that no value is present for Width, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


