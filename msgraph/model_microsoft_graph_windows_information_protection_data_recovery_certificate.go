/*
Partial Graph API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package msgraph

import (
	"encoding/json"
	"time"
)

// MicrosoftGraphWindowsInformationProtectionDataRecoveryCertificate Windows Information Protection DataRecoveryCertificate
type MicrosoftGraphWindowsInformationProtectionDataRecoveryCertificate struct {
	// Data recovery Certificate
	Certificate NullableString `json:"certificate,omitempty"`
	// Data recovery Certificate description
	Description NullableString `json:"description,omitempty"`
	// Data recovery Certificate expiration datetime
	ExpirationDateTime *time.Time `json:"expirationDateTime,omitempty"`
	// Data recovery Certificate subject name
	SubjectName NullableString `json:"subjectName,omitempty"`
}

// NewMicrosoftGraphWindowsInformationProtectionDataRecoveryCertificate instantiates a new MicrosoftGraphWindowsInformationProtectionDataRecoveryCertificate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphWindowsInformationProtectionDataRecoveryCertificate() *MicrosoftGraphWindowsInformationProtectionDataRecoveryCertificate {
	this := MicrosoftGraphWindowsInformationProtectionDataRecoveryCertificate{}
	return &this
}

// NewMicrosoftGraphWindowsInformationProtectionDataRecoveryCertificateWithDefaults instantiates a new MicrosoftGraphWindowsInformationProtectionDataRecoveryCertificate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphWindowsInformationProtectionDataRecoveryCertificateWithDefaults() *MicrosoftGraphWindowsInformationProtectionDataRecoveryCertificate {
	this := MicrosoftGraphWindowsInformationProtectionDataRecoveryCertificate{}
	return &this
}

// GetCertificate returns the Certificate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphWindowsInformationProtectionDataRecoveryCertificate) GetCertificate() string {
	if o == nil || o.Certificate.Get() == nil {
		var ret string
		return ret
	}
	return *o.Certificate.Get()
}

// GetCertificateOk returns a tuple with the Certificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphWindowsInformationProtectionDataRecoveryCertificate) GetCertificateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Certificate.Get(), o.Certificate.IsSet()
}

// HasCertificate returns a boolean if a field has been set.
func (o *MicrosoftGraphWindowsInformationProtectionDataRecoveryCertificate) HasCertificate() bool {
	if o != nil && o.Certificate.IsSet() {
		return true
	}

	return false
}

// SetCertificate gets a reference to the given NullableString and assigns it to the Certificate field.
func (o *MicrosoftGraphWindowsInformationProtectionDataRecoveryCertificate) SetCertificate(v string) {
	o.Certificate.Set(&v)
}
// SetCertificateNil sets the value for Certificate to be an explicit nil
func (o *MicrosoftGraphWindowsInformationProtectionDataRecoveryCertificate) SetCertificateNil() {
	o.Certificate.Set(nil)
}

// UnsetCertificate ensures that no value is present for Certificate, not even an explicit nil
func (o *MicrosoftGraphWindowsInformationProtectionDataRecoveryCertificate) UnsetCertificate() {
	o.Certificate.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphWindowsInformationProtectionDataRecoveryCertificate) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphWindowsInformationProtectionDataRecoveryCertificate) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *MicrosoftGraphWindowsInformationProtectionDataRecoveryCertificate) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *MicrosoftGraphWindowsInformationProtectionDataRecoveryCertificate) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *MicrosoftGraphWindowsInformationProtectionDataRecoveryCertificate) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *MicrosoftGraphWindowsInformationProtectionDataRecoveryCertificate) UnsetDescription() {
	o.Description.Unset()
}

// GetExpirationDateTime returns the ExpirationDateTime field value if set, zero value otherwise.
func (o *MicrosoftGraphWindowsInformationProtectionDataRecoveryCertificate) GetExpirationDateTime() time.Time {
	if o == nil || o.ExpirationDateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.ExpirationDateTime
}

// GetExpirationDateTimeOk returns a tuple with the ExpirationDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphWindowsInformationProtectionDataRecoveryCertificate) GetExpirationDateTimeOk() (*time.Time, bool) {
	if o == nil || o.ExpirationDateTime == nil {
		return nil, false
	}
	return o.ExpirationDateTime, true
}

// HasExpirationDateTime returns a boolean if a field has been set.
func (o *MicrosoftGraphWindowsInformationProtectionDataRecoveryCertificate) HasExpirationDateTime() bool {
	if o != nil && o.ExpirationDateTime != nil {
		return true
	}

	return false
}

// SetExpirationDateTime gets a reference to the given time.Time and assigns it to the ExpirationDateTime field.
func (o *MicrosoftGraphWindowsInformationProtectionDataRecoveryCertificate) SetExpirationDateTime(v time.Time) {
	o.ExpirationDateTime = &v
}

// GetSubjectName returns the SubjectName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphWindowsInformationProtectionDataRecoveryCertificate) GetSubjectName() string {
	if o == nil || o.SubjectName.Get() == nil {
		var ret string
		return ret
	}
	return *o.SubjectName.Get()
}

// GetSubjectNameOk returns a tuple with the SubjectName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphWindowsInformationProtectionDataRecoveryCertificate) GetSubjectNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SubjectName.Get(), o.SubjectName.IsSet()
}

// HasSubjectName returns a boolean if a field has been set.
func (o *MicrosoftGraphWindowsInformationProtectionDataRecoveryCertificate) HasSubjectName() bool {
	if o != nil && o.SubjectName.IsSet() {
		return true
	}

	return false
}

// SetSubjectName gets a reference to the given NullableString and assigns it to the SubjectName field.
func (o *MicrosoftGraphWindowsInformationProtectionDataRecoveryCertificate) SetSubjectName(v string) {
	o.SubjectName.Set(&v)
}
// SetSubjectNameNil sets the value for SubjectName to be an explicit nil
func (o *MicrosoftGraphWindowsInformationProtectionDataRecoveryCertificate) SetSubjectNameNil() {
	o.SubjectName.Set(nil)
}

// UnsetSubjectName ensures that no value is present for SubjectName, not even an explicit nil
func (o *MicrosoftGraphWindowsInformationProtectionDataRecoveryCertificate) UnsetSubjectName() {
	o.SubjectName.Unset()
}

func (o MicrosoftGraphWindowsInformationProtectionDataRecoveryCertificate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Certificate.IsSet() {
		toSerialize["certificate"] = o.Certificate.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.ExpirationDateTime != nil {
		toSerialize["expirationDateTime"] = o.ExpirationDateTime
	}
	if o.SubjectName.IsSet() {
		toSerialize["subjectName"] = o.SubjectName.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphWindowsInformationProtectionDataRecoveryCertificate struct {
	value *MicrosoftGraphWindowsInformationProtectionDataRecoveryCertificate
	isSet bool
}

func (v NullableMicrosoftGraphWindowsInformationProtectionDataRecoveryCertificate) Get() *MicrosoftGraphWindowsInformationProtectionDataRecoveryCertificate {
	return v.value
}

func (v *NullableMicrosoftGraphWindowsInformationProtectionDataRecoveryCertificate) Set(val *MicrosoftGraphWindowsInformationProtectionDataRecoveryCertificate) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphWindowsInformationProtectionDataRecoveryCertificate) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphWindowsInformationProtectionDataRecoveryCertificate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphWindowsInformationProtectionDataRecoveryCertificate(val *MicrosoftGraphWindowsInformationProtectionDataRecoveryCertificate) *NullableMicrosoftGraphWindowsInformationProtectionDataRecoveryCertificate {
	return &NullableMicrosoftGraphWindowsInformationProtectionDataRecoveryCertificate{value: val, isSet: true}
}

func (v NullableMicrosoftGraphWindowsInformationProtectionDataRecoveryCertificate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphWindowsInformationProtectionDataRecoveryCertificate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


