# MicrosoftGraphHostSecurityState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fqdn** | Pointer to **NullableString** | Host FQDN (Fully Qualified Domain Name) (for example, machine.company.com). | [optional] 
**IsAzureAdJoined** | Pointer to **NullableBool** |  | [optional] 
**IsAzureAdRegistered** | Pointer to **NullableBool** |  | [optional] 
**IsHybridAzureDomainJoined** | Pointer to **NullableBool** | True if the host is domain joined to an on-premises Active Directory domain. | [optional] 
**NetBiosName** | Pointer to **NullableString** | The local host name, without the DNS domain name. | [optional] 
**Os** | Pointer to **NullableString** | Host Operating System. (For example, Windows10, MacOS, RHEL, etc.). | [optional] 
**PrivateIpAddress** | Pointer to **NullableString** | Private (not routable) IPv4 or IPv6 address (see RFC 1918) at the time of the alert. | [optional] 
**PublicIpAddress** | Pointer to **NullableString** | Publicly routable IPv4 or IPv6 address (see RFC 1918) at time of the alert. | [optional] 
**RiskScore** | Pointer to **NullableString** | Provider-generated/calculated risk score of the host.  Recommended value range of 0-1, which equates to a percentage. | [optional] 

## Methods

### NewMicrosoftGraphHostSecurityState

`func NewMicrosoftGraphHostSecurityState() *MicrosoftGraphHostSecurityState`

NewMicrosoftGraphHostSecurityState instantiates a new MicrosoftGraphHostSecurityState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphHostSecurityStateWithDefaults

`func NewMicrosoftGraphHostSecurityStateWithDefaults() *MicrosoftGraphHostSecurityState`

NewMicrosoftGraphHostSecurityStateWithDefaults instantiates a new MicrosoftGraphHostSecurityState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFqdn

`func (o *MicrosoftGraphHostSecurityState) GetFqdn() string`

GetFqdn returns the Fqdn field if non-nil, zero value otherwise.

### GetFqdnOk

`func (o *MicrosoftGraphHostSecurityState) GetFqdnOk() (*string, bool)`

GetFqdnOk returns a tuple with the Fqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqdn

`func (o *MicrosoftGraphHostSecurityState) SetFqdn(v string)`

SetFqdn sets Fqdn field to given value.

### HasFqdn

`func (o *MicrosoftGraphHostSecurityState) HasFqdn() bool`

HasFqdn returns a boolean if a field has been set.

### SetFqdnNil

`func (o *MicrosoftGraphHostSecurityState) SetFqdnNil(b bool)`

 SetFqdnNil sets the value for Fqdn to be an explicit nil

### UnsetFqdn
`func (o *MicrosoftGraphHostSecurityState) UnsetFqdn()`

UnsetFqdn ensures that no value is present for Fqdn, not even an explicit nil
### GetIsAzureAdJoined

`func (o *MicrosoftGraphHostSecurityState) GetIsAzureAdJoined() bool`

GetIsAzureAdJoined returns the IsAzureAdJoined field if non-nil, zero value otherwise.

### GetIsAzureAdJoinedOk

`func (o *MicrosoftGraphHostSecurityState) GetIsAzureAdJoinedOk() (*bool, bool)`

GetIsAzureAdJoinedOk returns a tuple with the IsAzureAdJoined field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAzureAdJoined

`func (o *MicrosoftGraphHostSecurityState) SetIsAzureAdJoined(v bool)`

SetIsAzureAdJoined sets IsAzureAdJoined field to given value.

### HasIsAzureAdJoined

`func (o *MicrosoftGraphHostSecurityState) HasIsAzureAdJoined() bool`

HasIsAzureAdJoined returns a boolean if a field has been set.

### SetIsAzureAdJoinedNil

`func (o *MicrosoftGraphHostSecurityState) SetIsAzureAdJoinedNil(b bool)`

 SetIsAzureAdJoinedNil sets the value for IsAzureAdJoined to be an explicit nil

### UnsetIsAzureAdJoined
`func (o *MicrosoftGraphHostSecurityState) UnsetIsAzureAdJoined()`

UnsetIsAzureAdJoined ensures that no value is present for IsAzureAdJoined, not even an explicit nil
### GetIsAzureAdRegistered

