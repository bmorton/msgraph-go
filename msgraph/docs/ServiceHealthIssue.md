# ServiceHealthIssue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Classification** | Pointer to [**AnyOfmicrosoftGraphServiceHealthClassificationType**](anyOf&lt;microsoft.graph.serviceHealthClassificationType&gt;.md) | The type of service health issue. Possible values are: advisory, incident, unknownFutureValue. | [optional] 
**Feature** | Pointer to **NullableString** | The feature name of the service issue. | [optional] 
**FeatureGroup** | Pointer to **NullableString** | The feature group name of the service issue. | [optional] 
**ImpactDescription** | Pointer to **string** | The description of the service issue impact. | [optional] 
**IsResolved** | Pointer to **bool** | Indicates whether the issue is resolved. | [optional] 
**Origin** | Pointer to [**AnyOfmicrosoftGraphServiceHealthOrigin**](anyOf&lt;microsoft.graph.serviceHealthOrigin&gt;.md) | Indicates the origin of the service issue. Possible values are: microsoft, thirdParty, customer, unknownFutureValue. | [optional] 
**Posts** | Pointer to [**[]MicrosoftGraphServiceHealthIssuePost**](MicrosoftGraphServiceHealthIssuePost.md) | Collection of historical posts for the service issue. | [optional] 
**Service** | Pointer to **string** | Indicates the service affected by the issue. | [optional] 
**Status** | Pointer to [**AnyOfmicrosoftGraphServiceHealthStatus**](anyOf&lt;microsoft.graph.serviceHealthStatus&gt;.md) | The status of the service issue. Possible values are: serviceOperational, investigating, restoringService, verifyingService, serviceRestored, postIncidentReviewPublished, serviceDegradation, serviceInterruption, extendedRecovery, falsePositive, investigationSuspended, resolved, mitigatedExternal, mitigated, resolvedExternal, confirmed, reported, unknownFutureValue. | [optional] 

## Methods

### NewServiceHealthIssue

`func NewServiceHealthIssue() *ServiceHealthIssue`

NewServiceHealthIssue instantiates a new ServiceHealthIssue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceHealthIssueWithDefaults

`func NewServiceHealthIssueWithDefaults() *ServiceHealthIssue`

NewServiceHealthIssueWithDefaults instantiates a new ServiceHealthIssue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassification

`func (o *ServiceHealthIssue) GetClassification() AnyOfmicrosoftGraphServiceHealthClassificationType`

GetClassification returns the Classification field if non-nil, zero value otherwise.

### GetClassificationOk

`func (o *ServiceHealthIssue) GetClassificationOk() (*AnyOfmicrosoftGraphServiceHealthClassificationType, bool)`

GetClassificationOk returns a tuple with the Classification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassification

`func (o *ServiceHealthIssue) SetClassification(v AnyOfmicrosoftGraphServiceHealthClassificationType)`

SetClassification sets Classification field to given value.

### HasClassification

`func (o *ServiceHealthIssue) HasClassification() bool`

HasClassification returns a boolean if a field has been set.

### SetClassificationNil

`func (o *ServiceHealthIssue) SetClassificationNil(b bool)`

 SetClassificationNil sets the value for Classification to be an explicit nil

### UnsetClassification
`func (o *ServiceHealthIssue) UnsetClassification()`

UnsetClassification ensures that no value is present for Classification, not even an explicit nil
### GetFeature

`func (o *ServiceHealthIssue) GetFeature() string`

GetFeature returns the Feature field if non-nil, zero value otherwise.

### GetFeatureOk

`func (o *ServiceHealthIssue) GetFeatureOk() (*string, bool)`

GetFeatureOk returns a tuple with the Feature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeature

`func (o *ServiceHealthIssue) SetFeature(v string)`

SetFeature sets Feature field to given value.

### HasFeature

`func (o *ServiceHealthIssue) HasFeature() bool`

HasFeature returns a boolean if a field has been set.

### SetFeatureNil

`func (o *ServiceHealthIssue) SetFeatureNil(b bool)`

 SetFeatureNil sets the value for Feature to be an explicit nil

### UnsetFeature
`func (o *ServiceHealthIssue) UnsetFeature()`

UnsetFeature ensures that no value is present for Feature, not even an explicit nil
### GetFeatureGroup

`func (o *ServiceHealthIssue) GetFeatureGroup() string`

GetFeatureGroup returns the FeatureGroup field if non-nil, zero value otherwise.

### GetFeatureGroupOk

