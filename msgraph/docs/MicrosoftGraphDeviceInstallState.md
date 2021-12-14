# MicrosoftGraphDeviceInstallState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**DeviceId** | Pointer to **NullableString** | Device Id. | [optional] 
**DeviceName** | Pointer to **NullableString** | Device name. | [optional] 
**ErrorCode** | Pointer to **NullableString** | The error code for install failures. | [optional] 
**InstallState** | Pointer to [**AnyOfmicrosoftGraphInstallState**](anyOf&lt;microsoft.graph.installState&gt;.md) | The install state of the eBook. Possible values are: notApplicable, installed, failed, notInstalled, uninstallFailed, unknown. | [optional] 
**LastSyncDateTime** | Pointer to **time.Time** | Last sync date and time. | [optional] 
**OsDescription** | Pointer to **NullableString** | OS Description. | [optional] 
**OsVersion** | Pointer to **NullableString** | OS Version. | [optional] 
**UserName** | Pointer to **NullableString** | Device User Name. | [optional] 

## Methods

### NewMicrosoftGraphDeviceInstallState

`func NewMicrosoftGraphDeviceInstallState() *MicrosoftGraphDeviceInstallState`

NewMicrosoftGraphDeviceInstallState instantiates a new MicrosoftGraphDeviceInstallState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphDeviceInstallStateWithDefaults

`func NewMicrosoftGraphDeviceInstallStateWithDefaults() *MicrosoftGraphDeviceInstallState`

NewMicrosoftGraphDeviceInstallStateWithDefaults instantiates a new MicrosoftGraphDeviceInstallState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphDeviceInstallState) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphDeviceInstallState) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphDeviceInstallState) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphDeviceInstallState) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDeviceId

`func (o *MicrosoftGraphDeviceInstallState) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *MicrosoftGraphDeviceInstallState) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *MicrosoftGraphDeviceInstallState) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *MicrosoftGraphDeviceInstallState) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### SetDeviceIdNil

`func (o *MicrosoftGraphDeviceInstallState) SetDeviceIdNil(b bool)`

 SetDeviceIdNil sets the value for DeviceId to be an explicit nil

### UnsetDeviceId
`func (o *MicrosoftGraphDeviceInstallState) UnsetDeviceId()`

UnsetDeviceId ensures that no value is present for DeviceId, not even an explicit nil
### GetDeviceName

`func (o *MicrosoftGraphDeviceInstallState) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *MicrosoftGraphDeviceInstallState) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *MicrosoftGraphDeviceInstallState) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.

### HasDeviceName

`func (o *MicrosoftGraphDeviceInstallState) HasDeviceName() bool`

HasDeviceName returns a boolean if a field has been set.

### SetDeviceNameNil

`func (o *MicrosoftGraphDeviceInstallState) SetDeviceNameNil(b bool)`

 SetDeviceNameNil sets the value for DeviceName to be an explicit nil

### UnsetDeviceName
`func (o *MicrosoftGraphDeviceInstallState) UnsetDeviceName()`

UnsetDeviceName ensures that no value is present for DeviceName, not even an explicit nil
### GetErrorCode

