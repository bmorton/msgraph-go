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

// MicrosoftGraphTeamGuestSettings struct for MicrosoftGraphTeamGuestSettings
type MicrosoftGraphTeamGuestSettings struct {
	// If set to true, guests can add and update channels.
	AllowCreateUpdateChannels NullableBool `json:"allowCreateUpdateChannels,omitempty"`
	// If set to true, guests can delete channels.
	AllowDeleteChannels NullableBool `json:"allowDeleteChannels,omitempty"`
}

// NewMicrosoftGraphTeamGuestSettings instantiates a new MicrosoftGraphTeamGuestSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphTeamGuestSettings() *MicrosoftGraphTeamGuestSettings {
	this := MicrosoftGraphTeamGuestSettings{}
	return &this
}

// NewMicrosoftGraphTeamGuestSettingsWithDefaults instantiates a new MicrosoftGraphTeamGuestSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphTeamGuestSettingsWithDefaults() *MicrosoftGraphTeamGuestSettings {
	this := MicrosoftGraphTeamGuestSettings{}
	return &this
}

// GetAllowCreateUpdateChannels returns the AllowCreateUpdateChannels field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphTeamGuestSettings) GetAllowCreateUpdateChannels() bool {
	if o == nil || o.AllowCreateUpdateChannels.Get() == nil {
		var ret bool
		return ret
	}
	return *o.AllowCreateUpdateChannels.Get()
}

// GetAllowCreateUpdateChannelsOk returns a tuple with the AllowCreateUpdateChannels field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphTeamGuestSettings) GetAllowCreateUpdateChannelsOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AllowCreateUpdateChannels.Get(), o.AllowCreateUpdateChannels.IsSet()
}

// HasAllowCreateUpdateChannels returns a boolean if a field has been set.
func (o *MicrosoftGraphTeamGuestSettings) HasAllowCreateUpdateChannels() bool {
	if o != nil && o.AllowCreateUpdateChannels.IsSet() {
		return true
	}

	return false
}

// SetAllowCreateUpdateChannels gets a reference to the given NullableBool and assigns it to the AllowCreateUpdateChannels field.
func (o *MicrosoftGraphTeamGuestSettings) SetAllowCreateUpdateChannels(v bool) {
	o.AllowCreateUpdateChannels.Set(&v)
}
// SetAllowCreateUpdateChannelsNil sets the value for AllowCreateUpdateChannels to be an explicit nil
func (o *MicrosoftGraphTeamGuestSettings) SetAllowCreateUpdateChannelsNil() {
	o.AllowCreateUpdateChannels.Set(nil)
}

// UnsetAllowCreateUpdateChannels ensures that no value is present for AllowCreateUpdateChannels, not even an explicit nil
func (o *MicrosoftGraphTeamGuestSettings) UnsetAllowCreateUpdateChannels() {
	o.AllowCreateUpdateChannels.Unset()
}

// GetAllowDeleteChannels returns the AllowDeleteChannels field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphTeamGuestSettings) GetAllowDeleteChannels() bool {
	if o == nil || o.AllowDeleteChannels.Get() == nil {
		var ret bool
		return ret
	}
	return *o.AllowDeleteChannels.Get()
}

// GetAllowDeleteChannelsOk returns a tuple with the AllowDeleteChannels field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphTeamGuestSettings) GetAllowDeleteChannelsOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AllowDeleteChannels.Get(), o.AllowDeleteChannels.IsSet()
}

// HasAllowDeleteChannels returns a boolean if a field has been set.
func (o *MicrosoftGraphTeamGuestSettings) HasAllowDeleteChannels() bool {
	if o != nil && o.AllowDeleteChannels.IsSet() {
		return true
	}

	return false
}

// SetAllowDeleteChannels gets a reference to the given NullableBool and assigns it to the AllowDeleteChannels field.
func (o *MicrosoftGraphTeamGuestSettings) SetAllowDeleteChannels(v bool) {
	o.AllowDeleteChannels.Set(&v)
}
// SetAllowDeleteChannelsNil sets the value for AllowDeleteChannels to be an explicit nil
func (o *MicrosoftGraphTeamGuestSettings) SetAllowDeleteChannelsNil() {
	o.AllowDeleteChannels.Set(nil)
}

// UnsetAllowDeleteChannels ensures that no value is present for AllowDeleteChannels, not even an explicit nil
func (o *MicrosoftGraphTeamGuestSettings) UnsetAllowDeleteChannels() {
	o.AllowDeleteChannels.Unset()
}

func (o MicrosoftGraphTeamGuestSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AllowCreateUpdateChannels.IsSet() {
		toSerialize["allowCreateUpdateChannels"] = o.AllowCreateUpdateChannels.Get()
	}
	if o.AllowDeleteChannels.IsSet() {
		toSerialize["allowDeleteChannels"] = o.AllowDeleteChannels.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphTeamGuestSettings struct {
	value *MicrosoftGraphTeamGuestSettings
	isSet bool
}

func (v NullableMicrosoftGraphTeamGuestSettings) Get() *MicrosoftGraphTeamGuestSettings {
	return v.value
}

func (v *NullableMicrosoftGraphTeamGuestSettings) Set(val *MicrosoftGraphTeamGuestSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphTeamGuestSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphTeamGuestSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphTeamGuestSettings(val *MicrosoftGraphTeamGuestSettings) *NullableMicrosoftGraphTeamGuestSettings {
	return &NullableMicrosoftGraphTeamGuestSettings{value: val, isSet: true}
}

func (v NullableMicrosoftGraphTeamGuestSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphTeamGuestSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


