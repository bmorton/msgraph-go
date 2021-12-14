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

// MicrosoftGraphEducationSchool struct for MicrosoftGraphEducationSchool
type MicrosoftGraphEducationSchool struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
	// Organization description.
	Description NullableString `json:"description,omitempty"`
	// Organization display name.
	DisplayName *string `json:"displayName,omitempty"`
	// Source where this organization was created from. Possible values are: sis, manual.
	ExternalSource AnyOfmicrosoftGraphEducationExternalSource `json:"externalSource,omitempty"`
	// The name of the external source this resources was generated from.
	ExternalSourceDetail NullableString `json:"externalSourceDetail,omitempty"`
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

// NewMicrosoftGraphEducationSchool instantiates a new MicrosoftGraphEducationSchool object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphEducationSchool() *MicrosoftGraphEducationSchool {
	this := MicrosoftGraphEducationSchool{}
	return &this
}

// NewMicrosoftGraphEducationSchoolWithDefaults instantiates a new MicrosoftGraphEducationSchool object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphEducationSchoolWithDefaults() *MicrosoftGraphEducationSchool {
	this := MicrosoftGraphEducationSchool{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MicrosoftGraphEducationSchool) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphEducationSchool) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphEducationSchool) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MicrosoftGraphEducationSchool) SetId(v string) {
	o.Id = &v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphEducationSchool) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphEducationSchool) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *MicrosoftGraphEducationSchool) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *MicrosoftGraphEducationSchool) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *MicrosoftGraphEducationSchool) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *MicrosoftGraphEducationSchool) UnsetDescription() {
	o.Description.Unset()
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *MicrosoftGraphEducationSchool) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphEducationSchool) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *MicrosoftGraphEducationSchool) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *MicrosoftGraphEducationSchool) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetExternalSource returns the ExternalSource field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphEducationSchool) GetExternalSource() AnyOfmicrosoftGraphEducationExternalSource {
	if o == nil  {
		var ret AnyOfmicrosoftGraphEducationExternalSource
		return ret
	}
	return o.ExternalSource
}

// GetExternalSourceOk returns a tuple with the ExternalSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphEducationSchool) GetExternalSourceOk() (*AnyOfmicrosoftGraphEducationExternalSource, bool) {
	if o == nil || o.ExternalSource == nil {
		return nil, false
	}
	return &o.ExternalSource, true
}

// HasExternalSource returns a boolean if a field has been set.
func (o *MicrosoftGraphEducationSchool) HasExternalSource() bool {
	if o != nil && o.ExternalSource != nil {
		return true
	}

	return false
}

// SetExternalSource gets a reference to the given AnyOfmicrosoftGraphEducationExternalSource and assigns it to the ExternalSource field.
func (o *MicrosoftGraphEducationSchool) SetExternalSource(v AnyOfmicrosoftGraphEducationExternalSource) {
	o.ExternalSource = v
}

// GetExternalSourceDetail returns the ExternalSourceDetail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphEducationSchool) GetExternalSourceDetail() string {
	if o == nil || o.ExternalSourceDetail.Get() == nil {
		var ret string
		return ret
	}
	return *o.ExternalSourceDetail.Get()
}

// GetExternalSourceDetailOk returns a tuple with the ExternalSourceDetail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphEducationSchool) GetExternalSourceDetailOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ExternalSourceDetail.Get(), o.ExternalSourceDetail.IsSet()
}

// HasExternalSourceDetail returns a boolean if a field has been set.
func (o *MicrosoftGraphEducationSchool) HasExternalSourceDetail() bool {
	if o != nil && o.ExternalSourceDetail.IsSet() {
		return true
	}

	return false
}

// SetExternalSourceDetail gets a reference to the given NullableString and assigns it to the ExternalSourceDetail field.
func (o *MicrosoftGraphEducationSchool) SetExternalSourceDetail(v string) {
	o.ExternalSourceDetail.Set(&v)
}
// SetExternalSourceDetailNil sets the value for ExternalSourceDetail to be an explicit nil
func (o *MicrosoftGraphEducationSchool) SetExternalSourceDetailNil() {
	o.ExternalSourceDetail.Set(nil)
}

// UnsetExternalSourceDetail ensures that no value is present for ExternalSourceDetail, not even an explicit nil
func (o *MicrosoftGraphEducationSchool) UnsetExternalSourceDetail() {
	o.ExternalSourceDetail.Unset()
}

