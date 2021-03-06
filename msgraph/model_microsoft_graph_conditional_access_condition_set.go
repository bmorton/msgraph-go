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

// MicrosoftGraphConditionalAccessConditionSet struct for MicrosoftGraphConditionalAccessConditionSet
type MicrosoftGraphConditionalAccessConditionSet struct {
	// Applications and user actions included in and excluded from the policy. Required.
	Applications AnyOfmicrosoftGraphConditionalAccessApplications `json:"applications,omitempty"`
	// Client application types included in the policy. Possible values are: all, browser, mobileAppsAndDesktopClients, exchangeActiveSync, easSupported, other. Required.
	ClientAppTypes *[]AnyOfmicrosoftGraphConditionalAccessClientApp `json:"clientAppTypes,omitempty"`
	// Devices in the policy.
	Devices AnyOfmicrosoftGraphConditionalAccessDevices `json:"devices,omitempty"`
	// Locations included in and excluded from the policy.
	Locations AnyOfmicrosoftGraphConditionalAccessLocations `json:"locations,omitempty"`
	// Platforms included in and excluded from the policy.
	Platforms AnyOfmicrosoftGraphConditionalAccessPlatforms `json:"platforms,omitempty"`
	// Sign-in risk levels included in the policy. Possible values are: low, medium, high, hidden, none, unknownFutureValue. Required.
	SignInRiskLevels *[]AnyOfmicrosoftGraphRiskLevel `json:"signInRiskLevels,omitempty"`
	// User risk levels included in the policy. Possible values are: low, medium, high, hidden, none, unknownFutureValue. Required.
	UserRiskLevels *[]AnyOfmicrosoftGraphRiskLevel `json:"userRiskLevels,omitempty"`
	Users *MicrosoftGraphConditionalAccessUsers `json:"users,omitempty"`
}

// NewMicrosoftGraphConditionalAccessConditionSet instantiates a new MicrosoftGraphConditionalAccessConditionSet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphConditionalAccessConditionSet() *MicrosoftGraphConditionalAccessConditionSet {
	this := MicrosoftGraphConditionalAccessConditionSet{}
	return &this
}

// NewMicrosoftGraphConditionalAccessConditionSetWithDefaults instantiates a new MicrosoftGraphConditionalAccessConditionSet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphConditionalAccessConditionSetWithDefaults() *MicrosoftGraphConditionalAccessConditionSet {
	this := MicrosoftGraphConditionalAccessConditionSet{}
	return &this
}

// GetApplications returns the Applications field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphConditionalAccessConditionSet) GetApplications() AnyOfmicrosoftGraphConditionalAccessApplications {
	if o == nil  {
		var ret AnyOfmicrosoftGraphConditionalAccessApplications
		return ret
	}
	return o.Applications
}

// GetApplicationsOk returns a tuple with the Applications field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphConditionalAccessConditionSet) GetApplicationsOk() (*AnyOfmicrosoftGraphConditionalAccessApplications, bool) {
	if o == nil || o.Applications == nil {
		return nil, false
	}
	return &o.Applications, true
}

// HasApplications returns a boolean if a field has been set.
func (o *MicrosoftGraphConditionalAccessConditionSet) HasApplications() bool {
	if o != nil && o.Applications != nil {
		return true
	}

	return false
}

// SetApplications gets a reference to the given AnyOfmicrosoftGraphConditionalAccessApplications and assigns it to the Applications field.
func (o *MicrosoftGraphConditionalAccessConditionSet) SetApplications(v AnyOfmicrosoftGraphConditionalAccessApplications) {
	o.Applications = v
}

// GetClientAppTypes returns the ClientAppTypes field value if set, zero value otherwise.
func (o *MicrosoftGraphConditionalAccessConditionSet) GetClientAppTypes() []AnyOfmicrosoftGraphConditionalAccessClientApp {
	if o == nil || o.ClientAppTypes == nil {
		var ret []AnyOfmicrosoftGraphConditionalAccessClientApp
		return ret
	}
	return *o.ClientAppTypes
}

// GetClientAppTypesOk returns a tuple with the ClientAppTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphConditionalAccessConditionSet) GetClientAppTypesOk() (*[]AnyOfmicrosoftGraphConditionalAccessClientApp, bool) {
	if o == nil || o.ClientAppTypes == nil {
		return nil, false
	}
	return o.ClientAppTypes, true
}

// HasClientAppTypes returns a boolean if a field has been set.
func (o *MicrosoftGraphConditionalAccessConditionSet) HasClientAppTypes() bool {
	if o != nil && o.ClientAppTypes != nil {
		return true
	}

	return false
}

