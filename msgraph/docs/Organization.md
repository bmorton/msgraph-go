# Organization

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssignedPlans** | Pointer to [**[]MicrosoftGraphAssignedPlan**](MicrosoftGraphAssignedPlan.md) | The collection of service plans associated with the tenant. Not nullable. | [optional] 
**BusinessPhones** | Pointer to **[]string** | Telephone number for the organization. Although this is a string collection, only one number can be set for this property. | [optional] 
**City** | Pointer to **NullableString** | City name of the address for the organization. | [optional] 
**Country** | Pointer to **NullableString** | Country/region name of the address for the organization. | [optional] 
**CountryLetterCode** | Pointer to **NullableString** | Country or region abbreviation for the organization in ISO 3166-2 format. | [optional] 
**CreatedDateTime** | Pointer to **NullableTime** | Timestamp of when the organization was created. The value cannot be modified and is automatically populated when the organization is created. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only. | [optional] 
**DisplayName** | Pointer to **NullableString** | The display name for the tenant. | [optional] 
**MarketingNotificationEmails** | Pointer to **[]string** | Not nullable. | [optional] 
**OnPremisesLastSyncDateTime** | Pointer to **NullableTime** | The time and date at which the tenant was last synced with the on-premises directory. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only. | [optional] 
**OnPremisesSyncEnabled** | Pointer to **NullableBool** | true if this object is synced from an on-premises directory; false if this object was originally synced from an on-premises directory but is no longer synced. Nullable. null if this object has never been synced from an on-premises directory (default). | [optional] 
**PostalCode** | Pointer to **NullableString** | Postal code of the address for the organization. | [optional] 
**PreferredLanguage** | Pointer to **NullableString** | The preferred language for the organization. Should follow ISO 639-1 Code; for example, en. | [optional] 
**PrivacyProfile** | Pointer to [**AnyOfmicrosoftGraphPrivacyProfile**](anyOf&lt;microsoft.graph.privacyProfile&gt;.md) | The privacy profile of an organization. | [optional] 
**ProvisionedPlans** | Pointer to [**[]MicrosoftGraphProvisionedPlan**](MicrosoftGraphProvisionedPlan.md) | Not nullable. | [optional] 
**SecurityComplianceNotificationMails** | Pointer to **[]string** |  | [optional] 
**SecurityComplianceNotificationPhones** | Pointer to **[]string** |  | [optional] 
**State** | Pointer to **NullableString** | State name of the address for the organization. | [optional] 
**Street** | Pointer to **NullableString** | Street name of the address for organization. | [optional] 
**TechnicalNotificationMails** | Pointer to **[]string** | Not nullable. | [optional] 
**TenantType** | Pointer to **NullableString** |  | [optional] 
**VerifiedDomains** | Pointer to [**[]MicrosoftGraphVerifiedDomain**](MicrosoftGraphVerifiedDomain.md) | The collection of domains associated with this tenant. Not nullable. | [optional] 
**MobileDeviceManagementAuthority** | Pointer to [**AnyOfmicrosoftGraphMdmAuthority**](anyOf&lt;microsoft.graph.mdmAuthority&gt;.md) | Mobile device management authority. Possible values are: unknown, intune, sccm, office365. | [optional] 
**Branding** | Pointer to [**AnyOfmicrosoftGraphOrganizationalBranding**](anyOf&lt;microsoft.graph.organizationalBranding&gt;.md) |  | [optional] 
**CertificateBasedAuthConfiguration** | Pointer to [**[]MicrosoftGraphCertificateBasedAuthConfiguration**](MicrosoftGraphCertificateBasedAuthConfiguration.md) | Navigation property to manage certificate-based authentication configuration. Only a single instance of certificateBasedAuthConfiguration can be created in the collection. | [optional] 
**Extensions** | Pointer to [**[]MicrosoftGraphExtension**](MicrosoftGraphExtension.md) | The collection of open extensions defined for the organization. Read-only. Nullable. | [optional] 

## Methods

### NewOrganization

`func NewOrganization() *Organization`

NewOrganization instantiates a new Organization object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationWithDefaults

`func NewOrganizationWithDefaults() *Organization`