// GetAddress returns the Address field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphEducationSchool) GetAddress() AnyOfmicrosoftGraphPhysicalAddress {
	if o == nil  {
		var ret AnyOfmicrosoftGraphPhysicalAddress
		return ret
	}
	return o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphEducationSchool) GetAddressOk() (*AnyOfmicrosoftGraphPhysicalAddress, bool) {
	if o == nil || o.Address == nil {
		return nil, false
	}
	return &o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *MicrosoftGraphEducationSchool) HasAddress() bool {
	if o != nil && o.Address != nil {
		return true
	}

	return false
}

// SetAddress gets a reference to the given AnyOfmicrosoftGraphPhysicalAddress and assigns it to the Address field.
func (o *MicrosoftGraphEducationSchool) SetAddress(v AnyOfmicrosoftGraphPhysicalAddress) {
	o.Address = v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphEducationSchool) GetCreatedBy() AnyOfmicrosoftGraphIdentitySet {
	if o == nil  {
		var ret AnyOfmicrosoftGraphIdentitySet
		return ret
	}
	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphEducationSchool) GetCreatedByOk() (*AnyOfmicrosoftGraphIdentitySet, bool) {
	if o == nil || o.CreatedBy == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *MicrosoftGraphEducationSchool) HasCreatedBy() bool {
	if o != nil && o.CreatedBy != nil {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given AnyOfmicrosoftGraphIdentitySet and assigns it to the CreatedBy field.
func (o *MicrosoftGraphEducationSchool) SetCreatedBy(v AnyOfmicrosoftGraphIdentitySet) {
	o.CreatedBy = v
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphEducationSchool) GetExternalId() string {
	if o == nil || o.ExternalId.Get() == nil {
		var ret string
		return ret
	}
	return *o.ExternalId.Get()
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphEducationSchool) GetExternalIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ExternalId.Get(), o.ExternalId.IsSet()
}

// HasExternalId returns a boolean if a field has been set.
func (o *MicrosoftGraphEducationSchool) HasExternalId() bool {
	if o != nil && o.ExternalId.IsSet() {
		return true
	}

	return false
}

// SetExternalId gets a reference to the given NullableString and assigns it to the ExternalId field.
func (o *MicrosoftGraphEducationSchool) SetExternalId(v string) {
	o.ExternalId.Set(&v)
}
// SetExternalIdNil sets the value for ExternalId to be an explicit nil
func (o *MicrosoftGraphEducationSchool) SetExternalIdNil() {
	o.ExternalId.Set(nil)
}

// UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
func (o *MicrosoftGraphEducationSchool) UnsetExternalId() {
	o.ExternalId.Unset()
}

// GetExternalPrincipalId returns the ExternalPrincipalId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphEducationSchool) GetExternalPrincipalId() string {
	if o == nil || o.ExternalPrincipalId.Get() == nil {
		var ret string
		return ret
	}
	return *o.ExternalPrincipalId.Get()
}

// GetExternalPrincipalIdOk returns a tuple with the ExternalPrincipalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphEducationSchool) GetExternalPrincipalIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ExternalPrincipalId.Get(), o.ExternalPrincipalId.IsSet()
}

// HasExternalPrincipalId returns a boolean if a field has been set.
func (o *MicrosoftGraphEducationSchool) HasExternalPrincipalId() bool {
	if o != nil && o.ExternalPrincipalId.IsSet() {
		return true
	}

	return false
}

// SetExternalPrincipalId gets a reference to the given NullableString and assigns it to the ExternalPrincipalId field.
func (o *MicrosoftGraphEducationSchool) SetExternalPrincipalId(v string) {
	o.ExternalPrincipalId.Set(&v)
}
// SetExternalPrincipalIdNil sets the value for ExternalPrincipalId to be an explicit nil
func (o *MicrosoftGraphEducationSchool) SetExternalPrincipalIdNil() {
	o.ExternalPrincipalId.Set(nil)
}

// UnsetExternalPrincipalId ensures that no value is present for ExternalPrincipalId, not even an explicit nil
func (o *MicrosoftGraphEducationSchool) UnsetExternalPrincipalId() {
	o.ExternalPrincipalId.Unset()
}

