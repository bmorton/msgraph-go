# MicrosoftGraphPermissionScope

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdminConsentDescription** | Pointer to **NullableString** | A description of the delegated permissions, intended to be read by an administrator granting the permission on behalf of all users. This text appears in tenant-wide admin consent experiences. | [optional] 
**AdminConsentDisplayName** | Pointer to **NullableString** | The permission&#39;s title, intended to be read by an administrator granting the permission on behalf of all users. | [optional] 
**Id** | Pointer to **string** | Unique delegated permission identifier inside the collection of delegated permissions defined for a resource application. | [optional] 
**IsEnabled** | Pointer to **bool** | When creating or updating a permission, this property must be set to true (which is the default). To delete a permission, this property must first be set to false.  At that point, in a subsequent call, the permission may be removed. | [optional] 
**Origin** | Pointer to **NullableString** |  | [optional] 
**Type** | Pointer to **NullableString** | Specifies whether this delegated permission should be considered safe for non-admin users to consent to on behalf of themselves, or whether an administrator should be required for consent to the permissions. This will be the default behavior, but each customer can choose to customize the behavior in their organization (by allowing, restricting or limiting user consent to this delegated permission.) | [optional] 
**UserConsentDescription** | Pointer to **NullableString** | A description of the delegated permissions, intended to be read by a user granting the permission on their own behalf. This text appears in consent experiences where the user is consenting only on behalf of themselves. | [optional] 
**UserConsentDisplayName** | Pointer to **NullableString** | A title for the permission, intended to be read by a user granting the permission on their own behalf. This text appears in consent experiences where the user is consenting only on behalf of themselves. | [optional] 
**Value** | Pointer to **NullableString** | Specifies the value to include in the scp (scope) claim in access tokens. Must not exceed 120 characters in length. Allowed characters are : ! # $ % &amp; &#39; ( ) * + , - . / : ;  &#x3D;  ? @ [ ] ^ + _  {  } ~, as well as characters in the ranges 0-9, A-Z and a-z. Any other character, including the space character, are not allowed. May not begin with .. | [optional] 

## Methods

### NewMicrosoftGraphPermissionScope

`func NewMicrosoftGraphPermissionScope() *MicrosoftGraphPermissionScope`

NewMicrosoftGraphPermissionScope instantiates a new MicrosoftGraphPermissionScope object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphPermissionScopeWithDefaults

`func NewMicrosoftGraphPermissionScopeWithDefaults() *MicrosoftGraphPermissionScope`

NewMicrosoftGraphPermissionScopeWithDefaults instantiates a new MicrosoftGraphPermissionScope object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdminConsentDescription

`func (o *MicrosoftGraphPermissionScope) GetAdminConsentDescription() string`

GetAdminConsentDescription returns the AdminConsentDescription field if non-nil, zero value otherwise.

### GetAdminConsentDescriptionOk

`func (o *MicrosoftGraphPermissionScope) GetAdminConsentDescriptionOk() (*string, bool)`

GetAdminConsentDescriptionOk returns a tuple with the AdminConsentDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminConsentDescription

`func (o *MicrosoftGraphPermissionScope) SetAdminConsentDescription(v string)`

SetAdminConsentDescription sets AdminConsentDescription field to given value.

### HasAdminConsentDescription

`func (o *MicrosoftGraphPermissionScope) HasAdminConsentDescription() bool`

HasAdminConsentDescription returns a boolean if a field has been set.

### SetAdminConsentDescriptionNil

`func (o *MicrosoftGraphPermissionScope) SetAdminConsentDescriptionNil(b bool)`

 SetAdminConsentDescriptionNil sets the value for AdminConsentDescription to be an explicit nil

### UnsetAdminConsentDescription
`func (o *MicrosoftGraphPermissionScope) UnsetAdminConsentDescription()`

UnsetAdminConsentDescription ensures that no value is present for AdminConsentDescription, not even an explicit nil
### GetAdminConsentDisplayName

`func (o *MicrosoftGraphPermissionScope) GetAdminConsentDisplayName() string`

GetAdminConsentDisplayName returns the AdminConsentDisplayName field if non-nil, zero value otherwise.

### GetAdminConsentDisplayNameOk

`func (o *MicrosoftGraphPermissionScope) GetAdminConsentDisplayNameOk() (*string, bool)`

GetAdminConsentDisplayNameOk returns a tuple with the AdminConsentDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminConsentDisplayName

`func (o *MicrosoftGraphPermissionScope) SetAdminConsentDisplayName(v string)`

SetAdminConsentDisplayName sets AdminConsentDisplayName field to given value.

### HasAdminConsentDisplayName

`func (o *MicrosoftGraphPermissionScope) HasAdminConsentDisplayName() bool`

HasAdminConsentDisplayName returns a boolean if a field has been set.

### SetAdminConsentDisplayNameNil

