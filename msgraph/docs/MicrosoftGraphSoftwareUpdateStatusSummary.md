# MicrosoftGraphSoftwareUpdateStatusSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
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

### NewMicrosoftGraphSoftwareUpdateStatusSummary

`func NewMicrosoftGraphSoftwareUpdateStatusSummary() *MicrosoftGraphSoftwareUpdateStatusSummary`

NewMicrosoftGraphSoftwareUpdateStatusSummary instantiates a new MicrosoftGraphSoftwareUpdateStatusSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphSoftwareUpdateStatusSummaryWithDefaults

`func NewMicrosoftGraphSoftwareUpdateStatusSummaryWithDefaults() *MicrosoftGraphSoftwareUpdateStatusSummary`

NewMicrosoftGraphSoftwareUpdateStatusSummaryWithDefaults instantiates a new MicrosoftGraphSoftwareUpdateStatusSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphSoftwareUpdateStatusSummary) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphSoftwareUpdateStatusSummary) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphSoftwareUpdateStatusSummary) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphSoftwareUpdateStatusSummary) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCompliantDeviceCount

`func (o *MicrosoftGraphSoftwareUpdateStatusSummary) GetCompliantDeviceCount() int32`

GetCompliantDeviceCount returns the CompliantDeviceCount field if non-nil, zero value otherwise.

### GetCompliantDeviceCountOk

`func (o *MicrosoftGraphSoftwareUpdateStatusSummary) GetCompliantDeviceCountOk() (*int32, bool)`

GetCompliantDeviceCountOk returns a tuple with the CompliantDeviceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompliantDeviceCount

`func (o *MicrosoftGraphSoftwareUpdateStatusSummary) SetCompliantDeviceCount(v int32)`

SetCompliantDeviceCount sets CompliantDeviceCount field to given value.

### HasCompliantDeviceCount

`func (o *MicrosoftGraphSoftwareUpdateStatusSummary) HasCompliantDeviceCount() bool`

HasCompliantDeviceCount returns a boolean if a field has been set.

### GetCompliantUserCount

`func (o *MicrosoftGraphSoftwareUpdateStatusSummary) GetCompliantUserCount() int32`

GetCompliantUserCount returns the CompliantUserCount field if non-nil, zero value otherwise.

### GetCompliantUserCountOk

`func (o *MicrosoftGraphSoftwareUpdateStatusSummary) GetCompliantUserCountOk() (*int32, bool)`

GetCompliantUserCountOk returns a tuple with the CompliantUserCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompliantUserCount

`func (o *MicrosoftGraphSoftwareUpdateStatusSummary) SetCompliantUserCount(v int32)`

SetCompliantUserCount sets CompliantUserCount field to given value.

### HasCompliantUserCount

`func (o *MicrosoftGraphSoftwareUpdateStatusSummary) HasCompliantUserCount() bool`

HasCompliantUserCount returns a boolean if a field has been set.

### GetConflictDeviceCount

`func (o *MicrosoftGraphSoftwareUpdateStatusSummary) GetConflictDeviceCount() int32`

GetConflictDeviceCount returns the ConflictDeviceCount field if non-nil, zero value otherwise.

### GetConflictDeviceCountOk

`func (o *MicrosoftGraphSoftwareUpdateStatusSummary) GetConflictDeviceCountOk() (*int32, bool)`

GetConflictDeviceCountOk returns a tuple with the ConflictDeviceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConflictDeviceCount

`func (o *MicrosoftGraphSoftwareUpdateStatusSummary) SetConflictDeviceCount(v int32)`

SetConflictDeviceCount sets ConflictDeviceCount field to given value.

### HasConflictDeviceCount

`func (o *MicrosoftGraphSoftwareUpdateStatusSummary) HasConflictDeviceCount() bool`

HasConflictDeviceCount returns a boolean if a field has been set.

### GetConflictUserCount

`func (o *MicrosoftGraphSoftwareUpdateStatusSummary) GetConflictUserCount() int32`

GetConflictUserCount returns the ConflictUserCount field if non-nil, zero value otherwise.

### GetConflictUserCountOk

`func (o *MicrosoftGraphSoftwareUpdateStatusSummary) GetConflictUserCountOk() (*int32, bool)`

GetConflictUserCountOk returns a tuple with the ConflictUserCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConflictUserCount

`func (o *MicrosoftGraphSoftwareUpdateStatusSummary) SetConflictUserCount(v int32)`

SetConflictUserCount sets ConflictUserCount field to given value.

### HasConflictUserCount

`func (o *MicrosoftGraphSoftwareUpdateStatusSummary) HasConflictUserCount() bool`

HasConflictUserCount returns a boolean if a field has been set.

### GetDisplayName

