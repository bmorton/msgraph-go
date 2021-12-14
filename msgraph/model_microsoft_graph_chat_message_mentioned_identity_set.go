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

// MicrosoftGraphChatMessageMentionedIdentitySet struct for MicrosoftGraphChatMessageMentionedIdentitySet
type MicrosoftGraphChatMessageMentionedIdentitySet struct {
	// Optional. The application associated with this action.
	Application AnyOfmicrosoftGraphIdentity `json:"application,omitempty"`
	// Optional. The device associated with this action.
	Device AnyOfmicrosoftGraphIdentity `json:"device,omitempty"`
	// Optional. The user associated with this action.
	User AnyOfmicrosoftGraphIdentity `json:"user,omitempty"`
	// If present, represents a conversation (for example, team or channel) @mentioned in a message.
	Conversation AnyOfmicrosoftGraphTeamworkConversationIdentity `json:"conversation,omitempty"`
}

// NewMicrosoftGraphChatMessageMentionedIdentitySet instantiates a new MicrosoftGraphChatMessageMentionedIdentitySet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphChatMessageMentionedIdentitySet() *MicrosoftGraphChatMessageMentionedIdentitySet {
	this := MicrosoftGraphChatMessageMentionedIdentitySet{}
	return &this
}

// NewMicrosoftGraphChatMessageMentionedIdentitySetWithDefaults instantiates a new MicrosoftGraphChatMessageMentionedIdentitySet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphChatMessageMentionedIdentitySetWithDefaults() *MicrosoftGraphChatMessageMentionedIdentitySet {
	this := MicrosoftGraphChatMessageMentionedIdentitySet{}
	return &this
}

// GetApplication returns the Application field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphChatMessageMentionedIdentitySet) GetApplication() AnyOfmicrosoftGraphIdentity {
	if o == nil  {
		var ret AnyOfmicrosoftGraphIdentity
		return ret
	}
	return o.Application
}

// GetApplicationOk returns a tuple with the Application field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphChatMessageMentionedIdentitySet) GetApplicationOk() (*AnyOfmicrosoftGraphIdentity, bool) {
	if o == nil || o.Application == nil {
		return nil, false
	}
	return &o.Application, true
}

// HasApplication returns a boolean if a field has been set.
func (o *MicrosoftGraphChatMessageMentionedIdentitySet) HasApplication() bool {
	if o != nil && o.Application != nil {
		return true
	}

	return false
}

// SetApplication gets a reference to the given AnyOfmicrosoftGraphIdentity and assigns it to the Application field.
func (o *MicrosoftGraphChatMessageMentionedIdentitySet) SetApplication(v AnyOfmicrosoftGraphIdentity) {
	o.Application = v
}

// GetDevice returns the Device field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphChatMessageMentionedIdentitySet) GetDevice() AnyOfmicrosoftGraphIdentity {
	if o == nil  {
		var ret AnyOfmicrosoftGraphIdentity
		return ret
	}
	return o.Device
}

// GetDeviceOk returns a tuple with the Device field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphChatMessageMentionedIdentitySet) GetDeviceOk() (*AnyOfmicrosoftGraphIdentity, bool) {
	if o == nil || o.Device == nil {
		return nil, false
	}
	return &o.Device, true
}

// HasDevice returns a boolean if a field has been set.
func (o *MicrosoftGraphChatMessageMentionedIdentitySet) HasDevice() bool {
	if o != nil && o.Device != nil {
		return true
	}

	return false
}

// SetDevice gets a reference to the given AnyOfmicrosoftGraphIdentity and assigns it to the Device field.
func (o *MicrosoftGraphChatMessageMentionedIdentitySet) SetDevice(v AnyOfmicrosoftGraphIdentity) {
	o.Device = v
}

// GetUser returns the User field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphChatMessageMentionedIdentitySet) GetUser() AnyOfmicrosoftGraphIdentity {
	if o == nil  {
		var ret AnyOfmicrosoftGraphIdentity
		return ret
	}
	return o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphChatMessageMentionedIdentitySet) GetUserOk() (*AnyOfmicrosoftGraphIdentity, bool) {
	if o == nil || o.User == nil {
		return nil, false
	}
	return &o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *MicrosoftGraphChatMessageMentionedIdentitySet) HasUser() bool {
	if o != nil && o.User != nil {
		return true
	}

	return false
}

// SetUser gets a reference to the given AnyOfmicrosoftGraphIdentity and assigns it to the User field.
func (o *MicrosoftGraphChatMessageMentionedIdentitySet) SetUser(v AnyOfmicrosoftGraphIdentity) {
	o.User = v
}

// GetConversation returns the Conversation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphChatMessageMentionedIdentitySet) GetConversation() AnyOfmicrosoftGraphTeamworkConversationIdentity {
	if o == nil  {
		var ret AnyOfmicrosoftGraphTeamworkConversationIdentity
		return ret
	}
	return o.Conversation
}

// GetConversationOk returns a tuple with the Conversation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphChatMessageMentionedIdentitySet) GetConversationOk() (*AnyOfmicrosoftGraphTeamworkConversationIdentity, bool) {
	if o == nil || o.Conversation == nil {
		return nil, false
	}
	return &o.Conversation, true
}

// HasConversation returns a boolean if a field has been set.
func (o *MicrosoftGraphChatMessageMentionedIdentitySet) HasConversation() bool {
	if o != nil && o.Conversation != nil {
		return true
	}

	return false
}

// SetConversation gets a reference to the given AnyOfmicrosoftGraphTeamworkConversationIdentity and assigns it to the Conversation field.
func (o *MicrosoftGraphChatMessageMentionedIdentitySet) SetConversation(v AnyOfmicrosoftGraphTeamworkConversationIdentity) {
	o.Conversation = v
}

func (o MicrosoftGraphChatMessageMentionedIdentitySet) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Application != nil {
		toSerialize["application"] = o.Application
	}
	if o.Device != nil {
		toSerialize["device"] = o.Device
	}
	if o.User != nil {
		toSerialize["user"] = o.User
	}
	if o.Conversation != nil {
		toSerialize["conversation"] = o.Conversation
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphChatMessageMentionedIdentitySet struct {
	value *MicrosoftGraphChatMessageMentionedIdentitySet
	isSet bool
}

func (v NullableMicrosoftGraphChatMessageMentionedIdentitySet) Get() *MicrosoftGraphChatMessageMentionedIdentitySet {
	return v.value
}

func (v *NullableMicrosoftGraphChatMessageMentionedIdentitySet) Set(val *MicrosoftGraphChatMessageMentionedIdentitySet) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphChatMessageMentionedIdentitySet) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphChatMessageMentionedIdentitySet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphChatMessageMentionedIdentitySet(val *MicrosoftGraphChatMessageMentionedIdentitySet) *NullableMicrosoftGraphChatMessageMentionedIdentitySet {
	return &NullableMicrosoftGraphChatMessageMentionedIdentitySet{value: val, isSet: true}
}

func (v NullableMicrosoftGraphChatMessageMentionedIdentitySet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphChatMessageMentionedIdentitySet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

