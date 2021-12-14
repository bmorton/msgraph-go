# MicrosoftGraphMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**Categories** | Pointer to **[]string** | The categories associated with the item | [optional] 
**ChangeKey** | Pointer to **NullableString** | Identifies the version of the item. Every time the item is changed, changeKey changes as well. This allows Exchange to apply changes to the correct version of the object. Read-only. | [optional] 
**CreatedDateTime** | Pointer to **NullableTime** | The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z | [optional] 
**LastModifiedDateTime** | Pointer to **NullableTime** | The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z | [optional] 
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

### NewMicrosoftGraphMessage

`func NewMicrosoftGraphMessage() *MicrosoftGraphMessage`

NewMicrosoftGraphMessage instantiates a new MicrosoftGraphMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphMessageWithDefaults

`func NewMicrosoftGraphMessageWithDefaults() *MicrosoftGraphMessage`

NewMicrosoftGraphMessageWithDefaults instantiates a new MicrosoftGraphMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphMessage) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphMessage) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphMessage) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphMessage) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCategories

`func (o *MicrosoftGraphMessage) GetCategories() []*string`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *MicrosoftGraphMessage) GetCategoriesOk() (*[]*string, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *MicrosoftGraphMessage) SetCategories(v []*string)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *MicrosoftGraphMessage) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### GetChangeKey

`func (o *MicrosoftGraphMessage) GetChangeKey() string`

GetChangeKey returns the ChangeKey field if non-nil, zero value otherwise.

### GetChangeKeyOk

`func (o *MicrosoftGraphMessage) GetChangeKeyOk() (*string, bool)`

GetChangeKeyOk returns a tuple with the ChangeKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeKey

`func (o *MicrosoftGraphMessage) SetChangeKey(v string)`

SetChangeKey sets ChangeKey field to given value.

### HasChangeKey

`func (o *MicrosoftGraphMessage) HasChangeKey() bool`

HasChangeKey returns a boolean if a field has been set.

### SetChangeKeyNil

`func (o *MicrosoftGraphMessage) SetChangeKeyNil(b bool)`

 SetChangeKeyNil sets the value for ChangeKey to be an explicit nil

### UnsetChangeKey
`func (o *MicrosoftGraphMessage) UnsetChangeKey()`

UnsetChangeKey ensures that no value is present for ChangeKey, not even an explicit nil
### GetCreatedDateTime

`func (o *MicrosoftGraphMessage) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *MicrosoftGraphMessage) GetCreatedDateTimeOk() (*time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateTime

`func (o *MicrosoftGraphMessage) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime sets CreatedDateTime field to given value.

### HasCreatedDateTime

