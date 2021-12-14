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

// MicrosoftGraphUnifiedRoleAssignment struct for MicrosoftGraphUnifiedRoleAssignment
type MicrosoftGraphUnifiedRoleAssignment struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
	// Identifier of the app-specific scope when the assignment scope is app-specific.  Either this property or directoryScopeId is required. App scopes are scopes that are defined and understood by this application only. Use / for tenant-wide app scopes. Use directoryScopeId to limit the scope to particular directory objects, for example, administrative units. Supports $filter (eq, in).
	AppScopeId NullableString `json:"appScopeId,omitempty"`
	Condition NullableString `json:"condition,omitempty"`
	// Identifier of the directory object representing the scope of the assignment.  Either this property or appScopeId is required. The scope of an assignment determines the set of resources for which the principal has been granted access. Directory scopes are shared scopes stored in the directory that are understood by multiple applications. Use / for tenant-wide scope. Use appScopeId to limit the scope to an application only. Supports $filter (eq, in).
	DirectoryScopeId NullableString `json:"directoryScopeId,omitempty"`
	// Identifier of the principal to which the assignment is granted. Supports $filter (eq, in).
	PrincipalId NullableString `json:"principalId,omitempty"`
	// Identifier of the role definition the assignment is for. Read only. Supports $filter (eq, in).
	RoleDefinitionId NullableString `json:"roleDefinitionId,omitempty"`
	// Read-only property with details of the app specific scope when the assignment scope is app specific. Containment entity. Supports $expand.
	AppScope AnyOfmicrosoftGraphAppScope `json:"appScope,omitempty"`
	// The directory object that is the scope of the assignment. Read-only. Supports $expand.
	DirectoryScope AnyOfmicrosoftGraphDirectoryObject `json:"directoryScope,omitempty"`
	// Referencing the assigned principal. Read-only. Supports $expand.
	Principal AnyOfmicrosoftGraphDirectoryObject `json:"principal,omitempty"`
	// The roleDefinition the assignment is for.  Supports $expand. roleDefinition.Id will be auto expanded.
	RoleDefinition AnyOfmicrosoftGraphUnifiedRoleDefinition `json:"roleDefinition,omitempty"`
}

// NewMicrosoftGraphUnifiedRoleAssignment instantiates a new MicrosoftGraphUnifiedRoleAssignment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphUnifiedRoleAssignment() *MicrosoftGraphUnifiedRoleAssignment {
	this := MicrosoftGraphUnifiedRoleAssignment{}
	return &this
}

// NewMicrosoftGraphUnifiedRoleAssignmentWithDefaults instantiates a new MicrosoftGraphUnifiedRoleAssignment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphUnifiedRoleAssignmentWithDefaults() *MicrosoftGraphUnifiedRoleAssignment {
	this := MicrosoftGraphUnifiedRoleAssignment{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MicrosoftGraphUnifiedRoleAssignment) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphUnifiedRoleAssignment) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphUnifiedRoleAssignment) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MicrosoftGraphUnifiedRoleAssignment) SetId(v string) {
	o.Id = &v
}

// GetAppScopeId returns the AppScopeId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphUnifiedRoleAssignment) GetAppScopeId() string {
	if o == nil || o.AppScopeId.Get() == nil {
		var ret string
		return ret
	}
	return *o.AppScopeId.Get()
}

// GetAppScopeIdOk returns a tuple with the AppScopeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphUnifiedRoleAssignment) GetAppScopeIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AppScopeId.Get(), o.AppScopeId.IsSet()
}

// HasAppScopeId returns a boolean if a field has been set.
func (o *MicrosoftGraphUnifiedRoleAssignment) HasAppScopeId() bool {
	if o != nil && o.AppScopeId.IsSet() {
		return true
	}

	return false
}

// SetAppScopeId gets a reference to the given NullableString and assigns it to the AppScopeId field.
func (o *MicrosoftGraphUnifiedRoleAssignment) SetAppScopeId(v string) {
	o.AppScopeId.Set(&v)
}
// SetAppScopeIdNil sets the value for AppScopeId to be an explicit nil
func (o *MicrosoftGraphUnifiedRoleAssignment) SetAppScopeIdNil() {
	o.AppScopeId.Set(nil)
}

