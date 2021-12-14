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

// LocationConstraintItem struct for LocationConstraintItem
type LocationConstraintItem struct {
	// If set to true and the specified resource is busy, findMeetingTimes looks for another resource that is free. If set to false and the specified resource is busy, findMeetingTimes returns the resource best ranked in the user's cache without checking if it's free. Default is true.
	ResolveAvailability NullableBool `json:"resolveAvailability,omitempty"`
}

// NewLocationConstraintItem instantiates a new LocationConstraintItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLocationConstraintItem() *LocationConstraintItem {
	this := LocationConstraintItem{}
	return &this
}

// NewLocationConstraintItemWithDefaults instantiates a new LocationConstraintItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLocationConstraintItemWithDefaults() *LocationConstraintItem {
	this := LocationConstraintItem{}
	return &this
}

// GetResolveAvailability returns the ResolveAvailability field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LocationConstraintItem) GetResolveAvailability() bool {
	if o == nil || o.ResolveAvailability.Get() == nil {
		var ret bool
		return ret
	}
	return *o.ResolveAvailability.Get()
}

// GetResolveAvailabilityOk returns a tuple with the ResolveAvailability field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LocationConstraintItem) GetResolveAvailabilityOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ResolveAvailability.Get(), o.ResolveAvailability.IsSet()
}

// HasResolveAvailability returns a boolean if a field has been set.
func (o *LocationConstraintItem) HasResolveAvailability() bool {
	if o != nil && o.ResolveAvailability.IsSet() {
		return true
	}

	return false
}

// SetResolveAvailability gets a reference to the given NullableBool and assigns it to the ResolveAvailability field.
func (o *LocationConstraintItem) SetResolveAvailability(v bool) {
	o.ResolveAvailability.Set(&v)
}
// SetResolveAvailabilityNil sets the value for ResolveAvailability to be an explicit nil
func (o *LocationConstraintItem) SetResolveAvailabilityNil() {
	o.ResolveAvailability.Set(nil)
}

// UnsetResolveAvailability ensures that no value is present for ResolveAvailability, not even an explicit nil
func (o *LocationConstraintItem) UnsetResolveAvailability() {
	o.ResolveAvailability.Unset()
}

func (o LocationConstraintItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ResolveAvailability.IsSet() {
		toSerialize["resolveAvailability"] = o.ResolveAvailability.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableLocationConstraintItem struct {
	value *LocationConstraintItem
	isSet bool
}

func (v NullableLocationConstraintItem) Get() *LocationConstraintItem {
	return v.value
}

func (v *NullableLocationConstraintItem) Set(val *LocationConstraintItem) {
	v.value = val
	v.isSet = true
}

func (v NullableLocationConstraintItem) IsSet() bool {
	return v.isSet
}

func (v *NullableLocationConstraintItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocationConstraintItem(val *LocationConstraintItem) *NullableLocationConstraintItem {
	return &NullableLocationConstraintItem{value: val, isSet: true}
}

func (v NullableLocationConstraintItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocationConstraintItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


