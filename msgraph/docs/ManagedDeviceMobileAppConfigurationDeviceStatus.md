# ManagedDeviceMobileAppConfigurationDeviceStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ComplianceGracePeriodExpirationDateTime** | Pointer to **time.Time** | The DateTime when device compliance grace period expires | [optional] 
**DeviceDisplayName** | Pointer to **NullableString** | Device name of the DevicePolicyStatus. | [optional] 
**DeviceModel** | Pointer to **NullableString** | The device model that is being reported | [optional] 
**LastReportedDateTime** | Pointer to **time.Time** | Last modified date time of the policy report. | [optional] 
**Status** | Pointer to [**AnyOfmicrosoftGraphComplianceStatus**](anyOf&lt;microsoft.graph.complianceStatus&gt;.md) | Compliance status of the policy report. Possible values are: unknown, notApplicable, compliant, remediated, nonCompliant, error, conflict, notAssigned. | [optional] 
**UserName** | Pointer to **NullableString** | The User Name that is being reported | [optional] 
**UserPrincipalName** | Pointer to **NullableString** | UserPrincipalName. | [optional] 

## Methods

### NewManagedDeviceMobileAppConfigurationDeviceStatus

`func NewManagedDeviceMobileAppConfigurationDeviceStatus() *ManagedDeviceMobileAppConfigurationDeviceStatus`

NewManagedDeviceMobileAppConfigurationDeviceStatus instantiates a new ManagedDeviceMobileAppConfigurationDeviceStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagedDeviceMobileAppConfigurationDeviceStatusWithDefaults

`func NewManagedDeviceMobileAppConfigurationDeviceStatusWithDefaults() *ManagedDeviceMobileAppConfigurationDeviceStatus`

NewManagedDeviceMobileAppConfigurationDeviceStatusWithDefaults instantiates a new ManagedDeviceMobileAppConfigurationDeviceStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComplianceGracePeriodExpirationDateTime

`func (o *ManagedDeviceMobileAppConfigurationDeviceStatus) GetComplianceGracePeriodExpirationDateTime() time.Time`

GetComplianceGracePeriodExpirationDateTime returns the ComplianceGracePeriodExpirationDateTime field if non-nil, zero value otherwise.

### GetComplianceGracePeriodExpirationDateTimeOk

`func (o *ManagedDeviceMobileAppConfigurationDeviceStatus) GetComplianceGracePeriodExpirationDateTimeOk() (*time.Time, bool)`

GetComplianceGracePeriodExpirationDateTimeOk returns a tuple with the ComplianceGracePeriodExpirationDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplianceGracePeriodExpirationDateTime

`func (o *ManagedDeviceMobileAppConfigurationDeviceStatus) SetComplianceGracePeriodExpirationDateTime(v time.Time)`

SetComplianceGracePeriodExpirationDateTime sets ComplianceGracePeriodExpirationDateTime field to given value.

### HasComplianceGracePeriodExpirationDateTime

`func (o *ManagedDeviceMobileAppConfigurationDeviceStatus) HasComplianceGracePeriodExpirationDateTime() bool`

HasComplianceGracePeriodExpirationDateTime returns a boolean if a field has been set.

### GetDeviceDisplayName

`func (o *ManagedDeviceMobileAppConfigurationDeviceStatus) GetDeviceDisplayName() string`

GetDeviceDisplayName returns the DeviceDisplayName field if non-nil, zero value otherwise.

### GetDeviceDisplayNameOk

`func (o *ManagedDeviceMobileAppConfigurationDeviceStatus) GetDeviceDisplayNameOk() (*string, bool)`

GetDeviceDisplayNameOk returns a tuple with the DeviceDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceDisplayName

`func (o *ManagedDeviceMobileAppConfigurationDeviceStatus) SetDeviceDisplayName(v string)`

SetDeviceDisplayName sets DeviceDisplayName field to given value.

### HasDeviceDisplayName

`func (o *ManagedDeviceMobileAppConfigurationDeviceStatus) HasDeviceDisplayName() bool`

HasDeviceDisplayName returns a boolean if a field has been set.

### SetDeviceDisplayNameNil

`func (o *ManagedDeviceMobileAppConfigurationDeviceStatus) SetDeviceDisplayNameNil(b bool)`

 SetDeviceDisplayNameNil sets the value for DeviceDisplayName to be an explicit nil

### UnsetDeviceDisplayName
`func (o *ManagedDeviceMobileAppConfigurationDeviceStatus) UnsetDeviceDisplayName()`

UnsetDeviceDisplayName ensures that no value is present for DeviceDisplayName, not even an explicit nil
### GetDeviceModel

`func (o *ManagedDeviceMobileAppConfigurationDeviceStatus) GetDeviceModel() string`

