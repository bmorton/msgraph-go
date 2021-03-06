/*
Partial Graph API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package msgraph

import (
	"encoding/json"
	"time"
)

// InlineObject1198 struct for InlineObject1198
type InlineObject1198 struct {
	Type *string `json:"type,omitempty"`
	Scope NullableString `json:"scope,omitempty"`
	ExpirationDateTime NullableTime `json:"expirationDateTime,omitempty"`
	Password NullableString `json:"password,omitempty"`
	Message NullableString `json:"message,omitempty"`
}

// NewInlineObject1198 instantiates a new InlineObject1198 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject1198() *InlineObject1198 {
	this := InlineObject1198{}
	return &this
}

// NewInlineObject1198WithDefaults instantiates a new InlineObject1198 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject1198WithDefaults() *InlineObject1198 {
	this := InlineObject1198{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *InlineObject1198) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject1198) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *InlineObject1198) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *InlineObject1198) SetType(v string) {
	o.Type = &v
}

// GetScope returns the Scope field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1198) GetScope() string {
	if o == nil || o.Scope.Get() == nil {
		var ret string
		return ret
	}
	return *o.Scope.Get()
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1198) GetScopeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Scope.Get(), o.Scope.IsSet()
}

// HasScope returns a boolean if a field has been set.
func (o *InlineObject1198) HasScope() bool {
	if o != nil && o.Scope.IsSet() {
		return true
	}

	return false
}

// SetScope gets a reference to the given NullableString and assigns it to the Scope field.
func (o *InlineObject1198) SetScope(v string) {
	o.Scope.Set(&v)
}
// SetScopeNil sets the value for Scope to be an explicit nil
func (o *InlineObject1198) SetScopeNil() {
	o.Scope.Set(nil)
}

// UnsetScope ensures that no value is present for Scope, not even an explicit nil
func (o *InlineObject1198) UnsetScope() {
	o.Scope.Unset()
}

// GetExpirationDateTime returns the ExpirationDateTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1198) GetExpirationDateTime() time.Time {
	if o == nil || o.ExpirationDateTime.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.ExpirationDateTime.Get()
}

// GetExpirationDateTimeOk returns a tuple with the ExpirationDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1198) GetExpirationDateTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ExpirationDateTime.Get(), o.ExpirationDateTime.IsSet()
}

// HasExpirationDateTime returns a boolean if a field has been set.
func (o *InlineObject1198) HasExpirationDateTime() bool {
	if o != nil && o.ExpirationDateTime.IsSet() {
		return true
	}

	return false
}

// SetExpirationDateTime gets a reference to the given NullableTime and assigns it to the ExpirationDateTime field.
func (o *InlineObject1198) SetExpirationDateTime(v time.Time) {
	o.ExpirationDateTime.Set(&v)
}
// SetExpirationDateTimeNil sets the value for ExpirationDateTime to be an explicit nil
func (o *InlineObject1198) SetExpirationDateTimeNil() {
	o.ExpirationDateTime.Set(nil)
}

// UnsetExpirationDateTime ensures that no value is present for ExpirationDateTime, not even an explicit nil
func (o *InlineObject1198) UnsetExpirationDateTime() {
	o.ExpirationDateTime.Unset()
}

// GetPassword returns the Password field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1198) GetPassword() string {
	if o == nil || o.Password.Get() == nil {
		var ret string
		return ret
	}
	return *o.Password.Get()
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1198) GetPasswordOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Password.Get(), o.Password.IsSet()
}

// HasPassword returns a boolean if a field has been set.
func (o *InlineObject1198) HasPassword() bool {
	if o != nil && o.Password.IsSet() {
		return true
	}

	return false
}

// SetPassword gets a reference to the given NullableString and assigns it to the Password field.
func (o *InlineObject1198) SetPassword(v string) {
	o.Password.Set(&v)
}
// SetPasswordNil sets the value for Password to be an explicit nil
func (o *InlineObject1198) SetPasswordNil() {
	o.Password.Set(nil)
}

// UnsetPassword ensures that no value is present for Password, not even an explicit nil
func (o *InlineObject1198) UnsetPassword() {
	o.Password.Unset()
}

// GetMessage returns the Message field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1198) GetMessage() string {
	if o == nil || o.Message.Get() == nil {
		var ret string
		return ret
	}
	return *o.Message.Get()
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1198) GetMessageOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Message.Get(), o.Message.IsSet()
}

// HasMessage returns a boolean if a field has been set.
func (o *InlineObject1198) HasMessage() bool {
	if o != nil && o.Message.IsSet() {
		return true
	}

	return false
}

// SetMessage gets a reference to the given NullableString and assigns it to the Message field.
func (o *InlineObject1198) SetMessage(v string) {
	o.Message.Set(&v)
}
// SetMessageNil sets the value for Message to be an explicit nil
func (o *InlineObject1198) SetMessageNil() {
	o.Message.Set(nil)
}

// UnsetMessage ensures that no value is present for Message, not even an explicit nil
func (o *InlineObject1198) UnsetMessage() {
	o.Message.Unset()
}

func (o InlineObject1198) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Scope.IsSet() {
		toSerialize["scope"] = o.Scope.Get()
	}
	if o.ExpirationDateTime.IsSet() {
		toSerialize["expirationDateTime"] = o.ExpirationDateTime.Get()
	}
	if o.Password.IsSet() {
		toSerialize["password"] = o.Password.Get()
	}
	if o.Message.IsSet() {
		toSerialize["message"] = o.Message.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject1198 struct {
	value *InlineObject1198
	isSet bool
}

func (v NullableInlineObject1198) Get() *InlineObject1198 {
	return v.value
}

func (v *NullableInlineObject1198) Set(val *InlineObject1198) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject1198) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject1198) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject1198(val *InlineObject1198) *NullableInlineObject1198 {
	return &NullableInlineObject1198{value: val, isSet: true}
}

func (v NullableInlineObject1198) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject1198) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


