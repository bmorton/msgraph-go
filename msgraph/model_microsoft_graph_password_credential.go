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

// MicrosoftGraphPasswordCredential struct for MicrosoftGraphPasswordCredential
type MicrosoftGraphPasswordCredential struct {
	// Do not use.
	CustomKeyIdentifier NullableString `json:"customKeyIdentifier,omitempty"`
	// Friendly name for the password. Optional.
	DisplayName NullableString `json:"displayName,omitempty"`
	// The date and time at which the password expires represented using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Optional.
	EndDateTime NullableTime `json:"endDateTime,omitempty"`
	// Contains the first three characters of the password. Read-only.
	Hint NullableString `json:"hint,omitempty"`
	// The unique identifier for the password.
	KeyId NullableString `json:"keyId,omitempty"`
	// Read-only; Contains the strong passwords generated by Azure AD that are 16-64 characters in length. The generated password value is only returned during the initial POST request to addPassword. There is no way to retrieve this password in the future.
	SecretText NullableString `json:"secretText,omitempty"`
	// The date and time at which the password becomes valid. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Optional.
	StartDateTime NullableTime `json:"startDateTime,omitempty"`
}

// NewMicrosoftGraphPasswordCredential instantiates a new MicrosoftGraphPasswordCredential object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphPasswordCredential() *MicrosoftGraphPasswordCredential {
	this := MicrosoftGraphPasswordCredential{}
	return &this
}

// NewMicrosoftGraphPasswordCredentialWithDefaults instantiates a new MicrosoftGraphPasswordCredential object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphPasswordCredentialWithDefaults() *MicrosoftGraphPasswordCredential {
	this := MicrosoftGraphPasswordCredential{}
	return &this
}

// GetCustomKeyIdentifier returns the CustomKeyIdentifier field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphPasswordCredential) GetCustomKeyIdentifier() string {
	if o == nil || o.CustomKeyIdentifier.Get() == nil {
		var ret string
		return ret
	}
	return *o.CustomKeyIdentifier.Get()
}

// GetCustomKeyIdentifierOk returns a tuple with the CustomKeyIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphPasswordCredential) GetCustomKeyIdentifierOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CustomKeyIdentifier.Get(), o.CustomKeyIdentifier.IsSet()
}

// HasCustomKeyIdentifier returns a boolean if a field has been set.
func (o *MicrosoftGraphPasswordCredential) HasCustomKeyIdentifier() bool {
	if o != nil && o.CustomKeyIdentifier.IsSet() {
		return true
	}

	return false
}

// SetCustomKeyIdentifier gets a reference to the given NullableString and assigns it to the CustomKeyIdentifier field.
func (o *MicrosoftGraphPasswordCredential) SetCustomKeyIdentifier(v string) {
	o.CustomKeyIdentifier.Set(&v)
}
// SetCustomKeyIdentifierNil sets the value for CustomKeyIdentifier to be an explicit nil
func (o *MicrosoftGraphPasswordCredential) SetCustomKeyIdentifierNil() {
	o.CustomKeyIdentifier.Set(nil)
}

// UnsetCustomKeyIdentifier ensures that no value is present for CustomKeyIdentifier, not even an explicit nil
func (o *MicrosoftGraphPasswordCredential) UnsetCustomKeyIdentifier() {
	o.CustomKeyIdentifier.Unset()
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphPasswordCredential) GetDisplayName() string {
	if o == nil || o.DisplayName.Get() == nil {
		var ret string
		return ret
	}
	return *o.DisplayName.Get()
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphPasswordCredential) GetDisplayNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DisplayName.Get(), o.DisplayName.IsSet()
}

