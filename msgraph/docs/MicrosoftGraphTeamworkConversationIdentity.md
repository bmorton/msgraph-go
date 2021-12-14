# MicrosoftGraphTeamworkConversationIdentity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **NullableString** | The identity&#39;s display name. Note that this may not always be available or up to date. For example, if a user changes their display name, the API may show the new value in a future response, but the items associated with the user won&#39;t show up as having changed when using delta. | [optional] 
**Id** | Pointer to **NullableString** | Unique identifier for the identity. | [optional] 
**ConversationIdentityType** | Pointer to [**AnyOfmicrosoftGraphTeamworkConversationIdentityType**](anyOf&lt;microsoft.graph.teamworkConversationIdentityType&gt;.md) | Type of conversation. Possible values are: team, channel, chat, and unknownFutureValue. | [optional] 

## Methods

### NewMicrosoftGraphTeamworkConversationIdentity

`func NewMicrosoftGraphTeamworkConversationIdentity() *MicrosoftGraphTeamworkConversationIdentity`

NewMicrosoftGraphTeamworkConversationIdentity instantiates a new MicrosoftGraphTeamworkConversationIdentity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphTeamworkConversationIdentityWithDefaults

`func NewMicrosoftGraphTeamworkConversationIdentityWithDefaults() *MicrosoftGraphTeamworkConversationIdentity`

NewMicrosoftGraphTeamworkConversationIdentityWithDefaults instantiates a new MicrosoftGraphTeamworkConversationIdentity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *MicrosoftGraphTeamworkConversationIdentity) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphTeamworkConversationIdentity) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *MicrosoftGraphTeamworkConversationIdentity) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *MicrosoftGraphTeamworkConversationIdentity) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *MicrosoftGraphTeamworkConversationIdentity) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *MicrosoftGraphTeamworkConversationIdentity) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetId

`func (o *MicrosoftGraphTeamworkConversationIdentity) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphTeamworkConversationIdentity) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphTeamworkConversationIdentity) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphTeamworkConversationIdentity) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *MicrosoftGraphTeamworkConversationIdentity) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *MicrosoftGraphTeamworkConversationIdentity) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetConversationIdentityType

`func (o *MicrosoftGraphTeamworkConversationIdentity) GetConversationIdentityType() AnyOfmicrosoftGraphTeamworkConversationIdentityType`

GetConversationIdentityType returns the ConversationIdentityType field if non-nil, zero value otherwise.

### GetConversationIdentityTypeOk

`func (o *MicrosoftGraphTeamworkConversationIdentity) GetConversationIdentityTypeOk() (*AnyOfmicrosoftGraphTeamworkConversationIdentityType, bool)`

GetConversationIdentityTypeOk returns a tuple with the ConversationIdentityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversationIdentityType

`func (o *MicrosoftGraphTeamworkConversationIdentity) SetConversationIdentityType(v AnyOfmicrosoftGraphTeamworkConversationIdentityType)`

SetConversationIdentityType sets ConversationIdentityType field to given value.

### HasConversationIdentityType

`func (o *MicrosoftGraphTeamworkConversationIdentity) HasConversationIdentityType() bool`

HasConversationIdentityType returns a boolean if a field has been set.

### SetConversationIdentityTypeNil

`func (o *MicrosoftGraphTeamworkConversationIdentity) SetConversationIdentityTypeNil(b bool)`

 SetConversationIdentityTypeNil sets the value for ConversationIdentityType to be an explicit nil

### UnsetConversationIdentityType
`func (o *MicrosoftGraphTeamworkConversationIdentity) UnsetConversationIdentityType()`

UnsetConversationIdentityType ensures that no value is present for ConversationIdentityType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


