# MicrosoftGraphEducationUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**RelatedContacts** | Pointer to [**[]AnyOfmicrosoftGraphRelatedContact**](AnyOfmicrosoftGraphRelatedContact.md) | Related records related to the user. Possible relationships are parent, relative, aide, doctor, guardian, child, other, unknownFutureValue | [optional] 
**AccountEnabled** | Pointer to **NullableBool** | True if the account is enabled; otherwise, false. This property is required when a user is created. Supports $filter. | [optional] 
**AssignedLicenses** | Pointer to [**[]MicrosoftGraphAssignedLicense**](MicrosoftGraphAssignedLicense.md) | The licenses that are assigned to the user. Not nullable. | [optional] 
**AssignedPlans** | Pointer to [**[]MicrosoftGraphAssignedPlan**](MicrosoftGraphAssignedPlan.md) | The plans that are assigned to the user. Read-only. Not nullable. | [optional] 
**BusinessPhones** | Pointer to **[]string** | The telephone numbers for the user. Note: Although this is a string collection, only one number can be set for this property. | [optional] 
**CreatedBy** | Pointer to [**AnyOfmicrosoftGraphIdentitySet**](anyOf&lt;microsoft.graph.identitySet&gt;.md) | Entity who created the user. | [optional] 
**Department** | Pointer to **NullableString** | The name for the department in which the user works. Supports $filter. | [optional] 
**DisplayName** | Pointer to **NullableString** | The name displayed in the address book for the user. This is usually the combination of the user&#39;s first name, middle initial, and last name. This property is required when a user is created and it cannot be cleared during updates. Supports $filter and $orderby. | [optional] 
**ExternalSource** | Pointer to [**AnyOfmicrosoftGraphEducationExternalSource**](anyOf&lt;microsoft.graph.educationExternalSource&gt;.md) | Where this user was created from. Possible values are: sis, manual. | [optional] 
**ExternalSourceDetail** | Pointer to **NullableString** | The name of the external source this resources was generated from. | [optional] 
**GivenName** | Pointer to **NullableString** | The given name (first name) of the user. Supports $filter. | [optional] 
**Mail** | Pointer to **NullableString** | The SMTP address for the user; for example, jeff@contoso.onmicrosoft.com. Read-Only. Supports $filter. | [optional] 
**MailingAddress** | Pointer to [**AnyOfmicrosoftGraphPhysicalAddress**](anyOf&lt;microsoft.graph.physicalAddress&gt;.md) | Mail address of user. | [optional] 
**MailNickname** | Pointer to **NullableString** | The mail alias for the user. This property must be specified when a user is created. Supports $filter. | [optional] 
**MiddleName** | Pointer to **NullableString** | The middle name of user. | [optional] 
**MobilePhone** | Pointer to **NullableString** | The primary cellular telephone number for the user. | [optional] 
**OfficeLocation** | Pointer to **NullableString** |  | [optional] 
**OnPremisesInfo** | Pointer to [**AnyOfmicrosoftGraphEducationOnPremisesInfo**](anyOf&lt;microsoft.graph.educationOnPremisesInfo&gt;.md) | Additional information used to associate the Azure AD user with its Active Directory counterpart. | [optional] 
**PasswordPolicies** | Pointer to **NullableString** | Specifies password policies for the user. This value is an enumeration with one possible value being DisableStrongPassword, which allows weaker passwords than the default policy to be specified. DisablePasswordExpiration can also be specified. The two can be specified together; for example: DisablePasswordExpiration, DisableStrongPassword. | [optional] 
**PasswordProfile** | Pointer to [**AnyOfmicrosoftGraphPasswordProfile**](anyOf&lt;microsoft.graph.passwordProfile&gt;.md) | Specifies the password profile for the user. The profile contains the user&#39;s password. This property is required when a user is created. The password in the profile must satisfy minimum requirements as specified by the passwordPolicies property. By default, a strong password is required. | [optional] 
**PreferredLanguage** | Pointer to **NullableString** | The preferred language for the user. Should follow ISO 639-1 Code; for example, &#39;en-US&#39;. | [optional] 
**PrimaryRole** | Pointer to [**AnyOfmicrosoftGraphEducationUserRole**](anyOf&lt;microsoft.graph.educationUserRole&gt;.md) | Default role for a user. The user&#39;s role might be different in an individual class. Possible values are: student, teacher, none, unknownFutureValue. | [optional] 
**ProvisionedPlans** | Pointer to [**[]MicrosoftGraphProvisionedPlan**](MicrosoftGraphProvisionedPlan.md) | The plans that are provisioned for the user. Read-only. Not nullable. | [optional] 
**RefreshTokensValidFromDateTime** | Pointer to **NullableTime** |  | [optional] 
**ResidenceAddress** | Pointer to [**AnyOfmicrosoftGraphPhysicalAddress**](anyOf&lt;microsoft.graph.physicalAddress&gt;.md) | Address where user lives. | [optional] 
**ShowInAddressList** | Pointer to **NullableBool** | true if the Outlook global address list should contain this user, otherwise false. If not set, this will be treated as true. For users invited through the invitation manager, this property will be set to false. | [optional] 
**Student** | Pointer to [**AnyOfmicrosoftGraphEducationStudent**](anyOf&lt;microsoft.graph.educationStudent&gt;.md) | If the primary role is student, this block will contain student specific data. | [optional] 
**Surname** | Pointer to **NullableString** | The user&#39;s surname (family name or last name). Supports $filter. | [optional] 
**Teacher** | Pointer to [**AnyOfmicrosoftGraphEducationTeacher**](anyOf&lt;microsoft.graph.educationTeacher&gt;.md) | If the primary role is teacher, this block will contain teacher specific data. | [optional] 
**UsageLocation** | Pointer to **NullableString** | A two-letter country code (ISO standard 3166). Required for users who will be assigned licenses due to a legal requirement to check for availability of services in countries or regions. Examples include: &#39;US&#39;, &#39;JP&#39;, and &#39;GB&#39;. Not nullable. Supports $filter. | [optional] 
**UserPrincipalName** | Pointer to **NullableString** | The user principal name (UPN) of the user. The UPN is an Internet-style login name for the user based on the Internet standard RFC 822. By convention, this should map to the user&#39;s email name. The general format is alias@domain, where domain must be present in the tenant&#39;s collection of verified domains. This property is required when a user is created. The verified domains for the tenant can be accessed from the verifiedDomains property of organization. Supports $filter and $orderby. | [optional] 
**UserType** | Pointer to **NullableString** | A string value that can be used to classify user types in your directory, such as &#39;Member&#39; and &#39;Guest&#39;. Supports $filter. | [optional] 
**Rubrics** | Pointer to [**[]MicrosoftGraphEducationRubric**](MicrosoftGraphEducationRubric.md) |  | [optional] 
**Classes** | Pointer to [**[]MicrosoftGraphEducationClass**](MicrosoftGraphEducationClass.md) | Classes to which the user belongs. Nullable. | [optional] 
**Schools** | Pointer to [**[]MicrosoftGraphEducationSchool**](MicrosoftGraphEducationSchool.md) | Schools to which the user belongs. Nullable. | [optional] 
**TaughtClasses** | Pointer to [**[]MicrosoftGraphEducationClass**](MicrosoftGraphEducationClass.md) | Classes for which the user is a teacher. | [optional] 
**User** | Pointer to [**AnyOfmicrosoftGraphUser**](anyOf&lt;microsoft.graph.user&gt;.md) | The directory user corresponding to this user. | [optional] 

