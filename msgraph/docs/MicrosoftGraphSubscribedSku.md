# MicrosoftGraphSubscribedSku

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**AppliesTo** | Pointer to **NullableString** | For example, &#39;User&#39; or &#39;Company&#39;. | [optional] 
**CapabilityStatus** | Pointer to **NullableString** | Possible values are: Enabled, Warning, Suspended, Deleted, LockedOut. The capabilityStatus is Enabled if the prepaidUnits property has at least 1 unit that is enabled, and LockedOut if the customer cancelled their subscription. | [optional] 
**ConsumedUnits** | Pointer to **NullableInt32** | The number of licenses that have been assigned. | [optional] 
**PrepaidUnits** | Pointer to [**AnyOfmicrosoftGraphLicenseUnitsDetail**](anyOf&lt;microsoft.graph.licenseUnitsDetail&gt;.md) | Information about the number and status of prepaid licenses. | [optional] 
**ServicePlans** | Pointer to [**[]MicrosoftGraphServicePlanInfo**](MicrosoftGraphServicePlanInfo.md) | Information about the service plans that are available with the SKU. Not nullable | [optional] 
**SkuId** | Pointer to **NullableString** | The unique identifier (GUID) for the service SKU. | [optional] 
**SkuPartNumber** | Pointer to **NullableString** | The SKU part number; for example: &#39;AAD_PREMIUM&#39; or &#39;RMSBASIC&#39;. To get a list of commercial subscriptions that an organization has acquired, see List subscribedSkus. | [optional] 

## Methods

### NewMicrosoftGraphSubscribedSku

`func NewMicrosoftGraphSubscribedSku() *MicrosoftGraphSubscribedSku`

NewMicrosoftGraphSubscribedSku instantiates a new MicrosoftGraphSubscribedSku object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphSubscribedSkuWithDefaults

`func NewMicrosoftGraphSubscribedSkuWithDefaults() *MicrosoftGraphSubscribedSku`

NewMicrosoftGraphSubscribedSkuWithDefaults instantiates a new MicrosoftGraphSubscribedSku object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphSubscribedSku) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphSubscribedSku) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphSubscribedSku) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphSubscribedSku) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAppliesTo

`func (o *MicrosoftGraphSubscribedSku) GetAppliesTo() string`

GetAppliesTo returns the AppliesTo field if non-nil, zero value otherwise.

### GetAppliesToOk

`func (o *MicrosoftGraphSubscribedSku) GetAppliesToOk() (*string, bool)`

GetAppliesToOk returns a tuple with the AppliesTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliesTo

`func (o *MicrosoftGraphSubscribedSku) SetAppliesTo(v string)`

SetAppliesTo sets AppliesTo field to given value.

### HasAppliesTo

`func (o *MicrosoftGraphSubscribedSku) HasAppliesTo() bool`

HasAppliesTo returns a boolean if a field has been set.

### SetAppliesToNil

`func (o *MicrosoftGraphSubscribedSku) SetAppliesToNil(b bool)`

 SetAppliesToNil sets the value for AppliesTo to be an explicit nil

### UnsetAppliesTo
`func (o *MicrosoftGraphSubscribedSku) UnsetAppliesTo()`

UnsetAppliesTo ensures that no value is present for AppliesTo, not even an explicit nil
### GetCapabilityStatus

`func (o *MicrosoftGraphSubscribedSku) GetCapabilityStatus() string`

GetCapabilityStatus returns the CapabilityStatus field if non-nil, zero value otherwise.

### GetCapabilityStatusOk

`func (o *MicrosoftGraphSubscribedSku) GetCapabilityStatusOk() (*string, bool)`

GetCapabilityStatusOk returns a tuple with the CapabilityStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilityStatus

`func (o *MicrosoftGraphSubscribedSku) SetCapabilityStatus(v string)`

SetCapabilityStatus sets CapabilityStatus field to given value.

### HasCapabilityStatus

`func (o *MicrosoftGraphSubscribedSku) HasCapabilityStatus() bool`

HasCapabilityStatus returns a boolean if a field has been set.

### SetCapabilityStatusNil

`func (o *MicrosoftGraphSubscribedSku) SetCapabilityStatusNil(b bool)`

 SetCapabilityStatusNil sets the value for CapabilityStatus to be an explicit nil

### UnsetCapabilityStatus
`func (o *MicrosoftGraphSubscribedSku) UnsetCapabilityStatus()`

UnsetCapabilityStatus ensures that no value is present for CapabilityStatus, not even an explicit nil
### GetConsumedUnits

`func (o *MicrosoftGraphSubscribedSku) GetConsumedUnits() int32`

GetConsumedUnits returns the ConsumedUnits field if non-nil, zero value otherwise.

### GetConsumedUnitsOk

`func (o *MicrosoftGraphSubscribedSku) GetConsumedUnitsOk() (*int32, bool)`

GetConsumedUnitsOk returns a tuple with the ConsumedUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumedUnits

`func (o *MicrosoftGraphSubscribedSku) SetConsumedUnits(v int32)`

SetConsumedUnits sets ConsumedUnits field to given value.

