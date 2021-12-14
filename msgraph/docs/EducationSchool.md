# EducationSchool

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to [**AnyOfmicrosoftGraphPhysicalAddress**](anyOf&lt;microsoft.graph.physicalAddress&gt;.md) | Address of the school. | [optional] 
**CreatedBy** | Pointer to [**AnyOfmicrosoftGraphIdentitySet**](anyOf&lt;microsoft.graph.identitySet&gt;.md) | Entity who created the school. | [optional] 
**ExternalId** | Pointer to **NullableString** | ID of school in syncing system. | [optional] 
**ExternalPrincipalId** | Pointer to **NullableString** | ID of principal in syncing system. | [optional] 
**Fax** | Pointer to **NullableString** |  | [optional] 
**HighestGrade** | Pointer to **NullableString** | Highest grade taught. | [optional] 
**LowestGrade** | Pointer to **NullableString** | Lowest grade taught. | [optional] 
**Phone** | Pointer to **NullableString** | Phone number of school. | [optional] 
**PrincipalEmail** | Pointer to **NullableString** | Email address of the principal. | [optional] 
**PrincipalName** | Pointer to **NullableString** | Name of the principal. | [optional] 
**SchoolNumber** | Pointer to **NullableString** | School Number. | [optional] 
**AdministrativeUnit** | Pointer to [**AnyOfmicrosoftGraphAdministrativeUnit**](anyOf&lt;microsoft.graph.administrativeUnit&gt;.md) | The underlying administrativeUnit for this school. | [optional] 
**Classes** | Pointer to [**[]MicrosoftGraphEducationClass**](MicrosoftGraphEducationClass.md) | Classes taught at the school. Nullable. | [optional] 
**Users** | Pointer to [**[]MicrosoftGraphEducationUser**](MicrosoftGraphEducationUser.md) | Users in the school. Nullable. | [optional] 

## Methods

### NewEducationSchool

`func NewEducationSchool() *EducationSchool`

NewEducationSchool instantiates a new EducationSchool object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEducationSchoolWithDefaults

`func NewEducationSchoolWithDefaults() *EducationSchool`

NewEducationSchoolWithDefaults instantiates a new EducationSchool object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *EducationSchool) GetAddress() AnyOfmicrosoftGraphPhysicalAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *EducationSchool) GetAddressOk() (*AnyOfmicrosoftGraphPhysicalAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *EducationSchool) SetAddress(v AnyOfmicrosoftGraphPhysicalAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *EducationSchool) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### SetAddressNil

`func (o *EducationSchool) SetAddressNil(b bool)`

 SetAddressNil sets the value for Address to be an explicit nil

### UnsetAddress
`func (o *EducationSchool) UnsetAddress()`

UnsetAddress ensures that no value is present for Address, not even an explicit nil
### GetCreatedBy

`func (o *EducationSchool) GetCreatedBy() AnyOfmicrosoftGraphIdentitySet`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *EducationSchool) GetCreatedByOk() (*AnyOfmicrosoftGraphIdentitySet, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *EducationSchool) SetCreatedBy(v AnyOfmicrosoftGraphIdentitySet)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *EducationSchool) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedByNil

`func (o *EducationSchool) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *EducationSchool) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
### GetExternalId

`func (o *EducationSchool) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *EducationSchool) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *EducationSchool) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *EducationSchool) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *EducationSchool) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *EducationSchool) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetExternalPrincipalId

`func (o *EducationSchool) GetExternalPrincipalId() string`

GetExternalPrincipalId returns the ExternalPrincipalId field if non-nil, zero value otherwise.

### GetExternalPrincipalIdOk

`func (o *EducationSchool) GetExternalPrincipalIdOk() (*string, bool)`

GetExternalPrincipalIdOk returns a tuple with the ExternalPrincipalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalPrincipalId

`func (o *EducationSchool) SetExternalPrincipalId(v string)`

SetExternalPrincipalId sets ExternalPrincipalId field to given value.

### HasExternalPrincipalId

`func (o *EducationSchool) HasExternalPrincipalId() bool`

HasExternalPrincipalId returns a boolean if a field has been set.

### SetExternalPrincipalIdNil

`func (o *EducationSchool) SetExternalPrincipalIdNil(b bool)`

 SetExternalPrincipalIdNil sets the value for ExternalPrincipalId to be an explicit nil