## Methods

### NewMicrosoftGraphEducationUser

`func NewMicrosoftGraphEducationUser() *MicrosoftGraphEducationUser`

NewMicrosoftGraphEducationUser instantiates a new MicrosoftGraphEducationUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphEducationUserWithDefaults

`func NewMicrosoftGraphEducationUserWithDefaults() *MicrosoftGraphEducationUser`

NewMicrosoftGraphEducationUserWithDefaults instantiates a new MicrosoftGraphEducationUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphEducationUser) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphEducationUser) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphEducationUser) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphEducationUser) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRelatedContacts

`func (o *MicrosoftGraphEducationUser) GetRelatedContacts() []*AnyOfmicrosoftGraphRelatedContact`

GetRelatedContacts returns the RelatedContacts field if non-nil, zero value otherwise.

### GetRelatedContactsOk

`func (o *MicrosoftGraphEducationUser) GetRelatedContactsOk() (*[]*AnyOfmicrosoftGraphRelatedContact, bool)`

GetRelatedContactsOk returns a tuple with the RelatedContacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedContacts

`func (o *MicrosoftGraphEducationUser) SetRelatedContacts(v []*AnyOfmicrosoftGraphRelatedContact)`

SetRelatedContacts sets RelatedContacts field to given value.

### HasRelatedContacts

`func (o *MicrosoftGraphEducationUser) HasRelatedContacts() bool`

HasRelatedContacts returns a boolean if a field has been set.

### GetAccountEnabled

`func (o *MicrosoftGraphEducationUser) GetAccountEnabled() bool`

GetAccountEnabled returns the AccountEnabled field if non-nil, zero value otherwise.

### GetAccountEnabledOk

`func (o *MicrosoftGraphEducationUser) GetAccountEnabledOk() (*bool, bool)`

GetAccountEnabledOk returns a tuple with the AccountEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountEnabled

`func (o *MicrosoftGraphEducationUser) SetAccountEnabled(v bool)`

SetAccountEnabled sets AccountEnabled field to given value.

### HasAccountEnabled

`func (o *MicrosoftGraphEducationUser) HasAccountEnabled() bool`

HasAccountEnabled returns a boolean if a field has been set.

### SetAccountEnabledNil

`func (o *MicrosoftGraphEducationUser) SetAccountEnabledNil(b bool)`

 SetAccountEnabledNil sets the value for AccountEnabled to be an explicit nil

### UnsetAccountEnabled
`func (o *MicrosoftGraphEducationUser) UnsetAccountEnabled()`

UnsetAccountEnabled ensures that no value is present for AccountEnabled, not even an explicit nil
### GetAssignedLicenses

`func (o *MicrosoftGraphEducationUser) GetAssignedLicenses() []MicrosoftGraphAssignedLicense`

GetAssignedLicenses returns the AssignedLicenses field if non-nil, zero value otherwise.

### GetAssignedLicensesOk

`func (o *MicrosoftGraphEducationUser) GetAssignedLicensesOk() (*[]MicrosoftGraphAssignedLicense, bool)`

GetAssignedLicensesOk returns a tuple with the AssignedLicenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedLicenses

`func (o *MicrosoftGraphEducationUser) SetAssignedLicenses(v []MicrosoftGraphAssignedLicense)`

SetAssignedLicenses sets AssignedLicenses field to given value.

### HasAssignedLicenses

`func (o *MicrosoftGraphEducationUser) HasAssignedLicenses() bool`

HasAssignedLicenses returns a boolean if a field has been set.

### GetAssignedPlans

`func (o *MicrosoftGraphEducationUser) GetAssignedPlans() []MicrosoftGraphAssignedPlan`

GetAssignedPlans returns the AssignedPlans field if non-nil, zero value otherwise.

### GetAssignedPlansOk

`func (o *MicrosoftGraphEducationUser) GetAssignedPlansOk() (*[]MicrosoftGraphAssignedPlan, bool)`

GetAssignedPlansOk returns a tuple with the AssignedPlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedPlans

`func (o *MicrosoftGraphEducationUser) SetAssignedPlans(v []MicrosoftGraphAssignedPlan)`

SetAssignedPlans sets AssignedPlans field to given value.

### HasAssignedPlans

`func (o *MicrosoftGraphEducationUser) HasAssignedPlans() bool`

HasAssignedPlans returns a boolean if a field has been set.

### GetBusinessPhones

`func (o *MicrosoftGraphEducationUser) GetBusinessPhones() []string`

GetBusinessPhones returns the BusinessPhones field if non-nil, zero value otherwise.

### GetBusinessPhonesOk

`func (o *MicrosoftGraphEducationUser) GetBusinessPhonesOk() (*[]string, bool)`

GetBusinessPhonesOk returns a tuple with the BusinessPhones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessPhones

`func (o *MicrosoftGraphEducationUser) SetBusinessPhones(v []string)`

SetBusinessPhones sets BusinessPhones field to given value.

### HasBusinessPhones

`func (o *MicrosoftGraphEducationUser) HasBusinessPhones() bool`

HasBusinessPhones returns a boolean if a field has been set.

### GetCreatedBy

`func (o *MicrosoftGraphEducationUser) GetCreatedBy() AnyOfmicrosoftGraphIdentitySet`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *MicrosoftGraphEducationUser) GetCreatedByOk() (*AnyOfmicrosoftGraphIdentitySet, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *MicrosoftGraphEducationUser) SetCreatedBy(v AnyOfmicrosoftGraphIdentitySet)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *MicrosoftGraphEducationUser) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedByNil

`func (o *MicrosoftGraphEducationUser) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *MicrosoftGraphEducationUser) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
### GetDepartment

`func (o *MicrosoftGraphEducationUser) GetDepartment() string`

GetDepartment returns the Department field if non-nil, zero value otherwise.

### GetDepartmentOk

`func (o *MicrosoftGraphEducationUser) GetDepartmentOk() (*string, bool)`

GetDepartmentOk returns a tuple with the Department field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartment

`func (o *MicrosoftGraphEducationUser) SetDepartment(v string)`

SetDepartment sets Department field to given value.

### HasDepartment

`func (o *MicrosoftGraphEducationUser) HasDepartment() bool`

HasDepartment returns a boolean if a field has been set.

### SetDepartmentNil

`func (o *MicrosoftGraphEducationUser) SetDepartmentNil(b bool)`

 SetDepartmentNil sets the value for Department to be an explicit nil

### UnsetDepartment
`func (o *MicrosoftGraphEducationUser) UnsetDepartment()`

UnsetDepartment ensures that no value is present for Department, not even an explicit nil
### GetDisplayName

`func (o *MicrosoftGraphEducationUser) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphEducationUser) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *MicrosoftGraphEducationUser) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *MicrosoftGraphEducationUser) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *MicrosoftGraphEducationUser) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *MicrosoftGraphEducationUser) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetExternalSource

`func (o *MicrosoftGraphEducationUser) GetExternalSource() AnyOfmicrosoftGraphEducationExternalSource`

GetExternalSource returns the ExternalSource field if non-nil, zero value otherwise.

### GetExternalSourceOk

`func (o *MicrosoftGraphEducationUser) GetExternalSourceOk() (*AnyOfmicrosoftGraphEducationExternalSource, bool)`

GetExternalSourceOk returns a tuple with the ExternalSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalSource

`func (o *MicrosoftGraphEducationUser) SetExternalSource(v AnyOfmicrosoftGraphEducationExternalSource)`

SetExternalSource sets ExternalSource field to given value.

### HasExternalSource

`func (o *MicrosoftGraphEducationUser) HasExternalSource() bool`

HasExternalSource returns a boolean if a field has been set.

### SetExternalSourceNil

`func (o *MicrosoftGraphEducationUser) SetExternalSourceNil(b bool)`

 SetExternalSourceNil sets the value for ExternalSource to be an explicit nil

### UnsetExternalSource
`func (o *MicrosoftGraphEducationUser) UnsetExternalSource()`

UnsetExternalSource ensures that no value is present for ExternalSource, not even an explicit nil
### GetExternalSourceDetail

`func (o *MicrosoftGraphEducationUser) GetExternalSourceDetail() string`

GetExternalSourceDetail returns the ExternalSourceDetail field if non-nil, zero value otherwise.

### GetExternalSourceDetailOk

`func (o *MicrosoftGraphEducationUser) GetExternalSourceDetailOk() (*string, bool)`

GetExternalSourceDetailOk returns a tuple with the ExternalSourceDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalSourceDetail

`func (o *MicrosoftGraphEducationUser) SetExternalSourceDetail(v string)`

SetExternalSourceDetail sets ExternalSourceDetail field to given value.

### HasExternalSourceDetail

`func (o *MicrosoftGraphEducationUser) HasExternalSourceDetail() bool`

HasExternalSourceDetail returns a boolean if a field has been set.

### SetExternalSourceDetailNil

`func (o *MicrosoftGraphEducationUser) SetExternalSourceDetailNil(b bool)`

 SetExternalSourceDetailNil sets the value for ExternalSourceDetail to be an explicit nil

### UnsetExternalSourceDetail
`func (o *MicrosoftGraphEducationUser) UnsetExternalSourceDetail()`

UnsetExternalSourceDetail ensures that no value is present for ExternalSourceDetail, not even an explicit nil
### GetGivenName

`func (o *MicrosoftGraphEducationUser) GetGivenName() string`

GetGivenName returns the GivenName field if non-nil, zero value otherwise.

### GetGivenNameOk

`func (o *MicrosoftGraphEducationUser) GetGivenNameOk() (*string, bool)`

GetGivenNameOk returns a tuple with the GivenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGivenName

`func (o *MicrosoftGraphEducationUser) SetGivenName(v string)`

SetGivenName sets GivenName field to given value.

### HasGivenName

`func (o *MicrosoftGraphEducationUser) HasGivenName() bool`

HasGivenName returns a boolean if a field has been set.

### SetGivenNameNil

`func (o *MicrosoftGraphEducationUser) SetGivenNameNil(b bool)`

 SetGivenNameNil sets the value for GivenName to be an explicit nil

### UnsetGivenName
`func (o *MicrosoftGraphEducationUser) UnsetGivenName()`

UnsetGivenName ensures that no value is present for GivenName, not even an explicit nil
### GetMail

`func (o *MicrosoftGraphEducationUser) GetMail() string`

GetMail returns the Mail field if non-nil, zero value otherwise.

### GetMailOk

`func (o *MicrosoftGraphEducationUser) GetMailOk() (*string, bool)`

GetMailOk returns a tuple with the Mail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMail

`func (o *MicrosoftGraphEducationUser) SetMail(v string)`

SetMail sets Mail field to given value.

### HasMail

`func (o *MicrosoftGraphEducationUser) HasMail() bool`

HasMail returns a boolean if a field has been set.

### SetMailNil

`func (o *MicrosoftGraphEducationUser) SetMailNil(b bool)`

 SetMailNil sets the value for Mail to be an explicit nil

### UnsetMail
`func (o *MicrosoftGraphEducationUser) UnsetMail()`

UnsetMail ensures that no value is present for Mail, not even an explicit nil
### GetMailingAddress

`func (o *MicrosoftGraphEducationUser) GetMailingAddress() AnyOfmicrosoftGraphPhysicalAddress`

GetMailingAddress returns the MailingAddress field if non-nil, zero value otherwise.

### GetMailingAddressOk

`func (o *MicrosoftGraphEducationUser) GetMailingAddressOk() (*AnyOfmicrosoftGraphPhysicalAddress, bool)`

GetMailingAddressOk returns a tuple with the MailingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailingAddress

`func (o *MicrosoftGraphEducationUser) SetMailingAddress(v AnyOfmicrosoftGraphPhysicalAddress)`

SetMailingAddress sets MailingAddress field to given value.

### HasMailingAddress

`func (o *MicrosoftGraphEducationUser) HasMailingAddress() bool`

HasMailingAddress returns a boolean if a field has been set.

### SetMailingAddressNil

`func (o *MicrosoftGraphEducationUser) SetMailingAddressNil(b bool)`

 SetMailingAddressNil sets the value for MailingAddress to be an explicit nil

### UnsetMailingAddress
`func (o *MicrosoftGraphEducationUser) UnsetMailingAddress()`

UnsetMailingAddress ensures that no value is present for MailingAddress, not even an explicit nil
### GetMailNickname

`func (o *MicrosoftGraphEducationUser) GetMailNickname() string`

GetMailNickname returns the MailNickname field if non-nil, zero value otherwise.

### GetMailNicknameOk

`func (o *MicrosoftGraphEducationUser) GetMailNicknameOk() (*string, bool)`

GetMailNicknameOk returns a tuple with the MailNickname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailNickname

`func (o *MicrosoftGraphEducationUser) SetMailNickname(v string)`

SetMailNickname sets MailNickname field to given value.

### HasMailNickname

`func (o *MicrosoftGraphEducationUser) HasMailNickname() bool`

HasMailNickname returns a boolean if a field has been set.

### SetMailNicknameNil

`func (o *MicrosoftGraphEducationUser) SetMailNicknameNil(b bool)`

 SetMailNicknameNil sets the value for MailNickname to be an explicit nil

### UnsetMailNickname
`func (o *MicrosoftGraphEducationUser) UnsetMailNickname()`

UnsetMailNickname ensures that no value is present for MailNickname, not even an explicit nil
### GetMiddleName

`func (o *MicrosoftGraphEducationUser) GetMiddleName() string`

GetMiddleName returns the MiddleName field if non-nil, zero value otherwise.

### GetMiddleNameOk

`func (o *MicrosoftGraphEducationUser) GetMiddleNameOk() (*string, bool)`

GetMiddleNameOk returns a tuple with the MiddleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiddleName

`func (o *MicrosoftGraphEducationUser) SetMiddleName(v string)`

SetMiddleName sets MiddleName field to given value.

### HasMiddleName

`func (o *MicrosoftGraphEducationUser) HasMiddleName() bool`

HasMiddleName returns a boolean if a field has been set.

### SetMiddleNameNil

`func (o *MicrosoftGraphEducationUser) SetMiddleNameNil(b bool)`

 SetMiddleNameNil sets the value for MiddleName to be an explicit nil

### UnsetMiddleName
`func (o *MicrosoftGraphEducationUser) UnsetMiddleName()`

UnsetMiddleName ensures that no value is present for MiddleName, not even an explicit nil
### GetMobilePhone

`func (o *MicrosoftGraphEducationUser) GetMobilePhone() string`

GetMobilePhone returns the MobilePhone field if non-nil, zero value otherwise.

### GetMobilePhoneOk

`func (o *MicrosoftGraphEducationUser) GetMobilePhoneOk() (*string, bool)`

GetMobilePhoneOk returns a tuple with the MobilePhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobilePhone

`func (o *MicrosoftGraphEducationUser) SetMobilePhone(v string)`

SetMobilePhone sets MobilePhone field to given value.

### HasMobilePhone

`func (o *MicrosoftGraphEducationUser) HasMobilePhone() bool`

HasMobilePhone returns a boolean if a field has been set.

### SetMobilePhoneNil

`func (o *MicrosoftGraphEducationUser) SetMobilePhoneNil(b bool)`

 SetMobilePhoneNil sets the value for MobilePhone to be an explicit nil

