# Team

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Classification** | Pointer to **NullableString** | An optional label. Typically describes the data or business sensitivity of the team. Must match one of a pre-configured set in the tenant&#39;s directory. | [optional] 
**CreatedDateTime** | Pointer to **NullableTime** | Timestamp at which the team was created. | [optional] 
**Description** | Pointer to **NullableString** | An optional description for the team. Maximum length: 1024 characters. | [optional] 
**DisplayName** | Pointer to **NullableString** | The name of the team. | [optional] 
**FunSettings** | Pointer to [**AnyOfmicrosoftGraphTeamFunSettings**](anyOf&lt;microsoft.graph.teamFunSettings&gt;.md) | Settings to configure use of Giphy, memes, and stickers in the team. | [optional] 
**GuestSettings** | Pointer to [**AnyOfmicrosoftGraphTeamGuestSettings**](anyOf&lt;microsoft.graph.teamGuestSettings&gt;.md) | Settings to configure whether guests can create, update, or delete channels in the team. | [optional] 
**InternalId** | Pointer to **NullableString** | A unique ID for the team that has been used in a few places such as the audit log/Office 365 Management Activity API. | [optional] 
**IsArchived** | Pointer to **NullableBool** | Whether this team is in read-only mode. | [optional] 
**MemberSettings** | Pointer to [**AnyOfmicrosoftGraphTeamMemberSettings**](anyOf&lt;microsoft.graph.teamMemberSettings&gt;.md) | Settings to configure whether members can perform certain actions, for example, create channels and add bots, in the team. | [optional] 
**MessagingSettings** | Pointer to [**AnyOfmicrosoftGraphTeamMessagingSettings**](anyOf&lt;microsoft.graph.teamMessagingSettings&gt;.md) | Settings to configure messaging and mentions in the team. | [optional] 
**Specialization** | Pointer to [**AnyOfmicrosoftGraphTeamSpecialization**](anyOf&lt;microsoft.graph.teamSpecialization&gt;.md) | Optional. Indicates whether the team is intended for a particular use case.  Each team specialization has access to unique behaviors and experiences targeted to its use case. | [optional] 
**Visibility** | Pointer to [**AnyOfmicrosoftGraphTeamVisibilityType**](anyOf&lt;microsoft.graph.teamVisibilityType&gt;.md) | The visibility of the group and team. Defaults to Public. | [optional] 
**WebUrl** | Pointer to **NullableString** | A hyperlink that will go to the team in the Microsoft Teams client. This is the URL that you get when you right-click a team in the Microsoft Teams client and select Get link to team. This URL should be treated as an opaque blob, and not parsed. | [optional] 
**Channels** | Pointer to [**[]MicrosoftGraphChannel**](MicrosoftGraphChannel.md) | The collection of channels &amp; messages associated with the team. | [optional] 
**Group** | Pointer to [**AnyOfmicrosoftGraphGroup**](anyOf&lt;microsoft.graph.group&gt;.md) |  | [optional] 
**InstalledApps** | Pointer to [**[]MicrosoftGraphTeamsAppInstallation**](MicrosoftGraphTeamsAppInstallation.md) | The apps installed in this team. | [optional] 
**Members** | Pointer to [**[]MicrosoftGraphConversationMember**](MicrosoftGraphConversationMember.md) | Members and owners of the team. | [optional] 
**Operations** | Pointer to [**[]MicrosoftGraphTeamsAsyncOperation**](MicrosoftGraphTeamsAsyncOperation.md) | The async operations that ran or are running on this team. | [optional] 
**PrimaryChannel** | Pointer to [**AnyOfmicrosoftGraphChannel**](anyOf&lt;microsoft.graph.channel&gt;.md) | The general channel for the team. | [optional] 
**Template** | Pointer to [**AnyOfmicrosoftGraphTeamsTemplate**](anyOf&lt;microsoft.graph.teamsTemplate&gt;.md) | The template this team was created from. See available templates. | [optional] 
**Schedule** | Pointer to [**AnyOfmicrosoftGraphSchedule**](anyOf&lt;microsoft.graph.schedule&gt;.md) | The schedule of shifts for this team. | [optional] 

