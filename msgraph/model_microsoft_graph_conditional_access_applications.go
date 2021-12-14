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

// MicrosoftGraphConditionalAccessApplications struct for MicrosoftGraphConditionalAccessApplications
type MicrosoftGraphConditionalAccessApplications struct {
	// The list of application IDs explicitly excluded from the policy.
	ExcludeApplications *[]string `json:"excludeApplications,omitempty"`
	// The list of application IDs the policy applies to, unless explicitly excluded (in excludeApplications). Can also be set to All.
	IncludeApplications *[]string `json:"includeApplications,omitempty"`
	// Authentication context class references include. Supported values are c1 through c25.
	IncludeAuthenticationContextClassReferences *[]string `json:"includeAuthenticationContextClassReferences,omitempty"`
	// User actions to include. Supported values are urn:user:registersecurityinfo and urn:user:registerdevice
	IncludeUserActions *[]string `json:"includeUserActions,omitempty"`
}

// NewMicrosoftGraphConditionalAccessApplications instantiates a new MicrosoftGraphConditionalAccessApplications object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphConditionalAccessApplications() *MicrosoftGraphConditionalAccessApplications {
	this := MicrosoftGraphConditionalAccessApplications{}
	return &this
}

// NewMicrosoftGraphConditionalAccessApplicationsWithDefaults instantiates a new MicrosoftGraphConditionalAccessApplications object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphConditionalAccessApplicationsWithDefaults() *MicrosoftGraphConditionalAccessApplications {
	this := MicrosoftGraphConditionalAccessApplications{}
	return &this
}

// GetExcludeApplications returns the ExcludeApplications field value if set, zero value otherwise.
func (o *MicrosoftGraphConditionalAccessApplications) GetExcludeApplications() []string {
	if o == nil || o.ExcludeApplications == nil {
		var ret []string
		return ret
	}
	return *o.ExcludeApplications
}

// GetExcludeApplicationsOk returns a tuple with the ExcludeApplications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphConditionalAccessApplications) GetExcludeApplicationsOk() (*[]string, bool) {
	if o == nil || o.ExcludeApplications == nil {
		return nil, false
	}
	return o.ExcludeApplications, true
}

// HasExcludeApplications returns a boolean if a field has been set.
func (o *MicrosoftGraphConditionalAccessApplications) HasExcludeApplications() bool {
	if o != nil && o.ExcludeApplications != nil {
		return true
	}

	return false
}

// SetExcludeApplications gets a reference to the given []string and assigns it to the ExcludeApplications field.
func (o *MicrosoftGraphConditionalAccessApplications) SetExcludeApplications(v []string) {
	o.ExcludeApplications = &v
}

// GetIncludeApplications returns the IncludeApplications field value if set, zero value otherwise.
func (o *MicrosoftGraphConditionalAccessApplications) GetIncludeApplications() []string {
	if o == nil || o.IncludeApplications == nil {
		var ret []string
		return ret
	}
	return *o.IncludeApplications
}

// GetIncludeApplicationsOk returns a tuple with the IncludeApplications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphConditionalAccessApplications) GetIncludeApplicationsOk() (*[]string, bool) {
	if o == nil || o.IncludeApplications == nil {
		return nil, false
	}
	return o.IncludeApplications, true
}

// HasIncludeApplications returns a boolean if a field has been set.
func (o *MicrosoftGraphConditionalAccessApplications) HasIncludeApplications() bool {
	if o != nil && o.IncludeApplications != nil {
		return true
	}

	return false
}

// SetIncludeApplications gets a reference to the given []string and assigns it to the IncludeApplications field.
func (o *MicrosoftGraphConditionalAccessApplications) SetIncludeApplications(v []string) {
	o.IncludeApplications = &v
}

