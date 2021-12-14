# Schedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **NullableBool** | Indicates whether the schedule is enabled for the team. Required. | [optional] 
**OfferShiftRequestsEnabled** | Pointer to **NullableBool** | Indicates whether offer shift requests are enabled for the schedule. | [optional] 
**OpenShiftsEnabled** | Pointer to **NullableBool** | Indicates whether open shifts are enabled for the schedule. | [optional] 
**ProvisionStatus** | Pointer to [**AnyOfmicrosoftGraphOperationStatus**](anyOf&lt;microsoft.graph.operationStatus&gt;.md) | The status of the schedule provisioning. The possible values are notStarted, running, completed, failed. | [optional] 
**ProvisionStatusCode** | Pointer to **NullableString** | Additional information about why schedule provisioning failed. | [optional] 
**SwapShiftsRequestsEnabled** | Pointer to **NullableBool** | Indicates whether swap shifts requests are enabled for the schedule. | [optional] 
**TimeClockEnabled** | Pointer to **NullableBool** | Indicates whether time clock is enabled for the schedule. | [optional] 
**TimeOffRequestsEnabled** | Pointer to **NullableBool** | Indicates whether time off requests are enabled for the schedule. | [optional] 
**TimeZone** | Pointer to **NullableString** | Indicates the time zone of the schedule team using tz database format. Required. | [optional] 
**WorkforceIntegrationIds** | Pointer to **[]string** |  | [optional] 
**OfferShiftRequests** | Pointer to [**[]MicrosoftGraphOfferShiftRequest**](MicrosoftGraphOfferShiftRequest.md) |  | [optional] 
**OpenShiftChangeRequests** | Pointer to [**[]MicrosoftGraphOpenShiftChangeRequest**](MicrosoftGraphOpenShiftChangeRequest.md) |  | [optional] 
**OpenShifts** | Pointer to [**[]MicrosoftGraphOpenShift**](MicrosoftGraphOpenShift.md) |  | [optional] 
**SchedulingGroups** | Pointer to [**[]MicrosoftGraphSchedulingGroup**](MicrosoftGraphSchedulingGroup.md) | The logical grouping of users in the schedule (usually by role). | [optional] 
**Shifts** | Pointer to [**[]MicrosoftGraphShift**](MicrosoftGraphShift.md) | The shifts in the schedule. | [optional] 
**SwapShiftsChangeRequests** | Pointer to [**[]MicrosoftGraphSwapShiftsChangeRequest**](MicrosoftGraphSwapShiftsChangeRequest.md) |  | [optional] 
**TimeOffReasons** | Pointer to [**[]MicrosoftGraphTimeOffReason**](MicrosoftGraphTimeOffReason.md) | The set of reasons for a time off in the schedule. | [optional] 
**TimeOffRequests** | Pointer to [**[]MicrosoftGraphTimeOffRequest**](MicrosoftGraphTimeOffRequest.md) |  | [optional] 
**TimesOff** | Pointer to [**[]MicrosoftGraphTimeOff**](MicrosoftGraphTimeOff.md) | The instances of times off in the schedule. | [optional] 

## Methods

### NewSchedule

`func NewSchedule() *Schedule`

NewSchedule instantiates a new Schedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduleWithDefaults

`func NewScheduleWithDefaults() *Schedule`

NewScheduleWithDefaults instantiates a new Schedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *Schedule) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *Schedule) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *Schedule) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *Schedule) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### SetEnabledNil

`func (o *Schedule) SetEnabledNil(b bool)`

 SetEnabledNil sets the value for Enabled to be an explicit nil

### UnsetEnabled
`func (o *Schedule) UnsetEnabled()`

UnsetEnabled ensures that no value is present for Enabled, not even an explicit nil
### GetOfferShiftRequestsEnabled

`func (o *Schedule) GetOfferShiftRequestsEnabled() bool`

GetOfferShiftRequestsEnabled returns the OfferShiftRequestsEnabled field if non-nil, zero value otherwise.

### GetOfferShiftRequestsEnabledOk

`func (o *Schedule) GetOfferShiftRequestsEnabledOk() (*bool, bool)`

GetOfferShiftRequestsEnabledOk returns a tuple with the OfferShiftRequestsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfferShiftRequestsEnabled

`func (o *Schedule) SetOfferShiftRequestsEnabled(v bool)`

SetOfferShiftRequestsEnabled sets OfferShiftRequestsEnabled field to given value.

### HasOfferShiftRequestsEnabled

`func (o *Schedule) HasOfferShiftRequestsEnabled() bool`

HasOfferShiftRequestsEnabled returns a boolean if a field has been set.

### SetOfferShiftRequestsEnabledNil

`func (o *Schedule) SetOfferShiftRequestsEnabledNil(b bool)`

 SetOfferShiftRequestsEnabledNil sets the value for OfferShiftRequestsEnabled to be an explicit nil

### UnsetOfferShiftRequestsEnabled
`func (o *Schedule) UnsetOfferShiftRequestsEnabled()`

UnsetOfferShiftRequestsEnabled ensures that no value is present for OfferShiftRequestsEnabled, not even an explicit nil
### GetOpenShiftsEnabled

`func (o *Schedule) GetOpenShiftsEnabled() bool`

GetOpenShiftsEnabled returns the OpenShiftsEnabled field if non-nil, zero value otherwise.

### GetOpenShiftsEnabledOk

`func (o *Schedule) GetOpenShiftsEnabledOk() (*bool, bool)`

GetOpenShiftsEnabledOk returns a tuple with the OpenShiftsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenShiftsEnabled

`func (o *Schedule) SetOpenShiftsEnabled(v bool)`

SetOpenShiftsEnabled sets OpenShiftsEnabled field to given value.

### HasOpenShiftsEnabled

`func (o *Schedule) HasOpenShiftsEnabled() bool`

HasOpenShiftsEnabled returns a boolean if a field has been set.

### SetOpenShiftsEnabledNil

`func (o *Schedule) SetOpenShiftsEnabledNil(b bool)`

 SetOpenShiftsEnabledNil sets the value for OpenShiftsEnabled to be an explicit nil

### UnsetOpenShiftsEnabled
`func (o *Schedule) UnsetOpenShiftsEnabled()`

UnsetOpenShiftsEnabled ensures that no value is present for OpenShiftsEnabled, not even an explicit nil
### GetProvisionStatus

`func (o *Schedule) GetProvisionStatus() AnyOfmicrosoftGraphOperationStatus`

GetProvisionStatus returns the ProvisionStatus field if non-nil, zero value otherwise.

### GetProvisionStatusOk

`func (o *Schedule) GetProvisionStatusOk() (*AnyOfmicrosoftGraphOperationStatus, bool)`

GetProvisionStatusOk returns a tuple with the ProvisionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionStatus

`func (o *Schedule) SetProvisionStatus(v AnyOfmicrosoftGraphOperationStatus)`

SetProvisionStatus sets ProvisionStatus field to given value.

### HasProvisionStatus

`func (o *Schedule) HasProvisionStatus() bool`

HasProvisionStatus returns a boolean if a field has been set.

### SetProvisionStatusNil

`func (o *Schedule) SetProvisionStatusNil(b bool)`

 SetProvisionStatusNil sets the value for ProvisionStatus to be an explicit nil

### UnsetProvisionStatus
`func (o *Schedule) UnsetProvisionStatus()`

UnsetProvisionStatus ensures that no value is present for ProvisionStatus, not even an explicit nil
### GetProvisionStatusCode

`func (o *Schedule) GetProvisionStatusCode() string`

GetProvisionStatusCode returns the ProvisionStatusCode field if non-nil, zero value otherwise.

### GetProvisionStatusCodeOk

`func (o *Schedule) GetProvisionStatusCodeOk() (*string, bool)`

GetProvisionStatusCodeOk returns a tuple with the ProvisionStatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionStatusCode

`func (o *Schedule) SetProvisionStatusCode(v string)`

SetProvisionStatusCode sets ProvisionStatusCode field to given value.

### HasProvisionStatusCode

`func (o *Schedule) HasProvisionStatusCode() bool`

HasProvisionStatusCode returns a boolean if a field has been set.

### SetProvisionStatusCodeNil

`func (o *Schedule) SetProvisionStatusCodeNil(b bool)`

 SetProvisionStatusCodeNil sets the value for ProvisionStatusCode to be an explicit nil

