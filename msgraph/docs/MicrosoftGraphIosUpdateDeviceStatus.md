# MicrosoftGraphIosUpdateDeviceStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**ComplianceGracePeriodExpirationDateTime** | Pointer to **time.Time** | The DateTime when device compliance grace period expires | [optional] 
**DeviceDisplayName** | Pointer to **NullableString** | Device name of the DevicePolicyStatus. | [optional] 
**DeviceId** | Pointer to **NullableString** | The device id that is being reported. | [optional] 
**DeviceModel** | Pointer to **NullableString** | The device model that is being reported | [optional] 
**InstallStatus** | Pointer to [**AnyOfmicrosoftGraphIosUpdatesInstallStatus**](anyOf&lt;microsoft.graph.iosUpdatesInstallStatus&gt;.md) | The installation status of the policy report. Possible values are: success, available, idle, unknown, downloading, downloadFailed, downloadRequiresComputer, downloadInsufficientSpace, downloadInsufficientPower, downloadInsufficientNetwork, installing, installInsufficientSpace, installInsufficientPower, installPhoneCallInProgress, installFailed, notSupportedOperation, sharedDeviceUserLoggedInError, deviceOsHigherThanDesiredOsVersion. | [optional] 
**LastReportedDateTime** | Pointer to **time.Time** | Last modified date time of the policy report. | [optional] 
**OsVersion** | Pointer to **NullableString** | The device version that is being reported. | [optional] 
**Status** | Pointer to [**AnyOfmicrosoftGraphComplianceStatus**](anyOf&lt;microsoft.graph.complianceStatus&gt;.md) | Compliance status of the policy report. Possible values are: unknown, notApplicable, compliant, remediated, nonCompliant, error, conflict, notAssigned. | [optional] 
**UserId** | Pointer to **NullableString** | The User id that is being reported. | [optional] 
**UserName** | Pointer to **NullableString** | The User Name that is being reported | [optional] 
**UserPrincipalName** | Pointer to **NullableString** | UserPrincipalName. | [optional] 

## Methods

### NewMicrosoftGraphIosUpdateDeviceStatus

`func NewMicrosoftGraphIosUpdateDeviceStatus() *MicrosoftGraphIosUpdateDeviceStatus`

NewMicrosoftGraphIosUpdateDeviceStatus instantiates a new MicrosoftGraphIosUpdateDeviceStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphIosUpdateDeviceStatusWithDefaults

`func NewMicrosoftGraphIosUpdateDeviceStatusWithDefaults() *MicrosoftGraphIosUpdateDeviceStatus`

NewMicrosoftGraphIosUpdateDeviceStatusWithDefaults instantiates a new MicrosoftGraphIosUpdateDeviceStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphIosUpdateDeviceStatus) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphIosUpdateDeviceStatus) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphIosUpdateDeviceStatus) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphIosUpdateDeviceStatus) HasId() bool`

HasId returns a boolean if a field has been set.

### GetComplianceGracePeriodExpirationDateTime

`func (o *MicrosoftGraphIosUpdateDeviceStatus) GetComplianceGracePeriodExpirationDateTime() time.Time`

GetComplianceGracePeriodExpirationDateTime returns the ComplianceGracePeriodExpirationDateTime field if non-nil, zero value otherwise.

### GetComplianceGracePeriodExpirationDateTimeOk

`func (o *MicrosoftGraphIosUpdateDeviceStatus) GetComplianceGracePeriodExpirationDateTimeOk() (*time.Time, bool)`

GetComplianceGracePeriodExpirationDateTimeOk returns a tuple with the ComplianceGracePeriodExpirationDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplianceGracePeriodExpirationDateTime

`func (o *MicrosoftGraphIosUpdateDeviceStatus) SetComplianceGracePeriodExpirationDateTime(v time.Time)`

SetComplianceGracePeriodExpirationDateTime sets ComplianceGracePeriodExpirationDateTime field to given value.

### HasComplianceGracePeriodExpirationDateTime

`func (o *MicrosoftGraphIosUpdateDeviceStatus) HasComplianceGracePeriodExpirationDateTime() bool`

HasComplianceGracePeriodExpirationDateTime returns a boolean if a field has been set.

### GetDeviceDisplayName

`func (o *MicrosoftGraphIosUpdateDeviceStatus) GetDeviceDisplayName() string`

GetDeviceDisplayName returns the DeviceDisplayName field if non-nil, zero value otherwise.

### GetDeviceDisplayNameOk

`func (o *MicrosoftGraphIosUpdateDeviceStatus) GetDeviceDisplayNameOk() (*string, bool)`

GetDeviceDisplayNameOk returns a tuple with the DeviceDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceDisplayName

`func (o *MicrosoftGraphIosUpdateDeviceStatus) SetDeviceDisplayName(v string)`

SetDeviceDisplayName sets DeviceDisplayName field to given value.

### HasDeviceDisplayName

`func (o *MicrosoftGraphIosUpdateDeviceStatus) HasDeviceDisplayName() bool`

HasDeviceDisplayName returns a boolean if a field has been set.

### SetDeviceDisplayNameNil

`func (o *MicrosoftGraphIosUpdateDeviceStatus) SetDeviceDisplayNameNil(b bool)`

 SetDeviceDisplayNameNil sets the value for DeviceDisplayName to be an explicit nil

### UnsetDeviceDisplayName
`func (o *MicrosoftGraphIosUpdateDeviceStatus) UnsetDeviceDisplayName()`

UnsetDeviceDisplayName ensures that no value is present for DeviceDisplayName, not even an explicit nil
### GetDeviceId

`func (o *MicrosoftGraphIosUpdateDeviceStatus) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *MicrosoftGraphIosUpdateDeviceStatus) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *MicrosoftGraphIosUpdateDeviceStatus) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *MicrosoftGraphIosUpdateDeviceStatus) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### SetDeviceIdNil

