# MicrosoftGraphLicenseUnitsDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **NullableInt32** | The number of units that are enabled for the active subscription of the service SKU. | [optional] 
**Suspended** | Pointer to **NullableInt32** | The number of units that are suspended because the subscription of the service SKU has been cancelled. The units cannot be assigned but can still be reactivated before they are deleted. | [optional] 
**Warning** | Pointer to **NullableInt32** | The number of units that are in warning status. When the subscription of the service SKU has expired, the customer has a grace period to renew their subscription before it is cancelled (moved to a suspended state). | [optional] 

## Methods

### NewMicrosoftGraphLicenseUnitsDetail

`func NewMicrosoftGraphLicenseUnitsDetail() *MicrosoftGraphLicenseUnitsDetail`

NewMicrosoftGraphLicenseUnitsDetail instantiates a new MicrosoftGraphLicenseUnitsDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphLicenseUnitsDetailWithDefaults

`func NewMicrosoftGraphLicenseUnitsDetailWithDefaults() *MicrosoftGraphLicenseUnitsDetail`

NewMicrosoftGraphLicenseUnitsDetailWithDefaults instantiates a new MicrosoftGraphLicenseUnitsDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *MicrosoftGraphLicenseUnitsDetail) GetEnabled() int32`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *MicrosoftGraphLicenseUnitsDetail) GetEnabledOk() (*int32, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *MicrosoftGraphLicenseUnitsDetail) SetEnabled(v int32)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *MicrosoftGraphLicenseUnitsDetail) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### SetEnabledNil

`func (o *MicrosoftGraphLicenseUnitsDetail) SetEnabledNil(b bool)`

 SetEnabledNil sets the value for Enabled to be an explicit nil

### UnsetEnabled
`func (o *MicrosoftGraphLicenseUnitsDetail) UnsetEnabled()`

UnsetEnabled ensures that no value is present for Enabled, not even an explicit nil
### GetSuspended

`func (o *MicrosoftGraphLicenseUnitsDetail) GetSuspended() int32`

GetSuspended returns the Suspended field if non-nil, zero value otherwise.

### GetSuspendedOk

`func (o *MicrosoftGraphLicenseUnitsDetail) GetSuspendedOk() (*int32, bool)`

GetSuspendedOk returns a tuple with the Suspended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuspended

`func (o *MicrosoftGraphLicenseUnitsDetail) SetSuspended(v int32)`

SetSuspended sets Suspended field to given value.

### HasSuspended

`func (o *MicrosoftGraphLicenseUnitsDetail) HasSuspended() bool`

HasSuspended returns a boolean if a field has been set.

### SetSuspendedNil

`func (o *MicrosoftGraphLicenseUnitsDetail) SetSuspendedNil(b bool)`

 SetSuspendedNil sets the value for Suspended to be an explicit nil

### UnsetSuspended
`func (o *MicrosoftGraphLicenseUnitsDetail) UnsetSuspended()`

UnsetSuspended ensures that no value is present for Suspended, not even an explicit nil
### GetWarning

`func (o *MicrosoftGraphLicenseUnitsDetail) GetWarning() int32`

GetWarning returns the Warning field if non-nil, zero value otherwise.

### GetWarningOk

`func (o *MicrosoftGraphLicenseUnitsDetail) GetWarningOk() (*int32, bool)`

GetWarningOk returns a tuple with the Warning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarning

`func (o *MicrosoftGraphLicenseUnitsDetail) SetWarning(v int32)`

SetWarning sets Warning field to given value.

### HasWarning

`func (o *MicrosoftGraphLicenseUnitsDetail) HasWarning() bool`

HasWarning returns a boolean if a field has been set.

### SetWarningNil

`func (o *MicrosoftGraphLicenseUnitsDetail) SetWarningNil(b bool)`

 SetWarningNil sets the value for Warning to be an explicit nil

### UnsetWarning
`func (o *MicrosoftGraphLicenseUnitsDetail) UnsetWarning()`

UnsetWarning ensures that no value is present for Warning, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


