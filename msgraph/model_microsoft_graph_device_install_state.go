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

// MicrosoftGraphDeviceInstallState struct for MicrosoftGraphDeviceInstallState
type MicrosoftGraphDeviceInstallState struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
	// Device Id.
	DeviceId NullableString `json:"deviceId,omitempty"`
	// Device name.
	DeviceName NullableString `json:"deviceName,omitempty"`
	// The error code for install failures.
	ErrorCode NullableString `json:"errorCode,omitempty"`
	// The install state of the eBook. Possible values are: notApplicable, installed, failed, notInstalled, uninstallFailed, unknown.
	InstallState AnyOfmicrosoftGraphInstallState `json:"installState,omitempty"`
	// Last sync date and time.
	LastSyncDateTime *time.Time `json:"lastSyncDateTime,omitempty"`
	// OS Description.
	OsDescription NullableString `json:"osDescription,omitempty"`
	// OS Version.
	OsVersion NullableString `json:"osVersion,omitempty"`
	// Device User Name.
	UserName NullableString `json:"userName,omitempty"`
}

// NewMicrosoftGraphDeviceInstallState instantiates a new MicrosoftGraphDeviceInstallState object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphDeviceInstallState() *MicrosoftGraphDeviceInstallState {
	this := MicrosoftGraphDeviceInstallState{}
	return &this
}

// NewMicrosoftGraphDeviceInstallStateWithDefaults instantiates a new MicrosoftGraphDeviceInstallState object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphDeviceInstallStateWithDefaults() *MicrosoftGraphDeviceInstallState {
	this := MicrosoftGraphDeviceInstallState{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MicrosoftGraphDeviceInstallState) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphDeviceInstallState) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphDeviceInstallState) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MicrosoftGraphDeviceInstallState) SetId(v string) {
	o.Id = &v
}

// GetDeviceId returns the DeviceId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphDeviceInstallState) GetDeviceId() string {
	if o == nil || o.DeviceId.Get() == nil {
		var ret string
		return ret
	}
	return *o.DeviceId.Get()
}

// GetDeviceIdOk returns a tuple with the DeviceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphDeviceInstallState) GetDeviceIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DeviceId.Get(), o.DeviceId.IsSet()
}

// HasDeviceId returns a boolean if a field has been set.
func (o *MicrosoftGraphDeviceInstallState) HasDeviceId() bool {
	if o != nil && o.DeviceId.IsSet() {
		return true
	}

	return false
}

// SetDeviceId gets a reference to the given NullableString and assigns it to the DeviceId field.
func (o *MicrosoftGraphDeviceInstallState) SetDeviceId(v string) {
	o.DeviceId.Set(&v)
}
// SetDeviceIdNil sets the value for DeviceId to be an explicit nil
func (o *MicrosoftGraphDeviceInstallState) SetDeviceIdNil() {
	o.DeviceId.Set(nil)
}

// UnsetDeviceId ensures that no value is present for DeviceId, not even an explicit nil
func (o *MicrosoftGraphDeviceInstallState) UnsetDeviceId() {
	o.DeviceId.Unset()
}

// GetDeviceName returns the DeviceName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphDeviceInstallState) GetDeviceName() string {
	if o == nil || o.DeviceName.Get() == nil {
		var ret string
		return ret
	}
	return *o.DeviceName.Get()
}

// GetDeviceNameOk returns a tuple with the DeviceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphDeviceInstallState) GetDeviceNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DeviceName.Get(), o.DeviceName.IsSet()
}

// HasDeviceName returns a boolean if a field has been set.
func (o *MicrosoftGraphDeviceInstallState) HasDeviceName() bool {
	if o != nil && o.DeviceName.IsSet() {
		return true
	}

	return false
}

// SetDeviceName gets a reference to the given NullableString and assigns it to the DeviceName field.
func (o *MicrosoftGraphDeviceInstallState) SetDeviceName(v string) {
	o.DeviceName.Set(&v)
}
// SetDeviceNameNil sets the value for DeviceName to be an explicit nil
func (o *MicrosoftGraphDeviceInstallState) SetDeviceNameNil() {
	o.DeviceName.Set(nil)
}

