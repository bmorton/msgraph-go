# MicrosoftGraphMobileThreatDefenseConnector

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**AndroidDeviceBlockedOnMissingPartnerData** | Pointer to **bool** | For Android, set whether Intune must receive data from the data sync partner prior to marking a device compliant | [optional] 
**AndroidEnabled** | Pointer to **bool** | For Android, set whether data from the data sync partner should be used during compliance evaluations | [optional] 
**IosDeviceBlockedOnMissingPartnerData** | Pointer to **bool** | For IOS, set whether Intune must receive data from the data sync partner prior to marking a device compliant | [optional] 
**IosEnabled** | Pointer to **bool** | For IOS, get or set whether data from the data sync partner should be used during compliance evaluations | [optional] 
**LastHeartbeatDateTime** | Pointer to **time.Time** | DateTime of last Heartbeat recieved from the Data Sync Partner | [optional] 
**PartnerState** | Pointer to [**AnyOfmicrosoftGraphMobileThreatPartnerTenantState**](anyOf&lt;microsoft.graph.mobileThreatPartnerTenantState&gt;.md) | Data Sync Partner state for this account. Possible values are: unavailable, available, enabled, unresponsive. | [optional] 
**PartnerUnresponsivenessThresholdInDays** | Pointer to **int32** | Get or Set days the per tenant tolerance to unresponsiveness for this partner integration | [optional] 
**PartnerUnsupportedOsVersionBlocked** | Pointer to **bool** | Get or set whether to block devices on the enabled platforms that do not meet the minimum version requirements of the Data Sync Partner | [optional] 

## Methods

### NewMicrosoftGraphMobileThreatDefenseConnector

`func NewMicrosoftGraphMobileThreatDefenseConnector() *MicrosoftGraphMobileThreatDefenseConnector`

NewMicrosoftGraphMobileThreatDefenseConnector instantiates a new MicrosoftGraphMobileThreatDefenseConnector object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphMobileThreatDefenseConnectorWithDefaults

`func NewMicrosoftGraphMobileThreatDefenseConnectorWithDefaults() *MicrosoftGraphMobileThreatDefenseConnector`

NewMicrosoftGraphMobileThreatDefenseConnectorWithDefaults instantiates a new MicrosoftGraphMobileThreatDefenseConnector object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphMobileThreatDefenseConnector) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphMobileThreatDefenseConnector) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphMobileThreatDefenseConnector) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphMobileThreatDefenseConnector) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAndroidDeviceBlockedOnMissingPartnerData

`func (o *MicrosoftGraphMobileThreatDefenseConnector) GetAndroidDeviceBlockedOnMissingPartnerData() bool`

GetAndroidDeviceBlockedOnMissingPartnerData returns the AndroidDeviceBlockedOnMissingPartnerData field if non-nil, zero value otherwise.

### GetAndroidDeviceBlockedOnMissingPartnerDataOk

`func (o *MicrosoftGraphMobileThreatDefenseConnector) GetAndroidDeviceBlockedOnMissingPartnerDataOk() (*bool, bool)`

GetAndroidDeviceBlockedOnMissingPartnerDataOk returns a tuple with the AndroidDeviceBlockedOnMissingPartnerData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAndroidDeviceBlockedOnMissingPartnerData

`func (o *MicrosoftGraphMobileThreatDefenseConnector) SetAndroidDeviceBlockedOnMissingPartnerData(v bool)`

SetAndroidDeviceBlockedOnMissingPartnerData sets AndroidDeviceBlockedOnMissingPartnerData field to given value.

### HasAndroidDeviceBlockedOnMissingPartnerData

`func (o *MicrosoftGraphMobileThreatDefenseConnector) HasAndroidDeviceBlockedOnMissingPartnerData() bool`

HasAndroidDeviceBlockedOnMissingPartnerData returns a boolean if a field has been set.

### GetAndroidEnabled

