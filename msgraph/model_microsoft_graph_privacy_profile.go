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

// MicrosoftGraphPrivacyProfile struct for MicrosoftGraphPrivacyProfile
type MicrosoftGraphPrivacyProfile struct {
	// A valid smtp email address for the privacy statement contact. Not required.
	ContactEmail NullableString `json:"contactEmail,omitempty"`
	// A valid URL format that begins with http:// or https://. Maximum length is 255 characters. The URL that directs to the company's privacy statement. Not required.
	StatementUrl NullableString `json:"statementUrl,omitempty"`
}

// NewMicrosoftGraphPrivacyProfile instantiates a new MicrosoftGraphPrivacyProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphPrivacyProfile() *MicrosoftGraphPrivacyProfile {
	this := MicrosoftGraphPrivacyProfile{}
	return &this
}

// NewMicrosoftGraphPrivacyProfileWithDefaults instantiates a new MicrosoftGraphPrivacyProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphPrivacyProfileWithDefaults() *MicrosoftGraphPrivacyProfile {
	this := MicrosoftGraphPrivacyProfile{}
	return &this
}

// GetContactEmail returns the ContactEmail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphPrivacyProfile) GetContactEmail() string {
	if o == nil || o.ContactEmail.Get() == nil {
		var ret string
		return ret
	}
	return *o.ContactEmail.Get()
}

// GetContactEmailOk returns a tuple with the ContactEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphPrivacyProfile) GetContactEmailOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ContactEmail.Get(), o.ContactEmail.IsSet()
}

// HasContactEmail returns a boolean if a field has been set.
func (o *MicrosoftGraphPrivacyProfile) HasContactEmail() bool {
	if o != nil && o.ContactEmail.IsSet() {
		return true
	}

	return false
}

// SetContactEmail gets a reference to the given NullableString and assigns it to the ContactEmail field.
func (o *MicrosoftGraphPrivacyProfile) SetContactEmail(v string) {
	o.ContactEmail.Set(&v)
}
// SetContactEmailNil sets the value for ContactEmail to be an explicit nil
func (o *MicrosoftGraphPrivacyProfile) SetContactEmailNil() {
	o.ContactEmail.Set(nil)
}

// UnsetContactEmail ensures that no value is present for ContactEmail, not even an explicit nil
func (o *MicrosoftGraphPrivacyProfile) UnsetContactEmail() {
	o.ContactEmail.Unset()
}

// GetStatementUrl returns the StatementUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphPrivacyProfile) GetStatementUrl() string {
	if o == nil || o.StatementUrl.Get() == nil {
		var ret string
		return ret
	}
	return *o.StatementUrl.Get()
}

// GetStatementUrlOk returns a tuple with the StatementUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphPrivacyProfile) GetStatementUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.StatementUrl.Get(), o.StatementUrl.IsSet()
}

// HasStatementUrl returns a boolean if a field has been set.
func (o *MicrosoftGraphPrivacyProfile) HasStatementUrl() bool {
	if o != nil && o.StatementUrl.IsSet() {
		return true
	}

	return false
}

// SetStatementUrl gets a reference to the given NullableString and assigns it to the StatementUrl field.
func (o *MicrosoftGraphPrivacyProfile) SetStatementUrl(v string) {
	o.StatementUrl.Set(&v)
}
// SetStatementUrlNil sets the value for StatementUrl to be an explicit nil
func (o *MicrosoftGraphPrivacyProfile) SetStatementUrlNil() {
	o.StatementUrl.Set(nil)
}

// UnsetStatementUrl ensures that no value is present for StatementUrl, not even an explicit nil
func (o *MicrosoftGraphPrivacyProfile) UnsetStatementUrl() {
	o.StatementUrl.Unset()
}

func (o MicrosoftGraphPrivacyProfile) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ContactEmail.IsSet() {
		toSerialize["contactEmail"] = o.ContactEmail.Get()
	}
	if o.StatementUrl.IsSet() {
		toSerialize["statementUrl"] = o.StatementUrl.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphPrivacyProfile struct {
	value *MicrosoftGraphPrivacyProfile
	isSet bool
}

func (v NullableMicrosoftGraphPrivacyProfile) Get() *MicrosoftGraphPrivacyProfile {
	return v.value
}

func (v *NullableMicrosoftGraphPrivacyProfile) Set(val *MicrosoftGraphPrivacyProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphPrivacyProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphPrivacyProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphPrivacyProfile(val *MicrosoftGraphPrivacyProfile) *NullableMicrosoftGraphPrivacyProfile {
	return &NullableMicrosoftGraphPrivacyProfile{value: val, isSet: true}
}

func (v NullableMicrosoftGraphPrivacyProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphPrivacyProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


