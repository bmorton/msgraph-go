# MicrosoftGraphCloudAppSecurityState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestinationServiceIp** | Pointer to **NullableString** | Destination IP Address of the connection to the cloud application/service. | [optional] 
**DestinationServiceName** | Pointer to **NullableString** | Cloud application/service name (for example &#39;Salesforce&#39;, &#39;DropBox&#39;, etc.). | [optional] 
**RiskScore** | Pointer to **NullableString** | Provider-generated/calculated risk score of the Cloud Application/Service. Recommended value range of 0-1, which equates to a percentage. | [optional] 

## Methods

### NewMicrosoftGraphCloudAppSecurityState

`func NewMicrosoftGraphCloudAppSecurityState() *MicrosoftGraphCloudAppSecurityState`

NewMicrosoftGraphCloudAppSecurityState instantiates a new MicrosoftGraphCloudAppSecurityState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphCloudAppSecurityStateWithDefaults

`func NewMicrosoftGraphCloudAppSecurityStateWithDefaults() *MicrosoftGraphCloudAppSecurityState`

NewMicrosoftGraphCloudAppSecurityStateWithDefaults instantiates a new MicrosoftGraphCloudAppSecurityState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestinationServiceIp

`func (o *MicrosoftGraphCloudAppSecurityState) GetDestinationServiceIp() string`

GetDestinationServiceIp returns the DestinationServiceIp field if non-nil, zero value otherwise.

### GetDestinationServiceIpOk

`func (o *MicrosoftGraphCloudAppSecurityState) GetDestinationServiceIpOk() (*string, bool)`

GetDestinationServiceIpOk returns a tuple with the DestinationServiceIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationServiceIp

`func (o *MicrosoftGraphCloudAppSecurityState) SetDestinationServiceIp(v string)`

SetDestinationServiceIp sets DestinationServiceIp field to given value.

### HasDestinationServiceIp

`func (o *MicrosoftGraphCloudAppSecurityState) HasDestinationServiceIp() bool`

HasDestinationServiceIp returns a boolean if a field has been set.

### SetDestinationServiceIpNil

`func (o *MicrosoftGraphCloudAppSecurityState) SetDestinationServiceIpNil(b bool)`

 SetDestinationServiceIpNil sets the value for DestinationServiceIp to be an explicit nil

### UnsetDestinationServiceIp
`func (o *MicrosoftGraphCloudAppSecurityState) UnsetDestinationServiceIp()`

UnsetDestinationServiceIp ensures that no value is present for DestinationServiceIp, not even an explicit nil
### GetDestinationServiceName

`func (o *MicrosoftGraphCloudAppSecurityState) GetDestinationServiceName() string`

GetDestinationServiceName returns the DestinationServiceName field if non-nil, zero value otherwise.

### GetDestinationServiceNameOk

`func (o *MicrosoftGraphCloudAppSecurityState) GetDestinationServiceNameOk() (*string, bool)`

GetDestinationServiceNameOk returns a tuple with the DestinationServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationServiceName

`func (o *MicrosoftGraphCloudAppSecurityState) SetDestinationServiceName(v string)`

SetDestinationServiceName sets DestinationServiceName field to given value.

### HasDestinationServiceName

`func (o *MicrosoftGraphCloudAppSecurityState) HasDestinationServiceName() bool`

HasDestinationServiceName returns a boolean if a field has been set.

### SetDestinationServiceNameNil

`func (o *MicrosoftGraphCloudAppSecurityState) SetDestinationServiceNameNil(b bool)`

 SetDestinationServiceNameNil sets the value for DestinationServiceName to be an explicit nil

### UnsetDestinationServiceName
`func (o *MicrosoftGraphCloudAppSecurityState) UnsetDestinationServiceName()`

UnsetDestinationServiceName ensures that no value is present for DestinationServiceName, not even an explicit nil
### GetRiskScore

`func (o *MicrosoftGraphCloudAppSecurityState) GetRiskScore() string`

GetRiskScore returns the RiskScore field if non-nil, zero value otherwise.

### GetRiskScoreOk

`func (o *MicrosoftGraphCloudAppSecurityState) GetRiskScoreOk() (*string, bool)`

GetRiskScoreOk returns a tuple with the RiskScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskScore

`func (o *MicrosoftGraphCloudAppSecurityState) SetRiskScore(v string)`

SetRiskScore sets RiskScore field to given value.

### HasRiskScore

`func (o *MicrosoftGraphCloudAppSecurityState) HasRiskScore() bool`

HasRiskScore returns a boolean if a field has been set.

### SetRiskScoreNil

`func (o *MicrosoftGraphCloudAppSecurityState) SetRiskScoreNil(b bool)`

 SetRiskScoreNil sets the value for RiskScore to be an explicit nil

### UnsetRiskScore
`func (o *MicrosoftGraphCloudAppSecurityState) UnsetRiskScore()`

UnsetRiskScore ensures that no value is present for RiskScore, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


