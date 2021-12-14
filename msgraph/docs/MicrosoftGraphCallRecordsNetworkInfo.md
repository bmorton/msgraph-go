# MicrosoftGraphCallRecordsNetworkInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BandwidthLowEventRatio** | Pointer to [**AnyOfnumberstringstring**](anyOf&lt;number,string,string&gt;.md) | Fraction of the call that the media endpoint detected the available bandwidth or bandwidth policy was low enough to cause poor quality of the audio sent. | [optional] 
**BasicServiceSetIdentifier** | Pointer to **NullableString** | The wireless LAN basic service set identifier of the media endpoint used to connect to the network. | [optional] 
**ConnectionType** | Pointer to [**AnyOfmicrosoftGraphCallRecordsNetworkConnectionType**](anyOf&lt;microsoft.graph.callRecords.networkConnectionType&gt;.md) | Type of network used by the media endpoint. Possible values are: unknown, wired, wifi, mobile, tunnel, unknownFutureValue. | [optional] 
**DelayEventRatio** | Pointer to [**AnyOfnumberstringstring**](anyOf&lt;number,string,string&gt;.md) | Fraction of the call that the media endpoint detected the network delay was significant enough to impact the ability to have real-time two-way communication. | [optional] 
**DnsSuffix** | Pointer to **NullableString** | DNS suffix associated with the network adapter of the media endpoint. | [optional] 
**IpAddress** | Pointer to **NullableString** | IP address of the media endpoint. | [optional] 
**LinkSpeed** | Pointer to **NullableInt64** | Link speed in bits per second reported by the network adapter used by the media endpoint. | [optional] 
**MacAddress** | Pointer to **NullableString** | The media access control (MAC) address of the media endpoint&#39;s network device. | [optional] 
**Port** | Pointer to **NullableInt32** | Network port number used by media endpoint. | [optional] 
**ReceivedQualityEventRatio** | Pointer to [**AnyOfnumberstringstring**](anyOf&lt;number,string,string&gt;.md) | Fraction of the call that the media endpoint detected the network was causing poor quality of the audio received. | [optional] 
**ReflexiveIPAddress** | Pointer to **NullableString** | IP address of the media endpoint as seen by the media relay server. This is typically the public internet IP address associated to the endpoint. | [optional] 
**RelayIPAddress** | Pointer to **NullableString** | IP address of the media relay server allocated by the media endpoint. | [optional] 
**RelayPort** | Pointer to **NullableInt32** | Network port number allocated on the media relay server by the media endpoint. | [optional] 
**SentQualityEventRatio** | Pointer to [**AnyOfnumberstringstring**](anyOf&lt;number,string,string&gt;.md) | Fraction of the call that the media endpoint detected the network was causing poor quality of the audio sent. | [optional] 
**Subnet** | Pointer to **NullableString** | Subnet used for media stream by the media endpoint. | [optional] 
**WifiBand** | Pointer to [**AnyOfmicrosoftGraphCallRecordsWifiBand**](anyOf&lt;microsoft.graph.callRecords.wifiBand&gt;.md) | WiFi band used by the media endpoint. Possible values are: unknown, frequency24GHz, frequency50GHz, frequency60GHz, unknownFutureValue. | [optional] 
**WifiBatteryCharge** | Pointer to **NullableInt32** | Estimated remaining battery charge in percentage reported by the media endpoint. | [optional] 
**WifiChannel** | Pointer to **NullableInt32** | WiFi channel used by the media endpoint. | [optional] 
**WifiMicrosoftDriver** | Pointer to **NullableString** | Name of the Microsoft WiFi driver used by the media endpoint. Value may be localized based on the language used by endpoint. | [optional] 
**WifiMicrosoftDriverVersion** | Pointer to **NullableString** | Version of the Microsoft WiFi driver used by the media endpoint. | [optional] 
**WifiRadioType** | Pointer to [**AnyOfmicrosoftGraphCallRecordsWifiRadioType**](anyOf&lt;microsoft.graph.callRecords.wifiRadioType&gt;.md) | Type of WiFi radio used by the media endpoint. Possible values are: unknown, wifi80211a, wifi80211b, wifi80211g, wifi80211n, wifi80211ac, wifi80211ax, unknownFutureValue. | [optional] 
**WifiSignalStrength** | Pointer to **NullableInt32** | WiFi signal strength in percentage reported by the media endpoint. | [optional] 
**WifiVendorDriver** | Pointer to **NullableString** | Name of the WiFi driver used by the media endpoint. Value may be localized based on the language used by endpoint. | [optional] 
**WifiVendorDriverVersion** | Pointer to **NullableString** | Version of the WiFi driver used by the media endpoint. | [optional] 

