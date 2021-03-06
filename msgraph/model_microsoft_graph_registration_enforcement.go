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

// MicrosoftGraphRegistrationEnforcement struct for MicrosoftGraphRegistrationEnforcement
type MicrosoftGraphRegistrationEnforcement struct {
	// Run campaigns to remind users to set up targeted authentication methods.
	AuthenticationMethodsRegistrationCampaign AnyOfmicrosoftGraphAuthenticationMethodsRegistrationCampaign `json:"authenticationMethodsRegistrationCampaign,omitempty"`
}

// NewMicrosoftGraphRegistrationEnforcement instantiates a new MicrosoftGraphRegistrationEnforcement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphRegistrationEnforcement() *MicrosoftGraphRegistrationEnforcement {
	this := MicrosoftGraphRegistrationEnforcement{}
	return &this
}

// NewMicrosoftGraphRegistrationEnforcementWithDefaults instantiates a new MicrosoftGraphRegistrationEnforcement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphRegistrationEnforcementWithDefaults() *MicrosoftGraphRegistrationEnforcement {
	this := MicrosoftGraphRegistrationEnforcement{}
	return &this
}

// GetAuthenticationMethodsRegistrationCampaign returns the AuthenticationMethodsRegistrationCampaign field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphRegistrationEnforcement) GetAuthenticationMethodsRegistrationCampaign() AnyOfmicrosoftGraphAuthenticationMethodsRegistrationCampaign {
	if o == nil  {
		var ret AnyOfmicrosoftGraphAuthenticationMethodsRegistrationCampaign
		return ret
	}
	return o.AuthenticationMethodsRegistrationCampaign
}

// GetAuthenticationMethodsRegistrationCampaignOk returns a tuple with the AuthenticationMethodsRegistrationCampaign field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphRegistrationEnforcement) GetAuthenticationMethodsRegistrationCampaignOk() (*AnyOfmicrosoftGraphAuthenticationMethodsRegistrationCampaign, bool) {
	if o == nil || o.AuthenticationMethodsRegistrationCampaign == nil {
		return nil, false
	}
	return &o.AuthenticationMethodsRegistrationCampaign, true
}

// HasAuthenticationMethodsRegistrationCampaign returns a boolean if a field has been set.
func (o *MicrosoftGraphRegistrationEnforcement) HasAuthenticationMethodsRegistrationCampaign() bool {
	if o != nil && o.AuthenticationMethodsRegistrationCampaign != nil {
		return true
	}

	return false
}

// SetAuthenticationMethodsRegistrationCampaign gets a reference to the given AnyOfmicrosoftGraphAuthenticationMethodsRegistrationCampaign and assigns it to the AuthenticationMethodsRegistrationCampaign field.
func (o *MicrosoftGraphRegistrationEnforcement) SetAuthenticationMethodsRegistrationCampaign(v AnyOfmicrosoftGraphAuthenticationMethodsRegistrationCampaign) {
	o.AuthenticationMethodsRegistrationCampaign = v
}

func (o MicrosoftGraphRegistrationEnforcement) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AuthenticationMethodsRegistrationCampaign != nil {
		toSerialize["authenticationMethodsRegistrationCampaign"] = o.AuthenticationMethodsRegistrationCampaign
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphRegistrationEnforcement struct {
	value *MicrosoftGraphRegistrationEnforcement
	isSet bool
}

func (v NullableMicrosoftGraphRegistrationEnforcement) Get() *MicrosoftGraphRegistrationEnforcement {
	return v.value
}

func (v *NullableMicrosoftGraphRegistrationEnforcement) Set(val *MicrosoftGraphRegistrationEnforcement) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphRegistrationEnforcement) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphRegistrationEnforcement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphRegistrationEnforcement(val *MicrosoftGraphRegistrationEnforcement) *NullableMicrosoftGraphRegistrationEnforcement {
	return &NullableMicrosoftGraphRegistrationEnforcement{value: val, isSet: true}
}

func (v NullableMicrosoftGraphRegistrationEnforcement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphRegistrationEnforcement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


