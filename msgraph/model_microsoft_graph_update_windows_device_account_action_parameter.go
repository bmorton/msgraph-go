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

// MicrosoftGraphUpdateWindowsDeviceAccountActionParameter struct for MicrosoftGraphUpdateWindowsDeviceAccountActionParameter
type MicrosoftGraphUpdateWindowsDeviceAccountActionParameter struct {
	// Not yet documented
	CalendarSyncEnabled NullableBool `json:"calendarSyncEnabled,omitempty"`
	// Not yet documented
	DeviceAccount AnyOfmicrosoftGraphWindowsDeviceAccount `json:"deviceAccount,omitempty"`
	// Not yet documented
	DeviceAccountEmail NullableString `json:"deviceAccountEmail,omitempty"`
	// Not yet documented
	ExchangeServer NullableString `json:"exchangeServer,omitempty"`
	// Not yet documented
	PasswordRotationEnabled NullableBool `json:"passwordRotationEnabled,omitempty"`
	// Not yet documented
	SessionInitiationProtocalAddress NullableString `json:"sessionInitiationProtocalAddress,omitempty"`
}

// NewMicrosoftGraphUpdateWindowsDeviceAccountActionParameter instantiates a new MicrosoftGraphUpdateWindowsDeviceAccountActionParameter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphUpdateWindowsDeviceAccountActionParameter() *MicrosoftGraphUpdateWindowsDeviceAccountActionParameter {
	this := MicrosoftGraphUpdateWindowsDeviceAccountActionParameter{}
	return &this
}

// NewMicrosoftGraphUpdateWindowsDeviceAccountActionParameterWithDefaults instantiates a new MicrosoftGraphUpdateWindowsDeviceAccountActionParameter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphUpdateWindowsDeviceAccountActionParameterWithDefaults() *MicrosoftGraphUpdateWindowsDeviceAccountActionParameter {
	this := MicrosoftGraphUpdateWindowsDeviceAccountActionParameter{}
	return &this
}

// GetCalendarSyncEnabled returns the CalendarSyncEnabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphUpdateWindowsDeviceAccountActionParameter) GetCalendarSyncEnabled() bool {
	if o == nil || o.CalendarSyncEnabled.Get() == nil {
		var ret bool
		return ret
	}
	return *o.CalendarSyncEnabled.Get()
}

// GetCalendarSyncEnabledOk returns a tuple with the CalendarSyncEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphUpdateWindowsDeviceAccountActionParameter) GetCalendarSyncEnabledOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CalendarSyncEnabled.Get(), o.CalendarSyncEnabled.IsSet()
}

// HasCalendarSyncEnabled returns a boolean if a field has been set.
func (o *MicrosoftGraphUpdateWindowsDeviceAccountActionParameter) HasCalendarSyncEnabled() bool {
	if o != nil && o.CalendarSyncEnabled.IsSet() {
		return true
	}

	return false
}

// SetCalendarSyncEnabled gets a reference to the given NullableBool and assigns it to the CalendarSyncEnabled field.
func (o *MicrosoftGraphUpdateWindowsDeviceAccountActionParameter) SetCalendarSyncEnabled(v bool) {
	o.CalendarSyncEnabled.Set(&v)
}
// SetCalendarSyncEnabledNil sets the value for CalendarSyncEnabled to be an explicit nil
func (o *MicrosoftGraphUpdateWindowsDeviceAccountActionParameter) SetCalendarSyncEnabledNil() {
	o.CalendarSyncEnabled.Set(nil)
}

// UnsetCalendarSyncEnabled ensures that no value is present for CalendarSyncEnabled, not even an explicit nil
func (o *MicrosoftGraphUpdateWindowsDeviceAccountActionParameter) UnsetCalendarSyncEnabled() {
	o.CalendarSyncEnabled.Unset()
}

// GetDeviceAccount returns the DeviceAccount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphUpdateWindowsDeviceAccountActionParameter) GetDeviceAccount() AnyOfmicrosoftGraphWindowsDeviceAccount {
	if o == nil  {
		var ret AnyOfmicrosoftGraphWindowsDeviceAccount
		return ret
	}
	return o.DeviceAccount
}

// GetDeviceAccountOk returns a tuple with the DeviceAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphUpdateWindowsDeviceAccountActionParameter) GetDeviceAccountOk() (*AnyOfmicrosoftGraphWindowsDeviceAccount, bool) {
	if o == nil || o.DeviceAccount == nil {
		return nil, false
	}
	return &o.DeviceAccount, true
}

// HasDeviceAccount returns a boolean if a field has been set.
func (o *MicrosoftGraphUpdateWindowsDeviceAccountActionParameter) HasDeviceAccount() bool {
	if o != nil && o.DeviceAccount != nil {
		return true
	}

	return false
}

