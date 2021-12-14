# MicrosoftGraphAdministrativeUnit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**DeletedDateTime** | Pointer to **NullableTime** |  | [optional] 
**Description** | Pointer to **NullableString** | An optional description for the administrative unit. Supports $filter (eq, ne, in, startsWith). | [optional] 
**DisplayName** | Pointer to **NullableString** | Display name for the administrative unit. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values), $search, and $orderBy. | [optional] 
**Visibility** | Pointer to **NullableString** | Controls whether the administrative unit and its members are hidden or public. Can be set to HiddenMembership or Public. If not set, default behavior is Public. When set to HiddenMembership, only members of the administrative unit can list other members of the administrative unit. | [optional] 
**Members** | Pointer to [**[]MicrosoftGraphDirectoryObject**](MicrosoftGraphDirectoryObject.md) | Users and groups that are members of this administrative unit. HTTP Methods: GET (list members), POST (add members), DELETE (remove members). | [optional] 
**ScopedRoleMembers** | Pointer to [**[]MicrosoftGraphScopedRoleMembership**](MicrosoftGraphScopedRoleMembership.md) | Scoped-role members of this administrative unit.  HTTP Methods: GET (list scopedRoleMemberships), POST (add scopedRoleMembership), DELETE (remove scopedRoleMembership). | [optional] 
**Extensions** | Pointer to [**[]MicrosoftGraphExtension**](MicrosoftGraphExtension.md) | The collection of open extensions defined for this administrative unit. Nullable. | [optional] 

## Methods

### NewMicrosoftGraphAdministrativeUnit

`func NewMicrosoftGraphAdministrativeUnit() *MicrosoftGraphAdministrativeUnit`

NewMicrosoftGraphAdministrativeUnit instantiates a new MicrosoftGraphAdministrativeUnit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphAdministrativeUnitWithDefaults

`func NewMicrosoftGraphAdministrativeUnitWithDefaults() *MicrosoftGraphAdministrativeUnit`

NewMicrosoftGraphAdministrativeUnitWithDefaults instantiates a new MicrosoftGraphAdministrativeUnit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphAdministrativeUnit) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphAdministrativeUnit) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphAdministrativeUnit) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphAdministrativeUnit) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDeletedDateTime

`func (o *MicrosoftGraphAdministrativeUnit) GetDeletedDateTime() time.Time`

GetDeletedDateTime returns the DeletedDateTime field if non-nil, zero value otherwise.

### GetDeletedDateTimeOk

`func (o *MicrosoftGraphAdministrativeUnit) GetDeletedDateTimeOk() (*time.Time, bool)`

GetDeletedDateTimeOk returns a tuple with the DeletedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedDateTime

`func (o *MicrosoftGraphAdministrativeUnit) SetDeletedDateTime(v time.Time)`

SetDeletedDateTime sets DeletedDateTime field to given value.

### HasDeletedDateTime

`func (o *MicrosoftGraphAdministrativeUnit) HasDeletedDateTime() bool`

HasDeletedDateTime returns a boolean if a field has been set.

### SetDeletedDateTimeNil

`func (o *MicrosoftGraphAdministrativeUnit) SetDeletedDateTimeNil(b bool)`

 SetDeletedDateTimeNil sets the value for DeletedDateTime to be an explicit nil

### UnsetDeletedDateTime
`func (o *MicrosoftGraphAdministrativeUnit) UnsetDeletedDateTime()`

UnsetDeletedDateTime ensures that no value is present for DeletedDateTime, not even an explicit nil
### GetDescription

`func (o *MicrosoftGraphAdministrativeUnit) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MicrosoftGraphAdministrativeUnit) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MicrosoftGraphAdministrativeUnit) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MicrosoftGraphAdministrativeUnit) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *MicrosoftGraphAdministrativeUnit) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *MicrosoftGraphAdministrativeUnit) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDisplayName

`func (o *MicrosoftGraphAdministrativeUnit) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphAdministrativeUnit) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *MicrosoftGraphAdministrativeUnit) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *MicrosoftGraphAdministrativeUnit) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *MicrosoftGraphAdministrativeUnit) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *MicrosoftGraphAdministrativeUnit) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetVisibility

`func (o *MicrosoftGraphAdministrativeUnit) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *MicrosoftGraphAdministrativeUnit) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *MicrosoftGraphAdministrativeUnit) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *MicrosoftGraphAdministrativeUnit) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### SetVisibilityNil

`func (o *MicrosoftGraphAdministrativeUnit) SetVisibilityNil(b bool)`

 SetVisibilityNil sets the value for Visibility to be an explicit nil

### UnsetVisibility
`func (o *MicrosoftGraphAdministrativeUnit) UnsetVisibility()`

UnsetVisibility ensures that no value is present for Visibility, not even an explicit nil
### GetMembers

`func (o *MicrosoftGraphAdministrativeUnit) GetMembers() []MicrosoftGraphDirectoryObject`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *MicrosoftGraphAdministrativeUnit) GetMembersOk() (*[]MicrosoftGraphDirectoryObject, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *MicrosoftGraphAdministrativeUnit) SetMembers(v []MicrosoftGraphDirectoryObject)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *MicrosoftGraphAdministrativeUnit) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### GetScopedRoleMembers

`func (o *MicrosoftGraphAdministrativeUnit) GetScopedRoleMembers() []MicrosoftGraphScopedRoleMembership`

GetScopedRoleMembers returns the ScopedRoleMembers field if non-nil, zero value otherwise.

### GetScopedRoleMembersOk

`func (o *MicrosoftGraphAdministrativeUnit) GetScopedRoleMembersOk() (*[]MicrosoftGraphScopedRoleMembership, bool)`

GetScopedRoleMembersOk returns a tuple with the ScopedRoleMembers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopedRoleMembers

`func (o *MicrosoftGraphAdministrativeUnit) SetScopedRoleMembers(v []MicrosoftGraphScopedRoleMembership)`

SetScopedRoleMembers sets ScopedRoleMembers field to given value.

### HasScopedRoleMembers

`func (o *MicrosoftGraphAdministrativeUnit) HasScopedRoleMembers() bool`

HasScopedRoleMembers returns a boolean if a field has been set.

### GetExtensions

`func (o *MicrosoftGraphAdministrativeUnit) GetExtensions() []MicrosoftGraphExtension`

GetExtensions returns the Extensions field if non-nil, zero value otherwise.

### GetExtensionsOk

`func (o *MicrosoftGraphAdministrativeUnit) GetExtensionsOk() (*[]MicrosoftGraphExtension, bool)`

GetExtensionsOk returns a tuple with the Extensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensions

`func (o *MicrosoftGraphAdministrativeUnit) SetExtensions(v []MicrosoftGraphExtension)`

SetExtensions sets Extensions field to given value.

### HasExtensions

`func (o *MicrosoftGraphAdministrativeUnit) HasExtensions() bool`

HasExtensions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


