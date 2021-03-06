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

// MicrosoftGraphAppCatalogs struct for MicrosoftGraphAppCatalogs
type MicrosoftGraphAppCatalogs struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
	TeamsApps *[]MicrosoftGraphTeamsApp `json:"teamsApps,omitempty"`
}

// NewMicrosoftGraphAppCatalogs instantiates a new MicrosoftGraphAppCatalogs object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphAppCatalogs() *MicrosoftGraphAppCatalogs {
	this := MicrosoftGraphAppCatalogs{}
	return &this
}

// NewMicrosoftGraphAppCatalogsWithDefaults instantiates a new MicrosoftGraphAppCatalogs object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphAppCatalogsWithDefaults() *MicrosoftGraphAppCatalogs {
	this := MicrosoftGraphAppCatalogs{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MicrosoftGraphAppCatalogs) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphAppCatalogs) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphAppCatalogs) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MicrosoftGraphAppCatalogs) SetId(v string) {
	o.Id = &v
}

// GetTeamsApps returns the TeamsApps field value if set, zero value otherwise.
func (o *MicrosoftGraphAppCatalogs) GetTeamsApps() []MicrosoftGraphTeamsApp {
	if o == nil || o.TeamsApps == nil {
		var ret []MicrosoftGraphTeamsApp
		return ret
	}
	return *o.TeamsApps
}

// GetTeamsAppsOk returns a tuple with the TeamsApps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphAppCatalogs) GetTeamsAppsOk() (*[]MicrosoftGraphTeamsApp, bool) {
	if o == nil || o.TeamsApps == nil {
		return nil, false
	}
	return o.TeamsApps, true
}

// HasTeamsApps returns a boolean if a field has been set.
func (o *MicrosoftGraphAppCatalogs) HasTeamsApps() bool {
	if o != nil && o.TeamsApps != nil {
		return true
	}

	return false
}

// SetTeamsApps gets a reference to the given []MicrosoftGraphTeamsApp and assigns it to the TeamsApps field.
func (o *MicrosoftGraphAppCatalogs) SetTeamsApps(v []MicrosoftGraphTeamsApp) {
	o.TeamsApps = &v
}

func (o MicrosoftGraphAppCatalogs) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.TeamsApps != nil {
		toSerialize["teamsApps"] = o.TeamsApps
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphAppCatalogs struct {
	value *MicrosoftGraphAppCatalogs
	isSet bool
}

func (v NullableMicrosoftGraphAppCatalogs) Get() *MicrosoftGraphAppCatalogs {
	return v.value
}

func (v *NullableMicrosoftGraphAppCatalogs) Set(val *MicrosoftGraphAppCatalogs) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphAppCatalogs) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphAppCatalogs) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphAppCatalogs(val *MicrosoftGraphAppCatalogs) *NullableMicrosoftGraphAppCatalogs {
	return &NullableMicrosoftGraphAppCatalogs{value: val, isSet: true}
}

func (v NullableMicrosoftGraphAppCatalogs) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphAppCatalogs) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


