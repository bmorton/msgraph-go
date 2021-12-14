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

// MicrosoftGraphRecurrenceRange struct for MicrosoftGraphRecurrenceRange
type MicrosoftGraphRecurrenceRange struct {
	// The date to stop applying the recurrence pattern. Depending on the recurrence pattern of the event, the last occurrence of the meeting may not be this date. Required if type is endDate.
	EndDate NullableString `json:"endDate,omitempty"`
	// The number of times to repeat the event. Required and must be positive if type is numbered.
	NumberOfOccurrences *int32 `json:"numberOfOccurrences,omitempty"`
	// Time zone for the startDate and endDate properties. Optional. If not specified, the time zone of the event is used.
	RecurrenceTimeZone NullableString `json:"recurrenceTimeZone,omitempty"`
	// The date to start applying the recurrence pattern. The first occurrence of the meeting may be this date or later, depending on the recurrence pattern of the event. Must be the same value as the start property of the recurring event. Required.
	StartDate NullableString `json:"startDate,omitempty"`
	// The recurrence range. The possible values are: endDate, noEnd, numbered. Required.
	Type AnyOfmicrosoftGraphRecurrenceRangeType `json:"type,omitempty"`
}

// NewMicrosoftGraphRecurrenceRange instantiates a new MicrosoftGraphRecurrenceRange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphRecurrenceRange() *MicrosoftGraphRecurrenceRange {
	this := MicrosoftGraphRecurrenceRange{}
	return &this
}

// NewMicrosoftGraphRecurrenceRangeWithDefaults instantiates a new MicrosoftGraphRecurrenceRange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphRecurrenceRangeWithDefaults() *MicrosoftGraphRecurrenceRange {
	this := MicrosoftGraphRecurrenceRange{}
	return &this
}

// GetEndDate returns the EndDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphRecurrenceRange) GetEndDate() string {
	if o == nil || o.EndDate.Get() == nil {
		var ret string
		return ret
	}
	return *o.EndDate.Get()
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphRecurrenceRange) GetEndDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.EndDate.Get(), o.EndDate.IsSet()
}

// HasEndDate returns a boolean if a field has been set.
func (o *MicrosoftGraphRecurrenceRange) HasEndDate() bool {
	if o != nil && o.EndDate.IsSet() {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given NullableString and assigns it to the EndDate field.
func (o *MicrosoftGraphRecurrenceRange) SetEndDate(v string) {
	o.EndDate.Set(&v)
}
// SetEndDateNil sets the value for EndDate to be an explicit nil
func (o *MicrosoftGraphRecurrenceRange) SetEndDateNil() {
	o.EndDate.Set(nil)
}

// UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil
func (o *MicrosoftGraphRecurrenceRange) UnsetEndDate() {
	o.EndDate.Unset()
}

// GetNumberOfOccurrences returns the NumberOfOccurrences field value if set, zero value otherwise.
func (o *MicrosoftGraphRecurrenceRange) GetNumberOfOccurrences() int32 {
	if o == nil || o.NumberOfOccurrences == nil {
		var ret int32
		return ret
	}
	return *o.NumberOfOccurrences
}

// GetNumberOfOccurrencesOk returns a tuple with the NumberOfOccurrences field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphRecurrenceRange) GetNumberOfOccurrencesOk() (*int32, bool) {
	if o == nil || o.NumberOfOccurrences == nil {
		return nil, false
	}
	return o.NumberOfOccurrences, true
}

// HasNumberOfOccurrences returns a boolean if a field has been set.
func (o *MicrosoftGraphRecurrenceRange) HasNumberOfOccurrences() bool {
	if o != nil && o.NumberOfOccurrences != nil {
		return true
	}

	return false
}

// SetNumberOfOccurrences gets a reference to the given int32 and assigns it to the NumberOfOccurrences field.
func (o *MicrosoftGraphRecurrenceRange) SetNumberOfOccurrences(v int32) {
	o.NumberOfOccurrences = &v
}

// GetRecurrenceTimeZone returns the RecurrenceTimeZone field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphRecurrenceRange) GetRecurrenceTimeZone() string {
	if o == nil || o.RecurrenceTimeZone.Get() == nil {
		var ret string
		return ret
	}
	return *o.RecurrenceTimeZone.Get()
}

