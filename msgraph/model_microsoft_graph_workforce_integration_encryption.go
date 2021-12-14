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

// MicrosoftGraphWorkforceIntegrationEncryption struct for MicrosoftGraphWorkforceIntegrationEncryption
type MicrosoftGraphWorkforceIntegrationEncryption struct {
	// Possible values are: sharedSecret, unknownFutureValue.
	Protocol AnyOfmicrosoftGraphWorkforceIntegrationEncryptionProtocol `json:"protocol,omitempty"`
	// Encryption shared secret.
	Secret NullableString `json:"secret,omitempty"`
}

// NewMicrosoftGraphWorkforceIntegrationEncryption instantiates a new MicrosoftGraphWorkforceIntegrationEncryption object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphWorkforceIntegrationEncryption() *MicrosoftGraphWorkforceIntegrationEncryption {
	this := MicrosoftGraphWorkforceIntegrationEncryption{}
	return &this
}

// NewMicrosoftGraphWorkforceIntegrationEncryptionWithDefaults instantiates a new MicrosoftGraphWorkforceIntegrationEncryption object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphWorkforceIntegrationEncryptionWithDefaults() *MicrosoftGraphWorkforceIntegrationEncryption {
	this := MicrosoftGraphWorkforceIntegrationEncryption{}
	return &this
}

// GetProtocol returns the Protocol field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphWorkforceIntegrationEncryption) GetProtocol() AnyOfmicrosoftGraphWorkforceIntegrationEncryptionProtocol {
	if o == nil  {
		var ret AnyOfmicrosoftGraphWorkforceIntegrationEncryptionProtocol
		return ret
	}
	return o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphWorkforceIntegrationEncryption) GetProtocolOk() (*AnyOfmicrosoftGraphWorkforceIntegrationEncryptionProtocol, bool) {
	if o == nil || o.Protocol == nil {
		return nil, false
	}
	return &o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *MicrosoftGraphWorkforceIntegrationEncryption) HasProtocol() bool {
	if o != nil && o.Protocol != nil {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given AnyOfmicrosoftGraphWorkforceIntegrationEncryptionProtocol and assigns it to the Protocol field.
func (o *MicrosoftGraphWorkforceIntegrationEncryption) SetProtocol(v AnyOfmicrosoftGraphWorkforceIntegrationEncryptionProtocol) {
	o.Protocol = v
}

// GetSecret returns the Secret field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphWorkforceIntegrationEncryption) GetSecret() string {
	if o == nil || o.Secret.Get() == nil {
		var ret string
		return ret
	}
	return *o.Secret.Get()
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphWorkforceIntegrationEncryption) GetSecretOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Secret.Get(), o.Secret.IsSet()
}

// HasSecret returns a boolean if a field has been set.
func (o *MicrosoftGraphWorkforceIntegrationEncryption) HasSecret() bool {
	if o != nil && o.Secret.IsSet() {
		return true
	}

	return false
}

// SetSecret gets a reference to the given NullableString and assigns it to the Secret field.
func (o *MicrosoftGraphWorkforceIntegrationEncryption) SetSecret(v string) {
	o.Secret.Set(&v)
}
// SetSecretNil sets the value for Secret to be an explicit nil
func (o *MicrosoftGraphWorkforceIntegrationEncryption) SetSecretNil() {
	o.Secret.Set(nil)
}

// UnsetSecret ensures that no value is present for Secret, not even an explicit nil
func (o *MicrosoftGraphWorkforceIntegrationEncryption) UnsetSecret() {
	o.Secret.Unset()
}

func (o MicrosoftGraphWorkforceIntegrationEncryption) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Protocol != nil {
		toSerialize["protocol"] = o.Protocol
	}
	if o.Secret.IsSet() {
		toSerialize["secret"] = o.Secret.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphWorkforceIntegrationEncryption struct {
	value *MicrosoftGraphWorkforceIntegrationEncryption
	isSet bool
}

func (v NullableMicrosoftGraphWorkforceIntegrationEncryption) Get() *MicrosoftGraphWorkforceIntegrationEncryption {
	return v.value
}

func (v *NullableMicrosoftGraphWorkforceIntegrationEncryption) Set(val *MicrosoftGraphWorkforceIntegrationEncryption) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphWorkforceIntegrationEncryption) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphWorkforceIntegrationEncryption) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphWorkforceIntegrationEncryption(val *MicrosoftGraphWorkforceIntegrationEncryption) *NullableMicrosoftGraphWorkforceIntegrationEncryption {
	return &NullableMicrosoftGraphWorkforceIntegrationEncryption{value: val, isSet: true}
}

func (v NullableMicrosoftGraphWorkforceIntegrationEncryption) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphWorkforceIntegrationEncryption) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