GetDeviceModel returns the DeviceModel field if non-nil, zero value otherwise.

### GetDeviceModelOk

`func (o *ManagedDeviceMobileAppConfigurationDeviceStatus) GetDeviceModelOk() (*string, bool)`

GetDeviceModelOk returns a tuple with the DeviceModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceModel

`func (o *ManagedDeviceMobileAppConfigurationDeviceStatus) SetDeviceModel(v string)`

SetDeviceModel sets DeviceModel field to given value.

### HasDeviceModel

`func (o *ManagedDeviceMobileAppConfigurationDeviceStatus) HasDeviceModel() bool`

HasDeviceModel returns a boolean if a field has been set.

### SetDeviceModelNil

`func (o *ManagedDeviceMobileAppConfigurationDeviceStatus) SetDeviceModelNil(b bool)`

 SetDeviceModelNil sets the value for DeviceModel to be an explicit nil

### UnsetDeviceModel
`func (o *ManagedDeviceMobileAppConfigurationDeviceStatus) UnsetDeviceModel()`

UnsetDeviceModel ensures that no value is present for DeviceModel, not even an explicit nil
### GetLastReportedDateTime

`func (o *ManagedDeviceMobileAppConfigurationDeviceStatus) GetLastReportedDateTime() time.Time`

GetLastReportedDateTime returns the LastReportedDateTime field if non-nil, zero value otherwise.

### GetLastReportedDateTimeOk

`func (o *ManagedDeviceMobileAppConfigurationDeviceStatus) GetLastReportedDateTimeOk() (*time.Time, bool)`

GetLastReportedDateTimeOk returns a tuple with the LastReportedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastReportedDateTime

`func (o *ManagedDeviceMobileAppConfigurationDeviceStatus) SetLastReportedDateTime(v time.Time)`

SetLastReportedDateTime sets LastReportedDateTime field to given value.

### HasLastReportedDateTime

`func (o *ManagedDeviceMobileAppConfigurationDeviceStatus) HasLastReportedDateTime() bool`

HasLastReportedDateTime returns a boolean if a field has been set.

### GetStatus

`func (o *ManagedDeviceMobileAppConfigurationDeviceStatus) GetStatus() AnyOfmicrosoftGraphComplianceStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ManagedDeviceMobileAppConfigurationDeviceStatus) GetStatusOk() (*AnyOfmicrosoftGraphComplianceStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ManagedDeviceMobileAppConfigurationDeviceStatus) SetStatus(v AnyOfmicrosoftGraphComplianceStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ManagedDeviceMobileAppConfigurationDeviceStatus) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *ManagedDeviceMobileAppConfigurationDeviceStatus) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *ManagedDeviceMobileAppConfigurationDeviceStatus) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetUserName

`func (o *ManagedDeviceMobileAppConfigurationDeviceStatus) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *ManagedDeviceMobileAppConfigurationDeviceStatus) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *ManagedDeviceMobileAppConfigurationDeviceStatus) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *ManagedDeviceMobileAppConfigurationDeviceStatus) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### SetUserNameNil

`func (o *ManagedDeviceMobileAppConfigurationDeviceStatus) SetUserNameNil(b bool)`

 SetUserNameNil sets the value for UserName to be an explicit nil

### UnsetUserName
`func (o *ManagedDeviceMobileAppConfigurationDeviceStatus) UnsetUserName()`

UnsetUserName ensures that no value is present for UserName, not even an explicit nil
### GetUserPrincipalName

`func (o *ManagedDeviceMobileAppConfigurationDeviceStatus) GetUserPrincipalName() string`

GetUserPrincipalName returns the UserPrincipalName field if non-nil, zero value otherwise.

### GetUserPrincipalNameOk

`func (o *ManagedDeviceMobileAppConfigurationDeviceStatus) GetUserPrincipalNameOk() (*string, bool)`

GetUserPrincipalNameOk returns a tuple with the UserPrincipalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPrincipalName

`func (o *ManagedDeviceMobileAppConfigurationDeviceStatus) SetUserPrincipalName(v string)`

SetUserPrincipalName sets UserPrincipalName field to given value.

### HasUserPrincipalName

`func (o *ManagedDeviceMobileAppConfigurationDeviceStatus) HasUserPrincipalName() bool`

HasUserPrincipalName returns a boolean if a field has been set.

### SetUserPrincipalNameNil

`func (o *ManagedDeviceMobileAppConfigurationDeviceStatus) SetUserPrincipalNameNil(b bool)`

 SetUserPrincipalNameNil sets the value for UserPrincipalName to be an explicit nil

### UnsetUserPrincipalName
`func (o *ManagedDeviceMobileAppConfigurationDeviceStatus) UnsetUserPrincipalName()`

UnsetUserPrincipalName ensures that no value is present for UserPrincipalName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


