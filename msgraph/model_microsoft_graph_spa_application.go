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

// MicrosoftGraphSpaApplication struct for MicrosoftGraphSpaApplication
type MicrosoftGraphSpaApplication struct {
	// Specifies the URLs where user tokens are sent for sign-in, or the redirect URIs where OAuth 2.0 authorization codes and access tokens are sent.
	RedirectUris *[]string `json:"redirectUris,omitempty"`
}

// NewMicrosoftGraphSpaApplication instantiates a new MicrosoftGraphSpaApplication object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphSpaApplication() *MicrosoftGraphSpaApplication {
	this := MicrosoftGraphSpaApplication{}
	return &this
}

// NewMicrosoftGraphSpaApplicationWithDefaults instantiates a new MicrosoftGraphSpaApplication object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphSpaApplicationWithDefaults() *MicrosoftGraphSpaApplication {
	this := MicrosoftGraphSpaApplication{}
	return &this
}

// GetRedirectUris returns the RedirectUris field value if set, zero value otherwise.
func (o *MicrosoftGraphSpaApplication) GetRedirectUris() []string {
	if o == nil || o.RedirectUris == nil {
		var ret []string
		return ret
	}
	return *o.RedirectUris
}

// GetRedirectUrisOk returns a tuple with the RedirectUris field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphSpaApplication) GetRedirectUrisOk() (*[]string, bool) {
	if o == nil || o.RedirectUris == nil {
		return nil, false
	}
	return o.RedirectUris, true
}

// HasRedirectUris returns a boolean if a field has been set.
func (o *MicrosoftGraphSpaApplication) HasRedirectUris() bool {
	if o != nil && o.RedirectUris != nil {
		return true
	}

	return false
}

// SetRedirectUris gets a reference to the given []string and assigns it to the RedirectUris field.
func (o *MicrosoftGraphSpaApplication) SetRedirectUris(v []string) {
	o.RedirectUris = &v
}

func (o MicrosoftGraphSpaApplication) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.RedirectUris != nil {
		toSerialize["redirectUris"] = o.RedirectUris
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphSpaApplication struct {
	value *MicrosoftGraphSpaApplication
	isSet bool
}

func (v NullableMicrosoftGraphSpaApplication) Get() *MicrosoftGraphSpaApplication {
	return v.value
}

func (v *NullableMicrosoftGraphSpaApplication) Set(val *MicrosoftGraphSpaApplication) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphSpaApplication) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphSpaApplication) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphSpaApplication(val *MicrosoftGraphSpaApplication) *NullableMicrosoftGraphSpaApplication {
	return &NullableMicrosoftGraphSpaApplication{value: val, isSet: true}
}

func (v NullableMicrosoftGraphSpaApplication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphSpaApplication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


