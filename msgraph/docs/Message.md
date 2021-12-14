# Message

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BccRecipients** | Pointer to [**[]AnyOfmicrosoftGraphRecipient**](AnyOfmicrosoftGraphRecipient.md) | The Bcc: recipients for the message. | [optional] 
**Body** | Pointer to [**AnyOfmicrosoftGraphItemBody**](anyOf&lt;microsoft.graph.itemBody&gt;.md) | The body of the message. It can be in HTML or text format. Find out about safe HTML in a message body. | [optional] 
**BodyPreview** | Pointer to **NullableString** | The first 255 characters of the message body. It is in text format. | [optional] 
**CcRecipients** | Pointer to [**[]AnyOfmicrosoftGraphRecipient**](AnyOfmicrosoftGraphRecipient.md) | The Cc: recipients for the message. | [optional] 
**ConversationId** | Pointer to **NullableString** | The ID of the conversation the email belongs to. | [optional] 
**ConversationIndex** | Pointer to **NullableString** | Indicates the position of the message within the conversation. | [optional] 
**Flag** | Pointer to [**AnyOfmicrosoftGraphFollowupFlag**](anyOf&lt;microsoft.graph.followupFlag&gt;.md) | The flag value that indicates the status, start date, due date, or completion date for the message. | [optional] 
**From** | Pointer to [**AnyOfmicrosoftGraphRecipient**](anyOf&lt;microsoft.graph.recipient&gt;.md) | The owner of the mailbox from which the message is sent. In most cases, this value is the same as the sender property, except for sharing or delegation scenarios. The value must correspond to the actual mailbox used. Find out more about setting the from and sender properties of a message. | [optional] 
**HasAttachments** | Pointer to **NullableBool** | Indicates whether the message has attachments. This property doesn&#39;t include inline attachments, so if a message contains only inline attachments, this property is false. To verify the existence of inline attachments, parse the body property to look for a src attribute, such as &lt;IMG src&#x3D;&#39;cid:image001.jpg@01D26CD8.6C05F070&#39;&gt;. | [optional] 
**Importance** | Pointer to [**AnyOfmicrosoftGraphImportance**](anyOf&lt;microsoft.graph.importance&gt;.md) |  | [optional] 
**InferenceClassification** | Pointer to [**AnyOfmicrosoftGraphInferenceClassificationType**](anyOf&lt;microsoft.graph.inferenceClassificationType&gt;.md) |  | [optional] 
**InternetMessageHeaders** | Pointer to [**[]AnyOfmicrosoftGraphInternetMessageHeader**](AnyOfmicrosoftGraphInternetMessageHeader.md) |  | [optional] 
**InternetMessageId** | Pointer to **NullableString** |  | [optional] 
**IsDeliveryReceiptRequested** | Pointer to **NullableBool** |  | [optional] 
**IsDraft** | Pointer to **NullableBool** |  | [optional] 
**IsRead** | Pointer to **NullableBool** |  | [optional] 
**IsReadReceiptRequested** | Pointer to **NullableBool** |  | [optional] 
**ParentFolderId** | Pointer to **NullableString** |  | [optional] 
**ReceivedDateTime** | Pointer to **NullableTime** |  | [optional] 
**ReplyTo** | Pointer to [**[]AnyOfmicrosoftGraphRecipient**](AnyOfmicrosoftGraphRecipient.md) |  | [optional] 
**Sender** | Pointer to [**AnyOfmicrosoftGraphRecipient**](anyOf&lt;microsoft.graph.recipient&gt;.md) |  | [optional] 
**SentDateTime** | Pointer to **NullableTime** |  | [optional] 
**Subject** | Pointer to **NullableString** |  | [optional] 
**ToRecipients** | Pointer to [**[]AnyOfmicrosoftGraphRecipient**](AnyOfmicrosoftGraphRecipient.md) |  | [optional] 
**UniqueBody** | Pointer to [**AnyOfmicrosoftGraphItemBody**](anyOf&lt;microsoft.graph.itemBody&gt;.md) |  | [optional] 
**WebLink** | Pointer to **NullableString** |  | [optional] 
**Attachments** | Pointer to [**[]MicrosoftGraphAttachment**](MicrosoftGraphAttachment.md) | The fileAttachment and itemAttachment attachments for the message. | [optional] 
**Extensions** | Pointer to [**[]MicrosoftGraphExtension**](MicrosoftGraphExtension.md) | The collection of open extensions defined for the message. Nullable. | [optional] 
**MultiValueExtendedProperties** | Pointer to [**[]MicrosoftGraphMultiValueLegacyExtendedProperty**](MicrosoftGraphMultiValueLegacyExtendedProperty.md) | The collection of multi-value extended properties defined for the message. Nullable. | [optional] 
**SingleValueExtendedProperties** | Pointer to [**[]MicrosoftGraphSingleValueLegacyExtendedProperty**](MicrosoftGraphSingleValueLegacyExtendedProperty.md) | The collection of single-value extended properties defined for the message. Nullable. | [optional] 

