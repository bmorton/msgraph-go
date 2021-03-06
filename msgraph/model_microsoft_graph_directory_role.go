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

// MicrosoftGraphDirectoryRole struct for MicrosoftGraphDirectoryRole
type MicrosoftGraphDirectoryRole struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
	DeletedDateTime NullableTime `json:"deletedDateTime,omitempty"`
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

// NewMicrosoftGraphDirectoryRole instantiates a new MicrosoftGraphDirectoryRole object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphDirectoryRole() *MicrosoftGraphDirectoryRole {
	this := MicrosoftGraphDirectoryRole{}
	return &this
}

// NewMicrosoftGraphDirectoryRoleWithDefaults instantiates a new MicrosoftGraphDirectoryRole object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphDirectoryRoleWithDefaults() *MicrosoftGraphDirectoryRole {
	this := MicrosoftGraphDirectoryRole{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MicrosoftGraphDirectoryRole) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphDirectoryRole) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphDirectoryRole) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MicrosoftGraphDirectoryRole) SetId(v string) {
	o.Id = &v
}

// GetDeletedDateTime returns the DeletedDateTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphDirectoryRole) GetDeletedDateTime() time.Time {
	if o == nil || o.DeletedDateTime.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.DeletedDateTime.Get()
}

// GetDeletedDateTimeOk returns a tuple with the DeletedDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphDirectoryRole) GetDeletedDateTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DeletedDateTime.Get(), o.DeletedDateTime.IsSet()
}

// HasDeletedDateTime returns a boolean if a field has been set.
func (o *MicrosoftGraphDirectoryRole) HasDeletedDateTime() bool {
	if o != nil && o.DeletedDateTime.IsSet() {
		return true
	}

	return false
}

// SetDeletedDateTime gets a reference to the given NullableTime and assigns it to the DeletedDateTime field.
func (o *MicrosoftGraphDirectoryRole) SetDeletedDateTime(v time.Time) {
	o.DeletedDateTime.Set(&v)
}
// SetDeletedDateTimeNil sets the value for DeletedDateTime to be an explicit nil
func (o *MicrosoftGraphDirectoryRole) SetDeletedDateTimeNil() {
	o.DeletedDateTime.Set(nil)
}

// UnsetDeletedDateTime ensures that no value is present for DeletedDateTime, not even an explicit nil
func (o *MicrosoftGraphDirectoryRole) UnsetDeletedDateTime() {
	o.DeletedDateTime.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphDirectoryRole) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphDirectoryRole) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *MicrosoftGraphDirectoryRole) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *MicrosoftGraphDirectoryRole) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *MicrosoftGraphDirectoryRole) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *MicrosoftGraphDirectoryRole) UnsetDescription() {
	o.Description.Unset()
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphDirectoryRole) GetDisplayName() string {
	if o == nil || o.DisplayName.Get() == nil {
		var ret string
		return ret
	}
	return *o.DisplayName.Get()
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphDirectoryRole) GetDisplayNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DisplayName.Get(), o.DisplayName.IsSet()
}

// HasDisplayName returns a boolean if a field has been set.
func (o *MicrosoftGraphDirectoryRole) HasDisplayName() bool {
	if o != nil && o.DisplayName.IsSet() {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given NullableString and assigns it to the DisplayName field.
func (o *MicrosoftGraphDirectoryRole) SetDisplayName(v string) {
	o.DisplayName.Set(&v)
}
// SetDisplayNameNil sets the value for DisplayName to be an explicit nil
func (o *MicrosoftGraphDirectoryRole) SetDisplayNameNil() {
	o.DisplayName.Set(nil)
}

// UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
func (o *MicrosoftGraphDirectoryRole) UnsetDisplayName() {
	o.DisplayName.Unset()
}

// GetRoleTemplateId returns the RoleTemplateId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphDirectoryRole) GetRoleTemplateId() string {
	if o == nil || o.RoleTemplateId.Get() == nil {
		var ret string
		return ret
	}
	return *o.RoleTemplateId.Get()
}

// GetRoleTemplateIdOk returns a tuple with the RoleTemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphDirectoryRole) GetRoleTemplateIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RoleTemplateId.Get(), o.RoleTemplateId.IsSet()
}

