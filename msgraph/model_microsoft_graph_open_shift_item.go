/*
Partial Graph API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// MicrosoftGraphOpenShiftItem struct for MicrosoftGraphOpenShiftItem
type MicrosoftGraphOpenShiftItem struct {
	EndDateTime NullableTime `json:"endDateTime,omitempty"`
	StartDateTime NullableTime `json:"startDateTime,omitempty"`
	Theme AnyOfmicrosoftGraphScheduleEntityTheme `json:"theme,omitempty"`
	// An incremental part of a shift which can cover details of when and where an employee is during their shift. For example, an assignment or a scheduled break or lunch. Required.
	Activities *[]*AnyOfmicrosoftGraphShiftActivity `json:"activities,omitempty"`
	// The shift label of the shiftItem.
	DisplayName NullableString `json:"displayName,omitempty"`
	// The shift notes for the shiftItem.
	Notes NullableString `json:"notes,omitempty"`
	// Count of the number of slots for the given open shift.
	OpenSlotCount *int32 `json:"openSlotCount,omitempty"`
}

// NewMicrosoftGraphOpenShiftItem instantiates a new MicrosoftGraphOpenShiftItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphOpenShiftItem() *MicrosoftGraphOpenShiftItem {
	this := MicrosoftGraphOpenShiftItem{}
	return &this
}

// NewMicrosoftGraphOpenShiftItemWithDefaults instantiates a new MicrosoftGraphOpenShiftItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphOpenShiftItemWithDefaults() *MicrosoftGraphOpenShiftItem {
	this := MicrosoftGraphOpenShiftItem{}
	return &this
}

// GetEndDateTime returns the EndDateTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphOpenShiftItem) GetEndDateTime() time.Time {
	if o == nil || o.EndDateTime.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.EndDateTime.Get()
}

// GetEndDateTimeOk returns a tuple with the EndDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphOpenShiftItem) GetEndDateTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.EndDateTime.Get(), o.EndDateTime.IsSet()
}

// HasEndDateTime returns a boolean if a field has been set.
func (o *MicrosoftGraphOpenShiftItem) HasEndDateTime() bool {
	if o != nil && o.EndDateTime.IsSet() {
		return true
	}

	return false
}

// SetEndDateTime gets a reference to the given NullableTime and assigns it to the EndDateTime field.
func (o *MicrosoftGraphOpenShiftItem) SetEndDateTime(v time.Time) {
	o.EndDateTime.Set(&v)
}
// SetEndDateTimeNil sets the value for EndDateTime to be an explicit nil
func (o *MicrosoftGraphOpenShiftItem) SetEndDateTimeNil() {
	o.EndDateTime.Set(nil)
}

// UnsetEndDateTime ensures that no value is present for EndDateTime, not even an explicit nil
func (o *MicrosoftGraphOpenShiftItem) UnsetEndDateTime() {
	o.EndDateTime.Unset()
}

// GetStartDateTime returns the StartDateTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphOpenShiftItem) GetStartDateTime() time.Time {
	if o == nil || o.StartDateTime.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.StartDateTime.Get()
}

// GetStartDateTimeOk returns a tuple with the StartDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphOpenShiftItem) GetStartDateTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.StartDateTime.Get(), o.StartDateTime.IsSet()
}

// HasStartDateTime returns a boolean if a field has been set.
func (o *MicrosoftGraphOpenShiftItem) HasStartDateTime() bool {
	if o != nil && o.StartDateTime.IsSet() {
		return true
	}

	return false
}

// SetStartDateTime gets a reference to the given NullableTime and assigns it to the StartDateTime field.
func (o *MicrosoftGraphOpenShiftItem) SetStartDateTime(v time.Time) {
	o.StartDateTime.Set(&v)
}
// SetStartDateTimeNil sets the value for StartDateTime to be an explicit nil
func (o *MicrosoftGraphOpenShiftItem) SetStartDateTimeNil() {
	o.StartDateTime.Set(nil)
}

// UnsetStartDateTime ensures that no value is present for StartDateTime, not even an explicit nil
func (o *MicrosoftGraphOpenShiftItem) UnsetStartDateTime() {
	o.StartDateTime.Unset()
}

// GetTheme returns the Theme field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphOpenShiftItem) GetTheme() AnyOfmicrosoftGraphScheduleEntityTheme {
	if o == nil  {
		var ret AnyOfmicrosoftGraphScheduleEntityTheme
		return ret
	}
	return o.Theme
}

// GetThemeOk returns a tuple with the Theme field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphOpenShiftItem) GetThemeOk() (*AnyOfmicrosoftGraphScheduleEntityTheme, bool) {
	if o == nil || o.Theme == nil {
		return nil, false
	}
	return &o.Theme, true
}

// HasTheme returns a boolean if a field has been set.
func (o *MicrosoftGraphOpenShiftItem) HasTheme() bool {
	if o != nil && o.Theme != nil {
		return true
	}

	return false
}

// SetTheme gets a reference to the given AnyOfmicrosoftGraphScheduleEntityTheme and assigns it to the Theme field.
func (o *MicrosoftGraphOpenShiftItem) SetTheme(v AnyOfmicrosoftGraphScheduleEntityTheme) {
	o.Theme = v
}

// GetActivities returns the Activities field value if set, zero value otherwise.
func (o *MicrosoftGraphOpenShiftItem) GetActivities() []*AnyOfmicrosoftGraphShiftActivity {
	if o == nil || o.Activities == nil {
		var ret []*AnyOfmicrosoftGraphShiftActivity
		return ret
	}
	return *o.Activities
}

// GetActivitiesOk returns a tuple with the Activities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphOpenShiftItem) GetActivitiesOk() (*[]*AnyOfmicrosoftGraphShiftActivity, bool) {
	if o == nil || o.Activities == nil {
		return nil, false
	}
	return o.Activities, true
}

// HasActivities returns a boolean if a field has been set.
func (o *MicrosoftGraphOpenShiftItem) HasActivities() bool {
	if o != nil && o.Activities != nil {
		return true
	}

	return false
}

// SetActivities gets a reference to the given []*AnyOfmicrosoftGraphShiftActivity and assigns it to the Activities field.
func (o *MicrosoftGraphOpenShiftItem) SetActivities(v []*AnyOfmicrosoftGraphShiftActivity) {
	o.Activities = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphOpenShiftItem) GetDisplayName() string {
	if o == nil || o.DisplayName.Get() == nil {
		var ret string
		return ret
	}
	return *o.DisplayName.Get()
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphOpenShiftItem) GetDisplayNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DisplayName.Get(), o.DisplayName.IsSet()
}

// HasDisplayName returns a boolean if a field has been set.
func (o *MicrosoftGraphOpenShiftItem) HasDisplayName() bool {
	if o != nil && o.DisplayName.IsSet() {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given NullableString and assigns it to the DisplayName field.
func (o *MicrosoftGraphOpenShiftItem) SetDisplayName(v string) {
	o.DisplayName.Set(&v)
}
// SetDisplayNameNil sets the value for DisplayName to be an explicit nil
func (o *MicrosoftGraphOpenShiftItem) SetDisplayNameNil() {
	o.DisplayName.Set(nil)
}

// UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
func (o *MicrosoftGraphOpenShiftItem) UnsetDisplayName() {
	o.DisplayName.Unset()
}

// GetNotes returns the Notes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphOpenShiftItem) GetNotes() string {
	if o == nil || o.Notes.Get() == nil {
		var ret string
		return ret
	}
	return *o.Notes.Get()
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphOpenShiftItem) GetNotesOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Notes.Get(), o.Notes.IsSet()
}

// HasNotes returns a boolean if a field has been set.
func (o *MicrosoftGraphOpenShiftItem) HasNotes() bool {
	if o != nil && o.Notes.IsSet() {
		return true
	}

	return false
}

// SetNotes gets a reference to the given NullableString and assigns it to the Notes field.
func (o *MicrosoftGraphOpenShiftItem) SetNotes(v string) {
	o.Notes.Set(&v)
}
// SetNotesNil sets the value for Notes to be an explicit nil
func (o *MicrosoftGraphOpenShiftItem) SetNotesNil() {
	o.Notes.Set(nil)
}

// UnsetNotes ensures that no value is present for Notes, not even an explicit nil
func (o *MicrosoftGraphOpenShiftItem) UnsetNotes() {
	o.Notes.Unset()
}

// GetOpenSlotCount returns the OpenSlotCount field value if set, zero value otherwise.
func (o *MicrosoftGraphOpenShiftItem) GetOpenSlotCount() int32 {
	if o == nil || o.OpenSlotCount == nil {
		var ret int32
		return ret
	}
	return *o.OpenSlotCount
}

// GetOpenSlotCountOk returns a tuple with the OpenSlotCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphOpenShiftItem) GetOpenSlotCountOk() (*int32, bool) {
	if o == nil || o.OpenSlotCount == nil {
		return nil, false
	}
	return o.OpenSlotCount, true
}

// HasOpenSlotCount returns a boolean if a field has been set.
func (o *MicrosoftGraphOpenShiftItem) HasOpenSlotCount() bool {
	if o != nil && o.OpenSlotCount != nil {
		return true
	}

	return false
}

// SetOpenSlotCount gets a reference to the given int32 and assigns it to the OpenSlotCount field.
func (o *MicrosoftGraphOpenShiftItem) SetOpenSlotCount(v int32) {
	o.OpenSlotCount = &v
}

func (o MicrosoftGraphOpenShiftItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.EndDateTime.IsSet() {
		toSerialize["endDateTime"] = o.EndDateTime.Get()
	}
	if o.StartDateTime.IsSet() {
		toSerialize["startDateTime"] = o.StartDateTime.Get()
	}
	if o.Theme != nil {
		toSerialize["theme"] = o.Theme
	}
	if o.Activities != nil {
		toSerialize["activities"] = o.Activities
	}
	if o.DisplayName.IsSet() {
		toSerialize["displayName"] = o.DisplayName.Get()
	}
	if o.Notes.IsSet() {
		toSerialize["notes"] = o.Notes.Get()
	}
	if o.OpenSlotCount != nil {
		toSerialize["openSlotCount"] = o.OpenSlotCount
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphOpenShiftItem struct {
	value *MicrosoftGraphOpenShiftItem
	isSet bool
}

func (v NullableMicrosoftGraphOpenShiftItem) Get() *MicrosoftGraphOpenShiftItem {
	return v.value
}

func (v *NullableMicrosoftGraphOpenShiftItem) Set(val *MicrosoftGraphOpenShiftItem) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphOpenShiftItem) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphOpenShiftItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphOpenShiftItem(val *MicrosoftGraphOpenShiftItem) *NullableMicrosoftGraphOpenShiftItem {
	return &NullableMicrosoftGraphOpenShiftItem{value: val, isSet: true}
}

func (v NullableMicrosoftGraphOpenShiftItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphOpenShiftItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