## Methods

### NewMicrosoftGraphCallRecordsNetworkInfo

`func NewMicrosoftGraphCallRecordsNetworkInfo() *MicrosoftGraphCallRecordsNetworkInfo`

NewMicrosoftGraphCallRecordsNetworkInfo instantiates a new MicrosoftGraphCallRecordsNetworkInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphCallRecordsNetworkInfoWithDefaults

`func NewMicrosoftGraphCallRecordsNetworkInfoWithDefaults() *MicrosoftGraphCallRecordsNetworkInfo`

NewMicrosoftGraphCallRecordsNetworkInfoWithDefaults instantiates a new MicrosoftGraphCallRecordsNetworkInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBandwidthLowEventRatio

`func (o *MicrosoftGraphCallRecordsNetworkInfo) GetBandwidthLowEventRatio() AnyOfnumberstringstring`

GetBandwidthLowEventRatio returns the BandwidthLowEventRatio field if non-nil, zero value otherwise.

### GetBandwidthLowEventRatioOk

`func (o *MicrosoftGraphCallRecordsNetworkInfo) GetBandwidthLowEventRatioOk() (*AnyOfnumberstringstring, bool)`

GetBandwidthLowEventRatioOk returns a tuple with the BandwidthLowEventRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidthLowEventRatio

`func (o *MicrosoftGraphCallRecordsNetworkInfo) SetBandwidthLowEventRatio(v AnyOfnumberstringstring)`

SetBandwidthLowEventRatio sets BandwidthLowEventRatio field to given value.

### HasBandwidthLowEventRatio

`func (o *MicrosoftGraphCallRecordsNetworkInfo) HasBandwidthLowEventRatio() bool`

HasBandwidthLowEventRatio returns a boolean if a field has been set.

### SetBandwidthLowEventRatioNil

`func (o *MicrosoftGraphCallRecordsNetworkInfo) SetBandwidthLowEventRatioNil(b bool)`

 SetBandwidthLowEventRatioNil sets the value for BandwidthLowEventRatio to be an explicit nil

### UnsetBandwidthLowEventRatio
`func (o *MicrosoftGraphCallRecordsNetworkInfo) UnsetBandwidthLowEventRatio()`

UnsetBandwidthLowEventRatio ensures that no value is present for BandwidthLowEventRatio, not even an explicit nil
### GetBasicServiceSetIdentifier

`func (o *MicrosoftGraphCallRecordsNetworkInfo) GetBasicServiceSetIdentifier() string`

GetBasicServiceSetIdentifier returns the BasicServiceSetIdentifier field if non-nil, zero value otherwise.

### GetBasicServiceSetIdentifierOk

`func (o *MicrosoftGraphCallRecordsNetworkInfo) GetBasicServiceSetIdentifierOk() (*string, bool)`

GetBasicServiceSetIdentifierOk returns a tuple with the BasicServiceSetIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicServiceSetIdentifier

`func (o *MicrosoftGraphCallRecordsNetworkInfo) SetBasicServiceSetIdentifier(v string)`

SetBasicServiceSetIdentifier sets BasicServiceSetIdentifier field to given value.

### HasBasicServiceSetIdentifier

`func (o *MicrosoftGraphCallRecordsNetworkInfo) HasBasicServiceSetIdentifier() bool`

HasBasicServiceSetIdentifier returns a boolean if a field has been set.

### SetBasicServiceSetIdentifierNil

`func (o *MicrosoftGraphCallRecordsNetworkInfo) SetBasicServiceSetIdentifierNil(b bool)`

 SetBasicServiceSetIdentifierNil sets the value for BasicServiceSetIdentifier to be an explicit nil

### UnsetBasicServiceSetIdentifier
`func (o *MicrosoftGraphCallRecordsNetworkInfo) UnsetBasicServiceSetIdentifier()`

UnsetBasicServiceSetIdentifier ensures that no value is present for BasicServiceSetIdentifier, not even an explicit nil
### GetConnectionType

