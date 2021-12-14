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

// OnPremisesConditionalAccessSettings Singleton entity which represents the Exchange OnPremises Conditional Access Settings for a tenant.
type OnPremisesConditionalAccessSettings struct {
	// Indicates if on premises conditional access is enabled for this organization
	Enabled *bool `json:"enabled,omitempty"`
	// User groups that will be exempt by on premises conditional access. All users in these groups will be exempt from the conditional access policy.
	ExcludedGroups *[]string `json:"excludedGroups,omitempty"`
	// User groups that will be targeted by on premises conditional access. All users in these groups will be required to have mobile device managed and compliant for mail access.
	IncludedGroups *[]string `json:"includedGroups,omitempty"`
	// Override the default access rule when allowing a device to ensure access is granted.
	OverrideDefaultRule *bool `json:"overrideDefaultRule,omitempty"`
}

// NewOnPremisesConditionalAccessSettings instantiates a new OnPremisesConditionalAccessSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOnPremisesConditionalAccessSettings() *OnPremisesConditionalAccessSettings {
	this := OnPremisesConditionalAccessSettings{}
	return &this
}

// NewOnPremisesConditionalAccessSettingsWithDefaults instantiates a new OnPremisesConditionalAccessSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOnPremisesConditionalAccessSettingsWithDefaults() *OnPremisesConditionalAccessSettings {
	this := OnPremisesConditionalAccessSettings{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *OnPremisesConditionalAccessSettings) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnPremisesConditionalAccessSettings) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *OnPremisesConditionalAccessSettings) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *OnPremisesConditionalAccessSettings) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetExcludedGroups returns the ExcludedGroups field value if set, zero value otherwise.
func (o *OnPremisesConditionalAccessSettings) GetExcludedGroups() []string {
	if o == nil || o.ExcludedGroups == nil {
		var ret []string
		return ret
	}
	return *o.ExcludedGroups
}

// GetExcludedGroupsOk returns a tuple with the ExcludedGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnPremisesConditionalAccessSettings) GetExcludedGroupsOk() (*[]string, bool) {
	if o == nil || o.ExcludedGroups == nil {
		return nil, false
	}
	return o.ExcludedGroups, true
}

// HasExcludedGroups returns a boolean if a field has been set.
func (o *OnPremisesConditionalAccessSettings) HasExcludedGroups() bool {
	if o != nil && o.ExcludedGroups != nil {
		return true
	}

	return false
}

// SetExcludedGroups gets a reference to the given []string and assigns it to the ExcludedGroups field.
func (o *OnPremisesConditionalAccessSettings) SetExcludedGroups(v []string) {
	o.ExcludedGroups = &v
}

// GetIncludedGroups returns the IncludedGroups field value if set, zero value otherwise.
func (o *OnPremisesConditionalAccessSettings) GetIncludedGroups() []string {
	if o == nil || o.IncludedGroups == nil {
		var ret []string
		return ret
	}
	return *o.IncludedGroups
}

// GetIncludedGroupsOk returns a tuple with the IncludedGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnPremisesConditionalAccessSettings) GetIncludedGroupsOk() (*[]string, bool) {
	if o == nil || o.IncludedGroups == nil {
		return nil, false
	}
	return o.IncludedGroups, true
}

// HasIncludedGroups returns a boolean if a field has been set.
func (o *OnPremisesConditionalAccessSettings) HasIncludedGroups() bool {
	if o != nil && o.IncludedGroups != nil {
		return true
	}

	return false
}

// SetIncludedGroups gets a reference to the given []string and assigns it to the IncludedGroups field.
func (o *OnPremisesConditionalAccessSettings) SetIncludedGroups(v []string) {
	o.IncludedGroups = &v
}

// GetOverrideDefaultRule returns the OverrideDefaultRule field value if set, zero value otherwise.
func (o *OnPremisesConditionalAccessSettings) GetOverrideDefaultRule() bool {
	if o == nil || o.OverrideDefaultRule == nil {
		var ret bool
		return ret
	}
	return *o.OverrideDefaultRule
}

// GetOverrideDefaultRuleOk returns a tuple with the OverrideDefaultRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnPremisesConditionalAccessSettings) GetOverrideDefaultRuleOk() (*bool, bool) {
	if o == nil || o.OverrideDefaultRule == nil {
		return nil, false
	}
	return o.OverrideDefaultRule, true
}

// HasOverrideDefaultRule returns a boolean if a field has been set.
func (o *OnPremisesConditionalAccessSettings) HasOverrideDefaultRule() bool {
	if o != nil && o.OverrideDefaultRule != nil {
		return true
	}

	return false
}

// SetOverrideDefaultRule gets a reference to the given bool and assigns it to the OverrideDefaultRule field.
func (o *OnPremisesConditionalAccessSettings) SetOverrideDefaultRule(v bool) {
	o.OverrideDefaultRule = &v
}

func (o OnPremisesConditionalAccessSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.ExcludedGroups != nil {
		toSerialize["excludedGroups"] = o.ExcludedGroups
	}
	if o.IncludedGroups != nil {
		toSerialize["includedGroups"] = o.IncludedGroups
	}
	if o.OverrideDefaultRule != nil {
		toSerialize["overrideDefaultRule"] = o.OverrideDefaultRule
	}
	return json.Marshal(toSerialize)
}

type NullableOnPremisesConditionalAccessSettings struct {
	value *OnPremisesConditionalAccessSettings
	isSet bool
}

func (v NullableOnPremisesConditionalAccessSettings) Get() *OnPremisesConditionalAccessSettings {
	return v.value
}

func (v *NullableOnPremisesConditionalAccessSettings) Set(val *OnPremisesConditionalAccessSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableOnPremisesConditionalAccessSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableOnPremisesConditionalAccessSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOnPremisesConditionalAccessSettings(val *OnPremisesConditionalAccessSettings) *NullableOnPremisesConditionalAccessSettings {
	return &NullableOnPremisesConditionalAccessSettings{value: val, isSet: true}
}

func (v NullableOnPremisesConditionalAccessSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOnPremisesConditionalAccessSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


