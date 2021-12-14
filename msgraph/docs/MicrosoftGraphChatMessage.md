# MicrosoftGraphChatMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**Attachments** | Pointer to [**[]AnyOfmicrosoftGraphChatMessageAttachment**](AnyOfmicrosoftGraphChatMessageAttachment.md) | References to attached objects like files, tabs, meetings etc. | [optional] 
**Body** | Pointer to [**MicrosoftGraphItemBody**](MicrosoftGraphItemBody.md) |  | [optional] 
**ChannelIdentity** | Pointer to [**AnyOfmicrosoftGraphChannelIdentity**](anyOf&lt;microsoft.graph.channelIdentity&gt;.md) | If the message was sent in a channel, represents identity of the channel. | [optional] 
**ChatId** | Pointer to **NullableString** | If the message was sent in a chat, represents the identity of the chat. | [optional] 
**CreatedDateTime** | Pointer to **NullableTime** | Timestamp of when the chat message was created. | [optional] 
**DeletedDateTime** | Pointer to **NullableTime** | Read only. Timestamp at which the chat message was deleted, or null if not deleted. | [optional] 
**Etag** | Pointer to **NullableString** | Read-only. Version number of the chat message. | [optional] 
**EventDetail** | Pointer to [**AnyOfobject**](anyOf&lt;object&gt;.md) | Read-only. If present, represents details of an event that happened in a chat, a channel, or a team, for example, adding new members. For event messages, the messageType property will be set to systemEventMessage. | [optional] 
**From** | Pointer to [**AnyOfmicrosoftGraphChatMessageFromIdentitySet**](anyOf&lt;microsoft.graph.chatMessageFromIdentitySet&gt;.md) | Details of the sender of the chat message. Can only be set during migration. | [optional] 
**Importance** | Pointer to [**AnyOfmicrosoftGraphChatMessageImportance**](anyOf&lt;microsoft.graph.chatMessageImportance&gt;.md) | The importance of the chat message. The possible values are: normal, high, urgent. | [optional] 
**LastEditedDateTime** | Pointer to **NullableTime** | Read only. Timestamp when edits to the chat message were made. Triggers an &#39;Edited&#39; flag in the Teams UI. If no edits are made the value is null. | [optional] 
**LastModifiedDateTime** | Pointer to **NullableTime** | Read only. Timestamp when the chat message is created (initial setting) or modified, including when a reaction is added or removed. | [optional] 
**Locale** | Pointer to **string** | Locale of the chat message set by the client. Always set to en-us. | [optional] 
**Mentions** | Pointer to [**[]AnyOfmicrosoftGraphChatMessageMention**](AnyOfmicrosoftGraphChatMessageMention.md) | List of entities mentioned in the chat message. Supported entities are: user, bot, team, and channel. | [optional] 
**MessageType** | Pointer to [**AnyOfmicrosoftGraphChatMessageType**](anyOf&lt;microsoft.graph.chatMessageType&gt;.md) | The type of chat message. The possible values are: message, chatEvent, typing, unknownFutureValue, systemEventMessage. Note that you must use the Prefer: include-unknown-enum-members request header to get the following value in this evolvable enum: systemEventMessage. | [optional] 
**PolicyViolation** | Pointer to [**AnyOfmicrosoftGraphChatMessagePolicyViolation**](anyOf&lt;microsoft.graph.chatMessagePolicyViolation&gt;.md) | Defines the properties of a policy violation set by a data loss prevention (DLP) application. | [optional] 
**Reactions** | Pointer to [**[]AnyOfmicrosoftGraphChatMessageReaction**](AnyOfmicrosoftGraphChatMessageReaction.md) | Reactions for this chat message (for example, Like). | [optional] 
**ReplyToId** | Pointer to **NullableString** | Read-only. ID of the parent chat message or root chat message of the thread. (Only applies to chat messages in channels, not chats.) | [optional] 
**Subject** | Pointer to **NullableString** | The subject of the chat message, in plaintext. | [optional] 
**Summary** | Pointer to **NullableString** | Summary text of the chat message that could be used for push notifications and summary views or fall back views. Only applies to channel chat messages, not chat messages in a chat. | [optional] 
**WebUrl** | Pointer to **NullableString** | Read-only. Link to the message in Microsoft Teams. | [optional] 
**HostedContents** | Pointer to [**[]MicrosoftGraphChatMessageHostedContent**](MicrosoftGraphChatMessageHostedContent.md) | Content in a message hosted by Microsoft Teams - for example, images or code snippets. | [optional] 
**Replies** | Pointer to [**[]MicrosoftGraphChatMessage**](MicrosoftGraphChatMessage.md) | Replies for a specified message. | [optional] 

