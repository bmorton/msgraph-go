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

// DeviceComplianceScheduledActionForRule Scheduled Action for Rule
type DeviceComplianceScheduledActionForRule struct {
	// Name of the rule which this scheduled action applies to. Currently scheduled actions are created per policy instead of per rule, thus RuleName is always set to default value PasswordRequired.
	RuleName NullableString `json:"ruleName,omitempty"`
	// The list of scheduled action configurations for this compliance policy. Compliance policy must have one and only one block scheduled action.
	ScheduledActionConfigurations *[]MicrosoftGraphDeviceComplianceActionItem `json:"scheduledActionConfigurations,omitempty"`
}

// NewDeviceComplianceScheduledActionForRule instantiates a new DeviceComplianceScheduledActionForRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceComplianceScheduledActionForRule() *DeviceComplianceScheduledActionForRule {
	this := DeviceComplianceScheduledActionForRule{}
	return &this
}

// NewDeviceComplianceScheduledActionForRuleWithDefaults instantiates a new DeviceComplianceScheduledActionForRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceComplianceScheduledActionForRuleWithDefaults() *DeviceComplianceScheduledActionForRule {
	this := DeviceComplianceScheduledActionForRule{}
	return &this
}

// GetRuleName returns the RuleName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeviceComplianceScheduledActionForRule) GetRuleName() string {
	if o == nil || o.RuleName.Get() == nil {
		var ret string
		return ret
	}
	return *o.RuleName.Get()
}

// GetRuleNameOk returns a tuple with the RuleName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeviceComplianceScheduledActionForRule) GetRuleNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RuleName.Get(), o.RuleName.IsSet()
}

// HasRuleName returns a boolean if a field has been set.
func (o *DeviceComplianceScheduledActionForRule) HasRuleName() bool {
	if o != nil && o.RuleName.IsSet() {
		return true
	}

	return false
}

// SetRuleName gets a reference to the given NullableString and assigns it to the RuleName field.
func (o *DeviceComplianceScheduledActionForRule) SetRuleName(v string) {
	o.RuleName.Set(&v)
}
// SetRuleNameNil sets the value for RuleName to be an explicit nil
func (o *DeviceComplianceScheduledActionForRule) SetRuleNameNil() {
	o.RuleName.Set(nil)
}

// UnsetRuleName ensures that no value is present for RuleName, not even an explicit nil
func (o *DeviceComplianceScheduledActionForRule) UnsetRuleName() {
	o.RuleName.Unset()
}

// GetScheduledActionConfigurations returns the ScheduledActionConfigurations field value if set, zero value otherwise.
func (o *DeviceComplianceScheduledActionForRule) GetScheduledActionConfigurations() []MicrosoftGraphDeviceComplianceActionItem {
	if o == nil || o.ScheduledActionConfigurations == nil {
		var ret []MicrosoftGraphDeviceComplianceActionItem
		return ret
	}
	return *o.ScheduledActionConfigurations
}

// GetScheduledActionConfigurationsOk returns a tuple with the ScheduledActionConfigurations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceComplianceScheduledActionForRule) GetScheduledActionConfigurationsOk() (*[]MicrosoftGraphDeviceComplianceActionItem, bool) {
	if o == nil || o.ScheduledActionConfigurations == nil {
		return nil, false
	}
	return o.ScheduledActionConfigurations, true
}

// HasScheduledActionConfigurations returns a boolean if a field has been set.
func (o *DeviceComplianceScheduledActionForRule) HasScheduledActionConfigurations() bool {
	if o != nil && o.ScheduledActionConfigurations != nil {
		return true
	}

	return false
}

// SetScheduledActionConfigurations gets a reference to the given []MicrosoftGraphDeviceComplianceActionItem and assigns it to the ScheduledActionConfigurations field.
func (o *DeviceComplianceScheduledActionForRule) SetScheduledActionConfigurations(v []MicrosoftGraphDeviceComplianceActionItem) {
	o.ScheduledActionConfigurations = &v
}

func (o DeviceComplianceScheduledActionForRule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.RuleName.IsSet() {
		toSerialize["ruleName"] = o.RuleName.Get()
	}
	if o.ScheduledActionConfigurations != nil {
		toSerialize["scheduledActionConfigurations"] = o.ScheduledActionConfigurations
	}
	return json.Marshal(toSerialize)
}

type NullableDeviceComplianceScheduledActionForRule struct {
	value *DeviceComplianceScheduledActionForRule
	isSet bool
}

func (v NullableDeviceComplianceScheduledActionForRule) Get() *DeviceComplianceScheduledActionForRule {
	return v.value
}

func (v *NullableDeviceComplianceScheduledActionForRule) Set(val *DeviceComplianceScheduledActionForRule) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceComplianceScheduledActionForRule) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceComplianceScheduledActionForRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceComplianceScheduledActionForRule(val *DeviceComplianceScheduledActionForRule) *NullableDeviceComplianceScheduledActionForRule {
	return &NullableDeviceComplianceScheduledActionForRule{value: val, isSet: true}
}

func (v NullableDeviceComplianceScheduledActionForRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceComplianceScheduledActionForRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


