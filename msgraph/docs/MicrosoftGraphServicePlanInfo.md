# MicrosoftGraphServicePlanInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppliesTo** | Pointer to **NullableString** | The object the service plan can be assigned to. Possible values:&#39;User&#39; - service plan can be assigned to individual users.&#39;Company&#39; - service plan can be assigned to the entire tenant. | [optional] 
**ProvisioningStatus** | Pointer to **NullableString** | The provisioning status of the service plan. Possible values:&#39;Success&#39; - Service is fully provisioned.&#39;Disabled&#39; - Service has been disabled.&#39;PendingInput&#39; - Service is not yet provisioned; awaiting service confirmation.&#39;PendingActivation&#39; - Service is provisioned but requires explicit activation by administrator (for example, Intune_O365 service plan)&#39;PendingProvisioning&#39; - Microsoft has added a new service to the product SKU and it has not been activated in the tenant, yet. | [optional] 
**ServicePlanId** | Pointer to **NullableString** | The unique identifier of the service plan. | [optional] 
**ServicePlanName** | Pointer to **NullableString** | The name of the service plan. | [optional] 

## Methods

### NewMicrosoftGraphServicePlanInfo

`func NewMicrosoftGraphServicePlanInfo() *MicrosoftGraphServicePlanInfo`

NewMicrosoftGraphServicePlanInfo instantiates a new MicrosoftGraphServicePlanInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphServicePlanInfoWithDefaults

`func NewMicrosoftGraphServicePlanInfoWithDefaults() *MicrosoftGraphServicePlanInfo`

NewMicrosoftGraphServicePlanInfoWithDefaults instantiates a new MicrosoftGraphServicePlanInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppliesTo

`func (o *MicrosoftGraphServicePlanInfo) GetAppliesTo() string`

GetAppliesTo returns the AppliesTo field if non-nil, zero value otherwise.

### GetAppliesToOk

`func (o *MicrosoftGraphServicePlanInfo) GetAppliesToOk() (*string, bool)`

GetAppliesToOk returns a tuple with the AppliesTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliesTo

`func (o *MicrosoftGraphServicePlanInfo) SetAppliesTo(v string)`

SetAppliesTo sets AppliesTo field to given value.

### HasAppliesTo

`func (o *MicrosoftGraphServicePlanInfo) HasAppliesTo() bool`

HasAppliesTo returns a boolean if a field has been set.

### SetAppliesToNil

`func (o *MicrosoftGraphServicePlanInfo) SetAppliesToNil(b bool)`

 SetAppliesToNil sets the value for AppliesTo to be an explicit nil

### UnsetAppliesTo
`func (o *MicrosoftGraphServicePlanInfo) UnsetAppliesTo()`

UnsetAppliesTo ensures that no value is present for AppliesTo, not even an explicit nil
### GetProvisioningStatus

`func (o *MicrosoftGraphServicePlanInfo) GetProvisioningStatus() string`

GetProvisioningStatus returns the ProvisioningStatus field if non-nil, zero value otherwise.

### GetProvisioningStatusOk

`func (o *MicrosoftGraphServicePlanInfo) GetProvisioningStatusOk() (*string, bool)`

GetProvisioningStatusOk returns a tuple with the ProvisioningStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningStatus

`func (o *MicrosoftGraphServicePlanInfo) SetProvisioningStatus(v string)`

SetProvisioningStatus sets ProvisioningStatus field to given value.

### HasProvisioningStatus

`func (o *MicrosoftGraphServicePlanInfo) HasProvisioningStatus() bool`

HasProvisioningStatus returns a boolean if a field has been set.

### SetProvisioningStatusNil

`func (o *MicrosoftGraphServicePlanInfo) SetProvisioningStatusNil(b bool)`

 SetProvisioningStatusNil sets the value for ProvisioningStatus to be an explicit nil

### UnsetProvisioningStatus
`func (o *MicrosoftGraphServicePlanInfo) UnsetProvisioningStatus()`

UnsetProvisioningStatus ensures that no value is present for ProvisioningStatus, not even an explicit nil
### GetServicePlanId

`func (o *MicrosoftGraphServicePlanInfo) GetServicePlanId() string`

GetServicePlanId returns the ServicePlanId field if non-nil, zero value otherwise.

### GetServicePlanIdOk

`func (o *MicrosoftGraphServicePlanInfo) GetServicePlanIdOk() (*string, bool)`

GetServicePlanIdOk returns a tuple with the ServicePlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePlanId

`func (o *MicrosoftGraphServicePlanInfo) SetServicePlanId(v string)`

SetServicePlanId sets ServicePlanId field to given value.

### HasServicePlanId

`func (o *MicrosoftGraphServicePlanInfo) HasServicePlanId() bool`

HasServicePlanId returns a boolean if a field has been set.

### SetServicePlanIdNil

`func (o *MicrosoftGraphServicePlanInfo) SetServicePlanIdNil(b bool)`

 SetServicePlanIdNil sets the value for ServicePlanId to be an explicit nil

### UnsetServicePlanId
`func (o *MicrosoftGraphServicePlanInfo) UnsetServicePlanId()`

UnsetServicePlanId ensures that no value is present for ServicePlanId, not even an explicit nil
### GetServicePlanName

`func (o *MicrosoftGraphServicePlanInfo) GetServicePlanName() string`

GetServicePlanName returns the ServicePlanName field if non-nil, zero value otherwise.

### GetServicePlanNameOk

`func (o *MicrosoftGraphServicePlanInfo) GetServicePlanNameOk() (*string, bool)`

GetServicePlanNameOk returns a tuple with the ServicePlanName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePlanName

`func (o *MicrosoftGraphServicePlanInfo) SetServicePlanName(v string)`

SetServicePlanName sets ServicePlanName field to given value.

### HasServicePlanName

`func (o *MicrosoftGraphServicePlanInfo) HasServicePlanName() bool`

HasServicePlanName returns a boolean if a field has been set.

### SetServicePlanNameNil

`func (o *MicrosoftGraphServicePlanInfo) SetServicePlanNameNil(b bool)`

 SetServicePlanNameNil sets the value for ServicePlanName to be an explicit nil

### UnsetServicePlanName
`func (o *MicrosoftGraphServicePlanInfo) UnsetServicePlanName()`

UnsetServicePlanName ensures that no value is present for ServicePlanName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


