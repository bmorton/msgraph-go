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

// MicrosoftGraphIdentityProviderBase struct for MicrosoftGraphIdentityProviderBase
type MicrosoftGraphIdentityProviderBase struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
	// The display name of the identity provider.
	DisplayName NullableString `json:"displayName,omitempty"`
}

// NewMicrosoftGraphIdentityProviderBase instantiates a new MicrosoftGraphIdentityProviderBase object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphIdentityProviderBase() *MicrosoftGraphIdentityProviderBase {
	this := MicrosoftGraphIdentityProviderBase{}
	return &this
}

// NewMicrosoftGraphIdentityProviderBaseWithDefaults instantiates a new MicrosoftGraphIdentityProviderBase object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphIdentityProviderBaseWithDefaults() *MicrosoftGraphIdentityProviderBase {
	this := MicrosoftGraphIdentityProviderBase{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MicrosoftGraphIdentityProviderBase) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphIdentityProviderBase) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphIdentityProviderBase) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MicrosoftGraphIdentityProviderBase) SetId(v string) {
	o.Id = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphIdentityProviderBase) GetDisplayName() string {
	if o == nil || o.DisplayName.Get() == nil {
		var ret string
		return ret
	}
	return *o.DisplayName.Get()
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphIdentityProviderBase) GetDisplayNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DisplayName.Get(), o.DisplayName.IsSet()
}

// HasDisplayName returns a boolean if a field has been set.
func (o *MicrosoftGraphIdentityProviderBase) HasDisplayName() bool {
	if o != nil && o.DisplayName.IsSet() {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given NullableString and assigns it to the DisplayName field.
func (o *MicrosoftGraphIdentityProviderBase) SetDisplayName(v string) {
	o.DisplayName.Set(&v)
}
// SetDisplayNameNil sets the value for DisplayName to be an explicit nil
func (o *MicrosoftGraphIdentityProviderBase) SetDisplayNameNil() {
	o.DisplayName.Set(nil)
}

// UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
func (o *MicrosoftGraphIdentityProviderBase) UnsetDisplayName() {
	o.DisplayName.Unset()
}

func (o MicrosoftGraphIdentityProviderBase) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.DisplayName.IsSet() {
		toSerialize["displayName"] = o.DisplayName.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphIdentityProviderBase struct {
	value *MicrosoftGraphIdentityProviderBase
	isSet bool
}

func (v NullableMicrosoftGraphIdentityProviderBase) Get() *MicrosoftGraphIdentityProviderBase {
	return v.value
}

func (v *NullableMicrosoftGraphIdentityProviderBase) Set(val *MicrosoftGraphIdentityProviderBase) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphIdentityProviderBase) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphIdentityProviderBase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphIdentityProviderBase(val *MicrosoftGraphIdentityProviderBase) *NullableMicrosoftGraphIdentityProviderBase {
	return &NullableMicrosoftGraphIdentityProviderBase{value: val, isSet: true}
}

func (v NullableMicrosoftGraphIdentityProviderBase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphIdentityProviderBase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

