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

// MicrosoftGraphDeviceConfiguration struct for MicrosoftGraphDeviceConfiguration
type MicrosoftGraphDeviceConfiguration struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
	// DateTime the object was created.
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// Admin provided description of the Device Configuration.
	Description NullableString `json:"description,omitempty"`
	// Admin provided name of the device configuration.
	DisplayName *string `json:"displayName,omitempty"`
	// DateTime the object was last modified.
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
	// Version of the device configuration.
	Version *int32 `json:"version,omitempty"`
	// The list of assignments for the device configuration profile.
	Assignments *[]MicrosoftGraphDeviceConfigurationAssignment `json:"assignments,omitempty"`
	// Device Configuration Setting State Device Summary
	DeviceSettingStateSummaries *[]MicrosoftGraphSettingStateDeviceSummary `json:"deviceSettingStateSummaries,omitempty"`
	// Device configuration installation status by device.
	DeviceStatuses *[]MicrosoftGraphDeviceConfigurationDeviceStatus `json:"deviceStatuses,omitempty"`
	// Device Configuration devices status overview
	DeviceStatusOverview AnyOfmicrosoftGraphDeviceConfigurationDeviceOverview `json:"deviceStatusOverview,omitempty"`
	// Device configuration installation status by user.
	UserStatuses *[]MicrosoftGraphDeviceConfigurationUserStatus `json:"userStatuses,omitempty"`
	// Device Configuration users status overview
	UserStatusOverview AnyOfmicrosoftGraphDeviceConfigurationUserOverview `json:"userStatusOverview,omitempty"`
}

// NewMicrosoftGraphDeviceConfiguration instantiates a new MicrosoftGraphDeviceConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphDeviceConfiguration() *MicrosoftGraphDeviceConfiguration {
	this := MicrosoftGraphDeviceConfiguration{}
	return &this
}

// NewMicrosoftGraphDeviceConfigurationWithDefaults instantiates a new MicrosoftGraphDeviceConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphDeviceConfigurationWithDefaults() *MicrosoftGraphDeviceConfiguration {
	this := MicrosoftGraphDeviceConfiguration{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MicrosoftGraphDeviceConfiguration) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphDeviceConfiguration) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphDeviceConfiguration) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MicrosoftGraphDeviceConfiguration) SetId(v string) {
	o.Id = &v
}

// GetCreatedDateTime returns the CreatedDateTime field value if set, zero value otherwise.
func (o *MicrosoftGraphDeviceConfiguration) GetCreatedDateTime() time.Time {
	if o == nil || o.CreatedDateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedDateTime
}

// GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphDeviceConfiguration) GetCreatedDateTimeOk() (*time.Time, bool) {
	if o == nil || o.CreatedDateTime == nil {
		return nil, false
	}
	return o.CreatedDateTime, true
}

// HasCreatedDateTime returns a boolean if a field has been set.
func (o *MicrosoftGraphDeviceConfiguration) HasCreatedDateTime() bool {
	if o != nil && o.CreatedDateTime != nil {
		return true
	}

	return false
}

// SetCreatedDateTime gets a reference to the given time.Time and assigns it to the CreatedDateTime field.
func (o *MicrosoftGraphDeviceConfiguration) SetCreatedDateTime(v time.Time) {
	o.CreatedDateTime = &v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphDeviceConfiguration) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphDeviceConfiguration) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *MicrosoftGraphDeviceConfiguration) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *MicrosoftGraphDeviceConfiguration) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *MicrosoftGraphDeviceConfiguration) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *MicrosoftGraphDeviceConfiguration) UnsetDescription() {
	o.Description.Unset()
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *MicrosoftGraphDeviceConfiguration) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphDeviceConfiguration) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *MicrosoftGraphDeviceConfiguration) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *MicrosoftGraphDeviceConfiguration) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetLastModifiedDateTime returns the LastModifiedDateTime field value if set, zero value otherwise.
func (o *MicrosoftGraphDeviceConfiguration) GetLastModifiedDateTime() time.Time {
	if o == nil || o.LastModifiedDateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.LastModifiedDateTime
}

// GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphDeviceConfiguration) GetLastModifiedDateTimeOk() (*time.Time, bool) {
	if o == nil || o.LastModifiedDateTime == nil {
		return nil, false
	}
	return o.LastModifiedDateTime, true
}

// HasLastModifiedDateTime returns a boolean if a field has been set.
func (o *MicrosoftGraphDeviceConfiguration) HasLastModifiedDateTime() bool {
	if o != nil && o.LastModifiedDateTime != nil {
		return true
	}

	return false
}

