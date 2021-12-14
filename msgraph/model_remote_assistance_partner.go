/*
Partial Graph API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// RemoteAssistancePartner RemoteAssistPartner resources represent the metadata and status of a given Remote Assistance partner service.
type RemoteAssistancePartner struct {
	// Display name of the partner.
	DisplayName NullableString `json:"displayName,omitempty"`
	// Timestamp of the last request sent to Intune by the TEM partner.
	LastConnectionDateTime *time.Time `json:"lastConnectionDateTime,omitempty"`
	// A friendly description of the current TeamViewer connector status. Possible values are: notOnboarded, onboarding, onboarded.
	OnboardingStatus AnyOfmicrosoftGraphRemoteAssistanceOnboardingStatus `json:"onboardingStatus,omitempty"`
	// URL of the partner's onboarding portal, where an administrator can configure their Remote Assistance service.
	OnboardingUrl NullableString `json:"onboardingUrl,omitempty"`
}

// NewRemoteAssistancePartner instantiates a new RemoteAssistancePartner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRemoteAssistancePartner() *RemoteAssistancePartner {
	this := RemoteAssistancePartner{}
	return &this
}

// NewRemoteAssistancePartnerWithDefaults instantiates a new RemoteAssistancePartner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRemoteAssistancePartnerWithDefaults() *RemoteAssistancePartner {
	this := RemoteAssistancePartner{}
	return &this
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RemoteAssistancePartner) GetDisplayName() string {
	if o == nil || o.DisplayName.Get() == nil {
		var ret string
		return ret
	}
	return *o.DisplayName.Get()
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RemoteAssistancePartner) GetDisplayNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DisplayName.Get(), o.DisplayName.IsSet()
}

// HasDisplayName returns a boolean if a field has been set.
func (o *RemoteAssistancePartner) HasDisplayName() bool {
	if o != nil && o.DisplayName.IsSet() {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given NullableString and assigns it to the DisplayName field.
func (o *RemoteAssistancePartner) SetDisplayName(v string) {
	o.DisplayName.Set(&v)
}
// SetDisplayNameNil sets the value for DisplayName to be an explicit nil
func (o *RemoteAssistancePartner) SetDisplayNameNil() {
	o.DisplayName.Set(nil)
}

// UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
func (o *RemoteAssistancePartner) UnsetDisplayName() {
	o.DisplayName.Unset()
}

// GetLastConnectionDateTime returns the LastConnectionDateTime field value if set, zero value otherwise.
func (o *RemoteAssistancePartner) GetLastConnectionDateTime() time.Time {
	if o == nil || o.LastConnectionDateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.LastConnectionDateTime
}

// GetLastConnectionDateTimeOk returns a tuple with the LastConnectionDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteAssistancePartner) GetLastConnectionDateTimeOk() (*time.Time, bool) {
	if o == nil || o.LastConnectionDateTime == nil {
		return nil, false
	}
	return o.LastConnectionDateTime, true
}

// HasLastConnectionDateTime returns a boolean if a field has been set.
func (o *RemoteAssistancePartner) HasLastConnectionDateTime() bool {
	if o != nil && o.LastConnectionDateTime != nil {
		return true
	}

	return false
}

// SetLastConnectionDateTime gets a reference to the given time.Time and assigns it to the LastConnectionDateTime field.
func (o *RemoteAssistancePartner) SetLastConnectionDateTime(v time.Time) {
	o.LastConnectionDateTime = &v
}

// GetOnboardingStatus returns the OnboardingStatus field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RemoteAssistancePartner) GetOnboardingStatus() AnyOfmicrosoftGraphRemoteAssistanceOnboardingStatus {
	if o == nil  {
		var ret AnyOfmicrosoftGraphRemoteAssistanceOnboardingStatus
		return ret
	}
	return o.OnboardingStatus
}

// GetOnboardingStatusOk returns a tuple with the OnboardingStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RemoteAssistancePartner) GetOnboardingStatusOk() (*AnyOfmicrosoftGraphRemoteAssistanceOnboardingStatus, bool) {
	if o == nil || o.OnboardingStatus == nil {
		return nil, false
	}
	return &o.OnboardingStatus, true
}

// HasOnboardingStatus returns a boolean if a field has been set.
func (o *RemoteAssistancePartner) HasOnboardingStatus() bool {
	if o != nil && o.OnboardingStatus != nil {
		return true
	}

	return false
}

// SetOnboardingStatus gets a reference to the given AnyOfmicrosoftGraphRemoteAssistanceOnboardingStatus and assigns it to the OnboardingStatus field.
func (o *RemoteAssistancePartner) SetOnboardingStatus(v AnyOfmicrosoftGraphRemoteAssistanceOnboardingStatus) {
	o.OnboardingStatus = v
}

// GetOnboardingUrl returns the OnboardingUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RemoteAssistancePartner) GetOnboardingUrl() string {
	if o == nil || o.OnboardingUrl.Get() == nil {
		var ret string
		return ret
	}
	return *o.OnboardingUrl.Get()
}

// GetOnboardingUrlOk returns a tuple with the OnboardingUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RemoteAssistancePartner) GetOnboardingUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.OnboardingUrl.Get(), o.OnboardingUrl.IsSet()
}

// HasOnboardingUrl returns a boolean if a field has been set.
func (o *RemoteAssistancePartner) HasOnboardingUrl() bool {
	if o != nil && o.OnboardingUrl.IsSet() {
		return true
	}

	return false
}

// SetOnboardingUrl gets a reference to the given NullableString and assigns it to the OnboardingUrl field.
func (o *RemoteAssistancePartner) SetOnboardingUrl(v string) {
	o.OnboardingUrl.Set(&v)
}
// SetOnboardingUrlNil sets the value for OnboardingUrl to be an explicit nil
func (o *RemoteAssistancePartner) SetOnboardingUrlNil() {
	o.OnboardingUrl.Set(nil)
}

// UnsetOnboardingUrl ensures that no value is present for OnboardingUrl, not even an explicit nil
func (o *RemoteAssistancePartner) UnsetOnboardingUrl() {
	o.OnboardingUrl.Unset()
}

func (o RemoteAssistancePartner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DisplayName.IsSet() {
		toSerialize["displayName"] = o.DisplayName.Get()
	}
	if o.LastConnectionDateTime != nil {
		toSerialize["lastConnectionDateTime"] = o.LastConnectionDateTime
	}
	if o.OnboardingStatus != nil {
		toSerialize["onboardingStatus"] = o.OnboardingStatus
	}
	if o.OnboardingUrl.IsSet() {
		toSerialize["onboardingUrl"] = o.OnboardingUrl.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableRemoteAssistancePartner struct {
	value *RemoteAssistancePartner
	isSet bool
}

func (v NullableRemoteAssistancePartner) Get() *RemoteAssistancePartner {
	return v.value
}

func (v *NullableRemoteAssistancePartner) Set(val *RemoteAssistancePartner) {
	v.value = val
	v.isSet = true
}

func (v NullableRemoteAssistancePartner) IsSet() bool {
	return v.isSet
}

func (v *NullableRemoteAssistancePartner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRemoteAssistancePartner(val *RemoteAssistancePartner) *NullableRemoteAssistancePartner {
	return &NullableRemoteAssistancePartner{value: val, isSet: true}
}

func (v NullableRemoteAssistancePartner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRemoteAssistancePartner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


