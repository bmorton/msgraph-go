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

// UnifiedRoleDefinition struct for UnifiedRoleDefinition
type UnifiedRoleDefinition struct {
	// The description for the unifiedRoleDefinition. Read-only when isBuiltIn is true.
	Description NullableString `json:"description,omitempty"`
	// The display name for the unifiedRoleDefinition. Read-only when isBuiltIn is true. Required.  Supports $filter (eq, in).
	DisplayName NullableString `json:"displayName,omitempty"`
	// Flag indicating whether the role definition is part of the default set included in Azure Active Directory (Azure AD) or a custom definition. Read-only. Supports $filter (eq, in).
	IsBuiltIn NullableBool `json:"isBuiltIn,omitempty"`
	// Flag indicating whether the role is enabled for assignment. If false the role is not available for assignment. Read-only when isBuiltIn is true.
	IsEnabled NullableBool `json:"isEnabled,omitempty"`
	// List of the scopes or permissions the role definition applies to. Currently only / is supported. Read-only when isBuiltIn is true. DO NOT USE. This will be deprecated soon. Attach scope to role assignment.
	ResourceScopes *[]string `json:"resourceScopes,omitempty"`
	// List of permissions included in the role. Read-only when isBuiltIn is true. Required.
	RolePermissions *[]MicrosoftGraphUnifiedRolePermission `json:"rolePermissions,omitempty"`
	// Custom template identifier that can be set when isBuiltIn is false but is read-only when isBuiltIn is true. This identifier is typically used if one needs an identifier to be the same across different directories.
	TemplateId NullableString `json:"templateId,omitempty"`
	// Indicates version of the role definition. Read-only when isBuiltIn is true.
	Version NullableString `json:"version,omitempty"`
	// Read-only collection of role definitions that the given role definition inherits from. Only Azure AD built-in roles (isBuiltIn is true) support this attribute. Supports $expand.
	InheritsPermissionsFrom *[]MicrosoftGraphUnifiedRoleDefinition `json:"inheritsPermissionsFrom,omitempty"`
}

// NewUnifiedRoleDefinition instantiates a new UnifiedRoleDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnifiedRoleDefinition() *UnifiedRoleDefinition {
	this := UnifiedRoleDefinition{}
	return &this
}

// NewUnifiedRoleDefinitionWithDefaults instantiates a new UnifiedRoleDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnifiedRoleDefinitionWithDefaults() *UnifiedRoleDefinition {
	this := UnifiedRoleDefinition{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UnifiedRoleDefinition) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UnifiedRoleDefinition) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *UnifiedRoleDefinition) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *UnifiedRoleDefinition) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *UnifiedRoleDefinition) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *UnifiedRoleDefinition) UnsetDescription() {
	o.Description.Unset()
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UnifiedRoleDefinition) GetDisplayName() string {
	if o == nil || o.DisplayName.Get() == nil {
		var ret string
		return ret
	}
	return *o.DisplayName.Get()
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UnifiedRoleDefinition) GetDisplayNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DisplayName.Get(), o.DisplayName.IsSet()
}

// HasDisplayName returns a boolean if a field has been set.
func (o *UnifiedRoleDefinition) HasDisplayName() bool {
	if o != nil && o.DisplayName.IsSet() {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given NullableString and assigns it to the DisplayName field.
func (o *UnifiedRoleDefinition) SetDisplayName(v string) {
	o.DisplayName.Set(&v)
}
// SetDisplayNameNil sets the value for DisplayName to be an explicit nil
func (o *UnifiedRoleDefinition) SetDisplayNameNil() {
	o.DisplayName.Set(nil)
}

// UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
func (o *UnifiedRoleDefinition) UnsetDisplayName() {
	o.DisplayName.Unset()
}

// GetIsBuiltIn returns the IsBuiltIn field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UnifiedRoleDefinition) GetIsBuiltIn() bool {
	if o == nil || o.IsBuiltIn.Get() == nil {
		var ret bool
		return ret
	}
	return *o.IsBuiltIn.Get()
}

// GetIsBuiltInOk returns a tuple with the IsBuiltIn field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UnifiedRoleDefinition) GetIsBuiltInOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IsBuiltIn.Get(), o.IsBuiltIn.IsSet()
}

// HasIsBuiltIn returns a boolean if a field has been set.
func (o *UnifiedRoleDefinition) HasIsBuiltIn() bool {
	if o != nil && o.IsBuiltIn.IsSet() {
		return true
	}

	return false
}

