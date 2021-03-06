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

// DeviceComplianceSettingState Device compliance setting State for a given device.
type DeviceComplianceSettingState struct {
	// The DateTime when device compliance grace period expires
	ComplianceGracePeriodExpirationDateTime *time.Time `json:"complianceGracePeriodExpirationDateTime,omitempty"`
	// The Device Id that is being reported
	DeviceId NullableString `json:"deviceId,omitempty"`
	// The device model that is being reported
	DeviceModel NullableString `json:"deviceModel,omitempty"`
	// The Device Name that is being reported
	DeviceName NullableString `json:"deviceName,omitempty"`
	// The setting class name and property name.
	Setting NullableString `json:"setting,omitempty"`
	// The Setting Name that is being reported
	SettingName NullableString `json:"settingName,omitempty"`
	// The compliance state of the setting. Possible values are: unknown, notApplicable, compliant, remediated, nonCompliant, error, conflict, notAssigned.
	State AnyOfmicrosoftGraphComplianceStatus `json:"state,omitempty"`
	// The User email address that is being reported
	UserEmail NullableString `json:"userEmail,omitempty"`
	// The user Id that is being reported
	UserId NullableString `json:"userId,omitempty"`
	// The User Name that is being reported
	UserName NullableString `json:"userName,omitempty"`
	// The User PrincipalName that is being reported
	UserPrincipalName NullableString `json:"userPrincipalName,omitempty"`
}

// NewDeviceComplianceSettingState instantiates a new DeviceComplianceSettingState object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceComplianceSettingState() *DeviceComplianceSettingState {
	this := DeviceComplianceSettingState{}
	return &this
}

// NewDeviceComplianceSettingStateWithDefaults instantiates a new DeviceComplianceSettingState object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceComplianceSettingStateWithDefaults() *DeviceComplianceSettingState {
	this := DeviceComplianceSettingState{}
	return &this
}

// GetComplianceGracePeriodExpirationDateTime returns the ComplianceGracePeriodExpirationDateTime field value if set, zero value otherwise.
func (o *DeviceComplianceSettingState) GetComplianceGracePeriodExpirationDateTime() time.Time {
	if o == nil || o.ComplianceGracePeriodExpirationDateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.ComplianceGracePeriodExpirationDateTime
}

// GetComplianceGracePeriodExpirationDateTimeOk returns a tuple with the ComplianceGracePeriodExpirationDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceComplianceSettingState) GetComplianceGracePeriodExpirationDateTimeOk() (*time.Time, bool) {
	if o == nil || o.ComplianceGracePeriodExpirationDateTime == nil {
		return nil, false
	}
	return o.ComplianceGracePeriodExpirationDateTime, true
}

// HasComplianceGracePeriodExpirationDateTime returns a boolean if a field has been set.
func (o *DeviceComplianceSettingState) HasComplianceGracePeriodExpirationDateTime() bool {
	if o != nil && o.ComplianceGracePeriodExpirationDateTime != nil {
		return true
	}

	return false
}

// SetComplianceGracePeriodExpirationDateTime gets a reference to the given time.Time and assigns it to the ComplianceGracePeriodExpirationDateTime field.
func (o *DeviceComplianceSettingState) SetComplianceGracePeriodExpirationDateTime(v time.Time) {
	o.ComplianceGracePeriodExpirationDateTime = &v
}

// GetDeviceId returns the DeviceId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeviceComplianceSettingState) GetDeviceId() string {
	if o == nil || o.DeviceId.Get() == nil {
		var ret string
		return ret
	}
	return *o.DeviceId.Get()
}

// GetDeviceIdOk returns a tuple with the DeviceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeviceComplianceSettingState) GetDeviceIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DeviceId.Get(), o.DeviceId.IsSet()
}

// HasDeviceId returns a boolean if a field has been set.
func (o *DeviceComplianceSettingState) HasDeviceId() bool {
	if o != nil && o.DeviceId.IsSet() {
		return true
	}

	return false
}

// SetDeviceId gets a reference to the given NullableString and assigns it to the DeviceId field.
func (o *DeviceComplianceSettingState) SetDeviceId(v string) {
	o.DeviceId.Set(&v)
}
// SetDeviceIdNil sets the value for DeviceId to be an explicit nil
func (o *DeviceComplianceSettingState) SetDeviceIdNil() {
	o.DeviceId.Set(nil)
}

