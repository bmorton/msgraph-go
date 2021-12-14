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

// MicrosoftGraphIntuneBrand intuneBrand contains data which is used in customizing the appearance of the Company Portal applications as well as the end user web portal.
type MicrosoftGraphIntuneBrand struct {
	// Email address of the person/organization responsible for IT support.
	ContactITEmailAddress NullableString `json:"contactITEmailAddress,omitempty"`
	// Name of the person/organization responsible for IT support.
	ContactITName NullableString `json:"contactITName,omitempty"`
	// Text comments regarding the person/organization responsible for IT support.
	ContactITNotes NullableString `json:"contactITNotes,omitempty"`
	// Phone number of the person/organization responsible for IT support.
	ContactITPhoneNumber NullableString `json:"contactITPhoneNumber,omitempty"`
	// Logo image displayed in Company Portal apps which have a dark background behind the logo.
	DarkBackgroundLogo AnyOfmicrosoftGraphMimeContent `json:"darkBackgroundLogo,omitempty"`
	// Company/organization name that is displayed to end users.
	DisplayName NullableString `json:"displayName,omitempty"`
	// Logo image displayed in Company Portal apps which have a light background behind the logo.
	LightBackgroundLogo AnyOfmicrosoftGraphMimeContent `json:"lightBackgroundLogo,omitempty"`
	// Display name of the company/organization’s IT helpdesk site.
	OnlineSupportSiteName NullableString `json:"onlineSupportSiteName,omitempty"`
	// URL to the company/organization’s IT helpdesk site.
	OnlineSupportSiteUrl NullableString `json:"onlineSupportSiteUrl,omitempty"`
	// URL to the company/organization’s privacy policy.
	PrivacyUrl NullableString `json:"privacyUrl,omitempty"`
	// Boolean that represents whether the administrator-supplied display name will be shown next to the logo image.
	ShowDisplayNameNextToLogo *bool `json:"showDisplayNameNextToLogo,omitempty"`
	// Boolean that represents whether the administrator-supplied logo images are shown or not shown.
	ShowLogo *bool `json:"showLogo,omitempty"`
	// Boolean that represents whether the administrator-supplied display name will be shown next to the logo image.
	ShowNameNextToLogo *bool `json:"showNameNextToLogo,omitempty"`
	// Primary theme color used in the Company Portal applications and web portal.
	ThemeColor AnyOfmicrosoftGraphRgbColor `json:"themeColor,omitempty"`
}

// NewMicrosoftGraphIntuneBrand instantiates a new MicrosoftGraphIntuneBrand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphIntuneBrand() *MicrosoftGraphIntuneBrand {
	this := MicrosoftGraphIntuneBrand{}
	return &this
}

// NewMicrosoftGraphIntuneBrandWithDefaults instantiates a new MicrosoftGraphIntuneBrand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphIntuneBrandWithDefaults() *MicrosoftGraphIntuneBrand {
	this := MicrosoftGraphIntuneBrand{}
	return &this
}

// GetContactITEmailAddress returns the ContactITEmailAddress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphIntuneBrand) GetContactITEmailAddress() string {
	if o == nil || o.ContactITEmailAddress.Get() == nil {
		var ret string
		return ret
	}
	return *o.ContactITEmailAddress.Get()
}

// GetContactITEmailAddressOk returns a tuple with the ContactITEmailAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphIntuneBrand) GetContactITEmailAddressOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ContactITEmailAddress.Get(), o.ContactITEmailAddress.IsSet()
}

// HasContactITEmailAddress returns a boolean if a field has been set.
func (o *MicrosoftGraphIntuneBrand) HasContactITEmailAddress() bool {
	if o != nil && o.ContactITEmailAddress.IsSet() {
		return true
	}

	return false
}

// SetContactITEmailAddress gets a reference to the given NullableString and assigns it to the ContactITEmailAddress field.
func (o *MicrosoftGraphIntuneBrand) SetContactITEmailAddress(v string) {
	o.ContactITEmailAddress.Set(&v)
}
// SetContactITEmailAddressNil sets the value for ContactITEmailAddress to be an explicit nil
func (o *MicrosoftGraphIntuneBrand) SetContactITEmailAddressNil() {
	o.ContactITEmailAddress.Set(nil)
}

