# Domain

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthenticationType** | Pointer to **string** | Indicates the configured authentication type for the domain. The value is either Managed or Federated. Managed indicates a cloud managed domain where Azure AD performs user authentication. Federated indicates authentication is federated with an identity provider such as the tenant&#39;s on-premises Active Directory via Active Directory Federation Services. This property is read-only and is not nullable. | [optional] 
**AvailabilityStatus** | Pointer to **NullableString** | This property is always null except when the verify action is used. When the verify action is used, a domain entity is returned in the response. The availabilityStatus property of the domain entity in the response is either AvailableImmediately or EmailVerifiedDomainTakeoverScheduled. | [optional] 
**IsAdminManaged** | Pointer to **bool** | The value of the property is false if the DNS record management of the domain has been delegated to Microsoft 365. Otherwise, the value is true. Not nullable | [optional] 
**IsDefault** | Pointer to **bool** | true if this is the default domain that is used for user creation. There is only one default domain per company. Not nullable | [optional] 
**IsInitial** | Pointer to **bool** | true if this is the initial domain created by Microsoft Online Services (companyname.onmicrosoft.com). There is only one initial domain per company. Not nullable | [optional] 
**IsRoot** | Pointer to **bool** | true if the domain is a verified root domain. Otherwise, false if the domain is a subdomain or unverified. Not nullable | [optional] 
**IsVerified** | Pointer to **bool** | true if the domain has completed domain ownership verification. Not nullable | [optional] 
**Manufacturer** | Pointer to **NullableString** |  | [optional] 
**Model** | Pointer to **NullableString** |  | [optional] 
**PasswordNotificationWindowInDays** | Pointer to **NullableInt32** | Specifies the number of days before a user receives notification that their password will expire. If the property is not set, a default value of 14 days will be used. | [optional] 
**PasswordValidityPeriodInDays** | Pointer to **NullableInt32** | Specifies the length of time that a password is valid before it must be changed. If the property is not set, a default value of 90 days will be used. | [optional] 
**State** | Pointer to [**AnyOfmicrosoftGraphDomainState**](anyOf&lt;microsoft.graph.domainState&gt;.md) | Status of asynchronous operations scheduled for the domain. | [optional] 
**SupportedServices** | Pointer to **[]string** | The capabilities assigned to the domain. Can include 0, 1 or more of following values: Email, Sharepoint, EmailInternalRelayOnly, OfficeCommunicationsOnline, SharePointDefaultDomain, FullRedelegation, SharePointPublic, OrgIdAuthentication, Yammer, Intune. The values which you can add/remove using Graph API include: Email, OfficeCommunicationsOnline, Yammer. Not nullable | [optional] 
**DomainNameReferences** | Pointer to [**[]MicrosoftGraphDirectoryObject**](MicrosoftGraphDirectoryObject.md) | Read-only, Nullable | [optional] 
**ServiceConfigurationRecords** | Pointer to [**[]MicrosoftGraphDomainDnsRecord**](MicrosoftGraphDomainDnsRecord.md) | DNS records the customer adds to the DNS zone file of the domain before the domain can be used by Microsoft Online services. Read-only, Nullable | [optional] 
**VerificationDnsRecords** | Pointer to [**[]MicrosoftGraphDomainDnsRecord**](MicrosoftGraphDomainDnsRecord.md) | DNS records that the customer adds to the DNS zone file of the domain before the customer can complete domain ownership verification with Azure AD. Read-only, Nullable | [optional] 

## Methods

### NewDomain

`func NewDomain() *Domain`

NewDomain instantiates a new Domain object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainWithDefaults

`func NewDomainWithDefaults() *Domain`

NewDomainWithDefaults instantiates a new Domain object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthenticationType

`func (o *Domain) GetAuthenticationType() string`

GetAuthenticationType returns the AuthenticationType field if non-nil, zero value otherwise.

### GetAuthenticationTypeOk