`func (o *MicrosoftGraphMessage) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTimeNil

`func (o *MicrosoftGraphMessage) SetCreatedDateTimeNil(b bool)`

 SetCreatedDateTimeNil sets the value for CreatedDateTime to be an explicit nil

### UnsetCreatedDateTime
`func (o *MicrosoftGraphMessage) UnsetCreatedDateTime()`

UnsetCreatedDateTime ensures that no value is present for CreatedDateTime, not even an explicit nil
### GetLastModifiedDateTime

`func (o *MicrosoftGraphMessage) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *MicrosoftGraphMessage) GetLastModifiedDateTimeOk() (*time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedDateTime

`func (o *MicrosoftGraphMessage) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime sets LastModifiedDateTime field to given value.

### HasLastModifiedDateTime

`func (o *MicrosoftGraphMessage) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### SetLastModifiedDateTimeNil

`func (o *MicrosoftGraphMessage) SetLastModifiedDateTimeNil(b bool)`

 SetLastModifiedDateTimeNil sets the value for LastModifiedDateTime to be an explicit nil

### UnsetLastModifiedDateTime
`func (o *MicrosoftGraphMessage) UnsetLastModifiedDateTime()`

UnsetLastModifiedDateTime ensures that no value is present for LastModifiedDateTime, not even an explicit nil
### GetBccRecipients

`func (o *MicrosoftGraphMessage) GetBccRecipients() []*AnyOfmicrosoftGraphRecipient`

GetBccRecipients returns the BccRecipients field if non-nil, zero value otherwise.

### GetBccRecipientsOk

`func (o *MicrosoftGraphMessage) GetBccRecipientsOk() (*[]*AnyOfmicrosoftGraphRecipient, bool)`

GetBccRecipientsOk returns a tuple with the BccRecipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBccRecipients

`func (o *MicrosoftGraphMessage) SetBccRecipients(v []*AnyOfmicrosoftGraphRecipient)`

SetBccRecipients sets BccRecipients field to given value.

### HasBccRecipients

`func (o *MicrosoftGraphMessage) HasBccRecipients() bool`

HasBccRecipients returns a boolean if a field has been set.

### GetBody

`func (o *MicrosoftGraphMessage) GetBody() AnyOfmicrosoftGraphItemBody`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *MicrosoftGraphMessage) GetBodyOk() (*AnyOfmicrosoftGraphItemBody, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *MicrosoftGraphMessage) SetBody(v AnyOfmicrosoftGraphItemBody)`

SetBody sets Body field to given value.

### HasBody

`func (o *MicrosoftGraphMessage) HasBody() bool`

HasBody returns a boolean if a field has been set.

### SetBodyNil

`func (o *MicrosoftGraphMessage) SetBodyNil(b bool)`

 SetBodyNil sets the value for Body to be an explicit nil

### UnsetBody
`func (o *MicrosoftGraphMessage) UnsetBody()`

UnsetBody ensures that no value is present for Body, not even an explicit nil
### GetBodyPreview

`func (o *MicrosoftGraphMessage) GetBodyPreview() string`

GetBodyPreview returns the BodyPreview field if non-nil, zero value otherwise.

### GetBodyPreviewOk

`func (o *MicrosoftGraphMessage) GetBodyPreviewOk() (*string, bool)`

GetBodyPreviewOk returns a tuple with the BodyPreview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBodyPreview

`func (o *MicrosoftGraphMessage) SetBodyPreview(v string)`

SetBodyPreview sets BodyPreview field to given value.

### HasBodyPreview

`func (o *MicrosoftGraphMessage) HasBodyPreview() bool`

HasBodyPreview returns a boolean if a field has been set.

### SetBodyPreviewNil

`func (o *MicrosoftGraphMessage) SetBodyPreviewNil(b bool)`

 SetBodyPreviewNil sets the value for BodyPreview to be an explicit nil

### UnsetBodyPreview
`func (o *MicrosoftGraphMessage) UnsetBodyPreview()`

UnsetBodyPreview ensures that no value is present for BodyPreview, not even an explicit nil
### GetCcRecipients

`func (o *MicrosoftGraphMessage) GetCcRecipients() []*AnyOfmicrosoftGraphRecipient`

GetCcRecipients returns the CcRecipients field if non-nil, zero value otherwise.

### GetCcRecipientsOk

`func (o *MicrosoftGraphMessage) GetCcRecipientsOk() (*[]*AnyOfmicrosoftGraphRecipient, bool)`

GetCcRecipientsOk returns a tuple with the CcRecipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCcRecipients

`func (o *MicrosoftGraphMessage) SetCcRecipients(v []*AnyOfmicrosoftGraphRecipient)`

SetCcRecipients sets CcRecipients field to given value.

### HasCcRecipients

`func (o *MicrosoftGraphMessage) HasCcRecipients() bool`

HasCcRecipients returns a boolean if a field has been set.

### GetConversationId

`func (o *MicrosoftGraphMessage) GetConversationId() string`

GetConversationId returns the ConversationId field if non-nil, zero value otherwise.

### GetConversationIdOk

`func (o *MicrosoftGraphMessage) GetConversationIdOk() (*string, bool)`

GetConversationIdOk returns a tuple with the ConversationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversationId

`func (o *MicrosoftGraphMessage) SetConversationId(v string)`

SetConversationId sets ConversationId field to given value.

### HasConversationId

`func (o *MicrosoftGraphMessage) HasConversationId() bool`

HasConversationId returns a boolean if a field has been set.

### SetConversationIdNil

`func (o *MicrosoftGraphMessage) SetConversationIdNil(b bool)`

 SetConversationIdNil sets the value for ConversationId to be an explicit nil

### UnsetConversationId
`func (o *MicrosoftGraphMessage) UnsetConversationId()`

UnsetConversationId ensures that no value is present for ConversationId, not even an explicit nil
### GetConversationIndex

`func (o *MicrosoftGraphMessage) GetConversationIndex() string`

GetConversationIndex returns the ConversationIndex field if non-nil, zero value otherwise.

### GetConversationIndexOk

`func (o *MicrosoftGraphMessage) GetConversationIndexOk() (*string, bool)`

GetConversationIndexOk returns a tuple with the ConversationIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversationIndex

`func (o *MicrosoftGraphMessage) SetConversationIndex(v string)`

SetConversationIndex sets ConversationIndex field to given value.

### HasConversationIndex

`func (o *MicrosoftGraphMessage) HasConversationIndex() bool`

HasConversationIndex returns a boolean if a field has been set.

### SetConversationIndexNil

`func (o *MicrosoftGraphMessage) SetConversationIndexNil(b bool)`

 SetConversationIndexNil sets the value for ConversationIndex to be an explicit nil

### UnsetConversationIndex
`func (o *MicrosoftGraphMessage) UnsetConversationIndex()`

UnsetConversationIndex ensures that no value is present for ConversationIndex, not even an explicit nil
### GetFlag

`func (o *MicrosoftGraphMessage) GetFlag() AnyOfmicrosoftGraphFollowupFlag`

GetFlag returns the Flag field if non-nil, zero value otherwise.

### GetFlagOk

`func (o *MicrosoftGraphMessage) GetFlagOk() (*AnyOfmicrosoftGraphFollowupFlag, bool)`

GetFlagOk returns a tuple with the Flag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlag

`func (o *MicrosoftGraphMessage) SetFlag(v AnyOfmicrosoftGraphFollowupFlag)`

SetFlag sets Flag field to given value.

### HasFlag

`func (o *MicrosoftGraphMessage) HasFlag() bool`

HasFlag returns a boolean if a field has been set.

### SetFlagNil

`func (o *MicrosoftGraphMessage) SetFlagNil(b bool)`

 SetFlagNil sets the value for Flag to be an explicit nil

### UnsetFlag
`func (o *MicrosoftGraphMessage) UnsetFlag()`

UnsetFlag ensures that no value is present for Flag, not even an explicit nil
### GetFrom

`func (o *MicrosoftGraphMessage) GetFrom() AnyOfmicrosoftGraphRecipient`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *MicrosoftGraphMessage) GetFromOk() (*AnyOfmicrosoftGraphRecipient, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *MicrosoftGraphMessage) SetFrom(v AnyOfmicrosoftGraphRecipient)`

