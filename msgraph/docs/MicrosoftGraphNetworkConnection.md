# MicrosoftGraphNetworkConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationName** | Pointer to **NullableString** | Name of the application managing the network connection (for example, Facebook or SMTP). | [optional] 
**DestinationAddress** | Pointer to **NullableString** | Destination IP address (of the network connection). | [optional] 
**DestinationDomain** | Pointer to **NullableString** | Destination domain portion of the destination URL. (for example &#39;www.contoso.com&#39;). | [optional] 
**DestinationLocation** | Pointer to **NullableString** | Location (by IP address mapping) associated with the destination of a network connection. | [optional] 
**DestinationPort** | Pointer to **NullableString** | Destination port (of the network connection). | [optional] 
**DestinationUrl** | Pointer to **NullableString** | Network connection URL/URI string - excluding parameters. (for example &#39;www.contoso.com/products/default.html&#39;) | [optional] 
**Direction** | Pointer to [**AnyOfmicrosoftGraphConnectionDirection**](anyOf&lt;microsoft.graph.connectionDirection&gt;.md) | Network connection direction. Possible values are: unknown, inbound, outbound. | [optional] 
**DomainRegisteredDateTime** | Pointer to **NullableTime** | Date when the destination domain was registered. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z | [optional] 
**LocalDnsName** | Pointer to **NullableString** | The local DNS name resolution as it appears in the host&#39;s local DNS cache (for example, in case the &#39;hosts&#39; file was tampered with). | [optional] 
**NatDestinationAddress** | Pointer to **NullableString** | Network Address Translation destination IP address. | [optional] 
**NatDestinationPort** | Pointer to **NullableString** | Network Address Translation destination port. | [optional] 
**NatSourceAddress** | Pointer to **NullableString** | Network Address Translation source IP address. | [optional] 
**NatSourcePort** | Pointer to **NullableString** | Network Address Translation source port. | [optional] 
**Protocol** | Pointer to [**AnyOfmicrosoftGraphSecurityNetworkProtocol**](anyOf&lt;microsoft.graph.securityNetworkProtocol&gt;.md) | Network protocol. Possible values are: unknown, ip, icmp, igmp, ggp, ipv4, tcp, pup, udp, idp, ipv6, ipv6RoutingHeader, ipv6FragmentHeader, ipSecEncapsulatingSecurityPayload, ipSecAuthenticationHeader, icmpV6, ipv6NoNextHeader, ipv6DestinationOptions, nd, raw, ipx, spx, spxII. | [optional] 
**RiskScore** | Pointer to **NullableString** | Provider generated/calculated risk score of the network connection. Recommended value range of 0-1, which equates to a percentage. | [optional] 
**SourceAddress** | Pointer to **NullableString** | Source (i.e. origin) IP address (of the network connection). | [optional] 
**SourceLocation** | Pointer to **NullableString** | Location (by IP address mapping) associated with the source of a network connection. | [optional] 
**SourcePort** | Pointer to **NullableString** | Source (i.e. origin) IP port (of the network connection). | [optional] 
**Status** | Pointer to [**AnyOfmicrosoftGraphConnectionStatus**](anyOf&lt;microsoft.graph.connectionStatus&gt;.md) | Network connection status. Possible values are: unknown, attempted, succeeded, blocked, failed. | [optional] 
**UrlParameters** | Pointer to **NullableString** | Parameters (suffix) of the destination URL. | [optional] 

## Methods

### NewMicrosoftGraphNetworkConnection

`func NewMicrosoftGraphNetworkConnection() *MicrosoftGraphNetworkConnection`

NewMicrosoftGraphNetworkConnection instantiates a new MicrosoftGraphNetworkConnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphNetworkConnectionWithDefaults

`func NewMicrosoftGraphNetworkConnectionWithDefaults() *MicrosoftGraphNetworkConnection`

NewMicrosoftGraphNetworkConnectionWithDefaults instantiates a new MicrosoftGraphNetworkConnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationName

`func (o *MicrosoftGraphNetworkConnection) GetApplicationName() string`

GetApplicationName returns the ApplicationName field if non-nil, zero value otherwise.

### GetApplicationNameOk

`func (o *MicrosoftGraphNetworkConnection) GetApplicationNameOk() (*string, bool)`

GetApplicationNameOk returns a tuple with the ApplicationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationName

`func (o *MicrosoftGraphNetworkConnection) SetApplicationName(v string)`

SetApplicationName sets ApplicationName field to given value.

### HasApplicationName

`func (o *MicrosoftGraphNetworkConnection) HasApplicationName() bool`

HasApplicationName returns a boolean if a field has been set.

### SetApplicationNameNil

`func (o *MicrosoftGraphNetworkConnection) SetApplicationNameNil(b bool)`

 SetApplicationNameNil sets the value for ApplicationName to be an explicit nil

### UnsetApplicationName
`func (o *MicrosoftGraphNetworkConnection) UnsetApplicationName()`

UnsetApplicationName ensures that no value is present for ApplicationName, not even an explicit nil
### GetDestinationAddress

`func (o *MicrosoftGraphNetworkConnection) GetDestinationAddress() string`

GetDestinationAddress returns the DestinationAddress field if non-nil, zero value otherwise.

### GetDestinationAddressOk

`func (o *MicrosoftGraphNetworkConnection) GetDestinationAddressOk() (*string, bool)`

GetDestinationAddressOk returns a tuple with the DestinationAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationAddress

`func (o *MicrosoftGraphNetworkConnection) SetDestinationAddress(v string)`

SetDestinationAddress sets DestinationAddress field to given value.

### HasDestinationAddress

`func (o *MicrosoftGraphNetworkConnection) HasDestinationAddress() bool`

HasDestinationAddress returns a boolean if a field has been set.

### SetDestinationAddressNil

`func (o *MicrosoftGraphNetworkConnection) SetDestinationAddressNil(b bool)`

 SetDestinationAddressNil sets the value for DestinationAddress to be an explicit nil

### UnsetDestinationAddress
`func (o *MicrosoftGraphNetworkConnection) UnsetDestinationAddress()`

UnsetDestinationAddress ensures that no value is present for DestinationAddress, not even an explicit nil
### GetDestinationDomain

`func (o *MicrosoftGraphNetworkConnection) GetDestinationDomain() string`

GetDestinationDomain returns the DestinationDomain field if non-nil, zero value otherwise.

### GetDestinationDomainOk

`func (o *MicrosoftGraphNetworkConnection) GetDestinationDomainOk() (*string, bool)`

GetDestinationDomainOk returns a tuple with the DestinationDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationDomain

`func (o *MicrosoftGraphNetworkConnection) SetDestinationDomain(v string)`

SetDestinationDomain sets DestinationDomain field to given value.

### HasDestinationDomain

`func (o *MicrosoftGraphNetworkConnection) HasDestinationDomain() bool`

HasDestinationDomain returns a boolean if a field has been set.

### SetDestinationDomainNil

`func (o *MicrosoftGraphNetworkConnection) SetDestinationDomainNil(b bool)`

 SetDestinationDomainNil sets the value for DestinationDomain to be an explicit nil

### UnsetDestinationDomain
`func (o *MicrosoftGraphNetworkConnection) UnsetDestinationDomain()`

UnsetDestinationDomain ensures that no value is present for DestinationDomain, not even an explicit nil
### GetDestinationLocation

`func (o *MicrosoftGraphNetworkConnection) GetDestinationLocation() string`

GetDestinationLocation returns the DestinationLocation field if non-nil, zero value otherwise.

### GetDestinationLocationOk

`func (o *MicrosoftGraphNetworkConnection) GetDestinationLocationOk() (*string, bool)`

GetDestinationLocationOk returns a tuple with the DestinationLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationLocation

`func (o *MicrosoftGraphNetworkConnection) SetDestinationLocation(v string)`

SetDestinationLocation sets DestinationLocation field to given value.

### HasDestinationLocation

`func (o *MicrosoftGraphNetworkConnection) HasDestinationLocation() bool`

HasDestinationLocation returns a boolean if a field has been set.

### SetDestinationLocationNil

`func (o *MicrosoftGraphNetworkConnection) SetDestinationLocationNil(b bool)`

 SetDestinationLocationNil sets the value for DestinationLocation to be an explicit nil

### UnsetDestinationLocation
`func (o *MicrosoftGraphNetworkConnection) UnsetDestinationLocation()`

UnsetDestinationLocation ensures that no value is present for DestinationLocation, not even an explicit nil
### GetDestinationPort

`func (o *MicrosoftGraphNetworkConnection) GetDestinationPort() string`

GetDestinationPort returns the DestinationPort field if non-nil, zero value otherwise.

### GetDestinationPortOk

`func (o *MicrosoftGraphNetworkConnection) GetDestinationPortOk() (*string, bool)`

GetDestinationPortOk returns a tuple with the DestinationPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationPort

`func (o *MicrosoftGraphNetworkConnection) SetDestinationPort(v string)`

SetDestinationPort sets DestinationPort field to given value.

### HasDestinationPort

`func (o *MicrosoftGraphNetworkConnection) HasDestinationPort() bool`

HasDestinationPort returns a boolean if a field has been set.

### SetDestinationPortNil

`func (o *MicrosoftGraphNetworkConnection) SetDestinationPortNil(b bool)`

 SetDestinationPortNil sets the value for DestinationPort to be an explicit nil

### UnsetDestinationPort
`func (o *MicrosoftGraphNetworkConnection) UnsetDestinationPort()`

UnsetDestinationPort ensures that no value is present for DestinationPort, not even an explicit nil
### GetDestinationUrl

`func (o *MicrosoftGraphNetworkConnection) GetDestinationUrl() string`

GetDestinationUrl returns the DestinationUrl field if non-nil, zero value otherwise.

### GetDestinationUrlOk

`func (o *MicrosoftGraphNetworkConnection) GetDestinationUrlOk() (*string, bool)`

GetDestinationUrlOk returns a tuple with the DestinationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationUrl

`func (o *MicrosoftGraphNetworkConnection) SetDestinationUrl(v string)`

SetDestinationUrl sets DestinationUrl field to given value.

### HasDestinationUrl

`func (o *MicrosoftGraphNetworkConnection) HasDestinationUrl() bool`

HasDestinationUrl returns a boolean if a field has been set.

### SetDestinationUrlNil

`func (o *MicrosoftGraphNetworkConnection) SetDestinationUrlNil(b bool)`

 SetDestinationUrlNil sets the value for DestinationUrl to be an explicit nil

### UnsetDestinationUrl
`func (o *MicrosoftGraphNetworkConnection) UnsetDestinationUrl()`

UnsetDestinationUrl ensures that no value is present for DestinationUrl, not even an explicit nil
### GetDirection

`func (o *MicrosoftGraphNetworkConnection) GetDirection() AnyOfmicrosoftGraphConnectionDirection`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *MicrosoftGraphNetworkConnection) GetDirectionOk() (*AnyOfmicrosoftGraphConnectionDirection, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *MicrosoftGraphNetworkConnection) SetDirection(v AnyOfmicrosoftGraphConnectionDirection)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *MicrosoftGraphNetworkConnection) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### SetDirectionNil