// UnsetContactITEmailAddress ensures that no value is present for ContactITEmailAddress, not even an explicit nil
func (o *MicrosoftGraphIntuneBrand) UnsetContactITEmailAddress() {
	o.ContactITEmailAddress.Unset()
}

// GetContactITName returns the ContactITName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphIntuneBrand) GetContactITName() string {
	if o == nil || o.ContactITName.Get() == nil {
		var ret string
		return ret
	}
	return *o.ContactITName.Get()
}

// GetContactITNameOk returns a tuple with the ContactITName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphIntuneBrand) GetContactITNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ContactITName.Get(), o.ContactITName.IsSet()
}

// HasContactITName returns a boolean if a field has been set.
func (o *MicrosoftGraphIntuneBrand) HasContactITName() bool {
	if o != nil && o.ContactITName.IsSet() {
		return true
	}

	return false
}

// SetContactITName gets a reference to the given NullableString and assigns it to the ContactITName field.
func (o *MicrosoftGraphIntuneBrand) SetContactITName(v string) {
	o.ContactITName.Set(&v)
}
// SetContactITNameNil sets the value for ContactITName to be an explicit nil
func (o *MicrosoftGraphIntuneBrand) SetContactITNameNil() {
	o.ContactITName.Set(nil)
}

// UnsetContactITName ensures that no value is present for ContactITName, not even an explicit nil
func (o *MicrosoftGraphIntuneBrand) UnsetContactITName() {
	o.ContactITName.Unset()
}

// GetContactITNotes returns the ContactITNotes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphIntuneBrand) GetContactITNotes() string {
	if o == nil || o.ContactITNotes.Get() == nil {
		var ret string
		return ret
	}
	return *o.ContactITNotes.Get()
}

// GetContactITNotesOk returns a tuple with the ContactITNotes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphIntuneBrand) GetContactITNotesOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ContactITNotes.Get(), o.ContactITNotes.IsSet()
}

// HasContactITNotes returns a boolean if a field has been set.
func (o *MicrosoftGraphIntuneBrand) HasContactITNotes() bool {
	if o != nil && o.ContactITNotes.IsSet() {
		return true
	}

	return false
}

// SetContactITNotes gets a reference to the given NullableString and assigns it to the ContactITNotes field.
func (o *MicrosoftGraphIntuneBrand) SetContactITNotes(v string) {
	o.ContactITNotes.Set(&v)
}
// SetContactITNotesNil sets the value for ContactITNotes to be an explicit nil
func (o *MicrosoftGraphIntuneBrand) SetContactITNotesNil() {
	o.ContactITNotes.Set(nil)
}

// UnsetContactITNotes ensures that no value is present for ContactITNotes, not even an explicit nil
func (o *MicrosoftGraphIntuneBrand) UnsetContactITNotes() {
	o.ContactITNotes.Unset()
}

// GetContactITPhoneNumber returns the ContactITPhoneNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphIntuneBrand) GetContactITPhoneNumber() string {
	if o == nil || o.ContactITPhoneNumber.Get() == nil {
		var ret string
		return ret
	}
	return *o.ContactITPhoneNumber.Get()
}

// GetContactITPhoneNumberOk returns a tuple with the ContactITPhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphIntuneBrand) GetContactITPhoneNumberOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ContactITPhoneNumber.Get(), o.ContactITPhoneNumber.IsSet()
}

// HasContactITPhoneNumber returns a boolean if a field has been set.
func (o *MicrosoftGraphIntuneBrand) HasContactITPhoneNumber() bool {
	if o != nil && o.ContactITPhoneNumber.IsSet() {
		return true
	}

	return false
}

