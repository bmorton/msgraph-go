# Contact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssistantName** | Pointer to **NullableString** | The name of the contact&#39;s assistant. | [optional] 
**Birthday** | Pointer to **NullableTime** | The contact&#39;s birthday. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z | [optional] 
**BusinessAddress** | Pointer to [**AnyOfmicrosoftGraphPhysicalAddress**](anyOf&lt;microsoft.graph.physicalAddress&gt;.md) | The contact&#39;s business address. | [optional] 
**BusinessHomePage** | Pointer to **NullableString** | The business home page of the contact. | [optional] 
**BusinessPhones** | Pointer to **[]string** | The contact&#39;s business phone numbers. | [optional] 
**Children** | Pointer to **[]string** | The names of the contact&#39;s children. | [optional] 
**CompanyName** | Pointer to **NullableString** | The name of the contact&#39;s company. | [optional] 
**Department** | Pointer to **NullableString** | The contact&#39;s department. | [optional] 
**DisplayName** | Pointer to **NullableString** | The contact&#39;s display name. You can specify the display name in a create or update operation. Note that later updates to other properties may cause an automatically generated value to overwrite the displayName value you have specified. To preserve a pre-existing value, always include it as displayName in an update operation. | [optional] 
**EmailAddresses** | Pointer to [**[]AnyOfmicrosoftGraphEmailAddress**](AnyOfmicrosoftGraphEmailAddress.md) | The contact&#39;s email addresses. | [optional] 
**FileAs** | Pointer to **NullableString** | The name the contact is filed under. | [optional] 
**Generation** | Pointer to **NullableString** | The contact&#39;s generation. | [optional] 
**GivenName** | Pointer to **NullableString** | The contact&#39;s given name. | [optional] 
**HomeAddress** | Pointer to [**AnyOfmicrosoftGraphPhysicalAddress**](anyOf&lt;microsoft.graph.physicalAddress&gt;.md) | The contact&#39;s home address. | [optional] 
**HomePhones** | Pointer to **[]string** | The contact&#39;s home phone numbers. | [optional] 
**ImAddresses** | Pointer to **[]string** |  | [optional] 
**Initials** | Pointer to **NullableString** |  | [optional] 
**JobTitle** | Pointer to **NullableString** |  | [optional] 
**Manager** | Pointer to **NullableString** |  | [optional] 
**MiddleName** | Pointer to **NullableString** |  | [optional] 
**MobilePhone** | Pointer to **NullableString** |  | [optional] 
**NickName** | Pointer to **NullableString** |  | [optional] 
**OfficeLocation** | Pointer to **NullableString** |  | [optional] 
**OtherAddress** | Pointer to [**AnyOfmicrosoftGraphPhysicalAddress**](anyOf&lt;microsoft.graph.physicalAddress&gt;.md) |  | [optional] 
**ParentFolderId** | Pointer to **NullableString** |  | [optional] 
**PersonalNotes** | Pointer to **NullableString** |  | [optional] 
**Profession** | Pointer to **NullableString** |  | [optional] 
**SpouseName** | Pointer to **NullableString** |  | [optional] 
**Surname** | Pointer to **NullableString** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**YomiCompanyName** | Pointer to **NullableString** |  | [optional] 
**YomiGivenName** | Pointer to **NullableString** |  | [optional] 
**YomiSurname** | Pointer to **NullableString** |  | [optional] 
**Extensions** | Pointer to [**[]MicrosoftGraphExtension**](MicrosoftGraphExtension.md) | The collection of open extensions defined for the contact. Read-only. Nullable. | [optional] 
**MultiValueExtendedProperties** | Pointer to [**[]MicrosoftGraphMultiValueLegacyExtendedProperty**](MicrosoftGraphMultiValueLegacyExtendedProperty.md) | The collection of multi-value extended properties defined for the contact. Read-only. Nullable. | [optional] 
**Photo** | Pointer to [**AnyOfmicrosoftGraphProfilePhoto**](anyOf&lt;microsoft.graph.profilePhoto&gt;.md) | Optional contact picture. You can get or set a photo for a contact. | [optional] 
**SingleValueExtendedProperties** | Pointer to [**[]MicrosoftGraphSingleValueLegacyExtendedProperty**](MicrosoftGraphSingleValueLegacyExtendedProperty.md) | The collection of single-value extended properties defined for the contact. Read-only. Nullable. | [optional] 

