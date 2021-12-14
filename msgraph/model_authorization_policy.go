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

// AuthorizationPolicy struct for AuthorizationPolicy
type AuthorizationPolicy struct {
	// Indicates whether users can sign up for email based subscriptions.
	AllowedToSignUpEmailBasedSubscriptions *bool `json:"allowedToSignUpEmailBasedSubscriptions,omitempty"`
	// Indicates whether the Self-Serve Password Reset feature can be used by users on the tenant.
	AllowedToUseSSPR *bool `json:"allowedToUseSSPR,omitempty"`
	// Indicates whether a user can join the tenant by email validation.
	AllowEmailVerifiedUsersToJoinOrganization *bool `json:"allowEmailVerifiedUsersToJoinOrganization,omitempty"`
	// Indicates who can invite external users to the organization. Possible values are: none, adminsAndGuestInviters, adminsGuestInvitersAndAllMembers, everyone.  everyone is the default setting for all cloud environments except US Government. See more in the table below.
	AllowInvitesFrom AnyOfmicrosoftGraphAllowInvitesFrom `json:"allowInvitesFrom,omitempty"`
	// To disable the use of MSOL PowerShell set this property to true. This will also disable user-based access to the legacy service endpoint used by MSOL PowerShell. This does not affect Azure AD Connect or Microsoft Graph.
	BlockMsolPowerShell NullableBool `json:"blockMsolPowerShell,omitempty"`
	DefaultUserRolePermissions *MicrosoftGraphDefaultUserRolePermissions `json:"defaultUserRolePermissions,omitempty"`
	// Represents role templateId for the role that should be granted to guest user. Currently following roles are supported:  User (a0b1b346-4d3e-4e8b-98f8-753987be4970), Guest User (10dae51f-b6af-4016-8d66-8c2a99b929b3), and Restricted Guest User (2af84b1e-32c8-42b7-82bc-daa82404023b).
	GuestUserRoleId NullableString `json:"guestUserRoleId,omitempty"`
}

// NewAuthorizationPolicy instantiates a new AuthorizationPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthorizationPolicy() *AuthorizationPolicy {
	this := AuthorizationPolicy{}
	return &this
}

// NewAuthorizationPolicyWithDefaults instantiates a new AuthorizationPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthorizationPolicyWithDefaults() *AuthorizationPolicy {
	this := AuthorizationPolicy{}
	return &this
}

// GetAllowedToSignUpEmailBasedSubscriptions returns the AllowedToSignUpEmailBasedSubscriptions field value if set, zero value otherwise.
func (o *AuthorizationPolicy) GetAllowedToSignUpEmailBasedSubscriptions() bool {
	if o == nil || o.AllowedToSignUpEmailBasedSubscriptions == nil {
		var ret bool
		return ret
	}
	return *o.AllowedToSignUpEmailBasedSubscriptions
}

// GetAllowedToSignUpEmailBasedSubscriptionsOk returns a tuple with the AllowedToSignUpEmailBasedSubscriptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationPolicy) GetAllowedToSignUpEmailBasedSubscriptionsOk() (*bool, bool) {
	if o == nil || o.AllowedToSignUpEmailBasedSubscriptions == nil {
		return nil, false
	}
	return o.AllowedToSignUpEmailBasedSubscriptions, true
}

// HasAllowedToSignUpEmailBasedSubscriptions returns a boolean if a field has been set.
func (o *AuthorizationPolicy) HasAllowedToSignUpEmailBasedSubscriptions() bool {
	if o != nil && o.AllowedToSignUpEmailBasedSubscriptions != nil {
		return true
	}

	return false
}

// SetAllowedToSignUpEmailBasedSubscriptions gets a reference to the given bool and assigns it to the AllowedToSignUpEmailBasedSubscriptions field.
func (o *AuthorizationPolicy) SetAllowedToSignUpEmailBasedSubscriptions(v bool) {
	o.AllowedToSignUpEmailBasedSubscriptions = &v
}

// GetAllowedToUseSSPR returns the AllowedToUseSSPR field value if set, zero value otherwise.
func (o *AuthorizationPolicy) GetAllowedToUseSSPR() bool {
	if o == nil || o.AllowedToUseSSPR == nil {
		var ret bool
		return ret
	}
	return *o.AllowedToUseSSPR
}

// GetAllowedToUseSSPROk returns a tuple with the AllowedToUseSSPR field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationPolicy) GetAllowedToUseSSPROk() (*bool, bool) {
	if o == nil || o.AllowedToUseSSPR == nil {
		return nil, false
	}
	return o.AllowedToUseSSPR, true
}

// HasAllowedToUseSSPR returns a boolean if a field has been set.
func (o *AuthorizationPolicy) HasAllowedToUseSSPR() bool {
	if o != nil && o.AllowedToUseSSPR != nil {
		return true
	}

	return false
}

