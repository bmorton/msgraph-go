/*
Partial Graph API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package msgraph

import (
	"encoding/json"
	"time"
)

// MicrosoftGraphAuthorizationPolicy struct for MicrosoftGraphAuthorizationPolicy
type MicrosoftGraphAuthorizationPolicy struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
	DeletedDateTime NullableTime `json:"deletedDateTime,omitempty"`
	// Description for this policy. Required.
	Description NullableString `json:"description,omitempty"`
	// Display name for this policy. Required.
	DisplayName NullableString `json:"displayName,omitempty"`
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

// NewMicrosoftGraphAuthorizationPolicy instantiates a new MicrosoftGraphAuthorizationPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphAuthorizationPolicy() *MicrosoftGraphAuthorizationPolicy {
	this := MicrosoftGraphAuthorizationPolicy{}
	return &this
}

// NewMicrosoftGraphAuthorizationPolicyWithDefaults instantiates a new MicrosoftGraphAuthorizationPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphAuthorizationPolicyWithDefaults() *MicrosoftGraphAuthorizationPolicy {
	this := MicrosoftGraphAuthorizationPolicy{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MicrosoftGraphAuthorizationPolicy) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphAuthorizationPolicy) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphAuthorizationPolicy) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MicrosoftGraphAuthorizationPolicy) SetId(v string) {
	o.Id = &v
}

// GetDeletedDateTime returns the DeletedDateTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphAuthorizationPolicy) GetDeletedDateTime() time.Time {
	if o == nil || o.DeletedDateTime.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.DeletedDateTime.Get()
}

// GetDeletedDateTimeOk returns a tuple with the DeletedDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphAuthorizationPolicy) GetDeletedDateTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DeletedDateTime.Get(), o.DeletedDateTime.IsSet()
}

// HasDeletedDateTime returns a boolean if a field has been set.
func (o *MicrosoftGraphAuthorizationPolicy) HasDeletedDateTime() bool {
	if o != nil && o.DeletedDateTime.IsSet() {
		return true
	}

	return false
}

// SetDeletedDateTime gets a reference to the given NullableTime and assigns it to the DeletedDateTime field.
func (o *MicrosoftGraphAuthorizationPolicy) SetDeletedDateTime(v time.Time) {
	o.DeletedDateTime.Set(&v)
}
// SetDeletedDateTimeNil sets the value for DeletedDateTime to be an explicit nil
func (o *MicrosoftGraphAuthorizationPolicy) SetDeletedDateTimeNil() {
	o.DeletedDateTime.Set(nil)
}

// UnsetDeletedDateTime ensures that no value is present for DeletedDateTime, not even an explicit nil
func (o *MicrosoftGraphAuthorizationPolicy) UnsetDeletedDateTime() {
	o.DeletedDateTime.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphAuthorizationPolicy) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphAuthorizationPolicy) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *MicrosoftGraphAuthorizationPolicy) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *MicrosoftGraphAuthorizationPolicy) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *MicrosoftGraphAuthorizationPolicy) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *MicrosoftGraphAuthorizationPolicy) UnsetDescription() {
	o.Description.Unset()
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphAuthorizationPolicy) GetDisplayName() string {
	if o == nil || o.DisplayName.Get() == nil {
		var ret string
		return ret
	}
	return *o.DisplayName.Get()
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphAuthorizationPolicy) GetDisplayNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DisplayName.Get(), o.DisplayName.IsSet()
}

// HasDisplayName returns a boolean if a field has been set.
func (o *MicrosoftGraphAuthorizationPolicy) HasDisplayName() bool {
	if o != nil && o.DisplayName.IsSet() {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given NullableString and assigns it to the DisplayName field.
func (o *MicrosoftGraphAuthorizationPolicy) SetDisplayName(v string) {
	o.DisplayName.Set(&v)
}
// SetDisplayNameNil sets the value for DisplayName to be an explicit nil
func (o *MicrosoftGraphAuthorizationPolicy) SetDisplayNameNil() {
	o.DisplayName.Set(nil)
}

// UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
func (o *MicrosoftGraphAuthorizationPolicy) UnsetDisplayName() {
	o.DisplayName.Unset()
}

// GetAllowedToSignUpEmailBasedSubscriptions returns the AllowedToSignUpEmailBasedSubscriptions field value if set, zero value otherwise.
func (o *MicrosoftGraphAuthorizationPolicy) GetAllowedToSignUpEmailBasedSubscriptions() bool {
	if o == nil || o.AllowedToSignUpEmailBasedSubscriptions == nil {
		var ret bool
		return ret
	}
	return *o.AllowedToSignUpEmailBasedSubscriptions
}

// GetAllowedToSignUpEmailBasedSubscriptionsOk returns a tuple with the AllowedToSignUpEmailBasedSubscriptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphAuthorizationPolicy) GetAllowedToSignUpEmailBasedSubscriptionsOk() (*bool, bool) {
	if o == nil || o.AllowedToSignUpEmailBasedSubscriptions == nil {
		return nil, false
	}
	return o.AllowedToSignUpEmailBasedSubscriptions, true
}

// HasAllowedToSignUpEmailBasedSubscriptions returns a boolean if a field has been set.
func (o *MicrosoftGraphAuthorizationPolicy) HasAllowedToSignUpEmailBasedSubscriptions() bool {
	if o != nil && o.AllowedToSignUpEmailBasedSubscriptions != nil {
		return true
	}

	return false
}

// SetAllowedToSignUpEmailBasedSubscriptions gets a reference to the given bool and assigns it to the AllowedToSignUpEmailBasedSubscriptions field.
func (o *MicrosoftGraphAuthorizationPolicy) SetAllowedToSignUpEmailBasedSubscriptions(v bool) {
	o.AllowedToSignUpEmailBasedSubscriptions = &v
}

// GetAllowedToUseSSPR returns the AllowedToUseSSPR field value if set, zero value otherwise.
func (o *MicrosoftGraphAuthorizationPolicy) GetAllowedToUseSSPR() bool {
	if o == nil || o.AllowedToUseSSPR == nil {
		var ret bool
		return ret
	}
	return *o.AllowedToUseSSPR
}

// GetAllowedToUseSSPROk returns a tuple with the AllowedToUseSSPR field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphAuthorizationPolicy) GetAllowedToUseSSPROk() (*bool, bool) {
	if o == nil || o.AllowedToUseSSPR == nil {
		return nil, false
	}
	return o.AllowedToUseSSPR, true
}

// HasAllowedToUseSSPR returns a boolean if a field has been set.
func (o *MicrosoftGraphAuthorizationPolicy) HasAllowedToUseSSPR() bool {
	if o != nil && o.AllowedToUseSSPR != nil {
		return true
	}

	return false
}

// SetAllowedToUseSSPR gets a reference to the given bool and assigns it to the AllowedToUseSSPR field.
func (o *MicrosoftGraphAuthorizationPolicy) SetAllowedToUseSSPR(v bool) {
	o.AllowedToUseSSPR = &v
}

// GetAllowEmailVerifiedUsersToJoinOrganization returns the AllowEmailVerifiedUsersToJoinOrganization field value if set, zero value otherwise.
func (o *MicrosoftGraphAuthorizationPolicy) GetAllowEmailVerifiedUsersToJoinOrganization() bool {
	if o == nil || o.AllowEmailVerifiedUsersToJoinOrganization == nil {
		var ret bool
		return ret
	}
	return *o.AllowEmailVerifiedUsersToJoinOrganization
}

// GetAllowEmailVerifiedUsersToJoinOrganizationOk returns a tuple with the AllowEmailVerifiedUsersToJoinOrganization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphAuthorizationPolicy) GetAllowEmailVerifiedUsersToJoinOrganizationOk() (*bool, bool) {
	if o == nil || o.AllowEmailVerifiedUsersToJoinOrganization == nil {
		return nil, false
	}
	return o.AllowEmailVerifiedUsersToJoinOrganization, true
}

// HasAllowEmailVerifiedUsersToJoinOrganization returns a boolean if a field has been set.
func (o *MicrosoftGraphAuthorizationPolicy) HasAllowEmailVerifiedUsersToJoinOrganization() bool {
	if o != nil && o.AllowEmailVerifiedUsersToJoinOrganization != nil {
		return true
	}

	return false
}

// SetAllowEmailVerifiedUsersToJoinOrganization gets a reference to the given bool and assigns it to the AllowEmailVerifiedUsersToJoinOrganization field.
func (o *MicrosoftGraphAuthorizationPolicy) SetAllowEmailVerifiedUsersToJoinOrganization(v bool) {
	o.AllowEmailVerifiedUsersToJoinOrganization = &v
}

// GetAllowInvitesFrom returns the AllowInvitesFrom field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphAuthorizationPolicy) GetAllowInvitesFrom() AnyOfmicrosoftGraphAllowInvitesFrom {
	if o == nil  {
		var ret AnyOfmicrosoftGraphAllowInvitesFrom
		return ret
	}
	return o.AllowInvitesFrom
}

// GetAllowInvitesFromOk returns a tuple with the AllowInvitesFrom field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphAuthorizationPolicy) GetAllowInvitesFromOk() (*AnyOfmicrosoftGraphAllowInvitesFrom, bool) {
	if o == nil || o.AllowInvitesFrom == nil {
		return nil, false
	}
	return &o.AllowInvitesFrom, true
}

// HasAllowInvitesFrom returns a boolean if a field has been set.
func (o *MicrosoftGraphAuthorizationPolicy) HasAllowInvitesFrom() bool {
	if o != nil && o.AllowInvitesFrom != nil {
		return true
	}

	return false
}

// SetAllowInvitesFrom gets a reference to the given AnyOfmicrosoftGraphAllowInvitesFrom and assigns it to the AllowInvitesFrom field.
func (o *MicrosoftGraphAuthorizationPolicy) SetAllowInvitesFrom(v AnyOfmicrosoftGraphAllowInvitesFrom) {
	o.AllowInvitesFrom = v
}

// GetBlockMsolPowerShell returns the BlockMsolPowerShell field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphAuthorizationPolicy) GetBlockMsolPowerShell() bool {
	if o == nil || o.BlockMsolPowerShell.Get() == nil {
		var ret bool
		return ret
	}
	return *o.BlockMsolPowerShell.Get()
}

// GetBlockMsolPowerShellOk returns a tuple with the BlockMsolPowerShell field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphAuthorizationPolicy) GetBlockMsolPowerShellOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.BlockMsolPowerShell.Get(), o.BlockMsolPowerShell.IsSet()
}

// HasBlockMsolPowerShell returns a boolean if a field has been set.
func (o *MicrosoftGraphAuthorizationPolicy) HasBlockMsolPowerShell() bool {
	if o != nil && o.BlockMsolPowerShell.IsSet() {
		return true
	}

	return false
}

// SetBlockMsolPowerShell gets a reference to the given NullableBool and assigns it to the BlockMsolPowerShell field.
func (o *MicrosoftGraphAuthorizationPolicy) SetBlockMsolPowerShell(v bool) {
	o.BlockMsolPowerShell.Set(&v)
}
// SetBlockMsolPowerShellNil sets the value for BlockMsolPowerShell to be an explicit nil
func (o *MicrosoftGraphAuthorizationPolicy) SetBlockMsolPowerShellNil() {
	o.BlockMsolPowerShell.Set(nil)
}

// UnsetBlockMsolPowerShell ensures that no value is present for BlockMsolPowerShell, not even an explicit nil
func (o *MicrosoftGraphAuthorizationPolicy) UnsetBlockMsolPowerShell() {
	o.BlockMsolPowerShell.Unset()
}

// GetDefaultUserRolePermissions returns the DefaultUserRolePermissions field value if set, zero value otherwise.
func (o *MicrosoftGraphAuthorizationPolicy) GetDefaultUserRolePermissions() MicrosoftGraphDefaultUserRolePermissions {
	if o == nil || o.DefaultUserRolePermissions == nil {
		var ret MicrosoftGraphDefaultUserRolePermissions
		return ret
	}
	return *o.DefaultUserRolePermissions
}

// GetDefaultUserRolePermissionsOk returns a tuple with the DefaultUserRolePermissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphAuthorizationPolicy) GetDefaultUserRolePermissionsOk() (*MicrosoftGraphDefaultUserRolePermissions, bool) {
	if o == nil || o.DefaultUserRolePermissions == nil {
		return nil, false
	}
	return o.DefaultUserRolePermissions, true
}

// HasDefaultUserRolePermissions returns a boolean if a field has been set.
func (o *MicrosoftGraphAuthorizationPolicy) HasDefaultUserRolePermissions() bool {
	if o != nil && o.DefaultUserRolePermissions != nil {
		return true
	}

	return false
}

// SetDefaultUserRolePermissions gets a reference to the given MicrosoftGraphDefaultUserRolePermissions and assigns it to the DefaultUserRolePermissions field.
func (o *MicrosoftGraphAuthorizationPolicy) SetDefaultUserRolePermissions(v MicrosoftGraphDefaultUserRolePermissions) {
	o.DefaultUserRolePermissions = &v
}

// GetGuestUserRoleId returns the GuestUserRoleId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphAuthorizationPolicy) GetGuestUserRoleId() string {
	if o == nil || o.GuestUserRoleId.Get() == nil {
		var ret string
		return ret
	}
	return *o.GuestUserRoleId.Get()
}

// GetGuestUserRoleIdOk returns a tuple with the GuestUserRoleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphAuthorizationPolicy) GetGuestUserRoleIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.GuestUserRoleId.Get(), o.GuestUserRoleId.IsSet()
}

// HasGuestUserRoleId returns a boolean if a field has been set.
func (o *MicrosoftGraphAuthorizationPolicy) HasGuestUserRoleId() bool {
	if o != nil && o.GuestUserRoleId.IsSet() {
		return true
	}

	return false
}

// SetGuestUserRoleId gets a reference to the given NullableString and assigns it to the GuestUserRoleId field.
func (o *MicrosoftGraphAuthorizationPolicy) SetGuestUserRoleId(v string) {
	o.GuestUserRoleId.Set(&v)
}
// SetGuestUserRoleIdNil sets the value for GuestUserRoleId to be an explicit nil
func (o *MicrosoftGraphAuthorizationPolicy) SetGuestUserRoleIdNil() {
	o.GuestUserRoleId.Set(nil)
}

// UnsetGuestUserRoleId ensures that no value is present for GuestUserRoleId, not even an explicit nil
func (o *MicrosoftGraphAuthorizationPolicy) UnsetGuestUserRoleId() {
	o.GuestUserRoleId.Unset()
}

func (o MicrosoftGraphAuthorizationPolicy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.DeletedDateTime.IsSet() {
		toSerialize["deletedDateTime"] = o.DeletedDateTime.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.DisplayName.IsSet() {
		toSerialize["displayName"] = o.DisplayName.Get()
	}
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

type NullableMicrosoftGraphAuthorizationPolicy struct {
	value *MicrosoftGraphAuthorizationPolicy
	isSet bool
}

func (v NullableMicrosoftGraphAuthorizationPolicy) Get() *MicrosoftGraphAuthorizationPolicy {
	return v.value
}

func (v *NullableMicrosoftGraphAuthorizationPolicy) Set(val *MicrosoftGraphAuthorizationPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphAuthorizationPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphAuthorizationPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphAuthorizationPolicy(val *MicrosoftGraphAuthorizationPolicy) *NullableMicrosoftGraphAuthorizationPolicy {
	return &NullableMicrosoftGraphAuthorizationPolicy{value: val, isSet: true}
}

func (v NullableMicrosoftGraphAuthorizationPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphAuthorizationPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


