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

// MicrosoftGraphSamlSingleSignOnSettings struct for MicrosoftGraphSamlSingleSignOnSettings
type MicrosoftGraphSamlSingleSignOnSettings struct {
	// The relative URI the service provider would redirect to after completion of the single sign-on flow.
	RelayState NullableString `json:"relayState,omitempty"`
}

// NewMicrosoftGraphSamlSingleSignOnSettings instantiates a new MicrosoftGraphSamlSingleSignOnSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphSamlSingleSignOnSettings() *MicrosoftGraphSamlSingleSignOnSettings {
	this := MicrosoftGraphSamlSingleSignOnSettings{}
	return &this
}

// NewMicrosoftGraphSamlSingleSignOnSettingsWithDefaults instantiates a new MicrosoftGraphSamlSingleSignOnSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphSamlSingleSignOnSettingsWithDefaults() *MicrosoftGraphSamlSingleSignOnSettings {
	this := MicrosoftGraphSamlSingleSignOnSettings{}
	return &this
}

// GetRelayState returns the RelayState field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphSamlSingleSignOnSettings) GetRelayState() string {
	if o == nil || o.RelayState.Get() == nil {
		var ret string
		return ret
	}
	return *o.RelayState.Get()
}

// GetRelayStateOk returns a tuple with the RelayState field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphSamlSingleSignOnSettings) GetRelayStateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RelayState.Get(), o.RelayState.IsSet()
}

// HasRelayState returns a boolean if a field has been set.
func (o *MicrosoftGraphSamlSingleSignOnSettings) HasRelayState() bool {
	if o != nil && o.RelayState.IsSet() {
		return true
	}

	return false
}

// SetRelayState gets a reference to the given NullableString and assigns it to the RelayState field.
func (o *MicrosoftGraphSamlSingleSignOnSettings) SetRelayState(v string) {
	o.RelayState.Set(&v)
}
// SetRelayStateNil sets the value for RelayState to be an explicit nil
func (o *MicrosoftGraphSamlSingleSignOnSettings) SetRelayStateNil() {
	o.RelayState.Set(nil)
}

// UnsetRelayState ensures that no value is present for RelayState, not even an explicit nil
func (o *MicrosoftGraphSamlSingleSignOnSettings) UnsetRelayState() {
	o.RelayState.Unset()
}

func (o MicrosoftGraphSamlSingleSignOnSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.RelayState.IsSet() {
		toSerialize["relayState"] = o.RelayState.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphSamlSingleSignOnSettings struct {
	value *MicrosoftGraphSamlSingleSignOnSettings
	isSet bool
}

func (v NullableMicrosoftGraphSamlSingleSignOnSettings) Get() *MicrosoftGraphSamlSingleSignOnSettings {
	return v.value
}

func (v *NullableMicrosoftGraphSamlSingleSignOnSettings) Set(val *MicrosoftGraphSamlSingleSignOnSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphSamlSingleSignOnSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphSamlSingleSignOnSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphSamlSingleSignOnSettings(val *MicrosoftGraphSamlSingleSignOnSettings) *NullableMicrosoftGraphSamlSingleSignOnSettings {
	return &NullableMicrosoftGraphSamlSingleSignOnSettings{value: val, isSet: true}
}

func (v NullableMicrosoftGraphSamlSingleSignOnSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphSamlSingleSignOnSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


