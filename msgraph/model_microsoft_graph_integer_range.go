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

// MicrosoftGraphIntegerRange struct for MicrosoftGraphIntegerRange
type MicrosoftGraphIntegerRange struct {
	// The inclusive upper bound of the integer range.
	End NullableInt64 `json:"end,omitempty"`
	// The inclusive lower bound of the integer range.
	Start NullableInt64 `json:"start,omitempty"`
}

// NewMicrosoftGraphIntegerRange instantiates a new MicrosoftGraphIntegerRange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphIntegerRange() *MicrosoftGraphIntegerRange {
	this := MicrosoftGraphIntegerRange{}
	return &this
}

// NewMicrosoftGraphIntegerRangeWithDefaults instantiates a new MicrosoftGraphIntegerRange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphIntegerRangeWithDefaults() *MicrosoftGraphIntegerRange {
	this := MicrosoftGraphIntegerRange{}
	return &this
}

// GetEnd returns the End field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphIntegerRange) GetEnd() int64 {
	if o == nil || o.End.Get() == nil {
		var ret int64
		return ret
	}
	return *o.End.Get()
}

// GetEndOk returns a tuple with the End field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphIntegerRange) GetEndOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.End.Get(), o.End.IsSet()
}

// HasEnd returns a boolean if a field has been set.
func (o *MicrosoftGraphIntegerRange) HasEnd() bool {
	if o != nil && o.End.IsSet() {
		return true
	}

	return false
}

// SetEnd gets a reference to the given NullableInt64 and assigns it to the End field.
func (o *MicrosoftGraphIntegerRange) SetEnd(v int64) {
	o.End.Set(&v)
}
// SetEndNil sets the value for End to be an explicit nil
func (o *MicrosoftGraphIntegerRange) SetEndNil() {
	o.End.Set(nil)
}

// UnsetEnd ensures that no value is present for End, not even an explicit nil
func (o *MicrosoftGraphIntegerRange) UnsetEnd() {
	o.End.Unset()
}

// GetStart returns the Start field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphIntegerRange) GetStart() int64 {
	if o == nil || o.Start.Get() == nil {
		var ret int64
		return ret
	}
	return *o.Start.Get()
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphIntegerRange) GetStartOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Start.Get(), o.Start.IsSet()
}

// HasStart returns a boolean if a field has been set.
func (o *MicrosoftGraphIntegerRange) HasStart() bool {
	if o != nil && o.Start.IsSet() {
		return true
	}

	return false
}

// SetStart gets a reference to the given NullableInt64 and assigns it to the Start field.
func (o *MicrosoftGraphIntegerRange) SetStart(v int64) {
	o.Start.Set(&v)
}
// SetStartNil sets the value for Start to be an explicit nil
func (o *MicrosoftGraphIntegerRange) SetStartNil() {
	o.Start.Set(nil)
}

// UnsetStart ensures that no value is present for Start, not even an explicit nil
func (o *MicrosoftGraphIntegerRange) UnsetStart() {
	o.Start.Unset()
}

func (o MicrosoftGraphIntegerRange) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.End.IsSet() {
		toSerialize["end"] = o.End.Get()
	}
	if o.Start.IsSet() {
		toSerialize["start"] = o.Start.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphIntegerRange struct {
	value *MicrosoftGraphIntegerRange
	isSet bool
}

func (v NullableMicrosoftGraphIntegerRange) Get() *MicrosoftGraphIntegerRange {
	return v.value
}

func (v *NullableMicrosoftGraphIntegerRange) Set(val *MicrosoftGraphIntegerRange) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphIntegerRange) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphIntegerRange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphIntegerRange(val *MicrosoftGraphIntegerRange) *NullableMicrosoftGraphIntegerRange {
	return &NullableMicrosoftGraphIntegerRange{value: val, isSet: true}
}

func (v NullableMicrosoftGraphIntegerRange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphIntegerRange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


