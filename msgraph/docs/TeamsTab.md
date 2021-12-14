# TeamsTab

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Configuration** | Pointer to [**AnyOfmicrosoftGraphTeamsTabConfiguration**](anyOf&lt;microsoft.graph.teamsTabConfiguration&gt;.md) | Container for custom settings applied to a tab. The tab is considered configured only once this property is set. | [optional] 
**DisplayName** | Pointer to **NullableString** | Name of the tab. | [optional] 
**WebUrl** | Pointer to **NullableString** | Deep link URL of the tab instance. Read only. | [optional] 
**TeamsApp** | Pointer to [**AnyOfmicrosoftGraphTeamsApp**](anyOf&lt;microsoft.graph.teamsApp&gt;.md) | The application that is linked to the tab. This cannot be changed after tab creation. | [optional] 

## Methods

### NewTeamsTab

`func NewTeamsTab() *TeamsTab`

NewTeamsTab instantiates a new TeamsTab object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamsTabWithDefaults

`func NewTeamsTabWithDefaults() *TeamsTab`

NewTeamsTabWithDefaults instantiates a new TeamsTab object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfiguration

`func (o *TeamsTab) GetConfiguration() AnyOfmicrosoftGraphTeamsTabConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *TeamsTab) GetConfigurationOk() (*AnyOfmicrosoftGraphTeamsTabConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *TeamsTab) SetConfiguration(v AnyOfmicrosoftGraphTeamsTabConfiguration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *TeamsTab) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### SetConfigurationNil

`func (o *TeamsTab) SetConfigurationNil(b bool)`

 SetConfigurationNil sets the value for Configuration to be an explicit nil

### UnsetConfiguration
`func (o *TeamsTab) UnsetConfiguration()`

UnsetConfiguration ensures that no value is present for Configuration, not even an explicit nil
### GetDisplayName

`func (o *TeamsTab) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *TeamsTab) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *TeamsTab) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *TeamsTab) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *TeamsTab) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *TeamsTab) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetWebUrl

`func (o *TeamsTab) GetWebUrl() string`

GetWebUrl returns the WebUrl field if non-nil, zero value otherwise.

### GetWebUrlOk

`func (o *TeamsTab) GetWebUrlOk() (*string, bool)`

GetWebUrlOk returns a tuple with the WebUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebUrl

`func (o *TeamsTab) SetWebUrl(v string)`

SetWebUrl sets WebUrl field to given value.

### HasWebUrl

`func (o *TeamsTab) HasWebUrl() bool`

HasWebUrl returns a boolean if a field has been set.

### SetWebUrlNil

`func (o *TeamsTab) SetWebUrlNil(b bool)`

 SetWebUrlNil sets the value for WebUrl to be an explicit nil

### UnsetWebUrl
`func (o *TeamsTab) UnsetWebUrl()`

UnsetWebUrl ensures that no value is present for WebUrl, not even an explicit nil
### GetTeamsApp

`func (o *TeamsTab) GetTeamsApp() AnyOfmicrosoftGraphTeamsApp`

GetTeamsApp returns the TeamsApp field if non-nil, zero value otherwise.

### GetTeamsAppOk

`func (o *TeamsTab) GetTeamsAppOk() (*AnyOfmicrosoftGraphTeamsApp, bool)`

GetTeamsAppOk returns a tuple with the TeamsApp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamsApp

`func (o *TeamsTab) SetTeamsApp(v AnyOfmicrosoftGraphTeamsApp)`

SetTeamsApp sets TeamsApp field to given value.

### HasTeamsApp

`func (o *TeamsTab) HasTeamsApp() bool`

HasTeamsApp returns a boolean if a field has been set.

### SetTeamsAppNil

`func (o *TeamsTab) SetTeamsAppNil(b bool)`

 SetTeamsAppNil sets the value for TeamsApp to be an explicit nil

### UnsetTeamsApp
`func (o *TeamsTab) UnsetTeamsApp()`

UnsetTeamsApp ensures that no value is present for TeamsApp, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


