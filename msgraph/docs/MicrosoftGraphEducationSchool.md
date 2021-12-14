# MicrosoftGraphEducationSchool

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**Description** | Pointer to **NullableString** | Organization description. | [optional] 
**DisplayName** | Pointer to **string** | Organization display name. | [optional] 
**ExternalSource** | Pointer to [**AnyOfmicrosoftGraphEducationExternalSource**](anyOf&lt;microsoft.graph.educationExternalSource&gt;.md) | Source where this organization was created from. Possible values are: sis, manual. | [optional] 
**ExternalSourceDetail** | Pointer to **NullableString** | The name of the external source this resources was generated from. | [optional] 
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

### NewMicrosoftGraphEducationSchool

`func NewMicrosoftGraphEducationSchool() *MicrosoftGraphEducationSchool`

NewMicrosoftGraphEducationSchool instantiates a new MicrosoftGraphEducationSchool object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphEducationSchoolWithDefaults

`func NewMicrosoftGraphEducationSchoolWithDefaults() *MicrosoftGraphEducationSchool`

NewMicrosoftGraphEducationSchoolWithDefaults instantiates a new MicrosoftGraphEducationSchool object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphEducationSchool) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphEducationSchool) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphEducationSchool) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphEducationSchool) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDescription

`func (o *MicrosoftGraphEducationSchool) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MicrosoftGraphEducationSchool) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MicrosoftGraphEducationSchool) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MicrosoftGraphEducationSchool) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *MicrosoftGraphEducationSchool) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *MicrosoftGraphEducationSchool) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDisplayName

`func (o *MicrosoftGraphEducationSchool) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphEducationSchool) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *MicrosoftGraphEducationSchool) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *MicrosoftGraphEducationSchool) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetExternalSource

`func (o *MicrosoftGraphEducationSchool) GetExternalSource() AnyOfmicrosoftGraphEducationExternalSource`

GetExternalSource returns the ExternalSource field if non-nil, zero value otherwise.

### GetExternalSourceOk

`func (o *MicrosoftGraphEducationSchool) GetExternalSourceOk() (*AnyOfmicrosoftGraphEducationExternalSource, bool)`

GetExternalSourceOk returns a tuple with the ExternalSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalSource

`func (o *MicrosoftGraphEducationSchool) SetExternalSource(v AnyOfmicrosoftGraphEducationExternalSource)`

SetExternalSource sets ExternalSource field to given value.

### HasExternalSource

`func (o *MicrosoftGraphEducationSchool) HasExternalSource() bool`

HasExternalSource returns a boolean if a field has been set.

### SetExternalSourceNil

`func (o *MicrosoftGraphEducationSchool) SetExternalSourceNil(b bool)`

 SetExternalSourceNil sets the value for ExternalSource to be an explicit nil

### UnsetExternalSource
`func (o *MicrosoftGraphEducationSchool) UnsetExternalSource()`

UnsetExternalSource ensures that no value is present for ExternalSource, not even an explicit nil
### GetExternalSourceDetail

`func (o *MicrosoftGraphEducationSchool) GetExternalSourceDetail() string`

GetExternalSourceDetail returns the ExternalSourceDetail field if non-nil, zero value otherwise.

### GetExternalSourceDetailOk

`func (o *MicrosoftGraphEducationSchool) GetExternalSourceDetailOk() (*string, bool)`

GetExternalSourceDetailOk returns a tuple with the ExternalSourceDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalSourceDetail

`func (o *MicrosoftGraphEducationSchool) SetExternalSourceDetail(v string)`

SetExternalSourceDetail sets ExternalSourceDetail field to given value.

### HasExternalSourceDetail

`func (o *MicrosoftGraphEducationSchool) HasExternalSourceDetail() bool`

HasExternalSourceDetail returns a boolean if a field has been set.

### SetExternalSourceDetailNil

`func (o *MicrosoftGraphEducationSchool) SetExternalSourceDetailNil(b bool)`

 SetExternalSourceDetailNil sets the value for ExternalSourceDetail to be an explicit nil

### UnsetExternalSourceDetail
`func (o *MicrosoftGraphEducationSchool) UnsetExternalSourceDetail()`

UnsetExternalSourceDetail ensures that no value is present for ExternalSourceDetail, not even an explicit nil
### GetAddress