// SetClientAppTypes gets a reference to the given []AnyOfmicrosoftGraphConditionalAccessClientApp and assigns it to the ClientAppTypes field.
func (o *MicrosoftGraphConditionalAccessConditionSet) SetClientAppTypes(v []AnyOfmicrosoftGraphConditionalAccessClientApp) {
	o.ClientAppTypes = &v
}

// GetDevices returns the Devices field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphConditionalAccessConditionSet) GetDevices() AnyOfmicrosoftGraphConditionalAccessDevices {
	if o == nil  {
		var ret AnyOfmicrosoftGraphConditionalAccessDevices
		return ret
	}
	return o.Devices
}

// GetDevicesOk returns a tuple with the Devices field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphConditionalAccessConditionSet) GetDevicesOk() (*AnyOfmicrosoftGraphConditionalAccessDevices, bool) {
	if o == nil || o.Devices == nil {
		return nil, false
	}
	return &o.Devices, true
}

// HasDevices returns a boolean if a field has been set.
func (o *MicrosoftGraphConditionalAccessConditionSet) HasDevices() bool {
	if o != nil && o.Devices != nil {
		return true
	}

	return false
}

// SetDevices gets a reference to the given AnyOfmicrosoftGraphConditionalAccessDevices and assigns it to the Devices field.
func (o *MicrosoftGraphConditionalAccessConditionSet) SetDevices(v AnyOfmicrosoftGraphConditionalAccessDevices) {
	o.Devices = v
}

// GetLocations returns the Locations field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphConditionalAccessConditionSet) GetLocations() AnyOfmicrosoftGraphConditionalAccessLocations {
	if o == nil  {
		var ret AnyOfmicrosoftGraphConditionalAccessLocations
		return ret
	}
	return o.Locations
}

// GetLocationsOk returns a tuple with the Locations field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphConditionalAccessConditionSet) GetLocationsOk() (*AnyOfmicrosoftGraphConditionalAccessLocations, bool) {
	if o == nil || o.Locations == nil {
		return nil, false
	}
	return &o.Locations, true
}

// HasLocations returns a boolean if a field has been set.
func (o *MicrosoftGraphConditionalAccessConditionSet) HasLocations() bool {
	if o != nil && o.Locations != nil {
		return true
	}

	return false
}

// SetLocations gets a reference to the given AnyOfmicrosoftGraphConditionalAccessLocations and assigns it to the Locations field.
func (o *MicrosoftGraphConditionalAccessConditionSet) SetLocations(v AnyOfmicrosoftGraphConditionalAccessLocations) {
	o.Locations = v
}

// GetPlatforms returns the Platforms field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphConditionalAccessConditionSet) GetPlatforms() AnyOfmicrosoftGraphConditionalAccessPlatforms {
	if o == nil  {
		var ret AnyOfmicrosoftGraphConditionalAccessPlatforms
		return ret
	}
	return o.Platforms
}

// GetPlatformsOk returns a tuple with the Platforms field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphConditionalAccessConditionSet) GetPlatformsOk() (*AnyOfmicrosoftGraphConditionalAccessPlatforms, bool) {
	if o == nil || o.Platforms == nil {
		return nil, false
	}
	return &o.Platforms, true
}

// HasPlatforms returns a boolean if a field has been set.
func (o *MicrosoftGraphConditionalAccessConditionSet) HasPlatforms() bool {
	if o != nil && o.Platforms != nil {
		return true
	}

	return false
}

// SetPlatforms gets a reference to the given AnyOfmicrosoftGraphConditionalAccessPlatforms and assigns it to the Platforms field.
func (o *MicrosoftGraphConditionalAccessConditionSet) SetPlatforms(v AnyOfmicrosoftGraphConditionalAccessPlatforms) {
	o.Platforms = v
}

// GetSignInRiskLevels returns the SignInRiskLevels field value if set, zero value otherwise.
func (o *MicrosoftGraphConditionalAccessConditionSet) GetSignInRiskLevels() []AnyOfmicrosoftGraphRiskLevel {
	if o == nil || o.SignInRiskLevels == nil {
		var ret []AnyOfmicrosoftGraphRiskLevel
		return ret
	}
	return *o.SignInRiskLevels
}

// GetSignInRiskLevelsOk returns a tuple with the SignInRiskLevels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphConditionalAccessConditionSet) GetSignInRiskLevelsOk() (*[]AnyOfmicrosoftGraphRiskLevel, bool) {
	if o == nil || o.SignInRiskLevels == nil {
		return nil, false
	}
	return o.SignInRiskLevels, true
}

