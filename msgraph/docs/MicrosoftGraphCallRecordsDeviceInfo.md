# MicrosoftGraphCallRecordsDeviceInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CaptureDeviceDriver** | Pointer to **NullableString** | Name of the capture device driver used by the media endpoint. | [optional] 
**CaptureDeviceName** | Pointer to **NullableString** | Name of the capture device used by the media endpoint. | [optional] 
**CaptureNotFunctioningEventRatio** | Pointer to [**AnyOfnumberstringstring**](anyOf&lt;number,string,string&gt;.md) | Fraction of the call that the media endpoint detected the capture device was not working properly. | [optional] 
**CpuInsufficentEventRatio** | Pointer to [**AnyOfnumberstringstring**](anyOf&lt;number,string,string&gt;.md) | Fraction of the call that the media endpoint detected the CPU resources available were insufficient and caused poor quality of the audio sent and received. | [optional] 
**DeviceClippingEventRatio** | Pointer to [**AnyOfnumberstringstring**](anyOf&lt;number,string,string&gt;.md) | Fraction of the call that the media endpoint detected clipping in the captured audio that caused poor quality of the audio being sent. | [optional] 
**DeviceGlitchEventRatio** | Pointer to [**AnyOfnumberstringstring**](anyOf&lt;number,string,string&gt;.md) | Fraction of the call that the media endpoint detected glitches or gaps in the audio played or captured that caused poor quality of the audio being sent or received. | [optional] 
**HowlingEventCount** | Pointer to **NullableInt32** | Number of times during the call that the media endpoint detected howling or screeching audio. | [optional] 
**InitialSignalLevelRootMeanSquare** | Pointer to [**AnyOfnumberstringstring**](anyOf&lt;number,string,string&gt;.md) | The root mean square (RMS) of the incoming signal of up to the first 30 seconds of the call. | [optional] 
**LowSpeechLevelEventRatio** | Pointer to [**AnyOfnumberstringstring**](anyOf&lt;number,string,string&gt;.md) | Fraction of the call that the media endpoint detected low speech level that caused poor quality of the audio being sent. | [optional] 
**LowSpeechToNoiseEventRatio** | Pointer to [**AnyOfnumberstringstring**](anyOf&lt;number,string,string&gt;.md) | Fraction of the call that the media endpoint detected low speech to noise level that caused poor quality of the audio being sent. | [optional] 
**MicGlitchRate** | Pointer to [**AnyOfnumberstringstring**](anyOf&lt;number,string,string&gt;.md) | Glitches per 5 minute interval for the media endpoint&#39;s microphone. | [optional] 
**ReceivedNoiseLevel** | Pointer to **NullableInt32** | Average energy level of received audio for audio classified as mono noise or left channel of stereo noise by the media endpoint. | [optional] 
**ReceivedSignalLevel** | Pointer to **NullableInt32** | Average energy level of received audio for audio classified as mono speech, or left channel of stereo speech by the media endpoint. | [optional] 
**RenderDeviceDriver** | Pointer to **NullableString** | Name of the render device driver used by the media endpoint. | [optional] 
**RenderDeviceName** | Pointer to **NullableString** | Name of the render device used by the media endpoint. | [optional] 
**RenderMuteEventRatio** | Pointer to [**AnyOfnumberstringstring**](anyOf&lt;number,string,string&gt;.md) | Fraction of the call that media endpoint detected device render is muted. | [optional] 
**RenderNotFunctioningEventRatio** | Pointer to [**AnyOfnumberstringstring**](anyOf&lt;number,string,string&gt;.md) | Fraction of the call that the media endpoint detected the render device was not working properly. | [optional] 
**RenderZeroVolumeEventRatio** | Pointer to [**AnyOfnumberstringstring**](anyOf&lt;number,string,string&gt;.md) | Fraction of the call that media endpoint detected device render volume is set to 0. | [optional] 
**SentNoiseLevel** | Pointer to **NullableInt32** | Average energy level of sent audio for audio classified as mono noise or left channel of stereo noise by the media endpoint. | [optional] 
**SentSignalLevel** | Pointer to **NullableInt32** | Average energy level of sent audio for audio classified as mono speech, or left channel of stereo speech by the media endpoint. | [optional] 
**SpeakerGlitchRate** | Pointer to [**AnyOfnumberstringstring**](anyOf&lt;number,string,string&gt;.md) | Glitches per 5 minute internal for the media endpoint&#39;s loudspeaker. | [optional] 

## Methods

### NewMicrosoftGraphCallRecordsDeviceInfo

`func NewMicrosoftGraphCallRecordsDeviceInfo() *MicrosoftGraphCallRecordsDeviceInfo`

NewMicrosoftGraphCallRecordsDeviceInfo instantiates a new MicrosoftGraphCallRecordsDeviceInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphCallRecordsDeviceInfoWithDefaults

`func NewMicrosoftGraphCallRecordsDeviceInfoWithDefaults() *MicrosoftGraphCallRecordsDeviceInfo`

NewMicrosoftGraphCallRecordsDeviceInfoWithDefaults instantiates a new MicrosoftGraphCallRecordsDeviceInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCaptureDeviceDriver

`func (o *MicrosoftGraphCallRecordsDeviceInfo) GetCaptureDeviceDriver() string`

GetCaptureDeviceDriver returns the CaptureDeviceDriver field if non-nil, zero value otherwise.

### GetCaptureDeviceDriverOk

`func (o *MicrosoftGraphCallRecordsDeviceInfo) GetCaptureDeviceDriverOk() (*string, bool)`

GetCaptureDeviceDriverOk returns a tuple with the CaptureDeviceDriver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptureDeviceDriver

`func (o *MicrosoftGraphCallRecordsDeviceInfo) SetCaptureDeviceDriver(v string)`

SetCaptureDeviceDriver sets CaptureDeviceDriver field to given value.

### HasCaptureDeviceDriver

`func (o *MicrosoftGraphCallRecordsDeviceInfo) HasCaptureDeviceDriver() bool`

HasCaptureDeviceDriver returns a boolean if a field has been set.

### SetCaptureDeviceDriverNil

`func (o *MicrosoftGraphCallRecordsDeviceInfo) SetCaptureDeviceDriverNil(b bool)`

 SetCaptureDeviceDriverNil sets the value for CaptureDeviceDriver to be an explicit nil

### UnsetCaptureDeviceDriver
`func (o *MicrosoftGraphCallRecordsDeviceInfo) UnsetCaptureDeviceDriver()`

UnsetCaptureDeviceDriver ensures that no value is present for CaptureDeviceDriver, not even an explicit nil
### GetCaptureDeviceName

`func (o *MicrosoftGraphCallRecordsDeviceInfo) GetCaptureDeviceName() string`

GetCaptureDeviceName returns the CaptureDeviceName field if non-nil, zero value otherwise.

### GetCaptureDeviceNameOk

`func (o *MicrosoftGraphCallRecordsDeviceInfo) GetCaptureDeviceNameOk() (*string, bool)`

GetCaptureDeviceNameOk returns a tuple with the CaptureDeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptureDeviceName

`func (o *MicrosoftGraphCallRecordsDeviceInfo) SetCaptureDeviceName(v string)`

SetCaptureDeviceName sets CaptureDeviceName field to given value.

### HasCaptureDeviceName

`func (o *MicrosoftGraphCallRecordsDeviceInfo) HasCaptureDeviceName() bool`

HasCaptureDeviceName returns a boolean if a field has been set.

### SetCaptureDeviceNameNil

`func (o *MicrosoftGraphCallRecordsDeviceInfo) SetCaptureDeviceNameNil(b bool)`

 SetCaptureDeviceNameNil sets the value for CaptureDeviceName to be an explicit nil

### UnsetCaptureDeviceName
`func (o *MicrosoftGraphCallRecordsDeviceInfo) UnsetCaptureDeviceName()`

UnsetCaptureDeviceName ensures that no value is present for CaptureDeviceName, not even an explicit nil
### GetCaptureNotFunctioningEventRatio

`func (o *MicrosoftGraphCallRecordsDeviceInfo) GetCaptureNotFunctioningEventRatio() AnyOfnumberstringstring`

GetCaptureNotFunctioningEventRatio returns the CaptureNotFunctioningEventRatio field if non-nil, zero value otherwise.

### GetCaptureNotFunctioningEventRatioOk

`func (o *MicrosoftGraphCallRecordsDeviceInfo) GetCaptureNotFunctioningEventRatioOk() (*AnyOfnumberstringstring, bool)`

GetCaptureNotFunctioningEventRatioOk returns a tuple with the CaptureNotFunctioningEventRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptureNotFunctioningEventRatio

`func (o *MicrosoftGraphCallRecordsDeviceInfo) SetCaptureNotFunctioningEventRatio(v AnyOfnumberstringstring)`

SetCaptureNotFunctioningEventRatio sets CaptureNotFunctioningEventRatio field to given value.

### HasCaptureNotFunctioningEventRatio

`func (o *MicrosoftGraphCallRecordsDeviceInfo) HasCaptureNotFunctioningEventRatio() bool`

HasCaptureNotFunctioningEventRatio returns a boolean if a field has been set.

### SetCaptureNotFunctioningEventRatioNil

`func (o *MicrosoftGraphCallRecordsDeviceInfo) SetCaptureNotFunctioningEventRatioNil(b bool)`

 SetCaptureNotFunctioningEventRatioNil sets the value for CaptureNotFunctioningEventRatio to be an explicit nil

