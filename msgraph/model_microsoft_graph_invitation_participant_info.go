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

// MicrosoftGraphInvitationParticipantInfo struct for MicrosoftGraphInvitationParticipantInfo
type MicrosoftGraphInvitationParticipantInfo struct {
	Identity *MicrosoftGraphIdentitySet `json:"identity,omitempty"`
	// Optional. The call which the target identity is currently a part of. This call will be dropped once the participant is added.
	ReplacesCallId NullableString `json:"replacesCallId,omitempty"`
}

// NewMicrosoftGraphInvitationParticipantInfo instantiates a new MicrosoftGraphInvitationParticipantInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphInvitationParticipantInfo() *MicrosoftGraphInvitationParticipantInfo {
	this := MicrosoftGraphInvitationParticipantInfo{}
	return &this
}

// NewMicrosoftGraphInvitationParticipantInfoWithDefaults instantiates a new MicrosoftGraphInvitationParticipantInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphInvitationParticipantInfoWithDefaults() *MicrosoftGraphInvitationParticipantInfo {
	this := MicrosoftGraphInvitationParticipantInfo{}
	return &this
}

// GetIdentity returns the Identity field value if set, zero value otherwise.
func (o *MicrosoftGraphInvitationParticipantInfo) GetIdentity() MicrosoftGraphIdentitySet {
	if o == nil || o.Identity == nil {
		var ret MicrosoftGraphIdentitySet
		return ret
	}
	return *o.Identity
}

// GetIdentityOk returns a tuple with the Identity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphInvitationParticipantInfo) GetIdentityOk() (*MicrosoftGraphIdentitySet, bool) {
	if o == nil || o.Identity == nil {
		return nil, false
	}
	return o.Identity, true
}

// HasIdentity returns a boolean if a field has been set.
func (o *MicrosoftGraphInvitationParticipantInfo) HasIdentity() bool {
	if o != nil && o.Identity != nil {
		return true
	}

	return false
}

// SetIdentity gets a reference to the given MicrosoftGraphIdentitySet and assigns it to the Identity field.
func (o *MicrosoftGraphInvitationParticipantInfo) SetIdentity(v MicrosoftGraphIdentitySet) {
	o.Identity = &v
}

// GetReplacesCallId returns the ReplacesCallId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphInvitationParticipantInfo) GetReplacesCallId() string {
	if o == nil || o.ReplacesCallId.Get() == nil {
		var ret string
		return ret
	}
	return *o.ReplacesCallId.Get()
}

// GetReplacesCallIdOk returns a tuple with the ReplacesCallId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphInvitationParticipantInfo) GetReplacesCallIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ReplacesCallId.Get(), o.ReplacesCallId.IsSet()
}

// HasReplacesCallId returns a boolean if a field has been set.
func (o *MicrosoftGraphInvitationParticipantInfo) HasReplacesCallId() bool {
	if o != nil && o.ReplacesCallId.IsSet() {
		return true
	}

	return false
}

// SetReplacesCallId gets a reference to the given NullableString and assigns it to the ReplacesCallId field.
func (o *MicrosoftGraphInvitationParticipantInfo) SetReplacesCallId(v string) {
	o.ReplacesCallId.Set(&v)
}
// SetReplacesCallIdNil sets the value for ReplacesCallId to be an explicit nil
func (o *MicrosoftGraphInvitationParticipantInfo) SetReplacesCallIdNil() {
	o.ReplacesCallId.Set(nil)
}

// UnsetReplacesCallId ensures that no value is present for ReplacesCallId, not even an explicit nil
func (o *MicrosoftGraphInvitationParticipantInfo) UnsetReplacesCallId() {
	o.ReplacesCallId.Unset()
}

func (o MicrosoftGraphInvitationParticipantInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Identity != nil {
		toSerialize["identity"] = o.Identity
	}
	if o.ReplacesCallId.IsSet() {
		toSerialize["replacesCallId"] = o.ReplacesCallId.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphInvitationParticipantInfo struct {
	value *MicrosoftGraphInvitationParticipantInfo
	isSet bool
}

func (v NullableMicrosoftGraphInvitationParticipantInfo) Get() *MicrosoftGraphInvitationParticipantInfo {
	return v.value
}

func (v *NullableMicrosoftGraphInvitationParticipantInfo) Set(val *MicrosoftGraphInvitationParticipantInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphInvitationParticipantInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphInvitationParticipantInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphInvitationParticipantInfo(val *MicrosoftGraphInvitationParticipantInfo) *NullableMicrosoftGraphInvitationParticipantInfo {
	return &NullableMicrosoftGraphInvitationParticipantInfo{value: val, isSet: true}
}

func (v NullableMicrosoftGraphInvitationParticipantInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphInvitationParticipantInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

