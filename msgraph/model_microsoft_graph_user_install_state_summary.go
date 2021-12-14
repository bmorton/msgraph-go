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

// MicrosoftGraphUserInstallStateSummary struct for MicrosoftGraphUserInstallStateSummary
type MicrosoftGraphUserInstallStateSummary struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
	// Failed Device Count.
	FailedDeviceCount *int32 `json:"failedDeviceCount,omitempty"`
	// Installed Device Count.
	InstalledDeviceCount *int32 `json:"installedDeviceCount,omitempty"`
	// Not installed device count.
	NotInstalledDeviceCount *int32 `json:"notInstalledDeviceCount,omitempty"`
	// User name.
	UserName NullableString `json:"userName,omitempty"`
	// The install state of the eBook.
	DeviceStates *[]MicrosoftGraphDeviceInstallState `json:"deviceStates,omitempty"`
}

// NewMicrosoftGraphUserInstallStateSummary instantiates a new MicrosoftGraphUserInstallStateSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphUserInstallStateSummary() *MicrosoftGraphUserInstallStateSummary {
	this := MicrosoftGraphUserInstallStateSummary{}
	return &this
}

// NewMicrosoftGraphUserInstallStateSummaryWithDefaults instantiates a new MicrosoftGraphUserInstallStateSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphUserInstallStateSummaryWithDefaults() *MicrosoftGraphUserInstallStateSummary {
	this := MicrosoftGraphUserInstallStateSummary{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MicrosoftGraphUserInstallStateSummary) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphUserInstallStateSummary) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphUserInstallStateSummary) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MicrosoftGraphUserInstallStateSummary) SetId(v string) {
	o.Id = &v
}

// GetFailedDeviceCount returns the FailedDeviceCount field value if set, zero value otherwise.
func (o *MicrosoftGraphUserInstallStateSummary) GetFailedDeviceCount() int32 {
	if o == nil || o.FailedDeviceCount == nil {
		var ret int32
		return ret
	}
	return *o.FailedDeviceCount
}

// GetFailedDeviceCountOk returns a tuple with the FailedDeviceCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphUserInstallStateSummary) GetFailedDeviceCountOk() (*int32, bool) {
	if o == nil || o.FailedDeviceCount == nil {
		return nil, false
	}
	return o.FailedDeviceCount, true
}

// HasFailedDeviceCount returns a boolean if a field has been set.
func (o *MicrosoftGraphUserInstallStateSummary) HasFailedDeviceCount() bool {
	if o != nil && o.FailedDeviceCount != nil {
		return true
	}

	return false
}

// SetFailedDeviceCount gets a reference to the given int32 and assigns it to the FailedDeviceCount field.
func (o *MicrosoftGraphUserInstallStateSummary) SetFailedDeviceCount(v int32) {
	o.FailedDeviceCount = &v
}

// GetInstalledDeviceCount returns the InstalledDeviceCount field value if set, zero value otherwise.
func (o *MicrosoftGraphUserInstallStateSummary) GetInstalledDeviceCount() int32 {
	if o == nil || o.InstalledDeviceCount == nil {
		var ret int32
		return ret
	}
	return *o.InstalledDeviceCount
}

// GetInstalledDeviceCountOk returns a tuple with the InstalledDeviceCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphUserInstallStateSummary) GetInstalledDeviceCountOk() (*int32, bool) {
	if o == nil || o.InstalledDeviceCount == nil {
		return nil, false
	}
	return o.InstalledDeviceCount, true
}

// HasInstalledDeviceCount returns a boolean if a field has been set.
func (o *MicrosoftGraphUserInstallStateSummary) HasInstalledDeviceCount() bool {
	if o != nil && o.InstalledDeviceCount != nil {
		return true
	}

	return false
}

// SetInstalledDeviceCount gets a reference to the given int32 and assigns it to the InstalledDeviceCount field.
func (o *MicrosoftGraphUserInstallStateSummary) SetInstalledDeviceCount(v int32) {
	o.InstalledDeviceCount = &v
}

// GetNotInstalledDeviceCount returns the NotInstalledDeviceCount field value if set, zero value otherwise.
func (o *MicrosoftGraphUserInstallStateSummary) GetNotInstalledDeviceCount() int32 {
	if o == nil || o.NotInstalledDeviceCount == nil {
		var ret int32
		return ret
	}
	return *o.NotInstalledDeviceCount
}

// GetNotInstalledDeviceCountOk returns a tuple with the NotInstalledDeviceCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphUserInstallStateSummary) GetNotInstalledDeviceCountOk() (*int32, bool) {
	if o == nil || o.NotInstalledDeviceCount == nil {
		return nil, false
	}
	return o.NotInstalledDeviceCount, true
}