// UnsetAppScopeId ensures that no value is present for AppScopeId, not even an explicit nil
func (o *MicrosoftGraphUnifiedRoleAssignment) UnsetAppScopeId() {
	o.AppScopeId.Unset()
}

// GetCondition returns the Condition field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphUnifiedRoleAssignment) GetCondition() string {
	if o == nil || o.Condition.Get() == nil {
		var ret string
		return ret
	}
	return *o.Condition.Get()
}

// GetConditionOk returns a tuple with the Condition field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphUnifiedRoleAssignment) GetConditionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Condition.Get(), o.Condition.IsSet()
}

// HasCondition returns a boolean if a field has been set.
func (o *MicrosoftGraphUnifiedRoleAssignment) HasCondition() bool {
	if o != nil && o.Condition.IsSet() {
		return true
	}

	return false
}

// SetCondition gets a reference to the given NullableString and assigns it to the Condition field.
func (o *MicrosoftGraphUnifiedRoleAssignment) SetCondition(v string) {
	o.Condition.Set(&v)
}
// SetConditionNil sets the value for Condition to be an explicit nil
func (o *MicrosoftGraphUnifiedRoleAssignment) SetConditionNil() {
	o.Condition.Set(nil)
}

// UnsetCondition ensures that no value is present for Condition, not even an explicit nil
func (o *MicrosoftGraphUnifiedRoleAssignment) UnsetCondition() {
	o.Condition.Unset()
}

// GetDirectoryScopeId returns the DirectoryScopeId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphUnifiedRoleAssignment) GetDirectoryScopeId() string {
	if o == nil || o.DirectoryScopeId.Get() == nil {
		var ret string
		return ret
	}
	return *o.DirectoryScopeId.Get()
}

// GetDirectoryScopeIdOk returns a tuple with the DirectoryScopeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphUnifiedRoleAssignment) GetDirectoryScopeIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DirectoryScopeId.Get(), o.DirectoryScopeId.IsSet()
}

// HasDirectoryScopeId returns a boolean if a field has been set.
func (o *MicrosoftGraphUnifiedRoleAssignment) HasDirectoryScopeId() bool {
	if o != nil && o.DirectoryScopeId.IsSet() {
		return true
	}

	return false
}

// SetDirectoryScopeId gets a reference to the given NullableString and assigns it to the DirectoryScopeId field.
func (o *MicrosoftGraphUnifiedRoleAssignment) SetDirectoryScopeId(v string) {
	o.DirectoryScopeId.Set(&v)
}
// SetDirectoryScopeIdNil sets the value for DirectoryScopeId to be an explicit nil
func (o *MicrosoftGraphUnifiedRoleAssignment) SetDirectoryScopeIdNil() {
	o.DirectoryScopeId.Set(nil)
}

// UnsetDirectoryScopeId ensures that no value is present for DirectoryScopeId, not even an explicit nil
func (o *MicrosoftGraphUnifiedRoleAssignment) UnsetDirectoryScopeId() {
	o.DirectoryScopeId.Unset()
}

// GetPrincipalId returns the PrincipalId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphUnifiedRoleAssignment) GetPrincipalId() string {
	if o == nil || o.PrincipalId.Get() == nil {
		var ret string
		return ret
	}
	return *o.PrincipalId.Get()
}

// GetPrincipalIdOk returns a tuple with the PrincipalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphUnifiedRoleAssignment) GetPrincipalIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PrincipalId.Get(), o.PrincipalId.IsSet()
}

// HasPrincipalId returns a boolean if a field has been set.
func (o *MicrosoftGraphUnifiedRoleAssignment) HasPrincipalId() bool {
	if o != nil && o.PrincipalId.IsSet() {
		return true
	}

	return false
}

// SetPrincipalId gets a reference to the given NullableString and assigns it to the PrincipalId field.
func (o *MicrosoftGraphUnifiedRoleAssignment) SetPrincipalId(v string) {
	o.PrincipalId.Set(&v)
}
// SetPrincipalIdNil sets the value for PrincipalId to be an explicit nil
func (o *MicrosoftGraphUnifiedRoleAssignment) SetPrincipalIdNil() {
	o.PrincipalId.Set(nil)
}

