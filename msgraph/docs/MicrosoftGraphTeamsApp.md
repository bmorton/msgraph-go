# MicrosoftGraphTeamsApp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**DisplayName** | Pointer to **NullableString** | The name of the catalog app provided by the app developer in the Microsoft Teams zip app package. | [optional] 
**DistributionMethod** | Pointer to [**AnyOfmicrosoftGraphTeamsAppDistributionMethod**](anyOf&lt;microsoft.graph.teamsAppDistributionMethod&gt;.md) | The method of distribution for the app. Read-only. | [optional] 
**ExternalId** | Pointer to **NullableString** | The ID of the catalog provided by the app developer in the Microsoft Teams zip app package. | [optional] 
**AppDefinitions** | Pointer to [**[]MicrosoftGraphTeamsAppDefinition**](MicrosoftGraphTeamsAppDefinition.md) | The details for each version of the app. | [optional] 

## Methods

### NewMicrosoftGraphTeamsApp

`func NewMicrosoftGraphTeamsApp() *MicrosoftGraphTeamsApp`

NewMicrosoftGraphTeamsApp instantiates a new MicrosoftGraphTeamsApp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphTeamsAppWithDefaults

`func NewMicrosoftGraphTeamsAppWithDefaults() *MicrosoftGraphTeamsApp`

NewMicrosoftGraphTeamsAppWithDefaults instantiates a new MicrosoftGraphTeamsApp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphTeamsApp) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphTeamsApp) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphTeamsApp) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphTeamsApp) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDisplayName

`func (o *MicrosoftGraphTeamsApp) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphTeamsApp) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *MicrosoftGraphTeamsApp) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *MicrosoftGraphTeamsApp) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *MicrosoftGraphTeamsApp) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *MicrosoftGraphTeamsApp) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetDistributionMethod

`func (o *MicrosoftGraphTeamsApp) GetDistributionMethod() AnyOfmicrosoftGraphTeamsAppDistributionMethod`

GetDistributionMethod returns the DistributionMethod field if non-nil, zero value otherwise.

### GetDistributionMethodOk

`func (o *MicrosoftGraphTeamsApp) GetDistributionMethodOk() (*AnyOfmicrosoftGraphTeamsAppDistributionMethod, bool)`

GetDistributionMethodOk returns a tuple with the DistributionMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistributionMethod

`func (o *MicrosoftGraphTeamsApp) SetDistributionMethod(v AnyOfmicrosoftGraphTeamsAppDistributionMethod)`

SetDistributionMethod sets DistributionMethod field to given value.

### HasDistributionMethod

`func (o *MicrosoftGraphTeamsApp) HasDistributionMethod() bool`

HasDistributionMethod returns a boolean if a field has been set.

### SetDistributionMethodNil

`func (o *MicrosoftGraphTeamsApp) SetDistributionMethodNil(b bool)`

 SetDistributionMethodNil sets the value for DistributionMethod to be an explicit nil

### UnsetDistributionMethod
`func (o *MicrosoftGraphTeamsApp) UnsetDistributionMethod()`

UnsetDistributionMethod ensures that no value is present for DistributionMethod, not even an explicit nil
### GetExternalId

`func (o *MicrosoftGraphTeamsApp) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *MicrosoftGraphTeamsApp) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *MicrosoftGraphTeamsApp) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *MicrosoftGraphTeamsApp) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *MicrosoftGraphTeamsApp) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *MicrosoftGraphTeamsApp) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetAppDefinitions

`func (o *MicrosoftGraphTeamsApp) GetAppDefinitions() []MicrosoftGraphTeamsAppDefinition`

GetAppDefinitions returns the AppDefinitions field if non-nil, zero value otherwise.

### GetAppDefinitionsOk

`func (o *MicrosoftGraphTeamsApp) GetAppDefinitionsOk() (*[]MicrosoftGraphTeamsAppDefinition, bool)`

GetAppDefinitionsOk returns a tuple with the AppDefinitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppDefinitions

`func (o *MicrosoftGraphTeamsApp) SetAppDefinitions(v []MicrosoftGraphTeamsAppDefinition)`

SetAppDefinitions sets AppDefinitions field to given value.

### HasAppDefinitions

`func (o *MicrosoftGraphTeamsApp) HasAppDefinitions() bool`

HasAppDefinitions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