// HasDisplayName returns a boolean if a field has been set.
func (o *MicrosoftGraphPasswordCredential) HasDisplayName() bool {
	if o != nil && o.DisplayName.IsSet() {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given NullableString and assigns it to the DisplayName field.
func (o *MicrosoftGraphPasswordCredential) SetDisplayName(v string) {
	o.DisplayName.Set(&v)
}
// SetDisplayNameNil sets the value for DisplayName to be an explicit nil
func (o *MicrosoftGraphPasswordCredential) SetDisplayNameNil() {
	o.DisplayName.Set(nil)
}

// UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
func (o *MicrosoftGraphPasswordCredential) UnsetDisplayName() {
	o.DisplayName.Unset()
}

// GetEndDateTime returns the EndDateTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphPasswordCredential) GetEndDateTime() time.Time {
	if o == nil || o.EndDateTime.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.EndDateTime.Get()
}

// GetEndDateTimeOk returns a tuple with the EndDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphPasswordCredential) GetEndDateTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.EndDateTime.Get(), o.EndDateTime.IsSet()
}

// HasEndDateTime returns a boolean if a field has been set.
func (o *MicrosoftGraphPasswordCredential) HasEndDateTime() bool {
	if o != nil && o.EndDateTime.IsSet() {
		return true
	}

	return false
}

// SetEndDateTime gets a reference to the given NullableTime and assigns it to the EndDateTime field.
func (o *MicrosoftGraphPasswordCredential) SetEndDateTime(v time.Time) {
	o.EndDateTime.Set(&v)
}
// SetEndDateTimeNil sets the value for EndDateTime to be an explicit nil
func (o *MicrosoftGraphPasswordCredential) SetEndDateTimeNil() {
	o.EndDateTime.Set(nil)
}

// UnsetEndDateTime ensures that no value is present for EndDateTime, not even an explicit nil
func (o *MicrosoftGraphPasswordCredential) UnsetEndDateTime() {
	o.EndDateTime.Unset()
}

// GetHint returns the Hint field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphPasswordCredential) GetHint() string {
	if o == nil || o.Hint.Get() == nil {
		var ret string
		return ret
	}
	return *o.Hint.Get()
}

// GetHintOk returns a tuple with the Hint field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphPasswordCredential) GetHintOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Hint.Get(), o.Hint.IsSet()
}

// HasHint returns a boolean if a field has been set.
func (o *MicrosoftGraphPasswordCredential) HasHint() bool {
	if o != nil && o.Hint.IsSet() {
		return true
	}

	return false
}

// SetHint gets a reference to the given NullableString and assigns it to the Hint field.
func (o *MicrosoftGraphPasswordCredential) SetHint(v string) {
	o.Hint.Set(&v)
}
// SetHintNil sets the value for Hint to be an explicit nil
func (o *MicrosoftGraphPasswordCredential) SetHintNil() {
	o.Hint.Set(nil)
}

// UnsetHint ensures that no value is present for Hint, not even an explicit nil
func (o *MicrosoftGraphPasswordCredential) UnsetHint() {
	o.Hint.Unset()
}

// GetKeyId returns the KeyId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphPasswordCredential) GetKeyId() string {
	if o == nil || o.KeyId.Get() == nil {
		var ret string
		return ret
	}
	return *o.KeyId.Get()
}

// GetKeyIdOk returns a tuple with the KeyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphPasswordCredential) GetKeyIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.KeyId.Get(), o.KeyId.IsSet()
}

// HasKeyId returns a boolean if a field has been set.
func (o *MicrosoftGraphPasswordCredential) HasKeyId() bool {
	if o != nil && o.KeyId.IsSet() {
		return true
	}

	return false
}

// SetKeyId gets a reference to the given NullableString and assigns it to the KeyId field.
func (o *MicrosoftGraphPasswordCredential) SetKeyId(v string) {
	o.KeyId.Set(&v)
}
// SetKeyIdNil sets the value for KeyId to be an explicit nil
func (o *MicrosoftGraphPasswordCredential) SetKeyIdNil() {
	o.KeyId.Set(nil)
}

// UnsetKeyId ensures that no value is present for KeyId, not even an explicit nil
func (o *MicrosoftGraphPasswordCredential) UnsetKeyId() {
	o.KeyId.Unset()
}

