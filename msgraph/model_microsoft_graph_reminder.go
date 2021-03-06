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

// MicrosoftGraphReminder struct for MicrosoftGraphReminder
type MicrosoftGraphReminder struct {
	// Identifies the version of the reminder. Every time the reminder is changed, changeKey changes as well. This allows Exchange to apply changes to the correct version of the object.
	ChangeKey NullableString `json:"changeKey,omitempty"`
	// The date, time and time zone that the event ends.
	EventEndTime AnyOfmicrosoftGraphDateTimeTimeZone `json:"eventEndTime,omitempty"`
	// The unique ID of the event. Read only.
	EventId NullableString `json:"eventId,omitempty"`
	// The location of the event.
	EventLocation AnyOfmicrosoftGraphLocation `json:"eventLocation,omitempty"`
	// The date, time, and time zone that the event starts.
	EventStartTime AnyOfmicrosoftGraphDateTimeTimeZone `json:"eventStartTime,omitempty"`
	// The text of the event's subject line.
	EventSubject NullableString `json:"eventSubject,omitempty"`
	// The URL to open the event in Outlook on the web.The event will open in the browser if you are logged in to your mailbox via Outlook on the web. You will be prompted to login if you are not already logged in with the browser.This URL cannot be accessed from within an iFrame.
	EventWebLink NullableString `json:"eventWebLink,omitempty"`
	// The date, time, and time zone that the reminder is set to occur.
	ReminderFireTime AnyOfmicrosoftGraphDateTimeTimeZone `json:"reminderFireTime,omitempty"`
}

// NewMicrosoftGraphReminder instantiates a new MicrosoftGraphReminder object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphReminder() *MicrosoftGraphReminder {
	this := MicrosoftGraphReminder{}
	return &this
}

// NewMicrosoftGraphReminderWithDefaults instantiates a new MicrosoftGraphReminder object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphReminderWithDefaults() *MicrosoftGraphReminder {
	this := MicrosoftGraphReminder{}
	return &this
}

// GetChangeKey returns the ChangeKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphReminder) GetChangeKey() string {
	if o == nil || o.ChangeKey.Get() == nil {
		var ret string
		return ret
	}
	return *o.ChangeKey.Get()
}

// GetChangeKeyOk returns a tuple with the ChangeKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphReminder) GetChangeKeyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ChangeKey.Get(), o.ChangeKey.IsSet()
}

// HasChangeKey returns a boolean if a field has been set.
func (o *MicrosoftGraphReminder) HasChangeKey() bool {
	if o != nil && o.ChangeKey.IsSet() {
		return true
	}

	return false
}

// SetChangeKey gets a reference to the given NullableString and assigns it to the ChangeKey field.
func (o *MicrosoftGraphReminder) SetChangeKey(v string) {
	o.ChangeKey.Set(&v)
}
// SetChangeKeyNil sets the value for ChangeKey to be an explicit nil
func (o *MicrosoftGraphReminder) SetChangeKeyNil() {
	o.ChangeKey.Set(nil)
}

// UnsetChangeKey ensures that no value is present for ChangeKey, not even an explicit nil
func (o *MicrosoftGraphReminder) UnsetChangeKey() {
	o.ChangeKey.Unset()
}

// GetEventEndTime returns the EventEndTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphReminder) GetEventEndTime() AnyOfmicrosoftGraphDateTimeTimeZone {
	if o == nil  {
		var ret AnyOfmicrosoftGraphDateTimeTimeZone
		return ret
	}
	return o.EventEndTime
}

// GetEventEndTimeOk returns a tuple with the EventEndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphReminder) GetEventEndTimeOk() (*AnyOfmicrosoftGraphDateTimeTimeZone, bool) {
	if o == nil || o.EventEndTime == nil {
		return nil, false
	}
	return &o.EventEndTime, true
}

// HasEventEndTime returns a boolean if a field has been set.
func (o *MicrosoftGraphReminder) HasEventEndTime() bool {
	if o != nil && o.EventEndTime != nil {
		return true
	}

	return false
}

// SetEventEndTime gets a reference to the given AnyOfmicrosoftGraphDateTimeTimeZone and assigns it to the EventEndTime field.
func (o *MicrosoftGraphReminder) SetEventEndTime(v AnyOfmicrosoftGraphDateTimeTimeZone) {
	o.EventEndTime = v
}

