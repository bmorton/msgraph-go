# OrgContact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Addresses** | Pointer to [**[]AnyOfmicrosoftGraphPhysicalOfficeAddress**](AnyOfmicrosoftGraphPhysicalOfficeAddress.md) | Postal addresses for this organizational contact. For now a contact can only have one physical address. | [optional] 
**CompanyName** | Pointer to **NullableString** | Name of the company that this organizational contact belong to. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values). | [optional] 
**Department** | Pointer to **NullableString** | The name for the department in which the contact works. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values). | [optional] 
**DisplayName** | Pointer to **NullableString** | Display name for this organizational contact. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values), $search, and $orderBy. | [optional] 
**GivenName** | Pointer to **NullableString** | First name for this organizational contact. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values). | [optional] 
**JobTitle** | Pointer to **NullableString** | Job title for this organizational contact. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values). | [optional] 
**Mail** | Pointer to **NullableString** | The SMTP address for the contact, for example, &#39;jeff@contoso.onmicrosoft.com&#39;. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values). | [optional] 
**MailNickname** | Pointer to **NullableString** | Email alias (portion of email address pre-pending the @ symbol) for this organizational contact. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values). | [optional] 
**OnPremisesLastSyncDateTime** | Pointer to **NullableTime** | Date and time when this organizational contact was last synchronized from on-premises AD. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Supports $filter (eq, ne, not, ge, le, in). | [optional] 
**OnPremisesProvisioningErrors** | Pointer to [**[]AnyOfmicrosoftGraphOnPremisesProvisioningError**](AnyOfmicrosoftGraphOnPremisesProvisioningError.md) | List of any synchronization provisioning errors for this organizational contact. Supports $filter (eq, not). | [optional] 
**OnPremisesSyncEnabled** | Pointer to **NullableBool** | true if this object is synced from an on-premises directory; false if this object was originally synced from an on-premises directory but is no longer synced and now mastered in Exchange; null if this object has never been synced from an on-premises directory (default).  Supports $filter (eq, ne, not, in, and eq on null values). | [optional] 
**Phones** | Pointer to [**[]AnyOfmicrosoftGraphPhone**](AnyOfmicrosoftGraphPhone.md) | List of phones for this organizational contact. Phone types can be mobile, business, and businessFax. Only one of each type can ever be present in the collection. Supports $filter (eq, ne, not, in). | [optional] 
**ProxyAddresses** | Pointer to **[]string** | For example: &#39;SMTP: bob@contoso.com&#39;, &#39;smtp: bob@sales.contoso.com&#39;. The any operator is required for filter expressions on multi-valued properties. Supports $filter (eq, not, ge, le, startsWith). | [optional] 
**Surname** | Pointer to **NullableString** | Last name for this organizational contact. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values) | [optional] 
**DirectReports** | Pointer to [**[]MicrosoftGraphDirectoryObject**](MicrosoftGraphDirectoryObject.md) | The contact&#39;s direct reports. (The users and contacts that have their manager property set to this contact.) Read-only. Nullable. Supports $expand. | [optional] 
**Manager** | Pointer to [**AnyOfmicrosoftGraphDirectoryObject**](anyOf&lt;microsoft.graph.directoryObject&gt;.md) | The user or contact that is this contact&#39;s manager. Read-only. Supports $expand. | [optional] 
**MemberOf** | Pointer to [**[]MicrosoftGraphDirectoryObject**](MicrosoftGraphDirectoryObject.md) | Groups that this contact is a member of. Read-only. Nullable. Supports $expand. | [optional] 
**TransitiveMemberOf** | Pointer to [**[]MicrosoftGraphDirectoryObject**](MicrosoftGraphDirectoryObject.md) |  | [optional] 

## Methods

### NewOrgContact

`func NewOrgContact() *OrgContact`

NewOrgContact instantiates a new OrgContact object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgContactWithDefaults

`func NewOrgContactWithDefaults() *OrgContact`

NewOrgContactWithDefaults instantiates a new OrgContact object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddresses

`func (o *OrgContact) GetAddresses() []*AnyOfmicrosoftGraphPhysicalOfficeAddress`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *OrgContact) GetAddressesOk() (*[]*AnyOfmicrosoftGraphPhysicalOfficeAddress, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *OrgContact) SetAddresses(v []*AnyOfmicrosoftGraphPhysicalOfficeAddress)`

SetAddresses sets Addresses field to given value.

### HasAddresses

`func (o *OrgContact) HasAddresses() bool`

HasAddresses returns a boolean if a field has been set.

### GetCompanyName

`func (o *OrgContact) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *OrgContact) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *OrgContact) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *OrgContact) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### SetCompanyNameNil