// UnsetDeviceId ensures that no value is present for DeviceId, not even an explicit nil
func (o *DeviceComplianceSettingState) UnsetDeviceId() {
	o.DeviceId.Unset()
}

// GetDeviceModel returns the DeviceModel field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeviceComplianceSettingState) GetDeviceModel() string {
	if o == nil || o.DeviceModel.Get() == nil {
		var ret string
		return ret
	}
	return *o.DeviceModel.Get()
}

// GetDeviceModelOk returns a tuple with the DeviceModel field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeviceComplianceSettingState) GetDeviceModelOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DeviceModel.Get(), o.DeviceModel.IsSet()
}

// HasDeviceModel returns a boolean if a field has been set.
func (o *DeviceComplianceSettingState) HasDeviceModel() bool {
	if o != nil && o.DeviceModel.IsSet() {
		return true
	}

	return false
}

// SetDeviceModel gets a reference to the given NullableString and assigns it to the DeviceModel field.
func (o *DeviceComplianceSettingState) SetDeviceModel(v string) {
	o.DeviceModel.Set(&v)
}
// SetDeviceModelNil sets the value for DeviceModel to be an explicit nil
func (o *DeviceComplianceSettingState) SetDeviceModelNil() {
	o.DeviceModel.Set(nil)
}

// UnsetDeviceModel ensures that no value is present for DeviceModel, not even an explicit nil
func (o *DeviceComplianceSettingState) UnsetDeviceModel() {
	o.DeviceModel.Unset()
}

// GetDeviceName returns the DeviceName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeviceComplianceSettingState) GetDeviceName() string {
	if o == nil || o.DeviceName.Get() == nil {
		var ret string
		return ret
	}
	return *o.DeviceName.Get()
}

// GetDeviceNameOk returns a tuple with the DeviceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeviceComplianceSettingState) GetDeviceNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DeviceName.Get(), o.DeviceName.IsSet()
}

// HasDeviceName returns a boolean if a field has been set.
func (o *DeviceComplianceSettingState) HasDeviceName() bool {
	if o != nil && o.DeviceName.IsSet() {
		return true
	}

	return false
}

// SetDeviceName gets a reference to the given NullableString and assigns it to the DeviceName field.
func (o *DeviceComplianceSettingState) SetDeviceName(v string) {
	o.DeviceName.Set(&v)
}
// SetDeviceNameNil sets the value for DeviceName to be an explicit nil
func (o *DeviceComplianceSettingState) SetDeviceNameNil() {
	o.DeviceName.Set(nil)
}

// UnsetDeviceName ensures that no value is present for DeviceName, not even an explicit nil
func (o *DeviceComplianceSettingState) UnsetDeviceName() {
	o.DeviceName.Unset()
}

// GetSetting returns the Setting field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeviceComplianceSettingState) GetSetting() string {
	if o == nil || o.Setting.Get() == nil {
		var ret string
		return ret
	}
	return *o.Setting.Get()
}

// GetSettingOk returns a tuple with the Setting field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeviceComplianceSettingState) GetSettingOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Setting.Get(), o.Setting.IsSet()
}

// HasSetting returns a boolean if a field has been set.
func (o *DeviceComplianceSettingState) HasSetting() bool {
	if o != nil && o.Setting.IsSet() {
		return true
	}

	return false
}

// SetSetting gets a reference to the given NullableString and assigns it to the Setting field.
func (o *DeviceComplianceSettingState) SetSetting(v string) {
	o.Setting.Set(&v)
}
// SetSettingNil sets the value for Setting to be an explicit nil
func (o *DeviceComplianceSettingState) SetSettingNil() {
	o.Setting.Set(nil)
}

// UnsetSetting ensures that no value is present for Setting, not even an explicit nil
func (o *DeviceComplianceSettingState) UnsetSetting() {
	o.Setting.Unset()
}

// GetSettingName returns the SettingName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeviceComplianceSettingState) GetSettingName() string {
	if o == nil || o.SettingName.Get() == nil {
		var ret string
		return ret
	}
	return *o.SettingName.Get()
}

