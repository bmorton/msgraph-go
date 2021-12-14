# MicrosoftGraphMessageRuleActions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssignCategories** | Pointer to **[]string** | A list of categories to be assigned to a message. | [optional] 
**CopyToFolder** | Pointer to **NullableString** | The ID of a folder that a message is to be copied to. | [optional] 
**Delete** | Pointer to **NullableBool** | Indicates whether a message should be moved to the Deleted Items folder. | [optional] 
**ForwardAsAttachmentTo** | Pointer to [**[]AnyOfmicrosoftGraphRecipient**](AnyOfmicrosoftGraphRecipient.md) | The email addresses of the recipients to which a message should be forwarded as an attachment. | [optional] 
**ForwardTo** | Pointer to [**[]AnyOfmicrosoftGraphRecipient**](AnyOfmicrosoftGraphRecipient.md) | The email addresses of the recipients to which a message should be forwarded. | [optional] 
**MarkAsRead** | Pointer to **NullableBool** | Indicates whether a message should be marked as read. | [optional] 
**MarkImportance** | Pointer to [**AnyOfmicrosoftGraphImportance**](anyOf&lt;microsoft.graph.importance&gt;.md) | Sets the importance of the message, which can be: low, normal, high. | [optional] 
**MoveToFolder** | Pointer to **NullableString** | The ID of the folder that a message will be moved to. | [optional] 
**PermanentDelete** | Pointer to **NullableBool** | Indicates whether a message should be permanently deleted and not saved to the Deleted Items folder. | [optional] 
**RedirectTo** | Pointer to [**[]AnyOfmicrosoftGraphRecipient**](AnyOfmicrosoftGraphRecipient.md) | The email addresses to which a message should be redirected. | [optional] 
**StopProcessingRules** | Pointer to **NullableBool** | Indicates whether subsequent rules should be evaluated. | [optional] 

## Methods

### NewMicrosoftGraphMessageRuleActions

`func NewMicrosoftGraphMessageRuleActions() *MicrosoftGraphMessageRuleActions`

NewMicrosoftGraphMessageRuleActions instantiates a new MicrosoftGraphMessageRuleActions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphMessageRuleActionsWithDefaults

`func NewMicrosoftGraphMessageRuleActionsWithDefaults() *MicrosoftGraphMessageRuleActions`

NewMicrosoftGraphMessageRuleActionsWithDefaults instantiates a new MicrosoftGraphMessageRuleActions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssignCategories

`func (o *MicrosoftGraphMessageRuleActions) GetAssignCategories() []*string`

GetAssignCategories returns the AssignCategories field if non-nil, zero value otherwise.

### GetAssignCategoriesOk

`func (o *MicrosoftGraphMessageRuleActions) GetAssignCategoriesOk() (*[]*string, bool)`

GetAssignCategoriesOk returns a tuple with the AssignCategories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignCategories

`func (o *MicrosoftGraphMessageRuleActions) SetAssignCategories(v []*string)`

SetAssignCategories sets AssignCategories field to given value.

### HasAssignCategories

`func (o *MicrosoftGraphMessageRuleActions) HasAssignCategories() bool`

HasAssignCategories returns a boolean if a field has been set.

### GetCopyToFolder

`func (o *MicrosoftGraphMessageRuleActions) GetCopyToFolder() string`

GetCopyToFolder returns the CopyToFolder field if non-nil, zero value otherwise.

### GetCopyToFolderOk

`func (o *MicrosoftGraphMessageRuleActions) GetCopyToFolderOk() (*string, bool)`

GetCopyToFolderOk returns a tuple with the CopyToFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyToFolder

`func (o *MicrosoftGraphMessageRuleActions) SetCopyToFolder(v string)`

SetCopyToFolder sets CopyToFolder field to given value.

### HasCopyToFolder

`func (o *MicrosoftGraphMessageRuleActions) HasCopyToFolder() bool`

HasCopyToFolder returns a boolean if a field has been set.

### SetCopyToFolderNil

`func (o *MicrosoftGraphMessageRuleActions) SetCopyToFolderNil(b bool)`

 SetCopyToFolderNil sets the value for CopyToFolder to be an explicit nil

### UnsetCopyToFolder
`func (o *MicrosoftGraphMessageRuleActions) UnsetCopyToFolder()`

UnsetCopyToFolder ensures that no value is present for CopyToFolder, not even an explicit nil
### GetDelete

`func (o *MicrosoftGraphMessageRuleActions) GetDelete() bool`

GetDelete returns the Delete field if non-nil, zero value otherwise.

### GetDeleteOk

`func (o *MicrosoftGraphMessageRuleActions) GetDeleteOk() (*bool, bool)`

GetDeleteOk returns a tuple with the Delete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelete

`func (o *MicrosoftGraphMessageRuleActions) SetDelete(v bool)`