`func (o *OrgContact) SetCompanyNameNil(b bool)`

 SetCompanyNameNil sets the value for CompanyName to be an explicit nil

### UnsetCompanyName
`func (o *OrgContact) UnsetCompanyName()`

UnsetCompanyName ensures that no value is present for CompanyName, not even an explicit nil
### GetDepartment

`func (o *OrgContact) GetDepartment() string`

GetDepartment returns the Department field if non-nil, zero value otherwise.

### GetDepartmentOk

`func (o *OrgContact) GetDepartmentOk() (*string, bool)`

GetDepartmentOk returns a tuple with the Department field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartment

`func (o *OrgContact) SetDepartment(v string)`

SetDepartment sets Department field to given value.

### HasDepartment

`func (o *OrgContact) HasDepartment() bool`

HasDepartment returns a boolean if a field has been set.

### SetDepartmentNil

`func (o *OrgContact) SetDepartmentNil(b bool)`

 SetDepartmentNil sets the value for Department to be an explicit nil

### UnsetDepartment
`func (o *OrgContact) UnsetDepartment()`

UnsetDepartment ensures that no value is present for Department, not even an explicit nil
### GetDisplayName

`func (o *OrgContact) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *OrgContact) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *OrgContact) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *OrgContact) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *OrgContact) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *OrgContact) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetGivenName

`func (o *OrgContact) GetGivenName() string`

GetGivenName returns the GivenName field if non-nil, zero value otherwise.

### GetGivenNameOk

`func (o *OrgContact) GetGivenNameOk() (*string, bool)`

GetGivenNameOk returns a tuple with the GivenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGivenName

`func (o *OrgContact) SetGivenName(v string)`

SetGivenName sets GivenName field to given value.

### HasGivenName

`func (o *OrgContact) HasGivenName() bool`

HasGivenName returns a boolean if a field has been set.

### SetGivenNameNil

`func (o *OrgContact) SetGivenNameNil(b bool)`

 SetGivenNameNil sets the value for GivenName to be an explicit nil

### UnsetGivenName
`func (o *OrgContact) UnsetGivenName()`

UnsetGivenName ensures that no value is present for GivenName, not even an explicit nil
### GetJobTitle

`func (o *OrgContact) GetJobTitle() string`

GetJobTitle returns the JobTitle field if non-nil, zero value otherwise.

### GetJobTitleOk

`func (o *OrgContact) GetJobTitleOk() (*string, bool)`

GetJobTitleOk returns a tuple with the JobTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobTitle

`func (o *OrgContact) SetJobTitle(v string)`

SetJobTitle sets JobTitle field to given value.

### HasJobTitle

`func (o *OrgContact) HasJobTitle() bool`

HasJobTitle returns a boolean if a field has been set.

### SetJobTitleNil

`func (o *OrgContact) SetJobTitleNil(b bool)`

 SetJobTitleNil sets the value for JobTitle to be an explicit nil

### UnsetJobTitle
`func (o *OrgContact) UnsetJobTitle()`

UnsetJobTitle ensures that no value is present for JobTitle, not even an explicit nil
### GetMail

`func (o *OrgContact) GetMail() string`

GetMail returns the Mail field if non-nil, zero value otherwise.

### GetMailOk

`func (o *OrgContact) GetMailOk() (*string, bool)`

GetMailOk returns a tuple with the Mail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMail

`func (o *OrgContact) SetMail(v string)`

SetMail sets Mail field to given value.

### HasMail

`func (o *OrgContact) HasMail() bool`

HasMail returns a boolean if a field has been set.

### SetMailNil

`func (o *OrgContact) SetMailNil(b bool)`

 SetMailNil sets the value for Mail to be an explicit nil

### UnsetMail
`func (o *OrgContact) UnsetMail()`

UnsetMail ensures that no value is present for Mail, not even an explicit nil
### GetMailNickname

`func (o *OrgContact) GetMailNickname() string`

GetMailNickname returns the MailNickname field if non-nil, zero value otherwise.

### GetMailNicknameOk

`func (o *OrgContact) GetMailNicknameOk() (*string, bool)`

GetMailNicknameOk returns a tuple with the MailNickname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailNickname

`func (o *OrgContact) SetMailNickname(v string)`

SetMailNickname sets MailNickname field to given value.

### HasMailNickname

`func (o *OrgContact) HasMailNickname() bool`

HasMailNickname returns a boolean if a field has been set.

### SetMailNicknameNil

`func (o *OrgContact) SetMailNicknameNil(b bool)`

 SetMailNicknameNil sets the value for MailNickname to be an explicit nil

### UnsetMailNickname
`func (o *OrgContact) UnsetMailNickname()`

UnsetMailNickname ensures that no value is present for MailNickname, not even an explicit nil
### GetOnPremisesLastSyncDateTime

`func (o *OrgContact) GetOnPremisesLastSyncDateTime() time.Time`

GetOnPremisesLastSyncDateTime returns the OnPremisesLastSyncDateTime field if non-nil, zero value otherwise.

### GetOnPremisesLastSyncDateTimeOk

`func (o *OrgContact) GetOnPremisesLastSyncDateTimeOk() (*time.Time, bool)`

GetOnPremisesLastSyncDateTimeOk returns a tuple with the OnPremisesLastSyncDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnPremisesLastSyncDateTime

`func (o *OrgContact) SetOnPremisesLastSyncDateTime(v time.Time)`

SetOnPremisesLastSyncDateTime sets OnPremisesLastSyncDateTime field to given value.

### HasOnPremisesLastSyncDateTime

`func (o *OrgContact) HasOnPremisesLastSyncDateTime() bool`

HasOnPremisesLastSyncDateTime returns a boolean if a field has been set.

### SetOnPremisesLastSyncDateTimeNil

`func (o *OrgContact) SetOnPremisesLastSyncDateTimeNil(b bool)`

 SetOnPremisesLastSyncDateTimeNil sets the value for OnPremisesLastSyncDateTime to be an explicit nil

### UnsetOnPremisesLastSyncDateTime
`func (o *OrgContact) UnsetOnPremisesLastSyncDateTime()`

UnsetOnPremisesLastSyncDateTime ensures that no value is present for OnPremisesLastSyncDateTime, not even an explicit nil
### GetOnPremisesProvisioningErrors

`func (o *OrgContact) GetOnPremisesProvisioningErrors() []*AnyOfmicrosoftGraphOnPremisesProvisioningError`

GetOnPremisesProvisioningErrors returns the OnPremisesProvisioningErrors field if non-nil, zero value otherwise.

### GetOnPremisesProvisioningErrorsOk

`func (o *OrgContact) GetOnPremisesProvisioningErrorsOk() (*[]*AnyOfmicrosoftGraphOnPremisesProvisioningError, bool)`

GetOnPremisesProvisioningErrorsOk returns a tuple with the OnPremisesProvisioningErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnPremisesProvisioningErrors

`func (o *OrgContact) SetOnPremisesProvisioningErrors(v []*AnyOfmicrosoftGraphOnPremisesProvisioningError)`

SetOnPremisesProvisioningErrors sets OnPremisesProvisioningErrors field to given value.

### HasOnPremisesProvisioningErrors

`func (o *OrgContact) HasOnPremisesProvisioningErrors() bool`

HasOnPremisesProvisioningErrors returns a boolean if a field has been set.

### GetOnPremisesSyncEnabled

`func (o *OrgContact) GetOnPremisesSyncEnabled() bool`

GetOnPremisesSyncEnabled returns the OnPremisesSyncEnabled field if non-nil, zero value otherwise.

### GetOnPremisesSyncEnabledOk

`func (o *OrgContact) GetOnPremisesSyncEnabledOk() (*bool, bool)`

GetOnPremisesSyncEnabledOk returns a tuple with the OnPremisesSyncEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnPremisesSyncEnabled

`func (o *OrgContact) SetOnPremisesSyncEnabled(v bool)`

SetOnPremisesSyncEnabled sets OnPremisesSyncEnabled field to given value.

### HasOnPremisesSyncEnabled

`func (o *OrgContact) HasOnPremisesSyncEnabled() bool`

HasOnPremisesSyncEnabled returns a boolean if a field has been set.

### SetOnPremisesSyncEnabledNil

`func (o *OrgContact) SetOnPremisesSyncEnabledNil(b bool)`

 SetOnPremisesSyncEnabledNil sets the value for OnPremisesSyncEnabled to be an explicit nil

### UnsetOnPremisesSyncEnabled
`func (o *OrgContact) UnsetOnPremisesSyncEnabled()`

UnsetOnPremisesSyncEnabled ensures that no value is present for OnPremisesSyncEnabled, not even an explicit nil
### GetPhones

`func (o *OrgContact) GetPhones() []*AnyOfmicrosoftGraphPhone`

GetPhones returns the Phones field if non-nil, zero value otherwise.

### GetPhonesOk

`func (o *OrgContact) GetPhonesOk() (*[]*AnyOfmicrosoftGraphPhone, bool)`

GetPhonesOk returns a tuple with the Phones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhones

`func (o *OrgContact) SetPhones(v []*AnyOfmicrosoftGraphPhone)`

SetPhones sets Phones field to given value.

### HasPhones

`func (o *OrgContact) HasPhones() bool`

HasPhones returns a boolean if a field has been set.

### GetProxyAddresses

`func (o *OrgContact) GetProxyAddresses() []string`

GetProxyAddresses returns the ProxyAddresses field if non-nil, zero value otherwise.

### GetProxyAddressesOk

`func (o *OrgContact) GetProxyAddressesOk() (*[]string, bool)`

GetProxyAddressesOk returns a tuple with the ProxyAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyAddresses

`func (o *OrgContact) SetProxyAddresses(v []string)`

SetProxyAddresses sets ProxyAddresses field to given value.

### HasProxyAddresses

`func (o *OrgContact) HasProxyAddresses() bool`

HasProxyAddresses returns a boolean if a field has been set.

### GetSurname

`func (o *OrgContact) GetSurname() string`

GetSurname returns the Surname field if non-nil, zero value otherwise.

### GetSurnameOk

`func (o *OrgContact) GetSurnameOk() (*string, bool)`

GetSurnameOk returns a tuple with the Surname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurname

`func (o *OrgContact) SetSurname(v string)`

SetSurname sets Surname field to given value.

### HasSurname

`func (o *OrgContact) HasSurname() bool`

HasSurname returns a boolean if a field has been set.

### SetSurnameNil

`func (o *OrgContact) SetSurnameNil(b bool)`

 SetSurnameNil sets the value for Surname to be an explicit nil

### UnsetSurname
`func (o *OrgContact) UnsetSurname()`

UnsetSurname ensures that no value is present for Surname, not even an explicit nil
### GetDirectReports

`func (o *OrgContact) GetDirectReports() []MicrosoftGraphDirectoryObject`

GetDirectReports returns the DirectReports field if non-nil, zero value otherwise.

### GetDirectReportsOk

`func (o *OrgContact) GetDirectReportsOk() (*[]MicrosoftGraphDirectoryObject, bool)`

GetDirectReportsOk returns a tuple with the DirectReports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectReports

`func (o *OrgContact) SetDirectReports(v []MicrosoftGraphDirectoryObject)`

SetDirectReports sets DirectReports field to given value.

### HasDirectReports

`func (o *OrgContact) HasDirectReports() bool`

HasDirectReports returns a boolean if a field has been set.

### GetManager

`func (o *OrgContact) GetManager() AnyOfmicrosoftGraphDirectoryObject`

GetManager returns the Manager field if non-nil, zero value otherwise.

### GetManagerOk

`func (o *OrgContact) GetManagerOk() (*AnyOfmicrosoftGraphDirectoryObject, bool)`

GetManagerOk returns a tuple with the Manager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManager

`func (o *OrgContact) SetManager(v AnyOfmicrosoftGraphDirectoryObject)`

SetManager sets Manager field to given value.

### HasManager

`func (o *OrgContact) HasManager() bool`

HasManager returns a boolean if a field has been set.

### SetManagerNil

`func (o *OrgContact) SetManagerNil(b bool)`

 SetManagerNil sets the value for Manager to be an explicit nil

### UnsetManager
`func (o *OrgContact) UnsetManager()`

UnsetManager ensures that no value is present for Manager, not even an explicit nil
### GetMemberOf

`func (o *OrgContact) GetMemberOf() []MicrosoftGraphDirectoryObject`

GetMemberOf returns the MemberOf field if non-nil, zero value otherwise.

### GetMemberOfOk

`func (o *OrgContact) GetMemberOfOk() (*[]MicrosoftGraphDirectoryObject, bool)`

GetMemberOfOk returns a tuple with the MemberOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberOf

`func (o *OrgContact) SetMemberOf(v []MicrosoftGraphDirectoryObject)`

SetMemberOf sets MemberOf field to given value.

### HasMemberOf

`func (o *OrgContact) HasMemberOf() bool`

HasMemberOf returns a boolean if a field has been set.

### GetTransitiveMemberOf

`func (o *OrgContact) GetTransitiveMemberOf() []MicrosoftGraphDirectoryObject`

GetTransitiveMemberOf returns the TransitiveMemberOf field if non-nil, zero value otherwise.

### GetTransitiveMemberOfOk

`func (o *OrgContact) GetTransitiveMemberOfOk() (*[]MicrosoftGraphDirectoryObject, bool)`

GetTransitiveMemberOfOk returns a tuple with the TransitiveMemberOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransitiveMemberOf

`func (o *OrgContact) SetTransitiveMemberOf(v []MicrosoftGraphDirectoryObject)`

SetTransitiveMemberOf sets TransitiveMemberOf field to given value.

### HasTransitiveMemberOf

`func (o *OrgContact) HasTransitiveMemberOf() bool`

HasTransitiveMemberOf returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


