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

// InlineObject241 struct for InlineObject241
type InlineObject241 struct {
	DisplayName NullableString `json:"displayName,omitempty"`
	MailNickname NullableString `json:"mailNickname,omitempty"`
	OnBehalfOfUserId NullableString `json:"onBehalfOfUserId,omitempty"`
}

// NewInlineObject241 instantiates a new InlineObject241 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject241() *InlineObject241 {
	this := InlineObject241{}
	return &this
}

// NewInlineObject241WithDefaults instantiates a new InlineObject241 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject241WithDefaults() *InlineObject241 {
	this := InlineObject241{}
	return &this
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject241) GetDisplayName() string {
	if o == nil || o.DisplayName.Get() == nil {
		var ret string
		return ret
	}
	return *o.DisplayName.Get()
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject241) GetDisplayNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DisplayName.Get(), o.DisplayName.IsSet()
}

// HasDisplayName returns a boolean if a field has been set.
func (o *InlineObject241) HasDisplayName() bool {
	if o != nil && o.DisplayName.IsSet() {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given NullableString and assigns it to the DisplayName field.
func (o *InlineObject241) SetDisplayName(v string) {
	o.DisplayName.Set(&v)
}
// SetDisplayNameNil sets the value for DisplayName to be an explicit nil
func (o *InlineObject241) SetDisplayNameNil() {
	o.DisplayName.Set(nil)
}

// UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
func (o *InlineObject241) UnsetDisplayName() {
	o.DisplayName.Unset()
}

// GetMailNickname returns the MailNickname field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject241) GetMailNickname() string {
	if o == nil || o.MailNickname.Get() == nil {
		var ret string
		return ret
	}
	return *o.MailNickname.Get()
}

// GetMailNicknameOk returns a tuple with the MailNickname field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject241) GetMailNicknameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.MailNickname.Get(), o.MailNickname.IsSet()
}

// HasMailNickname returns a boolean if a field has been set.
func (o *InlineObject241) HasMailNickname() bool {
	if o != nil && o.MailNickname.IsSet() {
		return true
	}

	return false
}

// SetMailNickname gets a reference to the given NullableString and assigns it to the MailNickname field.
func (o *InlineObject241) SetMailNickname(v string) {
	o.MailNickname.Set(&v)
}
// SetMailNicknameNil sets the value for MailNickname to be an explicit nil
func (o *InlineObject241) SetMailNicknameNil() {
	o.MailNickname.Set(nil)
}

// UnsetMailNickname ensures that no value is present for MailNickname, not even an explicit nil
func (o *InlineObject241) UnsetMailNickname() {
	o.MailNickname.Unset()
}

// GetOnBehalfOfUserId returns the OnBehalfOfUserId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject241) GetOnBehalfOfUserId() string {
	if o == nil || o.OnBehalfOfUserId.Get() == nil {
		var ret string
		return ret
	}
	return *o.OnBehalfOfUserId.Get()
}

// GetOnBehalfOfUserIdOk returns a tuple with the OnBehalfOfUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject241) GetOnBehalfOfUserIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.OnBehalfOfUserId.Get(), o.OnBehalfOfUserId.IsSet()
}

// HasOnBehalfOfUserId returns a boolean if a field has been set.
func (o *InlineObject241) HasOnBehalfOfUserId() bool {
	if o != nil && o.OnBehalfOfUserId.IsSet() {
		return true
	}

	return false
}

// SetOnBehalfOfUserId gets a reference to the given NullableString and assigns it to the OnBehalfOfUserId field.
func (o *InlineObject241) SetOnBehalfOfUserId(v string) {
	o.OnBehalfOfUserId.Set(&v)
}
// SetOnBehalfOfUserIdNil sets the value for OnBehalfOfUserId to be an explicit nil
func (o *InlineObject241) SetOnBehalfOfUserIdNil() {
	o.OnBehalfOfUserId.Set(nil)
}

// UnsetOnBehalfOfUserId ensures that no value is present for OnBehalfOfUserId, not even an explicit nil
func (o *InlineObject241) UnsetOnBehalfOfUserId() {
	o.OnBehalfOfUserId.Unset()
}

func (o InlineObject241) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
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

type NullableInlineObject241 struct {
	value *InlineObject241
	isSet bool
}

func (v NullableInlineObject241) Get() *InlineObject241 {
	return v.value
}

func (v *NullableInlineObject241) Set(val *InlineObject241) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject241) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject241) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject241(val *InlineObject241) *NullableInlineObject241 {
	return &NullableInlineObject241{value: val, isSet: true}
}

func (v NullableInlineObject241) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject241) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