## Methods

### NewMessage

`func NewMessage() *Message`

NewMessage instantiates a new Message object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageWithDefaults

`func NewMessageWithDefaults() *Message`

NewMessageWithDefaults instantiates a new Message object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBccRecipients

`func (o *Message) GetBccRecipients() []*AnyOfmicrosoftGraphRecipient`

GetBccRecipients returns the BccRecipients field if non-nil, zero value otherwise.

### GetBccRecipientsOk

`func (o *Message) GetBccRecipientsOk() (*[]*AnyOfmicrosoftGraphRecipient, bool)`

GetBccRecipientsOk returns a tuple with the BccRecipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBccRecipients

`func (o *Message) SetBccRecipients(v []*AnyOfmicrosoftGraphRecipient)`

SetBccRecipients sets BccRecipients field to given value.

### HasBccRecipients

`func (o *Message) HasBccRecipients() bool`

HasBccRecipients returns a boolean if a field has been set.

### GetBody

`func (o *Message) GetBody() AnyOfmicrosoftGraphItemBody`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *Message) GetBodyOk() (*AnyOfmicrosoftGraphItemBody, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *Message) SetBody(v AnyOfmicrosoftGraphItemBody)`

SetBody sets Body field to given value.

### HasBody

`func (o *Message) HasBody() bool`

HasBody returns a boolean if a field has been set.

### SetBodyNil

`func (o *Message) SetBodyNil(b bool)`

 SetBodyNil sets the value for Body to be an explicit nil

### UnsetBody
`func (o *Message) UnsetBody()`

UnsetBody ensures that no value is present for Body, not even an explicit nil
### GetBodyPreview

`func (o *Message) GetBodyPreview() string`

GetBodyPreview returns the BodyPreview field if non-nil, zero value otherwise.

### GetBodyPreviewOk

`func (o *Message) GetBodyPreviewOk() (*string, bool)`

GetBodyPreviewOk returns a tuple with the BodyPreview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBodyPreview

`func (o *Message) SetBodyPreview(v string)`

SetBodyPreview sets BodyPreview field to given value.

### HasBodyPreview

`func (o *Message) HasBodyPreview() bool`

HasBodyPreview returns a boolean if a field has been set.

### SetBodyPreviewNil

`func (o *Message) SetBodyPreviewNil(b bool)`

 SetBodyPreviewNil sets the value for BodyPreview to be an explicit nil

### UnsetBodyPreview
`func (o *Message) UnsetBodyPreview()`

UnsetBodyPreview ensures that no value is present for BodyPreview, not even an explicit nil
### GetCcRecipients

`func (o *Message) GetCcRecipients() []*AnyOfmicrosoftGraphRecipient`

GetCcRecipients returns the CcRecipients field if non-nil, zero value otherwise.

### GetCcRecipientsOk

`func (o *Message) GetCcRecipientsOk() (*[]*AnyOfmicrosoftGraphRecipient, bool)`

GetCcRecipientsOk returns a tuple with the CcRecipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCcRecipients

`func (o *Message) SetCcRecipients(v []*AnyOfmicrosoftGraphRecipient)`

SetCcRecipients sets CcRecipients field to given value.

### HasCcRecipients

`func (o *Message) HasCcRecipients() bool`

HasCcRecipients returns a boolean if a field has been set.

### GetConversationId

`func (o *Message) GetConversationId() string`

GetConversationId returns the ConversationId field if non-nil, zero value otherwise.

### GetConversationIdOk

`func (o *Message) GetConversationIdOk() (*string, bool)`

GetConversationIdOk returns a tuple with the ConversationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversationId

`func (o *Message) SetConversationId(v string)`

SetConversationId sets ConversationId field to given value.

### HasConversationId

`func (o *Message) HasConversationId() bool`

HasConversationId returns a boolean if a field has been set.

### SetConversationIdNil

`func (o *Message) SetConversationIdNil(b bool)`

 SetConversationIdNil sets the value for ConversationId to be an explicit nil

### UnsetConversationId
`func (o *Message) UnsetConversationId()`

UnsetConversationId ensures that no value is present for ConversationId, not even an explicit nil
### GetConversationIndex

`func (o *Message) GetConversationIndex() string`

GetConversationIndex returns the ConversationIndex field if non-nil, zero value otherwise.

### GetConversationIndexOk

`func (o *Message) GetConversationIndexOk() (*string, bool)`

GetConversationIndexOk returns a tuple with the ConversationIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversationIndex

`func (o *Message) SetConversationIndex(v string)`

SetConversationIndex sets ConversationIndex field to given value.

### HasConversationIndex

`func (o *Message) HasConversationIndex() bool`

HasConversationIndex returns a boolean if a field has been set.

### SetConversationIndexNil

`func (o *Message) SetConversationIndexNil(b bool)`

 SetConversationIndexNil sets the value for ConversationIndex to be an explicit nil

### UnsetConversationIndex
`func (o *Message) UnsetConversationIndex()`

UnsetConversationIndex ensures that no value is present for ConversationIndex, not even an explicit nil
### GetFlag

`func (o *Message) GetFlag() AnyOfmicrosoftGraphFollowupFlag`

GetFlag returns the Flag field if non-nil, zero value otherwise.

### GetFlagOk

`func (o *Message) GetFlagOk() (*AnyOfmicrosoftGraphFollowupFlag, bool)`

GetFlagOk returns a tuple with the Flag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlag

`func (o *Message) SetFlag(v AnyOfmicrosoftGraphFollowupFlag)`

SetFlag sets Flag field to given value.

### HasFlag

`func (o *Message) HasFlag() bool`

HasFlag returns a boolean if a field has been set.

### SetFlagNil

`func (o *Message) SetFlagNil(b bool)`

 SetFlagNil sets the value for Flag to be an explicit nil

### UnsetFlag
`func (o *Message) UnsetFlag()`

UnsetFlag ensures that no value is present for Flag, not even an explicit nil
### GetFrom

`func (o *Message) GetFrom() AnyOfmicrosoftGraphRecipient`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *Message) GetFromOk() (*AnyOfmicrosoftGraphRecipient, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *Message) SetFrom(v AnyOfmicrosoftGraphRecipient)`