## Methods

### NewContact

`func NewContact() *Contact`

NewContact instantiates a new Contact object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContactWithDefaults

`func NewContactWithDefaults() *Contact`

NewContactWithDefaults instantiates a new Contact object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssistantName

`func (o *Contact) GetAssistantName() string`

GetAssistantName returns the AssistantName field if non-nil, zero value otherwise.

### GetAssistantNameOk

`func (o *Contact) GetAssistantNameOk() (*string, bool)`

GetAssistantNameOk returns a tuple with the AssistantName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssistantName

`func (o *Contact) SetAssistantName(v string)`

SetAssistantName sets AssistantName field to given value.

### HasAssistantName

`func (o *Contact) HasAssistantName() bool`

HasAssistantName returns a boolean if a field has been set.

### SetAssistantNameNil

`func (o *Contact) SetAssistantNameNil(b bool)`

 SetAssistantNameNil sets the value for AssistantName to be an explicit nil

### UnsetAssistantName
`func (o *Contact) UnsetAssistantName()`

UnsetAssistantName ensures that no value is present for AssistantName, not even an explicit nil
### GetBirthday

`func (o *Contact) GetBirthday() time.Time`

GetBirthday returns the Birthday field if non-nil, zero value otherwise.

### GetBirthdayOk

`func (o *Contact) GetBirthdayOk() (*time.Time, bool)`

GetBirthdayOk returns a tuple with the Birthday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBirthday

`func (o *Contact) SetBirthday(v time.Time)`

SetBirthday sets Birthday field to given value.

### HasBirthday

`func (o *Contact) HasBirthday() bool`

HasBirthday returns a boolean if a field has been set.

### SetBirthdayNil

`func (o *Contact) SetBirthdayNil(b bool)`

 SetBirthdayNil sets the value for Birthday to be an explicit nil

### UnsetBirthday
`func (o *Contact) UnsetBirthday()`

UnsetBirthday ensures that no value is present for Birthday, not even an explicit nil
### GetBusinessAddress

`func (o *Contact) GetBusinessAddress() AnyOfmicrosoftGraphPhysicalAddress`

GetBusinessAddress returns the BusinessAddress field if non-nil, zero value otherwise.

### GetBusinessAddressOk

`func (o *Contact) GetBusinessAddressOk() (*AnyOfmicrosoftGraphPhysicalAddress, bool)`

GetBusinessAddressOk returns a tuple with the BusinessAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessAddress

`func (o *Contact) SetBusinessAddress(v AnyOfmicrosoftGraphPhysicalAddress)`

SetBusinessAddress sets BusinessAddress field to given value.

### HasBusinessAddress

`func (o *Contact) HasBusinessAddress() bool`

HasBusinessAddress returns a boolean if a field has been set.

### SetBusinessAddressNil

`func (o *Contact) SetBusinessAddressNil(b bool)`

 SetBusinessAddressNil sets the value for BusinessAddress to be an explicit nil

### UnsetBusinessAddress
`func (o *Contact) UnsetBusinessAddress()`

UnsetBusinessAddress ensures that no value is present for BusinessAddress, not even an explicit nil
### GetBusinessHomePage

`func (o *Contact) GetBusinessHomePage() string`

GetBusinessHomePage returns the BusinessHomePage field if non-nil, zero value otherwise.

### GetBusinessHomePageOk

`func (o *Contact) GetBusinessHomePageOk() (*string, bool)`

GetBusinessHomePageOk returns a tuple with the BusinessHomePage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessHomePage

`func (o *Contact) SetBusinessHomePage(v string)`

SetBusinessHomePage sets BusinessHomePage field to given value.

### HasBusinessHomePage

`func (o *Contact) HasBusinessHomePage() bool`

HasBusinessHomePage returns a boolean if a field has been set.

### SetBusinessHomePageNil

`func (o *Contact) SetBusinessHomePageNil(b bool)`

 SetBusinessHomePageNil sets the value for BusinessHomePage to be an explicit nil

### UnsetBusinessHomePage
`func (o *Contact) UnsetBusinessHomePage()`

UnsetBusinessHomePage ensures that no value is present for BusinessHomePage, not even an explicit nil
### GetBusinessPhones

`func (o *Contact) GetBusinessPhones() []*string`

