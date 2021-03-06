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

// OrgContact struct for OrgContact
type OrgContact struct {
	// Postal addresses for this organizational contact. For now a contact can only have one physical address.
	Addresses *[]*AnyOfmicrosoftGraphPhysicalOfficeAddress `json:"addresses,omitempty"`
	// Name of the company that this organizational contact belong to. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values).
	CompanyName NullableString `json:"companyName,omitempty"`
	// The name for the department in which the contact works. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values).
	Department NullableString `json:"department,omitempty"`
	// Display name for this organizational contact. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values), $search, and $orderBy.
	DisplayName NullableString `json:"displayName,omitempty"`
	// First name for this organizational contact. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values).
	GivenName NullableString `json:"givenName,omitempty"`
	// Job title for this organizational contact. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values).
	JobTitle NullableString `json:"jobTitle,omitempty"`
	// The SMTP address for the contact, for example, 'jeff@contoso.onmicrosoft.com'. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values).
	Mail NullableString `json:"mail,omitempty"`
	// Email alias (portion of email address pre-pending the @ symbol) for this organizational contact. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values).
	MailNickname NullableString `json:"mailNickname,omitempty"`
	// Date and time when this organizational contact was last synchronized from on-premises AD. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Supports $filter (eq, ne, not, ge, le, in).
	OnPremisesLastSyncDateTime NullableTime `json:"onPremisesLastSyncDateTime,omitempty"`
	// List of any synchronization provisioning errors for this organizational contact. Supports $filter (eq, not).
	OnPremisesProvisioningErrors *[]*AnyOfmicrosoftGraphOnPremisesProvisioningError `json:"onPremisesProvisioningErrors,omitempty"`
	// true if this object is synced from an on-premises directory; false if this object was originally synced from an on-premises directory but is no longer synced and now mastered in Exchange; null if this object has never been synced from an on-premises directory (default).  Supports $filter (eq, ne, not, in, and eq on null values).
	OnPremisesSyncEnabled NullableBool `json:"onPremisesSyncEnabled,omitempty"`
	// List of phones for this organizational contact. Phone types can be mobile, business, and businessFax. Only one of each type can ever be present in the collection. Supports $filter (eq, ne, not, in).
	Phones *[]*AnyOfmicrosoftGraphPhone `json:"phones,omitempty"`
	// For example: 'SMTP: bob@contoso.com', 'smtp: bob@sales.contoso.com'. The any operator is required for filter expressions on multi-valued properties. Supports $filter (eq, not, ge, le, startsWith).
	ProxyAddresses *[]string `json:"proxyAddresses,omitempty"`
	// Last name for this organizational contact. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values)
	Surname NullableString `json:"surname,omitempty"`
	// The contact's direct reports. (The users and contacts that have their manager property set to this contact.) Read-only. Nullable. Supports $expand.
	DirectReports *[]MicrosoftGraphDirectoryObject `json:"directReports,omitempty"`
	// The user or contact that is this contact's manager. Read-only. Supports $expand.
	Manager AnyOfmicrosoftGraphDirectoryObject `json:"manager,omitempty"`
	// Groups that this contact is a member of. Read-only. Nullable. Supports $expand.
	MemberOf *[]MicrosoftGraphDirectoryObject `json:"memberOf,omitempty"`
	TransitiveMemberOf *[]MicrosoftGraphDirectoryObject `json:"transitiveMemberOf,omitempty"`
}

// NewOrgContact instantiates a new OrgContact object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrgContact() *OrgContact {
	this := OrgContact{}
	return &this
}

// NewOrgContactWithDefaults instantiates a new OrgContact object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrgContactWithDefaults() *OrgContact {
	this := OrgContact{}
	return &this
}

// GetAddresses returns the Addresses field value if set, zero value otherwise.
func (o *OrgContact) GetAddresses() []*AnyOfmicrosoftGraphPhysicalOfficeAddress {
	if o == nil || o.Addresses == nil {
		var ret []*AnyOfmicrosoftGraphPhysicalOfficeAddress
		return ret
	}
	return *o.Addresses
}

// GetAddressesOk returns a tuple with the Addresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgContact) GetAddressesOk() (*[]*AnyOfmicrosoftGraphPhysicalOfficeAddress, bool) {
	if o == nil || o.Addresses == nil {
		return nil, false
	}
	return o.Addresses, true
}

// HasAddresses returns a boolean if a field has been set.
func (o *OrgContact) HasAddresses() bool {
	if o != nil && o.Addresses != nil {
		return true
	}

	return false
}

