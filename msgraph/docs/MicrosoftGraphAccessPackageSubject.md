# MicrosoftGraphAccessPackageSubject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**DisplayName** | Pointer to **NullableString** | The display name of the subject. | [optional] 
**Email** | Pointer to **NullableString** | The email address of the subject. | [optional] 
**ObjectId** | Pointer to **NullableString** | The object identifier of the subject. null if the subject is not yet a user in the tenant. | [optional] 
**OnPremisesSecurityIdentifier** | Pointer to **NullableString** | A string representation of the principal&#39;s security identifier, if known, or null if the subject does not have a security identifier. | [optional] 
**PrincipalName** | Pointer to **NullableString** | The principal name, if known, of the subject. | [optional] 
**SubjectType** | Pointer to [**AnyOfmicrosoftGraphAccessPackageSubjectType**](anyOf&lt;microsoft.graph.accessPackageSubjectType&gt;.md) | The resource type of the subject. The possible values are: notSpecified, user, servicePrincipal, unknownFutureValue. | [optional] 
**ConnectedOrganization** | Pointer to [**AnyOfmicrosoftGraphConnectedOrganization**](anyOf&lt;microsoft.graph.connectedOrganization&gt;.md) | The connected organization of the subject. Read-only. Nullable. | [optional] 

## Methods

### NewMicrosoftGraphAccessPackageSubject

`func NewMicrosoftGraphAccessPackageSubject() *MicrosoftGraphAccessPackageSubject`

NewMicrosoftGraphAccessPackageSubject instantiates a new MicrosoftGraphAccessPackageSubject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphAccessPackageSubjectWithDefaults

`func NewMicrosoftGraphAccessPackageSubjectWithDefaults() *MicrosoftGraphAccessPackageSubject`

NewMicrosoftGraphAccessPackageSubjectWithDefaults instantiates a new MicrosoftGraphAccessPackageSubject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphAccessPackageSubject) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphAccessPackageSubject) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphAccessPackageSubject) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphAccessPackageSubject) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDisplayName

`func (o *MicrosoftGraphAccessPackageSubject) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphAccessPackageSubject) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *MicrosoftGraphAccessPackageSubject) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *MicrosoftGraphAccessPackageSubject) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *MicrosoftGraphAccessPackageSubject) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *MicrosoftGraphAccessPackageSubject) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetEmail

`func (o *MicrosoftGraphAccessPackageSubject) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *MicrosoftGraphAccessPackageSubject) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *MicrosoftGraphAccessPackageSubject) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *MicrosoftGraphAccessPackageSubject) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *MicrosoftGraphAccessPackageSubject) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *MicrosoftGraphAccessPackageSubject) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetObjectId

`func (o *MicrosoftGraphAccessPackageSubject) GetObjectId() string`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *MicrosoftGraphAccessPackageSubject) GetObjectIdOk() (*string, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *MicrosoftGraphAccessPackageSubject) SetObjectId(v string)`

SetObjectId sets ObjectId field to given value.

### HasObjectId

`func (o *MicrosoftGraphAccessPackageSubject) HasObjectId() bool`

HasObjectId returns a boolean if a field has been set.

### SetObjectIdNil

`func (o *MicrosoftGraphAccessPackageSubject) SetObjectIdNil(b bool)`

 SetObjectIdNil sets the value for ObjectId to be an explicit nil

### UnsetObjectId
`func (o *MicrosoftGraphAccessPackageSubject) UnsetObjectId()`

UnsetObjectId ensures that no value is present for ObjectId, not even an explicit nil
### GetOnPremisesSecurityIdentifier

`func (o *MicrosoftGraphAccessPackageSubject) GetOnPremisesSecurityIdentifier() string`

GetOnPremisesSecurityIdentifier returns the OnPremisesSecurityIdentifier field if non-nil, zero value otherwise.

### GetOnPremisesSecurityIdentifierOk

`func (o *MicrosoftGraphAccessPackageSubject) GetOnPremisesSecurityIdentifierOk() (*string, bool)`

GetOnPremisesSecurityIdentifierOk returns a tuple with the OnPremisesSecurityIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnPremisesSecurityIdentifier

`func (o *MicrosoftGraphAccessPackageSubject) SetOnPremisesSecurityIdentifier(v string)`

SetOnPremisesSecurityIdentifier sets OnPremisesSecurityIdentifier field to given value.

### HasOnPremisesSecurityIdentifier

`func (o *MicrosoftGraphAccessPackageSubject) HasOnPremisesSecurityIdentifier() bool`

HasOnPremisesSecurityIdentifier returns a boolean if a field has been set.

### SetOnPremisesSecurityIdentifierNil

`func (o *MicrosoftGraphAccessPackageSubject) SetOnPremisesSecurityIdentifierNil(b bool)`

 SetOnPremisesSecurityIdentifierNil sets the value for OnPremisesSecurityIdentifier to be an explicit nil

### UnsetOnPremisesSecurityIdentifier
`func (o *MicrosoftGraphAccessPackageSubject) UnsetOnPremisesSecurityIdentifier()`

UnsetOnPremisesSecurityIdentifier ensures that no value is present for OnPremisesSecurityIdentifier, not even an explicit nil
### GetPrincipalName

`func (o *MicrosoftGraphAccessPackageSubject) GetPrincipalName() string`

GetPrincipalName returns the PrincipalName field if non-nil, zero value otherwise.

### GetPrincipalNameOk

`func (o *MicrosoftGraphAccessPackageSubject) GetPrincipalNameOk() (*string, bool)`

GetPrincipalNameOk returns a tuple with the PrincipalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipalName

`func (o *MicrosoftGraphAccessPackageSubject) SetPrincipalName(v string)`

SetPrincipalName sets PrincipalName field to given value.

### HasPrincipalName

`func (o *MicrosoftGraphAccessPackageSubject) HasPrincipalName() bool`

HasPrincipalName returns a boolean if a field has been set.

### SetPrincipalNameNil

`func (o *MicrosoftGraphAccessPackageSubject) SetPrincipalNameNil(b bool)`

 SetPrincipalNameNil sets the value for PrincipalName to be an explicit nil

### UnsetPrincipalName
`func (o *MicrosoftGraphAccessPackageSubject) UnsetPrincipalName()`

UnsetPrincipalName ensures that no value is present for PrincipalName, not even an explicit nil
### GetSubjectType

`func (o *MicrosoftGraphAccessPackageSubject) GetSubjectType() AnyOfmicrosoftGraphAccessPackageSubjectType`

GetSubjectType returns the SubjectType field if non-nil, zero value otherwise.

### GetSubjectTypeOk

`func (o *MicrosoftGraphAccessPackageSubject) GetSubjectTypeOk() (*AnyOfmicrosoftGraphAccessPackageSubjectType, bool)`

GetSubjectTypeOk returns a tuple with the SubjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectType

`func (o *MicrosoftGraphAccessPackageSubject) SetSubjectType(v AnyOfmicrosoftGraphAccessPackageSubjectType)`

SetSubjectType sets SubjectType field to given value.

### HasSubjectType

`func (o *MicrosoftGraphAccessPackageSubject) HasSubjectType() bool`

HasSubjectType returns a boolean if a field has been set.

### SetSubjectTypeNil

`func (o *MicrosoftGraphAccessPackageSubject) SetSubjectTypeNil(b bool)`

 SetSubjectTypeNil sets the value for SubjectType to be an explicit nil

### UnsetSubjectType
`func (o *MicrosoftGraphAccessPackageSubject) UnsetSubjectType()`

UnsetSubjectType ensures that no value is present for SubjectType, not even an explicit nil
### GetConnectedOrganization

`func (o *MicrosoftGraphAccessPackageSubject) GetConnectedOrganization() AnyOfmicrosoftGraphConnectedOrganization`

GetConnectedOrganization returns the ConnectedOrganization field if non-nil, zero value otherwise.

### GetConnectedOrganizationOk

`func (o *MicrosoftGraphAccessPackageSubject) GetConnectedOrganizationOk() (*AnyOfmicrosoftGraphConnectedOrganization, bool)`

GetConnectedOrganizationOk returns a tuple with the ConnectedOrganization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedOrganization

`func (o *MicrosoftGraphAccessPackageSubject) SetConnectedOrganization(v AnyOfmicrosoftGraphConnectedOrganization)`

SetConnectedOrganization sets ConnectedOrganization field to given value.

### HasConnectedOrganization

`func (o *MicrosoftGraphAccessPackageSubject) HasConnectedOrganization() bool`

HasConnectedOrganization returns a boolean if a field has been set.

### SetConnectedOrganizationNil

`func (o *MicrosoftGraphAccessPackageSubject) SetConnectedOrganizationNil(b bool)`

 SetConnectedOrganizationNil sets the value for ConnectedOrganization to be an explicit nil

### UnsetConnectedOrganization
`func (o *MicrosoftGraphAccessPackageSubject) UnsetConnectedOrganization()`

UnsetConnectedOrganization ensures that no value is present for ConnectedOrganization, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


