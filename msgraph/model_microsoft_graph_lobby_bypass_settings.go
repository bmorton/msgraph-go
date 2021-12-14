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

// MicrosoftGraphLobbyBypassSettings struct for MicrosoftGraphLobbyBypassSettings
type MicrosoftGraphLobbyBypassSettings struct {
	// Specifies whether or not to always let dial-in callers bypass the lobby. Optional.
	IsDialInBypassEnabled NullableBool `json:"isDialInBypassEnabled,omitempty"`
	// Specifies the type of participants that are automatically admitted into a meeting, bypassing the lobby. Optional.
	Scope AnyOfmicrosoftGraphLobbyBypassScope `json:"scope,omitempty"`
}

// NewMicrosoftGraphLobbyBypassSettings instantiates a new MicrosoftGraphLobbyBypassSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphLobbyBypassSettings() *MicrosoftGraphLobbyBypassSettings {
	this := MicrosoftGraphLobbyBypassSettings{}
	return &this
}

// NewMicrosoftGraphLobbyBypassSettingsWithDefaults instantiates a new MicrosoftGraphLobbyBypassSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphLobbyBypassSettingsWithDefaults() *MicrosoftGraphLobbyBypassSettings {
	this := MicrosoftGraphLobbyBypassSettings{}
	return &this
}

// GetIsDialInBypassEnabled returns the IsDialInBypassEnabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphLobbyBypassSettings) GetIsDialInBypassEnabled() bool {
	if o == nil || o.IsDialInBypassEnabled.Get() == nil {
		var ret bool
		return ret
	}
	return *o.IsDialInBypassEnabled.Get()
}

// GetIsDialInBypassEnabledOk returns a tuple with the IsDialInBypassEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphLobbyBypassSettings) GetIsDialInBypassEnabledOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IsDialInBypassEnabled.Get(), o.IsDialInBypassEnabled.IsSet()
}

// HasIsDialInBypassEnabled returns a boolean if a field has been set.
func (o *MicrosoftGraphLobbyBypassSettings) HasIsDialInBypassEnabled() bool {
	if o != nil && o.IsDialInBypassEnabled.IsSet() {
		return true
	}

	return false
}

// SetIsDialInBypassEnabled gets a reference to the given NullableBool and assigns it to the IsDialInBypassEnabled field.
func (o *MicrosoftGraphLobbyBypassSettings) SetIsDialInBypassEnabled(v bool) {
	o.IsDialInBypassEnabled.Set(&v)
}
// SetIsDialInBypassEnabledNil sets the value for IsDialInBypassEnabled to be an explicit nil
func (o *MicrosoftGraphLobbyBypassSettings) SetIsDialInBypassEnabledNil() {
	o.IsDialInBypassEnabled.Set(nil)
}

// UnsetIsDialInBypassEnabled ensures that no value is present for IsDialInBypassEnabled, not even an explicit nil
func (o *MicrosoftGraphLobbyBypassSettings) UnsetIsDialInBypassEnabled() {
	o.IsDialInBypassEnabled.Unset()
}

// GetScope returns the Scope field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphLobbyBypassSettings) GetScope() AnyOfmicrosoftGraphLobbyBypassScope {
	if o == nil  {
		var ret AnyOfmicrosoftGraphLobbyBypassScope
		return ret
	}
	return o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphLobbyBypassSettings) GetScopeOk() (*AnyOfmicrosoftGraphLobbyBypassScope, bool) {
	if o == nil || o.Scope == nil {
		return nil, false
	}
	return &o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *MicrosoftGraphLobbyBypassSettings) HasScope() bool {
	if o != nil && o.Scope != nil {
		return true
	}

	return false
}

// SetScope gets a reference to the given AnyOfmicrosoftGraphLobbyBypassScope and assigns it to the Scope field.
func (o *MicrosoftGraphLobbyBypassSettings) SetScope(v AnyOfmicrosoftGraphLobbyBypassScope) {
	o.Scope = v
}

func (o MicrosoftGraphLobbyBypassSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.IsDialInBypassEnabled.IsSet() {
		toSerialize["isDialInBypassEnabled"] = o.IsDialInBypassEnabled.Get()
	}
	if o.Scope != nil {
		toSerialize["scope"] = o.Scope
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphLobbyBypassSettings struct {
	value *MicrosoftGraphLobbyBypassSettings
	isSet bool
}

func (v NullableMicrosoftGraphLobbyBypassSettings) Get() *MicrosoftGraphLobbyBypassSettings {
	return v.value
}

func (v *NullableMicrosoftGraphLobbyBypassSettings) Set(val *MicrosoftGraphLobbyBypassSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphLobbyBypassSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphLobbyBypassSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphLobbyBypassSettings(val *MicrosoftGraphLobbyBypassSettings) *NullableMicrosoftGraphLobbyBypassSettings {
	return &NullableMicrosoftGraphLobbyBypassSettings{value: val, isSet: true}
}

func (v NullableMicrosoftGraphLobbyBypassSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphLobbyBypassSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


