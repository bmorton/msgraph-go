# MicrosoftGraphQuota

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Deleted** | Pointer to **NullableInt64** | Total space consumed by files in the recycle bin, in bytes. Read-only. | [optional] 
**Remaining** | Pointer to **NullableInt64** | Total space remaining before reaching the quota limit, in bytes. Read-only. | [optional] 
**State** | Pointer to **NullableString** | Enumeration value that indicates the state of the storage space. Read-only. | [optional] 
**StoragePlanInformation** | Pointer to [**AnyOfmicrosoftGraphStoragePlanInformation**](anyOf&lt;microsoft.graph.storagePlanInformation&gt;.md) | Information about the drive&#39;s storage quota plans. Only in Personal OneDrive. | [optional] 
**Total** | Pointer to **NullableInt64** | Total allowed storage space, in bytes. Read-only. | [optional] 
**Used** | Pointer to **NullableInt64** | Total space used, in bytes. Read-only. | [optional] 

## Methods

### NewMicrosoftGraphQuota

`func NewMicrosoftGraphQuota() *MicrosoftGraphQuota`

NewMicrosoftGraphQuota instantiates a new MicrosoftGraphQuota object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphQuotaWithDefaults

`func NewMicrosoftGraphQuotaWithDefaults() *MicrosoftGraphQuota`

NewMicrosoftGraphQuotaWithDefaults instantiates a new MicrosoftGraphQuota object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeleted

`func (o *MicrosoftGraphQuota) GetDeleted() int64`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *MicrosoftGraphQuota) GetDeletedOk() (*int64, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *MicrosoftGraphQuota) SetDeleted(v int64)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *MicrosoftGraphQuota) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### SetDeletedNil

`func (o *MicrosoftGraphQuota) SetDeletedNil(b bool)`

 SetDeletedNil sets the value for Deleted to be an explicit nil

### UnsetDeleted
`func (o *MicrosoftGraphQuota) UnsetDeleted()`

UnsetDeleted ensures that no value is present for Deleted, not even an explicit nil
### GetRemaining

`func (o *MicrosoftGraphQuota) GetRemaining() int64`

GetRemaining returns the Remaining field if non-nil, zero value otherwise.

### GetRemainingOk

`func (o *MicrosoftGraphQuota) GetRemainingOk() (*int64, bool)`

GetRemainingOk returns a tuple with the Remaining field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemaining

`func (o *MicrosoftGraphQuota) SetRemaining(v int64)`

SetRemaining sets Remaining field to given value.

### HasRemaining

`func (o *MicrosoftGraphQuota) HasRemaining() bool`

HasRemaining returns a boolean if a field has been set.

### SetRemainingNil

`func (o *MicrosoftGraphQuota) SetRemainingNil(b bool)`

 SetRemainingNil sets the value for Remaining to be an explicit nil

### UnsetRemaining
`func (o *MicrosoftGraphQuota) UnsetRemaining()`

UnsetRemaining ensures that no value is present for Remaining, not even an explicit nil
### GetState

`func (o *MicrosoftGraphQuota) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *MicrosoftGraphQuota) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *MicrosoftGraphQuota) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *MicrosoftGraphQuota) HasState() bool`

HasState returns a boolean if a field has been set.

### SetStateNil

`func (o *MicrosoftGraphQuota) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *MicrosoftGraphQuota) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil
### GetStoragePlanInformation

`func (o *MicrosoftGraphQuota) GetStoragePlanInformation() AnyOfmicrosoftGraphStoragePlanInformation`

GetStoragePlanInformation returns the StoragePlanInformation field if non-nil, zero value otherwise.

### GetStoragePlanInformationOk

`func (o *MicrosoftGraphQuota) GetStoragePlanInformationOk() (*AnyOfmicrosoftGraphStoragePlanInformation, bool)`

GetStoragePlanInformationOk returns a tuple with the StoragePlanInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoragePlanInformation

`func (o *MicrosoftGraphQuota) SetStoragePlanInformation(v AnyOfmicrosoftGraphStoragePlanInformation)`

SetStoragePlanInformation sets StoragePlanInformation field to given value.

### HasStoragePlanInformation

`func (o *MicrosoftGraphQuota) HasStoragePlanInformation() bool`

HasStoragePlanInformation returns a boolean if a field has been set.

### SetStoragePlanInformationNil

`func (o *MicrosoftGraphQuota) SetStoragePlanInformationNil(b bool)`

 SetStoragePlanInformationNil sets the value for StoragePlanInformation to be an explicit nil

### UnsetStoragePlanInformation
`func (o *MicrosoftGraphQuota) UnsetStoragePlanInformation()`

UnsetStoragePlanInformation ensures that no value is present for StoragePlanInformation, not even an explicit nil
### GetTotal

`func (o *MicrosoftGraphQuota) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *MicrosoftGraphQuota) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *MicrosoftGraphQuota) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *MicrosoftGraphQuota) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### SetTotalNil

`func (o *MicrosoftGraphQuota) SetTotalNil(b bool)`

 SetTotalNil sets the value for Total to be an explicit nil

### UnsetTotal
`func (o *MicrosoftGraphQuota) UnsetTotal()`

UnsetTotal ensures that no value is present for Total, not even an explicit nil
### GetUsed

`func (o *MicrosoftGraphQuota) GetUsed() int64`

GetUsed returns the Used field if non-nil, zero value otherwise.

### GetUsedOk

`func (o *MicrosoftGraphQuota) GetUsedOk() (*int64, bool)`

GetUsedOk returns a tuple with the Used field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsed

`func (o *MicrosoftGraphQuota) SetUsed(v int64)`

SetUsed sets Used field to given value.

### HasUsed

`func (o *MicrosoftGraphQuota) HasUsed() bool`

HasUsed returns a boolean if a field has been set.

### SetUsedNil

`func (o *MicrosoftGraphQuota) SetUsedNil(b bool)`

 SetUsedNil sets the value for Used to be an explicit nil

### UnsetUsed
`func (o *MicrosoftGraphQuota) UnsetUsed()`

UnsetUsed ensures that no value is present for Used, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


