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

// UnifiedRoleAssignment struct for UnifiedRoleAssignment
type UnifiedRoleAssignment struct {
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

// NewUnifiedRoleAssignment instantiates a new UnifiedRoleAssignment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnifiedRoleAssignment() *UnifiedRoleAssignment {
	this := UnifiedRoleAssignment{}
	return &this
}

// NewUnifiedRoleAssignmentWithDefaults instantiates a new UnifiedRoleAssignment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnifiedRoleAssignmentWithDefaults() *UnifiedRoleAssignment {
	this := UnifiedRoleAssignment{}
	return &this
}

// GetAppScopeId returns the AppScopeId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UnifiedRoleAssignment) GetAppScopeId() string {
	if o == nil || o.AppScopeId.Get() == nil {
		var ret string
		return ret
	}
	return *o.AppScopeId.Get()
}

// GetAppScopeIdOk returns a tuple with the AppScopeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UnifiedRoleAssignment) GetAppScopeIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AppScopeId.Get(), o.AppScopeId.IsSet()
}

// HasAppScopeId returns a boolean if a field has been set.
func (o *UnifiedRoleAssignment) HasAppScopeId() bool {
	if o != nil && o.AppScopeId.IsSet() {
		return true
	}

	return false
}

// SetAppScopeId gets a reference to the given NullableString and assigns it to the AppScopeId field.
func (o *UnifiedRoleAssignment) SetAppScopeId(v string) {
	o.AppScopeId.Set(&v)
}
// SetAppScopeIdNil sets the value for AppScopeId to be an explicit nil
func (o *UnifiedRoleAssignment) SetAppScopeIdNil() {
	o.AppScopeId.Set(nil)
}

// UnsetAppScopeId ensures that no value is present for AppScopeId, not even an explicit nil
func (o *UnifiedRoleAssignment) UnsetAppScopeId() {
	o.AppScopeId.Unset()
}

// GetCondition returns the Condition field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UnifiedRoleAssignment) GetCondition() string {
	if o == nil || o.Condition.Get() == nil {
		var ret string
		return ret
	}
	return *o.Condition.Get()
}

// GetConditionOk returns a tuple with the Condition field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UnifiedRoleAssignment) GetConditionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Condition.Get(), o.Condition.IsSet()
}

// HasCondition returns a boolean if a field has been set.
func (o *UnifiedRoleAssignment) HasCondition() bool {
	if o != nil && o.Condition.IsSet() {
		return true
	}

	return false
}

// SetCondition gets a reference to the given NullableString and assigns it to the Condition field.
func (o *UnifiedRoleAssignment) SetCondition(v string) {
	o.Condition.Set(&v)
}
// SetConditionNil sets the value for Condition to be an explicit nil
func (o *UnifiedRoleAssignment) SetConditionNil() {
	o.Condition.Set(nil)
}

// UnsetCondition ensures that no value is present for Condition, not even an explicit nil
func (o *UnifiedRoleAssignment) UnsetCondition() {
	o.Condition.Unset()
}

// GetDirectoryScopeId returns the DirectoryScopeId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UnifiedRoleAssignment) GetDirectoryScopeId() string {
	if o == nil || o.DirectoryScopeId.Get() == nil {
		var ret string
		return ret
	}
	return *o.DirectoryScopeId.Get()
}

// GetDirectoryScopeIdOk returns a tuple with the DirectoryScopeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UnifiedRoleAssignment) GetDirectoryScopeIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DirectoryScopeId.Get(), o.DirectoryScopeId.IsSet()
}

// HasDirectoryScopeId returns a boolean if a field has been set.
func (o *UnifiedRoleAssignment) HasDirectoryScopeId() bool {
	if o != nil && o.DirectoryScopeId.IsSet() {
		return true
	}

	return false
}

// SetDirectoryScopeId gets a reference to the given NullableString and assigns it to the DirectoryScopeId field.
func (o *UnifiedRoleAssignment) SetDirectoryScopeId(v string) {
	o.DirectoryScopeId.Set(&v)
}
// SetDirectoryScopeIdNil sets the value for DirectoryScopeId to be an explicit nil
func (o *UnifiedRoleAssignment) SetDirectoryScopeIdNil() {
	o.DirectoryScopeId.Set(nil)
}

