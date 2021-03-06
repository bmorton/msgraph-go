/*
Partial Graph API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package msgraph

import (
	"encoding/json"
)

// DeviceConfigurationState Device Configuration State for a given device.
type DeviceConfigurationState struct {
	// The name of the policy for this policyBase
	DisplayName NullableString `json:"displayName,omitempty"`
	// Platform type that the policy applies to
	PlatformType AnyOfmicrosoftGraphPolicyPlatformType `json:"platformType,omitempty"`
	// Count of how many setting a policy holds
	SettingCount *int32 `json:"settingCount,omitempty"`
	SettingStates *[]*AnyOfmicrosoftGraphDeviceConfigurationSettingState `json:"settingStates,omitempty"`
	// The compliance state of the policy
	State AnyOfmicrosoftGraphComplianceStatus `json:"state,omitempty"`
	// The version of the policy
	Version *int32 `json:"version,omitempty"`
}

// NewDeviceConfigurationState instantiates a new DeviceConfigurationState object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceConfigurationState() *DeviceConfigurationState {
	this := DeviceConfigurationState{}
	return &this
}

// NewDeviceConfigurationStateWithDefaults instantiates a new DeviceConfigurationState object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceConfigurationStateWithDefaults() *DeviceConfigurationState {
	this := DeviceConfigurationState{}
	return &this
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeviceConfigurationState) GetDisplayName() string {
	if o == nil || o.DisplayName.Get() == nil {
		var ret string
		return ret
	}
	return *o.DisplayName.Get()
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeviceConfigurationState) GetDisplayNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DisplayName.Get(), o.DisplayName.IsSet()
}

// HasDisplayName returns a boolean if a field has been set.
func (o *DeviceConfigurationState) HasDisplayName() bool {
	if o != nil && o.DisplayName.IsSet() {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given NullableString and assigns it to the DisplayName field.
func (o *DeviceConfigurationState) SetDisplayName(v string) {
	o.DisplayName.Set(&v)
}
// SetDisplayNameNil sets the value for DisplayName to be an explicit nil
func (o *DeviceConfigurationState) SetDisplayNameNil() {
	o.DisplayName.Set(nil)
}

// UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
func (o *DeviceConfigurationState) UnsetDisplayName() {
	o.DisplayName.Unset()
}

// GetPlatformType returns the PlatformType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeviceConfigurationState) GetPlatformType() AnyOfmicrosoftGraphPolicyPlatformType {
	if o == nil  {
		var ret AnyOfmicrosoftGraphPolicyPlatformType
		return ret
	}
	return o.PlatformType
}

// GetPlatformTypeOk returns a tuple with the PlatformType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeviceConfigurationState) GetPlatformTypeOk() (*AnyOfmicrosoftGraphPolicyPlatformType, bool) {
	if o == nil || o.PlatformType == nil {
		return nil, false
	}
	return &o.PlatformType, true
}

// HasPlatformType returns a boolean if a field has been set.
func (o *DeviceConfigurationState) HasPlatformType() bool {
	if o != nil && o.PlatformType != nil {
		return true
	}

	return false
}

// SetPlatformType gets a reference to the given AnyOfmicrosoftGraphPolicyPlatformType and assigns it to the PlatformType field.
func (o *DeviceConfigurationState) SetPlatformType(v AnyOfmicrosoftGraphPolicyPlatformType) {
	o.PlatformType = v
}

// GetSettingCount returns the SettingCount field value if set, zero value otherwise.
func (o *DeviceConfigurationState) GetSettingCount() int32 {
	if o == nil || o.SettingCount == nil {
		var ret int32
		return ret
	}
	return *o.SettingCount
}

// GetSettingCountOk returns a tuple with the SettingCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceConfigurationState) GetSettingCountOk() (*int32, bool) {
	if o == nil || o.SettingCount == nil {
		return nil, false
	}
	return o.SettingCount, true
}

// HasSettingCount returns a boolean if a field has been set.
func (o *DeviceConfigurationState) HasSettingCount() bool {
	if o != nil && o.SettingCount != nil {
		return true
	}

	return false
}

// SetSettingCount gets a reference to the given int32 and assigns it to the SettingCount field.
func (o *DeviceConfigurationState) SetSettingCount(v int32) {
	o.SettingCount = &v
}

// GetSettingStates returns the SettingStates field value if set, zero value otherwise.
func (o *DeviceConfigurationState) GetSettingStates() []*AnyOfmicrosoftGraphDeviceConfigurationSettingState {
	if o == nil || o.SettingStates == nil {
		var ret []*AnyOfmicrosoftGraphDeviceConfigurationSettingState
		return ret
	}
	return *o.SettingStates
}

// GetSettingStatesOk returns a tuple with the SettingStates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceConfigurationState) GetSettingStatesOk() (*[]*AnyOfmicrosoftGraphDeviceConfigurationSettingState, bool) {
	if o == nil || o.SettingStates == nil {
		return nil, false
	}
	return o.SettingStates, true
}

// HasSettingStates returns a boolean if a field has been set.
func (o *DeviceConfigurationState) HasSettingStates() bool {
	if o != nil && o.SettingStates != nil {
		return true
	}

	return false
}

// SetSettingStates gets a reference to the given []*AnyOfmicrosoftGraphDeviceConfigurationSettingState and assigns it to the SettingStates field.
func (o *DeviceConfigurationState) SetSettingStates(v []*AnyOfmicrosoftGraphDeviceConfigurationSettingState) {
	o.SettingStates = &v
}

// GetState returns the State field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeviceConfigurationState) GetState() AnyOfmicrosoftGraphComplianceStatus {
	if o == nil  {
		var ret AnyOfmicrosoftGraphComplianceStatus
		return ret
	}
	return o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeviceConfigurationState) GetStateOk() (*AnyOfmicrosoftGraphComplianceStatus, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return &o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *DeviceConfigurationState) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given AnyOfmicrosoftGraphComplianceStatus and assigns it to the State field.
func (o *DeviceConfigurationState) SetState(v AnyOfmicrosoftGraphComplianceStatus) {
	o.State = v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *DeviceConfigurationState) GetVersion() int32 {
	if o == nil || o.Version == nil {
		var ret int32
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceConfigurationState) GetVersionOk() (*int32, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *DeviceConfigurationState) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int32 and assigns it to the Version field.
func (o *DeviceConfigurationState) SetVersion(v int32) {
	o.Version = &v
}

func (o DeviceConfigurationState) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DisplayName.IsSet() {
		toSerialize["displayName"] = o.DisplayName.Get()
	}
	if o.PlatformType != nil {
		toSerialize["platformType"] = o.PlatformType
	}
	if o.SettingCount != nil {
		toSerialize["settingCount"] = o.SettingCount
	}
	if o.SettingStates != nil {
		toSerialize["settingStates"] = o.SettingStates
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	return json.Marshal(toSerialize)
}

type NullableDeviceConfigurationState struct {
	value *DeviceConfigurationState
	isSet bool
}

func (v NullableDeviceConfigurationState) Get() *DeviceConfigurationState {
	return v.value
}

func (v *NullableDeviceConfigurationState) Set(val *DeviceConfigurationState) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceConfigurationState) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceConfigurationState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceConfigurationState(val *DeviceConfigurationState) *NullableDeviceConfigurationState {
	return &NullableDeviceConfigurationState{value: val, isSet: true}
}

func (v NullableDeviceConfigurationState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceConfigurationState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


