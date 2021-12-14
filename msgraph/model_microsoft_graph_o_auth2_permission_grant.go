/*
Partial Graph API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// MicrosoftGraphOAuth2PermissionGrant struct for MicrosoftGraphOAuth2PermissionGrant
type MicrosoftGraphOAuth2PermissionGrant struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
	// The id of the client service principal for the application which is authorized to act on behalf of a signed-in user when accessing an API. Required. Supports $filter (eq only).
	ClientId *string `json:"clientId,omitempty"`
	// Indicates if authorization is granted for the client application to impersonate all users or only a specific user. AllPrincipals indicates authorization to impersonate all users. Principal indicates authorization to impersonate a specific user. Consent on behalf of all users can be granted by an administrator. Non-admin users may be authorized to consent on behalf of themselves in some cases, for some delegated permissions. Required. Supports $filter (eq only).
	ConsentType NullableString `json:"consentType,omitempty"`
	// The id of the user on behalf of whom the client is authorized to access the resource, when consentType is Principal. If consentType is AllPrincipals this value is null. Required when consentType is Principal.
	PrincipalId NullableString `json:"principalId,omitempty"`
	// The id of the resource service principal to which access is authorized. This identifies the API which the client is authorized to attempt to call on behalf of a signed-in user.
	ResourceId *string `json:"resourceId,omitempty"`
	// A space-separated list of the claim values for delegated permissions which should be included in access tokens for the resource application (the API). For example, openid User.Read GroupMember.Read.All. Each claim value should match the value field of one of the delegated permissions defined by the API, listed in the publishedPermissionScopes property of the resource service principal.
	Scope NullableString `json:"scope,omitempty"`
}

// NewMicrosoftGraphOAuth2PermissionGrant instantiates a new MicrosoftGraphOAuth2PermissionGrant object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphOAuth2PermissionGrant() *MicrosoftGraphOAuth2PermissionGrant {
	this := MicrosoftGraphOAuth2PermissionGrant{}
	return &this
}

// NewMicrosoftGraphOAuth2PermissionGrantWithDefaults instantiates a new MicrosoftGraphOAuth2PermissionGrant object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphOAuth2PermissionGrantWithDefaults() *MicrosoftGraphOAuth2PermissionGrant {
	this := MicrosoftGraphOAuth2PermissionGrant{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MicrosoftGraphOAuth2PermissionGrant) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphOAuth2PermissionGrant) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphOAuth2PermissionGrant) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MicrosoftGraphOAuth2PermissionGrant) SetId(v string) {
	o.Id = &v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *MicrosoftGraphOAuth2PermissionGrant) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphOAuth2PermissionGrant) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *MicrosoftGraphOAuth2PermissionGrant) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *MicrosoftGraphOAuth2PermissionGrant) SetClientId(v string) {
	o.ClientId = &v
}

// GetConsentType returns the ConsentType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphOAuth2PermissionGrant) GetConsentType() string {
	if o == nil || o.ConsentType.Get() == nil {
		var ret string
		return ret
	}
	return *o.ConsentType.Get()
}

// GetConsentTypeOk returns a tuple with the ConsentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphOAuth2PermissionGrant) GetConsentTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ConsentType.Get(), o.ConsentType.IsSet()
}

// HasConsentType returns a boolean if a field has been set.
func (o *MicrosoftGraphOAuth2PermissionGrant) HasConsentType() bool {
	if o != nil && o.ConsentType.IsSet() {
		return true
	}

	return false
}

// SetConsentType gets a reference to the given NullableString and assigns it to the ConsentType field.
func (o *MicrosoftGraphOAuth2PermissionGrant) SetConsentType(v string) {
	o.ConsentType.Set(&v)
}
// SetConsentTypeNil sets the value for ConsentType to be an explicit nil
func (o *MicrosoftGraphOAuth2PermissionGrant) SetConsentTypeNil() {
	o.ConsentType.Set(nil)
}

// UnsetConsentType ensures that no value is present for ConsentType, not even an explicit nil
func (o *MicrosoftGraphOAuth2PermissionGrant) UnsetConsentType() {
	o.ConsentType.Unset()
}

// GetPrincipalId returns the PrincipalId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphOAuth2PermissionGrant) GetPrincipalId() string {
	if o == nil || o.PrincipalId.Get() == nil {
		var ret string
		return ret
	}
	return *o.PrincipalId.Get()
}

// GetPrincipalIdOk returns a tuple with the PrincipalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphOAuth2PermissionGrant) GetPrincipalIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PrincipalId.Get(), o.PrincipalId.IsSet()
}

// HasPrincipalId returns a boolean if a field has been set.
func (o *MicrosoftGraphOAuth2PermissionGrant) HasPrincipalId() bool {
	if o != nil && o.PrincipalId.IsSet() {
		return true
	}

	return false
}

// SetPrincipalId gets a reference to the given NullableString and assigns it to the PrincipalId field.
func (o *MicrosoftGraphOAuth2PermissionGrant) SetPrincipalId(v string) {
	o.PrincipalId.Set(&v)
}
// SetPrincipalIdNil sets the value for PrincipalId to be an explicit nil
func (o *MicrosoftGraphOAuth2PermissionGrant) SetPrincipalIdNil() {
	o.PrincipalId.Set(nil)
}

// UnsetPrincipalId ensures that no value is present for PrincipalId, not even an explicit nil
func (o *MicrosoftGraphOAuth2PermissionGrant) UnsetPrincipalId() {
	o.PrincipalId.Unset()
}

// GetResourceId returns the ResourceId field value if set, zero value otherwise.
func (o *MicrosoftGraphOAuth2PermissionGrant) GetResourceId() string {
	if o == nil || o.ResourceId == nil {
		var ret string
		return ret
	}
	return *o.ResourceId
}

// GetResourceIdOk returns a tuple with the ResourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphOAuth2PermissionGrant) GetResourceIdOk() (*string, bool) {
	if o == nil || o.ResourceId == nil {
		return nil, false
	}
	return o.ResourceId, true
}

// HasResourceId returns a boolean if a field has been set.
func (o *MicrosoftGraphOAuth2PermissionGrant) HasResourceId() bool {
	if o != nil && o.ResourceId != nil {
		return true
	}

	return false
}

// SetResourceId gets a reference to the given string and assigns it to the ResourceId field.
func (o *MicrosoftGraphOAuth2PermissionGrant) SetResourceId(v string) {
	o.ResourceId = &v
}

// GetScope returns the Scope field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphOAuth2PermissionGrant) GetScope() string {
	if o == nil || o.Scope.Get() == nil {
		var ret string
		return ret
	}
	return *o.Scope.Get()
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphOAuth2PermissionGrant) GetScopeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Scope.Get(), o.Scope.IsSet()
}

// HasScope returns a boolean if a field has been set.
func (o *MicrosoftGraphOAuth2PermissionGrant) HasScope() bool {
	if o != nil && o.Scope.IsSet() {
		return true
	}

	return false
}

// SetScope gets a reference to the given NullableString and assigns it to the Scope field.
func (o *MicrosoftGraphOAuth2PermissionGrant) SetScope(v string) {
	o.Scope.Set(&v)
}
// SetScopeNil sets the value for Scope to be an explicit nil
func (o *MicrosoftGraphOAuth2PermissionGrant) SetScopeNil() {
	o.Scope.Set(nil)
}

// UnsetScope ensures that no value is present for Scope, not even an explicit nil
func (o *MicrosoftGraphOAuth2PermissionGrant) UnsetScope() {
	o.Scope.Unset()
}

func (o MicrosoftGraphOAuth2PermissionGrant) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.ClientId != nil {
		toSerialize["clientId"] = o.ClientId
	}
	if o.ConsentType.IsSet() {
		toSerialize["consentType"] = o.ConsentType.Get()
	}
	if o.PrincipalId.IsSet() {
		toSerialize["principalId"] = o.PrincipalId.Get()
	}
	if o.ResourceId != nil {
		toSerialize["resourceId"] = o.ResourceId
	}
	if o.Scope.IsSet() {
		toSerialize["scope"] = o.Scope.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphOAuth2PermissionGrant struct {
	value *MicrosoftGraphOAuth2PermissionGrant
	isSet bool
}

func (v NullableMicrosoftGraphOAuth2PermissionGrant) Get() *MicrosoftGraphOAuth2PermissionGrant {
	return v.value
}

func (v *NullableMicrosoftGraphOAuth2PermissionGrant) Set(val *MicrosoftGraphOAuth2PermissionGrant) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphOAuth2PermissionGrant) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphOAuth2PermissionGrant) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphOAuth2PermissionGrant(val *MicrosoftGraphOAuth2PermissionGrant) *NullableMicrosoftGraphOAuth2PermissionGrant {
	return &NullableMicrosoftGraphOAuth2PermissionGrant{value: val, isSet: true}
}

func (v NullableMicrosoftGraphOAuth2PermissionGrant) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphOAuth2PermissionGrant) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


