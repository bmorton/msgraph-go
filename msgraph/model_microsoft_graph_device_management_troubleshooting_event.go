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

// MicrosoftGraphDeviceManagementTroubleshootingEvent struct for MicrosoftGraphDeviceManagementTroubleshootingEvent
type MicrosoftGraphDeviceManagementTroubleshootingEvent struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
	// Id used for tracing the failure in the service.
	CorrelationId NullableString `json:"correlationId,omitempty"`
	// Time when the event occurred .
	EventDateTime *time.Time `json:"eventDateTime,omitempty"`
}

// NewMicrosoftGraphDeviceManagementTroubleshootingEvent instantiates a new MicrosoftGraphDeviceManagementTroubleshootingEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphDeviceManagementTroubleshootingEvent() *MicrosoftGraphDeviceManagementTroubleshootingEvent {
	this := MicrosoftGraphDeviceManagementTroubleshootingEvent{}
	return &this
}

// NewMicrosoftGraphDeviceManagementTroubleshootingEventWithDefaults instantiates a new MicrosoftGraphDeviceManagementTroubleshootingEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphDeviceManagementTroubleshootingEventWithDefaults() *MicrosoftGraphDeviceManagementTroubleshootingEvent {
	this := MicrosoftGraphDeviceManagementTroubleshootingEvent{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MicrosoftGraphDeviceManagementTroubleshootingEvent) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphDeviceManagementTroubleshootingEvent) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphDeviceManagementTroubleshootingEvent) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MicrosoftGraphDeviceManagementTroubleshootingEvent) SetId(v string) {
	o.Id = &v
}

// GetCorrelationId returns the CorrelationId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphDeviceManagementTroubleshootingEvent) GetCorrelationId() string {
	if o == nil || o.CorrelationId.Get() == nil {
		var ret string
		return ret
	}
	return *o.CorrelationId.Get()
}

// GetCorrelationIdOk returns a tuple with the CorrelationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphDeviceManagementTroubleshootingEvent) GetCorrelationIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CorrelationId.Get(), o.CorrelationId.IsSet()
}

// HasCorrelationId returns a boolean if a field has been set.
func (o *MicrosoftGraphDeviceManagementTroubleshootingEvent) HasCorrelationId() bool {
	if o != nil && o.CorrelationId.IsSet() {
		return true
	}

	return false
}

// SetCorrelationId gets a reference to the given NullableString and assigns it to the CorrelationId field.
func (o *MicrosoftGraphDeviceManagementTroubleshootingEvent) SetCorrelationId(v string) {
	o.CorrelationId.Set(&v)
}
// SetCorrelationIdNil sets the value for CorrelationId to be an explicit nil
func (o *MicrosoftGraphDeviceManagementTroubleshootingEvent) SetCorrelationIdNil() {
	o.CorrelationId.Set(nil)
}

// UnsetCorrelationId ensures that no value is present for CorrelationId, not even an explicit nil
func (o *MicrosoftGraphDeviceManagementTroubleshootingEvent) UnsetCorrelationId() {
	o.CorrelationId.Unset()
}

// GetEventDateTime returns the EventDateTime field value if set, zero value otherwise.
func (o *MicrosoftGraphDeviceManagementTroubleshootingEvent) GetEventDateTime() time.Time {
	if o == nil || o.EventDateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.EventDateTime
}

// GetEventDateTimeOk returns a tuple with the EventDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphDeviceManagementTroubleshootingEvent) GetEventDateTimeOk() (*time.Time, bool) {
	if o == nil || o.EventDateTime == nil {
		return nil, false
	}
	return o.EventDateTime, true
}

// HasEventDateTime returns a boolean if a field has been set.
func (o *MicrosoftGraphDeviceManagementTroubleshootingEvent) HasEventDateTime() bool {
	if o != nil && o.EventDateTime != nil {
		return true
	}

	return false
}

// SetEventDateTime gets a reference to the given time.Time and assigns it to the EventDateTime field.
func (o *MicrosoftGraphDeviceManagementTroubleshootingEvent) SetEventDateTime(v time.Time) {
	o.EventDateTime = &v
}

func (o MicrosoftGraphDeviceManagementTroubleshootingEvent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.CorrelationId.IsSet() {
		toSerialize["correlationId"] = o.CorrelationId.Get()
	}
	if o.EventDateTime != nil {
		toSerialize["eventDateTime"] = o.EventDateTime
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphDeviceManagementTroubleshootingEvent struct {
	value *MicrosoftGraphDeviceManagementTroubleshootingEvent
	isSet bool
}

func (v NullableMicrosoftGraphDeviceManagementTroubleshootingEvent) Get() *MicrosoftGraphDeviceManagementTroubleshootingEvent {
	return v.value
}

func (v *NullableMicrosoftGraphDeviceManagementTroubleshootingEvent) Set(val *MicrosoftGraphDeviceManagementTroubleshootingEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphDeviceManagementTroubleshootingEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphDeviceManagementTroubleshootingEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphDeviceManagementTroubleshootingEvent(val *MicrosoftGraphDeviceManagementTroubleshootingEvent) *NullableMicrosoftGraphDeviceManagementTroubleshootingEvent {
	return &NullableMicrosoftGraphDeviceManagementTroubleshootingEvent{value: val, isSet: true}
}

func (v NullableMicrosoftGraphDeviceManagementTroubleshootingEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphDeviceManagementTroubleshootingEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


