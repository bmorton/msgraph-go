# DeviceCompliancePolicySettingStateSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompliantDeviceCount** | Pointer to **int32** | Number of compliant devices | [optional] 
**ConflictDeviceCount** | Pointer to **int32** | Number of conflict devices | [optional] 
**ErrorDeviceCount** | Pointer to **int32** | Number of error devices | [optional] 
**NonCompliantDeviceCount** | Pointer to **int32** | Number of NonCompliant devices | [optional] 
**NotApplicableDeviceCount** | Pointer to **int32** | Number of not applicable devices | [optional] 
**PlatformType** | Pointer to [**AnyOfmicrosoftGraphPolicyPlatformType**](anyOf&lt;microsoft.graph.policyPlatformType&gt;.md) | Setting platform. Possible values are: android, iOS, macOS, windowsPhone81, windows81AndLater, windows10AndLater, androidWorkProfile, all. | [optional] 
**RemediatedDeviceCount** | Pointer to **int32** | Number of remediated devices | [optional] 
**Setting** | Pointer to **NullableString** | The setting class name and property name. | [optional] 
**SettingName** | Pointer to **NullableString** | Name of the setting. | [optional] 
**UnknownDeviceCount** | Pointer to **int32** | Number of unknown devices | [optional] 
**DeviceComplianceSettingStates** | Pointer to [**[]MicrosoftGraphDeviceComplianceSettingState**](MicrosoftGraphDeviceComplianceSettingState.md) | Not yet documented | [optional] 

## Methods

### NewDeviceCompliancePolicySettingStateSummary

`func NewDeviceCompliancePolicySettingStateSummary() *DeviceCompliancePolicySettingStateSummary`

NewDeviceCompliancePolicySettingStateSummary instantiates a new DeviceCompliancePolicySettingStateSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceCompliancePolicySettingStateSummaryWithDefaults

`func NewDeviceCompliancePolicySettingStateSummaryWithDefaults() *DeviceCompliancePolicySettingStateSummary`

NewDeviceCompliancePolicySettingStateSummaryWithDefaults instantiates a new DeviceCompliancePolicySettingStateSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompliantDeviceCount

`func (o *DeviceCompliancePolicySettingStateSummary) GetCompliantDeviceCount() int32`

GetCompliantDeviceCount returns the CompliantDeviceCount field if non-nil, zero value otherwise.

### GetCompliantDeviceCountOk

`func (o *DeviceCompliancePolicySettingStateSummary) GetCompliantDeviceCountOk() (*int32, bool)`

GetCompliantDeviceCountOk returns a tuple with the CompliantDeviceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompliantDeviceCount

`func (o *DeviceCompliancePolicySettingStateSummary) SetCompliantDeviceCount(v int32)`

SetCompliantDeviceCount sets CompliantDeviceCount field to given value.

### HasCompliantDeviceCount

`func (o *DeviceCompliancePolicySettingStateSummary) HasCompliantDeviceCount() bool`

HasCompliantDeviceCount returns a boolean if a field has been set.

### GetConflictDeviceCount

`func (o *DeviceCompliancePolicySettingStateSummary) GetConflictDeviceCount() int32`

GetConflictDeviceCount returns the ConflictDeviceCount field if non-nil, zero value otherwise.

### GetConflictDeviceCountOk

`func (o *DeviceCompliancePolicySettingStateSummary) GetConflictDeviceCountOk() (*int32, bool)`

GetConflictDeviceCountOk returns a tuple with the ConflictDeviceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConflictDeviceCount

`func (o *DeviceCompliancePolicySettingStateSummary) SetConflictDeviceCount(v int32)`

SetConflictDeviceCount sets ConflictDeviceCount field to given value.

### HasConflictDeviceCount

`func (o *DeviceCompliancePolicySettingStateSummary) HasConflictDeviceCount() bool`

HasConflictDeviceCount returns a boolean if a field has been set.

### GetErrorDeviceCount

`func (o *DeviceCompliancePolicySettingStateSummary) GetErrorDeviceCount() int32`

GetErrorDeviceCount returns the ErrorDeviceCount field if non-nil, zero value otherwise.

### GetErrorDeviceCountOk

`func (o *DeviceCompliancePolicySettingStateSummary) GetErrorDeviceCountOk() (*int32, bool)`

GetErrorDeviceCountOk returns a tuple with the ErrorDeviceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorDeviceCount

`func (o *DeviceCompliancePolicySettingStateSummary) SetErrorDeviceCount(v int32)`

SetErrorDeviceCount sets ErrorDeviceCount field to given value.

### HasErrorDeviceCount

`func (o *DeviceCompliancePolicySettingStateSummary) HasErrorDeviceCount() bool`

HasErrorDeviceCount returns a boolean if a field has been set.

### GetNonCompliantDeviceCount

`func (o *DeviceCompliancePolicySettingStateSummary) GetNonCompliantDeviceCount() int32`

GetNonCompliantDeviceCount returns the NonCompliantDeviceCount field if non-nil, zero value otherwise.

### GetNonCompliantDeviceCountOk

`func (o *DeviceCompliancePolicySettingStateSummary) GetNonCompliantDeviceCountOk() (*int32, bool)`

GetNonCompliantDeviceCountOk returns a tuple with the NonCompliantDeviceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonCompliantDeviceCount

`func (o *DeviceCompliancePolicySettingStateSummary) SetNonCompliantDeviceCount(v int32)`

SetNonCompliantDeviceCount sets NonCompliantDeviceCount field to given value.

### HasNonCompliantDeviceCount

`func (o *DeviceCompliancePolicySettingStateSummary) HasNonCompliantDeviceCount() bool`

HasNonCompliantDeviceCount returns a boolean if a field has been set.

### GetNotApplicableDeviceCount

`func (o *DeviceCompliancePolicySettingStateSummary) GetNotApplicableDeviceCount() int32`

GetNotApplicableDeviceCount returns the NotApplicableDeviceCount field if non-nil, zero value otherwise.

### GetNotApplicableDeviceCountOk

`func (o *DeviceCompliancePolicySettingStateSummary) GetNotApplicableDeviceCountOk() (*int32, bool)`

GetNotApplicableDeviceCountOk returns a tuple with the NotApplicableDeviceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotApplicableDeviceCount

`func (o *DeviceCompliancePolicySettingStateSummary) SetNotApplicableDeviceCount(v int32)`

SetNotApplicableDeviceCount sets NotApplicableDeviceCount field to given value.

### HasNotApplicableDeviceCount

`func (o *DeviceCompliancePolicySettingStateSummary) HasNotApplicableDeviceCount() bool`

HasNotApplicableDeviceCount returns a boolean if a field has been set.

### GetPlatformType

`func (o *DeviceCompliancePolicySettingStateSummary) GetPlatformType() AnyOfmicrosoftGraphPolicyPlatformType`

GetPlatformType returns the PlatformType field if non-nil, zero value otherwise.

### GetPlatformTypeOk

`func (o *DeviceCompliancePolicySettingStateSummary) GetPlatformTypeOk() (*AnyOfmicrosoftGraphPolicyPlatformType, bool)`

GetPlatformTypeOk returns a tuple with the PlatformType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformType

`func (o *DeviceCompliancePolicySettingStateSummary) SetPlatformType(v AnyOfmicrosoftGraphPolicyPlatformType)`

SetPlatformType sets PlatformType field to given value.

### HasPlatformType

`func (o *DeviceCompliancePolicySettingStateSummary) HasPlatformType() bool`

HasPlatformType returns a boolean if a field has been set.

### SetPlatformTypeNil

`func (o *DeviceCompliancePolicySettingStateSummary) SetPlatformTypeNil(b bool)`

 SetPlatformTypeNil sets the value for PlatformType to be an explicit nil

### UnsetPlatformType
`func (o *DeviceCompliancePolicySettingStateSummary) UnsetPlatformType()`

UnsetPlatformType ensures that no value is present for PlatformType, not even an explicit nil
### GetRemediatedDeviceCount

`func (o *DeviceCompliancePolicySettingStateSummary) GetRemediatedDeviceCount() int32`

GetRemediatedDeviceCount returns the RemediatedDeviceCount field if non-nil, zero value otherwise.

### GetRemediatedDeviceCountOk

`func (o *DeviceCompliancePolicySettingStateSummary) GetRemediatedDeviceCountOk() (*int32, bool)`

GetRemediatedDeviceCountOk returns a tuple with the RemediatedDeviceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemediatedDeviceCount

`func (o *DeviceCompliancePolicySettingStateSummary) SetRemediatedDeviceCount(v int32)`

SetRemediatedDeviceCount sets RemediatedDeviceCount field to given value.

