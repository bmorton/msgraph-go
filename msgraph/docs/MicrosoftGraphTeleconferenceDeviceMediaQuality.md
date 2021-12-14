# MicrosoftGraphTeleconferenceDeviceMediaQuality

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AverageInboundJitter** | Pointer to **NullableString** | The average inbound stream network jitter. | [optional] 
**AverageInboundPacketLossRateInPercentage** | Pointer to [**AnyOfnumberstringstring**](anyOf&lt;number,string,string&gt;.md) | The average inbound stream packet loss rate in percentage (0-100). For example, 0.01 means 0.01%. | [optional] 
**AverageInboundRoundTripDelay** | Pointer to **NullableString** | The average inbound stream network round trip delay. | [optional] 
**AverageOutboundJitter** | Pointer to **NullableString** | The average outbound stream network jitter. | [optional] 
**AverageOutboundPacketLossRateInPercentage** | Pointer to [**AnyOfnumberstringstring**](anyOf&lt;number,string,string&gt;.md) | The average outbound stream packet loss rate in percentage (0-100). For example, 0.01 means 0.01%. | [optional] 
**AverageOutboundRoundTripDelay** | Pointer to **NullableString** | The average outbound stream network round trip delay. | [optional] 
**ChannelIndex** | Pointer to **int32** | The channel index of media. Indexing begins with 1.  If a media session contains 3 video modalities, channel indexes will be 1, 2, and 3. | [optional] 
**InboundPackets** | Pointer to **NullableInt64** | The total number of the inbound packets. | [optional] 
**LocalIPAddress** | Pointer to **NullableString** | the local IP address for the media session. | [optional] 
**LocalPort** | Pointer to **NullableInt32** | The local media port. | [optional] 
**MaximumInboundJitter** | Pointer to **NullableString** | The maximum inbound stream network jitter. | [optional] 
**MaximumInboundPacketLossRateInPercentage** | Pointer to [**AnyOfnumberstringstring**](anyOf&lt;number,string,string&gt;.md) | The maximum inbound stream packet loss rate in percentage (0-100). For example, 0.01 means 0.01%. | [optional] 
**MaximumInboundRoundTripDelay** | Pointer to **NullableString** | The maximum inbound stream network round trip delay. | [optional] 
**MaximumOutboundJitter** | Pointer to **NullableString** | The maximum outbound stream network jitter. | [optional] 
**MaximumOutboundPacketLossRateInPercentage** | Pointer to [**AnyOfnumberstringstring**](anyOf&lt;number,string,string&gt;.md) | The maximum outbound stream packet loss rate in percentage (0-100). For example, 0.01 means 0.01%. | [optional] 
**MaximumOutboundRoundTripDelay** | Pointer to **NullableString** | The maximum outbound stream network round trip delay. | [optional] 
**MediaDuration** | Pointer to **NullableString** | The total modality duration. If the media enabled and disabled multiple times, MediaDuration will the summation of all of the durations. | [optional] 
**NetworkLinkSpeedInBytes** | Pointer to **NullableInt64** | The network link speed in bytes | [optional] 
**OutboundPackets** | Pointer to **NullableInt64** | The total number of the outbound packets. | [optional] 
**RemoteIPAddress** | Pointer to **NullableString** | The remote IP address for the media session. | [optional] 
**RemotePort** | Pointer to **NullableInt32** | The remote media port. | [optional] 

## Methods

### NewMicrosoftGraphTeleconferenceDeviceMediaQuality

`func NewMicrosoftGraphTeleconferenceDeviceMediaQuality() *MicrosoftGraphTeleconferenceDeviceMediaQuality`

NewMicrosoftGraphTeleconferenceDeviceMediaQuality instantiates a new MicrosoftGraphTeleconferenceDeviceMediaQuality object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphTeleconferenceDeviceMediaQualityWithDefaults

`func NewMicrosoftGraphTeleconferenceDeviceMediaQualityWithDefaults() *MicrosoftGraphTeleconferenceDeviceMediaQuality`

NewMicrosoftGraphTeleconferenceDeviceMediaQualityWithDefaults instantiates a new MicrosoftGraphTeleconferenceDeviceMediaQuality object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAverageInboundJitter

`func (o *MicrosoftGraphTeleconferenceDeviceMediaQuality) GetAverageInboundJitter() string`

GetAverageInboundJitter returns the AverageInboundJitter field if non-nil, zero value otherwise.

### GetAverageInboundJitterOk

`func (o *MicrosoftGraphTeleconferenceDeviceMediaQuality) GetAverageInboundJitterOk() (*string, bool)`

GetAverageInboundJitterOk returns a tuple with the AverageInboundJitter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageInboundJitter

`func (o *MicrosoftGraphTeleconferenceDeviceMediaQuality) SetAverageInboundJitter(v string)`

SetAverageInboundJitter sets AverageInboundJitter field to given value.

### HasAverageInboundJitter

`func (o *MicrosoftGraphTeleconferenceDeviceMediaQuality) HasAverageInboundJitter() bool`

HasAverageInboundJitter returns a boolean if a field has been set.

### SetAverageInboundJitterNil

`func (o *MicrosoftGraphTeleconferenceDeviceMediaQuality) SetAverageInboundJitterNil(b bool)`

 SetAverageInboundJitterNil sets the value for AverageInboundJitter to be an explicit nil

### UnsetAverageInboundJitter
`func (o *MicrosoftGraphTeleconferenceDeviceMediaQuality) UnsetAverageInboundJitter()`

UnsetAverageInboundJitter ensures that no value is present for AverageInboundJitter, not even an explicit nil
### GetAverageInboundPacketLossRateInPercentage

`func (o *MicrosoftGraphTeleconferenceDeviceMediaQuality) GetAverageInboundPacketLossRateInPercentage() AnyOfnumberstringstring`

GetAverageInboundPacketLossRateInPercentage returns the AverageInboundPacketLossRateInPercentage field if non-nil, zero value otherwise.

### GetAverageInboundPacketLossRateInPercentageOk

`func (o *MicrosoftGraphTeleconferenceDeviceMediaQuality) GetAverageInboundPacketLossRateInPercentageOk() (*AnyOfnumberstringstring, bool)`

GetAverageInboundPacketLossRateInPercentageOk returns a tuple with the AverageInboundPacketLossRateInPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageInboundPacketLossRateInPercentage

`func (o *MicrosoftGraphTeleconferenceDeviceMediaQuality) SetAverageInboundPacketLossRateInPercentage(v AnyOfnumberstringstring)`

SetAverageInboundPacketLossRateInPercentage sets AverageInboundPacketLossRateInPercentage field to given value.

### HasAverageInboundPacketLossRateInPercentage

`func (o *MicrosoftGraphTeleconferenceDeviceMediaQuality) HasAverageInboundPacketLossRateInPercentage() bool`

HasAverageInboundPacketLossRateInPercentage returns a boolean if a field has been set.

### SetAverageInboundPacketLossRateInPercentageNil

`func (o *MicrosoftGraphTeleconferenceDeviceMediaQuality) SetAverageInboundPacketLossRateInPercentageNil(b bool)`

 SetAverageInboundPacketLossRateInPercentageNil sets the value for AverageInboundPacketLossRateInPercentage to be an explicit nil

### UnsetAverageInboundPacketLossRateInPercentage
`func (o *MicrosoftGraphTeleconferenceDeviceMediaQuality) UnsetAverageInboundPacketLossRateInPercentage()`

UnsetAverageInboundPacketLossRateInPercentage ensures that no value is present for AverageInboundPacketLossRateInPercentage, not even an explicit nil
### GetAverageInboundRoundTripDelay

`func (o *MicrosoftGraphTeleconferenceDeviceMediaQuality) GetAverageInboundRoundTripDelay() string`

GetAverageInboundRoundTripDelay returns the AverageInboundRoundTripDelay field if non-nil, zero value otherwise.

### GetAverageInboundRoundTripDelayOk

`func (o *MicrosoftGraphTeleconferenceDeviceMediaQuality) GetAverageInboundRoundTripDelayOk() (*string, bool)`

GetAverageInboundRoundTripDelayOk returns a tuple with the AverageInboundRoundTripDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageInboundRoundTripDelay

`func (o *MicrosoftGraphTeleconferenceDeviceMediaQuality) SetAverageInboundRoundTripDelay(v string)`

SetAverageInboundRoundTripDelay sets AverageInboundRoundTripDelay field to given value.

### HasAverageInboundRoundTripDelay

`func (o *MicrosoftGraphTeleconferenceDeviceMediaQuality) HasAverageInboundRoundTripDelay() bool`

HasAverageInboundRoundTripDelay returns a boolean if a field has been set.

### SetAverageInboundRoundTripDelayNil

`func (o *MicrosoftGraphTeleconferenceDeviceMediaQuality) SetAverageInboundRoundTripDelayNil(b bool)`

 SetAverageInboundRoundTripDelayNil sets the value for AverageInboundRoundTripDelay to be an explicit nil

### UnsetAverageInboundRoundTripDelay
`func (o *MicrosoftGraphTeleconferenceDeviceMediaQuality) UnsetAverageInboundRoundTripDelay()`

UnsetAverageInboundRoundTripDelay ensures that no value is present for AverageInboundRoundTripDelay, not even an explicit nil
### GetAverageOutboundJitter