### UnsetCaptureNotFunctioningEventRatio
`func (o *MicrosoftGraphCallRecordsDeviceInfo) UnsetCaptureNotFunctioningEventRatio()`

UnsetCaptureNotFunctioningEventRatio ensures that no value is present for CaptureNotFunctioningEventRatio, not even an explicit nil
### GetCpuInsufficentEventRatio

`func (o *MicrosoftGraphCallRecordsDeviceInfo) GetCpuInsufficentEventRatio() AnyOfnumberstringstring`

GetCpuInsufficentEventRatio returns the CpuInsufficentEventRatio field if non-nil, zero value otherwise.

### GetCpuInsufficentEventRatioOk

`func (o *MicrosoftGraphCallRecordsDeviceInfo) GetCpuInsufficentEventRatioOk() (*AnyOfnumberstringstring, bool)`

GetCpuInsufficentEventRatioOk returns a tuple with the CpuInsufficentEventRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuInsufficentEventRatio

`func (o *MicrosoftGraphCallRecordsDeviceInfo) SetCpuInsufficentEventRatio(v AnyOfnumberstringstring)`

SetCpuInsufficentEventRatio sets CpuInsufficentEventRatio field to given value.

### HasCpuInsufficentEventRatio

`func (o *MicrosoftGraphCallRecordsDeviceInfo) HasCpuInsufficentEventRatio() bool`

HasCpuInsufficentEventRatio returns a boolean if a field has been set.

### SetCpuInsufficentEventRatioNil

`func (o *MicrosoftGraphCallRecordsDeviceInfo) SetCpuInsufficentEventRatioNil(b bool)`

 SetCpuInsufficentEventRatioNil sets the value for CpuInsufficentEventRatio to be an explicit nil

### UnsetCpuInsufficentEventRatio
`func (o *MicrosoftGraphCallRecordsDeviceInfo) UnsetCpuInsufficentEventRatio()`

UnsetCpuInsufficentEventRatio ensures that no value is present for CpuInsufficentEventRatio, not even an explicit nil
### GetDeviceClippingEventRatio

`func (o *MicrosoftGraphCallRecordsDeviceInfo) GetDeviceClippingEventRatio() AnyOfnumberstringstring`

GetDeviceClippingEventRatio returns the DeviceClippingEventRatio field if non-nil, zero value otherwise.

### GetDeviceClippingEventRatioOk

`func (o *MicrosoftGraphCallRecordsDeviceInfo) GetDeviceClippingEventRatioOk() (*AnyOfnumberstringstring, bool)`

GetDeviceClippingEventRatioOk returns a tuple with the DeviceClippingEventRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceClippingEventRatio

`func (o *MicrosoftGraphCallRecordsDeviceInfo) SetDeviceClippingEventRatio(v AnyOfnumberstringstring)`

SetDeviceClippingEventRatio sets DeviceClippingEventRatio field to given value.

### HasDeviceClippingEventRatio

`func (o *MicrosoftGraphCallRecordsDeviceInfo) HasDeviceClippingEventRatio() bool`

HasDeviceClippingEventRatio returns a boolean if a field has been set.

### SetDeviceClippingEventRatioNil

`func (o *MicrosoftGraphCallRecordsDeviceInfo) SetDeviceClippingEventRatioNil(b bool)`

 SetDeviceClippingEventRatioNil sets the value for DeviceClippingEventRatio to be an explicit nil

### UnsetDeviceClippingEventRatio
`func (o *MicrosoftGraphCallRecordsDeviceInfo) UnsetDeviceClippingEventRatio()`

UnsetDeviceClippingEventRatio ensures that no value is present for DeviceClippingEventRatio, not even an explicit nil
### GetDeviceGlitchEventRatio

`func (o *MicrosoftGraphCallRecordsDeviceInfo) GetDeviceGlitchEventRatio() AnyOfnumberstringstring`

GetDeviceGlitchEventRatio returns the DeviceGlitchEventRatio field if non-nil, zero value otherwise.

### GetDeviceGlitchEventRatioOk

`func (o *MicrosoftGraphCallRecordsDeviceInfo) GetDeviceGlitchEventRatioOk() (*AnyOfnumberstringstring, bool)`

GetDeviceGlitchEventRatioOk returns a tuple with the DeviceGlitchEventRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceGlitchEventRatio

`func (o *MicrosoftGraphCallRecordsDeviceInfo) SetDeviceGlitchEventRatio(v AnyOfnumberstringstring)`

SetDeviceGlitchEventRatio sets DeviceGlitchEventRatio field to given value.

### HasDeviceGlitchEventRatio

`func (o *MicrosoftGraphCallRecordsDeviceInfo) HasDeviceGlitchEventRatio() bool`

HasDeviceGlitchEventRatio returns a boolean if a field has been set.

### SetDeviceGlitchEventRatioNil