`func (o *MicrosoftGraphIosUpdateDeviceStatus) SetDeviceIdNil(b bool)`

 SetDeviceIdNil sets the value for DeviceId to be an explicit nil

### UnsetDeviceId
`func (o *MicrosoftGraphIosUpdateDeviceStatus) UnsetDeviceId()`

UnsetDeviceId ensures that no value is present for DeviceId, not even an explicit nil
### GetDeviceModel

`func (o *MicrosoftGraphIosUpdateDeviceStatus) GetDeviceModel() string`

GetDeviceModel returns the DeviceModel field if non-nil, zero value otherwise.

### GetDeviceModelOk

`func (o *MicrosoftGraphIosUpdateDeviceStatus) GetDeviceModelOk() (*string, bool)`

GetDeviceModelOk returns a tuple with the DeviceModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceModel

`func (o *MicrosoftGraphIosUpdateDeviceStatus) SetDeviceModel(v string)`

SetDeviceModel sets DeviceModel field to given value.

### HasDeviceModel

`func (o *MicrosoftGraphIosUpdateDeviceStatus) HasDeviceModel() bool`

HasDeviceModel returns a boolean if a field has been set.

### SetDeviceModelNil

`func (o *MicrosoftGraphIosUpdateDeviceStatus) SetDeviceModelNil(b bool)`

 SetDeviceModelNil sets the value for DeviceModel to be an explicit nil

### UnsetDeviceModel
`func (o *MicrosoftGraphIosUpdateDeviceStatus) UnsetDeviceModel()`

UnsetDeviceModel ensures that no value is present for DeviceModel, not even an explicit nil
### GetInstallStatus

`func (o *MicrosoftGraphIosUpdateDeviceStatus) GetInstallStatus() AnyOfmicrosoftGraphIosUpdatesInstallStatus`

GetInstallStatus returns the InstallStatus field if non-nil, zero value otherwise.

### GetInstallStatusOk

`func (o *MicrosoftGraphIosUpdateDeviceStatus) GetInstallStatusOk() (*AnyOfmicrosoftGraphIosUpdatesInstallStatus, bool)`

GetInstallStatusOk returns a tuple with the InstallStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallStatus

`func (o *MicrosoftGraphIosUpdateDeviceStatus) SetInstallStatus(v AnyOfmicrosoftGraphIosUpdatesInstallStatus)`

SetInstallStatus sets InstallStatus field to given value.

### HasInstallStatus

`func (o *MicrosoftGraphIosUpdateDeviceStatus) HasInstallStatus() bool`

HasInstallStatus returns a boolean if a field has been set.

### SetInstallStatusNil

`func (o *MicrosoftGraphIosUpdateDeviceStatus) SetInstallStatusNil(b bool)`

 SetInstallStatusNil sets the value for InstallStatus to be an explicit nil

### UnsetInstallStatus
`func (o *MicrosoftGraphIosUpdateDeviceStatus) UnsetInstallStatus()`

UnsetInstallStatus ensures that no value is present for InstallStatus, not even an explicit nil
### GetLastReportedDateTime

`func (o *MicrosoftGraphIosUpdateDeviceStatus) GetLastReportedDateTime() time.Time`

GetLastReportedDateTime returns the LastReportedDateTime field if non-nil, zero value otherwise.

### GetLastReportedDateTimeOk

`func (o *MicrosoftGraphIosUpdateDeviceStatus) GetLastReportedDateTimeOk() (*time.Time, bool)`