// SetIsBuiltIn gets a reference to the given NullableBool and assigns it to the IsBuiltIn field.
func (o *UnifiedRoleDefinition) SetIsBuiltIn(v bool) {
	o.IsBuiltIn.Set(&v)
}
// SetIsBuiltInNil sets the value for IsBuiltIn to be an explicit nil
func (o *UnifiedRoleDefinition) SetIsBuiltInNil() {
	o.IsBuiltIn.Set(nil)
}

// UnsetIsBuiltIn ensures that no value is present for IsBuiltIn, not even an explicit nil
func (o *UnifiedRoleDefinition) UnsetIsBuiltIn() {
	o.IsBuiltIn.Unset()
}

// GetIsEnabled returns the IsEnabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UnifiedRoleDefinition) GetIsEnabled() bool {
	if o == nil || o.IsEnabled.Get() == nil {
		var ret bool
		return ret
	}
	return *o.IsEnabled.Get()
}

// GetIsEnabledOk returns a tuple with the IsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UnifiedRoleDefinition) GetIsEnabledOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IsEnabled.Get(), o.IsEnabled.IsSet()
}

// HasIsEnabled returns a boolean if a field has been set.
func (o *UnifiedRoleDefinition) HasIsEnabled() bool {
	if o != nil && o.IsEnabled.IsSet() {
		return true
	}

	return false
}

// SetIsEnabled gets a reference to the given NullableBool and assigns it to the IsEnabled field.
func (o *UnifiedRoleDefinition) SetIsEnabled(v bool) {
	o.IsEnabled.Set(&v)
}
// SetIsEnabledNil sets the value for IsEnabled to be an explicit nil
func (o *UnifiedRoleDefinition) SetIsEnabledNil() {
	o.IsEnabled.Set(nil)
}

// UnsetIsEnabled ensures that no value is present for IsEnabled, not even an explicit nil
func (o *UnifiedRoleDefinition) UnsetIsEnabled() {
	o.IsEnabled.Unset()
}

// GetResourceScopes returns the ResourceScopes field value if set, zero value otherwise.
func (o *UnifiedRoleDefinition) GetResourceScopes() []string {
	if o == nil || o.ResourceScopes == nil {
		var ret []string
		return ret
	}
	return *o.ResourceScopes
}

// GetResourceScopesOk returns a tuple with the ResourceScopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnifiedRoleDefinition) GetResourceScopesOk() (*[]string, bool) {
	if o == nil || o.ResourceScopes == nil {
		return nil, false
	}
	return o.ResourceScopes, true
}

// HasResourceScopes returns a boolean if a field has been set.
func (o *UnifiedRoleDefinition) HasResourceScopes() bool {
	if o != nil && o.ResourceScopes != nil {
		return true
	}

	return false
}

// SetResourceScopes gets a reference to the given []string and assigns it to the ResourceScopes field.
func (o *UnifiedRoleDefinition) SetResourceScopes(v []string) {
	o.ResourceScopes = &v
}

// GetRolePermissions returns the RolePermissions field value if set, zero value otherwise.
func (o *UnifiedRoleDefinition) GetRolePermissions() []MicrosoftGraphUnifiedRolePermission {
	if o == nil || o.RolePermissions == nil {
		var ret []MicrosoftGraphUnifiedRolePermission
		return ret
	}
	return *o.RolePermissions
}

// GetRolePermissionsOk returns a tuple with the RolePermissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnifiedRoleDefinition) GetRolePermissionsOk() (*[]MicrosoftGraphUnifiedRolePermission, bool) {
	if o == nil || o.RolePermissions == nil {
		return nil, false
	}
	return o.RolePermissions, true
}

// HasRolePermissions returns a boolean if a field has been set.
func (o *UnifiedRoleDefinition) HasRolePermissions() bool {
	if o != nil && o.RolePermissions != nil {
		return true
	}

	return false
}

// SetRolePermissions gets a reference to the given []MicrosoftGraphUnifiedRolePermission and assigns it to the RolePermissions field.
func (o *UnifiedRoleDefinition) SetRolePermissions(v []MicrosoftGraphUnifiedRolePermission) {
	o.RolePermissions = &v
}

// GetTemplateId returns the TemplateId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UnifiedRoleDefinition) GetTemplateId() string {
	if o == nil || o.TemplateId.Get() == nil {
		var ret string
		return ret
	}
	return *o.TemplateId.Get()
}

// GetTemplateIdOk returns a tuple with the TemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UnifiedRoleDefinition) GetTemplateIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TemplateId.Get(), o.TemplateId.IsSet()
}

// HasTemplateId returns a boolean if a field has been set.
func (o *UnifiedRoleDefinition) HasTemplateId() bool {
	if o != nil && o.TemplateId.IsSet() {
		return true
	}

	return false
}

