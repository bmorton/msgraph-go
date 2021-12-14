# MicrosoftGraphUnifiedRoleDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
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

### NewMicrosoftGraphUnifiedRoleDefinition

`func NewMicrosoftGraphUnifiedRoleDefinition() *MicrosoftGraphUnifiedRoleDefinition`

NewMicrosoftGraphUnifiedRoleDefinition instantiates a new MicrosoftGraphUnifiedRoleDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphUnifiedRoleDefinitionWithDefaults

`func NewMicrosoftGraphUnifiedRoleDefinitionWithDefaults() *MicrosoftGraphUnifiedRoleDefinition`

NewMicrosoftGraphUnifiedRoleDefinitionWithDefaults instantiates a new MicrosoftGraphUnifiedRoleDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphUnifiedRoleDefinition) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphUnifiedRoleDefinition) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphUnifiedRoleDefinition) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphUnifiedRoleDefinition) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDescription

`func (o *MicrosoftGraphUnifiedRoleDefinition) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MicrosoftGraphUnifiedRoleDefinition) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MicrosoftGraphUnifiedRoleDefinition) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MicrosoftGraphUnifiedRoleDefinition) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *MicrosoftGraphUnifiedRoleDefinition) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *MicrosoftGraphUnifiedRoleDefinition) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDisplayName

`func (o *MicrosoftGraphUnifiedRoleDefinition) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphUnifiedRoleDefinition) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *MicrosoftGraphUnifiedRoleDefinition) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *MicrosoftGraphUnifiedRoleDefinition) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *MicrosoftGraphUnifiedRoleDefinition) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *MicrosoftGraphUnifiedRoleDefinition) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetIsBuiltIn

`func (o *MicrosoftGraphUnifiedRoleDefinition) GetIsBuiltIn() bool`

GetIsBuiltIn returns the IsBuiltIn field if non-nil, zero value otherwise.

### GetIsBuiltInOk

`func (o *MicrosoftGraphUnifiedRoleDefinition) GetIsBuiltInOk() (*bool, bool)`

GetIsBuiltInOk returns a tuple with the IsBuiltIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBuiltIn

`func (o *MicrosoftGraphUnifiedRoleDefinition) SetIsBuiltIn(v bool)`

SetIsBuiltIn sets IsBuiltIn field to given value.

### HasIsBuiltIn

`func (o *MicrosoftGraphUnifiedRoleDefinition) HasIsBuiltIn() bool`

HasIsBuiltIn returns a boolean if a field has been set.

### SetIsBuiltInNil

`func (o *MicrosoftGraphUnifiedRoleDefinition) SetIsBuiltInNil(b bool)`

 SetIsBuiltInNil sets the value for IsBuiltIn to be an explicit nil

### UnsetIsBuiltIn
`func (o *MicrosoftGraphUnifiedRoleDefinition) UnsetIsBuiltIn()`

UnsetIsBuiltIn ensures that no value is present for IsBuiltIn, not even an explicit nil
### GetIsEnabled