// SetContactITPhoneNumber gets a reference to the given NullableString and assigns it to the ContactITPhoneNumber field.
func (o *MicrosoftGraphIntuneBrand) SetContactITPhoneNumber(v string) {
	o.ContactITPhoneNumber.Set(&v)
}
// SetContactITPhoneNumberNil sets the value for ContactITPhoneNumber to be an explicit nil
func (o *MicrosoftGraphIntuneBrand) SetContactITPhoneNumberNil() {
	o.ContactITPhoneNumber.Set(nil)
}

// UnsetContactITPhoneNumber ensures that no value is present for ContactITPhoneNumber, not even an explicit nil
func (o *MicrosoftGraphIntuneBrand) UnsetContactITPhoneNumber() {
	o.ContactITPhoneNumber.Unset()
}

// GetDarkBackgroundLogo returns the DarkBackgroundLogo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphIntuneBrand) GetDarkBackgroundLogo() AnyOfmicrosoftGraphMimeContent {
	if o == nil  {
		var ret AnyOfmicrosoftGraphMimeContent
		return ret
	}
	return o.DarkBackgroundLogo
}

// GetDarkBackgroundLogoOk returns a tuple with the DarkBackgroundLogo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphIntuneBrand) GetDarkBackgroundLogoOk() (*AnyOfmicrosoftGraphMimeContent, bool) {
	if o == nil || o.DarkBackgroundLogo == nil {
		return nil, false
	}
	return &o.DarkBackgroundLogo, true
}

// HasDarkBackgroundLogo returns a boolean if a field has been set.
func (o *MicrosoftGraphIntuneBrand) HasDarkBackgroundLogo() bool {
	if o != nil && o.DarkBackgroundLogo != nil {
		return true
	}

	return false
}

// SetDarkBackgroundLogo gets a reference to the given AnyOfmicrosoftGraphMimeContent and assigns it to the DarkBackgroundLogo field.
func (o *MicrosoftGraphIntuneBrand) SetDarkBackgroundLogo(v AnyOfmicrosoftGraphMimeContent) {
	o.DarkBackgroundLogo = v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphIntuneBrand) GetDisplayName() string {
	if o == nil || o.DisplayName.Get() == nil {
		var ret string
		return ret
	}
	return *o.DisplayName.Get()
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphIntuneBrand) GetDisplayNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DisplayName.Get(), o.DisplayName.IsSet()
}

// HasDisplayName returns a boolean if a field has been set.
func (o *MicrosoftGraphIntuneBrand) HasDisplayName() bool {
	if o != nil && o.DisplayName.IsSet() {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given NullableString and assigns it to the DisplayName field.
func (o *MicrosoftGraphIntuneBrand) SetDisplayName(v string) {
	o.DisplayName.Set(&v)
}
// SetDisplayNameNil sets the value for DisplayName to be an explicit nil
func (o *MicrosoftGraphIntuneBrand) SetDisplayNameNil() {
	o.DisplayName.Set(nil)
}

// UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
func (o *MicrosoftGraphIntuneBrand) UnsetDisplayName() {
	o.DisplayName.Unset()
}

// GetLightBackgroundLogo returns the LightBackgroundLogo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphIntuneBrand) GetLightBackgroundLogo() AnyOfmicrosoftGraphMimeContent {
	if o == nil  {
		var ret AnyOfmicrosoftGraphMimeContent
		return ret
	}
	return o.LightBackgroundLogo
}

// GetLightBackgroundLogoOk returns a tuple with the LightBackgroundLogo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphIntuneBrand) GetLightBackgroundLogoOk() (*AnyOfmicrosoftGraphMimeContent, bool) {
	if o == nil || o.LightBackgroundLogo == nil {
		return nil, false
	}
	return &o.LightBackgroundLogo, true
}

// HasLightBackgroundLogo returns a boolean if a field has been set.
func (o *MicrosoftGraphIntuneBrand) HasLightBackgroundLogo() bool {
	if o != nil && o.LightBackgroundLogo != nil {
		return true
	}

	return false
}

// SetLightBackgroundLogo gets a reference to the given AnyOfmicrosoftGraphMimeContent and assigns it to the LightBackgroundLogo field.
func (o *MicrosoftGraphIntuneBrand) SetLightBackgroundLogo(v AnyOfmicrosoftGraphMimeContent) {
	o.LightBackgroundLogo = v
}