### HasConsumedUnits

`func (o *MicrosoftGraphSubscribedSku) HasConsumedUnits() bool`

HasConsumedUnits returns a boolean if a field has been set.

### SetConsumedUnitsNil

`func (o *MicrosoftGraphSubscribedSku) SetConsumedUnitsNil(b bool)`

 SetConsumedUnitsNil sets the value for ConsumedUnits to be an explicit nil

### UnsetConsumedUnits
`func (o *MicrosoftGraphSubscribedSku) UnsetConsumedUnits()`

UnsetConsumedUnits ensures that no value is present for ConsumedUnits, not even an explicit nil
### GetPrepaidUnits

`func (o *MicrosoftGraphSubscribedSku) GetPrepaidUnits() AnyOfmicrosoftGraphLicenseUnitsDetail`

GetPrepaidUnits returns the PrepaidUnits field if non-nil, zero value otherwise.

### GetPrepaidUnitsOk

`func (o *MicrosoftGraphSubscribedSku) GetPrepaidUnitsOk() (*AnyOfmicrosoftGraphLicenseUnitsDetail, bool)`

GetPrepaidUnitsOk returns a tuple with the PrepaidUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrepaidUnits

`func (o *MicrosoftGraphSubscribedSku) SetPrepaidUnits(v AnyOfmicrosoftGraphLicenseUnitsDetail)`

SetPrepaidUnits sets PrepaidUnits field to given value.

### HasPrepaidUnits

`func (o *MicrosoftGraphSubscribedSku) HasPrepaidUnits() bool`

HasPrepaidUnits returns a boolean if a field has been set.

### SetPrepaidUnitsNil

`func (o *MicrosoftGraphSubscribedSku) SetPrepaidUnitsNil(b bool)`

 SetPrepaidUnitsNil sets the value for PrepaidUnits to be an explicit nil

### UnsetPrepaidUnits
`func (o *MicrosoftGraphSubscribedSku) UnsetPrepaidUnits()`

UnsetPrepaidUnits ensures that no value is present for PrepaidUnits, not even an explicit nil
### GetServicePlans

`func (o *MicrosoftGraphSubscribedSku) GetServicePlans() []MicrosoftGraphServicePlanInfo`

GetServicePlans returns the ServicePlans field if non-nil, zero value otherwise.

### GetServicePlansOk

`func (o *MicrosoftGraphSubscribedSku) GetServicePlansOk() (*[]MicrosoftGraphServicePlanInfo, bool)`

GetServicePlansOk returns a tuple with the ServicePlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePlans

`func (o *MicrosoftGraphSubscribedSku) SetServicePlans(v []MicrosoftGraphServicePlanInfo)`

SetServicePlans sets ServicePlans field to given value.

### HasServicePlans

`func (o *MicrosoftGraphSubscribedSku) HasServicePlans() bool`

HasServicePlans returns a boolean if a field has been set.

### GetSkuId

`func (o *MicrosoftGraphSubscribedSku) GetSkuId() string`

GetSkuId returns the SkuId field if non-nil, zero value otherwise.

### GetSkuIdOk

`func (o *MicrosoftGraphSubscribedSku) GetSkuIdOk() (*string, bool)`

GetSkuIdOk returns a tuple with the SkuId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkuId

`func (o *MicrosoftGraphSubscribedSku) SetSkuId(v string)`

SetSkuId sets SkuId field to given value.

### HasSkuId

`func (o *MicrosoftGraphSubscribedSku) HasSkuId() bool`

HasSkuId returns a boolean if a field has been set.

### SetSkuIdNil

`func (o *MicrosoftGraphSubscribedSku) SetSkuIdNil(b bool)`

 SetSkuIdNil sets the value for SkuId to be an explicit nil

### UnsetSkuId
`func (o *MicrosoftGraphSubscribedSku) UnsetSkuId()`

UnsetSkuId ensures that no value is present for SkuId, not even an explicit nil
### GetSkuPartNumber

`func (o *MicrosoftGraphSubscribedSku) GetSkuPartNumber() string`

GetSkuPartNumber returns the SkuPartNumber field if non-nil, zero value otherwise.

### GetSkuPartNumberOk

`func (o *MicrosoftGraphSubscribedSku) GetSkuPartNumberOk() (*string, bool)`

GetSkuPartNumberOk returns a tuple with the SkuPartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkuPartNumber

`func (o *MicrosoftGraphSubscribedSku) SetSkuPartNumber(v string)`

SetSkuPartNumber sets SkuPartNumber field to given value.

### HasSkuPartNumber

`func (o *MicrosoftGraphSubscribedSku) HasSkuPartNumber() bool`

HasSkuPartNumber returns a boolean if a field has been set.

### SetSkuPartNumberNil

`func (o *MicrosoftGraphSubscribedSku) SetSkuPartNumberNil(b bool)`

 SetSkuPartNumberNil sets the value for SkuPartNumber to be an explicit nil

### UnsetSkuPartNumber
`func (o *MicrosoftGraphSubscribedSku) UnsetSkuPartNumber()`

UnsetSkuPartNumber ensures that no value is present for SkuPartNumber, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


