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

// EducationSchool struct for EducationSchool
type EducationSchool struct {
	// Address of the school.
	Address AnyOfmicrosoftGraphPhysicalAddress `json:"address,omitempty"`
	// Entity who created the school.
	CreatedBy AnyOfmicrosoftGraphIdentitySet `json:"createdBy,omitempty"`
	// ID of school in syncing system.
	ExternalId NullableString `json:"externalId,omitempty"`
	// ID of principal in syncing system.
	ExternalPrincipalId NullableString `json:"externalPrincipalId,omitempty"`
	Fax NullableString `json:"fax,omitempty"`
	// Highest grade taught.
	HighestGrade NullableString `json:"highestGrade,omitempty"`
	// Lowest grade taught.
	LowestGrade NullableString `json:"lowestGrade,omitempty"`
	// Phone number of school.
	Phone NullableString `json:"phone,omitempty"`
	// Email address of the principal.
	PrincipalEmail NullableString `json:"principalEmail,omitempty"`
	// Name of the principal.
	PrincipalName NullableString `json:"principalName,omitempty"`
	// School Number.
	SchoolNumber NullableString `json:"schoolNumber,omitempty"`
	// The underlying administrativeUnit for this school.
	AdministrativeUnit AnyOfmicrosoftGraphAdministrativeUnit `json:"administrativeUnit,omitempty"`
	// Classes taught at the school. Nullable.
	Classes *[]MicrosoftGraphEducationClass `json:"classes,omitempty"`
	// Users in the school. Nullable.
	Users *[]MicrosoftGraphEducationUser `json:"users,omitempty"`
}

// NewEducationSchool instantiates a new EducationSchool object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEducationSchool() *EducationSchool {
	this := EducationSchool{}
	return &this
}

// NewEducationSchoolWithDefaults instantiates a new EducationSchool object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEducationSchoolWithDefaults() *EducationSchool {
	this := EducationSchool{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EducationSchool) GetAddress() AnyOfmicrosoftGraphPhysicalAddress {
	if o == nil  {
		var ret AnyOfmicrosoftGraphPhysicalAddress
		return ret
	}
	return o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EducationSchool) GetAddressOk() (*AnyOfmicrosoftGraphPhysicalAddress, bool) {
	if o == nil || o.Address == nil {
		return nil, false
	}
	return &o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *EducationSchool) HasAddress() bool {
	if o != nil && o.Address != nil {
		return true
	}

	return false
}

// SetAddress gets a reference to the given AnyOfmicrosoftGraphPhysicalAddress and assigns it to the Address field.
func (o *EducationSchool) SetAddress(v AnyOfmicrosoftGraphPhysicalAddress) {
	o.Address = v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EducationSchool) GetCreatedBy() AnyOfmicrosoftGraphIdentitySet {
	if o == nil  {
		var ret AnyOfmicrosoftGraphIdentitySet
		return ret
	}
	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EducationSchool) GetCreatedByOk() (*AnyOfmicrosoftGraphIdentitySet, bool) {
	if o == nil || o.CreatedBy == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *EducationSchool) HasCreatedBy() bool {
	if o != nil && o.CreatedBy != nil {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given AnyOfmicrosoftGraphIdentitySet and assigns it to the CreatedBy field.
func (o *EducationSchool) SetCreatedBy(v AnyOfmicrosoftGraphIdentitySet) {
	o.CreatedBy = v
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EducationSchool) GetExternalId() string {
	if o == nil || o.ExternalId.Get() == nil {
		var ret string
		return ret
	}
	return *o.ExternalId.Get()
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EducationSchool) GetExternalIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ExternalId.Get(), o.ExternalId.IsSet()
}

// HasExternalId returns a boolean if a field has been set.
func (o *EducationSchool) HasExternalId() bool {
	if o != nil && o.ExternalId.IsSet() {
		return true
	}

	return false
}

// SetExternalId gets a reference to the given NullableString and assigns it to the ExternalId field.
func (o *EducationSchool) SetExternalId(v string) {
	o.ExternalId.Set(&v)
}
// SetExternalIdNil sets the value for ExternalId to be an explicit nil
func (o *EducationSchool) SetExternalIdNil() {
	o.ExternalId.Set(nil)
}

// UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
func (o *EducationSchool) UnsetExternalId() {
	o.ExternalId.Unset()
}

// GetExternalPrincipalId returns the ExternalPrincipalId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EducationSchool) GetExternalPrincipalId() string {
	if o == nil || o.ExternalPrincipalId.Get() == nil {
		var ret string
		return ret
	}
	return *o.ExternalPrincipalId.Get()
}

// GetExternalPrincipalIdOk returns a tuple with the ExternalPrincipalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EducationSchool) GetExternalPrincipalIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ExternalPrincipalId.Get(), o.ExternalPrincipalId.IsSet()
}