// SetTemplateId gets a reference to the given NullableString and assigns it to the TemplateId field.
func (o *UnifiedRoleDefinition) SetTemplateId(v string) {
	o.TemplateId.Set(&v)
}
// SetTemplateIdNil sets the value for TemplateId to be an explicit nil
func (o *UnifiedRoleDefinition) SetTemplateIdNil() {
	o.TemplateId.Set(nil)
}

// UnsetTemplateId ensures that no value is present for TemplateId, not even an explicit nil
func (o *UnifiedRoleDefinition) UnsetTemplateId() {
	o.TemplateId.Unset()
}

// GetVersion returns the Version field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UnifiedRoleDefinition) GetVersion() string {
	if o == nil || o.Version.Get() == nil {
		var ret string
		return ret
	}
	return *o.Version.Get()
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UnifiedRoleDefinition) GetVersionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Version.Get(), o.Version.IsSet()
}

// HasVersion returns a boolean if a field has been set.
func (o *UnifiedRoleDefinition) HasVersion() bool {
	if o != nil && o.Version.IsSet() {
		return true
	}

	return false
}

// SetVersion gets a reference to the given NullableString and assigns it to the Version field.
func (o *UnifiedRoleDefinition) SetVersion(v string) {
	o.Version.Set(&v)
}
// SetVersionNil sets the value for Version to be an explicit nil
func (o *UnifiedRoleDefinition) SetVersionNil() {
	o.Version.Set(nil)
}

// UnsetVersion ensures that no value is present for Version, not even an explicit nil
func (o *UnifiedRoleDefinition) UnsetVersion() {
	o.Version.Unset()
}

// GetInheritsPermissionsFrom returns the InheritsPermissionsFrom field value if set, zero value otherwise.
func (o *UnifiedRoleDefinition) GetInheritsPermissionsFrom() []MicrosoftGraphUnifiedRoleDefinition {
	if o == nil || o.InheritsPermissionsFrom == nil {
		var ret []MicrosoftGraphUnifiedRoleDefinition
		return ret
	}
	return *o.InheritsPermissionsFrom
}

// GetInheritsPermissionsFromOk returns a tuple with the InheritsPermissionsFrom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnifiedRoleDefinition) GetInheritsPermissionsFromOk() (*[]MicrosoftGraphUnifiedRoleDefinition, bool) {
	if o == nil || o.InheritsPermissionsFrom == nil {
		return nil, false
	}
	return o.InheritsPermissionsFrom, true
}

// HasInheritsPermissionsFrom returns a boolean if a field has been set.
func (o *UnifiedRoleDefinition) HasInheritsPermissionsFrom() bool {
	if o != nil && o.InheritsPermissionsFrom != nil {
		return true
	}

	return false
}

// SetInheritsPermissionsFrom gets a reference to the given []MicrosoftGraphUnifiedRoleDefinition and assigns it to the InheritsPermissionsFrom field.
func (o *UnifiedRoleDefinition) SetInheritsPermissionsFrom(v []MicrosoftGraphUnifiedRoleDefinition) {
	o.InheritsPermissionsFrom = &v
}

func (o UnifiedRoleDefinition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.DisplayName.IsSet() {
		toSerialize["displayName"] = o.DisplayName.Get()
	}
	if o.IsBuiltIn.IsSet() {
		toSerialize["isBuiltIn"] = o.IsBuiltIn.Get()
	}
	if o.IsEnabled.IsSet() {
		toSerialize["isEnabled"] = o.IsEnabled.Get()
	}
	if o.ResourceScopes != nil {
		toSerialize["resourceScopes"] = o.ResourceScopes
	}
	if o.RolePermissions != nil {
		toSerialize["rolePermissions"] = o.RolePermissions
	}
	if o.TemplateId.IsSet() {
		toSerialize["templateId"] = o.TemplateId.Get()
	}
	if o.Version.IsSet() {
		toSerialize["version"] = o.Version.Get()
	}
	if o.InheritsPermissionsFrom != nil {
		toSerialize["inheritsPermissionsFrom"] = o.InheritsPermissionsFrom
	}
	return json.Marshal(toSerialize)
}

type NullableUnifiedRoleDefinition struct {
	value *UnifiedRoleDefinition
	isSet bool
}

func (v NullableUnifiedRoleDefinition) Get() *UnifiedRoleDefinition {
	return v.value
}

func (v *NullableUnifiedRoleDefinition) Set(val *UnifiedRoleDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableUnifiedRoleDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableUnifiedRoleDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnifiedRoleDefinition(val *UnifiedRoleDefinition) *NullableUnifiedRoleDefinition {
	return &NullableUnifiedRoleDefinition{value: val, isSet: true}
}

func (v NullableUnifiedRoleDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnifiedRoleDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


