# MicrosoftGraphMessageRulePredicates

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BodyContains** | Pointer to **[]string** | Represents the strings that should appear in the body of an incoming message in order for the condition or exception to apply. | [optional] 
**BodyOrSubjectContains** | Pointer to **[]string** | Represents the strings that should appear in the body or subject of an incoming message in order for the condition or exception to apply. | [optional] 
**Categories** | Pointer to **[]string** | Represents the categories that an incoming message should be labeled with in order for the condition or exception to apply. | [optional] 
**FromAddresses** | Pointer to [**[]AnyOfmicrosoftGraphRecipient**](AnyOfmicrosoftGraphRecipient.md) | Represents the specific sender email addresses of an incoming message in order for the condition or exception to apply. | [optional] 
**HasAttachments** | Pointer to **NullableBool** | Indicates whether an incoming message must have attachments in order for the condition or exception to apply. | [optional] 
**HeaderContains** | Pointer to **[]string** | Represents the strings that appear in the headers of an incoming message in order for the condition or exception to apply. | [optional] 
**Importance** | Pointer to [**AnyOfmicrosoftGraphImportance**](anyOf&lt;microsoft.graph.importance&gt;.md) | The importance that is stamped on an incoming message in order for the condition or exception to apply: low, normal, high. | [optional] 
**IsApprovalRequest** | Pointer to **NullableBool** | Indicates whether an incoming message must be an approval request in order for the condition or exception to apply. | [optional] 
**IsAutomaticForward** | Pointer to **NullableBool** | Indicates whether an incoming message must be automatically forwarded in order for the condition or exception to apply. | [optional] 
**IsAutomaticReply** | Pointer to **NullableBool** | Indicates whether an incoming message must be an auto reply in order for the condition or exception to apply. | [optional] 
**IsEncrypted** | Pointer to **NullableBool** | Indicates whether an incoming message must be encrypted in order for the condition or exception to apply. | [optional] 
**IsMeetingRequest** | Pointer to **NullableBool** | Indicates whether an incoming message must be a meeting request in order for the condition or exception to apply. | [optional] 
**IsMeetingResponse** | Pointer to **NullableBool** | Indicates whether an incoming message must be a meeting response in order for the condition or exception to apply. | [optional] 
**IsNonDeliveryReport** | Pointer to **NullableBool** | Indicates whether an incoming message must be a non-delivery report in order for the condition or exception to apply. | [optional] 
**IsPermissionControlled** | Pointer to **NullableBool** | Indicates whether an incoming message must be permission controlled (RMS-protected) in order for the condition or exception to apply. | [optional] 
**IsReadReceipt** | Pointer to **NullableBool** | Indicates whether an incoming message must be a read receipt in order for the condition or exception to apply. | [optional] 
**IsSigned** | Pointer to **NullableBool** | Indicates whether an incoming message must be S/MIME-signed in order for the condition or exception to apply. | [optional] 
**IsVoicemail** | Pointer to **NullableBool** | Indicates whether an incoming message must be a voice mail in order for the condition or exception to apply. | [optional] 
**MessageActionFlag** | Pointer to [**AnyOfmicrosoftGraphMessageActionFlag**](anyOf&lt;microsoft.graph.messageActionFlag&gt;.md) | Represents the flag-for-action value that appears on an incoming message in order for the condition or exception to apply. The possible values are: any, call, doNotForward, followUp, fyi, forward, noResponseNecessary, read, reply, replyToAll, review. | [optional] 
**NotSentToMe** | Pointer to **NullableBool** | Indicates whether the owner of the mailbox must not be a recipient of an incoming message in order for the condition or exception to apply. | [optional] 
**RecipientContains** | Pointer to **[]string** | Represents the strings that appear in either the toRecipients or ccRecipients properties of an incoming message in order for the condition or exception to apply. | [optional] 
**SenderContains** | Pointer to **[]string** | Represents the strings that appear in the from property of an incoming message in order for the condition or exception to apply. | [optional] 
**Sensitivity** | Pointer to [**AnyOfmicrosoftGraphSensitivity**](anyOf&lt;microsoft.graph.sensitivity&gt;.md) | Represents the sensitivity level that must be stamped on an incoming message in order for the condition or exception to apply. The possible values are: normal, personal, private, confidential. | [optional] 
**SentCcMe** | Pointer to **NullableBool** | Indicates whether the owner of the mailbox must be in the ccRecipients property of an incoming message in order for the condition or exception to apply. | [optional] 
**SentOnlyToMe** | Pointer to **NullableBool** | Indicates whether the owner of the mailbox must be the only recipient in an incoming message in order for the condition or exception to apply. | [optional] 
**SentToAddresses** | Pointer to [**[]AnyOfmicrosoftGraphRecipient**](AnyOfmicrosoftGraphRecipient.md) | Represents the email addresses that an incoming message must have been sent to in order for the condition or exception to apply. | [optional] 
**SentToMe** | Pointer to **NullableBool** | Indicates whether the owner of the mailbox must be in the toRecipients property of an incoming message in order for the condition or exception to apply. | [optional] 
**SentToOrCcMe** | Pointer to **NullableBool** | Indicates whether the owner of the mailbox must be in either a toRecipients or ccRecipients property of an incoming message in order for the condition or exception to apply. | [optional] 
**SubjectContains** | Pointer to **[]string** | Represents the strings that appear in the subject of an incoming message in order for the condition or exception to apply. | [optional] 
**WithinSizeRange** | Pointer to [**AnyOfmicrosoftGraphSizeRange**](anyOf&lt;microsoft.graph.sizeRange&gt;.md) | Represents the minimum and maximum sizes (in kilobytes) that an incoming message must fall in between in order for the condition or exception to apply. | [optional] 

