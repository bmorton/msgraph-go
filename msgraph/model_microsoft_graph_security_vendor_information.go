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

// MicrosoftGraphSecurityVendorInformation struct for MicrosoftGraphSecurityVendorInformation
type MicrosoftGraphSecurityVendorInformation struct {
	// Specific provider (product/service - not vendor company); for example, WindowsDefenderATP.
	Provider NullableString `json:"provider,omitempty"`
	// Version of the provider or subprovider, if it exists, that generated the alert. Required
	ProviderVersion NullableString `json:"providerVersion,omitempty"`
	// Specific subprovider (under aggregating provider); for example, WindowsDefenderATP.SmartScreen.
	SubProvider NullableString `json:"subProvider,omitempty"`
	// Name of the alert vendor (for example, Microsoft, Dell, FireEye). Required
	Vendor NullableString `json:"vendor,omitempty"`
}

// NewMicrosoftGraphSecurityVendorInformation instantiates a new MicrosoftGraphSecurityVendorInformation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphSecurityVendorInformation() *MicrosoftGraphSecurityVendorInformation {
	this := MicrosoftGraphSecurityVendorInformation{}
	return &this
}

// NewMicrosoftGraphSecurityVendorInformationWithDefaults instantiates a new MicrosoftGraphSecurityVendorInformation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphSecurityVendorInformationWithDefaults() *MicrosoftGraphSecurityVendorInformation {
	this := MicrosoftGraphSecurityVendorInformation{}
	return &this
}

// GetProvider returns the Provider field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphSecurityVendorInformation) GetProvider() string {
	if o == nil || o.Provider.Get() == nil {
		var ret string
		return ret
	}
	return *o.Provider.Get()
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphSecurityVendorInformation) GetProviderOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Provider.Get(), o.Provider.IsSet()
}

// HasProvider returns a boolean if a field has been set.
func (o *MicrosoftGraphSecurityVendorInformation) HasProvider() bool {
	if o != nil && o.Provider.IsSet() {
		return true
	}

	return false
}

// SetProvider gets a reference to the given NullableString and assigns it to the Provider field.
func (o *MicrosoftGraphSecurityVendorInformation) SetProvider(v string) {
	o.Provider.Set(&v)
}
// SetProviderNil sets the value for Provider to be an explicit nil
func (o *MicrosoftGraphSecurityVendorInformation) SetProviderNil() {
	o.Provider.Set(nil)
}

// UnsetProvider ensures that no value is present for Provider, not even an explicit nil
func (o *MicrosoftGraphSecurityVendorInformation) UnsetProvider() {
	o.Provider.Unset()
}

// GetProviderVersion returns the ProviderVersion field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphSecurityVendorInformation) GetProviderVersion() string {
	if o == nil || o.ProviderVersion.Get() == nil {
		var ret string
		return ret
	}
	return *o.ProviderVersion.Get()
}

// GetProviderVersionOk returns a tuple with the ProviderVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphSecurityVendorInformation) GetProviderVersionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ProviderVersion.Get(), o.ProviderVersion.IsSet()
}

// HasProviderVersion returns a boolean if a field has been set.
func (o *MicrosoftGraphSecurityVendorInformation) HasProviderVersion() bool {
	if o != nil && o.ProviderVersion.IsSet() {
		return true
	}

	return false
}

// SetProviderVersion gets a reference to the given NullableString and assigns it to the ProviderVersion field.
func (o *MicrosoftGraphSecurityVendorInformation) SetProviderVersion(v string) {
	o.ProviderVersion.Set(&v)
}
// SetProviderVersionNil sets the value for ProviderVersion to be an explicit nil
func (o *MicrosoftGraphSecurityVendorInformation) SetProviderVersionNil() {
	o.ProviderVersion.Set(nil)
}

// UnsetProviderVersion ensures that no value is present for ProviderVersion, not even an explicit nil
func (o *MicrosoftGraphSecurityVendorInformation) UnsetProviderVersion() {
	o.ProviderVersion.Unset()
}

