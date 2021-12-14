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

// DeviceAndAppManagementRoleAssignment The Role Assignment resource. Role assignments tie together a role definition with members and scopes. There can be one or more role assignments per role. This applies to custom and built-in roles.
type DeviceAndAppManagementRoleAssignment struct {
	// The list of ids of role member security groups. These are IDs from Azure Active Directory.
	Members *[]*string `json:"members,omitempty"`
}

// NewDeviceAndAppManagementRoleAssignment instantiates a new DeviceAndAppManagementRoleAssignment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceAndAppManagementRoleAssignment() *DeviceAndAppManagementRoleAssignment {
	this := DeviceAndAppManagementRoleAssignment{}
	return &this
}

// NewDeviceAndAppManagementRoleAssignmentWithDefaults instantiates a new DeviceAndAppManagementRoleAssignment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceAndAppManagementRoleAssignmentWithDefaults() *DeviceAndAppManagementRoleAssignment {
	this := DeviceAndAppManagementRoleAssignment{}
	return &this
}

// GetMembers returns the Members field value if set, zero value otherwise.
func (o *DeviceAndAppManagementRoleAssignment) GetMembers() []*string {
	if o == nil || o.Members == nil {
		var ret []*string
		return ret
	}
	return *o.Members
}

// GetMembersOk returns a tuple with the Members field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAndAppManagementRoleAssignment) GetMembersOk() (*[]*string, bool) {
	if o == nil || o.Members == nil {
		return nil, false
	}
	return o.Members, true
}

// HasMembers returns a boolean if a field has been set.
func (o *DeviceAndAppManagementRoleAssignment) HasMembers() bool {
	if o != nil && o.Members != nil {
		return true
	}

	return false
}

// SetMembers gets a reference to the given []*string and assigns it to the Members field.
func (o *DeviceAndAppManagementRoleAssignment) SetMembers(v []*string) {
	o.Members = &v
}

func (o DeviceAndAppManagementRoleAssignment) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Members != nil {
		toSerialize["members"] = o.Members
	}
	return json.Marshal(toSerialize)
}

type NullableDeviceAndAppManagementRoleAssignment struct {
	value *DeviceAndAppManagementRoleAssignment
	isSet bool
}

func (v NullableDeviceAndAppManagementRoleAssignment) Get() *DeviceAndAppManagementRoleAssignment {
	return v.value
}

func (v *NullableDeviceAndAppManagementRoleAssignment) Set(val *DeviceAndAppManagementRoleAssignment) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceAndAppManagementRoleAssignment) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceAndAppManagementRoleAssignment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceAndAppManagementRoleAssignment(val *DeviceAndAppManagementRoleAssignment) *NullableDeviceAndAppManagementRoleAssignment {
	return &NullableDeviceAndAppManagementRoleAssignment{value: val, isSet: true}
}

func (v NullableDeviceAndAppManagementRoleAssignment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceAndAppManagementRoleAssignment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