`func (o *ServiceHealthIssue) GetFeatureGroupOk() (*string, bool)`

GetFeatureGroupOk returns a tuple with the FeatureGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureGroup

`func (o *ServiceHealthIssue) SetFeatureGroup(v string)`

SetFeatureGroup sets FeatureGroup field to given value.

### HasFeatureGroup

`func (o *ServiceHealthIssue) HasFeatureGroup() bool`

HasFeatureGroup returns a boolean if a field has been set.

### SetFeatureGroupNil

`func (o *ServiceHealthIssue) SetFeatureGroupNil(b bool)`

 SetFeatureGroupNil sets the value for FeatureGroup to be an explicit nil

### UnsetFeatureGroup
`func (o *ServiceHealthIssue) UnsetFeatureGroup()`

UnsetFeatureGroup ensures that no value is present for FeatureGroup, not even an explicit nil
### GetImpactDescription

`func (o *ServiceHealthIssue) GetImpactDescription() string`

GetImpactDescription returns the ImpactDescription field if non-nil, zero value otherwise.

### GetImpactDescriptionOk

`func (o *ServiceHealthIssue) GetImpactDescriptionOk() (*string, bool)`

GetImpactDescriptionOk returns a tuple with the ImpactDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpactDescription

`func (o *ServiceHealthIssue) SetImpactDescription(v string)`

SetImpactDescription sets ImpactDescription field to given value.

### HasImpactDescription

`func (o *ServiceHealthIssue) HasImpactDescription() bool`

HasImpactDescription returns a boolean if a field has been set.

### GetIsResolved

`func (o *ServiceHealthIssue) GetIsResolved() bool`

GetIsResolved returns the IsResolved field if non-nil, zero value otherwise.

### GetIsResolvedOk

`func (o *ServiceHealthIssue) GetIsResolvedOk() (*bool, bool)`

GetIsResolvedOk returns a tuple with the IsResolved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsResolved

`func (o *ServiceHealthIssue) SetIsResolved(v bool)`

SetIsResolved sets IsResolved field to given value.

### HasIsResolved

`func (o *ServiceHealthIssue) HasIsResolved() bool`

HasIsResolved returns a boolean if a field has been set.

### GetOrigin

`func (o *ServiceHealthIssue) GetOrigin() AnyOfmicrosoftGraphServiceHealthOrigin`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *ServiceHealthIssue) GetOriginOk() (*AnyOfmicrosoftGraphServiceHealthOrigin, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *ServiceHealthIssue) SetOrigin(v AnyOfmicrosoftGraphServiceHealthOrigin)`

SetOrigin sets Origin field to given value.

### HasOrigin

`func (o *ServiceHealthIssue) HasOrigin() bool`

HasOrigin returns a boolean if a field has been set.

### SetOriginNil

`func (o *ServiceHealthIssue) SetOriginNil(b bool)`

 SetOriginNil sets the value for Origin to be an explicit nil

### UnsetOrigin
`func (o *ServiceHealthIssue) UnsetOrigin()`

UnsetOrigin ensures that no value is present for Origin, not even an explicit nil
### GetPosts

`func (o *ServiceHealthIssue) GetPosts() []MicrosoftGraphServiceHealthIssuePost`

GetPosts returns the Posts field if non-nil, zero value otherwise.

### GetPostsOk

`func (o *ServiceHealthIssue) GetPostsOk() (*[]MicrosoftGraphServiceHealthIssuePost, bool)`

GetPostsOk returns a tuple with the Posts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosts

`func (o *ServiceHealthIssue) SetPosts(v []MicrosoftGraphServiceHealthIssuePost)`

SetPosts sets Posts field to given value.

### HasPosts

`func (o *ServiceHealthIssue) HasPosts() bool`

HasPosts returns a boolean if a field has been set.

### GetService

`func (o *ServiceHealthIssue) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *ServiceHealthIssue) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *ServiceHealthIssue) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *ServiceHealthIssue) HasService() bool`

HasService returns a boolean if a field has been set.

### GetStatus

`func (o *ServiceHealthIssue) GetStatus() AnyOfmicrosoftGraphServiceHealthStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ServiceHealthIssue) GetStatusOk() (*AnyOfmicrosoftGraphServiceHealthStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ServiceHealthIssue) SetStatus(v AnyOfmicrosoftGraphServiceHealthStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ServiceHealthIssue) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *ServiceHealthIssue) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *ServiceHealthIssue) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


