# MicrosoftGraphOrganization

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**DeletedDateTime** | Pointer to **NullableTime** |  | [optional] 
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

### NewMicrosoftGraphOrganization

`func NewMicrosoftGraphOrganization() *MicrosoftGraphOrganization`

NewMicrosoftGraphOrganization instantiates a new MicrosoftGraphOrganization object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphOrganizationWithDefaults

`func NewMicrosoftGraphOrganizationWithDefaults() *MicrosoftGraphOrganization`

NewMicrosoftGraphOrganizationWithDefaults instantiates a new MicrosoftGraphOrganization object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphOrganization) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphOrganization) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphOrganization) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphOrganization) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDeletedDateTime

`func (o *MicrosoftGraphOrganization) GetDeletedDateTime() time.Time`

GetDeletedDateTime returns the DeletedDateTime field if non-nil, zero value otherwise.

### GetDeletedDateTimeOk

`func (o *MicrosoftGraphOrganization) GetDeletedDateTimeOk() (*time.Time, bool)`

GetDeletedDateTimeOk returns a tuple with the DeletedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedDateTime

`func (o *MicrosoftGraphOrganization) SetDeletedDateTime(v time.Time)`

SetDeletedDateTime sets DeletedDateTime field to given value.

### HasDeletedDateTime

`func (o *MicrosoftGraphOrganization) HasDeletedDateTime() bool`

HasDeletedDateTime returns a boolean if a field has been set.

### SetDeletedDateTimeNil

`func (o *MicrosoftGraphOrganization) SetDeletedDateTimeNil(b bool)`

 SetDeletedDateTimeNil sets the value for DeletedDateTime to be an explicit nil

### UnsetDeletedDateTime
`func (o *MicrosoftGraphOrganization) UnsetDeletedDateTime()`

UnsetDeletedDateTime ensures that no value is present for DeletedDateTime, not even an explicit nil
### GetAssignedPlans

`func (o *MicrosoftGraphOrganization) GetAssignedPlans() []MicrosoftGraphAssignedPlan`

GetAssignedPlans returns the AssignedPlans field if non-nil, zero value otherwise.

### GetAssignedPlansOk

`func (o *MicrosoftGraphOrganization) GetAssignedPlansOk() (*[]MicrosoftGraphAssignedPlan, bool)`

GetAssignedPlansOk returns a tuple with the AssignedPlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedPlans

`func (o *MicrosoftGraphOrganization) SetAssignedPlans(v []MicrosoftGraphAssignedPlan)`

SetAssignedPlans sets AssignedPlans field to given value.

### HasAssignedPlans

`func (o *MicrosoftGraphOrganization) HasAssignedPlans() bool`

HasAssignedPlans returns a boolean if a field has been set.

### GetBusinessPhones

`func (o *MicrosoftGraphOrganization) GetBusinessPhones() []string`

GetBusinessPhones returns the BusinessPhones field if non-nil, zero value otherwise.

### GetBusinessPhonesOk

`func (o *MicrosoftGraphOrganization) GetBusinessPhonesOk() (*[]string, bool)`

GetBusinessPhonesOk returns a tuple with the BusinessPhones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessPhones

`func (o *MicrosoftGraphOrganization) SetBusinessPhones(v []string)`

SetBusinessPhones sets BusinessPhones field to given value.

### HasBusinessPhones

`func (o *MicrosoftGraphOrganization) HasBusinessPhones() bool`

HasBusinessPhones returns a boolean if a field has been set.

### GetCity

`func (o *MicrosoftGraphOrganization) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *MicrosoftGraphOrganization) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *MicrosoftGraphOrganization) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *MicrosoftGraphOrganization) HasCity() bool`

HasCity returns a boolean if a field has been set.

### SetCityNil

`func (o *MicrosoftGraphOrganization) SetCityNil(b bool)`

 SetCityNil sets the value for City to be an explicit nil

### UnsetCity
`func (o *MicrosoftGraphOrganization) UnsetCity()`

UnsetCity ensures that no value is present for City, not even an explicit nil
### GetCountry

