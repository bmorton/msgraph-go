# MicrosoftGraphChannel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**CreatedDateTime** | Pointer to **NullableTime** | Read only. Timestamp at which the channel was created. | [optional] 
**Description** | Pointer to **NullableString** | Optional textual description for the channel. | [optional] 
**DisplayName** | Pointer to **string** | Channel name as it will appear to the user in Microsoft Teams. | [optional] 
**Email** | Pointer to **NullableString** | The email address for sending messages to the channel. Read-only. | [optional] 
**IsFavoriteByDefault** | Pointer to **NullableBool** | Indicates whether the channel should automatically be marked &#39;favorite&#39; for all members of the team. Can only be set programmatically with Create team. Default: false. | [optional] 
**MembershipType** | Pointer to [**AnyOfmicrosoftGraphChannelMembershipType**](anyOf&lt;microsoft.graph.channelMembershipType&gt;.md) | The type of the channel. Can be set during creation and can&#39;t be changed. Possible values are: standard - Channel inherits the list of members of the parent team; private - Channel can have members that are a subset of all the members on the parent team. | [optional] 
**WebUrl** | Pointer to **NullableString** | A hyperlink that will go to the channel in Microsoft Teams. This is the URL that you get when you right-click a channel in Microsoft Teams and select Get link to channel. This URL should be treated as an opaque blob, and not parsed. Read-only. | [optional] 
**FilesFolder** | Pointer to [**AnyOfmicrosoftGraphDriveItem**](anyOf&lt;microsoft.graph.driveItem&gt;.md) | Metadata for the location where the channel&#39;s files are stored. | [optional] 
**Members** | Pointer to [**[]MicrosoftGraphConversationMember**](MicrosoftGraphConversationMember.md) | A collection of membership records associated with the channel. | [optional] 
**Messages** | Pointer to [**[]MicrosoftGraphChatMessage**](MicrosoftGraphChatMessage.md) | A collection of all the messages in the channel. A navigation property. Nullable. | [optional] 
**Tabs** | Pointer to [**[]MicrosoftGraphTeamsTab**](MicrosoftGraphTeamsTab.md) | A collection of all the tabs in the channel. A navigation property. | [optional] 

## Methods

### NewMicrosoftGraphChannel

`func NewMicrosoftGraphChannel() *MicrosoftGraphChannel`

NewMicrosoftGraphChannel instantiates a new MicrosoftGraphChannel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphChannelWithDefaults

`func NewMicrosoftGraphChannelWithDefaults() *MicrosoftGraphChannel`

NewMicrosoftGraphChannelWithDefaults instantiates a new MicrosoftGraphChannel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphChannel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphChannel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphChannel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphChannel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedDateTime

`func (o *MicrosoftGraphChannel) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *MicrosoftGraphChannel) GetCreatedDateTimeOk() (*time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateTime

`func (o *MicrosoftGraphChannel) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime sets CreatedDateTime field to given value.

### HasCreatedDateTime

`func (o *MicrosoftGraphChannel) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTimeNil

`func (o *MicrosoftGraphChannel) SetCreatedDateTimeNil(b bool)`

 SetCreatedDateTimeNil sets the value for CreatedDateTime to be an explicit nil

### UnsetCreatedDateTime
`func (o *MicrosoftGraphChannel) UnsetCreatedDateTime()`

UnsetCreatedDateTime ensures that no value is present for CreatedDateTime, not even an explicit nil
### GetDescription

`func (o *MicrosoftGraphChannel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MicrosoftGraphChannel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MicrosoftGraphChannel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MicrosoftGraphChannel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *MicrosoftGraphChannel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *MicrosoftGraphChannel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDisplayName

`func (o *MicrosoftGraphChannel) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphChannel) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *MicrosoftGraphChannel) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *MicrosoftGraphChannel) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetEmail

`func (o *MicrosoftGraphChannel) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *MicrosoftGraphChannel) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *MicrosoftGraphChannel) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *MicrosoftGraphChannel) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *MicrosoftGraphChannel) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *MicrosoftGraphChannel) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetIsFavoriteByDefault

`func (o *MicrosoftGraphChannel) GetIsFavoriteByDefault() bool`

GetIsFavoriteByDefault returns the IsFavoriteByDefault field if non-nil, zero value otherwise.

### GetIsFavoriteByDefaultOk

`func (o *MicrosoftGraphChannel) GetIsFavoriteByDefaultOk() (*bool, bool)`

GetIsFavoriteByDefaultOk returns a tuple with the IsFavoriteByDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFavoriteByDefault

`func (o *MicrosoftGraphChannel) SetIsFavoriteByDefault(v bool)`

SetIsFavoriteByDefault sets IsFavoriteByDefault field to given value.

### HasIsFavoriteByDefault

`func (o *MicrosoftGraphChannel) HasIsFavoriteByDefault() bool`

HasIsFavoriteByDefault returns a boolean if a field has been set.

### SetIsFavoriteByDefaultNil

`func (o *MicrosoftGraphChannel) SetIsFavoriteByDefaultNil(b bool)`

 SetIsFavoriteByDefaultNil sets the value for IsFavoriteByDefault to be an explicit nil