`func (o *MicrosoftGraphCallRecordsNetworkInfo) GetConnectionType() AnyOfmicrosoftGraphCallRecordsNetworkConnectionType`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *MicrosoftGraphCallRecordsNetworkInfo) GetConnectionTypeOk() (*AnyOfmicrosoftGraphCallRecordsNetworkConnectionType, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *MicrosoftGraphCallRecordsNetworkInfo) SetConnectionType(v AnyOfmicrosoftGraphCallRecordsNetworkConnectionType)`

SetConnectionType sets ConnectionType field to given value.

### HasConnectionType

`func (o *MicrosoftGraphCallRecordsNetworkInfo) HasConnectionType() bool`

HasConnectionType returns a boolean if a field has been set.

### SetConnectionTypeNil

`func (o *MicrosoftGraphCallRecordsNetworkInfo) SetConnectionTypeNil(b bool)`

 SetConnectionTypeNil sets the value for ConnectionType to be an explicit nil

### UnsetConnectionType
`func (o *MicrosoftGraphCallRecordsNetworkInfo) UnsetConnectionType()`

UnsetConnectionType ensures that no value is present for ConnectionType, not even an explicit nil
### GetDelayEventRatio

`func (o *MicrosoftGraphCallRecordsNetworkInfo) GetDelayEventRatio() AnyOfnumberstringstring`

GetDelayEventRatio returns the DelayEventRatio field if non-nil, zero value otherwise.

### GetDelayEventRatioOk

`func (o *MicrosoftGraphCallRecordsNetworkInfo) GetDelayEventRatioOk() (*AnyOfnumberstringstring, bool)`

GetDelayEventRatioOk returns a tuple with the DelayEventRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelayEventRatio

`func (o *MicrosoftGraphCallRecordsNetworkInfo) SetDelayEventRatio(v AnyOfnumberstringstring)`

SetDelayEventRatio sets DelayEventRatio field to given value.

### HasDelayEventRatio

`func (o *MicrosoftGraphCallRecordsNetworkInfo) HasDelayEventRatio() bool`

HasDelayEventRatio returns a boolean if a field has been set.

### SetDelayEventRatioNil

`func (o *MicrosoftGraphCallRecordsNetworkInfo) SetDelayEventRatioNil(b bool)`

 SetDelayEventRatioNil sets the value for DelayEventRatio to be an explicit nil

### UnsetDelayEventRatio
`func (o *MicrosoftGraphCallRecordsNetworkInfo) UnsetDelayEventRatio()`

UnsetDelayEventRatio ensures that no value is present for DelayEventRatio, not even an explicit nil
### GetDnsSuffix

`func (o *MicrosoftGraphCallRecordsNetworkInfo) GetDnsSuffix() string`

GetDnsSuffix returns the DnsSuffix field if non-nil, zero value otherwise.

### GetDnsSuffixOk

`func (o *MicrosoftGraphCallRecordsNetworkInfo) GetDnsSuffixOk() (*string, bool)`

GetDnsSuffixOk returns a tuple with the DnsSuffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsSuffix

`func (o *MicrosoftGraphCallRecordsNetworkInfo) SetDnsSuffix(v string)`

SetDnsSuffix sets DnsSuffix field to given value.

### HasDnsSuffix

`func (o *MicrosoftGraphCallRecordsNetworkInfo) HasDnsSuffix() bool`

HasDnsSuffix returns a boolean if a field has been set.

### SetDnsSuffixNil

`func (o *MicrosoftGraphCallRecordsNetworkInfo) SetDnsSuffixNil(b bool)`

 SetDnsSuffixNil sets the value for DnsSuffix to be an explicit nil

### UnsetDnsSuffix
`func (o *MicrosoftGraphCallRecordsNetworkInfo) UnsetDnsSuffix()`

UnsetDnsSuffix ensures that no value is present for DnsSuffix, not even an explicit nil
### GetIpAddress

`func (o *MicrosoftGraphCallRecordsNetworkInfo) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *MicrosoftGraphCallRecordsNetworkInfo) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *MicrosoftGraphCallRecordsNetworkInfo) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *MicrosoftGraphCallRecordsNetworkInfo) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### SetIpAddressNil

`func (o *MicrosoftGraphCallRecordsNetworkInfo) SetIpAddressNil(b bool)`

 SetIpAddressNil sets the value for IpAddress to be an explicit nil

### UnsetIpAddress
`func (o *MicrosoftGraphCallRecordsNetworkInfo) UnsetIpAddress()`

UnsetIpAddress ensures that no value is present for IpAddress, not even an explicit nil
### GetLinkSpeed

`func (o *MicrosoftGraphCallRecordsNetworkInfo) GetLinkSpeed() int64`

GetLinkSpeed returns the LinkSpeed field if non-nil, zero value otherwise.

### GetLinkSpeedOk

`func (o *MicrosoftGraphCallRecordsNetworkInfo) GetLinkSpeedOk() (*int64, bool)`

GetLinkSpeedOk returns a tuple with the LinkSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkSpeed

`func (o *MicrosoftGraphCallRecordsNetworkInfo) SetLinkSpeed(v int64)`

SetLinkSpeed sets LinkSpeed field to given value.

### HasLinkSpeed

`func (o *MicrosoftGraphCallRecordsNetworkInfo) HasLinkSpeed() bool`