`func (o *MicrosoftGraphNetworkConnection) SetDirectionNil(b bool)`

 SetDirectionNil sets the value for Direction to be an explicit nil

### UnsetDirection
`func (o *MicrosoftGraphNetworkConnection) UnsetDirection()`

UnsetDirection ensures that no value is present for Direction, not even an explicit nil
### GetDomainRegisteredDateTime

`func (o *MicrosoftGraphNetworkConnection) GetDomainRegisteredDateTime() time.Time`

GetDomainRegisteredDateTime returns the DomainRegisteredDateTime field if non-nil, zero value otherwise.

### GetDomainRegisteredDateTimeOk

`func (o *MicrosoftGraphNetworkConnection) GetDomainRegisteredDateTimeOk() (*time.Time, bool)`

GetDomainRegisteredDateTimeOk returns a tuple with the DomainRegisteredDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainRegisteredDateTime

`func (o *MicrosoftGraphNetworkConnection) SetDomainRegisteredDateTime(v time.Time)`

SetDomainRegisteredDateTime sets DomainRegisteredDateTime field to given value.

### HasDomainRegisteredDateTime

`func (o *MicrosoftGraphNetworkConnection) HasDomainRegisteredDateTime() bool`

HasDomainRegisteredDateTime returns a boolean if a field has been set.