SetFrom sets From field to given value.

### HasFrom

`func (o *MicrosoftGraphMessage) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### SetFromNil

`func (o *MicrosoftGraphMessage) SetFromNil(b bool)`

 SetFromNil sets the value for From to be an explicit nil

### UnsetFrom
`func (o *MicrosoftGraphMessage) UnsetFrom()`

UnsetFrom ensures that no value is present for From, not even an explicit nil
### GetHasAttachments

`func (o *MicrosoftGraphMessage) GetHasAttachments() bool`

GetHasAttachments returns the HasAttachments field if non-nil, zero value otherwise.

### GetHasAttachmentsOk

`func (o *MicrosoftGraphMessage) GetHasAttachmentsOk() (*bool, bool)`

GetHasAttachmentsOk returns a tuple with the HasAttachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAttachments

`func (o *MicrosoftGraphMessage) SetHasAttachments(v bool)`

SetHasAttachments sets HasAttachments field to given value.

### HasHasAttachments

`func (o *MicrosoftGraphMessage) HasHasAttachments() bool`

HasHasAttachments returns a boolean if a field has been set.

### SetHasAttachmentsNil

`func (o *MicrosoftGraphMessage) SetHasAttachmentsNil(b bool)`

 SetHasAttachmentsNil sets the value for HasAttachments to be an explicit nil

### UnsetHasAttachments
`func (o *MicrosoftGraphMessage) UnsetHasAttachments()`

UnsetHasAttachments ensures that no value is present for HasAttachments, not even an explicit nil
### GetImportance

`func (o *MicrosoftGraphMessage) GetImportance() AnyOfmicrosoftGraphImportance`

GetImportance returns the Importance field if non-nil, zero value otherwise.

### GetImportanceOk

`func (o *MicrosoftGraphMessage) GetImportanceOk() (*AnyOfmicrosoftGraphImportance, bool)`

GetImportanceOk returns a tuple with the Importance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportance

`func (o *MicrosoftGraphMessage) SetImportance(v AnyOfmicrosoftGraphImportance)`

SetImportance sets Importance field to given value.

### HasImportance

`func (o *MicrosoftGraphMessage) HasImportance() bool`

HasImportance returns a boolean if a field has been set.

### SetImportanceNil

`func (o *MicrosoftGraphMessage) SetImportanceNil(b bool)`

 SetImportanceNil sets the value for Importance to be an explicit nil

### UnsetImportance
`func (o *MicrosoftGraphMessage) UnsetImportance()`

UnsetImportance ensures that no value is present for Importance, not even an explicit nil
### GetInferenceClassification

`func (o *MicrosoftGraphMessage) GetInferenceClassification() AnyOfmicrosoftGraphInferenceClassificationType`

GetInferenceClassification returns the InferenceClassification field if non-nil, zero value otherwise.

### GetInferenceClassificationOk

`func (o *MicrosoftGraphMessage) GetInferenceClassificationOk() (*AnyOfmicrosoftGraphInferenceClassificationType, bool)`

GetInferenceClassificationOk returns a tuple with the InferenceClassification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInferenceClassification

`func (o *MicrosoftGraphMessage) SetInferenceClassification(v AnyOfmicrosoftGraphInferenceClassificationType)`

SetInferenceClassification sets InferenceClassification field to given value.

### HasInferenceClassification

`func (o *MicrosoftGraphMessage) HasInferenceClassification() bool`

HasInferenceClassification returns a boolean if a field has been set.

### SetInferenceClassificationNil

`func (o *MicrosoftGraphMessage) SetInferenceClassificationNil(b bool)`

 SetInferenceClassificationNil sets the value for InferenceClassification to be an explicit nil

### UnsetInferenceClassification
`func (o *MicrosoftGraphMessage) UnsetInferenceClassification()`

UnsetInferenceClassification ensures that no value is present for InferenceClassification, not even an explicit nil
### GetInternetMessageHeaders

`func (o *MicrosoftGraphMessage) GetInternetMessageHeaders() []*AnyOfmicrosoftGraphInternetMessageHeader`

GetInternetMessageHeaders returns the InternetMessageHeaders field if non-nil, zero value otherwise.

### GetInternetMessageHeadersOk

`func (o *MicrosoftGraphMessage) GetInternetMessageHeadersOk() (*[]*AnyOfmicrosoftGraphInternetMessageHeader, bool)`

GetInternetMessageHeadersOk returns a tuple with the InternetMessageHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternetMessageHeaders

`func (o *MicrosoftGraphMessage) SetInternetMessageHeaders(v []*AnyOfmicrosoftGraphInternetMessageHeader)`

SetInternetMessageHeaders sets InternetMessageHeaders field to given value.

### HasInternetMessageHeaders

`func (o *MicrosoftGraphMessage) HasInternetMessageHeaders() bool`

HasInternetMessageHeaders returns a boolean if a field has been set.

### GetInternetMessageId

`func (o *MicrosoftGraphMessage) GetInternetMessageId() string`

GetInternetMessageId returns the InternetMessageId field if non-nil, zero value otherwise.

### GetInternetMessageIdOk

`func (o *MicrosoftGraphMessage) GetInternetMessageIdOk() (*string, bool)`

GetInternetMessageIdOk returns a tuple with the InternetMessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternetMessageId

`func (o *MicrosoftGraphMessage) SetInternetMessageId(v string)`

SetInternetMessageId sets InternetMessageId field to given value.

### HasInternetMessageId

`func (o *MicrosoftGraphMessage) HasInternetMessageId() bool`

HasInternetMessageId returns a boolean if a field has been set.

### SetInternetMessageIdNil

`func (o *MicrosoftGraphMessage) SetInternetMessageIdNil(b bool)`

 SetInternetMessageIdNil sets the value for InternetMessageId to be an explicit nil

### UnsetInternetMessageId
`func (o *MicrosoftGraphMessage) UnsetInternetMessageId()`

UnsetInternetMessageId ensures that no value is present for InternetMessageId, not even an explicit nil
### GetIsDeliveryReceiptRequested

`func (o *MicrosoftGraphMessage) GetIsDeliveryReceiptRequested() bool`

GetIsDeliveryReceiptRequested returns the IsDeliveryReceiptRequested field if non-nil, zero value otherwise.

### GetIsDeliveryReceiptRequestedOk

`func (o *MicrosoftGraphMessage) GetIsDeliveryReceiptRequestedOk() (*bool, bool)`

GetIsDeliveryReceiptRequestedOk returns a tuple with the IsDeliveryReceiptRequested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeliveryReceiptRequested

`func (o *MicrosoftGraphMessage) SetIsDeliveryReceiptRequested(v bool)`

SetIsDeliveryReceiptRequested sets IsDeliveryReceiptRequested field to given value.

### HasIsDeliveryReceiptRequested

`func (o *MicrosoftGraphMessage) HasIsDeliveryReceiptRequested() bool`

HasIsDeliveryReceiptRequested returns a boolean if a field has been set.

### SetIsDeliveryReceiptRequestedNil

`func (o *MicrosoftGraphMessage) SetIsDeliveryReceiptRequestedNil(b bool)`

 SetIsDeliveryReceiptRequestedNil sets the value for IsDeliveryReceiptRequested to be an explicit nil

### UnsetIsDeliveryReceiptRequested
`func (o *MicrosoftGraphMessage) UnsetIsDeliveryReceiptRequested()`

UnsetIsDeliveryReceiptRequested ensures that no value is present for IsDeliveryReceiptRequested, not even an explicit nil
### GetIsDraft

`func (o *MicrosoftGraphMessage) GetIsDraft() bool`

GetIsDraft returns the IsDraft field if non-nil, zero value otherwise.

### GetIsDraftOk

`func (o *MicrosoftGraphMessage) GetIsDraftOk() (*bool, bool)`

GetIsDraftOk returns a tuple with the IsDraft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDraft

`func (o *MicrosoftGraphMessage) SetIsDraft(v bool)`

SetIsDraft sets IsDraft field to given value.

### HasIsDraft

`func (o *MicrosoftGraphMessage) HasIsDraft() bool`

HasIsDraft returns a boolean if a field has been set.

### SetIsDraftNil

`func (o *MicrosoftGraphMessage) SetIsDraftNil(b bool)`

 SetIsDraftNil sets the value for IsDraft to be an explicit nil

### UnsetIsDraft
`func (o *MicrosoftGraphMessage) UnsetIsDraft()`

UnsetIsDraft ensures that no value is present for IsDraft, not even an explicit nil
### GetIsRead

`func (o *MicrosoftGraphMessage) GetIsRead() bool`

GetIsRead returns the IsRead field if non-nil, zero value otherwise.

### GetIsReadOk

`func (o *MicrosoftGraphMessage) GetIsReadOk() (*bool, bool)`

GetIsReadOk returns a tuple with the IsRead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRead

`func (o *MicrosoftGraphMessage) SetIsRead(v bool)`

SetIsRead sets IsRead field to given value.

### HasIsRead

`func (o *MicrosoftGraphMessage) HasIsRead() bool`

HasIsRead returns a boolean if a field has been set.

### SetIsReadNil

`func (o *MicrosoftGraphMessage) SetIsReadNil(b bool)`

 SetIsReadNil sets the value for IsRead to be an explicit nil

### UnsetIsRead
`func (o *MicrosoftGraphMessage) UnsetIsRead()`

UnsetIsRead ensures that no value is present for IsRead, not even an explicit nil
### GetIsReadReceiptRequested

`func (o *MicrosoftGraphMessage) GetIsReadReceiptRequested() bool`

GetIsReadReceiptRequested returns the IsReadReceiptRequested field if non-nil, zero value otherwise.

### GetIsReadReceiptRequestedOk

`func (o *MicrosoftGraphMessage) GetIsReadReceiptRequestedOk() (*bool, bool)`

GetIsReadReceiptRequestedOk returns a tuple with the IsReadReceiptRequested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReadReceiptRequested

`func (o *MicrosoftGraphMessage) SetIsReadReceiptRequested(v bool)`

SetIsReadReceiptRequested sets IsReadReceiptRequested field to given value.

### HasIsReadReceiptRequested

`func (o *MicrosoftGraphMessage) HasIsReadReceiptRequested() bool`

HasIsReadReceiptRequested returns a boolean if a field has been set.

### SetIsReadReceiptRequestedNil

`func (o *MicrosoftGraphMessage) SetIsReadReceiptRequestedNil(b bool)`

 SetIsReadReceiptRequestedNil sets the value for IsReadReceiptRequested to be an explicit nil

### UnsetIsReadReceiptRequested
`func (o *MicrosoftGraphMessage) UnsetIsReadReceiptRequested()`

UnsetIsReadReceiptRequested ensures that no value is present for IsReadReceiptRequested, not even an explicit nil
### GetParentFolderId

`func (o *MicrosoftGraphMessage) GetParentFolderId() string`

GetParentFolderId returns the ParentFolderId field if non-nil, zero value otherwise.

### GetParentFolderIdOk

`func (o *MicrosoftGraphMessage) GetParentFolderIdOk() (*string, bool)`

GetParentFolderIdOk returns a tuple with the ParentFolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentFolderId

`func (o *MicrosoftGraphMessage) SetParentFolderId(v string)`

SetParentFolderId sets ParentFolderId field to given value.

### HasParentFolderId

`func (o *MicrosoftGraphMessage) HasParentFolderId() bool`

HasParentFolderId returns a boolean if a field has been set.

### SetParentFolderIdNil

`func (o *MicrosoftGraphMessage) SetParentFolderIdNil(b bool)`

 SetParentFolderIdNil sets the value for ParentFolderId to be an explicit nil

### UnsetParentFolderId
`func (o *MicrosoftGraphMessage) UnsetParentFolderId()`

UnsetParentFolderId ensures that no value is present for ParentFolderId, not even an explicit nil
### GetReceivedDateTime

`func (o *MicrosoftGraphMessage) GetReceivedDateTime() time.Time`

GetReceivedDateTime returns the ReceivedDateTime field if non-nil, zero value otherwise.

### GetReceivedDateTimeOk

`func (o *MicrosoftGraphMessage) GetReceivedDateTimeOk() (*time.Time, bool)`

GetReceivedDateTimeOk returns a tuple with the ReceivedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivedDateTime

`func (o *MicrosoftGraphMessage) SetReceivedDateTime(v time.Time)`

SetReceivedDateTime sets ReceivedDateTime field to given value.

### HasReceivedDateTime

`func (o *MicrosoftGraphMessage) HasReceivedDateTime() bool`

HasReceivedDateTime returns a boolean if a field has been set.

### SetReceivedDateTimeNil

`func (o *MicrosoftGraphMessage) SetReceivedDateTimeNil(b bool)`

 SetReceivedDateTimeNil sets the value for ReceivedDateTime to be an explicit nil

### UnsetReceivedDateTime
`func (o *MicrosoftGraphMessage) UnsetReceivedDateTime()`

UnsetReceivedDateTime ensures that no value is present for ReceivedDateTime, not even an explicit nil
### GetReplyTo

`func (o *MicrosoftGraphMessage) GetReplyTo() []*AnyOfmicrosoftGraphRecipient`

GetReplyTo returns the ReplyTo field if non-nil, zero value otherwise.

### GetReplyToOk

`func (o *MicrosoftGraphMessage) GetReplyToOk() (*[]*AnyOfmicrosoftGraphRecipient, bool)`

GetReplyToOk returns a tuple with the ReplyTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyTo

`func (o *MicrosoftGraphMessage) SetReplyTo(v []*AnyOfmicrosoftGraphRecipient)`

SetReplyTo sets ReplyTo field to given value.

### HasReplyTo

`func (o *MicrosoftGraphMessage) HasReplyTo() bool`

HasReplyTo returns a boolean if a field has been set.

### GetSender

`func (o *MicrosoftGraphMessage) GetSender() AnyOfmicrosoftGraphRecipient`

GetSender returns the Sender field if non-nil, zero value otherwise.

### GetSenderOk

`func (o *MicrosoftGraphMessage) GetSenderOk() (*AnyOfmicrosoftGraphRecipient, bool)`

GetSenderOk returns a tuple with the Sender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSender

`func (o *MicrosoftGraphMessage) SetSender(v AnyOfmicrosoftGraphRecipient)`

SetSender sets Sender field to given value.

### HasSender

`func (o *MicrosoftGraphMessage) HasSender() bool`

HasSender returns a boolean if a field has been set.

### SetSenderNil

`func (o *MicrosoftGraphMessage) SetSenderNil(b bool)`

 SetSenderNil sets the value for Sender to be an explicit nil

### UnsetSender
`func (o *MicrosoftGraphMessage) UnsetSender()`

UnsetSender ensures that no value is present for Sender, not even an explicit nil
### GetSentDateTime

`func (o *MicrosoftGraphMessage) GetSentDateTime() time.Time`

GetSentDateTime returns the SentDateTime field if non-nil, zero value otherwise.

### GetSentDateTimeOk

`func (o *MicrosoftGraphMessage) GetSentDateTimeOk() (*time.Time, bool)`

GetSentDateTimeOk returns a tuple with the SentDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSentDateTime

`func (o *MicrosoftGraphMessage) SetSentDateTime(v time.Time)`

SetSentDateTime sets SentDateTime field to given value.

### HasSentDateTime

`func (o *MicrosoftGraphMessage) HasSentDateTime() bool`

HasSentDateTime returns a boolean if a field has been set.

### SetSentDateTimeNil

`func (o *MicrosoftGraphMessage) SetSentDateTimeNil(b bool)`

 SetSentDateTimeNil sets the value for SentDateTime to be an explicit nil

### UnsetSentDateTime
`func (o *MicrosoftGraphMessage) UnsetSentDateTime()`

UnsetSentDateTime ensures that no value is present for SentDateTime, not even an explicit nil
### GetSubject

`func (o *MicrosoftGraphMessage) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *MicrosoftGraphMessage) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *MicrosoftGraphMessage) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *MicrosoftGraphMessage) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### SetSubjectNil

