# LicenseDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServicePlans** | Pointer to [**[]MicrosoftGraphServicePlanInfo**](MicrosoftGraphServicePlanInfo.md) | Information about the service plans assigned with the license. Read-only, Not nullable | [optional] 
**SkuId** | Pointer to **NullableString** | Unique identifier (GUID) for the service SKU. Equal to the skuId property on the related SubscribedSku object. Read-only | [optional] 
**SkuPartNumber** | Pointer to **NullableString** | Unique SKU display name. Equal to the skuPartNumber on the related SubscribedSku object; for example: &#39;AAD_Premium&#39;. Read-only | [optional] 

## Methods

### NewLicenseDetails

`func NewLicenseDetails() *LicenseDetails`

NewLicenseDetails instantiates a new LicenseDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLicenseDetailsWithDefaults

`func NewLicenseDetailsWithDefaults() *LicenseDetails`

NewLicenseDetailsWithDefaults instantiates a new LicenseDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServicePlans

`func (o *LicenseDetails) GetServicePlans() []MicrosoftGraphServicePlanInfo`

GetServicePlans returns the ServicePlans field if non-nil, zero value otherwise.

### GetServicePlansOk

`func (o *LicenseDetails) GetServicePlansOk() (*[]MicrosoftGraphServicePlanInfo, bool)`

GetServicePlansOk returns a tuple with the ServicePlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePlans

`func (o *LicenseDetails) SetServicePlans(v []MicrosoftGraphServicePlanInfo)`

SetServicePlans sets ServicePlans field to given value.

### HasServicePlans

`func (o *LicenseDetails) HasServicePlans() bool`

HasServicePlans returns a boolean if a field has been set.

### GetSkuId

`func (o *LicenseDetails) GetSkuId() string`

GetSkuId returns the SkuId field if non-nil, zero value otherwise.

### GetSkuIdOk

`func (o *LicenseDetails) GetSkuIdOk() (*string, bool)`

GetSkuIdOk returns a tuple with the SkuId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkuId

`func (o *LicenseDetails) SetSkuId(v string)`

SetSkuId sets SkuId field to given value.

### HasSkuId

`func (o *LicenseDetails) HasSkuId() bool`

HasSkuId returns a boolean if a field has been set.

### SetSkuIdNil

`func (o *LicenseDetails) SetSkuIdNil(b bool)`

 SetSkuIdNil sets the value for SkuId to be an explicit nil

### UnsetSkuId
`func (o *LicenseDetails) UnsetSkuId()`

UnsetSkuId ensures that no value is present for SkuId, not even an explicit nil
### GetSkuPartNumber

`func (o *LicenseDetails) GetSkuPartNumber() string`

GetSkuPartNumber returns the SkuPartNumber field if non-nil, zero value otherwise.

### GetSkuPartNumberOk

`func (o *LicenseDetails) GetSkuPartNumberOk() (*string, bool)`

GetSkuPartNumberOk returns a tuple with the SkuPartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkuPartNumber

`func (o *LicenseDetails) SetSkuPartNumber(v string)`

SetSkuPartNumber sets SkuPartNumber field to given value.

### HasSkuPartNumber

`func (o *LicenseDetails) HasSkuPartNumber() bool`

HasSkuPartNumber returns a boolean if a field has been set.

### SetSkuPartNumberNil

`func (o *LicenseDetails) SetSkuPartNumberNil(b bool)`

 SetSkuPartNumberNil sets the value for SkuPartNumber to be an explicit nil

### UnsetSkuPartNumber
`func (o *LicenseDetails) UnsetSkuPartNumber()`

UnsetSkuPartNumber ensures that no value is present for SkuPartNumber, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


