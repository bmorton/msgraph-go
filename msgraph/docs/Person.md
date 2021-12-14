# Person

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Birthday** | Pointer to **NullableString** | The person&#39;s birthday. | [optional] 
**CompanyName** | Pointer to **NullableString** | The name of the person&#39;s company. | [optional] 
**Department** | Pointer to **NullableString** | The person&#39;s department. | [optional] 
**DisplayName** | Pointer to **NullableString** | The person&#39;s display name. | [optional] 
**GivenName** | Pointer to **NullableString** | The person&#39;s given name. | [optional] 
**ImAddress** | Pointer to **NullableString** | The instant message voice over IP (VOIP) session initiation protocol (SIP) address for the user. Read-only. | [optional] 
**IsFavorite** | Pointer to **NullableBool** | true if the user has flagged this person as a favorite. | [optional] 
**JobTitle** | Pointer to **NullableString** | The person&#39;s job title. | [optional] 
**OfficeLocation** | Pointer to **NullableString** | The location of the person&#39;s office. | [optional] 
**PersonNotes** | Pointer to **NullableString** | Free-form notes that the user has taken about this person. | [optional] 
**PersonType** | Pointer to [**AnyOfmicrosoftGraphPersonType**](anyOf&lt;microsoft.graph.personType&gt;.md) | The type of person. | [optional] 
**Phones** | Pointer to [**[]AnyOfmicrosoftGraphPhone**](AnyOfmicrosoftGraphPhone.md) | The person&#39;s phone numbers. | [optional] 
**PostalAddresses** | Pointer to [**[]AnyOfmicrosoftGraphLocation**](AnyOfmicrosoftGraphLocation.md) | The person&#39;s addresses. | [optional] 
**Profession** | Pointer to **NullableString** | The person&#39;s profession. | [optional] 
**ScoredEmailAddresses** | Pointer to [**[]AnyOfmicrosoftGraphScoredEmailAddress**](AnyOfmicrosoftGraphScoredEmailAddress.md) | The person&#39;s email addresses. | [optional] 
**Surname** | Pointer to **NullableString** | The person&#39;s surname. | [optional] 
**UserPrincipalName** | Pointer to **NullableString** | The user principal name (UPN) of the person. The UPN is an Internet-style login name for the person based on the Internet standard RFC 822. By convention, this should map to the person&#39;s email name. The general format is alias@domain. | [optional] 
**Websites** | Pointer to [**[]AnyOfmicrosoftGraphWebsite**](AnyOfmicrosoftGraphWebsite.md) | The person&#39;s websites. | [optional] 
**YomiCompany** | Pointer to **NullableString** | The phonetic Japanese name of the person&#39;s company. | [optional] 

## Methods

### NewPerson

`func NewPerson() *Person`

NewPerson instantiates a new Person object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPersonWithDefaults

`func NewPersonWithDefaults() *Person`

NewPersonWithDefaults instantiates a new Person object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBirthday

`func (o *Person) GetBirthday() string`

GetBirthday returns the Birthday field if non-nil, zero value otherwise.

### GetBirthdayOk

`func (o *Person) GetBirthdayOk() (*string, bool)`

GetBirthdayOk returns a tuple with the Birthday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBirthday

`func (o *Person) SetBirthday(v string)`

SetBirthday sets Birthday field to given value.

### HasBirthday

`func (o *Person) HasBirthday() bool`

HasBirthday returns a boolean if a field has been set.

### SetBirthdayNil

`func (o *Person) SetBirthdayNil(b bool)`

 SetBirthdayNil sets the value for Birthday to be an explicit nil

### UnsetBirthday
`func (o *Person) UnsetBirthday()`

UnsetBirthday ensures that no value is present for Birthday, not even an explicit nil
### GetCompanyName

`func (o *Person) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *Person) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *Person) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *Person) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### SetCompanyNameNil

`func (o *Person) SetCompanyNameNil(b bool)`

 SetCompanyNameNil sets the value for CompanyName to be an explicit nil

### UnsetCompanyName
`func (o *Person) UnsetCompanyName()`

UnsetCompanyName ensures that no value is present for CompanyName, not even an explicit nil
### GetDepartment

`func (o *Person) GetDepartment() string`

GetDepartment returns the Department field if non-nil, zero value otherwise.

### GetDepartmentOk

`func (o *Person) GetDepartmentOk() (*string, bool)`

GetDepartmentOk returns a tuple with the Department field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartment

`func (o *Person) SetDepartment(v string)`

SetDepartment sets Department field to given value.

### HasDepartment

`func (o *Person) HasDepartment() bool`

HasDepartment returns a boolean if a field has been set.

### SetDepartmentNil

`func (o *Person) SetDepartmentNil(b bool)`

 SetDepartmentNil sets the value for Department to be an explicit nil

### UnsetDepartment
`func (o *Person) UnsetDepartment()`

UnsetDepartment ensures that no value is present for Department, not even an explicit nil
### GetDisplayName

`func (o *Person) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *Person) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *Person) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *Person) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *Person) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *Person) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetGivenName

