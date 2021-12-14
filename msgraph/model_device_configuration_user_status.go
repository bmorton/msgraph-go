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

// DeviceConfigurationUserStatus struct for DeviceConfigurationUserStatus
type DeviceConfigurationUserStatus struct {
	// Devices count for that user.
	DevicesCount *int32 `json:"devicesCount,omitempty"`
	// Last modified date time of the policy report.
	LastReportedDateTime *time.Time `json:"lastReportedDateTime,omitempty"`
	// Compliance status of the policy report. Possible values are: unknown, notApplicable, compliant, remediated, nonCompliant, error, conflict, notAssigned.
	Status AnyOfmicrosoftGraphComplianceStatus `json:"status,omitempty"`
	// User name of the DevicePolicyStatus.
	UserDisplayName NullableString `json:"userDisplayName,omitempty"`
	// UserPrincipalName.
	UserPrincipalName NullableString `json:"userPrincipalName,omitempty"`
}

// NewDeviceConfigurationUserStatus instantiates a new DeviceConfigurationUserStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceConfigurationUserStatus() *DeviceConfigurationUserStatus {
	this := DeviceConfigurationUserStatus{}
	return &this
}

// NewDeviceConfigurationUserStatusWithDefaults instantiates a new DeviceConfigurationUserStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceConfigurationUserStatusWithDefaults() *DeviceConfigurationUserStatus {
	this := DeviceConfigurationUserStatus{}
	return &this
}

// GetDevicesCount returns the DevicesCount field value if set, zero value otherwise.
func (o *DeviceConfigurationUserStatus) GetDevicesCount() int32 {
	if o == nil || o.DevicesCount == nil {
		var ret int32
		return ret
	}
	return *o.DevicesCount
}

// GetDevicesCountOk returns a tuple with the DevicesCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceConfigurationUserStatus) GetDevicesCountOk() (*int32, bool) {
	if o == nil || o.DevicesCount == nil {
		return nil, false
	}
	return o.DevicesCount, true
}

// HasDevicesCount returns a boolean if a field has been set.
func (o *DeviceConfigurationUserStatus) HasDevicesCount() bool {
	if o != nil && o.DevicesCount != nil {
		return true
	}

	return false
}

// SetDevicesCount gets a reference to the given int32 and assigns it to the DevicesCount field.
func (o *DeviceConfigurationUserStatus) SetDevicesCount(v int32) {
	o.DevicesCount = &v
}

// GetLastReportedDateTime returns the LastReportedDateTime field value if set, zero value otherwise.
func (o *DeviceConfigurationUserStatus) GetLastReportedDateTime() time.Time {
	if o == nil || o.LastReportedDateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.LastReportedDateTime
}

// GetLastReportedDateTimeOk returns a tuple with the LastReportedDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceConfigurationUserStatus) GetLastReportedDateTimeOk() (*time.Time, bool) {
	if o == nil || o.LastReportedDateTime == nil {
		return nil, false
	}
	return o.LastReportedDateTime, true
}

// HasLastReportedDateTime returns a boolean if a field has been set.
func (o *DeviceConfigurationUserStatus) HasLastReportedDateTime() bool {
	if o != nil && o.LastReportedDateTime != nil {
		return true
	}

	return false
}

// SetLastReportedDateTime gets a reference to the given time.Time and assigns it to the LastReportedDateTime field.
func (o *DeviceConfigurationUserStatus) SetLastReportedDateTime(v time.Time) {
	o.LastReportedDateTime = &v
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeviceConfigurationUserStatus) GetStatus() AnyOfmicrosoftGraphComplianceStatus {
	if o == nil  {
		var ret AnyOfmicrosoftGraphComplianceStatus
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeviceConfigurationUserStatus) GetStatusOk() (*AnyOfmicrosoftGraphComplianceStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return &o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *DeviceConfigurationUserStatus) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given AnyOfmicrosoftGraphComplianceStatus and assigns it to the Status field.
func (o *DeviceConfigurationUserStatus) SetStatus(v AnyOfmicrosoftGraphComplianceStatus) {
	o.Status = v
}

// GetUserDisplayName returns the UserDisplayName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeviceConfigurationUserStatus) GetUserDisplayName() string {
	if o == nil || o.UserDisplayName.Get() == nil {
		var ret string
		return ret
	}
	return *o.UserDisplayName.Get()
}