`func (o *MicrosoftGraphUnifiedRoleDefinition) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *MicrosoftGraphUnifiedRoleDefinition) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *MicrosoftGraphUnifiedRoleDefinition) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *MicrosoftGraphUnifiedRoleDefinition) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### SetIsEnabledNil

`func (o *MicrosoftGraphUnifiedRoleDefinition) SetIsEnabledNil(b bool)`

 SetIsEnabledNil sets the value for IsEnabled to be an explicit nil

### UnsetIsEnabled
`func (o *MicrosoftGraphUnifiedRoleDefinition) UnsetIsEnabled()`

UnsetIsEnabled ensures that no value is present for IsEnabled, not even an explicit nil
### GetResourceScopes

`func (o *MicrosoftGraphUnifiedRoleDefinition) GetResourceScopes() []string`

GetResourceScopes returns the ResourceScopes field if non-nil, zero value otherwise.

### GetResourceScopesOk

`func (o *MicrosoftGraphUnifiedRoleDefinition) GetResourceScopesOk() (*[]string, bool)`

GetResourceScopesOk returns a tuple with the ResourceScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceScopes

`func (o *MicrosoftGraphUnifiedRoleDefinition) SetResourceScopes(v []string)`

SetResourceScopes sets ResourceScopes field to given value.

### HasResourceScopes

`func (o *MicrosoftGraphUnifiedRoleDefinition) HasResourceScopes() bool`

HasResourceScopes returns a boolean if a field has been set.

### GetRolePermissions

`func (o *MicrosoftGraphUnifiedRoleDefinition) GetRolePermissions() []MicrosoftGraphUnifiedRolePermission`

GetRolePermissions returns the RolePermissions field if non-nil, zero value otherwise.

### GetRolePermissionsOk

`func (o *MicrosoftGraphUnifiedRoleDefinition) GetRolePermissionsOk() (*[]MicrosoftGraphUnifiedRolePermission, bool)`

GetRolePermissionsOk returns a tuple with the RolePermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRolePermissions

`func (o *MicrosoftGraphUnifiedRoleDefinition) SetRolePermissions(v []MicrosoftGraphUnifiedRolePermission)`

SetRolePermissions sets RolePermissions field to given value.

### HasRolePermissions

`func (o *MicrosoftGraphUnifiedRoleDefinition) HasRolePermissions() bool`

HasRolePermissions returns a boolean if a field has been set.

### GetTemplateId

`func (o *MicrosoftGraphUnifiedRoleDefinition) GetTemplateId() string`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *MicrosoftGraphUnifiedRoleDefinition) GetTemplateIdOk() (*string, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *MicrosoftGraphUnifiedRoleDefinition) SetTemplateId(v string)`

SetTemplateId sets TemplateId field to given value.

### HasTemplateId

`func (o *MicrosoftGraphUnifiedRoleDefinition) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.

### SetTemplateIdNil

`func (o *MicrosoftGraphUnifiedRoleDefinition) SetTemplateIdNil(b bool)`

 SetTemplateIdNil sets the value for TemplateId to be an explicit nil

### UnsetTemplateId
`func (o *MicrosoftGraphUnifiedRoleDefinition) UnsetTemplateId()`

UnsetTemplateId ensures that no value is present for TemplateId, not even an explicit nil
### GetVersion

`func (o *MicrosoftGraphUnifiedRoleDefinition) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *MicrosoftGraphUnifiedRoleDefinition) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *MicrosoftGraphUnifiedRoleDefinition) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *MicrosoftGraphUnifiedRoleDefinition) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *MicrosoftGraphUnifiedRoleDefinition) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *MicrosoftGraphUnifiedRoleDefinition) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetInheritsPermissionsFrom

`func (o *MicrosoftGraphUnifiedRoleDefinition) GetInheritsPermissionsFrom() []MicrosoftGraphUnifiedRoleDefinition`

GetInheritsPermissionsFrom returns the InheritsPermissionsFrom field if non-nil, zero value otherwise.

### GetInheritsPermissionsFromOk

`func (o *MicrosoftGraphUnifiedRoleDefinition) GetInheritsPermissionsFromOk() (*[]MicrosoftGraphUnifiedRoleDefinition, bool)`

GetInheritsPermissionsFromOk returns a tuple with the InheritsPermissionsFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInheritsPermissionsFrom

`func (o *MicrosoftGraphUnifiedRoleDefinition) SetInheritsPermissionsFrom(v []MicrosoftGraphUnifiedRoleDefinition)`

SetInheritsPermissionsFrom sets InheritsPermissionsFrom field to given value.

### HasInheritsPermissionsFrom

`func (o *MicrosoftGraphUnifiedRoleDefinition) HasInheritsPermissionsFrom() bool`

HasInheritsPermissionsFrom returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