### UnsetMobilePhone
`func (o *MicrosoftGraphEducationUser) UnsetMobilePhone()`

UnsetMobilePhone ensures that no value is present for MobilePhone, not even an explicit nil
### GetOfficeLocation

`func (o *MicrosoftGraphEducationUser) GetOfficeLocation() string`

GetOfficeLocation returns the OfficeLocation field if non-nil, zero value otherwise.

### GetOfficeLocationOk

`func (o *MicrosoftGraphEducationUser) GetOfficeLocationOk() (*string, bool)`

GetOfficeLocationOk returns a tuple with the OfficeLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfficeLocation

`func (o *MicrosoftGraphEducationUser) SetOfficeLocation(v string)`

SetOfficeLocation sets OfficeLocation field to given value.

### HasOfficeLocation

`func (o *MicrosoftGraphEducationUser) HasOfficeLocation() bool`

HasOfficeLocation returns a boolean if a field has been set.

### SetOfficeLocationNil

`func (o *MicrosoftGraphEducationUser) SetOfficeLocationNil(b bool)`

 SetOfficeLocationNil sets the value for OfficeLocation to be an explicit nil

### UnsetOfficeLocation
`func (o *MicrosoftGraphEducationUser) UnsetOfficeLocation()`

UnsetOfficeLocation ensures that no value is present for OfficeLocation, not even an explicit nil
### GetOnPremisesInfo

`func (o *MicrosoftGraphEducationUser) GetOnPremisesInfo() AnyOfmicrosoftGraphEducationOnPremisesInfo`

GetOnPremisesInfo returns the OnPremisesInfo field if non-nil, zero value otherwise.

### GetOnPremisesInfoOk

`func (o *MicrosoftGraphEducationUser) GetOnPremisesInfoOk() (*AnyOfmicrosoftGraphEducationOnPremisesInfo, bool)`

GetOnPremisesInfoOk returns a tuple with the OnPremisesInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnPremisesInfo

`func (o *MicrosoftGraphEducationUser) SetOnPremisesInfo(v AnyOfmicrosoftGraphEducationOnPremisesInfo)`

SetOnPremisesInfo sets OnPremisesInfo field to given value.

### HasOnPremisesInfo

`func (o *MicrosoftGraphEducationUser) HasOnPremisesInfo() bool`

HasOnPremisesInfo returns a boolean if a field has been set.

### SetOnPremisesInfoNil

`func (o *MicrosoftGraphEducationUser) SetOnPremisesInfoNil(b bool)`

 SetOnPremisesInfoNil sets the value for OnPremisesInfo to be an explicit nil

### UnsetOnPremisesInfo
`func (o *MicrosoftGraphEducationUser) UnsetOnPremisesInfo()`

UnsetOnPremisesInfo ensures that no value is present for OnPremisesInfo, not even an explicit nil
### GetPasswordPolicies

`func (o *MicrosoftGraphEducationUser) GetPasswordPolicies() string`

GetPasswordPolicies returns the PasswordPolicies field if non-nil, zero value otherwise.

### GetPasswordPoliciesOk

`func (o *MicrosoftGraphEducationUser) GetPasswordPoliciesOk() (*string, bool)`

GetPasswordPoliciesOk returns a tuple with the PasswordPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordPolicies

`func (o *MicrosoftGraphEducationUser) SetPasswordPolicies(v string)`

SetPasswordPolicies sets PasswordPolicies field to given value.

### HasPasswordPolicies

`func (o *MicrosoftGraphEducationUser) HasPasswordPolicies() bool`

HasPasswordPolicies returns a boolean if a field has been set.

### SetPasswordPoliciesNil

`func (o *MicrosoftGraphEducationUser) SetPasswordPoliciesNil(b bool)`

 SetPasswordPoliciesNil sets the value for PasswordPolicies to be an explicit nil

### UnsetPasswordPolicies
`func (o *MicrosoftGraphEducationUser) UnsetPasswordPolicies()`

UnsetPasswordPolicies ensures that no value is present for PasswordPolicies, not even an explicit nil
### GetPasswordProfile

`func (o *MicrosoftGraphEducationUser) GetPasswordProfile() AnyOfmicrosoftGraphPasswordProfile`

GetPasswordProfile returns the PasswordProfile field if non-nil, zero value otherwise.

### GetPasswordProfileOk

`func (o *MicrosoftGraphEducationUser) GetPasswordProfileOk() (*AnyOfmicrosoftGraphPasswordProfile, bool)`

GetPasswordProfileOk returns a tuple with the PasswordProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordProfile

`func (o *MicrosoftGraphEducationUser) SetPasswordProfile(v AnyOfmicrosoftGraphPasswordProfile)`

SetPasswordProfile sets PasswordProfile field to given value.

### HasPasswordProfile

`func (o *MicrosoftGraphEducationUser) HasPasswordProfile() bool`

HasPasswordProfile returns a boolean if a field has been set.

### SetPasswordProfileNil

`func (o *MicrosoftGraphEducationUser) SetPasswordProfileNil(b bool)`

 SetPasswordProfileNil sets the value for PasswordProfile to be an explicit nil

### UnsetPasswordProfile
`func (o *MicrosoftGraphEducationUser) UnsetPasswordProfile()`

UnsetPasswordProfile ensures that no value is present for PasswordProfile, not even an explicit nil
### GetPreferredLanguage

`func (o *MicrosoftGraphEducationUser) GetPreferredLanguage() string`

GetPreferredLanguage returns the PreferredLanguage field if non-nil, zero value otherwise.

### GetPreferredLanguageOk

`func (o *MicrosoftGraphEducationUser) GetPreferredLanguageOk() (*string, bool)`