NewOrganizationWithDefaults instantiates a new Organization object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssignedPlans

`func (o *Organization) GetAssignedPlans() []MicrosoftGraphAssignedPlan`

GetAssignedPlans returns the AssignedPlans field if non-nil, zero value otherwise.

### GetAssignedPlansOk

`func (o *Organization) GetAssignedPlansOk() (*[]MicrosoftGraphAssignedPlan, bool)`

GetAssignedPlansOk returns a tuple with the AssignedPlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedPlans

`func (o *Organization) SetAssignedPlans(v []MicrosoftGraphAssignedPlan)`

SetAssignedPlans sets AssignedPlans field to given value.

### HasAssignedPlans

`func (o *Organization) HasAssignedPlans() bool`

HasAssignedPlans returns a boolean if a field has been set.

### GetBusinessPhones

`func (o *Organization) GetBusinessPhones() []string`

GetBusinessPhones returns the BusinessPhones field if non-nil, zero value otherwise.

### GetBusinessPhonesOk

`func (o *Organization) GetBusinessPhonesOk() (*[]string, bool)`

GetBusinessPhonesOk returns a tuple with the BusinessPhones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessPhones

`func (o *Organization) SetBusinessPhones(v []string)`

SetBusinessPhones sets BusinessPhones field to given value.

### HasBusinessPhones

`func (o *Organization) HasBusinessPhones() bool`

HasBusinessPhones returns a boolean if a field has been set.

### GetCity

`func (o *Organization) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *Organization) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *Organization) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *Organization) HasCity() bool`

HasCity returns a boolean if a field has been set.

### SetCityNil

`func (o *Organization) SetCityNil(b bool)`

 SetCityNil sets the value for City to be an explicit nil

### UnsetCity
`func (o *Organization) UnsetCity()`

UnsetCity ensures that no value is present for City, not even an explicit nil
### GetCountry

