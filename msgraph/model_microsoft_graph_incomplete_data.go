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

// MicrosoftGraphIncompleteData struct for MicrosoftGraphIncompleteData
type MicrosoftGraphIncompleteData struct {
	// The service does not have source data before the specified time.
	MissingDataBeforeDateTime NullableTime `json:"missingDataBeforeDateTime,omitempty"`
	// Some data was not recorded due to excessive activity.
	WasThrottled NullableBool `json:"wasThrottled,omitempty"`
}

// NewMicrosoftGraphIncompleteData instantiates a new MicrosoftGraphIncompleteData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphIncompleteData() *MicrosoftGraphIncompleteData {
	this := MicrosoftGraphIncompleteData{}
	return &this
}

// NewMicrosoftGraphIncompleteDataWithDefaults instantiates a new MicrosoftGraphIncompleteData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphIncompleteDataWithDefaults() *MicrosoftGraphIncompleteData {
	this := MicrosoftGraphIncompleteData{}
	return &this
}

// GetMissingDataBeforeDateTime returns the MissingDataBeforeDateTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphIncompleteData) GetMissingDataBeforeDateTime() time.Time {
	if o == nil || o.MissingDataBeforeDateTime.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.MissingDataBeforeDateTime.Get()
}

// GetMissingDataBeforeDateTimeOk returns a tuple with the MissingDataBeforeDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphIncompleteData) GetMissingDataBeforeDateTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.MissingDataBeforeDateTime.Get(), o.MissingDataBeforeDateTime.IsSet()
}

// HasMissingDataBeforeDateTime returns a boolean if a field has been set.
func (o *MicrosoftGraphIncompleteData) HasMissingDataBeforeDateTime() bool {
	if o != nil && o.MissingDataBeforeDateTime.IsSet() {
		return true
	}

	return false
}

// SetMissingDataBeforeDateTime gets a reference to the given NullableTime and assigns it to the MissingDataBeforeDateTime field.
func (o *MicrosoftGraphIncompleteData) SetMissingDataBeforeDateTime(v time.Time) {
	o.MissingDataBeforeDateTime.Set(&v)
}
// SetMissingDataBeforeDateTimeNil sets the value for MissingDataBeforeDateTime to be an explicit nil
func (o *MicrosoftGraphIncompleteData) SetMissingDataBeforeDateTimeNil() {
	o.MissingDataBeforeDateTime.Set(nil)
}

// UnsetMissingDataBeforeDateTime ensures that no value is present for MissingDataBeforeDateTime, not even an explicit nil
func (o *MicrosoftGraphIncompleteData) UnsetMissingDataBeforeDateTime() {
	o.MissingDataBeforeDateTime.Unset()
}

// GetWasThrottled returns the WasThrottled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphIncompleteData) GetWasThrottled() bool {
	if o == nil || o.WasThrottled.Get() == nil {
		var ret bool
		return ret
	}
	return *o.WasThrottled.Get()
}

// GetWasThrottledOk returns a tuple with the WasThrottled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphIncompleteData) GetWasThrottledOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.WasThrottled.Get(), o.WasThrottled.IsSet()
}

// HasWasThrottled returns a boolean if a field has been set.
func (o *MicrosoftGraphIncompleteData) HasWasThrottled() bool {
	if o != nil && o.WasThrottled.IsSet() {
		return true
	}

	return false
}

// SetWasThrottled gets a reference to the given NullableBool and assigns it to the WasThrottled field.
func (o *MicrosoftGraphIncompleteData) SetWasThrottled(v bool) {
	o.WasThrottled.Set(&v)
}
// SetWasThrottledNil sets the value for WasThrottled to be an explicit nil
func (o *MicrosoftGraphIncompleteData) SetWasThrottledNil() {
	o.WasThrottled.Set(nil)
}

// UnsetWasThrottled ensures that no value is present for WasThrottled, not even an explicit nil
func (o *MicrosoftGraphIncompleteData) UnsetWasThrottled() {
	o.WasThrottled.Unset()
}

func (o MicrosoftGraphIncompleteData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.MissingDataBeforeDateTime.IsSet() {
		toSerialize["missingDataBeforeDateTime"] = o.MissingDataBeforeDateTime.Get()
	}
	if o.WasThrottled.IsSet() {
		toSerialize["wasThrottled"] = o.WasThrottled.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphIncompleteData struct {
	value *MicrosoftGraphIncompleteData
	isSet bool
}

func (v NullableMicrosoftGraphIncompleteData) Get() *MicrosoftGraphIncompleteData {
	return v.value
}

func (v *NullableMicrosoftGraphIncompleteData) Set(val *MicrosoftGraphIncompleteData) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphIncompleteData) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphIncompleteData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphIncompleteData(val *MicrosoftGraphIncompleteData) *NullableMicrosoftGraphIncompleteData {
	return &NullableMicrosoftGraphIncompleteData{value: val, isSet: true}
}

func (v NullableMicrosoftGraphIncompleteData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphIncompleteData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