// HasExternalPrincipalId returns a boolean if a field has been set.
func (o *EducationSchool) HasExternalPrincipalId() bool {
	if o != nil && o.ExternalPrincipalId.IsSet() {
		return true
	}

	return false
}

// SetExternalPrincipalId gets a reference to the given NullableString and assigns it to the ExternalPrincipalId field.
func (o *EducationSchool) SetExternalPrincipalId(v string) {
	o.ExternalPrincipalId.Set(&v)
}
// SetExternalPrincipalIdNil sets the value for ExternalPrincipalId to be an explicit nil
func (o *EducationSchool) SetExternalPrincipalIdNil() {
	o.ExternalPrincipalId.Set(nil)
}

// UnsetExternalPrincipalId ensures that no value is present for ExternalPrincipalId, not even an explicit nil
func (o *EducationSchool) UnsetExternalPrincipalId() {
	o.ExternalPrincipalId.Unset()
}

// GetFax returns the Fax field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EducationSchool) GetFax() string {
	if o == nil || o.Fax.Get() == nil {
		var ret string
		return ret
	}
	return *o.Fax.Get()
}

// GetFaxOk returns a tuple with the Fax field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EducationSchool) GetFaxOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Fax.Get(), o.Fax.IsSet()
}

// HasFax returns a boolean if a field has been set.
func (o *EducationSchool) HasFax() bool {
	if o != nil && o.Fax.IsSet() {
		return true
	}

	return false
}

// SetFax gets a reference to the given NullableString and assigns it to the Fax field.
func (o *EducationSchool) SetFax(v string) {
	o.Fax.Set(&v)
}
// SetFaxNil sets the value for Fax to be an explicit nil
func (o *EducationSchool) SetFaxNil() {
	o.Fax.Set(nil)
}

// UnsetFax ensures that no value is present for Fax, not even an explicit nil
func (o *EducationSchool) UnsetFax() {
	o.Fax.Unset()
}

// GetHighestGrade returns the HighestGrade field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EducationSchool) GetHighestGrade() string {
	if o == nil || o.HighestGrade.Get() == nil {
		var ret string
		return ret
	}
	return *o.HighestGrade.Get()
}

// GetHighestGradeOk returns a tuple with the HighestGrade field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EducationSchool) GetHighestGradeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.HighestGrade.Get(), o.HighestGrade.IsSet()
}

// HasHighestGrade returns a boolean if a field has been set.
func (o *EducationSchool) HasHighestGrade() bool {
	if o != nil && o.HighestGrade.IsSet() {
		return true
	}

	return false
}

// SetHighestGrade gets a reference to the given NullableString and assigns it to the HighestGrade field.
func (o *EducationSchool) SetHighestGrade(v string) {
	o.HighestGrade.Set(&v)
}
// SetHighestGradeNil sets the value for HighestGrade to be an explicit nil
func (o *EducationSchool) SetHighestGradeNil() {
	o.HighestGrade.Set(nil)
}

// UnsetHighestGrade ensures that no value is present for HighestGrade, not even an explicit nil
func (o *EducationSchool) UnsetHighestGrade() {
	o.HighestGrade.Unset()
}

// GetLowestGrade returns the LowestGrade field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EducationSchool) GetLowestGrade() string {
	if o == nil || o.LowestGrade.Get() == nil {
		var ret string
		return ret
	}
	return *o.LowestGrade.Get()
}

// GetLowestGradeOk returns a tuple with the LowestGrade field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EducationSchool) GetLowestGradeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LowestGrade.Get(), o.LowestGrade.IsSet()
}

// HasLowestGrade returns a boolean if a field has been set.
func (o *EducationSchool) HasLowestGrade() bool {
	if o != nil && o.LowestGrade.IsSet() {
		return true
	}

	return false
}

// SetLowestGrade gets a reference to the given NullableString and assigns it to the LowestGrade field.
func (o *EducationSchool) SetLowestGrade(v string) {
	o.LowestGrade.Set(&v)
}
// SetLowestGradeNil sets the value for LowestGrade to be an explicit nil
func (o *EducationSchool) SetLowestGradeNil() {
	o.LowestGrade.Set(nil)
}

// UnsetLowestGrade ensures that no value is present for LowestGrade, not even an explicit nil
func (o *EducationSchool) UnsetLowestGrade() {
	o.LowestGrade.Unset()
}

// GetPhone returns the Phone field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EducationSchool) GetPhone() string {
	if o == nil || o.Phone.Get() == nil {
		var ret string
		return ret
	}
	return *o.Phone.Get()
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EducationSchool) GetPhoneOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Phone.Get(), o.Phone.IsSet()
}