### UnsetProvisionStatusCode
`func (o *Schedule) UnsetProvisionStatusCode()`

UnsetProvisionStatusCode ensures that no value is present for ProvisionStatusCode, not even an explicit nil
### GetSwapShiftsRequestsEnabled

`func (o *Schedule) GetSwapShiftsRequestsEnabled() bool`

GetSwapShiftsRequestsEnabled returns the SwapShiftsRequestsEnabled field if non-nil, zero value otherwise.

### GetSwapShiftsRequestsEnabledOk

`func (o *Schedule) GetSwapShiftsRequestsEnabledOk() (*bool, bool)`

GetSwapShiftsRequestsEnabledOk returns a tuple with the SwapShiftsRequestsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwapShiftsRequestsEnabled

`func (o *Schedule) SetSwapShiftsRequestsEnabled(v bool)`

SetSwapShiftsRequestsEnabled sets SwapShiftsRequestsEnabled field to given value.

### HasSwapShiftsRequestsEnabled

`func (o *Schedule) HasSwapShiftsRequestsEnabled() bool`

HasSwapShiftsRequestsEnabled returns a boolean if a field has been set.

### SetSwapShiftsRequestsEnabledNil

`func (o *Schedule) SetSwapShiftsRequestsEnabledNil(b bool)`

 SetSwapShiftsRequestsEnabledNil sets the value for SwapShiftsRequestsEnabled to be an explicit nil

### UnsetSwapShiftsRequestsEnabled
`func (o *Schedule) UnsetSwapShiftsRequestsEnabled()`

UnsetSwapShiftsRequestsEnabled ensures that no value is present for SwapShiftsRequestsEnabled, not even an explicit nil
### GetTimeClockEnabled

`func (o *Schedule) GetTimeClockEnabled() bool`

GetTimeClockEnabled returns the TimeClockEnabled field if non-nil, zero value otherwise.

### GetTimeClockEnabledOk

`func (o *Schedule) GetTimeClockEnabledOk() (*bool, bool)`

GetTimeClockEnabledOk returns a tuple with the TimeClockEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeClockEnabled

`func (o *Schedule) SetTimeClockEnabled(v bool)`

SetTimeClockEnabled sets TimeClockEnabled field to given value.

### HasTimeClockEnabled

`func (o *Schedule) HasTimeClockEnabled() bool`

HasTimeClockEnabled returns a boolean if a field has been set.

### SetTimeClockEnabledNil

`func (o *Schedule) SetTimeClockEnabledNil(b bool)`

 SetTimeClockEnabledNil sets the value for TimeClockEnabled to be an explicit nil

### UnsetTimeClockEnabled
`func (o *Schedule) UnsetTimeClockEnabled()`

UnsetTimeClockEnabled ensures that no value is present for TimeClockEnabled, not even an explicit nil
### GetTimeOffRequestsEnabled

`func (o *Schedule) GetTimeOffRequestsEnabled() bool`

GetTimeOffRequestsEnabled returns the TimeOffRequestsEnabled field if non-nil, zero value otherwise.

### GetTimeOffRequestsEnabledOk

`func (o *Schedule) GetTimeOffRequestsEnabledOk() (*bool, bool)`

GetTimeOffRequestsEnabledOk returns a tuple with the TimeOffRequestsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeOffRequestsEnabled

`func (o *Schedule) SetTimeOffRequestsEnabled(v bool)`

SetTimeOffRequestsEnabled sets TimeOffRequestsEnabled field to given value.

### HasTimeOffRequestsEnabled

`func (o *Schedule) HasTimeOffRequestsEnabled() bool`

HasTimeOffRequestsEnabled returns a boolean if a field has been set.

### SetTimeOffRequestsEnabledNil

`func (o *Schedule) SetTimeOffRequestsEnabledNil(b bool)`

 SetTimeOffRequestsEnabledNil sets the value for TimeOffRequestsEnabled to be an explicit nil

### UnsetTimeOffRequestsEnabled
`func (o *Schedule) UnsetTimeOffRequestsEnabled()`

UnsetTimeOffRequestsEnabled ensures that no value is present for TimeOffRequestsEnabled, not even an explicit nil
### GetTimeZone