`func (o *Person) GetGivenName() string`

GetGivenName returns the GivenName field if non-nil, zero value otherwise.

### GetGivenNameOk

`func (o *Person) GetGivenNameOk() (*string, bool)`

GetGivenNameOk returns a tuple with the GivenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGivenName

`func (o *Person) SetGivenName(v string)`

SetGivenName sets GivenName field to given value.

### HasGivenName

`func (o *Person) HasGivenName() bool`

HasGivenName returns a boolean if a field has been set.

### SetGivenNameNil

`func (o *Person) SetGivenNameNil(b bool)`

 SetGivenNameNil sets the value for GivenName to be an explicit nil

### UnsetGivenName
`func (o *Person) UnsetGivenName()`

UnsetGivenName ensures that no value is present for GivenName, not even an explicit nil
### GetImAddress

`func (o *Person) GetImAddress() string`

GetImAddress returns the ImAddress field if non-nil, zero value otherwise.

### GetImAddressOk

`func (o *Person) GetImAddressOk() (*string, bool)`

GetImAddressOk returns a tuple with the ImAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImAddress

`func (o *Person) SetImAddress(v string)`

SetImAddress sets ImAddress field to given value.

### HasImAddress

`func (o *Person) HasImAddress() bool`

HasImAddress returns a boolean if a field has been set.

### SetImAddressNil

`func (o *Person) SetImAddressNil(b bool)`

 SetImAddressNil sets the value for ImAddress to be an explicit nil

### UnsetImAddress
`func (o *Person) UnsetImAddress()`

UnsetImAddress ensures that no value is present for ImAddress, not even an explicit nil
### GetIsFavorite

`func (o *Person) GetIsFavorite() bool`

GetIsFavorite returns the IsFavorite field if non-nil, zero value otherwise.

### GetIsFavoriteOk

`func (o *Person) GetIsFavoriteOk() (*bool, bool)`

GetIsFavoriteOk returns a tuple with the IsFavorite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFavorite

`func (o *Person) SetIsFavorite(v bool)`

SetIsFavorite sets IsFavorite field to given value.

### HasIsFavorite

`func (o *Person) HasIsFavorite() bool`

HasIsFavorite returns a boolean if a field has been set.

### SetIsFavoriteNil

`func (o *Person) SetIsFavoriteNil(b bool)`

 SetIsFavoriteNil sets the value for IsFavorite to be an explicit nil

### UnsetIsFavorite
`func (o *Person) UnsetIsFavorite()`

UnsetIsFavorite ensures that no value is present for IsFavorite, not even an explicit nil
### GetJobTitle

`func (o *Person) GetJobTitle() string`

GetJobTitle returns the JobTitle field if non-nil, zero value otherwise.

### GetJobTitleOk

`func (o *Person) GetJobTitleOk() (*string, bool)`

GetJobTitleOk returns a tuple with the JobTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobTitle

`func (o *Person) SetJobTitle(v string)`

SetJobTitle sets JobTitle field to given value.

### HasJobTitle

`func (o *Person) HasJobTitle() bool`

HasJobTitle returns a boolean if a field has been set.

### SetJobTitleNil

`func (o *Person) SetJobTitleNil(b bool)`

 SetJobTitleNil sets the value for JobTitle to be an explicit nil

### UnsetJobTitle
`func (o *Person) UnsetJobTitle()`

UnsetJobTitle ensures that no value is present for JobTitle, not even an explicit nil
### GetOfficeLocation

`func (o *Person) GetOfficeLocation() string`

GetOfficeLocation returns the OfficeLocation field if non-nil, zero value otherwise.

### GetOfficeLocationOk

`func (o *Person) GetOfficeLocationOk() (*string, bool)`

GetOfficeLocationOk returns a tuple with the OfficeLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfficeLocation

`func (o *Person) SetOfficeLocation(v string)`

