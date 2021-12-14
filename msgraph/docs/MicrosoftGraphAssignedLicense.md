# MicrosoftGraphAssignedLicense

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisabledPlans** | Pointer to **[]string** | A collection of the unique identifiers for plans that have been disabled. | [optional] 
**SkuId** | Pointer to **NullableString** | The unique identifier for the SKU. | [optional] 

## Methods

### NewMicrosoftGraphAssignedLicense

`func NewMicrosoftGraphAssignedLicense() *MicrosoftGraphAssignedLicense`

NewMicrosoftGraphAssignedLicense instantiates a new MicrosoftGraphAssignedLicense object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphAssignedLicenseWithDefaults

`func NewMicrosoftGraphAssignedLicenseWithDefaults() *MicrosoftGraphAssignedLicense`

NewMicrosoftGraphAssignedLicenseWithDefaults instantiates a new MicrosoftGraphAssignedLicense object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisabledPlans

`func (o *MicrosoftGraphAssignedLicense) GetDisabledPlans() []string`

GetDisabledPlans returns the DisabledPlans field if non-nil, zero value otherwise.

### GetDisabledPlansOk

`func (o *MicrosoftGraphAssignedLicense) GetDisabledPlansOk() (*[]string, bool)`

GetDisabledPlansOk returns a tuple with the DisabledPlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledPlans

`func (o *MicrosoftGraphAssignedLicense) SetDisabledPlans(v []string)`

SetDisabledPlans sets DisabledPlans field to given value.

### HasDisabledPlans

`func (o *MicrosoftGraphAssignedLicense) HasDisabledPlans() bool`

HasDisabledPlans returns a boolean if a field has been set.

### GetSkuId

`func (o *MicrosoftGraphAssignedLicense) GetSkuId() string`

GetSkuId returns the SkuId field if non-nil, zero value otherwise.

### GetSkuIdOk

`func (o *MicrosoftGraphAssignedLicense) GetSkuIdOk() (*string, bool)`

GetSkuIdOk returns a tuple with the SkuId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkuId

`func (o *MicrosoftGraphAssignedLicense) SetSkuId(v string)`

SetSkuId sets SkuId field to given value.

### HasSkuId

`func (o *MicrosoftGraphAssignedLicense) HasSkuId() bool`

HasSkuId returns a boolean if a field has been set.

### SetSkuIdNil

`func (o *MicrosoftGraphAssignedLicense) SetSkuIdNil(b bool)`

 SetSkuIdNil sets the value for SkuId to be an explicit nil

### UnsetSkuId
`func (o *MicrosoftGraphAssignedLicense) UnsetSkuId()`

UnsetSkuId ensures that no value is present for SkuId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