// SetAddresses gets a reference to the given []*AnyOfmicrosoftGraphPhysicalOfficeAddress and assigns it to the Addresses field.
func (o *OrgContact) SetAddresses(v []*AnyOfmicrosoftGraphPhysicalOfficeAddress) {
	o.Addresses = &v
}

// GetCompanyName returns the CompanyName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrgContact) GetCompanyName() string {
	if o == nil || o.CompanyName.Get() == nil {
		var ret string
		return ret
	}
	return *o.CompanyName.Get()
}

// GetCompanyNameOk returns a tuple with the CompanyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrgContact) GetCompanyNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CompanyName.Get(), o.CompanyName.IsSet()
}

// HasCompanyName returns a boolean if a field has been set.
func (o *OrgContact) HasCompanyName() bool {
	if o != nil && o.CompanyName.IsSet() {
		return true
	}

	return false
}

// SetCompanyName gets a reference to the given NullableString and assigns it to the CompanyName field.
func (o *OrgContact) SetCompanyName(v string) {
	o.CompanyName.Set(&v)
}
// SetCompanyNameNil sets the value for CompanyName to be an explicit nil
func (o *OrgContact) SetCompanyNameNil() {
	o.CompanyName.Set(nil)
}

// UnsetCompanyName ensures that no value is present for CompanyName, not even an explicit nil
func (o *OrgContact) UnsetCompanyName() {
	o.CompanyName.Unset()
}

// GetDepartment returns the Department field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrgContact) GetDepartment() string {
	if o == nil || o.Department.Get() == nil {
		var ret string
		return ret
	}
	return *o.Department.Get()
}

// GetDepartmentOk returns a tuple with the Department field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrgContact) GetDepartmentOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Department.Get(), o.Department.IsSet()
}

// HasDepartment returns a boolean if a field has been set.
func (o *OrgContact) HasDepartment() bool {
	if o != nil && o.Department.IsSet() {
		return true
	}

	return false
}

// SetDepartment gets a reference to the given NullableString and assigns it to the Department field.
func (o *OrgContact) SetDepartment(v string) {
	o.Department.Set(&v)
}
// SetDepartmentNil sets the value for Department to be an explicit nil
func (o *OrgContact) SetDepartmentNil() {
	o.Department.Set(nil)
}

// UnsetDepartment ensures that no value is present for Department, not even an explicit nil
func (o *OrgContact) UnsetDepartment() {
	o.Department.Unset()
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrgContact) GetDisplayName() string {
	if o == nil || o.DisplayName.Get() == nil {
		var ret string
		return ret
	}
	return *o.DisplayName.Get()
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrgContact) GetDisplayNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DisplayName.Get(), o.DisplayName.IsSet()
}

// HasDisplayName returns a boolean if a field has been set.
func (o *OrgContact) HasDisplayName() bool {
	if o != nil && o.DisplayName.IsSet() {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given NullableString and assigns it to the DisplayName field.
func (o *OrgContact) SetDisplayName(v string) {
	o.DisplayName.Set(&v)
}
// SetDisplayNameNil sets the value for DisplayName to be an explicit nil
func (o *OrgContact) SetDisplayNameNil() {
	o.DisplayName.Set(nil)
}

// UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
func (o *OrgContact) UnsetDisplayName() {
	o.DisplayName.Unset()
}

// GetGivenName returns the GivenName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrgContact) GetGivenName() string {
	if o == nil || o.GivenName.Get() == nil {
		var ret string
		return ret
	}
	return *o.GivenName.Get()
}

// GetGivenNameOk returns a tuple with the GivenName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrgContact) GetGivenNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.GivenName.Get(), o.GivenName.IsSet()
}

// HasGivenName returns a boolean if a field has been set.
func (o *OrgContact) HasGivenName() bool {
	if o != nil && o.GivenName.IsSet() {
		return true
	}

	return false
}

// SetGivenName gets a reference to the given NullableString and assigns it to the GivenName field.
func (o *OrgContact) SetGivenName(v string) {
	o.GivenName.Set(&v)
}
// SetGivenNameNil sets the value for GivenName to be an explicit nil
func (o *OrgContact) SetGivenNameNil() {
	o.GivenName.Set(nil)
}

// UnsetGivenName ensures that no value is present for GivenName, not even an explicit nil
func (o *OrgContact) UnsetGivenName() {
	o.GivenName.Unset()
}

