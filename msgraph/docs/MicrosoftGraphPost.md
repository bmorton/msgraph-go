# MicrosoftGraphPost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**Categories** | Pointer to **[]string** | The categories associated with the item | [optional] 
**ChangeKey** | Pointer to **NullableString** | Identifies the version of the item. Every time the item is changed, changeKey changes as well. This allows Exchange to apply changes to the correct version of the object. Read-only. | [optional] 
**CreatedDateTime** | Pointer to **NullableTime** | The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z | [optional] 
**LastModifiedDateTime** | Pointer to **NullableTime** | The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z | [optional] 
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

### NewMicrosoftGraphPost

`func NewMicrosoftGraphPost() *MicrosoftGraphPost`

NewMicrosoftGraphPost instantiates a new MicrosoftGraphPost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphPostWithDefaults

`func NewMicrosoftGraphPostWithDefaults() *MicrosoftGraphPost`

NewMicrosoftGraphPostWithDefaults instantiates a new MicrosoftGraphPost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphPost) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphPost) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphPost) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphPost) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCategories

`func (o *MicrosoftGraphPost) GetCategories() []*string`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *MicrosoftGraphPost) GetCategoriesOk() (*[]*string, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *MicrosoftGraphPost) SetCategories(v []*string)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *MicrosoftGraphPost) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### GetChangeKey

`func (o *MicrosoftGraphPost) GetChangeKey() string`

GetChangeKey returns the ChangeKey field if non-nil, zero value otherwise.

### GetChangeKeyOk

`func (o *MicrosoftGraphPost) GetChangeKeyOk() (*string, bool)`

GetChangeKeyOk returns a tuple with the ChangeKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeKey

`func (o *MicrosoftGraphPost) SetChangeKey(v string)`

SetChangeKey sets ChangeKey field to given value.

### HasChangeKey

`func (o *MicrosoftGraphPost) HasChangeKey() bool`

HasChangeKey returns a boolean if a field has been set.

### SetChangeKeyNil

`func (o *MicrosoftGraphPost) SetChangeKeyNil(b bool)`

 SetChangeKeyNil sets the value for ChangeKey to be an explicit nil

### UnsetChangeKey
`func (o *MicrosoftGraphPost) UnsetChangeKey()`

UnsetChangeKey ensures that no value is present for ChangeKey, not even an explicit nil
### GetCreatedDateTime

`func (o *MicrosoftGraphPost) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *MicrosoftGraphPost) GetCreatedDateTimeOk() (*time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateTime

`func (o *MicrosoftGraphPost) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime sets CreatedDateTime field to given value.

### HasCreatedDateTime

`func (o *MicrosoftGraphPost) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTimeNil

`func (o *MicrosoftGraphPost) SetCreatedDateTimeNil(b bool)`

 SetCreatedDateTimeNil sets the value for CreatedDateTime to be an explicit nil

### UnsetCreatedDateTime
`func (o *MicrosoftGraphPost) UnsetCreatedDateTime()`

UnsetCreatedDateTime ensures that no value is present for CreatedDateTime, not even an explicit nil
### GetLastModifiedDateTime

`func (o *MicrosoftGraphPost) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *MicrosoftGraphPost) GetLastModifiedDateTimeOk() (*time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedDateTime

`func (o *MicrosoftGraphPost) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime sets LastModifiedDateTime field to given value.

### HasLastModifiedDateTime

`func (o *MicrosoftGraphPost) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### SetLastModifiedDateTimeNil

`func (o *MicrosoftGraphPost) SetLastModifiedDateTimeNil(b bool)`

 SetLastModifiedDateTimeNil sets the value for LastModifiedDateTime to be an explicit nil

### UnsetLastModifiedDateTime
`func (o *MicrosoftGraphPost) UnsetLastModifiedDateTime()`

UnsetLastModifiedDateTime ensures that no value is present for LastModifiedDateTime, not even an explicit nil
### GetBody

`func (o *MicrosoftGraphPost) GetBody() AnyOfmicrosoftGraphItemBody`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *MicrosoftGraphPost) GetBodyOk() (*AnyOfmicrosoftGraphItemBody, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *MicrosoftGraphPost) SetBody(v AnyOfmicrosoftGraphItemBody)`

SetBody sets Body field to given value.

### HasBody

`func (o *MicrosoftGraphPost) HasBody() bool`

HasBody returns a boolean if a field has been set.

### SetBodyNil

`func (o *MicrosoftGraphPost) SetBodyNil(b bool)`

 SetBodyNil sets the value for Body to be an explicit nil

### UnsetBody
`func (o *MicrosoftGraphPost) UnsetBody()`

UnsetBody ensures that no value is present for Body, not even an explicit nil
### GetConversationId

`func (o *MicrosoftGraphPost) GetConversationId() string`

GetConversationId returns the ConversationId field if non-nil, zero value otherwise.

### GetConversationIdOk

`func (o *MicrosoftGraphPost) GetConversationIdOk() (*string, bool)`

GetConversationIdOk returns a tuple with the ConversationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversationId

`func (o *MicrosoftGraphPost) SetConversationId(v string)`

SetConversationId sets ConversationId field to given value.

### HasConversationId

`func (o *MicrosoftGraphPost) HasConversationId() bool`

HasConversationId returns a boolean if a field has been set.

### SetConversationIdNil

`func (o *MicrosoftGraphPost) SetConversationIdNil(b bool)`

 SetConversationIdNil sets the value for ConversationId to be an explicit nil

### UnsetConversationId
`func (o *MicrosoftGraphPost) UnsetConversationId()`

UnsetConversationId ensures that no value is present for ConversationId, not even an explicit nil
### GetConversationThreadId

`func (o *MicrosoftGraphPost) GetConversationThreadId() string`

GetConversationThreadId returns the ConversationThreadId field if non-nil, zero value otherwise.

### GetConversationThreadIdOk

`func (o *MicrosoftGraphPost) GetConversationThreadIdOk() (*string, bool)`

GetConversationThreadIdOk returns a tuple with the ConversationThreadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversationThreadId

`func (o *MicrosoftGraphPost) SetConversationThreadId(v string)`

SetConversationThreadId sets ConversationThreadId field to given value.

### HasConversationThreadId

`func (o *MicrosoftGraphPost) HasConversationThreadId() bool`

HasConversationThreadId returns a boolean if a field has been set.

### SetConversationThreadIdNil

`func (o *MicrosoftGraphPost) SetConversationThreadIdNil(b bool)`

 SetConversationThreadIdNil sets the value for ConversationThreadId to be an explicit nil

### UnsetConversationThreadId
`func (o *MicrosoftGraphPost) UnsetConversationThreadId()`

UnsetConversationThreadId ensures that no value is present for ConversationThreadId, not even an explicit nil
### GetFrom

`func (o *MicrosoftGraphPost) GetFrom() MicrosoftGraphRecipient`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *MicrosoftGraphPost) GetFromOk() (*MicrosoftGraphRecipient, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *MicrosoftGraphPost) SetFrom(v MicrosoftGraphRecipient)`