`func (o *MicrosoftGraphCallRecordsDeviceInfo) SetDeviceGlitchEventRatioNil(b bool)`

 SetDeviceGlitchEventRatioNil sets the value for DeviceGlitchEventRatio to be an explicit nil

### UnsetDeviceGlitchEventRatio
`func (o *MicrosoftGraphCallRecordsDeviceInfo) UnsetDeviceGlitchEventRatio()`

UnsetDeviceGlitchEventRatio ensures that no value is present for DeviceGlitchEventRatio, not even an explicit nil
### GetHowlingEventCount

`func (o *MicrosoftGraphCallRecordsDeviceInfo) GetHowlingEventCount() int32`

GetHowlingEventCount returns the HowlingEventCount field if non-nil, zero value otherwise.

### GetHowlingEventCountOk

`func (o *MicrosoftGraphCallRecordsDeviceInfo) GetHowlingEventCountOk() (*int32, bool)`

GetHowlingEventCountOk returns a tuple with the HowlingEventCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHowlingEventCount

`func (o *MicrosoftGraphCallRecordsDeviceInfo) SetHowlingEventCount(v int32)`

SetHowlingEventCount sets HowlingEventCount field to given value.

### HasHowlingEventCount

`func (o *MicrosoftGraphCallRecordsDeviceInfo) HasHowlingEventCount() bool`

HasHowlingEventCount returns a boolean if a field has been set.

### SetHowlingEventCountNil

`func (o *MicrosoftGraphCallRecordsDeviceInfo) SetHowlingEventCountNil(b bool)`

 SetHowlingEventCountNil sets the value for HowlingEventCount to be an explicit nil

### UnsetHowlingEventCount
`func (o *MicrosoftGraphCallRecordsDeviceInfo) UnsetHowlingEventCount()`

UnsetHowlingEventCount ensures that no value is present for HowlingEventCount, not even an explicit nil
### GetInitialSignalLevelRootMeanSquare

`func (o *MicrosoftGraphCallRecordsDeviceInfo) GetInitialSignalLevelRootMeanSquare() AnyOfnumberstringstring`

GetInitialSignalLevelRootMeanSquare returns the InitialSignalLevelRootMeanSquare field if non-nil, zero value otherwise.

### GetInitialSignalLevelRootMeanSquareOk

`func (o *MicrosoftGraphCallRecordsDeviceInfo) GetInitialSignalLevelRootMeanSquareOk() (*AnyOfnumberstringstring, bool)`

GetInitialSignalLevelRootMeanSquareOk returns a tuple with the InitialSignalLevelRootMeanSquare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialSignalLevelRootMeanSquare

`func (o *MicrosoftGraphCallRecordsDeviceInfo) SetInitialSignalLevelRootMeanSquare(v AnyOfnumberstringstring)`

SetInitialSignalLevelRootMeanSquare sets InitialSignalLevelRootMeanSquare field to given value.

### HasInitialSignalLevelRootMeanSquare

`func (o *MicrosoftGraphCallRecordsDeviceInfo) HasInitialSignalLevelRootMeanSquare() bool`

HasInitialSignalLevelRootMeanSquare returns a boolean if a field has been set.

### SetInitialSignalLevelRootMeanSquareNil

`func (o *MicrosoftGraphCallRecordsDeviceInfo) SetInitialSignalLevelRootMeanSquareNil(b bool)`

 SetInitialSignalLevelRootMeanSquareNil sets the value for InitialSignalLevelRootMeanSquare to be an explicit nil

### UnsetInitialSignalLevelRootMeanSquare
`func (o *MicrosoftGraphCallRecordsDeviceInfo) UnsetInitialSignalLevelRootMeanSquare()`

UnsetInitialSignalLevelRootMeanSquare ensures that no value is present for InitialSignalLevelRootMeanSquare, not even an explicit nil
### GetLowSpeechLevelEventRatio

`func (o *MicrosoftGraphCallRecordsDeviceInfo) GetLowSpeechLevelEventRatio() AnyOfnumberstringstring`

GetLowSpeechLevelEventRatio returns the LowSpeechLevelEventRatio field if non-nil, zero value otherwise.

### GetLowSpeechLevelEventRatioOk

`func (o *MicrosoftGraphCallRecordsDeviceInfo) GetLowSpeechLevelEventRatioOk() (*AnyOfnumberstringstring, bool)`

GetLowSpeechLevelEventRatioOk returns a tuple with the LowSpeechLevelEventRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowSpeechLevelEventRatio

`func (o *MicrosoftGraphCallRecordsDeviceInfo) SetLowSpeechLevelEventRatio(v AnyOfnumberstringstring)`

SetLowSpeechLevelEventRatio sets LowSpeechLevelEventRatio field to given value.

### HasLowSpeechLevelEventRatio

`func (o *MicrosoftGraphCallRecordsDeviceInfo) HasLowSpeechLevelEventRatio() bool`