// SetAllowedToUseSSPR gets a reference to the given bool and assigns it to the AllowedToUseSSPR field.
func (o *AuthorizationPolicy) SetAllowedToUseSSPR(v bool) {
	o.AllowedToUseSSPR = &v
}

// GetAllowEmailVerifiedUsersToJoinOrganization returns the AllowEmailVerifiedUsersToJoinOrganization field value if set, zero value otherwise.
func (o *AuthorizationPolicy) GetAllowEmailVerifiedUsersToJoinOrganization() bool {
	if o == nil || o.AllowEmailVerifiedUsersToJoinOrganization == nil {
		var ret bool
		return ret
	}
	return *o.AllowEmailVerifiedUsersToJoinOrganization
}

// GetAllowEmailVerifiedUsersToJoinOrganizationOk returns a tuple with the AllowEmailVerifiedUsersToJoinOrganization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationPolicy) GetAllowEmailVerifiedUsersToJoinOrganizationOk() (*bool, bool) {
	if o == nil || o.AllowEmailVerifiedUsersToJoinOrganization == nil {
		return nil, false
	}
	return o.AllowEmailVerifiedUsersToJoinOrganization, true
}

// HasAllowEmailVerifiedUsersToJoinOrganization returns a boolean if a field has been set.
func (o *AuthorizationPolicy) HasAllowEmailVerifiedUsersToJoinOrganization() bool {
	if o != nil && o.AllowEmailVerifiedUsersToJoinOrganization != nil {
		return true
	}

	return false
}

// SetAllowEmailVerifiedUsersToJoinOrganization gets a reference to the given bool and assigns it to the AllowEmailVerifiedUsersToJoinOrganization field.
func (o *AuthorizationPolicy) SetAllowEmailVerifiedUsersToJoinOrganization(v bool) {
	o.AllowEmailVerifiedUsersToJoinOrganization = &v
}

// GetAllowInvitesFrom returns the AllowInvitesFrom field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuthorizationPolicy) GetAllowInvitesFrom() AnyOfmicrosoftGraphAllowInvitesFrom {
	if o == nil  {
		var ret AnyOfmicrosoftGraphAllowInvitesFrom
		return ret
	}
	return o.AllowInvitesFrom
}

// GetAllowInvitesFromOk returns a tuple with the AllowInvitesFrom field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuthorizationPolicy) GetAllowInvitesFromOk() (*AnyOfmicrosoftGraphAllowInvitesFrom, bool) {
	if o == nil || o.AllowInvitesFrom == nil {
		return nil, false
	}
	return &o.AllowInvitesFrom, true
}

// HasAllowInvitesFrom returns a boolean if a field has been set.
func (o *AuthorizationPolicy) HasAllowInvitesFrom() bool {
	if o != nil && o.AllowInvitesFrom != nil {
		return true
	}

	return false
}

// SetAllowInvitesFrom gets a reference to the given AnyOfmicrosoftGraphAllowInvitesFrom and assigns it to the AllowInvitesFrom field.
func (o *AuthorizationPolicy) SetAllowInvitesFrom(v AnyOfmicrosoftGraphAllowInvitesFrom) {
	o.AllowInvitesFrom = v
}

// GetBlockMsolPowerShell returns the BlockMsolPowerShell field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuthorizationPolicy) GetBlockMsolPowerShell() bool {
	if o == nil || o.BlockMsolPowerShell.Get() == nil {
		var ret bool
		return ret
	}
	return *o.BlockMsolPowerShell.Get()
}

// GetBlockMsolPowerShellOk returns a tuple with the BlockMsolPowerShell field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuthorizationPolicy) GetBlockMsolPowerShellOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.BlockMsolPowerShell.Get(), o.BlockMsolPowerShell.IsSet()
}

// HasBlockMsolPowerShell returns a boolean if a field has been set.
func (o *AuthorizationPolicy) HasBlockMsolPowerShell() bool {
	if o != nil && o.BlockMsolPowerShell.IsSet() {
		return true
	}

	return false
}

// SetBlockMsolPowerShell gets a reference to the given NullableBool and assigns it to the BlockMsolPowerShell field.
func (o *AuthorizationPolicy) SetBlockMsolPowerShell(v bool) {
	o.BlockMsolPowerShell.Set(&v)
}
// SetBlockMsolPowerShellNil sets the value for BlockMsolPowerShell to be an explicit nil
func (o *AuthorizationPolicy) SetBlockMsolPowerShellNil() {
	o.BlockMsolPowerShell.Set(nil)
}

