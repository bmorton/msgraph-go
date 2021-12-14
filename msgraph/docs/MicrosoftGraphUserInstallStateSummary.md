# MicrosoftGraphUserInstallStateSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**FailedDeviceCount** | Pointer to **int32** | Failed Device Count. | [optional] 
**InstalledDeviceCount** | Pointer to **int32** | Installed Device Count. | [optional] 
**NotInstalledDeviceCount** | Pointer to **int32** | Not installed device count. | [optional] 
**UserName** | Pointer to **NullableString** | User name. | [optional] 
**DeviceStates** | Pointer to [**[]MicrosoftGraphDeviceInstallState**](MicrosoftGraphDeviceInstallState.md) | The install state of the eBook. | [optional] 

## Methods

### NewMicrosoftGraphUserInstallStateSummary

`func NewMicrosoftGraphUserInstallStateSummary() *MicrosoftGraphUserInstallStateSummary`

NewMicrosoftGraphUserInstallStateSummary instantiates a new MicrosoftGraphUserInstallStateSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphUserInstallStateSummaryWithDefaults

`func NewMicrosoftGraphUserInstallStateSummaryWithDefaults() *MicrosoftGraphUserInstallStateSummary`

NewMicrosoftGraphUserInstallStateSummaryWithDefaults instantiates a new MicrosoftGraphUserInstallStateSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphUserInstallStateSummary) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphUserInstallStateSummary) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphUserInstallStateSummary) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphUserInstallStateSummary) HasId() bool`

HasId returns a boolean if a field has been set.

### GetFailedDeviceCount

`func (o *MicrosoftGraphUserInstallStateSummary) GetFailedDeviceCount() int32`

GetFailedDeviceCount returns the FailedDeviceCount field if non-nil, zero value otherwise.

### GetFailedDeviceCountOk

`func (o *MicrosoftGraphUserInstallStateSummary) GetFailedDeviceCountOk() (*int32, bool)`

GetFailedDeviceCountOk returns a tuple with the FailedDeviceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedDeviceCount

`func (o *MicrosoftGraphUserInstallStateSummary) SetFailedDeviceCount(v int32)`

SetFailedDeviceCount sets FailedDeviceCount field to given value.

### HasFailedDeviceCount

`func (o *MicrosoftGraphUserInstallStateSummary) HasFailedDeviceCount() bool`

HasFailedDeviceCount returns a boolean if a field has been set.

### GetInstalledDeviceCount

`func (o *MicrosoftGraphUserInstallStateSummary) GetInstalledDeviceCount() int32`

GetInstalledDeviceCount returns the InstalledDeviceCount field if non-nil, zero value otherwise.

### GetInstalledDeviceCountOk

`func (o *MicrosoftGraphUserInstallStateSummary) GetInstalledDeviceCountOk() (*int32, bool)`

GetInstalledDeviceCountOk returns a tuple with the InstalledDeviceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstalledDeviceCount

`func (o *MicrosoftGraphUserInstallStateSummary) SetInstalledDeviceCount(v int32)`

SetInstalledDeviceCount sets InstalledDeviceCount field to given value.

### HasInstalledDeviceCount

`func (o *MicrosoftGraphUserInstallStateSummary) HasInstalledDeviceCount() bool`

HasInstalledDeviceCount returns a boolean if a field has been set.

### GetNotInstalledDeviceCount

`func (o *MicrosoftGraphUserInstallStateSummary) GetNotInstalledDeviceCount() int32`

GetNotInstalledDeviceCount returns the NotInstalledDeviceCount field if non-nil, zero value otherwise.

### GetNotInstalledDeviceCountOk

`func (o *MicrosoftGraphUserInstallStateSummary) GetNotInstalledDeviceCountOk() (*int32, bool)`

GetNotInstalledDeviceCountOk returns a tuple with the NotInstalledDeviceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotInstalledDeviceCount

`func (o *MicrosoftGraphUserInstallStateSummary) SetNotInstalledDeviceCount(v int32)`

SetNotInstalledDeviceCount sets NotInstalledDeviceCount field to given value.

### HasNotInstalledDeviceCount

`func (o *MicrosoftGraphUserInstallStateSummary) HasNotInstalledDeviceCount() bool`

HasNotInstalledDeviceCount returns a boolean if a field has been set.

### GetUserName

`func (o *MicrosoftGraphUserInstallStateSummary) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *MicrosoftGraphUserInstallStateSummary) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *MicrosoftGraphUserInstallStateSummary) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *MicrosoftGraphUserInstallStateSummary) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### SetUserNameNil

`func (o *MicrosoftGraphUserInstallStateSummary) SetUserNameNil(b bool)`

 SetUserNameNil sets the value for UserName to be an explicit nil

### UnsetUserName
`func (o *MicrosoftGraphUserInstallStateSummary) UnsetUserName()`

UnsetUserName ensures that no value is present for UserName, not even an explicit nil
### GetDeviceStates

`func (o *MicrosoftGraphUserInstallStateSummary) GetDeviceStates() []MicrosoftGraphDeviceInstallState`

GetDeviceStates returns the DeviceStates field if non-nil, zero value otherwise.

### GetDeviceStatesOk

`func (o *MicrosoftGraphUserInstallStateSummary) GetDeviceStatesOk() (*[]MicrosoftGraphDeviceInstallState, bool)`

GetDeviceStatesOk returns a tuple with the DeviceStates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceStates

`func (o *MicrosoftGraphUserInstallStateSummary) SetDeviceStates(v []MicrosoftGraphDeviceInstallState)`

SetDeviceStates sets DeviceStates field to given value.

### HasDeviceStates

`func (o *MicrosoftGraphUserInstallStateSummary) HasDeviceStates() bool`

HasDeviceStates returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


