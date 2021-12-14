# MailFolder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChildFolderCount** | Pointer to **NullableInt32** | The number of immediate child mailFolders in the current mailFolder. | [optional] 
**DisplayName** | Pointer to **NullableString** | The mailFolder&#39;s display name. | [optional] 
**IsHidden** | Pointer to **NullableBool** | Indicates whether the mailFolder is hidden. This property can be set only when creating the folder. Find more information in Hidden mail folders. | [optional] 
**ParentFolderId** | Pointer to **NullableString** | The unique identifier for the mailFolder&#39;s parent mailFolder. | [optional] 
**TotalItemCount** | Pointer to **NullableInt32** | The number of items in the mailFolder. | [optional] 
**UnreadItemCount** | Pointer to **NullableInt32** | The number of items in the mailFolder marked as unread. | [optional] 
**ChildFolders** | Pointer to [**[]MicrosoftGraphMailFolder**](MicrosoftGraphMailFolder.md) | The collection of child folders in the mailFolder. | [optional] 
**MessageRules** | Pointer to [**[]MicrosoftGraphMessageRule**](MicrosoftGraphMessageRule.md) | The collection of rules that apply to the user&#39;s Inbox folder. | [optional] 
**Messages** | Pointer to [**[]MicrosoftGraphMessage**](MicrosoftGraphMessage.md) | The collection of messages in the mailFolder. | [optional] 
**MultiValueExtendedProperties** | Pointer to [**[]MicrosoftGraphMultiValueLegacyExtendedProperty**](MicrosoftGraphMultiValueLegacyExtendedProperty.md) | The collection of multi-value extended properties defined for the mailFolder. Read-only. Nullable. | [optional] 
**SingleValueExtendedProperties** | Pointer to [**[]MicrosoftGraphSingleValueLegacyExtendedProperty**](MicrosoftGraphSingleValueLegacyExtendedProperty.md) | The collection of single-value extended properties defined for the mailFolder. Read-only. Nullable. | [optional] 

## Methods

### NewMailFolder

`func NewMailFolder() *MailFolder`

NewMailFolder instantiates a new MailFolder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMailFolderWithDefaults

`func NewMailFolderWithDefaults() *MailFolder`

NewMailFolderWithDefaults instantiates a new MailFolder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChildFolderCount

`func (o *MailFolder) GetChildFolderCount() int32`

GetChildFolderCount returns the ChildFolderCount field if non-nil, zero value otherwise.

### GetChildFolderCountOk

`func (o *MailFolder) GetChildFolderCountOk() (*int32, bool)`

GetChildFolderCountOk returns a tuple with the ChildFolderCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildFolderCount

`func (o *MailFolder) SetChildFolderCount(v int32)`

SetChildFolderCount sets ChildFolderCount field to given value.

### HasChildFolderCount

`func (o *MailFolder) HasChildFolderCount() bool`

HasChildFolderCount returns a boolean if a field has been set.

### SetChildFolderCountNil

`func (o *MailFolder) SetChildFolderCountNil(b bool)`

 SetChildFolderCountNil sets the value for ChildFolderCount to be an explicit nil

### UnsetChildFolderCount
`func (o *MailFolder) UnsetChildFolderCount()`

UnsetChildFolderCount ensures that no value is present for ChildFolderCount, not even an explicit nil
### GetDisplayName