### UnsetIsFavoriteByDefault
`func (o *MicrosoftGraphChannel) UnsetIsFavoriteByDefault()`

UnsetIsFavoriteByDefault ensures that no value is present for IsFavoriteByDefault, not even an explicit nil
### GetMembershipType

`func (o *MicrosoftGraphChannel) GetMembershipType() AnyOfmicrosoftGraphChannelMembershipType`

GetMembershipType returns the MembershipType field if non-nil, zero value otherwise.

### GetMembershipTypeOk

`func (o *MicrosoftGraphChannel) GetMembershipTypeOk() (*AnyOfmicrosoftGraphChannelMembershipType, bool)`

GetMembershipTypeOk returns a tuple with the MembershipType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembershipType

`func (o *MicrosoftGraphChannel) SetMembershipType(v AnyOfmicrosoftGraphChannelMembershipType)`

SetMembershipType sets MembershipType field to given value.

### HasMembershipType

`func (o *MicrosoftGraphChannel) HasMembershipType() bool`

HasMembershipType returns a boolean if a field has been set.

### SetMembershipTypeNil

`func (o *MicrosoftGraphChannel) SetMembershipTypeNil(b bool)`

 SetMembershipTypeNil sets the value for MembershipType to be an explicit nil

### UnsetMembershipType
`func (o *MicrosoftGraphChannel) UnsetMembershipType()`

UnsetMembershipType ensures that no value is present for MembershipType, not even an explicit nil
### GetWebUrl

`func (o *MicrosoftGraphChannel) GetWebUrl() string`

GetWebUrl returns the WebUrl field if non-nil, zero value otherwise.

### GetWebUrlOk

`func (o *MicrosoftGraphChannel) GetWebUrlOk() (*string, bool)`

GetWebUrlOk returns a tuple with the WebUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebUrl

`func (o *MicrosoftGraphChannel) SetWebUrl(v string)`

SetWebUrl sets WebUrl field to given value.

### HasWebUrl

`func (o *MicrosoftGraphChannel) HasWebUrl() bool`

HasWebUrl returns a boolean if a field has been set.

### SetWebUrlNil

`func (o *MicrosoftGraphChannel) SetWebUrlNil(b bool)`

 SetWebUrlNil sets the value for WebUrl to be an explicit nil

### UnsetWebUrl
`func (o *MicrosoftGraphChannel) UnsetWebUrl()`

UnsetWebUrl ensures that no value is present for WebUrl, not even an explicit nil
### GetFilesFolder

`func (o *MicrosoftGraphChannel) GetFilesFolder() AnyOfmicrosoftGraphDriveItem`

GetFilesFolder returns the FilesFolder field if non-nil, zero value otherwise.

### GetFilesFolderOk

`func (o *MicrosoftGraphChannel) GetFilesFolderOk() (*AnyOfmicrosoftGraphDriveItem, bool)`

GetFilesFolderOk returns a tuple with the FilesFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilesFolder

`func (o *MicrosoftGraphChannel) SetFilesFolder(v AnyOfmicrosoftGraphDriveItem)`

SetFilesFolder sets FilesFolder field to given value.

### HasFilesFolder

`func (o *MicrosoftGraphChannel) HasFilesFolder() bool`

HasFilesFolder returns a boolean if a field has been set.

### SetFilesFolderNil

`func (o *MicrosoftGraphChannel) SetFilesFolderNil(b bool)`

 SetFilesFolderNil sets the value for FilesFolder to be an explicit nil

### UnsetFilesFolder
`func (o *MicrosoftGraphChannel) UnsetFilesFolder()`

UnsetFilesFolder ensures that no value is present for FilesFolder, not even an explicit nil
### GetMembers

`func (o *MicrosoftGraphChannel) GetMembers() []MicrosoftGraphConversationMember`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *MicrosoftGraphChannel) GetMembersOk() (*[]MicrosoftGraphConversationMember, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *MicrosoftGraphChannel) SetMembers(v []MicrosoftGraphConversationMember)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *MicrosoftGraphChannel) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### GetMessages

`func (o *MicrosoftGraphChannel) GetMessages() []MicrosoftGraphChatMessage`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *MicrosoftGraphChannel) GetMessagesOk() (*[]MicrosoftGraphChatMessage, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *MicrosoftGraphChannel) SetMessages(v []MicrosoftGraphChatMessage)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *MicrosoftGraphChannel) HasMessages() bool`

HasMessages returns a boolean if a field has been set.

### GetTabs

`func (o *MicrosoftGraphChannel) GetTabs() []MicrosoftGraphTeamsTab`

GetTabs returns the Tabs field if non-nil, zero value otherwise.

### GetTabsOk

`func (o *MicrosoftGraphChannel) GetTabsOk() (*[]MicrosoftGraphTeamsTab, bool)`

GetTabsOk returns a tuple with the Tabs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTabs

`func (o *MicrosoftGraphChannel) SetTabs(v []MicrosoftGraphTeamsTab)`

SetTabs sets Tabs field to given value.

### HasTabs

`func (o *MicrosoftGraphChannel) HasTabs() bool`

HasTabs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