`func (o *Domain) GetAuthenticationTypeOk() (*string, bool)`

GetAuthenticationTypeOk returns a tuple with the AuthenticationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationType

`func (o *Domain) SetAuthenticationType(v string)`

SetAuthenticationType sets AuthenticationType field to given value.

### HasAuthenticationType

`func (o *Domain) HasAuthenticationType() bool`

HasAuthenticationType returns a boolean if a field has been set.

### GetAvailabilityStatus

`func (o *Domain) GetAvailabilityStatus() string`

GetAvailabilityStatus returns the AvailabilityStatus field if non-nil, zero value otherwise.

### GetAvailabilityStatusOk

`func (o *Domain) GetAvailabilityStatusOk() (*string, bool)`

GetAvailabilityStatusOk returns a tuple with the AvailabilityStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityStatus

`func (o *Domain) SetAvailabilityStatus(v string)`

SetAvailabilityStatus sets AvailabilityStatus field to given value.

### HasAvailabilityStatus

`func (o *Domain) HasAvailabilityStatus() bool`

HasAvailabilityStatus returns a boolean if a field has been set.

### SetAvailabilityStatusNil

`func (o *Domain) SetAvailabilityStatusNil(b bool)`

 SetAvailabilityStatusNil sets the value for AvailabilityStatus to be an explicit nil

### UnsetAvailabilityStatus
`func (o *Domain) UnsetAvailabilityStatus()`

UnsetAvailabilityStatus ensures that no value is present for AvailabilityStatus, not even an explicit nil
### GetIsAdminManaged

`func (o *Domain) GetIsAdminManaged() bool`

GetIsAdminManaged returns the IsAdminManaged field if non-nil, zero value otherwise.

### GetIsAdminManagedOk

`func (o *Domain) GetIsAdminManagedOk() (*bool, bool)`

GetIsAdminManagedOk returns a tuple with the IsAdminManaged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAdminManaged

`func (o *Domain) SetIsAdminManaged(v bool)`

SetIsAdminManaged sets IsAdminManaged field to given value.

### HasIsAdminManaged

`func (o *Domain) HasIsAdminManaged() bool`

HasIsAdminManaged returns a boolean if a field has been set.

### GetIsDefault

