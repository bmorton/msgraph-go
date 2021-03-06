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

// MicrosoftGraphWebApplication struct for MicrosoftGraphWebApplication
type MicrosoftGraphWebApplication struct {
	// Home page or landing page of the application.
	HomePageUrl NullableString `json:"homePageUrl,omitempty"`
	// Specifies whether this web application can request tokens using the OAuth 2.0 implicit flow.
	ImplicitGrantSettings AnyOfmicrosoftGraphImplicitGrantSettings `json:"implicitGrantSettings,omitempty"`
	// Specifies the URL that will be used by Microsoft's authorization service to logout an user using front-channel, back-channel or SAML logout protocols.
	LogoutUrl NullableString `json:"logoutUrl,omitempty"`
	// Specifies the URLs where user tokens are sent for sign-in, or the redirect URIs where OAuth 2.0 authorization codes and access tokens are sent.
	RedirectUris *[]string `json:"redirectUris,omitempty"`
}

// NewMicrosoftGraphWebApplication instantiates a new MicrosoftGraphWebApplication object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphWebApplication() *MicrosoftGraphWebApplication {
	this := MicrosoftGraphWebApplication{}
	return &this
}

// NewMicrosoftGraphWebApplicationWithDefaults instantiates a new MicrosoftGraphWebApplication object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphWebApplicationWithDefaults() *MicrosoftGraphWebApplication {
	this := MicrosoftGraphWebApplication{}
	return &this
}

// GetHomePageUrl returns the HomePageUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphWebApplication) GetHomePageUrl() string {
	if o == nil || o.HomePageUrl.Get() == nil {
		var ret string
		return ret
	}
	return *o.HomePageUrl.Get()
}

// GetHomePageUrlOk returns a tuple with the HomePageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphWebApplication) GetHomePageUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.HomePageUrl.Get(), o.HomePageUrl.IsSet()
}

// HasHomePageUrl returns a boolean if a field has been set.
func (o *MicrosoftGraphWebApplication) HasHomePageUrl() bool {
	if o != nil && o.HomePageUrl.IsSet() {
		return true
	}

	return false
}

// SetHomePageUrl gets a reference to the given NullableString and assigns it to the HomePageUrl field.
func (o *MicrosoftGraphWebApplication) SetHomePageUrl(v string) {
	o.HomePageUrl.Set(&v)
}
// SetHomePageUrlNil sets the value for HomePageUrl to be an explicit nil
func (o *MicrosoftGraphWebApplication) SetHomePageUrlNil() {
	o.HomePageUrl.Set(nil)
}

// UnsetHomePageUrl ensures that no value is present for HomePageUrl, not even an explicit nil
func (o *MicrosoftGraphWebApplication) UnsetHomePageUrl() {
	o.HomePageUrl.Unset()
}

// GetImplicitGrantSettings returns the ImplicitGrantSettings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphWebApplication) GetImplicitGrantSettings() AnyOfmicrosoftGraphImplicitGrantSettings {
	if o == nil  {
		var ret AnyOfmicrosoftGraphImplicitGrantSettings
		return ret
	}
	return o.ImplicitGrantSettings
}

// GetImplicitGrantSettingsOk returns a tuple with the ImplicitGrantSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphWebApplication) GetImplicitGrantSettingsOk() (*AnyOfmicrosoftGraphImplicitGrantSettings, bool) {
	if o == nil || o.ImplicitGrantSettings == nil {
		return nil, false
	}
	return &o.ImplicitGrantSettings, true
}

// HasImplicitGrantSettings returns a boolean if a field has been set.
func (o *MicrosoftGraphWebApplication) HasImplicitGrantSettings() bool {
	if o != nil && o.ImplicitGrantSettings != nil {
		return true
	}

	return false
}

// SetImplicitGrantSettings gets a reference to the given AnyOfmicrosoftGraphImplicitGrantSettings and assigns it to the ImplicitGrantSettings field.
func (o *MicrosoftGraphWebApplication) SetImplicitGrantSettings(v AnyOfmicrosoftGraphImplicitGrantSettings) {
	o.ImplicitGrantSettings = v
}

// GetLogoutUrl returns the LogoutUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphWebApplication) GetLogoutUrl() string {
	if o == nil || o.LogoutUrl.Get() == nil {
		var ret string
		return ret
	}
	return *o.LogoutUrl.Get()
}

// GetLogoutUrlOk returns a tuple with the LogoutUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphWebApplication) GetLogoutUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LogoutUrl.Get(), o.LogoutUrl.IsSet()
}

// HasLogoutUrl returns a boolean if a field has been set.
func (o *MicrosoftGraphWebApplication) HasLogoutUrl() bool {
	if o != nil && o.LogoutUrl.IsSet() {
		return true
	}

	return false
}

// SetLogoutUrl gets a reference to the given NullableString and assigns it to the LogoutUrl field.
func (o *MicrosoftGraphWebApplication) SetLogoutUrl(v string) {
	o.LogoutUrl.Set(&v)
}
// SetLogoutUrlNil sets the value for LogoutUrl to be an explicit nil
func (o *MicrosoftGraphWebApplication) SetLogoutUrlNil() {
	o.LogoutUrl.Set(nil)
}

// UnsetLogoutUrl ensures that no value is present for LogoutUrl, not even an explicit nil
func (o *MicrosoftGraphWebApplication) UnsetLogoutUrl() {
	o.LogoutUrl.Unset()
}

// GetRedirectUris returns the RedirectUris field value if set, zero value otherwise.
func (o *MicrosoftGraphWebApplication) GetRedirectUris() []string {
	if o == nil || o.RedirectUris == nil {
		var ret []string
		return ret
	}
	return *o.RedirectUris
}

// GetRedirectUrisOk returns a tuple with the RedirectUris field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphWebApplication) GetRedirectUrisOk() (*[]string, bool) {
	if o == nil || o.RedirectUris == nil {
		return nil, false
	}
	return o.RedirectUris, true
}

// HasRedirectUris returns a boolean if a field has been set.
func (o *MicrosoftGraphWebApplication) HasRedirectUris() bool {
	if o != nil && o.RedirectUris != nil {
		return true
	}

	return false
}

// SetRedirectUris gets a reference to the given []string and assigns it to the RedirectUris field.
func (o *MicrosoftGraphWebApplication) SetRedirectUris(v []string) {
	o.RedirectUris = &v
}

func (o MicrosoftGraphWebApplication) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.HomePageUrl.IsSet() {
		toSerialize["homePageUrl"] = o.HomePageUrl.Get()
	}
	if o.ImplicitGrantSettings != nil {
		toSerialize["implicitGrantSettings"] = o.ImplicitGrantSettings
	}
	if o.LogoutUrl.IsSet() {
		toSerialize["logoutUrl"] = o.LogoutUrl.Get()
	}
	if o.RedirectUris != nil {
		toSerialize["redirectUris"] = o.RedirectUris
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphWebApplication struct {
	value *MicrosoftGraphWebApplication
	isSet bool
}

func (v NullableMicrosoftGraphWebApplication) Get() *MicrosoftGraphWebApplication {
	return v.value
}

func (v *NullableMicrosoftGraphWebApplication) Set(val *MicrosoftGraphWebApplication) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphWebApplication) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphWebApplication) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphWebApplication(val *MicrosoftGraphWebApplication) *NullableMicrosoftGraphWebApplication {
	return &NullableMicrosoftGraphWebApplication{value: val, isSet: true}
}

func (v NullableMicrosoftGraphWebApplication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphWebApplication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