## Methods

### NewMicrosoftGraphChatMessage

`func NewMicrosoftGraphChatMessage() *MicrosoftGraphChatMessage`

NewMicrosoftGraphChatMessage instantiates a new MicrosoftGraphChatMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphChatMessageWithDefaults

`func NewMicrosoftGraphChatMessageWithDefaults() *MicrosoftGraphChatMessage`

NewMicrosoftGraphChatMessageWithDefaults instantiates a new MicrosoftGraphChatMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphChatMessage) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphChatMessage) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphChatMessage) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphChatMessage) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAttachments

`func (o *MicrosoftGraphChatMessage) GetAttachments() []*AnyOfmicrosoftGraphChatMessageAttachment`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *MicrosoftGraphChatMessage) GetAttachmentsOk() (*[]*AnyOfmicrosoftGraphChatMessageAttachment, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *MicrosoftGraphChatMessage) SetAttachments(v []*AnyOfmicrosoftGraphChatMessageAttachment)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *MicrosoftGraphChatMessage) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetBody

`func (o *MicrosoftGraphChatMessage) GetBody() MicrosoftGraphItemBody`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *MicrosoftGraphChatMessage) GetBodyOk() (*MicrosoftGraphItemBody, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *MicrosoftGraphChatMessage) SetBody(v MicrosoftGraphItemBody)`

SetBody sets Body field to given value.

### HasBody

`func (o *MicrosoftGraphChatMessage) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetChannelIdentity

`func (o *MicrosoftGraphChatMessage) GetChannelIdentity() AnyOfmicrosoftGraphChannelIdentity`

GetChannelIdentity returns the ChannelIdentity field if non-nil, zero value otherwise.

### GetChannelIdentityOk

`func (o *MicrosoftGraphChatMessage) GetChannelIdentityOk() (*AnyOfmicrosoftGraphChannelIdentity, bool)`

GetChannelIdentityOk returns a tuple with the ChannelIdentity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelIdentity

`func (o *MicrosoftGraphChatMessage) SetChannelIdentity(v AnyOfmicrosoftGraphChannelIdentity)`

SetChannelIdentity sets ChannelIdentity field to given value.

### HasChannelIdentity

`func (o *MicrosoftGraphChatMessage) HasChannelIdentity() bool`

HasChannelIdentity returns a boolean if a field has been set.

### SetChannelIdentityNil

`func (o *MicrosoftGraphChatMessage) SetChannelIdentityNil(b bool)`

 SetChannelIdentityNil sets the value for ChannelIdentity to be an explicit nil

### UnsetChannelIdentity
`func (o *MicrosoftGraphChatMessage) UnsetChannelIdentity()`

UnsetChannelIdentity ensures that no value is present for ChannelIdentity, not even an explicit nil
### GetChatId

`func (o *MicrosoftGraphChatMessage) GetChatId() string`

GetChatId returns the ChatId field if non-nil, zero value otherwise.

### GetChatIdOk

`func (o *MicrosoftGraphChatMessage) GetChatIdOk() (*string, bool)`

GetChatIdOk returns a tuple with the ChatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatId

`func (o *MicrosoftGraphChatMessage) SetChatId(v string)`

SetChatId sets ChatId field to given value.

### HasChatId

`func (o *MicrosoftGraphChatMessage) HasChatId() bool`

HasChatId returns a boolean if a field has been set.

### SetChatIdNil

`func (o *MicrosoftGraphChatMessage) SetChatIdNil(b bool)`

 SetChatIdNil sets the value for ChatId to be an explicit nil

### UnsetChatId
`func (o *MicrosoftGraphChatMessage) UnsetChatId()`

UnsetChatId ensures that no value is present for ChatId, not even an explicit nil
### GetCreatedDateTime

`func (o *MicrosoftGraphChatMessage) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *MicrosoftGraphChatMessage) GetCreatedDateTimeOk() (*time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateTime

`func (o *MicrosoftGraphChatMessage) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime sets CreatedDateTime field to given value.