### SetDomainRegisteredDateTimeNil

`func (o *MicrosoftGraphNetworkConnection) SetDomainRegisteredDateTimeNil(b bool)`

 SetDomainRegisteredDateTimeNil sets the value for DomainRegisteredDateTime to be an explicit nil

### UnsetDomainRegisteredDateTime
`func (o *MicrosoftGraphNetworkConnection) UnsetDomainRegisteredDateTime()`

UnsetDomainRegisteredDateTime ensures that no value is present for DomainRegisteredDateTime, not even an explicit nil
### GetLocalDnsName

`func (o *MicrosoftGraphNetworkConnection) GetLocalDnsName() string`

GetLocalDnsName returns the LocalDnsName field if non-nil, zero value otherwise.

### GetLocalDnsNameOk

`func (o *MicrosoftGraphNetworkConnection) GetLocalDnsNameOk() (*string, bool)`

GetLocalDnsNameOk returns a tuple with the LocalDnsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalDnsName

`func (o *MicrosoftGraphNetworkConnection) SetLocalDnsName(v string)`

SetLocalDnsName sets LocalDnsName field to given value.

### HasLocalDnsName

`func (o *MicrosoftGraphNetworkConnection) HasLocalDnsName() bool`

HasLocalDnsName returns a boolean if a field has been set.