// GetIncludeAuthenticationContextClassReferences returns the IncludeAuthenticationContextClassReferences field value if set, zero value otherwise.
func (o *MicrosoftGraphConditionalAccessApplications) GetIncludeAuthenticationContextClassReferences() []string {
	if o == nil || o.IncludeAuthenticationContextClassReferences == nil {
		var ret []string
		return ret
	}
	return *o.IncludeAuthenticationContextClassReferences
}

// GetIncludeAuthenticationContextClassReferencesOk returns a tuple with the IncludeAuthenticationContextClassReferences field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphConditionalAccessApplications) GetIncludeAuthenticationContextClassReferencesOk() (*[]string, bool) {
	if o == nil || o.IncludeAuthenticationContextClassReferences == nil {
		return nil, false
	}
	return o.IncludeAuthenticationContextClassReferences, true
}

// HasIncludeAuthenticationContextClassReferences returns a boolean if a field has been set.
func (o *MicrosoftGraphConditionalAccessApplications) HasIncludeAuthenticationContextClassReferences() bool {
	if o != nil && o.IncludeAuthenticationContextClassReferences != nil {
		return true
	}

	return false
}

// SetIncludeAuthenticationContextClassReferences gets a reference to the given []string and assigns it to the IncludeAuthenticationContextClassReferences field.
func (o *MicrosoftGraphConditionalAccessApplications) SetIncludeAuthenticationContextClassReferences(v []string) {
	o.IncludeAuthenticationContextClassReferences = &v
}

// GetIncludeUserActions returns the IncludeUserActions field value if set, zero value otherwise.
func (o *MicrosoftGraphConditionalAccessApplications) GetIncludeUserActions() []string {
	if o == nil || o.IncludeUserActions == nil {
		var ret []string
		return ret
	}
	return *o.IncludeUserActions
}

// GetIncludeUserActionsOk returns a tuple with the IncludeUserActions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphConditionalAccessApplications) GetIncludeUserActionsOk() (*[]string, bool) {
	if o == nil || o.IncludeUserActions == nil {
		return nil, false
	}
	return o.IncludeUserActions, true
}

// HasIncludeUserActions returns a boolean if a field has been set.
func (o *MicrosoftGraphConditionalAccessApplications) HasIncludeUserActions() bool {
	if o != nil && o.IncludeUserActions != nil {
		return true
	}

	return false
}

// SetIncludeUserActions gets a reference to the given []string and assigns it to the IncludeUserActions field.
func (o *MicrosoftGraphConditionalAccessApplications) SetIncludeUserActions(v []string) {
	o.IncludeUserActions = &v
}

func (o MicrosoftGraphConditionalAccessApplications) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ExcludeApplications != nil {
		toSerialize["excludeApplications"] = o.ExcludeApplications
	}
	if o.IncludeApplications != nil {
		toSerialize["includeApplications"] = o.IncludeApplications
	}
	if o.IncludeAuthenticationContextClassReferences != nil {
		toSerialize["includeAuthenticationContextClassReferences"] = o.IncludeAuthenticationContextClassReferences
	}
	if o.IncludeUserActions != nil {
		toSerialize["includeUserActions"] = o.IncludeUserActions
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphConditionalAccessApplications struct {
	value *MicrosoftGraphConditionalAccessApplications
	isSet bool
}

func (v NullableMicrosoftGraphConditionalAccessApplications) Get() *MicrosoftGraphConditionalAccessApplications {
	return v.value
}

func (v *NullableMicrosoftGraphConditionalAccessApplications) Set(val *MicrosoftGraphConditionalAccessApplications) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphConditionalAccessApplications) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphConditionalAccessApplications) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphConditionalAccessApplications(val *MicrosoftGraphConditionalAccessApplications) *NullableMicrosoftGraphConditionalAccessApplications {
	return &NullableMicrosoftGraphConditionalAccessApplications{value: val, isSet: true}
}

func (v NullableMicrosoftGraphConditionalAccessApplications) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphConditionalAccessApplications) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


