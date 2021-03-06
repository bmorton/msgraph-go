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

// EBookInstallSummary Contains properties for the installation summary of a book for a device.
type EBookInstallSummary struct {
	// Number of Devices that have failed to install this book.
	FailedDeviceCount *int32 `json:"failedDeviceCount,omitempty"`
	// Number of Users that have 1 or more device that failed to install this book.
	FailedUserCount *int32 `json:"failedUserCount,omitempty"`
	// Number of Devices that have successfully installed this book.
	InstalledDeviceCount *int32 `json:"installedDeviceCount,omitempty"`
	// Number of Users whose devices have all succeeded to install this book.
	InstalledUserCount *int32 `json:"installedUserCount,omitempty"`
	// Number of Devices that does not have this book installed.
	NotInstalledDeviceCount *int32 `json:"notInstalledDeviceCount,omitempty"`
	// Number of Users that did not install this book.
	NotInstalledUserCount *int32 `json:"notInstalledUserCount,omitempty"`
}

// NewEBookInstallSummary instantiates a new EBookInstallSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEBookInstallSummary() *EBookInstallSummary {
	this := EBookInstallSummary{}
	return &this
}

// NewEBookInstallSummaryWithDefaults instantiates a new EBookInstallSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEBookInstallSummaryWithDefaults() *EBookInstallSummary {
	this := EBookInstallSummary{}
	return &this
}

// GetFailedDeviceCount returns the FailedDeviceCount field value if set, zero value otherwise.
func (o *EBookInstallSummary) GetFailedDeviceCount() int32 {
	if o == nil || o.FailedDeviceCount == nil {
		var ret int32
		return ret
	}
	return *o.FailedDeviceCount
}

// GetFailedDeviceCountOk returns a tuple with the FailedDeviceCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EBookInstallSummary) GetFailedDeviceCountOk() (*int32, bool) {
	if o == nil || o.FailedDeviceCount == nil {
		return nil, false
	}
	return o.FailedDeviceCount, true
}

// HasFailedDeviceCount returns a boolean if a field has been set.
func (o *EBookInstallSummary) HasFailedDeviceCount() bool {
	if o != nil && o.FailedDeviceCount != nil {
		return true
	}

	return false
}

// SetFailedDeviceCount gets a reference to the given int32 and assigns it to the FailedDeviceCount field.
func (o *EBookInstallSummary) SetFailedDeviceCount(v int32) {
	o.FailedDeviceCount = &v
}

// GetFailedUserCount returns the FailedUserCount field value if set, zero value otherwise.
func (o *EBookInstallSummary) GetFailedUserCount() int32 {
	if o == nil || o.FailedUserCount == nil {
		var ret int32
		return ret
	}
	return *o.FailedUserCount
}

// GetFailedUserCountOk returns a tuple with the FailedUserCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EBookInstallSummary) GetFailedUserCountOk() (*int32, bool) {
	if o == nil || o.FailedUserCount == nil {
		return nil, false
	}
	return o.FailedUserCount, true
}

// HasFailedUserCount returns a boolean if a field has been set.
func (o *EBookInstallSummary) HasFailedUserCount() bool {
	if o != nil && o.FailedUserCount != nil {
		return true
	}

	return false
}

// SetFailedUserCount gets a reference to the given int32 and assigns it to the FailedUserCount field.
func (o *EBookInstallSummary) SetFailedUserCount(v int32) {
	o.FailedUserCount = &v
}

// GetInstalledDeviceCount returns the InstalledDeviceCount field value if set, zero value otherwise.
func (o *EBookInstallSummary) GetInstalledDeviceCount() int32 {
	if o == nil || o.InstalledDeviceCount == nil {
		var ret int32
		return ret
	}
	return *o.InstalledDeviceCount
}

// GetInstalledDeviceCountOk returns a tuple with the InstalledDeviceCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EBookInstallSummary) GetInstalledDeviceCountOk() (*int32, bool) {
	if o == nil || o.InstalledDeviceCount == nil {
		return nil, false
	}
	return o.InstalledDeviceCount, true
}

// HasInstalledDeviceCount returns a boolean if a field has been set.
func (o *EBookInstallSummary) HasInstalledDeviceCount() bool {
	if o != nil && o.InstalledDeviceCount != nil {
		return true
	}

	return false
}

// SetInstalledDeviceCount gets a reference to the given int32 and assigns it to the InstalledDeviceCount field.
func (o *EBookInstallSummary) SetInstalledDeviceCount(v int32) {
	o.InstalledDeviceCount = &v
}