// HasRoleTemplateId returns a boolean if a field has been set.
func (o *MicrosoftGraphDirectoryRole) HasRoleTemplateId() bool {
	if o != nil && o.RoleTemplateId.IsSet() {
		return true
	}

	return false
}

// SetRoleTemplateId gets a reference to the given NullableString and assigns it to the RoleTemplateId field.
func (o *MicrosoftGraphDirectoryRole) SetRoleTemplateId(v string) {
	o.RoleTemplateId.Set(&v)
}
// SetRoleTemplateIdNil sets the value for RoleTemplateId to be an explicit nil
func (o *MicrosoftGraphDirectoryRole) SetRoleTemplateIdNil() {
	o.RoleTemplateId.Set(nil)
}

// UnsetRoleTemplateId ensures that no value is present for RoleTemplateId, not even an explicit nil
func (o *MicrosoftGraphDirectoryRole) UnsetRoleTemplateId() {
	o.RoleTemplateId.Unset()
}

// GetMembers returns the Members field value if set, zero value otherwise.
func (o *MicrosoftGraphDirectoryRole) GetMembers() []MicrosoftGraphDirectoryObject {
	if o == nil || o.Members == nil {
		var ret []MicrosoftGraphDirectoryObject
		return ret
	}
	return *o.Members
}

// GetMembersOk returns a tuple with the Members field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphDirectoryRole) GetMembersOk() (*[]MicrosoftGraphDirectoryObject, bool) {
	if o == nil || o.Members == nil {
		return nil, false
	}
	return o.Members, true
}

// HasMembers returns a boolean if a field has been set.
func (o *MicrosoftGraphDirectoryRole) HasMembers() bool {
	if o != nil && o.Members != nil {
		return true
	}

	return false
}

// SetMembers gets a reference to the given []MicrosoftGraphDirectoryObject and assigns it to the Members field.
func (o *MicrosoftGraphDirectoryRole) SetMembers(v []MicrosoftGraphDirectoryObject) {
	o.Members = &v
}

// GetScopedMembers returns the ScopedMembers field value if set, zero value otherwise.
func (o *MicrosoftGraphDirectoryRole) GetScopedMembers() []MicrosoftGraphScopedRoleMembership {
	if o == nil || o.ScopedMembers == nil {
		var ret []MicrosoftGraphScopedRoleMembership
		return ret
	}
	return *o.ScopedMembers
}

// GetScopedMembersOk returns a tuple with the ScopedMembers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphDirectoryRole) GetScopedMembersOk() (*[]MicrosoftGraphScopedRoleMembership, bool) {
	if o == nil || o.ScopedMembers == nil {
		return nil, false
	}
	return o.ScopedMembers, true
}

// HasScopedMembers returns a boolean if a field has been set.
func (o *MicrosoftGraphDirectoryRole) HasScopedMembers() bool {
	if o != nil && o.ScopedMembers != nil {
		return true
	}

	return false
}

// SetScopedMembers gets a reference to the given []MicrosoftGraphScopedRoleMembership and assigns it to the ScopedMembers field.
func (o *MicrosoftGraphDirectoryRole) SetScopedMembers(v []MicrosoftGraphScopedRoleMembership) {
	o.ScopedMembers = &v
}

func (o MicrosoftGraphDirectoryRole) MarshalJSON() ([]byte, error) {
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

type NullableMicrosoftGraphDirectoryRole struct {
	value *MicrosoftGraphDirectoryRole
	isSet bool
}

func (v NullableMicrosoftGraphDirectoryRole) Get() *MicrosoftGraphDirectoryRole {
	return v.value
}

func (v *NullableMicrosoftGraphDirectoryRole) Set(val *MicrosoftGraphDirectoryRole) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphDirectoryRole) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphDirectoryRole) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphDirectoryRole(val *MicrosoftGraphDirectoryRole) *NullableMicrosoftGraphDirectoryRole {
	return &NullableMicrosoftGraphDirectoryRole{value: val, isSet: true}
}

func (v NullableMicrosoftGraphDirectoryRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphDirectoryRole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


