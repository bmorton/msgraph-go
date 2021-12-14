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

// MicrosoftGraphTimeSlot struct for MicrosoftGraphTimeSlot
type MicrosoftGraphTimeSlot struct {
	End *MicrosoftGraphDateTimeTimeZone `json:"end,omitempty"`
	Start *MicrosoftGraphDateTimeTimeZone `json:"start,omitempty"`
}

// NewMicrosoftGraphTimeSlot instantiates a new MicrosoftGraphTimeSlot object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphTimeSlot() *MicrosoftGraphTimeSlot {
	this := MicrosoftGraphTimeSlot{}
	return &this
}

// NewMicrosoftGraphTimeSlotWithDefaults instantiates a new MicrosoftGraphTimeSlot object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphTimeSlotWithDefaults() *MicrosoftGraphTimeSlot {
	this := MicrosoftGraphTimeSlot{}
	return &this
}

// GetEnd returns the End field value if set, zero value otherwise.
func (o *MicrosoftGraphTimeSlot) GetEnd() MicrosoftGraphDateTimeTimeZone {
	if o == nil || o.End == nil {
		var ret MicrosoftGraphDateTimeTimeZone
		return ret
	}
	return *o.End
}

// GetEndOk returns a tuple with the End field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphTimeSlot) GetEndOk() (*MicrosoftGraphDateTimeTimeZone, bool) {
	if o == nil || o.End == nil {
		return nil, false
	}
	return o.End, true
}

// HasEnd returns a boolean if a field has been set.
func (o *MicrosoftGraphTimeSlot) HasEnd() bool {
	if o != nil && o.End != nil {
		return true
	}

	return false
}

// SetEnd gets a reference to the given MicrosoftGraphDateTimeTimeZone and assigns it to the End field.
func (o *MicrosoftGraphTimeSlot) SetEnd(v MicrosoftGraphDateTimeTimeZone) {
	o.End = &v
}

// GetStart returns the Start field value if set, zero value otherwise.
func (o *MicrosoftGraphTimeSlot) GetStart() MicrosoftGraphDateTimeTimeZone {
	if o == nil || o.Start == nil {
		var ret MicrosoftGraphDateTimeTimeZone
		return ret
	}
	return *o.Start
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphTimeSlot) GetStartOk() (*MicrosoftGraphDateTimeTimeZone, bool) {
	if o == nil || o.Start == nil {
		return nil, false
	}
	return o.Start, true
}

// HasStart returns a boolean if a field has been set.
func (o *MicrosoftGraphTimeSlot) HasStart() bool {
	if o != nil && o.Start != nil {
		return true
	}

	return false
}

// SetStart gets a reference to the given MicrosoftGraphDateTimeTimeZone and assigns it to the Start field.
func (o *MicrosoftGraphTimeSlot) SetStart(v MicrosoftGraphDateTimeTimeZone) {
	o.Start = &v
}

func (o MicrosoftGraphTimeSlot) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.End != nil {
		toSerialize["end"] = o.End
	}
	if o.Start != nil {
		toSerialize["start"] = o.Start
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphTimeSlot struct {
	value *MicrosoftGraphTimeSlot
	isSet bool
}

func (v NullableMicrosoftGraphTimeSlot) Get() *MicrosoftGraphTimeSlot {
	return v.value
}

func (v *NullableMicrosoftGraphTimeSlot) Set(val *MicrosoftGraphTimeSlot) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphTimeSlot) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphTimeSlot) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphTimeSlot(val *MicrosoftGraphTimeSlot) *NullableMicrosoftGraphTimeSlot {
	return &NullableMicrosoftGraphTimeSlot{value: val, isSet: true}
}

func (v NullableMicrosoftGraphTimeSlot) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphTimeSlot) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


