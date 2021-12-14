# MicrosoftGraphFeatureRolloutPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**Description** | Pointer to **NullableString** | A description for this feature rollout policy. | [optional] 
**DisplayName** | Pointer to **string** | The display name for this  feature rollout policy. | [optional] 
**Feature** | Pointer to [**AnyOfmicrosoftGraphStagedFeatureName**](anyOf&lt;microsoft.graph.stagedFeatureName&gt;.md) | Possible values are: passthroughAuthentication, seamlessSso, passwordHashSync, emailAsAlternateId, unknownFutureValue. | [optional] 
**IsAppliedToOrganization** | Pointer to **bool** | Indicates whether this feature rollout policy should be applied to the entire organization. | [optional] 
**IsEnabled** | Pointer to **bool** | Indicates whether the feature rollout is enabled. | [optional] 
**AppliesTo** | Pointer to [**[]MicrosoftGraphDirectoryObject**](MicrosoftGraphDirectoryObject.md) | Nullable. Specifies a list of directoryObjects that feature is enabled for. | [optional] 

## Methods

### NewMicrosoftGraphFeatureRolloutPolicy

`func NewMicrosoftGraphFeatureRolloutPolicy() *MicrosoftGraphFeatureRolloutPolicy`

NewMicrosoftGraphFeatureRolloutPolicy instantiates a new MicrosoftGraphFeatureRolloutPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphFeatureRolloutPolicyWithDefaults

`func NewMicrosoftGraphFeatureRolloutPolicyWithDefaults() *MicrosoftGraphFeatureRolloutPolicy`

NewMicrosoftGraphFeatureRolloutPolicyWithDefaults instantiates a new MicrosoftGraphFeatureRolloutPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphFeatureRolloutPolicy) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphFeatureRolloutPolicy) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphFeatureRolloutPolicy) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphFeatureRolloutPolicy) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDescription

`func (o *MicrosoftGraphFeatureRolloutPolicy) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MicrosoftGraphFeatureRolloutPolicy) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MicrosoftGraphFeatureRolloutPolicy) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MicrosoftGraphFeatureRolloutPolicy) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *MicrosoftGraphFeatureRolloutPolicy) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *MicrosoftGraphFeatureRolloutPolicy) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDisplayName

`func (o *MicrosoftGraphFeatureRolloutPolicy) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphFeatureRolloutPolicy) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *MicrosoftGraphFeatureRolloutPolicy) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *MicrosoftGraphFeatureRolloutPolicy) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetFeature

`func (o *MicrosoftGraphFeatureRolloutPolicy) GetFeature() AnyOfmicrosoftGraphStagedFeatureName`

GetFeature returns the Feature field if non-nil, zero value otherwise.

### GetFeatureOk

`func (o *MicrosoftGraphFeatureRolloutPolicy) GetFeatureOk() (*AnyOfmicrosoftGraphStagedFeatureName, bool)`

GetFeatureOk returns a tuple with the Feature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeature

`func (o *MicrosoftGraphFeatureRolloutPolicy) SetFeature(v AnyOfmicrosoftGraphStagedFeatureName)`

SetFeature sets Feature field to given value.

### HasFeature

`func (o *MicrosoftGraphFeatureRolloutPolicy) HasFeature() bool`

HasFeature returns a boolean if a field has been set.

### SetFeatureNil

`func (o *MicrosoftGraphFeatureRolloutPolicy) SetFeatureNil(b bool)`

 SetFeatureNil sets the value for Feature to be an explicit nil

### UnsetFeature
`func (o *MicrosoftGraphFeatureRolloutPolicy) UnsetFeature()`

UnsetFeature ensures that no value is present for Feature, not even an explicit nil
### GetIsAppliedToOrganization

`func (o *MicrosoftGraphFeatureRolloutPolicy) GetIsAppliedToOrganization() bool`

GetIsAppliedToOrganization returns the IsAppliedToOrganization field if non-nil, zero value otherwise.

### GetIsAppliedToOrganizationOk

`func (o *MicrosoftGraphFeatureRolloutPolicy) GetIsAppliedToOrganizationOk() (*bool, bool)`

GetIsAppliedToOrganizationOk returns a tuple with the IsAppliedToOrganization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAppliedToOrganization

`func (o *MicrosoftGraphFeatureRolloutPolicy) SetIsAppliedToOrganization(v bool)`

SetIsAppliedToOrganization sets IsAppliedToOrganization field to given value.

### HasIsAppliedToOrganization

`func (o *MicrosoftGraphFeatureRolloutPolicy) HasIsAppliedToOrganization() bool`

HasIsAppliedToOrganization returns a boolean if a field has been set.

### GetIsEnabled

`func (o *MicrosoftGraphFeatureRolloutPolicy) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *MicrosoftGraphFeatureRolloutPolicy) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *MicrosoftGraphFeatureRolloutPolicy) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *MicrosoftGraphFeatureRolloutPolicy) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetAppliesTo

`func (o *MicrosoftGraphFeatureRolloutPolicy) GetAppliesTo() []MicrosoftGraphDirectoryObject`

GetAppliesTo returns the AppliesTo field if non-nil, zero value otherwise.

### GetAppliesToOk

`func (o *MicrosoftGraphFeatureRolloutPolicy) GetAppliesToOk() (*[]MicrosoftGraphDirectoryObject, bool)`

GetAppliesToOk returns a tuple with the AppliesTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliesTo

`func (o *MicrosoftGraphFeatureRolloutPolicy) SetAppliesTo(v []MicrosoftGraphDirectoryObject)`

SetAppliesTo sets AppliesTo field to given value.

### HasAppliesTo

`func (o *MicrosoftGraphFeatureRolloutPolicy) HasAppliesTo() bool`

HasAppliesTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


