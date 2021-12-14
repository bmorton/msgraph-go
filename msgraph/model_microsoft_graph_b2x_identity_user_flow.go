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

// MicrosoftGraphB2xIdentityUserFlow struct for MicrosoftGraphB2xIdentityUserFlow
type MicrosoftGraphB2xIdentityUserFlow struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
	UserFlowType AnyOfmicrosoftGraphUserFlowType `json:"userFlowType,omitempty"`
	UserFlowTypeVersion AnyOfnumberstringstring `json:"userFlowTypeVersion,omitempty"`
	// Configuration for enabling an API connector for use as part of the self-service sign-up user flow. You can only obtain the value of this object using Get userFlowApiConnectorConfiguration.
	ApiConnectorConfiguration AnyOfmicrosoftGraphUserFlowApiConnectorConfiguration `json:"apiConnectorConfiguration,omitempty"`
	// The identity providers included in the user flow.
	IdentityProviders *[]MicrosoftGraphIdentityProvider `json:"identityProviders,omitempty"`
	// The languages supported for customization within the user flow. Language customization is enabled by default in self-service sign-up user flow. You cannot create custom languages in self-service sign-up user flows.
	Languages *[]MicrosoftGraphUserFlowLanguageConfiguration `json:"languages,omitempty"`
	// The user attribute assignments included in the user flow.
	UserAttributeAssignments *[]MicrosoftGraphIdentityUserFlowAttributeAssignment `json:"userAttributeAssignments,omitempty"`
	UserFlowIdentityProviders *[]MicrosoftGraphIdentityProviderBase `json:"userFlowIdentityProviders,omitempty"`
}

// NewMicrosoftGraphB2xIdentityUserFlow instantiates a new MicrosoftGraphB2xIdentityUserFlow object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphB2xIdentityUserFlow() *MicrosoftGraphB2xIdentityUserFlow {
	this := MicrosoftGraphB2xIdentityUserFlow{}
	return &this
}

// NewMicrosoftGraphB2xIdentityUserFlowWithDefaults instantiates a new MicrosoftGraphB2xIdentityUserFlow object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphB2xIdentityUserFlowWithDefaults() *MicrosoftGraphB2xIdentityUserFlow {
	this := MicrosoftGraphB2xIdentityUserFlow{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MicrosoftGraphB2xIdentityUserFlow) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphB2xIdentityUserFlow) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphB2xIdentityUserFlow) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MicrosoftGraphB2xIdentityUserFlow) SetId(v string) {
	o.Id = &v
}

// GetUserFlowType returns the UserFlowType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphB2xIdentityUserFlow) GetUserFlowType() AnyOfmicrosoftGraphUserFlowType {
	if o == nil  {
		var ret AnyOfmicrosoftGraphUserFlowType
		return ret
	}
	return o.UserFlowType
}

// GetUserFlowTypeOk returns a tuple with the UserFlowType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphB2xIdentityUserFlow) GetUserFlowTypeOk() (*AnyOfmicrosoftGraphUserFlowType, bool) {
	if o == nil || o.UserFlowType == nil {
		return nil, false
	}
	return &o.UserFlowType, true
}

// HasUserFlowType returns a boolean if a field has been set.
func (o *MicrosoftGraphB2xIdentityUserFlow) HasUserFlowType() bool {
	if o != nil && o.UserFlowType != nil {
		return true
	}

	return false
}

// SetUserFlowType gets a reference to the given AnyOfmicrosoftGraphUserFlowType and assigns it to the UserFlowType field.
func (o *MicrosoftGraphB2xIdentityUserFlow) SetUserFlowType(v AnyOfmicrosoftGraphUserFlowType) {
	o.UserFlowType = v
}

// GetUserFlowTypeVersion returns the UserFlowTypeVersion field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphB2xIdentityUserFlow) GetUserFlowTypeVersion() AnyOfnumberstringstring {
	if o == nil  {
		var ret AnyOfnumberstringstring
		return ret
	}
	return o.UserFlowTypeVersion
}

// GetUserFlowTypeVersionOk returns a tuple with the UserFlowTypeVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphB2xIdentityUserFlow) GetUserFlowTypeVersionOk() (*AnyOfnumberstringstring, bool) {
	if o == nil || o.UserFlowTypeVersion == nil {
		return nil, false
	}
	return &o.UserFlowTypeVersion, true
}

// HasUserFlowTypeVersion returns a boolean if a field has been set.
func (o *MicrosoftGraphB2xIdentityUserFlow) HasUserFlowTypeVersion() bool {
	if o != nil && o.UserFlowTypeVersion != nil {
		return true
	}

	return false
}

