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

// MicrosoftGraphManagedDeviceMobileAppConfigurationDeviceSummary struct for MicrosoftGraphManagedDeviceMobileAppConfigurationDeviceSummary
type MicrosoftGraphManagedDeviceMobileAppConfigurationDeviceSummary struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
	// Version of the policy for that overview
	ConfigurationVersion *int32 `json:"configurationVersion,omitempty"`
	// Number of error devices
	ErrorCount *int32 `json:"errorCount,omitempty"`
	// Number of failed devices
	FailedCount *int32 `json:"failedCount,omitempty"`
	// Last update time
	LastUpdateDateTime *time.Time `json:"lastUpdateDateTime,omitempty"`
	// Number of not applicable devices
	NotApplicableCount *int32 `json:"notApplicableCount,omitempty"`
	// Number of pending devices
	PendingCount *int32 `json:"pendingCount,omitempty"`
	// Number of succeeded devices
	SuccessCount *int32 `json:"successCount,omitempty"`
}

// NewMicrosoftGraphManagedDeviceMobileAppConfigurationDeviceSummary instantiates a new MicrosoftGraphManagedDeviceMobileAppConfigurationDeviceSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphManagedDeviceMobileAppConfigurationDeviceSummary() *MicrosoftGraphManagedDeviceMobileAppConfigurationDeviceSummary {
	this := MicrosoftGraphManagedDeviceMobileAppConfigurationDeviceSummary{}
	return &this
}

// NewMicrosoftGraphManagedDeviceMobileAppConfigurationDeviceSummaryWithDefaults instantiates a new MicrosoftGraphManagedDeviceMobileAppConfigurationDeviceSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphManagedDeviceMobileAppConfigurationDeviceSummaryWithDefaults() *MicrosoftGraphManagedDeviceMobileAppConfigurationDeviceSummary {
	this := MicrosoftGraphManagedDeviceMobileAppConfigurationDeviceSummary{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MicrosoftGraphManagedDeviceMobileAppConfigurationDeviceSummary) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphManagedDeviceMobileAppConfigurationDeviceSummary) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphManagedDeviceMobileAppConfigurationDeviceSummary) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MicrosoftGraphManagedDeviceMobileAppConfigurationDeviceSummary) SetId(v string) {
	o.Id = &v
}

// GetConfigurationVersion returns the ConfigurationVersion field value if set, zero value otherwise.
func (o *MicrosoftGraphManagedDeviceMobileAppConfigurationDeviceSummary) GetConfigurationVersion() int32 {
	if o == nil || o.ConfigurationVersion == nil {
		var ret int32
		return ret
	}
	return *o.ConfigurationVersion
}

// GetConfigurationVersionOk returns a tuple with the ConfigurationVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphManagedDeviceMobileAppConfigurationDeviceSummary) GetConfigurationVersionOk() (*int32, bool) {
	if o == nil || o.ConfigurationVersion == nil {
		return nil, false
	}
	return o.ConfigurationVersion, true
}

// HasConfigurationVersion returns a boolean if a field has been set.
func (o *MicrosoftGraphManagedDeviceMobileAppConfigurationDeviceSummary) HasConfigurationVersion() bool {
	if o != nil && o.ConfigurationVersion != nil {
		return true
	}

	return false
}

// SetConfigurationVersion gets a reference to the given int32 and assigns it to the ConfigurationVersion field.
func (o *MicrosoftGraphManagedDeviceMobileAppConfigurationDeviceSummary) SetConfigurationVersion(v int32) {
	o.ConfigurationVersion = &v
}

// GetErrorCount returns the ErrorCount field value if set, zero value otherwise.
func (o *MicrosoftGraphManagedDeviceMobileAppConfigurationDeviceSummary) GetErrorCount() int32 {
	if o == nil || o.ErrorCount == nil {
		var ret int32
		return ret
	}
	return *o.ErrorCount
}

