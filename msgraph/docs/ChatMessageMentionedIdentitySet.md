# ChatMessageMentionedIdentitySet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Conversation** | Pointer to [**AnyOfmicrosoftGraphTeamworkConversationIdentity**](anyOf&lt;microsoft.graph.teamworkConversationIdentity&gt;.md) | If present, represents a conversation (for example, team or channel) @mentioned in a message. | [optional] 

## Methods

### NewChatMessageMentionedIdentitySet

`func NewChatMessageMentionedIdentitySet() *ChatMessageMentionedIdentitySet`

NewChatMessageMentionedIdentitySet instantiates a new ChatMessageMentionedIdentitySet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatMessageMentionedIdentitySetWithDefaults

`func NewChatMessageMentionedIdentitySetWithDefaults() *ChatMessageMentionedIdentitySet`

NewChatMessageMentionedIdentitySetWithDefaults instantiates a new ChatMessageMentionedIdentitySet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConversation

`func (o *ChatMessageMentionedIdentitySet) GetConversation() AnyOfmicrosoftGraphTeamworkConversationIdentity`

GetConversation returns the Conversation field if non-nil, zero value otherwise.

### GetConversationOk

`func (o *ChatMessageMentionedIdentitySet) GetConversationOk() (*AnyOfmicrosoftGraphTeamworkConversationIdentity, bool)`

GetConversationOk returns a tuple with the Conversation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversation

`func (o *ChatMessageMentionedIdentitySet) SetConversation(v AnyOfmicrosoftGraphTeamworkConversationIdentity)`

SetConversation sets Conversation field to given value.

### HasConversation

`func (o *ChatMessageMentionedIdentitySet) HasConversation() bool`

HasConversation returns a boolean if a field has been set.

### SetConversationNil

`func (o *ChatMessageMentionedIdentitySet) SetConversationNil(b bool)`

 SetConversationNil sets the value for Conversation to be an explicit nil

### UnsetConversation
`func (o *ChatMessageMentionedIdentitySet) UnsetConversation()`

UnsetConversation ensures that no value is present for Conversation, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


