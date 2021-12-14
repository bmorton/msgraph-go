# MicrosoftGraphChat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**ChatType** | Pointer to [**AnyOfmicrosoftGraphChatType**](anyOf&lt;microsoft.graph.chatType&gt;.md) | Specifies the type of chat. Possible values are: group, oneOnOne, meeting, unknownFutureValue. | [optional] 
**CreatedDateTime** | Pointer to **NullableTime** | Date and time at which the chat was created. Read-only. | [optional] 
**LastUpdatedDateTime** | Pointer to **NullableTime** | Date and time at which the chat was renamed or list of members were last changed. Read-only. | [optional] 
**Topic** | Pointer to **NullableString** | (Optional) Subject or topic for the chat. Only available for group chats. | [optional] 
**InstalledApps** | Pointer to [**[]MicrosoftGraphTeamsAppInstallation**](MicrosoftGraphTeamsAppInstallation.md) | A collection of all the apps in the chat. Nullable. | [optional] 
**Members** | Pointer to [**[]MicrosoftGraphConversationMember**](MicrosoftGraphConversationMember.md) | A collection of all the members in the chat. Nullable. | [optional] 
**Messages** | Pointer to [**[]MicrosoftGraphChatMessage**](MicrosoftGraphChatMessage.md) | A collection of all the messages in the chat. Nullable. | [optional] 
**Tabs** | Pointer to [**[]MicrosoftGraphTeamsTab**](MicrosoftGraphTeamsTab.md) |  | [optional] 

## Methods

### NewMicrosoftGraphChat

`func NewMicrosoftGraphChat() *MicrosoftGraphChat`

NewMicrosoftGraphChat instantiates a new MicrosoftGraphChat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphChatWithDefaults

`func NewMicrosoftGraphChatWithDefaults() *MicrosoftGraphChat`