`func (o *MicrosoftGraphTeleconferenceDeviceMediaQuality) GetAverageOutboundJitter() string`

GetAverageOutboundJitter returns the AverageOutboundJitter field if non-nil, zero value otherwise.

### GetAverageOutboundJitterOk

`func (o *MicrosoftGraphTeleconferenceDeviceMediaQuality) GetAverageOutboundJitterOk() (*string, bool)`

GetAverageOutboundJitterOk returns a tuple with the AverageOutboundJitter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageOutboundJitter

`func (o *MicrosoftGraphTeleconferenceDeviceMediaQuality) SetAverageOutboundJitter(v string)`

SetAverageOutboundJitter sets AverageOutboundJitter field to given value.

### HasAverageOutboundJitter

`func (o *MicrosoftGraphTeleconferenceDeviceMediaQuality) HasAverageOutboundJitter() bool`

HasAverageOutboundJitter returns a boolean if a field has been set.

### SetAverageOutboundJitterNil

`func (o *MicrosoftGraphTeleconferenceDeviceMediaQuality) SetAverageOutboundJitterNil(b bool)`

 SetAverageOutboundJitterNil sets the value for AverageOutboundJitter to be an explicit nil

### UnsetAverageOutboundJitter
`func (o *MicrosoftGraphTeleconferenceDeviceMediaQuality) UnsetAverageOutboundJitter()`

UnsetAverageOutboundJitter ensures that no value is present for AverageOutboundJitter, not even an explicit nil
### GetAverageOutboundPacketLossRateInPercentage

`func (o *MicrosoftGraphTeleconferenceDeviceMediaQuality) GetAverageOutboundPacketLossRateInPercentage() AnyOfnumberstringstring`

GetAverageOutboundPacketLossRateInPercentage returns the AverageOutboundPacketLossRateInPercentage field if non-nil, zero value otherwise.

### GetAverageOutboundPacketLossRateInPercentageOk

`func (o *MicrosoftGraphTeleconferenceDeviceMediaQuality) GetAverageOutboundPacketLossRateInPercentageOk() (*AnyOfnumberstringstring, bool)`

GetAverageOutboundPacketLossRateInPercentageOk returns a tuple with the AverageOutboundPacketLossRateInPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageOutboundPacketLossRateInPercentage

`func (o *MicrosoftGraphTeleconferenceDeviceMediaQuality) SetAverageOutboundPacketLossRateInPercentage(v AnyOfnumberstringstring)`

SetAverageOutboundPacketLossRateInPercentage sets AverageOutboundPacketLossRateInPercentage field to given value.

### HasAverageOutboundPacketLossRateInPercentage

`func (o *MicrosoftGraphTeleconferenceDeviceMediaQuality) HasAverageOutboundPacketLossRateInPercentage() bool`

HasAverageOutboundPacketLossRateInPercentage returns a boolean if a field has been set.

### SetAverageOutboundPacketLossRateInPercentageNil

`func (o *MicrosoftGraphTeleconferenceDeviceMediaQuality) SetAverageOutboundPacketLossRateInPercentageNil(b bool)`

 SetAverageOutboundPacketLossRateInPercentageNil sets the value for AverageOutboundPacketLossRateInPercentage to be an explicit nil

### UnsetAverageOutboundPacketLossRateInPercentage
`func (o *MicrosoftGraphTeleconferenceDeviceMediaQuality) UnsetAverageOutboundPacketLossRateInPercentage()`

UnsetAverageOutboundPacketLossRateInPercentage ensures that no value is present for AverageOutboundPacketLossRateInPercentage, not even an explicit nil
### GetAverageOutboundRoundTripDelay

`func (o *MicrosoftGraphTeleconferenceDeviceMediaQuality) GetAverageOutboundRoundTripDelay() string`

GetAverageOutboundRoundTripDelay returns the AverageOutboundRoundTripDelay field if non-nil, zero value otherwise.

### GetAverageOutboundRoundTripDelayOk

`func (o *MicrosoftGraphTeleconferenceDeviceMediaQuality) GetAverageOutboundRoundTripDelayOk() (*string, bool)`

GetAverageOutboundRoundTripDelayOk returns a tuple with the AverageOutboundRoundTripDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageOutboundRoundTripDelay

`func (o *MicrosoftGraphTeleconferenceDeviceMediaQuality) SetAverageOutboundRoundTripDelay(v string)`

SetAverageOutboundRoundTripDelay sets AverageOutboundRoundTripDelay field to given value.

### HasAverageOutboundRoundTripDelay

`func (o *MicrosoftGraphTeleconferenceDeviceMediaQuality) HasAverageOutboundRoundTripDelay() bool`

