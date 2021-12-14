# MicrosoftGraphApiApplication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcceptMappedClaims** | Pointer to **NullableBool** | When true, allows an application to use claims mapping without specifying a custom signing key. | [optional] 
**KnownClientApplications** | Pointer to **[]string** | Used for bundling consent if you have a solution that contains two parts: a client app and a custom web API app. If you set the appID of the client app to this value, the user only consents once to the client app. Azure AD knows that consenting to the client means implicitly consenting to the web API and automatically provisions service principals for both APIs at the same time. Both the client and the web API app must be registered in the same tenant. | [optional] 
**Oauth2PermissionScopes** | Pointer to [**[]MicrosoftGraphPermissionScope**](MicrosoftGraphPermissionScope.md) | The definition of the delegated permissions exposed by the web API represented by this application registration. These delegated permissions may be requested by a client application, and may be granted by users or administrators during consent. Delegated permissions are sometimes referred to as OAuth 2.0 scopes. | [optional] 
**PreAuthorizedApplications** | Pointer to [**[]AnyOfmicrosoftGraphPreAuthorizedApplication**](AnyOfmicrosoftGraphPreAuthorizedApplication.md) | Lists the client applications that are pre-authorized with the specified delegated permissions to access this application&#39;s APIs. Users are not required to consent to any pre-authorized application (for the permissions specified). However, any additional permissions not listed in preAuthorizedApplications (requested through incremental consent for example) will require user consent. | [optional] 
**RequestedAccessTokenVersion** | Pointer to **NullableInt32** | Specifies the access token version expected by this resource. This changes the version and format of the JWT produced independent of the endpoint or client used to request the access token.  The endpoint used, v1.0 or v2.0, is chosen by the client and only impacts the version of id_tokens. Resources need to explicitly configure requestedAccessTokenVersion to indicate the supported access token format.  Possible values for requestedAccessTokenVersion are 1, 2, or null. If the value is null, this defaults to 1, which corresponds to the v1.0 endpoint.  If signInAudience on the application is configured as AzureADandPersonalMicrosoftAccount, the value for this property must be 2 | [optional] 

## Methods

### NewMicrosoftGraphApiApplication

`func NewMicrosoftGraphApiApplication() *MicrosoftGraphApiApplication`

NewMicrosoftGraphApiApplication instantiates a new MicrosoftGraphApiApplication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphApiApplicationWithDefaults

`func NewMicrosoftGraphApiApplicationWithDefaults() *MicrosoftGraphApiApplication`

NewMicrosoftGraphApiApplicationWithDefaults instantiates a new MicrosoftGraphApiApplication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcceptMappedClaims

`func (o *MicrosoftGraphApiApplication) GetAcceptMappedClaims() bool`

GetAcceptMappedClaims returns the AcceptMappedClaims field if non-nil, zero value otherwise.

### GetAcceptMappedClaimsOk

`func (o *MicrosoftGraphApiApplication) GetAcceptMappedClaimsOk() (*bool, bool)`

GetAcceptMappedClaimsOk returns a tuple with the AcceptMappedClaims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptMappedClaims

`func (o *MicrosoftGraphApiApplication) SetAcceptMappedClaims(v bool)`

SetAcceptMappedClaims sets AcceptMappedClaims field to given value.

### HasAcceptMappedClaims

`func (o *MicrosoftGraphApiApplication) HasAcceptMappedClaims() bool`

HasAcceptMappedClaims returns a boolean if a field has been set.

### SetAcceptMappedClaimsNil

`func (o *MicrosoftGraphApiApplication) SetAcceptMappedClaimsNil(b bool)`

 SetAcceptMappedClaimsNil sets the value for AcceptMappedClaims to be an explicit nil

### UnsetAcceptMappedClaims
`func (o *MicrosoftGraphApiApplication) UnsetAcceptMappedClaims()`

UnsetAcceptMappedClaims ensures that no value is present for AcceptMappedClaims, not even an explicit nil
### GetKnownClientApplications

`func (o *MicrosoftGraphApiApplication) GetKnownClientApplications() []*string`

GetKnownClientApplications returns the KnownClientApplications field if non-nil, zero value otherwise.

### GetKnownClientApplicationsOk