// GetFax returns the Fax field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphEducationSchool) GetFax() string {
	if o == nil || o.Fax.Get() == nil {
		var ret string
		return ret
	}
	return *o.Fax.Get()
}

// GetFaxOk returns a tuple with the Fax field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphEducationSchool) GetFaxOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Fax.Get(), o.Fax.IsSet()
}

// HasFax returns a boolean if a field has been set.
func (o *MicrosoftGraphEducationSchool) HasFax() bool {
	if o != nil && o.Fax.IsSet() {
		return true
	}

	return false
}

// SetFax gets a reference to the given NullableString and assigns it to the Fax field.
func (o *MicrosoftGraphEducationSchool) SetFax(v string) {
	o.Fax.Set(&v)
}
// SetFaxNil sets the value for Fax to be an explicit nil
func (o *MicrosoftGraphEducationSchool) SetFaxNil() {
	o.Fax.Set(nil)
}

// UnsetFax ensures that no value is present for Fax, not even an explicit nil
func (o *MicrosoftGraphEducationSchool) UnsetFax() {
	o.Fax.Unset()
}

// GetHighestGrade returns the HighestGrade field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphEducationSchool) GetHighestGrade() string {
	if o == nil || o.HighestGrade.Get() == nil {
		var ret string
		return ret
	}
	return *o.HighestGrade.Get()
}

// GetHighestGradeOk returns a tuple with the HighestGrade field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphEducationSchool) GetHighestGradeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.HighestGrade.Get(), o.HighestGrade.IsSet()
}

// HasHighestGrade returns a boolean if a field has been set.
func (o *MicrosoftGraphEducationSchool) HasHighestGrade() bool {
	if o != nil && o.HighestGrade.IsSet() {
		return true
	}

	return false
}

// SetHighestGrade gets a reference to the given NullableString and assigns it to the HighestGrade field.
func (o *MicrosoftGraphEducationSchool) SetHighestGrade(v string) {
	o.HighestGrade.Set(&v)
}
// SetHighestGradeNil sets the value for HighestGrade to be an explicit nil
func (o *MicrosoftGraphEducationSchool) SetHighestGradeNil() {
	o.HighestGrade.Set(nil)
}

// UnsetHighestGrade ensures that no value is present for HighestGrade, not even an explicit nil
func (o *MicrosoftGraphEducationSchool) UnsetHighestGrade() {
	o.HighestGrade.Unset()
}

// GetLowestGrade returns the LowestGrade field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphEducationSchool) GetLowestGrade() string {
	if o == nil || o.LowestGrade.Get() == nil {
		var ret string
		return ret
	}
	return *o.LowestGrade.Get()
}

// GetLowestGradeOk returns a tuple with the LowestGrade field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphEducationSchool) GetLowestGradeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LowestGrade.Get(), o.LowestGrade.IsSet()
}

// HasLowestGrade returns a boolean if a field has been set.
func (o *MicrosoftGraphEducationSchool) HasLowestGrade() bool {
	if o != nil && o.LowestGrade.IsSet() {
		return true
	}

	return false
}

// SetLowestGrade gets a reference to the given NullableString and assigns it to the LowestGrade field.
func (o *MicrosoftGraphEducationSchool) SetLowestGrade(v string) {
	o.LowestGrade.Set(&v)
}
// SetLowestGradeNil sets the value for LowestGrade to be an explicit nil
func (o *MicrosoftGraphEducationSchool) SetLowestGradeNil() {
	o.LowestGrade.Set(nil)
}

// UnsetLowestGrade ensures that no value is present for LowestGrade, not even an explicit nil
func (o *MicrosoftGraphEducationSchool) UnsetLowestGrade() {
	o.LowestGrade.Unset()
}

// GetPhone returns the Phone field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphEducationSchool) GetPhone() string {
	if o == nil || o.Phone.Get() == nil {
		var ret string
		return ret
	}
	return *o.Phone.Get()
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphEducationSchool) GetPhoneOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Phone.Get(), o.Phone.IsSet()
}

// HasPhone returns a boolean if a field has been set.
func (o *MicrosoftGraphEducationSchool) HasPhone() bool {
	if o != nil && o.Phone.IsSet() {
		return true
	}

	return false
}