`func (o *Schedule) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *Schedule) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *Schedule) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *Schedule) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### SetTimeZoneNil

`func (o *Schedule) SetTimeZoneNil(b bool)`

 SetTimeZoneNil sets the value for TimeZone to be an explicit nil

### UnsetTimeZone
`func (o *Schedule) UnsetTimeZone()`

UnsetTimeZone ensures that no value is present for TimeZone, not even an explicit nil
### GetWorkforceIntegrationIds

`func (o *Schedule) GetWorkforceIntegrationIds() []*string`

GetWorkforceIntegrationIds returns the WorkforceIntegrationIds field if non-nil, zero value otherwise.

### GetWorkforceIntegrationIdsOk

`func (o *Schedule) GetWorkforceIntegrationIdsOk() (*[]*string, bool)`

GetWorkforceIntegrationIdsOk returns a tuple with the WorkforceIntegrationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkforceIntegrationIds

`func (o *Schedule) SetWorkforceIntegrationIds(v []*string)`

SetWorkforceIntegrationIds sets WorkforceIntegrationIds field to given value.

### HasWorkforceIntegrationIds

`func (o *Schedule) HasWorkforceIntegrationIds() bool`

HasWorkforceIntegrationIds returns a boolean if a field has been set.

### GetOfferShiftRequests

`func (o *Schedule) GetOfferShiftRequests() []MicrosoftGraphOfferShiftRequest`

GetOfferShiftRequests returns the OfferShiftRequests field if non-nil, zero value otherwise.

### GetOfferShiftRequestsOk

`func (o *Schedule) GetOfferShiftRequestsOk() (*[]MicrosoftGraphOfferShiftRequest, bool)`

GetOfferShiftRequestsOk returns a tuple with the OfferShiftRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfferShiftRequests

`func (o *Schedule) SetOfferShiftRequests(v []MicrosoftGraphOfferShiftRequest)`

SetOfferShiftRequests sets OfferShiftRequests field to given value.

### HasOfferShiftRequests

`func (o *Schedule) HasOfferShiftRequests() bool`

HasOfferShiftRequests returns a boolean if a field has been set.

### GetOpenShiftChangeRequests

`func (o *Schedule) GetOpenShiftChangeRequests() []MicrosoftGraphOpenShiftChangeRequest`

GetOpenShiftChangeRequests returns the OpenShiftChangeRequests field if non-nil, zero value otherwise.

### GetOpenShiftChangeRequestsOk

`func (o *Schedule) GetOpenShiftChangeRequestsOk() (*[]MicrosoftGraphOpenShiftChangeRequest, bool)`

GetOpenShiftChangeRequestsOk returns a tuple with the OpenShiftChangeRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenShiftChangeRequests

`func (o *Schedule) SetOpenShiftChangeRequests(v []MicrosoftGraphOpenShiftChangeRequest)`

SetOpenShiftChangeRequests sets OpenShiftChangeRequests field to given value.

### HasOpenShiftChangeRequests

`func (o *Schedule) HasOpenShiftChangeRequests() bool`

HasOpenShiftChangeRequests returns a boolean if a field has been set.

### GetOpenShifts

`func (o *Schedule) GetOpenShifts() []MicrosoftGraphOpenShift`

GetOpenShifts returns the OpenShifts field if non-nil, zero value otherwise.

### GetOpenShiftsOk

`func (o *Schedule) GetOpenShiftsOk() (*[]MicrosoftGraphOpenShift, bool)`

GetOpenShiftsOk returns a tuple with the OpenShifts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenShifts

`func (o *Schedule) SetOpenShifts(v []MicrosoftGraphOpenShift)`

SetOpenShifts sets OpenShifts field to given value.

### HasOpenShifts

`func (o *Schedule) HasOpenShifts() bool`

HasOpenShifts returns a boolean if a field has been set.

### GetSchedulingGroups

`func (o *Schedule) GetSchedulingGroups() []MicrosoftGraphSchedulingGroup`

GetSchedulingGroups returns the SchedulingGroups field if non-nil, zero value otherwise.

### GetSchedulingGroupsOk

`func (o *Schedule) GetSchedulingGroupsOk() (*[]MicrosoftGraphSchedulingGroup, bool)`

GetSchedulingGroupsOk returns a tuple with the SchedulingGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedulingGroups

`func (o *Schedule) SetSchedulingGroups(v []MicrosoftGraphSchedulingGroup)`

SetSchedulingGroups sets SchedulingGroups field to given value.

### HasSchedulingGroups

`func (o *Schedule) HasSchedulingGroups() bool`

HasSchedulingGroups returns a boolean if a field has been set.

### GetShifts

`func (o *Schedule) GetShifts() []MicrosoftGraphShift`

GetShifts returns the Shifts field if non-nil, zero value otherwise.

### GetShiftsOk

`func (o *Schedule) GetShiftsOk() (*[]MicrosoftGraphShift, bool)`

GetShiftsOk returns a tuple with the Shifts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShifts

`func (o *Schedule) SetShifts(v []MicrosoftGraphShift)`

SetShifts sets Shifts field to given value.

### HasShifts

`func (o *Schedule) HasShifts() bool`

HasShifts returns a boolean if a field has been set.

### GetSwapShiftsChangeRequests

`func (o *Schedule) GetSwapShiftsChangeRequests() []MicrosoftGraphSwapShiftsChangeRequest`

GetSwapShiftsChangeRequests returns the SwapShiftsChangeRequests field if non-nil, zero value otherwise.

### GetSwapShiftsChangeRequestsOk

`func (o *Schedule) GetSwapShiftsChangeRequestsOk() (*[]MicrosoftGraphSwapShiftsChangeRequest, bool)`

GetSwapShiftsChangeRequestsOk returns a tuple with the SwapShiftsChangeRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwapShiftsChangeRequests

`func (o *Schedule) SetSwapShiftsChangeRequests(v []MicrosoftGraphSwapShiftsChangeRequest)`

SetSwapShiftsChangeRequests sets SwapShiftsChangeRequests field to given value.

### HasSwapShiftsChangeRequests

`func (o *Schedule) HasSwapShiftsChangeRequests() bool`

HasSwapShiftsChangeRequests returns a boolean if a field has been set.

### GetTimeOffReasons

`func (o *Schedule) GetTimeOffReasons() []MicrosoftGraphTimeOffReason`

GetTimeOffReasons returns the TimeOffReasons field if non-nil, zero value otherwise.

### GetTimeOffReasonsOk

`func (o *Schedule) GetTimeOffReasonsOk() (*[]MicrosoftGraphTimeOffReason, bool)`

GetTimeOffReasonsOk returns a tuple with the TimeOffReasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeOffReasons

`func (o *Schedule) SetTimeOffReasons(v []MicrosoftGraphTimeOffReason)`

SetTimeOffReasons sets TimeOffReasons field to given value.

### HasTimeOffReasons

`func (o *Schedule) HasTimeOffReasons() bool`

HasTimeOffReasons returns a boolean if a field has been set.

### GetTimeOffRequests

`func (o *Schedule) GetTimeOffRequests() []MicrosoftGraphTimeOffRequest`

GetTimeOffRequests returns the TimeOffRequests field if non-nil, zero value otherwise.

### GetTimeOffRequestsOk

`func (o *Schedule) GetTimeOffRequestsOk() (*[]MicrosoftGraphTimeOffRequest, bool)`

GetTimeOffRequestsOk returns a tuple with the TimeOffRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeOffRequests

`func (o *Schedule) SetTimeOffRequests(v []MicrosoftGraphTimeOffRequest)`

SetTimeOffRequests sets TimeOffRequests field to given value.

### HasTimeOffRequests

`func (o *Schedule) HasTimeOffRequests() bool`

HasTimeOffRequests returns a boolean if a field has been set.

### GetTimesOff

`func (o *Schedule) GetTimesOff() []MicrosoftGraphTimeOff`

GetTimesOff returns the TimesOff field if non-nil, zero value otherwise.

### GetTimesOffOk

`func (o *Schedule) GetTimesOffOk() (*[]MicrosoftGraphTimeOff, bool)`

GetTimesOffOk returns a tuple with the TimesOff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimesOff

`func (o *Schedule) SetTimesOff(v []MicrosoftGraphTimeOff)`

SetTimesOff sets TimesOff field to given value.

### HasTimesOff

`func (o *Schedule) HasTimesOff() bool`

HasTimesOff returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


