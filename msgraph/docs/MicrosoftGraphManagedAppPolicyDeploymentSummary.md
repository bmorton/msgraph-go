# MicrosoftGraphManagedAppPolicyDeploymentSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**ConfigurationDeployedUserCount** | Pointer to **int32** | Not yet documented | [optional] 
**ConfigurationDeploymentSummaryPerApp** | Pointer to [**[]AnyOfmicrosoftGraphManagedAppPolicyDeploymentSummaryPerApp**](AnyOfmicrosoftGraphManagedAppPolicyDeploymentSummaryPerApp.md) | Not yet documented | [optional] 
**DisplayName** | Pointer to **NullableString** | Not yet documented | [optional] 
**LastRefreshTime** | Pointer to **time.Time** | Not yet documented | [optional] 
**Version** | Pointer to **NullableString** | Version of the entity. | [optional] 

## Methods

### NewMicrosoftGraphManagedAppPolicyDeploymentSummary

`func NewMicrosoftGraphManagedAppPolicyDeploymentSummary() *MicrosoftGraphManagedAppPolicyDeploymentSummary`

NewMicrosoftGraphManagedAppPolicyDeploymentSummary instantiates a new MicrosoftGraphManagedAppPolicyDeploymentSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphManagedAppPolicyDeploymentSummaryWithDefaults

`func NewMicrosoftGraphManagedAppPolicyDeploymentSummaryWithDefaults() *MicrosoftGraphManagedAppPolicyDeploymentSummary`

NewMicrosoftGraphManagedAppPolicyDeploymentSummaryWithDefaults instantiates a new MicrosoftGraphManagedAppPolicyDeploymentSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphManagedAppPolicyDeploymentSummary) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphManagedAppPolicyDeploymentSummary) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphManagedAppPolicyDeploymentSummary) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphManagedAppPolicyDeploymentSummary) HasId() bool`

HasId returns a boolean if a field has been set.

### GetConfigurationDeployedUserCount

`func (o *MicrosoftGraphManagedAppPolicyDeploymentSummary) GetConfigurationDeployedUserCount() int32`

GetConfigurationDeployedUserCount returns the ConfigurationDeployedUserCount field if non-nil, zero value otherwise.

### GetConfigurationDeployedUserCountOk

`func (o *MicrosoftGraphManagedAppPolicyDeploymentSummary) GetConfigurationDeployedUserCountOk() (*int32, bool)`

GetConfigurationDeployedUserCountOk returns a tuple with the ConfigurationDeployedUserCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationDeployedUserCount

`func (o *MicrosoftGraphManagedAppPolicyDeploymentSummary) SetConfigurationDeployedUserCount(v int32)`

SetConfigurationDeployedUserCount sets ConfigurationDeployedUserCount field to given value.

### HasConfigurationDeployedUserCount

`func (o *MicrosoftGraphManagedAppPolicyDeploymentSummary) HasConfigurationDeployedUserCount() bool`

HasConfigurationDeployedUserCount returns a boolean if a field has been set.

### GetConfigurationDeploymentSummaryPerApp

`func (o *MicrosoftGraphManagedAppPolicyDeploymentSummary) GetConfigurationDeploymentSummaryPerApp() []*AnyOfmicrosoftGraphManagedAppPolicyDeploymentSummaryPerApp`

GetConfigurationDeploymentSummaryPerApp returns the ConfigurationDeploymentSummaryPerApp field if non-nil, zero value otherwise.

### GetConfigurationDeploymentSummaryPerAppOk

`func (o *MicrosoftGraphManagedAppPolicyDeploymentSummary) GetConfigurationDeploymentSummaryPerAppOk() (*[]*AnyOfmicrosoftGraphManagedAppPolicyDeploymentSummaryPerApp, bool)`

GetConfigurationDeploymentSummaryPerAppOk returns a tuple with the ConfigurationDeploymentSummaryPerApp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationDeploymentSummaryPerApp

`func (o *MicrosoftGraphManagedAppPolicyDeploymentSummary) SetConfigurationDeploymentSummaryPerApp(v []*AnyOfmicrosoftGraphManagedAppPolicyDeploymentSummaryPerApp)`

SetConfigurationDeploymentSummaryPerApp sets ConfigurationDeploymentSummaryPerApp field to given value.

### HasConfigurationDeploymentSummaryPerApp

`func (o *MicrosoftGraphManagedAppPolicyDeploymentSummary) HasConfigurationDeploymentSummaryPerApp() bool`

HasConfigurationDeploymentSummaryPerApp returns a boolean if a field has been set.

### GetDisplayName

`func (o *MicrosoftGraphManagedAppPolicyDeploymentSummary) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphManagedAppPolicyDeploymentSummary) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *MicrosoftGraphManagedAppPolicyDeploymentSummary) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *MicrosoftGraphManagedAppPolicyDeploymentSummary) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *MicrosoftGraphManagedAppPolicyDeploymentSummary) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *MicrosoftGraphManagedAppPolicyDeploymentSummary) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetLastRefreshTime

`func (o *MicrosoftGraphManagedAppPolicyDeploymentSummary) GetLastRefreshTime() time.Time`

GetLastRefreshTime returns the LastRefreshTime field if non-nil, zero value otherwise.

### GetLastRefreshTimeOk

`func (o *MicrosoftGraphManagedAppPolicyDeploymentSummary) GetLastRefreshTimeOk() (*time.Time, bool)`

GetLastRefreshTimeOk returns a tuple with the LastRefreshTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRefreshTime

`func (o *MicrosoftGraphManagedAppPolicyDeploymentSummary) SetLastRefreshTime(v time.Time)`

SetLastRefreshTime sets LastRefreshTime field to given value.

### HasLastRefreshTime

`func (o *MicrosoftGraphManagedAppPolicyDeploymentSummary) HasLastRefreshTime() bool`

HasLastRefreshTime returns a boolean if a field has been set.

### GetVersion

`func (o *MicrosoftGraphManagedAppPolicyDeploymentSummary) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *MicrosoftGraphManagedAppPolicyDeploymentSummary) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *MicrosoftGraphManagedAppPolicyDeploymentSummary) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *MicrosoftGraphManagedAppPolicyDeploymentSummary) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *MicrosoftGraphManagedAppPolicyDeploymentSummary) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *MicrosoftGraphManagedAppPolicyDeploymentSummary) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