// SetLastModifiedDateTime gets a reference to the given time.Time and assigns it to the LastModifiedDateTime field.
func (o *MicrosoftGraphDeviceConfiguration) SetLastModifiedDateTime(v time.Time) {
	o.LastModifiedDateTime = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *MicrosoftGraphDeviceConfiguration) GetVersion() int32 {
	if o == nil || o.Version == nil {
		var ret int32
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphDeviceConfiguration) GetVersionOk() (*int32, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *MicrosoftGraphDeviceConfiguration) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int32 and assigns it to the Version field.
func (o *MicrosoftGraphDeviceConfiguration) SetVersion(v int32) {
	o.Version = &v
}

// GetAssignments returns the Assignments field value if set, zero value otherwise.
func (o *MicrosoftGraphDeviceConfiguration) GetAssignments() []MicrosoftGraphDeviceConfigurationAssignment {
	if o == nil || o.Assignments == nil {
		var ret []MicrosoftGraphDeviceConfigurationAssignment
		return ret
	}
	return *o.Assignments
}

// GetAssignmentsOk returns a tuple with the Assignments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphDeviceConfiguration) GetAssignmentsOk() (*[]MicrosoftGraphDeviceConfigurationAssignment, bool) {
	if o == nil || o.Assignments == nil {
		return nil, false
	}
	return o.Assignments, true
}

// HasAssignments returns a boolean if a field has been set.
func (o *MicrosoftGraphDeviceConfiguration) HasAssignments() bool {
	if o != nil && o.Assignments != nil {
		return true
	}

	return false
}

// SetAssignments gets a reference to the given []MicrosoftGraphDeviceConfigurationAssignment and assigns it to the Assignments field.
func (o *MicrosoftGraphDeviceConfiguration) SetAssignments(v []MicrosoftGraphDeviceConfigurationAssignment) {
	o.Assignments = &v
}

// GetDeviceSettingStateSummaries returns the DeviceSettingStateSummaries field value if set, zero value otherwise.
func (o *MicrosoftGraphDeviceConfiguration) GetDeviceSettingStateSummaries() []MicrosoftGraphSettingStateDeviceSummary {
	if o == nil || o.DeviceSettingStateSummaries == nil {
		var ret []MicrosoftGraphSettingStateDeviceSummary
		return ret
	}
	return *o.DeviceSettingStateSummaries
}

// GetDeviceSettingStateSummariesOk returns a tuple with the DeviceSettingStateSummaries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphDeviceConfiguration) GetDeviceSettingStateSummariesOk() (*[]MicrosoftGraphSettingStateDeviceSummary, bool) {
	if o == nil || o.DeviceSettingStateSummaries == nil {
		return nil, false
	}
	return o.DeviceSettingStateSummaries, true
}

// HasDeviceSettingStateSummaries returns a boolean if a field has been set.
func (o *MicrosoftGraphDeviceConfiguration) HasDeviceSettingStateSummaries() bool {
	if o != nil && o.DeviceSettingStateSummaries != nil {
		return true
	}

	return false
}

// SetDeviceSettingStateSummaries gets a reference to the given []MicrosoftGraphSettingStateDeviceSummary and assigns it to the DeviceSettingStateSummaries field.
func (o *MicrosoftGraphDeviceConfiguration) SetDeviceSettingStateSummaries(v []MicrosoftGraphSettingStateDeviceSummary) {
	o.DeviceSettingStateSummaries = &v
}

// GetDeviceStatuses returns the DeviceStatuses field value if set, zero value otherwise.
func (o *MicrosoftGraphDeviceConfiguration) GetDeviceStatuses() []MicrosoftGraphDeviceConfigurationDeviceStatus {
	if o == nil || o.DeviceStatuses == nil {
		var ret []MicrosoftGraphDeviceConfigurationDeviceStatus
		return ret
	}
	return *o.DeviceStatuses
}

// GetDeviceStatusesOk returns a tuple with the DeviceStatuses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphDeviceConfiguration) GetDeviceStatusesOk() (*[]MicrosoftGraphDeviceConfigurationDeviceStatus, bool) {
	if o == nil || o.DeviceStatuses == nil {
		return nil, false
	}
	return o.DeviceStatuses, true
}

// HasDeviceStatuses returns a boolean if a field has been set.
func (o *MicrosoftGraphDeviceConfiguration) HasDeviceStatuses() bool {
	if o != nil && o.DeviceStatuses != nil {
		return true
	}

	return false
}

// SetDeviceStatuses gets a reference to the given []MicrosoftGraphDeviceConfigurationDeviceStatus and assigns it to the DeviceStatuses field.
func (o *MicrosoftGraphDeviceConfiguration) SetDeviceStatuses(v []MicrosoftGraphDeviceConfigurationDeviceStatus) {
	o.DeviceStatuses = &v
}

// GetDeviceStatusOverview returns the DeviceStatusOverview field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphDeviceConfiguration) GetDeviceStatusOverview() AnyOfmicrosoftGraphDeviceConfigurationDeviceOverview {
	if o == nil  {
		var ret AnyOfmicrosoftGraphDeviceConfigurationDeviceOverview
		return ret
	}
	return o.DeviceStatusOverview
}

// GetDeviceStatusOverviewOk returns a tuple with the DeviceStatusOverview field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphDeviceConfiguration) GetDeviceStatusOverviewOk() (*AnyOfmicrosoftGraphDeviceConfigurationDeviceOverview, bool) {
	if o == nil || o.DeviceStatusOverview == nil {
		return nil, false
	}
	return &o.DeviceStatusOverview, true
}