// GetUserDisplayNameOk returns a tuple with the UserDisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeviceConfigurationUserStatus) GetUserDisplayNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.UserDisplayName.Get(), o.UserDisplayName.IsSet()
}

// HasUserDisplayName returns a boolean if a field has been set.
func (o *DeviceConfigurationUserStatus) HasUserDisplayName() bool {
	if o != nil && o.UserDisplayName.IsSet() {
		return true
	}

	return false
}

// SetUserDisplayName gets a reference to the given NullableString and assigns it to the UserDisplayName field.
func (o *DeviceConfigurationUserStatus) SetUserDisplayName(v string) {
	o.UserDisplayName.Set(&v)
}
// SetUserDisplayNameNil sets the value for UserDisplayName to be an explicit nil
func (o *DeviceConfigurationUserStatus) SetUserDisplayNameNil() {
	o.UserDisplayName.Set(nil)
}

// UnsetUserDisplayName ensures that no value is present for UserDisplayName, not even an explicit nil
func (o *DeviceConfigurationUserStatus) UnsetUserDisplayName() {
	o.UserDisplayName.Unset()
}

// GetUserPrincipalName returns the UserPrincipalName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeviceConfigurationUserStatus) GetUserPrincipalName() string {
	if o == nil || o.UserPrincipalName.Get() == nil {
		var ret string
		return ret
	}
	return *o.UserPrincipalName.Get()
}

// GetUserPrincipalNameOk returns a tuple with the UserPrincipalName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeviceConfigurationUserStatus) GetUserPrincipalNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.UserPrincipalName.Get(), o.UserPrincipalName.IsSet()
}

// HasUserPrincipalName returns a boolean if a field has been set.
func (o *DeviceConfigurationUserStatus) HasUserPrincipalName() bool {
	if o != nil && o.UserPrincipalName.IsSet() {
		return true
	}

	return false
}

// SetUserPrincipalName gets a reference to the given NullableString and assigns it to the UserPrincipalName field.
func (o *DeviceConfigurationUserStatus) SetUserPrincipalName(v string) {
	o.UserPrincipalName.Set(&v)
}
// SetUserPrincipalNameNil sets the value for UserPrincipalName to be an explicit nil
func (o *DeviceConfigurationUserStatus) SetUserPrincipalNameNil() {
	o.UserPrincipalName.Set(nil)
}

// UnsetUserPrincipalName ensures that no value is present for UserPrincipalName, not even an explicit nil
func (o *DeviceConfigurationUserStatus) UnsetUserPrincipalName() {
	o.UserPrincipalName.Unset()
}

func (o DeviceConfigurationUserStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DevicesCount != nil {
		toSerialize["devicesCount"] = o.DevicesCount
	}
	if o.LastReportedDateTime != nil {
		toSerialize["lastReportedDateTime"] = o.LastReportedDateTime
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.UserDisplayName.IsSet() {
		toSerialize["userDisplayName"] = o.UserDisplayName.Get()
	}
	if o.UserPrincipalName.IsSet() {
		toSerialize["userPrincipalName"] = o.UserPrincipalName.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableDeviceConfigurationUserStatus struct {
	value *DeviceConfigurationUserStatus
	isSet bool
}

func (v NullableDeviceConfigurationUserStatus) Get() *DeviceConfigurationUserStatus {
	return v.value
}

func (v *NullableDeviceConfigurationUserStatus) Set(val *DeviceConfigurationUserStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceConfigurationUserStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceConfigurationUserStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceConfigurationUserStatus(val *DeviceConfigurationUserStatus) *NullableDeviceConfigurationUserStatus {
	return &NullableDeviceConfigurationUserStatus{value: val, isSet: true}
}

func (v NullableDeviceConfigurationUserStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceConfigurationUserStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


