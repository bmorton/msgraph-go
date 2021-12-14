# MicrosoftGraphChatMessagePolicyViolation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DlpAction** | Pointer to [**AnyOfmicrosoftGraphChatMessagePolicyViolationDlpActionTypes**](anyOf&lt;microsoft.graph.chatMessagePolicyViolationDlpActionTypes&gt;.md) | The action taken by the DLP provider on the message with sensitive content. Supported values are: NoneNotifySender -- Inform the sender of the violation but allow readers to read the message.BlockAccess -- Block readers from reading the message.BlockAccessExternal -- Block users outside the organization from reading the message, while allowing users within the organization to read the message. | [optional] 
**JustificationText** | Pointer to **NullableString** | Justification text provided by the sender of the message when overriding a policy violation. | [optional] 
**PolicyTip** | Pointer to [**AnyOfmicrosoftGraphChatMessagePolicyViolationPolicyTip**](anyOf&lt;microsoft.graph.chatMessagePolicyViolationPolicyTip&gt;.md) | Information to display to the message sender about why the message was flagged as a violation. | [optional] 
**UserAction** | Pointer to [**AnyOfmicrosoftGraphChatMessagePolicyViolationUserActionTypes**](anyOf&lt;microsoft.graph.chatMessagePolicyViolationUserActionTypes&gt;.md) | Indicates the action taken by the user on a message blocked by the DLP provider. Supported values are: NoneOverrideReportFalsePositiveWhen the DLP provider is updating the message for blocking sensitive content, userAction is not required. | [optional] 
**VerdictDetails** | Pointer to [**AnyOfmicrosoftGraphChatMessagePolicyViolationVerdictDetailsTypes**](anyOf&lt;microsoft.graph.chatMessagePolicyViolationVerdictDetailsTypes&gt;.md) | Indicates what actions the sender may take in response to the policy violation. Supported values are: NoneAllowFalsePositiveOverride -- Allows the sender to declare the policyViolation to be an error in the DLP app and its rules, and allow readers to see the message again if the dlpAction had hidden it.AllowOverrideWithoutJustification -- Allows the sender to overriide the DLP violation and allow readers to see the message again if the dlpAction had hidden it, without needing to provide an explanation for doing so. AllowOverrideWithJustification -- Allows the sender to overriide the DLP violation and allow readers to see the message again if the dlpAction had hidden it, after providing an explanation for doing so.AllowOverrideWithoutJustification and AllowOverrideWithJustification are mutually exclusive. | [optional] 

## Methods

### NewMicrosoftGraphChatMessagePolicyViolation

`func NewMicrosoftGraphChatMessagePolicyViolation() *MicrosoftGraphChatMessagePolicyViolation`

NewMicrosoftGraphChatMessagePolicyViolation instantiates a new MicrosoftGraphChatMessagePolicyViolation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphChatMessagePolicyViolationWithDefaults

`func NewMicrosoftGraphChatMessagePolicyViolationWithDefaults() *MicrosoftGraphChatMessagePolicyViolation`

NewMicrosoftGraphChatMessagePolicyViolationWithDefaults instantiates a new MicrosoftGraphChatMessagePolicyViolation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDlpAction

`func (o *MicrosoftGraphChatMessagePolicyViolation) GetDlpAction() AnyOfmicrosoftGraphChatMessagePolicyViolationDlpActionTypes`

GetDlpAction returns the DlpAction field if non-nil, zero value otherwise.

### GetDlpActionOk

`func (o *MicrosoftGraphChatMessagePolicyViolation) GetDlpActionOk() (*AnyOfmicrosoftGraphChatMessagePolicyViolationDlpActionTypes, bool)`

GetDlpActionOk returns a tuple with the DlpAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDlpAction

`func (o *MicrosoftGraphChatMessagePolicyViolation) SetDlpAction(v AnyOfmicrosoftGraphChatMessagePolicyViolationDlpActionTypes)`

SetDlpAction sets DlpAction field to given value.

### HasDlpAction

`func (o *MicrosoftGraphChatMessagePolicyViolation) HasDlpAction() bool`

HasDlpAction returns a boolean if a field has been set.

### SetDlpActionNil

`func (o *MicrosoftGraphChatMessagePolicyViolation) SetDlpActionNil(b bool)`

 SetDlpActionNil sets the value for DlpAction to be an explicit nil

### UnsetDlpAction
`func (o *MicrosoftGraphChatMessagePolicyViolation) UnsetDlpAction()`

UnsetDlpAction ensures that no value is present for DlpAction, not even an explicit nil
### GetJustificationText

`func (o *MicrosoftGraphChatMessagePolicyViolation) GetJustificationText() string`

GetJustificationText returns the JustificationText field if non-nil, zero value otherwise.

### GetJustificationTextOk

`func (o *MicrosoftGraphChatMessagePolicyViolation) GetJustificationTextOk() (*string, bool)`

GetJustificationTextOk returns a tuple with the JustificationText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJustificationText

`func (o *MicrosoftGraphChatMessagePolicyViolation) SetJustificationText(v string)`

SetJustificationText sets JustificationText field to given value.

### HasJustificationText

`func (o *MicrosoftGraphChatMessagePolicyViolation) HasJustificationText() bool`

HasJustificationText returns a boolean if a field has been set.

### SetJustificationTextNil

`func (o *MicrosoftGraphChatMessagePolicyViolation) SetJustificationTextNil(b bool)`

 SetJustificationTextNil sets the value for JustificationText to be an explicit nil

### UnsetJustificationText
`func (o *MicrosoftGraphChatMessagePolicyViolation) UnsetJustificationText()`

UnsetJustificationText ensures that no value is present for JustificationText, not even an explicit nil
### GetPolicyTip

`func (o *MicrosoftGraphChatMessagePolicyViolation) GetPolicyTip() AnyOfmicrosoftGraphChatMessagePolicyViolationPolicyTip`

GetPolicyTip returns the PolicyTip field if non-nil, zero value otherwise.

### GetPolicyTipOk

`func (o *MicrosoftGraphChatMessagePolicyViolation) GetPolicyTipOk() (*AnyOfmicrosoftGraphChatMessagePolicyViolationPolicyTip, bool)`

GetPolicyTipOk returns a tuple with the PolicyTip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyTip

`func (o *MicrosoftGraphChatMessagePolicyViolation) SetPolicyTip(v AnyOfmicrosoftGraphChatMessagePolicyViolationPolicyTip)`

SetPolicyTip sets PolicyTip field to given value.

### HasPolicyTip

`func (o *MicrosoftGraphChatMessagePolicyViolation) HasPolicyTip() bool`

HasPolicyTip returns a boolean if a field has been set.

### SetPolicyTipNil

`func (o *MicrosoftGraphChatMessagePolicyViolation) SetPolicyTipNil(b bool)`

 SetPolicyTipNil sets the value for PolicyTip to be an explicit nil

### UnsetPolicyTip
`func (o *MicrosoftGraphChatMessagePolicyViolation) UnsetPolicyTip()`

UnsetPolicyTip ensures that no value is present for PolicyTip, not even an explicit nil
### GetUserAction

`func (o *MicrosoftGraphChatMessagePolicyViolation) GetUserAction() AnyOfmicrosoftGraphChatMessagePolicyViolationUserActionTypes`

GetUserAction returns the UserAction field if non-nil, zero value otherwise.

### GetUserActionOk

`func (o *MicrosoftGraphChatMessagePolicyViolation) GetUserActionOk() (*AnyOfmicrosoftGraphChatMessagePolicyViolationUserActionTypes, bool)`

GetUserActionOk returns a tuple with the UserAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAction

`func (o *MicrosoftGraphChatMessagePolicyViolation) SetUserAction(v AnyOfmicrosoftGraphChatMessagePolicyViolationUserActionTypes)`

SetUserAction sets UserAction field to given value.

### HasUserAction

`func (o *MicrosoftGraphChatMessagePolicyViolation) HasUserAction() bool`

HasUserAction returns a boolean if a field has been set.

### SetUserActionNil

`func (o *MicrosoftGraphChatMessagePolicyViolation) SetUserActionNil(b bool)`

 SetUserActionNil sets the value for UserAction to be an explicit nil

### UnsetUserAction
`func (o *MicrosoftGraphChatMessagePolicyViolation) UnsetUserAction()`

UnsetUserAction ensures that no value is present for UserAction, not even an explicit nil
### GetVerdictDetails

`func (o *MicrosoftGraphChatMessagePolicyViolation) GetVerdictDetails() AnyOfmicrosoftGraphChatMessagePolicyViolationVerdictDetailsTypes`

GetVerdictDetails returns the VerdictDetails field if non-nil, zero value otherwise.

### GetVerdictDetailsOk

`func (o *MicrosoftGraphChatMessagePolicyViolation) GetVerdictDetailsOk() (*AnyOfmicrosoftGraphChatMessagePolicyViolationVerdictDetailsTypes, bool)`

GetVerdictDetailsOk returns a tuple with the VerdictDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerdictDetails

`func (o *MicrosoftGraphChatMessagePolicyViolation) SetVerdictDetails(v AnyOfmicrosoftGraphChatMessagePolicyViolationVerdictDetailsTypes)`

SetVerdictDetails sets VerdictDetails field to given value.

### HasVerdictDetails

`func (o *MicrosoftGraphChatMessagePolicyViolation) HasVerdictDetails() bool`

HasVerdictDetails returns a boolean if a field has been set.

### SetVerdictDetailsNil

`func (o *MicrosoftGraphChatMessagePolicyViolation) SetVerdictDetailsNil(b bool)`

 SetVerdictDetailsNil sets the value for VerdictDetails to be an explicit nil

### UnsetVerdictDetails
`func (o *MicrosoftGraphChatMessagePolicyViolation) UnsetVerdictDetails()`

UnsetVerdictDetails ensures that no value is present for VerdictDetails, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