// HasNotInstalledDeviceCount returns a boolean if a field has been set.
func (o *MicrosoftGraphUserInstallStateSummary) HasNotInstalledDeviceCount() bool {
	if o != nil && o.NotInstalledDeviceCount != nil {
		return true
	}

	return false
}

// SetNotInstalledDeviceCount gets a reference to the given int32 and assigns it to the NotInstalledDeviceCount field.
func (o *MicrosoftGraphUserInstallStateSummary) SetNotInstalledDeviceCount(v int32) {
	o.NotInstalledDeviceCount = &v
}

// GetUserName returns the UserName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphUserInstallStateSummary) GetUserName() string {
	if o == nil || o.UserName.Get() == nil {
		var ret string
		return ret
	}
	return *o.UserName.Get()
}

// GetUserNameOk returns a tuple with the UserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphUserInstallStateSummary) GetUserNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.UserName.Get(), o.UserName.IsSet()
}

// HasUserName returns a boolean if a field has been set.
func (o *MicrosoftGraphUserInstallStateSummary) HasUserName() bool {
	if o != nil && o.UserName.IsSet() {
		return true
	}

	return false
}

// SetUserName gets a reference to the given NullableString and assigns it to the UserName field.
func (o *MicrosoftGraphUserInstallStateSummary) SetUserName(v string) {
	o.UserName.Set(&v)
}
// SetUserNameNil sets the value for UserName to be an explicit nil
func (o *MicrosoftGraphUserInstallStateSummary) SetUserNameNil() {
	o.UserName.Set(nil)
}

// UnsetUserName ensures that no value is present for UserName, not even an explicit nil
func (o *MicrosoftGraphUserInstallStateSummary) UnsetUserName() {
	o.UserName.Unset()
}

// GetDeviceStates returns the DeviceStates field value if set, zero value otherwise.
func (o *MicrosoftGraphUserInstallStateSummary) GetDeviceStates() []MicrosoftGraphDeviceInstallState {
	if o == nil || o.DeviceStates == nil {
		var ret []MicrosoftGraphDeviceInstallState
		return ret
	}
	return *o.DeviceStates
}

// GetDeviceStatesOk returns a tuple with the DeviceStates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphUserInstallStateSummary) GetDeviceStatesOk() (*[]MicrosoftGraphDeviceInstallState, bool) {
	if o == nil || o.DeviceStates == nil {
		return nil, false
	}
	return o.DeviceStates, true
}

// HasDeviceStates returns a boolean if a field has been set.
func (o *MicrosoftGraphUserInstallStateSummary) HasDeviceStates() bool {
	if o != nil && o.DeviceStates != nil {
		return true
	}

	return false
}

// SetDeviceStates gets a reference to the given []MicrosoftGraphDeviceInstallState and assigns it to the DeviceStates field.
func (o *MicrosoftGraphUserInstallStateSummary) SetDeviceStates(v []MicrosoftGraphDeviceInstallState) {
	o.DeviceStates = &v
}

func (o MicrosoftGraphUserInstallStateSummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.FailedDeviceCount != nil {
		toSerialize["failedDeviceCount"] = o.FailedDeviceCount
	}
	if o.InstalledDeviceCount != nil {
		toSerialize["installedDeviceCount"] = o.InstalledDeviceCount
	}
	if o.NotInstalledDeviceCount != nil {
		toSerialize["notInstalledDeviceCount"] = o.NotInstalledDeviceCount
	}
	if o.UserName.IsSet() {
		toSerialize["userName"] = o.UserName.Get()
	}
	if o.DeviceStates != nil {
		toSerialize["deviceStates"] = o.DeviceStates
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphUserInstallStateSummary struct {
	value *MicrosoftGraphUserInstallStateSummary
	isSet bool
}

func (v NullableMicrosoftGraphUserInstallStateSummary) Get() *MicrosoftGraphUserInstallStateSummary {
	return v.value
}

func (v *NullableMicrosoftGraphUserInstallStateSummary) Set(val *MicrosoftGraphUserInstallStateSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphUserInstallStateSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphUserInstallStateSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphUserInstallStateSummary(val *MicrosoftGraphUserInstallStateSummary) *NullableMicrosoftGraphUserInstallStateSummary {
	return &NullableMicrosoftGraphUserInstallStateSummary{value: val, isSet: true}
}

func (v NullableMicrosoftGraphUserInstallStateSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphUserInstallStateSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


