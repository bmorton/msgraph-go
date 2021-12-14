# MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
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

### NewMicrosoftGraphDeviceCompliancePolicyDeviceStateSummary

`func NewMicrosoftGraphDeviceCompliancePolicyDeviceStateSummary() *MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary`

NewMicrosoftGraphDeviceCompliancePolicyDeviceStateSummary instantiates a new MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphDeviceCompliancePolicyDeviceStateSummaryWithDefaults

`func NewMicrosoftGraphDeviceCompliancePolicyDeviceStateSummaryWithDefaults() *MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary`

NewMicrosoftGraphDeviceCompliancePolicyDeviceStateSummaryWithDefaults instantiates a new MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCompliantDeviceCount

`func (o *MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary) GetCompliantDeviceCount() int32`

GetCompliantDeviceCount returns the CompliantDeviceCount field if non-nil, zero value otherwise.

### GetCompliantDeviceCountOk

`func (o *MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary) GetCompliantDeviceCountOk() (*int32, bool)`

GetCompliantDeviceCountOk returns a tuple with the CompliantDeviceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompliantDeviceCount

`func (o *MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary) SetCompliantDeviceCount(v int32)`

SetCompliantDeviceCount sets CompliantDeviceCount field to given value.

### HasCompliantDeviceCount

`func (o *MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary) HasCompliantDeviceCount() bool`

HasCompliantDeviceCount returns a boolean if a field has been set.

### GetConfigManagerCount

`func (o *MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary) GetConfigManagerCount() int32`

GetConfigManagerCount returns the ConfigManagerCount field if non-nil, zero value otherwise.

### GetConfigManagerCountOk

`func (o *MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary) GetConfigManagerCountOk() (*int32, bool)`

GetConfigManagerCountOk returns a tuple with the ConfigManagerCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigManagerCount

`func (o *MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary) SetConfigManagerCount(v int32)`

SetConfigManagerCount sets ConfigManagerCount field to given value.

### HasConfigManagerCount

`func (o *MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary) HasConfigManagerCount() bool`

HasConfigManagerCount returns a boolean if a field has been set.

### GetConflictDeviceCount

`func (o *MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary) GetConflictDeviceCount() int32`

GetConflictDeviceCount returns the ConflictDeviceCount field if non-nil, zero value otherwise.

### GetConflictDeviceCountOk

`func (o *MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary) GetConflictDeviceCountOk() (*int32, bool)`

GetConflictDeviceCountOk returns a tuple with the ConflictDeviceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConflictDeviceCount

`func (o *MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary) SetConflictDeviceCount(v int32)`

SetConflictDeviceCount sets ConflictDeviceCount field to given value.

### HasConflictDeviceCount

`func (o *MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary) HasConflictDeviceCount() bool`

HasConflictDeviceCount returns a boolean if a field has been set.

### GetErrorDeviceCount

`func (o *MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary) GetErrorDeviceCount() int32`

GetErrorDeviceCount returns the ErrorDeviceCount field if non-nil, zero value otherwise.

### GetErrorDeviceCountOk

`func (o *MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary) GetErrorDeviceCountOk() (*int32, bool)`

GetErrorDeviceCountOk returns a tuple with the ErrorDeviceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorDeviceCount

`func (o *MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary) SetErrorDeviceCount(v int32)`

SetErrorDeviceCount sets ErrorDeviceCount field to given value.

### HasErrorDeviceCount

`func (o *MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary) HasErrorDeviceCount() bool`

HasErrorDeviceCount returns a boolean if a field has been set.

### GetInGracePeriodCount

`func (o *MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary) GetInGracePeriodCount() int32`

GetInGracePeriodCount returns the InGracePeriodCount field if non-nil, zero value otherwise.

### GetInGracePeriodCountOk

`func (o *MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary) GetInGracePeriodCountOk() (*int32, bool)`

GetInGracePeriodCountOk returns a tuple with the InGracePeriodCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInGracePeriodCount

`func (o *MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary) SetInGracePeriodCount(v int32)`

SetInGracePeriodCount sets InGracePeriodCount field to given value.

### HasInGracePeriodCount

`func (o *MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary) HasInGracePeriodCount() bool`

HasInGracePeriodCount returns a boolean if a field has been set.

### GetNonCompliantDeviceCount

`func (o *MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary) GetNonCompliantDeviceCount() int32`

GetNonCompliantDeviceCount returns the NonCompliantDeviceCount field if non-nil, zero value otherwise.

### GetNonCompliantDeviceCountOk

`func (o *MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary) GetNonCompliantDeviceCountOk() (*int32, bool)`

GetNonCompliantDeviceCountOk returns a tuple with the NonCompliantDeviceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonCompliantDeviceCount

`func (o *MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary) SetNonCompliantDeviceCount(v int32)`

SetNonCompliantDeviceCount sets NonCompliantDeviceCount field to given value.

### HasNonCompliantDeviceCount

`func (o *MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary) HasNonCompliantDeviceCount() bool`

HasNonCompliantDeviceCount returns a boolean if a field has been set.

### GetNotApplicableDeviceCount

`func (o *MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary) GetNotApplicableDeviceCount() int32`

GetNotApplicableDeviceCount returns the NotApplicableDeviceCount field if non-nil, zero value otherwise.

### GetNotApplicableDeviceCountOk

`func (o *MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary) GetNotApplicableDeviceCountOk() (*int32, bool)`

GetNotApplicableDeviceCountOk returns a tuple with the NotApplicableDeviceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotApplicableDeviceCount

`func (o *MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary) SetNotApplicableDeviceCount(v int32)`

SetNotApplicableDeviceCount sets NotApplicableDeviceCount field to given value.

### HasNotApplicableDeviceCount

`func (o *MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary) HasNotApplicableDeviceCount() bool`

HasNotApplicableDeviceCount returns a boolean if a field has been set.

### GetRemediatedDeviceCount

`func (o *MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary) GetRemediatedDeviceCount() int32`

GetRemediatedDeviceCount returns the RemediatedDeviceCount field if non-nil, zero value otherwise.

### GetRemediatedDeviceCountOk

`func (o *MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary) GetRemediatedDeviceCountOk() (*int32, bool)`

GetRemediatedDeviceCountOk returns a tuple with the RemediatedDeviceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemediatedDeviceCount

`func (o *MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary) SetRemediatedDeviceCount(v int32)`

SetRemediatedDeviceCount sets RemediatedDeviceCount field to given value.

### HasRemediatedDeviceCount

`func (o *MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary) HasRemediatedDeviceCount() bool`

HasRemediatedDeviceCount returns a boolean if a field has been set.

### GetUnknownDeviceCount

`func (o *MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary) GetUnknownDeviceCount() int32`

GetUnknownDeviceCount returns the UnknownDeviceCount field if non-nil, zero value otherwise.

### GetUnknownDeviceCountOk

`func (o *MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary) GetUnknownDeviceCountOk() (*int32, bool)`

GetUnknownDeviceCountOk returns a tuple with the UnknownDeviceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnknownDeviceCount

`func (o *MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary) SetUnknownDeviceCount(v int32)`

SetUnknownDeviceCount sets UnknownDeviceCount field to given value.

### HasUnknownDeviceCount

`func (o *MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary) HasUnknownDeviceCount() bool`

HasUnknownDeviceCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


