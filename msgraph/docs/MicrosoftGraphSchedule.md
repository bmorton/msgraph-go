# MicrosoftGraphSchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
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

### NewMicrosoftGraphSchedule

`func NewMicrosoftGraphSchedule() *MicrosoftGraphSchedule`

NewMicrosoftGraphSchedule instantiates a new MicrosoftGraphSchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphScheduleWithDefaults

`func NewMicrosoftGraphScheduleWithDefaults() *MicrosoftGraphSchedule`

NewMicrosoftGraphScheduleWithDefaults instantiates a new MicrosoftGraphSchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphSchedule) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphSchedule) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphSchedule) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphSchedule) HasId() bool`

HasId returns a boolean if a field has been set.

### GetEnabled

`func (o *MicrosoftGraphSchedule) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *MicrosoftGraphSchedule) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *MicrosoftGraphSchedule) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *MicrosoftGraphSchedule) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### SetEnabledNil

`func (o *MicrosoftGraphSchedule) SetEnabledNil(b bool)`

 SetEnabledNil sets the value for Enabled to be an explicit nil

### UnsetEnabled
`func (o *MicrosoftGraphSchedule) UnsetEnabled()`

UnsetEnabled ensures that no value is present for Enabled, not even an explicit nil
### GetOfferShiftRequestsEnabled

`func (o *MicrosoftGraphSchedule) GetOfferShiftRequestsEnabled() bool`

GetOfferShiftRequestsEnabled returns the OfferShiftRequestsEnabled field if non-nil, zero value otherwise.

### GetOfferShiftRequestsEnabledOk

`func (o *MicrosoftGraphSchedule) GetOfferShiftRequestsEnabledOk() (*bool, bool)`

GetOfferShiftRequestsEnabledOk returns a tuple with the OfferShiftRequestsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfferShiftRequestsEnabled

`func (o *MicrosoftGraphSchedule) SetOfferShiftRequestsEnabled(v bool)`

SetOfferShiftRequestsEnabled sets OfferShiftRequestsEnabled field to given value.

### HasOfferShiftRequestsEnabled

`func (o *MicrosoftGraphSchedule) HasOfferShiftRequestsEnabled() bool`

HasOfferShiftRequestsEnabled returns a boolean if a field has been set.

### SetOfferShiftRequestsEnabledNil

`func (o *MicrosoftGraphSchedule) SetOfferShiftRequestsEnabledNil(b bool)`

 SetOfferShiftRequestsEnabledNil sets the value for OfferShiftRequestsEnabled to be an explicit nil

### UnsetOfferShiftRequestsEnabled
`func (o *MicrosoftGraphSchedule) UnsetOfferShiftRequestsEnabled()`

UnsetOfferShiftRequestsEnabled ensures that no value is present for OfferShiftRequestsEnabled, not even an explicit nil
### GetOpenShiftsEnabled

`func (o *MicrosoftGraphSchedule) GetOpenShiftsEnabled() bool`

GetOpenShiftsEnabled returns the OpenShiftsEnabled field if non-nil, zero value otherwise.

### GetOpenShiftsEnabledOk

`func (o *MicrosoftGraphSchedule) GetOpenShiftsEnabledOk() (*bool, bool)`

GetOpenShiftsEnabledOk returns a tuple with the OpenShiftsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenShiftsEnabled

`func (o *MicrosoftGraphSchedule) SetOpenShiftsEnabled(v bool)`

SetOpenShiftsEnabled sets OpenShiftsEnabled field to given value.

### HasOpenShiftsEnabled

`func (o *MicrosoftGraphSchedule) HasOpenShiftsEnabled() bool`

HasOpenShiftsEnabled returns a boolean if a field has been set.

### SetOpenShiftsEnabledNil

`func (o *MicrosoftGraphSchedule) SetOpenShiftsEnabledNil(b bool)`

 SetOpenShiftsEnabledNil sets the value for OpenShiftsEnabled to be an explicit nil

### UnsetOpenShiftsEnabled
`func (o *MicrosoftGraphSchedule) UnsetOpenShiftsEnabled()`

