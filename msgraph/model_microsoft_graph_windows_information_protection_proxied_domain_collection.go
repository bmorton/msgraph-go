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

// MicrosoftGraphWindowsInformationProtectionProxiedDomainCollection Windows Information Protection Proxied Domain Collection
type MicrosoftGraphWindowsInformationProtectionProxiedDomainCollection struct {
	// Display name
	DisplayName *string `json:"displayName,omitempty"`
	// Collection of proxied domains
	ProxiedDomains *[]MicrosoftGraphProxiedDomain `json:"proxiedDomains,omitempty"`
}

// NewMicrosoftGraphWindowsInformationProtectionProxiedDomainCollection instantiates a new MicrosoftGraphWindowsInformationProtectionProxiedDomainCollection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphWindowsInformationProtectionProxiedDomainCollection() *MicrosoftGraphWindowsInformationProtectionProxiedDomainCollection {
	this := MicrosoftGraphWindowsInformationProtectionProxiedDomainCollection{}
	return &this
}

// NewMicrosoftGraphWindowsInformationProtectionProxiedDomainCollectionWithDefaults instantiates a new MicrosoftGraphWindowsInformationProtectionProxiedDomainCollection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphWindowsInformationProtectionProxiedDomainCollectionWithDefaults() *MicrosoftGraphWindowsInformationProtectionProxiedDomainCollection {
	this := MicrosoftGraphWindowsInformationProtectionProxiedDomainCollection{}
	return &this
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *MicrosoftGraphWindowsInformationProtectionProxiedDomainCollection) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphWindowsInformationProtectionProxiedDomainCollection) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *MicrosoftGraphWindowsInformationProtectionProxiedDomainCollection) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *MicrosoftGraphWindowsInformationProtectionProxiedDomainCollection) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetProxiedDomains returns the ProxiedDomains field value if set, zero value otherwise.
func (o *MicrosoftGraphWindowsInformationProtectionProxiedDomainCollection) GetProxiedDomains() []MicrosoftGraphProxiedDomain {
	if o == nil || o.ProxiedDomains == nil {
		var ret []MicrosoftGraphProxiedDomain
		return ret
	}
	return *o.ProxiedDomains
}

// GetProxiedDomainsOk returns a tuple with the ProxiedDomains field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphWindowsInformationProtectionProxiedDomainCollection) GetProxiedDomainsOk() (*[]MicrosoftGraphProxiedDomain, bool) {
	if o == nil || o.ProxiedDomains == nil {
		return nil, false
	}
	return o.ProxiedDomains, true
}

// HasProxiedDomains returns a boolean if a field has been set.
func (o *MicrosoftGraphWindowsInformationProtectionProxiedDomainCollection) HasProxiedDomains() bool {
	if o != nil && o.ProxiedDomains != nil {
		return true
	}

	return false
}

// SetProxiedDomains gets a reference to the given []MicrosoftGraphProxiedDomain and assigns it to the ProxiedDomains field.
func (o *MicrosoftGraphWindowsInformationProtectionProxiedDomainCollection) SetProxiedDomains(v []MicrosoftGraphProxiedDomain) {
	o.ProxiedDomains = &v
}

func (o MicrosoftGraphWindowsInformationProtectionProxiedDomainCollection) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DisplayName != nil {
		toSerialize["displayName"] = o.DisplayName
	}
	if o.ProxiedDomains != nil {
		toSerialize["proxiedDomains"] = o.ProxiedDomains
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphWindowsInformationProtectionProxiedDomainCollection struct {
	value *MicrosoftGraphWindowsInformationProtectionProxiedDomainCollection
	isSet bool
}

func (v NullableMicrosoftGraphWindowsInformationProtectionProxiedDomainCollection) Get() *MicrosoftGraphWindowsInformationProtectionProxiedDomainCollection {
	return v.value
}

func (v *NullableMicrosoftGraphWindowsInformationProtectionProxiedDomainCollection) Set(val *MicrosoftGraphWindowsInformationProtectionProxiedDomainCollection) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphWindowsInformationProtectionProxiedDomainCollection) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphWindowsInformationProtectionProxiedDomainCollection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphWindowsInformationProtectionProxiedDomainCollection(val *MicrosoftGraphWindowsInformationProtectionProxiedDomainCollection) *NullableMicrosoftGraphWindowsInformationProtectionProxiedDomainCollection {
	return &NullableMicrosoftGraphWindowsInformationProtectionProxiedDomainCollection{value: val, isSet: true}
}

func (v NullableMicrosoftGraphWindowsInformationProtectionProxiedDomainCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphWindowsInformationProtectionProxiedDomainCollection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


