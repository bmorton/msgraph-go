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

// MicrosoftGraphManagedDeviceMobileAppConfigurationUserStatus struct for MicrosoftGraphManagedDeviceMobileAppConfigurationUserStatus
type MicrosoftGraphManagedDeviceMobileAppConfigurationUserStatus struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
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

// NewMicrosoftGraphManagedDeviceMobileAppConfigurationUserStatus instantiates a new MicrosoftGraphManagedDeviceMobileAppConfigurationUserStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphManagedDeviceMobileAppConfigurationUserStatus() *MicrosoftGraphManagedDeviceMobileAppConfigurationUserStatus {
	this := MicrosoftGraphManagedDeviceMobileAppConfigurationUserStatus{}
	return &this
}

// NewMicrosoftGraphManagedDeviceMobileAppConfigurationUserStatusWithDefaults instantiates a new MicrosoftGraphManagedDeviceMobileAppConfigurationUserStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphManagedDeviceMobileAppConfigurationUserStatusWithDefaults() *MicrosoftGraphManagedDeviceMobileAppConfigurationUserStatus {
	this := MicrosoftGraphManagedDeviceMobileAppConfigurationUserStatus{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MicrosoftGraphManagedDeviceMobileAppConfigurationUserStatus) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphManagedDeviceMobileAppConfigurationUserStatus) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphManagedDeviceMobileAppConfigurationUserStatus) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MicrosoftGraphManagedDeviceMobileAppConfigurationUserStatus) SetId(v string) {
	o.Id = &v
}

// GetDevicesCount returns the DevicesCount field value if set, zero value otherwise.
func (o *MicrosoftGraphManagedDeviceMobileAppConfigurationUserStatus) GetDevicesCount() int32 {
	if o == nil || o.DevicesCount == nil {
		var ret int32
		return ret
	}
	return *o.DevicesCount
}

// GetDevicesCountOk returns a tuple with the DevicesCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphManagedDeviceMobileAppConfigurationUserStatus) GetDevicesCountOk() (*int32, bool) {
	if o == nil || o.DevicesCount == nil {
		return nil, false
	}
	return o.DevicesCount, true
}

// HasDevicesCount returns a boolean if a field has been set.
func (o *MicrosoftGraphManagedDeviceMobileAppConfigurationUserStatus) HasDevicesCount() bool {
	if o != nil && o.DevicesCount != nil {
		return true
	}

	return false
}

// SetDevicesCount gets a reference to the given int32 and assigns it to the DevicesCount field.
func (o *MicrosoftGraphManagedDeviceMobileAppConfigurationUserStatus) SetDevicesCount(v int32) {
	o.DevicesCount = &v
}

// GetLastReportedDateTime returns the LastReportedDateTime field value if set, zero value otherwise.
func (o *MicrosoftGraphManagedDeviceMobileAppConfigurationUserStatus) GetLastReportedDateTime() time.Time {
	if o == nil || o.LastReportedDateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.LastReportedDateTime
}

// GetLastReportedDateTimeOk returns a tuple with the LastReportedDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphManagedDeviceMobileAppConfigurationUserStatus) GetLastReportedDateTimeOk() (*time.Time, bool) {
	if o == nil || o.LastReportedDateTime == nil {
		return nil, false
	}
	return o.LastReportedDateTime, true
}

// HasLastReportedDateTime returns a boolean if a field has been set.
func (o *MicrosoftGraphManagedDeviceMobileAppConfigurationUserStatus) HasLastReportedDateTime() bool {
	if o != nil && o.LastReportedDateTime != nil {
		return true
	}

	return false
}

// SetLastReportedDateTime gets a reference to the given time.Time and assigns it to the LastReportedDateTime field.
func (o *MicrosoftGraphManagedDeviceMobileAppConfigurationUserStatus) SetLastReportedDateTime(v time.Time) {
	o.LastReportedDateTime = &v
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphManagedDeviceMobileAppConfigurationUserStatus) GetStatus() AnyOfmicrosoftGraphComplianceStatus {
	if o == nil  {
		var ret AnyOfmicrosoftGraphComplianceStatus
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphManagedDeviceMobileAppConfigurationUserStatus) GetStatusOk() (*AnyOfmicrosoftGraphComplianceStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return &o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *MicrosoftGraphManagedDeviceMobileAppConfigurationUserStatus) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given AnyOfmicrosoftGraphComplianceStatus and assigns it to the Status field.
func (o *MicrosoftGraphManagedDeviceMobileAppConfigurationUserStatus) SetStatus(v AnyOfmicrosoftGraphComplianceStatus) {
	o.Status = v
}

// GetUserDisplayName returns the UserDisplayName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphManagedDeviceMobileAppConfigurationUserStatus) GetUserDisplayName() string {
	if o == nil || o.UserDisplayName.Get() == nil {
		var ret string
		return ret
	}
	return *o.UserDisplayName.Get()
}