UnsetOpenShiftsEnabled ensures that no value is present for OpenShiftsEnabled, not even an explicit nil
### GetProvisionStatus

`func (o *MicrosoftGraphSchedule) GetProvisionStatus() AnyOfmicrosoftGraphOperationStatus`

GetProvisionStatus returns the ProvisionStatus field if non-nil, zero value otherwise.

### GetProvisionStatusOk

`func (o *MicrosoftGraphSchedule) GetProvisionStatusOk() (*AnyOfmicrosoftGraphOperationStatus, bool)`

GetProvisionStatusOk returns a tuple with the ProvisionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionStatus

`func (o *MicrosoftGraphSchedule) SetProvisionStatus(v AnyOfmicrosoftGraphOperationStatus)`

SetProvisionStatus sets ProvisionStatus field to given value.

### HasProvisionStatus

`func (o *MicrosoftGraphSchedule) HasProvisionStatus() bool`

HasProvisionStatus returns a boolean if a field has been set.

### SetProvisionStatusNil

`func (o *MicrosoftGraphSchedule) SetProvisionStatusNil(b bool)`

 SetProvisionStatusNil sets the value for ProvisionStatus to be an explicit nil

### UnsetProvisionStatus
`func (o *MicrosoftGraphSchedule) UnsetProvisionStatus()`

UnsetProvisionStatus ensures that no value is present for ProvisionStatus, not even an explicit nil
### GetProvisionStatusCode

`func (o *MicrosoftGraphSchedule) GetProvisionStatusCode() string`

GetProvisionStatusCode returns the ProvisionStatusCode field if non-nil, zero value otherwise.

### GetProvisionStatusCodeOk

`func (o *MicrosoftGraphSchedule) GetProvisionStatusCodeOk() (*string, bool)`

GetProvisionStatusCodeOk returns a tuple with the ProvisionStatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionStatusCode

`func (o *MicrosoftGraphSchedule) SetProvisionStatusCode(v string)`

SetProvisionStatusCode sets ProvisionStatusCode field to given value.

### HasProvisionStatusCode

`func (o *MicrosoftGraphSchedule) HasProvisionStatusCode() bool`

HasProvisionStatusCode returns a boolean if a field has been set.

### SetProvisionStatusCodeNil

`func (o *MicrosoftGraphSchedule) SetProvisionStatusCodeNil(b bool)`

 SetProvisionStatusCodeNil sets the value for ProvisionStatusCode to be an explicit nil

### UnsetProvisionStatusCode
`func (o *MicrosoftGraphSchedule) UnsetProvisionStatusCode()`

UnsetProvisionStatusCode ensures that no value is present for ProvisionStatusCode, not even an explicit nil
### GetSwapShiftsRequestsEnabled

`func (o *MicrosoftGraphSchedule) GetSwapShiftsRequestsEnabled() bool`

GetSwapShiftsRequestsEnabled returns the SwapShiftsRequestsEnabled field if non-nil, zero value otherwise.

### GetSwapShiftsRequestsEnabledOk

`func (o *MicrosoftGraphSchedule) GetSwapShiftsRequestsEnabledOk() (*bool, bool)`

GetSwapShiftsRequestsEnabledOk returns a tuple with the SwapShiftsRequestsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwapShiftsRequestsEnabled

`func (o *MicrosoftGraphSchedule) SetSwapShiftsRequestsEnabled(v bool)`

SetSwapShiftsRequestsEnabled sets SwapShiftsRequestsEnabled field to given value.

### HasSwapShiftsRequestsEnabled

`func (o *MicrosoftGraphSchedule) HasSwapShiftsRequestsEnabled() bool`

HasSwapShiftsRequestsEnabled returns a boolean if a field has been set.

### SetSwapShiftsRequestsEnabledNil

`func (o *MicrosoftGraphSchedule) SetSwapShiftsRequestsEnabledNil(b bool)`

 SetSwapShiftsRequestsEnabledNil sets the value for SwapShiftsRequestsEnabled to be an explicit nil

### UnsetSwapShiftsRequestsEnabled
`func (o *MicrosoftGraphSchedule) UnsetSwapShiftsRequestsEnabled()`

