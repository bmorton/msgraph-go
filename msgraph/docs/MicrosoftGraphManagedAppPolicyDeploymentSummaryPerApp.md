# MicrosoftGraphManagedAppPolicyDeploymentSummaryPerApp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigurationAppliedUserCount** | Pointer to **int32** | Number of users the policy is applied. | [optional] 
**MobileAppIdentifier** | Pointer to [**AnyOfobject**](anyOf&lt;object&gt;.md) | Deployment of an app. | [optional] 

## Methods

### NewMicrosoftGraphManagedAppPolicyDeploymentSummaryPerApp

`func NewMicrosoftGraphManagedAppPolicyDeploymentSummaryPerApp() *MicrosoftGraphManagedAppPolicyDeploymentSummaryPerApp`

NewMicrosoftGraphManagedAppPolicyDeploymentSummaryPerApp instantiates a new MicrosoftGraphManagedAppPolicyDeploymentSummaryPerApp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphManagedAppPolicyDeploymentSummaryPerAppWithDefaults

`func NewMicrosoftGraphManagedAppPolicyDeploymentSummaryPerAppWithDefaults() *MicrosoftGraphManagedAppPolicyDeploymentSummaryPerApp`

NewMicrosoftGraphManagedAppPolicyDeploymentSummaryPerAppWithDefaults instantiates a new MicrosoftGraphManagedAppPolicyDeploymentSummaryPerApp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigurationAppliedUserCount

`func (o *MicrosoftGraphManagedAppPolicyDeploymentSummaryPerApp) GetConfigurationAppliedUserCount() int32`

GetConfigurationAppliedUserCount returns the ConfigurationAppliedUserCount field if non-nil, zero value otherwise.

### GetConfigurationAppliedUserCountOk

`func (o *MicrosoftGraphManagedAppPolicyDeploymentSummaryPerApp) GetConfigurationAppliedUserCountOk() (*int32, bool)`

GetConfigurationAppliedUserCountOk returns a tuple with the ConfigurationAppliedUserCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationAppliedUserCount

`func (o *MicrosoftGraphManagedAppPolicyDeploymentSummaryPerApp) SetConfigurationAppliedUserCount(v int32)`

SetConfigurationAppliedUserCount sets ConfigurationAppliedUserCount field to given value.

### HasConfigurationAppliedUserCount

`func (o *MicrosoftGraphManagedAppPolicyDeploymentSummaryPerApp) HasConfigurationAppliedUserCount() bool`

HasConfigurationAppliedUserCount returns a boolean if a field has been set.

### GetMobileAppIdentifier

`func (o *MicrosoftGraphManagedAppPolicyDeploymentSummaryPerApp) GetMobileAppIdentifier() AnyOfobject`

GetMobileAppIdentifier returns the MobileAppIdentifier field if non-nil, zero value otherwise.

### GetMobileAppIdentifierOk

`func (o *MicrosoftGraphManagedAppPolicyDeploymentSummaryPerApp) GetMobileAppIdentifierOk() (*AnyOfobject, bool)`

GetMobileAppIdentifierOk returns a tuple with the MobileAppIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobileAppIdentifier

`func (o *MicrosoftGraphManagedAppPolicyDeploymentSummaryPerApp) SetMobileAppIdentifier(v AnyOfobject)`

SetMobileAppIdentifier sets MobileAppIdentifier field to given value.

### HasMobileAppIdentifier

`func (o *MicrosoftGraphManagedAppPolicyDeploymentSummaryPerApp) HasMobileAppIdentifier() bool`

HasMobileAppIdentifier returns a boolean if a field has been set.

### SetMobileAppIdentifierNil

`func (o *MicrosoftGraphManagedAppPolicyDeploymentSummaryPerApp) SetMobileAppIdentifierNil(b bool)`

 SetMobileAppIdentifierNil sets the value for MobileAppIdentifier to be an explicit nil

### UnsetMobileAppIdentifier
`func (o *MicrosoftGraphManagedAppPolicyDeploymentSummaryPerApp) UnsetMobileAppIdentifier()`

UnsetMobileAppIdentifier ensures that no value is present for MobileAppIdentifier, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