SetOfficeLocation sets OfficeLocation field to given value.

### HasOfficeLocation

`func (o *Person) HasOfficeLocation() bool`

HasOfficeLocation returns a boolean if a field has been set.

### SetOfficeLocationNil

`func (o *Person) SetOfficeLocationNil(b bool)`

 SetOfficeLocationNil sets the value for OfficeLocation to be an explicit nil

### UnsetOfficeLocation
`func (o *Person) UnsetOfficeLocation()`

UnsetOfficeLocation ensures that no value is present for OfficeLocation, not even an explicit nil
### GetPersonNotes

`func (o *Person) GetPersonNotes() string`

GetPersonNotes returns the PersonNotes field if non-nil, zero value otherwise.

### GetPersonNotesOk

`func (o *Person) GetPersonNotesOk() (*string, bool)`

GetPersonNotesOk returns a tuple with the PersonNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonNotes

`func (o *Person) SetPersonNotes(v string)`

SetPersonNotes sets PersonNotes field to given value.

### HasPersonNotes

`func (o *Person) HasPersonNotes() bool`

HasPersonNotes returns a boolean if a field has been set.

### SetPersonNotesNil

`func (o *Person) SetPersonNotesNil(b bool)`

 SetPersonNotesNil sets the value for PersonNotes to be an explicit nil

### UnsetPersonNotes
`func (o *Person) UnsetPersonNotes()`

UnsetPersonNotes ensures that no value is present for PersonNotes, not even an explicit nil
### GetPersonType

`func (o *Person) GetPersonType() AnyOfmicrosoftGraphPersonType`

GetPersonType returns the PersonType field if non-nil, zero value otherwise.

### GetPersonTypeOk

`func (o *Person) GetPersonTypeOk() (*AnyOfmicrosoftGraphPersonType, bool)`

GetPersonTypeOk returns a tuple with the PersonType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonType

`func (o *Person) SetPersonType(v AnyOfmicrosoftGraphPersonType)`

SetPersonType sets PersonType field to given value.

### HasPersonType

`func (o *Person) HasPersonType() bool`

HasPersonType returns a boolean if a field has been set.

### SetPersonTypeNil

`func (o *Person) SetPersonTypeNil(b bool)`

 SetPersonTypeNil sets the value for PersonType to be an explicit nil

### UnsetPersonType
`func (o *Person) UnsetPersonType()`

UnsetPersonType ensures that no value is present for PersonType, not even an explicit nil
### GetPhones

`func (o *Person) GetPhones() []*AnyOfmicrosoftGraphPhone`

GetPhones returns the Phones field if non-nil, zero value otherwise.

### GetPhonesOk

`func (o *Person) GetPhonesOk() (*[]*AnyOfmicrosoftGraphPhone, bool)`

GetPhonesOk returns a tuple with the Phones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhones

`func (o *Person) SetPhones(v []*AnyOfmicrosoftGraphPhone)`

SetPhones sets Phones field to given value.

### HasPhones

`func (o *Person) HasPhones() bool`

HasPhones returns a boolean if a field has been set.

### GetPostalAddresses

`func (o *Person) GetPostalAddresses() []*AnyOfmicrosoftGraphLocation`

GetPostalAddresses returns the PostalAddresses field if non-nil, zero value otherwise.

### GetPostalAddressesOk

`func (o *Person) GetPostalAddressesOk() (*[]*AnyOfmicrosoftGraphLocation, bool)`

GetPostalAddressesOk returns a tuple with the PostalAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalAddresses

`func (o *Person) SetPostalAddresses(v []*AnyOfmicrosoftGraphLocation)`

SetPostalAddresses sets PostalAddresses field to given value.

### HasPostalAddresses

`func (o *Person) HasPostalAddresses() bool`

HasPostalAddresses returns a boolean if a field has been set.

### GetProfession

`func (o *Person) GetProfession() string`

GetProfession returns the Profession field if non-nil, zero value otherwise.

### GetProfessionOk

`func (o *Person) GetProfessionOk() (*string, bool)`

GetProfessionOk returns a tuple with the Profession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfession

`func (o *Person) SetProfession(v string)`

SetProfession sets Profession field to given value.

### HasProfession

`func (o *Person) HasProfession() bool`

HasProfession returns a boolean if a field has been set.

### SetProfessionNil

`func (o *Person) SetProfessionNil(b bool)`

 SetProfessionNil sets the value for Profession to be an explicit nil