`func (o *MicrosoftGraphOrganization) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *MicrosoftGraphOrganization) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *MicrosoftGraphOrganization) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *MicrosoftGraphOrganization) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### SetCountryNil

`func (o *MicrosoftGraphOrganization) SetCountryNil(b bool)`

 SetCountryNil sets the value for Country to be an explicit nil

### UnsetCountry
`func (o *MicrosoftGraphOrganization) UnsetCountry()`

UnsetCountry ensures that no value is present for Country, not even an explicit nil
### GetCountryLetterCode

`func (o *MicrosoftGraphOrganization) GetCountryLetterCode() string`

GetCountryLetterCode returns the CountryLetterCode field if non-nil, zero value otherwise.

### GetCountryLetterCodeOk

`func (o *MicrosoftGraphOrganization) GetCountryLetterCodeOk() (*string, bool)`

GetCountryLetterCodeOk returns a tuple with the CountryLetterCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryLetterCode

`func (o *MicrosoftGraphOrganization) SetCountryLetterCode(v string)`

SetCountryLetterCode sets CountryLetterCode field to given value.

### HasCountryLetterCode

`func (o *MicrosoftGraphOrganization) HasCountryLetterCode() bool`

HasCountryLetterCode returns a boolean if a field has been set.

### SetCountryLetterCodeNil

`func (o *MicrosoftGraphOrganization) SetCountryLetterCodeNil(b bool)`

 SetCountryLetterCodeNil sets the value for CountryLetterCode to be an explicit nil

### UnsetCountryLetterCode
`func (o *MicrosoftGraphOrganization) UnsetCountryLetterCode()`

UnsetCountryLetterCode ensures that no value is present for CountryLetterCode, not even an explicit nil
### GetCreatedDateTime

`func (o *MicrosoftGraphOrganization) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *MicrosoftGraphOrganization) GetCreatedDateTimeOk() (*time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateTime

`func (o *MicrosoftGraphOrganization) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime sets CreatedDateTime field to given value.

### HasCreatedDateTime

`func (o *MicrosoftGraphOrganization) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTimeNil

`func (o *MicrosoftGraphOrganization) SetCreatedDateTimeNil(b bool)`

 SetCreatedDateTimeNil sets the value for CreatedDateTime to be an explicit nil

### UnsetCreatedDateTime
`func (o *MicrosoftGraphOrganization) UnsetCreatedDateTime()`

UnsetCreatedDateTime ensures that no value is present for CreatedDateTime, not even an explicit nil
### GetDisplayName

`func (o *MicrosoftGraphOrganization) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphOrganization) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *MicrosoftGraphOrganization) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *MicrosoftGraphOrganization) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *MicrosoftGraphOrganization) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *MicrosoftGraphOrganization) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetMarketingNotificationEmails

`func (o *MicrosoftGraphOrganization) GetMarketingNotificationEmails() []string`

GetMarketingNotificationEmails returns the MarketingNotificationEmails field if non-nil, zero value otherwise.

### GetMarketingNotificationEmailsOk

`func (o *MicrosoftGraphOrganization) GetMarketingNotificationEmailsOk() (*[]string, bool)`

GetMarketingNotificationEmailsOk returns a tuple with the MarketingNotificationEmails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketingNotificationEmails

`func (o *MicrosoftGraphOrganization) SetMarketingNotificationEmails(v []string)`

SetMarketingNotificationEmails sets MarketingNotificationEmails field to given value.

### HasMarketingNotificationEmails

`func (o *MicrosoftGraphOrganization) HasMarketingNotificationEmails() bool`

HasMarketingNotificationEmails returns a boolean if a field has been set.

### GetOnPremisesLastSyncDateTime

`func (o *MicrosoftGraphOrganization) GetOnPremisesLastSyncDateTime() time.Time`

GetOnPremisesLastSyncDateTime returns the OnPremisesLastSyncDateTime field if non-nil, zero value otherwise.

### GetOnPremisesLastSyncDateTimeOk

`func (o *MicrosoftGraphOrganization) GetOnPremisesLastSyncDateTimeOk() (*time.Time, bool)`

GetOnPremisesLastSyncDateTimeOk returns a tuple with the OnPremisesLastSyncDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnPremisesLastSyncDateTime

`func (o *MicrosoftGraphOrganization) SetOnPremisesLastSyncDateTime(v time.Time)`

SetOnPremisesLastSyncDateTime sets OnPremisesLastSyncDateTime field to given value.

### HasOnPremisesLastSyncDateTime

`func (o *MicrosoftGraphOrganization) HasOnPremisesLastSyncDateTime() bool`

HasOnPremisesLastSyncDateTime returns a boolean if a field has been set.

### SetOnPremisesLastSyncDateTimeNil

`func (o *MicrosoftGraphOrganization) SetOnPremisesLastSyncDateTimeNil(b bool)`

 SetOnPremisesLastSyncDateTimeNil sets the value for OnPremisesLastSyncDateTime to be an explicit nil

### UnsetOnPremisesLastSyncDateTime
`func (o *MicrosoftGraphOrganization) UnsetOnPremisesLastSyncDateTime()`

UnsetOnPremisesLastSyncDateTime ensures that no value is present for OnPremisesLastSyncDateTime, not even an explicit nil
### GetOnPremisesSyncEnabled

`func (o *MicrosoftGraphOrganization) GetOnPremisesSyncEnabled() bool`

GetOnPremisesSyncEnabled returns the OnPremisesSyncEnabled field if non-nil, zero value otherwise.

### GetOnPremisesSyncEnabledOk

`func (o *MicrosoftGraphOrganization) GetOnPremisesSyncEnabledOk() (*bool, bool)`

GetOnPremisesSyncEnabledOk returns a tuple with the OnPremisesSyncEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnPremisesSyncEnabled

`func (o *MicrosoftGraphOrganization) SetOnPremisesSyncEnabled(v bool)`

SetOnPremisesSyncEnabled sets OnPremisesSyncEnabled field to given value.

### HasOnPremisesSyncEnabled

`func (o *MicrosoftGraphOrganization) HasOnPremisesSyncEnabled() bool`

HasOnPremisesSyncEnabled returns a boolean if a field has been set.

### SetOnPremisesSyncEnabledNil

`func (o *MicrosoftGraphOrganization) SetOnPremisesSyncEnabledNil(b bool)`

 SetOnPremisesSyncEnabledNil sets the value for OnPremisesSyncEnabled to be an explicit nil

### UnsetOnPremisesSyncEnabled
`func (o *MicrosoftGraphOrganization) UnsetOnPremisesSyncEnabled()`

UnsetOnPremisesSyncEnabled ensures that no value is present for OnPremisesSyncEnabled, not even an explicit nil
### GetPostalCode

`func (o *MicrosoftGraphOrganization) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *MicrosoftGraphOrganization) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *MicrosoftGraphOrganization) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *MicrosoftGraphOrganization) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### SetPostalCodeNil

