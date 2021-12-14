# AccessPackageSubject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **NullableString** | The display name of the subject. | [optional] 
**Email** | Pointer to **NullableString** | The email address of the subject. | [optional] 
**ObjectId** | Pointer to **NullableString** | The object identifier of the subject. null if the subject is not yet a user in the tenant. | [optional] 
**OnPremisesSecurityIdentifier** | Pointer to **NullableString** | A string representation of the principal&#39;s security identifier, if known, or null if the subject does not have a security identifier. | [optional] 
**PrincipalName** | Pointer to **NullableString** | The principal name, if known, of the subject. | [optional] 
**SubjectType** | Pointer to [**AnyOfmicrosoftGraphAccessPackageSubjectType**](anyOf&lt;microsoft.graph.accessPackageSubjectType&gt;.md) | The resource type of the subject. The possible values are: notSpecified, user, servicePrincipal, unknownFutureValue. | [optional] 
**ConnectedOrganization** | Pointer to [**AnyOfmicrosoftGraphConnectedOrganization**](anyOf&lt;microsoft.graph.connectedOrganization&gt;.md) | The connected organization of the subject. Read-only. Nullable. | [optional] 

## Methods

### NewAccessPackageSubject

`func NewAccessPackageSubject() *AccessPackageSubject`

NewAccessPackageSubject instantiates a new AccessPackageSubject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessPackageSubjectWithDefaults

`func NewAccessPackageSubjectWithDefaults() *AccessPackageSubject`

NewAccessPackageSubjectWithDefaults instantiates a new AccessPackageSubject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *AccessPackageSubject) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *AccessPackageSubject) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *AccessPackageSubject) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *AccessPackageSubject) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *AccessPackageSubject) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *AccessPackageSubject) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetEmail

`func (o *AccessPackageSubject) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *AccessPackageSubject) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *AccessPackageSubject) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *AccessPackageSubject) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *AccessPackageSubject) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *AccessPackageSubject) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetObjectId

`func (o *AccessPackageSubject) GetObjectId() string`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *AccessPackageSubject) GetObjectIdOk() (*string, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *AccessPackageSubject) SetObjectId(v string)`

SetObjectId sets ObjectId field to given value.

### HasObjectId

`func (o *AccessPackageSubject) HasObjectId() bool`

HasObjectId returns a boolean if a field has been set.

### SetObjectIdNil

`func (o *AccessPackageSubject) SetObjectIdNil(b bool)`

 SetObjectIdNil sets the value for ObjectId to be an explicit nil

### UnsetObjectId
`func (o *AccessPackageSubject) UnsetObjectId()`

UnsetObjectId ensures that no value is present for ObjectId, not even an explicit nil
### GetOnPremisesSecurityIdentifier

`func (o *AccessPackageSubject) GetOnPremisesSecurityIdentifier() string`

GetOnPremisesSecurityIdentifier returns the OnPremisesSecurityIdentifier field if non-nil, zero value otherwise.

### GetOnPremisesSecurityIdentifierOk

`func (o *AccessPackageSubject) GetOnPremisesSecurityIdentifierOk() (*string, bool)`

GetOnPremisesSecurityIdentifierOk returns a tuple with the OnPremisesSecurityIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnPremisesSecurityIdentifier

`func (o *AccessPackageSubject) SetOnPremisesSecurityIdentifier(v string)`

SetOnPremisesSecurityIdentifier sets OnPremisesSecurityIdentifier field to given value.

### HasOnPremisesSecurityIdentifier

`func (o *AccessPackageSubject) HasOnPremisesSecurityIdentifier() bool`

HasOnPremisesSecurityIdentifier returns a boolean if a field has been set.

### SetOnPremisesSecurityIdentifierNil

`func (o *AccessPackageSubject) SetOnPremisesSecurityIdentifierNil(b bool)`

 SetOnPremisesSecurityIdentifierNil sets the value for OnPremisesSecurityIdentifier to be an explicit nil

