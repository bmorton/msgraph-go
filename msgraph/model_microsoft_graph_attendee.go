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

// MicrosoftGraphAttendee struct for MicrosoftGraphAttendee
type MicrosoftGraphAttendee struct {
	// The recipient's email address.
	EmailAddress AnyOfmicrosoftGraphEmailAddress `json:"emailAddress,omitempty"`
	// The type of attendee. The possible values are: required, optional, resource. Currently if the attendee is a person, findMeetingTimes always considers the person is of the Required type.
	Type AnyOfmicrosoftGraphAttendeeType `json:"type,omitempty"`
	// An alternate date/time proposed by the attendee for a meeting request to start and end. If the attendee hasn't proposed another time, then this property is not included in a response of a GET event.
	ProposedNewTime AnyOfmicrosoftGraphTimeSlot `json:"proposedNewTime,omitempty"`
	// The attendee's response (none, accepted, declined, etc.) for the event and date-time that the response was sent.
	Status AnyOfmicrosoftGraphResponseStatus `json:"status,omitempty"`
}

// NewMicrosoftGraphAttendee instantiates a new MicrosoftGraphAttendee object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphAttendee() *MicrosoftGraphAttendee {
	this := MicrosoftGraphAttendee{}
	return &this
}

// NewMicrosoftGraphAttendeeWithDefaults instantiates a new MicrosoftGraphAttendee object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphAttendeeWithDefaults() *MicrosoftGraphAttendee {
	this := MicrosoftGraphAttendee{}
	return &this
}

// GetEmailAddress returns the EmailAddress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphAttendee) GetEmailAddress() AnyOfmicrosoftGraphEmailAddress {
	if o == nil  {
		var ret AnyOfmicrosoftGraphEmailAddress
		return ret
	}
	return o.EmailAddress
}

// GetEmailAddressOk returns a tuple with the EmailAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphAttendee) GetEmailAddressOk() (*AnyOfmicrosoftGraphEmailAddress, bool) {
	if o == nil || o.EmailAddress == nil {
		return nil, false
	}
	return &o.EmailAddress, true
}

// HasEmailAddress returns a boolean if a field has been set.
func (o *MicrosoftGraphAttendee) HasEmailAddress() bool {
	if o != nil && o.EmailAddress != nil {
		return true
	}

	return false
}

// SetEmailAddress gets a reference to the given AnyOfmicrosoftGraphEmailAddress and assigns it to the EmailAddress field.
func (o *MicrosoftGraphAttendee) SetEmailAddress(v AnyOfmicrosoftGraphEmailAddress) {
	o.EmailAddress = v
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphAttendee) GetType() AnyOfmicrosoftGraphAttendeeType {
	if o == nil  {
		var ret AnyOfmicrosoftGraphAttendeeType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphAttendee) GetTypeOk() (*AnyOfmicrosoftGraphAttendeeType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return &o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *MicrosoftGraphAttendee) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given AnyOfmicrosoftGraphAttendeeType and assigns it to the Type field.
func (o *MicrosoftGraphAttendee) SetType(v AnyOfmicrosoftGraphAttendeeType) {
	o.Type = v
}

// GetProposedNewTime returns the ProposedNewTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphAttendee) GetProposedNewTime() AnyOfmicrosoftGraphTimeSlot {
	if o == nil  {
		var ret AnyOfmicrosoftGraphTimeSlot
		return ret
	}
	return o.ProposedNewTime
}

// GetProposedNewTimeOk returns a tuple with the ProposedNewTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphAttendee) GetProposedNewTimeOk() (*AnyOfmicrosoftGraphTimeSlot, bool) {
	if o == nil || o.ProposedNewTime == nil {
		return nil, false
	}
	return &o.ProposedNewTime, true
}

// HasProposedNewTime returns a boolean if a field has been set.
func (o *MicrosoftGraphAttendee) HasProposedNewTime() bool {
	if o != nil && o.ProposedNewTime != nil {
		return true
	}

	return false
}

// SetProposedNewTime gets a reference to the given AnyOfmicrosoftGraphTimeSlot and assigns it to the ProposedNewTime field.
func (o *MicrosoftGraphAttendee) SetProposedNewTime(v AnyOfmicrosoftGraphTimeSlot) {
	o.ProposedNewTime = v
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphAttendee) GetStatus() AnyOfmicrosoftGraphResponseStatus {
	if o == nil  {
		var ret AnyOfmicrosoftGraphResponseStatus
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphAttendee) GetStatusOk() (*AnyOfmicrosoftGraphResponseStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return &o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *MicrosoftGraphAttendee) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given AnyOfmicrosoftGraphResponseStatus and assigns it to the Status field.
func (o *MicrosoftGraphAttendee) SetStatus(v AnyOfmicrosoftGraphResponseStatus) {
	o.Status = v
}

func (o MicrosoftGraphAttendee) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.EmailAddress != nil {
		toSerialize["emailAddress"] = o.EmailAddress
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.ProposedNewTime != nil {
		toSerialize["proposedNewTime"] = o.ProposedNewTime
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphAttendee struct {
	value *MicrosoftGraphAttendee
	isSet bool
}

func (v NullableMicrosoftGraphAttendee) Get() *MicrosoftGraphAttendee {
	return v.value
}

func (v *NullableMicrosoftGraphAttendee) Set(val *MicrosoftGraphAttendee) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphAttendee) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphAttendee) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphAttendee(val *MicrosoftGraphAttendee) *NullableMicrosoftGraphAttendee {
	return &NullableMicrosoftGraphAttendee{value: val, isSet: true}
}

func (v NullableMicrosoftGraphAttendee) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphAttendee) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


