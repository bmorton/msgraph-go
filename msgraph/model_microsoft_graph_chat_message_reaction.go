/*
Partial Graph API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// MicrosoftGraphChatMessageReaction struct for MicrosoftGraphChatMessageReaction
type MicrosoftGraphChatMessageReaction struct {
	// The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// Supported values are like, angry, sad, laugh, heart, surprised.
	ReactionType *string `json:"reactionType,omitempty"`
	User *MicrosoftGraphChatMessageReactionIdentitySet `json:"user,omitempty"`
}

// NewMicrosoftGraphChatMessageReaction instantiates a new MicrosoftGraphChatMessageReaction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphChatMessageReaction() *MicrosoftGraphChatMessageReaction {
	this := MicrosoftGraphChatMessageReaction{}
	return &this
}

// NewMicrosoftGraphChatMessageReactionWithDefaults instantiates a new MicrosoftGraphChatMessageReaction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphChatMessageReactionWithDefaults() *MicrosoftGraphChatMessageReaction {
	this := MicrosoftGraphChatMessageReaction{}
	return &this
}

// GetCreatedDateTime returns the CreatedDateTime field value if set, zero value otherwise.
func (o *MicrosoftGraphChatMessageReaction) GetCreatedDateTime() time.Time {
	if o == nil || o.CreatedDateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedDateTime
}

// GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphChatMessageReaction) GetCreatedDateTimeOk() (*time.Time, bool) {
	if o == nil || o.CreatedDateTime == nil {
		return nil, false
	}
	return o.CreatedDateTime, true
}

// HasCreatedDateTime returns a boolean if a field has been set.
func (o *MicrosoftGraphChatMessageReaction) HasCreatedDateTime() bool {
	if o != nil && o.CreatedDateTime != nil {
		return true
	}

	return false
}

// SetCreatedDateTime gets a reference to the given time.Time and assigns it to the CreatedDateTime field.
func (o *MicrosoftGraphChatMessageReaction) SetCreatedDateTime(v time.Time) {
	o.CreatedDateTime = &v
}

// GetReactionType returns the ReactionType field value if set, zero value otherwise.
func (o *MicrosoftGraphChatMessageReaction) GetReactionType() string {
	if o == nil || o.ReactionType == nil {
		var ret string
		return ret
	}
	return *o.ReactionType
}

// GetReactionTypeOk returns a tuple with the ReactionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphChatMessageReaction) GetReactionTypeOk() (*string, bool) {
	if o == nil || o.ReactionType == nil {
		return nil, false
	}
	return o.ReactionType, true
}

// HasReactionType returns a boolean if a field has been set.
func (o *MicrosoftGraphChatMessageReaction) HasReactionType() bool {
	if o != nil && o.ReactionType != nil {
		return true
	}

	return false
}

// SetReactionType gets a reference to the given string and assigns it to the ReactionType field.
func (o *MicrosoftGraphChatMessageReaction) SetReactionType(v string) {
	o.ReactionType = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *MicrosoftGraphChatMessageReaction) GetUser() MicrosoftGraphChatMessageReactionIdentitySet {
	if o == nil || o.User == nil {
		var ret MicrosoftGraphChatMessageReactionIdentitySet
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphChatMessageReaction) GetUserOk() (*MicrosoftGraphChatMessageReactionIdentitySet, bool) {
	if o == nil || o.User == nil {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *MicrosoftGraphChatMessageReaction) HasUser() bool {
	if o != nil && o.User != nil {
		return true
	}

	return false
}

// SetUser gets a reference to the given MicrosoftGraphChatMessageReactionIdentitySet and assigns it to the User field.
func (o *MicrosoftGraphChatMessageReaction) SetUser(v MicrosoftGraphChatMessageReactionIdentitySet) {
	o.User = &v
}

func (o MicrosoftGraphChatMessageReaction) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CreatedDateTime != nil {
		toSerialize["createdDateTime"] = o.CreatedDateTime
	}
	if o.ReactionType != nil {
		toSerialize["reactionType"] = o.ReactionType
	}
	if o.User != nil {
		toSerialize["user"] = o.User
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphChatMessageReaction struct {
	value *MicrosoftGraphChatMessageReaction
	isSet bool
}

func (v NullableMicrosoftGraphChatMessageReaction) Get() *MicrosoftGraphChatMessageReaction {
	return v.value
}

func (v *NullableMicrosoftGraphChatMessageReaction) Set(val *MicrosoftGraphChatMessageReaction) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphChatMessageReaction) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphChatMessageReaction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphChatMessageReaction(val *MicrosoftGraphChatMessageReaction) *NullableMicrosoftGraphChatMessageReaction {
	return &NullableMicrosoftGraphChatMessageReaction{value: val, isSet: true}
}

func (v NullableMicrosoftGraphChatMessageReaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphChatMessageReaction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