GetPreferredLanguageOk returns a tuple with the PreferredLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredLanguage

`func (o *MicrosoftGraphEducationUser) SetPreferredLanguage(v string)`

SetPreferredLanguage sets PreferredLanguage field to given value.

### HasPreferredLanguage

`func (o *MicrosoftGraphEducationUser) HasPreferredLanguage() bool`

HasPreferredLanguage returns a boolean if a field has been set.

### SetPreferredLanguageNil

`func (o *MicrosoftGraphEducationUser) SetPreferredLanguageNil(b bool)`

 SetPreferredLanguageNil sets the value for PreferredLanguage to be an explicit nil

### UnsetPreferredLanguage
`func (o *MicrosoftGraphEducationUser) UnsetPreferredLanguage()`

UnsetPreferredLanguage ensures that no value is present for PreferredLanguage, not even an explicit nil
### GetPrimaryRole

`func (o *MicrosoftGraphEducationUser) GetPrimaryRole() AnyOfmicrosoftGraphEducationUserRole`

GetPrimaryRole returns the PrimaryRole field if non-nil, zero value otherwise.

### GetPrimaryRoleOk

`func (o *MicrosoftGraphEducationUser) GetPrimaryRoleOk() (*AnyOfmicrosoftGraphEducationUserRole, bool)`

GetPrimaryRoleOk returns a tuple with the PrimaryRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryRole

`func (o *MicrosoftGraphEducationUser) SetPrimaryRole(v AnyOfmicrosoftGraphEducationUserRole)`

SetPrimaryRole sets PrimaryRole field to given value.

### HasPrimaryRole

`func (o *MicrosoftGraphEducationUser) HasPrimaryRole() bool`

HasPrimaryRole returns a boolean if a field has been set.

### SetPrimaryRoleNil

`func (o *MicrosoftGraphEducationUser) SetPrimaryRoleNil(b bool)`

 SetPrimaryRoleNil sets the value for PrimaryRole to be an explicit nil

### UnsetPrimaryRole
`func (o *MicrosoftGraphEducationUser) UnsetPrimaryRole()`

UnsetPrimaryRole ensures that no value is present for PrimaryRole, not even an explicit nil
### GetProvisionedPlans

`func (o *MicrosoftGraphEducationUser) GetProvisionedPlans() []MicrosoftGraphProvisionedPlan`

GetProvisionedPlans returns the ProvisionedPlans field if non-nil, zero value otherwise.

### GetProvisionedPlansOk

`func (o *MicrosoftGraphEducationUser) GetProvisionedPlansOk() (*[]MicrosoftGraphProvisionedPlan, bool)`

GetProvisionedPlansOk returns a tuple with the ProvisionedPlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionedPlans

`func (o *MicrosoftGraphEducationUser) SetProvisionedPlans(v []MicrosoftGraphProvisionedPlan)`

SetProvisionedPlans sets ProvisionedPlans field to given value.

### HasProvisionedPlans

`func (o *MicrosoftGraphEducationUser) HasProvisionedPlans() bool`

HasProvisionedPlans returns a boolean if a field has been set.

### GetRefreshTokensValidFromDateTime

`func (o *MicrosoftGraphEducationUser) GetRefreshTokensValidFromDateTime() time.Time`

GetRefreshTokensValidFromDateTime returns the RefreshTokensValidFromDateTime field if non-nil, zero value otherwise.

### GetRefreshTokensValidFromDateTimeOk

`func (o *MicrosoftGraphEducationUser) GetRefreshTokensValidFromDateTimeOk() (*time.Time, bool)`

GetRefreshTokensValidFromDateTimeOk returns a tuple with the RefreshTokensValidFromDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshTokensValidFromDateTime

`func (o *MicrosoftGraphEducationUser) SetRefreshTokensValidFromDateTime(v time.Time)`

SetRefreshTokensValidFromDateTime sets RefreshTokensValidFromDateTime field to given value.

### HasRefreshTokensValidFromDateTime

`func (o *MicrosoftGraphEducationUser) HasRefreshTokensValidFromDateTime() bool`

HasRefreshTokensValidFromDateTime returns a boolean if a field has been set.

### SetRefreshTokensValidFromDateTimeNil

`func (o *MicrosoftGraphEducationUser) SetRefreshTokensValidFromDateTimeNil(b bool)`

 SetRefreshTokensValidFromDateTimeNil sets the value for RefreshTokensValidFromDateTime to be an explicit nil

### UnsetRefreshTokensValidFromDateTime
`func (o *MicrosoftGraphEducationUser) UnsetRefreshTokensValidFromDateTime()`

UnsetRefreshTokensValidFromDateTime ensures that no value is present for RefreshTokensValidFromDateTime, not even an explicit nil
### GetResidenceAddress

`func (o *MicrosoftGraphEducationUser) GetResidenceAddress() AnyOfmicrosoftGraphPhysicalAddress`

GetResidenceAddress returns the ResidenceAddress field if non-nil, zero value otherwise.

### GetResidenceAddressOk

`func (o *MicrosoftGraphEducationUser) GetResidenceAddressOk() (*AnyOfmicrosoftGraphPhysicalAddress, bool)`

GetResidenceAddressOk returns a tuple with the ResidenceAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResidenceAddress

`func (o *MicrosoftGraphEducationUser) SetResidenceAddress(v AnyOfmicrosoftGraphPhysicalAddress)`

SetResidenceAddress sets ResidenceAddress field to given value.

### HasResidenceAddress

`func (o *MicrosoftGraphEducationUser) HasResidenceAddress() bool`

HasResidenceAddress returns a boolean if a field has been set.

### SetResidenceAddressNil