`func (o *MicrosoftGraphMobileThreatDefenseConnector) GetAndroidEnabled() bool`

GetAndroidEnabled returns the AndroidEnabled field if non-nil, zero value otherwise.

### GetAndroidEnabledOk

`func (o *MicrosoftGraphMobileThreatDefenseConnector) GetAndroidEnabledOk() (*bool, bool)`

GetAndroidEnabledOk returns a tuple with the AndroidEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAndroidEnabled

`func (o *MicrosoftGraphMobileThreatDefenseConnector) SetAndroidEnabled(v bool)`

SetAndroidEnabled sets AndroidEnabled field to given value.

### HasAndroidEnabled

`func (o *MicrosoftGraphMobileThreatDefenseConnector) HasAndroidEnabled() bool`

HasAndroidEnabled returns a boolean if a field has been set.

### GetIosDeviceBlockedOnMissingPartnerData

`func (o *MicrosoftGraphMobileThreatDefenseConnector) GetIosDeviceBlockedOnMissingPartnerData() bool`

GetIosDeviceBlockedOnMissingPartnerData returns the IosDeviceBlockedOnMissingPartnerData field if non-nil, zero value otherwise.

### GetIosDeviceBlockedOnMissingPartnerDataOk

`func (o *MicrosoftGraphMobileThreatDefenseConnector) GetIosDeviceBlockedOnMissingPartnerDataOk() (*bool, bool)`

GetIosDeviceBlockedOnMissingPartnerDataOk returns a tuple with the IosDeviceBlockedOnMissingPartnerData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIosDeviceBlockedOnMissingPartnerData

`func (o *MicrosoftGraphMobileThreatDefenseConnector) SetIosDeviceBlockedOnMissingPartnerData(v bool)`

SetIosDeviceBlockedOnMissingPartnerData sets IosDeviceBlockedOnMissingPartnerData field to given value.

### HasIosDeviceBlockedOnMissingPartnerData

`func (o *MicrosoftGraphMobileThreatDefenseConnector) HasIosDeviceBlockedOnMissingPartnerData() bool`

HasIosDeviceBlockedOnMissingPartnerData returns a boolean if a field has been set.

### GetIosEnabled

`func (o *MicrosoftGraphMobileThreatDefenseConnector) GetIosEnabled() bool`

GetIosEnabled returns the IosEnabled field if non-nil, zero value otherwise.

### GetIosEnabledOk

`func (o *MicrosoftGraphMobileThreatDefenseConnector) GetIosEnabledOk() (*bool, bool)`

GetIosEnabledOk returns a tuple with the IosEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIosEnabled

`func (o *MicrosoftGraphMobileThreatDefenseConnector) SetIosEnabled(v bool)`

SetIosEnabled sets IosEnabled field to given value.

### HasIosEnabled

`func (o *MicrosoftGraphMobileThreatDefenseConnector) HasIosEnabled() bool`

HasIosEnabled returns a boolean if a field has been set.

### GetLastHeartbeatDateTime

`func (o *MicrosoftGraphMobileThreatDefenseConnector) GetLastHeartbeatDateTime() time.Time`

GetLastHeartbeatDateTime returns the LastHeartbeatDateTime field if non-nil, zero value otherwise.

### GetLastHeartbeatDateTimeOk

`func (o *MicrosoftGraphMobileThreatDefenseConnector) GetLastHeartbeatDateTimeOk() (*time.Time, bool)`

GetLastHeartbeatDateTimeOk returns a tuple with the LastHeartbeatDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastHeartbeatDateTime

`func (o *MicrosoftGraphMobileThreatDefenseConnector) SetLastHeartbeatDateTime(v time.Time)`

SetLastHeartbeatDateTime sets LastHeartbeatDateTime field to given value.

### HasLastHeartbeatDateTime

`func (o *MicrosoftGraphMobileThreatDefenseConnector) HasLastHeartbeatDateTime() bool`

HasLastHeartbeatDateTime returns a boolean if a field has been set.