GetBusinessPhones returns the BusinessPhones field if non-nil, zero value otherwise.

### GetBusinessPhonesOk

`func (o *Contact) GetBusinessPhonesOk() (*[]*string, bool)`

GetBusinessPhonesOk returns a tuple with the BusinessPhones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessPhones

`func (o *Contact) SetBusinessPhones(v []*string)`

SetBusinessPhones sets BusinessPhones field to given value.

### HasBusinessPhones

`func (o *Contact) HasBusinessPhones() bool`

HasBusinessPhones returns a boolean if a field has been set.

### GetChildren

`func (o *Contact) GetChildren() []*string`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *Contact) GetChildrenOk() (*[]*string, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *Contact) SetChildren(v []*string)`

SetChildren sets Children field to given value.

### HasChildren

`func (o *Contact) HasChildren() bool`

HasChildren returns a boolean if a field has been set.

### GetCompanyName

`func (o *Contact) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *Contact) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *Contact) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *Contact) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### SetCompanyNameNil

`func (o *Contact) SetCompanyNameNil(b bool)`

 SetCompanyNameNil sets the value for CompanyName to be an explicit nil

### UnsetCompanyName
`func (o *Contact) UnsetCompanyName()`

UnsetCompanyName ensures that no value is present for CompanyName, not even an explicit nil
### GetDepartment

`func (o *Contact) GetDepartment() string`

GetDepartment returns the Department field if non-nil, zero value otherwise.

### GetDepartmentOk

`func (o *Contact) GetDepartmentOk() (*string, bool)`

GetDepartmentOk returns a tuple with the Department field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartment

`func (o *Contact) SetDepartment(v string)`

SetDepartment sets Department field to given value.

### HasDepartment

`func (o *Contact) HasDepartment() bool`

HasDepartment returns a boolean if a field has been set.

### SetDepartmentNil

`func (o *Contact) SetDepartmentNil(b bool)`

 SetDepartmentNil sets the value for Department to be an explicit nil

### UnsetDepartment
`func (o *Contact) UnsetDepartment()`

UnsetDepartment ensures that no value is present for Department, not even an explicit nil
### GetDisplayName

