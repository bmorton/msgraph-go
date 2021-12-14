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

// MicrosoftGraphSoftwareUpdateStatusSummary struct for MicrosoftGraphSoftwareUpdateStatusSummary
type MicrosoftGraphSoftwareUpdateStatusSummary struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
	// Number of compliant devices.
	CompliantDeviceCount *int32 `json:"compliantDeviceCount,omitempty"`
	// Number of compliant users.
	CompliantUserCount *int32 `json:"compliantUserCount,omitempty"`
	// Number of conflict devices.
	ConflictDeviceCount *int32 `json:"conflictDeviceCount,omitempty"`
	// Number of conflict users.
	ConflictUserCount *int32 `json:"conflictUserCount,omitempty"`
	// The name of the policy.
	DisplayName NullableString `json:"displayName,omitempty"`
	// Number of devices had error.
	ErrorDeviceCount *int32 `json:"errorDeviceCount,omitempty"`
	// Number of users had error.
	ErrorUserCount *int32 `json:"errorUserCount,omitempty"`
	// Number of non compliant devices.
	NonCompliantDeviceCount *int32 `json:"nonCompliantDeviceCount,omitempty"`
	// Number of non compliant users.
	NonCompliantUserCount *int32 `json:"nonCompliantUserCount,omitempty"`
	// Number of not applicable devices.
	NotApplicableDeviceCount *int32 `json:"notApplicableDeviceCount,omitempty"`
	// Number of not applicable users.
	NotApplicableUserCount *int32 `json:"notApplicableUserCount,omitempty"`
	// Number of remediated devices.
	RemediatedDeviceCount *int32 `json:"remediatedDeviceCount,omitempty"`
	// Number of remediated users.
	RemediatedUserCount *int32 `json:"remediatedUserCount,omitempty"`
	// Number of unknown devices.
	UnknownDeviceCount *int32 `json:"unknownDeviceCount,omitempty"`
	// Number of unknown users.
	UnknownUserCount *int32 `json:"unknownUserCount,omitempty"`
}

// NewMicrosoftGraphSoftwareUpdateStatusSummary instantiates a new MicrosoftGraphSoftwareUpdateStatusSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphSoftwareUpdateStatusSummary() *MicrosoftGraphSoftwareUpdateStatusSummary {
	this := MicrosoftGraphSoftwareUpdateStatusSummary{}
	return &this
}

// NewMicrosoftGraphSoftwareUpdateStatusSummaryWithDefaults instantiates a new MicrosoftGraphSoftwareUpdateStatusSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphSoftwareUpdateStatusSummaryWithDefaults() *MicrosoftGraphSoftwareUpdateStatusSummary {
	this := MicrosoftGraphSoftwareUpdateStatusSummary{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MicrosoftGraphSoftwareUpdateStatusSummary) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphSoftwareUpdateStatusSummary) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphSoftwareUpdateStatusSummary) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MicrosoftGraphSoftwareUpdateStatusSummary) SetId(v string) {
	o.Id = &v
}

// GetCompliantDeviceCount returns the CompliantDeviceCount field value if set, zero value otherwise.
func (o *MicrosoftGraphSoftwareUpdateStatusSummary) GetCompliantDeviceCount() int32 {
	if o == nil || o.CompliantDeviceCount == nil {
		var ret int32
		return ret
	}
	return *o.CompliantDeviceCount
}

// GetCompliantDeviceCountOk returns a tuple with the CompliantDeviceCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphSoftwareUpdateStatusSummary) GetCompliantDeviceCountOk() (*int32, bool) {
	if o == nil || o.CompliantDeviceCount == nil {
		return nil, false
	}
	return o.CompliantDeviceCount, true
}

// HasCompliantDeviceCount returns a boolean if a field has been set.
func (o *MicrosoftGraphSoftwareUpdateStatusSummary) HasCompliantDeviceCount() bool {
	if o != nil && o.CompliantDeviceCount != nil {
		return true
	}

	return false
}

