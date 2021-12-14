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

// MicrosoftGraphDomain struct for MicrosoftGraphDomain
type MicrosoftGraphDomain struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
	// Indicates the configured authentication type for the domain. The value is either Managed or Federated. Managed indicates a cloud managed domain where Azure AD performs user authentication. Federated indicates authentication is federated with an identity provider such as the tenant's on-premises Active Directory via Active Directory Federation Services. This property is read-only and is not nullable.
	AuthenticationType *string `json:"authenticationType,omitempty"`
	// This property is always null except when the verify action is used. When the verify action is used, a domain entity is returned in the response. The availabilityStatus property of the domain entity in the response is either AvailableImmediately or EmailVerifiedDomainTakeoverScheduled.
	AvailabilityStatus NullableString `json:"availabilityStatus,omitempty"`
	// The value of the property is false if the DNS record management of the domain has been delegated to Microsoft 365. Otherwise, the value is true. Not nullable
	IsAdminManaged *bool `json:"isAdminManaged,omitempty"`
	// true if this is the default domain that is used for user creation. There is only one default domain per company. Not nullable
	IsDefault *bool `json:"isDefault,omitempty"`
	// true if this is the initial domain created by Microsoft Online Services (companyname.onmicrosoft.com). There is only one initial domain per company. Not nullable
	IsInitial *bool `json:"isInitial,omitempty"`
	// true if the domain is a verified root domain. Otherwise, false if the domain is a subdomain or unverified. Not nullable
	IsRoot *bool `json:"isRoot,omitempty"`
	// true if the domain has completed domain ownership verification. Not nullable
	IsVerified *bool `json:"isVerified,omitempty"`
	Manufacturer NullableString `json:"manufacturer,omitempty"`
	Model NullableString `json:"model,omitempty"`
	// Specifies the number of days before a user receives notification that their password will expire. If the property is not set, a default value of 14 days will be used.
	PasswordNotificationWindowInDays NullableInt32 `json:"passwordNotificationWindowInDays,omitempty"`
	// Specifies the length of time that a password is valid before it must be changed. If the property is not set, a default value of 90 days will be used.
	PasswordValidityPeriodInDays NullableInt32 `json:"passwordValidityPeriodInDays,omitempty"`
	// Status of asynchronous operations scheduled for the domain.
	State AnyOfmicrosoftGraphDomainState `json:"state,omitempty"`
	// The capabilities assigned to the domain. Can include 0, 1 or more of following values: Email, Sharepoint, EmailInternalRelayOnly, OfficeCommunicationsOnline, SharePointDefaultDomain, FullRedelegation, SharePointPublic, OrgIdAuthentication, Yammer, Intune. The values which you can add/remove using Graph API include: Email, OfficeCommunicationsOnline, Yammer. Not nullable
	SupportedServices *[]string `json:"supportedServices,omitempty"`
	// Read-only, Nullable
	DomainNameReferences *[]MicrosoftGraphDirectoryObject `json:"domainNameReferences,omitempty"`
	// DNS records the customer adds to the DNS zone file of the domain before the domain can be used by Microsoft Online services. Read-only, Nullable
	ServiceConfigurationRecords *[]MicrosoftGraphDomainDnsRecord `json:"serviceConfigurationRecords,omitempty"`
	// DNS records that the customer adds to the DNS zone file of the domain before the customer can complete domain ownership verification with Azure AD. Read-only, Nullable
	VerificationDnsRecords *[]MicrosoftGraphDomainDnsRecord `json:"verificationDnsRecords,omitempty"`
}

// NewMicrosoftGraphDomain instantiates a new MicrosoftGraphDomain object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphDomain() *MicrosoftGraphDomain {
	this := MicrosoftGraphDomain{}
	return &this
}

// NewMicrosoftGraphDomainWithDefaults instantiates a new MicrosoftGraphDomain object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphDomainWithDefaults() *MicrosoftGraphDomain {
	this := MicrosoftGraphDomain{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MicrosoftGraphDomain) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphDomain) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphDomain) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MicrosoftGraphDomain) SetId(v string) {
	o.Id = &v
}