### HasCreatedDateTime

`func (o *MicrosoftGraphChatMessage) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTimeNil

`func (o *MicrosoftGraphChatMessage) SetCreatedDateTimeNil(b bool)`

 SetCreatedDateTimeNil sets the value for CreatedDateTime to be an explicit nil

### UnsetCreatedDateTime
`func (o *MicrosoftGraphChatMessage) UnsetCreatedDateTime()`

UnsetCreatedDateTime ensures that no value is present for CreatedDateTime, not even an explicit nil
### GetDeletedDateTime

`func (o *MicrosoftGraphChatMessage) GetDeletedDateTime() time.Time`

GetDeletedDateTime returns the DeletedDateTime field if non-nil, zero value otherwise.

### GetDeletedDateTimeOk

`func (o *MicrosoftGraphChatMessage) GetDeletedDateTimeOk() (*time.Time, bool)`

GetDeletedDateTimeOk returns a tuple with the DeletedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedDateTime

`func (o *MicrosoftGraphChatMessage) SetDeletedDateTime(v time.Time)`

SetDeletedDateTime sets DeletedDateTime field to given value.

### HasDeletedDateTime

`func (o *MicrosoftGraphChatMessage) HasDeletedDateTime() bool`

HasDeletedDateTime returns a boolean if a field has been set.

### SetDeletedDateTimeNil

`func (o *MicrosoftGraphChatMessage) SetDeletedDateTimeNil(b bool)`

 SetDeletedDateTimeNil sets the value for DeletedDateTime to be an explicit nil

### UnsetDeletedDateTime
`func (o *MicrosoftGraphChatMessage) UnsetDeletedDateTime()`

UnsetDeletedDateTime ensures that no value is present for DeletedDateTime, not even an explicit nil
### GetEtag

`func (o *MicrosoftGraphChatMessage) GetEtag() string`

GetEtag returns the Etag field if non-nil, zero value otherwise.

### GetEtagOk

`func (o *MicrosoftGraphChatMessage) GetEtagOk() (*string, bool)`

GetEtagOk returns a tuple with the Etag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEtag

`func (o *MicrosoftGraphChatMessage) SetEtag(v string)`

SetEtag sets Etag field to given value.

### HasEtag

`func (o *MicrosoftGraphChatMessage) HasEtag() bool`

HasEtag returns a boolean if a field has been set.

### SetEtagNil

`func (o *MicrosoftGraphChatMessage) SetEtagNil(b bool)`

 SetEtagNil sets the value for Etag to be an explicit nil

### UnsetEtag
`func (o *MicrosoftGraphChatMessage) UnsetEtag()`

UnsetEtag ensures that no value is present for Etag, not even an explicit nil
### GetEventDetail

`func (o *MicrosoftGraphChatMessage) GetEventDetail() AnyOfobject`

GetEventDetail returns the EventDetail field if non-nil, zero value otherwise.

### GetEventDetailOk

`func (o *MicrosoftGraphChatMessage) GetEventDetailOk() (*AnyOfobject, bool)`

GetEventDetailOk returns a tuple with the EventDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventDetail

`func (o *MicrosoftGraphChatMessage) SetEventDetail(v AnyOfobject)`

SetEventDetail sets EventDetail field to given value.

### HasEventDetail

`func (o *MicrosoftGraphChatMessage) HasEventDetail() bool`

HasEventDetail returns a boolean if a field has been set.

### SetEventDetailNil

`func (o *MicrosoftGraphChatMessage) SetEventDetailNil(b bool)`

 SetEventDetailNil sets the value for EventDetail to be an explicit nil

### UnsetEventDetail
`func (o *MicrosoftGraphChatMessage) UnsetEventDetail()`

UnsetEventDetail ensures that no value is present for EventDetail, not even an explicit nil
### GetFrom

`func (o *MicrosoftGraphChatMessage) GetFrom() AnyOfmicrosoftGraphChatMessageFromIdentitySet`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *MicrosoftGraphChatMessage) GetFromOk() (*AnyOfmicrosoftGraphChatMessageFromIdentitySet, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *MicrosoftGraphChatMessage) SetFrom(v AnyOfmicrosoftGraphChatMessageFromIdentitySet)`

