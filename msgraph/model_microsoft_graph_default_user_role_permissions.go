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

// MicrosoftGraphDefaultUserRolePermissions struct for MicrosoftGraphDefaultUserRolePermissions
type MicrosoftGraphDefaultUserRolePermissions struct {
	// Indicates whether the default user role can create applications.
	AllowedToCreateApps *bool `json:"allowedToCreateApps,omitempty"`
	// Indicates whether the default user role can create security groups.
	AllowedToCreateSecurityGroups *bool `json:"allowedToCreateSecurityGroups,omitempty"`
	// Indicates whether the default user role can read other users.
	AllowedToReadOtherUsers *bool `json:"allowedToReadOtherUsers,omitempty"`
	// Indicates if user consent to apps is allowed, and if it is, which permission to grant consent and which app consent policy (permissionGrantPolicy) govern the permission for users to grant consent. Value should be in the format managePermissionGrantsForSelf.{id}, where {id} is the id of a built-in or custom app consent policy. An empty list indicates user consent to apps is disabled.
	PermissionGrantPoliciesAssigned *[]*string `json:"permissionGrantPoliciesAssigned,omitempty"`
}

// NewMicrosoftGraphDefaultUserRolePermissions instantiates a new MicrosoftGraphDefaultUserRolePermissions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphDefaultUserRolePermissions() *MicrosoftGraphDefaultUserRolePermissions {
	this := MicrosoftGraphDefaultUserRolePermissions{}
	return &this
}

// NewMicrosoftGraphDefaultUserRolePermissionsWithDefaults instantiates a new MicrosoftGraphDefaultUserRolePermissions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphDefaultUserRolePermissionsWithDefaults() *MicrosoftGraphDefaultUserRolePermissions {
	this := MicrosoftGraphDefaultUserRolePermissions{}
	return &this
}

// GetAllowedToCreateApps returns the AllowedToCreateApps field value if set, zero value otherwise.
func (o *MicrosoftGraphDefaultUserRolePermissions) GetAllowedToCreateApps() bool {
	if o == nil || o.AllowedToCreateApps == nil {
		var ret bool
		return ret
	}
	return *o.AllowedToCreateApps
}

// GetAllowedToCreateAppsOk returns a tuple with the AllowedToCreateApps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphDefaultUserRolePermissions) GetAllowedToCreateAppsOk() (*bool, bool) {
	if o == nil || o.AllowedToCreateApps == nil {
		return nil, false
	}
	return o.AllowedToCreateApps, true
}

// HasAllowedToCreateApps returns a boolean if a field has been set.
func (o *MicrosoftGraphDefaultUserRolePermissions) HasAllowedToCreateApps() bool {
	if o != nil && o.AllowedToCreateApps != nil {
		return true
	}

	return false
}

// SetAllowedToCreateApps gets a reference to the given bool and assigns it to the AllowedToCreateApps field.
func (o *MicrosoftGraphDefaultUserRolePermissions) SetAllowedToCreateApps(v bool) {
	o.AllowedToCreateApps = &v
}

// GetAllowedToCreateSecurityGroups returns the AllowedToCreateSecurityGroups field value if set, zero value otherwise.
func (o *MicrosoftGraphDefaultUserRolePermissions) GetAllowedToCreateSecurityGroups() bool {
	if o == nil || o.AllowedToCreateSecurityGroups == nil {
		var ret bool
		return ret
	}
	return *o.AllowedToCreateSecurityGroups
}

// GetAllowedToCreateSecurityGroupsOk returns a tuple with the AllowedToCreateSecurityGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphDefaultUserRolePermissions) GetAllowedToCreateSecurityGroupsOk() (*bool, bool) {
	if o == nil || o.AllowedToCreateSecurityGroups == nil {
		return nil, false
	}
	return o.AllowedToCreateSecurityGroups, true
}

// HasAllowedToCreateSecurityGroups returns a boolean if a field has been set.
func (o *MicrosoftGraphDefaultUserRolePermissions) HasAllowedToCreateSecurityGroups() bool {
	if o != nil && o.AllowedToCreateSecurityGroups != nil {
		return true
	}

	return false
}

// SetAllowedToCreateSecurityGroups gets a reference to the given bool and assigns it to the AllowedToCreateSecurityGroups field.
func (o *MicrosoftGraphDefaultUserRolePermissions) SetAllowedToCreateSecurityGroups(v bool) {
	o.AllowedToCreateSecurityGroups = &v
}

