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

// MicrosoftGraphApiApplication struct for MicrosoftGraphApiApplication
type MicrosoftGraphApiApplication struct {
	// When true, allows an application to use claims mapping without specifying a custom signing key.
	AcceptMappedClaims NullableBool `json:"acceptMappedClaims,omitempty"`
	// Used for bundling consent if you have a solution that contains two parts: a client app and a custom web API app. If you set the appID of the client app to this value, the user only consents once to the client app. Azure AD knows that consenting to the client means implicitly consenting to the web API and automatically provisions service principals for both APIs at the same time. Both the client and the web API app must be registered in the same tenant.
	KnownClientApplications *[]*string `json:"knownClientApplications,omitempty"`
	// The definition of the delegated permissions exposed by the web API represented by this application registration. These delegated permissions may be requested by a client application, and may be granted by users or administrators during consent. Delegated permissions are sometimes referred to as OAuth 2.0 scopes.
	Oauth2PermissionScopes *[]MicrosoftGraphPermissionScope `json:"oauth2PermissionScopes,omitempty"`
	// Lists the client applications that are pre-authorized with the specified delegated permissions to access this application's APIs. Users are not required to consent to any pre-authorized application (for the permissions specified). However, any additional permissions not listed in preAuthorizedApplications (requested through incremental consent for example) will require user consent.
	PreAuthorizedApplications *[]*AnyOfmicrosoftGraphPreAuthorizedApplication `json:"preAuthorizedApplications,omitempty"`
	// Specifies the access token version expected by this resource. This changes the version and format of the JWT produced independent of the endpoint or client used to request the access token.  The endpoint used, v1.0 or v2.0, is chosen by the client and only impacts the version of id_tokens. Resources need to explicitly configure requestedAccessTokenVersion to indicate the supported access token format.  Possible values for requestedAccessTokenVersion are 1, 2, or null. If the value is null, this defaults to 1, which corresponds to the v1.0 endpoint.  If signInAudience on the application is configured as AzureADandPersonalMicrosoftAccount, the value for this property must be 2
	RequestedAccessTokenVersion NullableInt32 `json:"requestedAccessTokenVersion,omitempty"`
}

// NewMicrosoftGraphApiApplication instantiates a new MicrosoftGraphApiApplication object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphApiApplication() *MicrosoftGraphApiApplication {
	this := MicrosoftGraphApiApplication{}
	return &this
}

// NewMicrosoftGraphApiApplicationWithDefaults instantiates a new MicrosoftGraphApiApplication object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphApiApplicationWithDefaults() *MicrosoftGraphApiApplication {
	this := MicrosoftGraphApiApplication{}
	return &this
}

// GetAcceptMappedClaims returns the AcceptMappedClaims field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphApiApplication) GetAcceptMappedClaims() bool {
	if o == nil || o.AcceptMappedClaims.Get() == nil {
		var ret bool
		return ret
	}
	return *o.AcceptMappedClaims.Get()
}

// GetAcceptMappedClaimsOk returns a tuple with the AcceptMappedClaims field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphApiApplication) GetAcceptMappedClaimsOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AcceptMappedClaims.Get(), o.AcceptMappedClaims.IsSet()
}

// HasAcceptMappedClaims returns a boolean if a field has been set.
func (o *MicrosoftGraphApiApplication) HasAcceptMappedClaims() bool {
	if o != nil && o.AcceptMappedClaims.IsSet() {
		return true
	}

	return false
}

// SetAcceptMappedClaims gets a reference to the given NullableBool and assigns it to the AcceptMappedClaims field.
func (o *MicrosoftGraphApiApplication) SetAcceptMappedClaims(v bool) {
	o.AcceptMappedClaims.Set(&v)
}
// SetAcceptMappedClaimsNil sets the value for AcceptMappedClaims to be an explicit nil
func (o *MicrosoftGraphApiApplication) SetAcceptMappedClaimsNil() {
	o.AcceptMappedClaims.Set(nil)
}

// UnsetAcceptMappedClaims ensures that no value is present for AcceptMappedClaims, not even an explicit nil
func (o *MicrosoftGraphApiApplication) UnsetAcceptMappedClaims() {
	o.AcceptMappedClaims.Unset()
}

// GetKnownClientApplications returns the KnownClientApplications field value if set, zero value otherwise.
func (o *MicrosoftGraphApiApplication) GetKnownClientApplications() []*string {
	if o == nil || o.KnownClientApplications == nil {
		var ret []*string
		return ret
	}
	return *o.KnownClientApplications
}

// GetKnownClientApplicationsOk returns a tuple with the KnownClientApplications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphApiApplication) GetKnownClientApplicationsOk() (*[]*string, bool) {
	if o == nil || o.KnownClientApplications == nil {
		return nil, false
	}
	return o.KnownClientApplications, true
}

// HasKnownClientApplications returns a boolean if a field has been set.
func (o *MicrosoftGraphApiApplication) HasKnownClientApplications() bool {
	if o != nil && o.KnownClientApplications != nil {
		return true
	}

	return false
}

// SetKnownClientApplications gets a reference to the given []*string and assigns it to the KnownClientApplications field.
func (o *MicrosoftGraphApiApplication) SetKnownClientApplications(v []*string) {
	o.KnownClientApplications = &v
}

// GetOauth2PermissionScopes returns the Oauth2PermissionScopes field value if set, zero value otherwise.
func (o *MicrosoftGraphApiApplication) GetOauth2PermissionScopes() []MicrosoftGraphPermissionScope {
	if o == nil || o.Oauth2PermissionScopes == nil {
		var ret []MicrosoftGraphPermissionScope
		return ret
	}
	return *o.Oauth2PermissionScopes
}

