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

// Authentication struct for Authentication
type Authentication struct {
	Fido2Methods *[]MicrosoftGraphFido2AuthenticationMethod `json:"fido2Methods,omitempty"`
	Methods *[]MicrosoftGraphAuthenticationMethod `json:"methods,omitempty"`
	MicrosoftAuthenticatorMethods *[]MicrosoftGraphMicrosoftAuthenticatorAuthenticationMethod `json:"microsoftAuthenticatorMethods,omitempty"`
	WindowsHelloForBusinessMethods *[]MicrosoftGraphWindowsHelloForBusinessAuthenticationMethod `json:"windowsHelloForBusinessMethods,omitempty"`
}

// NewAuthentication instantiates a new Authentication object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthentication() *Authentication {
	this := Authentication{}
	return &this
}

// NewAuthenticationWithDefaults instantiates a new Authentication object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticationWithDefaults() *Authentication {
	this := Authentication{}
	return &this
}

// GetFido2Methods returns the Fido2Methods field value if set, zero value otherwise.
func (o *Authentication) GetFido2Methods() []MicrosoftGraphFido2AuthenticationMethod {
	if o == nil || o.Fido2Methods == nil {
		var ret []MicrosoftGraphFido2AuthenticationMethod
		return ret
	}
	return *o.Fido2Methods
}

// GetFido2MethodsOk returns a tuple with the Fido2Methods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Authentication) GetFido2MethodsOk() (*[]MicrosoftGraphFido2AuthenticationMethod, bool) {
	if o == nil || o.Fido2Methods == nil {
		return nil, false
	}
	return o.Fido2Methods, true
}

// HasFido2Methods returns a boolean if a field has been set.
func (o *Authentication) HasFido2Methods() bool {
	if o != nil && o.Fido2Methods != nil {
		return true
	}

	return false
}

// SetFido2Methods gets a reference to the given []MicrosoftGraphFido2AuthenticationMethod and assigns it to the Fido2Methods field.
func (o *Authentication) SetFido2Methods(v []MicrosoftGraphFido2AuthenticationMethod) {
	o.Fido2Methods = &v
}

// GetMethods returns the Methods field value if set, zero value otherwise.
func (o *Authentication) GetMethods() []MicrosoftGraphAuthenticationMethod {
	if o == nil || o.Methods == nil {
		var ret []MicrosoftGraphAuthenticationMethod
		return ret
	}
	return *o.Methods
}

// GetMethodsOk returns a tuple with the Methods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Authentication) GetMethodsOk() (*[]MicrosoftGraphAuthenticationMethod, bool) {
	if o == nil || o.Methods == nil {
		return nil, false
	}
	return o.Methods, true
}

// HasMethods returns a boolean if a field has been set.
func (o *Authentication) HasMethods() bool {
	if o != nil && o.Methods != nil {
		return true
	}

	return false
}

// SetMethods gets a reference to the given []MicrosoftGraphAuthenticationMethod and assigns it to the Methods field.
func (o *Authentication) SetMethods(v []MicrosoftGraphAuthenticationMethod) {
	o.Methods = &v
}

// GetMicrosoftAuthenticatorMethods returns the MicrosoftAuthenticatorMethods field value if set, zero value otherwise.
func (o *Authentication) GetMicrosoftAuthenticatorMethods() []MicrosoftGraphMicrosoftAuthenticatorAuthenticationMethod {
	if o == nil || o.MicrosoftAuthenticatorMethods == nil {
		var ret []MicrosoftGraphMicrosoftAuthenticatorAuthenticationMethod
		return ret
	}
	return *o.MicrosoftAuthenticatorMethods
}

// GetMicrosoftAuthenticatorMethodsOk returns a tuple with the MicrosoftAuthenticatorMethods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Authentication) GetMicrosoftAuthenticatorMethodsOk() (*[]MicrosoftGraphMicrosoftAuthenticatorAuthenticationMethod, bool) {
	if o == nil || o.MicrosoftAuthenticatorMethods == nil {
		return nil, false
	}
	return o.MicrosoftAuthenticatorMethods, true
}

// HasMicrosoftAuthenticatorMethods returns a boolean if a field has been set.
func (o *Authentication) HasMicrosoftAuthenticatorMethods() bool {
	if o != nil && o.MicrosoftAuthenticatorMethods != nil {
		return true
	}

	return false
}

// SetMicrosoftAuthenticatorMethods gets a reference to the given []MicrosoftGraphMicrosoftAuthenticatorAuthenticationMethod and assigns it to the MicrosoftAuthenticatorMethods field.
func (o *Authentication) SetMicrosoftAuthenticatorMethods(v []MicrosoftGraphMicrosoftAuthenticatorAuthenticationMethod) {
	o.MicrosoftAuthenticatorMethods = &v
}

// GetWindowsHelloForBusinessMethods returns the WindowsHelloForBusinessMethods field value if set, zero value otherwise.
func (o *Authentication) GetWindowsHelloForBusinessMethods() []MicrosoftGraphWindowsHelloForBusinessAuthenticationMethod {
	if o == nil || o.WindowsHelloForBusinessMethods == nil {
		var ret []MicrosoftGraphWindowsHelloForBusinessAuthenticationMethod
		return ret
	}
	return *o.WindowsHelloForBusinessMethods
}

// GetWindowsHelloForBusinessMethodsOk returns a tuple with the WindowsHelloForBusinessMethods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Authentication) GetWindowsHelloForBusinessMethodsOk() (*[]MicrosoftGraphWindowsHelloForBusinessAuthenticationMethod, bool) {
	if o == nil || o.WindowsHelloForBusinessMethods == nil {
		return nil, false
	}
	return o.WindowsHelloForBusinessMethods, true
}

// HasWindowsHelloForBusinessMethods returns a boolean if a field has been set.
func (o *Authentication) HasWindowsHelloForBusinessMethods() bool {
	if o != nil && o.WindowsHelloForBusinessMethods != nil {
		return true
	}

	return false
}

// SetWindowsHelloForBusinessMethods gets a reference to the given []MicrosoftGraphWindowsHelloForBusinessAuthenticationMethod and assigns it to the WindowsHelloForBusinessMethods field.
func (o *Authentication) SetWindowsHelloForBusinessMethods(v []MicrosoftGraphWindowsHelloForBusinessAuthenticationMethod) {
	o.WindowsHelloForBusinessMethods = &v
}

func (o Authentication) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Fido2Methods != nil {
		toSerialize["fido2Methods"] = o.Fido2Methods
	}
	if o.Methods != nil {
		toSerialize["methods"] = o.Methods
	}
	if o.MicrosoftAuthenticatorMethods != nil {
		toSerialize["microsoftAuthenticatorMethods"] = o.MicrosoftAuthenticatorMethods
	}
	if o.WindowsHelloForBusinessMethods != nil {
		toSerialize["windowsHelloForBusinessMethods"] = o.WindowsHelloForBusinessMethods
	}
	return json.Marshal(toSerialize)
}

type NullableAuthentication struct {
	value *Authentication
	isSet bool
}

func (v NullableAuthentication) Get() *Authentication {
	return v.value
}

func (v *NullableAuthentication) Set(val *Authentication) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthentication) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthentication) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthentication(val *Authentication) *NullableAuthentication {
	return &NullableAuthentication{value: val, isSet: true}
}

func (v NullableAuthentication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthentication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


