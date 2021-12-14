# SoftwareUpdateStatusSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompliantDeviceCount** | Pointer to **int32** | Number of compliant devices. | [optional] 
**CompliantUserCount** | Pointer to **int32** | Number of compliant users. | [optional] 
**ConflictDeviceCount** | Pointer to **int32** | Number of conflict devices. | [optional] 
**ConflictUserCount** | Pointer to **int32** | Number of conflict users. | [optional] 
**DisplayName** | Pointer to **NullableString** | The name of the policy. | [optional] 
**ErrorDeviceCount** | Pointer to **int32** | Number of devices had error. | [optional] 
**ErrorUserCount** | Pointer to **int32** | Number of users had error. | [optional] 
**NonCompliantDeviceCount** | Pointer to **int32** | Number of non compliant devices. | [optional] 
**NonCompliantUserCount** | Pointer to **int32** | Number of non compliant users. | [optional] 
**NotApplicableDeviceCount** | Pointer to **int32** | Number of not applicable devices. | [optional] 
**NotApplicableUserCount** | Pointer to **int32** | Number of not applicable users. | [optional] 
**RemediatedDeviceCount** | Pointer to **int32** | Number of remediated devices. | [optional] 
**RemediatedUserCount** | Pointer to **int32** | Number of remediated users. | [optional] 
**UnknownDeviceCount** | Pointer to **int32** | Number of unknown devices. | [optional] 
**UnknownUserCount** | Pointer to **int32** | Number of unknown users. | [optional] 

## Methods

### NewSoftwareUpdateStatusSummary

`func NewSoftwareUpdateStatusSummary() *SoftwareUpdateStatusSummary`

NewSoftwareUpdateStatusSummary instantiates a new SoftwareUpdateStatusSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSoftwareUpdateStatusSummaryWithDefaults

`func NewSoftwareUpdateStatusSummaryWithDefaults() *SoftwareUpdateStatusSummary`

NewSoftwareUpdateStatusSummaryWithDefaults instantiates a new SoftwareUpdateStatusSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompliantDeviceCount

`func (o *SoftwareUpdateStatusSummary) GetCompliantDeviceCount() int32`

GetCompliantDeviceCount returns the CompliantDeviceCount field if non-nil, zero value otherwise.

### GetCompliantDeviceCountOk

`func (o *SoftwareUpdateStatusSummary) GetCompliantDeviceCountOk() (*int32, bool)`

GetCompliantDeviceCountOk returns a tuple with the CompliantDeviceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompliantDeviceCount

`func (o *SoftwareUpdateStatusSummary) SetCompliantDeviceCount(v int32)`

SetCompliantDeviceCount sets CompliantDeviceCount field to given value.

### HasCompliantDeviceCount

`func (o *SoftwareUpdateStatusSummary) HasCompliantDeviceCount() bool`

HasCompliantDeviceCount returns a boolean if a field has been set.

### GetCompliantUserCount

`func (o *SoftwareUpdateStatusSummary) GetCompliantUserCount() int32`

GetCompliantUserCount returns the CompliantUserCount field if non-nil, zero value otherwise.

### GetCompliantUserCountOk

`func (o *SoftwareUpdateStatusSummary) GetCompliantUserCountOk() (*int32, bool)`

GetCompliantUserCountOk returns a tuple with the CompliantUserCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompliantUserCount

`func (o *SoftwareUpdateStatusSummary) SetCompliantUserCount(v int32)`

SetCompliantUserCount sets CompliantUserCount field to given value.

### HasCompliantUserCount

`func (o *SoftwareUpdateStatusSummary) HasCompliantUserCount() bool`

HasCompliantUserCount returns a boolean if a field has been set.

### GetConflictDeviceCount

`func (o *SoftwareUpdateStatusSummary) GetConflictDeviceCount() int32`

GetConflictDeviceCount returns the ConflictDeviceCount field if non-nil, zero value otherwise.

### GetConflictDeviceCountOk

`func (o *SoftwareUpdateStatusSummary) GetConflictDeviceCountOk() (*int32, bool)`

GetConflictDeviceCountOk returns a tuple with the ConflictDeviceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConflictDeviceCount

`func (o *SoftwareUpdateStatusSummary) SetConflictDeviceCount(v int32)`

SetConflictDeviceCount sets ConflictDeviceCount field to given value.

### HasConflictDeviceCount

`func (o *SoftwareUpdateStatusSummary) HasConflictDeviceCount() bool`

HasConflictDeviceCount returns a boolean if a field has been set.

### GetConflictUserCount

`func (o *SoftwareUpdateStatusSummary) GetConflictUserCount() int32`

GetConflictUserCount returns the ConflictUserCount field if non-nil, zero value otherwise.