`func (o *MicrosoftGraphOrganization) SetPostalCodeNil(b bool)`

 SetPostalCodeNil sets the value for PostalCode to be an explicit nil

### UnsetPostalCode
`func (o *MicrosoftGraphOrganization) UnsetPostalCode()`

UnsetPostalCode ensures that no value is present for PostalCode, not even an explicit nil
### GetPreferredLanguage

`func (o *MicrosoftGraphOrganization) GetPreferredLanguage() string`

GetPreferredLanguage returns the PreferredLanguage field if non-nil, zero value otherwise.

### GetPreferredLanguageOk

`func (o *MicrosoftGraphOrganization) GetPreferredLanguageOk() (*string, bool)`

GetPreferredLanguageOk returns a tuple with the PreferredLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredLanguage

`func (o *MicrosoftGraphOrganization) SetPreferredLanguage(v string)`

SetPreferredLanguage sets PreferredLanguage field to given value.

### HasPreferredLanguage

`func (o *MicrosoftGraphOrganization) HasPreferredLanguage() bool`

HasPreferredLanguage returns a boolean if a field has been set.

### SetPreferredLanguageNil

`func (o *MicrosoftGraphOrganization) SetPreferredLanguageNil(b bool)`

 SetPreferredLanguageNil sets the value for PreferredLanguage to be an explicit nil

### UnsetPreferredLanguage
`func (o *MicrosoftGraphOrganization) UnsetPreferredLanguage()`

UnsetPreferredLanguage ensures that no value is present for PreferredLanguage, not even an explicit nil
### GetPrivacyProfile

`func (o *MicrosoftGraphOrganization) GetPrivacyProfile() AnyOfmicrosoftGraphPrivacyProfile`

GetPrivacyProfile returns the PrivacyProfile field if non-nil, zero value otherwise.