// UnsetDeviceName ensures that no value is present for DeviceName, not even an explicit nil
func (o *MicrosoftGraphDeviceInstallState) UnsetDeviceName() {
	o.DeviceName.Unset()
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphDeviceInstallState) GetErrorCode() string {
	if o == nil || o.ErrorCode.Get() == nil {
		var ret string
		return ret
	}
	return *o.ErrorCode.Get()
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphDeviceInstallState) GetErrorCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ErrorCode.Get(), o.ErrorCode.IsSet()
}

// HasErrorCode returns a boolean if a field has been set.
func (o *MicrosoftGraphDeviceInstallState) HasErrorCode() bool {
	if o != nil && o.ErrorCode.IsSet() {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given NullableString and assigns it to the ErrorCode field.
func (o *MicrosoftGraphDeviceInstallState) SetErrorCode(v string) {
	o.ErrorCode.Set(&v)
}
// SetErrorCodeNil sets the value for ErrorCode to be an explicit nil
func (o *MicrosoftGraphDeviceInstallState) SetErrorCodeNil() {
	o.ErrorCode.Set(nil)
}

// UnsetErrorCode ensures that no value is present for ErrorCode, not even an explicit nil
func (o *MicrosoftGraphDeviceInstallState) UnsetErrorCode() {
	o.ErrorCode.Unset()
}

// GetInstallState returns the InstallState field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphDeviceInstallState) GetInstallState() AnyOfmicrosoftGraphInstallState {
	if o == nil  {
		var ret AnyOfmicrosoftGraphInstallState
		return ret
	}
	return o.InstallState
}

// GetInstallStateOk returns a tuple with the InstallState field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphDeviceInstallState) GetInstallStateOk() (*AnyOfmicrosoftGraphInstallState, bool) {
	if o == nil || o.InstallState == nil {
		return nil, false
	}
	return &o.InstallState, true
}

// HasInstallState returns a boolean if a field has been set.
func (o *MicrosoftGraphDeviceInstallState) HasInstallState() bool {
	if o != nil && o.InstallState != nil {
		return true
	}

	return false
}

// SetInstallState gets a reference to the given AnyOfmicrosoftGraphInstallState and assigns it to the InstallState field.
func (o *MicrosoftGraphDeviceInstallState) SetInstallState(v AnyOfmicrosoftGraphInstallState) {
	o.InstallState = v
}

// GetLastSyncDateTime returns the LastSyncDateTime field value if set, zero value otherwise.
func (o *MicrosoftGraphDeviceInstallState) GetLastSyncDateTime() time.Time {
	if o == nil || o.LastSyncDateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.LastSyncDateTime
}

// GetLastSyncDateTimeOk returns a tuple with the LastSyncDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphDeviceInstallState) GetLastSyncDateTimeOk() (*time.Time, bool) {
	if o == nil || o.LastSyncDateTime == nil {
		return nil, false
	}
	return o.LastSyncDateTime, true
}

// HasLastSyncDateTime returns a boolean if a field has been set.
func (o *MicrosoftGraphDeviceInstallState) HasLastSyncDateTime() bool {
	if o != nil && o.LastSyncDateTime != nil {
		return true
	}

	return false
}

// SetLastSyncDateTime gets a reference to the given time.Time and assigns it to the LastSyncDateTime field.
func (o *MicrosoftGraphDeviceInstallState) SetLastSyncDateTime(v time.Time) {
	o.LastSyncDateTime = &v
}

// GetOsDescription returns the OsDescription field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphDeviceInstallState) GetOsDescription() string {
	if o == nil || o.OsDescription.Get() == nil {
		var ret string
		return ret
	}
	return *o.OsDescription.Get()
}

// GetOsDescriptionOk returns a tuple with the OsDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphDeviceInstallState) GetOsDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.OsDescription.Get(), o.OsDescription.IsSet()
}

// HasOsDescription returns a boolean if a field has been set.
func (o *MicrosoftGraphDeviceInstallState) HasOsDescription() bool {
	if o != nil && o.OsDescription.IsSet() {
		return true
	}

	return false
}

// SetOsDescription gets a reference to the given NullableString and assigns it to the OsDescription field.
func (o *MicrosoftGraphDeviceInstallState) SetOsDescription(v string) {
	o.OsDescription.Set(&v)
}
// SetOsDescriptionNil sets the value for OsDescription to be an explicit nil
func (o *MicrosoftGraphDeviceInstallState) SetOsDescriptionNil() {
	o.OsDescription.Set(nil)
}