HasLinkSpeed returns a boolean if a field has been set.

### SetLinkSpeedNil

`func (o *MicrosoftGraphCallRecordsNetworkInfo) SetLinkSpeedNil(b bool)`

 SetLinkSpeedNil sets the value for LinkSpeed to be an explicit nil

### UnsetLinkSpeed
`func (o *MicrosoftGraphCallRecordsNetworkInfo) UnsetLinkSpeed()`

UnsetLinkSpeed ensures that no value is present for LinkSpeed, not even an explicit nil
### GetMacAddress

`func (o *MicrosoftGraphCallRecordsNetworkInfo) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *MicrosoftGraphCallRecordsNetworkInfo) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *MicrosoftGraphCallRecordsNetworkInfo) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *MicrosoftGraphCallRecordsNetworkInfo) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### SetMacAddressNil

`func (o *MicrosoftGraphCallRecordsNetworkInfo) SetMacAddressNil(b bool)`

 SetMacAddressNil sets the value for MacAddress to be an explicit nil

### UnsetMacAddress
`func (o *MicrosoftGraphCallRecordsNetworkInfo) UnsetMacAddress()`

UnsetMacAddress ensures that no value is present for MacAddress, not even an explicit nil
### GetPort

`func (o *MicrosoftGraphCallRecordsNetworkInfo) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *MicrosoftGraphCallRecordsNetworkInfo) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *MicrosoftGraphCallRecordsNetworkInfo) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *MicrosoftGraphCallRecordsNetworkInfo) HasPort() bool`

HasPort returns a boolean if a field has been set.

### SetPortNil

`func (o *MicrosoftGraphCallRecordsNetworkInfo) SetPortNil(b bool)`

 SetPortNil sets the value for Port to be an explicit nil

### UnsetPort
`func (o *MicrosoftGraphCallRecordsNetworkInfo) UnsetPort()`

UnsetPort ensures that no value is present for Port, not even an explicit nil
### GetReceivedQualityEventRatio

`func (o *MicrosoftGraphCallRecordsNetworkInfo) GetReceivedQualityEventRatio() AnyOfnumberstringstring`

GetReceivedQualityEventRatio returns the ReceivedQualityEventRatio field if non-nil, zero value otherwise.

### GetReceivedQualityEventRatioOk

`func (o *MicrosoftGraphCallRecordsNetworkInfo) GetReceivedQualityEventRatioOk() (*AnyOfnumberstringstring, bool)`

GetReceivedQualityEventRatioOk returns a tuple with the ReceivedQualityEventRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivedQualityEventRatio

`func (o *MicrosoftGraphCallRecordsNetworkInfo) SetReceivedQualityEventRatio(v AnyOfnumberstringstring)`

SetReceivedQualityEventRatio sets ReceivedQualityEventRatio field to given value.

### HasReceivedQualityEventRatio

`func (o *MicrosoftGraphCallRecordsNetworkInfo) HasReceivedQualityEventRatio() bool`

HasReceivedQualityEventRatio returns a boolean if a field has been set.

### SetReceivedQualityEventRatioNil

`func (o *MicrosoftGraphCallRecordsNetworkInfo) SetReceivedQualityEventRatioNil(b bool)`

 SetReceivedQualityEventRatioNil sets the value for ReceivedQualityEventRatio to be an explicit nil

### UnsetReceivedQualityEventRatio
`func (o *MicrosoftGraphCallRecordsNetworkInfo) UnsetReceivedQualityEventRatio()`

UnsetReceivedQualityEventRatio ensures that no value is present for ReceivedQualityEventRatio, not even an explicit nil
### GetReflexiveIPAddress

`func (o *MicrosoftGraphCallRecordsNetworkInfo) GetReflexiveIPAddress() string`

GetReflexiveIPAddress returns the ReflexiveIPAddress field if non-nil, zero value otherwise.

### GetReflexiveIPAddressOk

`func (o *MicrosoftGraphCallRecordsNetworkInfo) GetReflexiveIPAddressOk() (*string, bool)`

GetReflexiveIPAddressOk returns a tuple with the ReflexiveIPAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReflexiveIPAddress

`func (o *MicrosoftGraphCallRecordsNetworkInfo) SetReflexiveIPAddress(v string)`

SetReflexiveIPAddress sets ReflexiveIPAddress field to given value.

### HasReflexiveIPAddress

`func (o *MicrosoftGraphCallRecordsNetworkInfo) HasReflexiveIPAddress() bool`

HasReflexiveIPAddress returns a boolean if a field has been set.

### SetReflexiveIPAddressNil

`func (o *MicrosoftGraphCallRecordsNetworkInfo) SetReflexiveIPAddressNil(b bool)`

 SetReflexiveIPAddressNil sets the value for ReflexiveIPAddress to be an explicit nil

### UnsetReflexiveIPAddress
`func (o *MicrosoftGraphCallRecordsNetworkInfo) UnsetReflexiveIPAddress()`

UnsetReflexiveIPAddress ensures that no value is present for ReflexiveIPAddress, not even an explicit nil
### GetRelayIPAddress

`func (o *MicrosoftGraphCallRecordsNetworkInfo) GetRelayIPAddress() string`

GetRelayIPAddress returns the RelayIPAddress field if non-nil, zero value otherwise.

### GetRelayIPAddressOk

`func (o *MicrosoftGraphCallRecordsNetworkInfo) GetRelayIPAddressOk() (*string, bool)`

GetRelayIPAddressOk returns a tuple with the RelayIPAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelayIPAddress

`func (o *MicrosoftGraphCallRecordsNetworkInfo) SetRelayIPAddress(v string)`

SetRelayIPAddress sets RelayIPAddress field to given value.

### HasRelayIPAddress

`func (o *MicrosoftGraphCallRecordsNetworkInfo) HasRelayIPAddress() bool`

HasRelayIPAddress returns a boolean if a field has been set.

### SetRelayIPAddressNil

`func (o *MicrosoftGraphCallRecordsNetworkInfo) SetRelayIPAddressNil(b bool)`

 SetRelayIPAddressNil sets the value for RelayIPAddress to be an explicit nil

### UnsetRelayIPAddress
`func (o *MicrosoftGraphCallRecordsNetworkInfo) UnsetRelayIPAddress()`

UnsetRelayIPAddress ensures that no value is present for RelayIPAddress, not even an explicit nil
### GetRelayPort

`func (o *MicrosoftGraphCallRecordsNetworkInfo) GetRelayPort() int32`

GetRelayPort returns the RelayPort field if non-nil, zero value otherwise.

### GetRelayPortOk

`func (o *MicrosoftGraphCallRecordsNetworkInfo) GetRelayPortOk() (*int32, bool)`

GetRelayPortOk returns a tuple with the RelayPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelayPort

`func (o *MicrosoftGraphCallRecordsNetworkInfo) SetRelayPort(v int32)`

SetRelayPort sets RelayPort field to given value.

### HasRelayPort

`func (o *MicrosoftGraphCallRecordsNetworkInfo) HasRelayPort() bool`

HasRelayPort returns a boolean if a field has been set.

### SetRelayPortNil

`func (o *MicrosoftGraphCallRecordsNetworkInfo) SetRelayPortNil(b bool)`

 SetRelayPortNil sets the value for RelayPort to be an explicit nil

### UnsetRelayPort
`func (o *MicrosoftGraphCallRecordsNetworkInfo) UnsetRelayPort()`

UnsetRelayPort ensures that no value is present for RelayPort, not even an explicit nil
### GetSentQualityEventRatio

`func (o *MicrosoftGraphCallRecordsNetworkInfo) GetSentQualityEventRatio() AnyOfnumberstringstring`

GetSentQualityEventRatio returns the SentQualityEventRatio field if non-nil, zero value otherwise.

### GetSentQualityEventRatioOk

`func (o *MicrosoftGraphCallRecordsNetworkInfo) GetSentQualityEventRatioOk() (*AnyOfnumberstringstring, bool)`

GetSentQualityEventRatioOk returns a tuple with the SentQualityEventRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSentQualityEventRatio

`func (o *MicrosoftGraphCallRecordsNetworkInfo) SetSentQualityEventRatio(v AnyOfnumberstringstring)`

SetSentQualityEventRatio sets SentQualityEventRatio field to given value.

### HasSentQualityEventRatio

`func (o *MicrosoftGraphCallRecordsNetworkInfo) HasSentQualityEventRatio() bool`

HasSentQualityEventRatio returns a boolean if a field has been set.

### SetSentQualityEventRatioNil

`func (o *MicrosoftGraphCallRecordsNetworkInfo) SetSentQualityEventRatioNil(b bool)`

 SetSentQualityEventRatioNil sets the value for SentQualityEventRatio to be an explicit nil

### UnsetSentQualityEventRatio
`func (o *MicrosoftGraphCallRecordsNetworkInfo) UnsetSentQualityEventRatio()`

UnsetSentQualityEventRatio ensures that no value is present for SentQualityEventRatio, not even an explicit nil
### GetSubnet

`func (o *MicrosoftGraphCallRecordsNetworkInfo) GetSubnet() string`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *MicrosoftGraphCallRecordsNetworkInfo) GetSubnetOk() (*string, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *MicrosoftGraphCallRecordsNetworkInfo) SetSubnet(v string)`

