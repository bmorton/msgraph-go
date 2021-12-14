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

// DeviceConfigurationAssignment The device configuration assignment entity assigns an AAD group to a specific device configuration.
type DeviceConfigurationAssignment struct {
	// The assignment target for the device configuration.
	Target AnyOfobject `json:"target,omitempty"`
}

// NewDeviceConfigurationAssignment instantiates a new DeviceConfigurationAssignment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceConfigurationAssignment() *DeviceConfigurationAssignment {
	this := DeviceConfigurationAssignment{}
	return &this
}

// NewDeviceConfigurationAssignmentWithDefaults instantiates a new DeviceConfigurationAssignment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceConfigurationAssignmentWithDefaults() *DeviceConfigurationAssignment {
	this := DeviceConfigurationAssignment{}
	return &this
}

// GetTarget returns the Target field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeviceConfigurationAssignment) GetTarget() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeviceConfigurationAssignment) GetTargetOk() (*AnyOfobject, bool) {
	if o == nil || o.Target == nil {
		return nil, false
	}
	return &o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *DeviceConfigurationAssignment) HasTarget() bool {
	if o != nil && o.Target != nil {
		return true
	}

	return false
}

// SetTarget gets a reference to the given AnyOfobject and assigns it to the Target field.
func (o *DeviceConfigurationAssignment) SetTarget(v AnyOfobject) {
	o.Target = v
}

func (o DeviceConfigurationAssignment) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Target != nil {
		toSerialize["target"] = o.Target
	}
	return json.Marshal(toSerialize)
}

type NullableDeviceConfigurationAssignment struct {
	value *DeviceConfigurationAssignment
	isSet bool
}

func (v NullableDeviceConfigurationAssignment) Get() *DeviceConfigurationAssignment {
	return v.value
}

func (v *NullableDeviceConfigurationAssignment) Set(val *DeviceConfigurationAssignment) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceConfigurationAssignment) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceConfigurationAssignment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceConfigurationAssignment(val *DeviceConfigurationAssignment) *NullableDeviceConfigurationAssignment {
	return &NullableDeviceConfigurationAssignment{value: val, isSet: true}
}

func (v NullableDeviceConfigurationAssignment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceConfigurationAssignment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