`func (o *MicrosoftGraphMessage) SetSubjectNil(b bool)`

 SetSubjectNil sets the value for Subject to be an explicit nil

### UnsetSubject
`func (o *MicrosoftGraphMessage) UnsetSubject()`

UnsetSubject ensures that no value is present for Subject, not even an explicit nil
### GetToRecipients

`func (o *MicrosoftGraphMessage) GetToRecipients() []*AnyOfmicrosoftGraphRecipient`

GetToRecipients returns the ToRecipients field if non-nil, zero value otherwise.

### GetToRecipientsOk

`func (o *MicrosoftGraphMessage) GetToRecipientsOk() (*[]*AnyOfmicrosoftGraphRecipient, bool)`

GetToRecipientsOk returns a tuple with the ToRecipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToRecipients

`func (o *MicrosoftGraphMessage) SetToRecipients(v []*AnyOfmicrosoftGraphRecipient)`

SetToRecipients sets ToRecipients field to given value.

### HasToRecipients

`func (o *MicrosoftGraphMessage) HasToRecipients() bool`

HasToRecipients returns a boolean if a field has been set.

### GetUniqueBody

`func (o *MicrosoftGraphMessage) GetUniqueBody() AnyOfmicrosoftGraphItemBody`

GetUniqueBody returns the UniqueBody field if non-nil, zero value otherwise.