NewMicrosoftGraphChatWithDefaults instantiates a new MicrosoftGraphChat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphChat) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphChat) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphChat) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphChat) HasId() bool`

HasId returns a boolean if a field has been set.

### GetChatType

`func (o *MicrosoftGraphChat) GetChatType() AnyOfmicrosoftGraphChatType`

GetChatType returns the ChatType field if non-nil, zero value otherwise.

### GetChatTypeOk

`func (o *MicrosoftGraphChat) GetChatTypeOk() (*AnyOfmicrosoftGraphChatType, bool)`

GetChatTypeOk returns a tuple with the ChatType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatType

`func (o *MicrosoftGraphChat) SetChatType(v AnyOfmicrosoftGraphChatType)`

SetChatType sets ChatType field to given value.

### HasChatType

`func (o *MicrosoftGraphChat) HasChatType() bool`

HasChatType returns a boolean if a field has been set.

### SetChatTypeNil

`func (o *MicrosoftGraphChat) SetChatTypeNil(b bool)`

 SetChatTypeNil sets the value for ChatType to be an explicit nil

### UnsetChatType
`func (o *MicrosoftGraphChat) UnsetChatType()`

UnsetChatType ensures that no value is present for ChatType, not even an explicit nil
### GetCreatedDateTime

`func (o *MicrosoftGraphChat) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *MicrosoftGraphChat) GetCreatedDateTimeOk() (*time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateTime

`func (o *MicrosoftGraphChat) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime sets CreatedDateTime field to given value.

### HasCreatedDateTime

`func (o *MicrosoftGraphChat) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTimeNil

`func (o *MicrosoftGraphChat) SetCreatedDateTimeNil(b bool)`

 SetCreatedDateTimeNil sets the value for CreatedDateTime to be an explicit nil

### UnsetCreatedDateTime
`func (o *MicrosoftGraphChat) UnsetCreatedDateTime()`

UnsetCreatedDateTime ensures that no value is present for CreatedDateTime, not even an explicit nil
### GetLastUpdatedDateTime

`func (o *MicrosoftGraphChat) GetLastUpdatedDateTime() time.Time`

GetLastUpdatedDateTime returns the LastUpdatedDateTime field if non-nil, zero value otherwise.

### GetLastUpdatedDateTimeOk

`func (o *MicrosoftGraphChat) GetLastUpdatedDateTimeOk() (*time.Time, bool)`

GetLastUpdatedDateTimeOk returns a tuple with the LastUpdatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedDateTime

`func (o *MicrosoftGraphChat) SetLastUpdatedDateTime(v time.Time)`

SetLastUpdatedDateTime sets LastUpdatedDateTime field to given value.

### HasLastUpdatedDateTime

`func (o *MicrosoftGraphChat) HasLastUpdatedDateTime() bool`

HasLastUpdatedDateTime returns a boolean if a field has been set.

### SetLastUpdatedDateTimeNil

`func (o *MicrosoftGraphChat) SetLastUpdatedDateTimeNil(b bool)`

 SetLastUpdatedDateTimeNil sets the value for LastUpdatedDateTime to be an explicit nil

### UnsetLastUpdatedDateTime
`func (o *MicrosoftGraphChat) UnsetLastUpdatedDateTime()`

UnsetLastUpdatedDateTime ensures that no value is present for LastUpdatedDateTime, not even an explicit nil
### GetTopic

`func (o *MicrosoftGraphChat) GetTopic() string`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *MicrosoftGraphChat) GetTopicOk() (*string, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *MicrosoftGraphChat) SetTopic(v string)`

SetTopic sets Topic field to given value.

### HasTopic

`func (o *MicrosoftGraphChat) HasTopic() bool`

HasTopic returns a boolean if a field has been set.

### SetTopicNil

`func (o *MicrosoftGraphChat) SetTopicNil(b bool)`

 SetTopicNil sets the value for Topic to be an explicit nil

### UnsetTopic
`func (o *MicrosoftGraphChat) UnsetTopic()`

UnsetTopic ensures that no value is present for Topic, not even an explicit nil
### GetInstalledApps

`func (o *MicrosoftGraphChat) GetInstalledApps() []MicrosoftGraphTeamsAppInstallation`

GetInstalledApps returns the InstalledApps field if non-nil, zero value otherwise.

### GetInstalledAppsOk

`func (o *MicrosoftGraphChat) GetInstalledAppsOk() (*[]MicrosoftGraphTeamsAppInstallation, bool)`

GetInstalledAppsOk returns a tuple with the InstalledApps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstalledApps

`func (o *MicrosoftGraphChat) SetInstalledApps(v []MicrosoftGraphTeamsAppInstallation)`

SetInstalledApps sets InstalledApps field to given value.

### HasInstalledApps

`func (o *MicrosoftGraphChat) HasInstalledApps() bool`

HasInstalledApps returns a boolean if a field has been set.

### GetMembers

`func (o *MicrosoftGraphChat) GetMembers() []MicrosoftGraphConversationMember`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *MicrosoftGraphChat) GetMembersOk() (*[]MicrosoftGraphConversationMember, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *MicrosoftGraphChat) SetMembers(v []MicrosoftGraphConversationMember)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *MicrosoftGraphChat) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### GetMessages

`func (o *MicrosoftGraphChat) GetMessages() []MicrosoftGraphChatMessage`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *MicrosoftGraphChat) GetMessagesOk() (*[]MicrosoftGraphChatMessage, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *MicrosoftGraphChat) SetMessages(v []MicrosoftGraphChatMessage)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *MicrosoftGraphChat) HasMessages() bool`

HasMessages returns a boolean if a field has been set.

### GetTabs

`func (o *MicrosoftGraphChat) GetTabs() []MicrosoftGraphTeamsTab`

GetTabs returns the Tabs field if non-nil, zero value otherwise.

### GetTabsOk

`func (o *MicrosoftGraphChat) GetTabsOk() (*[]MicrosoftGraphTeamsTab, bool)`

GetTabsOk returns a tuple with the Tabs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTabs

`func (o *MicrosoftGraphChat) SetTabs(v []MicrosoftGraphTeamsTab)`

SetTabs sets Tabs field to given value.

### HasTabs

`func (o *MicrosoftGraphChat) HasTabs() bool`

HasTabs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


