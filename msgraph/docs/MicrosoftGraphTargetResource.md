# MicrosoftGraphTargetResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **NullableString** | Indicates the visible name defined for the resource. Typically specified when the resource is created. | [optional] 
**GroupType** | Pointer to [**AnyOfmicrosoftGraphGroupType**](anyOf&lt;microsoft.graph.groupType&gt;.md) | When type is set to Group, this indicates the group type. Possible values are: unifiedGroups, azureAD, and unknownFutureValue | [optional] 
**Id** | Pointer to **NullableString** | Indicates the unique ID of the resource. | [optional] 
**ModifiedProperties** | Pointer to [**[]AnyOfmicrosoftGraphModifiedProperty**](AnyOfmicrosoftGraphModifiedProperty.md) | Indicates name, old value and new value of each attribute that changed. Property values depend on the operation type. | [optional] 
**Type** | Pointer to **NullableString** | Describes the resource type.  Example values include Application, Group, ServicePrincipal, and User. | [optional] 
**UserPrincipalName** | Pointer to **NullableString** | When type is set to User, this includes the user name that initiated the action; null for other types. | [optional] 

## Methods

### NewMicrosoftGraphTargetResource

`func NewMicrosoftGraphTargetResource() *MicrosoftGraphTargetResource`

NewMicrosoftGraphTargetResource instantiates a new MicrosoftGraphTargetResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphTargetResourceWithDefaults

`func NewMicrosoftGraphTargetResourceWithDefaults() *MicrosoftGraphTargetResource`

NewMicrosoftGraphTargetResourceWithDefaults instantiates a new MicrosoftGraphTargetResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *MicrosoftGraphTargetResource) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphTargetResource) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *MicrosoftGraphTargetResource) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *MicrosoftGraphTargetResource) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *MicrosoftGraphTargetResource) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *MicrosoftGraphTargetResource) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetGroupType

`func (o *MicrosoftGraphTargetResource) GetGroupType() AnyOfmicrosoftGraphGroupType`

GetGroupType returns the GroupType field if non-nil, zero value otherwise.

### GetGroupTypeOk

`func (o *MicrosoftGraphTargetResource) GetGroupTypeOk() (*AnyOfmicrosoftGraphGroupType, bool)`

GetGroupTypeOk returns a tuple with the GroupType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupType

`func (o *MicrosoftGraphTargetResource) SetGroupType(v AnyOfmicrosoftGraphGroupType)`

SetGroupType sets GroupType field to given value.

### HasGroupType

`func (o *MicrosoftGraphTargetResource) HasGroupType() bool`

HasGroupType returns a boolean if a field has been set.

### SetGroupTypeNil

`func (o *MicrosoftGraphTargetResource) SetGroupTypeNil(b bool)`

 SetGroupTypeNil sets the value for GroupType to be an explicit nil

### UnsetGroupType
`func (o *MicrosoftGraphTargetResource) UnsetGroupType()`

UnsetGroupType ensures that no value is present for GroupType, not even an explicit nil
### GetId

`func (o *MicrosoftGraphTargetResource) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphTargetResource) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphTargetResource) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphTargetResource) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *MicrosoftGraphTargetResource) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *MicrosoftGraphTargetResource) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetModifiedProperties

`func (o *MicrosoftGraphTargetResource) GetModifiedProperties() []*AnyOfmicrosoftGraphModifiedProperty`

GetModifiedProperties returns the ModifiedProperties field if non-nil, zero value otherwise.

### GetModifiedPropertiesOk

`func (o *MicrosoftGraphTargetResource) GetModifiedPropertiesOk() (*[]*AnyOfmicrosoftGraphModifiedProperty, bool)`

GetModifiedPropertiesOk returns a tuple with the ModifiedProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedProperties

`func (o *MicrosoftGraphTargetResource) SetModifiedProperties(v []*AnyOfmicrosoftGraphModifiedProperty)`

SetModifiedProperties sets ModifiedProperties field to given value.

### HasModifiedProperties

`func (o *MicrosoftGraphTargetResource) HasModifiedProperties() bool`

HasModifiedProperties returns a boolean if a field has been set.

### GetType

`func (o *MicrosoftGraphTargetResource) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MicrosoftGraphTargetResource) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MicrosoftGraphTargetResource) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *MicrosoftGraphTargetResource) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *MicrosoftGraphTargetResource) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *MicrosoftGraphTargetResource) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetUserPrincipalName

`func (o *MicrosoftGraphTargetResource) GetUserPrincipalName() string`

GetUserPrincipalName returns the UserPrincipalName field if non-nil, zero value otherwise.

### GetUserPrincipalNameOk

`func (o *MicrosoftGraphTargetResource) GetUserPrincipalNameOk() (*string, bool)`

GetUserPrincipalNameOk returns a tuple with the UserPrincipalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPrincipalName

`func (o *MicrosoftGraphTargetResource) SetUserPrincipalName(v string)`

SetUserPrincipalName sets UserPrincipalName field to given value.

### HasUserPrincipalName

`func (o *MicrosoftGraphTargetResource) HasUserPrincipalName() bool`

HasUserPrincipalName returns a boolean if a field has been set.

### SetUserPrincipalNameNil

`func (o *MicrosoftGraphTargetResource) SetUserPrincipalNameNil(b bool)`

 SetUserPrincipalNameNil sets the value for UserPrincipalName to be an explicit nil

### UnsetUserPrincipalName
`func (o *MicrosoftGraphTargetResource) UnsetUserPrincipalName()`

UnsetUserPrincipalName ensures that no value is present for UserPrincipalName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