`func (o *MicrosoftGraphSoftwareUpdateStatusSummary) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphSoftwareUpdateStatusSummary) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *MicrosoftGraphSoftwareUpdateStatusSummary) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *MicrosoftGraphSoftwareUpdateStatusSummary) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *MicrosoftGraphSoftwareUpdateStatusSummary) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *MicrosoftGraphSoftwareUpdateStatusSummary) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetErrorDeviceCount

`func (o *MicrosoftGraphSoftwareUpdateStatusSummary) GetErrorDeviceCount() int32`

GetErrorDeviceCount returns the ErrorDeviceCount field if non-nil, zero value otherwise.

### GetErrorDeviceCountOk

`func (o *MicrosoftGraphSoftwareUpdateStatusSummary) GetErrorDeviceCountOk() (*int32, bool)`

GetErrorDeviceCountOk returns a tuple with the ErrorDeviceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorDeviceCount

`func (o *MicrosoftGraphSoftwareUpdateStatusSummary) SetErrorDeviceCount(v int32)`

SetErrorDeviceCount sets ErrorDeviceCount field to given value.

### HasErrorDeviceCount

`func (o *MicrosoftGraphSoftwareUpdateStatusSummary) HasErrorDeviceCount() bool`

HasErrorDeviceCount returns a boolean if a field has been set.

### GetErrorUserCount

`func (o *MicrosoftGraphSoftwareUpdateStatusSummary) GetErrorUserCount() int32`

GetErrorUserCount returns the ErrorUserCount field if non-nil, zero value otherwise.

### GetErrorUserCountOk

`func (o *MicrosoftGraphSoftwareUpdateStatusSummary) GetErrorUserCountOk() (*int32, bool)`

GetErrorUserCountOk returns a tuple with the ErrorUserCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorUserCount

`func (o *MicrosoftGraphSoftwareUpdateStatusSummary) SetErrorUserCount(v int32)`

SetErrorUserCount sets ErrorUserCount field to given value.

### HasErrorUserCount

`func (o *MicrosoftGraphSoftwareUpdateStatusSummary) HasErrorUserCount() bool`

HasErrorUserCount returns a boolean if a field has been set.

### GetNonCompliantDeviceCount

`func (o *MicrosoftGraphSoftwareUpdateStatusSummary) GetNonCompliantDeviceCount() int32`

GetNonCompliantDeviceCount returns the NonCompliantDeviceCount field if non-nil, zero value otherwise.

### GetNonCompliantDeviceCountOk

`func (o *MicrosoftGraphSoftwareUpdateStatusSummary) GetNonCompliantDeviceCountOk() (*int32, bool)`

GetNonCompliantDeviceCountOk returns a tuple with the NonCompliantDeviceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonCompliantDeviceCount

`func (o *MicrosoftGraphSoftwareUpdateStatusSummary) SetNonCompliantDeviceCount(v int32)`

SetNonCompliantDeviceCount sets NonCompliantDeviceCount field to given value.

### HasNonCompliantDeviceCount

`func (o *MicrosoftGraphSoftwareUpdateStatusSummary) HasNonCompliantDeviceCount() bool`

HasNonCompliantDeviceCount returns a boolean if a field has been set.

### GetNonCompliantUserCount

`func (o *MicrosoftGraphSoftwareUpdateStatusSummary) GetNonCompliantUserCount() int32`

GetNonCompliantUserCount returns the NonCompliantUserCount field if non-nil, zero value otherwise.

### GetNonCompliantUserCountOk

`func (o *MicrosoftGraphSoftwareUpdateStatusSummary) GetNonCompliantUserCountOk() (*int32, bool)`

GetNonCompliantUserCountOk returns a tuple with the NonCompliantUserCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonCompliantUserCount

`func (o *MicrosoftGraphSoftwareUpdateStatusSummary) SetNonCompliantUserCount(v int32)`

SetNonCompliantUserCount sets NonCompliantUserCount field to given value.

### HasNonCompliantUserCount

`func (o *MicrosoftGraphSoftwareUpdateStatusSummary) HasNonCompliantUserCount() bool`

HasNonCompliantUserCount returns a boolean if a field has been set.

### GetNotApplicableDeviceCount

`func (o *MicrosoftGraphSoftwareUpdateStatusSummary) GetNotApplicableDeviceCount() int32`

GetNotApplicableDeviceCount returns the NotApplicableDeviceCount field if non-nil, zero value otherwise.

### GetNotApplicableDeviceCountOk

`func (o *MicrosoftGraphSoftwareUpdateStatusSummary) GetNotApplicableDeviceCountOk() (*int32, bool)`

GetNotApplicableDeviceCountOk returns a tuple with the NotApplicableDeviceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotApplicableDeviceCount

`func (o *MicrosoftGraphSoftwareUpdateStatusSummary) SetNotApplicableDeviceCount(v int32)`

SetNotApplicableDeviceCount sets NotApplicableDeviceCount field to given value.

### HasNotApplicableDeviceCount

`func (o *MicrosoftGraphSoftwareUpdateStatusSummary) HasNotApplicableDeviceCount() bool`

HasNotApplicableDeviceCount returns a boolean if a field has been set.

### GetNotApplicableUserCount

`func (o *MicrosoftGraphSoftwareUpdateStatusSummary) GetNotApplicableUserCount() int32`

GetNotApplicableUserCount returns the NotApplicableUserCount field if non-nil, zero value otherwise.

### GetNotApplicableUserCountOk

`func (o *MicrosoftGraphSoftwareUpdateStatusSummary) GetNotApplicableUserCountOk() (*int32, bool)`

GetNotApplicableUserCountOk returns a tuple with the NotApplicableUserCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotApplicableUserCount

`func (o *MicrosoftGraphSoftwareUpdateStatusSummary) SetNotApplicableUserCount(v int32)`

SetNotApplicableUserCount sets NotApplicableUserCount field to given value.

### HasNotApplicableUserCount

`func (o *MicrosoftGraphSoftwareUpdateStatusSummary) HasNotApplicableUserCount() bool`

HasNotApplicableUserCount returns a boolean if a field has been set.

### GetRemediatedDeviceCount

`func (o *MicrosoftGraphSoftwareUpdateStatusSummary) GetRemediatedDeviceCount() int32`

GetRemediatedDeviceCount returns the RemediatedDeviceCount field if non-nil, zero value otherwise.

### GetRemediatedDeviceCountOk

`func (o *MicrosoftGraphSoftwareUpdateStatusSummary) GetRemediatedDeviceCountOk() (*int32, bool)`

GetRemediatedDeviceCountOk returns a tuple with the RemediatedDeviceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemediatedDeviceCount

`func (o *MicrosoftGraphSoftwareUpdateStatusSummary) SetRemediatedDeviceCount(v int32)`

SetRemediatedDeviceCount sets RemediatedDeviceCount field to given value.

### HasRemediatedDeviceCount

`func (o *MicrosoftGraphSoftwareUpdateStatusSummary) HasRemediatedDeviceCount() bool`

HasRemediatedDeviceCount returns a boolean if a field has been set.

### GetRemediatedUserCount

`func (o *MicrosoftGraphSoftwareUpdateStatusSummary) GetRemediatedUserCount() int32`

GetRemediatedUserCount returns the RemediatedUserCount field if non-nil, zero value otherwise.

### GetRemediatedUserCountOk

`func (o *MicrosoftGraphSoftwareUpdateStatusSummary) GetRemediatedUserCountOk() (*int32, bool)`

GetRemediatedUserCountOk returns a tuple with the RemediatedUserCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemediatedUserCount

`func (o *MicrosoftGraphSoftwareUpdateStatusSummary) SetRemediatedUserCount(v int32)`

SetRemediatedUserCount sets RemediatedUserCount field to given value.

### HasRemediatedUserCount

`func (o *MicrosoftGraphSoftwareUpdateStatusSummary) HasRemediatedUserCount() bool`

HasRemediatedUserCount returns a boolean if a field has been set.

### GetUnknownDeviceCount

`func (o *MicrosoftGraphSoftwareUpdateStatusSummary) GetUnknownDeviceCount() int32`

GetUnknownDeviceCount returns the UnknownDeviceCount field if non-nil, zero value otherwise.

### GetUnknownDeviceCountOk

`func (o *MicrosoftGraphSoftwareUpdateStatusSummary) GetUnknownDeviceCountOk() (*int32, bool)`

GetUnknownDeviceCountOk returns a tuple with the UnknownDeviceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnknownDeviceCount

`func (o *MicrosoftGraphSoftwareUpdateStatusSummary) SetUnknownDeviceCount(v int32)`

SetUnknownDeviceCount sets UnknownDeviceCount field to given value.

### HasUnknownDeviceCount

`func (o *MicrosoftGraphSoftwareUpdateStatusSummary) HasUnknownDeviceCount() bool`

HasUnknownDeviceCount returns a boolean if a field has been set.

### GetUnknownUserCount

`func (o *MicrosoftGraphSoftwareUpdateStatusSummary) GetUnknownUserCount() int32`

GetUnknownUserCount returns the UnknownUserCount field if non-nil, zero value otherwise.

### GetUnknownUserCountOk

`func (o *MicrosoftGraphSoftwareUpdateStatusSummary) GetUnknownUserCountOk() (*int32, bool)`

GetUnknownUserCountOk returns a tuple with the UnknownUserCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnknownUserCount

`func (o *MicrosoftGraphSoftwareUpdateStatusSummary) SetUnknownUserCount(v int32)`

SetUnknownUserCount sets UnknownUserCount field to given value.

### HasUnknownUserCount

`func (o *MicrosoftGraphSoftwareUpdateStatusSummary) HasUnknownUserCount() bool`

HasUnknownUserCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


