# MicrosoftGraphContact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**Categories** | Pointer to **[]string** | The categories associated with the item | [optional] 
**ChangeKey** | Pointer to **NullableString** | Identifies the version of the item. Every time the item is changed, changeKey changes as well. This allows Exchange to apply changes to the correct version of the object. Read-only. | [optional] 
**CreatedDateTime** | Pointer to **NullableTime** | The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z | [optional] 
**LastModifiedDateTime** | Pointer to **NullableTime** | The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z | [optional] 
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

### NewMicrosoftGraphContact

`func NewMicrosoftGraphContact() *MicrosoftGraphContact`

NewMicrosoftGraphContact instantiates a new MicrosoftGraphContact object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphContactWithDefaults

`func NewMicrosoftGraphContactWithDefaults() *MicrosoftGraphContact`

NewMicrosoftGraphContactWithDefaults instantiates a new MicrosoftGraphContact object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphContact) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphContact) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphContact) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphContact) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCategories

`func (o *MicrosoftGraphContact) GetCategories() []*string`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *MicrosoftGraphContact) GetCategoriesOk() (*[]*string, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *MicrosoftGraphContact) SetCategories(v []*string)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *MicrosoftGraphContact) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### GetChangeKey

`func (o *MicrosoftGraphContact) GetChangeKey() string`

GetChangeKey returns the ChangeKey field if non-nil, zero value otherwise.

### GetChangeKeyOk

`func (o *MicrosoftGraphContact) GetChangeKeyOk() (*string, bool)`

GetChangeKeyOk returns a tuple with the ChangeKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeKey

`func (o *MicrosoftGraphContact) SetChangeKey(v string)`

SetChangeKey sets ChangeKey field to given value.

### HasChangeKey

`func (o *MicrosoftGraphContact) HasChangeKey() bool`

HasChangeKey returns a boolean if a field has been set.

### SetChangeKeyNil

`func (o *MicrosoftGraphContact) SetChangeKeyNil(b bool)`

 SetChangeKeyNil sets the value for ChangeKey to be an explicit nil

### UnsetChangeKey
`func (o *MicrosoftGraphContact) UnsetChangeKey()`

UnsetChangeKey ensures that no value is present for ChangeKey, not even an explicit nil
### GetCreatedDateTime

`func (o *MicrosoftGraphContact) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *MicrosoftGraphContact) GetCreatedDateTimeOk() (*time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateTime

`func (o *MicrosoftGraphContact) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime sets CreatedDateTime field to given value.

### HasCreatedDateTime

