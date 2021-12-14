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

// MicrosoftGraphIdentityApiConnector struct for MicrosoftGraphIdentityApiConnector
type MicrosoftGraphIdentityApiConnector struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
	// The object which describes the authentication configuration details for calling the API. Basic and PKCS 12 client certificate are supported.
	AuthenticationConfiguration AnyOfobject `json:"authenticationConfiguration,omitempty"`
	// The name of the API connector.
	DisplayName NullableString `json:"displayName,omitempty"`
	// The URL of the API endpoint to call.
	TargetUrl NullableString `json:"targetUrl,omitempty"`
}

// NewMicrosoftGraphIdentityApiConnector instantiates a new MicrosoftGraphIdentityApiConnector object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphIdentityApiConnector() *MicrosoftGraphIdentityApiConnector {
	this := MicrosoftGraphIdentityApiConnector{}
	return &this
}

// NewMicrosoftGraphIdentityApiConnectorWithDefaults instantiates a new MicrosoftGraphIdentityApiConnector object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphIdentityApiConnectorWithDefaults() *MicrosoftGraphIdentityApiConnector {
	this := MicrosoftGraphIdentityApiConnector{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MicrosoftGraphIdentityApiConnector) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphIdentityApiConnector) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphIdentityApiConnector) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MicrosoftGraphIdentityApiConnector) SetId(v string) {
	o.Id = &v
}

// GetAuthenticationConfiguration returns the AuthenticationConfiguration field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphIdentityApiConnector) GetAuthenticationConfiguration() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.AuthenticationConfiguration
}

// GetAuthenticationConfigurationOk returns a tuple with the AuthenticationConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphIdentityApiConnector) GetAuthenticationConfigurationOk() (*AnyOfobject, bool) {
	if o == nil || o.AuthenticationConfiguration == nil {
		return nil, false
	}
	return &o.AuthenticationConfiguration, true
}

// HasAuthenticationConfiguration returns a boolean if a field has been set.
func (o *MicrosoftGraphIdentityApiConnector) HasAuthenticationConfiguration() bool {
	if o != nil && o.AuthenticationConfiguration != nil {
		return true
	}

	return false
}

// SetAuthenticationConfiguration gets a reference to the given AnyOfobject and assigns it to the AuthenticationConfiguration field.
func (o *MicrosoftGraphIdentityApiConnector) SetAuthenticationConfiguration(v AnyOfobject) {
	o.AuthenticationConfiguration = v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphIdentityApiConnector) GetDisplayName() string {
	if o == nil || o.DisplayName.Get() == nil {
		var ret string
		return ret
	}
	return *o.DisplayName.Get()
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphIdentityApiConnector) GetDisplayNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DisplayName.Get(), o.DisplayName.IsSet()
}

// HasDisplayName returns a boolean if a field has been set.
func (o *MicrosoftGraphIdentityApiConnector) HasDisplayName() bool {
	if o != nil && o.DisplayName.IsSet() {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given NullableString and assigns it to the DisplayName field.
func (o *MicrosoftGraphIdentityApiConnector) SetDisplayName(v string) {
	o.DisplayName.Set(&v)
}
// SetDisplayNameNil sets the value for DisplayName to be an explicit nil
func (o *MicrosoftGraphIdentityApiConnector) SetDisplayNameNil() {
	o.DisplayName.Set(nil)
}

// UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
func (o *MicrosoftGraphIdentityApiConnector) UnsetDisplayName() {
	o.DisplayName.Unset()
}

// GetTargetUrl returns the TargetUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphIdentityApiConnector) GetTargetUrl() string {
	if o == nil || o.TargetUrl.Get() == nil {
		var ret string
		return ret
	}
	return *o.TargetUrl.Get()
}

// GetTargetUrlOk returns a tuple with the TargetUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphIdentityApiConnector) GetTargetUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TargetUrl.Get(), o.TargetUrl.IsSet()
}

// HasTargetUrl returns a boolean if a field has been set.
func (o *MicrosoftGraphIdentityApiConnector) HasTargetUrl() bool {
	if o != nil && o.TargetUrl.IsSet() {
		return true
	}

	return false
}

// SetTargetUrl gets a reference to the given NullableString and assigns it to the TargetUrl field.
func (o *MicrosoftGraphIdentityApiConnector) SetTargetUrl(v string) {
	o.TargetUrl.Set(&v)
}
// SetTargetUrlNil sets the value for TargetUrl to be an explicit nil
func (o *MicrosoftGraphIdentityApiConnector) SetTargetUrlNil() {
	o.TargetUrl.Set(nil)
}

// UnsetTargetUrl ensures that no value is present for TargetUrl, not even an explicit nil
func (o *MicrosoftGraphIdentityApiConnector) UnsetTargetUrl() {
	o.TargetUrl.Unset()
}

func (o MicrosoftGraphIdentityApiConnector) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.AuthenticationConfiguration != nil {
		toSerialize["authenticationConfiguration"] = o.AuthenticationConfiguration
	}
	if o.DisplayName.IsSet() {
		toSerialize["displayName"] = o.DisplayName.Get()
	}
	if o.TargetUrl.IsSet() {
		toSerialize["targetUrl"] = o.TargetUrl.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphIdentityApiConnector struct {
	value *MicrosoftGraphIdentityApiConnector
	isSet bool
}

func (v NullableMicrosoftGraphIdentityApiConnector) Get() *MicrosoftGraphIdentityApiConnector {
	return v.value
}

func (v *NullableMicrosoftGraphIdentityApiConnector) Set(val *MicrosoftGraphIdentityApiConnector) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphIdentityApiConnector) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphIdentityApiConnector) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphIdentityApiConnector(val *MicrosoftGraphIdentityApiConnector) *NullableMicrosoftGraphIdentityApiConnector {
	return &NullableMicrosoftGraphIdentityApiConnector{value: val, isSet: true}
}

func (v NullableMicrosoftGraphIdentityApiConnector) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphIdentityApiConnector) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


