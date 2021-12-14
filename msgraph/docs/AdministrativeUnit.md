# AdministrativeUnit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **NullableString** | An optional description for the administrative unit. Supports $filter (eq, ne, in, startsWith). | [optional] 
**DisplayName** | Pointer to **NullableString** | Display name for the administrative unit. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values), $search, and $orderBy. | [optional] 
**Visibility** | Pointer to **NullableString** | Controls whether the administrative unit and its members are hidden or public. Can be set to HiddenMembership or Public. If not set, default behavior is Public. When set to HiddenMembership, only members of the administrative unit can list other members of the administrative unit. | [optional] 
**Members** | Pointer to [**[]MicrosoftGraphDirectoryObject**](MicrosoftGraphDirectoryObject.md) | Users and groups that are members of this administrative unit. HTTP Methods: GET (list members), POST (add members), DELETE (remove members). | [optional] 
**ScopedRoleMembers** | Pointer to [**[]MicrosoftGraphScopedRoleMembership**](MicrosoftGraphScopedRoleMembership.md) | Scoped-role members of this administrative unit.  HTTP Methods: GET (list scopedRoleMemberships), POST (add scopedRoleMembership), DELETE (remove scopedRoleMembership). | [optional] 
**Extensions** | Pointer to [**[]MicrosoftGraphExtension**](MicrosoftGraphExtension.md) | The collection of open extensions defined for this administrative unit. Nullable. | [optional] 

## Methods

### NewAdministrativeUnit

`func NewAdministrativeUnit() *AdministrativeUnit`

NewAdministrativeUnit instantiates a new AdministrativeUnit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdministrativeUnitWithDefaults

`func NewAdministrativeUnitWithDefaults() *AdministrativeUnit`

NewAdministrativeUnitWithDefaults instantiates a new AdministrativeUnit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *AdministrativeUnit) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AdministrativeUnit) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AdministrativeUnit) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AdministrativeUnit) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *AdministrativeUnit) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *AdministrativeUnit) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDisplayName

`func (o *AdministrativeUnit) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *AdministrativeUnit) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *AdministrativeUnit) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *AdministrativeUnit) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *AdministrativeUnit) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *AdministrativeUnit) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetVisibility

`func (o *AdministrativeUnit) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *AdministrativeUnit) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *AdministrativeUnit) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *AdministrativeUnit) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### SetVisibilityNil

`func (o *AdministrativeUnit) SetVisibilityNil(b bool)`

 SetVisibilityNil sets the value for Visibility to be an explicit nil

### UnsetVisibility
`func (o *AdministrativeUnit) UnsetVisibility()`

UnsetVisibility ensures that no value is present for Visibility, not even an explicit nil
### GetMembers

`func (o *AdministrativeUnit) GetMembers() []MicrosoftGraphDirectoryObject`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *AdministrativeUnit) GetMembersOk() (*[]MicrosoftGraphDirectoryObject, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *AdministrativeUnit) SetMembers(v []MicrosoftGraphDirectoryObject)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *AdministrativeUnit) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### GetScopedRoleMembers

`func (o *AdministrativeUnit) GetScopedRoleMembers() []MicrosoftGraphScopedRoleMembership`

GetScopedRoleMembers returns the ScopedRoleMembers field if non-nil, zero value otherwise.

### GetScopedRoleMembersOk

`func (o *AdministrativeUnit) GetScopedRoleMembersOk() (*[]MicrosoftGraphScopedRoleMembership, bool)`

GetScopedRoleMembersOk returns a tuple with the ScopedRoleMembers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopedRoleMembers

`func (o *AdministrativeUnit) SetScopedRoleMembers(v []MicrosoftGraphScopedRoleMembership)`

SetScopedRoleMembers sets ScopedRoleMembers field to given value.

### HasScopedRoleMembers

`func (o *AdministrativeUnit) HasScopedRoleMembers() bool`

HasScopedRoleMembers returns a boolean if a field has been set.

### GetExtensions

`func (o *AdministrativeUnit) GetExtensions() []MicrosoftGraphExtension`

GetExtensions returns the Extensions field if non-nil, zero value otherwise.

### GetExtensionsOk

`func (o *AdministrativeUnit) GetExtensionsOk() (*[]MicrosoftGraphExtension, bool)`

GetExtensionsOk returns a tuple with the Extensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensions

`func (o *AdministrativeUnit) SetExtensions(v []MicrosoftGraphExtension)`

SetExtensions sets Extensions field to given value.

### HasExtensions

`func (o *AdministrativeUnit) HasExtensions() bool`

HasExtensions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