// HasDeviceStatusOverview returns a boolean if a field has been set.
func (o *MicrosoftGraphDeviceConfiguration) HasDeviceStatusOverview() bool {
	if o != nil && o.DeviceStatusOverview != nil {
		return true
	}

	return false
}

// SetDeviceStatusOverview gets a reference to the given AnyOfmicrosoftGraphDeviceConfigurationDeviceOverview and assigns it to the DeviceStatusOverview field.
func (o *MicrosoftGraphDeviceConfiguration) SetDeviceStatusOverview(v AnyOfmicrosoftGraphDeviceConfigurationDeviceOverview) {
	o.DeviceStatusOverview = v
}

// GetUserStatuses returns the UserStatuses field value if set, zero value otherwise.
func (o *MicrosoftGraphDeviceConfiguration) GetUserStatuses() []MicrosoftGraphDeviceConfigurationUserStatus {
	if o == nil || o.UserStatuses == nil {
		var ret []MicrosoftGraphDeviceConfigurationUserStatus
		return ret
	}
	return *o.UserStatuses
}

// GetUserStatusesOk returns a tuple with the UserStatuses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphDeviceConfiguration) GetUserStatusesOk() (*[]MicrosoftGraphDeviceConfigurationUserStatus, bool) {
	if o == nil || o.UserStatuses == nil {
		return nil, false
	}
	return o.UserStatuses, true
}

// HasUserStatuses returns a boolean if a field has been set.
func (o *MicrosoftGraphDeviceConfiguration) HasUserStatuses() bool {
	if o != nil && o.UserStatuses != nil {
		return true
	}

	return false
}

// SetUserStatuses gets a reference to the given []MicrosoftGraphDeviceConfigurationUserStatus and assigns it to the UserStatuses field.
func (o *MicrosoftGraphDeviceConfiguration) SetUserStatuses(v []MicrosoftGraphDeviceConfigurationUserStatus) {
	o.UserStatuses = &v
}

// GetUserStatusOverview returns the UserStatusOverview field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphDeviceConfiguration) GetUserStatusOverview() AnyOfmicrosoftGraphDeviceConfigurationUserOverview {
	if o == nil  {
		var ret AnyOfmicrosoftGraphDeviceConfigurationUserOverview
		return ret
	}
	return o.UserStatusOverview
}

// GetUserStatusOverviewOk returns a tuple with the UserStatusOverview field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphDeviceConfiguration) GetUserStatusOverviewOk() (*AnyOfmicrosoftGraphDeviceConfigurationUserOverview, bool) {
	if o == nil || o.UserStatusOverview == nil {
		return nil, false
	}
	return &o.UserStatusOverview, true
}

// HasUserStatusOverview returns a boolean if a field has been set.
func (o *MicrosoftGraphDeviceConfiguration) HasUserStatusOverview() bool {
	if o != nil && o.UserStatusOverview != nil {
		return true
	}

	return false
}

// SetUserStatusOverview gets a reference to the given AnyOfmicrosoftGraphDeviceConfigurationUserOverview and assigns it to the UserStatusOverview field.
func (o *MicrosoftGraphDeviceConfiguration) SetUserStatusOverview(v AnyOfmicrosoftGraphDeviceConfigurationUserOverview) {
	o.UserStatusOverview = v
}

func (o MicrosoftGraphDeviceConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.CreatedDateTime != nil {
		toSerialize["createdDateTime"] = o.CreatedDateTime
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.DisplayName != nil {
		toSerialize["displayName"] = o.DisplayName
	}
	if o.LastModifiedDateTime != nil {
		toSerialize["lastModifiedDateTime"] = o.LastModifiedDateTime
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	if o.Assignments != nil {
		toSerialize["assignments"] = o.Assignments
	}
	if o.DeviceSettingStateSummaries != nil {
		toSerialize["deviceSettingStateSummaries"] = o.DeviceSettingStateSummaries
	}
	if o.DeviceStatuses != nil {
		toSerialize["deviceStatuses"] = o.DeviceStatuses
	}
	if o.DeviceStatusOverview != nil {
		toSerialize["deviceStatusOverview"] = o.DeviceStatusOverview
	}
	if o.UserStatuses != nil {
		toSerialize["userStatuses"] = o.UserStatuses
	}
	if o.UserStatusOverview != nil {
		toSerialize["userStatusOverview"] = o.UserStatusOverview
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphDeviceConfiguration struct {
	value *MicrosoftGraphDeviceConfiguration
	isSet bool
}

func (v NullableMicrosoftGraphDeviceConfiguration) Get() *MicrosoftGraphDeviceConfiguration {
	return v.value
}

func (v *NullableMicrosoftGraphDeviceConfiguration) Set(val *MicrosoftGraphDeviceConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphDeviceConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphDeviceConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphDeviceConfiguration(val *MicrosoftGraphDeviceConfiguration) *NullableMicrosoftGraphDeviceConfiguration {
	return &NullableMicrosoftGraphDeviceConfiguration{value: val, isSet: true}
}

func (v NullableMicrosoftGraphDeviceConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphDeviceConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


