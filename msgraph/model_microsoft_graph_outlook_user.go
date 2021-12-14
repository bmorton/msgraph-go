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

// MicrosoftGraphOutlookUser struct for MicrosoftGraphOutlookUser
type MicrosoftGraphOutlookUser struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
	// A list of categories defined for the user.
	MasterCategories *[]MicrosoftGraphOutlookCategory `json:"masterCategories,omitempty"`
}

// NewMicrosoftGraphOutlookUser instantiates a new MicrosoftGraphOutlookUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphOutlookUser() *MicrosoftGraphOutlookUser {
	this := MicrosoftGraphOutlookUser{}
	return &this
}

// NewMicrosoftGraphOutlookUserWithDefaults instantiates a new MicrosoftGraphOutlookUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphOutlookUserWithDefaults() *MicrosoftGraphOutlookUser {
	this := MicrosoftGraphOutlookUser{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MicrosoftGraphOutlookUser) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphOutlookUser) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphOutlookUser) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MicrosoftGraphOutlookUser) SetId(v string) {
	o.Id = &v
}

// GetMasterCategories returns the MasterCategories field value if set, zero value otherwise.
func (o *MicrosoftGraphOutlookUser) GetMasterCategories() []MicrosoftGraphOutlookCategory {
	if o == nil || o.MasterCategories == nil {
		var ret []MicrosoftGraphOutlookCategory
		return ret
	}
	return *o.MasterCategories
}

// GetMasterCategoriesOk returns a tuple with the MasterCategories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphOutlookUser) GetMasterCategoriesOk() (*[]MicrosoftGraphOutlookCategory, bool) {
	if o == nil || o.MasterCategories == nil {
		return nil, false
	}
	return o.MasterCategories, true
}

// HasMasterCategories returns a boolean if a field has been set.
func (o *MicrosoftGraphOutlookUser) HasMasterCategories() bool {
	if o != nil && o.MasterCategories != nil {
		return true
	}

	return false
}

// SetMasterCategories gets a reference to the given []MicrosoftGraphOutlookCategory and assigns it to the MasterCategories field.
func (o *MicrosoftGraphOutlookUser) SetMasterCategories(v []MicrosoftGraphOutlookCategory) {
	o.MasterCategories = &v
}

func (o MicrosoftGraphOutlookUser) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.MasterCategories != nil {
		toSerialize["masterCategories"] = o.MasterCategories
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphOutlookUser struct {
	value *MicrosoftGraphOutlookUser
	isSet bool
}

func (v NullableMicrosoftGraphOutlookUser) Get() *MicrosoftGraphOutlookUser {
	return v.value
}

func (v *NullableMicrosoftGraphOutlookUser) Set(val *MicrosoftGraphOutlookUser) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphOutlookUser) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphOutlookUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphOutlookUser(val *MicrosoftGraphOutlookUser) *NullableMicrosoftGraphOutlookUser {
	return &NullableMicrosoftGraphOutlookUser{value: val, isSet: true}
}

func (v NullableMicrosoftGraphOutlookUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphOutlookUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


