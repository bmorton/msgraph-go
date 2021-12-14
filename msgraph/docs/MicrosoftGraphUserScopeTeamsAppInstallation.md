# MicrosoftGraphUserScopeTeamsAppInstallation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**TeamsApp** | Pointer to [**AnyOfmicrosoftGraphTeamsApp**](anyOf&lt;microsoft.graph.teamsApp&gt;.md) | The app that is installed. | [optional] 
**TeamsAppDefinition** | Pointer to [**AnyOfmicrosoftGraphTeamsAppDefinition**](anyOf&lt;microsoft.graph.teamsAppDefinition&gt;.md) | The details of this version of the app. | [optional] 
**Chat** | Pointer to [**AnyOfmicrosoftGraphChat**](anyOf&lt;microsoft.graph.chat&gt;.md) | The chat between the user and Teams app. | [optional] 

## Methods

### NewMicrosoftGraphUserScopeTeamsAppInstallation

`func NewMicrosoftGraphUserScopeTeamsAppInstallation() *MicrosoftGraphUserScopeTeamsAppInstallation`

NewMicrosoftGraphUserScopeTeamsAppInstallation instantiates a new MicrosoftGraphUserScopeTeamsAppInstallation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphUserScopeTeamsAppInstallationWithDefaults

`func NewMicrosoftGraphUserScopeTeamsAppInstallationWithDefaults() *MicrosoftGraphUserScopeTeamsAppInstallation`

NewMicrosoftGraphUserScopeTeamsAppInstallationWithDefaults instantiates a new MicrosoftGraphUserScopeTeamsAppInstallation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphUserScopeTeamsAppInstallation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphUserScopeTeamsAppInstallation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphUserScopeTeamsAppInstallation) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphUserScopeTeamsAppInstallation) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTeamsApp

`func (o *MicrosoftGraphUserScopeTeamsAppInstallation) GetTeamsApp() AnyOfmicrosoftGraphTeamsApp`

GetTeamsApp returns the TeamsApp field if non-nil, zero value otherwise.

### GetTeamsAppOk

`func (o *MicrosoftGraphUserScopeTeamsAppInstallation) GetTeamsAppOk() (*AnyOfmicrosoftGraphTeamsApp, bool)`

GetTeamsAppOk returns a tuple with the TeamsApp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamsApp

`func (o *MicrosoftGraphUserScopeTeamsAppInstallation) SetTeamsApp(v AnyOfmicrosoftGraphTeamsApp)`

SetTeamsApp sets TeamsApp field to given value.

### HasTeamsApp

`func (o *MicrosoftGraphUserScopeTeamsAppInstallation) HasTeamsApp() bool`

HasTeamsApp returns a boolean if a field has been set.

### SetTeamsAppNil

`func (o *MicrosoftGraphUserScopeTeamsAppInstallation) SetTeamsAppNil(b bool)`

 SetTeamsAppNil sets the value for TeamsApp to be an explicit nil

### UnsetTeamsApp
`func (o *MicrosoftGraphUserScopeTeamsAppInstallation) UnsetTeamsApp()`

UnsetTeamsApp ensures that no value is present for TeamsApp, not even an explicit nil
### GetTeamsAppDefinition

`func (o *MicrosoftGraphUserScopeTeamsAppInstallation) GetTeamsAppDefinition() AnyOfmicrosoftGraphTeamsAppDefinition`

GetTeamsAppDefinition returns the TeamsAppDefinition field if non-nil, zero value otherwise.

### GetTeamsAppDefinitionOk

`func (o *MicrosoftGraphUserScopeTeamsAppInstallation) GetTeamsAppDefinitionOk() (*AnyOfmicrosoftGraphTeamsAppDefinition, bool)`

GetTeamsAppDefinitionOk returns a tuple with the TeamsAppDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamsAppDefinition

`func (o *MicrosoftGraphUserScopeTeamsAppInstallation) SetTeamsAppDefinition(v AnyOfmicrosoftGraphTeamsAppDefinition)`

SetTeamsAppDefinition sets TeamsAppDefinition field to given value.

### HasTeamsAppDefinition

`func (o *MicrosoftGraphUserScopeTeamsAppInstallation) HasTeamsAppDefinition() bool`

HasTeamsAppDefinition returns a boolean if a field has been set.

### SetTeamsAppDefinitionNil

`func (o *MicrosoftGraphUserScopeTeamsAppInstallation) SetTeamsAppDefinitionNil(b bool)`

 SetTeamsAppDefinitionNil sets the value for TeamsAppDefinition to be an explicit nil

### UnsetTeamsAppDefinition
`func (o *MicrosoftGraphUserScopeTeamsAppInstallation) UnsetTeamsAppDefinition()`

UnsetTeamsAppDefinition ensures that no value is present for TeamsAppDefinition, not even an explicit nil
### GetChat

`func (o *MicrosoftGraphUserScopeTeamsAppInstallation) GetChat() AnyOfmicrosoftGraphChat`

GetChat returns the Chat field if non-nil, zero value otherwise.

### GetChatOk

`func (o *MicrosoftGraphUserScopeTeamsAppInstallation) GetChatOk() (*AnyOfmicrosoftGraphChat, bool)`

GetChatOk returns a tuple with the Chat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChat

`func (o *MicrosoftGraphUserScopeTeamsAppInstallation) SetChat(v AnyOfmicrosoftGraphChat)`

SetChat sets Chat field to given value.

### HasChat

`func (o *MicrosoftGraphUserScopeTeamsAppInstallation) HasChat() bool`

HasChat returns a boolean if a field has been set.

### SetChatNil

`func (o *MicrosoftGraphUserScopeTeamsAppInstallation) SetChatNil(b bool)`

 SetChatNil sets the value for Chat to be an explicit nil

### UnsetChat
`func (o *MicrosoftGraphUserScopeTeamsAppInstallation) UnsetChat()`

UnsetChat ensures that no value is present for Chat, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


