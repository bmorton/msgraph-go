# MicrosoftGraphEducationClass

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**ClassCode** | Pointer to **NullableString** | Class code used by the school to identify the class. | [optional] 
**Course** | Pointer to [**AnyOfmicrosoftGraphEducationCourse**](anyOf&lt;microsoft.graph.educationCourse&gt;.md) | Course information for the class. | [optional] 
**CreatedBy** | Pointer to [**AnyOfmicrosoftGraphIdentitySet**](anyOf&lt;microsoft.graph.identitySet&gt;.md) | Entity who created the class | [optional] 
**Description** | Pointer to **NullableString** | Description of the class. | [optional] 
**DisplayName** | Pointer to **string** | Name of the class. | [optional] 
**ExternalId** | Pointer to **NullableString** | ID of the class from the syncing system. | [optional] 
**ExternalName** | Pointer to **NullableString** | Name of the class in the syncing system. | [optional] 
**ExternalSource** | Pointer to [**AnyOfmicrosoftGraphEducationExternalSource**](anyOf&lt;microsoft.graph.educationExternalSource&gt;.md) | How this class was created. Possible values are: sis, manual. | [optional] 
**ExternalSourceDetail** | Pointer to **NullableString** | The name of the external source this resources was generated from. | [optional] 
**Grade** | Pointer to **NullableString** | Grade level of the class. | [optional] 
**MailNickname** | Pointer to **string** | Mail name for sending email to all members, if this is enabled. | [optional] 
**Term** | Pointer to [**AnyOfmicrosoftGraphEducationTerm**](anyOf&lt;microsoft.graph.educationTerm&gt;.md) | Term for this class. | [optional] 
**AssignmentCategories** | Pointer to [**[]MicrosoftGraphEducationCategory**](MicrosoftGraphEducationCategory.md) |  | [optional] 
**AssignmentDefaults** | Pointer to [**AnyOfmicrosoftGraphEducationAssignmentDefaults**](anyOf&lt;microsoft.graph.educationAssignmentDefaults&gt;.md) |  | [optional] 
**Assignments** | Pointer to [**[]MicrosoftGraphEducationAssignment**](MicrosoftGraphEducationAssignment.md) | All assignments associated with this class. Nullable. | [optional] 
**AssignmentSettings** | Pointer to [**AnyOfmicrosoftGraphEducationAssignmentSettings**](anyOf&lt;microsoft.graph.educationAssignmentSettings&gt;.md) |  | [optional] 
**Group** | Pointer to [**AnyOfmicrosoftGraphGroup**](anyOf&lt;microsoft.graph.group&gt;.md) | The underlying Microsoft 365 group object. | [optional] 
**Members** | Pointer to [**[]MicrosoftGraphEducationUser**](MicrosoftGraphEducationUser.md) | All users in the class. Nullable. | [optional] 
**Schools** | Pointer to [**[]MicrosoftGraphEducationSchool**](MicrosoftGraphEducationSchool.md) | All schools that this class is associated with. Nullable. | [optional] 
**Teachers** | Pointer to [**[]MicrosoftGraphEducationUser**](MicrosoftGraphEducationUser.md) | All teachers in the class. Nullable. | [optional] 

## Methods

### NewMicrosoftGraphEducationClass

`func NewMicrosoftGraphEducationClass() *MicrosoftGraphEducationClass`

NewMicrosoftGraphEducationClass instantiates a new MicrosoftGraphEducationClass object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphEducationClassWithDefaults

`func NewMicrosoftGraphEducationClassWithDefaults() *MicrosoftGraphEducationClass`

