# Post

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Body** | Pointer to [**AnyOfmicrosoftGraphItemBody**](anyOf&lt;microsoft.graph.itemBody&gt;.md) | The contents of the post. This is a default property. This property can be null. | [optional] 
**ConversationId** | Pointer to **NullableString** | Unique ID of the conversation. Read-only. | [optional] 
**ConversationThreadId** | Pointer to **NullableString** | Unique ID of the conversation thread. Read-only. | [optional] 
**From** | Pointer to [**MicrosoftGraphRecipient**](MicrosoftGraphRecipient.md) |  | [optional] 
**HasAttachments** | Pointer to **bool** | Indicates whether the post has at least one attachment. This is a default property. | [optional] 
**NewParticipants** | Pointer to [**[]MicrosoftGraphRecipient**](MicrosoftGraphRecipient.md) | Conversation participants that were added to the thread as part of this post. | [optional] 
**ReceivedDateTime** | Pointer to **time.Time** | Specifies when the post was received. The DateTimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z | [optional] 
**Sender** | Pointer to [**AnyOfmicrosoftGraphRecipient**](anyOf&lt;microsoft.graph.recipient&gt;.md) | Contains the address of the sender. The value of Sender is assumed to be the address of the authenticated user in the case when Sender is not specified. This is a default property. | [optional] 
**Attachments** | Pointer to [**[]MicrosoftGraphAttachment**](MicrosoftGraphAttachment.md) | Read-only. Nullable. Supports $expand. | [optional] 
**Extensions** | Pointer to [**[]MicrosoftGraphExtension**](MicrosoftGraphExtension.md) | The collection of open extensions defined for the post. Read-only. Nullable. Supports $expand. | [optional] 
**InReplyTo** | Pointer to [**AnyOfmicrosoftGraphPost**](anyOf&lt;microsoft.graph.post&gt;.md) | Read-only. Supports $expand. | [optional] 
**MultiValueExtendedProperties** | Pointer to [**[]MicrosoftGraphMultiValueLegacyExtendedProperty**](MicrosoftGraphMultiValueLegacyExtendedProperty.md) | The collection of multi-value extended properties defined for the post. Read-only. Nullable. | [optional] 
**SingleValueExtendedProperties** | Pointer to [**[]MicrosoftGraphSingleValueLegacyExtendedProperty**](MicrosoftGraphSingleValueLegacyExtendedProperty.md) | The collection of single-value extended properties defined for the post. Read-only. Nullable. | [optional] 

## Methods

### NewPost

`func NewPost() *Post`

NewPost instantiates a new Post object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostWithDefaults

`func NewPostWithDefaults() *Post`

NewPostWithDefaults instantiates a new Post object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBody

`func (o *Post) GetBody() AnyOfmicrosoftGraphItemBody`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *Post) GetBodyOk() (*AnyOfmicrosoftGraphItemBody, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *Post) SetBody(v AnyOfmicrosoftGraphItemBody)`

SetBody sets Body field to given value.

### HasBody

`func (o *Post) HasBody() bool`

HasBody returns a boolean if a field has been set.

### SetBodyNil

`func (o *Post) SetBodyNil(b bool)`

 SetBodyNil sets the value for Body to be an explicit nil

### UnsetBody
`func (o *Post) UnsetBody()`

UnsetBody ensures that no value is present for Body, not even an explicit nil
### GetConversationId

`func (o *Post) GetConversationId() string`

GetConversationId returns the ConversationId field if non-nil, zero value otherwise.

### GetConversationIdOk

`func (o *Post) GetConversationIdOk() (*string, bool)`

GetConversationIdOk returns a tuple with the ConversationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversationId

`func (o *Post) SetConversationId(v string)`

SetConversationId sets ConversationId field to given value.

### HasConversationId

`func (o *Post) HasConversationId() bool`

HasConversationId returns a boolean if a field has been set.

### SetConversationIdNil

`func (o *Post) SetConversationIdNil(b bool)`

 SetConversationIdNil sets the value for ConversationId to be an explicit nil

### UnsetConversationId
`func (o *Post) UnsetConversationId()`

UnsetConversationId ensures that no value is present for ConversationId, not even an explicit nil
### GetConversationThreadId

`func (o *Post) GetConversationThreadId() string`

GetConversationThreadId returns the ConversationThreadId field if non-nil, zero value otherwise.

### GetConversationThreadIdOk

`func (o *Post) GetConversationThreadIdOk() (*string, bool)`

GetConversationThreadIdOk returns a tuple with the ConversationThreadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversationThreadId

`func (o *Post) SetConversationThreadId(v string)`

SetConversationThreadId sets ConversationThreadId field to given value.

### HasConversationThreadId

`func (o *Post) HasConversationThreadId() bool`

HasConversationThreadId returns a boolean if a field has been set.

### SetConversationThreadIdNil

`func (o *Post) SetConversationThreadIdNil(b bool)`

 SetConversationThreadIdNil sets the value for ConversationThreadId to be an explicit nil

### UnsetConversationThreadId
`func (o *Post) UnsetConversationThreadId()`

UnsetConversationThreadId ensures that no value is present for ConversationThreadId, not even an explicit nil
### GetFrom

`func (o *Post) GetFrom() MicrosoftGraphRecipient`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *Post) GetFromOk() (*MicrosoftGraphRecipient, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *Post) SetFrom(v MicrosoftGraphRecipient)`

