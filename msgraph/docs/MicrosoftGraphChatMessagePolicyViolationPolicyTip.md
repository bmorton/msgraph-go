# MicrosoftGraphChatMessagePolicyViolationPolicyTip

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ComplianceUrl** | Pointer to **NullableString** | The URL a user can visit to read about the data loss prevention policies for the organization. (ie, policies about what users shouldn&#39;t say in chats) | [optional] 
**GeneralText** | Pointer to **NullableString** | Explanatory text shown to the sender of the message. | [optional] 
**MatchedConditionDescriptions** | Pointer to **[]string** | The list of improper data in the message that was detected by the data loss prevention app. Each DLP app defines its own conditions, examples include &#39;Credit Card Number&#39; and &#39;Social Security Number&#39;. | [optional] 

## Methods

### NewMicrosoftGraphChatMessagePolicyViolationPolicyTip

`func NewMicrosoftGraphChatMessagePolicyViolationPolicyTip() *MicrosoftGraphChatMessagePolicyViolationPolicyTip`

NewMicrosoftGraphChatMessagePolicyViolationPolicyTip instantiates a new MicrosoftGraphChatMessagePolicyViolationPolicyTip object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphChatMessagePolicyViolationPolicyTipWithDefaults

`func NewMicrosoftGraphChatMessagePolicyViolationPolicyTipWithDefaults() *MicrosoftGraphChatMessagePolicyViolationPolicyTip`

NewMicrosoftGraphChatMessagePolicyViolationPolicyTipWithDefaults instantiates a new MicrosoftGraphChatMessagePolicyViolationPolicyTip object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComplianceUrl

`func (o *MicrosoftGraphChatMessagePolicyViolationPolicyTip) GetComplianceUrl() string`

GetComplianceUrl returns the ComplianceUrl field if non-nil, zero value otherwise.

### GetComplianceUrlOk

`func (o *MicrosoftGraphChatMessagePolicyViolationPolicyTip) GetComplianceUrlOk() (*string, bool)`

GetComplianceUrlOk returns a tuple with the ComplianceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplianceUrl

`func (o *MicrosoftGraphChatMessagePolicyViolationPolicyTip) SetComplianceUrl(v string)`

SetComplianceUrl sets ComplianceUrl field to given value.

### HasComplianceUrl

`func (o *MicrosoftGraphChatMessagePolicyViolationPolicyTip) HasComplianceUrl() bool`

HasComplianceUrl returns a boolean if a field has been set.

### SetComplianceUrlNil

`func (o *MicrosoftGraphChatMessagePolicyViolationPolicyTip) SetComplianceUrlNil(b bool)`

 SetComplianceUrlNil sets the value for ComplianceUrl to be an explicit nil

### UnsetComplianceUrl
`func (o *MicrosoftGraphChatMessagePolicyViolationPolicyTip) UnsetComplianceUrl()`

UnsetComplianceUrl ensures that no value is present for ComplianceUrl, not even an explicit nil
### GetGeneralText

`func (o *MicrosoftGraphChatMessagePolicyViolationPolicyTip) GetGeneralText() string`

GetGeneralText returns the GeneralText field if non-nil, zero value otherwise.

### GetGeneralTextOk

`func (o *MicrosoftGraphChatMessagePolicyViolationPolicyTip) GetGeneralTextOk() (*string, bool)`

GetGeneralTextOk returns a tuple with the GeneralText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneralText

`func (o *MicrosoftGraphChatMessagePolicyViolationPolicyTip) SetGeneralText(v string)`

SetGeneralText sets GeneralText field to given value.

### HasGeneralText

`func (o *MicrosoftGraphChatMessagePolicyViolationPolicyTip) HasGeneralText() bool`

HasGeneralText returns a boolean if a field has been set.

### SetGeneralTextNil

`func (o *MicrosoftGraphChatMessagePolicyViolationPolicyTip) SetGeneralTextNil(b bool)`

 SetGeneralTextNil sets the value for GeneralText to be an explicit nil

### UnsetGeneralText
`func (o *MicrosoftGraphChatMessagePolicyViolationPolicyTip) UnsetGeneralText()`

UnsetGeneralText ensures that no value is present for GeneralText, not even an explicit nil
### GetMatchedConditionDescriptions

`func (o *MicrosoftGraphChatMessagePolicyViolationPolicyTip) GetMatchedConditionDescriptions() []*string`

GetMatchedConditionDescriptions returns the MatchedConditionDescriptions field if non-nil, zero value otherwise.

### GetMatchedConditionDescriptionsOk

`func (o *MicrosoftGraphChatMessagePolicyViolationPolicyTip) GetMatchedConditionDescriptionsOk() (*[]*string, bool)`

GetMatchedConditionDescriptionsOk returns a tuple with the MatchedConditionDescriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchedConditionDescriptions

`func (o *MicrosoftGraphChatMessagePolicyViolationPolicyTip) SetMatchedConditionDescriptions(v []*string)`

SetMatchedConditionDescriptions sets MatchedConditionDescriptions field to given value.

### HasMatchedConditionDescriptions

`func (o *MicrosoftGraphChatMessagePolicyViolationPolicyTip) HasMatchedConditionDescriptions() bool`

HasMatchedConditionDescriptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