// SetCompliantDeviceCount gets a reference to the given int32 and assigns it to the CompliantDeviceCount field.
func (o *MicrosoftGraphSoftwareUpdateStatusSummary) SetCompliantDeviceCount(v int32) {
	o.CompliantDeviceCount = &v
}

// GetCompliantUserCount returns the CompliantUserCount field value if set, zero value otherwise.
func (o *MicrosoftGraphSoftwareUpdateStatusSummary) GetCompliantUserCount() int32 {
	if o == nil || o.CompliantUserCount == nil {
		var ret int32
		return ret
	}
	return *o.CompliantUserCount
}

// GetCompliantUserCountOk returns a tuple with the CompliantUserCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphSoftwareUpdateStatusSummary) GetCompliantUserCountOk() (*int32, bool) {
	if o == nil || o.CompliantUserCount == nil {
		return nil, false
	}
	return o.CompliantUserCount, true
}

// HasCompliantUserCount returns a boolean if a field has been set.
func (o *MicrosoftGraphSoftwareUpdateStatusSummary) HasCompliantUserCount() bool {
	if o != nil && o.CompliantUserCount != nil {
		return true
	}

	return false
}

// SetCompliantUserCount gets a reference to the given int32 and assigns it to the CompliantUserCount field.
func (o *MicrosoftGraphSoftwareUpdateStatusSummary) SetCompliantUserCount(v int32) {
	o.CompliantUserCount = &v
}

// GetConflictDeviceCount returns the ConflictDeviceCount field value if set, zero value otherwise.
func (o *MicrosoftGraphSoftwareUpdateStatusSummary) GetConflictDeviceCount() int32 {
	if o == nil || o.ConflictDeviceCount == nil {
		var ret int32
		return ret
	}
	return *o.ConflictDeviceCount
}

// GetConflictDeviceCountOk returns a tuple with the ConflictDeviceCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphSoftwareUpdateStatusSummary) GetConflictDeviceCountOk() (*int32, bool) {
	if o == nil || o.ConflictDeviceCount == nil {
		return nil, false
	}
	return o.ConflictDeviceCount, true
}

// HasConflictDeviceCount returns a boolean if a field has been set.
func (o *MicrosoftGraphSoftwareUpdateStatusSummary) HasConflictDeviceCount() bool {
	if o != nil && o.ConflictDeviceCount != nil {
		return true
	}

	return false
}

// SetConflictDeviceCount gets a reference to the given int32 and assigns it to the ConflictDeviceCount field.
func (o *MicrosoftGraphSoftwareUpdateStatusSummary) SetConflictDeviceCount(v int32) {
	o.ConflictDeviceCount = &v
}

// GetConflictUserCount returns the ConflictUserCount field value if set, zero value otherwise.
func (o *MicrosoftGraphSoftwareUpdateStatusSummary) GetConflictUserCount() int32 {
	if o == nil || o.ConflictUserCount == nil {
		var ret int32
		return ret
	}
	return *o.ConflictUserCount
}

// GetConflictUserCountOk returns a tuple with the ConflictUserCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphSoftwareUpdateStatusSummary) GetConflictUserCountOk() (*int32, bool) {
	if o == nil || o.ConflictUserCount == nil {
		return nil, false
	}
	return o.ConflictUserCount, true
}

// HasConflictUserCount returns a boolean if a field has been set.
func (o *MicrosoftGraphSoftwareUpdateStatusSummary) HasConflictUserCount() bool {
	if o != nil && o.ConflictUserCount != nil {
		return true
	}

	return false
}

// SetConflictUserCount gets a reference to the given int32 and assigns it to the ConflictUserCount field.
func (o *MicrosoftGraphSoftwareUpdateStatusSummary) SetConflictUserCount(v int32) {
	o.ConflictUserCount = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphSoftwareUpdateStatusSummary) GetDisplayName() string {
	if o == nil || o.DisplayName.Get() == nil {
		var ret string
		return ret
	}
	return *o.DisplayName.Get()
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphSoftwareUpdateStatusSummary) GetDisplayNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DisplayName.Get(), o.DisplayName.IsSet()
}