### SetLocalDnsNameNil

`func (o *MicrosoftGraphNetworkConnection) SetLocalDnsNameNil(b bool)`

 SetLocalDnsNameNil sets the value for LocalDnsName to be an explicit nil

### UnsetLocalDnsName
`func (o *MicrosoftGraphNetworkConnection) UnsetLocalDnsName()`

UnsetLocalDnsName ensures that no value is present for LocalDnsName, not even an explicit nil
### GetNatDestinationAddress

`func (o *MicrosoftGraphNetworkConnection) GetNatDestinationAddress() string`

GetNatDestinationAddress returns the NatDestinationAddress field if non-nil, zero value otherwise.

### GetNatDestinationAddressOk

`func (o *MicrosoftGraphNetworkConnection) GetNatDestinationAddressOk() (*string, bool)`

GetNatDestinationAddressOk returns a tuple with the NatDestinationAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNatDestinationAddress

`func (o *MicrosoftGraphNetworkConnection) SetNatDestinationAddress(v string)`

SetNatDestinationAddress sets NatDestinationAddress field to given value.

### HasNatDestinationAddress

`func (o *MicrosoftGraphNetworkConnection) HasNatDestinationAddress() bool`

HasNatDestinationAddress returns a boolean if a field has been set.

### SetNatDestinationAddressNil

`func (o *MicrosoftGraphNetworkConnection) SetNatDestinationAddressNil(b bool)`

 SetNatDestinationAddressNil sets the value for NatDestinationAddress to be an explicit nil

### UnsetNatDestinationAddress
`func (o *MicrosoftGraphNetworkConnection) UnsetNatDestinationAddress()`

UnsetNatDestinationAddress ensures that no value is present for NatDestinationAddress, not even an explicit nil
### GetNatDestinationPort

`func (o *MicrosoftGraphNetworkConnection) GetNatDestinationPort() string`

GetNatDestinationPort returns the NatDestinationPort field if non-nil, zero value otherwise.

### GetNatDestinationPortOk

`func (o *MicrosoftGraphNetworkConnection) GetNatDestinationPortOk() (*string, bool)`

GetNatDestinationPortOk returns a tuple with the NatDestinationPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNatDestinationPort

`func (o *MicrosoftGraphNetworkConnection) SetNatDestinationPort(v string)`

SetNatDestinationPort sets NatDestinationPort field to given value.

### HasNatDestinationPort

`func (o *MicrosoftGraphNetworkConnection) HasNatDestinationPort() bool`

HasNatDestinationPort returns a boolean if a field has been set.

### SetNatDestinationPortNil

`func (o *MicrosoftGraphNetworkConnection) SetNatDestinationPortNil(b bool)`

 SetNatDestinationPortNil sets the value for NatDestinationPort to be an explicit nil

### UnsetNatDestinationPort
`func (o *MicrosoftGraphNetworkConnection) UnsetNatDestinationPort()`

UnsetNatDestinationPort ensures that no value is present for NatDestinationPort, not even an explicit nil
### GetNatSourceAddress

`func (o *MicrosoftGraphNetworkConnection) GetNatSourceAddress() string`

GetNatSourceAddress returns the NatSourceAddress field if non-nil, zero value otherwise.

### GetNatSourceAddressOk

`func (o *MicrosoftGraphNetworkConnection) GetNatSourceAddressOk() (*string, bool)`

GetNatSourceAddressOk returns a tuple with the NatSourceAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNatSourceAddress

`func (o *MicrosoftGraphNetworkConnection) SetNatSourceAddress(v string)`