`func (o *MicrosoftGraphContact) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTimeNil

`func (o *MicrosoftGraphContact) SetCreatedDateTimeNil(b bool)`

 SetCreatedDateTimeNil sets the value for CreatedDateTime to be an explicit nil

### UnsetCreatedDateTime
`func (o *MicrosoftGraphContact) UnsetCreatedDateTime()`

UnsetCreatedDateTime ensures that no value is present for CreatedDateTime, not even an explicit nil
### GetLastModifiedDateTime

`func (o *MicrosoftGraphContact) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *MicrosoftGraphContact) GetLastModifiedDateTimeOk() (*time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedDateTime

`func (o *MicrosoftGraphContact) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime sets LastModifiedDateTime field to given value.

### HasLastModifiedDateTime

`func (o *MicrosoftGraphContact) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### SetLastModifiedDateTimeNil

`func (o *MicrosoftGraphContact) SetLastModifiedDateTimeNil(b bool)`

 SetLastModifiedDateTimeNil sets the value for LastModifiedDateTime to be an explicit nil

### UnsetLastModifiedDateTime
`func (o *MicrosoftGraphContact) UnsetLastModifiedDateTime()`

UnsetLastModifiedDateTime ensures that no value is present for LastModifiedDateTime, not even an explicit nil
### GetAssistantName

`func (o *MicrosoftGraphContact) GetAssistantName() string`

GetAssistantName returns the AssistantName field if non-nil, zero value otherwise.

### GetAssistantNameOk

`func (o *MicrosoftGraphContact) GetAssistantNameOk() (*string, bool)`

GetAssistantNameOk returns a tuple with the AssistantName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssistantName

`func (o *MicrosoftGraphContact) SetAssistantName(v string)`

SetAssistantName sets AssistantName field to given value.

### HasAssistantName

`func (o *MicrosoftGraphContact) HasAssistantName() bool`

HasAssistantName returns a boolean if a field has been set.

### SetAssistantNameNil

`func (o *MicrosoftGraphContact) SetAssistantNameNil(b bool)`

 SetAssistantNameNil sets the value for AssistantName to be an explicit nil

### UnsetAssistantName
`func (o *MicrosoftGraphContact) UnsetAssistantName()`

UnsetAssistantName ensures that no value is present for AssistantName, not even an explicit nil
### GetBirthday

`func (o *MicrosoftGraphContact) GetBirthday() time.Time`

GetBirthday returns the Birthday field if non-nil, zero value otherwise.

### GetBirthdayOk

`func (o *MicrosoftGraphContact) GetBirthdayOk() (*time.Time, bool)`

GetBirthdayOk returns a tuple with the Birthday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBirthday

`func (o *MicrosoftGraphContact) SetBirthday(v time.Time)`

SetBirthday sets Birthday field to given value.

### HasBirthday

`func (o *MicrosoftGraphContact) HasBirthday() bool`

HasBirthday returns a boolean if a field has been set.

### SetBirthdayNil

`func (o *MicrosoftGraphContact) SetBirthdayNil(b bool)`

 SetBirthdayNil sets the value for Birthday to be an explicit nil

### UnsetBirthday
`func (o *MicrosoftGraphContact) UnsetBirthday()`

UnsetBirthday ensures that no value is present for Birthday, not even an explicit nil
### GetBusinessAddress

`func (o *MicrosoftGraphContact) GetBusinessAddress() AnyOfmicrosoftGraphPhysicalAddress`

GetBusinessAddress returns the BusinessAddress field if non-nil, zero value otherwise.

### GetBusinessAddressOk

`func (o *MicrosoftGraphContact) GetBusinessAddressOk() (*AnyOfmicrosoftGraphPhysicalAddress, bool)`

GetBusinessAddressOk returns a tuple with the BusinessAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessAddress

`func (o *MicrosoftGraphContact) SetBusinessAddress(v AnyOfmicrosoftGraphPhysicalAddress)`

SetBusinessAddress sets BusinessAddress field to given value.

### HasBusinessAddress

`func (o *MicrosoftGraphContact) HasBusinessAddress() bool`

HasBusinessAddress returns a boolean if a field has been set.

### SetBusinessAddressNil

`func (o *MicrosoftGraphContact) SetBusinessAddressNil(b bool)`

 SetBusinessAddressNil sets the value for BusinessAddress to be an explicit nil

### UnsetBusinessAddress
`func (o *MicrosoftGraphContact) UnsetBusinessAddress()`

UnsetBusinessAddress ensures that no value is present for BusinessAddress, not even an explicit nil
### GetBusinessHomePage

`func (o *MicrosoftGraphContact) GetBusinessHomePage() string`

GetBusinessHomePage returns the BusinessHomePage field if non-nil, zero value otherwise.

### GetBusinessHomePageOk

`func (o *MicrosoftGraphContact) GetBusinessHomePageOk() (*string, bool)`

GetBusinessHomePageOk returns a tuple with the BusinessHomePage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessHomePage

`func (o *MicrosoftGraphContact) SetBusinessHomePage(v string)`

SetBusinessHomePage sets BusinessHomePage field to given value.

### HasBusinessHomePage

`func (o *MicrosoftGraphContact) HasBusinessHomePage() bool`

HasBusinessHomePage returns a boolean if a field has been set.

### SetBusinessHomePageNil

`func (o *MicrosoftGraphContact) SetBusinessHomePageNil(b bool)`

 SetBusinessHomePageNil sets the value for BusinessHomePage to be an explicit nil

### UnsetBusinessHomePage
`func (o *MicrosoftGraphContact) UnsetBusinessHomePage()`

UnsetBusinessHomePage ensures that no value is present for BusinessHomePage, not even an explicit nil
### GetBusinessPhones

`func (o *MicrosoftGraphContact) GetBusinessPhones() []*string`

GetBusinessPhones returns the BusinessPhones field if non-nil, zero value otherwise.

### GetBusinessPhonesOk

`func (o *MicrosoftGraphContact) GetBusinessPhonesOk() (*[]*string, bool)`

GetBusinessPhonesOk returns a tuple with the BusinessPhones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessPhones

`func (o *MicrosoftGraphContact) SetBusinessPhones(v []*string)`

SetBusinessPhones sets BusinessPhones field to given value.

### HasBusinessPhones

`func (o *MicrosoftGraphContact) HasBusinessPhones() bool`

HasBusinessPhones returns a boolean if a field has been set.

### GetChildren

`func (o *MicrosoftGraphContact) GetChildren() []*string`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *MicrosoftGraphContact) GetChildrenOk() (*[]*string, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *MicrosoftGraphContact) SetChildren(v []*string)`

SetChildren sets Children field to given value.

### HasChildren

`func (o *MicrosoftGraphContact) HasChildren() bool`

HasChildren returns a boolean if a field has been set.

### GetCompanyName

`func (o *MicrosoftGraphContact) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *MicrosoftGraphContact) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *MicrosoftGraphContact) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *MicrosoftGraphContact) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### SetCompanyNameNil