// GetSettingNameOk returns a tuple with the SettingName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeviceComplianceSettingState) GetSettingNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SettingName.Get(), o.SettingName.IsSet()
}

// HasSettingName returns a boolean if a field has been set.
func (o *DeviceComplianceSettingState) HasSettingName() bool {
	if o != nil && o.SettingName.IsSet() {
		return true
	}

	return false
}

// SetSettingName gets a reference to the given NullableString and assigns it to the SettingName field.
func (o *DeviceComplianceSettingState) SetSettingName(v string) {
	o.SettingName.Set(&v)
}
// SetSettingNameNil sets the value for SettingName to be an explicit nil
func (o *DeviceComplianceSettingState) SetSettingNameNil() {
	o.SettingName.Set(nil)
}

// UnsetSettingName ensures that no value is present for SettingName, not even an explicit nil
func (o *DeviceComplianceSettingState) UnsetSettingName() {
	o.SettingName.Unset()
}

// GetState returns the State field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeviceComplianceSettingState) GetState() AnyOfmicrosoftGraphComplianceStatus {
	if o == nil  {
		var ret AnyOfmicrosoftGraphComplianceStatus
		return ret
	}
	return o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeviceComplianceSettingState) GetStateOk() (*AnyOfmicrosoftGraphComplianceStatus, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return &o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *DeviceComplianceSettingState) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given AnyOfmicrosoftGraphComplianceStatus and assigns it to the State field.
func (o *DeviceComplianceSettingState) SetState(v AnyOfmicrosoftGraphComplianceStatus) {
	o.State = v
}

// GetUserEmail returns the UserEmail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeviceComplianceSettingState) GetUserEmail() string {
	if o == nil || o.UserEmail.Get() == nil {
		var ret string
		return ret
	}
	return *o.UserEmail.Get()
}

// GetUserEmailOk returns a tuple with the UserEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeviceComplianceSettingState) GetUserEmailOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.UserEmail.Get(), o.UserEmail.IsSet()
}

// HasUserEmail returns a boolean if a field has been set.
func (o *DeviceComplianceSettingState) HasUserEmail() bool {
	if o != nil && o.UserEmail.IsSet() {
		return true
	}

	return false
}

// SetUserEmail gets a reference to the given NullableString and assigns it to the UserEmail field.
func (o *DeviceComplianceSettingState) SetUserEmail(v string) {
	o.UserEmail.Set(&v)
}
// SetUserEmailNil sets the value for UserEmail to be an explicit nil
func (o *DeviceComplianceSettingState) SetUserEmailNil() {
	o.UserEmail.Set(nil)
}

// UnsetUserEmail ensures that no value is present for UserEmail, not even an explicit nil
func (o *DeviceComplianceSettingState) UnsetUserEmail() {
	o.UserEmail.Unset()
}

// GetUserId returns the UserId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeviceComplianceSettingState) GetUserId() string {
	if o == nil || o.UserId.Get() == nil {
		var ret string
		return ret
	}
	return *o.UserId.Get()
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeviceComplianceSettingState) GetUserIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.UserId.Get(), o.UserId.IsSet()
}

// HasUserId returns a boolean if a field has been set.
func (o *DeviceComplianceSettingState) HasUserId() bool {
	if o != nil && o.UserId.IsSet() {
		return true
	}

	return false
}

// SetUserId gets a reference to the given NullableString and assigns it to the UserId field.
func (o *DeviceComplianceSettingState) SetUserId(v string) {
	o.UserId.Set(&v)
}
// SetUserIdNil sets the value for UserId to be an explicit nil
func (o *DeviceComplianceSettingState) SetUserIdNil() {
	o.UserId.Set(nil)
}

// UnsetUserId ensures that no value is present for UserId, not even an explicit nil
func (o *DeviceComplianceSettingState) UnsetUserId() {
	o.UserId.Unset()
}

// GetUserName returns the UserName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeviceComplianceSettingState) GetUserName() string {
	if o == nil || o.UserName.Get() == nil {
		var ret string
		return ret
	}
	return *o.UserName.Get()
}