// HasDisplayName returns a boolean if a field has been set.
func (o *MicrosoftGraphSoftwareUpdateStatusSummary) HasDisplayName() bool {
	if o != nil && o.DisplayName.IsSet() {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given NullableString and assigns it to the DisplayName field.
func (o *MicrosoftGraphSoftwareUpdateStatusSummary) SetDisplayName(v string) {
	o.DisplayName.Set(&v)
}
// SetDisplayNameNil sets the value for DisplayName to be an explicit nil
func (o *MicrosoftGraphSoftwareUpdateStatusSummary) SetDisplayNameNil() {
	o.DisplayName.Set(nil)
}

// UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
func (o *MicrosoftGraphSoftwareUpdateStatusSummary) UnsetDisplayName() {
	o.DisplayName.Unset()
}

// GetErrorDeviceCount returns the ErrorDeviceCount field value if set, zero value otherwise.
func (o *MicrosoftGraphSoftwareUpdateStatusSummary) GetErrorDeviceCount() int32 {
	if o == nil || o.ErrorDeviceCount == nil {
		var ret int32
		return ret
	}
	return *o.ErrorDeviceCount
}

// GetErrorDeviceCountOk returns a tuple with the ErrorDeviceCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphSoftwareUpdateStatusSummary) GetErrorDeviceCountOk() (*int32, bool) {
	if o == nil || o.ErrorDeviceCount == nil {
		return nil, false
	}
	return o.ErrorDeviceCount, true
}

// HasErrorDeviceCount returns a boolean if a field has been set.
func (o *MicrosoftGraphSoftwareUpdateStatusSummary) HasErrorDeviceCount() bool {
	if o != nil && o.ErrorDeviceCount != nil {
		return true
	}

	return false
}

// SetErrorDeviceCount gets a reference to the given int32 and assigns it to the ErrorDeviceCount field.
func (o *MicrosoftGraphSoftwareUpdateStatusSummary) SetErrorDeviceCount(v int32) {
	o.ErrorDeviceCount = &v
}

// GetErrorUserCount returns the ErrorUserCount field value if set, zero value otherwise.
func (o *MicrosoftGraphSoftwareUpdateStatusSummary) GetErrorUserCount() int32 {
	if o == nil || o.ErrorUserCount == nil {
		var ret int32
		return ret
	}
	return *o.ErrorUserCount
}

// GetErrorUserCountOk returns a tuple with the ErrorUserCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphSoftwareUpdateStatusSummary) GetErrorUserCountOk() (*int32, bool) {
	if o == nil || o.ErrorUserCount == nil {
		return nil, false
	}
	return o.ErrorUserCount, true
}

// HasErrorUserCount returns a boolean if a field has been set.
func (o *MicrosoftGraphSoftwareUpdateStatusSummary) HasErrorUserCount() bool {
	if o != nil && o.ErrorUserCount != nil {
		return true
	}

	return false
}

// SetErrorUserCount gets a reference to the given int32 and assigns it to the ErrorUserCount field.
func (o *MicrosoftGraphSoftwareUpdateStatusSummary) SetErrorUserCount(v int32) {
	o.ErrorUserCount = &v
}

// GetNonCompliantDeviceCount returns the NonCompliantDeviceCount field value if set, zero value otherwise.
func (o *MicrosoftGraphSoftwareUpdateStatusSummary) GetNonCompliantDeviceCount() int32 {
	if o == nil || o.NonCompliantDeviceCount == nil {
		var ret int32
		return ret
	}
	return *o.NonCompliantDeviceCount
}

// GetNonCompliantDeviceCountOk returns a tuple with the NonCompliantDeviceCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphSoftwareUpdateStatusSummary) GetNonCompliantDeviceCountOk() (*int32, bool) {
	if o == nil || o.NonCompliantDeviceCount == nil {
		return nil, false
	}
	return o.NonCompliantDeviceCount, true
}

// HasNonCompliantDeviceCount returns a boolean if a field has been set.
func (o *MicrosoftGraphSoftwareUpdateStatusSummary) HasNonCompliantDeviceCount() bool {
	if o != nil && o.NonCompliantDeviceCount != nil {
		return true
	}

	return false
}

