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

// Agreement struct for Agreement
type Agreement struct {
	// Display name of the agreement. The display name is used for internal tracking of the agreement but is not shown to end users who view the agreement.
	DisplayName NullableString `json:"displayName,omitempty"`
	// Indicates whether end users are required to accept this agreement on every device that they access it from. The end user is required to register their device in Azure AD, if they haven't already done so.
	IsPerDeviceAcceptanceRequired NullableBool `json:"isPerDeviceAcceptanceRequired,omitempty"`
	// Indicates whether the user has to expand the agreement before accepting.
	IsViewingBeforeAcceptanceRequired NullableBool `json:"isViewingBeforeAcceptanceRequired,omitempty"`
	// Expiration schedule and frequency of agreement for all users.
	TermsExpiration AnyOfmicrosoftGraphTermsExpiration `json:"termsExpiration,omitempty"`
	// The duration after which the user must re-accept the terms of use. The value is represented in ISO 8601 format for durations.
	UserReacceptRequiredFrequency NullableString `json:"userReacceptRequiredFrequency,omitempty"`
	// Read-only. Information about acceptances of this agreement.
	Acceptances *[]MicrosoftGraphAgreementAcceptance `json:"acceptances,omitempty"`
	// Default PDF linked to this agreement.
	File AnyOfmicrosoftGraphAgreementFile `json:"file,omitempty"`
	// PDFs linked to this agreement. This property is in the process of being deprecated. Use the  file property instead.
	Files *[]MicrosoftGraphAgreementFileLocalization `json:"files,omitempty"`
}

// NewAgreement instantiates a new Agreement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAgreement() *Agreement {
	this := Agreement{}
	return &this
}

// NewAgreementWithDefaults instantiates a new Agreement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgreementWithDefaults() *Agreement {
	this := Agreement{}
	return &this
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Agreement) GetDisplayName() string {
	if o == nil || o.DisplayName.Get() == nil {
		var ret string
		return ret
	}
	return *o.DisplayName.Get()
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Agreement) GetDisplayNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DisplayName.Get(), o.DisplayName.IsSet()
}

// HasDisplayName returns a boolean if a field has been set.
func (o *Agreement) HasDisplayName() bool {
	if o != nil && o.DisplayName.IsSet() {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given NullableString and assigns it to the DisplayName field.
func (o *Agreement) SetDisplayName(v string) {
	o.DisplayName.Set(&v)
}
// SetDisplayNameNil sets the value for DisplayName to be an explicit nil
func (o *Agreement) SetDisplayNameNil() {
	o.DisplayName.Set(nil)
}

// UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
func (o *Agreement) UnsetDisplayName() {
	o.DisplayName.Unset()
}

// GetIsPerDeviceAcceptanceRequired returns the IsPerDeviceAcceptanceRequired field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Agreement) GetIsPerDeviceAcceptanceRequired() bool {
	if o == nil || o.IsPerDeviceAcceptanceRequired.Get() == nil {
		var ret bool
		return ret
	}
	return *o.IsPerDeviceAcceptanceRequired.Get()
}

// GetIsPerDeviceAcceptanceRequiredOk returns a tuple with the IsPerDeviceAcceptanceRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Agreement) GetIsPerDeviceAcceptanceRequiredOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IsPerDeviceAcceptanceRequired.Get(), o.IsPerDeviceAcceptanceRequired.IsSet()
}

// HasIsPerDeviceAcceptanceRequired returns a boolean if a field has been set.
func (o *Agreement) HasIsPerDeviceAcceptanceRequired() bool {
	if o != nil && o.IsPerDeviceAcceptanceRequired.IsSet() {
		return true
	}

	return false
}

// SetIsPerDeviceAcceptanceRequired gets a reference to the given NullableBool and assigns it to the IsPerDeviceAcceptanceRequired field.
func (o *Agreement) SetIsPerDeviceAcceptanceRequired(v bool) {
	o.IsPerDeviceAcceptanceRequired.Set(&v)
}
// SetIsPerDeviceAcceptanceRequiredNil sets the value for IsPerDeviceAcceptanceRequired to be an explicit nil
func (o *Agreement) SetIsPerDeviceAcceptanceRequiredNil() {
	o.IsPerDeviceAcceptanceRequired.Set(nil)
}

// UnsetIsPerDeviceAcceptanceRequired ensures that no value is present for IsPerDeviceAcceptanceRequired, not even an explicit nil
func (o *Agreement) UnsetIsPerDeviceAcceptanceRequired() {
	o.IsPerDeviceAcceptanceRequired.Unset()
}

