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

// DirectoryRole struct for DirectoryRole
type DirectoryRole struct {
	// The description for the directory role. Read-only.
	Description NullableString `json:"description,omitempty"`
	// The display name for the directory role. Read-only.
	DisplayName NullableString `json:"displayName,omitempty"`
	// The id of the directoryRoleTemplate that this role is based on. The property must be specified when activating a directory role in a tenant with a POST operation. After the directory role has been activated, the property is read only.
	RoleTemplateId NullableString `json:"roleTemplateId,omitempty"`
	// Users that are members of this directory role. HTTP Methods: GET, POST, DELETE. Read-only. Nullable.
	Members *[]MicrosoftGraphDirectoryObject `json:"members,omitempty"`
	// Members of this directory role that are scoped to administrative units. Read-only. Nullable.
	ScopedMembers *[]MicrosoftGraphScopedRoleMembership `json:"scopedMembers,omitempty"`
}

// NewDirectoryRole instantiates a new DirectoryRole object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDirectoryRole() *DirectoryRole {
	this := DirectoryRole{}
	return &this
}

// NewDirectoryRoleWithDefaults instantiates a new DirectoryRole object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDirectoryRoleWithDefaults() *DirectoryRole {
	this := DirectoryRole{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DirectoryRole) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DirectoryRole) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *DirectoryRole) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *DirectoryRole) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *DirectoryRole) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *DirectoryRole) UnsetDescription() {
	o.Description.Unset()
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DirectoryRole) GetDisplayName() string {
	if o == nil || o.DisplayName.Get() == nil {
		var ret string
		return ret
	}
	return *o.DisplayName.Get()
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DirectoryRole) GetDisplayNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DisplayName.Get(), o.DisplayName.IsSet()
}

// HasDisplayName returns a boolean if a field has been set.
func (o *DirectoryRole) HasDisplayName() bool {
	if o != nil && o.DisplayName.IsSet() {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given NullableString and assigns it to the DisplayName field.
func (o *DirectoryRole) SetDisplayName(v string) {
	o.DisplayName.Set(&v)
}
// SetDisplayNameNil sets the value for DisplayName to be an explicit nil
func (o *DirectoryRole) SetDisplayNameNil() {
	o.DisplayName.Set(nil)
}

// UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
func (o *DirectoryRole) UnsetDisplayName() {
	o.DisplayName.Unset()
}

// GetRoleTemplateId returns the RoleTemplateId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DirectoryRole) GetRoleTemplateId() string {
	if o == nil || o.RoleTemplateId.Get() == nil {
		var ret string
		return ret
	}
	return *o.RoleTemplateId.Get()
}

// GetRoleTemplateIdOk returns a tuple with the RoleTemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DirectoryRole) GetRoleTemplateIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RoleTemplateId.Get(), o.RoleTemplateId.IsSet()
}

// HasRoleTemplateId returns a boolean if a field has been set.
func (o *DirectoryRole) HasRoleTemplateId() bool {
	if o != nil && o.RoleTemplateId.IsSet() {
		return true
	}

	return false
}

// SetRoleTemplateId gets a reference to the given NullableString and assigns it to the RoleTemplateId field.
func (o *DirectoryRole) SetRoleTemplateId(v string) {
	o.RoleTemplateId.Set(&v)
}
// SetRoleTemplateIdNil sets the value for RoleTemplateId to be an explicit nil
func (o *DirectoryRole) SetRoleTemplateIdNil() {
	o.RoleTemplateId.Set(nil)
}

// UnsetRoleTemplateId ensures that no value is present for RoleTemplateId, not even an explicit nil
func (o *DirectoryRole) UnsetRoleTemplateId() {
	o.RoleTemplateId.Unset()
}

// GetMembers returns the Members field value if set, zero value otherwise.
func (o *DirectoryRole) GetMembers() []MicrosoftGraphDirectoryObject {
	if o == nil || o.Members == nil {
		var ret []MicrosoftGraphDirectoryObject
		return ret
	}
	return *o.Members
}

// GetMembersOk returns a tuple with the Members field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DirectoryRole) GetMembersOk() (*[]MicrosoftGraphDirectoryObject, bool) {
	if o == nil || o.Members == nil {
		return nil, false
	}
	return o.Members, true
}

// HasMembers returns a boolean if a field has been set.
func (o *DirectoryRole) HasMembers() bool {
	if o != nil && o.Members != nil {
		return true
	}

	return false
}

// SetMembers gets a reference to the given []MicrosoftGraphDirectoryObject and assigns it to the Members field.
func (o *DirectoryRole) SetMembers(v []MicrosoftGraphDirectoryObject) {
	o.Members = &v
}

// GetScopedMembers returns the ScopedMembers field value if set, zero value otherwise.
func (o *DirectoryRole) GetScopedMembers() []MicrosoftGraphScopedRoleMembership {
	if o == nil || o.ScopedMembers == nil {
		var ret []MicrosoftGraphScopedRoleMembership
		return ret
	}
	return *o.ScopedMembers
}

// GetScopedMembersOk returns a tuple with the ScopedMembers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DirectoryRole) GetScopedMembersOk() (*[]MicrosoftGraphScopedRoleMembership, bool) {
	if o == nil || o.ScopedMembers == nil {
		return nil, false
	}
	return o.ScopedMembers, true
}

// HasScopedMembers returns a boolean if a field has been set.
func (o *DirectoryRole) HasScopedMembers() bool {
	if o != nil && o.ScopedMembers != nil {
		return true
	}

	return false
}

// SetScopedMembers gets a reference to the given []MicrosoftGraphScopedRoleMembership and assigns it to the ScopedMembers field.
func (o *DirectoryRole) SetScopedMembers(v []MicrosoftGraphScopedRoleMembership) {
	o.ScopedMembers = &v
}

func (o DirectoryRole) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.DisplayName.IsSet() {
		toSerialize["displayName"] = o.DisplayName.Get()
	}
	if o.RoleTemplateId.IsSet() {
		toSerialize["roleTemplateId"] = o.RoleTemplateId.Get()
	}
	if o.Members != nil {
		toSerialize["members"] = o.Members
	}
	if o.ScopedMembers != nil {
		toSerialize["scopedMembers"] = o.ScopedMembers
	}
	return json.Marshal(toSerialize)
}

type NullableDirectoryRole struct {
	value *DirectoryRole
	isSet bool
}

func (v NullableDirectoryRole) Get() *DirectoryRole {
	return v.value
}

func (v *NullableDirectoryRole) Set(val *DirectoryRole) {
	v.value = val
	v.isSet = true
}

func (v NullableDirectoryRole) IsSet() bool {
	return v.isSet
}

func (v *NullableDirectoryRole) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDirectoryRole(val *DirectoryRole) *NullableDirectoryRole {
	return &NullableDirectoryRole{value: val, isSet: true}
}

func (v NullableDirectoryRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDirectoryRole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

