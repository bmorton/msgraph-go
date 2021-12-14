# SubscribedSku

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppliesTo** | Pointer to **NullableString** | For example, &#39;User&#39; or &#39;Company&#39;. | [optional] 
**CapabilityStatus** | Pointer to **NullableString** | Possible values are: Enabled, Warning, Suspended, Deleted, LockedOut. The capabilityStatus is Enabled if the prepaidUnits property has at least 1 unit that is enabled, and LockedOut if the customer cancelled their subscription. | [optional] 
**ConsumedUnits** | Pointer to **NullableInt32** | The number of licenses that have been assigned. | [optional] 
**PrepaidUnits** | Pointer to [**AnyOfmicrosoftGraphLicenseUnitsDetail**](anyOf&lt;microsoft.graph.licenseUnitsDetail&gt;.md) | Information about the number and status of prepaid licenses. | [optional] 
**ServicePlans** | Pointer to [**[]MicrosoftGraphServicePlanInfo**](MicrosoftGraphServicePlanInfo.md) | Information about the service plans that are available with the SKU. Not nullable | [optional] 
**SkuId** | Pointer to **NullableString** | The unique identifier (GUID) for the service SKU. | [optional] 
**SkuPartNumber** | Pointer to **NullableString** | The SKU part number; for example: &#39;AAD_PREMIUM&#39; or &#39;RMSBASIC&#39;. To get a list of commercial subscriptions that an organization has acquired, see List subscribedSkus. | [optional] 

## Methods

### NewSubscribedSku

`func NewSubscribedSku() *SubscribedSku`

NewSubscribedSku instantiates a new SubscribedSku object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscribedSkuWithDefaults

`func NewSubscribedSkuWithDefaults() *SubscribedSku`

NewSubscribedSkuWithDefaults instantiates a new SubscribedSku object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppliesTo

`func (o *SubscribedSku) GetAppliesTo() string`

GetAppliesTo returns the AppliesTo field if non-nil, zero value otherwise.

### GetAppliesToOk

`func (o *SubscribedSku) GetAppliesToOk() (*string, bool)`

GetAppliesToOk returns a tuple with the AppliesTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliesTo

`func (o *SubscribedSku) SetAppliesTo(v string)`

SetAppliesTo sets AppliesTo field to given value.

### HasAppliesTo

`func (o *SubscribedSku) HasAppliesTo() bool`

HasAppliesTo returns a boolean if a field has been set.

### SetAppliesToNil

`func (o *SubscribedSku) SetAppliesToNil(b bool)`

 SetAppliesToNil sets the value for AppliesTo to be an explicit nil

### UnsetAppliesTo
`func (o *SubscribedSku) UnsetAppliesTo()`

UnsetAppliesTo ensures that no value is present for AppliesTo, not even an explicit nil
### GetCapabilityStatus

`func (o *SubscribedSku) GetCapabilityStatus() string`

GetCapabilityStatus returns the CapabilityStatus field if non-nil, zero value otherwise.

### GetCapabilityStatusOk

`func (o *SubscribedSku) GetCapabilityStatusOk() (*string, bool)`

GetCapabilityStatusOk returns a tuple with the CapabilityStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilityStatus

`func (o *SubscribedSku) SetCapabilityStatus(v string)`

SetCapabilityStatus sets CapabilityStatus field to given value.

### HasCapabilityStatus

`func (o *SubscribedSku) HasCapabilityStatus() bool`

HasCapabilityStatus returns a boolean if a field has been set.

### SetCapabilityStatusNil

`func (o *SubscribedSku) SetCapabilityStatusNil(b bool)`

 SetCapabilityStatusNil sets the value for CapabilityStatus to be an explicit nil

### UnsetCapabilityStatus
`func (o *SubscribedSku) UnsetCapabilityStatus()`

UnsetCapabilityStatus ensures that no value is present for CapabilityStatus, not even an explicit nil
### GetConsumedUnits

`func (o *SubscribedSku) GetConsumedUnits() int32`

GetConsumedUnits returns the ConsumedUnits field if non-nil, zero value otherwise.

### GetConsumedUnitsOk

`func (o *SubscribedSku) GetConsumedUnitsOk() (*int32, bool)`

GetConsumedUnitsOk returns a tuple with the ConsumedUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumedUnits

`func (o *SubscribedSku) SetConsumedUnits(v int32)`

SetConsumedUnits sets ConsumedUnits field to given value.

### HasConsumedUnits

`func (o *SubscribedSku) HasConsumedUnits() bool`