// SetNonCompliantDeviceCount gets a reference to the given int32 and assigns it to the NonCompliantDeviceCount field.
func (o *MicrosoftGraphSoftwareUpdateStatusSummary) SetNonCompliantDeviceCount(v int32) {
	o.NonCompliantDeviceCount = &v
}

// GetNonCompliantUserCount returns the NonCompliantUserCount field value if set, zero value otherwise.
func (o *MicrosoftGraphSoftwareUpdateStatusSummary) GetNonCompliantUserCount() int32 {
	if o == nil || o.NonCompliantUserCount == nil {
		var ret int32
		return ret
	}
	return *o.NonCompliantUserCount
}

// GetNonCompliantUserCountOk returns a tuple with the NonCompliantUserCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphSoftwareUpdateStatusSummary) GetNonCompliantUserCountOk() (*int32, bool) {
	if o == nil || o.NonCompliantUserCount == nil {
		return nil, false
	}
	return o.NonCompliantUserCount, true
}

// HasNonCompliantUserCount returns a boolean if a field has been set.
func (o *MicrosoftGraphSoftwareUpdateStatusSummary) HasNonCompliantUserCount() bool {
	if o != nil && o.NonCompliantUserCount != nil {
		return true
	}

	return false
}

// SetNonCompliantUserCount gets a reference to the given int32 and assigns it to the NonCompliantUserCount field.
func (o *MicrosoftGraphSoftwareUpdateStatusSummary) SetNonCompliantUserCount(v int32) {
	o.NonCompliantUserCount = &v
}

// GetNotApplicableDeviceCount returns the NotApplicableDeviceCount field value if set, zero value otherwise.
func (o *MicrosoftGraphSoftwareUpdateStatusSummary) GetNotApplicableDeviceCount() int32 {
	if o == nil || o.NotApplicableDeviceCount == nil {
		var ret int32
		return ret
	}
	return *o.NotApplicableDeviceCount
}

// GetNotApplicableDeviceCountOk returns a tuple with the NotApplicableDeviceCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphSoftwareUpdateStatusSummary) GetNotApplicableDeviceCountOk() (*int32, bool) {
	if o == nil || o.NotApplicableDeviceCount == nil {
		return nil, false
	}
	return o.NotApplicableDeviceCount, true
}

// HasNotApplicableDeviceCount returns a boolean if a field has been set.
func (o *MicrosoftGraphSoftwareUpdateStatusSummary) HasNotApplicableDeviceCount() bool {
	if o != nil && o.NotApplicableDeviceCount != nil {
		return true
	}

	return false
}

// SetNotApplicableDeviceCount gets a reference to the given int32 and assigns it to the NotApplicableDeviceCount field.
func (o *MicrosoftGraphSoftwareUpdateStatusSummary) SetNotApplicableDeviceCount(v int32) {
	o.NotApplicableDeviceCount = &v
}

// GetNotApplicableUserCount returns the NotApplicableUserCount field value if set, zero value otherwise.
func (o *MicrosoftGraphSoftwareUpdateStatusSummary) GetNotApplicableUserCount() int32 {
	if o == nil || o.NotApplicableUserCount == nil {
		var ret int32
		return ret
	}
	return *o.NotApplicableUserCount
}

// GetNotApplicableUserCountOk returns a tuple with the NotApplicableUserCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphSoftwareUpdateStatusSummary) GetNotApplicableUserCountOk() (*int32, bool) {
	if o == nil || o.NotApplicableUserCount == nil {
		return nil, false
	}
	return o.NotApplicableUserCount, true
}

// HasNotApplicableUserCount returns a boolean if a field has been set.
func (o *MicrosoftGraphSoftwareUpdateStatusSummary) HasNotApplicableUserCount() bool {
	if o != nil && o.NotApplicableUserCount != nil {
		return true
	}

	return false
}

