# TeamsAppDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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

### NewTeamsAppDefinition

`func NewTeamsAppDefinition() *TeamsAppDefinition`

NewTeamsAppDefinition instantiates a new TeamsAppDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamsAppDefinitionWithDefaults

`func NewTeamsAppDefinitionWithDefaults() *TeamsAppDefinition`

NewTeamsAppDefinitionWithDefaults instantiates a new TeamsAppDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedBy

`func (o *TeamsAppDefinition) GetCreatedBy() AnyOfmicrosoftGraphIdentitySet`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *TeamsAppDefinition) GetCreatedByOk() (*AnyOfmicrosoftGraphIdentitySet, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *TeamsAppDefinition) SetCreatedBy(v AnyOfmicrosoftGraphIdentitySet)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *TeamsAppDefinition) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedByNil

`func (o *TeamsAppDefinition) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *TeamsAppDefinition) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
### GetDescription

`func (o *TeamsAppDefinition) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TeamsAppDefinition) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TeamsAppDefinition) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TeamsAppDefinition) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *TeamsAppDefinition) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *TeamsAppDefinition) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDisplayName

`func (o *TeamsAppDefinition) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *TeamsAppDefinition) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *TeamsAppDefinition) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *TeamsAppDefinition) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *TeamsAppDefinition) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *TeamsAppDefinition) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetLastModifiedDateTime

`func (o *TeamsAppDefinition) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *TeamsAppDefinition) GetLastModifiedDateTimeOk() (*time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedDateTime

`func (o *TeamsAppDefinition) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime sets LastModifiedDateTime field to given value.

### HasLastModifiedDateTime

`func (o *TeamsAppDefinition) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### SetLastModifiedDateTimeNil

`func (o *TeamsAppDefinition) SetLastModifiedDateTimeNil(b bool)`

 SetLastModifiedDateTimeNil sets the value for LastModifiedDateTime to be an explicit nil

### UnsetLastModifiedDateTime
`func (o *TeamsAppDefinition) UnsetLastModifiedDateTime()`

UnsetLastModifiedDateTime ensures that no value is present for LastModifiedDateTime, not even an explicit nil
### GetPublishingState

`func (o *TeamsAppDefinition) GetPublishingState() AnyOfmicrosoftGraphTeamsAppPublishingState`

GetPublishingState returns the PublishingState field if non-nil, zero value otherwise.

### GetPublishingStateOk

`func (o *TeamsAppDefinition) GetPublishingStateOk() (*AnyOfmicrosoftGraphTeamsAppPublishingState, bool)`

GetPublishingStateOk returns a tuple with the PublishingState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishingState

`func (o *TeamsAppDefinition) SetPublishingState(v AnyOfmicrosoftGraphTeamsAppPublishingState)`

SetPublishingState sets PublishingState field to given value.

### HasPublishingState

`func (o *TeamsAppDefinition) HasPublishingState() bool`

HasPublishingState returns a boolean if a field has been set.

### SetPublishingStateNil

`func (o *TeamsAppDefinition) SetPublishingStateNil(b bool)`

 SetPublishingStateNil sets the value for PublishingState to be an explicit nil

### UnsetPublishingState
`func (o *TeamsAppDefinition) UnsetPublishingState()`

UnsetPublishingState ensures that no value is present for PublishingState, not even an explicit nil
### GetShortDescription

`func (o *TeamsAppDefinition) GetShortDescription() string`

GetShortDescription returns the ShortDescription field if non-nil, zero value otherwise.

### GetShortDescriptionOk

`func (o *TeamsAppDefinition) GetShortDescriptionOk() (*string, bool)`

GetShortDescriptionOk returns a tuple with the ShortDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortDescription

`func (o *TeamsAppDefinition) SetShortDescription(v string)`

SetShortDescription sets ShortDescription field to given value.

### HasShortDescription

`func (o *TeamsAppDefinition) HasShortDescription() bool`

HasShortDescription returns a boolean if a field has been set.

### SetShortDescriptionNil

`func (o *TeamsAppDefinition) SetShortDescriptionNil(b bool)`

 SetShortDescriptionNil sets the value for ShortDescription to be an explicit nil

### UnsetShortDescription
`func (o *TeamsAppDefinition) UnsetShortDescription()`

UnsetShortDescription ensures that no value is present for ShortDescription, not even an explicit nil
### GetTeamsAppId

`func (o *TeamsAppDefinition) GetTeamsAppId() string`

GetTeamsAppId returns the TeamsAppId field if non-nil, zero value otherwise.

### GetTeamsAppIdOk

`func (o *TeamsAppDefinition) GetTeamsAppIdOk() (*string, bool)`

GetTeamsAppIdOk returns a tuple with the TeamsAppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamsAppId

`func (o *TeamsAppDefinition) SetTeamsAppId(v string)`

SetTeamsAppId sets TeamsAppId field to given value.

### HasTeamsAppId

`func (o *TeamsAppDefinition) HasTeamsAppId() bool`

HasTeamsAppId returns a boolean if a field has been set.

### SetTeamsAppIdNil

`func (o *TeamsAppDefinition) SetTeamsAppIdNil(b bool)`

 SetTeamsAppIdNil sets the value for TeamsAppId to be an explicit nil

### UnsetTeamsAppId
`func (o *TeamsAppDefinition) UnsetTeamsAppId()`

UnsetTeamsAppId ensures that no value is present for TeamsAppId, not even an explicit nil
### GetVersion

`func (o *TeamsAppDefinition) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *TeamsAppDefinition) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *TeamsAppDefinition) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *TeamsAppDefinition) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *TeamsAppDefinition) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *TeamsAppDefinition) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetBot

`func (o *TeamsAppDefinition) GetBot() AnyOfmicrosoftGraphTeamworkBot`

GetBot returns the Bot field if non-nil, zero value otherwise.

### GetBotOk

`func (o *TeamsAppDefinition) GetBotOk() (*AnyOfmicrosoftGraphTeamworkBot, bool)`

GetBotOk returns a tuple with the Bot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBot

`func (o *TeamsAppDefinition) SetBot(v AnyOfmicrosoftGraphTeamworkBot)`

SetBot sets Bot field to given value.

### HasBot

`func (o *TeamsAppDefinition) HasBot() bool`

HasBot returns a boolean if a field has been set.

### SetBotNil

`func (o *TeamsAppDefinition) SetBotNil(b bool)`

 SetBotNil sets the value for Bot to be an explicit nil

### UnsetBot
`func (o *TeamsAppDefinition) UnsetBot()`

UnsetBot ensures that no value is present for Bot, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