// GetJobTitle returns the JobTitle field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrgContact) GetJobTitle() string {
	if o == nil || o.JobTitle.Get() == nil {
		var ret string
		return ret
	}
	return *o.JobTitle.Get()
}

// GetJobTitleOk returns a tuple with the JobTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrgContact) GetJobTitleOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.JobTitle.Get(), o.JobTitle.IsSet()
}

// HasJobTitle returns a boolean if a field has been set.
func (o *OrgContact) HasJobTitle() bool {
	if o != nil && o.JobTitle.IsSet() {
		return true
	}

	return false
}

// SetJobTitle gets a reference to the given NullableString and assigns it to the JobTitle field.
func (o *OrgContact) SetJobTitle(v string) {
	o.JobTitle.Set(&v)
}
// SetJobTitleNil sets the value for JobTitle to be an explicit nil
func (o *OrgContact) SetJobTitleNil() {
	o.JobTitle.Set(nil)
}

// UnsetJobTitle ensures that no value is present for JobTitle, not even an explicit nil
func (o *OrgContact) UnsetJobTitle() {
	o.JobTitle.Unset()
}

// GetMail returns the Mail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrgContact) GetMail() string {
	if o == nil || o.Mail.Get() == nil {
		var ret string
		return ret
	}
	return *o.Mail.Get()
}

// GetMailOk returns a tuple with the Mail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrgContact) GetMailOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Mail.Get(), o.Mail.IsSet()
}

// HasMail returns a boolean if a field has been set.
func (o *OrgContact) HasMail() bool {
	if o != nil && o.Mail.IsSet() {
		return true
	}

	return false
}

// SetMail gets a reference to the given NullableString and assigns it to the Mail field.
func (o *OrgContact) SetMail(v string) {
	o.Mail.Set(&v)
}
// SetMailNil sets the value for Mail to be an explicit nil
func (o *OrgContact) SetMailNil() {
	o.Mail.Set(nil)
}

// UnsetMail ensures that no value is present for Mail, not even an explicit nil
func (o *OrgContact) UnsetMail() {
	o.Mail.Unset()
}

// GetMailNickname returns the MailNickname field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrgContact) GetMailNickname() string {
	if o == nil || o.MailNickname.Get() == nil {
		var ret string
		return ret
	}
	return *o.MailNickname.Get()
}

// GetMailNicknameOk returns a tuple with the MailNickname field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrgContact) GetMailNicknameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.MailNickname.Get(), o.MailNickname.IsSet()
}

// HasMailNickname returns a boolean if a field has been set.
func (o *OrgContact) HasMailNickname() bool {
	if o != nil && o.MailNickname.IsSet() {
		return true
	}

	return false
}

// SetMailNickname gets a reference to the given NullableString and assigns it to the MailNickname field.
func (o *OrgContact) SetMailNickname(v string) {
	o.MailNickname.Set(&v)
}
// SetMailNicknameNil sets the value for MailNickname to be an explicit nil
func (o *OrgContact) SetMailNicknameNil() {
	o.MailNickname.Set(nil)
}

// UnsetMailNickname ensures that no value is present for MailNickname, not even an explicit nil
func (o *OrgContact) UnsetMailNickname() {
	o.MailNickname.Unset()
}

// GetOnPremisesLastSyncDateTime returns the OnPremisesLastSyncDateTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrgContact) GetOnPremisesLastSyncDateTime() time.Time {
	if o == nil || o.OnPremisesLastSyncDateTime.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.OnPremisesLastSyncDateTime.Get()
}

// GetOnPremisesLastSyncDateTimeOk returns a tuple with the OnPremisesLastSyncDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrgContact) GetOnPremisesLastSyncDateTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.OnPremisesLastSyncDateTime.Get(), o.OnPremisesLastSyncDateTime.IsSet()
}

// HasOnPremisesLastSyncDateTime returns a boolean if a field has been set.
func (o *OrgContact) HasOnPremisesLastSyncDateTime() bool {
	if o != nil && o.OnPremisesLastSyncDateTime.IsSet() {
		return true
	}

	return false
}

// SetOnPremisesLastSyncDateTime gets a reference to the given NullableTime and assigns it to the OnPremisesLastSyncDateTime field.
func (o *OrgContact) SetOnPremisesLastSyncDateTime(v time.Time) {
	o.OnPremisesLastSyncDateTime.Set(&v)
}
// SetOnPremisesLastSyncDateTimeNil sets the value for OnPremisesLastSyncDateTime to be an explicit nil
func (o *OrgContact) SetOnPremisesLastSyncDateTimeNil() {
	o.OnPremisesLastSyncDateTime.Set(nil)
}

