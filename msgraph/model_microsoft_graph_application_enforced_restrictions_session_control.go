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

// MicrosoftGraphApplicationEnforcedRestrictionsSessionControl struct for MicrosoftGraphApplicationEnforcedRestrictionsSessionControl
type MicrosoftGraphApplicationEnforcedRestrictionsSessionControl struct {
	// Specifies whether the session control is enabled.
	IsEnabled NullableBool `json:"isEnabled,omitempty"`
}

// NewMicrosoftGraphApplicationEnforcedRestrictionsSessionControl instantiates a new MicrosoftGraphApplicationEnforcedRestrictionsSessionControl object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphApplicationEnforcedRestrictionsSessionControl() *MicrosoftGraphApplicationEnforcedRestrictionsSessionControl {
	this := MicrosoftGraphApplicationEnforcedRestrictionsSessionControl{}
	return &this
}

// NewMicrosoftGraphApplicationEnforcedRestrictionsSessionControlWithDefaults instantiates a new MicrosoftGraphApplicationEnforcedRestrictionsSessionControl object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphApplicationEnforcedRestrictionsSessionControlWithDefaults() *MicrosoftGraphApplicationEnforcedRestrictionsSessionControl {
	this := MicrosoftGraphApplicationEnforcedRestrictionsSessionControl{}
	return &this
}

// GetIsEnabled returns the IsEnabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphApplicationEnforcedRestrictionsSessionControl) GetIsEnabled() bool {
	if o == nil || o.IsEnabled.Get() == nil {
		var ret bool
		return ret
	}
	return *o.IsEnabled.Get()
}

// GetIsEnabledOk returns a tuple with the IsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphApplicationEnforcedRestrictionsSessionControl) GetIsEnabledOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IsEnabled.Get(), o.IsEnabled.IsSet()
}

// HasIsEnabled returns a boolean if a field has been set.
func (o *MicrosoftGraphApplicationEnforcedRestrictionsSessionControl) HasIsEnabled() bool {
	if o != nil && o.IsEnabled.IsSet() {
		return true
	}

	return false
}

// SetIsEnabled gets a reference to the given NullableBool and assigns it to the IsEnabled field.
func (o *MicrosoftGraphApplicationEnforcedRestrictionsSessionControl) SetIsEnabled(v bool) {
	o.IsEnabled.Set(&v)
}
// SetIsEnabledNil sets the value for IsEnabled to be an explicit nil
func (o *MicrosoftGraphApplicationEnforcedRestrictionsSessionControl) SetIsEnabledNil() {
	o.IsEnabled.Set(nil)
}

// UnsetIsEnabled ensures that no value is present for IsEnabled, not even an explicit nil
func (o *MicrosoftGraphApplicationEnforcedRestrictionsSessionControl) UnsetIsEnabled() {
	o.IsEnabled.Unset()
}

func (o MicrosoftGraphApplicationEnforcedRestrictionsSessionControl) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.IsEnabled.IsSet() {
		toSerialize["isEnabled"] = o.IsEnabled.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphApplicationEnforcedRestrictionsSessionControl struct {
	value *MicrosoftGraphApplicationEnforcedRestrictionsSessionControl
	isSet bool
}

func (v NullableMicrosoftGraphApplicationEnforcedRestrictionsSessionControl) Get() *MicrosoftGraphApplicationEnforcedRestrictionsSessionControl {
	return v.value
}

func (v *NullableMicrosoftGraphApplicationEnforcedRestrictionsSessionControl) Set(val *MicrosoftGraphApplicationEnforcedRestrictionsSessionControl) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphApplicationEnforcedRestrictionsSessionControl) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphApplicationEnforcedRestrictionsSessionControl) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphApplicationEnforcedRestrictionsSessionControl(val *MicrosoftGraphApplicationEnforcedRestrictionsSessionControl) *NullableMicrosoftGraphApplicationEnforcedRestrictionsSessionControl {
	return &NullableMicrosoftGraphApplicationEnforcedRestrictionsSessionControl{value: val, isSet: true}
}

func (v NullableMicrosoftGraphApplicationEnforcedRestrictionsSessionControl) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphApplicationEnforcedRestrictionsSessionControl) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


