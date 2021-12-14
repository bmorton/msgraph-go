# SettingStateDeviceSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompliantDeviceCount** | Pointer to **int32** | Device Compliant count for the setting | [optional] 
**ConflictDeviceCount** | Pointer to **int32** | Device conflict error count for the setting | [optional] 
**ErrorDeviceCount** | Pointer to **int32** | Device error count for the setting | [optional] 
**InstancePath** | Pointer to **NullableString** | Name of the InstancePath for the setting | [optional] 
**NonCompliantDeviceCount** | Pointer to **int32** | Device NonCompliant count for the setting | [optional] 
**NotApplicableDeviceCount** | Pointer to **int32** | Device Not Applicable count for the setting | [optional] 
**RemediatedDeviceCount** | Pointer to **int32** | Device Compliant count for the setting | [optional] 
**SettingName** | Pointer to **NullableString** | Name of the setting | [optional] 
**UnknownDeviceCount** | Pointer to **int32** | Device Unkown count for the setting | [optional] 

## Methods

### NewSettingStateDeviceSummary

`func NewSettingStateDeviceSummary() *SettingStateDeviceSummary`

NewSettingStateDeviceSummary instantiates a new SettingStateDeviceSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSettingStateDeviceSummaryWithDefaults

`func NewSettingStateDeviceSummaryWithDefaults() *SettingStateDeviceSummary`

NewSettingStateDeviceSummaryWithDefaults instantiates a new SettingStateDeviceSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompliantDeviceCount

`func (o *SettingStateDeviceSummary) GetCompliantDeviceCount() int32`

GetCompliantDeviceCount returns the CompliantDeviceCount field if non-nil, zero value otherwise.

### GetCompliantDeviceCountOk

`func (o *SettingStateDeviceSummary) GetCompliantDeviceCountOk() (*int32, bool)`

GetCompliantDeviceCountOk returns a tuple with the CompliantDeviceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompliantDeviceCount

`func (o *SettingStateDeviceSummary) SetCompliantDeviceCount(v int32)`

SetCompliantDeviceCount sets CompliantDeviceCount field to given value.

### HasCompliantDeviceCount

`func (o *SettingStateDeviceSummary) HasCompliantDeviceCount() bool`

HasCompliantDeviceCount returns a boolean if a field has been set.

### GetConflictDeviceCount

`func (o *SettingStateDeviceSummary) GetConflictDeviceCount() int32`

GetConflictDeviceCount returns the ConflictDeviceCount field if non-nil, zero value otherwise.

### GetConflictDeviceCountOk

`func (o *SettingStateDeviceSummary) GetConflictDeviceCountOk() (*int32, bool)`

GetConflictDeviceCountOk returns a tuple with the ConflictDeviceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConflictDeviceCount

`func (o *SettingStateDeviceSummary) SetConflictDeviceCount(v int32)`

SetConflictDeviceCount sets ConflictDeviceCount field to given value.

### HasConflictDeviceCount

`func (o *SettingStateDeviceSummary) HasConflictDeviceCount() bool`

HasConflictDeviceCount returns a boolean if a field has been set.

### GetErrorDeviceCount

`func (o *SettingStateDeviceSummary) GetErrorDeviceCount() int32`

GetErrorDeviceCount returns the ErrorDeviceCount field if non-nil, zero value otherwise.

### GetErrorDeviceCountOk

`func (o *SettingStateDeviceSummary) GetErrorDeviceCountOk() (*int32, bool)`

GetErrorDeviceCountOk returns a tuple with the ErrorDeviceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorDeviceCount

`func (o *SettingStateDeviceSummary) SetErrorDeviceCount(v int32)`

SetErrorDeviceCount sets ErrorDeviceCount field to given value.

### HasErrorDeviceCount

`func (o *SettingStateDeviceSummary) HasErrorDeviceCount() bool`

HasErrorDeviceCount returns a boolean if a field has been set.

### GetInstancePath

`func (o *SettingStateDeviceSummary) GetInstancePath() string`

GetInstancePath returns the InstancePath field if non-nil, zero value otherwise.

### GetInstancePathOk

`func (o *SettingStateDeviceSummary) GetInstancePathOk() (*string, bool)`

GetInstancePathOk returns a tuple with the InstancePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstancePath

`func (o *SettingStateDeviceSummary) SetInstancePath(v string)`

SetInstancePath sets InstancePath field to given value.

### HasInstancePath

`func (o *SettingStateDeviceSummary) HasInstancePath() bool`

HasInstancePath returns a boolean if a field has been set.

### SetInstancePathNil

`func (o *SettingStateDeviceSummary) SetInstancePathNil(b bool)`

 SetInstancePathNil sets the value for InstancePath to be an explicit nil

