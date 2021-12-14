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

// SecureScore struct for SecureScore
type SecureScore struct {
	// Active user count of the given tenant.
	ActiveUserCount NullableInt32 `json:"activeUserCount,omitempty"`
	// Average score by different scopes (for example, average by industry, average by seating) and control category (Identity, Data, Device, Apps, Infrastructure) within the scope.
	AverageComparativeScores *[]*AnyOfmicrosoftGraphAverageComparativeScore `json:"averageComparativeScores,omitempty"`
	// GUID string for tenant ID.
	AzureTenantId *string `json:"azureTenantId,omitempty"`
	// Contains tenant scores for a set of controls.
	ControlScores *[]*AnyOfmicrosoftGraphControlScore `json:"controlScores,omitempty"`
	// The date when the entity is created.
	CreatedDateTime NullableTime `json:"createdDateTime,omitempty"`
	// Tenant current attained score on specified date.
	CurrentScore AnyOfnumberstringstring `json:"currentScore,omitempty"`
	// Microsoft-provided services for the tenant (for example, Exchange online, Skype, Sharepoint).
	EnabledServices *[]*string `json:"enabledServices,omitempty"`
	// Licensed user count of the given tenant.
	LicensedUserCount NullableInt32 `json:"licensedUserCount,omitempty"`
	// Tenant maximum possible score on specified date.
	MaxScore AnyOfnumberstringstring `json:"maxScore,omitempty"`
	// Complex type containing details about the security product/service vendor, provider, and subprovider (for example, vendor=Microsoft; provider=SecureScore). Required.
	VendorInformation AnyOfmicrosoftGraphSecurityVendorInformation `json:"vendorInformation,omitempty"`
}

// NewSecureScore instantiates a new SecureScore object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecureScore() *SecureScore {
	this := SecureScore{}
	return &this
}

// NewSecureScoreWithDefaults instantiates a new SecureScore object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecureScoreWithDefaults() *SecureScore {
	this := SecureScore{}
	return &this
}

// GetActiveUserCount returns the ActiveUserCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SecureScore) GetActiveUserCount() int32 {
	if o == nil || o.ActiveUserCount.Get() == nil {
		var ret int32
		return ret
	}
	return *o.ActiveUserCount.Get()
}

// GetActiveUserCountOk returns a tuple with the ActiveUserCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SecureScore) GetActiveUserCountOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ActiveUserCount.Get(), o.ActiveUserCount.IsSet()
}

// HasActiveUserCount returns a boolean if a field has been set.
func (o *SecureScore) HasActiveUserCount() bool {
	if o != nil && o.ActiveUserCount.IsSet() {
		return true
	}

	return false
}

// SetActiveUserCount gets a reference to the given NullableInt32 and assigns it to the ActiveUserCount field.
func (o *SecureScore) SetActiveUserCount(v int32) {
	o.ActiveUserCount.Set(&v)
}
// SetActiveUserCountNil sets the value for ActiveUserCount to be an explicit nil
func (o *SecureScore) SetActiveUserCountNil() {
	o.ActiveUserCount.Set(nil)
}

// UnsetActiveUserCount ensures that no value is present for ActiveUserCount, not even an explicit nil
func (o *SecureScore) UnsetActiveUserCount() {
	o.ActiveUserCount.Unset()
}

// GetAverageComparativeScores returns the AverageComparativeScores field value if set, zero value otherwise.
func (o *SecureScore) GetAverageComparativeScores() []*AnyOfmicrosoftGraphAverageComparativeScore {
	if o == nil || o.AverageComparativeScores == nil {
		var ret []*AnyOfmicrosoftGraphAverageComparativeScore
		return ret
	}
	return *o.AverageComparativeScores
}

// GetAverageComparativeScoresOk returns a tuple with the AverageComparativeScores field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecureScore) GetAverageComparativeScoresOk() (*[]*AnyOfmicrosoftGraphAverageComparativeScore, bool) {
	if o == nil || o.AverageComparativeScores == nil {
		return nil, false
	}
	return o.AverageComparativeScores, true
}

// HasAverageComparativeScores returns a boolean if a field has been set.
func (o *SecureScore) HasAverageComparativeScores() bool {
	if o != nil && o.AverageComparativeScores != nil {
		return true
	}

	return false
}

// SetAverageComparativeScores gets a reference to the given []*AnyOfmicrosoftGraphAverageComparativeScore and assigns it to the AverageComparativeScores field.
func (o *SecureScore) SetAverageComparativeScores(v []*AnyOfmicrosoftGraphAverageComparativeScore) {
	o.AverageComparativeScores = &v
}

// GetAzureTenantId returns the AzureTenantId field value if set, zero value otherwise.
func (o *SecureScore) GetAzureTenantId() string {
	if o == nil || o.AzureTenantId == nil {
		var ret string
		return ret
	}
	return *o.AzureTenantId
}

