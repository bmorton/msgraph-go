# Chat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChatType** | Pointer to [**AnyOfmicrosoftGraphChatType**](anyOf&lt;microsoft.graph.chatType&gt;.md) | Specifies the type of chat. Possible values are: group, oneOnOne, meeting, unknownFutureValue. | [optional] 
**CreatedDateTime** | Pointer to **NullableTime** | Date and time at which the chat was created. Read-only. | [optional] 
**LastUpdatedDateTime** | Pointer to **NullableTime** | Date and time at which the chat was renamed or list of members were last changed. Read-only. | [optional] 
**Topic** | Pointer to **NullableString** | (Optional) Subject or topic for the chat. Only available for group chats. | [optional] 
**InstalledApps** | Pointer to [**[]MicrosoftGraphTeamsAppInstallation**](MicrosoftGraphTeamsAppInstallation.md) | A collection of all the apps in the chat. Nullable. | [optional] 
**Members** | Pointer to [**[]MicrosoftGraphConversationMember**](MicrosoftGraphConversationMember.md) | A collection of all the members in the chat. Nullable. | [optional] 
**Messages** | Pointer to [**[]MicrosoftGraphChatMessage**](MicrosoftGraphChatMessage.md) | A collection of all the messages in the chat. Nullable. | [optional] 
**Tabs** | Pointer to [**[]MicrosoftGraphTeamsTab**](MicrosoftGraphTeamsTab.md) |  | [optional] 

## Methods

### NewChat

`func NewChat() *Chat`

NewChat instantiates a new Chat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatWithDefaults

`func NewChatWithDefaults() *Chat`

NewChatWithDefaults instantiates a new Chat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChatType

`func (o *Chat) GetChatType() AnyOfmicrosoftGraphChatType`

GetChatType returns the ChatType field if non-nil, zero value otherwise.

### GetChatTypeOk

`func (o *Chat) GetChatTypeOk() (*AnyOfmicrosoftGraphChatType, bool)`

GetChatTypeOk returns a tuple with the ChatType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatType

`func (o *Chat) SetChatType(v AnyOfmicrosoftGraphChatType)`

SetChatType sets ChatType field to given value.

### HasChatType

`func (o *Chat) HasChatType() bool`

HasChatType returns a boolean if a field has been set.

### SetChatTypeNil

`func (o *Chat) SetChatTypeNil(b bool)`

 SetChatTypeNil sets the value for ChatType to be an explicit nil

### UnsetChatType
`func (o *Chat) UnsetChatType()`

UnsetChatType ensures that no value is present for ChatType, not even an explicit nil
### GetCreatedDateTime

`func (o *Chat) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *Chat) GetCreatedDateTimeOk() (*time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateTime

`func (o *Chat) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime sets CreatedDateTime field to given value.

### HasCreatedDateTime

`func (o *Chat) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTimeNil

`func (o *Chat) SetCreatedDateTimeNil(b bool)`

 SetCreatedDateTimeNil sets the value for CreatedDateTime to be an explicit nil

### UnsetCreatedDateTime
`func (o *Chat) UnsetCreatedDateTime()`

UnsetCreatedDateTime ensures that no value is present for CreatedDateTime, not even an explicit nil
### GetLastUpdatedDateTime

`func (o *Chat) GetLastUpdatedDateTime() time.Time`

GetLastUpdatedDateTime returns the LastUpdatedDateTime field if non-nil, zero value otherwise.

### GetLastUpdatedDateTimeOk

`func (o *Chat) GetLastUpdatedDateTimeOk() (*time.Time, bool)`

GetLastUpdatedDateTimeOk returns a tuple with the LastUpdatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedDateTime

`func (o *Chat) SetLastUpdatedDateTime(v time.Time)`

SetLastUpdatedDateTime sets LastUpdatedDateTime field to given value.

### HasLastUpdatedDateTime

`func (o *Chat) HasLastUpdatedDateTime() bool`

HasLastUpdatedDateTime returns a boolean if a field has been set.

### SetLastUpdatedDateTimeNil

`func (o *Chat) SetLastUpdatedDateTimeNil(b bool)`

 SetLastUpdatedDateTimeNil sets the value for LastUpdatedDateTime to be an explicit nil

### UnsetLastUpdatedDateTime
`func (o *Chat) UnsetLastUpdatedDateTime()`

UnsetLastUpdatedDateTime ensures that no value is present for LastUpdatedDateTime, not even an explicit nil
### GetTopic

`func (o *Chat) GetTopic() string`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *Chat) GetTopicOk() (*string, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *Chat) SetTopic(v string)`

SetTopic sets Topic field to given value.

### HasTopic

`func (o *Chat) HasTopic() bool`

HasTopic returns a boolean if a field has been set.

### SetTopicNil

`func (o *Chat) SetTopicNil(b bool)`

 SetTopicNil sets the value for Topic to be an explicit nil

### UnsetTopic
`func (o *Chat) UnsetTopic()`

UnsetTopic ensures that no value is present for Topic, not even an explicit nil
### GetInstalledApps

`func (o *Chat) GetInstalledApps() []MicrosoftGraphTeamsAppInstallation`

GetInstalledApps returns the InstalledApps field if non-nil, zero value otherwise.

### GetInstalledAppsOk

`func (o *Chat) GetInstalledAppsOk() (*[]MicrosoftGraphTeamsAppInstallation, bool)`

GetInstalledAppsOk returns a tuple with the InstalledApps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstalledApps

`func (o *Chat) SetInstalledApps(v []MicrosoftGraphTeamsAppInstallation)`

SetInstalledApps sets InstalledApps field to given value.

### HasInstalledApps

`func (o *Chat) HasInstalledApps() bool`

HasInstalledApps returns a boolean if a field has been set.

### GetMembers

`func (o *Chat) GetMembers() []MicrosoftGraphConversationMember`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *Chat) GetMembersOk() (*[]MicrosoftGraphConversationMember, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *Chat) SetMembers(v []MicrosoftGraphConversationMember)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *Chat) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### GetMessages

`func (o *Chat) GetMessages() []MicrosoftGraphChatMessage`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *Chat) GetMessagesOk() (*[]MicrosoftGraphChatMessage, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *Chat) SetMessages(v []MicrosoftGraphChatMessage)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *Chat) HasMessages() bool`

HasMessages returns a boolean if a field has been set.

### GetTabs

`func (o *Chat) GetTabs() []MicrosoftGraphTeamsTab`

GetTabs returns the Tabs field if non-nil, zero value otherwise.

### GetTabsOk

`func (o *Chat) GetTabsOk() (*[]MicrosoftGraphTeamsTab, bool)`

GetTabsOk returns a tuple with the Tabs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTabs

`func (o *Chat) SetTabs(v []MicrosoftGraphTeamsTab)`

SetTabs sets Tabs field to given value.

### HasTabs

`func (o *Chat) HasTabs() bool`

HasTabs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