// UnsetPrincipalId ensures that no value is present for PrincipalId, not even an explicit nil
func (o *MicrosoftGraphUnifiedRoleAssignment) UnsetPrincipalId() {
	o.PrincipalId.Unset()
}

// GetRoleDefinitionId returns the RoleDefinitionId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphUnifiedRoleAssignment) GetRoleDefinitionId() string {
	if o == nil || o.RoleDefinitionId.Get() == nil {
		var ret string
		return ret
	}
	return *o.RoleDefinitionId.Get()
}

// GetRoleDefinitionIdOk returns a tuple with the RoleDefinitionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphUnifiedRoleAssignment) GetRoleDefinitionIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RoleDefinitionId.Get(), o.RoleDefinitionId.IsSet()
}

// HasRoleDefinitionId returns a boolean if a field has been set.
func (o *MicrosoftGraphUnifiedRoleAssignment) HasRoleDefinitionId() bool {
	if o != nil && o.RoleDefinitionId.IsSet() {
		return true
	}

	return false
}

// SetRoleDefinitionId gets a reference to the given NullableString and assigns it to the RoleDefinitionId field.
func (o *MicrosoftGraphUnifiedRoleAssignment) SetRoleDefinitionId(v string) {
	o.RoleDefinitionId.Set(&v)
}
// SetRoleDefinitionIdNil sets the value for RoleDefinitionId to be an explicit nil
func (o *MicrosoftGraphUnifiedRoleAssignment) SetRoleDefinitionIdNil() {
	o.RoleDefinitionId.Set(nil)
}

// UnsetRoleDefinitionId ensures that no value is present for RoleDefinitionId, not even an explicit nil
func (o *MicrosoftGraphUnifiedRoleAssignment) UnsetRoleDefinitionId() {
	o.RoleDefinitionId.Unset()
}

// GetAppScope returns the AppScope field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphUnifiedRoleAssignment) GetAppScope() AnyOfmicrosoftGraphAppScope {
	if o == nil  {
		var ret AnyOfmicrosoftGraphAppScope
		return ret
	}
	return o.AppScope
}

// GetAppScopeOk returns a tuple with the AppScope field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphUnifiedRoleAssignment) GetAppScopeOk() (*AnyOfmicrosoftGraphAppScope, bool) {
	if o == nil || o.AppScope == nil {
		return nil, false
	}
	return &o.AppScope, true
}

// HasAppScope returns a boolean if a field has been set.
func (o *MicrosoftGraphUnifiedRoleAssignment) HasAppScope() bool {
	if o != nil && o.AppScope != nil {
		return true
	}

	return false
}

// SetAppScope gets a reference to the given AnyOfmicrosoftGraphAppScope and assigns it to the AppScope field.
func (o *MicrosoftGraphUnifiedRoleAssignment) SetAppScope(v AnyOfmicrosoftGraphAppScope) {
	o.AppScope = v
}

// GetDirectoryScope returns the DirectoryScope field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphUnifiedRoleAssignment) GetDirectoryScope() AnyOfmicrosoftGraphDirectoryObject {
	if o == nil  {
		var ret AnyOfmicrosoftGraphDirectoryObject
		return ret
	}
	return o.DirectoryScope
}

// GetDirectoryScopeOk returns a tuple with the DirectoryScope field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphUnifiedRoleAssignment) GetDirectoryScopeOk() (*AnyOfmicrosoftGraphDirectoryObject, bool) {
	if o == nil || o.DirectoryScope == nil {
		return nil, false
	}
	return &o.DirectoryScope, true
}

// HasDirectoryScope returns a boolean if a field has been set.
func (o *MicrosoftGraphUnifiedRoleAssignment) HasDirectoryScope() bool {
	if o != nil && o.DirectoryScope != nil {
		return true
	}

	return false
}

// SetDirectoryScope gets a reference to the given AnyOfmicrosoftGraphDirectoryObject and assigns it to the DirectoryScope field.
func (o *MicrosoftGraphUnifiedRoleAssignment) SetDirectoryScope(v AnyOfmicrosoftGraphDirectoryObject) {
	o.DirectoryScope = v
}