`func (o *MicrosoftGraphContact) SetCompanyNameNil(b bool)`

 SetCompanyNameNil sets the value for CompanyName to be an explicit nil

### UnsetCompanyName
`func (o *MicrosoftGraphContact) UnsetCompanyName()`

UnsetCompanyName ensures that no value is present for CompanyName, not even an explicit nil
### GetDepartment

`func (o *MicrosoftGraphContact) GetDepartment() string`

GetDepartment returns the Department field if non-nil, zero value otherwise.

### GetDepartmentOk

`func (o *MicrosoftGraphContact) GetDepartmentOk() (*string, bool)`

GetDepartmentOk returns a tuple with the Department field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartment

`func (o *MicrosoftGraphContact) SetDepartment(v string)`

SetDepartment sets Department field to given value.

### HasDepartment

`func (o *MicrosoftGraphContact) HasDepartment() bool`

HasDepartment returns a boolean if a field has been set.

### SetDepartmentNil

`func (o *MicrosoftGraphContact) SetDepartmentNil(b bool)`

 SetDepartmentNil sets the value for Department to be an explicit nil

### UnsetDepartment
`func (o *MicrosoftGraphContact) UnsetDepartment()`

UnsetDepartment ensures that no value is present for Department, not even an explicit nil
### GetDisplayName