// UnsetDirectoryScopeId ensures that no value is present for DirectoryScopeId, not even an explicit nil
func (o *UnifiedRoleAssignment) UnsetDirectoryScopeId() {
	o.DirectoryScopeId.Unset()
}

// GetPrincipalId returns the PrincipalId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UnifiedRoleAssignment) GetPrincipalId() string {
	if o == nil || o.PrincipalId.Get() == nil {
		var ret string
		return ret
	}
	return *o.PrincipalId.Get()
}

// GetPrincipalIdOk returns a tuple with the PrincipalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UnifiedRoleAssignment) GetPrincipalIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PrincipalId.Get(), o.PrincipalId.IsSet()
}

// HasPrincipalId returns a boolean if a field has been set.
func (o *UnifiedRoleAssignment) HasPrincipalId() bool {
	if o != nil && o.PrincipalId.IsSet() {
		return true
	}

	return false
}

// SetPrincipalId gets a reference to the given NullableString and assigns it to the PrincipalId field.
func (o *UnifiedRoleAssignment) SetPrincipalId(v string) {
	o.PrincipalId.Set(&v)
}
// SetPrincipalIdNil sets the value for PrincipalId to be an explicit nil
func (o *UnifiedRoleAssignment) SetPrincipalIdNil() {
	o.PrincipalId.Set(nil)
}

// UnsetPrincipalId ensures that no value is present for PrincipalId, not even an explicit nil
func (o *UnifiedRoleAssignment) UnsetPrincipalId() {
	o.PrincipalId.Unset()
}

// GetRoleDefinitionId returns the RoleDefinitionId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UnifiedRoleAssignment) GetRoleDefinitionId() string {
	if o == nil || o.RoleDefinitionId.Get() == nil {
		var ret string
		return ret
	}
	return *o.RoleDefinitionId.Get()
}

// GetRoleDefinitionIdOk returns a tuple with the RoleDefinitionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UnifiedRoleAssignment) GetRoleDefinitionIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RoleDefinitionId.Get(), o.RoleDefinitionId.IsSet()
}

// HasRoleDefinitionId returns a boolean if a field has been set.
func (o *UnifiedRoleAssignment) HasRoleDefinitionId() bool {
	if o != nil && o.RoleDefinitionId.IsSet() {
		return true
	}

	return false
}

// SetRoleDefinitionId gets a reference to the given NullableString and assigns it to the RoleDefinitionId field.
func (o *UnifiedRoleAssignment) SetRoleDefinitionId(v string) {
	o.RoleDefinitionId.Set(&v)
}
// SetRoleDefinitionIdNil sets the value for RoleDefinitionId to be an explicit nil
func (o *UnifiedRoleAssignment) SetRoleDefinitionIdNil() {
	o.RoleDefinitionId.Set(nil)
}

// UnsetRoleDefinitionId ensures that no value is present for RoleDefinitionId, not even an explicit nil
func (o *UnifiedRoleAssignment) UnsetRoleDefinitionId() {
	o.RoleDefinitionId.Unset()
}

// GetAppScope returns the AppScope field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UnifiedRoleAssignment) GetAppScope() AnyOfmicrosoftGraphAppScope {
	if o == nil  {
		var ret AnyOfmicrosoftGraphAppScope
		return ret
	}
	return o.AppScope
}

// GetAppScopeOk returns a tuple with the AppScope field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UnifiedRoleAssignment) GetAppScopeOk() (*AnyOfmicrosoftGraphAppScope, bool) {
	if o == nil || o.AppScope == nil {
		return nil, false
	}
	return &o.AppScope, true
}

// HasAppScope returns a boolean if a field has been set.
func (o *UnifiedRoleAssignment) HasAppScope() bool {
	if o != nil && o.AppScope != nil {
		return true
	}

	return false
}

// SetAppScope gets a reference to the given AnyOfmicrosoftGraphAppScope and assigns it to the AppScope field.
func (o *UnifiedRoleAssignment) SetAppScope(v AnyOfmicrosoftGraphAppScope) {
	o.AppScope = v
}