SetFrom sets From field to given value.

### HasFrom

`func (o *Message) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### SetFromNil

`func (o *Message) SetFromNil(b bool)`

 SetFromNil sets the value for From to be an explicit nil

### UnsetFrom
`func (o *Message) UnsetFrom()`

UnsetFrom ensures that no value is present for From, not even an explicit nil
### GetHasAttachments

`func (o *Message) GetHasAttachments() bool`

GetHasAttachments returns the HasAttachments field if non-nil, zero value otherwise.

### GetHasAttachmentsOk

`func (o *Message) GetHasAttachmentsOk() (*bool, bool)`

GetHasAttachmentsOk returns a tuple with the HasAttachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAttachments

`func (o *Message) SetHasAttachments(v bool)`

SetHasAttachments sets HasAttachments field to given value.

### HasHasAttachments

`func (o *Message) HasHasAttachments() bool`

HasHasAttachments returns a boolean if a field has been set.

### SetHasAttachmentsNil

`func (o *Message) SetHasAttachmentsNil(b bool)`

 SetHasAttachmentsNil sets the value for HasAttachments to be an explicit nil

### UnsetHasAttachments
`func (o *Message) UnsetHasAttachments()`

UnsetHasAttachments ensures that no value is present for HasAttachments, not even an explicit nil
### GetImportance

`func (o *Message) GetImportance() AnyOfmicrosoftGraphImportance`

GetImportance returns the Importance field if non-nil, zero value otherwise.

### GetImportanceOk

`func (o *Message) GetImportanceOk() (*AnyOfmicrosoftGraphImportance, bool)`

GetImportanceOk returns a tuple with the Importance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportance

`func (o *Message) SetImportance(v AnyOfmicrosoftGraphImportance)`

SetImportance sets Importance field to given value.

### HasImportance

`func (o *Message) HasImportance() bool`

HasImportance returns a boolean if a field has been set.

### SetImportanceNil

`func (o *Message) SetImportanceNil(b bool)`

 SetImportanceNil sets the value for Importance to be an explicit nil

### UnsetImportance
`func (o *Message) UnsetImportance()`

UnsetImportance ensures that no value is present for Importance, not even an explicit nil
### GetInferenceClassification

`func (o *Message) GetInferenceClassification() AnyOfmicrosoftGraphInferenceClassificationType`

GetInferenceClassification returns the InferenceClassification field if non-nil, zero value otherwise.

### GetInferenceClassificationOk

`func (o *Message) GetInferenceClassificationOk() (*AnyOfmicrosoftGraphInferenceClassificationType, bool)`

GetInferenceClassificationOk returns a tuple with the InferenceClassification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInferenceClassification

`func (o *Message) SetInferenceClassification(v AnyOfmicrosoftGraphInferenceClassificationType)`

SetInferenceClassification sets InferenceClassification field to given value.

### HasInferenceClassification

`func (o *Message) HasInferenceClassification() bool`

HasInferenceClassification returns a boolean if a field has been set.

### SetInferenceClassificationNil

`func (o *Message) SetInferenceClassificationNil(b bool)`

 SetInferenceClassificationNil sets the value for InferenceClassification to be an explicit nil

### UnsetInferenceClassification
`func (o *Message) UnsetInferenceClassification()`

UnsetInferenceClassification ensures that no value is present for InferenceClassification, not even an explicit nil
### GetInternetMessageHeaders

`func (o *Message) GetInternetMessageHeaders() []*AnyOfmicrosoftGraphInternetMessageHeader`

GetInternetMessageHeaders returns the InternetMessageHeaders field if non-nil, zero value otherwise.

### GetInternetMessageHeadersOk

`func (o *Message) GetInternetMessageHeadersOk() (*[]*AnyOfmicrosoftGraphInternetMessageHeader, bool)`

GetInternetMessageHeadersOk returns a tuple with the InternetMessageHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternetMessageHeaders

`func (o *Message) SetInternetMessageHeaders(v []*AnyOfmicrosoftGraphInternetMessageHeader)`

SetInternetMessageHeaders sets InternetMessageHeaders field to given value.

### HasInternetMessageHeaders

`func (o *Message) HasInternetMessageHeaders() bool`

HasInternetMessageHeaders returns a boolean if a field has been set.

### GetInternetMessageId

`func (o *Message) GetInternetMessageId() string`

GetInternetMessageId returns the InternetMessageId field if non-nil, zero value otherwise.

### GetInternetMessageIdOk

`func (o *Message) GetInternetMessageIdOk() (*string, bool)`

GetInternetMessageIdOk returns a tuple with the InternetMessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternetMessageId

`func (o *Message) SetInternetMessageId(v string)`

SetInternetMessageId sets InternetMessageId field to given value.

### HasInternetMessageId

`func (o *Message) HasInternetMessageId() bool`

HasInternetMessageId returns a boolean if a field has been set.

### SetInternetMessageIdNil

`func (o *Message) SetInternetMessageIdNil(b bool)`

 SetInternetMessageIdNil sets the value for InternetMessageId to be an explicit nil

### UnsetInternetMessageId
`func (o *Message) UnsetInternetMessageId()`

UnsetInternetMessageId ensures that no value is present for InternetMessageId, not even an explicit nil
### GetIsDeliveryReceiptRequested

`func (o *Message) GetIsDeliveryReceiptRequested() bool`

GetIsDeliveryReceiptRequested returns the IsDeliveryReceiptRequested field if non-nil, zero value otherwise.

### GetIsDeliveryReceiptRequestedOk

`func (o *Message) GetIsDeliveryReceiptRequestedOk() (*bool, bool)`

GetIsDeliveryReceiptRequestedOk returns a tuple with the IsDeliveryReceiptRequested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeliveryReceiptRequested

`func (o *Message) SetIsDeliveryReceiptRequested(v bool)`

SetIsDeliveryReceiptRequested sets IsDeliveryReceiptRequested field to given value.

### HasIsDeliveryReceiptRequested

`func (o *Message) HasIsDeliveryReceiptRequested() bool`

HasIsDeliveryReceiptRequested returns a boolean if a field has been set.

### SetIsDeliveryReceiptRequestedNil

`func (o *Message) SetIsDeliveryReceiptRequestedNil(b bool)`

 SetIsDeliveryReceiptRequestedNil sets the value for IsDeliveryReceiptRequested to be an explicit nil

### UnsetIsDeliveryReceiptRequested
`func (o *Message) UnsetIsDeliveryReceiptRequested()`

UnsetIsDeliveryReceiptRequested ensures that no value is present for IsDeliveryReceiptRequested, not even an explicit nil
### GetIsDraft

`func (o *Message) GetIsDraft() bool`

GetIsDraft returns the IsDraft field if non-nil, zero value otherwise.

### GetIsDraftOk

`func (o *Message) GetIsDraftOk() (*bool, bool)`

GetIsDraftOk returns a tuple with the IsDraft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDraft

`func (o *Message) SetIsDraft(v bool)`

SetIsDraft sets IsDraft field to given value.

### HasIsDraft

`func (o *Message) HasIsDraft() bool`

HasIsDraft returns a boolean if a field has been set.

### SetIsDraftNil

`func (o *Message) SetIsDraftNil(b bool)`

 SetIsDraftNil sets the value for IsDraft to be an explicit nil

### UnsetIsDraft
`func (o *Message) UnsetIsDraft()`

UnsetIsDraft ensures that no value is present for IsDraft, not even an explicit nil
### GetIsRead

`func (o *Message) GetIsRead() bool`

GetIsRead returns the IsRead field if non-nil, zero value otherwise.

### GetIsReadOk

`func (o *Message) GetIsReadOk() (*bool, bool)`

GetIsReadOk returns a tuple with the IsRead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRead

`func (o *Message) SetIsRead(v bool)`

SetIsRead sets IsRead field to given value.

### HasIsRead

`func (o *Message) HasIsRead() bool`

HasIsRead returns a boolean if a field has been set.

### SetIsReadNil

`func (o *Message) SetIsReadNil(b bool)`

 SetIsReadNil sets the value for IsRead to be an explicit nil

### UnsetIsRead
`func (o *Message) UnsetIsRead()`

UnsetIsRead ensures that no value is present for IsRead, not even an explicit nil
### GetIsReadReceiptRequested

`func (o *Message) GetIsReadReceiptRequested() bool`

GetIsReadReceiptRequested returns the IsReadReceiptRequested field if non-nil, zero value otherwise.

### GetIsReadReceiptRequestedOk

`func (o *Message) GetIsReadReceiptRequestedOk() (*bool, bool)`

GetIsReadReceiptRequestedOk returns a tuple with the IsReadReceiptRequested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReadReceiptRequested

`func (o *Message) SetIsReadReceiptRequested(v bool)`

SetIsReadReceiptRequested sets IsReadReceiptRequested field to given value.

### HasIsReadReceiptRequested

`func (o *Message) HasIsReadReceiptRequested() bool`

HasIsReadReceiptRequested returns a boolean if a field has been set.

### SetIsReadReceiptRequestedNil

`func (o *Message) SetIsReadReceiptRequestedNil(b bool)`

 SetIsReadReceiptRequestedNil sets the value for IsReadReceiptRequested to be an explicit nil

### UnsetIsReadReceiptRequested
`func (o *Message) UnsetIsReadReceiptRequested()`

UnsetIsReadReceiptRequested ensures that no value is present for IsReadReceiptRequested, not even an explicit nil
### GetParentFolderId

`func (o *Message) GetParentFolderId() string`

GetParentFolderId returns the ParentFolderId field if non-nil, zero value otherwise.

### GetParentFolderIdOk

`func (o *Message) GetParentFolderIdOk() (*string, bool)`

GetParentFolderIdOk returns a tuple with the ParentFolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentFolderId

`func (o *Message) SetParentFolderId(v string)`

SetParentFolderId sets ParentFolderId field to given value.

### HasParentFolderId

`func (o *Message) HasParentFolderId() bool`

HasParentFolderId returns a boolean if a field has been set.

### SetParentFolderIdNil

`func (o *Message) SetParentFolderIdNil(b bool)`

 SetParentFolderIdNil sets the value for ParentFolderId to be an explicit nil

### UnsetParentFolderId
`func (o *Message) UnsetParentFolderId()`

UnsetParentFolderId ensures that no value is present for ParentFolderId, not even an explicit nil
### GetReceivedDateTime

`func (o *Message) GetReceivedDateTime() time.Time`

GetReceivedDateTime returns the ReceivedDateTime field if non-nil, zero value otherwise.

### GetReceivedDateTimeOk

`func (o *Message) GetReceivedDateTimeOk() (*time.Time, bool)`

GetReceivedDateTimeOk returns a tuple with the ReceivedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivedDateTime

`func (o *Message) SetReceivedDateTime(v time.Time)`

SetReceivedDateTime sets ReceivedDateTime field to given value.

### HasReceivedDateTime

`func (o *Message) HasReceivedDateTime() bool`

HasReceivedDateTime returns a boolean if a field has been set.

### SetReceivedDateTimeNil

`func (o *Message) SetReceivedDateTimeNil(b bool)`

 SetReceivedDateTimeNil sets the value for ReceivedDateTime to be an explicit nil

### UnsetReceivedDateTime
`func (o *Message) UnsetReceivedDateTime()`

UnsetReceivedDateTime ensures that no value is present for ReceivedDateTime, not even an explicit nil
### GetReplyTo

`func (o *Message) GetReplyTo() []*AnyOfmicrosoftGraphRecipient`

GetReplyTo returns the ReplyTo field if non-nil, zero value otherwise.

### GetReplyToOk

`func (o *Message) GetReplyToOk() (*[]*AnyOfmicrosoftGraphRecipient, bool)`

GetReplyToOk returns a tuple with the ReplyTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyTo

`func (o *Message) SetReplyTo(v []*AnyOfmicrosoftGraphRecipient)`

SetReplyTo sets ReplyTo field to given value.

### HasReplyTo

`func (o *Message) HasReplyTo() bool`

HasReplyTo returns a boolean if a field has been set.

### GetSender

`func (o *Message) GetSender() AnyOfmicrosoftGraphRecipient`

GetSender returns the Sender field if non-nil, zero value otherwise.

### GetSenderOk

`func (o *Message) GetSenderOk() (*AnyOfmicrosoftGraphRecipient, bool)`

GetSenderOk returns a tuple with the Sender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSender

`func (o *Message) SetSender(v AnyOfmicrosoftGraphRecipient)`

SetSender sets Sender field to given value.

### HasSender

`func (o *Message) HasSender() bool`

HasSender returns a boolean if a field has been set.

### SetSenderNil

`func (o *Message) SetSenderNil(b bool)`

 SetSenderNil sets the value for Sender to be an explicit nil

### UnsetSender
`func (o *Message) UnsetSender()`

UnsetSender ensures that no value is present for Sender, not even an explicit nil
### GetSentDateTime

`func (o *Message) GetSentDateTime() time.Time`

GetSentDateTime returns the SentDateTime field if non-nil, zero value otherwise.

### GetSentDateTimeOk

`func (o *Message) GetSentDateTimeOk() (*time.Time, bool)`

GetSentDateTimeOk returns a tuple with the SentDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSentDateTime

`func (o *Message) SetSentDateTime(v time.Time)`

SetSentDateTime sets SentDateTime field to given value.

### HasSentDateTime

`func (o *Message) HasSentDateTime() bool`

HasSentDateTime returns a boolean if a field has been set.

### SetSentDateTimeNil

`func (o *Message) SetSentDateTimeNil(b bool)`

 SetSentDateTimeNil sets the value for SentDateTime to be an explicit nil

### UnsetSentDateTime
`func (o *Message) UnsetSentDateTime()`

UnsetSentDateTime ensures that no value is present for SentDateTime, not even an explicit nil
### GetSubject

`func (o *Message) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *Message) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *Message) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *Message) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### SetSubjectNil

