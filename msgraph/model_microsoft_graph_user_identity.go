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

// MicrosoftGraphUserIdentity struct for MicrosoftGraphUserIdentity
type MicrosoftGraphUserIdentity struct {
	// The identity's display name. Note that this may not always be available or up to date. For example, if a user changes their display name, the API may show the new value in a future response, but the items associated with the user won't show up as having changed when using delta.
	DisplayName NullableString `json:"displayName,omitempty"`
	// Unique identifier for the identity.
	Id NullableString `json:"id,omitempty"`
	// Indicates the client IP address used by user performing the activity (audit log only).
	IpAddress NullableString `json:"ipAddress,omitempty"`
	// The userPrincipalName attribute of the user.
	UserPrincipalName NullableString `json:"userPrincipalName,omitempty"`
}

// NewMicrosoftGraphUserIdentity instantiates a new MicrosoftGraphUserIdentity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphUserIdentity() *MicrosoftGraphUserIdentity {
	this := MicrosoftGraphUserIdentity{}
	return &this
}

// NewMicrosoftGraphUserIdentityWithDefaults instantiates a new MicrosoftGraphUserIdentity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphUserIdentityWithDefaults() *MicrosoftGraphUserIdentity {
	this := MicrosoftGraphUserIdentity{}
	return &this
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphUserIdentity) GetDisplayName() string {
	if o == nil || o.DisplayName.Get() == nil {
		var ret string
		return ret
	}
	return *o.DisplayName.Get()
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphUserIdentity) GetDisplayNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DisplayName.Get(), o.DisplayName.IsSet()
}

// HasDisplayName returns a boolean if a field has been set.
func (o *MicrosoftGraphUserIdentity) HasDisplayName() bool {
	if o != nil && o.DisplayName.IsSet() {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given NullableString and assigns it to the DisplayName field.
func (o *MicrosoftGraphUserIdentity) SetDisplayName(v string) {
	o.DisplayName.Set(&v)
}
// SetDisplayNameNil sets the value for DisplayName to be an explicit nil
func (o *MicrosoftGraphUserIdentity) SetDisplayNameNil() {
	o.DisplayName.Set(nil)
}

// UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
func (o *MicrosoftGraphUserIdentity) UnsetDisplayName() {
	o.DisplayName.Unset()
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphUserIdentity) GetId() string {
	if o == nil || o.Id.Get() == nil {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphUserIdentity) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphUserIdentity) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *MicrosoftGraphUserIdentity) SetId(v string) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *MicrosoftGraphUserIdentity) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *MicrosoftGraphUserIdentity) UnsetId() {
	o.Id.Unset()
}

// GetIpAddress returns the IpAddress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphUserIdentity) GetIpAddress() string {
	if o == nil || o.IpAddress.Get() == nil {
		var ret string
		return ret
	}
	return *o.IpAddress.Get()
}

// GetIpAddressOk returns a tuple with the IpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphUserIdentity) GetIpAddressOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IpAddress.Get(), o.IpAddress.IsSet()
}

// HasIpAddress returns a boolean if a field has been set.
func (o *MicrosoftGraphUserIdentity) HasIpAddress() bool {
	if o != nil && o.IpAddress.IsSet() {
		return true
	}

	return false
}

// SetIpAddress gets a reference to the given NullableString and assigns it to the IpAddress field.
func (o *MicrosoftGraphUserIdentity) SetIpAddress(v string) {
	o.IpAddress.Set(&v)
}
// SetIpAddressNil sets the value for IpAddress to be an explicit nil
func (o *MicrosoftGraphUserIdentity) SetIpAddressNil() {
	o.IpAddress.Set(nil)
}

// UnsetIpAddress ensures that no value is present for IpAddress, not even an explicit nil
func (o *MicrosoftGraphUserIdentity) UnsetIpAddress() {
	o.IpAddress.Unset()
}

// GetUserPrincipalName returns the UserPrincipalName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphUserIdentity) GetUserPrincipalName() string {
	if o == nil || o.UserPrincipalName.Get() == nil {
		var ret string
		return ret
	}
	return *o.UserPrincipalName.Get()
}

// GetUserPrincipalNameOk returns a tuple with the UserPrincipalName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphUserIdentity) GetUserPrincipalNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.UserPrincipalName.Get(), o.UserPrincipalName.IsSet()
}

// HasUserPrincipalName returns a boolean if a field has been set.
func (o *MicrosoftGraphUserIdentity) HasUserPrincipalName() bool {
	if o != nil && o.UserPrincipalName.IsSet() {
		return true
	}

	return false
}

// SetUserPrincipalName gets a reference to the given NullableString and assigns it to the UserPrincipalName field.
func (o *MicrosoftGraphUserIdentity) SetUserPrincipalName(v string) {
	o.UserPrincipalName.Set(&v)
}
// SetUserPrincipalNameNil sets the value for UserPrincipalName to be an explicit nil
func (o *MicrosoftGraphUserIdentity) SetUserPrincipalNameNil() {
	o.UserPrincipalName.Set(nil)
}

// UnsetUserPrincipalName ensures that no value is present for UserPrincipalName, not even an explicit nil
func (o *MicrosoftGraphUserIdentity) UnsetUserPrincipalName() {
	o.UserPrincipalName.Unset()
}

func (o MicrosoftGraphUserIdentity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DisplayName.IsSet() {
		toSerialize["displayName"] = o.DisplayName.Get()
	}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.IpAddress.IsSet() {
		toSerialize["ipAddress"] = o.IpAddress.Get()
	}
	if o.UserPrincipalName.IsSet() {
		toSerialize["userPrincipalName"] = o.UserPrincipalName.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphUserIdentity struct {
	value *MicrosoftGraphUserIdentity
	isSet bool
}

func (v NullableMicrosoftGraphUserIdentity) Get() *MicrosoftGraphUserIdentity {
	return v.value
}

func (v *NullableMicrosoftGraphUserIdentity) Set(val *MicrosoftGraphUserIdentity) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphUserIdentity) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphUserIdentity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphUserIdentity(val *MicrosoftGraphUserIdentity) *NullableMicrosoftGraphUserIdentity {
	return &NullableMicrosoftGraphUserIdentity{value: val, isSet: true}
}

func (v NullableMicrosoftGraphUserIdentity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphUserIdentity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


