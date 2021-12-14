# MicrosoftGraphSettingStateDeviceSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
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

### NewMicrosoftGraphSettingStateDeviceSummary

`func NewMicrosoftGraphSettingStateDeviceSummary() *MicrosoftGraphSettingStateDeviceSummary`

NewMicrosoftGraphSettingStateDeviceSummary instantiates a new MicrosoftGraphSettingStateDeviceSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphSettingStateDeviceSummaryWithDefaults

`func NewMicrosoftGraphSettingStateDeviceSummaryWithDefaults() *MicrosoftGraphSettingStateDeviceSummary`

NewMicrosoftGraphSettingStateDeviceSummaryWithDefaults instantiates a new MicrosoftGraphSettingStateDeviceSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphSettingStateDeviceSummary) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphSettingStateDeviceSummary) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphSettingStateDeviceSummary) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphSettingStateDeviceSummary) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCompliantDeviceCount

`func (o *MicrosoftGraphSettingStateDeviceSummary) GetCompliantDeviceCount() int32`

GetCompliantDeviceCount returns the CompliantDeviceCount field if non-nil, zero value otherwise.

### GetCompliantDeviceCountOk

`func (o *MicrosoftGraphSettingStateDeviceSummary) GetCompliantDeviceCountOk() (*int32, bool)`

GetCompliantDeviceCountOk returns a tuple with the CompliantDeviceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompliantDeviceCount

`func (o *MicrosoftGraphSettingStateDeviceSummary) SetCompliantDeviceCount(v int32)`

SetCompliantDeviceCount sets CompliantDeviceCount field to given value.

### HasCompliantDeviceCount

`func (o *MicrosoftGraphSettingStateDeviceSummary) HasCompliantDeviceCount() bool`

HasCompliantDeviceCount returns a boolean if a field has been set.

### GetConflictDeviceCount

`func (o *MicrosoftGraphSettingStateDeviceSummary) GetConflictDeviceCount() int32`

GetConflictDeviceCount returns the ConflictDeviceCount field if non-nil, zero value otherwise.

### GetConflictDeviceCountOk

`func (o *MicrosoftGraphSettingStateDeviceSummary) GetConflictDeviceCountOk() (*int32, bool)`

GetConflictDeviceCountOk returns a tuple with the ConflictDeviceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConflictDeviceCount

`func (o *MicrosoftGraphSettingStateDeviceSummary) SetConflictDeviceCount(v int32)`

SetConflictDeviceCount sets ConflictDeviceCount field to given value.

### HasConflictDeviceCount

`func (o *MicrosoftGraphSettingStateDeviceSummary) HasConflictDeviceCount() bool`

HasConflictDeviceCount returns a boolean if a field has been set.

### GetErrorDeviceCount

`func (o *MicrosoftGraphSettingStateDeviceSummary) GetErrorDeviceCount() int32`

GetErrorDeviceCount returns the ErrorDeviceCount field if non-nil, zero value otherwise.

### GetErrorDeviceCountOk

`func (o *MicrosoftGraphSettingStateDeviceSummary) GetErrorDeviceCountOk() (*int32, bool)`

GetErrorDeviceCountOk returns a tuple with the ErrorDeviceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorDeviceCount

`func (o *MicrosoftGraphSettingStateDeviceSummary) SetErrorDeviceCount(v int32)`

SetErrorDeviceCount sets ErrorDeviceCount field to given value.

### HasErrorDeviceCount

`func (o *MicrosoftGraphSettingStateDeviceSummary) HasErrorDeviceCount() bool`

HasErrorDeviceCount returns a boolean if a field has been set.

### GetInstancePath

`func (o *MicrosoftGraphSettingStateDeviceSummary) GetInstancePath() string`

GetInstancePath returns the InstancePath field if non-nil, zero value otherwise.

### GetInstancePathOk

`func (o *MicrosoftGraphSettingStateDeviceSummary) GetInstancePathOk() (*string, bool)`

GetInstancePathOk returns a tuple with the InstancePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstancePath

`func (o *MicrosoftGraphSettingStateDeviceSummary) SetInstancePath(v string)`

SetInstancePath sets InstancePath field to given value.

### HasInstancePath

`func (o *MicrosoftGraphSettingStateDeviceSummary) HasInstancePath() bool`

HasInstancePath returns a boolean if a field has been set.

### SetInstancePathNil

`func (o *MicrosoftGraphSettingStateDeviceSummary) SetInstancePathNil(b bool)`

 SetInstancePathNil sets the value for InstancePath to be an explicit nil

### UnsetInstancePath
`func (o *MicrosoftGraphSettingStateDeviceSummary) UnsetInstancePath()`

UnsetInstancePath ensures that no value is present for InstancePath, not even an explicit nil
### GetNonCompliantDeviceCount

`func (o *MicrosoftGraphSettingStateDeviceSummary) GetNonCompliantDeviceCount() int32`

