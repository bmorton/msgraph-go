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

// IdentityProviderBase struct for IdentityProviderBase
type IdentityProviderBase struct {
	// The display name of the identity provider.
	DisplayName NullableString `json:"displayName,omitempty"`
}

// NewIdentityProviderBase instantiates a new IdentityProviderBase object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityProviderBase() *IdentityProviderBase {
	this := IdentityProviderBase{}
	return &this
}

// NewIdentityProviderBaseWithDefaults instantiates a new IdentityProviderBase object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityProviderBaseWithDefaults() *IdentityProviderBase {
	this := IdentityProviderBase{}
	return &this
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IdentityProviderBase) GetDisplayName() string {
	if o == nil || o.DisplayName.Get() == nil {
		var ret string
		return ret
	}
	return *o.DisplayName.Get()
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentityProviderBase) GetDisplayNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DisplayName.Get(), o.DisplayName.IsSet()
}

// HasDisplayName returns a boolean if a field has been set.
func (o *IdentityProviderBase) HasDisplayName() bool {
	if o != nil && o.DisplayName.IsSet() {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given NullableString and assigns it to the DisplayName field.
func (o *IdentityProviderBase) SetDisplayName(v string) {
	o.DisplayName.Set(&v)
}
// SetDisplayNameNil sets the value for DisplayName to be an explicit nil
func (o *IdentityProviderBase) SetDisplayNameNil() {
	o.DisplayName.Set(nil)
}

// UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
func (o *IdentityProviderBase) UnsetDisplayName() {
	o.DisplayName.Unset()
}

func (o IdentityProviderBase) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DisplayName.IsSet() {
		toSerialize["displayName"] = o.DisplayName.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableIdentityProviderBase struct {
	value *IdentityProviderBase
	isSet bool
}

func (v NullableIdentityProviderBase) Get() *IdentityProviderBase {
	return v.value
}

func (v *NullableIdentityProviderBase) Set(val *IdentityProviderBase) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityProviderBase) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityProviderBase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityProviderBase(val *IdentityProviderBase) *NullableIdentityProviderBase {
	return &NullableIdentityProviderBase{value: val, isSet: true}
}

func (v NullableIdentityProviderBase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityProviderBase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

