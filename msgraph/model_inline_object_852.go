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

// InlineObject852 struct for InlineObject852
type InlineObject852 struct {
	NotifyTeam NullableBool `json:"notifyTeam,omitempty"`
	StartDateTime NullableTime `json:"startDateTime,omitempty"`
	EndDateTime NullableTime `json:"endDateTime,omitempty"`
}

// NewInlineObject852 instantiates a new InlineObject852 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject852() *InlineObject852 {
	this := InlineObject852{}
	var notifyTeam bool = false
	this.NotifyTeam = *NewNullableBool(&notifyTeam)
	return &this
}

// NewInlineObject852WithDefaults instantiates a new InlineObject852 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject852WithDefaults() *InlineObject852 {
	this := InlineObject852{}
	var notifyTeam bool = false
	this.NotifyTeam = *NewNullableBool(&notifyTeam)
	return &this
}

// GetNotifyTeam returns the NotifyTeam field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject852) GetNotifyTeam() bool {
	if o == nil || o.NotifyTeam.Get() == nil {
		var ret bool
		return ret
	}
	return *o.NotifyTeam.Get()
}

// GetNotifyTeamOk returns a tuple with the NotifyTeam field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject852) GetNotifyTeamOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.NotifyTeam.Get(), o.NotifyTeam.IsSet()
}

// HasNotifyTeam returns a boolean if a field has been set.
func (o *InlineObject852) HasNotifyTeam() bool {
	if o != nil && o.NotifyTeam.IsSet() {
		return true
	}

	return false
}

// SetNotifyTeam gets a reference to the given NullableBool and assigns it to the NotifyTeam field.
func (o *InlineObject852) SetNotifyTeam(v bool) {
	o.NotifyTeam.Set(&v)
}
// SetNotifyTeamNil sets the value for NotifyTeam to be an explicit nil
func (o *InlineObject852) SetNotifyTeamNil() {
	o.NotifyTeam.Set(nil)
}

// UnsetNotifyTeam ensures that no value is present for NotifyTeam, not even an explicit nil
func (o *InlineObject852) UnsetNotifyTeam() {
	o.NotifyTeam.Unset()
}

// GetStartDateTime returns the StartDateTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject852) GetStartDateTime() time.Time {
	if o == nil || o.StartDateTime.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.StartDateTime.Get()
}

// GetStartDateTimeOk returns a tuple with the StartDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject852) GetStartDateTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.StartDateTime.Get(), o.StartDateTime.IsSet()
}

// HasStartDateTime returns a boolean if a field has been set.
func (o *InlineObject852) HasStartDateTime() bool {
	if o != nil && o.StartDateTime.IsSet() {
		return true
	}

	return false
}

// SetStartDateTime gets a reference to the given NullableTime and assigns it to the StartDateTime field.
func (o *InlineObject852) SetStartDateTime(v time.Time) {
	o.StartDateTime.Set(&v)
}
// SetStartDateTimeNil sets the value for StartDateTime to be an explicit nil
func (o *InlineObject852) SetStartDateTimeNil() {
	o.StartDateTime.Set(nil)
}

// UnsetStartDateTime ensures that no value is present for StartDateTime, not even an explicit nil
func (o *InlineObject852) UnsetStartDateTime() {
	o.StartDateTime.Unset()
}

// GetEndDateTime returns the EndDateTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject852) GetEndDateTime() time.Time {
	if o == nil || o.EndDateTime.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.EndDateTime.Get()
}

// GetEndDateTimeOk returns a tuple with the EndDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject852) GetEndDateTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.EndDateTime.Get(), o.EndDateTime.IsSet()
}

// HasEndDateTime returns a boolean if a field has been set.
func (o *InlineObject852) HasEndDateTime() bool {
	if o != nil && o.EndDateTime.IsSet() {
		return true
	}

	return false
}

// SetEndDateTime gets a reference to the given NullableTime and assigns it to the EndDateTime field.
func (o *InlineObject852) SetEndDateTime(v time.Time) {
	o.EndDateTime.Set(&v)
}
// SetEndDateTimeNil sets the value for EndDateTime to be an explicit nil
func (o *InlineObject852) SetEndDateTimeNil() {
	o.EndDateTime.Set(nil)
}

// UnsetEndDateTime ensures that no value is present for EndDateTime, not even an explicit nil
func (o *InlineObject852) UnsetEndDateTime() {
	o.EndDateTime.Unset()
}

func (o InlineObject852) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.NotifyTeam.IsSet() {
		toSerialize["notifyTeam"] = o.NotifyTeam.Get()
	}
	if o.StartDateTime.IsSet() {
		toSerialize["startDateTime"] = o.StartDateTime.Get()
	}
	if o.EndDateTime.IsSet() {
		toSerialize["endDateTime"] = o.EndDateTime.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject852 struct {
	value *InlineObject852
	isSet bool
}

func (v NullableInlineObject852) Get() *InlineObject852 {
	return v.value
}

func (v *NullableInlineObject852) Set(val *InlineObject852) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject852) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject852) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject852(val *InlineObject852) *NullableInlineObject852 {
	return &NullableInlineObject852{value: val, isSet: true}
}

func (v NullableInlineObject852) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject852) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


