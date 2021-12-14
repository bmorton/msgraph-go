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

// MicrosoftGraphDeviceConfigurationDeviceStatus struct for MicrosoftGraphDeviceConfigurationDeviceStatus
type MicrosoftGraphDeviceConfigurationDeviceStatus struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
	// The DateTime when device compliance grace period expires
	ComplianceGracePeriodExpirationDateTime *time.Time `json:"complianceGracePeriodExpirationDateTime,omitempty"`
	// Device name of the DevicePolicyStatus.
	DeviceDisplayName NullableString `json:"deviceDisplayName,omitempty"`
	// The device model that is being reported
	DeviceModel NullableString `json:"deviceModel,omitempty"`
	// Last modified date time of the policy report.
	LastReportedDateTime *time.Time `json:"lastReportedDateTime,omitempty"`
	// Compliance status of the policy report. Possible values are: unknown, notApplicable, compliant, remediated, nonCompliant, error, conflict, notAssigned.
	Status AnyOfmicrosoftGraphComplianceStatus `json:"status,omitempty"`
	// The User Name that is being reported
	UserName NullableString `json:"userName,omitempty"`
	// UserPrincipalName.
	UserPrincipalName NullableString `json:"userPrincipalName,omitempty"`
}

// NewMicrosoftGraphDeviceConfigurationDeviceStatus instantiates a new MicrosoftGraphDeviceConfigurationDeviceStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphDeviceConfigurationDeviceStatus() *MicrosoftGraphDeviceConfigurationDeviceStatus {
	this := MicrosoftGraphDeviceConfigurationDeviceStatus{}
	return &this
}

// NewMicrosoftGraphDeviceConfigurationDeviceStatusWithDefaults instantiates a new MicrosoftGraphDeviceConfigurationDeviceStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphDeviceConfigurationDeviceStatusWithDefaults() *MicrosoftGraphDeviceConfigurationDeviceStatus {
	this := MicrosoftGraphDeviceConfigurationDeviceStatus{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MicrosoftGraphDeviceConfigurationDeviceStatus) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphDeviceConfigurationDeviceStatus) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphDeviceConfigurationDeviceStatus) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MicrosoftGraphDeviceConfigurationDeviceStatus) SetId(v string) {
	o.Id = &v
}

// GetComplianceGracePeriodExpirationDateTime returns the ComplianceGracePeriodExpirationDateTime field value if set, zero value otherwise.
func (o *MicrosoftGraphDeviceConfigurationDeviceStatus) GetComplianceGracePeriodExpirationDateTime() time.Time {
	if o == nil || o.ComplianceGracePeriodExpirationDateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.ComplianceGracePeriodExpirationDateTime
}

// GetComplianceGracePeriodExpirationDateTimeOk returns a tuple with the ComplianceGracePeriodExpirationDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphDeviceConfigurationDeviceStatus) GetComplianceGracePeriodExpirationDateTimeOk() (*time.Time, bool) {
	if o == nil || o.ComplianceGracePeriodExpirationDateTime == nil {
		return nil, false
	}
	return o.ComplianceGracePeriodExpirationDateTime, true
}

// HasComplianceGracePeriodExpirationDateTime returns a boolean if a field has been set.
func (o *MicrosoftGraphDeviceConfigurationDeviceStatus) HasComplianceGracePeriodExpirationDateTime() bool {
	if o != nil && o.ComplianceGracePeriodExpirationDateTime != nil {
		return true
	}

	return false
}

// SetComplianceGracePeriodExpirationDateTime gets a reference to the given time.Time and assigns it to the ComplianceGracePeriodExpirationDateTime field.
func (o *MicrosoftGraphDeviceConfigurationDeviceStatus) SetComplianceGracePeriodExpirationDateTime(v time.Time) {
	o.ComplianceGracePeriodExpirationDateTime = &v
}

// GetDeviceDisplayName returns the DeviceDisplayName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphDeviceConfigurationDeviceStatus) GetDeviceDisplayName() string {
	if o == nil || o.DeviceDisplayName.Get() == nil {
		var ret string
		return ret
	}
	return *o.DeviceDisplayName.Get()
}

// GetDeviceDisplayNameOk returns a tuple with the DeviceDisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphDeviceConfigurationDeviceStatus) GetDeviceDisplayNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DeviceDisplayName.Get(), o.DeviceDisplayName.IsSet()
}

