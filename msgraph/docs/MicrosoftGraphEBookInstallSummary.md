# MicrosoftGraphEBookInstallSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**FailedDeviceCount** | Pointer to **int32** | Number of Devices that have failed to install this book. | [optional] 
**FailedUserCount** | Pointer to **int32** | Number of Users that have 1 or more device that failed to install this book. | [optional] 
**InstalledDeviceCount** | Pointer to **int32** | Number of Devices that have successfully installed this book. | [optional] 
**InstalledUserCount** | Pointer to **int32** | Number of Users whose devices have all succeeded to install this book. | [optional] 
**NotInstalledDeviceCount** | Pointer to **int32** | Number of Devices that does not have this book installed. | [optional] 
**NotInstalledUserCount** | Pointer to **int32** | Number of Users that did not install this book. | [optional] 

## Methods

### NewMicrosoftGraphEBookInstallSummary

`func NewMicrosoftGraphEBookInstallSummary() *MicrosoftGraphEBookInstallSummary`

NewMicrosoftGraphEBookInstallSummary instantiates a new MicrosoftGraphEBookInstallSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphEBookInstallSummaryWithDefaults

`func NewMicrosoftGraphEBookInstallSummaryWithDefaults() *MicrosoftGraphEBookInstallSummary`

NewMicrosoftGraphEBookInstallSummaryWithDefaults instantiates a new MicrosoftGraphEBookInstallSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphEBookInstallSummary) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphEBookInstallSummary) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphEBookInstallSummary) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphEBookInstallSummary) HasId() bool`

HasId returns a boolean if a field has been set.

### GetFailedDeviceCount

`func (o *MicrosoftGraphEBookInstallSummary) GetFailedDeviceCount() int32`

GetFailedDeviceCount returns the FailedDeviceCount field if non-nil, zero value otherwise.

### GetFailedDeviceCountOk

`func (o *MicrosoftGraphEBookInstallSummary) GetFailedDeviceCountOk() (*int32, bool)`

GetFailedDeviceCountOk returns a tuple with the FailedDeviceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedDeviceCount

`func (o *MicrosoftGraphEBookInstallSummary) SetFailedDeviceCount(v int32)`

SetFailedDeviceCount sets FailedDeviceCount field to given value.

### HasFailedDeviceCount

`func (o *MicrosoftGraphEBookInstallSummary) HasFailedDeviceCount() bool`

HasFailedDeviceCount returns a boolean if a field has been set.

### GetFailedUserCount

`func (o *MicrosoftGraphEBookInstallSummary) GetFailedUserCount() int32`

GetFailedUserCount returns the FailedUserCount field if non-nil, zero value otherwise.

### GetFailedUserCountOk

`func (o *MicrosoftGraphEBookInstallSummary) GetFailedUserCountOk() (*int32, bool)`

GetFailedUserCountOk returns a tuple with the FailedUserCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedUserCount

`func (o *MicrosoftGraphEBookInstallSummary) SetFailedUserCount(v int32)`

SetFailedUserCount sets FailedUserCount field to given value.

### HasFailedUserCount

`func (o *MicrosoftGraphEBookInstallSummary) HasFailedUserCount() bool`

HasFailedUserCount returns a boolean if a field has been set.

### GetInstalledDeviceCount

`func (o *MicrosoftGraphEBookInstallSummary) GetInstalledDeviceCount() int32`

GetInstalledDeviceCount returns the InstalledDeviceCount field if non-nil, zero value otherwise.

### GetInstalledDeviceCountOk

`func (o *MicrosoftGraphEBookInstallSummary) GetInstalledDeviceCountOk() (*int32, bool)`

GetInstalledDeviceCountOk returns a tuple with the InstalledDeviceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstalledDeviceCount

`func (o *MicrosoftGraphEBookInstallSummary) SetInstalledDeviceCount(v int32)`

SetInstalledDeviceCount sets InstalledDeviceCount field to given value.

### HasInstalledDeviceCount

`func (o *MicrosoftGraphEBookInstallSummary) HasInstalledDeviceCount() bool`

HasInstalledDeviceCount returns a boolean if a field has been set.

### GetInstalledUserCount

`func (o *MicrosoftGraphEBookInstallSummary) GetInstalledUserCount() int32`

GetInstalledUserCount returns the InstalledUserCount field if non-nil, zero value otherwise.

### GetInstalledUserCountOk

`func (o *MicrosoftGraphEBookInstallSummary) GetInstalledUserCountOk() (*int32, bool)`

GetInstalledUserCountOk returns a tuple with the InstalledUserCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstalledUserCount

`func (o *MicrosoftGraphEBookInstallSummary) SetInstalledUserCount(v int32)`

SetInstalledUserCount sets InstalledUserCount field to given value.

### HasInstalledUserCount

`func (o *MicrosoftGraphEBookInstallSummary) HasInstalledUserCount() bool`

HasInstalledUserCount returns a boolean if a field has been set.

### GetNotInstalledDeviceCount

`func (o *MicrosoftGraphEBookInstallSummary) GetNotInstalledDeviceCount() int32`

GetNotInstalledDeviceCount returns the NotInstalledDeviceCount field if non-nil, zero value otherwise.

### GetNotInstalledDeviceCountOk

`func (o *MicrosoftGraphEBookInstallSummary) GetNotInstalledDeviceCountOk() (*int32, bool)`

GetNotInstalledDeviceCountOk returns a tuple with the NotInstalledDeviceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotInstalledDeviceCount

`func (o *MicrosoftGraphEBookInstallSummary) SetNotInstalledDeviceCount(v int32)`

SetNotInstalledDeviceCount sets NotInstalledDeviceCount field to given value.

### HasNotInstalledDeviceCount

`func (o *MicrosoftGraphEBookInstallSummary) HasNotInstalledDeviceCount() bool`

HasNotInstalledDeviceCount returns a boolean if a field has been set.

### GetNotInstalledUserCount

`func (o *MicrosoftGraphEBookInstallSummary) GetNotInstalledUserCount() int32`

GetNotInstalledUserCount returns the NotInstalledUserCount field if non-nil, zero value otherwise.

### GetNotInstalledUserCountOk

`func (o *MicrosoftGraphEBookInstallSummary) GetNotInstalledUserCountOk() (*int32, bool)`

GetNotInstalledUserCountOk returns a tuple with the NotInstalledUserCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotInstalledUserCount

`func (o *MicrosoftGraphEBookInstallSummary) SetNotInstalledUserCount(v int32)`

SetNotInstalledUserCount sets NotInstalledUserCount field to given value.

### HasNotInstalledUserCount

`func (o *MicrosoftGraphEBookInstallSummary) HasNotInstalledUserCount() bool`

HasNotInstalledUserCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


