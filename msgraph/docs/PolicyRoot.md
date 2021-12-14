# PolicyRoot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthenticationMethodsPolicy** | Pointer to [**AnyOfmicrosoftGraphAuthenticationMethodsPolicy**](anyOf&lt;microsoft.graph.authenticationMethodsPolicy&gt;.md) | The authentication methods and the users that are allowed to use them to sign in and perform multi-factor authentication (MFA) in Azure Active Directory (Azure AD). | [optional] 
**AuthenticationFlowsPolicy** | Pointer to [**AnyOfmicrosoftGraphAuthenticationFlowsPolicy**](anyOf&lt;microsoft.graph.authenticationFlowsPolicy&gt;.md) | The policy configuration of the self-service sign-up experience of external users. | [optional] 
**ActivityBasedTimeoutPolicies** | Pointer to [**[]MicrosoftGraphActivityBasedTimeoutPolicy**](MicrosoftGraphActivityBasedTimeoutPolicy.md) | The policy that controls the idle time out for web sessions for applications. | [optional] 
**AuthorizationPolicy** | Pointer to [**AnyOfmicrosoftGraphAuthorizationPolicy**](anyOf&lt;microsoft.graph.authorizationPolicy&gt;.md) | The policy that controls Azure AD authorization settings. | [optional] 
**ClaimsMappingPolicies** | Pointer to [**[]MicrosoftGraphClaimsMappingPolicy**](MicrosoftGraphClaimsMappingPolicy.md) | The claim-mapping policies for WS-Fed, SAML, OAuth 2.0, and OpenID Connect protocols, for tokens issued to a specific application. | [optional] 
**HomeRealmDiscoveryPolicies** | Pointer to [**[]MicrosoftGraphHomeRealmDiscoveryPolicy**](MicrosoftGraphHomeRealmDiscoveryPolicy.md) | The policy to control Azure AD authentication behavior for federated users. | [optional] 
**PermissionGrantPolicies** | Pointer to [**[]MicrosoftGraphPermissionGrantPolicy**](MicrosoftGraphPermissionGrantPolicy.md) | The policy that specifies the conditions under which consent can be granted. | [optional] 
**TokenIssuancePolicies** | Pointer to [**[]MicrosoftGraphTokenIssuancePolicy**](MicrosoftGraphTokenIssuancePolicy.md) | The policy that specifies the characteristics of SAML tokens issued by Azure AD. | [optional] 
**TokenLifetimePolicies** | Pointer to [**[]MicrosoftGraphTokenLifetimePolicy**](MicrosoftGraphTokenLifetimePolicy.md) | The policy that controls the lifetime of a JWT access token, an ID token, or a SAML 1.1/2.0 token issued by Azure AD. | [optional] 
**FeatureRolloutPolicies** | Pointer to [**[]MicrosoftGraphFeatureRolloutPolicy**](MicrosoftGraphFeatureRolloutPolicy.md) | The feature rollout policy associated with a directory object. | [optional] 
**AdminConsentRequestPolicy** | Pointer to [**AnyOfmicrosoftGraphAdminConsentRequestPolicy**](anyOf&lt;microsoft.graph.adminConsentRequestPolicy&gt;.md) | The policy by which consent requests are created and managed for the entire tenant. | [optional] 
**ConditionalAccessPolicies** | Pointer to [**[]MicrosoftGraphConditionalAccessPolicy**](MicrosoftGraphConditionalAccessPolicy.md) | The custom rules that define an access scenario. | [optional] 
**IdentitySecurityDefaultsEnforcementPolicy** | Pointer to [**AnyOfmicrosoftGraphIdentitySecurityDefaultsEnforcementPolicy**](anyOf&lt;microsoft.graph.identitySecurityDefaultsEnforcementPolicy&gt;.md) | The policy that represents the security defaults that protect against common attacks. | [optional] 

## Methods

### NewPolicyRoot

`func NewPolicyRoot() *PolicyRoot`

NewPolicyRoot instantiates a new PolicyRoot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyRootWithDefaults

`func NewPolicyRootWithDefaults() *PolicyRoot`

NewPolicyRootWithDefaults instantiates a new PolicyRoot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthenticationMethodsPolicy

`func (o *PolicyRoot) GetAuthenticationMethodsPolicy() AnyOfmicrosoftGraphAuthenticationMethodsPolicy`

GetAuthenticationMethodsPolicy returns the AuthenticationMethodsPolicy field if non-nil, zero value otherwise.

### GetAuthenticationMethodsPolicyOk

`func (o *PolicyRoot) GetAuthenticationMethodsPolicyOk() (*AnyOfmicrosoftGraphAuthenticationMethodsPolicy, bool)`

GetAuthenticationMethodsPolicyOk returns a tuple with the AuthenticationMethodsPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationMethodsPolicy

`func (o *PolicyRoot) SetAuthenticationMethodsPolicy(v AnyOfmicrosoftGraphAuthenticationMethodsPolicy)`

SetAuthenticationMethodsPolicy sets AuthenticationMethodsPolicy field to given value.

### HasAuthenticationMethodsPolicy

`func (o *PolicyRoot) HasAuthenticationMethodsPolicy() bool`

HasAuthenticationMethodsPolicy returns a boolean if a field has been set.

### SetAuthenticationMethodsPolicyNil

`func (o *PolicyRoot) SetAuthenticationMethodsPolicyNil(b bool)`

 SetAuthenticationMethodsPolicyNil sets the value for AuthenticationMethodsPolicy to be an explicit nil

### UnsetAuthenticationMethodsPolicy
`func (o *PolicyRoot) UnsetAuthenticationMethodsPolicy()`

UnsetAuthenticationMethodsPolicy ensures that no value is present for AuthenticationMethodsPolicy, not even an explicit nil
### GetAuthenticationFlowsPolicy

`func (o *PolicyRoot) GetAuthenticationFlowsPolicy() AnyOfmicrosoftGraphAuthenticationFlowsPolicy`

GetAuthenticationFlowsPolicy returns the AuthenticationFlowsPolicy field if non-nil, zero value otherwise.

### GetAuthenticationFlowsPolicyOk

`func (o *PolicyRoot) GetAuthenticationFlowsPolicyOk() (*AnyOfmicrosoftGraphAuthenticationFlowsPolicy, bool)`

GetAuthenticationFlowsPolicyOk returns a tuple with the AuthenticationFlowsPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationFlowsPolicy

`func (o *PolicyRoot) SetAuthenticationFlowsPolicy(v AnyOfmicrosoftGraphAuthenticationFlowsPolicy)`

SetAuthenticationFlowsPolicy sets AuthenticationFlowsPolicy field to given value.

### HasAuthenticationFlowsPolicy

`func (o *PolicyRoot) HasAuthenticationFlowsPolicy() bool`

HasAuthenticationFlowsPolicy returns a boolean if a field has been set.

### SetAuthenticationFlowsPolicyNil

`func (o *PolicyRoot) SetAuthenticationFlowsPolicyNil(b bool)`

 SetAuthenticationFlowsPolicyNil sets the value for AuthenticationFlowsPolicy to be an explicit nil

### UnsetAuthenticationFlowsPolicy
`func (o *PolicyRoot) UnsetAuthenticationFlowsPolicy()`

UnsetAuthenticationFlowsPolicy ensures that no value is present for AuthenticationFlowsPolicy, not even an explicit nil
### GetActivityBasedTimeoutPolicies

`func (o *PolicyRoot) GetActivityBasedTimeoutPolicies() []MicrosoftGraphActivityBasedTimeoutPolicy`

GetActivityBasedTimeoutPolicies returns the ActivityBasedTimeoutPolicies field if non-nil, zero value otherwise.

### GetActivityBasedTimeoutPoliciesOk

`func (o *PolicyRoot) GetActivityBasedTimeoutPoliciesOk() (*[]MicrosoftGraphActivityBasedTimeoutPolicy, bool)`

GetActivityBasedTimeoutPoliciesOk returns a tuple with the ActivityBasedTimeoutPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityBasedTimeoutPolicies

`func (o *PolicyRoot) SetActivityBasedTimeoutPolicies(v []MicrosoftGraphActivityBasedTimeoutPolicy)`

