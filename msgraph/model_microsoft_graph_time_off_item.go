/*
Partial Graph API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package msgraph

import (
	"encoding/json"
	"time"
)

// MicrosoftGraphTimeOffItem struct for MicrosoftGraphTimeOffItem
type MicrosoftGraphTimeOffItem struct {
	EndDateTime NullableTime `json:"endDateTime,omitempty"`
	StartDateTime NullableTime `json:"startDateTime,omitempty"`
	Theme AnyOfmicrosoftGraphScheduleEntityTheme `json:"theme,omitempty"`
	// ID of the timeOffReason for this timeOffItem. Required.
	TimeOffReasonId NullableString `json:"timeOffReasonId,omitempty"`
}

// NewMicrosoftGraphTimeOffItem instantiates a new MicrosoftGraphTimeOffItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphTimeOffItem() *MicrosoftGraphTimeOffItem {
	this := MicrosoftGraphTimeOffItem{}
	return &this
}

// NewMicrosoftGraphTimeOffItemWithDefaults instantiates a new MicrosoftGraphTimeOffItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphTimeOffItemWithDefaults() *MicrosoftGraphTimeOffItem {
	this := MicrosoftGraphTimeOffItem{}
	return &this
}

// GetEndDateTime returns the EndDateTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphTimeOffItem) GetEndDateTime() time.Time {
	if o == nil || o.EndDateTime.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.EndDateTime.Get()
}

// GetEndDateTimeOk returns a tuple with the EndDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphTimeOffItem) GetEndDateTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.EndDateTime.Get(), o.EndDateTime.IsSet()
}

// HasEndDateTime returns a boolean if a field has been set.
func (o *MicrosoftGraphTimeOffItem) HasEndDateTime() bool {
	if o != nil && o.EndDateTime.IsSet() {
		return true
	}

	return false
}

// SetEndDateTime gets a reference to the given NullableTime and assigns it to the EndDateTime field.
func (o *MicrosoftGraphTimeOffItem) SetEndDateTime(v time.Time) {
	o.EndDateTime.Set(&v)
}
// SetEndDateTimeNil sets the value for EndDateTime to be an explicit nil
func (o *MicrosoftGraphTimeOffItem) SetEndDateTimeNil() {
	o.EndDateTime.Set(nil)
}

// UnsetEndDateTime ensures that no value is present for EndDateTime, not even an explicit nil
func (o *MicrosoftGraphTimeOffItem) UnsetEndDateTime() {
	o.EndDateTime.Unset()
}

// GetStartDateTime returns the StartDateTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphTimeOffItem) GetStartDateTime() time.Time {
	if o == nil || o.StartDateTime.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.StartDateTime.Get()
}

// GetStartDateTimeOk returns a tuple with the StartDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphTimeOffItem) GetStartDateTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.StartDateTime.Get(), o.StartDateTime.IsSet()
}

// HasStartDateTime returns a boolean if a field has been set.
func (o *MicrosoftGraphTimeOffItem) HasStartDateTime() bool {
	if o != nil && o.StartDateTime.IsSet() {
		return true
	}

	return false
}

// SetStartDateTime gets a reference to the given NullableTime and assigns it to the StartDateTime field.
func (o *MicrosoftGraphTimeOffItem) SetStartDateTime(v time.Time) {
	o.StartDateTime.Set(&v)
}
// SetStartDateTimeNil sets the value for StartDateTime to be an explicit nil
func (o *MicrosoftGraphTimeOffItem) SetStartDateTimeNil() {
	o.StartDateTime.Set(nil)
}

// UnsetStartDateTime ensures that no value is present for StartDateTime, not even an explicit nil
func (o *MicrosoftGraphTimeOffItem) UnsetStartDateTime() {
	o.StartDateTime.Unset()
}

// GetTheme returns the Theme field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphTimeOffItem) GetTheme() AnyOfmicrosoftGraphScheduleEntityTheme {
	if o == nil  {
		var ret AnyOfmicrosoftGraphScheduleEntityTheme
		return ret
	}
	return o.Theme
}

// GetThemeOk returns a tuple with the Theme field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphTimeOffItem) GetThemeOk() (*AnyOfmicrosoftGraphScheduleEntityTheme, bool) {
	if o == nil || o.Theme == nil {
		return nil, false
	}
	return &o.Theme, true
}

// HasTheme returns a boolean if a field has been set.
func (o *MicrosoftGraphTimeOffItem) HasTheme() bool {
	if o != nil && o.Theme != nil {
		return true
	}

	return false
}

// SetTheme gets a reference to the given AnyOfmicrosoftGraphScheduleEntityTheme and assigns it to the Theme field.
func (o *MicrosoftGraphTimeOffItem) SetTheme(v AnyOfmicrosoftGraphScheduleEntityTheme) {
	o.Theme = v
}

// GetTimeOffReasonId returns the TimeOffReasonId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphTimeOffItem) GetTimeOffReasonId() string {
	if o == nil || o.TimeOffReasonId.Get() == nil {
		var ret string
		return ret
	}
	return *o.TimeOffReasonId.Get()
}

// GetTimeOffReasonIdOk returns a tuple with the TimeOffReasonId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphTimeOffItem) GetTimeOffReasonIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TimeOffReasonId.Get(), o.TimeOffReasonId.IsSet()
}

// HasTimeOffReasonId returns a boolean if a field has been set.
func (o *MicrosoftGraphTimeOffItem) HasTimeOffReasonId() bool {
	if o != nil && o.TimeOffReasonId.IsSet() {
		return true
	}

	return false
}

// SetTimeOffReasonId gets a reference to the given NullableString and assigns it to the TimeOffReasonId field.
func (o *MicrosoftGraphTimeOffItem) SetTimeOffReasonId(v string) {
	o.TimeOffReasonId.Set(&v)
}
// SetTimeOffReasonIdNil sets the value for TimeOffReasonId to be an explicit nil
func (o *MicrosoftGraphTimeOffItem) SetTimeOffReasonIdNil() {
	o.TimeOffReasonId.Set(nil)
}

// UnsetTimeOffReasonId ensures that no value is present for TimeOffReasonId, not even an explicit nil
func (o *MicrosoftGraphTimeOffItem) UnsetTimeOffReasonId() {
	o.TimeOffReasonId.Unset()
}

func (o MicrosoftGraphTimeOffItem) MarshalJSON() ([]byte, error) {
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
	if o.TimeOffReasonId.IsSet() {
		toSerialize["timeOffReasonId"] = o.TimeOffReasonId.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphTimeOffItem struct {
	value *MicrosoftGraphTimeOffItem
	isSet bool
}

func (v NullableMicrosoftGraphTimeOffItem) Get() *MicrosoftGraphTimeOffItem {
	return v.value
}

func (v *NullableMicrosoftGraphTimeOffItem) Set(val *MicrosoftGraphTimeOffItem) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphTimeOffItem) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphTimeOffItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphTimeOffItem(val *MicrosoftGraphTimeOffItem) *NullableMicrosoftGraphTimeOffItem {
	return &NullableMicrosoftGraphTimeOffItem{value: val, isSet: true}
}

func (v NullableMicrosoftGraphTimeOffItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphTimeOffItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