### UnsetExternalPrincipalId
`func (o *EducationSchool) UnsetExternalPrincipalId()`

UnsetExternalPrincipalId ensures that no value is present for ExternalPrincipalId, not even an explicit nil
### GetFax

`func (o *EducationSchool) GetFax() string`

GetFax returns the Fax field if non-nil, zero value otherwise.

### GetFaxOk

`func (o *EducationSchool) GetFaxOk() (*string, bool)`

GetFaxOk returns a tuple with the Fax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFax

`func (o *EducationSchool) SetFax(v string)`

SetFax sets Fax field to given value.

### HasFax

`func (o *EducationSchool) HasFax() bool`

HasFax returns a boolean if a field has been set.

### SetFaxNil

`func (o *EducationSchool) SetFaxNil(b bool)`

 SetFaxNil sets the value for Fax to be an explicit nil

### UnsetFax
`func (o *EducationSchool) UnsetFax()`

UnsetFax ensures that no value is present for Fax, not even an explicit nil
### GetHighestGrade

`func (o *EducationSchool) GetHighestGrade() string`

GetHighestGrade returns the HighestGrade field if non-nil, zero value otherwise.

### GetHighestGradeOk

`func (o *EducationSchool) GetHighestGradeOk() (*string, bool)`

GetHighestGradeOk returns a tuple with the HighestGrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighestGrade

`func (o *EducationSchool) SetHighestGrade(v string)`

SetHighestGrade sets HighestGrade field to given value.

### HasHighestGrade

`func (o *EducationSchool) HasHighestGrade() bool`

HasHighestGrade returns a boolean if a field has been set.

### SetHighestGradeNil

`func (o *EducationSchool) SetHighestGradeNil(b bool)`

 SetHighestGradeNil sets the value for HighestGrade to be an explicit nil

### UnsetHighestGrade
`func (o *EducationSchool) UnsetHighestGrade()`

UnsetHighestGrade ensures that no value is present for HighestGrade, not even an explicit nil
### GetLowestGrade

`func (o *EducationSchool) GetLowestGrade() string`

GetLowestGrade returns the LowestGrade field if non-nil, zero value otherwise.

### GetLowestGradeOk

`func (o *EducationSchool) GetLowestGradeOk() (*string, bool)`

GetLowestGradeOk returns a tuple with the LowestGrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowestGrade

`func (o *EducationSchool) SetLowestGrade(v string)`

SetLowestGrade sets LowestGrade field to given value.

### HasLowestGrade

`func (o *EducationSchool) HasLowestGrade() bool`

HasLowestGrade returns a boolean if a field has been set.

### SetLowestGradeNil

`func (o *EducationSchool) SetLowestGradeNil(b bool)`

 SetLowestGradeNil sets the value for LowestGrade to be an explicit nil

### UnsetLowestGrade
`func (o *EducationSchool) UnsetLowestGrade()`

UnsetLowestGrade ensures that no value is present for LowestGrade, not even an explicit nil
### GetPhone

