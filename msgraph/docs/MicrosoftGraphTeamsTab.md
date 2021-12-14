# MicrosoftGraphTeamsTab

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**Configuration** | Pointer to [**AnyOfmicrosoftGraphTeamsTabConfiguration**](anyOf&lt;microsoft.graph.teamsTabConfiguration&gt;.md) | Container for custom settings applied to a tab. The tab is considered configured only once this property is set. | [optional] 
**DisplayName** | Pointer to **NullableString** | Name of the tab. | [optional] 
**WebUrl** | Pointer to **NullableString** | Deep link URL of the tab instance. Read only. | [optional] 
**TeamsApp** | Pointer to [**AnyOfmicrosoftGraphTeamsApp**](anyOf&lt;microsoft.graph.teamsApp&gt;.md) | The application that is linked to the tab. This cannot be changed after tab creation. | [optional] 

## Methods

### NewMicrosoftGraphTeamsTab

`func NewMicrosoftGraphTeamsTab() *MicrosoftGraphTeamsTab`

NewMicrosoftGraphTeamsTab instantiates a new MicrosoftGraphTeamsTab object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphTeamsTabWithDefaults

`func NewMicrosoftGraphTeamsTabWithDefaults() *MicrosoftGraphTeamsTab`

NewMicrosoftGraphTeamsTabWithDefaults instantiates a new MicrosoftGraphTeamsTab object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphTeamsTab) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphTeamsTab) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphTeamsTab) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphTeamsTab) HasId() bool`

HasId returns a boolean if a field has been set.

### GetConfiguration

`func (o *MicrosoftGraphTeamsTab) GetConfiguration() AnyOfmicrosoftGraphTeamsTabConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *MicrosoftGraphTeamsTab) GetConfigurationOk() (*AnyOfmicrosoftGraphTeamsTabConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *MicrosoftGraphTeamsTab) SetConfiguration(v AnyOfmicrosoftGraphTeamsTabConfiguration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *MicrosoftGraphTeamsTab) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### SetConfigurationNil

`func (o *MicrosoftGraphTeamsTab) SetConfigurationNil(b bool)`

 SetConfigurationNil sets the value for Configuration to be an explicit nil

### UnsetConfiguration
`func (o *MicrosoftGraphTeamsTab) UnsetConfiguration()`

UnsetConfiguration ensures that no value is present for Configuration, not even an explicit nil
### GetDisplayName

`func (o *MicrosoftGraphTeamsTab) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphTeamsTab) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *MicrosoftGraphTeamsTab) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *MicrosoftGraphTeamsTab) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *MicrosoftGraphTeamsTab) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *MicrosoftGraphTeamsTab) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetWebUrl

`func (o *MicrosoftGraphTeamsTab) GetWebUrl() string`

GetWebUrl returns the WebUrl field if non-nil, zero value otherwise.

### GetWebUrlOk

`func (o *MicrosoftGraphTeamsTab) GetWebUrlOk() (*string, bool)`

GetWebUrlOk returns a tuple with the WebUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebUrl

`func (o *MicrosoftGraphTeamsTab) SetWebUrl(v string)`

SetWebUrl sets WebUrl field to given value.

### HasWebUrl

`func (o *MicrosoftGraphTeamsTab) HasWebUrl() bool`

HasWebUrl returns a boolean if a field has been set.

### SetWebUrlNil

`func (o *MicrosoftGraphTeamsTab) SetWebUrlNil(b bool)`

 SetWebUrlNil sets the value for WebUrl to be an explicit nil

### UnsetWebUrl
`func (o *MicrosoftGraphTeamsTab) UnsetWebUrl()`

UnsetWebUrl ensures that no value is present for WebUrl, not even an explicit nil
### GetTeamsApp

`func (o *MicrosoftGraphTeamsTab) GetTeamsApp() AnyOfmicrosoftGraphTeamsApp`

GetTeamsApp returns the TeamsApp field if non-nil, zero value otherwise.

### GetTeamsAppOk

`func (o *MicrosoftGraphTeamsTab) GetTeamsAppOk() (*AnyOfmicrosoftGraphTeamsApp, bool)`

GetTeamsAppOk returns a tuple with the TeamsApp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamsApp

`func (o *MicrosoftGraphTeamsTab) SetTeamsApp(v AnyOfmicrosoftGraphTeamsApp)`

SetTeamsApp sets TeamsApp field to given value.

### HasTeamsApp

`func (o *MicrosoftGraphTeamsTab) HasTeamsApp() bool`

HasTeamsApp returns a boolean if a field has been set.

### SetTeamsAppNil

`func (o *MicrosoftGraphTeamsTab) SetTeamsAppNil(b bool)`

 SetTeamsAppNil sets the value for TeamsApp to be an explicit nil

### UnsetTeamsApp
`func (o *MicrosoftGraphTeamsTab) UnsetTeamsApp()`

UnsetTeamsApp ensures that no value is present for TeamsApp, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


