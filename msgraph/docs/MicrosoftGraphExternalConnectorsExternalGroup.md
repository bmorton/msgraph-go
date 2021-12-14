# MicrosoftGraphExternalConnectorsExternalGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**Description** | Pointer to **NullableString** | The description of the external group. Optional. | [optional] 
**DisplayName** | Pointer to **NullableString** | The friendly name of the external group. Optional. | [optional] 
**Members** | Pointer to [**[]MicrosoftGraphExternalConnectorsIdentity**](MicrosoftGraphExternalConnectorsIdentity.md) | A member added to an externalGroup. You can add Azure Active Directory users, Azure Active Directory groups, or an externalGroup as members. | [optional] 

## Methods

### NewMicrosoftGraphExternalConnectorsExternalGroup

`func NewMicrosoftGraphExternalConnectorsExternalGroup() *MicrosoftGraphExternalConnectorsExternalGroup`

NewMicrosoftGraphExternalConnectorsExternalGroup instantiates a new MicrosoftGraphExternalConnectorsExternalGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphExternalConnectorsExternalGroupWithDefaults

`func NewMicrosoftGraphExternalConnectorsExternalGroupWithDefaults() *MicrosoftGraphExternalConnectorsExternalGroup`

NewMicrosoftGraphExternalConnectorsExternalGroupWithDefaults instantiates a new MicrosoftGraphExternalConnectorsExternalGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphExternalConnectorsExternalGroup) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphExternalConnectorsExternalGroup) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphExternalConnectorsExternalGroup) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphExternalConnectorsExternalGroup) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDescription

`func (o *MicrosoftGraphExternalConnectorsExternalGroup) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MicrosoftGraphExternalConnectorsExternalGroup) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MicrosoftGraphExternalConnectorsExternalGroup) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MicrosoftGraphExternalConnectorsExternalGroup) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *MicrosoftGraphExternalConnectorsExternalGroup) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *MicrosoftGraphExternalConnectorsExternalGroup) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDisplayName

`func (o *MicrosoftGraphExternalConnectorsExternalGroup) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphExternalConnectorsExternalGroup) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *MicrosoftGraphExternalConnectorsExternalGroup) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *MicrosoftGraphExternalConnectorsExternalGroup) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *MicrosoftGraphExternalConnectorsExternalGroup) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *MicrosoftGraphExternalConnectorsExternalGroup) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetMembers

`func (o *MicrosoftGraphExternalConnectorsExternalGroup) GetMembers() []MicrosoftGraphExternalConnectorsIdentity`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *MicrosoftGraphExternalConnectorsExternalGroup) GetMembersOk() (*[]MicrosoftGraphExternalConnectorsIdentity, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *MicrosoftGraphExternalConnectorsExternalGroup) SetMembers(v []MicrosoftGraphExternalConnectorsIdentity)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *MicrosoftGraphExternalConnectorsExternalGroup) HasMembers() bool`

HasMembers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