// UnsetOsDescription ensures that no value is present for OsDescription, not even an explicit nil
func (o *MicrosoftGraphDeviceInstallState) UnsetOsDescription() {
	o.OsDescription.Unset()
}

// GetOsVersion returns the OsVersion field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphDeviceInstallState) GetOsVersion() string {
	if o == nil || o.OsVersion.Get() == nil {
		var ret string
		return ret
	}
	return *o.OsVersion.Get()
}

// GetOsVersionOk returns a tuple with the OsVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphDeviceInstallState) GetOsVersionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.OsVersion.Get(), o.OsVersion.IsSet()
}

// HasOsVersion returns a boolean if a field has been set.
func (o *MicrosoftGraphDeviceInstallState) HasOsVersion() bool {
	if o != nil && o.OsVersion.IsSet() {
		return true
	}

	return false
}

// SetOsVersion gets a reference to the given NullableString and assigns it to the OsVersion field.
func (o *MicrosoftGraphDeviceInstallState) SetOsVersion(v string) {
	o.OsVersion.Set(&v)
}
// SetOsVersionNil sets the value for OsVersion to be an explicit nil
func (o *MicrosoftGraphDeviceInstallState) SetOsVersionNil() {
	o.OsVersion.Set(nil)
}

// UnsetOsVersion ensures that no value is present for OsVersion, not even an explicit nil
func (o *MicrosoftGraphDeviceInstallState) UnsetOsVersion() {
	o.OsVersion.Unset()
}

// GetUserName returns the UserName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphDeviceInstallState) GetUserName() string {
	if o == nil || o.UserName.Get() == nil {
		var ret string
		return ret
	}
	return *o.UserName.Get()
}

// GetUserNameOk returns a tuple with the UserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphDeviceInstallState) GetUserNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.UserName.Get(), o.UserName.IsSet()
}

// HasUserName returns a boolean if a field has been set.
func (o *MicrosoftGraphDeviceInstallState) HasUserName() bool {
	if o != nil && o.UserName.IsSet() {
		return true
	}

	return false
}

// SetUserName gets a reference to the given NullableString and assigns it to the UserName field.
func (o *MicrosoftGraphDeviceInstallState) SetUserName(v string) {
	o.UserName.Set(&v)
}
// SetUserNameNil sets the value for UserName to be an explicit nil
func (o *MicrosoftGraphDeviceInstallState) SetUserNameNil() {
	o.UserName.Set(nil)
}

// UnsetUserName ensures that no value is present for UserName, not even an explicit nil
func (o *MicrosoftGraphDeviceInstallState) UnsetUserName() {
	o.UserName.Unset()
}

func (o MicrosoftGraphDeviceInstallState) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.DeviceId.IsSet() {
		toSerialize["deviceId"] = o.DeviceId.Get()
	}
	if o.DeviceName.IsSet() {
		toSerialize["deviceName"] = o.DeviceName.Get()
	}
	if o.ErrorCode.IsSet() {
		toSerialize["errorCode"] = o.ErrorCode.Get()
	}
	if o.InstallState != nil {
		toSerialize["installState"] = o.InstallState
	}
	if o.LastSyncDateTime != nil {
		toSerialize["lastSyncDateTime"] = o.LastSyncDateTime
	}
	if o.OsDescription.IsSet() {
		toSerialize["osDescription"] = o.OsDescription.Get()
	}
	if o.OsVersion.IsSet() {
		toSerialize["osVersion"] = o.OsVersion.Get()
	}
	if o.UserName.IsSet() {
		toSerialize["userName"] = o.UserName.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphDeviceInstallState struct {
	value *MicrosoftGraphDeviceInstallState
	isSet bool
}

func (v NullableMicrosoftGraphDeviceInstallState) Get() *MicrosoftGraphDeviceInstallState {
	return v.value
}

func (v *NullableMicrosoftGraphDeviceInstallState) Set(val *MicrosoftGraphDeviceInstallState) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphDeviceInstallState) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphDeviceInstallState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphDeviceInstallState(val *MicrosoftGraphDeviceInstallState) *NullableMicrosoftGraphDeviceInstallState {
	return &NullableMicrosoftGraphDeviceInstallState{value: val, isSet: true}
}

func (v NullableMicrosoftGraphDeviceInstallState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphDeviceInstallState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