`func (o *MicrosoftGraphApiApplication) GetKnownClientApplicationsOk() (*[]*string, bool)`

GetKnownClientApplicationsOk returns a tuple with the KnownClientApplications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKnownClientApplications

`func (o *MicrosoftGraphApiApplication) SetKnownClientApplications(v []*string)`

SetKnownClientApplications sets KnownClientApplications field to given value.

### HasKnownClientApplications

`func (o *MicrosoftGraphApiApplication) HasKnownClientApplications() bool`

HasKnownClientApplications returns a boolean if a field has been set.

### GetOauth2PermissionScopes

`func (o *MicrosoftGraphApiApplication) GetOauth2PermissionScopes() []MicrosoftGraphPermissionScope`

GetOauth2PermissionScopes returns the Oauth2PermissionScopes field if non-nil, zero value otherwise.

### GetOauth2PermissionScopesOk

`func (o *MicrosoftGraphApiApplication) GetOauth2PermissionScopesOk() (*[]MicrosoftGraphPermissionScope, bool)`

GetOauth2PermissionScopesOk returns a tuple with the Oauth2PermissionScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauth2PermissionScopes

`func (o *MicrosoftGraphApiApplication) SetOauth2PermissionScopes(v []MicrosoftGraphPermissionScope)`

SetOauth2PermissionScopes sets Oauth2PermissionScopes field to given value.

### HasOauth2PermissionScopes

`func (o *MicrosoftGraphApiApplication) HasOauth2PermissionScopes() bool`

HasOauth2PermissionScopes returns a boolean if a field has been set.

### GetPreAuthorizedApplications

`func (o *MicrosoftGraphApiApplication) GetPreAuthorizedApplications() []*AnyOfmicrosoftGraphPreAuthorizedApplication`

GetPreAuthorizedApplications returns the PreAuthorizedApplications field if non-nil, zero value otherwise.

### GetPreAuthorizedApplicationsOk

`func (o *MicrosoftGraphApiApplication) GetPreAuthorizedApplicationsOk() (*[]*AnyOfmicrosoftGraphPreAuthorizedApplication, bool)`

GetPreAuthorizedApplicationsOk returns a tuple with the PreAuthorizedApplications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreAuthorizedApplications

`func (o *MicrosoftGraphApiApplication) SetPreAuthorizedApplications(v []*AnyOfmicrosoftGraphPreAuthorizedApplication)`

SetPreAuthorizedApplications sets PreAuthorizedApplications field to given value.

### HasPreAuthorizedApplications

`func (o *MicrosoftGraphApiApplication) HasPreAuthorizedApplications() bool`

HasPreAuthorizedApplications returns a boolean if a field has been set.

### GetRequestedAccessTokenVersion

`func (o *MicrosoftGraphApiApplication) GetRequestedAccessTokenVersion() int32`

GetRequestedAccessTokenVersion returns the RequestedAccessTokenVersion field if non-nil, zero value otherwise.

### GetRequestedAccessTokenVersionOk

`func (o *MicrosoftGraphApiApplication) GetRequestedAccessTokenVersionOk() (*int32, bool)`

GetRequestedAccessTokenVersionOk returns a tuple with the RequestedAccessTokenVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedAccessTokenVersion

`func (o *MicrosoftGraphApiApplication) SetRequestedAccessTokenVersion(v int32)`

SetRequestedAccessTokenVersion sets RequestedAccessTokenVersion field to given value.

### HasRequestedAccessTokenVersion

`func (o *MicrosoftGraphApiApplication) HasRequestedAccessTokenVersion() bool`

HasRequestedAccessTokenVersion returns a boolean if a field has been set.

### SetRequestedAccessTokenVersionNil

`func (o *MicrosoftGraphApiApplication) SetRequestedAccessTokenVersionNil(b bool)`

 SetRequestedAccessTokenVersionNil sets the value for RequestedAccessTokenVersion to be an explicit nil

### UnsetRequestedAccessTokenVersion
`func (o *MicrosoftGraphApiApplication) UnsetRequestedAccessTokenVersion()`

UnsetRequestedAccessTokenVersion ensures that no value is present for RequestedAccessTokenVersion, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


