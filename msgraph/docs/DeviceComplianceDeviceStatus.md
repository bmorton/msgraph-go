# DeviceComplianceDeviceStatus

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

### NewDeviceComplianceDeviceStatus

`func NewDeviceComplianceDeviceStatus() *DeviceComplianceDeviceStatus`

NewDeviceComplianceDeviceStatus instantiates a new DeviceComplianceDeviceStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceComplianceDeviceStatusWithDefaults

`func NewDeviceComplianceDeviceStatusWithDefaults() *DeviceComplianceDeviceStatus`

NewDeviceComplianceDeviceStatusWithDefaults instantiates a new DeviceComplianceDeviceStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComplianceGracePeriodExpirationDateTime

`func (o *DeviceComplianceDeviceStatus) GetComplianceGracePeriodExpirationDateTime() time.Time`

GetComplianceGracePeriodExpirationDateTime returns the ComplianceGracePeriodExpirationDateTime field if non-nil, zero value otherwise.

### GetComplianceGracePeriodExpirationDateTimeOk

`func (o *DeviceComplianceDeviceStatus) GetComplianceGracePeriodExpirationDateTimeOk() (*time.Time, bool)`

GetComplianceGracePeriodExpirationDateTimeOk returns a tuple with the ComplianceGracePeriodExpirationDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplianceGracePeriodExpirationDateTime

`func (o *DeviceComplianceDeviceStatus) SetComplianceGracePeriodExpirationDateTime(v time.Time)`

SetComplianceGracePeriodExpirationDateTime sets ComplianceGracePeriodExpirationDateTime field to given value.

### HasComplianceGracePeriodExpirationDateTime

`func (o *DeviceComplianceDeviceStatus) HasComplianceGracePeriodExpirationDateTime() bool`

HasComplianceGracePeriodExpirationDateTime returns a boolean if a field has been set.

### GetDeviceDisplayName

`func (o *DeviceComplianceDeviceStatus) GetDeviceDisplayName() string`

GetDeviceDisplayName returns the DeviceDisplayName field if non-nil, zero value otherwise.

### GetDeviceDisplayNameOk

`func (o *DeviceComplianceDeviceStatus) GetDeviceDisplayNameOk() (*string, bool)`

GetDeviceDisplayNameOk returns a tuple with the DeviceDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceDisplayName

`func (o *DeviceComplianceDeviceStatus) SetDeviceDisplayName(v string)`

SetDeviceDisplayName sets DeviceDisplayName field to given value.

### HasDeviceDisplayName

`func (o *DeviceComplianceDeviceStatus) HasDeviceDisplayName() bool`

HasDeviceDisplayName returns a boolean if a field has been set.

### SetDeviceDisplayNameNil

`func (o *DeviceComplianceDeviceStatus) SetDeviceDisplayNameNil(b bool)`

 SetDeviceDisplayNameNil sets the value for DeviceDisplayName to be an explicit nil

### UnsetDeviceDisplayName
`func (o *DeviceComplianceDeviceStatus) UnsetDeviceDisplayName()`

UnsetDeviceDisplayName ensures that no value is present for DeviceDisplayName, not even an explicit nil
### GetDeviceModel

`func (o *DeviceComplianceDeviceStatus) GetDeviceModel() string`

GetDeviceModel returns the DeviceModel field if non-nil, zero value otherwise.

### GetDeviceModelOk

`func (o *DeviceComplianceDeviceStatus) GetDeviceModelOk() (*string, bool)`

GetDeviceModelOk returns a tuple with the DeviceModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceModel

`func (o *DeviceComplianceDeviceStatus) SetDeviceModel(v string)`

SetDeviceModel sets DeviceModel field to given value.

### HasDeviceModel

`func (o *DeviceComplianceDeviceStatus) HasDeviceModel() bool`

HasDeviceModel returns a boolean if a field has been set.

### SetDeviceModelNil

`func (o *DeviceComplianceDeviceStatus) SetDeviceModelNil(b bool)`

 SetDeviceModelNil sets the value for DeviceModel to be an explicit nil

### UnsetDeviceModel
`func (o *DeviceComplianceDeviceStatus) UnsetDeviceModel()`

UnsetDeviceModel ensures that no value is present for DeviceModel, not even an explicit nil
### GetLastReportedDateTime

`func (o *DeviceComplianceDeviceStatus) GetLastReportedDateTime() time.Time`

GetLastReportedDateTime returns the LastReportedDateTime field if non-nil, zero value otherwise.

### GetLastReportedDateTimeOk

`func (o *DeviceComplianceDeviceStatus) GetLastReportedDateTimeOk() (*time.Time, bool)`

GetLastReportedDateTimeOk returns a tuple with the LastReportedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastReportedDateTime

`func (o *DeviceComplianceDeviceStatus) SetLastReportedDateTime(v time.Time)`

SetLastReportedDateTime sets LastReportedDateTime field to given value.

### HasLastReportedDateTime

`func (o *DeviceComplianceDeviceStatus) HasLastReportedDateTime() bool`

HasLastReportedDateTime returns a boolean if a field has been set.

### GetStatus

`func (o *DeviceComplianceDeviceStatus) GetStatus() AnyOfmicrosoftGraphComplianceStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DeviceComplianceDeviceStatus) GetStatusOk() (*AnyOfmicrosoftGraphComplianceStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DeviceComplianceDeviceStatus) SetStatus(v AnyOfmicrosoftGraphComplianceStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DeviceComplianceDeviceStatus) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *DeviceComplianceDeviceStatus) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *DeviceComplianceDeviceStatus) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetUserName

`func (o *DeviceComplianceDeviceStatus) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *DeviceComplianceDeviceStatus) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *DeviceComplianceDeviceStatus) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *DeviceComplianceDeviceStatus) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### SetUserNameNil

`func (o *DeviceComplianceDeviceStatus) SetUserNameNil(b bool)`

 SetUserNameNil sets the value for UserName to be an explicit nil

### UnsetUserName
`func (o *DeviceComplianceDeviceStatus) UnsetUserName()`

UnsetUserName ensures that no value is present for UserName, not even an explicit nil
### GetUserPrincipalName

`func (o *DeviceComplianceDeviceStatus) GetUserPrincipalName() string`

GetUserPrincipalName returns the UserPrincipalName field if non-nil, zero value otherwise.

### GetUserPrincipalNameOk

`func (o *DeviceComplianceDeviceStatus) GetUserPrincipalNameOk() (*string, bool)`

GetUserPrincipalNameOk returns a tuple with the UserPrincipalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPrincipalName

`func (o *DeviceComplianceDeviceStatus) SetUserPrincipalName(v string)`

SetUserPrincipalName sets UserPrincipalName field to given value.

### HasUserPrincipalName

`func (o *DeviceComplianceDeviceStatus) HasUserPrincipalName() bool`

HasUserPrincipalName returns a boolean if a field has been set.

### SetUserPrincipalNameNil

`func (o *DeviceComplianceDeviceStatus) SetUserPrincipalNameNil(b bool)`

 SetUserPrincipalNameNil sets the value for UserPrincipalName to be an explicit nil

### UnsetUserPrincipalName
`func (o *DeviceComplianceDeviceStatus) UnsetUserPrincipalName()`

UnsetUserPrincipalName ensures that no value is present for UserPrincipalName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


