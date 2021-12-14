# TeamworkConversationIdentity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConversationIdentityType** | Pointer to [**AnyOfmicrosoftGraphTeamworkConversationIdentityType**](anyOf&lt;microsoft.graph.teamworkConversationIdentityType&gt;.md) | Type of conversation. Possible values are: team, channel, chat, and unknownFutureValue. | [optional] 

## Methods

### NewTeamworkConversationIdentity

`func NewTeamworkConversationIdentity() *TeamworkConversationIdentity`

NewTeamworkConversationIdentity instantiates a new TeamworkConversationIdentity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamworkConversationIdentityWithDefaults

`func NewTeamworkConversationIdentityWithDefaults() *TeamworkConversationIdentity`

NewTeamworkConversationIdentityWithDefaults instantiates a new TeamworkConversationIdentity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConversationIdentityType

`func (o *TeamworkConversationIdentity) GetConversationIdentityType() AnyOfmicrosoftGraphTeamworkConversationIdentityType`

GetConversationIdentityType returns the ConversationIdentityType field if non-nil, zero value otherwise.

### GetConversationIdentityTypeOk

`func (o *TeamworkConversationIdentity) GetConversationIdentityTypeOk() (*AnyOfmicrosoftGraphTeamworkConversationIdentityType, bool)`

GetConversationIdentityTypeOk returns a tuple with the ConversationIdentityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversationIdentityType

`func (o *TeamworkConversationIdentity) SetConversationIdentityType(v AnyOfmicrosoftGraphTeamworkConversationIdentityType)`

SetConversationIdentityType sets ConversationIdentityType field to given value.

### HasConversationIdentityType

`func (o *TeamworkConversationIdentity) HasConversationIdentityType() bool`

HasConversationIdentityType returns a boolean if a field has been set.

### SetConversationIdentityTypeNil

`func (o *TeamworkConversationIdentity) SetConversationIdentityTypeNil(b bool)`

 SetConversationIdentityTypeNil sets the value for ConversationIdentityType to be an explicit nil

### UnsetConversationIdentityType
`func (o *TeamworkConversationIdentity) UnsetConversationIdentityType()`

UnsetConversationIdentityType ensures that no value is present for ConversationIdentityType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