// GetEventId returns the EventId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphReminder) GetEventId() string {
	if o == nil || o.EventId.Get() == nil {
		var ret string
		return ret
	}
	return *o.EventId.Get()
}

// GetEventIdOk returns a tuple with the EventId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphReminder) GetEventIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.EventId.Get(), o.EventId.IsSet()
}

// HasEventId returns a boolean if a field has been set.
func (o *MicrosoftGraphReminder) HasEventId() bool {
	if o != nil && o.EventId.IsSet() {
		return true
	}

	return false
}

// SetEventId gets a reference to the given NullableString and assigns it to the EventId field.
func (o *MicrosoftGraphReminder) SetEventId(v string) {
	o.EventId.Set(&v)
}
// SetEventIdNil sets the value for EventId to be an explicit nil
func (o *MicrosoftGraphReminder) SetEventIdNil() {
	o.EventId.Set(nil)
}

// UnsetEventId ensures that no value is present for EventId, not even an explicit nil
func (o *MicrosoftGraphReminder) UnsetEventId() {
	o.EventId.Unset()
}

// GetEventLocation returns the EventLocation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphReminder) GetEventLocation() AnyOfmicrosoftGraphLocation {
	if o == nil  {
		var ret AnyOfmicrosoftGraphLocation
		return ret
	}
	return o.EventLocation
}

// GetEventLocationOk returns a tuple with the EventLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphReminder) GetEventLocationOk() (*AnyOfmicrosoftGraphLocation, bool) {
	if o == nil || o.EventLocation == nil {
		return nil, false
	}
	return &o.EventLocation, true
}

// HasEventLocation returns a boolean if a field has been set.
func (o *MicrosoftGraphReminder) HasEventLocation() bool {
	if o != nil && o.EventLocation != nil {
		return true
	}

	return false
}

// SetEventLocation gets a reference to the given AnyOfmicrosoftGraphLocation and assigns it to the EventLocation field.
func (o *MicrosoftGraphReminder) SetEventLocation(v AnyOfmicrosoftGraphLocation) {
	o.EventLocation = v
}

// GetEventStartTime returns the EventStartTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphReminder) GetEventStartTime() AnyOfmicrosoftGraphDateTimeTimeZone {
	if o == nil  {
		var ret AnyOfmicrosoftGraphDateTimeTimeZone
		return ret
	}
	return o.EventStartTime
}

// GetEventStartTimeOk returns a tuple with the EventStartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphReminder) GetEventStartTimeOk() (*AnyOfmicrosoftGraphDateTimeTimeZone, bool) {
	if o == nil || o.EventStartTime == nil {
		return nil, false
	}
	return &o.EventStartTime, true
}

// HasEventStartTime returns a boolean if a field has been set.
func (o *MicrosoftGraphReminder) HasEventStartTime() bool {
	if o != nil && o.EventStartTime != nil {
		return true
	}

	return false
}

// SetEventStartTime gets a reference to the given AnyOfmicrosoftGraphDateTimeTimeZone and assigns it to the EventStartTime field.
func (o *MicrosoftGraphReminder) SetEventStartTime(v AnyOfmicrosoftGraphDateTimeTimeZone) {
	o.EventStartTime = v
}

// GetEventSubject returns the EventSubject field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphReminder) GetEventSubject() string {
	if o == nil || o.EventSubject.Get() == nil {
		var ret string
		return ret
	}
	return *o.EventSubject.Get()
}

// GetEventSubjectOk returns a tuple with the EventSubject field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphReminder) GetEventSubjectOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.EventSubject.Get(), o.EventSubject.IsSet()
}

// HasEventSubject returns a boolean if a field has been set.
func (o *MicrosoftGraphReminder) HasEventSubject() bool {
	if o != nil && o.EventSubject.IsSet() {
		return true
	}

	return false
}

// SetEventSubject gets a reference to the given NullableString and assigns it to the EventSubject field.
func (o *MicrosoftGraphReminder) SetEventSubject(v string) {
	o.EventSubject.Set(&v)
}
// SetEventSubjectNil sets the value for EventSubject to be an explicit nil
func (o *MicrosoftGraphReminder) SetEventSubjectNil() {
	o.EventSubject.Set(nil)
}