// GetIsViewingBeforeAcceptanceRequired returns the IsViewingBeforeAcceptanceRequired field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Agreement) GetIsViewingBeforeAcceptanceRequired() bool {
	if o == nil || o.IsViewingBeforeAcceptanceRequired.Get() == nil {
		var ret bool
		return ret
	}
	return *o.IsViewingBeforeAcceptanceRequired.Get()
}

// GetIsViewingBeforeAcceptanceRequiredOk returns a tuple with the IsViewingBeforeAcceptanceRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Agreement) GetIsViewingBeforeAcceptanceRequiredOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IsViewingBeforeAcceptanceRequired.Get(), o.IsViewingBeforeAcceptanceRequired.IsSet()
}

// HasIsViewingBeforeAcceptanceRequired returns a boolean if a field has been set.
func (o *Agreement) HasIsViewingBeforeAcceptanceRequired() bool {
	if o != nil && o.IsViewingBeforeAcceptanceRequired.IsSet() {
		return true
	}

	return false
}

// SetIsViewingBeforeAcceptanceRequired gets a reference to the given NullableBool and assigns it to the IsViewingBeforeAcceptanceRequired field.
func (o *Agreement) SetIsViewingBeforeAcceptanceRequired(v bool) {
	o.IsViewingBeforeAcceptanceRequired.Set(&v)
}
// SetIsViewingBeforeAcceptanceRequiredNil sets the value for IsViewingBeforeAcceptanceRequired to be an explicit nil
func (o *Agreement) SetIsViewingBeforeAcceptanceRequiredNil() {
	o.IsViewingBeforeAcceptanceRequired.Set(nil)
}

// UnsetIsViewingBeforeAcceptanceRequired ensures that no value is present for IsViewingBeforeAcceptanceRequired, not even an explicit nil
func (o *Agreement) UnsetIsViewingBeforeAcceptanceRequired() {
	o.IsViewingBeforeAcceptanceRequired.Unset()
}

// GetTermsExpiration returns the TermsExpiration field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Agreement) GetTermsExpiration() AnyOfmicrosoftGraphTermsExpiration {
	if o == nil  {
		var ret AnyOfmicrosoftGraphTermsExpiration
		return ret
	}
	return o.TermsExpiration
}

// GetTermsExpirationOk returns a tuple with the TermsExpiration field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Agreement) GetTermsExpirationOk() (*AnyOfmicrosoftGraphTermsExpiration, bool) {
	if o == nil || o.TermsExpiration == nil {
		return nil, false
	}
	return &o.TermsExpiration, true
}

// HasTermsExpiration returns a boolean if a field has been set.
func (o *Agreement) HasTermsExpiration() bool {
	if o != nil && o.TermsExpiration != nil {
		return true
	}

	return false
}

// SetTermsExpiration gets a reference to the given AnyOfmicrosoftGraphTermsExpiration and assigns it to the TermsExpiration field.
func (o *Agreement) SetTermsExpiration(v AnyOfmicrosoftGraphTermsExpiration) {
	o.TermsExpiration = v
}

// GetUserReacceptRequiredFrequency returns the UserReacceptRequiredFrequency field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Agreement) GetUserReacceptRequiredFrequency() string {
	if o == nil || o.UserReacceptRequiredFrequency.Get() == nil {
		var ret string
		return ret
	}
	return *o.UserReacceptRequiredFrequency.Get()
}

// GetUserReacceptRequiredFrequencyOk returns a tuple with the UserReacceptRequiredFrequency field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Agreement) GetUserReacceptRequiredFrequencyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.UserReacceptRequiredFrequency.Get(), o.UserReacceptRequiredFrequency.IsSet()
}

// HasUserReacceptRequiredFrequency returns a boolean if a field has been set.
func (o *Agreement) HasUserReacceptRequiredFrequency() bool {
	if o != nil && o.UserReacceptRequiredFrequency.IsSet() {
		return true
	}

	return false
}

// SetUserReacceptRequiredFrequency gets a reference to the given NullableString and assigns it to the UserReacceptRequiredFrequency field.
func (o *Agreement) SetUserReacceptRequiredFrequency(v string) {
	o.UserReacceptRequiredFrequency.Set(&v)
}
// SetUserReacceptRequiredFrequencyNil sets the value for UserReacceptRequiredFrequency to be an explicit nil
func (o *Agreement) SetUserReacceptRequiredFrequencyNil() {
	o.UserReacceptRequiredFrequency.Set(nil)
}

// UnsetUserReacceptRequiredFrequency ensures that no value is present for UserReacceptRequiredFrequency, not even an explicit nil
func (o *Agreement) UnsetUserReacceptRequiredFrequency() {
	o.UserReacceptRequiredFrequency.Unset()
}