`func (o *EducationSchool) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *EducationSchool) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *EducationSchool) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *EducationSchool) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### SetPhoneNil

`func (o *EducationSchool) SetPhoneNil(b bool)`

 SetPhoneNil sets the value for Phone to be an explicit nil

### UnsetPhone
`func (o *EducationSchool) UnsetPhone()`

UnsetPhone ensures that no value is present for Phone, not even an explicit nil
### GetPrincipalEmail

`func (o *EducationSchool) GetPrincipalEmail() string`

GetPrincipalEmail returns the PrincipalEmail field if non-nil, zero value otherwise.

### GetPrincipalEmailOk

`func (o *EducationSchool) GetPrincipalEmailOk() (*string, bool)`

GetPrincipalEmailOk returns a tuple with the PrincipalEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipalEmail

`func (o *EducationSchool) SetPrincipalEmail(v string)`

SetPrincipalEmail sets PrincipalEmail field to given value.

### HasPrincipalEmail

`func (o *EducationSchool) HasPrincipalEmail() bool`

HasPrincipalEmail returns a boolean if a field has been set.

### SetPrincipalEmailNil

`func (o *EducationSchool) SetPrincipalEmailNil(b bool)`

 SetPrincipalEmailNil sets the value for PrincipalEmail to be an explicit nil

### UnsetPrincipalEmail
`func (o *EducationSchool) UnsetPrincipalEmail()`

UnsetPrincipalEmail ensures that no value is present for PrincipalEmail, not even an explicit nil
### GetPrincipalName

`func (o *EducationSchool) GetPrincipalName() string`

GetPrincipalName returns the PrincipalName field if non-nil, zero value otherwise.

### GetPrincipalNameOk

`func (o *EducationSchool) GetPrincipalNameOk() (*string, bool)`

GetPrincipalNameOk returns a tuple with the PrincipalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipalName

`func (o *EducationSchool) SetPrincipalName(v string)`

SetPrincipalName sets PrincipalName field to given value.

### HasPrincipalName

`func (o *EducationSchool) HasPrincipalName() bool`

HasPrincipalName returns a boolean if a field has been set.

### SetPrincipalNameNil

`func (o *EducationSchool) SetPrincipalNameNil(b bool)`

 SetPrincipalNameNil sets the value for PrincipalName to be an explicit nil

### UnsetPrincipalName
`func (o *EducationSchool) UnsetPrincipalName()`

UnsetPrincipalName ensures that no value is present for PrincipalName, not even an explicit nil
### GetSchoolNumber

`func (o *EducationSchool) GetSchoolNumber() string`

GetSchoolNumber returns the SchoolNumber field if non-nil, zero value otherwise.

### GetSchoolNumberOk

`func (o *EducationSchool) GetSchoolNumberOk() (*string, bool)`

GetSchoolNumberOk returns a tuple with the SchoolNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchoolNumber

`func (o *EducationSchool) SetSchoolNumber(v string)`

SetSchoolNumber sets SchoolNumber field to given value.

### HasSchoolNumber

`func (o *EducationSchool) HasSchoolNumber() bool`

HasSchoolNumber returns a boolean if a field has been set.

### SetSchoolNumberNil

`func (o *EducationSchool) SetSchoolNumberNil(b bool)`

 SetSchoolNumberNil sets the value for SchoolNumber to be an explicit nil

### UnsetSchoolNumber
`func (o *EducationSchool) UnsetSchoolNumber()`

UnsetSchoolNumber ensures that no value is present for SchoolNumber, not even an explicit nil
### GetAdministrativeUnit

`func (o *EducationSchool) GetAdministrativeUnit() AnyOfmicrosoftGraphAdministrativeUnit`

GetAdministrativeUnit returns the AdministrativeUnit field if non-nil, zero value otherwise.

### GetAdministrativeUnitOk

`func (o *EducationSchool) GetAdministrativeUnitOk() (*AnyOfmicrosoftGraphAdministrativeUnit, bool)`

GetAdministrativeUnitOk returns a tuple with the AdministrativeUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdministrativeUnit

`func (o *EducationSchool) SetAdministrativeUnit(v AnyOfmicrosoftGraphAdministrativeUnit)`

SetAdministrativeUnit sets AdministrativeUnit field to given value.

### HasAdministrativeUnit

`func (o *EducationSchool) HasAdministrativeUnit() bool`

HasAdministrativeUnit returns a boolean if a field has been set.

### SetAdministrativeUnitNil

`func (o *EducationSchool) SetAdministrativeUnitNil(b bool)`

 SetAdministrativeUnitNil sets the value for AdministrativeUnit to be an explicit nil

### UnsetAdministrativeUnit
`func (o *EducationSchool) UnsetAdministrativeUnit()`

UnsetAdministrativeUnit ensures that no value is present for AdministrativeUnit, not even an explicit nil
### GetClasses

`func (o *EducationSchool) GetClasses() []MicrosoftGraphEducationClass`

GetClasses returns the Classes field if non-nil, zero value otherwise.

### GetClassesOk

`func (o *EducationSchool) GetClassesOk() (*[]MicrosoftGraphEducationClass, bool)`

GetClassesOk returns a tuple with the Classes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClasses

`func (o *EducationSchool) SetClasses(v []MicrosoftGraphEducationClass)`

SetClasses sets Classes field to given value.

### HasClasses

`func (o *EducationSchool) HasClasses() bool`

HasClasses returns a boolean if a field has been set.

### GetUsers

`func (o *EducationSchool) GetUsers() []MicrosoftGraphEducationUser`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *EducationSchool) GetUsersOk() (*[]MicrosoftGraphEducationUser, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *EducationSchool) SetUsers(v []MicrosoftGraphEducationUser)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *EducationSchool) HasUsers() bool`

HasUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