## Methods

### NewMicrosoftGraphMessageRulePredicates

`func NewMicrosoftGraphMessageRulePredicates() *MicrosoftGraphMessageRulePredicates`

NewMicrosoftGraphMessageRulePredicates instantiates a new MicrosoftGraphMessageRulePredicates object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphMessageRulePredicatesWithDefaults

`func NewMicrosoftGraphMessageRulePredicatesWithDefaults() *MicrosoftGraphMessageRulePredicates`

NewMicrosoftGraphMessageRulePredicatesWithDefaults instantiates a new MicrosoftGraphMessageRulePredicates object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBodyContains

`func (o *MicrosoftGraphMessageRulePredicates) GetBodyContains() []*string`

GetBodyContains returns the BodyContains field if non-nil, zero value otherwise.

### GetBodyContainsOk

`func (o *MicrosoftGraphMessageRulePredicates) GetBodyContainsOk() (*[]*string, bool)`

GetBodyContainsOk returns a tuple with the BodyContains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBodyContains

`func (o *MicrosoftGraphMessageRulePredicates) SetBodyContains(v []*string)`

SetBodyContains sets BodyContains field to given value.

### HasBodyContains

`func (o *MicrosoftGraphMessageRulePredicates) HasBodyContains() bool`

HasBodyContains returns a boolean if a field has been set.

### GetBodyOrSubjectContains

`func (o *MicrosoftGraphMessageRulePredicates) GetBodyOrSubjectContains() []*string`

GetBodyOrSubjectContains returns the BodyOrSubjectContains field if non-nil, zero value otherwise.

### GetBodyOrSubjectContainsOk

`func (o *MicrosoftGraphMessageRulePredicates) GetBodyOrSubjectContainsOk() (*[]*string, bool)`

GetBodyOrSubjectContainsOk returns a tuple with the BodyOrSubjectContains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBodyOrSubjectContains

`func (o *MicrosoftGraphMessageRulePredicates) SetBodyOrSubjectContains(v []*string)`

SetBodyOrSubjectContains sets BodyOrSubjectContains field to given value.

### HasBodyOrSubjectContains

`func (o *MicrosoftGraphMessageRulePredicates) HasBodyOrSubjectContains() bool`

HasBodyOrSubjectContains returns a boolean if a field has been set.

### GetCategories

`func (o *MicrosoftGraphMessageRulePredicates) GetCategories() []*string`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *MicrosoftGraphMessageRulePredicates) GetCategoriesOk() (*[]*string, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *MicrosoftGraphMessageRulePredicates) SetCategories(v []*string)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *MicrosoftGraphMessageRulePredicates) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### GetFromAddresses

`func (o *MicrosoftGraphMessageRulePredicates) GetFromAddresses() []*AnyOfmicrosoftGraphRecipient`

GetFromAddresses returns the FromAddresses field if non-nil, zero value otherwise.

### GetFromAddressesOk

`func (o *MicrosoftGraphMessageRulePredicates) GetFromAddressesOk() (*[]*AnyOfmicrosoftGraphRecipient, bool)`

GetFromAddressesOk returns a tuple with the FromAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromAddresses

`func (o *MicrosoftGraphMessageRulePredicates) SetFromAddresses(v []*AnyOfmicrosoftGraphRecipient)`

SetFromAddresses sets FromAddresses field to given value.

### HasFromAddresses

`func (o *MicrosoftGraphMessageRulePredicates) HasFromAddresses() bool`

HasFromAddresses returns a boolean if a field has been set.

### GetHasAttachments

`func (o *MicrosoftGraphMessageRulePredicates) GetHasAttachments() bool`

GetHasAttachments returns the HasAttachments field if non-nil, zero value otherwise.

### GetHasAttachmentsOk

`func (o *MicrosoftGraphMessageRulePredicates) GetHasAttachmentsOk() (*bool, bool)`

GetHasAttachmentsOk returns a tuple with the HasAttachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAttachments

`func (o *MicrosoftGraphMessageRulePredicates) SetHasAttachments(v bool)`

SetHasAttachments sets HasAttachments field to given value.

### HasHasAttachments

`func (o *MicrosoftGraphMessageRulePredicates) HasHasAttachments() bool`

HasHasAttachments returns a boolean if a field has been set.

### SetHasAttachmentsNil