// GetAuthenticationType returns the AuthenticationType field value if set, zero value otherwise.
func (o *MicrosoftGraphDomain) GetAuthenticationType() string {
	if o == nil || o.AuthenticationType == nil {
		var ret string
		return ret
	}
	return *o.AuthenticationType
}

// GetAuthenticationTypeOk returns a tuple with the AuthenticationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphDomain) GetAuthenticationTypeOk() (*string, bool) {
	if o == nil || o.AuthenticationType == nil {
		return nil, false
	}
	return o.AuthenticationType, true
}

// HasAuthenticationType returns a boolean if a field has been set.
func (o *MicrosoftGraphDomain) HasAuthenticationType() bool {
	if o != nil && o.AuthenticationType != nil {
		return true
	}

	return false
}

// SetAuthenticationType gets a reference to the given string and assigns it to the AuthenticationType field.
func (o *MicrosoftGraphDomain) SetAuthenticationType(v string) {
	o.AuthenticationType = &v
}

// GetAvailabilityStatus returns the AvailabilityStatus field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphDomain) GetAvailabilityStatus() string {
	if o == nil || o.AvailabilityStatus.Get() == nil {
		var ret string
		return ret
	}
	return *o.AvailabilityStatus.Get()
}

// GetAvailabilityStatusOk returns a tuple with the AvailabilityStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphDomain) GetAvailabilityStatusOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AvailabilityStatus.Get(), o.AvailabilityStatus.IsSet()
}

// HasAvailabilityStatus returns a boolean if a field has been set.
func (o *MicrosoftGraphDomain) HasAvailabilityStatus() bool {
	if o != nil && o.AvailabilityStatus.IsSet() {
		return true
	}

	return false
}

// SetAvailabilityStatus gets a reference to the given NullableString and assigns it to the AvailabilityStatus field.
func (o *MicrosoftGraphDomain) SetAvailabilityStatus(v string) {
	o.AvailabilityStatus.Set(&v)
}
// SetAvailabilityStatusNil sets the value for AvailabilityStatus to be an explicit nil
func (o *MicrosoftGraphDomain) SetAvailabilityStatusNil() {
	o.AvailabilityStatus.Set(nil)
}

// UnsetAvailabilityStatus ensures that no value is present for AvailabilityStatus, not even an explicit nil
func (o *MicrosoftGraphDomain) UnsetAvailabilityStatus() {
	o.AvailabilityStatus.Unset()
}

// GetIsAdminManaged returns the IsAdminManaged field value if set, zero value otherwise.
func (o *MicrosoftGraphDomain) GetIsAdminManaged() bool {
	if o == nil || o.IsAdminManaged == nil {
		var ret bool
		return ret
	}
	return *o.IsAdminManaged
}

// GetIsAdminManagedOk returns a tuple with the IsAdminManaged field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphDomain) GetIsAdminManagedOk() (*bool, bool) {
	if o == nil || o.IsAdminManaged == nil {
		return nil, false
	}
	return o.IsAdminManaged, true
}

// HasIsAdminManaged returns a boolean if a field has been set.
func (o *MicrosoftGraphDomain) HasIsAdminManaged() bool {
	if o != nil && o.IsAdminManaged != nil {
		return true
	}

	return false
}

// SetIsAdminManaged gets a reference to the given bool and assigns it to the IsAdminManaged field.
func (o *MicrosoftGraphDomain) SetIsAdminManaged(v bool) {
	o.IsAdminManaged = &v
}

// GetIsDefault returns the IsDefault field value if set, zero value otherwise.
func (o *MicrosoftGraphDomain) GetIsDefault() bool {
	if o == nil || o.IsDefault == nil {
		var ret bool
		return ret
	}
	return *o.IsDefault
}

// GetIsDefaultOk returns a tuple with the IsDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphDomain) GetIsDefaultOk() (*bool, bool) {
	if o == nil || o.IsDefault == nil {
		return nil, false
	}
	return o.IsDefault, true
}

// HasIsDefault returns a boolean if a field has been set.
func (o *MicrosoftGraphDomain) HasIsDefault() bool {
	if o != nil && o.IsDefault != nil {
		return true
	}

	return false
}

