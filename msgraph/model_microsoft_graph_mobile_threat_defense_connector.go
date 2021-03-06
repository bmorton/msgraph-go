/*
Partial Graph API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package msgraph

import (
	"encoding/json"
	"time"
)

// MicrosoftGraphMobileThreatDefenseConnector struct for MicrosoftGraphMobileThreatDefenseConnector
type MicrosoftGraphMobileThreatDefenseConnector struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
	// For Android, set whether Intune must receive data from the data sync partner prior to marking a device compliant
	AndroidDeviceBlockedOnMissingPartnerData *bool `json:"androidDeviceBlockedOnMissingPartnerData,omitempty"`
	// For Android, set whether data from the data sync partner should be used during compliance evaluations
	AndroidEnabled *bool `json:"androidEnabled,omitempty"`
	// For IOS, set whether Intune must receive data from the data sync partner prior to marking a device compliant
	IosDeviceBlockedOnMissingPartnerData *bool `json:"iosDeviceBlockedOnMissingPartnerData,omitempty"`
	// For IOS, get or set whether data from the data sync partner should be used during compliance evaluations
	IosEnabled *bool `json:"iosEnabled,omitempty"`
	// DateTime of last Heartbeat recieved from the Data Sync Partner
	LastHeartbeatDateTime *time.Time `json:"lastHeartbeatDateTime,omitempty"`
	// Data Sync Partner state for this account. Possible values are: unavailable, available, enabled, unresponsive.
	PartnerState AnyOfmicrosoftGraphMobileThreatPartnerTenantState `json:"partnerState,omitempty"`
	// Get or Set days the per tenant tolerance to unresponsiveness for this partner integration
	PartnerUnresponsivenessThresholdInDays *int32 `json:"partnerUnresponsivenessThresholdInDays,omitempty"`
	// Get or set whether to block devices on the enabled platforms that do not meet the minimum version requirements of the Data Sync Partner
	PartnerUnsupportedOsVersionBlocked *bool `json:"partnerUnsupportedOsVersionBlocked,omitempty"`
}

// NewMicrosoftGraphMobileThreatDefenseConnector instantiates a new MicrosoftGraphMobileThreatDefenseConnector object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphMobileThreatDefenseConnector() *MicrosoftGraphMobileThreatDefenseConnector {
	this := MicrosoftGraphMobileThreatDefenseConnector{}
	return &this
}

// NewMicrosoftGraphMobileThreatDefenseConnectorWithDefaults instantiates a new MicrosoftGraphMobileThreatDefenseConnector object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphMobileThreatDefenseConnectorWithDefaults() *MicrosoftGraphMobileThreatDefenseConnector {
	this := MicrosoftGraphMobileThreatDefenseConnector{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MicrosoftGraphMobileThreatDefenseConnector) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphMobileThreatDefenseConnector) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphMobileThreatDefenseConnector) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MicrosoftGraphMobileThreatDefenseConnector) SetId(v string) {
	o.Id = &v
}

// GetAndroidDeviceBlockedOnMissingPartnerData returns the AndroidDeviceBlockedOnMissingPartnerData field value if set, zero value otherwise.
func (o *MicrosoftGraphMobileThreatDefenseConnector) GetAndroidDeviceBlockedOnMissingPartnerData() bool {
	if o == nil || o.AndroidDeviceBlockedOnMissingPartnerData == nil {
		var ret bool
		return ret
	}
	return *o.AndroidDeviceBlockedOnMissingPartnerData
}

// GetAndroidDeviceBlockedOnMissingPartnerDataOk returns a tuple with the AndroidDeviceBlockedOnMissingPartnerData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphMobileThreatDefenseConnector) GetAndroidDeviceBlockedOnMissingPartnerDataOk() (*bool, bool) {
	if o == nil || o.AndroidDeviceBlockedOnMissingPartnerData == nil {
		return nil, false
	}
	return o.AndroidDeviceBlockedOnMissingPartnerData, true
}

// HasAndroidDeviceBlockedOnMissingPartnerData returns a boolean if a field has been set.
func (o *MicrosoftGraphMobileThreatDefenseConnector) HasAndroidDeviceBlockedOnMissingPartnerData() bool {
	if o != nil && o.AndroidDeviceBlockedOnMissingPartnerData != nil {
		return true
	}

	return false
}

// SetAndroidDeviceBlockedOnMissingPartnerData gets a reference to the given bool and assigns it to the AndroidDeviceBlockedOnMissingPartnerData field.
func (o *MicrosoftGraphMobileThreatDefenseConnector) SetAndroidDeviceBlockedOnMissingPartnerData(v bool) {
	o.AndroidDeviceBlockedOnMissingPartnerData = &v
}

// GetAndroidEnabled returns the AndroidEnabled field value if set, zero value otherwise.
func (o *MicrosoftGraphMobileThreatDefenseConnector) GetAndroidEnabled() bool {
	if o == nil || o.AndroidEnabled == nil {
		var ret bool
		return ret
	}
	return *o.AndroidEnabled
}

// GetAndroidEnabledOk returns a tuple with the AndroidEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphMobileThreatDefenseConnector) GetAndroidEnabledOk() (*bool, bool) {
	if o == nil || o.AndroidEnabled == nil {
		return nil, false
	}
	return o.AndroidEnabled, true
}

// HasAndroidEnabled returns a boolean if a field has been set.
func (o *MicrosoftGraphMobileThreatDefenseConnector) HasAndroidEnabled() bool {
	if o != nil && o.AndroidEnabled != nil {
		return true
	}

	return false
}

// SetAndroidEnabled gets a reference to the given bool and assigns it to the AndroidEnabled field.
func (o *MicrosoftGraphMobileThreatDefenseConnector) SetAndroidEnabled(v bool) {
	o.AndroidEnabled = &v
}

// GetIosDeviceBlockedOnMissingPartnerData returns the IosDeviceBlockedOnMissingPartnerData field value if set, zero value otherwise.
func (o *MicrosoftGraphMobileThreatDefenseConnector) GetIosDeviceBlockedOnMissingPartnerData() bool {
	if o == nil || o.IosDeviceBlockedOnMissingPartnerData == nil {
		var ret bool
		return ret
	}
	return *o.IosDeviceBlockedOnMissingPartnerData
}

// GetIosDeviceBlockedOnMissingPartnerDataOk returns a tuple with the IosDeviceBlockedOnMissingPartnerData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphMobileThreatDefenseConnector) GetIosDeviceBlockedOnMissingPartnerDataOk() (*bool, bool) {
	if o == nil || o.IosDeviceBlockedOnMissingPartnerData == nil {
		return nil, false
	}
	return o.IosDeviceBlockedOnMissingPartnerData, true
}

// HasIosDeviceBlockedOnMissingPartnerData returns a boolean if a field has been set.
func (o *MicrosoftGraphMobileThreatDefenseConnector) HasIosDeviceBlockedOnMissingPartnerData() bool {
	if o != nil && o.IosDeviceBlockedOnMissingPartnerData != nil {
		return true
	}

	return false
}

// SetIosDeviceBlockedOnMissingPartnerData gets a reference to the given bool and assigns it to the IosDeviceBlockedOnMissingPartnerData field.
func (o *MicrosoftGraphMobileThreatDefenseConnector) SetIosDeviceBlockedOnMissingPartnerData(v bool) {
	o.IosDeviceBlockedOnMissingPartnerData = &v
}

// GetIosEnabled returns the IosEnabled field value if set, zero value otherwise.
func (o *MicrosoftGraphMobileThreatDefenseConnector) GetIosEnabled() bool {
	if o == nil || o.IosEnabled == nil {
		var ret bool
		return ret
	}
	return *o.IosEnabled
}

// GetIosEnabledOk returns a tuple with the IosEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphMobileThreatDefenseConnector) GetIosEnabledOk() (*bool, bool) {
	if o == nil || o.IosEnabled == nil {
		return nil, false
	}
	return o.IosEnabled, true
}

// HasIosEnabled returns a boolean if a field has been set.
func (o *MicrosoftGraphMobileThreatDefenseConnector) HasIosEnabled() bool {
	if o != nil && o.IosEnabled != nil {
		return true
	}

	return false
}

// SetIosEnabled gets a reference to the given bool and assigns it to the IosEnabled field.
func (o *MicrosoftGraphMobileThreatDefenseConnector) SetIosEnabled(v bool) {
	o.IosEnabled = &v
}

// GetLastHeartbeatDateTime returns the LastHeartbeatDateTime field value if set, zero value otherwise.
func (o *MicrosoftGraphMobileThreatDefenseConnector) GetLastHeartbeatDateTime() time.Time {
	if o == nil || o.LastHeartbeatDateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.LastHeartbeatDateTime
}

// GetLastHeartbeatDateTimeOk returns a tuple with the LastHeartbeatDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphMobileThreatDefenseConnector) GetLastHeartbeatDateTimeOk() (*time.Time, bool) {
	if o == nil || o.LastHeartbeatDateTime == nil {
		return nil, false
	}
	return o.LastHeartbeatDateTime, true
}

// HasLastHeartbeatDateTime returns a boolean if a field has been set.
func (o *MicrosoftGraphMobileThreatDefenseConnector) HasLastHeartbeatDateTime() bool {
	if o != nil && o.LastHeartbeatDateTime != nil {
		return true
	}

	return false
}

// SetLastHeartbeatDateTime gets a reference to the given time.Time and assigns it to the LastHeartbeatDateTime field.
func (o *MicrosoftGraphMobileThreatDefenseConnector) SetLastHeartbeatDateTime(v time.Time) {
	o.LastHeartbeatDateTime = &v
}

// GetPartnerState returns the PartnerState field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphMobileThreatDefenseConnector) GetPartnerState() AnyOfmicrosoftGraphMobileThreatPartnerTenantState {
	if o == nil  {
		var ret AnyOfmicrosoftGraphMobileThreatPartnerTenantState
		return ret
	}
	return o.PartnerState
}

// GetPartnerStateOk returns a tuple with the PartnerState field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphMobileThreatDefenseConnector) GetPartnerStateOk() (*AnyOfmicrosoftGraphMobileThreatPartnerTenantState, bool) {
	if o == nil || o.PartnerState == nil {
		return nil, false
	}
	return &o.PartnerState, true
}

// HasPartnerState returns a boolean if a field has been set.
func (o *MicrosoftGraphMobileThreatDefenseConnector) HasPartnerState() bool {
	if o != nil && o.PartnerState != nil {
		return true
	}

	return false
}

// SetPartnerState gets a reference to the given AnyOfmicrosoftGraphMobileThreatPartnerTenantState and assigns it to the PartnerState field.
func (o *MicrosoftGraphMobileThreatDefenseConnector) SetPartnerState(v AnyOfmicrosoftGraphMobileThreatPartnerTenantState) {
	o.PartnerState = v
}

// GetPartnerUnresponsivenessThresholdInDays returns the PartnerUnresponsivenessThresholdInDays field value if set, zero value otherwise.
func (o *MicrosoftGraphMobileThreatDefenseConnector) GetPartnerUnresponsivenessThresholdInDays() int32 {
	if o == nil || o.PartnerUnresponsivenessThresholdInDays == nil {
		var ret int32
		return ret
	}
	return *o.PartnerUnresponsivenessThresholdInDays
}

// GetPartnerUnresponsivenessThresholdInDaysOk returns a tuple with the PartnerUnresponsivenessThresholdInDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphMobileThreatDefenseConnector) GetPartnerUnresponsivenessThresholdInDaysOk() (*int32, bool) {
	if o == nil || o.PartnerUnresponsivenessThresholdInDays == nil {
		return nil, false
	}
	return o.PartnerUnresponsivenessThresholdInDays, true
}

// HasPartnerUnresponsivenessThresholdInDays returns a boolean if a field has been set.
func (o *MicrosoftGraphMobileThreatDefenseConnector) HasPartnerUnresponsivenessThresholdInDays() bool {
	if o != nil && o.PartnerUnresponsivenessThresholdInDays != nil {
		return true
	}

	return false
}

// SetPartnerUnresponsivenessThresholdInDays gets a reference to the given int32 and assigns it to the PartnerUnresponsivenessThresholdInDays field.
func (o *MicrosoftGraphMobileThreatDefenseConnector) SetPartnerUnresponsivenessThresholdInDays(v int32) {
	o.PartnerUnresponsivenessThresholdInDays = &v
}

// GetPartnerUnsupportedOsVersionBlocked returns the PartnerUnsupportedOsVersionBlocked field value if set, zero value otherwise.
func (o *MicrosoftGraphMobileThreatDefenseConnector) GetPartnerUnsupportedOsVersionBlocked() bool {
	if o == nil || o.PartnerUnsupportedOsVersionBlocked == nil {
		var ret bool
		return ret
	}
	return *o.PartnerUnsupportedOsVersionBlocked
}

// GetPartnerUnsupportedOsVersionBlockedOk returns a tuple with the PartnerUnsupportedOsVersionBlocked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphMobileThreatDefenseConnector) GetPartnerUnsupportedOsVersionBlockedOk() (*bool, bool) {
	if o == nil || o.PartnerUnsupportedOsVersionBlocked == nil {
		return nil, false
	}
	return o.PartnerUnsupportedOsVersionBlocked, true
}

// HasPartnerUnsupportedOsVersionBlocked returns a boolean if a field has been set.
func (o *MicrosoftGraphMobileThreatDefenseConnector) HasPartnerUnsupportedOsVersionBlocked() bool {
	if o != nil && o.PartnerUnsupportedOsVersionBlocked != nil {
		return true
	}

	return false
}

// SetPartnerUnsupportedOsVersionBlocked gets a reference to the given bool and assigns it to the PartnerUnsupportedOsVersionBlocked field.
func (o *MicrosoftGraphMobileThreatDefenseConnector) SetPartnerUnsupportedOsVersionBlocked(v bool) {
	o.PartnerUnsupportedOsVersionBlocked = &v
}

func (o MicrosoftGraphMobileThreatDefenseConnector) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.AndroidDeviceBlockedOnMissingPartnerData != nil {
		toSerialize["androidDeviceBlockedOnMissingPartnerData"] = o.AndroidDeviceBlockedOnMissingPartnerData
	}
	if o.AndroidEnabled != nil {
		toSerialize["androidEnabled"] = o.AndroidEnabled
	}
	if o.IosDeviceBlockedOnMissingPartnerData != nil {
		toSerialize["iosDeviceBlockedOnMissingPartnerData"] = o.IosDeviceBlockedOnMissingPartnerData
	}
	if o.IosEnabled != nil {
		toSerialize["iosEnabled"] = o.IosEnabled
	}
	if o.LastHeartbeatDateTime != nil {
		toSerialize["lastHeartbeatDateTime"] = o.LastHeartbeatDateTime
	}
	if o.PartnerState != nil {
		toSerialize["partnerState"] = o.PartnerState
	}
	if o.PartnerUnresponsivenessThresholdInDays != nil {
		toSerialize["partnerUnresponsivenessThresholdInDays"] = o.PartnerUnresponsivenessThresholdInDays
	}
	if o.PartnerUnsupportedOsVersionBlocked != nil {
		toSerialize["partnerUnsupportedOsVersionBlocked"] = o.PartnerUnsupportedOsVersionBlocked
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphMobileThreatDefenseConnector struct {
	value *MicrosoftGraphMobileThreatDefenseConnector
	isSet bool
}

func (v NullableMicrosoftGraphMobileThreatDefenseConnector) Get() *MicrosoftGraphMobileThreatDefenseConnector {
	return v.value
}

func (v *NullableMicrosoftGraphMobileThreatDefenseConnector) Set(val *MicrosoftGraphMobileThreatDefenseConnector) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphMobileThreatDefenseConnector) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphMobileThreatDefenseConnector) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphMobileThreatDefenseConnector(val *MicrosoftGraphMobileThreatDefenseConnector) *NullableMicrosoftGraphMobileThreatDefenseConnector {
	return &NullableMicrosoftGraphMobileThreatDefenseConnector{value: val, isSet: true}
}

func (v NullableMicrosoftGraphMobileThreatDefenseConnector) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphMobileThreatDefenseConnector) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


