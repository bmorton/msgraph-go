# MicrosoftGraphChatMessageMentionedIdentitySet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Application** | Pointer to [**AnyOfmicrosoftGraphIdentity**](anyOf&lt;microsoft.graph.identity&gt;.md) | Optional. The application associated with this action. | [optional] 
**Device** | Pointer to [**AnyOfmicrosoftGraphIdentity**](anyOf&lt;microsoft.graph.identity&gt;.md) | Optional. The device associated with this action. | [optional] 
**User** | Pointer to [**AnyOfmicrosoftGraphIdentity**](anyOf&lt;microsoft.graph.identity&gt;.md) | Optional. The user associated with this action. | [optional] 
**Conversation** | Pointer to [**AnyOfmicrosoftGraphTeamworkConversationIdentity**](anyOf&lt;microsoft.graph.teamworkConversationIdentity&gt;.md) | If present, represents a conversation (for example, team or channel) @mentioned in a message. | [optional] 

## Methods

### NewMicrosoftGraphChatMessageMentionedIdentitySet

`func NewMicrosoftGraphChatMessageMentionedIdentitySet() *MicrosoftGraphChatMessageMentionedIdentitySet`

NewMicrosoftGraphChatMessageMentionedIdentitySet instantiates a new MicrosoftGraphChatMessageMentionedIdentitySet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphChatMessageMentionedIdentitySetWithDefaults

`func NewMicrosoftGraphChatMessageMentionedIdentitySetWithDefaults() *MicrosoftGraphChatMessageMentionedIdentitySet`

NewMicrosoftGraphChatMessageMentionedIdentitySetWithDefaults instantiates a new MicrosoftGraphChatMessageMentionedIdentitySet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplication

`func (o *MicrosoftGraphChatMessageMentionedIdentitySet) GetApplication() AnyOfmicrosoftGraphIdentity`

GetApplication returns the Application field if non-nil, zero value otherwise.

### GetApplicationOk

`func (o *MicrosoftGraphChatMessageMentionedIdentitySet) GetApplicationOk() (*AnyOfmicrosoftGraphIdentity, bool)`

GetApplicationOk returns a tuple with the Application field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplication

`func (o *MicrosoftGraphChatMessageMentionedIdentitySet) SetApplication(v AnyOfmicrosoftGraphIdentity)`

SetApplication sets Application field to given value.

### HasApplication

`func (o *MicrosoftGraphChatMessageMentionedIdentitySet) HasApplication() bool`

HasApplication returns a boolean if a field has been set.

### SetApplicationNil

`func (o *MicrosoftGraphChatMessageMentionedIdentitySet) SetApplicationNil(b bool)`

 SetApplicationNil sets the value for Application to be an explicit nil

### UnsetApplication
`func (o *MicrosoftGraphChatMessageMentionedIdentitySet) UnsetApplication()`

UnsetApplication ensures that no value is present for Application, not even an explicit nil
### GetDevice

`func (o *MicrosoftGraphChatMessageMentionedIdentitySet) GetDevice() AnyOfmicrosoftGraphIdentity`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *MicrosoftGraphChatMessageMentionedIdentitySet) GetDeviceOk() (*AnyOfmicrosoftGraphIdentity, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *MicrosoftGraphChatMessageMentionedIdentitySet) SetDevice(v AnyOfmicrosoftGraphIdentity)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *MicrosoftGraphChatMessageMentionedIdentitySet) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### SetDeviceNil

`func (o *MicrosoftGraphChatMessageMentionedIdentitySet) SetDeviceNil(b bool)`

 SetDeviceNil sets the value for Device to be an explicit nil

### UnsetDevice
`func (o *MicrosoftGraphChatMessageMentionedIdentitySet) UnsetDevice()`

UnsetDevice ensures that no value is present for Device, not even an explicit nil
### GetUser

`func (o *MicrosoftGraphChatMessageMentionedIdentitySet) GetUser() AnyOfmicrosoftGraphIdentity`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *MicrosoftGraphChatMessageMentionedIdentitySet) GetUserOk() (*AnyOfmicrosoftGraphIdentity, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *MicrosoftGraphChatMessageMentionedIdentitySet) SetUser(v AnyOfmicrosoftGraphIdentity)`

SetUser sets User field to given value.

### HasUser

`func (o *MicrosoftGraphChatMessageMentionedIdentitySet) HasUser() bool`

HasUser returns a boolean if a field has been set.

### SetUserNil

`func (o *MicrosoftGraphChatMessageMentionedIdentitySet) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *MicrosoftGraphChatMessageMentionedIdentitySet) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil
### GetConversation

`func (o *MicrosoftGraphChatMessageMentionedIdentitySet) GetConversation() AnyOfmicrosoftGraphTeamworkConversationIdentity`

GetConversation returns the Conversation field if non-nil, zero value otherwise.

### GetConversationOk

`func (o *MicrosoftGraphChatMessageMentionedIdentitySet) GetConversationOk() (*AnyOfmicrosoftGraphTeamworkConversationIdentity, bool)`

GetConversationOk returns a tuple with the Conversation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversation

`func (o *MicrosoftGraphChatMessageMentionedIdentitySet) SetConversation(v AnyOfmicrosoftGraphTeamworkConversationIdentity)`

SetConversation sets Conversation field to given value.

### HasConversation

`func (o *MicrosoftGraphChatMessageMentionedIdentitySet) HasConversation() bool`

HasConversation returns a boolean if a field has been set.

### SetConversationNil

`func (o *MicrosoftGraphChatMessageMentionedIdentitySet) SetConversationNil(b bool)`

 SetConversationNil sets the value for Conversation to be an explicit nil

### UnsetConversation
`func (o *MicrosoftGraphChatMessageMentionedIdentitySet) UnsetConversation()`

UnsetConversation ensures that no value is present for Conversation, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