`func (o *MailFolder) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MailFolder) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *MailFolder) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *MailFolder) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *MailFolder) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *MailFolder) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetIsHidden

`func (o *MailFolder) GetIsHidden() bool`

GetIsHidden returns the IsHidden field if non-nil, zero value otherwise.

### GetIsHiddenOk

`func (o *MailFolder) GetIsHiddenOk() (*bool, bool)`

GetIsHiddenOk returns a tuple with the IsHidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHidden

`func (o *MailFolder) SetIsHidden(v bool)`

SetIsHidden sets IsHidden field to given value.

### HasIsHidden

`func (o *MailFolder) HasIsHidden() bool`

HasIsHidden returns a boolean if a field has been set.

### SetIsHiddenNil

`func (o *MailFolder) SetIsHiddenNil(b bool)`

 SetIsHiddenNil sets the value for IsHidden to be an explicit nil

### UnsetIsHidden
`func (o *MailFolder) UnsetIsHidden()`

UnsetIsHidden ensures that no value is present for IsHidden, not even an explicit nil
### GetParentFolderId

`func (o *MailFolder) GetParentFolderId() string`

GetParentFolderId returns the ParentFolderId field if non-nil, zero value otherwise.

### GetParentFolderIdOk

`func (o *MailFolder) GetParentFolderIdOk() (*string, bool)`

GetParentFolderIdOk returns a tuple with the ParentFolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentFolderId

`func (o *MailFolder) SetParentFolderId(v string)`

SetParentFolderId sets ParentFolderId field to given value.

### HasParentFolderId

`func (o *MailFolder) HasParentFolderId() bool`

HasParentFolderId returns a boolean if a field has been set.

### SetParentFolderIdNil

`func (o *MailFolder) SetParentFolderIdNil(b bool)`

 SetParentFolderIdNil sets the value for ParentFolderId to be an explicit nil

### UnsetParentFolderId
`func (o *MailFolder) UnsetParentFolderId()`

UnsetParentFolderId ensures that no value is present for ParentFolderId, not even an explicit nil
### GetTotalItemCount

`func (o *MailFolder) GetTotalItemCount() int32`

GetTotalItemCount returns the TotalItemCount field if non-nil, zero value otherwise.

### GetTotalItemCountOk

`func (o *MailFolder) GetTotalItemCountOk() (*int32, bool)`

GetTotalItemCountOk returns a tuple with the TotalItemCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalItemCount

`func (o *MailFolder) SetTotalItemCount(v int32)`

SetTotalItemCount sets TotalItemCount field to given value.

### HasTotalItemCount

`func (o *MailFolder) HasTotalItemCount() bool`

HasTotalItemCount returns a boolean if a field has been set.

### SetTotalItemCountNil

`func (o *MailFolder) SetTotalItemCountNil(b bool)`

 SetTotalItemCountNil sets the value for TotalItemCount to be an explicit nil

### UnsetTotalItemCount
`func (o *MailFolder) UnsetTotalItemCount()`

UnsetTotalItemCount ensures that no value is present for TotalItemCount, not even an explicit nil
### GetUnreadItemCount

`func (o *MailFolder) GetUnreadItemCount() int32`

GetUnreadItemCount returns the UnreadItemCount field if non-nil, zero value otherwise.

### GetUnreadItemCountOk

`func (o *MailFolder) GetUnreadItemCountOk() (*int32, bool)`

GetUnreadItemCountOk returns a tuple with the UnreadItemCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnreadItemCount

`func (o *MailFolder) SetUnreadItemCount(v int32)`

SetUnreadItemCount sets UnreadItemCount field to given value.

### HasUnreadItemCount

`func (o *MailFolder) HasUnreadItemCount() bool`

HasUnreadItemCount returns a boolean if a field has been set.

### SetUnreadItemCountNil

`func (o *MailFolder) SetUnreadItemCountNil(b bool)`

 SetUnreadItemCountNil sets the value for UnreadItemCount to be an explicit nil

### UnsetUnreadItemCount
`func (o *MailFolder) UnsetUnreadItemCount()`

UnsetUnreadItemCount ensures that no value is present for UnreadItemCount, not even an explicit nil
### GetChildFolders

`func (o *MailFolder) GetChildFolders() []MicrosoftGraphMailFolder`

GetChildFolders returns the ChildFolders field if non-nil, zero value otherwise.

### GetChildFoldersOk

`func (o *MailFolder) GetChildFoldersOk() (*[]MicrosoftGraphMailFolder, bool)`

GetChildFoldersOk returns a tuple with the ChildFolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildFolders

`func (o *MailFolder) SetChildFolders(v []MicrosoftGraphMailFolder)`

SetChildFolders sets ChildFolders field to given value.

### HasChildFolders

`func (o *MailFolder) HasChildFolders() bool`

HasChildFolders returns a boolean if a field has been set.

### GetMessageRules

`func (o *MailFolder) GetMessageRules() []MicrosoftGraphMessageRule`

GetMessageRules returns the MessageRules field if non-nil, zero value otherwise.

### GetMessageRulesOk

`func (o *MailFolder) GetMessageRulesOk() (*[]MicrosoftGraphMessageRule, bool)`

GetMessageRulesOk returns a tuple with the MessageRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageRules

`func (o *MailFolder) SetMessageRules(v []MicrosoftGraphMessageRule)`

SetMessageRules sets MessageRules field to given value.

### HasMessageRules

`func (o *MailFolder) HasMessageRules() bool`

HasMessageRules returns a boolean if a field has been set.

### GetMessages

`func (o *MailFolder) GetMessages() []MicrosoftGraphMessage`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *MailFolder) GetMessagesOk() (*[]MicrosoftGraphMessage, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *MailFolder) SetMessages(v []MicrosoftGraphMessage)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *MailFolder) HasMessages() bool`

HasMessages returns a boolean if a field has been set.

### GetMultiValueExtendedProperties

`func (o *MailFolder) GetMultiValueExtendedProperties() []MicrosoftGraphMultiValueLegacyExtendedProperty`

GetMultiValueExtendedProperties returns the MultiValueExtendedProperties field if non-nil, zero value otherwise.

### GetMultiValueExtendedPropertiesOk

`func (o *MailFolder) GetMultiValueExtendedPropertiesOk() (*[]MicrosoftGraphMultiValueLegacyExtendedProperty, bool)`

GetMultiValueExtendedPropertiesOk returns a tuple with the MultiValueExtendedProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiValueExtendedProperties

`func (o *MailFolder) SetMultiValueExtendedProperties(v []MicrosoftGraphMultiValueLegacyExtendedProperty)`

SetMultiValueExtendedProperties sets MultiValueExtendedProperties field to given value.

### HasMultiValueExtendedProperties

`func (o *MailFolder) HasMultiValueExtendedProperties() bool`

HasMultiValueExtendedProperties returns a boolean if a field has been set.

### GetSingleValueExtendedProperties

`func (o *MailFolder) GetSingleValueExtendedProperties() []MicrosoftGraphSingleValueLegacyExtendedProperty`

GetSingleValueExtendedProperties returns the SingleValueExtendedProperties field if non-nil, zero value otherwise.

### GetSingleValueExtendedPropertiesOk

`func (o *MailFolder) GetSingleValueExtendedPropertiesOk() (*[]MicrosoftGraphSingleValueLegacyExtendedProperty, bool)`

GetSingleValueExtendedPropertiesOk returns a tuple with the SingleValueExtendedProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSingleValueExtendedProperties

`func (o *MailFolder) SetSingleValueExtendedProperties(v []MicrosoftGraphSingleValueLegacyExtendedProperty)`

SetSingleValueExtendedProperties sets SingleValueExtendedProperties field to given value.

### HasSingleValueExtendedProperties

`func (o *MailFolder) HasSingleValueExtendedProperties() bool`

HasSingleValueExtendedProperties returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


