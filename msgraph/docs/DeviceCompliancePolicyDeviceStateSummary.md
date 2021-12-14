# DeviceCompliancePolicyDeviceStateSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompliantDeviceCount** | Pointer to **int32** | Number of compliant devices | [optional] 
**ConfigManagerCount** | Pointer to **int32** | Number of devices that have compliance managed by System Center Configuration Manager | [optional] 
**ConflictDeviceCount** | Pointer to **int32** | Number of conflict devices | [optional] 
**ErrorDeviceCount** | Pointer to **int32** | Number of error devices | [optional] 
**InGracePeriodCount** | Pointer to **int32** | Number of devices that are in grace period | [optional] 
**NonCompliantDeviceCount** | Pointer to **int32** | Number of NonCompliant devices | [optional] 
**NotApplicableDeviceCount** | Pointer to **int32** | Number of not applicable devices | [optional] 
**RemediatedDeviceCount** | Pointer to **int32** | Number of remediated devices | [optional] 
**UnknownDeviceCount** | Pointer to **int32** | Number of unknown devices | [optional] 

## Methods

### NewDeviceCompliancePolicyDeviceStateSummary

`func NewDeviceCompliancePolicyDeviceStateSummary() *DeviceCompliancePolicyDeviceStateSummary`

NewDeviceCompliancePolicyDeviceStateSummary instantiates a new DeviceCompliancePolicyDeviceStateSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceCompliancePolicyDeviceStateSummaryWithDefaults

`func NewDeviceCompliancePolicyDeviceStateSummaryWithDefaults() *DeviceCompliancePolicyDeviceStateSummary`

NewDeviceCompliancePolicyDeviceStateSummaryWithDefaults instantiates a new DeviceCompliancePolicyDeviceStateSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompliantDeviceCount

`func (o *DeviceCompliancePolicyDeviceStateSummary) GetCompliantDeviceCount() int32`

GetCompliantDeviceCount returns the CompliantDeviceCount field if non-nil, zero value otherwise.

### GetCompliantDeviceCountOk

`func (o *DeviceCompliancePolicyDeviceStateSummary) GetCompliantDeviceCountOk() (*int32, bool)`

GetCompliantDeviceCountOk returns a tuple with the CompliantDeviceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompliantDeviceCount

`func (o *DeviceCompliancePolicyDeviceStateSummary) SetCompliantDeviceCount(v int32)`

SetCompliantDeviceCount sets CompliantDeviceCount field to given value.

### HasCompliantDeviceCount

`func (o *DeviceCompliancePolicyDeviceStateSummary) HasCompliantDeviceCount() bool`

HasCompliantDeviceCount returns a boolean if a field has been set.

### GetConfigManagerCount

`func (o *DeviceCompliancePolicyDeviceStateSummary) GetConfigManagerCount() int32`

GetConfigManagerCount returns the ConfigManagerCount field if non-nil, zero value otherwise.

### GetConfigManagerCountOk

`func (o *DeviceCompliancePolicyDeviceStateSummary) GetConfigManagerCountOk() (*int32, bool)`

GetConfigManagerCountOk returns a tuple with the ConfigManagerCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigManagerCount

`func (o *DeviceCompliancePolicyDeviceStateSummary) SetConfigManagerCount(v int32)`

SetConfigManagerCount sets ConfigManagerCount field to given value.

### HasConfigManagerCount

`func (o *DeviceCompliancePolicyDeviceStateSummary) HasConfigManagerCount() bool`

HasConfigManagerCount returns a boolean if a field has been set.

### GetConflictDeviceCount

`func (o *DeviceCompliancePolicyDeviceStateSummary) GetConflictDeviceCount() int32`

GetConflictDeviceCount returns the ConflictDeviceCount field if non-nil, zero value otherwise.

### GetConflictDeviceCountOk

`func (o *DeviceCompliancePolicyDeviceStateSummary) GetConflictDeviceCountOk() (*int32, bool)`

GetConflictDeviceCountOk returns a tuple with the ConflictDeviceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConflictDeviceCount

`func (o *DeviceCompliancePolicyDeviceStateSummary) SetConflictDeviceCount(v int32)`

SetConflictDeviceCount sets ConflictDeviceCount field to given value.

### HasConflictDeviceCount

`func (o *DeviceCompliancePolicyDeviceStateSummary) HasConflictDeviceCount() bool`

HasConflictDeviceCount returns a boolean if a field has been set.

### GetErrorDeviceCount

`func (o *DeviceCompliancePolicyDeviceStateSummary) GetErrorDeviceCount() int32`

GetErrorDeviceCount returns the ErrorDeviceCount field if non-nil, zero value otherwise.

### GetErrorDeviceCountOk

`func (o *DeviceCompliancePolicyDeviceStateSummary) GetErrorDeviceCountOk() (*int32, bool)`

GetErrorDeviceCountOk returns a tuple with the ErrorDeviceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorDeviceCount

`func (o *DeviceCompliancePolicyDeviceStateSummary) SetErrorDeviceCount(v int32)`