// UnsetOnPremisesLastSyncDateTime ensures that no value is present for OnPremisesLastSyncDateTime, not even an explicit nil
func (o *OrgContact) UnsetOnPremisesLastSyncDateTime() {
	o.OnPremisesLastSyncDateTime.Unset()
}

// GetOnPremisesProvisioningErrors returns the OnPremisesProvisioningErrors field value if set, zero value otherwise.
func (o *OrgContact) GetOnPremisesProvisioningErrors() []*AnyOfmicrosoftGraphOnPremisesProvisioningError {
	if o == nil || o.OnPremisesProvisioningErrors == nil {
		var ret []*AnyOfmicrosoftGraphOnPremisesProvisioningError
		return ret
	}
	return *o.OnPremisesProvisioningErrors
}

// GetOnPremisesProvisioningErrorsOk returns a tuple with the OnPremisesProvisioningErrors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgContact) GetOnPremisesProvisioningErrorsOk() (*[]*AnyOfmicrosoftGraphOnPremisesProvisioningError, bool) {
	if o == nil || o.OnPremisesProvisioningErrors == nil {
		return nil, false
	}
	return o.OnPremisesProvisioningErrors, true
}

// HasOnPremisesProvisioningErrors returns a boolean if a field has been set.
func (o *OrgContact) HasOnPremisesProvisioningErrors() bool {
	if o != nil && o.OnPremisesProvisioningErrors != nil {
		return true
	}

	return false
}

// SetOnPremisesProvisioningErrors gets a reference to the given []*AnyOfmicrosoftGraphOnPremisesProvisioningError and assigns it to the OnPremisesProvisioningErrors field.
func (o *OrgContact) SetOnPremisesProvisioningErrors(v []*AnyOfmicrosoftGraphOnPremisesProvisioningError) {
	o.OnPremisesProvisioningErrors = &v
}

// GetOnPremisesSyncEnabled returns the OnPremisesSyncEnabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrgContact) GetOnPremisesSyncEnabled() bool {
	if o == nil || o.OnPremisesSyncEnabled.Get() == nil {
		var ret bool
		return ret
	}
	return *o.OnPremisesSyncEnabled.Get()
}

// GetOnPremisesSyncEnabledOk returns a tuple with the OnPremisesSyncEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrgContact) GetOnPremisesSyncEnabledOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.OnPremisesSyncEnabled.Get(), o.OnPremisesSyncEnabled.IsSet()
}

// HasOnPremisesSyncEnabled returns a boolean if a field has been set.
func (o *OrgContact) HasOnPremisesSyncEnabled() bool {
	if o != nil && o.OnPremisesSyncEnabled.IsSet() {
		return true
	}

	return false
}

// SetOnPremisesSyncEnabled gets a reference to the given NullableBool and assigns it to the OnPremisesSyncEnabled field.
func (o *OrgContact) SetOnPremisesSyncEnabled(v bool) {
	o.OnPremisesSyncEnabled.Set(&v)
}
// SetOnPremisesSyncEnabledNil sets the value for OnPremisesSyncEnabled to be an explicit nil
func (o *OrgContact) SetOnPremisesSyncEnabledNil() {
	o.OnPremisesSyncEnabled.Set(nil)
}

// UnsetOnPremisesSyncEnabled ensures that no value is present for OnPremisesSyncEnabled, not even an explicit nil
func (o *OrgContact) UnsetOnPremisesSyncEnabled() {
	o.OnPremisesSyncEnabled.Unset()
}

// GetPhones returns the Phones field value if set, zero value otherwise.
func (o *OrgContact) GetPhones() []*AnyOfmicrosoftGraphPhone {
	if o == nil || o.Phones == nil {
		var ret []*AnyOfmicrosoftGraphPhone
		return ret
	}
	return *o.Phones
}

// GetPhonesOk returns a tuple with the Phones field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgContact) GetPhonesOk() (*[]*AnyOfmicrosoftGraphPhone, bool) {
	if o == nil || o.Phones == nil {
		return nil, false
	}
	return o.Phones, true
}

// HasPhones returns a boolean if a field has been set.
func (o *OrgContact) HasPhones() bool {
	if o != nil && o.Phones != nil {
		return true
	}

	return false
}

// SetPhones gets a reference to the given []*AnyOfmicrosoftGraphPhone and assigns it to the Phones field.
func (o *OrgContact) SetPhones(v []*AnyOfmicrosoftGraphPhone) {
	o.Phones = &v
}

