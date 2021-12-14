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

// MicrosoftGraphWindowsDeviceAccount struct for MicrosoftGraphWindowsDeviceAccount
type MicrosoftGraphWindowsDeviceAccount struct {
	// Not yet documented
	Password NullableString `json:"password,omitempty"`
}

// NewMicrosoftGraphWindowsDeviceAccount instantiates a new MicrosoftGraphWindowsDeviceAccount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphWindowsDeviceAccount() *MicrosoftGraphWindowsDeviceAccount {
	this := MicrosoftGraphWindowsDeviceAccount{}
	return &this
}

// NewMicrosoftGraphWindowsDeviceAccountWithDefaults instantiates a new MicrosoftGraphWindowsDeviceAccount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphWindowsDeviceAccountWithDefaults() *MicrosoftGraphWindowsDeviceAccount {
	this := MicrosoftGraphWindowsDeviceAccount{}
	return &this
}

// GetPassword returns the Password field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphWindowsDeviceAccount) GetPassword() string {
	if o == nil || o.Password.Get() == nil {
		var ret string
		return ret
	}
	return *o.Password.Get()
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphWindowsDeviceAccount) GetPasswordOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Password.Get(), o.Password.IsSet()
}

// HasPassword returns a boolean if a field has been set.
func (o *MicrosoftGraphWindowsDeviceAccount) HasPassword() bool {
	if o != nil && o.Password.IsSet() {
		return true
	}

	return false
}

// SetPassword gets a reference to the given NullableString and assigns it to the Password field.
func (o *MicrosoftGraphWindowsDeviceAccount) SetPassword(v string) {
	o.Password.Set(&v)
}
// SetPasswordNil sets the value for Password to be an explicit nil
func (o *MicrosoftGraphWindowsDeviceAccount) SetPasswordNil() {
	o.Password.Set(nil)
}

// UnsetPassword ensures that no value is present for Password, not even an explicit nil
func (o *MicrosoftGraphWindowsDeviceAccount) UnsetPassword() {
	o.Password.Unset()
}

func (o MicrosoftGraphWindowsDeviceAccount) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Password.IsSet() {
		toSerialize["password"] = o.Password.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphWindowsDeviceAccount struct {
	value *MicrosoftGraphWindowsDeviceAccount
	isSet bool
}

func (v NullableMicrosoftGraphWindowsDeviceAccount) Get() *MicrosoftGraphWindowsDeviceAccount {
	return v.value
}

func (v *NullableMicrosoftGraphWindowsDeviceAccount) Set(val *MicrosoftGraphWindowsDeviceAccount) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphWindowsDeviceAccount) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphWindowsDeviceAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphWindowsDeviceAccount(val *MicrosoftGraphWindowsDeviceAccount) *NullableMicrosoftGraphWindowsDeviceAccount {
	return &NullableMicrosoftGraphWindowsDeviceAccount{value: val, isSet: true}
}

func (v NullableMicrosoftGraphWindowsDeviceAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphWindowsDeviceAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

