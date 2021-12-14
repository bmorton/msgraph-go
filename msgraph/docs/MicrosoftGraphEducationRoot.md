# MicrosoftGraphEducationRoot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Classes** | Pointer to [**[]MicrosoftGraphEducationClass**](MicrosoftGraphEducationClass.md) |  | [optional] 
**Me** | Pointer to [**AnyOfmicrosoftGraphEducationUser**](anyOf&lt;microsoft.graph.educationUser&gt;.md) |  | [optional] 
**Schools** | Pointer to [**[]MicrosoftGraphEducationSchool**](MicrosoftGraphEducationSchool.md) |  | [optional] 
**Users** | Pointer to [**[]MicrosoftGraphEducationUser**](MicrosoftGraphEducationUser.md) |  | [optional] 

## Methods

### NewMicrosoftGraphEducationRoot

`func NewMicrosoftGraphEducationRoot() *MicrosoftGraphEducationRoot`

NewMicrosoftGraphEducationRoot instantiates a new MicrosoftGraphEducationRoot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphEducationRootWithDefaults

`func NewMicrosoftGraphEducationRootWithDefaults() *MicrosoftGraphEducationRoot`

NewMicrosoftGraphEducationRootWithDefaults instantiates a new MicrosoftGraphEducationRoot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClasses

`func (o *MicrosoftGraphEducationRoot) GetClasses() []MicrosoftGraphEducationClass`

GetClasses returns the Classes field if non-nil, zero value otherwise.

### GetClassesOk

`func (o *MicrosoftGraphEducationRoot) GetClassesOk() (*[]MicrosoftGraphEducationClass, bool)`

GetClassesOk returns a tuple with the Classes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClasses

`func (o *MicrosoftGraphEducationRoot) SetClasses(v []MicrosoftGraphEducationClass)`

SetClasses sets Classes field to given value.

### HasClasses

`func (o *MicrosoftGraphEducationRoot) HasClasses() bool`

HasClasses returns a boolean if a field has been set.

### GetMe

`func (o *MicrosoftGraphEducationRoot) GetMe() AnyOfmicrosoftGraphEducationUser`

GetMe returns the Me field if non-nil, zero value otherwise.

### GetMeOk

`func (o *MicrosoftGraphEducationRoot) GetMeOk() (*AnyOfmicrosoftGraphEducationUser, bool)`

GetMeOk returns a tuple with the Me field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMe

`func (o *MicrosoftGraphEducationRoot) SetMe(v AnyOfmicrosoftGraphEducationUser)`

SetMe sets Me field to given value.

### HasMe

`func (o *MicrosoftGraphEducationRoot) HasMe() bool`

HasMe returns a boolean if a field has been set.

### SetMeNil

`func (o *MicrosoftGraphEducationRoot) SetMeNil(b bool)`

 SetMeNil sets the value for Me to be an explicit nil

### UnsetMe
`func (o *MicrosoftGraphEducationRoot) UnsetMe()`

UnsetMe ensures that no value is present for Me, not even an explicit nil
### GetSchools

`func (o *MicrosoftGraphEducationRoot) GetSchools() []MicrosoftGraphEducationSchool`

GetSchools returns the Schools field if non-nil, zero value otherwise.

### GetSchoolsOk

`func (o *MicrosoftGraphEducationRoot) GetSchoolsOk() (*[]MicrosoftGraphEducationSchool, bool)`

GetSchoolsOk returns a tuple with the Schools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchools

`func (o *MicrosoftGraphEducationRoot) SetSchools(v []MicrosoftGraphEducationSchool)`

SetSchools sets Schools field to given value.

### HasSchools

`func (o *MicrosoftGraphEducationRoot) HasSchools() bool`

HasSchools returns a boolean if a field has been set.

### GetUsers

`func (o *MicrosoftGraphEducationRoot) GetUsers() []MicrosoftGraphEducationUser`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *MicrosoftGraphEducationRoot) GetUsersOk() (*[]MicrosoftGraphEducationUser, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *MicrosoftGraphEducationRoot) SetUsers(v []MicrosoftGraphEducationUser)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *MicrosoftGraphEducationRoot) HasUsers() bool`

HasUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