HasLowSpeechLevelEventRatio returns a boolean if a field has been set.

### SetLowSpeechLevelEventRatioNil

`func (o *MicrosoftGraphCallRecordsDeviceInfo) SetLowSpeechLevelEventRatioNil(b bool)`

 SetLowSpeechLevelEventRatioNil sets the value for LowSpeechLevelEventRatio to be an explicit nil

### UnsetLowSpeechLevelEventRatio
`func (o *MicrosoftGraphCallRecordsDeviceInfo) UnsetLowSpeechLevelEventRatio()`

UnsetLowSpeechLevelEventRatio ensures that no value is present for LowSpeechLevelEventRatio, not even an explicit nil
### GetLowSpeechToNoiseEventRatio

`func (o *MicrosoftGraphCallRecordsDeviceInfo) GetLowSpeechToNoiseEventRatio() AnyOfnumberstringstring`

GetLowSpeechToNoiseEventRatio returns the LowSpeechToNoiseEventRatio field if non-nil, zero value otherwise.

### GetLowSpeechToNoiseEventRatioOk

`func (o *MicrosoftGraphCallRecordsDeviceInfo) GetLowSpeechToNoiseEventRatioOk() (*AnyOfnumberstringstring, bool)`

GetLowSpeechToNoiseEventRatioOk returns a tuple with the LowSpeechToNoiseEventRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowSpeechToNoiseEventRatio

`func (o *MicrosoftGraphCallRecordsDeviceInfo) SetLowSpeechToNoiseEventRatio(v AnyOfnumberstringstring)`

SetLowSpeechToNoiseEventRatio sets LowSpeechToNoiseEventRatio field to given value.

### HasLowSpeechToNoiseEventRatio

`func (o *MicrosoftGraphCallRecordsDeviceInfo) HasLowSpeechToNoiseEventRatio() bool`

HasLowSpeechToNoiseEventRatio returns a boolean if a field has been set.

### SetLowSpeechToNoiseEventRatioNil

`func (o *MicrosoftGraphCallRecordsDeviceInfo) SetLowSpeechToNoiseEventRatioNil(b bool)`

 SetLowSpeechToNoiseEventRatioNil sets the value for LowSpeechToNoiseEventRatio to be an explicit nil

### UnsetLowSpeechToNoiseEventRatio
`func (o *MicrosoftGraphCallRecordsDeviceInfo) UnsetLowSpeechToNoiseEventRatio()`

UnsetLowSpeechToNoiseEventRatio ensures that no value is present for LowSpeechToNoiseEventRatio, not even an explicit nil
### GetMicGlitchRate

`func (o *MicrosoftGraphCallRecordsDeviceInfo) GetMicGlitchRate() AnyOfnumberstringstring`

GetMicGlitchRate returns the MicGlitchRate field if non-nil, zero value otherwise.

### GetMicGlitchRateOk

`func (o *MicrosoftGraphCallRecordsDeviceInfo) GetMicGlitchRateOk() (*AnyOfnumberstringstring, bool)`

GetMicGlitchRateOk returns a tuple with the MicGlitchRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicGlitchRate

`func (o *MicrosoftGraphCallRecordsDeviceInfo) SetMicGlitchRate(v AnyOfnumberstringstring)`

SetMicGlitchRate sets MicGlitchRate field to given value.

### HasMicGlitchRate

`func (o *MicrosoftGraphCallRecordsDeviceInfo) HasMicGlitchRate() bool`

HasMicGlitchRate returns a boolean if a field has been set.

### SetMicGlitchRateNil

`func (o *MicrosoftGraphCallRecordsDeviceInfo) SetMicGlitchRateNil(b bool)`

 SetMicGlitchRateNil sets the value for MicGlitchRate to be an explicit nil

### UnsetMicGlitchRate
`func (o *MicrosoftGraphCallRecordsDeviceInfo) UnsetMicGlitchRate()`

UnsetMicGlitchRate ensures that no value is present for MicGlitchRate, not even an explicit nil
### GetReceivedNoiseLevel

`func (o *MicrosoftGraphCallRecordsDeviceInfo) GetReceivedNoiseLevel() int32`

GetReceivedNoiseLevel returns the ReceivedNoiseLevel field if non-nil, zero value otherwise.

### GetReceivedNoiseLevelOk

`func (o *MicrosoftGraphCallRecordsDeviceInfo) GetReceivedNoiseLevelOk() (*int32, bool)`

GetReceivedNoiseLevelOk returns a tuple with the ReceivedNoiseLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivedNoiseLevel

`func (o *MicrosoftGraphCallRecordsDeviceInfo) SetReceivedNoiseLevel(v int32)`

SetReceivedNoiseLevel sets ReceivedNoiseLevel field to given value.

### HasReceivedNoiseLevel