SetSubnet sets Subnet field to given value.

### HasSubnet

`func (o *MicrosoftGraphCallRecordsNetworkInfo) HasSubnet() bool`

HasSubnet returns a boolean if a field has been set.

### SetSubnetNil

`func (o *MicrosoftGraphCallRecordsNetworkInfo) SetSubnetNil(b bool)`

 SetSubnetNil sets the value for Subnet to be an explicit nil

### UnsetSubnet
`func (o *MicrosoftGraphCallRecordsNetworkInfo) UnsetSubnet()`

UnsetSubnet ensures that no value is present for Subnet, not even an explicit nil
### GetWifiBand

`func (o *MicrosoftGraphCallRecordsNetworkInfo) GetWifiBand() AnyOfmicrosoftGraphCallRecordsWifiBand`

GetWifiBand returns the WifiBand field if non-nil, zero value otherwise.

### GetWifiBandOk

`func (o *MicrosoftGraphCallRecordsNetworkInfo) GetWifiBandOk() (*AnyOfmicrosoftGraphCallRecordsWifiBand, bool)`

GetWifiBandOk returns a tuple with the WifiBand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWifiBand

`func (o *MicrosoftGraphCallRecordsNetworkInfo) SetWifiBand(v AnyOfmicrosoftGraphCallRecordsWifiBand)`