// SetUserFlowTypeVersion gets a reference to the given AnyOfnumberstringstring and assigns it to the UserFlowTypeVersion field.
func (o *MicrosoftGraphB2xIdentityUserFlow) SetUserFlowTypeVersion(v AnyOfnumberstringstring) {
	o.UserFlowTypeVersion = v
}

// GetApiConnectorConfiguration returns the ApiConnectorConfiguration field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphB2xIdentityUserFlow) GetApiConnectorConfiguration() AnyOfmicrosoftGraphUserFlowApiConnectorConfiguration {
	if o == nil  {
		var ret AnyOfmicrosoftGraphUserFlowApiConnectorConfiguration
		return ret
	}
	return o.ApiConnectorConfiguration
}

// GetApiConnectorConfigurationOk returns a tuple with the ApiConnectorConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphB2xIdentityUserFlow) GetApiConnectorConfigurationOk() (*AnyOfmicrosoftGraphUserFlowApiConnectorConfiguration, bool) {
	if o == nil || o.ApiConnectorConfiguration == nil {
		return nil, false
	}
	return &o.ApiConnectorConfiguration, true
}

// HasApiConnectorConfiguration returns a boolean if a field has been set.
func (o *MicrosoftGraphB2xIdentityUserFlow) HasApiConnectorConfiguration() bool {
	if o != nil && o.ApiConnectorConfiguration != nil {
		return true
	}

	return false
}

// SetApiConnectorConfiguration gets a reference to the given AnyOfmicrosoftGraphUserFlowApiConnectorConfiguration and assigns it to the ApiConnectorConfiguration field.
func (o *MicrosoftGraphB2xIdentityUserFlow) SetApiConnectorConfiguration(v AnyOfmicrosoftGraphUserFlowApiConnectorConfiguration) {
	o.ApiConnectorConfiguration = v
}

// GetIdentityProviders returns the IdentityProviders field value if set, zero value otherwise.
func (o *MicrosoftGraphB2xIdentityUserFlow) GetIdentityProviders() []MicrosoftGraphIdentityProvider {
	if o == nil || o.IdentityProviders == nil {
		var ret []MicrosoftGraphIdentityProvider
		return ret
	}
	return *o.IdentityProviders
}

// GetIdentityProvidersOk returns a tuple with the IdentityProviders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphB2xIdentityUserFlow) GetIdentityProvidersOk() (*[]MicrosoftGraphIdentityProvider, bool) {
	if o == nil || o.IdentityProviders == nil {
		return nil, false
	}
	return o.IdentityProviders, true
}

// HasIdentityProviders returns a boolean if a field has been set.
func (o *MicrosoftGraphB2xIdentityUserFlow) HasIdentityProviders() bool {
	if o != nil && o.IdentityProviders != nil {
		return true
	}

	return false
}

// SetIdentityProviders gets a reference to the given []MicrosoftGraphIdentityProvider and assigns it to the IdentityProviders field.
func (o *MicrosoftGraphB2xIdentityUserFlow) SetIdentityProviders(v []MicrosoftGraphIdentityProvider) {
	o.IdentityProviders = &v
}

// GetLanguages returns the Languages field value if set, zero value otherwise.
func (o *MicrosoftGraphB2xIdentityUserFlow) GetLanguages() []MicrosoftGraphUserFlowLanguageConfiguration {
	if o == nil || o.Languages == nil {
		var ret []MicrosoftGraphUserFlowLanguageConfiguration
		return ret
	}
	return *o.Languages
}

// GetLanguagesOk returns a tuple with the Languages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphB2xIdentityUserFlow) GetLanguagesOk() (*[]MicrosoftGraphUserFlowLanguageConfiguration, bool) {
	if o == nil || o.Languages == nil {
		return nil, false
	}
	return o.Languages, true
}

// HasLanguages returns a boolean if a field has been set.
func (o *MicrosoftGraphB2xIdentityUserFlow) HasLanguages() bool {
	if o != nil && o.Languages != nil {
		return true
	}

	return false
}

// SetLanguages gets a reference to the given []MicrosoftGraphUserFlowLanguageConfiguration and assigns it to the Languages field.
func (o *MicrosoftGraphB2xIdentityUserFlow) SetLanguages(v []MicrosoftGraphUserFlowLanguageConfiguration) {
	o.Languages = &v
}

// GetUserAttributeAssignments returns the UserAttributeAssignments field value if set, zero value otherwise.
func (o *MicrosoftGraphB2xIdentityUserFlow) GetUserAttributeAssignments() []MicrosoftGraphIdentityUserFlowAttributeAssignment {
	if o == nil || o.UserAttributeAssignments == nil {
		var ret []MicrosoftGraphIdentityUserFlowAttributeAssignment
		return ret
	}
	return *o.UserAttributeAssignments
}

