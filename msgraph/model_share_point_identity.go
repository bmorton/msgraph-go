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

// SharePointIdentity struct for SharePointIdentity
type SharePointIdentity struct {
	// The sign in name of the SharePoint identity.
	LoginName NullableString `json:"loginName,omitempty"`
}

// NewSharePointIdentity instantiates a new SharePointIdentity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSharePointIdentity() *SharePointIdentity {
	this := SharePointIdentity{}
	return &this
}

// NewSharePointIdentityWithDefaults instantiates a new SharePointIdentity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSharePointIdentityWithDefaults() *SharePointIdentity {
	this := SharePointIdentity{}
	return &this
}

// GetLoginName returns the LoginName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SharePointIdentity) GetLoginName() string {
	if o == nil || o.LoginName.Get() == nil {
		var ret string
		return ret
	}
	return *o.LoginName.Get()
}

// GetLoginNameOk returns a tuple with the LoginName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SharePointIdentity) GetLoginNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LoginName.Get(), o.LoginName.IsSet()
}

// HasLoginName returns a boolean if a field has been set.
func (o *SharePointIdentity) HasLoginName() bool {
	if o != nil && o.LoginName.IsSet() {
		return true
	}

	return false
}

// SetLoginName gets a reference to the given NullableString and assigns it to the LoginName field.
func (o *SharePointIdentity) SetLoginName(v string) {
	o.LoginName.Set(&v)
}
// SetLoginNameNil sets the value for LoginName to be an explicit nil
func (o *SharePointIdentity) SetLoginNameNil() {
	o.LoginName.Set(nil)
}

// UnsetLoginName ensures that no value is present for LoginName, not even an explicit nil
func (o *SharePointIdentity) UnsetLoginName() {
	o.LoginName.Unset()
}

func (o SharePointIdentity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.LoginName.IsSet() {
		toSerialize["loginName"] = o.LoginName.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableSharePointIdentity struct {
	value *SharePointIdentity
	isSet bool
}

func (v NullableSharePointIdentity) Get() *SharePointIdentity {
	return v.value
}

func (v *NullableSharePointIdentity) Set(val *SharePointIdentity) {
	v.value = val
	v.isSet = true
}

func (v NullableSharePointIdentity) IsSet() bool {
	return v.isSet
}

func (v *NullableSharePointIdentity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSharePointIdentity(val *SharePointIdentity) *NullableSharePointIdentity {
	return &NullableSharePointIdentity{value: val, isSet: true}
}

func (v NullableSharePointIdentity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSharePointIdentity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


