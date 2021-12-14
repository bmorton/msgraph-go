# MicrosoftGraphDeviceConfigurationDeviceStateSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**CompliantDeviceCount** | Pointer to **int32** | Number of compliant devices | [optional] 
**ConflictDeviceCount** | Pointer to **int32** | Number of conflict devices | [optional] 
**ErrorDeviceCount** | Pointer to **int32** | Number of error devices | [optional] 
**NonCompliantDeviceCount** | Pointer to **int32** | Number of NonCompliant devices | [optional] 
**NotApplicableDeviceCount** | Pointer to **int32** | Number of not applicable devices | [optional] 
**RemediatedDeviceCount** | Pointer to **int32** | Number of remediated devices | [optional] 
**UnknownDeviceCount** | Pointer to **int32** | Number of unknown devices | [optional] 

## Methods

### NewMicrosoftGraphDeviceConfigurationDeviceStateSummary

`func NewMicrosoftGraphDeviceConfigurationDeviceStateSummary() *MicrosoftGraphDeviceConfigurationDeviceStateSummary`

NewMicrosoftGraphDeviceConfigurationDeviceStateSummary instantiates a new MicrosoftGraphDeviceConfigurationDeviceStateSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphDeviceConfigurationDeviceStateSummaryWithDefaults

`func NewMicrosoftGraphDeviceConfigurationDeviceStateSummaryWithDefaults() *MicrosoftGraphDeviceConfigurationDeviceStateSummary`

NewMicrosoftGraphDeviceConfigurationDeviceStateSummaryWithDefaults instantiates a new MicrosoftGraphDeviceConfigurationDeviceStateSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphDeviceConfigurationDeviceStateSummary) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphDeviceConfigurationDeviceStateSummary) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphDeviceConfigurationDeviceStateSummary) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphDeviceConfigurationDeviceStateSummary) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCompliantDeviceCount

`func (o *MicrosoftGraphDeviceConfigurationDeviceStateSummary) GetCompliantDeviceCount() int32`

GetCompliantDeviceCount returns the CompliantDeviceCount field if non-nil, zero value otherwise.

### GetCompliantDeviceCountOk

`func (o *MicrosoftGraphDeviceConfigurationDeviceStateSummary) GetCompliantDeviceCountOk() (*int32, bool)`

GetCompliantDeviceCountOk returns a tuple with the CompliantDeviceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompliantDeviceCount

`func (o *MicrosoftGraphDeviceConfigurationDeviceStateSummary) SetCompliantDeviceCount(v int32)`

SetCompliantDeviceCount sets CompliantDeviceCount field to given value.

### HasCompliantDeviceCount

`func (o *MicrosoftGraphDeviceConfigurationDeviceStateSummary) HasCompliantDeviceCount() bool`

HasCompliantDeviceCount returns a boolean if a field has been set.

### GetConflictDeviceCount

`func (o *MicrosoftGraphDeviceConfigurationDeviceStateSummary) GetConflictDeviceCount() int32`

GetConflictDeviceCount returns the ConflictDeviceCount field if non-nil, zero value otherwise.

### GetConflictDeviceCountOk

`func (o *MicrosoftGraphDeviceConfigurationDeviceStateSummary) GetConflictDeviceCountOk() (*int32, bool)`

GetConflictDeviceCountOk returns a tuple with the ConflictDeviceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConflictDeviceCount

`func (o *MicrosoftGraphDeviceConfigurationDeviceStateSummary) SetConflictDeviceCount(v int32)`

SetConflictDeviceCount sets ConflictDeviceCount field to given value.

### HasConflictDeviceCount

`func (o *MicrosoftGraphDeviceConfigurationDeviceStateSummary) HasConflictDeviceCount() bool`

HasConflictDeviceCount returns a boolean if a field has been set.

### GetErrorDeviceCount

`func (o *MicrosoftGraphDeviceConfigurationDeviceStateSummary) GetErrorDeviceCount() int32`

GetErrorDeviceCount returns the ErrorDeviceCount field if non-nil, zero value otherwise.

### GetErrorDeviceCountOk

`func (o *MicrosoftGraphDeviceConfigurationDeviceStateSummary) GetErrorDeviceCountOk() (*int32, bool)`

GetErrorDeviceCountOk returns a tuple with the ErrorDeviceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorDeviceCount

`func (o *MicrosoftGraphDeviceConfigurationDeviceStateSummary) SetErrorDeviceCount(v int32)`