## Methods

### NewTeam

`func NewTeam() *Team`

NewTeam instantiates a new Team object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamWithDefaults

`func NewTeamWithDefaults() *Team`

NewTeamWithDefaults instantiates a new Team object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassification

`func (o *Team) GetClassification() string`

GetClassification returns the Classification field if non-nil, zero value otherwise.

### GetClassificationOk

`func (o *Team) GetClassificationOk() (*string, bool)`

GetClassificationOk returns a tuple with the Classification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassification

`func (o *Team) SetClassification(v string)`

SetClassification sets Classification field to given value.

### HasClassification

`func (o *Team) HasClassification() bool`

HasClassification returns a boolean if a field has been set.

### SetClassificationNil

`func (o *Team) SetClassificationNil(b bool)`

 SetClassificationNil sets the value for Classification to be an explicit nil

### UnsetClassification
`func (o *Team) UnsetClassification()`

UnsetClassification ensures that no value is present for Classification, not even an explicit nil
### GetCreatedDateTime

`func (o *Team) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *Team) GetCreatedDateTimeOk() (*time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateTime

`func (o *Team) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime sets CreatedDateTime field to given value.

### HasCreatedDateTime

`func (o *Team) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTimeNil

`func (o *Team) SetCreatedDateTimeNil(b bool)`

 SetCreatedDateTimeNil sets the value for CreatedDateTime to be an explicit nil

### UnsetCreatedDateTime
`func (o *Team) UnsetCreatedDateTime()`

UnsetCreatedDateTime ensures that no value is present for CreatedDateTime, not even an explicit nil
### GetDescription

`func (o *Team) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Team) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Team) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Team) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *Team) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Team) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDisplayName

`func (o *Team) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *Team) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *Team) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *Team) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *Team) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *Team) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetFunSettings

`func (o *Team) GetFunSettings() AnyOfmicrosoftGraphTeamFunSettings`

GetFunSettings returns the FunSettings field if non-nil, zero value otherwise.

### GetFunSettingsOk

`func (o *Team) GetFunSettingsOk() (*AnyOfmicrosoftGraphTeamFunSettings, bool)`

GetFunSettingsOk returns a tuple with the FunSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunSettings

`func (o *Team) SetFunSettings(v AnyOfmicrosoftGraphTeamFunSettings)`

SetFunSettings sets FunSettings field to given value.

### HasFunSettings

`func (o *Team) HasFunSettings() bool`

HasFunSettings returns a boolean if a field has been set.

### SetFunSettingsNil

`func (o *Team) SetFunSettingsNil(b bool)`

 SetFunSettingsNil sets the value for FunSettings to be an explicit nil

### UnsetFunSettings
`func (o *Team) UnsetFunSettings()`

UnsetFunSettings ensures that no value is present for FunSettings, not even an explicit nil
### GetGuestSettings

`func (o *Team) GetGuestSettings() AnyOfmicrosoftGraphTeamGuestSettings`

GetGuestSettings returns the GuestSettings field if non-nil, zero value otherwise.

### GetGuestSettingsOk

`func (o *Team) GetGuestSettingsOk() (*AnyOfmicrosoftGraphTeamGuestSettings, bool)`

GetGuestSettingsOk returns a tuple with the GuestSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestSettings

`func (o *Team) SetGuestSettings(v AnyOfmicrosoftGraphTeamGuestSettings)`

SetGuestSettings sets GuestSettings field to given value.

### HasGuestSettings

`func (o *Team) HasGuestSettings() bool`

HasGuestSettings returns a boolean if a field has been set.

### SetGuestSettingsNil

`func (o *Team) SetGuestSettingsNil(b bool)`

 SetGuestSettingsNil sets the value for GuestSettings to be an explicit nil

### UnsetGuestSettings
`func (o *Team) UnsetGuestSettings()`

UnsetGuestSettings ensures that no value is present for GuestSettings, not even an explicit nil
### GetInternalId

`func (o *Team) GetInternalId() string`

GetInternalId returns the InternalId field if non-nil, zero value otherwise.

### GetInternalIdOk

`func (o *Team) GetInternalIdOk() (*string, bool)`

GetInternalIdOk returns a tuple with the InternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalId

`func (o *Team) SetInternalId(v string)`

SetInternalId sets InternalId field to given value.

### HasInternalId

`func (o *Team) HasInternalId() bool`

HasInternalId returns a boolean if a field has been set.

### SetInternalIdNil

`func (o *Team) SetInternalIdNil(b bool)`

 SetInternalIdNil sets the value for InternalId to be an explicit nil

### UnsetInternalId
`func (o *Team) UnsetInternalId()`

UnsetInternalId ensures that no value is present for InternalId, not even an explicit nil
### GetIsArchived

`func (o *Team) GetIsArchived() bool`

GetIsArchived returns the IsArchived field if non-nil, zero value otherwise.

### GetIsArchivedOk

`func (o *Team) GetIsArchivedOk() (*bool, bool)`

GetIsArchivedOk returns a tuple with the IsArchived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsArchived

`func (o *Team) SetIsArchived(v bool)`

SetIsArchived sets IsArchived field to given value.

### HasIsArchived

`func (o *Team) HasIsArchived() bool`

HasIsArchived returns a boolean if a field has been set.

### SetIsArchivedNil

`func (o *Team) SetIsArchivedNil(b bool)`

 SetIsArchivedNil sets the value for IsArchived to be an explicit nil

### UnsetIsArchived
`func (o *Team) UnsetIsArchived()`

UnsetIsArchived ensures that no value is present for IsArchived, not even an explicit nil
### GetMemberSettings

`func (o *Team) GetMemberSettings() AnyOfmicrosoftGraphTeamMemberSettings`

GetMemberSettings returns the MemberSettings field if non-nil, zero value otherwise.

### GetMemberSettingsOk

`func (o *Team) GetMemberSettingsOk() (*AnyOfmicrosoftGraphTeamMemberSettings, bool)`

GetMemberSettingsOk returns a tuple with the MemberSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberSettings

`func (o *Team) SetMemberSettings(v AnyOfmicrosoftGraphTeamMemberSettings)`

SetMemberSettings sets MemberSettings field to given value.

### HasMemberSettings

`func (o *Team) HasMemberSettings() bool`

HasMemberSettings returns a boolean if a field has been set.

### SetMemberSettingsNil

`func (o *Team) SetMemberSettingsNil(b bool)`

 SetMemberSettingsNil sets the value for MemberSettings to be an explicit nil

### UnsetMemberSettings
`func (o *Team) UnsetMemberSettings()`

UnsetMemberSettings ensures that no value is present for MemberSettings, not even an explicit nil
### GetMessagingSettings

`func (o *Team) GetMessagingSettings() AnyOfmicrosoftGraphTeamMessagingSettings`

GetMessagingSettings returns the MessagingSettings field if non-nil, zero value otherwise.

### GetMessagingSettingsOk

`func (o *Team) GetMessagingSettingsOk() (*AnyOfmicrosoftGraphTeamMessagingSettings, bool)`

GetMessagingSettingsOk returns a tuple with the MessagingSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessagingSettings

`func (o *Team) SetMessagingSettings(v AnyOfmicrosoftGraphTeamMessagingSettings)`

SetMessagingSettings sets MessagingSettings field to given value.

### HasMessagingSettings

`func (o *Team) HasMessagingSettings() bool`

HasMessagingSettings returns a boolean if a field has been set.

### SetMessagingSettingsNil

`func (o *Team) SetMessagingSettingsNil(b bool)`

 SetMessagingSettingsNil sets the value for MessagingSettings to be an explicit nil

### UnsetMessagingSettings
`func (o *Team) UnsetMessagingSettings()`

UnsetMessagingSettings ensures that no value is present for MessagingSettings, not even an explicit nil
### GetSpecialization

`func (o *Team) GetSpecialization() AnyOfmicrosoftGraphTeamSpecialization`

GetSpecialization returns the Specialization field if non-nil, zero value otherwise.

### GetSpecializationOk

`func (o *Team) GetSpecializationOk() (*AnyOfmicrosoftGraphTeamSpecialization, bool)`

GetSpecializationOk returns a tuple with the Specialization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialization

`func (o *Team) SetSpecialization(v AnyOfmicrosoftGraphTeamSpecialization)`

SetSpecialization sets Specialization field to given value.

### HasSpecialization

`func (o *Team) HasSpecialization() bool`

HasSpecialization returns a boolean if a field has been set.

### SetSpecializationNil

`func (o *Team) SetSpecializationNil(b bool)`

 SetSpecializationNil sets the value for Specialization to be an explicit nil

### UnsetSpecialization
`func (o *Team) UnsetSpecialization()`

UnsetSpecialization ensures that no value is present for Specialization, not even an explicit nil
### GetVisibility

`func (o *Team) GetVisibility() AnyOfmicrosoftGraphTeamVisibilityType`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *Team) GetVisibilityOk() (*AnyOfmicrosoftGraphTeamVisibilityType, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *Team) SetVisibility(v AnyOfmicrosoftGraphTeamVisibilityType)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *Team) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### SetVisibilityNil

`func (o *Team) SetVisibilityNil(b bool)`

 SetVisibilityNil sets the value for Visibility to be an explicit nil

### UnsetVisibility
`func (o *Team) UnsetVisibility()`

UnsetVisibility ensures that no value is present for Visibility, not even an explicit nil
### GetWebUrl

`func (o *Team) GetWebUrl() string`

GetWebUrl returns the WebUrl field if non-nil, zero value otherwise.

### GetWebUrlOk

`func (o *Team) GetWebUrlOk() (*string, bool)`

GetWebUrlOk returns a tuple with the WebUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebUrl

`func (o *Team) SetWebUrl(v string)`

SetWebUrl sets WebUrl field to given value.

### HasWebUrl

`func (o *Team) HasWebUrl() bool`

HasWebUrl returns a boolean if a field has been set.

### SetWebUrlNil

`func (o *Team) SetWebUrlNil(b bool)`

 SetWebUrlNil sets the value for WebUrl to be an explicit nil

### UnsetWebUrl
`func (o *Team) UnsetWebUrl()`

UnsetWebUrl ensures that no value is present for WebUrl, not even an explicit nil
### GetChannels

`func (o *Team) GetChannels() []MicrosoftGraphChannel`

GetChannels returns the Channels field if non-nil, zero value otherwise.

### GetChannelsOk

`func (o *Team) GetChannelsOk() (*[]MicrosoftGraphChannel, bool)`

GetChannelsOk returns a tuple with the Channels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannels

`func (o *Team) SetChannels(v []MicrosoftGraphChannel)`

SetChannels sets Channels field to given value.

### HasChannels

`func (o *Team) HasChannels() bool`

HasChannels returns a boolean if a field has been set.

### GetGroup

`func (o *Team) GetGroup() AnyOfmicrosoftGraphGroup`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *Team) GetGroupOk() (*AnyOfmicrosoftGraphGroup, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *Team) SetGroup(v AnyOfmicrosoftGraphGroup)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *Team) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### SetGroupNil

`func (o *Team) SetGroupNil(b bool)`

 SetGroupNil sets the value for Group to be an explicit nil

### UnsetGroup
`func (o *Team) UnsetGroup()`

UnsetGroup ensures that no value is present for Group, not even an explicit nil
### GetInstalledApps

`func (o *Team) GetInstalledApps() []MicrosoftGraphTeamsAppInstallation`

GetInstalledApps returns the InstalledApps field if non-nil, zero value otherwise.

### GetInstalledAppsOk

`func (o *Team) GetInstalledAppsOk() (*[]MicrosoftGraphTeamsAppInstallation, bool)`

GetInstalledAppsOk returns a tuple with the InstalledApps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstalledApps

`func (o *Team) SetInstalledApps(v []MicrosoftGraphTeamsAppInstallation)`

SetInstalledApps sets InstalledApps field to given value.

### HasInstalledApps

`func (o *Team) HasInstalledApps() bool`

HasInstalledApps returns a boolean if a field has been set.

### GetMembers

`func (o *Team) GetMembers() []MicrosoftGraphConversationMember`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *Team) GetMembersOk() (*[]MicrosoftGraphConversationMember, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *Team) SetMembers(v []MicrosoftGraphConversationMember)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *Team) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### GetOperations

`func (o *Team) GetOperations() []MicrosoftGraphTeamsAsyncOperation`

GetOperations returns the Operations field if non-nil, zero value otherwise.

### GetOperationsOk

`func (o *Team) GetOperationsOk() (*[]MicrosoftGraphTeamsAsyncOperation, bool)`

GetOperationsOk returns a tuple with the Operations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperations

`func (o *Team) SetOperations(v []MicrosoftGraphTeamsAsyncOperation)`

SetOperations sets Operations field to given value.

### HasOperations

`func (o *Team) HasOperations() bool`

HasOperations returns a boolean if a field has been set.

### GetPrimaryChannel

`func (o *Team) GetPrimaryChannel() AnyOfmicrosoftGraphChannel`

GetPrimaryChannel returns the PrimaryChannel field if non-nil, zero value otherwise.

### GetPrimaryChannelOk

`func (o *Team) GetPrimaryChannelOk() (*AnyOfmicrosoftGraphChannel, bool)`

GetPrimaryChannelOk returns a tuple with the PrimaryChannel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryChannel

`func (o *Team) SetPrimaryChannel(v AnyOfmicrosoftGraphChannel)`

SetPrimaryChannel sets PrimaryChannel field to given value.

### HasPrimaryChannel

`func (o *Team) HasPrimaryChannel() bool`

HasPrimaryChannel returns a boolean if a field has been set.

### SetPrimaryChannelNil

`func (o *Team) SetPrimaryChannelNil(b bool)`

 SetPrimaryChannelNil sets the value for PrimaryChannel to be an explicit nil

### UnsetPrimaryChannel
`func (o *Team) UnsetPrimaryChannel()`

UnsetPrimaryChannel ensures that no value is present for PrimaryChannel, not even an explicit nil
### GetTemplate

`func (o *Team) GetTemplate() AnyOfmicrosoftGraphTeamsTemplate`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *Team) GetTemplateOk() (*AnyOfmicrosoftGraphTeamsTemplate, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *Team) SetTemplate(v AnyOfmicrosoftGraphTeamsTemplate)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *Team) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### SetTemplateNil

`func (o *Team) SetTemplateNil(b bool)`

 SetTemplateNil sets the value for Template to be an explicit nil

### UnsetTemplate
`func (o *Team) UnsetTemplate()`

UnsetTemplate ensures that no value is present for Template, not even an explicit nil
### GetSchedule

`func (o *Team) GetSchedule() AnyOfmicrosoftGraphSchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *Team) GetScheduleOk() (*AnyOfmicrosoftGraphSchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *Team) SetSchedule(v AnyOfmicrosoftGraphSchedule)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *Team) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### SetScheduleNil

`func (o *Team) SetScheduleNil(b bool)`

 SetScheduleNil sets the value for Schedule to be an explicit nil

### UnsetSchedule
`func (o *Team) UnsetSchedule()`

UnsetSchedule ensures that no value is present for Schedule, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


