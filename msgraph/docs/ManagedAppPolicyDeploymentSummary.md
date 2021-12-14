# ManagedAppPolicyDeploymentSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigurationDeployedUserCount** | Pointer to **int32** | Not yet documented | [optional] 
**ConfigurationDeploymentSummaryPerApp** | Pointer to [**[]AnyOfmicrosoftGraphManagedAppPolicyDeploymentSummaryPerApp**](AnyOfmicrosoftGraphManagedAppPolicyDeploymentSummaryPerApp.md) | Not yet documented | [optional] 
**DisplayName** | Pointer to **NullableString** | Not yet documented | [optional] 
**LastRefreshTime** | Pointer to **time.Time** | Not yet documented | [optional] 
**Version** | Pointer to **NullableString** | Version of the entity. | [optional] 

## Methods

### NewManagedAppPolicyDeploymentSummary

`func NewManagedAppPolicyDeploymentSummary() *ManagedAppPolicyDeploymentSummary`

NewManagedAppPolicyDeploymentSummary instantiates a new ManagedAppPolicyDeploymentSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagedAppPolicyDeploymentSummaryWithDefaults

`func NewManagedAppPolicyDeploymentSummaryWithDefaults() *ManagedAppPolicyDeploymentSummary`

NewManagedAppPolicyDeploymentSummaryWithDefaults instantiates a new ManagedAppPolicyDeploymentSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigurationDeployedUserCount

`func (o *ManagedAppPolicyDeploymentSummary) GetConfigurationDeployedUserCount() int32`

GetConfigurationDeployedUserCount returns the ConfigurationDeployedUserCount field if non-nil, zero value otherwise.

### GetConfigurationDeployedUserCountOk

`func (o *ManagedAppPolicyDeploymentSummary) GetConfigurationDeployedUserCountOk() (*int32, bool)`

GetConfigurationDeployedUserCountOk returns a tuple with the ConfigurationDeployedUserCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationDeployedUserCount

`func (o *ManagedAppPolicyDeploymentSummary) SetConfigurationDeployedUserCount(v int32)`

SetConfigurationDeployedUserCount sets ConfigurationDeployedUserCount field to given value.

### HasConfigurationDeployedUserCount

`func (o *ManagedAppPolicyDeploymentSummary) HasConfigurationDeployedUserCount() bool`

HasConfigurationDeployedUserCount returns a boolean if a field has been set.

### GetConfigurationDeploymentSummaryPerApp

`func (o *ManagedAppPolicyDeploymentSummary) GetConfigurationDeploymentSummaryPerApp() []*AnyOfmicrosoftGraphManagedAppPolicyDeploymentSummaryPerApp`

GetConfigurationDeploymentSummaryPerApp returns the ConfigurationDeploymentSummaryPerApp field if non-nil, zero value otherwise.

### GetConfigurationDeploymentSummaryPerAppOk

`func (o *ManagedAppPolicyDeploymentSummary) GetConfigurationDeploymentSummaryPerAppOk() (*[]*AnyOfmicrosoftGraphManagedAppPolicyDeploymentSummaryPerApp, bool)`

GetConfigurationDeploymentSummaryPerAppOk returns a tuple with the ConfigurationDeploymentSummaryPerApp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationDeploymentSummaryPerApp

`func (o *ManagedAppPolicyDeploymentSummary) SetConfigurationDeploymentSummaryPerApp(v []*AnyOfmicrosoftGraphManagedAppPolicyDeploymentSummaryPerApp)`

SetConfigurationDeploymentSummaryPerApp sets ConfigurationDeploymentSummaryPerApp field to given value.

### HasConfigurationDeploymentSummaryPerApp

`func (o *ManagedAppPolicyDeploymentSummary) HasConfigurationDeploymentSummaryPerApp() bool`

HasConfigurationDeploymentSummaryPerApp returns a boolean if a field has been set.

### GetDisplayName

`func (o *ManagedAppPolicyDeploymentSummary) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ManagedAppPolicyDeploymentSummary) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ManagedAppPolicyDeploymentSummary) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ManagedAppPolicyDeploymentSummary) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *ManagedAppPolicyDeploymentSummary) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *ManagedAppPolicyDeploymentSummary) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetLastRefreshTime

`func (o *ManagedAppPolicyDeploymentSummary) GetLastRefreshTime() time.Time`

GetLastRefreshTime returns the LastRefreshTime field if non-nil, zero value otherwise.

### GetLastRefreshTimeOk

`func (o *ManagedAppPolicyDeploymentSummary) GetLastRefreshTimeOk() (*time.Time, bool)`

GetLastRefreshTimeOk returns a tuple with the LastRefreshTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRefreshTime

`func (o *ManagedAppPolicyDeploymentSummary) SetLastRefreshTime(v time.Time)`

SetLastRefreshTime sets LastRefreshTime field to given value.

### HasLastRefreshTime

`func (o *ManagedAppPolicyDeploymentSummary) HasLastRefreshTime() bool`

HasLastRefreshTime returns a boolean if a field has been set.

### GetVersion

`func (o *ManagedAppPolicyDeploymentSummary) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ManagedAppPolicyDeploymentSummary) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ManagedAppPolicyDeploymentSummary) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ManagedAppPolicyDeploymentSummary) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *ManagedAppPolicyDeploymentSummary) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *ManagedAppPolicyDeploymentSummary) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