// UnsetBlockMsolPowerShell ensures that no value is present for BlockMsolPowerShell, not even an explicit nil
func (o *AuthorizationPolicy) UnsetBlockMsolPowerShell() {
	o.BlockMsolPowerShell.Unset()
}

// GetDefaultUserRolePermissions returns the DefaultUserRolePermissions field value if set, zero value otherwise.
func (o *AuthorizationPolicy) GetDefaultUserRolePermissions() MicrosoftGraphDefaultUserRolePermissions {
	if o == nil || o.DefaultUserRolePermissions == nil {
		var ret MicrosoftGraphDefaultUserRolePermissions
		return ret
	}
	return *o.DefaultUserRolePermissions
}

// GetDefaultUserRolePermissionsOk returns a tuple with the DefaultUserRolePermissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationPolicy) GetDefaultUserRolePermissionsOk() (*MicrosoftGraphDefaultUserRolePermissions, bool) {
	if o == nil || o.DefaultUserRolePermissions == nil {
		return nil, false
	}
	return o.DefaultUserRolePermissions, true
}

// HasDefaultUserRolePermissions returns a boolean if a field has been set.
func (o *AuthorizationPolicy) HasDefaultUserRolePermissions() bool {
	if o != nil && o.DefaultUserRolePermissions != nil {
		return true
	}

	return false
}

// SetDefaultUserRolePermissions gets a reference to the given MicrosoftGraphDefaultUserRolePermissions and assigns it to the DefaultUserRolePermissions field.
func (o *AuthorizationPolicy) SetDefaultUserRolePermissions(v MicrosoftGraphDefaultUserRolePermissions) {
	o.DefaultUserRolePermissions = &v
}

// GetGuestUserRoleId returns the GuestUserRoleId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuthorizationPolicy) GetGuestUserRoleId() string {
	if o == nil || o.GuestUserRoleId.Get() == nil {
		var ret string
		return ret
	}
	return *o.GuestUserRoleId.Get()
}

// GetGuestUserRoleIdOk returns a tuple with the GuestUserRoleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuthorizationPolicy) GetGuestUserRoleIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.GuestUserRoleId.Get(), o.GuestUserRoleId.IsSet()
}

// HasGuestUserRoleId returns a boolean if a field has been set.
func (o *AuthorizationPolicy) HasGuestUserRoleId() bool {
	if o != nil && o.GuestUserRoleId.IsSet() {
		return true
	}

	return false
}

// SetGuestUserRoleId gets a reference to the given NullableString and assigns it to the GuestUserRoleId field.
func (o *AuthorizationPolicy) SetGuestUserRoleId(v string) {
	o.GuestUserRoleId.Set(&v)
}
// SetGuestUserRoleIdNil sets the value for GuestUserRoleId to be an explicit nil
func (o *AuthorizationPolicy) SetGuestUserRoleIdNil() {
	o.GuestUserRoleId.Set(nil)
}

// UnsetGuestUserRoleId ensures that no value is present for GuestUserRoleId, not even an explicit nil
func (o *AuthorizationPolicy) UnsetGuestUserRoleId() {
	o.GuestUserRoleId.Unset()
}

func (o AuthorizationPolicy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AllowedToSignUpEmailBasedSubscriptions != nil {
		toSerialize["allowedToSignUpEmailBasedSubscriptions"] = o.AllowedToSignUpEmailBasedSubscriptions
	}
	if o.AllowedToUseSSPR != nil {
		toSerialize["allowedToUseSSPR"] = o.AllowedToUseSSPR
	}
	if o.AllowEmailVerifiedUsersToJoinOrganization != nil {
		toSerialize["allowEmailVerifiedUsersToJoinOrganization"] = o.AllowEmailVerifiedUsersToJoinOrganization
	}
	if o.AllowInvitesFrom != nil {
		toSerialize["allowInvitesFrom"] = o.AllowInvitesFrom
	}
	if o.BlockMsolPowerShell.IsSet() {
		toSerialize["blockMsolPowerShell"] = o.BlockMsolPowerShell.Get()
	}
	if o.DefaultUserRolePermissions != nil {
		toSerialize["defaultUserRolePermissions"] = o.DefaultUserRolePermissions
	}
	if o.GuestUserRoleId.IsSet() {
		toSerialize["guestUserRoleId"] = o.GuestUserRoleId.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableAuthorizationPolicy struct {
	value *AuthorizationPolicy
	isSet bool
}

func (v NullableAuthorizationPolicy) Get() *AuthorizationPolicy {
	return v.value
}

func (v *NullableAuthorizationPolicy) Set(val *AuthorizationPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthorizationPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthorizationPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthorizationPolicy(val *AuthorizationPolicy) *NullableAuthorizationPolicy {
	return &NullableAuthorizationPolicy{value: val, isSet: true}
}

func (v NullableAuthorizationPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthorizationPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

