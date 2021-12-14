# MicrosoftGraphTeleconferenceDeviceQuality

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallChainId** | Pointer to **string** | A unique identifier for all  the participant calls in a conference or a unique identifier for two participant calls in P2P call. This needs to be copied over from Microsoft.Graph.Call.CallChainId. | [optional] 
**CloudServiceDeploymentEnvironment** | Pointer to **NullableString** | A geo-region where the service is deployed, such as ProdNoam. | [optional] 
**CloudServiceDeploymentId** | Pointer to **NullableString** | A unique deployment identifier assigned by Azure. | [optional] 
**CloudServiceInstanceName** | Pointer to **NullableString** | The Azure deployed cloud service instance name, such as FrontEnd_IN_3. | [optional] 
**CloudServiceName** | Pointer to **NullableString** | The Azure deployed cloud service name, such as contoso.cloudapp.net. | [optional] 
**DeviceDescription** | Pointer to **string** | Any additional description, such as VTC Bldg 30/21. | [optional] 
**DeviceName** | Pointer to **string** | The user media agent name, such as Cisco SX80. | [optional] 
**MediaLegId** | Pointer to **string** | A unique identifier for a specific media leg of a participant in a conference.  One participant can have multiple media leg identifiers if retargeting happens. CVI partner assigns this value. | [optional] 
**MediaQualityList** | Pointer to [**[]MicrosoftGraphTeleconferenceDeviceMediaQuality**](MicrosoftGraphTeleconferenceDeviceMediaQuality.md) | The list of media qualities in a media session (call), such as audio quality, video quality, and/or screen sharing quality. | [optional] 
**ParticipantId** | Pointer to **string** | A unique identifier for a specific participant in a conference. The CVI partner needs to copy over Call.MyParticipantId to this property. | [optional] 

## Methods

### NewMicrosoftGraphTeleconferenceDeviceQuality

`func NewMicrosoftGraphTeleconferenceDeviceQuality() *MicrosoftGraphTeleconferenceDeviceQuality`

NewMicrosoftGraphTeleconferenceDeviceQuality instantiates a new MicrosoftGraphTeleconferenceDeviceQuality object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphTeleconferenceDeviceQualityWithDefaults

`func NewMicrosoftGraphTeleconferenceDeviceQualityWithDefaults() *MicrosoftGraphTeleconferenceDeviceQuality`

NewMicrosoftGraphTeleconferenceDeviceQualityWithDefaults instantiates a new MicrosoftGraphTeleconferenceDeviceQuality object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallChainId

`func (o *MicrosoftGraphTeleconferenceDeviceQuality) GetCallChainId() string`

GetCallChainId returns the CallChainId field if non-nil, zero value otherwise.

### GetCallChainIdOk

`func (o *MicrosoftGraphTeleconferenceDeviceQuality) GetCallChainIdOk() (*string, bool)`

GetCallChainIdOk returns a tuple with the CallChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallChainId

`func (o *MicrosoftGraphTeleconferenceDeviceQuality) SetCallChainId(v string)`

SetCallChainId sets CallChainId field to given value.

### HasCallChainId

`func (o *MicrosoftGraphTeleconferenceDeviceQuality) HasCallChainId() bool`

HasCallChainId returns a boolean if a field has been set.

### GetCloudServiceDeploymentEnvironment

`func (o *MicrosoftGraphTeleconferenceDeviceQuality) GetCloudServiceDeploymentEnvironment() string`

GetCloudServiceDeploymentEnvironment returns the CloudServiceDeploymentEnvironment field if non-nil, zero value otherwise.

### GetCloudServiceDeploymentEnvironmentOk

`func (o *MicrosoftGraphTeleconferenceDeviceQuality) GetCloudServiceDeploymentEnvironmentOk() (*string, bool)`

GetCloudServiceDeploymentEnvironmentOk returns a tuple with the CloudServiceDeploymentEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudServiceDeploymentEnvironment

`func (o *MicrosoftGraphTeleconferenceDeviceQuality) SetCloudServiceDeploymentEnvironment(v string)`

SetCloudServiceDeploymentEnvironment sets CloudServiceDeploymentEnvironment field to given value.

### HasCloudServiceDeploymentEnvironment

`func (o *MicrosoftGraphTeleconferenceDeviceQuality) HasCloudServiceDeploymentEnvironment() bool`

HasCloudServiceDeploymentEnvironment returns a boolean if a field has been set.