// GetInstalledUserCount returns the InstalledUserCount field value if set, zero value otherwise.
func (o *EBookInstallSummary) GetInstalledUserCount() int32 {
	if o == nil || o.InstalledUserCount == nil {
		var ret int32
		return ret
	}
	return *o.InstalledUserCount
}

// GetInstalledUserCountOk returns a tuple with the InstalledUserCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EBookInstallSummary) GetInstalledUserCountOk() (*int32, bool) {
	if o == nil || o.InstalledUserCount == nil {
		return nil, false
	}
	return o.InstalledUserCount, true
}

// HasInstalledUserCount returns a boolean if a field has been set.
func (o *EBookInstallSummary) HasInstalledUserCount() bool {
	if o != nil && o.InstalledUserCount != nil {
		return true
	}

	return false
}

// SetInstalledUserCount gets a reference to the given int32 and assigns it to the InstalledUserCount field.
func (o *EBookInstallSummary) SetInstalledUserCount(v int32) {
	o.InstalledUserCount = &v
}

// GetNotInstalledDeviceCount returns the NotInstalledDeviceCount field value if set, zero value otherwise.
func (o *EBookInstallSummary) GetNotInstalledDeviceCount() int32 {
	if o == nil || o.NotInstalledDeviceCount == nil {
		var ret int32
		return ret
	}
	return *o.NotInstalledDeviceCount
}

// GetNotInstalledDeviceCountOk returns a tuple with the NotInstalledDeviceCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EBookInstallSummary) GetNotInstalledDeviceCountOk() (*int32, bool) {
	if o == nil || o.NotInstalledDeviceCount == nil {
		return nil, false
	}
	return o.NotInstalledDeviceCount, true
}

// HasNotInstalledDeviceCount returns a boolean if a field has been set.
func (o *EBookInstallSummary) HasNotInstalledDeviceCount() bool {
	if o != nil && o.NotInstalledDeviceCount != nil {
		return true
	}

	return false
}

// SetNotInstalledDeviceCount gets a reference to the given int32 and assigns it to the NotInstalledDeviceCount field.
func (o *EBookInstallSummary) SetNotInstalledDeviceCount(v int32) {
	o.NotInstalledDeviceCount = &v
}

// GetNotInstalledUserCount returns the NotInstalledUserCount field value if set, zero value otherwise.
func (o *EBookInstallSummary) GetNotInstalledUserCount() int32 {
	if o == nil || o.NotInstalledUserCount == nil {
		var ret int32
		return ret
	}
	return *o.NotInstalledUserCount
}

// GetNotInstalledUserCountOk returns a tuple with the NotInstalledUserCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EBookInstallSummary) GetNotInstalledUserCountOk() (*int32, bool) {
	if o == nil || o.NotInstalledUserCount == nil {
		return nil, false
	}
	return o.NotInstalledUserCount, true
}

// HasNotInstalledUserCount returns a boolean if a field has been set.
func (o *EBookInstallSummary) HasNotInstalledUserCount() bool {
	if o != nil && o.NotInstalledUserCount != nil {
		return true
	}

	return false
}

// SetNotInstalledUserCount gets a reference to the given int32 and assigns it to the NotInstalledUserCount field.
func (o *EBookInstallSummary) SetNotInstalledUserCount(v int32) {
	o.NotInstalledUserCount = &v
}

func (o EBookInstallSummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.FailedDeviceCount != nil {
		toSerialize["failedDeviceCount"] = o.FailedDeviceCount
	}
	if o.FailedUserCount != nil {
		toSerialize["failedUserCount"] = o.FailedUserCount
	}
	if o.InstalledDeviceCount != nil {
		toSerialize["installedDeviceCount"] = o.InstalledDeviceCount
	}
	if o.InstalledUserCount != nil {
		toSerialize["installedUserCount"] = o.InstalledUserCount
	}
	if o.NotInstalledDeviceCount != nil {
		toSerialize["notInstalledDeviceCount"] = o.NotInstalledDeviceCount
	}
	if o.NotInstalledUserCount != nil {
		toSerialize["notInstalledUserCount"] = o.NotInstalledUserCount
	}
	return json.Marshal(toSerialize)
}

type NullableEBookInstallSummary struct {
	value *EBookInstallSummary
	isSet bool
}

func (v NullableEBookInstallSummary) Get() *EBookInstallSummary {
	return v.value
}

func (v *NullableEBookInstallSummary) Set(val *EBookInstallSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableEBookInstallSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableEBookInstallSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEBookInstallSummary(val *EBookInstallSummary) *NullableEBookInstallSummary {
	return &NullableEBookInstallSummary{value: val, isSet: true}
}

func (v NullableEBookInstallSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEBookInstallSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


