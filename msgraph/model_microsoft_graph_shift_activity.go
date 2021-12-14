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

// MicrosoftGraphShiftActivity struct for MicrosoftGraphShiftActivity
type MicrosoftGraphShiftActivity struct {
	// Customer defined code for the shiftActivity. Required.
	Code NullableString `json:"code,omitempty"`
	// The name of the shiftActivity. Required.
	DisplayName NullableString `json:"displayName,omitempty"`
	// The end date and time for the shiftActivity. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Required.
	EndDateTime NullableTime `json:"endDateTime,omitempty"`
	// Indicates whether the microsoft.graph.user should be paid for the activity during their shift. Required.
	IsPaid NullableBool `json:"isPaid,omitempty"`
	// The start date and time for the shiftActivity. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Required.
	StartDateTime NullableTime `json:"startDateTime,omitempty"`
	Theme AnyOfmicrosoftGraphScheduleEntityTheme `json:"theme,omitempty"`
}

// NewMicrosoftGraphShiftActivity instantiates a new MicrosoftGraphShiftActivity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphShiftActivity() *MicrosoftGraphShiftActivity {
	this := MicrosoftGraphShiftActivity{}
	return &this
}

// NewMicrosoftGraphShiftActivityWithDefaults instantiates a new MicrosoftGraphShiftActivity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphShiftActivityWithDefaults() *MicrosoftGraphShiftActivity {
	this := MicrosoftGraphShiftActivity{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphShiftActivity) GetCode() string {
	if o == nil || o.Code.Get() == nil {
		var ret string
		return ret
	}
	return *o.Code.Get()
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphShiftActivity) GetCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Code.Get(), o.Code.IsSet()
}

// HasCode returns a boolean if a field has been set.
func (o *MicrosoftGraphShiftActivity) HasCode() bool {
	if o != nil && o.Code.IsSet() {
		return true
	}

	return false
}

// SetCode gets a reference to the given NullableString and assigns it to the Code field.
func (o *MicrosoftGraphShiftActivity) SetCode(v string) {
	o.Code.Set(&v)
}
// SetCodeNil sets the value for Code to be an explicit nil
func (o *MicrosoftGraphShiftActivity) SetCodeNil() {
	o.Code.Set(nil)
}

// UnsetCode ensures that no value is present for Code, not even an explicit nil
func (o *MicrosoftGraphShiftActivity) UnsetCode() {
	o.Code.Unset()
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphShiftActivity) GetDisplayName() string {
	if o == nil || o.DisplayName.Get() == nil {
		var ret string
		return ret
	}
	return *o.DisplayName.Get()
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphShiftActivity) GetDisplayNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DisplayName.Get(), o.DisplayName.IsSet()
}

// HasDisplayName returns a boolean if a field has been set.
func (o *MicrosoftGraphShiftActivity) HasDisplayName() bool {
	if o != nil && o.DisplayName.IsSet() {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given NullableString and assigns it to the DisplayName field.
func (o *MicrosoftGraphShiftActivity) SetDisplayName(v string) {
	o.DisplayName.Set(&v)
}
// SetDisplayNameNil sets the value for DisplayName to be an explicit nil
func (o *MicrosoftGraphShiftActivity) SetDisplayNameNil() {
	o.DisplayName.Set(nil)
}

// UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
func (o *MicrosoftGraphShiftActivity) UnsetDisplayName() {
	o.DisplayName.Unset()
}

// GetEndDateTime returns the EndDateTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphShiftActivity) GetEndDateTime() time.Time {
	if o == nil || o.EndDateTime.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.EndDateTime.Get()
}

// GetEndDateTimeOk returns a tuple with the EndDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphShiftActivity) GetEndDateTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.EndDateTime.Get(), o.EndDateTime.IsSet()
}

// HasEndDateTime returns a boolean if a field has been set.
func (o *MicrosoftGraphShiftActivity) HasEndDateTime() bool {
	if o != nil && o.EndDateTime.IsSet() {
		return true
	}

	return false
}

// SetEndDateTime gets a reference to the given NullableTime and assigns it to the EndDateTime field.
func (o *MicrosoftGraphShiftActivity) SetEndDateTime(v time.Time) {
	o.EndDateTime.Set(&v)
}
// SetEndDateTimeNil sets the value for EndDateTime to be an explicit nil
func (o *MicrosoftGraphShiftActivity) SetEndDateTimeNil() {
	o.EndDateTime.Set(nil)
}

// UnsetEndDateTime ensures that no value is present for EndDateTime, not even an explicit nil
func (o *MicrosoftGraphShiftActivity) UnsetEndDateTime() {
	o.EndDateTime.Unset()
}

// GetIsPaid returns the IsPaid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphShiftActivity) GetIsPaid() bool {
	if o == nil || o.IsPaid.Get() == nil {
		var ret bool
		return ret
	}
	return *o.IsPaid.Get()
}