// GetOauth2PermissionScopesOk returns a tuple with the Oauth2PermissionScopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphApiApplication) GetOauth2PermissionScopesOk() (*[]MicrosoftGraphPermissionScope, bool) {
	if o == nil || o.Oauth2PermissionScopes == nil {
		return nil, false
	}
	return o.Oauth2PermissionScopes, true
}

// HasOauth2PermissionScopes returns a boolean if a field has been set.
func (o *MicrosoftGraphApiApplication) HasOauth2PermissionScopes() bool {
	if o != nil && o.Oauth2PermissionScopes != nil {
		return true
	}

	return false
}

// SetOauth2PermissionScopes gets a reference to the given []MicrosoftGraphPermissionScope and assigns it to the Oauth2PermissionScopes field.
func (o *MicrosoftGraphApiApplication) SetOauth2PermissionScopes(v []MicrosoftGraphPermissionScope) {
	o.Oauth2PermissionScopes = &v
}

// GetPreAuthorizedApplications returns the PreAuthorizedApplications field value if set, zero value otherwise.
func (o *MicrosoftGraphApiApplication) GetPreAuthorizedApplications() []*AnyOfmicrosoftGraphPreAuthorizedApplication {
	if o == nil || o.PreAuthorizedApplications == nil {
		var ret []*AnyOfmicrosoftGraphPreAuthorizedApplication
		return ret
	}
	return *o.PreAuthorizedApplications
}

// GetPreAuthorizedApplicationsOk returns a tuple with the PreAuthorizedApplications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphApiApplication) GetPreAuthorizedApplicationsOk() (*[]*AnyOfmicrosoftGraphPreAuthorizedApplication, bool) {
	if o == nil || o.PreAuthorizedApplications == nil {
		return nil, false
	}
	return o.PreAuthorizedApplications, true
}

// HasPreAuthorizedApplications returns a boolean if a field has been set.
func (o *MicrosoftGraphApiApplication) HasPreAuthorizedApplications() bool {
	if o != nil && o.PreAuthorizedApplications != nil {
		return true
	}

	return false
}

// SetPreAuthorizedApplications gets a reference to the given []*AnyOfmicrosoftGraphPreAuthorizedApplication and assigns it to the PreAuthorizedApplications field.
func (o *MicrosoftGraphApiApplication) SetPreAuthorizedApplications(v []*AnyOfmicrosoftGraphPreAuthorizedApplication) {
	o.PreAuthorizedApplications = &v
}

// GetRequestedAccessTokenVersion returns the RequestedAccessTokenVersion field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphApiApplication) GetRequestedAccessTokenVersion() int32 {
	if o == nil || o.RequestedAccessTokenVersion.Get() == nil {
		var ret int32
		return ret
	}
	return *o.RequestedAccessTokenVersion.Get()
}

// GetRequestedAccessTokenVersionOk returns a tuple with the RequestedAccessTokenVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphApiApplication) GetRequestedAccessTokenVersionOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RequestedAccessTokenVersion.Get(), o.RequestedAccessTokenVersion.IsSet()
}

// HasRequestedAccessTokenVersion returns a boolean if a field has been set.
func (o *MicrosoftGraphApiApplication) HasRequestedAccessTokenVersion() bool {
	if o != nil && o.RequestedAccessTokenVersion.IsSet() {
		return true
	}

	return false
}

// SetRequestedAccessTokenVersion gets a reference to the given NullableInt32 and assigns it to the RequestedAccessTokenVersion field.
func (o *MicrosoftGraphApiApplication) SetRequestedAccessTokenVersion(v int32) {
	o.RequestedAccessTokenVersion.Set(&v)
}
// SetRequestedAccessTokenVersionNil sets the value for RequestedAccessTokenVersion to be an explicit nil
func (o *MicrosoftGraphApiApplication) SetRequestedAccessTokenVersionNil() {
	o.RequestedAccessTokenVersion.Set(nil)
}

// UnsetRequestedAccessTokenVersion ensures that no value is present for RequestedAccessTokenVersion, not even an explicit nil
func (o *MicrosoftGraphApiApplication) UnsetRequestedAccessTokenVersion() {
	o.RequestedAccessTokenVersion.Unset()
}

func (o MicrosoftGraphApiApplication) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AcceptMappedClaims.IsSet() {
		toSerialize["acceptMappedClaims"] = o.AcceptMappedClaims.Get()
	}
	if o.KnownClientApplications != nil {
		toSerialize["knownClientApplications"] = o.KnownClientApplications
	}
	if o.Oauth2PermissionScopes != nil {
		toSerialize["oauth2PermissionScopes"] = o.Oauth2PermissionScopes
	}
	if o.PreAuthorizedApplications != nil {
		toSerialize["preAuthorizedApplications"] = o.PreAuthorizedApplications
	}
	if o.RequestedAccessTokenVersion.IsSet() {
		toSerialize["requestedAccessTokenVersion"] = o.RequestedAccessTokenVersion.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphApiApplication struct {
	value *MicrosoftGraphApiApplication
	isSet bool
}

func (v NullableMicrosoftGraphApiApplication) Get() *MicrosoftGraphApiApplication {
	return v.value
}

func (v *NullableMicrosoftGraphApiApplication) Set(val *MicrosoftGraphApiApplication) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphApiApplication) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphApiApplication) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphApiApplication(val *MicrosoftGraphApiApplication) *NullableMicrosoftGraphApiApplication {
	return &NullableMicrosoftGraphApiApplication{value: val, isSet: true}
}

func (v NullableMicrosoftGraphApiApplication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphApiApplication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