// GetOnlineSupportSiteName returns the OnlineSupportSiteName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphIntuneBrand) GetOnlineSupportSiteName() string {
	if o == nil || o.OnlineSupportSiteName.Get() == nil {
		var ret string
		return ret
	}
	return *o.OnlineSupportSiteName.Get()
}

// GetOnlineSupportSiteNameOk returns a tuple with the OnlineSupportSiteName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphIntuneBrand) GetOnlineSupportSiteNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.OnlineSupportSiteName.Get(), o.OnlineSupportSiteName.IsSet()
}

// HasOnlineSupportSiteName returns a boolean if a field has been set.
func (o *MicrosoftGraphIntuneBrand) HasOnlineSupportSiteName() bool {
	if o != nil && o.OnlineSupportSiteName.IsSet() {
		return true
	}

	return false
}

// SetOnlineSupportSiteName gets a reference to the given NullableString and assigns it to the OnlineSupportSiteName field.
func (o *MicrosoftGraphIntuneBrand) SetOnlineSupportSiteName(v string) {
	o.OnlineSupportSiteName.Set(&v)
}
// SetOnlineSupportSiteNameNil sets the value for OnlineSupportSiteName to be an explicit nil
func (o *MicrosoftGraphIntuneBrand) SetOnlineSupportSiteNameNil() {
	o.OnlineSupportSiteName.Set(nil)
}

// UnsetOnlineSupportSiteName ensures that no value is present for OnlineSupportSiteName, not even an explicit nil
func (o *MicrosoftGraphIntuneBrand) UnsetOnlineSupportSiteName() {
	o.OnlineSupportSiteName.Unset()
}

// GetOnlineSupportSiteUrl returns the OnlineSupportSiteUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphIntuneBrand) GetOnlineSupportSiteUrl() string {
	if o == nil || o.OnlineSupportSiteUrl.Get() == nil {
		var ret string
		return ret
	}
	return *o.OnlineSupportSiteUrl.Get()
}

// GetOnlineSupportSiteUrlOk returns a tuple with the OnlineSupportSiteUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphIntuneBrand) GetOnlineSupportSiteUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.OnlineSupportSiteUrl.Get(), o.OnlineSupportSiteUrl.IsSet()
}

// HasOnlineSupportSiteUrl returns a boolean if a field has been set.
func (o *MicrosoftGraphIntuneBrand) HasOnlineSupportSiteUrl() bool {
	if o != nil && o.OnlineSupportSiteUrl.IsSet() {
		return true
	}

	return false
}

// SetOnlineSupportSiteUrl gets a reference to the given NullableString and assigns it to the OnlineSupportSiteUrl field.
func (o *MicrosoftGraphIntuneBrand) SetOnlineSupportSiteUrl(v string) {
	o.OnlineSupportSiteUrl.Set(&v)
}
// SetOnlineSupportSiteUrlNil sets the value for OnlineSupportSiteUrl to be an explicit nil
func (o *MicrosoftGraphIntuneBrand) SetOnlineSupportSiteUrlNil() {
	o.OnlineSupportSiteUrl.Set(nil)
}

// UnsetOnlineSupportSiteUrl ensures that no value is present for OnlineSupportSiteUrl, not even an explicit nil
func (o *MicrosoftGraphIntuneBrand) UnsetOnlineSupportSiteUrl() {
	o.OnlineSupportSiteUrl.Unset()
}

// GetPrivacyUrl returns the PrivacyUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphIntuneBrand) GetPrivacyUrl() string {
	if o == nil || o.PrivacyUrl.Get() == nil {
		var ret string
		return ret
	}
	return *o.PrivacyUrl.Get()
}

// GetPrivacyUrlOk returns a tuple with the PrivacyUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphIntuneBrand) GetPrivacyUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PrivacyUrl.Get(), o.PrivacyUrl.IsSet()
}

// HasPrivacyUrl returns a boolean if a field has been set.
func (o *MicrosoftGraphIntuneBrand) HasPrivacyUrl() bool {
	if o != nil && o.PrivacyUrl.IsSet() {
		return true
	}

	return false
}