`func (o *MicrosoftGraphDeviceInstallState) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *MicrosoftGraphDeviceInstallState) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *MicrosoftGraphDeviceInstallState) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *MicrosoftGraphDeviceInstallState) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### SetErrorCodeNil

`func (o *MicrosoftGraphDeviceInstallState) SetErrorCodeNil(b bool)`

 SetErrorCodeNil sets the value for ErrorCode to be an explicit nil

### UnsetErrorCode
`func (o *MicrosoftGraphDeviceInstallState) UnsetErrorCode()`

UnsetErrorCode ensures that no value is present for ErrorCode, not even an explicit nil
### GetInstallState

`func (o *MicrosoftGraphDeviceInstallState) GetInstallState() AnyOfmicrosoftGraphInstallState`

GetInstallState returns the InstallState field if non-nil, zero value otherwise.

### GetInstallStateOk

`func (o *MicrosoftGraphDeviceInstallState) GetInstallStateOk() (*AnyOfmicrosoftGraphInstallState, bool)`

GetInstallStateOk returns a tuple with the InstallState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallState

`func (o *MicrosoftGraphDeviceInstallState) SetInstallState(v AnyOfmicrosoftGraphInstallState)`

SetInstallState sets InstallState field to given value.

### HasInstallState

`func (o *MicrosoftGraphDeviceInstallState) HasInstallState() bool`

HasInstallState returns a boolean if a field has been set.

### SetInstallStateNil

`func (o *MicrosoftGraphDeviceInstallState) SetInstallStateNil(b bool)`

 SetInstallStateNil sets the value for InstallState to be an explicit nil

### UnsetInstallState
`func (o *MicrosoftGraphDeviceInstallState) UnsetInstallState()`

UnsetInstallState ensures that no value is present for InstallState, not even an explicit nil
### GetLastSyncDateTime

`func (o *MicrosoftGraphDeviceInstallState) GetLastSyncDateTime() time.Time`

GetLastSyncDateTime returns the LastSyncDateTime field if non-nil, zero value otherwise.

### GetLastSyncDateTimeOk

`func (o *MicrosoftGraphDeviceInstallState) GetLastSyncDateTimeOk() (*time.Time, bool)`

GetLastSyncDateTimeOk returns a tuple with the LastSyncDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSyncDateTime

`func (o *MicrosoftGraphDeviceInstallState) SetLastSyncDateTime(v time.Time)`

SetLastSyncDateTime sets LastSyncDateTime field to given value.

### HasLastSyncDateTime

`func (o *MicrosoftGraphDeviceInstallState) HasLastSyncDateTime() bool`

HasLastSyncDateTime returns a boolean if a field has been set.

### GetOsDescription

`func (o *MicrosoftGraphDeviceInstallState) GetOsDescription() string`

GetOsDescription returns the OsDescription field if non-nil, zero value otherwise.

### GetOsDescriptionOk

`func (o *MicrosoftGraphDeviceInstallState) GetOsDescriptionOk() (*string, bool)`

GetOsDescriptionOk returns a tuple with the OsDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsDescription

`func (o *MicrosoftGraphDeviceInstallState) SetOsDescription(v string)`

SetOsDescription sets OsDescription field to given value.

### HasOsDescription

`func (o *MicrosoftGraphDeviceInstallState) HasOsDescription() bool`

HasOsDescription returns a boolean if a field has been set.

### SetOsDescriptionNil

`func (o *MicrosoftGraphDeviceInstallState) SetOsDescriptionNil(b bool)`

 SetOsDescriptionNil sets the value for OsDescription to be an explicit nil

### UnsetOsDescription
`func (o *MicrosoftGraphDeviceInstallState) UnsetOsDescription()`

UnsetOsDescription ensures that no value is present for OsDescription, not even an explicit nil
### GetOsVersion

`func (o *MicrosoftGraphDeviceInstallState) GetOsVersion() string`

GetOsVersion returns the OsVersion field if non-nil, zero value otherwise.

### GetOsVersionOk

`func (o *MicrosoftGraphDeviceInstallState) GetOsVersionOk() (*string, bool)`

GetOsVersionOk returns a tuple with the OsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsVersion

`func (o *MicrosoftGraphDeviceInstallState) SetOsVersion(v string)`

SetOsVersion sets OsVersion field to given value.

### HasOsVersion

`func (o *MicrosoftGraphDeviceInstallState) HasOsVersion() bool`

HasOsVersion returns a boolean if a field has been set.

### SetOsVersionNil

`func (o *MicrosoftGraphDeviceInstallState) SetOsVersionNil(b bool)`

 SetOsVersionNil sets the value for OsVersion to be an explicit nil

### UnsetOsVersion
`func (o *MicrosoftGraphDeviceInstallState) UnsetOsVersion()`

UnsetOsVersion ensures that no value is present for OsVersion, not even an explicit nil
### GetUserName

`func (o *MicrosoftGraphDeviceInstallState) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *MicrosoftGraphDeviceInstallState) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *MicrosoftGraphDeviceInstallState) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *MicrosoftGraphDeviceInstallState) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### SetUserNameNil

`func (o *MicrosoftGraphDeviceInstallState) SetUserNameNil(b bool)`

 SetUserNameNil sets the value for UserName to be an explicit nil

### UnsetUserName
`func (o *MicrosoftGraphDeviceInstallState) UnsetUserName()`

UnsetUserName ensures that no value is present for UserName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