`func (o *MicrosoftGraphEducationUser) SetResidenceAddressNil(b bool)`

 SetResidenceAddressNil sets the value for ResidenceAddress to be an explicit nil

### UnsetResidenceAddress
`func (o *MicrosoftGraphEducationUser) UnsetResidenceAddress()`

UnsetResidenceAddress ensures that no value is present for ResidenceAddress, not even an explicit nil
### GetShowInAddressList

`func (o *MicrosoftGraphEducationUser) GetShowInAddressList() bool`

GetShowInAddressList returns the ShowInAddressList field if non-nil, zero value otherwise.

### GetShowInAddressListOk

`func (o *MicrosoftGraphEducationUser) GetShowInAddressListOk() (*bool, bool)`

GetShowInAddressListOk returns a tuple with the ShowInAddressList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowInAddressList

`func (o *MicrosoftGraphEducationUser) SetShowInAddressList(v bool)`

SetShowInAddressList sets ShowInAddressList field to given value.

### HasShowInAddressList

`func (o *MicrosoftGraphEducationUser) HasShowInAddressList() bool`

HasShowInAddressList returns a boolean if a field has been set.

### SetShowInAddressListNil

`func (o *MicrosoftGraphEducationUser) SetShowInAddressListNil(b bool)`

 SetShowInAddressListNil sets the value for ShowInAddressList to be an explicit nil

### UnsetShowInAddressList
`func (o *MicrosoftGraphEducationUser) UnsetShowInAddressList()`

UnsetShowInAddressList ensures that no value is present for ShowInAddressList, not even an explicit nil
### GetStudent

`func (o *MicrosoftGraphEducationUser) GetStudent() AnyOfmicrosoftGraphEducationStudent`

GetStudent returns the Student field if non-nil, zero value otherwise.

### GetStudentOk

`func (o *MicrosoftGraphEducationUser) GetStudentOk() (*AnyOfmicrosoftGraphEducationStudent, bool)`

GetStudentOk returns a tuple with the Student field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStudent

`func (o *MicrosoftGraphEducationUser) SetStudent(v AnyOfmicrosoftGraphEducationStudent)`

SetStudent sets Student field to given value.

### HasStudent

`func (o *MicrosoftGraphEducationUser) HasStudent() bool`

HasStudent returns a boolean if a field has been set.

### SetStudentNil

`func (o *MicrosoftGraphEducationUser) SetStudentNil(b bool)`

 SetStudentNil sets the value for Student to be an explicit nil

### UnsetStudent
`func (o *MicrosoftGraphEducationUser) UnsetStudent()`

UnsetStudent ensures that no value is present for Student, not even an explicit nil
### GetSurname

`func (o *MicrosoftGraphEducationUser) GetSurname() string`

GetSurname returns the Surname field if non-nil, zero value otherwise.

### GetSurnameOk

`func (o *MicrosoftGraphEducationUser) GetSurnameOk() (*string, bool)`

GetSurnameOk returns a tuple with the Surname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurname

`func (o *MicrosoftGraphEducationUser) SetSurname(v string)`

SetSurname sets Surname field to given value.

### HasSurname

`func (o *MicrosoftGraphEducationUser) HasSurname() bool`

HasSurname returns a boolean if a field has been set.

### SetSurnameNil

`func (o *MicrosoftGraphEducationUser) SetSurnameNil(b bool)`

 SetSurnameNil sets the value for Surname to be an explicit nil

### UnsetSurname
`func (o *MicrosoftGraphEducationUser) UnsetSurname()`

UnsetSurname ensures that no value is present for Surname, not even an explicit nil
### GetTeacher

`func (o *MicrosoftGraphEducationUser) GetTeacher() AnyOfmicrosoftGraphEducationTeacher`

GetTeacher returns the Teacher field if non-nil, zero value otherwise.

### GetTeacherOk

`func (o *MicrosoftGraphEducationUser) GetTeacherOk() (*AnyOfmicrosoftGraphEducationTeacher, bool)`

GetTeacherOk returns a tuple with the Teacher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeacher

`func (o *MicrosoftGraphEducationUser) SetTeacher(v AnyOfmicrosoftGraphEducationTeacher)`

SetTeacher sets Teacher field to given value.

### HasTeacher

`func (o *MicrosoftGraphEducationUser) HasTeacher() bool`

HasTeacher returns a boolean if a field has been set.

### SetTeacherNil

`func (o *MicrosoftGraphEducationUser) SetTeacherNil(b bool)`

 SetTeacherNil sets the value for Teacher to be an explicit nil

### UnsetTeacher
`func (o *MicrosoftGraphEducationUser) UnsetTeacher()`

UnsetTeacher ensures that no value is present for Teacher, not even an explicit nil
### GetUsageLocation

`func (o *MicrosoftGraphEducationUser) GetUsageLocation() string`

GetUsageLocation returns the UsageLocation field if non-nil, zero value otherwise.

### GetUsageLocationOk

`func (o *MicrosoftGraphEducationUser) GetUsageLocationOk() (*string, bool)`

GetUsageLocationOk returns a tuple with the UsageLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageLocation

`func (o *MicrosoftGraphEducationUser) SetUsageLocation(v string)`

SetUsageLocation sets UsageLocation field to given value.

### HasUsageLocation

`func (o *MicrosoftGraphEducationUser) HasUsageLocation() bool`

HasUsageLocation returns a boolean if a field has been set.

### SetUsageLocationNil

`func (o *MicrosoftGraphEducationUser) SetUsageLocationNil(b bool)`

 SetUsageLocationNil sets the value for UsageLocation to be an explicit nil

### UnsetUsageLocation
`func (o *MicrosoftGraphEducationUser) UnsetUsageLocation()`

UnsetUsageLocation ensures that no value is present for UsageLocation, not even an explicit nil
### GetUserPrincipalName

