# ManagedDeviceOverview

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceExchangeAccessStateSummary** | Pointer to [**AnyOfmicrosoftGraphDeviceExchangeAccessStateSummary**](anyOf&lt;microsoft.graph.deviceExchangeAccessStateSummary&gt;.md) | Distribution of Exchange Access State in Intune | [optional] 
**DeviceOperatingSystemSummary** | Pointer to [**AnyOfmicrosoftGraphDeviceOperatingSystemSummary**](anyOf&lt;microsoft.graph.deviceOperatingSystemSummary&gt;.md) | Device operating system summary. | [optional] 
**DualEnrolledDeviceCount** | Pointer to **int32** | The number of devices enrolled in both MDM and EAS | [optional] 
**EnrolledDeviceCount** | Pointer to **int32** | Total enrolled device count. Does not include PC devices managed via Intune PC Agent | [optional] 
**MdmEnrolledCount** | Pointer to **int32** | The number of devices enrolled in MDM | [optional] 

## Methods

### NewManagedDeviceOverview

`func NewManagedDeviceOverview() *ManagedDeviceOverview`

NewManagedDeviceOverview instantiates a new ManagedDeviceOverview object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagedDeviceOverviewWithDefaults

`func NewManagedDeviceOverviewWithDefaults() *ManagedDeviceOverview`

NewManagedDeviceOverviewWithDefaults instantiates a new ManagedDeviceOverview object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceExchangeAccessStateSummary

`func (o *ManagedDeviceOverview) GetDeviceExchangeAccessStateSummary() AnyOfmicrosoftGraphDeviceExchangeAccessStateSummary`

GetDeviceExchangeAccessStateSummary returns the DeviceExchangeAccessStateSummary field if non-nil, zero value otherwise.

### GetDeviceExchangeAccessStateSummaryOk

`func (o *ManagedDeviceOverview) GetDeviceExchangeAccessStateSummaryOk() (*AnyOfmicrosoftGraphDeviceExchangeAccessStateSummary, bool)`

GetDeviceExchangeAccessStateSummaryOk returns a tuple with the DeviceExchangeAccessStateSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceExchangeAccessStateSummary

`func (o *ManagedDeviceOverview) SetDeviceExchangeAccessStateSummary(v AnyOfmicrosoftGraphDeviceExchangeAccessStateSummary)`

SetDeviceExchangeAccessStateSummary sets DeviceExchangeAccessStateSummary field to given value.

### HasDeviceExchangeAccessStateSummary

`func (o *ManagedDeviceOverview) HasDeviceExchangeAccessStateSummary() bool`

HasDeviceExchangeAccessStateSummary returns a boolean if a field has been set.

### SetDeviceExchangeAccessStateSummaryNil

`func (o *ManagedDeviceOverview) SetDeviceExchangeAccessStateSummaryNil(b bool)`

 SetDeviceExchangeAccessStateSummaryNil sets the value for DeviceExchangeAccessStateSummary to be an explicit nil

### UnsetDeviceExchangeAccessStateSummary
`func (o *ManagedDeviceOverview) UnsetDeviceExchangeAccessStateSummary()`

UnsetDeviceExchangeAccessStateSummary ensures that no value is present for DeviceExchangeAccessStateSummary, not even an explicit nil
### GetDeviceOperatingSystemSummary

`func (o *ManagedDeviceOverview) GetDeviceOperatingSystemSummary() AnyOfmicrosoftGraphDeviceOperatingSystemSummary`

GetDeviceOperatingSystemSummary returns the DeviceOperatingSystemSummary field if non-nil, zero value otherwise.

### GetDeviceOperatingSystemSummaryOk

`func (o *ManagedDeviceOverview) GetDeviceOperatingSystemSummaryOk() (*AnyOfmicrosoftGraphDeviceOperatingSystemSummary, bool)`

GetDeviceOperatingSystemSummaryOk returns a tuple with the DeviceOperatingSystemSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceOperatingSystemSummary

`func (o *ManagedDeviceOverview) SetDeviceOperatingSystemSummary(v AnyOfmicrosoftGraphDeviceOperatingSystemSummary)`

SetDeviceOperatingSystemSummary sets DeviceOperatingSystemSummary field to given value.

### HasDeviceOperatingSystemSummary

`func (o *ManagedDeviceOverview) HasDeviceOperatingSystemSummary() bool`

HasDeviceOperatingSystemSummary returns a boolean if a field has been set.

### SetDeviceOperatingSystemSummaryNil

`func (o *ManagedDeviceOverview) SetDeviceOperatingSystemSummaryNil(b bool)`

 SetDeviceOperatingSystemSummaryNil sets the value for DeviceOperatingSystemSummary to be an explicit nil

### UnsetDeviceOperatingSystemSummary
`func (o *ManagedDeviceOverview) UnsetDeviceOperatingSystemSummary()`

UnsetDeviceOperatingSystemSummary ensures that no value is present for DeviceOperatingSystemSummary, not even an explicit nil
### GetDualEnrolledDeviceCount

`func (o *ManagedDeviceOverview) GetDualEnrolledDeviceCount() int32`

GetDualEnrolledDeviceCount returns the DualEnrolledDeviceCount field if non-nil, zero value otherwise.

### GetDualEnrolledDeviceCountOk

`func (o *ManagedDeviceOverview) GetDualEnrolledDeviceCountOk() (*int32, bool)`

GetDualEnrolledDeviceCountOk returns a tuple with the DualEnrolledDeviceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDualEnrolledDeviceCount

`func (o *ManagedDeviceOverview) SetDualEnrolledDeviceCount(v int32)`

SetDualEnrolledDeviceCount sets DualEnrolledDeviceCount field to given value.

### HasDualEnrolledDeviceCount

`func (o *ManagedDeviceOverview) HasDualEnrolledDeviceCount() bool`

HasDualEnrolledDeviceCount returns a boolean if a field has been set.

### GetEnrolledDeviceCount

`func (o *ManagedDeviceOverview) GetEnrolledDeviceCount() int32`

GetEnrolledDeviceCount returns the EnrolledDeviceCount field if non-nil, zero value otherwise.

### GetEnrolledDeviceCountOk

`func (o *ManagedDeviceOverview) GetEnrolledDeviceCountOk() (*int32, bool)`

GetEnrolledDeviceCountOk returns a tuple with the EnrolledDeviceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrolledDeviceCount

`func (o *ManagedDeviceOverview) SetEnrolledDeviceCount(v int32)`

SetEnrolledDeviceCount sets EnrolledDeviceCount field to given value.

### HasEnrolledDeviceCount

`func (o *ManagedDeviceOverview) HasEnrolledDeviceCount() bool`

HasEnrolledDeviceCount returns a boolean if a field has been set.

### GetMdmEnrolledCount

`func (o *ManagedDeviceOverview) GetMdmEnrolledCount() int32`

GetMdmEnrolledCount returns the MdmEnrolledCount field if non-nil, zero value otherwise.

### GetMdmEnrolledCountOk

`func (o *ManagedDeviceOverview) GetMdmEnrolledCountOk() (*int32, bool)`

GetMdmEnrolledCountOk returns a tuple with the MdmEnrolledCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMdmEnrolledCount

`func (o *ManagedDeviceOverview) SetMdmEnrolledCount(v int32)`

SetMdmEnrolledCount sets MdmEnrolledCount field to given value.

### HasMdmEnrolledCount

`func (o *ManagedDeviceOverview) HasMdmEnrolledCount() bool`

HasMdmEnrolledCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