### GetPrivacyProfileOk

`func (o *MicrosoftGraphOrganization) GetPrivacyProfileOk() (*AnyOfmicrosoftGraphPrivacyProfile, bool)`

GetPrivacyProfileOk returns a tuple with the PrivacyProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacyProfile

`func (o *MicrosoftGraphOrganization) SetPrivacyProfile(v AnyOfmicrosoftGraphPrivacyProfile)`

SetPrivacyProfile sets PrivacyProfile field to given value.

### HasPrivacyProfile

`func (o *MicrosoftGraphOrganization) HasPrivacyProfile() bool`

HasPrivacyProfile returns a boolean if a field has been set.

### SetPrivacyProfileNil

`func (o *MicrosoftGraphOrganization) SetPrivacyProfileNil(b bool)`

 SetPrivacyProfileNil sets the value for PrivacyProfile to be an explicit nil

### UnsetPrivacyProfile
`func (o *MicrosoftGraphOrganization) UnsetPrivacyProfile()`

UnsetPrivacyProfile ensures that no value is present for PrivacyProfile, not even an explicit nil
### GetProvisionedPlans

`func (o *MicrosoftGraphOrganization) GetProvisionedPlans() []MicrosoftGraphProvisionedPlan`

GetProvisionedPlans returns the ProvisionedPlans field if non-nil, zero value otherwise.

### GetProvisionedPlansOk

`func (o *MicrosoftGraphOrganization) GetProvisionedPlansOk() (*[]MicrosoftGraphProvisionedPlan, bool)`

GetProvisionedPlansOk returns a tuple with the ProvisionedPlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionedPlans

`func (o *MicrosoftGraphOrganization) SetProvisionedPlans(v []MicrosoftGraphProvisionedPlan)`

SetProvisionedPlans sets ProvisionedPlans field to given value.

### HasProvisionedPlans

`func (o *MicrosoftGraphOrganization) HasProvisionedPlans() bool`

HasProvisionedPlans returns a boolean if a field has been set.

### GetSecurityComplianceNotificationMails

`func (o *MicrosoftGraphOrganization) GetSecurityComplianceNotificationMails() []string`

GetSecurityComplianceNotificationMails returns the SecurityComplianceNotificationMails field if non-nil, zero value otherwise.

### GetSecurityComplianceNotificationMailsOk

`func (o *MicrosoftGraphOrganization) GetSecurityComplianceNotificationMailsOk() (*[]string, bool)`

GetSecurityComplianceNotificationMailsOk returns a tuple with the SecurityComplianceNotificationMails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityComplianceNotificationMails

`func (o *MicrosoftGraphOrganization) SetSecurityComplianceNotificationMails(v []string)`

SetSecurityComplianceNotificationMails sets SecurityComplianceNotificationMails field to given value.

### HasSecurityComplianceNotificationMails

`func (o *MicrosoftGraphOrganization) HasSecurityComplianceNotificationMails() bool`

HasSecurityComplianceNotificationMails returns a boolean if a field has been set.

### GetSecurityComplianceNotificationPhones

`func (o *MicrosoftGraphOrganization) GetSecurityComplianceNotificationPhones() []string`

GetSecurityComplianceNotificationPhones returns the SecurityComplianceNotificationPhones field if non-nil, zero value otherwise.

### GetSecurityComplianceNotificationPhonesOk

`func (o *MicrosoftGraphOrganization) GetSecurityComplianceNotificationPhonesOk() (*[]string, bool)`

GetSecurityComplianceNotificationPhonesOk returns a tuple with the SecurityComplianceNotificationPhones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityComplianceNotificationPhones

`func (o *MicrosoftGraphOrganization) SetSecurityComplianceNotificationPhones(v []string)`

SetSecurityComplianceNotificationPhones sets SecurityComplianceNotificationPhones field to given value.

### HasSecurityComplianceNotificationPhones

`func (o *MicrosoftGraphOrganization) HasSecurityComplianceNotificationPhones() bool`

HasSecurityComplianceNotificationPhones returns a boolean if a field has been set.

### GetState

`func (o *MicrosoftGraphOrganization) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *MicrosoftGraphOrganization) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *MicrosoftGraphOrganization) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *MicrosoftGraphOrganization) HasState() bool`