SetWifiBand sets WifiBand field to given value.

### HasWifiBand

`func (o *MicrosoftGraphCallRecordsNetworkInfo) HasWifiBand() bool`

HasWifiBand returns a boolean if a field has been set.

### SetWifiBandNil

`func (o *MicrosoftGraphCallRecordsNetworkInfo) SetWifiBandNil(b bool)`

 SetWifiBandNil sets the value for WifiBand to be an explicit nil

### UnsetWifiBand
`func (o *MicrosoftGraphCallRecordsNetworkInfo) UnsetWifiBand()`

UnsetWifiBand ensures that no value is present for WifiBand, not even an explicit nil
### GetWifiBatteryCharge

`func (o *MicrosoftGraphCallRecordsNetworkInfo) GetWifiBatteryCharge() int32`

GetWifiBatteryCharge returns the WifiBatteryCharge field if non-nil, zero value otherwise.

### GetWifiBatteryChargeOk

`func (o *MicrosoftGraphCallRecordsNetworkInfo) GetWifiBatteryChargeOk() (*int32, bool)`

GetWifiBatteryChargeOk returns a tuple with the WifiBatteryCharge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWifiBatteryCharge

`func (o *MicrosoftGraphCallRecordsNetworkInfo) SetWifiBatteryCharge(v int32)`

SetWifiBatteryCharge sets WifiBatteryCharge field to given value.

### HasWifiBatteryCharge

`func (o *MicrosoftGraphCallRecordsNetworkInfo) HasWifiBatteryCharge() bool`

HasWifiBatteryCharge returns a boolean if a field has been set.

### SetWifiBatteryChargeNil

`func (o *MicrosoftGraphCallRecordsNetworkInfo) SetWifiBatteryChargeNil(b bool)`

 SetWifiBatteryChargeNil sets the value for WifiBatteryCharge to be an explicit nil

### UnsetWifiBatteryCharge
`func (o *MicrosoftGraphCallRecordsNetworkInfo) UnsetWifiBatteryCharge()`

UnsetWifiBatteryCharge ensures that no value is present for WifiBatteryCharge, not even an explicit nil
### GetWifiChannel

`func (o *MicrosoftGraphCallRecordsNetworkInfo) GetWifiChannel() int32`

GetWifiChannel returns the WifiChannel field if non-nil, zero value otherwise.

### GetWifiChannelOk

`func (o *MicrosoftGraphCallRecordsNetworkInfo) GetWifiChannelOk() (*int32, bool)`

GetWifiChannelOk returns a tuple with the WifiChannel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWifiChannel

`func (o *MicrosoftGraphCallRecordsNetworkInfo) SetWifiChannel(v int32)`

SetWifiChannel sets WifiChannel field to given value.

### HasWifiChannel

`func (o *MicrosoftGraphCallRecordsNetworkInfo) HasWifiChannel() bool`

HasWifiChannel returns a boolean if a field has been set.

### SetWifiChannelNil

`func (o *MicrosoftGraphCallRecordsNetworkInfo) SetWifiChannelNil(b bool)`

 SetWifiChannelNil sets the value for WifiChannel to be an explicit nil

### UnsetWifiChannel
`func (o *MicrosoftGraphCallRecordsNetworkInfo) UnsetWifiChannel()`