`func (o *Organization) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *Organization) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *Organization) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *Organization) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### SetCountryNil

`func (o *Organization) SetCountryNil(b bool)`

 SetCountryNil sets the value for Country to be an explicit nil

### UnsetCountry
`func (o *Organization) UnsetCountry()`

UnsetCountry ensures that no value is present for Country, not even an explicit nil
### GetCountryLetterCode

`func (o *Organization) GetCountryLetterCode() string`

GetCountryLetterCode returns the CountryLetterCode field if non-nil, zero value otherwise.

### GetCountryLetterCodeOk

`func (o *Organization) GetCountryLetterCodeOk() (*string, bool)`

GetCountryLetterCodeOk returns a tuple with the CountryLetterCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryLetterCode

`func (o *Organization) SetCountryLetterCode(v string)`

SetCountryLetterCode sets CountryLetterCode field to given value.

### HasCountryLetterCode

`func (o *Organization) HasCountryLetterCode() bool`

HasCountryLetterCode returns a boolean if a field has been set.

### SetCountryLetterCodeNil

`func (o *Organization) SetCountryLetterCodeNil(b bool)`

 SetCountryLetterCodeNil sets the value for CountryLetterCode to be an explicit nil

### UnsetCountryLetterCode
`func (o *Organization) UnsetCountryLetterCode()`

UnsetCountryLetterCode ensures that no value is present for CountryLetterCode, not even an explicit nil
### GetCreatedDateTime

`func (o *Organization) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *Organization) GetCreatedDateTimeOk() (*time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateTime

`func (o *Organization) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime sets CreatedDateTime field to given value.

### HasCreatedDateTime

`func (o *Organization) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTimeNil

`func (o *Organization) SetCreatedDateTimeNil(b bool)`

 SetCreatedDateTimeNil sets the value for CreatedDateTime to be an explicit nil

### UnsetCreatedDateTime
`func (o *Organization) UnsetCreatedDateTime()`

UnsetCreatedDateTime ensures that no value is present for CreatedDateTime, not even an explicit nil
### GetDisplayName

`func (o *Organization) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *Organization) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *Organization) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *Organization) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *Organization) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *Organization) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetMarketingNotificationEmails

`func (o *Organization) GetMarketingNotificationEmails() []string`

GetMarketingNotificationEmails returns the MarketingNotificationEmails field if non-nil, zero value otherwise.

### GetMarketingNotificationEmailsOk

`func (o *Organization) GetMarketingNotificationEmailsOk() (*[]string, bool)`

GetMarketingNotificationEmailsOk returns a tuple with the MarketingNotificationEmails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketingNotificationEmails

`func (o *Organization) SetMarketingNotificationEmails(v []string)`

SetMarketingNotificationEmails sets MarketingNotificationEmails field to given value.

### HasMarketingNotificationEmails

`func (o *Organization) HasMarketingNotificationEmails() bool`

HasMarketingNotificationEmails returns a boolean if a field has been set.

### GetOnPremisesLastSyncDateTime

`func (o *Organization) GetOnPremisesLastSyncDateTime() time.Time`

GetOnPremisesLastSyncDateTime returns the OnPremisesLastSyncDateTime field if non-nil, zero value otherwise.

### GetOnPremisesLastSyncDateTimeOk

`func (o *Organization) GetOnPremisesLastSyncDateTimeOk() (*time.Time, bool)`

GetOnPremisesLastSyncDateTimeOk returns a tuple with the OnPremisesLastSyncDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnPremisesLastSyncDateTime

`func (o *Organization) SetOnPremisesLastSyncDateTime(v time.Time)`

SetOnPremisesLastSyncDateTime sets OnPremisesLastSyncDateTime field to given value.

### HasOnPremisesLastSyncDateTime

`func (o *Organization) HasOnPremisesLastSyncDateTime() bool`

HasOnPremisesLastSyncDateTime returns a boolean if a field has been set.

### SetOnPremisesLastSyncDateTimeNil

`func (o *Organization) SetOnPremisesLastSyncDateTimeNil(b bool)`

 SetOnPremisesLastSyncDateTimeNil sets the value for OnPremisesLastSyncDateTime to be an explicit nil

### UnsetOnPremisesLastSyncDateTime
`func (o *Organization) UnsetOnPremisesLastSyncDateTime()`

UnsetOnPremisesLastSyncDateTime ensures that no value is present for OnPremisesLastSyncDateTime, not even an explicit nil
### GetOnPremisesSyncEnabled

`func (o *Organization) GetOnPremisesSyncEnabled() bool`

GetOnPremisesSyncEnabled returns the OnPremisesSyncEnabled field if non-nil, zero value otherwise.

### GetOnPremisesSyncEnabledOk

`func (o *Organization) GetOnPremisesSyncEnabledOk() (*bool, bool)`

GetOnPremisesSyncEnabledOk returns a tuple with the OnPremisesSyncEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnPremisesSyncEnabled

`func (o *Organization) SetOnPremisesSyncEnabled(v bool)`

SetOnPremisesSyncEnabled sets OnPremisesSyncEnabled field to given value.

### HasOnPremisesSyncEnabled

`func (o *Organization) HasOnPremisesSyncEnabled() bool`

HasOnPremisesSyncEnabled returns a boolean if a field has been set.

### SetOnPremisesSyncEnabledNil

`func (o *Organization) SetOnPremisesSyncEnabledNil(b bool)`

 SetOnPremisesSyncEnabledNil sets the value for OnPremisesSyncEnabled to be an explicit nil

### UnsetOnPremisesSyncEnabled
`func (o *Organization) UnsetOnPremisesSyncEnabled()`

UnsetOnPremisesSyncEnabled ensures that no value is present for OnPremisesSyncEnabled, not even an explicit nil
### GetPostalCode

`func (o *Organization) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *Organization) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *Organization) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *Organization) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### SetPostalCodeNil

`func (o *Organization) SetPostalCodeNil(b bool)`

 SetPostalCodeNil sets the value for PostalCode to be an explicit nil

### UnsetPostalCode
`func (o *Organization) UnsetPostalCode()`

UnsetPostalCode ensures that no value is present for PostalCode, not even an explicit nil
### GetPreferredLanguage

`func (o *Organization) GetPreferredLanguage() string`

GetPreferredLanguage returns the PreferredLanguage field if non-nil, zero value otherwise.