`func (o *MicrosoftGraphCallRecordsDeviceInfo) HasReceivedNoiseLevel() bool`

HasReceivedNoiseLevel returns a boolean if a field has been set.

### SetReceivedNoiseLevelNil

`func (o *MicrosoftGraphCallRecordsDeviceInfo) SetReceivedNoiseLevelNil(b bool)`

 SetReceivedNoiseLevelNil sets the value for ReceivedNoiseLevel to be an explicit nil

### UnsetReceivedNoiseLevel
`func (o *MicrosoftGraphCallRecordsDeviceInfo) UnsetReceivedNoiseLevel()`

UnsetReceivedNoiseLevel ensures that no value is present for ReceivedNoiseLevel, not even an explicit nil
### GetReceivedSignalLevel

`func (o *MicrosoftGraphCallRecordsDeviceInfo) GetReceivedSignalLevel() int32`

GetReceivedSignalLevel returns the ReceivedSignalLevel field if non-nil, zero value otherwise.

### GetReceivedSignalLevelOk

`func (o *MicrosoftGraphCallRecordsDeviceInfo) GetReceivedSignalLevelOk() (*int32, bool)`

GetReceivedSignalLevelOk returns a tuple with the ReceivedSignalLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivedSignalLevel

`func (o *MicrosoftGraphCallRecordsDeviceInfo) SetReceivedSignalLevel(v int32)`

SetReceivedSignalLevel sets ReceivedSignalLevel field to given value.

### HasReceivedSignalLevel

`func (o *MicrosoftGraphCallRecordsDeviceInfo) HasReceivedSignalLevel() bool`

HasReceivedSignalLevel returns a boolean if a field has been set.

### SetReceivedSignalLevelNil

`func (o *MicrosoftGraphCallRecordsDeviceInfo) SetReceivedSignalLevelNil(b bool)`

 SetReceivedSignalLevelNil sets the value for ReceivedSignalLevel to be an explicit nil

### UnsetReceivedSignalLevel
`func (o *MicrosoftGraphCallRecordsDeviceInfo) UnsetReceivedSignalLevel()`

UnsetReceivedSignalLevel ensures that no value is present for ReceivedSignalLevel, not even an explicit nil
### GetRenderDeviceDriver

`func (o *MicrosoftGraphCallRecordsDeviceInfo) GetRenderDeviceDriver() string`

GetRenderDeviceDriver returns the RenderDeviceDriver field if non-nil, zero value otherwise.

### GetRenderDeviceDriverOk

`func (o *MicrosoftGraphCallRecordsDeviceInfo) GetRenderDeviceDriverOk() (*string, bool)`

GetRenderDeviceDriverOk returns a tuple with the RenderDeviceDriver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenderDeviceDriver

`func (o *MicrosoftGraphCallRecordsDeviceInfo) SetRenderDeviceDriver(v string)`

SetRenderDeviceDriver sets RenderDeviceDriver field to given value.

### HasRenderDeviceDriver

`func (o *MicrosoftGraphCallRecordsDeviceInfo) HasRenderDeviceDriver() bool`

HasRenderDeviceDriver returns a boolean if a field has been set.

### SetRenderDeviceDriverNil

`func (o *MicrosoftGraphCallRecordsDeviceInfo) SetRenderDeviceDriverNil(b bool)`

 SetRenderDeviceDriverNil sets the value for RenderDeviceDriver to be an explicit nil

### UnsetRenderDeviceDriver
`func (o *MicrosoftGraphCallRecordsDeviceInfo) UnsetRenderDeviceDriver()`

UnsetRenderDeviceDriver ensures that no value is present for RenderDeviceDriver, not even an explicit nil
### GetRenderDeviceName

`func (o *MicrosoftGraphCallRecordsDeviceInfo) GetRenderDeviceName() string`

GetRenderDeviceName returns the RenderDeviceName field if non-nil, zero value otherwise.

### GetRenderDeviceNameOk

`func (o *MicrosoftGraphCallRecordsDeviceInfo) GetRenderDeviceNameOk() (*string, bool)`

GetRenderDeviceNameOk returns a tuple with the RenderDeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenderDeviceName

`func (o *MicrosoftGraphCallRecordsDeviceInfo) SetRenderDeviceName(v string)`

SetRenderDeviceName sets RenderDeviceName field to given value.

### HasRenderDeviceName

`func (o *MicrosoftGraphCallRecordsDeviceInfo) HasRenderDeviceName() bool`

HasRenderDeviceName returns a boolean if a field has been set.

### SetRenderDeviceNameNil

`func (o *MicrosoftGraphCallRecordsDeviceInfo) SetRenderDeviceNameNil(b bool)`

 SetRenderDeviceNameNil sets the value for RenderDeviceName to be an explicit nil

### UnsetRenderDeviceName
`func (o *MicrosoftGraphCallRecordsDeviceInfo) UnsetRenderDeviceName()`

