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

// InlineObject586 struct for InlineObject586
type InlineObject586 struct {
	CurrentPassword NullableString `json:"currentPassword,omitempty"`
	NewPassword NullableString `json:"newPassword,omitempty"`
}

// NewInlineObject586 instantiates a new InlineObject586 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject586() *InlineObject586 {
	this := InlineObject586{}
	return &this
}

// NewInlineObject586WithDefaults instantiates a new InlineObject586 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject586WithDefaults() *InlineObject586 {
	this := InlineObject586{}
	return &this
}

// GetCurrentPassword returns the CurrentPassword field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject586) GetCurrentPassword() string {
	if o == nil || o.CurrentPassword.Get() == nil {
		var ret string
		return ret
	}
	return *o.CurrentPassword.Get()
}

// GetCurrentPasswordOk returns a tuple with the CurrentPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject586) GetCurrentPasswordOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CurrentPassword.Get(), o.CurrentPassword.IsSet()
}

// HasCurrentPassword returns a boolean if a field has been set.
func (o *InlineObject586) HasCurrentPassword() bool {
	if o != nil && o.CurrentPassword.IsSet() {
		return true
	}

	return false
}

// SetCurrentPassword gets a reference to the given NullableString and assigns it to the CurrentPassword field.
func (o *InlineObject586) SetCurrentPassword(v string) {
	o.CurrentPassword.Set(&v)
}
// SetCurrentPasswordNil sets the value for CurrentPassword to be an explicit nil
func (o *InlineObject586) SetCurrentPasswordNil() {
	o.CurrentPassword.Set(nil)
}

// UnsetCurrentPassword ensures that no value is present for CurrentPassword, not even an explicit nil
func (o *InlineObject586) UnsetCurrentPassword() {
	o.CurrentPassword.Unset()
}

// GetNewPassword returns the NewPassword field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject586) GetNewPassword() string {
	if o == nil || o.NewPassword.Get() == nil {
		var ret string
		return ret
	}
	return *o.NewPassword.Get()
}

// GetNewPasswordOk returns a tuple with the NewPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject586) GetNewPasswordOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.NewPassword.Get(), o.NewPassword.IsSet()
}

// HasNewPassword returns a boolean if a field has been set.
func (o *InlineObject586) HasNewPassword() bool {
	if o != nil && o.NewPassword.IsSet() {
		return true
	}

	return false
}

// SetNewPassword gets a reference to the given NullableString and assigns it to the NewPassword field.
func (o *InlineObject586) SetNewPassword(v string) {
	o.NewPassword.Set(&v)
}
// SetNewPasswordNil sets the value for NewPassword to be an explicit nil
func (o *InlineObject586) SetNewPasswordNil() {
	o.NewPassword.Set(nil)
}

// UnsetNewPassword ensures that no value is present for NewPassword, not even an explicit nil
func (o *InlineObject586) UnsetNewPassword() {
	o.NewPassword.Unset()
}

func (o InlineObject586) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CurrentPassword.IsSet() {
		toSerialize["currentPassword"] = o.CurrentPassword.Get()
	}
	if o.NewPassword.IsSet() {
		toSerialize["newPassword"] = o.NewPassword.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject586 struct {
	value *InlineObject586
	isSet bool
}

func (v NullableInlineObject586) Get() *InlineObject586 {
	return v.value
}

func (v *NullableInlineObject586) Set(val *InlineObject586) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject586) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject586) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject586(val *InlineObject586) *NullableInlineObject586 {
	return &NullableInlineObject586{value: val, isSet: true}
}

func (v NullableInlineObject586) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject586) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