// SetPhone gets a reference to the given NullableString and assigns it to the Phone field.
func (o *MicrosoftGraphEducationSchool) SetPhone(v string) {
	o.Phone.Set(&v)
}
// SetPhoneNil sets the value for Phone to be an explicit nil
func (o *MicrosoftGraphEducationSchool) SetPhoneNil() {
	o.Phone.Set(nil)
}

// UnsetPhone ensures that no value is present for Phone, not even an explicit nil
func (o *MicrosoftGraphEducationSchool) UnsetPhone() {
	o.Phone.Unset()
}

// GetPrincipalEmail returns the PrincipalEmail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphEducationSchool) GetPrincipalEmail() string {
	if o == nil || o.PrincipalEmail.Get() == nil {
		var ret string
		return ret
	}
	return *o.PrincipalEmail.Get()
}

// GetPrincipalEmailOk returns a tuple with the PrincipalEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphEducationSchool) GetPrincipalEmailOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PrincipalEmail.Get(), o.PrincipalEmail.IsSet()
}

// HasPrincipalEmail returns a boolean if a field has been set.
func (o *MicrosoftGraphEducationSchool) HasPrincipalEmail() bool {
	if o != nil && o.PrincipalEmail.IsSet() {
		return true
	}

	return false
}

// SetPrincipalEmail gets a reference to the given NullableString and assigns it to the PrincipalEmail field.
func (o *MicrosoftGraphEducationSchool) SetPrincipalEmail(v string) {
	o.PrincipalEmail.Set(&v)
}
// SetPrincipalEmailNil sets the value for PrincipalEmail to be an explicit nil
func (o *MicrosoftGraphEducationSchool) SetPrincipalEmailNil() {
	o.PrincipalEmail.Set(nil)
}

// UnsetPrincipalEmail ensures that no value is present for PrincipalEmail, not even an explicit nil
func (o *MicrosoftGraphEducationSchool) UnsetPrincipalEmail() {
	o.PrincipalEmail.Unset()
}

// GetPrincipalName returns the PrincipalName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphEducationSchool) GetPrincipalName() string {
	if o == nil || o.PrincipalName.Get() == nil {
		var ret string
		return ret
	}
	return *o.PrincipalName.Get()
}

// GetPrincipalNameOk returns a tuple with the PrincipalName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphEducationSchool) GetPrincipalNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PrincipalName.Get(), o.PrincipalName.IsSet()
}

// HasPrincipalName returns a boolean if a field has been set.
func (o *MicrosoftGraphEducationSchool) HasPrincipalName() bool {
	if o != nil && o.PrincipalName.IsSet() {
		return true
	}

	return false
}

// SetPrincipalName gets a reference to the given NullableString and assigns it to the PrincipalName field.
func (o *MicrosoftGraphEducationSchool) SetPrincipalName(v string) {
	o.PrincipalName.Set(&v)
}
// SetPrincipalNameNil sets the value for PrincipalName to be an explicit nil
func (o *MicrosoftGraphEducationSchool) SetPrincipalNameNil() {
	o.PrincipalName.Set(nil)
}

// UnsetPrincipalName ensures that no value is present for PrincipalName, not even an explicit nil
func (o *MicrosoftGraphEducationSchool) UnsetPrincipalName() {
	o.PrincipalName.Unset()
}

// GetSchoolNumber returns the SchoolNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphEducationSchool) GetSchoolNumber() string {
	if o == nil || o.SchoolNumber.Get() == nil {
		var ret string
		return ret
	}
	return *o.SchoolNumber.Get()
}

// GetSchoolNumberOk returns a tuple with the SchoolNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphEducationSchool) GetSchoolNumberOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SchoolNumber.Get(), o.SchoolNumber.IsSet()
}

// HasSchoolNumber returns a boolean if a field has been set.
func (o *MicrosoftGraphEducationSchool) HasSchoolNumber() bool {
	if o != nil && o.SchoolNumber.IsSet() {
		return true
	}

	return false
}

// SetSchoolNumber gets a reference to the given NullableString and assigns it to the SchoolNumber field.
func (o *MicrosoftGraphEducationSchool) SetSchoolNumber(v string) {
	o.SchoolNumber.Set(&v)
}
// SetSchoolNumberNil sets the value for SchoolNumber to be an explicit nil
func (o *MicrosoftGraphEducationSchool) SetSchoolNumberNil() {
	o.SchoolNumber.Set(nil)
}