`func (o *MicrosoftGraphContact) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphContact) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *MicrosoftGraphContact) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *MicrosoftGraphContact) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *MicrosoftGraphContact) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *MicrosoftGraphContact) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetEmailAddresses

`func (o *MicrosoftGraphContact) GetEmailAddresses() []*AnyOfmicrosoftGraphEmailAddress`

GetEmailAddresses returns the EmailAddresses field if non-nil, zero value otherwise.

### GetEmailAddressesOk

`func (o *MicrosoftGraphContact) GetEmailAddressesOk() (*[]*AnyOfmicrosoftGraphEmailAddress, bool)`

GetEmailAddressesOk returns a tuple with the EmailAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddresses

`func (o *MicrosoftGraphContact) SetEmailAddresses(v []*AnyOfmicrosoftGraphEmailAddress)`

SetEmailAddresses sets EmailAddresses field to given value.

### HasEmailAddresses

`func (o *MicrosoftGraphContact) HasEmailAddresses() bool`

HasEmailAddresses returns a boolean if a field has been set.

### GetFileAs

`func (o *MicrosoftGraphContact) GetFileAs() string`

GetFileAs returns the FileAs field if non-nil, zero value otherwise.

### GetFileAsOk

`func (o *MicrosoftGraphContact) GetFileAsOk() (*string, bool)`

GetFileAsOk returns a tuple with the FileAs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileAs

`func (o *MicrosoftGraphContact) SetFileAs(v string)`

SetFileAs sets FileAs field to given value.

### HasFileAs

`func (o *MicrosoftGraphContact) HasFileAs() bool`

HasFileAs returns a boolean if a field has been set.

### SetFileAsNil

`func (o *MicrosoftGraphContact) SetFileAsNil(b bool)`

 SetFileAsNil sets the value for FileAs to be an explicit nil

### UnsetFileAs
`func (o *MicrosoftGraphContact) UnsetFileAs()`

UnsetFileAs ensures that no value is present for FileAs, not even an explicit nil
### GetGeneration

`func (o *MicrosoftGraphContact) GetGeneration() string`

GetGeneration returns the Generation field if non-nil, zero value otherwise.

### GetGenerationOk

`func (o *MicrosoftGraphContact) GetGenerationOk() (*string, bool)`

GetGenerationOk returns a tuple with the Generation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneration

`func (o *MicrosoftGraphContact) SetGeneration(v string)`

SetGeneration sets Generation field to given value.

### HasGeneration

`func (o *MicrosoftGraphContact) HasGeneration() bool`

HasGeneration returns a boolean if a field has been set.

### SetGenerationNil

`func (o *MicrosoftGraphContact) SetGenerationNil(b bool)`

 SetGenerationNil sets the value for Generation to be an explicit nil

### UnsetGeneration
`func (o *MicrosoftGraphContact) UnsetGeneration()`

UnsetGeneration ensures that no value is present for Generation, not even an explicit nil
### GetGivenName

`func (o *MicrosoftGraphContact) GetGivenName() string`

GetGivenName returns the GivenName field if non-nil, zero value otherwise.

### GetGivenNameOk

`func (o *MicrosoftGraphContact) GetGivenNameOk() (*string, bool)`

GetGivenNameOk returns a tuple with the GivenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGivenName

`func (o *MicrosoftGraphContact) SetGivenName(v string)`

SetGivenName sets GivenName field to given value.

### HasGivenName

`func (o *MicrosoftGraphContact) HasGivenName() bool`

HasGivenName returns a boolean if a field has been set.

### SetGivenNameNil

`func (o *MicrosoftGraphContact) SetGivenNameNil(b bool)`

 SetGivenNameNil sets the value for GivenName to be an explicit nil

### UnsetGivenName
`func (o *MicrosoftGraphContact) UnsetGivenName()`

UnsetGivenName ensures that no value is present for GivenName, not even an explicit nil
### GetHomeAddress

`func (o *MicrosoftGraphContact) GetHomeAddress() AnyOfmicrosoftGraphPhysicalAddress`

GetHomeAddress returns the HomeAddress field if non-nil, zero value otherwise.

### GetHomeAddressOk

`func (o *MicrosoftGraphContact) GetHomeAddressOk() (*AnyOfmicrosoftGraphPhysicalAddress, bool)`

GetHomeAddressOk returns a tuple with the HomeAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomeAddress

`func (o *MicrosoftGraphContact) SetHomeAddress(v AnyOfmicrosoftGraphPhysicalAddress)`

SetHomeAddress sets HomeAddress field to given value.

### HasHomeAddress

`func (o *MicrosoftGraphContact) HasHomeAddress() bool`

HasHomeAddress returns a boolean if a field has been set.

### SetHomeAddressNil

`func (o *MicrosoftGraphContact) SetHomeAddressNil(b bool)`

 SetHomeAddressNil sets the value for HomeAddress to be an explicit nil

### UnsetHomeAddress
`func (o *MicrosoftGraphContact) UnsetHomeAddress()`

UnsetHomeAddress ensures that no value is present for HomeAddress, not even an explicit nil
### GetHomePhones

`func (o *MicrosoftGraphContact) GetHomePhones() []*string`

GetHomePhones returns the HomePhones field if non-nil, zero value otherwise.

### GetHomePhonesOk

`func (o *MicrosoftGraphContact) GetHomePhonesOk() (*[]*string, bool)`

GetHomePhonesOk returns a tuple with the HomePhones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomePhones

`func (o *MicrosoftGraphContact) SetHomePhones(v []*string)`

SetHomePhones sets HomePhones field to given value.

### HasHomePhones

`func (o *MicrosoftGraphContact) HasHomePhones() bool`

HasHomePhones returns a boolean if a field has been set.

### GetImAddresses

`func (o *MicrosoftGraphContact) GetImAddresses() []*string`

GetImAddresses returns the ImAddresses field if non-nil, zero value otherwise.

### GetImAddressesOk

`func (o *MicrosoftGraphContact) GetImAddressesOk() (*[]*string, bool)`

GetImAddressesOk returns a tuple with the ImAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImAddresses

`func (o *MicrosoftGraphContact) SetImAddresses(v []*string)`

SetImAddresses sets ImAddresses field to given value.

### HasImAddresses

`func (o *MicrosoftGraphContact) HasImAddresses() bool`

HasImAddresses returns a boolean if a field has been set.

### GetInitials

`func (o *MicrosoftGraphContact) GetInitials() string`

GetInitials returns the Initials field if non-nil, zero value otherwise.

### GetInitialsOk

`func (o *MicrosoftGraphContact) GetInitialsOk() (*string, bool)`

GetInitialsOk returns a tuple with the Initials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitials

`func (o *MicrosoftGraphContact) SetInitials(v string)`

SetInitials sets Initials field to given value.

### HasInitials

`func (o *MicrosoftGraphContact) HasInitials() bool`

HasInitials returns a boolean if a field has been set.

### SetInitialsNil

`func (o *MicrosoftGraphContact) SetInitialsNil(b bool)`

 SetInitialsNil sets the value for Initials to be an explicit nil

### UnsetInitials
`func (o *MicrosoftGraphContact) UnsetInitials()`

UnsetInitials ensures that no value is present for Initials, not even an explicit nil
### GetJobTitle

`func (o *MicrosoftGraphContact) GetJobTitle() string`

GetJobTitle returns the JobTitle field if non-nil, zero value otherwise.

### GetJobTitleOk

`func (o *MicrosoftGraphContact) GetJobTitleOk() (*string, bool)`

GetJobTitleOk returns a tuple with the JobTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobTitle

`func (o *MicrosoftGraphContact) SetJobTitle(v string)`

SetJobTitle sets JobTitle field to given value.

### HasJobTitle

`func (o *MicrosoftGraphContact) HasJobTitle() bool`

HasJobTitle returns a boolean if a field has been set.

### SetJobTitleNil

`func (o *MicrosoftGraphContact) SetJobTitleNil(b bool)`

 SetJobTitleNil sets the value for JobTitle to be an explicit nil

### UnsetJobTitle
`func (o *MicrosoftGraphContact) UnsetJobTitle()`

UnsetJobTitle ensures that no value is present for JobTitle, not even an explicit nil
### GetManager

`func (o *MicrosoftGraphContact) GetManager() string`

GetManager returns the Manager field if non-nil, zero value otherwise.

### GetManagerOk

`func (o *MicrosoftGraphContact) GetManagerOk() (*string, bool)`

GetManagerOk returns a tuple with the Manager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManager

`func (o *MicrosoftGraphContact) SetManager(v string)`

SetManager sets Manager field to given value.

### HasManager

`func (o *MicrosoftGraphContact) HasManager() bool`

HasManager returns a boolean if a field has been set.

### SetManagerNil

`func (o *MicrosoftGraphContact) SetManagerNil(b bool)`

 SetManagerNil sets the value for Manager to be an explicit nil

### UnsetManager
`func (o *MicrosoftGraphContact) UnsetManager()`

UnsetManager ensures that no value is present for Manager, not even an explicit nil
### GetMiddleName

`func (o *MicrosoftGraphContact) GetMiddleName() string`

GetMiddleName returns the MiddleName field if non-nil, zero value otherwise.

### GetMiddleNameOk

`func (o *MicrosoftGraphContact) GetMiddleNameOk() (*string, bool)`

GetMiddleNameOk returns a tuple with the MiddleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiddleName

`func (o *MicrosoftGraphContact) SetMiddleName(v string)`

SetMiddleName sets MiddleName field to given value.

### HasMiddleName

`func (o *MicrosoftGraphContact) HasMiddleName() bool`

HasMiddleName returns a boolean if a field has been set.

### SetMiddleNameNil

`func (o *MicrosoftGraphContact) SetMiddleNameNil(b bool)`

 SetMiddleNameNil sets the value for MiddleName to be an explicit nil

### UnsetMiddleName
`func (o *MicrosoftGraphContact) UnsetMiddleName()`

UnsetMiddleName ensures that no value is present for MiddleName, not even an explicit nil
### GetMobilePhone

`func (o *MicrosoftGraphContact) GetMobilePhone() string`

GetMobilePhone returns the MobilePhone field if non-nil, zero value otherwise.

### GetMobilePhoneOk

`func (o *MicrosoftGraphContact) GetMobilePhoneOk() (*string, bool)`

GetMobilePhoneOk returns a tuple with the MobilePhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobilePhone

`func (o *MicrosoftGraphContact) SetMobilePhone(v string)`

SetMobilePhone sets MobilePhone field to given value.

### HasMobilePhone

`func (o *MicrosoftGraphContact) HasMobilePhone() bool`

HasMobilePhone returns a boolean if a field has been set.

### SetMobilePhoneNil

`func (o *MicrosoftGraphContact) SetMobilePhoneNil(b bool)`

 SetMobilePhoneNil sets the value for MobilePhone to be an explicit nil

### UnsetMobilePhone
`func (o *MicrosoftGraphContact) UnsetMobilePhone()`

UnsetMobilePhone ensures that no value is present for MobilePhone, not even an explicit nil
### GetNickName

`func (o *MicrosoftGraphContact) GetNickName() string`

GetNickName returns the NickName field if non-nil, zero value otherwise.

### GetNickNameOk

`func (o *MicrosoftGraphContact) GetNickNameOk() (*string, bool)`

GetNickNameOk returns a tuple with the NickName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNickName

`func (o *MicrosoftGraphContact) SetNickName(v string)`

SetNickName sets NickName field to given value.

### HasNickName

`func (o *MicrosoftGraphContact) HasNickName() bool`

HasNickName returns a boolean if a field has been set.

### SetNickNameNil

`func (o *MicrosoftGraphContact) SetNickNameNil(b bool)`

 SetNickNameNil sets the value for NickName to be an explicit nil

### UnsetNickName
`func (o *MicrosoftGraphContact) UnsetNickName()`

UnsetNickName ensures that no value is present for NickName, not even an explicit nil
### GetOfficeLocation

`func (o *MicrosoftGraphContact) GetOfficeLocation() string`

GetOfficeLocation returns the OfficeLocation field if non-nil, zero value otherwise.

### GetOfficeLocationOk

`func (o *MicrosoftGraphContact) GetOfficeLocationOk() (*string, bool)`

GetOfficeLocationOk returns a tuple with the OfficeLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfficeLocation

`func (o *MicrosoftGraphContact) SetOfficeLocation(v string)`

SetOfficeLocation sets OfficeLocation field to given value.

### HasOfficeLocation

`func (o *MicrosoftGraphContact) HasOfficeLocation() bool`

HasOfficeLocation returns a boolean if a field has been set.

### SetOfficeLocationNil

`func (o *MicrosoftGraphContact) SetOfficeLocationNil(b bool)`

 SetOfficeLocationNil sets the value for OfficeLocation to be an explicit nil

### UnsetOfficeLocation
`func (o *MicrosoftGraphContact) UnsetOfficeLocation()`

UnsetOfficeLocation ensures that no value is present for OfficeLocation, not even an explicit nil
### GetOtherAddress

`func (o *MicrosoftGraphContact) GetOtherAddress() AnyOfmicrosoftGraphPhysicalAddress`

GetOtherAddress returns the OtherAddress field if non-nil, zero value otherwise.

### GetOtherAddressOk

`func (o *MicrosoftGraphContact) GetOtherAddressOk() (*AnyOfmicrosoftGraphPhysicalAddress, bool)`

GetOtherAddressOk returns a tuple with the OtherAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherAddress

`func (o *MicrosoftGraphContact) SetOtherAddress(v AnyOfmicrosoftGraphPhysicalAddress)`

SetOtherAddress sets OtherAddress field to given value.

### HasOtherAddress

`func (o *MicrosoftGraphContact) HasOtherAddress() bool`

HasOtherAddress returns a boolean if a field has been set.

### SetOtherAddressNil

`func (o *MicrosoftGraphContact) SetOtherAddressNil(b bool)`

 SetOtherAddressNil sets the value for OtherAddress to be an explicit nil

### UnsetOtherAddress
`func (o *MicrosoftGraphContact) UnsetOtherAddress()`

UnsetOtherAddress ensures that no value is present for OtherAddress, not even an explicit nil
### GetParentFolderId

`func (o *MicrosoftGraphContact) GetParentFolderId() string`

GetParentFolderId returns the ParentFolderId field if non-nil, zero value otherwise.

### GetParentFolderIdOk

`func (o *MicrosoftGraphContact) GetParentFolderIdOk() (*string, bool)`

GetParentFolderIdOk returns a tuple with the ParentFolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentFolderId

`func (o *MicrosoftGraphContact) SetParentFolderId(v string)`

SetParentFolderId sets ParentFolderId field to given value.

### HasParentFolderId

`func (o *MicrosoftGraphContact) HasParentFolderId() bool`

HasParentFolderId returns a boolean if a field has been set.

### SetParentFolderIdNil

`func (o *MicrosoftGraphContact) SetParentFolderIdNil(b bool)`

 SetParentFolderIdNil sets the value for ParentFolderId to be an explicit nil

### UnsetParentFolderId
`func (o *MicrosoftGraphContact) UnsetParentFolderId()`

UnsetParentFolderId ensures that no value is present for ParentFolderId, not even an explicit nil
### GetPersonalNotes

`func (o *MicrosoftGraphContact) GetPersonalNotes() string`

GetPersonalNotes returns the PersonalNotes field if non-nil, zero value otherwise.

### GetPersonalNotesOk

`func (o *MicrosoftGraphContact) GetPersonalNotesOk() (*string, bool)`

GetPersonalNotesOk returns a tuple with the PersonalNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonalNotes

`func (o *MicrosoftGraphContact) SetPersonalNotes(v string)`

SetPersonalNotes sets PersonalNotes field to given value.

### HasPersonalNotes

`func (o *MicrosoftGraphContact) HasPersonalNotes() bool`

HasPersonalNotes returns a boolean if a field has been set.

### SetPersonalNotesNil

`func (o *MicrosoftGraphContact) SetPersonalNotesNil(b bool)`

 SetPersonalNotesNil sets the value for PersonalNotes to be an explicit nil

### UnsetPersonalNotes
`func (o *MicrosoftGraphContact) UnsetPersonalNotes()`

UnsetPersonalNotes ensures that no value is present for PersonalNotes, not even an explicit nil
### GetProfession

`func (o *MicrosoftGraphContact) GetProfession() string`

GetProfession returns the Profession field if non-nil, zero value otherwise.

### GetProfessionOk

`func (o *MicrosoftGraphContact) GetProfessionOk() (*string, bool)`

GetProfessionOk returns a tuple with the Profession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfession

`func (o *MicrosoftGraphContact) SetProfession(v string)`

SetProfession sets Profession field to given value.

### HasProfession

`func (o *MicrosoftGraphContact) HasProfession() bool`

HasProfession returns a boolean if a field has been set.

### SetProfessionNil

`func (o *MicrosoftGraphContact) SetProfessionNil(b bool)`

 SetProfessionNil sets the value for Profession to be an explicit nil

### UnsetProfession
`func (o *MicrosoftGraphContact) UnsetProfession()`

UnsetProfession ensures that no value is present for Profession, not even an explicit nil
### GetSpouseName

`func (o *MicrosoftGraphContact) GetSpouseName() string`

GetSpouseName returns the SpouseName field if non-nil, zero value otherwise.

### GetSpouseNameOk

`func (o *MicrosoftGraphContact) GetSpouseNameOk() (*string, bool)`

GetSpouseNameOk returns a tuple with the SpouseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpouseName

`func (o *MicrosoftGraphContact) SetSpouseName(v string)`

SetSpouseName sets SpouseName field to given value.

### HasSpouseName

`func (o *MicrosoftGraphContact) HasSpouseName() bool`

HasSpouseName returns a boolean if a field has been set.

### SetSpouseNameNil

`func (o *MicrosoftGraphContact) SetSpouseNameNil(b bool)`

 SetSpouseNameNil sets the value for SpouseName to be an explicit nil

### UnsetSpouseName
`func (o *MicrosoftGraphContact) UnsetSpouseName()`

UnsetSpouseName ensures that no value is present for SpouseName, not even an explicit nil
### GetSurname

`func (o *MicrosoftGraphContact) GetSurname() string`

GetSurname returns the Surname field if non-nil, zero value otherwise.

### GetSurnameOk

`func (o *MicrosoftGraphContact) GetSurnameOk() (*string, bool)`

GetSurnameOk returns a tuple with the Surname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurname

`func (o *MicrosoftGraphContact) SetSurname(v string)`

SetSurname sets Surname field to given value.

### HasSurname

`func (o *MicrosoftGraphContact) HasSurname() bool`

HasSurname returns a boolean if a field has been set.

### SetSurnameNil

`func (o *MicrosoftGraphContact) SetSurnameNil(b bool)`

 SetSurnameNil sets the value for Surname to be an explicit nil

### UnsetSurname
`func (o *MicrosoftGraphContact) UnsetSurname()`

UnsetSurname ensures that no value is present for Surname, not even an explicit nil
### GetTitle

`func (o *MicrosoftGraphContact) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *MicrosoftGraphContact) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *MicrosoftGraphContact) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *MicrosoftGraphContact) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *MicrosoftGraphContact) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *MicrosoftGraphContact) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetYomiCompanyName

`func (o *MicrosoftGraphContact) GetYomiCompanyName() string`

GetYomiCompanyName returns the YomiCompanyName field if non-nil, zero value otherwise.

### GetYomiCompanyNameOk

`func (o *MicrosoftGraphContact) GetYomiCompanyNameOk() (*string, bool)`

GetYomiCompanyNameOk returns a tuple with the YomiCompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYomiCompanyName

`func (o *MicrosoftGraphContact) SetYomiCompanyName(v string)`

SetYomiCompanyName sets YomiCompanyName field to given value.

### HasYomiCompanyName

`func (o *MicrosoftGraphContact) HasYomiCompanyName() bool`

HasYomiCompanyName returns a boolean if a field has been set.

### SetYomiCompanyNameNil

`func (o *MicrosoftGraphContact) SetYomiCompanyNameNil(b bool)`

 SetYomiCompanyNameNil sets the value for YomiCompanyName to be an explicit nil

### UnsetYomiCompanyName
`func (o *MicrosoftGraphContact) UnsetYomiCompanyName()`

UnsetYomiCompanyName ensures that no value is present for YomiCompanyName, not even an explicit nil
### GetYomiGivenName

`func (o *MicrosoftGraphContact) GetYomiGivenName() string`

GetYomiGivenName returns the YomiGivenName field if non-nil, zero value otherwise.