`func (o *MicrosoftGraphMessageRulePredicates) SetHasAttachmentsNil(b bool)`

 SetHasAttachmentsNil sets the value for HasAttachments to be an explicit nil

### UnsetHasAttachments
`func (o *MicrosoftGraphMessageRulePredicates) UnsetHasAttachments()`

UnsetHasAttachments ensures that no value is present for HasAttachments, not even an explicit nil
### GetHeaderContains

`func (o *MicrosoftGraphMessageRulePredicates) GetHeaderContains() []*string`

GetHeaderContains returns the HeaderContains field if non-nil, zero value otherwise.

### GetHeaderContainsOk

`func (o *MicrosoftGraphMessageRulePredicates) GetHeaderContainsOk() (*[]*string, bool)`

GetHeaderContainsOk returns a tuple with the HeaderContains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderContains

`func (o *MicrosoftGraphMessageRulePredicates) SetHeaderContains(v []*string)`

SetHeaderContains sets HeaderContains field to given value.

### HasHeaderContains

`func (o *MicrosoftGraphMessageRulePredicates) HasHeaderContains() bool`

HasHeaderContains returns a boolean if a field has been set.

### GetImportance

`func (o *MicrosoftGraphMessageRulePredicates) GetImportance() AnyOfmicrosoftGraphImportance`

GetImportance returns the Importance field if non-nil, zero value otherwise.

### GetImportanceOk

`func (o *MicrosoftGraphMessageRulePredicates) GetImportanceOk() (*AnyOfmicrosoftGraphImportance, bool)`

GetImportanceOk returns a tuple with the Importance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportance

`func (o *MicrosoftGraphMessageRulePredicates) SetImportance(v AnyOfmicrosoftGraphImportance)`

SetImportance sets Importance field to given value.

### HasImportance

`func (o *MicrosoftGraphMessageRulePredicates) HasImportance() bool`

HasImportance returns a boolean if a field has been set.

### SetImportanceNil

`func (o *MicrosoftGraphMessageRulePredicates) SetImportanceNil(b bool)`

 SetImportanceNil sets the value for Importance to be an explicit nil

### UnsetImportance
`func (o *MicrosoftGraphMessageRulePredicates) UnsetImportance()`

UnsetImportance ensures that no value is present for Importance, not even an explicit nil
### GetIsApprovalRequest

`func (o *MicrosoftGraphMessageRulePredicates) GetIsApprovalRequest() bool`

GetIsApprovalRequest returns the IsApprovalRequest field if non-nil, zero value otherwise.

### GetIsApprovalRequestOk

`func (o *MicrosoftGraphMessageRulePredicates) GetIsApprovalRequestOk() (*bool, bool)`

GetIsApprovalRequestOk returns a tuple with the IsApprovalRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsApprovalRequest

`func (o *MicrosoftGraphMessageRulePredicates) SetIsApprovalRequest(v bool)`

SetIsApprovalRequest sets IsApprovalRequest field to given value.

### HasIsApprovalRequest

`func (o *MicrosoftGraphMessageRulePredicates) HasIsApprovalRequest() bool`

HasIsApprovalRequest returns a boolean if a field has been set.

### SetIsApprovalRequestNil

`func (o *MicrosoftGraphMessageRulePredicates) SetIsApprovalRequestNil(b bool)`

 SetIsApprovalRequestNil sets the value for IsApprovalRequest to be an explicit nil

### UnsetIsApprovalRequest
`func (o *MicrosoftGraphMessageRulePredicates) UnsetIsApprovalRequest()`

UnsetIsApprovalRequest ensures that no value is present for IsApprovalRequest, not even an explicit nil
### GetIsAutomaticForward

`func (o *MicrosoftGraphMessageRulePredicates) GetIsAutomaticForward() bool`

GetIsAutomaticForward returns the IsAutomaticForward field if non-nil, zero value otherwise.

### GetIsAutomaticForwardOk

`func (o *MicrosoftGraphMessageRulePredicates) GetIsAutomaticForwardOk() (*bool, bool)`

GetIsAutomaticForwardOk returns a tuple with the IsAutomaticForward field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutomaticForward

`func (o *MicrosoftGraphMessageRulePredicates) SetIsAutomaticForward(v bool)`

SetIsAutomaticForward sets IsAutomaticForward field to given value.

### HasIsAutomaticForward

`func (o *MicrosoftGraphMessageRulePredicates) HasIsAutomaticForward() bool`

HasIsAutomaticForward returns a boolean if a field has been set.

### SetIsAutomaticForwardNil

`func (o *MicrosoftGraphMessageRulePredicates) SetIsAutomaticForwardNil(b bool)`

 SetIsAutomaticForwardNil sets the value for IsAutomaticForward to be an explicit nil

### UnsetIsAutomaticForward
`func (o *MicrosoftGraphMessageRulePredicates) UnsetIsAutomaticForward()`

UnsetIsAutomaticForward ensures that no value is present for IsAutomaticForward, not even an explicit nil
### GetIsAutomaticReply