### SetCloudServiceDeploymentEnvironmentNil

`func (o *MicrosoftGraphTeleconferenceDeviceQuality) SetCloudServiceDeploymentEnvironmentNil(b bool)`

 SetCloudServiceDeploymentEnvironmentNil sets the value for CloudServiceDeploymentEnvironment to be an explicit nil

### UnsetCloudServiceDeploymentEnvironment
`func (o *MicrosoftGraphTeleconferenceDeviceQuality) UnsetCloudServiceDeploymentEnvironment()`

UnsetCloudServiceDeploymentEnvironment ensures that no value is present for CloudServiceDeploymentEnvironment, not even an explicit nil
### GetCloudServiceDeploymentId

`func (o *MicrosoftGraphTeleconferenceDeviceQuality) GetCloudServiceDeploymentId() string`

GetCloudServiceDeploymentId returns the CloudServiceDeploymentId field if non-nil, zero value otherwise.

### GetCloudServiceDeploymentIdOk

`func (o *MicrosoftGraphTeleconferenceDeviceQuality) GetCloudServiceDeploymentIdOk() (*string, bool)`

GetCloudServiceDeploymentIdOk returns a tuple with the CloudServiceDeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudServiceDeploymentId

`func (o *MicrosoftGraphTeleconferenceDeviceQuality) SetCloudServiceDeploymentId(v string)`

SetCloudServiceDeploymentId sets CloudServiceDeploymentId field to given value.

### HasCloudServiceDeploymentId

`func (o *MicrosoftGraphTeleconferenceDeviceQuality) HasCloudServiceDeploymentId() bool`

HasCloudServiceDeploymentId returns a boolean if a field has been set.

### SetCloudServiceDeploymentIdNil

`func (o *MicrosoftGraphTeleconferenceDeviceQuality) SetCloudServiceDeploymentIdNil(b bool)`

 SetCloudServiceDeploymentIdNil sets the value for CloudServiceDeploymentId to be an explicit nil

### UnsetCloudServiceDeploymentId
`func (o *MicrosoftGraphTeleconferenceDeviceQuality) UnsetCloudServiceDeploymentId()`

UnsetCloudServiceDeploymentId ensures that no value is present for CloudServiceDeploymentId, not even an explicit nil
### GetCloudServiceInstanceName

`func (o *MicrosoftGraphTeleconferenceDeviceQuality) GetCloudServiceInstanceName() string`

GetCloudServiceInstanceName returns the CloudServiceInstanceName field if non-nil, zero value otherwise.

### GetCloudServiceInstanceNameOk

`func (o *MicrosoftGraphTeleconferenceDeviceQuality) GetCloudServiceInstanceNameOk() (*string, bool)`

GetCloudServiceInstanceNameOk returns a tuple with the CloudServiceInstanceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudServiceInstanceName

`func (o *MicrosoftGraphTeleconferenceDeviceQuality) SetCloudServiceInstanceName(v string)`

SetCloudServiceInstanceName sets CloudServiceInstanceName field to given value.

### HasCloudServiceInstanceName

`func (o *MicrosoftGraphTeleconferenceDeviceQuality) HasCloudServiceInstanceName() bool`

HasCloudServiceInstanceName returns a boolean if a field has been set.

### SetCloudServiceInstanceNameNil

`func (o *MicrosoftGraphTeleconferenceDeviceQuality) SetCloudServiceInstanceNameNil(b bool)`

 SetCloudServiceInstanceNameNil sets the value for CloudServiceInstanceName to be an explicit nil

### UnsetCloudServiceInstanceName
`func (o *MicrosoftGraphTeleconferenceDeviceQuality) UnsetCloudServiceInstanceName()`

UnsetCloudServiceInstanceName ensures that no value is present for CloudServiceInstanceName, not even an explicit nil
### GetCloudServiceName

`func (o *MicrosoftGraphTeleconferenceDeviceQuality) GetCloudServiceName() string`

GetCloudServiceName returns the CloudServiceName field if non-nil, zero value otherwise.

### GetCloudServiceNameOk

`func (o *MicrosoftGraphTeleconferenceDeviceQuality) GetCloudServiceNameOk() (*string, bool)`

GetCloudServiceNameOk returns a tuple with the CloudServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudServiceName

`func (o *MicrosoftGraphTeleconferenceDeviceQuality) SetCloudServiceName(v string)`

SetCloudServiceName sets CloudServiceName field to given value.

### HasCloudServiceName

