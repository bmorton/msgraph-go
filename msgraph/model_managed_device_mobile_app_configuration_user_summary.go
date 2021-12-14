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

// ManagedDeviceMobileAppConfigurationUserSummary Contains properties, inherited properties and actions for an MDM mobile app configuration user status summary.
type ManagedDeviceMobileAppConfigurationUserSummary struct {
	// Version of the policy for that overview
	ConfigurationVersion *int32 `json:"configurationVersion,omitempty"`
	// Number of error Users
	ErrorCount *int32 `json:"errorCount,omitempty"`
	// Number of failed Users
	FailedCount *int32 `json:"failedCount,omitempty"`
	// Last update time
	LastUpdateDateTime *time.Time `json:"lastUpdateDateTime,omitempty"`
	// Number of not applicable users
	NotApplicableCount *int32 `json:"notApplicableCount,omitempty"`
	// Number of pending Users
	PendingCount *int32 `json:"pendingCount,omitempty"`
	// Number of succeeded Users
	SuccessCount *int32 `json:"successCount,omitempty"`
}

// NewManagedDeviceMobileAppConfigurationUserSummary instantiates a new ManagedDeviceMobileAppConfigurationUserSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManagedDeviceMobileAppConfigurationUserSummary() *ManagedDeviceMobileAppConfigurationUserSummary {
	this := ManagedDeviceMobileAppConfigurationUserSummary{}
	return &this
}

// NewManagedDeviceMobileAppConfigurationUserSummaryWithDefaults instantiates a new ManagedDeviceMobileAppConfigurationUserSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManagedDeviceMobileAppConfigurationUserSummaryWithDefaults() *ManagedDeviceMobileAppConfigurationUserSummary {
	this := ManagedDeviceMobileAppConfigurationUserSummary{}
	return &this
}

// GetConfigurationVersion returns the ConfigurationVersion field value if set, zero value otherwise.
func (o *ManagedDeviceMobileAppConfigurationUserSummary) GetConfigurationVersion() int32 {
	if o == nil || o.ConfigurationVersion == nil {
		var ret int32
		return ret
	}
	return *o.ConfigurationVersion
}

// GetConfigurationVersionOk returns a tuple with the ConfigurationVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedDeviceMobileAppConfigurationUserSummary) GetConfigurationVersionOk() (*int32, bool) {
	if o == nil || o.ConfigurationVersion == nil {
		return nil, false
	}
	return o.ConfigurationVersion, true
}

// HasConfigurationVersion returns a boolean if a field has been set.
func (o *ManagedDeviceMobileAppConfigurationUserSummary) HasConfigurationVersion() bool {
	if o != nil && o.ConfigurationVersion != nil {
		return true
	}

	return false
}

// SetConfigurationVersion gets a reference to the given int32 and assigns it to the ConfigurationVersion field.
func (o *ManagedDeviceMobileAppConfigurationUserSummary) SetConfigurationVersion(v int32) {
	o.ConfigurationVersion = &v
}

// GetErrorCount returns the ErrorCount field value if set, zero value otherwise.
func (o *ManagedDeviceMobileAppConfigurationUserSummary) GetErrorCount() int32 {
	if o == nil || o.ErrorCount == nil {
		var ret int32
		return ret
	}
	return *o.ErrorCount
}

// GetErrorCountOk returns a tuple with the ErrorCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedDeviceMobileAppConfigurationUserSummary) GetErrorCountOk() (*int32, bool) {
	if o == nil || o.ErrorCount == nil {
		return nil, false
	}
	return o.ErrorCount, true
}

// HasErrorCount returns a boolean if a field has been set.
func (o *ManagedDeviceMobileAppConfigurationUserSummary) HasErrorCount() bool {
	if o != nil && o.ErrorCount != nil {
		return true
	}

	return false
}

// SetErrorCount gets a reference to the given int32 and assigns it to the ErrorCount field.
func (o *ManagedDeviceMobileAppConfigurationUserSummary) SetErrorCount(v int32) {
	o.ErrorCount = &v
}

// GetFailedCount returns the FailedCount field value if set, zero value otherwise.
func (o *ManagedDeviceMobileAppConfigurationUserSummary) GetFailedCount() int32 {
	if o == nil || o.FailedCount == nil {
		var ret int32
		return ret
	}
	return *o.FailedCount
}

// GetFailedCountOk returns a tuple with the FailedCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedDeviceMobileAppConfigurationUserSummary) GetFailedCountOk() (*int32, bool) {
	if o == nil || o.FailedCount == nil {
		return nil, false
	}
	return o.FailedCount, true
}

// HasFailedCount returns a boolean if a field has been set.
func (o *ManagedDeviceMobileAppConfigurationUserSummary) HasFailedCount() bool {
	if o != nil && o.FailedCount != nil {
		return true
	}

	return false
}

// SetFailedCount gets a reference to the given int32 and assigns it to the FailedCount field.
func (o *ManagedDeviceMobileAppConfigurationUserSummary) SetFailedCount(v int32) {
	o.FailedCount = &v
}

// GetLastUpdateDateTime returns the LastUpdateDateTime field value if set, zero value otherwise.
func (o *ManagedDeviceMobileAppConfigurationUserSummary) GetLastUpdateDateTime() time.Time {
	if o == nil || o.LastUpdateDateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUpdateDateTime
}

// GetLastUpdateDateTimeOk returns a tuple with the LastUpdateDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedDeviceMobileAppConfigurationUserSummary) GetLastUpdateDateTimeOk() (*time.Time, bool) {
	if o == nil || o.LastUpdateDateTime == nil {
		return nil, false
	}
	return o.LastUpdateDateTime, true
}