// GetUserAttributeAssignmentsOk returns a tuple with the UserAttributeAssignments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphB2xIdentityUserFlow) GetUserAttributeAssignmentsOk() (*[]MicrosoftGraphIdentityUserFlowAttributeAssignment, bool) {
	if o == nil || o.UserAttributeAssignments == nil {
		return nil, false
	}
	return o.UserAttributeAssignments, true
}

// HasUserAttributeAssignments returns a boolean if a field has been set.
func (o *MicrosoftGraphB2xIdentityUserFlow) HasUserAttributeAssignments() bool {
	if o != nil && o.UserAttributeAssignments != nil {
		return true
	}

	return false
}

// SetUserAttributeAssignments gets a reference to the given []MicrosoftGraphIdentityUserFlowAttributeAssignment and assigns it to the UserAttributeAssignments field.
func (o *MicrosoftGraphB2xIdentityUserFlow) SetUserAttributeAssignments(v []MicrosoftGraphIdentityUserFlowAttributeAssignment) {
	o.UserAttributeAssignments = &v
}

// GetUserFlowIdentityProviders returns the UserFlowIdentityProviders field value if set, zero value otherwise.
func (o *MicrosoftGraphB2xIdentityUserFlow) GetUserFlowIdentityProviders() []MicrosoftGraphIdentityProviderBase {
	if o == nil || o.UserFlowIdentityProviders == nil {
		var ret []MicrosoftGraphIdentityProviderBase
		return ret
	}
	return *o.UserFlowIdentityProviders
}

// GetUserFlowIdentityProvidersOk returns a tuple with the UserFlowIdentityProviders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphB2xIdentityUserFlow) GetUserFlowIdentityProvidersOk() (*[]MicrosoftGraphIdentityProviderBase, bool) {
	if o == nil || o.UserFlowIdentityProviders == nil {
		return nil, false
	}
	return o.UserFlowIdentityProviders, true
}

// HasUserFlowIdentityProviders returns a boolean if a field has been set.
func (o *MicrosoftGraphB2xIdentityUserFlow) HasUserFlowIdentityProviders() bool {
	if o != nil && o.UserFlowIdentityProviders != nil {
		return true
	}

	return false
}

// SetUserFlowIdentityProviders gets a reference to the given []MicrosoftGraphIdentityProviderBase and assigns it to the UserFlowIdentityProviders field.
func (o *MicrosoftGraphB2xIdentityUserFlow) SetUserFlowIdentityProviders(v []MicrosoftGraphIdentityProviderBase) {
	o.UserFlowIdentityProviders = &v
}

func (o MicrosoftGraphB2xIdentityUserFlow) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.UserFlowType != nil {
		toSerialize["userFlowType"] = o.UserFlowType
	}
	if o.UserFlowTypeVersion != nil {
		toSerialize["userFlowTypeVersion"] = o.UserFlowTypeVersion
	}
	if o.ApiConnectorConfiguration != nil {
		toSerialize["apiConnectorConfiguration"] = o.ApiConnectorConfiguration
	}
	if o.IdentityProviders != nil {
		toSerialize["identityProviders"] = o.IdentityProviders
	}
	if o.Languages != nil {
		toSerialize["languages"] = o.Languages
	}
	if o.UserAttributeAssignments != nil {
		toSerialize["userAttributeAssignments"] = o.UserAttributeAssignments
	}
	if o.UserFlowIdentityProviders != nil {
		toSerialize["userFlowIdentityProviders"] = o.UserFlowIdentityProviders
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphB2xIdentityUserFlow struct {
	value *MicrosoftGraphB2xIdentityUserFlow
	isSet bool
}

func (v NullableMicrosoftGraphB2xIdentityUserFlow) Get() *MicrosoftGraphB2xIdentityUserFlow {
	return v.value
}

func (v *NullableMicrosoftGraphB2xIdentityUserFlow) Set(val *MicrosoftGraphB2xIdentityUserFlow) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphB2xIdentityUserFlow) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphB2xIdentityUserFlow) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphB2xIdentityUserFlow(val *MicrosoftGraphB2xIdentityUserFlow) *NullableMicrosoftGraphB2xIdentityUserFlow {
	return &NullableMicrosoftGraphB2xIdentityUserFlow{value: val, isSet: true}
}

func (v NullableMicrosoftGraphB2xIdentityUserFlow) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphB2xIdentityUserFlow) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


