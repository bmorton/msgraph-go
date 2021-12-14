# MicrosoftGraphTeam

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
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

### NewMicrosoftGraphTeam

`func NewMicrosoftGraphTeam() *MicrosoftGraphTeam`

NewMicrosoftGraphTeam instantiates a new MicrosoftGraphTeam object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphTeamWithDefaults

`func NewMicrosoftGraphTeamWithDefaults() *MicrosoftGraphTeam`

NewMicrosoftGraphTeamWithDefaults instantiates a new MicrosoftGraphTeam object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphTeam) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphTeam) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphTeam) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphTeam) HasId() bool`

HasId returns a boolean if a field has been set.

### GetClassification

`func (o *MicrosoftGraphTeam) GetClassification() string`

GetClassification returns the Classification field if non-nil, zero value otherwise.

### GetClassificationOk

`func (o *MicrosoftGraphTeam) GetClassificationOk() (*string, bool)`

GetClassificationOk returns a tuple with the Classification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassification

`func (o *MicrosoftGraphTeam) SetClassification(v string)`

SetClassification sets Classification field to given value.

### HasClassification

`func (o *MicrosoftGraphTeam) HasClassification() bool`

HasClassification returns a boolean if a field has been set.

### SetClassificationNil

`func (o *MicrosoftGraphTeam) SetClassificationNil(b bool)`

 SetClassificationNil sets the value for Classification to be an explicit nil

### UnsetClassification
`func (o *MicrosoftGraphTeam) UnsetClassification()`

UnsetClassification ensures that no value is present for Classification, not even an explicit nil
### GetCreatedDateTime

`func (o *MicrosoftGraphTeam) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *MicrosoftGraphTeam) GetCreatedDateTimeOk() (*time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateTime

`func (o *MicrosoftGraphTeam) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime sets CreatedDateTime field to given value.

### HasCreatedDateTime

`func (o *MicrosoftGraphTeam) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTimeNil

`func (o *MicrosoftGraphTeam) SetCreatedDateTimeNil(b bool)`

 SetCreatedDateTimeNil sets the value for CreatedDateTime to be an explicit nil

### UnsetCreatedDateTime
`func (o *MicrosoftGraphTeam) UnsetCreatedDateTime()`

UnsetCreatedDateTime ensures that no value is present for CreatedDateTime, not even an explicit nil
### GetDescription

`func (o *MicrosoftGraphTeam) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MicrosoftGraphTeam) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MicrosoftGraphTeam) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MicrosoftGraphTeam) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *MicrosoftGraphTeam) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *MicrosoftGraphTeam) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDisplayName

`func (o *MicrosoftGraphTeam) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphTeam) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *MicrosoftGraphTeam) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *MicrosoftGraphTeam) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *MicrosoftGraphTeam) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *MicrosoftGraphTeam) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetFunSettings

`func (o *MicrosoftGraphTeam) GetFunSettings() AnyOfmicrosoftGraphTeamFunSettings`

GetFunSettings returns the FunSettings field if non-nil, zero value otherwise.

### GetFunSettingsOk

`func (o *MicrosoftGraphTeam) GetFunSettingsOk() (*AnyOfmicrosoftGraphTeamFunSettings, bool)`

GetFunSettingsOk returns a tuple with the FunSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunSettings

`func (o *MicrosoftGraphTeam) SetFunSettings(v AnyOfmicrosoftGraphTeamFunSettings)`

SetFunSettings sets FunSettings field to given value.

### HasFunSettings

`func (o *MicrosoftGraphTeam) HasFunSettings() bool`

HasFunSettings returns a boolean if a field has been set.

### SetFunSettingsNil

`func (o *MicrosoftGraphTeam) SetFunSettingsNil(b bool)`

 SetFunSettingsNil sets the value for FunSettings to be an explicit nil

### UnsetFunSettings
`func (o *MicrosoftGraphTeam) UnsetFunSettings()`

UnsetFunSettings ensures that no value is present for FunSettings, not even an explicit nil
### GetGuestSettings

`func (o *MicrosoftGraphTeam) GetGuestSettings() AnyOfmicrosoftGraphTeamGuestSettings`

GetGuestSettings returns the GuestSettings field if non-nil, zero value otherwise.

### GetGuestSettingsOk

`func (o *MicrosoftGraphTeam) GetGuestSettingsOk() (*AnyOfmicrosoftGraphTeamGuestSettings, bool)`

GetGuestSettingsOk returns a tuple with the GuestSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestSettings

`func (o *MicrosoftGraphTeam) SetGuestSettings(v AnyOfmicrosoftGraphTeamGuestSettings)`

SetGuestSettings sets GuestSettings field to given value.

### HasGuestSettings