`func (o *MicrosoftGraphEducationSchool) GetAddress() AnyOfmicrosoftGraphPhysicalAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *MicrosoftGraphEducationSchool) GetAddressOk() (*AnyOfmicrosoftGraphPhysicalAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *MicrosoftGraphEducationSchool) SetAddress(v AnyOfmicrosoftGraphPhysicalAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *MicrosoftGraphEducationSchool) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### SetAddressNil

`func (o *MicrosoftGraphEducationSchool) SetAddressNil(b bool)`

 SetAddressNil sets the value for Address to be an explicit nil

### UnsetAddress
`func (o *MicrosoftGraphEducationSchool) UnsetAddress()`

UnsetAddress ensures that no value is present for Address, not even an explicit nil
### GetCreatedBy

`func (o *MicrosoftGraphEducationSchool) GetCreatedBy() AnyOfmicrosoftGraphIdentitySet`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *MicrosoftGraphEducationSchool) GetCreatedByOk() (*AnyOfmicrosoftGraphIdentitySet, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *MicrosoftGraphEducationSchool) SetCreatedBy(v AnyOfmicrosoftGraphIdentitySet)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *MicrosoftGraphEducationSchool) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedByNil

`func (o *MicrosoftGraphEducationSchool) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *MicrosoftGraphEducationSchool) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
### GetExternalId

`func (o *MicrosoftGraphEducationSchool) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *MicrosoftGraphEducationSchool) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *MicrosoftGraphEducationSchool) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *MicrosoftGraphEducationSchool) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *MicrosoftGraphEducationSchool) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *MicrosoftGraphEducationSchool) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetExternalPrincipalId

`func (o *MicrosoftGraphEducationSchool) GetExternalPrincipalId() string`

GetExternalPrincipalId returns the ExternalPrincipalId field if non-nil, zero value otherwise.

### GetExternalPrincipalIdOk

`func (o *MicrosoftGraphEducationSchool) GetExternalPrincipalIdOk() (*string, bool)`

GetExternalPrincipalIdOk returns a tuple with the ExternalPrincipalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalPrincipalId

`func (o *MicrosoftGraphEducationSchool) SetExternalPrincipalId(v string)`

SetExternalPrincipalId sets ExternalPrincipalId field to given value.

### HasExternalPrincipalId

`func (o *MicrosoftGraphEducationSchool) HasExternalPrincipalId() bool`

HasExternalPrincipalId returns a boolean if a field has been set.

### SetExternalPrincipalIdNil

`func (o *MicrosoftGraphEducationSchool) SetExternalPrincipalIdNil(b bool)`

 SetExternalPrincipalIdNil sets the value for ExternalPrincipalId to be an explicit nil

### UnsetExternalPrincipalId
`func (o *MicrosoftGraphEducationSchool) UnsetExternalPrincipalId()`

UnsetExternalPrincipalId ensures that no value is present for ExternalPrincipalId, not even an explicit nil
### GetFax

`func (o *MicrosoftGraphEducationSchool) GetFax() string`

GetFax returns the Fax field if non-nil, zero value otherwise.

### GetFaxOk

`func (o *MicrosoftGraphEducationSchool) GetFaxOk() (*string, bool)`

GetFaxOk returns a tuple with the Fax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFax

`func (o *MicrosoftGraphEducationSchool) SetFax(v string)`

SetFax sets Fax field to given value.

### HasFax

`func (o *MicrosoftGraphEducationSchool) HasFax() bool`

HasFax returns a boolean if a field has been set.

### SetFaxNil

`func (o *MicrosoftGraphEducationSchool) SetFaxNil(b bool)`

 SetFaxNil sets the value for Fax to be an explicit nil

### UnsetFax
`func (o *MicrosoftGraphEducationSchool) UnsetFax()`

UnsetFax ensures that no value is present for Fax, not even an explicit nil
### GetHighestGrade

`func (o *MicrosoftGraphEducationSchool) GetHighestGrade() string`

GetHighestGrade returns the HighestGrade field if non-nil, zero value otherwise.

### GetHighestGradeOk

`func (o *MicrosoftGraphEducationSchool) GetHighestGradeOk() (*string, bool)`

GetHighestGradeOk returns a tuple with the HighestGrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighestGrade

`func (o *MicrosoftGraphEducationSchool) SetHighestGrade(v string)`

SetHighestGrade sets HighestGrade field to given value.

### HasHighestGrade

`func (o *MicrosoftGraphEducationSchool) HasHighestGrade() bool`

HasHighestGrade returns a boolean if a field has been set.

### SetHighestGradeNil

`func (o *MicrosoftGraphEducationSchool) SetHighestGradeNil(b bool)`

 SetHighestGradeNil sets the value for HighestGrade to be an explicit nil

### UnsetHighestGrade
`func (o *MicrosoftGraphEducationSchool) UnsetHighestGrade()`

UnsetHighestGrade ensures that no value is present for HighestGrade, not even an explicit nil
### GetLowestGrade

`func (o *MicrosoftGraphEducationSchool) GetLowestGrade() string`

GetLowestGrade returns the LowestGrade field if non-nil, zero value otherwise.

### GetLowestGradeOk

`func (o *MicrosoftGraphEducationSchool) GetLowestGradeOk() (*string, bool)`

GetLowestGradeOk returns a tuple with the LowestGrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowestGrade

`func (o *MicrosoftGraphEducationSchool) SetLowestGrade(v string)`

SetLowestGrade sets LowestGrade field to given value.

### HasLowestGrade

`func (o *MicrosoftGraphEducationSchool) HasLowestGrade() bool`

HasLowestGrade returns a boolean if a field has been set.

### SetLowestGradeNil

`func (o *MicrosoftGraphEducationSchool) SetLowestGradeNil(b bool)`

 SetLowestGradeNil sets the value for LowestGrade to be an explicit nil

### UnsetLowestGrade
`func (o *MicrosoftGraphEducationSchool) UnsetLowestGrade()`

UnsetLowestGrade ensures that no value is present for LowestGrade, not even an explicit nil
### GetPhone

`func (o *MicrosoftGraphEducationSchool) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *MicrosoftGraphEducationSchool) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *MicrosoftGraphEducationSchool) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *MicrosoftGraphEducationSchool) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### SetPhoneNil