UnsetWifiChannel ensures that no value is present for WifiChannel, not even an explicit nil
### GetWifiMicrosoftDriver

`func (o *MicrosoftGraphCallRecordsNetworkInfo) GetWifiMicrosoftDriver() string`

GetWifiMicrosoftDriver returns the WifiMicrosoftDriver field if non-nil, zero value otherwise.

### GetWifiMicrosoftDriverOk

`func (o *MicrosoftGraphCallRecordsNetworkInfo) GetWifiMicrosoftDriverOk() (*string, bool)`

GetWifiMicrosoftDriverOk returns a tuple with the WifiMicrosoftDriver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWifiMicrosoftDriver

`func (o *MicrosoftGraphCallRecordsNetworkInfo) SetWifiMicrosoftDriver(v string)`

SetWifiMicrosoftDriver sets WifiMicrosoftDriver field to given value.

### HasWifiMicrosoftDriver

`func (o *MicrosoftGraphCallRecordsNetworkInfo) HasWifiMicrosoftDriver() bool`

HasWifiMicrosoftDriver returns a boolean if a field has been set.

### SetWifiMicrosoftDriverNil

`func (o *MicrosoftGraphCallRecordsNetworkInfo) SetWifiMicrosoftDriverNil(b bool)`

 SetWifiMicrosoftDriverNil sets the value for WifiMicrosoftDriver to be an explicit nil

### UnsetWifiMicrosoftDriver
`func (o *MicrosoftGraphCallRecordsNetworkInfo) UnsetWifiMicrosoftDriver()`

UnsetWifiMicrosoftDriver ensures that no value is present for WifiMicrosoftDriver, not even an explicit nil
### GetWifiMicrosoftDriverVersion

`func (o *MicrosoftGraphCallRecordsNetworkInfo) GetWifiMicrosoftDriverVersion() string`

GetWifiMicrosoftDriverVersion returns the WifiMicrosoftDriverVersion field if non-nil, zero value otherwise.

### GetWifiMicrosoftDriverVersionOk

`func (o *MicrosoftGraphCallRecordsNetworkInfo) GetWifiMicrosoftDriverVersionOk() (*string, bool)`

GetWifiMicrosoftDriverVersionOk returns a tuple with the WifiMicrosoftDriverVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWifiMicrosoftDriverVersion

`func (o *MicrosoftGraphCallRecordsNetworkInfo) SetWifiMicrosoftDriverVersion(v string)`

SetWifiMicrosoftDriverVersion sets WifiMicrosoftDriverVersion field to given value.

### HasWifiMicrosoftDriverVersion

`func (o *MicrosoftGraphCallRecordsNetworkInfo) HasWifiMicrosoftDriverVersion() bool`

HasWifiMicrosoftDriverVersion returns a boolean if a field has been set.

### SetWifiMicrosoftDriverVersionNil

`func (o *MicrosoftGraphCallRecordsNetworkInfo) SetWifiMicrosoftDriverVersionNil(b bool)`

 SetWifiMicrosoftDriverVersionNil sets the value for WifiMicrosoftDriverVersion to be an explicit nil

### UnsetWifiMicrosoftDriverVersion
`func (o *MicrosoftGraphCallRecordsNetworkInfo) UnsetWifiMicrosoftDriverVersion()`

UnsetWifiMicrosoftDriverVersion ensures that no value is present for WifiMicrosoftDriverVersion, not even an explicit nil
### GetWifiRadioType

`func (o *MicrosoftGraphCallRecordsNetworkInfo) GetWifiRadioType() AnyOfmicrosoftGraphCallRecordsWifiRadioType`

GetWifiRadioType returns the WifiRadioType field if non-nil, zero value otherwise.

### GetWifiRadioTypeOk

`func (o *MicrosoftGraphCallRecordsNetworkInfo) GetWifiRadioTypeOk() (*AnyOfmicrosoftGraphCallRecordsWifiRadioType, bool)`

GetWifiRadioTypeOk returns a tuple with the WifiRadioType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWifiRadioType

`func (o *MicrosoftGraphCallRecordsNetworkInfo) SetWifiRadioType(v AnyOfmicrosoftGraphCallRecordsWifiRadioType)`

SetWifiRadioType sets WifiRadioType field to given value.

### HasWifiRadioType

`func (o *MicrosoftGraphCallRecordsNetworkInfo) HasWifiRadioType() bool`

HasWifiRadioType returns a boolean if a field has been set.

### SetWifiRadioTypeNil

`func (o *MicrosoftGraphCallRecordsNetworkInfo) SetWifiRadioTypeNil(b bool)`

 SetWifiRadioTypeNil sets the value for WifiRadioType to be an explicit nil