// SetNotApplicableUserCount gets a reference to the given int32 and assigns it to the NotApplicableUserCount field.
func (o *MicrosoftGraphSoftwareUpdateStatusSummary) SetNotApplicableUserCount(v int32) {
	o.NotApplicableUserCount = &v
}

// GetRemediatedDeviceCount returns the RemediatedDeviceCount field value if set, zero value otherwise.
func (o *MicrosoftGraphSoftwareUpdateStatusSummary) GetRemediatedDeviceCount() int32 {
	if o == nil || o.RemediatedDeviceCount == nil {
		var ret int32
		return ret
	}
	return *o.RemediatedDeviceCount
}

// GetRemediatedDeviceCountOk returns a tuple with the RemediatedDeviceCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphSoftwareUpdateStatusSummary) GetRemediatedDeviceCountOk() (*int32, bool) {
	if o == nil || o.RemediatedDeviceCount == nil {
		return nil, false
	}
	return o.RemediatedDeviceCount, true
}

// HasRemediatedDeviceCount returns a boolean if a field has been set.
func (o *MicrosoftGraphSoftwareUpdateStatusSummary) HasRemediatedDeviceCount() bool {
	if o != nil && o.RemediatedDeviceCount != nil {
		return true
	}

	return false
}

// SetRemediatedDeviceCount gets a reference to the given int32 and assigns it to the RemediatedDeviceCount field.
func (o *MicrosoftGraphSoftwareUpdateStatusSummary) SetRemediatedDeviceCount(v int32) {
	o.RemediatedDeviceCount = &v
}

// GetRemediatedUserCount returns the RemediatedUserCount field value if set, zero value otherwise.
func (o *MicrosoftGraphSoftwareUpdateStatusSummary) GetRemediatedUserCount() int32 {
	if o == nil || o.RemediatedUserCount == nil {
		var ret int32
		return ret
	}
	return *o.RemediatedUserCount
}

// GetRemediatedUserCountOk returns a tuple with the RemediatedUserCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphSoftwareUpdateStatusSummary) GetRemediatedUserCountOk() (*int32, bool) {
	if o == nil || o.RemediatedUserCount == nil {
		return nil, false
	}
	return o.RemediatedUserCount, true
}

// HasRemediatedUserCount returns a boolean if a field has been set.
func (o *MicrosoftGraphSoftwareUpdateStatusSummary) HasRemediatedUserCount() bool {
	if o != nil && o.RemediatedUserCount != nil {
		return true
	}

	return false
}

// SetRemediatedUserCount gets a reference to the given int32 and assigns it to the RemediatedUserCount field.
func (o *MicrosoftGraphSoftwareUpdateStatusSummary) SetRemediatedUserCount(v int32) {
	o.RemediatedUserCount = &v
}

// GetUnknownDeviceCount returns the UnknownDeviceCount field value if set, zero value otherwise.
func (o *MicrosoftGraphSoftwareUpdateStatusSummary) GetUnknownDeviceCount() int32 {
	if o == nil || o.UnknownDeviceCount == nil {
		var ret int32
		return ret
	}
	return *o.UnknownDeviceCount
}

// GetUnknownDeviceCountOk returns a tuple with the UnknownDeviceCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphSoftwareUpdateStatusSummary) GetUnknownDeviceCountOk() (*int32, bool) {
	if o == nil || o.UnknownDeviceCount == nil {
		return nil, false
	}
	return o.UnknownDeviceCount, true
}

// HasUnknownDeviceCount returns a boolean if a field has been set.
func (o *MicrosoftGraphSoftwareUpdateStatusSummary) HasUnknownDeviceCount() bool {
	if o != nil && o.UnknownDeviceCount != nil {
		return true
	}

	return false
}

// SetUnknownDeviceCount gets a reference to the given int32 and assigns it to the UnknownDeviceCount field.
func (o *MicrosoftGraphSoftwareUpdateStatusSummary) SetUnknownDeviceCount(v int32) {
	o.UnknownDeviceCount = &v
}