SetActivityBasedTimeoutPolicies sets ActivityBasedTimeoutPolicies field to given value.

### HasActivityBasedTimeoutPolicies

`func (o *PolicyRoot) HasActivityBasedTimeoutPolicies() bool`

HasActivityBasedTimeoutPolicies returns a boolean if a field has been set.

### GetAuthorizationPolicy

`func (o *PolicyRoot) GetAuthorizationPolicy() AnyOfmicrosoftGraphAuthorizationPolicy`

GetAuthorizationPolicy returns the AuthorizationPolicy field if non-nil, zero value otherwise.

### GetAuthorizationPolicyOk

`func (o *PolicyRoot) GetAuthorizationPolicyOk() (*AnyOfmicrosoftGraphAuthorizationPolicy, bool)`

GetAuthorizationPolicyOk returns a tuple with the AuthorizationPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationPolicy

`func (o *PolicyRoot) SetAuthorizationPolicy(v AnyOfmicrosoftGraphAuthorizationPolicy)`

SetAuthorizationPolicy sets AuthorizationPolicy field to given value.

### HasAuthorizationPolicy

`func (o *PolicyRoot) HasAuthorizationPolicy() bool`

HasAuthorizationPolicy returns a boolean if a field has been set.

### SetAuthorizationPolicyNil

`func (o *PolicyRoot) SetAuthorizationPolicyNil(b bool)`

 SetAuthorizationPolicyNil sets the value for AuthorizationPolicy to be an explicit nil

### UnsetAuthorizationPolicy
`func (o *PolicyRoot) UnsetAuthorizationPolicy()`

UnsetAuthorizationPolicy ensures that no value is present for AuthorizationPolicy, not even an explicit nil
### GetClaimsMappingPolicies

`func (o *PolicyRoot) GetClaimsMappingPolicies() []MicrosoftGraphClaimsMappingPolicy`

GetClaimsMappingPolicies returns the ClaimsMappingPolicies field if non-nil, zero value otherwise.

### GetClaimsMappingPoliciesOk

`func (o *PolicyRoot) GetClaimsMappingPoliciesOk() (*[]MicrosoftGraphClaimsMappingPolicy, bool)`

GetClaimsMappingPoliciesOk returns a tuple with the ClaimsMappingPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimsMappingPolicies

`func (o *PolicyRoot) SetClaimsMappingPolicies(v []MicrosoftGraphClaimsMappingPolicy)`

SetClaimsMappingPolicies sets ClaimsMappingPolicies field to given value.

### HasClaimsMappingPolicies

`func (o *PolicyRoot) HasClaimsMappingPolicies() bool`

HasClaimsMappingPolicies returns a boolean if a field has been set.

### GetHomeRealmDiscoveryPolicies

`func (o *PolicyRoot) GetHomeRealmDiscoveryPolicies() []MicrosoftGraphHomeRealmDiscoveryPolicy`

GetHomeRealmDiscoveryPolicies returns the HomeRealmDiscoveryPolicies field if non-nil, zero value otherwise.

### GetHomeRealmDiscoveryPoliciesOk

`func (o *PolicyRoot) GetHomeRealmDiscoveryPoliciesOk() (*[]MicrosoftGraphHomeRealmDiscoveryPolicy, bool)`

GetHomeRealmDiscoveryPoliciesOk returns a tuple with the HomeRealmDiscoveryPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomeRealmDiscoveryPolicies

`func (o *PolicyRoot) SetHomeRealmDiscoveryPolicies(v []MicrosoftGraphHomeRealmDiscoveryPolicy)`

SetHomeRealmDiscoveryPolicies sets HomeRealmDiscoveryPolicies field to given value.

### HasHomeRealmDiscoveryPolicies

`func (o *PolicyRoot) HasHomeRealmDiscoveryPolicies() bool`

HasHomeRealmDiscoveryPolicies returns a boolean if a field has been set.

### GetPermissionGrantPolicies

`func (o *PolicyRoot) GetPermissionGrantPolicies() []MicrosoftGraphPermissionGrantPolicy`

GetPermissionGrantPolicies returns the PermissionGrantPolicies field if non-nil, zero value otherwise.

### GetPermissionGrantPoliciesOk

