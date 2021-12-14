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

// MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary struct for MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary
type MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
	// Number of compliant devices
	CompliantDeviceCount *int32 `json:"compliantDeviceCount,omitempty"`
	// Number of devices that have compliance managed by System Center Configuration Manager
	ConfigManagerCount *int32 `json:"configManagerCount,omitempty"`
	// Number of conflict devices
	ConflictDeviceCount *int32 `json:"conflictDeviceCount,omitempty"`
	// Number of error devices
	ErrorDeviceCount *int32 `json:"errorDeviceCount,omitempty"`
	// Number of devices that are in grace period
	InGracePeriodCount *int32 `json:"inGracePeriodCount,omitempty"`
	// Number of NonCompliant devices
	NonCompliantDeviceCount *int32 `json:"nonCompliantDeviceCount,omitempty"`
	// Number of not applicable devices
	NotApplicableDeviceCount *int32 `json:"notApplicableDeviceCount,omitempty"`
	// Number of remediated devices
	RemediatedDeviceCount *int32 `json:"remediatedDeviceCount,omitempty"`
	// Number of unknown devices
	UnknownDeviceCount *int32 `json:"unknownDeviceCount,omitempty"`
}

// NewMicrosoftGraphDeviceCompliancePolicyDeviceStateSummary instantiates a new MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphDeviceCompliancePolicyDeviceStateSummary() *MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary {
	this := MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary{}
	return &this
}

// NewMicrosoftGraphDeviceCompliancePolicyDeviceStateSummaryWithDefaults instantiates a new MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphDeviceCompliancePolicyDeviceStateSummaryWithDefaults() *MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary {
	this := MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary) SetId(v string) {
	o.Id = &v
}

// GetCompliantDeviceCount returns the CompliantDeviceCount field value if set, zero value otherwise.
func (o *MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary) GetCompliantDeviceCount() int32 {
	if o == nil || o.CompliantDeviceCount == nil {
		var ret int32
		return ret
	}
	return *o.CompliantDeviceCount
}

// GetCompliantDeviceCountOk returns a tuple with the CompliantDeviceCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary) GetCompliantDeviceCountOk() (*int32, bool) {
	if o == nil || o.CompliantDeviceCount == nil {
		return nil, false
	}
	return o.CompliantDeviceCount, true
}

// HasCompliantDeviceCount returns a boolean if a field has been set.
func (o *MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary) HasCompliantDeviceCount() bool {
	if o != nil && o.CompliantDeviceCount != nil {
		return true
	}

	return false
}

// SetCompliantDeviceCount gets a reference to the given int32 and assigns it to the CompliantDeviceCount field.
func (o *MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary) SetCompliantDeviceCount(v int32) {
	o.CompliantDeviceCount = &v
}

// GetConfigManagerCount returns the ConfigManagerCount field value if set, zero value otherwise.
func (o *MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary) GetConfigManagerCount() int32 {
	if o == nil || o.ConfigManagerCount == nil {
		var ret int32
		return ret
	}
	return *o.ConfigManagerCount
}

// GetConfigManagerCountOk returns a tuple with the ConfigManagerCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary) GetConfigManagerCountOk() (*int32, bool) {
	if o == nil || o.ConfigManagerCount == nil {
		return nil, false
	}
	return o.ConfigManagerCount, true
}

// HasConfigManagerCount returns a boolean if a field has been set.
func (o *MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary) HasConfigManagerCount() bool {
	if o != nil && o.ConfigManagerCount != nil {
		return true
	}

	return false
}

// SetConfigManagerCount gets a reference to the given int32 and assigns it to the ConfigManagerCount field.
func (o *MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary) SetConfigManagerCount(v int32) {
	o.ConfigManagerCount = &v
}

// GetConflictDeviceCount returns the ConflictDeviceCount field value if set, zero value otherwise.
func (o *MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary) GetConflictDeviceCount() int32 {
	if o == nil || o.ConflictDeviceCount == nil {
		var ret int32
		return ret
	}
	return *o.ConflictDeviceCount
}

// GetConflictDeviceCountOk returns a tuple with the ConflictDeviceCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary) GetConflictDeviceCountOk() (*int32, bool) {
	if o == nil || o.ConflictDeviceCount == nil {
		return nil, false
	}
	return o.ConflictDeviceCount, true
}

// HasConflictDeviceCount returns a boolean if a field has been set.
func (o *MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary) HasConflictDeviceCount() bool {
	if o != nil && o.ConflictDeviceCount != nil {
		return true
	}

	return false
}

