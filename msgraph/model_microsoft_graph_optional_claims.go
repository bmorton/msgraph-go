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

// MicrosoftGraphOptionalClaims struct for MicrosoftGraphOptionalClaims
type MicrosoftGraphOptionalClaims struct {
	// The optional claims returned in the JWT access token.
	AccessToken *[]*AnyOfmicrosoftGraphOptionalClaim `json:"accessToken,omitempty"`
	// The optional claims returned in the JWT ID token.
	IdToken *[]*AnyOfmicrosoftGraphOptionalClaim `json:"idToken,omitempty"`
	// The optional claims returned in the SAML token.
	Saml2Token *[]*AnyOfmicrosoftGraphOptionalClaim `json:"saml2Token,omitempty"`
}

// NewMicrosoftGraphOptionalClaims instantiates a new MicrosoftGraphOptionalClaims object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphOptionalClaims() *MicrosoftGraphOptionalClaims {
	this := MicrosoftGraphOptionalClaims{}
	return &this
}

// NewMicrosoftGraphOptionalClaimsWithDefaults instantiates a new MicrosoftGraphOptionalClaims object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphOptionalClaimsWithDefaults() *MicrosoftGraphOptionalClaims {
	this := MicrosoftGraphOptionalClaims{}
	return &this
}

// GetAccessToken returns the AccessToken field value if set, zero value otherwise.
func (o *MicrosoftGraphOptionalClaims) GetAccessToken() []*AnyOfmicrosoftGraphOptionalClaim {
	if o == nil || o.AccessToken == nil {
		var ret []*AnyOfmicrosoftGraphOptionalClaim
		return ret
	}
	return *o.AccessToken
}

// GetAccessTokenOk returns a tuple with the AccessToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphOptionalClaims) GetAccessTokenOk() (*[]*AnyOfmicrosoftGraphOptionalClaim, bool) {
	if o == nil || o.AccessToken == nil {
		return nil, false
	}
	return o.AccessToken, true
}

// HasAccessToken returns a boolean if a field has been set.
func (o *MicrosoftGraphOptionalClaims) HasAccessToken() bool {
	if o != nil && o.AccessToken != nil {
		return true
	}

	return false
}

// SetAccessToken gets a reference to the given []*AnyOfmicrosoftGraphOptionalClaim and assigns it to the AccessToken field.
func (o *MicrosoftGraphOptionalClaims) SetAccessToken(v []*AnyOfmicrosoftGraphOptionalClaim) {
	o.AccessToken = &v
}

// GetIdToken returns the IdToken field value if set, zero value otherwise.
func (o *MicrosoftGraphOptionalClaims) GetIdToken() []*AnyOfmicrosoftGraphOptionalClaim {
	if o == nil || o.IdToken == nil {
		var ret []*AnyOfmicrosoftGraphOptionalClaim
		return ret
	}
	return *o.IdToken
}

// GetIdTokenOk returns a tuple with the IdToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphOptionalClaims) GetIdTokenOk() (*[]*AnyOfmicrosoftGraphOptionalClaim, bool) {
	if o == nil || o.IdToken == nil {
		return nil, false
	}
	return o.IdToken, true
}

// HasIdToken returns a boolean if a field has been set.
func (o *MicrosoftGraphOptionalClaims) HasIdToken() bool {
	if o != nil && o.IdToken != nil {
		return true
	}

	return false
}

// SetIdToken gets a reference to the given []*AnyOfmicrosoftGraphOptionalClaim and assigns it to the IdToken field.
func (o *MicrosoftGraphOptionalClaims) SetIdToken(v []*AnyOfmicrosoftGraphOptionalClaim) {
	o.IdToken = &v
}

// GetSaml2Token returns the Saml2Token field value if set, zero value otherwise.
func (o *MicrosoftGraphOptionalClaims) GetSaml2Token() []*AnyOfmicrosoftGraphOptionalClaim {
	if o == nil || o.Saml2Token == nil {
		var ret []*AnyOfmicrosoftGraphOptionalClaim
		return ret
	}
	return *o.Saml2Token
}

// GetSaml2TokenOk returns a tuple with the Saml2Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphOptionalClaims) GetSaml2TokenOk() (*[]*AnyOfmicrosoftGraphOptionalClaim, bool) {
	if o == nil || o.Saml2Token == nil {
		return nil, false
	}
	return o.Saml2Token, true
}

// HasSaml2Token returns a boolean if a field has been set.
func (o *MicrosoftGraphOptionalClaims) HasSaml2Token() bool {
	if o != nil && o.Saml2Token != nil {
		return true
	}

	return false
}

// SetSaml2Token gets a reference to the given []*AnyOfmicrosoftGraphOptionalClaim and assigns it to the Saml2Token field.
func (o *MicrosoftGraphOptionalClaims) SetSaml2Token(v []*AnyOfmicrosoftGraphOptionalClaim) {
	o.Saml2Token = &v
}

func (o MicrosoftGraphOptionalClaims) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccessToken != nil {
		toSerialize["accessToken"] = o.AccessToken
	}
	if o.IdToken != nil {
		toSerialize["idToken"] = o.IdToken
	}
	if o.Saml2Token != nil {
		toSerialize["saml2Token"] = o.Saml2Token
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphOptionalClaims struct {
	value *MicrosoftGraphOptionalClaims
	isSet bool
}

func (v NullableMicrosoftGraphOptionalClaims) Get() *MicrosoftGraphOptionalClaims {
	return v.value
}

func (v *NullableMicrosoftGraphOptionalClaims) Set(val *MicrosoftGraphOptionalClaims) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphOptionalClaims) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphOptionalClaims) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphOptionalClaims(val *MicrosoftGraphOptionalClaims) *NullableMicrosoftGraphOptionalClaims {
	return &NullableMicrosoftGraphOptionalClaims{value: val, isSet: true}
}

func (v NullableMicrosoftGraphOptionalClaims) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphOptionalClaims) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