// HasDeviceDisplayName returns a boolean if a field has been set.
func (o *MicrosoftGraphDeviceConfigurationDeviceStatus) HasDeviceDisplayName() bool {
	if o != nil && o.DeviceDisplayName.IsSet() {
		return true
	}

	return false
}

// SetDeviceDisplayName gets a reference to the given NullableString and assigns it to the DeviceDisplayName field.
func (o *MicrosoftGraphDeviceConfigurationDeviceStatus) SetDeviceDisplayName(v string) {
	o.DeviceDisplayName.Set(&v)
}
// SetDeviceDisplayNameNil sets the value for DeviceDisplayName to be an explicit nil
func (o *MicrosoftGraphDeviceConfigurationDeviceStatus) SetDeviceDisplayNameNil() {
	o.DeviceDisplayName.Set(nil)
}

// UnsetDeviceDisplayName ensures that no value is present for DeviceDisplayName, not even an explicit nil
func (o *MicrosoftGraphDeviceConfigurationDeviceStatus) UnsetDeviceDisplayName() {
	o.DeviceDisplayName.Unset()
}

// GetDeviceModel returns the DeviceModel field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphDeviceConfigurationDeviceStatus) GetDeviceModel() string {
	if o == nil || o.DeviceModel.Get() == nil {
		var ret string
		return ret
	}
	return *o.DeviceModel.Get()
}

// GetDeviceModelOk returns a tuple with the DeviceModel field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphDeviceConfigurationDeviceStatus) GetDeviceModelOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DeviceModel.Get(), o.DeviceModel.IsSet()
}

// HasDeviceModel returns a boolean if a field has been set.
func (o *MicrosoftGraphDeviceConfigurationDeviceStatus) HasDeviceModel() bool {
	if o != nil && o.DeviceModel.IsSet() {
		return true
	}

	return false
}

// SetDeviceModel gets a reference to the given NullableString and assigns it to the DeviceModel field.
func (o *MicrosoftGraphDeviceConfigurationDeviceStatus) SetDeviceModel(v string) {
	o.DeviceModel.Set(&v)
}
// SetDeviceModelNil sets the value for DeviceModel to be an explicit nil
func (o *MicrosoftGraphDeviceConfigurationDeviceStatus) SetDeviceModelNil() {
	o.DeviceModel.Set(nil)
}

// UnsetDeviceModel ensures that no value is present for DeviceModel, not even an explicit nil
func (o *MicrosoftGraphDeviceConfigurationDeviceStatus) UnsetDeviceModel() {
	o.DeviceModel.Unset()
}

// GetLastReportedDateTime returns the LastReportedDateTime field value if set, zero value otherwise.
func (o *MicrosoftGraphDeviceConfigurationDeviceStatus) GetLastReportedDateTime() time.Time {
	if o == nil || o.LastReportedDateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.LastReportedDateTime
}

// GetLastReportedDateTimeOk returns a tuple with the LastReportedDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphDeviceConfigurationDeviceStatus) GetLastReportedDateTimeOk() (*time.Time, bool) {
	if o == nil || o.LastReportedDateTime == nil {
		return nil, false
	}
	return o.LastReportedDateTime, true
}

// HasLastReportedDateTime returns a boolean if a field has been set.
func (o *MicrosoftGraphDeviceConfigurationDeviceStatus) HasLastReportedDateTime() bool {
	if o != nil && o.LastReportedDateTime != nil {
		return true
	}

	return false
}

// SetLastReportedDateTime gets a reference to the given time.Time and assigns it to the LastReportedDateTime field.
func (o *MicrosoftGraphDeviceConfigurationDeviceStatus) SetLastReportedDateTime(v time.Time) {
	o.LastReportedDateTime = &v
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphDeviceConfigurationDeviceStatus) GetStatus() AnyOfmicrosoftGraphComplianceStatus {
	if o == nil  {
		var ret AnyOfmicrosoftGraphComplianceStatus
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphDeviceConfigurationDeviceStatus) GetStatusOk() (*AnyOfmicrosoftGraphComplianceStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return &o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *MicrosoftGraphDeviceConfigurationDeviceStatus) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given AnyOfmicrosoftGraphComplianceStatus and assigns it to the Status field.
func (o *MicrosoftGraphDeviceConfigurationDeviceStatus) SetStatus(v AnyOfmicrosoftGraphComplianceStatus) {
	o.Status = v
}

// GetUserName returns the UserName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphDeviceConfigurationDeviceStatus) GetUserName() string {
	if o == nil || o.UserName.Get() == nil {
		var ret string
		return ret
	}
	return *o.UserName.Get()
}