NewMicrosoftGraphEducationClassWithDefaults instantiates a new MicrosoftGraphEducationClass object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphEducationClass) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphEducationClass) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphEducationClass) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphEducationClass) HasId() bool`

HasId returns a boolean if a field has been set.

### GetClassCode

`func (o *MicrosoftGraphEducationClass) GetClassCode() string`

GetClassCode returns the ClassCode field if non-nil, zero value otherwise.

### GetClassCodeOk

`func (o *MicrosoftGraphEducationClass) GetClassCodeOk() (*string, bool)`

GetClassCodeOk returns a tuple with the ClassCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassCode

`func (o *MicrosoftGraphEducationClass) SetClassCode(v string)`

SetClassCode sets ClassCode field to given value.

### HasClassCode

`func (o *MicrosoftGraphEducationClass) HasClassCode() bool`

HasClassCode returns a boolean if a field has been set.

### SetClassCodeNil

`func (o *MicrosoftGraphEducationClass) SetClassCodeNil(b bool)`

 SetClassCodeNil sets the value for ClassCode to be an explicit nil

### UnsetClassCode
`func (o *MicrosoftGraphEducationClass) UnsetClassCode()`

UnsetClassCode ensures that no value is present for ClassCode, not even an explicit nil
### GetCourse

`func (o *MicrosoftGraphEducationClass) GetCourse() AnyOfmicrosoftGraphEducationCourse`

GetCourse returns the Course field if non-nil, zero value otherwise.

### GetCourseOk

`func (o *MicrosoftGraphEducationClass) GetCourseOk() (*AnyOfmicrosoftGraphEducationCourse, bool)`

GetCourseOk returns a tuple with the Course field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCourse

`func (o *MicrosoftGraphEducationClass) SetCourse(v AnyOfmicrosoftGraphEducationCourse)`

SetCourse sets Course field to given value.

### HasCourse

`func (o *MicrosoftGraphEducationClass) HasCourse() bool`

HasCourse returns a boolean if a field has been set.

### SetCourseNil

`func (o *MicrosoftGraphEducationClass) SetCourseNil(b bool)`

 SetCourseNil sets the value for Course to be an explicit nil

### UnsetCourse
`func (o *MicrosoftGraphEducationClass) UnsetCourse()`

UnsetCourse ensures that no value is present for Course, not even an explicit nil
### GetCreatedBy

`func (o *MicrosoftGraphEducationClass) GetCreatedBy() AnyOfmicrosoftGraphIdentitySet`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *MicrosoftGraphEducationClass) GetCreatedByOk() (*AnyOfmicrosoftGraphIdentitySet, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *MicrosoftGraphEducationClass) SetCreatedBy(v AnyOfmicrosoftGraphIdentitySet)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *MicrosoftGraphEducationClass) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedByNil

`func (o *MicrosoftGraphEducationClass) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *MicrosoftGraphEducationClass) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
### GetDescription

`func (o *MicrosoftGraphEducationClass) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MicrosoftGraphEducationClass) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MicrosoftGraphEducationClass) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MicrosoftGraphEducationClass) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *MicrosoftGraphEducationClass) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *MicrosoftGraphEducationClass) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDisplayName

`func (o *MicrosoftGraphEducationClass) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphEducationClass) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *MicrosoftGraphEducationClass) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *MicrosoftGraphEducationClass) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetExternalId

`func (o *MicrosoftGraphEducationClass) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *MicrosoftGraphEducationClass) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *MicrosoftGraphEducationClass) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *MicrosoftGraphEducationClass) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *MicrosoftGraphEducationClass) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *MicrosoftGraphEducationClass) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetExternalName

`func (o *MicrosoftGraphEducationClass) GetExternalName() string`

GetExternalName returns the ExternalName field if non-nil, zero value otherwise.

### GetExternalNameOk

`func (o *MicrosoftGraphEducationClass) GetExternalNameOk() (*string, bool)`

GetExternalNameOk returns a tuple with the ExternalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalName

`func (o *MicrosoftGraphEducationClass) SetExternalName(v string)`

SetExternalName sets ExternalName field to given value.

### HasExternalName

`func (o *MicrosoftGraphEducationClass) HasExternalName() bool`

HasExternalName returns a boolean if a field has been set.

### SetExternalNameNil

`func (o *MicrosoftGraphEducationClass) SetExternalNameNil(b bool)`

 SetExternalNameNil sets the value for ExternalName to be an explicit nil

### UnsetExternalName
`func (o *MicrosoftGraphEducationClass) UnsetExternalName()`

UnsetExternalName ensures that no value is present for ExternalName, not even an explicit nil
### GetExternalSource

`func (o *MicrosoftGraphEducationClass) GetExternalSource() AnyOfmicrosoftGraphEducationExternalSource`

GetExternalSource returns the ExternalSource field if non-nil, zero value otherwise.

### GetExternalSourceOk

`func (o *MicrosoftGraphEducationClass) GetExternalSourceOk() (*AnyOfmicrosoftGraphEducationExternalSource, bool)`

GetExternalSourceOk returns a tuple with the ExternalSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalSource

`func (o *MicrosoftGraphEducationClass) SetExternalSource(v AnyOfmicrosoftGraphEducationExternalSource)`

SetExternalSource sets ExternalSource field to given value.

### HasExternalSource

`func (o *MicrosoftGraphEducationClass) HasExternalSource() bool`

HasExternalSource returns a boolean if a field has been set.

### SetExternalSourceNil

`func (o *MicrosoftGraphEducationClass) SetExternalSourceNil(b bool)`

 SetExternalSourceNil sets the value for ExternalSource to be an explicit nil

### UnsetExternalSource
`func (o *MicrosoftGraphEducationClass) UnsetExternalSource()`

UnsetExternalSource ensures that no value is present for ExternalSource, not even an explicit nil
### GetExternalSourceDetail

`func (o *MicrosoftGraphEducationClass) GetExternalSourceDetail() string`

GetExternalSourceDetail returns the ExternalSourceDetail field if non-nil, zero value otherwise.

### GetExternalSourceDetailOk

`func (o *MicrosoftGraphEducationClass) GetExternalSourceDetailOk() (*string, bool)`

GetExternalSourceDetailOk returns a tuple with the ExternalSourceDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalSourceDetail

`func (o *MicrosoftGraphEducationClass) SetExternalSourceDetail(v string)`

SetExternalSourceDetail sets ExternalSourceDetail field to given value.

### HasExternalSourceDetail

`func (o *MicrosoftGraphEducationClass) HasExternalSourceDetail() bool`

HasExternalSourceDetail returns a boolean if a field has been set.

### SetExternalSourceDetailNil

`func (o *MicrosoftGraphEducationClass) SetExternalSourceDetailNil(b bool)`

 SetExternalSourceDetailNil sets the value for ExternalSourceDetail to be an explicit nil

### UnsetExternalSourceDetail
`func (o *MicrosoftGraphEducationClass) UnsetExternalSourceDetail()`

UnsetExternalSourceDetail ensures that no value is present for ExternalSourceDetail, not even an explicit nil
### GetGrade

`func (o *MicrosoftGraphEducationClass) GetGrade() string`

GetGrade returns the Grade field if non-nil, zero value otherwise.

### GetGradeOk

`func (o *MicrosoftGraphEducationClass) GetGradeOk() (*string, bool)`

GetGradeOk returns a tuple with the Grade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrade

`func (o *MicrosoftGraphEducationClass) SetGrade(v string)`

SetGrade sets Grade field to given value.

### HasGrade

`func (o *MicrosoftGraphEducationClass) HasGrade() bool`

HasGrade returns a boolean if a field has been set.

### SetGradeNil

`func (o *MicrosoftGraphEducationClass) SetGradeNil(b bool)`

 SetGradeNil sets the value for Grade to be an explicit nil

### UnsetGrade
`func (o *MicrosoftGraphEducationClass) UnsetGrade()`

UnsetGrade ensures that no value is present for Grade, not even an explicit nil
### GetMailNickname

`func (o *MicrosoftGraphEducationClass) GetMailNickname() string`

GetMailNickname returns the MailNickname field if non-nil, zero value otherwise.

### GetMailNicknameOk

`func (o *MicrosoftGraphEducationClass) GetMailNicknameOk() (*string, bool)`

GetMailNicknameOk returns a tuple with the MailNickname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailNickname

`func (o *MicrosoftGraphEducationClass) SetMailNickname(v string)`

SetMailNickname sets MailNickname field to given value.

### HasMailNickname

`func (o *MicrosoftGraphEducationClass) HasMailNickname() bool`

HasMailNickname returns a boolean if a field has been set.

### GetTerm

`func (o *MicrosoftGraphEducationClass) GetTerm() AnyOfmicrosoftGraphEducationTerm`

GetTerm returns the Term field if non-nil, zero value otherwise.

### GetTermOk

`func (o *MicrosoftGraphEducationClass) GetTermOk() (*AnyOfmicrosoftGraphEducationTerm, bool)`

GetTermOk returns a tuple with the Term field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerm

`func (o *MicrosoftGraphEducationClass) SetTerm(v AnyOfmicrosoftGraphEducationTerm)`

SetTerm sets Term field to given value.

### HasTerm

`func (o *MicrosoftGraphEducationClass) HasTerm() bool`

HasTerm returns a boolean if a field has been set.

### SetTermNil

`func (o *MicrosoftGraphEducationClass) SetTermNil(b bool)`

 SetTermNil sets the value for Term to be an explicit nil

### UnsetTerm
`func (o *MicrosoftGraphEducationClass) UnsetTerm()`

UnsetTerm ensures that no value is present for Term, not even an explicit nil
### GetAssignmentCategories

`func (o *MicrosoftGraphEducationClass) GetAssignmentCategories() []MicrosoftGraphEducationCategory`

GetAssignmentCategories returns the AssignmentCategories field if non-nil, zero value otherwise.

### GetAssignmentCategoriesOk

`func (o *MicrosoftGraphEducationClass) GetAssignmentCategoriesOk() (*[]MicrosoftGraphEducationCategory, bool)`

GetAssignmentCategoriesOk returns a tuple with the AssignmentCategories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignmentCategories

`func (o *MicrosoftGraphEducationClass) SetAssignmentCategories(v []MicrosoftGraphEducationCategory)`

SetAssignmentCategories sets AssignmentCategories field to given value.

### HasAssignmentCategories

`func (o *MicrosoftGraphEducationClass) HasAssignmentCategories() bool`

HasAssignmentCategories returns a boolean if a field has been set.

### GetAssignmentDefaults

`func (o *MicrosoftGraphEducationClass) GetAssignmentDefaults() AnyOfmicrosoftGraphEducationAssignmentDefaults`

GetAssignmentDefaults returns the AssignmentDefaults field if non-nil, zero value otherwise.

### GetAssignmentDefaultsOk

`func (o *MicrosoftGraphEducationClass) GetAssignmentDefaultsOk() (*AnyOfmicrosoftGraphEducationAssignmentDefaults, bool)`

GetAssignmentDefaultsOk returns a tuple with the AssignmentDefaults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignmentDefaults

`func (o *MicrosoftGraphEducationClass) SetAssignmentDefaults(v AnyOfmicrosoftGraphEducationAssignmentDefaults)`

SetAssignmentDefaults sets AssignmentDefaults field to given value.

### HasAssignmentDefaults

`func (o *MicrosoftGraphEducationClass) HasAssignmentDefaults() bool`

HasAssignmentDefaults returns a boolean if a field has been set.

### SetAssignmentDefaultsNil

`func (o *MicrosoftGraphEducationClass) SetAssignmentDefaultsNil(b bool)`

 SetAssignmentDefaultsNil sets the value for AssignmentDefaults to be an explicit nil

### UnsetAssignmentDefaults
`func (o *MicrosoftGraphEducationClass) UnsetAssignmentDefaults()`

UnsetAssignmentDefaults ensures that no value is present for AssignmentDefaults, not even an explicit nil
### GetAssignments

`func (o *MicrosoftGraphEducationClass) GetAssignments() []MicrosoftGraphEducationAssignment`

GetAssignments returns the Assignments field if non-nil, zero value otherwise.

### GetAssignmentsOk

`func (o *MicrosoftGraphEducationClass) GetAssignmentsOk() (*[]MicrosoftGraphEducationAssignment, bool)`

GetAssignmentsOk returns a tuple with the Assignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignments

`func (o *MicrosoftGraphEducationClass) SetAssignments(v []MicrosoftGraphEducationAssignment)`

SetAssignments sets Assignments field to given value.

### HasAssignments

`func (o *MicrosoftGraphEducationClass) HasAssignments() bool`

HasAssignments returns a boolean if a field has been set.

### GetAssignmentSettings

`func (o *MicrosoftGraphEducationClass) GetAssignmentSettings() AnyOfmicrosoftGraphEducationAssignmentSettings`

GetAssignmentSettings returns the AssignmentSettings field if non-nil, zero value otherwise.

### GetAssignmentSettingsOk

`func (o *MicrosoftGraphEducationClass) GetAssignmentSettingsOk() (*AnyOfmicrosoftGraphEducationAssignmentSettings, bool)`

GetAssignmentSettingsOk returns a tuple with the AssignmentSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignmentSettings

`func (o *MicrosoftGraphEducationClass) SetAssignmentSettings(v AnyOfmicrosoftGraphEducationAssignmentSettings)`

SetAssignmentSettings sets AssignmentSettings field to given value.

### HasAssignmentSettings

`func (o *MicrosoftGraphEducationClass) HasAssignmentSettings() bool`

HasAssignmentSettings returns a boolean if a field has been set.

### SetAssignmentSettingsNil

`func (o *MicrosoftGraphEducationClass) SetAssignmentSettingsNil(b bool)`

 SetAssignmentSettingsNil sets the value for AssignmentSettings to be an explicit nil

### UnsetAssignmentSettings
`func (o *MicrosoftGraphEducationClass) UnsetAssignmentSettings()`

UnsetAssignmentSettings ensures that no value is present for AssignmentSettings, not even an explicit nil
### GetGroup

`func (o *MicrosoftGraphEducationClass) GetGroup() AnyOfmicrosoftGraphGroup`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *MicrosoftGraphEducationClass) GetGroupOk() (*AnyOfmicrosoftGraphGroup, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *MicrosoftGraphEducationClass) SetGroup(v AnyOfmicrosoftGraphGroup)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *MicrosoftGraphEducationClass) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### SetGroupNil

`func (o *MicrosoftGraphEducationClass) SetGroupNil(b bool)`

 SetGroupNil sets the value for Group to be an explicit nil

### UnsetGroup
`func (o *MicrosoftGraphEducationClass) UnsetGroup()`

UnsetGroup ensures that no value is present for Group, not even an explicit nil
### GetMembers

`func (o *MicrosoftGraphEducationClass) GetMembers() []MicrosoftGraphEducationUser`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *MicrosoftGraphEducationClass) GetMembersOk() (*[]MicrosoftGraphEducationUser, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *MicrosoftGraphEducationClass) SetMembers(v []MicrosoftGraphEducationUser)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *MicrosoftGraphEducationClass) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### GetSchools

`func (o *MicrosoftGraphEducationClass) GetSchools() []MicrosoftGraphEducationSchool`

GetSchools returns the Schools field if non-nil, zero value otherwise.

### GetSchoolsOk

`func (o *MicrosoftGraphEducationClass) GetSchoolsOk() (*[]MicrosoftGraphEducationSchool, bool)`

GetSchoolsOk returns a tuple with the Schools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchools

`func (o *MicrosoftGraphEducationClass) SetSchools(v []MicrosoftGraphEducationSchool)`

SetSchools sets Schools field to given value.

### HasSchools

`func (o *MicrosoftGraphEducationClass) HasSchools() bool`

HasSchools returns a boolean if a field has been set.

### GetTeachers

`func (o *MicrosoftGraphEducationClass) GetTeachers() []MicrosoftGraphEducationUser`

GetTeachers returns the Teachers field if non-nil, zero value otherwise.

### GetTeachersOk

`func (o *MicrosoftGraphEducationClass) GetTeachersOk() (*[]MicrosoftGraphEducationUser, bool)`

GetTeachersOk returns a tuple with the Teachers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeachers

`func (o *MicrosoftGraphEducationClass) SetTeachers(v []MicrosoftGraphEducationUser)`

SetTeachers sets Teachers field to given value.

### HasTeachers

`func (o *MicrosoftGraphEducationClass) HasTeachers() bool`

HasTeachers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