SetNatSourceAddress sets NatSourceAddress field to given value.

### HasNatSourceAddress

`func (o *MicrosoftGraphNetworkConnection) HasNatSourceAddress() bool`

HasNatSourceAddress returns a boolean if a field has been set.

### SetNatSourceAddressNil

`func (o *MicrosoftGraphNetworkConnection) SetNatSourceAddressNil(b bool)`

 SetNatSourceAddressNil sets the value for NatSourceAddress to be an explicit nil

### UnsetNatSourceAddress
`func (o *MicrosoftGraphNetworkConnection) UnsetNatSourceAddress()`

UnsetNatSourceAddress ensures that no value is present for NatSourceAddress, not even an explicit nil
### GetNatSourcePort

`func (o *MicrosoftGraphNetworkConnection) GetNatSourcePort() string`

GetNatSourcePort returns the NatSourcePort field if non-nil, zero value otherwise.

### GetNatSourcePortOk

`func (o *MicrosoftGraphNetworkConnection) GetNatSourcePortOk() (*string, bool)`

GetNatSourcePortOk returns a tuple with the NatSourcePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNatSourcePort

`func (o *MicrosoftGraphNetworkConnection) SetNatSourcePort(v string)`

SetNatSourcePort sets NatSourcePort field to given value.

### HasNatSourcePort

`func (o *MicrosoftGraphNetworkConnection) HasNatSourcePort() bool`

HasNatSourcePort returns a boolean if a field has been set.

### SetNatSourcePortNil

`func (o *MicrosoftGraphNetworkConnection) SetNatSourcePortNil(b bool)`

 SetNatSourcePortNil sets the value for NatSourcePort to be an explicit nil

### UnsetNatSourcePort
`func (o *MicrosoftGraphNetworkConnection) UnsetNatSourcePort()`

UnsetNatSourcePort ensures that no value is present for NatSourcePort, not even an explicit nil
### GetProtocol

`func (o *MicrosoftGraphNetworkConnection) GetProtocol() AnyOfmicrosoftGraphSecurityNetworkProtocol`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *MicrosoftGraphNetworkConnection) GetProtocolOk() (*AnyOfmicrosoftGraphSecurityNetworkProtocol, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *MicrosoftGraphNetworkConnection) SetProtocol(v AnyOfmicrosoftGraphSecurityNetworkProtocol)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *MicrosoftGraphNetworkConnection) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### SetProtocolNil

`func (o *MicrosoftGraphNetworkConnection) SetProtocolNil(b bool)`

 SetProtocolNil sets the value for Protocol to be an explicit nil

### UnsetProtocol
`func (o *MicrosoftGraphNetworkConnection) UnsetProtocol()`

UnsetProtocol ensures that no value is present for Protocol, not even an explicit nil
### GetRiskScore

`func (o *MicrosoftGraphNetworkConnection) GetRiskScore() string`

GetRiskScore returns the RiskScore field if non-nil, zero value otherwise.

### GetRiskScoreOk

`func (o *MicrosoftGraphNetworkConnection) GetRiskScoreOk() (*string, bool)`

GetRiskScoreOk returns a tuple with the RiskScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskScore

`func (o *MicrosoftGraphNetworkConnection) SetRiskScore(v string)`

SetRiskScore sets RiskScore field to given value.

### HasRiskScore

`func (o *MicrosoftGraphNetworkConnection) HasRiskScore() bool`

HasRiskScore returns a boolean if a field has been set.

### SetRiskScoreNil

`func (o *MicrosoftGraphNetworkConnection) SetRiskScoreNil(b bool)`

 SetRiskScoreNil sets the value for RiskScore to be an explicit nil

### UnsetRiskScore
`func (o *MicrosoftGraphNetworkConnection) UnsetRiskScore()`

UnsetRiskScore ensures that no value is present for RiskScore, not even an explicit nil
### GetSourceAddress

`func (o *MicrosoftGraphNetworkConnection) GetSourceAddress() string`

GetSourceAddress returns the SourceAddress field if non-nil, zero value otherwise.

### GetSourceAddressOk

`func (o *MicrosoftGraphNetworkConnection) GetSourceAddressOk() (*string, bool)`

GetSourceAddressOk returns a tuple with the SourceAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAddress

`func (o *MicrosoftGraphNetworkConnection) SetSourceAddress(v string)`

SetSourceAddress sets SourceAddress field to given value.

### HasSourceAddress

