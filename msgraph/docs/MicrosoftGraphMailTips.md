# MicrosoftGraphMailTips

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutomaticReplies** | Pointer to [**AnyOfmicrosoftGraphAutomaticRepliesMailTips**](anyOf&lt;microsoft.graph.automaticRepliesMailTips&gt;.md) | Mail tips for automatic reply if it has been set up by the recipient. | [optional] 
**CustomMailTip** | Pointer to **NullableString** | A custom mail tip that can be set on the recipient&#39;s mailbox. | [optional] 
**DeliveryRestricted** | Pointer to **NullableBool** | Whether the recipient&#39;s mailbox is restricted, for example, accepting messages from only a predefined list of senders, rejecting messages from a predefined list of senders, or accepting messages from only authenticated senders. | [optional] 
**EmailAddress** | Pointer to [**AnyOfmicrosoftGraphEmailAddress**](anyOf&lt;microsoft.graph.emailAddress&gt;.md) | The email address of the recipient to get mailtips for. | [optional] 
**Error** | Pointer to [**AnyOfmicrosoftGraphMailTipsError**](anyOf&lt;microsoft.graph.mailTipsError&gt;.md) | Errors that occur during the getMailTips action. | [optional] 
**ExternalMemberCount** | Pointer to **NullableInt32** | The number of external members if the recipient is a distribution list. | [optional] 
**IsModerated** | Pointer to **NullableBool** | Whether sending messages to the recipient requires approval. For example, if the recipient is a large distribution list and a moderator has been set up to approve messages sent to that distribution list, or if sending messages to a recipient requires approval of the recipient&#39;s manager. | [optional] 
**MailboxFull** | Pointer to **NullableBool** | The mailbox full status of the recipient. | [optional] 
**MaxMessageSize** | Pointer to **NullableInt32** | The maximum message size that has been configured for the recipient&#39;s organization or mailbox. | [optional] 
**RecipientScope** | Pointer to [**AnyOfmicrosoftGraphRecipientScopeType**](anyOf&lt;microsoft.graph.recipientScopeType&gt;.md) | The scope of the recipient. Possible values are: none, internal, external, externalPartner, externalNonParther. For example, an administrator can set another organization to be its &#39;partner&#39;. The scope is useful if an administrator wants certain mailtips to be accessible to certain scopes. It&#39;s also useful to senders to inform them that their message may leave the organization, helping them make the correct decisions about wording, tone and content. | [optional] 
**RecipientSuggestions** | Pointer to [**[]AnyOfmicrosoftGraphRecipient**](AnyOfmicrosoftGraphRecipient.md) | Recipients suggested based on previous contexts where they appear in the same message. | [optional] 
**TotalMemberCount** | Pointer to **NullableInt32** | The number of members if the recipient is a distribution list. | [optional] 

## Methods

### NewMicrosoftGraphMailTips

`func NewMicrosoftGraphMailTips() *MicrosoftGraphMailTips`

NewMicrosoftGraphMailTips instantiates a new MicrosoftGraphMailTips object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphMailTipsWithDefaults

`func NewMicrosoftGraphMailTipsWithDefaults() *MicrosoftGraphMailTips`

NewMicrosoftGraphMailTipsWithDefaults instantiates a new MicrosoftGraphMailTips object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutomaticReplies

`func (o *MicrosoftGraphMailTips) GetAutomaticReplies() AnyOfmicrosoftGraphAutomaticRepliesMailTips`

GetAutomaticReplies returns the AutomaticReplies field if non-nil, zero value otherwise.

### GetAutomaticRepliesOk

`func (o *MicrosoftGraphMailTips) GetAutomaticRepliesOk() (*AnyOfmicrosoftGraphAutomaticRepliesMailTips, bool)`

GetAutomaticRepliesOk returns a tuple with the AutomaticReplies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomaticReplies

`func (o *MicrosoftGraphMailTips) SetAutomaticReplies(v AnyOfmicrosoftGraphAutomaticRepliesMailTips)`

SetAutomaticReplies sets AutomaticReplies field to given value.

### HasAutomaticReplies

`func (o *MicrosoftGraphMailTips) HasAutomaticReplies() bool`