### UnsetOnPremisesSecurityIdentifier
`func (o *AccessPackageSubject) UnsetOnPremisesSecurityIdentifier()`

UnsetOnPremisesSecurityIdentifier ensures that no value is present for OnPremisesSecurityIdentifier, not even an explicit nil
### GetPrincipalName

`func (o *AccessPackageSubject) GetPrincipalName() string`

GetPrincipalName returns the PrincipalName field if non-nil, zero value otherwise.

### GetPrincipalNameOk

`func (o *AccessPackageSubject) GetPrincipalNameOk() (*string, bool)`

GetPrincipalNameOk returns a tuple with the PrincipalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipalName

`func (o *AccessPackageSubject) SetPrincipalName(v string)`

SetPrincipalName sets PrincipalName field to given value.

### HasPrincipalName

`func (o *AccessPackageSubject) HasPrincipalName() bool`

HasPrincipalName returns a boolean if a field has been set.

### SetPrincipalNameNil

`func (o *AccessPackageSubject) SetPrincipalNameNil(b bool)`

 SetPrincipalNameNil sets the value for PrincipalName to be an explicit nil

### UnsetPrincipalName
`func (o *AccessPackageSubject) UnsetPrincipalName()`

UnsetPrincipalName ensures that no value is present for PrincipalName, not even an explicit nil
### GetSubjectType

`func (o *AccessPackageSubject) GetSubjectType() AnyOfmicrosoftGraphAccessPackageSubjectType`

GetSubjectType returns the SubjectType field if non-nil, zero value otherwise.

### GetSubjectTypeOk

`func (o *AccessPackageSubject) GetSubjectTypeOk() (*AnyOfmicrosoftGraphAccessPackageSubjectType, bool)`

GetSubjectTypeOk returns a tuple with the SubjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectType

`func (o *AccessPackageSubject) SetSubjectType(v AnyOfmicrosoftGraphAccessPackageSubjectType)`

SetSubjectType sets SubjectType field to given value.

### HasSubjectType

`func (o *AccessPackageSubject) HasSubjectType() bool`

HasSubjectType returns a boolean if a field has been set.

### SetSubjectTypeNil

`func (o *AccessPackageSubject) SetSubjectTypeNil(b bool)`

 SetSubjectTypeNil sets the value for SubjectType to be an explicit nil

### UnsetSubjectType
`func (o *AccessPackageSubject) UnsetSubjectType()`

UnsetSubjectType ensures that no value is present for SubjectType, not even an explicit nil
### GetConnectedOrganization

`func (o *AccessPackageSubject) GetConnectedOrganization() AnyOfmicrosoftGraphConnectedOrganization`

GetConnectedOrganization returns the ConnectedOrganization field if non-nil, zero value otherwise.

### GetConnectedOrganizationOk

`func (o *AccessPackageSubject) GetConnectedOrganizationOk() (*AnyOfmicrosoftGraphConnectedOrganization, bool)`

GetConnectedOrganizationOk returns a tuple with the ConnectedOrganization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedOrganization

`func (o *AccessPackageSubject) SetConnectedOrganization(v AnyOfmicrosoftGraphConnectedOrganization)`

SetConnectedOrganization sets ConnectedOrganization field to given value.

### HasConnectedOrganization

`func (o *AccessPackageSubject) HasConnectedOrganization() bool`

HasConnectedOrganization returns a boolean if a field has been set.

### SetConnectedOrganizationNil

`func (o *AccessPackageSubject) SetConnectedOrganizationNil(b bool)`

 SetConnectedOrganizationNil sets the value for ConnectedOrganization to be an explicit nil

### UnsetConnectedOrganization
`func (o *AccessPackageSubject) UnsetConnectedOrganization()`

UnsetConnectedOrganization ensures that no value is present for ConnectedOrganization, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


