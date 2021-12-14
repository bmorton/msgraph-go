# UnifiedRoleDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **NullableString** | The description for the unifiedRoleDefinition. Read-only when isBuiltIn is true. | [optional] 
**DisplayName** | Pointer to **NullableString** | The display name for the unifiedRoleDefinition. Read-only when isBuiltIn is true. Required.  Supports $filter (eq, in). | [optional] 
**IsBuiltIn** | Pointer to **NullableBool** | Flag indicating whether the role definition is part of the default set included in Azure Active Directory (Azure AD) or a custom definition. Read-only. Supports $filter (eq, in). | [optional] 
**IsEnabled** | Pointer to **NullableBool** | Flag indicating whether the role is enabled for assignment. If false the role is not available for assignment. Read-only when isBuiltIn is true. | [optional] 
**ResourceScopes** | Pointer to **[]string** | List of the scopes or permissions the role definition applies to. Currently only / is supported. Read-only when isBuiltIn is true. DO NOT USE. This will be deprecated soon. Attach scope to role assignment. | [optional] 
**RolePermissions** | Pointer to [**[]MicrosoftGraphUnifiedRolePermission**](MicrosoftGraphUnifiedRolePermission.md) | List of permissions included in the role. Read-only when isBuiltIn is true. Required. | [optional] 
**TemplateId** | Pointer to **NullableString** | Custom template identifier that can be set when isBuiltIn is false but is read-only when isBuiltIn is true. This identifier is typically used if one needs an identifier to be the same across different directories. | [optional] 
**Version** | Pointer to **NullableString** | Indicates version of the role definition. Read-only when isBuiltIn is true. | [optional] 
**InheritsPermissionsFrom** | Pointer to [**[]MicrosoftGraphUnifiedRoleDefinition**](MicrosoftGraphUnifiedRoleDefinition.md) | Read-only collection of role definitions that the given role definition inherits from. Only Azure AD built-in roles (isBuiltIn is true) support this attribute. Supports $expand. | [optional] 

## Methods

### NewUnifiedRoleDefinition

`func NewUnifiedRoleDefinition() *UnifiedRoleDefinition`

NewUnifiedRoleDefinition instantiates a new UnifiedRoleDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnifiedRoleDefinitionWithDefaults

`func NewUnifiedRoleDefinitionWithDefaults() *UnifiedRoleDefinition`

NewUnifiedRoleDefinitionWithDefaults instantiates a new UnifiedRoleDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *UnifiedRoleDefinition) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UnifiedRoleDefinition) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UnifiedRoleDefinition) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UnifiedRoleDefinition) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *UnifiedRoleDefinition) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *UnifiedRoleDefinition) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDisplayName

`func (o *UnifiedRoleDefinition) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *UnifiedRoleDefinition) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *UnifiedRoleDefinition) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *UnifiedRoleDefinition) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *UnifiedRoleDefinition) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *UnifiedRoleDefinition) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetIsBuiltIn

`func (o *UnifiedRoleDefinition) GetIsBuiltIn() bool`

GetIsBuiltIn returns the IsBuiltIn field if non-nil, zero value otherwise.

### GetIsBuiltInOk

`func (o *UnifiedRoleDefinition) GetIsBuiltInOk() (*bool, bool)`

GetIsBuiltInOk returns a tuple with the IsBuiltIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBuiltIn

`func (o *UnifiedRoleDefinition) SetIsBuiltIn(v bool)`

SetIsBuiltIn sets IsBuiltIn field to given value.

### HasIsBuiltIn

`func (o *UnifiedRoleDefinition) HasIsBuiltIn() bool`

HasIsBuiltIn returns a boolean if a field has been set.

### SetIsBuiltInNil

`func (o *UnifiedRoleDefinition) SetIsBuiltInNil(b bool)`

 SetIsBuiltInNil sets the value for IsBuiltIn to be an explicit nil

### UnsetIsBuiltIn
`func (o *UnifiedRoleDefinition) UnsetIsBuiltIn()`

UnsetIsBuiltIn ensures that no value is present for IsBuiltIn, not even an explicit nil
### GetIsEnabled

`func (o *UnifiedRoleDefinition) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *UnifiedRoleDefinition) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *UnifiedRoleDefinition) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *UnifiedRoleDefinition) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### SetIsEnabledNil