`func (o *MicrosoftGraphTeleconferenceDeviceQuality) HasCloudServiceName() bool`

HasCloudServiceName returns a boolean if a field has been set.

### SetCloudServiceNameNil

`func (o *MicrosoftGraphTeleconferenceDeviceQuality) SetCloudServiceNameNil(b bool)`

 SetCloudServiceNameNil sets the value for CloudServiceName to be an explicit nil

### UnsetCloudServiceName
`func (o *MicrosoftGraphTeleconferenceDeviceQuality) UnsetCloudServiceName()`

UnsetCloudServiceName ensures that no value is present for CloudServiceName, not even an explicit nil
### GetDeviceDescription

`func (o *MicrosoftGraphTeleconferenceDeviceQuality) GetDeviceDescription() string`

GetDeviceDescription returns the DeviceDescription field if non-nil, zero value otherwise.

### GetDeviceDescriptionOk

`func (o *MicrosoftGraphTeleconferenceDeviceQuality) GetDeviceDescriptionOk() (*string, bool)`

GetDeviceDescriptionOk returns a tuple with the DeviceDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceDescription

`func (o *MicrosoftGraphTeleconferenceDeviceQuality) SetDeviceDescription(v string)`

SetDeviceDescription sets DeviceDescription field to given value.

### HasDeviceDescription

`func (o *MicrosoftGraphTeleconferenceDeviceQuality) HasDeviceDescription() bool`

HasDeviceDescription returns a boolean if a field has been set.

### GetDeviceName

`func (o *MicrosoftGraphTeleconferenceDeviceQuality) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *MicrosoftGraphTeleconferenceDeviceQuality) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *MicrosoftGraphTeleconferenceDeviceQuality) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.

### HasDeviceName

`func (o *MicrosoftGraphTeleconferenceDeviceQuality) HasDeviceName() bool`

HasDeviceName returns a boolean if a field has been set.

### GetMediaLegId

`func (o *MicrosoftGraphTeleconferenceDeviceQuality) GetMediaLegId() string`

GetMediaLegId returns the MediaLegId field if non-nil, zero value otherwise.

### GetMediaLegIdOk

`func (o *MicrosoftGraphTeleconferenceDeviceQuality) GetMediaLegIdOk() (*string, bool)`

GetMediaLegIdOk returns a tuple with the MediaLegId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaLegId

`func (o *MicrosoftGraphTeleconferenceDeviceQuality) SetMediaLegId(v string)`

SetMediaLegId sets MediaLegId field to given value.

### HasMediaLegId

`func (o *MicrosoftGraphTeleconferenceDeviceQuality) HasMediaLegId() bool`

HasMediaLegId returns a boolean if a field has been set.

### GetMediaQualityList

`func (o *MicrosoftGraphTeleconferenceDeviceQuality) GetMediaQualityList() []MicrosoftGraphTeleconferenceDeviceMediaQuality`

GetMediaQualityList returns the MediaQualityList field if non-nil, zero value otherwise.

### GetMediaQualityListOk

`func (o *MicrosoftGraphTeleconferenceDeviceQuality) GetMediaQualityListOk() (*[]MicrosoftGraphTeleconferenceDeviceMediaQuality, bool)`

GetMediaQualityListOk returns a tuple with the MediaQualityList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaQualityList

`func (o *MicrosoftGraphTeleconferenceDeviceQuality) SetMediaQualityList(v []MicrosoftGraphTeleconferenceDeviceMediaQuality)`

SetMediaQualityList sets MediaQualityList field to given value.

### HasMediaQualityList

`func (o *MicrosoftGraphTeleconferenceDeviceQuality) HasMediaQualityList() bool`

HasMediaQualityList returns a boolean if a field has been set.

### GetParticipantId

`func (o *MicrosoftGraphTeleconferenceDeviceQuality) GetParticipantId() string`

GetParticipantId returns the ParticipantId field if non-nil, zero value otherwise.

### GetParticipantIdOk

`func (o *MicrosoftGraphTeleconferenceDeviceQuality) GetParticipantIdOk() (*string, bool)`

GetParticipantIdOk returns a tuple with the ParticipantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipantId

`func (o *MicrosoftGraphTeleconferenceDeviceQuality) SetParticipantId(v string)`

SetParticipantId sets ParticipantId field to given value.

### HasParticipantId

`func (o *MicrosoftGraphTeleconferenceDeviceQuality) HasParticipantId() bool`

HasParticipantId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


