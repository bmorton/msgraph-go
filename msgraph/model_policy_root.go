/*
Partial Graph API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package msgraph

import (
	"encoding/json"
)

// PolicyRoot struct for PolicyRoot
type PolicyRoot struct {
	// The authentication methods and the users that are allowed to use them to sign in and perform multi-factor authentication (MFA) in Azure Active Directory (Azure AD).
	AuthenticationMethodsPolicy AnyOfmicrosoftGraphAuthenticationMethodsPolicy `json:"authenticationMethodsPolicy,omitempty"`
	// The policy configuration of the self-service sign-up experience of external users.
	AuthenticationFlowsPolicy AnyOfmicrosoftGraphAuthenticationFlowsPolicy `json:"authenticationFlowsPolicy,omitempty"`
	// The policy that controls the idle time out for web sessions for applications.
	ActivityBasedTimeoutPolicies *[]MicrosoftGraphActivityBasedTimeoutPolicy `json:"activityBasedTimeoutPolicies,omitempty"`
	// The policy that controls Azure AD authorization settings.
	AuthorizationPolicy AnyOfmicrosoftGraphAuthorizationPolicy `json:"authorizationPolicy,omitempty"`
	// The claim-mapping policies for WS-Fed, SAML, OAuth 2.0, and OpenID Connect protocols, for tokens issued to a specific application.
	ClaimsMappingPolicies *[]MicrosoftGraphClaimsMappingPolicy `json:"claimsMappingPolicies,omitempty"`
	// The policy to control Azure AD authentication behavior for federated users.
	HomeRealmDiscoveryPolicies *[]MicrosoftGraphHomeRealmDiscoveryPolicy `json:"homeRealmDiscoveryPolicies,omitempty"`
	// The policy that specifies the conditions under which consent can be granted.
	PermissionGrantPolicies *[]MicrosoftGraphPermissionGrantPolicy `json:"permissionGrantPolicies,omitempty"`
	// The policy that specifies the characteristics of SAML tokens issued by Azure AD.
	TokenIssuancePolicies *[]MicrosoftGraphTokenIssuancePolicy `json:"tokenIssuancePolicies,omitempty"`
	// The policy that controls the lifetime of a JWT access token, an ID token, or a SAML 1.1/2.0 token issued by Azure AD.
	TokenLifetimePolicies *[]MicrosoftGraphTokenLifetimePolicy `json:"tokenLifetimePolicies,omitempty"`
	// The feature rollout policy associated with a directory object.
	FeatureRolloutPolicies *[]MicrosoftGraphFeatureRolloutPolicy `json:"featureRolloutPolicies,omitempty"`
	// The policy by which consent requests are created and managed for the entire tenant.
	AdminConsentRequestPolicy AnyOfmicrosoftGraphAdminConsentRequestPolicy `json:"adminConsentRequestPolicy,omitempty"`
	// The custom rules that define an access scenario.
	ConditionalAccessPolicies *[]MicrosoftGraphConditionalAccessPolicy `json:"conditionalAccessPolicies,omitempty"`
	// The policy that represents the security defaults that protect against common attacks.
	IdentitySecurityDefaultsEnforcementPolicy AnyOfmicrosoftGraphIdentitySecurityDefaultsEnforcementPolicy `json:"identitySecurityDefaultsEnforcementPolicy,omitempty"`
}

// NewPolicyRoot instantiates a new PolicyRoot object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPolicyRoot() *PolicyRoot {
	this := PolicyRoot{}
	return &this
}

// NewPolicyRootWithDefaults instantiates a new PolicyRoot object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPolicyRootWithDefaults() *PolicyRoot {
	this := PolicyRoot{}
	return &this
}

// GetAuthenticationMethodsPolicy returns the AuthenticationMethodsPolicy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PolicyRoot) GetAuthenticationMethodsPolicy() AnyOfmicrosoftGraphAuthenticationMethodsPolicy {
	if o == nil  {
		var ret AnyOfmicrosoftGraphAuthenticationMethodsPolicy
		return ret
	}
	return o.AuthenticationMethodsPolicy
}

// GetAuthenticationMethodsPolicyOk returns a tuple with the AuthenticationMethodsPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PolicyRoot) GetAuthenticationMethodsPolicyOk() (*AnyOfmicrosoftGraphAuthenticationMethodsPolicy, bool) {
	if o == nil || o.AuthenticationMethodsPolicy == nil {
		return nil, false
	}
	return &o.AuthenticationMethodsPolicy, true
}

// HasAuthenticationMethodsPolicy returns a boolean if a field has been set.
func (o *PolicyRoot) HasAuthenticationMethodsPolicy() bool {
	if o != nil && o.AuthenticationMethodsPolicy != nil {
		return true
	}

	return false
}

// SetAuthenticationMethodsPolicy gets a reference to the given AnyOfmicrosoftGraphAuthenticationMethodsPolicy and assigns it to the AuthenticationMethodsPolicy field.
func (o *PolicyRoot) SetAuthenticationMethodsPolicy(v AnyOfmicrosoftGraphAuthenticationMethodsPolicy) {
	o.AuthenticationMethodsPolicy = v
}

// GetAuthenticationFlowsPolicy returns the AuthenticationFlowsPolicy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PolicyRoot) GetAuthenticationFlowsPolicy() AnyOfmicrosoftGraphAuthenticationFlowsPolicy {
	if o == nil  {
		var ret AnyOfmicrosoftGraphAuthenticationFlowsPolicy
		return ret
	}
	return o.AuthenticationFlowsPolicy
}

// GetAuthenticationFlowsPolicyOk returns a tuple with the AuthenticationFlowsPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PolicyRoot) GetAuthenticationFlowsPolicyOk() (*AnyOfmicrosoftGraphAuthenticationFlowsPolicy, bool) {
	if o == nil || o.AuthenticationFlowsPolicy == nil {
		return nil, false
	}
	return &o.AuthenticationFlowsPolicy, true
}

// HasAuthenticationFlowsPolicy returns a boolean if a field has been set.
func (o *PolicyRoot) HasAuthenticationFlowsPolicy() bool {
	if o != nil && o.AuthenticationFlowsPolicy != nil {
		return true
	}

	return false
}

// SetAuthenticationFlowsPolicy gets a reference to the given AnyOfmicrosoftGraphAuthenticationFlowsPolicy and assigns it to the AuthenticationFlowsPolicy field.
func (o *PolicyRoot) SetAuthenticationFlowsPolicy(v AnyOfmicrosoftGraphAuthenticationFlowsPolicy) {
	o.AuthenticationFlowsPolicy = v
}

// GetActivityBasedTimeoutPolicies returns the ActivityBasedTimeoutPolicies field value if set, zero value otherwise.
func (o *PolicyRoot) GetActivityBasedTimeoutPolicies() []MicrosoftGraphActivityBasedTimeoutPolicy {
	if o == nil || o.ActivityBasedTimeoutPolicies == nil {
		var ret []MicrosoftGraphActivityBasedTimeoutPolicy
		return ret
	}
	return *o.ActivityBasedTimeoutPolicies
}

// GetActivityBasedTimeoutPoliciesOk returns a tuple with the ActivityBasedTimeoutPolicies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyRoot) GetActivityBasedTimeoutPoliciesOk() (*[]MicrosoftGraphActivityBasedTimeoutPolicy, bool) {
	if o == nil || o.ActivityBasedTimeoutPolicies == nil {
		return nil, false
	}
	return o.ActivityBasedTimeoutPolicies, true
}

// HasActivityBasedTimeoutPolicies returns a boolean if a field has been set.
func (o *PolicyRoot) HasActivityBasedTimeoutPolicies() bool {
	if o != nil && o.ActivityBasedTimeoutPolicies != nil {
		return true
	}

	return false
}

// SetActivityBasedTimeoutPolicies gets a reference to the given []MicrosoftGraphActivityBasedTimeoutPolicy and assigns it to the ActivityBasedTimeoutPolicies field.
func (o *PolicyRoot) SetActivityBasedTimeoutPolicies(v []MicrosoftGraphActivityBasedTimeoutPolicy) {
	o.ActivityBasedTimeoutPolicies = &v
}

// GetAuthorizationPolicy returns the AuthorizationPolicy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PolicyRoot) GetAuthorizationPolicy() AnyOfmicrosoftGraphAuthorizationPolicy {
	if o == nil  {
		var ret AnyOfmicrosoftGraphAuthorizationPolicy
		return ret
	}
	return o.AuthorizationPolicy
}

// GetAuthorizationPolicyOk returns a tuple with the AuthorizationPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PolicyRoot) GetAuthorizationPolicyOk() (*AnyOfmicrosoftGraphAuthorizationPolicy, bool) {
	if o == nil || o.AuthorizationPolicy == nil {
		return nil, false
	}
	return &o.AuthorizationPolicy, true
}

// HasAuthorizationPolicy returns a boolean if a field has been set.
func (o *PolicyRoot) HasAuthorizationPolicy() bool {
	if o != nil && o.AuthorizationPolicy != nil {
		return true
	}

	return false
}

// SetAuthorizationPolicy gets a reference to the given AnyOfmicrosoftGraphAuthorizationPolicy and assigns it to the AuthorizationPolicy field.
func (o *PolicyRoot) SetAuthorizationPolicy(v AnyOfmicrosoftGraphAuthorizationPolicy) {
	o.AuthorizationPolicy = v
}

// GetClaimsMappingPolicies returns the ClaimsMappingPolicies field value if set, zero value otherwise.
func (o *PolicyRoot) GetClaimsMappingPolicies() []MicrosoftGraphClaimsMappingPolicy {
	if o == nil || o.ClaimsMappingPolicies == nil {
		var ret []MicrosoftGraphClaimsMappingPolicy
		return ret
	}
	return *o.ClaimsMappingPolicies
}

// GetClaimsMappingPoliciesOk returns a tuple with the ClaimsMappingPolicies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyRoot) GetClaimsMappingPoliciesOk() (*[]MicrosoftGraphClaimsMappingPolicy, bool) {
	if o == nil || o.ClaimsMappingPolicies == nil {
		return nil, false
	}
	return o.ClaimsMappingPolicies, true
}

// HasClaimsMappingPolicies returns a boolean if a field has been set.
func (o *PolicyRoot) HasClaimsMappingPolicies() bool {
	if o != nil && o.ClaimsMappingPolicies != nil {
		return true
	}

	return false
}

// SetClaimsMappingPolicies gets a reference to the given []MicrosoftGraphClaimsMappingPolicy and assigns it to the ClaimsMappingPolicies field.
func (o *PolicyRoot) SetClaimsMappingPolicies(v []MicrosoftGraphClaimsMappingPolicy) {
	o.ClaimsMappingPolicies = &v
}

// GetHomeRealmDiscoveryPolicies returns the HomeRealmDiscoveryPolicies field value if set, zero value otherwise.
func (o *PolicyRoot) GetHomeRealmDiscoveryPolicies() []MicrosoftGraphHomeRealmDiscoveryPolicy {
	if o == nil || o.HomeRealmDiscoveryPolicies == nil {
		var ret []MicrosoftGraphHomeRealmDiscoveryPolicy
		return ret
	}
	return *o.HomeRealmDiscoveryPolicies
}

// GetHomeRealmDiscoveryPoliciesOk returns a tuple with the HomeRealmDiscoveryPolicies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyRoot) GetHomeRealmDiscoveryPoliciesOk() (*[]MicrosoftGraphHomeRealmDiscoveryPolicy, bool) {
	if o == nil || o.HomeRealmDiscoveryPolicies == nil {
		return nil, false
	}
	return o.HomeRealmDiscoveryPolicies, true
}

// HasHomeRealmDiscoveryPolicies returns a boolean if a field has been set.
func (o *PolicyRoot) HasHomeRealmDiscoveryPolicies() bool {
	if o != nil && o.HomeRealmDiscoveryPolicies != nil {
		return true
	}

	return false
}

// SetHomeRealmDiscoveryPolicies gets a reference to the given []MicrosoftGraphHomeRealmDiscoveryPolicy and assigns it to the HomeRealmDiscoveryPolicies field.
func (o *PolicyRoot) SetHomeRealmDiscoveryPolicies(v []MicrosoftGraphHomeRealmDiscoveryPolicy) {
	o.HomeRealmDiscoveryPolicies = &v
}

// GetPermissionGrantPolicies returns the PermissionGrantPolicies field value if set, zero value otherwise.
func (o *PolicyRoot) GetPermissionGrantPolicies() []MicrosoftGraphPermissionGrantPolicy {
	if o == nil || o.PermissionGrantPolicies == nil {
		var ret []MicrosoftGraphPermissionGrantPolicy
		return ret
	}
	return *o.PermissionGrantPolicies
}

// GetPermissionGrantPoliciesOk returns a tuple with the PermissionGrantPolicies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyRoot) GetPermissionGrantPoliciesOk() (*[]MicrosoftGraphPermissionGrantPolicy, bool) {
	if o == nil || o.PermissionGrantPolicies == nil {
		return nil, false
	}
	return o.PermissionGrantPolicies, true
}

// HasPermissionGrantPolicies returns a boolean if a field has been set.
func (o *PolicyRoot) HasPermissionGrantPolicies() bool {
	if o != nil && o.PermissionGrantPolicies != nil {
		return true
	}

	return false
}

// SetPermissionGrantPolicies gets a reference to the given []MicrosoftGraphPermissionGrantPolicy and assigns it to the PermissionGrantPolicies field.
func (o *PolicyRoot) SetPermissionGrantPolicies(v []MicrosoftGraphPermissionGrantPolicy) {
	o.PermissionGrantPolicies = &v
}

// GetTokenIssuancePolicies returns the TokenIssuancePolicies field value if set, zero value otherwise.
func (o *PolicyRoot) GetTokenIssuancePolicies() []MicrosoftGraphTokenIssuancePolicy {
	if o == nil || o.TokenIssuancePolicies == nil {
		var ret []MicrosoftGraphTokenIssuancePolicy
		return ret
	}
	return *o.TokenIssuancePolicies
}

// GetTokenIssuancePoliciesOk returns a tuple with the TokenIssuancePolicies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyRoot) GetTokenIssuancePoliciesOk() (*[]MicrosoftGraphTokenIssuancePolicy, bool) {
	if o == nil || o.TokenIssuancePolicies == nil {
		return nil, false
	}
	return o.TokenIssuancePolicies, true
}

// HasTokenIssuancePolicies returns a boolean if a field has been set.
func (o *PolicyRoot) HasTokenIssuancePolicies() bool {
	if o != nil && o.TokenIssuancePolicies != nil {
		return true
	}

	return false
}

// SetTokenIssuancePolicies gets a reference to the given []MicrosoftGraphTokenIssuancePolicy and assigns it to the TokenIssuancePolicies field.
func (o *PolicyRoot) SetTokenIssuancePolicies(v []MicrosoftGraphTokenIssuancePolicy) {
	o.TokenIssuancePolicies = &v
}

// GetTokenLifetimePolicies returns the TokenLifetimePolicies field value if set, zero value otherwise.
func (o *PolicyRoot) GetTokenLifetimePolicies() []MicrosoftGraphTokenLifetimePolicy {
	if o == nil || o.TokenLifetimePolicies == nil {
		var ret []MicrosoftGraphTokenLifetimePolicy
		return ret
	}
	return *o.TokenLifetimePolicies
}

// GetTokenLifetimePoliciesOk returns a tuple with the TokenLifetimePolicies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyRoot) GetTokenLifetimePoliciesOk() (*[]MicrosoftGraphTokenLifetimePolicy, bool) {
	if o == nil || o.TokenLifetimePolicies == nil {
		return nil, false
	}
	return o.TokenLifetimePolicies, true
}

// HasTokenLifetimePolicies returns a boolean if a field has been set.
func (o *PolicyRoot) HasTokenLifetimePolicies() bool {
	if o != nil && o.TokenLifetimePolicies != nil {
		return true
	}

	return false
}

// SetTokenLifetimePolicies gets a reference to the given []MicrosoftGraphTokenLifetimePolicy and assigns it to the TokenLifetimePolicies field.
func (o *PolicyRoot) SetTokenLifetimePolicies(v []MicrosoftGraphTokenLifetimePolicy) {
	o.TokenLifetimePolicies = &v
}

// GetFeatureRolloutPolicies returns the FeatureRolloutPolicies field value if set, zero value otherwise.
func (o *PolicyRoot) GetFeatureRolloutPolicies() []MicrosoftGraphFeatureRolloutPolicy {
	if o == nil || o.FeatureRolloutPolicies == nil {
		var ret []MicrosoftGraphFeatureRolloutPolicy
		return ret
	}
	return *o.FeatureRolloutPolicies
}

// GetFeatureRolloutPoliciesOk returns a tuple with the FeatureRolloutPolicies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyRoot) GetFeatureRolloutPoliciesOk() (*[]MicrosoftGraphFeatureRolloutPolicy, bool) {
	if o == nil || o.FeatureRolloutPolicies == nil {
		return nil, false
	}
	return o.FeatureRolloutPolicies, true
}

// HasFeatureRolloutPolicies returns a boolean if a field has been set.
func (o *PolicyRoot) HasFeatureRolloutPolicies() bool {
	if o != nil && o.FeatureRolloutPolicies != nil {
		return true
	}

	return false
}

// SetFeatureRolloutPolicies gets a reference to the given []MicrosoftGraphFeatureRolloutPolicy and assigns it to the FeatureRolloutPolicies field.
func (o *PolicyRoot) SetFeatureRolloutPolicies(v []MicrosoftGraphFeatureRolloutPolicy) {
	o.FeatureRolloutPolicies = &v
}

// GetAdminConsentRequestPolicy returns the AdminConsentRequestPolicy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PolicyRoot) GetAdminConsentRequestPolicy() AnyOfmicrosoftGraphAdminConsentRequestPolicy {
	if o == nil  {
		var ret AnyOfmicrosoftGraphAdminConsentRequestPolicy
		return ret
	}
	return o.AdminConsentRequestPolicy
}

// GetAdminConsentRequestPolicyOk returns a tuple with the AdminConsentRequestPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PolicyRoot) GetAdminConsentRequestPolicyOk() (*AnyOfmicrosoftGraphAdminConsentRequestPolicy, bool) {
	if o == nil || o.AdminConsentRequestPolicy == nil {
		return nil, false
	}
	return &o.AdminConsentRequestPolicy, true
}

// HasAdminConsentRequestPolicy returns a boolean if a field has been set.
func (o *PolicyRoot) HasAdminConsentRequestPolicy() bool {
	if o != nil && o.AdminConsentRequestPolicy != nil {
		return true
	}

	return false
}

// SetAdminConsentRequestPolicy gets a reference to the given AnyOfmicrosoftGraphAdminConsentRequestPolicy and assigns it to the AdminConsentRequestPolicy field.
func (o *PolicyRoot) SetAdminConsentRequestPolicy(v AnyOfmicrosoftGraphAdminConsentRequestPolicy) {
	o.AdminConsentRequestPolicy = v
}

// GetConditionalAccessPolicies returns the ConditionalAccessPolicies field value if set, zero value otherwise.
func (o *PolicyRoot) GetConditionalAccessPolicies() []MicrosoftGraphConditionalAccessPolicy {
	if o == nil || o.ConditionalAccessPolicies == nil {
		var ret []MicrosoftGraphConditionalAccessPolicy
		return ret
	}
	return *o.ConditionalAccessPolicies
}

// GetConditionalAccessPoliciesOk returns a tuple with the ConditionalAccessPolicies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyRoot) GetConditionalAccessPoliciesOk() (*[]MicrosoftGraphConditionalAccessPolicy, bool) {
	if o == nil || o.ConditionalAccessPolicies == nil {
		return nil, false
	}
	return o.ConditionalAccessPolicies, true
}

// HasConditionalAccessPolicies returns a boolean if a field has been set.
func (o *PolicyRoot) HasConditionalAccessPolicies() bool {
	if o != nil && o.ConditionalAccessPolicies != nil {
		return true
	}

	return false
}

// SetConditionalAccessPolicies gets a reference to the given []MicrosoftGraphConditionalAccessPolicy and assigns it to the ConditionalAccessPolicies field.
func (o *PolicyRoot) SetConditionalAccessPolicies(v []MicrosoftGraphConditionalAccessPolicy) {
	o.ConditionalAccessPolicies = &v
}

// GetIdentitySecurityDefaultsEnforcementPolicy returns the IdentitySecurityDefaultsEnforcementPolicy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PolicyRoot) GetIdentitySecurityDefaultsEnforcementPolicy() AnyOfmicrosoftGraphIdentitySecurityDefaultsEnforcementPolicy {
	if o == nil  {
		var ret AnyOfmicrosoftGraphIdentitySecurityDefaultsEnforcementPolicy
		return ret
	}
	return o.IdentitySecurityDefaultsEnforcementPolicy
}

// GetIdentitySecurityDefaultsEnforcementPolicyOk returns a tuple with the IdentitySecurityDefaultsEnforcementPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PolicyRoot) GetIdentitySecurityDefaultsEnforcementPolicyOk() (*AnyOfmicrosoftGraphIdentitySecurityDefaultsEnforcementPolicy, bool) {
	if o == nil || o.IdentitySecurityDefaultsEnforcementPolicy == nil {
		return nil, false
	}
	return &o.IdentitySecurityDefaultsEnforcementPolicy, true
}

// HasIdentitySecurityDefaultsEnforcementPolicy returns a boolean if a field has been set.
func (o *PolicyRoot) HasIdentitySecurityDefaultsEnforcementPolicy() bool {
	if o != nil && o.IdentitySecurityDefaultsEnforcementPolicy != nil {
		return true
	}

	return false
}

// SetIdentitySecurityDefaultsEnforcementPolicy gets a reference to the given AnyOfmicrosoftGraphIdentitySecurityDefaultsEnforcementPolicy and assigns it to the IdentitySecurityDefaultsEnforcementPolicy field.
func (o *PolicyRoot) SetIdentitySecurityDefaultsEnforcementPolicy(v AnyOfmicrosoftGraphIdentitySecurityDefaultsEnforcementPolicy) {
	o.IdentitySecurityDefaultsEnforcementPolicy = v
}

func (o PolicyRoot) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AuthenticationMethodsPolicy != nil {
		toSerialize["authenticationMethodsPolicy"] = o.AuthenticationMethodsPolicy
	}
	if o.AuthenticationFlowsPolicy != nil {
		toSerialize["authenticationFlowsPolicy"] = o.AuthenticationFlowsPolicy
	}
	if o.ActivityBasedTimeoutPolicies != nil {
		toSerialize["activityBasedTimeoutPolicies"] = o.ActivityBasedTimeoutPolicies
	}
	if o.AuthorizationPolicy != nil {
		toSerialize["authorizationPolicy"] = o.AuthorizationPolicy
	}
	if o.ClaimsMappingPolicies != nil {
		toSerialize["claimsMappingPolicies"] = o.ClaimsMappingPolicies
	}
	if o.HomeRealmDiscoveryPolicies != nil {
		toSerialize["homeRealmDiscoveryPolicies"] = o.HomeRealmDiscoveryPolicies
	}
	if o.PermissionGrantPolicies != nil {
		toSerialize["permissionGrantPolicies"] = o.PermissionGrantPolicies
	}
	if o.TokenIssuancePolicies != nil {
		toSerialize["tokenIssuancePolicies"] = o.TokenIssuancePolicies
	}
	if o.TokenLifetimePolicies != nil {
		toSerialize["tokenLifetimePolicies"] = o.TokenLifetimePolicies
	}
	if o.FeatureRolloutPolicies != nil {
		toSerialize["featureRolloutPolicies"] = o.FeatureRolloutPolicies
	}
	if o.AdminConsentRequestPolicy != nil {
		toSerialize["adminConsentRequestPolicy"] = o.AdminConsentRequestPolicy
	}
	if o.ConditionalAccessPolicies != nil {
		toSerialize["conditionalAccessPolicies"] = o.ConditionalAccessPolicies
	}
	if o.IdentitySecurityDefaultsEnforcementPolicy != nil {
		toSerialize["identitySecurityDefaultsEnforcementPolicy"] = o.IdentitySecurityDefaultsEnforcementPolicy
	}
	return json.Marshal(toSerialize)
}

type NullablePolicyRoot struct {
	value *PolicyRoot
	isSet bool
}

func (v NullablePolicyRoot) Get() *PolicyRoot {
	return v.value
}

func (v *NullablePolicyRoot) Set(val *PolicyRoot) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicyRoot) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicyRoot) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicyRoot(val *PolicyRoot) *NullablePolicyRoot {
	return &NullablePolicyRoot{value: val, isSet: true}
}

func (v NullablePolicyRoot) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicyRoot) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


