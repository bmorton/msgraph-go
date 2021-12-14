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

// ManagedAppStatus Represents app protection and configuration status for the organization.
type ManagedAppStatus struct {
	// Friendly name of the status report.
	DisplayName NullableString `json:"displayName,omitempty"`
	// Version of the entity.
	Version NullableString `json:"version,omitempty"`
}

// NewManagedAppStatus instantiates a new ManagedAppStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManagedAppStatus() *ManagedAppStatus {
	this := ManagedAppStatus{}
	return &this
}

// NewManagedAppStatusWithDefaults instantiates a new ManagedAppStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManagedAppStatusWithDefaults() *ManagedAppStatus {
	this := ManagedAppStatus{}
	return &this
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ManagedAppStatus) GetDisplayName() string {
	if o == nil || o.DisplayName.Get() == nil {
		var ret string
		return ret
	}
	return *o.DisplayName.Get()
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ManagedAppStatus) GetDisplayNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DisplayName.Get(), o.DisplayName.IsSet()
}

// HasDisplayName returns a boolean if a field has been set.
func (o *ManagedAppStatus) HasDisplayName() bool {
	if o != nil && o.DisplayName.IsSet() {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given NullableString and assigns it to the DisplayName field.
func (o *ManagedAppStatus) SetDisplayName(v string) {
	o.DisplayName.Set(&v)
}
// SetDisplayNameNil sets the value for DisplayName to be an explicit nil
func (o *ManagedAppStatus) SetDisplayNameNil() {
	o.DisplayName.Set(nil)
}

// UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
func (o *ManagedAppStatus) UnsetDisplayName() {
	o.DisplayName.Unset()
}

// GetVersion returns the Version field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ManagedAppStatus) GetVersion() string {
	if o == nil || o.Version.Get() == nil {
		var ret string
		return ret
	}
	return *o.Version.Get()
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ManagedAppStatus) GetVersionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Version.Get(), o.Version.IsSet()
}

// HasVersion returns a boolean if a field has been set.
func (o *ManagedAppStatus) HasVersion() bool {
	if o != nil && o.Version.IsSet() {
		return true
	}

	return false
}

// SetVersion gets a reference to the given NullableString and assigns it to the Version field.
func (o *ManagedAppStatus) SetVersion(v string) {
	o.Version.Set(&v)
}
// SetVersionNil sets the value for Version to be an explicit nil
func (o *ManagedAppStatus) SetVersionNil() {
	o.Version.Set(nil)
}

// UnsetVersion ensures that no value is present for Version, not even an explicit nil
func (o *ManagedAppStatus) UnsetVersion() {
	o.Version.Unset()
}

func (o ManagedAppStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DisplayName.IsSet() {
		toSerialize["displayName"] = o.DisplayName.Get()
	}
	if o.Version.IsSet() {
		toSerialize["version"] = o.Version.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableManagedAppStatus struct {
	value *ManagedAppStatus
	isSet bool
}

func (v NullableManagedAppStatus) Get() *ManagedAppStatus {
	return v.value
}

func (v *NullableManagedAppStatus) Set(val *ManagedAppStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableManagedAppStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableManagedAppStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManagedAppStatus(val *ManagedAppStatus) *NullableManagedAppStatus {
	return &NullableManagedAppStatus{value: val, isSet: true}
}

func (v NullableManagedAppStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManagedAppStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


