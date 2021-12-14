# MicrosoftGraphLicenseDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**ServicePlans** | Pointer to [**[]MicrosoftGraphServicePlanInfo**](MicrosoftGraphServicePlanInfo.md) | Information about the service plans assigned with the license. Read-only, Not nullable | [optional] 
**SkuId** | Pointer to **NullableString** | Unique identifier (GUID) for the service SKU. Equal to the skuId property on the related SubscribedSku object. Read-only | [optional] 
**SkuPartNumber** | Pointer to **NullableString** | Unique SKU display name. Equal to the skuPartNumber on the related SubscribedSku object; for example: &#39;AAD_Premium&#39;. Read-only | [optional] 

## Methods

### NewMicrosoftGraphLicenseDetails

`func NewMicrosoftGraphLicenseDetails() *MicrosoftGraphLicenseDetails`

NewMicrosoftGraphLicenseDetails instantiates a new MicrosoftGraphLicenseDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphLicenseDetailsWithDefaults

`func NewMicrosoftGraphLicenseDetailsWithDefaults() *MicrosoftGraphLicenseDetails`

NewMicrosoftGraphLicenseDetailsWithDefaults instantiates a new MicrosoftGraphLicenseDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphLicenseDetails) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphLicenseDetails) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphLicenseDetails) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphLicenseDetails) HasId() bool`

HasId returns a boolean if a field has been set.

### GetServicePlans

`func (o *MicrosoftGraphLicenseDetails) GetServicePlans() []MicrosoftGraphServicePlanInfo`

GetServicePlans returns the ServicePlans field if non-nil, zero value otherwise.

### GetServicePlansOk

`func (o *MicrosoftGraphLicenseDetails) GetServicePlansOk() (*[]MicrosoftGraphServicePlanInfo, bool)`

GetServicePlansOk returns a tuple with the ServicePlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePlans

`func (o *MicrosoftGraphLicenseDetails) SetServicePlans(v []MicrosoftGraphServicePlanInfo)`

SetServicePlans sets ServicePlans field to given value.

### HasServicePlans

`func (o *MicrosoftGraphLicenseDetails) HasServicePlans() bool`

HasServicePlans returns a boolean if a field has been set.

### GetSkuId

`func (o *MicrosoftGraphLicenseDetails) GetSkuId() string`

GetSkuId returns the SkuId field if non-nil, zero value otherwise.

### GetSkuIdOk

`func (o *MicrosoftGraphLicenseDetails) GetSkuIdOk() (*string, bool)`

GetSkuIdOk returns a tuple with the SkuId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkuId

`func (o *MicrosoftGraphLicenseDetails) SetSkuId(v string)`

SetSkuId sets SkuId field to given value.

### HasSkuId

`func (o *MicrosoftGraphLicenseDetails) HasSkuId() bool`

HasSkuId returns a boolean if a field has been set.

### SetSkuIdNil

`func (o *MicrosoftGraphLicenseDetails) SetSkuIdNil(b bool)`

 SetSkuIdNil sets the value for SkuId to be an explicit nil

### UnsetSkuId
`func (o *MicrosoftGraphLicenseDetails) UnsetSkuId()`

UnsetSkuId ensures that no value is present for SkuId, not even an explicit nil
### GetSkuPartNumber

`func (o *MicrosoftGraphLicenseDetails) GetSkuPartNumber() string`

GetSkuPartNumber returns the SkuPartNumber field if non-nil, zero value otherwise.

### GetSkuPartNumberOk

`func (o *MicrosoftGraphLicenseDetails) GetSkuPartNumberOk() (*string, bool)`

GetSkuPartNumberOk returns a tuple with the SkuPartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkuPartNumber

`func (o *MicrosoftGraphLicenseDetails) SetSkuPartNumber(v string)`

SetSkuPartNumber sets SkuPartNumber field to given value.

### HasSkuPartNumber

`func (o *MicrosoftGraphLicenseDetails) HasSkuPartNumber() bool`

HasSkuPartNumber returns a boolean if a field has been set.

### SetSkuPartNumberNil

`func (o *MicrosoftGraphLicenseDetails) SetSkuPartNumberNil(b bool)`

 SetSkuPartNumberNil sets the value for SkuPartNumber to be an explicit nil

### UnsetSkuPartNumber
`func (o *MicrosoftGraphLicenseDetails) UnsetSkuPartNumber()`

UnsetSkuPartNumber ensures that no value is present for SkuPartNumber, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