// SetIsDefault gets a reference to the given bool and assigns it to the IsDefault field.
func (o *MicrosoftGraphDomain) SetIsDefault(v bool) {
	o.IsDefault = &v
}

// GetIsInitial returns the IsInitial field value if set, zero value otherwise.
func (o *MicrosoftGraphDomain) GetIsInitial() bool {
	if o == nil || o.IsInitial == nil {
		var ret bool
		return ret
	}
	return *o.IsInitial
}

// GetIsInitialOk returns a tuple with the IsInitial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphDomain) GetIsInitialOk() (*bool, bool) {
	if o == nil || o.IsInitial == nil {
		return nil, false
	}
	return o.IsInitial, true
}

// HasIsInitial returns a boolean if a field has been set.
func (o *MicrosoftGraphDomain) HasIsInitial() bool {
	if o != nil && o.IsInitial != nil {
		return true
	}

	return false
}

// SetIsInitial gets a reference to the given bool and assigns it to the IsInitial field.
func (o *MicrosoftGraphDomain) SetIsInitial(v bool) {
	o.IsInitial = &v
}

// GetIsRoot returns the IsRoot field value if set, zero value otherwise.
func (o *MicrosoftGraphDomain) GetIsRoot() bool {
	if o == nil || o.IsRoot == nil {
		var ret bool
		return ret
	}
	return *o.IsRoot
}

// GetIsRootOk returns a tuple with the IsRoot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphDomain) GetIsRootOk() (*bool, bool) {
	if o == nil || o.IsRoot == nil {
		return nil, false
	}
	return o.IsRoot, true
}

// HasIsRoot returns a boolean if a field has been set.
func (o *MicrosoftGraphDomain) HasIsRoot() bool {
	if o != nil && o.IsRoot != nil {
		return true
	}

	return false
}

// SetIsRoot gets a reference to the given bool and assigns it to the IsRoot field.
func (o *MicrosoftGraphDomain) SetIsRoot(v bool) {
	o.IsRoot = &v
}

// GetIsVerified returns the IsVerified field value if set, zero value otherwise.
func (o *MicrosoftGraphDomain) GetIsVerified() bool {
	if o == nil || o.IsVerified == nil {
		var ret bool
		return ret
	}
	return *o.IsVerified
}

// GetIsVerifiedOk returns a tuple with the IsVerified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphDomain) GetIsVerifiedOk() (*bool, bool) {
	if o == nil || o.IsVerified == nil {
		return nil, false
	}
	return o.IsVerified, true
}

// HasIsVerified returns a boolean if a field has been set.
func (o *MicrosoftGraphDomain) HasIsVerified() bool {
	if o != nil && o.IsVerified != nil {
		return true
	}

	return false
}

// SetIsVerified gets a reference to the given bool and assigns it to the IsVerified field.
func (o *MicrosoftGraphDomain) SetIsVerified(v bool) {
	o.IsVerified = &v
}

// GetManufacturer returns the Manufacturer field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphDomain) GetManufacturer() string {
	if o == nil || o.Manufacturer.Get() == nil {
		var ret string
		return ret
	}
	return *o.Manufacturer.Get()
}

// GetManufacturerOk returns a tuple with the Manufacturer field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphDomain) GetManufacturerOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Manufacturer.Get(), o.Manufacturer.IsSet()
}

// HasManufacturer returns a boolean if a field has been set.
func (o *MicrosoftGraphDomain) HasManufacturer() bool {
	if o != nil && o.Manufacturer.IsSet() {
		return true
	}

	return false
}

// SetManufacturer gets a reference to the given NullableString and assigns it to the Manufacturer field.
func (o *MicrosoftGraphDomain) SetManufacturer(v string) {
	o.Manufacturer.Set(&v)
}
// SetManufacturerNil sets the value for Manufacturer to be an explicit nil
func (o *MicrosoftGraphDomain) SetManufacturerNil() {
	o.Manufacturer.Set(nil)
}

// UnsetManufacturer ensures that no value is present for Manufacturer, not even an explicit nil
func (o *MicrosoftGraphDomain) UnsetManufacturer() {
	o.Manufacturer.Unset()
}

