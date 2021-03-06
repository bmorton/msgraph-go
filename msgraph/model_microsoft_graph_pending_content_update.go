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

// MicrosoftGraphPendingContentUpdate struct for MicrosoftGraphPendingContentUpdate
type MicrosoftGraphPendingContentUpdate struct {
	// Date and time the pending binary operation was queued in UTC time. Read-only.
	QueuedDateTime NullableTime `json:"queuedDateTime,omitempty"`
}

// NewMicrosoftGraphPendingContentUpdate instantiates a new MicrosoftGraphPendingContentUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphPendingContentUpdate() *MicrosoftGraphPendingContentUpdate {
	this := MicrosoftGraphPendingContentUpdate{}
	return &this
}

// NewMicrosoftGraphPendingContentUpdateWithDefaults instantiates a new MicrosoftGraphPendingContentUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphPendingContentUpdateWithDefaults() *MicrosoftGraphPendingContentUpdate {
	this := MicrosoftGraphPendingContentUpdate{}
	return &this
}

// GetQueuedDateTime returns the QueuedDateTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphPendingContentUpdate) GetQueuedDateTime() time.Time {
	if o == nil || o.QueuedDateTime.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.QueuedDateTime.Get()
}

// GetQueuedDateTimeOk returns a tuple with the QueuedDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphPendingContentUpdate) GetQueuedDateTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.QueuedDateTime.Get(), o.QueuedDateTime.IsSet()
}

// HasQueuedDateTime returns a boolean if a field has been set.
func (o *MicrosoftGraphPendingContentUpdate) HasQueuedDateTime() bool {
	if o != nil && o.QueuedDateTime.IsSet() {
		return true
	}

	return false
}

// SetQueuedDateTime gets a reference to the given NullableTime and assigns it to the QueuedDateTime field.
func (o *MicrosoftGraphPendingContentUpdate) SetQueuedDateTime(v time.Time) {
	o.QueuedDateTime.Set(&v)
}
// SetQueuedDateTimeNil sets the value for QueuedDateTime to be an explicit nil
func (o *MicrosoftGraphPendingContentUpdate) SetQueuedDateTimeNil() {
	o.QueuedDateTime.Set(nil)
}

// UnsetQueuedDateTime ensures that no value is present for QueuedDateTime, not even an explicit nil
func (o *MicrosoftGraphPendingContentUpdate) UnsetQueuedDateTime() {
	o.QueuedDateTime.Unset()
}

func (o MicrosoftGraphPendingContentUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.QueuedDateTime.IsSet() {
		toSerialize["queuedDateTime"] = o.QueuedDateTime.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphPendingContentUpdate struct {
	value *MicrosoftGraphPendingContentUpdate
	isSet bool
}

func (v NullableMicrosoftGraphPendingContentUpdate) Get() *MicrosoftGraphPendingContentUpdate {
	return v.value
}

func (v *NullableMicrosoftGraphPendingContentUpdate) Set(val *MicrosoftGraphPendingContentUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphPendingContentUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphPendingContentUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphPendingContentUpdate(val *MicrosoftGraphPendingContentUpdate) *NullableMicrosoftGraphPendingContentUpdate {
	return &NullableMicrosoftGraphPendingContentUpdate{value: val, isSet: true}
}

func (v NullableMicrosoftGraphPendingContentUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphPendingContentUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