### GetUniqueBodyOk

`func (o *MicrosoftGraphMessage) GetUniqueBodyOk() (*AnyOfmicrosoftGraphItemBody, bool)`

GetUniqueBodyOk returns a tuple with the UniqueBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueBody

`func (o *MicrosoftGraphMessage) SetUniqueBody(v AnyOfmicrosoftGraphItemBody)`

SetUniqueBody sets UniqueBody field to given value.

### HasUniqueBody

`func (o *MicrosoftGraphMessage) HasUniqueBody() bool`

HasUniqueBody returns a boolean if a field has been set.

### SetUniqueBodyNil

`func (o *MicrosoftGraphMessage) SetUniqueBodyNil(b bool)`

 SetUniqueBodyNil sets the value for UniqueBody to be an explicit nil

### UnsetUniqueBody
`func (o *MicrosoftGraphMessage) UnsetUniqueBody()`

UnsetUniqueBody ensures that no value is present for UniqueBody, not even an explicit nil
### GetWebLink

`func (o *MicrosoftGraphMessage) GetWebLink() string`

GetWebLink returns the WebLink field if non-nil, zero value otherwise.

### GetWebLinkOk

`func (o *MicrosoftGraphMessage) GetWebLinkOk() (*string, bool)`

GetWebLinkOk returns a tuple with the WebLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebLink

`func (o *MicrosoftGraphMessage) SetWebLink(v string)`

SetWebLink sets WebLink field to given value.

### HasWebLink

`func (o *MicrosoftGraphMessage) HasWebLink() bool`

HasWebLink returns a boolean if a field has been set.

### SetWebLinkNil

`func (o *MicrosoftGraphMessage) SetWebLinkNil(b bool)`

 SetWebLinkNil sets the value for WebLink to be an explicit nil

### UnsetWebLink
`func (o *MicrosoftGraphMessage) UnsetWebLink()`

UnsetWebLink ensures that no value is present for WebLink, not even an explicit nil
### GetAttachments

`func (o *MicrosoftGraphMessage) GetAttachments() []MicrosoftGraphAttachment`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *MicrosoftGraphMessage) GetAttachmentsOk() (*[]MicrosoftGraphAttachment, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *MicrosoftGraphMessage) SetAttachments(v []MicrosoftGraphAttachment)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *MicrosoftGraphMessage) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetExtensions

`func (o *MicrosoftGraphMessage) GetExtensions() []MicrosoftGraphExtension`

GetExtensions returns the Extensions field if non-nil, zero value otherwise.

### GetExtensionsOk

`func (o *MicrosoftGraphMessage) GetExtensionsOk() (*[]MicrosoftGraphExtension, bool)`