// HasPhone returns a boolean if a field has been set.
func (o *EducationSchool) HasPhone() bool {
	if o != nil && o.Phone.IsSet() {
		return true
	}

	return false
}

// SetPhone gets a reference to the given NullableString and assigns it to the Phone field.
func (o *EducationSchool) SetPhone(v string) {
	o.Phone.Set(&v)
}
// SetPhoneNil sets the value for Phone to be an explicit nil
func (o *EducationSchool) SetPhoneNil() {
	o.Phone.Set(nil)
}

// UnsetPhone ensures that no value is present for Phone, not even an explicit nil
func (o *EducationSchool) UnsetPhone() {
	o.Phone.Unset()
}

// GetPrincipalEmail returns the PrincipalEmail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EducationSchool) GetPrincipalEmail() string {
	if o == nil || o.PrincipalEmail.Get() == nil {
		var ret string
		return ret
	}
	return *o.PrincipalEmail.Get()
}

// GetPrincipalEmailOk returns a tuple with the PrincipalEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EducationSchool) GetPrincipalEmailOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PrincipalEmail.Get(), o.PrincipalEmail.IsSet()
}

// HasPrincipalEmail returns a boolean if a field has been set.
func (o *EducationSchool) HasPrincipalEmail() bool {
	if o != nil && o.PrincipalEmail.IsSet() {
		return true
	}

	return false
}

// SetPrincipalEmail gets a reference to the given NullableString and assigns it to the PrincipalEmail field.
func (o *EducationSchool) SetPrincipalEmail(v string) {
	o.PrincipalEmail.Set(&v)
}
// SetPrincipalEmailNil sets the value for PrincipalEmail to be an explicit nil
func (o *EducationSchool) SetPrincipalEmailNil() {
	o.PrincipalEmail.Set(nil)
}

// UnsetPrincipalEmail ensures that no value is present for PrincipalEmail, not even an explicit nil
func (o *EducationSchool) UnsetPrincipalEmail() {
	o.PrincipalEmail.Unset()
}

// GetPrincipalName returns the PrincipalName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EducationSchool) GetPrincipalName() string {
	if o == nil || o.PrincipalName.Get() == nil {
		var ret string
		return ret
	}
	return *o.PrincipalName.Get()
}

// GetPrincipalNameOk returns a tuple with the PrincipalName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EducationSchool) GetPrincipalNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PrincipalName.Get(), o.PrincipalName.IsSet()
}

// HasPrincipalName returns a boolean if a field has been set.
func (o *EducationSchool) HasPrincipalName() bool {
	if o != nil && o.PrincipalName.IsSet() {
		return true
	}

	return false
}

// SetPrincipalName gets a reference to the given NullableString and assigns it to the PrincipalName field.
func (o *EducationSchool) SetPrincipalName(v string) {
	o.PrincipalName.Set(&v)
}
// SetPrincipalNameNil sets the value for PrincipalName to be an explicit nil
func (o *EducationSchool) SetPrincipalNameNil() {
	o.PrincipalName.Set(nil)
}

// UnsetPrincipalName ensures that no value is present for PrincipalName, not even an explicit nil
func (o *EducationSchool) UnsetPrincipalName() {
	o.PrincipalName.Unset()
}

// GetSchoolNumber returns the SchoolNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EducationSchool) GetSchoolNumber() string {
	if o == nil || o.SchoolNumber.Get() == nil {
		var ret string
		return ret
	}
	return *o.SchoolNumber.Get()
}

// GetSchoolNumberOk returns a tuple with the SchoolNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EducationSchool) GetSchoolNumberOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SchoolNumber.Get(), o.SchoolNumber.IsSet()
}

// HasSchoolNumber returns a boolean if a field has been set.
func (o *EducationSchool) HasSchoolNumber() bool {
	if o != nil && o.SchoolNumber.IsSet() {
		return true
	}

	return false
}

// SetSchoolNumber gets a reference to the given NullableString and assigns it to the SchoolNumber field.
func (o *EducationSchool) SetSchoolNumber(v string) {
	o.SchoolNumber.Set(&v)
}
// SetSchoolNumberNil sets the value for SchoolNumber to be an explicit nil
func (o *EducationSchool) SetSchoolNumberNil() {
	o.SchoolNumber.Set(nil)
}

// UnsetSchoolNumber ensures that no value is present for SchoolNumber, not even an explicit nil
func (o *EducationSchool) UnsetSchoolNumber() {
	o.SchoolNumber.Unset()
}

// GetAdministrativeUnit returns the AdministrativeUnit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EducationSchool) GetAdministrativeUnit() AnyOfmicrosoftGraphAdministrativeUnit {
	if o == nil  {
		var ret AnyOfmicrosoftGraphAdministrativeUnit
		return ret
	}
	return o.AdministrativeUnit
}

