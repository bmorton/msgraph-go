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

// AgreementFile struct for AgreementFile
type AgreementFile struct {
	Localizations *[]MicrosoftGraphAgreementFileLocalization `json:"localizations,omitempty"`
}

// NewAgreementFile instantiates a new AgreementFile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAgreementFile() *AgreementFile {
	this := AgreementFile{}
	return &this
}

// NewAgreementFileWithDefaults instantiates a new AgreementFile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgreementFileWithDefaults() *AgreementFile {
	this := AgreementFile{}
	return &this
}

// GetLocalizations returns the Localizations field value if set, zero value otherwise.
func (o *AgreementFile) GetLocalizations() []MicrosoftGraphAgreementFileLocalization {
	if o == nil || o.Localizations == nil {
		var ret []MicrosoftGraphAgreementFileLocalization
		return ret
	}
	return *o.Localizations
}

// GetLocalizationsOk returns a tuple with the Localizations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgreementFile) GetLocalizationsOk() (*[]MicrosoftGraphAgreementFileLocalization, bool) {
	if o == nil || o.Localizations == nil {
		return nil, false
	}
	return o.Localizations, true
}

// HasLocalizations returns a boolean if a field has been set.
func (o *AgreementFile) HasLocalizations() bool {
	if o != nil && o.Localizations != nil {
		return true
	}

	return false
}

// SetLocalizations gets a reference to the given []MicrosoftGraphAgreementFileLocalization and assigns it to the Localizations field.
func (o *AgreementFile) SetLocalizations(v []MicrosoftGraphAgreementFileLocalization) {
	o.Localizations = &v
}

func (o AgreementFile) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Localizations != nil {
		toSerialize["localizations"] = o.Localizations
	}
	return json.Marshal(toSerialize)
}

type NullableAgreementFile struct {
	value *AgreementFile
	isSet bool
}

func (v NullableAgreementFile) Get() *AgreementFile {
	return v.value
}

func (v *NullableAgreementFile) Set(val *AgreementFile) {
	v.value = val
	v.isSet = true
}

func (v NullableAgreementFile) IsSet() bool {
	return v.isSet
}

func (v *NullableAgreementFile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgreementFile(val *AgreementFile) *NullableAgreementFile {
	return &NullableAgreementFile{value: val, isSet: true}
}

func (v NullableAgreementFile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgreementFile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