`func (o *MicrosoftGraphNetworkConnection) HasSourceAddress() bool`

HasSourceAddress returns a boolean if a field has been set.

### SetSourceAddressNil

`func (o *MicrosoftGraphNetworkConnection) SetSourceAddressNil(b bool)`

 SetSourceAddressNil sets the value for SourceAddress to be an explicit nil

### UnsetSourceAddress
`func (o *MicrosoftGraphNetworkConnection) UnsetSourceAddress()`

UnsetSourceAddress ensures that no value is present for SourceAddress, not even an explicit nil
### GetSourceLocation

`func (o *MicrosoftGraphNetworkConnection) GetSourceLocation() string`

GetSourceLocation returns the SourceLocation field if non-nil, zero value otherwise.

### GetSourceLocationOk

`func (o *MicrosoftGraphNetworkConnection) GetSourceLocationOk() (*string, bool)`

GetSourceLocationOk returns a tuple with the SourceLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceLocation

`func (o *MicrosoftGraphNetworkConnection) SetSourceLocation(v string)`

SetSourceLocation sets SourceLocation field to given value.

### HasSourceLocation

`func (o *MicrosoftGraphNetworkConnection) HasSourceLocation() bool`

HasSourceLocation returns a boolean if a field has been set.

### SetSourceLocationNil

`func (o *MicrosoftGraphNetworkConnection) SetSourceLocationNil(b bool)`

 SetSourceLocationNil sets the value for SourceLocation to be an explicit nil

### UnsetSourceLocation
`func (o *MicrosoftGraphNetworkConnection) UnsetSourceLocation()`

UnsetSourceLocation ensures that no value is present for SourceLocation, not even an explicit nil
### GetSourcePort

`func (o *MicrosoftGraphNetworkConnection) GetSourcePort() string`

GetSourcePort returns the SourcePort field if non-nil, zero value otherwise.

### GetSourcePortOk

`func (o *MicrosoftGraphNetworkConnection) GetSourcePortOk() (*string, bool)`

GetSourcePortOk returns a tuple with the SourcePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourcePort

`func (o *MicrosoftGraphNetworkConnection) SetSourcePort(v string)`

SetSourcePort sets SourcePort field to given value.

### HasSourcePort

`func (o *MicrosoftGraphNetworkConnection) HasSourcePort() bool`

HasSourcePort returns a boolean if a field has been set.

### SetSourcePortNil

`func (o *MicrosoftGraphNetworkConnection) SetSourcePortNil(b bool)`

 SetSourcePortNil sets the value for SourcePort to be an explicit nil

### UnsetSourcePort
`func (o *MicrosoftGraphNetworkConnection) UnsetSourcePort()`

UnsetSourcePort ensures that no value is present for SourcePort, not even an explicit nil
### GetStatus

`func (o *MicrosoftGraphNetworkConnection) GetStatus() AnyOfmicrosoftGraphConnectionStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MicrosoftGraphNetworkConnection) GetStatusOk() (*AnyOfmicrosoftGraphConnectionStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MicrosoftGraphNetworkConnection) SetStatus(v AnyOfmicrosoftGraphConnectionStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MicrosoftGraphNetworkConnection) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *MicrosoftGraphNetworkConnection) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *MicrosoftGraphNetworkConnection) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetUrlParameters

`func (o *MicrosoftGraphNetworkConnection) GetUrlParameters() string`

GetUrlParameters returns the UrlParameters field if non-nil, zero value otherwise.

### GetUrlParametersOk

`func (o *MicrosoftGraphNetworkConnection) GetUrlParametersOk() (*string, bool)`

GetUrlParametersOk returns a tuple with the UrlParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlParameters

`func (o *MicrosoftGraphNetworkConnection) SetUrlParameters(v string)`

SetUrlParameters sets UrlParameters field to given value.

### HasUrlParameters

`func (o *MicrosoftGraphNetworkConnection) HasUrlParameters() bool`

HasUrlParameters returns a boolean if a field has been set.

### SetUrlParametersNil

`func (o *MicrosoftGraphNetworkConnection) SetUrlParametersNil(b bool)`

 SetUrlParametersNil sets the value for UrlParameters to be an explicit nil

### UnsetUrlParameters
`func (o *MicrosoftGraphNetworkConnection) UnsetUrlParameters()`

UnsetUrlParameters ensures that no value is present for UrlParameters, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