`func (o *MicrosoftGraphMessageRulePredicates) GetIsAutomaticReply() bool`

GetIsAutomaticReply returns the IsAutomaticReply field if non-nil, zero value otherwise.

### GetIsAutomaticReplyOk

`func (o *MicrosoftGraphMessageRulePredicates) GetIsAutomaticReplyOk() (*bool, bool)`

GetIsAutomaticReplyOk returns a tuple with the IsAutomaticReply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutomaticReply

`func (o *MicrosoftGraphMessageRulePredicates) SetIsAutomaticReply(v bool)`

SetIsAutomaticReply sets IsAutomaticReply field to given value.

### HasIsAutomaticReply

`func (o *MicrosoftGraphMessageRulePredicates) HasIsAutomaticReply() bool`

HasIsAutomaticReply returns a boolean if a field has been set.

### SetIsAutomaticReplyNil

`func (o *MicrosoftGraphMessageRulePredicates) SetIsAutomaticReplyNil(b bool)`

 SetIsAutomaticReplyNil sets the value for IsAutomaticReply to be an explicit nil

### UnsetIsAutomaticReply
`func (o *MicrosoftGraphMessageRulePredicates) UnsetIsAutomaticReply()`

UnsetIsAutomaticReply ensures that no value is present for IsAutomaticReply, not even an explicit nil
### GetIsEncrypted

`func (o *MicrosoftGraphMessageRulePredicates) GetIsEncrypted() bool`

GetIsEncrypted returns the IsEncrypted field if non-nil, zero value otherwise.

### GetIsEncryptedOk

`func (o *MicrosoftGraphMessageRulePredicates) GetIsEncryptedOk() (*bool, bool)`

GetIsEncryptedOk returns a tuple with the IsEncrypted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEncrypted

`func (o *MicrosoftGraphMessageRulePredicates) SetIsEncrypted(v bool)`

SetIsEncrypted sets IsEncrypted field to given value.

### HasIsEncrypted

`func (o *MicrosoftGraphMessageRulePredicates) HasIsEncrypted() bool`

HasIsEncrypted returns a boolean if a field has been set.

### SetIsEncryptedNil

`func (o *MicrosoftGraphMessageRulePredicates) SetIsEncryptedNil(b bool)`

 SetIsEncryptedNil sets the value for IsEncrypted to be an explicit nil

### UnsetIsEncrypted
`func (o *MicrosoftGraphMessageRulePredicates) UnsetIsEncrypted()`

UnsetIsEncrypted ensures that no value is present for IsEncrypted, not even an explicit nil
### GetIsMeetingRequest

`func (o *MicrosoftGraphMessageRulePredicates) GetIsMeetingRequest() bool`

GetIsMeetingRequest returns the IsMeetingRequest field if non-nil, zero value otherwise.

### GetIsMeetingRequestOk

`func (o *MicrosoftGraphMessageRulePredicates) GetIsMeetingRequestOk() (*bool, bool)`

GetIsMeetingRequestOk returns a tuple with the IsMeetingRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMeetingRequest

`func (o *MicrosoftGraphMessageRulePredicates) SetIsMeetingRequest(v bool)`

SetIsMeetingRequest sets IsMeetingRequest field to given value.

### HasIsMeetingRequest

`func (o *MicrosoftGraphMessageRulePredicates) HasIsMeetingRequest() bool`

HasIsMeetingRequest returns a boolean if a field has been set.

### SetIsMeetingRequestNil

`func (o *MicrosoftGraphMessageRulePredicates) SetIsMeetingRequestNil(b bool)`

 SetIsMeetingRequestNil sets the value for IsMeetingRequest to be an explicit nil

### UnsetIsMeetingRequest
`func (o *MicrosoftGraphMessageRulePredicates) UnsetIsMeetingRequest()`

UnsetIsMeetingRequest ensures that no value is present for IsMeetingRequest, not even an explicit nil
### GetIsMeetingResponse

`func (o *MicrosoftGraphMessageRulePredicates) GetIsMeetingResponse() bool`

GetIsMeetingResponse returns the IsMeetingResponse field if non-nil, zero value otherwise.

### GetIsMeetingResponseOk

`func (o *MicrosoftGraphMessageRulePredicates) GetIsMeetingResponseOk() (*bool, bool)`

GetIsMeetingResponseOk returns a tuple with the IsMeetingResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMeetingResponse

`func (o *MicrosoftGraphMessageRulePredicates) SetIsMeetingResponse(v bool)`

SetIsMeetingResponse sets IsMeetingResponse field to given value.

### HasIsMeetingResponse

`func (o *MicrosoftGraphMessageRulePredicates) HasIsMeetingResponse() bool`

HasIsMeetingResponse returns a boolean if a field has been set.

### SetIsMeetingResponseNil

`func (o *MicrosoftGraphMessageRulePredicates) SetIsMeetingResponseNil(b bool)`

 SetIsMeetingResponseNil sets the value for IsMeetingResponse to be an explicit nil