### UnsetInstancePath
`func (o *SettingStateDeviceSummary) UnsetInstancePath()`

UnsetInstancePath ensures that no value is present for InstancePath, not even an explicit nil
### GetNonCompliantDeviceCount

`func (o *SettingStateDeviceSummary) GetNonCompliantDeviceCount() int32`

GetNonCompliantDeviceCount returns the NonCompliantDeviceCount field if non-nil, zero value otherwise.

### GetNonCompliantDeviceCountOk

`func (o *SettingStateDeviceSummary) GetNonCompliantDeviceCountOk() (*int32, bool)`

GetNonCompliantDeviceCountOk returns a tuple with the NonCompliantDeviceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonCompliantDeviceCount

`func (o *SettingStateDeviceSummary) SetNonCompliantDeviceCount(v int32)`

SetNonCompliantDeviceCount sets NonCompliantDeviceCount field to given value.

### HasNonCompliantDeviceCount

`func (o *SettingStateDeviceSummary) HasNonCompliantDeviceCount() bool`

HasNonCompliantDeviceCount returns a boolean if a field has been set.

### GetNotApplicableDeviceCount

`func (o *SettingStateDeviceSummary) GetNotApplicableDeviceCount() int32`

GetNotApplicableDeviceCount returns the NotApplicableDeviceCount field if non-nil, zero value otherwise.

### GetNotApplicableDeviceCountOk

`func (o *SettingStateDeviceSummary) GetNotApplicableDeviceCountOk() (*int32, bool)`

GetNotApplicableDeviceCountOk returns a tuple with the NotApplicableDeviceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotApplicableDeviceCount

`func (o *SettingStateDeviceSummary) SetNotApplicableDeviceCount(v int32)`

SetNotApplicableDeviceCount sets NotApplicableDeviceCount field to given value.

### HasNotApplicableDeviceCount

`func (o *SettingStateDeviceSummary) HasNotApplicableDeviceCount() bool`

HasNotApplicableDeviceCount returns a boolean if a field has been set.

### GetRemediatedDeviceCount

`func (o *SettingStateDeviceSummary) GetRemediatedDeviceCount() int32`

GetRemediatedDeviceCount returns the RemediatedDeviceCount field if non-nil, zero value otherwise.

### GetRemediatedDeviceCountOk

`func (o *SettingStateDeviceSummary) GetRemediatedDeviceCountOk() (*int32, bool)`

GetRemediatedDeviceCountOk returns a tuple with the RemediatedDeviceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemediatedDeviceCount

`func (o *SettingStateDeviceSummary) SetRemediatedDeviceCount(v int32)`

SetRemediatedDeviceCount sets RemediatedDeviceCount field to given value.

### HasRemediatedDeviceCount

`func (o *SettingStateDeviceSummary) HasRemediatedDeviceCount() bool`

HasRemediatedDeviceCount returns a boolean if a field has been set.

### GetSettingName

`func (o *SettingStateDeviceSummary) GetSettingName() string`

GetSettingName returns the SettingName field if non-nil, zero value otherwise.

### GetSettingNameOk

`func (o *SettingStateDeviceSummary) GetSettingNameOk() (*string, bool)`

GetSettingNameOk returns a tuple with the SettingName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettingName

`func (o *SettingStateDeviceSummary) SetSettingName(v string)`

SetSettingName sets SettingName field to given value.

### HasSettingName

`func (o *SettingStateDeviceSummary) HasSettingName() bool`

HasSettingName returns a boolean if a field has been set.

### SetSettingNameNil

`func (o *SettingStateDeviceSummary) SetSettingNameNil(b bool)`

 SetSettingNameNil sets the value for SettingName to be an explicit nil

### UnsetSettingName
`func (o *SettingStateDeviceSummary) UnsetSettingName()`

UnsetSettingName ensures that no value is present for SettingName, not even an explicit nil
### GetUnknownDeviceCount

`func (o *SettingStateDeviceSummary) GetUnknownDeviceCount() int32`

GetUnknownDeviceCount returns the UnknownDeviceCount field if non-nil, zero value otherwise.

### GetUnknownDeviceCountOk

`func (o *SettingStateDeviceSummary) GetUnknownDeviceCountOk() (*int32, bool)`

GetUnknownDeviceCountOk returns a tuple with the UnknownDeviceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnknownDeviceCount

`func (o *SettingStateDeviceSummary) SetUnknownDeviceCount(v int32)`

SetUnknownDeviceCount sets UnknownDeviceCount field to given value.

### HasUnknownDeviceCount

`func (o *SettingStateDeviceSummary) HasUnknownDeviceCount() bool`

HasUnknownDeviceCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


