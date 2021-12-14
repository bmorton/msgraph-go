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

// AgreementFileLocalization struct for AgreementFileLocalization
type AgreementFileLocalization struct {
	Versions *[]MicrosoftGraphAgreementFileVersion `json:"versions,omitempty"`
}

// NewAgreementFileLocalization instantiates a new AgreementFileLocalization object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAgreementFileLocalization() *AgreementFileLocalization {
	this := AgreementFileLocalization{}
	return &this
}

// NewAgreementFileLocalizationWithDefaults instantiates a new AgreementFileLocalization object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgreementFileLocalizationWithDefaults() *AgreementFileLocalization {
	this := AgreementFileLocalization{}
	return &this
}

// GetVersions returns the Versions field value if set, zero value otherwise.
func (o *AgreementFileLocalization) GetVersions() []MicrosoftGraphAgreementFileVersion {
	if o == nil || o.Versions == nil {
		var ret []MicrosoftGraphAgreementFileVersion
		return ret
	}
	return *o.Versions
}

// GetVersionsOk returns a tuple with the Versions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgreementFileLocalization) GetVersionsOk() (*[]MicrosoftGraphAgreementFileVersion, bool) {
	if o == nil || o.Versions == nil {
		return nil, false
	}
	return o.Versions, true
}

// HasVersions returns a boolean if a field has been set.
func (o *AgreementFileLocalization) HasVersions() bool {
	if o != nil && o.Versions != nil {
		return true
	}

	return false
}

// SetVersions gets a reference to the given []MicrosoftGraphAgreementFileVersion and assigns it to the Versions field.
func (o *AgreementFileLocalization) SetVersions(v []MicrosoftGraphAgreementFileVersion) {
	o.Versions = &v
}

func (o AgreementFileLocalization) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Versions != nil {
		toSerialize["versions"] = o.Versions
	}
	return json.Marshal(toSerialize)
}

type NullableAgreementFileLocalization struct {
	value *AgreementFileLocalization
	isSet bool
}

func (v NullableAgreementFileLocalization) Get() *AgreementFileLocalization {
	return v.value
}

func (v *NullableAgreementFileLocalization) Set(val *AgreementFileLocalization) {
	v.value = val
	v.isSet = true
}

func (v NullableAgreementFileLocalization) IsSet() bool {
	return v.isSet
}

func (v *NullableAgreementFileLocalization) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgreementFileLocalization(val *AgreementFileLocalization) *NullableAgreementFileLocalization {
	return &NullableAgreementFileLocalization{value: val, isSet: true}
}

func (v NullableAgreementFileLocalization) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgreementFileLocalization) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