// GetUserNameOk returns a tuple with the UserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphDeviceConfigurationDeviceStatus) GetUserNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.UserName.Get(), o.UserName.IsSet()
}

// HasUserName returns a boolean if a field has been set.
func (o *MicrosoftGraphDeviceConfigurationDeviceStatus) HasUserName() bool {
	if o != nil && o.UserName.IsSet() {
		return true
	}

	return false
}

// SetUserName gets a reference to the given NullableString and assigns it to the UserName field.
func (o *MicrosoftGraphDeviceConfigurationDeviceStatus) SetUserName(v string) {
	o.UserName.Set(&v)
}
// SetUserNameNil sets the value for UserName to be an explicit nil
func (o *MicrosoftGraphDeviceConfigurationDeviceStatus) SetUserNameNil() {
	o.UserName.Set(nil)
}

// UnsetUserName ensures that no value is present for UserName, not even an explicit nil
func (o *MicrosoftGraphDeviceConfigurationDeviceStatus) UnsetUserName() {
	o.UserName.Unset()
}

// GetUserPrincipalName returns the UserPrincipalName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphDeviceConfigurationDeviceStatus) GetUserPrincipalName() string {
	if o == nil || o.UserPrincipalName.Get() == nil {
		var ret string
		return ret
	}
	return *o.UserPrincipalName.Get()
}

// GetUserPrincipalNameOk returns a tuple with the UserPrincipalName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphDeviceConfigurationDeviceStatus) GetUserPrincipalNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.UserPrincipalName.Get(), o.UserPrincipalName.IsSet()
}

// HasUserPrincipalName returns a boolean if a field has been set.
func (o *MicrosoftGraphDeviceConfigurationDeviceStatus) HasUserPrincipalName() bool {
	if o != nil && o.UserPrincipalName.IsSet() {
		return true
	}

	return false
}

// SetUserPrincipalName gets a reference to the given NullableString and assigns it to the UserPrincipalName field.
func (o *MicrosoftGraphDeviceConfigurationDeviceStatus) SetUserPrincipalName(v string) {
	o.UserPrincipalName.Set(&v)
}
// SetUserPrincipalNameNil sets the value for UserPrincipalName to be an explicit nil
func (o *MicrosoftGraphDeviceConfigurationDeviceStatus) SetUserPrincipalNameNil() {
	o.UserPrincipalName.Set(nil)
}

// UnsetUserPrincipalName ensures that no value is present for UserPrincipalName, not even an explicit nil
func (o *MicrosoftGraphDeviceConfigurationDeviceStatus) UnsetUserPrincipalName() {
	o.UserPrincipalName.Unset()
}

func (o MicrosoftGraphDeviceConfigurationDeviceStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.ComplianceGracePeriodExpirationDateTime != nil {
		toSerialize["complianceGracePeriodExpirationDateTime"] = o.ComplianceGracePeriodExpirationDateTime
	}
	if o.DeviceDisplayName.IsSet() {
		toSerialize["deviceDisplayName"] = o.DeviceDisplayName.Get()
	}
	if o.DeviceModel.IsSet() {
		toSerialize["deviceModel"] = o.DeviceModel.Get()
	}
	if o.LastReportedDateTime != nil {
		toSerialize["lastReportedDateTime"] = o.LastReportedDateTime
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.UserName.IsSet() {
		toSerialize["userName"] = o.UserName.Get()
	}
	if o.UserPrincipalName.IsSet() {
		toSerialize["userPrincipalName"] = o.UserPrincipalName.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphDeviceConfigurationDeviceStatus struct {
	value *MicrosoftGraphDeviceConfigurationDeviceStatus
	isSet bool
}

func (v NullableMicrosoftGraphDeviceConfigurationDeviceStatus) Get() *MicrosoftGraphDeviceConfigurationDeviceStatus {
	return v.value
}

func (v *NullableMicrosoftGraphDeviceConfigurationDeviceStatus) Set(val *MicrosoftGraphDeviceConfigurationDeviceStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphDeviceConfigurationDeviceStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphDeviceConfigurationDeviceStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphDeviceConfigurationDeviceStatus(val *MicrosoftGraphDeviceConfigurationDeviceStatus) *NullableMicrosoftGraphDeviceConfigurationDeviceStatus {
	return &NullableMicrosoftGraphDeviceConfigurationDeviceStatus{value: val, isSet: true}
}

func (v NullableMicrosoftGraphDeviceConfigurationDeviceStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphDeviceConfigurationDeviceStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