// SetDeviceAccount gets a reference to the given AnyOfmicrosoftGraphWindowsDeviceAccount and assigns it to the DeviceAccount field.
func (o *MicrosoftGraphUpdateWindowsDeviceAccountActionParameter) SetDeviceAccount(v AnyOfmicrosoftGraphWindowsDeviceAccount) {
	o.DeviceAccount = v
}

// GetDeviceAccountEmail returns the DeviceAccountEmail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphUpdateWindowsDeviceAccountActionParameter) GetDeviceAccountEmail() string {
	if o == nil || o.DeviceAccountEmail.Get() == nil {
		var ret string
		return ret
	}
	return *o.DeviceAccountEmail.Get()
}

// GetDeviceAccountEmailOk returns a tuple with the DeviceAccountEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphUpdateWindowsDeviceAccountActionParameter) GetDeviceAccountEmailOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DeviceAccountEmail.Get(), o.DeviceAccountEmail.IsSet()
}

// HasDeviceAccountEmail returns a boolean if a field has been set.
func (o *MicrosoftGraphUpdateWindowsDeviceAccountActionParameter) HasDeviceAccountEmail() bool {
	if o != nil && o.DeviceAccountEmail.IsSet() {
		return true
	}

	return false
}

// SetDeviceAccountEmail gets a reference to the given NullableString and assigns it to the DeviceAccountEmail field.
func (o *MicrosoftGraphUpdateWindowsDeviceAccountActionParameter) SetDeviceAccountEmail(v string) {
	o.DeviceAccountEmail.Set(&v)
}
// SetDeviceAccountEmailNil sets the value for DeviceAccountEmail to be an explicit nil
func (o *MicrosoftGraphUpdateWindowsDeviceAccountActionParameter) SetDeviceAccountEmailNil() {
	o.DeviceAccountEmail.Set(nil)
}

// UnsetDeviceAccountEmail ensures that no value is present for DeviceAccountEmail, not even an explicit nil
func (o *MicrosoftGraphUpdateWindowsDeviceAccountActionParameter) UnsetDeviceAccountEmail() {
	o.DeviceAccountEmail.Unset()
}

// GetExchangeServer returns the ExchangeServer field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphUpdateWindowsDeviceAccountActionParameter) GetExchangeServer() string {
	if o == nil || o.ExchangeServer.Get() == nil {
		var ret string
		return ret
	}
	return *o.ExchangeServer.Get()
}

// GetExchangeServerOk returns a tuple with the ExchangeServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphUpdateWindowsDeviceAccountActionParameter) GetExchangeServerOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ExchangeServer.Get(), o.ExchangeServer.IsSet()
}

// HasExchangeServer returns a boolean if a field has been set.
func (o *MicrosoftGraphUpdateWindowsDeviceAccountActionParameter) HasExchangeServer() bool {
	if o != nil && o.ExchangeServer.IsSet() {
		return true
	}

	return false
}

// SetExchangeServer gets a reference to the given NullableString and assigns it to the ExchangeServer field.
func (o *MicrosoftGraphUpdateWindowsDeviceAccountActionParameter) SetExchangeServer(v string) {
	o.ExchangeServer.Set(&v)
}
// SetExchangeServerNil sets the value for ExchangeServer to be an explicit nil
func (o *MicrosoftGraphUpdateWindowsDeviceAccountActionParameter) SetExchangeServerNil() {
	o.ExchangeServer.Set(nil)
}

// UnsetExchangeServer ensures that no value is present for ExchangeServer, not even an explicit nil
func (o *MicrosoftGraphUpdateWindowsDeviceAccountActionParameter) UnsetExchangeServer() {
	o.ExchangeServer.Unset()
}

// GetPasswordRotationEnabled returns the PasswordRotationEnabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphUpdateWindowsDeviceAccountActionParameter) GetPasswordRotationEnabled() bool {
	if o == nil || o.PasswordRotationEnabled.Get() == nil {
		var ret bool
		return ret
	}
	return *o.PasswordRotationEnabled.Get()
}

// GetPasswordRotationEnabledOk returns a tuple with the PasswordRotationEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphUpdateWindowsDeviceAccountActionParameter) GetPasswordRotationEnabledOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PasswordRotationEnabled.Get(), o.PasswordRotationEnabled.IsSet()
}

// HasPasswordRotationEnabled returns a boolean if a field has been set.
func (o *MicrosoftGraphUpdateWindowsDeviceAccountActionParameter) HasPasswordRotationEnabled() bool {
	if o != nil && o.PasswordRotationEnabled.IsSet() {
		return true
	}

	return false
}