### GetConflictUserCountOk

`func (o *SoftwareUpdateStatusSummary) GetConflictUserCountOk() (*int32, bool)`

GetConflictUserCountOk returns a tuple with the ConflictUserCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConflictUserCount

`func (o *SoftwareUpdateStatusSummary) SetConflictUserCount(v int32)`

SetConflictUserCount sets ConflictUserCount field to given value.

### HasConflictUserCount

`func (o *SoftwareUpdateStatusSummary) HasConflictUserCount() bool`

HasConflictUserCount returns a boolean if a field has been set.

### GetDisplayName

`func (o *SoftwareUpdateStatusSummary) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *SoftwareUpdateStatusSummary) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *SoftwareUpdateStatusSummary) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *SoftwareUpdateStatusSummary) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *SoftwareUpdateStatusSummary) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *SoftwareUpdateStatusSummary) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetErrorDeviceCount

`func (o *SoftwareUpdateStatusSummary) GetErrorDeviceCount() int32`

GetErrorDeviceCount returns the ErrorDeviceCount field if non-nil, zero value otherwise.

### GetErrorDeviceCountOk

`func (o *SoftwareUpdateStatusSummary) GetErrorDeviceCountOk() (*int32, bool)`

GetErrorDeviceCountOk returns a tuple with the ErrorDeviceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorDeviceCount

`func (o *SoftwareUpdateStatusSummary) SetErrorDeviceCount(v int32)`

SetErrorDeviceCount sets ErrorDeviceCount field to given value.

### HasErrorDeviceCount

`func (o *SoftwareUpdateStatusSummary) HasErrorDeviceCount() bool`

HasErrorDeviceCount returns a boolean if a field has been set.

### GetErrorUserCount

`func (o *SoftwareUpdateStatusSummary) GetErrorUserCount() int32`

GetErrorUserCount returns the ErrorUserCount field if non-nil, zero value otherwise.

### GetErrorUserCountOk

`func (o *SoftwareUpdateStatusSummary) GetErrorUserCountOk() (*int32, bool)`

GetErrorUserCountOk returns a tuple with the ErrorUserCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorUserCount

`func (o *SoftwareUpdateStatusSummary) SetErrorUserCount(v int32)`

SetErrorUserCount sets ErrorUserCount field to given value.

### HasErrorUserCount

`func (o *SoftwareUpdateStatusSummary) HasErrorUserCount() bool`

HasErrorUserCount returns a boolean if a field has been set.

### GetNonCompliantDeviceCount

`func (o *SoftwareUpdateStatusSummary) GetNonCompliantDeviceCount() int32`

GetNonCompliantDeviceCount returns the NonCompliantDeviceCount field if non-nil, zero value otherwise.

### GetNonCompliantDeviceCountOk

`func (o *SoftwareUpdateStatusSummary) GetNonCompliantDeviceCountOk() (*int32, bool)`

GetNonCompliantDeviceCountOk returns a tuple with the NonCompliantDeviceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonCompliantDeviceCount

`func (o *SoftwareUpdateStatusSummary) SetNonCompliantDeviceCount(v int32)`

SetNonCompliantDeviceCount sets NonCompliantDeviceCount field to given value.

### HasNonCompliantDeviceCount

`func (o *SoftwareUpdateStatusSummary) HasNonCompliantDeviceCount() bool`

HasNonCompliantDeviceCount returns a boolean if a field has been set.

### GetNonCompliantUserCount

`func (o *SoftwareUpdateStatusSummary) GetNonCompliantUserCount() int32`

GetNonCompliantUserCount returns the NonCompliantUserCount field if non-nil, zero value otherwise.

### GetNonCompliantUserCountOk

`func (o *SoftwareUpdateStatusSummary) GetNonCompliantUserCountOk() (*int32, bool)`

GetNonCompliantUserCountOk returns a tuple with the NonCompliantUserCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonCompliantUserCount

`func (o *SoftwareUpdateStatusSummary) SetNonCompliantUserCount(v int32)`

SetNonCompliantUserCount sets NonCompliantUserCount field to given value.

### HasNonCompliantUserCount

`func (o *SoftwareUpdateStatusSummary) HasNonCompliantUserCount() bool`

HasNonCompliantUserCount returns a boolean if a field has been set.

### GetNotApplicableDeviceCount

`func (o *SoftwareUpdateStatusSummary) GetNotApplicableDeviceCount() int32`

GetNotApplicableDeviceCount returns the NotApplicableDeviceCount field if non-nil, zero value otherwise.

### GetNotApplicableDeviceCountOk

`func (o *SoftwareUpdateStatusSummary) GetNotApplicableDeviceCountOk() (*int32, bool)`

GetNotApplicableDeviceCountOk returns a tuple with the NotApplicableDeviceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotApplicableDeviceCount

`func (o *SoftwareUpdateStatusSummary) SetNotApplicableDeviceCount(v int32)`

SetNotApplicableDeviceCount sets NotApplicableDeviceCount field to given value.

### HasNotApplicableDeviceCount

`func (o *SoftwareUpdateStatusSummary) HasNotApplicableDeviceCount() bool`

HasNotApplicableDeviceCount returns a boolean if a field has been set.

### GetNotApplicableUserCount

`func (o *SoftwareUpdateStatusSummary) GetNotApplicableUserCount() int32`

GetNotApplicableUserCount returns the NotApplicableUserCount field if non-nil, zero value otherwise.

### GetNotApplicableUserCountOk

`func (o *SoftwareUpdateStatusSummary) GetNotApplicableUserCountOk() (*int32, bool)`

GetNotApplicableUserCountOk returns a tuple with the NotApplicableUserCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotApplicableUserCount

`func (o *SoftwareUpdateStatusSummary) SetNotApplicableUserCount(v int32)`

SetNotApplicableUserCount sets NotApplicableUserCount field to given value.

### HasNotApplicableUserCount

`func (o *SoftwareUpdateStatusSummary) HasNotApplicableUserCount() bool`

HasNotApplicableUserCount returns a boolean if a field has been set.

### GetRemediatedDeviceCount

`func (o *SoftwareUpdateStatusSummary) GetRemediatedDeviceCount() int32`

GetRemediatedDeviceCount returns the RemediatedDeviceCount field if non-nil, zero value otherwise.

### GetRemediatedDeviceCountOk

`func (o *SoftwareUpdateStatusSummary) GetRemediatedDeviceCountOk() (*int32, bool)`

GetRemediatedDeviceCountOk returns a tuple with the RemediatedDeviceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemediatedDeviceCount

`func (o *SoftwareUpdateStatusSummary) SetRemediatedDeviceCount(v int32)`

SetRemediatedDeviceCount sets RemediatedDeviceCount field to given value.

### HasRemediatedDeviceCount

`func (o *SoftwareUpdateStatusSummary) HasRemediatedDeviceCount() bool`

HasRemediatedDeviceCount returns a boolean if a field has been set.

### GetRemediatedUserCount

`func (o *SoftwareUpdateStatusSummary) GetRemediatedUserCount() int32`

GetRemediatedUserCount returns the RemediatedUserCount field if non-nil, zero value otherwise.

### GetRemediatedUserCountOk

`func (o *SoftwareUpdateStatusSummary) GetRemediatedUserCountOk() (*int32, bool)`

GetRemediatedUserCountOk returns a tuple with the RemediatedUserCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemediatedUserCount

`func (o *SoftwareUpdateStatusSummary) SetRemediatedUserCount(v int32)`

SetRemediatedUserCount sets RemediatedUserCount field to given value.

### HasRemediatedUserCount

`func (o *SoftwareUpdateStatusSummary) HasRemediatedUserCount() bool`

HasRemediatedUserCount returns a boolean if a field has been set.

### GetUnknownDeviceCount

`func (o *SoftwareUpdateStatusSummary) GetUnknownDeviceCount() int32`

GetUnknownDeviceCount returns the UnknownDeviceCount field if non-nil, zero value otherwise.

### GetUnknownDeviceCountOk

`func (o *SoftwareUpdateStatusSummary) GetUnknownDeviceCountOk() (*int32, bool)`

GetUnknownDeviceCountOk returns a tuple with the UnknownDeviceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnknownDeviceCount

`func (o *SoftwareUpdateStatusSummary) SetUnknownDeviceCount(v int32)`

SetUnknownDeviceCount sets UnknownDeviceCount field to given value.

### HasUnknownDeviceCount

`func (o *SoftwareUpdateStatusSummary) HasUnknownDeviceCount() bool`

HasUnknownDeviceCount returns a boolean if a field has been set.

### GetUnknownUserCount

`func (o *SoftwareUpdateStatusSummary) GetUnknownUserCount() int32`

GetUnknownUserCount returns the UnknownUserCount field if non-nil, zero value otherwise.

### GetUnknownUserCountOk

`func (o *SoftwareUpdateStatusSummary) GetUnknownUserCountOk() (*int32, bool)`

GetUnknownUserCountOk returns a tuple with the UnknownUserCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnknownUserCount

`func (o *SoftwareUpdateStatusSummary) SetUnknownUserCount(v int32)`

SetUnknownUserCount sets UnknownUserCount field to given value.

### HasUnknownUserCount

`func (o *SoftwareUpdateStatusSummary) HasUnknownUserCount() bool`

HasUnknownUserCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


