# MicrosoftGraphDeviceConfigurationUserStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**DevicesCount** | Pointer to **int32** | Devices count for that user. | [optional] 
**LastReportedDateTime** | Pointer to **time.Time** | Last modified date time of the policy report. | [optional] 
**Status** | Pointer to [**AnyOfmicrosoftGraphComplianceStatus**](anyOf&lt;microsoft.graph.complianceStatus&gt;.md) | Compliance status of the policy report. Possible values are: unknown, notApplicable, compliant, remediated, nonCompliant, error, conflict, notAssigned. | [optional] 
**UserDisplayName** | Pointer to **NullableString** | User name of the DevicePolicyStatus. | [optional] 
**UserPrincipalName** | Pointer to **NullableString** | UserPrincipalName. | [optional] 

## Methods

### NewMicrosoftGraphDeviceConfigurationUserStatus

`func NewMicrosoftGraphDeviceConfigurationUserStatus() *MicrosoftGraphDeviceConfigurationUserStatus`

NewMicrosoftGraphDeviceConfigurationUserStatus instantiates a new MicrosoftGraphDeviceConfigurationUserStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphDeviceConfigurationUserStatusWithDefaults

`func NewMicrosoftGraphDeviceConfigurationUserStatusWithDefaults() *MicrosoftGraphDeviceConfigurationUserStatus`

NewMicrosoftGraphDeviceConfigurationUserStatusWithDefaults instantiates a new MicrosoftGraphDeviceConfigurationUserStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphDeviceConfigurationUserStatus) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphDeviceConfigurationUserStatus) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphDeviceConfigurationUserStatus) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphDeviceConfigurationUserStatus) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDevicesCount

`func (o *MicrosoftGraphDeviceConfigurationUserStatus) GetDevicesCount() int32`

GetDevicesCount returns the DevicesCount field if non-nil, zero value otherwise.

### GetDevicesCountOk

`func (o *MicrosoftGraphDeviceConfigurationUserStatus) GetDevicesCountOk() (*int32, bool)`

GetDevicesCountOk returns a tuple with the DevicesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevicesCount

`func (o *MicrosoftGraphDeviceConfigurationUserStatus) SetDevicesCount(v int32)`

SetDevicesCount sets DevicesCount field to given value.

### HasDevicesCount

`func (o *MicrosoftGraphDeviceConfigurationUserStatus) HasDevicesCount() bool`

HasDevicesCount returns a boolean if a field has been set.

### GetLastReportedDateTime

`func (o *MicrosoftGraphDeviceConfigurationUserStatus) GetLastReportedDateTime() time.Time`

GetLastReportedDateTime returns the LastReportedDateTime field if non-nil, zero value otherwise.

### GetLastReportedDateTimeOk

`func (o *MicrosoftGraphDeviceConfigurationUserStatus) GetLastReportedDateTimeOk() (*time.Time, bool)`

GetLastReportedDateTimeOk returns a tuple with the LastReportedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastReportedDateTime

`func (o *MicrosoftGraphDeviceConfigurationUserStatus) SetLastReportedDateTime(v time.Time)`

SetLastReportedDateTime sets LastReportedDateTime field to given value.

### HasLastReportedDateTime

`func (o *MicrosoftGraphDeviceConfigurationUserStatus) HasLastReportedDateTime() bool`

HasLastReportedDateTime returns a boolean if a field has been set.

### GetStatus

`func (o *MicrosoftGraphDeviceConfigurationUserStatus) GetStatus() AnyOfmicrosoftGraphComplianceStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MicrosoftGraphDeviceConfigurationUserStatus) GetStatusOk() (*AnyOfmicrosoftGraphComplianceStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MicrosoftGraphDeviceConfigurationUserStatus) SetStatus(v AnyOfmicrosoftGraphComplianceStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MicrosoftGraphDeviceConfigurationUserStatus) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *MicrosoftGraphDeviceConfigurationUserStatus) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *MicrosoftGraphDeviceConfigurationUserStatus) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetUserDisplayName

`func (o *MicrosoftGraphDeviceConfigurationUserStatus) GetUserDisplayName() string`

GetUserDisplayName returns the UserDisplayName field if non-nil, zero value otherwise.

### GetUserDisplayNameOk

`func (o *MicrosoftGraphDeviceConfigurationUserStatus) GetUserDisplayNameOk() (*string, bool)`

GetUserDisplayNameOk returns a tuple with the UserDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDisplayName

`func (o *MicrosoftGraphDeviceConfigurationUserStatus) SetUserDisplayName(v string)`

SetUserDisplayName sets UserDisplayName field to given value.

### HasUserDisplayName

`func (o *MicrosoftGraphDeviceConfigurationUserStatus) HasUserDisplayName() bool`

HasUserDisplayName returns a boolean if a field has been set.

### SetUserDisplayNameNil

`func (o *MicrosoftGraphDeviceConfigurationUserStatus) SetUserDisplayNameNil(b bool)`

 SetUserDisplayNameNil sets the value for UserDisplayName to be an explicit nil

### UnsetUserDisplayName
`func (o *MicrosoftGraphDeviceConfigurationUserStatus) UnsetUserDisplayName()`

UnsetUserDisplayName ensures that no value is present for UserDisplayName, not even an explicit nil
### GetUserPrincipalName

`func (o *MicrosoftGraphDeviceConfigurationUserStatus) GetUserPrincipalName() string`

GetUserPrincipalName returns the UserPrincipalName field if non-nil, zero value otherwise.

### GetUserPrincipalNameOk

`func (o *MicrosoftGraphDeviceConfigurationUserStatus) GetUserPrincipalNameOk() (*string, bool)`

GetUserPrincipalNameOk returns a tuple with the UserPrincipalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPrincipalName

`func (o *MicrosoftGraphDeviceConfigurationUserStatus) SetUserPrincipalName(v string)`

SetUserPrincipalName sets UserPrincipalName field to given value.

### HasUserPrincipalName

`func (o *MicrosoftGraphDeviceConfigurationUserStatus) HasUserPrincipalName() bool`

HasUserPrincipalName returns a boolean if a field has been set.

### SetUserPrincipalNameNil

`func (o *MicrosoftGraphDeviceConfigurationUserStatus) SetUserPrincipalNameNil(b bool)`

 SetUserPrincipalNameNil sets the value for UserPrincipalName to be an explicit nil

### UnsetUserPrincipalName
`func (o *MicrosoftGraphDeviceConfigurationUserStatus) UnsetUserPrincipalName()`

UnsetUserPrincipalName ensures that no value is present for UserPrincipalName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