### GetPreferredLanguageOk

`func (o *Organization) GetPreferredLanguageOk() (*string, bool)`

GetPreferredLanguageOk returns a tuple with the PreferredLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredLanguage

`func (o *Organization) SetPreferredLanguage(v string)`

SetPreferredLanguage sets PreferredLanguage field to given value.

### HasPreferredLanguage

`func (o *Organization) HasPreferredLanguage() bool`

HasPreferredLanguage returns a boolean if a field has been set.

### SetPreferredLanguageNil

`func (o *Organization) SetPreferredLanguageNil(b bool)`

 SetPreferredLanguageNil sets the value for PreferredLanguage to be an explicit nil

### UnsetPreferredLanguage
`func (o *Organization) UnsetPreferredLanguage()`

UnsetPreferredLanguage ensures that no value is present for PreferredLanguage, not even an explicit nil
### GetPrivacyProfile

`func (o *Organization) GetPrivacyProfile() AnyOfmicrosoftGraphPrivacyProfile`

GetPrivacyProfile returns the PrivacyProfile field if non-nil, zero value otherwise.

### GetPrivacyProfileOk

`func (o *Organization) GetPrivacyProfileOk() (*AnyOfmicrosoftGraphPrivacyProfile, bool)`

GetPrivacyProfileOk returns a tuple with the PrivacyProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacyProfile

`func (o *Organization) SetPrivacyProfile(v AnyOfmicrosoftGraphPrivacyProfile)`

SetPrivacyProfile sets PrivacyProfile field to given value.

### HasPrivacyProfile

`func (o *Organization) HasPrivacyProfile() bool`

HasPrivacyProfile returns a boolean if a field has been set.

### SetPrivacyProfileNil

`func (o *Organization) SetPrivacyProfileNil(b bool)`

 SetPrivacyProfileNil sets the value for PrivacyProfile to be an explicit nil

### UnsetPrivacyProfile
`func (o *Organization) UnsetPrivacyProfile()`

UnsetPrivacyProfile ensures that no value is present for PrivacyProfile, not even an explicit nil
### GetProvisionedPlans

`func (o *Organization) GetProvisionedPlans() []MicrosoftGraphProvisionedPlan`

GetProvisionedPlans returns the ProvisionedPlans field if non-nil, zero value otherwise.

### GetProvisionedPlansOk

`func (o *Organization) GetProvisionedPlansOk() (*[]MicrosoftGraphProvisionedPlan, bool)`

GetProvisionedPlansOk returns a tuple with the ProvisionedPlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionedPlans

`func (o *Organization) SetProvisionedPlans(v []MicrosoftGraphProvisionedPlan)`

SetProvisionedPlans sets ProvisionedPlans field to given value.

### HasProvisionedPlans

`func (o *Organization) HasProvisionedPlans() bool`

HasProvisionedPlans returns a boolean if a field has been set.

### GetSecurityComplianceNotificationMails

`func (o *Organization) GetSecurityComplianceNotificationMails() []string`

GetSecurityComplianceNotificationMails returns the SecurityComplianceNotificationMails field if non-nil, zero value otherwise.

### GetSecurityComplianceNotificationMailsOk

`func (o *Organization) GetSecurityComplianceNotificationMailsOk() (*[]string, bool)`

GetSecurityComplianceNotificationMailsOk returns a tuple with the SecurityComplianceNotificationMails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityComplianceNotificationMails

`func (o *Organization) SetSecurityComplianceNotificationMails(v []string)`

SetSecurityComplianceNotificationMails sets SecurityComplianceNotificationMails field to given value.

### HasSecurityComplianceNotificationMails

`func (o *Organization) HasSecurityComplianceNotificationMails() bool`

HasSecurityComplianceNotificationMails returns a boolean if a field has been set.

### GetSecurityComplianceNotificationPhones

`func (o *Organization) GetSecurityComplianceNotificationPhones() []string`

GetSecurityComplianceNotificationPhones returns the SecurityComplianceNotificationPhones field if non-nil, zero value otherwise.

### GetSecurityComplianceNotificationPhonesOk