`func (o *MicrosoftGraphEducationSchool) SetPhoneNil(b bool)`

 SetPhoneNil sets the value for Phone to be an explicit nil

### UnsetPhone
`func (o *MicrosoftGraphEducationSchool) UnsetPhone()`

UnsetPhone ensures that no value is present for Phone, not even an explicit nil
### GetPrincipalEmail

`func (o *MicrosoftGraphEducationSchool) GetPrincipalEmail() string`

GetPrincipalEmail returns the PrincipalEmail field if non-nil, zero value otherwise.

### GetPrincipalEmailOk

`func (o *MicrosoftGraphEducationSchool) GetPrincipalEmailOk() (*string, bool)`

GetPrincipalEmailOk returns a tuple with the PrincipalEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipalEmail

`func (o *MicrosoftGraphEducationSchool) SetPrincipalEmail(v string)`

SetPrincipalEmail sets PrincipalEmail field to given value.

### HasPrincipalEmail

`func (o *MicrosoftGraphEducationSchool) HasPrincipalEmail() bool`

HasPrincipalEmail returns a boolean if a field has been set.

### SetPrincipalEmailNil

`func (o *MicrosoftGraphEducationSchool) SetPrincipalEmailNil(b bool)`

 SetPrincipalEmailNil sets the value for PrincipalEmail to be an explicit nil

### UnsetPrincipalEmail
`func (o *MicrosoftGraphEducationSchool) UnsetPrincipalEmail()`

UnsetPrincipalEmail ensures that no value is present for PrincipalEmail, not even an explicit nil
### GetPrincipalName

`func (o *MicrosoftGraphEducationSchool) GetPrincipalName() string`

GetPrincipalName returns the PrincipalName field if non-nil, zero value otherwise.

### GetPrincipalNameOk

`func (o *MicrosoftGraphEducationSchool) GetPrincipalNameOk() (*string, bool)`

GetPrincipalNameOk returns a tuple with the PrincipalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipalName

`func (o *MicrosoftGraphEducationSchool) SetPrincipalName(v string)`

SetPrincipalName sets PrincipalName field to given value.

### HasPrincipalName

`func (o *MicrosoftGraphEducationSchool) HasPrincipalName() bool`