### GetYomiGivenNameOk

`func (o *MicrosoftGraphContact) GetYomiGivenNameOk() (*string, bool)`

GetYomiGivenNameOk returns a tuple with the YomiGivenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYomiGivenName

`func (o *MicrosoftGraphContact) SetYomiGivenName(v string)`

SetYomiGivenName sets YomiGivenName field to given value.

### HasYomiGivenName

`func (o *MicrosoftGraphContact) HasYomiGivenName() bool`

HasYomiGivenName returns a boolean if a field has been set.

### SetYomiGivenNameNil

`func (o *MicrosoftGraphContact) SetYomiGivenNameNil(b bool)`

 SetYomiGivenNameNil sets the value for YomiGivenName to be an explicit nil

### UnsetYomiGivenName
`func (o *MicrosoftGraphContact) UnsetYomiGivenName()`

UnsetYomiGivenName ensures that no value is present for YomiGivenName, not even an explicit nil
### GetYomiSurname

`func (o *MicrosoftGraphContact) GetYomiSurname() string`

GetYomiSurname returns the YomiSurname field if non-nil, zero value otherwise.

### GetYomiSurnameOk

`func (o *MicrosoftGraphContact) GetYomiSurnameOk() (*string, bool)`

GetYomiSurnameOk returns a tuple with the YomiSurname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYomiSurname

`func (o *MicrosoftGraphContact) SetYomiSurname(v string)`

SetYomiSurname sets YomiSurname field to given value.

### HasYomiSurname

`func (o *MicrosoftGraphContact) HasYomiSurname() bool`

HasYomiSurname returns a boolean if a field has been set.

### SetYomiSurnameNil

`func (o *MicrosoftGraphContact) SetYomiSurnameNil(b bool)`

 SetYomiSurnameNil sets the value for YomiSurname to be an explicit nil

### UnsetYomiSurname
`func (o *MicrosoftGraphContact) UnsetYomiSurname()`

UnsetYomiSurname ensures that no value is present for YomiSurname, not even an explicit nil
### GetExtensions

`func (o *MicrosoftGraphContact) GetExtensions() []MicrosoftGraphExtension`

GetExtensions returns the Extensions field if non-nil, zero value otherwise.

### GetExtensionsOk

`func (o *MicrosoftGraphContact) GetExtensionsOk() (*[]MicrosoftGraphExtension, bool)`

GetExtensionsOk returns a tuple with the Extensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensions

`func (o *MicrosoftGraphContact) SetExtensions(v []MicrosoftGraphExtension)`

SetExtensions sets Extensions field to given value.

### HasExtensions

`func (o *MicrosoftGraphContact) HasExtensions() bool`