`func (o *MicrosoftGraphTeam) HasGuestSettings() bool`

HasGuestSettings returns a boolean if a field has been set.

### SetGuestSettingsNil

`func (o *MicrosoftGraphTeam) SetGuestSettingsNil(b bool)`

 SetGuestSettingsNil sets the value for GuestSettings to be an explicit nil

### UnsetGuestSettings
`func (o *MicrosoftGraphTeam) UnsetGuestSettings()`

UnsetGuestSettings ensures that no value is present for GuestSettings, not even an explicit nil
### GetInternalId

`func (o *MicrosoftGraphTeam) GetInternalId() string`

GetInternalId returns the InternalId field if non-nil, zero value otherwise.

### GetInternalIdOk

`func (o *MicrosoftGraphTeam) GetInternalIdOk() (*string, bool)`

GetInternalIdOk returns a tuple with the InternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalId

`func (o *MicrosoftGraphTeam) SetInternalId(v string)`

SetInternalId sets InternalId field to given value.

### HasInternalId

`func (o *MicrosoftGraphTeam) HasInternalId() bool`

HasInternalId returns a boolean if a field has been set.

### SetInternalIdNil

`func (o *MicrosoftGraphTeam) SetInternalIdNil(b bool)`

 SetInternalIdNil sets the value for InternalId to be an explicit nil

### UnsetInternalId
`func (o *MicrosoftGraphTeam) UnsetInternalId()`

UnsetInternalId ensures that no value is present for InternalId, not even an explicit nil
### GetIsArchived

`func (o *MicrosoftGraphTeam) GetIsArchived() bool`

GetIsArchived returns the IsArchived field if non-nil, zero value otherwise.

### GetIsArchivedOk

`func (o *MicrosoftGraphTeam) GetIsArchivedOk() (*bool, bool)`

GetIsArchivedOk returns a tuple with the IsArchived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsArchived

`func (o *MicrosoftGraphTeam) SetIsArchived(v bool)`

SetIsArchived sets IsArchived field to given value.

### HasIsArchived

`func (o *MicrosoftGraphTeam) HasIsArchived() bool`

HasIsArchived returns a boolean if a field has been set.

### SetIsArchivedNil

`func (o *MicrosoftGraphTeam) SetIsArchivedNil(b bool)`

 SetIsArchivedNil sets the value for IsArchived to be an explicit nil

### UnsetIsArchived
`func (o *MicrosoftGraphTeam) UnsetIsArchived()`

UnsetIsArchived ensures that no value is present for IsArchived, not even an explicit nil
### GetMemberSettings

`func (o *MicrosoftGraphTeam) GetMemberSettings() AnyOfmicrosoftGraphTeamMemberSettings`

GetMemberSettings returns the MemberSettings field if non-nil, zero value otherwise.

### GetMemberSettingsOk

`func (o *MicrosoftGraphTeam) GetMemberSettingsOk() (*AnyOfmicrosoftGraphTeamMemberSettings, bool)`

GetMemberSettingsOk returns a tuple with the MemberSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberSettings

`func (o *MicrosoftGraphTeam) SetMemberSettings(v AnyOfmicrosoftGraphTeamMemberSettings)`

SetMemberSettings sets MemberSettings field to given value.

### HasMemberSettings

`func (o *MicrosoftGraphTeam) HasMemberSettings() bool`

HasMemberSettings returns a boolean if a field has been set.

### SetMemberSettingsNil

`func (o *MicrosoftGraphTeam) SetMemberSettingsNil(b bool)`

 SetMemberSettingsNil sets the value for MemberSettings to be an explicit nil

### UnsetMemberSettings
`func (o *MicrosoftGraphTeam) UnsetMemberSettings()`

UnsetMemberSettings ensures that no value is present for MemberSettings, not even an explicit nil
### GetMessagingSettings

`func (o *MicrosoftGraphTeam) GetMessagingSettings() AnyOfmicrosoftGraphTeamMessagingSettings`

GetMessagingSettings returns the MessagingSettings field if non-nil, zero value otherwise.

### GetMessagingSettingsOk

`func (o *MicrosoftGraphTeam) GetMessagingSettingsOk() (*AnyOfmicrosoftGraphTeamMessagingSettings, bool)`

GetMessagingSettingsOk returns a tuple with the MessagingSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessagingSettings

`func (o *MicrosoftGraphTeam) SetMessagingSettings(v AnyOfmicrosoftGraphTeamMessagingSettings)`

SetMessagingSettings sets MessagingSettings field to given value.

### HasMessagingSettings

`func (o *MicrosoftGraphTeam) HasMessagingSettings() bool`

HasMessagingSettings returns a boolean if a field has been set.

### SetMessagingSettingsNil