`func (o *Message) SetSubjectNil(b bool)`

 SetSubjectNil sets the value for Subject to be an explicit nil

### UnsetSubject
`func (o *Message) UnsetSubject()`

UnsetSubject ensures that no value is present for Subject, not even an explicit nil
### GetToRecipients

`func (o *Message) GetToRecipients() []*AnyOfmicrosoftGraphRecipient`

GetToRecipients returns the ToRecipients field if non-nil, zero value otherwise.

### GetToRecipientsOk

`func (o *Message) GetToRecipientsOk() (*[]*AnyOfmicrosoftGraphRecipient, bool)`

GetToRecipientsOk returns a tuple with the ToRecipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToRecipients

`func (o *Message) SetToRecipients(v []*AnyOfmicrosoftGraphRecipient)`

SetToRecipients sets ToRecipients field to given value.

### HasToRecipients

`func (o *Message) HasToRecipients() bool`

HasToRecipients returns a boolean if a field has been set.

### GetUniqueBody

`func (o *Message) GetUniqueBody() AnyOfmicrosoftGraphItemBody`

GetUniqueBody returns the UniqueBody field if non-nil, zero value otherwise.

### GetUniqueBodyOk

`func (o *Message) GetUniqueBodyOk() (*AnyOfmicrosoftGraphItemBody, bool)`

