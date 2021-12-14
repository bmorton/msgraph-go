# BitlockerRecoveryKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedDateTime** | Pointer to **time.Time** | The date and time when the key was originally backed up to Azure Active Directory. Not nullable. | [optional] 
**DeviceId** | Pointer to **NullableString** | Identifier of the device the BitLocker key is originally backed up from. Supports $filter (eq). | [optional] 
**Key** | Pointer to **string** | The BitLocker recovery key. Returned only on $select. Not nullable. | [optional] 
**VolumeType** | Pointer to [**AnyOfmicrosoftGraphVolumeType**](anyOf&lt;microsoft.graph.volumeType&gt;.md) | Indicates the type of volume the BitLocker key is associated with. The possible values are: 1 (for operatingSystemVolume), 2 (for fixedDataVolume), 3 (for removableDataVolume), and 4 (for unknownFutureValue). | [optional] 

## Methods

### NewBitlockerRecoveryKey

`func NewBitlockerRecoveryKey() *BitlockerRecoveryKey`

NewBitlockerRecoveryKey instantiates a new BitlockerRecoveryKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBitlockerRecoveryKeyWithDefaults

`func NewBitlockerRecoveryKeyWithDefaults() *BitlockerRecoveryKey`

NewBitlockerRecoveryKeyWithDefaults instantiates a new BitlockerRecoveryKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedDateTime

`func (o *BitlockerRecoveryKey) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *BitlockerRecoveryKey) GetCreatedDateTimeOk() (*time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateTime

`func (o *BitlockerRecoveryKey) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime sets CreatedDateTime field to given value.

### HasCreatedDateTime

`func (o *BitlockerRecoveryKey) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### GetDeviceId

`func (o *BitlockerRecoveryKey) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *BitlockerRecoveryKey) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *BitlockerRecoveryKey) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *BitlockerRecoveryKey) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### SetDeviceIdNil

`func (o *BitlockerRecoveryKey) SetDeviceIdNil(b bool)`

 SetDeviceIdNil sets the value for DeviceId to be an explicit nil

### UnsetDeviceId
`func (o *BitlockerRecoveryKey) UnsetDeviceId()`

UnsetDeviceId ensures that no value is present for DeviceId, not even an explicit nil
### GetKey

`func (o *BitlockerRecoveryKey) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *BitlockerRecoveryKey) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *BitlockerRecoveryKey) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *BitlockerRecoveryKey) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetVolumeType

`func (o *BitlockerRecoveryKey) GetVolumeType() AnyOfmicrosoftGraphVolumeType`

GetVolumeType returns the VolumeType field if non-nil, zero value otherwise.

### GetVolumeTypeOk

`func (o *BitlockerRecoveryKey) GetVolumeTypeOk() (*AnyOfmicrosoftGraphVolumeType, bool)`

GetVolumeTypeOk returns a tuple with the VolumeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeType

`func (o *BitlockerRecoveryKey) SetVolumeType(v AnyOfmicrosoftGraphVolumeType)`

SetVolumeType sets VolumeType field to given value.

### HasVolumeType

`func (o *BitlockerRecoveryKey) HasVolumeType() bool`

HasVolumeType returns a boolean if a field has been set.

### SetVolumeTypeNil

`func (o *BitlockerRecoveryKey) SetVolumeTypeNil(b bool)`

 SetVolumeTypeNil sets the value for VolumeType to be an explicit nil

### UnsetVolumeType
`func (o *BitlockerRecoveryKey) UnsetVolumeType()`

UnsetVolumeType ensures that no value is present for VolumeType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


