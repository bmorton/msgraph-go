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

// InlineObject1195 struct for InlineObject1195
type InlineObject1195 struct {
	EntityType NullableString `json:"entityType,omitempty"`
	DisplayName NullableString `json:"displayName,omitempty"`
	MailNickname NullableString `json:"mailNickname,omitempty"`
	OnBehalfOfUserId NullableString `json:"onBehalfOfUserId,omitempty"`
}

// NewInlineObject1195 instantiates a new InlineObject1195 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject1195() *InlineObject1195 {
	this := InlineObject1195{}
	return &this
}

// NewInlineObject1195WithDefaults instantiates a new InlineObject1195 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject1195WithDefaults() *InlineObject1195 {
	this := InlineObject1195{}
	return &this
}

// GetEntityType returns the EntityType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1195) GetEntityType() string {
	if o == nil || o.EntityType.Get() == nil {
		var ret string
		return ret
	}
	return *o.EntityType.Get()
}

// GetEntityTypeOk returns a tuple with the EntityType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1195) GetEntityTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.EntityType.Get(), o.EntityType.IsSet()
}

// HasEntityType returns a boolean if a field has been set.
func (o *InlineObject1195) HasEntityType() bool {
	if o != nil && o.EntityType.IsSet() {
		return true
	}

	return false
}

// SetEntityType gets a reference to the given NullableString and assigns it to the EntityType field.
func (o *InlineObject1195) SetEntityType(v string) {
	o.EntityType.Set(&v)
}
// SetEntityTypeNil sets the value for EntityType to be an explicit nil
func (o *InlineObject1195) SetEntityTypeNil() {
	o.EntityType.Set(nil)
}

// UnsetEntityType ensures that no value is present for EntityType, not even an explicit nil
func (o *InlineObject1195) UnsetEntityType() {
	o.EntityType.Unset()
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1195) GetDisplayName() string {
	if o == nil || o.DisplayName.Get() == nil {
		var ret string
		return ret
	}
	return *o.DisplayName.Get()
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1195) GetDisplayNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DisplayName.Get(), o.DisplayName.IsSet()
}

// HasDisplayName returns a boolean if a field has been set.
func (o *InlineObject1195) HasDisplayName() bool {
	if o != nil && o.DisplayName.IsSet() {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given NullableString and assigns it to the DisplayName field.
func (o *InlineObject1195) SetDisplayName(v string) {
	o.DisplayName.Set(&v)
}
// SetDisplayNameNil sets the value for DisplayName to be an explicit nil
func (o *InlineObject1195) SetDisplayNameNil() {
	o.DisplayName.Set(nil)
}

// UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
func (o *InlineObject1195) UnsetDisplayName() {
	o.DisplayName.Unset()
}

// GetMailNickname returns the MailNickname field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1195) GetMailNickname() string {
	if o == nil || o.MailNickname.Get() == nil {
		var ret string
		return ret
	}
	return *o.MailNickname.Get()
}

// GetMailNicknameOk returns a tuple with the MailNickname field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1195) GetMailNicknameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.MailNickname.Get(), o.MailNickname.IsSet()
}

// HasMailNickname returns a boolean if a field has been set.
func (o *InlineObject1195) HasMailNickname() bool {
	if o != nil && o.MailNickname.IsSet() {
		return true
	}

	return false
}

// SetMailNickname gets a reference to the given NullableString and assigns it to the MailNickname field.
func (o *InlineObject1195) SetMailNickname(v string) {
	o.MailNickname.Set(&v)
}
// SetMailNicknameNil sets the value for MailNickname to be an explicit nil
func (o *InlineObject1195) SetMailNicknameNil() {
	o.MailNickname.Set(nil)
}

// UnsetMailNickname ensures that no value is present for MailNickname, not even an explicit nil
func (o *InlineObject1195) UnsetMailNickname() {
	o.MailNickname.Unset()
}

// GetOnBehalfOfUserId returns the OnBehalfOfUserId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1195) GetOnBehalfOfUserId() string {
	if o == nil || o.OnBehalfOfUserId.Get() == nil {
		var ret string
		return ret
	}
	return *o.OnBehalfOfUserId.Get()
}

// GetOnBehalfOfUserIdOk returns a tuple with the OnBehalfOfUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1195) GetOnBehalfOfUserIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.OnBehalfOfUserId.Get(), o.OnBehalfOfUserId.IsSet()
}

// HasOnBehalfOfUserId returns a boolean if a field has been set.
func (o *InlineObject1195) HasOnBehalfOfUserId() bool {
	if o != nil && o.OnBehalfOfUserId.IsSet() {
		return true
	}

	return false
}

// SetOnBehalfOfUserId gets a reference to the given NullableString and assigns it to the OnBehalfOfUserId field.
func (o *InlineObject1195) SetOnBehalfOfUserId(v string) {
	o.OnBehalfOfUserId.Set(&v)
}
// SetOnBehalfOfUserIdNil sets the value for OnBehalfOfUserId to be an explicit nil
func (o *InlineObject1195) SetOnBehalfOfUserIdNil() {
	o.OnBehalfOfUserId.Set(nil)
}

// UnsetOnBehalfOfUserId ensures that no value is present for OnBehalfOfUserId, not even an explicit nil
func (o *InlineObject1195) UnsetOnBehalfOfUserId() {
	o.OnBehalfOfUserId.Unset()
}

func (o InlineObject1195) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.EntityType.IsSet() {
		toSerialize["entityType"] = o.EntityType.Get()
	}
	if o.DisplayName.IsSet() {
		toSerialize["displayName"] = o.DisplayName.Get()
	}
	if o.MailNickname.IsSet() {
		toSerialize["mailNickname"] = o.MailNickname.Get()
	}
	if o.OnBehalfOfUserId.IsSet() {
		toSerialize["onBehalfOfUserId"] = o.OnBehalfOfUserId.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject1195 struct {
	value *InlineObject1195
	isSet bool
}

func (v NullableInlineObject1195) Get() *InlineObject1195 {
	return v.value
}

func (v *NullableInlineObject1195) Set(val *InlineObject1195) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject1195) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject1195) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject1195(val *InlineObject1195) *NullableInlineObject1195 {
	return &NullableInlineObject1195{value: val, isSet: true}
}

func (v NullableInlineObject1195) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject1195) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


