# MicrosoftGraphOAuth2PermissionGrant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**ClientId** | Pointer to **string** | The id of the client service principal for the application which is authorized to act on behalf of a signed-in user when accessing an API. Required. Supports $filter (eq only). | [optional] 
**ConsentType** | Pointer to **NullableString** | Indicates if authorization is granted for the client application to impersonate all users or only a specific user. AllPrincipals indicates authorization to impersonate all users. Principal indicates authorization to impersonate a specific user. Consent on behalf of all users can be granted by an administrator. Non-admin users may be authorized to consent on behalf of themselves in some cases, for some delegated permissions. Required. Supports $filter (eq only). | [optional] 
**PrincipalId** | Pointer to **NullableString** | The id of the user on behalf of whom the client is authorized to access the resource, when consentType is Principal. If consentType is AllPrincipals this value is null. Required when consentType is Principal. | [optional] 
**ResourceId** | Pointer to **string** | The id of the resource service principal to which access is authorized. This identifies the API which the client is authorized to attempt to call on behalf of a signed-in user. | [optional] 
**Scope** | Pointer to **NullableString** | A space-separated list of the claim values for delegated permissions which should be included in access tokens for the resource application (the API). For example, openid User.Read GroupMember.Read.All. Each claim value should match the value field of one of the delegated permissions defined by the API, listed in the publishedPermissionScopes property of the resource service principal. | [optional] 

## Methods

### NewMicrosoftGraphOAuth2PermissionGrant

`func NewMicrosoftGraphOAuth2PermissionGrant() *MicrosoftGraphOAuth2PermissionGrant`

NewMicrosoftGraphOAuth2PermissionGrant instantiates a new MicrosoftGraphOAuth2PermissionGrant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphOAuth2PermissionGrantWithDefaults

`func NewMicrosoftGraphOAuth2PermissionGrantWithDefaults() *MicrosoftGraphOAuth2PermissionGrant`

NewMicrosoftGraphOAuth2PermissionGrantWithDefaults instantiates a new MicrosoftGraphOAuth2PermissionGrant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphOAuth2PermissionGrant) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphOAuth2PermissionGrant) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphOAuth2PermissionGrant) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphOAuth2PermissionGrant) HasId() bool`

HasId returns a boolean if a field has been set.

### GetClientId

`func (o *MicrosoftGraphOAuth2PermissionGrant) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *MicrosoftGraphOAuth2PermissionGrant) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *MicrosoftGraphOAuth2PermissionGrant) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *MicrosoftGraphOAuth2PermissionGrant) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetConsentType

`func (o *MicrosoftGraphOAuth2PermissionGrant) GetConsentType() string`

GetConsentType returns the ConsentType field if non-nil, zero value otherwise.

### GetConsentTypeOk

`func (o *MicrosoftGraphOAuth2PermissionGrant) GetConsentTypeOk() (*string, bool)`

GetConsentTypeOk returns a tuple with the ConsentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsentType

`func (o *MicrosoftGraphOAuth2PermissionGrant) SetConsentType(v string)`

SetConsentType sets ConsentType field to given value.

### HasConsentType

`func (o *MicrosoftGraphOAuth2PermissionGrant) HasConsentType() bool`

HasConsentType returns a boolean if a field has been set.

### SetConsentTypeNil

`func (o *MicrosoftGraphOAuth2PermissionGrant) SetConsentTypeNil(b bool)`

 SetConsentTypeNil sets the value for ConsentType to be an explicit nil

### UnsetConsentType
`func (o *MicrosoftGraphOAuth2PermissionGrant) UnsetConsentType()`

UnsetConsentType ensures that no value is present for ConsentType, not even an explicit nil
### GetPrincipalId

`func (o *MicrosoftGraphOAuth2PermissionGrant) GetPrincipalId() string`

GetPrincipalId returns the PrincipalId field if non-nil, zero value otherwise.

### GetPrincipalIdOk

`func (o *MicrosoftGraphOAuth2PermissionGrant) GetPrincipalIdOk() (*string, bool)`

GetPrincipalIdOk returns a tuple with the PrincipalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipalId

`func (o *MicrosoftGraphOAuth2PermissionGrant) SetPrincipalId(v string)`

SetPrincipalId sets PrincipalId field to given value.

### HasPrincipalId

`func (o *MicrosoftGraphOAuth2PermissionGrant) HasPrincipalId() bool`

HasPrincipalId returns a boolean if a field has been set.

### SetPrincipalIdNil

`func (o *MicrosoftGraphOAuth2PermissionGrant) SetPrincipalIdNil(b bool)`

 SetPrincipalIdNil sets the value for PrincipalId to be an explicit nil

### UnsetPrincipalId
`func (o *MicrosoftGraphOAuth2PermissionGrant) UnsetPrincipalId()`

UnsetPrincipalId ensures that no value is present for PrincipalId, not even an explicit nil
### GetResourceId

`func (o *MicrosoftGraphOAuth2PermissionGrant) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *MicrosoftGraphOAuth2PermissionGrant) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *MicrosoftGraphOAuth2PermissionGrant) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.

### HasResourceId

`func (o *MicrosoftGraphOAuth2PermissionGrant) HasResourceId() bool`

HasResourceId returns a boolean if a field has been set.

### GetScope

`func (o *MicrosoftGraphOAuth2PermissionGrant) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *MicrosoftGraphOAuth2PermissionGrant) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *MicrosoftGraphOAuth2PermissionGrant) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *MicrosoftGraphOAuth2PermissionGrant) HasScope() bool`

HasScope returns a boolean if a field has been set.

### SetScopeNil

`func (o *MicrosoftGraphOAuth2PermissionGrant) SetScopeNil(b bool)`

 SetScopeNil sets the value for Scope to be an explicit nil

### UnsetScope
`func (o *MicrosoftGraphOAuth2PermissionGrant) UnsetScope()`

UnsetScope ensures that no value is present for Scope, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


