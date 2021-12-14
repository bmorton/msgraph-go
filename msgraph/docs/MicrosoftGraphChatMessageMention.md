# MicrosoftGraphChatMessageMention

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableInt32** | Index of an entity being mentioned in the specified chatMessage. Matches the {index} value in the corresponding &lt;at id&#x3D;&#39;{index}&#39;&gt; tag in the message body. | [optional] 
**Mentioned** | Pointer to [**AnyOfmicrosoftGraphChatMessageMentionedIdentitySet**](anyOf&lt;microsoft.graph.chatMessageMentionedIdentitySet&gt;.md) | The entity (user, application, team, or channel) that was @mentioned. | [optional] 
**MentionText** | Pointer to **NullableString** | String used to represent the mention. For example, a user&#39;s display name, a team name. | [optional] 

## Methods

### NewMicrosoftGraphChatMessageMention

`func NewMicrosoftGraphChatMessageMention() *MicrosoftGraphChatMessageMention`

NewMicrosoftGraphChatMessageMention instantiates a new MicrosoftGraphChatMessageMention object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphChatMessageMentionWithDefaults

`func NewMicrosoftGraphChatMessageMentionWithDefaults() *MicrosoftGraphChatMessageMention`

NewMicrosoftGraphChatMessageMentionWithDefaults instantiates a new MicrosoftGraphChatMessageMention object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphChatMessageMention) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphChatMessageMention) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphChatMessageMention) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphChatMessageMention) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *MicrosoftGraphChatMessageMention) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *MicrosoftGraphChatMessageMention) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetMentioned

`func (o *MicrosoftGraphChatMessageMention) GetMentioned() AnyOfmicrosoftGraphChatMessageMentionedIdentitySet`

GetMentioned returns the Mentioned field if non-nil, zero value otherwise.

### GetMentionedOk

`func (o *MicrosoftGraphChatMessageMention) GetMentionedOk() (*AnyOfmicrosoftGraphChatMessageMentionedIdentitySet, bool)`

GetMentionedOk returns a tuple with the Mentioned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMentioned

`func (o *MicrosoftGraphChatMessageMention) SetMentioned(v AnyOfmicrosoftGraphChatMessageMentionedIdentitySet)`

SetMentioned sets Mentioned field to given value.

### HasMentioned

`func (o *MicrosoftGraphChatMessageMention) HasMentioned() bool`

HasMentioned returns a boolean if a field has been set.

### SetMentionedNil

`func (o *MicrosoftGraphChatMessageMention) SetMentionedNil(b bool)`

 SetMentionedNil sets the value for Mentioned to be an explicit nil

### UnsetMentioned
`func (o *MicrosoftGraphChatMessageMention) UnsetMentioned()`

UnsetMentioned ensures that no value is present for Mentioned, not even an explicit nil
### GetMentionText

`func (o *MicrosoftGraphChatMessageMention) GetMentionText() string`

GetMentionText returns the MentionText field if non-nil, zero value otherwise.

### GetMentionTextOk

`func (o *MicrosoftGraphChatMessageMention) GetMentionTextOk() (*string, bool)`

GetMentionTextOk returns a tuple with the MentionText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMentionText

`func (o *MicrosoftGraphChatMessageMention) SetMentionText(v string)`

SetMentionText sets MentionText field to given value.

### HasMentionText

`func (o *MicrosoftGraphChatMessageMention) HasMentionText() bool`

HasMentionText returns a boolean if a field has been set.

### SetMentionTextNil

`func (o *MicrosoftGraphChatMessageMention) SetMentionTextNil(b bool)`

 SetMentionTextNil sets the value for MentionText to be an explicit nil

### UnsetMentionText
`func (o *MicrosoftGraphChatMessageMention) UnsetMentionText()`

UnsetMentionText ensures that no value is present for MentionText, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


