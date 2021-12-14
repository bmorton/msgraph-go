# MicrosoftGraphInvitedUserMessageInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CcRecipients** | Pointer to [**[]AnyOfmicrosoftGraphRecipient**](AnyOfmicrosoftGraphRecipient.md) | Additional recipients the invitation message should be sent to. Currently only 1 additional recipient is supported. | [optional] 
**CustomizedMessageBody** | Pointer to **NullableString** | Customized message body you want to send if you don&#39;t want the default message. | [optional] 
**MessageLanguage** | Pointer to **NullableString** | The language you want to send the default message in. If the customizedMessageBody is specified, this property is ignored, and the message is sent using the customizedMessageBody. The language format should be in ISO 639. The default is en-US. | [optional] 

## Methods

### NewMicrosoftGraphInvitedUserMessageInfo

`func NewMicrosoftGraphInvitedUserMessageInfo() *MicrosoftGraphInvitedUserMessageInfo`

NewMicrosoftGraphInvitedUserMessageInfo instantiates a new MicrosoftGraphInvitedUserMessageInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphInvitedUserMessageInfoWithDefaults

`func NewMicrosoftGraphInvitedUserMessageInfoWithDefaults() *MicrosoftGraphInvitedUserMessageInfo`

NewMicrosoftGraphInvitedUserMessageInfoWithDefaults instantiates a new MicrosoftGraphInvitedUserMessageInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCcRecipients

`func (o *MicrosoftGraphInvitedUserMessageInfo) GetCcRecipients() []*AnyOfmicrosoftGraphRecipient`

GetCcRecipients returns the CcRecipients field if non-nil, zero value otherwise.

### GetCcRecipientsOk

`func (o *MicrosoftGraphInvitedUserMessageInfo) GetCcRecipientsOk() (*[]*AnyOfmicrosoftGraphRecipient, bool)`

GetCcRecipientsOk returns a tuple with the CcRecipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCcRecipients

`func (o *MicrosoftGraphInvitedUserMessageInfo) SetCcRecipients(v []*AnyOfmicrosoftGraphRecipient)`

SetCcRecipients sets CcRecipients field to given value.

### HasCcRecipients

`func (o *MicrosoftGraphInvitedUserMessageInfo) HasCcRecipients() bool`

HasCcRecipients returns a boolean if a field has been set.

### GetCustomizedMessageBody

`func (o *MicrosoftGraphInvitedUserMessageInfo) GetCustomizedMessageBody() string`

GetCustomizedMessageBody returns the CustomizedMessageBody field if non-nil, zero value otherwise.

### GetCustomizedMessageBodyOk

`func (o *MicrosoftGraphInvitedUserMessageInfo) GetCustomizedMessageBodyOk() (*string, bool)`

GetCustomizedMessageBodyOk returns a tuple with the CustomizedMessageBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomizedMessageBody

`func (o *MicrosoftGraphInvitedUserMessageInfo) SetCustomizedMessageBody(v string)`

SetCustomizedMessageBody sets CustomizedMessageBody field to given value.

### HasCustomizedMessageBody

`func (o *MicrosoftGraphInvitedUserMessageInfo) HasCustomizedMessageBody() bool`

HasCustomizedMessageBody returns a boolean if a field has been set.

### SetCustomizedMessageBodyNil

`func (o *MicrosoftGraphInvitedUserMessageInfo) SetCustomizedMessageBodyNil(b bool)`

 SetCustomizedMessageBodyNil sets the value for CustomizedMessageBody to be an explicit nil

### UnsetCustomizedMessageBody
`func (o *MicrosoftGraphInvitedUserMessageInfo) UnsetCustomizedMessageBody()`

UnsetCustomizedMessageBody ensures that no value is present for CustomizedMessageBody, not even an explicit nil
### GetMessageLanguage

`func (o *MicrosoftGraphInvitedUserMessageInfo) GetMessageLanguage() string`

GetMessageLanguage returns the MessageLanguage field if non-nil, zero value otherwise.

### GetMessageLanguageOk

`func (o *MicrosoftGraphInvitedUserMessageInfo) GetMessageLanguageOk() (*string, bool)`

GetMessageLanguageOk returns a tuple with the MessageLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageLanguage

`func (o *MicrosoftGraphInvitedUserMessageInfo) SetMessageLanguage(v string)`

SetMessageLanguage sets MessageLanguage field to given value.

### HasMessageLanguage

`func (o *MicrosoftGraphInvitedUserMessageInfo) HasMessageLanguage() bool`

HasMessageLanguage returns a boolean if a field has been set.

### SetMessageLanguageNil

`func (o *MicrosoftGraphInvitedUserMessageInfo) SetMessageLanguageNil(b bool)`

 SetMessageLanguageNil sets the value for MessageLanguage to be an explicit nil

### UnsetMessageLanguage
`func (o *MicrosoftGraphInvitedUserMessageInfo) UnsetMessageLanguage()`

UnsetMessageLanguage ensures that no value is present for MessageLanguage, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