// GetUserDisplayNameOk returns a tuple with the UserDisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphManagedDeviceMobileAppConfigurationUserStatus) GetUserDisplayNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.UserDisplayName.Get(), o.UserDisplayName.IsSet()
}

// HasUserDisplayName returns a boolean if a field has been set.
func (o *MicrosoftGraphManagedDeviceMobileAppConfigurationUserStatus) HasUserDisplayName() bool {
	if o != nil && o.UserDisplayName.IsSet() {
		return true
	}

	return false
}

// SetUserDisplayName gets a reference to the given NullableString and assigns it to the UserDisplayName field.
func (o *MicrosoftGraphManagedDeviceMobileAppConfigurationUserStatus) SetUserDisplayName(v string) {
	o.UserDisplayName.Set(&v)
}
// SetUserDisplayNameNil sets the value for UserDisplayName to be an explicit nil
func (o *MicrosoftGraphManagedDeviceMobileAppConfigurationUserStatus) SetUserDisplayNameNil() {
	o.UserDisplayName.Set(nil)
}

// UnsetUserDisplayName ensures that no value is present for UserDisplayName, not even an explicit nil
func (o *MicrosoftGraphManagedDeviceMobileAppConfigurationUserStatus) UnsetUserDisplayName() {
	o.UserDisplayName.Unset()
}

// GetUserPrincipalName returns the UserPrincipalName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphManagedDeviceMobileAppConfigurationUserStatus) GetUserPrincipalName() string {
	if o == nil || o.UserPrincipalName.Get() == nil {
		var ret string
		return ret
	}
	return *o.UserPrincipalName.Get()
}

// GetUserPrincipalNameOk returns a tuple with the UserPrincipalName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphManagedDeviceMobileAppConfigurationUserStatus) GetUserPrincipalNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.UserPrincipalName.Get(), o.UserPrincipalName.IsSet()
}

// HasUserPrincipalName returns a boolean if a field has been set.
func (o *MicrosoftGraphManagedDeviceMobileAppConfigurationUserStatus) HasUserPrincipalName() bool {
	if o != nil && o.UserPrincipalName.IsSet() {
		return true
	}

	return false
}

// SetUserPrincipalName gets a reference to the given NullableString and assigns it to the UserPrincipalName field.
func (o *MicrosoftGraphManagedDeviceMobileAppConfigurationUserStatus) SetUserPrincipalName(v string) {
	o.UserPrincipalName.Set(&v)
}
// SetUserPrincipalNameNil sets the value for UserPrincipalName to be an explicit nil
func (o *MicrosoftGraphManagedDeviceMobileAppConfigurationUserStatus) SetUserPrincipalNameNil() {
	o.UserPrincipalName.Set(nil)
}

// UnsetUserPrincipalName ensures that no value is present for UserPrincipalName, not even an explicit nil
func (o *MicrosoftGraphManagedDeviceMobileAppConfigurationUserStatus) UnsetUserPrincipalName() {
	o.UserPrincipalName.Unset()
}

func (o MicrosoftGraphManagedDeviceMobileAppConfigurationUserStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
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

type NullableMicrosoftGraphManagedDeviceMobileAppConfigurationUserStatus struct {
	value *MicrosoftGraphManagedDeviceMobileAppConfigurationUserStatus
	isSet bool
}

func (v NullableMicrosoftGraphManagedDeviceMobileAppConfigurationUserStatus) Get() *MicrosoftGraphManagedDeviceMobileAppConfigurationUserStatus {
	return v.value
}

func (v *NullableMicrosoftGraphManagedDeviceMobileAppConfigurationUserStatus) Set(val *MicrosoftGraphManagedDeviceMobileAppConfigurationUserStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphManagedDeviceMobileAppConfigurationUserStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphManagedDeviceMobileAppConfigurationUserStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphManagedDeviceMobileAppConfigurationUserStatus(val *MicrosoftGraphManagedDeviceMobileAppConfigurationUserStatus) *NullableMicrosoftGraphManagedDeviceMobileAppConfigurationUserStatus {
	return &NullableMicrosoftGraphManagedDeviceMobileAppConfigurationUserStatus{value: val, isSet: true}
}

func (v NullableMicrosoftGraphManagedDeviceMobileAppConfigurationUserStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphManagedDeviceMobileAppConfigurationUserStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


