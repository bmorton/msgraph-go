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

// MicrosoftGraphTimeConstraint struct for MicrosoftGraphTimeConstraint
type MicrosoftGraphTimeConstraint struct {
	// The nature of the activity, optional. The possible values are: work, personal, unrestricted, or unknown.
	ActivityDomain AnyOfmicrosoftGraphActivityDomain `json:"activityDomain,omitempty"`
	TimeSlots *[]*AnyOfmicrosoftGraphTimeSlot `json:"timeSlots,omitempty"`
}

// NewMicrosoftGraphTimeConstraint instantiates a new MicrosoftGraphTimeConstraint object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphTimeConstraint() *MicrosoftGraphTimeConstraint {
	this := MicrosoftGraphTimeConstraint{}
	return &this
}

// NewMicrosoftGraphTimeConstraintWithDefaults instantiates a new MicrosoftGraphTimeConstraint object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphTimeConstraintWithDefaults() *MicrosoftGraphTimeConstraint {
	this := MicrosoftGraphTimeConstraint{}
	return &this
}

// GetActivityDomain returns the ActivityDomain field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphTimeConstraint) GetActivityDomain() AnyOfmicrosoftGraphActivityDomain {
	if o == nil  {
		var ret AnyOfmicrosoftGraphActivityDomain
		return ret
	}
	return o.ActivityDomain
}

// GetActivityDomainOk returns a tuple with the ActivityDomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphTimeConstraint) GetActivityDomainOk() (*AnyOfmicrosoftGraphActivityDomain, bool) {
	if o == nil || o.ActivityDomain == nil {
		return nil, false
	}
	return &o.ActivityDomain, true
}

// HasActivityDomain returns a boolean if a field has been set.
func (o *MicrosoftGraphTimeConstraint) HasActivityDomain() bool {
	if o != nil && o.ActivityDomain != nil {
		return true
	}

	return false
}

// SetActivityDomain gets a reference to the given AnyOfmicrosoftGraphActivityDomain and assigns it to the ActivityDomain field.
func (o *MicrosoftGraphTimeConstraint) SetActivityDomain(v AnyOfmicrosoftGraphActivityDomain) {
	o.ActivityDomain = v
}

// GetTimeSlots returns the TimeSlots field value if set, zero value otherwise.
func (o *MicrosoftGraphTimeConstraint) GetTimeSlots() []*AnyOfmicrosoftGraphTimeSlot {
	if o == nil || o.TimeSlots == nil {
		var ret []*AnyOfmicrosoftGraphTimeSlot
		return ret
	}
	return *o.TimeSlots
}

// GetTimeSlotsOk returns a tuple with the TimeSlots field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphTimeConstraint) GetTimeSlotsOk() (*[]*AnyOfmicrosoftGraphTimeSlot, bool) {
	if o == nil || o.TimeSlots == nil {
		return nil, false
	}
	return o.TimeSlots, true
}

// HasTimeSlots returns a boolean if a field has been set.
func (o *MicrosoftGraphTimeConstraint) HasTimeSlots() bool {
	if o != nil && o.TimeSlots != nil {
		return true
	}

	return false
}

// SetTimeSlots gets a reference to the given []*AnyOfmicrosoftGraphTimeSlot and assigns it to the TimeSlots field.
func (o *MicrosoftGraphTimeConstraint) SetTimeSlots(v []*AnyOfmicrosoftGraphTimeSlot) {
	o.TimeSlots = &v
}

func (o MicrosoftGraphTimeConstraint) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ActivityDomain != nil {
		toSerialize["activityDomain"] = o.ActivityDomain
	}
	if o.TimeSlots != nil {
		toSerialize["timeSlots"] = o.TimeSlots
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphTimeConstraint struct {
	value *MicrosoftGraphTimeConstraint
	isSet bool
}

func (v NullableMicrosoftGraphTimeConstraint) Get() *MicrosoftGraphTimeConstraint {
	return v.value
}

func (v *NullableMicrosoftGraphTimeConstraint) Set(val *MicrosoftGraphTimeConstraint) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphTimeConstraint) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphTimeConstraint) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphTimeConstraint(val *MicrosoftGraphTimeConstraint) *NullableMicrosoftGraphTimeConstraint {
	return &NullableMicrosoftGraphTimeConstraint{value: val, isSet: true}
}

func (v NullableMicrosoftGraphTimeConstraint) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphTimeConstraint) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


