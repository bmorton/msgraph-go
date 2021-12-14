# MicrosoftGraphLobbyBypassSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsDialInBypassEnabled** | Pointer to **NullableBool** | Specifies whether or not to always let dial-in callers bypass the lobby. Optional. | [optional] 
**Scope** | Pointer to [**AnyOfmicrosoftGraphLobbyBypassScope**](anyOf&lt;microsoft.graph.lobbyBypassScope&gt;.md) | Specifies the type of participants that are automatically admitted into a meeting, bypassing the lobby. Optional. | [optional] 

## Methods

### NewMicrosoftGraphLobbyBypassSettings

`func NewMicrosoftGraphLobbyBypassSettings() *MicrosoftGraphLobbyBypassSettings`

NewMicrosoftGraphLobbyBypassSettings instantiates a new MicrosoftGraphLobbyBypassSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphLobbyBypassSettingsWithDefaults

`func NewMicrosoftGraphLobbyBypassSettingsWithDefaults() *MicrosoftGraphLobbyBypassSettings`

NewMicrosoftGraphLobbyBypassSettingsWithDefaults instantiates a new MicrosoftGraphLobbyBypassSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsDialInBypassEnabled

`func (o *MicrosoftGraphLobbyBypassSettings) GetIsDialInBypassEnabled() bool`

GetIsDialInBypassEnabled returns the IsDialInBypassEnabled field if non-nil, zero value otherwise.

### GetIsDialInBypassEnabledOk

`func (o *MicrosoftGraphLobbyBypassSettings) GetIsDialInBypassEnabledOk() (*bool, bool)`

GetIsDialInBypassEnabledOk returns a tuple with the IsDialInBypassEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDialInBypassEnabled

`func (o *MicrosoftGraphLobbyBypassSettings) SetIsDialInBypassEnabled(v bool)`

SetIsDialInBypassEnabled sets IsDialInBypassEnabled field to given value.

### HasIsDialInBypassEnabled

`func (o *MicrosoftGraphLobbyBypassSettings) HasIsDialInBypassEnabled() bool`

HasIsDialInBypassEnabled returns a boolean if a field has been set.

### SetIsDialInBypassEnabledNil

`func (o *MicrosoftGraphLobbyBypassSettings) SetIsDialInBypassEnabledNil(b bool)`

 SetIsDialInBypassEnabledNil sets the value for IsDialInBypassEnabled to be an explicit nil

### UnsetIsDialInBypassEnabled
`func (o *MicrosoftGraphLobbyBypassSettings) UnsetIsDialInBypassEnabled()`

UnsetIsDialInBypassEnabled ensures that no value is present for IsDialInBypassEnabled, not even an explicit nil
### GetScope

`func (o *MicrosoftGraphLobbyBypassSettings) GetScope() AnyOfmicrosoftGraphLobbyBypassScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *MicrosoftGraphLobbyBypassSettings) GetScopeOk() (*AnyOfmicrosoftGraphLobbyBypassScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *MicrosoftGraphLobbyBypassSettings) SetScope(v AnyOfmicrosoftGraphLobbyBypassScope)`

SetScope sets Scope field to given value.

### HasScope

`func (o *MicrosoftGraphLobbyBypassSettings) HasScope() bool`

HasScope returns a boolean if a field has been set.

### SetScopeNil

`func (o *MicrosoftGraphLobbyBypassSettings) SetScopeNil(b bool)`

 SetScopeNil sets the value for Scope to be an explicit nil

### UnsetScope
`func (o *MicrosoftGraphLobbyBypassSettings) UnsetScope()`

UnsetScope ensures that no value is present for Scope, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


