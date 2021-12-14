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

// MicrosoftGraphAttendeeAvailability struct for MicrosoftGraphAttendeeAvailability
type MicrosoftGraphAttendeeAvailability struct {
	// The email address and type of attendee - whether it's a person or a resource, and whether required or optional if it's a person.
	Attendee AnyOfmicrosoftGraphAttendeeBase `json:"attendee,omitempty"`
	// The availability status of the attendee. The possible values are: free, tentative, busy, oof, workingElsewhere, unknown.
	Availability AnyOfmicrosoftGraphFreeBusyStatus `json:"availability,omitempty"`
}

// NewMicrosoftGraphAttendeeAvailability instantiates a new MicrosoftGraphAttendeeAvailability object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphAttendeeAvailability() *MicrosoftGraphAttendeeAvailability {
	this := MicrosoftGraphAttendeeAvailability{}
	return &this
}

// NewMicrosoftGraphAttendeeAvailabilityWithDefaults instantiates a new MicrosoftGraphAttendeeAvailability object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphAttendeeAvailabilityWithDefaults() *MicrosoftGraphAttendeeAvailability {
	this := MicrosoftGraphAttendeeAvailability{}
	return &this
}

// GetAttendee returns the Attendee field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphAttendeeAvailability) GetAttendee() AnyOfmicrosoftGraphAttendeeBase {
	if o == nil  {
		var ret AnyOfmicrosoftGraphAttendeeBase
		return ret
	}
	return o.Attendee
}

// GetAttendeeOk returns a tuple with the Attendee field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphAttendeeAvailability) GetAttendeeOk() (*AnyOfmicrosoftGraphAttendeeBase, bool) {
	if o == nil || o.Attendee == nil {
		return nil, false
	}
	return &o.Attendee, true
}

// HasAttendee returns a boolean if a field has been set.
func (o *MicrosoftGraphAttendeeAvailability) HasAttendee() bool {
	if o != nil && o.Attendee != nil {
		return true
	}

	return false
}

// SetAttendee gets a reference to the given AnyOfmicrosoftGraphAttendeeBase and assigns it to the Attendee field.
func (o *MicrosoftGraphAttendeeAvailability) SetAttendee(v AnyOfmicrosoftGraphAttendeeBase) {
	o.Attendee = v
}

// GetAvailability returns the Availability field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphAttendeeAvailability) GetAvailability() AnyOfmicrosoftGraphFreeBusyStatus {
	if o == nil  {
		var ret AnyOfmicrosoftGraphFreeBusyStatus
		return ret
	}
	return o.Availability
}

// GetAvailabilityOk returns a tuple with the Availability field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphAttendeeAvailability) GetAvailabilityOk() (*AnyOfmicrosoftGraphFreeBusyStatus, bool) {
	if o == nil || o.Availability == nil {
		return nil, false
	}
	return &o.Availability, true
}

// HasAvailability returns a boolean if a field has been set.
func (o *MicrosoftGraphAttendeeAvailability) HasAvailability() bool {
	if o != nil && o.Availability != nil {
		return true
	}

	return false
}

// SetAvailability gets a reference to the given AnyOfmicrosoftGraphFreeBusyStatus and assigns it to the Availability field.
func (o *MicrosoftGraphAttendeeAvailability) SetAvailability(v AnyOfmicrosoftGraphFreeBusyStatus) {
	o.Availability = v
}

func (o MicrosoftGraphAttendeeAvailability) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Attendee != nil {
		toSerialize["attendee"] = o.Attendee
	}
	if o.Availability != nil {
		toSerialize["availability"] = o.Availability
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphAttendeeAvailability struct {
	value *MicrosoftGraphAttendeeAvailability
	isSet bool
}

func (v NullableMicrosoftGraphAttendeeAvailability) Get() *MicrosoftGraphAttendeeAvailability {
	return v.value
}

func (v *NullableMicrosoftGraphAttendeeAvailability) Set(val *MicrosoftGraphAttendeeAvailability) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphAttendeeAvailability) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphAttendeeAvailability) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphAttendeeAvailability(val *MicrosoftGraphAttendeeAvailability) *NullableMicrosoftGraphAttendeeAvailability {
	return &NullableMicrosoftGraphAttendeeAvailability{value: val, isSet: true}
}

func (v NullableMicrosoftGraphAttendeeAvailability) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphAttendeeAvailability) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