### HasRemediatedDeviceCount

`func (o *DeviceCompliancePolicySettingStateSummary) HasRemediatedDeviceCount() bool`

HasRemediatedDeviceCount returns a boolean if a field has been set.

### GetSetting

`func (o *DeviceCompliancePolicySettingStateSummary) GetSetting() string`

GetSetting returns the Setting field if non-nil, zero value otherwise.

### GetSettingOk

`func (o *DeviceCompliancePolicySettingStateSummary) GetSettingOk() (*string, bool)`

GetSettingOk returns a tuple with the Setting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetting

`func (o *DeviceCompliancePolicySettingStateSummary) SetSetting(v string)`

SetSetting sets Setting field to given value.

### HasSetting

`func (o *DeviceCompliancePolicySettingStateSummary) HasSetting() bool`

HasSetting returns a boolean if a field has been set.

### SetSettingNil

`func (o *DeviceCompliancePolicySettingStateSummary) SetSettingNil(b bool)`

 SetSettingNil sets the value for Setting to be an explicit nil

### UnsetSetting
`func (o *DeviceCompliancePolicySettingStateSummary) UnsetSetting()`

UnsetSetting ensures that no value is present for Setting, not even an explicit nil
### GetSettingName

`func (o *DeviceCompliancePolicySettingStateSummary) GetSettingName() string`

GetSettingName returns the SettingName field if non-nil, zero value otherwise.

### GetSettingNameOk

`func (o *DeviceCompliancePolicySettingStateSummary) GetSettingNameOk() (*string, bool)`

GetSettingNameOk returns a tuple with the SettingName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettingName

`func (o *DeviceCompliancePolicySettingStateSummary) SetSettingName(v string)`

SetSettingName sets SettingName field to given value.

### HasSettingName

`func (o *DeviceCompliancePolicySettingStateSummary) HasSettingName() bool`

HasSettingName returns a boolean if a field has been set.

### SetSettingNameNil

`func (o *DeviceCompliancePolicySettingStateSummary) SetSettingNameNil(b bool)`

 SetSettingNameNil sets the value for SettingName to be an explicit nil

### UnsetSettingName
`func (o *DeviceCompliancePolicySettingStateSummary) UnsetSettingName()`

UnsetSettingName ensures that no value is present for SettingName, not even an explicit nil
### GetUnknownDeviceCount

`func (o *DeviceCompliancePolicySettingStateSummary) GetUnknownDeviceCount() int32`

GetUnknownDeviceCount returns the UnknownDeviceCount field if non-nil, zero value otherwise.

### GetUnknownDeviceCountOk

`func (o *DeviceCompliancePolicySettingStateSummary) GetUnknownDeviceCountOk() (*int32, bool)`

GetUnknownDeviceCountOk returns a tuple with the UnknownDeviceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnknownDeviceCount

`func (o *DeviceCompliancePolicySettingStateSummary) SetUnknownDeviceCount(v int32)`

SetUnknownDeviceCount sets UnknownDeviceCount field to given value.

### HasUnknownDeviceCount

`func (o *DeviceCompliancePolicySettingStateSummary) HasUnknownDeviceCount() bool`

HasUnknownDeviceCount returns a boolean if a field has been set.

### GetDeviceComplianceSettingStates

`func (o *DeviceCompliancePolicySettingStateSummary) GetDeviceComplianceSettingStates() []MicrosoftGraphDeviceComplianceSettingState`

GetDeviceComplianceSettingStates returns the DeviceComplianceSettingStates field if non-nil, zero value otherwise.

### GetDeviceComplianceSettingStatesOk

`func (o *DeviceCompliancePolicySettingStateSummary) GetDeviceComplianceSettingStatesOk() (*[]MicrosoftGraphDeviceComplianceSettingState, bool)`

GetDeviceComplianceSettingStatesOk returns a tuple with the DeviceComplianceSettingStates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceComplianceSettingStates

`func (o *DeviceCompliancePolicySettingStateSummary) SetDeviceComplianceSettingStates(v []MicrosoftGraphDeviceComplianceSettingState)`

SetDeviceComplianceSettingStates sets DeviceComplianceSettingStates field to given value.

### HasDeviceComplianceSettingStates

`func (o *DeviceCompliancePolicySettingStateSummary) HasDeviceComplianceSettingStates() bool`

HasDeviceComplianceSettingStates returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