GetNonCompliantDeviceCount returns the NonCompliantDeviceCount field if non-nil, zero value otherwise.

### GetNonCompliantDeviceCountOk

`func (o *MicrosoftGraphSettingStateDeviceSummary) GetNonCompliantDeviceCountOk() (*int32, bool)`

GetNonCompliantDeviceCountOk returns a tuple with the NonCompliantDeviceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonCompliantDeviceCount

`func (o *MicrosoftGraphSettingStateDeviceSummary) SetNonCompliantDeviceCount(v int32)`

SetNonCompliantDeviceCount sets NonCompliantDeviceCount field to given value.

### HasNonCompliantDeviceCount

`func (o *MicrosoftGraphSettingStateDeviceSummary) HasNonCompliantDeviceCount() bool`

HasNonCompliantDeviceCount returns a boolean if a field has been set.

### GetNotApplicableDeviceCount

`func (o *MicrosoftGraphSettingStateDeviceSummary) GetNotApplicableDeviceCount() int32`

GetNotApplicableDeviceCount returns the NotApplicableDeviceCount field if non-nil, zero value otherwise.

### GetNotApplicableDeviceCountOk

`func (o *MicrosoftGraphSettingStateDeviceSummary) GetNotApplicableDeviceCountOk() (*int32, bool)`

GetNotApplicableDeviceCountOk returns a tuple with the NotApplicableDeviceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotApplicableDeviceCount

`func (o *MicrosoftGraphSettingStateDeviceSummary) SetNotApplicableDeviceCount(v int32)`

SetNotApplicableDeviceCount sets NotApplicableDeviceCount field to given value.

### HasNotApplicableDeviceCount

`func (o *MicrosoftGraphSettingStateDeviceSummary) HasNotApplicableDeviceCount() bool`

HasNotApplicableDeviceCount returns a boolean if a field has been set.

### GetRemediatedDeviceCount

`func (o *MicrosoftGraphSettingStateDeviceSummary) GetRemediatedDeviceCount() int32`

GetRemediatedDeviceCount returns the RemediatedDeviceCount field if non-nil, zero value otherwise.

### GetRemediatedDeviceCountOk

`func (o *MicrosoftGraphSettingStateDeviceSummary) GetRemediatedDeviceCountOk() (*int32, bool)`

GetRemediatedDeviceCountOk returns a tuple with the RemediatedDeviceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemediatedDeviceCount

`func (o *MicrosoftGraphSettingStateDeviceSummary) SetRemediatedDeviceCount(v int32)`

SetRemediatedDeviceCount sets RemediatedDeviceCount field to given value.

### HasRemediatedDeviceCount

`func (o *MicrosoftGraphSettingStateDeviceSummary) HasRemediatedDeviceCount() bool`

HasRemediatedDeviceCount returns a boolean if a field has been set.

### GetSettingName

`func (o *MicrosoftGraphSettingStateDeviceSummary) GetSettingName() string`

GetSettingName returns the SettingName field if non-nil, zero value otherwise.

### GetSettingNameOk

`func (o *MicrosoftGraphSettingStateDeviceSummary) GetSettingNameOk() (*string, bool)`

GetSettingNameOk returns a tuple with the SettingName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettingName

`func (o *MicrosoftGraphSettingStateDeviceSummary) SetSettingName(v string)`

SetSettingName sets SettingName field to given value.

### HasSettingName

`func (o *MicrosoftGraphSettingStateDeviceSummary) HasSettingName() bool`

HasSettingName returns a boolean if a field has been set.

### SetSettingNameNil

`func (o *MicrosoftGraphSettingStateDeviceSummary) SetSettingNameNil(b bool)`

 SetSettingNameNil sets the value for SettingName to be an explicit nil

### UnsetSettingName
`func (o *MicrosoftGraphSettingStateDeviceSummary) UnsetSettingName()`

UnsetSettingName ensures that no value is present for SettingName, not even an explicit nil
### GetUnknownDeviceCount

`func (o *MicrosoftGraphSettingStateDeviceSummary) GetUnknownDeviceCount() int32`

GetUnknownDeviceCount returns the UnknownDeviceCount field if non-nil, zero value otherwise.

### GetUnknownDeviceCountOk

`func (o *MicrosoftGraphSettingStateDeviceSummary) GetUnknownDeviceCountOk() (*int32, bool)`

GetUnknownDeviceCountOk returns a tuple with the UnknownDeviceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnknownDeviceCount

`func (o *MicrosoftGraphSettingStateDeviceSummary) SetUnknownDeviceCount(v int32)`

SetUnknownDeviceCount sets UnknownDeviceCount field to given value.

### HasUnknownDeviceCount

`func (o *MicrosoftGraphSettingStateDeviceSummary) HasUnknownDeviceCount() bool`

HasUnknownDeviceCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