// GetErrorCountOk returns a tuple with the ErrorCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphManagedDeviceMobileAppConfigurationDeviceSummary) GetErrorCountOk() (*int32, bool) {
	if o == nil || o.ErrorCount == nil {
		return nil, false
	}
	return o.ErrorCount, true
}

// HasErrorCount returns a boolean if a field has been set.
func (o *MicrosoftGraphManagedDeviceMobileAppConfigurationDeviceSummary) HasErrorCount() bool {
	if o != nil && o.ErrorCount != nil {
		return true
	}

	return false
}

// SetErrorCount gets a reference to the given int32 and assigns it to the ErrorCount field.
func (o *MicrosoftGraphManagedDeviceMobileAppConfigurationDeviceSummary) SetErrorCount(v int32) {
	o.ErrorCount = &v
}

// GetFailedCount returns the FailedCount field value if set, zero value otherwise.
func (o *MicrosoftGraphManagedDeviceMobileAppConfigurationDeviceSummary) GetFailedCount() int32 {
	if o == nil || o.FailedCount == nil {
		var ret int32
		return ret
	}
	return *o.FailedCount
}

// GetFailedCountOk returns a tuple with the FailedCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphManagedDeviceMobileAppConfigurationDeviceSummary) GetFailedCountOk() (*int32, bool) {
	if o == nil || o.FailedCount == nil {
		return nil, false
	}
	return o.FailedCount, true
}

// HasFailedCount returns a boolean if a field has been set.
func (o *MicrosoftGraphManagedDeviceMobileAppConfigurationDeviceSummary) HasFailedCount() bool {
	if o != nil && o.FailedCount != nil {
		return true
	}

	return false
}

// SetFailedCount gets a reference to the given int32 and assigns it to the FailedCount field.
func (o *MicrosoftGraphManagedDeviceMobileAppConfigurationDeviceSummary) SetFailedCount(v int32) {
	o.FailedCount = &v
}

// GetLastUpdateDateTime returns the LastUpdateDateTime field value if set, zero value otherwise.
func (o *MicrosoftGraphManagedDeviceMobileAppConfigurationDeviceSummary) GetLastUpdateDateTime() time.Time {
	if o == nil || o.LastUpdateDateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUpdateDateTime
}

// GetLastUpdateDateTimeOk returns a tuple with the LastUpdateDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphManagedDeviceMobileAppConfigurationDeviceSummary) GetLastUpdateDateTimeOk() (*time.Time, bool) {
	if o == nil || o.LastUpdateDateTime == nil {
		return nil, false
	}
	return o.LastUpdateDateTime, true
}

// HasLastUpdateDateTime returns a boolean if a field has been set.
func (o *MicrosoftGraphManagedDeviceMobileAppConfigurationDeviceSummary) HasLastUpdateDateTime() bool {
	if o != nil && o.LastUpdateDateTime != nil {
		return true
	}

	return false
}

// SetLastUpdateDateTime gets a reference to the given time.Time and assigns it to the LastUpdateDateTime field.
func (o *MicrosoftGraphManagedDeviceMobileAppConfigurationDeviceSummary) SetLastUpdateDateTime(v time.Time) {
	o.LastUpdateDateTime = &v
}

// GetNotApplicableCount returns the NotApplicableCount field value if set, zero value otherwise.
func (o *MicrosoftGraphManagedDeviceMobileAppConfigurationDeviceSummary) GetNotApplicableCount() int32 {
	if o == nil || o.NotApplicableCount == nil {
		var ret int32
		return ret
	}
	return *o.NotApplicableCount
}

// GetNotApplicableCountOk returns a tuple with the NotApplicableCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphManagedDeviceMobileAppConfigurationDeviceSummary) GetNotApplicableCountOk() (*int32, bool) {
	if o == nil || o.NotApplicableCount == nil {
		return nil, false
	}
	return o.NotApplicableCount, true
}

// HasNotApplicableCount returns a boolean if a field has been set.
func (o *MicrosoftGraphManagedDeviceMobileAppConfigurationDeviceSummary) HasNotApplicableCount() bool {
	if o != nil && o.NotApplicableCount != nil {
		return true
	}

	return false
}

