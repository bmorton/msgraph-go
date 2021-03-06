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

// OpenShiftItem struct for OpenShiftItem
type OpenShiftItem struct {
	// Count of the number of slots for the given open shift.
	OpenSlotCount *int32 `json:"openSlotCount,omitempty"`
}

// NewOpenShiftItem instantiates a new OpenShiftItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOpenShiftItem() *OpenShiftItem {
	this := OpenShiftItem{}
	return &this
}

// NewOpenShiftItemWithDefaults instantiates a new OpenShiftItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpenShiftItemWithDefaults() *OpenShiftItem {
	this := OpenShiftItem{}
	return &this
}

// GetOpenSlotCount returns the OpenSlotCount field value if set, zero value otherwise.
func (o *OpenShiftItem) GetOpenSlotCount() int32 {
	if o == nil || o.OpenSlotCount == nil {
		var ret int32
		return ret
	}
	return *o.OpenSlotCount
}

// GetOpenSlotCountOk returns a tuple with the OpenSlotCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenShiftItem) GetOpenSlotCountOk() (*int32, bool) {
	if o == nil || o.OpenSlotCount == nil {
		return nil, false
	}
	return o.OpenSlotCount, true
}

// HasOpenSlotCount returns a boolean if a field has been set.
func (o *OpenShiftItem) HasOpenSlotCount() bool {
	if o != nil && o.OpenSlotCount != nil {
		return true
	}

	return false
}

// SetOpenSlotCount gets a reference to the given int32 and assigns it to the OpenSlotCount field.
func (o *OpenShiftItem) SetOpenSlotCount(v int32) {
	o.OpenSlotCount = &v
}

func (o OpenShiftItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.OpenSlotCount != nil {
		toSerialize["openSlotCount"] = o.OpenSlotCount
	}
	return json.Marshal(toSerialize)
}

type NullableOpenShiftItem struct {
	value *OpenShiftItem
	isSet bool
}

func (v NullableOpenShiftItem) Get() *OpenShiftItem {
	return v.value
}

func (v *NullableOpenShiftItem) Set(val *OpenShiftItem) {
	v.value = val
	v.isSet = true
}

func (v NullableOpenShiftItem) IsSet() bool {
	return v.isSet
}

func (v *NullableOpenShiftItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpenShiftItem(val *OpenShiftItem) *NullableOpenShiftItem {
	return &NullableOpenShiftItem{value: val, isSet: true}
}

func (v NullableOpenShiftItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpenShiftItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