### UnsetIsMeetingResponse
`func (o *MicrosoftGraphMessageRulePredicates) UnsetIsMeetingResponse()`

UnsetIsMeetingResponse ensures that no value is present for IsMeetingResponse, not even an explicit nil
### GetIsNonDeliveryReport

`func (o *MicrosoftGraphMessageRulePredicates) GetIsNonDeliveryReport() bool`

GetIsNonDeliveryReport returns the IsNonDeliveryReport field if non-nil, zero value otherwise.

### GetIsNonDeliveryReportOk

`func (o *MicrosoftGraphMessageRulePredicates) GetIsNonDeliveryReportOk() (*bool, bool)`

GetIsNonDeliveryReportOk returns a tuple with the IsNonDeliveryReport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNonDeliveryReport

`func (o *MicrosoftGraphMessageRulePredicates) SetIsNonDeliveryReport(v bool)`

SetIsNonDeliveryReport sets IsNonDeliveryReport field to given value.

### HasIsNonDeliveryReport

`func (o *MicrosoftGraphMessageRulePredicates) HasIsNonDeliveryReport() bool`

HasIsNonDeliveryReport returns a boolean if a field has been set.

### SetIsNonDeliveryReportNil

`func (o *MicrosoftGraphMessageRulePredicates) SetIsNonDeliveryReportNil(b bool)`

 SetIsNonDeliveryReportNil sets the value for IsNonDeliveryReport to be an explicit nil

### UnsetIsNonDeliveryReport
`func (o *MicrosoftGraphMessageRulePredicates) UnsetIsNonDeliveryReport()`

UnsetIsNonDeliveryReport ensures that no value is present for IsNonDeliveryReport, not even an explicit nil
### GetIsPermissionControlled

`func (o *MicrosoftGraphMessageRulePredicates) GetIsPermissionControlled() bool`

GetIsPermissionControlled returns the IsPermissionControlled field if non-nil, zero value otherwise.

### GetIsPermissionControlledOk

`func (o *MicrosoftGraphMessageRulePredicates) GetIsPermissionControlledOk() (*bool, bool)`

GetIsPermissionControlledOk returns a tuple with the IsPermissionControlled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPermissionControlled

`func (o *MicrosoftGraphMessageRulePredicates) SetIsPermissionControlled(v bool)`

SetIsPermissionControlled sets IsPermissionControlled field to given value.

### HasIsPermissionControlled

`func (o *MicrosoftGraphMessageRulePredicates) HasIsPermissionControlled() bool`

HasIsPermissionControlled returns a boolean if a field has been set.

### SetIsPermissionControlledNil

`func (o *MicrosoftGraphMessageRulePredicates) SetIsPermissionControlledNil(b bool)`

 SetIsPermissionControlledNil sets the value for IsPermissionControlled to be an explicit nil

### UnsetIsPermissionControlled
`func (o *MicrosoftGraphMessageRulePredicates) UnsetIsPermissionControlled()`

UnsetIsPermissionControlled ensures that no value is present for IsPermissionControlled, not even an explicit nil
### GetIsReadReceipt

`func (o *MicrosoftGraphMessageRulePredicates) GetIsReadReceipt() bool`

GetIsReadReceipt returns the IsReadReceipt field if non-nil, zero value otherwise.

### GetIsReadReceiptOk

`func (o *MicrosoftGraphMessageRulePredicates) GetIsReadReceiptOk() (*bool, bool)`

GetIsReadReceiptOk returns a tuple with the IsReadReceipt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReadReceipt

`func (o *MicrosoftGraphMessageRulePredicates) SetIsReadReceipt(v bool)`

SetIsReadReceipt sets IsReadReceipt field to given value.

### HasIsReadReceipt

`func (o *MicrosoftGraphMessageRulePredicates) HasIsReadReceipt() bool`

HasIsReadReceipt returns a boolean if a field has been set.

### SetIsReadReceiptNil

`func (o *MicrosoftGraphMessageRulePredicates) SetIsReadReceiptNil(b bool)`

 SetIsReadReceiptNil sets the value for IsReadReceipt to be an explicit nil

### UnsetIsReadReceipt
`func (o *MicrosoftGraphMessageRulePredicates) UnsetIsReadReceipt()`

UnsetIsReadReceipt ensures that no value is present for IsReadReceipt, not even an explicit nil
### GetIsSigned

`func (o *MicrosoftGraphMessageRulePredicates) GetIsSigned() bool`

GetIsSigned returns the IsSigned field if non-nil, zero value otherwise.

### GetIsSignedOk

`func (o *MicrosoftGraphMessageRulePredicates) GetIsSignedOk() (*bool, bool)`

GetIsSignedOk returns a tuple with the IsSigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSigned

`func (o *MicrosoftGraphMessageRulePredicates) SetIsSigned(v bool)`

SetIsSigned sets IsSigned field to given value.

### HasIsSigned

`func (o *MicrosoftGraphMessageRulePredicates) HasIsSigned() bool`

HasIsSigned returns a boolean if a field has been set.

### SetIsSignedNil