SetErrorDeviceCount sets ErrorDeviceCount field to given value.

### HasErrorDeviceCount

`func (o *MicrosoftGraphDeviceConfigurationDeviceStateSummary) HasErrorDeviceCount() bool`

HasErrorDeviceCount returns a boolean if a field has been set.

### GetNonCompliantDeviceCount

`func (o *MicrosoftGraphDeviceConfigurationDeviceStateSummary) GetNonCompliantDeviceCount() int32`

GetNonCompliantDeviceCount returns the NonCompliantDeviceCount field if non-nil, zero value otherwise.

### GetNonCompliantDeviceCountOk

`func (o *MicrosoftGraphDeviceConfigurationDeviceStateSummary) GetNonCompliantDeviceCountOk() (*int32, bool)`

GetNonCompliantDeviceCountOk returns a tuple with the NonCompliantDeviceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonCompliantDeviceCount

`func (o *MicrosoftGraphDeviceConfigurationDeviceStateSummary) SetNonCompliantDeviceCount(v int32)`

SetNonCompliantDeviceCount sets NonCompliantDeviceCount field to given value.

### HasNonCompliantDeviceCount

`func (o *MicrosoftGraphDeviceConfigurationDeviceStateSummary) HasNonCompliantDeviceCount() bool`

HasNonCompliantDeviceCount returns a boolean if a field has been set.

### GetNotApplicableDeviceCount

`func (o *MicrosoftGraphDeviceConfigurationDeviceStateSummary) GetNotApplicableDeviceCount() int32`

GetNotApplicableDeviceCount returns the NotApplicableDeviceCount field if non-nil, zero value otherwise.

### GetNotApplicableDeviceCountOk

`func (o *MicrosoftGraphDeviceConfigurationDeviceStateSummary) GetNotApplicableDeviceCountOk() (*int32, bool)`

GetNotApplicableDeviceCountOk returns a tuple with the NotApplicableDeviceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotApplicableDeviceCount

`func (o *MicrosoftGraphDeviceConfigurationDeviceStateSummary) SetNotApplicableDeviceCount(v int32)`

SetNotApplicableDeviceCount sets NotApplicableDeviceCount field to given value.

### HasNotApplicableDeviceCount

`func (o *MicrosoftGraphDeviceConfigurationDeviceStateSummary) HasNotApplicableDeviceCount() bool`

HasNotApplicableDeviceCount returns a boolean if a field has been set.

### GetRemediatedDeviceCount

`func (o *MicrosoftGraphDeviceConfigurationDeviceStateSummary) GetRemediatedDeviceCount() int32`

GetRemediatedDeviceCount returns the RemediatedDeviceCount field if non-nil, zero value otherwise.

### GetRemediatedDeviceCountOk

`func (o *MicrosoftGraphDeviceConfigurationDeviceStateSummary) GetRemediatedDeviceCountOk() (*int32, bool)`

GetRemediatedDeviceCountOk returns a tuple with the RemediatedDeviceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemediatedDeviceCount

`func (o *MicrosoftGraphDeviceConfigurationDeviceStateSummary) SetRemediatedDeviceCount(v int32)`

SetRemediatedDeviceCount sets RemediatedDeviceCount field to given value.

### HasRemediatedDeviceCount

`func (o *MicrosoftGraphDeviceConfigurationDeviceStateSummary) HasRemediatedDeviceCount() bool`

HasRemediatedDeviceCount returns a boolean if a field has been set.

### GetUnknownDeviceCount

`func (o *MicrosoftGraphDeviceConfigurationDeviceStateSummary) GetUnknownDeviceCount() int32`

GetUnknownDeviceCount returns the UnknownDeviceCount field if non-nil, zero value otherwise.

### GetUnknownDeviceCountOk

`func (o *MicrosoftGraphDeviceConfigurationDeviceStateSummary) GetUnknownDeviceCountOk() (*int32, bool)`

GetUnknownDeviceCountOk returns a tuple with the UnknownDeviceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnknownDeviceCount

`func (o *MicrosoftGraphDeviceConfigurationDeviceStateSummary) SetUnknownDeviceCount(v int32)`

SetUnknownDeviceCount sets UnknownDeviceCount field to given value.

### HasUnknownDeviceCount

`func (o *MicrosoftGraphDeviceConfigurationDeviceStateSummary) HasUnknownDeviceCount() bool`

HasUnknownDeviceCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


