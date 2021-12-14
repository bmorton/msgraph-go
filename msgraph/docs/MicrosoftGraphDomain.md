# MicrosoftGraphDomain

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
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

### NewMicrosoftGraphDomain

`func NewMicrosoftGraphDomain() *MicrosoftGraphDomain`

NewMicrosoftGraphDomain instantiates a new MicrosoftGraphDomain object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphDomainWithDefaults

`func NewMicrosoftGraphDomainWithDefaults() *MicrosoftGraphDomain`

NewMicrosoftGraphDomainWithDefaults instantiates a new MicrosoftGraphDomain object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphDomain) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphDomain) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphDomain) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphDomain) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAuthenticationType

`func (o *MicrosoftGraphDomain) GetAuthenticationType() string`

GetAuthenticationType returns the AuthenticationType field if non-nil, zero value otherwise.

### GetAuthenticationTypeOk

`func (o *MicrosoftGraphDomain) GetAuthenticationTypeOk() (*string, bool)`

GetAuthenticationTypeOk returns a tuple with the AuthenticationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationType

`func (o *MicrosoftGraphDomain) SetAuthenticationType(v string)`

SetAuthenticationType sets AuthenticationType field to given value.

### HasAuthenticationType

`func (o *MicrosoftGraphDomain) HasAuthenticationType() bool`

HasAuthenticationType returns a boolean if a field has been set.

### GetAvailabilityStatus

`func (o *MicrosoftGraphDomain) GetAvailabilityStatus() string`

GetAvailabilityStatus returns the AvailabilityStatus field if non-nil, zero value otherwise.

### GetAvailabilityStatusOk

`func (o *MicrosoftGraphDomain) GetAvailabilityStatusOk() (*string, bool)`

GetAvailabilityStatusOk returns a tuple with the AvailabilityStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityStatus

`func (o *MicrosoftGraphDomain) SetAvailabilityStatus(v string)`

SetAvailabilityStatus sets AvailabilityStatus field to given value.

### HasAvailabilityStatus

`func (o *MicrosoftGraphDomain) HasAvailabilityStatus() bool`

HasAvailabilityStatus returns a boolean if a field has been set.

### SetAvailabilityStatusNil

`func (o *MicrosoftGraphDomain) SetAvailabilityStatusNil(b bool)`

 SetAvailabilityStatusNil sets the value for AvailabilityStatus to be an explicit nil

### UnsetAvailabilityStatus
`func (o *MicrosoftGraphDomain) UnsetAvailabilityStatus()`

UnsetAvailabilityStatus ensures that no value is present for AvailabilityStatus, not even an explicit nil
### GetIsAdminManaged

`func (o *MicrosoftGraphDomain) GetIsAdminManaged() bool`

GetIsAdminManaged returns the IsAdminManaged field if non-nil, zero value otherwise.

### GetIsAdminManagedOk

`func (o *MicrosoftGraphDomain) GetIsAdminManagedOk() (*bool, bool)`

GetIsAdminManagedOk returns a tuple with the IsAdminManaged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAdminManaged

`func (o *MicrosoftGraphDomain) SetIsAdminManaged(v bool)`

SetIsAdminManaged sets IsAdminManaged field to given value.

### HasIsAdminManaged

`func (o *MicrosoftGraphDomain) HasIsAdminManaged() bool`

HasIsAdminManaged returns a boolean if a field has been set.

### GetIsDefault

