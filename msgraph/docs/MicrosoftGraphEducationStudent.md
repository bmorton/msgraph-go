# MicrosoftGraphEducationStudent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BirthDate** | Pointer to **NullableString** | Birth date of the student. | [optional] 
**ExternalId** | Pointer to **NullableString** | ID of the student in the source system. | [optional] 
**Gender** | Pointer to [**AnyOfmicrosoftGraphEducationGender**](anyOf&lt;microsoft.graph.educationGender&gt;.md) | The possible values are: female, male, other, unknownFutureValue. | [optional] 
**Grade** | Pointer to **NullableString** | Current grade level of the student. | [optional] 
**GraduationYear** | Pointer to **NullableString** | Year the student is graduating from the school. | [optional] 
**StudentNumber** | Pointer to **NullableString** | Student Number. | [optional] 

## Methods

### NewMicrosoftGraphEducationStudent

`func NewMicrosoftGraphEducationStudent() *MicrosoftGraphEducationStudent`

NewMicrosoftGraphEducationStudent instantiates a new MicrosoftGraphEducationStudent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphEducationStudentWithDefaults

`func NewMicrosoftGraphEducationStudentWithDefaults() *MicrosoftGraphEducationStudent`

NewMicrosoftGraphEducationStudentWithDefaults instantiates a new MicrosoftGraphEducationStudent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBirthDate

`func (o *MicrosoftGraphEducationStudent) GetBirthDate() string`

GetBirthDate returns the BirthDate field if non-nil, zero value otherwise.

### GetBirthDateOk

`func (o *MicrosoftGraphEducationStudent) GetBirthDateOk() (*string, bool)`

GetBirthDateOk returns a tuple with the BirthDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBirthDate

`func (o *MicrosoftGraphEducationStudent) SetBirthDate(v string)`

SetBirthDate sets BirthDate field to given value.

### HasBirthDate

`func (o *MicrosoftGraphEducationStudent) HasBirthDate() bool`

HasBirthDate returns a boolean if a field has been set.

### SetBirthDateNil

`func (o *MicrosoftGraphEducationStudent) SetBirthDateNil(b bool)`

 SetBirthDateNil sets the value for BirthDate to be an explicit nil

### UnsetBirthDate
`func (o *MicrosoftGraphEducationStudent) UnsetBirthDate()`

UnsetBirthDate ensures that no value is present for BirthDate, not even an explicit nil
### GetExternalId

`func (o *MicrosoftGraphEducationStudent) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *MicrosoftGraphEducationStudent) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *MicrosoftGraphEducationStudent) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *MicrosoftGraphEducationStudent) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *MicrosoftGraphEducationStudent) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *MicrosoftGraphEducationStudent) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetGender

`func (o *MicrosoftGraphEducationStudent) GetGender() AnyOfmicrosoftGraphEducationGender`

GetGender returns the Gender field if non-nil, zero value otherwise.

### GetGenderOk

`func (o *MicrosoftGraphEducationStudent) GetGenderOk() (*AnyOfmicrosoftGraphEducationGender, bool)`

GetGenderOk returns a tuple with the Gender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGender

`func (o *MicrosoftGraphEducationStudent) SetGender(v AnyOfmicrosoftGraphEducationGender)`

SetGender sets Gender field to given value.

### HasGender

`func (o *MicrosoftGraphEducationStudent) HasGender() bool`

HasGender returns a boolean if a field has been set.

### SetGenderNil

`func (o *MicrosoftGraphEducationStudent) SetGenderNil(b bool)`

 SetGenderNil sets the value for Gender to be an explicit nil

### UnsetGender
`func (o *MicrosoftGraphEducationStudent) UnsetGender()`

UnsetGender ensures that no value is present for Gender, not even an explicit nil
### GetGrade

`func (o *MicrosoftGraphEducationStudent) GetGrade() string`

GetGrade returns the Grade field if non-nil, zero value otherwise.

### GetGradeOk

`func (o *MicrosoftGraphEducationStudent) GetGradeOk() (*string, bool)`

GetGradeOk returns a tuple with the Grade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrade

`func (o *MicrosoftGraphEducationStudent) SetGrade(v string)`

SetGrade sets Grade field to given value.

### HasGrade

`func (o *MicrosoftGraphEducationStudent) HasGrade() bool`

HasGrade returns a boolean if a field has been set.

### SetGradeNil

`func (o *MicrosoftGraphEducationStudent) SetGradeNil(b bool)`

 SetGradeNil sets the value for Grade to be an explicit nil

### UnsetGrade
`func (o *MicrosoftGraphEducationStudent) UnsetGrade()`

UnsetGrade ensures that no value is present for Grade, not even an explicit nil
### GetGraduationYear

`func (o *MicrosoftGraphEducationStudent) GetGraduationYear() string`

GetGraduationYear returns the GraduationYear field if non-nil, zero value otherwise.

### GetGraduationYearOk

`func (o *MicrosoftGraphEducationStudent) GetGraduationYearOk() (*string, bool)`

GetGraduationYearOk returns a tuple with the GraduationYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGraduationYear

`func (o *MicrosoftGraphEducationStudent) SetGraduationYear(v string)`

SetGraduationYear sets GraduationYear field to given value.

### HasGraduationYear

`func (o *MicrosoftGraphEducationStudent) HasGraduationYear() bool`

HasGraduationYear returns a boolean if a field has been set.

### SetGraduationYearNil

`func (o *MicrosoftGraphEducationStudent) SetGraduationYearNil(b bool)`

 SetGraduationYearNil sets the value for GraduationYear to be an explicit nil

### UnsetGraduationYear
`func (o *MicrosoftGraphEducationStudent) UnsetGraduationYear()`

UnsetGraduationYear ensures that no value is present for GraduationYear, not even an explicit nil
### GetStudentNumber

`func (o *MicrosoftGraphEducationStudent) GetStudentNumber() string`

GetStudentNumber returns the StudentNumber field if non-nil, zero value otherwise.

### GetStudentNumberOk

`func (o *MicrosoftGraphEducationStudent) GetStudentNumberOk() (*string, bool)`

GetStudentNumberOk returns a tuple with the StudentNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStudentNumber

`func (o *MicrosoftGraphEducationStudent) SetStudentNumber(v string)`

SetStudentNumber sets StudentNumber field to given value.

### HasStudentNumber

`func (o *MicrosoftGraphEducationStudent) HasStudentNumber() bool`

HasStudentNumber returns a boolean if a field has been set.

### SetStudentNumberNil

`func (o *MicrosoftGraphEducationStudent) SetStudentNumberNil(b bool)`

 SetStudentNumberNil sets the value for StudentNumber to be an explicit nil

### UnsetStudentNumber
`func (o *MicrosoftGraphEducationStudent) UnsetStudentNumber()`

UnsetStudentNumber ensures that no value is present for StudentNumber, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