HasAverageOutboundRoundTripDelay returns a boolean if a field has been set.

### SetAverageOutboundRoundTripDelayNil

`func (o *MicrosoftGraphTeleconferenceDeviceMediaQuality) SetAverageOutboundRoundTripDelayNil(b bool)`

 SetAverageOutboundRoundTripDelayNil sets the value for AverageOutboundRoundTripDelay to be an explicit nil

### UnsetAverageOutboundRoundTripDelay
`func (o *MicrosoftGraphTeleconferenceDeviceMediaQuality) UnsetAverageOutboundRoundTripDelay()`

UnsetAverageOutboundRoundTripDelay ensures that no value is present for AverageOutboundRoundTripDelay, not even an explicit nil
### GetChannelIndex

`func (o *MicrosoftGraphTeleconferenceDeviceMediaQuality) GetChannelIndex() int32`

GetChannelIndex returns the ChannelIndex field if non-nil, zero value otherwise.

### GetChannelIndexOk

`func (o *MicrosoftGraphTeleconferenceDeviceMediaQuality) GetChannelIndexOk() (*int32, bool)`

GetChannelIndexOk returns a tuple with the ChannelIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelIndex

`func (o *MicrosoftGraphTeleconferenceDeviceMediaQuality) SetChannelIndex(v int32)`

SetChannelIndex sets ChannelIndex field to given value.

### HasChannelIndex

`func (o *MicrosoftGraphTeleconferenceDeviceMediaQuality) HasChannelIndex() bool`

HasChannelIndex returns a boolean if a field has been set.

### GetInboundPackets

`func (o *MicrosoftGraphTeleconferenceDeviceMediaQuality) GetInboundPackets() int64`

GetInboundPackets returns the InboundPackets field if non-nil, zero value otherwise.

### GetInboundPacketsOk

`func (o *MicrosoftGraphTeleconferenceDeviceMediaQuality) GetInboundPacketsOk() (*int64, bool)`

GetInboundPacketsOk returns a tuple with the InboundPackets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInboundPackets

`func (o *MicrosoftGraphTeleconferenceDeviceMediaQuality) SetInboundPackets(v int64)`

SetInboundPackets sets InboundPackets field to given value.

### HasInboundPackets

`func (o *MicrosoftGraphTeleconferenceDeviceMediaQuality) HasInboundPackets() bool`

HasInboundPackets returns a boolean if a field has been set.

### SetInboundPacketsNil

`func (o *MicrosoftGraphTeleconferenceDeviceMediaQuality) SetInboundPacketsNil(b bool)`

 SetInboundPacketsNil sets the value for InboundPackets to be an explicit nil

### UnsetInboundPackets
`func (o *MicrosoftGraphTeleconferenceDeviceMediaQuality) UnsetInboundPackets()`

UnsetInboundPackets ensures that no value is present for InboundPackets, not even an explicit nil
### GetLocalIPAddress

`func (o *MicrosoftGraphTeleconferenceDeviceMediaQuality) GetLocalIPAddress() string`

GetLocalIPAddress returns the LocalIPAddress field if non-nil, zero value otherwise.

### GetLocalIPAddressOk

`func (o *MicrosoftGraphTeleconferenceDeviceMediaQuality) GetLocalIPAddressOk() (*string, bool)`

GetLocalIPAddressOk returns a tuple with the LocalIPAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalIPAddress

`func (o *MicrosoftGraphTeleconferenceDeviceMediaQuality) SetLocalIPAddress(v string)`

SetLocalIPAddress sets LocalIPAddress field to given value.

### HasLocalIPAddress

`func (o *MicrosoftGraphTeleconferenceDeviceMediaQuality) HasLocalIPAddress() bool`

HasLocalIPAddress returns a boolean if a field has been set.

### SetLocalIPAddressNil

`func (o *MicrosoftGraphTeleconferenceDeviceMediaQuality) SetLocalIPAddressNil(b bool)`

 SetLocalIPAddressNil sets the value for LocalIPAddress to be an explicit nil

### UnsetLocalIPAddress
`func (o *MicrosoftGraphTeleconferenceDeviceMediaQuality) UnsetLocalIPAddress()`

UnsetLocalIPAddress ensures that no value is present for LocalIPAddress, not even an explicit nil
### GetLocalPort

`func (o *MicrosoftGraphTeleconferenceDeviceMediaQuality) GetLocalPort() int32`

GetLocalPort returns the LocalPort field if non-nil, zero value otherwise.

### GetLocalPortOk

`func (o *MicrosoftGraphTeleconferenceDeviceMediaQuality) GetLocalPortOk() (*int32, bool)`

GetLocalPortOk returns a tuple with the LocalPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalPort

`func (o *MicrosoftGraphTeleconferenceDeviceMediaQuality) SetLocalPort(v int32)`

SetLocalPort sets LocalPort field to given value.

### HasLocalPort

`func (o *MicrosoftGraphTeleconferenceDeviceMediaQuality) HasLocalPort() bool`

HasLocalPort returns a boolean if a field has been set.

### SetLocalPortNil

`func (o *MicrosoftGraphTeleconferenceDeviceMediaQuality) SetLocalPortNil(b bool)`

 SetLocalPortNil sets the value for LocalPort to be an explicit nil

### UnsetLocalPort
`func (o *MicrosoftGraphTeleconferenceDeviceMediaQuality) UnsetLocalPort()`

UnsetLocalPort ensures that no value is present for LocalPort, not even an explicit nil
### GetMaximumInboundJitter

`func (o *MicrosoftGraphTeleconferenceDeviceMediaQuality) GetMaximumInboundJitter() string`

GetMaximumInboundJitter returns the MaximumInboundJitter field if non-nil, zero value otherwise.

### GetMaximumInboundJitterOk

`func (o *MicrosoftGraphTeleconferenceDeviceMediaQuality) GetMaximumInboundJitterOk() (*string, bool)`

GetMaximumInboundJitterOk returns a tuple with the MaximumInboundJitter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumInboundJitter

`func (o *MicrosoftGraphTeleconferenceDeviceMediaQuality) SetMaximumInboundJitter(v string)`

SetMaximumInboundJitter sets MaximumInboundJitter field to given value.

### HasMaximumInboundJitter

`func (o *MicrosoftGraphTeleconferenceDeviceMediaQuality) HasMaximumInboundJitter() bool`

HasMaximumInboundJitter returns a boolean if a field has been set.

### SetMaximumInboundJitterNil

`func (o *MicrosoftGraphTeleconferenceDeviceMediaQuality) SetMaximumInboundJitterNil(b bool)`

 SetMaximumInboundJitterNil sets the value for MaximumInboundJitter to be an explicit nil

### UnsetMaximumInboundJitter
`func (o *MicrosoftGraphTeleconferenceDeviceMediaQuality) UnsetMaximumInboundJitter()`

UnsetMaximumInboundJitter ensures that no value is present for MaximumInboundJitter, not even an explicit nil
### GetMaximumInboundPacketLossRateInPercentage

`func (o *MicrosoftGraphTeleconferenceDeviceMediaQuality) GetMaximumInboundPacketLossRateInPercentage() AnyOfnumberstringstring`

GetMaximumInboundPacketLossRateInPercentage returns the MaximumInboundPacketLossRateInPercentage field if non-nil, zero value otherwise.

### GetMaximumInboundPacketLossRateInPercentageOk

`func (o *MicrosoftGraphTeleconferenceDeviceMediaQuality) GetMaximumInboundPacketLossRateInPercentageOk() (*AnyOfnumberstringstring, bool)`

GetMaximumInboundPacketLossRateInPercentageOk returns a tuple with the MaximumInboundPacketLossRateInPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumInboundPacketLossRateInPercentage

`func (o *MicrosoftGraphTeleconferenceDeviceMediaQuality) SetMaximumInboundPacketLossRateInPercentage(v AnyOfnumberstringstring)`

SetMaximumInboundPacketLossRateInPercentage sets MaximumInboundPacketLossRateInPercentage field to given value.

### HasMaximumInboundPacketLossRateInPercentage

`func (o *MicrosoftGraphTeleconferenceDeviceMediaQuality) HasMaximumInboundPacketLossRateInPercentage() bool`

HasMaximumInboundPacketLossRateInPercentage returns a boolean if a field has been set.

### SetMaximumInboundPacketLossRateInPercentageNil

`func (o *MicrosoftGraphTeleconferenceDeviceMediaQuality) SetMaximumInboundPacketLossRateInPercentageNil(b bool)`

 SetMaximumInboundPacketLossRateInPercentageNil sets the value for MaximumInboundPacketLossRateInPercentage to be an explicit nil

### UnsetMaximumInboundPacketLossRateInPercentage
`func (o *MicrosoftGraphTeleconferenceDeviceMediaQuality) UnsetMaximumInboundPacketLossRateInPercentage()`

UnsetMaximumInboundPacketLossRateInPercentage ensures that no value is present for MaximumInboundPacketLossRateInPercentage, not even an explicit nil
### GetMaximumInboundRoundTripDelay

`func (o *MicrosoftGraphTeleconferenceDeviceMediaQuality) GetMaximumInboundRoundTripDelay() string`

GetMaximumInboundRoundTripDelay returns the MaximumInboundRoundTripDelay field if non-nil, zero value otherwise.

### GetMaximumInboundRoundTripDelayOk