`func (o *Domain) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *Domain) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *Domain) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *Domain) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetIsInitial

`func (o *Domain) GetIsInitial() bool`

GetIsInitial returns the IsInitial field if non-nil, zero value otherwise.

### GetIsInitialOk

`func (o *Domain) GetIsInitialOk() (*bool, bool)`

GetIsInitialOk returns a tuple with the IsInitial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInitial

`func (o *Domain) SetIsInitial(v bool)`

SetIsInitial sets IsInitial field to given value.

### HasIsInitial

`func (o *Domain) HasIsInitial() bool`

HasIsInitial returns a boolean if a field has been set.

### GetIsRoot

`func (o *Domain) GetIsRoot() bool`

GetIsRoot returns the IsRoot field if non-nil, zero value otherwise.

### GetIsRootOk

`func (o *Domain) GetIsRootOk() (*bool, bool)`

GetIsRootOk returns a tuple with the IsRoot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRoot

`func (o *Domain) SetIsRoot(v bool)`

SetIsRoot sets IsRoot field to given value.

### HasIsRoot

`func (o *Domain) HasIsRoot() bool`

HasIsRoot returns a boolean if a field has been set.

### GetIsVerified

`func (o *Domain) GetIsVerified() bool`

GetIsVerified returns the IsVerified field if non-nil, zero value otherwise.

### GetIsVerifiedOk

`func (o *Domain) GetIsVerifiedOk() (*bool, bool)`

GetIsVerifiedOk returns a tuple with the IsVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVerified

`func (o *Domain) SetIsVerified(v bool)`

SetIsVerified sets IsVerified field to given value.

### HasIsVerified

`func (o *Domain) HasIsVerified() bool`

HasIsVerified returns a boolean if a field has been set.

### GetManufacturer

`func (o *Domain) GetManufacturer() string`

GetManufacturer returns the Manufacturer field if non-nil, zero value otherwise.

### GetManufacturerOk

`func (o *Domain) GetManufacturerOk() (*string, bool)`

GetManufacturerOk returns a tuple with the Manufacturer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturer

`func (o *Domain) SetManufacturer(v string)`

SetManufacturer sets Manufacturer field to given value.

### HasManufacturer

`func (o *Domain) HasManufacturer() bool`

HasManufacturer returns a boolean if a field has been set.

### SetManufacturerNil

`func (o *Domain) SetManufacturerNil(b bool)`

 SetManufacturerNil sets the value for Manufacturer to be an explicit nil

### UnsetManufacturer
`func (o *Domain) UnsetManufacturer()`

UnsetManufacturer ensures that no value is present for Manufacturer, not even an explicit nil
### GetModel

`func (o *Domain) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *Domain) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *Domain) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *Domain) HasModel() bool`

HasModel returns a boolean if a field has been set.

### SetModelNil

`func (o *Domain) SetModelNil(b bool)`

 SetModelNil sets the value for Model to be an explicit nil

### UnsetModel
`func (o *Domain) UnsetModel()`

UnsetModel ensures that no value is present for Model, not even an explicit nil
### GetPasswordNotificationWindowInDays

`func (o *Domain) GetPasswordNotificationWindowInDays() int32`

GetPasswordNotificationWindowInDays returns the PasswordNotificationWindowInDays field if non-nil, zero value otherwise.

### GetPasswordNotificationWindowInDaysOk

`func (o *Domain) GetPasswordNotificationWindowInDaysOk() (*int32, bool)`

GetPasswordNotificationWindowInDaysOk returns a tuple with the PasswordNotificationWindowInDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordNotificationWindowInDays

`func (o *Domain) SetPasswordNotificationWindowInDays(v int32)`

SetPasswordNotificationWindowInDays sets PasswordNotificationWindowInDays field to given value.

### HasPasswordNotificationWindowInDays

`func (o *Domain) HasPasswordNotificationWindowInDays() bool`

HasPasswordNotificationWindowInDays returns a boolean if a field has been set.

### SetPasswordNotificationWindowInDaysNil

`func (o *Domain) SetPasswordNotificationWindowInDaysNil(b bool)`

 SetPasswordNotificationWindowInDaysNil sets the value for PasswordNotificationWindowInDays to be an explicit nil

### UnsetPasswordNotificationWindowInDays
`func (o *Domain) UnsetPasswordNotificationWindowInDays()`

UnsetPasswordNotificationWindowInDays ensures that no value is present for PasswordNotificationWindowInDays, not even an explicit nil
### GetPasswordValidityPeriodInDays

`func (o *Domain) GetPasswordValidityPeriodInDays() int32`

GetPasswordValidityPeriodInDays returns the PasswordValidityPeriodInDays field if non-nil, zero value otherwise.

### GetPasswordValidityPeriodInDaysOk

`func (o *Domain) GetPasswordValidityPeriodInDaysOk() (*int32, bool)`

GetPasswordValidityPeriodInDaysOk returns a tuple with the PasswordValidityPeriodInDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordValidityPeriodInDays

`func (o *Domain) SetPasswordValidityPeriodInDays(v int32)`

SetPasswordValidityPeriodInDays sets PasswordValidityPeriodInDays field to given value.

### HasPasswordValidityPeriodInDays

`func (o *Domain) HasPasswordValidityPeriodInDays() bool`

HasPasswordValidityPeriodInDays returns a boolean if a field has been set.

### SetPasswordValidityPeriodInDaysNil

`func (o *Domain) SetPasswordValidityPeriodInDaysNil(b bool)`

 SetPasswordValidityPeriodInDaysNil sets the value for PasswordValidityPeriodInDays to be an explicit nil

### UnsetPasswordValidityPeriodInDays
`func (o *Domain) UnsetPasswordValidityPeriodInDays()`

UnsetPasswordValidityPeriodInDays ensures that no value is present for PasswordValidityPeriodInDays, not even an explicit nil
### GetState

`func (o *Domain) GetState() AnyOfmicrosoftGraphDomainState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Domain) GetStateOk() (*AnyOfmicrosoftGraphDomainState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Domain) SetState(v AnyOfmicrosoftGraphDomainState)`

SetState sets State field to given value.

### HasState

`func (o *Domain) HasState() bool`

HasState returns a boolean if a field has been set.

### SetStateNil

`func (o *Domain) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *Domain) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil
### GetSupportedServices