UnsetSwapShiftsRequestsEnabled ensures that no value is present for SwapShiftsRequestsEnabled, not even an explicit nil
### GetTimeClockEnabled

`func (o *MicrosoftGraphSchedule) GetTimeClockEnabled() bool`

GetTimeClockEnabled returns the TimeClockEnabled field if non-nil, zero value otherwise.

### GetTimeClockEnabledOk

`func (o *MicrosoftGraphSchedule) GetTimeClockEnabledOk() (*bool, bool)`

GetTimeClockEnabledOk returns a tuple with the TimeClockEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeClockEnabled

`func (o *MicrosoftGraphSchedule) SetTimeClockEnabled(v bool)`

SetTimeClockEnabled sets TimeClockEnabled field to given value.

### HasTimeClockEnabled

`func (o *MicrosoftGraphSchedule) HasTimeClockEnabled() bool`

HasTimeClockEnabled returns a boolean if a field has been set.

### SetTimeClockEnabledNil

`func (o *MicrosoftGraphSchedule) SetTimeClockEnabledNil(b bool)`

 SetTimeClockEnabledNil sets the value for TimeClockEnabled to be an explicit nil

### UnsetTimeClockEnabled
`func (o *MicrosoftGraphSchedule) UnsetTimeClockEnabled()`

UnsetTimeClockEnabled ensures that no value is present for TimeClockEnabled, not even an explicit nil
### GetTimeOffRequestsEnabled

`func (o *MicrosoftGraphSchedule) GetTimeOffRequestsEnabled() bool`

GetTimeOffRequestsEnabled returns the TimeOffRequestsEnabled field if non-nil, zero value otherwise.

### GetTimeOffRequestsEnabledOk

`func (o *MicrosoftGraphSchedule) GetTimeOffRequestsEnabledOk() (*bool, bool)`

GetTimeOffRequestsEnabledOk returns a tuple with the TimeOffRequestsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeOffRequestsEnabled

`func (o *MicrosoftGraphSchedule) SetTimeOffRequestsEnabled(v bool)`

SetTimeOffRequestsEnabled sets TimeOffRequestsEnabled field to given value.

### HasTimeOffRequestsEnabled

`func (o *MicrosoftGraphSchedule) HasTimeOffRequestsEnabled() bool`

HasTimeOffRequestsEnabled returns a boolean if a field has been set.

### SetTimeOffRequestsEnabledNil

`func (o *MicrosoftGraphSchedule) SetTimeOffRequestsEnabledNil(b bool)`

 SetTimeOffRequestsEnabledNil sets the value for TimeOffRequestsEnabled to be an explicit nil

### UnsetTimeOffRequestsEnabled
`func (o *MicrosoftGraphSchedule) UnsetTimeOffRequestsEnabled()`

UnsetTimeOffRequestsEnabled ensures that no value is present for TimeOffRequestsEnabled, not even an explicit nil
### GetTimeZone