// GetSubProvider returns the SubProvider field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphSecurityVendorInformation) GetSubProvider() string {
	if o == nil || o.SubProvider.Get() == nil {
		var ret string
		return ret
	}
	return *o.SubProvider.Get()
}

// GetSubProviderOk returns a tuple with the SubProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphSecurityVendorInformation) GetSubProviderOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SubProvider.Get(), o.SubProvider.IsSet()
}

// HasSubProvider returns a boolean if a field has been set.
func (o *MicrosoftGraphSecurityVendorInformation) HasSubProvider() bool {
	if o != nil && o.SubProvider.IsSet() {
		return true
	}

	return false
}

// SetSubProvider gets a reference to the given NullableString and assigns it to the SubProvider field.
func (o *MicrosoftGraphSecurityVendorInformation) SetSubProvider(v string) {
	o.SubProvider.Set(&v)
}
// SetSubProviderNil sets the value for SubProvider to be an explicit nil
func (o *MicrosoftGraphSecurityVendorInformation) SetSubProviderNil() {
	o.SubProvider.Set(nil)
}

// UnsetSubProvider ensures that no value is present for SubProvider, not even an explicit nil
func (o *MicrosoftGraphSecurityVendorInformation) UnsetSubProvider() {
	o.SubProvider.Unset()
}

// GetVendor returns the Vendor field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphSecurityVendorInformation) GetVendor() string {
	if o == nil || o.Vendor.Get() == nil {
		var ret string
		return ret
	}
	return *o.Vendor.Get()
}

// GetVendorOk returns a tuple with the Vendor field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphSecurityVendorInformation) GetVendorOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Vendor.Get(), o.Vendor.IsSet()
}

// HasVendor returns a boolean if a field has been set.
func (o *MicrosoftGraphSecurityVendorInformation) HasVendor() bool {
	if o != nil && o.Vendor.IsSet() {
		return true
	}

	return false
}

// SetVendor gets a reference to the given NullableString and assigns it to the Vendor field.
func (o *MicrosoftGraphSecurityVendorInformation) SetVendor(v string) {
	o.Vendor.Set(&v)
}
// SetVendorNil sets the value for Vendor to be an explicit nil
func (o *MicrosoftGraphSecurityVendorInformation) SetVendorNil() {
	o.Vendor.Set(nil)
}

// UnsetVendor ensures that no value is present for Vendor, not even an explicit nil
func (o *MicrosoftGraphSecurityVendorInformation) UnsetVendor() {
	o.Vendor.Unset()
}

func (o MicrosoftGraphSecurityVendorInformation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Provider.IsSet() {
		toSerialize["provider"] = o.Provider.Get()
	}
	if o.ProviderVersion.IsSet() {
		toSerialize["providerVersion"] = o.ProviderVersion.Get()
	}
	if o.SubProvider.IsSet() {
		toSerialize["subProvider"] = o.SubProvider.Get()
	}
	if o.Vendor.IsSet() {
		toSerialize["vendor"] = o.Vendor.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphSecurityVendorInformation struct {
	value *MicrosoftGraphSecurityVendorInformation
	isSet bool
}

func (v NullableMicrosoftGraphSecurityVendorInformation) Get() *MicrosoftGraphSecurityVendorInformation {
	return v.value
}

func (v *NullableMicrosoftGraphSecurityVendorInformation) Set(val *MicrosoftGraphSecurityVendorInformation) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphSecurityVendorInformation) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphSecurityVendorInformation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphSecurityVendorInformation(val *MicrosoftGraphSecurityVendorInformation) *NullableMicrosoftGraphSecurityVendorInformation {
	return &NullableMicrosoftGraphSecurityVendorInformation{value: val, isSet: true}
}

func (v NullableMicrosoftGraphSecurityVendorInformation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphSecurityVendorInformation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

