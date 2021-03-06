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

// MicrosoftGraphSharePointIdentity struct for MicrosoftGraphSharePointIdentity
type MicrosoftGraphSharePointIdentity struct {
	// The identity's display name. Note that this may not always be available or up to date. For example, if a user changes their display name, the API may show the new value in a future response, but the items associated with the user won't show up as having changed when using delta.
	DisplayName NullableString `json:"displayName,omitempty"`
	// Unique identifier for the identity.
	Id NullableString `json:"id,omitempty"`
	// The sign in name of the SharePoint identity.
	LoginName NullableString `json:"loginName,omitempty"`
}

// NewMicrosoftGraphSharePointIdentity instantiates a new MicrosoftGraphSharePointIdentity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphSharePointIdentity() *MicrosoftGraphSharePointIdentity {
	this := MicrosoftGraphSharePointIdentity{}
	return &this
}

// NewMicrosoftGraphSharePointIdentityWithDefaults instantiates a new MicrosoftGraphSharePointIdentity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphSharePointIdentityWithDefaults() *MicrosoftGraphSharePointIdentity {
	this := MicrosoftGraphSharePointIdentity{}
	return &this
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphSharePointIdentity) GetDisplayName() string {
	if o == nil || o.DisplayName.Get() == nil {
		var ret string
		return ret
	}
	return *o.DisplayName.Get()
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphSharePointIdentity) GetDisplayNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DisplayName.Get(), o.DisplayName.IsSet()
}

// HasDisplayName returns a boolean if a field has been set.
func (o *MicrosoftGraphSharePointIdentity) HasDisplayName() bool {
	if o != nil && o.DisplayName.IsSet() {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given NullableString and assigns it to the DisplayName field.
func (o *MicrosoftGraphSharePointIdentity) SetDisplayName(v string) {
	o.DisplayName.Set(&v)
}
// SetDisplayNameNil sets the value for DisplayName to be an explicit nil
func (o *MicrosoftGraphSharePointIdentity) SetDisplayNameNil() {
	o.DisplayName.Set(nil)
}

// UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
func (o *MicrosoftGraphSharePointIdentity) UnsetDisplayName() {
	o.DisplayName.Unset()
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphSharePointIdentity) GetId() string {
	if o == nil || o.Id.Get() == nil {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphSharePointIdentity) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphSharePointIdentity) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *MicrosoftGraphSharePointIdentity) SetId(v string) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *MicrosoftGraphSharePointIdentity) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *MicrosoftGraphSharePointIdentity) UnsetId() {
	o.Id.Unset()
}

// GetLoginName returns the LoginName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphSharePointIdentity) GetLoginName() string {
	if o == nil || o.LoginName.Get() == nil {
		var ret string
		return ret
	}
	return *o.LoginName.Get()
}

// GetLoginNameOk returns a tuple with the LoginName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphSharePointIdentity) GetLoginNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LoginName.Get(), o.LoginName.IsSet()
}

// HasLoginName returns a boolean if a field has been set.
func (o *MicrosoftGraphSharePointIdentity) HasLoginName() bool {
	if o != nil && o.LoginName.IsSet() {
		return true
	}

	return false
}

// SetLoginName gets a reference to the given NullableString and assigns it to the LoginName field.
func (o *MicrosoftGraphSharePointIdentity) SetLoginName(v string) {
	o.LoginName.Set(&v)
}
// SetLoginNameNil sets the value for LoginName to be an explicit nil
func (o *MicrosoftGraphSharePointIdentity) SetLoginNameNil() {
	o.LoginName.Set(nil)
}

// UnsetLoginName ensures that no value is present for LoginName, not even an explicit nil
func (o *MicrosoftGraphSharePointIdentity) UnsetLoginName() {
	o.LoginName.Unset()
}

func (o MicrosoftGraphSharePointIdentity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DisplayName.IsSet() {
		toSerialize["displayName"] = o.DisplayName.Get()
	}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.LoginName.IsSet() {
		toSerialize["loginName"] = o.LoginName.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphSharePointIdentity struct {
	value *MicrosoftGraphSharePointIdentity
	isSet bool
}

func (v NullableMicrosoftGraphSharePointIdentity) Get() *MicrosoftGraphSharePointIdentity {
	return v.value
}

func (v *NullableMicrosoftGraphSharePointIdentity) Set(val *MicrosoftGraphSharePointIdentity) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphSharePointIdentity) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphSharePointIdentity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphSharePointIdentity(val *MicrosoftGraphSharePointIdentity) *NullableMicrosoftGraphSharePointIdentity {
	return &NullableMicrosoftGraphSharePointIdentity{value: val, isSet: true}
}

func (v NullableMicrosoftGraphSharePointIdentity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphSharePointIdentity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