// SetPrivacyUrl gets a reference to the given NullableString and assigns it to the PrivacyUrl field.
func (o *MicrosoftGraphIntuneBrand) SetPrivacyUrl(v string) {
	o.PrivacyUrl.Set(&v)
}
// SetPrivacyUrlNil sets the value for PrivacyUrl to be an explicit nil
func (o *MicrosoftGraphIntuneBrand) SetPrivacyUrlNil() {
	o.PrivacyUrl.Set(nil)
}

// UnsetPrivacyUrl ensures that no value is present for PrivacyUrl, not even an explicit nil
func (o *MicrosoftGraphIntuneBrand) UnsetPrivacyUrl() {
	o.PrivacyUrl.Unset()
}

// GetShowDisplayNameNextToLogo returns the ShowDisplayNameNextToLogo field value if set, zero value otherwise.
func (o *MicrosoftGraphIntuneBrand) GetShowDisplayNameNextToLogo() bool {
	if o == nil || o.ShowDisplayNameNextToLogo == nil {
		var ret bool
		return ret
	}
	return *o.ShowDisplayNameNextToLogo
}

// GetShowDisplayNameNextToLogoOk returns a tuple with the ShowDisplayNameNextToLogo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphIntuneBrand) GetShowDisplayNameNextToLogoOk() (*bool, bool) {
	if o == nil || o.ShowDisplayNameNextToLogo == nil {
		return nil, false
	}
	return o.ShowDisplayNameNextToLogo, true
}

// HasShowDisplayNameNextToLogo returns a boolean if a field has been set.
func (o *MicrosoftGraphIntuneBrand) HasShowDisplayNameNextToLogo() bool {
	if o != nil && o.ShowDisplayNameNextToLogo != nil {
		return true
	}

	return false
}

// SetShowDisplayNameNextToLogo gets a reference to the given bool and assigns it to the ShowDisplayNameNextToLogo field.
func (o *MicrosoftGraphIntuneBrand) SetShowDisplayNameNextToLogo(v bool) {
	o.ShowDisplayNameNextToLogo = &v
}

// GetShowLogo returns the ShowLogo field value if set, zero value otherwise.
func (o *MicrosoftGraphIntuneBrand) GetShowLogo() bool {
	if o == nil || o.ShowLogo == nil {
		var ret bool
		return ret
	}
	return *o.ShowLogo
}

// GetShowLogoOk returns a tuple with the ShowLogo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphIntuneBrand) GetShowLogoOk() (*bool, bool) {
	if o == nil || o.ShowLogo == nil {
		return nil, false
	}
	return o.ShowLogo, true
}

// HasShowLogo returns a boolean if a field has been set.
func (o *MicrosoftGraphIntuneBrand) HasShowLogo() bool {
	if o != nil && o.ShowLogo != nil {
		return true
	}

	return false
}

// SetShowLogo gets a reference to the given bool and assigns it to the ShowLogo field.
func (o *MicrosoftGraphIntuneBrand) SetShowLogo(v bool) {
	o.ShowLogo = &v
}

// GetShowNameNextToLogo returns the ShowNameNextToLogo field value if set, zero value otherwise.
func (o *MicrosoftGraphIntuneBrand) GetShowNameNextToLogo() bool {
	if o == nil || o.ShowNameNextToLogo == nil {
		var ret bool
		return ret
	}
	return *o.ShowNameNextToLogo
}

// GetShowNameNextToLogoOk returns a tuple with the ShowNameNextToLogo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphIntuneBrand) GetShowNameNextToLogoOk() (*bool, bool) {
	if o == nil || o.ShowNameNextToLogo == nil {
		return nil, false
	}
	return o.ShowNameNextToLogo, true
}

// HasShowNameNextToLogo returns a boolean if a field has been set.
func (o *MicrosoftGraphIntuneBrand) HasShowNameNextToLogo() bool {
	if o != nil && o.ShowNameNextToLogo != nil {
		return true
	}

	return false
}