// GetModel returns the Model field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphDomain) GetModel() string {
	if o == nil || o.Model.Get() == nil {
		var ret string
		return ret
	}
	return *o.Model.Get()
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphDomain) GetModelOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Model.Get(), o.Model.IsSet()
}

// HasModel returns a boolean if a field has been set.
func (o *MicrosoftGraphDomain) HasModel() bool {
	if o != nil && o.Model.IsSet() {
		return true
	}

	return false
}

// SetModel gets a reference to the given NullableString and assigns it to the Model field.
func (o *MicrosoftGraphDomain) SetModel(v string) {
	o.Model.Set(&v)
}
// SetModelNil sets the value for Model to be an explicit nil
func (o *MicrosoftGraphDomain) SetModelNil() {
	o.Model.Set(nil)
}

// UnsetModel ensures that no value is present for Model, not even an explicit nil
func (o *MicrosoftGraphDomain) UnsetModel() {
	o.Model.Unset()
}

// GetPasswordNotificationWindowInDays returns the PasswordNotificationWindowInDays field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphDomain) GetPasswordNotificationWindowInDays() int32 {
	if o == nil || o.PasswordNotificationWindowInDays.Get() == nil {
		var ret int32
		return ret
	}
	return *o.PasswordNotificationWindowInDays.Get()
}

// GetPasswordNotificationWindowInDaysOk returns a tuple with the PasswordNotificationWindowInDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphDomain) GetPasswordNotificationWindowInDaysOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PasswordNotificationWindowInDays.Get(), o.PasswordNotificationWindowInDays.IsSet()
}

// HasPasswordNotificationWindowInDays returns a boolean if a field has been set.
func (o *MicrosoftGraphDomain) HasPasswordNotificationWindowInDays() bool {
	if o != nil && o.PasswordNotificationWindowInDays.IsSet() {
		return true
	}

	return false
}

// SetPasswordNotificationWindowInDays gets a reference to the given NullableInt32 and assigns it to the PasswordNotificationWindowInDays field.
func (o *MicrosoftGraphDomain) SetPasswordNotificationWindowInDays(v int32) {
	o.PasswordNotificationWindowInDays.Set(&v)
}
// SetPasswordNotificationWindowInDaysNil sets the value for PasswordNotificationWindowInDays to be an explicit nil
func (o *MicrosoftGraphDomain) SetPasswordNotificationWindowInDaysNil() {
	o.PasswordNotificationWindowInDays.Set(nil)
}

// UnsetPasswordNotificationWindowInDays ensures that no value is present for PasswordNotificationWindowInDays, not even an explicit nil
func (o *MicrosoftGraphDomain) UnsetPasswordNotificationWindowInDays() {
	o.PasswordNotificationWindowInDays.Unset()
}

// GetPasswordValidityPeriodInDays returns the PasswordValidityPeriodInDays field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphDomain) GetPasswordValidityPeriodInDays() int32 {
	if o == nil || o.PasswordValidityPeriodInDays.Get() == nil {
		var ret int32
		return ret
	}
	return *o.PasswordValidityPeriodInDays.Get()
}

// GetPasswordValidityPeriodInDaysOk returns a tuple with the PasswordValidityPeriodInDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphDomain) GetPasswordValidityPeriodInDaysOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PasswordValidityPeriodInDays.Get(), o.PasswordValidityPeriodInDays.IsSet()
}

// HasPasswordValidityPeriodInDays returns a boolean if a field has been set.
func (o *MicrosoftGraphDomain) HasPasswordValidityPeriodInDays() bool {
	if o != nil && o.PasswordValidityPeriodInDays.IsSet() {
		return true
	}

	return false
}

// SetPasswordValidityPeriodInDays gets a reference to the given NullableInt32 and assigns it to the PasswordValidityPeriodInDays field.
func (o *MicrosoftGraphDomain) SetPasswordValidityPeriodInDays(v int32) {
	o.PasswordValidityPeriodInDays.Set(&v)
}
// SetPasswordValidityPeriodInDaysNil sets the value for PasswordValidityPeriodInDays to be an explicit nil
func (o *MicrosoftGraphDomain) SetPasswordValidityPeriodInDaysNil() {
	o.PasswordValidityPeriodInDays.Set(nil)
}

