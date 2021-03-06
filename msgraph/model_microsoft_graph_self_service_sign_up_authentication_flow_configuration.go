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

// MicrosoftGraphSelfServiceSignUpAuthenticationFlowConfiguration struct for MicrosoftGraphSelfServiceSignUpAuthenticationFlowConfiguration
type MicrosoftGraphSelfServiceSignUpAuthenticationFlowConfiguration struct {
	// Indicates whether self-service sign-up flow is enabled or disabled. The default value is false. This property is not a key. Required.
	IsEnabled *bool `json:"isEnabled,omitempty"`
}

// NewMicrosoftGraphSelfServiceSignUpAuthenticationFlowConfiguration instantiates a new MicrosoftGraphSelfServiceSignUpAuthenticationFlowConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphSelfServiceSignUpAuthenticationFlowConfiguration() *MicrosoftGraphSelfServiceSignUpAuthenticationFlowConfiguration {
	this := MicrosoftGraphSelfServiceSignUpAuthenticationFlowConfiguration{}
	return &this
}

// NewMicrosoftGraphSelfServiceSignUpAuthenticationFlowConfigurationWithDefaults instantiates a new MicrosoftGraphSelfServiceSignUpAuthenticationFlowConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphSelfServiceSignUpAuthenticationFlowConfigurationWithDefaults() *MicrosoftGraphSelfServiceSignUpAuthenticationFlowConfiguration {
	this := MicrosoftGraphSelfServiceSignUpAuthenticationFlowConfiguration{}
	return &this
}

// GetIsEnabled returns the IsEnabled field value if set, zero value otherwise.
func (o *MicrosoftGraphSelfServiceSignUpAuthenticationFlowConfiguration) GetIsEnabled() bool {
	if o == nil || o.IsEnabled == nil {
		var ret bool
		return ret
	}
	return *o.IsEnabled
}

// GetIsEnabledOk returns a tuple with the IsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphSelfServiceSignUpAuthenticationFlowConfiguration) GetIsEnabledOk() (*bool, bool) {
	if o == nil || o.IsEnabled == nil {
		return nil, false
	}
	return o.IsEnabled, true
}

// HasIsEnabled returns a boolean if a field has been set.
func (o *MicrosoftGraphSelfServiceSignUpAuthenticationFlowConfiguration) HasIsEnabled() bool {
	if o != nil && o.IsEnabled != nil {
		return true
	}

	return false
}

// SetIsEnabled gets a reference to the given bool and assigns it to the IsEnabled field.
func (o *MicrosoftGraphSelfServiceSignUpAuthenticationFlowConfiguration) SetIsEnabled(v bool) {
	o.IsEnabled = &v
}

func (o MicrosoftGraphSelfServiceSignUpAuthenticationFlowConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.IsEnabled != nil {
		toSerialize["isEnabled"] = o.IsEnabled
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphSelfServiceSignUpAuthenticationFlowConfiguration struct {
	value *MicrosoftGraphSelfServiceSignUpAuthenticationFlowConfiguration
	isSet bool
}

func (v NullableMicrosoftGraphSelfServiceSignUpAuthenticationFlowConfiguration) Get() *MicrosoftGraphSelfServiceSignUpAuthenticationFlowConfiguration {
	return v.value
}

func (v *NullableMicrosoftGraphSelfServiceSignUpAuthenticationFlowConfiguration) Set(val *MicrosoftGraphSelfServiceSignUpAuthenticationFlowConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphSelfServiceSignUpAuthenticationFlowConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphSelfServiceSignUpAuthenticationFlowConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphSelfServiceSignUpAuthenticationFlowConfiguration(val *MicrosoftGraphSelfServiceSignUpAuthenticationFlowConfiguration) *NullableMicrosoftGraphSelfServiceSignUpAuthenticationFlowConfiguration {
	return &NullableMicrosoftGraphSelfServiceSignUpAuthenticationFlowConfiguration{value: val, isSet: true}
}

func (v NullableMicrosoftGraphSelfServiceSignUpAuthenticationFlowConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphSelfServiceSignUpAuthenticationFlowConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