`func (o *PolicyRoot) GetPermissionGrantPoliciesOk() (*[]MicrosoftGraphPermissionGrantPolicy, bool)`

GetPermissionGrantPoliciesOk returns a tuple with the PermissionGrantPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionGrantPolicies

`func (o *PolicyRoot) SetPermissionGrantPolicies(v []MicrosoftGraphPermissionGrantPolicy)`

SetPermissionGrantPolicies sets PermissionGrantPolicies field to given value.

### HasPermissionGrantPolicies

`func (o *PolicyRoot) HasPermissionGrantPolicies() bool`

HasPermissionGrantPolicies returns a boolean if a field has been set.

### GetTokenIssuancePolicies

`func (o *PolicyRoot) GetTokenIssuancePolicies() []MicrosoftGraphTokenIssuancePolicy`

GetTokenIssuancePolicies returns the TokenIssuancePolicies field if non-nil, zero value otherwise.

### GetTokenIssuancePoliciesOk

`func (o *PolicyRoot) GetTokenIssuancePoliciesOk() (*[]MicrosoftGraphTokenIssuancePolicy, bool)`

GetTokenIssuancePoliciesOk returns a tuple with the TokenIssuancePolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenIssuancePolicies

`func (o *PolicyRoot) SetTokenIssuancePolicies(v []MicrosoftGraphTokenIssuancePolicy)`

SetTokenIssuancePolicies sets TokenIssuancePolicies field to given value.

### HasTokenIssuancePolicies

`func (o *PolicyRoot) HasTokenIssuancePolicies() bool`

HasTokenIssuancePolicies returns a boolean if a field has been set.

### GetTokenLifetimePolicies

`func (o *PolicyRoot) GetTokenLifetimePolicies() []MicrosoftGraphTokenLifetimePolicy`

GetTokenLifetimePolicies returns the TokenLifetimePolicies field if non-nil, zero value otherwise.

### GetTokenLifetimePoliciesOk

`func (o *PolicyRoot) GetTokenLifetimePoliciesOk() (*[]MicrosoftGraphTokenLifetimePolicy, bool)`

GetTokenLifetimePoliciesOk returns a tuple with the TokenLifetimePolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenLifetimePolicies

`func (o *PolicyRoot) SetTokenLifetimePolicies(v []MicrosoftGraphTokenLifetimePolicy)`

SetTokenLifetimePolicies sets TokenLifetimePolicies field to given value.

### HasTokenLifetimePolicies

`func (o *PolicyRoot) HasTokenLifetimePolicies() bool`

HasTokenLifetimePolicies returns a boolean if a field has been set.

### GetFeatureRolloutPolicies

`func (o *PolicyRoot) GetFeatureRolloutPolicies() []MicrosoftGraphFeatureRolloutPolicy`

GetFeatureRolloutPolicies returns the FeatureRolloutPolicies field if non-nil, zero value otherwise.

### GetFeatureRolloutPoliciesOk

`func (o *PolicyRoot) GetFeatureRolloutPoliciesOk() (*[]MicrosoftGraphFeatureRolloutPolicy, bool)`

GetFeatureRolloutPoliciesOk returns a tuple with the FeatureRolloutPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureRolloutPolicies

`func (o *PolicyRoot) SetFeatureRolloutPolicies(v []MicrosoftGraphFeatureRolloutPolicy)`

SetFeatureRolloutPolicies sets FeatureRolloutPolicies field to given value.

### HasFeatureRolloutPolicies

`func (o *PolicyRoot) HasFeatureRolloutPolicies() bool`

HasFeatureRolloutPolicies returns a boolean if a field has been set.

### GetAdminConsentRequestPolicy

`func (o *PolicyRoot) GetAdminConsentRequestPolicy() AnyOfmicrosoftGraphAdminConsentRequestPolicy`

GetAdminConsentRequestPolicy returns the AdminConsentRequestPolicy field if non-nil, zero value otherwise.

### GetAdminConsentRequestPolicyOk

`func (o *PolicyRoot) GetAdminConsentRequestPolicyOk() (*AnyOfmicrosoftGraphAdminConsentRequestPolicy, bool)`

