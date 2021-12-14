/*
Partial Graph API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// Schedule struct for Schedule
type Schedule struct {
	// Indicates whether the schedule is enabled for the team. Required.
	Enabled NullableBool `json:"enabled,omitempty"`
	// Indicates whether offer shift requests are enabled for the schedule.
	OfferShiftRequestsEnabled NullableBool `json:"offerShiftRequestsEnabled,omitempty"`
	// Indicates whether open shifts are enabled for the schedule.
	OpenShiftsEnabled NullableBool `json:"openShiftsEnabled,omitempty"`
	// The status of the schedule provisioning. The possible values are notStarted, running, completed, failed.
	ProvisionStatus AnyOfmicrosoftGraphOperationStatus `json:"provisionStatus,omitempty"`
	// Additional information about why schedule provisioning failed.
	ProvisionStatusCode NullableString `json:"provisionStatusCode,omitempty"`
	// Indicates whether swap shifts requests are enabled for the schedule.
	SwapShiftsRequestsEnabled NullableBool `json:"swapShiftsRequestsEnabled,omitempty"`
	// Indicates whether time clock is enabled for the schedule.
	TimeClockEnabled NullableBool `json:"timeClockEnabled,omitempty"`
	// Indicates whether time off requests are enabled for the schedule.
	TimeOffRequestsEnabled NullableBool `json:"timeOffRequestsEnabled,omitempty"`
	// Indicates the time zone of the schedule team using tz database format. Required.
	TimeZone NullableString `json:"timeZone,omitempty"`
	WorkforceIntegrationIds *[]*string `json:"workforceIntegrationIds,omitempty"`
	OfferShiftRequests *[]MicrosoftGraphOfferShiftRequest `json:"offerShiftRequests,omitempty"`
	OpenShiftChangeRequests *[]MicrosoftGraphOpenShiftChangeRequest `json:"openShiftChangeRequests,omitempty"`
	OpenShifts *[]MicrosoftGraphOpenShift `json:"openShifts,omitempty"`
	// The logical grouping of users in the schedule (usually by role).
	SchedulingGroups *[]MicrosoftGraphSchedulingGroup `json:"schedulingGroups,omitempty"`
	// The shifts in the schedule.
	Shifts *[]MicrosoftGraphShift `json:"shifts,omitempty"`
	SwapShiftsChangeRequests *[]MicrosoftGraphSwapShiftsChangeRequest `json:"swapShiftsChangeRequests,omitempty"`
	// The set of reasons for a time off in the schedule.
	TimeOffReasons *[]MicrosoftGraphTimeOffReason `json:"timeOffReasons,omitempty"`
	TimeOffRequests *[]MicrosoftGraphTimeOffRequest `json:"timeOffRequests,omitempty"`
	// The instances of times off in the schedule.
	TimesOff *[]MicrosoftGraphTimeOff `json:"timesOff,omitempty"`
}

// NewSchedule instantiates a new Schedule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSchedule() *Schedule {
	this := Schedule{}
	return &this
}

// NewScheduleWithDefaults instantiates a new Schedule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScheduleWithDefaults() *Schedule {
	this := Schedule{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Schedule) GetEnabled() bool {
	if o == nil || o.Enabled.Get() == nil {
		var ret bool
		return ret
	}
	return *o.Enabled.Get()
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Schedule) GetEnabledOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Enabled.Get(), o.Enabled.IsSet()
}

// HasEnabled returns a boolean if a field has been set.
func (o *Schedule) HasEnabled() bool {
	if o != nil && o.Enabled.IsSet() {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given NullableBool and assigns it to the Enabled field.
func (o *Schedule) SetEnabled(v bool) {
	o.Enabled.Set(&v)
}
// SetEnabledNil sets the value for Enabled to be an explicit nil
func (o *Schedule) SetEnabledNil() {
	o.Enabled.Set(nil)
}

// UnsetEnabled ensures that no value is present for Enabled, not even an explicit nil
func (o *Schedule) UnsetEnabled() {
	o.Enabled.Unset()
}

// GetOfferShiftRequestsEnabled returns the OfferShiftRequestsEnabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Schedule) GetOfferShiftRequestsEnabled() bool {
	if o == nil || o.OfferShiftRequestsEnabled.Get() == nil {
		var ret bool
		return ret
	}
	return *o.OfferShiftRequestsEnabled.Get()
}

// GetOfferShiftRequestsEnabledOk returns a tuple with the OfferShiftRequestsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Schedule) GetOfferShiftRequestsEnabledOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.OfferShiftRequestsEnabled.Get(), o.OfferShiftRequestsEnabled.IsSet()
}

// HasOfferShiftRequestsEnabled returns a boolean if a field has been set.
func (o *Schedule) HasOfferShiftRequestsEnabled() bool {
	if o != nil && o.OfferShiftRequestsEnabled.IsSet() {
		return true
	}

	return false
}

// SetOfferShiftRequestsEnabled gets a reference to the given NullableBool and assigns it to the OfferShiftRequestsEnabled field.
func (o *Schedule) SetOfferShiftRequestsEnabled(v bool) {
	o.OfferShiftRequestsEnabled.Set(&v)
}
// SetOfferShiftRequestsEnabledNil sets the value for OfferShiftRequestsEnabled to be an explicit nil
func (o *Schedule) SetOfferShiftRequestsEnabledNil() {
	o.OfferShiftRequestsEnabled.Set(nil)
}

// UnsetOfferShiftRequestsEnabled ensures that no value is present for OfferShiftRequestsEnabled, not even an explicit nil
func (o *Schedule) UnsetOfferShiftRequestsEnabled() {
	o.OfferShiftRequestsEnabled.Unset()
}

// GetOpenShiftsEnabled returns the OpenShiftsEnabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Schedule) GetOpenShiftsEnabled() bool {
	if o == nil || o.OpenShiftsEnabled.Get() == nil {
		var ret bool
		return ret
	}
	return *o.OpenShiftsEnabled.Get()
}

// GetOpenShiftsEnabledOk returns a tuple with the OpenShiftsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Schedule) GetOpenShiftsEnabledOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.OpenShiftsEnabled.Get(), o.OpenShiftsEnabled.IsSet()
}

// HasOpenShiftsEnabled returns a boolean if a field has been set.
func (o *Schedule) HasOpenShiftsEnabled() bool {
	if o != nil && o.OpenShiftsEnabled.IsSet() {
		return true
	}

	return false
}

// SetOpenShiftsEnabled gets a reference to the given NullableBool and assigns it to the OpenShiftsEnabled field.
func (o *Schedule) SetOpenShiftsEnabled(v bool) {
	o.OpenShiftsEnabled.Set(&v)
}
// SetOpenShiftsEnabledNil sets the value for OpenShiftsEnabled to be an explicit nil
func (o *Schedule) SetOpenShiftsEnabledNil() {
	o.OpenShiftsEnabled.Set(nil)
}

// UnsetOpenShiftsEnabled ensures that no value is present for OpenShiftsEnabled, not even an explicit nil
func (o *Schedule) UnsetOpenShiftsEnabled() {
	o.OpenShiftsEnabled.Unset()
}

// GetProvisionStatus returns the ProvisionStatus field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Schedule) GetProvisionStatus() AnyOfmicrosoftGraphOperationStatus {
	if o == nil  {
		var ret AnyOfmicrosoftGraphOperationStatus
		return ret
	}
	return o.ProvisionStatus
}

// GetProvisionStatusOk returns a tuple with the ProvisionStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Schedule) GetProvisionStatusOk() (*AnyOfmicrosoftGraphOperationStatus, bool) {
	if o == nil || o.ProvisionStatus == nil {
		return nil, false
	}
	return &o.ProvisionStatus, true
}

// HasProvisionStatus returns a boolean if a field has been set.
func (o *Schedule) HasProvisionStatus() bool {
	if o != nil && o.ProvisionStatus != nil {
		return true
	}

	return false
}

// SetProvisionStatus gets a reference to the given AnyOfmicrosoftGraphOperationStatus and assigns it to the ProvisionStatus field.
func (o *Schedule) SetProvisionStatus(v AnyOfmicrosoftGraphOperationStatus) {
	o.ProvisionStatus = v
}

// GetProvisionStatusCode returns the ProvisionStatusCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Schedule) GetProvisionStatusCode() string {
	if o == nil || o.ProvisionStatusCode.Get() == nil {
		var ret string
		return ret
	}
	return *o.ProvisionStatusCode.Get()
}

// GetProvisionStatusCodeOk returns a tuple with the ProvisionStatusCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Schedule) GetProvisionStatusCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ProvisionStatusCode.Get(), o.ProvisionStatusCode.IsSet()
}

// HasProvisionStatusCode returns a boolean if a field has been set.
func (o *Schedule) HasProvisionStatusCode() bool {
	if o != nil && o.ProvisionStatusCode.IsSet() {
		return true
	}

	return false
}

// SetProvisionStatusCode gets a reference to the given NullableString and assigns it to the ProvisionStatusCode field.
func (o *Schedule) SetProvisionStatusCode(v string) {
	o.ProvisionStatusCode.Set(&v)
}
// SetProvisionStatusCodeNil sets the value for ProvisionStatusCode to be an explicit nil
func (o *Schedule) SetProvisionStatusCodeNil() {
	o.ProvisionStatusCode.Set(nil)
}

// UnsetProvisionStatusCode ensures that no value is present for ProvisionStatusCode, not even an explicit nil
func (o *Schedule) UnsetProvisionStatusCode() {
	o.ProvisionStatusCode.Unset()
}

// GetSwapShiftsRequestsEnabled returns the SwapShiftsRequestsEnabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Schedule) GetSwapShiftsRequestsEnabled() bool {
	if o == nil || o.SwapShiftsRequestsEnabled.Get() == nil {
		var ret bool
		return ret
	}
	return *o.SwapShiftsRequestsEnabled.Get()
}

// GetSwapShiftsRequestsEnabledOk returns a tuple with the SwapShiftsRequestsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Schedule) GetSwapShiftsRequestsEnabledOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SwapShiftsRequestsEnabled.Get(), o.SwapShiftsRequestsEnabled.IsSet()
}

// HasSwapShiftsRequestsEnabled returns a boolean if a field has been set.
func (o *Schedule) HasSwapShiftsRequestsEnabled() bool {
	if o != nil && o.SwapShiftsRequestsEnabled.IsSet() {
		return true
	}

	return false
}

// SetSwapShiftsRequestsEnabled gets a reference to the given NullableBool and assigns it to the SwapShiftsRequestsEnabled field.
func (o *Schedule) SetSwapShiftsRequestsEnabled(v bool) {
	o.SwapShiftsRequestsEnabled.Set(&v)
}
// SetSwapShiftsRequestsEnabledNil sets the value for SwapShiftsRequestsEnabled to be an explicit nil
func (o *Schedule) SetSwapShiftsRequestsEnabledNil() {
	o.SwapShiftsRequestsEnabled.Set(nil)
}

// UnsetSwapShiftsRequestsEnabled ensures that no value is present for SwapShiftsRequestsEnabled, not even an explicit nil
func (o *Schedule) UnsetSwapShiftsRequestsEnabled() {
	o.SwapShiftsRequestsEnabled.Unset()
}

// GetTimeClockEnabled returns the TimeClockEnabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Schedule) GetTimeClockEnabled() bool {
	if o == nil || o.TimeClockEnabled.Get() == nil {
		var ret bool
		return ret
	}
	return *o.TimeClockEnabled.Get()
}

// GetTimeClockEnabledOk returns a tuple with the TimeClockEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Schedule) GetTimeClockEnabledOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TimeClockEnabled.Get(), o.TimeClockEnabled.IsSet()
}

// HasTimeClockEnabled returns a boolean if a field has been set.
func (o *Schedule) HasTimeClockEnabled() bool {
	if o != nil && o.TimeClockEnabled.IsSet() {
		return true
	}

	return false
}

// SetTimeClockEnabled gets a reference to the given NullableBool and assigns it to the TimeClockEnabled field.
func (o *Schedule) SetTimeClockEnabled(v bool) {
	o.TimeClockEnabled.Set(&v)
}
// SetTimeClockEnabledNil sets the value for TimeClockEnabled to be an explicit nil
func (o *Schedule) SetTimeClockEnabledNil() {
	o.TimeClockEnabled.Set(nil)
}

// UnsetTimeClockEnabled ensures that no value is present for TimeClockEnabled, not even an explicit nil
func (o *Schedule) UnsetTimeClockEnabled() {
	o.TimeClockEnabled.Unset()
}

// GetTimeOffRequestsEnabled returns the TimeOffRequestsEnabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Schedule) GetTimeOffRequestsEnabled() bool {
	if o == nil || o.TimeOffRequestsEnabled.Get() == nil {
		var ret bool
		return ret
	}
	return *o.TimeOffRequestsEnabled.Get()
}

// GetTimeOffRequestsEnabledOk returns a tuple with the TimeOffRequestsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Schedule) GetTimeOffRequestsEnabledOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TimeOffRequestsEnabled.Get(), o.TimeOffRequestsEnabled.IsSet()
}

// HasTimeOffRequestsEnabled returns a boolean if a field has been set.
func (o *Schedule) HasTimeOffRequestsEnabled() bool {
	if o != nil && o.TimeOffRequestsEnabled.IsSet() {
		return true
	}

	return false
}

// SetTimeOffRequestsEnabled gets a reference to the given NullableBool and assigns it to the TimeOffRequestsEnabled field.
func (o *Schedule) SetTimeOffRequestsEnabled(v bool) {
	o.TimeOffRequestsEnabled.Set(&v)
}
// SetTimeOffRequestsEnabledNil sets the value for TimeOffRequestsEnabled to be an explicit nil
func (o *Schedule) SetTimeOffRequestsEnabledNil() {
	o.TimeOffRequestsEnabled.Set(nil)
}

// UnsetTimeOffRequestsEnabled ensures that no value is present for TimeOffRequestsEnabled, not even an explicit nil
func (o *Schedule) UnsetTimeOffRequestsEnabled() {
	o.TimeOffRequestsEnabled.Unset()
}

// GetTimeZone returns the TimeZone field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Schedule) GetTimeZone() string {
	if o == nil || o.TimeZone.Get() == nil {
		var ret string
		return ret
	}
	return *o.TimeZone.Get()
}

// GetTimeZoneOk returns a tuple with the TimeZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Schedule) GetTimeZoneOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TimeZone.Get(), o.TimeZone.IsSet()
}

// HasTimeZone returns a boolean if a field has been set.
func (o *Schedule) HasTimeZone() bool {
	if o != nil && o.TimeZone.IsSet() {
		return true
	}

	return false
}

// SetTimeZone gets a reference to the given NullableString and assigns it to the TimeZone field.
func (o *Schedule) SetTimeZone(v string) {
	o.TimeZone.Set(&v)
}
// SetTimeZoneNil sets the value for TimeZone to be an explicit nil
func (o *Schedule) SetTimeZoneNil() {
	o.TimeZone.Set(nil)
}

// UnsetTimeZone ensures that no value is present for TimeZone, not even an explicit nil
func (o *Schedule) UnsetTimeZone() {
	o.TimeZone.Unset()
}

// GetWorkforceIntegrationIds returns the WorkforceIntegrationIds field value if set, zero value otherwise.
func (o *Schedule) GetWorkforceIntegrationIds() []*string {
	if o == nil || o.WorkforceIntegrationIds == nil {
		var ret []*string
		return ret
	}
	return *o.WorkforceIntegrationIds
}

// GetWorkforceIntegrationIdsOk returns a tuple with the WorkforceIntegrationIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Schedule) GetWorkforceIntegrationIdsOk() (*[]*string, bool) {
	if o == nil || o.WorkforceIntegrationIds == nil {
		return nil, false
	}
	return o.WorkforceIntegrationIds, true
}

// HasWorkforceIntegrationIds returns a boolean if a field has been set.
func (o *Schedule) HasWorkforceIntegrationIds() bool {
	if o != nil && o.WorkforceIntegrationIds != nil {
		return true
	}

	return false
}

// SetWorkforceIntegrationIds gets a reference to the given []*string and assigns it to the WorkforceIntegrationIds field.
func (o *Schedule) SetWorkforceIntegrationIds(v []*string) {
	o.WorkforceIntegrationIds = &v
}

// GetOfferShiftRequests returns the OfferShiftRequests field value if set, zero value otherwise.
func (o *Schedule) GetOfferShiftRequests() []MicrosoftGraphOfferShiftRequest {
	if o == nil || o.OfferShiftRequests == nil {
		var ret []MicrosoftGraphOfferShiftRequest
		return ret
	}
	return *o.OfferShiftRequests
}

// GetOfferShiftRequestsOk returns a tuple with the OfferShiftRequests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Schedule) GetOfferShiftRequestsOk() (*[]MicrosoftGraphOfferShiftRequest, bool) {
	if o == nil || o.OfferShiftRequests == nil {
		return nil, false
	}
	return o.OfferShiftRequests, true
}

// HasOfferShiftRequests returns a boolean if a field has been set.
func (o *Schedule) HasOfferShiftRequests() bool {
	if o != nil && o.OfferShiftRequests != nil {
		return true
	}

	return false
}

// SetOfferShiftRequests gets a reference to the given []MicrosoftGraphOfferShiftRequest and assigns it to the OfferShiftRequests field.
func (o *Schedule) SetOfferShiftRequests(v []MicrosoftGraphOfferShiftRequest) {
	o.OfferShiftRequests = &v
}

// GetOpenShiftChangeRequests returns the OpenShiftChangeRequests field value if set, zero value otherwise.
func (o *Schedule) GetOpenShiftChangeRequests() []MicrosoftGraphOpenShiftChangeRequest {
	if o == nil || o.OpenShiftChangeRequests == nil {
		var ret []MicrosoftGraphOpenShiftChangeRequest
		return ret
	}
	return *o.OpenShiftChangeRequests
}

// GetOpenShiftChangeRequestsOk returns a tuple with the OpenShiftChangeRequests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Schedule) GetOpenShiftChangeRequestsOk() (*[]MicrosoftGraphOpenShiftChangeRequest, bool) {
	if o == nil || o.OpenShiftChangeRequests == nil {
		return nil, false
	}
	return o.OpenShiftChangeRequests, true
}

// HasOpenShiftChangeRequests returns a boolean if a field has been set.
func (o *Schedule) HasOpenShiftChangeRequests() bool {
	if o != nil && o.OpenShiftChangeRequests != nil {
		return true
	}

	return false
}

// SetOpenShiftChangeRequests gets a reference to the given []MicrosoftGraphOpenShiftChangeRequest and assigns it to the OpenShiftChangeRequests field.
func (o *Schedule) SetOpenShiftChangeRequests(v []MicrosoftGraphOpenShiftChangeRequest) {
	o.OpenShiftChangeRequests = &v
}

// GetOpenShifts returns the OpenShifts field value if set, zero value otherwise.
func (o *Schedule) GetOpenShifts() []MicrosoftGraphOpenShift {
	if o == nil || o.OpenShifts == nil {
		var ret []MicrosoftGraphOpenShift
		return ret
	}
	return *o.OpenShifts
}

// GetOpenShiftsOk returns a tuple with the OpenShifts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Schedule) GetOpenShiftsOk() (*[]MicrosoftGraphOpenShift, bool) {
	if o == nil || o.OpenShifts == nil {
		return nil, false
	}
	return o.OpenShifts, true
}

// HasOpenShifts returns a boolean if a field has been set.
func (o *Schedule) HasOpenShifts() bool {
	if o != nil && o.OpenShifts != nil {
		return true
	}

	return false
}

// SetOpenShifts gets a reference to the given []MicrosoftGraphOpenShift and assigns it to the OpenShifts field.
func (o *Schedule) SetOpenShifts(v []MicrosoftGraphOpenShift) {
	o.OpenShifts = &v
}

// GetSchedulingGroups returns the SchedulingGroups field value if set, zero value otherwise.
func (o *Schedule) GetSchedulingGroups() []MicrosoftGraphSchedulingGroup {
	if o == nil || o.SchedulingGroups == nil {
		var ret []MicrosoftGraphSchedulingGroup
		return ret
	}
	return *o.SchedulingGroups
}

// GetSchedulingGroupsOk returns a tuple with the SchedulingGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Schedule) GetSchedulingGroupsOk() (*[]MicrosoftGraphSchedulingGroup, bool) {
	if o == nil || o.SchedulingGroups == nil {
		return nil, false
	}
	return o.SchedulingGroups, true
}

// HasSchedulingGroups returns a boolean if a field has been set.
func (o *Schedule) HasSchedulingGroups() bool {
	if o != nil && o.SchedulingGroups != nil {
		return true
	}

	return false
}

// SetSchedulingGroups gets a reference to the given []MicrosoftGraphSchedulingGroup and assigns it to the SchedulingGroups field.
func (o *Schedule) SetSchedulingGroups(v []MicrosoftGraphSchedulingGroup) {
	o.SchedulingGroups = &v
}

// GetShifts returns the Shifts field value if set, zero value otherwise.
func (o *Schedule) GetShifts() []MicrosoftGraphShift {
	if o == nil || o.Shifts == nil {
		var ret []MicrosoftGraphShift
		return ret
	}
	return *o.Shifts
}

// GetShiftsOk returns a tuple with the Shifts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Schedule) GetShiftsOk() (*[]MicrosoftGraphShift, bool) {
	if o == nil || o.Shifts == nil {
		return nil, false
	}
	return o.Shifts, true
}

// HasShifts returns a boolean if a field has been set.
func (o *Schedule) HasShifts() bool {
	if o != nil && o.Shifts != nil {
		return true
	}

	return false
}

// SetShifts gets a reference to the given []MicrosoftGraphShift and assigns it to the Shifts field.
func (o *Schedule) SetShifts(v []MicrosoftGraphShift) {
	o.Shifts = &v
}

// GetSwapShiftsChangeRequests returns the SwapShiftsChangeRequests field value if set, zero value otherwise.
func (o *Schedule) GetSwapShiftsChangeRequests() []MicrosoftGraphSwapShiftsChangeRequest {
	if o == nil || o.SwapShiftsChangeRequests == nil {
		var ret []MicrosoftGraphSwapShiftsChangeRequest
		return ret
	}
	return *o.SwapShiftsChangeRequests
}

// GetSwapShiftsChangeRequestsOk returns a tuple with the SwapShiftsChangeRequests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Schedule) GetSwapShiftsChangeRequestsOk() (*[]MicrosoftGraphSwapShiftsChangeRequest, bool) {
	if o == nil || o.SwapShiftsChangeRequests == nil {
		return nil, false
	}
	return o.SwapShiftsChangeRequests, true
}

// HasSwapShiftsChangeRequests returns a boolean if a field has been set.
func (o *Schedule) HasSwapShiftsChangeRequests() bool {
	if o != nil && o.SwapShiftsChangeRequests != nil {
		return true
	}

	return false
}

// SetSwapShiftsChangeRequests gets a reference to the given []MicrosoftGraphSwapShiftsChangeRequest and assigns it to the SwapShiftsChangeRequests field.
func (o *Schedule) SetSwapShiftsChangeRequests(v []MicrosoftGraphSwapShiftsChangeRequest) {
	o.SwapShiftsChangeRequests = &v
}

// GetTimeOffReasons returns the TimeOffReasons field value if set, zero value otherwise.
func (o *Schedule) GetTimeOffReasons() []MicrosoftGraphTimeOffReason {
	if o == nil || o.TimeOffReasons == nil {
		var ret []MicrosoftGraphTimeOffReason
		return ret
	}
	return *o.TimeOffReasons
}

// GetTimeOffReasonsOk returns a tuple with the TimeOffReasons field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Schedule) GetTimeOffReasonsOk() (*[]MicrosoftGraphTimeOffReason, bool) {
	if o == nil || o.TimeOffReasons == nil {
		return nil, false
	}
	return o.TimeOffReasons, true
}

// HasTimeOffReasons returns a boolean if a field has been set.
func (o *Schedule) HasTimeOffReasons() bool {
	if o != nil && o.TimeOffReasons != nil {
		return true
	}

	return false
}

// SetTimeOffReasons gets a reference to the given []MicrosoftGraphTimeOffReason and assigns it to the TimeOffReasons field.
func (o *Schedule) SetTimeOffReasons(v []MicrosoftGraphTimeOffReason) {
	o.TimeOffReasons = &v
}

// GetTimeOffRequests returns the TimeOffRequests field value if set, zero value otherwise.
func (o *Schedule) GetTimeOffRequests() []MicrosoftGraphTimeOffRequest {
	if o == nil || o.TimeOffRequests == nil {
		var ret []MicrosoftGraphTimeOffRequest
		return ret
	}
	return *o.TimeOffRequests
}

// GetTimeOffRequestsOk returns a tuple with the TimeOffRequests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Schedule) GetTimeOffRequestsOk() (*[]MicrosoftGraphTimeOffRequest, bool) {
	if o == nil || o.TimeOffRequests == nil {
		return nil, false
	}
	return o.TimeOffRequests, true
}

// HasTimeOffRequests returns a boolean if a field has been set.
func (o *Schedule) HasTimeOffRequests() bool {
	if o != nil && o.TimeOffRequests != nil {
		return true
	}

	return false
}

// SetTimeOffRequests gets a reference to the given []MicrosoftGraphTimeOffRequest and assigns it to the TimeOffRequests field.
func (o *Schedule) SetTimeOffRequests(v []MicrosoftGraphTimeOffRequest) {
	o.TimeOffRequests = &v
}

// GetTimesOff returns the TimesOff field value if set, zero value otherwise.
func (o *Schedule) GetTimesOff() []MicrosoftGraphTimeOff {
	if o == nil || o.TimesOff == nil {
		var ret []MicrosoftGraphTimeOff
		return ret
	}
	return *o.TimesOff
}

// GetTimesOffOk returns a tuple with the TimesOff field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Schedule) GetTimesOffOk() (*[]MicrosoftGraphTimeOff, bool) {
	if o == nil || o.TimesOff == nil {
		return nil, false
	}
	return o.TimesOff, true
}

// HasTimesOff returns a boolean if a field has been set.
func (o *Schedule) HasTimesOff() bool {
	if o != nil && o.TimesOff != nil {
		return true
	}

	return false
}

// SetTimesOff gets a reference to the given []MicrosoftGraphTimeOff and assigns it to the TimesOff field.
func (o *Schedule) SetTimesOff(v []MicrosoftGraphTimeOff) {
	o.TimesOff = &v
}

func (o Schedule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Enabled.IsSet() {
		toSerialize["enabled"] = o.Enabled.Get()
	}
	if o.OfferShiftRequestsEnabled.IsSet() {
		toSerialize["offerShiftRequestsEnabled"] = o.OfferShiftRequestsEnabled.Get()
	}
	if o.OpenShiftsEnabled.IsSet() {
		toSerialize["openShiftsEnabled"] = o.OpenShiftsEnabled.Get()
	}
	if o.ProvisionStatus != nil {
		toSerialize["provisionStatus"] = o.ProvisionStatus
	}
	if o.ProvisionStatusCode.IsSet() {
		toSerialize["provisionStatusCode"] = o.ProvisionStatusCode.Get()
	}
	if o.SwapShiftsRequestsEnabled.IsSet() {
		toSerialize["swapShiftsRequestsEnabled"] = o.SwapShiftsRequestsEnabled.Get()
	}
	if o.TimeClockEnabled.IsSet() {
		toSerialize["timeClockEnabled"] = o.TimeClockEnabled.Get()
	}
	if o.TimeOffRequestsEnabled.IsSet() {
		toSerialize["timeOffRequestsEnabled"] = o.TimeOffRequestsEnabled.Get()
	}
	if o.TimeZone.IsSet() {
		toSerialize["timeZone"] = o.TimeZone.Get()
	}
	if o.WorkforceIntegrationIds != nil {
		toSerialize["workforceIntegrationIds"] = o.WorkforceIntegrationIds
	}
	if o.OfferShiftRequests != nil {
		toSerialize["offerShiftRequests"] = o.OfferShiftRequests
	}
	if o.OpenShiftChangeRequests != nil {
		toSerialize["openShiftChangeRequests"] = o.OpenShiftChangeRequests
	}
	if o.OpenShifts != nil {
		toSerialize["openShifts"] = o.OpenShifts
	}
	if o.SchedulingGroups != nil {
		toSerialize["schedulingGroups"] = o.SchedulingGroups
	}
	if o.Shifts != nil {
		toSerialize["shifts"] = o.Shifts
	}
	if o.SwapShiftsChangeRequests != nil {
		toSerialize["swapShiftsChangeRequests"] = o.SwapShiftsChangeRequests
	}
	if o.TimeOffReasons != nil {
		toSerialize["timeOffReasons"] = o.TimeOffReasons
	}
	if o.TimeOffRequests != nil {
		toSerialize["timeOffRequests"] = o.TimeOffRequests
	}
	if o.TimesOff != nil {
		toSerialize["timesOff"] = o.TimesOff
	}
	return json.Marshal(toSerialize)
}

type NullableSchedule struct {
	value *Schedule
	isSet bool
}

func (v NullableSchedule) Get() *Schedule {
	return v.value
}

func (v *NullableSchedule) Set(val *Schedule) {
	v.value = val
	v.isSet = true
}

func (v NullableSchedule) IsSet() bool {
	return v.isSet
}

func (v *NullableSchedule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSchedule(val *Schedule) *NullableSchedule {
	return &NullableSchedule{value: val, isSet: true}
}

func (v NullableSchedule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSchedule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