`func (o *MicrosoftGraphSchedule) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *MicrosoftGraphSchedule) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *MicrosoftGraphSchedule) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *MicrosoftGraphSchedule) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### SetTimeZoneNil

`func (o *MicrosoftGraphSchedule) SetTimeZoneNil(b bool)`

 SetTimeZoneNil sets the value for TimeZone to be an explicit nil

### UnsetTimeZone
`func (o *MicrosoftGraphSchedule) UnsetTimeZone()`

UnsetTimeZone ensures that no value is present for TimeZone, not even an explicit nil
### GetWorkforceIntegrationIds

`func (o *MicrosoftGraphSchedule) GetWorkforceIntegrationIds() []*string`

GetWorkforceIntegrationIds returns the WorkforceIntegrationIds field if non-nil, zero value otherwise.

### GetWorkforceIntegrationIdsOk

`func (o *MicrosoftGraphSchedule) GetWorkforceIntegrationIdsOk() (*[]*string, bool)`

GetWorkforceIntegrationIdsOk returns a tuple with the WorkforceIntegrationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkforceIntegrationIds

`func (o *MicrosoftGraphSchedule) SetWorkforceIntegrationIds(v []*string)`

SetWorkforceIntegrationIds sets WorkforceIntegrationIds field to given value.

### HasWorkforceIntegrationIds

`func (o *MicrosoftGraphSchedule) HasWorkforceIntegrationIds() bool`

HasWorkforceIntegrationIds returns a boolean if a field has been set.

### GetOfferShiftRequests

`func (o *MicrosoftGraphSchedule) GetOfferShiftRequests() []MicrosoftGraphOfferShiftRequest`

GetOfferShiftRequests returns the OfferShiftRequests field if non-nil, zero value otherwise.

### GetOfferShiftRequestsOk

`func (o *MicrosoftGraphSchedule) GetOfferShiftRequestsOk() (*[]MicrosoftGraphOfferShiftRequest, bool)`

GetOfferShiftRequestsOk returns a tuple with the OfferShiftRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfferShiftRequests

`func (o *MicrosoftGraphSchedule) SetOfferShiftRequests(v []MicrosoftGraphOfferShiftRequest)`

SetOfferShiftRequests sets OfferShiftRequests field to given value.

### HasOfferShiftRequests

`func (o *MicrosoftGraphSchedule) HasOfferShiftRequests() bool`

HasOfferShiftRequests returns a boolean if a field has been set.

### GetOpenShiftChangeRequests

`func (o *MicrosoftGraphSchedule) GetOpenShiftChangeRequests() []MicrosoftGraphOpenShiftChangeRequest`

GetOpenShiftChangeRequests returns the OpenShiftChangeRequests field if non-nil, zero value otherwise.

### GetOpenShiftChangeRequestsOk

`func (o *MicrosoftGraphSchedule) GetOpenShiftChangeRequestsOk() (*[]MicrosoftGraphOpenShiftChangeRequest, bool)`

GetOpenShiftChangeRequestsOk returns a tuple with the OpenShiftChangeRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenShiftChangeRequests

`func (o *MicrosoftGraphSchedule) SetOpenShiftChangeRequests(v []MicrosoftGraphOpenShiftChangeRequest)`

SetOpenShiftChangeRequests sets OpenShiftChangeRequests field to given value.

### HasOpenShiftChangeRequests

`func (o *MicrosoftGraphSchedule) HasOpenShiftChangeRequests() bool`

HasOpenShiftChangeRequests returns a boolean if a field has been set.

### GetOpenShifts

`func (o *MicrosoftGraphSchedule) GetOpenShifts() []MicrosoftGraphOpenShift`

GetOpenShifts returns the OpenShifts field if non-nil, zero value otherwise.

### GetOpenShiftsOk

`func (o *MicrosoftGraphSchedule) GetOpenShiftsOk() (*[]MicrosoftGraphOpenShift, bool)`

GetOpenShiftsOk returns a tuple with the OpenShifts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenShifts

`func (o *MicrosoftGraphSchedule) SetOpenShifts(v []MicrosoftGraphOpenShift)`

SetOpenShifts sets OpenShifts field to given value.

### HasOpenShifts

`func (o *MicrosoftGraphSchedule) HasOpenShifts() bool`

HasOpenShifts returns a boolean if a field has been set.

### GetSchedulingGroups

`func (o *MicrosoftGraphSchedule) GetSchedulingGroups() []MicrosoftGraphSchedulingGroup`

GetSchedulingGroups returns the SchedulingGroups field if non-nil, zero value otherwise.

### GetSchedulingGroupsOk

`func (o *MicrosoftGraphSchedule) GetSchedulingGroupsOk() (*[]MicrosoftGraphSchedulingGroup, bool)`

GetSchedulingGroupsOk returns a tuple with the SchedulingGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedulingGroups

`func (o *MicrosoftGraphSchedule) SetSchedulingGroups(v []MicrosoftGraphSchedulingGroup)`

SetSchedulingGroups sets SchedulingGroups field to given value.

### HasSchedulingGroups

`func (o *MicrosoftGraphSchedule) HasSchedulingGroups() bool`

HasSchedulingGroups returns a boolean if a field has been set.

### GetShifts

`func (o *MicrosoftGraphSchedule) GetShifts() []MicrosoftGraphShift`

GetShifts returns the Shifts field if non-nil, zero value otherwise.

### GetShiftsOk

`func (o *MicrosoftGraphSchedule) GetShiftsOk() (*[]MicrosoftGraphShift, bool)`

GetShiftsOk returns a tuple with the Shifts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShifts

`func (o *MicrosoftGraphSchedule) SetShifts(v []MicrosoftGraphShift)`

SetShifts sets Shifts field to given value.

### HasShifts

`func (o *MicrosoftGraphSchedule) HasShifts() bool`

HasShifts returns a boolean if a field has been set.

### GetSwapShiftsChangeRequests

`func (o *MicrosoftGraphSchedule) GetSwapShiftsChangeRequests() []MicrosoftGraphSwapShiftsChangeRequest`

GetSwapShiftsChangeRequests returns the SwapShiftsChangeRequests field if non-nil, zero value otherwise.

### GetSwapShiftsChangeRequestsOk

`func (o *MicrosoftGraphSchedule) GetSwapShiftsChangeRequestsOk() (*[]MicrosoftGraphSwapShiftsChangeRequest, bool)`

GetSwapShiftsChangeRequestsOk returns a tuple with the SwapShiftsChangeRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwapShiftsChangeRequests

`func (o *MicrosoftGraphSchedule) SetSwapShiftsChangeRequests(v []MicrosoftGraphSwapShiftsChangeRequest)`

SetSwapShiftsChangeRequests sets SwapShiftsChangeRequests field to given value.

### HasSwapShiftsChangeRequests

`func (o *MicrosoftGraphSchedule) HasSwapShiftsChangeRequests() bool`

HasSwapShiftsChangeRequests returns a boolean if a field has been set.

### GetTimeOffReasons

`func (o *MicrosoftGraphSchedule) GetTimeOffReasons() []MicrosoftGraphTimeOffReason`

GetTimeOffReasons returns the TimeOffReasons field if non-nil, zero value otherwise.

### GetTimeOffReasonsOk

`func (o *MicrosoftGraphSchedule) GetTimeOffReasonsOk() (*[]MicrosoftGraphTimeOffReason, bool)`

GetTimeOffReasonsOk returns a tuple with the TimeOffReasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeOffReasons

`func (o *MicrosoftGraphSchedule) SetTimeOffReasons(v []MicrosoftGraphTimeOffReason)`

SetTimeOffReasons sets TimeOffReasons field to given value.

### HasTimeOffReasons

`func (o *MicrosoftGraphSchedule) HasTimeOffReasons() bool`

HasTimeOffReasons returns a boolean if a field has been set.

### GetTimeOffRequests

`func (o *MicrosoftGraphSchedule) GetTimeOffRequests() []MicrosoftGraphTimeOffRequest`

GetTimeOffRequests returns the TimeOffRequests field if non-nil, zero value otherwise.

### GetTimeOffRequestsOk

`func (o *MicrosoftGraphSchedule) GetTimeOffRequestsOk() (*[]MicrosoftGraphTimeOffRequest, bool)`

GetTimeOffRequestsOk returns a tuple with the TimeOffRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeOffRequests

`func (o *MicrosoftGraphSchedule) SetTimeOffRequests(v []MicrosoftGraphTimeOffRequest)`

SetTimeOffRequests sets TimeOffRequests field to given value.

### HasTimeOffRequests

`func (o *MicrosoftGraphSchedule) HasTimeOffRequests() bool`

HasTimeOffRequests returns a boolean if a field has been set.

### GetTimesOff

`func (o *MicrosoftGraphSchedule) GetTimesOff() []MicrosoftGraphTimeOff`

GetTimesOff returns the TimesOff field if non-nil, zero value otherwise.

### GetTimesOffOk

`func (o *MicrosoftGraphSchedule) GetTimesOffOk() (*[]MicrosoftGraphTimeOff, bool)`

GetTimesOffOk returns a tuple with the TimesOff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimesOff

`func (o *MicrosoftGraphSchedule) SetTimesOff(v []MicrosoftGraphTimeOff)`

SetTimesOff sets TimesOff field to given value.

### HasTimesOff

`func (o *MicrosoftGraphSchedule) HasTimesOff() bool`

HasTimesOff returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