SetDelete sets Delete field to given value.

### HasDelete

`func (o *MicrosoftGraphMessageRuleActions) HasDelete() bool`

HasDelete returns a boolean if a field has been set.

### SetDeleteNil

`func (o *MicrosoftGraphMessageRuleActions) SetDeleteNil(b bool)`

 SetDeleteNil sets the value for Delete to be an explicit nil

### UnsetDelete
`func (o *MicrosoftGraphMessageRuleActions) UnsetDelete()`

UnsetDelete ensures that no value is present for Delete, not even an explicit nil
### GetForwardAsAttachmentTo

`func (o *MicrosoftGraphMessageRuleActions) GetForwardAsAttachmentTo() []*AnyOfmicrosoftGraphRecipient`

GetForwardAsAttachmentTo returns the ForwardAsAttachmentTo field if non-nil, zero value otherwise.

### GetForwardAsAttachmentToOk

`func (o *MicrosoftGraphMessageRuleActions) GetForwardAsAttachmentToOk() (*[]*AnyOfmicrosoftGraphRecipient, bool)`

GetForwardAsAttachmentToOk returns a tuple with the ForwardAsAttachmentTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardAsAttachmentTo

`func (o *MicrosoftGraphMessageRuleActions) SetForwardAsAttachmentTo(v []*AnyOfmicrosoftGraphRecipient)`

SetForwardAsAttachmentTo sets ForwardAsAttachmentTo field to given value.

### HasForwardAsAttachmentTo

`func (o *MicrosoftGraphMessageRuleActions) HasForwardAsAttachmentTo() bool`

HasForwardAsAttachmentTo returns a boolean if a field has been set.

### GetForwardTo

`func (o *MicrosoftGraphMessageRuleActions) GetForwardTo() []*AnyOfmicrosoftGraphRecipient`

GetForwardTo returns the ForwardTo field if non-nil, zero value otherwise.

### GetForwardToOk

`func (o *MicrosoftGraphMessageRuleActions) GetForwardToOk() (*[]*AnyOfmicrosoftGraphRecipient, bool)`

GetForwardToOk returns a tuple with the ForwardTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardTo

`func (o *MicrosoftGraphMessageRuleActions) SetForwardTo(v []*AnyOfmicrosoftGraphRecipient)`

SetForwardTo sets ForwardTo field to given value.

### HasForwardTo

`func (o *MicrosoftGraphMessageRuleActions) HasForwardTo() bool`

HasForwardTo returns a boolean if a field has been set.

### GetMarkAsRead

`func (o *MicrosoftGraphMessageRuleActions) GetMarkAsRead() bool`

GetMarkAsRead returns the MarkAsRead field if non-nil, zero value otherwise.

### GetMarkAsReadOk

`func (o *MicrosoftGraphMessageRuleActions) GetMarkAsReadOk() (*bool, bool)`

GetMarkAsReadOk returns a tuple with the MarkAsRead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkAsRead

`func (o *MicrosoftGraphMessageRuleActions) SetMarkAsRead(v bool)`

SetMarkAsRead sets MarkAsRead field to given value.

### HasMarkAsRead

`func (o *MicrosoftGraphMessageRuleActions) HasMarkAsRead() bool`

HasMarkAsRead returns a boolean if a field has been set.

### SetMarkAsReadNil

`func (o *MicrosoftGraphMessageRuleActions) SetMarkAsReadNil(b bool)`

 SetMarkAsReadNil sets the value for MarkAsRead to be an explicit nil

### UnsetMarkAsRead
`func (o *MicrosoftGraphMessageRuleActions) UnsetMarkAsRead()`

UnsetMarkAsRead ensures that no value is present for MarkAsRead, not even an explicit nil
### GetMarkImportance

`func (o *MicrosoftGraphMessageRuleActions) GetMarkImportance() AnyOfmicrosoftGraphImportance`

GetMarkImportance returns the MarkImportance field if non-nil, zero value otherwise.

### GetMarkImportanceOk

`func (o *MicrosoftGraphMessageRuleActions) GetMarkImportanceOk() (*AnyOfmicrosoftGraphImportance, bool)`

GetMarkImportanceOk returns a tuple with the MarkImportance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkImportance

`func (o *MicrosoftGraphMessageRuleActions) SetMarkImportance(v AnyOfmicrosoftGraphImportance)`

SetMarkImportance sets MarkImportance field to given value.

### HasMarkImportance

`func (o *MicrosoftGraphMessageRuleActions) HasMarkImportance() bool`

HasMarkImportance returns a boolean if a field has been set.

### SetMarkImportanceNil

`func (o *MicrosoftGraphMessageRuleActions) SetMarkImportanceNil(b bool)`

 SetMarkImportanceNil sets the value for MarkImportance to be an explicit nil