`func (o *MicrosoftGraphHostSecurityState) GetIsAzureAdRegistered() bool`

GetIsAzureAdRegistered returns the IsAzureAdRegistered field if non-nil, zero value otherwise.

### GetIsAzureAdRegisteredOk

`func (o *MicrosoftGraphHostSecurityState) GetIsAzureAdRegisteredOk() (*bool, bool)`

GetIsAzureAdRegisteredOk returns a tuple with the IsAzureAdRegistered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAzureAdRegistered

`func (o *MicrosoftGraphHostSecurityState) SetIsAzureAdRegistered(v bool)`

SetIsAzureAdRegistered sets IsAzureAdRegistered field to given value.

### HasIsAzureAdRegistered

`func (o *MicrosoftGraphHostSecurityState) HasIsAzureAdRegistered() bool`

HasIsAzureAdRegistered returns a boolean if a field has been set.

### SetIsAzureAdRegisteredNil

`func (o *MicrosoftGraphHostSecurityState) SetIsAzureAdRegisteredNil(b bool)`

 SetIsAzureAdRegisteredNil sets the value for IsAzureAdRegistered to be an explicit nil

### UnsetIsAzureAdRegistered
`func (o *MicrosoftGraphHostSecurityState) UnsetIsAzureAdRegistered()`

UnsetIsAzureAdRegistered ensures that no value is present for IsAzureAdRegistered, not even an explicit nil
### GetIsHybridAzureDomainJoined

`func (o *MicrosoftGraphHostSecurityState) GetIsHybridAzureDomainJoined() bool`

GetIsHybridAzureDomainJoined returns the IsHybridAzureDomainJoined field if non-nil, zero value otherwise.

### GetIsHybridAzureDomainJoinedOk

`func (o *MicrosoftGraphHostSecurityState) GetIsHybridAzureDomainJoinedOk() (*bool, bool)`

GetIsHybridAzureDomainJoinedOk returns a tuple with the IsHybridAzureDomainJoined field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHybridAzureDomainJoined

`func (o *MicrosoftGraphHostSecurityState) SetIsHybridAzureDomainJoined(v bool)`

SetIsHybridAzureDomainJoined sets IsHybridAzureDomainJoined field to given value.

### HasIsHybridAzureDomainJoined

`func (o *MicrosoftGraphHostSecurityState) HasIsHybridAzureDomainJoined() bool`

HasIsHybridAzureDomainJoined returns a boolean if a field has been set.

### SetIsHybridAzureDomainJoinedNil

`func (o *MicrosoftGraphHostSecurityState) SetIsHybridAzureDomainJoinedNil(b bool)`

 SetIsHybridAzureDomainJoinedNil sets the value for IsHybridAzureDomainJoined to be an explicit nil

### UnsetIsHybridAzureDomainJoined
`func (o *MicrosoftGraphHostSecurityState) UnsetIsHybridAzureDomainJoined()`

UnsetIsHybridAzureDomainJoined ensures that no value is present for IsHybridAzureDomainJoined, not even an explicit nil
### GetNetBiosName

`func (o *MicrosoftGraphHostSecurityState) GetNetBiosName() string`

GetNetBiosName returns the NetBiosName field if non-nil, zero value otherwise.

### GetNetBiosNameOk

`func (o *MicrosoftGraphHostSecurityState) GetNetBiosNameOk() (*string, bool)`

GetNetBiosNameOk returns a tuple with the NetBiosName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetBiosName

`func (o *MicrosoftGraphHostSecurityState) SetNetBiosName(v string)`

SetNetBiosName sets NetBiosName field to given value.

### HasNetBiosName

`func (o *MicrosoftGraphHostSecurityState) HasNetBiosName() bool`

HasNetBiosName returns a boolean if a field has been set.

### SetNetBiosNameNil

`func (o *MicrosoftGraphHostSecurityState) SetNetBiosNameNil(b bool)`

 SetNetBiosNameNil sets the value for NetBiosName to be an explicit nil

### UnsetNetBiosName
`func (o *MicrosoftGraphHostSecurityState) UnsetNetBiosName()`

UnsetNetBiosName ensures that no value is present for NetBiosName, not even an explicit nil
### GetOs

