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

// MicrosoftGraphUserScopeTeamsAppInstallation struct for MicrosoftGraphUserScopeTeamsAppInstallation
type MicrosoftGraphUserScopeTeamsAppInstallation struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
	// The app that is installed.
	TeamsApp AnyOfmicrosoftGraphTeamsApp `json:"teamsApp,omitempty"`
	// The details of this version of the app.
	TeamsAppDefinition AnyOfmicrosoftGraphTeamsAppDefinition `json:"teamsAppDefinition,omitempty"`
	// The chat between the user and Teams app.
	Chat AnyOfmicrosoftGraphChat `json:"chat,omitempty"`
}

// NewMicrosoftGraphUserScopeTeamsAppInstallation instantiates a new MicrosoftGraphUserScopeTeamsAppInstallation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphUserScopeTeamsAppInstallation() *MicrosoftGraphUserScopeTeamsAppInstallation {
	this := MicrosoftGraphUserScopeTeamsAppInstallation{}
	return &this
}

// NewMicrosoftGraphUserScopeTeamsAppInstallationWithDefaults instantiates a new MicrosoftGraphUserScopeTeamsAppInstallation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphUserScopeTeamsAppInstallationWithDefaults() *MicrosoftGraphUserScopeTeamsAppInstallation {
	this := MicrosoftGraphUserScopeTeamsAppInstallation{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MicrosoftGraphUserScopeTeamsAppInstallation) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphUserScopeTeamsAppInstallation) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphUserScopeTeamsAppInstallation) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MicrosoftGraphUserScopeTeamsAppInstallation) SetId(v string) {
	o.Id = &v
}

// GetTeamsApp returns the TeamsApp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphUserScopeTeamsAppInstallation) GetTeamsApp() AnyOfmicrosoftGraphTeamsApp {
	if o == nil  {
		var ret AnyOfmicrosoftGraphTeamsApp
		return ret
	}
	return o.TeamsApp
}

// GetTeamsAppOk returns a tuple with the TeamsApp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphUserScopeTeamsAppInstallation) GetTeamsAppOk() (*AnyOfmicrosoftGraphTeamsApp, bool) {
	if o == nil || o.TeamsApp == nil {
		return nil, false
	}
	return &o.TeamsApp, true
}

// HasTeamsApp returns a boolean if a field has been set.
func (o *MicrosoftGraphUserScopeTeamsAppInstallation) HasTeamsApp() bool {
	if o != nil && o.TeamsApp != nil {
		return true
	}

	return false
}

// SetTeamsApp gets a reference to the given AnyOfmicrosoftGraphTeamsApp and assigns it to the TeamsApp field.
func (o *MicrosoftGraphUserScopeTeamsAppInstallation) SetTeamsApp(v AnyOfmicrosoftGraphTeamsApp) {
	o.TeamsApp = v
}

// GetTeamsAppDefinition returns the TeamsAppDefinition field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphUserScopeTeamsAppInstallation) GetTeamsAppDefinition() AnyOfmicrosoftGraphTeamsAppDefinition {
	if o == nil  {
		var ret AnyOfmicrosoftGraphTeamsAppDefinition
		return ret
	}
	return o.TeamsAppDefinition
}

// GetTeamsAppDefinitionOk returns a tuple with the TeamsAppDefinition field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphUserScopeTeamsAppInstallation) GetTeamsAppDefinitionOk() (*AnyOfmicrosoftGraphTeamsAppDefinition, bool) {
	if o == nil || o.TeamsAppDefinition == nil {
		return nil, false
	}
	return &o.TeamsAppDefinition, true
}

// HasTeamsAppDefinition returns a boolean if a field has been set.
func (o *MicrosoftGraphUserScopeTeamsAppInstallation) HasTeamsAppDefinition() bool {
	if o != nil && o.TeamsAppDefinition != nil {
		return true
	}

	return false
}

// SetTeamsAppDefinition gets a reference to the given AnyOfmicrosoftGraphTeamsAppDefinition and assigns it to the TeamsAppDefinition field.
func (o *MicrosoftGraphUserScopeTeamsAppInstallation) SetTeamsAppDefinition(v AnyOfmicrosoftGraphTeamsAppDefinition) {
	o.TeamsAppDefinition = v
}

// GetChat returns the Chat field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphUserScopeTeamsAppInstallation) GetChat() AnyOfmicrosoftGraphChat {
	if o == nil  {
		var ret AnyOfmicrosoftGraphChat
		return ret
	}
	return o.Chat
}

// GetChatOk returns a tuple with the Chat field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphUserScopeTeamsAppInstallation) GetChatOk() (*AnyOfmicrosoftGraphChat, bool) {
	if o == nil || o.Chat == nil {
		return nil, false
	}
	return &o.Chat, true
}

// HasChat returns a boolean if a field has been set.
func (o *MicrosoftGraphUserScopeTeamsAppInstallation) HasChat() bool {
	if o != nil && o.Chat != nil {
		return true
	}

	return false
}

// SetChat gets a reference to the given AnyOfmicrosoftGraphChat and assigns it to the Chat field.
func (o *MicrosoftGraphUserScopeTeamsAppInstallation) SetChat(v AnyOfmicrosoftGraphChat) {
	o.Chat = v
}

func (o MicrosoftGraphUserScopeTeamsAppInstallation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.TeamsApp != nil {
		toSerialize["teamsApp"] = o.TeamsApp
	}
	if o.TeamsAppDefinition != nil {
		toSerialize["teamsAppDefinition"] = o.TeamsAppDefinition
	}
	if o.Chat != nil {
		toSerialize["chat"] = o.Chat
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphUserScopeTeamsAppInstallation struct {
	value *MicrosoftGraphUserScopeTeamsAppInstallation
	isSet bool
}

func (v NullableMicrosoftGraphUserScopeTeamsAppInstallation) Get() *MicrosoftGraphUserScopeTeamsAppInstallation {
	return v.value
}

func (v *NullableMicrosoftGraphUserScopeTeamsAppInstallation) Set(val *MicrosoftGraphUserScopeTeamsAppInstallation) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphUserScopeTeamsAppInstallation) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphUserScopeTeamsAppInstallation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphUserScopeTeamsAppInstallation(val *MicrosoftGraphUserScopeTeamsAppInstallation) *NullableMicrosoftGraphUserScopeTeamsAppInstallation {
	return &NullableMicrosoftGraphUserScopeTeamsAppInstallation{value: val, isSet: true}
}

func (v NullableMicrosoftGraphUserScopeTeamsAppInstallation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphUserScopeTeamsAppInstallation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