`func (o *Contact) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *Contact) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *Contact) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *Contact) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *Contact) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *Contact) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetEmailAddresses

`func (o *Contact) GetEmailAddresses() []*AnyOfmicrosoftGraphEmailAddress`

GetEmailAddresses returns the EmailAddresses field if non-nil, zero value otherwise.

### GetEmailAddressesOk

`func (o *Contact) GetEmailAddressesOk() (*[]*AnyOfmicrosoftGraphEmailAddress, bool)`

GetEmailAddressesOk returns a tuple with the EmailAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddresses

`func (o *Contact) SetEmailAddresses(v []*AnyOfmicrosoftGraphEmailAddress)`

SetEmailAddresses sets EmailAddresses field to given value.

### HasEmailAddresses

`func (o *Contact) HasEmailAddresses() bool`

HasEmailAddresses returns a boolean if a field has been set.

### GetFileAs

`func (o *Contact) GetFileAs() string`

GetFileAs returns the FileAs field if non-nil, zero value otherwise.

### GetFileAsOk

`func (o *Contact) GetFileAsOk() (*string, bool)`

GetFileAsOk returns a tuple with the FileAs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileAs

`func (o *Contact) SetFileAs(v string)`

SetFileAs sets FileAs field to given value.

### HasFileAs

`func (o *Contact) HasFileAs() bool`

HasFileAs returns a boolean if a field has been set.

### SetFileAsNil

`func (o *Contact) SetFileAsNil(b bool)`

 SetFileAsNil sets the value for FileAs to be an explicit nil

### UnsetFileAs
`func (o *Contact) UnsetFileAs()`

UnsetFileAs ensures that no value is present for FileAs, not even an explicit nil
### GetGeneration

`func (o *Contact) GetGeneration() string`

GetGeneration returns the Generation field if non-nil, zero value otherwise.

### GetGenerationOk

`func (o *Contact) GetGenerationOk() (*string, bool)`

GetGenerationOk returns a tuple with the Generation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneration

`func (o *Contact) SetGeneration(v string)`

SetGeneration sets Generation field to given value.

### HasGeneration

`func (o *Contact) HasGeneration() bool`

HasGeneration returns a boolean if a field has been set.

### SetGenerationNil

`func (o *Contact) SetGenerationNil(b bool)`

 SetGenerationNil sets the value for Generation to be an explicit nil

### UnsetGeneration
`func (o *Contact) UnsetGeneration()`

UnsetGeneration ensures that no value is present for Generation, not even an explicit nil
### GetGivenName

`func (o *Contact) GetGivenName() string`

GetGivenName returns the GivenName field if non-nil, zero value otherwise.

### GetGivenNameOk

`func (o *Contact) GetGivenNameOk() (*string, bool)`

GetGivenNameOk returns a tuple with the GivenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGivenName

`func (o *Contact) SetGivenName(v string)`

SetGivenName sets GivenName field to given value.

### HasGivenName

`func (o *Contact) HasGivenName() bool`

HasGivenName returns a boolean if a field has been set.

### SetGivenNameNil

`func (o *Contact) SetGivenNameNil(b bool)`

 SetGivenNameNil sets the value for GivenName to be an explicit nil

### UnsetGivenName
`func (o *Contact) UnsetGivenName()`

UnsetGivenName ensures that no value is present for GivenName, not even an explicit nil
### GetHomeAddress

`func (o *Contact) GetHomeAddress() AnyOfmicrosoftGraphPhysicalAddress`

GetHomeAddress returns the HomeAddress field if non-nil, zero value otherwise.

### GetHomeAddressOk

`func (o *Contact) GetHomeAddressOk() (*AnyOfmicrosoftGraphPhysicalAddress, bool)`

GetHomeAddressOk returns a tuple with the HomeAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomeAddress

`func (o *Contact) SetHomeAddress(v AnyOfmicrosoftGraphPhysicalAddress)`

SetHomeAddress sets HomeAddress field to given value.

### HasHomeAddress

`func (o *Contact) HasHomeAddress() bool`

HasHomeAddress returns a boolean if a field has been set.

### SetHomeAddressNil

`func (o *Contact) SetHomeAddressNil(b bool)`

 SetHomeAddressNil sets the value for HomeAddress to be an explicit nil

### UnsetHomeAddress
`func (o *Contact) UnsetHomeAddress()`

UnsetHomeAddress ensures that no value is present for HomeAddress, not even an explicit nil
### GetHomePhones

`func (o *Contact) GetHomePhones() []*string`

GetHomePhones returns the HomePhones field if non-nil, zero value otherwise.

### GetHomePhonesOk

`func (o *Contact) GetHomePhonesOk() (*[]*string, bool)`

GetHomePhonesOk returns a tuple with the HomePhones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomePhones

`func (o *Contact) SetHomePhones(v []*string)`

SetHomePhones sets HomePhones field to given value.

### HasHomePhones

`func (o *Contact) HasHomePhones() bool`

HasHomePhones returns a boolean if a field has been set.

### GetImAddresses

`func (o *Contact) GetImAddresses() []*string`

GetImAddresses returns the ImAddresses field if non-nil, zero value otherwise.

### GetImAddressesOk

`func (o *Contact) GetImAddressesOk() (*[]*string, bool)`

GetImAddressesOk returns a tuple with the ImAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImAddresses

`func (o *Contact) SetImAddresses(v []*string)`

SetImAddresses sets ImAddresses field to given value.

### HasImAddresses

`func (o *Contact) HasImAddresses() bool`

HasImAddresses returns a boolean if a field has been set.

### GetInitials

`func (o *Contact) GetInitials() string`

GetInitials returns the Initials field if non-nil, zero value otherwise.

### GetInitialsOk

`func (o *Contact) GetInitialsOk() (*string, bool)`

GetInitialsOk returns a tuple with the Initials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitials

`func (o *Contact) SetInitials(v string)`

SetInitials sets Initials field to given value.

### HasInitials

`func (o *Contact) HasInitials() bool`

HasInitials returns a boolean if a field has been set.

### SetInitialsNil

`func (o *Contact) SetInitialsNil(b bool)`

 SetInitialsNil sets the value for Initials to be an explicit nil

### UnsetInitials
`func (o *Contact) UnsetInitials()`

UnsetInitials ensures that no value is present for Initials, not even an explicit nil
### GetJobTitle

`func (o *Contact) GetJobTitle() string`

GetJobTitle returns the JobTitle field if non-nil, zero value otherwise.

### GetJobTitleOk

`func (o *Contact) GetJobTitleOk() (*string, bool)`

GetJobTitleOk returns a tuple with the JobTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobTitle

`func (o *Contact) SetJobTitle(v string)`

SetJobTitle sets JobTitle field to given value.

### HasJobTitle

`func (o *Contact) HasJobTitle() bool`

HasJobTitle returns a boolean if a field has been set.

### SetJobTitleNil

`func (o *Contact) SetJobTitleNil(b bool)`

 SetJobTitleNil sets the value for JobTitle to be an explicit nil

### UnsetJobTitle
`func (o *Contact) UnsetJobTitle()`

UnsetJobTitle ensures that no value is present for JobTitle, not even an explicit nil
### GetManager

`func (o *Contact) GetManager() string`

GetManager returns the Manager field if non-nil, zero value otherwise.

### GetManagerOk

`func (o *Contact) GetManagerOk() (*string, bool)`

GetManagerOk returns a tuple with the Manager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManager

`func (o *Contact) SetManager(v string)`

SetManager sets Manager field to given value.

### HasManager

`func (o *Contact) HasManager() bool`

HasManager returns a boolean if a field has been set.

### SetManagerNil

`func (o *Contact) SetManagerNil(b bool)`

 SetManagerNil sets the value for Manager to be an explicit nil

### UnsetManager
`func (o *Contact) UnsetManager()`

UnsetManager ensures that no value is present for Manager, not even an explicit nil
### GetMiddleName

`func (o *Contact) GetMiddleName() string`

GetMiddleName returns the MiddleName field if non-nil, zero value otherwise.

### GetMiddleNameOk

`func (o *Contact) GetMiddleNameOk() (*string, bool)`

GetMiddleNameOk returns a tuple with the MiddleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiddleName

`func (o *Contact) SetMiddleName(v string)`

SetMiddleName sets MiddleName field to given value.

### HasMiddleName

`func (o *Contact) HasMiddleName() bool`

HasMiddleName returns a boolean if a field has been set.

### SetMiddleNameNil

`func (o *Contact) SetMiddleNameNil(b bool)`

 SetMiddleNameNil sets the value for MiddleName to be an explicit nil

### UnsetMiddleName
`func (o *Contact) UnsetMiddleName()`

UnsetMiddleName ensures that no value is present for MiddleName, not even an explicit nil
### GetMobilePhone

`func (o *Contact) GetMobilePhone() string`

GetMobilePhone returns the MobilePhone field if non-nil, zero value otherwise.

### GetMobilePhoneOk

`func (o *Contact) GetMobilePhoneOk() (*string, bool)`

GetMobilePhoneOk returns a tuple with the MobilePhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobilePhone

`func (o *Contact) SetMobilePhone(v string)`

SetMobilePhone sets MobilePhone field to given value.

### HasMobilePhone

`func (o *Contact) HasMobilePhone() bool`

HasMobilePhone returns a boolean if a field has been set.

### SetMobilePhoneNil

`func (o *Contact) SetMobilePhoneNil(b bool)`

 SetMobilePhoneNil sets the value for MobilePhone to be an explicit nil

### UnsetMobilePhone
`func (o *Contact) UnsetMobilePhone()`

UnsetMobilePhone ensures that no value is present for MobilePhone, not even an explicit nil
### GetNickName

`func (o *Contact) GetNickName() string`

GetNickName returns the NickName field if non-nil, zero value otherwise.

### GetNickNameOk

`func (o *Contact) GetNickNameOk() (*string, bool)`

GetNickNameOk returns a tuple with the NickName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNickName

`func (o *Contact) SetNickName(v string)`

SetNickName sets NickName field to given value.

### HasNickName

`func (o *Contact) HasNickName() bool`

HasNickName returns a boolean if a field has been set.

### SetNickNameNil

`func (o *Contact) SetNickNameNil(b bool)`

 SetNickNameNil sets the value for NickName to be an explicit nil

### UnsetNickName
`func (o *Contact) UnsetNickName()`

UnsetNickName ensures that no value is present for NickName, not even an explicit nil
### GetOfficeLocation

`func (o *Contact) GetOfficeLocation() string`

GetOfficeLocation returns the OfficeLocation field if non-nil, zero value otherwise.

### GetOfficeLocationOk

`func (o *Contact) GetOfficeLocationOk() (*string, bool)`

GetOfficeLocationOk returns a tuple with the OfficeLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfficeLocation

`func (o *Contact) SetOfficeLocation(v string)`

SetOfficeLocation sets OfficeLocation field to given value.

### HasOfficeLocation

`func (o *Contact) HasOfficeLocation() bool`

HasOfficeLocation returns a boolean if a field has been set.

### SetOfficeLocationNil

`func (o *Contact) SetOfficeLocationNil(b bool)`

 SetOfficeLocationNil sets the value for OfficeLocation to be an explicit nil

### UnsetOfficeLocation
`func (o *Contact) UnsetOfficeLocation()`

UnsetOfficeLocation ensures that no value is present for OfficeLocation, not even an explicit nil
### GetOtherAddress

`func (o *Contact) GetOtherAddress() AnyOfmicrosoftGraphPhysicalAddress`

GetOtherAddress returns the OtherAddress field if non-nil, zero value otherwise.

### GetOtherAddressOk

`func (o *Contact) GetOtherAddressOk() (*AnyOfmicrosoftGraphPhysicalAddress, bool)`

GetOtherAddressOk returns a tuple with the OtherAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherAddress

`func (o *Contact) SetOtherAddress(v AnyOfmicrosoftGraphPhysicalAddress)`

SetOtherAddress sets OtherAddress field to given value.

### HasOtherAddress

`func (o *Contact) HasOtherAddress() bool`

HasOtherAddress returns a boolean if a field has been set.

### SetOtherAddressNil

`func (o *Contact) SetOtherAddressNil(b bool)`

 SetOtherAddressNil sets the value for OtherAddress to be an explicit nil

### UnsetOtherAddress
`func (o *Contact) UnsetOtherAddress()`

UnsetOtherAddress ensures that no value is present for OtherAddress, not even an explicit nil
### GetParentFolderId

`func (o *Contact) GetParentFolderId() string`

GetParentFolderId returns the ParentFolderId field if non-nil, zero value otherwise.

### GetParentFolderIdOk

`func (o *Contact) GetParentFolderIdOk() (*string, bool)`

GetParentFolderIdOk returns a tuple with the ParentFolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentFolderId

`func (o *Contact) SetParentFolderId(v string)`

SetParentFolderId sets ParentFolderId field to given value.

### HasParentFolderId

`func (o *Contact) HasParentFolderId() bool`

HasParentFolderId returns a boolean if a field has been set.

### SetParentFolderIdNil

`func (o *Contact) SetParentFolderIdNil(b bool)`

 SetParentFolderIdNil sets the value for ParentFolderId to be an explicit nil

### UnsetParentFolderId
`func (o *Contact) UnsetParentFolderId()`

UnsetParentFolderId ensures that no value is present for ParentFolderId, not even an explicit nil
### GetPersonalNotes

`func (o *Contact) GetPersonalNotes() string`

GetPersonalNotes returns the PersonalNotes field if non-nil, zero value otherwise.

### GetPersonalNotesOk

`func (o *Contact) GetPersonalNotesOk() (*string, bool)`

GetPersonalNotesOk returns a tuple with the PersonalNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonalNotes

`func (o *Contact) SetPersonalNotes(v string)`

SetPersonalNotes sets PersonalNotes field to given value.

### HasPersonalNotes

`func (o *Contact) HasPersonalNotes() bool`

HasPersonalNotes returns a boolean if a field has been set.

### SetPersonalNotesNil

`func (o *Contact) SetPersonalNotesNil(b bool)`

 SetPersonalNotesNil sets the value for PersonalNotes to be an explicit nil

### UnsetPersonalNotes
`func (o *Contact) UnsetPersonalNotes()`

UnsetPersonalNotes ensures that no value is present for PersonalNotes, not even an explicit nil
### GetProfession

`func (o *Contact) GetProfession() string`

GetProfession returns the Profession field if non-nil, zero value otherwise.

### GetProfessionOk

`func (o *Contact) GetProfessionOk() (*string, bool)`

GetProfessionOk returns a tuple with the Profession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfession

`func (o *Contact) SetProfession(v string)`

SetProfession sets Profession field to given value.

### HasProfession

`func (o *Contact) HasProfession() bool`

HasProfession returns a boolean if a field has been set.

### SetProfessionNil

`func (o *Contact) SetProfessionNil(b bool)`

 SetProfessionNil sets the value for Profession to be an explicit nil

### UnsetProfession
`func (o *Contact) UnsetProfession()`

UnsetProfession ensures that no value is present for Profession, not even an explicit nil
### GetSpouseName

`func (o *Contact) GetSpouseName() string`

GetSpouseName returns the SpouseName field if non-nil, zero value otherwise.

### GetSpouseNameOk

`func (o *Contact) GetSpouseNameOk() (*string, bool)`

GetSpouseNameOk returns a tuple with the SpouseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpouseName

`func (o *Contact) SetSpouseName(v string)`

SetSpouseName sets SpouseName field to given value.

### HasSpouseName

`func (o *Contact) HasSpouseName() bool`

HasSpouseName returns a boolean if a field has been set.

### SetSpouseNameNil

`func (o *Contact) SetSpouseNameNil(b bool)`

 SetSpouseNameNil sets the value for SpouseName to be an explicit nil

### UnsetSpouseName
`func (o *Contact) UnsetSpouseName()`

UnsetSpouseName ensures that no value is present for SpouseName, not even an explicit nil
### GetSurname

`func (o *Contact) GetSurname() string`

GetSurname returns the Surname field if non-nil, zero value otherwise.

### GetSurnameOk

`func (o *Contact) GetSurnameOk() (*string, bool)`

GetSurnameOk returns a tuple with the Surname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurname

`func (o *Contact) SetSurname(v string)`

SetSurname sets Surname field to given value.

### HasSurname

`func (o *Contact) HasSurname() bool`

HasSurname returns a boolean if a field has been set.

### SetSurnameNil

`func (o *Contact) SetSurnameNil(b bool)`

 SetSurnameNil sets the value for Surname to be an explicit nil

### UnsetSurname
`func (o *Contact) UnsetSurname()`

UnsetSurname ensures that no value is present for Surname, not even an explicit nil
### GetTitle

`func (o *Contact) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Contact) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Contact) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *Contact) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *Contact) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *Contact) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetYomiCompanyName

`func (o *Contact) GetYomiCompanyName() string`

GetYomiCompanyName returns the YomiCompanyName field if non-nil, zero value otherwise.

### GetYomiCompanyNameOk

`func (o *Contact) GetYomiCompanyNameOk() (*string, bool)`

GetYomiCompanyNameOk returns a tuple with the YomiCompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYomiCompanyName

`func (o *Contact) SetYomiCompanyName(v string)`

SetYomiCompanyName sets YomiCompanyName field to given value.

### HasYomiCompanyName

`func (o *Contact) HasYomiCompanyName() bool`

HasYomiCompanyName returns a boolean if a field has been set.

### SetYomiCompanyNameNil

`func (o *Contact) SetYomiCompanyNameNil(b bool)`

 SetYomiCompanyNameNil sets the value for YomiCompanyName to be an explicit nil

### UnsetYomiCompanyName
`func (o *Contact) UnsetYomiCompanyName()`

UnsetYomiCompanyName ensures that no value is present for YomiCompanyName, not even an explicit nil
### GetYomiGivenName

`func (o *Contact) GetYomiGivenName() string`

GetYomiGivenName returns the YomiGivenName field if non-nil, zero value otherwise.

### GetYomiGivenNameOk

`func (o *Contact) GetYomiGivenNameOk() (*string, bool)`