`func (o *MicrosoftGraphTeleconferenceDeviceMediaQuality) GetMaximumInboundRoundTripDelayOk() (*string, bool)`

GetMaximumInboundRoundTripDelayOk returns a tuple with the MaximumInboundRoundTripDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumInboundRoundTripDelay

`func (o *MicrosoftGraphTeleconferenceDeviceMediaQuality) SetMaximumInboundRoundTripDelay(v string)`

SetMaximumInboundRoundTripDelay sets MaximumInboundRoundTripDelay field to given value.

### HasMaximumInboundRoundTripDelay

`func (o *MicrosoftGraphTeleconferenceDeviceMediaQuality) HasMaximumInboundRoundTripDelay() bool`

HasMaximumInboundRoundTripDelay returns a boolean if a field has been set.

### SetMaximumInboundRoundTripDelayNil

`func (o *MicrosoftGraphTeleconferenceDeviceMediaQuality) SetMaximumInboundRoundTripDelayNil(b bool)`

 SetMaximumInboundRoundTripDelayNil sets the value for MaximumInboundRoundTripDelay to be an explicit nil

### UnsetMaximumInboundRoundTripDelay
`func (o *MicrosoftGraphTeleconferenceDeviceMediaQuality) UnsetMaximumInboundRoundTripDelay()`

UnsetMaximumInboundRoundTripDelay ensures that no value is present for MaximumInboundRoundTripDelay, not even an explicit nil
### GetMaximumOutboundJitter

`func (o *MicrosoftGraphTeleconferenceDeviceMediaQuality) GetMaximumOutboundJitter() string`

GetMaximumOutboundJitter returns the MaximumOutboundJitter field if non-nil, zero value otherwise.

### GetMaximumOutboundJitterOk

`func (o *MicrosoftGraphTeleconferenceDeviceMediaQuality) GetMaximumOutboundJitterOk() (*string, bool)`

GetMaximumOutboundJitterOk returns a tuple with the MaximumOutboundJitter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumOutboundJitter

`func (o *MicrosoftGraphTeleconferenceDeviceMediaQuality) SetMaximumOutboundJitter(v string)`

SetMaximumOutboundJitter sets MaximumOutboundJitter field to given value.

### HasMaximumOutboundJitter

`func (o *MicrosoftGraphTeleconferenceDeviceMediaQuality) HasMaximumOutboundJitter() bool`

HasMaximumOutboundJitter returns a boolean if a field has been set.

### SetMaximumOutboundJitterNil

`func (o *MicrosoftGraphTeleconferenceDeviceMediaQuality) SetMaximumOutboundJitterNil(b bool)`

 SetMaximumOutboundJitterNil sets the value for MaximumOutboundJitter to be an explicit nil

### UnsetMaximumOutboundJitter
`func (o *MicrosoftGraphTeleconferenceDeviceMediaQuality) UnsetMaximumOutboundJitter()`

UnsetMaximumOutboundJitter ensures that no value is present for MaximumOutboundJitter, not even an explicit nil
### GetMaximumOutboundPacketLossRateInPercentage

`func (o *MicrosoftGraphTeleconferenceDeviceMediaQuality) GetMaximumOutboundPacketLossRateInPercentage() AnyOfnumberstringstring`

GetMaximumOutboundPacketLossRateInPercentage returns the MaximumOutboundPacketLossRateInPercentage field if non-nil, zero value otherwise.

### GetMaximumOutboundPacketLossRateInPercentageOk

`func (o *MicrosoftGraphTeleconferenceDeviceMediaQuality) GetMaximumOutboundPacketLossRateInPercentageOk() (*AnyOfnumberstringstring, bool)`

GetMaximumOutboundPacketLossRateInPercentageOk returns a tuple with the MaximumOutboundPacketLossRateInPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumOutboundPacketLossRateInPercentage

`func (o *MicrosoftGraphTeleconferenceDeviceMediaQuality) SetMaximumOutboundPacketLossRateInPercentage(v AnyOfnumberstringstring)`

SetMaximumOutboundPacketLossRateInPercentage sets MaximumOutboundPacketLossRateInPercentage field to given value.

### HasMaximumOutboundPacketLossRateInPercentage

`func (o *MicrosoftGraphTeleconferenceDeviceMediaQuality) HasMaximumOutboundPacketLossRateInPercentage() bool`

HasMaximumOutboundPacketLossRateInPercentage returns a boolean if a field has been set.

### SetMaximumOutboundPacketLossRateInPercentageNil

`func (o *MicrosoftGraphTeleconferenceDeviceMediaQuality) SetMaximumOutboundPacketLossRateInPercentageNil(b bool)`

 SetMaximumOutboundPacketLossRateInPercentageNil sets the value for MaximumOutboundPacketLossRateInPercentage to be an explicit nil

