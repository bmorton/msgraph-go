# MicrosoftGraphUserTeamwork

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**InstalledApps** | Pointer to [**[]MicrosoftGraphUserScopeTeamsAppInstallation**](MicrosoftGraphUserScopeTeamsAppInstallation.md) | The apps installed in the personal scope of this user. | [optional] 

## Methods

### NewMicrosoftGraphUserTeamwork

`func NewMicrosoftGraphUserTeamwork() *MicrosoftGraphUserTeamwork`

NewMicrosoftGraphUserTeamwork instantiates a new MicrosoftGraphUserTeamwork object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphUserTeamworkWithDefaults

`func NewMicrosoftGraphUserTeamworkWithDefaults() *MicrosoftGraphUserTeamwork`

NewMicrosoftGraphUserTeamworkWithDefaults instantiates a new MicrosoftGraphUserTeamwork object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphUserTeamwork) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphUserTeamwork) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphUserTeamwork) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphUserTeamwork) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInstalledApps

`func (o *MicrosoftGraphUserTeamwork) GetInstalledApps() []MicrosoftGraphUserScopeTeamsAppInstallation`

GetInstalledApps returns the InstalledApps field if non-nil, zero value otherwise.

### GetInstalledAppsOk

`func (o *MicrosoftGraphUserTeamwork) GetInstalledAppsOk() (*[]MicrosoftGraphUserScopeTeamsAppInstallation, bool)`

GetInstalledAppsOk returns a tuple with the InstalledApps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstalledApps

`func (o *MicrosoftGraphUserTeamwork) SetInstalledApps(v []MicrosoftGraphUserScopeTeamsAppInstallation)`

SetInstalledApps sets InstalledApps field to given value.

### HasInstalledApps

`func (o *MicrosoftGraphUserTeamwork) HasInstalledApps() bool`

HasInstalledApps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


