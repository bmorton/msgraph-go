# MicrosoftGraphDeviceConfigurationUserOverview

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**ConfigurationVersion** | Pointer to **int32** | Version of the policy for that overview | [optional] 
**ErrorCount** | Pointer to **int32** | Number of error Users | [optional] 
**FailedCount** | Pointer to **int32** | Number of failed Users | [optional] 
**LastUpdateDateTime** | Pointer to **time.Time** | Last update time | [optional] 
**NotApplicableCount** | Pointer to **int32** | Number of not applicable users | [optional] 
**PendingCount** | Pointer to **int32** | Number of pending Users | [optional] 
**SuccessCount** | Pointer to **int32** | Number of succeeded Users | [optional] 

## Methods

### NewMicrosoftGraphDeviceConfigurationUserOverview

`func NewMicrosoftGraphDeviceConfigurationUserOverview() *MicrosoftGraphDeviceConfigurationUserOverview`

NewMicrosoftGraphDeviceConfigurationUserOverview instantiates a new MicrosoftGraphDeviceConfigurationUserOverview object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphDeviceConfigurationUserOverviewWithDefaults

`func NewMicrosoftGraphDeviceConfigurationUserOverviewWithDefaults() *MicrosoftGraphDeviceConfigurationUserOverview`

NewMicrosoftGraphDeviceConfigurationUserOverviewWithDefaults instantiates a new MicrosoftGraphDeviceConfigurationUserOverview object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphDeviceConfigurationUserOverview) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphDeviceConfigurationUserOverview) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphDeviceConfigurationUserOverview) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphDeviceConfigurationUserOverview) HasId() bool`

HasId returns a boolean if a field has been set.

### GetConfigurationVersion

`func (o *MicrosoftGraphDeviceConfigurationUserOverview) GetConfigurationVersion() int32`

GetConfigurationVersion returns the ConfigurationVersion field if non-nil, zero value otherwise.

### GetConfigurationVersionOk

`func (o *MicrosoftGraphDeviceConfigurationUserOverview) GetConfigurationVersionOk() (*int32, bool)`

GetConfigurationVersionOk returns a tuple with the ConfigurationVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationVersion

`func (o *MicrosoftGraphDeviceConfigurationUserOverview) SetConfigurationVersion(v int32)`

SetConfigurationVersion sets ConfigurationVersion field to given value.

### HasConfigurationVersion

`func (o *MicrosoftGraphDeviceConfigurationUserOverview) HasConfigurationVersion() bool`

HasConfigurationVersion returns a boolean if a field has been set.

### GetErrorCount

`func (o *MicrosoftGraphDeviceConfigurationUserOverview) GetErrorCount() int32`

GetErrorCount returns the ErrorCount field if non-nil, zero value otherwise.

### GetErrorCountOk

`func (o *MicrosoftGraphDeviceConfigurationUserOverview) GetErrorCountOk() (*int32, bool)`

GetErrorCountOk returns a tuple with the ErrorCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCount

`func (o *MicrosoftGraphDeviceConfigurationUserOverview) SetErrorCount(v int32)`

SetErrorCount sets ErrorCount field to given value.

### HasErrorCount

`func (o *MicrosoftGraphDeviceConfigurationUserOverview) HasErrorCount() bool`

HasErrorCount returns a boolean if a field has been set.

### GetFailedCount

`func (o *MicrosoftGraphDeviceConfigurationUserOverview) GetFailedCount() int32`

GetFailedCount returns the FailedCount field if non-nil, zero value otherwise.

### GetFailedCountOk

`func (o *MicrosoftGraphDeviceConfigurationUserOverview) GetFailedCountOk() (*int32, bool)`

GetFailedCountOk returns a tuple with the FailedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedCount

`func (o *MicrosoftGraphDeviceConfigurationUserOverview) SetFailedCount(v int32)`

SetFailedCount sets FailedCount field to given value.

### HasFailedCount

`func (o *MicrosoftGraphDeviceConfigurationUserOverview) HasFailedCount() bool`

HasFailedCount returns a boolean if a field has been set.

### GetLastUpdateDateTime

`func (o *MicrosoftGraphDeviceConfigurationUserOverview) GetLastUpdateDateTime() time.Time`

GetLastUpdateDateTime returns the LastUpdateDateTime field if non-nil, zero value otherwise.

### GetLastUpdateDateTimeOk

`func (o *MicrosoftGraphDeviceConfigurationUserOverview) GetLastUpdateDateTimeOk() (*time.Time, bool)`

GetLastUpdateDateTimeOk returns a tuple with the LastUpdateDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdateDateTime

`func (o *MicrosoftGraphDeviceConfigurationUserOverview) SetLastUpdateDateTime(v time.Time)`

SetLastUpdateDateTime sets LastUpdateDateTime field to given value.

### HasLastUpdateDateTime

`func (o *MicrosoftGraphDeviceConfigurationUserOverview) HasLastUpdateDateTime() bool`

HasLastUpdateDateTime returns a boolean if a field has been set.

### GetNotApplicableCount

`func (o *MicrosoftGraphDeviceConfigurationUserOverview) GetNotApplicableCount() int32`

GetNotApplicableCount returns the NotApplicableCount field if non-nil, zero value otherwise.

### GetNotApplicableCountOk

`func (o *MicrosoftGraphDeviceConfigurationUserOverview) GetNotApplicableCountOk() (*int32, bool)`

GetNotApplicableCountOk returns a tuple with the NotApplicableCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotApplicableCount

`func (o *MicrosoftGraphDeviceConfigurationUserOverview) SetNotApplicableCount(v int32)`

SetNotApplicableCount sets NotApplicableCount field to given value.

### HasNotApplicableCount

`func (o *MicrosoftGraphDeviceConfigurationUserOverview) HasNotApplicableCount() bool`

HasNotApplicableCount returns a boolean if a field has been set.

### GetPendingCount

`func (o *MicrosoftGraphDeviceConfigurationUserOverview) GetPendingCount() int32`

GetPendingCount returns the PendingCount field if non-nil, zero value otherwise.

### GetPendingCountOk

`func (o *MicrosoftGraphDeviceConfigurationUserOverview) GetPendingCountOk() (*int32, bool)`

GetPendingCountOk returns a tuple with the PendingCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingCount

`func (o *MicrosoftGraphDeviceConfigurationUserOverview) SetPendingCount(v int32)`

SetPendingCount sets PendingCount field to given value.

### HasPendingCount

`func (o *MicrosoftGraphDeviceConfigurationUserOverview) HasPendingCount() bool`

HasPendingCount returns a boolean if a field has been set.

### GetSuccessCount

`func (o *MicrosoftGraphDeviceConfigurationUserOverview) GetSuccessCount() int32`

GetSuccessCount returns the SuccessCount field if non-nil, zero value otherwise.

### GetSuccessCountOk

`func (o *MicrosoftGraphDeviceConfigurationUserOverview) GetSuccessCountOk() (*int32, bool)`

GetSuccessCountOk returns a tuple with the SuccessCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessCount

`func (o *MicrosoftGraphDeviceConfigurationUserOverview) SetSuccessCount(v int32)`

SetSuccessCount sets SuccessCount field to given value.

### HasSuccessCount

`func (o *MicrosoftGraphDeviceConfigurationUserOverview) HasSuccessCount() bool`

HasSuccessCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