`func (o *MicrosoftGraphTeam) SetMessagingSettingsNil(b bool)`

 SetMessagingSettingsNil sets the value for MessagingSettings to be an explicit nil

### UnsetMessagingSettings
`func (o *MicrosoftGraphTeam) UnsetMessagingSettings()`

UnsetMessagingSettings ensures that no value is present for MessagingSettings, not even an explicit nil
### GetSpecialization

`func (o *MicrosoftGraphTeam) GetSpecialization() AnyOfmicrosoftGraphTeamSpecialization`

GetSpecialization returns the Specialization field if non-nil, zero value otherwise.

### GetSpecializationOk

`func (o *MicrosoftGraphTeam) GetSpecializationOk() (*AnyOfmicrosoftGraphTeamSpecialization, bool)`

GetSpecializationOk returns a tuple with the Specialization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialization

`func (o *MicrosoftGraphTeam) SetSpecialization(v AnyOfmicrosoftGraphTeamSpecialization)`

SetSpecialization sets Specialization field to given value.

### HasSpecialization

`func (o *MicrosoftGraphTeam) HasSpecialization() bool`

HasSpecialization returns a boolean if a field has been set.

### SetSpecializationNil

`func (o *MicrosoftGraphTeam) SetSpecializationNil(b bool)`

 SetSpecializationNil sets the value for Specialization to be an explicit nil

### UnsetSpecialization
`func (o *MicrosoftGraphTeam) UnsetSpecialization()`

UnsetSpecialization ensures that no value is present for Specialization, not even an explicit nil
### GetVisibility

`func (o *MicrosoftGraphTeam) GetVisibility() AnyOfmicrosoftGraphTeamVisibilityType`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *MicrosoftGraphTeam) GetVisibilityOk() (*AnyOfmicrosoftGraphTeamVisibilityType, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *MicrosoftGraphTeam) SetVisibility(v AnyOfmicrosoftGraphTeamVisibilityType)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *MicrosoftGraphTeam) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### SetVisibilityNil

`func (o *MicrosoftGraphTeam) SetVisibilityNil(b bool)`

 SetVisibilityNil sets the value for Visibility to be an explicit nil

### UnsetVisibility
`func (o *MicrosoftGraphTeam) UnsetVisibility()`

UnsetVisibility ensures that no value is present for Visibility, not even an explicit nil
### GetWebUrl

`func (o *MicrosoftGraphTeam) GetWebUrl() string`

GetWebUrl returns the WebUrl field if non-nil, zero value otherwise.

### GetWebUrlOk

`func (o *MicrosoftGraphTeam) GetWebUrlOk() (*string, bool)`

GetWebUrlOk returns a tuple with the WebUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebUrl

`func (o *MicrosoftGraphTeam) SetWebUrl(v string)`

SetWebUrl sets WebUrl field to given value.

### HasWebUrl

`func (o *MicrosoftGraphTeam) HasWebUrl() bool`

HasWebUrl returns a boolean if a field has been set.

### SetWebUrlNil

`func (o *MicrosoftGraphTeam) SetWebUrlNil(b bool)`

 SetWebUrlNil sets the value for WebUrl to be an explicit nil

### UnsetWebUrl
`func (o *MicrosoftGraphTeam) UnsetWebUrl()`

UnsetWebUrl ensures that no value is present for WebUrl, not even an explicit nil
### GetChannels

`func (o *MicrosoftGraphTeam) GetChannels() []MicrosoftGraphChannel`

GetChannels returns the Channels field if non-nil, zero value otherwise.

### GetChannelsOk

`func (o *MicrosoftGraphTeam) GetChannelsOk() (*[]MicrosoftGraphChannel, bool)`

GetChannelsOk returns a tuple with the Channels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannels

`func (o *MicrosoftGraphTeam) SetChannels(v []MicrosoftGraphChannel)`

SetChannels sets Channels field to given value.

### HasChannels

`func (o *MicrosoftGraphTeam) HasChannels() bool`

HasChannels returns a boolean if a field has been set.

### GetGroup

`func (o *MicrosoftGraphTeam) GetGroup() AnyOfmicrosoftGraphGroup`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *MicrosoftGraphTeam) GetGroupOk() (*AnyOfmicrosoftGraphGroup, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *MicrosoftGraphTeam) SetGroup(v AnyOfmicrosoftGraphGroup)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *MicrosoftGraphTeam) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### SetGroupNil

`func (o *MicrosoftGraphTeam) SetGroupNil(b bool)`

 SetGroupNil sets the value for Group to be an explicit nil

### UnsetGroup
`func (o *MicrosoftGraphTeam) UnsetGroup()`

UnsetGroup ensures that no value is present for Group, not even an explicit nil
### GetInstalledApps