// UnsetPasswordValidityPeriodInDays ensures that no value is present for PasswordValidityPeriodInDays, not even an explicit nil
func (o *MicrosoftGraphDomain) UnsetPasswordValidityPeriodInDays() {
	o.PasswordValidityPeriodInDays.Unset()
}

// GetState returns the State field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphDomain) GetState() AnyOfmicrosoftGraphDomainState {
	if o == nil  {
		var ret AnyOfmicrosoftGraphDomainState
		return ret
	}
	return o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphDomain) GetStateOk() (*AnyOfmicrosoftGraphDomainState, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return &o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *MicrosoftGraphDomain) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given AnyOfmicrosoftGraphDomainState and assigns it to the State field.
func (o *MicrosoftGraphDomain) SetState(v AnyOfmicrosoftGraphDomainState) {
	o.State = v
}

// GetSupportedServices returns the SupportedServices field value if set, zero value otherwise.
func (o *MicrosoftGraphDomain) GetSupportedServices() []string {
	if o == nil || o.SupportedServices == nil {
		var ret []string
		return ret
	}
	return *o.SupportedServices
}

// GetSupportedServicesOk returns a tuple with the SupportedServices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphDomain) GetSupportedServicesOk() (*[]string, bool) {
	if o == nil || o.SupportedServices == nil {
		return nil, false
	}
	return o.SupportedServices, true
}

// HasSupportedServices returns a boolean if a field has been set.
func (o *MicrosoftGraphDomain) HasSupportedServices() bool {
	if o != nil && o.SupportedServices != nil {
		return true
	}

	return false
}

// SetSupportedServices gets a reference to the given []string and assigns it to the SupportedServices field.
func (o *MicrosoftGraphDomain) SetSupportedServices(v []string) {
	o.SupportedServices = &v
}

// GetDomainNameReferences returns the DomainNameReferences field value if set, zero value otherwise.
func (o *MicrosoftGraphDomain) GetDomainNameReferences() []MicrosoftGraphDirectoryObject {
	if o == nil || o.DomainNameReferences == nil {
		var ret []MicrosoftGraphDirectoryObject
		return ret
	}
	return *o.DomainNameReferences
}

// GetDomainNameReferencesOk returns a tuple with the DomainNameReferences field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphDomain) GetDomainNameReferencesOk() (*[]MicrosoftGraphDirectoryObject, bool) {
	if o == nil || o.DomainNameReferences == nil {
		return nil, false
	}
	return o.DomainNameReferences, true
}

// HasDomainNameReferences returns a boolean if a field has been set.
func (o *MicrosoftGraphDomain) HasDomainNameReferences() bool {
	if o != nil && o.DomainNameReferences != nil {
		return true
	}

	return false
}

// SetDomainNameReferences gets a reference to the given []MicrosoftGraphDirectoryObject and assigns it to the DomainNameReferences field.
func (o *MicrosoftGraphDomain) SetDomainNameReferences(v []MicrosoftGraphDirectoryObject) {
	o.DomainNameReferences = &v
}

// GetServiceConfigurationRecords returns the ServiceConfigurationRecords field value if set, zero value otherwise.
func (o *MicrosoftGraphDomain) GetServiceConfigurationRecords() []MicrosoftGraphDomainDnsRecord {
	if o == nil || o.ServiceConfigurationRecords == nil {
		var ret []MicrosoftGraphDomainDnsRecord
		return ret
	}
	return *o.ServiceConfigurationRecords
}

// GetServiceConfigurationRecordsOk returns a tuple with the ServiceConfigurationRecords field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphDomain) GetServiceConfigurationRecordsOk() (*[]MicrosoftGraphDomainDnsRecord, bool) {
	if o == nil || o.ServiceConfigurationRecords == nil {
		return nil, false
	}
	return o.ServiceConfigurationRecords, true
}

// HasServiceConfigurationRecords returns a boolean if a field has been set.
func (o *MicrosoftGraphDomain) HasServiceConfigurationRecords() bool {
	if o != nil && o.ServiceConfigurationRecords != nil {
		return true
	}

	return false
}