// GetUnknownUserCount returns the UnknownUserCount field value if set, zero value otherwise.
func (o *MicrosoftGraphSoftwareUpdateStatusSummary) GetUnknownUserCount() int32 {
	if o == nil || o.UnknownUserCount == nil {
		var ret int32
		return ret
	}
	return *o.UnknownUserCount
}

// GetUnknownUserCountOk returns a tuple with the UnknownUserCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphSoftwareUpdateStatusSummary) GetUnknownUserCountOk() (*int32, bool) {
	if o == nil || o.UnknownUserCount == nil {
		return nil, false
	}
	return o.UnknownUserCount, true
}

// HasUnknownUserCount returns a boolean if a field has been set.
func (o *MicrosoftGraphSoftwareUpdateStatusSummary) HasUnknownUserCount() bool {
	if o != nil && o.UnknownUserCount != nil {
		return true
	}

	return false
}

// SetUnknownUserCount gets a reference to the given int32 and assigns it to the UnknownUserCount field.
func (o *MicrosoftGraphSoftwareUpdateStatusSummary) SetUnknownUserCount(v int32) {
	o.UnknownUserCount = &v
}

func (o MicrosoftGraphSoftwareUpdateStatusSummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.CompliantDeviceCount != nil {
		toSerialize["compliantDeviceCount"] = o.CompliantDeviceCount
	}
	if o.CompliantUserCount != nil {
		toSerialize["compliantUserCount"] = o.CompliantUserCount
	}
	if o.ConflictDeviceCount != nil {
		toSerialize["conflictDeviceCount"] = o.ConflictDeviceCount
	}
	if o.ConflictUserCount != nil {
		toSerialize["conflictUserCount"] = o.ConflictUserCount
	}
	if o.DisplayName.IsSet() {
		toSerialize["displayName"] = o.DisplayName.Get()
	}
	if o.ErrorDeviceCount != nil {
		toSerialize["errorDeviceCount"] = o.ErrorDeviceCount
	}
	if o.ErrorUserCount != nil {
		toSerialize["errorUserCount"] = o.ErrorUserCount
	}
	if o.NonCompliantDeviceCount != nil {
		toSerialize["nonCompliantDeviceCount"] = o.NonCompliantDeviceCount
	}
	if o.NonCompliantUserCount != nil {
		toSerialize["nonCompliantUserCount"] = o.NonCompliantUserCount
	}
	if o.NotApplicableDeviceCount != nil {
		toSerialize["notApplicableDeviceCount"] = o.NotApplicableDeviceCount
	}
	if o.NotApplicableUserCount != nil {
		toSerialize["notApplicableUserCount"] = o.NotApplicableUserCount
	}
	if o.RemediatedDeviceCount != nil {
		toSerialize["remediatedDeviceCount"] = o.RemediatedDeviceCount
	}
	if o.RemediatedUserCount != nil {
		toSerialize["remediatedUserCount"] = o.RemediatedUserCount
	}
	if o.UnknownDeviceCount != nil {
		toSerialize["unknownDeviceCount"] = o.UnknownDeviceCount
	}
	if o.UnknownUserCount != nil {
		toSerialize["unknownUserCount"] = o.UnknownUserCount
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphSoftwareUpdateStatusSummary struct {
	value *MicrosoftGraphSoftwareUpdateStatusSummary
	isSet bool
}

func (v NullableMicrosoftGraphSoftwareUpdateStatusSummary) Get() *MicrosoftGraphSoftwareUpdateStatusSummary {
	return v.value
}

func (v *NullableMicrosoftGraphSoftwareUpdateStatusSummary) Set(val *MicrosoftGraphSoftwareUpdateStatusSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphSoftwareUpdateStatusSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphSoftwareUpdateStatusSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphSoftwareUpdateStatusSummary(val *MicrosoftGraphSoftwareUpdateStatusSummary) *NullableMicrosoftGraphSoftwareUpdateStatusSummary {
	return &NullableMicrosoftGraphSoftwareUpdateStatusSummary{value: val, isSet: true}
}

func (v NullableMicrosoftGraphSoftwareUpdateStatusSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphSoftwareUpdateStatusSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