// UnsetEventSubject ensures that no value is present for EventSubject, not even an explicit nil
func (o *MicrosoftGraphReminder) UnsetEventSubject() {
	o.EventSubject.Unset()
}

// GetEventWebLink returns the EventWebLink field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphReminder) GetEventWebLink() string {
	if o == nil || o.EventWebLink.Get() == nil {
		var ret string
		return ret
	}
	return *o.EventWebLink.Get()
}

// GetEventWebLinkOk returns a tuple with the EventWebLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphReminder) GetEventWebLinkOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.EventWebLink.Get(), o.EventWebLink.IsSet()
}

// HasEventWebLink returns a boolean if a field has been set.
func (o *MicrosoftGraphReminder) HasEventWebLink() bool {
	if o != nil && o.EventWebLink.IsSet() {
		return true
	}

	return false
}

// SetEventWebLink gets a reference to the given NullableString and assigns it to the EventWebLink field.
func (o *MicrosoftGraphReminder) SetEventWebLink(v string) {
	o.EventWebLink.Set(&v)
}
// SetEventWebLinkNil sets the value for EventWebLink to be an explicit nil
func (o *MicrosoftGraphReminder) SetEventWebLinkNil() {
	o.EventWebLink.Set(nil)
}

// UnsetEventWebLink ensures that no value is present for EventWebLink, not even an explicit nil
func (o *MicrosoftGraphReminder) UnsetEventWebLink() {
	o.EventWebLink.Unset()
}

// GetReminderFireTime returns the ReminderFireTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphReminder) GetReminderFireTime() AnyOfmicrosoftGraphDateTimeTimeZone {
	if o == nil  {
		var ret AnyOfmicrosoftGraphDateTimeTimeZone
		return ret
	}
	return o.ReminderFireTime
}

// GetReminderFireTimeOk returns a tuple with the ReminderFireTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphReminder) GetReminderFireTimeOk() (*AnyOfmicrosoftGraphDateTimeTimeZone, bool) {
	if o == nil || o.ReminderFireTime == nil {
		return nil, false
	}
	return &o.ReminderFireTime, true
}

// HasReminderFireTime returns a boolean if a field has been set.
func (o *MicrosoftGraphReminder) HasReminderFireTime() bool {
	if o != nil && o.ReminderFireTime != nil {
		return true
	}

	return false
}

// SetReminderFireTime gets a reference to the given AnyOfmicrosoftGraphDateTimeTimeZone and assigns it to the ReminderFireTime field.
func (o *MicrosoftGraphReminder) SetReminderFireTime(v AnyOfmicrosoftGraphDateTimeTimeZone) {
	o.ReminderFireTime = v
}

func (o MicrosoftGraphReminder) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ChangeKey.IsSet() {
		toSerialize["changeKey"] = o.ChangeKey.Get()
	}
	if o.EventEndTime != nil {
		toSerialize["eventEndTime"] = o.EventEndTime
	}
	if o.EventId.IsSet() {
		toSerialize["eventId"] = o.EventId.Get()
	}
	if o.EventLocation != nil {
		toSerialize["eventLocation"] = o.EventLocation
	}
	if o.EventStartTime != nil {
		toSerialize["eventStartTime"] = o.EventStartTime
	}
	if o.EventSubject.IsSet() {
		toSerialize["eventSubject"] = o.EventSubject.Get()
	}
	if o.EventWebLink.IsSet() {
		toSerialize["eventWebLink"] = o.EventWebLink.Get()
	}
	if o.ReminderFireTime != nil {
		toSerialize["reminderFireTime"] = o.ReminderFireTime
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphReminder struct {
	value *MicrosoftGraphReminder
	isSet bool
}

func (v NullableMicrosoftGraphReminder) Get() *MicrosoftGraphReminder {
	return v.value
}

func (v *NullableMicrosoftGraphReminder) Set(val *MicrosoftGraphReminder) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphReminder) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphReminder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphReminder(val *MicrosoftGraphReminder) *NullableMicrosoftGraphReminder {
	return &NullableMicrosoftGraphReminder{value: val, isSet: true}
}

func (v NullableMicrosoftGraphReminder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphReminder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