`func (o *Domain) GetSupportedServices() []string`

GetSupportedServices returns the SupportedServices field if non-nil, zero value otherwise.

### GetSupportedServicesOk

`func (o *Domain) GetSupportedServicesOk() (*[]string, bool)`

GetSupportedServicesOk returns a tuple with the SupportedServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedServices

`func (o *Domain) SetSupportedServices(v []string)`

SetSupportedServices sets SupportedServices field to given value.

### HasSupportedServices

`func (o *Domain) HasSupportedServices() bool`

HasSupportedServices returns a boolean if a field has been set.

### GetDomainNameReferences

`func (o *Domain) GetDomainNameReferences() []MicrosoftGraphDirectoryObject`

GetDomainNameReferences returns the DomainNameReferences field if non-nil, zero value otherwise.

### GetDomainNameReferencesOk

`func (o *Domain) GetDomainNameReferencesOk() (*[]MicrosoftGraphDirectoryObject, bool)`

GetDomainNameReferencesOk returns a tuple with the DomainNameReferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainNameReferences

`func (o *Domain) SetDomainNameReferences(v []MicrosoftGraphDirectoryObject)`

SetDomainNameReferences sets DomainNameReferences field to given value.

### HasDomainNameReferences

`func (o *Domain) HasDomainNameReferences() bool`

HasDomainNameReferences returns a boolean if a field has been set.

### GetServiceConfigurationRecords

`func (o *Domain) GetServiceConfigurationRecords() []MicrosoftGraphDomainDnsRecord`

GetServiceConfigurationRecords returns the ServiceConfigurationRecords field if non-nil, zero value otherwise.

### GetServiceConfigurationRecordsOk

`func (o *Domain) GetServiceConfigurationRecordsOk() (*[]MicrosoftGraphDomainDnsRecord, bool)`

GetServiceConfigurationRecordsOk returns a tuple with the ServiceConfigurationRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceConfigurationRecords

`func (o *Domain) SetServiceConfigurationRecords(v []MicrosoftGraphDomainDnsRecord)`

SetServiceConfigurationRecords sets ServiceConfigurationRecords field to given value.

### HasServiceConfigurationRecords

`func (o *Domain) HasServiceConfigurationRecords() bool`

HasServiceConfigurationRecords returns a boolean if a field has been set.

### GetVerificationDnsRecords

`func (o *Domain) GetVerificationDnsRecords() []MicrosoftGraphDomainDnsRecord`

GetVerificationDnsRecords returns the VerificationDnsRecords field if non-nil, zero value otherwise.

### GetVerificationDnsRecordsOk

`func (o *Domain) GetVerificationDnsRecordsOk() (*[]MicrosoftGraphDomainDnsRecord, bool)`

GetVerificationDnsRecordsOk returns a tuple with the VerificationDnsRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationDnsRecords

`func (o *Domain) SetVerificationDnsRecords(v []MicrosoftGraphDomainDnsRecord)`

SetVerificationDnsRecords sets VerificationDnsRecords field to given value.

### HasVerificationDnsRecords

`func (o *Domain) HasVerificationDnsRecords() bool`

HasVerificationDnsRecords returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