### UnsetMarkImportance
`func (o *MicrosoftGraphMessageRuleActions) UnsetMarkImportance()`

UnsetMarkImportance ensures that no value is present for MarkImportance, not even an explicit nil
### GetMoveToFolder

`func (o *MicrosoftGraphMessageRuleActions) GetMoveToFolder() string`

GetMoveToFolder returns the MoveToFolder field if non-nil, zero value otherwise.

### GetMoveToFolderOk

`func (o *MicrosoftGraphMessageRuleActions) GetMoveToFolderOk() (*string, bool)`

GetMoveToFolderOk returns a tuple with the MoveToFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoveToFolder

`func (o *MicrosoftGraphMessageRuleActions) SetMoveToFolder(v string)`

SetMoveToFolder sets MoveToFolder field to given value.

### HasMoveToFolder

`func (o *MicrosoftGraphMessageRuleActions) HasMoveToFolder() bool`

HasMoveToFolder returns a boolean if a field has been set.

### SetMoveToFolderNil

`func (o *MicrosoftGraphMessageRuleActions) SetMoveToFolderNil(b bool)`

 SetMoveToFolderNil sets the value for MoveToFolder to be an explicit nil

### UnsetMoveToFolder
`func (o *MicrosoftGraphMessageRuleActions) UnsetMoveToFolder()`

UnsetMoveToFolder ensures that no value is present for MoveToFolder, not even an explicit nil
### GetPermanentDelete

`func (o *MicrosoftGraphMessageRuleActions) GetPermanentDelete() bool`

GetPermanentDelete returns the PermanentDelete field if non-nil, zero value otherwise.

### GetPermanentDeleteOk

`func (o *MicrosoftGraphMessageRuleActions) GetPermanentDeleteOk() (*bool, bool)`

GetPermanentDeleteOk returns a tuple with the PermanentDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermanentDelete

`func (o *MicrosoftGraphMessageRuleActions) SetPermanentDelete(v bool)`

SetPermanentDelete sets PermanentDelete field to given value.

### HasPermanentDelete

`func (o *MicrosoftGraphMessageRuleActions) HasPermanentDelete() bool`

HasPermanentDelete returns a boolean if a field has been set.

### SetPermanentDeleteNil

`func (o *MicrosoftGraphMessageRuleActions) SetPermanentDeleteNil(b bool)`

 SetPermanentDeleteNil sets the value for PermanentDelete to be an explicit nil

### UnsetPermanentDelete
`func (o *MicrosoftGraphMessageRuleActions) UnsetPermanentDelete()`

UnsetPermanentDelete ensures that no value is present for PermanentDelete, not even an explicit nil
### GetRedirectTo

`func (o *MicrosoftGraphMessageRuleActions) GetRedirectTo() []*AnyOfmicrosoftGraphRecipient`

GetRedirectTo returns the RedirectTo field if non-nil, zero value otherwise.

### GetRedirectToOk

`func (o *MicrosoftGraphMessageRuleActions) GetRedirectToOk() (*[]*AnyOfmicrosoftGraphRecipient, bool)`

GetRedirectToOk returns a tuple with the RedirectTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectTo

`func (o *MicrosoftGraphMessageRuleActions) SetRedirectTo(v []*AnyOfmicrosoftGraphRecipient)`

SetRedirectTo sets RedirectTo field to given value.

### HasRedirectTo

`func (o *MicrosoftGraphMessageRuleActions) HasRedirectTo() bool`

HasRedirectTo returns a boolean if a field has been set.

### GetStopProcessingRules

`func (o *MicrosoftGraphMessageRuleActions) GetStopProcessingRules() bool`

GetStopProcessingRules returns the StopProcessingRules field if non-nil, zero value otherwise.

### GetStopProcessingRulesOk

`func (o *MicrosoftGraphMessageRuleActions) GetStopProcessingRulesOk() (*bool, bool)`

GetStopProcessingRulesOk returns a tuple with the StopProcessingRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopProcessingRules

`func (o *MicrosoftGraphMessageRuleActions) SetStopProcessingRules(v bool)`

SetStopProcessingRules sets StopProcessingRules field to given value.

### HasStopProcessingRules

`func (o *MicrosoftGraphMessageRuleActions) HasStopProcessingRules() bool`

HasStopProcessingRules returns a boolean if a field has been set.

### SetStopProcessingRulesNil

`func (o *MicrosoftGraphMessageRuleActions) SetStopProcessingRulesNil(b bool)`

 SetStopProcessingRulesNil sets the value for StopProcessingRules to be an explicit nil

### UnsetStopProcessingRules
`func (o *MicrosoftGraphMessageRuleActions) UnsetStopProcessingRules()`

UnsetStopProcessingRules ensures that no value is present for StopProcessingRules, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