`func (o *MicrosoftGraphDomain) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *MicrosoftGraphDomain) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *MicrosoftGraphDomain) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *MicrosoftGraphDomain) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetIsInitial

`func (o *MicrosoftGraphDomain) GetIsInitial() bool`

GetIsInitial returns the IsInitial field if non-nil, zero value otherwise.

### GetIsInitialOk

`func (o *MicrosoftGraphDomain) GetIsInitialOk() (*bool, bool)`

GetIsInitialOk returns a tuple with the IsInitial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInitial

`func (o *MicrosoftGraphDomain) SetIsInitial(v bool)`

SetIsInitial sets IsInitial field to given value.

### HasIsInitial

`func (o *MicrosoftGraphDomain) HasIsInitial() bool`

HasIsInitial returns a boolean if a field has been set.

### GetIsRoot

`func (o *MicrosoftGraphDomain) GetIsRoot() bool`

GetIsRoot returns the IsRoot field if non-nil, zero value otherwise.

### GetIsRootOk

`func (o *MicrosoftGraphDomain) GetIsRootOk() (*bool, bool)`

GetIsRootOk returns a tuple with the IsRoot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRoot

`func (o *MicrosoftGraphDomain) SetIsRoot(v bool)`

SetIsRoot sets IsRoot field to given value.

### HasIsRoot

`func (o *MicrosoftGraphDomain) HasIsRoot() bool`

HasIsRoot returns a boolean if a field has been set.

### GetIsVerified

`func (o *MicrosoftGraphDomain) GetIsVerified() bool`

GetIsVerified returns the IsVerified field if non-nil, zero value otherwise.

### GetIsVerifiedOk

`func (o *MicrosoftGraphDomain) GetIsVerifiedOk() (*bool, bool)`

GetIsVerifiedOk returns a tuple with the IsVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVerified

`func (o *MicrosoftGraphDomain) SetIsVerified(v bool)`

SetIsVerified sets IsVerified field to given value.

### HasIsVerified

`func (o *MicrosoftGraphDomain) HasIsVerified() bool`

HasIsVerified returns a boolean if a field has been set.

### GetManufacturer

`func (o *MicrosoftGraphDomain) GetManufacturer() string`

GetManufacturer returns the Manufacturer field if non-nil, zero value otherwise.

### GetManufacturerOk

`func (o *MicrosoftGraphDomain) GetManufacturerOk() (*string, bool)`

GetManufacturerOk returns a tuple with the Manufacturer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturer

`func (o *MicrosoftGraphDomain) SetManufacturer(v string)`

SetManufacturer sets Manufacturer field to given value.

### HasManufacturer

`func (o *MicrosoftGraphDomain) HasManufacturer() bool`

HasManufacturer returns a boolean if a field has been set.

### SetManufacturerNil

`func (o *MicrosoftGraphDomain) SetManufacturerNil(b bool)`

 SetManufacturerNil sets the value for Manufacturer to be an explicit nil

### UnsetManufacturer
`func (o *MicrosoftGraphDomain) UnsetManufacturer()`

UnsetManufacturer ensures that no value is present for Manufacturer, not even an explicit nil
### GetModel

`func (o *MicrosoftGraphDomain) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *MicrosoftGraphDomain) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *MicrosoftGraphDomain) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *MicrosoftGraphDomain) HasModel() bool`

HasModel returns a boolean if a field has been set.

### SetModelNil

`func (o *MicrosoftGraphDomain) SetModelNil(b bool)`

 SetModelNil sets the value for Model to be an explicit nil

### UnsetModel
`func (o *MicrosoftGraphDomain) UnsetModel()`

UnsetModel ensures that no value is present for Model, not even an explicit nil
### GetPasswordNotificationWindowInDays

`func (o *MicrosoftGraphDomain) GetPasswordNotificationWindowInDays() int32`

GetPasswordNotificationWindowInDays returns the PasswordNotificationWindowInDays field if non-nil, zero value otherwise.

### GetPasswordNotificationWindowInDaysOk

`func (o *MicrosoftGraphDomain) GetPasswordNotificationWindowInDaysOk() (*int32, bool)`

GetPasswordNotificationWindowInDaysOk returns a tuple with the PasswordNotificationWindowInDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordNotificationWindowInDays

`func (o *MicrosoftGraphDomain) SetPasswordNotificationWindowInDays(v int32)`

SetPasswordNotificationWindowInDays sets PasswordNotificationWindowInDays field to given value.

### HasPasswordNotificationWindowInDays

`func (o *MicrosoftGraphDomain) HasPasswordNotificationWindowInDays() bool`

HasPasswordNotificationWindowInDays returns a boolean if a field has been set.

### SetPasswordNotificationWindowInDaysNil

`func (o *MicrosoftGraphDomain) SetPasswordNotificationWindowInDaysNil(b bool)`

 SetPasswordNotificationWindowInDaysNil sets the value for PasswordNotificationWindowInDays to be an explicit nil

### UnsetPasswordNotificationWindowInDays
`func (o *MicrosoftGraphDomain) UnsetPasswordNotificationWindowInDays()`

UnsetPasswordNotificationWindowInDays ensures that no value is present for PasswordNotificationWindowInDays, not even an explicit nil
### GetPasswordValidityPeriodInDays

`func (o *MicrosoftGraphDomain) GetPasswordValidityPeriodInDays() int32`

GetPasswordValidityPeriodInDays returns the PasswordValidityPeriodInDays field if non-nil, zero value otherwise.

### GetPasswordValidityPeriodInDaysOk

`func (o *MicrosoftGraphDomain) GetPasswordValidityPeriodInDaysOk() (*int32, bool)`

GetPasswordValidityPeriodInDaysOk returns a tuple with the PasswordValidityPeriodInDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordValidityPeriodInDays

`func (o *MicrosoftGraphDomain) SetPasswordValidityPeriodInDays(v int32)`

SetPasswordValidityPeriodInDays sets PasswordValidityPeriodInDays field to given value.

### HasPasswordValidityPeriodInDays

`func (o *MicrosoftGraphDomain) HasPasswordValidityPeriodInDays() bool`

HasPasswordValidityPeriodInDays returns a boolean if a field has been set.

### SetPasswordValidityPeriodInDaysNil

`func (o *MicrosoftGraphDomain) SetPasswordValidityPeriodInDaysNil(b bool)`

 SetPasswordValidityPeriodInDaysNil sets the value for PasswordValidityPeriodInDays to be an explicit nil

### UnsetPasswordValidityPeriodInDays
`func (o *MicrosoftGraphDomain) UnsetPasswordValidityPeriodInDays()`

UnsetPasswordValidityPeriodInDays ensures that no value is present for PasswordValidityPeriodInDays, not even an explicit nil
### GetState

`func (o *MicrosoftGraphDomain) GetState() AnyOfmicrosoftGraphDomainState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *MicrosoftGraphDomain) GetStateOk() (*AnyOfmicrosoftGraphDomainState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *MicrosoftGraphDomain) SetState(v AnyOfmicrosoftGraphDomainState)`

SetState sets State field to given value.

### HasState

`func (o *MicrosoftGraphDomain) HasState() bool`

HasState returns a boolean if a field has been set.

### SetStateNil

`func (o *MicrosoftGraphDomain) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *MicrosoftGraphDomain) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil
### GetSupportedServices