`func (o *MicrosoftGraphMessageRulePredicates) SetIsSignedNil(b bool)`

 SetIsSignedNil sets the value for IsSigned to be an explicit nil

### UnsetIsSigned
`func (o *MicrosoftGraphMessageRulePredicates) UnsetIsSigned()`

UnsetIsSigned ensures that no value is present for IsSigned, not even an explicit nil
### GetIsVoicemail

`func (o *MicrosoftGraphMessageRulePredicates) GetIsVoicemail() bool`

GetIsVoicemail returns the IsVoicemail field if non-nil, zero value otherwise.

### GetIsVoicemailOk

`func (o *MicrosoftGraphMessageRulePredicates) GetIsVoicemailOk() (*bool, bool)`

GetIsVoicemailOk returns a tuple with the IsVoicemail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVoicemail

`func (o *MicrosoftGraphMessageRulePredicates) SetIsVoicemail(v bool)`

SetIsVoicemail sets IsVoicemail field to given value.

### HasIsVoicemail

`func (o *MicrosoftGraphMessageRulePredicates) HasIsVoicemail() bool`

HasIsVoicemail returns a boolean if a field has been set.

### SetIsVoicemailNil

`func (o *MicrosoftGraphMessageRulePredicates) SetIsVoicemailNil(b bool)`

 SetIsVoicemailNil sets the value for IsVoicemail to be an explicit nil

### UnsetIsVoicemail
`func (o *MicrosoftGraphMessageRulePredicates) UnsetIsVoicemail()`

UnsetIsVoicemail ensures that no value is present for IsVoicemail, not even an explicit nil
### GetMessageActionFlag

`func (o *MicrosoftGraphMessageRulePredicates) GetMessageActionFlag() AnyOfmicrosoftGraphMessageActionFlag`

GetMessageActionFlag returns the MessageActionFlag field if non-nil, zero value otherwise.

### GetMessageActionFlagOk

`func (o *MicrosoftGraphMessageRulePredicates) GetMessageActionFlagOk() (*AnyOfmicrosoftGraphMessageActionFlag, bool)`

GetMessageActionFlagOk returns a tuple with the MessageActionFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageActionFlag

`func (o *MicrosoftGraphMessageRulePredicates) SetMessageActionFlag(v AnyOfmicrosoftGraphMessageActionFlag)`

SetMessageActionFlag sets MessageActionFlag field to given value.

### HasMessageActionFlag

`func (o *MicrosoftGraphMessageRulePredicates) HasMessageActionFlag() bool`

HasMessageActionFlag returns a boolean if a field has been set.

### SetMessageActionFlagNil

`func (o *MicrosoftGraphMessageRulePredicates) SetMessageActionFlagNil(b bool)`

 SetMessageActionFlagNil sets the value for MessageActionFlag to be an explicit nil

### UnsetMessageActionFlag
`func (o *MicrosoftGraphMessageRulePredicates) UnsetMessageActionFlag()`

UnsetMessageActionFlag ensures that no value is present for MessageActionFlag, not even an explicit nil
### GetNotSentToMe

`func (o *MicrosoftGraphMessageRulePredicates) GetNotSentToMe() bool`

GetNotSentToMe returns the NotSentToMe field if non-nil, zero value otherwise.

### GetNotSentToMeOk

`func (o *MicrosoftGraphMessageRulePredicates) GetNotSentToMeOk() (*bool, bool)`

GetNotSentToMeOk returns a tuple with the NotSentToMe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotSentToMe

`func (o *MicrosoftGraphMessageRulePredicates) SetNotSentToMe(v bool)`

SetNotSentToMe sets NotSentToMe field to given value.

### HasNotSentToMe

`func (o *MicrosoftGraphMessageRulePredicates) HasNotSentToMe() bool`

HasNotSentToMe returns a boolean if a field has been set.

### SetNotSentToMeNil

`func (o *MicrosoftGraphMessageRulePredicates) SetNotSentToMeNil(b bool)`

 SetNotSentToMeNil sets the value for NotSentToMe to be an explicit nil

### UnsetNotSentToMe
`func (o *MicrosoftGraphMessageRulePredicates) UnsetNotSentToMe()`

UnsetNotSentToMe ensures that no value is present for NotSentToMe, not even an explicit nil
### GetRecipientContains

`func (o *MicrosoftGraphMessageRulePredicates) GetRecipientContains() []*string`

GetRecipientContains returns the RecipientContains field if non-nil, zero value otherwise.

### GetRecipientContainsOk

`func (o *MicrosoftGraphMessageRulePredicates) GetRecipientContainsOk() (*[]*string, bool)`

GetRecipientContainsOk returns a tuple with the RecipientContains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientContains

`func (o *MicrosoftGraphMessageRulePredicates) SetRecipientContains(v []*string)`

SetRecipientContains sets RecipientContains field to given value.

### HasRecipientContains

`func (o *MicrosoftGraphMessageRulePredicates) HasRecipientContains() bool`

HasRecipientContains returns a boolean if a field has been set.

### GetSenderContains