// GetPrincipal returns the Principal field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphUnifiedRoleAssignment) GetPrincipal() AnyOfmicrosoftGraphDirectoryObject {
	if o == nil  {
		var ret AnyOfmicrosoftGraphDirectoryObject
		return ret
	}
	return o.Principal
}

// GetPrincipalOk returns a tuple with the Principal field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphUnifiedRoleAssignment) GetPrincipalOk() (*AnyOfmicrosoftGraphDirectoryObject, bool) {
	if o == nil || o.Principal == nil {
		return nil, false
	}
	return &o.Principal, true
}

// HasPrincipal returns a boolean if a field has been set.
func (o *MicrosoftGraphUnifiedRoleAssignment) HasPrincipal() bool {
	if o != nil && o.Principal != nil {
		return true
	}

	return false
}

// SetPrincipal gets a reference to the given AnyOfmicrosoftGraphDirectoryObject and assigns it to the Principal field.
func (o *MicrosoftGraphUnifiedRoleAssignment) SetPrincipal(v AnyOfmicrosoftGraphDirectoryObject) {
	o.Principal = v
}

// GetRoleDefinition returns the RoleDefinition field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphUnifiedRoleAssignment) GetRoleDefinition() AnyOfmicrosoftGraphUnifiedRoleDefinition {
	if o == nil  {
		var ret AnyOfmicrosoftGraphUnifiedRoleDefinition
		return ret
	}
	return o.RoleDefinition
}

// GetRoleDefinitionOk returns a tuple with the RoleDefinition field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphUnifiedRoleAssignment) GetRoleDefinitionOk() (*AnyOfmicrosoftGraphUnifiedRoleDefinition, bool) {
	if o == nil || o.RoleDefinition == nil {
		return nil, false
	}
	return &o.RoleDefinition, true
}

// HasRoleDefinition returns a boolean if a field has been set.
func (o *MicrosoftGraphUnifiedRoleAssignment) HasRoleDefinition() bool {
	if o != nil && o.RoleDefinition != nil {
		return true
	}

	return false
}

// SetRoleDefinition gets a reference to the given AnyOfmicrosoftGraphUnifiedRoleDefinition and assigns it to the RoleDefinition field.
func (o *MicrosoftGraphUnifiedRoleAssignment) SetRoleDefinition(v AnyOfmicrosoftGraphUnifiedRoleDefinition) {
	o.RoleDefinition = v
}

func (o MicrosoftGraphUnifiedRoleAssignment) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.AppScopeId.IsSet() {
		toSerialize["appScopeId"] = o.AppScopeId.Get()
	}
	if o.Condition.IsSet() {
		toSerialize["condition"] = o.Condition.Get()
	}
	if o.DirectoryScopeId.IsSet() {
		toSerialize["directoryScopeId"] = o.DirectoryScopeId.Get()
	}
	if o.PrincipalId.IsSet() {
		toSerialize["principalId"] = o.PrincipalId.Get()
	}
	if o.RoleDefinitionId.IsSet() {
		toSerialize["roleDefinitionId"] = o.RoleDefinitionId.Get()
	}
	if o.AppScope != nil {
		toSerialize["appScope"] = o.AppScope
	}
	if o.DirectoryScope != nil {
		toSerialize["directoryScope"] = o.DirectoryScope
	}
	if o.Principal != nil {
		toSerialize["principal"] = o.Principal
	}
	if o.RoleDefinition != nil {
		toSerialize["roleDefinition"] = o.RoleDefinition
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphUnifiedRoleAssignment struct {
	value *MicrosoftGraphUnifiedRoleAssignment
	isSet bool
}

func (v NullableMicrosoftGraphUnifiedRoleAssignment) Get() *MicrosoftGraphUnifiedRoleAssignment {
	return v.value
}

func (v *NullableMicrosoftGraphUnifiedRoleAssignment) Set(val *MicrosoftGraphUnifiedRoleAssignment) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphUnifiedRoleAssignment) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphUnifiedRoleAssignment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphUnifiedRoleAssignment(val *MicrosoftGraphUnifiedRoleAssignment) *NullableMicrosoftGraphUnifiedRoleAssignment {
	return &NullableMicrosoftGraphUnifiedRoleAssignment{value: val, isSet: true}
}

func (v NullableMicrosoftGraphUnifiedRoleAssignment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphUnifiedRoleAssignment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


