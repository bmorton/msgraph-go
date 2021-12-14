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

// InlineObject360 struct for InlineObject360
type InlineObject360 struct {
	Schedules *[]*string `json:"Schedules,omitempty"`
	EndTime AnyOfmicrosoftGraphDateTimeTimeZone `json:"EndTime,omitempty"`
	StartTime AnyOfmicrosoftGraphDateTimeTimeZone `json:"StartTime,omitempty"`
	AvailabilityViewInterval NullableInt32 `json:"AvailabilityViewInterval,omitempty"`
}

// NewInlineObject360 instantiates a new InlineObject360 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject360() *InlineObject360 {
	this := InlineObject360{}
	return &this
}

// NewInlineObject360WithDefaults instantiates a new InlineObject360 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject360WithDefaults() *InlineObject360 {
	this := InlineObject360{}
	return &this
}

// GetSchedules returns the Schedules field value if set, zero value otherwise.
func (o *InlineObject360) GetSchedules() []*string {
	if o == nil || o.Schedules == nil {
		var ret []*string
		return ret
	}
	return *o.Schedules
}

// GetSchedulesOk returns a tuple with the Schedules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject360) GetSchedulesOk() (*[]*string, bool) {
	if o == nil || o.Schedules == nil {
		return nil, false
	}
	return o.Schedules, true
}

// HasSchedules returns a boolean if a field has been set.
func (o *InlineObject360) HasSchedules() bool {
	if o != nil && o.Schedules != nil {
		return true
	}

	return false
}

// SetSchedules gets a reference to the given []*string and assigns it to the Schedules field.
func (o *InlineObject360) SetSchedules(v []*string) {
	o.Schedules = &v
}

// GetEndTime returns the EndTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject360) GetEndTime() AnyOfmicrosoftGraphDateTimeTimeZone {
	if o == nil  {
		var ret AnyOfmicrosoftGraphDateTimeTimeZone
		return ret
	}
	return o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject360) GetEndTimeOk() (*AnyOfmicrosoftGraphDateTimeTimeZone, bool) {
	if o == nil || o.EndTime == nil {
		return nil, false
	}
	return &o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *InlineObject360) HasEndTime() bool {
	if o != nil && o.EndTime != nil {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given AnyOfmicrosoftGraphDateTimeTimeZone and assigns it to the EndTime field.
func (o *InlineObject360) SetEndTime(v AnyOfmicrosoftGraphDateTimeTimeZone) {
	o.EndTime = v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject360) GetStartTime() AnyOfmicrosoftGraphDateTimeTimeZone {
	if o == nil  {
		var ret AnyOfmicrosoftGraphDateTimeTimeZone
		return ret
	}
	return o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject360) GetStartTimeOk() (*AnyOfmicrosoftGraphDateTimeTimeZone, bool) {
	if o == nil || o.StartTime == nil {
		return nil, false
	}
	return &o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *InlineObject360) HasStartTime() bool {
	if o != nil && o.StartTime != nil {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given AnyOfmicrosoftGraphDateTimeTimeZone and assigns it to the StartTime field.
func (o *InlineObject360) SetStartTime(v AnyOfmicrosoftGraphDateTimeTimeZone) {
	o.StartTime = v
}

// GetAvailabilityViewInterval returns the AvailabilityViewInterval field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject360) GetAvailabilityViewInterval() int32 {
	if o == nil || o.AvailabilityViewInterval.Get() == nil {
		var ret int32
		return ret
	}
	return *o.AvailabilityViewInterval.Get()
}

// GetAvailabilityViewIntervalOk returns a tuple with the AvailabilityViewInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject360) GetAvailabilityViewIntervalOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AvailabilityViewInterval.Get(), o.AvailabilityViewInterval.IsSet()
}

// HasAvailabilityViewInterval returns a boolean if a field has been set.
func (o *InlineObject360) HasAvailabilityViewInterval() bool {
	if o != nil && o.AvailabilityViewInterval.IsSet() {
		return true
	}

	return false
}

// SetAvailabilityViewInterval gets a reference to the given NullableInt32 and assigns it to the AvailabilityViewInterval field.
func (o *InlineObject360) SetAvailabilityViewInterval(v int32) {
	o.AvailabilityViewInterval.Set(&v)
}
// SetAvailabilityViewIntervalNil sets the value for AvailabilityViewInterval to be an explicit nil
func (o *InlineObject360) SetAvailabilityViewIntervalNil() {
	o.AvailabilityViewInterval.Set(nil)
}

// UnsetAvailabilityViewInterval ensures that no value is present for AvailabilityViewInterval, not even an explicit nil
func (o *InlineObject360) UnsetAvailabilityViewInterval() {
	o.AvailabilityViewInterval.Unset()
}

func (o InlineObject360) MarshalJSON() ([]byte, error) {
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

type NullableInlineObject360 struct {
	value *InlineObject360
	isSet bool
}

func (v NullableInlineObject360) Get() *InlineObject360 {
	return v.value
}

func (v *NullableInlineObject360) Set(val *InlineObject360) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject360) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject360) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject360(val *InlineObject360) *NullableInlineObject360 {
	return &NullableInlineObject360{value: val, isSet: true}
}

func (v NullableInlineObject360) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject360) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