SetFrom sets From field to given value.

### HasFrom

`func (o *MicrosoftGraphChatMessage) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### SetFromNil

`func (o *MicrosoftGraphChatMessage) SetFromNil(b bool)`

 SetFromNil sets the value for From to be an explicit nil

### UnsetFrom
`func (o *MicrosoftGraphChatMessage) UnsetFrom()`

UnsetFrom ensures that no value is present for From, not even an explicit nil
### GetImportance

`func (o *MicrosoftGraphChatMessage) GetImportance() AnyOfmicrosoftGraphChatMessageImportance`

GetImportance returns the Importance field if non-nil, zero value otherwise.

### GetImportanceOk

`func (o *MicrosoftGraphChatMessage) GetImportanceOk() (*AnyOfmicrosoftGraphChatMessageImportance, bool)`

GetImportanceOk returns a tuple with the Importance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportance

`func (o *MicrosoftGraphChatMessage) SetImportance(v AnyOfmicrosoftGraphChatMessageImportance)`

SetImportance sets Importance field to given value.

### HasImportance

`func (o *MicrosoftGraphChatMessage) HasImportance() bool`

HasImportance returns a boolean if a field has been set.

### SetImportanceNil

`func (o *MicrosoftGraphChatMessage) SetImportanceNil(b bool)`

 SetImportanceNil sets the value for Importance to be an explicit nil

### UnsetImportance
`func (o *MicrosoftGraphChatMessage) UnsetImportance()`

UnsetImportance ensures that no value is present for Importance, not even an explicit nil
### GetLastEditedDateTime

`func (o *MicrosoftGraphChatMessage) GetLastEditedDateTime() time.Time`

GetLastEditedDateTime returns the LastEditedDateTime field if non-nil, zero value otherwise.

### GetLastEditedDateTimeOk

`func (o *MicrosoftGraphChatMessage) GetLastEditedDateTimeOk() (*time.Time, bool)`

GetLastEditedDateTimeOk returns a tuple with the LastEditedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEditedDateTime

`func (o *MicrosoftGraphChatMessage) SetLastEditedDateTime(v time.Time)`

SetLastEditedDateTime sets LastEditedDateTime field to given value.

### HasLastEditedDateTime

`func (o *MicrosoftGraphChatMessage) HasLastEditedDateTime() bool`

HasLastEditedDateTime returns a boolean if a field has been set.

### SetLastEditedDateTimeNil

`func (o *MicrosoftGraphChatMessage) SetLastEditedDateTimeNil(b bool)`

 SetLastEditedDateTimeNil sets the value for LastEditedDateTime to be an explicit nil

### UnsetLastEditedDateTime
`func (o *MicrosoftGraphChatMessage) UnsetLastEditedDateTime()`

UnsetLastEditedDateTime ensures that no value is present for LastEditedDateTime, not even an explicit nil
### GetLastModifiedDateTime

`func (o *MicrosoftGraphChatMessage) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *MicrosoftGraphChatMessage) GetLastModifiedDateTimeOk() (*time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedDateTime

`func (o *MicrosoftGraphChatMessage) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime sets LastModifiedDateTime field to given value.

### HasLastModifiedDateTime

`func (o *MicrosoftGraphChatMessage) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### SetLastModifiedDateTimeNil

`func (o *MicrosoftGraphChatMessage) SetLastModifiedDateTimeNil(b bool)`

 SetLastModifiedDateTimeNil sets the value for LastModifiedDateTime to be an explicit nil

### UnsetLastModifiedDateTime
`func (o *MicrosoftGraphChatMessage) UnsetLastModifiedDateTime()`

UnsetLastModifiedDateTime ensures that no value is present for LastModifiedDateTime, not even an explicit nil
### GetLocale

`func (o *MicrosoftGraphChatMessage) GetLocale() string`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *MicrosoftGraphChatMessage) GetLocaleOk() (*string, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *MicrosoftGraphChatMessage) SetLocale(v string)`

SetLocale sets Locale field to given value.

### HasLocale

`func (o *MicrosoftGraphChatMessage) HasLocale() bool`

HasLocale returns a boolean if a field has been set.

### GetMentions

`func (o *MicrosoftGraphChatMessage) GetMentions() []*AnyOfmicrosoftGraphChatMessageMention`

GetMentions returns the Mentions field if non-nil, zero value otherwise.

### GetMentionsOk

`func (o *MicrosoftGraphChatMessage) GetMentionsOk() (*[]*AnyOfmicrosoftGraphChatMessageMention, bool)`

GetMentionsOk returns a tuple with the Mentions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMentions

`func (o *MicrosoftGraphChatMessage) SetMentions(v []*AnyOfmicrosoftGraphChatMessageMention)`

SetMentions sets Mentions field to given value.

### HasMentions

`func (o *MicrosoftGraphChatMessage) HasMentions() bool`

HasMentions returns a boolean if a field has been set.

### GetMessageType

`func (o *MicrosoftGraphChatMessage) GetMessageType() AnyOfmicrosoftGraphChatMessageType`

GetMessageType returns the MessageType field if non-nil, zero value otherwise.

### GetMessageTypeOk

`func (o *MicrosoftGraphChatMessage) GetMessageTypeOk() (*AnyOfmicrosoftGraphChatMessageType, bool)`

GetMessageTypeOk returns a tuple with the MessageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageType

`func (o *MicrosoftGraphChatMessage) SetMessageType(v AnyOfmicrosoftGraphChatMessageType)`

SetMessageType sets MessageType field to given value.

### HasMessageType

`func (o *MicrosoftGraphChatMessage) HasMessageType() bool`

HasMessageType returns a boolean if a field has been set.

### SetMessageTypeNil

`func (o *MicrosoftGraphChatMessage) SetMessageTypeNil(b bool)`

 SetMessageTypeNil sets the value for MessageType to be an explicit nil

### UnsetMessageType
`func (o *MicrosoftGraphChatMessage) UnsetMessageType()`

UnsetMessageType ensures that no value is present for MessageType, not even an explicit nil
### GetPolicyViolation

`func (o *MicrosoftGraphChatMessage) GetPolicyViolation() AnyOfmicrosoftGraphChatMessagePolicyViolation`

GetPolicyViolation returns the PolicyViolation field if non-nil, zero value otherwise.

### GetPolicyViolationOk

`func (o *MicrosoftGraphChatMessage) GetPolicyViolationOk() (*AnyOfmicrosoftGraphChatMessagePolicyViolation, bool)`

GetPolicyViolationOk returns a tuple with the PolicyViolation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyViolation

`func (o *MicrosoftGraphChatMessage) SetPolicyViolation(v AnyOfmicrosoftGraphChatMessagePolicyViolation)`

SetPolicyViolation sets PolicyViolation field to given value.

### HasPolicyViolation

`func (o *MicrosoftGraphChatMessage) HasPolicyViolation() bool`

HasPolicyViolation returns a boolean if a field has been set.

### SetPolicyViolationNil

`func (o *MicrosoftGraphChatMessage) SetPolicyViolationNil(b bool)`

 SetPolicyViolationNil sets the value for PolicyViolation to be an explicit nil

### UnsetPolicyViolation
`func (o *MicrosoftGraphChatMessage) UnsetPolicyViolation()`

UnsetPolicyViolation ensures that no value is present for PolicyViolation, not even an explicit nil
### GetReactions

`func (o *MicrosoftGraphChatMessage) GetReactions() []*AnyOfmicrosoftGraphChatMessageReaction`

GetReactions returns the Reactions field if non-nil, zero value otherwise.

### GetReactionsOk

`func (o *MicrosoftGraphChatMessage) GetReactionsOk() (*[]*AnyOfmicrosoftGraphChatMessageReaction, bool)`

GetReactionsOk returns a tuple with the Reactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReactions

`func (o *MicrosoftGraphChatMessage) SetReactions(v []*AnyOfmicrosoftGraphChatMessageReaction)`

SetReactions sets Reactions field to given value.

### HasReactions

`func (o *MicrosoftGraphChatMessage) HasReactions() bool`

HasReactions returns a boolean if a field has been set.

### GetReplyToId

`func (o *MicrosoftGraphChatMessage) GetReplyToId() string`

GetReplyToId returns the ReplyToId field if non-nil, zero value otherwise.

### GetReplyToIdOk

`func (o *MicrosoftGraphChatMessage) GetReplyToIdOk() (*string, bool)`

GetReplyToIdOk returns a tuple with the ReplyToId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyToId

`func (o *MicrosoftGraphChatMessage) SetReplyToId(v string)`

SetReplyToId sets ReplyToId field to given value.

### HasReplyToId

`func (o *MicrosoftGraphChatMessage) HasReplyToId() bool`

HasReplyToId returns a boolean if a field has been set.

### SetReplyToIdNil

`func (o *MicrosoftGraphChatMessage) SetReplyToIdNil(b bool)`

 SetReplyToIdNil sets the value for ReplyToId to be an explicit nil

### UnsetReplyToId
`func (o *MicrosoftGraphChatMessage) UnsetReplyToId()`

UnsetReplyToId ensures that no value is present for ReplyToId, not even an explicit nil
### GetSubject

`func (o *MicrosoftGraphChatMessage) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *MicrosoftGraphChatMessage) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *MicrosoftGraphChatMessage) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *MicrosoftGraphChatMessage) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### SetSubjectNil