// SetConflictDeviceCount gets a reference to the given int32 and assigns it to the ConflictDeviceCount field.
func (o *MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary) SetConflictDeviceCount(v int32) {
	o.ConflictDeviceCount = &v
}

// GetErrorDeviceCount returns the ErrorDeviceCount field value if set, zero value otherwise.
func (o *MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary) GetErrorDeviceCount() int32 {
	if o == nil || o.ErrorDeviceCount == nil {
		var ret int32
		return ret
	}
	return *o.ErrorDeviceCount
}

// GetErrorDeviceCountOk returns a tuple with the ErrorDeviceCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary) GetErrorDeviceCountOk() (*int32, bool) {
	if o == nil || o.ErrorDeviceCount == nil {
		return nil, false
	}
	return o.ErrorDeviceCount, true
}

// HasErrorDeviceCount returns a boolean if a field has been set.
func (o *MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary) HasErrorDeviceCount() bool {
	if o != nil && o.ErrorDeviceCount != nil {
		return true
	}

	return false
}

// SetErrorDeviceCount gets a reference to the given int32 and assigns it to the ErrorDeviceCount field.
func (o *MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary) SetErrorDeviceCount(v int32) {
	o.ErrorDeviceCount = &v
}

// GetInGracePeriodCount returns the InGracePeriodCount field value if set, zero value otherwise.
func (o *MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary) GetInGracePeriodCount() int32 {
	if o == nil || o.InGracePeriodCount == nil {
		var ret int32
		return ret
	}
	return *o.InGracePeriodCount
}

// GetInGracePeriodCountOk returns a tuple with the InGracePeriodCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary) GetInGracePeriodCountOk() (*int32, bool) {
	if o == nil || o.InGracePeriodCount == nil {
		return nil, false
	}
	return o.InGracePeriodCount, true
}

// HasInGracePeriodCount returns a boolean if a field has been set.
func (o *MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary) HasInGracePeriodCount() bool {
	if o != nil && o.InGracePeriodCount != nil {
		return true
	}

	return false
}

// SetInGracePeriodCount gets a reference to the given int32 and assigns it to the InGracePeriodCount field.
func (o *MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary) SetInGracePeriodCount(v int32) {
	o.InGracePeriodCount = &v
}

// GetNonCompliantDeviceCount returns the NonCompliantDeviceCount field value if set, zero value otherwise.
func (o *MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary) GetNonCompliantDeviceCount() int32 {
	if o == nil || o.NonCompliantDeviceCount == nil {
		var ret int32
		return ret
	}
	return *o.NonCompliantDeviceCount
}

// GetNonCompliantDeviceCountOk returns a tuple with the NonCompliantDeviceCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary) GetNonCompliantDeviceCountOk() (*int32, bool) {
	if o == nil || o.NonCompliantDeviceCount == nil {
		return nil, false
	}
	return o.NonCompliantDeviceCount, true
}

// HasNonCompliantDeviceCount returns a boolean if a field has been set.
func (o *MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary) HasNonCompliantDeviceCount() bool {
	if o != nil && o.NonCompliantDeviceCount != nil {
		return true
	}

	return false
}

// SetNonCompliantDeviceCount gets a reference to the given int32 and assigns it to the NonCompliantDeviceCount field.
func (o *MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary) SetNonCompliantDeviceCount(v int32) {
	o.NonCompliantDeviceCount = &v
}

// GetNotApplicableDeviceCount returns the NotApplicableDeviceCount field value if set, zero value otherwise.
func (o *MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary) GetNotApplicableDeviceCount() int32 {
	if o == nil || o.NotApplicableDeviceCount == nil {
		var ret int32
		return ret
	}
	return *o.NotApplicableDeviceCount
}

// GetNotApplicableDeviceCountOk returns a tuple with the NotApplicableDeviceCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary) GetNotApplicableDeviceCountOk() (*int32, bool) {
	if o == nil || o.NotApplicableDeviceCount == nil {
		return nil, false
	}
	return o.NotApplicableDeviceCount, true
}

// HasNotApplicableDeviceCount returns a boolean if a field has been set.
func (o *MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary) HasNotApplicableDeviceCount() bool {
	if o != nil && o.NotApplicableDeviceCount != nil {
		return true
	}

	return false
}

