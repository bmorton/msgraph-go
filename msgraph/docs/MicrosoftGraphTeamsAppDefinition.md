# MicrosoftGraphTeamsAppDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**CreatedBy** | Pointer to [**AnyOfmicrosoftGraphIdentitySet**](anyOf&lt;microsoft.graph.identitySet&gt;.md) |  | [optional] 
**Description** | Pointer to **NullableString** | Verbose description of the application. | [optional] 
**DisplayName** | Pointer to **NullableString** | The name of the app provided by the app developer. | [optional] 
**LastModifiedDateTime** | Pointer to **NullableTime** |  | [optional] 
**PublishingState** | Pointer to [**AnyOfmicrosoftGraphTeamsAppPublishingState**](anyOf&lt;microsoft.graph.teamsAppPublishingState&gt;.md) | The published status of a specific version of a Teams app. Possible values are:submitted — The specific version of the Teams app has been submitted and is under review. published  — The request to publish the specific version of the Teams app has been approved by the admin and the app is published.  rejected — The request to publish the specific version of the Teams app was rejected by the admin. | [optional] 
**ShortDescription** | Pointer to **NullableString** | Short description of the application. | [optional] 
**TeamsAppId** | Pointer to **NullableString** | The ID from the Teams app manifest. | [optional] 
**Version** | Pointer to **NullableString** | The version number of the application. | [optional] 
**Bot** | Pointer to [**AnyOfmicrosoftGraphTeamworkBot**](anyOf&lt;microsoft.graph.teamworkBot&gt;.md) | The details of the bot specified in the Teams app manifest. | [optional] 

## Methods

### NewMicrosoftGraphTeamsAppDefinition

`func NewMicrosoftGraphTeamsAppDefinition() *MicrosoftGraphTeamsAppDefinition`

NewMicrosoftGraphTeamsAppDefinition instantiates a new MicrosoftGraphTeamsAppDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphTeamsAppDefinitionWithDefaults

`func NewMicrosoftGraphTeamsAppDefinitionWithDefaults() *MicrosoftGraphTeamsAppDefinition`

NewMicrosoftGraphTeamsAppDefinitionWithDefaults instantiates a new MicrosoftGraphTeamsAppDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphTeamsAppDefinition) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphTeamsAppDefinition) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphTeamsAppDefinition) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphTeamsAppDefinition) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedBy

`func (o *MicrosoftGraphTeamsAppDefinition) GetCreatedBy() AnyOfmicrosoftGraphIdentitySet`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *MicrosoftGraphTeamsAppDefinition) GetCreatedByOk() (*AnyOfmicrosoftGraphIdentitySet, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *MicrosoftGraphTeamsAppDefinition) SetCreatedBy(v AnyOfmicrosoftGraphIdentitySet)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *MicrosoftGraphTeamsAppDefinition) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedByNil

`func (o *MicrosoftGraphTeamsAppDefinition) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *MicrosoftGraphTeamsAppDefinition) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
### GetDescription

`func (o *MicrosoftGraphTeamsAppDefinition) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MicrosoftGraphTeamsAppDefinition) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MicrosoftGraphTeamsAppDefinition) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MicrosoftGraphTeamsAppDefinition) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *MicrosoftGraphTeamsAppDefinition) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *MicrosoftGraphTeamsAppDefinition) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDisplayName

`func (o *MicrosoftGraphTeamsAppDefinition) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphTeamsAppDefinition) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *MicrosoftGraphTeamsAppDefinition) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *MicrosoftGraphTeamsAppDefinition) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *MicrosoftGraphTeamsAppDefinition) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *MicrosoftGraphTeamsAppDefinition) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetLastModifiedDateTime

`func (o *MicrosoftGraphTeamsAppDefinition) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *MicrosoftGraphTeamsAppDefinition) GetLastModifiedDateTimeOk() (*time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedDateTime

`func (o *MicrosoftGraphTeamsAppDefinition) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime sets LastModifiedDateTime field to given value.

### HasLastModifiedDateTime

`func (o *MicrosoftGraphTeamsAppDefinition) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### SetLastModifiedDateTimeNil

`func (o *MicrosoftGraphTeamsAppDefinition) SetLastModifiedDateTimeNil(b bool)`

 SetLastModifiedDateTimeNil sets the value for LastModifiedDateTime to be an explicit nil

### UnsetLastModifiedDateTime
`func (o *MicrosoftGraphTeamsAppDefinition) UnsetLastModifiedDateTime()`

UnsetLastModifiedDateTime ensures that no value is present for LastModifiedDateTime, not even an explicit nil
### GetPublishingState

`func (o *MicrosoftGraphTeamsAppDefinition) GetPublishingState() AnyOfmicrosoftGraphTeamsAppPublishingState`

GetPublishingState returns the PublishingState field if non-nil, zero value otherwise.

### GetPublishingStateOk

`func (o *MicrosoftGraphTeamsAppDefinition) GetPublishingStateOk() (*AnyOfmicrosoftGraphTeamsAppPublishingState, bool)`

GetPublishingStateOk returns a tuple with the PublishingState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishingState

`func (o *MicrosoftGraphTeamsAppDefinition) SetPublishingState(v AnyOfmicrosoftGraphTeamsAppPublishingState)`

SetPublishingState sets PublishingState field to given value.

### HasPublishingState

`func (o *MicrosoftGraphTeamsAppDefinition) HasPublishingState() bool`

HasPublishingState returns a boolean if a field has been set.

### SetPublishingStateNil

`func (o *MicrosoftGraphTeamsAppDefinition) SetPublishingStateNil(b bool)`

 SetPublishingStateNil sets the value for PublishingState to be an explicit nil

### UnsetPublishingState
`func (o *MicrosoftGraphTeamsAppDefinition) UnsetPublishingState()`

UnsetPublishingState ensures that no value is present for PublishingState, not even an explicit nil
### GetShortDescription

`func (o *MicrosoftGraphTeamsAppDefinition) GetShortDescription() string`

GetShortDescription returns the ShortDescription field if non-nil, zero value otherwise.

### GetShortDescriptionOk

`func (o *MicrosoftGraphTeamsAppDefinition) GetShortDescriptionOk() (*string, bool)`

GetShortDescriptionOk returns a tuple with the ShortDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortDescription

`func (o *MicrosoftGraphTeamsAppDefinition) SetShortDescription(v string)`

SetShortDescription sets ShortDescription field to given value.

### HasShortDescription

`func (o *MicrosoftGraphTeamsAppDefinition) HasShortDescription() bool`

HasShortDescription returns a boolean if a field has been set.

### SetShortDescriptionNil

`func (o *MicrosoftGraphTeamsAppDefinition) SetShortDescriptionNil(b bool)`

 SetShortDescriptionNil sets the value for ShortDescription to be an explicit nil

### UnsetShortDescription
`func (o *MicrosoftGraphTeamsAppDefinition) UnsetShortDescription()`

UnsetShortDescription ensures that no value is present for ShortDescription, not even an explicit nil
### GetTeamsAppId

`func (o *MicrosoftGraphTeamsAppDefinition) GetTeamsAppId() string`

GetTeamsAppId returns the TeamsAppId field if non-nil, zero value otherwise.

### GetTeamsAppIdOk

`func (o *MicrosoftGraphTeamsAppDefinition) GetTeamsAppIdOk() (*string, bool)`

GetTeamsAppIdOk returns a tuple with the TeamsAppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamsAppId

`func (o *MicrosoftGraphTeamsAppDefinition) SetTeamsAppId(v string)`

SetTeamsAppId sets TeamsAppId field to given value.

### HasTeamsAppId

`func (o *MicrosoftGraphTeamsAppDefinition) HasTeamsAppId() bool`

HasTeamsAppId returns a boolean if a field has been set.

### SetTeamsAppIdNil

`func (o *MicrosoftGraphTeamsAppDefinition) SetTeamsAppIdNil(b bool)`

 SetTeamsAppIdNil sets the value for TeamsAppId to be an explicit nil

### UnsetTeamsAppId
`func (o *MicrosoftGraphTeamsAppDefinition) UnsetTeamsAppId()`

UnsetTeamsAppId ensures that no value is present for TeamsAppId, not even an explicit nil
### GetVersion

`func (o *MicrosoftGraphTeamsAppDefinition) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *MicrosoftGraphTeamsAppDefinition) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *MicrosoftGraphTeamsAppDefinition) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *MicrosoftGraphTeamsAppDefinition) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *MicrosoftGraphTeamsAppDefinition) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *MicrosoftGraphTeamsAppDefinition) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetBot

`func (o *MicrosoftGraphTeamsAppDefinition) GetBot() AnyOfmicrosoftGraphTeamworkBot`

GetBot returns the Bot field if non-nil, zero value otherwise.

### GetBotOk

`func (o *MicrosoftGraphTeamsAppDefinition) GetBotOk() (*AnyOfmicrosoftGraphTeamworkBot, bool)`

GetBotOk returns a tuple with the Bot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBot

`func (o *MicrosoftGraphTeamsAppDefinition) SetBot(v AnyOfmicrosoftGraphTeamworkBot)`

SetBot sets Bot field to given value.

### HasBot

`func (o *MicrosoftGraphTeamsAppDefinition) HasBot() bool`

HasBot returns a boolean if a field has been set.

### SetBotNil

`func (o *MicrosoftGraphTeamsAppDefinition) SetBotNil(b bool)`

 SetBotNil sets the value for Bot to be an explicit nil

### UnsetBot
`func (o *MicrosoftGraphTeamsAppDefinition) UnsetBot()`

UnsetBot ensures that no value is present for Bot, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


