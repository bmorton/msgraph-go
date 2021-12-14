# MicrosoftGraphProvisionedPlan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CapabilityStatus** | Pointer to **NullableString** | For example, &#39;Enabled&#39;. | [optional] 
**ProvisioningStatus** | Pointer to **NullableString** | For example, &#39;Success&#39;. | [optional] 
**Service** | Pointer to **NullableString** | The name of the service; for example, &#39;AccessControlS2S&#39; | [optional] 

## Methods

### NewMicrosoftGraphProvisionedPlan

`func NewMicrosoftGraphProvisionedPlan() *MicrosoftGraphProvisionedPlan`

NewMicrosoftGraphProvisionedPlan instantiates a new MicrosoftGraphProvisionedPlan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphProvisionedPlanWithDefaults

`func NewMicrosoftGraphProvisionedPlanWithDefaults() *MicrosoftGraphProvisionedPlan`

NewMicrosoftGraphProvisionedPlanWithDefaults instantiates a new MicrosoftGraphProvisionedPlan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCapabilityStatus

`func (o *MicrosoftGraphProvisionedPlan) GetCapabilityStatus() string`

GetCapabilityStatus returns the CapabilityStatus field if non-nil, zero value otherwise.

### GetCapabilityStatusOk

`func (o *MicrosoftGraphProvisionedPlan) GetCapabilityStatusOk() (*string, bool)`

GetCapabilityStatusOk returns a tuple with the CapabilityStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilityStatus

`func (o *MicrosoftGraphProvisionedPlan) SetCapabilityStatus(v string)`

SetCapabilityStatus sets CapabilityStatus field to given value.

### HasCapabilityStatus

`func (o *MicrosoftGraphProvisionedPlan) HasCapabilityStatus() bool`

HasCapabilityStatus returns a boolean if a field has been set.

### SetCapabilityStatusNil

`func (o *MicrosoftGraphProvisionedPlan) SetCapabilityStatusNil(b bool)`

 SetCapabilityStatusNil sets the value for CapabilityStatus to be an explicit nil

### UnsetCapabilityStatus
`func (o *MicrosoftGraphProvisionedPlan) UnsetCapabilityStatus()`

UnsetCapabilityStatus ensures that no value is present for CapabilityStatus, not even an explicit nil
### GetProvisioningStatus

`func (o *MicrosoftGraphProvisionedPlan) GetProvisioningStatus() string`

GetProvisioningStatus returns the ProvisioningStatus field if non-nil, zero value otherwise.

### GetProvisioningStatusOk

`func (o *MicrosoftGraphProvisionedPlan) GetProvisioningStatusOk() (*string, bool)`

GetProvisioningStatusOk returns a tuple with the ProvisioningStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningStatus

`func (o *MicrosoftGraphProvisionedPlan) SetProvisioningStatus(v string)`

SetProvisioningStatus sets ProvisioningStatus field to given value.

### HasProvisioningStatus

`func (o *MicrosoftGraphProvisionedPlan) HasProvisioningStatus() bool`

HasProvisioningStatus returns a boolean if a field has been set.

### SetProvisioningStatusNil

`func (o *MicrosoftGraphProvisionedPlan) SetProvisioningStatusNil(b bool)`

 SetProvisioningStatusNil sets the value for ProvisioningStatus to be an explicit nil

### UnsetProvisioningStatus
`func (o *MicrosoftGraphProvisionedPlan) UnsetProvisioningStatus()`

UnsetProvisioningStatus ensures that no value is present for ProvisioningStatus, not even an explicit nil
### GetService

`func (o *MicrosoftGraphProvisionedPlan) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *MicrosoftGraphProvisionedPlan) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *MicrosoftGraphProvisionedPlan) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *MicrosoftGraphProvisionedPlan) HasService() bool`

HasService returns a boolean if a field has been set.

### SetServiceNil

`func (o *MicrosoftGraphProvisionedPlan) SetServiceNil(b bool)`

 SetServiceNil sets the value for Service to be an explicit nil

### UnsetService
`func (o *MicrosoftGraphProvisionedPlan) UnsetService()`

UnsetService ensures that no value is present for Service, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


