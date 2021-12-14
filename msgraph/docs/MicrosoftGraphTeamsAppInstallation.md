# MicrosoftGraphTeamsAppInstallation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**TeamsApp** | Pointer to [**AnyOfmicrosoftGraphTeamsApp**](anyOf&lt;microsoft.graph.teamsApp&gt;.md) | The app that is installed. | [optional] 
**TeamsAppDefinition** | Pointer to [**AnyOfmicrosoftGraphTeamsAppDefinition**](anyOf&lt;microsoft.graph.teamsAppDefinition&gt;.md) | The details of this version of the app. | [optional] 

## Methods

### NewMicrosoftGraphTeamsAppInstallation

`func NewMicrosoftGraphTeamsAppInstallation() *MicrosoftGraphTeamsAppInstallation`

NewMicrosoftGraphTeamsAppInstallation instantiates a new MicrosoftGraphTeamsAppInstallation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphTeamsAppInstallationWithDefaults

`func NewMicrosoftGraphTeamsAppInstallationWithDefaults() *MicrosoftGraphTeamsAppInstallation`

NewMicrosoftGraphTeamsAppInstallationWithDefaults instantiates a new MicrosoftGraphTeamsAppInstallation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphTeamsAppInstallation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphTeamsAppInstallation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphTeamsAppInstallation) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphTeamsAppInstallation) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTeamsApp

`func (o *MicrosoftGraphTeamsAppInstallation) GetTeamsApp() AnyOfmicrosoftGraphTeamsApp`

GetTeamsApp returns the TeamsApp field if non-nil, zero value otherwise.

### GetTeamsAppOk

`func (o *MicrosoftGraphTeamsAppInstallation) GetTeamsAppOk() (*AnyOfmicrosoftGraphTeamsApp, bool)`

GetTeamsAppOk returns a tuple with the TeamsApp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamsApp

`func (o *MicrosoftGraphTeamsAppInstallation) SetTeamsApp(v AnyOfmicrosoftGraphTeamsApp)`

SetTeamsApp sets TeamsApp field to given value.

### HasTeamsApp

`func (o *MicrosoftGraphTeamsAppInstallation) HasTeamsApp() bool`

HasTeamsApp returns a boolean if a field has been set.

### SetTeamsAppNil

`func (o *MicrosoftGraphTeamsAppInstallation) SetTeamsAppNil(b bool)`

 SetTeamsAppNil sets the value for TeamsApp to be an explicit nil

### UnsetTeamsApp
`func (o *MicrosoftGraphTeamsAppInstallation) UnsetTeamsApp()`

UnsetTeamsApp ensures that no value is present for TeamsApp, not even an explicit nil
### GetTeamsAppDefinition

`func (o *MicrosoftGraphTeamsAppInstallation) GetTeamsAppDefinition() AnyOfmicrosoftGraphTeamsAppDefinition`

GetTeamsAppDefinition returns the TeamsAppDefinition field if non-nil, zero value otherwise.

### GetTeamsAppDefinitionOk

`func (o *MicrosoftGraphTeamsAppInstallation) GetTeamsAppDefinitionOk() (*AnyOfmicrosoftGraphTeamsAppDefinition, bool)`

GetTeamsAppDefinitionOk returns a tuple with the TeamsAppDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamsAppDefinition

`func (o *MicrosoftGraphTeamsAppInstallation) SetTeamsAppDefinition(v AnyOfmicrosoftGraphTeamsAppDefinition)`

SetTeamsAppDefinition sets TeamsAppDefinition field to given value.

### HasTeamsAppDefinition

`func (o *MicrosoftGraphTeamsAppInstallation) HasTeamsAppDefinition() bool`

HasTeamsAppDefinition returns a boolean if a field has been set.

### SetTeamsAppDefinitionNil

`func (o *MicrosoftGraphTeamsAppInstallation) SetTeamsAppDefinitionNil(b bool)`

 SetTeamsAppDefinitionNil sets the value for TeamsAppDefinition to be an explicit nil

### UnsetTeamsAppDefinition
`func (o *MicrosoftGraphTeamsAppInstallation) UnsetTeamsAppDefinition()`

UnsetTeamsAppDefinition ensures that no value is present for TeamsAppDefinition, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