UnsetRenderDeviceName ensures that no value is present for RenderDeviceName, not even an explicit nil
### GetRenderMuteEventRatio

`func (o *MicrosoftGraphCallRecordsDeviceInfo) GetRenderMuteEventRatio() AnyOfnumberstringstring`

GetRenderMuteEventRatio returns the RenderMuteEventRatio field if non-nil, zero value otherwise.

### GetRenderMuteEventRatioOk

`func (o *MicrosoftGraphCallRecordsDeviceInfo) GetRenderMuteEventRatioOk() (*AnyOfnumberstringstring, bool)`

GetRenderMuteEventRatioOk returns a tuple with the RenderMuteEventRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenderMuteEventRatio

`func (o *MicrosoftGraphCallRecordsDeviceInfo) SetRenderMuteEventRatio(v AnyOfnumberstringstring)`

SetRenderMuteEventRatio sets RenderMuteEventRatio field to given value.

### HasRenderMuteEventRatio

`func (o *MicrosoftGraphCallRecordsDeviceInfo) HasRenderMuteEventRatio() bool`

HasRenderMuteEventRatio returns a boolean if a field has been set.

### SetRenderMuteEventRatioNil

`func (o *MicrosoftGraphCallRecordsDeviceInfo) SetRenderMuteEventRatioNil(b bool)`

 SetRenderMuteEventRatioNil sets the value for RenderMuteEventRatio to be an explicit nil

### UnsetRenderMuteEventRatio
`func (o *MicrosoftGraphCallRecordsDeviceInfo) UnsetRenderMuteEventRatio()`

UnsetRenderMuteEventRatio ensures that no value is present for RenderMuteEventRatio, not even an explicit nil
### GetRenderNotFunctioningEventRatio

`func (o *MicrosoftGraphCallRecordsDeviceInfo) GetRenderNotFunctioningEventRatio() AnyOfnumberstringstring`

GetRenderNotFunctioningEventRatio returns the RenderNotFunctioningEventRatio field if non-nil, zero value otherwise.

### GetRenderNotFunctioningEventRatioOk

`func (o *MicrosoftGraphCallRecordsDeviceInfo) GetRenderNotFunctioningEventRatioOk() (*AnyOfnumberstringstring, bool)`

GetRenderNotFunctioningEventRatioOk returns a tuple with the RenderNotFunctioningEventRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenderNotFunctioningEventRatio

`func (o *MicrosoftGraphCallRecordsDeviceInfo) SetRenderNotFunctioningEventRatio(v AnyOfnumberstringstring)`

SetRenderNotFunctioningEventRatio sets RenderNotFunctioningEventRatio field to given value.

### HasRenderNotFunctioningEventRatio

`func (o *MicrosoftGraphCallRecordsDeviceInfo) HasRenderNotFunctioningEventRatio() bool`

HasRenderNotFunctioningEventRatio returns a boolean if a field has been set.

### SetRenderNotFunctioningEventRatioNil

`func (o *MicrosoftGraphCallRecordsDeviceInfo) SetRenderNotFunctioningEventRatioNil(b bool)`

 SetRenderNotFunctioningEventRatioNil sets the value for RenderNotFunctioningEventRatio to be an explicit nil

### UnsetRenderNotFunctioningEventRatio
`func (o *MicrosoftGraphCallRecordsDeviceInfo) UnsetRenderNotFunctioningEventRatio()`

UnsetRenderNotFunctioningEventRatio ensures that no value is present for RenderNotFunctioningEventRatio, not even an explicit nil
### GetRenderZeroVolumeEventRatio

`func (o *MicrosoftGraphCallRecordsDeviceInfo) GetRenderZeroVolumeEventRatio() AnyOfnumberstringstring`

GetRenderZeroVolumeEventRatio returns the RenderZeroVolumeEventRatio field if non-nil, zero value otherwise.

### GetRenderZeroVolumeEventRatioOk

`func (o *MicrosoftGraphCallRecordsDeviceInfo) GetRenderZeroVolumeEventRatioOk() (*AnyOfnumberstringstring, bool)`

GetRenderZeroVolumeEventRatioOk returns a tuple with the RenderZeroVolumeEventRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenderZeroVolumeEventRatio

`func (o *MicrosoftGraphCallRecordsDeviceInfo) SetRenderZeroVolumeEventRatio(v AnyOfnumberstringstring)`

SetRenderZeroVolumeEventRatio sets RenderZeroVolumeEventRatio field to given value.

### HasRenderZeroVolumeEventRatio

`func (o *MicrosoftGraphCallRecordsDeviceInfo) HasRenderZeroVolumeEventRatio() bool`

HasRenderZeroVolumeEventRatio returns a boolean if a field has been set.

### SetRenderZeroVolumeEventRatioNil

