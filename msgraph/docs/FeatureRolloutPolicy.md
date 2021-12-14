# FeatureRolloutPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **NullableString** | A description for this feature rollout policy. | [optional] 
**DisplayName** | Pointer to **string** | The display name for this  feature rollout policy. | [optional] 
**Feature** | Pointer to [**AnyOfmicrosoftGraphStagedFeatureName**](anyOf&lt;microsoft.graph.stagedFeatureName&gt;.md) | Possible values are: passthroughAuthentication, seamlessSso, passwordHashSync, emailAsAlternateId, unknownFutureValue. | [optional] 
**IsAppliedToOrganization** | Pointer to **bool** | Indicates whether this feature rollout policy should be applied to the entire organization. | [optional] 
**IsEnabled** | Pointer to **bool** | Indicates whether the feature rollout is enabled. | [optional] 
**AppliesTo** | Pointer to [**[]MicrosoftGraphDirectoryObject**](MicrosoftGraphDirectoryObject.md) | Nullable. Specifies a list of directoryObjects that feature is enabled for. | [optional] 

## Methods

### NewFeatureRolloutPolicy

`func NewFeatureRolloutPolicy() *FeatureRolloutPolicy`

NewFeatureRolloutPolicy instantiates a new FeatureRolloutPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeatureRolloutPolicyWithDefaults

`func NewFeatureRolloutPolicyWithDefaults() *FeatureRolloutPolicy`

NewFeatureRolloutPolicyWithDefaults instantiates a new FeatureRolloutPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *FeatureRolloutPolicy) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FeatureRolloutPolicy) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FeatureRolloutPolicy) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FeatureRolloutPolicy) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *FeatureRolloutPolicy) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *FeatureRolloutPolicy) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDisplayName

`func (o *FeatureRolloutPolicy) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *FeatureRolloutPolicy) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *FeatureRolloutPolicy) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *FeatureRolloutPolicy) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetFeature

`func (o *FeatureRolloutPolicy) GetFeature() AnyOfmicrosoftGraphStagedFeatureName`

GetFeature returns the Feature field if non-nil, zero value otherwise.

### GetFeatureOk

`func (o *FeatureRolloutPolicy) GetFeatureOk() (*AnyOfmicrosoftGraphStagedFeatureName, bool)`

GetFeatureOk returns a tuple with the Feature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeature

`func (o *FeatureRolloutPolicy) SetFeature(v AnyOfmicrosoftGraphStagedFeatureName)`

SetFeature sets Feature field to given value.

### HasFeature

`func (o *FeatureRolloutPolicy) HasFeature() bool`

HasFeature returns a boolean if a field has been set.

### SetFeatureNil

`func (o *FeatureRolloutPolicy) SetFeatureNil(b bool)`

 SetFeatureNil sets the value for Feature to be an explicit nil

### UnsetFeature
`func (o *FeatureRolloutPolicy) UnsetFeature()`

UnsetFeature ensures that no value is present for Feature, not even an explicit nil
### GetIsAppliedToOrganization

`func (o *FeatureRolloutPolicy) GetIsAppliedToOrganization() bool`

GetIsAppliedToOrganization returns the IsAppliedToOrganization field if non-nil, zero value otherwise.

### GetIsAppliedToOrganizationOk

`func (o *FeatureRolloutPolicy) GetIsAppliedToOrganizationOk() (*bool, bool)`

GetIsAppliedToOrganizationOk returns a tuple with the IsAppliedToOrganization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAppliedToOrganization

`func (o *FeatureRolloutPolicy) SetIsAppliedToOrganization(v bool)`

SetIsAppliedToOrganization sets IsAppliedToOrganization field to given value.

### HasIsAppliedToOrganization

`func (o *FeatureRolloutPolicy) HasIsAppliedToOrganization() bool`

HasIsAppliedToOrganization returns a boolean if a field has been set.

### GetIsEnabled

`func (o *FeatureRolloutPolicy) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *FeatureRolloutPolicy) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *FeatureRolloutPolicy) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *FeatureRolloutPolicy) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetAppliesTo

`func (o *FeatureRolloutPolicy) GetAppliesTo() []MicrosoftGraphDirectoryObject`

GetAppliesTo returns the AppliesTo field if non-nil, zero value otherwise.

### GetAppliesToOk

`func (o *FeatureRolloutPolicy) GetAppliesToOk() (*[]MicrosoftGraphDirectoryObject, bool)`

GetAppliesToOk returns a tuple with the AppliesTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliesTo

`func (o *FeatureRolloutPolicy) SetAppliesTo(v []MicrosoftGraphDirectoryObject)`

SetAppliesTo sets AppliesTo field to given value.

### HasAppliesTo

`func (o *FeatureRolloutPolicy) HasAppliesTo() bool`

HasAppliesTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