// GetAzureTenantIdOk returns a tuple with the AzureTenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecureScore) GetAzureTenantIdOk() (*string, bool) {
	if o == nil || o.AzureTenantId == nil {
		return nil, false
	}
	return o.AzureTenantId, true
}

// HasAzureTenantId returns a boolean if a field has been set.
func (o *SecureScore) HasAzureTenantId() bool {
	if o != nil && o.AzureTenantId != nil {
		return true
	}

	return false
}

// SetAzureTenantId gets a reference to the given string and assigns it to the AzureTenantId field.
func (o *SecureScore) SetAzureTenantId(v string) {
	o.AzureTenantId = &v
}

// GetControlScores returns the ControlScores field value if set, zero value otherwise.
func (o *SecureScore) GetControlScores() []*AnyOfmicrosoftGraphControlScore {
	if o == nil || o.ControlScores == nil {
		var ret []*AnyOfmicrosoftGraphControlScore
		return ret
	}
	return *o.ControlScores
}

// GetControlScoresOk returns a tuple with the ControlScores field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecureScore) GetControlScoresOk() (*[]*AnyOfmicrosoftGraphControlScore, bool) {
	if o == nil || o.ControlScores == nil {
		return nil, false
	}
	return o.ControlScores, true
}

// HasControlScores returns a boolean if a field has been set.
func (o *SecureScore) HasControlScores() bool {
	if o != nil && o.ControlScores != nil {
		return true
	}

	return false
}

// SetControlScores gets a reference to the given []*AnyOfmicrosoftGraphControlScore and assigns it to the ControlScores field.
func (o *SecureScore) SetControlScores(v []*AnyOfmicrosoftGraphControlScore) {
	o.ControlScores = &v
}

// GetCreatedDateTime returns the CreatedDateTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SecureScore) GetCreatedDateTime() time.Time {
	if o == nil || o.CreatedDateTime.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedDateTime.Get()
}

// GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SecureScore) GetCreatedDateTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CreatedDateTime.Get(), o.CreatedDateTime.IsSet()
}

// HasCreatedDateTime returns a boolean if a field has been set.
func (o *SecureScore) HasCreatedDateTime() bool {
	if o != nil && o.CreatedDateTime.IsSet() {
		return true
	}

	return false
}

// SetCreatedDateTime gets a reference to the given NullableTime and assigns it to the CreatedDateTime field.
func (o *SecureScore) SetCreatedDateTime(v time.Time) {
	o.CreatedDateTime.Set(&v)
}
// SetCreatedDateTimeNil sets the value for CreatedDateTime to be an explicit nil
func (o *SecureScore) SetCreatedDateTimeNil() {
	o.CreatedDateTime.Set(nil)
}

// UnsetCreatedDateTime ensures that no value is present for CreatedDateTime, not even an explicit nil
func (o *SecureScore) UnsetCreatedDateTime() {
	o.CreatedDateTime.Unset()
}

// GetCurrentScore returns the CurrentScore field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SecureScore) GetCurrentScore() AnyOfnumberstringstring {
	if o == nil  {
		var ret AnyOfnumberstringstring
		return ret
	}
	return o.CurrentScore
}

// GetCurrentScoreOk returns a tuple with the CurrentScore field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SecureScore) GetCurrentScoreOk() (*AnyOfnumberstringstring, bool) {
	if o == nil || o.CurrentScore == nil {
		return nil, false
	}
	return &o.CurrentScore, true
}

// HasCurrentScore returns a boolean if a field has been set.
func (o *SecureScore) HasCurrentScore() bool {
	if o != nil && o.CurrentScore != nil {
		return true
	}

	return false
}

// SetCurrentScore gets a reference to the given AnyOfnumberstringstring and assigns it to the CurrentScore field.
func (o *SecureScore) SetCurrentScore(v AnyOfnumberstringstring) {
	o.CurrentScore = v
}

// GetEnabledServices returns the EnabledServices field value if set, zero value otherwise.
func (o *SecureScore) GetEnabledServices() []*string {
	if o == nil || o.EnabledServices == nil {
		var ret []*string
		return ret
	}
	return *o.EnabledServices
}

// GetEnabledServicesOk returns a tuple with the EnabledServices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecureScore) GetEnabledServicesOk() (*[]*string, bool) {
	if o == nil || o.EnabledServices == nil {
		return nil, false
	}
	return o.EnabledServices, true
}

// HasEnabledServices returns a boolean if a field has been set.
func (o *SecureScore) HasEnabledServices() bool {
	if o != nil && o.EnabledServices != nil {
		return true
	}

	return false
}

// SetEnabledServices gets a reference to the given []*string and assigns it to the EnabledServices field.
func (o *SecureScore) SetEnabledServices(v []*string) {
	o.EnabledServices = &v
}

// GetLicensedUserCount returns the LicensedUserCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SecureScore) GetLicensedUserCount() int32 {
	if o == nil || o.LicensedUserCount.Get() == nil {
		var ret int32
		return ret
	}
	return *o.LicensedUserCount.Get()
}