`func (o *MicrosoftGraphChatMessage) SetSubjectNil(b bool)`

 SetSubjectNil sets the value for Subject to be an explicit nil

### UnsetSubject
`func (o *MicrosoftGraphChatMessage) UnsetSubject()`

UnsetSubject ensures that no value is present for Subject, not even an explicit nil
### GetSummary

`func (o *MicrosoftGraphChatMessage) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *MicrosoftGraphChatMessage) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *MicrosoftGraphChatMessage) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *MicrosoftGraphChatMessage) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### SetSummaryNil

`func (o *MicrosoftGraphChatMessage) SetSummaryNil(b bool)`

 SetSummaryNil sets the value for Summary to be an explicit nil

### UnsetSummary
`func (o *MicrosoftGraphChatMessage) UnsetSummary()`

UnsetSummary ensures that no value is present for Summary, not even an explicit nil
### GetWebUrl

`func (o *MicrosoftGraphChatMessage) GetWebUrl() string`

GetWebUrl returns the WebUrl field if non-nil, zero value otherwise.

### GetWebUrlOk

`func (o *MicrosoftGraphChatMessage) GetWebUrlOk() (*string, bool)`

GetWebUrlOk returns a tuple with the WebUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebUrl

`func (o *MicrosoftGraphChatMessage) SetWebUrl(v string)`

SetWebUrl sets WebUrl field to given value.

### HasWebUrl

`func (o *MicrosoftGraphChatMessage) HasWebUrl() bool`

HasWebUrl returns a boolean if a field has been set.

### SetWebUrlNil

`func (o *MicrosoftGraphChatMessage) SetWebUrlNil(b bool)`

 SetWebUrlNil sets the value for WebUrl to be an explicit nil

### UnsetWebUrl
`func (o *MicrosoftGraphChatMessage) UnsetWebUrl()`

UnsetWebUrl ensures that no value is present for WebUrl, not even an explicit nil
### GetHostedContents

`func (o *MicrosoftGraphChatMessage) GetHostedContents() []MicrosoftGraphChatMessageHostedContent`

GetHostedContents returns the HostedContents field if non-nil, zero value otherwise.

### GetHostedContentsOk

`func (o *MicrosoftGraphChatMessage) GetHostedContentsOk() (*[]MicrosoftGraphChatMessageHostedContent, bool)`

GetHostedContentsOk returns a tuple with the HostedContents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostedContents

`func (o *MicrosoftGraphChatMessage) SetHostedContents(v []MicrosoftGraphChatMessageHostedContent)`

SetHostedContents sets HostedContents field to given value.

### HasHostedContents

`func (o *MicrosoftGraphChatMessage) HasHostedContents() bool`

HasHostedContents returns a boolean if a field has been set.

### GetReplies

`func (o *MicrosoftGraphChatMessage) GetReplies() []MicrosoftGraphChatMessage`

GetReplies returns the Replies field if non-nil, zero value otherwise.

### GetRepliesOk

`func (o *MicrosoftGraphChatMessage) GetRepliesOk() (*[]MicrosoftGraphChatMessage, bool)`

GetRepliesOk returns a tuple with the Replies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplies

`func (o *MicrosoftGraphChatMessage) SetReplies(v []MicrosoftGraphChatMessage)`

SetReplies sets Replies field to given value.

### HasReplies

`func (o *MicrosoftGraphChatMessage) HasReplies() bool`

HasReplies returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