// HasLastUpdateDateTime returns a boolean if a field has been set.
func (o *ManagedDeviceMobileAppConfigurationUserSummary) HasLastUpdateDateTime() bool {
	if o != nil && o.LastUpdateDateTime != nil {
		return true
	}

	return false
}

// SetLastUpdateDateTime gets a reference to the given time.Time and assigns it to the LastUpdateDateTime field.
func (o *ManagedDeviceMobileAppConfigurationUserSummary) SetLastUpdateDateTime(v time.Time) {
	o.LastUpdateDateTime = &v
}

// GetNotApplicableCount returns the NotApplicableCount field value if set, zero value otherwise.
func (o *ManagedDeviceMobileAppConfigurationUserSummary) GetNotApplicableCount() int32 {
	if o == nil || o.NotApplicableCount == nil {
		var ret int32
		return ret
	}
	return *o.NotApplicableCount
}

// GetNotApplicableCountOk returns a tuple with the NotApplicableCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedDeviceMobileAppConfigurationUserSummary) GetNotApplicableCountOk() (*int32, bool) {
	if o == nil || o.NotApplicableCount == nil {
		return nil, false
	}
	return o.NotApplicableCount, true
}

// HasNotApplicableCount returns a boolean if a field has been set.
func (o *ManagedDeviceMobileAppConfigurationUserSummary) HasNotApplicableCount() bool {
	if o != nil && o.NotApplicableCount != nil {
		return true
	}

	return false
}

// SetNotApplicableCount gets a reference to the given int32 and assigns it to the NotApplicableCount field.
func (o *ManagedDeviceMobileAppConfigurationUserSummary) SetNotApplicableCount(v int32) {
	o.NotApplicableCount = &v
}

// GetPendingCount returns the PendingCount field value if set, zero value otherwise.
func (o *ManagedDeviceMobileAppConfigurationUserSummary) GetPendingCount() int32 {
	if o == nil || o.PendingCount == nil {
		var ret int32
		return ret
	}
	return *o.PendingCount
}

// GetPendingCountOk returns a tuple with the PendingCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedDeviceMobileAppConfigurationUserSummary) GetPendingCountOk() (*int32, bool) {
	if o == nil || o.PendingCount == nil {
		return nil, false
	}
	return o.PendingCount, true
}

// HasPendingCount returns a boolean if a field has been set.
func (o *ManagedDeviceMobileAppConfigurationUserSummary) HasPendingCount() bool {
	if o != nil && o.PendingCount != nil {
		return true
	}

	return false
}

// SetPendingCount gets a reference to the given int32 and assigns it to the PendingCount field.
func (o *ManagedDeviceMobileAppConfigurationUserSummary) SetPendingCount(v int32) {
	o.PendingCount = &v
}

// GetSuccessCount returns the SuccessCount field value if set, zero value otherwise.
func (o *ManagedDeviceMobileAppConfigurationUserSummary) GetSuccessCount() int32 {
	if o == nil || o.SuccessCount == nil {
		var ret int32
		return ret
	}
	return *o.SuccessCount
}

// GetSuccessCountOk returns a tuple with the SuccessCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedDeviceMobileAppConfigurationUserSummary) GetSuccessCountOk() (*int32, bool) {
	if o == nil || o.SuccessCount == nil {
		return nil, false
	}
	return o.SuccessCount, true
}

// HasSuccessCount returns a boolean if a field has been set.
func (o *ManagedDeviceMobileAppConfigurationUserSummary) HasSuccessCount() bool {
	if o != nil && o.SuccessCount != nil {
		return true
	}

	return false
}

// SetSuccessCount gets a reference to the given int32 and assigns it to the SuccessCount field.
func (o *ManagedDeviceMobileAppConfigurationUserSummary) SetSuccessCount(v int32) {
	o.SuccessCount = &v
}

func (o ManagedDeviceMobileAppConfigurationUserSummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ConfigurationVersion != nil {
		toSerialize["configurationVersion"] = o.ConfigurationVersion
	}
	if o.ErrorCount != nil {
		toSerialize["errorCount"] = o.ErrorCount
	}
	if o.FailedCount != nil {
		toSerialize["failedCount"] = o.FailedCount
	}
	if o.LastUpdateDateTime != nil {
		toSerialize["lastUpdateDateTime"] = o.LastUpdateDateTime
	}
	if o.NotApplicableCount != nil {
		toSerialize["notApplicableCount"] = o.NotApplicableCount
	}
	if o.PendingCount != nil {
		toSerialize["pendingCount"] = o.PendingCount
	}
	if o.SuccessCount != nil {
		toSerialize["successCount"] = o.SuccessCount
	}
	return json.Marshal(toSerialize)
}

type NullableManagedDeviceMobileAppConfigurationUserSummary struct {
	value *ManagedDeviceMobileAppConfigurationUserSummary
	isSet bool
}

func (v NullableManagedDeviceMobileAppConfigurationUserSummary) Get() *ManagedDeviceMobileAppConfigurationUserSummary {
	return v.value
}

func (v *NullableManagedDeviceMobileAppConfigurationUserSummary) Set(val *ManagedDeviceMobileAppConfigurationUserSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableManagedDeviceMobileAppConfigurationUserSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableManagedDeviceMobileAppConfigurationUserSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManagedDeviceMobileAppConfigurationUserSummary(val *ManagedDeviceMobileAppConfigurationUserSummary) *NullableManagedDeviceMobileAppConfigurationUserSummary {
	return &NullableManagedDeviceMobileAppConfigurationUserSummary{value: val, isSet: true}
}

func (v NullableManagedDeviceMobileAppConfigurationUserSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManagedDeviceMobileAppConfigurationUserSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

