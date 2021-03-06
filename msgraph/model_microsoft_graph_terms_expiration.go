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

// MicrosoftGraphTermsExpiration struct for MicrosoftGraphTermsExpiration
type MicrosoftGraphTermsExpiration struct {
	// Represents the frequency at which the terms will expire, after its first expiration as set in startDateTime. The value is represented in ISO 8601 format for durations. For example, PT1M represents a time period of 1 month.
	Frequency NullableString `json:"frequency,omitempty"`
	// The DateTime when the agreement is set to expire for all users. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	StartDateTime NullableTime `json:"startDateTime,omitempty"`
}

// NewMicrosoftGraphTermsExpiration instantiates a new MicrosoftGraphTermsExpiration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphTermsExpiration() *MicrosoftGraphTermsExpiration {
	this := MicrosoftGraphTermsExpiration{}
	return &this
}

// NewMicrosoftGraphTermsExpirationWithDefaults instantiates a new MicrosoftGraphTermsExpiration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphTermsExpirationWithDefaults() *MicrosoftGraphTermsExpiration {
	this := MicrosoftGraphTermsExpiration{}
	return &this
}

// GetFrequency returns the Frequency field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphTermsExpiration) GetFrequency() string {
	if o == nil || o.Frequency.Get() == nil {
		var ret string
		return ret
	}
	return *o.Frequency.Get()
}

// GetFrequencyOk returns a tuple with the Frequency field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphTermsExpiration) GetFrequencyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Frequency.Get(), o.Frequency.IsSet()
}

// HasFrequency returns a boolean if a field has been set.
func (o *MicrosoftGraphTermsExpiration) HasFrequency() bool {
	if o != nil && o.Frequency.IsSet() {
		return true
	}

	return false
}

// SetFrequency gets a reference to the given NullableString and assigns it to the Frequency field.
func (o *MicrosoftGraphTermsExpiration) SetFrequency(v string) {
	o.Frequency.Set(&v)
}
// SetFrequencyNil sets the value for Frequency to be an explicit nil
func (o *MicrosoftGraphTermsExpiration) SetFrequencyNil() {
	o.Frequency.Set(nil)
}

// UnsetFrequency ensures that no value is present for Frequency, not even an explicit nil
func (o *MicrosoftGraphTermsExpiration) UnsetFrequency() {
	o.Frequency.Unset()
}

// GetStartDateTime returns the StartDateTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphTermsExpiration) GetStartDateTime() time.Time {
	if o == nil || o.StartDateTime.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.StartDateTime.Get()
}

// GetStartDateTimeOk returns a tuple with the StartDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphTermsExpiration) GetStartDateTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.StartDateTime.Get(), o.StartDateTime.IsSet()
}

// HasStartDateTime returns a boolean if a field has been set.
func (o *MicrosoftGraphTermsExpiration) HasStartDateTime() bool {
	if o != nil && o.StartDateTime.IsSet() {
		return true
	}

	return false
}

// SetStartDateTime gets a reference to the given NullableTime and assigns it to the StartDateTime field.
func (o *MicrosoftGraphTermsExpiration) SetStartDateTime(v time.Time) {
	o.StartDateTime.Set(&v)
}
// SetStartDateTimeNil sets the value for StartDateTime to be an explicit nil
func (o *MicrosoftGraphTermsExpiration) SetStartDateTimeNil() {
	o.StartDateTime.Set(nil)
}

// UnsetStartDateTime ensures that no value is present for StartDateTime, not even an explicit nil
func (o *MicrosoftGraphTermsExpiration) UnsetStartDateTime() {
	o.StartDateTime.Unset()
}

func (o MicrosoftGraphTermsExpiration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Frequency.IsSet() {
		toSerialize["frequency"] = o.Frequency.Get()
	}
	if o.StartDateTime.IsSet() {
		toSerialize["startDateTime"] = o.StartDateTime.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphTermsExpiration struct {
	value *MicrosoftGraphTermsExpiration
	isSet bool
}

func (v NullableMicrosoftGraphTermsExpiration) Get() *MicrosoftGraphTermsExpiration {
	return v.value
}

func (v *NullableMicrosoftGraphTermsExpiration) Set(val *MicrosoftGraphTermsExpiration) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphTermsExpiration) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphTermsExpiration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphTermsExpiration(val *MicrosoftGraphTermsExpiration) *NullableMicrosoftGraphTermsExpiration {
	return &NullableMicrosoftGraphTermsExpiration{value: val, isSet: true}
}

func (v NullableMicrosoftGraphTermsExpiration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphTermsExpiration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


