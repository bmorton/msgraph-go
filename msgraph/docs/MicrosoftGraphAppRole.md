# MicrosoftGraphAppRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedMemberTypes** | Pointer to **[]string** | Specifies whether this app role can be assigned to users and groups (by setting to [&#39;User&#39;]), to other application&#39;s (by setting to [&#39;Application&#39;], or both (by setting to [&#39;User&#39;, &#39;Application&#39;]). App roles supporting assignment to other applications&#39; service principals are also known as application permissions. The &#39;Application&#39; value is only supported for app roles defined on application entities. | [optional] 
**Description** | Pointer to **NullableString** | The description for the app role. This is displayed when the app role is being assigned and, if the app role functions as an application permission, during  consent experiences. | [optional] 
**DisplayName** | Pointer to **NullableString** | Display name for the permission that appears in the app role assignment and consent experiences. | [optional] 
**Id** | Pointer to **string** | Unique role identifier inside the appRoles collection. When creating a new app role, a new Guid identifier must be provided. | [optional] 
**IsEnabled** | Pointer to **bool** | When creating or updating an app role, this must be set to true (which is the default). To delete a role, this must first be set to false.  At that point, in a subsequent call, this role may be removed. | [optional] 
**Origin** | Pointer to **NullableString** | Specifies if the app role is defined on the application object or on the servicePrincipal entity. Must not be included in any POST or PATCH requests. Read-only. | [optional] 
**Value** | Pointer to **NullableString** | Specifies the value to include in the roles claim in ID tokens and access tokens authenticating an assigned user or service principal. Must not exceed 120 characters in length. Allowed characters are : ! # $ % &amp; &#39; ( ) * + , - . / : ;  &#x3D;  ? @ [ ] ^ + _  {  } ~, as well as characters in the ranges 0-9, A-Z and a-z. Any other character, including the space character, are not allowed. May not begin with .. | [optional] 

## Methods

### NewMicrosoftGraphAppRole

`func NewMicrosoftGraphAppRole() *MicrosoftGraphAppRole`

NewMicrosoftGraphAppRole instantiates a new MicrosoftGraphAppRole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphAppRoleWithDefaults

`func NewMicrosoftGraphAppRoleWithDefaults() *MicrosoftGraphAppRole`

NewMicrosoftGraphAppRoleWithDefaults instantiates a new MicrosoftGraphAppRole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedMemberTypes

`func (o *MicrosoftGraphAppRole) GetAllowedMemberTypes() []string`

GetAllowedMemberTypes returns the AllowedMemberTypes field if non-nil, zero value otherwise.

### GetAllowedMemberTypesOk

`func (o *MicrosoftGraphAppRole) GetAllowedMemberTypesOk() (*[]string, bool)`

GetAllowedMemberTypesOk returns a tuple with the AllowedMemberTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedMemberTypes

`func (o *MicrosoftGraphAppRole) SetAllowedMemberTypes(v []string)`

SetAllowedMemberTypes sets AllowedMemberTypes field to given value.

### HasAllowedMemberTypes

`func (o *MicrosoftGraphAppRole) HasAllowedMemberTypes() bool`

HasAllowedMemberTypes returns a boolean if a field has been set.

### GetDescription

`func (o *MicrosoftGraphAppRole) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MicrosoftGraphAppRole) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MicrosoftGraphAppRole) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MicrosoftGraphAppRole) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *MicrosoftGraphAppRole) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *MicrosoftGraphAppRole) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDisplayName

`func (o *MicrosoftGraphAppRole) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphAppRole) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *MicrosoftGraphAppRole) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *MicrosoftGraphAppRole) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *MicrosoftGraphAppRole) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *MicrosoftGraphAppRole) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetId

`func (o *MicrosoftGraphAppRole) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphAppRole) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphAppRole) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphAppRole) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsEnabled

`func (o *MicrosoftGraphAppRole) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *MicrosoftGraphAppRole) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *MicrosoftGraphAppRole) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *MicrosoftGraphAppRole) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetOrigin

`func (o *MicrosoftGraphAppRole) GetOrigin() string`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *MicrosoftGraphAppRole) GetOriginOk() (*string, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *MicrosoftGraphAppRole) SetOrigin(v string)`

SetOrigin sets Origin field to given value.

### HasOrigin

`func (o *MicrosoftGraphAppRole) HasOrigin() bool`

HasOrigin returns a boolean if a field has been set.

### SetOriginNil

`func (o *MicrosoftGraphAppRole) SetOriginNil(b bool)`

 SetOriginNil sets the value for Origin to be an explicit nil

### UnsetOrigin
`func (o *MicrosoftGraphAppRole) UnsetOrigin()`

UnsetOrigin ensures that no value is present for Origin, not even an explicit nil
### GetValue

`func (o *MicrosoftGraphAppRole) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *MicrosoftGraphAppRole) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *MicrosoftGraphAppRole) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *MicrosoftGraphAppRole) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *MicrosoftGraphAppRole) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *MicrosoftGraphAppRole) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