// GetAdministrativeUnitOk returns a tuple with the AdministrativeUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EducationSchool) GetAdministrativeUnitOk() (*AnyOfmicrosoftGraphAdministrativeUnit, bool) {
	if o == nil || o.AdministrativeUnit == nil {
		return nil, false
	}
	return &o.AdministrativeUnit, true
}

// HasAdministrativeUnit returns a boolean if a field has been set.
func (o *EducationSchool) HasAdministrativeUnit() bool {
	if o != nil && o.AdministrativeUnit != nil {
		return true
	}

	return false
}

// SetAdministrativeUnit gets a reference to the given AnyOfmicrosoftGraphAdministrativeUnit and assigns it to the AdministrativeUnit field.
func (o *EducationSchool) SetAdministrativeUnit(v AnyOfmicrosoftGraphAdministrativeUnit) {
	o.AdministrativeUnit = v
}

// GetClasses returns the Classes field value if set, zero value otherwise.
func (o *EducationSchool) GetClasses() []MicrosoftGraphEducationClass {
	if o == nil || o.Classes == nil {
		var ret []MicrosoftGraphEducationClass
		return ret
	}
	return *o.Classes
}

// GetClassesOk returns a tuple with the Classes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EducationSchool) GetClassesOk() (*[]MicrosoftGraphEducationClass, bool) {
	if o == nil || o.Classes == nil {
		return nil, false
	}
	return o.Classes, true
}

// HasClasses returns a boolean if a field has been set.
func (o *EducationSchool) HasClasses() bool {
	if o != nil && o.Classes != nil {
		return true
	}

	return false
}

// SetClasses gets a reference to the given []MicrosoftGraphEducationClass and assigns it to the Classes field.
func (o *EducationSchool) SetClasses(v []MicrosoftGraphEducationClass) {
	o.Classes = &v
}

// GetUsers returns the Users field value if set, zero value otherwise.
func (o *EducationSchool) GetUsers() []MicrosoftGraphEducationUser {
	if o == nil || o.Users == nil {
		var ret []MicrosoftGraphEducationUser
		return ret
	}
	return *o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EducationSchool) GetUsersOk() (*[]MicrosoftGraphEducationUser, bool) {
	if o == nil || o.Users == nil {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *EducationSchool) HasUsers() bool {
	if o != nil && o.Users != nil {
		return true
	}

	return false
}

// SetUsers gets a reference to the given []MicrosoftGraphEducationUser and assigns it to the Users field.
func (o *EducationSchool) SetUsers(v []MicrosoftGraphEducationUser) {
	o.Users = &v
}

func (o EducationSchool) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Address != nil {
		toSerialize["address"] = o.Address
	}
	if o.CreatedBy != nil {
		toSerialize["createdBy"] = o.CreatedBy
	}
	if o.ExternalId.IsSet() {
		toSerialize["externalId"] = o.ExternalId.Get()
	}
	if o.ExternalPrincipalId.IsSet() {
		toSerialize["externalPrincipalId"] = o.ExternalPrincipalId.Get()
	}
	if o.Fax.IsSet() {
		toSerialize["fax"] = o.Fax.Get()
	}
	if o.HighestGrade.IsSet() {
		toSerialize["highestGrade"] = o.HighestGrade.Get()
	}
	if o.LowestGrade.IsSet() {
		toSerialize["lowestGrade"] = o.LowestGrade.Get()
	}
	if o.Phone.IsSet() {
		toSerialize["phone"] = o.Phone.Get()
	}
	if o.PrincipalEmail.IsSet() {
		toSerialize["principalEmail"] = o.PrincipalEmail.Get()
	}
	if o.PrincipalName.IsSet() {
		toSerialize["principalName"] = o.PrincipalName.Get()
	}
	if o.SchoolNumber.IsSet() {
		toSerialize["schoolNumber"] = o.SchoolNumber.Get()
	}
	if o.AdministrativeUnit != nil {
		toSerialize["administrativeUnit"] = o.AdministrativeUnit
	}
	if o.Classes != nil {
		toSerialize["classes"] = o.Classes
	}
	if o.Users != nil {
		toSerialize["users"] = o.Users
	}
	return json.Marshal(toSerialize)
}

type NullableEducationSchool struct {
	value *EducationSchool
	isSet bool
}

func (v NullableEducationSchool) Get() *EducationSchool {
	return v.value
}

func (v *NullableEducationSchool) Set(val *EducationSchool) {
	v.value = val
	v.isSet = true
}

func (v NullableEducationSchool) IsSet() bool {
	return v.isSet
}

func (v *NullableEducationSchool) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEducationSchool(val *EducationSchool) *NullableEducationSchool {
	return &NullableEducationSchool{value: val, isSet: true}
}

func (v NullableEducationSchool) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEducationSchool) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