HasConsumedUnits returns a boolean if a field has been set.

### SetConsumedUnitsNil

`func (o *SubscribedSku) SetConsumedUnitsNil(b bool)`

 SetConsumedUnitsNil sets the value for ConsumedUnits to be an explicit nil

### UnsetConsumedUnits
`func (o *SubscribedSku) UnsetConsumedUnits()`

UnsetConsumedUnits ensures that no value is present for ConsumedUnits, not even an explicit nil
### GetPrepaidUnits

`func (o *SubscribedSku) GetPrepaidUnits() AnyOfmicrosoftGraphLicenseUnitsDetail`

GetPrepaidUnits returns the PrepaidUnits field if non-nil, zero value otherwise.

### GetPrepaidUnitsOk

`func (o *SubscribedSku) GetPrepaidUnitsOk() (*AnyOfmicrosoftGraphLicenseUnitsDetail, bool)`

GetPrepaidUnitsOk returns a tuple with the PrepaidUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrepaidUnits

`func (o *SubscribedSku) SetPrepaidUnits(v AnyOfmicrosoftGraphLicenseUnitsDetail)`

SetPrepaidUnits sets PrepaidUnits field to given value.

### HasPrepaidUnits

`func (o *SubscribedSku) HasPrepaidUnits() bool`

HasPrepaidUnits returns a boolean if a field has been set.

### SetPrepaidUnitsNil

`func (o *SubscribedSku) SetPrepaidUnitsNil(b bool)`

 SetPrepaidUnitsNil sets the value for PrepaidUnits to be an explicit nil

### UnsetPrepaidUnits
`func (o *SubscribedSku) UnsetPrepaidUnits()`

UnsetPrepaidUnits ensures that no value is present for PrepaidUnits, not even an explicit nil
### GetServicePlans

`func (o *SubscribedSku) GetServicePlans() []MicrosoftGraphServicePlanInfo`

GetServicePlans returns the ServicePlans field if non-nil, zero value otherwise.

### GetServicePlansOk

`func (o *SubscribedSku) GetServicePlansOk() (*[]MicrosoftGraphServicePlanInfo, bool)`

GetServicePlansOk returns a tuple with the ServicePlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePlans

`func (o *SubscribedSku) SetServicePlans(v []MicrosoftGraphServicePlanInfo)`

SetServicePlans sets ServicePlans field to given value.

### HasServicePlans

`func (o *SubscribedSku) HasServicePlans() bool`

HasServicePlans returns a boolean if a field has been set.

### GetSkuId

`func (o *SubscribedSku) GetSkuId() string`

GetSkuId returns the SkuId field if non-nil, zero value otherwise.

### GetSkuIdOk

`func (o *SubscribedSku) GetSkuIdOk() (*string, bool)`

GetSkuIdOk returns a tuple with the SkuId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkuId

`func (o *SubscribedSku) SetSkuId(v string)`

SetSkuId sets SkuId field to given value.

### HasSkuId

`func (o *SubscribedSku) HasSkuId() bool`

HasSkuId returns a boolean if a field has been set.

### SetSkuIdNil

`func (o *SubscribedSku) SetSkuIdNil(b bool)`

 SetSkuIdNil sets the value for SkuId to be an explicit nil

### UnsetSkuId
`func (o *SubscribedSku) UnsetSkuId()`

UnsetSkuId ensures that no value is present for SkuId, not even an explicit nil
### GetSkuPartNumber

`func (o *SubscribedSku) GetSkuPartNumber() string`

GetSkuPartNumber returns the SkuPartNumber field if non-nil, zero value otherwise.

### GetSkuPartNumberOk

`func (o *SubscribedSku) GetSkuPartNumberOk() (*string, bool)`

GetSkuPartNumberOk returns a tuple with the SkuPartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkuPartNumber

`func (o *SubscribedSku) SetSkuPartNumber(v string)`

SetSkuPartNumber sets SkuPartNumber field to given value.

### HasSkuPartNumber

`func (o *SubscribedSku) HasSkuPartNumber() bool`

HasSkuPartNumber returns a boolean if a field has been set.

### SetSkuPartNumberNil

`func (o *SubscribedSku) SetSkuPartNumberNil(b bool)`

 SetSkuPartNumberNil sets the value for SkuPartNumber to be an explicit nil

### UnsetSkuPartNumber
`func (o *SubscribedSku) UnsetSkuPartNumber()`

UnsetSkuPartNumber ensures that no value is present for SkuPartNumber, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