// GetDirectoryScope returns the DirectoryScope field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UnifiedRoleAssignment) GetDirectoryScope() AnyOfmicrosoftGraphDirectoryObject {
	if o == nil  {
		var ret AnyOfmicrosoftGraphDirectoryObject
		return ret
	}
	return o.DirectoryScope
}

// GetDirectoryScopeOk returns a tuple with the DirectoryScope field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UnifiedRoleAssignment) GetDirectoryScopeOk() (*AnyOfmicrosoftGraphDirectoryObject, bool) {
	if o == nil || o.DirectoryScope == nil {
		return nil, false
	}
	return &o.DirectoryScope, true
}

// HasDirectoryScope returns a boolean if a field has been set.
func (o *UnifiedRoleAssignment) HasDirectoryScope() bool {
	if o != nil && o.DirectoryScope != nil {
		return true
	}

	return false
}

// SetDirectoryScope gets a reference to the given AnyOfmicrosoftGraphDirectoryObject and assigns it to the DirectoryScope field.
func (o *UnifiedRoleAssignment) SetDirectoryScope(v AnyOfmicrosoftGraphDirectoryObject) {
	o.DirectoryScope = v
}

// GetPrincipal returns the Principal field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UnifiedRoleAssignment) GetPrincipal() AnyOfmicrosoftGraphDirectoryObject {
	if o == nil  {
		var ret AnyOfmicrosoftGraphDirectoryObject
		return ret
	}
	return o.Principal
}

// GetPrincipalOk returns a tuple with the Principal field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UnifiedRoleAssignment) GetPrincipalOk() (*AnyOfmicrosoftGraphDirectoryObject, bool) {
	if o == nil || o.Principal == nil {
		return nil, false
	}
	return &o.Principal, true
}

// HasPrincipal returns a boolean if a field has been set.
func (o *UnifiedRoleAssignment) HasPrincipal() bool {
	if o != nil && o.Principal != nil {
		return true
	}

	return false
}

// SetPrincipal gets a reference to the given AnyOfmicrosoftGraphDirectoryObject and assigns it to the Principal field.
func (o *UnifiedRoleAssignment) SetPrincipal(v AnyOfmicrosoftGraphDirectoryObject) {
	o.Principal = v
}

// GetRoleDefinition returns the RoleDefinition field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UnifiedRoleAssignment) GetRoleDefinition() AnyOfmicrosoftGraphUnifiedRoleDefinition {
	if o == nil  {
		var ret AnyOfmicrosoftGraphUnifiedRoleDefinition
		return ret
	}
	return o.RoleDefinition
}

// GetRoleDefinitionOk returns a tuple with the RoleDefinition field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UnifiedRoleAssignment) GetRoleDefinitionOk() (*AnyOfmicrosoftGraphUnifiedRoleDefinition, bool) {
	if o == nil || o.RoleDefinition == nil {
		return nil, false
	}
	return &o.RoleDefinition, true
}

// HasRoleDefinition returns a boolean if a field has been set.
func (o *UnifiedRoleAssignment) HasRoleDefinition() bool {
	if o != nil && o.RoleDefinition != nil {
		return true
	}

	return false
}

// SetRoleDefinition gets a reference to the given AnyOfmicrosoftGraphUnifiedRoleDefinition and assigns it to the RoleDefinition field.
func (o *UnifiedRoleAssignment) SetRoleDefinition(v AnyOfmicrosoftGraphUnifiedRoleDefinition) {
	o.RoleDefinition = v
}

func (o UnifiedRoleAssignment) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
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

type NullableUnifiedRoleAssignment struct {
	value *UnifiedRoleAssignment
	isSet bool
}

func (v NullableUnifiedRoleAssignment) Get() *UnifiedRoleAssignment {
	return v.value
}

func (v *NullableUnifiedRoleAssignment) Set(val *UnifiedRoleAssignment) {
	v.value = val
	v.isSet = true
}

func (v NullableUnifiedRoleAssignment) IsSet() bool {
	return v.isSet
}

func (v *NullableUnifiedRoleAssignment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnifiedRoleAssignment(val *UnifiedRoleAssignment) *NullableUnifiedRoleAssignment {
	return &NullableUnifiedRoleAssignment{value: val, isSet: true}
}

func (v NullableUnifiedRoleAssignment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnifiedRoleAssignment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


