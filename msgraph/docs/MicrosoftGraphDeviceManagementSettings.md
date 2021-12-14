# MicrosoftGraphDeviceManagementSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceComplianceCheckinThresholdDays** | Pointer to **int32** | The number of days a device is allowed to go without checking in to remain compliant. | [optional] 
**IsScheduledActionEnabled** | Pointer to **bool** | Is feature enabled or not for scheduled action for rule. | [optional] 
**SecureByDefault** | Pointer to **bool** | Device should be noncompliant when there is no compliance policy targeted when this is true | [optional] 

## Methods

### NewMicrosoftGraphDeviceManagementSettings

`func NewMicrosoftGraphDeviceManagementSettings() *MicrosoftGraphDeviceManagementSettings`

NewMicrosoftGraphDeviceManagementSettings instantiates a new MicrosoftGraphDeviceManagementSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphDeviceManagementSettingsWithDefaults

`func NewMicrosoftGraphDeviceManagementSettingsWithDefaults() *MicrosoftGraphDeviceManagementSettings`

NewMicrosoftGraphDeviceManagementSettingsWithDefaults instantiates a new MicrosoftGraphDeviceManagementSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceComplianceCheckinThresholdDays

`func (o *MicrosoftGraphDeviceManagementSettings) GetDeviceComplianceCheckinThresholdDays() int32`

GetDeviceComplianceCheckinThresholdDays returns the DeviceComplianceCheckinThresholdDays field if non-nil, zero value otherwise.

### GetDeviceComplianceCheckinThresholdDaysOk

`func (o *MicrosoftGraphDeviceManagementSettings) GetDeviceComplianceCheckinThresholdDaysOk() (*int32, bool)`

GetDeviceComplianceCheckinThresholdDaysOk returns a tuple with the DeviceComplianceCheckinThresholdDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceComplianceCheckinThresholdDays

`func (o *MicrosoftGraphDeviceManagementSettings) SetDeviceComplianceCheckinThresholdDays(v int32)`

SetDeviceComplianceCheckinThresholdDays sets DeviceComplianceCheckinThresholdDays field to given value.

### HasDeviceComplianceCheckinThresholdDays

`func (o *MicrosoftGraphDeviceManagementSettings) HasDeviceComplianceCheckinThresholdDays() bool`

HasDeviceComplianceCheckinThresholdDays returns a boolean if a field has been set.

### GetIsScheduledActionEnabled

`func (o *MicrosoftGraphDeviceManagementSettings) GetIsScheduledActionEnabled() bool`

GetIsScheduledActionEnabled returns the IsScheduledActionEnabled field if non-nil, zero value otherwise.

### GetIsScheduledActionEnabledOk

`func (o *MicrosoftGraphDeviceManagementSettings) GetIsScheduledActionEnabledOk() (*bool, bool)`

GetIsScheduledActionEnabledOk returns a tuple with the IsScheduledActionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsScheduledActionEnabled

`func (o *MicrosoftGraphDeviceManagementSettings) SetIsScheduledActionEnabled(v bool)`

SetIsScheduledActionEnabled sets IsScheduledActionEnabled field to given value.

### HasIsScheduledActionEnabled

`func (o *MicrosoftGraphDeviceManagementSettings) HasIsScheduledActionEnabled() bool`

HasIsScheduledActionEnabled returns a boolean if a field has been set.

### GetSecureByDefault

`func (o *MicrosoftGraphDeviceManagementSettings) GetSecureByDefault() bool`

GetSecureByDefault returns the SecureByDefault field if non-nil, zero value otherwise.

### GetSecureByDefaultOk

`func (o *MicrosoftGraphDeviceManagementSettings) GetSecureByDefaultOk() (*bool, bool)`

GetSecureByDefaultOk returns a tuple with the SecureByDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureByDefault

`func (o *MicrosoftGraphDeviceManagementSettings) SetSecureByDefault(v bool)`

SetSecureByDefault sets SecureByDefault field to given value.

### HasSecureByDefault

`func (o *MicrosoftGraphDeviceManagementSettings) HasSecureByDefault() bool`

HasSecureByDefault returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