GetLastReportedDateTimeOk returns a tuple with the LastReportedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastReportedDateTime

`func (o *MicrosoftGraphIosUpdateDeviceStatus) SetLastReportedDateTime(v time.Time)`

SetLastReportedDateTime sets LastReportedDateTime field to given value.

### HasLastReportedDateTime

`func (o *MicrosoftGraphIosUpdateDeviceStatus) HasLastReportedDateTime() bool`

HasLastReportedDateTime returns a boolean if a field has been set.

### GetOsVersion

`func (o *MicrosoftGraphIosUpdateDeviceStatus) GetOsVersion() string`

GetOsVersion returns the OsVersion field if non-nil, zero value otherwise.

### GetOsVersionOk

`func (o *MicrosoftGraphIosUpdateDeviceStatus) GetOsVersionOk() (*string, bool)`

GetOsVersionOk returns a tuple with the OsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsVersion

`func (o *MicrosoftGraphIosUpdateDeviceStatus) SetOsVersion(v string)`

SetOsVersion sets OsVersion field to given value.

### HasOsVersion

`func (o *MicrosoftGraphIosUpdateDeviceStatus) HasOsVersion() bool`

HasOsVersion returns a boolean if a field has been set.

### SetOsVersionNil

`func (o *MicrosoftGraphIosUpdateDeviceStatus) SetOsVersionNil(b bool)`

 SetOsVersionNil sets the value for OsVersion to be an explicit nil

### UnsetOsVersion
`func (o *MicrosoftGraphIosUpdateDeviceStatus) UnsetOsVersion()`

UnsetOsVersion ensures that no value is present for OsVersion, not even an explicit nil
### GetStatus

`func (o *MicrosoftGraphIosUpdateDeviceStatus) GetStatus() AnyOfmicrosoftGraphComplianceStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MicrosoftGraphIosUpdateDeviceStatus) GetStatusOk() (*AnyOfmicrosoftGraphComplianceStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MicrosoftGraphIosUpdateDeviceStatus) SetStatus(v AnyOfmicrosoftGraphComplianceStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MicrosoftGraphIosUpdateDeviceStatus) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *MicrosoftGraphIosUpdateDeviceStatus) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *MicrosoftGraphIosUpdateDeviceStatus) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetUserId

`func (o *MicrosoftGraphIosUpdateDeviceStatus) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *MicrosoftGraphIosUpdateDeviceStatus) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *MicrosoftGraphIosUpdateDeviceStatus) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *MicrosoftGraphIosUpdateDeviceStatus) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *MicrosoftGraphIosUpdateDeviceStatus) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *MicrosoftGraphIosUpdateDeviceStatus) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil
### GetUserName

`func (o *MicrosoftGraphIosUpdateDeviceStatus) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *MicrosoftGraphIosUpdateDeviceStatus) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *MicrosoftGraphIosUpdateDeviceStatus) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *MicrosoftGraphIosUpdateDeviceStatus) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### SetUserNameNil

`func (o *MicrosoftGraphIosUpdateDeviceStatus) SetUserNameNil(b bool)`

 SetUserNameNil sets the value for UserName to be an explicit nil

### UnsetUserName
`func (o *MicrosoftGraphIosUpdateDeviceStatus) UnsetUserName()`

UnsetUserName ensures that no value is present for UserName, not even an explicit nil
### GetUserPrincipalName

`func (o *MicrosoftGraphIosUpdateDeviceStatus) GetUserPrincipalName() string`

GetUserPrincipalName returns the UserPrincipalName field if non-nil, zero value otherwise.

### GetUserPrincipalNameOk

`func (o *MicrosoftGraphIosUpdateDeviceStatus) GetUserPrincipalNameOk() (*string, bool)`

GetUserPrincipalNameOk returns a tuple with the UserPrincipalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPrincipalName

`func (o *MicrosoftGraphIosUpdateDeviceStatus) SetUserPrincipalName(v string)`

SetUserPrincipalName sets UserPrincipalName field to given value.

### HasUserPrincipalName

`func (o *MicrosoftGraphIosUpdateDeviceStatus) HasUserPrincipalName() bool`

HasUserPrincipalName returns a boolean if a field has been set.

### SetUserPrincipalNameNil

`func (o *MicrosoftGraphIosUpdateDeviceStatus) SetUserPrincipalNameNil(b bool)`

 SetUserPrincipalNameNil sets the value for UserPrincipalName to be an explicit nil

### UnsetUserPrincipalName
`func (o *MicrosoftGraphIosUpdateDeviceStatus) UnsetUserPrincipalName()`

UnsetUserPrincipalName ensures that no value is present for UserPrincipalName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