`func (o *MicrosoftGraphCallRecordsDeviceInfo) SetRenderZeroVolumeEventRatioNil(b bool)`

 SetRenderZeroVolumeEventRatioNil sets the value for RenderZeroVolumeEventRatio to be an explicit nil

### UnsetRenderZeroVolumeEventRatio
`func (o *MicrosoftGraphCallRecordsDeviceInfo) UnsetRenderZeroVolumeEventRatio()`

UnsetRenderZeroVolumeEventRatio ensures that no value is present for RenderZeroVolumeEventRatio, not even an explicit nil
### GetSentNoiseLevel

`func (o *MicrosoftGraphCallRecordsDeviceInfo) GetSentNoiseLevel() int32`

GetSentNoiseLevel returns the SentNoiseLevel field if non-nil, zero value otherwise.

### GetSentNoiseLevelOk

`func (o *MicrosoftGraphCallRecordsDeviceInfo) GetSentNoiseLevelOk() (*int32, bool)`

GetSentNoiseLevelOk returns a tuple with the SentNoiseLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSentNoiseLevel

`func (o *MicrosoftGraphCallRecordsDeviceInfo) SetSentNoiseLevel(v int32)`

SetSentNoiseLevel sets SentNoiseLevel field to given value.

### HasSentNoiseLevel

`func (o *MicrosoftGraphCallRecordsDeviceInfo) HasSentNoiseLevel() bool`

HasSentNoiseLevel returns a boolean if a field has been set.

### SetSentNoiseLevelNil

`func (o *MicrosoftGraphCallRecordsDeviceInfo) SetSentNoiseLevelNil(b bool)`

 SetSentNoiseLevelNil sets the value for SentNoiseLevel to be an explicit nil

### UnsetSentNoiseLevel
`func (o *MicrosoftGraphCallRecordsDeviceInfo) UnsetSentNoiseLevel()`

UnsetSentNoiseLevel ensures that no value is present for SentNoiseLevel, not even an explicit nil
### GetSentSignalLevel

`func (o *MicrosoftGraphCallRecordsDeviceInfo) GetSentSignalLevel() int32`

GetSentSignalLevel returns the SentSignalLevel field if non-nil, zero value otherwise.

### GetSentSignalLevelOk

`func (o *MicrosoftGraphCallRecordsDeviceInfo) GetSentSignalLevelOk() (*int32, bool)`

GetSentSignalLevelOk returns a tuple with the SentSignalLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSentSignalLevel

`func (o *MicrosoftGraphCallRecordsDeviceInfo) SetSentSignalLevel(v int32)`

SetSentSignalLevel sets SentSignalLevel field to given value.

### HasSentSignalLevel

`func (o *MicrosoftGraphCallRecordsDeviceInfo) HasSentSignalLevel() bool`

HasSentSignalLevel returns a boolean if a field has been set.

### SetSentSignalLevelNil

`func (o *MicrosoftGraphCallRecordsDeviceInfo) SetSentSignalLevelNil(b bool)`

 SetSentSignalLevelNil sets the value for SentSignalLevel to be an explicit nil

### UnsetSentSignalLevel
`func (o *MicrosoftGraphCallRecordsDeviceInfo) UnsetSentSignalLevel()`

UnsetSentSignalLevel ensures that no value is present for SentSignalLevel, not even an explicit nil
### GetSpeakerGlitchRate

`func (o *MicrosoftGraphCallRecordsDeviceInfo) GetSpeakerGlitchRate() AnyOfnumberstringstring`

GetSpeakerGlitchRate returns the SpeakerGlitchRate field if non-nil, zero value otherwise.

### GetSpeakerGlitchRateOk

`func (o *MicrosoftGraphCallRecordsDeviceInfo) GetSpeakerGlitchRateOk() (*AnyOfnumberstringstring, bool)`

GetSpeakerGlitchRateOk returns a tuple with the SpeakerGlitchRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeakerGlitchRate

`func (o *MicrosoftGraphCallRecordsDeviceInfo) SetSpeakerGlitchRate(v AnyOfnumberstringstring)`

SetSpeakerGlitchRate sets SpeakerGlitchRate field to given value.

### HasSpeakerGlitchRate

`func (o *MicrosoftGraphCallRecordsDeviceInfo) HasSpeakerGlitchRate() bool`

HasSpeakerGlitchRate returns a boolean if a field has been set.

### SetSpeakerGlitchRateNil

`func (o *MicrosoftGraphCallRecordsDeviceInfo) SetSpeakerGlitchRateNil(b bool)`

 SetSpeakerGlitchRateNil sets the value for SpeakerGlitchRate to be an explicit nil

### UnsetSpeakerGlitchRate
`func (o *MicrosoftGraphCallRecordsDeviceInfo) UnsetSpeakerGlitchRate()`

UnsetSpeakerGlitchRate ensures that no value is present for SpeakerGlitchRate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