HasPrincipalName returns a boolean if a field has been set.

### SetPrincipalNameNil

`func (o *MicrosoftGraphEducationSchool) SetPrincipalNameNil(b bool)`

 SetPrincipalNameNil sets the value for PrincipalName to be an explicit nil

### UnsetPrincipalName
`func (o *MicrosoftGraphEducationSchool) UnsetPrincipalName()`

UnsetPrincipalName ensures that no value is present for PrincipalName, not even an explicit nil
### GetSchoolNumber

`func (o *MicrosoftGraphEducationSchool) GetSchoolNumber() string`

GetSchoolNumber returns the SchoolNumber field if non-nil, zero value otherwise.

### GetSchoolNumberOk

`func (o *MicrosoftGraphEducationSchool) GetSchoolNumberOk() (*string, bool)`

GetSchoolNumberOk returns a tuple with the SchoolNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchoolNumber

`func (o *MicrosoftGraphEducationSchool) SetSchoolNumber(v string)`

SetSchoolNumber sets SchoolNumber field to given value.

### HasSchoolNumber

`func (o *MicrosoftGraphEducationSchool) HasSchoolNumber() bool`

HasSchoolNumber returns a boolean if a field has been set.

### SetSchoolNumberNil

`func (o *MicrosoftGraphEducationSchool) SetSchoolNumberNil(b bool)`

 SetSchoolNumberNil sets the value for SchoolNumber to be an explicit nil

### UnsetSchoolNumber
`func (o *MicrosoftGraphEducationSchool) UnsetSchoolNumber()`

UnsetSchoolNumber ensures that no value is present for SchoolNumber, not even an explicit nil
### GetAdministrativeUnit

`func (o *MicrosoftGraphEducationSchool) GetAdministrativeUnit() AnyOfmicrosoftGraphAdministrativeUnit`

GetAdministrativeUnit returns the AdministrativeUnit field if non-nil, zero value otherwise.

### GetAdministrativeUnitOk

`func (o *MicrosoftGraphEducationSchool) GetAdministrativeUnitOk() (*AnyOfmicrosoftGraphAdministrativeUnit, bool)`

GetAdministrativeUnitOk returns a tuple with the AdministrativeUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdministrativeUnit

`func (o *MicrosoftGraphEducationSchool) SetAdministrativeUnit(v AnyOfmicrosoftGraphAdministrativeUnit)`

SetAdministrativeUnit sets AdministrativeUnit field to given value.

### HasAdministrativeUnit

`func (o *MicrosoftGraphEducationSchool) HasAdministrativeUnit() bool`

HasAdministrativeUnit returns a boolean if a field has been set.

### SetAdministrativeUnitNil

`func (o *MicrosoftGraphEducationSchool) SetAdministrativeUnitNil(b bool)`

 SetAdministrativeUnitNil sets the value for AdministrativeUnit to be an explicit nil

### UnsetAdministrativeUnit
`func (o *MicrosoftGraphEducationSchool) UnsetAdministrativeUnit()`

UnsetAdministrativeUnit ensures that no value is present for AdministrativeUnit, not even an explicit nil
### GetClasses

`func (o *MicrosoftGraphEducationSchool) GetClasses() []MicrosoftGraphEducationClass`

GetClasses returns the Classes field if non-nil, zero value otherwise.

### GetClassesOk

`func (o *MicrosoftGraphEducationSchool) GetClassesOk() (*[]MicrosoftGraphEducationClass, bool)`

GetClassesOk returns a tuple with the Classes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClasses

`func (o *MicrosoftGraphEducationSchool) SetClasses(v []MicrosoftGraphEducationClass)`

SetClasses sets Classes field to given value.

### HasClasses

`func (o *MicrosoftGraphEducationSchool) HasClasses() bool`

HasClasses returns a boolean if a field has been set.

### GetUsers

`func (o *MicrosoftGraphEducationSchool) GetUsers() []MicrosoftGraphEducationUser`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *MicrosoftGraphEducationSchool) GetUsersOk() (*[]MicrosoftGraphEducationUser, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *MicrosoftGraphEducationSchool) SetUsers(v []MicrosoftGraphEducationUser)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *MicrosoftGraphEducationSchool) HasUsers() bool`

HasUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


