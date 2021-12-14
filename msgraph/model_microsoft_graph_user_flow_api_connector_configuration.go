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

// MicrosoftGraphUserFlowApiConnectorConfiguration struct for MicrosoftGraphUserFlowApiConnectorConfiguration
type MicrosoftGraphUserFlowApiConnectorConfiguration struct {
	PostAttributeCollection AnyOfmicrosoftGraphIdentityApiConnector `json:"postAttributeCollection,omitempty"`
	PostFederationSignup AnyOfmicrosoftGraphIdentityApiConnector `json:"postFederationSignup,omitempty"`
}

// NewMicrosoftGraphUserFlowApiConnectorConfiguration instantiates a new MicrosoftGraphUserFlowApiConnectorConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphUserFlowApiConnectorConfiguration() *MicrosoftGraphUserFlowApiConnectorConfiguration {
	this := MicrosoftGraphUserFlowApiConnectorConfiguration{}
	return &this
}

// NewMicrosoftGraphUserFlowApiConnectorConfigurationWithDefaults instantiates a new MicrosoftGraphUserFlowApiConnectorConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphUserFlowApiConnectorConfigurationWithDefaults() *MicrosoftGraphUserFlowApiConnectorConfiguration {
	this := MicrosoftGraphUserFlowApiConnectorConfiguration{}
	return &this
}

// GetPostAttributeCollection returns the PostAttributeCollection field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphUserFlowApiConnectorConfiguration) GetPostAttributeCollection() AnyOfmicrosoftGraphIdentityApiConnector {
	if o == nil  {
		var ret AnyOfmicrosoftGraphIdentityApiConnector
		return ret
	}
	return o.PostAttributeCollection
}

// GetPostAttributeCollectionOk returns a tuple with the PostAttributeCollection field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphUserFlowApiConnectorConfiguration) GetPostAttributeCollectionOk() (*AnyOfmicrosoftGraphIdentityApiConnector, bool) {
	if o == nil || o.PostAttributeCollection == nil {
		return nil, false
	}
	return &o.PostAttributeCollection, true
}

// HasPostAttributeCollection returns a boolean if a field has been set.
func (o *MicrosoftGraphUserFlowApiConnectorConfiguration) HasPostAttributeCollection() bool {
	if o != nil && o.PostAttributeCollection != nil {
		return true
	}

	return false
}

// SetPostAttributeCollection gets a reference to the given AnyOfmicrosoftGraphIdentityApiConnector and assigns it to the PostAttributeCollection field.
func (o *MicrosoftGraphUserFlowApiConnectorConfiguration) SetPostAttributeCollection(v AnyOfmicrosoftGraphIdentityApiConnector) {
	o.PostAttributeCollection = v
}

// GetPostFederationSignup returns the PostFederationSignup field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphUserFlowApiConnectorConfiguration) GetPostFederationSignup() AnyOfmicrosoftGraphIdentityApiConnector {
	if o == nil  {
		var ret AnyOfmicrosoftGraphIdentityApiConnector
		return ret
	}
	return o.PostFederationSignup
}

// GetPostFederationSignupOk returns a tuple with the PostFederationSignup field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphUserFlowApiConnectorConfiguration) GetPostFederationSignupOk() (*AnyOfmicrosoftGraphIdentityApiConnector, bool) {
	if o == nil || o.PostFederationSignup == nil {
		return nil, false
	}
	return &o.PostFederationSignup, true
}

// HasPostFederationSignup returns a boolean if a field has been set.
func (o *MicrosoftGraphUserFlowApiConnectorConfiguration) HasPostFederationSignup() bool {
	if o != nil && o.PostFederationSignup != nil {
		return true
	}

	return false
}

// SetPostFederationSignup gets a reference to the given AnyOfmicrosoftGraphIdentityApiConnector and assigns it to the PostFederationSignup field.
func (o *MicrosoftGraphUserFlowApiConnectorConfiguration) SetPostFederationSignup(v AnyOfmicrosoftGraphIdentityApiConnector) {
	o.PostFederationSignup = v
}

func (o MicrosoftGraphUserFlowApiConnectorConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PostAttributeCollection != nil {
		toSerialize["postAttributeCollection"] = o.PostAttributeCollection
	}
	if o.PostFederationSignup != nil {
		toSerialize["postFederationSignup"] = o.PostFederationSignup
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphUserFlowApiConnectorConfiguration struct {
	value *MicrosoftGraphUserFlowApiConnectorConfiguration
	isSet bool
}

func (v NullableMicrosoftGraphUserFlowApiConnectorConfiguration) Get() *MicrosoftGraphUserFlowApiConnectorConfiguration {
	return v.value
}

func (v *NullableMicrosoftGraphUserFlowApiConnectorConfiguration) Set(val *MicrosoftGraphUserFlowApiConnectorConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphUserFlowApiConnectorConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphUserFlowApiConnectorConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphUserFlowApiConnectorConfiguration(val *MicrosoftGraphUserFlowApiConnectorConfiguration) *NullableMicrosoftGraphUserFlowApiConnectorConfiguration {
	return &NullableMicrosoftGraphUserFlowApiConnectorConfiguration{value: val, isSet: true}
}

func (v NullableMicrosoftGraphUserFlowApiConnectorConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphUserFlowApiConnectorConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