// GetLicensedUserCountOk returns a tuple with the LicensedUserCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SecureScore) GetLicensedUserCountOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LicensedUserCount.Get(), o.LicensedUserCount.IsSet()
}

// HasLicensedUserCount returns a boolean if a field has been set.
func (o *SecureScore) HasLicensedUserCount() bool {
	if o != nil && o.LicensedUserCount.IsSet() {
		return true
	}

	return false
}

// SetLicensedUserCount gets a reference to the given NullableInt32 and assigns it to the LicensedUserCount field.
func (o *SecureScore) SetLicensedUserCount(v int32) {
	o.LicensedUserCount.Set(&v)
}
// SetLicensedUserCountNil sets the value for LicensedUserCount to be an explicit nil
func (o *SecureScore) SetLicensedUserCountNil() {
	o.LicensedUserCount.Set(nil)
}

// UnsetLicensedUserCount ensures that no value is present for LicensedUserCount, not even an explicit nil
func (o *SecureScore) UnsetLicensedUserCount() {
	o.LicensedUserCount.Unset()
}

// GetMaxScore returns the MaxScore field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SecureScore) GetMaxScore() AnyOfnumberstringstring {
	if o == nil  {
		var ret AnyOfnumberstringstring
		return ret
	}
	return o.MaxScore
}

// GetMaxScoreOk returns a tuple with the MaxScore field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SecureScore) GetMaxScoreOk() (*AnyOfnumberstringstring, bool) {
	if o == nil || o.MaxScore == nil {
		return nil, false
	}
	return &o.MaxScore, true
}

// HasMaxScore returns a boolean if a field has been set.
func (o *SecureScore) HasMaxScore() bool {
	if o != nil && o.MaxScore != nil {
		return true
	}

	return false
}

// SetMaxScore gets a reference to the given AnyOfnumberstringstring and assigns it to the MaxScore field.
func (o *SecureScore) SetMaxScore(v AnyOfnumberstringstring) {
	o.MaxScore = v
}

// GetVendorInformation returns the VendorInformation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SecureScore) GetVendorInformation() AnyOfmicrosoftGraphSecurityVendorInformation {
	if o == nil  {
		var ret AnyOfmicrosoftGraphSecurityVendorInformation
		return ret
	}
	return o.VendorInformation
}

// GetVendorInformationOk returns a tuple with the VendorInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SecureScore) GetVendorInformationOk() (*AnyOfmicrosoftGraphSecurityVendorInformation, bool) {
	if o == nil || o.VendorInformation == nil {
		return nil, false
	}
	return &o.VendorInformation, true
}

// HasVendorInformation returns a boolean if a field has been set.
func (o *SecureScore) HasVendorInformation() bool {
	if o != nil && o.VendorInformation != nil {
		return true
	}

	return false
}

// SetVendorInformation gets a reference to the given AnyOfmicrosoftGraphSecurityVendorInformation and assigns it to the VendorInformation field.
func (o *SecureScore) SetVendorInformation(v AnyOfmicrosoftGraphSecurityVendorInformation) {
	o.VendorInformation = v
}

func (o SecureScore) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ActiveUserCount.IsSet() {
		toSerialize["activeUserCount"] = o.ActiveUserCount.Get()
	}
	if o.AverageComparativeScores != nil {
		toSerialize["averageComparativeScores"] = o.AverageComparativeScores
	}
	if o.AzureTenantId != nil {
		toSerialize["azureTenantId"] = o.AzureTenantId
	}
	if o.ControlScores != nil {
		toSerialize["controlScores"] = o.ControlScores
	}
	if o.CreatedDateTime.IsSet() {
		toSerialize["createdDateTime"] = o.CreatedDateTime.Get()
	}
	if o.CurrentScore != nil {
		toSerialize["currentScore"] = o.CurrentScore
	}
	if o.EnabledServices != nil {
		toSerialize["enabledServices"] = o.EnabledServices
	}
	if o.LicensedUserCount.IsSet() {
		toSerialize["licensedUserCount"] = o.LicensedUserCount.Get()
	}
	if o.MaxScore != nil {
		toSerialize["maxScore"] = o.MaxScore
	}
	if o.VendorInformation != nil {
		toSerialize["vendorInformation"] = o.VendorInformation
	}
	return json.Marshal(toSerialize)
}

type NullableSecureScore struct {
	value *SecureScore
	isSet bool
}

func (v NullableSecureScore) Get() *SecureScore {
	return v.value
}

func (v *NullableSecureScore) Set(val *SecureScore) {
	v.value = val
	v.isSet = true
}

func (v NullableSecureScore) IsSet() bool {
	return v.isSet
}

func (v *NullableSecureScore) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecureScore(val *SecureScore) *NullableSecureScore {
	return &NullableSecureScore{value: val, isSet: true}
}

func (v NullableSecureScore) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecureScore) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