`func (o *MicrosoftGraphTeam) GetInstalledApps() []MicrosoftGraphTeamsAppInstallation`

GetInstalledApps returns the InstalledApps field if non-nil, zero value otherwise.

### GetInstalledAppsOk

`func (o *MicrosoftGraphTeam) GetInstalledAppsOk() (*[]MicrosoftGraphTeamsAppInstallation, bool)`

GetInstalledAppsOk returns a tuple with the InstalledApps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstalledApps

`func (o *MicrosoftGraphTeam) SetInstalledApps(v []MicrosoftGraphTeamsAppInstallation)`

SetInstalledApps sets InstalledApps field to given value.

### HasInstalledApps

`func (o *MicrosoftGraphTeam) HasInstalledApps() bool`

HasInstalledApps returns a boolean if a field has been set.

### GetMembers

`func (o *MicrosoftGraphTeam) GetMembers() []MicrosoftGraphConversationMember`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *MicrosoftGraphTeam) GetMembersOk() (*[]MicrosoftGraphConversationMember, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *MicrosoftGraphTeam) SetMembers(v []MicrosoftGraphConversationMember)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *MicrosoftGraphTeam) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### GetOperations

`func (o *MicrosoftGraphTeam) GetOperations() []MicrosoftGraphTeamsAsyncOperation`

GetOperations returns the Operations field if non-nil, zero value otherwise.

### GetOperationsOk

`func (o *MicrosoftGraphTeam) GetOperationsOk() (*[]MicrosoftGraphTeamsAsyncOperation, bool)`

GetOperationsOk returns a tuple with the Operations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperations

`func (o *MicrosoftGraphTeam) SetOperations(v []MicrosoftGraphTeamsAsyncOperation)`

SetOperations sets Operations field to given value.

### HasOperations

`func (o *MicrosoftGraphTeam) HasOperations() bool`

HasOperations returns a boolean if a field has been set.

### GetPrimaryChannel

`func (o *MicrosoftGraphTeam) GetPrimaryChannel() AnyOfmicrosoftGraphChannel`

GetPrimaryChannel returns the PrimaryChannel field if non-nil, zero value otherwise.

### GetPrimaryChannelOk

`func (o *MicrosoftGraphTeam) GetPrimaryChannelOk() (*AnyOfmicrosoftGraphChannel, bool)`

GetPrimaryChannelOk returns a tuple with the PrimaryChannel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryChannel

`func (o *MicrosoftGraphTeam) SetPrimaryChannel(v AnyOfmicrosoftGraphChannel)`

SetPrimaryChannel sets PrimaryChannel field to given value.

### HasPrimaryChannel

`func (o *MicrosoftGraphTeam) HasPrimaryChannel() bool`

HasPrimaryChannel returns a boolean if a field has been set.

### SetPrimaryChannelNil

`func (o *MicrosoftGraphTeam) SetPrimaryChannelNil(b bool)`

 SetPrimaryChannelNil sets the value for PrimaryChannel to be an explicit nil

### UnsetPrimaryChannel
`func (o *MicrosoftGraphTeam) UnsetPrimaryChannel()`

UnsetPrimaryChannel ensures that no value is present for PrimaryChannel, not even an explicit nil
### GetTemplate

`func (o *MicrosoftGraphTeam) GetTemplate() AnyOfmicrosoftGraphTeamsTemplate`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *MicrosoftGraphTeam) GetTemplateOk() (*AnyOfmicrosoftGraphTeamsTemplate, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *MicrosoftGraphTeam) SetTemplate(v AnyOfmicrosoftGraphTeamsTemplate)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *MicrosoftGraphTeam) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### SetTemplateNil

`func (o *MicrosoftGraphTeam) SetTemplateNil(b bool)`

 SetTemplateNil sets the value for Template to be an explicit nil

### UnsetTemplate
`func (o *MicrosoftGraphTeam) UnsetTemplate()`

UnsetTemplate ensures that no value is present for Template, not even an explicit nil
### GetSchedule

`func (o *MicrosoftGraphTeam) GetSchedule() AnyOfmicrosoftGraphSchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *MicrosoftGraphTeam) GetScheduleOk() (*AnyOfmicrosoftGraphSchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *MicrosoftGraphTeam) SetSchedule(v AnyOfmicrosoftGraphSchedule)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *MicrosoftGraphTeam) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### SetScheduleNil

`func (o *MicrosoftGraphTeam) SetScheduleNil(b bool)`

 SetScheduleNil sets the value for Schedule to be an explicit nil

### UnsetSchedule
`func (o *MicrosoftGraphTeam) UnsetSchedule()`

UnsetSchedule ensures that no value is present for Schedule, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