GetAdminConsentRequestPolicyOk returns a tuple with the AdminConsentRequestPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminConsentRequestPolicy

`func (o *PolicyRoot) SetAdminConsentRequestPolicy(v AnyOfmicrosoftGraphAdminConsentRequestPolicy)`

SetAdminConsentRequestPolicy sets AdminConsentRequestPolicy field to given value.

### HasAdminConsentRequestPolicy

`func (o *PolicyRoot) HasAdminConsentRequestPolicy() bool`

HasAdminConsentRequestPolicy returns a boolean if a field has been set.

### SetAdminConsentRequestPolicyNil

`func (o *PolicyRoot) SetAdminConsentRequestPolicyNil(b bool)`

 SetAdminConsentRequestPolicyNil sets the value for AdminConsentRequestPolicy to be an explicit nil

### UnsetAdminConsentRequestPolicy
`func (o *PolicyRoot) UnsetAdminConsentRequestPolicy()`

UnsetAdminConsentRequestPolicy ensures that no value is present for AdminConsentRequestPolicy, not even an explicit nil
### GetConditionalAccessPolicies

`func (o *PolicyRoot) GetConditionalAccessPolicies() []MicrosoftGraphConditionalAccessPolicy`

GetConditionalAccessPolicies returns the ConditionalAccessPolicies field if non-nil, zero value otherwise.

### GetConditionalAccessPoliciesOk

`func (o *PolicyRoot) GetConditionalAccessPoliciesOk() (*[]MicrosoftGraphConditionalAccessPolicy, bool)`

GetConditionalAccessPoliciesOk returns a tuple with the ConditionalAccessPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionalAccessPolicies

`func (o *PolicyRoot) SetConditionalAccessPolicies(v []MicrosoftGraphConditionalAccessPolicy)`

SetConditionalAccessPolicies sets ConditionalAccessPolicies field to given value.

### HasConditionalAccessPolicies

`func (o *PolicyRoot) HasConditionalAccessPolicies() bool`

HasConditionalAccessPolicies returns a boolean if a field has been set.

### GetIdentitySecurityDefaultsEnforcementPolicy

`func (o *PolicyRoot) GetIdentitySecurityDefaultsEnforcementPolicy() AnyOfmicrosoftGraphIdentitySecurityDefaultsEnforcementPolicy`

GetIdentitySecurityDefaultsEnforcementPolicy returns the IdentitySecurityDefaultsEnforcementPolicy field if non-nil, zero value otherwise.

### GetIdentitySecurityDefaultsEnforcementPolicyOk

`func (o *PolicyRoot) GetIdentitySecurityDefaultsEnforcementPolicyOk() (*AnyOfmicrosoftGraphIdentitySecurityDefaultsEnforcementPolicy, bool)`

GetIdentitySecurityDefaultsEnforcementPolicyOk returns a tuple with the IdentitySecurityDefaultsEnforcementPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentitySecurityDefaultsEnforcementPolicy

`func (o *PolicyRoot) SetIdentitySecurityDefaultsEnforcementPolicy(v AnyOfmicrosoftGraphIdentitySecurityDefaultsEnforcementPolicy)`

SetIdentitySecurityDefaultsEnforcementPolicy sets IdentitySecurityDefaultsEnforcementPolicy field to given value.

### HasIdentitySecurityDefaultsEnforcementPolicy

`func (o *PolicyRoot) HasIdentitySecurityDefaultsEnforcementPolicy() bool`

HasIdentitySecurityDefaultsEnforcementPolicy returns a boolean if a field has been set.

### SetIdentitySecurityDefaultsEnforcementPolicyNil

`func (o *PolicyRoot) SetIdentitySecurityDefaultsEnforcementPolicyNil(b bool)`

 SetIdentitySecurityDefaultsEnforcementPolicyNil sets the value for IdentitySecurityDefaultsEnforcementPolicy to be an explicit nil

### UnsetIdentitySecurityDefaultsEnforcementPolicy
`func (o *PolicyRoot) UnsetIdentitySecurityDefaultsEnforcementPolicy()`

UnsetIdentitySecurityDefaultsEnforcementPolicy ensures that no value is present for IdentitySecurityDefaultsEnforcementPolicy, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