// GetProxyAddresses returns the ProxyAddresses field value if set, zero value otherwise.
func (o *OrgContact) GetProxyAddresses() []string {
	if o == nil || o.ProxyAddresses == nil {
		var ret []string
		return ret
	}
	return *o.ProxyAddresses
}

// GetProxyAddressesOk returns a tuple with the ProxyAddresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgContact) GetProxyAddressesOk() (*[]string, bool) {
	if o == nil || o.ProxyAddresses == nil {
		return nil, false
	}
	return o.ProxyAddresses, true
}

// HasProxyAddresses returns a boolean if a field has been set.
func (o *OrgContact) HasProxyAddresses() bool {
	if o != nil && o.ProxyAddresses != nil {
		return true
	}

	return false
}

// SetProxyAddresses gets a reference to the given []string and assigns it to the ProxyAddresses field.
func (o *OrgContact) SetProxyAddresses(v []string) {
	o.ProxyAddresses = &v
}

// GetSurname returns the Surname field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrgContact) GetSurname() string {
	if o == nil || o.Surname.Get() == nil {
		var ret string
		return ret
	}
	return *o.Surname.Get()
}

// GetSurnameOk returns a tuple with the Surname field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrgContact) GetSurnameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Surname.Get(), o.Surname.IsSet()
}

// HasSurname returns a boolean if a field has been set.
func (o *OrgContact) HasSurname() bool {
	if o != nil && o.Surname.IsSet() {
		return true
	}

	return false
}

// SetSurname gets a reference to the given NullableString and assigns it to the Surname field.
func (o *OrgContact) SetSurname(v string) {
	o.Surname.Set(&v)
}
// SetSurnameNil sets the value for Surname to be an explicit nil
func (o *OrgContact) SetSurnameNil() {
	o.Surname.Set(nil)
}

// UnsetSurname ensures that no value is present for Surname, not even an explicit nil
func (o *OrgContact) UnsetSurname() {
	o.Surname.Unset()
}

// GetDirectReports returns the DirectReports field value if set, zero value otherwise.
func (o *OrgContact) GetDirectReports() []MicrosoftGraphDirectoryObject {
	if o == nil || o.DirectReports == nil {
		var ret []MicrosoftGraphDirectoryObject
		return ret
	}
	return *o.DirectReports
}

// GetDirectReportsOk returns a tuple with the DirectReports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgContact) GetDirectReportsOk() (*[]MicrosoftGraphDirectoryObject, bool) {
	if o == nil || o.DirectReports == nil {
		return nil, false
	}
	return o.DirectReports, true
}

// HasDirectReports returns a boolean if a field has been set.
func (o *OrgContact) HasDirectReports() bool {
	if o != nil && o.DirectReports != nil {
		return true
	}

	return false
}

// SetDirectReports gets a reference to the given []MicrosoftGraphDirectoryObject and assigns it to the DirectReports field.
func (o *OrgContact) SetDirectReports(v []MicrosoftGraphDirectoryObject) {
	o.DirectReports = &v
}

// GetManager returns the Manager field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrgContact) GetManager() AnyOfmicrosoftGraphDirectoryObject {
	if o == nil  {
		var ret AnyOfmicrosoftGraphDirectoryObject
		return ret
	}
	return o.Manager
}

// GetManagerOk returns a tuple with the Manager field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrgContact) GetManagerOk() (*AnyOfmicrosoftGraphDirectoryObject, bool) {
	if o == nil || o.Manager == nil {
		return nil, false
	}
	return &o.Manager, true
}

// HasManager returns a boolean if a field has been set.
func (o *OrgContact) HasManager() bool {
	if o != nil && o.Manager != nil {
		return true
	}

	return false
}

// SetManager gets a reference to the given AnyOfmicrosoftGraphDirectoryObject and assigns it to the Manager field.
func (o *OrgContact) SetManager(v AnyOfmicrosoftGraphDirectoryObject) {
	o.Manager = v
}

// GetMemberOf returns the MemberOf field value if set, zero value otherwise.
func (o *OrgContact) GetMemberOf() []MicrosoftGraphDirectoryObject {
	if o == nil || o.MemberOf == nil {
		var ret []MicrosoftGraphDirectoryObject
		return ret
	}
	return *o.MemberOf
}

// GetMemberOfOk returns a tuple with the MemberOf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgContact) GetMemberOfOk() (*[]MicrosoftGraphDirectoryObject, bool) {
	if o == nil || o.MemberOf == nil {
		return nil, false
	}
	return o.MemberOf, true
}

