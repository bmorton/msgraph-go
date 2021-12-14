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

// MicrosoftGraphConditionalAccessLocations struct for MicrosoftGraphConditionalAccessLocations
type MicrosoftGraphConditionalAccessLocations struct {
	// Location IDs excluded from scope of policy.
	ExcludeLocations *[]string `json:"excludeLocations,omitempty"`
	// Location IDs in scope of policy unless explicitly excluded, All, or AllTrusted.
	IncludeLocations *[]string `json:"includeLocations,omitempty"`
}

// NewMicrosoftGraphConditionalAccessLocations instantiates a new MicrosoftGraphConditionalAccessLocations object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphConditionalAccessLocations() *MicrosoftGraphConditionalAccessLocations {
	this := MicrosoftGraphConditionalAccessLocations{}
	return &this
}

// NewMicrosoftGraphConditionalAccessLocationsWithDefaults instantiates a new MicrosoftGraphConditionalAccessLocations object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphConditionalAccessLocationsWithDefaults() *MicrosoftGraphConditionalAccessLocations {
	this := MicrosoftGraphConditionalAccessLocations{}
	return &this
}

// GetExcludeLocations returns the ExcludeLocations field value if set, zero value otherwise.
func (o *MicrosoftGraphConditionalAccessLocations) GetExcludeLocations() []string {
	if o == nil || o.ExcludeLocations == nil {
		var ret []string
		return ret
	}
	return *o.ExcludeLocations
}

// GetExcludeLocationsOk returns a tuple with the ExcludeLocations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphConditionalAccessLocations) GetExcludeLocationsOk() (*[]string, bool) {
	if o == nil || o.ExcludeLocations == nil {
		return nil, false
	}
	return o.ExcludeLocations, true
}

// HasExcludeLocations returns a boolean if a field has been set.
func (o *MicrosoftGraphConditionalAccessLocations) HasExcludeLocations() bool {
	if o != nil && o.ExcludeLocations != nil {
		return true
	}

	return false
}

// SetExcludeLocations gets a reference to the given []string and assigns it to the ExcludeLocations field.
func (o *MicrosoftGraphConditionalAccessLocations) SetExcludeLocations(v []string) {
	o.ExcludeLocations = &v
}

// GetIncludeLocations returns the IncludeLocations field value if set, zero value otherwise.
func (o *MicrosoftGraphConditionalAccessLocations) GetIncludeLocations() []string {
	if o == nil || o.IncludeLocations == nil {
		var ret []string
		return ret
	}
	return *o.IncludeLocations
}

// GetIncludeLocationsOk returns a tuple with the IncludeLocations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphConditionalAccessLocations) GetIncludeLocationsOk() (*[]string, bool) {
	if o == nil || o.IncludeLocations == nil {
		return nil, false
	}
	return o.IncludeLocations, true
}

// HasIncludeLocations returns a boolean if a field has been set.
func (o *MicrosoftGraphConditionalAccessLocations) HasIncludeLocations() bool {
	if o != nil && o.IncludeLocations != nil {
		return true
	}

	return false
}

// SetIncludeLocations gets a reference to the given []string and assigns it to the IncludeLocations field.
func (o *MicrosoftGraphConditionalAccessLocations) SetIncludeLocations(v []string) {
	o.IncludeLocations = &v
}

func (o MicrosoftGraphConditionalAccessLocations) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ExcludeLocations != nil {
		toSerialize["excludeLocations"] = o.ExcludeLocations
	}
	if o.IncludeLocations != nil {
		toSerialize["includeLocations"] = o.IncludeLocations
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphConditionalAccessLocations struct {
	value *MicrosoftGraphConditionalAccessLocations
	isSet bool
}

func (v NullableMicrosoftGraphConditionalAccessLocations) Get() *MicrosoftGraphConditionalAccessLocations {
	return v.value
}

func (v *NullableMicrosoftGraphConditionalAccessLocations) Set(val *MicrosoftGraphConditionalAccessLocations) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphConditionalAccessLocations) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphConditionalAccessLocations) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphConditionalAccessLocations(val *MicrosoftGraphConditionalAccessLocations) *NullableMicrosoftGraphConditionalAccessLocations {
	return &NullableMicrosoftGraphConditionalAccessLocations{value: val, isSet: true}
}

func (v NullableMicrosoftGraphConditionalAccessLocations) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphConditionalAccessLocations) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