### UnsetMaximumOutboundPacketLossRateInPercentage
`func (o *MicrosoftGraphTeleconferenceDeviceMediaQuality) UnsetMaximumOutboundPacketLossRateInPercentage()`

UnsetMaximumOutboundPacketLossRateInPercentage ensures that no value is present for MaximumOutboundPacketLossRateInPercentage, not even an explicit nil
### GetMaximumOutboundRoundTripDelay

`func (o *MicrosoftGraphTeleconferenceDeviceMediaQuality) GetMaximumOutboundRoundTripDelay() string`

GetMaximumOutboundRoundTripDelay returns the MaximumOutboundRoundTripDelay field if non-nil, zero value otherwise.

### GetMaximumOutboundRoundTripDelayOk

`func (o *MicrosoftGraphTeleconferenceDeviceMediaQuality) GetMaximumOutboundRoundTripDelayOk() (*string, bool)`

GetMaximumOutboundRoundTripDelayOk returns a tuple with the MaximumOutboundRoundTripDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumOutboundRoundTripDelay

`func (o *MicrosoftGraphTeleconferenceDeviceMediaQuality) SetMaximumOutboundRoundTripDelay(v string)`

SetMaximumOutboundRoundTripDelay sets MaximumOutboundRoundTripDelay field to given value.

### HasMaximumOutboundRoundTripDelay

`func (o *MicrosoftGraphTeleconferenceDeviceMediaQuality) HasMaximumOutboundRoundTripDelay() bool`

HasMaximumOutboundRoundTripDelay returns a boolean if a field has been set.

### SetMaximumOutboundRoundTripDelayNil

`func (o *MicrosoftGraphTeleconferenceDeviceMediaQuality) SetMaximumOutboundRoundTripDelayNil(b bool)`

 SetMaximumOutboundRoundTripDelayNil sets the value for MaximumOutboundRoundTripDelay to be an explicit nil

### UnsetMaximumOutboundRoundTripDelay
`func (o *MicrosoftGraphTeleconferenceDeviceMediaQuality) UnsetMaximumOutboundRoundTripDelay()`

UnsetMaximumOutboundRoundTripDelay ensures that no value is present for MaximumOutboundRoundTripDelay, not even an explicit nil
### GetMediaDuration

`func (o *MicrosoftGraphTeleconferenceDeviceMediaQuality) GetMediaDuration() string`

GetMediaDuration returns the MediaDuration field if non-nil, zero value otherwise.

### GetMediaDurationOk

`func (o *MicrosoftGraphTeleconferenceDeviceMediaQuality) GetMediaDurationOk() (*string, bool)`

GetMediaDurationOk returns a tuple with the MediaDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaDuration

`func (o *MicrosoftGraphTeleconferenceDeviceMediaQuality) SetMediaDuration(v string)`

SetMediaDuration sets MediaDuration field to given value.

### HasMediaDuration

`func (o *MicrosoftGraphTeleconferenceDeviceMediaQuality) HasMediaDuration() bool`

HasMediaDuration returns a boolean if a field has been set.

### SetMediaDurationNil

`func (o *MicrosoftGraphTeleconferenceDeviceMediaQuality) SetMediaDurationNil(b bool)`

 SetMediaDurationNil sets the value for MediaDuration to be an explicit nil

### UnsetMediaDuration
`func (o *MicrosoftGraphTeleconferenceDeviceMediaQuality) UnsetMediaDuration()`

UnsetMediaDuration ensures that no value is present for MediaDuration, not even an explicit nil
### GetNetworkLinkSpeedInBytes

`func (o *MicrosoftGraphTeleconferenceDeviceMediaQuality) GetNetworkLinkSpeedInBytes() int64`

GetNetworkLinkSpeedInBytes returns the NetworkLinkSpeedInBytes field if non-nil, zero value otherwise.

### GetNetworkLinkSpeedInBytesOk

`func (o *MicrosoftGraphTeleconferenceDeviceMediaQuality) GetNetworkLinkSpeedInBytesOk() (*int64, bool)`

GetNetworkLinkSpeedInBytesOk returns a tuple with the NetworkLinkSpeedInBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkLinkSpeedInBytes

`func (o *MicrosoftGraphTeleconferenceDeviceMediaQuality) SetNetworkLinkSpeedInBytes(v int64)`

SetNetworkLinkSpeedInBytes sets NetworkLinkSpeedInBytes field to given value.

### HasNetworkLinkSpeedInBytes

`func (o *MicrosoftGraphTeleconferenceDeviceMediaQuality) HasNetworkLinkSpeedInBytes() bool`