// SetServiceConfigurationRecords gets a reference to the given []MicrosoftGraphDomainDnsRecord and assigns it to the ServiceConfigurationRecords field.
func (o *MicrosoftGraphDomain) SetServiceConfigurationRecords(v []MicrosoftGraphDomainDnsRecord) {
	o.ServiceConfigurationRecords = &v
}

// GetVerificationDnsRecords returns the VerificationDnsRecords field value if set, zero value otherwise.
func (o *MicrosoftGraphDomain) GetVerificationDnsRecords() []MicrosoftGraphDomainDnsRecord {
	if o == nil || o.VerificationDnsRecords == nil {
		var ret []MicrosoftGraphDomainDnsRecord
		return ret
	}
	return *o.VerificationDnsRecords
}

// GetVerificationDnsRecordsOk returns a tuple with the VerificationDnsRecords field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphDomain) GetVerificationDnsRecordsOk() (*[]MicrosoftGraphDomainDnsRecord, bool) {
	if o == nil || o.VerificationDnsRecords == nil {
		return nil, false
	}
	return o.VerificationDnsRecords, true
}

// HasVerificationDnsRecords returns a boolean if a field has been set.
func (o *MicrosoftGraphDomain) HasVerificationDnsRecords() bool {
	if o != nil && o.VerificationDnsRecords != nil {
		return true
	}

	return false
}

// SetVerificationDnsRecords gets a reference to the given []MicrosoftGraphDomainDnsRecord and assigns it to the VerificationDnsRecords field.
func (o *MicrosoftGraphDomain) SetVerificationDnsRecords(v []MicrosoftGraphDomainDnsRecord) {
	o.VerificationDnsRecords = &v
}

func (o MicrosoftGraphDomain) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.AuthenticationType != nil {
		toSerialize["authenticationType"] = o.AuthenticationType
	}
	if o.AvailabilityStatus.IsSet() {
		toSerialize["availabilityStatus"] = o.AvailabilityStatus.Get()
	}
	if o.IsAdminManaged != nil {
		toSerialize["isAdminManaged"] = o.IsAdminManaged
	}
	if o.IsDefault != nil {
		toSerialize["isDefault"] = o.IsDefault
	}
	if o.IsInitial != nil {
		toSerialize["isInitial"] = o.IsInitial
	}
	if o.IsRoot != nil {
		toSerialize["isRoot"] = o.IsRoot
	}
	if o.IsVerified != nil {
		toSerialize["isVerified"] = o.IsVerified
	}
	if o.Manufacturer.IsSet() {
		toSerialize["manufacturer"] = o.Manufacturer.Get()
	}
	if o.Model.IsSet() {
		toSerialize["model"] = o.Model.Get()
	}
	if o.PasswordNotificationWindowInDays.IsSet() {
		toSerialize["passwordNotificationWindowInDays"] = o.PasswordNotificationWindowInDays.Get()
	}
	if o.PasswordValidityPeriodInDays.IsSet() {
		toSerialize["passwordValidityPeriodInDays"] = o.PasswordValidityPeriodInDays.Get()
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.SupportedServices != nil {
		toSerialize["supportedServices"] = o.SupportedServices
	}
	if o.DomainNameReferences != nil {
		toSerialize["domainNameReferences"] = o.DomainNameReferences
	}
	if o.ServiceConfigurationRecords != nil {
		toSerialize["serviceConfigurationRecords"] = o.ServiceConfigurationRecords
	}
	if o.VerificationDnsRecords != nil {
		toSerialize["verificationDnsRecords"] = o.VerificationDnsRecords
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphDomain struct {
	value *MicrosoftGraphDomain
	isSet bool
}

func (v NullableMicrosoftGraphDomain) Get() *MicrosoftGraphDomain {
	return v.value
}

func (v *NullableMicrosoftGraphDomain) Set(val *MicrosoftGraphDomain) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphDomain) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphDomain) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphDomain(val *MicrosoftGraphDomain) *NullableMicrosoftGraphDomain {
	return &NullableMicrosoftGraphDomain{value: val, isSet: true}
}

func (v NullableMicrosoftGraphDomain) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphDomain) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

