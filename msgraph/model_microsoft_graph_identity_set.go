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

// MicrosoftGraphIdentitySet struct for MicrosoftGraphIdentitySet
type MicrosoftGraphIdentitySet struct {
	// Optional. The application associated with this action.
	Application AnyOfmicrosoftGraphIdentity `json:"application,omitempty"`
	// Optional. The device associated with this action.
	Device AnyOfmicrosoftGraphIdentity `json:"device,omitempty"`
	// Optional. The user associated with this action.
	User AnyOfmicrosoftGraphIdentity `json:"user,omitempty"`
}

// NewMicrosoftGraphIdentitySet instantiates a new MicrosoftGraphIdentitySet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphIdentitySet() *MicrosoftGraphIdentitySet {
	this := MicrosoftGraphIdentitySet{}
	return &this
}

// NewMicrosoftGraphIdentitySetWithDefaults instantiates a new MicrosoftGraphIdentitySet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphIdentitySetWithDefaults() *MicrosoftGraphIdentitySet {
	this := MicrosoftGraphIdentitySet{}
	return &this
}

// GetApplication returns the Application field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphIdentitySet) GetApplication() AnyOfmicrosoftGraphIdentity {
	if o == nil  {
		var ret AnyOfmicrosoftGraphIdentity
		return ret
	}
	return o.Application
}

// GetApplicationOk returns a tuple with the Application field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphIdentitySet) GetApplicationOk() (*AnyOfmicrosoftGraphIdentity, bool) {
	if o == nil || o.Application == nil {
		return nil, false
	}
	return &o.Application, true
}

// HasApplication returns a boolean if a field has been set.
func (o *MicrosoftGraphIdentitySet) HasApplication() bool {
	if o != nil && o.Application != nil {
		return true
	}

	return false
}

// SetApplication gets a reference to the given AnyOfmicrosoftGraphIdentity and assigns it to the Application field.
func (o *MicrosoftGraphIdentitySet) SetApplication(v AnyOfmicrosoftGraphIdentity) {
	o.Application = v
}

// GetDevice returns the Device field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphIdentitySet) GetDevice() AnyOfmicrosoftGraphIdentity {
	if o == nil  {
		var ret AnyOfmicrosoftGraphIdentity
		return ret
	}
	return o.Device
}

// GetDeviceOk returns a tuple with the Device field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphIdentitySet) GetDeviceOk() (*AnyOfmicrosoftGraphIdentity, bool) {
	if o == nil || o.Device == nil {
		return nil, false
	}
	return &o.Device, true
}

// HasDevice returns a boolean if a field has been set.
func (o *MicrosoftGraphIdentitySet) HasDevice() bool {
	if o != nil && o.Device != nil {
		return true
	}

	return false
}

// SetDevice gets a reference to the given AnyOfmicrosoftGraphIdentity and assigns it to the Device field.
func (o *MicrosoftGraphIdentitySet) SetDevice(v AnyOfmicrosoftGraphIdentity) {
	o.Device = v
}

// GetUser returns the User field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphIdentitySet) GetUser() AnyOfmicrosoftGraphIdentity {
	if o == nil  {
		var ret AnyOfmicrosoftGraphIdentity
		return ret
	}
	return o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphIdentitySet) GetUserOk() (*AnyOfmicrosoftGraphIdentity, bool) {
	if o == nil || o.User == nil {
		return nil, false
	}
	return &o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *MicrosoftGraphIdentitySet) HasUser() bool {
	if o != nil && o.User != nil {
		return true
	}

	return false
}

// SetUser gets a reference to the given AnyOfmicrosoftGraphIdentity and assigns it to the User field.
func (o *MicrosoftGraphIdentitySet) SetUser(v AnyOfmicrosoftGraphIdentity) {
	o.User = v
}

func (o MicrosoftGraphIdentitySet) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Application != nil {
		toSerialize["application"] = o.Application
	}
	if o.Device != nil {
		toSerialize["device"] = o.Device
	}
	if o.User != nil {
		toSerialize["user"] = o.User
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphIdentitySet struct {
	value *MicrosoftGraphIdentitySet
	isSet bool
}

func (v NullableMicrosoftGraphIdentitySet) Get() *MicrosoftGraphIdentitySet {
	return v.value
}

func (v *NullableMicrosoftGraphIdentitySet) Set(val *MicrosoftGraphIdentitySet) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphIdentitySet) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphIdentitySet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphIdentitySet(val *MicrosoftGraphIdentitySet) *NullableMicrosoftGraphIdentitySet {
	return &NullableMicrosoftGraphIdentitySet{value: val, isSet: true}
}

func (v NullableMicrosoftGraphIdentitySet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphIdentitySet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

