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

// UserSettings struct for UserSettings
type UserSettings struct {
	// Reflects the Office Delve organization level setting. When set to true, the organization doesn't have access to Office Delve. This setting is read-only and can only be changed by administrators in the SharePoint admin center.
	ContributionToContentDiscoveryAsOrganizationDisabled *bool `json:"contributionToContentDiscoveryAsOrganizationDisabled,omitempty"`
	// When set to true, documents in the user's Office Delve are disabled. Users can control this setting in Office Delve.
	ContributionToContentDiscoveryDisabled *bool `json:"contributionToContentDiscoveryDisabled,omitempty"`
	// The shift preferences for the user.
	ShiftPreferences AnyOfmicrosoftGraphShiftPreferences `json:"shiftPreferences,omitempty"`
}

// NewUserSettings instantiates a new UserSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserSettings() *UserSettings {
	this := UserSettings{}
	return &this
}

// NewUserSettingsWithDefaults instantiates a new UserSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserSettingsWithDefaults() *UserSettings {
	this := UserSettings{}
	return &this
}

// GetContributionToContentDiscoveryAsOrganizationDisabled returns the ContributionToContentDiscoveryAsOrganizationDisabled field value if set, zero value otherwise.
func (o *UserSettings) GetContributionToContentDiscoveryAsOrganizationDisabled() bool {
	if o == nil || o.ContributionToContentDiscoveryAsOrganizationDisabled == nil {
		var ret bool
		return ret
	}
	return *o.ContributionToContentDiscoveryAsOrganizationDisabled
}

// GetContributionToContentDiscoveryAsOrganizationDisabledOk returns a tuple with the ContributionToContentDiscoveryAsOrganizationDisabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSettings) GetContributionToContentDiscoveryAsOrganizationDisabledOk() (*bool, bool) {
	if o == nil || o.ContributionToContentDiscoveryAsOrganizationDisabled == nil {
		return nil, false
	}
	return o.ContributionToContentDiscoveryAsOrganizationDisabled, true
}

// HasContributionToContentDiscoveryAsOrganizationDisabled returns a boolean if a field has been set.
func (o *UserSettings) HasContributionToContentDiscoveryAsOrganizationDisabled() bool {
	if o != nil && o.ContributionToContentDiscoveryAsOrganizationDisabled != nil {
		return true
	}

	return false
}

// SetContributionToContentDiscoveryAsOrganizationDisabled gets a reference to the given bool and assigns it to the ContributionToContentDiscoveryAsOrganizationDisabled field.
func (o *UserSettings) SetContributionToContentDiscoveryAsOrganizationDisabled(v bool) {
	o.ContributionToContentDiscoveryAsOrganizationDisabled = &v
}

// GetContributionToContentDiscoveryDisabled returns the ContributionToContentDiscoveryDisabled field value if set, zero value otherwise.
func (o *UserSettings) GetContributionToContentDiscoveryDisabled() bool {
	if o == nil || o.ContributionToContentDiscoveryDisabled == nil {
		var ret bool
		return ret
	}
	return *o.ContributionToContentDiscoveryDisabled
}

// GetContributionToContentDiscoveryDisabledOk returns a tuple with the ContributionToContentDiscoveryDisabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSettings) GetContributionToContentDiscoveryDisabledOk() (*bool, bool) {
	if o == nil || o.ContributionToContentDiscoveryDisabled == nil {
		return nil, false
	}
	return o.ContributionToContentDiscoveryDisabled, true
}

// HasContributionToContentDiscoveryDisabled returns a boolean if a field has been set.
func (o *UserSettings) HasContributionToContentDiscoveryDisabled() bool {
	if o != nil && o.ContributionToContentDiscoveryDisabled != nil {
		return true
	}

	return false
}

// SetContributionToContentDiscoveryDisabled gets a reference to the given bool and assigns it to the ContributionToContentDiscoveryDisabled field.
func (o *UserSettings) SetContributionToContentDiscoveryDisabled(v bool) {
	o.ContributionToContentDiscoveryDisabled = &v
}

// GetShiftPreferences returns the ShiftPreferences field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserSettings) GetShiftPreferences() AnyOfmicrosoftGraphShiftPreferences {
	if o == nil  {
		var ret AnyOfmicrosoftGraphShiftPreferences
		return ret
	}
	return o.ShiftPreferences
}

// GetShiftPreferencesOk returns a tuple with the ShiftPreferences field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserSettings) GetShiftPreferencesOk() (*AnyOfmicrosoftGraphShiftPreferences, bool) {
	if o == nil || o.ShiftPreferences == nil {
		return nil, false
	}
	return &o.ShiftPreferences, true
}

// HasShiftPreferences returns a boolean if a field has been set.
func (o *UserSettings) HasShiftPreferences() bool {
	if o != nil && o.ShiftPreferences != nil {
		return true
	}

	return false
}

// SetShiftPreferences gets a reference to the given AnyOfmicrosoftGraphShiftPreferences and assigns it to the ShiftPreferences field.
func (o *UserSettings) SetShiftPreferences(v AnyOfmicrosoftGraphShiftPreferences) {
	o.ShiftPreferences = v
}

func (o UserSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ContributionToContentDiscoveryAsOrganizationDisabled != nil {
		toSerialize["contributionToContentDiscoveryAsOrganizationDisabled"] = o.ContributionToContentDiscoveryAsOrganizationDisabled
	}
	if o.ContributionToContentDiscoveryDisabled != nil {
		toSerialize["contributionToContentDiscoveryDisabled"] = o.ContributionToContentDiscoveryDisabled
	}
	if o.ShiftPreferences != nil {
		toSerialize["shiftPreferences"] = o.ShiftPreferences
	}
	return json.Marshal(toSerialize)
}

type NullableUserSettings struct {
	value *UserSettings
	isSet bool
}

func (v NullableUserSettings) Get() *UserSettings {
	return v.value
}

func (v *NullableUserSettings) Set(val *UserSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableUserSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableUserSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserSettings(val *UserSettings) *NullableUserSettings {
	return &NullableUserSettings{value: val, isSet: true}
}

func (v NullableUserSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


