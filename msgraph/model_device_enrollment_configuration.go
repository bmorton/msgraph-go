/*
Partial Graph API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// DeviceEnrollmentConfiguration The Base Class of Device Enrollment Configuration
type DeviceEnrollmentConfiguration struct {
	// Created date time in UTC of the device enrollment configuration
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// The description of the device enrollment configuration
	Description NullableString `json:"description,omitempty"`
	// The display name of the device enrollment configuration
	DisplayName NullableString `json:"displayName,omitempty"`
	// Last modified date time in UTC of the device enrollment configuration
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
	// Priority is used when a user exists in multiple groups that are assigned enrollment configuration. Users are subject only to the configuration with the lowest priority value.
	Priority *int32 `json:"priority,omitempty"`
	// The version of the device enrollment configuration
	Version *int32 `json:"version,omitempty"`
	// The list of group assignments for the device configuration profile
	Assignments *[]MicrosoftGraphEnrollmentConfigurationAssignment `json:"assignments,omitempty"`
}

// NewDeviceEnrollmentConfiguration instantiates a new DeviceEnrollmentConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceEnrollmentConfiguration() *DeviceEnrollmentConfiguration {
	this := DeviceEnrollmentConfiguration{}
	return &this
}

// NewDeviceEnrollmentConfigurationWithDefaults instantiates a new DeviceEnrollmentConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceEnrollmentConfigurationWithDefaults() *DeviceEnrollmentConfiguration {
	this := DeviceEnrollmentConfiguration{}
	return &this
}

// GetCreatedDateTime returns the CreatedDateTime field value if set, zero value otherwise.
func (o *DeviceEnrollmentConfiguration) GetCreatedDateTime() time.Time {
	if o == nil || o.CreatedDateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedDateTime
}

// GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceEnrollmentConfiguration) GetCreatedDateTimeOk() (*time.Time, bool) {
	if o == nil || o.CreatedDateTime == nil {
		return nil, false
	}
	return o.CreatedDateTime, true
}

// HasCreatedDateTime returns a boolean if a field has been set.
func (o *DeviceEnrollmentConfiguration) HasCreatedDateTime() bool {
	if o != nil && o.CreatedDateTime != nil {
		return true
	}

	return false
}

// SetCreatedDateTime gets a reference to the given time.Time and assigns it to the CreatedDateTime field.
func (o *DeviceEnrollmentConfiguration) SetCreatedDateTime(v time.Time) {
	o.CreatedDateTime = &v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeviceEnrollmentConfiguration) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeviceEnrollmentConfiguration) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *DeviceEnrollmentConfiguration) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *DeviceEnrollmentConfiguration) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *DeviceEnrollmentConfiguration) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *DeviceEnrollmentConfiguration) UnsetDescription() {
	o.Description.Unset()
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeviceEnrollmentConfiguration) GetDisplayName() string {
	if o == nil || o.DisplayName.Get() == nil {
		var ret string
		return ret
	}
	return *o.DisplayName.Get()
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeviceEnrollmentConfiguration) GetDisplayNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DisplayName.Get(), o.DisplayName.IsSet()
}

// HasDisplayName returns a boolean if a field has been set.
func (o *DeviceEnrollmentConfiguration) HasDisplayName() bool {
	if o != nil && o.DisplayName.IsSet() {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given NullableString and assigns it to the DisplayName field.
func (o *DeviceEnrollmentConfiguration) SetDisplayName(v string) {
	o.DisplayName.Set(&v)
}
// SetDisplayNameNil sets the value for DisplayName to be an explicit nil
func (o *DeviceEnrollmentConfiguration) SetDisplayNameNil() {
	o.DisplayName.Set(nil)
}

// UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
func (o *DeviceEnrollmentConfiguration) UnsetDisplayName() {
	o.DisplayName.Unset()
}

// GetLastModifiedDateTime returns the LastModifiedDateTime field value if set, zero value otherwise.
func (o *DeviceEnrollmentConfiguration) GetLastModifiedDateTime() time.Time {
	if o == nil || o.LastModifiedDateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.LastModifiedDateTime
}

// GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceEnrollmentConfiguration) GetLastModifiedDateTimeOk() (*time.Time, bool) {
	if o == nil || o.LastModifiedDateTime == nil {
		return nil, false
	}
	return o.LastModifiedDateTime, true
}

// HasLastModifiedDateTime returns a boolean if a field has been set.
func (o *DeviceEnrollmentConfiguration) HasLastModifiedDateTime() bool {
	if o != nil && o.LastModifiedDateTime != nil {
		return true
	}

	return false
}

// SetLastModifiedDateTime gets a reference to the given time.Time and assigns it to the LastModifiedDateTime field.
func (o *DeviceEnrollmentConfiguration) SetLastModifiedDateTime(v time.Time) {
	o.LastModifiedDateTime = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *DeviceEnrollmentConfiguration) GetPriority() int32 {
	if o == nil || o.Priority == nil {
		var ret int32
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceEnrollmentConfiguration) GetPriorityOk() (*int32, bool) {
	if o == nil || o.Priority == nil {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *DeviceEnrollmentConfiguration) HasPriority() bool {
	if o != nil && o.Priority != nil {
		return true
	}

	return false
}

// SetPriority gets a reference to the given int32 and assigns it to the Priority field.
func (o *DeviceEnrollmentConfiguration) SetPriority(v int32) {
	o.Priority = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *DeviceEnrollmentConfiguration) GetVersion() int32 {
	if o == nil || o.Version == nil {
		var ret int32
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceEnrollmentConfiguration) GetVersionOk() (*int32, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *DeviceEnrollmentConfiguration) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int32 and assigns it to the Version field.
func (o *DeviceEnrollmentConfiguration) SetVersion(v int32) {
	o.Version = &v
}

// GetAssignments returns the Assignments field value if set, zero value otherwise.
func (o *DeviceEnrollmentConfiguration) GetAssignments() []MicrosoftGraphEnrollmentConfigurationAssignment {
	if o == nil || o.Assignments == nil {
		var ret []MicrosoftGraphEnrollmentConfigurationAssignment
		return ret
	}
	return *o.Assignments
}

// GetAssignmentsOk returns a tuple with the Assignments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceEnrollmentConfiguration) GetAssignmentsOk() (*[]MicrosoftGraphEnrollmentConfigurationAssignment, bool) {
	if o == nil || o.Assignments == nil {
		return nil, false
	}
	return o.Assignments, true
}

// HasAssignments returns a boolean if a field has been set.
func (o *DeviceEnrollmentConfiguration) HasAssignments() bool {
	if o != nil && o.Assignments != nil {
		return true
	}

	return false
}

// SetAssignments gets a reference to the given []MicrosoftGraphEnrollmentConfigurationAssignment and assigns it to the Assignments field.
func (o *DeviceEnrollmentConfiguration) SetAssignments(v []MicrosoftGraphEnrollmentConfigurationAssignment) {
	o.Assignments = &v
}

func (o DeviceEnrollmentConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CreatedDateTime != nil {
		toSerialize["createdDateTime"] = o.CreatedDateTime
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.DisplayName.IsSet() {
		toSerialize["displayName"] = o.DisplayName.Get()
	}
	if o.LastModifiedDateTime != nil {
		toSerialize["lastModifiedDateTime"] = o.LastModifiedDateTime
	}
	if o.Priority != nil {
		toSerialize["priority"] = o.Priority
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	if o.Assignments != nil {
		toSerialize["assignments"] = o.Assignments
	}
	return json.Marshal(toSerialize)
}

type NullableDeviceEnrollmentConfiguration struct {
	value *DeviceEnrollmentConfiguration
	isSet bool
}

func (v NullableDeviceEnrollmentConfiguration) Get() *DeviceEnrollmentConfiguration {
	return v.value
}

func (v *NullableDeviceEnrollmentConfiguration) Set(val *DeviceEnrollmentConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceEnrollmentConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceEnrollmentConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceEnrollmentConfiguration(val *DeviceEnrollmentConfiguration) *NullableDeviceEnrollmentConfiguration {
	return &NullableDeviceEnrollmentConfiguration{value: val, isSet: true}
}

func (v NullableDeviceEnrollmentConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceEnrollmentConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