GetUniqueBodyOk returns a tuple with the UniqueBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueBody

`func (o *Message) SetUniqueBody(v AnyOfmicrosoftGraphItemBody)`

SetUniqueBody sets UniqueBody field to given value.

### HasUniqueBody

`func (o *Message) HasUniqueBody() bool`

HasUniqueBody returns a boolean if a field has been set.

### SetUniqueBodyNil

`func (o *Message) SetUniqueBodyNil(b bool)`

 SetUniqueBodyNil sets the value for UniqueBody to be an explicit nil

### UnsetUniqueBody
`func (o *Message) UnsetUniqueBody()`

UnsetUniqueBody ensures that no value is present for UniqueBody, not even an explicit nil
### GetWebLink

`func (o *Message) GetWebLink() string`

GetWebLink returns the WebLink field if non-nil, zero value otherwise.

### GetWebLinkOk

`func (o *Message) GetWebLinkOk() (*string, bool)`

GetWebLinkOk returns a tuple with the WebLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebLink

`func (o *Message) SetWebLink(v string)`

SetWebLink sets WebLink field to given value.

### HasWebLink

`func (o *Message) HasWebLink() bool`

HasWebLink returns a boolean if a field has been set.

### SetWebLinkNil

`func (o *Message) SetWebLinkNil(b bool)`

 SetWebLinkNil sets the value for WebLink to be an explicit nil