// GetAllowedToReadOtherUsers returns the AllowedToReadOtherUsers field value if set, zero value otherwise.
func (o *MicrosoftGraphDefaultUserRolePermissions) GetAllowedToReadOtherUsers() bool {
	if o == nil || o.AllowedToReadOtherUsers == nil {
		var ret bool
		return ret
	}
	return *o.AllowedToReadOtherUsers
}

// GetAllowedToReadOtherUsersOk returns a tuple with the AllowedToReadOtherUsers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphDefaultUserRolePermissions) GetAllowedToReadOtherUsersOk() (*bool, bool) {
	if o == nil || o.AllowedToReadOtherUsers == nil {
		return nil, false
	}
	return o.AllowedToReadOtherUsers, true
}

// HasAllowedToReadOtherUsers returns a boolean if a field has been set.
func (o *MicrosoftGraphDefaultUserRolePermissions) HasAllowedToReadOtherUsers() bool {
	if o != nil && o.AllowedToReadOtherUsers != nil {
		return true
	}

	return false
}

// SetAllowedToReadOtherUsers gets a reference to the given bool and assigns it to the AllowedToReadOtherUsers field.
func (o *MicrosoftGraphDefaultUserRolePermissions) SetAllowedToReadOtherUsers(v bool) {
	o.AllowedToReadOtherUsers = &v
}

// GetPermissionGrantPoliciesAssigned returns the PermissionGrantPoliciesAssigned field value if set, zero value otherwise.
func (o *MicrosoftGraphDefaultUserRolePermissions) GetPermissionGrantPoliciesAssigned() []*string {
	if o == nil || o.PermissionGrantPoliciesAssigned == nil {
		var ret []*string
		return ret
	}
	return *o.PermissionGrantPoliciesAssigned
}

// GetPermissionGrantPoliciesAssignedOk returns a tuple with the PermissionGrantPoliciesAssigned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphDefaultUserRolePermissions) GetPermissionGrantPoliciesAssignedOk() (*[]*string, bool) {
	if o == nil || o.PermissionGrantPoliciesAssigned == nil {
		return nil, false
	}
	return o.PermissionGrantPoliciesAssigned, true
}

// HasPermissionGrantPoliciesAssigned returns a boolean if a field has been set.
func (o *MicrosoftGraphDefaultUserRolePermissions) HasPermissionGrantPoliciesAssigned() bool {
	if o != nil && o.PermissionGrantPoliciesAssigned != nil {
		return true
	}

	return false
}

// SetPermissionGrantPoliciesAssigned gets a reference to the given []*string and assigns it to the PermissionGrantPoliciesAssigned field.
func (o *MicrosoftGraphDefaultUserRolePermissions) SetPermissionGrantPoliciesAssigned(v []*string) {
	o.PermissionGrantPoliciesAssigned = &v
}

func (o MicrosoftGraphDefaultUserRolePermissions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AllowedToCreateApps != nil {
		toSerialize["allowedToCreateApps"] = o.AllowedToCreateApps
	}
	if o.AllowedToCreateSecurityGroups != nil {
		toSerialize["allowedToCreateSecurityGroups"] = o.AllowedToCreateSecurityGroups
	}
	if o.AllowedToReadOtherUsers != nil {
		toSerialize["allowedToReadOtherUsers"] = o.AllowedToReadOtherUsers
	}
	if o.PermissionGrantPoliciesAssigned != nil {
		toSerialize["permissionGrantPoliciesAssigned"] = o.PermissionGrantPoliciesAssigned
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphDefaultUserRolePermissions struct {
	value *MicrosoftGraphDefaultUserRolePermissions
	isSet bool
}

func (v NullableMicrosoftGraphDefaultUserRolePermissions) Get() *MicrosoftGraphDefaultUserRolePermissions {
	return v.value
}

func (v *NullableMicrosoftGraphDefaultUserRolePermissions) Set(val *MicrosoftGraphDefaultUserRolePermissions) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphDefaultUserRolePermissions) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphDefaultUserRolePermissions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphDefaultUserRolePermissions(val *MicrosoftGraphDefaultUserRolePermissions) *NullableMicrosoftGraphDefaultUserRolePermissions {
	return &NullableMicrosoftGraphDefaultUserRolePermissions{value: val, isSet: true}
}

func (v NullableMicrosoftGraphDefaultUserRolePermissions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphDefaultUserRolePermissions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