// HasMemberOf returns a boolean if a field has been set.
func (o *OrgContact) HasMemberOf() bool {
	if o != nil && o.MemberOf != nil {
		return true
	}

	return false
}

// SetMemberOf gets a reference to the given []MicrosoftGraphDirectoryObject and assigns it to the MemberOf field.
func (o *OrgContact) SetMemberOf(v []MicrosoftGraphDirectoryObject) {
	o.MemberOf = &v
}

// GetTransitiveMemberOf returns the TransitiveMemberOf field value if set, zero value otherwise.
func (o *OrgContact) GetTransitiveMemberOf() []MicrosoftGraphDirectoryObject {
	if o == nil || o.TransitiveMemberOf == nil {
		var ret []MicrosoftGraphDirectoryObject
		return ret
	}
	return *o.TransitiveMemberOf
}

// GetTransitiveMemberOfOk returns a tuple with the TransitiveMemberOf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgContact) GetTransitiveMemberOfOk() (*[]MicrosoftGraphDirectoryObject, bool) {
	if o == nil || o.TransitiveMemberOf == nil {
		return nil, false
	}
	return o.TransitiveMemberOf, true
}

// HasTransitiveMemberOf returns a boolean if a field has been set.
func (o *OrgContact) HasTransitiveMemberOf() bool {
	if o != nil && o.TransitiveMemberOf != nil {
		return true
	}

	return false
}

// SetTransitiveMemberOf gets a reference to the given []MicrosoftGraphDirectoryObject and assigns it to the TransitiveMemberOf field.
func (o *OrgContact) SetTransitiveMemberOf(v []MicrosoftGraphDirectoryObject) {
	o.TransitiveMemberOf = &v
}

func (o OrgContact) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Addresses != nil {
		toSerialize["addresses"] = o.Addresses
	}
	if o.CompanyName.IsSet() {
		toSerialize["companyName"] = o.CompanyName.Get()
	}
	if o.Department.IsSet() {
		toSerialize["department"] = o.Department.Get()
	}
	if o.DisplayName.IsSet() {
		toSerialize["displayName"] = o.DisplayName.Get()
	}
	if o.GivenName.IsSet() {
		toSerialize["givenName"] = o.GivenName.Get()
	}
	if o.JobTitle.IsSet() {
		toSerialize["jobTitle"] = o.JobTitle.Get()
	}
	if o.Mail.IsSet() {
		toSerialize["mail"] = o.Mail.Get()
	}
	if o.MailNickname.IsSet() {
		toSerialize["mailNickname"] = o.MailNickname.Get()
	}
	if o.OnPremisesLastSyncDateTime.IsSet() {
		toSerialize["onPremisesLastSyncDateTime"] = o.OnPremisesLastSyncDateTime.Get()
	}
	if o.OnPremisesProvisioningErrors != nil {
		toSerialize["onPremisesProvisioningErrors"] = o.OnPremisesProvisioningErrors
	}
	if o.OnPremisesSyncEnabled.IsSet() {
		toSerialize["onPremisesSyncEnabled"] = o.OnPremisesSyncEnabled.Get()
	}
	if o.Phones != nil {
		toSerialize["phones"] = o.Phones
	}
	if o.ProxyAddresses != nil {
		toSerialize["proxyAddresses"] = o.ProxyAddresses
	}
	if o.Surname.IsSet() {
		toSerialize["surname"] = o.Surname.Get()
	}
	if o.DirectReports != nil {
		toSerialize["directReports"] = o.DirectReports
	}
	if o.Manager != nil {
		toSerialize["manager"] = o.Manager
	}
	if o.MemberOf != nil {
		toSerialize["memberOf"] = o.MemberOf
	}
	if o.TransitiveMemberOf != nil {
		toSerialize["transitiveMemberOf"] = o.TransitiveMemberOf
	}
	return json.Marshal(toSerialize)
}

type NullableOrgContact struct {
	value *OrgContact
	isSet bool
}

func (v NullableOrgContact) Get() *OrgContact {
	return v.value
}

func (v *NullableOrgContact) Set(val *OrgContact) {
	v.value = val
	v.isSet = true
}

func (v NullableOrgContact) IsSet() bool {
	return v.isSet
}

func (v *NullableOrgContact) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrgContact(val *OrgContact) *NullableOrgContact {
	return &NullableOrgContact{value: val, isSet: true}
}

func (v NullableOrgContact) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrgContact) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


