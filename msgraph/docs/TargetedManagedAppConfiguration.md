# TargetedManagedAppConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeployedAppCount** | Pointer to **int32** | Count of apps to which the current policy is deployed. | [optional] 
**IsAssigned** | Pointer to **bool** | Indicates if the policy is deployed to any inclusion groups or not. | [optional] 
**Apps** | Pointer to [**[]MicrosoftGraphManagedMobileApp**](MicrosoftGraphManagedMobileApp.md) | List of apps to which the policy is deployed. | [optional] 
**Assignments** | Pointer to [**[]MicrosoftGraphTargetedManagedAppPolicyAssignment**](MicrosoftGraphTargetedManagedAppPolicyAssignment.md) | Navigation property to list of inclusion and exclusion groups to which the policy is deployed. | [optional] 
**DeploymentSummary** | Pointer to [**AnyOfmicrosoftGraphManagedAppPolicyDeploymentSummary**](anyOf&lt;microsoft.graph.managedAppPolicyDeploymentSummary&gt;.md) | Navigation property to deployment summary of the configuration. | [optional] 

## Methods

### NewTargetedManagedAppConfiguration

`func NewTargetedManagedAppConfiguration() *TargetedManagedAppConfiguration`

NewTargetedManagedAppConfiguration instantiates a new TargetedManagedAppConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTargetedManagedAppConfigurationWithDefaults

`func NewTargetedManagedAppConfigurationWithDefaults() *TargetedManagedAppConfiguration`

NewTargetedManagedAppConfigurationWithDefaults instantiates a new TargetedManagedAppConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeployedAppCount

`func (o *TargetedManagedAppConfiguration) GetDeployedAppCount() int32`

GetDeployedAppCount returns the DeployedAppCount field if non-nil, zero value otherwise.

### GetDeployedAppCountOk

`func (o *TargetedManagedAppConfiguration) GetDeployedAppCountOk() (*int32, bool)`

GetDeployedAppCountOk returns a tuple with the DeployedAppCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployedAppCount

`func (o *TargetedManagedAppConfiguration) SetDeployedAppCount(v int32)`

SetDeployedAppCount sets DeployedAppCount field to given value.

### HasDeployedAppCount

`func (o *TargetedManagedAppConfiguration) HasDeployedAppCount() bool`

HasDeployedAppCount returns a boolean if a field has been set.

### GetIsAssigned

`func (o *TargetedManagedAppConfiguration) GetIsAssigned() bool`

GetIsAssigned returns the IsAssigned field if non-nil, zero value otherwise.

### GetIsAssignedOk

`func (o *TargetedManagedAppConfiguration) GetIsAssignedOk() (*bool, bool)`

GetIsAssignedOk returns a tuple with the IsAssigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAssigned

`func (o *TargetedManagedAppConfiguration) SetIsAssigned(v bool)`

SetIsAssigned sets IsAssigned field to given value.

### HasIsAssigned

`func (o *TargetedManagedAppConfiguration) HasIsAssigned() bool`

HasIsAssigned returns a boolean if a field has been set.

### GetApps

`func (o *TargetedManagedAppConfiguration) GetApps() []MicrosoftGraphManagedMobileApp`

GetApps returns the Apps field if non-nil, zero value otherwise.

### GetAppsOk

`func (o *TargetedManagedAppConfiguration) GetAppsOk() (*[]MicrosoftGraphManagedMobileApp, bool)`

GetAppsOk returns a tuple with the Apps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApps

`func (o *TargetedManagedAppConfiguration) SetApps(v []MicrosoftGraphManagedMobileApp)`

SetApps sets Apps field to given value.

### HasApps

`func (o *TargetedManagedAppConfiguration) HasApps() bool`

HasApps returns a boolean if a field has been set.

### GetAssignments

`func (o *TargetedManagedAppConfiguration) GetAssignments() []MicrosoftGraphTargetedManagedAppPolicyAssignment`

GetAssignments returns the Assignments field if non-nil, zero value otherwise.

### GetAssignmentsOk

`func (o *TargetedManagedAppConfiguration) GetAssignmentsOk() (*[]MicrosoftGraphTargetedManagedAppPolicyAssignment, bool)`

GetAssignmentsOk returns a tuple with the Assignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignments

`func (o *TargetedManagedAppConfiguration) SetAssignments(v []MicrosoftGraphTargetedManagedAppPolicyAssignment)`

SetAssignments sets Assignments field to given value.

### HasAssignments

`func (o *TargetedManagedAppConfiguration) HasAssignments() bool`

HasAssignments returns a boolean if a field has been set.

### GetDeploymentSummary

`func (o *TargetedManagedAppConfiguration) GetDeploymentSummary() AnyOfmicrosoftGraphManagedAppPolicyDeploymentSummary`

GetDeploymentSummary returns the DeploymentSummary field if non-nil, zero value otherwise.

### GetDeploymentSummaryOk

`func (o *TargetedManagedAppConfiguration) GetDeploymentSummaryOk() (*AnyOfmicrosoftGraphManagedAppPolicyDeploymentSummary, bool)`

GetDeploymentSummaryOk returns a tuple with the DeploymentSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentSummary

`func (o *TargetedManagedAppConfiguration) SetDeploymentSummary(v AnyOfmicrosoftGraphManagedAppPolicyDeploymentSummary)`

SetDeploymentSummary sets DeploymentSummary field to given value.

### HasDeploymentSummary

`func (o *TargetedManagedAppConfiguration) HasDeploymentSummary() bool`

HasDeploymentSummary returns a boolean if a field has been set.

### SetDeploymentSummaryNil

`func (o *TargetedManagedAppConfiguration) SetDeploymentSummaryNil(b bool)`

 SetDeploymentSummaryNil sets the value for DeploymentSummary to be an explicit nil

### UnsetDeploymentSummary
`func (o *TargetedManagedAppConfiguration) UnsetDeploymentSummary()`

UnsetDeploymentSummary ensures that no value is present for DeploymentSummary, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