`func (o *MicrosoftGraphPermissionScope) SetAdminConsentDisplayNameNil(b bool)`

 SetAdminConsentDisplayNameNil sets the value for AdminConsentDisplayName to be an explicit nil

### UnsetAdminConsentDisplayName
`func (o *MicrosoftGraphPermissionScope) UnsetAdminConsentDisplayName()`

UnsetAdminConsentDisplayName ensures that no value is present for AdminConsentDisplayName, not even an explicit nil
### GetId

`func (o *MicrosoftGraphPermissionScope) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphPermissionScope) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphPermissionScope) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphPermissionScope) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsEnabled

`func (o *MicrosoftGraphPermissionScope) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *MicrosoftGraphPermissionScope) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *MicrosoftGraphPermissionScope) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *MicrosoftGraphPermissionScope) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetOrigin

`func (o *MicrosoftGraphPermissionScope) GetOrigin() string`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *MicrosoftGraphPermissionScope) GetOriginOk() (*string, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *MicrosoftGraphPermissionScope) SetOrigin(v string)`

SetOrigin sets Origin field to given value.

### HasOrigin

`func (o *MicrosoftGraphPermissionScope) HasOrigin() bool`

HasOrigin returns a boolean if a field has been set.

### SetOriginNil

`func (o *MicrosoftGraphPermissionScope) SetOriginNil(b bool)`

 SetOriginNil sets the value for Origin to be an explicit nil

### UnsetOrigin
`func (o *MicrosoftGraphPermissionScope) UnsetOrigin()`

UnsetOrigin ensures that no value is present for Origin, not even an explicit nil
### GetType

`func (o *MicrosoftGraphPermissionScope) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MicrosoftGraphPermissionScope) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MicrosoftGraphPermissionScope) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *MicrosoftGraphPermissionScope) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *MicrosoftGraphPermissionScope) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *MicrosoftGraphPermissionScope) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetUserConsentDescription

`func (o *MicrosoftGraphPermissionScope) GetUserConsentDescription() string`

GetUserConsentDescription returns the UserConsentDescription field if non-nil, zero value otherwise.

### GetUserConsentDescriptionOk

`func (o *MicrosoftGraphPermissionScope) GetUserConsentDescriptionOk() (*string, bool)`

GetUserConsentDescriptionOk returns a tuple with the UserConsentDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserConsentDescription

`func (o *MicrosoftGraphPermissionScope) SetUserConsentDescription(v string)`

SetUserConsentDescription sets UserConsentDescription field to given value.

### HasUserConsentDescription

`func (o *MicrosoftGraphPermissionScope) HasUserConsentDescription() bool`

HasUserConsentDescription returns a boolean if a field has been set.

### SetUserConsentDescriptionNil

`func (o *MicrosoftGraphPermissionScope) SetUserConsentDescriptionNil(b bool)`

 SetUserConsentDescriptionNil sets the value for UserConsentDescription to be an explicit nil

### UnsetUserConsentDescription
`func (o *MicrosoftGraphPermissionScope) UnsetUserConsentDescription()`

UnsetUserConsentDescription ensures that no value is present for UserConsentDescription, not even an explicit nil
### GetUserConsentDisplayName

`func (o *MicrosoftGraphPermissionScope) GetUserConsentDisplayName() string`

GetUserConsentDisplayName returns the UserConsentDisplayName field if non-nil, zero value otherwise.

### GetUserConsentDisplayNameOk

`func (o *MicrosoftGraphPermissionScope) GetUserConsentDisplayNameOk() (*string, bool)`

GetUserConsentDisplayNameOk returns a tuple with the UserConsentDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserConsentDisplayName

`func (o *MicrosoftGraphPermissionScope) SetUserConsentDisplayName(v string)`

SetUserConsentDisplayName sets UserConsentDisplayName field to given value.

### HasUserConsentDisplayName

`func (o *MicrosoftGraphPermissionScope) HasUserConsentDisplayName() bool`

HasUserConsentDisplayName returns a boolean if a field has been set.

### SetUserConsentDisplayNameNil

`func (o *MicrosoftGraphPermissionScope) SetUserConsentDisplayNameNil(b bool)`

 SetUserConsentDisplayNameNil sets the value for UserConsentDisplayName to be an explicit nil

### UnsetUserConsentDisplayName
`func (o *MicrosoftGraphPermissionScope) UnsetUserConsentDisplayName()`

UnsetUserConsentDisplayName ensures that no value is present for UserConsentDisplayName, not even an explicit nil
### GetValue

`func (o *MicrosoftGraphPermissionScope) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *MicrosoftGraphPermissionScope) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *MicrosoftGraphPermissionScope) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *MicrosoftGraphPermissionScope) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *MicrosoftGraphPermissionScope) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *MicrosoftGraphPermissionScope) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