HasAutomaticReplies returns a boolean if a field has been set.

### SetAutomaticRepliesNil

`func (o *MicrosoftGraphMailTips) SetAutomaticRepliesNil(b bool)`

 SetAutomaticRepliesNil sets the value for AutomaticReplies to be an explicit nil

### UnsetAutomaticReplies
`func (o *MicrosoftGraphMailTips) UnsetAutomaticReplies()`

UnsetAutomaticReplies ensures that no value is present for AutomaticReplies, not even an explicit nil
### GetCustomMailTip

`func (o *MicrosoftGraphMailTips) GetCustomMailTip() string`

GetCustomMailTip returns the CustomMailTip field if non-nil, zero value otherwise.

### GetCustomMailTipOk

`func (o *MicrosoftGraphMailTips) GetCustomMailTipOk() (*string, bool)`

GetCustomMailTipOk returns a tuple with the CustomMailTip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomMailTip

`func (o *MicrosoftGraphMailTips) SetCustomMailTip(v string)`

SetCustomMailTip sets CustomMailTip field to given value.

### HasCustomMailTip

`func (o *MicrosoftGraphMailTips) HasCustomMailTip() bool`

HasCustomMailTip returns a boolean if a field has been set.

### SetCustomMailTipNil

`func (o *MicrosoftGraphMailTips) SetCustomMailTipNil(b bool)`

 SetCustomMailTipNil sets the value for CustomMailTip to be an explicit nil

### UnsetCustomMailTip
`func (o *MicrosoftGraphMailTips) UnsetCustomMailTip()`

UnsetCustomMailTip ensures that no value is present for CustomMailTip, not even an explicit nil
### GetDeliveryRestricted

`func (o *MicrosoftGraphMailTips) GetDeliveryRestricted() bool`

GetDeliveryRestricted returns the DeliveryRestricted field if non-nil, zero value otherwise.

### GetDeliveryRestrictedOk

`func (o *MicrosoftGraphMailTips) GetDeliveryRestrictedOk() (*bool, bool)`

GetDeliveryRestrictedOk returns a tuple with the DeliveryRestricted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryRestricted

`func (o *MicrosoftGraphMailTips) SetDeliveryRestricted(v bool)`

SetDeliveryRestricted sets DeliveryRestricted field to given value.

### HasDeliveryRestricted

`func (o *MicrosoftGraphMailTips) HasDeliveryRestricted() bool`

HasDeliveryRestricted returns a boolean if a field has been set.

### SetDeliveryRestrictedNil

`func (o *MicrosoftGraphMailTips) SetDeliveryRestrictedNil(b bool)`

 SetDeliveryRestrictedNil sets the value for DeliveryRestricted to be an explicit nil

### UnsetDeliveryRestricted
`func (o *MicrosoftGraphMailTips) UnsetDeliveryRestricted()`

UnsetDeliveryRestricted ensures that no value is present for DeliveryRestricted, not even an explicit nil
### GetEmailAddress

`func (o *MicrosoftGraphMailTips) GetEmailAddress() AnyOfmicrosoftGraphEmailAddress`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *MicrosoftGraphMailTips) GetEmailAddressOk() (*AnyOfmicrosoftGraphEmailAddress, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *MicrosoftGraphMailTips) SetEmailAddress(v AnyOfmicrosoftGraphEmailAddress)`

SetEmailAddress sets EmailAddress field to given value.

### HasEmailAddress

`func (o *MicrosoftGraphMailTips) HasEmailAddress() bool`

HasEmailAddress returns a boolean if a field has been set.

### SetEmailAddressNil

`func (o *MicrosoftGraphMailTips) SetEmailAddressNil(b bool)`

 SetEmailAddressNil sets the value for EmailAddress to be an explicit nil

### UnsetEmailAddress
`func (o *MicrosoftGraphMailTips) UnsetEmailAddress()`

UnsetEmailAddress ensures that no value is present for EmailAddress, not even an explicit nil
### GetError

`func (o *MicrosoftGraphMailTips) GetError() AnyOfmicrosoftGraphMailTipsError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *MicrosoftGraphMailTips) GetErrorOk() (*AnyOfmicrosoftGraphMailTipsError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *MicrosoftGraphMailTips) SetError(v AnyOfmicrosoftGraphMailTipsError)`

SetError sets Error field to given value.

### HasError

`func (o *MicrosoftGraphMailTips) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *MicrosoftGraphMailTips) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *MicrosoftGraphMailTips) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetExternalMemberCount

`func (o *MicrosoftGraphMailTips) GetExternalMemberCount() int32`

GetExternalMemberCount returns the ExternalMemberCount field if non-nil, zero value otherwise.

### GetExternalMemberCountOk

`func (o *MicrosoftGraphMailTips) GetExternalMemberCountOk() (*int32, bool)`

GetExternalMemberCountOk returns a tuple with the ExternalMemberCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalMemberCount

`func (o *MicrosoftGraphMailTips) SetExternalMemberCount(v int32)`

SetExternalMemberCount sets ExternalMemberCount field to given value.

### HasExternalMemberCount

`func (o *MicrosoftGraphMailTips) HasExternalMemberCount() bool`

HasExternalMemberCount returns a boolean if a field has been set.

### SetExternalMemberCountNil

`func (o *MicrosoftGraphMailTips) SetExternalMemberCountNil(b bool)`

 SetExternalMemberCountNil sets the value for ExternalMemberCount to be an explicit nil

### UnsetExternalMemberCount
`func (o *MicrosoftGraphMailTips) UnsetExternalMemberCount()`

UnsetExternalMemberCount ensures that no value is present for ExternalMemberCount, not even an explicit nil
### GetIsModerated

`func (o *MicrosoftGraphMailTips) GetIsModerated() bool`

GetIsModerated returns the IsModerated field if non-nil, zero value otherwise.

### GetIsModeratedOk

`func (o *MicrosoftGraphMailTips) GetIsModeratedOk() (*bool, bool)`

GetIsModeratedOk returns a tuple with the IsModerated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsModerated

`func (o *MicrosoftGraphMailTips) SetIsModerated(v bool)`

SetIsModerated sets IsModerated field to given value.

### HasIsModerated

`func (o *MicrosoftGraphMailTips) HasIsModerated() bool`

HasIsModerated returns a boolean if a field has been set.

### SetIsModeratedNil

`func (o *MicrosoftGraphMailTips) SetIsModeratedNil(b bool)`

 SetIsModeratedNil sets the value for IsModerated to be an explicit nil

### UnsetIsModerated
`func (o *MicrosoftGraphMailTips) UnsetIsModerated()`

UnsetIsModerated ensures that no value is present for IsModerated, not even an explicit nil
### GetMailboxFull

`func (o *MicrosoftGraphMailTips) GetMailboxFull() bool`

GetMailboxFull returns the MailboxFull field if non-nil, zero value otherwise.

### GetMailboxFullOk

`func (o *MicrosoftGraphMailTips) GetMailboxFullOk() (*bool, bool)`

GetMailboxFullOk returns a tuple with the MailboxFull field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailboxFull

`func (o *MicrosoftGraphMailTips) SetMailboxFull(v bool)`

SetMailboxFull sets MailboxFull field to given value.

### HasMailboxFull

`func (o *MicrosoftGraphMailTips) HasMailboxFull() bool`

HasMailboxFull returns a boolean if a field has been set.

### SetMailboxFullNil

`func (o *MicrosoftGraphMailTips) SetMailboxFullNil(b bool)`

 SetMailboxFullNil sets the value for MailboxFull to be an explicit nil

### UnsetMailboxFull
`func (o *MicrosoftGraphMailTips) UnsetMailboxFull()`

UnsetMailboxFull ensures that no value is present for MailboxFull, not even an explicit nil
### GetMaxMessageSize

`func (o *MicrosoftGraphMailTips) GetMaxMessageSize() int32`

GetMaxMessageSize returns the MaxMessageSize field if non-nil, zero value otherwise.

### GetMaxMessageSizeOk

`func (o *MicrosoftGraphMailTips) GetMaxMessageSizeOk() (*int32, bool)`

GetMaxMessageSizeOk returns a tuple with the MaxMessageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMessageSize

`func (o *MicrosoftGraphMailTips) SetMaxMessageSize(v int32)`

SetMaxMessageSize sets MaxMessageSize field to given value.

### HasMaxMessageSize

`func (o *MicrosoftGraphMailTips) HasMaxMessageSize() bool`

HasMaxMessageSize returns a boolean if a field has been set.

### SetMaxMessageSizeNil

`func (o *MicrosoftGraphMailTips) SetMaxMessageSizeNil(b bool)`

 SetMaxMessageSizeNil sets the value for MaxMessageSize to be an explicit nil

### UnsetMaxMessageSize
`func (o *MicrosoftGraphMailTips) UnsetMaxMessageSize()`

UnsetMaxMessageSize ensures that no value is present for MaxMessageSize, not even an explicit nil
### GetRecipientScope

`func (o *MicrosoftGraphMailTips) GetRecipientScope() AnyOfmicrosoftGraphRecipientScopeType`

GetRecipientScope returns the RecipientScope field if non-nil, zero value otherwise.

### GetRecipientScopeOk

`func (o *MicrosoftGraphMailTips) GetRecipientScopeOk() (*AnyOfmicrosoftGraphRecipientScopeType, bool)`

GetRecipientScopeOk returns a tuple with the RecipientScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientScope

`func (o *MicrosoftGraphMailTips) SetRecipientScope(v AnyOfmicrosoftGraphRecipientScopeType)`

SetRecipientScope sets RecipientScope field to given value.

### HasRecipientScope

`func (o *MicrosoftGraphMailTips) HasRecipientScope() bool`

HasRecipientScope returns a boolean if a field has been set.

### SetRecipientScopeNil

`func (o *MicrosoftGraphMailTips) SetRecipientScopeNil(b bool)`

 SetRecipientScopeNil sets the value for RecipientScope to be an explicit nil

### UnsetRecipientScope
`func (o *MicrosoftGraphMailTips) UnsetRecipientScope()`

UnsetRecipientScope ensures that no value is present for RecipientScope, not even an explicit nil
### GetRecipientSuggestions

`func (o *MicrosoftGraphMailTips) GetRecipientSuggestions() []*AnyOfmicrosoftGraphRecipient`

GetRecipientSuggestions returns the RecipientSuggestions field if non-nil, zero value otherwise.

### GetRecipientSuggestionsOk

`func (o *MicrosoftGraphMailTips) GetRecipientSuggestionsOk() (*[]*AnyOfmicrosoftGraphRecipient, bool)`

GetRecipientSuggestionsOk returns a tuple with the RecipientSuggestions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientSuggestions

`func (o *MicrosoftGraphMailTips) SetRecipientSuggestions(v []*AnyOfmicrosoftGraphRecipient)`

SetRecipientSuggestions sets RecipientSuggestions field to given value.

### HasRecipientSuggestions

`func (o *MicrosoftGraphMailTips) HasRecipientSuggestions() bool`

HasRecipientSuggestions returns a boolean if a field has been set.

### GetTotalMemberCount

`func (o *MicrosoftGraphMailTips) GetTotalMemberCount() int32`

GetTotalMemberCount returns the TotalMemberCount field if non-nil, zero value otherwise.

### GetTotalMemberCountOk

`func (o *MicrosoftGraphMailTips) GetTotalMemberCountOk() (*int32, bool)`

GetTotalMemberCountOk returns a tuple with the TotalMemberCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalMemberCount

`func (o *MicrosoftGraphMailTips) SetTotalMemberCount(v int32)`

SetTotalMemberCount sets TotalMemberCount field to given value.

### HasTotalMemberCount

`func (o *MicrosoftGraphMailTips) HasTotalMemberCount() bool`

HasTotalMemberCount returns a boolean if a field has been set.

### SetTotalMemberCountNil

`func (o *MicrosoftGraphMailTips) SetTotalMemberCountNil(b bool)`

 SetTotalMemberCountNil sets the value for TotalMemberCount to be an explicit nil

### UnsetTotalMemberCount
`func (o *MicrosoftGraphMailTips) UnsetTotalMemberCount()`

UnsetTotalMemberCount ensures that no value is present for TotalMemberCount, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