// SetPasswordRotationEnabled gets a reference to the given NullableBool and assigns it to the PasswordRotationEnabled field.
func (o *MicrosoftGraphUpdateWindowsDeviceAccountActionParameter) SetPasswordRotationEnabled(v bool) {
	o.PasswordRotationEnabled.Set(&v)
}
// SetPasswordRotationEnabledNil sets the value for PasswordRotationEnabled to be an explicit nil
func (o *MicrosoftGraphUpdateWindowsDeviceAccountActionParameter) SetPasswordRotationEnabledNil() {
	o.PasswordRotationEnabled.Set(nil)
}

// UnsetPasswordRotationEnabled ensures that no value is present for PasswordRotationEnabled, not even an explicit nil
func (o *MicrosoftGraphUpdateWindowsDeviceAccountActionParameter) UnsetPasswordRotationEnabled() {
	o.PasswordRotationEnabled.Unset()
}

// GetSessionInitiationProtocalAddress returns the SessionInitiationProtocalAddress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphUpdateWindowsDeviceAccountActionParameter) GetSessionInitiationProtocalAddress() string {
	if o == nil || o.SessionInitiationProtocalAddress.Get() == nil {
		var ret string
		return ret
	}
	return *o.SessionInitiationProtocalAddress.Get()
}

// GetSessionInitiationProtocalAddressOk returns a tuple with the SessionInitiationProtocalAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphUpdateWindowsDeviceAccountActionParameter) GetSessionInitiationProtocalAddressOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SessionInitiationProtocalAddress.Get(), o.SessionInitiationProtocalAddress.IsSet()
}

// HasSessionInitiationProtocalAddress returns a boolean if a field has been set.
func (o *MicrosoftGraphUpdateWindowsDeviceAccountActionParameter) HasSessionInitiationProtocalAddress() bool {
	if o != nil && o.SessionInitiationProtocalAddress.IsSet() {
		return true
	}

	return false
}

// SetSessionInitiationProtocalAddress gets a reference to the given NullableString and assigns it to the SessionInitiationProtocalAddress field.
func (o *MicrosoftGraphUpdateWindowsDeviceAccountActionParameter) SetSessionInitiationProtocalAddress(v string) {
	o.SessionInitiationProtocalAddress.Set(&v)
}
// SetSessionInitiationProtocalAddressNil sets the value for SessionInitiationProtocalAddress to be an explicit nil
func (o *MicrosoftGraphUpdateWindowsDeviceAccountActionParameter) SetSessionInitiationProtocalAddressNil() {
	o.SessionInitiationProtocalAddress.Set(nil)
}

// UnsetSessionInitiationProtocalAddress ensures that no value is present for SessionInitiationProtocalAddress, not even an explicit nil
func (o *MicrosoftGraphUpdateWindowsDeviceAccountActionParameter) UnsetSessionInitiationProtocalAddress() {
	o.SessionInitiationProtocalAddress.Unset()
}

func (o MicrosoftGraphUpdateWindowsDeviceAccountActionParameter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CalendarSyncEnabled.IsSet() {
		toSerialize["calendarSyncEnabled"] = o.CalendarSyncEnabled.Get()
	}
	if o.DeviceAccount != nil {
		toSerialize["deviceAccount"] = o.DeviceAccount
	}
	if o.DeviceAccountEmail.IsSet() {
		toSerialize["deviceAccountEmail"] = o.DeviceAccountEmail.Get()
	}
	if o.ExchangeServer.IsSet() {
		toSerialize["exchangeServer"] = o.ExchangeServer.Get()
	}
	if o.PasswordRotationEnabled.IsSet() {
		toSerialize["passwordRotationEnabled"] = o.PasswordRotationEnabled.Get()
	}
	if o.SessionInitiationProtocalAddress.IsSet() {
		toSerialize["sessionInitiationProtocalAddress"] = o.SessionInitiationProtocalAddress.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphUpdateWindowsDeviceAccountActionParameter struct {
	value *MicrosoftGraphUpdateWindowsDeviceAccountActionParameter
	isSet bool
}

func (v NullableMicrosoftGraphUpdateWindowsDeviceAccountActionParameter) Get() *MicrosoftGraphUpdateWindowsDeviceAccountActionParameter {
	return v.value
}

func (v *NullableMicrosoftGraphUpdateWindowsDeviceAccountActionParameter) Set(val *MicrosoftGraphUpdateWindowsDeviceAccountActionParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphUpdateWindowsDeviceAccountActionParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphUpdateWindowsDeviceAccountActionParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphUpdateWindowsDeviceAccountActionParameter(val *MicrosoftGraphUpdateWindowsDeviceAccountActionParameter) *NullableMicrosoftGraphUpdateWindowsDeviceAccountActionParameter {
	return &NullableMicrosoftGraphUpdateWindowsDeviceAccountActionParameter{value: val, isSet: true}
}

func (v NullableMicrosoftGraphUpdateWindowsDeviceAccountActionParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphUpdateWindowsDeviceAccountActionParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