`func (o *Organization) GetSecurityComplianceNotificationPhonesOk() (*[]string, bool)`

GetSecurityComplianceNotificationPhonesOk returns a tuple with the SecurityComplianceNotificationPhones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityComplianceNotificationPhones

`func (o *Organization) SetSecurityComplianceNotificationPhones(v []string)`

SetSecurityComplianceNotificationPhones sets SecurityComplianceNotificationPhones field to given value.

### HasSecurityComplianceNotificationPhones

`func (o *Organization) HasSecurityComplianceNotificationPhones() bool`

HasSecurityComplianceNotificationPhones returns a boolean if a field has been set.

### GetState

`func (o *Organization) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Organization) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Organization) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *Organization) HasState() bool`

HasState returns a boolean if a field has been set.

### SetStateNil

`func (o *Organization) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *Organization) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil
### GetStreet

`func (o *Organization) GetStreet() string`

GetStreet returns the Street field if non-nil, zero value otherwise.

### GetStreetOk

`func (o *Organization) GetStreetOk() (*string, bool)`

GetStreetOk returns a tuple with the Street field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreet

`func (o *Organization) SetStreet(v string)`

SetStreet sets Street field to given value.

### HasStreet

`func (o *Organization) HasStreet() bool`

HasStreet returns a boolean if a field has been set.

### SetStreetNil

`func (o *Organization) SetStreetNil(b bool)`

 SetStreetNil sets the value for Street to be an explicit nil

### UnsetStreet
`func (o *Organization) UnsetStreet()`

UnsetStreet ensures that no value is present for Street, not even an explicit nil
### GetTechnicalNotificationMails

`func (o *Organization) GetTechnicalNotificationMails() []string`

GetTechnicalNotificationMails returns the TechnicalNotificationMails field if non-nil, zero value otherwise.

### GetTechnicalNotificationMailsOk

`func (o *Organization) GetTechnicalNotificationMailsOk() (*[]string, bool)`

GetTechnicalNotificationMailsOk returns a tuple with the TechnicalNotificationMails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTechnicalNotificationMails

`func (o *Organization) SetTechnicalNotificationMails(v []string)`

SetTechnicalNotificationMails sets TechnicalNotificationMails field to given value.

### HasTechnicalNotificationMails

`func (o *Organization) HasTechnicalNotificationMails() bool`

HasTechnicalNotificationMails returns a boolean if a field has been set.

### GetTenantType

`func (o *Organization) GetTenantType() string`

GetTenantType returns the TenantType field if non-nil, zero value otherwise.

### GetTenantTypeOk

`func (o *Organization) GetTenantTypeOk() (*string, bool)`

GetTenantTypeOk returns a tuple with the TenantType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantType

`func (o *Organization) SetTenantType(v string)`

SetTenantType sets TenantType field to given value.

### HasTenantType

`func (o *Organization) HasTenantType() bool`

HasTenantType returns a boolean if a field has been set.

### SetTenantTypeNil

`func (o *Organization) SetTenantTypeNil(b bool)`

 SetTenantTypeNil sets the value for TenantType to be an explicit nil

### UnsetTenantType
`func (o *Organization) UnsetTenantType()`

UnsetTenantType ensures that no value is present for TenantType, not even an explicit nil
### GetVerifiedDomains

`func (o *Organization) GetVerifiedDomains() []MicrosoftGraphVerifiedDomain`

GetVerifiedDomains returns the VerifiedDomains field if non-nil, zero value otherwise.

### GetVerifiedDomainsOk

`func (o *Organization) GetVerifiedDomainsOk() (*[]MicrosoftGraphVerifiedDomain, bool)`

GetVerifiedDomainsOk returns a tuple with the VerifiedDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifiedDomains

`func (o *Organization) SetVerifiedDomains(v []MicrosoftGraphVerifiedDomain)`

SetVerifiedDomains sets VerifiedDomains field to given value.

### HasVerifiedDomains

`func (o *Organization) HasVerifiedDomains() bool`

HasVerifiedDomains returns a boolean if a field has been set.

### GetMobileDeviceManagementAuthority

`func (o *Organization) GetMobileDeviceManagementAuthority() AnyOfmicrosoftGraphMdmAuthority`

GetMobileDeviceManagementAuthority returns the MobileDeviceManagementAuthority field if non-nil, zero value otherwise.

### GetMobileDeviceManagementAuthorityOk

`func (o *Organization) GetMobileDeviceManagementAuthorityOk() (*AnyOfmicrosoftGraphMdmAuthority, bool)`

GetMobileDeviceManagementAuthorityOk returns a tuple with the MobileDeviceManagementAuthority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobileDeviceManagementAuthority

`func (o *Organization) SetMobileDeviceManagementAuthority(v AnyOfmicrosoftGraphMdmAuthority)`

SetMobileDeviceManagementAuthority sets MobileDeviceManagementAuthority field to given value.

### HasMobileDeviceManagementAuthority

`func (o *Organization) HasMobileDeviceManagementAuthority() bool`

HasMobileDeviceManagementAuthority returns a boolean if a field has been set.

### SetMobileDeviceManagementAuthorityNil

`func (o *Organization) SetMobileDeviceManagementAuthorityNil(b bool)`

 SetMobileDeviceManagementAuthorityNil sets the value for MobileDeviceManagementAuthority to be an explicit nil

### UnsetMobileDeviceManagementAuthority
`func (o *Organization) UnsetMobileDeviceManagementAuthority()`

UnsetMobileDeviceManagementAuthority ensures that no value is present for MobileDeviceManagementAuthority, not even an explicit nil
### GetBranding

`func (o *Organization) GetBranding() AnyOfmicrosoftGraphOrganizationalBranding`

GetBranding returns the Branding field if non-nil, zero value otherwise.

### GetBrandingOk

`func (o *Organization) GetBrandingOk() (*AnyOfmicrosoftGraphOrganizationalBranding, bool)`

GetBrandingOk returns a tuple with the Branding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranding

`func (o *Organization) SetBranding(v AnyOfmicrosoftGraphOrganizationalBranding)`

SetBranding sets Branding field to given value.

### HasBranding

`func (o *Organization) HasBranding() bool`

HasBranding returns a boolean if a field has been set.

### SetBrandingNil

`func (o *Organization) SetBrandingNil(b bool)`

 SetBrandingNil sets the value for Branding to be an explicit nil

### UnsetBranding
`func (o *Organization) UnsetBranding()`

UnsetBranding ensures that no value is present for Branding, not even an explicit nil
### GetCertificateBasedAuthConfiguration

`func (o *Organization) GetCertificateBasedAuthConfiguration() []MicrosoftGraphCertificateBasedAuthConfiguration`

GetCertificateBasedAuthConfiguration returns the CertificateBasedAuthConfiguration field if non-nil, zero value otherwise.

### GetCertificateBasedAuthConfigurationOk

`func (o *Organization) GetCertificateBasedAuthConfigurationOk() (*[]MicrosoftGraphCertificateBasedAuthConfiguration, bool)`

GetCertificateBasedAuthConfigurationOk returns a tuple with the CertificateBasedAuthConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateBasedAuthConfiguration

`func (o *Organization) SetCertificateBasedAuthConfiguration(v []MicrosoftGraphCertificateBasedAuthConfiguration)`

SetCertificateBasedAuthConfiguration sets CertificateBasedAuthConfiguration field to given value.

### HasCertificateBasedAuthConfiguration

`func (o *Organization) HasCertificateBasedAuthConfiguration() bool`

HasCertificateBasedAuthConfiguration returns a boolean if a field has been set.

### GetExtensions

`func (o *Organization) GetExtensions() []MicrosoftGraphExtension`

GetExtensions returns the Extensions field if non-nil, zero value otherwise.

### GetExtensionsOk

`func (o *Organization) GetExtensionsOk() (*[]MicrosoftGraphExtension, bool)`

GetExtensionsOk returns a tuple with the Extensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensions

`func (o *Organization) SetExtensions(v []MicrosoftGraphExtension)`

SetExtensions sets Extensions field to given value.

### HasExtensions

`func (o *Organization) HasExtensions() bool`

HasExtensions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