// GetRecurrenceTimeZoneOk returns a tuple with the RecurrenceTimeZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphRecurrenceRange) GetRecurrenceTimeZoneOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RecurrenceTimeZone.Get(), o.RecurrenceTimeZone.IsSet()
}

// HasRecurrenceTimeZone returns a boolean if a field has been set.
func (o *MicrosoftGraphRecurrenceRange) HasRecurrenceTimeZone() bool {
	if o != nil && o.RecurrenceTimeZone.IsSet() {
		return true
	}

	return false
}

// SetRecurrenceTimeZone gets a reference to the given NullableString and assigns it to the RecurrenceTimeZone field.
func (o *MicrosoftGraphRecurrenceRange) SetRecurrenceTimeZone(v string) {
	o.RecurrenceTimeZone.Set(&v)
}
// SetRecurrenceTimeZoneNil sets the value for RecurrenceTimeZone to be an explicit nil
func (o *MicrosoftGraphRecurrenceRange) SetRecurrenceTimeZoneNil() {
	o.RecurrenceTimeZone.Set(nil)
}

// UnsetRecurrenceTimeZone ensures that no value is present for RecurrenceTimeZone, not even an explicit nil
func (o *MicrosoftGraphRecurrenceRange) UnsetRecurrenceTimeZone() {
	o.RecurrenceTimeZone.Unset()
}

// GetStartDate returns the StartDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphRecurrenceRange) GetStartDate() string {
	if o == nil || o.StartDate.Get() == nil {
		var ret string
		return ret
	}
	return *o.StartDate.Get()
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphRecurrenceRange) GetStartDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.StartDate.Get(), o.StartDate.IsSet()
}

// HasStartDate returns a boolean if a field has been set.
func (o *MicrosoftGraphRecurrenceRange) HasStartDate() bool {
	if o != nil && o.StartDate.IsSet() {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given NullableString and assigns it to the StartDate field.
func (o *MicrosoftGraphRecurrenceRange) SetStartDate(v string) {
	o.StartDate.Set(&v)
}
// SetStartDateNil sets the value for StartDate to be an explicit nil
func (o *MicrosoftGraphRecurrenceRange) SetStartDateNil() {
	o.StartDate.Set(nil)
}

// UnsetStartDate ensures that no value is present for StartDate, not even an explicit nil
func (o *MicrosoftGraphRecurrenceRange) UnsetStartDate() {
	o.StartDate.Unset()
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphRecurrenceRange) GetType() AnyOfmicrosoftGraphRecurrenceRangeType {
	if o == nil  {
		var ret AnyOfmicrosoftGraphRecurrenceRangeType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphRecurrenceRange) GetTypeOk() (*AnyOfmicrosoftGraphRecurrenceRangeType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return &o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *MicrosoftGraphRecurrenceRange) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given AnyOfmicrosoftGraphRecurrenceRangeType and assigns it to the Type field.
func (o *MicrosoftGraphRecurrenceRange) SetType(v AnyOfmicrosoftGraphRecurrenceRangeType) {
	o.Type = v
}

func (o MicrosoftGraphRecurrenceRange) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.EndDate.IsSet() {
		toSerialize["endDate"] = o.EndDate.Get()
	}
	if o.NumberOfOccurrences != nil {
		toSerialize["numberOfOccurrences"] = o.NumberOfOccurrences
	}
	if o.RecurrenceTimeZone.IsSet() {
		toSerialize["recurrenceTimeZone"] = o.RecurrenceTimeZone.Get()
	}
	if o.StartDate.IsSet() {
		toSerialize["startDate"] = o.StartDate.Get()
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphRecurrenceRange struct {
	value *MicrosoftGraphRecurrenceRange
	isSet bool
}

func (v NullableMicrosoftGraphRecurrenceRange) Get() *MicrosoftGraphRecurrenceRange {
	return v.value
}

func (v *NullableMicrosoftGraphRecurrenceRange) Set(val *MicrosoftGraphRecurrenceRange) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphRecurrenceRange) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphRecurrenceRange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphRecurrenceRange(val *MicrosoftGraphRecurrenceRange) *NullableMicrosoftGraphRecurrenceRange {
	return &NullableMicrosoftGraphRecurrenceRange{value: val, isSet: true}
}

func (v NullableMicrosoftGraphRecurrenceRange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphRecurrenceRange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