// GetIsPaidOk returns a tuple with the IsPaid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphShiftActivity) GetIsPaidOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IsPaid.Get(), o.IsPaid.IsSet()
}

// HasIsPaid returns a boolean if a field has been set.
func (o *MicrosoftGraphShiftActivity) HasIsPaid() bool {
	if o != nil && o.IsPaid.IsSet() {
		return true
	}

	return false
}

// SetIsPaid gets a reference to the given NullableBool and assigns it to the IsPaid field.
func (o *MicrosoftGraphShiftActivity) SetIsPaid(v bool) {
	o.IsPaid.Set(&v)
}
// SetIsPaidNil sets the value for IsPaid to be an explicit nil
func (o *MicrosoftGraphShiftActivity) SetIsPaidNil() {
	o.IsPaid.Set(nil)
}

// UnsetIsPaid ensures that no value is present for IsPaid, not even an explicit nil
func (o *MicrosoftGraphShiftActivity) UnsetIsPaid() {
	o.IsPaid.Unset()
}

// GetStartDateTime returns the StartDateTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphShiftActivity) GetStartDateTime() time.Time {
	if o == nil || o.StartDateTime.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.StartDateTime.Get()
}

// GetStartDateTimeOk returns a tuple with the StartDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphShiftActivity) GetStartDateTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.StartDateTime.Get(), o.StartDateTime.IsSet()
}

// HasStartDateTime returns a boolean if a field has been set.
func (o *MicrosoftGraphShiftActivity) HasStartDateTime() bool {
	if o != nil && o.StartDateTime.IsSet() {
		return true
	}

	return false
}

// SetStartDateTime gets a reference to the given NullableTime and assigns it to the StartDateTime field.
func (o *MicrosoftGraphShiftActivity) SetStartDateTime(v time.Time) {
	o.StartDateTime.Set(&v)
}
// SetStartDateTimeNil sets the value for StartDateTime to be an explicit nil
func (o *MicrosoftGraphShiftActivity) SetStartDateTimeNil() {
	o.StartDateTime.Set(nil)
}

// UnsetStartDateTime ensures that no value is present for StartDateTime, not even an explicit nil
func (o *MicrosoftGraphShiftActivity) UnsetStartDateTime() {
	o.StartDateTime.Unset()
}

// GetTheme returns the Theme field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphShiftActivity) GetTheme() AnyOfmicrosoftGraphScheduleEntityTheme {
	if o == nil  {
		var ret AnyOfmicrosoftGraphScheduleEntityTheme
		return ret
	}
	return o.Theme
}

// GetThemeOk returns a tuple with the Theme field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphShiftActivity) GetThemeOk() (*AnyOfmicrosoftGraphScheduleEntityTheme, bool) {
	if o == nil || o.Theme == nil {
		return nil, false
	}
	return &o.Theme, true
}

// HasTheme returns a boolean if a field has been set.
func (o *MicrosoftGraphShiftActivity) HasTheme() bool {
	if o != nil && o.Theme != nil {
		return true
	}

	return false
}

// SetTheme gets a reference to the given AnyOfmicrosoftGraphScheduleEntityTheme and assigns it to the Theme field.
func (o *MicrosoftGraphShiftActivity) SetTheme(v AnyOfmicrosoftGraphScheduleEntityTheme) {
	o.Theme = v
}

func (o MicrosoftGraphShiftActivity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Code.IsSet() {
		toSerialize["code"] = o.Code.Get()
	}
	if o.DisplayName.IsSet() {
		toSerialize["displayName"] = o.DisplayName.Get()
	}
	if o.EndDateTime.IsSet() {
		toSerialize["endDateTime"] = o.EndDateTime.Get()
	}
	if o.IsPaid.IsSet() {
		toSerialize["isPaid"] = o.IsPaid.Get()
	}
	if o.StartDateTime.IsSet() {
		toSerialize["startDateTime"] = o.StartDateTime.Get()
	}
	if o.Theme != nil {
		toSerialize["theme"] = o.Theme
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphShiftActivity struct {
	value *MicrosoftGraphShiftActivity
	isSet bool
}

func (v NullableMicrosoftGraphShiftActivity) Get() *MicrosoftGraphShiftActivity {
	return v.value
}

func (v *NullableMicrosoftGraphShiftActivity) Set(val *MicrosoftGraphShiftActivity) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphShiftActivity) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphShiftActivity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphShiftActivity(val *MicrosoftGraphShiftActivity) *NullableMicrosoftGraphShiftActivity {
	return &NullableMicrosoftGraphShiftActivity{value: val, isSet: true}
}

func (v NullableMicrosoftGraphShiftActivity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphShiftActivity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