`func (o *MicrosoftGraphHostSecurityState) GetOs() string`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *MicrosoftGraphHostSecurityState) GetOsOk() (*string, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOs

`func (o *MicrosoftGraphHostSecurityState) SetOs(v string)`

SetOs sets Os field to given value.

### HasOs

`func (o *MicrosoftGraphHostSecurityState) HasOs() bool`

HasOs returns a boolean if a field has been set.

### SetOsNil

`func (o *MicrosoftGraphHostSecurityState) SetOsNil(b bool)`

 SetOsNil sets the value for Os to be an explicit nil

### UnsetOs
`func (o *MicrosoftGraphHostSecurityState) UnsetOs()`

UnsetOs ensures that no value is present for Os, not even an explicit nil
### GetPrivateIpAddress

`func (o *MicrosoftGraphHostSecurityState) GetPrivateIpAddress() string`

GetPrivateIpAddress returns the PrivateIpAddress field if non-nil, zero value otherwise.

### GetPrivateIpAddressOk

`func (o *MicrosoftGraphHostSecurityState) GetPrivateIpAddressOk() (*string, bool)`

GetPrivateIpAddressOk returns a tuple with the PrivateIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateIpAddress

`func (o *MicrosoftGraphHostSecurityState) SetPrivateIpAddress(v string)`

SetPrivateIpAddress sets PrivateIpAddress field to given value.

### HasPrivateIpAddress

`func (o *MicrosoftGraphHostSecurityState) HasPrivateIpAddress() bool`

HasPrivateIpAddress returns a boolean if a field has been set.

### SetPrivateIpAddressNil

`func (o *MicrosoftGraphHostSecurityState) SetPrivateIpAddressNil(b bool)`

 SetPrivateIpAddressNil sets the value for PrivateIpAddress to be an explicit nil

### UnsetPrivateIpAddress
`func (o *MicrosoftGraphHostSecurityState) UnsetPrivateIpAddress()`

UnsetPrivateIpAddress ensures that no value is present for PrivateIpAddress, not even an explicit nil
### GetPublicIpAddress

`func (o *MicrosoftGraphHostSecurityState) GetPublicIpAddress() string`

GetPublicIpAddress returns the PublicIpAddress field if non-nil, zero value otherwise.

### GetPublicIpAddressOk

`func (o *MicrosoftGraphHostSecurityState) GetPublicIpAddressOk() (*string, bool)`

GetPublicIpAddressOk returns a tuple with the PublicIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIpAddress

`func (o *MicrosoftGraphHostSecurityState) SetPublicIpAddress(v string)`

SetPublicIpAddress sets PublicIpAddress field to given value.

### HasPublicIpAddress

`func (o *MicrosoftGraphHostSecurityState) HasPublicIpAddress() bool`

HasPublicIpAddress returns a boolean if a field has been set.

### SetPublicIpAddressNil

`func (o *MicrosoftGraphHostSecurityState) SetPublicIpAddressNil(b bool)`

 SetPublicIpAddressNil sets the value for PublicIpAddress to be an explicit nil

### UnsetPublicIpAddress
`func (o *MicrosoftGraphHostSecurityState) UnsetPublicIpAddress()`

UnsetPublicIpAddress ensures that no value is present for PublicIpAddress, not even an explicit nil
### GetRiskScore

`func (o *MicrosoftGraphHostSecurityState) GetRiskScore() string`

GetRiskScore returns the RiskScore field if non-nil, zero value otherwise.

### GetRiskScoreOk

`func (o *MicrosoftGraphHostSecurityState) GetRiskScoreOk() (*string, bool)`

GetRiskScoreOk returns a tuple with the RiskScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskScore

`func (o *MicrosoftGraphHostSecurityState) SetRiskScore(v string)`

SetRiskScore sets RiskScore field to given value.

### HasRiskScore

`func (o *MicrosoftGraphHostSecurityState) HasRiskScore() bool`

HasRiskScore returns a boolean if a field has been set.

### SetRiskScoreNil

`func (o *MicrosoftGraphHostSecurityState) SetRiskScoreNil(b bool)`

 SetRiskScoreNil sets the value for RiskScore to be an explicit nil

### UnsetRiskScore
`func (o *MicrosoftGraphHostSecurityState) UnsetRiskScore()`

UnsetRiskScore ensures that no value is present for RiskScore, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