`func (o *UnifiedRoleDefinition) SetIsEnabledNil(b bool)`

 SetIsEnabledNil sets the value for IsEnabled to be an explicit nil

### UnsetIsEnabled
`func (o *UnifiedRoleDefinition) UnsetIsEnabled()`

UnsetIsEnabled ensures that no value is present for IsEnabled, not even an explicit nil
### GetResourceScopes

`func (o *UnifiedRoleDefinition) GetResourceScopes() []string`

GetResourceScopes returns the ResourceScopes field if non-nil, zero value otherwise.

### GetResourceScopesOk

`func (o *UnifiedRoleDefinition) GetResourceScopesOk() (*[]string, bool)`

GetResourceScopesOk returns a tuple with the ResourceScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceScopes

`func (o *UnifiedRoleDefinition) SetResourceScopes(v []string)`

SetResourceScopes sets ResourceScopes field to given value.

### HasResourceScopes

`func (o *UnifiedRoleDefinition) HasResourceScopes() bool`

HasResourceScopes returns a boolean if a field has been set.

### GetRolePermissions

`func (o *UnifiedRoleDefinition) GetRolePermissions() []MicrosoftGraphUnifiedRolePermission`

GetRolePermissions returns the RolePermissions field if non-nil, zero value otherwise.

### GetRolePermissionsOk

`func (o *UnifiedRoleDefinition) GetRolePermissionsOk() (*[]MicrosoftGraphUnifiedRolePermission, bool)`

GetRolePermissionsOk returns a tuple with the RolePermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRolePermissions

`func (o *UnifiedRoleDefinition) SetRolePermissions(v []MicrosoftGraphUnifiedRolePermission)`

SetRolePermissions sets RolePermissions field to given value.

### HasRolePermissions

`func (o *UnifiedRoleDefinition) HasRolePermissions() bool`

HasRolePermissions returns a boolean if a field has been set.

### GetTemplateId

`func (o *UnifiedRoleDefinition) GetTemplateId() string`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *UnifiedRoleDefinition) GetTemplateIdOk() (*string, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *UnifiedRoleDefinition) SetTemplateId(v string)`

SetTemplateId sets TemplateId field to given value.

### HasTemplateId

`func (o *UnifiedRoleDefinition) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.

### SetTemplateIdNil

`func (o *UnifiedRoleDefinition) SetTemplateIdNil(b bool)`

 SetTemplateIdNil sets the value for TemplateId to be an explicit nil

### UnsetTemplateId
`func (o *UnifiedRoleDefinition) UnsetTemplateId()`

UnsetTemplateId ensures that no value is present for TemplateId, not even an explicit nil
### GetVersion

`func (o *UnifiedRoleDefinition) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *UnifiedRoleDefinition) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *UnifiedRoleDefinition) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *UnifiedRoleDefinition) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *UnifiedRoleDefinition) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *UnifiedRoleDefinition) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetInheritsPermissionsFrom

`func (o *UnifiedRoleDefinition) GetInheritsPermissionsFrom() []MicrosoftGraphUnifiedRoleDefinition`

GetInheritsPermissionsFrom returns the InheritsPermissionsFrom field if non-nil, zero value otherwise.

### GetInheritsPermissionsFromOk

`func (o *UnifiedRoleDefinition) GetInheritsPermissionsFromOk() (*[]MicrosoftGraphUnifiedRoleDefinition, bool)`

GetInheritsPermissionsFromOk returns a tuple with the InheritsPermissionsFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInheritsPermissionsFrom

`func (o *UnifiedRoleDefinition) SetInheritsPermissionsFrom(v []MicrosoftGraphUnifiedRoleDefinition)`

SetInheritsPermissionsFrom sets InheritsPermissionsFrom field to given value.

### HasInheritsPermissionsFrom

`func (o *UnifiedRoleDefinition) HasInheritsPermissionsFrom() bool`

HasInheritsPermissionsFrom returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


