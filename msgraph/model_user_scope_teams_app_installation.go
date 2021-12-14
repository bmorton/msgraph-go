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

// UserScopeTeamsAppInstallation struct for UserScopeTeamsAppInstallation
type UserScopeTeamsAppInstallation struct {
	// The chat between the user and Teams app.
	Chat AnyOfmicrosoftGraphChat `json:"chat,omitempty"`
}

// NewUserScopeTeamsAppInstallation instantiates a new UserScopeTeamsAppInstallation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserScopeTeamsAppInstallation() *UserScopeTeamsAppInstallation {
	this := UserScopeTeamsAppInstallation{}
	return &this
}

// NewUserScopeTeamsAppInstallationWithDefaults instantiates a new UserScopeTeamsAppInstallation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserScopeTeamsAppInstallationWithDefaults() *UserScopeTeamsAppInstallation {
	this := UserScopeTeamsAppInstallation{}
	return &this
}

// GetChat returns the Chat field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserScopeTeamsAppInstallation) GetChat() AnyOfmicrosoftGraphChat {
	if o == nil  {
		var ret AnyOfmicrosoftGraphChat
		return ret
	}
	return o.Chat
}

// GetChatOk returns a tuple with the Chat field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserScopeTeamsAppInstallation) GetChatOk() (*AnyOfmicrosoftGraphChat, bool) {
	if o == nil || o.Chat == nil {
		return nil, false
	}
	return &o.Chat, true
}

// HasChat returns a boolean if a field has been set.
func (o *UserScopeTeamsAppInstallation) HasChat() bool {
	if o != nil && o.Chat != nil {
		return true
	}

	return false
}

// SetChat gets a reference to the given AnyOfmicrosoftGraphChat and assigns it to the Chat field.
func (o *UserScopeTeamsAppInstallation) SetChat(v AnyOfmicrosoftGraphChat) {
	o.Chat = v
}

func (o UserScopeTeamsAppInstallation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Chat != nil {
		toSerialize["chat"] = o.Chat
	}
	return json.Marshal(toSerialize)
}

type NullableUserScopeTeamsAppInstallation struct {
	value *UserScopeTeamsAppInstallation
	isSet bool
}

func (v NullableUserScopeTeamsAppInstallation) Get() *UserScopeTeamsAppInstallation {
	return v.value
}

func (v *NullableUserScopeTeamsAppInstallation) Set(val *UserScopeTeamsAppInstallation) {
	v.value = val
	v.isSet = true
}

func (v NullableUserScopeTeamsAppInstallation) IsSet() bool {
	return v.isSet
}

func (v *NullableUserScopeTeamsAppInstallation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserScopeTeamsAppInstallation(val *UserScopeTeamsAppInstallation) *NullableUserScopeTeamsAppInstallation {
	return &NullableUserScopeTeamsAppInstallation{value: val, isSet: true}
}

func (v NullableUserScopeTeamsAppInstallation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserScopeTeamsAppInstallation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