SetErrorDeviceCount sets ErrorDeviceCount field to given value.

### HasErrorDeviceCount

`func (o *DeviceCompliancePolicyDeviceStateSummary) HasErrorDeviceCount() bool`

HasErrorDeviceCount returns a boolean if a field has been set.

### GetInGracePeriodCount

`func (o *DeviceCompliancePolicyDeviceStateSummary) GetInGracePeriodCount() int32`

GetInGracePeriodCount returns the InGracePeriodCount field if non-nil, zero value otherwise.

### GetInGracePeriodCountOk

`func (o *DeviceCompliancePolicyDeviceStateSummary) GetInGracePeriodCountOk() (*int32, bool)`

GetInGracePeriodCountOk returns a tuple with the InGracePeriodCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInGracePeriodCount

`func (o *DeviceCompliancePolicyDeviceStateSummary) SetInGracePeriodCount(v int32)`

SetInGracePeriodCount sets InGracePeriodCount field to given value.

### HasInGracePeriodCount

`func (o *DeviceCompliancePolicyDeviceStateSummary) HasInGracePeriodCount() bool`

HasInGracePeriodCount returns a boolean if a field has been set.

### GetNonCompliantDeviceCount

`func (o *DeviceCompliancePolicyDeviceStateSummary) GetNonCompliantDeviceCount() int32`

GetNonCompliantDeviceCount returns the NonCompliantDeviceCount field if non-nil, zero value otherwise.

### GetNonCompliantDeviceCountOk

`func (o *DeviceCompliancePolicyDeviceStateSummary) GetNonCompliantDeviceCountOk() (*int32, bool)`

GetNonCompliantDeviceCountOk returns a tuple with the NonCompliantDeviceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonCompliantDeviceCount

`func (o *DeviceCompliancePolicyDeviceStateSummary) SetNonCompliantDeviceCount(v int32)`

SetNonCompliantDeviceCount sets NonCompliantDeviceCount field to given value.

### HasNonCompliantDeviceCount

`func (o *DeviceCompliancePolicyDeviceStateSummary) HasNonCompliantDeviceCount() bool`

HasNonCompliantDeviceCount returns a boolean if a field has been set.

### GetNotApplicableDeviceCount

`func (o *DeviceCompliancePolicyDeviceStateSummary) GetNotApplicableDeviceCount() int32`

GetNotApplicableDeviceCount returns the NotApplicableDeviceCount field if non-nil, zero value otherwise.

### GetNotApplicableDeviceCountOk

`func (o *DeviceCompliancePolicyDeviceStateSummary) GetNotApplicableDeviceCountOk() (*int32, bool)`

GetNotApplicableDeviceCountOk returns a tuple with the NotApplicableDeviceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotApplicableDeviceCount

`func (o *DeviceCompliancePolicyDeviceStateSummary) SetNotApplicableDeviceCount(v int32)`

SetNotApplicableDeviceCount sets NotApplicableDeviceCount field to given value.

### HasNotApplicableDeviceCount

`func (o *DeviceCompliancePolicyDeviceStateSummary) HasNotApplicableDeviceCount() bool`

HasNotApplicableDeviceCount returns a boolean if a field has been set.

### GetRemediatedDeviceCount

`func (o *DeviceCompliancePolicyDeviceStateSummary) GetRemediatedDeviceCount() int32`

GetRemediatedDeviceCount returns the RemediatedDeviceCount field if non-nil, zero value otherwise.

### GetRemediatedDeviceCountOk

`func (o *DeviceCompliancePolicyDeviceStateSummary) GetRemediatedDeviceCountOk() (*int32, bool)`

GetRemediatedDeviceCountOk returns a tuple with the RemediatedDeviceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemediatedDeviceCount

`func (o *DeviceCompliancePolicyDeviceStateSummary) SetRemediatedDeviceCount(v int32)`

SetRemediatedDeviceCount sets RemediatedDeviceCount field to given value.

### HasRemediatedDeviceCount

`func (o *DeviceCompliancePolicyDeviceStateSummary) HasRemediatedDeviceCount() bool`

HasRemediatedDeviceCount returns a boolean if a field has been set.

### GetUnknownDeviceCount

`func (o *DeviceCompliancePolicyDeviceStateSummary) GetUnknownDeviceCount() int32`

GetUnknownDeviceCount returns the UnknownDeviceCount field if non-nil, zero value otherwise.

### GetUnknownDeviceCountOk

`func (o *DeviceCompliancePolicyDeviceStateSummary) GetUnknownDeviceCountOk() (*int32, bool)`

GetUnknownDeviceCountOk returns a tuple with the UnknownDeviceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnknownDeviceCount

`func (o *DeviceCompliancePolicyDeviceStateSummary) SetUnknownDeviceCount(v int32)`

SetUnknownDeviceCount sets UnknownDeviceCount field to given value.

### HasUnknownDeviceCount

`func (o *DeviceCompliancePolicyDeviceStateSummary) HasUnknownDeviceCount() bool`

HasUnknownDeviceCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