// GetSecretText returns the SecretText field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphPasswordCredential) GetSecretText() string {
	if o == nil || o.SecretText.Get() == nil {
		var ret string
		return ret
	}
	return *o.SecretText.Get()
}

// GetSecretTextOk returns a tuple with the SecretText field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphPasswordCredential) GetSecretTextOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SecretText.Get(), o.SecretText.IsSet()
}

// HasSecretText returns a boolean if a field has been set.
func (o *MicrosoftGraphPasswordCredential) HasSecretText() bool {
	if o != nil && o.SecretText.IsSet() {
		return true
	}

	return false
}

// SetSecretText gets a reference to the given NullableString and assigns it to the SecretText field.
func (o *MicrosoftGraphPasswordCredential) SetSecretText(v string) {
	o.SecretText.Set(&v)
}
// SetSecretTextNil sets the value for SecretText to be an explicit nil
func (o *MicrosoftGraphPasswordCredential) SetSecretTextNil() {
	o.SecretText.Set(nil)
}

// UnsetSecretText ensures that no value is present for SecretText, not even an explicit nil
func (o *MicrosoftGraphPasswordCredential) UnsetSecretText() {
	o.SecretText.Unset()
}

// GetStartDateTime returns the StartDateTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphPasswordCredential) GetStartDateTime() time.Time {
	if o == nil || o.StartDateTime.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.StartDateTime.Get()
}

// GetStartDateTimeOk returns a tuple with the StartDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphPasswordCredential) GetStartDateTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.StartDateTime.Get(), o.StartDateTime.IsSet()
}

// HasStartDateTime returns a boolean if a field has been set.
func (o *MicrosoftGraphPasswordCredential) HasStartDateTime() bool {
	if o != nil && o.StartDateTime.IsSet() {
		return true
	}

	return false
}

// SetStartDateTime gets a reference to the given NullableTime and assigns it to the StartDateTime field.
func (o *MicrosoftGraphPasswordCredential) SetStartDateTime(v time.Time) {
	o.StartDateTime.Set(&v)
}
// SetStartDateTimeNil sets the value for StartDateTime to be an explicit nil
func (o *MicrosoftGraphPasswordCredential) SetStartDateTimeNil() {
	o.StartDateTime.Set(nil)
}

// UnsetStartDateTime ensures that no value is present for StartDateTime, not even an explicit nil
func (o *MicrosoftGraphPasswordCredential) UnsetStartDateTime() {
	o.StartDateTime.Unset()
}

func (o MicrosoftGraphPasswordCredential) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CustomKeyIdentifier.IsSet() {
		toSerialize["customKeyIdentifier"] = o.CustomKeyIdentifier.Get()
	}
	if o.DisplayName.IsSet() {
		toSerialize["displayName"] = o.DisplayName.Get()
	}
	if o.EndDateTime.IsSet() {
		toSerialize["endDateTime"] = o.EndDateTime.Get()
	}
	if o.Hint.IsSet() {
		toSerialize["hint"] = o.Hint.Get()
	}
	if o.KeyId.IsSet() {
		toSerialize["keyId"] = o.KeyId.Get()
	}
	if o.SecretText.IsSet() {
		toSerialize["secretText"] = o.SecretText.Get()
	}
	if o.StartDateTime.IsSet() {
		toSerialize["startDateTime"] = o.StartDateTime.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphPasswordCredential struct {
	value *MicrosoftGraphPasswordCredential
	isSet bool
}

func (v NullableMicrosoftGraphPasswordCredential) Get() *MicrosoftGraphPasswordCredential {
	return v.value
}

func (v *NullableMicrosoftGraphPasswordCredential) Set(val *MicrosoftGraphPasswordCredential) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphPasswordCredential) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphPasswordCredential) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphPasswordCredential(val *MicrosoftGraphPasswordCredential) *NullableMicrosoftGraphPasswordCredential {
	return &NullableMicrosoftGraphPasswordCredential{value: val, isSet: true}
}

func (v NullableMicrosoftGraphPasswordCredential) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphPasswordCredential) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