### UnsetWebLink
`func (o *Message) UnsetWebLink()`

UnsetWebLink ensures that no value is present for WebLink, not even an explicit nil
### GetAttachments

`func (o *Message) GetAttachments() []MicrosoftGraphAttachment`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *Message) GetAttachmentsOk() (*[]MicrosoftGraphAttachment, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *Message) SetAttachments(v []MicrosoftGraphAttachment)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *Message) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetExtensions

`func (o *Message) GetExtensions() []MicrosoftGraphExtension`

GetExtensions returns the Extensions field if non-nil, zero value otherwise.

### GetExtensionsOk

`func (o *Message) GetExtensionsOk() (*[]MicrosoftGraphExtension, bool)`

GetExtensionsOk returns a tuple with the Extensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensions

`func (o *Message) SetExtensions(v []MicrosoftGraphExtension)`

SetExtensions sets Extensions field to given value.

### HasExtensions

`func (o *Message) HasExtensions() bool`

HasExtensions returns a boolean if a field has been set.

### GetMultiValueExtendedProperties

`func (o *Message) GetMultiValueExtendedProperties() []MicrosoftGraphMultiValueLegacyExtendedProperty`

GetMultiValueExtendedProperties returns the MultiValueExtendedProperties field if non-nil, zero value otherwise.