`func (o *MicrosoftGraphEducationUser) GetUserPrincipalName() string`

GetUserPrincipalName returns the UserPrincipalName field if non-nil, zero value otherwise.

### GetUserPrincipalNameOk

`func (o *MicrosoftGraphEducationUser) GetUserPrincipalNameOk() (*string, bool)`

GetUserPrincipalNameOk returns a tuple with the UserPrincipalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPrincipalName

`func (o *MicrosoftGraphEducationUser) SetUserPrincipalName(v string)`

SetUserPrincipalName sets UserPrincipalName field to given value.

### HasUserPrincipalName

`func (o *MicrosoftGraphEducationUser) HasUserPrincipalName() bool`

HasUserPrincipalName returns a boolean if a field has been set.

### SetUserPrincipalNameNil

`func (o *MicrosoftGraphEducationUser) SetUserPrincipalNameNil(b bool)`

 SetUserPrincipalNameNil sets the value for UserPrincipalName to be an explicit nil

### UnsetUserPrincipalName
`func (o *MicrosoftGraphEducationUser) UnsetUserPrincipalName()`

UnsetUserPrincipalName ensures that no value is present for UserPrincipalName, not even an explicit nil
### GetUserType

`func (o *MicrosoftGraphEducationUser) GetUserType() string`

GetUserType returns the UserType field if non-nil, zero value otherwise.

### GetUserTypeOk

`func (o *MicrosoftGraphEducationUser) GetUserTypeOk() (*string, bool)`

GetUserTypeOk returns a tuple with the UserType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserType

`func (o *MicrosoftGraphEducationUser) SetUserType(v string)`

SetUserType sets UserType field to given value.

### HasUserType

`func (o *MicrosoftGraphEducationUser) HasUserType() bool`

HasUserType returns a boolean if a field has been set.

### SetUserTypeNil

`func (o *MicrosoftGraphEducationUser) SetUserTypeNil(b bool)`

 SetUserTypeNil sets the value for UserType to be an explicit nil

### UnsetUserType
`func (o *MicrosoftGraphEducationUser) UnsetUserType()`

UnsetUserType ensures that no value is present for UserType, not even an explicit nil
### GetRubrics

`func (o *MicrosoftGraphEducationUser) GetRubrics() []MicrosoftGraphEducationRubric`

GetRubrics returns the Rubrics field if non-nil, zero value otherwise.

### GetRubricsOk

`func (o *MicrosoftGraphEducationUser) GetRubricsOk() (*[]MicrosoftGraphEducationRubric, bool)`

GetRubricsOk returns a tuple with the Rubrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRubrics

`func (o *MicrosoftGraphEducationUser) SetRubrics(v []MicrosoftGraphEducationRubric)`

SetRubrics sets Rubrics field to given value.

### HasRubrics

`func (o *MicrosoftGraphEducationUser) HasRubrics() bool`

HasRubrics returns a boolean if a field has been set.

### GetClasses

`func (o *MicrosoftGraphEducationUser) GetClasses() []MicrosoftGraphEducationClass`

GetClasses returns the Classes field if non-nil, zero value otherwise.

### GetClassesOk

`func (o *MicrosoftGraphEducationUser) GetClassesOk() (*[]MicrosoftGraphEducationClass, bool)`

GetClassesOk returns a tuple with the Classes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClasses

`func (o *MicrosoftGraphEducationUser) SetClasses(v []MicrosoftGraphEducationClass)`

SetClasses sets Classes field to given value.

### HasClasses

`func (o *MicrosoftGraphEducationUser) HasClasses() bool`

HasClasses returns a boolean if a field has been set.

### GetSchools

`func (o *MicrosoftGraphEducationUser) GetSchools() []MicrosoftGraphEducationSchool`

GetSchools returns the Schools field if non-nil, zero value otherwise.

### GetSchoolsOk

`func (o *MicrosoftGraphEducationUser) GetSchoolsOk() (*[]MicrosoftGraphEducationSchool, bool)`

GetSchoolsOk returns a tuple with the Schools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchools

`func (o *MicrosoftGraphEducationUser) SetSchools(v []MicrosoftGraphEducationSchool)`

SetSchools sets Schools field to given value.

### HasSchools

`func (o *MicrosoftGraphEducationUser) HasSchools() bool`

HasSchools returns a boolean if a field has been set.

### GetTaughtClasses

`func (o *MicrosoftGraphEducationUser) GetTaughtClasses() []MicrosoftGraphEducationClass`

GetTaughtClasses returns the TaughtClasses field if non-nil, zero value otherwise.

### GetTaughtClassesOk

`func (o *MicrosoftGraphEducationUser) GetTaughtClassesOk() (*[]MicrosoftGraphEducationClass, bool)`

GetTaughtClassesOk returns a tuple with the TaughtClasses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaughtClasses

`func (o *MicrosoftGraphEducationUser) SetTaughtClasses(v []MicrosoftGraphEducationClass)`

SetTaughtClasses sets TaughtClasses field to given value.

### HasTaughtClasses

`func (o *MicrosoftGraphEducationUser) HasTaughtClasses() bool`

HasTaughtClasses returns a boolean if a field has been set.

### GetUser

`func (o *MicrosoftGraphEducationUser) GetUser() AnyOfmicrosoftGraphUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *MicrosoftGraphEducationUser) GetUserOk() (*AnyOfmicrosoftGraphUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *MicrosoftGraphEducationUser) SetUser(v AnyOfmicrosoftGraphUser)`

SetUser sets User field to given value.

### HasUser

`func (o *MicrosoftGraphEducationUser) HasUser() bool`

HasUser returns a boolean if a field has been set.

### SetUserNil

`func (o *MicrosoftGraphEducationUser) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *MicrosoftGraphEducationUser) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