SetFrom sets From field to given value.

### HasFrom

`func (o *MicrosoftGraphPost) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetHasAttachments

`func (o *MicrosoftGraphPost) GetHasAttachments() bool`

GetHasAttachments returns the HasAttachments field if non-nil, zero value otherwise.

### GetHasAttachmentsOk

`func (o *MicrosoftGraphPost) GetHasAttachmentsOk() (*bool, bool)`

GetHasAttachmentsOk returns a tuple with the HasAttachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAttachments

`func (o *MicrosoftGraphPost) SetHasAttachments(v bool)`

SetHasAttachments sets HasAttachments field to given value.

### HasHasAttachments

`func (o *MicrosoftGraphPost) HasHasAttachments() bool`

HasHasAttachments returns a boolean if a field has been set.

### GetNewParticipants

`func (o *MicrosoftGraphPost) GetNewParticipants() []MicrosoftGraphRecipient`

GetNewParticipants returns the NewParticipants field if non-nil, zero value otherwise.

### GetNewParticipantsOk

`func (o *MicrosoftGraphPost) GetNewParticipantsOk() (*[]MicrosoftGraphRecipient, bool)`

GetNewParticipantsOk returns a tuple with the NewParticipants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewParticipants

`func (o *MicrosoftGraphPost) SetNewParticipants(v []MicrosoftGraphRecipient)`

SetNewParticipants sets NewParticipants field to given value.

### HasNewParticipants

`func (o *MicrosoftGraphPost) HasNewParticipants() bool`

HasNewParticipants returns a boolean if a field has been set.

### GetReceivedDateTime

`func (o *MicrosoftGraphPost) GetReceivedDateTime() time.Time`

GetReceivedDateTime returns the ReceivedDateTime field if non-nil, zero value otherwise.

### GetReceivedDateTimeOk

`func (o *MicrosoftGraphPost) GetReceivedDateTimeOk() (*time.Time, bool)`

GetReceivedDateTimeOk returns a tuple with the ReceivedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivedDateTime

`func (o *MicrosoftGraphPost) SetReceivedDateTime(v time.Time)`

SetReceivedDateTime sets ReceivedDateTime field to given value.

### HasReceivedDateTime

`func (o *MicrosoftGraphPost) HasReceivedDateTime() bool`

HasReceivedDateTime returns a boolean if a field has been set.

### GetSender

`func (o *MicrosoftGraphPost) GetSender() AnyOfmicrosoftGraphRecipient`

GetSender returns the Sender field if non-nil, zero value otherwise.

### GetSenderOk

`func (o *MicrosoftGraphPost) GetSenderOk() (*AnyOfmicrosoftGraphRecipient, bool)`

GetSenderOk returns a tuple with the Sender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSender

`func (o *MicrosoftGraphPost) SetSender(v AnyOfmicrosoftGraphRecipient)`

SetSender sets Sender field to given value.

### HasSender

`func (o *MicrosoftGraphPost) HasSender() bool`

HasSender returns a boolean if a field has been set.

### SetSenderNil

`func (o *MicrosoftGraphPost) SetSenderNil(b bool)`

 SetSenderNil sets the value for Sender to be an explicit nil

### UnsetSender
`func (o *MicrosoftGraphPost) UnsetSender()`

UnsetSender ensures that no value is present for Sender, not even an explicit nil
### GetAttachments

`func (o *MicrosoftGraphPost) GetAttachments() []MicrosoftGraphAttachment`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *MicrosoftGraphPost) GetAttachmentsOk() (*[]MicrosoftGraphAttachment, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *MicrosoftGraphPost) SetAttachments(v []MicrosoftGraphAttachment)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *MicrosoftGraphPost) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetExtensions

`func (o *MicrosoftGraphPost) GetExtensions() []MicrosoftGraphExtension`

GetExtensions returns the Extensions field if non-nil, zero value otherwise.

### GetExtensionsOk

`func (o *MicrosoftGraphPost) GetExtensionsOk() (*[]MicrosoftGraphExtension, bool)`

GetExtensionsOk returns a tuple with the Extensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensions

`func (o *MicrosoftGraphPost) SetExtensions(v []MicrosoftGraphExtension)`

SetExtensions sets Extensions field to given value.

### HasExtensions

`func (o *MicrosoftGraphPost) HasExtensions() bool`

HasExtensions returns a boolean if a field has been set.

### GetInReplyTo

`func (o *MicrosoftGraphPost) GetInReplyTo() AnyOfmicrosoftGraphPost`

GetInReplyTo returns the InReplyTo field if non-nil, zero value otherwise.

### GetInReplyToOk

`func (o *MicrosoftGraphPost) GetInReplyToOk() (*AnyOfmicrosoftGraphPost, bool)`

GetInReplyToOk returns a tuple with the InReplyTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInReplyTo

`func (o *MicrosoftGraphPost) SetInReplyTo(v AnyOfmicrosoftGraphPost)`

SetInReplyTo sets InReplyTo field to given value.

### HasInReplyTo

`func (o *MicrosoftGraphPost) HasInReplyTo() bool`

HasInReplyTo returns a boolean if a field has been set.

### SetInReplyToNil

`func (o *MicrosoftGraphPost) SetInReplyToNil(b bool)`

 SetInReplyToNil sets the value for InReplyTo to be an explicit nil

### UnsetInReplyTo
`func (o *MicrosoftGraphPost) UnsetInReplyTo()`

UnsetInReplyTo ensures that no value is present for InReplyTo, not even an explicit nil
### GetMultiValueExtendedProperties

`func (o *MicrosoftGraphPost) GetMultiValueExtendedProperties() []MicrosoftGraphMultiValueLegacyExtendedProperty`

GetMultiValueExtendedProperties returns the MultiValueExtendedProperties field if non-nil, zero value otherwise.

### GetMultiValueExtendedPropertiesOk

`func (o *MicrosoftGraphPost) GetMultiValueExtendedPropertiesOk() (*[]MicrosoftGraphMultiValueLegacyExtendedProperty, bool)`

GetMultiValueExtendedPropertiesOk returns a tuple with the MultiValueExtendedProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiValueExtendedProperties

`func (o *MicrosoftGraphPost) SetMultiValueExtendedProperties(v []MicrosoftGraphMultiValueLegacyExtendedProperty)`

SetMultiValueExtendedProperties sets MultiValueExtendedProperties field to given value.

### HasMultiValueExtendedProperties

`func (o *MicrosoftGraphPost) HasMultiValueExtendedProperties() bool`

HasMultiValueExtendedProperties returns a boolean if a field has been set.

### GetSingleValueExtendedProperties

`func (o *MicrosoftGraphPost) GetSingleValueExtendedProperties() []MicrosoftGraphSingleValueLegacyExtendedProperty`

GetSingleValueExtendedProperties returns the SingleValueExtendedProperties field if non-nil, zero value otherwise.

### GetSingleValueExtendedPropertiesOk

`func (o *MicrosoftGraphPost) GetSingleValueExtendedPropertiesOk() (*[]MicrosoftGraphSingleValueLegacyExtendedProperty, bool)`

GetSingleValueExtendedPropertiesOk returns a tuple with the SingleValueExtendedProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSingleValueExtendedProperties

`func (o *MicrosoftGraphPost) SetSingleValueExtendedProperties(v []MicrosoftGraphSingleValueLegacyExtendedProperty)`

SetSingleValueExtendedProperties sets SingleValueExtendedProperties field to given value.

### HasSingleValueExtendedProperties

`func (o *MicrosoftGraphPost) HasSingleValueExtendedProperties() bool`

HasSingleValueExtendedProperties returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


