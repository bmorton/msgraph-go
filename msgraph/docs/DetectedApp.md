# DetectedApp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceCount** | Pointer to **int32** | The number of devices that have installed this application | [optional] 
**DisplayName** | Pointer to **NullableString** | Name of the discovered application. Read-only | [optional] 
**SizeInByte** | Pointer to **int64** | Discovered application size in bytes. Read-only | [optional] 
**Version** | Pointer to **NullableString** | Version of the discovered application. Read-only | [optional] 
**ManagedDevices** | Pointer to [**[]MicrosoftGraphManagedDevice**](MicrosoftGraphManagedDevice.md) | The devices that have the discovered application installed | [optional] 

## Methods

### NewDetectedApp

`func NewDetectedApp() *DetectedApp`

NewDetectedApp instantiates a new DetectedApp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDetectedAppWithDefaults

`func NewDetectedAppWithDefaults() *DetectedApp`

NewDetectedAppWithDefaults instantiates a new DetectedApp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceCount

`func (o *DetectedApp) GetDeviceCount() int32`

GetDeviceCount returns the DeviceCount field if non-nil, zero value otherwise.

### GetDeviceCountOk

`func (o *DetectedApp) GetDeviceCountOk() (*int32, bool)`

GetDeviceCountOk returns a tuple with the DeviceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceCount

`func (o *DetectedApp) SetDeviceCount(v int32)`

SetDeviceCount sets DeviceCount field to given value.

### HasDeviceCount

`func (o *DetectedApp) HasDeviceCount() bool`

HasDeviceCount returns a boolean if a field has been set.

### GetDisplayName

`func (o *DetectedApp) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *DetectedApp) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *DetectedApp) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *DetectedApp) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *DetectedApp) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *DetectedApp) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetSizeInByte

`func (o *DetectedApp) GetSizeInByte() int64`

GetSizeInByte returns the SizeInByte field if non-nil, zero value otherwise.

### GetSizeInByteOk

`func (o *DetectedApp) GetSizeInByteOk() (*int64, bool)`

GetSizeInByteOk returns a tuple with the SizeInByte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeInByte

`func (o *DetectedApp) SetSizeInByte(v int64)`

SetSizeInByte sets SizeInByte field to given value.

### HasSizeInByte

`func (o *DetectedApp) HasSizeInByte() bool`

HasSizeInByte returns a boolean if a field has been set.

### GetVersion

`func (o *DetectedApp) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *DetectedApp) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *DetectedApp) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *DetectedApp) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *DetectedApp) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *DetectedApp) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetManagedDevices

`func (o *DetectedApp) GetManagedDevices() []MicrosoftGraphManagedDevice`

GetManagedDevices returns the ManagedDevices field if non-nil, zero value otherwise.

### GetManagedDevicesOk

`func (o *DetectedApp) GetManagedDevicesOk() (*[]MicrosoftGraphManagedDevice, bool)`

GetManagedDevicesOk returns a tuple with the ManagedDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedDevices

`func (o *DetectedApp) SetManagedDevices(v []MicrosoftGraphManagedDevice)`

SetManagedDevices sets ManagedDevices field to given value.

### HasManagedDevices

`func (o *DetectedApp) HasManagedDevices() bool`

HasManagedDevices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