HasNetworkLinkSpeedInBytes returns a boolean if a field has been set.

### SetNetworkLinkSpeedInBytesNil

`func (o *MicrosoftGraphTeleconferenceDeviceMediaQuality) SetNetworkLinkSpeedInBytesNil(b bool)`

 SetNetworkLinkSpeedInBytesNil sets the value for NetworkLinkSpeedInBytes to be an explicit nil

### UnsetNetworkLinkSpeedInBytes
`func (o *MicrosoftGraphTeleconferenceDeviceMediaQuality) UnsetNetworkLinkSpeedInBytes()`

UnsetNetworkLinkSpeedInBytes ensures that no value is present for NetworkLinkSpeedInBytes, not even an explicit nil
### GetOutboundPackets

`func (o *MicrosoftGraphTeleconferenceDeviceMediaQuality) GetOutboundPackets() int64`

GetOutboundPackets returns the OutboundPackets field if non-nil, zero value otherwise.

### GetOutboundPacketsOk

`func (o *MicrosoftGraphTeleconferenceDeviceMediaQuality) GetOutboundPacketsOk() (*int64, bool)`

GetOutboundPacketsOk returns a tuple with the OutboundPackets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutboundPackets

`func (o *MicrosoftGraphTeleconferenceDeviceMediaQuality) SetOutboundPackets(v int64)`

SetOutboundPackets sets OutboundPackets field to given value.

### HasOutboundPackets

`func (o *MicrosoftGraphTeleconferenceDeviceMediaQuality) HasOutboundPackets() bool`

HasOutboundPackets returns a boolean if a field has been set.

### SetOutboundPacketsNil

`func (o *MicrosoftGraphTeleconferenceDeviceMediaQuality) SetOutboundPacketsNil(b bool)`

 SetOutboundPacketsNil sets the value for OutboundPackets to be an explicit nil

### UnsetOutboundPackets
`func (o *MicrosoftGraphTeleconferenceDeviceMediaQuality) UnsetOutboundPackets()`

UnsetOutboundPackets ensures that no value is present for OutboundPackets, not even an explicit nil
### GetRemoteIPAddress

`func (o *MicrosoftGraphTeleconferenceDeviceMediaQuality) GetRemoteIPAddress() string`

GetRemoteIPAddress returns the RemoteIPAddress field if non-nil, zero value otherwise.

### GetRemoteIPAddressOk

`func (o *MicrosoftGraphTeleconferenceDeviceMediaQuality) GetRemoteIPAddressOk() (*string, bool)`

GetRemoteIPAddressOk returns a tuple with the RemoteIPAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteIPAddress

`func (o *MicrosoftGraphTeleconferenceDeviceMediaQuality) SetRemoteIPAddress(v string)`

SetRemoteIPAddress sets RemoteIPAddress field to given value.

### HasRemoteIPAddress

`func (o *MicrosoftGraphTeleconferenceDeviceMediaQuality) HasRemoteIPAddress() bool`

HasRemoteIPAddress returns a boolean if a field has been set.

### SetRemoteIPAddressNil

`func (o *MicrosoftGraphTeleconferenceDeviceMediaQuality) SetRemoteIPAddressNil(b bool)`

 SetRemoteIPAddressNil sets the value for RemoteIPAddress to be an explicit nil

### UnsetRemoteIPAddress
`func (o *MicrosoftGraphTeleconferenceDeviceMediaQuality) UnsetRemoteIPAddress()`

UnsetRemoteIPAddress ensures that no value is present for RemoteIPAddress, not even an explicit nil
### GetRemotePort

`func (o *MicrosoftGraphTeleconferenceDeviceMediaQuality) GetRemotePort() int32`

GetRemotePort returns the RemotePort field if non-nil, zero value otherwise.

### GetRemotePortOk

`func (o *MicrosoftGraphTeleconferenceDeviceMediaQuality) GetRemotePortOk() (*int32, bool)`

GetRemotePortOk returns a tuple with the RemotePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemotePort

`func (o *MicrosoftGraphTeleconferenceDeviceMediaQuality) SetRemotePort(v int32)`

SetRemotePort sets RemotePort field to given value.

### HasRemotePort

`func (o *MicrosoftGraphTeleconferenceDeviceMediaQuality) HasRemotePort() bool`

HasRemotePort returns a boolean if a field has been set.

### SetRemotePortNil

`func (o *MicrosoftGraphTeleconferenceDeviceMediaQuality) SetRemotePortNil(b bool)`

 SetRemotePortNil sets the value for RemotePort to be an explicit nil

### UnsetRemotePort
`func (o *MicrosoftGraphTeleconferenceDeviceMediaQuality) UnsetRemotePort()`

UnsetRemotePort ensures that no value is present for RemotePort, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