// GetAcceptances returns the Acceptances field value if set, zero value otherwise.
func (o *Agreement) GetAcceptances() []MicrosoftGraphAgreementAcceptance {
	if o == nil || o.Acceptances == nil {
		var ret []MicrosoftGraphAgreementAcceptance
		return ret
	}
	return *o.Acceptances
}

// GetAcceptancesOk returns a tuple with the Acceptances field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Agreement) GetAcceptancesOk() (*[]MicrosoftGraphAgreementAcceptance, bool) {
	if o == nil || o.Acceptances == nil {
		return nil, false
	}
	return o.Acceptances, true
}

// HasAcceptances returns a boolean if a field has been set.
func (o *Agreement) HasAcceptances() bool {
	if o != nil && o.Acceptances != nil {
		return true
	}

	return false
}

// SetAcceptances gets a reference to the given []MicrosoftGraphAgreementAcceptance and assigns it to the Acceptances field.
func (o *Agreement) SetAcceptances(v []MicrosoftGraphAgreementAcceptance) {
	o.Acceptances = &v
}

// GetFile returns the File field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Agreement) GetFile() AnyOfmicrosoftGraphAgreementFile {
	if o == nil  {
		var ret AnyOfmicrosoftGraphAgreementFile
		return ret
	}
	return o.File
}

// GetFileOk returns a tuple with the File field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Agreement) GetFileOk() (*AnyOfmicrosoftGraphAgreementFile, bool) {
	if o == nil || o.File == nil {
		return nil, false
	}
	return &o.File, true
}

// HasFile returns a boolean if a field has been set.
func (o *Agreement) HasFile() bool {
	if o != nil && o.File != nil {
		return true
	}

	return false
}

// SetFile gets a reference to the given AnyOfmicrosoftGraphAgreementFile and assigns it to the File field.
func (o *Agreement) SetFile(v AnyOfmicrosoftGraphAgreementFile) {
	o.File = v
}

// GetFiles returns the Files field value if set, zero value otherwise.
func (o *Agreement) GetFiles() []MicrosoftGraphAgreementFileLocalization {
	if o == nil || o.Files == nil {
		var ret []MicrosoftGraphAgreementFileLocalization
		return ret
	}
	return *o.Files
}

// GetFilesOk returns a tuple with the Files field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Agreement) GetFilesOk() (*[]MicrosoftGraphAgreementFileLocalization, bool) {
	if o == nil || o.Files == nil {
		return nil, false
	}
	return o.Files, true
}

// HasFiles returns a boolean if a field has been set.
func (o *Agreement) HasFiles() bool {
	if o != nil && o.Files != nil {
		return true
	}

	return false
}

// SetFiles gets a reference to the given []MicrosoftGraphAgreementFileLocalization and assigns it to the Files field.
func (o *Agreement) SetFiles(v []MicrosoftGraphAgreementFileLocalization) {
	o.Files = &v
}

func (o Agreement) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DisplayName.IsSet() {
		toSerialize["displayName"] = o.DisplayName.Get()
	}
	if o.IsPerDeviceAcceptanceRequired.IsSet() {
		toSerialize["isPerDeviceAcceptanceRequired"] = o.IsPerDeviceAcceptanceRequired.Get()
	}
	if o.IsViewingBeforeAcceptanceRequired.IsSet() {
		toSerialize["isViewingBeforeAcceptanceRequired"] = o.IsViewingBeforeAcceptanceRequired.Get()
	}
	if o.TermsExpiration != nil {
		toSerialize["termsExpiration"] = o.TermsExpiration
	}
	if o.UserReacceptRequiredFrequency.IsSet() {
		toSerialize["userReacceptRequiredFrequency"] = o.UserReacceptRequiredFrequency.Get()
	}
	if o.Acceptances != nil {
		toSerialize["acceptances"] = o.Acceptances
	}
	if o.File != nil {
		toSerialize["file"] = o.File
	}
	if o.Files != nil {
		toSerialize["files"] = o.Files
	}
	return json.Marshal(toSerialize)
}

type NullableAgreement struct {
	value *Agreement
	isSet bool
}

func (v NullableAgreement) Get() *Agreement {
	return v.value
}

func (v *NullableAgreement) Set(val *Agreement) {
	v.value = val
	v.isSet = true
}

func (v NullableAgreement) IsSet() bool {
	return v.isSet
}

func (v *NullableAgreement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgreement(val *Agreement) *NullableAgreement {
	return &NullableAgreement{value: val, isSet: true}
}

func (v NullableAgreement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgreement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


