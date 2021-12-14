# MicrosoftGraphDetectedApp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**DeviceCount** | Pointer to **int32** | The number of devices that have installed this application | [optional] 
**DisplayName** | Pointer to **NullableString** | Name of the discovered application. Read-only | [optional] 
**SizeInByte** | Pointer to **int64** | Discovered application size in bytes. Read-only | [optional] 
**Version** | Pointer to **NullableString** | Version of the discovered application. Read-only | [optional] 
**ManagedDevices** | Pointer to [**[]MicrosoftGraphManagedDevice**](MicrosoftGraphManagedDevice.md) | The devices that have the discovered application installed | [optional] 

## Methods

### NewMicrosoftGraphDetectedApp

`func NewMicrosoftGraphDetectedApp() *MicrosoftGraphDetectedApp`

NewMicrosoftGraphDetectedApp instantiates a new MicrosoftGraphDetectedApp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphDetectedAppWithDefaults

`func NewMicrosoftGraphDetectedAppWithDefaults() *MicrosoftGraphDetectedApp`

NewMicrosoftGraphDetectedAppWithDefaults instantiates a new MicrosoftGraphDetectedApp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphDetectedApp) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphDetectedApp) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphDetectedApp) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphDetectedApp) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDeviceCount

`func (o *MicrosoftGraphDetectedApp) GetDeviceCount() int32`

GetDeviceCount returns the DeviceCount field if non-nil, zero value otherwise.

### GetDeviceCountOk

`func (o *MicrosoftGraphDetectedApp) GetDeviceCountOk() (*int32, bool)`

GetDeviceCountOk returns a tuple with the DeviceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceCount

`func (o *MicrosoftGraphDetectedApp) SetDeviceCount(v int32)`

SetDeviceCount sets DeviceCount field to given value.

### HasDeviceCount

`func (o *MicrosoftGraphDetectedApp) HasDeviceCount() bool`

HasDeviceCount returns a boolean if a field has been set.

### GetDisplayName

`func (o *MicrosoftGraphDetectedApp) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphDetectedApp) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *MicrosoftGraphDetectedApp) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *MicrosoftGraphDetectedApp) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *MicrosoftGraphDetectedApp) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *MicrosoftGraphDetectedApp) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetSizeInByte

`func (o *MicrosoftGraphDetectedApp) GetSizeInByte() int64`

GetSizeInByte returns the SizeInByte field if non-nil, zero value otherwise.

### GetSizeInByteOk

`func (o *MicrosoftGraphDetectedApp) GetSizeInByteOk() (*int64, bool)`

GetSizeInByteOk returns a tuple with the SizeInByte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeInByte

`func (o *MicrosoftGraphDetectedApp) SetSizeInByte(v int64)`

SetSizeInByte sets SizeInByte field to given value.

### HasSizeInByte

`func (o *MicrosoftGraphDetectedApp) HasSizeInByte() bool`

HasSizeInByte returns a boolean if a field has been set.

### GetVersion

`func (o *MicrosoftGraphDetectedApp) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *MicrosoftGraphDetectedApp) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *MicrosoftGraphDetectedApp) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *MicrosoftGraphDetectedApp) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *MicrosoftGraphDetectedApp) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *MicrosoftGraphDetectedApp) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetManagedDevices

`func (o *MicrosoftGraphDetectedApp) GetManagedDevices() []MicrosoftGraphManagedDevice`

GetManagedDevices returns the ManagedDevices field if non-nil, zero value otherwise.

### GetManagedDevicesOk

`func (o *MicrosoftGraphDetectedApp) GetManagedDevicesOk() (*[]MicrosoftGraphManagedDevice, bool)`

GetManagedDevicesOk returns a tuple with the ManagedDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedDevices

`func (o *MicrosoftGraphDetectedApp) SetManagedDevices(v []MicrosoftGraphManagedDevice)`

SetManagedDevices sets ManagedDevices field to given value.

### HasManagedDevices

`func (o *MicrosoftGraphDetectedApp) HasManagedDevices() bool`

HasManagedDevices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


