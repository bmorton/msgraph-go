/*
Partial Graph API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package msgraph

import (
	"encoding/json"
)

// InlineObject223 struct for InlineObject223
type InlineObject223 struct {
	Schedules *[]*string `json:"Schedules,omitempty"`
	EndTime AnyOfmicrosoftGraphDateTimeTimeZone `json:"EndTime,omitempty"`
	StartTime AnyOfmicrosoftGraphDateTimeTimeZone `json:"StartTime,omitempty"`
	AvailabilityViewInterval NullableInt32 `json:"AvailabilityViewInterval,omitempty"`
}

// NewInlineObject223 instantiates a new InlineObject223 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject223() *InlineObject223 {
	this := InlineObject223{}
	return &this
}

// NewInlineObject223WithDefaults instantiates a new InlineObject223 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject223WithDefaults() *InlineObject223 {
	this := InlineObject223{}
	return &this
}

// GetSchedules returns the Schedules field value if set, zero value otherwise.
func (o *InlineObject223) GetSchedules() []*string {
	if o == nil || o.Schedules == nil {
		var ret []*string
		return ret
	}
	return *o.Schedules
}

// GetSchedulesOk returns a tuple with the Schedules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject223) GetSchedulesOk() (*[]*string, bool) {
	if o == nil || o.Schedules == nil {
		return nil, false
	}
	return o.Schedules, true
}

// HasSchedules returns a boolean if a field has been set.
func (o *InlineObject223) HasSchedules() bool {
	if o != nil && o.Schedules != nil {
		return true
	}

	return false
}

// SetSchedules gets a reference to the given []*string and assigns it to the Schedules field.
func (o *InlineObject223) SetSchedules(v []*string) {
	o.Schedules = &v
}

// GetEndTime returns the EndTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject223) GetEndTime() AnyOfmicrosoftGraphDateTimeTimeZone {
	if o == nil  {
		var ret AnyOfmicrosoftGraphDateTimeTimeZone
		return ret
	}
	return o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject223) GetEndTimeOk() (*AnyOfmicrosoftGraphDateTimeTimeZone, bool) {
	if o == nil || o.EndTime == nil {
		return nil, false
	}
	return &o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *InlineObject223) HasEndTime() bool {
	if o != nil && o.EndTime != nil {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given AnyOfmicrosoftGraphDateTimeTimeZone and assigns it to the EndTime field.
func (o *InlineObject223) SetEndTime(v AnyOfmicrosoftGraphDateTimeTimeZone) {
	o.EndTime = v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject223) GetStartTime() AnyOfmicrosoftGraphDateTimeTimeZone {
	if o == nil  {
		var ret AnyOfmicrosoftGraphDateTimeTimeZone
		return ret
	}
	return o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject223) GetStartTimeOk() (*AnyOfmicrosoftGraphDateTimeTimeZone, bool) {
	if o == nil || o.StartTime == nil {
		return nil, false
	}
	return &o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *InlineObject223) HasStartTime() bool {
	if o != nil && o.StartTime != nil {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given AnyOfmicrosoftGraphDateTimeTimeZone and assigns it to the StartTime field.
func (o *InlineObject223) SetStartTime(v AnyOfmicrosoftGraphDateTimeTimeZone) {
	o.StartTime = v
}

// GetAvailabilityViewInterval returns the AvailabilityViewInterval field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject223) GetAvailabilityViewInterval() int32 {
	if o == nil || o.AvailabilityViewInterval.Get() == nil {
		var ret int32
		return ret
	}
	return *o.AvailabilityViewInterval.Get()
}

// GetAvailabilityViewIntervalOk returns a tuple with the AvailabilityViewInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject223) GetAvailabilityViewIntervalOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AvailabilityViewInterval.Get(), o.AvailabilityViewInterval.IsSet()
}

// HasAvailabilityViewInterval returns a boolean if a field has been set.
func (o *InlineObject223) HasAvailabilityViewInterval() bool {
	if o != nil && o.AvailabilityViewInterval.IsSet() {
		return true
	}

	return false
}

// SetAvailabilityViewInterval gets a reference to the given NullableInt32 and assigns it to the AvailabilityViewInterval field.
func (o *InlineObject223) SetAvailabilityViewInterval(v int32) {
	o.AvailabilityViewInterval.Set(&v)
}
// SetAvailabilityViewIntervalNil sets the value for AvailabilityViewInterval to be an explicit nil
func (o *InlineObject223) SetAvailabilityViewIntervalNil() {
	o.AvailabilityViewInterval.Set(nil)
}

// UnsetAvailabilityViewInterval ensures that no value is present for AvailabilityViewInterval, not even an explicit nil
func (o *InlineObject223) UnsetAvailabilityViewInterval() {
	o.AvailabilityViewInterval.Unset()
}

func (o InlineObject223) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Schedules != nil {
		toSerialize["Schedules"] = o.Schedules
	}
	if o.EndTime != nil {
		toSerialize["EndTime"] = o.EndTime
	}
	if o.StartTime != nil {
		toSerialize["StartTime"] = o.StartTime
	}
	if o.AvailabilityViewInterval.IsSet() {
		toSerialize["AvailabilityViewInterval"] = o.AvailabilityViewInterval.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject223 struct {
	value *InlineObject223
	isSet bool
}

func (v NullableInlineObject223) Get() *InlineObject223 {
	return v.value
}

func (v *NullableInlineObject223) Set(val *InlineObject223) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject223) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject223) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject223(val *InlineObject223) *NullableInlineObject223 {
	return &NullableInlineObject223{value: val, isSet: true}
}

func (v NullableInlineObject223) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject223) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