### UnsetProfession
`func (o *Person) UnsetProfession()`

UnsetProfession ensures that no value is present for Profession, not even an explicit nil
### GetScoredEmailAddresses

`func (o *Person) GetScoredEmailAddresses() []*AnyOfmicrosoftGraphScoredEmailAddress`

GetScoredEmailAddresses returns the ScoredEmailAddresses field if non-nil, zero value otherwise.

### GetScoredEmailAddressesOk

`func (o *Person) GetScoredEmailAddressesOk() (*[]*AnyOfmicrosoftGraphScoredEmailAddress, bool)`

GetScoredEmailAddressesOk returns a tuple with the ScoredEmailAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScoredEmailAddresses

`func (o *Person) SetScoredEmailAddresses(v []*AnyOfmicrosoftGraphScoredEmailAddress)`

SetScoredEmailAddresses sets ScoredEmailAddresses field to given value.

### HasScoredEmailAddresses

`func (o *Person) HasScoredEmailAddresses() bool`

HasScoredEmailAddresses returns a boolean if a field has been set.

### GetSurname

`func (o *Person) GetSurname() string`

GetSurname returns the Surname field if non-nil, zero value otherwise.

### GetSurnameOk

`func (o *Person) GetSurnameOk() (*string, bool)`

GetSurnameOk returns a tuple with the Surname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurname

`func (o *Person) SetSurname(v string)`

SetSurname sets Surname field to given value.

### HasSurname

`func (o *Person) HasSurname() bool`

HasSurname returns a boolean if a field has been set.

### SetSurnameNil

`func (o *Person) SetSurnameNil(b bool)`

 SetSurnameNil sets the value for Surname to be an explicit nil

### UnsetSurname
`func (o *Person) UnsetSurname()`

UnsetSurname ensures that no value is present for Surname, not even an explicit nil
### GetUserPrincipalName

`func (o *Person) GetUserPrincipalName() string`

GetUserPrincipalName returns the UserPrincipalName field if non-nil, zero value otherwise.

### GetUserPrincipalNameOk

`func (o *Person) GetUserPrincipalNameOk() (*string, bool)`

GetUserPrincipalNameOk returns a tuple with the UserPrincipalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPrincipalName

`func (o *Person) SetUserPrincipalName(v string)`

SetUserPrincipalName sets UserPrincipalName field to given value.

### HasUserPrincipalName

`func (o *Person) HasUserPrincipalName() bool`

HasUserPrincipalName returns a boolean if a field has been set.

### SetUserPrincipalNameNil

`func (o *Person) SetUserPrincipalNameNil(b bool)`

 SetUserPrincipalNameNil sets the value for UserPrincipalName to be an explicit nil

### UnsetUserPrincipalName
`func (o *Person) UnsetUserPrincipalName()`

UnsetUserPrincipalName ensures that no value is present for UserPrincipalName, not even an explicit nil
### GetWebsites

`func (o *Person) GetWebsites() []*AnyOfmicrosoftGraphWebsite`

GetWebsites returns the Websites field if non-nil, zero value otherwise.

### GetWebsitesOk

`func (o *Person) GetWebsitesOk() (*[]*AnyOfmicrosoftGraphWebsite, bool)`

GetWebsitesOk returns a tuple with the Websites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsites

`func (o *Person) SetWebsites(v []*AnyOfmicrosoftGraphWebsite)`

SetWebsites sets Websites field to given value.

### HasWebsites

`func (o *Person) HasWebsites() bool`

HasWebsites returns a boolean if a field has been set.

### GetYomiCompany

`func (o *Person) GetYomiCompany() string`

GetYomiCompany returns the YomiCompany field if non-nil, zero value otherwise.

### GetYomiCompanyOk

`func (o *Person) GetYomiCompanyOk() (*string, bool)`

GetYomiCompanyOk returns a tuple with the YomiCompany field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYomiCompany

`func (o *Person) SetYomiCompany(v string)`

SetYomiCompany sets YomiCompany field to given value.

### HasYomiCompany

`func (o *Person) HasYomiCompany() bool`

HasYomiCompany returns a boolean if a field has been set.

### SetYomiCompanyNil

`func (o *Person) SetYomiCompanyNil(b bool)`

 SetYomiCompanyNil sets the value for YomiCompany to be an explicit nil

### UnsetYomiCompany
`func (o *Person) UnsetYomiCompany()`

UnsetYomiCompany ensures that no value is present for YomiCompany, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