// UnsetSchoolNumber ensures that no value is present for SchoolNumber, not even an explicit nil
func (o *MicrosoftGraphEducationSchool) UnsetSchoolNumber() {
	o.SchoolNumber.Unset()
}

// GetAdministrativeUnit returns the AdministrativeUnit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphEducationSchool) GetAdministrativeUnit() AnyOfmicrosoftGraphAdministrativeUnit {
	if o == nil  {
		var ret AnyOfmicrosoftGraphAdministrativeUnit
		return ret
	}
	return o.AdministrativeUnit
}

// GetAdministrativeUnitOk returns a tuple with the AdministrativeUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphEducationSchool) GetAdministrativeUnitOk() (*AnyOfmicrosoftGraphAdministrativeUnit, bool) {
	if o == nil || o.AdministrativeUnit == nil {
		return nil, false
	}
	return &o.AdministrativeUnit, true
}

// HasAdministrativeUnit returns a boolean if a field has been set.
func (o *MicrosoftGraphEducationSchool) HasAdministrativeUnit() bool {
	if o != nil && o.AdministrativeUnit != nil {
		return true
	}

	return false
}

// SetAdministrativeUnit gets a reference to the given AnyOfmicrosoftGraphAdministrativeUnit and assigns it to the AdministrativeUnit field.
func (o *MicrosoftGraphEducationSchool) SetAdministrativeUnit(v AnyOfmicrosoftGraphAdministrativeUnit) {
	o.AdministrativeUnit = v
}

// GetClasses returns the Classes field value if set, zero value otherwise.
func (o *MicrosoftGraphEducationSchool) GetClasses() []MicrosoftGraphEducationClass {
	if o == nil || o.Classes == nil {
		var ret []MicrosoftGraphEducationClass
		return ret
	}
	return *o.Classes
}

// GetClassesOk returns a tuple with the Classes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphEducationSchool) GetClassesOk() (*[]MicrosoftGraphEducationClass, bool) {
	if o == nil || o.Classes == nil {
		return nil, false
	}
	return o.Classes, true
}

// HasClasses returns a boolean if a field has been set.
func (o *MicrosoftGraphEducationSchool) HasClasses() bool {
	if o != nil && o.Classes != nil {
		return true
	}

	return false
}

// SetClasses gets a reference to the given []MicrosoftGraphEducationClass and assigns it to the Classes field.
func (o *MicrosoftGraphEducationSchool) SetClasses(v []MicrosoftGraphEducationClass) {
	o.Classes = &v
}

// GetUsers returns the Users field value if set, zero value otherwise.
func (o *MicrosoftGraphEducationSchool) GetUsers() []MicrosoftGraphEducationUser {
	if o == nil || o.Users == nil {
		var ret []MicrosoftGraphEducationUser
		return ret
	}
	return *o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphEducationSchool) GetUsersOk() (*[]MicrosoftGraphEducationUser, bool) {
	if o == nil || o.Users == nil {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *MicrosoftGraphEducationSchool) HasUsers() bool {
	if o != nil && o.Users != nil {
		return true
	}

	return false
}

// SetUsers gets a reference to the given []MicrosoftGraphEducationUser and assigns it to the Users field.
func (o *MicrosoftGraphEducationSchool) SetUsers(v []MicrosoftGraphEducationUser) {
	o.Users = &v
}

func (o MicrosoftGraphEducationSchool) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.DisplayName != nil {
		toSerialize["displayName"] = o.DisplayName
	}
	if o.ExternalSource != nil {
		toSerialize["externalSource"] = o.ExternalSource
	}
	if o.ExternalSourceDetail.IsSet() {
		toSerialize["externalSourceDetail"] = o.ExternalSourceDetail.Get()
	}
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

type NullableMicrosoftGraphEducationSchool struct {
	value *MicrosoftGraphEducationSchool
	isSet bool
}

func (v NullableMicrosoftGraphEducationSchool) Get() *MicrosoftGraphEducationSchool {
	return v.value
}

func (v *NullableMicrosoftGraphEducationSchool) Set(val *MicrosoftGraphEducationSchool) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphEducationSchool) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphEducationSchool) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphEducationSchool(val *MicrosoftGraphEducationSchool) *NullableMicrosoftGraphEducationSchool {
	return &NullableMicrosoftGraphEducationSchool{value: val, isSet: true}
}

func (v NullableMicrosoftGraphEducationSchool) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphEducationSchool) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