SetFrom sets From field to given value.

### HasFrom

`func (o *Post) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetHasAttachments

`func (o *Post) GetHasAttachments() bool`

GetHasAttachments returns the HasAttachments field if non-nil, zero value otherwise.

### GetHasAttachmentsOk

`func (o *Post) GetHasAttachmentsOk() (*bool, bool)`

GetHasAttachmentsOk returns a tuple with the HasAttachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAttachments

`func (o *Post) SetHasAttachments(v bool)`

SetHasAttachments sets HasAttachments field to given value.

### HasHasAttachments

`func (o *Post) HasHasAttachments() bool`

HasHasAttachments returns a boolean if a field has been set.

### GetNewParticipants

`func (o *Post) GetNewParticipants() []MicrosoftGraphRecipient`

GetNewParticipants returns the NewParticipants field if non-nil, zero value otherwise.

### GetNewParticipantsOk

`func (o *Post) GetNewParticipantsOk() (*[]MicrosoftGraphRecipient, bool)`

GetNewParticipantsOk returns a tuple with the NewParticipants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewParticipants

`func (o *Post) SetNewParticipants(v []MicrosoftGraphRecipient)`

SetNewParticipants sets NewParticipants field to given value.

### HasNewParticipants

`func (o *Post) HasNewParticipants() bool`

HasNewParticipants returns a boolean if a field has been set.

### GetReceivedDateTime

`func (o *Post) GetReceivedDateTime() time.Time`

GetReceivedDateTime returns the ReceivedDateTime field if non-nil, zero value otherwise.

### GetReceivedDateTimeOk

`func (o *Post) GetReceivedDateTimeOk() (*time.Time, bool)`

GetReceivedDateTimeOk returns a tuple with the ReceivedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivedDateTime

`func (o *Post) SetReceivedDateTime(v time.Time)`

SetReceivedDateTime sets ReceivedDateTime field to given value.

### HasReceivedDateTime

`func (o *Post) HasReceivedDateTime() bool`

HasReceivedDateTime returns a boolean if a field has been set.

### GetSender

`func (o *Post) GetSender() AnyOfmicrosoftGraphRecipient`

GetSender returns the Sender field if non-nil, zero value otherwise.

### GetSenderOk

`func (o *Post) GetSenderOk() (*AnyOfmicrosoftGraphRecipient, bool)`

GetSenderOk returns a tuple with the Sender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSender

`func (o *Post) SetSender(v AnyOfmicrosoftGraphRecipient)`

SetSender sets Sender field to given value.

### HasSender

`func (o *Post) HasSender() bool`

HasSender returns a boolean if a field has been set.

### SetSenderNil

`func (o *Post) SetSenderNil(b bool)`

 SetSenderNil sets the value for Sender to be an explicit nil

### UnsetSender
`func (o *Post) UnsetSender()`

UnsetSender ensures that no value is present for Sender, not even an explicit nil
### GetAttachments

`func (o *Post) GetAttachments() []MicrosoftGraphAttachment`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *Post) GetAttachmentsOk() (*[]MicrosoftGraphAttachment, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *Post) SetAttachments(v []MicrosoftGraphAttachment)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *Post) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetExtensions

`func (o *Post) GetExtensions() []MicrosoftGraphExtension`

GetExtensions returns the Extensions field if non-nil, zero value otherwise.

### GetExtensionsOk

`func (o *Post) GetExtensionsOk() (*[]MicrosoftGraphExtension, bool)`

GetExtensionsOk returns a tuple with the Extensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensions

`func (o *Post) SetExtensions(v []MicrosoftGraphExtension)`

SetExtensions sets Extensions field to given value.

### HasExtensions

`func (o *Post) HasExtensions() bool`

HasExtensions returns a boolean if a field has been set.

### GetInReplyTo

`func (o *Post) GetInReplyTo() AnyOfmicrosoftGraphPost`

GetInReplyTo returns the InReplyTo field if non-nil, zero value otherwise.

### GetInReplyToOk

`func (o *Post) GetInReplyToOk() (*AnyOfmicrosoftGraphPost, bool)`

GetInReplyToOk returns a tuple with the InReplyTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInReplyTo

`func (o *Post) SetInReplyTo(v AnyOfmicrosoftGraphPost)`

SetInReplyTo sets InReplyTo field to given value.

### HasInReplyTo

`func (o *Post) HasInReplyTo() bool`

HasInReplyTo returns a boolean if a field has been set.

### SetInReplyToNil

`func (o *Post) SetInReplyToNil(b bool)`

 SetInReplyToNil sets the value for InReplyTo to be an explicit nil

### UnsetInReplyTo
`func (o *Post) UnsetInReplyTo()`

UnsetInReplyTo ensures that no value is present for InReplyTo, not even an explicit nil
### GetMultiValueExtendedProperties

`func (o *Post) GetMultiValueExtendedProperties() []MicrosoftGraphMultiValueLegacyExtendedProperty`

GetMultiValueExtendedProperties returns the MultiValueExtendedProperties field if non-nil, zero value otherwise.

### GetMultiValueExtendedPropertiesOk

`func (o *Post) GetMultiValueExtendedPropertiesOk() (*[]MicrosoftGraphMultiValueLegacyExtendedProperty, bool)`

GetMultiValueExtendedPropertiesOk returns a tuple with the MultiValueExtendedProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiValueExtendedProperties

`func (o *Post) SetMultiValueExtendedProperties(v []MicrosoftGraphMultiValueLegacyExtendedProperty)`

SetMultiValueExtendedProperties sets MultiValueExtendedProperties field to given value.

### HasMultiValueExtendedProperties

`func (o *Post) HasMultiValueExtendedProperties() bool`

HasMultiValueExtendedProperties returns a boolean if a field has been set.

### GetSingleValueExtendedProperties

`func (o *Post) GetSingleValueExtendedProperties() []MicrosoftGraphSingleValueLegacyExtendedProperty`

GetSingleValueExtendedProperties returns the SingleValueExtendedProperties field if non-nil, zero value otherwise.

### GetSingleValueExtendedPropertiesOk

`func (o *Post) GetSingleValueExtendedPropertiesOk() (*[]MicrosoftGraphSingleValueLegacyExtendedProperty, bool)`

GetSingleValueExtendedPropertiesOk returns a tuple with the SingleValueExtendedProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSingleValueExtendedProperties

`func (o *Post) SetSingleValueExtendedProperties(v []MicrosoftGraphSingleValueLegacyExtendedProperty)`

SetSingleValueExtendedProperties sets SingleValueExtendedProperties field to given value.

### HasSingleValueExtendedProperties

`func (o *Post) HasSingleValueExtendedProperties() bool`

HasSingleValueExtendedProperties returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