GetYomiGivenNameOk returns a tuple with the YomiGivenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYomiGivenName

`func (o *Contact) SetYomiGivenName(v string)`

SetYomiGivenName sets YomiGivenName field to given value.

### HasYomiGivenName

`func (o *Contact) HasYomiGivenName() bool`

HasYomiGivenName returns a boolean if a field has been set.

### SetYomiGivenNameNil

`func (o *Contact) SetYomiGivenNameNil(b bool)`

 SetYomiGivenNameNil sets the value for YomiGivenName to be an explicit nil

### UnsetYomiGivenName
`func (o *Contact) UnsetYomiGivenName()`

UnsetYomiGivenName ensures that no value is present for YomiGivenName, not even an explicit nil
### GetYomiSurname

`func (o *Contact) GetYomiSurname() string`

GetYomiSurname returns the YomiSurname field if non-nil, zero value otherwise.

### GetYomiSurnameOk

`func (o *Contact) GetYomiSurnameOk() (*string, bool)`

GetYomiSurnameOk returns a tuple with the YomiSurname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYomiSurname

`func (o *Contact) SetYomiSurname(v string)`

SetYomiSurname sets YomiSurname field to given value.

### HasYomiSurname

`func (o *Contact) HasYomiSurname() bool`

HasYomiSurname returns a boolean if a field has been set.