### GetPartnerState

`func (o *MicrosoftGraphMobileThreatDefenseConnector) GetPartnerState() AnyOfmicrosoftGraphMobileThreatPartnerTenantState`

GetPartnerState returns the PartnerState field if non-nil, zero value otherwise.

### GetPartnerStateOk

`func (o *MicrosoftGraphMobileThreatDefenseConnector) GetPartnerStateOk() (*AnyOfmicrosoftGraphMobileThreatPartnerTenantState, bool)`

GetPartnerStateOk returns a tuple with the PartnerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerState

`func (o *MicrosoftGraphMobileThreatDefenseConnector) SetPartnerState(v AnyOfmicrosoftGraphMobileThreatPartnerTenantState)`

SetPartnerState sets PartnerState field to given value.

### HasPartnerState

`func (o *MicrosoftGraphMobileThreatDefenseConnector) HasPartnerState() bool`

HasPartnerState returns a boolean if a field has been set.

### SetPartnerStateNil

`func (o *MicrosoftGraphMobileThreatDefenseConnector) SetPartnerStateNil(b bool)`

 SetPartnerStateNil sets the value for PartnerState to be an explicit nil

### UnsetPartnerState
`func (o *MicrosoftGraphMobileThreatDefenseConnector) UnsetPartnerState()`

UnsetPartnerState ensures that no value is present for PartnerState, not even an explicit nil
### GetPartnerUnresponsivenessThresholdInDays

`func (o *MicrosoftGraphMobileThreatDefenseConnector) GetPartnerUnresponsivenessThresholdInDays() int32`

GetPartnerUnresponsivenessThresholdInDays returns the PartnerUnresponsivenessThresholdInDays field if non-nil, zero value otherwise.

### GetPartnerUnresponsivenessThresholdInDaysOk

`func (o *MicrosoftGraphMobileThreatDefenseConnector) GetPartnerUnresponsivenessThresholdInDaysOk() (*int32, bool)`

GetPartnerUnresponsivenessThresholdInDaysOk returns a tuple with the PartnerUnresponsivenessThresholdInDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerUnresponsivenessThresholdInDays

`func (o *MicrosoftGraphMobileThreatDefenseConnector) SetPartnerUnresponsivenessThresholdInDays(v int32)`

SetPartnerUnresponsivenessThresholdInDays sets PartnerUnresponsivenessThresholdInDays field to given value.

### HasPartnerUnresponsivenessThresholdInDays

`func (o *MicrosoftGraphMobileThreatDefenseConnector) HasPartnerUnresponsivenessThresholdInDays() bool`

HasPartnerUnresponsivenessThresholdInDays returns a boolean if a field has been set.

### GetPartnerUnsupportedOsVersionBlocked

`func (o *MicrosoftGraphMobileThreatDefenseConnector) GetPartnerUnsupportedOsVersionBlocked() bool`

GetPartnerUnsupportedOsVersionBlocked returns the PartnerUnsupportedOsVersionBlocked field if non-nil, zero value otherwise.

### GetPartnerUnsupportedOsVersionBlockedOk

`func (o *MicrosoftGraphMobileThreatDefenseConnector) GetPartnerUnsupportedOsVersionBlockedOk() (*bool, bool)`

GetPartnerUnsupportedOsVersionBlockedOk returns a tuple with the PartnerUnsupportedOsVersionBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerUnsupportedOsVersionBlocked

`func (o *MicrosoftGraphMobileThreatDefenseConnector) SetPartnerUnsupportedOsVersionBlocked(v bool)`

SetPartnerUnsupportedOsVersionBlocked sets PartnerUnsupportedOsVersionBlocked field to given value.

### HasPartnerUnsupportedOsVersionBlocked

`func (o *MicrosoftGraphMobileThreatDefenseConnector) HasPartnerUnsupportedOsVersionBlocked() bool`

HasPartnerUnsupportedOsVersionBlocked returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