### UnsetWifiRadioType
`func (o *MicrosoftGraphCallRecordsNetworkInfo) UnsetWifiRadioType()`

UnsetWifiRadioType ensures that no value is present for WifiRadioType, not even an explicit nil
### GetWifiSignalStrength

`func (o *MicrosoftGraphCallRecordsNetworkInfo) GetWifiSignalStrength() int32`

GetWifiSignalStrength returns the WifiSignalStrength field if non-nil, zero value otherwise.

### GetWifiSignalStrengthOk

`func (o *MicrosoftGraphCallRecordsNetworkInfo) GetWifiSignalStrengthOk() (*int32, bool)`

GetWifiSignalStrengthOk returns a tuple with the WifiSignalStrength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWifiSignalStrength

`func (o *MicrosoftGraphCallRecordsNetworkInfo) SetWifiSignalStrength(v int32)`

SetWifiSignalStrength sets WifiSignalStrength field to given value.

### HasWifiSignalStrength

`func (o *MicrosoftGraphCallRecordsNetworkInfo) HasWifiSignalStrength() bool`

HasWifiSignalStrength returns a boolean if a field has been set.

### SetWifiSignalStrengthNil

`func (o *MicrosoftGraphCallRecordsNetworkInfo) SetWifiSignalStrengthNil(b bool)`

 SetWifiSignalStrengthNil sets the value for WifiSignalStrength to be an explicit nil

### UnsetWifiSignalStrength
`func (o *MicrosoftGraphCallRecordsNetworkInfo) UnsetWifiSignalStrength()`

UnsetWifiSignalStrength ensures that no value is present for WifiSignalStrength, not even an explicit nil
### GetWifiVendorDriver

`func (o *MicrosoftGraphCallRecordsNetworkInfo) GetWifiVendorDriver() string`

GetWifiVendorDriver returns the WifiVendorDriver field if non-nil, zero value otherwise.

### GetWifiVendorDriverOk

`func (o *MicrosoftGraphCallRecordsNetworkInfo) GetWifiVendorDriverOk() (*string, bool)`

GetWifiVendorDriverOk returns a tuple with the WifiVendorDriver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWifiVendorDriver

`func (o *MicrosoftGraphCallRecordsNetworkInfo) SetWifiVendorDriver(v string)`

SetWifiVendorDriver sets WifiVendorDriver field to given value.

### HasWifiVendorDriver

`func (o *MicrosoftGraphCallRecordsNetworkInfo) HasWifiVendorDriver() bool`

HasWifiVendorDriver returns a boolean if a field has been set.

### SetWifiVendorDriverNil

`func (o *MicrosoftGraphCallRecordsNetworkInfo) SetWifiVendorDriverNil(b bool)`

 SetWifiVendorDriverNil sets the value for WifiVendorDriver to be an explicit nil

### UnsetWifiVendorDriver
`func (o *MicrosoftGraphCallRecordsNetworkInfo) UnsetWifiVendorDriver()`

UnsetWifiVendorDriver ensures that no value is present for WifiVendorDriver, not even an explicit nil
### GetWifiVendorDriverVersion

`func (o *MicrosoftGraphCallRecordsNetworkInfo) GetWifiVendorDriverVersion() string`

GetWifiVendorDriverVersion returns the WifiVendorDriverVersion field if non-nil, zero value otherwise.

### GetWifiVendorDriverVersionOk

`func (o *MicrosoftGraphCallRecordsNetworkInfo) GetWifiVendorDriverVersionOk() (*string, bool)`

GetWifiVendorDriverVersionOk returns a tuple with the WifiVendorDriverVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWifiVendorDriverVersion

`func (o *MicrosoftGraphCallRecordsNetworkInfo) SetWifiVendorDriverVersion(v string)`

SetWifiVendorDriverVersion sets WifiVendorDriverVersion field to given value.

### HasWifiVendorDriverVersion

`func (o *MicrosoftGraphCallRecordsNetworkInfo) HasWifiVendorDriverVersion() bool`

HasWifiVendorDriverVersion returns a boolean if a field has been set.

### SetWifiVendorDriverVersionNil

`func (o *MicrosoftGraphCallRecordsNetworkInfo) SetWifiVendorDriverVersionNil(b bool)`

 SetWifiVendorDriverVersionNil sets the value for WifiVendorDriverVersion to be an explicit nil

### UnsetWifiVendorDriverVersion
`func (o *MicrosoftGraphCallRecordsNetworkInfo) UnsetWifiVendorDriverVersion()`

UnsetWifiVendorDriverVersion ensures that no value is present for WifiVendorDriverVersion, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