// SetNotApplicableDeviceCount gets a reference to the given int32 and assigns it to the NotApplicableDeviceCount field.
func (o *MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary) SetNotApplicableDeviceCount(v int32) {
	o.NotApplicableDeviceCount = &v
}

// GetRemediatedDeviceCount returns the RemediatedDeviceCount field value if set, zero value otherwise.
func (o *MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary) GetRemediatedDeviceCount() int32 {
	if o == nil || o.RemediatedDeviceCount == nil {
		var ret int32
		return ret
	}
	return *o.RemediatedDeviceCount
}

// GetRemediatedDeviceCountOk returns a tuple with the RemediatedDeviceCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary) GetRemediatedDeviceCountOk() (*int32, bool) {
	if o == nil || o.RemediatedDeviceCount == nil {
		return nil, false
	}
	return o.RemediatedDeviceCount, true
}

// HasRemediatedDeviceCount returns a boolean if a field has been set.
func (o *MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary) HasRemediatedDeviceCount() bool {
	if o != nil && o.RemediatedDeviceCount != nil {
		return true
	}

	return false
}

// SetRemediatedDeviceCount gets a reference to the given int32 and assigns it to the RemediatedDeviceCount field.
func (o *MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary) SetRemediatedDeviceCount(v int32) {
	o.RemediatedDeviceCount = &v
}

// GetUnknownDeviceCount returns the UnknownDeviceCount field value if set, zero value otherwise.
func (o *MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary) GetUnknownDeviceCount() int32 {
	if o == nil || o.UnknownDeviceCount == nil {
		var ret int32
		return ret
	}
	return *o.UnknownDeviceCount
}

// GetUnknownDeviceCountOk returns a tuple with the UnknownDeviceCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary) GetUnknownDeviceCountOk() (*int32, bool) {
	if o == nil || o.UnknownDeviceCount == nil {
		return nil, false
	}
	return o.UnknownDeviceCount, true
}

// HasUnknownDeviceCount returns a boolean if a field has been set.
func (o *MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary) HasUnknownDeviceCount() bool {
	if o != nil && o.UnknownDeviceCount != nil {
		return true
	}

	return false
}

// SetUnknownDeviceCount gets a reference to the given int32 and assigns it to the UnknownDeviceCount field.
func (o *MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary) SetUnknownDeviceCount(v int32) {
	o.UnknownDeviceCount = &v
}

func (o MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.CompliantDeviceCount != nil {
		toSerialize["compliantDeviceCount"] = o.CompliantDeviceCount
	}
	if o.ConfigManagerCount != nil {
		toSerialize["configManagerCount"] = o.ConfigManagerCount
	}
	if o.ConflictDeviceCount != nil {
		toSerialize["conflictDeviceCount"] = o.ConflictDeviceCount
	}
	if o.ErrorDeviceCount != nil {
		toSerialize["errorDeviceCount"] = o.ErrorDeviceCount
	}
	if o.InGracePeriodCount != nil {
		toSerialize["inGracePeriodCount"] = o.InGracePeriodCount
	}
	if o.NonCompliantDeviceCount != nil {
		toSerialize["nonCompliantDeviceCount"] = o.NonCompliantDeviceCount
	}
	if o.NotApplicableDeviceCount != nil {
		toSerialize["notApplicableDeviceCount"] = o.NotApplicableDeviceCount
	}
	if o.RemediatedDeviceCount != nil {
		toSerialize["remediatedDeviceCount"] = o.RemediatedDeviceCount
	}
	if o.UnknownDeviceCount != nil {
		toSerialize["unknownDeviceCount"] = o.UnknownDeviceCount
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphDeviceCompliancePolicyDeviceStateSummary struct {
	value *MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary
	isSet bool
}

func (v NullableMicrosoftGraphDeviceCompliancePolicyDeviceStateSummary) Get() *MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary {
	return v.value
}

func (v *NullableMicrosoftGraphDeviceCompliancePolicyDeviceStateSummary) Set(val *MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphDeviceCompliancePolicyDeviceStateSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphDeviceCompliancePolicyDeviceStateSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphDeviceCompliancePolicyDeviceStateSummary(val *MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary) *NullableMicrosoftGraphDeviceCompliancePolicyDeviceStateSummary {
	return &NullableMicrosoftGraphDeviceCompliancePolicyDeviceStateSummary{value: val, isSet: true}
}

func (v NullableMicrosoftGraphDeviceCompliancePolicyDeviceStateSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphDeviceCompliancePolicyDeviceStateSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


