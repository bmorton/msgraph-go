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

// MicrosoftGraphExpirationPattern struct for MicrosoftGraphExpirationPattern
type MicrosoftGraphExpirationPattern struct {
	// The requestor's desired duration of access represented in ISO 8601 format for durations. For example, PT3H refers to three hours.  If specified in a request, endDateTime should not be present and the type property should be set to afterDuration.
	Duration NullableString `json:"duration,omitempty"`
	// Timestamp of date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	EndDateTime NullableTime `json:"endDateTime,omitempty"`
	// The requestor's desired expiration pattern type. The possible values are: notSpecified, noExpiration, afterDateTime, afterDuration.
	Type AnyOfmicrosoftGraphExpirationPatternType `json:"type,omitempty"`
}

// NewMicrosoftGraphExpirationPattern instantiates a new MicrosoftGraphExpirationPattern object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphExpirationPattern() *MicrosoftGraphExpirationPattern {
	this := MicrosoftGraphExpirationPattern{}
	return &this
}

// NewMicrosoftGraphExpirationPatternWithDefaults instantiates a new MicrosoftGraphExpirationPattern object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphExpirationPatternWithDefaults() *MicrosoftGraphExpirationPattern {
	this := MicrosoftGraphExpirationPattern{}
	return &this
}

// GetDuration returns the Duration field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphExpirationPattern) GetDuration() string {
	if o == nil || o.Duration.Get() == nil {
		var ret string
		return ret
	}
	return *o.Duration.Get()
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphExpirationPattern) GetDurationOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Duration.Get(), o.Duration.IsSet()
}

// HasDuration returns a boolean if a field has been set.
func (o *MicrosoftGraphExpirationPattern) HasDuration() bool {
	if o != nil && o.Duration.IsSet() {
		return true
	}

	return false
}

// SetDuration gets a reference to the given NullableString and assigns it to the Duration field.
func (o *MicrosoftGraphExpirationPattern) SetDuration(v string) {
	o.Duration.Set(&v)
}
// SetDurationNil sets the value for Duration to be an explicit nil
func (o *MicrosoftGraphExpirationPattern) SetDurationNil() {
	o.Duration.Set(nil)
}

// UnsetDuration ensures that no value is present for Duration, not even an explicit nil
func (o *MicrosoftGraphExpirationPattern) UnsetDuration() {
	o.Duration.Unset()
}

// GetEndDateTime returns the EndDateTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphExpirationPattern) GetEndDateTime() time.Time {
	if o == nil || o.EndDateTime.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.EndDateTime.Get()
}

// GetEndDateTimeOk returns a tuple with the EndDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphExpirationPattern) GetEndDateTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.EndDateTime.Get(), o.EndDateTime.IsSet()
}

// HasEndDateTime returns a boolean if a field has been set.
func (o *MicrosoftGraphExpirationPattern) HasEndDateTime() bool {
	if o != nil && o.EndDateTime.IsSet() {
		return true
	}

	return false
}

// SetEndDateTime gets a reference to the given NullableTime and assigns it to the EndDateTime field.
func (o *MicrosoftGraphExpirationPattern) SetEndDateTime(v time.Time) {
	o.EndDateTime.Set(&v)
}
// SetEndDateTimeNil sets the value for EndDateTime to be an explicit nil
func (o *MicrosoftGraphExpirationPattern) SetEndDateTimeNil() {
	o.EndDateTime.Set(nil)
}

// UnsetEndDateTime ensures that no value is present for EndDateTime, not even an explicit nil
func (o *MicrosoftGraphExpirationPattern) UnsetEndDateTime() {
	o.EndDateTime.Unset()
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphExpirationPattern) GetType() AnyOfmicrosoftGraphExpirationPatternType {
	if o == nil  {
		var ret AnyOfmicrosoftGraphExpirationPatternType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphExpirationPattern) GetTypeOk() (*AnyOfmicrosoftGraphExpirationPatternType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return &o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *MicrosoftGraphExpirationPattern) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given AnyOfmicrosoftGraphExpirationPatternType and assigns it to the Type field.
func (o *MicrosoftGraphExpirationPattern) SetType(v AnyOfmicrosoftGraphExpirationPatternType) {
	o.Type = v
}

func (o MicrosoftGraphExpirationPattern) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Duration.IsSet() {
		toSerialize["duration"] = o.Duration.Get()
	}
	if o.EndDateTime.IsSet() {
		toSerialize["endDateTime"] = o.EndDateTime.Get()
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphExpirationPattern struct {
	value *MicrosoftGraphExpirationPattern
	isSet bool
}

func (v NullableMicrosoftGraphExpirationPattern) Get() *MicrosoftGraphExpirationPattern {
	return v.value
}

func (v *NullableMicrosoftGraphExpirationPattern) Set(val *MicrosoftGraphExpirationPattern) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphExpirationPattern) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphExpirationPattern) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphExpirationPattern(val *MicrosoftGraphExpirationPattern) *NullableMicrosoftGraphExpirationPattern {
	return &NullableMicrosoftGraphExpirationPattern{value: val, isSet: true}
}

func (v NullableMicrosoftGraphExpirationPattern) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphExpirationPattern) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