### SetYomiSurnameNil

`func (o *Contact) SetYomiSurnameNil(b bool)`

 SetYomiSurnameNil sets the value for YomiSurname to be an explicit nil

### UnsetYomiSurname
`func (o *Contact) UnsetYomiSurname()`

UnsetYomiSurname ensures that no value is present for YomiSurname, not even an explicit nil
### GetExtensions

`func (o *Contact) GetExtensions() []MicrosoftGraphExtension`

GetExtensions returns the Extensions field if non-nil, zero value otherwise.

### GetExtensionsOk

`func (o *Contact) GetExtensionsOk() (*[]MicrosoftGraphExtension, bool)`

GetExtensionsOk returns a tuple with the Extensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensions

`func (o *Contact) SetExtensions(v []MicrosoftGraphExtension)`

SetExtensions sets Extensions field to given value.

### HasExtensions

`func (o *Contact) HasExtensions() bool`

HasExtensions returns a boolean if a field has been set.

### GetMultiValueExtendedProperties

`func (o *Contact) GetMultiValueExtendedProperties() []MicrosoftGraphMultiValueLegacyExtendedProperty`

GetMultiValueExtendedProperties returns the MultiValueExtendedProperties field if non-nil, zero value otherwise.

### GetMultiValueExtendedPropertiesOk

