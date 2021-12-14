# ExternalGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **NullableString** | The description of the external group. Optional. | [optional] 
**DisplayName** | Pointer to **NullableString** | The friendly name of the external group. Optional. | [optional] 
**Members** | Pointer to [**[]MicrosoftGraphExternalConnectorsIdentity**](MicrosoftGraphExternalConnectorsIdentity.md) | A member added to an externalGroup. You can add Azure Active Directory users, Azure Active Directory groups, or an externalGroup as members. | [optional] 

## Methods

### NewExternalGroup

`func NewExternalGroup() *ExternalGroup`

NewExternalGroup instantiates a new ExternalGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalGroupWithDefaults

`func NewExternalGroupWithDefaults() *ExternalGroup`

NewExternalGroupWithDefaults instantiates a new ExternalGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *ExternalGroup) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ExternalGroup) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ExternalGroup) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ExternalGroup) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ExternalGroup) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ExternalGroup) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDisplayName

`func (o *ExternalGroup) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ExternalGroup) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ExternalGroup) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ExternalGroup) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *ExternalGroup) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *ExternalGroup) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetMembers

`func (o *ExternalGroup) GetMembers() []MicrosoftGraphExternalConnectorsIdentity`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *ExternalGroup) GetMembersOk() (*[]MicrosoftGraphExternalConnectorsIdentity, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *ExternalGroup) SetMembers(v []MicrosoftGraphExternalConnectorsIdentity)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *ExternalGroup) HasMembers() bool`

HasMembers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