### GetMultiValueExtendedPropertiesOk

`func (o *Message) GetMultiValueExtendedPropertiesOk() (*[]MicrosoftGraphMultiValueLegacyExtendedProperty, bool)`

GetMultiValueExtendedPropertiesOk returns a tuple with the MultiValueExtendedProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiValueExtendedProperties

`func (o *Message) SetMultiValueExtendedProperties(v []MicrosoftGraphMultiValueLegacyExtendedProperty)`

SetMultiValueExtendedProperties sets MultiValueExtendedProperties field to given value.

### HasMultiValueExtendedProperties

`func (o *Message) HasMultiValueExtendedProperties() bool`

HasMultiValueExtendedProperties returns a boolean if a field has been set.

### GetSingleValueExtendedProperties

`func (o *Message) GetSingleValueExtendedProperties() []MicrosoftGraphSingleValueLegacyExtendedProperty`

GetSingleValueExtendedProperties returns the SingleValueExtendedProperties field if non-nil, zero value otherwise.

### GetSingleValueExtendedPropertiesOk

`func (o *Message) GetSingleValueExtendedPropertiesOk() (*[]MicrosoftGraphSingleValueLegacyExtendedProperty, bool)`

GetSingleValueExtendedPropertiesOk returns a tuple with the SingleValueExtendedProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSingleValueExtendedProperties

`func (o *Message) SetSingleValueExtendedProperties(v []MicrosoftGraphSingleValueLegacyExtendedProperty)`

SetSingleValueExtendedProperties sets SingleValueExtendedProperties field to given value.

### HasSingleValueExtendedProperties

`func (o *Message) HasSingleValueExtendedProperties() bool`

HasSingleValueExtendedProperties returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


