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

// InlineObject1083 struct for InlineObject1083
type InlineObject1083 struct {
	StorageLocation NullableString `json:"storageLocation,omitempty"`
}

// NewInlineObject1083 instantiates a new InlineObject1083 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject1083() *InlineObject1083 {
	this := InlineObject1083{}
	return &this
}

// NewInlineObject1083WithDefaults instantiates a new InlineObject1083 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject1083WithDefaults() *InlineObject1083 {
	this := InlineObject1083{}
	return &this
}

// GetStorageLocation returns the StorageLocation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1083) GetStorageLocation() string {
	if o == nil || o.StorageLocation.Get() == nil {
		var ret string
		return ret
	}
	return *o.StorageLocation.Get()
}

// GetStorageLocationOk returns a tuple with the StorageLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1083) GetStorageLocationOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.StorageLocation.Get(), o.StorageLocation.IsSet()
}

// HasStorageLocation returns a boolean if a field has been set.
func (o *InlineObject1083) HasStorageLocation() bool {
	if o != nil && o.StorageLocation.IsSet() {
		return true
	}

	return false
}

// SetStorageLocation gets a reference to the given NullableString and assigns it to the StorageLocation field.
func (o *InlineObject1083) SetStorageLocation(v string) {
	o.StorageLocation.Set(&v)
}
// SetStorageLocationNil sets the value for StorageLocation to be an explicit nil
func (o *InlineObject1083) SetStorageLocationNil() {
	o.StorageLocation.Set(nil)
}

// UnsetStorageLocation ensures that no value is present for StorageLocation, not even an explicit nil
func (o *InlineObject1083) UnsetStorageLocation() {
	o.StorageLocation.Unset()
}

func (o InlineObject1083) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.StorageLocation.IsSet() {
		toSerialize["storageLocation"] = o.StorageLocation.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject1083 struct {
	value *InlineObject1083
	isSet bool
}

func (v NullableInlineObject1083) Get() *InlineObject1083 {
	return v.value
}

func (v *NullableInlineObject1083) Set(val *InlineObject1083) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject1083) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject1083) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject1083(val *InlineObject1083) *NullableInlineObject1083 {
	return &NullableInlineObject1083{value: val, isSet: true}
}

func (v NullableInlineObject1083) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject1083) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