HasState returns a boolean if a field has been set.

### SetStateNil

`func (o *MicrosoftGraphOrganization) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *MicrosoftGraphOrganization) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil
### GetStreet

`func (o *MicrosoftGraphOrganization) GetStreet() string`

GetStreet returns the Street field if non-nil, zero value otherwise.

### GetStreetOk

`func (o *MicrosoftGraphOrganization) GetStreetOk() (*string, bool)`

GetStreetOk returns a tuple with the Street field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreet

`func (o *MicrosoftGraphOrganization) SetStreet(v string)`

SetStreet sets Street field to given value.

### HasStreet

`func (o *MicrosoftGraphOrganization) HasStreet() bool`

HasStreet returns a boolean if a field has been set.

### SetStreetNil

`func (o *MicrosoftGraphOrganization) SetStreetNil(b bool)`

 SetStreetNil sets the value for Street to be an explicit nil

### UnsetStreet
`func (o *MicrosoftGraphOrganization) UnsetStreet()`

UnsetStreet ensures that no value is present for Street, not even an explicit nil
### GetTechnicalNotificationMails

`func (o *MicrosoftGraphOrganization) GetTechnicalNotificationMails() []string`

GetTechnicalNotificationMails returns the TechnicalNotificationMails field if non-nil, zero value otherwise.

### GetTechnicalNotificationMailsOk

`func (o *MicrosoftGraphOrganization) GetTechnicalNotificationMailsOk() (*[]string, bool)`

GetTechnicalNotificationMailsOk returns a tuple with the TechnicalNotificationMails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTechnicalNotificationMails

`func (o *MicrosoftGraphOrganization) SetTechnicalNotificationMails(v []string)`

SetTechnicalNotificationMails sets TechnicalNotificationMails field to given value.

### HasTechnicalNotificationMails

`func (o *MicrosoftGraphOrganization) HasTechnicalNotificationMails() bool`

HasTechnicalNotificationMails returns a boolean if a field has been set.

### GetTenantType

`func (o *MicrosoftGraphOrganization) GetTenantType() string`

GetTenantType returns the TenantType field if non-nil, zero value otherwise.

### GetTenantTypeOk

`func (o *MicrosoftGraphOrganization) GetTenantTypeOk() (*string, bool)`

GetTenantTypeOk returns a tuple with the TenantType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantType

`func (o *MicrosoftGraphOrganization) SetTenantType(v string)`

SetTenantType sets TenantType field to given value.

### HasTenantType

`func (o *MicrosoftGraphOrganization) HasTenantType() bool`

HasTenantType returns a boolean if a field has been set.

### SetTenantTypeNil

`func (o *MicrosoftGraphOrganization) SetTenantTypeNil(b bool)`

 SetTenantTypeNil sets the value for TenantType to be an explicit nil

### UnsetTenantType
`func (o *MicrosoftGraphOrganization) UnsetTenantType()`

UnsetTenantType ensures that no value is present for TenantType, not even an explicit nil
### GetVerifiedDomains

`func (o *MicrosoftGraphOrganization) GetVerifiedDomains() []MicrosoftGraphVerifiedDomain`

GetVerifiedDomains returns the VerifiedDomains field if non-nil, zero value otherwise.

### GetVerifiedDomainsOk

`func (o *MicrosoftGraphOrganization) GetVerifiedDomainsOk() (*[]MicrosoftGraphVerifiedDomain, bool)`

GetVerifiedDomainsOk returns a tuple with the VerifiedDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifiedDomains

`func (o *MicrosoftGraphOrganization) SetVerifiedDomains(v []MicrosoftGraphVerifiedDomain)`

SetVerifiedDomains sets VerifiedDomains field to given value.

### HasVerifiedDomains

`func (o *MicrosoftGraphOrganization) HasVerifiedDomains() bool`

HasVerifiedDomains returns a boolean if a field has been set.

### GetMobileDeviceManagementAuthority

`func (o *MicrosoftGraphOrganization) GetMobileDeviceManagementAuthority() AnyOfmicrosoftGraphMdmAuthority`

GetMobileDeviceManagementAuthority returns the MobileDeviceManagementAuthority field if non-nil, zero value otherwise.

### GetMobileDeviceManagementAuthorityOk

`func (o *MicrosoftGraphOrganization) GetMobileDeviceManagementAuthorityOk() (*AnyOfmicrosoftGraphMdmAuthority, bool)`

GetMobileDeviceManagementAuthorityOk returns a tuple with the MobileDeviceManagementAuthority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobileDeviceManagementAuthority

`func (o *MicrosoftGraphOrganization) SetMobileDeviceManagementAuthority(v AnyOfmicrosoftGraphMdmAuthority)`

SetMobileDeviceManagementAuthority sets MobileDeviceManagementAuthority field to given value.

### HasMobileDeviceManagementAuthority

`func (o *MicrosoftGraphOrganization) HasMobileDeviceManagementAuthority() bool`

HasMobileDeviceManagementAuthority returns a boolean if a field has been set.

### SetMobileDeviceManagementAuthorityNil

`func (o *MicrosoftGraphOrganization) SetMobileDeviceManagementAuthorityNil(b bool)`

 SetMobileDeviceManagementAuthorityNil sets the value for MobileDeviceManagementAuthority to be an explicit nil

### UnsetMobileDeviceManagementAuthority
`func (o *MicrosoftGraphOrganization) UnsetMobileDeviceManagementAuthority()`

UnsetMobileDeviceManagementAuthority ensures that no value is present for MobileDeviceManagementAuthority, not even an explicit nil
### GetBranding

`func (o *MicrosoftGraphOrganization) GetBranding() AnyOfmicrosoftGraphOrganizationalBranding`

GetBranding returns the Branding field if non-nil, zero value otherwise.

### GetBrandingOk

`func (o *MicrosoftGraphOrganization) GetBrandingOk() (*AnyOfmicrosoftGraphOrganizationalBranding, bool)`

GetBrandingOk returns a tuple with the Branding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranding

`func (o *MicrosoftGraphOrganization) SetBranding(v AnyOfmicrosoftGraphOrganizationalBranding)`

SetBranding sets Branding field to given value.

### HasBranding

`func (o *MicrosoftGraphOrganization) HasBranding() bool`

HasBranding returns a boolean if a field has been set.

### SetBrandingNil

`func (o *MicrosoftGraphOrganization) SetBrandingNil(b bool)`

 SetBrandingNil sets the value for Branding to be an explicit nil

### UnsetBranding
`func (o *MicrosoftGraphOrganization) UnsetBranding()`

UnsetBranding ensures that no value is present for Branding, not even an explicit nil
### GetCertificateBasedAuthConfiguration

`func (o *MicrosoftGraphOrganization) GetCertificateBasedAuthConfiguration() []MicrosoftGraphCertificateBasedAuthConfiguration`

GetCertificateBasedAuthConfiguration returns the CertificateBasedAuthConfiguration field if non-nil, zero value otherwise.

### GetCertificateBasedAuthConfigurationOk

`func (o *MicrosoftGraphOrganization) GetCertificateBasedAuthConfigurationOk() (*[]MicrosoftGraphCertificateBasedAuthConfiguration, bool)`

GetCertificateBasedAuthConfigurationOk returns a tuple with the CertificateBasedAuthConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateBasedAuthConfiguration

`func (o *MicrosoftGraphOrganization) SetCertificateBasedAuthConfiguration(v []MicrosoftGraphCertificateBasedAuthConfiguration)`

SetCertificateBasedAuthConfiguration sets CertificateBasedAuthConfiguration field to given value.

### HasCertificateBasedAuthConfiguration

`func (o *MicrosoftGraphOrganization) HasCertificateBasedAuthConfiguration() bool`

HasCertificateBasedAuthConfiguration returns a boolean if a field has been set.

### GetExtensions

`func (o *MicrosoftGraphOrganization) GetExtensions() []MicrosoftGraphExtension`

GetExtensions returns the Extensions field if non-nil, zero value otherwise.

### GetExtensionsOk

`func (o *MicrosoftGraphOrganization) GetExtensionsOk() (*[]MicrosoftGraphExtension, bool)`

GetExtensionsOk returns a tuple with the Extensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensions

`func (o *MicrosoftGraphOrganization) SetExtensions(v []MicrosoftGraphExtension)`

SetExtensions sets Extensions field to given value.

### HasExtensions

`func (o *MicrosoftGraphOrganization) HasExtensions() bool`

HasExtensions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


