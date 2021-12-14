/*
Partial Graph API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// Fido2AuthenticationMethod struct for Fido2AuthenticationMethod
type Fido2AuthenticationMethod struct {
	// Authenticator Attestation GUID, an identifier that indicates the type (e.g. make and model) of the authenticator.
	AaGuid NullableString `json:"aaGuid,omitempty"`
	// The attestation certificate(s) attached to this security key.
	AttestationCertificates *[]*string `json:"attestationCertificates,omitempty"`
	// The attestation level of this FIDO2 security key. Possible values are: attested, or notAttested.
	AttestationLevel AnyOfmicrosoftGraphAttestationLevel `json:"attestationLevel,omitempty"`
	// The timestamp when this key was registered to the user.
	CreatedDateTime NullableTime `json:"createdDateTime,omitempty"`
	// The display name of the key as given by the user.
	DisplayName NullableString `json:"displayName,omitempty"`
	// The manufacturer-assigned model of the FIDO2 security key.
	Model NullableString `json:"model,omitempty"`
}

// NewFido2AuthenticationMethod instantiates a new Fido2AuthenticationMethod object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFido2AuthenticationMethod() *Fido2AuthenticationMethod {
	this := Fido2AuthenticationMethod{}
	return &this
}

// NewFido2AuthenticationMethodWithDefaults instantiates a new Fido2AuthenticationMethod object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFido2AuthenticationMethodWithDefaults() *Fido2AuthenticationMethod {
	this := Fido2AuthenticationMethod{}
	return &this
}

// GetAaGuid returns the AaGuid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Fido2AuthenticationMethod) GetAaGuid() string {
	if o == nil || o.AaGuid.Get() == nil {
		var ret string
		return ret
	}
	return *o.AaGuid.Get()
}

// GetAaGuidOk returns a tuple with the AaGuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Fido2AuthenticationMethod) GetAaGuidOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AaGuid.Get(), o.AaGuid.IsSet()
}

// HasAaGuid returns a boolean if a field has been set.
func (o *Fido2AuthenticationMethod) HasAaGuid() bool {
	if o != nil && o.AaGuid.IsSet() {
		return true
	}

	return false
}

// SetAaGuid gets a reference to the given NullableString and assigns it to the AaGuid field.
func (o *Fido2AuthenticationMethod) SetAaGuid(v string) {
	o.AaGuid.Set(&v)
}
// SetAaGuidNil sets the value for AaGuid to be an explicit nil
func (o *Fido2AuthenticationMethod) SetAaGuidNil() {
	o.AaGuid.Set(nil)
}

// UnsetAaGuid ensures that no value is present for AaGuid, not even an explicit nil
func (o *Fido2AuthenticationMethod) UnsetAaGuid() {
	o.AaGuid.Unset()
}

// GetAttestationCertificates returns the AttestationCertificates field value if set, zero value otherwise.
func (o *Fido2AuthenticationMethod) GetAttestationCertificates() []*string {
	if o == nil || o.AttestationCertificates == nil {
		var ret []*string
		return ret
	}
	return *o.AttestationCertificates
}

// GetAttestationCertificatesOk returns a tuple with the AttestationCertificates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Fido2AuthenticationMethod) GetAttestationCertificatesOk() (*[]*string, bool) {
	if o == nil || o.AttestationCertificates == nil {
		return nil, false
	}
	return o.AttestationCertificates, true
}

// HasAttestationCertificates returns a boolean if a field has been set.
func (o *Fido2AuthenticationMethod) HasAttestationCertificates() bool {
	if o != nil && o.AttestationCertificates != nil {
		return true
	}

	return false
}

// SetAttestationCertificates gets a reference to the given []*string and assigns it to the AttestationCertificates field.
func (o *Fido2AuthenticationMethod) SetAttestationCertificates(v []*string) {
	o.AttestationCertificates = &v
}

// GetAttestationLevel returns the AttestationLevel field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Fido2AuthenticationMethod) GetAttestationLevel() AnyOfmicrosoftGraphAttestationLevel {
	if o == nil  {
		var ret AnyOfmicrosoftGraphAttestationLevel
		return ret
	}
	return o.AttestationLevel
}

// GetAttestationLevelOk returns a tuple with the AttestationLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Fido2AuthenticationMethod) GetAttestationLevelOk() (*AnyOfmicrosoftGraphAttestationLevel, bool) {
	if o == nil || o.AttestationLevel == nil {
		return nil, false
	}
	return &o.AttestationLevel, true
}

// HasAttestationLevel returns a boolean if a field has been set.
func (o *Fido2AuthenticationMethod) HasAttestationLevel() bool {
	if o != nil && o.AttestationLevel != nil {
		return true
	}

	return false
}

// SetAttestationLevel gets a reference to the given AnyOfmicrosoftGraphAttestationLevel and assigns it to the AttestationLevel field.
func (o *Fido2AuthenticationMethod) SetAttestationLevel(v AnyOfmicrosoftGraphAttestationLevel) {
	o.AttestationLevel = v
}

// GetCreatedDateTime returns the CreatedDateTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Fido2AuthenticationMethod) GetCreatedDateTime() time.Time {
	if o == nil || o.CreatedDateTime.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedDateTime.Get()
}

// GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Fido2AuthenticationMethod) GetCreatedDateTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CreatedDateTime.Get(), o.CreatedDateTime.IsSet()
}

// HasCreatedDateTime returns a boolean if a field has been set.
func (o *Fido2AuthenticationMethod) HasCreatedDateTime() bool {
	if o != nil && o.CreatedDateTime.IsSet() {
		return true
	}

	return false
}

// SetCreatedDateTime gets a reference to the given NullableTime and assigns it to the CreatedDateTime field.
func (o *Fido2AuthenticationMethod) SetCreatedDateTime(v time.Time) {
	o.CreatedDateTime.Set(&v)
}
// SetCreatedDateTimeNil sets the value for CreatedDateTime to be an explicit nil
func (o *Fido2AuthenticationMethod) SetCreatedDateTimeNil() {
	o.CreatedDateTime.Set(nil)
}

// UnsetCreatedDateTime ensures that no value is present for CreatedDateTime, not even an explicit nil
func (o *Fido2AuthenticationMethod) UnsetCreatedDateTime() {
	o.CreatedDateTime.Unset()
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Fido2AuthenticationMethod) GetDisplayName() string {
	if o == nil || o.DisplayName.Get() == nil {
		var ret string
		return ret
	}
	return *o.DisplayName.Get()
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Fido2AuthenticationMethod) GetDisplayNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DisplayName.Get(), o.DisplayName.IsSet()
}

// HasDisplayName returns a boolean if a field has been set.
func (o *Fido2AuthenticationMethod) HasDisplayName() bool {
	if o != nil && o.DisplayName.IsSet() {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given NullableString and assigns it to the DisplayName field.
func (o *Fido2AuthenticationMethod) SetDisplayName(v string) {
	o.DisplayName.Set(&v)
}
// SetDisplayNameNil sets the value for DisplayName to be an explicit nil
func (o *Fido2AuthenticationMethod) SetDisplayNameNil() {
	o.DisplayName.Set(nil)
}

// UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
func (o *Fido2AuthenticationMethod) UnsetDisplayName() {
	o.DisplayName.Unset()
}

// GetModel returns the Model field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Fido2AuthenticationMethod) GetModel() string {
	if o == nil || o.Model.Get() == nil {
		var ret string
		return ret
	}
	return *o.Model.Get()
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Fido2AuthenticationMethod) GetModelOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Model.Get(), o.Model.IsSet()
}

// HasModel returns a boolean if a field has been set.
func (o *Fido2AuthenticationMethod) HasModel() bool {
	if o != nil && o.Model.IsSet() {
		return true
	}

	return false
}

// SetModel gets a reference to the given NullableString and assigns it to the Model field.
func (o *Fido2AuthenticationMethod) SetModel(v string) {
	o.Model.Set(&v)
}
// SetModelNil sets the value for Model to be an explicit nil
func (o *Fido2AuthenticationMethod) SetModelNil() {
	o.Model.Set(nil)
}

// UnsetModel ensures that no value is present for Model, not even an explicit nil
func (o *Fido2AuthenticationMethod) UnsetModel() {
	o.Model.Unset()
}

func (o Fido2AuthenticationMethod) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AaGuid.IsSet() {
		toSerialize["aaGuid"] = o.AaGuid.Get()
	}
	if o.AttestationCertificates != nil {
		toSerialize["attestationCertificates"] = o.AttestationCertificates
	}
	if o.AttestationLevel != nil {
		toSerialize["attestationLevel"] = o.AttestationLevel
	}
	if o.CreatedDateTime.IsSet() {
		toSerialize["createdDateTime"] = o.CreatedDateTime.Get()
	}
	if o.DisplayName.IsSet() {
		toSerialize["displayName"] = o.DisplayName.Get()
	}
	if o.Model.IsSet() {
		toSerialize["model"] = o.Model.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableFido2AuthenticationMethod struct {
	value *Fido2AuthenticationMethod
	isSet bool
}

func (v NullableFido2AuthenticationMethod) Get() *Fido2AuthenticationMethod {
	return v.value
}

func (v *NullableFido2AuthenticationMethod) Set(val *Fido2AuthenticationMethod) {
	v.value = val
	v.isSet = true
}

func (v NullableFido2AuthenticationMethod) IsSet() bool {
	return v.isSet
}

func (v *NullableFido2AuthenticationMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFido2AuthenticationMethod(val *Fido2AuthenticationMethod) *NullableFido2AuthenticationMethod {
	return &NullableFido2AuthenticationMethod{value: val, isSet: true}
}

func (v NullableFido2AuthenticationMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFido2AuthenticationMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

