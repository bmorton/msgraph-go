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

// MicrosoftGraphTeamsApp struct for MicrosoftGraphTeamsApp
type MicrosoftGraphTeamsApp struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
	// The name of the catalog app provided by the app developer in the Microsoft Teams zip app package.
	DisplayName NullableString `json:"displayName,omitempty"`
	// The method of distribution for the app. Read-only.
	DistributionMethod AnyOfmicrosoftGraphTeamsAppDistributionMethod `json:"distributionMethod,omitempty"`
	// The ID of the catalog provided by the app developer in the Microsoft Teams zip app package.
	ExternalId NullableString `json:"externalId,omitempty"`
	// The details for each version of the app.
	AppDefinitions *[]MicrosoftGraphTeamsAppDefinition `json:"appDefinitions,omitempty"`
}

// NewMicrosoftGraphTeamsApp instantiates a new MicrosoftGraphTeamsApp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphTeamsApp() *MicrosoftGraphTeamsApp {
	this := MicrosoftGraphTeamsApp{}
	return &this
}

// NewMicrosoftGraphTeamsAppWithDefaults instantiates a new MicrosoftGraphTeamsApp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphTeamsAppWithDefaults() *MicrosoftGraphTeamsApp {
	this := MicrosoftGraphTeamsApp{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MicrosoftGraphTeamsApp) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphTeamsApp) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphTeamsApp) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MicrosoftGraphTeamsApp) SetId(v string) {
	o.Id = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphTeamsApp) GetDisplayName() string {
	if o == nil || o.DisplayName.Get() == nil {
		var ret string
		return ret
	}
	return *o.DisplayName.Get()
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphTeamsApp) GetDisplayNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DisplayName.Get(), o.DisplayName.IsSet()
}

// HasDisplayName returns a boolean if a field has been set.
func (o *MicrosoftGraphTeamsApp) HasDisplayName() bool {
	if o != nil && o.DisplayName.IsSet() {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given NullableString and assigns it to the DisplayName field.
func (o *MicrosoftGraphTeamsApp) SetDisplayName(v string) {
	o.DisplayName.Set(&v)
}
// SetDisplayNameNil sets the value for DisplayName to be an explicit nil
func (o *MicrosoftGraphTeamsApp) SetDisplayNameNil() {
	o.DisplayName.Set(nil)
}

// UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
func (o *MicrosoftGraphTeamsApp) UnsetDisplayName() {
	o.DisplayName.Unset()
}

// GetDistributionMethod returns the DistributionMethod field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphTeamsApp) GetDistributionMethod() AnyOfmicrosoftGraphTeamsAppDistributionMethod {
	if o == nil  {
		var ret AnyOfmicrosoftGraphTeamsAppDistributionMethod
		return ret
	}
	return o.DistributionMethod
}

// GetDistributionMethodOk returns a tuple with the DistributionMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphTeamsApp) GetDistributionMethodOk() (*AnyOfmicrosoftGraphTeamsAppDistributionMethod, bool) {
	if o == nil || o.DistributionMethod == nil {
		return nil, false
	}
	return &o.DistributionMethod, true
}

// HasDistributionMethod returns a boolean if a field has been set.
func (o *MicrosoftGraphTeamsApp) HasDistributionMethod() bool {
	if o != nil && o.DistributionMethod != nil {
		return true
	}

	return false
}

// SetDistributionMethod gets a reference to the given AnyOfmicrosoftGraphTeamsAppDistributionMethod and assigns it to the DistributionMethod field.
func (o *MicrosoftGraphTeamsApp) SetDistributionMethod(v AnyOfmicrosoftGraphTeamsAppDistributionMethod) {
	o.DistributionMethod = v
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphTeamsApp) GetExternalId() string {
	if o == nil || o.ExternalId.Get() == nil {
		var ret string
		return ret
	}
	return *o.ExternalId.Get()
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphTeamsApp) GetExternalIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ExternalId.Get(), o.ExternalId.IsSet()
}

// HasExternalId returns a boolean if a field has been set.
func (o *MicrosoftGraphTeamsApp) HasExternalId() bool {
	if o != nil && o.ExternalId.IsSet() {
		return true
	}

	return false
}

// SetExternalId gets a reference to the given NullableString and assigns it to the ExternalId field.
func (o *MicrosoftGraphTeamsApp) SetExternalId(v string) {
	o.ExternalId.Set(&v)
}
// SetExternalIdNil sets the value for ExternalId to be an explicit nil
func (o *MicrosoftGraphTeamsApp) SetExternalIdNil() {
	o.ExternalId.Set(nil)
}

// UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
func (o *MicrosoftGraphTeamsApp) UnsetExternalId() {
	o.ExternalId.Unset()
}

// GetAppDefinitions returns the AppDefinitions field value if set, zero value otherwise.
func (o *MicrosoftGraphTeamsApp) GetAppDefinitions() []MicrosoftGraphTeamsAppDefinition {
	if o == nil || o.AppDefinitions == nil {
		var ret []MicrosoftGraphTeamsAppDefinition
		return ret
	}
	return *o.AppDefinitions
}

// GetAppDefinitionsOk returns a tuple with the AppDefinitions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphTeamsApp) GetAppDefinitionsOk() (*[]MicrosoftGraphTeamsAppDefinition, bool) {
	if o == nil || o.AppDefinitions == nil {
		return nil, false
	}
	return o.AppDefinitions, true
}

// HasAppDefinitions returns a boolean if a field has been set.
func (o *MicrosoftGraphTeamsApp) HasAppDefinitions() bool {
	if o != nil && o.AppDefinitions != nil {
		return true
	}

	return false
}

// SetAppDefinitions gets a reference to the given []MicrosoftGraphTeamsAppDefinition and assigns it to the AppDefinitions field.
func (o *MicrosoftGraphTeamsApp) SetAppDefinitions(v []MicrosoftGraphTeamsAppDefinition) {
	o.AppDefinitions = &v
}

func (o MicrosoftGraphTeamsApp) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.DisplayName.IsSet() {
		toSerialize["displayName"] = o.DisplayName.Get()
	}
	if o.DistributionMethod != nil {
		toSerialize["distributionMethod"] = o.DistributionMethod
	}
	if o.ExternalId.IsSet() {
		toSerialize["externalId"] = o.ExternalId.Get()
	}
	if o.AppDefinitions != nil {
		toSerialize["appDefinitions"] = o.AppDefinitions
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphTeamsApp struct {
	value *MicrosoftGraphTeamsApp
	isSet bool
}

func (v NullableMicrosoftGraphTeamsApp) Get() *MicrosoftGraphTeamsApp {
	return v.value
}

func (v *NullableMicrosoftGraphTeamsApp) Set(val *MicrosoftGraphTeamsApp) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphTeamsApp) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphTeamsApp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphTeamsApp(val *MicrosoftGraphTeamsApp) *NullableMicrosoftGraphTeamsApp {
	return &NullableMicrosoftGraphTeamsApp{value: val, isSet: true}
}

func (v NullableMicrosoftGraphTeamsApp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphTeamsApp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