// SetNotApplicableCount gets a reference to the given int32 and assigns it to the NotApplicableCount field.
func (o *MicrosoftGraphManagedDeviceMobileAppConfigurationDeviceSummary) SetNotApplicableCount(v int32) {
	o.NotApplicableCount = &v
}

// GetPendingCount returns the PendingCount field value if set, zero value otherwise.
func (o *MicrosoftGraphManagedDeviceMobileAppConfigurationDeviceSummary) GetPendingCount() int32 {
	if o == nil || o.PendingCount == nil {
		var ret int32
		return ret
	}
	return *o.PendingCount
}

// GetPendingCountOk returns a tuple with the PendingCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphManagedDeviceMobileAppConfigurationDeviceSummary) GetPendingCountOk() (*int32, bool) {
	if o == nil || o.PendingCount == nil {
		return nil, false
	}
	return o.PendingCount, true
}

// HasPendingCount returns a boolean if a field has been set.
func (o *MicrosoftGraphManagedDeviceMobileAppConfigurationDeviceSummary) HasPendingCount() bool {
	if o != nil && o.PendingCount != nil {
		return true
	}

	return false
}

// SetPendingCount gets a reference to the given int32 and assigns it to the PendingCount field.
func (o *MicrosoftGraphManagedDeviceMobileAppConfigurationDeviceSummary) SetPendingCount(v int32) {
	o.PendingCount = &v
}

// GetSuccessCount returns the SuccessCount field value if set, zero value otherwise.
func (o *MicrosoftGraphManagedDeviceMobileAppConfigurationDeviceSummary) GetSuccessCount() int32 {
	if o == nil || o.SuccessCount == nil {
		var ret int32
		return ret
	}
	return *o.SuccessCount
}

// GetSuccessCountOk returns a tuple with the SuccessCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphManagedDeviceMobileAppConfigurationDeviceSummary) GetSuccessCountOk() (*int32, bool) {
	if o == nil || o.SuccessCount == nil {
		return nil, false
	}
	return o.SuccessCount, true
}

// HasSuccessCount returns a boolean if a field has been set.
func (o *MicrosoftGraphManagedDeviceMobileAppConfigurationDeviceSummary) HasSuccessCount() bool {
	if o != nil && o.SuccessCount != nil {
		return true
	}

	return false
}

// SetSuccessCount gets a reference to the given int32 and assigns it to the SuccessCount field.
func (o *MicrosoftGraphManagedDeviceMobileAppConfigurationDeviceSummary) SetSuccessCount(v int32) {
	o.SuccessCount = &v
}

func (o MicrosoftGraphManagedDeviceMobileAppConfigurationDeviceSummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
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

type NullableMicrosoftGraphManagedDeviceMobileAppConfigurationDeviceSummary struct {
	value *MicrosoftGraphManagedDeviceMobileAppConfigurationDeviceSummary
	isSet bool
}

func (v NullableMicrosoftGraphManagedDeviceMobileAppConfigurationDeviceSummary) Get() *MicrosoftGraphManagedDeviceMobileAppConfigurationDeviceSummary {
	return v.value
}

func (v *NullableMicrosoftGraphManagedDeviceMobileAppConfigurationDeviceSummary) Set(val *MicrosoftGraphManagedDeviceMobileAppConfigurationDeviceSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphManagedDeviceMobileAppConfigurationDeviceSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphManagedDeviceMobileAppConfigurationDeviceSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphManagedDeviceMobileAppConfigurationDeviceSummary(val *MicrosoftGraphManagedDeviceMobileAppConfigurationDeviceSummary) *NullableMicrosoftGraphManagedDeviceMobileAppConfigurationDeviceSummary {
	return &NullableMicrosoftGraphManagedDeviceMobileAppConfigurationDeviceSummary{value: val, isSet: true}
}

func (v NullableMicrosoftGraphManagedDeviceMobileAppConfigurationDeviceSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphManagedDeviceMobileAppConfigurationDeviceSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


