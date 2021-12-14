# MicrosoftGraphConditionalAccessConditionSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Applications** | Pointer to [**AnyOfmicrosoftGraphConditionalAccessApplications**](anyOf&lt;microsoft.graph.conditionalAccessApplications&gt;.md) | Applications and user actions included in and excluded from the policy. Required. | [optional] 
**ClientAppTypes** | Pointer to [**[]AnyOfmicrosoftGraphConditionalAccessClientApp**](AnyOfmicrosoftGraphConditionalAccessClientApp.md) | Client application types included in the policy. Possible values are: all, browser, mobileAppsAndDesktopClients, exchangeActiveSync, easSupported, other. Required. | [optional] 
**Devices** | Pointer to [**AnyOfmicrosoftGraphConditionalAccessDevices**](anyOf&lt;microsoft.graph.conditionalAccessDevices&gt;.md) | Devices in the policy. | [optional] 
**Locations** | Pointer to [**AnyOfmicrosoftGraphConditionalAccessLocations**](anyOf&lt;microsoft.graph.conditionalAccessLocations&gt;.md) | Locations included in and excluded from the policy. | [optional] 
**Platforms** | Pointer to [**AnyOfmicrosoftGraphConditionalAccessPlatforms**](anyOf&lt;microsoft.graph.conditionalAccessPlatforms&gt;.md) | Platforms included in and excluded from the policy. | [optional] 
**SignInRiskLevels** | Pointer to [**[]AnyOfmicrosoftGraphRiskLevel**](AnyOfmicrosoftGraphRiskLevel.md) | Sign-in risk levels included in the policy. Possible values are: low, medium, high, hidden, none, unknownFutureValue. Required. | [optional] 
**UserRiskLevels** | Pointer to [**[]AnyOfmicrosoftGraphRiskLevel**](AnyOfmicrosoftGraphRiskLevel.md) | User risk levels included in the policy. Possible values are: low, medium, high, hidden, none, unknownFutureValue. Required. | [optional] 
**Users** | Pointer to [**MicrosoftGraphConditionalAccessUsers**](MicrosoftGraphConditionalAccessUsers.md) |  | [optional] 

## Methods

### NewMicrosoftGraphConditionalAccessConditionSet

`func NewMicrosoftGraphConditionalAccessConditionSet() *MicrosoftGraphConditionalAccessConditionSet`

NewMicrosoftGraphConditionalAccessConditionSet instantiates a new MicrosoftGraphConditionalAccessConditionSet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphConditionalAccessConditionSetWithDefaults

`func NewMicrosoftGraphConditionalAccessConditionSetWithDefaults() *MicrosoftGraphConditionalAccessConditionSet`

NewMicrosoftGraphConditionalAccessConditionSetWithDefaults instantiates a new MicrosoftGraphConditionalAccessConditionSet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplications

`func (o *MicrosoftGraphConditionalAccessConditionSet) GetApplications() AnyOfmicrosoftGraphConditionalAccessApplications`

GetApplications returns the Applications field if non-nil, zero value otherwise.

### GetApplicationsOk

`func (o *MicrosoftGraphConditionalAccessConditionSet) GetApplicationsOk() (*AnyOfmicrosoftGraphConditionalAccessApplications, bool)`

GetApplicationsOk returns a tuple with the Applications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplications

`func (o *MicrosoftGraphConditionalAccessConditionSet) SetApplications(v AnyOfmicrosoftGraphConditionalAccessApplications)`

SetApplications sets Applications field to given value.

### HasApplications

`func (o *MicrosoftGraphConditionalAccessConditionSet) HasApplications() bool`

HasApplications returns a boolean if a field has been set.

### SetApplicationsNil

`func (o *MicrosoftGraphConditionalAccessConditionSet) SetApplicationsNil(b bool)`

 SetApplicationsNil sets the value for Applications to be an explicit nil

### UnsetApplications
`func (o *MicrosoftGraphConditionalAccessConditionSet) UnsetApplications()`

UnsetApplications ensures that no value is present for Applications, not even an explicit nil
### GetClientAppTypes

`func (o *MicrosoftGraphConditionalAccessConditionSet) GetClientAppTypes() []AnyOfmicrosoftGraphConditionalAccessClientApp`

GetClientAppTypes returns the ClientAppTypes field if non-nil, zero value otherwise.

### GetClientAppTypesOk

`func (o *MicrosoftGraphConditionalAccessConditionSet) GetClientAppTypesOk() (*[]AnyOfmicrosoftGraphConditionalAccessClientApp, bool)`

GetClientAppTypesOk returns a tuple with the ClientAppTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientAppTypes

`func (o *MicrosoftGraphConditionalAccessConditionSet) SetClientAppTypes(v []AnyOfmicrosoftGraphConditionalAccessClientApp)`

SetClientAppTypes sets ClientAppTypes field to given value.

### HasClientAppTypes

`func (o *MicrosoftGraphConditionalAccessConditionSet) HasClientAppTypes() bool`

HasClientAppTypes returns a boolean if a field has been set.

### GetDevices

`func (o *MicrosoftGraphConditionalAccessConditionSet) GetDevices() AnyOfmicrosoftGraphConditionalAccessDevices`

GetDevices returns the Devices field if non-nil, zero value otherwise.

### GetDevicesOk

`func (o *MicrosoftGraphConditionalAccessConditionSet) GetDevicesOk() (*AnyOfmicrosoftGraphConditionalAccessDevices, bool)`

GetDevicesOk returns a tuple with the Devices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevices

`func (o *MicrosoftGraphConditionalAccessConditionSet) SetDevices(v AnyOfmicrosoftGraphConditionalAccessDevices)`