GetExtensionsOk returns a tuple with the Extensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensions

`func (o *MicrosoftGraphMessage) SetExtensions(v []MicrosoftGraphExtension)`

SetExtensions sets Extensions field to given value.

### HasExtensions

`func (o *MicrosoftGraphMessage) HasExtensions() bool`

HasExtensions returns a boolean if a field has been set.

### GetMultiValueExtendedProperties

`func (o *MicrosoftGraphMessage) GetMultiValueExtendedProperties() []MicrosoftGraphMultiValueLegacyExtendedProperty`

GetMultiValueExtendedProperties returns the MultiValueExtendedProperties field if non-nil, zero value otherwise.

### GetMultiValueExtendedPropertiesOk

`func (o *MicrosoftGraphMessage) GetMultiValueExtendedPropertiesOk() (*[]MicrosoftGraphMultiValueLegacyExtendedProperty, bool)`

GetMultiValueExtendedPropertiesOk returns a tuple with the MultiValueExtendedProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiValueExtendedProperties

`func (o *MicrosoftGraphMessage) SetMultiValueExtendedProperties(v []MicrosoftGraphMultiValueLegacyExtendedProperty)`

SetMultiValueExtendedProperties sets MultiValueExtendedProperties field to given value.

### HasMultiValueExtendedProperties

`func (o *MicrosoftGraphMessage) HasMultiValueExtendedProperties() bool`

HasMultiValueExtendedProperties returns a boolean if a field has been set.

### GetSingleValueExtendedProperties

`func (o *MicrosoftGraphMessage) GetSingleValueExtendedProperties() []MicrosoftGraphSingleValueLegacyExtendedProperty`

GetSingleValueExtendedProperties returns the SingleValueExtendedProperties field if non-nil, zero value otherwise.

### GetSingleValueExtendedPropertiesOk

`func (o *MicrosoftGraphMessage) GetSingleValueExtendedPropertiesOk() (*[]MicrosoftGraphSingleValueLegacyExtendedProperty, bool)`

GetSingleValueExtendedPropertiesOk returns a tuple with the SingleValueExtendedProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSingleValueExtendedProperties

`func (o *MicrosoftGraphMessage) SetSingleValueExtendedProperties(v []MicrosoftGraphSingleValueLegacyExtendedProperty)`

SetSingleValueExtendedProperties sets SingleValueExtendedProperties field to given value.

### HasSingleValueExtendedProperties

`func (o *MicrosoftGraphMessage) HasSingleValueExtendedProperties() bool`

HasSingleValueExtendedProperties returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