`func (o *Contact) GetMultiValueExtendedPropertiesOk() (*[]MicrosoftGraphMultiValueLegacyExtendedProperty, bool)`

GetMultiValueExtendedPropertiesOk returns a tuple with the MultiValueExtendedProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiValueExtendedProperties

`func (o *Contact) SetMultiValueExtendedProperties(v []MicrosoftGraphMultiValueLegacyExtendedProperty)`

SetMultiValueExtendedProperties sets MultiValueExtendedProperties field to given value.

### HasMultiValueExtendedProperties

`func (o *Contact) HasMultiValueExtendedProperties() bool`

HasMultiValueExtendedProperties returns a boolean if a field has been set.

### GetPhoto

`func (o *Contact) GetPhoto() AnyOfmicrosoftGraphProfilePhoto`

GetPhoto returns the Photo field if non-nil, zero value otherwise.

### GetPhotoOk

`func (o *Contact) GetPhotoOk() (*AnyOfmicrosoftGraphProfilePhoto, bool)`

GetPhotoOk returns a tuple with the Photo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoto

`func (o *Contact) SetPhoto(v AnyOfmicrosoftGraphProfilePhoto)`

SetPhoto sets Photo field to given value.

### HasPhoto

`func (o *Contact) HasPhoto() bool`

HasPhoto returns a boolean if a field has been set.

### SetPhotoNil

`func (o *Contact) SetPhotoNil(b bool)`

 SetPhotoNil sets the value for Photo to be an explicit nil

### UnsetPhoto
`func (o *Contact) UnsetPhoto()`

UnsetPhoto ensures that no value is present for Photo, not even an explicit nil
### GetSingleValueExtendedProperties

`func (o *Contact) GetSingleValueExtendedProperties() []MicrosoftGraphSingleValueLegacyExtendedProperty`

GetSingleValueExtendedProperties returns the SingleValueExtendedProperties field if non-nil, zero value otherwise.

### GetSingleValueExtendedPropertiesOk

`func (o *Contact) GetSingleValueExtendedPropertiesOk() (*[]MicrosoftGraphSingleValueLegacyExtendedProperty, bool)`

GetSingleValueExtendedPropertiesOk returns a tuple with the SingleValueExtendedProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSingleValueExtendedProperties

`func (o *Contact) SetSingleValueExtendedProperties(v []MicrosoftGraphSingleValueLegacyExtendedProperty)`

SetSingleValueExtendedProperties sets SingleValueExtendedProperties field to given value.

### HasSingleValueExtendedProperties

`func (o *Contact) HasSingleValueExtendedProperties() bool`

HasSingleValueExtendedProperties returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