`func (o *MicrosoftGraphDomain) GetSupportedServices() []string`

GetSupportedServices returns the SupportedServices field if non-nil, zero value otherwise.

### GetSupportedServicesOk

`func (o *MicrosoftGraphDomain) GetSupportedServicesOk() (*[]string, bool)`

GetSupportedServicesOk returns a tuple with the SupportedServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedServices

`func (o *MicrosoftGraphDomain) SetSupportedServices(v []string)`

SetSupportedServices sets SupportedServices field to given value.

### HasSupportedServices

`func (o *MicrosoftGraphDomain) HasSupportedServices() bool`

HasSupportedServices returns a boolean if a field has been set.

### GetDomainNameReferences

`func (o *MicrosoftGraphDomain) GetDomainNameReferences() []MicrosoftGraphDirectoryObject`

GetDomainNameReferences returns the DomainNameReferences field if non-nil, zero value otherwise.

### GetDomainNameReferencesOk

`func (o *MicrosoftGraphDomain) GetDomainNameReferencesOk() (*[]MicrosoftGraphDirectoryObject, bool)`

GetDomainNameReferencesOk returns a tuple with the DomainNameReferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainNameReferences

`func (o *MicrosoftGraphDomain) SetDomainNameReferences(v []MicrosoftGraphDirectoryObject)`

SetDomainNameReferences sets DomainNameReferences field to given value.

### HasDomainNameReferences

`func (o *MicrosoftGraphDomain) HasDomainNameReferences() bool`

HasDomainNameReferences returns a boolean if a field has been set.

### GetServiceConfigurationRecords

`func (o *MicrosoftGraphDomain) GetServiceConfigurationRecords() []MicrosoftGraphDomainDnsRecord`

GetServiceConfigurationRecords returns the ServiceConfigurationRecords field if non-nil, zero value otherwise.

### GetServiceConfigurationRecordsOk

`func (o *MicrosoftGraphDomain) GetServiceConfigurationRecordsOk() (*[]MicrosoftGraphDomainDnsRecord, bool)`

GetServiceConfigurationRecordsOk returns a tuple with the ServiceConfigurationRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceConfigurationRecords

`func (o *MicrosoftGraphDomain) SetServiceConfigurationRecords(v []MicrosoftGraphDomainDnsRecord)`

SetServiceConfigurationRecords sets ServiceConfigurationRecords field to given value.

### HasServiceConfigurationRecords

`func (o *MicrosoftGraphDomain) HasServiceConfigurationRecords() bool`

HasServiceConfigurationRecords returns a boolean if a field has been set.

### GetVerificationDnsRecords

`func (o *MicrosoftGraphDomain) GetVerificationDnsRecords() []MicrosoftGraphDomainDnsRecord`

GetVerificationDnsRecords returns the VerificationDnsRecords field if non-nil, zero value otherwise.

### GetVerificationDnsRecordsOk

`func (o *MicrosoftGraphDomain) GetVerificationDnsRecordsOk() (*[]MicrosoftGraphDomainDnsRecord, bool)`

GetVerificationDnsRecordsOk returns a tuple with the VerificationDnsRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationDnsRecords

`func (o *MicrosoftGraphDomain) SetVerificationDnsRecords(v []MicrosoftGraphDomainDnsRecord)`

SetVerificationDnsRecords sets VerificationDnsRecords field to given value.

### HasVerificationDnsRecords

`func (o *MicrosoftGraphDomain) HasVerificationDnsRecords() bool`

HasVerificationDnsRecords returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