// GetUserNameOk returns a tuple with the UserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeviceComplianceSettingState) GetUserNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.UserName.Get(), o.UserName.IsSet()
}

// HasUserName returns a boolean if a field has been set.
func (o *DeviceComplianceSettingState) HasUserName() bool {
	if o != nil && o.UserName.IsSet() {
		return true
	}

	return false
}

// SetUserName gets a reference to the given NullableString and assigns it to the UserName field.
func (o *DeviceComplianceSettingState) SetUserName(v string) {
	o.UserName.Set(&v)
}
// SetUserNameNil sets the value for UserName to be an explicit nil
func (o *DeviceComplianceSettingState) SetUserNameNil() {
	o.UserName.Set(nil)
}

// UnsetUserName ensures that no value is present for UserName, not even an explicit nil
func (o *DeviceComplianceSettingState) UnsetUserName() {
	o.UserName.Unset()
}

// GetUserPrincipalName returns the UserPrincipalName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeviceComplianceSettingState) GetUserPrincipalName() string {
	if o == nil || o.UserPrincipalName.Get() == nil {
		var ret string
		return ret
	}
	return *o.UserPrincipalName.Get()
}

// GetUserPrincipalNameOk returns a tuple with the UserPrincipalName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeviceComplianceSettingState) GetUserPrincipalNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.UserPrincipalName.Get(), o.UserPrincipalName.IsSet()
}

// HasUserPrincipalName returns a boolean if a field has been set.
func (o *DeviceComplianceSettingState) HasUserPrincipalName() bool {
	if o != nil && o.UserPrincipalName.IsSet() {
		return true
	}

	return false
}

// SetUserPrincipalName gets a reference to the given NullableString and assigns it to the UserPrincipalName field.
func (o *DeviceComplianceSettingState) SetUserPrincipalName(v string) {
	o.UserPrincipalName.Set(&v)
}
// SetUserPrincipalNameNil sets the value for UserPrincipalName to be an explicit nil
func (o *DeviceComplianceSettingState) SetUserPrincipalNameNil() {
	o.UserPrincipalName.Set(nil)
}

// UnsetUserPrincipalName ensures that no value is present for UserPrincipalName, not even an explicit nil
func (o *DeviceComplianceSettingState) UnsetUserPrincipalName() {
	o.UserPrincipalName.Unset()
}

func (o DeviceComplianceSettingState) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ComplianceGracePeriodExpirationDateTime != nil {
		toSerialize["complianceGracePeriodExpirationDateTime"] = o.ComplianceGracePeriodExpirationDateTime
	}
	if o.DeviceId.IsSet() {
		toSerialize["deviceId"] = o.DeviceId.Get()
	}
	if o.DeviceModel.IsSet() {
		toSerialize["deviceModel"] = o.DeviceModel.Get()
	}
	if o.DeviceName.IsSet() {
		toSerialize["deviceName"] = o.DeviceName.Get()
	}
	if o.Setting.IsSet() {
		toSerialize["setting"] = o.Setting.Get()
	}
	if o.SettingName.IsSet() {
		toSerialize["settingName"] = o.SettingName.Get()
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.UserEmail.IsSet() {
		toSerialize["userEmail"] = o.UserEmail.Get()
	}
	if o.UserId.IsSet() {
		toSerialize["userId"] = o.UserId.Get()
	}
	if o.UserName.IsSet() {
		toSerialize["userName"] = o.UserName.Get()
	}
	if o.UserPrincipalName.IsSet() {
		toSerialize["userPrincipalName"] = o.UserPrincipalName.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableDeviceComplianceSettingState struct {
	value *DeviceComplianceSettingState
	isSet bool
}

func (v NullableDeviceComplianceSettingState) Get() *DeviceComplianceSettingState {
	return v.value
}

func (v *NullableDeviceComplianceSettingState) Set(val *DeviceComplianceSettingState) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceComplianceSettingState) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceComplianceSettingState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceComplianceSettingState(val *DeviceComplianceSettingState) *NullableDeviceComplianceSettingState {
	return &NullableDeviceComplianceSettingState{value: val, isSet: true}
}

func (v NullableDeviceComplianceSettingState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceComplianceSettingState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


