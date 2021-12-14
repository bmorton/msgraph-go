# DeviceInstallState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceId** | Pointer to **NullableString** | Device Id. | [optional] 
**DeviceName** | Pointer to **NullableString** | Device name. | [optional] 
**ErrorCode** | Pointer to **NullableString** | The error code for install failures. | [optional] 
**InstallState** | Pointer to [**AnyOfmicrosoftGraphInstallState**](anyOf&lt;microsoft.graph.installState&gt;.md) | The install state of the eBook. Possible values are: notApplicable, installed, failed, notInstalled, uninstallFailed, unknown. | [optional] 
**LastSyncDateTime** | Pointer to **time.Time** | Last sync date and time. | [optional] 
**OsDescription** | Pointer to **NullableString** | OS Description. | [optional] 
**OsVersion** | Pointer to **NullableString** | OS Version. | [optional] 
**UserName** | Pointer to **NullableString** | Device User Name. | [optional] 

## Methods

### NewDeviceInstallState

`func NewDeviceInstallState() *DeviceInstallState`

NewDeviceInstallState instantiates a new DeviceInstallState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceInstallStateWithDefaults

`func NewDeviceInstallStateWithDefaults() *DeviceInstallState`

NewDeviceInstallStateWithDefaults instantiates a new DeviceInstallState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceId

`func (o *DeviceInstallState) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *DeviceInstallState) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *DeviceInstallState) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *DeviceInstallState) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### SetDeviceIdNil

`func (o *DeviceInstallState) SetDeviceIdNil(b bool)`

 SetDeviceIdNil sets the value for DeviceId to be an explicit nil

### UnsetDeviceId
`func (o *DeviceInstallState) UnsetDeviceId()`

UnsetDeviceId ensures that no value is present for DeviceId, not even an explicit nil
### GetDeviceName

`func (o *DeviceInstallState) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *DeviceInstallState) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *DeviceInstallState) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.

### HasDeviceName

`func (o *DeviceInstallState) HasDeviceName() bool`

HasDeviceName returns a boolean if a field has been set.

### SetDeviceNameNil

`func (o *DeviceInstallState) SetDeviceNameNil(b bool)`

 SetDeviceNameNil sets the value for DeviceName to be an explicit nil

### UnsetDeviceName
`func (o *DeviceInstallState) UnsetDeviceName()`

UnsetDeviceName ensures that no value is present for DeviceName, not even an explicit nil
### GetErrorCode

`func (o *DeviceInstallState) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *DeviceInstallState) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *DeviceInstallState) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *DeviceInstallState) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### SetErrorCodeNil

`func (o *DeviceInstallState) SetErrorCodeNil(b bool)`

 SetErrorCodeNil sets the value for ErrorCode to be an explicit nil

### UnsetErrorCode
`func (o *DeviceInstallState) UnsetErrorCode()`

UnsetErrorCode ensures that no value is present for ErrorCode, not even an explicit nil
### GetInstallState

`func (o *DeviceInstallState) GetInstallState() AnyOfmicrosoftGraphInstallState`

GetInstallState returns the InstallState field if non-nil, zero value otherwise.

### GetInstallStateOk

`func (o *DeviceInstallState) GetInstallStateOk() (*AnyOfmicrosoftGraphInstallState, bool)`

GetInstallStateOk returns a tuple with the InstallState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallState

`func (o *DeviceInstallState) SetInstallState(v AnyOfmicrosoftGraphInstallState)`

SetInstallState sets InstallState field to given value.

### HasInstallState

`func (o *DeviceInstallState) HasInstallState() bool`

HasInstallState returns a boolean if a field has been set.

### SetInstallStateNil

`func (o *DeviceInstallState) SetInstallStateNil(b bool)`

 SetInstallStateNil sets the value for InstallState to be an explicit nil

### UnsetInstallState
`func (o *DeviceInstallState) UnsetInstallState()`

UnsetInstallState ensures that no value is present for InstallState, not even an explicit nil
### GetLastSyncDateTime

`func (o *DeviceInstallState) GetLastSyncDateTime() time.Time`

GetLastSyncDateTime returns the LastSyncDateTime field if non-nil, zero value otherwise.

### GetLastSyncDateTimeOk

`func (o *DeviceInstallState) GetLastSyncDateTimeOk() (*time.Time, bool)`

GetLastSyncDateTimeOk returns a tuple with the LastSyncDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSyncDateTime

`func (o *DeviceInstallState) SetLastSyncDateTime(v time.Time)`

SetLastSyncDateTime sets LastSyncDateTime field to given value.

### HasLastSyncDateTime

`func (o *DeviceInstallState) HasLastSyncDateTime() bool`

HasLastSyncDateTime returns a boolean if a field has been set.

### GetOsDescription

`func (o *DeviceInstallState) GetOsDescription() string`

GetOsDescription returns the OsDescription field if non-nil, zero value otherwise.

### GetOsDescriptionOk

`func (o *DeviceInstallState) GetOsDescriptionOk() (*string, bool)`

GetOsDescriptionOk returns a tuple with the OsDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsDescription

`func (o *DeviceInstallState) SetOsDescription(v string)`

SetOsDescription sets OsDescription field to given value.

### HasOsDescription

`func (o *DeviceInstallState) HasOsDescription() bool`

HasOsDescription returns a boolean if a field has been set.

### SetOsDescriptionNil

`func (o *DeviceInstallState) SetOsDescriptionNil(b bool)`

 SetOsDescriptionNil sets the value for OsDescription to be an explicit nil

### UnsetOsDescription
`func (o *DeviceInstallState) UnsetOsDescription()`

UnsetOsDescription ensures that no value is present for OsDescription, not even an explicit nil
### GetOsVersion

`func (o *DeviceInstallState) GetOsVersion() string`

GetOsVersion returns the OsVersion field if non-nil, zero value otherwise.

### GetOsVersionOk

`func (o *DeviceInstallState) GetOsVersionOk() (*string, bool)`

GetOsVersionOk returns a tuple with the OsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsVersion

`func (o *DeviceInstallState) SetOsVersion(v string)`

SetOsVersion sets OsVersion field to given value.

### HasOsVersion

`func (o *DeviceInstallState) HasOsVersion() bool`

HasOsVersion returns a boolean if a field has been set.

### SetOsVersionNil

`func (o *DeviceInstallState) SetOsVersionNil(b bool)`

 SetOsVersionNil sets the value for OsVersion to be an explicit nil

### UnsetOsVersion
`func (o *DeviceInstallState) UnsetOsVersion()`

UnsetOsVersion ensures that no value is present for OsVersion, not even an explicit nil
### GetUserName

`func (o *DeviceInstallState) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *DeviceInstallState) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *DeviceInstallState) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *DeviceInstallState) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### SetUserNameNil

`func (o *DeviceInstallState) SetUserNameNil(b bool)`

 SetUserNameNil sets the value for UserName to be an explicit nil

### UnsetUserName
`func (o *DeviceInstallState) UnsetUserName()`

UnsetUserName ensures that no value is present for UserName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