HasExtensions returns a boolean if a field has been set.

### GetMultiValueExtendedProperties

`func (o *MicrosoftGraphContact) GetMultiValueExtendedProperties() []MicrosoftGraphMultiValueLegacyExtendedProperty`

GetMultiValueExtendedProperties returns the MultiValueExtendedProperties field if non-nil, zero value otherwise.

### GetMultiValueExtendedPropertiesOk

`func (o *MicrosoftGraphContact) GetMultiValueExtendedPropertiesOk() (*[]MicrosoftGraphMultiValueLegacyExtendedProperty, bool)`

GetMultiValueExtendedPropertiesOk returns a tuple with the MultiValueExtendedProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiValueExtendedProperties

`func (o *MicrosoftGraphContact) SetMultiValueExtendedProperties(v []MicrosoftGraphMultiValueLegacyExtendedProperty)`

SetMultiValueExtendedProperties sets MultiValueExtendedProperties field to given value.

### HasMultiValueExtendedProperties

`func (o *MicrosoftGraphContact) HasMultiValueExtendedProperties() bool`

HasMultiValueExtendedProperties returns a boolean if a field has been set.

### GetPhoto

`func (o *MicrosoftGraphContact) GetPhoto() AnyOfmicrosoftGraphProfilePhoto`

GetPhoto returns the Photo field if non-nil, zero value otherwise.

### GetPhotoOk

`func (o *MicrosoftGraphContact) GetPhotoOk() (*AnyOfmicrosoftGraphProfilePhoto, bool)`

GetPhotoOk returns a tuple with the Photo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoto

`func (o *MicrosoftGraphContact) SetPhoto(v AnyOfmicrosoftGraphProfilePhoto)`

SetPhoto sets Photo field to given value.

### HasPhoto

`func (o *MicrosoftGraphContact) HasPhoto() bool`

HasPhoto returns a boolean if a field has been set.

### SetPhotoNil

`func (o *MicrosoftGraphContact) SetPhotoNil(b bool)`

 SetPhotoNil sets the value for Photo to be an explicit nil

### UnsetPhoto
`func (o *MicrosoftGraphContact) UnsetPhoto()`

UnsetPhoto ensures that no value is present for Photo, not even an explicit nil
### GetSingleValueExtendedProperties

`func (o *MicrosoftGraphContact) GetSingleValueExtendedProperties() []MicrosoftGraphSingleValueLegacyExtendedProperty`

GetSingleValueExtendedProperties returns the SingleValueExtendedProperties field if non-nil, zero value otherwise.

### GetSingleValueExtendedPropertiesOk

`func (o *MicrosoftGraphContact) GetSingleValueExtendedPropertiesOk() (*[]MicrosoftGraphSingleValueLegacyExtendedProperty, bool)`

GetSingleValueExtendedPropertiesOk returns a tuple with the SingleValueExtendedProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSingleValueExtendedProperties

`func (o *MicrosoftGraphContact) SetSingleValueExtendedProperties(v []MicrosoftGraphSingleValueLegacyExtendedProperty)`

SetSingleValueExtendedProperties sets SingleValueExtendedProperties field to given value.

### HasSingleValueExtendedProperties

`func (o *MicrosoftGraphContact) HasSingleValueExtendedProperties() bool`

HasSingleValueExtendedProperties returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