// HasSignInRiskLevels returns a boolean if a field has been set.
func (o *MicrosoftGraphConditionalAccessConditionSet) HasSignInRiskLevels() bool {
	if o != nil && o.SignInRiskLevels != nil {
		return true
	}

	return false
}

// SetSignInRiskLevels gets a reference to the given []AnyOfmicrosoftGraphRiskLevel and assigns it to the SignInRiskLevels field.
func (o *MicrosoftGraphConditionalAccessConditionSet) SetSignInRiskLevels(v []AnyOfmicrosoftGraphRiskLevel) {
	o.SignInRiskLevels = &v
}

// GetUserRiskLevels returns the UserRiskLevels field value if set, zero value otherwise.
func (o *MicrosoftGraphConditionalAccessConditionSet) GetUserRiskLevels() []AnyOfmicrosoftGraphRiskLevel {
	if o == nil || o.UserRiskLevels == nil {
		var ret []AnyOfmicrosoftGraphRiskLevel
		return ret
	}
	return *o.UserRiskLevels
}

// GetUserRiskLevelsOk returns a tuple with the UserRiskLevels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphConditionalAccessConditionSet) GetUserRiskLevelsOk() (*[]AnyOfmicrosoftGraphRiskLevel, bool) {
	if o == nil || o.UserRiskLevels == nil {
		return nil, false
	}
	return o.UserRiskLevels, true
}

// HasUserRiskLevels returns a boolean if a field has been set.
func (o *MicrosoftGraphConditionalAccessConditionSet) HasUserRiskLevels() bool {
	if o != nil && o.UserRiskLevels != nil {
		return true
	}

	return false
}

// SetUserRiskLevels gets a reference to the given []AnyOfmicrosoftGraphRiskLevel and assigns it to the UserRiskLevels field.
func (o *MicrosoftGraphConditionalAccessConditionSet) SetUserRiskLevels(v []AnyOfmicrosoftGraphRiskLevel) {
	o.UserRiskLevels = &v
}

// GetUsers returns the Users field value if set, zero value otherwise.
func (o *MicrosoftGraphConditionalAccessConditionSet) GetUsers() MicrosoftGraphConditionalAccessUsers {
	if o == nil || o.Users == nil {
		var ret MicrosoftGraphConditionalAccessUsers
		return ret
	}
	return *o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphConditionalAccessConditionSet) GetUsersOk() (*MicrosoftGraphConditionalAccessUsers, bool) {
	if o == nil || o.Users == nil {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *MicrosoftGraphConditionalAccessConditionSet) HasUsers() bool {
	if o != nil && o.Users != nil {
		return true
	}

	return false
}

// SetUsers gets a reference to the given MicrosoftGraphConditionalAccessUsers and assigns it to the Users field.
func (o *MicrosoftGraphConditionalAccessConditionSet) SetUsers(v MicrosoftGraphConditionalAccessUsers) {
	o.Users = &v
}

func (o MicrosoftGraphConditionalAccessConditionSet) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Applications != nil {
		toSerialize["applications"] = o.Applications
	}
	if o.ClientAppTypes != nil {
		toSerialize["clientAppTypes"] = o.ClientAppTypes
	}
	if o.Devices != nil {
		toSerialize["devices"] = o.Devices
	}
	if o.Locations != nil {
		toSerialize["locations"] = o.Locations
	}
	if o.Platforms != nil {
		toSerialize["platforms"] = o.Platforms
	}
	if o.SignInRiskLevels != nil {
		toSerialize["signInRiskLevels"] = o.SignInRiskLevels
	}
	if o.UserRiskLevels != nil {
		toSerialize["userRiskLevels"] = o.UserRiskLevels
	}
	if o.Users != nil {
		toSerialize["users"] = o.Users
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphConditionalAccessConditionSet struct {
	value *MicrosoftGraphConditionalAccessConditionSet
	isSet bool
}

func (v NullableMicrosoftGraphConditionalAccessConditionSet) Get() *MicrosoftGraphConditionalAccessConditionSet {
	return v.value
}

func (v *NullableMicrosoftGraphConditionalAccessConditionSet) Set(val *MicrosoftGraphConditionalAccessConditionSet) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphConditionalAccessConditionSet) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphConditionalAccessConditionSet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphConditionalAccessConditionSet(val *MicrosoftGraphConditionalAccessConditionSet) *NullableMicrosoftGraphConditionalAccessConditionSet {
	return &NullableMicrosoftGraphConditionalAccessConditionSet{value: val, isSet: true}
}

func (v NullableMicrosoftGraphConditionalAccessConditionSet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphConditionalAccessConditionSet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