// SetShowNameNextToLogo gets a reference to the given bool and assigns it to the ShowNameNextToLogo field.
func (o *MicrosoftGraphIntuneBrand) SetShowNameNextToLogo(v bool) {
	o.ShowNameNextToLogo = &v
}

// GetThemeColor returns the ThemeColor field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphIntuneBrand) GetThemeColor() AnyOfmicrosoftGraphRgbColor {
	if o == nil  {
		var ret AnyOfmicrosoftGraphRgbColor
		return ret
	}
	return o.ThemeColor
}

// GetThemeColorOk returns a tuple with the ThemeColor field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphIntuneBrand) GetThemeColorOk() (*AnyOfmicrosoftGraphRgbColor, bool) {
	if o == nil || o.ThemeColor == nil {
		return nil, false
	}
	return &o.ThemeColor, true
}

// HasThemeColor returns a boolean if a field has been set.
func (o *MicrosoftGraphIntuneBrand) HasThemeColor() bool {
	if o != nil && o.ThemeColor != nil {
		return true
	}

	return false
}

// SetThemeColor gets a reference to the given AnyOfmicrosoftGraphRgbColor and assigns it to the ThemeColor field.
func (o *MicrosoftGraphIntuneBrand) SetThemeColor(v AnyOfmicrosoftGraphRgbColor) {
	o.ThemeColor = v
}

func (o MicrosoftGraphIntuneBrand) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ContactITEmailAddress.IsSet() {
		toSerialize["contactITEmailAddress"] = o.ContactITEmailAddress.Get()
	}
	if o.ContactITName.IsSet() {
		toSerialize["contactITName"] = o.ContactITName.Get()
	}
	if o.ContactITNotes.IsSet() {
		toSerialize["contactITNotes"] = o.ContactITNotes.Get()
	}
	if o.ContactITPhoneNumber.IsSet() {
		toSerialize["contactITPhoneNumber"] = o.ContactITPhoneNumber.Get()
	}
	if o.DarkBackgroundLogo != nil {
		toSerialize["darkBackgroundLogo"] = o.DarkBackgroundLogo
	}
	if o.DisplayName.IsSet() {
		toSerialize["displayName"] = o.DisplayName.Get()
	}
	if o.LightBackgroundLogo != nil {
		toSerialize["lightBackgroundLogo"] = o.LightBackgroundLogo
	}
	if o.OnlineSupportSiteName.IsSet() {
		toSerialize["onlineSupportSiteName"] = o.OnlineSupportSiteName.Get()
	}
	if o.OnlineSupportSiteUrl.IsSet() {
		toSerialize["onlineSupportSiteUrl"] = o.OnlineSupportSiteUrl.Get()
	}
	if o.PrivacyUrl.IsSet() {
		toSerialize["privacyUrl"] = o.PrivacyUrl.Get()
	}
	if o.ShowDisplayNameNextToLogo != nil {
		toSerialize["showDisplayNameNextToLogo"] = o.ShowDisplayNameNextToLogo
	}
	if o.ShowLogo != nil {
		toSerialize["showLogo"] = o.ShowLogo
	}
	if o.ShowNameNextToLogo != nil {
		toSerialize["showNameNextToLogo"] = o.ShowNameNextToLogo
	}
	if o.ThemeColor != nil {
		toSerialize["themeColor"] = o.ThemeColor
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphIntuneBrand struct {
	value *MicrosoftGraphIntuneBrand
	isSet bool
}

func (v NullableMicrosoftGraphIntuneBrand) Get() *MicrosoftGraphIntuneBrand {
	return v.value
}

func (v *NullableMicrosoftGraphIntuneBrand) Set(val *MicrosoftGraphIntuneBrand) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphIntuneBrand) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphIntuneBrand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphIntuneBrand(val *MicrosoftGraphIntuneBrand) *NullableMicrosoftGraphIntuneBrand {
	return &NullableMicrosoftGraphIntuneBrand{value: val, isSet: true}
}

func (v NullableMicrosoftGraphIntuneBrand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphIntuneBrand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