`func (o *MicrosoftGraphMessageRulePredicates) GetSenderContains() []*string`

GetSenderContains returns the SenderContains field if non-nil, zero value otherwise.

### GetSenderContainsOk

`func (o *MicrosoftGraphMessageRulePredicates) GetSenderContainsOk() (*[]*string, bool)`

GetSenderContainsOk returns a tuple with the SenderContains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderContains

`func (o *MicrosoftGraphMessageRulePredicates) SetSenderContains(v []*string)`

SetSenderContains sets SenderContains field to given value.

### HasSenderContains

`func (o *MicrosoftGraphMessageRulePredicates) HasSenderContains() bool`

HasSenderContains returns a boolean if a field has been set.

### GetSensitivity

`func (o *MicrosoftGraphMessageRulePredicates) GetSensitivity() AnyOfmicrosoftGraphSensitivity`

GetSensitivity returns the Sensitivity field if non-nil, zero value otherwise.

### GetSensitivityOk

`func (o *MicrosoftGraphMessageRulePredicates) GetSensitivityOk() (*AnyOfmicrosoftGraphSensitivity, bool)`

GetSensitivityOk returns a tuple with the Sensitivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSensitivity

`func (o *MicrosoftGraphMessageRulePredicates) SetSensitivity(v AnyOfmicrosoftGraphSensitivity)`

SetSensitivity sets Sensitivity field to given value.

### HasSensitivity

`func (o *MicrosoftGraphMessageRulePredicates) HasSensitivity() bool`

HasSensitivity returns a boolean if a field has been set.

### SetSensitivityNil

`func (o *MicrosoftGraphMessageRulePredicates) SetSensitivityNil(b bool)`

 SetSensitivityNil sets the value for Sensitivity to be an explicit nil

### UnsetSensitivity
`func (o *MicrosoftGraphMessageRulePredicates) UnsetSensitivity()`

UnsetSensitivity ensures that no value is present for Sensitivity, not even an explicit nil
### GetSentCcMe

`func (o *MicrosoftGraphMessageRulePredicates) GetSentCcMe() bool`

GetSentCcMe returns the SentCcMe field if non-nil, zero value otherwise.

### GetSentCcMeOk

`func (o *MicrosoftGraphMessageRulePredicates) GetSentCcMeOk() (*bool, bool)`

GetSentCcMeOk returns a tuple with the SentCcMe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSentCcMe

`func (o *MicrosoftGraphMessageRulePredicates) SetSentCcMe(v bool)`

SetSentCcMe sets SentCcMe field to given value.

### HasSentCcMe

`func (o *MicrosoftGraphMessageRulePredicates) HasSentCcMe() bool`

HasSentCcMe returns a boolean if a field has been set.

### SetSentCcMeNil

`func (o *MicrosoftGraphMessageRulePredicates) SetSentCcMeNil(b bool)`

 SetSentCcMeNil sets the value for SentCcMe to be an explicit nil

### UnsetSentCcMe
`func (o *MicrosoftGraphMessageRulePredicates) UnsetSentCcMe()`

UnsetSentCcMe ensures that no value is present for SentCcMe, not even an explicit nil
### GetSentOnlyToMe

`func (o *MicrosoftGraphMessageRulePredicates) GetSentOnlyToMe() bool`

GetSentOnlyToMe returns the SentOnlyToMe field if non-nil, zero value otherwise.

### GetSentOnlyToMeOk

`func (o *MicrosoftGraphMessageRulePredicates) GetSentOnlyToMeOk() (*bool, bool)`

GetSentOnlyToMeOk returns a tuple with the SentOnlyToMe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSentOnlyToMe

`func (o *MicrosoftGraphMessageRulePredicates) SetSentOnlyToMe(v bool)`

SetSentOnlyToMe sets SentOnlyToMe field to given value.

### HasSentOnlyToMe

`func (o *MicrosoftGraphMessageRulePredicates) HasSentOnlyToMe() bool`

HasSentOnlyToMe returns a boolean if a field has been set.

### SetSentOnlyToMeNil

`func (o *MicrosoftGraphMessageRulePredicates) SetSentOnlyToMeNil(b bool)`

 SetSentOnlyToMeNil sets the value for SentOnlyToMe to be an explicit nil

### UnsetSentOnlyToMe
`func (o *MicrosoftGraphMessageRulePredicates) UnsetSentOnlyToMe()`

UnsetSentOnlyToMe ensures that no value is present for SentOnlyToMe, not even an explicit nil
### GetSentToAddresses

`func (o *MicrosoftGraphMessageRulePredicates) GetSentToAddresses() []*AnyOfmicrosoftGraphRecipient`

GetSentToAddresses returns the SentToAddresses field if non-nil, zero value otherwise.

### GetSentToAddressesOk

`func (o *MicrosoftGraphMessageRulePredicates) GetSentToAddressesOk() (*[]*AnyOfmicrosoftGraphRecipient, bool)`

GetSentToAddressesOk returns a tuple with the SentToAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSentToAddresses

`func (o *MicrosoftGraphMessageRulePredicates) SetSentToAddresses(v []*AnyOfmicrosoftGraphRecipient)`

SetSentToAddresses sets SentToAddresses field to given value.

### HasSentToAddresses

`func (o *MicrosoftGraphMessageRulePredicates) HasSentToAddresses() bool`

HasSentToAddresses returns a boolean if a field has been set.

### GetSentToMe

`func (o *MicrosoftGraphMessageRulePredicates) GetSentToMe() bool`

GetSentToMe returns the SentToMe field if non-nil, zero value otherwise.

### GetSentToMeOk

`func (o *MicrosoftGraphMessageRulePredicates) GetSentToMeOk() (*bool, bool)`

GetSentToMeOk returns a tuple with the SentToMe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSentToMe

`func (o *MicrosoftGraphMessageRulePredicates) SetSentToMe(v bool)`

SetSentToMe sets SentToMe field to given value.

### HasSentToMe

`func (o *MicrosoftGraphMessageRulePredicates) HasSentToMe() bool`

HasSentToMe returns a boolean if a field has been set.

### SetSentToMeNil

`func (o *MicrosoftGraphMessageRulePredicates) SetSentToMeNil(b bool)`

 SetSentToMeNil sets the value for SentToMe to be an explicit nil

### UnsetSentToMe
`func (o *MicrosoftGraphMessageRulePredicates) UnsetSentToMe()`

UnsetSentToMe ensures that no value is present for SentToMe, not even an explicit nil
### GetSentToOrCcMe

`func (o *MicrosoftGraphMessageRulePredicates) GetSentToOrCcMe() bool`

GetSentToOrCcMe returns the SentToOrCcMe field if non-nil, zero value otherwise.

### GetSentToOrCcMeOk

`func (o *MicrosoftGraphMessageRulePredicates) GetSentToOrCcMeOk() (*bool, bool)`

GetSentToOrCcMeOk returns a tuple with the SentToOrCcMe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSentToOrCcMe

`func (o *MicrosoftGraphMessageRulePredicates) SetSentToOrCcMe(v bool)`

SetSentToOrCcMe sets SentToOrCcMe field to given value.

### HasSentToOrCcMe

`func (o *MicrosoftGraphMessageRulePredicates) HasSentToOrCcMe() bool`

HasSentToOrCcMe returns a boolean if a field has been set.

### SetSentToOrCcMeNil

`func (o *MicrosoftGraphMessageRulePredicates) SetSentToOrCcMeNil(b bool)`

 SetSentToOrCcMeNil sets the value for SentToOrCcMe to be an explicit nil

### UnsetSentToOrCcMe
`func (o *MicrosoftGraphMessageRulePredicates) UnsetSentToOrCcMe()`

UnsetSentToOrCcMe ensures that no value is present for SentToOrCcMe, not even an explicit nil
### GetSubjectContains

`func (o *MicrosoftGraphMessageRulePredicates) GetSubjectContains() []*string`

GetSubjectContains returns the SubjectContains field if non-nil, zero value otherwise.

### GetSubjectContainsOk

`func (o *MicrosoftGraphMessageRulePredicates) GetSubjectContainsOk() (*[]*string, bool)`

GetSubjectContainsOk returns a tuple with the SubjectContains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectContains

`func (o *MicrosoftGraphMessageRulePredicates) SetSubjectContains(v []*string)`

SetSubjectContains sets SubjectContains field to given value.

### HasSubjectContains

`func (o *MicrosoftGraphMessageRulePredicates) HasSubjectContains() bool`

HasSubjectContains returns a boolean if a field has been set.

### GetWithinSizeRange

`func (o *MicrosoftGraphMessageRulePredicates) GetWithinSizeRange() AnyOfmicrosoftGraphSizeRange`

GetWithinSizeRange returns the WithinSizeRange field if non-nil, zero value otherwise.

### GetWithinSizeRangeOk

`func (o *MicrosoftGraphMessageRulePredicates) GetWithinSizeRangeOk() (*AnyOfmicrosoftGraphSizeRange, bool)`

GetWithinSizeRangeOk returns a tuple with the WithinSizeRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithinSizeRange

`func (o *MicrosoftGraphMessageRulePredicates) SetWithinSizeRange(v AnyOfmicrosoftGraphSizeRange)`

SetWithinSizeRange sets WithinSizeRange field to given value.

### HasWithinSizeRange

`func (o *MicrosoftGraphMessageRulePredicates) HasWithinSizeRange() bool`

HasWithinSizeRange returns a boolean if a field has been set.

### SetWithinSizeRangeNil

`func (o *MicrosoftGraphMessageRulePredicates) SetWithinSizeRangeNil(b bool)`

 SetWithinSizeRangeNil sets the value for WithinSizeRange to be an explicit nil

### UnsetWithinSizeRange
`func (o *MicrosoftGraphMessageRulePredicates) UnsetWithinSizeRange()`

UnsetWithinSizeRange ensures that no value is present for WithinSizeRange, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