SetDevices sets Devices field to given value.

### HasDevices

`func (o *MicrosoftGraphConditionalAccessConditionSet) HasDevices() bool`

HasDevices returns a boolean if a field has been set.

### SetDevicesNil

`func (o *MicrosoftGraphConditionalAccessConditionSet) SetDevicesNil(b bool)`

 SetDevicesNil sets the value for Devices to be an explicit nil

### UnsetDevices
`func (o *MicrosoftGraphConditionalAccessConditionSet) UnsetDevices()`

UnsetDevices ensures that no value is present for Devices, not even an explicit nil
### GetLocations

`func (o *MicrosoftGraphConditionalAccessConditionSet) GetLocations() AnyOfmicrosoftGraphConditionalAccessLocations`

GetLocations returns the Locations field if non-nil, zero value otherwise.

### GetLocationsOk

`func (o *MicrosoftGraphConditionalAccessConditionSet) GetLocationsOk() (*AnyOfmicrosoftGraphConditionalAccessLocations, bool)`

GetLocationsOk returns a tuple with the Locations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocations

`func (o *MicrosoftGraphConditionalAccessConditionSet) SetLocations(v AnyOfmicrosoftGraphConditionalAccessLocations)`

SetLocations sets Locations field to given value.

### HasLocations

`func (o *MicrosoftGraphConditionalAccessConditionSet) HasLocations() bool`

HasLocations returns a boolean if a field has been set.

### SetLocationsNil

`func (o *MicrosoftGraphConditionalAccessConditionSet) SetLocationsNil(b bool)`

 SetLocationsNil sets the value for Locations to be an explicit nil

### UnsetLocations
`func (o *MicrosoftGraphConditionalAccessConditionSet) UnsetLocations()`

UnsetLocations ensures that no value is present for Locations, not even an explicit nil
### GetPlatforms

`func (o *MicrosoftGraphConditionalAccessConditionSet) GetPlatforms() AnyOfmicrosoftGraphConditionalAccessPlatforms`

GetPlatforms returns the Platforms field if non-nil, zero value otherwise.

### GetPlatformsOk

`func (o *MicrosoftGraphConditionalAccessConditionSet) GetPlatformsOk() (*AnyOfmicrosoftGraphConditionalAccessPlatforms, bool)`

GetPlatformsOk returns a tuple with the Platforms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatforms

`func (o *MicrosoftGraphConditionalAccessConditionSet) SetPlatforms(v AnyOfmicrosoftGraphConditionalAccessPlatforms)`

SetPlatforms sets Platforms field to given value.

### HasPlatforms

`func (o *MicrosoftGraphConditionalAccessConditionSet) HasPlatforms() bool`

HasPlatforms returns a boolean if a field has been set.

### SetPlatformsNil

`func (o *MicrosoftGraphConditionalAccessConditionSet) SetPlatformsNil(b bool)`

 SetPlatformsNil sets the value for Platforms to be an explicit nil

### UnsetPlatforms
`func (o *MicrosoftGraphConditionalAccessConditionSet) UnsetPlatforms()`

UnsetPlatforms ensures that no value is present for Platforms, not even an explicit nil
### GetSignInRiskLevels

`func (o *MicrosoftGraphConditionalAccessConditionSet) GetSignInRiskLevels() []AnyOfmicrosoftGraphRiskLevel`

GetSignInRiskLevels returns the SignInRiskLevels field if non-nil, zero value otherwise.

### GetSignInRiskLevelsOk

`func (o *MicrosoftGraphConditionalAccessConditionSet) GetSignInRiskLevelsOk() (*[]AnyOfmicrosoftGraphRiskLevel, bool)`

GetSignInRiskLevelsOk returns a tuple with the SignInRiskLevels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignInRiskLevels

`func (o *MicrosoftGraphConditionalAccessConditionSet) SetSignInRiskLevels(v []AnyOfmicrosoftGraphRiskLevel)`

SetSignInRiskLevels sets SignInRiskLevels field to given value.

### HasSignInRiskLevels

`func (o *MicrosoftGraphConditionalAccessConditionSet) HasSignInRiskLevels() bool`

HasSignInRiskLevels returns a boolean if a field has been set.

### GetUserRiskLevels

`func (o *MicrosoftGraphConditionalAccessConditionSet) GetUserRiskLevels() []AnyOfmicrosoftGraphRiskLevel`

GetUserRiskLevels returns the UserRiskLevels field if non-nil, zero value otherwise.

### GetUserRiskLevelsOk

`func (o *MicrosoftGraphConditionalAccessConditionSet) GetUserRiskLevelsOk() (*[]AnyOfmicrosoftGraphRiskLevel, bool)`

GetUserRiskLevelsOk returns a tuple with the UserRiskLevels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserRiskLevels

`func (o *MicrosoftGraphConditionalAccessConditionSet) SetUserRiskLevels(v []AnyOfmicrosoftGraphRiskLevel)`

SetUserRiskLevels sets UserRiskLevels field to given value.

### HasUserRiskLevels

`func (o *MicrosoftGraphConditionalAccessConditionSet) HasUserRiskLevels() bool`

HasUserRiskLevels returns a boolean if a field has been set.

### GetUsers

`func (o *MicrosoftGraphConditionalAccessConditionSet) GetUsers() MicrosoftGraphConditionalAccessUsers`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *MicrosoftGraphConditionalAccessConditionSet) GetUsersOk() (*MicrosoftGraphConditionalAccessUsers, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *MicrosoftGraphConditionalAccessConditionSet) SetUsers(v MicrosoftGraphConditionalAccessUsers)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *MicrosoftGraphConditionalAccessConditionSet) HasUsers() bool`

HasUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


